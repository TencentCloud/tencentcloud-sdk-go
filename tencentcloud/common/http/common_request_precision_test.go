//go:build tencentcloud_common_number_precision
// +build tencentcloud_common_number_precision

package common

import (
	"strings"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

const bigInt = "20220101123456789012"

func TestCommonRequest_RawJSONPreservesNumberPrecision(t *testing.T) {
	payload := `{"InstanceId":` + bigInt + `,"Nested":{"Offset":` + bigInt + `},"List":[` + bigInt + `]}`

	for _, tc := range []struct {
		name string
		data interface{}
	}{
		{"string", payload},
		{"bytes", []byte(payload)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			cr := NewCommonRequest("cvm", "2017-03-12", "DescribeInstances")
			if err := cr.SetActionParameters(tc.data); err != nil {
				t.Fatalf("SetActionParameters: %v", err)
			}

			got, err := json.Marshal(cr)
			if err != nil {
				t.Fatalf("Marshal: %v", err)
			}
			if n := strings.Count(string(got), bigInt); n != 3 {
				t.Fatalf("expected %s to survive in all 3 positions, got %d occurrence(s): %s",
					bigInt, n, got)
			}
		})
	}
}

func TestCommonRequest_RawJSONIsIndependentOfCallerBuffer(t *testing.T) {
	buf := []byte(`{"A":1}`)
	cr := NewCommonRequest("cvm", "2017-03-12", "DescribeInstances")
	if err := cr.SetActionParameters(buf); err != nil {
		t.Fatalf("SetActionParameters: %v", err)
	}

	// The caller still owns buf and may reuse it.
	copy(buf, `{"B":2}`)

	got, err := json.Marshal(cr)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(got) != `{"A":1}` {
		t.Fatalf("request payload aliases the caller's buffer: got %s", got)
	}
}

func TestCommonRequest_SetActionParametersReplaces(t *testing.T) {
	cr := NewCommonRequest("cvm", "2017-03-12", "DescribeInstances")

	if err := cr.SetActionParameters(`{"A":1}`); err != nil {
		t.Fatalf("SetActionParameters(raw): %v", err)
	}
	if err := cr.SetActionParameters(map[string]interface{}{"B": 2}); err != nil {
		t.Fatalf("SetActionParameters(map): %v", err)
	}
	got, err := json.Marshal(cr)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(got) != `{"B":2}` {
		t.Fatalf("map must fully replace a previous raw payload, got %s", got)
	}

	if err := cr.SetActionParameters(`{"C":3}`); err != nil {
		t.Fatalf("SetActionParameters(raw): %v", err)
	}
	got, err = json.Marshal(cr)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(got) != `{"C":3}` {
		t.Fatalf("raw payload must fully replace a previous map, got %s", got)
	}
}

// A map[string]interface{} carrying real Go integers was never lossy and must
// stay exact.
func TestCommonRequest_MapKeepsInt64(t *testing.T) {
	cr := NewCommonRequest("cvm", "2017-03-12", "DescribeInstances")
	if err := cr.SetActionParameters(map[string]interface{}{
		"InstanceId": int64(20220101123456789),
	}); err != nil {
		t.Fatalf("SetActionParameters: %v", err)
	}
	got, err := json.Marshal(cr)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(got) != `{"InstanceId":20220101123456789}` {
		t.Fatalf("unexpected payload: %s", got)
	}
}

func TestCommonRequest_RawJSONRejectsInvalidInput(t *testing.T) {
	for _, tc := range []struct {
		name string
		data interface{}
	}{
		{"truncated object", `{"a":"1"`},
		{"array", `[1,2]`},
		{"bare string", `"nope"`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			cr := NewCommonRequest("cvm", "2017-03-12", "DescribeInstances")
			if err := cr.SetActionParameters(tc.data); err == nil {
				t.Fatal("expected ClientError.ParseJsonError, got nil")
			}
		})
	}
}

// Raw JSON can never produce a []byte value, so it must not be mistaken for an
// octet-stream payload.
func TestCommonRequest_RawJSONIsNotOctetStream(t *testing.T) {
	cr := NewCommonRequest("cvm", "2017-03-12", "DescribeInstances")
	cr.SetHeader(map[string]string{"Content-Type": octetStream})
	if err := cr.SetActionParameters(`{"OctetStreamBody":"aGk="}`); err != nil {
		t.Fatalf("SetActionParameters: %v", err)
	}
	if cr.IsOctetStream() {
		t.Fatal("raw JSON must not be treated as an octet-stream body")
	}
	if cr.GetOctetStreamBody() != nil {
		t.Fatal("GetOctetStreamBody must be nil for a raw JSON payload")
	}
}

