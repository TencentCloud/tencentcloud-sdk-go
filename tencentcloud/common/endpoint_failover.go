package common

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

// EndpointFailover manages per-host circuit breakers for endpoint failover.
type EndpointFailover struct {
	mu          sync.Mutex
	breakers    map[string]*circuitBreaker
	tldFamilies [][]string
}

var defaultEndpointFailover = &EndpointFailover{
	breakers: map[string]*circuitBreaker{},
	tldFamilies: [][]string{
		{"ai.tencentcloudapi.com", "ai.tencentcloudapi.cn", "ai.tencentcloudapi.com.cn"},
		{"internal.tencentcloudapi.com", "internal.tencentcloudapi.cn", "internal.tencentcloudapi.com.cn"},
		{"tencentcloudapi.com", "tencentcloudapi.cn", "tencentcloudapi.com.cn"},
	},
}

type tldMatch struct {
	family     []string
	familyIdx  int
	tldIdx     int
	fullPrefix string
}

type candidate struct {
	host    string
	breaker *circuitBreaker
	gen     uint64
}

// breakerFor returns the circuit breaker for host, creating one if needed.
func (f *EndpointFailover) breakerFor(host string) *circuitBreaker {
	f.mu.Lock()
	defer f.mu.Unlock()
	if b, ok := f.breakers[host]; ok {
		return b
	}
	b := defaultRegionBreaker()
	f.breakers[host] = b
	return b
}

// ----------------------------------------------------------------------
// Host classification
// ----------------------------------------------------------------------

func (f *EndpointFailover) tldMatchOf(host string) *tldMatch {
	if host == "" {
		return nil
	}
	for fi, family := range f.tldFamilies {
		for ti, suffix := range family {
			if !strings.HasSuffix(host, "."+suffix) {
				continue
			}
			prefix := host[:len(host)-len(suffix)-1]
			if prefix == "" || f.hasEmptyLabel(prefix) {
				return nil
			}
			return &tldMatch{
				family:     family,
				familyIdx:  fi,
				tldIdx:     ti,
				fullPrefix: prefix,
			}
		}
	}
	return nil
}

func (f *EndpointFailover) hasEmptyLabel(prefix string) bool {
	if strings.HasPrefix(prefix, ".") || strings.HasSuffix(prefix, ".") {
		return true
	}
	return strings.Contains(prefix, "..")
}

func (f *EndpointFailover) serviceOf(host string) string {
	if i := strings.IndexByte(host, '.'); i >= 0 {
		return host[:i]
	}
	return host
}

// ----------------------------------------------------------------------
// Candidate selection
// ----------------------------------------------------------------------

func (f *EndpointFailover) selectHost(req *http.Request, c *Client) candidate {
	urlHost := req.URL.Host
	headerHost := req.Host

	candidateHosts := f.buildCandidateHosts(urlHost, headerHost, c.profile, req.Context())
	// Track the urlHost breaker result from the first iteration so we can
	// reuse it in the fallback without calling beforeRequest a second time.
	var urlCand candidate
	for i, host := range candidateHosts {
		b := f.breakerFor(host)
		gen, err := b.beforeRequest()
		if i == 0 {
			urlCand = candidate{host: host, breaker: b, gen: gen}
		}
		if err == nil {
			return candidate{host: host, breaker: b, gen: gen}
		}
	}
	// All breakers open: fall back to urlHost using the gen already recorded.
	return urlCand
}

// candidateFor selects a candidate host based on profile and originHost,
// checking circuit breakers. Used by tests and for direct host selection.
func (f *EndpointFailover) candidateFor(profile *profile.ClientProfile, originHost string) *candidate {
	if backupEP := f.backupEndpointOf(profile); backupEP != "" {
		b := f.breakerFor(originHost)
		gen, err := b.beforeRequest()
		if err == nil {
			return &candidate{host: originHost, breaker: b, gen: gen}
		}
		backupHost := f.serviceOf(originHost) + "." + backupEP
		b = f.breakerFor(backupHost)
		gen, err = b.beforeRequest()
		if err == nil {
			return &candidate{host: backupHost, breaker: b, gen: gen}
		}
		return nil
	}

	m := f.tldMatchOf(originHost)
	if m == nil {
		return nil
	}

	for i := 0; i < len(m.family); i++ {
		t := (m.tldIdx + i) % len(m.family)
		host := m.fullPrefix + "." + m.family[t]
		b := f.breakerFor(host)
		gen, err := b.beforeRequest()
		if err == nil {
			return &candidate{host: host, breaker: b, gen: gen}
		}
	}
	return nil
}

