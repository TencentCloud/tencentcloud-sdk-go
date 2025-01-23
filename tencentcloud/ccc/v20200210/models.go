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

package v20200210

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AICallExtractConfigElement struct {
	// 配置项类型，包括
	// Text 文本
	// Selector 选项
	// Boolean 布尔值
	// Number 数字
	InfoType *string `json:"InfoType,omitnil,omitempty" name:"InfoType"`

	// 配置项名称，不可重复
	InfoName *string `json:"InfoName,omitnil,omitempty" name:"InfoName"`

	// 配置项具体内容
	InfoContent *string `json:"InfoContent,omitnil,omitempty" name:"InfoContent"`

	// 配置项提取内容示例
	Examples []*string `json:"Examples,omitnil,omitempty" name:"Examples"`

	// InfoType 为 Selector，需要配置此字段
	Choices []*string `json:"Choices,omitnil,omitempty" name:"Choices"`
}

type AICallExtractResultElement struct {
	// 提取信息的类型
	// Text 文本
	// Selector 选项
	// Boolean 布尔值
	// Number 数字
	InfoType *string `json:"InfoType,omitnil,omitempty" name:"InfoType"`

	// 提取信息的名称
	InfoName *string `json:"InfoName,omitnil,omitempty" name:"InfoName"`

	// 提取信息的具体描述
	InfoContent *string `json:"InfoContent,omitnil,omitempty" name:"InfoContent"`

	// 提取信息的具体结果
	Result *AICallExtractResultInfo `json:"Result,omitnil,omitempty" name:"Result"`
}

type AICallExtractResultInfo struct {
	// 提取的类型是文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 提取的内型是选项
	Chosen []*string `json:"Chosen,omitnil,omitempty" name:"Chosen"`

	// 提取类型是布尔值
	Boolean *bool `json:"Boolean,omitnil,omitempty" name:"Boolean"`

	// 提取类型是数字
	Number *float64 `json:"Number,omitnil,omitempty" name:"Number"`
}

type AITransferItem struct {
	// 转人工的function calling 名称
	TransferFunctionName *string `json:"TransferFunctionName,omitnil,omitempty" name:"TransferFunctionName"`

	// TransferFunctionEnable为true时生效；transfer_to_human function calling的desc，默认为 "Transfer to human when the user has to transfer to human (like says transfer to human) or you are instructed to do so."
	TransferFunctionDesc *string `json:"TransferFunctionDesc,omitnil,omitempty" name:"TransferFunctionDesc"`

	// 转人工的技能组ID
	TransferSkillGroupId *uint64 `json:"TransferSkillGroupId,omitnil,omitempty" name:"TransferSkillGroupId"`
}

// Predefined struct for user
type AbortPredictiveDialingCampaignRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务 ID
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

type AbortPredictiveDialingCampaignRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务 ID
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

func (r *AbortPredictiveDialingCampaignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AbortPredictiveDialingCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CampaignId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AbortPredictiveDialingCampaignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AbortPredictiveDialingCampaignResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AbortPredictiveDialingCampaignResponse struct {
	*tchttp.BaseResponse
	Response *AbortPredictiveDialingCampaignResponseParams `json:"Response"`
}

func (r *AbortPredictiveDialingCampaignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AbortPredictiveDialingCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActiveCarrierPrivilegeNumber struct {
	// 实例Id
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 主叫号码
	Caller *string `json:"Caller,omitnil,omitempty" name:"Caller"`

	// 被叫号码
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// 生效unix时间戳(秒)
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type AsrData struct {
	// 用户方
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 消息内容
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 时间戳
	//
	// Deprecated: Timestamp is deprecated.
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 句子开始时间，Unix 毫秒时间戳
	Start *int64 `json:"Start,omitnil,omitempty" name:"Start"`

	// 句子结束时间，Unix 毫秒时间戳
	End *int64 `json:"End,omitnil,omitempty" name:"End"`
}

type AudioFileInfo struct {
	// 文件ID
	FileId *uint64 `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 文件别名
	CustomFileName *string `json:"CustomFileName,omitnil,omitempty" name:"CustomFileName"`

	// 文件名
	AudioFileName *string `json:"AudioFileName,omitnil,omitempty" name:"AudioFileName"`

	// 审核状态，0-未审核，1-审核通过，2-审核拒绝
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type AutoCalloutTaskCalleeInfo struct {
	// 被叫号码
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// 呼叫状态 0初始 1已接听 2未接听 3呼叫中 4待重试
	State *uint64 `json:"State,omitnil,omitempty" name:"State"`

	// 会话ID列表
	Sessions []*string `json:"Sessions,omitnil,omitempty" name:"Sessions"`
}

type AutoCalloutTaskInfo struct {
	// 任务名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 被叫数量
	CalleeCount *uint64 `json:"CalleeCount,omitnil,omitempty" name:"CalleeCount"`

	// 主叫号码列表
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 起始时间戳
	NotBefore *int64 `json:"NotBefore,omitnil,omitempty" name:"NotBefore"`

	// 结束时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotAfter *int64 `json:"NotAfter,omitnil,omitempty" name:"NotAfter"`

	// 任务使用的IvrId
	IvrId *uint64 `json:"IvrId,omitnil,omitempty" name:"IvrId"`

	// 任务状态：
	// 0初始：任务创建，呼叫未开始
	// 1运行中
	// 2 已完成：任务中所有呼叫完成
	// 3结束中：任务到期，但仍有部分呼叫未结束
	// 4已结束：任务到期终止
	State *uint64 `json:"State,omitnil,omitempty" name:"State"`

	// 任务Id
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type BindNumberCallOutSkillGroupRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 待绑定的号码
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// 待绑定的技能组Id列表
	SkillGroupIds []*uint64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`
}

type BindNumberCallOutSkillGroupRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 待绑定的号码
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// 待绑定的技能组Id列表
	SkillGroupIds []*uint64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`
}

func (r *BindNumberCallOutSkillGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindNumberCallOutSkillGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Number")
	delete(f, "SkillGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindNumberCallOutSkillGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindNumberCallOutSkillGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindNumberCallOutSkillGroupResponse struct {
	*tchttp.BaseResponse
	Response *BindNumberCallOutSkillGroupResponseParams `json:"Response"`
}

func (r *BindNumberCallOutSkillGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindNumberCallOutSkillGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindStaffSkillGroupListRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 座席邮箱
	StaffEmail *string `json:"StaffEmail,omitnil,omitempty" name:"StaffEmail"`

	// 绑定技能组列表
	//
	// Deprecated: SkillGroupList is deprecated.
	SkillGroupList []*int64 `json:"SkillGroupList,omitnil,omitempty" name:"SkillGroupList"`

	// 绑定技能组列表(必填)
	StaffSkillGroupList []*StaffSkillGroupList `json:"StaffSkillGroupList,omitnil,omitempty" name:"StaffSkillGroupList"`
}

type BindStaffSkillGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 座席邮箱
	StaffEmail *string `json:"StaffEmail,omitnil,omitempty" name:"StaffEmail"`

	// 绑定技能组列表
	SkillGroupList []*int64 `json:"SkillGroupList,omitnil,omitempty" name:"SkillGroupList"`

	// 绑定技能组列表(必填)
	StaffSkillGroupList []*StaffSkillGroupList `json:"StaffSkillGroupList,omitnil,omitempty" name:"StaffSkillGroupList"`
}

func (r *BindStaffSkillGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindStaffSkillGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StaffEmail")
	delete(f, "SkillGroupList")
	delete(f, "StaffSkillGroupList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindStaffSkillGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindStaffSkillGroupListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindStaffSkillGroupListResponse struct {
	*tchttp.BaseResponse
	Response *BindStaffSkillGroupListResponseParams `json:"Response"`
}

func (r *BindStaffSkillGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindStaffSkillGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CallInMetrics struct {
	// IVR驻留数量
	IvrCount *int64 `json:"IvrCount,omitnil,omitempty" name:"IvrCount"`

	// 排队中数量
	QueueCount *int64 `json:"QueueCount,omitnil,omitempty" name:"QueueCount"`

	// 振铃中数量
	RingCount *int64 `json:"RingCount,omitnil,omitempty" name:"RingCount"`

	// 接通中数量
	AcceptCount *int64 `json:"AcceptCount,omitnil,omitempty" name:"AcceptCount"`

	// 客服转接外线中数量
	TransferOuterCount *int64 `json:"TransferOuterCount,omitnil,omitempty" name:"TransferOuterCount"`

	// 最大排队时长
	MaxQueueDuration *int64 `json:"MaxQueueDuration,omitnil,omitempty" name:"MaxQueueDuration"`

	// 平均排队时长
	AvgQueueDuration *int64 `json:"AvgQueueDuration,omitnil,omitempty" name:"AvgQueueDuration"`

	// 最大振铃时长
	MaxRingDuration *int64 `json:"MaxRingDuration,omitnil,omitempty" name:"MaxRingDuration"`

	// 平均振铃时长
	AvgRingDuration *int64 `json:"AvgRingDuration,omitnil,omitempty" name:"AvgRingDuration"`

	// 最大接通时长
	MaxAcceptDuration *int64 `json:"MaxAcceptDuration,omitnil,omitempty" name:"MaxAcceptDuration"`

	// 平均接通时长
	AvgAcceptDuration *int64 `json:"AvgAcceptDuration,omitnil,omitempty" name:"AvgAcceptDuration"`
}

type CallInNumberMetrics struct {
	// 线路号码
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// 线路相关指标
	Metrics *CallInMetrics `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// 所属技能组相关指标
	SkillGroupMetrics []*CallInSkillGroupMetrics `json:"SkillGroupMetrics,omitnil,omitempty" name:"SkillGroupMetrics"`
}

type CallInSkillGroupMetrics struct {
	// 技能组ID
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// 数据指标
	Metrics *CallInMetrics `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// 技能组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type CalleeAttribute struct {
	// 被叫号码
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// 随路数据
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`

	// 参数
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`
}

type CarrierPrivilegeNumberApplicant struct {
	// 实例Id
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 申请单Id
	ApplicantId *uint64 `json:"ApplicantId,omitnil,omitempty" name:"ApplicantId"`

	// 主叫号码列表
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 被叫号码列表
	Callees []*string `json:"Callees,omitnil,omitempty" name:"Callees"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 审批状态:1 待审核、2 通过、3 拒绝
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 创建时间，unix时间戳(秒)
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间，unix时间戳(秒)
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type CompanyApplyInfo struct {
	// 申请人身份，0-公司法定代表人，1-经办人（受法定代表人委托）
	ApplicantType *int64 `json:"ApplicantType,omitnil,omitempty" name:"ApplicantType"`

	// 企业名称
	CompanyName *string `json:"CompanyName,omitnil,omitempty" name:"CompanyName"`

	// 统一社会信用代码
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// 营业执照扫描件(加盖公章)。(支持jpg、png、gif、jpeg格式的图片，每张图片应大于50K，不超过5MB，模板参见控制台:https://console.cloud.tencent.com/ccc/enterprise/update)
	BusinessIdPicUrl *string `json:"BusinessIdPicUrl,omitnil,omitempty" name:"BusinessIdPicUrl"`

	// 法定代表人名称
	CorporationName *string `json:"CorporationName,omitnil,omitempty" name:"CorporationName"`

	// 法定代表人身份证号码
	CorporationId *string `json:"CorporationId,omitnil,omitempty" name:"CorporationId"`

	// 法定代表人身份证正反面扫描件。(支持jpg、png、gif、jpeg格式的图片，每张图片应大于50K，不超过5MB，模板参见控制台:https://console.cloud.tencent.com/ccc/enterprise/update)
	CorporationIdPicUrl *string `json:"CorporationIdPicUrl,omitnil,omitempty" name:"CorporationIdPicUrl"`

	// 安全合规使用承诺书。(支持jpg、png、gif、jpeg格式的图片，每张图片应大于50K，不超过5MB，模板参见控制台:https://console.cloud.tencent.com/ccc/enterprise/update)
	NetworkCommitmentPicUrl *string `json:"NetworkCommitmentPicUrl,omitnil,omitempty" name:"NetworkCommitmentPicUrl"`

	// 是否与腾讯云账号的资质一致,0-不一致,1-一致
	IsEqualTencentCloud *int64 `json:"IsEqualTencentCloud,omitnil,omitempty" name:"IsEqualTencentCloud"`

	// 法定代表人手机号
	CorporationMobile *string `json:"CorporationMobile,omitnil,omitempty" name:"CorporationMobile"`

	// 法定代表人手机号码实名认证。(支持jpg、png、gif、jpeg格式的图片，每张图片应大于50K，不超过5MB，模板参见控制台:https://console.cloud.tencent.com/ccc/enterprise/update)
	CorporationMobilePicUrl *string `json:"CorporationMobilePicUrl,omitnil,omitempty" name:"CorporationMobilePicUrl"`

	// 通话话术。(支持doc、docx格式的文档不超过50MB，模板参见控制台:https://console.cloud.tencent.com/ccc/enterprise/update)
	UseDescribeFileUrl *string `json:"UseDescribeFileUrl,omitnil,omitempty" name:"UseDescribeFileUrl"`

	// 公司授权函。(支持jpg、png、gif、jpeg格式的图片，每张图片应大于50K，不超过5MB，模板参见控制台:https://console.cloud.tencent.com/ccc/enterprise/update)
	CompanyAuthLetterPicUrl *string `json:"CompanyAuthLetterPicUrl,omitnil,omitempty" name:"CompanyAuthLetterPicUrl"`

	// 电话受理单。(支持jpg、png、gif、jpeg格式的图片，每张图片应大于50K，不超过5MB，模板参见控制台:https://console.cloud.tencent.com/ccc/enterprise/update)
	AcceptPicUrl *string `json:"AcceptPicUrl,omitnil,omitempty" name:"AcceptPicUrl"`

	// 法定代表人手持身份证照，申请人类型为法定代表人时必填。(支持jpg、png、gif、jpeg格式的图片，每张图片应大于50K，不超过5MB，模板参见控制台:https://console.cloud.tencent.com/ccc/enterprise/update)
	CorporationHoldingOnIdPicUrl *string `json:"CorporationHoldingOnIdPicUrl,omitnil,omitempty" name:"CorporationHoldingOnIdPicUrl"`

	// 经办人名称，申请人类型为经办人时必填。
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 经办人证件号码，申请人类型为经办人时必填。
	OperatorId *string `json:"OperatorId,omitnil,omitempty" name:"OperatorId"`

	// 经办人身份证正反面扫描件，申请人类型为经办人时必填。(支持jpg、png、gif、jpeg格式的图片，每张图片应大于50K，不超过5MB，模板参见控制台:https://console.cloud.tencent.com/ccc/enterprise/update)
	OperatorIdPicUrl *string `json:"OperatorIdPicUrl,omitnil,omitempty" name:"OperatorIdPicUrl"`

	// 经办人手持身份证照，申请人类型为经办人时必填。(支持jpg、png、gif、jpeg格式的图片，每张图片应大于50K，不超过5MB，模板参见控制台:https://console.cloud.tencent.com/ccc/enterprise/update)
	OperatorHoldingOnIdPicUrl *string `json:"OperatorHoldingOnIdPicUrl,omitnil,omitempty" name:"OperatorHoldingOnIdPicUrl"`

	// 委托授权书，申请人类型为经办人时必填。(支持jpg、png、gif、jpeg格式的图片，每张图片应大于50K，不超过5MB，模板参见控制台:https://console.cloud.tencent.com/ccc/enterprise/update)
	CommissionPicUrl *string `json:"CommissionPicUrl,omitnil,omitempty" name:"CommissionPicUrl"`

	// 经办人手机号，申请人类型为经办人时必填。
	OperatorMobile *string `json:"OperatorMobile,omitnil,omitempty" name:"OperatorMobile"`

	// 经办人邮箱，申请人类型为经办人时必填。
	OperatorEmail *string `json:"OperatorEmail,omitnil,omitempty" name:"OperatorEmail"`

	// 经办人手机号码实名认证，申请人类型为经办人时必填。(支持jpg、png、gif、jpeg格式的图片，每张图片应大于50K，不超过5MB，模板参见控制台:https://console.cloud.tencent.com/ccc/enterprise/update)
	OperatorMobilePicUrl *string `json:"OperatorMobilePicUrl,omitnil,omitempty" name:"OperatorMobilePicUrl"`
}

type CompanyStateInfo struct {
	// 申请单ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公司名称
	CompanyName *string `json:"CompanyName,omitnil,omitempty" name:"CompanyName"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 审核时间
	CheckTime *int64 `json:"CheckTime,omitnil,omitempty" name:"CheckTime"`

	// 审核备注
	CheckMsg *string `json:"CheckMsg,omitnil,omitempty" name:"CheckMsg"`

	// 审核状态，1-待审核，2-审核通过，3-驳回
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 公司统一社会信用代码
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// 修改时间
	ModifyTime *int64 `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 合同编号
	ContractNo *string `json:"ContractNo,omitnil,omitempty" name:"ContractNo"`
}

// Predefined struct for user
type CreateAIAgentCallRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// AI智能体ID
	AIAgentId *uint64 `json:"AIAgentId,omitnil,omitempty" name:"AIAgentId"`

	// 被叫号码
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// 主叫号码列表
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 提示词变量
	PromptVariables []*Variable `json:"PromptVariables,omitnil,omitempty" name:"PromptVariables"`
}

type CreateAIAgentCallRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// AI智能体ID
	AIAgentId *uint64 `json:"AIAgentId,omitnil,omitempty" name:"AIAgentId"`

	// 被叫号码
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// 主叫号码列表
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 提示词变量
	PromptVariables []*Variable `json:"PromptVariables,omitnil,omitempty" name:"PromptVariables"`
}

func (r *CreateAIAgentCallRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIAgentCallRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "AIAgentId")
	delete(f, "Callee")
	delete(f, "Callers")
	delete(f, "PromptVariables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAIAgentCallRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIAgentCallResponseParams struct {
	// 新创建的会话 ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAIAgentCallResponse struct {
	*tchttp.BaseResponse
	Response *CreateAIAgentCallResponseParams `json:"Response"`
}

func (r *CreateAIAgentCallResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIAgentCallResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAICallRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 被叫号码
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// 用于设定AI人设、说话规则、任务等的全局提示词。示例：## 人设您是人民医院友善、和蔼的随访医生李医生，正在给患者小明的家长打电话，原因是医院要求小明2024-08-08回院复查手术恢复情况，但小明没有来。您需要按照任务流程对小明家长进行电话随访调查。## 要求简洁回复：使用简练语言，每次最多询问一个问题，不要在一个回复中询问多个问题。富有变化：尽量使表达富有变化，表达机械重复。自然亲切：使用日常语言，尽量显得专业并亲切。提到时间时使用口语表述，如下周三、6月18日。积极主动：尝试引导对话，每个回复通常以问题或下一步建议来结尾。询问清楚：如果对方部分回答了您的问题，或者回答很模糊，请通过追问来确保回答的完整明确。遵循任务：当对方的回答偏离了您的任务时，及时引导对方回到任务中。不要从头开始重复，从偏离的地方继续询问。诚实可靠：对于客户的提问，如果不确定请务必不要编造，礼貌告知对方不清楚。不要捏造患者未提及的症状史、用药史、治疗史。其他注意点：避免提到病情恶化、恢复不理想或疾病名称等使用会使患者感到紧张的表述。不要问患者已经直接或间接回答过的问题，例如患者已经说没有不适症状，那就不要再问手术部位是否有红肿疼痛症状的问题。##任务： 1.自我介绍您是人民医院负责随访的李医生，并说明致电的目的。2.询问被叫方是否是小明家长。 - 如果不是小明家长，请礼貌表达歉意，并使用 call_end 挂断电话。- 如果小明家长没空，请礼貌告诉对方稍后会重新致电，并使用 end_call 挂断电话。3.询问小明出院后水肿情况如何，较出院时是否有变化。- 如果水肿变严重，直接跳转步骤7。4.询问出院后是否给小朋友量过体温，是否出现过发烧情况。- 如果没有量过体温，请礼貌告诉家长出院后三个月内需要每天观察体温。- 如果出现过发烧，请直接跳转步骤7。5.询问出院后是否给小朋友按时服药。- 如果没有按时服药，请友善提醒家长严格按医嘱服用药物，避免影响手术效果。6.询问小朋友在饮食上是否做到低盐低脂，适量吃优质蛋白如鸡蛋、牛奶、瘦肉等。- 如果没有做到，请友善提醒家长低盐低脂和优质蛋白有助小朋友尽快恢复。7.告知家长医生要求6月18日回院复查，但没看到有相关复诊记录。提醒家长尽快前往医院体检复查血化验、尿常规。8.询问家长是否有问题需要咨询，如果没有请礼貌道别并用call_end挂断电话。
	SystemPrompt *string `json:"SystemPrompt,omitnil,omitempty" name:"SystemPrompt"`

	// 模型接口协议类型，目前兼容三种协议类型：
	// 
	// - OpenAI协议(包括GPT、混元、DeepSeek等)："openai"
	// - Azure协议："azure"
	// - Minimax协议："minimax"
	LLMType *string `json:"LLMType,omitnil,omitempty" name:"LLMType"`

	// 模型名称，如
	// 
	// - OpenAI协议
	// "gpt-4o-mini","gpt-4o"，"hunyuan-standard", "hunyuan-turbo"，"deepseek-chat"；
	// 
	// - Azure协议
	// "gpt-4o-mini", "gpt-4o"；
	// 
	// - Minmax协议
	// "deepseek-chat".
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 模型API密钥，获取鉴权信息方式请参见各模型官网
	// 
	// - OpenAI协议：[GPT](https://help.openai.com/en/articles/4936850-where-do-i-find-my-openai-api-key)，[混元](https://cloud.tencent.com/document/product/1729/111008)，[DeepSeek](https://api-docs.deepseek.com/zh-cn/)；
	// 
	// - Azure协议：[Azure GPT](https://learn.microsoft.com/en-us/azure/ai-services/openai/chatgpt-quickstart?tabs=command-line%2Ctypescript%2Cpython-new&pivots=programming-language-studio#key-settings)；
	// 
	// - Minimax：[Minimax](https://platform.minimaxi.com/document/Fast%20access?key=66701cf51d57f38758d581b2)
	APIKey *string `json:"APIKey,omitnil,omitempty" name:"APIKey"`

	// 模型接口地址
	// 
	// - OpenAI协议
	// GPT："https://api.openai.com/v1/"
	// 混元："https://api.hunyuan.cloud.tencent.com/v1"
	// Deepseek："https://api.deepseek.com/v1"
	// 
	// - Azure协议
	//  "https://{your-resource-name}.openai.azure.com?api-version={api-version}"
	// 
	// - Minimax协议
	// "https://api.minimax.chat/v1"
	APIUrl *string `json:"APIUrl,omitnil,omitempty" name:"APIUrl"`

	// 默认提供以下音色参数值可选择，如需自定义音色VoiceType请留空并在参数CustomTTSConfig中配置
	// 
	// 汉语：
	// ZhiMei：智美，客服女声
	// ZhiXi： 智希 通用女声
	// ZhiQi：智琪 客服女声
	// ZhiTian：智甜 女童声
	// AiXiaoJing：爱小静 对话女声
	// 
	// 英语:
	// WeRose：英文女声
	// Monika：英文女声
	// 
	// 日语：
	// Nanami
	// 
	// 韩语：
	// SunHi
	// 
	// 印度尼西亚语(印度尼西亚)：
	// Gadis
	// 
	// 马来语（马来西亚）:
	// Yasmin
	// 
	//  泰米尔语（马来西亚）:
	// Kani
	// 
	// 泰语（泰国）:
	// Achara
	// 
	// 越南语(越南):
	// HoaiMy
	// 
	VoiceType *string `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`

	// 主叫号码列表
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 用于设定AI座席欢迎语。
	WelcomeMessage *string `json:"WelcomeMessage,omitnil,omitempty" name:"WelcomeMessage"`

	// 0：使用welcomeMessage(为空时，被叫先说话；不为空时，机器人先说话)
	// 1:   使用ai根据prompt自动生成welcomeMessage并先说话
	WelcomeType *int64 `json:"WelcomeType,omitnil,omitempty" name:"WelcomeType"`

	// 0: 默认可打断， 1：高优先不可打断
	WelcomeMessagePriority *int64 `json:"WelcomeMessagePriority,omitnil,omitempty" name:"WelcomeMessagePriority"`

	// 最大等待时长(毫秒)，默认60秒，超过这个时间用户没说话，自动挂断
	MaxDuration *int64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`

	// 语音识别支持的语言, 默认是"zh" 中文,
	// 填写数组,最长4个语言，第一个语言为主要识别语言，后面为可选语言，
	// 注意:主要语言为中国方言时，可选语言无效
	// 目前全量支持的语言如下，等号左面是语言英文名，右面是Language字段需要填写的值，该值遵循ISO639：
	// 1. Chinese = "zh" # 中文
	// 2. Chinese_TW = "zh-TW" # 中国台湾
	// 3. Chinese_DIALECT = "zh-dialect" # 中国方言
	// 4. English = "en" # 英语
	// 5. Vietnamese = "vi" # 越南语
	// 6. Japanese = "ja" # 日语
	// 7. Korean = "ko" # 汉语
	// 8. Indonesia = "id" # 印度尼西亚语
	// 9. Thai = "th" # 泰语
	// 10. Portuguese = "pt" # 葡萄牙语
	// 11. Turkish = "tr" # 土耳其语
	// 12. Arabic = "ar" # 阿拉伯语
	// 13. Spanish = "es" # 西班牙语
	// 14. Hindi = "hi" # 印地语
	// 15. French = "fr" # 法语
	// 16. Malay = "ms" # 马来语
	// 17. Filipino = "fil" # 菲律宾语
	// 18. German = "de" # 德语
	// 19. Italian = "it" # 意大利语
	// 20. Russian = "ru" # 俄语
	Languages []*string `json:"Languages,omitnil,omitempty" name:"Languages"`

	// 打断AI说话模式，默认为0，0表示自动打断，1表示不打断。
	InterruptMode *int64 `json:"InterruptMode,omitnil,omitempty" name:"InterruptMode"`

	// InterruptMode为0时使用，单位为毫秒，默认为500ms。表示服务端检测到持续InterruptSpeechDuration毫秒的人声则进行打断。
	InterruptSpeechDuration *int64 `json:"InterruptSpeechDuration,omitnil,omitempty" name:"InterruptSpeechDuration"`

	// 模型是否支持(或者开启)call_end function calling
	EndFunctionEnable *bool `json:"EndFunctionEnable,omitnil,omitempty" name:"EndFunctionEnable"`

	// EndFunctionEnable为true时生效；call_end function calling的desc，默认为 "End the call when user has to leave (like says bye) or you are instructed to do so."
	EndFunctionDesc *string `json:"EndFunctionDesc,omitnil,omitempty" name:"EndFunctionDesc"`

	// 模型是否支持(或者开启)transfer_to_human function calling
	TransferFunctionEnable *bool `json:"TransferFunctionEnable,omitnil,omitempty" name:"TransferFunctionEnable"`

	// TransferFunctionEnable为true的时候生效: 转人工配置
	TransferItems []*AITransferItem `json:"TransferItems,omitnil,omitempty" name:"TransferItems"`

	// 用户多久没说话提示时长,最小10秒,默认10秒
	NotifyDuration *int64 `json:"NotifyDuration,omitnil,omitempty" name:"NotifyDuration"`

	// 用户NotifyDuration没说话，AI提示的语句，默认是"抱歉，我没听清。您可以重复下吗？"
	NotifyMessage *string `json:"NotifyMessage,omitnil,omitempty" name:"NotifyMessage"`

	// 最大触发AI提示音次数，默认为不限制
	NotifyMaxCount *uint64 `json:"NotifyMaxCount,omitnil,omitempty" name:"NotifyMaxCount"`

	// <p>和VoiceType字段需要选填一个，这里是使用自己自定义的TTS，VoiceType是系统内置的一些音色</p>
	// <ul>
	// <li>Tencent TTS<br>
	// 配置请参考<a href="https://cloud.tencent.com/document/product/1073/92668#55924b56-1a73-4663-a7a1-a8dd82d6e823" target="_blank">腾讯云TTS文档链接</a></li>
	// </ul>
	// <div><div class="v-md-pre-wrapper copy-code-mode v-md-pre-wrapper- extra-class"><pre class="v-md-prism-"><code>{ 
	//        &quot;TTSType&quot;: &quot;tencent&quot;, // String TTS类型, 目前支持&quot;tencent&quot; 和 “minixmax”， 其他的厂商支持中
	//        &quot;AppId&quot;: &quot;您的应用ID&quot;, // String 必填
	//        &quot;SecretId&quot;: &quot;您的密钥ID&quot;, // String 必填
	//        &quot;SecretKey&quot;:  &quot;您的密钥Key&quot;, // String 必填
	//        &quot;VoiceType&quot;: 101001, // Integer  必填，音色 ID，包括标准音色与精品音色，精品音色拟真度更高，价格不同于标准音色，请参见语音合成计费概述。完整的音色 ID 列表请参见语音合成音色列表。
	//        &quot;Speed&quot;: 1.25, // Integer 非必填，语速，范围：[-2，6]，分别对应不同语速： -2: 代表0.6倍 -1: 代表0.8倍 0: 代表1.0倍（默认） 1: 代表1.2倍 2: 代表1.5倍  6: 代表2.5倍  如果需要更细化的语速，可以保留小数点后 2 位，例如0.5/1.25/2.81等。 参数值与实际语速转换，可参考 语速转换
	//        &quot;Volume&quot;: 5, // Integer 非必填，音量大小，范围：[0，10]，分别对应11个等级的音量，默认值为0，代表正常音量。
	//        &quot;PrimaryLanguage&quot;: 1, // Integer 可选 主要语言 1-中文（默认） 2-英文 3-日文
	//        &quot;FastVoiceType&quot;: &quot;xxxx&quot;   //  可选参数， 快速声音复刻的参数 
	//   }
	// </code></pre>
	// 
	//   </div></div><ul>
	// <li>Minimax TTS<br>
	// 配置请参考<a href="https://platform.minimaxi.com/document/T2A%20V2?key=66719005a427f0c8a5701643" target="_blank">Minimax TTS文档链接</a>。注意Minimax TTS存在频率限制，超频可能会导致回答卡顿，<a href="https://platform.minimaxi.com/document/Rate%20limits?key=66b19417290299a26b234572" target="_blank">Minimax TTS频率限制相关文档链接</a>。</li>
	// </ul>
	// <div><div class="v-md-pre-wrapper copy-code-mode v-md-pre-wrapper- extra-class"><pre class="v-md-prism-"><code>{
	//         &quot;TTSType&quot;: &quot;minimax&quot;,  // String TTS类型, 
	//         &quot;Model&quot;: &quot;speech-01-turbo&quot;,
	//         &quot;APIUrl&quot;: &quot;https://api.minimax.chat/v1/t2a_v2&quot;,
	//         &quot;APIKey&quot;: &quot;eyxxxx&quot;,
	//         &quot;GroupId&quot;: &quot;181000000000000&quot;,
	//         &quot;VoiceType&quot;:&quot;female-tianmei-jingpin&quot;,
	//         &quot;Speed&quot;: 1.2
	// }
	// </code></pre>
	// </div></div><ul>
	// <li>火山 TTS</li>
	// </ul>
	// <p>配置音色类型参考<a href="https://www.volcengine.com/docs/6561/162929" target="_blank">火山TTS文档链接</a><br>
	// 语音合成音色列表–语音技术-火山引擎<br>
	// 大模型语音合成音色列表–语音技术-火山引擎</p>
	// <div><div class="v-md-pre-wrapper copy-code-mode v-md-pre-wrapper- extra-class"><pre class="v-md-prism-"><code>{
	//     &quot;TTSType&quot;: &quot;volcengine&quot;,  // 必填：String TTS类型
	//     &quot;AppId&quot; : &quot;xxxxxxxx&quot;,   // 必填：String 火山引擎分配的Appid
	//     &quot;Token&quot; : &quot;TY9d4sQXHxxxxxxx&quot;, // 必填： String类型 火山引擎的访问token
	//     &quot;Speed&quot; : 1.0,            // 可选参数 语速，默认为1.0
	//     &quot;Volume&quot;: 1.0,            // 可选参数， 音量大小， 默认为1.0
	//     &quot;Cluster&quot; : &quot;volcano_tts&quot;, // 可选参数，业务集群, 默认是 volcano_tts
	//     &quot;VoiceType&quot; : &quot;zh_male_aojiaobazong_moon_bigtts&quot;   // 音色类型， 默认为大模型语音合成的音色。 如果使用普通语音合成，则需要填写对应的音色类型。 音色类型填写错误会导致没有声音。
	// }
	// </code></pre>
	// 
	// </div></div><ul>
	// <li>Azure TTS<br>
	// 配置请参考<a href="https://docs.azure.cn/zh-cn/ai-services/speech-service/speech-synthesis-markup-voice" target="_blank">AzureTTS文档链接</a></li>
	// </ul>
	// <div><div class="v-md-pre-wrapper copy-code-mode v-md-pre-wrapper- extra-class"><pre class="v-md-prism-"><code>{
	//     &quot;TTSType&quot;: &quot;azure&quot;, // 必填：String TTS类型
	//     &quot;SubscriptionKey&quot;: &quot;xxxxxxxx&quot;, // 必填：String 订阅的Key
	//     &quot;Region&quot;: &quot;chinanorth3&quot;,  // 必填：String 订阅的地区
	//     &quot;VoiceName&quot;: &quot;zh-CN-XiaoxiaoNeural&quot;, // 必填：String 音色名必填
	//     &quot;Language&quot;: &quot;zh-CN&quot;, // 必填：String 合成的语言  
	//     &quot;Rate&quot;: 1 // 选填：float 语速  0.5～2 默认为 1
	// }
	// </code></pre>
	// 
	// </div></div><ul>
	// <li>自定义</li>
	// </ul>
	// <p>TTS<br>
	// 具体协议规范请参考<a href="https://doc.weixin.qq.com/doc/w3_ANQAiAbdAFwHILbJBmtSqSbV1WZ3L?scode=AJEAIQdfAAo5a1xajYANQAiAbdAFw" target="_blank">腾讯文档</a></p>
	// <div><div class="v-md-pre-wrapper copy-code-mode v-md-pre-wrapper- extra-class"><pre class="v-md-prism-"><code>{
	//   &quot;TTSType&quot;: &quot;custom&quot;, // String 必填
	//   &quot;APIKey&quot;: &quot;ApiKey&quot;, // String 必填 用来鉴权
	//   &quot;APIUrl&quot;: &quot;http://0.0.0.0:8080/stream-audio&quot; // String，必填，TTS API URL
	//   &quot;AudioFormat&quot;: &quot;wav&quot;, // String, 非必填，期望输出的音频格式，如mp3， ogg_opus，pcm，wav，默认为 wav，目前只支持pcm和wav，
	//   &quot;SampleRate&quot;: 16000,  // Integer，非必填，音频采样率，默认为16000(16k)，推荐值为16000
	//   &quot;AudioChannel&quot;: 1,    // Integer，非必填，音频通道数，取值：1 或 2  默认为1  
	// }
	// </code></pre>
	// 
	// </div></div>
	CustomTTSConfig *string `json:"CustomTTSConfig,omitnil,omitempty" name:"CustomTTSConfig"`

	// 提示词变量
	PromptVariables []*Variable `json:"PromptVariables,omitnil,omitempty" name:"PromptVariables"`

	// 语音识别vad的时间，范围为240-2000，默认为1000，单位为ms。更小的值会让语音识别分句更快。
	VadSilenceTime *int64 `json:"VadSilenceTime,omitnil,omitempty" name:"VadSilenceTime"`

	// 通话内容提取配置
	ExtractConfig []*AICallExtractConfigElement `json:"ExtractConfig,omitnil,omitempty" name:"ExtractConfig"`
}

type CreateAICallRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 被叫号码
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// 用于设定AI人设、说话规则、任务等的全局提示词。示例：## 人设您是人民医院友善、和蔼的随访医生李医生，正在给患者小明的家长打电话，原因是医院要求小明2024-08-08回院复查手术恢复情况，但小明没有来。您需要按照任务流程对小明家长进行电话随访调查。## 要求简洁回复：使用简练语言，每次最多询问一个问题，不要在一个回复中询问多个问题。富有变化：尽量使表达富有变化，表达机械重复。自然亲切：使用日常语言，尽量显得专业并亲切。提到时间时使用口语表述，如下周三、6月18日。积极主动：尝试引导对话，每个回复通常以问题或下一步建议来结尾。询问清楚：如果对方部分回答了您的问题，或者回答很模糊，请通过追问来确保回答的完整明确。遵循任务：当对方的回答偏离了您的任务时，及时引导对方回到任务中。不要从头开始重复，从偏离的地方继续询问。诚实可靠：对于客户的提问，如果不确定请务必不要编造，礼貌告知对方不清楚。不要捏造患者未提及的症状史、用药史、治疗史。其他注意点：避免提到病情恶化、恢复不理想或疾病名称等使用会使患者感到紧张的表述。不要问患者已经直接或间接回答过的问题，例如患者已经说没有不适症状，那就不要再问手术部位是否有红肿疼痛症状的问题。##任务： 1.自我介绍您是人民医院负责随访的李医生，并说明致电的目的。2.询问被叫方是否是小明家长。 - 如果不是小明家长，请礼貌表达歉意，并使用 call_end 挂断电话。- 如果小明家长没空，请礼貌告诉对方稍后会重新致电，并使用 end_call 挂断电话。3.询问小明出院后水肿情况如何，较出院时是否有变化。- 如果水肿变严重，直接跳转步骤7。4.询问出院后是否给小朋友量过体温，是否出现过发烧情况。- 如果没有量过体温，请礼貌告诉家长出院后三个月内需要每天观察体温。- 如果出现过发烧，请直接跳转步骤7。5.询问出院后是否给小朋友按时服药。- 如果没有按时服药，请友善提醒家长严格按医嘱服用药物，避免影响手术效果。6.询问小朋友在饮食上是否做到低盐低脂，适量吃优质蛋白如鸡蛋、牛奶、瘦肉等。- 如果没有做到，请友善提醒家长低盐低脂和优质蛋白有助小朋友尽快恢复。7.告知家长医生要求6月18日回院复查，但没看到有相关复诊记录。提醒家长尽快前往医院体检复查血化验、尿常规。8.询问家长是否有问题需要咨询，如果没有请礼貌道别并用call_end挂断电话。
	SystemPrompt *string `json:"SystemPrompt,omitnil,omitempty" name:"SystemPrompt"`

	// 模型接口协议类型，目前兼容三种协议类型：
	// 
	// - OpenAI协议(包括GPT、混元、DeepSeek等)："openai"
	// - Azure协议："azure"
	// - Minimax协议："minimax"
	LLMType *string `json:"LLMType,omitnil,omitempty" name:"LLMType"`

	// 模型名称，如
	// 
	// - OpenAI协议
	// "gpt-4o-mini","gpt-4o"，"hunyuan-standard", "hunyuan-turbo"，"deepseek-chat"；
	// 
	// - Azure协议
	// "gpt-4o-mini", "gpt-4o"；
	// 
	// - Minmax协议
	// "deepseek-chat".
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 模型API密钥，获取鉴权信息方式请参见各模型官网
	// 
	// - OpenAI协议：[GPT](https://help.openai.com/en/articles/4936850-where-do-i-find-my-openai-api-key)，[混元](https://cloud.tencent.com/document/product/1729/111008)，[DeepSeek](https://api-docs.deepseek.com/zh-cn/)；
	// 
	// - Azure协议：[Azure GPT](https://learn.microsoft.com/en-us/azure/ai-services/openai/chatgpt-quickstart?tabs=command-line%2Ctypescript%2Cpython-new&pivots=programming-language-studio#key-settings)；
	// 
	// - Minimax：[Minimax](https://platform.minimaxi.com/document/Fast%20access?key=66701cf51d57f38758d581b2)
	APIKey *string `json:"APIKey,omitnil,omitempty" name:"APIKey"`

	// 模型接口地址
	// 
	// - OpenAI协议
	// GPT："https://api.openai.com/v1/"
	// 混元："https://api.hunyuan.cloud.tencent.com/v1"
	// Deepseek："https://api.deepseek.com/v1"
	// 
	// - Azure协议
	//  "https://{your-resource-name}.openai.azure.com?api-version={api-version}"
	// 
	// - Minimax协议
	// "https://api.minimax.chat/v1"
	APIUrl *string `json:"APIUrl,omitnil,omitempty" name:"APIUrl"`

	// 默认提供以下音色参数值可选择，如需自定义音色VoiceType请留空并在参数CustomTTSConfig中配置
	// 
	// 汉语：
	// ZhiMei：智美，客服女声
	// ZhiXi： 智希 通用女声
	// ZhiQi：智琪 客服女声
	// ZhiTian：智甜 女童声
	// AiXiaoJing：爱小静 对话女声
	// 
	// 英语:
	// WeRose：英文女声
	// Monika：英文女声
	// 
	// 日语：
	// Nanami
	// 
	// 韩语：
	// SunHi
	// 
	// 印度尼西亚语(印度尼西亚)：
	// Gadis
	// 
	// 马来语（马来西亚）:
	// Yasmin
	// 
	//  泰米尔语（马来西亚）:
	// Kani
	// 
	// 泰语（泰国）:
	// Achara
	// 
	// 越南语(越南):
	// HoaiMy
	// 
	VoiceType *string `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`

	// 主叫号码列表
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 用于设定AI座席欢迎语。
	WelcomeMessage *string `json:"WelcomeMessage,omitnil,omitempty" name:"WelcomeMessage"`

	// 0：使用welcomeMessage(为空时，被叫先说话；不为空时，机器人先说话)
	// 1:   使用ai根据prompt自动生成welcomeMessage并先说话
	WelcomeType *int64 `json:"WelcomeType,omitnil,omitempty" name:"WelcomeType"`

	// 0: 默认可打断， 1：高优先不可打断
	WelcomeMessagePriority *int64 `json:"WelcomeMessagePriority,omitnil,omitempty" name:"WelcomeMessagePriority"`

	// 最大等待时长(毫秒)，默认60秒，超过这个时间用户没说话，自动挂断
	MaxDuration *int64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`

	// 语音识别支持的语言, 默认是"zh" 中文,
	// 填写数组,最长4个语言，第一个语言为主要识别语言，后面为可选语言，
	// 注意:主要语言为中国方言时，可选语言无效
	// 目前全量支持的语言如下，等号左面是语言英文名，右面是Language字段需要填写的值，该值遵循ISO639：
	// 1. Chinese = "zh" # 中文
	// 2. Chinese_TW = "zh-TW" # 中国台湾
	// 3. Chinese_DIALECT = "zh-dialect" # 中国方言
	// 4. English = "en" # 英语
	// 5. Vietnamese = "vi" # 越南语
	// 6. Japanese = "ja" # 日语
	// 7. Korean = "ko" # 汉语
	// 8. Indonesia = "id" # 印度尼西亚语
	// 9. Thai = "th" # 泰语
	// 10. Portuguese = "pt" # 葡萄牙语
	// 11. Turkish = "tr" # 土耳其语
	// 12. Arabic = "ar" # 阿拉伯语
	// 13. Spanish = "es" # 西班牙语
	// 14. Hindi = "hi" # 印地语
	// 15. French = "fr" # 法语
	// 16. Malay = "ms" # 马来语
	// 17. Filipino = "fil" # 菲律宾语
	// 18. German = "de" # 德语
	// 19. Italian = "it" # 意大利语
	// 20. Russian = "ru" # 俄语
	Languages []*string `json:"Languages,omitnil,omitempty" name:"Languages"`

	// 打断AI说话模式，默认为0，0表示自动打断，1表示不打断。
	InterruptMode *int64 `json:"InterruptMode,omitnil,omitempty" name:"InterruptMode"`

	// InterruptMode为0时使用，单位为毫秒，默认为500ms。表示服务端检测到持续InterruptSpeechDuration毫秒的人声则进行打断。
	InterruptSpeechDuration *int64 `json:"InterruptSpeechDuration,omitnil,omitempty" name:"InterruptSpeechDuration"`

	// 模型是否支持(或者开启)call_end function calling
	EndFunctionEnable *bool `json:"EndFunctionEnable,omitnil,omitempty" name:"EndFunctionEnable"`

	// EndFunctionEnable为true时生效；call_end function calling的desc，默认为 "End the call when user has to leave (like says bye) or you are instructed to do so."
	EndFunctionDesc *string `json:"EndFunctionDesc,omitnil,omitempty" name:"EndFunctionDesc"`

	// 模型是否支持(或者开启)transfer_to_human function calling
	TransferFunctionEnable *bool `json:"TransferFunctionEnable,omitnil,omitempty" name:"TransferFunctionEnable"`

	// TransferFunctionEnable为true的时候生效: 转人工配置
	TransferItems []*AITransferItem `json:"TransferItems,omitnil,omitempty" name:"TransferItems"`

	// 用户多久没说话提示时长,最小10秒,默认10秒
	NotifyDuration *int64 `json:"NotifyDuration,omitnil,omitempty" name:"NotifyDuration"`

	// 用户NotifyDuration没说话，AI提示的语句，默认是"抱歉，我没听清。您可以重复下吗？"
	NotifyMessage *string `json:"NotifyMessage,omitnil,omitempty" name:"NotifyMessage"`

	// 最大触发AI提示音次数，默认为不限制
	NotifyMaxCount *uint64 `json:"NotifyMaxCount,omitnil,omitempty" name:"NotifyMaxCount"`

	// <p>和VoiceType字段需要选填一个，这里是使用自己自定义的TTS，VoiceType是系统内置的一些音色</p>
	// <ul>
	// <li>Tencent TTS<br>
	// 配置请参考<a href="https://cloud.tencent.com/document/product/1073/92668#55924b56-1a73-4663-a7a1-a8dd82d6e823" target="_blank">腾讯云TTS文档链接</a></li>
	// </ul>
	// <div><div class="v-md-pre-wrapper copy-code-mode v-md-pre-wrapper- extra-class"><pre class="v-md-prism-"><code>{ 
	//        &quot;TTSType&quot;: &quot;tencent&quot;, // String TTS类型, 目前支持&quot;tencent&quot; 和 “minixmax”， 其他的厂商支持中
	//        &quot;AppId&quot;: &quot;您的应用ID&quot;, // String 必填
	//        &quot;SecretId&quot;: &quot;您的密钥ID&quot;, // String 必填
	//        &quot;SecretKey&quot;:  &quot;您的密钥Key&quot;, // String 必填
	//        &quot;VoiceType&quot;: 101001, // Integer  必填，音色 ID，包括标准音色与精品音色，精品音色拟真度更高，价格不同于标准音色，请参见语音合成计费概述。完整的音色 ID 列表请参见语音合成音色列表。
	//        &quot;Speed&quot;: 1.25, // Integer 非必填，语速，范围：[-2，6]，分别对应不同语速： -2: 代表0.6倍 -1: 代表0.8倍 0: 代表1.0倍（默认） 1: 代表1.2倍 2: 代表1.5倍  6: 代表2.5倍  如果需要更细化的语速，可以保留小数点后 2 位，例如0.5/1.25/2.81等。 参数值与实际语速转换，可参考 语速转换
	//        &quot;Volume&quot;: 5, // Integer 非必填，音量大小，范围：[0，10]，分别对应11个等级的音量，默认值为0，代表正常音量。
	//        &quot;PrimaryLanguage&quot;: 1, // Integer 可选 主要语言 1-中文（默认） 2-英文 3-日文
	//        &quot;FastVoiceType&quot;: &quot;xxxx&quot;   //  可选参数， 快速声音复刻的参数 
	//   }
	// </code></pre>
	// 
	//   </div></div><ul>
	// <li>Minimax TTS<br>
	// 配置请参考<a href="https://platform.minimaxi.com/document/T2A%20V2?key=66719005a427f0c8a5701643" target="_blank">Minimax TTS文档链接</a>。注意Minimax TTS存在频率限制，超频可能会导致回答卡顿，<a href="https://platform.minimaxi.com/document/Rate%20limits?key=66b19417290299a26b234572" target="_blank">Minimax TTS频率限制相关文档链接</a>。</li>
	// </ul>
	// <div><div class="v-md-pre-wrapper copy-code-mode v-md-pre-wrapper- extra-class"><pre class="v-md-prism-"><code>{
	//         &quot;TTSType&quot;: &quot;minimax&quot;,  // String TTS类型, 
	//         &quot;Model&quot;: &quot;speech-01-turbo&quot;,
	//         &quot;APIUrl&quot;: &quot;https://api.minimax.chat/v1/t2a_v2&quot;,
	//         &quot;APIKey&quot;: &quot;eyxxxx&quot;,
	//         &quot;GroupId&quot;: &quot;181000000000000&quot;,
	//         &quot;VoiceType&quot;:&quot;female-tianmei-jingpin&quot;,
	//         &quot;Speed&quot;: 1.2
	// }
	// </code></pre>
	// </div></div><ul>
	// <li>火山 TTS</li>
	// </ul>
	// <p>配置音色类型参考<a href="https://www.volcengine.com/docs/6561/162929" target="_blank">火山TTS文档链接</a><br>
	// 语音合成音色列表–语音技术-火山引擎<br>
	// 大模型语音合成音色列表–语音技术-火山引擎</p>
	// <div><div class="v-md-pre-wrapper copy-code-mode v-md-pre-wrapper- extra-class"><pre class="v-md-prism-"><code>{
	//     &quot;TTSType&quot;: &quot;volcengine&quot;,  // 必填：String TTS类型
	//     &quot;AppId&quot; : &quot;xxxxxxxx&quot;,   // 必填：String 火山引擎分配的Appid
	//     &quot;Token&quot; : &quot;TY9d4sQXHxxxxxxx&quot;, // 必填： String类型 火山引擎的访问token
	//     &quot;Speed&quot; : 1.0,            // 可选参数 语速，默认为1.0
	//     &quot;Volume&quot;: 1.0,            // 可选参数， 音量大小， 默认为1.0
	//     &quot;Cluster&quot; : &quot;volcano_tts&quot;, // 可选参数，业务集群, 默认是 volcano_tts
	//     &quot;VoiceType&quot; : &quot;zh_male_aojiaobazong_moon_bigtts&quot;   // 音色类型， 默认为大模型语音合成的音色。 如果使用普通语音合成，则需要填写对应的音色类型。 音色类型填写错误会导致没有声音。
	// }
	// </code></pre>
	// 
	// </div></div><ul>
	// <li>Azure TTS<br>
	// 配置请参考<a href="https://docs.azure.cn/zh-cn/ai-services/speech-service/speech-synthesis-markup-voice" target="_blank">AzureTTS文档链接</a></li>
	// </ul>
	// <div><div class="v-md-pre-wrapper copy-code-mode v-md-pre-wrapper- extra-class"><pre class="v-md-prism-"><code>{
	//     &quot;TTSType&quot;: &quot;azure&quot;, // 必填：String TTS类型
	//     &quot;SubscriptionKey&quot;: &quot;xxxxxxxx&quot;, // 必填：String 订阅的Key
	//     &quot;Region&quot;: &quot;chinanorth3&quot;,  // 必填：String 订阅的地区
	//     &quot;VoiceName&quot;: &quot;zh-CN-XiaoxiaoNeural&quot;, // 必填：String 音色名必填
	//     &quot;Language&quot;: &quot;zh-CN&quot;, // 必填：String 合成的语言  
	//     &quot;Rate&quot;: 1 // 选填：float 语速  0.5～2 默认为 1
	// }
	// </code></pre>
	// 
	// </div></div><ul>
	// <li>自定义</li>
	// </ul>
	// <p>TTS<br>
	// 具体协议规范请参考<a href="https://doc.weixin.qq.com/doc/w3_ANQAiAbdAFwHILbJBmtSqSbV1WZ3L?scode=AJEAIQdfAAo5a1xajYANQAiAbdAFw" target="_blank">腾讯文档</a></p>
	// <div><div class="v-md-pre-wrapper copy-code-mode v-md-pre-wrapper- extra-class"><pre class="v-md-prism-"><code>{
	//   &quot;TTSType&quot;: &quot;custom&quot;, // String 必填
	//   &quot;APIKey&quot;: &quot;ApiKey&quot;, // String 必填 用来鉴权
	//   &quot;APIUrl&quot;: &quot;http://0.0.0.0:8080/stream-audio&quot; // String，必填，TTS API URL
	//   &quot;AudioFormat&quot;: &quot;wav&quot;, // String, 非必填，期望输出的音频格式，如mp3， ogg_opus，pcm，wav，默认为 wav，目前只支持pcm和wav，
	//   &quot;SampleRate&quot;: 16000,  // Integer，非必填，音频采样率，默认为16000(16k)，推荐值为16000
	//   &quot;AudioChannel&quot;: 1,    // Integer，非必填，音频通道数，取值：1 或 2  默认为1  
	// }
	// </code></pre>
	// 
	// </div></div>
	CustomTTSConfig *string `json:"CustomTTSConfig,omitnil,omitempty" name:"CustomTTSConfig"`

	// 提示词变量
	PromptVariables []*Variable `json:"PromptVariables,omitnil,omitempty" name:"PromptVariables"`

	// 语音识别vad的时间，范围为240-2000，默认为1000，单位为ms。更小的值会让语音识别分句更快。
	VadSilenceTime *int64 `json:"VadSilenceTime,omitnil,omitempty" name:"VadSilenceTime"`

	// 通话内容提取配置
	ExtractConfig []*AICallExtractConfigElement `json:"ExtractConfig,omitnil,omitempty" name:"ExtractConfig"`
}

func (r *CreateAICallRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAICallRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Callee")
	delete(f, "SystemPrompt")
	delete(f, "LLMType")
	delete(f, "Model")
	delete(f, "APIKey")
	delete(f, "APIUrl")
	delete(f, "VoiceType")
	delete(f, "Callers")
	delete(f, "WelcomeMessage")
	delete(f, "WelcomeType")
	delete(f, "WelcomeMessagePriority")
	delete(f, "MaxDuration")
	delete(f, "Languages")
	delete(f, "InterruptMode")
	delete(f, "InterruptSpeechDuration")
	delete(f, "EndFunctionEnable")
	delete(f, "EndFunctionDesc")
	delete(f, "TransferFunctionEnable")
	delete(f, "TransferItems")
	delete(f, "NotifyDuration")
	delete(f, "NotifyMessage")
	delete(f, "NotifyMaxCount")
	delete(f, "CustomTTSConfig")
	delete(f, "PromptVariables")
	delete(f, "VadSilenceTime")
	delete(f, "ExtractConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAICallRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAICallResponseParams struct {
	// 新创建的会话 ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAICallResponse struct {
	*tchttp.BaseResponse
	Response *CreateAICallResponseParams `json:"Response"`
}

func (r *CreateAICallResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAICallResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAdminURLRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 管理员账号
	SeatUserId *string `json:"SeatUserId,omitnil,omitempty" name:"SeatUserId"`
}

type CreateAdminURLRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 管理员账号
	SeatUserId *string `json:"SeatUserId,omitnil,omitempty" name:"SeatUserId"`
}

func (r *CreateAdminURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAdminURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SeatUserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAdminURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAdminURLResponseParams struct {
	// 登录链接
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAdminURLResponse struct {
	*tchttp.BaseResponse
	Response *CreateAdminURLResponseParams `json:"Response"`
}

func (r *CreateAdminURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAdminURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoCalloutTaskRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务起始时间戳，Unix 秒级时间戳
	NotBefore *int64 `json:"NotBefore,omitnil,omitempty" name:"NotBefore"`

	// 被叫号码列表
	Callees []*string `json:"Callees,omitnil,omitempty" name:"Callees"`

	// 主叫号码列表
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 呼叫使用的Ivr
	IvrId *uint64 `json:"IvrId,omitnil,omitempty" name:"IvrId"`

	// 任务名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 任务停止时间戳，Unix 秒级时间戳
	NotAfter *int64 `json:"NotAfter,omitnil,omitempty" name:"NotAfter"`

	// 最大尝试次数，1-3 次
	Tries *uint64 `json:"Tries,omitnil,omitempty" name:"Tries"`

	// 自定义变量（仅高级版支持）
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`

	// UUI
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`

	// 被叫属性
	CalleeAttributes []*CalleeAttribute `json:"CalleeAttributes,omitnil,omitempty" name:"CalleeAttributes"`
}

type CreateAutoCalloutTaskRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务起始时间戳，Unix 秒级时间戳
	NotBefore *int64 `json:"NotBefore,omitnil,omitempty" name:"NotBefore"`

	// 被叫号码列表
	Callees []*string `json:"Callees,omitnil,omitempty" name:"Callees"`

	// 主叫号码列表
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 呼叫使用的Ivr
	IvrId *uint64 `json:"IvrId,omitnil,omitempty" name:"IvrId"`

	// 任务名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 任务停止时间戳，Unix 秒级时间戳
	NotAfter *int64 `json:"NotAfter,omitnil,omitempty" name:"NotAfter"`

	// 最大尝试次数，1-3 次
	Tries *uint64 `json:"Tries,omitnil,omitempty" name:"Tries"`

	// 自定义变量（仅高级版支持）
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`

	// UUI
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`

	// 被叫属性
	CalleeAttributes []*CalleeAttribute `json:"CalleeAttributes,omitnil,omitempty" name:"CalleeAttributes"`
}

func (r *CreateAutoCalloutTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoCalloutTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "NotBefore")
	delete(f, "Callees")
	delete(f, "Callers")
	delete(f, "IvrId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "NotAfter")
	delete(f, "Tries")
	delete(f, "Variables")
	delete(f, "UUI")
	delete(f, "CalleeAttributes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAutoCalloutTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoCalloutTaskResponseParams struct {
	// 任务Id
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAutoCalloutTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAutoCalloutTaskResponseParams `json:"Response"`
}

func (r *CreateAutoCalloutTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoCalloutTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCCCSkillGroupRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 技能组名称
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`

	// 技能组类型0-电话，1-在线，3-音频，4-视频
	SkillGroupType *int64 `json:"SkillGroupType,omitnil,omitempty" name:"SkillGroupType"`

	// 技能组接待人数上限（该技能组中1个座席可接待的人数上限）默认为1。1、若技能组类型为在线，则接待上限可设置为1及以上
	// 2、若技能组类型为电话、音频、视频，则接待上线必须只能为1
	MaxConcurrency *uint64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`
}

type CreateCCCSkillGroupRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 技能组名称
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`

	// 技能组类型0-电话，1-在线，3-音频，4-视频
	SkillGroupType *int64 `json:"SkillGroupType,omitnil,omitempty" name:"SkillGroupType"`

	// 技能组接待人数上限（该技能组中1个座席可接待的人数上限）默认为1。1、若技能组类型为在线，则接待上限可设置为1及以上
	// 2、若技能组类型为电话、音频、视频，则接待上线必须只能为1
	MaxConcurrency *uint64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`
}

func (r *CreateCCCSkillGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCCSkillGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SkillGroupName")
	delete(f, "SkillGroupType")
	delete(f, "MaxConcurrency")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCCCSkillGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCCCSkillGroupResponseParams struct {
	// 技能组ID
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCCCSkillGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateCCCSkillGroupResponseParams `json:"Response"`
}

func (r *CreateCCCSkillGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCCSkillGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCallOutSessionRequestParams struct {
	// 应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 客服用户 ID，一般为客服邮箱
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 被叫号码，须带 0086 前缀
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// 主叫号码（废弃，使用Callers），须带 0086 前缀
	Caller *string `json:"Caller,omitnil,omitempty" name:"Caller"`

	// 指定主叫号码列表，如果前面的号码失败了会自动换成下一个号码，须带 0086 前缀
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 是否强制使用手机外呼，当前只支持 true，若为 true 请确保已配置白名单
	IsForceUseMobile *bool `json:"IsForceUseMobile,omitnil,omitempty" name:"IsForceUseMobile"`

	// 自定义数据，长度限制 1024 字节
	//
	// Deprecated: Uui is deprecated.
	Uui *string `json:"Uui,omitnil,omitempty" name:"Uui"`

	// 自定义数据，长度限制 1024 字节
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`
}

type CreateCallOutSessionRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 客服用户 ID，一般为客服邮箱
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 被叫号码，须带 0086 前缀
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// 主叫号码（废弃，使用Callers），须带 0086 前缀
	Caller *string `json:"Caller,omitnil,omitempty" name:"Caller"`

	// 指定主叫号码列表，如果前面的号码失败了会自动换成下一个号码，须带 0086 前缀
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 是否强制使用手机外呼，当前只支持 true，若为 true 请确保已配置白名单
	IsForceUseMobile *bool `json:"IsForceUseMobile,omitnil,omitempty" name:"IsForceUseMobile"`

	// 自定义数据，长度限制 1024 字节
	Uui *string `json:"Uui,omitnil,omitempty" name:"Uui"`

	// 自定义数据，长度限制 1024 字节
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`
}

func (r *CreateCallOutSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCallOutSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "UserId")
	delete(f, "Callee")
	delete(f, "Caller")
	delete(f, "Callers")
	delete(f, "IsForceUseMobile")
	delete(f, "Uui")
	delete(f, "UUI")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCallOutSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCallOutSessionResponseParams struct {
	// 新创建的会话 ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCallOutSessionResponse struct {
	*tchttp.BaseResponse
	Response *CreateCallOutSessionResponseParams `json:"Response"`
}

func (r *CreateCallOutSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCallOutSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCarrierPrivilegeNumberApplicantRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 主叫号码，必须为实例中存在的号码，格式为0086xxxx（暂时只支持国内号码）
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 被叫号码，必须为实例中坐席绑定的手机号码，格式为0086xxxx（暂时只支持国内号码）
	Callees []*string `json:"Callees,omitnil,omitempty" name:"Callees"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateCarrierPrivilegeNumberApplicantRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 主叫号码，必须为实例中存在的号码，格式为0086xxxx（暂时只支持国内号码）
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 被叫号码，必须为实例中坐席绑定的手机号码，格式为0086xxxx（暂时只支持国内号码）
	Callees []*string `json:"Callees,omitnil,omitempty" name:"Callees"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateCarrierPrivilegeNumberApplicantRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCarrierPrivilegeNumberApplicantRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Callers")
	delete(f, "Callees")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCarrierPrivilegeNumberApplicantRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCarrierPrivilegeNumberApplicantResponseParams struct {
	// 申请单Id
	ApplicantId *uint64 `json:"ApplicantId,omitnil,omitempty" name:"ApplicantId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCarrierPrivilegeNumberApplicantResponse struct {
	*tchttp.BaseResponse
	Response *CreateCarrierPrivilegeNumberApplicantResponseParams `json:"Response"`
}

func (r *CreateCarrierPrivilegeNumberApplicantResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCarrierPrivilegeNumberApplicantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCompanyApplyRequestParams struct {
	// 企业资质信息
	CompanyInfo *CompanyApplyInfo `json:"CompanyInfo,omitnil,omitempty" name:"CompanyInfo"`
}

type CreateCompanyApplyRequest struct {
	*tchttp.BaseRequest
	
	// 企业资质信息
	CompanyInfo *CompanyApplyInfo `json:"CompanyInfo,omitnil,omitempty" name:"CompanyInfo"`
}

func (r *CreateCompanyApplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCompanyApplyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCompanyApplyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCompanyApplyResponseParams struct {
	// 申请单ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCompanyApplyResponse struct {
	*tchttp.BaseResponse
	Response *CreateCompanyApplyResponseParams `json:"Response"`
}

func (r *CreateCompanyApplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCompanyApplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExtensionRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`

	// 分机名称
	ExtensionName *string `json:"ExtensionName,omitnil,omitempty" name:"ExtensionName"`

	// 绑定的技能组列表
	SkillGroupIds []*uint64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`

	// 绑定的坐席邮箱
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`
}

type CreateExtensionRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`

	// 分机名称
	ExtensionName *string `json:"ExtensionName,omitnil,omitempty" name:"ExtensionName"`

	// 绑定的技能组列表
	SkillGroupIds []*uint64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`

	// 绑定的坐席邮箱
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`
}

func (r *CreateExtensionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExtensionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "ExtensionId")
	delete(f, "ExtensionName")
	delete(f, "SkillGroupIds")
	delete(f, "Relation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExtensionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExtensionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateExtensionResponse struct {
	*tchttp.BaseResponse
	Response *CreateExtensionResponseParams `json:"Response"`
}

func (r *CreateExtensionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExtensionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIVRSessionRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 被叫
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// 指定的 IVR Id，目前支持呼入和自动外呼两种类型
	IVRId *int64 `json:"IVRId,omitnil,omitempty" name:"IVRId"`

	// 主叫号码列表
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 自定义变量
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`

	// 用户数据
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`
}

type CreateIVRSessionRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 被叫
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// 指定的 IVR Id，目前支持呼入和自动外呼两种类型
	IVRId *int64 `json:"IVRId,omitnil,omitempty" name:"IVRId"`

	// 主叫号码列表
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 自定义变量
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`

	// 用户数据
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`
}

func (r *CreateIVRSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIVRSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Callee")
	delete(f, "IVRId")
	delete(f, "Callers")
	delete(f, "Variables")
	delete(f, "UUI")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIVRSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIVRSessionResponseParams struct {
	// 新创建的会话 ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIVRSessionResponse struct {
	*tchttp.BaseResponse
	Response *CreateIVRSessionResponseParams `json:"Response"`
}

func (r *CreateIVRSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIVRSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOwnNumberApplyRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// SIP通道ID
	SipTrunkId *int64 `json:"SipTrunkId,omitnil,omitempty" name:"SipTrunkId"`

	// 线路相关参数
	DetailList []*OwnNumberApplyDetailItem `json:"DetailList,omitnil,omitempty" name:"DetailList"`

	// 送号前缀
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`
}

type CreateOwnNumberApplyRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// SIP通道ID
	SipTrunkId *int64 `json:"SipTrunkId,omitnil,omitempty" name:"SipTrunkId"`

	// 线路相关参数
	DetailList []*OwnNumberApplyDetailItem `json:"DetailList,omitnil,omitempty" name:"DetailList"`

	// 送号前缀
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`
}

func (r *CreateOwnNumberApplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOwnNumberApplyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SipTrunkId")
	delete(f, "DetailList")
	delete(f, "Prefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOwnNumberApplyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOwnNumberApplyResponseParams struct {
	// 审批单号
	ApplyId *uint64 `json:"ApplyId,omitnil,omitempty" name:"ApplyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOwnNumberApplyResponse struct {
	*tchttp.BaseResponse
	Response *CreateOwnNumberApplyResponseParams `json:"Response"`
}

func (r *CreateOwnNumberApplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOwnNumberApplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePredictiveDialingCampaignRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 被叫列表，支持 E.164 或不带国家码形式的号码
	Callees []*string `json:"Callees,omitnil,omitempty" name:"Callees"`

	// 主叫列表，使用管理端展示的号码格式
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 被叫呼叫顺序 0 随机 1 顺序
	CallOrder *int64 `json:"CallOrder,omitnil,omitempty" name:"CallOrder"`

	// 使用的座席技能组 ID
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// 相同应用内多个任务运行优先级，从高到底 1 - 5
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 预期呼损率，百分比，5 - 50
	ExpectedAbandonRate *int64 `json:"ExpectedAbandonRate,omitnil,omitempty" name:"ExpectedAbandonRate"`

	// 呼叫重试间隔时间，单位秒，60 - 86400
	RetryInterval *int64 `json:"RetryInterval,omitnil,omitempty" name:"RetryInterval"`

	// 任务启动时间，Unix 时间戳，到此时间后会自动启动任务
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间，Unix 时间戳，到此时间后会自动终止任务
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指定的 IVR Id
	IVRId *int64 `json:"IVRId,omitnil,omitempty" name:"IVRId"`

	// 呼叫重试次数，0 - 2
	RetryTimes *int64 `json:"RetryTimes,omitnil,omitempty" name:"RetryTimes"`

	// 自定义变量
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`

	// UUI
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`

	// 被叫属性
	CalleeAttributes []*CalleeAttribute `json:"CalleeAttributes,omitnil,omitempty" name:"CalleeAttributes"`
}

type CreatePredictiveDialingCampaignRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 被叫列表，支持 E.164 或不带国家码形式的号码
	Callees []*string `json:"Callees,omitnil,omitempty" name:"Callees"`

	// 主叫列表，使用管理端展示的号码格式
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 被叫呼叫顺序 0 随机 1 顺序
	CallOrder *int64 `json:"CallOrder,omitnil,omitempty" name:"CallOrder"`

	// 使用的座席技能组 ID
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// 相同应用内多个任务运行优先级，从高到底 1 - 5
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 预期呼损率，百分比，5 - 50
	ExpectedAbandonRate *int64 `json:"ExpectedAbandonRate,omitnil,omitempty" name:"ExpectedAbandonRate"`

	// 呼叫重试间隔时间，单位秒，60 - 86400
	RetryInterval *int64 `json:"RetryInterval,omitnil,omitempty" name:"RetryInterval"`

	// 任务启动时间，Unix 时间戳，到此时间后会自动启动任务
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间，Unix 时间戳，到此时间后会自动终止任务
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指定的 IVR Id
	IVRId *int64 `json:"IVRId,omitnil,omitempty" name:"IVRId"`

	// 呼叫重试次数，0 - 2
	RetryTimes *int64 `json:"RetryTimes,omitnil,omitempty" name:"RetryTimes"`

	// 自定义变量
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`

	// UUI
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`

	// 被叫属性
	CalleeAttributes []*CalleeAttribute `json:"CalleeAttributes,omitnil,omitempty" name:"CalleeAttributes"`
}

func (r *CreatePredictiveDialingCampaignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePredictiveDialingCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Name")
	delete(f, "Callees")
	delete(f, "Callers")
	delete(f, "CallOrder")
	delete(f, "SkillGroupId")
	delete(f, "Priority")
	delete(f, "ExpectedAbandonRate")
	delete(f, "RetryInterval")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "IVRId")
	delete(f, "RetryTimes")
	delete(f, "Variables")
	delete(f, "UUI")
	delete(f, "CalleeAttributes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePredictiveDialingCampaignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePredictiveDialingCampaignResponseParams struct {
	// 生成的任务 ID
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePredictiveDialingCampaignResponse struct {
	*tchttp.BaseResponse
	Response *CreatePredictiveDialingCampaignResponseParams `json:"Response"`
}

func (r *CreatePredictiveDialingCampaignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePredictiveDialingCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSDKLoginTokenRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 座席账号。
	SeatUserId *string `json:"SeatUserId,omitnil,omitempty" name:"SeatUserId"`

	// 生成的token是否一次性校验
	OnlyOnce *bool `json:"OnlyOnce,omitnil,omitempty" name:"OnlyOnce"`
}

type CreateSDKLoginTokenRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 座席账号。
	SeatUserId *string `json:"SeatUserId,omitnil,omitempty" name:"SeatUserId"`

	// 生成的token是否一次性校验
	OnlyOnce *bool `json:"OnlyOnce,omitnil,omitempty" name:"OnlyOnce"`
}

func (r *CreateSDKLoginTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSDKLoginTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SeatUserId")
	delete(f, "OnlyOnce")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSDKLoginTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSDKLoginTokenResponseParams struct {
	// SDK 登录 Token。
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 过期时间戳，Unix 时间戳。
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// SDK 加载路径会随着 SDK 的发布而变动。
	SdkURL *string `json:"SdkURL,omitnil,omitempty" name:"SdkURL"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSDKLoginTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateSDKLoginTokenResponseParams `json:"Response"`
}

func (r *CreateSDKLoginTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSDKLoginTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStaffRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 客服信息，个数不超过 10
	Staffs []*SeatUserInfo `json:"Staffs,omitnil,omitempty" name:"Staffs"`

	// 是否发送密码邮件，默认true
	SendPassword *bool `json:"SendPassword,omitnil,omitempty" name:"SendPassword"`
}

type CreateStaffRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 客服信息，个数不超过 10
	Staffs []*SeatUserInfo `json:"Staffs,omitnil,omitempty" name:"Staffs"`

	// 是否发送密码邮件，默认true
	SendPassword *bool `json:"SendPassword,omitnil,omitempty" name:"SendPassword"`
}

func (r *CreateStaffRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStaffRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Staffs")
	delete(f, "SendPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStaffRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStaffResponseParams struct {
	// 错误坐席列表及错误信息
	ErrorStaffList []*ErrStaffItem `json:"ErrorStaffList,omitnil,omitempty" name:"ErrorStaffList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStaffResponse struct {
	*tchttp.BaseResponse
	Response *CreateStaffResponseParams `json:"Response"`
}

func (r *CreateStaffResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStaffResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserSigRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 用户 ID，该值必须与 ClientData 字段中 Uid 的值一致
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 有效期，单位秒，不超过 1 小时
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 用户签名数据，必填字段，为标准 JSON 格式
	ClientData *string `json:"ClientData,omitnil,omitempty" name:"ClientData"`
}

type CreateUserSigRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 用户 ID，该值必须与 ClientData 字段中 Uid 的值一致
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 有效期，单位秒，不超过 1 小时
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 用户签名数据，必填字段，为标准 JSON 格式
	ClientData *string `json:"ClientData,omitnil,omitempty" name:"ClientData"`
}

func (r *CreateUserSigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserSigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Uid")
	delete(f, "ExpiredTime")
	delete(f, "ClientData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserSigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserSigResponseParams struct {
	// 签名结果
	UserSig *string `json:"UserSig,omitnil,omitempty" name:"UserSig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserSigResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserSigResponseParams `json:"Response"`
}

func (r *CreateUserSigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserSigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteExtensionRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`
}

type DeleteExtensionRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`
}

func (r *DeleteExtensionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExtensionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "ExtensionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteExtensionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteExtensionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteExtensionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteExtensionResponseParams `json:"Response"`
}

func (r *DeleteExtensionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExtensionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePredictiveDialingCampaignRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务 ID
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

type DeletePredictiveDialingCampaignRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务 ID
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

func (r *DeletePredictiveDialingCampaignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePredictiveDialingCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CampaignId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePredictiveDialingCampaignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePredictiveDialingCampaignResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePredictiveDialingCampaignResponse struct {
	*tchttp.BaseResponse
	Response *DeletePredictiveDialingCampaignResponseParams `json:"Response"`
}

func (r *DeletePredictiveDialingCampaignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePredictiveDialingCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStaffRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 待删除客服邮箱列表，一次最大支持200个。
	StaffList []*string `json:"StaffList,omitnil,omitempty" name:"StaffList"`
}

type DeleteStaffRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 待删除客服邮箱列表，一次最大支持200个。
	StaffList []*string `json:"StaffList,omitnil,omitempty" name:"StaffList"`
}

func (r *DeleteStaffRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStaffRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StaffList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStaffRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStaffResponseParams struct {
	// 无法删除的状态为在线的客服列表
	OnlineStaffList []*string `json:"OnlineStaffList,omitnil,omitempty" name:"OnlineStaffList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteStaffResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStaffResponseParams `json:"Response"`
}

func (r *DeleteStaffResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStaffResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAICallExtractResultRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 会话 ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 查找起始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查找结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeAICallExtractResultRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 会话 ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 查找起始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查找结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeAICallExtractResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAICallExtractResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SessionId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAICallExtractResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAICallExtractResultResponseParams struct {
	// 结果列表
	ResultList []*AICallExtractResultElement `json:"ResultList,omitnil,omitempty" name:"ResultList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAICallExtractResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAICallExtractResultResponseParams `json:"Response"`
}

func (r *DescribeAICallExtractResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAICallExtractResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActiveCarrierPrivilegeNumberRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 默认0
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 默认10，最大100
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 筛选条件 Name支持PhoneNumber(按号码模糊查找)
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeActiveCarrierPrivilegeNumberRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 默认0
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 默认10，最大100
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 筛选条件 Name支持PhoneNumber(按号码模糊查找)
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeActiveCarrierPrivilegeNumberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActiveCarrierPrivilegeNumberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeActiveCarrierPrivilegeNumberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActiveCarrierPrivilegeNumberResponseParams struct {
	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 生效列表
	ActiveCarrierPrivilegeNumbers []*ActiveCarrierPrivilegeNumber `json:"ActiveCarrierPrivilegeNumbers,omitnil,omitempty" name:"ActiveCarrierPrivilegeNumbers"`

	// 待审核单号
	PendingApplicantIds []*uint64 `json:"PendingApplicantIds,omitnil,omitempty" name:"PendingApplicantIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeActiveCarrierPrivilegeNumberResponse struct {
	*tchttp.BaseResponse
	Response *DescribeActiveCarrierPrivilegeNumberResponseParams `json:"Response"`
}

func (r *DescribeActiveCarrierPrivilegeNumberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActiveCarrierPrivilegeNumberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoCalloutTaskRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务Id
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeAutoCalloutTaskRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务Id
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeAutoCalloutTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoCalloutTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoCalloutTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoCalloutTaskResponseParams struct {
	// 任务名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 任务起始时间戳
	NotBefore *int64 `json:"NotBefore,omitnil,omitempty" name:"NotBefore"`

	// 任务结束时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotAfter *int64 `json:"NotAfter,omitnil,omitempty" name:"NotAfter"`

	// 主叫列表
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 被叫信息列表
	Callees []*AutoCalloutTaskCalleeInfo `json:"Callees,omitnil,omitempty" name:"Callees"`

	// 任务使用的IvrId
	IvrId *uint64 `json:"IvrId,omitnil,omitempty" name:"IvrId"`

	// 任务状态 0初始 1运行中 2已完成 3结束中 4已终止
	State *uint64 `json:"State,omitnil,omitempty" name:"State"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoCalloutTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoCalloutTaskResponseParams `json:"Response"`
}

func (r *DescribeAutoCalloutTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoCalloutTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoCalloutTasksRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type DescribeAutoCalloutTasksRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

func (r *DescribeAutoCalloutTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoCalloutTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoCalloutTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoCalloutTasksResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务列表
	Tasks []*AutoCalloutTaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoCalloutTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoCalloutTasksResponseParams `json:"Response"`
}

func (r *DescribeAutoCalloutTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoCalloutTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCCBuyInfoListRequestParams struct {
	// 应用ID列表，不传时查询所有应用
	SdkAppIds []*int64 `json:"SdkAppIds,omitnil,omitempty" name:"SdkAppIds"`
}

type DescribeCCCBuyInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID列表，不传时查询所有应用
	SdkAppIds []*int64 `json:"SdkAppIds,omitnil,omitempty" name:"SdkAppIds"`
}

func (r *DescribeCCCBuyInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCCBuyInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCCBuyInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCCBuyInfoListResponseParams struct {
	// 应用总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 应用购买信息列表
	SdkAppIdBuyList []*SdkAppIdBuyInfo `json:"SdkAppIdBuyList,omitnil,omitempty" name:"SdkAppIdBuyList"`

	// 套餐包购买信息列表
	PackageBuyList []*PackageBuyInfo `json:"PackageBuyList,omitnil,omitempty" name:"PackageBuyList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCCCBuyInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCCBuyInfoListResponseParams `json:"Response"`
}

func (r *DescribeCCCBuyInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCCBuyInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallInMetricsRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 是否返回技能组维度信息，默认“是”
	EnabledSkillGroup *bool `json:"EnabledSkillGroup,omitnil,omitempty" name:"EnabledSkillGroup"`

	// 是否返回线路维度信息，默认“否”
	EnabledNumber *bool `json:"EnabledNumber,omitnil,omitempty" name:"EnabledNumber"`

	// 筛选技能组列表
	GroupIdList []*int64 `json:"GroupIdList,omitnil,omitempty" name:"GroupIdList"`
}

type DescribeCallInMetricsRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 是否返回技能组维度信息，默认“是”
	EnabledSkillGroup *bool `json:"EnabledSkillGroup,omitnil,omitempty" name:"EnabledSkillGroup"`

	// 是否返回线路维度信息，默认“否”
	EnabledNumber *bool `json:"EnabledNumber,omitnil,omitempty" name:"EnabledNumber"`

	// 筛选技能组列表
	GroupIdList []*int64 `json:"GroupIdList,omitnil,omitempty" name:"GroupIdList"`
}

func (r *DescribeCallInMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallInMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "EnabledSkillGroup")
	delete(f, "EnabledNumber")
	delete(f, "GroupIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCallInMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallInMetricsResponseParams struct {
	// 时间戳
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 总体指标
	TotalMetrics *CallInMetrics `json:"TotalMetrics,omitnil,omitempty" name:"TotalMetrics"`

	// 线路维度指标
	NumberMetrics []*CallInNumberMetrics `json:"NumberMetrics,omitnil,omitempty" name:"NumberMetrics"`

	// 技能组维度指标
	SkillGroupMetrics []*CallInSkillGroupMetrics `json:"SkillGroupMetrics,omitnil,omitempty" name:"SkillGroupMetrics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCallInMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCallInMetricsResponseParams `json:"Response"`
}

func (r *DescribeCallInMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallInMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCarrierPrivilegeNumberApplicantsRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 默认0，从0开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 默认10，最大100
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 筛选条件,Name支持ApplicantId,PhoneNumber(按号码模糊查找)
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeCarrierPrivilegeNumberApplicantsRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 默认0，从0开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 默认10，最大100
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 筛选条件,Name支持ApplicantId,PhoneNumber(按号码模糊查找)
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeCarrierPrivilegeNumberApplicantsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCarrierPrivilegeNumberApplicantsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCarrierPrivilegeNumberApplicantsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCarrierPrivilegeNumberApplicantsResponseParams struct {
	// 筛选出的总申请单数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 申请单列表
	Applicants []*CarrierPrivilegeNumberApplicant `json:"Applicants,omitnil,omitempty" name:"Applicants"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCarrierPrivilegeNumberApplicantsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCarrierPrivilegeNumberApplicantsResponseParams `json:"Response"`
}

func (r *DescribeCarrierPrivilegeNumberApplicantsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCarrierPrivilegeNumberApplicantsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChatMessagesRequestParams struct {
	// 应用 ID，可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 实例 ID（废弃）
	//
	// Deprecated: InstanceId is deprecated.
	InstanceId *int64 `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务记录ID（废弃）
	//
	// Deprecated: CdrId is deprecated.
	CdrId *string `json:"CdrId,omitnil,omitempty" name:"CdrId"`

	// 返回记录条数，最大为100 默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回记录偏移，默认为 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 1为从早到晚，2为从晚到早，默认为2
	Order *int64 `json:"Order,omitnil,omitempty" name:"Order"`

	// 服务记录 SessionID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type DescribeChatMessagesRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID，可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 实例 ID（废弃）
	InstanceId *int64 `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务记录ID（废弃）
	CdrId *string `json:"CdrId,omitnil,omitempty" name:"CdrId"`

	// 返回记录条数，最大为100 默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回记录偏移，默认为 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 1为从早到晚，2为从晚到早，默认为2
	Order *int64 `json:"Order,omitnil,omitempty" name:"Order"`

	// 服务记录 SessionID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *DescribeChatMessagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChatMessagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "InstanceId")
	delete(f, "CdrId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChatMessagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChatMessagesResponseParams struct {
	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 消息列表
	Messages []*MessageBody `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeChatMessagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChatMessagesResponseParams `json:"Response"`
}

func (r *DescribeChatMessagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChatMessagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompanyListRequestParams struct {
	// 分页尺寸，上限 100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 公司名称
	CompanyName []*string `json:"CompanyName,omitnil,omitempty" name:"CompanyName"`

	// 审核状态，1-待审核，2-审核通过，3-驳回
	State []*int64 `json:"State,omitnil,omitempty" name:"State"`

	// 申请ID
	ApplyID []*int64 `json:"ApplyID,omitnil,omitempty" name:"ApplyID"`
}

type DescribeCompanyListRequest struct {
	*tchttp.BaseRequest
	
	// 分页尺寸，上限 100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 公司名称
	CompanyName []*string `json:"CompanyName,omitnil,omitempty" name:"CompanyName"`

	// 审核状态，1-待审核，2-审核通过，3-驳回
	State []*int64 `json:"State,omitnil,omitempty" name:"State"`

	// 申请ID
	ApplyID []*int64 `json:"ApplyID,omitnil,omitempty" name:"ApplyID"`
}

func (r *DescribeCompanyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompanyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "CompanyName")
	delete(f, "State")
	delete(f, "ApplyID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCompanyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompanyListResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 企业资质审核信息
	CompanyInfo []*CompanyStateInfo `json:"CompanyInfo,omitnil,omitempty" name:"CompanyInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCompanyListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCompanyListResponseParams `json:"Response"`
}

func (r *DescribeCompanyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompanyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtensionRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`
}

type DescribeExtensionRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`
}

func (r *DescribeExtensionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtensionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "ExtensionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExtensionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtensionResponseParams struct {
	// 分机号
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`

	// 域名
	ExtensionDomain *string `json:"ExtensionDomain,omitnil,omitempty" name:"ExtensionDomain"`

	// 注册密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 代理服务器地址
	OutboundProxy *string `json:"OutboundProxy,omitnil,omitempty" name:"OutboundProxy"`

	// 传输协议
	Transport *string `json:"Transport,omitnil,omitempty" name:"Transport"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExtensionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExtensionResponseParams `json:"Response"`
}

func (r *DescribeExtensionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtensionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtensionsRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页页号（从0开始）
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 筛选分机号列表
	ExtensionIds []*string `json:"ExtensionIds,omitnil,omitempty" name:"ExtensionIds"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 模糊查询字段（模糊查询分机号、分机名称、坐席邮箱、坐席名称）
	FuzzingKeyWord *string `json:"FuzzingKeyWord,omitnil,omitempty" name:"FuzzingKeyWord"`

	// 是否需要返回话机当前状态
	IsNeedStatus *bool `json:"IsNeedStatus,omitnil,omitempty" name:"IsNeedStatus"`
}

type DescribeExtensionsRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页页号（从0开始）
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 筛选分机号列表
	ExtensionIds []*string `json:"ExtensionIds,omitnil,omitempty" name:"ExtensionIds"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 模糊查询字段（模糊查询分机号、分机名称、坐席邮箱、坐席名称）
	FuzzingKeyWord *string `json:"FuzzingKeyWord,omitnil,omitempty" name:"FuzzingKeyWord"`

	// 是否需要返回话机当前状态
	IsNeedStatus *bool `json:"IsNeedStatus,omitnil,omitempty" name:"IsNeedStatus"`
}

func (r *DescribeExtensionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtensionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageNumber")
	delete(f, "ExtensionIds")
	delete(f, "PageSize")
	delete(f, "FuzzingKeyWord")
	delete(f, "IsNeedStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExtensionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtensionsResponseParams struct {
	// 查询总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 话机信息列表
	ExtensionList []*ExtensionInfo `json:"ExtensionList,omitnil,omitempty" name:"ExtensionList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExtensionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExtensionsResponseParams `json:"Response"`
}

func (r *DescribeExtensionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtensionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIMCdrListRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 起始时间（必填），Unix 秒级时间戳
	StartTimestamp *int64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// 结束时间（必填），Unix 秒级时间戳
	EndTimestamp *int64 `json:"EndTimestamp,omitnil,omitempty" name:"EndTimestamp"`

	// 返回记录条数，最大为100默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回记录偏移，默认为 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 1为全媒体，2为文本客服，不填则查询全部
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeIMCdrListRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 起始时间（必填），Unix 秒级时间戳
	StartTimestamp *int64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// 结束时间（必填），Unix 秒级时间戳
	EndTimestamp *int64 `json:"EndTimestamp,omitnil,omitempty" name:"EndTimestamp"`

	// 返回记录条数，最大为100默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回记录偏移，默认为 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 1为全媒体，2为文本客服，不填则查询全部
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeIMCdrListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIMCdrListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTimestamp")
	delete(f, "EndTimestamp")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIMCdrListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIMCdrListResponseParams struct {
	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 服务记录列表
	IMCdrList []*IMCdrInfo `json:"IMCdrList,omitnil,omitempty" name:"IMCdrList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIMCdrListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIMCdrListResponseParams `json:"Response"`
}

func (r *DescribeIMCdrListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIMCdrListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIMCdrsRequestParams struct {
	// 起始时间（必填），Unix 秒级时间戳
	StartTimestamp *int64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// 结束时间（必填），Unix 秒级时间戳
	EndTimestamp *int64 `json:"EndTimestamp,omitnil,omitempty" name:"EndTimestamp"`

	// 实例 ID（废弃）
	//
	// Deprecated: InstanceId is deprecated.
	InstanceId *int64 `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 返回记录条数，最大为100默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回记录偏移，默认为 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 1为全媒体，2为文本客服，不填则查询全部
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeIMCdrsRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间（必填），Unix 秒级时间戳
	StartTimestamp *int64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// 结束时间（必填），Unix 秒级时间戳
	EndTimestamp *int64 `json:"EndTimestamp,omitnil,omitempty" name:"EndTimestamp"`

	// 实例 ID（废弃）
	InstanceId *int64 `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 返回记录条数，最大为100默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回记录偏移，默认为 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 1为全媒体，2为文本客服，不填则查询全部
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeIMCdrsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIMCdrsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTimestamp")
	delete(f, "EndTimestamp")
	delete(f, "InstanceId")
	delete(f, "SdkAppId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIMCdrsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIMCdrsResponseParams struct {
	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 服务记录列表
	//
	// Deprecated: IMCdrs is deprecated.
	IMCdrs []*IMCdrInfo `json:"IMCdrs,omitnil,omitempty" name:"IMCdrs"`

	// 服务记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	IMCdrList []*IMCdrInfo `json:"IMCdrList,omitnil,omitempty" name:"IMCdrList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIMCdrsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIMCdrsResponseParams `json:"Response"`
}

func (r *DescribeIMCdrsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIMCdrsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIvrAudioListRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 50
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 文件别名
	CustomFileName []*string `json:"CustomFileName,omitnil,omitempty" name:"CustomFileName"`

	// 文件名
	AudioFileName []*string `json:"AudioFileName,omitnil,omitempty" name:"AudioFileName"`

	// 文件ID
	FileId []*uint64 `json:"FileId,omitnil,omitempty" name:"FileId"`
}

type DescribeIvrAudioListRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 50
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 文件别名
	CustomFileName []*string `json:"CustomFileName,omitnil,omitempty" name:"CustomFileName"`

	// 文件名
	AudioFileName []*string `json:"AudioFileName,omitnil,omitempty" name:"AudioFileName"`

	// 文件ID
	FileId []*uint64 `json:"FileId,omitnil,omitempty" name:"FileId"`
}

func (r *DescribeIvrAudioListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIvrAudioListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "CustomFileName")
	delete(f, "AudioFileName")
	delete(f, "FileId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIvrAudioListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIvrAudioListResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 文件信息
	FileInfo []*AudioFileInfo `json:"FileInfo,omitnil,omitempty" name:"FileInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIvrAudioListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIvrAudioListResponseParams `json:"Response"`
}

func (r *DescribeIvrAudioListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIvrAudioListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNumbersRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 页数，从0开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小，默认20
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeNumbersRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 页数，从0开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小，默认20
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeNumbersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNumbersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNumbersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNumbersResponseParams struct {
	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 号码列表
	Numbers []*NumberInfo `json:"Numbers,omitnil,omitempty" name:"Numbers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNumbersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNumbersResponseParams `json:"Response"`
}

func (r *DescribeNumbersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNumbersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePSTNActiveSessionListRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 数据偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回的数据条数，最大 25
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePSTNActiveSessionListRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 数据偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回的数据条数，最大 25
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribePSTNActiveSessionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePSTNActiveSessionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePSTNActiveSessionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePSTNActiveSessionListResponseParams struct {
	// 列表总条数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 列表内容
	Sessions []*PSTNSessionInfo `json:"Sessions,omitnil,omitempty" name:"Sessions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePSTNActiveSessionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePSTNActiveSessionListResponseParams `json:"Response"`
}

func (r *DescribePSTNActiveSessionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePSTNActiveSessionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePredictiveDialingCampaignRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务 ID
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

type DescribePredictiveDialingCampaignRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务 ID
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

func (r *DescribePredictiveDialingCampaignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePredictiveDialingCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CampaignId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePredictiveDialingCampaignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePredictiveDialingCampaignResponseParams struct {
	// 任务 ID
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`

	// 任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 被叫呼叫顺序 0 随机 1 顺序
	CallOrder *int64 `json:"CallOrder,omitnil,omitempty" name:"CallOrder"`

	// 使用的座席技能组 ID
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// 指定的 IVR ID
	IVRId *int64 `json:"IVRId,omitnil,omitempty" name:"IVRId"`

	// 相同应用内多个任务运行优先级，从高到底 1 - 5
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 预期呼损率，百分比，5 - 50
	ExpectedAbandonRate *int64 `json:"ExpectedAbandonRate,omitnil,omitempty" name:"ExpectedAbandonRate"`

	// 呼叫重试次数，0 - 2
	RetryTimes *int64 `json:"RetryTimes,omitnil,omitempty" name:"RetryTimes"`

	// 呼叫重试间隔时间，单位秒，60 - 86400
	RetryInterval *int64 `json:"RetryInterval,omitnil,omitempty" name:"RetryInterval"`

	// 任务启动时间，Unix 时间戳，到此时间后会自动启动任务
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间，Unix 时间戳，到此时间后会自动终止任务
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePredictiveDialingCampaignResponse struct {
	*tchttp.BaseResponse
	Response *DescribePredictiveDialingCampaignResponseParams `json:"Response"`
}

func (r *DescribePredictiveDialingCampaignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePredictiveDialingCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePredictiveDialingCampaignsElement struct {
	// 任务 ID
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`

	// 任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务状态 0 待开始 1 进行中 2 已暂停 3 已终止 4 已完成
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务状态原因 0 正常 1 手动结束 2 超时结束
	StatusReason *int64 `json:"StatusReason,omitnil,omitempty" name:"StatusReason"`

	// 被叫号码个数
	CalleeCount *int64 `json:"CalleeCount,omitnil,omitempty" name:"CalleeCount"`

	// 已完成的被叫个数
	FinishedCalleeCount *int64 `json:"FinishedCalleeCount,omitnil,omitempty" name:"FinishedCalleeCount"`

	// 相同应用内多个任务运行优先级，从高到底 1 - 5
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 使用的座席技能组 ID
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`
}

// Predefined struct for user
type DescribePredictiveDialingCampaignsRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页尺寸，最大为 100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 查询任务列表名称关键字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 查询任务列表技能组 ID
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`
}

type DescribePredictiveDialingCampaignsRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页尺寸，最大为 100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 查询任务列表名称关键字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 查询任务列表技能组 ID
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`
}

func (r *DescribePredictiveDialingCampaignsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePredictiveDialingCampaignsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Name")
	delete(f, "SkillGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePredictiveDialingCampaignsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePredictiveDialingCampaignsResponseParams struct {
	// 数据总量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 数据
	CampaignList []*DescribePredictiveDialingCampaignsElement `json:"CampaignList,omitnil,omitempty" name:"CampaignList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePredictiveDialingCampaignsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePredictiveDialingCampaignsResponseParams `json:"Response"`
}

func (r *DescribePredictiveDialingCampaignsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePredictiveDialingCampaignsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePredictiveDialingSessionsRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 生成的任务 ID
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`

	// 分页尺寸，最大为 1000
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type DescribePredictiveDialingSessionsRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 生成的任务 ID
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`

	// 分页尺寸，最大为 1000
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

func (r *DescribePredictiveDialingSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePredictiveDialingSessionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CampaignId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePredictiveDialingSessionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePredictiveDialingSessionsResponseParams struct {
	// 数据总量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 呼叫的 session id 列表，通过 https://cloud.tencent.com/document/product/679/47714 可以批量获取呼叫详细话单
	SessionList []*string `json:"SessionList,omitnil,omitempty" name:"SessionList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePredictiveDialingSessionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePredictiveDialingSessionsResponseParams `json:"Response"`
}

func (r *DescribePredictiveDialingSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePredictiveDialingSessionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProtectedTelCdrRequestParams struct {
	// 起始时间戳，Unix 秒级时间戳
	StartTimeStamp *int64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 秒级时间戳
	EndTimeStamp *int64 `json:"EndTimeStamp,omitnil,omitempty" name:"EndTimeStamp"`

	// 应用 ID，可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 100
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type DescribeProtectedTelCdrRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间戳，Unix 秒级时间戳
	StartTimeStamp *int64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 秒级时间戳
	EndTimeStamp *int64 `json:"EndTimeStamp,omitnil,omitempty" name:"EndTimeStamp"`

	// 应用 ID，可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 100
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

func (r *DescribeProtectedTelCdrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProtectedTelCdrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTimeStamp")
	delete(f, "EndTimeStamp")
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProtectedTelCdrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProtectedTelCdrResponseParams struct {
	// 话单记录总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 话单记录
	//
	// Deprecated: TelCdrs is deprecated.
	TelCdrs []*TelCdrInfo `json:"TelCdrs,omitnil,omitempty" name:"TelCdrs"`

	// 话单记录
	TelCdrList []*TelCdrInfo `json:"TelCdrList,omitnil,omitempty" name:"TelCdrList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProtectedTelCdrResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProtectedTelCdrResponseParams `json:"Response"`
}

func (r *DescribeProtectedTelCdrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProtectedTelCdrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillGroupInfoListRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 技能组ID，查询单个技能组时使用
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// 查询修改时间大于等于ModifiedTime的技能组时使用
	ModifiedTime *int64 `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// 技能组名称
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`
}

type DescribeSkillGroupInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 技能组ID，查询单个技能组时使用
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// 查询修改时间大于等于ModifiedTime的技能组时使用
	ModifiedTime *int64 `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// 技能组名称
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`
}

func (r *DescribeSkillGroupInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillGroupInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "SkillGroupId")
	delete(f, "ModifiedTime")
	delete(f, "SkillGroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSkillGroupInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillGroupInfoListResponseParams struct {
	// 技能组总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 技能组信息列表
	SkillGroupList []*SkillGroupInfoItem `json:"SkillGroupList,omitnil,omitempty" name:"SkillGroupList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSkillGroupInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSkillGroupInfoListResponseParams `json:"Response"`
}

func (r *DescribeSkillGroupInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillGroupInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStaffInfoListRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 9999
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 坐席账号，查询单个坐席时使用
	StaffMail *string `json:"StaffMail,omitnil,omitempty" name:"StaffMail"`

	// 查询修改时间大于等于ModifiedTime的坐席时使用
	ModifiedTime *int64 `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// 技能组ID
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`
}

type DescribeStaffInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 9999
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 坐席账号，查询单个坐席时使用
	StaffMail *string `json:"StaffMail,omitnil,omitempty" name:"StaffMail"`

	// 查询修改时间大于等于ModifiedTime的坐席时使用
	ModifiedTime *int64 `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// 技能组ID
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`
}

func (r *DescribeStaffInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaffInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "StaffMail")
	delete(f, "ModifiedTime")
	delete(f, "SkillGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStaffInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStaffInfoListResponseParams struct {
	// 坐席用户总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 坐席用户信息列表
	StaffList []*StaffInfo `json:"StaffList,omitnil,omitempty" name:"StaffList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStaffInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStaffInfoListResponseParams `json:"Response"`
}

func (r *DescribeStaffInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaffInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStaffStatusMetricsRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 筛选坐席列表，默认不传返回全部坐席信息
	StaffList []*string `json:"StaffList,omitnil,omitempty" name:"StaffList"`

	// 筛选技能组ID列表
	GroupIdList []*int64 `json:"GroupIdList,omitnil,omitempty" name:"GroupIdList"`

	// 筛选坐席状态列表 座席状态 free 示闲 | busy 忙碌 | rest 小休 | notReady 示忙 | afterCallWork 话后调整 | offline 离线 
	StatusList []*string `json:"StatusList,omitnil,omitempty" name:"StatusList"`
}

type DescribeStaffStatusMetricsRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 筛选坐席列表，默认不传返回全部坐席信息
	StaffList []*string `json:"StaffList,omitnil,omitempty" name:"StaffList"`

	// 筛选技能组ID列表
	GroupIdList []*int64 `json:"GroupIdList,omitnil,omitempty" name:"GroupIdList"`

	// 筛选坐席状态列表 座席状态 free 示闲 | busy 忙碌 | rest 小休 | notReady 示忙 | afterCallWork 话后调整 | offline 离线 
	StatusList []*string `json:"StatusList,omitnil,omitempty" name:"StatusList"`
}

func (r *DescribeStaffStatusMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaffStatusMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StaffList")
	delete(f, "GroupIdList")
	delete(f, "StatusList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStaffStatusMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStaffStatusMetricsResponseParams struct {
	// 坐席状态实时信息
	Metrics []*StaffStatusMetrics `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStaffStatusMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStaffStatusMetricsResponseParams `json:"Response"`
}

func (r *DescribeStaffStatusMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaffStatusMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTelCallInfoRequestParams struct {
	// 起始时间戳，Unix 时间戳(查询维度仅支持天，例如查询5月1日应该传startTime:"2023-05-01 00:00:00","endTime":"2023-05-01 23:59:59"的时间戳,查5月1日和5月2日的应该传startTime:"2023-05-01 00:00:00","endTime":"2023-05-02 23:59:59"的时间戳)
	StartTimeStamp *int64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 时间戳，查询时间范围最大为90天(查询维度仅支持天，例如查询5月1日应该传startTime:"2023-05-01 00:00:00","endTime":"2023-05-01 23:59:59"的时间戳,查5月1日和5月2日的应该传startTime:"2023-05-01 00:00:00","endTime":"2023-05-02 23:59:59"的时间戳)
	EndTimeStamp *int64 `json:"EndTimeStamp,omitnil,omitempty" name:"EndTimeStamp"`

	// 应用ID列表，多个ID时，返回值为多个ID使用总和
	SdkAppIdList []*int64 `json:"SdkAppIdList,omitnil,omitempty" name:"SdkAppIdList"`
}

type DescribeTelCallInfoRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间戳，Unix 时间戳(查询维度仅支持天，例如查询5月1日应该传startTime:"2023-05-01 00:00:00","endTime":"2023-05-01 23:59:59"的时间戳,查5月1日和5月2日的应该传startTime:"2023-05-01 00:00:00","endTime":"2023-05-02 23:59:59"的时间戳)
	StartTimeStamp *int64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 时间戳，查询时间范围最大为90天(查询维度仅支持天，例如查询5月1日应该传startTime:"2023-05-01 00:00:00","endTime":"2023-05-01 23:59:59"的时间戳,查5月1日和5月2日的应该传startTime:"2023-05-01 00:00:00","endTime":"2023-05-02 23:59:59"的时间戳)
	EndTimeStamp *int64 `json:"EndTimeStamp,omitnil,omitempty" name:"EndTimeStamp"`

	// 应用ID列表，多个ID时，返回值为多个ID使用总和
	SdkAppIdList []*int64 `json:"SdkAppIdList,omitnil,omitempty" name:"SdkAppIdList"`
}

func (r *DescribeTelCallInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelCallInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTimeStamp")
	delete(f, "EndTimeStamp")
	delete(f, "SdkAppIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTelCallInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTelCallInfoResponseParams struct {
	// 呼出套餐包消耗分钟数
	TelCallOutCount *int64 `json:"TelCallOutCount,omitnil,omitempty" name:"TelCallOutCount"`

	// 呼入套餐包消耗分钟数
	TelCallInCount *int64 `json:"TelCallInCount,omitnil,omitempty" name:"TelCallInCount"`

	// 坐席使用统计个数
	SeatUsedCount *int64 `json:"SeatUsedCount,omitnil,omitempty" name:"SeatUsedCount"`

	// 音频套餐包消耗分钟数
	//
	// Deprecated: VoipCallInCount is deprecated.
	VoipCallInCount *int64 `json:"VoipCallInCount,omitnil,omitempty" name:"VoipCallInCount"`

	// 音频套餐包消耗分钟数
	VOIPCallInCount *int64 `json:"VOIPCallInCount,omitnil,omitempty" name:"VOIPCallInCount"`

	// 离线语音转文字套餐包消耗分钟数
	AsrOfflineCount *int64 `json:"AsrOfflineCount,omitnil,omitempty" name:"AsrOfflineCount"`

	// 实时语音转文字套餐包消耗分钟数
	AsrRealtimeCount *int64 `json:"AsrRealtimeCount,omitnil,omitempty" name:"AsrRealtimeCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTelCallInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTelCallInfoResponseParams `json:"Response"`
}

func (r *DescribeTelCallInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelCallInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTelCdrRequestParams struct {
	// 起始时间戳，Unix 秒级时间戳，最大支持近180天。
	StartTimeStamp *int64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 秒级时间戳，结束时间与开始时间的区间范围小于90天。
	EndTimeStamp *int64 `json:"EndTimeStamp,omitnil,omitempty" name:"EndTimeStamp"`

	// 实例 ID（废弃）
	//
	// Deprecated: InstanceId is deprecated.
	InstanceId *int64 `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回数据条数，上限（废弃）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移（废弃）
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页尺寸（必填），上限 100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码（必填），从 0 开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 按手机号筛选
	Phones []*string `json:"Phones,omitnil,omitempty" name:"Phones"`

	// 按SessionId筛选
	SessionIds []*string `json:"SessionIds,omitnil,omitempty" name:"SessionIds"`
}

type DescribeTelCdrRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间戳，Unix 秒级时间戳，最大支持近180天。
	StartTimeStamp *int64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 秒级时间戳，结束时间与开始时间的区间范围小于90天。
	EndTimeStamp *int64 `json:"EndTimeStamp,omitnil,omitempty" name:"EndTimeStamp"`

	// 实例 ID（废弃）
	InstanceId *int64 `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回数据条数，上限（废弃）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移（废弃）
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页尺寸（必填），上限 100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码（必填），从 0 开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 按手机号筛选
	Phones []*string `json:"Phones,omitnil,omitempty" name:"Phones"`

	// 按SessionId筛选
	SessionIds []*string `json:"SessionIds,omitnil,omitempty" name:"SessionIds"`
}

func (r *DescribeTelCdrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelCdrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTimeStamp")
	delete(f, "EndTimeStamp")
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Phones")
	delete(f, "SessionIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTelCdrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTelCdrResponseParams struct {
	// 话单记录总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 话单记录
	//
	// Deprecated: TelCdrs is deprecated.
	TelCdrs []*TelCdrInfo `json:"TelCdrs,omitnil,omitempty" name:"TelCdrs"`

	// 话单记录
	TelCdrList []*TelCdrInfo `json:"TelCdrList,omitnil,omitempty" name:"TelCdrList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTelCdrResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTelCdrResponseParams `json:"Response"`
}

func (r *DescribeTelCdrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelCdrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTelRecordAsrRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 会话 ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type DescribeTelRecordAsrRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 会话 ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *DescribeTelRecordAsrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelRecordAsrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTelRecordAsrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTelRecordAsrResponseParams struct {
	// 录音转文本信息
	AsrDataList []*AsrData `json:"AsrDataList,omitnil,omitempty" name:"AsrDataList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTelRecordAsrResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTelRecordAsrResponseParams `json:"Response"`
}

func (r *DescribeTelRecordAsrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelRecordAsrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTelSessionRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 会话 ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type DescribeTelSessionRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 会话 ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *DescribeTelSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTelSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTelSessionResponseParams struct {
	// 会话信息
	Session *PSTNSession `json:"Session,omitnil,omitempty" name:"Session"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTelSessionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTelSessionResponseParams `json:"Response"`
}

func (r *DescribeTelSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableCCCPhoneNumberRequestParams struct {
	// 号码列表，0086开头
	PhoneNumbers []*string `json:"PhoneNumbers,omitnil,omitempty" name:"PhoneNumbers"`

	// 停用开关，0启用 1停用
	Disabled *int64 `json:"Disabled,omitnil,omitempty" name:"Disabled"`

	// TCCC 实例应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DisableCCCPhoneNumberRequest struct {
	*tchttp.BaseRequest
	
	// 号码列表，0086开头
	PhoneNumbers []*string `json:"PhoneNumbers,omitnil,omitempty" name:"PhoneNumbers"`

	// 停用开关，0启用 1停用
	Disabled *int64 `json:"Disabled,omitnil,omitempty" name:"Disabled"`

	// TCCC 实例应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

func (r *DisableCCCPhoneNumberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableCCCPhoneNumberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PhoneNumbers")
	delete(f, "Disabled")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableCCCPhoneNumberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableCCCPhoneNumberResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableCCCPhoneNumberResponse struct {
	*tchttp.BaseResponse
	Response *DisableCCCPhoneNumberResponseParams `json:"Response"`
}

func (r *DisableCCCPhoneNumberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableCCCPhoneNumberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ErrStaffItem struct {
	// 座席邮箱地址
	StaffEmail *string `json:"StaffEmail,omitnil,omitempty" name:"StaffEmail"`

	// 错误码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误描述
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type ExtensionInfo struct {
	// 实例ID
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分机全名
	FullExtensionId *string `json:"FullExtensionId,omitnil,omitempty" name:"FullExtensionId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`

	// 所属技能组列表
	SkillGroupId *string `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// 分机名称
	ExtensionName *string `json:"ExtensionName,omitnil,omitempty" name:"ExtensionName"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后修改时间
	ModifyTime *int64 `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 话机状态(0 离线、100 空闲、200忙碌）
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否注册
	Register *bool `json:"Register,omitnil,omitempty" name:"Register"`

	// 绑定座席邮箱
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`

	// 绑定座席名称
	RelationName *string `json:"RelationName,omitnil,omitempty" name:"RelationName"`
}

type Filter struct {
	// 筛选字段名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 筛选条件值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type HangUpCallRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type HangUpCallRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *HangUpCallRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HangUpCallRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "HangUpCallRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type HangUpCallResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type HangUpCallResponse struct {
	*tchttp.BaseResponse
	Response *HangUpCallResponseParams `json:"Response"`
}

func (r *HangUpCallResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HangUpCallResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IMCdrInfo struct {
	// 服务记录ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 服务时长秒数
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 结束状态
	// 0 异常结束
	// 1 正常结束
	// 3 无座席在线
	// 17 座席放弃接听
	// 100 黑名单
	// 101 座席手动转接
	// 102 IVR阶段放弃
	// 108 用户超时自动结束
	// 109 用户主动结束
	EndStatus *int64 `json:"EndStatus,omitnil,omitempty" name:"EndStatus"`

	// 用户昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 服务类型 1为全媒体，2为文本客服
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 客服ID
	StaffId *string `json:"StaffId,omitnil,omitempty" name:"StaffId"`

	// 服务时间戳
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 技能组ID
	SkillGroupId *string `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// 技能组名称
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`

	// 满意度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Satisfaction *IMSatisfaction `json:"Satisfaction,omitnil,omitempty" name:"Satisfaction"`

	// 用户ID
	ClientUserId *string `json:"ClientUserId,omitnil,omitempty" name:"ClientUserId"`
}

type IMSatisfaction struct {
	// 满意度值
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 满意度标签
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`
}

type IVRKeyPressedElement struct {
	// 命中的关键字或者按键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 按键关联的标签
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// Unix 毫秒时间戳
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 节点标签
	NodeLabel *string `json:"NodeLabel,omitnil,omitempty" name:"NodeLabel"`

	// 用户原始输入
	OriginalContent *string `json:"OriginalContent,omitnil,omitempty" name:"OriginalContent"`

	// TTS 提示音内容
	TTSPrompt *string `json:"TTSPrompt,omitnil,omitempty" name:"TTSPrompt"`
}

type Message struct {
	// 消息类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 消息内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type MessageBody struct {
	// 消息时间戳
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 发消息的用户ID
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 消息列表
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`
}

// Predefined struct for user
type ModifyCompanyApplyRequestParams struct {
	// 申请单ID(只能修改状态为“驳回”或者“待审核”的申请单)
	ApplyId *int64 `json:"ApplyId,omitnil,omitempty" name:"ApplyId"`

	// 企业资质信息
	CompanyInfo *CompanyApplyInfo `json:"CompanyInfo,omitnil,omitempty" name:"CompanyInfo"`
}

type ModifyCompanyApplyRequest struct {
	*tchttp.BaseRequest
	
	// 申请单ID(只能修改状态为“驳回”或者“待审核”的申请单)
	ApplyId *int64 `json:"ApplyId,omitnil,omitempty" name:"ApplyId"`

	// 企业资质信息
	CompanyInfo *CompanyApplyInfo `json:"CompanyInfo,omitnil,omitempty" name:"CompanyInfo"`
}

func (r *ModifyCompanyApplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCompanyApplyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplyId")
	delete(f, "CompanyInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCompanyApplyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCompanyApplyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCompanyApplyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCompanyApplyResponseParams `json:"Response"`
}

func (r *ModifyCompanyApplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCompanyApplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyExtensionRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`

	// 分机名称
	ExtensionName *string `json:"ExtensionName,omitnil,omitempty" name:"ExtensionName"`

	// 所属技能组列表
	SkillGroupIds []*int64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`

	// 绑定坐席邮箱账号
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`
}

type ModifyExtensionRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`

	// 分机名称
	ExtensionName *string `json:"ExtensionName,omitnil,omitempty" name:"ExtensionName"`

	// 所属技能组列表
	SkillGroupIds []*int64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`

	// 绑定坐席邮箱账号
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`
}

func (r *ModifyExtensionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyExtensionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "ExtensionId")
	delete(f, "ExtensionName")
	delete(f, "SkillGroupIds")
	delete(f, "Relation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyExtensionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyExtensionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyExtensionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyExtensionResponseParams `json:"Response"`
}

func (r *ModifyExtensionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyExtensionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOwnNumberApplyRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 线路相关参数
	DetailList []*OwnNumberApplyDetailItem `json:"DetailList,omitnil,omitempty" name:"DetailList"`

	// 审批单号
	ApplyId *uint64 `json:"ApplyId,omitnil,omitempty" name:"ApplyId"`

	// 送号前缀
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`
}

type ModifyOwnNumberApplyRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 线路相关参数
	DetailList []*OwnNumberApplyDetailItem `json:"DetailList,omitnil,omitempty" name:"DetailList"`

	// 审批单号
	ApplyId *uint64 `json:"ApplyId,omitnil,omitempty" name:"ApplyId"`

	// 送号前缀
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`
}

func (r *ModifyOwnNumberApplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOwnNumberApplyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "DetailList")
	delete(f, "ApplyId")
	delete(f, "Prefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOwnNumberApplyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOwnNumberApplyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyOwnNumberApplyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyOwnNumberApplyResponseParams `json:"Response"`
}

func (r *ModifyOwnNumberApplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOwnNumberApplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStaffPasswordRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 座席邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 设置的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type ModifyStaffPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 座席邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 设置的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *ModifyStaffPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStaffPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Email")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStaffPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStaffPasswordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyStaffPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStaffPasswordResponseParams `json:"Response"`
}

func (r *ModifyStaffPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStaffPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStaffRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 座席账户
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 座席名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 座席手机号（带0086前缀,示例：008618011111111）
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 座席昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 座席工号
	StaffNo *string `json:"StaffNo,omitnil,omitempty" name:"StaffNo"`

	// 绑定技能组ID列表
	SkillGroupIds []*int64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`

	// 是否开启手机外呼开关
	UseMobileCallOut *bool `json:"UseMobileCallOut,omitnil,omitempty" name:"UseMobileCallOut"`

	// 手机接听模式 0 - 关闭 | 1 - 仅离线 | 2 - 始终
	UseMobileAccept *int64 `json:"UseMobileAccept,omitnil,omitempty" name:"UseMobileAccept"`

	// 座席分机号（1 到 8 打头，4 - 6 位）
	ExtensionNumber *string `json:"ExtensionNumber,omitnil,omitempty" name:"ExtensionNumber"`
}

type ModifyStaffRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 座席账户
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 座席名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 座席手机号（带0086前缀,示例：008618011111111）
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 座席昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 座席工号
	StaffNo *string `json:"StaffNo,omitnil,omitempty" name:"StaffNo"`

	// 绑定技能组ID列表
	SkillGroupIds []*int64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`

	// 是否开启手机外呼开关
	UseMobileCallOut *bool `json:"UseMobileCallOut,omitnil,omitempty" name:"UseMobileCallOut"`

	// 手机接听模式 0 - 关闭 | 1 - 仅离线 | 2 - 始终
	UseMobileAccept *int64 `json:"UseMobileAccept,omitnil,omitempty" name:"UseMobileAccept"`

	// 座席分机号（1 到 8 打头，4 - 6 位）
	ExtensionNumber *string `json:"ExtensionNumber,omitnil,omitempty" name:"ExtensionNumber"`
}

func (r *ModifyStaffRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStaffRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Email")
	delete(f, "Name")
	delete(f, "Phone")
	delete(f, "Nick")
	delete(f, "StaffNo")
	delete(f, "SkillGroupIds")
	delete(f, "UseMobileCallOut")
	delete(f, "UseMobileAccept")
	delete(f, "ExtensionNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStaffRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStaffResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyStaffResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStaffResponseParams `json:"Response"`
}

func (r *ModifyStaffResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStaffResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NumberInfo struct {
	// 号码
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// 绑定的外呼技能组
	CallOutSkillGroupIds []*uint64 `json:"CallOutSkillGroupIds,omitnil,omitempty" name:"CallOutSkillGroupIds"`

	// 号码状态，1-正常，2-欠费停用，4-管理员停用，5-违规停用
	State *int64 `json:"State,omitnil,omitempty" name:"State"`
}

type OwnNumberApplyDetailItem struct {
	// 号码类型：0-呼入|1-呼出|2-呼入呼出
	CallType *int64 `json:"CallType,omitnil,omitempty" name:"CallType"`

	// 线路号码
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 最大并发呼叫数
	MaxCallCount *int64 `json:"MaxCallCount,omitnil,omitempty" name:"MaxCallCount"`

	// 每秒最大并发数
	MaxCallPSec *int64 `json:"MaxCallPSec,omitnil,omitempty" name:"MaxCallPSec"`

	// 呼出被叫格式，使用 {+E.164} 或 {E.164}, 
	OutboundCalleeFormat *string `json:"OutboundCalleeFormat,omitnil,omitempty" name:"OutboundCalleeFormat"`
}

type PSTNSession struct {
	// 会话 ID
	SessionID *string `json:"SessionID,omitnil,omitempty" name:"SessionID"`

	// 会话临时房间 ID
	RoomID *string `json:"RoomID,omitnil,omitempty" name:"RoomID"`

	// 主叫
	Caller *string `json:"Caller,omitnil,omitempty" name:"Caller"`

	// 被叫
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// 开始时间，Unix 时间戳
	StartTimestamp *int64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// 振铃时间，Unix 时间戳
	RingTimestamp *int64 `json:"RingTimestamp,omitnil,omitempty" name:"RingTimestamp"`

	// 接听时间，Unix 时间戳
	AcceptTimestamp *int64 `json:"AcceptTimestamp,omitnil,omitempty" name:"AcceptTimestamp"`

	// 座席邮箱
	StaffEmail *string `json:"StaffEmail,omitnil,omitempty" name:"StaffEmail"`

	// 座席工号
	StaffNumber *string `json:"StaffNumber,omitnil,omitempty" name:"StaffNumber"`

	// 会话状态
	// ringing 振铃中
	// seatJoining  等待座席接听
	// inProgress 进行中
	// finished 已完成
	SessionStatus *string `json:"SessionStatus,omitnil,omitempty" name:"SessionStatus"`

	// 会话呼叫方向， 0 呼入 | 1 - 呼出
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 转外线使用的号码（转外线主叫）
	OutBoundCaller *string `json:"OutBoundCaller,omitnil,omitempty" name:"OutBoundCaller"`

	// 转外线被叫
	OutBoundCallee *string `json:"OutBoundCallee,omitnil,omitempty" name:"OutBoundCallee"`

	// 主叫号码保护ID，开启号码保护映射功能时有效，且Caller字段置空
	ProtectedCaller *string `json:"ProtectedCaller,omitnil,omitempty" name:"ProtectedCaller"`

	// 被叫号码保护ID，开启号码保护映射功能时有效，且Callee字段置空
	ProtectedCallee *string `json:"ProtectedCallee,omitnil,omitempty" name:"ProtectedCallee"`
}

type PSTNSessionInfo struct {
	// 会话 ID
	SessionID *string `json:"SessionID,omitnil,omitempty" name:"SessionID"`

	// 会话临时房间 ID
	RoomID *string `json:"RoomID,omitnil,omitempty" name:"RoomID"`

	// 主叫
	Caller *string `json:"Caller,omitnil,omitempty" name:"Caller"`

	// 被叫
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// 开始时间，Unix 时间戳
	StartTimestamp *string `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// 接听时间，Unix 时间戳
	AcceptTimestamp *string `json:"AcceptTimestamp,omitnil,omitempty" name:"AcceptTimestamp"`

	// 座席邮箱
	StaffEmail *string `json:"StaffEmail,omitnil,omitempty" name:"StaffEmail"`

	// 座席工号
	StaffNumber *string `json:"StaffNumber,omitnil,omitempty" name:"StaffNumber"`

	// 座席状态 inProgress 进行中
	SessionStatus *string `json:"SessionStatus,omitnil,omitempty" name:"SessionStatus"`

	// 会话呼叫方向， 0 呼入 | 1 - 呼出
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 振铃时间，Unix 时间戳
	RingTimestamp *int64 `json:"RingTimestamp,omitnil,omitempty" name:"RingTimestamp"`

	// 主叫号码保护ID，开启号码保护映射功能时有效，且Caller字段置空
	ProtectedCaller *string `json:"ProtectedCaller,omitnil,omitempty" name:"ProtectedCaller"`

	// 被叫号码保护ID，开启号码保护映射功能时有效，且Callee字段置空
	ProtectedCallee *string `json:"ProtectedCallee,omitnil,omitempty" name:"ProtectedCallee"`
}

type PackageBuyInfo struct {
	// 套餐包Id
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 套餐包类型，0-外呼套餐包|1-400呼入套餐包
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 套餐包总量
	CapacitySize *int64 `json:"CapacitySize,omitnil,omitempty" name:"CapacitySize"`

	// 套餐包剩余量
	CapacityRemain *int64 `json:"CapacityRemain,omitnil,omitempty" name:"CapacityRemain"`

	// 购买时间戳
	BuyTime *int64 `json:"BuyTime,omitnil,omitempty" name:"BuyTime"`

	// 截止时间戳
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

// Predefined struct for user
type PausePredictiveDialingCampaignRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务 ID
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

type PausePredictiveDialingCampaignRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务 ID
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

func (r *PausePredictiveDialingCampaignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PausePredictiveDialingCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CampaignId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PausePredictiveDialingCampaignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PausePredictiveDialingCampaignResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PausePredictiveDialingCampaignResponse struct {
	*tchttp.BaseResponse
	Response *PausePredictiveDialingCampaignResponseParams `json:"Response"`
}

func (r *PausePredictiveDialingCampaignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PausePredictiveDialingCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PhoneNumBuyInfo struct {
	// 电话号码
	PhoneNum *string `json:"PhoneNum,omitnil,omitempty" name:"PhoneNum"`

	// 号码类型，0-固话|1-虚商号码|2-运营商号码|3-400号码
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 号码呼叫类型，1-呼入|2-呼出|3-呼入呼出
	CallType *int64 `json:"CallType,omitnil,omitempty" name:"CallType"`

	// 购买时间戳
	BuyTime *int64 `json:"BuyTime,omitnil,omitempty" name:"BuyTime"`

	// 截止时间戳
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 号码状态，1正常|2欠费停用|4管理员停用|5违规停用
	State *int64 `json:"State,omitnil,omitempty" name:"State"`
}

// Predefined struct for user
type ResetExtensionPasswordRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`
}

type ResetExtensionPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`
}

func (r *ResetExtensionPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetExtensionPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "ExtensionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetExtensionPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetExtensionPasswordResponseParams struct {
	// 重置后密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetExtensionPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetExtensionPasswordResponseParams `json:"Response"`
}

func (r *ResetExtensionPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetExtensionPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumePredictiveDialingCampaignRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务 ID
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

type ResumePredictiveDialingCampaignRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务 ID
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

func (r *ResumePredictiveDialingCampaignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumePredictiveDialingCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CampaignId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumePredictiveDialingCampaignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumePredictiveDialingCampaignResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumePredictiveDialingCampaignResponse struct {
	*tchttp.BaseResponse
	Response *ResumePredictiveDialingCampaignResponseParams `json:"Response"`
}

func (r *ResumePredictiveDialingCampaignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumePredictiveDialingCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SdkAppIdBuyInfo struct {
	// 应用ID
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 座席购买数（还在有效期内）
	StaffBuyNum *int64 `json:"StaffBuyNum,omitnil,omitempty" name:"StaffBuyNum"`

	// 座席购买列表 （还在有效期内）
	StaffBuyList []*StaffBuyInfo `json:"StaffBuyList,omitnil,omitempty" name:"StaffBuyList"`

	// 号码购买列表
	PhoneNumBuyList []*PhoneNumBuyInfo `json:"PhoneNumBuyList,omitnil,omitempty" name:"PhoneNumBuyList"`

	// 办公电话购买数（还在有效期内）
	SipBuyNum *int64 `json:"SipBuyNum,omitnil,omitempty" name:"SipBuyNum"`
}

type SeatUserInfo struct {
	// 座席名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 座席邮箱
	Mail *string `json:"Mail,omitnil,omitempty" name:"Mail"`

	// 工号
	StaffNumber *string `json:"StaffNumber,omitnil,omitempty" name:"StaffNumber"`

	// 座席电话号码（带0086前缀）
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 座席昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 座席关联的技能组列表
	SkillGroupNameList []*string `json:"SkillGroupNameList,omitnil,omitempty" name:"SkillGroupNameList"`

	// 1:管理员
	// 2:质检员
	// 3:普通座席
	// else:自定义角色ID
	Role *int64 `json:"Role,omitnil,omitempty" name:"Role"`

	// 座席分机号（1 到 8 打头，4 - 6 位）
	ExtensionNumber *string `json:"ExtensionNumber,omitnil,omitempty" name:"ExtensionNumber"`
}

type ServeParticipant struct {
	// 座席邮箱
	Mail *string `json:"Mail,omitnil,omitempty" name:"Mail"`

	// 座席电话
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 振铃时间戳，Unix 秒级时间戳
	RingTimestamp *int64 `json:"RingTimestamp,omitnil,omitempty" name:"RingTimestamp"`

	// 接听时间戳，Unix 秒级时间戳
	AcceptTimestamp *int64 `json:"AcceptTimestamp,omitnil,omitempty" name:"AcceptTimestamp"`

	// 结束时间戳，Unix 秒级时间戳
	EndedTimestamp *int64 `json:"EndedTimestamp,omitnil,omitempty" name:"EndedTimestamp"`

	// 录音 ID，能够索引到座席侧的录音
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 参与者类型，"staffSeat", "outboundSeat", "staffPhoneSeat"
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 转接来源座席信息
	TransferFrom *string `json:"TransferFrom,omitnil,omitempty" name:"TransferFrom"`

	// 转接来源参与者类型，取值与 Type 一致
	TransferFromType *string `json:"TransferFromType,omitnil,omitempty" name:"TransferFromType"`

	// 转接去向座席信息
	TransferTo *string `json:"TransferTo,omitnil,omitempty" name:"TransferTo"`

	// 转接去向参与者类型，取值与 Type 一致
	TransferToType *string `json:"TransferToType,omitnil,omitempty" name:"TransferToType"`

	// 技能组 ID
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// 结束状态
	EndStatusString *string `json:"EndStatusString,omitnil,omitempty" name:"EndStatusString"`

	// 录音 URL
	RecordURL *string `json:"RecordURL,omitnil,omitempty" name:"RecordURL"`

	// 参与者序号，从 0 开始
	Sequence *int64 `json:"Sequence,omitnil,omitempty" name:"Sequence"`

	// 开始时间戳，Unix 秒级时间戳
	StartTimestamp *int64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// 技能组名称
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`

	// 录音转存第三方COS地址
	CustomRecordURL *string `json:"CustomRecordURL,omitnil,omitempty" name:"CustomRecordURL"`
}

type SkillGroupInfoItem struct {
	// 技能组ID
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// 技能组名称
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`

	// （废弃）类型：IM、TEL、ALL（全媒体）
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 会话分配策略
	RoutePolicy *string `json:"RoutePolicy,omitnil,omitempty" name:"RoutePolicy"`

	// 会话分配是否优先上次服务座席
	UsingLastSeat *int64 `json:"UsingLastSeat,omitnil,omitempty" name:"UsingLastSeat"`

	// 单客服最大并发数（电话类型默认1）
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`

	// 最后修改时间
	LastModifyTimestamp *int64 `json:"LastModifyTimestamp,omitnil,omitempty" name:"LastModifyTimestamp"`

	// 技能组类型0-电话，1-在线，3-音频，4-视频	
	SkillGroupType *int64 `json:"SkillGroupType,omitnil,omitempty" name:"SkillGroupType"`

	// 技能组内线号码
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`
}

type SkillGroupItem struct {
	// 技能组ID
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// 技能组名称
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`

	// 优先级
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 类型：IM、TEL、ALL（全媒体）
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type StaffBuyInfo struct {
	// 购买座席数量
	Num *int64 `json:"Num,omitnil,omitempty" name:"Num"`

	// 购买时间戳
	BuyTime *int64 `json:"BuyTime,omitnil,omitempty" name:"BuyTime"`

	// 截止时间戳
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 购买办公电话数量
	SipNum *int64 `json:"SipNum,omitnil,omitempty" name:"SipNum"`
}

type StaffInfo struct {
	// 座席名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 座席邮箱
	Mail *string `json:"Mail,omitnil,omitempty" name:"Mail"`

	// 座席电话号码
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 座席昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 座席工号
	StaffNumber *string `json:"StaffNumber,omitnil,omitempty" name:"StaffNumber"`

	// 用户角色id
	RoleId *uint64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 所属技能组列表
	SkillGroupList []*SkillGroupItem `json:"SkillGroupList,omitnil,omitempty" name:"SkillGroupList"`

	// 最后修改时间
	LastModifyTimestamp *int64 `json:"LastModifyTimestamp,omitnil,omitempty" name:"LastModifyTimestamp"`

	// 座席分机号（1 到 8 打头，4 - 6 位）
	ExtensionNumber *string `json:"ExtensionNumber,omitnil,omitempty" name:"ExtensionNumber"`
}

type StaffSkillGroupList struct {
	// 技能组ID
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// 座席在技能组中的优先级（1为最高，5最低，默认3）
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

type StaffStatusExtra struct {
	// im - 文本 | tel - 电话 | all - 全媒体
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// in - 呼入 | out - 呼出
	Direct *string `json:"Direct,omitnil,omitempty" name:"Direct"`
}

type StaffStatusMetrics struct {
	// 座席邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 座席状态 free 示闲 | busy 忙碌 | rest 小休 | notReady 示忙 | afterCallWork 话后调整 | offline 离线
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 座席状态补充信息
	StatusExtra *StaffStatusExtra `json:"StatusExtra,omitnil,omitempty" name:"StatusExtra"`

	// 当天在线总时长
	OnlineDuration *int64 `json:"OnlineDuration,omitnil,omitempty" name:"OnlineDuration"`

	// 当天示闲总时长
	FreeDuration *int64 `json:"FreeDuration,omitnil,omitempty" name:"FreeDuration"`

	// 当天忙碌总时长
	BusyDuration *int64 `json:"BusyDuration,omitnil,omitempty" name:"BusyDuration"`

	// 当天示忙总时长
	NotReadyDuration *int64 `json:"NotReadyDuration,omitnil,omitempty" name:"NotReadyDuration"`

	// 当天小休总时长
	RestDuration *int64 `json:"RestDuration,omitnil,omitempty" name:"RestDuration"`

	// 当天话后调整总时长
	AfterCallWorkDuration *int64 `json:"AfterCallWorkDuration,omitnil,omitempty" name:"AfterCallWorkDuration"`

	// 小休原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 是否预约小休
	ReserveRest *bool `json:"ReserveRest,omitnil,omitempty" name:"ReserveRest"`

	// 是否预约示忙
	ReserveNotReady *bool `json:"ReserveNotReady,omitnil,omitempty" name:"ReserveNotReady"`

	// 手机接听模式： 0 - 关闭 | 1 - 仅离线 | 2- 始终
	UseMobileAccept *int64 `json:"UseMobileAccept,omitnil,omitempty" name:"UseMobileAccept"`

	// 手机外呼开关
	UseMobileCallOut *bool `json:"UseMobileCallOut,omitnil,omitempty" name:"UseMobileCallOut"`

	// 最近一次上线时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOnlineTimestamp *int64 `json:"LastOnlineTimestamp,omitnil,omitempty" name:"LastOnlineTimestamp"`

	// 最近一次状态时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastStatusTimestamp *int64 `json:"LastStatusTimestamp,omitnil,omitempty" name:"LastStatusTimestamp"`
}

// Predefined struct for user
type StopAutoCalloutTaskRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务Id
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type StopAutoCalloutTaskRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务Id
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *StopAutoCalloutTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopAutoCalloutTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopAutoCalloutTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopAutoCalloutTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopAutoCalloutTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopAutoCalloutTaskResponseParams `json:"Response"`
}

func (r *StopAutoCalloutTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopAutoCalloutTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TelCdrInfo struct {
	// 主叫号码
	Caller *string `json:"Caller,omitnil,omitempty" name:"Caller"`

	// 被叫号码
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// 呼叫发起时间戳，Unix 时间戳
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 呼入呼出方向 0 呼入 1 呼出
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 通话时长
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 录音信息
	RecordURL *string `json:"RecordURL,omitnil,omitempty" name:"RecordURL"`

	// 录音 ID
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 座席信息
	SeatUser *SeatUserInfo `json:"SeatUser,omitnil,omitempty" name:"SeatUser"`

	// EndStatus与EndStatusString一一对应，具体枚举如下：
	// 
	// **场景	         EndStatus	EndStatusString	状态说明**
	// 
	// 电话呼入&呼出	1	        ok	                        正常结束
	// 
	// 电话呼入&呼出	0	        error	                系统错误
	// 
	// 电话呼入	             102	        ivrGiveUp	        IVR 期间用户放弃
	// 
	// 电话呼入	             103	        waitingGiveUp	       会话排队期间用户放弃
	// 
	// 电话呼入	             104	        ringingGiveUp	       会话振铃期间用户放弃
	// 
	// 电话呼入	             105	        noSeatOnline	       无座席在线
	// 
	// 电话呼入              106	       notWorkTime	       非工作时间   
	// 
	// 电话呼入	            107	       ivrEnd	               IVR 后直接结束
	// 
	// 电话呼入	            100	      blackList 呼入黑名单 
	// 
	// 电话呼出               2	              unconnected	未接通
	// 
	// 电话呼出             108	        restrictedCallee	被叫因高风险受限
	// 
	// 电话呼出             109	        tooManyRequest	    超频
	// 
	// 电话呼出             110	        restrictedArea	    外呼盲区
	// 
	// 电话呼出             111	        restrictedTime	外呼时间限制
	//                          
	// 电话呼出             201            unknown	未知状态
	// 
	// 电话呼出             202            notAnswer	未接听
	// 
	// 电话呼出            203	    userReject	拒接挂断
	// 
	// 电话呼出	          204	    powerOff	关机
	// 
	// 电话呼出           205            numberNotExist	空号
	// 
	// 电话呼出	         206	           busy	通话中
	// 
	// 电话呼出   	        207	           outOfCredit	欠费
	// 
	// 电话呼出	         208	           operatorError	运营商线路异常
	// 
	// 电话呼出         	209	           callerCancel	主叫取消
	// 
	// 电话呼出	        210	           notInService	不在服务区
	// 
	// 电话呼入&呼出	211    clientError    客户端错误
	EndStatus *int64 `json:"EndStatus,omitnil,omitempty" name:"EndStatus"`

	// 技能组名称
	SkillGroup *string `json:"SkillGroup,omitnil,omitempty" name:"SkillGroup"`

	// 主叫归属地
	CallerLocation *string `json:"CallerLocation,omitnil,omitempty" name:"CallerLocation"`

	// IVR 阶段耗时
	IVRDuration *int64 `json:"IVRDuration,omitnil,omitempty" name:"IVRDuration"`

	// 振铃时间戳，UNIX 秒级时间戳
	RingTimestamp *int64 `json:"RingTimestamp,omitnil,omitempty" name:"RingTimestamp"`

	// 接听时间戳，UNIX 秒级时间戳
	AcceptTimestamp *int64 `json:"AcceptTimestamp,omitnil,omitempty" name:"AcceptTimestamp"`

	// 结束时间戳，UNIX 秒级时间戳
	EndedTimestamp *int64 `json:"EndedTimestamp,omitnil,omitempty" name:"EndedTimestamp"`

	// IVR 按键信息 ，e.g. ["1","2","3"]
	IVRKeyPressed []*string `json:"IVRKeyPressed,omitnil,omitempty" name:"IVRKeyPressed"`

	// 挂机方 seat 座席 user 用户 system 系统
	HungUpSide *string `json:"HungUpSide,omitnil,omitempty" name:"HungUpSide"`

	// 服务参与者列表
	ServeParticipants []*ServeParticipant `json:"ServeParticipants,omitnil,omitempty" name:"ServeParticipants"`

	// 技能组ID
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// EndStatus与EndStatusString一一对应，具体枚举如下：
	// 
	// **场景	         EndStatus	EndStatusString	状态说明**
	// 
	// 电话呼入&呼出	1	        ok	                        正常结束
	// 
	// 电话呼入&呼出	0	        error	                系统错误
	// 
	// 电话呼入	             102	        ivrGiveUp	        IVR 期间用户放弃
	// 
	// 电话呼入	             103	        waitingGiveUp	       会话排队期间用户放弃
	// 
	// 电话呼入	             104	        ringingGiveUp	       会话振铃期间用户放弃
	// 
	// 电话呼入	             105	        noSeatOnline	       无座席在线
	// 
	// 电话呼入              106	       notWorkTime	       非工作时间   
	// 
	// 电话呼入	            107	       ivrEnd	               IVR 后直接结束
	// 
	// 电话呼入	            100	      blackList 呼入黑名单 
	// 
	// 电话呼出               2	              unconnected	未接通
	// 
	// 电话呼出             108	        restrictedCallee	被叫因高风险受限
	// 
	// 电话呼出             109	        tooManyRequest	    超频
	// 
	// 电话呼出             110	        restrictedArea	    外呼盲区
	// 
	// 电话呼出             111	        restrictedTime	外呼时间限制
	//                          
	// 电话呼出             201            unknown	未知状态
	// 
	// 电话呼出             202            notAnswer	未接听
	// 
	// 电话呼出            203	    userReject	拒接挂断
	// 
	// 电话呼出	          204	    powerOff	关机
	// 
	// 电话呼出           205            numberNotExist	空号
	// 
	// 电话呼出	         206	           busy	通话中
	// 
	// 电话呼出   	        207	           outOfCredit	欠费
	// 
	// 电话呼出	         208	           operatorError	运营商线路异常
	// 
	// 电话呼出         	209	           callerCancel	主叫取消
	// 
	// 电话呼出	        210	           notInService	不在服务区
	// 
	// 电话呼入&呼出	211    clientError    客户端错误
	EndStatusString *string `json:"EndStatusString,omitnil,omitempty" name:"EndStatusString"`

	// 会话开始时间戳，UNIX 秒级时间戳
	StartTimestamp *int64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// 进入排队时间，Unix 秒级时间戳
	QueuedTimestamp *int64 `json:"QueuedTimestamp,omitnil,omitempty" name:"QueuedTimestamp"`

	// 后置IVR按键信息（e.g. [{"Key":"1","Label":"非常满意"}]）
	PostIVRKeyPressed []*IVRKeyPressedElement `json:"PostIVRKeyPressed,omitnil,omitempty" name:"PostIVRKeyPressed"`

	// 排队技能组Id
	QueuedSkillGroupId *int64 `json:"QueuedSkillGroupId,omitnil,omitempty" name:"QueuedSkillGroupId"`

	// 会话 ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 主叫号码保护ID，开启号码保护映射功能时有效，且Caller字段置空
	ProtectedCaller *string `json:"ProtectedCaller,omitnil,omitempty" name:"ProtectedCaller"`

	// 被叫号码保护ID，开启号码保护映射功能时有效，且Callee字段置空
	ProtectedCallee *string `json:"ProtectedCallee,omitnil,omitempty" name:"ProtectedCallee"`

	// 客户自定义数据（User-to-User Interface）
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Uui is deprecated.
	Uui *string `json:"Uui,omitnil,omitempty" name:"Uui"`

	// 客户自定义数据（User-to-User Interface）
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`

	// IVR按键信息（e.g. [{"Key":"1","Label":"非常满意"}]）
	IVRKeyPressedEx []*IVRKeyPressedElement `json:"IVRKeyPressedEx,omitnil,omitempty" name:"IVRKeyPressedEx"`

	// 获取录音ASR文本信息地址
	AsrUrl *string `json:"AsrUrl,omitnil,omitempty" name:"AsrUrl"`

	// AsrUrl的状态：Complete
	// 已完成;
	// Processing
	// 正在生成中;
	// NotExists
	// 无记录(未开启生成离线asr或者无套餐包)
	AsrStatus *string `json:"AsrStatus,omitnil,omitempty" name:"AsrStatus"`

	// 录音转存第三方COS地址
	CustomRecordURL *string `json:"CustomRecordURL,omitnil,omitempty" name:"CustomRecordURL"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 排队技能组名称
	QueuedSkillGroupName *string `json:"QueuedSkillGroupName,omitnil,omitempty" name:"QueuedSkillGroupName"`

	// 通话中语音留言录音URL
	VoicemailRecordURL []*string `json:"VoicemailRecordURL,omitnil,omitempty" name:"VoicemailRecordURL"`

	// 通话中语音留言ASR文本信息地址
	VoicemailAsrURL []*string `json:"VoicemailAsrURL,omitnil,omitempty" name:"VoicemailAsrURL"`
}

// Predefined struct for user
type UnbindNumberCallOutSkillGroupRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 待解绑的号码
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// 待解绑的技能组Id列表
	SkillGroupIds []*uint64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`
}

type UnbindNumberCallOutSkillGroupRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 待解绑的号码
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// 待解绑的技能组Id列表
	SkillGroupIds []*uint64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`
}

func (r *UnbindNumberCallOutSkillGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindNumberCallOutSkillGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Number")
	delete(f, "SkillGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindNumberCallOutSkillGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindNumberCallOutSkillGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindNumberCallOutSkillGroupResponse struct {
	*tchttp.BaseResponse
	Response *UnbindNumberCallOutSkillGroupResponseParams `json:"Response"`
}

func (r *UnbindNumberCallOutSkillGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindNumberCallOutSkillGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindStaffSkillGroupListRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 客服邮箱
	StaffEmail *string `json:"StaffEmail,omitnil,omitempty" name:"StaffEmail"`

	// 解绑技能组列表
	SkillGroupList []*int64 `json:"SkillGroupList,omitnil,omitempty" name:"SkillGroupList"`
}

type UnbindStaffSkillGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 客服邮箱
	StaffEmail *string `json:"StaffEmail,omitnil,omitempty" name:"StaffEmail"`

	// 解绑技能组列表
	SkillGroupList []*int64 `json:"SkillGroupList,omitnil,omitempty" name:"SkillGroupList"`
}

func (r *UnbindStaffSkillGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindStaffSkillGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StaffEmail")
	delete(f, "SkillGroupList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindStaffSkillGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindStaffSkillGroupListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindStaffSkillGroupListResponse struct {
	*tchttp.BaseResponse
	Response *UnbindStaffSkillGroupListResponseParams `json:"Response"`
}

func (r *UnbindStaffSkillGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindStaffSkillGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCCCSkillGroupRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 技能组ID
	SkillGroupID *int64 `json:"SkillGroupID,omitnil,omitempty" name:"SkillGroupID"`

	// 修改后的技能组名字
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`

	// 修改后的最大并发数,同振最大为2
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`

	// true同振，false顺振
	RingAll *bool `json:"RingAll,omitnil,omitempty" name:"RingAll"`
}

type UpdateCCCSkillGroupRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 技能组ID
	SkillGroupID *int64 `json:"SkillGroupID,omitnil,omitempty" name:"SkillGroupID"`

	// 修改后的技能组名字
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`

	// 修改后的最大并发数,同振最大为2
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`

	// true同振，false顺振
	RingAll *bool `json:"RingAll,omitnil,omitempty" name:"RingAll"`
}

func (r *UpdateCCCSkillGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCCCSkillGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SkillGroupID")
	delete(f, "SkillGroupName")
	delete(f, "MaxConcurrency")
	delete(f, "RingAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCCCSkillGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCCCSkillGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCCCSkillGroupResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCCCSkillGroupResponseParams `json:"Response"`
}

func (r *UpdateCCCSkillGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCCCSkillGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePredictiveDialingCampaignRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 生成的任务 ID
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`

	// 任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 被叫列表，支持 E.164 或不带国家码形式的号码
	Callees []*string `json:"Callees,omitnil,omitempty" name:"Callees"`

	// 主叫列表，使用管理端展示的号码格式
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 被叫呼叫顺序 0 随机 1 顺序
	CallOrder *int64 `json:"CallOrder,omitnil,omitempty" name:"CallOrder"`

	// 使用的座席技能组 ID
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// 相同应用内多个任务运行优先级，从高到底 1 - 5
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 预期呼损率，百分比，5 - 50	
	ExpectedAbandonRate *int64 `json:"ExpectedAbandonRate,omitnil,omitempty" name:"ExpectedAbandonRate"`

	// 呼叫重试间隔时间，单位秒，60 - 86400
	RetryInterval *int64 `json:"RetryInterval,omitnil,omitempty" name:"RetryInterval"`

	// 任务启动时间，Unix 时间戳，到此时间后会自动启动任务
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间，Unix 时间戳，到此时间后会自动终止任务
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指定的 IVR ID
	IVRId *int64 `json:"IVRId,omitnil,omitempty" name:"IVRId"`

	// 呼叫重试次数，0 - 2
	RetryTimes *int64 `json:"RetryTimes,omitnil,omitempty" name:"RetryTimes"`

	// 自定义变量
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`

	// 	UUI
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`

	// 被叫属性
	CalleeAttributes []*CalleeAttribute `json:"CalleeAttributes,omitnil,omitempty" name:"CalleeAttributes"`
}

type UpdatePredictiveDialingCampaignRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 生成的任务 ID
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`

	// 任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 被叫列表，支持 E.164 或不带国家码形式的号码
	Callees []*string `json:"Callees,omitnil,omitempty" name:"Callees"`

	// 主叫列表，使用管理端展示的号码格式
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// 被叫呼叫顺序 0 随机 1 顺序
	CallOrder *int64 `json:"CallOrder,omitnil,omitempty" name:"CallOrder"`

	// 使用的座席技能组 ID
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// 相同应用内多个任务运行优先级，从高到底 1 - 5
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 预期呼损率，百分比，5 - 50	
	ExpectedAbandonRate *int64 `json:"ExpectedAbandonRate,omitnil,omitempty" name:"ExpectedAbandonRate"`

	// 呼叫重试间隔时间，单位秒，60 - 86400
	RetryInterval *int64 `json:"RetryInterval,omitnil,omitempty" name:"RetryInterval"`

	// 任务启动时间，Unix 时间戳，到此时间后会自动启动任务
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间，Unix 时间戳，到此时间后会自动终止任务
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指定的 IVR ID
	IVRId *int64 `json:"IVRId,omitnil,omitempty" name:"IVRId"`

	// 呼叫重试次数，0 - 2
	RetryTimes *int64 `json:"RetryTimes,omitnil,omitempty" name:"RetryTimes"`

	// 自定义变量
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`

	// 	UUI
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`

	// 被叫属性
	CalleeAttributes []*CalleeAttribute `json:"CalleeAttributes,omitnil,omitempty" name:"CalleeAttributes"`
}

func (r *UpdatePredictiveDialingCampaignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePredictiveDialingCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CampaignId")
	delete(f, "Name")
	delete(f, "Callees")
	delete(f, "Callers")
	delete(f, "CallOrder")
	delete(f, "SkillGroupId")
	delete(f, "Priority")
	delete(f, "ExpectedAbandonRate")
	delete(f, "RetryInterval")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "IVRId")
	delete(f, "RetryTimes")
	delete(f, "Variables")
	delete(f, "UUI")
	delete(f, "CalleeAttributes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePredictiveDialingCampaignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePredictiveDialingCampaignResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdatePredictiveDialingCampaignResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePredictiveDialingCampaignResponseParams `json:"Response"`
}

func (r *UpdatePredictiveDialingCampaignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePredictiveDialingCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadAudioInfo struct {
	// 文件别名（可重复）
	CustomFileName *string `json:"CustomFileName,omitnil,omitempty" name:"CustomFileName"`

	// 音频文件链接。(支持mp3和wav格式，文件不超过5MB)
	AudioUrl *string `json:"AudioUrl,omitnil,omitempty" name:"AudioUrl"`
}

type UploadIvrAudioFailedInfo struct {
	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 失败原因
	FailedMsg *string `json:"FailedMsg,omitnil,omitempty" name:"FailedMsg"`
}

// Predefined struct for user
type UploadIvrAudioRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 音频文件列表
	AudioList []*UploadAudioInfo `json:"AudioList,omitnil,omitempty" name:"AudioList"`
}

type UploadIvrAudioRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 音频文件列表
	AudioList []*UploadAudioInfo `json:"AudioList,omitnil,omitempty" name:"AudioList"`
}

func (r *UploadIvrAudioRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadIvrAudioRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "AudioList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadIvrAudioRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadIvrAudioResponseParams struct {
	// 上传失败的文件列表
	FailedFileList []*UploadIvrAudioFailedInfo `json:"FailedFileList,omitnil,omitempty" name:"FailedFileList"`

	// 上传成功文件列表
	SuccessFileList []*AudioFileInfo `json:"SuccessFileList,omitnil,omitempty" name:"SuccessFileList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadIvrAudioResponse struct {
	*tchttp.BaseResponse
	Response *UploadIvrAudioResponseParams `json:"Response"`
}

func (r *UploadIvrAudioResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadIvrAudioResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Variable struct {
	// 变量名
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 变量值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}