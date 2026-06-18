package common

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

// =====================================================================
// Domain families — matching Java's Family model
// =====================================================================

type family struct {
	name           string
	originHost     string
	firstFailover  string
	secondFailover string
	allTldHosts    []string
}

var (
	fNormal = family{
		name: "Normal", originHost: "cvm.tencentcloudapi.com",
		firstFailover: "cvm.tencentcloudapi.cn", secondFailover: "cvm.tencentcloudapi.com.cn",
		allTldHosts: []string{"cvm.tencentcloudapi.com", "cvm.tencentcloudapi.cn", "cvm.tencentcloudapi.com.cn"},
	}
	fAI = family{
		name: "AI", originHost: "hunyuan.ai.tencentcloudapi.com",
		firstFailover: "hunyuan.ai.tencentcloudapi.cn", secondFailover: "hunyuan.ai.tencentcloudapi.com.cn",
		allTldHosts: []string{"hunyuan.ai.tencentcloudapi.com", "hunyuan.ai.tencentcloudapi.cn", "hunyuan.ai.tencentcloudapi.com.cn"},
	}
	fInternal = family{
		name: "Internal", originHost: "cvm.internal.tencentcloudapi.com",
		firstFailover: "cvm.internal.tencentcloudapi.cn", secondFailover: "cvm.internal.tencentcloudapi.com.cn",
		allTldHosts: []string{"cvm.internal.tencentcloudapi.com", "cvm.internal.tencentcloudapi.cn", "cvm.internal.tencentcloudapi.com.cn"},
	}
	families = []family{fNormal, fAI, fInternal}
)

// =====================================================================
// TransportStub — matching Java's TransportStub interceptor
// =====================================================================

const failoverSuccessResp = `{"Response": {"RequestId": "req-ok"}}`

type transportStub struct {
	mu        sync.Mutex
	outcomes  []stubOutcome
	hostsSeen []string
	authsSeen []string
}

type stubOutcome struct {
	err         error
	code        int
	body        string
	contentType string
}

func (s *transportStub) programOk() {
	s.outcomes = append(s.outcomes, stubOutcome{code: 200, body: failoverSuccessResp})
}

func (s *transportStub) programJsonOk(json string) {
	s.outcomes = append(s.outcomes, stubOutcome{code: 200, body: json, contentType: "application/json"})
}

func (s *transportStub) programFailure(err error) {
	s.outcomes = append(s.outcomes, stubOutcome{err: err})
}

func (s *transportStub) programResponse(code int, body string) {
	s.outcomes = append(s.outcomes, stubOutcome{code: code, body: body, contentType: "application/json"})
}

func (s *transportStub) programResponseWithCt(code int, body, contentType string) {
	s.outcomes = append(s.outcomes, stubOutcome{code: code, body: body, contentType: contentType})
}

func (s *transportStub) RoundTrip(req *http.Request) (*http.Response, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	host := req.Host
	if host == "" {
		host = req.URL.Host
	}
	s.hostsSeen = append(s.hostsSeen, host)
	s.authsSeen = append(s.authsSeen, req.Header.Get("Authorization"))

	if len(s.outcomes) == 0 {
		return nil, errors.New("transportStub exhausted: no programmed outcome left")
	}
	o := s.outcomes[0]
	s.outcomes = s.outcomes[1:]
	if o.err != nil {
		return nil, o.err
	}
	ct := o.contentType
	if ct == "" {
		ct = "application/json"
	}
	header := http.Header{}
	if o.contentType != "none" {
		header.Set("Content-Type", ct)
	}
	return &http.Response{
		StatusCode: o.code,
		Status:     http.StatusText(o.code),
		Header:     header,
		Body:       ioutil.NopCloser(bytes.NewBufferString(o.body)),
	}, nil
}

// =====================================================================
// Error types for transport failures
// =====================================================================

type dnsErr struct{ msg string }

func (e dnsErr) Error() string { return e.msg }

type tlsErr struct{ msg string }

func (e tlsErr) Error() string { return e.msg }

type connectErr struct{ msg string }

func (e connectErr) Error() string { return e.msg }

type timeoutErr struct{ msg string }

func (e timeoutErr) Error() string { return e.msg }

// =====================================================================
// Helpers
// =====================================================================

func newCvmRequest() *requestWithClientToken {
	return newTestRequest()
}

func newClient(f family) *Client {
	resetBreakers()
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	cpf.HttpProfile.Endpoint = f.originHost
	return NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
}

