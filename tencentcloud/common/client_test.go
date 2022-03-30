package common

import (
	"bytes"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"io/ioutil"
	"net/http"
	"testing"
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

func TestClient_PacketSizeLimit(t *testing.T) {
	cli, err := NewClientWithSecretId("", "", regions.Guangzhou)
	if err != nil {
		t.Fatal(err)
	}
	cli = cli.WithProfile(profile.NewClientProfile())

	const (
		KB = 1024
		MB = KB * 1024
	)

	testCases := []struct {
		httpMethod string
		signMethod string
		size       int64
		limit      int64
		limited    bool
	}{
		//{
		//	httpMethod: http.MethodGet,
		//	signMethod: SHA1,
		//	size:       1 * KB,
		//	limit:      0,
		//	limited:    false,
		//},
		//{
		//	httpMethod: http.MethodGet,
		//	signMethod: SHA1,
		//	size:       1 * KB,
		//	limit:      5,
		//	limited:    true,
		//},
		//{
		//	httpMethod: http.MethodGet,
		//	signMethod: SHA1,
		//	size:       1 * KB,
		//	// limit should > ( size * 2 + len(http://domain.com/path) ) because url-encode double the data size
		//	limit:   4 * KB,
		//	limited: false,
		//},
		//{
		//	httpMethod: http.MethodGet,
		//	signMethod: SHA1,
		//	size:       33 * KB,
		//	limit:      0,
		//	limited:    true,
		//},
		//{
		//	httpMethod: http.MethodGet,
		//	signMethod: SHA1,
		//	size:       33 * KB,
		//	limit:      34 * KB,
		//	limited:    true,
		//},
		//
		//{
		//	httpMethod: http.MethodGet,
		//	signMethod: SHA256,
		//	size:       1 * KB,
		//	limit:      0,
		//	limited:    false,
		//},
		//{
		//	httpMethod: http.MethodGet,
		//	signMethod: SHA256,
		//	size:       1 * KB,
		//	limit:      5,
		//	limited:    true,
		//},
		//{
		//	httpMethod: http.MethodGet,
		//	signMethod: SHA256,
		//	size:       1 * KB,
		//	limit:      4 * KB,
		//	limited:    false,
		//},
		//{
		//	httpMethod: http.MethodGet,
		//	signMethod: SHA256,
		//	size:       33 * KB,
		//	limit:      0,
		//	limited:    true,
		//},
		//{
		//	httpMethod: http.MethodGet,
		//	signMethod: SHA256,
		//	size:       33 * KB,
		//	limit:      34 * KB,
		//	limited:    true,
		//},
		//
		//{
		//	httpMethod: http.MethodGet,
		//	signMethod: "TC3-HMAC-SHA256",
		//	size:       1 * KB,
		//	limit:      0,
		//	limited:    false,
		//},
		//{
		//	httpMethod: http.MethodGet,
		//	signMethod: "TC3-HMAC-SHA256",
		//	size:       1 * KB,
		//	limit:      5,
		//	limited:    true,
		//},
		//{
		//	httpMethod: http.MethodGet,
		//	signMethod: "TC3-HMAC-SHA256",
		//	size:       1 * KB,
		//	limit:      4 * KB,
		//	limited:    false,
		//},
		//{
		//	httpMethod: http.MethodGet,
		//	signMethod: "TC3-HMAC-SHA256",
		//	size:       33 * KB,
		//	limit:      0,
		//	limited:    true,
		//},
		//{
		//	httpMethod: http.MethodGet,
		//	signMethod: "TC3-HMAC-SHA256",
		//	size:       33 * KB,
		//	limit:      34 * KB,
		//	limited:    true,
		//},
		//
		//{
		//	httpMethod: http.MethodPost,
		//	signMethod: SHA1,
		//	size:       1 * KB,
		//	limit:      0,
		//	limited:    false,
		//},
		//{
		//	httpMethod: http.MethodPost,
		//	signMethod: SHA1,
		//	size:       1 * KB,
		//	limit:      5,
		//	limited:    true,
		//},
		//{
		//	httpMethod: http.MethodPost,
		//	signMethod: SHA1,
		//	size:       1 * KB,
		//	limit:      4 * KB,
		//	limited:    false,
		//},
		//{
		//	httpMethod: http.MethodPost,
		//	signMethod: SHA1,
		//	size:       2 * MB,
		//	limit:      0,
		//	limited:    true,
		//},
		//{
		//	httpMethod: http.MethodPost,
		//	signMethod: SHA1,
		//	size:       2 * MB,
		//	limit:      6 * MB,
		//	limited:    true,
		//},
		//
		//{
		//	httpMethod: http.MethodPost,
		//	signMethod: SHA256,
		//	size:       1 * KB,
		//	limit:      0,
		//	limited:    false,
		//},
		//{
		//	httpMethod: http.MethodPost,
		//	signMethod: SHA256,
		//	size:       1 * KB,
		//	limit:      5,
		//	limited:    true,
		//},
		//{
		//	httpMethod: http.MethodPost,
		//	signMethod: SHA256,
		//	size:       1 * KB,
		//	limit:      4 * KB,
		//	limited:    false,
		//},
		//{
		//	httpMethod: http.MethodPost,
		//	signMethod: SHA256,
		//	size:       2 * MB,
		//	limit:      0,
		//	limited:    true,
		//},
		//{
		//	httpMethod: http.MethodPost,
		//	signMethod: SHA256,
		//	size:       2 * MB,
		//	limit:      6 * MB,
		//	limited:    true,
		//},

		{
			httpMethod: http.MethodPost,
			signMethod: "TC3-HMAC-SHA256",
			size:       1 * KB,
			limit:      0,
			limited:    false,
		},
		{
			httpMethod: http.MethodPost,
			signMethod: "TC3-HMAC-SHA256",
			size:       1 * KB,
			limit:      5,
			limited:    true,
		},
		{
			httpMethod: http.MethodPost,
			signMethod: "TC3-HMAC-SHA256",
			size:       1 * KB,
			limit:      4 * KB,
			limited:    false,
		},
		{
			httpMethod: http.MethodPost,
			signMethod: "TC3-HMAC-SHA256",
			size:       11 * MB,
			limit:      0,
			limited:    true,
		},
		{
			httpMethod: http.MethodPost,
			signMethod: "TC3-HMAC-SHA256",
			size:       11 * MB,
			limit:      33 * MB,
			limited:    true,
		},
	}

	for _, tc := range testCases {
		cli.signMethod = tc.signMethod
		r := newTestRequest()
		r.SetHttpMethod(tc.httpMethod)

		if tc.httpMethod == http.MethodGet {
			r.GetParams()["someParams"] = longString(tc.size)
		} else {
			err = r.SetActionParameters(map[string]interface{}{"someParams": longString(tc.size)})
			if err != nil {
				t.Fatal(err)
			}
		}

		r.SetPacketSizeLimit(tc.limit)
		rsp := tchttp.NewCommonResponse()
		err = cli.Send(r, rsp)

		_, limited := err.(PacketTooLargeError)
		if limited != tc.limited {
			t.Fatalf("httpMethod:%s, signMethod:%s, size:%d, limit:%d, expected_limited: %v, got: %v",
				tc.httpMethod, tc.signMethod, tc.size, tc.limit, tc.limited, limited)
		}
	}
}

func longString(n int64) string {
	p := make([]byte, n)
	for i := int64(0); i < n; i++ {
		p[i] = 'a'
	}
	return string(p)
}
