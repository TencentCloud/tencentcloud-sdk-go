//go:build tencentcloud_common_number_precision
// +build tencentcloud_common_number_precision

package common

import (
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

func TestCommonResponse_GetBodyPreservesNumberPrecision(t *testing.T) {
	body := `{"Response":{"Total":` + bigInt + `,"RequestId":"req-1"}}`

	resp := NewCommonResponse()
	if err := json.Unmarshal([]byte(body), &resp); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if got := string(resp.GetBody()); got != body {
		t.Fatalf("GetBody must return the payload verbatim:\n want %s\n got  %s", body, got)
	}
}

func TestCommonResponse_GetBodyIsIndependentOfDecoderBuffer(t *testing.T) {
	buf := []byte(`{"Response":{"RequestId":"req-1"}}`)
	resp := NewCommonResponse()
	if err := resp.UnmarshalJSON(buf); err != nil {
		t.Fatalf("UnmarshalJSON: %v", err)
	}
	for i := range buf {
		buf[i] = 'x'
	}
	if got := string(resp.GetBody()); got != `{"Response":{"RequestId":"req-1"}}` {
		t.Fatalf("GetBody aliases the decoder buffer: got %s", got)
	}
}

func TestCommonResponse_GetBodyDefaultsToEmptyObject(t *testing.T) {
	if got := string(NewCommonResponse().GetBody()); got != "{}" {
		t.Fatalf("expected {} for an unpopulated response, got %s", got)
	}

	// make() produces a non-nil zero-length slice, so GetBody must key off
	// length rather than nil.
	resp := NewCommonResponse()
	resp.body = make(json.RawMessage, 0)
	if got := string(resp.GetBody()); got != "{}" {
		t.Fatalf("expected {} for a zero-length body, got %q", got)
	}
}

func TestCommonResponse_RejectsInvalidJSON(t *testing.T) {
	resp := NewCommonResponse()
	if err := resp.UnmarshalJSON([]byte(`{"Response":`)); err == nil {
		t.Fatal("expected ClientError.ParseJsonError, got nil")
	}
}