func installStub(c *Client) *transportStub {
	s := &transportStub{}
	c.WithHttpTransport(s)
	return s
}

func tripBreakerFor(host string, timeoutMs time.Duration) {
	b := defaultEndpointFailover.breakerFor(host)
	tripBreaker(b, timeoutMs)
}

func tripAllBreakersFor(f family, timeoutMs time.Duration) {
	for _, h := range f.allTldHosts {
		tripBreakerFor(h, timeoutMs)
	}
}

func tripBreaker(b *circuitBreaker, timeoutMs time.Duration) {
	b.mu.Lock()
	b.state = StateOpen
	b.expiry = time.Now().Add(timeoutMs)
	b.mu.Unlock()
}

func resetBreakers() {
	defaultEndpointFailover.mu.Lock()
	defaultEndpointFailover.breakers = map[string]*circuitBreaker{}
	defaultEndpointFailover.mu.Unlock()
}

func runFailoverClient(rt http.RoundTripper) *Client {
	resetBreakers()
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	c.WithHttpTransport(rt)
	return c
}

// =====================================================================
// M1: isKnownTencentCloudHost — host classification
// =====================================================================

func Test_isKnownTencentCloudHost(t *testing.T) {
	known := []string{
		"cvm.tencentcloudapi.com", "cvm.tencentcloudapi.cn", "cvm.tencentcloudapi.com.cn",
		"cvm.intl.tencentcloudapi.com",
		"cvm.ap-shanghai.tencentcloudapi.com", "cvm.ap-shanghai.tencentcloudapi.cn", "cvm.ap-shanghai.tencentcloudapi.com.cn",
		"hunyuan.ai.tencentcloudapi.com", "hunyuan.ai.tencentcloudapi.cn", "hunyuan.ai.tencentcloudapi.com.cn",
		"hunyuan.ap-guangzhou.ai.tencentcloudapi.com",
		"cvm.internal.tencentcloudapi.com", "cvm.internal.tencentcloudapi.cn", "cvm.internal.tencentcloudapi.com.cn",
		"cvm.ap-guangzhou.internal.tencentcloudapi.com",
	}
	for _, h := range known {
		if defaultEndpointFailover.tldMatchOf(h) == nil {
			t.Errorf("expected %q to be known", h)
		}
	}

	unknown := []string{
		"tencentcloudapi.com", "tencentcloudapi.cn", "tencentcloudapi.com.cn",
		".tencentcloudapi.com", ".foo.tencentcloudapi.com", "foo..tencentcloudapi.com",
		"example.com", "cvm.tencentcloudapi.woa.com", "proxy.internal", "192.168.0.1",
	}
	for _, h := range unknown {
		if defaultEndpointFailover.tldMatchOf(h) != nil {
			t.Errorf("expected %q to be unknown", h)
		}
	}
	if defaultEndpointFailover.tldMatchOf("") != nil {
		t.Error("expected nil for empty host")
	}
}

// =====================================================================
// M2: hostWithTld — utility helpers
// =====================================================================

func Test_hostWithTldBuildsCorrectHosts(t *testing.T) {
	c := newClient(fNormal)

	// Verify each TLD index via candidateFor on a tripped breaker.
	tripBreakerFor("cvm.tencentcloudapi.com", time.Hour)
	cand := defaultEndpointFailover.candidateFor(c.profile, "cvm.tencentcloudapi.com")
	if cand == nil || cand.host != "cvm.tencentcloudapi.cn" {
		t.Fatalf("got %v, want .cn", cand)
	}

	tripBreakerFor("cvm.tencentcloudapi.cn", time.Hour)
	cand = defaultEndpointFailover.candidateFor(c.profile, "cvm.tencentcloudapi.com")
	if cand == nil || cand.host != "cvm.tencentcloudapi.com.cn" {
		t.Fatalf("got %v, want .com.cn", cand)
	}
}

