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

package v20190103

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddMetricScaleStrategyRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 1表示按负载规则扩缩容，2表示按时间规则扩缩容。必须填写，并且和下面的规则策略匹配
	StrategyType *int64 `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 按负载扩容的规则。
	LoadAutoScaleStrategy *LoadAutoScaleStrategy `json:"LoadAutoScaleStrategy,omitnil,omitempty" name:"LoadAutoScaleStrategy"`

	// 按时间扩缩容的规则。
	TimeAutoScaleStrategy *TimeAutoScaleStrategy `json:"TimeAutoScaleStrategy,omitnil,omitempty" name:"TimeAutoScaleStrategy"`
}

type AddMetricScaleStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 1表示按负载规则扩缩容，2表示按时间规则扩缩容。必须填写，并且和下面的规则策略匹配
	StrategyType *int64 `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 按负载扩容的规则。
	LoadAutoScaleStrategy *LoadAutoScaleStrategy `json:"LoadAutoScaleStrategy,omitnil,omitempty" name:"LoadAutoScaleStrategy"`

	// 按时间扩缩容的规则。
	TimeAutoScaleStrategy *TimeAutoScaleStrategy `json:"TimeAutoScaleStrategy,omitnil,omitempty" name:"TimeAutoScaleStrategy"`
}

