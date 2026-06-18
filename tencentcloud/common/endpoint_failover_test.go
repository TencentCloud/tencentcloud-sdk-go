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
// Pure host classification
// =====================================================================

func Test_tldMatchOf_plainCom(t *testing.T) {
	m := DefaultEndpointFailover.tldMatchOf("cvm.tencentcloudapi.com")
	if m == nil {
		t.Fatal("expected non-nil match")
	}
	if m.familyIdx != 2 || m.tldIdx != 0 || m.fullPrefix != "cvm" {
		t.Fatalf("got %+v, want familyIdx=2 tldIdx=0 fullPrefix=cvm", m)
	}
}

func Test_tldMatchOf_plainCn(t *testing.T) {
	m := DefaultEndpointFailover.tldMatchOf("cvm.tencentcloudapi.cn")
	if m == nil || m.tldIdx != 1 {
		t.Fatalf("got %+v, want tldIdx=1", m)
	}
}

func Test_tldMatchOf_plainComCn(t *testing.T) {
	m := DefaultEndpointFailover.tldMatchOf("cvm.tencentcloudapi.com.cn")
	if m == nil || m.tldIdx != 2 {
		t.Fatalf("got %+v, want tldIdx=2", m)
	}
}

func Test_tldMatchOf_aiFamily(t *testing.T) {
	m := DefaultEndpointFailover.tldMatchOf("hunyuan.ai.tencentcloudapi.com")
	if m == nil || m.familyIdx != 0 || m.tldIdx != 0 || m.fullPrefix != "hunyuan" {
		t.Fatalf("got %+v, want familyIdx=0 tldIdx=0 fullPrefix=hunyuan", m)
	}
}

func Test_tldMatchOf_internalFamily(t *testing.T) {
	m := DefaultEndpointFailover.tldMatchOf("cvm.internal.tencentcloudapi.com")
	if m == nil || m.familyIdx != 1 || m.fullPrefix != "cvm" {
		t.Fatalf("got %+v, want familyIdx=1 fullPrefix=cvm", m)
	}
}

func Test_tldMatchOf_regionPinned(t *testing.T) {
	m := DefaultEndpointFailover.tldMatchOf("cvm.ap-guangzhou.tencentcloudapi.com")
	if m == nil {
		t.Fatal("expected non-nil match for region-pinned host")
	}
	// fullPrefix must preserve the region label
	if m.fullPrefix != "cvm.ap-guangzhou" {
		t.Fatalf("got fullPrefix=%q, want cvm.ap-guangzhou", m.fullPrefix)
	}
	if m.tldIdx != 0 {
		t.Fatalf("got tldIdx=%d, want 0 (.com)", m.tldIdx)
	}
}

func Test_tldMatchOf_unknownHost(t *testing.T) {
	if m := DefaultEndpointFailover.tldMatchOf("example.com"); m != nil {
		t.Fatalf("expected nil for example.com, got %+v", m)
	}
}

func Test_tldMatchOf_emptyHost(t *testing.T) {
	if m := DefaultEndpointFailover.tldMatchOf(""); m != nil {
		t.Fatalf("expected nil for empty host, got %+v", m)
	}
}

func Test_tldMatchOf_woaSuffixIsUnknown(t *testing.T) {
	if m := DefaultEndpointFailover.tldMatchOf("cvm.tencentcloudapi.woa.com"); m != nil {
		t.Fatalf("expected nil for non-Tencent suffix, got %+v", m)
	}
}

func Test_tldMatchOf_doubleDotIsRejected(t *testing.T) {
	// Malformed hostnames with ".." in the prefix must not match.
	if m := DefaultEndpointFailover.tldMatchOf("cvm..tencentcloudapi.com"); m != nil {
		t.Fatalf("expected nil for double-dot host, got %+v", m)
	}
}

// =====================================================================
// Plan generation
// =====================================================================

func newFailoverEnabledClient() *Client {
	resetBreakers()
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	return NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
}

