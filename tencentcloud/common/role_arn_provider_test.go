package common

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

type roleArnRoundTripper func(*http.Request) (*http.Response, error)

func (f roleArnRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request)
}

func TestRoleArnProviderWithCredential(t *testing.T) {
	const (
		sourceSecretID = "source-secret-id"
		sourceToken    = "source-token"
	)

	oldDefaultHTTPClient := DefaultHttpClient
	defer func() { DefaultHttpClient = oldDefaultHTTPClient }()

	var requestCredential, requestToken string
	DefaultHttpClient = &http.Client{Transport: roleArnRoundTripper(func(request *http.Request) (*http.Response, error) {
		requestCredential = request.Header.Get("Authorization")
		requestToken = request.Header.Get("X-TC-Token")
		body := `{"Response":{"Credentials":{"Token":"assumed-token","TmpSecretId":"assumed-secret-id","TmpSecretKey":"assumed-secret-key"},"ExpiredTime":` +
			strconv.FormatInt(time.Now().Add(time.Hour).Unix(), 10) + `,"RequestId":"request-id"}}`
		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     "200 OK",
			Header:     make(http.Header),
			Body:       ioutil.NopCloser(strings.NewReader(body)),
			Request:    request,
		}, nil
	})}

	sourceCredential := NewTokenCredential(sourceSecretID, "source-secret-key", sourceToken)
	provider := NewRoleArnProviderWithCredential(sourceCredential, "target-role-arn", "session-name", 900)
	provider.Endpoint = "sts.test"

	credential, err := provider.GetCredential()
	if err != nil {
		t.Fatalf("GetCredential() error = %v", err)
	}
	if !strings.Contains(requestCredential, "Credential="+sourceSecretID+"/") {
		t.Errorf("Authorization does not use the supplied credential: %q", requestCredential)
	}
	if requestToken != sourceToken {
		t.Errorf("X-TC-Token = %q, want %q", requestToken, sourceToken)
	}
	secretID, secretKey, token := credential.GetCredential()
	if secretID != "assumed-secret-id" || secretKey != "assumed-secret-key" || token != "assumed-token" {
		t.Errorf("GetCredential() = (%q, %q, %q), want assumed credentials", secretID, secretKey, token)
	}
}

func TestNewRoleArnProviderUsesLongTermCredential(t *testing.T) {
	provider := NewRoleArnProvider("secret-id", "secret-key", "role-arn", "session-name", 900)
	secretID, secretKey, token := provider.credential.GetCredential()
	if secretID != "secret-id" || secretKey != "secret-key" || token != "" {
		t.Errorf("credential = (%q, %q, %q), want long-term credential", secretID, secretKey, token)
	}
}

func TestDefaultRoleArnProviderWithCredential(t *testing.T) {
	credential := NewTokenCredential("secret-id", "secret-key", "token")
	provider := DefaultRoleArnProviderWithCredential(credential, "role-arn")

	if provider.credential != credential {
		t.Error("provider does not use the supplied credential")
	}
	if provider.roleArn != "role-arn" {
		t.Errorf("roleArn = %q, want %q", provider.roleArn, "role-arn")
	}
	if !strings.HasPrefix(provider.roleSessionName, defaultSessionName) {
		t.Errorf("roleSessionName = %q, want prefix %q", provider.roleSessionName, defaultSessionName)
	}
	if provider.durationSeconds != defaultDurationSeconds {
		t.Errorf("durationSeconds = %d, want %d", provider.durationSeconds, defaultDurationSeconds)
	}
}