func (f *EndpointFailover) buildCandidateHosts(urlHost, headerHost string, p *profile.ClientProfile, ctx context.Context) []string {
	hosts := []string{urlHost}

	// 1. backup endpoint
	if backupEp := f.backupEndpointOf(p); backupEp != "" && headerHost != "" {
		hosts = append(hosts, f.serviceOf(headerHost)+"."+backupEp)
		return hosts
	}

	// 2. TLD-based failover
	m := f.tldMatchOf(urlHost)
	if m == nil {
		return hosts
	}
	for i := 1; i < len(m.family); i++ {
		t := (m.tldIdx + i) % len(m.family)
		hosts = append(hosts, m.fullPrefix+"."+m.family[t])
	}
	return hosts
}

func (f *EndpointFailover) backupEndpointOf(p *profile.ClientProfile) string {
	if p == nil {
		return ""
	}
	if p.BackupEndpoint != "" {
		return p.BackupEndpoint
	}
	return p.BackupEndPoint
}

// ----------------------------------------------------------------------
// Execution
// ----------------------------------------------------------------------

func (f *EndpointFailover) send(req *http.Request, c *Client) (*http.Response, error) {
	cand := f.selectHost(req, c)

	urlHost := req.URL.Host
	if cand.host != urlHost {
		req = f.resignForHost(req, urlHost, cand.host, c.credential, c.signMethod, c.unsignedPayload)
	}
	resp, err := c.sendHttp(req)
	healthy := f.isResponseHealthy(resp, err)
	cand.breaker.afterRequest(cand.gen, healthy)
	return resp, err
}

func (f *EndpointFailover) isResponseHealthy(resp *http.Response, err error) bool {
	if err != nil || resp == nil || resp.StatusCode != 200 {
		return false
	}
	if !f.isJsonContentType(resp) {
		return true
	}
	if resp.Body == nil {
		return false
	}
	bodyBytes, readErr := io.ReadAll(resp.Body)
	resp.Body.Close()
	if readErr != nil {
		return false
	}
	if !f.isValidJson(bodyBytes) {
		return false
	}
	resp.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	return true
}

func (f *EndpointFailover) isJsonContentType(resp *http.Response) bool {
	ct := resp.Header.Get("Content-Type")
	return strings.Contains(strings.ToLower(ct), "application/json")
}

func (f *EndpointFailover) isValidJson(data []byte) bool {
	return json.Valid(data)
}

// ----------------------------------------------------------------------
// Request re-signing
// ----------------------------------------------------------------------

func (f *EndpointFailover) resignForHost(orig *http.Request, originHost, targetHost string, cred CredentialIface, signMethod string, unsignedPayload bool) *http.Request {
	if signMethod == "TC3-HMAC-SHA256" {
		return f.resignV3(orig, targetHost, cred, unsignedPayload)
	}
	if signMethod == "HmacSHA1" || signMethod == "HmacSHA256" {
		return f.resignV1(orig, originHost, targetHost, cred, signMethod)
	}
	// SKIP or unknown: just rewrite host.
	return f.rewriteHost(orig, targetHost)
}

func (f *EndpointFailover) rewriteHost(orig *http.Request, targetHost string) *http.Request {
	req := orig.Clone(orig.Context())
	req.Host = targetHost
	req.URL.Host = targetHost
	return req
}

func (f *EndpointFailover) resignV3(orig *http.Request, targetHost string, cred CredentialIface, unsignedPayload bool) *http.Request {
	secId, secKey, token := cred.GetCredential()
	bodyBytes := f.readBody(orig)
	contentType := orig.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/x-www-form-urlencoded"
	}

	canonicalURI := orig.URL.EscapedPath()
	if canonicalURI == "" {
		canonicalURI = "/"
	}
	canonicalQueryString := ""
	if orig.Method == "GET" {
		canonicalQueryString = orig.URL.RawQuery
	}
	canonicalHeaders := fmt.Sprintf("content-type:%s\nhost:%s\n", contentType, targetHost)
	signedHeaders := "content-type;host"

	hashedPayload := sha256hex(string(bodyBytes))
	if unsignedPayload {
		hashedPayload = sha256hex("UNSIGNED-PAYLOAD")
	}

	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		orig.Method, canonicalURI, canonicalQueryString,
		canonicalHeaders, signedHeaders, hashedPayload)

	now := time.Now()
	timestamp := fmt.Sprintf("%d", now.Unix())
	date := now.UTC().Format("2006-01-02")
	service := strings.SplitN(targetHost, ".", 2)[0]
	credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, service)
	hashedCanonicalRequest := sha256hex(canonicalRequest)

	stringToSign := fmt.Sprintf("TC3-HMAC-SHA256\n%s\n%s\n%s",
		timestamp, credentialScope, hashedCanonicalRequest)

	secretDate := hmacsha256(date, "TC3"+secKey)
	secretService := hmacsha256(service, secretDate)
	secretSigning := hmacsha256("tc3_request", secretService)
	signature := hex.EncodeToString([]byte(hmacsha256(stringToSign, secretSigning)))

	authorization := fmt.Sprintf("TC3-HMAC-SHA256 Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		secId, credentialScope, signedHeaders, signature)

	req := orig.Clone(orig.Context())
	req.Host = targetHost
	req.URL.Host = targetHost
	req.Header.Set("Host", targetHost)
	req.Header.Set("Authorization", authorization)
	req.Header.Set("X-TC-Timestamp", timestamp)
	if token != "" {
		req.Header.Set("X-TC-Token", token)
	} else {
		req.Header.Del("X-TC-Token")
	}
	return req
}