func Test_candidateFor_backupEndpointOriginAllowed(t *testing.T) {
	c := newFailoverEnabledClient()
	c.profile.BackupEndpoint = "ap-guangzhou.tencentcloudapi.com"

	cand := DefaultEndpointFailover.candidateFor(c.profile, "cvm.ap-shanghai.tencentcloudapi.com")
	if cand == nil {
		t.Fatal("expected origin candidate when breakers are healthy")
	}
	if cand.host != "cvm.ap-shanghai.tencentcloudapi.com" {
		t.Fatalf("got %s, want origin", cand.host)
	}
}

func Test_candidateFor_backupEndpointOriginOpenReturnsBackup(t *testing.T) {
	c := newFailoverEnabledClient()
	c.profile.BackupEndpoint = "ap-guangzhou.tencentcloudapi.com"
	tripBreakerForTest(DefaultEndpointFailover.breakerFor("cvm.ap-shanghai.tencentcloudapi.com"))

	cand := DefaultEndpointFailover.candidateFor(c.profile, "cvm.ap-shanghai.tencentcloudapi.com")
	if cand == nil {
		t.Fatal("expected backup candidate when origin breaker is open")
	}
	if cand.host != "cvm.ap-guangzhou.tencentcloudapi.com" {
		t.Fatalf("got %s, want backup", cand.host)
	}
}

func Test_candidateFor_plainHostReturnsOrigin(t *testing.T) {
	c := newFailoverEnabledClient()
	cand := DefaultEndpointFailover.candidateFor(c.profile, "cvm.tencentcloudapi.com")
	if cand == nil {
		t.Fatal("expected origin candidate")
	}
	if cand.host != "cvm.tencentcloudapi.com" {
		t.Fatalf("got %s, want cvm.tencentcloudapi.com", cand.host)
	}
}

func Test_candidateFor_aiFamilyStaysInFamily(t *testing.T) {
	c := newFailoverEnabledClient()
	cand := DefaultEndpointFailover.candidateFor(c.profile, "hunyuan.ai.tencentcloudapi.com")
	if cand == nil {
		t.Fatal("expected candidate in ai. family")
	}
	if !strings.Contains(cand.host, ".ai.") {
		t.Fatalf("candidate %s leaked outside ai. family", cand.host)
	}
}

func Test_candidateFor_internalFamilyStaysInFamily(t *testing.T) {
	c := newFailoverEnabledClient()
	cand := DefaultEndpointFailover.candidateFor(c.profile, "cvm.internal.tencentcloudapi.com")
	if cand == nil {
		t.Fatal("expected candidate in internal. family")
	}
	if !strings.Contains(cand.host, ".internal.") {
		t.Fatalf("candidate %s leaked outside internal. family", cand.host)
	}
}

func Test_candidateFor_regionPinnedOrigin(t *testing.T) {
	c := newFailoverEnabledClient()
	cand := DefaultEndpointFailover.candidateFor(c.profile, "cvm.ap-guangzhou.tencentcloudapi.com")
	if cand == nil {
		t.Fatal("expected candidate for region-pinned host")
	}
	if cand.host != "cvm.ap-guangzhou.tencentcloudapi.com" {
		t.Fatalf("got %s, want origin", cand.host)
	}
}

func Test_candidateFor_regionPinnedAiFamily(t *testing.T) {
	c := newFailoverEnabledClient()
	cand := DefaultEndpointFailover.candidateFor(c.profile, "hunyuan.ap-guangzhou.ai.tencentcloudapi.com")
	if cand == nil {
		t.Fatal("expected candidate for region-pinned ai host")
	}
	if !strings.Contains(cand.host, ".ai.") {
		t.Fatalf("candidate %s leaked outside ai. family", cand.host)
	}
}

func Test_candidateFor_unknownHostReturnsNil(t *testing.T) {
	c := newFailoverEnabledClient()
	if cand := DefaultEndpointFailover.candidateFor(c.profile, "example.com"); cand != nil {
		t.Fatalf("unknown host should return nil, got %+v", cand)
	}
}

// =====================================================================
// End-to-end with mock RoundTripper
// =====================================================================

const failoverSuccessResp = `{"Response": {"RequestId": "req-ok"}}`

// scriptedRT is a programmable RoundTripper. Each call pops the next
// outcome from the queue and either returns a programmed response or fails
// with a programmed error. Records every request's URL host and the
// Authorization header for inspection.
type scriptedRT struct {
	mu        sync.Mutex
	outcomes  []rtOutcome
	hostsSeen []string
	authsSeen []string
}

