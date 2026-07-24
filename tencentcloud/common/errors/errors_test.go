package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestTencentCloudSDKError_Error(t *testing.T) {
	cases := []struct {
		name string
		err  *TencentCloudSDKError
		want string
	}{
		{
			name: "no request id, no cause",
			err:  &TencentCloudSDKError{Code: "ClientError.NetworkError", Message: "Fail to get response"},
			want: "[TencentCloudSDKError] Code=ClientError.NetworkError, Message=Fail to get response",
		},
		{
			name: "with request id, no cause",
			err:  &TencentCloudSDKError{Code: "RequestLimitExceeded", Message: "rate limited", RequestId: "req-1"},
			want: "[TencentCloudSDKError] Code=RequestLimitExceeded, Message=rate limited, RequestId=req-1",
		},
		{
			name: "no request id, with cause",
			err:  &TencentCloudSDKError{Code: "ClientError.NetworkError", Message: "Fail to get response", cause: fmt.Errorf("dial tcp: connection refused")},
			want: "[TencentCloudSDKError] Code=ClientError.NetworkError, Message=Fail to get response: dial tcp: connection refused",
		},
		{
			name: "with request id, with cause",
			err:  &TencentCloudSDKError{Code: "ClientError.ParseJsonError", Message: "bad json", RequestId: "req-2", cause: fmt.Errorf("unexpected end of JSON input")},
			want: "[TencentCloudSDKError] Code=ClientError.ParseJsonError, Message=bad json, RequestId=req-2: unexpected end of JSON input",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.err.Error(); got != tc.want {
				t.Fatalf("Error() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestTencentCloudSDKError_Unwrap(t *testing.T) {
	t.Run("nil cause returns nil", func(t *testing.T) {
		e := &TencentCloudSDKError{Code: "c", Message: "m"}
		if got := e.Unwrap(); got != nil {
			t.Fatalf("Unwrap() = %v, want nil", got)
		}
	})

	t.Run("cause is returned", func(t *testing.T) {
		cause := fmt.Errorf("underlying")
		e := &TencentCloudSDKError{Code: "c", Message: "m", cause: cause}
		if got := e.Unwrap(); got != cause {
			t.Fatalf("Unwrap() = %v, want %v", got, cause)
		}
	})

	t.Run("errors.Is traverses chain", func(t *testing.T) {
		sentinel := fmt.Errorf("sentinel")
		e := NewTencentCloudSDKErrorWithCause("c", "m", "", sentinel)
		if !errors.Is(e, sentinel) {
			t.Fatalf("errors.Is(e, sentinel) = false, want true")
		}
	})

	t.Run("errors.As traverses chain to net.Error-like", func(t *testing.T) {
		type timeoutErr struct{ msg string }
		// plain error type to verify As plumbing reaches through the wrapper
		cause := &timeoutErrImpl{msg: "boom"}
		e := NewTencentCloudSDKErrorWithCause("c", "m", "", cause)
		var got *timeoutErrImpl
		if !errors.As(e, &got) {
			t.Fatalf("errors.As(e, &timeoutErrImpl) = false, want true")
		}
		if got.msg != "boom" {
			t.Fatalf("got.msg = %q, want %q", got.msg, "boom")
		}
	})
}

func TestNewTencentCloudSDKError_NoCause(t *testing.T) {
	e := NewTencentCloudSDKError("c", "m", "r")
	sdk, ok := e.(*TencentCloudSDKError)
	if !ok {
		t.Fatalf("expected *TencentCloudSDKError, got %T", e)
	}
	if sdk.Unwrap() != nil {
		t.Fatalf("Unwrap() = %v, want nil for error created without cause", sdk.Unwrap())
	}
	if got := e.Error(); got != "[TencentCloudSDKError] Code=c, Message=m, RequestId=r" {
		t.Fatalf("Error() = %q, want no cause suffix", got)
	}
}

func TestNewTencentCloudSDKErrorWithCause(t *testing.T) {
	cause := fmt.Errorf("root")
	e := NewTencentCloudSDKErrorWithCause("c", "m", "r", cause)
	sdk, ok := e.(*TencentCloudSDKError)
	if !ok {
		t.Fatalf("expected *TencentCloudSDKError, got %T", e)
	}
	if sdk.Unwrap() != cause {
		t.Fatalf("Unwrap() = %v, want %v", sdk.Unwrap(), cause)
	}
	// Error() must include cause text.
	if got := e.Error(); got != "[TencentCloudSDKError] Code=c, Message=m, RequestId=r: root" {
		t.Fatalf("Error() = %q, want cause suffix", got)
	}
}

// timeoutErrImpl is a plain error type used to verify errors.As reaches
// through TencentCloudSDKError's Unwrap chain.
type timeoutErrImpl struct{ msg string }

func (e *timeoutErrImpl) Error() string { return e.msg }
