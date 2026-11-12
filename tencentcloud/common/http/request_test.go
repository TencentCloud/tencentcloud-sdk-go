package common

import (
	"strconv"
	"testing"
)

func TestBaseRequest_Header(t *testing.T) {
	r := &BaseRequest{}

	const (
		traceKey = "X-TC-TraceId"
		traceVal = "ffe0c072-8a5d-4e17-8887-a8a60252abca"
	)

	if r.GetHeader() != nil {
		t.Fatal("default header MUST be nil")
	}

	r.SetHeader(nil)
	if r.GetHeader() != nil {
		t.Fatal("SetHeader(nil) MUST not replace nil map with empty map")
	}

	r.SetHeader(map[string]string{traceKey: traceVal})
	if r.GetHeader()[traceKey] != traceVal {
		t.Fatal("SetHeader failed")
	}

	r.SetHeader(nil)
	if r.GetHeader() == nil {
		t.Fatal("SetHeader(nil) MUST not overwrite existing header (for backward compatibility)")
	}

	if r.GetHeader()[traceKey] != traceVal {
		t.Fatal("SetHeader(nil) MUST not overwrite existing header (for backward compatibility)")
	}
}

func TestRandomNonce(t *testing.T) {
	seen := make(map[string]struct{}, 100)
	for i := 0; i < 100; i++ {
		nonce := randomNonce()
		if _, err := strconv.ParseInt(nonce, 10, 64); err != nil {
			t.Fatalf("nonce %q is not a valid non-negative int64: %v", nonce, err)
		}
		if _, ok := seen[nonce]; ok {
			t.Fatalf("duplicate nonce generated: %q", nonce)
		}
		seen[nonce] = struct{}{}
	}
}
