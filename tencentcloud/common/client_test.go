package common

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

type requestWithClientToken struct {
	tchttp.CommonRequest
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
}

func newTestRequest() *requestWithClientToken {
	var testRequest = &requestWithClientToken{
		CommonRequest: *tchttp.NewCommonRequest("cvm", "2017-03-12", "RunInstances"),
		ClientToken:   nil,
	}
	return testRequest
}

type retryErr struct{}

func (r retryErr) Error() string   { return "retry error" }
func (r retryErr) Timeout() bool   { return true }
func (r retryErr) Temporary() bool { return true }

var (
	successResp   = `{"Response": {"RequestId": ""}}`
	rateLimitResp = `{"Response": {"RequestId": "", "Error": {"Code": "RequestLimitExceeded"}}}`
)

type mockRT struct {
	NetworkFailures   int
	RateLimitFailures int

	NetworkTries   int
	RateLimitTries int
}

func (s *mockRT) RoundTrip(request *http.Request) (*http.Response, error) {
	if s.NetworkTries < s.NetworkFailures {
		s.NetworkTries++
		return nil, retryErr{}
	}

	if s.RateLimitTries < s.RateLimitFailures {
		s.RateLimitTries++
		return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBufferString(rateLimitResp))}, nil
	}

	return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBufferString(successResp))}, nil
}

type testCase struct {
	msg      string
	prof     *profile.ClientProfile
	request  tchttp.Request
	response tchttp.Response
	specific mockRT
	expected mockRT
	success  bool
}

func TestNormalSucceedRequest(t *testing.T) {
	test(t, testCase{
		prof:     profile.NewClientProfile(),
		request:  newTestRequest(),
		response: tchttp.NewCommonResponse(),
		specific: mockRT{},
		expected: mockRT{},
		success:  true,
	})
}

func TestNetworkFailedButSucceedAfterRetry(t *testing.T) {
	prof := profile.NewClientProfile()
	prof.NetworkFailureMaxRetries = 1
	prof.NetworkFailureRetryDuration = profile.ConstantDurationFunc(0)

	test(t, testCase{
		prof:     prof,
		request:  newTestRequest(),
		response: tchttp.NewCommonResponse(),
		specific: mockRT{NetworkFailures: 1},
		expected: mockRT{NetworkTries: 1},
		success:  true,
	})
}

func TestNetworkFailedAndShouldNotRetry(t *testing.T) {
	prof := profile.NewClientProfile()
	prof.NetworkFailureMaxRetries = 1
	prof.NetworkFailureRetryDuration = profile.ConstantDurationFunc(0)

	test(t, testCase{
		prof:     prof,
		request:  tchttp.NewCommonRequest("cvm", "2017-03-12", "DescribeInstances"),
		response: tchttp.NewCommonResponse(),
		specific: mockRT{NetworkFailures: 2},
		expected: mockRT{NetworkTries: 1},
		success:  false,
	})
}

func TestNetworkFailedAfterRetry(t *testing.T) {
	prof := profile.NewClientProfile()
	prof.NetworkFailureMaxRetries = 1
	prof.NetworkFailureRetryDuration = profile.ConstantDurationFunc(0)

	test(t, testCase{
		prof:     prof,
		request:  newTestRequest(),
		response: tchttp.NewCommonResponse(),
		specific: mockRT{NetworkFailures: 2},
		expected: mockRT{NetworkTries: 2},
		success:  false,
	})
}

func TestRateLimitButSucceedAfterRetry(t *testing.T) {
	prof := profile.NewClientProfile()
	prof.RateLimitExceededMaxRetries = 1
	prof.RateLimitExceededRetryDuration = profile.ConstantDurationFunc(0)

	test(t, testCase{
		prof:     prof,
		request:  newTestRequest(),
		response: tchttp.NewCommonResponse(),
		specific: mockRT{RateLimitFailures: 1},
		expected: mockRT{RateLimitTries: 1},
		success:  true,
	})
}

func TestRateLimitExceededAfterRetry(t *testing.T) {
	prof := profile.NewClientProfile()
	prof.RateLimitExceededMaxRetries = 1
	prof.RateLimitExceededRetryDuration = profile.ConstantDurationFunc(0)

	test(t, testCase{
		prof:     prof,
		request:  newTestRequest(),
		response: tchttp.NewCommonResponse(),
		specific: mockRT{RateLimitFailures: 3},
		expected: mockRT{RateLimitTries: 2},
		success:  false,
	})
}

func TestBothFailuresOccurredButSucceedAfterRetry(t *testing.T) {
	prof := profile.NewClientProfile()
	prof.NetworkFailureMaxRetries = 1
	prof.NetworkFailureRetryDuration = profile.ConstantDurationFunc(0)
	prof.RateLimitExceededMaxRetries = 1
	prof.RateLimitExceededRetryDuration = profile.ConstantDurationFunc(0)

	test(t, testCase{
		prof:     prof,
		request:  newTestRequest(),
		response: tchttp.NewCommonResponse(),
		specific: mockRT{RateLimitFailures: 1, NetworkFailures: 1},
		expected: mockRT{RateLimitTries: 1, NetworkTries: 1},
		success:  true,
	})
}

func test(t *testing.T, tc testCase) {
	credential := NewCredential("", "")
	client := NewCommonClient(credential, regions.Guangzhou, tc.prof)

	client.WithHttpTransport(&tc.specific)

	err := client.Send(tc.request, tc.response)
	if tc.success != (err == nil) {
		t.Fatalf("unexpected failed on request: %+v", err)
	}

	if tc.expected.RateLimitTries != tc.specific.RateLimitTries {
		t.Fatalf("unexpected rate limit retry, expected %d, got %d",
			tc.expected.RateLimitTries, tc.specific.RateLimitTries)
	}

	if tc.expected.NetworkTries != tc.specific.NetworkTries {
		t.Fatalf("unexpected network failure retry, expected %d, got %d",
			tc.expected.NetworkTries, tc.specific.NetworkTries)
	}
}

