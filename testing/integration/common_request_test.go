package integration

import (
	"encoding/json"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"testing"
)

func TestCommonRequestMarshal(t *testing.T) {
	crn := common.NewCommonRequest("cvm", "2017-03-12", "DescribeInstances")
	_ = crn.SetActionParameters(map[string]interface{}{
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
}
