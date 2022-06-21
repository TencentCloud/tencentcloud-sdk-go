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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddSignStatus struct {
	// 签名Id。
	SignId *uint64 `json:"SignId,omitempty" name:"SignId"`

	// 签名申请Id。
	SignApplyId *uint64 `json:"SignApplyId,omitempty" name:"SignApplyId"`
}

// Predefined struct for user
type AddSmsSignRequestParams struct {
	// 签名名称。
	// 注：不能重复申请已通过或待审核的签名。
	SignName *string `json:"SignName,omitempty" name:"SignName"`

	// 签名类型。其中每种类型后面标注了其可选的 DocumentType（证明类型）：
	// 0：公司，可选 DocumentType 有（0，1）。
	// 1：APP，可选 DocumentType 有（0，1，2，3，4） 。
	// 2：网站，可选 DocumentType 有（0，1，2，3，5）。
	// 3：公众号，可选 DocumentType 有（0，1，2，3，8）。
	// 4：商标，可选 DocumentType 有（7）。
	// 5：政府/机关事业单位/其他机构，可选 DocumentType 有（2，3）。
	// 6：小程序，可选 DocumentType 有（0，1，2，3，6）。
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
	// 8：公众号设置页面截图（个人认证公众号）。
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

type AddSmsSignRequest struct {
	*tchttp.BaseRequest
	
	// 签名名称。
	// 注：不能重复申请已通过或待审核的签名。
	SignName *string `json:"SignName,omitempty" name:"SignName"`

	// 签名类型。其中每种类型后面标注了其可选的 DocumentType（证明类型）：
	// 0：公司，可选 DocumentType 有（0，1）。
	// 1：APP，可选 DocumentType 有（0，1，2，3，4） 。
	// 2：网站，可选 DocumentType 有（0，1，2，3，5）。
	// 3：公众号，可选 DocumentType 有（0，1，2，3，8）。
	// 4：商标，可选 DocumentType 有（7）。
	// 5：政府/机关事业单位/其他机构，可选 DocumentType 有（2，3）。
	// 6：小程序，可选 DocumentType 有（0，1，2，3，6）。
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
	// 8：公众号设置页面截图（个人认证公众号）。
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
	delete(f, "UsedMethod")
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
	AddSignStatus *AddSignStatus `json:"AddSignStatus,omitempty" name:"AddSignStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	AddTemplateStatus *AddTemplateStatus `json:"AddTemplateStatus,omitempty" name:"AddTemplateStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CallbackStatusStatisticsRequestParams struct {
	// 开始时间，yyyymmddhh 需要拉取的起始时间，精确到小时。
	StartDateTime *uint64 `json:"StartDateTime,omitempty" name:"StartDateTime"`

	// 结束时间，yyyymmddhh 需要拉取的截止时间，精确到小时。
	// 注：EndDataTime 必须大于 StartDateTime。
	EndDataTime *uint64 `json:"EndDataTime,omitempty" name:"EndDataTime"`

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/smsv2) 添加应用后生成的实际SdkAppid，示例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`

	// 最大上限。
	// 注：目前固定设置为0。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type CallbackStatusStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，yyyymmddhh 需要拉取的起始时间，精确到小时。
	StartDateTime *uint64 `json:"StartDateTime,omitempty" name:"StartDateTime"`

	// 结束时间，yyyymmddhh 需要拉取的截止时间，精确到小时。
	// 注：EndDataTime 必须大于 StartDateTime。
	EndDataTime *uint64 `json:"EndDataTime,omitempty" name:"EndDataTime"`

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/smsv2) 添加应用后生成的实际SdkAppid，示例如1400006666。
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallbackStatusStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartDateTime")
	delete(f, "EndDataTime")
	delete(f, "SmsSdkAppid")
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
	CallbackStatusStatistics *CallbackStatusStatistics `json:"CallbackStatusStatistics,omitempty" name:"CallbackStatusStatistics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DeleteStatus *string `json:"DeleteStatus,omitempty" name:"DeleteStatus"`

	// 删除时间，UNIX 时间戳（单位：秒）。
	DeleteTime *uint64 `json:"DeleteTime,omitempty" name:"DeleteTime"`
}

