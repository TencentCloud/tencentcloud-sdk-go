package common

import (
	"errors"
	"testing"
	"time"

	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
)

func Test_checkDomain(t *testing.T) {
	type args struct {
		endpoint string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid endpoint", args{endpoint: "cvm.tencentcloudapi.com"}, true},
		{"valid endpoint", args{endpoint: "cvm.ap-beijing.tencentcloudapi.com"}, true},
		{"invalid endpoint", args{endpoint: "cvm.tencentcloud.com"}, false},
		{"invalid endpoint", args{endpoint: "cvm.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkEndpoint(tt.args.endpoint); got != tt.want {
				t.Errorf("checkEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_renewUrl(t *testing.T) {
	type args struct {
		oldDomain string
		region    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"success3", args{
			oldDomain: "cvm.tencentcloudapi.com",
			region:    "ap-beijing",
		}, "cvm.ap-beijing.tencentcloudapi.com"},
		{"success4", args{
			oldDomain: "cvm.ap-beijing.tencentcloudapi.com",
			region:    "ap-shanghai",
		}, "cvm.ap-shanghai.tencentcloudapi.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renewUrl(tt.args.oldDomain, tt.args.region); got != tt.want {
				t.Errorf("renewUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test_counter_onFailure_incrementsConsecutiveFailures guards against the typo in
// counter.onFailure where consecutiveSuccesses was being zeroed twice instead of
// incrementing consecutiveFailures. Symptom: the "consecutive failures" trip
// branch in readyToOpen was unreachable dead code.
func Test_counter_onFailure_incrementsConsecutiveFailures(t *testing.T) {
	c := newRegionCounter()
	for i := 0; i < 6; i++ {
		c.onFailure()
	}
	if c.consecutiveFailures != 6 {
		t.Fatalf("after 6 onFailure calls, consecutiveFailures = %d, want 6", c.consecutiveFailures)
	}
	if c.failures != 6 || c.all != 6 {
		t.Fatalf("after 6 onFailure calls, failures=%d all=%d, want 6/6", c.failures, c.all)
	}
}

// Test_counter_onSuccess_resetsConsecutiveFailures verifies that a single success
// resets consecutiveFailures back to zero (the well-known burst recovery contract).
func Test_counter_onSuccess_resetsConsecutiveFailures(t *testing.T) {
	c := newRegionCounter()
	c.onFailure()
	c.onFailure()
	c.onSuccess()
	if c.consecutiveFailures != 0 {
		t.Fatalf("after onSuccess, consecutiveFailures = %d, want 0", c.consecutiveFailures)
	}
}

// Test_counter_clear_resetsConsecutiveFailures guards against counter.clear() forgetting
// to reset consecutiveFailures. Without this, generation rollovers carry over a stale
// failure count.
func Test_counter_clear_resetsConsecutiveFailures(t *testing.T) {
	c := newRegionCounter()
	c.onFailure()
	c.onFailure()
	c.clear()
	if c.consecutiveFailures != 0 {
		t.Fatalf("after clear, consecutiveFailures = %d, want 0", c.consecutiveFailures)
	}
	if c.failures != 0 || c.all != 0 || c.consecutiveSuccesses != 0 {
		t.Fatalf("after clear, failures=%d all=%d consecutiveSuccesses=%d, want all 0",
			c.failures, c.all, c.consecutiveSuccesses)
	}
}

// Test_circuitBreaker_tripsOnConsecutiveFailures verifies the consecutive-failures
// trip branch fires once consecutiveFailures reaches maxFailNum.
func Test_circuitBreaker_tripsOnConsecutiveFailures(t *testing.T) {
	cb := newRegionBreaker(breakerSetting{
		backupEndpoint:    defaultBackupEndpoint,
		maxFailNum:        5,
		maxFailPercentage: 75,
		windowInterval:    1 * time.Hour,
		timeout:           60 * time.Second,
		maxRequests:       defaultMaxRequests,
	})
	for i := 0; i < 5; i++ {
		gen, err := cb.beforeRequest()
		if err != nil {
			t.Fatalf("iteration %d: beforeRequest returned err=%v before breaker should open", i, err)
		}
		cb.afterRequest(gen, false)
	}
	// After 5 consecutive failures the breaker must be open.
	_, err := cb.beforeRequest()
	if err != errOpenState {
		t.Fatalf("after 5 consecutive failures, breaker should be open, got err=%v", err)
	}
}

// Test_isBreakerSuccess_nilErrIsSuccess is the regression test for the bug where
// sendWithRegionBreaker treated nil err as failure. nil err must be classified as
// success, otherwise every successful API call counts as a failure to the breaker.
func Test_isBreakerSuccess_nilErrIsSuccess(t *testing.T) {
	if !isBreakerSuccess(nil) {
		t.Fatal("isBreakerSuccess(nil) = false, want true (nil error is a successful call)")
	}
}

// Test_isBreakerSuccess_sdkErrorWithRequestIdIsSuccess: a structured server error
// with a RequestId still indicates the region answered, so the breaker should treat it
// as success unless it is an InternalError.
func Test_isBreakerSuccess_sdkErrorWithRequestIdIsSuccess(t *testing.T) {
	e := tcerr.NewTencentCloudSDKError("AuthFailure.SignatureExpire", "expired", "req-123")
	if !isBreakerSuccess(e) {
		t.Fatal("isBreakerSuccess(AuthFailure with RequestId) = false, want true")
	}
}

// Test_isBreakerSuccess_internalErrorIsFailure: InternalError indicates region-side
// fault; treat as a failure for breaker accounting.
func Test_isBreakerSuccess_internalErrorIsFailure(t *testing.T) {
	e := tcerr.NewTencentCloudSDKError("InternalError", "boom", "req-123")
	if isBreakerSuccess(e) {
		t.Fatal("isBreakerSuccess(InternalError) = true, want false")
	}
}

// Test_isBreakerSuccess_sdkErrorWithoutRequestIdIsFailure: missing RequestId means
// the response never reached the user from the region (locally-fabricated error,
// network failure mid-flight, etc.), so the breaker should consider it a failure.
func Test_isBreakerSuccess_sdkErrorWithoutRequestIdIsFailure(t *testing.T) {
	e := tcerr.NewTencentCloudSDKError("ClientError.NetworkError", "boom", "")
	if isBreakerSuccess(e) {
		t.Fatal("isBreakerSuccess(SDKError without RequestId) = true, want false")
	}
}

// Test_isBreakerSuccess_nonSDKErrorIsFailure: arbitrary network/transport error must
// be a failure.
func Test_isBreakerSuccess_nonSDKErrorIsFailure(t *testing.T) {
	if isBreakerSuccess(errors.New("connection refused")) {
		t.Fatal("isBreakerSuccess(non-SDK error) = true, want false")
	}
}

// Test_circuitBreaker_doesNotOpenOnSuccesses is the integration-level regression test:
// 5 successful calls in a row must NOT open the breaker, otherwise users get silent
// wrong-region routing.
func Test_circuitBreaker_doesNotOpenOnSuccesses(t *testing.T) {
	cb := defaultRegionBreaker()
	for i := 0; i < 10; i++ {
		gen, err := cb.beforeRequest()
		if err != nil {
			t.Fatalf("iteration %d: breaker opened spuriously, err=%v", i, err)
		}
		// Simulate a successful call: nil error -> isBreakerSuccess returns true.
		cb.afterRequest(gen, isBreakerSuccess(nil))
	}
	// Confirm state is still closed.
	cb.mu.Lock()
	st := cb.state
	cb.mu.Unlock()
	if st != StateClosed {
		t.Fatalf("after 10 successes, breaker state = %v, want StateClosed", st)
	}
}

// Test_circuitBreaker_generationRollover_clearsConsecutiveFailures guards against
// consecutiveFailures surviving a generation rollover.
func Test_circuitBreaker_generationRollover_clearsConsecutiveFailures(t *testing.T) {
	cb := newRegionBreaker(breakerSetting{
		backupEndpoint:    defaultBackupEndpoint,
		maxFailNum:        5,
		maxFailPercentage: 75,
		windowInterval:    1 * time.Millisecond,
		timeout:           60 * time.Second,
		maxRequests:       defaultMaxRequests,
	})
	// Record a few failures (not enough to trip).
	for i := 0; i < 3; i++ {
		gen, err := cb.beforeRequest()
		if err != nil {
			t.Fatalf("iteration %d: unexpected err=%v", i, err)
		}
		cb.afterRequest(gen, false)
	}
	// Wait past windowInterval so the next beforeRequest rolls a new generation.
	time.Sleep(5 * time.Millisecond)
	_, err := cb.beforeRequest()
	if err != nil {
		t.Fatalf("after rollover, unexpected err=%v", err)
	}
	cb.mu.Lock()
	cf := cb.counter.consecutiveFailures
	all := cb.counter.all
	failures := cb.counter.failures
	cb.mu.Unlock()
	if cf != 0 || all != 0 || failures != 0 {
		t.Fatalf("after generation rollover counter not cleared: consecutiveFailures=%d failures=%d all=%d",
			cf, failures, all)
	}
}

// Test_circuitBreaker_halfOpenRequiresMaxRequests verifies that the half-open ->
// closed transition waits for maxRequests consecutive successes, rather than
// closing on the first success. Regression test for the missing maxRequests
// default in defaultRegionBreaker.
func Test_circuitBreaker_halfOpenRequiresMaxRequests(t *testing.T) {
	cb := newRegionBreaker(breakerSetting{
		backupEndpoint:    defaultBackupEndpoint,
		maxFailNum:        5,
		maxFailPercentage: 75,
		windowInterval:    1 * time.Hour,
		timeout:           1 * time.Millisecond, // trip fast for the test
		maxRequests:       3,
	})
	// Trip the breaker with 5 consecutive failures.
	for i := 0; i < 5; i++ {
		gen, _ := cb.beforeRequest()
		cb.afterRequest(gen, false)
	}
	if _, err := cb.beforeRequest(); err != errOpenState {
		t.Fatalf("expected breaker open after 5 failures, got %v", err)
	}
	// Wait past timeout so next beforeRequest moves to half-open.
	time.Sleep(5 * time.Millisecond)
	// Send 2 successes: not enough, breaker should still be half-open.
	for i := 0; i < 2; i++ {
		gen, err := cb.beforeRequest()
		if err != nil {
			t.Fatalf("half-open iteration %d: unexpected err=%v", i, err)
		}
		cb.afterRequest(gen, true)
	}
	cb.mu.Lock()
	st := cb.state
	cb.mu.Unlock()
	if st != StateHalfOpen {
		t.Fatalf("after 2 successes in half-open, state = %v, want StateHalfOpen", st)
	}
	// Third success should close.
	gen, _ := cb.beforeRequest()
	cb.afterRequest(gen, true)
	cb.mu.Lock()
	st = cb.state
	cb.mu.Unlock()
	if st != StateClosed {
		t.Fatalf("after 3 successes in half-open, state = %v, want StateClosed", st)
	}
}
