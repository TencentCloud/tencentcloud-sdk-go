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

package v20170312

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Activity struct {

	// 活动ID
	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

	// 计算节点ID
	ComputeNodeId *string `json:"ComputeNodeId,omitempty" name:"ComputeNodeId"`

	// 计算节点活动类型，创建或者销毁
	ComputeNodeActivityType *string `json:"ComputeNodeActivityType,omitempty" name:"ComputeNodeActivityType"`

	// 计算环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 起因
	Cause *string `json:"Cause,omitempty" name:"Cause"`

	// 活动状态
	ActivityState *string `json:"ActivityState,omitempty" name:"ActivityState"`

	// 状态原因
	StateReason *string `json:"StateReason,omitempty" name:"StateReason"`

	// 活动开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 活动结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 云服务器实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type AgentRunningMode struct {

	// 场景类型，支持WINDOWS
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// 运行Agent的User
	User *string `json:"User,omitempty" name:"User"`

	// 运行Agent的Session
	Session *string `json:"Session,omitempty" name:"Session"`
}

type AnonymousComputeEnv struct {

	// 计算环境管理类型
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// 计算环境具体参数
	EnvData *EnvData `json:"EnvData,omitempty" name:"EnvData"`

	// 数据盘挂载选项
	MountDataDisks []*MountDataDisk `json:"MountDataDisks,omitempty" name:"MountDataDisks" list`

	// agent运行模式，适用于Windows系统
	AgentRunningMode *AgentRunningMode `json:"AgentRunningMode,omitempty" name:"AgentRunningMode"`
}

type Application struct {

	// 任务执行命令
	Command *string `json:"Command,omitempty" name:"Command"`

	// 应用程序的交付方式，包括PACKAGE、LOCAL 两种取值，分别指远程存储的软件包、计算环境本地。
	DeliveryForm *string `json:"DeliveryForm,omitempty" name:"DeliveryForm"`

	// 应用程序软件包的远程存储路径
	PackagePath *string `json:"PackagePath,omitempty" name:"PackagePath"`

	// 应用使用Docker的相关配置。在使用Docker配置的情况下，DeliveryForm 为 LOCAL 表示直接使用Docker镜像内部的应用软件，通过Docker方式运行；DeliveryForm 为 PACKAGE，表示将远程应用包注入到Docker镜像后，通过Docker方式运行。为避免Docker不同版本的兼容性问题，Docker安装包及相关依赖由Batch统一负责，对于已安装Docker的自定义镜像，请卸载后再使用Docker特性。
	Docker *Docker `json:"Docker,omitempty" name:"Docker"`
}

type Authentication struct {

	// 授权场景，例如COS
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// SecretId
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// SecretKey
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
}

type ComputeEnvCreateInfo struct {

	// 计算环境 ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 计算环境名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`

	// 计算环境描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvDescription *string `json:"EnvDescription,omitempty" name:"EnvDescription"`

	// 计算环境类型，仅支持“MANAGED”类型
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// 计算环境参数
	EnvData *EnvData `json:"EnvData,omitempty" name:"EnvData"`

	// 数据盘挂载选项
	// 注意：此字段可能返回 null，表示取不到有效值。
	MountDataDisks []*MountDataDisk `json:"MountDataDisks,omitempty" name:"MountDataDisks" list`

	// 输入映射
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputMappings []*InputMapping `json:"InputMappings,omitempty" name:"InputMappings" list`

	// 授权信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Authentications []*Authentication `json:"Authentications,omitempty" name:"Authentications" list`

	// 通知信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Notifications []*Notification `json:"Notifications,omitempty" name:"Notifications" list`

	// 计算节点期望个数
	DesiredComputeNodeCount *uint64 `json:"DesiredComputeNodeCount,omitempty" name:"DesiredComputeNodeCount"`
}

type ComputeEnvData struct {

	// CVM实例类型列表
	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes" list`
}

type ComputeEnvView struct {

	// 计算环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 计算环境名称
	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`

	// 位置信息
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 计算节点统计指标
	ComputeNodeMetrics *ComputeNodeMetrics `json:"ComputeNodeMetrics,omitempty" name:"ComputeNodeMetrics"`

	// 计算环境类型
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// 计算节点期望个数
	DesiredComputeNodeCount *uint64 `json:"DesiredComputeNodeCount,omitempty" name:"DesiredComputeNodeCount"`
}

type ComputeNode struct {

	// 计算节点ID
	ComputeNodeId *string `json:"ComputeNodeId,omitempty" name:"ComputeNodeId"`

	// 计算节点实例ID，对于CVM场景，即为CVM的InstanceId
	ComputeNodeInstanceId *string `json:"ComputeNodeInstanceId,omitempty" name:"ComputeNodeInstanceId"`

	// 计算节点状态
	ComputeNodeState *string `json:"ComputeNodeState,omitempty" name:"ComputeNodeState"`

	// CPU核数
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存容量，单位GiB
	Mem *uint64 `json:"Mem,omitempty" name:"Mem"`

	// 资源创建完成时间
	ResourceCreatedTime *string `json:"ResourceCreatedTime,omitempty" name:"ResourceCreatedTime"`

	// 计算节点运行  TaskInstance 可用容量。0表示计算节点忙碌。
	TaskInstanceNumAvailable *uint64 `json:"TaskInstanceNumAvailable,omitempty" name:"TaskInstanceNumAvailable"`

	// Batch Agent 版本
	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`

	// 实例内网IP
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`

	// 实例公网IP
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses" list`
}

type ComputeNodeMetrics struct {

	// 已经完成提交的计算节点数量
	SubmittedCount *string `json:"SubmittedCount,omitempty" name:"SubmittedCount"`

	// 创建中的计算节点数量
	CreatingCount *string `json:"CreatingCount,omitempty" name:"CreatingCount"`

	// 创建失败的计算节点数量
	CreationFailedCount *string `json:"CreationFailedCount,omitempty" name:"CreationFailedCount"`

	// 完成创建的计算节点数量
	CreatedCount *string `json:"CreatedCount,omitempty" name:"CreatedCount"`

	// 运行中的计算节点数量
	RunningCount *string `json:"RunningCount,omitempty" name:"RunningCount"`

	// 销毁中的计算节点数量
	DeletingCount *string `json:"DeletingCount,omitempty" name:"DeletingCount"`

	// 异常的计算节点数量
	AbnormalCount *string `json:"AbnormalCount,omitempty" name:"AbnormalCount"`
}