// Predefined struct for user
type DeleteSmsSignRequestParams struct {
	// 待删除的签名 ID。
	SignId *uint64 `json:"SignId,omitempty" name:"SignId"`
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
	DeleteSignStatus *DeleteSignStatus `json:"DeleteSignStatus,omitempty" name:"DeleteSignStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TemplateId *uint64 `json:"TemplateId,omitempty" name:"TemplateId"`
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
	DeleteTemplateStatus *DeleteTemplateStatus `json:"DeleteTemplateStatus,omitempty" name:"DeleteTemplateStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 1：表示审核中。
	// -1：表示审核未通过或审核失败。
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 审核回复，审核人员审核后给出的回复，通常是审核未通过的原因。
	ReviewReply *string `json:"ReviewReply,omitempty" name:"ReviewReply"`

	// 签名名称。
	SignName *string `json:"SignName,omitempty" name:"SignName"`

	// 提交审核时间，UNIX 时间戳（单位：秒）。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type DescribeSmsSignListRequestParams struct {
	// 签名 ID 数组。
	SignIdSet []*uint64 `json:"SignIdSet,omitempty" name:"SignIdSet"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`
}

type DescribeSmsSignListRequest struct {
	*tchttp.BaseRequest
	
	// 签名 ID 数组。
	SignIdSet []*uint64 `json:"SignIdSet,omitempty" name:"SignIdSet"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`
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
	DescribeSignListStatusSet []*DescribeSignListStatus `json:"DescribeSignListStatusSet,omitempty" name:"DescribeSignListStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 模板 ID 数组。
	TemplateIdSet []*uint64 `json:"TemplateIdSet,omitempty" name:"TemplateIdSet"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`
}

type DescribeSmsTemplateListRequest struct {
	*tchttp.BaseRequest
	
	// 模板 ID 数组。
	TemplateIdSet []*uint64 `json:"TemplateIdSet,omitempty" name:"TemplateIdSet"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`
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
	delete(f, "TemplateIdSet")
	delete(f, "International")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSmsTemplateListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmsTemplateListResponseParams struct {
	// 获取短信模板信息响应
	DescribeTemplateStatusSet []*DescribeTemplateListStatus `json:"DescribeTemplateStatusSet,omitempty" name:"DescribeTemplateStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 模板Id
	TemplateId *uint64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`

	// 申请签名状态。其中：
	// 0：表示审核通过。
	// 1：表示审核中。
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

// Predefined struct for user
type ModifySmsSignRequestParams struct {
	// 待修改的签名 ID。
	SignId *uint64 `json:"SignId,omitempty" name:"SignId"`

	// 签名名称。
	SignName *string `json:"SignName,omitempty" name:"SignName"`

	// 签名类型。其中每种类型后面标注了其可选的 DocumentType（证明类型）：
	// 0：公司，可选 DocumentType 有（0，1）。
	// 1：APP，可选 DocumentType 有（0，1，2，3，4） 。
	// 2：网站，可选 DocumentType 有（0，1，2，3，5）。
	// 3：公众号，可选 DocumentType 有（0，1，2，3，8）。
	// 4：商标，可选 DocumentType 有（7）。
	// 5：政府/机关事业单位/其他机构，可选 DocumentType 有（2，3）。
	// 6：小程序，可选 DocumentType 有（0，1，2，3，6）。
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
	// 8：公众号设置页面截图（个人认证公众号）。
	DocumentType *uint64 `json:"DocumentType,omitempty" name:"DocumentType"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	// 注：需要和待修改签名International值保持一致，该参数不能直接修改国内签名到国际签名。
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

type ModifySmsSignRequest struct {
	*tchttp.BaseRequest
	
	// 待修改的签名 ID。
	SignId *uint64 `json:"SignId,omitempty" name:"SignId"`

	// 签名名称。
	SignName *string `json:"SignName,omitempty" name:"SignName"`

	// 签名类型。其中每种类型后面标注了其可选的 DocumentType（证明类型）：
	// 0：公司，可选 DocumentType 有（0，1）。
	// 1：APP，可选 DocumentType 有（0，1，2，3，4） 。
	// 2：网站，可选 DocumentType 有（0，1，2，3，5）。
	// 3：公众号，可选 DocumentType 有（0，1，2，3，8）。
	// 4：商标，可选 DocumentType 有（7）。
	// 5：政府/机关事业单位/其他机构，可选 DocumentType 有（2，3）。
	// 6：小程序，可选 DocumentType 有（0，1，2，3，6）。
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
	// 8：公众号设置页面截图（个人认证公众号）。
	DocumentType *uint64 `json:"DocumentType,omitempty" name:"DocumentType"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	// 注：需要和待修改签名International值保持一致，该参数不能直接修改国内签名到国际签名。
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
	delete(f, "UsedMethod")
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
	ModifySignStatus *ModifySignStatus `json:"ModifySignStatus,omitempty" name:"ModifySignStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ModifyTemplateStatus *ModifyTemplateStatus `json:"ModifyTemplateStatus,omitempty" name:"ModifyTemplateStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type PullSmsReplyStatusByPhoneNumberRequestParams struct {
	// 拉取起始时间，UNIX 时间戳（时间：秒）。
	// 注：最大可拉取当前时期7天前的数据。
	SendDateTime *uint64 `json:"SendDateTime,omitempty" name:"SendDateTime"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 拉取最大条数，最多 100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下发目的手机号码，依据 e.164 标准为：+[国家（或地区）码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/smsv2) 添加应用后生成的实际SdkAppid，例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`

	// 拉取截止时间，UNIX 时间戳（时间：秒）。
	EndDateTime *uint64 `json:"EndDateTime,omitempty" name:"EndDateTime"`
}

type PullSmsReplyStatusByPhoneNumberRequest struct {
	*tchttp.BaseRequest
	
	// 拉取起始时间，UNIX 时间戳（时间：秒）。
	// 注：最大可拉取当前时期7天前的数据。
	SendDateTime *uint64 `json:"SendDateTime,omitempty" name:"SendDateTime"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 拉取最大条数，最多 100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下发目的手机号码，依据 e.164 标准为：+[国家（或地区）码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/smsv2) 添加应用后生成的实际SdkAppid，例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`

	// 拉取截止时间，UNIX 时间戳（时间：秒）。
	EndDateTime *uint64 `json:"EndDateTime,omitempty" name:"EndDateTime"`
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
	delete(f, "SendDateTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PhoneNumber")
	delete(f, "SmsSdkAppid")
	delete(f, "EndDateTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PullSmsReplyStatusByPhoneNumberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullSmsReplyStatusByPhoneNumberResponseParams struct {
	// 回复状态响应集合。
	PullSmsReplyStatusSet []*PullSmsReplyStatus `json:"PullSmsReplyStatusSet,omitempty" name:"PullSmsReplyStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 短信 SdkAppid 在 [短信控制台](https://console.cloud.tencent.com/smsv2) 添加应用后生成的实际 SdkAppid，例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`
}

type PullSmsReplyStatusRequest struct {
	*tchttp.BaseRequest
	
	// 拉取最大条数，最多100条。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 短信 SdkAppid 在 [短信控制台](https://console.cloud.tencent.com/smsv2) 添加应用后生成的实际 SdkAppid，例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`
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
	delete(f, "SmsSdkAppid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PullSmsReplyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullSmsReplyStatusResponseParams struct {
	// 回复状态响应集合。
	PullSmsReplyStatusSet []*PullSmsReplyStatus `json:"PullSmsReplyStatusSet,omitempty" name:"PullSmsReplyStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type PullSmsSendStatusByPhoneNumberRequestParams struct {
	// 拉取起始时间，UNIX 时间戳（时间：秒）。
	// 注：最大可拉取当前时期7天前的数据。
	SendDateTime *uint64 `json:"SendDateTime,omitempty" name:"SendDateTime"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 拉取最大条数，最多 100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下发目的手机号码，依据 e.164 标准为：+[国家（或地区）码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/smsv2) 添加应用后生成的实际SdkAppid，例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`

	// 拉取截止时间，UNIX 时间戳（时间：秒）。
	EndDateTime *uint64 `json:"EndDateTime,omitempty" name:"EndDateTime"`
}

type PullSmsSendStatusByPhoneNumberRequest struct {
	*tchttp.BaseRequest
	
	// 拉取起始时间，UNIX 时间戳（时间：秒）。
	// 注：最大可拉取当前时期7天前的数据。
	SendDateTime *uint64 `json:"SendDateTime,omitempty" name:"SendDateTime"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 拉取最大条数，最多 100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下发目的手机号码，依据 e.164 标准为：+[国家（或地区）码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/smsv2) 添加应用后生成的实际SdkAppid，例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`

	// 拉取截止时间，UNIX 时间戳（时间：秒）。
	EndDateTime *uint64 `json:"EndDateTime,omitempty" name:"EndDateTime"`
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
	delete(f, "SendDateTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PhoneNumber")
	delete(f, "SmsSdkAppid")
	delete(f, "EndDateTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PullSmsSendStatusByPhoneNumberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullSmsSendStatusByPhoneNumberResponseParams struct {
	// 下发状态响应集合。
	PullSmsSendStatusSet []*PullSmsSendStatus `json:"PullSmsSendStatusSet,omitempty" name:"PullSmsSendStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/smsv2) 添加应用后生成的实际SdkAppid，例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`
}

type PullSmsSendStatusRequest struct {
	*tchttp.BaseRequest
	
	// 拉取最大条数，最多100条。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/smsv2) 添加应用后生成的实际SdkAppid，例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`
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
	delete(f, "SmsSdkAppid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PullSmsSendStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullSmsSendStatusResponseParams struct {
	// 下发状态响应集合。
	PullSmsSendStatusSet []*PullSmsSendStatus `json:"PullSmsSendStatusSet,omitempty" name:"PullSmsSendStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type SendSmsRequestParams struct {
	// 下发手机号码，采用 e.164 标准，格式为+[国家或地区码][手机号]，单次请求最多支持200个手机号且要求全为境内手机号或全为境外手机号。
	// 例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumberSet []*string `json:"PhoneNumberSet,omitempty" name:"PhoneNumberSet"`

	// 模板 ID，必须填写已审核通过的模板 ID。模板ID可登录 [短信控制台](https://console.cloud.tencent.com/smsv2) 查看，若向境外手机号发送短信，仅支持使用国际/港澳台短信模板。
	TemplateID *string `json:"TemplateID,omitempty" name:"TemplateID"`

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/smsv2)  添加应用后生成的实际SdkAppid，示例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`

	// 短信签名内容，使用 UTF-8 编码，必须填写已审核通过的签名，签名信息可登录 [短信控制台](https://console.cloud.tencent.com/smsv2)  查看。注：国内短信为必填参数。
	Sign *string `json:"Sign,omitempty" name:"Sign"`

	// 模板参数，若无模板参数，则设置为空。
	TemplateParamSet []*string `json:"TemplateParamSet,omitempty" name:"TemplateParamSet"`

	// 短信码号扩展号，默认未开通，如需开通请联系 [sms helper](https://cloud.tencent.com/document/product/382/3773)。
	ExtendCode *string `json:"ExtendCode,omitempty" name:"ExtendCode"`

	// 用户的 session 内容，可以携带用户侧 ID 等上下文信息，server 会原样返回。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// 国内短信无senderid，无需填写该项；若需开通国际/港澳台短信senderid，请联系smshelper。
	SenderId *string `json:"SenderId,omitempty" name:"SenderId"`
}

type SendSmsRequest struct {
	*tchttp.BaseRequest
	
	// 下发手机号码，采用 e.164 标准，格式为+[国家或地区码][手机号]，单次请求最多支持200个手机号且要求全为境内手机号或全为境外手机号。
	// 例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumberSet []*string `json:"PhoneNumberSet,omitempty" name:"PhoneNumberSet"`

	// 模板 ID，必须填写已审核通过的模板 ID。模板ID可登录 [短信控制台](https://console.cloud.tencent.com/smsv2) 查看，若向境外手机号发送短信，仅支持使用国际/港澳台短信模板。
	TemplateID *string `json:"TemplateID,omitempty" name:"TemplateID"`

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/smsv2)  添加应用后生成的实际SdkAppid，示例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`

	// 短信签名内容，使用 UTF-8 编码，必须填写已审核通过的签名，签名信息可登录 [短信控制台](https://console.cloud.tencent.com/smsv2)  查看。注：国内短信为必填参数。
	Sign *string `json:"Sign,omitempty" name:"Sign"`

	// 模板参数，若无模板参数，则设置为空。
	TemplateParamSet []*string `json:"TemplateParamSet,omitempty" name:"TemplateParamSet"`

	// 短信码号扩展号，默认未开通，如需开通请联系 [sms helper](https://cloud.tencent.com/document/product/382/3773)。
	ExtendCode *string `json:"ExtendCode,omitempty" name:"ExtendCode"`

	// 用户的 session 内容，可以携带用户侧 ID 等上下文信息，server 会原样返回。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// 国内短信无senderid，无需填写该项；若需开通国际/港澳台短信senderid，请联系smshelper。
	SenderId *string `json:"SenderId,omitempty" name:"SenderId"`
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
	delete(f, "TemplateID")
	delete(f, "SmsSdkAppid")
	delete(f, "Sign")
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
	SendStatusSet []*SendStatus `json:"SendStatusSet,omitempty" name:"SendStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

	// 国家码或地区码，例如CN,US等，对于未识别出国家码或者地区码，默认返回DEF,具体支持列表请参考国际/港澳台计费总览。
	IsoCode *string `json:"IsoCode,omitempty" name:"IsoCode"`
}

type SendStatusStatistics struct {
	// 短信计费条数统计，例如提交成功量为100条，其中有20条是长短信（长度为80字）被拆分成2条，则计费条数为： ```80 * 1 + 20 * 2 = 120``` 条。
	FeeCount *uint64 `json:"FeeCount,omitempty" name:"FeeCount"`

	// 短信提交量统计。
	RequestCount *uint64 `json:"RequestCount,omitempty" name:"RequestCount"`

	// 短信提交成功量统计。
	RequestSuccessCount *uint64 `json:"RequestSuccessCount,omitempty" name:"RequestSuccessCount"`
}

// Predefined struct for user
type SendStatusStatisticsRequestParams struct {
	// 拉取起始时间，yyyymmddhh 需要拉取的起始时间，精确到小时。
	StartDateTime *uint64 `json:"StartDateTime,omitempty" name:"StartDateTime"`

	// 结束时间，yyyymmddhh 需要拉取的截止时间，精确到小时
	// 注：EndDataTime 必须大于 StartDateTime。
	EndDataTime *uint64 `json:"EndDataTime,omitempty" name:"EndDataTime"`

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/smsv2) 添加应用后生成的实际SdkAppid，示例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`

	// 最大上限。
	// 注：目前固定设置为0。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type SendStatusStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 拉取起始时间，yyyymmddhh 需要拉取的起始时间，精确到小时。
	StartDateTime *uint64 `json:"StartDateTime,omitempty" name:"StartDateTime"`

	// 结束时间，yyyymmddhh 需要拉取的截止时间，精确到小时
	// 注：EndDataTime 必须大于 StartDateTime。
	EndDataTime *uint64 `json:"EndDataTime,omitempty" name:"EndDataTime"`

	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/smsv2) 添加应用后生成的实际SdkAppid，示例如1400006666。
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendStatusStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartDateTime")
	delete(f, "EndDataTime")
	delete(f, "SmsSdkAppid")
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
	SendStatusStatistics *SendStatusStatistics `json:"SendStatusStatistics,omitempty" name:"SendStatusStatistics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type SmsPackagesStatisticsRequestParams struct {
	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/smsv2) 添加应用后生成的实际SdkAppid，示例如1400006666。
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`

	// 最大上限(需要拉取的套餐包个数)。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量。
	// 注：目前固定设置为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type SmsPackagesStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 短信SdkAppid在 [短信控制台](https://console.cloud.tencent.com/smsv2) 添加应用后生成的实际SdkAppid，示例如1400006666。
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SmsPackagesStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SmsSdkAppid")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SmsPackagesStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SmsPackagesStatisticsResponseParams struct {
	// 发送数据统计响应包体。
	SmsPackagesStatisticsSet []*SmsPackagesStatistics `json:"SmsPackagesStatisticsSet,omitempty" name:"SmsPackagesStatisticsSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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