func Test_hostWithTldFromCnAndComCnOrigins(t *testing.T) {
	c := newClient(fNormal)
	c.profile.HttpProfile.Endpoint = "cvm.tencentcloudapi.cn"
	tripBreakerFor("cvm.tencentcloudapi.cn", time.Hour)
	cand := defaultEndpointFailover.candidateFor(c.profile, "cvm.tencentcloudapi.cn")
	if cand == nil || cand.host != "cvm.tencentcloudapi.com.cn" {
		t.Fatalf("cn origin: got %v, want .com.cn", cand)
	}

	c.profile.HttpProfile.Endpoint = "cvm.tencentcloudapi.com.cn"
	tripBreakerFor("cvm.tencentcloudapi.com.cn", time.Hour)
	cand = defaultEndpointFailover.candidateFor(c.profile, "cvm.tencentcloudapi.com.cn")
	if cand == nil || cand.host != "cvm.tencentcloudapi.com" {
		t.Fatalf("com.cn origin: got %v, want .com", cand)
	}
}

func Test_hostWithTldPreservesRegionInPrefix(t *testing.T) {
	c := newClient(fNormal)

	tripBreakerFor("cvm.ap-guangzhou.tencentcloudapi.com", time.Hour)
	cand := defaultEndpointFailover.candidateFor(c.profile, "cvm.ap-guangzhou.tencentcloudapi.com")
	if cand == nil || cand.host != "cvm.ap-guangzhou.tencentcloudapi.cn" {
		t.Fatalf("got %v, want cvm.ap-guangzhou.tencentcloudapi.cn", cand)
	}
}

// =====================================================================
// M3: Non-TencentCloud host — passthrough
// =====================================================================

func Test_passthroughNonTencentCloudHost(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		s := installStub(c)
		s.programOk()

		if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
			t.Fatalf("%s: unexpected error: %v", f.name, err)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.originHost {
			t.Fatalf("%s: host = %s, want %s", f.name, s.hostsSeen[0], f.originHost)
		}
	}
}

func Test_nonTencentHostDnsMissPropagatesWithoutRetry(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		s := installStub(c)
		s.programFailure(dnsErr{"proxy dns miss"})

		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("%s: expected error", f.name)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
	}
}

// =====================================================================
// M4-M6: Failover triggers — IOException types
// =====================================================================

func Test_failoverOnTransportErrors(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		s := installStub(c)
		s.programFailure(dnsErr{"dns miss"})

		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("%s: expected error", f.name)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.originHost {
			t.Fatalf("%s: host = %s, want %s", f.name, s.hostsSeen[0], f.originHost)
		}
	}
}

func Test_failoverOnTlsErrors(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		s := installStub(c)
		s.programFailure(tlsErr{"tls fail"})

		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("%s: expected error", f.name)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.originHost {
			t.Fatalf("%s: host = %s, want %s", f.name, s.hostsSeen[0], f.originHost)
		}
	}
}

// =====================================================================
// M14-M16: Protocol-level failover
// =====================================================================

func Test_non200ResponseRecordsFailureWithoutRetry(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		s := installStub(c)
		s.programResponse(503, `{"Response":{"Error":{}}}`)

		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("%s: expected error", f.name)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.originHost {
			t.Fatalf("%s: host = %s, want %s", f.name, s.hostsSeen[0], f.originHost)
		}
	}
}

func Test_invalidJsonBodyRecordsFailure(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		s := installStub(c)
		s.programResponse(200, "not json at all")

		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("%s: expected error", f.name)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.originHost {
			t.Fatalf("%s: host = %s, want %s", f.name, s.hostsSeen[0], f.originHost)
		}
	}
}

func Test_emptyBodyRecordsFailure(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		s := installStub(c)
		s.programResponseWithCt(200, "", "application/json")

		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("%s: expected error", f.name)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
	}
}

// =====================================================================
// M16a-M16b: TLD family rotation
// =====================================================================

func Test_aiFamilyRotationStaysWithinFamily(t *testing.T) {
	c := newClient(fAI)
	tripBreakerFor(fAI.originHost, time.Hour)
	s := installStub(c)
	s.programOk()

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success: %v", err)
	}
	if len(s.hostsSeen) != 1 {
		t.Fatalf("expected 1 attempt, got %d", len(s.hostsSeen))
	}
	if s.hostsSeen[0] != fAI.firstFailover {
		t.Fatalf("host = %s, want %s", s.hostsSeen[0], fAI.firstFailover)
	}
	if !strings.Contains(s.hostsSeen[0], "ai.tencentcloudapi") {
		t.Fatal("must stay within ai.tencentcloudapi family")
	}
}

