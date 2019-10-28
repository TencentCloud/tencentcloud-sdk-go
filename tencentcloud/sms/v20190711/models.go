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

package v20190711

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type PullSmsReplyStatus struct {

	// 通道扩展码，默认没有开通（需要填空）
	ExtendCode *string `json:"ExtendCode,omitempty" name:"ExtendCode"`

	// 国家（或地区）码
	NationCode *string `json:"NationCode,omitempty" name:"NationCode"`

	// 手机号码（ e.164 标准）
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 短信签名
	Sign *string `json:"Sign,omitempty" name:"Sign"`

	// 用户回复的内容
	ReplyContent *string `json:"ReplyContent,omitempty" name:"ReplyContent"`

	// 回复时间，UNIX 时间戳（单位：秒）
	ReplyTime *uint64 `json:"ReplyTime,omitempty" name:"ReplyTime"`
}

type PullSmsReplyStatusByPhoneNumberRequest struct {
	*tchttp.BaseRequest

	// 拉取起始时间，UNIX 时间戳（时间：秒）
	SendDateTime *string `json:"SendDateTime,omitempty" name:"SendDateTime"`

	// 偏移量
	// 注：目前固定设置为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 拉取最大条数，最多 100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下发目的手机号码，依据 e.164 标准为：+[国家（或地区）码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 短信SdkAppid在[短信控制台](https://console.cloud.tencent.com/sms/smslist) 添加应用后生成的实际SdkAppid,示例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`
}

func (r *PullSmsReplyStatusByPhoneNumberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullSmsReplyStatusByPhoneNumberRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullSmsReplyStatusByPhoneNumberResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 回复状态响应集合
		PullSmsReplyStatusSet []*PullSmsReplyStatus `json:"PullSmsReplyStatusSet,omitempty" name:"PullSmsReplyStatusSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PullSmsReplyStatusByPhoneNumberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullSmsReplyStatusByPhoneNumberResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullSmsReplyStatusRequest struct {
	*tchttp.BaseRequest

	// 拉取最大条数，最多100条
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 短信SdkAppid在[短信控制台](https://console.cloud.tencent.com/sms/smslist) 添加应用后生成的实际SdkAppid,示例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`
}

func (r *PullSmsReplyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullSmsReplyStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullSmsReplyStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 回复状态响应集合
		PullSmsReplyStatusSet []*PullSmsReplyStatus `json:"PullSmsReplyStatusSet,omitempty" name:"PullSmsReplyStatusSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PullSmsReplyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullSmsReplyStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullSmsSendStatus struct {

	// 用户实际接收到短信的时间
	UserReceiveTime *string `json:"UserReceiveTime,omitempty" name:"UserReceiveTime"`

	// 国家（或地区）码
	NationCode *string `json:"NationCode,omitempty" name:"NationCode"`

	// 手机号码（ e.164 标准）
	PurePhoneNumber *string `json:"PurePhoneNumber,omitempty" name:"PurePhoneNumber"`

	// 手机号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 本次发送标识 ID
	SerialNo *string `json:"SerialNo,omitempty" name:"SerialNo"`

	// 实际是否收到短信接收状态，SUCCESS（成功）、FAIL（失败）
	ReportStatus *string `json:"ReportStatus,omitempty" name:"ReportStatus"`

	// 用户接收短信状态描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type PullSmsSendStatusByPhoneNumberRequest struct {
	*tchttp.BaseRequest

	// 拉取起始时间，UNIX 时间戳（时间：秒）
	SendDateTime *string `json:"SendDateTime,omitempty" name:"SendDateTime"`

	// 偏移量
	// 注：目前固定设置为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 拉取最大条数，最多 100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下发目的手机号码，依据 e.164 标准为：+[国家（或地区）码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 短信SdkAppid在[短信控制台](https://console.cloud.tencent.com/sms/smslist) 添加应用后生成的实际SdkAppid,示例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`
}

func (r *PullSmsSendStatusByPhoneNumberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullSmsSendStatusByPhoneNumberRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullSmsSendStatusByPhoneNumberResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 下发状态响应集合
		PullSmsSendStatusSet []*PullSmsSendStatus `json:"PullSmsSendStatusSet,omitempty" name:"PullSmsSendStatusSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PullSmsSendStatusByPhoneNumberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullSmsSendStatusByPhoneNumberResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullSmsSendStatusRequest struct {
	*tchttp.BaseRequest

	// 拉取最大条数，最多100条
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 短信SdkAppid在[短信控制台](https://console.cloud.tencent.com/sms/smslist) 添加应用后生成的实际SdkAppid,示例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`
}

func (r *PullSmsSendStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullSmsSendStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullSmsSendStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 下发状态响应集合
		PullSmsSendStatusSet []*PullSmsSendStatus `json:"PullSmsSendStatusSet,omitempty" name:"PullSmsSendStatusSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PullSmsSendStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullSmsSendStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendSmsRequest struct {
	*tchttp.BaseRequest

	// 下发手机号码，采用 e.164 标准，+[国家或地区码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。最多不要超过200个手机号。
	PhoneNumberSet []*string `json:"PhoneNumberSet,omitempty" name:"PhoneNumberSet" list`

	// 模板 ID，必须填写已审核通过的模板 ID。模板ID可登录[短信控制台](https://console.cloud.tencent.com/sms/smslist)查看。
	TemplateID *string `json:"TemplateID,omitempty" name:"TemplateID"`

	// 短信SdkAppid在[短信控制台](https://console.cloud.tencent.com/sms/smslist) 添加应用后生成的实际SdkAppid,示例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`

	// 短信签名内容，使用 UTF-8 编码，必须填写已审核通过的签名
	// 签名信息可登录[短信控制台](https://console.cloud.tencent.com/sms/smslist) 查看。
	Sign *string `json:"Sign,omitempty" name:"Sign"`

	// 模板参数，若无模板参数，则设置为空。
	TemplateParamSet []*string `json:"TemplateParamSet,omitempty" name:"TemplateParamSet" list`

	// 短信码号扩展号，默认未开通，如需开通请联系 [sms helper](https://cloud.tencent.com/document/product/382/3773)
	ExtendCode *uint64 `json:"ExtendCode,omitempty" name:"ExtendCode"`

	// 用户的 session 内容，可以携带用户侧ID等上下文信息,server 会原样返回
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// 国际/港澳台短信senderid，国内短信填空。
	// 默认未开通，如需开通请联系 [sms helper](https://cloud.tencent.com/document/product/382/3773)
	SenderId *string `json:"SenderId,omitempty" name:"SenderId"`
}

func (r *SendSmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendSmsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendSmsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 短信发送状态
		SendStatusSet []*SendStatus `json:"SendStatusSet,omitempty" name:"SendStatusSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendSmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendSmsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendStatus struct {

	// 发送流水号
	SerialNo *string `json:"SerialNo,omitempty" name:"SerialNo"`

	// 手机号码,e.164标准，+[国家或地区码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 计费条数，计费规则请查询[计费策略](https://cloud.tencent.com/document/product/382/36135)
	Fee *uint64 `json:"Fee,omitempty" name:"Fee"`

	// 用户Session内容
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// 短信请求错误码，具体含义请参考错误码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 短信请求错误码描述
	Message *string `json:"Message,omitempty" name:"Message"`
}
