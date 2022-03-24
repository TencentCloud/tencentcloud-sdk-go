package common

import "testing"

func TestBaseRequest_Header(t *testing.T) {
	r := &BaseRequest{}

	if r.GetHeader() != nil {
		t.Fatal("default header MUST be nil")
	}

	// SetHeader(nil) MUST not replace nil map with empty map
	r.SetHeader(nil)
	if r.GetHeader() != nil {
		t.Fatal("SetHeader(nil) MUST not replace nil map with empty map")
	}

	r.SetHeader(map[string]string{"X-Key-1": "X-Value-1"})
	if r.GetHeader()["X-Key-1"] != "X-Value-1" {
		t.Fatal("SetHeader failed")
	}

	r.SetHeader(nil)
	if r.GetHeader() == nil {
		t.Fatal("SetHeader(nil) MUST not overwrite existing header (for backward compatibility)")
	}

	if r.GetHeader()["X-Key-1"] != "X-Value-1" {
		t.Fatal("SetHeader(nil) MUST not overwrite existing header (for backward compatibility)")
	}
}