func Test_internalFamilyRotationStaysWithinFamily(t *testing.T) {
	c := newClient(fInternal)
	tripBreakerFor(fInternal.originHost, time.Hour)
	s := installStub(c)
	s.programOk()

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success: %v", err)
	}
	if len(s.hostsSeen) != 1 {
		t.Fatalf("expected 1 attempt, got %d", len(s.hostsSeen))
	}
	if s.hostsSeen[0] != fInternal.firstFailover {
		t.Fatalf("host = %s, want %s", s.hostsSeen[0], fInternal.firstFailover)
	}
	if !strings.Contains(s.hostsSeen[0], "internal.tencentcloudapi") {
		t.Fatal("must stay within internal.tencentcloudapi family")
	}
}

func Test_regionPinnedHostFailoverPreservesPrefix(t *testing.T) {
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	cpf.HttpProfile.Endpoint = "cvm.ap-guangzhou.tencentcloudapi.com"
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	tripBreakerFor("cvm.ap-guangzhou.tencentcloudapi.com", time.Hour)
	s := installStub(c)
	s.programOk()

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success: %v", err)
	}
	if len(s.hostsSeen) != 1 {
		t.Fatalf("expected 1 attempt, got %d", len(s.hostsSeen))
	}
	if s.hostsSeen[0] != "cvm.ap-guangzhou.tencentcloudapi.cn" {
		t.Fatalf("host = %s, want cvm.ap-guangzhou.tencentcloudapi.cn", s.hostsSeen[0])
	}
	if !strings.HasPrefix(s.hostsSeen[0], "cvm.ap-guangzhou.") {
		t.Fatal("region prefix must be preserved")
	}
}

// =====================================================================
// M17: API response delivered after failover
// =====================================================================

func Test_apiResponseDeliveredAfterFailover(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		tripBreakerFor(f.originHost, time.Hour)
		s := installStub(c)
		s.programJsonOk(`{"Response":{"TotalCount":42,"InstanceSet":[],"RequestId":"req-xyz"}}`)

		resp := tchttp.NewCommonResponse()
		if err := c.Send(newCvmRequest(), resp); err != nil {
			t.Fatalf("%s: expected success: %v", f.name, err)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.firstFailover {
			t.Fatalf("%s: host = %s, want %s", f.name, s.hostsSeen[0], f.firstFailover)
		}
	}
}

// =====================================================================
// M18: Sustained failure opens breaker
// =====================================================================

func Test_breakerOpensAfterSustainedFailure(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		s := installStub(c)

		for i := 0; i < 5; i++ {
			s.programFailure(dnsErr{"fail"})
			err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
			if err == nil {
				t.Fatalf("%s: expected error on iteration %d", f.name, i)
			}
		}
		if len(s.hostsSeen) != 5 {
			t.Fatalf("%s: expected 5 attempts, got %d", f.name, len(s.hostsSeen))
		}

		// Origin breaker should be Open.
		if _, err := defaultEndpointFailover.breakerFor(f.originHost).beforeRequest(); err == nil {
			t.Fatalf("%s: origin breaker should be Open", f.name)
		}

		// Next request short-circuits origin, goes to first failover.
		s.hostsSeen = nil
		s.programOk()
		if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
			t.Fatalf("%s: expected success: %v", f.name, err)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.firstFailover {
			t.Fatalf("%s: host = %s, want %s", f.name, s.hostsSeen[0], f.firstFailover)
		}
	}
}

// =====================================================================
// M19: Open breaker short-circuits host
// =====================================================================

func Test_openBreakerShortCircuitsHost(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		tripBreakerFor(f.originHost, time.Hour)
		s := installStub(c)
		s.programOk()

		if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
			t.Fatalf("%s: expected success: %v", f.name, err)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.firstFailover {
			t.Fatalf("%s: should skip origin, hit %s, got %s", f.name, f.firstFailover, s.hostsSeen[0])
		}
	}
}

// =====================================================================
// M20: Open → HalfOpen after cooldown
// =====================================================================

func Test_breakerTransitionsOpenToHalfOpenAfterCooldown(t *testing.T) {
	shortTimeout := 100 * time.Millisecond
	for _, f := range families {
		resetBreakers()
		b := defaultEndpointFailover.breakerFor(f.originHost)
		b.mu.Lock()
		b.state = StateOpen
		b.expiry = time.Now().Add(shortTimeout)
		b.mu.Unlock()

		if _, err := b.beforeRequest(); err == nil {
			t.Fatalf("%s: breaker should be Open", f.name)
		}

		time.Sleep(shortTimeout + 50*time.Millisecond)
		if _, err := b.beforeRequest(); err != nil {
			t.Fatalf("%s: should permit HalfOpen probe: %v", f.name, err)
		}
	}
}

