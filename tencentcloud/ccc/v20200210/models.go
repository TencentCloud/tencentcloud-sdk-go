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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ActiveCarrierPrivilegeNumber struct {
	// 实例Id
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 主叫号码
	Caller *string `json:"Caller,omitempty" name:"Caller"`

	// 被叫号码
	Callee *string `json:"Callee,omitempty" name:"Callee"`

	// 生效unix时间戳(秒)
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type AutoCalloutTaskCalleeInfo struct {
	// 被叫号码
	Callee *string `json:"Callee,omitempty" name:"Callee"`

	// 呼叫状态 0初始 1已接听 2未接听 3呼叫中 4待重试
	State *uint64 `json:"State,omitempty" name:"State"`

	// 会话ID列表
	Sessions []*string `json:"Sessions,omitempty" name:"Sessions"`
}

type AutoCalloutTaskInfo struct {
	// 任务名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 被叫数量
	CalleeCount *uint64 `json:"CalleeCount,omitempty" name:"CalleeCount"`

	// 主叫号码列表
	Callers []*string `json:"Callers,omitempty" name:"Callers"`

	// 起始时间戳
	NotBefore *int64 `json:"NotBefore,omitempty" name:"NotBefore"`

	// 结束时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotAfter *int64 `json:"NotAfter,omitempty" name:"NotAfter"`

	// 任务使用的IvrId
	IvrId *uint64 `json:"IvrId,omitempty" name:"IvrId"`

	// 任务状态0初始 1运行中 2已完成 3结束中 4已结束
	State *uint64 `json:"State,omitempty" name:"State"`

	// 任务Id
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

// Predefined struct for user
type BindNumberCallOutSkillGroupRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 待绑定的号码
	Number *string `json:"Number,omitempty" name:"Number"`

	// 待绑定的技能组Id列表
	SkillGroupIds []*uint64 `json:"SkillGroupIds,omitempty" name:"SkillGroupIds"`
}

type BindNumberCallOutSkillGroupRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 待绑定的号码
	Number *string `json:"Number,omitempty" name:"Number"`

	// 待绑定的技能组Id列表
	SkillGroupIds []*uint64 `json:"SkillGroupIds,omitempty" name:"SkillGroupIds"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 坐席邮箱
	StaffEmail *string `json:"StaffEmail,omitempty" name:"StaffEmail"`

	// 绑定技能组列表
	SkillGroupList []*int64 `json:"SkillGroupList,omitempty" name:"SkillGroupList"`
}

type BindStaffSkillGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 坐席邮箱
	StaffEmail *string `json:"StaffEmail,omitempty" name:"StaffEmail"`

	// 绑定技能组列表
	SkillGroupList []*int64 `json:"SkillGroupList,omitempty" name:"SkillGroupList"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindStaffSkillGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindStaffSkillGroupListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	IvrCount *int64 `json:"IvrCount,omitempty" name:"IvrCount"`

	// 排队中数量
	QueueCount *int64 `json:"QueueCount,omitempty" name:"QueueCount"`

	// 振铃中数量
	RingCount *int64 `json:"RingCount,omitempty" name:"RingCount"`

	// 接通中数量
	AcceptCount *int64 `json:"AcceptCount,omitempty" name:"AcceptCount"`

	// 客服转接外线中数量
	TransferOuterCount *int64 `json:"TransferOuterCount,omitempty" name:"TransferOuterCount"`

	// 最大排队时长
	MaxQueueDuration *int64 `json:"MaxQueueDuration,omitempty" name:"MaxQueueDuration"`

	// 平均排队时长
	AvgQueueDuration *int64 `json:"AvgQueueDuration,omitempty" name:"AvgQueueDuration"`

	// 最大振铃时长
	MaxRingDuration *int64 `json:"MaxRingDuration,omitempty" name:"MaxRingDuration"`

	// 平均振铃时长
	AvgRingDuration *int64 `json:"AvgRingDuration,omitempty" name:"AvgRingDuration"`

	// 最大接通时长
	MaxAcceptDuration *int64 `json:"MaxAcceptDuration,omitempty" name:"MaxAcceptDuration"`

	// 平均接通时长
	AvgAcceptDuration *int64 `json:"AvgAcceptDuration,omitempty" name:"AvgAcceptDuration"`
}

type CallInNumberMetrics struct {
	// 线路号码
	Number *string `json:"Number,omitempty" name:"Number"`

	// 线路相关指标
	Metrics *CallInMetrics `json:"Metrics,omitempty" name:"Metrics"`

	// 所属技能组相关指标
	SkillGroupMetrics []*CallInSkillGroupMetrics `json:"SkillGroupMetrics,omitempty" name:"SkillGroupMetrics"`
}

