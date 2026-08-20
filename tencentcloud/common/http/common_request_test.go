package common

import (
	"testing"

	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

func TestCommonRequest_SetActionParameters(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			t.Fatalf("panic on SetActionParameters: %+v", e)
		}
	}()
	testCase := []struct {
		data    interface{}
		errCode string
	}{
		{[]byte("{\"a\":\"1\"}"), ""},
		{"{\"a\":\"1\"}", ""},
		{map[string]interface{}{"a": "1"}, ""},
		{[]byte("{\"a\":\"1\""), "ClientError.ParseJsonError"},
		{123, "ClientError.InvalidParameter"},
	}
	cr := &CommonRequest{}
	for _, tc := range testCase {
		err := cr.SetActionParameters(tc.data)
		if err != nil {
			if te, ok := err.(*tcerr.TencentCloudSDKError); ok {
				if te.GetCode() != tc.errCode {
					t.Fatalf("SetActionParameters failed: expected %+v, got %+v", tc.errCode, te.GetCode())
				}
			} else {
				t.Fatalf("SetActionParameters failed: expected %+v, got %T", "TencentCloudSDKError", err)
			}
		} else {
			if tc.errCode != "" {
				t.Fatalf("SetActionParameters failed: expected %+v, got %+v", tc.errCode, "")
			}
		}
	}
}

func TestCommonRequest_JSONMarshal(t *testing.T) {
	crn := NewCommonRequest("cvm", "2017-03-12", "DescribeInstances")
	_ = crn.SetActionParameters(map[string]interface{}{
		"a": 1,
		"b": map[string]interface{}{
			"b1": 2,
			"b2": "b2",
		},
	})

	bytes, err := json.MarshalIndent(crn, "", "\t")
	if err != nil || len(bytes) == 0 {
		t.Fatal(err)
	}
}

func TestCommonRequest_SetActionParametersNumberBoundary(t *testing.T) {
	testCase := []struct {
		name         string
		data         interface{}
		errCode      string
		expectedBody string
	}{
		{
			name:         "int64 max in string",
			data:         `{"Int":9223372036854775807}`,
			expectedBody: `{"Int":9223372036854775807}`,
		},
		{
			name:         "uint64 max",
			data:         `{"Uint":18446744073709551615}`,
			expectedBody: `{"Uint":18446744073709551615}`,
		},
		{
			name:         "int64 boundary array",
			data:         `{"Arr":[9223372036854775807,-9223372036854775808]}`,
			expectedBody: `{"Arr":[9223372036854775807,-9223372036854775808]}`,
		},
		{
			name:         "int64 max in bytes",
			data:         []byte(`{"Int":9223372036854775807}`),
			expectedBody: `{"Int":9223372036854775807}`,
		},
		{
			name:         "float beyond float64 range",
			data:         `{"Float":1.79E309}`,
			expectedBody: `{"Float":1.79E309}`,
		},
		{
			name:         "negative float beyond float64 range",
			data:         `{"Float":-1.79E309}`,
			expectedBody: `{"Float":-1.79E309}`,
		},
		{
			name:         "normal numbers unchanged",
			data:         `{"a":1,"b":1.5,"c":-2}`,
			expectedBody: `{"a":1,"b":1.5,"c":-2}`,
		},
		{
			name:         "nested object keeps big int",
			data:         `{"Outer":{"Int":9223372036854775807}}`,
			expectedBody: `{"Outer":{"Int":9223372036854775807}}`,
		},
		{
			name:    "trailing data still rejected",
			data:    `{"a":"1"} trailing`,
			errCode: "ClientError.ParseJsonError",
		},
		{
			name:    "empty string still rejected",
			data:    ``,
			errCode: "ClientError.ParseJsonError",
		},
		{
			name:    "truncated json still rejected",
			data:    `{"a":"1"`,
			errCode: "ClientError.ParseJsonError",
		},
	}
	for _, tc := range testCase {
		cr := NewCommonRequest("cvm", "2017-03-12", "DescribeInstances")
		err := cr.SetActionParameters(tc.data)
		if tc.errCode != "" {
			if err == nil {
				t.Fatalf("[%s] expected error %s, got nil", tc.name, tc.errCode)
			}
			te, ok := err.(*tcerr.TencentCloudSDKError)
			if !ok {
				t.Fatalf("[%s] expected TencentCloudSDKError, got %T", tc.name, err)
			}
			if te.GetCode() != tc.errCode {
				t.Fatalf("[%s] expected error code %s, got %s", tc.name, tc.errCode, te.GetCode())
			}
			continue
		}
		if err != nil {
			t.Fatalf("[%s] unexpected error: %v", tc.name, err)
		}
		body, err := cr.MarshalJSON()
		if err != nil {
			t.Fatalf("[%s] marshal error: %v", tc.name, err)
		}
		if string(body) != tc.expectedBody {
			t.Fatalf("[%s] marshal mismatch:\n expected: %s\n got:      %s", tc.name, tc.expectedBody, body)
		}
	}
}

func TestCommonRequest_IsOctetStream(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			t.Fatalf("panic on IsOctetStream: %+v", e)
		}
	}()
	cr1 := &CommonRequest{
		BaseRequest: &BaseRequest{
			header: map[string]string{
				"Content-Type": "text/plain",
			},
		}}
	cr2 := &CommonRequest{
		BaseRequest: &BaseRequest{
			header: map[string]string{
				"Content-Type": octetStream,
			},
		},
		actionParameters: map[string]interface{}{
			"octetstreambody": []byte{},
		},
	}
	cr3 := &CommonRequest{
		BaseRequest: &BaseRequest{
			header: map[string]string{
				"Content-Type": octetStream,
			},
		},
		actionParameters: map[string]interface{}{
			"OctetStreamBody": []string{},
		},
	}
	cr4 := &CommonRequest{
		BaseRequest: &BaseRequest{
			header: map[string]string{
				"Content-Type": octetStream,
			},
		},
		actionParameters: map[string]interface{}{
			"OctetStreamBody": []byte{},
		},
	}

	testCase := map[*CommonRequest]bool{
		cr1: false,
		cr2: false,
		cr3: false,
		cr4: true,
	}
	for cr, expected := range testCase {
		if val := cr.IsOctetStream(); val != expected {
			t.Fatalf("IsOctetStream failed: expected %+v, got %+v", expected, val)
		}
	}
}

func TestCommonRequest_SetOctetStreamParameters(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			t.Fatalf("panic on SetOctetStreamParameters: %+v", e)
		}
	}()
	type param struct {
		header map[string]string
		body   []byte
	}
	p1 := &param{
		header: map[string]string{
			"Content-Type": "text/plain",
		},
		body: []byte{},
	}
	p2 := &param{
		header: map[string]string{
			"Content-Type": octetStream,
		},
		body: []byte{},
	}
	testCase := map[*param]bool{
		p1: true,
		p2: true,
	}
	cr := &CommonRequest{}
	for p, wanted := range testCase {
		cr.SetOctetStreamParameters(p.header, p.body)
		if val := cr.IsOctetStream(); val != wanted {
			t.Fatalf("SetOctetStreamParameters failed: expected %+v, got %+v", wanted, val)
		}
	}
}

func TestCommonRequest_Header(t *testing.T) {
	r := &CommonRequest{}

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
