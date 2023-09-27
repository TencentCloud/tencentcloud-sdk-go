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

package v20210111

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AddSignStatus struct {
	// 签名ID。
	SignId *uint64 `json:"SignId,omitnil" name:"SignId"`
}

// Predefined struct for user
type AddSmsSignRequestParams struct {
	// 签名名称。
	// 注：不能重复申请已通过或待审核的签名。
	SignName *string `json:"SignName,omitnil" name:"SignName"`

	// 签名类型。其中每种类型后面标注了其可选的 DocumentType（证明类型）：
	// 0：公司，可选 DocumentType 有（0，1）。
	// 1：APP，可选 DocumentType 有（0，1，2，3，4） 。
	// 2：网站，可选 DocumentType 有（0，1，2，3，5）。
	// 3：公众号，可选 DocumentType 有（0，1，2，3，8）。
	// 4：商标，可选 DocumentType 有（7）。
	// 5：政府/机关事业单位/其他机构，可选 DocumentType 有（2，3）。
	// 6：小程序，可选 DocumentType 有（0，1，2，3，6）。
	// 注：必须按照对应关系选择证明类型，否则会审核失败。
	SignType *uint64 `json:"SignType,omitnil" name:"SignType"`

	// 证明类型：
	// 0：三证合一。
	// 1：企业营业执照。
	// 2：组织机构代码证书。
	// 3：社会信用代码证书。
	// 4：应用后台管理截图（个人开发APP）。
	// 5：网站备案后台截图（个人开发网站）。
	// 6：小程序设置页面截图（个人认证小程序）。
	// 7：商标注册书。
	// 8：公众号设置页面截图（个人认证公众号）。
	DocumentType *uint64 `json:"DocumentType,omitnil" name:"DocumentType"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitnil" name:"International"`

	// 签名用途：
	// 0：自用。
	// 1：他用。
	SignPurpose *uint64 `json:"SignPurpose,omitnil" name:"SignPurpose"`

	// 签名对应的资质证明图片需先进行 base64 编码格式转换，将转换后的字符串去掉前缀`data:image/jpeg;base64,`再赋值给该参数。
	ProofImage *string `json:"ProofImage,omitnil" name:"ProofImage"`

	// 委托授权证明。选择 SignPurpose 为他用之后需要提交委托的授权证明。
	// 图片需先进行 base64 编码格式转换，将转换后的字符串去掉前缀`data:image/jpeg;base64,`再赋值给该参数。
	// 注：只有 SignPurpose 在选择为 1（他用）时，这个字段才会生效。
	CommissionImage *string `json:"CommissionImage,omitnil" name:"CommissionImage"`

	// 签名的申请备注。
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type AddSmsSignRequest struct {
	*tchttp.BaseRequest
	
	// 签名名称。
	// 注：不能重复申请已通过或待审核的签名。
	SignName *string `json:"SignName,omitnil" name:"SignName"`

	// 签名类型。其中每种类型后面标注了其可选的 DocumentType（证明类型）：
	// 0：公司，可选 DocumentType 有（0，1）。
	// 1：APP，可选 DocumentType 有（0，1，2，3，4） 。
	// 2：网站，可选 DocumentType 有（0，1，2，3，5）。
	// 3：公众号，可选 DocumentType 有（0，1，2，3，8）。
	// 4：商标，可选 DocumentType 有（7）。
	// 5：政府/机关事业单位/其他机构，可选 DocumentType 有（2，3）。
	// 6：小程序，可选 DocumentType 有（0，1，2，3，6）。
	// 注：必须按照对应关系选择证明类型，否则会审核失败。
	SignType *uint64 `json:"SignType,omitnil" name:"SignType"`

	// 证明类型：
	// 0：三证合一。
	// 1：企业营业执照。
	// 2：组织机构代码证书。
	// 3：社会信用代码证书。
	// 4：应用后台管理截图（个人开发APP）。
	// 5：网站备案后台截图（个人开发网站）。
	// 6：小程序设置页面截图（个人认证小程序）。
	// 7：商标注册书。
	// 8：公众号设置页面截图（个人认证公众号）。
	DocumentType *uint64 `json:"DocumentType,omitnil" name:"DocumentType"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitnil" name:"International"`

	// 签名用途：
	// 0：自用。
	// 1：他用。
	SignPurpose *uint64 `json:"SignPurpose,omitnil" name:"SignPurpose"`

	// 签名对应的资质证明图片需先进行 base64 编码格式转换，将转换后的字符串去掉前缀`data:image/jpeg;base64,`再赋值给该参数。
	ProofImage *string `json:"ProofImage,omitnil" name:"ProofImage"`

	// 委托授权证明。选择 SignPurpose 为他用之后需要提交委托的授权证明。
	// 图片需先进行 base64 编码格式转换，将转换后的字符串去掉前缀`data:image/jpeg;base64,`再赋值给该参数。
	// 注：只有 SignPurpose 在选择为 1（他用）时，这个字段才会生效。
	CommissionImage *string `json:"CommissionImage,omitnil" name:"CommissionImage"`

	// 签名的申请备注。
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

func (r *AddSmsSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSmsSignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SignName")
	delete(f, "SignType")
	delete(f, "DocumentType")
	delete(f, "International")
	delete(f, "SignPurpose")
	delete(f, "ProofImage")
	delete(f, "CommissionImage")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddSmsSignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddSmsSignResponseParams struct {
	// 添加签名响应
	AddSignStatus *AddSignStatus `json:"AddSignStatus,omitnil" name:"AddSignStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddSmsSignResponse struct {
	*tchttp.BaseResponse
	Response *AddSmsSignResponseParams `json:"Response"`
}

func (r *AddSmsSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSmsSignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddSmsTemplateRequestParams struct {
	// 模板名称。
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 模板内容。
	TemplateContent *string `json:"TemplateContent,omitnil" name:"TemplateContent"`

	// 短信类型，0表示普通短信, 1表示营销短信。
	SmsType *uint64 `json:"SmsType,omitnil" name:"SmsType"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitnil" name:"International"`

	// 模板备注，例如申请原因，使用场景等。
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type AddSmsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称。
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 模板内容。
	TemplateContent *string `json:"TemplateContent,omitnil" name:"TemplateContent"`

	// 短信类型，0表示普通短信, 1表示营销短信。
	SmsType *uint64 `json:"SmsType,omitnil" name:"SmsType"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitnil" name:"International"`

	// 模板备注，例如申请原因，使用场景等。
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

func (r *AddSmsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSmsTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "TemplateContent")
	delete(f, "SmsType")
	delete(f, "International")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddSmsTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddSmsTemplateResponseParams struct {
	// 添加短信模板响应包体
	AddTemplateStatus *AddTemplateStatus `json:"AddTemplateStatus,omitnil" name:"AddTemplateStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddSmsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *AddSmsTemplateResponseParams `json:"Response"`
}

func (r *AddSmsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSmsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddTemplateStatus struct {
	// 模板ID。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

type CallbackStatusStatistics struct {
	// 短信回执量统计。
	CallbackCount *uint64 `json:"CallbackCount,omitnil" name:"CallbackCount"`

	// 短信提交成功量统计。
	RequestSuccessCount *uint64 `json:"RequestSuccessCount,omitnil" name:"RequestSuccessCount"`

	// 短信回执失败量统计。
	CallbackFailCount *uint64 `json:"CallbackFailCount,omitnil" name:"CallbackFailCount"`

	// 短信回执成功量统计。
	CallbackSuccessCount *uint64 `json:"CallbackSuccessCount,omitnil" name:"CallbackSuccessCount"`

	// 运营商内部错误统计。
	InternalErrorCount *uint64 `json:"InternalErrorCount,omitnil" name:"InternalErrorCount"`

	// 号码无效或空号统计。
	InvalidNumberCount *uint64 `json:"InvalidNumberCount,omitnil" name:"InvalidNumberCount"`

	// 停机、关机等错误统计。
	ShutdownErrorCount *uint64 `json:"ShutdownErrorCount,omitnil" name:"ShutdownErrorCount"`

	// 号码拉入黑名单统计。
	BlackListCount *uint64 `json:"BlackListCount,omitnil" name:"BlackListCount"`

	// 运营商频率限制统计。
	FrequencyLimitCount *uint64 `json:"FrequencyLimitCount,omitnil" name:"FrequencyLimitCount"`
}

// Predefined struct for user
type CallbackStatusStatisticsRequestParams struct {
	// 起始时间，格式为yyyymmddhh，精确到小时，例如2021050113，表示2021年5月1号13时。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间，格式为yyyymmddhh，精确到小时，例如2021050118，表示2021年5月1号18时。
	// 注：EndTime 必须大于 BeginTime，且相差不超过32天。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 短信 SdkAppId 在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage)  添加应用后生成的实际 SdkAppId，示例如1400006666。
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil" name:"SmsSdkAppId"`

	// 最大上限。
	// 注：目前固定设置为0。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type CallbackStatusStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间，格式为yyyymmddhh，精确到小时，例如2021050113，表示2021年5月1号13时。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间，格式为yyyymmddhh，精确到小时，例如2021050118，表示2021年5月1号18时。
	// 注：EndTime 必须大于 BeginTime，且相差不超过32天。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 短信 SdkAppId 在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage)  添加应用后生成的实际 SdkAppId，示例如1400006666。
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil" name:"SmsSdkAppId"`

	// 最大上限。
	// 注：目前固定设置为0。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *CallbackStatusStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallbackStatusStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "SmsSdkAppId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CallbackStatusStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CallbackStatusStatisticsResponseParams struct {
	// 回执数据统计响应包体。
	CallbackStatusStatistics *CallbackStatusStatistics `json:"CallbackStatusStatistics,omitnil" name:"CallbackStatusStatistics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CallbackStatusStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *CallbackStatusStatisticsResponseParams `json:"Response"`
}

func (r *CallbackStatusStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallbackStatusStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSignStatus struct {
	// 删除状态信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteStatus *string `json:"DeleteStatus,omitnil" name:"DeleteStatus"`

	// 删除时间，UNIX 时间戳（单位：秒）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteTime *uint64 `json:"DeleteTime,omitnil" name:"DeleteTime"`
}

// Predefined struct for user
type DeleteSmsSignRequestParams struct {
	// 待删除的签名 ID。
	SignId *uint64 `json:"SignId,omitnil" name:"SignId"`
}

type DeleteSmsSignRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的签名 ID。
	SignId *uint64 `json:"SignId,omitnil" name:"SignId"`
}

func (r *DeleteSmsSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSmsSignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SignId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSmsSignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSmsSignResponseParams struct {
	// 删除签名响应
	DeleteSignStatus *DeleteSignStatus `json:"DeleteSignStatus,omitnil" name:"DeleteSignStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteSmsSignResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSmsSignResponseParams `json:"Response"`
}

func (r *DeleteSmsSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSmsSignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSmsTemplateRequestParams struct {
	// 待删除的模板 ID。
	TemplateId *uint64 `json:"TemplateId,omitnil" name:"TemplateId"`
}

type DeleteSmsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的模板 ID。
	TemplateId *uint64 `json:"TemplateId,omitnil" name:"TemplateId"`
}

func (r *DeleteSmsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSmsTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSmsTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSmsTemplateResponseParams struct {
	// 删除模板响应
	DeleteTemplateStatus *DeleteTemplateStatus `json:"DeleteTemplateStatus,omitnil" name:"DeleteTemplateStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteSmsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSmsTemplateResponseParams `json:"Response"`
}

func (r *DeleteSmsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSmsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTemplateStatus struct {
	// 删除状态信息。
	DeleteStatus *string `json:"DeleteStatus,omitnil" name:"DeleteStatus"`

	// 删除时间，UNIX 时间戳（单位：秒）。
	DeleteTime *uint64 `json:"DeleteTime,omitnil" name:"DeleteTime"`
}

// Predefined struct for user
type DescribePhoneNumberInfoRequestParams struct {
	// 查询手机号码，采用 E.164 标准，格式为+[国家或地区码][手机号]，单次请求最多支持200个手机号。
	// 例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumberSet []*string `json:"PhoneNumberSet,omitnil" name:"PhoneNumberSet"`
}

type DescribePhoneNumberInfoRequest struct {
	*tchttp.BaseRequest
	
	// 查询手机号码，采用 E.164 标准，格式为+[国家或地区码][手机号]，单次请求最多支持200个手机号。
	// 例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumberSet []*string `json:"PhoneNumberSet,omitnil" name:"PhoneNumberSet"`
}

func (r *DescribePhoneNumberInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePhoneNumberInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PhoneNumberSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePhoneNumberInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePhoneNumberInfoResponseParams struct {
	// 获取号码信息。
	PhoneNumberInfoSet []*PhoneNumberInfo `json:"PhoneNumberInfoSet,omitnil" name:"PhoneNumberInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePhoneNumberInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribePhoneNumberInfoResponseParams `json:"Response"`
}

func (r *DescribePhoneNumberInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePhoneNumberInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSignListStatus struct {
	// 签名ID。
	SignId *uint64 `json:"SignId,omitnil" name:"SignId"`

	// 是否国际/港澳台短信，其中0表示国内短信，1表示国际/港澳台短信。
	International *uint64 `json:"International,omitnil" name:"International"`

	// 申请签名状态，其中0表示审核通过，1表示审核中。
	// -1：表示审核未通过或审核失败。
	StatusCode *int64 `json:"StatusCode,omitnil" name:"StatusCode"`

	// 审核回复，审核人员审核后给出的回复，通常是审核未通过的原因。
	ReviewReply *string `json:"ReviewReply,omitnil" name:"ReviewReply"`

	// 签名名称。
	SignName *string `json:"SignName,omitnil" name:"SignName"`

	// 提交审核时间，UNIX 时间戳（单位：秒）。
	CreateTime *uint64 `json:"CreateTime,omitnil" name:"CreateTime"`
}

// Predefined struct for user
type DescribeSmsSignListRequestParams struct {
	// 签名 ID 数组。
	// 注：默认数组最大长度100。
	SignIdSet []*uint64 `json:"SignIdSet,omitnil" name:"SignIdSet"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitnil" name:"International"`
}

type DescribeSmsSignListRequest struct {
	*tchttp.BaseRequest
	
	// 签名 ID 数组。
	// 注：默认数组最大长度100。
	SignIdSet []*uint64 `json:"SignIdSet,omitnil" name:"SignIdSet"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitnil" name:"International"`
}

func (r *DescribeSmsSignListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsSignListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SignIdSet")
	delete(f, "International")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSmsSignListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmsSignListResponseParams struct {
	// 获取签名信息响应
	DescribeSignListStatusSet []*DescribeSignListStatus `json:"DescribeSignListStatusSet,omitnil" name:"DescribeSignListStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSmsSignListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSmsSignListResponseParams `json:"Response"`
}

func (r *DescribeSmsSignListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsSignListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmsTemplateListRequestParams struct {
	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitnil" name:"International"`

	// 模板 ID 数组。数组为空时默认查询模板列表信息，请使用 Limit 和 Offset 字段设置查询范围。
	// <dx-alert infotype="notice" title="注意">默认数组长度最大100</dx-alert>
	TemplateIdSet []*uint64 `json:"TemplateIdSet,omitnil" name:"TemplateIdSet"`

	// 最大上限，最多100。
	// 注：默认为0，TemplateIdSet 为空时启用。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量。
	// 注：默认为0，TemplateIdSet 为空时启用。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeSmsTemplateListRequest struct {
	*tchttp.BaseRequest
	
	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitnil" name:"International"`

	// 模板 ID 数组。数组为空时默认查询模板列表信息，请使用 Limit 和 Offset 字段设置查询范围。
	// <dx-alert infotype="notice" title="注意">默认数组长度最大100</dx-alert>
	TemplateIdSet []*uint64 `json:"TemplateIdSet,omitnil" name:"TemplateIdSet"`

	// 最大上限，最多100。
	// 注：默认为0，TemplateIdSet 为空时启用。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量。
	// 注：默认为0，TemplateIdSet 为空时启用。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeSmsTemplateListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsTemplateListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "International")
	delete(f, "TemplateIdSet")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSmsTemplateListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmsTemplateListResponseParams struct {
	// 获取短信模板信息响应
	DescribeTemplateStatusSet []*DescribeTemplateListStatus `json:"DescribeTemplateStatusSet,omitnil" name:"DescribeTemplateStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSmsTemplateListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSmsTemplateListResponseParams `json:"Response"`
}

func (r *DescribeSmsTemplateListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsTemplateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateListStatus struct {
	// 模板ID。
	TemplateId *uint64 `json:"TemplateId,omitnil" name:"TemplateId"`

	// 是否国际/港澳台短信，其中0表示国内短信，1表示国际/港澳台短信，3表示该模板既支持国内短信也支持国际/港澳台短信。
	International *uint64 `json:"International,omitnil" name:"International"`

	// 申请模板状态，其中0表示审核通过且已生效，1表示审核中，2表示审核通过待生效，-1表示审核未通过或审核失败。注：只有状态值为0时该模板才能使用。
	StatusCode *int64 `json:"StatusCode,omitnil" name:"StatusCode"`

	// 审核回复，审核人员审核后给出的回复，通常是审核未通过的原因。
	ReviewReply *string `json:"ReviewReply,omitnil" name:"ReviewReply"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 提交审核时间，UNIX 时间戳（单位：秒）。
	CreateTime *uint64 `json:"CreateTime,omitnil" name:"CreateTime"`

	// 模板内容。
	TemplateContent *string `json:"TemplateContent,omitnil" name:"TemplateContent"`
}

type ModifySignStatus struct {
	// 签名ID。
	SignId *uint64 `json:"SignId,omitnil" name:"SignId"`
}

// Predefined struct for user
type ModifySmsSignRequestParams struct {
	// 待修改的签名 ID。
	SignId *uint64 `json:"SignId,omitnil" name:"SignId"`

	// 签名名称。
	SignName *string `json:"SignName,omitnil" name:"SignName"`

	// 签名类型。其中每种类型后面标注了其可选的 DocumentType（证明类型）：
	// 0：公司，可选 DocumentType 有（0，1）。
	// 1：APP，可选 DocumentType 有（0，1，2，3，4） 。
	// 2：网站，可选 DocumentType 有（0，1，2，3，5）。
	// 3：公众号，可选 DocumentType 有（0，1，2，3，8）。
	// 4：商标，可选 DocumentType 有（7）。
	// 5：政府/机关事业单位/其他机构，可选 DocumentType 有（2，3）。
	// 6：小程序，可选 DocumentType 有（0，1，2，3，6）。
	// 注：必须按照对应关系选择证明类型，否则会审核失败。
	SignType *uint64 `json:"SignType,omitnil" name:"SignType"`

	// 证明类型：
	// 0：三证合一。
	// 1：企业营业执照。
	// 2：组织机构代码证书。
	// 3：社会信用代码证书。
	// 4：应用后台管理截图（个人开发APP）。
	// 5：网站备案后台截图（个人开发网站）。
	// 6：小程序设置页面截图（个人认证小程序）。
	// 7：商标注册书。
	// 8：公众号设置页面截图（个人认证公众号）。
	DocumentType *uint64 `json:"DocumentType,omitnil" name:"DocumentType"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	// 注：需要和待修改签名International值保持一致，该参数不能直接修改国内签名到国际签名。
	International *uint64 `json:"International,omitnil" name:"International"`

	// 签名用途：
	// 0：自用。
	// 1：他用。
	SignPurpose *uint64 `json:"SignPurpose,omitnil" name:"SignPurpose"`

	// 签名对应的资质证明图片需先进行 base64 编码格式转换，将转换后的字符串去掉前缀`data:image/jpeg;base64,`再赋值给该参数。
	ProofImage *string `json:"ProofImage,omitnil" name:"ProofImage"`

	// 委托授权证明。选择 SignPurpose 为他用之后需要提交委托的授权证明。
	// 图片需先进行 base64 编码格式转换，将转换后的字符串去掉前缀`data:image/jpeg;base64,`再赋值给该参数。
	// 注：只有 SignPurpose 在选择为 1（他用）时，这个字段才会生效。
	CommissionImage *string `json:"CommissionImage,omitnil" name:"CommissionImage"`

	// 签名的申请备注。
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type ModifySmsSignRequest struct {
	*tchttp.BaseRequest
	
	// 待修改的签名 ID。
	SignId *uint64 `json:"SignId,omitnil" name:"SignId"`

	// 签名名称。
	SignName *string `json:"SignName,omitnil" name:"SignName"`

	// 签名类型。其中每种类型后面标注了其可选的 DocumentType（证明类型）：
	// 0：公司，可选 DocumentType 有（0，1）。
	// 1：APP，可选 DocumentType 有（0，1，2，3，4） 。
	// 2：网站，可选 DocumentType 有（0，1，2，3，5）。
	// 3：公众号，可选 DocumentType 有（0，1，2，3，8）。
	// 4：商标，可选 DocumentType 有（7）。
	// 5：政府/机关事业单位/其他机构，可选 DocumentType 有（2，3）。
	// 6：小程序，可选 DocumentType 有（0，1，2，3，6）。
	// 注：必须按照对应关系选择证明类型，否则会审核失败。
	SignType *uint64 `json:"SignType,omitnil" name:"SignType"`

	// 证明类型：
	// 0：三证合一。
	// 1：企业营业执照。
	// 2：组织机构代码证书。
	// 3：社会信用代码证书。
	// 4：应用后台管理截图（个人开发APP）。
	// 5：网站备案后台截图（个人开发网站）。
	// 6：小程序设置页面截图（个人认证小程序）。
	// 7：商标注册书。
	// 8：公众号设置页面截图（个人认证公众号）。
	DocumentType *uint64 `json:"DocumentType,omitnil" name:"DocumentType"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	// 注：需要和待修改签名International值保持一致，该参数不能直接修改国内签名到国际签名。
	International *uint64 `json:"International,omitnil" name:"International"`

	// 签名用途：
	// 0：自用。
	// 1：他用。
	SignPurpose *uint64 `json:"SignPurpose,omitnil" name:"SignPurpose"`

	// 签名对应的资质证明图片需先进行 base64 编码格式转换，将转换后的字符串去掉前缀`data:image/jpeg;base64,`再赋值给该参数。
	ProofImage *string `json:"ProofImage,omitnil" name:"ProofImage"`

	// 委托授权证明。选择 SignPurpose 为他用之后需要提交委托的授权证明。
	// 图片需先进行 base64 编码格式转换，将转换后的字符串去掉前缀`data:image/jpeg;base64,`再赋值给该参数。
	// 注：只有 SignPurpose 在选择为 1（他用）时，这个字段才会生效。
	CommissionImage *string `json:"CommissionImage,omitnil" name:"CommissionImage"`

	// 签名的申请备注。
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

func (r *ModifySmsSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySmsSignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SignId")
	delete(f, "SignName")
	delete(f, "SignType")
	delete(f, "DocumentType")
	delete(f, "International")
	delete(f, "SignPurpose")
	delete(f, "ProofImage")
	delete(f, "CommissionImage")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySmsSignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySmsSignResponseParams struct {
	// 修改签名响应
	ModifySignStatus *ModifySignStatus `json:"ModifySignStatus,omitnil" name:"ModifySignStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySmsSignResponse struct {
	*tchttp.BaseResponse
	Response *ModifySmsSignResponseParams `json:"Response"`
}

func (r *ModifySmsSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySmsSignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySmsTemplateRequestParams struct {
	// 待修改模板的ID。
	TemplateId *uint64 `json:"TemplateId,omitnil" name:"TemplateId"`

	// 新的模板名称。
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 新的模板内容。
	TemplateContent *string `json:"TemplateContent,omitnil" name:"TemplateContent"`

	// 短信类型，0表示普通短信, 1表示营销短信。
	SmsType *uint64 `json:"SmsType,omitnil" name:"SmsType"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitnil" name:"International"`

	// 模板备注，例如申请原因，使用场景等。
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type ModifySmsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 待修改模板的ID。
	TemplateId *uint64 `json:"TemplateId,omitnil" name:"TemplateId"`

	// 新的模板名称。
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 新的模板内容。
	TemplateContent *string `json:"TemplateContent,omitnil" name:"TemplateContent"`

	// 短信类型，0表示普通短信, 1表示营销短信。
	SmsType *uint64 `json:"SmsType,omitnil" name:"SmsType"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitnil" name:"International"`

	// 模板备注，例如申请原因，使用场景等。
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

func (r *ModifySmsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySmsTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TemplateName")
	delete(f, "TemplateContent")
	delete(f, "SmsType")
	delete(f, "International")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySmsTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySmsTemplateResponseParams struct {
	// 修改模板参数响应
	ModifyTemplateStatus *ModifyTemplateStatus `json:"ModifyTemplateStatus,omitnil" name:"ModifyTemplateStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySmsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifySmsTemplateResponseParams `json:"Response"`
}

func (r *ModifySmsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySmsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTemplateStatus struct {
	// 模板ID。
	TemplateId *uint64 `json:"TemplateId,omitnil" name:"TemplateId"`
}

type PhoneNumberInfo struct {
	// 号码信息查询错误码，查询成功返回 "Ok"。
	Code *string `json:"Code,omitnil" name:"Code"`

	// 号码信息查询错误码描述。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 国家（或地区）码。
	NationCode *string `json:"NationCode,omitnil" name:"NationCode"`

	// 用户号码，去除国家或地区码前缀的普通格式，示例如：13711112222。
	SubscriberNumber *string `json:"SubscriberNumber,omitnil" name:"SubscriberNumber"`

	// 解析后的规范的 E.164 号码，与下发短信的号码解析结果一致。解析失败时会原样返回。
	PhoneNumber *string `json:"PhoneNumber,omitnil" name:"PhoneNumber"`

	// 国家码或地区码，例如 CN、US 等，对于未识别出国家码或者地区码，默认返回 DEF。
	IsoCode *string `json:"IsoCode,omitnil" name:"IsoCode"`

	// 国家码或地区名，例如 China，可参考 [国际/港澳台短信价格总览](https://cloud.tencent.com/document/product/382/18051#.E6.97.A5.E7.BB.93.E5.90.8E.E4.BB.98.E8.B4.B9.3Ca-id.3D.22post-payment.22.3E.3C.2Fa.3E)
	IsoName *string `json:"IsoName,omitnil" name:"IsoName"`
}

type PullSmsReplyStatus struct {
	// 短信码号扩展号，默认未开通，如需开通请联系 [sms helper](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	ExtendCode *string `json:"ExtendCode,omitnil" name:"ExtendCode"`

	// 国家（或地区）码。
	CountryCode *string `json:"CountryCode,omitnil" name:"CountryCode"`

	// 手机号码，E.164标准，+[国家或地区码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumber *string `json:"PhoneNumber,omitnil" name:"PhoneNumber"`

	// 短信签名名称。
	SignName *string `json:"SignName,omitnil" name:"SignName"`

	// 用户回复的内容。
	ReplyContent *string `json:"ReplyContent,omitnil" name:"ReplyContent"`

	// 回复时间，UNIX 时间戳（单位：秒）。
	ReplyTime *uint64 `json:"ReplyTime,omitnil" name:"ReplyTime"`

	// 用户号码，普通格式，示例如：13711112222。
	SubscriberNumber *string `json:"SubscriberNumber,omitnil" name:"SubscriberNumber"`
}

// Predefined struct for user
type PullSmsReplyStatusByPhoneNumberRequestParams struct {
	// 拉取起始时间，UNIX 时间戳（时间：秒）。
	// 注：最大可拉取当前时期前7天的数据。
	BeginTime *uint64 `json:"BeginTime,omitnil" name:"BeginTime"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 拉取最大条数，最多 100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 下发目的手机号码，依据 E.164 标准为：+[国家（或地区）码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumber *string `json:"PhoneNumber,omitnil" name:"PhoneNumber"`

	// 短信 SdkAppId 在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage)  添加应用后生成的实际 SdkAppId，示例如1400006666。
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil" name:"SmsSdkAppId"`

	// 拉取截止时间，UNIX 时间戳（时间：秒）。
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`
}

type PullSmsReplyStatusByPhoneNumberRequest struct {
	*tchttp.BaseRequest
	
	// 拉取起始时间，UNIX 时间戳（时间：秒）。
	// 注：最大可拉取当前时期前7天的数据。
	BeginTime *uint64 `json:"BeginTime,omitnil" name:"BeginTime"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 拉取最大条数，最多 100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 下发目的手机号码，依据 E.164 标准为：+[国家（或地区）码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumber *string `json:"PhoneNumber,omitnil" name:"PhoneNumber"`

	// 短信 SdkAppId 在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage)  添加应用后生成的实际 SdkAppId，示例如1400006666。
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil" name:"SmsSdkAppId"`

	// 拉取截止时间，UNIX 时间戳（时间：秒）。
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *PullSmsReplyStatusByPhoneNumberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullSmsReplyStatusByPhoneNumberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PhoneNumber")
	delete(f, "SmsSdkAppId")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PullSmsReplyStatusByPhoneNumberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullSmsReplyStatusByPhoneNumberResponseParams struct {
	// 回复状态响应集合。
	PullSmsReplyStatusSet []*PullSmsReplyStatus `json:"PullSmsReplyStatusSet,omitnil" name:"PullSmsReplyStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PullSmsReplyStatusByPhoneNumberResponse struct {
	*tchttp.BaseResponse
	Response *PullSmsReplyStatusByPhoneNumberResponseParams `json:"Response"`
}

func (r *PullSmsReplyStatusByPhoneNumberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullSmsReplyStatusByPhoneNumberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullSmsReplyStatusRequestParams struct {
	// 拉取最大条数，最多100条。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 短信 SdkAppId 在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage) 添加应用后生成的实际 SdkAppId，例如1400006666。
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil" name:"SmsSdkAppId"`
}

type PullSmsReplyStatusRequest struct {
	*tchttp.BaseRequest
	
	// 拉取最大条数，最多100条。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 短信 SdkAppId 在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage) 添加应用后生成的实际 SdkAppId，例如1400006666。
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil" name:"SmsSdkAppId"`
}

func (r *PullSmsReplyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullSmsReplyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "SmsSdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PullSmsReplyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullSmsReplyStatusResponseParams struct {
	// 回复状态响应集合。
	PullSmsReplyStatusSet []*PullSmsReplyStatus `json:"PullSmsReplyStatusSet,omitnil" name:"PullSmsReplyStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PullSmsReplyStatusResponse struct {
	*tchttp.BaseResponse
	Response *PullSmsReplyStatusResponseParams `json:"Response"`
}

func (r *PullSmsReplyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullSmsReplyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PullSmsSendStatus struct {
	// 用户实际接收到短信的时间，UNIX 时间戳（单位：秒）。
	UserReceiveTime *uint64 `json:"UserReceiveTime,omitnil" name:"UserReceiveTime"`

	// 国家（或地区）码。
	CountryCode *string `json:"CountryCode,omitnil" name:"CountryCode"`

	// 用户号码，普通格式，示例如：13711112222。
	SubscriberNumber *string `json:"SubscriberNumber,omitnil" name:"SubscriberNumber"`

	// 手机号码，E.164标准，+[国家或地区码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumber *string `json:"PhoneNumber,omitnil" name:"PhoneNumber"`

	// 本次发送标识 ID。
	SerialNo *string `json:"SerialNo,omitnil" name:"SerialNo"`

	// 实际是否收到短信接收状态，SUCCESS（成功）、FAIL（失败）。
	ReportStatus *string `json:"ReportStatus,omitnil" name:"ReportStatus"`

	// 用户接收短信状态描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 用户的 session 内容。与请求中的 SessionContext 一致，默认为空，如需开通请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionContext *string `json:"SessionContext,omitnil" name:"SessionContext"`
}

// Predefined struct for user
type PullSmsSendStatusByPhoneNumberRequestParams struct {
	// 拉取起始时间，UNIX 时间戳（时间：秒）。
	// 注：最大可拉取当前时期前7天的数据。
	BeginTime *uint64 `json:"BeginTime,omitnil" name:"BeginTime"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 拉取最大条数，最多 100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 下发目的手机号码，依据 E.164 标准为：+[国家（或地区）码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumber *string `json:"PhoneNumber,omitnil" name:"PhoneNumber"`

	// 短信 SdkAppId 在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage)  添加应用后生成的实际 SdkAppId，示例如1400006666。
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil" name:"SmsSdkAppId"`

	// 拉取截止时间，UNIX 时间戳（时间：秒）。
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`
}

type PullSmsSendStatusByPhoneNumberRequest struct {
	*tchttp.BaseRequest
	
	// 拉取起始时间，UNIX 时间戳（时间：秒）。
	// 注：最大可拉取当前时期前7天的数据。
	BeginTime *uint64 `json:"BeginTime,omitnil" name:"BeginTime"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 拉取最大条数，最多 100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 下发目的手机号码，依据 E.164 标准为：+[国家（或地区）码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumber *string `json:"PhoneNumber,omitnil" name:"PhoneNumber"`

	// 短信 SdkAppId 在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage)  添加应用后生成的实际 SdkAppId，示例如1400006666。
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil" name:"SmsSdkAppId"`

	// 拉取截止时间，UNIX 时间戳（时间：秒）。
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *PullSmsSendStatusByPhoneNumberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullSmsSendStatusByPhoneNumberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PhoneNumber")
	delete(f, "SmsSdkAppId")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PullSmsSendStatusByPhoneNumberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullSmsSendStatusByPhoneNumberResponseParams struct {
	// 下发状态响应集合。
	PullSmsSendStatusSet []*PullSmsSendStatus `json:"PullSmsSendStatusSet,omitnil" name:"PullSmsSendStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PullSmsSendStatusByPhoneNumberResponse struct {
	*tchttp.BaseResponse
	Response *PullSmsSendStatusByPhoneNumberResponseParams `json:"Response"`
}

func (r *PullSmsSendStatusByPhoneNumberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullSmsSendStatusByPhoneNumberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullSmsSendStatusRequestParams struct {
	// 拉取最大条数，最多100条。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 短信 SdkAppId 在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage) 添加应用后生成的实际 SdkAppId，例如1400006666。
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil" name:"SmsSdkAppId"`
}

type PullSmsSendStatusRequest struct {
	*tchttp.BaseRequest
	
	// 拉取最大条数，最多100条。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 短信 SdkAppId 在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage) 添加应用后生成的实际 SdkAppId，例如1400006666。
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil" name:"SmsSdkAppId"`
}

func (r *PullSmsSendStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullSmsSendStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "SmsSdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PullSmsSendStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullSmsSendStatusResponseParams struct {
	// 下发状态响应集合。
	PullSmsSendStatusSet []*PullSmsSendStatus `json:"PullSmsSendStatusSet,omitnil" name:"PullSmsSendStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PullSmsSendStatusResponse struct {
	*tchttp.BaseResponse
	Response *PullSmsSendStatusResponseParams `json:"Response"`
}

func (r *PullSmsSendStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullSmsSendStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportConversionRequestParams struct {
	// 短信应用ID。在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage)  添加应用后生成的实际 SdkAppId，示例如1400006666。
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil" name:"SmsSdkAppId"`

	// 发送短信返回的流水号。
	SerialNo *string `json:"SerialNo,omitnil" name:"SerialNo"`

	// 用户回填时间，UNIX 时间戳（单位：秒）。
	ConversionTime *uint64 `json:"ConversionTime,omitnil" name:"ConversionTime"`
}

type ReportConversionRequest struct {
	*tchttp.BaseRequest
	
	// 短信应用ID。在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage)  添加应用后生成的实际 SdkAppId，示例如1400006666。
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil" name:"SmsSdkAppId"`

	// 发送短信返回的流水号。
	SerialNo *string `json:"SerialNo,omitnil" name:"SerialNo"`

	// 用户回填时间，UNIX 时间戳（单位：秒）。
	ConversionTime *uint64 `json:"ConversionTime,omitnil" name:"ConversionTime"`
}

func (r *ReportConversionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportConversionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SmsSdkAppId")
	delete(f, "SerialNo")
	delete(f, "ConversionTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportConversionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportConversionResponseParams struct {
	// 转化率上报响应包体。
	ReportConversionStatus *ReportConversionStatus `json:"ReportConversionStatus,omitnil" name:"ReportConversionStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ReportConversionResponse struct {
	*tchttp.BaseResponse
	Response *ReportConversionResponseParams `json:"Response"`
}

func (r *ReportConversionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportConversionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReportConversionStatus struct {
	// 错误码。上报成功返回 ok。
	Code *string `json:"Code,omitnil" name:"Code"`

	// 错误码描述。
	Message *string `json:"Message,omitnil" name:"Message"`
}

// Predefined struct for user
type SendSmsRequestParams struct {
	// 下发手机号码，采用 E.164 标准，格式为+[国家或地区码][手机号]，单次请求最多支持200个手机号且要求全为境内手机号或全为境外手机号。
	// 例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	// 注：发送国内短信格式还支持0086、86或无任何国家或地区码的11位手机号码，前缀默认为+86。
	PhoneNumberSet []*string `json:"PhoneNumberSet,omitnil" name:"PhoneNumberSet"`

	// 短信 SdkAppId，在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage)  添加应用后生成的实际 SdkAppId，示例如1400006666。
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil" name:"SmsSdkAppId"`

	// 模板 ID，必须填写已审核通过的模板 ID。模板 ID 可前往 [国内短信](https://console.cloud.tencent.com/smsv2/csms-template) 或 [国际/港澳台短信](https://console.cloud.tencent.com/smsv2/isms-template) 的正文模板管理查看，若向境外手机号发送短信，仅支持使用国际/港澳台短信模板。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 短信签名内容，使用 UTF-8 编码，必须填写已审核通过的签名，例如：腾讯云，签名信息可前往 [国内短信](https://console.cloud.tencent.com/smsv2/csms-sign) 或 [国际/港澳台短信](https://console.cloud.tencent.com/smsv2/isms-sign) 的签名管理查看。
	// <dx-alert infotype="notice" title="注意">发送国内短信该参数必填，且需填写签名内容而非签名ID。</dx-alert>
	SignName *string `json:"SignName,omitnil" name:"SignName"`

	// 模板参数，若无模板参数，则设置为空。
	// <dx-alert infotype="notice" title="注意">模板参数的个数需要与 TemplateId 对应模板的变量个数保持一致。</dx-alert>
	TemplateParamSet []*string `json:"TemplateParamSet,omitnil" name:"TemplateParamSet"`

	// 短信码号扩展号，默认未开通，如需开通请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	ExtendCode *string `json:"ExtendCode,omitnil" name:"ExtendCode"`

	// 用户的 session 内容，可以携带用户侧 ID 等上下文信息，server 会原样返回。注意长度需小于512字节。
	SessionContext *string `json:"SessionContext,omitnil" name:"SessionContext"`

	// 国内短信无需填写该项；国际/港澳台短信已申请独立 SenderId 需要填写该字段，默认使用公共 SenderId，无需填写该字段。
	// 注：月度使用量达到指定量级可申请独立 SenderId 使用，详情请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	SenderId *string `json:"SenderId,omitnil" name:"SenderId"`
}

type SendSmsRequest struct {
	*tchttp.BaseRequest
	
	// 下发手机号码，采用 E.164 标准，格式为+[国家或地区码][手机号]，单次请求最多支持200个手机号且要求全为境内手机号或全为境外手机号。
	// 例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	// 注：发送国内短信格式还支持0086、86或无任何国家或地区码的11位手机号码，前缀默认为+86。
	PhoneNumberSet []*string `json:"PhoneNumberSet,omitnil" name:"PhoneNumberSet"`

	// 短信 SdkAppId，在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage)  添加应用后生成的实际 SdkAppId，示例如1400006666。
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil" name:"SmsSdkAppId"`

	// 模板 ID，必须填写已审核通过的模板 ID。模板 ID 可前往 [国内短信](https://console.cloud.tencent.com/smsv2/csms-template) 或 [国际/港澳台短信](https://console.cloud.tencent.com/smsv2/isms-template) 的正文模板管理查看，若向境外手机号发送短信，仅支持使用国际/港澳台短信模板。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 短信签名内容，使用 UTF-8 编码，必须填写已审核通过的签名，例如：腾讯云，签名信息可前往 [国内短信](https://console.cloud.tencent.com/smsv2/csms-sign) 或 [国际/港澳台短信](https://console.cloud.tencent.com/smsv2/isms-sign) 的签名管理查看。
	// <dx-alert infotype="notice" title="注意">发送国内短信该参数必填，且需填写签名内容而非签名ID。</dx-alert>
	SignName *string `json:"SignName,omitnil" name:"SignName"`

	// 模板参数，若无模板参数，则设置为空。
	// <dx-alert infotype="notice" title="注意">模板参数的个数需要与 TemplateId 对应模板的变量个数保持一致。</dx-alert>
	TemplateParamSet []*string `json:"TemplateParamSet,omitnil" name:"TemplateParamSet"`

	// 短信码号扩展号，默认未开通，如需开通请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	ExtendCode *string `json:"ExtendCode,omitnil" name:"ExtendCode"`

	// 用户的 session 内容，可以携带用户侧 ID 等上下文信息，server 会原样返回。注意长度需小于512字节。
	SessionContext *string `json:"SessionContext,omitnil" name:"SessionContext"`

	// 国内短信无需填写该项；国际/港澳台短信已申请独立 SenderId 需要填写该字段，默认使用公共 SenderId，无需填写该字段。
	// 注：月度使用量达到指定量级可申请独立 SenderId 使用，详情请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	SenderId *string `json:"SenderId,omitnil" name:"SenderId"`
}

func (r *SendSmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendSmsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PhoneNumberSet")
	delete(f, "SmsSdkAppId")
	delete(f, "TemplateId")
	delete(f, "SignName")
	delete(f, "TemplateParamSet")
	delete(f, "ExtendCode")
	delete(f, "SessionContext")
	delete(f, "SenderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendSmsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendSmsResponseParams struct {
	// 短信发送状态。
	SendStatusSet []*SendStatus `json:"SendStatusSet,omitnil" name:"SendStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SendSmsResponse struct {
	*tchttp.BaseResponse
	Response *SendSmsResponseParams `json:"Response"`
}

func (r *SendSmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendSmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendStatus struct {
	// 发送流水号。
	SerialNo *string `json:"SerialNo,omitnil" name:"SerialNo"`

	// 手机号码，E.164标准，+[国家或地区码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumber *string `json:"PhoneNumber,omitnil" name:"PhoneNumber"`

	// 计费条数，计费规则请查询 [计费策略](https://cloud.tencent.com/document/product/382/36135)。
	Fee *uint64 `json:"Fee,omitnil" name:"Fee"`

	// 用户 session 内容。
	SessionContext *string `json:"SessionContext,omitnil" name:"SessionContext"`

	// 短信请求错误码，具体含义请参考 [错误码](https://cloud.tencent.com/document/api/382/55981#6.-.E9.94.99.E8.AF.AF.E7.A0.81)，发送成功返回 "Ok"。
	Code *string `json:"Code,omitnil" name:"Code"`

	// 短信请求错误码描述。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 国家码或地区码，例如 CN、US 等，对于未识别出国家码或者地区码，默认返回 DEF，具体支持列表请参考 [国际/港澳台短信价格总览](https://cloud.tencent.com/document/product/382/18051)。
	IsoCode *string `json:"IsoCode,omitnil" name:"IsoCode"`
}

type SendStatusStatistics struct {
	// 短信计费条数统计，例如提交成功量为100条，其中有20条是长短信（长度为80字）被拆分成2条，则计费条数为： ```80 * 1 + 20 * 2 = 120``` 条。
	FeeCount *uint64 `json:"FeeCount,omitnil" name:"FeeCount"`

	// 短信提交量统计。
	RequestCount *uint64 `json:"RequestCount,omitnil" name:"RequestCount"`

	// 短信提交成功量统计。
	RequestSuccessCount *uint64 `json:"RequestSuccessCount,omitnil" name:"RequestSuccessCount"`
}

// Predefined struct for user
type SendStatusStatisticsRequestParams struct {
	// 起始时间，格式为yyyymmddhh，精确到小时，例如2021050113，表示2021年5月1号13时。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间，格式为yyyymmddhh，精确到小时，例如2021050118，表示2021年5月1号18时。
	// 注：EndTime 必须大于 BeginTime。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 短信 SdkAppId 在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage)  添加应用后生成的实际 SdkAppId，示例如1400006666。
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil" name:"SmsSdkAppId"`

	// 最大上限。
	// 注：目前固定设置为0。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type SendStatusStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间，格式为yyyymmddhh，精确到小时，例如2021050113，表示2021年5月1号13时。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间，格式为yyyymmddhh，精确到小时，例如2021050118，表示2021年5月1号18时。
	// 注：EndTime 必须大于 BeginTime。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 短信 SdkAppId 在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage)  添加应用后生成的实际 SdkAppId，示例如1400006666。
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil" name:"SmsSdkAppId"`

	// 最大上限。
	// 注：目前固定设置为0。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *SendStatusStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendStatusStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "SmsSdkAppId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendStatusStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendStatusStatisticsResponseParams struct {
	// 发送数据统计响应包体。
	SendStatusStatistics *SendStatusStatistics `json:"SendStatusStatistics,omitnil" name:"SendStatusStatistics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SendStatusStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *SendStatusStatisticsResponseParams `json:"Response"`
}

func (r *SendStatusStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendStatusStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SmsPackagesStatistics struct {
	// 套餐包创建时间，UNIX 时间戳（单位：秒）。
	PackageCreateTime *uint64 `json:"PackageCreateTime,omitnil" name:"PackageCreateTime"`

	// 套餐包生效时间，UNIX 时间戳（单位：秒）。
	PackageEffectiveTime *uint64 `json:"PackageEffectiveTime,omitnil" name:"PackageEffectiveTime"`

	// 套餐包过期时间，UNIX 时间戳（单位：秒）。
	PackageExpiredTime *uint64 `json:"PackageExpiredTime,omitnil" name:"PackageExpiredTime"`

	// 套餐包条数。
	PackageAmount *uint64 `json:"PackageAmount,omitnil" name:"PackageAmount"`

	// 套餐包类别，0表示赠送套餐包，1表示购买套餐包。
	PackageType *uint64 `json:"PackageType,omitnil" name:"PackageType"`

	// 套餐包 ID。
	PackageId *uint64 `json:"PackageId,omitnil" name:"PackageId"`

	// 当前使用套餐包条数。
	CurrentUsage *uint64 `json:"CurrentUsage,omitnil" name:"CurrentUsage"`
}

// Predefined struct for user
type SmsPackagesStatisticsRequestParams struct {
	// 短信 SdkAppId 在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage)  添加应用后生成的实际 SdkAppId，示例如1400006666。
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil" name:"SmsSdkAppId"`

	// 最大上限(需要拉取的套餐包个数)。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 起始时间，格式为yyyymmddhh，精确到小时，例如2021050113，表示2021年5月1号13时。
	// 注：接口会返回 BeginTime 到 EndTime 之间创建的套餐包的统计信息。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间，格式为yyyymmddhh，精确到小时，例如2021050118，表示2021年5月1号18时。
	// 注：EndTime 必须大于 BeginTime 且小于当前时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type SmsPackagesStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 短信 SdkAppId 在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage)  添加应用后生成的实际 SdkAppId，示例如1400006666。
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil" name:"SmsSdkAppId"`

	// 最大上限(需要拉取的套餐包个数)。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 起始时间，格式为yyyymmddhh，精确到小时，例如2021050113，表示2021年5月1号13时。
	// 注：接口会返回 BeginTime 到 EndTime 之间创建的套餐包的统计信息。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间，格式为yyyymmddhh，精确到小时，例如2021050118，表示2021年5月1号18时。
	// 注：EndTime 必须大于 BeginTime 且小于当前时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *SmsPackagesStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SmsPackagesStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SmsSdkAppId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SmsPackagesStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SmsPackagesStatisticsResponseParams struct {
	// 发送数据统计响应包体。
	SmsPackagesStatisticsSet []*SmsPackagesStatistics `json:"SmsPackagesStatisticsSet,omitnil" name:"SmsPackagesStatisticsSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SmsPackagesStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *SmsPackagesStatisticsResponseParams `json:"Response"`
}

func (r *SmsPackagesStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SmsPackagesStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}