type CallInSkillGroupMetrics struct {
	// 技能组ID
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// 数据指标
	Metrics *CallInMetrics `json:"Metrics,omitempty" name:"Metrics"`

	// 技能组名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type CarrierPrivilegeNumberApplicant struct {
	// 实例Id
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 申请单Id
	ApplicantId *uint64 `json:"ApplicantId,omitempty" name:"ApplicantId"`

	// 主叫号码列表
	Callers []*string `json:"Callers,omitempty" name:"Callers"`

	// 被叫号码列表
	Callees []*string `json:"Callees,omitempty" name:"Callees"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 审批状态:1 待审核、2 通过、3 拒绝
	State *int64 `json:"State,omitempty" name:"State"`

	// 创建时间，unix时间戳(秒)
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间，unix时间戳(秒)
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type CreateAutoCalloutTaskRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 任务起始时间戳，Unix 秒级时间戳
	NotBefore *int64 `json:"NotBefore,omitempty" name:"NotBefore"`

	// 被叫号码列表
	Callees []*string `json:"Callees,omitempty" name:"Callees"`

	// 主叫号码列表
	Callers []*string `json:"Callers,omitempty" name:"Callers"`

	// 呼叫使用的Ivr
	IvrId *uint64 `json:"IvrId,omitempty" name:"IvrId"`

	// 任务名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 任务停止时间戳，Unix 秒级时间戳
	NotAfter *int64 `json:"NotAfter,omitempty" name:"NotAfter"`

	// 最大尝试次数
	Tries *uint64 `json:"Tries,omitempty" name:"Tries"`

	// 自定义变量（仅高级版支持）
	Variables []*Variable `json:"Variables,omitempty" name:"Variables"`
}

type CreateAutoCalloutTaskRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 任务起始时间戳，Unix 秒级时间戳
	NotBefore *int64 `json:"NotBefore,omitempty" name:"NotBefore"`

	// 被叫号码列表
	Callees []*string `json:"Callees,omitempty" name:"Callees"`

	// 主叫号码列表
	Callers []*string `json:"Callers,omitempty" name:"Callers"`

	// 呼叫使用的Ivr
	IvrId *uint64 `json:"IvrId,omitempty" name:"IvrId"`

	// 任务名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 任务停止时间戳，Unix 秒级时间戳
	NotAfter *int64 `json:"NotAfter,omitempty" name:"NotAfter"`

	// 最大尝试次数
	Tries *uint64 `json:"Tries,omitempty" name:"Tries"`

	// 自定义变量（仅高级版支持）
	Variables []*Variable `json:"Variables,omitempty" name:"Variables"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAutoCalloutTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoCalloutTaskResponseParams struct {
	// 任务Id
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 应用 ID（必填）
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 技能组名称
	SkillGroupName *string `json:"SkillGroupName,omitempty" name:"SkillGroupName"`

	// 技能组类型0-电话，1-在线，3-音频，4-视频
	SkillGroupType *int64 `json:"SkillGroupType,omitempty" name:"SkillGroupType"`

	// 技能组接待人数上限（该技能组中1个座席可接待的人数上限）默认为1。1、若技能组类型为在线，则接待上限可设置为1及以上
	// 2、若技能组类型为电话、音频、视频，则接待上线必须只能为1
	MaxConcurrency *uint64 `json:"MaxConcurrency,omitempty" name:"MaxConcurrency"`
}

type CreateCCCSkillGroupRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填）
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 技能组名称
	SkillGroupName *string `json:"SkillGroupName,omitempty" name:"SkillGroupName"`

	// 技能组类型0-电话，1-在线，3-音频，4-视频
	SkillGroupType *int64 `json:"SkillGroupType,omitempty" name:"SkillGroupType"`

	// 技能组接待人数上限（该技能组中1个座席可接待的人数上限）默认为1。1、若技能组类型为在线，则接待上限可设置为1及以上
	// 2、若技能组类型为电话、音频、视频，则接待上线必须只能为1
	MaxConcurrency *uint64 `json:"MaxConcurrency,omitempty" name:"MaxConcurrency"`
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
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 客服用户 ID，一般为客服邮箱
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 被叫号码，须带 0086 前缀
	Callee *string `json:"Callee,omitempty" name:"Callee"`

	// 主叫号码，须带 0086 前缀
	Caller *string `json:"Caller,omitempty" name:"Caller"`

	// 是否强制使用手机外呼，当前只支持 true，若为 true 请确保已配置白名单
	IsForceUseMobile *bool `json:"IsForceUseMobile,omitempty" name:"IsForceUseMobile"`

	// 自定义数据，长度限制 1024 字节
	Uui *string `json:"Uui,omitempty" name:"Uui"`
}

type CreateCallOutSessionRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 客服用户 ID，一般为客服邮箱
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 被叫号码，须带 0086 前缀
	Callee *string `json:"Callee,omitempty" name:"Callee"`

	// 主叫号码，须带 0086 前缀
	Caller *string `json:"Caller,omitempty" name:"Caller"`

	// 是否强制使用手机外呼，当前只支持 true，若为 true 请确保已配置白名单
	IsForceUseMobile *bool `json:"IsForceUseMobile,omitempty" name:"IsForceUseMobile"`

	// 自定义数据，长度限制 1024 字节
	Uui *string `json:"Uui,omitempty" name:"Uui"`
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
	delete(f, "IsForceUseMobile")
	delete(f, "Uui")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCallOutSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCallOutSessionResponseParams struct {
	// 新创建的会话 ID
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 主叫号码，必须为实例中存在的号码，格式为0086xxxx（暂时只支持国内号码）
	Callers []*string `json:"Callers,omitempty" name:"Callers"`

	// 被叫号码，必须为实例中坐席绑定的手机号码，格式为0086xxxx（暂时只支持国内号码）
	Callees []*string `json:"Callees,omitempty" name:"Callees"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateCarrierPrivilegeNumberApplicantRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 主叫号码，必须为实例中存在的号码，格式为0086xxxx（暂时只支持国内号码）
	Callers []*string `json:"Callers,omitempty" name:"Callers"`

	// 被叫号码，必须为实例中坐席绑定的手机号码，格式为0086xxxx（暂时只支持国内号码）
	Callees []*string `json:"Callees,omitempty" name:"Callees"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`
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
	ApplicantId *uint64 `json:"ApplicantId,omitempty" name:"ApplicantId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateExtensionRequestParams struct {
	// TCCC 实例应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitempty" name:"ExtensionId"`

	// 分机名称
	ExtensionName *string `json:"ExtensionName,omitempty" name:"ExtensionName"`
}

type CreateExtensionRequest struct {
	*tchttp.BaseRequest
	
	// TCCC 实例应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitempty" name:"ExtensionId"`

	// 分机名称
	ExtensionName *string `json:"ExtensionName,omitempty" name:"ExtensionName"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExtensionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExtensionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateSDKLoginTokenRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 坐席账号。
	SeatUserId *string `json:"SeatUserId,omitempty" name:"SeatUserId"`
}

type CreateSDKLoginTokenRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 坐席账号。
	SeatUserId *string `json:"SeatUserId,omitempty" name:"SeatUserId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSDKLoginTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSDKLoginTokenResponseParams struct {
	// SDK 登录 Token。
	Token *string `json:"Token,omitempty" name:"Token"`

	// 过期时间戳，Unix 时间戳。
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// SDK 加载路径会随着 SDK 的发布而变动。
	SdkURL *string `json:"SdkURL,omitempty" name:"SdkURL"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 客服信息，个数不超过 10
	Staffs []*SeatUserInfo `json:"Staffs,omitempty" name:"Staffs"`

	// 是否发送密码邮件，默认true
	SendPassword *bool `json:"SendPassword,omitempty" name:"SendPassword"`
}

type CreateStaffRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 客服信息，个数不超过 10
	Staffs []*SeatUserInfo `json:"Staffs,omitempty" name:"Staffs"`

	// 是否发送密码邮件，默认true
	SendPassword *bool `json:"SendPassword,omitempty" name:"SendPassword"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorStaffList []*ErrStaffItem `json:"ErrorStaffList,omitempty" name:"ErrorStaffList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户 ID
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 有效期，单位秒，不超过 1 小时
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 用户签名数据
	ClientData *string `json:"ClientData,omitempty" name:"ClientData"`
}

type CreateUserSigRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户 ID
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 有效期，单位秒，不超过 1 小时
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 用户签名数据
	ClientData *string `json:"ClientData,omitempty" name:"ClientData"`
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
	UserSig *string `json:"UserSig,omitempty" name:"UserSig"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// TCCC 实例应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitempty" name:"ExtensionId"`
}

