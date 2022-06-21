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

package v20200902

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type SendCodeVoiceRequestParams struct {
	// 验证码，仅支持填写数字，实际播报语音时，会自动在数字前补充语音文本"您的验证码是"。
	CodeMessage *string `json:"CodeMessage,omitempty" name:"CodeMessage"`

	// 被叫手机号码，采用 e.164 标准，格式为+[国家或地区码][用户号码]。
	// 例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	CalledNumber *string `json:"CalledNumber,omitempty" name:"CalledNumber"`

	// 在[语音控制台](https://console.cloud.tencent.com/vms)添加应用后生成的实际SdkAppid，示例如1400006666。
	VoiceSdkAppid *string `json:"VoiceSdkAppid,omitempty" name:"VoiceSdkAppid"`

	// 播放次数，可选，最多3次，默认2次。
	PlayTimes *uint64 `json:"PlayTimes,omitempty" name:"PlayTimes"`

	// 用户的 session 内容，腾讯 server 回包中会原样返回。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`
}

type SendCodeVoiceRequest struct {
	*tchttp.BaseRequest
	
	// 验证码，仅支持填写数字，实际播报语音时，会自动在数字前补充语音文本"您的验证码是"。
	CodeMessage *string `json:"CodeMessage,omitempty" name:"CodeMessage"`

	// 被叫手机号码，采用 e.164 标准，格式为+[国家或地区码][用户号码]。
	// 例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	CalledNumber *string `json:"CalledNumber,omitempty" name:"CalledNumber"`

	// 在[语音控制台](https://console.cloud.tencent.com/vms)添加应用后生成的实际SdkAppid，示例如1400006666。
	VoiceSdkAppid *string `json:"VoiceSdkAppid,omitempty" name:"VoiceSdkAppid"`

	// 播放次数，可选，最多3次，默认2次。
	PlayTimes *uint64 `json:"PlayTimes,omitempty" name:"PlayTimes"`

	// 用户的 session 内容，腾讯 server 回包中会原样返回。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`
}

func (r *SendCodeVoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendCodeVoiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CodeMessage")
	delete(f, "CalledNumber")
	delete(f, "VoiceSdkAppid")
	delete(f, "PlayTimes")
	delete(f, "SessionContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendCodeVoiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendCodeVoiceResponseParams struct {
	// 语音验证码发送状态。
	SendStatus *SendStatus `json:"SendStatus,omitempty" name:"SendStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SendCodeVoiceResponse struct {
	*tchttp.BaseResponse
	Response *SendCodeVoiceResponseParams `json:"Response"`
}

func (r *SendCodeVoiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendCodeVoiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendStatus struct {
	// 标识本次发送 ID，标识一次下发记录。
	CallId *string `json:"CallId,omitempty" name:"CallId"`

	// 用户的 session 内容，腾讯 server 回包中会原样返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`
}

// Predefined struct for user
type SendTtsVoiceRequestParams struct {
	// 模板 ID，在控制台审核通过的模板 ID。
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 被叫手机号码，采用 e.164 标准，格式为+[国家或地区码][用户号码]。
	// 例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	CalledNumber *string `json:"CalledNumber,omitempty" name:"CalledNumber"`

	// 在[语音控制台](https://console.cloud.tencent.com/vms)添加应用后生成的实际SdkAppid，示例如1400006666。
	VoiceSdkAppid *string `json:"VoiceSdkAppid,omitempty" name:"VoiceSdkAppid"`

	// 模板参数，若模板没有参数，请提供为空数组。
	// 注：语音消息的内容长度不超过350字。
	TemplateParamSet []*string `json:"TemplateParamSet,omitempty" name:"TemplateParamSet"`

	// 播放次数，可选，最多3次，默认2次。
	PlayTimes *uint64 `json:"PlayTimes,omitempty" name:"PlayTimes"`

	// 用户的 session 内容，腾讯 server 回包中会原样返回。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`
}

type SendTtsVoiceRequest struct {
	*tchttp.BaseRequest
	
	// 模板 ID，在控制台审核通过的模板 ID。
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 被叫手机号码，采用 e.164 标准，格式为+[国家或地区码][用户号码]。
	// 例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	CalledNumber *string `json:"CalledNumber,omitempty" name:"CalledNumber"`

	// 在[语音控制台](https://console.cloud.tencent.com/vms)添加应用后生成的实际SdkAppid，示例如1400006666。
	VoiceSdkAppid *string `json:"VoiceSdkAppid,omitempty" name:"VoiceSdkAppid"`

	// 模板参数，若模板没有参数，请提供为空数组。
	// 注：语音消息的内容长度不超过350字。
	TemplateParamSet []*string `json:"TemplateParamSet,omitempty" name:"TemplateParamSet"`

	// 播放次数，可选，最多3次，默认2次。
	PlayTimes *uint64 `json:"PlayTimes,omitempty" name:"PlayTimes"`

	// 用户的 session 内容，腾讯 server 回包中会原样返回。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`
}

func (r *SendTtsVoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendTtsVoiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "CalledNumber")
	delete(f, "VoiceSdkAppid")
	delete(f, "TemplateParamSet")
	delete(f, "PlayTimes")
	delete(f, "SessionContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendTtsVoiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendTtsVoiceResponseParams struct {
	// 语音通知发送状态。
	SendStatus *SendStatus `json:"SendStatus,omitempty" name:"SendStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SendTtsVoiceResponse struct {
	*tchttp.BaseResponse
	Response *SendTtsVoiceResponseParams `json:"Response"`
}

func (r *SendTtsVoiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendTtsVoiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}