package common

import (
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
)

type actionParameters map[string]interface{}

type CommonRequest struct {
	*BaseRequest
	*actionParameters
}

func NewCommonRequest(service, version, action string) (request *CommonRequest) {
	request = &CommonRequest{
		BaseRequest:      &BaseRequest{},
		actionParameters: &actionParameters{},
	}
	request.Init().WithApiInfo(service, version, action)
	return
}

func (cr *CommonRequest) SetActionParameters(data interface{}) error {
	if data == nil {
		return nil
	}
	switch data.(type) {
	case []byte:
		if err := json.Unmarshal(data.([]byte), cr.actionParameters); err != nil {
			msg := fmt.Sprintf("Fail to parse contenst %s to json,because: %s", data.([]byte), err)
			return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
		} else {
			return nil
		}
	case string:
		if err := json.Unmarshal([]byte(data.(string)), cr.actionParameters); err != nil {
			msg := fmt.Sprintf("Fail to parse contenst %s to json,because: %s", data.(string), err)
			return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
		} else {
			return nil
		}
	case map[string]interface{}:
		*cr.actionParameters = data.(map[string]interface{})
		return nil
	default:
		msg := fmt.Sprintf("Invalid data type:%T, must be one of the following: []byte, string, map[string]interface{}", data)
		return errors.NewTencentCloudSDKError("ClientError.InvalidParameter", msg, "")
	}
}

func (cr *CommonRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(cr.actionParameters)
}
