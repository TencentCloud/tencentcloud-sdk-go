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

type AddSignStatus struct {

	// 签名Id。
	SignId *uint64 `json:"SignId,omitempty" name:"SignId"`

	// 签名申请Id。
	SignApplyId *uint64 `json:"SignApplyId,omitempty" name:"SignApplyId"`
}

type AddSmsSignRequest struct {
	*tchttp.BaseRequest

	// 签名名称。
	SignName *string `json:"SignName,omitempty" name:"SignName"`

	// 签名类型。其中每种类型后面标注了其可选的 DocumentType（证明类型）：
	// 0：公司（0，1，2，3）。
	// 1：APP（0，1，2，3，4） 。
	// 2：网站（0，1，2，3，5）。
	// 3：公众号或者小程序（0，1，2，3，6）。
	// 4：商标（7）。
	// 5：政府/机关事业单位/其他机构（2，3）。
	// 注：必须按照对应关系选择证明类型，否则会审核失败。
	SignType *uint64 `json:"SignType,omitempty" name:"SignType"`

	// 证明类型：
	// 0：三证合一。
	// 1：企业营业执照。
	// 2：组织机构代码证书。
	// 3：社会信用代码证书。
	// 4：应用后台管理截图（个人开发APP）。
	// 5：网站备案后台截图（个人开发网站）。
	// 6：小程序设置页面截图（个人认证小程序）。
	// 7：商标注册书。
	DocumentType *uint64 `json:"DocumentType,omitempty" name:"DocumentType"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`

	// 签名用途：
	// 0：自用。
	// 1：他用。
	UsedMethod *uint64 `json:"UsedMethod,omitempty" name:"UsedMethod"`

	// 签名对应的资质证明图片需先进行 base64 编码格式转换，将转换后的字符串去掉前缀`data:image/jpeg;base64,`再赋值给该参数。
	ProofImage *string `json:"ProofImage,omitempty" name:"ProofImage"`

	// 委托授权证明。选择 UsedMethod 为他用之后需要提交委托的授权证明。
	// 图片需先进行 base64 编码格式转换，将转换后的字符串去掉前缀`data:image/jpeg;base64,`再赋值给该参数。
	// 注：只有 UsedMethod 在选择为 1（他用）时，这个字段才会生效。
	CommissionImage *string `json:"CommissionImage,omitempty" name:"CommissionImage"`

	// 签名的申请备注。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *AddSmsSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddSmsSignRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddSmsSignResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 添加签名响应
		AddSignStatus *AddSignStatus `json:"AddSignStatus,omitempty" name:"AddSignStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSmsSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddSmsSignResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddSmsTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 模板内容。
	TemplateContent *string `json:"TemplateContent,omitempty" name:"TemplateContent"`

	// 短信类型，0表示普通短信, 1表示营销短信。
	SmsType *uint64 `json:"SmsType,omitempty" name:"SmsType"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`

	// 模板备注，例如申请原因，使用场景等。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *AddSmsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddSmsTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddSmsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 添加短信模板响应包体
		AddTemplateStatus *AddTemplateStatus `json:"AddTemplateStatus,omitempty" name:"AddTemplateStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSmsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddSmsTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddTemplateStatus struct {

	// 模板参数
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type CallbackStatusStatistics struct {

	// 短信回执量统计。
	CallbackCount *uint64 `json:"CallbackCount,omitempty" name:"CallbackCount"`

	// 短信提交成功量统计。
	RequestSuccessCount *uint64 `json:"RequestSuccessCount,omitempty" name:"RequestSuccessCount"`

	// 短信回执失败量统计。
	CallbackFailCount *uint64 `json:"CallbackFailCount,omitempty" name:"CallbackFailCount"`

	// 短信回执成功量统计。
	CallbackSuccessCount *uint64 `json:"CallbackSuccessCount,omitempty" name:"CallbackSuccessCount"`

	// 运营商内部错误统计。
	InternalErrorCount *uint64 `json:"InternalErrorCount,omitempty" name:"InternalErrorCount"`

	// 号码无效或空号统计。
	InvalidNumberCount *uint64 `json:"InvalidNumberCount,omitempty" name:"InvalidNumberCount"`

	// 停机、关机等错误统计。
	ShutdownErrorCount *uint64 `json:"ShutdownErrorCount,omitempty" name:"ShutdownErrorCount"`

	// 号码拉入黑名单统计。
	BlackListCount *uint64 `json:"BlackListCount,omitempty" name:"BlackListCount"`

	// 运营商频率限制统计。
	FrequencyLimitCount *uint64 `json:"FrequencyLimitCount,omitempty" name:"FrequencyLimitCount"`
}

type CallbackStatusStatisticsRequest struct {
	*tchttp.BaseRequest

	// 开始时间，yyyymmddhh 需要拉取的起始时间，精确到小时。
	StartDateTime *uint64 `json:"StartDateTime,omitempty" name:"StartDateTime"`

	// 结束时间，yyyymmddhh 需要拉取的截止时间，精确到小时。
	// 注：EndDataTime 必须大于 StartDateTime。
	EndDataTime *uint64 `json:"EndDataTime,omitempty" name:"EndDataTime"`

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/sms/smslist) 添加应用后生成的实际SdkAppid，示例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`

	// 最大上限。
	// 注：目前固定设置为0。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *CallbackStatusStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CallbackStatusStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CallbackStatusStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 回执数据统计响应包体。
		CallbackStatusStatistics *CallbackStatusStatistics `json:"CallbackStatusStatistics,omitempty" name:"CallbackStatusStatistics"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CallbackStatusStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CallbackStatusStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSignStatus struct {

	// 删除状态信息。
	DeleteStatus *string `json:"DeleteStatus,omitempty" name:"DeleteStatus"`

	// 删除时间，UNIX 时间戳（单位：秒）。
	DeleteTime *uint64 `json:"DeleteTime,omitempty" name:"DeleteTime"`
}

type DeleteSmsSignRequest struct {
	*tchttp.BaseRequest

	// 待删除的签名 ID。
	SignId *uint64 `json:"SignId,omitempty" name:"SignId"`
}

func (r *DeleteSmsSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSmsSignRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSmsSignResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除签名响应
		DeleteSignStatus *DeleteSignStatus `json:"DeleteSignStatus,omitempty" name:"DeleteSignStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSmsSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSmsSignResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSmsTemplateRequest struct {
	*tchttp.BaseRequest

	// 待删除的模板 ID。
	TemplateId *uint64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteSmsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSmsTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSmsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除模板响应
		DeleteTemplateStatus *DeleteTemplateStatus `json:"DeleteTemplateStatus,omitempty" name:"DeleteTemplateStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSmsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSmsTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTemplateStatus struct {

	// 删除状态信息。
	DeleteStatus *string `json:"DeleteStatus,omitempty" name:"DeleteStatus"`

	// 删除时间，UNIX 时间戳（单位：秒）。
	DeleteTime *uint64 `json:"DeleteTime,omitempty" name:"DeleteTime"`
}

type DescribeSignListStatus struct {

	// 签名Id
	SignId *uint64 `json:"SignId,omitempty" name:"SignId"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`

	// 申请签名状态。其中：
	// 0：表示审核通过。
	// -1：表示审核未通过或审核失败。
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 审核回复，审核人员审核后给出的回复，通常是审核未通过的原因。
	ReviewReply *string `json:"ReviewReply,omitempty" name:"ReviewReply"`

	// 签名名称。
	SignName *string `json:"SignName,omitempty" name:"SignName"`

	// 提交审核时间，UNIX 时间戳（单位：秒）。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeSmsSignListRequest struct {
	*tchttp.BaseRequest

	// 签名 ID 数组。
	SignIdSet []*uint64 `json:"SignIdSet,omitempty" name:"SignIdSet" list`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`
}

func (r *DescribeSmsSignListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSmsSignListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSmsSignListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 获取签名信息响应
		DescribeSignListStatusSet []*DescribeSignListStatus `json:"DescribeSignListStatusSet,omitempty" name:"DescribeSignListStatusSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSmsSignListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSmsSignListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSmsTemplateListRequest struct {
	*tchttp.BaseRequest

	// 模板 ID 数组。
	TemplateIdSet []*uint64 `json:"TemplateIdSet,omitempty" name:"TemplateIdSet" list`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`
}

func (r *DescribeSmsTemplateListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSmsTemplateListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSmsTemplateListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 获取短信签名信息响应
		DescribeTemplateStatusSet []*DescribeTemplateListStatus `json:"DescribeTemplateStatusSet,omitempty" name:"DescribeTemplateStatusSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSmsTemplateListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSmsTemplateListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateListStatus struct {

	// 模板Id
	TemplateId *uint64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`

	// 申请签名状态。其中：
	// 0：表示审核通过。
	// -1：表示审核未通过或审核失败。
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 审核回复，审核人员审核后给出的回复，通常是审核未通过的原因。
	ReviewReply *string `json:"ReviewReply,omitempty" name:"ReviewReply"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 提交审核时间，UNIX 时间戳（单位：秒）。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ModifySignStatus struct {

	// 签名Id
	SignId *uint64 `json:"SignId,omitempty" name:"SignId"`

	// 签名修改申请Id
	SignApplyId *string `json:"SignApplyId,omitempty" name:"SignApplyId"`
}

type ModifySmsSignRequest struct {
	*tchttp.BaseRequest

	// 待修改的签名 ID。
	SignId *uint64 `json:"SignId,omitempty" name:"SignId"`

	// 签名名称。
	SignName *string `json:"SignName,omitempty" name:"SignName"`

	// 签名类型。其中每种类型后面标注了其可选的 DocumentType（证明类型）：
	// 0：公司（0，1，2，3）。
	// 1：APP（0，1，2，3，4） 。
	// 2：网站（0，1，2，3，5）。
	// 3：公众号或者小程序（0，1，2，3，6）。
	// 4：商标（7）。
	// 5：政府/机关事业单位/其他机构（2，3）。
	// 注：必须按照对应关系选择证明类型，否则会审核失败。
	SignType *uint64 `json:"SignType,omitempty" name:"SignType"`

	// 证明类型：
	// 0：三证合一。
	// 1：企业营业执照。
	// 2：组织机构代码证书。
	// 3：社会信用代码证书。
	// 4：应用后台管理截图(个人开发APP)。
	// 5：网站备案后台截图(个人开发网站)。
	// 6：小程序设置页面截图(个人认证小程序)。
	// 7：商标注册书。
	DocumentType *uint64 `json:"DocumentType,omitempty" name:"DocumentType"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`

	// 签名用途：
	// 0：自用。
	// 1：他用。
	UsedMethod *uint64 `json:"UsedMethod,omitempty" name:"UsedMethod"`

	// 签名对应的资质证明图片需先进行 base64 编码格式转换，将转换后的字符串去掉前缀`data:image/jpeg;base64,`再赋值给该参数。
	ProofImage *string `json:"ProofImage,omitempty" name:"ProofImage"`

	// 委托授权证明。选择 UsedMethod 为他用之后需要提交委托的授权证明。
	// 图片需先进行 base64 编码格式转换，将转换后的字符串去掉前缀`data:image/jpeg;base64,`再赋值给该参数。
	// 注：只有 UsedMethod 在选择为 1（他用）时，这个字段才会生效。
	CommissionImage *string `json:"CommissionImage,omitempty" name:"CommissionImage"`

	// 签名的申请备注。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifySmsSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySmsSignRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySmsSignResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改签名响应
		ModifySignStatus *ModifySignStatus `json:"ModifySignStatus,omitempty" name:"ModifySignStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySmsSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySmsSignResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySmsTemplateRequest struct {
	*tchttp.BaseRequest

	// 待修改的模板的模板 ID。
	TemplateId *uint64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 新的模板名称。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 新的模板内容。
	TemplateContent *string `json:"TemplateContent,omitempty" name:"TemplateContent"`

	// 短信类型，0表示普通短信, 1表示营销短信。
	SmsType *uint64 `json:"SmsType,omitempty" name:"SmsType"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`

	// 模板备注，例如申请原因，使用场景等。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifySmsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySmsTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySmsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改模板参数响应
		ModifyTemplateStatus *ModifyTemplateStatus `json:"ModifyTemplateStatus,omitempty" name:"ModifyTemplateStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySmsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySmsTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTemplateStatus struct {

	// 模板参数
	TemplateId *uint64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

type PullSmsReplyStatus struct {

	// 短信码号扩展号，默认未开通，如需开通请联系 [sms helper](https://cloud.tencent.com/document/product/382/3773)。
	ExtendCode *string `json:"ExtendCode,omitempty" name:"ExtendCode"`

	// 国家（或地区）码。
	NationCode *string `json:"NationCode,omitempty" name:"NationCode"`

	// 手机号码,e.164标准，+[国家或地区码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 短信签名。
	Sign *string `json:"Sign,omitempty" name:"Sign"`

	// 用户回复的内容。
	ReplyContent *string `json:"ReplyContent,omitempty" name:"ReplyContent"`

	// 回复时间（例如：2019-10-08 17:18:37）。
	ReplyTime *string `json:"ReplyTime,omitempty" name:"ReplyTime"`

	// 回复时间，UNIX 时间戳（单位：秒）。
	ReplyUnixTime *uint64 `json:"ReplyUnixTime,omitempty" name:"ReplyUnixTime"`
}

type PullSmsReplyStatusByPhoneNumberRequest struct {
	*tchttp.BaseRequest

	// 拉取起始时间，UNIX 时间戳（时间：秒）。
	SendDateTime *uint64 `json:"SendDateTime,omitempty" name:"SendDateTime"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 拉取最大条数，最多 100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下发目的手机号码，依据 e.164 标准为：+[国家（或地区）码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/sms/smslist) 添加应用后生成的实际SdkAppid，例如1400006666。
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

		// 回复状态响应集合。
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

	// 拉取最大条数，最多100条。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 短信 SdkAppid 在 [短信控制台](https://console.cloud.tencent.com/sms/smslist) 添加应用后生成的实际 SdkAppid，例如1400006666。
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

		// 回复状态响应集合。
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

	// 用户实际接收到短信的时间。
	UserReceiveTime *string `json:"UserReceiveTime,omitempty" name:"UserReceiveTime"`

	// 用户实际接收到短信的时间，UNIX 时间戳（单位：秒）。
	UserReceiveUnixTime *uint64 `json:"UserReceiveUnixTime,omitempty" name:"UserReceiveUnixTime"`

	// 国家（或地区）码。
	NationCode *string `json:"NationCode,omitempty" name:"NationCode"`

	// 手机号码,e.164标准，+[国家或地区码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PurePhoneNumber *string `json:"PurePhoneNumber,omitempty" name:"PurePhoneNumber"`

	// 手机号码，普通格式，示例如：13711112222。
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 本次发送标识 ID。
	SerialNo *string `json:"SerialNo,omitempty" name:"SerialNo"`

	// 实际是否收到短信接收状态，SUCCESS（成功）、FAIL（失败）。
	ReportStatus *string `json:"ReportStatus,omitempty" name:"ReportStatus"`

	// 用户接收短信状态描述。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type PullSmsSendStatusByPhoneNumberRequest struct {
	*tchttp.BaseRequest

	// 拉取起始时间，UNIX 时间戳（时间：秒）。
	SendDateTime *uint64 `json:"SendDateTime,omitempty" name:"SendDateTime"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 拉取最大条数，最多 100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下发目的手机号码，依据 e.164 标准为：+[国家（或地区）码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/sms/smslist) 添加应用后生成的实际SdkAppid，例如1400006666。
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

		// 下发状态响应集合。
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

	// 拉取最大条数，最多100条。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/sms/smslist) 添加应用后生成的实际SdkAppid，例如1400006666。
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

		// 下发状态响应集合。
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

	// 下发手机号码，采用 e.164 标准，格式为+[国家或地区码][手机号]，单次请求最多支持200个手机号且要求全为境内手机号或全为境外手机号。
	// 例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumberSet []*string `json:"PhoneNumberSet,omitempty" name:"PhoneNumberSet" list`

	// 模板 ID，必须填写已审核通过的模板 ID。模板ID可登录 [短信控制台](https://console.cloud.tencent.com/sms/smslist) 查看。
	TemplateID *string `json:"TemplateID,omitempty" name:"TemplateID"`

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/sms/smslist)  添加应用后生成的实际SdkAppid，示例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`

	// 短信签名内容，使用 UTF-8 编码，必须填写已审核通过的签名，签名信息可登录 [短信控制台](https://console.cloud.tencent.com/sms/smslist)  查看。注：国内短信为必填参数。
	Sign *string `json:"Sign,omitempty" name:"Sign"`

	// 模板参数，若无模板参数，则设置为空。
	TemplateParamSet []*string `json:"TemplateParamSet,omitempty" name:"TemplateParamSet" list`

	// 短信码号扩展号，默认未开通，如需开通请联系 [sms helper](https://cloud.tencent.com/document/product/382/3773)。
	ExtendCode *string `json:"ExtendCode,omitempty" name:"ExtendCode"`

	// 用户的 session 内容，可以携带用户侧 ID 等上下文信息，server 会原样返回。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// 国际/港澳台短信 senderid，国内短信填空，默认未开通，如需开通请联系 [sms helper](https://cloud.tencent.com/document/product/382/3773)。
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

		// 短信发送状态。
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

	// 发送流水号。
	SerialNo *string `json:"SerialNo,omitempty" name:"SerialNo"`

	// 手机号码,e.164标准，+[国家或地区码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 计费条数，计费规则请查询 [计费策略](https://cloud.tencent.com/document/product/382/36135)。
	Fee *uint64 `json:"Fee,omitempty" name:"Fee"`

	// 用户Session内容。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// 短信请求错误码，具体含义请参考错误码。
	Code *string `json:"Code,omitempty" name:"Code"`

	// 短信请求错误码描述。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type SendStatusStatistics struct {

	// 短信计费条数统计，例如提交成功量为100条，其中有20条是长短信（长度为80字）被拆分成2条，则计费条数为： ```80 * 1 + 20 * 2 = 120``` 条。
	FeeCount *uint64 `json:"FeeCount,omitempty" name:"FeeCount"`

	// 短信提交量统计。
	RequestCount *uint64 `json:"RequestCount,omitempty" name:"RequestCount"`

	// 短信提交成功量统计。
	RequestSuccessCount *uint64 `json:"RequestSuccessCount,omitempty" name:"RequestSuccessCount"`
}

type SendStatusStatisticsRequest struct {
	*tchttp.BaseRequest

	// 拉取起始时间，yyyymmddhh 需要拉取的起始时间，精确到小时。
	StartDateTime *uint64 `json:"StartDateTime,omitempty" name:"StartDateTime"`

	// 结束时间，yyyymmddhh 需要拉取的截止时间，精确到小时
	// 注：EndDataTime 必须大于 StartDateTime。
	EndDataTime *uint64 `json:"EndDataTime,omitempty" name:"EndDataTime"`

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/sms/smslist) 添加应用后生成的实际SdkAppid，示例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`

	// 最大上限。
	// 注：目前固定设置为0。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *SendStatusStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendStatusStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendStatusStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发送数据统计响应包体。
		SendStatusStatistics *SendStatusStatistics `json:"SendStatusStatistics,omitempty" name:"SendStatusStatistics"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendStatusStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendStatusStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SmsPackagesStatistics struct {

	// 套餐包创建时间，标准时间，例如：2019-10-08 17:18:37。
	PackageCreateTime *string `json:"PackageCreateTime,omitempty" name:"PackageCreateTime"`

	// 套餐包创建时间，UNIX 时间戳（单位：秒）。
	PackageCreateUnixTime *uint64 `json:"PackageCreateUnixTime,omitempty" name:"PackageCreateUnixTime"`

	// 套餐包生效时间，标准时间，例如：2019-10-08 17:18:37。
	PackageEffectiveTime *string `json:"PackageEffectiveTime,omitempty" name:"PackageEffectiveTime"`

	// 套餐包生效时间，UNIX 时间戳（单位：秒）。
	PackageEffectiveUnixTime *uint64 `json:"PackageEffectiveUnixTime,omitempty" name:"PackageEffectiveUnixTime"`

	// 套餐包过期时间，标准时间，例如：2019-10-08 17:18:37。
	PackageExpiredTime *string `json:"PackageExpiredTime,omitempty" name:"PackageExpiredTime"`

	// 套餐包过期时间，UNIX 时间戳（单位：秒）。
	PackageExpiredUnixTime *uint64 `json:"PackageExpiredUnixTime,omitempty" name:"PackageExpiredUnixTime"`

	// 套餐包条数。
	AmountOfPackage *uint64 `json:"AmountOfPackage,omitempty" name:"AmountOfPackage"`

	// 0表示赠送套餐包，1表示购买套餐包。
	TypeOfPackage *uint64 `json:"TypeOfPackage,omitempty" name:"TypeOfPackage"`

	// 套餐包 ID。
	PackageId *uint64 `json:"PackageId,omitempty" name:"PackageId"`

	// 当前使用量。
	CurrentUsage *uint64 `json:"CurrentUsage,omitempty" name:"CurrentUsage"`
}

type SmsPackagesStatisticsRequest struct {
	*tchttp.BaseRequest

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/sms/smslist) 添加应用后生成的实际SdkAppid，示例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`

	// 最大上限(需要拉取的套餐包个数)。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *SmsPackagesStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SmsPackagesStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SmsPackagesStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发送数据统计响应包体。
		SmsPackagesStatisticsSet []*SmsPackagesStatistics `json:"SmsPackagesStatisticsSet,omitempty" name:"SmsPackagesStatisticsSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SmsPackagesStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SmsPackagesStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