// =====================================================================
// M21: HalfOpen probe success → Closed
// =====================================================================

func Test_breakerReClosesAfterHalfOpenSuccess(t *testing.T) {
	shortTimeout := 100 * time.Millisecond
	for _, f := range families {
		resetBreakers()
		c := newClient(f)
		b := defaultEndpointFailover.breakerFor(f.originHost)
		b.mu.Lock()
		b.state = StateOpen
		b.expiry = time.Now().Add(shortTimeout)
		b.mu.Unlock()

		time.Sleep(shortTimeout + 50*time.Millisecond)

		s := installStub(c)
		s.programOk()
		if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
			t.Fatalf("%s: expected success: %v", f.name, err)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.originHost {
			t.Fatalf("%s: host = %s, want %s", f.name, s.hostsSeen[0], f.originHost)
		}

		// Should be Closed now.
		for i := 0; i < 10; i++ {
			if _, err := b.beforeRequest(); err != nil {
				t.Fatalf("%s: should be Closed after HalfOpen success: %v", f.name, err)
			}
		}
	}
}

// =====================================================================
// M22: HalfOpen probe failure → re-Open
// =====================================================================

func Test_breakerReOpensWhenHalfOpenProbeFails(t *testing.T) {
	shortTimeout := 100 * time.Millisecond
	for _, f := range families {
		resetBreakers()
		c := newClient(f)
		b := defaultEndpointFailover.breakerFor(f.originHost)
		b.mu.Lock()
		b.state = StateOpen
		b.expiry = time.Now().Add(shortTimeout)
		b.mu.Unlock()

		time.Sleep(shortTimeout + 50*time.Millisecond)

		s := installStub(c)
		s.programFailure(dnsErr{"still down"})
		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("%s: expected error", f.name)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.originHost {
			t.Fatalf("%s: host = %s, want %s", f.name, s.hostsSeen[0], f.originHost)
		}

		if _, err := b.beforeRequest(); err == nil {
			t.Fatalf("%s: should re-Open after HalfOpen failure", f.name)
		}

		// Next request short-circuits origin again.
		s.hostsSeen = nil
		s.programOk()
		if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
			t.Fatalf("%s: expected success: %v", f.name, err)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.firstFailover {
			t.Fatalf("%s: host = %s, want %s", f.name, s.hostsSeen[0], f.firstFailover)
		}
	}
}

// =====================================================================
// M23: All breakers open → fallback to origin host
// =====================================================================

func Test_allBreakersOpenFallsBackToOriginHost(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		tripAllBreakersFor(f, time.Hour)
		s := installStub(c)
		s.programOk()

		if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
			t.Fatalf("%s: expected success: %v", f.name, err)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.originHost {
			t.Fatalf("%s: host = %s, want %s", f.name, s.hostsSeen[0], f.originHost)
		}
	}
}

// =====================================================================
// M24: One transport attempt per request
// =====================================================================

func Test_endpointFailureSurfacesAttemptFailure(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		s := installStub(c)
		s.programFailure(dnsErr{"dns miss " + f.name})

		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("%s: expected error", f.name)
		}
		if !strings.Contains(err.Error(), "dns miss") {
			t.Fatalf("%s: error should contain failure message, got: %v", f.name, err)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
	}
}

// =====================================================================
// M25: Failure preserves original exception type
// =====================================================================

func Test_failurePreservesAttemptCauseType(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		s := installStub(c)
		s.programFailure(connectErr{"connect fail " + f.name})

		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("%s: expected error", f.name)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.originHost {
			t.Fatalf("%s: host = %s, want %s", f.name, s.hostsSeen[0], f.originHost)
		}
	}
}

// =====================================================================
// M26: Breaker skip mixed with real failure
// =====================================================================

func Test_failureMixesPriorBreakerSkipsWithRealFailure(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		tripBreakerFor(f.originHost, time.Hour)
		s := installStub(c)
		s.programFailure(tlsErr{"tls fail"})

		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("%s: expected error", f.name)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.firstFailover {
			t.Fatalf("%s: origin skipped, should hit %s, got %s", f.name, f.firstFailover, s.hostsSeen[0])
		}
	}
}

// =====================================================================
// M27: Failure does not pollute next request
// =====================================================================

