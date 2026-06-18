package common

import (
	"bytes"
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

const defaultHalfOpenMaxRequests = 1

// tldFamilies enumerates the Tencent Cloud TLD families for TLD-rotation mode.
// Families are ordered more-specific first so matchFamily returns the first hit.
var tldFamilies = [][]string{
	{"ai.tencentcloudapi.com", "ai.tencentcloudapi.cn", "ai.tencentcloudapi.com.cn"},
	{"internal.tencentcloudapi.com", "internal.tencentcloudapi.cn", "internal.tencentcloudapi.com.cn"},
	{"tencentcloudapi.com", "tencentcloudapi.cn", "tencentcloudapi.com.cn"},
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

// EndpointFailover manages per-host circuit breakers for endpoint failover.
// A single instance is shared across all Client instances via the package-level
// DefaultEndpointFailover variable.
type EndpointFailover struct {
	mu                   sync.Mutex
	breakers             map[string]*circuitBreaker
	DisableRegionBreaker bool
}

var DefaultEndpointFailover = &EndpointFailover{
	breakers: map[string]*circuitBreaker{},
}

// breakerFor returns the circuit breaker for host, creating one if needed.
func (f *EndpointFailover) breakerFor(host string) *circuitBreaker {
	f.mu.Lock()
	defer f.mu.Unlock()
	if b, ok := f.breakers[host]; ok {
		return b
	}
	b := newRegionBreaker(breakerSetting{
		maxFailNum:        defaultMaxFailNum,
		maxFailPercentage: defaultMaxFailPercentage,
		windowInterval:    defaultWindowLength,
		timeout:           defaultTimeout,
		maxRequests:       defaultHalfOpenMaxRequests,
	})
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
	for fi, family := range tldFamilies {
		for ti, suffix := range family {
			if !strings.HasSuffix(host, "."+suffix) {
				continue
			}
			prefix := host[:len(host)-len(suffix)-1]
			if prefix == "" || hasEmptyLabel(prefix) {
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

func hasEmptyLabel(prefix string) bool {
	if strings.HasPrefix(prefix, ".") || strings.HasSuffix(prefix, ".") {
		return true
	}
	return strings.Contains(prefix, "..")
}

func serviceOf(host string) string {
	if i := strings.IndexByte(host, '.'); i >= 0 {
		return host[:i]
	}
	return host
}

// ----------------------------------------------------------------------
// Candidate selection
// ----------------------------------------------------------------------

func (f *EndpointFailover) candidateFor(profile *profile.ClientProfile, originHost string) *candidate {
	if backupEP := f.backupEndpointOf(profile); backupEP != "" {
		b := f.breakerFor(originHost)
		gen, err := b.beforeRequest()
		if err == nil {
			return &candidate{host: originHost, breaker: b, gen: gen}
		}
		backupHost := serviceOf(originHost) + "." + backupEP
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

func (f *EndpointFailover) send(req *http.Request, profile *profile.ClientProfile, cred CredentialIface, signMethod string, unsignedPayload bool, sendHttp func(*http.Request) (*http.Response, error)) (*http.Response, error) {
	if f.DisableRegionBreaker {
		return sendHttp(req)
	}

	originHost := req.Host
	if originHost == "" {
		originHost = req.URL.Host
	}

	cand := f.candidateFor(profile, originHost)
	if cand == nil {
		return sendHttp(req)
	}

	if cand.host != originHost {
		req = f.resignForHost(req, originHost, cand.host, cred, signMethod, unsignedPayload)
	} else {
		req.Host = cand.host
		req.URL.Host = cand.host
	}
	resp, err := sendHttp(req)
	if !f.isResponseHealthy(resp, err) {
		cand.breaker.afterRequest(cand.gen, false)
		return resp, err
	}
	cand.breaker.afterRequest(cand.gen, true)

	buf, bodyErr := f.bufferResponseBody(resp)
	if bodyErr != nil {
		cand.breaker.afterRequest(cand.gen, false)
		return nil, bodyErr
	}
	return buf, nil
}

func (f *EndpointFailover) isResponseHealthy(resp *http.Response, err error) bool {
	if err != nil || resp == nil {
		return false
	}
	return resp.StatusCode == 200
}

func (f *EndpointFailover) bufferResponseBody(resp *http.Response) (*http.Response, error) {
	if resp.Body == nil {
		return nil, fmt.Errorf("response has no body")
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	if f.isJsonContentType(resp) && !f.isValidJson(bodyBytes) {
		return nil, fmt.Errorf("response body is not valid JSON")
	}
	resp.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	return resp, nil
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

	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	t, _ := time.Parse("2006-01-02", time.Now().UTC().Format("2006-01-02"))
	date := t.Format("2006-01-02")
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
	fo := DefaultEndpointFailover
	fo.DisableRegionBreaker = c.profile.DisableRegionBreaker
	return fo.send(req, c.profile, c.credential, c.signMethod, c.unsignedPayload, c.sendHttp)
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
