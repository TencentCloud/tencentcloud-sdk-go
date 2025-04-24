package json

import (
	"bytes"
	"testing"
)

type omitBehaviourTestType struct {
	F1 *int   `json:"F1,omitnil,omitempty"`
	F2 int    `json:"F2,omitnil,omitempty"`
	F3 []*int `json:"F3,omitnil,omitempty"`
}

func TestOmitBehaviourOmitNil(t *testing.T) {
	OmitBehaviour = OmitNil
	var iVar int
	testCases := []struct {
		F1    *int
		F2    int
		F3    []*int
		HasF1 bool
		HasF2 bool
		HasF3 bool
	}{
		{
			F1:    nil,
			F2:    0,
			F3:    nil,
			HasF1: false,
			HasF2: true,
			HasF3: false,
		},
		{
			F1:    &iVar,
			F2:    0,
			F3:    nil,
			HasF1: true,
			HasF2: true,
			HasF3: false,
		},
		{
			F1:    nil,
			F2:    1,
			F3:    nil,
			HasF1: false,
			HasF2: true,
			HasF3: false,
		},
		{
			F1:    nil,
			F2:    0,
			F3:    []*int{},
			HasF1: false,
			HasF2: true,
			HasF3: true,
		},
	}

	for i, testCase := range testCases {
		d1, err := Marshal(omitBehaviourTestType{
			F1: testCase.F1,
			F2: testCase.F2,
			F3: testCase.F3,
		})
		if err != nil {
			t.Fatal(err)
		}

		hasF1 := bytes.Index(d1, []byte("F1")) >= 0
		hasF2 := bytes.Index(d1, []byte("F2")) >= 0
		hasF3 := bytes.Index(d1, []byte("F3")) >= 0
		if hasF1 != testCase.HasF1 {
			t.Fatalf("i=%d TestCase=%v, HasF1=%v, HasF2=%v, HasF3=%v", i, testCase, hasF1, hasF2, hasF3)
		}
		if hasF2 != testCase.HasF2 {
			t.Fatalf("i=%d TestCase=%v, HasF1=%v, HasF2=%v, HasF3=%v", i, testCase, hasF1, hasF2, hasF3)
		}
		if hasF3 != testCase.HasF3 {
			t.Fatalf("i=%d TestCase=%v, HasF1=%v, HasF2=%v, HasF3=%v", i, testCase, hasF1, hasF2, hasF3)
		}
	}
}

func TestOmitBehaviourOmitEmpty(t *testing.T) {
	OmitBehaviour = OmitEmpty
	var iVar int
	testCases := []struct {
		F1    *int
		F2    int
		F3    []*int
		HasF1 bool
		HasF2 bool
		HasF3 bool
	}{
		{
			F1:    nil,
			F2:    0,
			F3:    nil,
			HasF1: false,
			HasF2: false,
			HasF3: false,
		},
		{
			F1:    &iVar,
			F2:    0,
			F3:    nil,
			HasF1: true,
			HasF2: false,
			HasF3: false,
		},
		{
			F1:    nil,
			F2:    1,
			F3:    nil,
			HasF1: false,
			HasF2: true,
			HasF3: false,
		},
		{
			F1:    nil,
			F2:    0,
			F3:    []*int{},
			HasF1: false,
			HasF2: false,
			HasF3: false,
		},
	}

	for i, testCase := range testCases {
		d1, err := Marshal(omitBehaviourTestType{
			F1: testCase.F1,
			F2: testCase.F2,
			F3: testCase.F3,
		})
		if err != nil {
			t.Fatal(err)
		}

		hasF1 := bytes.Index(d1, []byte("F1")) >= 0
		hasF2 := bytes.Index(d1, []byte("F2")) >= 0
		hasF3 := bytes.Index(d1, []byte("F3")) >= 0
		if hasF1 != testCase.HasF1 {
			t.Fatalf("i=%d TestCase=%v, HasF1=%v, HasF2=%v, HasF3=%v", i, testCase, hasF1, hasF2, hasF3)
		}
		if hasF2 != testCase.HasF2 {
			t.Fatalf("i=%d TestCase=%v, HasF1=%v, HasF2=%v, HasF3=%v", i, testCase, hasF1, hasF2, hasF3)
		}
		if hasF3 != testCase.HasF3 {
			t.Fatalf("i=%d TestCase=%v, HasF1=%v, HasF2=%v, HasF3=%v", i, testCase, hasF1, hasF2, hasF3)
		}
	}
}
