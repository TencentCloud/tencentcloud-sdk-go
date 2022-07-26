// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20220613

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type BatchContent struct {
	// 消息体
	Body *string `json:"Body,omitempty" name:"Body"`

	// 消息的键名
	Key *string `json:"Key,omitempty" name:"Key"`
}

// Predefined struct for user
type SendMessageRequestParams struct {
	// 接入资源ID
	DataHubId *string `json:"DataHubId,omitempty" name:"DataHubId"`

	// 批量消息
	Message []*BatchContent `json:"Message,omitempty" name:"Message"`
}

type SendMessageRequest struct {
	*tchttp.BaseRequest
	
	// 接入资源ID
	DataHubId *string `json:"DataHubId,omitempty" name:"DataHubId"`

	// 批量消息
	Message []*BatchContent `json:"Message,omitempty" name:"Message"`
}

func (r *SendMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataHubId")
	delete(f, "Message")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendMessageResponseParams struct {
	// 消息ID
	MessageId []*string `json:"MessageId,omitempty" name:"MessageId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SendMessageResponse struct {
	*tchttp.BaseResponse
	Response *SendMessageResponseParams `json:"Response"`
}

func (r *SendMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}