type DeleteExtensionRequest struct {
	*tchttp.BaseRequest
	
	// TCCC 实例应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitempty" name:"ExtensionId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeleteStaffRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 待删除客服邮箱列表
	StaffList []*string `json:"StaffList,omitempty" name:"StaffList"`
}

type DeleteStaffRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 待删除客服邮箱列表
	StaffList []*string `json:"StaffList,omitempty" name:"StaffList"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnlineStaffList []*string `json:"OnlineStaffList,omitempty" name:"OnlineStaffList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeActiveCarrierPrivilegeNumberRequestParams struct {
	// 实例Id
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 默认0
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 默认10，最大100
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 筛选条件 Name支持PhoneNumber(按号码模糊查找)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeActiveCarrierPrivilegeNumberRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 默认0
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 默认10，最大100
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 筛选条件 Name支持PhoneNumber(按号码模糊查找)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 生效列表
	ActiveCarrierPrivilegeNumbers []*ActiveCarrierPrivilegeNumber `json:"ActiveCarrierPrivilegeNumbers,omitempty" name:"ActiveCarrierPrivilegeNumbers"`

	// 待审核单号
	PendingApplicantIds []*uint64 `json:"PendingApplicantIds,omitempty" name:"PendingApplicantIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 任务Id
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeAutoCalloutTaskRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 任务Id
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
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
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 任务起始时间戳
	NotBefore *int64 `json:"NotBefore,omitempty" name:"NotBefore"`

	// 任务结束时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotAfter *int64 `json:"NotAfter,omitempty" name:"NotAfter"`

	// 主叫列表
	Callers []*string `json:"Callers,omitempty" name:"Callers"`

	// 被叫信息列表
	Callees []*AutoCalloutTaskCalleeInfo `json:"Callees,omitempty" name:"Callees"`

	// 任务使用的IvrId
	IvrId *uint64 `json:"IvrId,omitempty" name:"IvrId"`

	// 任务状态 0初始 1运行中 2已完成 3结束中 4已终止
	State *uint64 `json:"State,omitempty" name:"State"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
}

type DescribeAutoCalloutTasksRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务列表
	Tasks []*AutoCalloutTaskInfo `json:"Tasks,omitempty" name:"Tasks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SdkAppIds []*int64 `json:"SdkAppIds,omitempty" name:"SdkAppIds"`
}

type DescribeCCCBuyInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID列表，不传时查询所有应用
	SdkAppIds []*int64 `json:"SdkAppIds,omitempty" name:"SdkAppIds"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 应用购买信息列表
	SdkAppIdBuyList []*SdkAppIdBuyInfo `json:"SdkAppIdBuyList,omitempty" name:"SdkAppIdBuyList"`

	// 套餐包购买信息列表
	PackageBuyList []*PackageBuyInfo `json:"PackageBuyList,omitempty" name:"PackageBuyList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 是否返回技能组维度信息，默认“是”
	EnabledSkillGroup *bool `json:"EnabledSkillGroup,omitempty" name:"EnabledSkillGroup"`

	// 是否返回线路维度信息，默认“否”
	EnabledNumber *bool `json:"EnabledNumber,omitempty" name:"EnabledNumber"`
}

type DescribeCallInMetricsRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 是否返回技能组维度信息，默认“是”
	EnabledSkillGroup *bool `json:"EnabledSkillGroup,omitempty" name:"EnabledSkillGroup"`

	// 是否返回线路维度信息，默认“否”
	EnabledNumber *bool `json:"EnabledNumber,omitempty" name:"EnabledNumber"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCallInMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallInMetricsResponseParams struct {
	// 时间戳
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 总体指标
	TotalMetrics *CallInMetrics `json:"TotalMetrics,omitempty" name:"TotalMetrics"`

	// 线路维度指标
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumberMetrics []*CallInNumberMetrics `json:"NumberMetrics,omitempty" name:"NumberMetrics"`

	// 技能组维度指标
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupMetrics []*CallInSkillGroupMetrics `json:"SkillGroupMetrics,omitempty" name:"SkillGroupMetrics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例Id
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 默认0，从0开始
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 默认10，最大100
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 筛选条件,Name支持ApplicantId,PhoneNumber(按号码模糊查找)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeCarrierPrivilegeNumberApplicantsRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 默认0，从0开始
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 默认10，最大100
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 筛选条件,Name支持ApplicantId,PhoneNumber(按号码模糊查找)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 申请单列表
	Applicants []*CarrierPrivilegeNumberApplicant `json:"Applicants,omitempty" name:"Applicants"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例 ID（废弃）
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 服务记录ID（废弃）
	CdrId *string `json:"CdrId,omitempty" name:"CdrId"`

	// 返回记录条数，最大为100 默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回记录偏移，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 1为从早到晚，2为从晚到早，默认为2
	Order *int64 `json:"Order,omitempty" name:"Order"`

	// 服务记录 SessionID（必填）
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type DescribeChatMessagesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID（废弃）
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 服务记录ID（废弃）
	CdrId *string `json:"CdrId,omitempty" name:"CdrId"`

	// 返回记录条数，最大为100 默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回记录偏移，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 1为从早到晚，2为从晚到早，默认为2
	Order *int64 `json:"Order,omitempty" name:"Order"`

	// 服务记录 SessionID（必填）
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
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
	delete(f, "InstanceId")
	delete(f, "SdkAppId")
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 消息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Messages []*MessageBody `json:"Messages,omitempty" name:"Messages"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeExtensionRequestParams struct {
	// TCCC 实例应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitempty" name:"ExtensionId"`
}

type DescribeExtensionRequest struct {
	*tchttp.BaseRequest
	
	// TCCC 实例应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitempty" name:"ExtensionId"`
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
	ExtensionId *string `json:"ExtensionId,omitempty" name:"ExtensionId"`

	// 域名
	ExtensionDomain *string `json:"ExtensionDomain,omitempty" name:"ExtensionDomain"`

	// 注册密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 代理服务器地址
	OutboundProxy *string `json:"OutboundProxy,omitempty" name:"OutboundProxy"`

	// 传输协议
	Transport *string `json:"Transport,omitempty" name:"Transport"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// TCCC 实例应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页页号（从0开始）
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 筛选分机号列表
	ExtensionIds []*string `json:"ExtensionIds,omitempty" name:"ExtensionIds"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 模糊查询字段（模糊查询分机号、分机名称、坐席邮箱、坐席名称）
	FuzzingKeyWord *string `json:"FuzzingKeyWord,omitempty" name:"FuzzingKeyWord"`

	// 是否需要返回话机当前状态
	IsNeedStatus *bool `json:"IsNeedStatus,omitempty" name:"IsNeedStatus"`
}