// actionParameters is a tagged union: it must hold exactly one of the two
// supported shapes at any time, and every setter must replace the other one.
func TestCommonRequest_ActionParametersShape(t *testing.T) {
	assertMap := func(t *testing.T, cr *CommonRequest) map[string]interface{} {
		t.Helper()
		switch parameter := cr.actionParameters.(type) {
		case map[string]interface{}:
			return parameter
		case json.RawMessage:
			t.Fatal("expected a parameter map, got a raw JSON payload")
		default:
			t.Fatalf("expected a parameter map, got %T", parameter)
		}
		return nil
	}
	assertRaw := func(t *testing.T, cr *CommonRequest) json.RawMessage {
		t.Helper()
		switch parameter := cr.actionParameters.(type) {
		case json.RawMessage:
			return parameter
		case map[string]interface{}:
			t.Fatal("expected a raw JSON payload, got a parameter map")
		default:
			t.Fatalf("expected a raw JSON payload, got %T", parameter)
		}
		return nil
	}

	t.Run("new request starts as an empty map", func(t *testing.T) {
		cr := NewCommonRequest("cvm", "2017-03-12", "DescribeInstances")
		if m := assertMap(t, cr); len(m) != 0 {
			t.Fatalf("expected an empty map, got %+v", m)
		}
		got, err := json.Marshal(cr)
		if err != nil {
			t.Fatalf("Marshal: %v", err)
		}
		if string(got) != "{}" {
			t.Fatalf("expected {}, got %s", got)
		}
	})

	t.Run("raw replaces map", func(t *testing.T) {
		cr := NewCommonRequest("cvm", "2017-03-12", "DescribeInstances")
		if err := cr.SetActionParameters(map[string]interface{}{"A": 1}); err != nil {
			t.Fatalf("SetActionParameters(map): %v", err)
		}
		assertMap(t, cr)
		if err := cr.SetActionParameters(`{"B":2}`); err != nil {
			t.Fatalf("SetActionParameters(raw): %v", err)
		}
		if raw := assertRaw(t, cr); string(raw) != `{"B":2}` {
			t.Fatalf("unexpected raw payload: %s", raw)
		}
	})

	t.Run("map replaces raw", func(t *testing.T) {
		cr := NewCommonRequest("cvm", "2017-03-12", "DescribeInstances")
		if err := cr.SetActionParameters([]byte(`{"A":1}`)); err != nil {
			t.Fatalf("SetActionParameters(raw): %v", err)
		}
		assertRaw(t, cr)
		if err := cr.SetActionParameters(map[string]interface{}{"B": 2}); err != nil {
			t.Fatalf("SetActionParameters(map): %v", err)
		}
		if m := assertMap(t, cr); len(m) != 1 || m["B"] != 2 {
			t.Fatalf("unexpected parameter map: %+v", m)
		}
	})

	t.Run("octet stream replaces raw", func(t *testing.T) {
		cr := NewCommonRequest("cvm", "2017-03-12", "DescribeInstances")
		if err := cr.SetActionParameters(`{"A":1}`); err != nil {
			t.Fatalf("SetActionParameters(raw): %v", err)
		}
		cr.SetOctetStreamParameters(nil, []byte("hi"))
		assertMap(t, cr)
		if !cr.IsOctetStream() {
			t.Fatal("expected an octet-stream request")
		}
		if got := cr.GetOctetStreamBody(); string(got) != "hi" {
			t.Fatalf("unexpected octet-stream body: %q", got)
		}
	})

	t.Run("raw replaces octet stream", func(t *testing.T) {
		cr := NewCommonRequest("cvm", "2017-03-12", "DescribeInstances")
		cr.SetOctetStreamParameters(nil, []byte("hi"))
		if err := cr.SetActionParameters(`{"A":1}`); err != nil {
			t.Fatalf("SetActionParameters(raw): %v", err)
		}
		assertRaw(t, cr)
		// The Content-Type header set by SetOctetStreamParameters survives, so
		// only the payload shape can rule the octet-stream path out.
		if cr.GetHeader()["Content-Type"] != octetStream {
			t.Fatal("expected the octet-stream Content-Type to survive")
		}
		if cr.IsOctetStream() {
			t.Fatal("a raw JSON payload must not be treated as an octet-stream body")
		}
		if cr.GetOctetStreamBody() != nil {
			t.Fatal("GetOctetStreamBody must be nil for a raw JSON payload")
		}
	})

	t.Run("rejected input keeps the previous payload", func(t *testing.T) {
		cr := NewCommonRequest("cvm", "2017-03-12", "DescribeInstances")
		if err := cr.SetActionParameters(map[string]interface{}{"A": 1}); err != nil {
			t.Fatalf("SetActionParameters(map): %v", err)
		}
		if err := cr.SetActionParameters(`[1,2]`); err == nil {
			t.Fatal("expected ClientError.ParseJsonError, got nil")
		}
		if err := cr.SetActionParameters(123); err == nil {
			t.Fatal("expected ClientError.InvalidParameter, got nil")
		}
		if m := assertMap(t, cr); len(m) != 1 || m["A"] != 1 {
			t.Fatalf("a rejected payload must not clobber the previous one, got %+v", m)
		}
	})

	t.Run("nil data keeps the previous payload", func(t *testing.T) {
		cr := NewCommonRequest("cvm", "2017-03-12", "DescribeInstances")
		if err := cr.SetActionParameters(`{"A":1}`); err != nil {
			t.Fatalf("SetActionParameters(raw): %v", err)
		}
		if err := cr.SetActionParameters(nil); err != nil {
			t.Fatalf("SetActionParameters(nil): %v", err)
		}
		if raw := assertRaw(t, cr); string(raw) != `{"A":1}` {
			t.Fatalf("unexpected raw payload: %s", raw)
		}
	})
}

// GetOctetStreamBody must not panic when the well-known key holds a value of
// an unexpected type: the assertion is guarded by IsOctetStream.
func TestCommonRequest_OctetStreamBodyWrongType(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			t.Fatalf("panic on GetOctetStreamBody: %+v", e)
		}
	}()

	cr := &CommonRequest{
		BaseRequest: &BaseRequest{
			header: map[string]string{"Content-Type": octetStream},
		},
		actionParameters: map[string]interface{}{
			"OctetStreamBody": "not bytes",
		},
	}
	if cr.GetOctetStreamBody() != nil {
		t.Fatal("GetOctetStreamBody must be nil when the value is not []byte")
	}
}
