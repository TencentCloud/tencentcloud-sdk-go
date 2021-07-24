package common

import (
	"encoding/json"
	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"testing"
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

func TestCommonRequest_IsOctetStream(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			t.Fatalf("panic on IsOctetStream: %+v", e)
		}
	}()
	cr1 := &CommonRequest{header: map[string]string{
		"Content-Type": "text/plain",
	}}
	cr2 := &CommonRequest{
		header: map[string]string{
			"Content-Type": octetStream,
		},
		actionParameters: map[string]interface{}{
			"octetstreambody": []byte{},
		},
	}
	cr3 := &CommonRequest{
		header: map[string]string{
			"Content-Type": octetStream,
		},
		actionParameters: map[string]interface{}{
			"OctetStreamBody": []string{},
		},
	}
	cr4 := &CommonRequest{
		header: map[string]string{
			"Content-Type": octetStream,
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
