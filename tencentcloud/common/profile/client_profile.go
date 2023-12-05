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
	// Valid choices: HmacSHA1, HmacSHA256, TC3-HMAC-SHA256.
	// Default value is TC3-HMAC-SHA256.
	SignMethod      string
	UnsignedPayload bool
	// Valid choices: zh-CN, en-US.
	// Default value is zh-CN.
	Language string
	Debug    bool
	// define Whether to enable Regional auto switch
	DisableRegionBreaker bool

	// Deprecated. Use BackupEndpoint instead.
	BackupEndPoint string
	BackupEndpoint string

	// define how to retry request
	NetworkFailureMaxRetries       int
	NetworkFailureRetryDuration    DurationFunc
	RateLimitExceededMaxRetries    int
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
