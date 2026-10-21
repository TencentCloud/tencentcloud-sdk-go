//go:build tencentcloud_common_number_precision
// +build tencentcloud_common_number_precision

package common

import (
	"fmt"

	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

const (
	octetStream = "application/octet-stream"
)

type actionParameters interface{}

type CommonRequest struct {
	*BaseRequest
	actionParameters
}

func NewCommonRequest(service, version, action string) (request *CommonRequest) {
	request = &CommonRequest{
		BaseRequest:      &BaseRequest{},
		actionParameters: map[string]interface{}{},
	}
	request.Init().WithApiInfo(service, version, action)
	return
}

// SetActionParameters set common request's actionParameters to your data.
// note: your data Must be a json-formatted string or byte array or map[string]interface{}
// note: you could not call SetActionParameters and SetOctetStreamParameters at once
func (cr *CommonRequest) SetActionParameters(data interface{}) error {
	if data == nil {
		return nil
	}
	switch v := data.(type) {
	case []byte:
		return cr.setRawParameters(v)
	case string:
		return cr.setRawParameters([]byte(v))
	case map[string]interface{}:
		cr.actionParameters = v
	default:
		msg := fmt.Sprintf("Invalid data type:%T, must be one of the following: []byte, string, map[string]interface{}", data)
		return tcerr.NewTencentCloudSDKError("ClientError.InvalidParameter", msg, "")
	}
	return nil
}

// setRawParameters validates data and stores.
func (cr *CommonRequest) setRawParameters(data []byte) error {
	if !isJSONObject(data) {
		msg := fmt.Sprintf("Not a valid json object: %s", string(data))
		return tcerr.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}

	raw := make(json.RawMessage, len(data))
	copy(raw, data)
	cr.actionParameters = raw
	return nil
}

func (cr *CommonRequest) IsOctetStream() bool {
	v, ok := cr.GetHeader()["Content-Type"]
	if !ok || v != octetStream {
		return false
	}

	switch parameter := cr.actionParameters.(type) {
	case map[string]interface{}:
		_, ok := parameter["OctetStreamBody"].([]byte)
		return ok
	default:
		return false
	}
}

func (cr *CommonRequest) SetHeader(header map[string]string) {
	if header == nil {
		return
	}
	if cr.BaseRequest == nil {
		cr.BaseRequest = &BaseRequest{}
	}
	cr.BaseRequest.SetHeader(header)
}

func (cr *CommonRequest) GetHeader() map[string]string {
	if cr.BaseRequest == nil {
		return nil
	}
	return cr.BaseRequest.GetHeader()
}

// SetOctetStreamParameters set request body to your data, and set head Content-Type to application/octet-stream
// note: you could not call SetActionParameters and SetOctetStreamParameters on the same request
func (cr *CommonRequest) SetOctetStreamParameters(header map[string]string, body []byte) {
	parameter := map[string]interface{}{}
	if header == nil {
		header = map[string]string{}
	}
	header["Content-Type"] = octetStream
	cr.SetHeader(header)
	parameter["OctetStreamBody"] = body
	cr.actionParameters = parameter
}

func (cr *CommonRequest) GetOctetStreamBody() []byte {
	if !cr.IsOctetStream() {
		return nil
	}

	switch parameter := cr.actionParameters.(type) {
	case map[string]interface{}:
		body, _ := parameter["OctetStreamBody"].([]byte)
		return body
	default:
		return nil
	}
}

func (cr *CommonRequest) MarshalJSON() ([]byte, error) {
	switch parameter := cr.actionParameters.(type) {
	case nil:
		return json.Marshal(nil)
	case json.RawMessage:
		return parameter, nil
	case map[string]interface{}:
		return json.Marshal(parameter)
	default:
		return nil, fmt.Errorf("unsupported action parameters type: %T", parameter)
	}
}