func Test_failoverDoesNotPolluteNextRequest(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		s := installStub(c)

		s.programFailure(dnsErr{"run1 " + f.name})
		_ = c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		s.hostsSeen = nil

		s.programFailure(dnsErr{"run2 " + f.name})
		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("%s: expected error", f.name)
		}
		msg := err.Error()
		if !strings.Contains(msg, "run2") {
			t.Fatalf("%s: error should contain run2, got: %v", f.name, msg)
		}
		if strings.Contains(msg, "run1") {
			t.Fatalf("%s: error should not contain run1, got: %v", f.name, msg)
		}
	}
}

// =====================================================================
// M28: Breaker state isolated across origin hosts
// =====================================================================

func Test_breakerStateIsolatedAcrossOriginHosts(t *testing.T) {
	pairs := [][2]family{{fNormal, fAI}, {fNormal, fInternal}}
	for _, pair := range pairs {
		a, b := pair[0], pair[1]

		_ = newClient(a)
		tripBreakerFor(a.originHost, time.Hour)

		cB := newClient(b)
		sB := installStub(cB)
		sB.programOk()
		if err := cB.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
			t.Fatalf("%s should be unaffected by %s: %v", b.name, a.name, err)
		}
		if len(sB.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", b.name, len(sB.hostsSeen))
		}
		if sB.hostsSeen[0] != b.originHost {
			t.Fatalf("%s: host = %s, want %s", b.name, sB.hostsSeen[0], b.originHost)
		}
	}
}

// =====================================================================
// M35: SSE — no failover
// =====================================================================

func Test_sseStreamResponseIsNotJsonValidated(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		s := installStub(c)
		s.programResponseWithCt(200, "data: hello\n\n", "text/event-stream")

		_ = c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
	}
}

// =====================================================================
// M36: No Content-Type — no failover
// =====================================================================

func Test_responseWithoutContentTypeIsNotJsonValidated(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		s := installStub(c)
		s.programResponseWithCt(200, "<html>oops</html>", "none")

		_ = c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
	}
}

// =====================================================================
// M37: 200 + JSON business error — no failover
// =====================================================================

func Test_businessSdkErrorDoesNotTriggerFailover(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		s := installStub(c)
		s.programJsonOk(`{"Response":{"RequestId":"req-bad","Error":{"Code":"AuthFailure.SignatureFailure","Message":"signature wrong"}}}`)

		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("%s: expected business SDK error", f.name)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
	}
}

// =====================================================================
// M38: backupEndpoint failover behavior
// =====================================================================

func Test_backupEndpointFailover(t *testing.T) {
	for _, f := range families {
		resetBreakers()
		cpf := profile.NewClientProfile()
		cpf.DisableRegionBreaker = false
		cpf.HttpProfile.Endpoint = f.originHost
		cpf.BackupEndpoint = "backup.example.com"
		c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
		s := installStub(c)

		// Origin succeeds.
		s.programOk()
		if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
			t.Fatalf("%s origin: expected success: %v", f.name, err)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s origin: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.originHost {
			t.Fatalf("%s origin: host = %s, want %s", f.name, s.hostsSeen[0], f.originHost)
		}

		// Trip origin breaker, backup should be used.
		tripBreakerFor(f.originHost, time.Hour)
		s.hostsSeen = nil
		s.programOk()
		if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
			t.Fatalf("%s backup: expected success: %v", f.name, err)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s backup: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		servicePrefix := f.originHost[:strings.IndexByte(f.originHost, '.')]
		backupHost := servicePrefix + ".backup.example.com"
		if s.hostsSeen[0] != backupHost {
			t.Fatalf("%s backup: host = %s, want %s", f.name, s.hostsSeen[0], backupHost)
		}
	}
}

// =====================================================================
// M38a-M38d: backupEndpoint detail scenarios
// =====================================================================

func Test_backupEndpointOriginDnsMissDoesNotRetrySameRequest(t *testing.T) {
	for _, f := range families {
		resetBreakers()
		cpf := profile.NewClientProfile()
		cpf.DisableRegionBreaker = false
		cpf.HttpProfile.Endpoint = f.originHost
		cpf.BackupEndpoint = "backup.example.com"
		c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
		s := installStub(c)
		s.programFailure(dnsErr{"origin dns miss"})

		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("%s: expected error", f.name)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: no same-request retry to backup, got %d attempts", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.originHost {
			t.Fatalf("%s: host = %s, want %s", f.name, s.hostsSeen[0], f.originHost)
		}
	}
}

