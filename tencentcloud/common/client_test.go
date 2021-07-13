package common_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

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
		request:  cvm.NewRunInstancesRequest(),
		response: cvm.NewRunInstancesResponse(),
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
		request:  cvm.NewRunInstancesRequest(),
		response: cvm.NewRunInstancesResponse(),
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
		request:  cvm.NewDescribeInstancesRequest(),
		response: cvm.NewDescribeInstancesResponse(),
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
		request:  cvm.NewRunInstancesRequest(),
		response: cvm.NewRunInstancesResponse(),
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
		request:  cvm.NewRunInstancesRequest(),
		response: cvm.NewRunInstancesResponse(),
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
		request:  cvm.NewRunInstancesRequest(),
		response: cvm.NewRunInstancesResponse(),
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
		request:  cvm.NewRunInstancesRequest(),
		response: cvm.NewRunInstancesResponse(),
		specific: mockRT{RateLimitFailures: 1, NetworkFailures: 1},
		expected: mockRT{RateLimitTries: 1, NetworkTries: 1},
		success:  true,
	})
}

func test(t *testing.T, tc testCase) {
	credential := common.NewCredential("", "")
	client, err := cvm.NewClient(credential, regions.Guangzhou, tc.prof)
	if err != nil {
		t.Fatalf("failed on createing client: %+v", err)
	}

	client.WithHttpTransport(&tc.specific)

	err = client.Send(tc.request, tc.response)
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
