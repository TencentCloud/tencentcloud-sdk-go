package common

import (
	stdJson "encoding/json"
	commonJson "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

var (
	commonJsonMarshal = commonJson.Marshal
	stdJsonMarshal    = stdJson.Marshal
)

func marshalFunc(omitNil bool) func(interface{}) ([]byte, error) {
	if omitNil {
		return commonJsonMarshal
	}
	return stdJsonMarshal
}