type CreateComputeEnvRequest struct {
	*tchttp.BaseRequest

	// 计算环境信息
	ComputeEnv *NamedComputeEnv `json:"ComputeEnv,omitempty" name:"ComputeEnv"`

	// 位置信息
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 用于保证请求幂等性的字符串。该字符串由用户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
}

func (r *CreateComputeEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateComputeEnvRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateComputeEnvResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 计算环境ID
		EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateComputeEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateComputeEnvResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTaskTemplateRequest struct {
	*tchttp.BaseRequest

	// 任务模板名称
	TaskTemplateName *string `json:"TaskTemplateName,omitempty" name:"TaskTemplateName"`

	// 任务模板内容，参数要求与任务一致
	TaskTemplateInfo *Task `json:"TaskTemplateInfo,omitempty" name:"TaskTemplateInfo"`

	// 任务模板描述
	TaskTemplateDescription *string `json:"TaskTemplateDescription,omitempty" name:"TaskTemplateDescription"`
}

func (r *CreateTaskTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTaskTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTaskTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务模板ID
		TaskTemplateId *string `json:"TaskTemplateId,omitempty" name:"TaskTemplateId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTaskTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTaskTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {

	// 数据盘大小，单位：GB。最小调整步长为10G，不同数据盘类型取值范围不同，具体限制详见：[CVM实例配置](/document/product/213/2177)。默认值为0，表示不购买数据盘。更多限制详见产品文档。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 数据盘类型。数据盘类型限制详见[CVM实例配置](/document/product/213/2177)。取值范围：<br><li>LOCAL_BASIC：本地硬盘<br><li>LOCAL_SSD：本地SSD硬盘<br><li>CLOUD_BASIC：普通云硬盘<br><li>CLOUD_PREMIUM：高性能云硬盘<br><li>CLOUD_SSD：SSD云硬盘<br><br>默认取值：LOCAL_BASIC。<br><br>该参数对`ResizeInstanceDisk`接口无效。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 数据盘ID。LOCAL_BASIC 和 LOCAL_SSD 类型没有ID。暂时不支持该参数。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 数据盘是否随子机销毁。取值范围：
	// <li>TRUE：子机销毁时，销毁数据盘，只支持按小时后付费云盘
	// <li>FALSE：子机销毁时，保留数据盘<br>
	// 默认取值：TRUE<br>
	// 该参数目前仅用于 `RunInstances` 接口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
}

type DeleteComputeEnvRequest struct {
	*tchttp.BaseRequest

	// 计算环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *DeleteComputeEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteComputeEnvRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteComputeEnvResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteComputeEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteComputeEnvResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteJobRequest struct {
	*tchttp.BaseRequest

	// 作业ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DeleteJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTaskTemplatesRequest struct {
	*tchttp.BaseRequest

	// 用于删除任务模板信息
	TaskTemplateIds []*string `json:"TaskTemplateIds,omitempty" name:"TaskTemplateIds" list`
}

func (r *DeleteTaskTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTaskTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTaskTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTaskTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTaskTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Dependence struct {

	// 依赖关系的起点任务名称
	StartTask *string `json:"StartTask,omitempty" name:"StartTask"`

	// 依赖关系的终点任务名称
	EndTask *string `json:"EndTask,omitempty" name:"EndTask"`
}

type DescribeAvailableCvmInstanceTypesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li> zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。</li>
	// <li> instance-family String - 是否必填：否 -（过滤条件）按照机型系列过滤。实例机型系列形如：S1、I1、M1等。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeAvailableCvmInstanceTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAvailableCvmInstanceTypesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableCvmInstanceTypesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机型配置数组
		InstanceTypeConfigSet []*InstanceTypeConfig `json:"InstanceTypeConfigSet,omitempty" name:"InstanceTypeConfigSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableCvmInstanceTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAvailableCvmInstanceTypesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeComputeEnvActivitiesRequest struct {
	*tchttp.BaseRequest

	// 计算环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件
	// <li> compute-node-id - String - 是否必填：否 -（过滤条件）按照计算节点ID过滤。</li>
	Filters *Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComputeEnvActivitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeComputeEnvActivitiesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeComputeEnvActivitiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 计算环境中的活动列表
		ActivitySet []*Activity `json:"ActivitySet,omitempty" name:"ActivitySet" list`

		// 活动数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComputeEnvActivitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeComputeEnvActivitiesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeComputeEnvCreateInfoRequest struct {
	*tchttp.BaseRequest

	// 计算环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *DescribeComputeEnvCreateInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeComputeEnvCreateInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeComputeEnvCreateInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 计算环境 ID
		EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

		// 计算环境名称
		EnvName *string `json:"EnvName,omitempty" name:"EnvName"`

		// 计算环境描述
	// 注意：此字段可能返回 null，表示取不到有效值。
		EnvDescription *string `json:"EnvDescription,omitempty" name:"EnvDescription"`

		// 计算环境类型，仅支持“MANAGED”类型
		EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

		// 计算环境参数
		EnvData *EnvData `json:"EnvData,omitempty" name:"EnvData"`

		// 数据盘挂载选项
		MountDataDisks []*MountDataDisk `json:"MountDataDisks,omitempty" name:"MountDataDisks" list`

		// 输入映射
		InputMappings []*InputMapping `json:"InputMappings,omitempty" name:"InputMappings" list`

		// 授权信息
		Authentications []*Authentication `json:"Authentications,omitempty" name:"Authentications" list`

		// 通知信息
		Notifications []*Notification `json:"Notifications,omitempty" name:"Notifications" list`

		// 计算节点期望个数
		DesiredComputeNodeCount *int64 `json:"DesiredComputeNodeCount,omitempty" name:"DesiredComputeNodeCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComputeEnvCreateInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeComputeEnvCreateInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeComputeEnvCreateInfosRequest struct {
	*tchttp.BaseRequest

	// 计算环境ID
	EnvIds []*string `json:"EnvIds,omitempty" name:"EnvIds" list`

	// 过滤条件
	// <li> zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。</li>
	// <li> env-id - String - 是否必填：否 -（过滤条件）按照计算环境ID过滤。</li>
	// <li> env-name - String - 是否必填：否 -（过滤条件）按照计算环境名称过滤。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeComputeEnvCreateInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeComputeEnvCreateInfosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeComputeEnvCreateInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 计算环境数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 计算环境创建信息列表
		ComputeEnvCreateInfoSet []*ComputeEnvCreateInfo `json:"ComputeEnvCreateInfoSet,omitempty" name:"ComputeEnvCreateInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComputeEnvCreateInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeComputeEnvCreateInfosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeComputeEnvRequest struct {
	*tchttp.BaseRequest

	// 计算环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *DescribeComputeEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeComputeEnvRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeComputeEnvResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 计算环境ID
		EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

		// 计算环境名称
		EnvName *string `json:"EnvName,omitempty" name:"EnvName"`

		// 位置信息
		Placement *Placement `json:"Placement,omitempty" name:"Placement"`

		// 计算环境创建时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 计算节点列表信息
		ComputeNodeSet []*ComputeNode `json:"ComputeNodeSet,omitempty" name:"ComputeNodeSet" list`

		// 计算节点统计指标
		ComputeNodeMetrics *ComputeNodeMetrics `json:"ComputeNodeMetrics,omitempty" name:"ComputeNodeMetrics"`

		// 计算节点期望个数
		DesiredComputeNodeCount *uint64 `json:"DesiredComputeNodeCount,omitempty" name:"DesiredComputeNodeCount"`

		// 计算环境类型
		EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComputeEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeComputeEnvResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeComputeEnvsRequest struct {
	*tchttp.BaseRequest

	// 计算环境ID
	EnvIds []*string `json:"EnvIds,omitempty" name:"EnvIds" list`

	// 过滤条件
	// <li> zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。</li>
	// <li> env-id - String - 是否必填：否 -（过滤条件）按照计算环境ID过滤。</li>
	// <li> env-name - String - 是否必填：否 -（过滤条件）按照计算环境名称过滤。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeComputeEnvsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeComputeEnvsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeComputeEnvsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 计算环境列表
		ComputeEnvSet []*ComputeEnvView `json:"ComputeEnvSet,omitempty" name:"ComputeEnvSet" list`

		// 计算环境数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComputeEnvsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeComputeEnvsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCvmZoneInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li> zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。</li>
	// <li> instance-family String - 是否必填：否 -（过滤条件）按照机型系列过滤。实例机型系列形如：S1、I1、M1等。</li>
	// <li> instance-type - String - 是否必填：否 - （过滤条件）按照机型过滤。</li>
	// <li> instance-charge-type - String - 是否必填：否 -（过滤条件）按照实例计费模式过滤。 ( POSTPAID_BY_HOUR：表示后付费，即按量计费机型 | SPOTPAID：表示竞价付费机型。 )  </li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeCvmZoneInstanceConfigInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCvmZoneInstanceConfigInfosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCvmZoneInstanceConfigInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可用区机型配置列表。
		InstanceTypeQuotaSet []*InstanceTypeQuotaItem `json:"InstanceTypeQuotaSet,omitempty" name:"InstanceTypeQuotaSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCvmZoneInstanceConfigInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCvmZoneInstanceConfigInfosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeJobRequest struct {
	*tchttp.BaseRequest

	// 作业标识
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 作业ID
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 作业名称
		JobName *string `json:"JobName,omitempty" name:"JobName"`

		// 可用区信息
		Zone *string `json:"Zone,omitempty" name:"Zone"`

		// 作业优先级
		Priority *int64 `json:"Priority,omitempty" name:"Priority"`

		// 作业状态
		JobState *string `json:"JobState,omitempty" name:"JobState"`

		// 创建时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 结束时间
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 任务视图信息
		TaskSet []*TaskView `json:"TaskSet,omitempty" name:"TaskSet" list`

		// 任务间依赖信息
		DependenceSet []*Dependence `json:"DependenceSet,omitempty" name:"DependenceSet" list`

		// 任务统计指标
		TaskMetrics *TaskMetrics `json:"TaskMetrics,omitempty" name:"TaskMetrics"`

		// 任务实例统计指标
		TaskInstanceMetrics *TaskInstanceView `json:"TaskInstanceMetrics,omitempty" name:"TaskInstanceMetrics"`

		// 作业失败原因
		StateReason *string `json:"StateReason,omitempty" name:"StateReason"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeJobSubmitInfoRequest struct {
	*tchttp.BaseRequest

	// 作业ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeJobSubmitInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeJobSubmitInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeJobSubmitInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 作业ID
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 作业名称
		JobName *string `json:"JobName,omitempty" name:"JobName"`

		// 作业描述
		JobDescription *string `json:"JobDescription,omitempty" name:"JobDescription"`

		// 作业优先级，任务（Task）和任务实例（TaskInstance）会继承作业优先级
		Priority *int64 `json:"Priority,omitempty" name:"Priority"`

		// 任务信息
		Tasks []*Task `json:"Tasks,omitempty" name:"Tasks" list`

		// 依赖信息
		Dependences []*Dependence `json:"Dependences,omitempty" name:"Dependences" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJobSubmitInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeJobSubmitInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeJobsRequest struct {
	*tchttp.BaseRequest

	// 作业ID
	JobIds []*string `json:"JobIds,omitempty" name:"JobIds" list`

	// 过滤条件
	// <li> job-id - String - 是否必填：否 -（过滤条件）按照作业ID过滤。</li>
	// <li> job-name - String - 是否必填：否 -（过滤条件）按照作业名称过滤。</li>
	// <li> job-state - String - 是否必填：否 -（过滤条件）按照作业状态过滤。</li>
	// <li> zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeJobsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeJobsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 作业列表
		JobSet []*JobView `json:"JobSet,omitempty" name:"JobSet" list`

		// 符合条件的作业数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeJobsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskLogsRequest struct {
	*tchttp.BaseRequest

	// 作业ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务实例集合
	TaskInstanceIndexes []*uint64 `json:"TaskInstanceIndexes,omitempty" name:"TaskInstanceIndexes" list`

	// 起始任务实例
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 最大任务实例数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTaskLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务实例总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 任务实例日志详情集合
		TaskInstanceLogSet []*TaskInstanceLog `json:"TaskInstanceLogSet,omitempty" name:"TaskInstanceLogSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskRequest struct {
	*tchttp.BaseRequest

	// 作业ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量。默认取值100，最大取值1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，详情如下：
	// <li> task-instance-type - String - 是否必填： 否 - 按照任务实例状态进行过滤（SUBMITTED：已提交；PENDING：等待中；RUNNABLE：可运行；STARTING：启动中；RUNNING：运行中；SUCCEED：成功；FAILED：失败；FAILED_INTERRUPTED：失败后保留实例）。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 作业ID
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 任务名称
		TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

		// 任务状态
		TaskState *string `json:"TaskState,omitempty" name:"TaskState"`

		// 创建时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 结束时间
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 任务实例总数
		TaskInstanceTotalCount *int64 `json:"TaskInstanceTotalCount,omitempty" name:"TaskInstanceTotalCount"`

		// 任务实例信息
		TaskInstanceSet []*TaskInstanceView `json:"TaskInstanceSet,omitempty" name:"TaskInstanceSet" list`

		// 任务实例统计指标
		TaskInstanceMetrics *TaskInstanceMetrics `json:"TaskInstanceMetrics,omitempty" name:"TaskInstanceMetrics"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskTemplatesRequest struct {
	*tchttp.BaseRequest

	// 任务模板ID
	TaskTemplateIds []*string `json:"TaskTemplateIds,omitempty" name:"TaskTemplateIds" list`

	// 过滤条件
	// <li> task-template-name - String - 是否必填：否 -（过滤条件）按照任务模板名称过滤。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTaskTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务模板列表
		TaskTemplateSet []*TaskTemplateView `json:"TaskTemplateSet,omitempty" name:"TaskTemplateSet" list`

		// 任务模板数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Docker struct {

	// Docker Hub 用户名或 Tencent Registry 用户名
	User *string `json:"User,omitempty" name:"User"`

	// Docker Hub 密码或 Tencent Registry 密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// Docker Hub填写“[user/repo]:[tag]”，Tencent Registry填写“ccr.ccs.tencentyun.com/[namespace/repo]:[tag]”
	Image *string `json:"Image,omitempty" name:"Image"`

	// Docker Hub 可以不填，但确保具有公网访问能力。或者是 Tencent Registry 服务地址“ccr.ccs.tencentyun.com”
	Server *string `json:"Server,omitempty" name:"Server"`
}

type EnhancedService struct {

	// 开启云安全服务。若不指定该参数，则默认开启云安全服务。
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`

	// 开启云监控服务。若不指定该参数，则默认开启云监控服务。
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`
}

type EnvData struct {

	// CVM实例类型，不能与InstanceTypes和InstanceTypeOptions同时出现。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// CVM镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 实例系统盘配置信息
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// 私有网络相关信息配置
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// 公网带宽相关信息设置
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// CVM实例显示名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例登录设置
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 实例所属安全组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// CVM实例计费类型<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>SPOTPAID：竞价付费<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 实例的市场相关选项，如竞价实例相关参数
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitempty" name:"InstanceMarketOptions"`

	// CVM实例类型列表，不能与InstanceType和InstanceTypeOptions同时出现。指定该字段后，计算节点按照机型先后顺序依次尝试创建，直到实例创建成功，结束遍历过程。最多支持10个机型。
	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes" list`

	// CVM实例机型配置。不能与InstanceType和InstanceTypes同时出现。
	InstanceTypeOptions *InstanceTypeOptions `json:"InstanceTypeOptions,omitempty" name:"InstanceTypeOptions"`
}

type EnvVar struct {

	// 环境变量名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 环境变量取值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type EventConfig struct {

	// 事件类型，包括：<br/><li>“JOB_RUNNING”：作业运行，适用于"SubmitJob"。</li><li>“JOB_SUCCEED”：作业成功，适用于"SubmitJob"。</li><li>“JOB_FAILED”：作业失败，适用于"SubmitJob"。</li><li>“JOB_FAILED_INTERRUPTED”：作业失败，保留实例，适用于"SubmitJob"。</li><li>“TASK_RUNNING”：任务运行，适用于"SubmitJob"。</li><li>“TASK_SUCCEED”：任务成功，适用于"SubmitJob"。</li><li>“TASK_FAILED”：任务失败，适用于"SubmitJob"。</li><li>“TASK_FAILED_INTERRUPTED”：任务失败，保留实例，适用于"SubmitJob"。</li><li>“TASK_INSTANCE_RUNNING”：任务实例运行，适用于"SubmitJob"。</li><li>“TASK_INSTANCE_SUCCEED”：任务实例成功，适用于"SubmitJob"。</li><li>“TASK_INSTANCE_FAILED”：任务实例失败，适用于"SubmitJob"。</li><li>“TASK_INSTANCE_FAILED_INTERRUPTED”：任务实例失败，保留实例，适用于"SubmitJob"。</li><li>“COMPUTE_ENV_CREATED”：计算环境已创建，适用于"CreateComputeEnv"。</li><li>“COMPUTE_ENV_DELETED”：计算环境已删除，适用于"CreateComputeEnv"。</li><li>“COMPUTE_NODE_CREATED”：计算节点已创建，适用于"CreateComputeEnv"和"SubmitJob"。</li><li>“COMPUTE_NODE_CREATION_FAILED”：计算节点创建失败，适用于"CreateComputeEnv"和"SubmitJob"。</li><li>“COMPUTE_NODE_RUNNING”：计算节点运行中，适用于"CreateComputeEnv"和"SubmitJob"。</li><li>“COMPUTE_NODE_ABNORMAL”：计算节点异常，适用于"CreateComputeEnv"和"SubmitJob"。</li><li>“COMPUTE_NODE_DELETING”：计算节点已删除，适用于"CreateComputeEnv"和"SubmitJob"。</li>
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// 自定义键值对
	EventVars []*EventVar `json:"EventVars,omitempty" name:"EventVars" list`
}

type EventVar struct {

	// 自定义键
	Name *string `json:"Name,omitempty" name:"Name"`

	// 自定义值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Externals struct {

	// 释放地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseAddress *bool `json:"ReleaseAddress,omitempty" name:"ReleaseAddress"`

	// 不支持的网络类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnsupportNetworks []*string `json:"UnsupportNetworks,omitempty" name:"UnsupportNetworks" list`

	// HDD本地存储属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageBlockAttr *StorageBlock `json:"StorageBlockAttr,omitempty" name:"StorageBlockAttr"`
}

type Filter struct {

	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type InputMapping struct {

	// 源端路径
	SourcePath *string `json:"SourcePath,omitempty" name:"SourcePath"`

	// 目的端路径
	DestinationPath *string `json:"DestinationPath,omitempty" name:"DestinationPath"`

	// 挂载配置项参数
	MountOptionParameter *string `json:"MountOptionParameter,omitempty" name:"MountOptionParameter"`
}

type InstanceMarketOptionsRequest struct {
	*tchttp.BaseRequest

	// 竞价相关选项
	SpotOptions *SpotMarketOptions `json:"SpotOptions,omitempty" name:"SpotOptions"`

	// 市场选项类型，当前只支持取值：spot
	MarketType *string `json:"MarketType,omitempty" name:"MarketType"`
}

func (r *InstanceMarketOptionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InstanceMarketOptionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstanceTypeConfig struct {

	// 内存容量，单位：`GB`。
	Mem *int64 `json:"Mem,omitempty" name:"Mem"`

	// CPU核数，单位：核。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 实例机型。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例机型系列。
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
}

type InstanceTypeOptions struct {

	// CPU核数。
	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`

	// 内存值，单位GB。
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 实例机型类别，可选参数：“ALL”、“GENERAL_2”、“GENERAL_3”、“COMPUTE_2”和“COMPUTE_3”。默认值“ALL”。
	InstanceCategories []*string `json:"InstanceCategories,omitempty" name:"InstanceCategories" list`
}

type InstanceTypeQuotaItem struct {

	// 可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例机型。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例计费模式。取值范围： <br><li>PREPAID：表示预付费，即包年包月<br><li>POSTPAID_BY_HOUR：表示后付费，即按量计费<br><li>CDHPAID：表示[CDH](https://cloud.tencent.com/document/product/416)付费，即只对CDH计费，不对CDH上的实例计费。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 网卡类型，例如：25代表25G网卡
	NetworkCard *int64 `json:"NetworkCard,omitempty" name:"NetworkCard"`

	// 扩展属性。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Externals *Externals `json:"Externals,omitempty" name:"Externals"`

	// 实例的CPU核数，单位：核。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 实例内存容量，单位：`GB`。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例机型系列。
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`

	// 机型名称。
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// 本地磁盘规格列表。
	LocalDiskTypeList []*LocalDiskType `json:"LocalDiskTypeList,omitempty" name:"LocalDiskTypeList" list`

	// 实例是否售卖。取值范围： <br><li>SELL：表示实例可购买<br><li>SOLD_OUT：表示实例已售罄。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例的售卖价格。
	Price *ItemPrice `json:"Price,omitempty" name:"Price"`
}

type InternetAccessible struct {

	// 网络计费类型。取值范围：<br><li>BANDWIDTH_PREPAID：预付费按带宽结算<br><li>TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费<br><li>BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费<br><li>BANDWIDTH_PACKAGE：带宽包用户<br>默认取值：非带宽包用户默认与子机付费类型保持一致。
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// 公网出带宽上限，单位：Mbps。默认值：0Mbps。不同机型带宽上限范围不一致，具体限制详见[购买网络带宽](/document/product/213/509)。
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// 是否分配公网IP。取值范围：<br><li>TRUE：表示分配公网IP<br><li>FALSE：表示不分配公网IP<br><br>当公网带宽大于0Mbps时，可自由选择开通与否，默认开通公网IP；当公网带宽为0，则不允许分配公网IP。
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty" name:"PublicIpAssigned"`

	// 带宽包ID。可通过[`DescribeBandwidthPackages`](https://cloud.tencent.com/document/api/215/19209)接口返回值中的`BandwidthPackageId`获取。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
}

type ItemPrice struct {

	// 后续单价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// 后续计价单元，可取值范围： <br><li>HOUR：表示计价单元是按每小时来计算。当前涉及该计价单元的场景有：实例按小时后付费（POSTPAID_BY_HOUR）、带宽按小时后付费（BANDWIDTH_POSTPAID_BY_HOUR）：<br><li>GB：表示计价单元是按每GB来计算。当前涉及该计价单元的场景有：流量按小时后付费（TRAFFIC_POSTPAID_BY_HOUR）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`

	// 预支费用的原价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 预支费用的折扣价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
}

type Job struct {

	// 任务信息
	Tasks []*Task `json:"Tasks,omitempty" name:"Tasks" list`

	// 作业名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 作业描述
	JobDescription *string `json:"JobDescription,omitempty" name:"JobDescription"`

	// 作业优先级，任务（Task）和任务实例（TaskInstance）会继承作业优先级
	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`

	// 依赖信息
	Dependences []*Dependence `json:"Dependences,omitempty" name:"Dependences" list`

	// 通知信息
	Notifications []*Notification `json:"Notifications,omitempty" name:"Notifications" list`

	// 对于存在依赖关系的任务中，后序任务执行对于前序任务的依赖条件。取值范围包括 PRE_TASK_SUCCEED，PRE_TASK_AT_LEAST_PARTLY_SUCCEED，PRE_TASK_FINISHED，默认值为PRE_TASK_SUCCEED。
	TaskExecutionDependOn *string `json:"TaskExecutionDependOn,omitempty" name:"TaskExecutionDependOn"`

	// 表示创建 CVM 失败按照何种策略处理。取值范围包括 FAILED，RUNNABLE。FAILED 表示创建 CVM 失败按照一次执行失败处理，RUNNABLE 表示创建 CVM 失败按照继续等待处理。默认值为FAILED。StateIfCreateCvmFailed对于提交的指定计算环境的作业无效。
	StateIfCreateCvmFailed *string `json:"StateIfCreateCvmFailed,omitempty" name:"StateIfCreateCvmFailed"`
}

type JobView struct {

	// 作业ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 作业名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 作业状态
	JobState *string `json:"JobState,omitempty" name:"JobState"`

	// 作业优先级
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 位置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务统计指标
	TaskMetrics *TaskMetrics `json:"TaskMetrics,omitempty" name:"TaskMetrics"`
}

type LocalDiskType struct {

	// 本地磁盘类型。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 本地磁盘属性。
	PartitionType *string `json:"PartitionType,omitempty" name:"PartitionType"`

	// 本地磁盘最小值。
	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`

	// 本地磁盘最大值。
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

type LoginSettings struct {

	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>Linux实例密码必须8到16位，至少包括两项[a-z，A-Z]、[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? / ]中的特殊符号。<br><li>Windows实例密码必须12到16位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = { } [ ] : ; ' , . ? /]中的特殊符号。<br><br>若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`

	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：<br><li>TRUE：表示保持镜像的登录设置<br><li>FALSE：表示不保持镜像的登录设置<br><br>默认取值：FALSE。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

type ModifyComputeEnvRequest struct {
	*tchttp.BaseRequest

	// 计算环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 计算节点期望个数
	DesiredComputeNodeCount *int64 `json:"DesiredComputeNodeCount,omitempty" name:"DesiredComputeNodeCount"`

	// 计算环境名称
	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`

	// 计算环境描述
	EnvDescription *string `json:"EnvDescription,omitempty" name:"EnvDescription"`

	// 计算环境属性数据
	EnvData *ComputeEnvData `json:"EnvData,omitempty" name:"EnvData"`
}

func (r *ModifyComputeEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyComputeEnvRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyComputeEnvResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyComputeEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyComputeEnvResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTaskTemplateRequest struct {
	*tchttp.BaseRequest

	// 任务模板ID
	TaskTemplateId *string `json:"TaskTemplateId,omitempty" name:"TaskTemplateId"`

	// 任务模板名称
	TaskTemplateName *string `json:"TaskTemplateName,omitempty" name:"TaskTemplateName"`

	// 任务模板描述
	TaskTemplateDescription *string `json:"TaskTemplateDescription,omitempty" name:"TaskTemplateDescription"`

	// 任务模板信息
	TaskTemplateInfo *Task `json:"TaskTemplateInfo,omitempty" name:"TaskTemplateInfo"`
}

func (r *ModifyTaskTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTaskTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTaskTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTaskTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTaskTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MountDataDisk struct {

	// 挂载点，Linux系统合法路径，或Windows系统盘符,比如"H:\\"
	LocalPath *string `json:"LocalPath,omitempty" name:"LocalPath"`

	// 文件系统类型，Linux系统下支持"EXT3"和"EXT4"两种，默认"EXT3"；Windows系统下仅支持"NTFS"
	FileSystemType *string `json:"FileSystemType,omitempty" name:"FileSystemType"`
}

type NamedComputeEnv struct {

	// 计算环境名称
	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`

	// 计算节点期望个数
	DesiredComputeNodeCount *int64 `json:"DesiredComputeNodeCount,omitempty" name:"DesiredComputeNodeCount"`

	// 计算环境描述
	EnvDescription *string `json:"EnvDescription,omitempty" name:"EnvDescription"`

	// 计算环境管理类型
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// 计算环境具体参数
	EnvData *EnvData `json:"EnvData,omitempty" name:"EnvData"`

	// 数据盘挂载选项
	MountDataDisks []*MountDataDisk `json:"MountDataDisks,omitempty" name:"MountDataDisks" list`

	// 授权信息
	Authentications []*Authentication `json:"Authentications,omitempty" name:"Authentications" list`

	// 输入映射信息
	InputMappings []*InputMapping `json:"InputMappings,omitempty" name:"InputMappings" list`

	// agent运行模式，适用于Windows系统
	AgentRunningMode *AgentRunningMode `json:"AgentRunningMode,omitempty" name:"AgentRunningMode"`

	// 通知信息
	Notifications *Notification `json:"Notifications,omitempty" name:"Notifications"`

	// 非活跃节点处理策略，默认“RECREATE”，即对于实例创建失败或异常退还的计算节点，定期重新创建实例资源。
	ActionIfComputeNodeInactive *string `json:"ActionIfComputeNodeInactive,omitempty" name:"ActionIfComputeNodeInactive"`
}

type Notification struct {

	// CMQ主题名字，要求主题名有效且关联订阅
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 事件配置
	EventConfigs []*EventConfig `json:"EventConfigs,omitempty" name:"EventConfigs" list`
}

type OutputMapping struct {

	// 源端路径
	SourcePath *string `json:"SourcePath,omitempty" name:"SourcePath"`

	// 目的端路径
	DestinationPath *string `json:"DestinationPath,omitempty" name:"DestinationPath"`
}

type OutputMappingConfig struct {

	// 存储类型，仅支持COS
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// 并行worker数量
	WorkerNum *int64 `json:"WorkerNum,omitempty" name:"WorkerNum"`

	// worker分块大小
	WorkerPartSize *int64 `json:"WorkerPartSize,omitempty" name:"WorkerPartSize"`
}

type Placement struct {

	// 实例所属的[可用区](/document/product/213/9452#zone)ID。该参数也可以通过调用  [DescribeZones](/document/api/213/9455) 的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例所属项目ID。该参数可以通过调用 [DescribeProject](/document/api/378/4400) 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例所属的专用宿主机ID列表。如果您有购买专用宿主机并且指定了该参数，则您购买的实例就会随机的部署在这些专用宿主机上。
	HostIds []*string `json:"HostIds,omitempty" name:"HostIds" list`
}

type RedirectInfo struct {

	// 标准输出重定向路径
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" name:"StdoutRedirectPath"`

	// 标准错误重定向路径
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" name:"StderrRedirectPath"`

	// 标准输出重定向文件名，支持三个占位符${BATCH_JOB_ID}、${BATCH_TASK_NAME}、${BATCH_TASK_INSTANCE_INDEX}
	StdoutRedirectFileName *string `json:"StdoutRedirectFileName,omitempty" name:"StdoutRedirectFileName"`

	// 标准错误重定向文件名，支持三个占位符${BATCH_JOB_ID}、${BATCH_TASK_NAME}、${BATCH_TASK_INSTANCE_INDEX}
	StderrRedirectFileName *string `json:"StderrRedirectFileName,omitempty" name:"StderrRedirectFileName"`
}

type RedirectLocalInfo struct {

	// 标准输出重定向本地路径
	StdoutLocalPath *string `json:"StdoutLocalPath,omitempty" name:"StdoutLocalPath"`

	// 标准错误重定向本地路径
	StderrLocalPath *string `json:"StderrLocalPath,omitempty" name:"StderrLocalPath"`

	// 标准输出重定向本地文件名，支持三个占位符${BATCH_JOB_ID}、${BATCH_TASK_NAME}、${BATCH_TASK_INSTANCE_INDEX}
	StdoutLocalFileName *string `json:"StdoutLocalFileName,omitempty" name:"StdoutLocalFileName"`

	// 标准错误重定向本地文件名，支持三个占位符${BATCH_JOB_ID}、${BATCH_TASK_NAME}、${BATCH_TASK_INSTANCE_INDEX}
	StderrLocalFileName *string `json:"StderrLocalFileName,omitempty" name:"StderrLocalFileName"`
}

type RetryJobsRequest struct {
	*tchttp.BaseRequest

	// 作业ID列表。
	JobIds []*string `json:"JobIds,omitempty" name:"JobIds" list`
}

func (r *RetryJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RetryJobsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RetryJobsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RetryJobsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunMonitorServiceEnabled struct {

	// 是否开启[云监控](/document/product/248)服务。取值范围：<br><li>TRUE：表示开启云监控服务<br><li>FALSE：表示不开启云监控服务<br><br>默认取值：TRUE。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {

	// 是否开启[云安全](/document/product/296)服务。取值范围：<br><li>TRUE：表示开启云安全服务<br><li>FALSE：表示不开启云安全服务<br><br>默认取值：TRUE。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type SpotMarketOptions struct {

	// 竞价出价
	MaxPrice *string `json:"MaxPrice,omitempty" name:"MaxPrice"`

	// 竞价请求类型，当前仅支持类型：one-time
	SpotInstanceType *string `json:"SpotInstanceType,omitempty" name:"SpotInstanceType"`
}

type StorageBlock struct {

	// HDD本地存储类型，值为：LOCAL_PRO.
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// HDD本地存储的最小容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`

	// HDD本地存储的最大容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

type SubmitJobRequest struct {
	*tchttp.BaseRequest

	// 作业所提交的位置信息。通过该参数可以指定作业关联CVM所属可用区等信息。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 作业信息
	Job *Job `json:"Job,omitempty" name:"Job"`

	// 用于保证请求幂等性的字符串。该字符串由用户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
}

func (r *SubmitJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 当通过本接口来提交作业时会返回该参数，表示一个作业ID。返回作业ID列表并不代表作业解析/运行成功，可根据 DescribeJob 接口查询其状态。
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SystemDisk struct {

	// 系统盘类型。系统盘类型限制详见[CVM实例配置](/document/product/213/2177)。取值范围：<br><li>LOCAL_BASIC：本地硬盘<br><li>LOCAL_SSD：本地SSD硬盘<br><li>CLOUD_BASIC：普通云硬盘<br><li>CLOUD_SSD：SSD云硬盘<br><li>CLOUD_PREMIUM：高性能云硬盘<br><br>默认取值：CLOUD_BASIC。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 系统盘ID。LOCAL_BASIC 和 LOCAL_SSD 类型没有ID。暂时不支持该参数。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 系统盘大小，单位：GB。默认值为 50
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type Task struct {

	// 应用程序信息
	Application *Application `json:"Application,omitempty" name:"Application"`

	// 任务名称，在一个作业内部唯一
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务实例运行个数
	TaskInstanceNum *uint64 `json:"TaskInstanceNum,omitempty" name:"TaskInstanceNum"`

	// 运行环境信息，ComputeEnv 和 EnvId 必须指定一个（且只有一个）参数。
	ComputeEnv *AnonymousComputeEnv `json:"ComputeEnv,omitempty" name:"ComputeEnv"`

	// 计算环境ID，ComputeEnv 和 EnvId 必须指定一个（且只有一个）参数。
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 重定向信息
	RedirectInfo *RedirectInfo `json:"RedirectInfo,omitempty" name:"RedirectInfo"`

	// 重定向本地信息
	RedirectLocalInfo *RedirectLocalInfo `json:"RedirectLocalInfo,omitempty" name:"RedirectLocalInfo"`

	// 输入映射
	InputMappings []*InputMapping `json:"InputMappings,omitempty" name:"InputMappings" list`

	// 输出映射
	OutputMappings []*OutputMapping `json:"OutputMappings,omitempty" name:"OutputMappings" list`

	// 输出映射配置
	OutputMappingConfigs []*OutputMappingConfig `json:"OutputMappingConfigs,omitempty" name:"OutputMappingConfigs" list`

	// 自定义环境变量
	EnvVars []*EnvVar `json:"EnvVars,omitempty" name:"EnvVars" list`

	// 授权信息
	Authentications []*Authentication `json:"Authentications,omitempty" name:"Authentications" list`

	// TaskInstance失败后处理方式，取值包括TERMINATE（默认）、INTERRUPT、FAST_INTERRUPT。
	FailedAction *string `json:"FailedAction,omitempty" name:"FailedAction"`

	// 任务失败后的最大重试次数，默认为0
	MaxRetryCount *uint64 `json:"MaxRetryCount,omitempty" name:"MaxRetryCount"`

	// 任务启动后的超时时间，单位秒，默认为86400秒
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`

	// 任务最大并发数限制，默认没有限制。
	MaxConcurrentNum *uint64 `json:"MaxConcurrentNum,omitempty" name:"MaxConcurrentNum"`
}

type TaskInstanceLog struct {

	// 任务实例
	TaskInstanceIndex *uint64 `json:"TaskInstanceIndex,omitempty" name:"TaskInstanceIndex"`

	// 标准输出日志（Base64编码）
	// 注意：此字段可能返回 null，表示取不到有效值。
	StdoutLog *string `json:"StdoutLog,omitempty" name:"StdoutLog"`

	// 标准错误日志（Base64编码）
	// 注意：此字段可能返回 null，表示取不到有效值。
	StderrLog *string `json:"StderrLog,omitempty" name:"StderrLog"`

	// 标准输出重定向路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" name:"StdoutRedirectPath"`

	// 标准错误重定向路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" name:"StderrRedirectPath"`

	// 标准输出重定向文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	StdoutRedirectFileName *string `json:"StdoutRedirectFileName,omitempty" name:"StdoutRedirectFileName"`

	// 标准错误重定向文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	StderrRedirectFileName *string `json:"StderrRedirectFileName,omitempty" name:"StderrRedirectFileName"`
}

type TaskInstanceMetrics struct {

	// Submitted个数
	SubmittedCount *int64 `json:"SubmittedCount,omitempty" name:"SubmittedCount"`

	// Pending个数
	PendingCount *int64 `json:"PendingCount,omitempty" name:"PendingCount"`

	// Runnable个数
	RunnableCount *int64 `json:"RunnableCount,omitempty" name:"RunnableCount"`

	// Starting个数
	StartingCount *int64 `json:"StartingCount,omitempty" name:"StartingCount"`

	// Running个数
	RunningCount *int64 `json:"RunningCount,omitempty" name:"RunningCount"`

	// Succeed个数
	SucceedCount *int64 `json:"SucceedCount,omitempty" name:"SucceedCount"`

	// FailedInterrupted个数
	FailedInterruptedCount *int64 `json:"FailedInterruptedCount,omitempty" name:"FailedInterruptedCount"`

	// Failed个数
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`
}

type TaskInstanceView struct {

	// 任务实例索引
	TaskInstanceIndex *int64 `json:"TaskInstanceIndex,omitempty" name:"TaskInstanceIndex"`

	// 任务实例状态
	TaskInstanceState *string `json:"TaskInstanceState,omitempty" name:"TaskInstanceState"`

	// 应用程序执行结束的exit code
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExitCode *int64 `json:"ExitCode,omitempty" name:"ExitCode"`

	// 任务实例状态原因，任务实例失败时，会记录失败原因
	StateReason *string `json:"StateReason,omitempty" name:"StateReason"`

	// 任务实例运行时所在计算节点（例如CVM）的InstanceId。任务实例未运行或者完结时，本字段为空。任务实例重试时，本字段会随之变化
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComputeNodeInstanceId *string `json:"ComputeNodeInstanceId,omitempty" name:"ComputeNodeInstanceId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LaunchTime *string `json:"LaunchTime,omitempty" name:"LaunchTime"`

	// 开始运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningTime *string `json:"RunningTime,omitempty" name:"RunningTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 重定向信息
	RedirectInfo *RedirectInfo `json:"RedirectInfo,omitempty" name:"RedirectInfo"`

	// 任务实例状态原因详情，任务实例失败时，会记录失败原因
	StateDetailedReason *string `json:"StateDetailedReason,omitempty" name:"StateDetailedReason"`
}

type TaskMetrics struct {

	// Submitted个数
	SubmittedCount *int64 `json:"SubmittedCount,omitempty" name:"SubmittedCount"`

	// Pending个数
	PendingCount *int64 `json:"PendingCount,omitempty" name:"PendingCount"`

	// Runnable个数
	RunnableCount *int64 `json:"RunnableCount,omitempty" name:"RunnableCount"`

	// Starting个数
	StartingCount *int64 `json:"StartingCount,omitempty" name:"StartingCount"`

	// Running个数
	RunningCount *int64 `json:"RunningCount,omitempty" name:"RunningCount"`

	// Succeed个数
	SucceedCount *int64 `json:"SucceedCount,omitempty" name:"SucceedCount"`

	// FailedInterrupted个数
	FailedInterruptedCount *int64 `json:"FailedInterruptedCount,omitempty" name:"FailedInterruptedCount"`

	// Failed个数
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`
}

type TaskTemplateView struct {

	// 任务模板ID
	TaskTemplateId *string `json:"TaskTemplateId,omitempty" name:"TaskTemplateId"`

	// 任务模板名称
	TaskTemplateName *string `json:"TaskTemplateName,omitempty" name:"TaskTemplateName"`

	// 任务模板描述
	TaskTemplateDescription *string `json:"TaskTemplateDescription,omitempty" name:"TaskTemplateDescription"`

	// 任务模板信息
	TaskTemplateInfo *Task `json:"TaskTemplateInfo,omitempty" name:"TaskTemplateInfo"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type TaskView struct {

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务状态
	TaskState *string `json:"TaskState,omitempty" name:"TaskState"`

	// 开始时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type TerminateComputeNodeRequest struct {
	*tchttp.BaseRequest

	// 计算环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 计算节点ID
	ComputeNodeId *string `json:"ComputeNodeId,omitempty" name:"ComputeNodeId"`
}

func (r *TerminateComputeNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateComputeNodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateComputeNodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateComputeNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateComputeNodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateComputeNodesRequest struct {
	*tchttp.BaseRequest

	// 计算环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 计算节点ID列表
	ComputeNodeIds []*string `json:"ComputeNodeIds,omitempty" name:"ComputeNodeIds" list`
}

func (r *TerminateComputeNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateComputeNodesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateComputeNodesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateComputeNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateComputeNodesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateJobRequest struct {
	*tchttp.BaseRequest

	// 作业ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *TerminateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateTaskInstanceRequest struct {
	*tchttp.BaseRequest

	// 作业ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务实例索引
	TaskInstanceIndex *int64 `json:"TaskInstanceIndex,omitempty" name:"TaskInstanceIndex"`
}

func (r *TerminateTaskInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateTaskInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateTaskInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateTaskInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateTaskInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VirtualPrivateCloud struct {

	// 私有网络ID，形如`vpc-xxx`。有效的VpcId可通过登录[控制台](https://console.cloud.tencent.com/vpc/vpc?rid=1)查询；也可以调用接口 [DescribeVpcEx](/document/api/215/1372) ，从接口返回中的`unVpcId`字段获取。若在创建子机时VpcId与SubnetId同时传入`DEFAULT`，则强制使用默认vpc网络。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络子网ID，形如`subnet-xxx`。有效的私有网络子网ID可通过登录[控制台](https://console.cloud.tencent.com/vpc/subnet?rid=1)查询；也可以调用接口  [DescribeSubnets](/document/api/215/15784) ，从接口返回中的`unSubnetId`字段获取。若在创建子机时SubnetId与VpcId同时传入`DEFAULT`，则强制使用默认vpc网络。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 是否用作公网网关。公网网关只有在实例拥有公网IP以及处于私有网络下时才能正常使用。取值范围：<br><li>TRUE：表示用作公网网关<br><li>FALSE：表示不用作公网网关<br><br>默认取值：FALSE。
	AsVpcGateway *bool `json:"AsVpcGateway,omitempty" name:"AsVpcGateway"`

	// 私有网络子网 IP 数组，在创建实例、修改实例vpc属性操作中可使用此参数。当前仅批量创建多台实例时支持传入相同子网的多个 IP。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`
}
