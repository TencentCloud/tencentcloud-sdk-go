package common

import (
	"encoding/json"
	"testing"
)

func TestCommonRequestMarshal(t *testing.T) {
	crn := NewCommonRequest("cvm", "2017-03-12", "DescribeInstances")
	_ = crn.SetActionRequest(map[string]interface{}{
		"a": 1,
		"b": map[string]interface{}{
			"b1": 2,
			"b2": "b2",
		},
	})

	bytes, err := json.MarshalIndent(crn, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}