type rtOutcome struct {
	err         error
	code        int
	body        string
	contentType string // empty → use default "application/json"
}

func (r *scriptedRT) programOk() {
	r.outcomes = append(r.outcomes, rtOutcome{code: 200, body: failoverSuccessResp})
}

func (r *scriptedRT) programResponse(code int, body string) {
	r.outcomes = append(r.outcomes, rtOutcome{code: code, body: body})
}

func (r *scriptedRT) programResponseWithCt(code int, body, contentType string) {
	r.outcomes = append(r.outcomes, rtOutcome{code: code, body: body, contentType: contentType})
}

func (r *scriptedRT) programErr(err error) {
	r.outcomes = append(r.outcomes, rtOutcome{err: err})
}

func (r *scriptedRT) RoundTrip(req *http.Request) (*http.Response, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	host := req.Host
	if host == "" {
		host = req.URL.Host
	}
	r.hostsSeen = append(r.hostsSeen, host)
	r.authsSeen = append(r.authsSeen, req.Header.Get("Authorization"))

	if len(r.outcomes) == 0 {
		return nil, errors.New("scriptedRT exhausted: no programmed outcome left")
	}
	o := r.outcomes[0]
	r.outcomes = r.outcomes[1:]
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

// dnsErr satisfies isBreakerSuccess(false): a non-SDK error has no RequestId
// so it counts as a breaker failure and triggers the next candidate.
type dnsErr struct{ msg string }

func (e dnsErr) Error() string { return e.msg }

type tlsErr struct{ msg string }

func (e tlsErr) Error() string { return e.msg }

type connectErr struct{ msg string }

func (e connectErr) Error() string { return e.msg }

type timeoutErr struct{ msg string }

func (e timeoutErr) Error() string { return e.msg }

func newCvmRequest() *requestWithClientToken {
	return newTestRequest()
}

func runFailoverClient(rt http.RoundTripper) *Client {
	resetBreakers()
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	c.WithHttpTransport(rt)
	return c
}

// DNS miss → error, breaker records failure. Next request picks .cn.
func Test_failover_dnsMissOnComRecordsFailureNextRequestUsesCn(t *testing.T) {
	rt := &scriptedRT{}
	c := runFailoverClient(rt)

	// 5 failures to open the .com breaker.
	for i := 0; i < 5; i++ {
		rt.programErr(dnsErr{"dns miss"})
		_ = c.Send(newCvmRequest(), tchttp.NewCommonResponse())
	}

	// 6th request: .com open → .cn.
	rt.programOk()
	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("6th request should succeed on .cn: %v", err)
	}
	if rt.hostsSeen[5] != "cvm.tencentcloudapi.cn" {
		t.Fatalf("6th host = %s, want .cn", rt.hostsSeen[5])
	}
}

// Cyclic rotation: .cn origin, breaker open → next request picks .com.cn
func Test_failover_cyclicRotationFromCn(t *testing.T) {
	resetBreakers()
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.cn"
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)

	tripBreakerForTest(DefaultEndpointFailover.breakerFor("cvm.tencentcloudapi.cn"))

	rt := &scriptedRT{}
	rt.programOk()
	c.WithHttpTransport(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success on .com.cn: %v", err)
	}
	if rt.hostsSeen[0] != "cvm.tencentcloudapi.com.cn" {
		t.Fatalf("host = %s, want .com.cn", rt.hostsSeen[0])
	}
}

// Cyclic rotation: .com.cn origin, breaker open → next request picks .com
func Test_failover_cyclicRotationFromComCn(t *testing.T) {
	resetBreakers()
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com.cn"
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)

	tripBreakerForTest(DefaultEndpointFailover.breakerFor("cvm.tencentcloudapi.com.cn"))

	rt := &scriptedRT{}
	rt.programOk()
	c.WithHttpTransport(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success on .com: %v", err)
	}
	if rt.hostsSeen[0] != "cvm.tencentcloudapi.com" {
		t.Fatalf("host = %s, want .com", rt.hostsSeen[0])
	}
}

