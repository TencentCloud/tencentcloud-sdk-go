package profile

import (
	"math"
	"time"
)

type DurationFunc func(index int) time.Duration

func ConstantDurationFunc(duration time.Duration) DurationFunc {
	return func(int) time.Duration {
		return duration
	}
}

func ExponentialBackoff(index int) time.Duration {
	seconds := math.Pow(2, float64(index))
	return time.Duration(seconds) * time.Second
}

type ClientProfile struct {
	HttpProfile *HttpProfile
	// SignMethod specifies the signing algorithm used for request signatures.
	// Valid options: "HmacSHA1", "HmacSHA256", "TC3-HMAC-SHA256".
	// Default value is "TC3-HMAC-SHA256".
	SignMethod string

	// UnsignedPayload indicates whether the payload is unsigned.
	UnsignedPayload bool

	// Language specifies the language of the client profile.
	// Valid options: "zh-CN" (Chinese), "en-US" (English).
	// Default value is "zh-CN".
	Language string

	// Debug enables or disables debug output for client operations.
	Debug bool

	// DisableRegionBreaker determines whether to enable the Regional auto switch.
	// The SDK uses ap-guangzhou.tencentcloudapi.com as the default backup endpoint.
	// You can override this by specifying the BackupEndpoint.
	DisableRegionBreaker bool

	// BackupEndPoint specifies an alternative endpoint to use by region breaker.
	// Deprecated: Use BackupEndpoint instead.
	BackupEndPoint string
	BackupEndpoint string

	// NetworkFailureMaxRetries defines the maximum number of retries when a network failure occurs.
	NetworkFailureMaxRetries int
	// NetworkFailureRetryDuration is a function that determines the retry duration between network failure retries.
	NetworkFailureRetryDuration DurationFunc

	// RateLimitExceededMaxRetries defines the maximum number of retries when rate limit is exceeded.
	RateLimitExceededMaxRetries int
	// RateLimitExceededRetryDuration is a function that determines the retry duration between rate limit exceeded retries.
	RateLimitExceededRetryDuration DurationFunc

	// Configure this client to retry or not when a connectivity problem is encountered.
	// This differs from NetworkFailureMaxRetries that it will retry regardless of any request.
	// NetworkFailureMaxRetries only retry request with field "ClientToken".
	UnsafeRetryOnConnectionFailure bool
}

func NewClientProfile() *ClientProfile {
	return &ClientProfile{
		HttpProfile:     NewHttpProfile(),
		SignMethod:      "TC3-HMAC-SHA256",
		UnsignedPayload: false,
		Language:        "zh-CN",
		Debug:           false,
		// now is true, will become to false in future
		DisableRegionBreaker: true,
		BackupEndPoint:       "",
		BackupEndpoint:       "",
	}
}