func TestClient_withRegionBreaker(t *testing.T) {
	cpf := profile.NewClientProfile()
	//cpf.Debug =true
	cpf.DisableRegionBreaker = false
	cpf.BackupEndpoint = ""
	c := (&Client{}).Init("")
	c.WithProfile(cpf)
	if c.rb.backupEndpoint != "ap-guangzhou.tencentcloudapi.com" {
		t.Errorf("want %s ,got %s", "ap-beijing", c.rb.backupEndpoint)
	}
	if c.rb.maxFailNum != defaultMaxFailNum {
		t.Errorf("want %d ,got %d", defaultMaxFailNum, c.rb.maxFailNum)
	}
}

// breakerRT is a RoundTripper that always returns a successful empty
// CVM-style response. It is the simplest fixture that exercises the full
// Send → sendWithRegionBreaker → sendWithSignature → response-parse chain
// while letting Send return nil err.
type breakerRT struct {
	calls int
}

func (s *breakerRT) RoundTrip(*http.Request) (*http.Response, error) {
	s.calls++
	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(successResp)),
	}, nil
}

// newBreakerEnabledClient builds a Client wired through the same paths Send()
// uses in production, with the regional breaker turned on. The returned Client
// uses rt as its HTTP transport so test code can drive Send() without network.
func newBreakerEnabledClient(rt http.RoundTripper) *Client {
	cpf := profile.NewClientProfile()
	cpf.DisableRegionBreaker = false
	c := NewCommonClient(NewCredential("id", "key"), regions.Guangzhou, cpf)
	c.WithHttpTransport(rt)
	return c
}

// TestClient_breakerStaysClosedOnSuccesses is the end-to-end regression test
// for report/feature_4_bug.md Bug 1 ("circuit breaker counts every successful
// call as a failure"). It drives the full Client.Send path — including
// sendWithRegionBreaker — against a mock transport that always returns 200,
// and asserts the breaker remains closed after far more than maxFailNum
// successes. With the bug present, the breaker would open after 5 successful
// calls and the 6th call would be re-routed to ap-guangzhou.tencentcloudapi.com.
func TestClient_breakerStaysClosedOnSuccesses(t *testing.T) {
	rt := &breakerRT{}
	c := newBreakerEnabledClient(rt)

	const calls = 12 // > defaultMaxFailNum (5) and > the consecutive-failures threshold (6)
	for i := 0; i < calls; i++ {
		req := newTestRequest()
		resp := tchttp.NewCommonResponse()
		if err := c.Send(req, resp); err != nil {
			t.Fatalf("call %d: unexpected error: %v", i, err)
		}
		// The breaker rewrites request.SetDomain only when the breaker is open.
		// If the bug were present, calls 6..N would target the backup endpoint.
		if got := req.GetDomain(); got != "" && got != "cvm.tencentcloudapi.com" {
			// Allow the empty-domain default; reject anything with the backup host.
			if got == "cvm.ap-guangzhou.tencentcloudapi.com" && c.GetRegion() != "ap-guangzhou" {
				t.Fatalf("call %d: request was re-routed to backup endpoint %q (breaker opened spuriously)", i, got)
			}
		}
	}

	if rt.calls != calls {
		t.Fatalf("transport saw %d calls, want %d", rt.calls, calls)
	}

	c.rb.mu.Lock()
	st := c.rb.state
	c.rb.mu.Unlock()
	if st != StateClosed {
		t.Fatalf("after %d successful calls, breaker state = %v, want StateClosed (Bug 1 regression)", calls, st)
	}
}

// failingRT returns a transport-level error every call, simulating a region
// outage. We use it to verify the breaker DOES open when calls genuinely fail
// — i.e. that the fix didn't make the breaker permanently inert.
type failingRT struct{ calls int }

func (s *failingRT) RoundTrip(*http.Request) (*http.Response, error) {
	s.calls++
	return nil, retryErr{}
}

// TestClient_breakerOpensOnFailures is the dual of the previous test: with
// the fixed classifier, transport-level failures must still trip the breaker
// and re-route subsequent traffic to the backup endpoint. This guards against
// an over-correction that would make the breaker never trip.
func TestClient_breakerOpensOnFailures(t *testing.T) {
	rt := &failingRT{}
	c := newBreakerEnabledClient(rt)

	// Drive enough failures to trip both the rate-based branch (5 failures, 100% rate)
	// and the consecutive-failures branch (>5).
	for i := 0; i < 7; i++ {
		req := newTestRequest()
		resp := tchttp.NewCommonResponse()
		_ = c.Send(req, resp) // failures expected; we only care about breaker state
	}

	c.rb.mu.Lock()
	st := c.rb.state
	c.rb.mu.Unlock()
	if st != StateOpen {
		t.Fatalf("after 7 transport failures, breaker state = %v, want StateOpen", st)
	}

	// One more call should be re-routed to the backup endpoint.
	req := newTestRequest()
	resp := tchttp.NewCommonResponse()
	_ = c.Send(req, resp)
	if got := req.GetDomain(); got != "cvm."+c.rb.backupEndpoint {
		t.Fatalf("after breaker tripped, request domain = %q, want %q (re-routed to backup)",
			got, "cvm."+c.rb.backupEndpoint)
	}
}