// All TLDs fail across successive requests.
func Test_failover_allTldsFailAcrossRequests(t *testing.T) {
	rt := &scriptedRT{}
	rt.programErr(dnsErr{"com fail"})
	rt.programErr(dnsErr{"cn fail"})
	rt.programErr(dnsErr{"comcn fail"})
	c := runFailoverClient(rt)

	// Three requests, each hits a different TLD (cyclic rotation).
	for i := 0; i < 3; i++ {
		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("request %d should fail", i)
		}
	}
	if len(rt.hostsSeen) != 3 {
		t.Fatalf("expected 3 attempts across 3 requests, got %d", len(rt.hostsSeen))
	}
}

// Business error AuthFailure → no failover, breaker not tripped.
func Test_failover_businessErrorPropagatesImmediately(t *testing.T) {
	rt := &scriptedRT{}
	rt.programResponse(200, `{"Response":{"Error":{"Code":"AuthFailure.SignatureExpire","Message":"expired"},"RequestId":"req-123"}}`)
	c := runFailoverClient(rt)

	err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
	if err == nil {
		t.Fatal("expected SDK error")
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("business error must not trigger failover, attempts = %d", len(rt.hostsSeen))
	}
}

// Non-200 → breaker failure. 5 failures open breaker → next picks .cn.
func Test_failover_non200RecordsFailureNextRequestUsesCn(t *testing.T) {
	rt := &scriptedRT{}
	c := runFailoverClient(rt)

	for i := 0; i < 5; i++ {
		rt.programResponse(503, `{"Response":{"Error":{}}}`)
		_ = c.Send(newCvmRequest(), tchttp.NewCommonResponse())
	}

	rt.programOk()
	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("6th request should succeed on .cn: %v", err)
	}
	if rt.hostsSeen[5] != "cvm.tencentcloudapi.cn" {
		t.Fatalf("6th host = %s, want .cn", rt.hostsSeen[5])
	}
}

// Invalid JSON 200: breaker sees 200 → healthy → no failover.
func Test_failover_invalidJsonDoesNotTriggerFailover(t *testing.T) {
	rt := &scriptedRT{}
	rt.programResponse(200, `<html>blocked</html>`)
	c := runFailoverClient(rt)

	_ = c.Send(newCvmRequest(), tchttp.NewCommonResponse())
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("200 = healthy, expected 1 attempt, got %d", len(rt.hostsSeen))
	}
}

// Breaker open on .com → next request picks .cn.
func Test_failover_breakerOpenOnComShortCircuits(t *testing.T) {
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	rt := &scriptedRT{}
	c.WithHttpTransport(rt)

	tripBreakerForTest(DefaultEndpointFailover.breakerFor("cvm.tencentcloudapi.com"))

	rt.programOk()
	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success on .cn: %v", err)
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("expected 1 attempt, got %d", len(rt.hostsSeen))
	}
	if rt.hostsSeen[0] != "cvm.tencentcloudapi.cn" {
		t.Fatalf("host = %s, want .cn", rt.hostsSeen[0])
	}
}

// Happy path: single request, single attempt.
func Test_failover_happyPathOneAttempt(t *testing.T) {
	rt := &scriptedRT{}
	rt.programOk()
	c := runFailoverClient(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success: %v", err)
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("happy path made %d attempts, want 1", len(rt.hostsSeen))
	}
}

// Failover disabled → pass through.
func Test_failover_disabled_goesThrough(t *testing.T) {
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = true
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	rt := &scriptedRT{}
	rt.programErr(dnsErr{"dns miss"})
	c.WithHttpTransport(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err == nil {
		t.Fatal("expected error")
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("expected 1 attempt, got %d", len(rt.hostsSeen))
	}
}

// Sustained failure opens breaker, next request picks .cn.
func Test_failover_breakerOpensAfterSustainedFailure(t *testing.T) {
	rt := &scriptedRT{}
	c := runFailoverClient(rt)

	// 5 failures on .com open its breaker.
	for i := 0; i < 5; i++ {
		rt.programErr(dnsErr{"dns miss"})
		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatal("expected error")
		}
	}

	// 6th request: .com open → .cn.
	rt.programOk()
	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("6th request should succeed on .cn: %v", err)
	}
	if len(rt.hostsSeen) != 6 {
		t.Fatalf("expected 6 total attempts, got %d", len(rt.hostsSeen))
	}
}

