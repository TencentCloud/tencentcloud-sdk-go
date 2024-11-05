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

package v20190422

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type CCN struct {
	// 私有网络 ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网 ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 云联网 ID，如 ccn-rahigzjd
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`
}

// Predefined struct for user
type CheckSavepointRequestParams struct {
	// 作业 id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 快照资源 id
	SerialId *string `json:"SerialId,omitnil,omitempty" name:"SerialId"`

	// 快照类型 1: savepoint；2: checkpoint；3: cancelWithSavepoint
	RecordType *int64 `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 快照路径，目前只支持 cos 路径
	SavepointPath *string `json:"SavepointPath,omitnil,omitempty" name:"SavepointPath"`

	// 工作空间 id
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type CheckSavepointRequest struct {
	*tchttp.BaseRequest
	
	// 作业 id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 快照资源 id
	SerialId *string `json:"SerialId,omitnil,omitempty" name:"SerialId"`

	// 快照类型 1: savepoint；2: checkpoint；3: cancelWithSavepoint
	RecordType *int64 `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 快照路径，目前只支持 cos 路径
	SavepointPath *string `json:"SavepointPath,omitnil,omitempty" name:"SavepointPath"`

	// 工作空间 id
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *CheckSavepointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckSavepointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "SerialId")
	delete(f, "RecordType")
	delete(f, "SavepointPath")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckSavepointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckSavepointResponseParams struct {
	// 资源 id
	SerialId *string `json:"SerialId,omitnil,omitempty" name:"SerialId"`

	// 1=可用，2=不可用
	SavepointStatus *int64 `json:"SavepointStatus,omitnil,omitempty" name:"SavepointStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckSavepointResponse struct {
	*tchttp.BaseResponse
	Response *CheckSavepointResponseParams `json:"Response"`
}

func (r *CheckSavepointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckSavepointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClazzLevel struct {
	// java类全路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Clazz *string `json:"Clazz,omitnil,omitempty" name:"Clazz"`

	// 日志级别  TRACE，DEBUG、INFO、WARN、ERROR
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`
}

type Cluster struct {
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 用户 AppID
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 主账号 UIN
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 创建者 UIN
	CreatorUin *string `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 集群状态, 1 未初始化,，3 初始化中，2 运行中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 描述
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 集群创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后一次操作集群的时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// CU 数量
	CuNum *uint64 `json:"CuNum,omitnil,omitempty" name:"CuNum"`

	// CU 内存规格
	CuMem *uint64 `json:"CuMem,omitnil,omitempty" name:"CuMem"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 网络
	CCNs []*CCN `json:"CCNs,omitnil,omitempty" name:"CCNs"`

	// 网络
	NetEnvironmentType *uint64 `json:"NetEnvironmentType,omitnil,omitempty" name:"NetEnvironmentType"`

	// 空闲 CU
	FreeCuNum *int64 `json:"FreeCuNum,omitnil,omitempty" name:"FreeCuNum"`

	// 集群绑定的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 集群隔离时间; 没隔离时间，则为 -
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolatedTime *string `json:"IsolatedTime,omitnil,omitempty" name:"IsolatedTime"`

	// 集群过期时间; 没过期概念，则为 -
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 距离过期还有多少秒; 没过期概念，则为 -
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecondsUntilExpiry *string `json:"SecondsUntilExpiry,omitnil,omitempty" name:"SecondsUntilExpiry"`

	// 自动续费标记，0 表示默认状态 (用户未设置，即初始状态，用户开通了预付费不停服特权会进行自动续费)， 1 表示自动续费，2表示明确不自动续费(用户设置)
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 集群的默认 COS 存储桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultCOSBucket *string `json:"DefaultCOSBucket,omitnil,omitempty" name:"DefaultCOSBucket"`

	// 集群的CLS 日志集 LogSet
	// 注意：此字段可能返回 null，表示取不到有效值。
	CLSLogSet *string `json:"CLSLogSet,omitnil,omitempty" name:"CLSLogSet"`

	// 集群的CLS 日志主题 TopicId
	// 注意：此字段可能返回 null，表示取不到有效值。
	CLSTopicId *string `json:"CLSTopicId,omitnil,omitempty" name:"CLSTopicId"`

	// 集群的CLS 日志集  名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	CLSLogName *string `json:"CLSLogName,omitnil,omitempty" name:"CLSLogName"`

	// 集群的CLS 日志主题  名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	CLSTopicName *string `json:"CLSTopicName,omitnil,omitempty" name:"CLSTopicName"`

	// 集群的版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *ClusterVersion `json:"Version,omitnil,omitempty" name:"Version"`

	// 细粒度资源下的空闲CU
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreeCu *float64 `json:"FreeCu,omitnil,omitempty" name:"FreeCu"`

	// 集群的默认日志采集配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultLogCollectConf *string `json:"DefaultLogCollectConf,omitnil,omitempty" name:"DefaultLogCollectConf"`

	// 取值：0-没有设置，1-已设置，2-不允许设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomizedDNSEnabled *int64 `json:"CustomizedDNSEnabled,omitnil,omitempty" name:"CustomizedDNSEnabled"`

	// 空间信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Correlations []*WorkSpaceClusterItem `json:"Correlations,omitnil,omitempty" name:"Correlations"`

	// 运行CU
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningCu *float64 `json:"RunningCu,omitnil,omitempty" name:"RunningCu"`

	// 0 后付费,1 预付费
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 前端区分 集群是否需要2CU逻辑 因为历史集群 变配不需要, default 1  新集群都需要
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNeedManageNode *int64 `json:"IsNeedManageNode,omitnil,omitempty" name:"IsNeedManageNode"`

	// session集群信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterSessions []*ClusterSession `json:"ClusterSessions,omitnil,omitempty" name:"ClusterSessions"`

	// V3版本 = 2
	// 注意：此字段可能返回 null，表示取不到有效值。
	ArchGeneration *uint64 `json:"ArchGeneration,omitnil,omitempty" name:"ArchGeneration"`

	// 0:TKE, 1:EKS
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *uint64 `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 订单信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Orders []*Order `json:"Orders,omitnil,omitempty" name:"Orders"`

	// Gateway信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SqlGateways []*SqlGatewayItem `json:"SqlGateways,omitnil,omitempty" name:"SqlGateways"`

	// 0 公网访问 // 1 内网访问	
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebUIType *int64 `json:"WebUIType,omitnil,omitempty" name:"WebUIType"`

	// 2 独享集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 子eks集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubEks *SubEks `json:"SubEks,omitnil,omitempty" name:"SubEks"`

	// 上级集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentSerialId *string `json:"AgentSerialId,omitnil,omitempty" name:"AgentSerialId"`

	// 资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingResourceMode *string `json:"BillingResourceMode,omitnil,omitempty" name:"BillingResourceMode"`

	// Cu比例
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemRatio *int64 `json:"MemRatio,omitnil,omitempty" name:"MemRatio"`
}

type ClusterGroupSetItem struct {
	// clusterGroup 的 SerialId
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 账号 APPID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 主账号 UIN
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 创建账号 UIN
	CreatorUin *string `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// CU 数量
	CuNum *int64 `json:"CuNum,omitnil,omitempty" name:"CuNum"`

	// CU 内存规格
	CuMem *int64 `json:"CuMem,omitnil,omitempty" name:"CuMem"`

	// 集群状态, 1 未初始化,，3 初始化中，2 运行中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 集群创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后一次操作集群的时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 描述
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 网络
	NetEnvironmentType *int64 `json:"NetEnvironmentType,omitnil,omitempty" name:"NetEnvironmentType"`

	// 空闲 CU
	FreeCuNum *int64 `json:"FreeCuNum,omitnil,omitempty" name:"FreeCuNum"`

	// 细粒度资源下的空闲CU
	FreeCu *float64 `json:"FreeCu,omitnil,omitempty" name:"FreeCu"`

	// 运行中CU
	RunningCu *float64 `json:"RunningCu,omitnil,omitempty" name:"RunningCu"`

	// 付费模式
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 弹性
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubEks *SubEks `json:"SubEks,omitnil,omitempty" name:"SubEks"`

	// 默认 "" 包销模式 "exclusiveSale"
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingResourceMode *string `json:"BillingResourceMode,omitnil,omitempty" name:"BillingResourceMode"`
}

type ClusterSession struct {
	// 集群SerialId
	ClusterGroupSerialId *string `json:"ClusterGroupSerialId,omitnil,omitempty" name:"ClusterGroupSerialId"`

	// 创建者appId
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 创建者主账号
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 创建者账号
	CreatorUin *string `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 区域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Session集群状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Session集群消耗的cu数量
	CuNum *float64 `json:"CuNum,omitnil,omitempty" name:"CuNum"`

	// Session集群的Flink版本
	FlinkVersion *string `json:"FlinkVersion,omitnil,omitempty" name:"FlinkVersion"`

	// session集群FlinkUi地址
	WebUIUrl *string `json:"WebUIUrl,omitnil,omitempty" name:"WebUIUrl"`

	// session集群高级参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 引用资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceRefs []*SessionClusterRefItem `json:"ResourceRefs,omitnil,omitempty" name:"ResourceRefs"`

	// JobManager的规格
	JobManagerCuSpec *float64 `json:"JobManagerCuSpec,omitnil,omitempty" name:"JobManagerCuSpec"`

	// TaskManager的规格
	TaskManagerCuSpec *float64 `json:"TaskManagerCuSpec,omitnil,omitempty" name:"TaskManagerCuSpec"`

	// TaskManager启动的数量
	TaskManagerNum *int64 `json:"TaskManagerNum,omitnil,omitempty" name:"TaskManagerNum"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// JobManagerCpu
	JobManagerCpu *float64 `json:"JobManagerCpu,omitnil,omitempty" name:"JobManagerCpu"`

	// JobManagerMem
	JobManagerMem *float64 `json:"JobManagerMem,omitnil,omitempty" name:"JobManagerMem"`

	// TaskManagerCpu
	TaskManagerCpu *float64 `json:"TaskManagerCpu,omitnil,omitempty" name:"TaskManagerCpu"`

	// TaskManagerMem
	TaskManagerMem *float64 `json:"TaskManagerMem,omitnil,omitempty" name:"TaskManagerMem"`
}

type ClusterVersion struct {
	// 集群的Flink版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Flink *string `json:"Flink,omitnil,omitempty" name:"Flink"`

	// 集群支持的Flink版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportedFlink []*string `json:"SupportedFlink,omitnil,omitempty" name:"SupportedFlink"`
}

type CopyJobItem struct {
	// 需要复制的作业serial id
	SourceId *string `json:"SourceId,omitnil,omitempty" name:"SourceId"`

	// 目标集群的cluster serial id
	TargetClusterId *string `json:"TargetClusterId,omitnil,omitempty" name:"TargetClusterId"`

	// 需要复制的作业名称
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 新作业的名称
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// 新作业的目录id
	TargetFolderId *string `json:"TargetFolderId,omitnil,omitempty" name:"TargetFolderId"`

	// 源作业类型
	JobType *int64 `json:"JobType,omitnil,omitempty" name:"JobType"`
}

type CopyJobResult struct {
	// 原作业id
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 原作业名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// 新作业名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetJobName *string `json:"TargetJobName,omitnil,omitempty" name:"TargetJobName"`

	// 新作业id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetJobId *string `json:"TargetJobId,omitnil,omitempty" name:"TargetJobId"`

	// 失败时候的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 0 成功  -1 失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// 目标集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 目标集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 作业类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobType *int64 `json:"JobType,omitnil,omitempty" name:"JobType"`
}

// Predefined struct for user
type CopyJobsRequestParams struct {
	// 复制明细列表
	JobItems []*CopyJobItem `json:"JobItems,omitnil,omitempty" name:"JobItems"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 目标工作空间 SerialId
	TargetWorkspaceId *string `json:"TargetWorkspaceId,omitnil,omitempty" name:"TargetWorkspaceId"`
}

type CopyJobsRequest struct {
	*tchttp.BaseRequest
	
	// 复制明细列表
	JobItems []*CopyJobItem `json:"JobItems,omitnil,omitempty" name:"JobItems"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 目标工作空间 SerialId
	TargetWorkspaceId *string `json:"TargetWorkspaceId,omitnil,omitempty" name:"TargetWorkspaceId"`
}

func (r *CopyJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobItems")
	delete(f, "WorkSpaceId")
	delete(f, "TargetWorkspaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyJobsResponseParams struct {
	// 成功条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 失败条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailCount *int64 `json:"FailCount,omitnil,omitempty" name:"FailCount"`

	// 结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CopyJobsResults []*CopyJobResult `json:"CopyJobsResults,omitnil,omitempty" name:"CopyJobsResults"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CopyJobsResponse struct {
	*tchttp.BaseResponse
	Response *CopyJobsResponseParams `json:"Response"`
}

func (r *CopyJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFolderRequestParams struct {
	// 新建文件夹名
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 新建文件夹的父目录ID（根目录为"root"）
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 文件夹类型，0是任务文件夹，1是依赖文件夹
	FolderType *int64 `json:"FolderType,omitnil,omitempty" name:"FolderType"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type CreateFolderRequest struct {
	*tchttp.BaseRequest
	
	// 新建文件夹名
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 新建文件夹的父目录ID（根目录为"root"）
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 文件夹类型，0是任务文件夹，1是依赖文件夹
	FolderType *int64 `json:"FolderType,omitnil,omitempty" name:"FolderType"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *CreateFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FolderName")
	delete(f, "ParentId")
	delete(f, "FolderType")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFolderResponseParams struct {
	// 新建文件夹的唯一ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFolderResponse struct {
	*tchttp.BaseResponse
	Response *CreateFolderResponseParams `json:"Response"`
}

func (r *CreateFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateJobConfigRequestParams struct {
	// 作业Id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 主类
	EntrypointClass *string `json:"EntrypointClass,omitnil,omitempty" name:"EntrypointClass"`

	// 主类入参
	ProgramArgs *string `json:"ProgramArgs,omitnil,omitempty" name:"ProgramArgs"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 资源引用数组
	ResourceRefs []*ResourceRef `json:"ResourceRefs,omitnil,omitempty" name:"ResourceRefs"`

	// 作业默认并行度
	DefaultParallelism *uint64 `json:"DefaultParallelism,omitnil,omitempty" name:"DefaultParallelism"`

	// 系统参数
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 1: 作业配置达到上限之后，自动删除可删除的最早版本
	AutoDelete *int64 `json:"AutoDelete,omitnil,omitempty" name:"AutoDelete"`

	// 作业使用的 COS 存储桶名
	COSBucket *string `json:"COSBucket,omitnil,omitempty" name:"COSBucket"`

	// 是否采集作业日志
	LogCollect *bool `json:"LogCollect,omitnil,omitempty" name:"LogCollect"`

	// JobManager规格
	JobManagerSpec *float64 `json:"JobManagerSpec,omitnil,omitempty" name:"JobManagerSpec"`

	// TaskManager规格
	TaskManagerSpec *float64 `json:"TaskManagerSpec,omitnil,omitempty" name:"TaskManagerSpec"`

	// CLS日志集ID
	ClsLogsetId *string `json:"ClsLogsetId,omitnil,omitempty" name:"ClsLogsetId"`

	// CLS日志主题ID
	ClsTopicId *string `json:"ClsTopicId,omitnil,omitempty" name:"ClsTopicId"`

	// 日志采集类型 2：CLS；3：COS
	LogCollectType *int64 `json:"LogCollectType,omitnil,omitempty" name:"LogCollectType"`

	// pyflink作业运行时使用的python版本
	PythonVersion *string `json:"PythonVersion,omitnil,omitempty" name:"PythonVersion"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 日志级别
	LogLevel *string `json:"LogLevel,omitnil,omitempty" name:"LogLevel"`

	// Oceanus 平台恢复作业开关 1:开启 -1: 关闭
	AutoRecover *int64 `json:"AutoRecover,omitnil,omitempty" name:"AutoRecover"`

	// 类日志级别
	ClazzLevels []*ClazzLevel `json:"ClazzLevels,omitnil,omitempty" name:"ClazzLevels"`

	// 是否打开专家模式
	ExpertModeOn *bool `json:"ExpertModeOn,omitnil,omitempty" name:"ExpertModeOn"`

	// 专家模式的配置
	ExpertModeConfiguration *ExpertModeConfiguration `json:"ExpertModeConfiguration,omitnil,omitempty" name:"ExpertModeConfiguration"`

	// trace链路
	TraceModeOn *bool `json:"TraceModeOn,omitnil,omitempty" name:"TraceModeOn"`

	// trace链路配置
	TraceModeConfiguration *TraceModeConfiguration `json:"TraceModeConfiguration,omitnil,omitempty" name:"TraceModeConfiguration"`

	// checkpoint保留个数
	CheckpointRetainedNum *int64 `json:"CheckpointRetainedNum,omitnil,omitempty" name:"CheckpointRetainedNum"`

	// 算子拓扑图
	JobGraph *JobGraph `json:"JobGraph,omitnil,omitempty" name:"JobGraph"`

	// es索引名称
	EsServerlessIndex *string `json:"EsServerlessIndex,omitnil,omitempty" name:"EsServerlessIndex"`

	// es索引空间
	EsServerlessSpace *string `json:"EsServerlessSpace,omitnil,omitempty" name:"EsServerlessSpace"`

	// flink版本
	FlinkVersion *string `json:"FlinkVersion,omitnil,omitempty" name:"FlinkVersion"`
}

type CreateJobConfigRequest struct {
	*tchttp.BaseRequest
	
	// 作业Id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 主类
	EntrypointClass *string `json:"EntrypointClass,omitnil,omitempty" name:"EntrypointClass"`

	// 主类入参
	ProgramArgs *string `json:"ProgramArgs,omitnil,omitempty" name:"ProgramArgs"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 资源引用数组
	ResourceRefs []*ResourceRef `json:"ResourceRefs,omitnil,omitempty" name:"ResourceRefs"`

	// 作业默认并行度
	DefaultParallelism *uint64 `json:"DefaultParallelism,omitnil,omitempty" name:"DefaultParallelism"`

	// 系统参数
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 1: 作业配置达到上限之后，自动删除可删除的最早版本
	AutoDelete *int64 `json:"AutoDelete,omitnil,omitempty" name:"AutoDelete"`

	// 作业使用的 COS 存储桶名
	COSBucket *string `json:"COSBucket,omitnil,omitempty" name:"COSBucket"`

	// 是否采集作业日志
	LogCollect *bool `json:"LogCollect,omitnil,omitempty" name:"LogCollect"`

	// JobManager规格
	JobManagerSpec *float64 `json:"JobManagerSpec,omitnil,omitempty" name:"JobManagerSpec"`

	// TaskManager规格
	TaskManagerSpec *float64 `json:"TaskManagerSpec,omitnil,omitempty" name:"TaskManagerSpec"`

	// CLS日志集ID
	ClsLogsetId *string `json:"ClsLogsetId,omitnil,omitempty" name:"ClsLogsetId"`

	// CLS日志主题ID
	ClsTopicId *string `json:"ClsTopicId,omitnil,omitempty" name:"ClsTopicId"`

	// 日志采集类型 2：CLS；3：COS
	LogCollectType *int64 `json:"LogCollectType,omitnil,omitempty" name:"LogCollectType"`

	// pyflink作业运行时使用的python版本
	PythonVersion *string `json:"PythonVersion,omitnil,omitempty" name:"PythonVersion"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 日志级别
	LogLevel *string `json:"LogLevel,omitnil,omitempty" name:"LogLevel"`

	// Oceanus 平台恢复作业开关 1:开启 -1: 关闭
	AutoRecover *int64 `json:"AutoRecover,omitnil,omitempty" name:"AutoRecover"`

	// 类日志级别
	ClazzLevels []*ClazzLevel `json:"ClazzLevels,omitnil,omitempty" name:"ClazzLevels"`

	// 是否打开专家模式
	ExpertModeOn *bool `json:"ExpertModeOn,omitnil,omitempty" name:"ExpertModeOn"`

	// 专家模式的配置
	ExpertModeConfiguration *ExpertModeConfiguration `json:"ExpertModeConfiguration,omitnil,omitempty" name:"ExpertModeConfiguration"`

	// trace链路
	TraceModeOn *bool `json:"TraceModeOn,omitnil,omitempty" name:"TraceModeOn"`

	// trace链路配置
	TraceModeConfiguration *TraceModeConfiguration `json:"TraceModeConfiguration,omitnil,omitempty" name:"TraceModeConfiguration"`

	// checkpoint保留个数
	CheckpointRetainedNum *int64 `json:"CheckpointRetainedNum,omitnil,omitempty" name:"CheckpointRetainedNum"`

	// 算子拓扑图
	JobGraph *JobGraph `json:"JobGraph,omitnil,omitempty" name:"JobGraph"`

	// es索引名称
	EsServerlessIndex *string `json:"EsServerlessIndex,omitnil,omitempty" name:"EsServerlessIndex"`

	// es索引空间
	EsServerlessSpace *string `json:"EsServerlessSpace,omitnil,omitempty" name:"EsServerlessSpace"`

	// flink版本
	FlinkVersion *string `json:"FlinkVersion,omitnil,omitempty" name:"FlinkVersion"`
}

func (r *CreateJobConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateJobConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "EntrypointClass")
	delete(f, "ProgramArgs")
	delete(f, "Remark")
	delete(f, "ResourceRefs")
	delete(f, "DefaultParallelism")
	delete(f, "Properties")
	delete(f, "AutoDelete")
	delete(f, "COSBucket")
	delete(f, "LogCollect")
	delete(f, "JobManagerSpec")
	delete(f, "TaskManagerSpec")
	delete(f, "ClsLogsetId")
	delete(f, "ClsTopicId")
	delete(f, "LogCollectType")
	delete(f, "PythonVersion")
	delete(f, "WorkSpaceId")
	delete(f, "LogLevel")
	delete(f, "AutoRecover")
	delete(f, "ClazzLevels")
	delete(f, "ExpertModeOn")
	delete(f, "ExpertModeConfiguration")
	delete(f, "TraceModeOn")
	delete(f, "TraceModeConfiguration")
	delete(f, "CheckpointRetainedNum")
	delete(f, "JobGraph")
	delete(f, "EsServerlessIndex")
	delete(f, "EsServerlessSpace")
	delete(f, "FlinkVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateJobConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateJobConfigResponseParams struct {
	// 作业配置版本号
	Version *uint64 `json:"Version,omitnil,omitempty" name:"Version"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateJobConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateJobConfigResponseParams `json:"Response"`
}

func (r *CreateJobConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateJobConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateJobRequestParams struct {
	// 作业名称，允许输入长度小于50个字符的中文、英文、数字、-（横线）、_（下划线）、.（点），且符号必须半角字符。注意作业名不能和现有作业同名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 作业的类型，1 表示 SQL 作业，2 表示 JAR 作业
	JobType *int64 `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 集群的类型，1 表示共享集群，2 表示独享集群
	ClusterType *int64 `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 当 ClusterType=2 时，必选，用来指定该作业提交的独享集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 设置每 CU 的内存规格，单位为 GB，支持 2、4、8、16（需申请开通白名单后使用）。默认为 4，即 1 CU 对应 4 GB 的运行内存
	CuMem *uint64 `json:"CuMem,omitnil,omitempty" name:"CuMem"`

	// 作业的备注信息，可以随意设置
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 作业名所属文件夹ID，根目录为"root"
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 作业运行的Flink版本
	FlinkVersion *string `json:"FlinkVersion,omitnil,omitempty" name:"FlinkVersion"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 作业标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 作业描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateJobRequest struct {
	*tchttp.BaseRequest
	
	// 作业名称，允许输入长度小于50个字符的中文、英文、数字、-（横线）、_（下划线）、.（点），且符号必须半角字符。注意作业名不能和现有作业同名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 作业的类型，1 表示 SQL 作业，2 表示 JAR 作业
	JobType *int64 `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 集群的类型，1 表示共享集群，2 表示独享集群
	ClusterType *int64 `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 当 ClusterType=2 时，必选，用来指定该作业提交的独享集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 设置每 CU 的内存规格，单位为 GB，支持 2、4、8、16（需申请开通白名单后使用）。默认为 4，即 1 CU 对应 4 GB 的运行内存
	CuMem *uint64 `json:"CuMem,omitnil,omitempty" name:"CuMem"`

	// 作业的备注信息，可以随意设置
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 作业名所属文件夹ID，根目录为"root"
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 作业运行的Flink版本
	FlinkVersion *string `json:"FlinkVersion,omitnil,omitempty" name:"FlinkVersion"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 作业标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 作业描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "JobType")
	delete(f, "ClusterType")
	delete(f, "ClusterId")
	delete(f, "CuMem")
	delete(f, "Remark")
	delete(f, "FolderId")
	delete(f, "FlinkVersion")
	delete(f, "WorkSpaceId")
	delete(f, "Tags")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateJobResponseParams struct {
	// 作业Id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateJobResponseParams `json:"Response"`
}

func (r *CreateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceConfigRequestParams struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 位置信息
	ResourceLoc *ResourceLoc `json:"ResourceLoc,omitnil,omitempty" name:"ResourceLoc"`

	// 资源描述信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 1： 资源版本达到上限，自动删除最早可删除的版本
	AutoDelete *int64 `json:"AutoDelete,omitnil,omitempty" name:"AutoDelete"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type CreateResourceConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 位置信息
	ResourceLoc *ResourceLoc `json:"ResourceLoc,omitnil,omitempty" name:"ResourceLoc"`

	// 资源描述信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 1： 资源版本达到上限，自动删除最早可删除的版本
	AutoDelete *int64 `json:"AutoDelete,omitnil,omitempty" name:"AutoDelete"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *CreateResourceConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "ResourceLoc")
	delete(f, "Remark")
	delete(f, "AutoDelete")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResourceConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceConfigResponseParams struct {
	// 资源版本ID
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateResourceConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateResourceConfigResponseParams `json:"Response"`
}

func (r *CreateResourceConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceRequestParams struct {
	// 资源位置
	ResourceLoc *ResourceLoc `json:"ResourceLoc,omitnil,omitempty" name:"ResourceLoc"`

	// 资源类型。目前只支持 JAR，取值为 1
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源描述
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 资源名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资源版本描述
	ResourceConfigRemark *string `json:"ResourceConfigRemark,omitnil,omitempty" name:"ResourceConfigRemark"`

	// 目录ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type CreateResourceRequest struct {
	*tchttp.BaseRequest
	
	// 资源位置
	ResourceLoc *ResourceLoc `json:"ResourceLoc,omitnil,omitempty" name:"ResourceLoc"`

	// 资源类型。目前只支持 JAR，取值为 1
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源描述
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 资源名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资源版本描述
	ResourceConfigRemark *string `json:"ResourceConfigRemark,omitnil,omitempty" name:"ResourceConfigRemark"`

	// 目录ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *CreateResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceLoc")
	delete(f, "ResourceType")
	delete(f, "Remark")
	delete(f, "Name")
	delete(f, "ResourceConfigRemark")
	delete(f, "FolderId")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceResponseParams struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源版本
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateResourceResponse struct {
	*tchttp.BaseResponse
	Response *CreateResourceResponseParams `json:"Response"`
}

func (r *CreateResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkSpaceRequestParams struct {
	// 工作空间名称
	WorkSpaceName *string `json:"WorkSpaceName,omitnil,omitempty" name:"WorkSpaceName"`

	// 项目空间备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateWorkSpaceRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间名称
	WorkSpaceName *string `json:"WorkSpaceName,omitnil,omitempty" name:"WorkSpaceName"`

	// 项目空间备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateWorkSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkSpaceName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkSpaceResponseParams struct {
	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWorkSpaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkSpaceResponseParams `json:"Response"`
}

func (r *CreateWorkSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFoldersRequestParams struct {
	// 需删除的文件夹唯一ID
	FolderIds []*string `json:"FolderIds,omitnil,omitempty" name:"FolderIds"`

	// 文件夹类型，0是任务文件夹，1是依赖文件夹
	FolderType *int64 `json:"FolderType,omitnil,omitempty" name:"FolderType"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type DeleteFoldersRequest struct {
	*tchttp.BaseRequest
	
	// 需删除的文件夹唯一ID
	FolderIds []*string `json:"FolderIds,omitnil,omitempty" name:"FolderIds"`

	// 文件夹类型，0是任务文件夹，1是依赖文件夹
	FolderType *int64 `json:"FolderType,omitnil,omitempty" name:"FolderType"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *DeleteFoldersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFoldersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FolderIds")
	delete(f, "FolderType")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFoldersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFoldersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFoldersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFoldersResponseParams `json:"Response"`
}

func (r *DeleteFoldersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFoldersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteJobConfigsRequestParams struct {
	// 作业ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 作业配置版本数组
	JobConfigVersions []*int64 `json:"JobConfigVersions,omitnil,omitempty" name:"JobConfigVersions"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type DeleteJobConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 作业ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 作业配置版本数组
	JobConfigVersions []*int64 `json:"JobConfigVersions,omitnil,omitempty" name:"JobConfigVersions"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *DeleteJobConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteJobConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobConfigVersions")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteJobConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteJobConfigsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteJobConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteJobConfigsResponseParams `json:"Response"`
}

func (r *DeleteJobConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteJobConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteJobsRequestParams struct {
	// 作业Id列表
	JobIds []*string `json:"JobIds,omitnil,omitempty" name:"JobIds"`

	// 工作空间Id
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 作业名称列表
	JobNames []*string `json:"JobNames,omitnil,omitempty" name:"JobNames"`
}

type DeleteJobsRequest struct {
	*tchttp.BaseRequest
	
	// 作业Id列表
	JobIds []*string `json:"JobIds,omitnil,omitempty" name:"JobIds"`

	// 工作空间Id
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 作业名称列表
	JobNames []*string `json:"JobNames,omitnil,omitempty" name:"JobNames"`
}

func (r *DeleteJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobIds")
	delete(f, "WorkSpaceId")
	delete(f, "JobNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteJobsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteJobsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteJobsResponseParams `json:"Response"`
}

func (r *DeleteJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceConfigsRequestParams struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源版本数组
	ResourceConfigVersions []*int64 `json:"ResourceConfigVersions,omitnil,omitempty" name:"ResourceConfigVersions"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type DeleteResourceConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源版本数组
	ResourceConfigVersions []*int64 `json:"ResourceConfigVersions,omitnil,omitempty" name:"ResourceConfigVersions"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *DeleteResourceConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "ResourceConfigVersions")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteResourceConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceConfigsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteResourceConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteResourceConfigsResponseParams `json:"Response"`
}

func (r *DeleteResourceConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourcesRequestParams struct {
	// 待删除资源ID列表
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type DeleteResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 待删除资源ID列表
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *DeleteResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceIds")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourcesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteResourcesResponseParams `json:"Response"`
}

func (r *DeleteResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableConfigRequestParams struct {
	// 作业ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 调试作业ID
	DebugId *int64 `json:"DebugId,omitnil,omitempty" name:"DebugId"`

	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type DeleteTableConfigRequest struct {
	*tchttp.BaseRequest
	
	// 作业ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 调试作业ID
	DebugId *int64 `json:"DebugId,omitnil,omitempty" name:"DebugId"`

	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *DeleteTableConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "DebugId")
	delete(f, "TableName")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTableConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTableConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTableConfigResponseParams `json:"Response"`
}

func (r *DeleteTableConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkSpaceRequestParams struct {
	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type DeleteWorkSpaceRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *DeleteWorkSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkSpaceResponseParams struct {
	// 是否删除
	Delete *bool `json:"Delete,omitnil,omitempty" name:"Delete"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWorkSpaceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWorkSpaceResponseParams `json:"Response"`
}

func (r *DeleteWorkSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersRequestParams struct {
	// 按照一个或者多个集群 ID 查询，每次请求的集群上限为 100
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 偏移量，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 请求的集群数量，默认 20，最大值 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 集群信息结果排序规则，1 按时间降序，2 按照时间升序，3  按照状态排序
	OrderType *int64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 过滤规则
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个集群 ID 查询，每次请求的集群上限为 100
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 偏移量，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 请求的集群数量，默认 20，最大值 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 集群信息结果排序规则，1 按时间降序，2 按照时间升序，3  按照状态排序
	OrderType *int64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 过滤规则
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderType")
	delete(f, "Filters")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersResponseParams struct {
	// 集群总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群列表
	ClusterSet []*Cluster `json:"ClusterSet,omitnil,omitempty" name:"ClusterSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClustersResponseParams `json:"Response"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFolderRequestParams struct {
	// folder id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// workspace id
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 1:资源文件夹
	// 其他:作业文件夹
	FolderType *int64 `json:"FolderType,omitnil,omitempty" name:"FolderType"`
}

type DescribeFolderRequest struct {
	*tchttp.BaseRequest
	
	// folder id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// workspace id
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 1:资源文件夹
	// 其他:作业文件夹
	FolderType *int64 `json:"FolderType,omitnil,omitempty" name:"FolderType"`
}

func (r *DescribeFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FolderId")
	delete(f, "WorkSpaceId")
	delete(f, "FolderType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFolderResponseParams struct {
	// folder id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// folder name
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 父文件夹id
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 文件夹类型
	FolderType *int64 `json:"FolderType,omitnil,omitempty" name:"FolderType"`

	// workspace id
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 子文件夹信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubFolderInfo []*SubFolderInfo `json:"SubFolderInfo,omitnil,omitempty" name:"SubFolderInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFolderResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFolderResponseParams `json:"Response"`
}

func (r *DescribeFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobConfigsRequestParams struct {
	// 作业Id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 作业配置版本
	JobConfigVersions []*uint64 `json:"JobConfigVersions,omitnil,omitempty" name:"JobConfigVersions"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认20，最大100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// true 表示只展示草稿
	OnlyDraft *bool `json:"OnlyDraft,omitnil,omitempty" name:"OnlyDraft"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type DescribeJobConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 作业Id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 作业配置版本
	JobConfigVersions []*uint64 `json:"JobConfigVersions,omitnil,omitempty" name:"JobConfigVersions"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认20，最大100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// true 表示只展示草稿
	OnlyDraft *bool `json:"OnlyDraft,omitnil,omitempty" name:"OnlyDraft"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *DescribeJobConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobConfigVersions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "OnlyDraft")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobConfigsResponseParams struct {
	// 总的配置版本数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 作业配置列表
	JobConfigSet []*JobConfig `json:"JobConfigSet,omitnil,omitempty" name:"JobConfigSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeJobConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobConfigsResponseParams `json:"Response"`
}

func (r *DescribeJobConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobEventsRequestParams struct {
	// 作业的 ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 筛选条件：起始 Unix 时间戳（秒）
	StartTimestamp *uint64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// 筛选条件：结束 Unix 时间戳（秒）
	EndTimestamp *uint64 `json:"EndTimestamp,omitnil,omitempty" name:"EndTimestamp"`

	// 事件类型。如果不传则返回所有类型的数据
	Types []*string `json:"Types,omitnil,omitempty" name:"Types"`

	// 运行实例 ID 数组
	RunningOrderIds []*uint64 `json:"RunningOrderIds,omitnil,omitempty" name:"RunningOrderIds"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type DescribeJobEventsRequest struct {
	*tchttp.BaseRequest
	
	// 作业的 ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 筛选条件：起始 Unix 时间戳（秒）
	StartTimestamp *uint64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// 筛选条件：结束 Unix 时间戳（秒）
	EndTimestamp *uint64 `json:"EndTimestamp,omitnil,omitempty" name:"EndTimestamp"`

	// 事件类型。如果不传则返回所有类型的数据
	Types []*string `json:"Types,omitnil,omitempty" name:"Types"`

	// 运行实例 ID 数组
	RunningOrderIds []*uint64 `json:"RunningOrderIds,omitnil,omitempty" name:"RunningOrderIds"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *DescribeJobEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "StartTimestamp")
	delete(f, "EndTimestamp")
	delete(f, "Types")
	delete(f, "RunningOrderIds")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobEventsResponseParams struct {
	// 该作业指定范围内的事件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Events []*JobEvent `json:"Events,omitnil,omitempty" name:"Events"`

	// 该作业指定范围内运行实例 ID 数组，仅当入参没有传入 RunningOrderIds 参数时才会返回。倒序输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningOrderIds []*uint64 `json:"RunningOrderIds,omitnil,omitempty" name:"RunningOrderIds"`

	// 事件的总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeJobEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobEventsResponseParams `json:"Response"`
}

func (r *DescribeJobEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobRuntimeInfoRequestParams struct {
	// 作业ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 工作空间ID
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 作业运行信息 key
	IncludeInfo []*string `json:"IncludeInfo,omitnil,omitempty" name:"IncludeInfo"`
}

type DescribeJobRuntimeInfoRequest struct {
	*tchttp.BaseRequest
	
	// 作业ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 工作空间ID
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 作业运行信息 key
	IncludeInfo []*string `json:"IncludeInfo,omitnil,omitempty" name:"IncludeInfo"`
}

func (r *DescribeJobRuntimeInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobRuntimeInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "WorkSpaceId")
	delete(f, "IncludeInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobRuntimeInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobRuntimeInfoResponseParams struct {
	// 作业运行时信息
	JobRuntimeInfo []*JobRuntimeInfo `json:"JobRuntimeInfo,omitnil,omitempty" name:"JobRuntimeInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeJobRuntimeInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobRuntimeInfoResponseParams `json:"Response"`
}

func (r *DescribeJobRuntimeInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobRuntimeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobSavepointRequestParams struct {
	// 作业 SerialId
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 分页参数，单页总数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数，偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 2 是checkpoint
	// 1 是触发savepoint
	// 3 停止触发的savepoint
	RecordTypes []*int64 `json:"RecordTypes,omitnil,omitempty" name:"RecordTypes"`
}

type DescribeJobSavepointRequest struct {
	*tchttp.BaseRequest
	
	// 作业 SerialId
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 分页参数，单页总数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数，偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 2 是checkpoint
	// 1 是触发savepoint
	// 3 停止触发的savepoint
	RecordTypes []*int64 `json:"RecordTypes,omitnil,omitempty" name:"RecordTypes"`
}

func (r *DescribeJobSavepointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobSavepointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "WorkSpaceId")
	delete(f, "RecordTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobSavepointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobSavepointResponseParams struct {
	// 快照列表总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalNumber *int64 `json:"TotalNumber,omitnil,omitempty" name:"TotalNumber"`

	// 快照列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Savepoint []*Savepoint `json:"Savepoint,omitnil,omitempty" name:"Savepoint"`

	// 进行中的快照列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningSavepoint []*Savepoint `json:"RunningSavepoint,omitnil,omitempty" name:"RunningSavepoint"`

	// 进行中的快照列表总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningTotalNumber *int64 `json:"RunningTotalNumber,omitnil,omitempty" name:"RunningTotalNumber"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeJobSavepointResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobSavepointResponseParams `json:"Response"`
}

func (r *DescribeJobSavepointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobSavepointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobSubmissionLogRequestParams struct {
	// 作业ID，例如：cql-6v1jkxrn
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 起始时间，unix时间戳，毫秒级，例如：1611754219108
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，unix时间戳，毫秒级，例如：1611754219108
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 作业运行的实例ID, 例如：1,2,3。默认为0，表示未选中任何实例，搜索该时间段内最近的一个实例的日志
	RunningOrderId *int64 `json:"RunningOrderId,omitnil,omitempty" name:"RunningOrderId"`

	// 日志搜索的关键词，默认为空
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 日志搜索的游标，可透传上次返回的值，默认为空
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 时间戳排序规则，asc - 升序，desc - 降序。默认为升序
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 搜索的日志条数上限值，最大为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeJobSubmissionLogRequest struct {
	*tchttp.BaseRequest
	
	// 作业ID，例如：cql-6v1jkxrn
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 起始时间，unix时间戳，毫秒级，例如：1611754219108
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，unix时间戳，毫秒级，例如：1611754219108
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 作业运行的实例ID, 例如：1,2,3。默认为0，表示未选中任何实例，搜索该时间段内最近的一个实例的日志
	RunningOrderId *int64 `json:"RunningOrderId,omitnil,omitempty" name:"RunningOrderId"`

	// 日志搜索的关键词，默认为空
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 日志搜索的游标，可透传上次返回的值，默认为空
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 时间戳排序规则，asc - 升序，desc - 降序。默认为升序
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 搜索的日志条数上限值，最大为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeJobSubmissionLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobSubmissionLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RunningOrderId")
	delete(f, "Keyword")
	delete(f, "Cursor")
	delete(f, "OrderType")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobSubmissionLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobSubmissionLogResponseParams struct {
	// 日志搜索的游标，需要搜索更多时透传这个值
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 是否返回了所有的日志记录
	ListOver *bool `json:"ListOver,omitnil,omitempty" name:"ListOver"`

	// 作业启动的requestId
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobRequestId *string `json:"JobRequestId,omitnil,omitempty" name:"JobRequestId"`

	// 该时间段内符合关键字的所有的作业实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobInstanceList []*JobInstanceForSubmissionLog `json:"JobInstanceList,omitnil,omitempty" name:"JobInstanceList"`

	// 废弃，请使用LogContentList
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogList []*string `json:"LogList,omitnil,omitempty" name:"LogList"`

	// 日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogContentList []*LogContent `json:"LogContentList,omitnil,omitempty" name:"LogContentList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeJobSubmissionLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobSubmissionLogResponseParams `json:"Response"`
}

func (r *DescribeJobSubmissionLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobSubmissionLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobsRequestParams struct {
	// 按照一个或者多个作业ID查询。作业ID形如：cql-11112222，每次请求的作业上限为100。参数不支持同时指定JobIds和Filters。
	JobIds []*string `json:"JobIds,omitnil,omitempty" name:"JobIds"`

	// 过滤条件，支持的 Filter.Name 为：作业名 Name、作业状态 Status、所属集群 ClusterId、作业id JobId、集群名称 ClusterName。 每次请求的 Filters 个数的上限为 5，Filter.Values 的个数上限为 5。参数不支持同时指定 JobIds 和 Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 查询额外的作业信息,例如 JobEventInfo	
	ExtraResult []*string `json:"ExtraResult,omitnil,omitempty" name:"ExtraResult"`
}

type DescribeJobsRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个作业ID查询。作业ID形如：cql-11112222，每次请求的作业上限为100。参数不支持同时指定JobIds和Filters。
	JobIds []*string `json:"JobIds,omitnil,omitempty" name:"JobIds"`

	// 过滤条件，支持的 Filter.Name 为：作业名 Name、作业状态 Status、所属集群 ClusterId、作业id JobId、集群名称 ClusterName。 每次请求的 Filters 个数的上限为 5，Filter.Values 的个数上限为 5。参数不支持同时指定 JobIds 和 Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 查询额外的作业信息,例如 JobEventInfo	
	ExtraResult []*string `json:"ExtraResult,omitnil,omitempty" name:"ExtraResult"`
}

func (r *DescribeJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "WorkSpaceId")
	delete(f, "ExtraResult")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobsResponseParams struct {
	// 作业总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 作业列表
	JobSet []*JobV1 `json:"JobSet,omitnil,omitempty" name:"JobSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobsResponseParams `json:"Response"`
}

func (r *DescribeJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceConfigsRequestParams struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 偏移量，仅当设置 Limit 时该参数有效
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回值大小，不填则返回全量数据
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 资源配置Versions集合
	ResourceConfigVersions []*int64 `json:"ResourceConfigVersions,omitnil,omitempty" name:"ResourceConfigVersions"`

	// 作业配置版本
	JobConfigVersion *int64 `json:"JobConfigVersion,omitnil,omitempty" name:"JobConfigVersion"`

	// 作业ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type DescribeResourceConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 偏移量，仅当设置 Limit 时该参数有效
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回值大小，不填则返回全量数据
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 资源配置Versions集合
	ResourceConfigVersions []*int64 `json:"ResourceConfigVersions,omitnil,omitempty" name:"ResourceConfigVersions"`

	// 作业配置版本
	JobConfigVersion *int64 `json:"JobConfigVersion,omitnil,omitempty" name:"JobConfigVersion"`

	// 作业ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *DescribeResourceConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ResourceConfigVersions")
	delete(f, "JobConfigVersion")
	delete(f, "JobId")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceConfigsResponseParams struct {
	// 资源配置描述数组
	ResourceConfigSet []*ResourceConfigItem `json:"ResourceConfigSet,omitnil,omitempty" name:"ResourceConfigSet"`

	// 资源配置数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourceConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceConfigsResponseParams `json:"Response"`
}

func (r *DescribeResourceConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceRelatedJobsRequestParams struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 默认0;   1： 按照作业版本创建时间降序
	DESCByJobConfigCreateTime *int64 `json:"DESCByJobConfigCreateTime,omitnil,omitempty" name:"DESCByJobConfigCreateTime"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 资源版本号
	ResourceConfigVersion *int64 `json:"ResourceConfigVersion,omitnil,omitempty" name:"ResourceConfigVersion"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type DescribeResourceRelatedJobsRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 默认0;   1： 按照作业版本创建时间降序
	DESCByJobConfigCreateTime *int64 `json:"DESCByJobConfigCreateTime,omitnil,omitempty" name:"DESCByJobConfigCreateTime"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 资源版本号
	ResourceConfigVersion *int64 `json:"ResourceConfigVersion,omitnil,omitempty" name:"ResourceConfigVersion"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *DescribeResourceRelatedJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceRelatedJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "DESCByJobConfigCreateTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ResourceConfigVersion")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceRelatedJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceRelatedJobsResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 关联作业信息
	RefJobInfos []*ResourceRefJobInfo `json:"RefJobInfos,omitnil,omitempty" name:"RefJobInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourceRelatedJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceRelatedJobsResponseParams `json:"Response"`
}

func (r *DescribeResourceRelatedJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceRelatedJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesRequestParams struct {
	// 需要查询的资源ID数组，数量不超过100个。如果填写了该参数则忽略Filters参数。
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 偏移量，仅当设置 Limit 参数时有效
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 条数限制。如果不填，默认返回 20 条
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <li><strong>ResourceName</strong></li>
	// <p style="padding-left: 30px;">按照资源名字过滤，支持模糊过滤。传入的过滤名字不超过5个</p><p style="padding-left: 30px;">类型: String</p><p style="padding-left: 30px;">必选: 否</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type DescribeResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的资源ID数组，数量不超过100个。如果填写了该参数则忽略Filters参数。
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 偏移量，仅当设置 Limit 参数时有效
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 条数限制。如果不填，默认返回 20 条
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <li><strong>ResourceName</strong></li>
	// <p style="padding-left: 30px;">按照资源名字过滤，支持模糊过滤。传入的过滤名字不超过5个</p><p style="padding-left: 30px;">类型: String</p><p style="padding-left: 30px;">必选: 否</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *DescribeResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesResponseParams struct {
	// 资源详细信息集合
	ResourceSet []*ResourceItem `json:"ResourceSet,omitnil,omitempty" name:"ResourceSet"`

	// 总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcesResponseParams `json:"Response"`
}

func (r *DescribeResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSystemResourcesRequestParams struct {
	// 需要查询的资源ID数组
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 偏移量，仅当设置 Limit 参数时有效
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 条数限制，默认返回 20 条
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询资源配置列表， 如果不填写，返回该 ResourceIds.N 下所有作业配置列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 查询对应Flink版本的内置connector
	FlinkVersion *string `json:"FlinkVersion,omitnil,omitempty" name:"FlinkVersion"`
}

type DescribeSystemResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的资源ID数组
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 偏移量，仅当设置 Limit 参数时有效
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 条数限制，默认返回 20 条
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询资源配置列表， 如果不填写，返回该 ResourceIds.N 下所有作业配置列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 查询对应Flink版本的内置connector
	FlinkVersion *string `json:"FlinkVersion,omitnil,omitempty" name:"FlinkVersion"`
}

func (r *DescribeSystemResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSystemResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "ClusterId")
	delete(f, "FlinkVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSystemResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSystemResourcesResponseParams struct {
	// 资源详细信息集合
	ResourceSet []*SystemResourceItem `json:"ResourceSet,omitnil,omitempty" name:"ResourceSet"`

	// 总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSystemResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSystemResourcesResponseParams `json:"Response"`
}

func (r *DescribeSystemResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSystemResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTreeJobsRequestParams struct {
	// 筛选条件字段
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 工作空间 Serialid
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type DescribeTreeJobsRequest struct {
	*tchttp.BaseRequest
	
	// 筛选条件字段
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 工作空间 Serialid
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *DescribeTreeJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTreeJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTreeJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTreeJobsResponseParams struct {
	// 父节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 当前文件夹ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 当前文件夹名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 当前文件夹下的作业列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobSet []*TreeJobSets `json:"JobSet,omitnil,omitempty" name:"JobSet"`

	// 迭代子目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Children []*DescribeTreeJobsRsp `json:"Children,omitnil,omitempty" name:"Children"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTreeJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTreeJobsResponseParams `json:"Response"`
}

func (r *DescribeTreeJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTreeJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTreeJobsRsp struct {
	// 父节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 当前文件夹ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 当前文件夹名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 当前文件夹下的作业集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobSet []*TreeJobSets `json:"JobSet,omitnil,omitempty" name:"JobSet"`

	// 迭代子目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Children []*DescribeTreeJobsRsp `json:"Children,omitnil,omitempty" name:"Children"`

	// 请求ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

// Predefined struct for user
type DescribeTreeResourcesRequestParams struct {
	// 筛选条件字段
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 分页游标
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 单页显示数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTreeResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 筛选条件字段
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 分页游标
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 单页显示数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeTreeResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTreeResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "WorkSpaceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTreeResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTreeResourcesResponseParams struct {
	// 父节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 文件夹ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 文件夹名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 文件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*TreeResourceItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 子目录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Children []*DescribeTreeResourcesRsp `json:"Children,omitnil,omitempty" name:"Children"`

	// 资源总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTreeResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTreeResourcesResponseParams `json:"Response"`
}

func (r *DescribeTreeResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTreeResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTreeResourcesRsp struct {
	// 父节点ID
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 文件夹ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 文件夹名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 文件夹下资源数字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*TreeResourceItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 子节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Children []*DescribeTreeResourcesRsp `json:"Children,omitnil,omitempty" name:"Children"`

	// 资源总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type DescribeWorkSpacesRequestParams struct {
	// 偏移量，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 1 按照创建时间降序排序(默认) 2.按照创建时间升序排序，3. 按照状态降序排序 4. 按照状态升序排序 默认为0
	OrderType *int64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 请求的集群数量，默认 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤规则
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeWorkSpacesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 1 按照创建时间降序排序(默认) 2.按照创建时间升序排序，3. 按照状态降序排序 4. 按照状态升序排序 默认为0
	OrderType *int64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 请求的集群数量，默认 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤规则
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeWorkSpacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkSpacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "OrderType")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkSpacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkSpacesResponseParams struct {
	// 空间详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkSpaceSetItem []*WorkSpaceSetItem `json:"WorkSpaceSetItem,omitnil,omitempty" name:"WorkSpaceSetItem"`

	// 空间总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkSpacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkSpacesResponseParams `json:"Response"`
}

func (r *DescribeWorkSpacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkSpacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpertModeConfiguration struct {
	// Job graph
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobGraph *JobGraph `json:"JobGraph,omitnil,omitempty" name:"JobGraph"`

	// Node configuration
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeConfig []*NodeConfig `json:"NodeConfig,omitnil,omitempty" name:"NodeConfig"`

	// Slot sharing groups
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotSharingGroups []*SlotSharingGroup `json:"SlotSharingGroups,omitnil,omitempty" name:"SlotSharingGroups"`
}

// Predefined struct for user
type FetchSqlGatewayStatementResultRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Sql Gateway会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// sql的查询id
	OperationHandleId *string `json:"OperationHandleId,omitnil,omitempty" name:"OperationHandleId"`

	// 下一条结果的获取url，首次获取执行结果时可以为空，当获取下一批查询结果时需要传递
	ResultUri *string `json:"ResultUri,omitnil,omitempty" name:"ResultUri"`
}

type FetchSqlGatewayStatementResultRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Sql Gateway会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// sql的查询id
	OperationHandleId *string `json:"OperationHandleId,omitnil,omitempty" name:"OperationHandleId"`

	// 下一条结果的获取url，首次获取执行结果时可以为空，当获取下一批查询结果时需要传递
	ResultUri *string `json:"ResultUri,omitnil,omitempty" name:"ResultUri"`
}

func (r *FetchSqlGatewayStatementResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchSqlGatewayStatementResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SessionId")
	delete(f, "OperationHandleId")
	delete(f, "ResultUri")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FetchSqlGatewayStatementResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FetchSqlGatewayStatementResultResponseParams struct {
	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage []*string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 返回类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultType *string `json:"ResultType,omitnil,omitempty" name:"ResultType"`

	// 是否DQL结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsQueryResult *bool `json:"IsQueryResult,omitnil,omitempty" name:"IsQueryResult"`

	// 结果类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultKind *string `json:"ResultKind,omitnil,omitempty" name:"ResultKind"`

	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results *StatementResult `json:"Results,omitnil,omitempty" name:"Results"`

	// 下一次请求的uri
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextResultUri *string `json:"NextResultUri,omitnil,omitempty" name:"NextResultUri"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FetchSqlGatewayStatementResultResponse struct {
	*tchttp.BaseResponse
	Response *FetchSqlGatewayStatementResultResponseParams `json:"Response"`
}

func (r *FetchSqlGatewayStatementResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchSqlGatewayStatementResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 要过滤的字段
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段的过滤值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type GatewayRefItem struct {
	// 空间唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 资源唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`

	// 引用类型，0:用户资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type GetMetaTableRequestParams struct {
	// 目录名
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// 库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 表名
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 空间唯一标识
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type GetMetaTableRequest struct {
	*tchttp.BaseRequest
	
	// 目录名
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// 库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 表名
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 空间唯一标识
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *GetMetaTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMetaTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Catalog")
	delete(f, "Database")
	delete(f, "Table")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetMetaTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetMetaTableResponseParams struct {
	// 元数据表唯一标识
	SerialId *string `json:"SerialId,omitnil,omitempty" name:"SerialId"`

	// 目录名
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// 库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 表名
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 建表语句,使用 Base64 编码。
	// 例如
	// Q1JFQVRFIFRBQkxFIGRhdGFnZW5fc291cmNlX3RhYmxlICggCiAgICBpZCBJTlQsIAogICAgbmFtZSBTVFJJTkcgCikgV0lUSCAoCidjb25uZWN0b3InPSdkYXRhZ2VuJywKJ3Jvd3MtcGVyLXNlY29uZCcgPSAnMScKKTs=
	DDL *string `json:"DDL,omitnil,omitempty" name:"DDL"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetMetaTableResponse struct {
	*tchttp.BaseResponse
	Response *GetMetaTableResponseParams `json:"Response"`
}

func (r *GetMetaTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMetaTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JobConfig struct {
	// 作业Id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 主类
	// 注意：此字段可能返回 null，表示取不到有效值。
	EntrypointClass *string `json:"EntrypointClass,omitnil,omitempty" name:"EntrypointClass"`

	// 主类入参
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramArgs *string `json:"ProgramArgs,omitnil,omitempty" name:"ProgramArgs"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 作业配置创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 作业配置的版本号
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`

	// 作业默认并行度
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultParallelism *uint64 `json:"DefaultParallelism,omitnil,omitempty" name:"DefaultParallelism"`

	// 系统参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 引用资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceRefDetails []*ResourceRefDetail `json:"ResourceRefDetails,omitnil,omitempty" name:"ResourceRefDetails"`

	// 创建者uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorUin *string `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 作业配置上次启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 作业绑定的存储桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	COSBucket *string `json:"COSBucket,omitnil,omitempty" name:"COSBucket"`

	// 是否启用日志收集，0-未启用，1-已启用，2-历史集群未设置日志集，3-历史集群已开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogCollect *int64 `json:"LogCollect,omitnil,omitempty" name:"LogCollect"`

	// 作业的最大并行度
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxParallelism *uint64 `json:"MaxParallelism,omitnil,omitempty" name:"MaxParallelism"`

	// JobManager规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobManagerSpec *float64 `json:"JobManagerSpec,omitnil,omitempty" name:"JobManagerSpec"`

	// TaskManager规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskManagerSpec *float64 `json:"TaskManagerSpec,omitnil,omitempty" name:"TaskManagerSpec"`

	// CLS日志集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClsLogsetId *string `json:"ClsLogsetId,omitnil,omitempty" name:"ClsLogsetId"`

	// CLS日志主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClsTopicId *string `json:"ClsTopicId,omitnil,omitempty" name:"ClsTopicId"`

	// pyflink作业运行的python版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PythonVersion *string `json:"PythonVersion,omitnil,omitempty" name:"PythonVersion"`

	// Oceanus 平台恢复作业开关 1:开启 -1: 关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRecover *int64 `json:"AutoRecover,omitnil,omitempty" name:"AutoRecover"`

	// 日志级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogLevel *string `json:"LogLevel,omitnil,omitempty" name:"LogLevel"`

	// 类日志级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClazzLevels []*ClazzLevel `json:"ClazzLevels,omitnil,omitempty" name:"ClazzLevels"`

	// 是否开启专家模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpertModeOn *bool `json:"ExpertModeOn,omitnil,omitempty" name:"ExpertModeOn"`

	// 专家模式的配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpertModeConfiguration *ExpertModeConfiguration `json:"ExpertModeConfiguration,omitnil,omitempty" name:"ExpertModeConfiguration"`

	// trace链路
	// 注意：此字段可能返回 null，表示取不到有效值。
	TraceModeOn *bool `json:"TraceModeOn,omitnil,omitempty" name:"TraceModeOn"`

	// trace链路配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TraceModeConfiguration *TraceModeConfiguration `json:"TraceModeConfiguration,omitnil,omitempty" name:"TraceModeConfiguration"`

	// checkpoint保留个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckpointRetainedNum *int64 `json:"CheckpointRetainedNum,omitnil,omitempty" name:"CheckpointRetainedNum"`

	// 算子拓扑图
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobGraph *JobGraph `json:"JobGraph,omitnil,omitempty" name:"JobGraph"`

	// es索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsServerlessIndex *string `json:"EsServerlessIndex,omitnil,omitempty" name:"EsServerlessIndex"`

	// es空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsServerlessSpace *string `json:"EsServerlessSpace,omitnil,omitempty" name:"EsServerlessSpace"`

	// es索引中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// es空间中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceName *string `json:"WorkspaceName,omitnil,omitempty" name:"WorkspaceName"`

	// flink 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlinkVersion *string `json:"FlinkVersion,omitnil,omitempty" name:"FlinkVersion"`
}

type JobEvent struct {
	// 内部定义的事件类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 事件类型的说明文字
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 事件发生的 Unix 时间戳（秒）
	Timestamp *uint64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 事件发生时的运行 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningOrderId *uint64 `json:"RunningOrderId,omitnil,omitempty" name:"RunningOrderId"`

	// 事件的一些可选说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 异常事件的排查手册链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	SolutionLink *string `json:"SolutionLink,omitnil,omitempty" name:"SolutionLink"`
}

type JobEventInfo struct {
	// 异常事件总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorEventTotal *int64 `json:"ErrorEventTotal,omitnil,omitempty" name:"ErrorEventTotal"`
}

type JobGraph struct {
	// 运行图的点集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nodes []*JobGraphNode `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// 运行图的边集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Edges []*JobGraphEdge `json:"Edges,omitnil,omitempty" name:"Edges"`
}

type JobGraphEdge struct {
	// 边的起始节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 边的目标节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Target *int64 `json:"Target,omitnil,omitempty" name:"Target"`
}

type JobGraphNode struct {
	// 节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 节点描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 节点并行度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Parallelism *int64 `json:"Parallelism,omitnil,omitempty" name:"Parallelism"`
}

type JobInstanceForSubmissionLog struct {
	// 实例的Id, 按照启动的时间顺序，从1开始
	RunningOrderId *int64 `json:"RunningOrderId,omitnil,omitempty" name:"RunningOrderId"`

	// 作业实例的启动时间
	JobInstanceStartTime *string `json:"JobInstanceStartTime,omitnil,omitempty" name:"JobInstanceStartTime"`

	// 作业实例启动的时间（毫秒）
	StartingMillis *int64 `json:"StartingMillis,omitnil,omitempty" name:"StartingMillis"`
}

type JobRuntimeInfo struct {
	// 运行信息的key，目前支持：TaskManagers：taskmanager pod 列表； StreamGraph：作业对应的 StreamGraph；SubTasks：作业的 subtask 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 运行信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type JobV1 struct {
	// 作业ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 用户AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 创建者UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorUin *string `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 作业名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 作业类型，1：sql作业，2：Jar作业
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobType *int64 `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 作业状态，1：未初始化，2：未发布，3：操作中，4：运行中，5：停止，6：暂停，-1：故障
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 作业创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 作业启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 作业停止时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StopTime *string `json:"StopTime,omitnil,omitempty" name:"StopTime"`

	// 作业更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 作业累计运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalRunMillis *int64 `json:"TotalRunMillis,omitnil,omitempty" name:"TotalRunMillis"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 操作错误提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOpResult *string `json:"LastOpResult,omitnil,omitempty" name:"LastOpResult"`

	// 集群名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 最新配置版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestJobConfigVersion *int64 `json:"LatestJobConfigVersion,omitnil,omitempty" name:"LatestJobConfigVersion"`

	// 已发布的配置版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublishedJobConfigVersion *int64 `json:"PublishedJobConfigVersion,omitnil,omitempty" name:"PublishedJobConfigVersion"`

	// 运行的CU数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningCuNum *int64 `json:"RunningCuNum,omitnil,omitempty" name:"RunningCuNum"`

	// 作业内存规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	CuMem *int64 `json:"CuMem,omitnil,omitempty" name:"CuMem"`

	// 作业状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 运行状态时表示单次运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentRunMillis *int64 `json:"CurrentRunMillis,omitnil,omitempty" name:"CurrentRunMillis"`

	// 作业所在的集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 作业管理WEB UI 入口
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebUIUrl *string `json:"WebUIUrl,omitnil,omitempty" name:"WebUIUrl"`

	// 作业所在集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerType *int64 `json:"SchedulerType,omitnil,omitempty" name:"SchedulerType"`

	// 作业所在集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterStatus *int64 `json:"ClusterStatus,omitnil,omitempty" name:"ClusterStatus"`

	// 细粒度下的运行的CU数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningCu *float64 `json:"RunningCu,omitnil,omitempty" name:"RunningCu"`

	// 作业运行的 Flink 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlinkVersion *string `json:"FlinkVersion,omitnil,omitempty" name:"FlinkVersion"`

	// 工作空间 SerialId
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 工作空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkSpaceName *string `json:"WorkSpaceName,omitnil,omitempty" name:"WorkSpaceName"`

	// 作业标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 作业异常事件信息	
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventInfo *JobEventInfo `json:"EventInfo,omitnil,omitempty" name:"EventInfo"`

	// 描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 0:代表没开启调优任务，1:开启智能调优，2:代表定时调优
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScalingType *int64 `json:"ScalingType,omitnil,omitempty" name:"ScalingType"`

	// 使用CPU数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningCpu *float64 `json:"RunningCpu,omitnil,omitempty" name:"RunningCpu"`

	// 使用内存数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningMem *float64 `json:"RunningMem,omitnil,omitempty" name:"RunningMem"`
}

type LogContent struct {
	// 日志内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Log *string `json:"Log,omitnil,omitempty" name:"Log"`

	// 毫秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 日志组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgId *string `json:"PkgId,omitnil,omitempty" name:"PkgId"`

	// 日志Id，在日志组范围里唯一
	PkgLogId *int64 `json:"PkgLogId,omitnil,omitempty" name:"PkgLogId"`

	// 日志所属的容器名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerName *string `json:"ContainerName,omitnil,omitempty" name:"ContainerName"`
}

type LogicalType struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否允许为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	NullAble *bool `json:"NullAble,omitnil,omitempty" name:"NullAble"`

	// 长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Length *int64 `json:"Length,omitnil,omitempty" name:"Length"`
}

// Predefined struct for user
type ModifyFolderRequestParams struct {
	// 文件夹ID（必填）
	SourceFolderId *string `json:"SourceFolderId,omitnil,omitempty" name:"SourceFolderId"`

	// 如需拖拽文件夹，需传入目标文件夹ID
	TargetFolderId *string `json:"TargetFolderId,omitnil,omitempty" name:"TargetFolderId"`

	// 如需修改文件夹名，需传入FolderName字段
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 文件夹类型，0是任务文件夹，1是依赖文件夹
	FolderType *int64 `json:"FolderType,omitnil,omitempty" name:"FolderType"`

	// 批量移动的作业serial id 列表
	SourceJobIds []*string `json:"SourceJobIds,omitnil,omitempty" name:"SourceJobIds"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type ModifyFolderRequest struct {
	*tchttp.BaseRequest
	
	// 文件夹ID（必填）
	SourceFolderId *string `json:"SourceFolderId,omitnil,omitempty" name:"SourceFolderId"`

	// 如需拖拽文件夹，需传入目标文件夹ID
	TargetFolderId *string `json:"TargetFolderId,omitnil,omitempty" name:"TargetFolderId"`

	// 如需修改文件夹名，需传入FolderName字段
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 文件夹类型，0是任务文件夹，1是依赖文件夹
	FolderType *int64 `json:"FolderType,omitnil,omitempty" name:"FolderType"`

	// 批量移动的作业serial id 列表
	SourceJobIds []*string `json:"SourceJobIds,omitnil,omitempty" name:"SourceJobIds"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *ModifyFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceFolderId")
	delete(f, "TargetFolderId")
	delete(f, "FolderName")
	delete(f, "FolderType")
	delete(f, "SourceJobIds")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFolderResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyFolderResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFolderResponseParams `json:"Response"`
}

func (r *ModifyFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyJobRequestParams struct {
	// 作业Id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 作业名称，支持长度小于50的中文/英文/数字/”-”/”_”/”.”，不能重名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 拖拽文件需传入此参数
	TargetFolderId *string `json:"TargetFolderId,omitnil,omitempty" name:"TargetFolderId"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 作业描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyJobRequest struct {
	*tchttp.BaseRequest
	
	// 作业Id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 作业名称，支持长度小于50的中文/英文/数字/”-”/”_”/”.”，不能重名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 拖拽文件需传入此参数
	TargetFolderId *string `json:"TargetFolderId,omitnil,omitempty" name:"TargetFolderId"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 作业描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "Name")
	delete(f, "Remark")
	delete(f, "TargetFolderId")
	delete(f, "WorkSpaceId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyJobResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyJobResponse struct {
	*tchttp.BaseResponse
	Response *ModifyJobResponseParams `json:"Response"`
}

func (r *ModifyJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkSpaceRequestParams struct {
	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 待修改的工作空间名称
	WorkSpaceName *string `json:"WorkSpaceName,omitnil,omitempty" name:"WorkSpaceName"`

	// 待修改的工作空间备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyWorkSpaceRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 待修改的工作空间名称
	WorkSpaceName *string `json:"WorkSpaceName,omitnil,omitempty" name:"WorkSpaceName"`

	// 待修改的工作空间备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyWorkSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkSpaceId")
	delete(f, "WorkSpaceName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWorkSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkSpaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyWorkSpaceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWorkSpaceResponseParams `json:"Response"`
}

func (r *ModifyWorkSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeConfig struct {
	// Node ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Node parallelism
	// 注意：此字段可能返回 null，表示取不到有效值。
	Parallelism *int64 `json:"Parallelism,omitnil,omitempty" name:"Parallelism"`

	// Slot sharing group
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotSharingGroup *string `json:"SlotSharingGroup,omitnil,omitempty" name:"SlotSharingGroup"`

	// Configuration properties
	// 注意：此字段可能返回 null，表示取不到有效值。
	Configuration []*Property `json:"Configuration,omitnil,omitempty" name:"Configuration"`

	// 节点的状态ttl配置, 多个用 ; 分割
	// 注意：此字段可能返回 null，表示取不到有效值。
	StateTTL *string `json:"StateTTL,omitnil,omitempty" name:"StateTTL"`
}

type Order struct {
	// 创建、续费、扩缩容 1 2 3
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 自动续费 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 操作人的UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// 最终集群的CU数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComputeCu *int64 `json:"ComputeCu,omitnil,omitempty" name:"ComputeCu"`

	// 订单的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderTime *string `json:"OrderTime,omitnil,omitempty" name:"OrderTime"`
}

type Property struct {
	// 系统配置的Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 系统配置的Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type RefJobStatusCountItem struct {
	// 作业状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobStatus *int64 `json:"JobStatus,omitnil,omitempty" name:"JobStatus"`

	// 作业数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type ResourceConfigItem struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源类型
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源所属地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 资源所属AppId
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 主账号Uin
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 子账号Uin
	CreatorUin *string `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 资源位置描述
	ResourceLoc *ResourceLoc `json:"ResourceLoc,omitnil,omitempty" name:"ResourceLoc"`

	// 资源创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 资源版本
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`

	// 资源描述
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 资源状态：0: 资源同步中，1:资源已就绪
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 关联作业个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefJobCount *int64 `json:"RefJobCount,omitnil,omitempty" name:"RefJobCount"`

	// 分状态统计关联作业数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefJobStatusCountSet []*RefJobStatusCountItem `json:"RefJobStatusCountSet,omitnil,omitempty" name:"RefJobStatusCountSet"`
}

type ResourceItem struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资源类型
	ResourceType *uint64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源位置
	ResourceLoc *ResourceLoc `json:"ResourceLoc,omitnil,omitempty" name:"ResourceLoc"`

	// 资源地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 应用ID
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 主账号Uin
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 子账号Uin
	CreatorUin *string `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 资源创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 资源最后更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 资源的资源版本ID
	LatestResourceConfigVersion *int64 `json:"LatestResourceConfigVersion,omitnil,omitempty" name:"LatestResourceConfigVersion"`

	// 资源备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 版本个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionCount *int64 `json:"VersionCount,omitnil,omitempty" name:"VersionCount"`

	// 关联作业数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefJobCount *int64 `json:"RefJobCount,omitnil,omitempty" name:"RefJobCount"`

	// 作业运行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsJobRun *int64 `json:"IsJobRun,omitnil,omitempty" name:"IsJobRun"`

	// 文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 工作空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkSpaceId *int64 `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 分状态统计关联作业数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefJobStatusCountSet []*RefJobStatusCountItem `json:"RefJobStatusCountSet,omitnil,omitempty" name:"RefJobStatusCountSet"`
}

type ResourceLoc struct {
	// 资源位置的存储类型，目前只支持1:COS
	StorageType *int64 `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 描述资源位置的json
	Param *ResourceLocParam `json:"Param,omitnil,omitempty" name:"Param"`
}

type ResourceLocParam struct {
	// 资源bucket
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 资源路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 资源所在地域，如果不填，则使用Resource的Region
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type ResourceRef struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源版本ID，-1表示使用最新版本
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`

	// 引用资源类型，例如主资源设置为1，代表main class所在的jar包
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type ResourceRefDetail struct {
	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源版本，-1表示使用最新版本
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`

	// 资源名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 1: 主资源
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 1: 系统内置资源
	SystemProvide *int64 `json:"SystemProvide,omitnil,omitempty" name:"SystemProvide"`
}

type ResourceRefJobInfo struct {
	// Job id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Job配置版本
	JobConfigVersion *int64 `json:"JobConfigVersion,omitnil,omitempty" name:"JobConfigVersion"`

	// 资源版本
	ResourceVersion *int64 `json:"ResourceVersion,omitnil,omitempty" name:"ResourceVersion"`
}

type ResultColumn struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 本地类型描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogicalType *LogicalType `json:"LogicalType,omitnil,omitempty" name:"LogicalType"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type ResultData struct {
	// 操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`
}

type RoleAuth struct {
	// 用户 AppID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 工作空间 SerialId
	WorkSpaceSerialId *string `json:"WorkSpaceSerialId,omitnil,omitempty" name:"WorkSpaceSerialId"`

	// 主账号 UIN
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 创建者 UIN
	CreatorUin *string `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 绑定授权的 UIN
	AuthSubAccountUin *string `json:"AuthSubAccountUin,omitnil,omitempty" name:"AuthSubAccountUin"`

	// 对应 role表的id
	Permission *int64 `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后一次操作时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 2 启用 1 停用
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 工作空间id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkSpaceId *int64 `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 权限名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

type RunJobDescription struct {
	// 作业Id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 运行类型，1：启动，2：恢复
	RunType *int64 `json:"RunType,omitnil,omitempty" name:"RunType"`

	// 兼容旧版 SQL 类型作业启动参数：指定数据源消费起始时间点（建议传值）
	// 保证参数为 LATEST、EARLIEST、T+Timestamp （例:T1557394288000）
	StartMode *string `json:"StartMode,omitnil,omitempty" name:"StartMode"`

	// 当前作业的某个版本
	// （不传默认为非草稿的作业版本）
	JobConfigVersion *uint64 `json:"JobConfigVersion,omitnil,omitempty" name:"JobConfigVersion"`

	// Savepoint路径
	SavepointPath *string `json:"SavepointPath,omitnil,omitempty" name:"SavepointPath"`

	// Savepoint的Id
	SavepointId *string `json:"SavepointId,omitnil,omitempty" name:"SavepointId"`

	// 使用历史版本系统依赖
	UseOldSystemConnector *bool `json:"UseOldSystemConnector,omitnil,omitempty" name:"UseOldSystemConnector"`

	// 自定义时间戳
	CustomTimestamp *int64 `json:"CustomTimestamp,omitnil,omitempty" name:"CustomTimestamp"`

	// timestamp; latest-offset;  earliest-offset; 任选一种
	KafkaScanMode *string `json:"KafkaScanMode,omitnil,omitempty" name:"KafkaScanMode"`
}

// Predefined struct for user
type RunJobsRequestParams struct {
	// 批量启动作业的描述信息
	RunJobDescriptions []*RunJobDescription `json:"RunJobDescriptions,omitnil,omitempty" name:"RunJobDescriptions"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type RunJobsRequest struct {
	*tchttp.BaseRequest
	
	// 批量启动作业的描述信息
	RunJobDescriptions []*RunJobDescription `json:"RunJobDescriptions,omitnil,omitempty" name:"RunJobDescriptions"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *RunJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RunJobDescriptions")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunJobsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunJobsResponse struct {
	*tchttp.BaseResponse
	Response *RunJobsResponseParams `json:"Response"`
}

func (r *RunJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunSqlGatewayStatementRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 需要执行的sql，该sql会被Sql Gateway执行，当前支持的是paimon修改需求，因此主要是DDL语句
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Sql Gateway会话ID，可不填，如果不填则会自动创建一个会话ID，每个会话ID都有一个存活时间，测试环境为10分钟，线上默认是30分钟
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type RunSqlGatewayStatementRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 需要执行的sql，该sql会被Sql Gateway执行，当前支持的是paimon修改需求，因此主要是DDL语句
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Sql Gateway会话ID，可不填，如果不填则会自动创建一个会话ID，每个会话ID都有一个存活时间，测试环境为10分钟，线上默认是30分钟
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *RunSqlGatewayStatementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunSqlGatewayStatementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Sql")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunSqlGatewayStatementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunSqlGatewayStatementResponseParams struct {
	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage []*string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 会话id，若入参未传，则返回自动创建的会话id，若入参已经传递，则返回值与原传入值一致
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 返回执行id，可以根据该执行id和会话id获取执行结果
	OperationHandleId *string `json:"OperationHandleId,omitnil,omitempty" name:"OperationHandleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunSqlGatewayStatementResponse struct {
	*tchttp.BaseResponse
	Response *RunSqlGatewayStatementResponseParams `json:"Response"`
}

func (r *RunSqlGatewayStatementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunSqlGatewayStatementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Savepoint struct {
	// 主键
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionId *int64 `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// 状态 1: Active; 2: Expired; 3: InProgress; 4: Failed; 5: Timeout
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 快照类型 1: savepoint；2: checkpoint；3: cancelWithSavepoint
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordType *int64 `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 运行作业实例的顺序 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobRuntimeId *int64 `json:"JobRuntimeId,omitnil,omitempty" name:"JobRuntimeId"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 固定超时时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 快照 serialId
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerialId *string `json:"SerialId,omitnil,omitempty" name:"SerialId"`

	// 耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeConsuming *int64 `json:"TimeConsuming,omitnil,omitempty" name:"TimeConsuming"`

	// 快照路径状态 1：可用；2：不可用；
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathStatus *int64 `json:"PathStatus,omitnil,omitempty" name:"PathStatus"`
}

type SessionClusterRefItem struct {
	// 空间唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 资源唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`

	// 引用类型，0:用户资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type SlotSharingGroup struct {
	// SlotSharingGroup的名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// SlotSharingGroup的规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spec *SlotSharingGroupSpec `json:"Spec,omitnil,omitempty" name:"Spec"`

	// SlotSharingGroup的描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// SlotSharingGroup的配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Configuration []*Property `json:"Configuration,omitnil,omitempty" name:"Configuration"`
}

type SlotSharingGroupSpec struct {
	// 适用的cpu
	// 注意：此字段可能返回 null，表示取不到有效值。
	CPU *float64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// 默认为b, 支持单位有 b, kb, mb, gb
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeapMemory *string `json:"HeapMemory,omitnil,omitempty" name:"HeapMemory"`

	// 默认为b, 支持单位有 b, kb, mb, gb
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffHeapMemory *string `json:"OffHeapMemory,omitnil,omitempty" name:"OffHeapMemory"`

	// 默认为b, 支持单位有 b, kb, mb, gb
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManagedMemory *string `json:"ManagedMemory,omitnil,omitempty" name:"ManagedMemory"`
}

type SqlGatewayItem struct {
	// 唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerialId *string `json:"SerialId,omitnil,omitempty" name:"SerialId"`

	// Flink内核版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlinkVersion *string `json:"FlinkVersion,omitnil,omitempty" name:"FlinkVersion"`

	// 状态，1.停止 2. 开启中 3. 开启 4. 开启失败 5. 停止中
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorUin *string `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 引用资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceRefs []*GatewayRefItem `json:"ResourceRefs,omitnil,omitempty" name:"ResourceRefs"`

	// Cu规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	CuSpec *float64 `json:"CuSpec,omitnil,omitempty" name:"CuSpec"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 配置参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`
}

type StatementResult struct {
	// 返回结果列
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*ResultColumn `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	RowFormat *string `json:"RowFormat,omitnil,omitempty" name:"RowFormat"`

	// 结果值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ResultData `json:"Data,omitnil,omitempty" name:"Data"`
}

type StopJobDescription struct {
	// 作业Id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 停止类型，1 停止 2 暂停
	StopType *int64 `json:"StopType,omitnil,omitempty" name:"StopType"`
}

// Predefined struct for user
type StopJobsRequestParams struct {
	// 批量停止作业的描述信息
	StopJobDescriptions []*StopJobDescription `json:"StopJobDescriptions,omitnil,omitempty" name:"StopJobDescriptions"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type StopJobsRequest struct {
	*tchttp.BaseRequest
	
	// 批量停止作业的描述信息
	StopJobDescriptions []*StopJobDescription `json:"StopJobDescriptions,omitnil,omitempty" name:"StopJobDescriptions"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *StopJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StopJobDescriptions")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopJobsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopJobsResponse struct {
	*tchttp.BaseResponse
	Response *StopJobsResponseParams `json:"Response"`
}

func (r *StopJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubEks struct {
	// 集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerialId *string `json:"SerialId,omitnil,omitempty" name:"SerialId"`

	// cu数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CuNum *int64 `json:"CuNum,omitnil,omitempty" name:"CuNum"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 运行cu
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningCu *float64 `json:"RunningCu,omitnil,omitempty" name:"RunningCu"`
}

type SubFolderInfo struct {
	// folder id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// folder name
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

type SystemResourceItem struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资源类型。1 表示 JAR 包，目前只支持该值。
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 资源所属地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 资源的最新版本
	LatestResourceConfigVersion *int64 `json:"LatestResourceConfigVersion,omitnil,omitempty" name:"LatestResourceConfigVersion"`
}

type Tag struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TraceModeConfiguration struct {
	// 如1%转换为0.01
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rate *string `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 按照算子ID顺序配置，可以对每个算子配置IN、OUT、IN_AND_OUT三个值，分别表示采集输入数据、采集输出数据、同时采集输入和输出数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type TreeJobSets struct {
	// 作业Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 作业名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 作业类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobType *int64 `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 作业占用资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningCu *float64 `json:"RunningCu,omitnil,omitempty" name:"RunningCu"`

	// 作业状态 启动或者停止或者暂停
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type TreeResourceItem struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资源类型
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 目录ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 分状态统计关联作业数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefJobStatusCountSet []*RefJobStatusCountItem `json:"RefJobStatusCountSet,omitnil,omitempty" name:"RefJobStatusCountSet"`
}

// Predefined struct for user
type TriggerJobSavepointRequestParams struct {
	// 作业 SerialId
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

type TriggerJobSavepointRequest struct {
	*tchttp.BaseRequest
	
	// 作业 SerialId
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`
}

func (r *TriggerJobSavepointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerJobSavepointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "Description")
	delete(f, "WorkSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TriggerJobSavepointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TriggerJobSavepointResponseParams struct {
	// 是否成功
	SavepointTrigger *bool `json:"SavepointTrigger,omitnil,omitempty" name:"SavepointTrigger"`

	// 错误消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 快照路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinalSavepointPath *string `json:"FinalSavepointPath,omitnil,omitempty" name:"FinalSavepointPath"`

	// 快照 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SavepointId *string `json:"SavepointId,omitnil,omitempty" name:"SavepointId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TriggerJobSavepointResponse struct {
	*tchttp.BaseResponse
	Response *TriggerJobSavepointResponseParams `json:"Response"`
}

func (r *TriggerJobSavepointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerJobSavepointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WorkSpaceClusterItem struct {
	// 集群 ID
	ClusterGroupId *int64 `json:"ClusterGroupId,omitnil,omitempty" name:"ClusterGroupId"`

	// 集群 SerialId
	ClusterGroupSerialId *string `json:"ClusterGroupSerialId,omitnil,omitempty" name:"ClusterGroupSerialId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 工作空间名称
	WorkSpaceName *string `json:"WorkSpaceName,omitnil,omitempty" name:"WorkSpaceName"`

	// 绑定状态  2 绑定 1  解除绑定
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目ID string类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectIdStr *string `json:"ProjectIdStr,omitnil,omitempty" name:"ProjectIdStr"`
}

type WorkSpaceSetItem struct {
	// 工作空间 SerialId
	SerialId *string `json:"SerialId,omitnil,omitempty" name:"SerialId"`

	// 用户 APPID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 主账号 UIN
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 创建者 UIN
	CreatorUin *string `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 工作空间名称
	WorkSpaceName *string `json:"WorkSpaceName,omitnil,omitempty" name:"WorkSpaceName"`

	// 区域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 1 未初始化 2 可用  -1 已删除
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 工作空间描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 工作空间包含集群信息
	ClusterGroupSetItem []*ClusterGroupSetItem `json:"ClusterGroupSetItem,omitnil,omitempty" name:"ClusterGroupSetItem"`

	// 工作空间角色的信息
	RoleAuth []*RoleAuth `json:"RoleAuth,omitnil,omitempty" name:"RoleAuth"`

	// 工作空间成员数量
	RoleAuthCount *int64 `json:"RoleAuthCount,omitnil,omitempty" name:"RoleAuthCount"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitnil,omitempty" name:"WorkSpaceId"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobsCount *int64 `json:"JobsCount,omitnil,omitempty" name:"JobsCount"`
}