func (r *AddMetricScaleStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddMetricScaleStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StrategyType")
	delete(f, "LoadAutoScaleStrategy")
	delete(f, "TimeAutoScaleStrategy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddMetricScaleStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddMetricScaleStrategyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddMetricScaleStrategyResponse struct {
	*tchttp.BaseResponse
	Response *AddMetricScaleStrategyResponseParams `json:"Response"`
}

func (r *AddMetricScaleStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddMetricScaleStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNodeResourceConfigRequestParams struct {
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点类型 CORE TASK ROUTER
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源详情
	ResourceConfig *Resource `json:"ResourceConfig,omitnil,omitempty" name:"ResourceConfig"`

	// 付费模式
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 是否默认配置,DEFAULT,BACKUP,不填默认不是默认配置
	IsDefault *string `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 地域ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 添加多个规格时，第1个规格详情在ResourceConfig参数，第2-n个在MultipleResourceConfig参数
	MultipleResourceConfig []*Resource `json:"MultipleResourceConfig,omitnil,omitempty" name:"MultipleResourceConfig"`

	// 类型为ComputeResource和EMR以及默认，默认为EMR
	ResourceBaseType *string `json:"ResourceBaseType,omitnil,omitempty" name:"ResourceBaseType"`

	// 计算资源id
	ComputeResourceId *string `json:"ComputeResourceId,omitnil,omitempty" name:"ComputeResourceId"`

	// 硬件类型
	HardwareResourceType *string `json:"HardwareResourceType,omitnil,omitempty" name:"HardwareResourceType"`
}

type AddNodeResourceConfigRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点类型 CORE TASK ROUTER
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源详情
	ResourceConfig *Resource `json:"ResourceConfig,omitnil,omitempty" name:"ResourceConfig"`

	// 付费模式
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 是否默认配置,DEFAULT,BACKUP,不填默认不是默认配置
	IsDefault *string `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 地域ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 添加多个规格时，第1个规格详情在ResourceConfig参数，第2-n个在MultipleResourceConfig参数
	MultipleResourceConfig []*Resource `json:"MultipleResourceConfig,omitnil,omitempty" name:"MultipleResourceConfig"`

	// 类型为ComputeResource和EMR以及默认，默认为EMR
	ResourceBaseType *string `json:"ResourceBaseType,omitnil,omitempty" name:"ResourceBaseType"`

	// 计算资源id
	ComputeResourceId *string `json:"ComputeResourceId,omitnil,omitempty" name:"ComputeResourceId"`

	// 硬件类型
	HardwareResourceType *string `json:"HardwareResourceType,omitnil,omitempty" name:"HardwareResourceType"`
}

func (r *AddNodeResourceConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNodeResourceConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceType")
	delete(f, "ResourceConfig")
	delete(f, "PayMode")
	delete(f, "IsDefault")
	delete(f, "ZoneId")
	delete(f, "MultipleResourceConfig")
	delete(f, "ResourceBaseType")
	delete(f, "ComputeResourceId")
	delete(f, "HardwareResourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddNodeResourceConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNodeResourceConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddNodeResourceConfigResponse struct {
	*tchttp.BaseResponse
	Response *AddNodeResourceConfigResponseParams `json:"Response"`
}

func (r *AddNodeResourceConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNodeResourceConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUsersForUserManagerRequestParams struct {
	// 集群字符串ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户信息列表
	UserManagerUserList []*UserInfoForUserManager `json:"UserManagerUserList,omitnil,omitempty" name:"UserManagerUserList"`
}

type AddUsersForUserManagerRequest struct {
	*tchttp.BaseRequest
	
	// 集群字符串ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户信息列表
	UserManagerUserList []*UserInfoForUserManager `json:"UserManagerUserList,omitnil,omitempty" name:"UserManagerUserList"`
}

func (r *AddUsersForUserManagerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUsersForUserManagerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserManagerUserList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUsersForUserManagerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUsersForUserManagerResponseParams struct {
	// 添加成功的用户列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessUserList []*string `json:"SuccessUserList,omitnil,omitempty" name:"SuccessUserList"`

	// 添加失败的用户列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedUserList []*string `json:"FailedUserList,omitnil,omitempty" name:"FailedUserList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddUsersForUserManagerResponse struct {
	*tchttp.BaseResponse
	Response *AddUsersForUserManagerResponseParams `json:"Response"`
}

func (r *AddUsersForUserManagerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUsersForUserManagerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllNodeResourceSpec struct {
	// 描述Master节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterResourceSpec *NodeResourceSpec `json:"MasterResourceSpec,omitnil,omitempty" name:"MasterResourceSpec"`

	// 描述Core节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoreResourceSpec *NodeResourceSpec `json:"CoreResourceSpec,omitnil,omitempty" name:"CoreResourceSpec"`

	// 描述Taskr节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskResourceSpec *NodeResourceSpec `json:"TaskResourceSpec,omitnil,omitempty" name:"TaskResourceSpec"`

	// 描述Common节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommonResourceSpec *NodeResourceSpec `json:"CommonResourceSpec,omitnil,omitempty" name:"CommonResourceSpec"`

	// Master节点数量
	MasterCount *int64 `json:"MasterCount,omitnil,omitempty" name:"MasterCount"`

	// Corer节点数量
	CoreCount *int64 `json:"CoreCount,omitnil,omitempty" name:"CoreCount"`

	// Task节点数量
	TaskCount *int64 `json:"TaskCount,omitnil,omitempty" name:"TaskCount"`

	// Common节点数量
	CommonCount *int64 `json:"CommonCount,omitnil,omitempty" name:"CommonCount"`
}

type ApplicationStatics struct {
	// 队列名
	Queue *string `json:"Queue,omitnil,omitempty" name:"Queue"`

	// 用户名
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 作业类型
	ApplicationType *string `json:"ApplicationType,omitnil,omitempty" name:"ApplicationType"`

	// SumMemorySeconds含义
	SumMemorySeconds *int64 `json:"SumMemorySeconds,omitnil,omitempty" name:"SumMemorySeconds"`

	// SumVCoreSeconds含义
	SumVCoreSeconds *int64 `json:"SumVCoreSeconds,omitnil,omitempty" name:"SumVCoreSeconds"`

	// SumHDFSBytesWritten（带单位）
	SumHDFSBytesWritten *string `json:"SumHDFSBytesWritten,omitnil,omitempty" name:"SumHDFSBytesWritten"`

	// SumHDFSBytesRead（待单位）
	SumHDFSBytesRead *string `json:"SumHDFSBytesRead,omitnil,omitempty" name:"SumHDFSBytesRead"`

	// 作业数
	CountApps *int64 `json:"CountApps,omitnil,omitempty" name:"CountApps"`
}

type Arg struct {
	// key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type AttachDisksRequestParams struct {
	// EMR集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要挂载的云盘ID
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 挂载模式，取值范围：
	// AUTO_RENEW：自动续费
	// ALIGN_DEADLINE：自动对其到期时间
	AlignType *string `json:"AlignType,omitnil,omitempty" name:"AlignType"`

	// 需要挂载的cvm节点id列表
	CvmInstanceIds []*string `json:"CvmInstanceIds,omitnil,omitempty" name:"CvmInstanceIds"`

	// 是否是新购云盘进行挂载
	CreateDisk *bool `json:"CreateDisk,omitnil,omitempty" name:"CreateDisk"`

	// 新购云盘规格
	DiskSpec *NodeSpecDiskV2 `json:"DiskSpec,omitnil,omitempty" name:"DiskSpec"`

	// 可选参数，不传该参数则仅执行挂载操作。传入True时，会在挂载成功后将云硬盘设置为随云主机销毁模式，仅对按量计费云硬盘有效。
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitnil,omitempty" name:"DeleteWithInstance"`
}

type AttachDisksRequest struct {
	*tchttp.BaseRequest
	
	// EMR集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要挂载的云盘ID
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 挂载模式，取值范围：
	// AUTO_RENEW：自动续费
	// ALIGN_DEADLINE：自动对其到期时间
	AlignType *string `json:"AlignType,omitnil,omitempty" name:"AlignType"`

	// 需要挂载的cvm节点id列表
	CvmInstanceIds []*string `json:"CvmInstanceIds,omitnil,omitempty" name:"CvmInstanceIds"`

	// 是否是新购云盘进行挂载
	CreateDisk *bool `json:"CreateDisk,omitnil,omitempty" name:"CreateDisk"`

	// 新购云盘规格
	DiskSpec *NodeSpecDiskV2 `json:"DiskSpec,omitnil,omitempty" name:"DiskSpec"`

	// 可选参数，不传该参数则仅执行挂载操作。传入True时，会在挂载成功后将云硬盘设置为随云主机销毁模式，仅对按量计费云硬盘有效。
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitnil,omitempty" name:"DeleteWithInstance"`
}

func (r *AttachDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DiskIds")
	delete(f, "AlignType")
	delete(f, "CvmInstanceIds")
	delete(f, "CreateDisk")
	delete(f, "DiskSpec")
	delete(f, "DeleteWithInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachDisksResponseParams struct {
	// 流程id
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachDisksResponse struct {
	*tchttp.BaseResponse
	Response *AttachDisksResponseParams `json:"Response"`
}

func (r *AttachDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoScaleRecord struct {
	// 扩缩容规则名。
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// "SCALE_OUT"和"SCALE_IN"，分别表示扩容和缩容。
	ScaleAction *string `json:"ScaleAction,omitnil,omitempty" name:"ScaleAction"`

	// 取值为"SUCCESS","FAILED","PART_SUCCESS","IN_PROCESS"，分别表示成功、失败、部分成功和流程中。
	ActionStatus *string `json:"ActionStatus,omitnil,omitempty" name:"ActionStatus"`

	// 流程触发时间。
	ActionTime *string `json:"ActionTime,omitnil,omitempty" name:"ActionTime"`

	// 扩缩容相关描述信息。
	ScaleInfo *string `json:"ScaleInfo,omitnil,omitempty" name:"ScaleInfo"`

	// 只在ScaleAction为SCALE_OUT时有效。
	ExpectScaleNum *int64 `json:"ExpectScaleNum,omitnil,omitempty" name:"ExpectScaleNum"`

	// 流程结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 策略类型，按负载或者按时间，1表示负载伸缩，2表示时间伸缩
	StrategyType *int64 `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 扩容时所使用规格信息。
	SpecInfo *string `json:"SpecInfo,omitnil,omitempty" name:"SpecInfo"`

	// 补偿扩容，0表示不开启，1表示开启
	CompensateFlag *int64 `json:"CompensateFlag,omitnil,omitempty" name:"CompensateFlag"`

	// 补偿次数
	CompensateCount *int64 `json:"CompensateCount,omitnil,omitempty" name:"CompensateCount"`

	// 重试次数
	RetryCount *uint64 `json:"RetryCount,omitnil,omitempty" name:"RetryCount"`

	// 重试信息
	RetryInfo *string `json:"RetryInfo,omitnil,omitempty" name:"RetryInfo"`

	// 重试英文描述
	RetryEnReason *string `json:"RetryEnReason,omitnil,omitempty" name:"RetryEnReason"`

	// 重试描述
	RetryReason *string `json:"RetryReason,omitnil,omitempty" name:"RetryReason"`
}

type AutoScaleResourceConf struct {
	// 配置ID。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 集群实例ID。
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 自动扩缩容保留最小实例数。
	ScaleLowerBound *int64 `json:"ScaleLowerBound,omitnil,omitempty" name:"ScaleLowerBound"`

	// 自动扩缩容最大实例数。
	ScaleUpperBound *int64 `json:"ScaleUpperBound,omitnil,omitempty" name:"ScaleUpperBound"`

	// 扩容规则类型，1为按负载指标扩容规则，2为按时间扩容规则
	StrategyType *int64 `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 下次能可扩容时间。
	NextTimeCanScale *uint64 `json:"NextTimeCanScale,omitnil,omitempty" name:"NextTimeCanScale"`

	// 优雅缩容开关
	GraceDownFlag *bool `json:"GraceDownFlag,omitnil,omitempty" name:"GraceDownFlag"`

	// "CVM"表示规格全部使用CVM相关类型，"POD"表示规格使用容器相关类型,默认为"CVM"。
	HardwareType *string `json:"HardwareType,omitnil,omitempty" name:"HardwareType"`

	// "POSTPAY"表示只使用按量计费，"SPOT_FIRST"表示竞价实例优先，只有HardwareType为"HOST"时支持竞价实例优先，"POD"只支持纯按量计费。
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 竞价实例优先的场景下，按量计费资源数量的最低百分比，整数
	PostPayPercentMin *int64 `json:"PostPayPercentMin,omitnil,omitempty" name:"PostPayPercentMin"`

	// 预设资源类型为HOST时，支持勾选“资源不足时切换POD”；支持取消勾选；默认不勾选（0），勾选（1)
	ChangeToPod *int64 `json:"ChangeToPod,omitnil,omitempty" name:"ChangeToPod"`

	// 伸缩组名
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 标签
	YarnNodeLabel *string `json:"YarnNodeLabel,omitnil,omitempty" name:"YarnNodeLabel"`

	// 伸缩组状态
	GroupStatus *int64 `json:"GroupStatus,omitnil,omitempty" name:"GroupStatus"`

	// 并行伸缩 0关闭；1开启
	Parallel *int64 `json:"Parallel,omitnil,omitempty" name:"Parallel"`

	// 是否支持MNode
	EnableMNode *int64 `json:"EnableMNode,omitnil,omitempty" name:"EnableMNode"`
}

type BootstrapAction struct {
	// 脚本位置，支持cos上的文件，且只支持https协议。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 执行时间。
	// resourceAfter 表示在机器资源申请成功后执行。
	// clusterBefore 表示在集群初始化前执行。
	// clusterAfter 表示在集群初始化后执行。
	WhenRun *string `json:"WhenRun,omitnil,omitempty" name:"WhenRun"`

	// 脚本参数
	Args []*string `json:"Args,omitnil,omitempty" name:"Args"`
}

type CBSInstance struct {
	// 云硬盘ID
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 云硬盘类型
	DiskUsage *string `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// 云硬盘名称
	DiskName *string `json:"DiskName,omitnil,omitempty" name:"DiskName"`

	// 云硬盘大小
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 云盘介质类型
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 是否跟随实例删除
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitnil,omitempty" name:"DeleteWithInstance"`

	// 云硬盘收费类型
	DiskChargeType *string `json:"DiskChargeType,omitnil,omitempty" name:"DiskChargeType"`

	// 云硬盘运行状态
	DiskState *string `json:"DiskState,omitnil,omitempty" name:"DiskState"`

	// 是否自动续费
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 到期时间
	DeadlineTime *string `json:"DeadlineTime,omitnil,omitempty" name:"DeadlineTime"`

	// 云盘是否挂载到云主机上
	Attached *bool `json:"Attached,omitnil,omitempty" name:"Attached"`

	// 当前时间距离盘到期的天数
	DifferDaysOfDeadline *int64 `json:"DifferDaysOfDeadline,omitnil,omitempty" name:"DifferDaysOfDeadline"`

	// 该云盘当前被挂载到的CVM实例InstanceId
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// 云硬盘挂载的云主机ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 云盘是否为共享型云盘。
	Shareable *bool `json:"Shareable,omitnil,omitempty" name:"Shareable"`
}

type CLBSetting struct {
	// CLB类型，PUBLIC_IP表示支持公网CLB和INTERNAL_IP表示支持内网CLB字段 
	CLBType *string `json:"CLBType,omitnil,omitempty" name:"CLBType"`

	// Vpc和子网信息设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil,omitempty" name:"VPCSettings"`
}

type COSSettings struct {
	// COS SecretId
	CosSecretId *string `json:"CosSecretId,omitnil,omitempty" name:"CosSecretId"`

	// COS SecrectKey
	CosSecretKey *string `json:"CosSecretKey,omitnil,omitempty" name:"CosSecretKey"`

	// 日志存储在COS上的路径
	LogOnCosPath *string `json:"LogOnCosPath,omitnil,omitempty" name:"LogOnCosPath"`
}

type CapacityGlobalConfig struct {
	// 是否开启了标签调度
	EnableLabel *bool `json:"EnableLabel,omitnil,omitempty" name:"EnableLabel"`

	// 如果开启了标签调度，标签信息存放的路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelDir *string `json:"LabelDir,omitnil,omitempty" name:"LabelDir"`

	// 是否覆盖用户指定队列，为true表示覆盖。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueMappingOverride *bool `json:"QueueMappingOverride,omitnil,omitempty" name:"QueueMappingOverride"`

	// 高级设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultSettings []*DefaultSetting `json:"DefaultSettings,omitnil,omitempty" name:"DefaultSettings"`
}

type CdbInfo struct {
	// 数据库实例
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 数据库IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 数据库端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 数据库内存规格
	MemSize *int64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// 数据库磁盘规格
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 服务标识
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 申请时间
	ApplyTime *string `json:"ApplyTime,omitnil,omitempty" name:"ApplyTime"`

	// 付费类型
	PayType *int64 `json:"PayType,omitnil,omitempty" name:"PayType"`

	// 过期标识
	ExpireFlag *bool `json:"ExpireFlag,omitnil,omitempty" name:"ExpireFlag"`

	// 数据库状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 续费标识
	IsAutoRenew *int64 `json:"IsAutoRenew,omitnil,omitempty" name:"IsAutoRenew"`

	// 数据库字符串
	SerialNo *string `json:"SerialNo,omitnil,omitempty" name:"SerialNo"`

	// ZoneId
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// RegionId
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`
}

type CloudResource struct {
	// 组件角色名
	ComponentName *string `json:"ComponentName,omitnil,omitempty" name:"ComponentName"`

	// pod请求数量
	PodNumber *int64 `json:"PodNumber,omitnil,omitempty" name:"PodNumber"`

	// Cpu请求数量最大值
	LimitCpu *int64 `json:"LimitCpu,omitnil,omitempty" name:"LimitCpu"`

	// 内存请求数量最大值
	LimitMemory *int64 `json:"LimitMemory,omitnil,omitempty" name:"LimitMemory"`

	// 服务名称，如HIVE
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 数据卷目录设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeDir *VolumeSetting `json:"VolumeDir,omitnil,omitempty" name:"VolumeDir"`

	// 组件外部访问设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalAccess *ExternalAccess `json:"ExternalAccess,omitnil,omitempty" name:"ExternalAccess"`

	// 节点亲和性设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Affinity *NodeAffinity `json:"Affinity,omitnil,omitempty" name:"Affinity"`

	// 所选数据盘信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Disks []*Disk `json:"Disks,omitnil,omitempty" name:"Disks"`
}

type ClusterExternalServiceInfo struct {
	// 依赖关系，0:被其他集群依赖，1:依赖其他集群
	DependType *int64 `json:"DependType,omitnil,omitempty" name:"DependType"`

	// 共用组件
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 共用集群
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 共用集群状态
	ClusterStatus *int64 `json:"ClusterStatus,omitnil,omitempty" name:"ClusterStatus"`
}

type ClusterIDToFlowID struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 流程id
	FlowId *uint64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type ClusterInstancesInfo struct {
	// ID号
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Ftitle is deprecated.
	Ftitle *string `json:"Ftitle,omitnil,omitempty" name:"Ftitle"`

	// 集群名
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地区ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户APPID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 集群VPCID
	VpcId *int64 `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *int64 `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例的状态码。取值范围：
	// <li>2：表示集群运行中。</li>
	// <li>3：表示集群创建中。</li>
	// <li>4：表示集群扩容中。</li>
	// <li>5：表示集群增加router节点中。</li>
	// <li>6：表示集群安装组件中。</li>
	// <li>7：表示集群执行命令中。</li>
	// <li>8：表示重启服务中。</li>
	// <li>9：表示进入维护中。</li>
	// <li>10：表示服务暂停中。</li>
	// <li>11：表示退出维护中。</li>
	// <li>12：表示退出暂停中。</li>
	// <li>13：表示配置下发中。</li>
	// <li>14：表示销毁集群中。</li>
	// <li>15：表示销毁core节点中。</li>
	// <li>16：销毁task节点中。</li>
	// <li>17：表示销毁router节点中。</li>
	// <li>18：表示更改webproxy密码中。</li>
	// <li>19：表示集群隔离中。</li>
	// <li>20：表示集群冲正中。</li>
	// <li>21：表示集群回收中。</li>
	// <li>22：表示变配等待中。</li>
	// <li>23：表示集群已隔离。</li>
	// <li>24：表示缩容节点中。</li>
	// <li>33：表示集群等待退费中。</li>
	// <li>34：表示集群已退费。</li>
	// <li>301：表示创建失败。</li>
	// <li>302：表示扩容失败。</li>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 添加时间
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 已经运行时间
	RunTime *string `json:"RunTime,omitnil,omitempty" name:"RunTime"`

	// 集群产品配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Config is deprecated.
	Config *EmrProductConfigOutter `json:"Config,omitnil,omitempty" name:"Config"`

	// 主节点外网IP
	MasterIp *string `json:"MasterIp,omitnil,omitempty" name:"MasterIp"`

	// EMR版本
	EmrVersion *string `json:"EmrVersion,omitnil,omitempty" name:"EmrVersion"`

	// 收费类型
	ChargeType *int64 `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 交易版本
	TradeVersion *int64 `json:"TradeVersion,omitnil,omitempty" name:"TradeVersion"`

	// 资源订单ID
	ResourceOrderId *int64 `json:"ResourceOrderId,omitnil,omitempty" name:"ResourceOrderId"`

	// 是否计费集群
	IsTradeCluster *int64 `json:"IsTradeCluster,omitnil,omitempty" name:"IsTradeCluster"`

	// 集群错误状态告警信息
	AlarmInfo *string `json:"AlarmInfo,omitnil,omitempty" name:"AlarmInfo"`

	// 是否采用新架构
	IsWoodpeckerCluster *int64 `json:"IsWoodpeckerCluster,omitnil,omitempty" name:"IsWoodpeckerCluster"`

	// 元数据库信息
	MetaDb *string `json:"MetaDb,omitnil,omitempty" name:"MetaDb"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Hive元数据信息
	HiveMetaDb *string `json:"HiveMetaDb,omitnil,omitempty" name:"HiveMetaDb"`

	// 集群类型:EMR,CLICKHOUSE,DRUID
	ServiceClass *string `json:"ServiceClass,omitnil,omitempty" name:"ServiceClass"`

	// 集群所有节点的别名序列化
	AliasInfo *string `json:"AliasInfo,omitnil,omitempty" name:"AliasInfo"`

	// 集群版本Id
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 地区ID
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 场景名称
	SceneName *string `json:"SceneName,omitnil,omitempty" name:"SceneName"`

	// 场景化集群类型
	SceneServiceClass *string `json:"SceneServiceClass,omitnil,omitempty" name:"SceneServiceClass"`

	// 场景化EMR版本
	SceneEmrVersion *string `json:"SceneEmrVersion,omitnil,omitempty" name:"SceneEmrVersion"`

	// 场景化集群类型
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// vpc name
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// subnet name
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// 集群依赖关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterExternalServiceInfo []*ClusterExternalServiceInfo `json:"ClusterExternalServiceInfo,omitnil,omitempty" name:"ClusterExternalServiceInfo"`

	// 集群vpcid 字符串类型
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 子网id 字符串类型
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopologyInfoList []*TopologyInfo `json:"TopologyInfoList,omitnil,omitempty" name:"TopologyInfoList"`

	// 是否是跨AZ集群
	IsMultiZoneCluster *bool `json:"IsMultiZoneCluster,omitnil,omitempty" name:"IsMultiZoneCluster"`

	// 是否开通异常节点自动补偿
	IsCvmReplace *bool `json:"IsCvmReplace,omitnil,omitempty" name:"IsCvmReplace"`

	// 标题
	ClusterTitle *string `json:"ClusterTitle,omitnil,omitempty" name:"ClusterTitle"`

	// 集群产品配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigDetail *EmrProductConfigDetail `json:"ConfigDetail,omitnil,omitempty" name:"ConfigDetail"`

	// 集群绑定的文件系统数
	BindFileSystemNum *int64 `json:"BindFileSystemNum,omitnil,omitempty" name:"BindFileSystemNum"`

	// rss集群的绑定列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterRelationInfoList []*ClusterRelationMeta `json:"ClusterRelationInfoList,omitnil,omitempty" name:"ClusterRelationInfoList"`
}

type ClusterRelationMeta struct {
	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterIdList []*string `json:"ClusterIdList,omitnil,omitempty" name:"ClusterIdList"`
}

type ClusterSetting struct {
	// 付费方式。
	// PREPAID 包年包月。
	// POSTPAID_BY_HOUR 按量计费，默认方式。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 是否为HA集群。
	SupportHA *bool `json:"SupportHA,omitnil,omitempty" name:"SupportHA"`

	// 集群所使用的安全组，目前仅支持一个。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 实例位置。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 实例所在VPC。
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil,omitempty" name:"VPCSettings"`

	// 实例登录配置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例标签，示例：["{\"TagKey\":\"test-tag1\",\"TagValue\":\"001\"}","{\"TagKey\":\"test-tag2\",\"TagValue\":\"002\"}"]。
	TagSpecification []*string `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 元数据库配置。
	MetaDB *MetaDbInfo `json:"MetaDB,omitnil,omitempty" name:"MetaDB"`

	// 实例硬件配置。
	ResourceSpec *JobFlowResourceSpec `json:"ResourceSpec,omitnil,omitempty" name:"ResourceSpec"`

	// 是否申请公网IP，默认为false。
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitnil,omitempty" name:"PublicIpAssigned"`

	// 包年包月配置，只对包年包月集群生效。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 集群置放群组。
	DisasterRecoverGroupIds *string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 是否使用cbs加密。
	CbsEncryptFlag *bool `json:"CbsEncryptFlag,omitnil,omitempty" name:"CbsEncryptFlag"`

	// 是否使用远程登录，默认为false。
	RemoteTcpDefaultPort *bool `json:"RemoteTcpDefaultPort,omitnil,omitempty" name:"RemoteTcpDefaultPort"`
}

type ComponentBasicRestartInfo struct {
	// 进程名，必填，如NameNode
	ComponentName *string `json:"ComponentName,omitnil,omitempty" name:"ComponentName"`

	// 操作的IP列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`
}

type ConfigModifyInfoV2 struct {
	// 操作类型，可选值：
	// 
	// - 0：新建队列
	// - 1：编辑-全量覆盖
	// - 2：新建子队列
	// - 3：删除
	// - 4：克隆，与新建子队列的行为一样，特别的对于`fair`，可以复制子队列到新建队列
	// - 6：编辑-增量更新
	OpType *uint64 `json:"OpType,omitnil,omitempty" name:"OpType"`

	// 队列名称，不支持修改。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 新建队列 传root的MyId；新建子队列 传 选中队列的 myId；克隆 要传 选中队列 parentId
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 编辑、删除 传选中队列的 myId。克隆只有在调度器是`fair`时才需要传，用来复制子队列到新队列。
	MyId *string `json:"MyId,omitnil,omitempty" name:"MyId"`

	// 基础配置信息。key的取值与**DescribeYarnQueue**返回的字段一致。
	// ###### 公平调度器
	// key的取值信息如下：
	// 
	// - type，父队列，取值为 **parent** 或 **null**
	// - aclSubmitApps，提交访问控制，取值为**AclForYarnQueue类型的json串**或**null**
	// - aclAdministerApps，管理访问控制，取值为**AclForYarnQueue类型的json串**或**null**
	// - minSharePreemptionTimeout，最小共享优先权超时时间，取值为**数字字符串**或**null**
	// - fairSharePreemptionTimeout，公平份额抢占超时时间，取值为**数字字符串**或**null**
	// - fairSharePreemptionThreshold，公平份额抢占阈值，取值为**数字字符串**或**null**，其中数字的范围是（0，1]
	// - allowPreemptionFrom，抢占模式，取值为**布尔字符串**或**null**
	// - schedulingPolicy，调度策略，取值为**drf**、**fair**、**fifo**或**null**
	// 
	// ```
	// type AclForYarnQueue struct {
	// 	User  *string `json:"user"` //用户名
	// 	Group *string `json:"group"`//组名
	// }
	// ```
	// ###### 容量调度器
	// key的取值信息如下：
	// 
	// - state，队列状态，取值为**STOPPED**或**RUNNING**
	// - default-node-label-expression，默认标签表达式，取值为**标签**或**null**
	// - acl_submit_applications，提交访问控制，取值为**AclForYarnQueue类型的json串**或**null**
	// - acl_administer_queue，管理访问控制，取值为**AclForYarnQueue类型的json串**或**null**
	// - maximum-allocation-mb，分配Container最大内存数量，取值为**数字字符串**或**null**
	// - maximum-allocation-vcores，Container最大vCore数量，取值为**数字字符串**或**null**
	// ```
	// type AclForYarnQueue struct {
	// 	User  *string `json:"user"` //用户名
	// 	Group *string `json:"group"`//组名
	// }
	// ```
	// 注意：此字段可能返回 null，表示取不到有效值。
	BasicParams *ItemSeq `json:"BasicParams,omitnil,omitempty" name:"BasicParams"`

	// 配置集信息，取值见该复杂类型的参数说明。配置集是计划模式在队列中表现，表示的是不同时间段不同的配置值，所有队列的配置集名称都一样，对于单个队列，每个配置集中的标签与参数都一样，只是参数值不同。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigSetParams []*ConfigSetInfo `json:"ConfigSetParams,omitnil,omitempty" name:"ConfigSetParams"`

	// 容量调度专用，`OpType`为`6`时才生效，表示要删除这个队列中的哪些标签。优先级高于ConfigSetParams中的LabelParams。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteLables []*string `json:"DeleteLables,omitnil,omitempty" name:"DeleteLables"`
}

type ConfigSetInfo struct {
	// 配置集名称
	ConfigSet *string `json:"ConfigSet,omitnil,omitempty" name:"ConfigSet"`

	// 容量调度器会使用，里面设置了标签相关的配置。key的取值与**DescribeYarnQueue**返回的字段一致。
	// key的取值信息如下：
	// - labelName，标签名称，标签管理里的标签。
	// - capacity，容量，取值为**数字字符串**
	// - maximum-capacity，最大容量，取值为**数字字符串**
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelParams []*ItemSeq `json:"LabelParams,omitnil,omitempty" name:"LabelParams"`

	// 设置配置集相关的参数。key的取值与**DescribeYarnQueue**返回的字段一致。
	// ###### 公平调度器
	// key的取值信息如下：
	// - minResources，最大资源量，取值为**YarnResource类型的json串**或**null**
	// - maxResources，最大资源量，取值为**YarnResource类型的json串**或**null**
	// - maxChildResources，能够分配给为未声明子队列的最大资源量，取值为**数字字符串**或**null**
	// - maxRunningApps，最高可同时处于运行的App数量，取值为**数字字符串**或**null**
	// - weight，权重，取值为**数字字符串**或**null**
	// - maxAMShare，App Master最大份额，取值为**数字字符串**或**null**，其中数字的范围是[0，1]或-1
	// 
	// ```
	// type YarnResource struct {
	// 	Vcores *int `json:"vcores"`
	// 	Memory *int `json:"memory"`
	// 	Type *string `json:"type"` // 取值为`percent`或`null`当值为`percent`时，表示使用的百分比，否则就是使用的绝对数值。只有maxResources、maxChildResources才可以取值为`percent`
	// }
	// ```
	// 
	// ###### 容量调度器
	// key的取值信息如下：
	// - minimum-user-limit-percent，用户最小容量，取值为**YarnResource类型的json串**或**null**，其中数字的范围是[0，100]
	// - user-limit-factor，用户资源因子，取值为**YarnResource类型的json串**或**null**
	// - maximum-applications，最大应用数Max-Applications，取值为**数字字符串**或**null**，其中数字为正整数
	// - maximum-am-resource-percent，最大AM比例，取值为**数字字符串**或**null**，其中数字的范围是[0，1]或-1
	// - default-application-priority，资源池优先级，取值为**数字字符串**或**null**，其中数字为正整数
	// 注意：此字段可能返回 null，表示取不到有效值。
	BasicParams []*Item `json:"BasicParams,omitnil,omitempty" name:"BasicParams"`
}

type Configuration struct {
	// 配置文件名，支持SPARK、HIVE、HDFS、YARN的部分配置文件自定义。
	Classification *string `json:"Classification,omitnil,omitempty" name:"Classification"`

	// 配置参数通过KV的形式传入，部分文件支持自定义，可以通过特殊的键"content"传入所有内容。
	Properties *string `json:"Properties,omitnil,omitempty" name:"Properties"`
}

// Predefined struct for user
type CreateCloudInstanceRequestParams struct {
	// 实例名称。
	// <li>长度限制为6-36个字符。</li>
	// <li>只允许包含中文、字母、数字、-、_。</li>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 容器集群类型，取值范围
	// <li>EMR容器集群实例: EMR-TKE</li>
	ClusterClass *string `json:"ClusterClass,omitnil,omitempty" name:"ClusterClass"`

	// 部署的组件列表，不同的EMR产品ID（ProductId：具体含义参考入参ProductId字段）对应不同可选组件列表，不同产品版本可选组件列表查询：[组件版本](https://cloud.tencent.com/document/product/589/20279) ；
	Software []*string `json:"Software,omitnil,omitempty" name:"Software"`

	// 容器平台类型，取值范围
	// <li>EMR容器集群实例: tke</li>
	PlatFormType *string `json:"PlatFormType,omitnil,omitempty" name:"PlatFormType"`

	// cos存储桶
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// 容器集群id
	EksClusterId *string `json:"EksClusterId,omitnil,omitempty" name:"EksClusterId"`

	// 产品Id，不同产品ID表示不同的EMR产品版本。取值范围：
	// <li>60:表示EMR-TKE-V1.1.0</li>
	// <li>55:表示EMR-TKE-V1.0.1</li>
	// <li>52:表示EMR-TKE-V1.0.0</li>
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 客户端token，唯一随机标识，时效5分钟，需要调用者指定 防止客户端重新创建资源，小于等于64个字符，例如 a9a90aa6----fae36063280
	// 示例值：a9a90aa6----fae36063280
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil,omitempty" name:"VPCSettings"`

	// 所有组件角色及其对应的Pod资源请求信息
	CloudResources []*CloudResource `json:"CloudResources,omitnil,omitempty" name:"CloudResources"`

	// 安全组Id，为空默认创建新的安全组
	SgId *string `json:"SgId,omitnil,omitempty" name:"SgId"`

	// 元数据库信息
	// MetaDB信息，当MetaType选择EMR_NEW_META时，MetaDataJdbcUrl MetaDataUser MetaDataPass UnifyMetaInstanceId不用填
	// 当MetaType选择EMR_EXIT_META时，填写UnifyMetaInstanceId
	// 当MetaType选择USER_CUSTOM_META时，填写MetaDataJdbcUrl MetaDataUser MetaDataPass
	MetaDBInfo *CustomMetaDBInfo `json:"MetaDBInfo,omitnil,omitempty" name:"MetaDBInfo"`

	// 标签信息
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 登陆密码，LoginSettings中的Password字段
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 共享服务信息
	ExternalService []*ExternalService `json:"ExternalService,omitnil,omitempty" name:"ExternalService"`

	// 可用区id
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type CreateCloudInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例名称。
	// <li>长度限制为6-36个字符。</li>
	// <li>只允许包含中文、字母、数字、-、_。</li>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 容器集群类型，取值范围
	// <li>EMR容器集群实例: EMR-TKE</li>
	ClusterClass *string `json:"ClusterClass,omitnil,omitempty" name:"ClusterClass"`

	// 部署的组件列表，不同的EMR产品ID（ProductId：具体含义参考入参ProductId字段）对应不同可选组件列表，不同产品版本可选组件列表查询：[组件版本](https://cloud.tencent.com/document/product/589/20279) ；
	Software []*string `json:"Software,omitnil,omitempty" name:"Software"`

	// 容器平台类型，取值范围
	// <li>EMR容器集群实例: tke</li>
	PlatFormType *string `json:"PlatFormType,omitnil,omitempty" name:"PlatFormType"`

	// cos存储桶
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// 容器集群id
	EksClusterId *string `json:"EksClusterId,omitnil,omitempty" name:"EksClusterId"`

	// 产品Id，不同产品ID表示不同的EMR产品版本。取值范围：
	// <li>60:表示EMR-TKE-V1.1.0</li>
	// <li>55:表示EMR-TKE-V1.0.1</li>
	// <li>52:表示EMR-TKE-V1.0.0</li>
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 客户端token，唯一随机标识，时效5分钟，需要调用者指定 防止客户端重新创建资源，小于等于64个字符，例如 a9a90aa6----fae36063280
	// 示例值：a9a90aa6----fae36063280
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil,omitempty" name:"VPCSettings"`

	// 所有组件角色及其对应的Pod资源请求信息
	CloudResources []*CloudResource `json:"CloudResources,omitnil,omitempty" name:"CloudResources"`

	// 安全组Id，为空默认创建新的安全组
	SgId *string `json:"SgId,omitnil,omitempty" name:"SgId"`

	// 元数据库信息
	// MetaDB信息，当MetaType选择EMR_NEW_META时，MetaDataJdbcUrl MetaDataUser MetaDataPass UnifyMetaInstanceId不用填
	// 当MetaType选择EMR_EXIT_META时，填写UnifyMetaInstanceId
	// 当MetaType选择USER_CUSTOM_META时，填写MetaDataJdbcUrl MetaDataUser MetaDataPass
	MetaDBInfo *CustomMetaDBInfo `json:"MetaDBInfo,omitnil,omitempty" name:"MetaDBInfo"`

	// 标签信息
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 登陆密码，LoginSettings中的Password字段
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 共享服务信息
	ExternalService []*ExternalService `json:"ExternalService,omitnil,omitempty" name:"ExternalService"`

	// 可用区id
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *CreateCloudInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "ClusterClass")
	delete(f, "Software")
	delete(f, "PlatFormType")
	delete(f, "CosBucket")
	delete(f, "EksClusterId")
	delete(f, "ProductId")
	delete(f, "ClientToken")
	delete(f, "VPCSettings")
	delete(f, "CloudResources")
	delete(f, "SgId")
	delete(f, "MetaDBInfo")
	delete(f, "Tags")
	delete(f, "LoginSettings")
	delete(f, "ExternalService")
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudInstanceResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloudInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudInstanceResponseParams `json:"Response"`
}

func (r *CreateCloudInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterRequestParams struct {
	// EMR产品版本名称如EMR-V2.3.0 表示2.3.0版本的EMR， 当前支持产品版本名称查询：[产品版本名称](https://cloud.tencent.com/document/product/589/66338)
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// 是否开启节点高可用。取值范围：
	// <li>true：表示开启节点高可用。</li>
	// <li>false：表示不开启节点高可用。</li>
	EnableSupportHAFlag *bool `json:"EnableSupportHAFlag,omitnil,omitempty" name:"EnableSupportHAFlag"`

	// 实例名称。
	// <li>长度限制为6-36个字符。</li>
	// <li>只允许包含中文、字母、数字、-、_。</li>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例计费模式。取值范围：
	// <li>PREPAID：预付费，即包年包月。</li>
	// <li>POSTPAID_BY_HOUR：按小时后付费。</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例登录设置。通过该参数可以设置所购买节点的登录方式密码或者密钥。
	// <li>设置密钥时，密码仅用于组件原生WebUI快捷入口登录。</li>
	// <li>未设置密钥时，密码用于登录所购节点以及组件原生WebUI快捷入口登录。</li>
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 集群应用场景以及支持部署组件配置
	SceneSoftwareConfig *SceneSoftwareConfig `json:"SceneSoftwareConfig,omitnil,omitempty" name:"SceneSoftwareConfig"`

	// 即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 实例所属安全组的ID，形如sg-xxxxxxxx。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的SecurityGroupId字段来获取。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// [引导操作](https://cloud.tencent.com/document/product/589/35656)脚本设置。
	ScriptBootstrapActionConfig []*ScriptBootstrapActionConfig `json:"ScriptBootstrapActionConfig,omitnil,omitempty" name:"ScriptBootstrapActionConfig"`

	// 唯一随机标识，时效性为5分钟，需要调用者指定 防止客户端重复创建资源，例如 a9a90aa6-****-****-****-fae360632808
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否开启集群Master节点公网。取值范围：
	// <li>NEED_MASTER_WAN：表示开启集群Master节点公网。</li>
	// <li>NOT_NEED_MASTER_WAN：表示不开启。</li>默认开启集群Master节点公网。
	NeedMasterWan *string `json:"NeedMasterWan,omitnil,omitempty" name:"NeedMasterWan"`

	// 是否开启外网远程登录。（在SecurityGroupId不为空时，该参数无效）不填默认为不开启 取值范围：
	// <li>true：表示开启</li>
	// <li>false：表示不开启</li>
	EnableRemoteLoginFlag *bool `json:"EnableRemoteLoginFlag,omitnil,omitempty" name:"EnableRemoteLoginFlag"`

	// 是否开启Kerberos认证。默认不开启 取值范围：
	// <li>true：表示开启</li>
	// <li>false：表示不开启</li>
	EnableKerberosFlag *bool `json:"EnableKerberosFlag,omitnil,omitempty" name:"EnableKerberosFlag"`

	// [自定义软件配置](https://cloud.tencent.com/document/product/589/35655?from_cn_redirect=1)
	CustomConf *string `json:"CustomConf,omitnil,omitempty" name:"CustomConf"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的实例。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 分散置放群组ID列表，当前只支持指定一个。
	// 该参数可以通过调用 [DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/product/213/17810)的返回值中的DisasterRecoverGroupId字段来获取。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 是否开启集群维度CBS加密。默认不加密 取值范围：
	// <li>true：表示加密</li>
	// <li>false：表示不加密</li>
	EnableCbsEncryptFlag *bool `json:"EnableCbsEncryptFlag,omitnil,omitempty" name:"EnableCbsEncryptFlag"`

	// MetaDB信息，当MetaType选择EMR_NEW_META时，MetaDataJdbcUrl MetaDataUser MetaDataPass UnifyMetaInstanceId不用填
	// 当MetaType选择EMR_EXIT_META时，填写UnifyMetaInstanceId
	// 当MetaType选择USER_CUSTOM_META时，填写MetaDataJdbcUrl MetaDataUser MetaDataPass
	MetaDBInfo *CustomMetaDBInfo `json:"MetaDBInfo,omitnil,omitempty" name:"MetaDBInfo"`

	// 共享组件信息
	DependService []*DependService `json:"DependService,omitnil,omitempty" name:"DependService"`

	// 节点资源的规格，有几个可用区，就填几个，按顺序第一个为主可用区，第二个为备可用区，第三个为仲裁可用区。如果没有开启跨AZ，则长度为1即可。
	ZoneResourceConfiguration []*ZoneResourceConfiguration `json:"ZoneResourceConfiguration,omitnil,omitempty" name:"ZoneResourceConfiguration"`

	// cos桶路径，创建StarRocks存算分离集群时用到
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest
	
	// EMR产品版本名称如EMR-V2.3.0 表示2.3.0版本的EMR， 当前支持产品版本名称查询：[产品版本名称](https://cloud.tencent.com/document/product/589/66338)
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// 是否开启节点高可用。取值范围：
	// <li>true：表示开启节点高可用。</li>
	// <li>false：表示不开启节点高可用。</li>
	EnableSupportHAFlag *bool `json:"EnableSupportHAFlag,omitnil,omitempty" name:"EnableSupportHAFlag"`

	// 实例名称。
	// <li>长度限制为6-36个字符。</li>
	// <li>只允许包含中文、字母、数字、-、_。</li>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例计费模式。取值范围：
	// <li>PREPAID：预付费，即包年包月。</li>
	// <li>POSTPAID_BY_HOUR：按小时后付费。</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例登录设置。通过该参数可以设置所购买节点的登录方式密码或者密钥。
	// <li>设置密钥时，密码仅用于组件原生WebUI快捷入口登录。</li>
	// <li>未设置密钥时，密码用于登录所购节点以及组件原生WebUI快捷入口登录。</li>
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 集群应用场景以及支持部署组件配置
	SceneSoftwareConfig *SceneSoftwareConfig `json:"SceneSoftwareConfig,omitnil,omitempty" name:"SceneSoftwareConfig"`

	// 即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 实例所属安全组的ID，形如sg-xxxxxxxx。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的SecurityGroupId字段来获取。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// [引导操作](https://cloud.tencent.com/document/product/589/35656)脚本设置。
	ScriptBootstrapActionConfig []*ScriptBootstrapActionConfig `json:"ScriptBootstrapActionConfig,omitnil,omitempty" name:"ScriptBootstrapActionConfig"`

	// 唯一随机标识，时效性为5分钟，需要调用者指定 防止客户端重复创建资源，例如 a9a90aa6-****-****-****-fae360632808
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否开启集群Master节点公网。取值范围：
	// <li>NEED_MASTER_WAN：表示开启集群Master节点公网。</li>
	// <li>NOT_NEED_MASTER_WAN：表示不开启。</li>默认开启集群Master节点公网。
	NeedMasterWan *string `json:"NeedMasterWan,omitnil,omitempty" name:"NeedMasterWan"`

	// 是否开启外网远程登录。（在SecurityGroupId不为空时，该参数无效）不填默认为不开启 取值范围：
	// <li>true：表示开启</li>
	// <li>false：表示不开启</li>
	EnableRemoteLoginFlag *bool `json:"EnableRemoteLoginFlag,omitnil,omitempty" name:"EnableRemoteLoginFlag"`

	// 是否开启Kerberos认证。默认不开启 取值范围：
	// <li>true：表示开启</li>
	// <li>false：表示不开启</li>
	EnableKerberosFlag *bool `json:"EnableKerberosFlag,omitnil,omitempty" name:"EnableKerberosFlag"`

	// [自定义软件配置](https://cloud.tencent.com/document/product/589/35655?from_cn_redirect=1)
	CustomConf *string `json:"CustomConf,omitnil,omitempty" name:"CustomConf"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的实例。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 分散置放群组ID列表，当前只支持指定一个。
	// 该参数可以通过调用 [DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/product/213/17810)的返回值中的DisasterRecoverGroupId字段来获取。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 是否开启集群维度CBS加密。默认不加密 取值范围：
	// <li>true：表示加密</li>
	// <li>false：表示不加密</li>
	EnableCbsEncryptFlag *bool `json:"EnableCbsEncryptFlag,omitnil,omitempty" name:"EnableCbsEncryptFlag"`

	// MetaDB信息，当MetaType选择EMR_NEW_META时，MetaDataJdbcUrl MetaDataUser MetaDataPass UnifyMetaInstanceId不用填
	// 当MetaType选择EMR_EXIT_META时，填写UnifyMetaInstanceId
	// 当MetaType选择USER_CUSTOM_META时，填写MetaDataJdbcUrl MetaDataUser MetaDataPass
	MetaDBInfo *CustomMetaDBInfo `json:"MetaDBInfo,omitnil,omitempty" name:"MetaDBInfo"`

	// 共享组件信息
	DependService []*DependService `json:"DependService,omitnil,omitempty" name:"DependService"`

	// 节点资源的规格，有几个可用区，就填几个，按顺序第一个为主可用区，第二个为备可用区，第三个为仲裁可用区。如果没有开启跨AZ，则长度为1即可。
	ZoneResourceConfiguration []*ZoneResourceConfiguration `json:"ZoneResourceConfiguration,omitnil,omitempty" name:"ZoneResourceConfiguration"`

	// cos桶路径，创建StarRocks存算分离集群时用到
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`
}

func (r *CreateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductVersion")
	delete(f, "EnableSupportHAFlag")
	delete(f, "InstanceName")
	delete(f, "InstanceChargeType")
	delete(f, "LoginSettings")
	delete(f, "SceneSoftwareConfig")
	delete(f, "InstanceChargePrepaid")
	delete(f, "SecurityGroupIds")
	delete(f, "ScriptBootstrapActionConfig")
	delete(f, "ClientToken")
	delete(f, "NeedMasterWan")
	delete(f, "EnableRemoteLoginFlag")
	delete(f, "EnableKerberosFlag")
	delete(f, "CustomConf")
	delete(f, "Tags")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "EnableCbsEncryptFlag")
	delete(f, "MetaDBInfo")
	delete(f, "DependService")
	delete(f, "ZoneResourceConfiguration")
	delete(f, "CosBucket")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterResponseParams `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceRequestParams struct {
	// 产品ID，不同产品ID表示不同的EMR产品版本。取值范围：
	// 51:表示STARROCKS-V1.4.0
	// 54:表示STARROCKS-V2.0.0
	// 27:表示KAFKA-V1.0.0
	// 50:表示KAFKA-V2.0.0
	// 16:表示EMR-V2.3.0
	// 20:表示EMR-V2.5.0
	// 30:表示EMR-V2.6.0
	// 38:表示EMR-V2.7.0
	// 25:表示EMR-V3.1.0
	// 33:表示EMR-V3.2.1
	// 34:表示EMR-V3.3.0
	// 37:表示EMR-V3.4.0
	// 44:表示EMR-V3.5.0
	// 53:表示EMR-V3.6.0
	ProductId *uint64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 部署的组件列表。不同的EMR产品ID（ProductId：具体含义参考入参ProductId字段）对应不同可选组件列表，不同产品版本可选组件列表查询：[组件版本](https://cloud.tencent.com/document/product/589/20279) ；
	// 填写实例值：hive、flink。
	Software []*string `json:"Software,omitnil,omitempty" name:"Software"`

	// 是否开启节点高可用。取值范围：
	// <li>0：表示不开启节点高可用。</li>
	// <li>1：表示开启节点高可用。</li>
	SupportHA *uint64 `json:"SupportHA,omitnil,omitempty" name:"SupportHA"`

	// 实例名称。
	// <li>长度限制为6-36个字符。</li>
	// <li>只允许包含中文、字母、数字、-、_。</li>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 购买实例的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 购买实例的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 实例登录设置。通过该参数可以设置所购买节点的登录方式密码或者密钥。
	// <li>设置密钥时，密码仅用于组件原生WebUI快捷入口登录。</li>
	// <li>未设置密钥时，密码用于登录所购节点以及组件原生WebUI快捷入口登录。</li>
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil,omitempty" name:"VPCSettings"`

	// 节点资源的规格。
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitnil,omitempty" name:"ResourceSpec"`

	// 开启COS访问需要设置的参数。
	COSSettings *COSSettings `json:"COSSettings,omitnil,omitempty" name:"COSSettings"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 实例所属安全组的ID，形如sg-xxxxxxxx。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的SecurityGroupId字段来获取。
	SgId *string `json:"SgId,omitnil,omitempty" name:"SgId"`

	// [引导操作](https://cloud.tencent.com/document/product/589/35656)脚本设置。
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitnil,omitempty" name:"PreExecutedFileSettings"`

	// 包年包月实例是否自动续费。取值范围：
	// <li>0：表示不自动续费。</li>
	// <li>1：表示自动续费。</li>
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 唯一随机标识，时效5分钟，需要调用者指定 防止客户端重新创建资源，例如 a9a90aa6-****-****-****-fae36063280
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否开启集群Master节点公网。取值范围：
	// <li>NEED_MASTER_WAN：表示开启集群Master节点公网。</li>
	// <li>NOT_NEED_MASTER_WAN：表示不开启。</li>默认开启集群Master节点公网。
	NeedMasterWan *string `json:"NeedMasterWan,omitnil,omitempty" name:"NeedMasterWan"`

	// 是否需要开启外网远程登录，即22号端口。在SgId不为空时，该参数无效。
	RemoteLoginAtCreate *int64 `json:"RemoteLoginAtCreate,omitnil,omitempty" name:"RemoteLoginAtCreate"`

	// 是否开启安全集群。0表示不开启，非0表示开启。
	CheckSecurity *int64 `json:"CheckSecurity,omitnil,omitempty" name:"CheckSecurity"`

	// 访问外部文件系统。
	ExtendFsField *string `json:"ExtendFsField,omitnil,omitempty" name:"ExtendFsField"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的实例。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 分散置放群组ID列表，当前只支持指定一个。
	// 该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/product/213/15486 ) 的返回值中的SecurityGroupId字段来获取。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 集群维度CBS加密盘，默认0表示不加密，1表示加密
	CbsEncrypt *uint64 `json:"CbsEncrypt,omitnil,omitempty" name:"CbsEncrypt"`

	// hive共享元数据库类型。取值范围：
	// <li>EMR_DEFAULT_META：表示集群默认创建</li>
	// <li>EMR_EXIST_META：表示集群使用指定EMR-MetaDB。</li>
	// <li>USER_CUSTOM_META：表示集群使用自定义MetaDB。</li>
	MetaType *string `json:"MetaType,omitnil,omitempty" name:"MetaType"`

	// EMR-MetaDB实例
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitnil,omitempty" name:"UnifyMetaInstanceId"`

	// 自定义MetaDB信息
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitnil,omitempty" name:"MetaDBInfo"`

	// 自定义应用角色。
	ApplicationRole *string `json:"ApplicationRole,omitnil,omitempty" name:"ApplicationRole"`

	// 场景化取值：
	// Hadoop-Kudu
	// Hadoop-Zookeeper
	// Hadoop-Presto
	// Hadoop-Hbase
	SceneName *string `json:"SceneName,omitnil,omitempty" name:"SceneName"`

	// 共享组件信息
	ExternalService []*ExternalService `json:"ExternalService,omitnil,omitempty" name:"ExternalService"`

	// 如果为0，则MultiZone、MultiDeployStrategy、MultiZoneSettings是disable的状态，如果为1，则废弃ResourceSpec，使用MultiZoneSettings。
	VersionID *int64 `json:"VersionID,omitnil,omitempty" name:"VersionID"`

	// true表示开启跨AZ部署；仅为新建集群时的用户参数，后续不支持调整。
	MultiZone *bool `json:"MultiZone,omitnil,omitempty" name:"MultiZone"`

	// 节点资源的规格，有几个可用区，就填几个，按顺序第一个为主可用区，第二个为备可用区，第三个为仲裁可用区。如果没有开启跨AZ，则长度为1即可。
	MultiZoneSettings []*MultiZoneSetting `json:"MultiZoneSettings,omitnil,omitempty" name:"MultiZoneSettings"`

	// cos桶路径，创建StarRocks存算分离集群时用到
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID，不同产品ID表示不同的EMR产品版本。取值范围：
	// 51:表示STARROCKS-V1.4.0
	// 54:表示STARROCKS-V2.0.0
	// 27:表示KAFKA-V1.0.0
	// 50:表示KAFKA-V2.0.0
	// 16:表示EMR-V2.3.0
	// 20:表示EMR-V2.5.0
	// 30:表示EMR-V2.6.0
	// 38:表示EMR-V2.7.0
	// 25:表示EMR-V3.1.0
	// 33:表示EMR-V3.2.1
	// 34:表示EMR-V3.3.0
	// 37:表示EMR-V3.4.0
	// 44:表示EMR-V3.5.0
	// 53:表示EMR-V3.6.0
	ProductId *uint64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 部署的组件列表。不同的EMR产品ID（ProductId：具体含义参考入参ProductId字段）对应不同可选组件列表，不同产品版本可选组件列表查询：[组件版本](https://cloud.tencent.com/document/product/589/20279) ；
	// 填写实例值：hive、flink。
	Software []*string `json:"Software,omitnil,omitempty" name:"Software"`

	// 是否开启节点高可用。取值范围：
	// <li>0：表示不开启节点高可用。</li>
	// <li>1：表示开启节点高可用。</li>
	SupportHA *uint64 `json:"SupportHA,omitnil,omitempty" name:"SupportHA"`

	// 实例名称。
	// <li>长度限制为6-36个字符。</li>
	// <li>只允许包含中文、字母、数字、-、_。</li>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 购买实例的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 购买实例的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 实例登录设置。通过该参数可以设置所购买节点的登录方式密码或者密钥。
	// <li>设置密钥时，密码仅用于组件原生WebUI快捷入口登录。</li>
	// <li>未设置密钥时，密码用于登录所购节点以及组件原生WebUI快捷入口登录。</li>
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil,omitempty" name:"VPCSettings"`

	// 节点资源的规格。
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitnil,omitempty" name:"ResourceSpec"`

	// 开启COS访问需要设置的参数。
	COSSettings *COSSettings `json:"COSSettings,omitnil,omitempty" name:"COSSettings"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 实例所属安全组的ID，形如sg-xxxxxxxx。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的SecurityGroupId字段来获取。
	SgId *string `json:"SgId,omitnil,omitempty" name:"SgId"`

	// [引导操作](https://cloud.tencent.com/document/product/589/35656)脚本设置。
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitnil,omitempty" name:"PreExecutedFileSettings"`

	// 包年包月实例是否自动续费。取值范围：
	// <li>0：表示不自动续费。</li>
	// <li>1：表示自动续费。</li>
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 唯一随机标识，时效5分钟，需要调用者指定 防止客户端重新创建资源，例如 a9a90aa6-****-****-****-fae36063280
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否开启集群Master节点公网。取值范围：
	// <li>NEED_MASTER_WAN：表示开启集群Master节点公网。</li>
	// <li>NOT_NEED_MASTER_WAN：表示不开启。</li>默认开启集群Master节点公网。
	NeedMasterWan *string `json:"NeedMasterWan,omitnil,omitempty" name:"NeedMasterWan"`

	// 是否需要开启外网远程登录，即22号端口。在SgId不为空时，该参数无效。
	RemoteLoginAtCreate *int64 `json:"RemoteLoginAtCreate,omitnil,omitempty" name:"RemoteLoginAtCreate"`

	// 是否开启安全集群。0表示不开启，非0表示开启。
	CheckSecurity *int64 `json:"CheckSecurity,omitnil,omitempty" name:"CheckSecurity"`

	// 访问外部文件系统。
	ExtendFsField *string `json:"ExtendFsField,omitnil,omitempty" name:"ExtendFsField"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的实例。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 分散置放群组ID列表，当前只支持指定一个。
	// 该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/product/213/15486 ) 的返回值中的SecurityGroupId字段来获取。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 集群维度CBS加密盘，默认0表示不加密，1表示加密
	CbsEncrypt *uint64 `json:"CbsEncrypt,omitnil,omitempty" name:"CbsEncrypt"`

	// hive共享元数据库类型。取值范围：
	// <li>EMR_DEFAULT_META：表示集群默认创建</li>
	// <li>EMR_EXIST_META：表示集群使用指定EMR-MetaDB。</li>
	// <li>USER_CUSTOM_META：表示集群使用自定义MetaDB。</li>
	MetaType *string `json:"MetaType,omitnil,omitempty" name:"MetaType"`

	// EMR-MetaDB实例
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitnil,omitempty" name:"UnifyMetaInstanceId"`

	// 自定义MetaDB信息
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitnil,omitempty" name:"MetaDBInfo"`

	// 自定义应用角色。
	ApplicationRole *string `json:"ApplicationRole,omitnil,omitempty" name:"ApplicationRole"`

	// 场景化取值：
	// Hadoop-Kudu
	// Hadoop-Zookeeper
	// Hadoop-Presto
	// Hadoop-Hbase
	SceneName *string `json:"SceneName,omitnil,omitempty" name:"SceneName"`

	// 共享组件信息
	ExternalService []*ExternalService `json:"ExternalService,omitnil,omitempty" name:"ExternalService"`

	// 如果为0，则MultiZone、MultiDeployStrategy、MultiZoneSettings是disable的状态，如果为1，则废弃ResourceSpec，使用MultiZoneSettings。
	VersionID *int64 `json:"VersionID,omitnil,omitempty" name:"VersionID"`

	// true表示开启跨AZ部署；仅为新建集群时的用户参数，后续不支持调整。
	MultiZone *bool `json:"MultiZone,omitnil,omitempty" name:"MultiZone"`

	// 节点资源的规格，有几个可用区，就填几个，按顺序第一个为主可用区，第二个为备可用区，第三个为仲裁可用区。如果没有开启跨AZ，则长度为1即可。
	MultiZoneSettings []*MultiZoneSetting `json:"MultiZoneSettings,omitnil,omitempty" name:"MultiZoneSettings"`

	// cos桶路径，创建StarRocks存算分离集群时用到
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Software")
	delete(f, "SupportHA")
	delete(f, "InstanceName")
	delete(f, "PayMode")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "LoginSettings")
	delete(f, "VPCSettings")
	delete(f, "ResourceSpec")
	delete(f, "COSSettings")
	delete(f, "Placement")
	delete(f, "SgId")
	delete(f, "PreExecutedFileSettings")
	delete(f, "AutoRenew")
	delete(f, "ClientToken")
	delete(f, "NeedMasterWan")
	delete(f, "RemoteLoginAtCreate")
	delete(f, "CheckSecurity")
	delete(f, "ExtendFsField")
	delete(f, "Tags")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "CbsEncrypt")
	delete(f, "MetaType")
	delete(f, "UnifyMetaInstanceId")
	delete(f, "MetaDBInfo")
	delete(f, "ApplicationRole")
	delete(f, "SceneName")
	delete(f, "ExternalService")
	delete(f, "VersionID")
	delete(f, "MultiZone")
	delete(f, "MultiZoneSettings")
	delete(f, "CosBucket")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceResponseParams struct {
	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceResponseParams `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSLInstanceRequestParams struct {
	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例计费模式，0表示后付费，即按量计费。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 实例存储类型，填写CLOUD_HSSD，表示性能云存储。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 实例单节点磁盘容量，单位GB，单节点磁盘容量需大于等于100，小于等于250*CPU核心数，容量调整步长为100。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 实例节点规格，可填写4C16G、8C32G、16C64G、32C128G，不区分大小写。
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 实例可用区详细配置，当前支持多可用区，可用区数量只能为1或3，包含区域名称，VPC信息、节点数量，其中所有区域节点总数需大于等于3，小于等于50。
	ZoneSettings []*ZoneSetting `json:"ZoneSettings,omitnil,omitempty" name:"ZoneSettings"`

	// 实例要绑定的标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 预付费参数
	PrePaySetting *PrePaySetting `json:"PrePaySetting,omitnil,omitempty" name:"PrePaySetting"`
}

type CreateSLInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例计费模式，0表示后付费，即按量计费。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 实例存储类型，填写CLOUD_HSSD，表示性能云存储。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 实例单节点磁盘容量，单位GB，单节点磁盘容量需大于等于100，小于等于250*CPU核心数，容量调整步长为100。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 实例节点规格，可填写4C16G、8C32G、16C64G、32C128G，不区分大小写。
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 实例可用区详细配置，当前支持多可用区，可用区数量只能为1或3，包含区域名称，VPC信息、节点数量，其中所有区域节点总数需大于等于3，小于等于50。
	ZoneSettings []*ZoneSetting `json:"ZoneSettings,omitnil,omitempty" name:"ZoneSettings"`

	// 实例要绑定的标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 预付费参数
	PrePaySetting *PrePaySetting `json:"PrePaySetting,omitnil,omitempty" name:"PrePaySetting"`
}

func (r *CreateSLInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSLInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "PayMode")
	delete(f, "DiskType")
	delete(f, "DiskSize")
	delete(f, "NodeType")
	delete(f, "ZoneSettings")
	delete(f, "Tags")
	delete(f, "PrePaySetting")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSLInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSLInstanceResponseParams struct {
	// 实例唯一标识符（字符串表示）
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSLInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateSLInstanceResponseParams `json:"Response"`
}

func (r *CreateSLInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSLInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomMetaDBInfo struct {
	// 自定义MetaDB的JDBC连接，示例: jdbc:mysql://10.10.10.10:3306/dbname
	MetaDataJdbcUrl *string `json:"MetaDataJdbcUrl,omitnil,omitempty" name:"MetaDataJdbcUrl"`

	// 自定义MetaDB用户名
	MetaDataUser *string `json:"MetaDataUser,omitnil,omitempty" name:"MetaDataUser"`

	// 自定义MetaDB密码
	MetaDataPass *string `json:"MetaDataPass,omitnil,omitempty" name:"MetaDataPass"`

	// hive共享元数据库类型。取值范围：
	// <li>EMR_DEFAULT_META：表示集群默认创建</li>
	// <li>EMR_EXIST_META：表示集群使用指定EMR-MetaDB。</li>
	// <li>USER_CUSTOM_META：表示集群使用自定义MetaDB。</li>
	MetaType *string `json:"MetaType,omitnil,omitempty" name:"MetaType"`

	// EMR-MetaDB实例
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitnil,omitempty" name:"UnifyMetaInstanceId"`
}

type CustomMetaInfo struct {
	// 自定义MetaDB的JDBC连接，请以 jdbc:mysql:// 开头
	MetaDataJdbcUrl *string `json:"MetaDataJdbcUrl,omitnil,omitempty" name:"MetaDataJdbcUrl"`

	// 自定义MetaDB用户名
	MetaDataUser *string `json:"MetaDataUser,omitnil,omitempty" name:"MetaDataUser"`

	// 自定义MetaDB密码
	MetaDataPass *string `json:"MetaDataPass,omitnil,omitempty" name:"MetaDataPass"`
}

type CustomServiceDefine struct {
	// 自定义参数key
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义参数value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type DAGInfo struct {
	// 查询ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// DAG类型，目前只支持starrocks
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 返回的DAG的JSON字符串
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type DayRepeatStrategy struct {
	// 重复任务执行的具体时刻，例如"01:02:00"
	ExecuteAtTimeOfDay *string `json:"ExecuteAtTimeOfDay,omitnil,omitempty" name:"ExecuteAtTimeOfDay"`

	// 每隔Step天执行一次
	Step *uint64 `json:"Step,omitnil,omitempty" name:"Step"`
}

type DefaultSetting struct {
	// 名称，作为入参的key
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 提示
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// key，用于展示，该配置对应与配置文件中的配置项
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Name对应的值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type DeleteAutoScaleStrategyRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自动扩缩容规则类型，1表示按照负载指标扩缩容，2表示按照时间规则扩缩容。
	StrategyType *int64 `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 规则ID。
	StrategyId *int64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 伸缩组Id
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DeleteAutoScaleStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自动扩缩容规则类型，1表示按照负载指标扩缩容，2表示按照时间规则扩缩容。
	StrategyType *int64 `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 规则ID。
	StrategyId *int64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 伸缩组Id
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *DeleteAutoScaleStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoScaleStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StrategyType")
	delete(f, "StrategyId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAutoScaleStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAutoScaleStrategyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAutoScaleStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAutoScaleStrategyResponseParams `json:"Response"`
}

func (r *DeleteAutoScaleStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoScaleStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNodeResourceConfigRequestParams struct {
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点配置Id
	ResourceConfigId *uint64 `json:"ResourceConfigId,omitnil,omitempty" name:"ResourceConfigId"`

	// 节点类型 CORE TASK ROUTER
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 类型为ComputeResource和EMR以及默认，默认为EMR
	ResourceBaseType *string `json:"ResourceBaseType,omitnil,omitempty" name:"ResourceBaseType"`

	// 计算资源id
	ComputeResourceId *string `json:"ComputeResourceId,omitnil,omitempty" name:"ComputeResourceId"`
}

type DeleteNodeResourceConfigRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点配置Id
	ResourceConfigId *uint64 `json:"ResourceConfigId,omitnil,omitempty" name:"ResourceConfigId"`

	// 节点类型 CORE TASK ROUTER
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 类型为ComputeResource和EMR以及默认，默认为EMR
	ResourceBaseType *string `json:"ResourceBaseType,omitnil,omitempty" name:"ResourceBaseType"`

	// 计算资源id
	ComputeResourceId *string `json:"ComputeResourceId,omitnil,omitempty" name:"ComputeResourceId"`
}

func (r *DeleteNodeResourceConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNodeResourceConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceConfigId")
	delete(f, "ResourceType")
	delete(f, "ResourceBaseType")
	delete(f, "ComputeResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNodeResourceConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNodeResourceConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNodeResourceConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNodeResourceConfigResponseParams `json:"Response"`
}

func (r *DeleteNodeResourceConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNodeResourceConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserManagerUserListRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群用户名列表
	UserNameList []*string `json:"UserNameList,omitnil,omitempty" name:"UserNameList"`

	// tke/eks集群id，容器集群传
	TkeClusterId *string `json:"TkeClusterId,omitnil,omitempty" name:"TkeClusterId"`

	// 默认空，容器版传"native"
	DisplayStrategy *string `json:"DisplayStrategy,omitnil,omitempty" name:"DisplayStrategy"`

	// 用户组
	UserGroupList []*UserAndGroup `json:"UserGroupList,omitnil,omitempty" name:"UserGroupList"`

	// 是否删除家目录，只针对cvm集群
	DeleteHomeDir *bool `json:"DeleteHomeDir,omitnil,omitempty" name:"DeleteHomeDir"`
}

type DeleteUserManagerUserListRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群用户名列表
	UserNameList []*string `json:"UserNameList,omitnil,omitempty" name:"UserNameList"`

	// tke/eks集群id，容器集群传
	TkeClusterId *string `json:"TkeClusterId,omitnil,omitempty" name:"TkeClusterId"`

	// 默认空，容器版传"native"
	DisplayStrategy *string `json:"DisplayStrategy,omitnil,omitempty" name:"DisplayStrategy"`

	// 用户组
	UserGroupList []*UserAndGroup `json:"UserGroupList,omitnil,omitempty" name:"UserGroupList"`

	// 是否删除家目录，只针对cvm集群
	DeleteHomeDir *bool `json:"DeleteHomeDir,omitnil,omitempty" name:"DeleteHomeDir"`
}

func (r *DeleteUserManagerUserListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserManagerUserListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserNameList")
	delete(f, "TkeClusterId")
	delete(f, "DisplayStrategy")
	delete(f, "UserGroupList")
	delete(f, "DeleteHomeDir")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserManagerUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserManagerUserListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserManagerUserListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserManagerUserListResponseParams `json:"Response"`
}

func (r *DeleteUserManagerUserListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserManagerUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DependService struct {
	// 共用组件名
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 共用组件集群
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

// Predefined struct for user
type DeployYarnConfRequestParams struct {
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeployYarnConfRequest struct {
	*tchttp.BaseRequest
	
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeployYarnConfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployYarnConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployYarnConfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployYarnConfResponseParams struct {
	// 启动流程后的流程ID，可以使用[DescribeClusterFlowStatusDetail](https://cloud.tencent.com/document/product/589/107224)接口来获取流程状态
	FlowId *uint64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeployYarnConfResponse struct {
	*tchttp.BaseResponse
	Response *DeployYarnConfResponseParams `json:"Response"`
}

func (r *DeployYarnConfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployYarnConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScaleGroupGlobalConfRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeAutoScaleGroupGlobalConfRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeAutoScaleGroupGlobalConfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScaleGroupGlobalConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoScaleGroupGlobalConfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScaleGroupGlobalConfResponseParams struct {
	// 集群所有伸缩组全局信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupGlobalConfs []*GroupGlobalConfs `json:"GroupGlobalConfs,omitnil,omitempty" name:"GroupGlobalConfs"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoScaleGroupGlobalConfResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoScaleGroupGlobalConfResponseParams `json:"Response"`
}

func (r *DescribeAutoScaleGroupGlobalConfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScaleGroupGlobalConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScaleRecordsRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 记录过滤参数，目前仅能为“StartTime”,“EndTime”和“StrategyName”、ActionStatus、ScaleAction。
	// StartTime和EndTime支持2006-01-02 15:04:05 或者2006/01/02 15:04:05的时间格式
	// ActionStatus：0:INITED,1:SUCCESS, 2:FAILED,3:LIMITED_SUCCESSED,4:IN_PROCESS,5:IN_RETRY
	// ScaleAction：1:扩容  2:缩容
	Filters []*KeyValue `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页参数。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数。最大支持100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 表示是自动(0)还是托管伸缩(1)
	RecordSource *int64 `json:"RecordSource,omitnil,omitempty" name:"RecordSource"`

	// 是否升序，1:升序，0:降序
	Asc *int64 `json:"Asc,omitnil,omitempty" name:"Asc"`
}

type DescribeAutoScaleRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 记录过滤参数，目前仅能为“StartTime”,“EndTime”和“StrategyName”、ActionStatus、ScaleAction。
	// StartTime和EndTime支持2006-01-02 15:04:05 或者2006/01/02 15:04:05的时间格式
	// ActionStatus：0:INITED,1:SUCCESS, 2:FAILED,3:LIMITED_SUCCESSED,4:IN_PROCESS,5:IN_RETRY
	// ScaleAction：1:扩容  2:缩容
	Filters []*KeyValue `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页参数。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数。最大支持100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 表示是自动(0)还是托管伸缩(1)
	RecordSource *int64 `json:"RecordSource,omitnil,omitempty" name:"RecordSource"`

	// 是否升序，1:升序，0:降序
	Asc *int64 `json:"Asc,omitnil,omitempty" name:"Asc"`
}

func (r *DescribeAutoScaleRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScaleRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "RecordSource")
	delete(f, "Asc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoScaleRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScaleRecordsResponseParams struct {
	// 总扩缩容记录数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 记录列表。
	RecordList []*AutoScaleRecord `json:"RecordList,omitnil,omitempty" name:"RecordList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoScaleRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoScaleRecordsResponseParams `json:"Response"`
}

func (r *DescribeAutoScaleRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScaleRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScaleStrategiesRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 伸缩组id
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DescribeAutoScaleStrategiesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 伸缩组id
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *DescribeAutoScaleStrategiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScaleStrategiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoScaleStrategiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScaleStrategiesResponseParams struct {
	// 按负载伸缩规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadAutoScaleStrategies []*LoadAutoScaleStrategy `json:"LoadAutoScaleStrategies,omitnil,omitempty" name:"LoadAutoScaleStrategies"`

	// 按时间伸缩规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeBasedAutoScaleStrategies []*TimeAutoScaleStrategy `json:"TimeBasedAutoScaleStrategies,omitnil,omitempty" name:"TimeBasedAutoScaleStrategies"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoScaleStrategiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoScaleStrategiesResponseParams `json:"Response"`
}

func (r *DescribeAutoScaleStrategiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScaleStrategiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterFlowStatusDetailRequestParams struct {
	// EMR实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 流程相关参数
	FlowParam *FlowParam `json:"FlowParam,omitnil,omitempty" name:"FlowParam"`

	// 是否返回任务额外信息
	// 默认: false
	NeedExtraDetail *bool `json:"NeedExtraDetail,omitnil,omitempty" name:"NeedExtraDetail"`
}

type DescribeClusterFlowStatusDetailRequest struct {
	*tchttp.BaseRequest
	
	// EMR实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 流程相关参数
	FlowParam *FlowParam `json:"FlowParam,omitnil,omitempty" name:"FlowParam"`

	// 是否返回任务额外信息
	// 默认: false
	NeedExtraDetail *bool `json:"NeedExtraDetail,omitnil,omitempty" name:"NeedExtraDetail"`
}

func (r *DescribeClusterFlowStatusDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterFlowStatusDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FlowParam")
	delete(f, "NeedExtraDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterFlowStatusDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterFlowStatusDetailResponseParams struct {
	// 任务步骤详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	StageDetails []*StageInfoDetail `json:"StageDetails,omitnil,omitempty" name:"StageDetails"`

	// 任务参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowDesc []*FlowParamsDesc `json:"FlowDesc,omitnil,omitempty" name:"FlowDesc"`

	// 任务名称
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 总任务流程进度：
	// 例如：0.8
	FlowTotalProgress *float64 `json:"FlowTotalProgress,omitnil,omitempty" name:"FlowTotalProgress"`

	// 定义流程总状态：
	// 0:初始化，
	// 1:运行中，
	// 2:完成，
	// 3:完成（存在跳过步骤），
	// -1:失败，
	// -3:阻塞，
	FlowTotalStatus *int64 `json:"FlowTotalStatus,omitnil,omitempty" name:"FlowTotalStatus"`

	// 流程额外信息
	// NeedExtraDetail为true时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowExtraDetail []*FlowExtraDetail `json:"FlowExtraDetail,omitnil,omitempty" name:"FlowExtraDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterFlowStatusDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterFlowStatusDetailResponseParams `json:"Response"`
}

func (r *DescribeClusterFlowStatusDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterFlowStatusDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterNodesRequestParams struct {
	// 集群实例ID,实例ID形如: emr-xxxxxxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点标识，取值为：
	// <li>all：表示获取全部类型节点，cdb信息除外。</li>
	// <li>master：表示获取master节点信息。</li>
	// <li>core：表示获取core节点信息。</li>
	// <li>task：表示获取task节点信息。</li>
	// <li>common：表示获取common节点信息。</li>
	// <li>router：表示获取router节点信息。</li>
	// <li>db：表示获取正常状态的cdb信息。</li>
	// <li>recyle：表示获取回收站隔离中的节点信息，包括cdb信息。</li>
	// <li>renew：表示获取所有待续费的节点信息，包括cdb信息，自动续费节点不会返回。</li>
	// 注意：现在只支持以上取值，输入其他值会导致错误。
	NodeFlag *string `json:"NodeFlag,omitnil,omitempty" name:"NodeFlag"`

	// 导出全部节点信息csv时是否携带cdb信息
	ExportDb *bool `json:"ExportDb,omitnil,omitempty" name:"ExportDb"`

	// 页编号，默认值为0，表示第一页。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页返回数量，默认值为100，最大值为100。
	// 如果offset和limit都不填，或者都填0，则返回全部数据
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 资源类型:支持all/host/pod，默认为all
	HardwareResourceType *string `json:"HardwareResourceType,omitnil,omitempty" name:"HardwareResourceType"`

	// 支持搜索的字段
	SearchFields []*SearchItem `json:"SearchFields,omitnil,omitempty" name:"SearchFields"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 是否升序，1:升序，0:降序
	Asc *int64 `json:"Asc,omitnil,omitempty" name:"Asc"`
}

type DescribeClusterNodesRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID,实例ID形如: emr-xxxxxxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点标识，取值为：
	// <li>all：表示获取全部类型节点，cdb信息除外。</li>
	// <li>master：表示获取master节点信息。</li>
	// <li>core：表示获取core节点信息。</li>
	// <li>task：表示获取task节点信息。</li>
	// <li>common：表示获取common节点信息。</li>
	// <li>router：表示获取router节点信息。</li>
	// <li>db：表示获取正常状态的cdb信息。</li>
	// <li>recyle：表示获取回收站隔离中的节点信息，包括cdb信息。</li>
	// <li>renew：表示获取所有待续费的节点信息，包括cdb信息，自动续费节点不会返回。</li>
	// 注意：现在只支持以上取值，输入其他值会导致错误。
	NodeFlag *string `json:"NodeFlag,omitnil,omitempty" name:"NodeFlag"`

	// 导出全部节点信息csv时是否携带cdb信息
	ExportDb *bool `json:"ExportDb,omitnil,omitempty" name:"ExportDb"`

	// 页编号，默认值为0，表示第一页。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页返回数量，默认值为100，最大值为100。
	// 如果offset和limit都不填，或者都填0，则返回全部数据
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 资源类型:支持all/host/pod，默认为all
	HardwareResourceType *string `json:"HardwareResourceType,omitnil,omitempty" name:"HardwareResourceType"`

	// 支持搜索的字段
	SearchFields []*SearchItem `json:"SearchFields,omitnil,omitempty" name:"SearchFields"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 是否升序，1:升序，0:降序
	Asc *int64 `json:"Asc,omitnil,omitempty" name:"Asc"`
}

func (r *DescribeClusterNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeFlag")
	delete(f, "ExportDb")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "HardwareResourceType")
	delete(f, "SearchFields")
	delete(f, "OrderField")
	delete(f, "Asc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterNodesResponseParams struct {
	// 查询到的节点总数
	TotalCnt *int64 `json:"TotalCnt,omitnil,omitempty" name:"TotalCnt"`

	// 节点详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeList []*NodeHardwareInfo `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// 用户所有的标签键列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKeys []*string `json:"TagKeys,omitnil,omitempty" name:"TagKeys"`

	// 资源类型列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	HardwareResourceTypeList []*string `json:"HardwareResourceTypeList,omitnil,omitempty" name:"HardwareResourceTypeList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterNodesResponseParams `json:"Response"`
}

func (r *DescribeClusterNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCvmQuotaRequestParams struct {
	// EMR集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 区ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type DescribeCvmQuotaRequest struct {
	*tchttp.BaseRequest
	
	// EMR集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 区ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *DescribeCvmQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCvmQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCvmQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCvmQuotaResponseParams struct {
	// 后付费配额列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostPaidQuotaSet []*QuotaEntity `json:"PostPaidQuotaSet,omitnil,omitempty" name:"PostPaidQuotaSet"`

	// 竞价实例配额列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpotPaidQuotaSet []*QuotaEntity `json:"SpotPaidQuotaSet,omitnil,omitempty" name:"SpotPaidQuotaSet"`

	// eks配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	EksQuotaSet []*PodSaleSpec `json:"EksQuotaSet,omitnil,omitempty" name:"EksQuotaSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCvmQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCvmQuotaResponseParams `json:"Response"`
}

func (r *DescribeCvmQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCvmQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDAGInfoRequestParams struct {
	// 集群ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// DAG类型，目前只支持STARROCKS
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 查询ID列表,最大长度为1
	IDList []*string `json:"IDList,omitnil,omitempty" name:"IDList"`
}

type DescribeDAGInfoRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// DAG类型，目前只支持STARROCKS
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 查询ID列表,最大长度为1
	IDList []*string `json:"IDList,omitnil,omitempty" name:"IDList"`
}

func (r *DescribeDAGInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDAGInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "Type")
	delete(f, "IDList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDAGInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDAGInfoResponseParams struct {
	// 总数，分页查询时使用
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Starrocks 查询信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DAGInfoList []*DAGInfo `json:"DAGInfoList,omitnil,omitempty" name:"DAGInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDAGInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDAGInfoResponseParams `json:"Response"`
}

func (r *DescribeDAGInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDAGInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEmrApplicationStaticsRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 起始时间，时间戳（秒）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，时间戳（秒）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 过滤的队列名
	Queues []*string `json:"Queues,omitnil,omitempty" name:"Queues"`

	// 过滤的用户名
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// 过滤的作业类型
	ApplicationTypes []*string `json:"ApplicationTypes,omitnil,omitempty" name:"ApplicationTypes"`

	// 分组字段，可选：queue, user, applicationType
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// 排序字段，可选：sumMemorySeconds, sumVCoreSeconds, sumHDFSBytesWritten, sumHDFSBytesRead
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 是否顺序排序，0-逆序，1-正序
	IsAsc *int64 `json:"IsAsc,omitnil,omitempty" name:"IsAsc"`

	// 页号
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页容量，范围为[10,100]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeEmrApplicationStaticsRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 起始时间，时间戳（秒）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，时间戳（秒）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 过滤的队列名
	Queues []*string `json:"Queues,omitnil,omitempty" name:"Queues"`

	// 过滤的用户名
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// 过滤的作业类型
	ApplicationTypes []*string `json:"ApplicationTypes,omitnil,omitempty" name:"ApplicationTypes"`

	// 分组字段，可选：queue, user, applicationType
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// 排序字段，可选：sumMemorySeconds, sumVCoreSeconds, sumHDFSBytesWritten, sumHDFSBytesRead
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 是否顺序排序，0-逆序，1-正序
	IsAsc *int64 `json:"IsAsc,omitnil,omitempty" name:"IsAsc"`

	// 页号
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页容量，范围为[10,100]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeEmrApplicationStaticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEmrApplicationStaticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Queues")
	delete(f, "Users")
	delete(f, "ApplicationTypes")
	delete(f, "GroupBy")
	delete(f, "OrderBy")
	delete(f, "IsAsc")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEmrApplicationStaticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEmrApplicationStaticsResponseParams struct {
	// 作业统计信息
	Statics []*ApplicationStatics `json:"Statics,omitnil,omitempty" name:"Statics"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 可选择的队列名
	Queues []*string `json:"Queues,omitnil,omitempty" name:"Queues"`

	// 可选择的用户名
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// 可选择的作业类型
	ApplicationTypes []*string `json:"ApplicationTypes,omitnil,omitempty" name:"ApplicationTypes"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEmrApplicationStaticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEmrApplicationStaticsResponseParams `json:"Response"`
}

func (r *DescribeEmrApplicationStaticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEmrApplicationStaticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEmrOverviewMetricsRequestParams struct {
	// 结束时间
	End *int64 `json:"End,omitnil,omitempty" name:"End"`

	// 指标名，NODE.CPU：节点平均CPU利用率和总核数；NODE.CPU.SLHBASE：Serverless实例平均CPU利用率和总核数；HDFS.NN.CAPACITY：存储使用率和总量
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 粒度 30s-max 1m-max 1h-max等
	Downsample *string `json:"Downsample,omitnil,omitempty" name:"Downsample"`

	// 起始时间，画饼状图时不传
	Start *int64 `json:"Start,omitnil,omitempty" name:"Start"`

	// 聚合方法，扩展用，这里目前不用传
	Aggregator *string `json:"Aggregator,omitnil,omitempty" name:"Aggregator"`

	// 指标要查询的具体type 如："{"type":"CapacityTotal|CapacityRemaining"}"
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeEmrOverviewMetricsRequest struct {
	*tchttp.BaseRequest
	
	// 结束时间
	End *int64 `json:"End,omitnil,omitempty" name:"End"`

	// 指标名，NODE.CPU：节点平均CPU利用率和总核数；NODE.CPU.SLHBASE：Serverless实例平均CPU利用率和总核数；HDFS.NN.CAPACITY：存储使用率和总量
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 粒度 30s-max 1m-max 1h-max等
	Downsample *string `json:"Downsample,omitnil,omitempty" name:"Downsample"`

	// 起始时间，画饼状图时不传
	Start *int64 `json:"Start,omitnil,omitempty" name:"Start"`

	// 聚合方法，扩展用，这里目前不用传
	Aggregator *string `json:"Aggregator,omitnil,omitempty" name:"Aggregator"`

	// 指标要查询的具体type 如："{"type":"CapacityTotal|CapacityRemaining"}"
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeEmrOverviewMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEmrOverviewMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "End")
	delete(f, "Metric")
	delete(f, "InstanceId")
	delete(f, "Downsample")
	delete(f, "Start")
	delete(f, "Aggregator")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEmrOverviewMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEmrOverviewMetricsResponseParams struct {
	// 指标数据明细
	Result []*OverviewMetricData `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEmrOverviewMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEmrOverviewMetricsResponseParams `json:"Response"`
}

func (r *DescribeEmrOverviewMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEmrOverviewMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalConfigRequestParams struct {
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeGlobalConfigRequest struct {
	*tchttp.BaseRequest
	
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeGlobalConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGlobalConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalConfigResponseParams struct {
	// 是否开启了资源调度功能
	EnableResourceSchedule *bool `json:"EnableResourceSchedule,omitnil,omitempty" name:"EnableResourceSchedule"`

	// 当前生效的资源调度器
	ActiveScheduler *string `json:"ActiveScheduler,omitnil,omitempty" name:"ActiveScheduler"`

	// 公平调度器的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CapacityGlobalConfig *CapacityGlobalConfig `json:"CapacityGlobalConfig,omitnil,omitempty" name:"CapacityGlobalConfig"`

	// 容量调度器的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FairGlobalConfig *FairGlobalConfig `json:"FairGlobalConfig,omitnil,omitempty" name:"FairGlobalConfig"`

	// 最新的资源调度器
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGlobalConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGlobalConfigResponseParams `json:"Response"`
}

func (r *DescribeGlobalConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHBaseTableOverviewRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页查询编号偏移量，从0开始	
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询时的分页大小，最小1，最大100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 表名称，模糊匹配
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 排序的字段，有默认值
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 默认为降序，asc代表升序，desc代表降序
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

type DescribeHBaseTableOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页查询编号偏移量，从0开始	
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询时的分页大小，最小1，最大100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 表名称，模糊匹配
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 排序的字段，有默认值
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 默认为降序，asc代表升序，desc代表降序
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

func (r *DescribeHBaseTableOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHBaseTableOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Table")
	delete(f, "OrderField")
	delete(f, "OrderType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHBaseTableOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHBaseTableOverviewResponseParams struct {
	// 概览数据数组
	TableMonitorList []*OverviewRow `json:"TableMonitorList,omitnil,omitempty" name:"TableMonitorList"`

	// 概览数据数组长度
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 表schema信息
	SchemaList []*TableSchemaItem `json:"SchemaList,omitnil,omitempty" name:"SchemaList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHBaseTableOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHBaseTableOverviewResponseParams `json:"Response"`
}

func (r *DescribeHBaseTableOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHBaseTableOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHDFSStorageInfoRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 获取查询信息开始时间 (s)
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 获取查询信息结束时间 (s)
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeHDFSStorageInfoRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 获取查询信息开始时间 (s)
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 获取查询信息结束时间 (s)
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeHDFSStorageInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHDFSStorageInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHDFSStorageInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHDFSStorageInfoResponseParams struct {
	// 采样时间
	SampleTime *int64 `json:"SampleTime,omitnil,omitempty" name:"SampleTime"`

	// hdfs存储详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSummaryDistribution []*StorageSummaryDistribution `json:"StorageSummaryDistribution,omitnil,omitempty" name:"StorageSummaryDistribution"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHDFSStorageInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHDFSStorageInfoResponseParams `json:"Response"`
}

func (r *DescribeHDFSStorageInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHDFSStorageInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHiveQueriesRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 起始时间秒
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间秒，EndTime-StartTime不得超过1天秒数86400
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页起始偏移，从0开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，合法范围[1,100]
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 执行状态,ERROR等
	State []*string `json:"State,omitnil,omitempty" name:"State"`

	// 结束时间大于的时间点
	EndTimeGte *uint64 `json:"EndTimeGte,omitnil,omitempty" name:"EndTimeGte"`

	// 结束时间小于时间点
	EndTimeLte *uint64 `json:"EndTimeLte,omitnil,omitempty" name:"EndTimeLte"`
}

type DescribeHiveQueriesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 起始时间秒
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间秒，EndTime-StartTime不得超过1天秒数86400
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页起始偏移，从0开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，合法范围[1,100]
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 执行状态,ERROR等
	State []*string `json:"State,omitnil,omitempty" name:"State"`

	// 结束时间大于的时间点
	EndTimeGte *uint64 `json:"EndTimeGte,omitnil,omitempty" name:"EndTimeGte"`

	// 结束时间小于时间点
	EndTimeLte *uint64 `json:"EndTimeLte,omitnil,omitempty" name:"EndTimeLte"`
}

func (r *DescribeHiveQueriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHiveQueriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "State")
	delete(f, "EndTimeGte")
	delete(f, "EndTimeLte")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHiveQueriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHiveQueriesResponseParams struct {
	// 总条数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 结果列表
	Results []*HiveQuery `json:"Results,omitnil,omitempty" name:"Results"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHiveQueriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHiveQueriesResponseParams `json:"Response"`
}

func (r *DescribeHiveQueriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHiveQueriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImpalaQueriesRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 起始时间秒
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间秒，EndTime-StartTime不得超过1天秒数86400
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页起始偏移，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，合法范围[1,100]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 执行状态，CREATED、INITIALIZED、COMPILED、RUNNING、FINISHED、EXCEPTION
	State []*string `json:"State,omitnil,omitempty" name:"State"`

	// 结束时间大于的时间点
	EndTimeGte *uint64 `json:"EndTimeGte,omitnil,omitempty" name:"EndTimeGte"`

	// 结束时间小于的时间点
	EndTimeLte *uint64 `json:"EndTimeLte,omitnil,omitempty" name:"EndTimeLte"`
}

type DescribeImpalaQueriesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 起始时间秒
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间秒，EndTime-StartTime不得超过1天秒数86400
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页起始偏移，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，合法范围[1,100]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 执行状态，CREATED、INITIALIZED、COMPILED、RUNNING、FINISHED、EXCEPTION
	State []*string `json:"State,omitnil,omitempty" name:"State"`

	// 结束时间大于的时间点
	EndTimeGte *uint64 `json:"EndTimeGte,omitnil,omitempty" name:"EndTimeGte"`

	// 结束时间小于的时间点
	EndTimeLte *uint64 `json:"EndTimeLte,omitnil,omitempty" name:"EndTimeLte"`
}

func (r *DescribeImpalaQueriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImpalaQueriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "State")
	delete(f, "EndTimeGte")
	delete(f, "EndTimeLte")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImpalaQueriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImpalaQueriesResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 结果列表
	Results []*ImpalaQuery `json:"Results,omitnil,omitempty" name:"Results"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImpalaQueriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImpalaQueriesResponseParams `json:"Response"`
}

func (r *DescribeImpalaQueriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImpalaQueriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInsightListRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 获取的洞察结果开始时间，此时间针对对App或者Hive查询的开始时间的过滤
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 获取的洞察结果结束时间，此时间针对对App或者Hive查询的开始时间的过滤
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页查询时的分页大小，最小1，最大100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页查询时的页号，从1开始
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 查询类型,支持HIVE,SPARK,DLC_SPARK,SPARK_SQL,SCHEDULE,MAPREDUCE,TRINO等类型,默认查询全部
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeInsightListRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 获取的洞察结果开始时间，此时间针对对App或者Hive查询的开始时间的过滤
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 获取的洞察结果结束时间，此时间针对对App或者Hive查询的开始时间的过滤
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页查询时的分页大小，最小1，最大100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页查询时的页号，从1开始
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 查询类型,支持HIVE,SPARK,DLC_SPARK,SPARK_SQL,SCHEDULE,MAPREDUCE,TRINO等类型,默认查询全部
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeInsightListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInsightListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "Page")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInsightListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInsightListResponseParams struct {
	// 总数，分页查询时使用
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 洞察结果数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultList []*InsightResult `json:"ResultList,omitnil,omitempty" name:"ResultList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInsightListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInsightListResponseParams `json:"Response"`
}

func (r *DescribeInsightListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInsightListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceRenewNodesRequestParams struct {
	// 集群实例ID,实例ID形如: emr-xxxxxxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceRenewNodesRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID,实例ID形如: emr-xxxxxxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceRenewNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceRenewNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceRenewNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceRenewNodesResponseParams struct {
	// 查询到的节点总数
	TotalCnt *int64 `json:"TotalCnt,omitnil,omitempty" name:"TotalCnt"`

	// 节点详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeList []*RenewInstancesInfo `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// 用户所有的标签键列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaInfo []*string `json:"MetaInfo,omitnil,omitempty" name:"MetaInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceRenewNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceRenewNodesResponseParams `json:"Response"`
}

func (r *DescribeInstanceRenewNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceRenewNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesListRequestParams struct {
	// 集群筛选策略。取值范围：<li>clusterList：表示查询除了已销毁集群之外的集群列表。</li><li>monitorManage：表示查询除了已销毁、创建中以及创建失败的集群之外的集群列表。</li><li>cloudHardwareManage/componentManage：目前这两个取值为预留取值，暂时和monitorManage表示同样的含义。</li>
	DisplayStrategy *string `json:"DisplayStrategy,omitnil,omitempty" name:"DisplayStrategy"`

	// 页编号，默认值为0，表示第一页。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页返回数量，默认值为100，最大值为100。
	// 如果limit和offset都为0，则查询全部记录；
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段。取值范围：<li>clusterId：表示按照实例ID排序。</li><li>addTime：表示按照实例创建时间排序。</li><li>status：表示按照实例的状态码排序。</li>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 按照OrderField升序或者降序进行排序。取值范围：<li>0：表示升序。</li><li>1：表示降序。</li>默认值为0。
	Asc *int64 `json:"Asc,omitnil,omitempty" name:"Asc"`

	// 自定义查询过滤器。示例：<li>根据ClusterId过滤实例：[{"Name":"ClusterId","Values":["emr-xxxxxxxx"]}]</li><li>根据clusterName过滤实例：[{"Name": "ClusterName","Values": ["cluster_name"]}]</li><li>根据ClusterStatus过滤实例：[{"Name": "ClusterStatus","Values": ["2"]}]</li>
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeInstancesListRequest struct {
	*tchttp.BaseRequest
	
	// 集群筛选策略。取值范围：<li>clusterList：表示查询除了已销毁集群之外的集群列表。</li><li>monitorManage：表示查询除了已销毁、创建中以及创建失败的集群之外的集群列表。</li><li>cloudHardwareManage/componentManage：目前这两个取值为预留取值，暂时和monitorManage表示同样的含义。</li>
	DisplayStrategy *string `json:"DisplayStrategy,omitnil,omitempty" name:"DisplayStrategy"`

	// 页编号，默认值为0，表示第一页。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页返回数量，默认值为100，最大值为100。
	// 如果limit和offset都为0，则查询全部记录；
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段。取值范围：<li>clusterId：表示按照实例ID排序。</li><li>addTime：表示按照实例创建时间排序。</li><li>status：表示按照实例的状态码排序。</li>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 按照OrderField升序或者降序进行排序。取值范围：<li>0：表示升序。</li><li>1：表示降序。</li>默认值为0。
	Asc *int64 `json:"Asc,omitnil,omitempty" name:"Asc"`

	// 自定义查询过滤器。示例：<li>根据ClusterId过滤实例：[{"Name":"ClusterId","Values":["emr-xxxxxxxx"]}]</li><li>根据clusterName过滤实例：[{"Name": "ClusterName","Values": ["cluster_name"]}]</li><li>根据ClusterStatus过滤实例：[{"Name": "ClusterStatus","Values": ["2"]}]</li>
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeInstancesListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisplayStrategy")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Asc")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesListResponseParams struct {
	// 符合条件的实例总数。
	TotalCnt *int64 `json:"TotalCnt,omitnil,omitempty" name:"TotalCnt"`

	// 集群实例列表
	InstancesList []*EmrListInstance `json:"InstancesList,omitnil,omitempty" name:"InstancesList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesListResponseParams `json:"Response"`
}

func (r *DescribeInstancesListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 集群筛选策略。取值范围：
	// <li>clusterList：表示查询除了已销毁集群之外的集群列表。</li>
	// <li>monitorManage：表示查询除了已销毁、创建中以及创建失败的集群之外的集群列表。</li>
	// <li>cloudHardwareManage/componentManage：目前这两个取值为预留取值，暂时和monitorManage表示同样的含义。</li>
	DisplayStrategy *string `json:"DisplayStrategy,omitnil,omitempty" name:"DisplayStrategy"`

	// 按照一个或者多个实例ID查询。实例ID形如: emr-xxxxxxxx 。(此参数的具体格式可参考API[简介](https://cloud.tencent.com/document/api/213/15688)的 Ids.N 一节)。如果不填写实例ID，返回该APPID下所有实例列表。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 页编号，默认值为0，表示第一页。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页返回数量，默认值为100，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 建议必填-1，表示拉取所有项目下的集群。
	// 不填默认值为0，表示拉取默认项目下的集群。
	// 实例所属项目ID。该参数可以通过调用 [DescribeProjects](https://cloud.tencent.com/document/product/651/78725) 的返回值中的 projectId 字段来获取。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 排序字段。取值范围：
	// <li>clusterId：表示按照实例ID排序。</li>
	// <li>addTime：表示按照实例创建时间排序。</li>
	// <li>status：表示按照实例的状态码排序。</li>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 按照OrderField升序或者降序进行排序。取值范围：
	// <li>0：表示降序。</li>
	// <li>1：表示升序。</li>默认值为0。
	Asc *int64 `json:"Asc,omitnil,omitempty" name:"Asc"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群筛选策略。取值范围：
	// <li>clusterList：表示查询除了已销毁集群之外的集群列表。</li>
	// <li>monitorManage：表示查询除了已销毁、创建中以及创建失败的集群之外的集群列表。</li>
	// <li>cloudHardwareManage/componentManage：目前这两个取值为预留取值，暂时和monitorManage表示同样的含义。</li>
	DisplayStrategy *string `json:"DisplayStrategy,omitnil,omitempty" name:"DisplayStrategy"`

	// 按照一个或者多个实例ID查询。实例ID形如: emr-xxxxxxxx 。(此参数的具体格式可参考API[简介](https://cloud.tencent.com/document/api/213/15688)的 Ids.N 一节)。如果不填写实例ID，返回该APPID下所有实例列表。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 页编号，默认值为0，表示第一页。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页返回数量，默认值为100，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 建议必填-1，表示拉取所有项目下的集群。
	// 不填默认值为0，表示拉取默认项目下的集群。
	// 实例所属项目ID。该参数可以通过调用 [DescribeProjects](https://cloud.tencent.com/document/product/651/78725) 的返回值中的 projectId 字段来获取。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 排序字段。取值范围：
	// <li>clusterId：表示按照实例ID排序。</li>
	// <li>addTime：表示按照实例创建时间排序。</li>
	// <li>status：表示按照实例的状态码排序。</li>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 按照OrderField升序或者降序进行排序。取值范围：
	// <li>0：表示降序。</li>
	// <li>1：表示升序。</li>默认值为0。
	Asc *int64 `json:"Asc,omitnil,omitempty" name:"Asc"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisplayStrategy")
	delete(f, "InstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProjectId")
	delete(f, "OrderField")
	delete(f, "Asc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 符合条件的实例总数。
	TotalCnt *int64 `json:"TotalCnt,omitnil,omitempty" name:"TotalCnt"`

	// EMR实例详细信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterList []*ClusterInstancesInfo `json:"ClusterList,omitnil,omitempty" name:"ClusterList"`

	// 实例关联的标签键列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKeys []*string `json:"TagKeys,omitnil,omitempty" name:"TagKeys"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesResponseParams `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobFlowRequestParams struct {
	// 流程任务Id，RunJobFlow接口返回的值。
	JobFlowId *int64 `json:"JobFlowId,omitnil,omitempty" name:"JobFlowId"`
}

type DescribeJobFlowRequest struct {
	*tchttp.BaseRequest
	
	// 流程任务Id，RunJobFlow接口返回的值。
	JobFlowId *int64 `json:"JobFlowId,omitnil,omitempty" name:"JobFlowId"`
}

func (r *DescribeJobFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobFlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobFlowResponseParams struct {
	// 流程任务状态，可以为以下值：
	// JobFlowInit，流程任务初始化。
	// JobFlowResourceApplied，资源申请中，通常为JobFlow需要新建集群时的状态。
	// JobFlowResourceReady，执行流程任务的资源就绪。
	// JobFlowStepsRunning，流程任务步骤已提交。
	// JobFlowStepsComplete，流程任务步骤已完成。
	// JobFlowTerminating，流程任务所需资源销毁中。
	// JobFlowFinish，流程任务已完成。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 流程任务步骤结果。
	Details []*JobResult `json:"Details,omitnil,omitempty" name:"Details"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeJobFlowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobFlowResponseParams `json:"Response"`
}

func (r *DescribeJobFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKyuubiQueryInfoRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 获取查询信息开始时间 (s)
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 获取查询信息结束时间 (s)
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页查询时的分页大小，最小1，最大100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页查询时的页号，从1开始
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type DescribeKyuubiQueryInfoRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 获取查询信息开始时间 (s)
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 获取查询信息结束时间 (s)
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页查询时的分页大小，最小1，最大100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页查询时的页号，从1开始
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

func (r *DescribeKyuubiQueryInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKyuubiQueryInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "Page")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKyuubiQueryInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKyuubiQueryInfoResponseParams struct {
	// 总数，分页查询时使用
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Kyuubi查询信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	KyuubiQueryInfoList []*KyuubiQueryInfo `json:"KyuubiQueryInfoList,omitnil,omitempty" name:"KyuubiQueryInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKyuubiQueryInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKyuubiQueryInfoResponseParams `json:"Response"`
}

func (r *DescribeKyuubiQueryInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKyuubiQueryInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodeDataDisksRequestParams struct {
	// EMR集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点CVM实例Id列表
	CvmInstanceIds []*string `json:"CvmInstanceIds,omitnil,omitempty" name:"CvmInstanceIds"`
}

type DescribeNodeDataDisksRequest struct {
	*tchttp.BaseRequest
	
	// EMR集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点CVM实例Id列表
	CvmInstanceIds []*string `json:"CvmInstanceIds,omitnil,omitempty" name:"CvmInstanceIds"`
}

func (r *DescribeNodeDataDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeDataDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CvmInstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNodeDataDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodeDataDisksResponseParams struct {
	// 总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 云盘列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CBSList []*CBSInstance `json:"CBSList,omitnil,omitempty" name:"CBSList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNodeDataDisksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNodeDataDisksResponseParams `json:"Response"`
}

func (r *DescribeNodeDataDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeDataDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodeResourceConfigFastRequestParams struct {
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点类型 CORE TASK ROUTER ALL
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 计费类型
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 可用区ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 类型为ComputeResource和EMR以及默认，默认为EMR
	ResourceBaseType *string `json:"ResourceBaseType,omitnil,omitempty" name:"ResourceBaseType"`

	// 计算资源id
	ComputeResourceId *string `json:"ComputeResourceId,omitnil,omitempty" name:"ComputeResourceId"`

	// 硬件类型
	HardwareResourceType *string `json:"HardwareResourceType,omitnil,omitempty" name:"HardwareResourceType"`
}

type DescribeNodeResourceConfigFastRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点类型 CORE TASK ROUTER ALL
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 计费类型
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 可用区ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 类型为ComputeResource和EMR以及默认，默认为EMR
	ResourceBaseType *string `json:"ResourceBaseType,omitnil,omitempty" name:"ResourceBaseType"`

	// 计算资源id
	ComputeResourceId *string `json:"ComputeResourceId,omitnil,omitempty" name:"ComputeResourceId"`

	// 硬件类型
	HardwareResourceType *string `json:"HardwareResourceType,omitnil,omitempty" name:"HardwareResourceType"`
}

func (r *DescribeNodeResourceConfigFastRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeResourceConfigFastRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceType")
	delete(f, "PayMode")
	delete(f, "ZoneId")
	delete(f, "ResourceBaseType")
	delete(f, "ComputeResourceId")
	delete(f, "HardwareResourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNodeResourceConfigFastRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodeResourceConfigFastResponseParams struct {
	// DescribeResourceConfig接口返回值
	Data []*DescribeResourceConfig `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNodeResourceConfigFastResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNodeResourceConfigFastResponseParams `json:"Response"`
}

func (r *DescribeNodeResourceConfigFastResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeResourceConfigFastResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceConfig struct {
	// 规格管理类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 规格管理数据
	ResourceData []*NodeResource `json:"ResourceData,omitnil,omitempty" name:"ResourceData"`
}

// Predefined struct for user
type DescribeResourceScheduleDiffDetailRequestParams struct {
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询的变更明细对应的调度器，可选值为fair、capacity。如果不传或者传空会使用最新的调度器
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`
}

type DescribeResourceScheduleDiffDetailRequest struct {
	*tchttp.BaseRequest
	
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询的变更明细对应的调度器，可选值为fair、capacity。如果不传或者传空会使用最新的调度器
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`
}

func (r *DescribeResourceScheduleDiffDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceScheduleDiffDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Scheduler")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceScheduleDiffDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceScheduleDiffDetailResponseParams struct {
	// 变化项的明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details []*DiffDetail `json:"Details,omitnil,omitempty" name:"Details"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourceScheduleDiffDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceScheduleDiffDetailResponseParams `json:"Response"`
}

func (r *DescribeResourceScheduleDiffDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceScheduleDiffDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceScheduleRequestParams struct {
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeResourceScheduleRequest struct {
	*tchttp.BaseRequest
	
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeResourceScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceScheduleResponseParams struct {
	// 资源调度功能是否开启
	OpenSwitch *bool `json:"OpenSwitch,omitnil,omitempty" name:"OpenSwitch"`

	// 正在使用的资源调度器
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// 公平调度器的信息
	FSInfo *string `json:"FSInfo,omitnil,omitempty" name:"FSInfo"`

	// 容量调度器的信息
	CSInfo *string `json:"CSInfo,omitnil,omitempty" name:"CSInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourceScheduleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceScheduleResponseParams `json:"Response"`
}

func (r *DescribeResourceScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSLInstanceListRequestParams struct {
	// 实例筛选策略。取值范围：<li>clusterList：表示查询除了已销毁实例之外的实例列表。</li><li>monitorManage：表示查询除了已销毁、创建中以及创建失败的实例之外的实例列表。</li>
	DisplayStrategy *string `json:"DisplayStrategy,omitnil,omitempty" name:"DisplayStrategy"`

	// 页编号，默认值为0，表示第一页。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页返回数量，默认值为10，最大值为100。	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段。取值范围：<li>clusterId：表示按照实例ID排序。</li><li>addTime：表示按照实例创建时间排序。</li><li>status：表示按照实例的状态码排序。</li>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 按照OrderField升序或者降序进行排序。取值范围：<li>0：表示升序。</li><li>1：表示降序。</li>默认值为0。
	Asc *int64 `json:"Asc,omitnil,omitempty" name:"Asc"`

	// 自定义查询过滤器。示例：<li>根据ClusterId过滤实例：[{"Name":"ClusterId","Values":["emr-xxxxxxxx"]}]</li><li>根据clusterName过滤实例：[{"Name": "ClusterName","Values": ["cluster_name"]}]</li><li>根据ClusterStatus过滤实例：[{"Name": "ClusterStatus","Values": ["2"]}]</li>
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeSLInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 实例筛选策略。取值范围：<li>clusterList：表示查询除了已销毁实例之外的实例列表。</li><li>monitorManage：表示查询除了已销毁、创建中以及创建失败的实例之外的实例列表。</li>
	DisplayStrategy *string `json:"DisplayStrategy,omitnil,omitempty" name:"DisplayStrategy"`

	// 页编号，默认值为0，表示第一页。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页返回数量，默认值为10，最大值为100。	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段。取值范围：<li>clusterId：表示按照实例ID排序。</li><li>addTime：表示按照实例创建时间排序。</li><li>status：表示按照实例的状态码排序。</li>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 按照OrderField升序或者降序进行排序。取值范围：<li>0：表示升序。</li><li>1：表示降序。</li>默认值为0。
	Asc *int64 `json:"Asc,omitnil,omitempty" name:"Asc"`

	// 自定义查询过滤器。示例：<li>根据ClusterId过滤实例：[{"Name":"ClusterId","Values":["emr-xxxxxxxx"]}]</li><li>根据clusterName过滤实例：[{"Name": "ClusterName","Values": ["cluster_name"]}]</li><li>根据ClusterStatus过滤实例：[{"Name": "ClusterStatus","Values": ["2"]}]</li>
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeSLInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSLInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisplayStrategy")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Asc")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSLInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSLInstanceListResponseParams struct {
	// 符合条件的实例总数。	
	TotalCnt *int64 `json:"TotalCnt,omitnil,omitempty" name:"TotalCnt"`

	// 实例信息列表，如果进行了分页，只显示当前分页的示例信息列表。
	InstancesList []*SLInstanceInfo `json:"InstancesList,omitnil,omitempty" name:"InstancesList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSLInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSLInstanceListResponseParams `json:"Response"`
}

func (r *DescribeSLInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSLInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSLInstanceRequestParams struct {
	// 实例唯一标识符（字符串表示）
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeSLInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例唯一标识符（字符串表示）
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeSLInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSLInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSLInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSLInstanceResponseParams struct {
	// 实例字符串标识。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例计费模式。0表示后付费，即按量计费，1表示预付费，即包年包月。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 实例存储类型。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 实例单节点磁盘容量，单位GB。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 实例节点规格。
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 实例可用区详细配置，包含可用区名称，VPC信息、节点数量。
	ZoneSettings []*ZoneSetting `json:"ZoneSettings,omitnil,omitempty" name:"ZoneSettings"`

	// 实例绑定的标签列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 实例数字标识。
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例区域ID。
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 实例主可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例过期时间，后付费返回0000-00-00 00:00:00
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 实例隔离时间，未隔离返回0000-00-00 00:00:00。
	IsolateTime *string `json:"IsolateTime,omitnil,omitempty" name:"IsolateTime"`

	// 实例创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例状态码，-2:  "TERMINATED", 2:   "RUNNING", 14:  "TERMINATING", 19:  "ISOLATING", 22:  "ADJUSTING", 201: "ISOLATED"。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 自动续费标记， 0：表示通知即将过期，但不自动续费 1：表示通知即将过期，而且自动续费 2：表示不通知即将过期，也不自动续费，若业务无续费概念为0
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 实例节点总数。
	NodeNum *int64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSLInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSLInstanceResponseParams `json:"Response"`
}

func (r *DescribeSLInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSLInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceNodeInfosRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 页码
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索字段
	SearchText *string `json:"SearchText,omitnil,omitempty" name:"SearchText"`

	// '配置状态，-2：配置失败，-1:配置过期，1：已同步', -99 '全部'
	ConfStatus *int64 `json:"ConfStatus,omitnil,omitempty" name:"ConfStatus"`

	// 过滤条件：维护状态
	// 0代表所有状态
	// 1代表正常模式
	// 2代表维护模式
	MaintainStateId *int64 `json:"MaintainStateId,omitnil,omitempty" name:"MaintainStateId"`

	// 过滤条件：操作状态
	// 0代表所有状态
	// 1代表已启动
	// 2代表已停止
	OperatorStateId *int64 `json:"OperatorStateId,omitnil,omitempty" name:"OperatorStateId"`

	// 过滤条件：健康状态
	// "0"代表不可用
	// "1"代表良好
	// "-2"代表未知
	// "-99"代表所有
	// "-3"代表存在隐患
	// "-4"代表未探测
	HealthStateId *string `json:"HealthStateId,omitnil,omitempty" name:"HealthStateId"`

	// 服务组件名称，都是大写例如YARN
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 节点名称master,core,task,common,router
	NodeTypeName *string `json:"NodeTypeName,omitnil,omitempty" name:"NodeTypeName"`

	// 过滤条件：dn是否处于维护状态
	// 0代表所有状态
	// 1代表处于维护状态
	DataNodeMaintenanceId *int64 `json:"DataNodeMaintenanceId,omitnil,omitempty" name:"DataNodeMaintenanceId"`

	// 支持搜索的字段，目前支持 SearchType	：ipv4
	SearchFields []*SearchItem `json:"SearchFields,omitnil,omitempty" name:"SearchFields"`
}

type DescribeServiceNodeInfosRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 页码
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索字段
	SearchText *string `json:"SearchText,omitnil,omitempty" name:"SearchText"`

	// '配置状态，-2：配置失败，-1:配置过期，1：已同步', -99 '全部'
	ConfStatus *int64 `json:"ConfStatus,omitnil,omitempty" name:"ConfStatus"`

	// 过滤条件：维护状态
	// 0代表所有状态
	// 1代表正常模式
	// 2代表维护模式
	MaintainStateId *int64 `json:"MaintainStateId,omitnil,omitempty" name:"MaintainStateId"`

	// 过滤条件：操作状态
	// 0代表所有状态
	// 1代表已启动
	// 2代表已停止
	OperatorStateId *int64 `json:"OperatorStateId,omitnil,omitempty" name:"OperatorStateId"`

	// 过滤条件：健康状态
	// "0"代表不可用
	// "1"代表良好
	// "-2"代表未知
	// "-99"代表所有
	// "-3"代表存在隐患
	// "-4"代表未探测
	HealthStateId *string `json:"HealthStateId,omitnil,omitempty" name:"HealthStateId"`

	// 服务组件名称，都是大写例如YARN
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 节点名称master,core,task,common,router
	NodeTypeName *string `json:"NodeTypeName,omitnil,omitempty" name:"NodeTypeName"`

	// 过滤条件：dn是否处于维护状态
	// 0代表所有状态
	// 1代表处于维护状态
	DataNodeMaintenanceId *int64 `json:"DataNodeMaintenanceId,omitnil,omitempty" name:"DataNodeMaintenanceId"`

	// 支持搜索的字段，目前支持 SearchType	：ipv4
	SearchFields []*SearchItem `json:"SearchFields,omitnil,omitempty" name:"SearchFields"`
}

func (r *DescribeServiceNodeInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceNodeInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchText")
	delete(f, "ConfStatus")
	delete(f, "MaintainStateId")
	delete(f, "OperatorStateId")
	delete(f, "HealthStateId")
	delete(f, "ServiceName")
	delete(f, "NodeTypeName")
	delete(f, "DataNodeMaintenanceId")
	delete(f, "SearchFields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceNodeInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceNodeInfosResponseParams struct {
	// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCnt *int64 `json:"TotalCnt,omitnil,omitempty" name:"TotalCnt"`

	// 进程信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceNodeList []*ServiceNodeDetailInfo `json:"ServiceNodeList,omitnil,omitempty" name:"ServiceNodeList"`

	// 集群所有节点的别名序列化
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliasInfo *string `json:"AliasInfo,omitnil,omitempty" name:"AliasInfo"`

	// 支持的FlagNode列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportNodeFlagFilterList []*string `json:"SupportNodeFlagFilterList,omitnil,omitempty" name:"SupportNodeFlagFilterList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServiceNodeInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceNodeInfosResponseParams `json:"Response"`
}

func (r *DescribeServiceNodeInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceNodeInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkQueriesRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页起始偏移，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，合法范围[1,100]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 执行状态:RUNNING,COMPLETED,FAILED
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`
}

type DescribeSparkQueriesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页起始偏移，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，合法范围[1,100]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 执行状态:RUNNING,COMPLETED,FAILED
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *DescribeSparkQueriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkQueriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSparkQueriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkQueriesResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*SparkQuery `json:"Results,omitnil,omitempty" name:"Results"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSparkQueriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSparkQueriesResponseParams `json:"Response"`
}

func (r *DescribeSparkQueriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkQueriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStarRocksQueryInfoRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 获取查询信息开始时间 (s)
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 获取查询信息结束时间 (s)
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页查询时的分页大小，最小1，最大100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页查询时的页号，从1开始
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type DescribeStarRocksQueryInfoRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 获取查询信息开始时间 (s)
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 获取查询信息结束时间 (s)
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页查询时的分页大小，最小1，最大100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页查询时的页号，从1开始
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

func (r *DescribeStarRocksQueryInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStarRocksQueryInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "Page")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStarRocksQueryInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStarRocksQueryInfoResponseParams struct {
	// 总数，分页查询时使用
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Starrocks 查询信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	StarRocksQueryInfoList []*StarRocksQueryInfo `json:"StarRocksQueryInfoList,omitnil,omitempty" name:"StarRocksQueryInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStarRocksQueryInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStarRocksQueryInfoResponseParams `json:"Response"`
}

func (r *DescribeStarRocksQueryInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStarRocksQueryInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrinoQueryInfoRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 获取查询信息开始时间 (s)
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 获取查询信息结束时间 (s)
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页查询时的分页大小，最小1，最大100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页查询时的页号，从1开始
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type DescribeTrinoQueryInfoRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 获取查询信息开始时间 (s)
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 获取查询信息结束时间 (s)
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页查询时的分页大小，最小1，最大100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页查询时的页号，从1开始
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

func (r *DescribeTrinoQueryInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrinoQueryInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "Page")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrinoQueryInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrinoQueryInfoResponseParams struct {
	// 总数，分页查询时使用
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 查询结果数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryInfoList []*TrinoQueryInfo `json:"QueryInfoList,omitnil,omitempty" name:"QueryInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrinoQueryInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrinoQueryInfoResponseParams `json:"Response"`
}

func (r *DescribeTrinoQueryInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrinoQueryInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersForUserManagerRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 页码
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 分页的大小。
	// 默认查询全部；PageNo和PageSize不合理的设置，都是查询全部
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询用户列表过滤器
	UserManagerFilter *UserManagerFilter `json:"UserManagerFilter,omitnil,omitempty" name:"UserManagerFilter"`

	// 是否需要keytab文件的信息，仅对开启kerberos的集群有效，默认为false
	NeedKeytabInfo *bool `json:"NeedKeytabInfo,omitnil,omitempty" name:"NeedKeytabInfo"`
}

type DescribeUsersForUserManagerRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 页码
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 分页的大小。
	// 默认查询全部；PageNo和PageSize不合理的设置，都是查询全部
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询用户列表过滤器
	UserManagerFilter *UserManagerFilter `json:"UserManagerFilter,omitnil,omitempty" name:"UserManagerFilter"`

	// 是否需要keytab文件的信息，仅对开启kerberos的集群有效，默认为false
	NeedKeytabInfo *bool `json:"NeedKeytabInfo,omitnil,omitempty" name:"NeedKeytabInfo"`
}

func (r *DescribeUsersForUserManagerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersForUserManagerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "UserManagerFilter")
	delete(f, "NeedKeytabInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsersForUserManagerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersForUserManagerResponseParams struct {
	// 总数
	TotalCnt *int64 `json:"TotalCnt,omitnil,omitempty" name:"TotalCnt"`

	// 用户信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserManagerUserList []*UserManagerUserBriefInfo `json:"UserManagerUserList,omitnil,omitempty" name:"UserManagerUserList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUsersForUserManagerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsersForUserManagerResponseParams `json:"Response"`
}

func (r *DescribeUsersForUserManagerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersForUserManagerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeYarnApplicationsRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 起始时间秒
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间秒，EndTime-StartTime不得超过1天秒数86400
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页偏移量，Offset=0表示第一页；如果limit=100，Offset=1，则表示第二页，数据第101条开始查询，返回100条数据；如果limit=100，Offset=2，则表示第三页，数据第201条开始查询，返回100条数据。依次类推
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，合法范围[1,100]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeYarnApplicationsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 起始时间秒
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间秒，EndTime-StartTime不得超过1天秒数86400
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页偏移量，Offset=0表示第一页；如果limit=100，Offset=1，则表示第二页，数据第101条开始查询，返回100条数据；如果limit=100，Offset=2，则表示第三页，数据第201条开始查询，返回100条数据。依次类推
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，合法范围[1,100]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeYarnApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeYarnApplicationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeYarnApplicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeYarnApplicationsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 结果列表
	Results []*YarnApplication `json:"Results,omitnil,omitempty" name:"Results"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeYarnApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeYarnApplicationsResponseParams `json:"Response"`
}

func (r *DescribeYarnApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeYarnApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeYarnQueueRequestParams struct {
	// 集群Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 调度器，可选值：
	// 
	// 1. capacity
	// 2. fair
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`
}

type DescribeYarnQueueRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 调度器，可选值：
	// 
	// 1. capacity
	// 2. fair
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`
}

func (r *DescribeYarnQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeYarnQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Scheduler")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeYarnQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeYarnQueueResponseParams struct {
	// 队列信息。是一个对象转成的json字符串，对应的golang结构体如下所示，例如`QueueWithConfigSetForFairScheduler`的第一个字段`Name`：```Name                         string                               `json:"name"` //队列名称```- `Name`：字段名- `string`：字段类型- `json:"name"`：表示在序列化和反序列化`json`时，对应的`json key`，下面以`json key`来指代- `//`：后面的注释内容对应页面上看到的名称字段类型以`*`开头的表示取值可能为json规范下的null，不同的语言需要使用能表达null的类型来接收，例如java的包装类型；字段类型以`[]`开头的表示是数组类型；`json key`在调用`ModifyYarnQueueV2 `接口也会使用。- 公平调度器```type QueueWithConfigSetForFairScheduler struct {	Name                         string                               `json:"name"` //队列名称	MyId                         string                  `json:"myId"` // 队列id，用于编辑、删除、克隆时使用	ParentId                     string                  `json:"parentId"`  // 父队列Id	Type                         *string                              `json:"type"` // 队列归属。parent或空，当确定某个队列是父队列，且没有子队列时，才可以设置，通常用来支持放置策略nestedUserQueue	AclSubmitApps                *AclForYarnQueue                     `json:"aclSubmitApps"` // 提交访问控制	AclAdministerApps            *AclForYarnQueue                     `json:"aclAdministerApps"` // 管理访问控制	MinSharePreemptionTimeout    *int                                 `json:"minSharePreemptionTimeout"` // 最小共享优先权超时时间	FairSharePreemptionTimeout   *int                                 `json:"fairSharePreemptionTimeout"` // 公平份额抢占超时时间	FairSharePreemptionThreshold *float32                             `json:"fairSharePreemptionThreshold"` // 公平份额抢占阈值。取值 （0，1]	AllowPreemptionFrom          *bool                                `json:"allowPreemptionFrom"`                                        // 抢占模式	SchedulingPolicy             *string                              `json:"schedulingPolicy"`  // 调度策略，取值有drf、fair、fifo	IsDefault                    *bool                                `json:"isDefault"` // 是否是root.default队列	IsRoot                       *bool                                `json:"isRoot"` // 是否是root队列	ConfigSets                   []ConfigSetForFairScheduler          `json:"configSets"` // 配置集设置	Children                     []QueueWithConfigSetForFairScheduler `json:"queues"` // 子队列信息。递归}type AclForYarnQueue struct {	User  *string `json:"user"` //用户名	Group *string `json:"group"`//组名}type ConfigSetForFairScheduler struct {	Name              string        `json:"name"` // 配置集名称	MinResources      *YarnResource `json:"minResources"` // 最小资源量	MaxResources      *YarnResource `json:"maxResources"` // 最大资源量	MaxChildResources *YarnResource `json:"maxChildResources"` // 能够分配给为未声明子队列的最大资源量	MaxRunningApps    *int          `json:"maxRunningApps"` // 最高可同时处于运行的App数量	Weight            *float32      `json:"weight"`                   // 权重	MaxAMShare        *float32      `json:"maxAMShare"` // App Master最大份额}type YarnResource struct {	Vcores *int `json:"vcores"`	Memory *int `json:"memory"`	Type *string `json:"type"` // 当值为`percent`时，表示使用的百分比，否则就是使用的绝对数值}```- 容量调度器```type QueueForCapacitySchedulerV3 struct {	Name                       string                `json:"name"` // 队列名称	MyId                       string                `json:"myId"` // 队列id，用于编辑、删除、克隆时使用	ParentId                   string                `json:"parentId"` // 父队列Id	Configs                    []ConfigForCapacityV3 `json:"configs"` //配置集设置	State                      *string         `json:"state"` // 资源池状态	DefaultNodeLabelExpression *string               `json:"default-node-label-expression"` // 默认标签表达式	AclSubmitApps              *AclForYarnQueue      `json:"acl_submit_applications"` // 提交访问控制	AclAdminQueue              *AclForYarnQueue      `json:"acl_administer_queue"` //管理访问控制	MaxAllocationMB *int32 `json:"maximum-allocation-mb"` // 分配Container最大内存数量	MaxAllocationVcores *int32                         `json:"maximum-allocation-vcores"` // Container最大vCore数量	IsDefault           *bool                          `json:"isDefault"`// 是否是root.default队列	IsRoot              *bool                          `json:"isRoot"` // 是否是root队列	Queues              []*QueueForCapacitySchedulerV3 `json:"queues"`//子队列信息。递归}type ConfigForCapacityV3 struct {	Name                string          `json:"configName"` // 配置集名称	Labels              []CapacityLabel `json:"labels"` // 标签信息	MinUserLimitPercent *int32          `json:"minimum-user-limit-percent"` // 用户最小容量	UserLimitFactor     *float32        `json:"user-limit-factor" valid:"rangeExcludeLeft(0|)"`  // 用户资源因子	MaxApps *int32 `json:"maximum-applications" valid:"rangeExcludeLeft(0|)"` // 最大应用数Max-Applications	MaxAmPercent               *float32 `json:"maximum-am-resource-percent"` // 最大AM比例	DefaultApplicationPriority *int32   `json:"default-application-priority"` // 资源池优先级}type CapacityLabel struct {	Name        string   `json:"labelName"`	Capacity    *float32 `json:"capacity"`  // 容量	MaxCapacity *float32 `json:"maximum-capacity"` //最大容量}type AclForYarnQueue struct {	User  *string `json:"user"` //用户名	Group *string `json:"group"`//组名}```
	Queue *string `json:"Queue,omitnil,omitempty" name:"Queue"`

	// 版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeYarnQueueResponse struct {
	*tchttp.BaseResponse
	Response *DescribeYarnQueueResponseParams `json:"Response"`
}

func (r *DescribeYarnQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeYarnQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeYarnScheduleHistoryRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 页码
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页大小
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 调度器类型 可选值为“ALL”，"Capacity Scheduler", "Fair Scheduler"
	SchedulerType *string `json:"SchedulerType,omitnil,omitempty" name:"SchedulerType"`

	// 任务类型0:等待执行，1:执行中，2：完成，-1:失败 ，-99:全部
	TaskState *int64 `json:"TaskState,omitnil,omitempty" name:"TaskState"`
}

type DescribeYarnScheduleHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 页码
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页大小
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 调度器类型 可选值为“ALL”，"Capacity Scheduler", "Fair Scheduler"
	SchedulerType *string `json:"SchedulerType,omitnil,omitempty" name:"SchedulerType"`

	// 任务类型0:等待执行，1:执行中，2：完成，-1:失败 ，-99:全部
	TaskState *int64 `json:"TaskState,omitnil,omitempty" name:"TaskState"`
}

func (r *DescribeYarnScheduleHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeYarnScheduleHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SchedulerType")
	delete(f, "TaskState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeYarnScheduleHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeYarnScheduleHistoryResponseParams struct {
	// 任务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*SchedulerTaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 任务详情总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 调度类型筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerNameList []*string `json:"SchedulerNameList,omitnil,omitempty" name:"SchedulerNameList"`

	// 状态筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	StateList []*int64 `json:"StateList,omitnil,omitempty" name:"StateList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeYarnScheduleHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeYarnScheduleHistoryResponseParams `json:"Response"`
}

func (r *DescribeYarnScheduleHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeYarnScheduleHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiffDetail struct {
	// tab页的头
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 变化项的个数
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 要渲染的明细数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rows []*DiffDetailItem `json:"Rows,omitnil,omitempty" name:"Rows"`

	// 要渲染的头部信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Header []*DiffHeader `json:"Header,omitnil,omitempty" name:"Header"`
}

type DiffDetailItem struct {
	// 属性
	Attribute *string `json:"Attribute,omitnil,omitempty" name:"Attribute"`

	// 当前生效
	InEffect *string `json:"InEffect,omitnil,omitempty" name:"InEffect"`

	// 待生效
	PendingEffectiveness *string `json:"PendingEffectiveness,omitnil,omitempty" name:"PendingEffectiveness"`

	// 操作
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 队列
	Queue *string `json:"Queue,omitnil,omitempty" name:"Queue"`

	// 配置集
	ConfigSet *string `json:"ConfigSet,omitnil,omitempty" name:"ConfigSet"`

	// 标签
	LabelName *string `json:"LabelName,omitnil,omitempty" name:"LabelName"`

	// 当前所在位置
	InEffectIndex *string `json:"InEffectIndex,omitnil,omitempty" name:"InEffectIndex"`

	// 待生效的位置
	PendingEffectIndex *string `json:"PendingEffectIndex,omitnil,omitempty" name:"PendingEffectIndex"`

	// 计划模式名称
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// 标签
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 放置规则
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

type DiffHeader struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// ID，前端会使用
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type Disk struct {
	// 数据盘类型，创建EMR容器集群实例可选
	// <li> SSD云盘: CLOUD_SSD</li>
	// <li>高效云盘: CLOUD_PREMIUM</li>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 单块大小GB
	DiskCapacity *int64 `json:"DiskCapacity,omitnil,omitempty" name:"DiskCapacity"`

	// 数据盘数量
	DiskNumber *int64 `json:"DiskNumber,omitnil,omitempty" name:"DiskNumber"`
}

type DiskGroup struct {
	// 磁盘规格。
	Spec *DiskSpec `json:"Spec,omitnil,omitempty" name:"Spec"`

	// 同类型磁盘数量。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type DiskSpec struct {
	// 磁盘类型。
	// LOCAL_BASIC  本地盘。
	// CLOUD_BASIC 云硬盘。
	// LOCAL_SSD 本地SSD。
	// CLOUD_SSD 云SSD。
	// CLOUD_PREMIUM 高效云盘。
	// CLOUD_HSSD 增强型云SSD。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 磁盘大小，单位GB。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

type DiskSpecInfo struct {
	// 磁盘数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 系统盘类型 取值范围：
	// <li>CLOUD_SSD：表示云SSD。</li>
	// <li>CLOUD_PREMIUM：表示高效云盘。</li>
	// <li>CLOUD_BASIC：表示云硬盘。</li>
	// <li>LOCAL_BASIC：表示本地盘。</li>
	// <li>LOCAL_SSD：表示本地SSD。</li>
	// 
	// 数据盘类型 取值范围：
	// <li>CLOUD_SSD：表示云SSD。</li>
	// <li>CLOUD_PREMIUM：表示高效云盘。</li>
	// <li>CLOUD_BASIC：表示云硬盘。</li>
	// <li>LOCAL_BASIC：表示本地盘。</li>
	// <li>LOCAL_SSD：表示本地SSD。</li>
	// <li>CLOUD_HSSD：表示增强型SSD云硬盘。</li>
	// <li>CLOUD_THROUGHPUT：表示吞吐型云硬盘。</li>
	// <li>CLOUD_TSSD：表示极速型SSD云硬盘。</li>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 数据容量，单位为GB
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 额外性能
	ExtraPerformance *int64 `json:"ExtraPerformance,omitnil,omitempty" name:"ExtraPerformance"`
}

type Dps struct {
	// 时间戳
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 采样值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type DynamicPodSpec struct {
	// 需求最小cpu核数
	RequestCpu *float64 `json:"RequestCpu,omitnil,omitempty" name:"RequestCpu"`

	// 需求最大cpu核数
	LimitCpu *float64 `json:"LimitCpu,omitnil,omitempty" name:"LimitCpu"`

	// 需求最小memory，单位MB
	RequestMemory *float64 `json:"RequestMemory,omitnil,omitempty" name:"RequestMemory"`

	// 需求最大memory，单位MB
	LimitMemory *float64 `json:"LimitMemory,omitnil,omitempty" name:"LimitMemory"`
}

type EmrListInstance struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 常见状态描述：集群生产中,集群运行中,集群创建中,集群已关闭,集群已删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 集群名字
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群地域
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户APPID
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 创建时间
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 运行时间
	RunTime *string `json:"RunTime,omitnil,omitempty" name:"RunTime"`

	// 集群IP
	MasterIp *string `json:"MasterIp,omitnil,omitempty" name:"MasterIp"`

	// 集群版本
	EmrVersion *string `json:"EmrVersion,omitnil,omitempty" name:"EmrVersion"`

	// 集群计费类型
	ChargeType *uint64 `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// emr ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 产品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *uint64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *uint64 `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *uint64 `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 地区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 状态码, 取值为-2(集群已删除), -1(集群已关闭), 0(集群生产中), 2(集群运行中), 3(集群创建中)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 告警信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmInfo *string `json:"AlarmInfo,omitnil,omitempty" name:"AlarmInfo"`

	// 是否是woodpecker集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsWoodpeckerCluster *uint64 `json:"IsWoodpeckerCluster,omitnil,omitempty" name:"IsWoodpeckerCluster"`

	// Vpc中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 子网中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// 字符串VpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 字符串子网
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterClass *string `json:"ClusterClass,omitnil,omitempty" name:"ClusterClass"`

	// 是否为跨AZ集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsMultiZoneCluster *bool `json:"IsMultiZoneCluster,omitnil,omitempty" name:"IsMultiZoneCluster"`

	// 是否手戳集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsHandsCluster *bool `json:"IsHandsCluster,omitnil,omitempty" name:"IsHandsCluster"`

	// 体外客户端组件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutSideSoftInfo []*SoftDependInfo `json:"OutSideSoftInfo,omitnil,omitempty" name:"OutSideSoftInfo"`

	// 当前集群的应用场景是否支持体外客户端
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSupportOutsideCluster *bool `json:"IsSupportOutsideCluster,omitnil,omitempty" name:"IsSupportOutsideCluster"`

	// 是否专有集群场景集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDedicatedCluster *bool `json:"IsDedicatedCluster,omitnil,omitempty" name:"IsDedicatedCluster"`
}

type EmrPrice struct {
	// 刊例价格
	OriginalCost *string `json:"OriginalCost,omitnil,omitempty" name:"OriginalCost"`

	// 折扣价格
	DiscountCost *string `json:"DiscountCost,omitnil,omitempty" name:"DiscountCost"`

	// 单位
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 询价配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriceSpec *PriceResource `json:"PriceSpec,omitnil,omitempty" name:"PriceSpec"`

	// 是否支持竞价实例
	SupportSpotPaid *bool `json:"SupportSpotPaid,omitnil,omitempty" name:"SupportSpotPaid"`
}

type EmrProductConfigDetail struct {
	// 软件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SoftInfo []*string `json:"SoftInfo,omitnil,omitempty" name:"SoftInfo"`

	// Master节点个数
	MasterNodeSize *int64 `json:"MasterNodeSize,omitnil,omitempty" name:"MasterNodeSize"`

	// Core节点个数
	CoreNodeSize *int64 `json:"CoreNodeSize,omitnil,omitempty" name:"CoreNodeSize"`

	// Task节点个数
	TaskNodeSize *int64 `json:"TaskNodeSize,omitnil,omitempty" name:"TaskNodeSize"`

	// Common节点个数
	ComNodeSize *int64 `json:"ComNodeSize,omitnil,omitempty" name:"ComNodeSize"`

	// Master节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterResource *ResourceDetail `json:"MasterResource,omitnil,omitempty" name:"MasterResource"`

	// Core节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoreResource *ResourceDetail `json:"CoreResource,omitnil,omitempty" name:"CoreResource"`

	// Task节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskResource *ResourceDetail `json:"TaskResource,omitnil,omitempty" name:"TaskResource"`

	// Common节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComResource *ResourceDetail `json:"ComResource,omitnil,omitempty" name:"ComResource"`

	// 是否使用COS
	OnCos *bool `json:"OnCos,omitnil,omitempty" name:"OnCos"`

	// 收费类型
	ChargeType *int64 `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// Router节点个数
	RouterNodeSize *int64 `json:"RouterNodeSize,omitnil,omitempty" name:"RouterNodeSize"`

	// 是否支持HA
	SupportHA *bool `json:"SupportHA,omitnil,omitempty" name:"SupportHA"`

	// 是否支持安全模式
	SecurityOn *bool `json:"SecurityOn,omitnil,omitempty" name:"SecurityOn"`

	// 安全组名称
	SecurityGroup *string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 是否开启Cbs加密
	CbsEncrypt *int64 `json:"CbsEncrypt,omitnil,omitempty" name:"CbsEncrypt"`

	// 自定义应用角色。
	ApplicationRole *string `json:"ApplicationRole,omitnil,omitempty" name:"ApplicationRole"`

	// 安全组
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`

	// SSH密钥Id
	PublicKeyId *string `json:"PublicKeyId,omitnil,omitempty" name:"PublicKeyId"`
}

type EmrProductConfigOutter struct {
	// 软件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SoftInfo []*string `json:"SoftInfo,omitnil,omitempty" name:"SoftInfo"`

	// Master节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterNodeSize *int64 `json:"MasterNodeSize,omitnil,omitempty" name:"MasterNodeSize"`

	// Core节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoreNodeSize *int64 `json:"CoreNodeSize,omitnil,omitempty" name:"CoreNodeSize"`

	// Task节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskNodeSize *int64 `json:"TaskNodeSize,omitnil,omitempty" name:"TaskNodeSize"`

	// Common节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComNodeSize *int64 `json:"ComNodeSize,omitnil,omitempty" name:"ComNodeSize"`

	// Master节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterResource *OutterResource `json:"MasterResource,omitnil,omitempty" name:"MasterResource"`

	// Core节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoreResource *OutterResource `json:"CoreResource,omitnil,omitempty" name:"CoreResource"`

	// Task节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskResource *OutterResource `json:"TaskResource,omitnil,omitempty" name:"TaskResource"`

	// Common节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComResource *OutterResource `json:"ComResource,omitnil,omitempty" name:"ComResource"`

	// 是否使用COS
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnCos *bool `json:"OnCos,omitnil,omitempty" name:"OnCos"`

	// 收费类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *int64 `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// Router节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouterNodeSize *int64 `json:"RouterNodeSize,omitnil,omitempty" name:"RouterNodeSize"`

	// 是否支持HA
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportHA *bool `json:"SupportHA,omitnil,omitempty" name:"SupportHA"`

	// 是否支持安全模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityOn *bool `json:"SecurityOn,omitnil,omitempty" name:"SecurityOn"`

	// 集群初始安全组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroup *string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 是否开启Cbs加密
	// 注意：此字段可能返回 null，表示取不到有效值。
	CbsEncrypt *int64 `json:"CbsEncrypt,omitnil,omitempty" name:"CbsEncrypt"`

	// 自定义应用角色。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationRole *string `json:"ApplicationRole,omitnil,omitempty" name:"ApplicationRole"`

	// 安全组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`

	// SSH密钥Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicKeyId *string `json:"PublicKeyId,omitnil,omitempty" name:"PublicKeyId"`
}

type Execution struct {
	// 任务类型，目前支持以下类型。
	// 1. “MR”，将通过hadoop jar的方式提交。
	// 2. "HIVE"，将通过hive -f的方式提交。
	// 3. "SPARK"，将通过spark-submit的方式提交。
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 任务参数，提供除提交指令以外的参数。
	Args []*string `json:"Args,omitnil,omitempty" name:"Args"`
}

type ExternalAccess struct {
	// 外部访问类型，当前仅支持CLB字段
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// CLB设置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CLBServer *CLBSetting `json:"CLBServer,omitnil,omitempty" name:"CLBServer"`
}

type ExternalService struct {
	// 共用组件类型，EMR/CUSTOM
	ShareType *string `json:"ShareType,omitnil,omitempty" name:"ShareType"`

	// 共用组件名
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 共用组件集群
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自定义参数集合
	CustomServiceDefineList []*CustomServiceDefine `json:"CustomServiceDefineList,omitnil,omitempty" name:"CustomServiceDefineList"`
}

type FairGlobalConfig struct {
	// 对应与页面的<p>程序上限</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserMaxAppsDefault *int64 `json:"UserMaxAppsDefault,omitnil,omitempty" name:"UserMaxAppsDefault"`
}

type Filters struct {
	// 字段名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤字段值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FlowExtraDetail struct {
	// 额外信息Title
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*FlowParamsDesc `json:"Detail,omitnil,omitempty" name:"Detail"`
}

type FlowParam struct {
	// 流程参数key
	// TraceId：通过TraceId查询
	// FlowId： 通过FlowId查询
	FKey *string `json:"FKey,omitnil,omitempty" name:"FKey"`

	// 参数value
	FValue *string `json:"FValue,omitnil,omitempty" name:"FValue"`
}

type FlowParamsDesc struct {
	// 参数key
	PKey *string `json:"PKey,omitnil,omitempty" name:"PKey"`

	// 参数value
	PValue *string `json:"PValue,omitnil,omitempty" name:"PValue"`
}

type GroupGlobalConfs struct {
	// 伸缩组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupGlobalConf *AutoScaleResourceConf `json:"GroupGlobalConf,omitnil,omitempty" name:"GroupGlobalConf"`

	// 当前伸缩组扩容出来的节点数量。
	CurrentNodes *int64 `json:"CurrentNodes,omitnil,omitempty" name:"CurrentNodes"`

	// 当前伸缩组扩容出来的后付费节点数量。
	CurrentPostPaidNodes *int64 `json:"CurrentPostPaidNodes,omitnil,omitempty" name:"CurrentPostPaidNodes"`

	// 当前伸缩组扩容出来的竞价实例节点数量。
	CurrentSpotPaidNodes *int64 `json:"CurrentSpotPaidNodes,omitnil,omitempty" name:"CurrentSpotPaidNodes"`
}

type HealthStatus struct {
	// 运行正常
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 运行正常
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 运行正常
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type HiveQuery struct {
	// 查询语句
	Statement *string `json:"Statement,omitnil,omitempty" name:"Statement"`

	// 执行时长
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 开始时间毫秒
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间毫秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 状态
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// appId列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobIds []*string `json:"JobIds,omitnil,omitempty" name:"JobIds"`

	// 执行引擎
	ExecutionEngine *string `json:"ExecutionEngine,omitnil,omitempty" name:"ExecutionEngine"`

	// 查询ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type HostPathVolumeSource struct {
	// 主机路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 主机路径类型，当前默认DirectoryOrCreate
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type HostVolumeContext struct {
	// Pod挂载宿主机的目录。资源对宿主机的挂载点，指定的挂载点对应了宿主机的路径，该挂载点在Pod中作为数据存储目录使用
	VolumePath *string `json:"VolumePath,omitnil,omitempty" name:"VolumePath"`
}

type ImpalaQuery struct {
	// 执行语句
	Statement *string `json:"Statement,omitnil,omitempty" name:"Statement"`

	// 查询ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 运行时间
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 执行状态
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 获取行数
	RowsFetched *int64 `json:"RowsFetched,omitnil,omitempty" name:"RowsFetched"`

	// 用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 默认DB
	DefaultDB *string `json:"DefaultDB,omitnil,omitempty" name:"DefaultDB"`

	// 执行的Coordinator节点
	Coordinator *string `json:"Coordinator,omitnil,omitempty" name:"Coordinator"`

	// 单节点内存峰值
	MaxNodePeakMemoryUsage *string `json:"MaxNodePeakMemoryUsage,omitnil,omitempty" name:"MaxNodePeakMemoryUsage"`

	// 查询类型
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 扫描的HDFS行数
	ScanHDFSRows *int64 `json:"ScanHDFSRows,omitnil,omitempty" name:"ScanHDFSRows"`

	// 扫描的Kudu行数
	ScanKUDURows *int64 `json:"ScanKUDURows,omitnil,omitempty" name:"ScanKUDURows"`

	// 扫描的总行数
	ScanRowsTotal *int64 `json:"ScanRowsTotal,omitnil,omitempty" name:"ScanRowsTotal"`

	// 读取的总字节数
	TotalBytesRead *int64 `json:"TotalBytesRead,omitnil,omitempty" name:"TotalBytesRead"`

	// 发送的总字节数
	TotalBytesSent *int64 `json:"TotalBytesSent,omitnil,omitempty" name:"TotalBytesSent"`

	// CPU总时间
	TotalCpuTime *int64 `json:"TotalCpuTime,omitnil,omitempty" name:"TotalCpuTime"`

	// 内部数据发送总量(Bytes)
	TotalInnerBytesSent *int64 `json:"TotalInnerBytesSent,omitnil,omitempty" name:"TotalInnerBytesSent"`

	// 内部扫描数据发送总量(Bytes)
	TotalScanBytesSent *int64 `json:"TotalScanBytesSent,omitnil,omitempty" name:"TotalScanBytesSent"`

	// 预估单节点内存
	EstimatedPerHostMemBytes *int64 `json:"EstimatedPerHostMemBytes,omitnil,omitempty" name:"EstimatedPerHostMemBytes"`

	// 从缓存中获取的数据行数
	NumRowsFetchedFromCache *int64 `json:"NumRowsFetchedFromCache,omitnil,omitempty" name:"NumRowsFetchedFromCache"`

	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 单节点内存峰值和(Bytes)
	PerNodePeakMemoryBytesSum *int64 `json:"PerNodePeakMemoryBytesSum,omitnil,omitempty" name:"PerNodePeakMemoryBytesSum"`

	// 后端个数
	BackendsCount *int64 `json:"BackendsCount,omitnil,omitempty" name:"BackendsCount"`

	// fragment数
	FragmentInstancesCount *int64 `json:"FragmentInstancesCount,omitnil,omitempty" name:"FragmentInstancesCount"`

	// 剩余未完成Fragment数
	RemainingFragmentCount *int64 `json:"RemainingFragmentCount,omitnil,omitempty" name:"RemainingFragmentCount"`
}

// Predefined struct for user
type InquirePriceRenewEmrRequestParams struct {
	// 实例续费的时长。需要结合TimeUnit一起使用。1表示续费一个月
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 待续费集群ID列表。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 实例计费模式。此处只支持取值为1，表示包年包月。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 实例续费的时间单位。取值范围：
	// <li>m：表示月份。</li>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`
}

type InquirePriceRenewEmrRequest struct {
	*tchttp.BaseRequest
	
	// 实例续费的时长。需要结合TimeUnit一起使用。1表示续费一个月
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 待续费集群ID列表。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 实例计费模式。此处只支持取值为1，表示包年包月。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 实例续费的时间单位。取值范围：
	// <li>m：表示月份。</li>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`
}

func (r *InquirePriceRenewEmrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewEmrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeSpan")
	delete(f, "InstanceId")
	delete(f, "Placement")
	delete(f, "PayMode")
	delete(f, "TimeUnit")
	delete(f, "Currency")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceRenewEmrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewEmrResponseParams struct {
	// 原价，单位为元。
	OriginalCost *float64 `json:"OriginalCost,omitnil,omitempty" name:"OriginalCost"`

	// 折扣价，单位为元。
	DiscountCost *float64 `json:"DiscountCost,omitnil,omitempty" name:"DiscountCost"`

	// 实例续费的时间单位。取值范围：
	// <li>m：表示月份。</li>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 实例续费的时长。
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceRenewEmrResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceRenewEmrResponseParams `json:"Response"`
}

func (r *InquirePriceRenewEmrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewEmrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateInstanceRequestParams struct {
	// 购买实例的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 购买实例的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 是否开启节点高可用。取值范围：
	// <li>0：表示不开启节点高可用。</li>
	// <li>1：表示开启节点高可用。</li>
	SupportHA *uint64 `json:"SupportHA,omitnil,omitempty" name:"SupportHA"`

	// 部署的组件列表。不同的EMR产品ID（ProductId：具体含义参考入参ProductId字段）需要选择不同的必选组件：<li>ProductId为2(EMR-V2.0.1)的时候，必选组件包括：hdfs-2.7.3,yarn-2.7.3,zookeeper-3.4.9,knox-1.2.0</li><li>ProductId为16(EMR-V2.3.0)的时候，必选组件包括：hdfs-2.8.5,yarn-2.8.5,zookeeper-3.5.5,knox-1.2.0</li><li>ProductId为20(EMR-V2.5.0)的时候，必选组件包括：hdfs-2.8.5,yarn-2.8.5,zookeeper-3.6.1,knox-1.2.0</li><li>ProductId为30(EMR-V2.6.0)的时候，必选组件包括：hdfs-2.8.5,yarn-2.8.5,zookeeper-3.6.1,openldap-2.4.44,knox-1.2.0</li><li>ProductId为38(EMR-V2.7.0)的时候，必选组件包括：hdfs-2.8.5,yarn-2.8.5,zookeeper-3.6.3,openldap-2.4.44,knox-1.6.1</li><li>ProductId为57(EMR-V2.8.0)的时候，必选组件包括：hdfs-2.8.5,yarn-2.8.5,zookeeper-3.6.3,openldap-2.4.44,knox-1.6.1</li><li>ProductId为7(EMR-V3.0.0)的时候，必选组件包括：hdfs-3.1.2,yarn-3.1.2,zookeeper-3.4.9,knox-1.2.0</li><li>ProductId为25(EMR-V3.1.0)的时候，必选组件包括：hdfs-3.1.2,yarn-3.1.2,zookeeper-3.6.1,knox-1.2.0</li><li>ProductId为31(EMR-V3.1.1)的时候，必选组件包括：hdfs-3.1.2,yarn-3.1.2,zookeeper-3.6.1,knox-1.2.0</li><li>ProductId为28(EMR-V3.2.0)的时候，必选组件包括：hdfs-3.2.2,yarn-3.2.2,zookeeper-3.6.1,knox-1.2.0</li><li>ProductId为33(EMR-V3.2.1)的时候，必选组件包括：hdfs-3.2.2,yarn-3.2.2,zookeeper-3.6.1,openldap-2.4.44,knox-1.2.0</li><li>ProductId为34(EMR-V3.3.0)的时候，必选组件包括：hdfs-3.2.2,yarn-3.2.2,zookeeper-3.6.1,openldap-2.4.44,knox-1.2.0</li><li>ProductId为37(EMR-V3.4.0)的时候，必选组件包括：hdfs-3.2.2,yarn-3.2.2,zookeeper-3.6.3,openldap-2.4.44,knox-1.6.1</li><li>ProductId为44(EMR-V3.5.0)的时候，必选组件包括：hdfs-3.2.2,yarn-3.2.2,zookeeper-3.6.3,openldap-2.4.44,knox-1.6.1</li><li>ProductId为53(EMR-V3.6.0)的时候，必选组件包括：hdfs-3.2.2,yarn-3.2.2,zookeeper-3.6.3,openldap-2.4.44,knox-1.6.1</li><li>ProductId为58(EMR-V3.6.1)的时候，必选组件包括：hdfs-3.2.2,yarn-3.2.2,zookeeper-3.6.3,openldap-2.4.46,knox-1.6.1</li><li>ProductId为47(EMR-V4.0.0)的时候，必选组件包括：hdfs-3.2.2,yarn-3.2.2,zookeeper-3.6.3,openldap-2.4.44,knox-1.6.1</li>
	Software []*string `json:"Software,omitnil,omitempty" name:"Software"`

	// 询价的节点规格。
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitnil,omitempty" name:"ResourceSpec"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil,omitempty" name:"VPCSettings"`

	// hive共享元数据库类型。取值范围：
	// <li>EMR_NEW_META：表示集群默认创建</li>
	// <li>EMR_EXIT_METE：表示集群使用指定EMR-MetaDB。</li>
	// <li>USER_CUSTOM_META：表示集群使用自定义MetaDB。</li>
	MetaType *string `json:"MetaType,omitnil,omitempty" name:"MetaType"`

	// EMR-MetaDB实例
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitnil,omitempty" name:"UnifyMetaInstanceId"`

	// 自定义MetaDB信息
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitnil,omitempty" name:"MetaDBInfo"`

	// 产品ID，不同产品ID表示不同的EMR产品版本。取值范围：<li>2：表示EMR-V2.0.1</li><li>16：表示EMR-V2.3.0</li><li>20：表示EMR-V2.5.0</li><li>30：表示EMR-V2.6.0</li><li>38：表示EMR-V2.7.0</li><li>57：表示EMR-V2.8.0</li><li>7：表示EMR-V3.0.0</li><li>25：表示EMR-V3.1.0</li><li>31：表示EMR-V3.1.1</li><li>28：表示EMR-V3.2.0</li><li>33：表示EMR-V3.2.1</li><li>34：表示EMR-V3.3.0</li><li>37：表示EMR-V3.4.0</li><li>44：表示EMR-V3.5.0</li><li>53：表示EMR-V3.6.0</li><li>58：表示EMR-V3.6.1</li><li>47：表示EMR-V4.0.0</li>
	ProductId *uint64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 场景化取值：Hadoop-Kudu，Hadoop-Zookeeper，Hadoop-Presto，Hadoop-Hbase
	SceneName *string `json:"SceneName,omitnil,omitempty" name:"SceneName"`

	// 共用组件信息
	ExternalService []*ExternalService `json:"ExternalService,omitnil,omitempty" name:"ExternalService"`

	// 当前默认值为0，跨AZ特性支持后为1
	VersionID *uint64 `json:"VersionID,omitnil,omitempty" name:"VersionID"`

	// 可用区的规格信息
	MultiZoneSettings []*MultiZoneSetting `json:"MultiZoneSettings,omitnil,omitempty" name:"MultiZoneSettings"`
}

type InquiryPriceCreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 购买实例的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 购买实例的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 是否开启节点高可用。取值范围：
	// <li>0：表示不开启节点高可用。</li>
	// <li>1：表示开启节点高可用。</li>
	SupportHA *uint64 `json:"SupportHA,omitnil,omitempty" name:"SupportHA"`

	// 部署的组件列表。不同的EMR产品ID（ProductId：具体含义参考入参ProductId字段）需要选择不同的必选组件：<li>ProductId为2(EMR-V2.0.1)的时候，必选组件包括：hdfs-2.7.3,yarn-2.7.3,zookeeper-3.4.9,knox-1.2.0</li><li>ProductId为16(EMR-V2.3.0)的时候，必选组件包括：hdfs-2.8.5,yarn-2.8.5,zookeeper-3.5.5,knox-1.2.0</li><li>ProductId为20(EMR-V2.5.0)的时候，必选组件包括：hdfs-2.8.5,yarn-2.8.5,zookeeper-3.6.1,knox-1.2.0</li><li>ProductId为30(EMR-V2.6.0)的时候，必选组件包括：hdfs-2.8.5,yarn-2.8.5,zookeeper-3.6.1,openldap-2.4.44,knox-1.2.0</li><li>ProductId为38(EMR-V2.7.0)的时候，必选组件包括：hdfs-2.8.5,yarn-2.8.5,zookeeper-3.6.3,openldap-2.4.44,knox-1.6.1</li><li>ProductId为57(EMR-V2.8.0)的时候，必选组件包括：hdfs-2.8.5,yarn-2.8.5,zookeeper-3.6.3,openldap-2.4.44,knox-1.6.1</li><li>ProductId为7(EMR-V3.0.0)的时候，必选组件包括：hdfs-3.1.2,yarn-3.1.2,zookeeper-3.4.9,knox-1.2.0</li><li>ProductId为25(EMR-V3.1.0)的时候，必选组件包括：hdfs-3.1.2,yarn-3.1.2,zookeeper-3.6.1,knox-1.2.0</li><li>ProductId为31(EMR-V3.1.1)的时候，必选组件包括：hdfs-3.1.2,yarn-3.1.2,zookeeper-3.6.1,knox-1.2.0</li><li>ProductId为28(EMR-V3.2.0)的时候，必选组件包括：hdfs-3.2.2,yarn-3.2.2,zookeeper-3.6.1,knox-1.2.0</li><li>ProductId为33(EMR-V3.2.1)的时候，必选组件包括：hdfs-3.2.2,yarn-3.2.2,zookeeper-3.6.1,openldap-2.4.44,knox-1.2.0</li><li>ProductId为34(EMR-V3.3.0)的时候，必选组件包括：hdfs-3.2.2,yarn-3.2.2,zookeeper-3.6.1,openldap-2.4.44,knox-1.2.0</li><li>ProductId为37(EMR-V3.4.0)的时候，必选组件包括：hdfs-3.2.2,yarn-3.2.2,zookeeper-3.6.3,openldap-2.4.44,knox-1.6.1</li><li>ProductId为44(EMR-V3.5.0)的时候，必选组件包括：hdfs-3.2.2,yarn-3.2.2,zookeeper-3.6.3,openldap-2.4.44,knox-1.6.1</li><li>ProductId为53(EMR-V3.6.0)的时候，必选组件包括：hdfs-3.2.2,yarn-3.2.2,zookeeper-3.6.3,openldap-2.4.44,knox-1.6.1</li><li>ProductId为58(EMR-V3.6.1)的时候，必选组件包括：hdfs-3.2.2,yarn-3.2.2,zookeeper-3.6.3,openldap-2.4.46,knox-1.6.1</li><li>ProductId为47(EMR-V4.0.0)的时候，必选组件包括：hdfs-3.2.2,yarn-3.2.2,zookeeper-3.6.3,openldap-2.4.44,knox-1.6.1</li>
	Software []*string `json:"Software,omitnil,omitempty" name:"Software"`

	// 询价的节点规格。
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitnil,omitempty" name:"ResourceSpec"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil,omitempty" name:"VPCSettings"`

	// hive共享元数据库类型。取值范围：
	// <li>EMR_NEW_META：表示集群默认创建</li>
	// <li>EMR_EXIT_METE：表示集群使用指定EMR-MetaDB。</li>
	// <li>USER_CUSTOM_META：表示集群使用自定义MetaDB。</li>
	MetaType *string `json:"MetaType,omitnil,omitempty" name:"MetaType"`

	// EMR-MetaDB实例
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitnil,omitempty" name:"UnifyMetaInstanceId"`

	// 自定义MetaDB信息
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitnil,omitempty" name:"MetaDBInfo"`

	// 产品ID，不同产品ID表示不同的EMR产品版本。取值范围：<li>2：表示EMR-V2.0.1</li><li>16：表示EMR-V2.3.0</li><li>20：表示EMR-V2.5.0</li><li>30：表示EMR-V2.6.0</li><li>38：表示EMR-V2.7.0</li><li>57：表示EMR-V2.8.0</li><li>7：表示EMR-V3.0.0</li><li>25：表示EMR-V3.1.0</li><li>31：表示EMR-V3.1.1</li><li>28：表示EMR-V3.2.0</li><li>33：表示EMR-V3.2.1</li><li>34：表示EMR-V3.3.0</li><li>37：表示EMR-V3.4.0</li><li>44：表示EMR-V3.5.0</li><li>53：表示EMR-V3.6.0</li><li>58：表示EMR-V3.6.1</li><li>47：表示EMR-V4.0.0</li>
	ProductId *uint64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 场景化取值：Hadoop-Kudu，Hadoop-Zookeeper，Hadoop-Presto，Hadoop-Hbase
	SceneName *string `json:"SceneName,omitnil,omitempty" name:"SceneName"`

	// 共用组件信息
	ExternalService []*ExternalService `json:"ExternalService,omitnil,omitempty" name:"ExternalService"`

	// 当前默认值为0，跨AZ特性支持后为1
	VersionID *uint64 `json:"VersionID,omitnil,omitempty" name:"VersionID"`

	// 可用区的规格信息
	MultiZoneSettings []*MultiZoneSetting `json:"MultiZoneSettings,omitnil,omitempty" name:"MultiZoneSettings"`
}

func (r *InquiryPriceCreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	delete(f, "Currency")
	delete(f, "PayMode")
	delete(f, "SupportHA")
	delete(f, "Software")
	delete(f, "ResourceSpec")
	delete(f, "Placement")
	delete(f, "VPCSettings")
	delete(f, "MetaType")
	delete(f, "UnifyMetaInstanceId")
	delete(f, "MetaDBInfo")
	delete(f, "ProductId")
	delete(f, "SceneName")
	delete(f, "ExternalService")
	delete(f, "VersionID")
	delete(f, "MultiZoneSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateInstanceResponseParams struct {
	// 原价，单位为元。
	OriginalCost *float64 `json:"OriginalCost,omitnil,omitempty" name:"OriginalCost"`

	// 折扣价，单位为元。
	DiscountCost *float64 `json:"DiscountCost,omitnil,omitempty" name:"DiscountCost"`

	// 购买实例的时间单位。取值范围：
	// <li>s：表示秒。</li>
	// <li>m：表示月份。</li>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 购买实例的时长。
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 价格清单
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriceList []*ZoneDetailPriceResult `json:"PriceList,omitnil,omitempty" name:"PriceList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceCreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceCreateInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceCreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewInstanceRequestParams struct {
	// 实例续费的时长。需要结合TimeUnit一起使用。1表示续费一个月
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 实例计费模式。此处只支持取值为1，表示包年包月。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 待续费节点的资源ID列表。资源ID形如：emr-vm-xxxxxxxx。有效的资源ID可通过登录[控制台](https://console.cloud.tencent.com/emr)查询。
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 实例续费的时间单位。取值范围：
	// <li>m：表示月份。</li>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 是否按量转包年包月。0：否，1：是。
	ModifyPayMode *int64 `json:"ModifyPayMode,omitnil,omitempty" name:"ModifyPayMode"`

	// 是否需要每个节点续费价格
	NeedDetail *bool `json:"NeedDetail,omitnil,omitempty" name:"NeedDetail"`

	// 集群id，如果需要集群所有包年包月节点续费信息，可以填写该参数
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type InquiryPriceRenewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例续费的时长。需要结合TimeUnit一起使用。1表示续费一个月
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 实例计费模式。此处只支持取值为1，表示包年包月。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 待续费节点的资源ID列表。资源ID形如：emr-vm-xxxxxxxx。有效的资源ID可通过登录[控制台](https://console.cloud.tencent.com/emr)查询。
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 实例续费的时间单位。取值范围：
	// <li>m：表示月份。</li>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 是否按量转包年包月。0：否，1：是。
	ModifyPayMode *int64 `json:"ModifyPayMode,omitnil,omitempty" name:"ModifyPayMode"`

	// 是否需要每个节点续费价格
	NeedDetail *bool `json:"NeedDetail,omitnil,omitempty" name:"NeedDetail"`

	// 集群id，如果需要集群所有包年包月节点续费信息，可以填写该参数
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *InquiryPriceRenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeSpan")
	delete(f, "PayMode")
	delete(f, "ResourceIds")
	delete(f, "TimeUnit")
	delete(f, "Currency")
	delete(f, "Placement")
	delete(f, "ModifyPayMode")
	delete(f, "NeedDetail")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceRenewInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewInstanceResponseParams struct {
	// 原价，单位为元。
	OriginalCost *float64 `json:"OriginalCost,omitnil,omitempty" name:"OriginalCost"`

	// 折扣价，单位为元。
	DiscountCost *float64 `json:"DiscountCost,omitnil,omitempty" name:"DiscountCost"`

	// 实例续费的时间单位。取值范围：
	// <li>m：表示月份。</li>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 实例续费的时长。
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 价格详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriceDetail []*PriceDetail `json:"PriceDetail,omitnil,omitempty" name:"PriceDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceRenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceRenewInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceRenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceScaleOutInstanceRequestParams struct {
	// 扩容的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 扩容的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 实例所属的可用区ID，例如100003。该参数可以通过调用 [DescribeZones](https://cloud.tencent.com/document/api/213/15707) 的返回值中的ZoneId字段来获取。
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 扩容的Core节点数量。
	CoreCount *uint64 `json:"CoreCount,omitnil,omitempty" name:"CoreCount"`

	// 扩容的Task节点数量。
	TaskCount *uint64 `json:"TaskCount,omitnil,omitempty" name:"TaskCount"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 扩容的Router节点数量。
	RouterCount *uint64 `json:"RouterCount,omitnil,omitempty" name:"RouterCount"`

	// 扩容的Master节点数量。
	MasterCount *uint64 `json:"MasterCount,omitnil,omitempty" name:"MasterCount"`

	// 类型为ComputeResource和EMR以及默认，默认为EMR
	ResourceBaseType *string `json:"ResourceBaseType,omitnil,omitempty" name:"ResourceBaseType"`

	// 计算资源id
	ComputeResourceId *string `json:"ComputeResourceId,omitnil,omitempty" name:"ComputeResourceId"`

	// 扩容资源类型
	HardwareResourceType *string `json:"HardwareResourceType,omitnil,omitempty" name:"HardwareResourceType"`
}

type InquiryPriceScaleOutInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 扩容的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 扩容的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 实例所属的可用区ID，例如100003。该参数可以通过调用 [DescribeZones](https://cloud.tencent.com/document/api/213/15707) 的返回值中的ZoneId字段来获取。
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 扩容的Core节点数量。
	CoreCount *uint64 `json:"CoreCount,omitnil,omitempty" name:"CoreCount"`

	// 扩容的Task节点数量。
	TaskCount *uint64 `json:"TaskCount,omitnil,omitempty" name:"TaskCount"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 扩容的Router节点数量。
	RouterCount *uint64 `json:"RouterCount,omitnil,omitempty" name:"RouterCount"`

	// 扩容的Master节点数量。
	MasterCount *uint64 `json:"MasterCount,omitnil,omitempty" name:"MasterCount"`

	// 类型为ComputeResource和EMR以及默认，默认为EMR
	ResourceBaseType *string `json:"ResourceBaseType,omitnil,omitempty" name:"ResourceBaseType"`

	// 计算资源id
	ComputeResourceId *string `json:"ComputeResourceId,omitnil,omitempty" name:"ComputeResourceId"`

	// 扩容资源类型
	HardwareResourceType *string `json:"HardwareResourceType,omitnil,omitempty" name:"HardwareResourceType"`
}

func (r *InquiryPriceScaleOutInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceScaleOutInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	delete(f, "ZoneId")
	delete(f, "PayMode")
	delete(f, "InstanceId")
	delete(f, "CoreCount")
	delete(f, "TaskCount")
	delete(f, "Currency")
	delete(f, "RouterCount")
	delete(f, "MasterCount")
	delete(f, "ResourceBaseType")
	delete(f, "ComputeResourceId")
	delete(f, "HardwareResourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceScaleOutInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceScaleOutInstanceResponseParams struct {
	// 原价，单位为元。
	OriginalCost *string `json:"OriginalCost,omitnil,omitempty" name:"OriginalCost"`

	// 折扣价，单位为元。
	DiscountCost *string `json:"DiscountCost,omitnil,omitempty" name:"DiscountCost"`

	// 扩容的时间单位。取值范围：
	// <li>s：表示秒。</li>
	// <li>m：表示月份。</li>
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 询价的节点规格。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriceSpec *PriceResource `json:"PriceSpec,omitnil,omitempty" name:"PriceSpec"`

	// 对应入参MultipleResources中多个规格的询价结果，其它出参返回的是第一个规格的询价结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultipleEmrPrice []*EmrPrice `json:"MultipleEmrPrice,omitnil,omitempty" name:"MultipleEmrPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceScaleOutInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceScaleOutInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceScaleOutInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceScaleOutInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpdateInstanceRequestParams struct {
	// 变配的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 变配的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 节点变配的目标配置。
	UpdateSpec *UpdateInstanceSettings `json:"UpdateSpec,omitnil,omitempty" name:"UpdateSpec"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 批量变配资源ID列表
	ResourceIdList []*string `json:"ResourceIdList,omitnil,omitempty" name:"ResourceIdList"`
}

type InquiryPriceUpdateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 变配的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 变配的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 节点变配的目标配置。
	UpdateSpec *UpdateInstanceSettings `json:"UpdateSpec,omitnil,omitempty" name:"UpdateSpec"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 批量变配资源ID列表
	ResourceIdList []*string `json:"ResourceIdList,omitnil,omitempty" name:"ResourceIdList"`
}

func (r *InquiryPriceUpdateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpdateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	delete(f, "PayMode")
	delete(f, "UpdateSpec")
	delete(f, "Placement")
	delete(f, "Currency")
	delete(f, "ResourceIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceUpdateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpdateInstanceResponseParams struct {
	// 原价，单位为元。
	OriginalCost *float64 `json:"OriginalCost,omitnil,omitempty" name:"OriginalCost"`

	// 折扣价，单位为元。
	DiscountCost *float64 `json:"DiscountCost,omitnil,omitempty" name:"DiscountCost"`

	// 变配的时间单位。取值范围：
	// <li>s：表示秒。</li>
	// <li>m：表示月份。</li>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 变配的时长。
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 价格详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriceDetail []*PriceDetail `json:"PriceDetail,omitnil,omitempty" name:"PriceDetail"`

	// 新配置价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewConfigPrice *PriceResult `json:"NewConfigPrice,omitnil,omitempty" name:"NewConfigPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceUpdateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceUpdateInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceUpdateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpdateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InsightResult struct {
	// 当Type为HIVE时，是Hive查询ID，当Type为MAPREDUCE，SPARK，TEZ时则是YarnAppID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 洞察应用的类型，HIVE,SPARK,MAPREDUCE,TEZ
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 洞察规则ID
	// HIVE-ScanManyMeta:元数据扫描过多
	// HIVE-ScanManyPartition:大表扫描
	// HIVE-SlowCompile:编译耗时过长
	// HIVE-UnSuitableConfig:不合理参数
	// MAPREDUCE-MapperDataSkew:Map数据倾斜
	// MAPREDUCE-MapperMemWaste:MapMemory资源浪费
	// MAPREDUCE-MapperSlowTask:Map慢Task
	// MAPREDUCE-MapperTaskGC:MapperTaskGC
	// MAPREDUCE-MemExceeded:峰值内存超限
	// MAPREDUCE-ReducerDataSkew:Reduce数据倾斜
	// MAPREDUCE-ReducerMemWaste:ReduceMemory资源浪费
	// MAPREDUCE-ReducerSlowTask:Reduce慢Task
	// MAPREDUCE-ReducerTaskGC:ReducerTaskGC
	// MAPREDUCE-SchedulingDelay:调度延迟
	// SPARK-CpuWaste:CPU资源浪费
	// SPARK-DataSkew:数据倾斜
	// SPARK-ExecutorGC:ExecutorGC
	// SPARK-MemExceeded:峰值内存超限
	// SPARK-MemWaste:Memory资源浪费
	// SPARK-ScheduleOverhead:ScheduleOverhead
	// SPARK-ScheduleSkew:调度倾斜
	// SPARK-SlowTask:慢Task
	// TEZ-DataSkew:数据倾斜
	// TEZ-MapperDataSkew:Map数据倾斜
	// TEZ-ReducerDataSkew:Reduce数据倾斜
	// TEZ-TezMemWaste:Memory资源浪费
	// TEZ-TezSlowTask:慢Task
	// TEZ-TezTaskGC:TasksGC
	RuleID *string `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// 洞察规则名字，可参考RuleID的说明
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 洞察规则解释
	RuleExplain *string `json:"RuleExplain,omitnil,omitempty" name:"RuleExplain"`

	// 详情
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 建议信息
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 洞察异常衡量值，同类型的洞察项越大越严重，不同类型的洞察项无对比意义
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`

	// 调度任务执行ID
	ScheduleTaskExecID *string `json:"ScheduleTaskExecID,omitnil,omitempty" name:"ScheduleTaskExecID"`

	// 调度流，DAG
	ScheduleFlowName *string `json:"ScheduleFlowName,omitnil,omitempty" name:"ScheduleFlowName"`

	// 调度flow中的某个task节点
	ScheduleTaskName *string `json:"ScheduleTaskName,omitnil,omitempty" name:"ScheduleTaskName"`

	// Yarn任务的部分核心配置
	JobConf *string `json:"JobConf,omitnil,omitempty" name:"JobConf"`
}

type InstanceChargePrepaid struct {
	// 包年包月时间，默认为1，单位：月。
	// 取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10,11, 12, 24, 36, 48, 60。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 是否自动续费，默认为否。
	// <li>true：是</li>
	// <li>false：否</li>
	RenewFlag *bool `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type Item struct {
	// 健值
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ItemSeq struct {
	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*Item `json:"Items,omitnil,omitempty" name:"Items"`
}

type JobFlowResource struct {
	// 机器类型描述。
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// 机器类型描述，可参考CVM的该含义。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 标签KV对。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 磁盘描述列表。
	DiskGroups []*DiskGroup `json:"DiskGroups,omitnil,omitempty" name:"DiskGroups"`
}

type JobFlowResourceSpec struct {
	// 主节点数量。
	MasterCount *int64 `json:"MasterCount,omitnil,omitempty" name:"MasterCount"`

	// 主节点配置。
	MasterResourceSpec *JobFlowResource `json:"MasterResourceSpec,omitnil,omitempty" name:"MasterResourceSpec"`

	// Core节点数量
	CoreCount *int64 `json:"CoreCount,omitnil,omitempty" name:"CoreCount"`

	// Core节点配置。
	CoreResourceSpec *JobFlowResource `json:"CoreResourceSpec,omitnil,omitempty" name:"CoreResourceSpec"`

	// Task节点数量。
	TaskCount *int64 `json:"TaskCount,omitnil,omitempty" name:"TaskCount"`

	// Common节点数量。
	CommonCount *int64 `json:"CommonCount,omitnil,omitempty" name:"CommonCount"`

	// Task节点配置。
	TaskResourceSpec *JobFlowResource `json:"TaskResourceSpec,omitnil,omitempty" name:"TaskResourceSpec"`

	// Common节点配置。
	CommonResourceSpec *JobFlowResource `json:"CommonResourceSpec,omitnil,omitempty" name:"CommonResourceSpec"`
}

type JobResult struct {
	// 任务步骤名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务步骤失败时的处理策略，可以为以下值：
	// "CONTINUE"，跳过当前失败步骤，继续后续步骤。
	// “TERMINATE_CLUSTER”，终止当前及后续步骤，并销毁集群。
	// “CANCEL_AND_WAIT”，取消当前步骤并阻塞等待处理。
	ActionOnFailure *string `json:"ActionOnFailure,omitnil,omitempty" name:"ActionOnFailure"`

	// 当前步骤的状态，可以为以下值：
	// “JobFlowStepStatusInit”，初始化状态，等待执行。
	// “JobFlowStepStatusRunning”，任务步骤正在执行。
	// “JobFlowStepStatusFailed”，任务步骤执行失败。
	// “JobFlowStepStatusSucceed”，任务步骤执行成功。
	JobState *string `json:"JobState,omitnil,omitempty" name:"JobState"`

	// YARN任务ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
}

type KeyValue struct {
	// 键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type KyuubiQueryInfo struct {
	// 提交IP
	ClientIP *string `json:"ClientIP,omitnil,omitempty" name:"ClientIP"`

	// 执行时长
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Engine Id
	EngineID *string `json:"EngineID,omitnil,omitempty" name:"EngineID"`

	// 计算引擎
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Session Id
	SessionID *string `json:"SessionID,omitnil,omitempty" name:"SessionID"`

	// 开始时间
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 执行状态
	ExecutionState *string `json:"ExecutionState,omitnil,omitempty" name:"ExecutionState"`

	// 执行语句
	ExecutionStatement *string `json:"ExecutionStatement,omitnil,omitempty" name:"ExecutionStatement"`

	// Statement Id
	StatementID *string `json:"StatementID,omitnil,omitempty" name:"StatementID"`

	// 提交用户
	User *string `json:"User,omitnil,omitempty" name:"User"`
}

type LoadAutoScaleStrategy struct {
	// 规则ID。
	StrategyId *int64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 规则名称。
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 规则生效冷却时间。
	CalmDownTime *int64 `json:"CalmDownTime,omitnil,omitempty" name:"CalmDownTime"`

	// 扩缩容动作，1表示扩容，2表示缩容。
	ScaleAction *int64 `json:"ScaleAction,omitnil,omitempty" name:"ScaleAction"`

	// 每次规则生效时的扩缩容数量。
	ScaleNum *int64 `json:"ScaleNum,omitnil,omitempty" name:"ScaleNum"`

	// 指标处理方法，1表示MAX，2表示MIN，3表示AVG。
	ProcessMethod *int64 `json:"ProcessMethod,omitnil,omitempty" name:"ProcessMethod"`

	// 规则优先级，添加时无效，默认为自增。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 规则状态，1表示启动，3表示禁用。
	StrategyStatus *int64 `json:"StrategyStatus,omitnil,omitempty" name:"StrategyStatus"`

	// 规则扩容指定 yarn node label
	YarnNodeLabel *string `json:"YarnNodeLabel,omitnil,omitempty" name:"YarnNodeLabel"`

	// 规则生效的有效时间
	PeriodValid *string `json:"PeriodValid,omitnil,omitempty" name:"PeriodValid"`

	// 优雅缩容开关
	GraceDownFlag *bool `json:"GraceDownFlag,omitnil,omitempty" name:"GraceDownFlag"`

	// 优雅缩容等待时间
	GraceDownTime *int64 `json:"GraceDownTime,omitnil,omitempty" name:"GraceDownTime"`

	// 绑定标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 预设配置组
	ConfigGroupAssigned *string `json:"ConfigGroupAssigned,omitnil,omitempty" name:"ConfigGroupAssigned"`

	// 扩容资源计算方法，"DEFAULT","INSTANCE", "CPU", "MEMORYGB"。
	// "DEFAULT"表示默认方式，与"INSTANCE"意义相同。
	// "INSTANCE"表示按照节点计算，默认方式。
	// "CPU"表示按照机器的核数计算。
	// "MEMORYGB"表示按照机器内存数计算。
	MeasureMethod *string `json:"MeasureMethod,omitnil,omitempty" name:"MeasureMethod"`

	// 节点部署服务列表，例如["HDFS-3.1.2","YARN-3.1.2"]。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SoftDeployDesc []*string `json:"SoftDeployDesc,omitnil,omitempty" name:"SoftDeployDesc"`

	// 启动进程列表，例如["NodeManager"]。
	ServiceNodeDesc *string `json:"ServiceNodeDesc,omitnil,omitempty" name:"ServiceNodeDesc"`

	// 启动进程列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceNodeInfo []*int64 `json:"ServiceNodeInfo,omitnil,omitempty" name:"ServiceNodeInfo"`

	// 节点部署服务列表。部署服务仅填写HDFS、YARN。[组件名对应的映射关系表](https://cloud.tencent.com/document/product/589/98760)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SoftDeployInfo []*int64 `json:"SoftDeployInfo,omitnil,omitempty" name:"SoftDeployInfo"`

	// 多指标触发条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadMetricsConditions *LoadMetricsConditions `json:"LoadMetricsConditions,omitnil,omitempty" name:"LoadMetricsConditions"`

	// 伸缩组Id
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// soft例如yarn
	Soft *string `json:"Soft,omitnil,omitempty" name:"Soft"`
}

type LoadMetricsCondition struct {
	// 规则统计周期，提供1min,3min,5min。
	StatisticPeriod *int64 `json:"StatisticPeriod,omitnil,omitempty" name:"StatisticPeriod"`

	// 触发次数，当连续触发超过TriggerThreshold次后才开始扩缩容。
	TriggerThreshold *int64 `json:"TriggerThreshold,omitnil,omitempty" name:"TriggerThreshold"`

	// 扩缩容负载指标。
	LoadMetrics *string `json:"LoadMetrics,omitnil,omitempty" name:"LoadMetrics"`

	// 规则元数据记录ID。
	MetricId *int64 `json:"MetricId,omitnil,omitempty" name:"MetricId"`

	// 触发条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conditions []*TriggerCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`
}

type LoadMetricsConditions struct {
	// 触发规则条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadMetrics []*LoadMetricsCondition `json:"LoadMetrics,omitnil,omitempty" name:"LoadMetrics"`
}

type LoginSettings struct {
	// 实例登录密码，8-16个字符，包含大写字母、小写字母、数字和特殊字符四种，特殊符号仅支持!@%^*，密码第一位不能为特殊字符
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 密钥ID。关联密钥后，就可以通过对应的私钥来访问实例；PublicKeyId可通过接口[DescribeKeyPairs](https://cloud.tencent.com/document/api/213/15699)获取
	PublicKeyId *string `json:"PublicKeyId,omitnil,omitempty" name:"PublicKeyId"`
}

type MetaDbInfo struct {
	// 元数据类型。
	MetaType *string `json:"MetaType,omitnil,omitempty" name:"MetaType"`

	// 统一元数据库实例ID。
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitnil,omitempty" name:"UnifyMetaInstanceId"`

	// 自建元数据库信息。
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitnil,omitempty" name:"MetaDBInfo"`
}

type MetricTags struct {
	// 指标单位
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 指标Type
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type ModifyAutoRenewFlagRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例ID
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// NOTIFY_AND_MANUAL_RENEW：表示通知即将过期，但不自动续费  NOTIFY_AND_AUTO_RENEW：表示通知即将过期，而且自动续费  DISABLE_NOTIFY_AND_MANUAL_RENEW：表示不通知即将过期，也不自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 计算资源id
	ComputeResourceId *string `json:"ComputeResourceId,omitnil,omitempty" name:"ComputeResourceId"`
}

type ModifyAutoRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例ID
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// NOTIFY_AND_MANUAL_RENEW：表示通知即将过期，但不自动续费  NOTIFY_AND_AUTO_RENEW：表示通知即将过期，而且自动续费  DISABLE_NOTIFY_AND_MANUAL_RENEW：表示不通知即将过期，也不自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 计算资源id
	ComputeResourceId *string `json:"ComputeResourceId,omitnil,omitempty" name:"ComputeResourceId"`
}

func (r *ModifyAutoRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceIds")
	delete(f, "RenewFlag")
	delete(f, "ComputeResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoRenewFlagResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAutoRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoRenewFlagResponseParams `json:"Response"`
}

func (r *ModifyAutoRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoScaleStrategyRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自动扩缩容规则类型，1表示按负载指标扩缩容，2表示按时间扩缩容。
	StrategyType *int64 `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 按负载扩缩容的指标。
	LoadAutoScaleStrategies []*LoadAutoScaleStrategy `json:"LoadAutoScaleStrategies,omitnil,omitempty" name:"LoadAutoScaleStrategies"`

	// 按时间扩缩容的规则。
	TimeAutoScaleStrategies []*TimeAutoScaleStrategy `json:"TimeAutoScaleStrategies,omitnil,omitempty" name:"TimeAutoScaleStrategies"`

	// 伸缩组Id
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type ModifyAutoScaleStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自动扩缩容规则类型，1表示按负载指标扩缩容，2表示按时间扩缩容。
	StrategyType *int64 `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 按负载扩缩容的指标。
	LoadAutoScaleStrategies []*LoadAutoScaleStrategy `json:"LoadAutoScaleStrategies,omitnil,omitempty" name:"LoadAutoScaleStrategies"`

	// 按时间扩缩容的规则。
	TimeAutoScaleStrategies []*TimeAutoScaleStrategy `json:"TimeAutoScaleStrategies,omitnil,omitempty" name:"TimeAutoScaleStrategies"`

	// 伸缩组Id
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *ModifyAutoScaleStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoScaleStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StrategyType")
	delete(f, "LoadAutoScaleStrategies")
	delete(f, "TimeAutoScaleStrategies")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoScaleStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoScaleStrategyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAutoScaleStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoScaleStrategyResponseParams `json:"Response"`
}

func (r *ModifyAutoScaleStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoScaleStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGlobalConfigRequestParams struct {
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改的配置列表。其中Key的取值与`DescribeGlobalConfig`接口的出参一一对应，不区分大小写（如果报错找不到Key，以出参为准），分别为：
	// 1. 开启或关闭资源调度：enableResourceSchedule；在关闭时会有一个同步的选项，Key为sync，取值为true或false。
	// 2. 调度器类型：scheduler。
	// 2. 开启或关闭标签：enableLabel，取值为true或false。
	// 2. 标签目录：labelDir。
	// 3. 是否覆盖用户指定队列：queueMappingOverride，取值为true、false。
	// 4. 程序上限：userMaxAppsDefault。
	// 5. 动态配置项：`DescribeGlobalConfig`接口返回的DefaultSettings中的Name字段。
	// Value的取值都是字符串，对于**是否覆盖用户指定队列**、**程序上限**，json规范中的null表示清空该配置的值。支持修改单个配置项的值。对于**动态配置项**则需要全量传递以进行覆盖。
	Items []*Item `json:"Items,omitnil,omitempty" name:"Items"`
}

type ModifyGlobalConfigRequest struct {
	*tchttp.BaseRequest
	
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改的配置列表。其中Key的取值与`DescribeGlobalConfig`接口的出参一一对应，不区分大小写（如果报错找不到Key，以出参为准），分别为：
	// 1. 开启或关闭资源调度：enableResourceSchedule；在关闭时会有一个同步的选项，Key为sync，取值为true或false。
	// 2. 调度器类型：scheduler。
	// 2. 开启或关闭标签：enableLabel，取值为true或false。
	// 2. 标签目录：labelDir。
	// 3. 是否覆盖用户指定队列：queueMappingOverride，取值为true、false。
	// 4. 程序上限：userMaxAppsDefault。
	// 5. 动态配置项：`DescribeGlobalConfig`接口返回的DefaultSettings中的Name字段。
	// Value的取值都是字符串，对于**是否覆盖用户指定队列**、**程序上限**，json规范中的null表示清空该配置的值。支持修改单个配置项的值。对于**动态配置项**则需要全量传递以进行覆盖。
	Items []*Item `json:"Items,omitnil,omitempty" name:"Items"`
}

func (r *ModifyGlobalConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGlobalConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Items")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGlobalConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGlobalConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGlobalConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGlobalConfigResponseParams `json:"Response"`
}

func (r *ModifyGlobalConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGlobalConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceBasicRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 用来标注修改计算资源
	ResourceBaseType *string `json:"ResourceBaseType,omitnil,omitempty" name:"ResourceBaseType"`

	// 需要修改的计算资源id，与ResourceBaseType 配合使用
	ComputeResourceId *string `json:"ComputeResourceId,omitnil,omitempty" name:"ComputeResourceId"`
}

type ModifyInstanceBasicRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 用来标注修改计算资源
	ResourceBaseType *string `json:"ResourceBaseType,omitnil,omitempty" name:"ResourceBaseType"`

	// 需要修改的计算资源id，与ResourceBaseType 配合使用
	ComputeResourceId *string `json:"ComputeResourceId,omitnil,omitempty" name:"ComputeResourceId"`
}

func (r *ModifyInstanceBasicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceBasicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterName")
	delete(f, "ResourceBaseType")
	delete(f, "ComputeResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceBasicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceBasicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceBasicResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceBasicResponseParams `json:"Response"`
}

func (r *ModifyInstanceBasicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceBasicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPodNumRequestParams struct {
	// 集群Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务编号
	ServiceGroup *int64 `json:"ServiceGroup,omitnil,omitempty" name:"ServiceGroup"`

	// 角色编号
	ServiceType *int64 `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 期望Pod数量
	PodNum *int64 `json:"PodNum,omitnil,omitempty" name:"PodNum"`
}

type ModifyPodNumRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务编号
	ServiceGroup *int64 `json:"ServiceGroup,omitnil,omitempty" name:"ServiceGroup"`

	// 角色编号
	ServiceType *int64 `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 期望Pod数量
	PodNum *int64 `json:"PodNum,omitnil,omitempty" name:"PodNum"`
}

func (r *ModifyPodNumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPodNumRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ServiceGroup")
	delete(f, "ServiceType")
	delete(f, "PodNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPodNumRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPodNumResponseParams struct {
	// 集群Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 流程Id
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPodNumResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPodNumResponseParams `json:"Response"`
}

func (r *ModifyPodNumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPodNumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePoolsRequestParams struct {
	// emr集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 取值范围：
	// <li>fair:代表公平调度标识</li>
	// <li>capacity:代表容量调度标识</li>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`
}

type ModifyResourcePoolsRequest struct {
	*tchttp.BaseRequest
	
	// emr集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 取值范围：
	// <li>fair:代表公平调度标识</li>
	// <li>capacity:代表容量调度标识</li>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`
}

func (r *ModifyResourcePoolsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePoolsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Key")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourcePoolsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePoolsResponseParams struct {
	// false表示不是草稿，提交刷新请求成功
	IsDraft *bool `json:"IsDraft,omitnil,omitempty" name:"IsDraft"`

	// 扩展字段，暂时没用
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourcePoolsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourcePoolsResponseParams `json:"Response"`
}

func (r *ModifyResourcePoolsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePoolsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 计费类型
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 变配CPU
	NewCpu *int64 `json:"NewCpu,omitnil,omitempty" name:"NewCpu"`

	// 变配内存
	NewMem *int64 `json:"NewMem,omitnil,omitempty" name:"NewMem"`

	// Token
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 变配机器规格
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 节点ID列表
	ResourceIdList []*string `json:"ResourceIdList,omitnil,omitempty" name:"ResourceIdList"`
}

type ModifyResourceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 计费类型
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 变配CPU
	NewCpu *int64 `json:"NewCpu,omitnil,omitempty" name:"NewCpu"`

	// 变配内存
	NewMem *int64 `json:"NewMem,omitnil,omitempty" name:"NewMem"`

	// Token
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 变配机器规格
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 节点ID列表
	ResourceIdList []*string `json:"ResourceIdList,omitnil,omitempty" name:"ResourceIdList"`
}

func (r *ModifyResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PayMode")
	delete(f, "NewCpu")
	delete(f, "NewMem")
	delete(f, "ClientToken")
	delete(f, "InstanceType")
	delete(f, "ResourceIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceResponseParams struct {
	// 流程traceId
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceResponseParams `json:"Response"`
}

func (r *ModifyResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceScheduleConfigRequestParams struct {
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 业务标识，fair表示编辑公平的配置项，fairPlan表示编辑执行计划，capacity表示编辑容量的配置项
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 修改后的模块消息
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ModifyResourceScheduleConfigRequest struct {
	*tchttp.BaseRequest
	
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 业务标识，fair表示编辑公平的配置项，fairPlan表示编辑执行计划，capacity表示编辑容量的配置项
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 修改后的模块消息
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

func (r *ModifyResourceScheduleConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceScheduleConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Key")
	delete(f, "Value")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceScheduleConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceScheduleConfigResponseParams struct {
	// true为草稿，表示还没有刷新资源池
	IsDraft *bool `json:"IsDraft,omitnil,omitempty" name:"IsDraft"`

	// 校验错误信息，如果不为空，则说明校验失败，配置没有成功
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 返回数据
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourceScheduleConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceScheduleConfigResponseParams `json:"Response"`
}

func (r *ModifyResourceScheduleConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceScheduleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceSchedulerRequestParams struct {
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 老的调度器:fair
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`

	// 新的调度器:capacity
	NewValue *string `json:"NewValue,omitnil,omitempty" name:"NewValue"`
}

type ModifyResourceSchedulerRequest struct {
	*tchttp.BaseRequest
	
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 老的调度器:fair
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`

	// 新的调度器:capacity
	NewValue *string `json:"NewValue,omitnil,omitempty" name:"NewValue"`
}

func (r *ModifyResourceSchedulerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceSchedulerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OldValue")
	delete(f, "NewValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceSchedulerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceSchedulerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourceSchedulerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceSchedulerResponseParams `json:"Response"`
}

func (r *ModifyResourceSchedulerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceSchedulerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceTags struct {
	// 集群id 或者 cvm id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源6段式表达式
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 资源前缀
	ResourcePrefix *string `json:"ResourcePrefix,omitnil,omitempty" name:"ResourcePrefix"`

	// ap-beijing
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`

	// emr
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 删除的标签列表
	DeleteTags []*Tag `json:"DeleteTags,omitnil,omitempty" name:"DeleteTags"`

	// 添加的标签列表
	AddTags []*Tag `json:"AddTags,omitnil,omitempty" name:"AddTags"`

	// 修改的标签列表
	ModifyTags []*Tag `json:"ModifyTags,omitnil,omitempty" name:"ModifyTags"`
}

// Predefined struct for user
type ModifyResourcesTagsRequestParams struct {
	// 标签类型，取值Cluster或者Node
	ModifyType *string `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`

	// 标签信息
	ModifyResourceTagsInfoList []*ModifyResourceTags `json:"ModifyResourceTagsInfoList,omitnil,omitempty" name:"ModifyResourceTagsInfoList"`
}

type ModifyResourcesTagsRequest struct {
	*tchttp.BaseRequest
	
	// 标签类型，取值Cluster或者Node
	ModifyType *string `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`

	// 标签信息
	ModifyResourceTagsInfoList []*ModifyResourceTags `json:"ModifyResourceTagsInfoList,omitnil,omitempty" name:"ModifyResourceTagsInfoList"`
}

func (r *ModifyResourcesTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcesTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModifyType")
	delete(f, "ModifyResourceTagsInfoList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourcesTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcesTagsResponseParams struct {
	// 成功的资源id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessList []*string `json:"SuccessList,omitnil,omitempty" name:"SuccessList"`

	// 失败的资源id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailList []*string `json:"FailList,omitnil,omitempty" name:"FailList"`

	// 部分成功的资源id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartSuccessList []*string `json:"PartSuccessList,omitnil,omitempty" name:"PartSuccessList"`

	// 集群id与流程id的映射列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterToFlowIdList []*ClusterIDToFlowID `json:"ClusterToFlowIdList,omitnil,omitempty" name:"ClusterToFlowIdList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourcesTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourcesTagsResponseParams `json:"Response"`
}

func (r *ModifyResourcesTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcesTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySLInstanceBasicRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`
}

type ModifySLInstanceBasicRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`
}

func (r *ModifySLInstanceBasicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySLInstanceBasicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySLInstanceBasicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySLInstanceBasicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySLInstanceBasicResponse struct {
	*tchttp.BaseResponse
	Response *ModifySLInstanceBasicResponseParams `json:"Response"`
}

func (r *ModifySLInstanceBasicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySLInstanceBasicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySLInstanceRequestParams struct {
	// 实例唯一标识符（字符串表示）。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要变更的区域名称。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 该区域变配后的目标节点数量，所有区域节点总数应大于等于3，小于等于50。
	NodeNum *int64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`
}

type ModifySLInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例唯一标识符（字符串表示）。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要变更的区域名称。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 该区域变配后的目标节点数量，所有区域节点总数应大于等于3，小于等于50。
	NodeNum *int64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`
}

func (r *ModifySLInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySLInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Zone")
	delete(f, "NodeNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySLInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySLInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySLInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifySLInstanceResponseParams `json:"Response"`
}

func (r *ModifySLInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySLInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserManagerPwdRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 密码
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`
}

type ModifyUserManagerPwdRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 密码
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`
}

func (r *ModifyUserManagerPwdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserManagerPwdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "PassWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserManagerPwdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserManagerPwdResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserManagerPwdResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserManagerPwdResponseParams `json:"Response"`
}

func (r *ModifyUserManagerPwdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserManagerPwdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyYarnDeployRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 切换后的调度器，可选值为fair、capacity
	NewScheduler *string `json:"NewScheduler,omitnil,omitempty" name:"NewScheduler"`

	// 现在使用的调度器，可选值为fair、capacity
	OldScheduler *string `json:"OldScheduler,omitnil,omitempty" name:"OldScheduler"`
}

type ModifyYarnDeployRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 切换后的调度器，可选值为fair、capacity
	NewScheduler *string `json:"NewScheduler,omitnil,omitempty" name:"NewScheduler"`

	// 现在使用的调度器，可选值为fair、capacity
	OldScheduler *string `json:"OldScheduler,omitnil,omitempty" name:"OldScheduler"`
}

func (r *ModifyYarnDeployRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyYarnDeployRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NewScheduler")
	delete(f, "OldScheduler")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyYarnDeployRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyYarnDeployResponseParams struct {
	// 为false不点亮部署生效、重置
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDraft *bool `json:"IsDraft,omitnil,omitempty" name:"IsDraft"`

	// 错误信息，预留
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyYarnDeployResponse struct {
	*tchttp.BaseResponse
	Response *ModifyYarnDeployResponseParams `json:"Response"`
}

func (r *ModifyYarnDeployResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyYarnDeployResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyYarnQueueV2RequestParams struct {
	// 集群Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 调度器类型。可选值：
	// 
	// 1. capacity
	// 2. fair
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// 资源池数据
	ConfigModifyInfoList []*ConfigModifyInfoV2 `json:"ConfigModifyInfoList,omitnil,omitempty" name:"ConfigModifyInfoList"`
}

type ModifyYarnQueueV2Request struct {
	*tchttp.BaseRequest
	
	// 集群Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 调度器类型。可选值：
	// 
	// 1. capacity
	// 2. fair
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// 资源池数据
	ConfigModifyInfoList []*ConfigModifyInfoV2 `json:"ConfigModifyInfoList,omitnil,omitempty" name:"ConfigModifyInfoList"`
}

func (r *ModifyYarnQueueV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyYarnQueueV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Scheduler")
	delete(f, "ConfigModifyInfoList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyYarnQueueV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyYarnQueueV2ResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyYarnQueueV2Response struct {
	*tchttp.BaseResponse
	Response *ModifyYarnQueueV2ResponseParams `json:"Response"`
}

func (r *ModifyYarnQueueV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyYarnQueueV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonthRepeatStrategy struct {
	// 重复任务执行的具体时刻，例如"01:02:00"
	ExecuteAtTimeOfDay *string `json:"ExecuteAtTimeOfDay,omitnil,omitempty" name:"ExecuteAtTimeOfDay"`

	// 每月中的天数时间段描述，长度只能为2，例如[2,10]表示每月2-10号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DaysOfMonthRange []*uint64 `json:"DaysOfMonthRange,omitnil,omitempty" name:"DaysOfMonthRange"`
}

type MultiDisk struct {
	// 云盘类型
	// <li>CLOUD_SSD：表示云SSD。</li>
	// <li>CLOUD_PREMIUM：表示高效云盘。</li>
	// <li>CLOUD_HSSD：表示增强型SSD云硬盘。</li>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 云盘大小
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 该类型云盘个数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type MultiDiskMC struct {
	// 该类型云盘个数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 磁盘类型
	// 1  :本地盘
	// 2  :云硬盘
	// 3  : 本地SSD
	// 4  : 云SSD
	// 5  : 高效云盘
	// 6  : 增强型SSD云硬盘
	// 11 : 吞吐型云硬盘
	// 12 : 极速型SSD云硬盘
	// 13 : 通用型SSD云硬盘
	// 14 : 大数据型云硬盘
	// 15 : 高IO型云硬盘
	// 16 : 远端SSD盘
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 磁盘大小
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// 云盘大小,单位b
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`
}

type MultiZoneSetting struct {
	// "master"、"standby"、"third-party"
	ZoneTag *string `json:"ZoneTag,omitnil,omitempty" name:"ZoneTag"`

	// 无
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil,omitempty" name:"VPCSettings"`

	// 无
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 无
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitnil,omitempty" name:"ResourceSpec"`
}

type NewResourceSpec struct {
	// 描述Master节点资源
	MasterResourceSpec *Resource `json:"MasterResourceSpec,omitnil,omitempty" name:"MasterResourceSpec"`

	// 描述Core节点资源
	CoreResourceSpec *Resource `json:"CoreResourceSpec,omitnil,omitempty" name:"CoreResourceSpec"`

	// 描述Task节点资源
	TaskResourceSpec *Resource `json:"TaskResourceSpec,omitnil,omitempty" name:"TaskResourceSpec"`

	// Master节点数量
	MasterCount *int64 `json:"MasterCount,omitnil,omitempty" name:"MasterCount"`

	// Core节点数量
	CoreCount *int64 `json:"CoreCount,omitnil,omitempty" name:"CoreCount"`

	// Task节点数量
	TaskCount *int64 `json:"TaskCount,omitnil,omitempty" name:"TaskCount"`

	// 描述Common节点资源
	CommonResourceSpec *Resource `json:"CommonResourceSpec,omitnil,omitempty" name:"CommonResourceSpec"`

	// Common节点数量
	CommonCount *int64 `json:"CommonCount,omitnil,omitempty" name:"CommonCount"`
}

type NodeAffinity struct {
	// 节点亲和性-强制调度设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequiredDuringSchedulingIgnoredDuringExecution *NodeSelector `json:"RequiredDuringSchedulingIgnoredDuringExecution,omitnil,omitempty" name:"RequiredDuringSchedulingIgnoredDuringExecution"`

	// 节点亲和性-容忍调度
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredSchedulingTerm `json:"PreferredDuringSchedulingIgnoredDuringExecution,omitnil,omitempty" name:"PreferredDuringSchedulingIgnoredDuringExecution"`
}

type NodeDetailPriceResult struct {
	// 节点类型 master core task common router mysql
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点组成部分价格详情
	PartDetailPrice []*PartDetailPriceItem `json:"PartDetailPrice,omitnil,omitempty" name:"PartDetailPrice"`
}

type NodeHardwareInfo struct {
	// 用户APPID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 序列号
	SerialNo *string `json:"SerialNo,omitnil,omitempty" name:"SerialNo"`

	// 机器实例ID
	OrderNo *string `json:"OrderNo,omitnil,omitempty" name:"OrderNo"`

	// master节点绑定外网IP
	WanIp *string `json:"WanIp,omitnil,omitempty" name:"WanIp"`

	// 节点类型。0:common节点；1:master节点
	// ；2:core节点；3:task节点
	Flag *int64 `json:"Flag,omitnil,omitempty" name:"Flag"`

	// 节点规格
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// 节点核数
	CpuNum *int64 `json:"CpuNum,omitnil,omitempty" name:"CpuNum"`

	// 节点内存,单位b
	MemSize *int64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// 节点内存描述，单位GB
	MemDesc *string `json:"MemDesc,omitnil,omitempty" name:"MemDesc"`

	// 节点所在region
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 节点所在Zone
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 申请时间
	ApplyTime *string `json:"ApplyTime,omitnil,omitempty" name:"ApplyTime"`

	// 释放时间
	FreeTime *string `json:"FreeTime,omitnil,omitempty" name:"FreeTime"`

	// 硬盘大小
	DiskSize *string `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 节点描述
	NameTag *string `json:"NameTag,omitnil,omitempty" name:"NameTag"`

	// 节点部署服务
	Services *string `json:"Services,omitnil,omitempty" name:"Services"`

	// 磁盘类型，1 :本地盘 2 :云硬盘 3 : 本地SSD 4 : 云SSD 5 : 高效云盘 6 : 增强型SSD云硬盘 11 : 吞吐型云硬盘 12 : 极速型SSD云硬盘 13 : 通用型SSD云硬盘 14 : 大数据型云硬盘 15 : 高IO型云硬盘 16 : 远端SSD盘
	StorageType *int64 `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 系统盘大小，单位GB
	RootSize *int64 `json:"RootSize,omitnil,omitempty" name:"RootSize"`

	// 付费类型，0：按量计费；1：包年包月
	ChargeType *int64 `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 数据库IP
	CdbIp *string `json:"CdbIp,omitnil,omitempty" name:"CdbIp"`

	// 数据库端口
	CdbPort *int64 `json:"CdbPort,omitnil,omitempty" name:"CdbPort"`

	// 硬盘容量,单位b
	HwDiskSize *int64 `json:"HwDiskSize,omitnil,omitempty" name:"HwDiskSize"`

	// 硬盘容量描述
	HwDiskSizeDesc *string `json:"HwDiskSizeDesc,omitnil,omitempty" name:"HwDiskSizeDesc"`

	// 内存容量，单位b
	HwMemSize *int64 `json:"HwMemSize,omitnil,omitempty" name:"HwMemSize"`

	// 内存容量描述
	HwMemSizeDesc *string `json:"HwMemSizeDesc,omitnil,omitempty" name:"HwMemSizeDesc"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 节点资源ID
	EmrResourceId *string `json:"EmrResourceId,omitnil,omitempty" name:"EmrResourceId"`

	// 续费标志
	IsAutoRenew *int64 `json:"IsAutoRenew,omitnil,omitempty" name:"IsAutoRenew"`

	// 设备标识
	DeviceClass *string `json:"DeviceClass,omitnil,omitempty" name:"DeviceClass"`

	// 支持变配
	Mutable *int64 `json:"Mutable,omitnil,omitempty" name:"Mutable"`

	// 多云盘
	// 注意：此字段可能返回 null，表示取不到有效值。
	MCMultiDisk []*MultiDiskMC `json:"MCMultiDisk,omitnil,omitempty" name:"MCMultiDisk"`

	// 数据库信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdbNodeInfo *CdbInfo `json:"CdbNodeInfo,omitnil,omitempty" name:"CdbNodeInfo"`

	// 内网IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 此节点是否可销毁，1可销毁，0不可销毁
	Destroyable *int64 `json:"Destroyable,omitnil,omitempty" name:"Destroyable"`

	// 节点绑定的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否是自动扩缩容节点，0为普通节点，1为自动扩缩容节点。
	AutoFlag *int64 `json:"AutoFlag,omitnil,omitempty" name:"AutoFlag"`

	// 资源类型, host/pod
	HardwareResourceType *string `json:"HardwareResourceType,omitnil,omitempty" name:"HardwareResourceType"`

	// 是否浮动规格，1是，0否
	IsDynamicSpec *int64 `json:"IsDynamicSpec,omitnil,omitempty" name:"IsDynamicSpec"`

	// 浮动规格值json字符串
	DynamicPodSpec *string `json:"DynamicPodSpec,omitnil,omitempty" name:"DynamicPodSpec"`

	// 是否支持变更计费类型 1是，0否
	SupportModifyPayMode *int64 `json:"SupportModifyPayMode,omitnil,omitempty" name:"SupportModifyPayMode"`

	// 系统盘类型，1 :本地盘 2 :云硬盘 3 : 本地SSD 4 : 云SSD 5 : 高效云盘 6 : 增强型SSD云硬盘 11 : 吞吐型云硬盘 12 : 极速型SSD云硬盘 13 : 通用型SSD云硬盘 14 : 大数据型云硬盘 15 : 高IO型云硬盘 16 : 远端SSD盘
	RootStorageType *int64 `json:"RootStorageType,omitnil,omitempty" name:"RootStorageType"`

	// 可用区信息
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 子网
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetInfo *SubnetInfo `json:"SubnetInfo,omitnil,omitempty" name:"SubnetInfo"`

	// 客户端
	Clients *string `json:"Clients,omitnil,omitempty" name:"Clients"`

	// 系统当前时间
	CurrentTime *string `json:"CurrentTime,omitnil,omitempty" name:"CurrentTime"`

	// 是否用于联邦 ,1是，0否
	IsFederation *int64 `json:"IsFederation,omitnil,omitempty" name:"IsFederation"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 服务
	ServiceClient *string `json:"ServiceClient,omitnil,omitempty" name:"ServiceClient"`

	// 该实例是否开启实例保护，true为开启 false为关闭
	DisableApiTermination *bool `json:"DisableApiTermination,omitnil,omitempty" name:"DisableApiTermination"`

	// 0表示老计费，1表示新计费
	TradeVersion *int64 `json:"TradeVersion,omitnil,omitempty" name:"TradeVersion"`

	// 各组件状态，Zookeeper:STARTED,ResourceManager:STARTED，STARTED已启动，STOPED已停止
	ServicesStatus *string `json:"ServicesStatus,omitnil,omitempty" name:"ServicesStatus"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 共享集群id
	SharedClusterId *string `json:"SharedClusterId,omitnil,omitempty" name:"SharedClusterId"`

	// 共享集群id描述
	SharedClusterIdDesc *string `json:"SharedClusterIdDesc,omitnil,omitempty" name:"SharedClusterIdDesc"`

	// 是否是定时销毁资源
	TimingResource *bool `json:"TimingResource,omitnil,omitempty" name:"TimingResource"`

	// 资源类型（HardwareResourceType）为pod时，对应的TKE集群id
	TkeClusterId *string `json:"TkeClusterId,omitnil,omitempty" name:"TkeClusterId"`
}

type NodeResource struct {
	// 配置Id
	ResourceConfigId *uint64 `json:"ResourceConfigId,omitnil,omitempty" name:"ResourceConfigId"`

	// Resource
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *Resource `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 是否默认配置,DEFAULT,BACKUP
	IsDefault *string `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 该类型剩余
	MaxResourceNum *uint64 `json:"MaxResourceNum,omitnil,omitempty" name:"MaxResourceNum"`

	// 支持的包销时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrepaidUnderwritePeriods []*int64 `json:"PrepaidUnderwritePeriods,omitnil,omitempty" name:"PrepaidUnderwritePeriods"`
}

type NodeResourceSpec struct {
	// 规格类型，如S2.MEDIUM8
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 系统盘，系统盘个数不超过1块
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemDisk []*DiskSpecInfo `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 需要绑定的标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 云数据盘，云数据盘总个数不超过15块
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataDisk []*DiskSpecInfo `json:"DataDisk,omitnil,omitempty" name:"DataDisk"`

	// 本地数据盘
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalDataDisk []*DiskSpecInfo `json:"LocalDataDisk,omitnil,omitempty" name:"LocalDataDisk"`
}

type NodeSelector struct {
	// Pod强制调度节点选择条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeSelectorTerms []*NodeSelectorTerm `json:"NodeSelectorTerms,omitnil,omitempty" name:"NodeSelectorTerms"`
}

type NodeSelectorRequirement struct {
	// 节点选择项Key值
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 节点选择项Operator值，支持In, NotIn, Exists, DoesNotExist. Gt, and Lt.
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 节点选择项Values值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type NodeSelectorTerm struct {
	// 节点选择项表达式集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchExpressions []*NodeSelectorRequirement `json:"MatchExpressions,omitnil,omitempty" name:"MatchExpressions"`
}

type NodeSpecDiskV2 struct {
	// 数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 磁盘类型
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 指定磁盘大小
	DefaultDiskSize *int64 `json:"DefaultDiskSize,omitnil,omitempty" name:"DefaultDiskSize"`
}

type NotRepeatStrategy struct {
	// 该次任务执行的具体完整时间，格式为"2020-07-13 00:00:00"
	ExecuteAt *string `json:"ExecuteAt,omitnil,omitempty" name:"ExecuteAt"`
}

type OpScope struct {
	// 操作范围，要操作的服务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceInfoList []*ServiceBasicRestartInfo `json:"ServiceInfoList,omitnil,omitempty" name:"ServiceInfoList"`
}

type OutterResource struct {
	// 规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// 规格名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 硬盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *int64 `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 硬盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 系统盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootSize *int64 `json:"RootSize,omitnil,omitempty" name:"RootSize"`

	// 内存大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemSize *int64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// CPU个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 硬盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type OverviewMetricData struct {
	// 指标名
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 第一个数据时间戳
	First *int64 `json:"First,omitnil,omitempty" name:"First"`

	// 最后一个数据时间戳
	Last *int64 `json:"Last,omitnil,omitempty" name:"Last"`

	// 采样点时间间隔
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 采样点数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataPoints []*string `json:"DataPoints,omitnil,omitempty" name:"DataPoints"`

	// 指标tags
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags *MetricTags `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type OverviewRow struct {
	// 表名字
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 读请求次数
	ReadRequestCount *float64 `json:"ReadRequestCount,omitnil,omitempty" name:"ReadRequestCount"`

	// 写请求次数
	WriteRequestCount *float64 `json:"WriteRequestCount,omitnil,omitempty" name:"WriteRequestCount"`

	// 当前memstore的size
	MemstoreSize *float64 `json:"MemstoreSize,omitnil,omitempty" name:"MemstoreSize"`

	// 当前region中StroreFile的size
	StoreFileSize *float64 `json:"StoreFileSize,omitnil,omitempty" name:"StoreFileSize"`

	// regions，点击可跳转
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`
}

type PartDetailPriceItem struct {
	// 类型包括：节点->node、系统盘->rootDisk、云数据盘->dataDisk、metaDB
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 单价（原价）
	Price *float64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 单价（折扣价）
	RealCost *float64 `json:"RealCost,omitnil,omitempty" name:"RealCost"`

	// 总价（折扣价）
	RealTotalCost *float64 `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 折扣
	Policy *float64 `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 数量
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`
}

type Period struct {
	// 时间跨度
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 时间单位，"m"代表月。
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
}

type PersistentVolumeContext struct {
	// 磁盘大小，单位为GB。
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 磁盘类型。CLOUD_PREMIUM;CLOUD_SSD
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 磁盘数量
	DiskNum *int64 `json:"DiskNum,omitnil,omitempty" name:"DiskNum"`

	// 云盘额外性能
	ExtraPerformance *int64 `json:"ExtraPerformance,omitnil,omitempty" name:"ExtraPerformance"`
}

type Placement struct {
	// 实例所属的可用区，例如ap-guangzhou-1。该参数也可以通过调用[DescribeZones](https://cloud.tencent.com/document/product/213/15707) 的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例所属项目ID。该参数可以通过调用[DescribeProject](https://cloud.tencent.com/document/api/651/78725) 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type PodNewParameter struct {
	// TKE或EKS集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自定义权限
	// 如：
	// {
	//   "apiVersion": "v1",
	//   "clusters": [
	//     {
	//       "cluster": {
	//         "certificate-authority-data": "xxxxxx==",
	//         "server": "https://xxxxx.com"
	//       },
	//       "name": "cls-xxxxx"
	//     }
	//   ],
	//   "contexts": [
	//     {
	//       "context": {
	//         "cluster": "cls-xxxxx",
	//         "user": "100014xxxxx"
	//       },
	//       "name": "cls-a44yhcxxxxxxxxxx"
	//     }
	//   ],
	//   "current-context": "cls-a4xxxx-context-default",
	//   "kind": "Config",
	//   "preferences": {},
	//   "users": [
	//     {
	//       "name": "100014xxxxx",
	//       "user": {
	//         "client-certificate-data": "xxxxxx",
	//         "client-key-data": "xxxxxx"
	//       }
	//     }
	//   ]
	// }
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 自定义参数
	// 如：
	// {
	//     "apiVersion": "apps/v1",
	//     "kind": "Deployment",
	//     "metadata": {
	//       "name": "test-deployment",
	//       "labels": {
	//         "app": "test"
	//       }
	//     },
	//     "spec": {
	//       "replicas": 3,
	//       "selector": {
	//         "matchLabels": {
	//           "app": "test-app"
	//         }
	//       },
	//       "template": {
	//         "metadata": {
	//           "annotations": {
	//             "your-organization.com/department-v1": "test-example-v1",
	//             "your-organization.com/department-v2": "test-example-v2"
	//           },
	//           "labels": {
	//             "app": "test-app",
	//             "environment": "production"
	//           }
	//         },
	//         "spec": {
	//           "nodeSelector": {
	//             "your-organization/node-test": "test-node"
	//           },
	//           "containers": [
	//             {
	//               "name": "nginx",
	//               "image": "nginx:1.14.2",
	//               "ports": [
	//                 {
	//                   "containerPort": 80
	//                 }
	//               ]
	//             }
	//           ],
	//           "affinity": {
	//             "nodeAffinity": {
	//               "requiredDuringSchedulingIgnoredDuringExecution": {
	//                 "nodeSelectorTerms": [
	//                   {
	//                     "matchExpressions": [
	//                       {
	//                         "key": "disk-type",
	//                         "operator": "In",
	//                         "values": [
	//                           "ssd",
	//                           "sas"
	//                         ]
	//                       },
	//                       {
	//                         "key": "cpu-num",
	//                         "operator": "Gt",
	//                         "values": [
	//                           "6"
	//                         ]
	//                       }
	//                     ]
	//                   }
	//                 ]
	//               }
	//             }
	//           }
	//         }
	//       }
	//     }
	//   }
	Parameter *string `json:"Parameter,omitnil,omitempty" name:"Parameter"`
}

type PodNewSpec struct {
	// 外部资源提供者的标识符，例如"cls-a1cd23fa"。
	ResourceProviderIdentifier *string `json:"ResourceProviderIdentifier,omitnil,omitempty" name:"ResourceProviderIdentifier"`

	// 外部资源提供者类型，例如"tke",当前仅支持"tke"。
	ResourceProviderType *string `json:"ResourceProviderType,omitnil,omitempty" name:"ResourceProviderType"`

	// 资源的用途，即节点类型，当前仅支持"TASK"。
	NodeFlag *string `json:"NodeFlag,omitnil,omitempty" name:"NodeFlag"`

	// CPU核数。
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存大小，单位为GB。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Eks集群-CPU类型，当前支持"intel"和"amd"
	CpuType *string `json:"CpuType,omitnil,omitempty" name:"CpuType"`

	// Pod节点数据目录挂载信息。
	PodVolumes []*PodVolume `json:"PodVolumes,omitnil,omitempty" name:"PodVolumes"`

	// 是否浮动规格，默认否
	// <li>true：代表是</li>
	// <li>false：代表否</li>
	EnableDynamicSpecFlag *bool `json:"EnableDynamicSpecFlag,omitnil,omitempty" name:"EnableDynamicSpecFlag"`

	// 浮动规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	DynamicPodSpec *DynamicPodSpec `json:"DynamicPodSpec,omitnil,omitempty" name:"DynamicPodSpec"`

	// 代表vpc网络唯一id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 代表vpc子网唯一id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// pod name
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`
}

type PodParameter struct {
	// TKE或EKS集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 自定义权限
	// 如：
	// {
	//   "apiVersion": "v1",
	//   "clusters": [
	//     {
	//       "cluster": {
	//         "certificate-authority-data": "xxxxxx==",
	//         "server": "https://xxxxx.com"
	//       },
	//       "name": "cls-xxxxx"
	//     }
	//   ],
	//   "contexts": [
	//     {
	//       "context": {
	//         "cluster": "cls-xxxxx",
	//         "user": "100014xxxxx"
	//       },
	//       "name": "cls-a44yhcxxxxxxxxxx"
	//     }
	//   ],
	//   "current-context": "cls-a4xxxx-context-default",
	//   "kind": "Config",
	//   "preferences": {},
	//   "users": [
	//     {
	//       "name": "100014xxxxx",
	//       "user": {
	//         "client-certificate-data": "xxxxxx",
	//         "client-key-data": "xxxxxx"
	//       }
	//     }
	//   ]
	// }
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 自定义参数
	// 如：
	// {
	//     "apiVersion": "apps/v1",
	//     "kind": "Deployment",
	//     "metadata": {
	//       "name": "test-deployment",
	//       "labels": {
	//         "app": "test"
	//       }
	//     },
	//     "spec": {
	//       "replicas": 3,
	//       "selector": {
	//         "matchLabels": {
	//           "app": "test-app"
	//         }
	//       },
	//       "template": {
	//         "metadata": {
	//           "annotations": {
	//             "your-organization.com/department-v1": "test-example-v1",
	//             "your-organization.com/department-v2": "test-example-v2"
	//           },
	//           "labels": {
	//             "app": "test-app",
	//             "environment": "production"
	//           }
	//         },
	//         "spec": {
	//           "nodeSelector": {
	//             "your-organization/node-test": "test-node"
	//           },
	//           "containers": [
	//             {
	//               "name": "nginx",
	//               "image": "nginx:1.14.2",
	//               "ports": [
	//                 {
	//                   "containerPort": 80
	//                 }
	//               ]
	//             }
	//           ],
	//           "affinity": {
	//             "nodeAffinity": {
	//               "requiredDuringSchedulingIgnoredDuringExecution": {
	//                 "nodeSelectorTerms": [
	//                   {
	//                     "matchExpressions": [
	//                       {
	//                         "key": "disk-type",
	//                         "operator": "In",
	//                         "values": [
	//                           "ssd",
	//                           "sas"
	//                         ]
	//                       },
	//                       {
	//                         "key": "cpu-num",
	//                         "operator": "Gt",
	//                         "values": [
	//                           "6"
	//                         ]
	//                       }
	//                     ]
	//                   }
	//                 ]
	//               }
	//             }
	//           }
	//         }
	//       }
	//     }
	//   }
	Parameter *string `json:"Parameter,omitnil,omitempty" name:"Parameter"`
}

type PodSaleSpec struct {
	// 可售卖的资源规格，仅为以下值:"TASK","CORE","MASTER","ROUTER"。
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// Cpu核数。
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存数量，单位为GB。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 该规格资源可申请的最大数量。
	Number *uint64 `json:"Number,omitnil,omitempty" name:"Number"`
}

type PodSpec struct {
	// 外部资源提供者的标识符，例如"cls-a1cd23fa"。
	ResourceProviderIdentifier *string `json:"ResourceProviderIdentifier,omitnil,omitempty" name:"ResourceProviderIdentifier"`

	// 外部资源提供者类型，例如"tke",当前仅支持"tke"。
	ResourceProviderType *string `json:"ResourceProviderType,omitnil,omitempty" name:"ResourceProviderType"`

	// 资源的用途，即节点类型，当前仅支持"TASK"。
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// CPU核数。
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存大小，单位为GB。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 资源对宿主机的挂载点，指定的挂载点对应了宿主机的路径，该挂载点在Pod中作为数据存储目录使用。弃用
	DataVolumes []*string `json:"DataVolumes,omitnil,omitempty" name:"DataVolumes"`

	// Eks集群-CPU类型，当前支持"intel"和"amd"
	CpuType *string `json:"CpuType,omitnil,omitempty" name:"CpuType"`

	// Pod节点数据目录挂载信息。
	PodVolumes []*PodVolume `json:"PodVolumes,omitnil,omitempty" name:"PodVolumes"`

	// 是否浮动规格，1是，0否
	IsDynamicSpec *uint64 `json:"IsDynamicSpec,omitnil,omitempty" name:"IsDynamicSpec"`

	// 浮动规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	DynamicPodSpec *DynamicPodSpec `json:"DynamicPodSpec,omitnil,omitempty" name:"DynamicPodSpec"`

	// 代表vpc网络唯一id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 代表vpc子网唯一id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// pod name
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`
}

type PodSpecInfo struct {
	// 使用Pod资源扩容时，指定的Pod规格以及来源等信息
	PodSpec *PodNewSpec `json:"PodSpec,omitnil,omitempty" name:"PodSpec"`

	// POD自定义权限和自定义参数
	PodParameter *PodNewParameter `json:"PodParameter,omitnil,omitempty" name:"PodParameter"`
}

type PodState struct {
	// pod的名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// pod uuid
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// pod的状态
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// pod处于该状态原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// pod所属集群
	OwnerCluster *string `json:"OwnerCluster,omitnil,omitempty" name:"OwnerCluster"`

	// pod内存大小
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`
}

type PodVolume struct {
	// 存储类型，可为"pvc"，"hostpath"。
	VolumeType *string `json:"VolumeType,omitnil,omitempty" name:"VolumeType"`

	// 当VolumeType为"pvc"时，该字段生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PVCVolume *PersistentVolumeContext `json:"PVCVolume,omitnil,omitempty" name:"PVCVolume"`

	// 当VolumeType为"hostpath"时，该字段生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostVolume *HostVolumeContext `json:"HostVolume,omitnil,omitempty" name:"HostVolume"`
}

type PreExecuteFileSettings struct {
	// 脚本在COS上路径，已废弃
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 执行脚本参数
	Args []*string `json:"Args,omitnil,omitempty" name:"Args"`

	// COS的Bucket名称，已废弃
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// COS的Region名称，已废弃
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// COS的Domain数据，已废弃
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 执行顺序
	RunOrder *int64 `json:"RunOrder,omitnil,omitempty" name:"RunOrder"`

	// resourceAfter 或 clusterAfter
	WhenRun *string `json:"WhenRun,omitnil,omitempty" name:"WhenRun"`

	// 脚本文件名，已废弃
	CosFileName *string `json:"CosFileName,omitnil,omitempty" name:"CosFileName"`

	// 脚本的cos地址
	CosFileURI *string `json:"CosFileURI,omitnil,omitempty" name:"CosFileURI"`

	// cos的SecretId
	CosSecretId *string `json:"CosSecretId,omitnil,omitempty" name:"CosSecretId"`

	// Cos的SecretKey
	CosSecretKey *string `json:"CosSecretKey,omitnil,omitempty" name:"CosSecretKey"`

	// cos的appid，已废弃
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type PrePaySetting struct {
	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Period *Period `json:"Period,omitnil,omitempty" name:"Period"`

	// 自动续费标记，0：表示通知即将过期，但不自动续费 1：表示通知即将过期，而且自动续费 2：表示不通知即将过期，也不自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

type PreferredSchedulingTerm struct {
	// 权重，范围1-100
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 节点选择表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Preference *NodeSelectorTerm `json:"Preference,omitnil,omitempty" name:"Preference"`
}

type PriceDetail struct {
	// 节点ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 价格计算公式
	Formula *string `json:"Formula,omitnil,omitempty" name:"Formula"`

	// 原价
	OriginalCost *float64 `json:"OriginalCost,omitnil,omitempty" name:"OriginalCost"`

	// 折扣价
	DiscountCost *float64 `json:"DiscountCost,omitnil,omitempty" name:"DiscountCost"`
}

type PriceResource struct {
	// 需要的规格
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// 硬盘类型
	StorageType *uint64 `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 硬盘类型
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 系统盘大小
	RootSize *int64 `json:"RootSize,omitnil,omitempty" name:"RootSize"`

	// 内存大小
	MemSize *int64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// 核心数量
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 硬盘大小
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 云盘列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiDisks []*MultiDisk `json:"MultiDisks,omitnil,omitempty" name:"MultiDisks"`

	// 磁盘数量
	DiskCnt *int64 `json:"DiskCnt,omitnil,omitempty" name:"DiskCnt"`

	// 规格
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 磁盘数量
	DiskNum *int64 `json:"DiskNum,omitnil,omitempty" name:"DiskNum"`

	// 本地盘的数量
	LocalDiskNum *int64 `json:"LocalDiskNum,omitnil,omitempty" name:"LocalDiskNum"`
}

type PriceResult struct {
	// 原价
	OriginalCost *float64 `json:"OriginalCost,omitnil,omitempty" name:"OriginalCost"`

	// 折扣价
	DiscountCost *float64 `json:"DiscountCost,omitnil,omitempty" name:"DiscountCost"`
}

type QuotaEntity struct {
	// 已使用配额
	UsedQuota *int64 `json:"UsedQuota,omitnil,omitempty" name:"UsedQuota"`

	// 剩余配额
	RemainingQuota *int64 `json:"RemainingQuota,omitnil,omitempty" name:"RemainingQuota"`

	// 总配额
	TotalQuota *int64 `json:"TotalQuota,omitnil,omitempty" name:"TotalQuota"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type RenewInstancesInfo struct {
	// 节点资源ID
	EmrResourceId *string `json:"EmrResourceId,omitnil,omitempty" name:"EmrResourceId"`

	// 节点类型。0:common节点；1:master节点
	// ；2:core节点；3:task节点
	Flag *int64 `json:"Flag,omitnil,omitempty" name:"Flag"`

	// 内网IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 节点内存描述
	MemDesc *string `json:"MemDesc,omitnil,omitempty" name:"MemDesc"`

	// 节点核数
	CpuNum *int64 `json:"CpuNum,omitnil,omitempty" name:"CpuNum"`

	// 硬盘大小
	DiskSize *string `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 节点规格
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// 磁盘类型
	StorageType *int64 `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 系统盘大小
	RootSize *int64 `json:"RootSize,omitnil,omitempty" name:"RootSize"`

	// 系统盘类型
	RootStorageType *int64 `json:"RootStorageType,omitnil,omitempty" name:"RootStorageType"`

	// 数据盘信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MCMultiDisk []*MultiDiskMC `json:"MCMultiDisk,omitnil,omitempty" name:"MCMultiDisk"`
}

type RepeatStrategy struct {
	// 取值范围"DAY","DOW","DOM","NONE"，分别表示按天重复、按周重复、按月重复和一次执行。必须填写
	RepeatType *string `json:"RepeatType,omitnil,omitempty" name:"RepeatType"`

	// 按天重复规则，当RepeatType为"DAY"时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	DayRepeat *DayRepeatStrategy `json:"DayRepeat,omitnil,omitempty" name:"DayRepeat"`

	// 按周重复规则，当RepeatType为"DOW"时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeekRepeat *WeekRepeatStrategy `json:"WeekRepeat,omitnil,omitempty" name:"WeekRepeat"`

	// 按月重复规则，当RepeatType为"DOM"时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonthRepeat *MonthRepeatStrategy `json:"MonthRepeat,omitnil,omitempty" name:"MonthRepeat"`

	// 一次执行规则，当RepeatType为"NONE"时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotRepeat *NotRepeatStrategy `json:"NotRepeat,omitnil,omitempty" name:"NotRepeat"`

	// 规则过期时间，超过该时间后，规则将自动置为暂停状态，形式为"2020-07-23 00:00:00"。必须填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Expire *string `json:"Expire,omitnil,omitempty" name:"Expire"`

	// 周期性规则开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

// Predefined struct for user
type ResetYarnConfigRequestParams struct {
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要重置的配置别名，可选值：
	// 
	// - capacityLabel：重置标签管理的配置
	// - fair：重置公平调度的配置
	// - capacity：重置容量调度的配置
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`
}

type ResetYarnConfigRequest struct {
	*tchttp.BaseRequest
	
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要重置的配置别名，可选值：
	// 
	// - capacityLabel：重置标签管理的配置
	// - fair：重置公平调度的配置
	// - capacity：重置容量调度的配置
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`
}

func (r *ResetYarnConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetYarnConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Key")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetYarnConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetYarnConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetYarnConfigResponse struct {
	*tchttp.BaseResponse
	Response *ResetYarnConfigResponseParams `json:"Response"`
}

func (r *ResetYarnConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetYarnConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeDataDisksRequestParams struct {
	// EMR集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要扩容的云盘ID
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 需要扩充的容量值，容量值需要大于原容量，并且为10的整数倍
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 需要扩容的节点ID列表
	CvmInstanceIds []*string `json:"CvmInstanceIds,omitnil,omitempty" name:"CvmInstanceIds"`
}

type ResizeDataDisksRequest struct {
	*tchttp.BaseRequest
	
	// EMR集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要扩容的云盘ID
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 需要扩充的容量值，容量值需要大于原容量，并且为10的整数倍
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 需要扩容的节点ID列表
	CvmInstanceIds []*string `json:"CvmInstanceIds,omitnil,omitempty" name:"CvmInstanceIds"`
}

func (r *ResizeDataDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeDataDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DiskIds")
	delete(f, "DiskSize")
	delete(f, "CvmInstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResizeDataDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeDataDisksResponseParams struct {
	// 流程Id
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResizeDataDisksResponse struct {
	*tchttp.BaseResponse
	Response *ResizeDataDisksResponseParams `json:"Response"`
}

func (r *ResizeDataDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeDataDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Resource struct {
	// 节点规格描述，如CVM.SA2。
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// 存储类型
	// 取值范围：
	// <li>4：表示云SSD。</li>
	// <li>5：表示高效云盘。</li>
	// <li>6：表示增强型SSD云硬盘。</li>
	// <li>11：表示吞吐型云硬盘。</li>
	// <li>12：表示极速型SSD云硬盘。</li>：创建时该类型无效，会根据数据盘类型和节点类型自动判断
	StorageType *int64 `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 磁盘类型
	// 取值范围：
	// <li>CLOUD_SSD：表示云SSD。</li>
	// <li>CLOUD_PREMIUM：表示高效云盘。</li>
	// <li>CLOUD_BASIC：表示云硬盘。</li>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 内存容量,单位为M
	MemSize *int64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// CPU核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 数据盘容量
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 系统盘容量
	RootSize *int64 `json:"RootSize,omitnil,omitempty" name:"RootSize"`

	// 云盘列表，当数据盘为一块云盘时，直接使用DiskType和DiskSize参数，超出部分使用MultiDisks
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiDisks []*MultiDisk `json:"MultiDisks,omitnil,omitempty" name:"MultiDisks"`

	// 需要绑定的标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 规格类型，如S2.MEDIUM8
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 本地盘数量，该字段已废弃
	LocalDiskNum *uint64 `json:"LocalDiskNum,omitnil,omitempty" name:"LocalDiskNum"`

	// 本地盘数量，如2
	DiskNum *uint64 `json:"DiskNum,omitnil,omitempty" name:"DiskNum"`
}

type ResourceDetail struct {
	// 规格
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// 规格名
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 硬盘类型
	StorageType *int64 `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 硬盘类型
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 系统盘大小
	RootSize *int64 `json:"RootSize,omitnil,omitempty" name:"RootSize"`

	// 内存大小
	MemSize *int64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// CPU个数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 硬盘大小
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 规格
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type RestartPolicy struct {
	// 重启策略名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略展示名称。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 策略描述。
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// 批量重启节点数可选范围。
	BatchSizeRange []*int64 `json:"BatchSizeRange,omitnil,omitempty" name:"BatchSizeRange"`

	// 是否是默认策略。
	IsDefault *string `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`
}

// Predefined struct for user
type RunJobFlowRequestParams struct {
	// 作业名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否新创建集群。
	// true，新创建集群，则使用Instance中的参数进行集群创建。
	// false，使用已有集群，则通过InstanceId传入。
	CreateCluster *bool `json:"CreateCluster,omitnil,omitempty" name:"CreateCluster"`

	// 作业流程执行步骤。
	Steps []*Step `json:"Steps,omitnil,omitempty" name:"Steps"`

	// 作业流程正常完成时，集群的处理方式，可选择:
	// Terminate 销毁集群。
	// Reserve 保留集群。
	InstancePolicy *string `json:"InstancePolicy,omitnil,omitempty" name:"InstancePolicy"`

	// 只有CreateCluster为true时生效，目前只支持EMR版本，例如EMR-2.2.0，不支持ClickHouse和Druid版本。
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// 只在CreateCluster为true时生效。
	// true 表示安装kerberos，false表示不安装kerberos。
	SecurityClusterFlag *bool `json:"SecurityClusterFlag,omitnil,omitempty" name:"SecurityClusterFlag"`

	// 只在CreateCluster为true时生效。
	// 新建集群时，要安装的软件列表。
	Software []*string `json:"Software,omitnil,omitempty" name:"Software"`

	// 引导脚本。
	BootstrapActions []*BootstrapAction `json:"BootstrapActions,omitnil,omitempty" name:"BootstrapActions"`

	// 指定配置创建集群。
	Configurations []*Configuration `json:"Configurations,omitnil,omitempty" name:"Configurations"`

	// 作业日志保存地址。
	LogUri *string `json:"LogUri,omitnil,omitempty" name:"LogUri"`

	// 只在CreateCluster为false时生效。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自定义应用角色，大数据应用访问外部服务时使用的角色，默认为"EME_QCSRole"。
	ApplicationRole *string `json:"ApplicationRole,omitnil,omitempty" name:"ApplicationRole"`

	// 重入标签，用来可重入检查，防止在一段时间内，创建相同的流程作业。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 只在CreateCluster为true时生效，使用该配置创建集群。
	Instance *ClusterSetting `json:"Instance,omitnil,omitempty" name:"Instance"`
}

type RunJobFlowRequest struct {
	*tchttp.BaseRequest
	
	// 作业名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否新创建集群。
	// true，新创建集群，则使用Instance中的参数进行集群创建。
	// false，使用已有集群，则通过InstanceId传入。
	CreateCluster *bool `json:"CreateCluster,omitnil,omitempty" name:"CreateCluster"`

	// 作业流程执行步骤。
	Steps []*Step `json:"Steps,omitnil,omitempty" name:"Steps"`

	// 作业流程正常完成时，集群的处理方式，可选择:
	// Terminate 销毁集群。
	// Reserve 保留集群。
	InstancePolicy *string `json:"InstancePolicy,omitnil,omitempty" name:"InstancePolicy"`

	// 只有CreateCluster为true时生效，目前只支持EMR版本，例如EMR-2.2.0，不支持ClickHouse和Druid版本。
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// 只在CreateCluster为true时生效。
	// true 表示安装kerberos，false表示不安装kerberos。
	SecurityClusterFlag *bool `json:"SecurityClusterFlag,omitnil,omitempty" name:"SecurityClusterFlag"`

	// 只在CreateCluster为true时生效。
	// 新建集群时，要安装的软件列表。
	Software []*string `json:"Software,omitnil,omitempty" name:"Software"`

	// 引导脚本。
	BootstrapActions []*BootstrapAction `json:"BootstrapActions,omitnil,omitempty" name:"BootstrapActions"`

	// 指定配置创建集群。
	Configurations []*Configuration `json:"Configurations,omitnil,omitempty" name:"Configurations"`

	// 作业日志保存地址。
	LogUri *string `json:"LogUri,omitnil,omitempty" name:"LogUri"`

	// 只在CreateCluster为false时生效。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自定义应用角色，大数据应用访问外部服务时使用的角色，默认为"EME_QCSRole"。
	ApplicationRole *string `json:"ApplicationRole,omitnil,omitempty" name:"ApplicationRole"`

	// 重入标签，用来可重入检查，防止在一段时间内，创建相同的流程作业。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 只在CreateCluster为true时生效，使用该配置创建集群。
	Instance *ClusterSetting `json:"Instance,omitnil,omitempty" name:"Instance"`
}

func (r *RunJobFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunJobFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "CreateCluster")
	delete(f, "Steps")
	delete(f, "InstancePolicy")
	delete(f, "ProductVersion")
	delete(f, "SecurityClusterFlag")
	delete(f, "Software")
	delete(f, "BootstrapActions")
	delete(f, "Configurations")
	delete(f, "LogUri")
	delete(f, "InstanceId")
	delete(f, "ApplicationRole")
	delete(f, "ClientToken")
	delete(f, "Instance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunJobFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunJobFlowResponseParams struct {
	// 作业流程ID。
	JobFlowId *int64 `json:"JobFlowId,omitnil,omitempty" name:"JobFlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunJobFlowResponse struct {
	*tchttp.BaseResponse
	Response *RunJobFlowResponseParams `json:"Response"`
}

func (r *RunJobFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunJobFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SLInstanceInfo struct {
	// 集群实例字符串ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群实例数字ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 健康状态
	HealthStatus *string `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// 实例名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 地域ID
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 主可用区ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 主可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 用户APPID
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 主可用区私有网络ID
	VpcId *uint64 `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 主可用区子网ID
	SubnetId *uint64 `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 状态码
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 集群计费类型。0表示按量计费，1表示包年包月
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 多可用区信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneSettings []*ZoneSetting `json:"ZoneSettings,omitnil,omitempty" name:"ZoneSettings"`

	// 实例标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 自动续费标记， 0：表示通知即将过期，但不自动续费 1：表示通知即将过期，而且自动续费 2：表示不通知即将过期，也不自动续费，若业务无续费概念，设置为0
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 隔离时间，未隔离返回0000-00-00 00:00:00。
	IsolateTime *string `json:"IsolateTime,omitnil,omitempty" name:"IsolateTime"`

	// 过期时间，后付费返回0000-00-00 00:00:00
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

// Predefined struct for user
type ScaleOutClusterRequestParams struct {
	// 节点计费模式。取值范围：
	// <li>PREPAID：预付费，即包年包月。</li>
	// <li>POSTPAID_BY_HOUR：按小时后付费。</li>
	// <li>SPOTPAID：竞价付费（仅支持TASK节点）。</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 集群实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 扩容节点类型以及数量
	ScaleOutNodeConfig *ScaleOutNodeConfig `json:"ScaleOutNodeConfig,omitnil,omitempty" name:"ScaleOutNodeConfig"`

	// 唯一随机标识，时效5分钟，需要调用者指定 防止客户端重新创建资源，例如 a9a90aa6-****-****-****-fae36063280
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// [引导操作](https://cloud.tencent.com/document/product/589/35656)脚本设置。
	ScriptBootstrapActionConfig []*ScriptBootstrapActionConfig `json:"ScriptBootstrapActionConfig,omitnil,omitempty" name:"ScriptBootstrapActionConfig"`

	// 扩容部署服务，新增节点将默认继承当前节点类型中所部署服务，部署服务含默认可选服务，该参数仅支持可选服务填写，如：存量task节点已部署HDFS、YARN、impala；使用api扩容task节不部署impala时，部署服务仅填写HDFS、YARN。[组件名对应的映射关系表](https://cloud.tencent.com/document/product/589/98760)。
	SoftDeployInfo []*int64 `json:"SoftDeployInfo,omitnil,omitempty" name:"SoftDeployInfo"`

	// 部署进程，默认部署扩容服务的全部进程，支持修改部署进程，如：当前task节点部署服务为：HDFS、YARN、impala，默认部署服务为：DataNode,NodeManager,ImpalaServer，若用户需修改部署进程信息，部署进程：	DataNode,NodeManager,ImpalaServerCoordinator或DataNode,NodeManager,ImpalaServerExecutor。[进程名对应的映射关系表](https://cloud.tencent.com/document/product/589/98760)。
	ServiceNodeInfo []*int64 `json:"ServiceNodeInfo,omitnil,omitempty" name:"ServiceNodeInfo"`

	// 分散置放群组ID列表，当前只支持指定一个。
	// 该参数可以通过调用 [DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/product/213/17810)的返回值中的DisasterRecoverGroupId字段来获取。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 扩容节点绑定标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 扩容所选资源类型，可选范围为"host","pod"，host为普通的CVM资源，Pod为TKE集群或EKS集群提供的资源
	HardwareSourceType *string `json:"HardwareSourceType,omitnil,omitempty" name:"HardwareSourceType"`

	// Pod相关资源信息
	PodSpecInfo *PodSpecInfo `json:"PodSpecInfo,omitnil,omitempty" name:"PodSpecInfo"`

	// 使用clickhouse集群扩容时，选择的机器分组名称
	ClickHouseClusterName *string `json:"ClickHouseClusterName,omitnil,omitempty" name:"ClickHouseClusterName"`

	// 使用clickhouse集群扩容时，选择的机器分组类型。new为新增，old为选择旧分组
	ClickHouseClusterType *string `json:"ClickHouseClusterType,omitnil,omitempty" name:"ClickHouseClusterType"`

	// 扩容指定 Yarn Node Label
	YarnNodeLabel *string `json:"YarnNodeLabel,omitnil,omitempty" name:"YarnNodeLabel"`

	// 扩容后是否启动服务，默认取值否
	// <li>true：是</li>
	// <li>false：否</li>
	EnableStartServiceFlag *bool `json:"EnableStartServiceFlag,omitnil,omitempty" name:"EnableStartServiceFlag"`

	// 规格设置
	ResourceSpec *NodeResourceSpec `json:"ResourceSpec,omitnil,omitempty" name:"ResourceSpec"`

	// 实例所属的可用区，例如ap-guangzhou-1。该参数也可以通过调用[DescribeZones](https://cloud.tencent.com/document/product/213/15707) 的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 子网，默认是集群创建时的子网
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 扩容指定配置组
	ScaleOutServiceConfGroupsInfo []*ScaleOutServiceConfGroupsInfo `json:"ScaleOutServiceConfGroupsInfo,omitnil,omitempty" name:"ScaleOutServiceConfGroupsInfo"`
}

type ScaleOutClusterRequest struct {
	*tchttp.BaseRequest
	
	// 节点计费模式。取值范围：
	// <li>PREPAID：预付费，即包年包月。</li>
	// <li>POSTPAID_BY_HOUR：按小时后付费。</li>
	// <li>SPOTPAID：竞价付费（仅支持TASK节点）。</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 集群实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 扩容节点类型以及数量
	ScaleOutNodeConfig *ScaleOutNodeConfig `json:"ScaleOutNodeConfig,omitnil,omitempty" name:"ScaleOutNodeConfig"`

	// 唯一随机标识，时效5分钟，需要调用者指定 防止客户端重新创建资源，例如 a9a90aa6-****-****-****-fae36063280
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// [引导操作](https://cloud.tencent.com/document/product/589/35656)脚本设置。
	ScriptBootstrapActionConfig []*ScriptBootstrapActionConfig `json:"ScriptBootstrapActionConfig,omitnil,omitempty" name:"ScriptBootstrapActionConfig"`

	// 扩容部署服务，新增节点将默认继承当前节点类型中所部署服务，部署服务含默认可选服务，该参数仅支持可选服务填写，如：存量task节点已部署HDFS、YARN、impala；使用api扩容task节不部署impala时，部署服务仅填写HDFS、YARN。[组件名对应的映射关系表](https://cloud.tencent.com/document/product/589/98760)。
	SoftDeployInfo []*int64 `json:"SoftDeployInfo,omitnil,omitempty" name:"SoftDeployInfo"`

	// 部署进程，默认部署扩容服务的全部进程，支持修改部署进程，如：当前task节点部署服务为：HDFS、YARN、impala，默认部署服务为：DataNode,NodeManager,ImpalaServer，若用户需修改部署进程信息，部署进程：	DataNode,NodeManager,ImpalaServerCoordinator或DataNode,NodeManager,ImpalaServerExecutor。[进程名对应的映射关系表](https://cloud.tencent.com/document/product/589/98760)。
	ServiceNodeInfo []*int64 `json:"ServiceNodeInfo,omitnil,omitempty" name:"ServiceNodeInfo"`

	// 分散置放群组ID列表，当前只支持指定一个。
	// 该参数可以通过调用 [DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/product/213/17810)的返回值中的DisasterRecoverGroupId字段来获取。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 扩容节点绑定标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 扩容所选资源类型，可选范围为"host","pod"，host为普通的CVM资源，Pod为TKE集群或EKS集群提供的资源
	HardwareSourceType *string `json:"HardwareSourceType,omitnil,omitempty" name:"HardwareSourceType"`

	// Pod相关资源信息
	PodSpecInfo *PodSpecInfo `json:"PodSpecInfo,omitnil,omitempty" name:"PodSpecInfo"`

	// 使用clickhouse集群扩容时，选择的机器分组名称
	ClickHouseClusterName *string `json:"ClickHouseClusterName,omitnil,omitempty" name:"ClickHouseClusterName"`

	// 使用clickhouse集群扩容时，选择的机器分组类型。new为新增，old为选择旧分组
	ClickHouseClusterType *string `json:"ClickHouseClusterType,omitnil,omitempty" name:"ClickHouseClusterType"`

	// 扩容指定 Yarn Node Label
	YarnNodeLabel *string `json:"YarnNodeLabel,omitnil,omitempty" name:"YarnNodeLabel"`

	// 扩容后是否启动服务，默认取值否
	// <li>true：是</li>
	// <li>false：否</li>
	EnableStartServiceFlag *bool `json:"EnableStartServiceFlag,omitnil,omitempty" name:"EnableStartServiceFlag"`

	// 规格设置
	ResourceSpec *NodeResourceSpec `json:"ResourceSpec,omitnil,omitempty" name:"ResourceSpec"`

	// 实例所属的可用区，例如ap-guangzhou-1。该参数也可以通过调用[DescribeZones](https://cloud.tencent.com/document/product/213/15707) 的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 子网，默认是集群创建时的子网
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 扩容指定配置组
	ScaleOutServiceConfGroupsInfo []*ScaleOutServiceConfGroupsInfo `json:"ScaleOutServiceConfGroupsInfo,omitnil,omitempty" name:"ScaleOutServiceConfGroupsInfo"`
}

func (r *ScaleOutClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceChargeType")
	delete(f, "InstanceId")
	delete(f, "ScaleOutNodeConfig")
	delete(f, "ClientToken")
	delete(f, "InstanceChargePrepaid")
	delete(f, "ScriptBootstrapActionConfig")
	delete(f, "SoftDeployInfo")
	delete(f, "ServiceNodeInfo")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "Tags")
	delete(f, "HardwareSourceType")
	delete(f, "PodSpecInfo")
	delete(f, "ClickHouseClusterName")
	delete(f, "ClickHouseClusterType")
	delete(f, "YarnNodeLabel")
	delete(f, "EnableStartServiceFlag")
	delete(f, "ResourceSpec")
	delete(f, "Zone")
	delete(f, "SubnetId")
	delete(f, "ScaleOutServiceConfGroupsInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleOutClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutClusterResponseParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 客户端Token。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 扩容流程ID。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 查询流程状态，流程额外信息
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// 订单号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 大订单号。
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleOutClusterResponse struct {
	*tchttp.BaseResponse
	Response *ScaleOutClusterResponseParams `json:"Response"`
}

func (r *ScaleOutClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceRequestParams struct {
	// 扩容的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 扩容的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 唯一随机标识，时效5分钟，需要调用者指定 防止客户端重新创建资源，例如 a9a90aa6-****-****-****-fae36063280
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 引导操作脚本设置。
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitnil,omitempty" name:"PreExecutedFileSettings"`

	// 扩容的Task节点数量。
	TaskCount *uint64 `json:"TaskCount,omitnil,omitempty" name:"TaskCount"`

	// 扩容的Core节点数量。
	CoreCount *uint64 `json:"CoreCount,omitnil,omitempty" name:"CoreCount"`

	// 扩容时不需要安装的进程。
	UnNecessaryNodeList []*uint64 `json:"UnNecessaryNodeList,omitnil,omitempty" name:"UnNecessaryNodeList"`

	// 扩容的Router节点数量。
	RouterCount *uint64 `json:"RouterCount,omitnil,omitempty" name:"RouterCount"`

	// 部署的服务。
	// <li>SoftDeployInfo和ServiceNodeInfo是同组参数，和UnNecessaryNodeList参数互斥。</li>
	// <li>建议使用SoftDeployInfo和ServiceNodeInfo组合。</li>
	SoftDeployInfo []*uint64 `json:"SoftDeployInfo,omitnil,omitempty" name:"SoftDeployInfo"`

	// 启动的进程。
	ServiceNodeInfo []*uint64 `json:"ServiceNodeInfo,omitnil,omitempty" name:"ServiceNodeInfo"`

	// 分散置放群组ID列表，当前仅支持指定一个。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 扩容节点绑定标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 扩容所选资源类型，可选范围为"host","pod"，host为普通的CVM资源，Pod为TKE集群或EKS集群提供的资源
	HardwareResourceType *string `json:"HardwareResourceType,omitnil,omitempty" name:"HardwareResourceType"`

	// 使用Pod资源扩容时，指定的Pod规格以及来源等信息
	PodSpec *PodSpec `json:"PodSpec,omitnil,omitempty" name:"PodSpec"`

	// 使用clickhouse集群扩容时，选择的机器分组名称
	ClickHouseClusterName *string `json:"ClickHouseClusterName,omitnil,omitempty" name:"ClickHouseClusterName"`

	// 使用clickhouse集群扩容时，选择的机器分组类型。new为新增，old为选择旧分组
	ClickHouseClusterType *string `json:"ClickHouseClusterType,omitnil,omitempty" name:"ClickHouseClusterType"`

	// 规则扩容指定 yarn node label
	YarnNodeLabel *string `json:"YarnNodeLabel,omitnil,omitempty" name:"YarnNodeLabel"`

	// POD自定义权限和自定义参数
	PodParameter *PodParameter `json:"PodParameter,omitnil,omitempty" name:"PodParameter"`

	// 扩容的Master节点的数量。
	// 使用clickhouse集群扩容时，该参数不生效。
	// 使用kafka集群扩容时，该参数不生效。
	// 当HardwareResourceType=POD时，该参数不生效。
	MasterCount *uint64 `json:"MasterCount,omitnil,omitempty" name:"MasterCount"`

	// 扩容后是否启动服务，true：启动，false：不启动
	StartServiceAfterScaleOut *string `json:"StartServiceAfterScaleOut,omitnil,omitempty" name:"StartServiceAfterScaleOut"`

	// 可用区，默认是集群的主可用区
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 子网，默认是集群创建时的子网
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 预设配置组
	ScaleOutServiceConfAssign *string `json:"ScaleOutServiceConfAssign,omitnil,omitempty" name:"ScaleOutServiceConfAssign"`

	// 0表示关闭自动续费，1表示开启自动续费
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 类型为ComputeResource和EMR以及默认，默认为EMR,类型为EMR时,InstanceId生效,类型为ComputeResource时,使用ComputeResourceId标识
	ResourceBaseType *string `json:"ResourceBaseType,omitnil,omitempty" name:"ResourceBaseType"`

	// 计算资源id
	ComputeResourceId *string `json:"ComputeResourceId,omitnil,omitempty" name:"ComputeResourceId"`
}

type ScaleOutInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 扩容的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 扩容的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 唯一随机标识，时效5分钟，需要调用者指定 防止客户端重新创建资源，例如 a9a90aa6-****-****-****-fae36063280
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 引导操作脚本设置。
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitnil,omitempty" name:"PreExecutedFileSettings"`

	// 扩容的Task节点数量。
	TaskCount *uint64 `json:"TaskCount,omitnil,omitempty" name:"TaskCount"`

	// 扩容的Core节点数量。
	CoreCount *uint64 `json:"CoreCount,omitnil,omitempty" name:"CoreCount"`

	// 扩容时不需要安装的进程。
	UnNecessaryNodeList []*uint64 `json:"UnNecessaryNodeList,omitnil,omitempty" name:"UnNecessaryNodeList"`

	// 扩容的Router节点数量。
	RouterCount *uint64 `json:"RouterCount,omitnil,omitempty" name:"RouterCount"`

	// 部署的服务。
	// <li>SoftDeployInfo和ServiceNodeInfo是同组参数，和UnNecessaryNodeList参数互斥。</li>
	// <li>建议使用SoftDeployInfo和ServiceNodeInfo组合。</li>
	SoftDeployInfo []*uint64 `json:"SoftDeployInfo,omitnil,omitempty" name:"SoftDeployInfo"`

	// 启动的进程。
	ServiceNodeInfo []*uint64 `json:"ServiceNodeInfo,omitnil,omitempty" name:"ServiceNodeInfo"`

	// 分散置放群组ID列表，当前仅支持指定一个。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 扩容节点绑定标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 扩容所选资源类型，可选范围为"host","pod"，host为普通的CVM资源，Pod为TKE集群或EKS集群提供的资源
	HardwareResourceType *string `json:"HardwareResourceType,omitnil,omitempty" name:"HardwareResourceType"`

	// 使用Pod资源扩容时，指定的Pod规格以及来源等信息
	PodSpec *PodSpec `json:"PodSpec,omitnil,omitempty" name:"PodSpec"`

	// 使用clickhouse集群扩容时，选择的机器分组名称
	ClickHouseClusterName *string `json:"ClickHouseClusterName,omitnil,omitempty" name:"ClickHouseClusterName"`

	// 使用clickhouse集群扩容时，选择的机器分组类型。new为新增，old为选择旧分组
	ClickHouseClusterType *string `json:"ClickHouseClusterType,omitnil,omitempty" name:"ClickHouseClusterType"`

	// 规则扩容指定 yarn node label
	YarnNodeLabel *string `json:"YarnNodeLabel,omitnil,omitempty" name:"YarnNodeLabel"`

	// POD自定义权限和自定义参数
	PodParameter *PodParameter `json:"PodParameter,omitnil,omitempty" name:"PodParameter"`

	// 扩容的Master节点的数量。
	// 使用clickhouse集群扩容时，该参数不生效。
	// 使用kafka集群扩容时，该参数不生效。
	// 当HardwareResourceType=POD时，该参数不生效。
	MasterCount *uint64 `json:"MasterCount,omitnil,omitempty" name:"MasterCount"`

	// 扩容后是否启动服务，true：启动，false：不启动
	StartServiceAfterScaleOut *string `json:"StartServiceAfterScaleOut,omitnil,omitempty" name:"StartServiceAfterScaleOut"`

	// 可用区，默认是集群的主可用区
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 子网，默认是集群创建时的子网
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 预设配置组
	ScaleOutServiceConfAssign *string `json:"ScaleOutServiceConfAssign,omitnil,omitempty" name:"ScaleOutServiceConfAssign"`

	// 0表示关闭自动续费，1表示开启自动续费
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 类型为ComputeResource和EMR以及默认，默认为EMR,类型为EMR时,InstanceId生效,类型为ComputeResource时,使用ComputeResourceId标识
	ResourceBaseType *string `json:"ResourceBaseType,omitnil,omitempty" name:"ResourceBaseType"`

	// 计算资源id
	ComputeResourceId *string `json:"ComputeResourceId,omitnil,omitempty" name:"ComputeResourceId"`
}

func (r *ScaleOutInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	delete(f, "InstanceId")
	delete(f, "PayMode")
	delete(f, "ClientToken")
	delete(f, "PreExecutedFileSettings")
	delete(f, "TaskCount")
	delete(f, "CoreCount")
	delete(f, "UnNecessaryNodeList")
	delete(f, "RouterCount")
	delete(f, "SoftDeployInfo")
	delete(f, "ServiceNodeInfo")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "Tags")
	delete(f, "HardwareResourceType")
	delete(f, "PodSpec")
	delete(f, "ClickHouseClusterName")
	delete(f, "ClickHouseClusterType")
	delete(f, "YarnNodeLabel")
	delete(f, "PodParameter")
	delete(f, "MasterCount")
	delete(f, "StartServiceAfterScaleOut")
	delete(f, "ZoneId")
	delete(f, "SubnetId")
	delete(f, "ScaleOutServiceConfAssign")
	delete(f, "AutoRenew")
	delete(f, "ResourceBaseType")
	delete(f, "ComputeResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleOutInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceResponseParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 订单号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 客户端Token。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 扩容流程ID。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 大订单号。
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// 扩容TraceId
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleOutInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ScaleOutInstanceResponseParams `json:"Response"`
}

func (r *ScaleOutInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScaleOutNodeConfig struct {
	// 扩容节点类型取值范围：
	//   <li>MASTER</li>
	//   <li>TASK</li>
	//   <li>CORE</li>
	//   <li>ROUTER</li>
	NodeFlag *string `json:"NodeFlag,omitnil,omitempty" name:"NodeFlag"`

	// 扩容节点数量
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`
}

type ScaleOutServiceConfGroupsInfo struct {
	// 组件版本名称 如 HDFS-2.8.5
	ServiceComponentName *string `json:"ServiceComponentName,omitnil,omitempty" name:"ServiceComponentName"`

	// 配置组名 如hdfs-core-defaultGroup    ConfGroupName参数传入 代表配置组维度 
	//                                                              ConfGroupName参数不传 默认 代表集群维度
	ConfGroupName *string `json:"ConfGroupName,omitnil,omitempty" name:"ConfGroupName"`
}

type SceneSoftwareConfig struct {
	// 部署的组件列表。不同的EMR产品版本ProductVersion 对应不同可选组件列表，不同产品版本可选组件列表查询：[组件版本](https://cloud.tencent.com/document/product/589/20279) ；
	// 填写实例值：hive、flink。
	Software []*string `json:"Software,omitnil,omitempty" name:"Software"`

	// 默认Hadoop-Default,[场景查询](https://cloud.tencent.com/document/product/589/14624)场景化取值范围：
	// Hadoop-Kudu
	// Hadoop-Zookeeper
	// Hadoop-Presto
	// Hadoop-Hbase
	// Hadoop-Default
	SceneName *string `json:"SceneName,omitnil,omitempty" name:"SceneName"`
}

type SchedulerTaskDetail struct {
	// 步骤
	Step *string `json:"Step,omitnil,omitempty" name:"Step"`

	// 进度
	Progress *string `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 失败信息
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`

	// 用来获取详情的id
	JobId *uint64 `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type SchedulerTaskInfo struct {
	// 调度器类型
	SchedulerName *string `json:"SchedulerName,omitnil,omitempty" name:"SchedulerName"`

	// 操作类型
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 开始时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 状态
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details []*SchedulerTaskDetail `json:"Details,omitnil,omitempty" name:"Details"`
}

type ScriptBootstrapActionConfig struct {
	// 脚本的cos地址，参照格式：https://beijing-111111.cos.ap-beijing.myqcloud.com/data/test.sh查询cos存储桶列表：[存储桶列表](https://console.cloud.tencent.com/cos/bucket)
	CosFileURI *string `json:"CosFileURI,omitnil,omitempty" name:"CosFileURI"`

	// 引导脚步执行时机范围
	// <li>resourceAfter：节点初始化后</li>
	// <li>clusterAfter：集群启动后</li>
	// <li>clusterBefore：集群启动前</li>
	ExecutionMoment *string `json:"ExecutionMoment,omitnil,omitempty" name:"ExecutionMoment"`

	// 执行脚本参数，参数格式请遵循标准Shell规范
	Args []*string `json:"Args,omitnil,omitempty" name:"Args"`

	// 脚本文件名
	CosFileName *string `json:"CosFileName,omitnil,omitempty" name:"CosFileName"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type SearchItem struct {
	// 支持搜索的类型
	SearchType *string `json:"SearchType,omitnil,omitempty" name:"SearchType"`

	// 支持搜索的值
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`
}

type ServiceBasicRestartInfo struct {
	// 服务名，必填，如HDFS
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 如果没传，则表示所有进程
	ComponentInfoList []*ComponentBasicRestartInfo `json:"ComponentInfoList,omitnil,omitempty" name:"ComponentInfoList"`
}

type ServiceNodeDetailInfo struct {
	// 进程所在节点IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 进程类型
	NodeType *int64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 进程名称
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 服务组件状态
	ServiceStatus *int64 `json:"ServiceStatus,omitnil,omitempty" name:"ServiceStatus"`

	// 进程监控状态
	MonitorStatus *int64 `json:"MonitorStatus,omitnil,omitempty" name:"MonitorStatus"`

	// 服务组件状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 进程端口信息
	PortsInfo *string `json:"PortsInfo,omitnil,omitempty" name:"PortsInfo"`

	// 最近重启时间
	LastRestartTime *string `json:"LastRestartTime,omitnil,omitempty" name:"LastRestartTime"`

	// 节点类型
	Flag *int64 `json:"Flag,omitnil,omitempty" name:"Flag"`

	// 配置组ID
	ConfGroupId *int64 `json:"ConfGroupId,omitnil,omitempty" name:"ConfGroupId"`

	// 配置组名称
	ConfGroupName *string `json:"ConfGroupName,omitnil,omitempty" name:"ConfGroupName"`

	// 节点是否需要重启
	ConfStatus *int64 `json:"ConfStatus,omitnil,omitempty" name:"ConfStatus"`

	// 进程探测信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceDetectionInfo []*ServiceProcessFunctionInfo `json:"ServiceDetectionInfo,omitnil,omitempty" name:"ServiceDetectionInfo"`

	// 节点类型
	NodeFlagFilter *string `json:"NodeFlagFilter,omitnil,omitempty" name:"NodeFlagFilter"`

	// 进程健康状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthStatus *HealthStatus `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// 角色是否支持监控
	IsSupportRoleMonitor *bool `json:"IsSupportRoleMonitor,omitnil,omitempty" name:"IsSupportRoleMonitor"`

	// 暂停策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	StopPolicies []*RestartPolicy `json:"StopPolicies,omitnil,omitempty" name:"StopPolicies"`

	// 测试环境api强校验，现网没有，emrcc接口返回有。不加会报错
	HAState *string `json:"HAState,omitnil,omitempty" name:"HAState"`

	// NameService名称
	NameService *string `json:"NameService,omitnil,omitempty" name:"NameService"`

	// 是否支持联邦
	IsFederation *bool `json:"IsFederation,omitnil,omitempty" name:"IsFederation"`

	// datanode是否是维护状态
	DataNodeMaintenanceState *int64 `json:"DataNodeMaintenanceState,omitnil,omitempty" name:"DataNodeMaintenanceState"`
}

type ServiceProcessFunctionInfo struct {
	// 探测告警级别
	DetectAlert *string `json:"DetectAlert,omitnil,omitempty" name:"DetectAlert"`

	// 探测功能描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: DetetcFunctionKey is deprecated.
	DetetcFunctionKey *string `json:"DetetcFunctionKey,omitnil,omitempty" name:"DetetcFunctionKey"`

	// 探测功能结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: DetetcFunctionValue is deprecated.
	DetetcFunctionValue *string `json:"DetetcFunctionValue,omitnil,omitempty" name:"DetetcFunctionValue"`

	// 探测结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: DetetcTime is deprecated.
	DetetcTime *string `json:"DetetcTime,omitnil,omitempty" name:"DetetcTime"`

	// 探测功能描述
	DetectFunctionKey *string `json:"DetectFunctionKey,omitnil,omitempty" name:"DetectFunctionKey"`

	// 探测功能结果
	DetectFunctionValue *string `json:"DetectFunctionValue,omitnil,omitempty" name:"DetectFunctionValue"`

	// 探测结果
	DetectTime *string `json:"DetectTime,omitnil,omitempty" name:"DetectTime"`
}

// Predefined struct for user
type SetNodeResourceConfigDefaultRequestParams struct {
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置Id
	ResourceConfigId *uint64 `json:"ResourceConfigId,omitnil,omitempty" name:"ResourceConfigId"`

	// 规格节点类型 CORE TASK ROUTER
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 类型为ComputeResource和EMR以及默认，默认为EMR
	ResourceBaseType *string `json:"ResourceBaseType,omitnil,omitempty" name:"ResourceBaseType"`

	// 计算资源id
	ComputeResourceId *string `json:"ComputeResourceId,omitnil,omitempty" name:"ComputeResourceId"`

	// 硬件类型
	HardwareResourceType *string `json:"HardwareResourceType,omitnil,omitempty" name:"HardwareResourceType"`
}

type SetNodeResourceConfigDefaultRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置Id
	ResourceConfigId *uint64 `json:"ResourceConfigId,omitnil,omitempty" name:"ResourceConfigId"`

	// 规格节点类型 CORE TASK ROUTER
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 类型为ComputeResource和EMR以及默认，默认为EMR
	ResourceBaseType *string `json:"ResourceBaseType,omitnil,omitempty" name:"ResourceBaseType"`

	// 计算资源id
	ComputeResourceId *string `json:"ComputeResourceId,omitnil,omitempty" name:"ComputeResourceId"`

	// 硬件类型
	HardwareResourceType *string `json:"HardwareResourceType,omitnil,omitempty" name:"HardwareResourceType"`
}

func (r *SetNodeResourceConfigDefaultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNodeResourceConfigDefaultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceConfigId")
	delete(f, "ResourceType")
	delete(f, "ResourceBaseType")
	delete(f, "ComputeResourceId")
	delete(f, "HardwareResourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetNodeResourceConfigDefaultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetNodeResourceConfigDefaultResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetNodeResourceConfigDefaultResponse struct {
	*tchttp.BaseResponse
	Response *SetNodeResourceConfigDefaultResponseParams `json:"Response"`
}

func (r *SetNodeResourceConfigDefaultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNodeResourceConfigDefaultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShortNodeInfo struct {
	// 节点类型，Master/Core/Task/Router/Common
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点数量
	NodeSize *uint64 `json:"NodeSize,omitnil,omitempty" name:"NodeSize"`
}

type SoftDependInfo struct {
	// 组件名称
	SoftName *string `json:"SoftName,omitnil,omitempty" name:"SoftName"`

	// 是否必选
	Required *bool `json:"Required,omitnil,omitempty" name:"Required"`
}

type SparkQuery struct {
	// 执行语句
	Statement *string `json:"Statement,omitnil,omitempty" name:"Statement"`

	// 执行时长（单位毫秒）
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 执行状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 查询ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 扫描分区数
	ScanPartitionNum *int64 `json:"ScanPartitionNum,omitnil,omitempty" name:"ScanPartitionNum"`

	// 扫描总行数
	ScanRowNum *int64 `json:"ScanRowNum,omitnil,omitempty" name:"ScanRowNum"`

	// 扫描总文件数
	ScanFileNum *int64 `json:"ScanFileNum,omitnil,omitempty" name:"ScanFileNum"`

	// 查询扫描总数据量(单位B)
	ScanTotalData *int64 `json:"ScanTotalData,omitnil,omitempty" name:"ScanTotalData"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId []*string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 输出总行数
	OutputRowNum *int64 `json:"OutputRowNum,omitnil,omitempty" name:"OutputRowNum"`

	// 输出总文件数
	OutputFileNum *int64 `json:"OutputFileNum,omitnil,omitempty" name:"OutputFileNum"`

	// 输出分区数
	OutputPartitionNum *int64 `json:"OutputPartitionNum,omitnil,omitempty" name:"OutputPartitionNum"`

	// 输出总数据量（单位B）
	OutputTotalData *int64 `json:"OutputTotalData,omitnil,omitempty" name:"OutputTotalData"`

	// 开始时间
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type StageInfoDetail struct {
	// 步骤
	Stage *string `json:"Stage,omitnil,omitempty" name:"Stage"`

	// 步骤名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否展示
	IsShow *bool `json:"IsShow,omitnil,omitempty" name:"IsShow"`

	// 是否子流程
	IsSubFlow *bool `json:"IsSubFlow,omitnil,omitempty" name:"IsSubFlow"`

	// 子流程标签
	SubFlowFlag *string `json:"SubFlowFlag,omitnil,omitempty" name:"SubFlowFlag"`

	// 步骤运行状态：0:未开始 1:进行中 2:已完成 3:部分完成  -1:失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 步骤运行状态描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 运行进度
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Starttime *string `json:"Starttime,omitnil,omitempty" name:"Starttime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Endtime *string `json:"Endtime,omitnil,omitempty" name:"Endtime"`

	// 是否有详情信息
	HadWoodDetail *bool `json:"HadWoodDetail,omitnil,omitempty" name:"HadWoodDetail"`

	// Wood子流程Id
	WoodJobId *uint64 `json:"WoodJobId,omitnil,omitempty" name:"WoodJobId"`

	// 多语言版本Key
	LanguageKey *string `json:"LanguageKey,omitnil,omitempty" name:"LanguageKey"`

	// 如果stage失败，失败原因
	FailedReason *string `json:"FailedReason,omitnil,omitempty" name:"FailedReason"`

	// 步骤耗时
	TimeConsuming *string `json:"TimeConsuming,omitnil,omitempty" name:"TimeConsuming"`
}

type StarRocksQueryInfo struct {
	// 提交IP
	ClientIP *string `json:"ClientIP,omitnil,omitempty" name:"ClientIP"`

	// CPU总时间(ns)
	CPUCost *int64 `json:"CPUCost,omitnil,omitempty" name:"CPUCost"`

	// 默认DB
	DefaultDB *string `json:"DefaultDB,omitnil,omitempty" name:"DefaultDB"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 执行IP
	ExecutionIP *string `json:"ExecutionIP,omitnil,omitempty" name:"ExecutionIP"`

	// 查询ID
	QueryID *string `json:"QueryID,omitnil,omitempty" name:"QueryID"`

	// 查询类型
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 消耗总内存(bytes)
	MemCost *int64 `json:"MemCost,omitnil,omitempty" name:"MemCost"`

	// plan阶段CPU占用(ns)
	PlanCpuCosts *int64 `json:"PlanCpuCosts,omitnil,omitempty" name:"PlanCpuCosts"`

	// plan阶段内存占用(bytes)
	PlanMemCosts *int64 `json:"PlanMemCosts,omitnil,omitempty" name:"PlanMemCosts"`

	// 执行时长
	QueryTime *int64 `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// 资源组
	ResourceGroup *string `json:"ResourceGroup,omitnil,omitempty" name:"ResourceGroup"`

	// 获取行数
	ReturnRows *int64 `json:"ReturnRows,omitnil,omitempty" name:"ReturnRows"`

	// 扫描数据量(bytes)
	ScanBytes *int64 `json:"ScanBytes,omitnil,omitempty" name:"ScanBytes"`

	// 扫描行数
	ScanRows *int64 `json:"ScanRows,omitnil,omitempty" name:"ScanRows"`

	// 开始时间
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 执行状态
	ExecutionState *string `json:"ExecutionState,omitnil,omitempty" name:"ExecutionState"`

	// 执行语句
	ExecutionStatement *string `json:"ExecutionStatement,omitnil,omitempty" name:"ExecutionStatement"`

	// 用户
	User *string `json:"User,omitnil,omitempty" name:"User"`
}

// Predefined struct for user
type StartStopServiceOrMonitorRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 操作类型，当前支持
	// <li>StartService：启动服务</li>
	// <li>StopService：停止服务</li>
	// <li>StartMonitor：退出维护</li>
	// <li>StopMonitor：进入维护</li>
	// <li>RestartService：重启服务 如果操作类型选择重启服务 StrategyConfig操作策略则是必填项</li>
	OpType *string `json:"OpType,omitnil,omitempty" name:"OpType"`

	// 操作范围
	OpScope *OpScope `json:"OpScope,omitnil,omitempty" name:"OpScope"`

	// 操作策略
	StrategyConfig *StrategyConfig `json:"StrategyConfig,omitnil,omitempty" name:"StrategyConfig"`

	// 暂停服务时用的参数
	StopParams *StopParams `json:"StopParams,omitnil,omitempty" name:"StopParams"`

	// 当OpType为<li>StopMonitor</li>才有用，true表示进入维护模式但是仍然监控进程但是不拉起进程
	KeepMonitorButNotRecoverProcess *bool `json:"KeepMonitorButNotRecoverProcess,omitnil,omitempty" name:"KeepMonitorButNotRecoverProcess"`
}

type StartStopServiceOrMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 操作类型，当前支持
	// <li>StartService：启动服务</li>
	// <li>StopService：停止服务</li>
	// <li>StartMonitor：退出维护</li>
	// <li>StopMonitor：进入维护</li>
	// <li>RestartService：重启服务 如果操作类型选择重启服务 StrategyConfig操作策略则是必填项</li>
	OpType *string `json:"OpType,omitnil,omitempty" name:"OpType"`

	// 操作范围
	OpScope *OpScope `json:"OpScope,omitnil,omitempty" name:"OpScope"`

	// 操作策略
	StrategyConfig *StrategyConfig `json:"StrategyConfig,omitnil,omitempty" name:"StrategyConfig"`

	// 暂停服务时用的参数
	StopParams *StopParams `json:"StopParams,omitnil,omitempty" name:"StopParams"`

	// 当OpType为<li>StopMonitor</li>才有用，true表示进入维护模式但是仍然监控进程但是不拉起进程
	KeepMonitorButNotRecoverProcess *bool `json:"KeepMonitorButNotRecoverProcess,omitnil,omitempty" name:"KeepMonitorButNotRecoverProcess"`
}

func (r *StartStopServiceOrMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartStopServiceOrMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OpType")
	delete(f, "OpScope")
	delete(f, "StrategyConfig")
	delete(f, "StopParams")
	delete(f, "KeepMonitorButNotRecoverProcess")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartStopServiceOrMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartStopServiceOrMonitorResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartStopServiceOrMonitorResponse struct {
	*tchttp.BaseResponse
	Response *StartStopServiceOrMonitorResponseParams `json:"Response"`
}

func (r *StartStopServiceOrMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartStopServiceOrMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Step struct {
	// 执行步骤名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 执行动作。
	ExecutionStep *Execution `json:"ExecutionStep,omitnil,omitempty" name:"ExecutionStep"`

	// 执行失败策略。
	// 1. TERMINATE_CLUSTER 执行失败时退出并销毁集群。
	// 2. CONTINUE 执行失败时跳过并执行后续步骤。
	ActionOnFailure *string `json:"ActionOnFailure,omitnil,omitempty" name:"ActionOnFailure"`

	// 指定执行Step时的用户名，非必须，默认为hadoop。
	User *string `json:"User,omitnil,omitempty" name:"User"`
}

type StopParams struct {
	// 安全模式：safe
	// 默认模式：default
	StopPolicy *string `json:"StopPolicy,omitnil,omitempty" name:"StopPolicy"`

	// 线程数
	ThreadCount *int64 `json:"ThreadCount,omitnil,omitempty" name:"ThreadCount"`
}

type StorageSummaryDistribution struct {
	// 数据项
	MetricItem *string `json:"MetricItem,omitnil,omitempty" name:"MetricItem"`

	// 数据项描述
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 采样值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dps []*Dps `json:"Dps,omitnil,omitempty" name:"Dps"`
}

type StrategyConfig struct {
	// 0:关闭滚动重启
	// 1:开启滚动启动
	RollingRestartSwitch *int64 `json:"RollingRestartSwitch,omitnil,omitempty" name:"RollingRestartSwitch"`

	// 滚动重启每批次的重启数量，最大重启台数为 99999 台
	BatchSize *int64 `json:"BatchSize,omitnil,omitempty" name:"BatchSize"`

	// 滚动重启每批停止等待时间 ,最大间隔为 5 分钟 单位是秒
	TimeWait *int64 `json:"TimeWait,omitnil,omitempty" name:"TimeWait"`

	// 操作失败处理策略，0:失败阻塞, 1:失败自动跳过
	DealOnFail *int64 `json:"DealOnFail,omitnil,omitempty" name:"DealOnFail"`

	// 指令需要指定的参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Args []*Arg `json:"Args,omitnil,omitempty" name:"Args"`
}

type SubnetInfo struct {
	// 子网信息（名字）
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// 子网信息（ID）
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

// Predefined struct for user
type SyncPodStateRequestParams struct {
	// EmrService中pod状态信息
	Message *PodState `json:"Message,omitnil,omitempty" name:"Message"`
}

type SyncPodStateRequest struct {
	*tchttp.BaseRequest
	
	// EmrService中pod状态信息
	Message *PodState `json:"Message,omitnil,omitempty" name:"Message"`
}

func (r *SyncPodStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncPodStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Message")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncPodStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncPodStateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SyncPodStateResponse struct {
	*tchttp.BaseResponse
	Response *SyncPodStateResponseParams `json:"Response"`
}

func (r *SyncPodStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncPodStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TableSchemaItem struct {
	// 列标识
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否可按该列排序
	Sortable *bool `json:"Sortable,omitnil,omitempty" name:"Sortable"`

	// 是否可筛选
	WithFilter *bool `json:"WithFilter,omitnil,omitempty" name:"WithFilter"`

	// 筛选的候选集
	// 注意：此字段可能返回 null，表示取不到有效值。
	Candidates []*string `json:"Candidates,omitnil,omitempty" name:"Candidates"`

	// 是否可点击
	Clickable *bool `json:"Clickable,omitnil,omitempty" name:"Clickable"`

	// 展示的名字
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`
}

type Tag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type TerminateClusterNodesRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 销毁资源列表
	CvmInstanceIds []*string `json:"CvmInstanceIds,omitnil,omitempty" name:"CvmInstanceIds"`

	// 销毁节点类型取值范围：
	//   <li>MASTER</li>
	//   <li>TASK</li>
	//   <li>CORE</li>
	//   <li>ROUTER</li>
	NodeFlag *string `json:"NodeFlag,omitnil,omitempty" name:"NodeFlag"`

	// 优雅缩容开关
	//   <li>true:开启</li>
	//   <li>false:不开启</li>
	GraceDownFlag *bool `json:"GraceDownFlag,omitnil,omitempty" name:"GraceDownFlag"`

	// 优雅缩容等待时间,时间范围60到1800  单位秒
	GraceDownTime *int64 `json:"GraceDownTime,omitnil,omitempty" name:"GraceDownTime"`
}

type TerminateClusterNodesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 销毁资源列表
	CvmInstanceIds []*string `json:"CvmInstanceIds,omitnil,omitempty" name:"CvmInstanceIds"`

	// 销毁节点类型取值范围：
	//   <li>MASTER</li>
	//   <li>TASK</li>
	//   <li>CORE</li>
	//   <li>ROUTER</li>
	NodeFlag *string `json:"NodeFlag,omitnil,omitempty" name:"NodeFlag"`

	// 优雅缩容开关
	//   <li>true:开启</li>
	//   <li>false:不开启</li>
	GraceDownFlag *bool `json:"GraceDownFlag,omitnil,omitempty" name:"GraceDownFlag"`

	// 优雅缩容等待时间,时间范围60到1800  单位秒
	GraceDownTime *int64 `json:"GraceDownTime,omitnil,omitempty" name:"GraceDownTime"`
}

func (r *TerminateClusterNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateClusterNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CvmInstanceIds")
	delete(f, "NodeFlag")
	delete(f, "GraceDownFlag")
	delete(f, "GraceDownTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateClusterNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateClusterNodesResponseParams struct {
	// 缩容流程ID。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateClusterNodesResponse struct {
	*tchttp.BaseResponse
	Response *TerminateClusterNodesResponseParams `json:"Response"`
}

func (r *TerminateClusterNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateClusterNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateInstanceRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 销毁节点ID。该参数为预留参数，用户无需配置。
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 类型为ComputeResource和EMR以及默认，默认为EMR,类型为EMR时,InstanceId生效,类型为ComputeResource时,使用ComputeResourceId标识
	ResourceBaseType *string `json:"ResourceBaseType,omitnil,omitempty" name:"ResourceBaseType"`

	// 计算资源ID
	ComputeResourceId *string `json:"ComputeResourceId,omitnil,omitempty" name:"ComputeResourceId"`
}

type TerminateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 销毁节点ID。该参数为预留参数，用户无需配置。
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 类型为ComputeResource和EMR以及默认，默认为EMR,类型为EMR时,InstanceId生效,类型为ComputeResource时,使用ComputeResourceId标识
	ResourceBaseType *string `json:"ResourceBaseType,omitnil,omitempty" name:"ResourceBaseType"`

	// 计算资源ID
	ComputeResourceId *string `json:"ComputeResourceId,omitnil,omitempty" name:"ComputeResourceId"`
}

func (r *TerminateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceIds")
	delete(f, "ResourceBaseType")
	delete(f, "ComputeResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *TerminateInstanceResponseParams `json:"Response"`
}

func (r *TerminateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateSLInstanceRequestParams struct {
	// 实例唯一标识符（字符串表示）
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type TerminateSLInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例唯一标识符（字符串表示）
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *TerminateSLInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateSLInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateSLInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateSLInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateSLInstanceResponse struct {
	*tchttp.BaseResponse
	Response *TerminateSLInstanceResponseParams `json:"Response"`
}

func (r *TerminateSLInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateSLInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateTasksRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 待销毁节点的资源ID列表。资源ID形如：emr-vm-xxxxxxxx。有效的资源ID可通过登录[控制台](https://console.cloud.tencent.com/emr)查询。
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`
}

type TerminateTasksRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 待销毁节点的资源ID列表。资源ID形如：emr-vm-xxxxxxxx。有效的资源ID可通过登录[控制台](https://console.cloud.tencent.com/emr)查询。
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`
}

func (r *TerminateTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateTasksResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateTasksResponse struct {
	*tchttp.BaseResponse
	Response *TerminateTasksResponseParams `json:"Response"`
}

func (r *TerminateTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TimeAutoScaleStrategy struct {
	// 策略名字，集群内唯一。
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 策略触发后的冷却时间，该段时间内，将不能触发弹性扩缩容。
	IntervalTime *uint64 `json:"IntervalTime,omitnil,omitempty" name:"IntervalTime"`

	// 扩缩容动作，1表示扩容，2表示缩容。
	ScaleAction *uint64 `json:"ScaleAction,omitnil,omitempty" name:"ScaleAction"`

	// 扩缩容数量。
	ScaleNum *uint64 `json:"ScaleNum,omitnil,omitempty" name:"ScaleNum"`

	// 规则状态，1表示有效，2表示无效，3表示暂停。必须填写
	StrategyStatus *uint64 `json:"StrategyStatus,omitnil,omitempty" name:"StrategyStatus"`

	// 规则优先级，越小越高。
	Priority *uint64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 当多条规则同时触发，其中某些未真正执行时，在该时间范围内，将会重试。
	RetryValidTime *uint64 `json:"RetryValidTime,omitnil,omitempty" name:"RetryValidTime"`

	// 时间扩缩容重复策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepeatStrategy *RepeatStrategy `json:"RepeatStrategy,omitnil,omitempty" name:"RepeatStrategy"`

	// 策略唯一ID。
	StrategyId *uint64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 优雅缩容开关
	GraceDownFlag *bool `json:"GraceDownFlag,omitnil,omitempty" name:"GraceDownFlag"`

	// 优雅缩容等待时间
	GraceDownTime *int64 `json:"GraceDownTime,omitnil,omitempty" name:"GraceDownTime"`

	// 绑定标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 预设配置组
	ConfigGroupAssigned *string `json:"ConfigGroupAssigned,omitnil,omitempty" name:"ConfigGroupAssigned"`

	// 扩容资源计算方法，"DEFAULT","INSTANCE", "CPU", "MEMORYGB"。
	// "DEFAULT"表示默认方式，与"INSTANCE"意义相同。
	// "INSTANCE"表示按照节点计算，默认方式。
	// "CPU"表示按照机器的核数计算。
	// "MEMORYGB"表示按照机器内存数计算。
	MeasureMethod *string `json:"MeasureMethod,omitnil,omitempty" name:"MeasureMethod"`

	// 销毁策略, "DEFAULT",默认销毁策略，由缩容规则触发缩容，"TIMING"表示定时销毁
	TerminatePolicy *string `json:"TerminatePolicy,omitnil,omitempty" name:"TerminatePolicy"`

	// 最长使用时间， 秒数，最短1小时，最长24小时
	MaxUse *int64 `json:"MaxUse,omitnil,omitempty" name:"MaxUse"`

	// 节点部署服务列表。部署服务仅填写HDFS、YARN。[组件名对应的映射关系表](https://cloud.tencent.com/document/product/589/98760)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SoftDeployInfo []*int64 `json:"SoftDeployInfo,omitnil,omitempty" name:"SoftDeployInfo"`

	// 启动进程列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceNodeInfo []*int64 `json:"ServiceNodeInfo,omitnil,omitempty" name:"ServiceNodeInfo"`

	// 补偿扩容，0表示不开启，1表示开启
	CompensateFlag *int64 `json:"CompensateFlag,omitnil,omitempty" name:"CompensateFlag"`

	// 伸缩组id
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type TopologyInfo struct {
	// 可用区ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 可用区信息
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 子网信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetInfoList []*SubnetInfo `json:"SubnetInfoList,omitnil,omitempty" name:"SubnetInfoList"`

	// 节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeInfoList []*ShortNodeInfo `json:"NodeInfoList,omitnil,omitempty" name:"NodeInfoList"`
}

type TriggerCondition struct {
	// 条件比较方法，1表示大于，2表示小于，3表示大于等于，4表示小于等于。
	CompareMethod *int64 `json:"CompareMethod,omitnil,omitempty" name:"CompareMethod"`

	// 条件阈值。
	Threshold *float64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`
}

type TrinoQueryInfo struct {
	// catalog
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// 提交IP
	ClientIpAddr *string `json:"ClientIpAddr,omitnil,omitempty" name:"ClientIpAddr"`

	// 切片数
	CompletedSplits *string `json:"CompletedSplits,omitnil,omitempty" name:"CompletedSplits"`

	// CPU时间
	CpuTime *int64 `json:"CpuTime,omitnil,omitempty" name:"CpuTime"`

	// 累计内存
	CumulativeMemory *int64 `json:"CumulativeMemory,omitnil,omitempty" name:"CumulativeMemory"`

	// 执行时长
	DurationMillis *int64 `json:"DurationMillis,omitnil,omitempty" name:"DurationMillis"`

	// 结束时间 (s)
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 内部传输量
	InternalNetworkBytes *int64 `json:"InternalNetworkBytes,omitnil,omitempty" name:"InternalNetworkBytes"`

	// 输出字节数
	OutputBytes *int64 `json:"OutputBytes,omitnil,omitempty" name:"OutputBytes"`

	// 峰值内存量
	PeakUserMemoryBytes *int64 `json:"PeakUserMemoryBytes,omitnil,omitempty" name:"PeakUserMemoryBytes"`

	// 物理输入量
	PhysicalInputBytes *int64 `json:"PhysicalInputBytes,omitnil,omitempty" name:"PhysicalInputBytes"`

	// 处理输入量
	ProcessedInputBytes *int64 `json:"ProcessedInputBytes,omitnil,omitempty" name:"ProcessedInputBytes"`

	// 编译时长
	SqlCompileTime *int64 `json:"SqlCompileTime,omitnil,omitempty" name:"SqlCompileTime"`

	// 开始时间 (s)
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 执行状态
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 执行语句
	Statement *string `json:"Statement,omitnil,omitempty" name:"Statement"`

	// 提交用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 写入字节数
	WrittenBytes *int64 `json:"WrittenBytes,omitnil,omitempty" name:"WrittenBytes"`
}

type UpdateInstanceSettings struct {
	// 内存容量，单位为G
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// CPU核数
	CPUCores *uint64 `json:"CPUCores,omitnil,omitempty" name:"CPUCores"`

	// 机器资源ID（EMR测资源标识）
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 变配机器规格
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type UserAndGroup struct {
	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户组
	UserGroup *string `json:"UserGroup,omitnil,omitempty" name:"UserGroup"`
}

type UserInfoForUserManager struct {
	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户所属的组
	UserGroup *string `json:"UserGroup,omitnil,omitempty" name:"UserGroup"`

	// 密码
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// 备注
	ReMark *string `json:"ReMark,omitnil,omitempty" name:"ReMark"`
}

type UserManagerFilter struct {
	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户来源
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`
}

type UserManagerUserBriefInfo struct {
	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户所属的组
	UserGroup *string `json:"UserGroup,omitnil,omitempty" name:"UserGroup"`

	// Manager表示管理员、NormalUser表示普通用户
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 用户创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否可以下载用户对应的keytab文件，对开启kerberos的集群才有意义
	SupportDownLoadKeyTab *bool `json:"SupportDownLoadKeyTab,omitnil,omitempty" name:"SupportDownLoadKeyTab"`

	// keytab文件的下载地址
	DownLoadKeyTabUrl *string `json:"DownLoadKeyTabUrl,omitnil,omitempty" name:"DownLoadKeyTabUrl"`
}

type VPCSettings struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type VirtualPrivateCloud struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type VolumeSetting struct {
	// 数据卷类型
	// <li>HOST_PATH表示支持本机路径</li>
	// <li>NEW_PVC表示新建PVC</li>
	// 组件角色支持的数据卷类型可参考 EMR on TKE 集群部署说明：[部署说明](https://cloud.tencent.com/document/product/589/94254)
	VolumeType *string `json:"VolumeType,omitnil,omitempty" name:"VolumeType"`

	// 主机路径信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostPath *HostPathVolumeSource `json:"HostPath,omitnil,omitempty" name:"HostPath"`
}

type WeekRepeatStrategy struct {
	// 重复任务执行的具体时刻，例如"01:02:00"
	ExecuteAtTimeOfDay *string `json:"ExecuteAtTimeOfDay,omitnil,omitempty" name:"ExecuteAtTimeOfDay"`

	// 每周几的数字描述，例如，[1,3,4]表示每周周一、周三、周四。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DaysOfWeek []*uint64 `json:"DaysOfWeek,omitnil,omitempty" name:"DaysOfWeek"`
}

type YarnApplication struct {
	// 应用ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 应用名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 队列
	Queue *string `json:"Queue,omitnil,omitempty" name:"Queue"`

	// 应用类型
	ApplicationType *string `json:"ApplicationType,omitnil,omitempty" name:"ApplicationType"`

	// 运行时间
	ElapsedTime *string `json:"ElapsedTime,omitnil,omitempty" name:"ElapsedTime"`

	// 状态
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 最终状态
	FinalStatus *string `json:"FinalStatus,omitnil,omitempty" name:"FinalStatus"`

	// 进度
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 开始时间毫秒
	StartedTime *int64 `json:"StartedTime,omitnil,omitempty" name:"StartedTime"`

	// 结束时间毫秒
	FinishedTime *int64 `json:"FinishedTime,omitnil,omitempty" name:"FinishedTime"`

	// 申请内存MB
	AllocatedMB *int64 `json:"AllocatedMB,omitnil,omitempty" name:"AllocatedMB"`

	// 申请VCores
	AllocatedVCores *int64 `json:"AllocatedVCores,omitnil,omitempty" name:"AllocatedVCores"`

	// 运行的Containers数
	RunningContainers *int64 `json:"RunningContainers,omitnil,omitempty" name:"RunningContainers"`

	// 内存MB*时间秒
	MemorySeconds *int64 `json:"MemorySeconds,omitnil,omitempty" name:"MemorySeconds"`

	// VCores*时间秒
	VCoreSeconds *int64 `json:"VCoreSeconds,omitnil,omitempty" name:"VCoreSeconds"`

	// 队列资源占比
	QueueUsagePercentage *float64 `json:"QueueUsagePercentage,omitnil,omitempty" name:"QueueUsagePercentage"`

	// 集群资源占比
	ClusterUsagePercentage *float64 `json:"ClusterUsagePercentage,omitnil,omitempty" name:"ClusterUsagePercentage"`

	// 预占用的内存
	PreemptedResourceMB *int64 `json:"PreemptedResourceMB,omitnil,omitempty" name:"PreemptedResourceMB"`

	// 预占用的VCore
	PreemptedResourceVCores *int64 `json:"PreemptedResourceVCores,omitnil,omitempty" name:"PreemptedResourceVCores"`

	// 预占的非应用程序主节点容器数量
	NumNonAMContainerPreempted *int64 `json:"NumNonAMContainerPreempted,omitnil,omitempty" name:"NumNonAMContainerPreempted"`

	// AM预占用的容器数量
	NumAMContainerPreempted *int64 `json:"NumAMContainerPreempted,omitnil,omitempty" name:"NumAMContainerPreempted"`

	// Map总数
	MapsTotal *int64 `json:"MapsTotal,omitnil,omitempty" name:"MapsTotal"`

	// 完成的Map数
	MapsCompleted *int64 `json:"MapsCompleted,omitnil,omitempty" name:"MapsCompleted"`

	// Reduce总数
	ReducesTotal *int64 `json:"ReducesTotal,omitnil,omitempty" name:"ReducesTotal"`

	// 完成的Reduce数
	ReducesCompleted *int64 `json:"ReducesCompleted,omitnil,omitempty" name:"ReducesCompleted"`

	// 平均Map时间
	AvgMapTime *int64 `json:"AvgMapTime,omitnil,omitempty" name:"AvgMapTime"`

	// 平均Reduce时间
	AvgReduceTime *int64 `json:"AvgReduceTime,omitnil,omitempty" name:"AvgReduceTime"`

	// 平均Shuffle时间毫秒
	AvgShuffleTime *int64 `json:"AvgShuffleTime,omitnil,omitempty" name:"AvgShuffleTime"`

	// 平均Merge时间毫秒
	AvgMergeTime *int64 `json:"AvgMergeTime,omitnil,omitempty" name:"AvgMergeTime"`

	// 失败的Reduce执行次数
	FailedReduceAttempts *int64 `json:"FailedReduceAttempts,omitnil,omitempty" name:"FailedReduceAttempts"`

	// Kill的Reduce执行次数
	KilledReduceAttempts *int64 `json:"KilledReduceAttempts,omitnil,omitempty" name:"KilledReduceAttempts"`

	// 成功的Reduce执行次数
	SuccessfulReduceAttempts *int64 `json:"SuccessfulReduceAttempts,omitnil,omitempty" name:"SuccessfulReduceAttempts"`

	// 失败的Map执行次数
	FailedMapAttempts *int64 `json:"FailedMapAttempts,omitnil,omitempty" name:"FailedMapAttempts"`

	// Kill的Map执行次数
	KilledMapAttempts *int64 `json:"KilledMapAttempts,omitnil,omitempty" name:"KilledMapAttempts"`

	// 成功的Map执行次数
	SuccessfulMapAttempts *int64 `json:"SuccessfulMapAttempts,omitnil,omitempty" name:"SuccessfulMapAttempts"`

	// GC毫秒
	GcTimeMillis *int64 `json:"GcTimeMillis,omitnil,omitempty" name:"GcTimeMillis"`

	// Map使用的VCore毫秒
	VCoreMillisMaps *int64 `json:"VCoreMillisMaps,omitnil,omitempty" name:"VCoreMillisMaps"`

	// Map使用的内存毫秒
	MbMillisMaps *int64 `json:"MbMillisMaps,omitnil,omitempty" name:"MbMillisMaps"`

	// Reduce使用的VCore毫秒
	VCoreMillisReduces *int64 `json:"VCoreMillisReduces,omitnil,omitempty" name:"VCoreMillisReduces"`

	// Reduce使用的内存毫秒
	MbMillisReduces *int64 `json:"MbMillisReduces,omitnil,omitempty" name:"MbMillisReduces"`

	// 启动Map的总数
	TotalLaunchedMaps *int64 `json:"TotalLaunchedMaps,omitnil,omitempty" name:"TotalLaunchedMaps"`

	// 启动Reduce的总数
	TotalLaunchedReduces *int64 `json:"TotalLaunchedReduces,omitnil,omitempty" name:"TotalLaunchedReduces"`

	// Map输入记录数
	MapInputRecords *int64 `json:"MapInputRecords,omitnil,omitempty" name:"MapInputRecords"`

	// Map输出记录数
	MapOutputRecords *int64 `json:"MapOutputRecords,omitnil,omitempty" name:"MapOutputRecords"`

	// Reduce输入记录数
	ReduceInputRecords *int64 `json:"ReduceInputRecords,omitnil,omitempty" name:"ReduceInputRecords"`

	// Reduce输出记录数
	ReduceOutputRecords *int64 `json:"ReduceOutputRecords,omitnil,omitempty" name:"ReduceOutputRecords"`

	// HDFS写入字节数
	HDFSBytesWritten *int64 `json:"HDFSBytesWritten,omitnil,omitempty" name:"HDFSBytesWritten"`

	// HDFS读取字节数
	HDFSBytesRead *int64 `json:"HDFSBytesRead,omitnil,omitempty" name:"HDFSBytesRead"`
}

type ZoneDetailPriceResult struct {
	// 可用区Id
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 不同节点的价格详情
	NodeDetailPrice []*NodeDetailPriceResult `json:"NodeDetailPrice,omitnil,omitempty" name:"NodeDetailPrice"`
}

type ZoneResourceConfiguration struct {
	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil,omitempty" name:"VirtualPrivateCloud"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 所有节点资源的规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllNodeResourceSpec *AllNodeResourceSpec `json:"AllNodeResourceSpec,omitnil,omitempty" name:"AllNodeResourceSpec"`

	// 如果是单可用区，ZoneTag可以不用填， 如果是双AZ部署，第一个可用区ZoneTag选择master，第二个可用区ZoneTag选择standby，如果是三AZ部署，第一个可用区ZoneTag选择master，第二个可用区ZoneTag选择standby，第三个可用区ZoneTag选择third-party，取值范围：
	//   <li>master</li>
	//   <li>standby</li>
	//   <li>third-party</li>
	ZoneTag *string `json:"ZoneTag,omitnil,omitempty" name:"ZoneTag"`
}

type ZoneSetting struct {
	// 可用区名称
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 可用区VPC和子网
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil,omitempty" name:"VPCSettings"`

	// 可用区节点数量
	NodeNum *int64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`
}