// Unknown host → pass through.
func Test_failover_passThroughForUnknownHost(t *testing.T) {
	rt := &scriptedRT{}
	rt.programOk()
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	cpf.HttpProfile.Endpoint = "example.com"
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	c.WithHttpTransport(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("expected 1 attempt, got %d", len(rt.hostsSeen))
	}
}

// Non-Tencent host → pass through.
func Test_failover_nonTencentHostWithoutBackupDoesNotFailOver(t *testing.T) {
	rt := &scriptedRT{}
	rt.programErr(dnsErr{"dns miss"})
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	cpf.HttpProfile.Endpoint = "proxy.example.com"
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	c.WithHttpTransport(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err == nil {
		t.Fatal("expected error")
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("expected 1 attempt, got %d", len(rt.hostsSeen))
	}
}

// .com / .cn / .com.cn origins: each DNS miss → 1 attempt.
func Test_failover_allEndpointsDnsMissOneAttempt(t *testing.T) {
	for _, endpoint := range []string{
		"cvm.tencentcloudapi.com",
		"cvm.tencentcloudapi.cn",
		"cvm.tencentcloudapi.com.cn",
	} {
		rt := &scriptedRT{}
		rt.programErr(dnsErr{"dns miss"})
		cpf := profile.NewClientProfile()
		cpf.DisableRegionBreaker = false
		cpf.HttpProfile.Endpoint = endpoint
		c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
		c.WithHttpTransport(rt)

		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("%s: expected error", endpoint)
		}
		if len(rt.hostsSeen) != 1 {
			t.Fatalf("%s: expected 1 attempt, got %d", endpoint, len(rt.hostsSeen))
		}
	}
}

// Region-pinned host DNS miss → 1 attempt, breaker records failure.
func Test_failover_regionPinnedHostDnsMissOneAttempt(t *testing.T) {
	rt := &scriptedRT{}
	rt.programErr(dnsErr{"dns miss"})
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	cpf.HttpProfile.Endpoint = "cvm.ap-guangzhou.tencentcloudapi.com"
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	c.WithHttpTransport(rt)

	err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
	if err == nil {
		t.Fatal("expected error")
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("expected 1 attempt, got %d", len(rt.hostsSeen))
	}
}

// AI family: breaker open → next request stays in ai. family.
func Test_failover_openBreakerUsesNextAiFamilyHost(t *testing.T) {
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	cpf.HttpProfile.Endpoint = "hunyuan.ai.tencentcloudapi.com"
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)

	tripBreakerForTest(DefaultEndpointFailover.breakerFor("hunyuan.ai.tencentcloudapi.com"))

	rt := &scriptedRT{}
	rt.programOk()
	c.WithHttpTransport(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success: %v", err)
	}
	if rt.hostsSeen[0] != "hunyuan.ai.tencentcloudapi.cn" {
		t.Fatalf("host = %s, want ai. family .cn", rt.hostsSeen[0])
	}
}

// Internal family: breaker open → stays in internal. family.
func Test_failover_openBreakerUsesNextInternalFamilyHost(t *testing.T) {
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	cpf.HttpProfile.Endpoint = "cvm.internal.tencentcloudapi.com"
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)

	tripBreakerForTest(DefaultEndpointFailover.breakerFor("cvm.internal.tencentcloudapi.com"))

	rt := &scriptedRT{}
	rt.programOk()
	c.WithHttpTransport(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success: %v", err)
	}
	if rt.hostsSeen[0] != "cvm.internal.tencentcloudapi.cn" {
		t.Fatalf("host = %s, want internal. family .cn", rt.hostsSeen[0])
	}
}

