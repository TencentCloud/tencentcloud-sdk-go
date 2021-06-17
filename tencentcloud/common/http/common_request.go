package common

import (
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
)

type actionRequest map[string]interface{}

type CommonRequest struct {
	*BaseRequest
	*actionRequest
}

func NewCommonRequest(service, version, action string) (request *CommonRequest) {
	request = &CommonRequest{
		BaseRequest:   &BaseRequest{},
		actionRequest: &actionRequest{},
	}
	request.Init().WithApiInfo(service, version, action)
	return
}

func (cr *CommonRequest) SetActionRequest(data interface{}) error {
	if data == nil {
		return nil
	}
	switch data.(type) {
	case []byte:
		if err := json.Unmarshal(data.([]byte), cr.actionRequest); err != nil {
			msg := fmt.Sprintf("Fail to parse contenst %s to json,because: %s", data.([]byte), err)
			return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
		} else {
			return nil
		}
	case string:
		if err := json.Unmarshal([]byte(data.(string)), cr.actionRequest); err != nil {
			msg := fmt.Sprintf("Fail to parse contenst %s to json,because: %s", data.(string), err)
			return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
		} else {
			return nil
		}
	case map[string]interface{}:
		*cr.actionRequest = data.(map[string]interface{})
		return nil
	default:
		msg := fmt.Sprintf("Invalid data type:%T, must be one of the following: []byte, string, map[string]interface{}", data)
		return errors.NewTencentCloudSDKError("ClientError.InvalidParameter", msg, "")
	}
}

func (cr *CommonRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(cr.actionRequest)
}

func (cb *actionRequest) Raw() []byte {
	raw, _ := json.Marshal(cb)
	return raw
}