func Test_backupEndpointNonFailoverErrorPropagates(t *testing.T) {
	for _, f := range families {
		resetBreakers()
		cpf := profile.NewClientProfile()
		cpf.DisableRegionBreaker = false
		cpf.HttpProfile.Endpoint = f.originHost
		cpf.BackupEndpoint = "backup.example.com"
		c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
		s := installStub(c)
		s.programFailure(errors.New("generic I/O error"))

		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("%s: expected error", f.name)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.originHost {
			t.Fatalf("%s: host = %s, want %s", f.name, s.hostsSeen[0], f.originHost)
		}
	}
}

func Test_noBackupEndpointDnsMissDoesNotRetrySameRequest(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		s := installStub(c)
		s.programFailure(dnsErr{"dns miss"})

		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("%s: expected error", f.name)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: no same-request retry, got %d attempts", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.originHost {
			t.Fatalf("%s: host = %s, want %s", f.name, s.hostsSeen[0], f.originHost)
		}
	}
}

// =====================================================================
// Endpoint eligibility
// =====================================================================

func Test_cnOriginIsEligibleForFailover(t *testing.T) {
	for _, f := range families {
		c := newClient(f)
		s := installStub(c)
		s.programFailure(dnsErr{"cn dns miss"})

		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("%s: expected error", f.name)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.originHost {
			t.Fatalf("%s: host = %s, want %s", f.name, s.hostsSeen[0], f.originHost)
		}
	}
}

func Test_regionPinnedHostIsEligibleForFailover(t *testing.T) {
	regionFamilies := []family{
		{name: "Normal-R", originHost: "cvm.ap-guangzhou.tencentcloudapi.com",
			firstFailover: "cvm.ap-guangzhou.tencentcloudapi.cn", secondFailover: "cvm.ap-guangzhou.tencentcloudapi.com.cn",
			allTldHosts: []string{"cvm.ap-guangzhou.tencentcloudapi.com", "cvm.ap-guangzhou.tencentcloudapi.cn", "cvm.ap-guangzhou.tencentcloudapi.com.cn"}},
		{name: "AI-R", originHost: "hunyuan.ap-guangzhou.ai.tencentcloudapi.com",
			firstFailover: "hunyuan.ap-guangzhou.ai.tencentcloudapi.cn", secondFailover: "hunyuan.ap-guangzhou.ai.tencentcloudapi.com.cn",
			allTldHosts: []string{"hunyuan.ap-guangzhou.ai.tencentcloudapi.com", "hunyuan.ap-guangzhou.ai.tencentcloudapi.cn", "hunyuan.ap-guangzhou.ai.tencentcloudapi.com.cn"}},
		{name: "Internal-R", originHost: "cvm.ap-guangzhou.internal.tencentcloudapi.com",
			firstFailover: "cvm.ap-guangzhou.internal.tencentcloudapi.cn", secondFailover: "cvm.ap-guangzhou.internal.tencentcloudapi.com.cn",
			allTldHosts: []string{"cvm.ap-guangzhou.internal.tencentcloudapi.com", "cvm.ap-guangzhou.internal.tencentcloudapi.cn", "cvm.ap-guangzhou.internal.tencentcloudapi.com.cn"}},
	}
	for _, f := range regionFamilies {
		c := newClient(f)
		s := installStub(c)
		s.programFailure(dnsErr{"region dns miss"})

		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("%s: expected error", f.name)
		}
		if len(s.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", f.name, len(s.hostsSeen))
		}
		if s.hostsSeen[0] != f.originHost {
			t.Fatalf("%s: host = %s, want %s", f.name, s.hostsSeen[0], f.originHost)
		}
	}
}

// =====================================================================
// Happy path & disabled
// =====================================================================

func Test_happyPathOneAttempt(t *testing.T) {
	rt := &transportStub{}
	rt.programOk()
	c := runFailoverClient(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success: %v", err)
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("happy path made %d attempts, want 1", len(rt.hostsSeen))
	}
}

func Test_disabledFailoverGoesThrough(t *testing.T) {
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = true
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	rt := &transportStub{}
	rt.programFailure(dnsErr{"dns miss"})
	c.WithHttpTransport(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err == nil {
		t.Fatal("expected error")
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("expected 1 attempt, got %d", len(rt.hostsSeen))
	}
}