// Transport errors → breaker failure, 1 attempt.
func Test_failover_transportErrorsOneAttempt(t *testing.T) {
	errors := []error{
		tlsErr{"tls handshake failed"},
		connectErr{"connection refused"},
		timeoutErr{"read timed out"},
	}
	for _, e := range errors {
		rt := &scriptedRT{}
		rt.programErr(e)
		c := runFailoverClient(rt)

		err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
		if err == nil {
			t.Fatalf("expected error for %v", e)
		}
		if len(rt.hostsSeen) != 1 {
			t.Fatalf("expected 1 attempt for %v, got %d", e, len(rt.hostsSeen))
		}
	}
}

// Response delivered after breaker selects alternate host.
func Test_failover_apiResponseDeliveredAfterSelectingAlternateHost(t *testing.T) {
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	tripBreakerForTest(DefaultEndpointFailover.breakerFor("cvm.tencentcloudapi.com"))

	rt := &scriptedRT{}
	rt.programResponse(200, `{"Response":{"RequestId":"req-xyz"}}`)
	c.WithHttpTransport(rt)

	resp := tchttp.NewCommonResponse()
	if err := c.Send(newCvmRequest(), resp); err != nil {
		t.Fatalf("expected success: %v", err)
	}
	if rt.hostsSeen[0] != "cvm.tencentcloudapi.cn" {
		t.Fatalf("expected .cn host, got %s", rt.hostsSeen[0])
	}
}

// Non-200 → 1 attempt, breaker records failure.
func Test_failover_non200OneAttempt(t *testing.T) {
	rt := &scriptedRT{}
	rt.programResponse(503, `{"Response":{"Error":{}}}`)
	c := runFailoverClient(rt)

	err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
	if err == nil {
		t.Fatal("expected error")
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("expected 1 attempt, got %d", len(rt.hostsSeen))
	}
}

// Non-200 error message contains status.
func Test_failover_non200ResponsePropagatesOriginalError(t *testing.T) {
	rt := &scriptedRT{}
	rt.programResponse(502, `{}`)
	c := runFailoverClient(rt)

	err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "502") && !strings.Contains(err.Error(), "Bad Gateway") {
		t.Fatalf("error should mention 502 or Bad Gateway, got: %v", err)
	}
}

// SSE 200 → no failover.
func Test_failover_sseStreamResponseIsNotFailoverTriggered(t *testing.T) {
	rt := &scriptedRT{}
	rt.programResponseWithCt(200, "data: hello\n\n", "text/event-stream")
	c := runFailoverClient(rt)

	_ = c.Send(newCvmRequest(), tchttp.NewCommonResponse())
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("SSE 200 should not trigger failover, got %d attempts", len(rt.hostsSeen))
	}
}

// Business SDK error → no failover.
func Test_failover_businessSdkErrorDoesNotTriggerFailover(t *testing.T) {
	rt := &scriptedRT{}
	rt.programResponse(200, `{"Response":{"RequestId":"req-bad","Error":{"Code":"AuthFailure.SignatureFailure","Message":"signature wrong"}}}`)
	c := runFailoverClient(rt)

	err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
	if err == nil {
		t.Fatal("expected business SDK error")
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("business error must not trigger failover, got %d attempts", len(rt.hostsSeen))
	}
}

// Error message contains failure details.
func Test_failover_endpointFailureSurfacesAttemptFailure(t *testing.T) {
	rt := &scriptedRT{}
	rt.programErr(dnsErr{"first dns miss"})
	c := runFailoverClient(rt)

	err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "first dns miss") {
		t.Fatalf("error should contain original message, got: %v", err)
	}
}

// .com open + .cn fails → only .cn attempted (single candidate per request).
func Test_failover_breakerSkipPlusCandidateFailure(t *testing.T) {
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	tripBreakerForTest(DefaultEndpointFailover.breakerFor("cvm.tencentcloudapi.com"))

	rt := &scriptedRT{}
	rt.programErr(tlsErr{"cn tls fail"})
	c.WithHttpTransport(rt)

	err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
	if err == nil {
		t.Fatal("expected error")
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("expected 1 attempt, got %d", len(rt.hostsSeen))
	}
	if rt.hostsSeen[0] != "cvm.tencentcloudapi.cn" {
		t.Fatalf("host = %s, want .cn", rt.hostsSeen[0])
	}
}

