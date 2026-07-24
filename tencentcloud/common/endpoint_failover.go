package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

type endpointFailover struct {
	mu          sync.Mutex
	breakers    map[string]*circuitBreaker
	tldFamilies [][]string
}

var defaultEndpointFailover = &endpointFailover{
	breakers: map[string]*circuitBreaker{},
	tldFamilies: [][]string{
		{"ai.tencentcloudapi.com", "ai.tencentcloudapi.com.cn", "ai.tencentcloudapi.cn"},
		{"internal.tencentcloudapi.com", "internal.tencentcloudapi.com.cn", "internal.tencentcloudapi.cn"},
		{"tencentcloudapi.com", "tencentcloudapi.com.cn", "tencentcloudapi.cn"},
	},
}

type tldMatch struct {
	family []string
	tldIdx int
}

type candidate struct {
	host    string
	breaker *circuitBreaker
	gen     uint64
}

func (f *endpointFailover) breakerFor(host string) *circuitBreaker {
	f.mu.Lock()
	defer f.mu.Unlock()
	if b, ok := f.breakers[host]; ok {
		return b
	}
	b := defaultRegionBreaker()
	f.breakers[host] = b
	return b
}

// suffix alone (e.g. "tencentcloudapi.com") is not a valid endpoint.
func (f *endpointFailover) tldMatchOf(host string) *tldMatch {
	if host == "" {
		return nil
	}
	for _, family := range f.tldFamilies {
		for ti, suffix := range family {
			if !strings.HasSuffix(host, "."+suffix) {
				continue
			}
			prefix := host[:len(host)-len(suffix)-1]
			if prefix == "" || f.hasEmptyLabel(prefix) {
				return nil
			}
			return &tldMatch{
				family: family,
				tldIdx: ti,
			}
		}
	}
	return nil
}

func (f *endpointFailover) hasEmptyLabel(prefix string) bool {
	if strings.HasPrefix(prefix, ".") || strings.HasSuffix(prefix, ".") {
		return true
	}
	return strings.Contains(prefix, "..")
}

func (f *endpointFailover) serviceOf(host string) string {
	if i := strings.IndexByte(host, '.'); i >= 0 {
		return host[:i]
	}
	return host
}

func (f *endpointFailover) selectHost(req *http.Request, c *Client) candidate {
	urlHost := req.URL.Host

	// fast path: avoid building candidate list when origin is healthy.
	originB := f.breakerFor(urlHost)
	originGen, originErr := originB.beforeRequest()
	if originErr == nil {
		return candidate{host: urlHost, breaker: originB, gen: originGen}
	}

	// service comes from the first label of the origin host — e.g.
	// "cvm.ap-guangzhou.tencentcloudapi.com" → "cvm". This keeps the
	// failover host's service prefix aligned with what the origin was
	// actually called as, rather than rb.service which is the API
	// service name and may diverge from the host prefix.
	service := f.serviceOf(urlHost)

	candidateHosts := f.buildCandidateHosts(urlHost, service, c.profile)
	for _, host := range candidateHosts {
		b := f.breakerFor(host)
		gen, err := b.beforeRequest()
		if err == nil {
			return candidate{host: host, breaker: b, gen: gen}
		}
	}
	// all breakers open: fall back to origin.
	return candidate{host: urlHost, breaker: originB, gen: originGen}
}

func (f *endpointFailover) buildCandidateHosts(urlHost, service string, p *profile.ClientProfile) []string {
	if backupEp := f.backupEndpointOf(p); backupEp != "" {
		if service == "" {
			return nil
		}
		return []string{service + "." + backupEp}
	}

	if service == "" {
		return nil
	}
	m := f.tldMatchOf(urlHost)
	if m == nil {
		return nil
	}
	hosts := make([]string, 0, len(m.family)-1)
	for i := 1; i < len(m.family); i++ {
		t := (m.tldIdx + i) % len(m.family)
		hosts = append(hosts, service+"."+m.family[t])
	}
	return hosts
}

func (f *endpointFailover) backupEndpointOf(p *profile.ClientProfile) string {
	if p.BackupEndpoint != "" {
		return p.BackupEndpoint
	}
	return p.BackupEndPoint
}

func (f *endpointFailover) send(req *http.Request, c *Client) (*http.Response, error) {
	rb := requestBuilderFromContext(req.Context())
	if rb == nil {
		return nil, tcerr.NewTencentCloudSDKError("ClientError.EndpointFailoverError", "requestBuilder missing from context", "")
	}

	cand := f.selectHost(req, c)

	urlHost := req.URL.Host
	if cand.host != urlHost {
		var err error
		if req, err = f.resignForHost(req, cand.host, rb); err != nil {
			return nil, err
		}
	}
	resp, err := c.sendHttp(req)
	healthy := f.isResponseHealthy(resp, err)
	cand.breaker.afterRequest(cand.gen, healthy)
	return resp, err
}

// JSON body must be peeked to detect truncated/malformed responses that HTTP
// 200 still carries — tencentcloud returns 200 with an error body.
func (f *endpointFailover) isResponseHealthy(resp *http.Response, err error) bool {
	if err != nil || resp == nil || resp.StatusCode != 200 {
		return false
	}
	if !f.isJsonContentType(resp) {
		return true
	}
	if resp.Body == nil {
		return false
	}
	bodyBytes, readErr := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if readErr != nil {
		return false
	}
	if !f.isValidJson(bodyBytes) {
		return false
	}
	resp.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
	return true
}

func (f *endpointFailover) isJsonContentType(resp *http.Response) bool {
	ct := resp.Header.Get("Content-Type")
	return strings.Contains(strings.ToLower(ct), "application/json")
}

func (f *endpointFailover) isValidJson(data []byte) bool {
	return json.Valid(data)
}

// resignForHost rebuilds orig against targetHost via requestBuilder.
// Scheme is forced to https — failover endpoints are https-only.
func (f *endpointFailover) resignForHost(orig *http.Request, targetHost string, rb *requestBuilder) (*http.Request, error) {
	rb.url.Host = targetHost
	rb.url.Scheme = tchttp.HTTPS
	rb.host = targetHost
	rb.ctx = orig.Context()
	return rb.build()
}

func (c *Client) sendWithEndpointFailover(req *http.Request) (*http.Response, error) {
	if c.profile.DisableRegionBreaker {
		return c.sendHttp(req)
	}
	return defaultEndpointFailover.send(req, c)
}