func (f *EndpointFailover) resignV1(orig *http.Request, originHost, targetHost string, cred CredentialIface, signMethod string) *http.Request {
	secId, secKey, token := cred.GetCredential()

	var params map[string]string
	if orig.Method == "GET" {
		params = f.parseQueryParams(orig.URL.RawQuery)
	} else {
		params = f.parseFormParams(string(f.readBody(orig)))
	}
	delete(params, "Signature")
	params["SecretId"] = secId
	if token != "" {
		params["Token"] = token
	} else {
		delete(params, "Token")
	}
	params["SignatureMethod"] = signMethod

	sortedKeys := make([]string, 0, len(params))
	for k := range params {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	var plainBuf bytes.Buffer
	plainBuf.WriteString(orig.Method)
	plainBuf.WriteString(targetHost)
	plainBuf.WriteString(orig.URL.EscapedPath())
	plainBuf.WriteString("?")
	for _, k := range sortedKeys {
		plainBuf.WriteString(k)
		plainBuf.WriteString("=")
		plainBuf.WriteString(params[k])
		plainBuf.WriteString("&")
	}
	plainBuf.Truncate(plainBuf.Len() - 1)

	sig := Sign(plainBuf.String(), secKey, signMethod)
	params["Signature"] = sig

	if orig.Method == "GET" {
		var qs []string
		for _, k := range sortedKeys {
			qs = append(qs, k+"="+params[k])
		}
		qs = append(qs, "Signature="+sig)
		newURL := *orig.URL
		newURL.Host = targetHost
		newURL.RawQuery = strings.Join(qs, "&")
		req := orig.Clone(orig.Context())
		req.URL = &newURL
		req.Host = targetHost
		return req
	}

	var bodyBuf bytes.Buffer
	for _, k := range sortedKeys {
		bodyBuf.WriteString(k)
		bodyBuf.WriteString("=")
		bodyBuf.WriteString(params[k])
		bodyBuf.WriteString("&")
	}
	bodyBuf.WriteString("Signature=")
	bodyBuf.WriteString(sig)

	newURL := *orig.URL
	newURL.Host = targetHost
	req := orig.Clone(orig.Context())
	req.URL = &newURL
	req.Host = targetHost
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Body = io.NopCloser(bytes.NewReader(bodyBuf.Bytes()))
	req.ContentLength = int64(bodyBuf.Len())
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(bodyBuf.Bytes())), nil
	}
	return req
}

func (f *EndpointFailover) readBody(req *http.Request) []byte {
	if req.Body == nil {
		return nil
	}
	b, _ := io.ReadAll(req.Body)
	req.Body.Close()
	req.Body = io.NopCloser(bytes.NewReader(b))
	return b
}

func (f *EndpointFailover) parseQueryParams(rawQuery string) map[string]string {
	m := map[string]string{}
	if rawQuery == "" {
		return m
	}
	for _, pair := range strings.Split(rawQuery, "&") {
		kv := strings.SplitN(pair, "=", 2)
		k := kv[0]
		v := ""
		if len(kv) == 2 {
			v = kv[1]
		}
		m[k] = v
	}
	return m
}

func (f *EndpointFailover) parseFormParams(body string) map[string]string {
	m := map[string]string{}
	for _, pair := range strings.Split(body, "&") {
		kv := strings.SplitN(pair, "=", 2)
		k := kv[0]
		v := ""
		if len(kv) == 2 {
			v = kv[1]
		}
		m[k] = v
	}
	return m
}

// ----------------------------------------------------------------------
// Client integration
// ----------------------------------------------------------------------

func (c *Client) sendWithEndpointFailover(req *http.Request) (*http.Response, error) {
	if c.profile.DisableRegionBreaker {
		return c.sendHttp(req)
	}
	return defaultEndpointFailover.send(req, c)
}

// isBreakerSuccess is used by circuit_breaker_test.go.
func (f *EndpointFailover) isBreakerSuccess(err error) bool {
	if err == nil {
		return true
	}
	e, ok := err.(*tcerr.TencentCloudSDKError)
	if !ok {
		return false
	}
	if e.GetRequestId() == "" {
		return false
	}
	if e.GetCode() == "InternalError" {
		return false
	}
	return true
}