// All breakers open → pass through to origin (candidateFor returns nil).
func Test_failover_allBreakersOpenFallsThroughToOrigin(t *testing.T) {
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)

	for _, h := range []string{"cvm.tencentcloudapi.com", "cvm.tencentcloudapi.cn", "cvm.tencentcloudapi.com.cn"} {
		tripBreakerForTest(DefaultEndpointFailover.breakerFor(h))
	}

	rt := &scriptedRT{}
	rt.programOk()
	c.WithHttpTransport(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success on origin pass-through: %v", err)
	}
	if rt.hostsSeen[0] != "cvm.tencentcloudapi.com" {
		t.Fatalf("host = %s, want origin", rt.hostsSeen[0])
	}
}

// Error from run2 does not contain run1 messages.
func Test_failover_failureDoesNotPolluteNextRequestErrors(t *testing.T) {
	rt := &scriptedRT{}
	c := runFailoverClient(rt)

	rt.programErr(dnsErr{"run1 fail"})
	_ = c.Send(newCvmRequest(), tchttp.NewCommonResponse())

	rt.programErr(dnsErr{"run2 fail"})
	err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
	if err == nil {
		t.Fatal("expected error")
	}
	msg := err.Error()
	if strings.Contains(msg, "run1") {
		t.Fatalf("error from run2 must not contain run1 messages, got: %v", msg)
	}
	if !strings.Contains(msg, "run2") {
		t.Fatalf("error from run2 must mention run2, got: %v", msg)
	}
}

// Breaker Open → HalfOpen after cooldown.
func Test_failover_breakerTransitionsOpenToHalfOpenAfterCooldown(t *testing.T) {
	b := DefaultEndpointFailover.breakerFor("test-cooldown.tencentcloudapi.com")
	b.mu.Lock()
	b.state = StateOpen
	b.expiry = time.Now().Add(50 * time.Millisecond)
	b.mu.Unlock()

	_, err := b.beforeRequest()
	if err == nil {
		t.Fatal("breaker should be Open")
	}

	time.Sleep(100 * time.Millisecond)
	_, err = b.beforeRequest()
	if err != nil {
		t.Fatalf("breaker should permit probe after cooldown: %v", err)
	}
}

// HalfOpen success → Closed.
func Test_failover_breakerReClosesAfterHalfOpenSuccess(t *testing.T) {
	b := DefaultEndpointFailover.breakerFor("test-close.tencentcloudapi.com")
	b.mu.Lock()
	b.state = StateOpen
	b.expiry = time.Now().Add(50 * time.Millisecond)
	b.mu.Unlock()

	time.Sleep(100 * time.Millisecond)

	gen, err := b.beforeRequest()
	if err != nil {
		t.Fatalf("probe should be allowed: %v", err)
	}
	b.afterRequest(gen, true)

	if _, err := b.beforeRequest(); err != nil {
		t.Fatalf("breaker should be Closed after HalfOpen success: %v", err)
	}
}

// HalfOpen failure → Open.
func Test_failover_breakerReOpensWhenHalfOpenProbeFails(t *testing.T) {
	b := DefaultEndpointFailover.breakerFor("test-reopen.tencentcloudapi.com")
	b.mu.Lock()
	b.state = StateOpen
	b.expiry = time.Now().Add(50 * time.Millisecond)
	b.mu.Unlock()

	time.Sleep(100 * time.Millisecond)

	gen, err := b.beforeRequest()
	if err != nil {
		t.Fatalf("probe should be allowed: %v", err)
	}
	b.afterRequest(gen, false)

	if _, err := b.beforeRequest(); err == nil {
		t.Fatal("HalfOpen failure must re-Open the breaker")
	}
}

// Breaker state isolated across hosts.
func Test_failover_breakerStateIsolatedAcrossOriginHosts(t *testing.T) {
	tripBreakerForTest(DefaultEndpointFailover.breakerFor("cvm.tencentcloudapi.com"))

	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	cpf.HttpProfile.Endpoint = "cls.tencentcloudapi.com"
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	rt := &scriptedRT{}
	rt.programOk()
	c.WithHttpTransport(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("cls request should succeed: %v", err)
	}
	if rt.hostsSeen[0] != "cls.tencentcloudapi.com" {
		t.Fatalf("host = %s, want cls.tencentcloudapi.com", rt.hostsSeen[0])
	}
}

