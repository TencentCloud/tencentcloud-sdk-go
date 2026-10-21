//go:build tencentcloud_common_number_precision
// +build tencentcloud_common_number_precision

package common

import (
	"bytes"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type CommonResponse struct {
	*BaseResponse

	body json.RawMessage
}

func NewCommonResponse() (response *CommonResponse) {
	response = &CommonResponse{
		BaseResponse: &BaseResponse{},
	}
	return
}

func (r *CommonResponse) UnmarshalJSON(data []byte) error {
	if !isJSONObject(data) {
		msg := fmt.Sprintf("Not a valid json object: %s", string(data))
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}

	r.body = make(json.RawMessage, len(data))
	copy(r.body, data)
	return nil
}

func (r *CommonResponse) GetBody() []byte {
	// json.Marshal will also return "{}" if r.body is empty
	if len(r.body) == 0 {
		return []byte("{}")
	}
	return r.body
}

// isJSONObject checks if data is a valid JSON object
// use json.Valid instead of json.Unmarshal to avoid allocating unnecessary memory
// json.Valid returns true if data is a valid JSON array or single value
// add TrimLeft + trimmed[0] != '{' to check if data is a valid JSON object
func isJSONObject(data []byte) bool {
	trimmed := bytes.TrimLeft(data, " \t\r\n")
	if len(trimmed) == 0 || trimmed[0] != '{' {
		return false
	}
	return json.Valid(data)
}
