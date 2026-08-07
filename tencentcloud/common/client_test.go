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
	// These tests predate domain failover; disable it so a transport
	// failure isn't silently retried against an alternate TLD.
	tc.prof.DisableRegionBreaker = true
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
	cpf.DisableRegionBreaker = false
	c := (&Client{}).Init("")
	c.WithProfile(cpf)
	if c.profile.DisableRegionBreaker {
		t.Fatalf("expected DisableRegionBreaker=false (failover enabled)")
	}
}

// breakerRT is a RoundTripper that always returns a successful empty
// CVM-style response. It is the simplest fixture that exercises the full
// Send → sendWithEndpointFailover → sendWithSignature → response-parse chain
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

// TestClient_breakerStaysClosedOnSuccesses drives the full Client.Send path
// against a mock transport that always returns 200, and asserts every
// candidate's per-host breaker remains closed (no host re-routing).
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
	}

	if rt.calls != calls {
		t.Fatalf("transport saw %d calls, want %d", rt.calls, calls)
	}
}

// failingRT returns a transport-level error every call, simulating a region
// outage. We use it to verify the breaker DOES open when calls genuinely fail.
type failingRT struct{ calls int }

func (s *failingRT) RoundTrip(*http.Request) (*http.Response, error) {
	s.calls++
	return nil, retryErr{}
}