// Backup endpoint: origin succeeds → uses origin.
func Test_failover_backupEndpoint_originSucceeds_usesOrigin(t *testing.T) {
	resetBreakers()
	rt := &scriptedRT{}
	rt.programOk()
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	cpf.BackupEndpoint = "ap-guangzhou.tencentcloudapi.com"
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	c.WithHttpTransport(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success: %v", err)
	}
	if rt.hostsSeen[0] != "cvm.tencentcloudapi.com" {
		t.Fatalf("host = %s, want origin", rt.hostsSeen[0])
	}
}

// Backup endpoint: origin DNS miss → 1 attempt, breaker records failure.
func Test_failover_backupEndpoint_originDnsMissOneAttempt(t *testing.T) {
	resetBreakers()
	rt := &scriptedRT{}
	rt.programErr(dnsErr{"cvm.tencentcloudapi.com dns miss"})
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	cpf.BackupEndpoint = "ap-guangzhou.tencentcloudapi.com"
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	c.WithHttpTransport(rt)

	err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
	if err == nil {
		t.Fatal("expected error")
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("expected 1 attempt, got %d", len(rt.hostsSeen))
	}
}

// Backup endpoint: origin open → backup selected.
func Test_failover_backupEndpoint_breakerOpen_skipsOriginGoesToBackup(t *testing.T) {
	resetBreakers()
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	cpf.BackupEndpoint = "ap-guangzhou.tencentcloudapi.com"
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	tripBreakerForTest(DefaultEndpointFailover.breakerFor("cvm.tencentcloudapi.com"))

	rt := &scriptedRT{}
	rt.programOk()
	c.WithHttpTransport(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success on backup: %v", err)
	}
	if rt.hostsSeen[0] != "cvm.ap-guangzhou.tencentcloudapi.com" {
		t.Fatalf("host = %s, want backup", rt.hostsSeen[0])
	}
}

// No backup, DNS miss → 1 attempt.
func Test_failover_noBackupEndpoint_dnsMissOneAttempt(t *testing.T) {
	resetBreakers()
	rt := &scriptedRT{}
	rt.programErr(dnsErr{"cvm.tencentcloudapi.com"})
	c := runFailoverClient(rt)

	err := c.Send(newCvmRequest(), tchttp.NewCommonResponse())
	if err == nil {
		t.Fatal("expected error")
	}
	if len(rt.hostsSeen) != 1 {
		t.Fatalf("expected 1 attempt, got %d", len(rt.hostsSeen))
	}
}

// Region-pinned: breaker open → rotates to .cn.
func Test_failover_regionPinnedBreakerOpenRotatesToCn(t *testing.T) {
	resetBreakers()
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	cpf.HttpProfile.Endpoint = "cvm.ap-guangzhou.tencentcloudapi.com"
	c := NewCommonClient(NewCredential("AKIDTEST", "SKTEST"), regions.Guangzhou, cpf)
	tripBreakerForTest(DefaultEndpointFailover.breakerFor("cvm.ap-guangzhou.tencentcloudapi.com"))

	rt := &scriptedRT{}
	rt.programOk()
	c.WithHttpTransport(rt)

	if err := c.Send(newCvmRequest(), tchttp.NewCommonResponse()); err != nil {
		t.Fatalf("expected success on .cn: %v", err)
	}
	if rt.hostsSeen[0] != "cvm.ap-guangzhou.tencentcloudapi.cn" {
		t.Fatalf("host = %s, want .cn with prefix", rt.hostsSeen[0])
	}
}

// =====================================================================
// Helpers
// =====================================================================

// tripBreakerForTest drives the breaker's internal counters until it is Open.
// Using direct state manipulation is faster than running 6 real loops.
func tripBreakerForTest(b *circuitBreaker) {
	b.mu.Lock()
	b.state = StateOpen
	b.expiry = time.Now().Add(time.Hour)
	b.mu.Unlock()
}

func resetBreakers() {
	DefaultEndpointFailover.mu.Lock()
	DefaultEndpointFailover.breakers = map[string]*circuitBreaker{}
	DefaultEndpointFailover.mu.Unlock()
}
