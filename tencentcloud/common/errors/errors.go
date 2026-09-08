package errors

import (
	"fmt"
)

type TencentCloudSDKError struct {
	Code      string
	Message   string
	RequestId string
	cause     error
}

func (e *TencentCloudSDKError) Error() string {
	var cause string
	if e.cause != nil {
		cause = ": " + e.cause.Error()
	}
	if e.RequestId == "" {
		return fmt.Sprintf("[TencentCloudSDKError] Code=%s, Message=%s%s", e.Code, e.Message, cause)
	}
	return fmt.Sprintf("[TencentCloudSDKError] Code=%s, Message=%s, RequestId=%s%s", e.Code, e.Message, e.RequestId, cause)
}

func (e *TencentCloudSDKError) Unwrap() error {
	return e.cause
}

func NewTencentCloudSDKError(code, message, requestId string) error {
	return &TencentCloudSDKError{
		Code:      code,
		Message:   message,
		RequestId: requestId,
	}
}

// NewTencentCloudSDKErrorWithCause wraps cause as a TencentCloudSDKError so
// callers get a typed SDK error while the original error remains reachable
// via errors.Is / errors.As / Unwrap. Use this instead of NewTencentCloudSDKError
// whenever you are converting an existing Go error (network error, json parse
// error, etc.) — otherwise the original error type is lost.
func NewTencentCloudSDKErrorWithCause(code, message, requestId string, cause error) error {
	return &TencentCloudSDKError{
		Code:      code,
		Message:   message,
		RequestId: requestId,
		cause:     cause,
	}
}

func (e *TencentCloudSDKError) GetCode() string {
	return e.Code
}

func (e *TencentCloudSDKError) GetMessage() string {
	return e.Message
}

func (e *TencentCloudSDKError) GetRequestId() string {
	return e.RequestId
}