type DescribeExtensionsRequest struct {
	*tchttp.BaseRequest
	
	// TCCC 实例应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页页号（从0开始）
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 筛选分机号列表
	ExtensionIds []*string `json:"ExtensionIds,omitempty" name:"ExtensionIds"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 模糊查询字段（模糊查询分机号、分机名称、坐席邮箱、坐席名称）
	FuzzingKeyWord *string `json:"FuzzingKeyWord,omitempty" name:"FuzzingKeyWord"`

	// 是否需要返回话机当前状态
	IsNeedStatus *bool `json:"IsNeedStatus,omitempty" name:"IsNeedStatus"`
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
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 话机信息列表
	ExtensionList []*ExtensionInfo `json:"ExtensionList,omitempty" name:"ExtensionList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeIMCdrsRequestParams struct {
	// 起始时间（必填），Unix 秒级时间戳
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 结束时间（必填），Unix 秒级时间戳
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" name:"EndTimestamp"`

	// 实例 ID（废弃）
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 返回记录条数，最大为100默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回记录偏移，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 1为全媒体，2为文本客服，不填则查询全部
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type DescribeIMCdrsRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间（必填），Unix 秒级时间戳
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 结束时间（必填），Unix 秒级时间戳
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" name:"EndTimestamp"`

	// 实例 ID（废弃）
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 返回记录条数，最大为100默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回记录偏移，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 1为全媒体，2为文本客服，不填则查询全部
	Type *int64 `json:"Type,omitempty" name:"Type"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 服务记录列表
	IMCdrs []*IMCdrInfo `json:"IMCdrs,omitempty" name:"IMCdrs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeNumbersRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 页数，从0开始
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 分页大小，默认20
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeNumbersRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 页数，从0开始
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 分页大小，默认20
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 号码列表
	Numbers []*NumberInfo `json:"Numbers,omitempty" name:"Numbers"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 数据偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回的数据条数，最大 25
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribePSTNActiveSessionListRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 数据偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回的数据条数，最大 25
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 列表内容
	Sessions []*PSTNSessionInfo `json:"Sessions,omitempty" name:"Sessions"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeProtectedTelCdrRequestParams struct {
	// 起始时间戳，Unix 秒级时间戳
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 秒级时间戳
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`

	// 应用 ID，可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 100
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
}

type DescribeProtectedTelCdrRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间戳，Unix 秒级时间戳
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 秒级时间戳
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`

	// 应用 ID，可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 100
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 话单记录
	TelCdrs []*TelCdrInfo `json:"TelCdrs,omitempty" name:"TelCdrs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 100
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 技能组ID，查询单个技能组时使用
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// 查询修改时间大于等于ModifiedTime的技能组时使用
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 技能组名称
	SkillGroupName *string `json:"SkillGroupName,omitempty" name:"SkillGroupName"`
}

type DescribeSkillGroupInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 100
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 技能组ID，查询单个技能组时使用
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// 查询修改时间大于等于ModifiedTime的技能组时使用
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 技能组名称
	SkillGroupName *string `json:"SkillGroupName,omitempty" name:"SkillGroupName"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 技能组信息列表
	SkillGroupList []*SkillGroupInfoItem `json:"SkillGroupList,omitempty" name:"SkillGroupList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 9999
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 坐席账号，查询单个坐席时使用
	StaffMail *string `json:"StaffMail,omitempty" name:"StaffMail"`

	// 查询修改时间大于等于ModifiedTime的坐席时使用
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 技能组ID
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`
}

type DescribeStaffInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 9999
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 坐席账号，查询单个坐席时使用
	StaffMail *string `json:"StaffMail,omitempty" name:"StaffMail"`

	// 查询修改时间大于等于ModifiedTime的坐席时使用
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 技能组ID
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 坐席用户信息列表
	StaffList []*StaffInfo `json:"StaffList,omitempty" name:"StaffList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 筛选坐席列表，默认不传返回全部坐席信息
	StaffList []*string `json:"StaffList,omitempty" name:"StaffList"`
}

type DescribeStaffStatusMetricsRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 筛选坐席列表，默认不传返回全部坐席信息
	StaffList []*string `json:"StaffList,omitempty" name:"StaffList"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStaffStatusMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStaffStatusMetricsResponseParams struct {
	// 坐席状态实时信息
	Metrics []*StaffStatusMetrics `json:"Metrics,omitempty" name:"Metrics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 起始时间戳，Unix 时间戳
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 时间戳，查询时间范围最大为90天
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`

	// 应用ID列表，多个ID时，返回值为多个ID使用总和
	SdkAppIdList []*int64 `json:"SdkAppIdList,omitempty" name:"SdkAppIdList"`
}

type DescribeTelCallInfoRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间戳，Unix 时间戳
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 时间戳，查询时间范围最大为90天
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`

	// 应用ID列表，多个ID时，返回值为多个ID使用总和
	SdkAppIdList []*int64 `json:"SdkAppIdList,omitempty" name:"SdkAppIdList"`
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
	TelCallOutCount *int64 `json:"TelCallOutCount,omitempty" name:"TelCallOutCount"`

	// 呼入套餐包消耗分钟数
	TelCallInCount *int64 `json:"TelCallInCount,omitempty" name:"TelCallInCount"`

	// 坐席使用统计个数
	SeatUsedCount *int64 `json:"SeatUsedCount,omitempty" name:"SeatUsedCount"`

	// 音频套餐包消耗分钟数
	VoipCallInCount *int64 `json:"VoipCallInCount,omitempty" name:"VoipCallInCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 起始时间戳，Unix 秒级时间戳
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 秒级时间戳
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`

	// 实例 ID（废弃）
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 返回数据条数，上限（废弃）
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移（废弃）
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页尺寸（必填），上限 100
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页页码（必填），从 0 开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 按手机号筛选
	Phones []*string `json:"Phones,omitempty" name:"Phones"`

	// 按SessionId筛选
	SessionIds []*string `json:"SessionIds,omitempty" name:"SessionIds"`
}

type DescribeTelCdrRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间戳，Unix 秒级时间戳
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 秒级时间戳
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`

	// 实例 ID（废弃）
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 返回数据条数，上限（废弃）
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移（废弃）
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页尺寸（必填），上限 100
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页页码（必填），从 0 开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 按手机号筛选
	Phones []*string `json:"Phones,omitempty" name:"Phones"`

	// 按SessionId筛选
	SessionIds []*string `json:"SessionIds,omitempty" name:"SessionIds"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 话单记录
	TelCdrs []*TelCdrInfo `json:"TelCdrs,omitempty" name:"TelCdrs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeTelSessionRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 会话 ID
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type DescribeTelSessionRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 会话 ID
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
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
	Session *PSTNSession `json:"Session,omitempty" name:"Session"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PhoneNumbers []*string `json:"PhoneNumbers,omitempty" name:"PhoneNumbers"`

	// 停用开关，0启用 1停用
	Disabled *int64 `json:"Disabled,omitempty" name:"Disabled"`

	// TCCC 实例应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type DisableCCCPhoneNumberRequest struct {
	*tchttp.BaseRequest
	
	// 号码列表，0086开头
	PhoneNumbers []*string `json:"PhoneNumbers,omitempty" name:"PhoneNumbers"`

	// 停用开关，0启用 1停用
	Disabled *int64 `json:"Disabled,omitempty" name:"Disabled"`

	// TCCC 实例应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 坐席邮箱地址
	StaffEmail *string `json:"StaffEmail,omitempty" name:"StaffEmail"`

	// 错误码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 错误描述
	Message *string `json:"Message,omitempty" name:"Message"`
}

type ExtensionInfo struct {
	// 实例ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分机全名
	FullExtensionId *string `json:"FullExtensionId,omitempty" name:"FullExtensionId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitempty" name:"ExtensionId"`

	// 所属技能组列表
	SkillGroupId *string `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// 分机名称
	ExtensionName *string `json:"ExtensionName,omitempty" name:"ExtensionName"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最后修改时间
	ModifyTime *int64 `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 话机状态(0 离线、100 空闲、200忙碌）
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 是否注册
	Register *bool `json:"Register,omitempty" name:"Register"`

	// 绑定坐席邮箱
	Relation *string `json:"Relation,omitempty" name:"Relation"`

	// 绑定坐席名称
	RelationName *string `json:"RelationName,omitempty" name:"RelationName"`
}

type Filter struct {
	// 筛选字段名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 筛选条件值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

// Predefined struct for user
type HangUpCallRequestParams struct {
	// TCCC 实例应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 会话ID
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type HangUpCallRequest struct {
	*tchttp.BaseRequest
	
	// TCCC 实例应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 会话ID
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Id *string `json:"Id,omitempty" name:"Id"`

	// 服务时长秒数
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 结束状态
	// 0 异常结束
	// 1 正常结束
	// 3 无坐席在线
	// 17 坐席放弃接听
	// 100 黑名单
	// 101 坐席手动转接
	// 102 IVR阶段放弃
	// 108 用户超时自动结束
	EndStatus *int64 `json:"EndStatus,omitempty" name:"EndStatus"`

	// 用户昵称
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// 服务类型 1为全媒体，2为文本客服
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 客服ID
	StaffId *string `json:"StaffId,omitempty" name:"StaffId"`

	// 服务时间戳
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 会话ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 技能组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupId *string `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// 技能组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupName *string `json:"SkillGroupName,omitempty" name:"SkillGroupName"`

	// 满意度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Satisfaction *IMSatisfaction `json:"Satisfaction,omitempty" name:"Satisfaction"`
}

type IMSatisfaction struct {
	// 满意度值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 满意度标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`
}

type IVRKeyPressedElement struct {
	// 按键
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 按键关联的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`
}

type Message struct {
	// 消息类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 消息内容
	Content *string `json:"Content,omitempty" name:"Content"`
}

type MessageBody struct {
	// 消息时间戳
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 发消息的用户ID
	From *string `json:"From,omitempty" name:"From"`

	// 消息列表
	Messages []*Message `json:"Messages,omitempty" name:"Messages"`
}

// Predefined struct for user
type ModifyExtensionRequestParams struct {
	// TCCC 实例应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitempty" name:"ExtensionId"`

	// 分机名称
	ExtensionName *string `json:"ExtensionName,omitempty" name:"ExtensionName"`

	// 所属技能组列表
	SkillGroupIds []*int64 `json:"SkillGroupIds,omitempty" name:"SkillGroupIds"`

	// 绑定坐席邮箱账号
	Relation *string `json:"Relation,omitempty" name:"Relation"`
}

type ModifyExtensionRequest struct {
	*tchttp.BaseRequest
	
	// TCCC 实例应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitempty" name:"ExtensionId"`

	// 分机名称
	ExtensionName *string `json:"ExtensionName,omitempty" name:"ExtensionName"`

	// 所属技能组列表
	SkillGroupIds []*int64 `json:"SkillGroupIds,omitempty" name:"SkillGroupIds"`

	// 绑定坐席邮箱账号
	Relation *string `json:"Relation,omitempty" name:"Relation"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyStaffRequestParams struct {
	// 应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 坐席账户
	Email *string `json:"Email,omitempty" name:"Email"`

	// 坐席名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 坐席手机号（带0086前缀,示例：008618011111111）
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 坐席昵称
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// 绑定技能组ID列表
	SkillGroupIds []*int64 `json:"SkillGroupIds,omitempty" name:"SkillGroupIds"`

	// 是否开启手机外呼开关
	UseMobileCallOut *bool `json:"UseMobileCallOut,omitempty" name:"UseMobileCallOut"`

	// 手机接听模式 0 - 关闭 | 1 - 仅离线 | 2 - 始终
	UseMobileAccept *int64 `json:"UseMobileAccept,omitempty" name:"UseMobileAccept"`
}

type ModifyStaffRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 坐席账户
	Email *string `json:"Email,omitempty" name:"Email"`

	// 坐席名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 坐席手机号（带0086前缀,示例：008618011111111）
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 坐席昵称
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// 绑定技能组ID列表
	SkillGroupIds []*int64 `json:"SkillGroupIds,omitempty" name:"SkillGroupIds"`

	// 是否开启手机外呼开关
	UseMobileCallOut *bool `json:"UseMobileCallOut,omitempty" name:"UseMobileCallOut"`

	// 手机接听模式 0 - 关闭 | 1 - 仅离线 | 2 - 始终
	UseMobileAccept *int64 `json:"UseMobileAccept,omitempty" name:"UseMobileAccept"`
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
	delete(f, "SkillGroupIds")
	delete(f, "UseMobileCallOut")
	delete(f, "UseMobileAccept")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStaffRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStaffResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Number *string `json:"Number,omitempty" name:"Number"`

	// 绑定的外呼技能组
	CallOutSkillGroupIds []*uint64 `json:"CallOutSkillGroupIds,omitempty" name:"CallOutSkillGroupIds"`
}

type PSTNSession struct {
	// 会话 ID
	SessionID *string `json:"SessionID,omitempty" name:"SessionID"`

	// 会话临时房间 ID
	RoomID *string `json:"RoomID,omitempty" name:"RoomID"`

	// 主叫
	Caller *string `json:"Caller,omitempty" name:"Caller"`

	// 被叫
	Callee *string `json:"Callee,omitempty" name:"Callee"`

	// 开始时间，Unix 时间戳
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 振铃时间，Unix 时间戳
	RingTimestamp *int64 `json:"RingTimestamp,omitempty" name:"RingTimestamp"`

	// 接听时间，Unix 时间戳
	AcceptTimestamp *int64 `json:"AcceptTimestamp,omitempty" name:"AcceptTimestamp"`

	// 坐席邮箱
	StaffEmail *string `json:"StaffEmail,omitempty" name:"StaffEmail"`

	// 坐席工号
	StaffNumber *string `json:"StaffNumber,omitempty" name:"StaffNumber"`

	// 会话状态
	// ringing 振铃中
	// seatJoining  等待坐席接听
	// inProgress 进行中
	// finished 已完成
	SessionStatus *string `json:"SessionStatus,omitempty" name:"SessionStatus"`

	// 会话呼叫方向， 0 呼入 | 1 - 呼出
	Direction *int64 `json:"Direction,omitempty" name:"Direction"`

	// 转外线使用的号码（转外线主叫）
	OutBoundCaller *string `json:"OutBoundCaller,omitempty" name:"OutBoundCaller"`

	// 转外线被叫
	OutBoundCallee *string `json:"OutBoundCallee,omitempty" name:"OutBoundCallee"`

	// 主叫号码保护ID，开启号码保护映射功能时有效，且Caller字段置空
	ProtectedCaller *string `json:"ProtectedCaller,omitempty" name:"ProtectedCaller"`

	// 被叫号码保护ID，开启号码保护映射功能时有效，且Callee字段置空
	ProtectedCallee *string `json:"ProtectedCallee,omitempty" name:"ProtectedCallee"`
}

type PSTNSessionInfo struct {
	// 会话 ID
	SessionID *string `json:"SessionID,omitempty" name:"SessionID"`

	// 会话临时房间 ID
	RoomID *string `json:"RoomID,omitempty" name:"RoomID"`

	// 主叫
	Caller *string `json:"Caller,omitempty" name:"Caller"`

	// 被叫
	Callee *string `json:"Callee,omitempty" name:"Callee"`

	// 开始时间，Unix 时间戳
	StartTimestamp *string `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 接听时间，Unix 时间戳
	AcceptTimestamp *string `json:"AcceptTimestamp,omitempty" name:"AcceptTimestamp"`

	// 坐席邮箱
	StaffEmail *string `json:"StaffEmail,omitempty" name:"StaffEmail"`

	// 坐席工号
	StaffNumber *string `json:"StaffNumber,omitempty" name:"StaffNumber"`

	// 坐席状态 inProgress 进行中
	SessionStatus *string `json:"SessionStatus,omitempty" name:"SessionStatus"`

	// 会话呼叫方向， 0 呼入 | 1 - 呼出
	Direction *int64 `json:"Direction,omitempty" name:"Direction"`

	// 振铃时间，Unix 时间戳
	RingTimestamp *int64 `json:"RingTimestamp,omitempty" name:"RingTimestamp"`

	// 主叫号码保护ID，开启号码保护映射功能时有效，且Caller字段置空
	ProtectedCaller *string `json:"ProtectedCaller,omitempty" name:"ProtectedCaller"`

	// 被叫号码保护ID，开启号码保护映射功能时有效，且Callee字段置空
	ProtectedCallee *string `json:"ProtectedCallee,omitempty" name:"ProtectedCallee"`
}

type PackageBuyInfo struct {
	// 套餐包Id
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 套餐包类型，0-外呼套餐包|1-400呼入套餐包
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 套餐包总量
	CapacitySize *int64 `json:"CapacitySize,omitempty" name:"CapacitySize"`

	// 套餐包剩余量
	CapacityRemain *int64 `json:"CapacityRemain,omitempty" name:"CapacityRemain"`

	// 购买时间戳
	BuyTime *int64 `json:"BuyTime,omitempty" name:"BuyTime"`

	// 截至时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type PhoneNumBuyInfo struct {
	// 电话号码
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// 号码类型，0-固话|1-虚商号码|2-运营商号码|3-400号码
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 号码呼叫类型，1-呼入|2-呼出|3-呼入呼出
	CallType *int64 `json:"CallType,omitempty" name:"CallType"`

	// 购买时间戳
	BuyTime *int64 `json:"BuyTime,omitempty" name:"BuyTime"`

	// 截至时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 号码状态，1正常|2欠费停用|4管理员停用|5违规停用
	State *int64 `json:"State,omitempty" name:"State"`
}

// Predefined struct for user
type ResetExtensionPasswordRequestParams struct {
	// TCCC 实例应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitempty" name:"ExtensionId"`
}

type ResetExtensionPasswordRequest struct {
	*tchttp.BaseRequest
	
	// TCCC 实例应用 ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分机号
	ExtensionId *string `json:"ExtensionId,omitempty" name:"ExtensionId"`
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
	Password *string `json:"Password,omitempty" name:"Password"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type SdkAppIdBuyInfo struct {
	// 应用ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 应用名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 坐席购买数（还在有效期内）
	StaffBuyNum *int64 `json:"StaffBuyNum,omitempty" name:"StaffBuyNum"`

	// 坐席购买列表 （还在有效期内）
	StaffBuyList []*StaffBuyInfo `json:"StaffBuyList,omitempty" name:"StaffBuyList"`

	// 号码购买列表
	PhoneNumBuyList []*PhoneNumBuyInfo `json:"PhoneNumBuyList,omitempty" name:"PhoneNumBuyList"`
}

type SeatUserInfo struct {
	// 坐席名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 坐席邮箱
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// 坐席电话号码（带0086前缀）
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 坐席昵称
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 坐席关联的技能组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupNameList []*string `json:"SkillGroupNameList,omitempty" name:"SkillGroupNameList"`

	// 工号
	// 注意：此字段可能返回 null，表示取不到有效值。
	StaffNumber *string `json:"StaffNumber,omitempty" name:"StaffNumber"`
}

type ServeParticipant struct {
	// 坐席邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// 坐席电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 振铃时间戳，Unix 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	RingTimestamp *int64 `json:"RingTimestamp,omitempty" name:"RingTimestamp"`

	// 接听时间戳，Unix 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	AcceptTimestamp *int64 `json:"AcceptTimestamp,omitempty" name:"AcceptTimestamp"`

	// 结束时间戳，Unix 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndedTimestamp *int64 `json:"EndedTimestamp,omitempty" name:"EndedTimestamp"`

	// 录音 ID，能够索引到坐席侧的录音
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// 参与者类型，"staffSeat", "outboundSeat", "staffPhoneSeat"
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 转接来源坐席信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferFrom *string `json:"TransferFrom,omitempty" name:"TransferFrom"`

	// 转接去向坐席信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferTo *string `json:"TransferTo,omitempty" name:"TransferTo"`

	// 转接去向参与者类型，取值与 Type 一致
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferToType *string `json:"TransferToType,omitempty" name:"TransferToType"`

	// 技能组 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// 结束状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndStatusString *string `json:"EndStatusString,omitempty" name:"EndStatusString"`

	// 录音 URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordURL *string `json:"RecordURL,omitempty" name:"RecordURL"`

	// 参与者序号，从 0 开始
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sequence *int64 `json:"Sequence,omitempty" name:"Sequence"`

	// 开始时间戳，Unix 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 技能组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupName *string `json:"SkillGroupName,omitempty" name:"SkillGroupName"`

	// 录音转存第三方COS地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomRecordURL *string `json:"CustomRecordURL,omitempty" name:"CustomRecordURL"`
}

type SkillGroupInfoItem struct {
	// 技能组ID
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// 技能组名称
	SkillGroupName *string `json:"SkillGroupName,omitempty" name:"SkillGroupName"`

	// （废弃）类型：IM、TEL、ALL（全媒体）
	Type *string `json:"Type,omitempty" name:"Type"`

	// 会话分配策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoutePolicy *string `json:"RoutePolicy,omitempty" name:"RoutePolicy"`

	// 会话分配是否优先上次服务坐席
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsingLastSeat *int64 `json:"UsingLastSeat,omitempty" name:"UsingLastSeat"`

	// 单客服最大并发数（电话类型默认1）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" name:"MaxConcurrency"`

	// 最后修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyTimestamp *int64 `json:"LastModifyTimestamp,omitempty" name:"LastModifyTimestamp"`

	// 技能组类型0-电话，1-在线，3-音频，4-视频	
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupType *int64 `json:"SkillGroupType,omitempty" name:"SkillGroupType"`
}

type SkillGroupItem struct {
	// 技能组ID
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// 技能组名称
	SkillGroupName *string `json:"SkillGroupName,omitempty" name:"SkillGroupName"`

	// 优先级
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 类型：IM、TEL、ALL（全媒体）
	Type *string `json:"Type,omitempty" name:"Type"`
}

type StaffBuyInfo struct {
	// 购买坐席数量
	Num *int64 `json:"Num,omitempty" name:"Num"`

	// 购买时间戳
	BuyTime *int64 `json:"BuyTime,omitempty" name:"BuyTime"`

	// 截至时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type StaffInfo struct {
	// 坐席名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 坐席邮箱
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// 坐席电话号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 坐席昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// 坐席工号
	// 注意：此字段可能返回 null，表示取不到有效值。
	StaffNumber *string `json:"StaffNumber,omitempty" name:"StaffNumber"`

	// 所属技能组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupList []*SkillGroupItem `json:"SkillGroupList,omitempty" name:"SkillGroupList"`

	// 最后修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyTimestamp *int64 `json:"LastModifyTimestamp,omitempty" name:"LastModifyTimestamp"`
}

type StaffStatusExtra struct {
	// im - 文本 | tel - 电话 | all - 全媒体
	Type *string `json:"Type,omitempty" name:"Type"`

	// in - 呼入 | out - 呼出
	Direct *string `json:"Direct,omitempty" name:"Direct"`
}

type StaffStatusMetrics struct {
	// 坐席邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 坐席状态 free 示闲 | busy 忙碌 | rest 小休 | notReady 示忙 | afterCallWork 话后调整 | offline 离线
	Status *string `json:"Status,omitempty" name:"Status"`

	// 坐席状态补充信息
	StatusExtra *StaffStatusExtra `json:"StatusExtra,omitempty" name:"StatusExtra"`

	// 当天在线总时长
	OnlineDuration *int64 `json:"OnlineDuration,omitempty" name:"OnlineDuration"`

	// 当天示闲总时长
	FreeDuration *int64 `json:"FreeDuration,omitempty" name:"FreeDuration"`

	// 当天忙碌总时长
	BusyDuration *int64 `json:"BusyDuration,omitempty" name:"BusyDuration"`

	// 当天示忙总时长
	NotReadyDuration *int64 `json:"NotReadyDuration,omitempty" name:"NotReadyDuration"`

	// 当天小休总时长
	RestDuration *int64 `json:"RestDuration,omitempty" name:"RestDuration"`

	// 当天话后调整总时长
	AfterCallWorkDuration *int64 `json:"AfterCallWorkDuration,omitempty" name:"AfterCallWorkDuration"`

	// 小休原因
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 是否预约小休
	ReserveRest *bool `json:"ReserveRest,omitempty" name:"ReserveRest"`

	// 是否预约示忙
	ReserveNotReady *bool `json:"ReserveNotReady,omitempty" name:"ReserveNotReady"`

	// 手机接听模式： 0 - 关闭 | 1 - 仅离线 | 2- 始终
	UseMobileAccept *int64 `json:"UseMobileAccept,omitempty" name:"UseMobileAccept"`

	// 手机外呼开关
	UseMobileCallOut *bool `json:"UseMobileCallOut,omitempty" name:"UseMobileCallOut"`
}

// Predefined struct for user
type StopAutoCalloutTaskRequestParams struct {
	// 呼叫中心实例Id
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 任务Id
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type StopAutoCalloutTaskRequest struct {
	*tchttp.BaseRequest
	
	// 呼叫中心实例Id
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 任务Id
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Caller *string `json:"Caller,omitempty" name:"Caller"`

	// 被叫号码
	Callee *string `json:"Callee,omitempty" name:"Callee"`

	// 呼叫发起时间戳，Unix 时间戳
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// 呼入呼出方向 0 呼入 1 呼出
	Direction *int64 `json:"Direction,omitempty" name:"Direction"`

	// 通话时长
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 录音信息
	RecordURL *string `json:"RecordURL,omitempty" name:"RecordURL"`

	// 坐席信息
	SeatUser *SeatUserInfo `json:"SeatUser,omitempty" name:"SeatUser"`

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
	// 电话呼入	             105	        noSeatOnline	       无坐席在线
	// 
	// 电话呼入              106	       notWorkTime	       非工作时间   
	// 
	// 电话呼入	            107	       ivrEnd	               IVR 后直接结束
	// 
	// 电话呼入	            100	      CallinBlockedContact  呼入黑名单 
	// 
	// 电话呼出               2	              unconnected	未接通
	//                          
	// 电话呼出             201            unknown	未知状态
	// 
	// 电话呼出            203	    userReject	拒接挂断
	// 
	// 电话呼出	          204	    powerOff	关机
	// 
	// 电话呼出           205            numberNotExist	空号
	// 
	// 电话呼出	         206	           busy	通话中
	// 
	// 电话呼出   	 207	           outOfCredit	欠费
	// 
	// 电话呼出	         208	           operatorError	运营商线路异常
	// 
	// 电话呼出         	209	           callerCancel	主叫取消
	// 
	// 电话呼出	        210	           notInService	不在服务区
	EndStatus *int64 `json:"EndStatus,omitempty" name:"EndStatus"`

	// 技能组名称
	SkillGroup *string `json:"SkillGroup,omitempty" name:"SkillGroup"`

	// 主叫归属地
	CallerLocation *string `json:"CallerLocation,omitempty" name:"CallerLocation"`

	// IVR 阶段耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	IVRDuration *int64 `json:"IVRDuration,omitempty" name:"IVRDuration"`

	// 振铃时间戳，UNIX 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	RingTimestamp *int64 `json:"RingTimestamp,omitempty" name:"RingTimestamp"`

	// 接听时间戳，UNIX 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	AcceptTimestamp *int64 `json:"AcceptTimestamp,omitempty" name:"AcceptTimestamp"`

	// 结束时间戳，UNIX 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndedTimestamp *int64 `json:"EndedTimestamp,omitempty" name:"EndedTimestamp"`

	// IVR 按键信息 ，e.g. ["1","2","3"]
	// 注意：此字段可能返回 null，表示取不到有效值。
	IVRKeyPressed []*string `json:"IVRKeyPressed,omitempty" name:"IVRKeyPressed"`

	// 挂机方 seat 坐席 user 用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	HungUpSide *string `json:"HungUpSide,omitempty" name:"HungUpSide"`

	// 服务参与者列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServeParticipants []*ServeParticipant `json:"ServeParticipants,omitempty" name:"ServeParticipants"`

	// 技能组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

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
	// 电话呼入	             105	        noSeatOnline	       无坐席在线
	// 
	// 电话呼入              106	       notWorkTime	       非工作时间   
	// 
	// 电话呼入	            107	       ivrEnd	               IVR 后直接结束
	// 
	// 电话呼入	            100	      CallinBlockedContact  呼入黑名单 
	// 
	// 电话呼出               2	              unconnected	未接通
	//                          
	// 电话呼出             201            unknown	未知状态
	// 听
	// 电话呼出            203	    userReject	拒接挂断
	// 
	// 电话呼出	          204	    powerOff	关机
	// 
	// 电话呼出           205            numberNotExist	空号
	// 
	// 电话呼出	         206	           busy	通话中
	// 
	// 电话呼出   	 207	           outOfCredit	欠费
	// 
	// 电话呼出	         208	           operatorError	运营商线路异常
	// 
	// 电话呼出         	209	           callerCancel	主叫取消
	// 
	// 电话呼出	        210	           notInService	不在服务区
	// 
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndStatusString *string `json:"EndStatusString,omitempty" name:"EndStatusString"`

	// 会话开始时间戳，UNIX 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 进入排队时间，Unix 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueuedTimestamp *int64 `json:"QueuedTimestamp,omitempty" name:"QueuedTimestamp"`

	// 后置IVR按键信息（e.g. [{"Key":"1","Label":"非常满意"}]）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostIVRKeyPressed []*IVRKeyPressedElement `json:"PostIVRKeyPressed,omitempty" name:"PostIVRKeyPressed"`

	// 排队技能组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueuedSkillGroupId *int64 `json:"QueuedSkillGroupId,omitempty" name:"QueuedSkillGroupId"`

	// 会话 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 主叫号码保护ID，开启号码保护映射功能时有效，且Caller字段置空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectedCaller *string `json:"ProtectedCaller,omitempty" name:"ProtectedCaller"`

	// 被叫号码保护ID，开启号码保护映射功能时有效，且Callee字段置空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectedCallee *string `json:"ProtectedCallee,omitempty" name:"ProtectedCallee"`

	// 客户自定义数据（User-to-User Interface）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uui *string `json:"Uui,omitempty" name:"Uui"`

	// IVR按键信息（e.g. [{"Key":"1","Label":"非常满意"}]）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IVRKeyPressedEx []*IVRKeyPressedElement `json:"IVRKeyPressedEx,omitempty" name:"IVRKeyPressedEx"`

	// 获取录音ASR文本信息地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrUrl *string `json:"AsrUrl,omitempty" name:"AsrUrl"`

	// 录音转存第三方COS地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomRecordURL *string `json:"CustomRecordURL,omitempty" name:"CustomRecordURL"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 排队技能组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueuedSkillGroupName *string `json:"QueuedSkillGroupName,omitempty" name:"QueuedSkillGroupName"`

	// 通话中语音留言录音URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoicemailRecordURL []*string `json:"VoicemailRecordURL,omitempty" name:"VoicemailRecordURL"`
}

// Predefined struct for user
type UnbindNumberCallOutSkillGroupRequestParams struct {
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 待解绑的号码
	Number *string `json:"Number,omitempty" name:"Number"`

	// 待解绑的技能组Id列表
	SkillGroupIds []*uint64 `json:"SkillGroupIds,omitempty" name:"SkillGroupIds"`
}

type UnbindNumberCallOutSkillGroupRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 待解绑的号码
	Number *string `json:"Number,omitempty" name:"Number"`

	// 待解绑的技能组Id列表
	SkillGroupIds []*uint64 `json:"SkillGroupIds,omitempty" name:"SkillGroupIds"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 客服邮箱
	StaffEmail *string `json:"StaffEmail,omitempty" name:"StaffEmail"`

	// 解绑技能组列表
	SkillGroupList []*int64 `json:"SkillGroupList,omitempty" name:"SkillGroupList"`
}

type UnbindStaffSkillGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID（必填），可以查看 https://console.cloud.tencent.com/ccc
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 客服邮箱
	StaffEmail *string `json:"StaffEmail,omitempty" name:"StaffEmail"`

	// 解绑技能组列表
	SkillGroupList []*int64 `json:"SkillGroupList,omitempty" name:"SkillGroupList"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type Variable struct {
	// 变量名
	Key *string `json:"Key,omitempty" name:"Key"`

	// 变量值
	Value *string `json:"Value,omitempty" name:"Value"`
}