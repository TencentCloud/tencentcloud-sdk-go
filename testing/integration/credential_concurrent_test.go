package integration

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

var (
	lastTimestamp     int64
	lastTimestampTime time.Time
	timestampMutex    sync.Mutex
)

func getTimestamp() int64 {
	timestampMutex.Lock()
	defer timestampMutex.Unlock()

	now := time.Now()
	if now.Sub(lastTimestampTime) < time.Second {
		return lastTimestamp
	}

	lastTimestamp = now.Unix()
	lastTimestampTime = now
	return lastTimestamp
}

type mockMetadataTransport struct{}

func (m *mockMetadataTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	url := req.URL.String()

	if strings.HasSuffix(url, "/cam/security-credentials/") {
		body := io.NopCloser(strings.NewReader("mock-role"))
		return &http.Response{
			StatusCode: 200,
			Body:       body,
			Header:     make(http.Header),
		}, nil
	}

	if strings.HasSuffix(url, "/cam/security-credentials/mock-role") {
		n := getTimestamp()
		expireTime := n + 1
		resp := fmt.Sprintf(`{
			"TmpSecretId": "%d",
			"TmpSecretKey": "%d",
			"Token": "%d",
			"ExpiredTime": %d,
			"Code": "Success"
		}`, n, n, n, expireTime)
		body := io.NopCloser(strings.NewReader(resp))
		return &http.Response{StatusCode: 200, Body: body, Header: make(http.Header)}, nil
	}

	if strings.Contains(url, "sts.tencentcloudapi.com") {
		n := getTimestamp()
		expire := time.Now().Add(time.Second)
		resp := fmt.Sprintf(`{
			"Response": {
				"ExpiredTime": %d,
				"Expiration": "%s",
				"Credentials": {
					"Token": "%d",
					"TmpSecretId": "%d",
					"TmpSecretKey": "%d"
				},
				"RequestId": "%d"
			}
		}`, expire.Unix(), expire.Format(time.RFC3339), n, n, n, n)
		body := io.NopCloser(strings.NewReader(resp))
		return &http.Response{StatusCode: 200, Body: body, Header: make(http.Header)}, nil
	}

	return nil, fmt.Errorf("unexpected request: %s", url)
}

func runConcurrentTest(t *testing.T, getCred func() (string, string, string), duration time.Duration, threads int) int64 {
	var inconsistencies int64
	var wg sync.WaitGroup
	stop := make(chan struct{})

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				default:
					sid, skey, token := getCred()
					if sid != skey || skey != token {
						atomic.AddInt64(&inconsistencies, 1)
						t.Logf("[Goroutine %d] Inconsistent: %s / %s / %s", id, sid, skey, token)
					}
					time.Sleep(1 * time.Millisecond)
				}
			}
		}(i)
	}

	time.Sleep(duration)
	close(stop)
	wg.Wait()
	return atomic.LoadInt64(&inconsistencies)
}

func TestCvmRoleCredentialConcurrentSafeRefresh(t *testing.T) {
	http.DefaultTransport = &mockMetadataTransport{}

	provider := common.DefaultCvmRoleProvider()
	credIface, err := provider.GetCredential()
	if err != nil {
		t.Fatalf("failed to init credential: %v", err)
	}
	cred := credIface.(*common.CvmRoleCredential)

	incons := runConcurrentTest(t, cred.GetCredential, 10*time.Second, 1000)

	if incons > 0 {
		t.Errorf("TestCvmRoleCredentialConcurrentSafeRefresh failed: %d inconsistencies", incons)
	} else {
		t.Logf("TestCvmRoleCredentialConcurrentSafeRefresh passed: no inconsistencies")
	}
}

func TestRoleCredentialConcurrentSafeRefresh(t *testing.T) {
	http.DefaultTransport = &mockMetadataTransport{}

	provider := common.DefaultRoleArnProvider("mock-secret-id", "mock-secret-key", "mock-role-arn")
	credIface, err := provider.GetCredential()
	if err != nil {
		t.Fatalf("failed to init credential: %v", err)
	}
	cred := credIface.(*common.RoleArnCredential)

	incons := runConcurrentTest(t, cred.GetCredential, 10*time.Second, 1000)

	if incons > 0 {
		t.Errorf("TestRoleCredentialConcurrentSafeRefresh failed: %d inconsistencies", incons)
	} else {
		t.Logf("TestRoleCredentialConcurrentSafeRefresh passed: no inconsistencies")
	}
}

func TestOIDCRoleArnProviderConcurrentSafeRefresh(t *testing.T) {
	http.DefaultTransport = &mockMetadataTransport{}

	provider := common.NewOIDCRoleArnProvider("ap-guangzhou", "mock-provider-id",
		"mock-token", "mock-role-arn", "mock-role-session-name", 7200)

	credIface, err := provider.GetCredential()
	if err != nil {
		t.Fatalf("failed to init OIDC credential: %v", err)
	}
	cred := credIface.(*common.RoleArnCredential)

	incons := runConcurrentTest(t, cred.GetCredential, 10*time.Second, 1000)

	if incons > 0 {
		t.Errorf("TestOIDCRoleArnProviderConcurrentSafeRefresh failed: %d inconsistencies", incons)
	} else {
		t.Logf("TestOIDCRoleArnProviderConcurrentSafeRefresh passed: no inconsistencies")
	}
}

func TestTkeOIDCRoleArnProviderConcurrentSafeRefresh(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "mock-token-*.txt")
	if err != nil {
		t.Fatalf("failed to create temp token file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.WriteString("mock-token")
	tmpFile.Close()

	os.Setenv("TKE_REGION", "ap-guangzhou")
	os.Setenv("TKE_PROVIDER_ID", "mock-provider-id")
	os.Setenv("TKE_WEB_IDENTITY_TOKEN_FILE", tmpFile.Name())
	os.Setenv("TKE_ROLE_ARN", "mock-role-arn")

	http.DefaultTransport = &mockMetadataTransport{}

	provider, err := common.DefaultTkeOIDCRoleArnProvider()
	if err != nil {
		t.Fatalf("failed to create tke OIDC provider: %v", err)
	}
	credIface, err := provider.GetCredential()
	if err != nil {
		t.Fatalf("failed to init tke OIDC credential: %v", err)
	}
	cred := credIface.(*common.RoleArnCredential)

	incons := runConcurrentTest(t, cred.GetCredential, 10*time.Second, 1000)

	if incons > 0 {
		t.Errorf("TestOIDCRoleArnProviderConcurrentSafeRefresh failed: %d inconsistencies", incons)
	} else {
		t.Logf("TestOIDCRoleArnProviderConcurrentSafeRefresh passed: no inconsistencies")
	}
}
