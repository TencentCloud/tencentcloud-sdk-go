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

package v20210914

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ApplicationBasicConfig struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 工作负载类型
	WorkflowKind *string `json:"WorkflowKind,omitnil,omitempty" name:"WorkflowKind"`

	// 标签信息
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// Grid唯一Key
	GridUniqKey *string `json:"GridUniqKey,omitnil,omitempty" name:"GridUniqKey"`

	// NodeSelector标签
	NodeSelector []*Label `json:"NodeSelector,omitnil,omitempty" name:"NodeSelector"`

	// 实例数
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 可用实例数
	AvailableReplicas *int64 `json:"AvailableReplicas,omitnil,omitempty" name:"AvailableReplicas"`

	// 是否开启service环境变量注入pod
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableServiceLinks *bool `json:"EnableServiceLinks,omitnil,omitempty" name:"EnableServiceLinks"`
}

type ApplicationBasicInfo struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 管理URL地址
	ManageUrl *string `json:"ManageUrl,omitnil,omitempty" name:"ManageUrl"`

	// 描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否允许可视化修改
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowVisualModify *bool `json:"AllowVisualModify,omitnil,omitempty" name:"AllowVisualModify"`
}

type ApplicationDeployMode struct {
	// 1:指定节点部署 2:单元部署
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 资源ID
	ResourceID *uint64 `json:"ResourceID,omitnil,omitempty" name:"ResourceID"`

	// 资源名
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`
}

type ApplicationStatusInfo struct {
	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 应用状态(1:待部署 2:部署中 3:运行中 4:待更新 5:更新中 6:待删除 7:删除中 8:已删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 管理地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManageUrl *string `json:"ManageUrl,omitnil,omitempty" name:"ManageUrl"`

	// 负载类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkloadKind *string `json:"WorkloadKind,omitnil,omitempty" name:"WorkloadKind"`

	// 应用部署模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployMode *ApplicationDeployMode `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 期望Pod数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 运行Pod数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableReplicas *int64 `json:"AvailableReplicas,omitnil,omitempty" name:"AvailableReplicas"`
}

// Predefined struct for user
type BuildMessageRouteRequestParams struct {
	// 路由名字
	RouteName *string `json:"RouteName,omitnil,omitempty" name:"RouteName"`

	// 源产品id
	SourceProductID *string `json:"SourceProductID,omitnil,omitempty" name:"SourceProductID"`

	// 源设备名列表
	SourceDeviceNameList []*string `json:"SourceDeviceNameList,omitnil,omitempty" name:"SourceDeviceNameList"`

	// 第一个字符为 "0"或"1"，"1"表示自定义topic
	TopicFilter *string `json:"TopicFilter,omitnil,omitempty" name:"TopicFilter"`

	// http或mqtt-broker
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 源单元id列表
	SourceUnitIDList []*string `json:"SourceUnitIDList,omitnil,omitempty" name:"SourceUnitIDList"`

	// 描述
	Descript *string `json:"Descript,omitnil,omitempty" name:"Descript"`

	// 无
	TargetOptions *string `json:"TargetOptions,omitnil,omitempty" name:"TargetOptions"`
}

type BuildMessageRouteRequest struct {
	*tchttp.BaseRequest
	
	// 路由名字
	RouteName *string `json:"RouteName,omitnil,omitempty" name:"RouteName"`

	// 源产品id
	SourceProductID *string `json:"SourceProductID,omitnil,omitempty" name:"SourceProductID"`

	// 源设备名列表
	SourceDeviceNameList []*string `json:"SourceDeviceNameList,omitnil,omitempty" name:"SourceDeviceNameList"`

	// 第一个字符为 "0"或"1"，"1"表示自定义topic
	TopicFilter *string `json:"TopicFilter,omitnil,omitempty" name:"TopicFilter"`

	// http或mqtt-broker
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 源单元id列表
	SourceUnitIDList []*string `json:"SourceUnitIDList,omitnil,omitempty" name:"SourceUnitIDList"`

	// 描述
	Descript *string `json:"Descript,omitnil,omitempty" name:"Descript"`

	// 无
	TargetOptions *string `json:"TargetOptions,omitnil,omitempty" name:"TargetOptions"`
}

func (r *BuildMessageRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BuildMessageRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteName")
	delete(f, "SourceProductID")
	delete(f, "SourceDeviceNameList")
	delete(f, "TopicFilter")
	delete(f, "Mode")
	delete(f, "SourceUnitIDList")
	delete(f, "Descript")
	delete(f, "TargetOptions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BuildMessageRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BuildMessageRouteResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BuildMessageRouteResponse struct {
	*tchttp.BaseResponse
	Response *BuildMessageRouteResponseParams `json:"Response"`
}

func (r *BuildMessageRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BuildMessageRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigMapBasicInfo struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type Container struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 镜像名
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 镜像版本
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`

	// 镜像拉取策略(Always|Never|IfNotPresent)
	ImagePullPolicy *string `json:"ImagePullPolicy,omitnil,omitempty" name:"ImagePullPolicy"`

	// 卷挂载配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeMounts []*VolumeMount `json:"VolumeMounts,omitnil,omitempty" name:"VolumeMounts"`

	// cpu最低配置
	CpuRequest *string `json:"CpuRequest,omitnil,omitempty" name:"CpuRequest"`

	// cpu最高限制
	CpuLimit *string `json:"CpuLimit,omitnil,omitempty" name:"CpuLimit"`

	// 内存最低要求
	MemoryRequest *string `json:"MemoryRequest,omitnil,omitempty" name:"MemoryRequest"`

	// 内存最高要求
	MemoryLimit *string `json:"MemoryLimit,omitnil,omitempty" name:"MemoryLimit"`

	// 内存单位
	MemoryUnit *string `json:"MemoryUnit,omitnil,omitempty" name:"MemoryUnit"`

	// gpu最高限制
	GpuLimit *string `json:"GpuLimit,omitnil,omitempty" name:"GpuLimit"`

	// 资源配置
	ResourceMapCloud []*KeyValueObj `json:"ResourceMapCloud,omitnil,omitempty" name:"ResourceMapCloud"`

	// 环境配置
	Envs []*Env `json:"Envs,omitnil,omitempty" name:"Envs"`

	// 工作目录
	WorkingDir *string `json:"WorkingDir,omitnil,omitempty" name:"WorkingDir"`

	// 命令
	Commands []*string `json:"Commands,omitnil,omitempty" name:"Commands"`

	// 参数
	Args []*string `json:"Args,omitnil,omitempty" name:"Args"`

	// 安全配置
	SecurityContext *SecurityContext `json:"SecurityContext,omitnil,omitempty" name:"SecurityContext"`

	// 就绪探针配置
	ReadinessProbe *Probe `json:"ReadinessProbe,omitnil,omitempty" name:"ReadinessProbe"`
}

// Predefined struct for user
type CreateConfigMapRequestParams struct {
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// ConfigMap名称
	ConfigMapName *string `json:"ConfigMapName,omitnil,omitempty" name:"ConfigMapName"`

	// ConfigMap内容
	ConfigMapData []*KeyValueObj `json:"ConfigMapData,omitnil,omitempty" name:"ConfigMapData"`

	// ConfigMap命名空间,默认：default
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil,omitempty" name:"ConfigMapNamespace"`
}

type CreateConfigMapRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// ConfigMap名称
	ConfigMapName *string `json:"ConfigMapName,omitnil,omitempty" name:"ConfigMapName"`

	// ConfigMap内容
	ConfigMapData []*KeyValueObj `json:"ConfigMapData,omitnil,omitempty" name:"ConfigMapData"`

	// ConfigMap命名空间,默认：default
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil,omitempty" name:"ConfigMapNamespace"`
}

func (r *CreateConfigMapRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigMapRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitID")
	delete(f, "ConfigMapName")
	delete(f, "ConfigMapData")
	delete(f, "ConfigMapNamespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConfigMapRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigMapResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConfigMapResponse struct {
	*tchttp.BaseResponse
	Response *CreateConfigMapResponseParams `json:"Response"`
}

func (r *CreateConfigMapResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigMapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeNodeBatchRequestParams struct {
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 节点信息
	Nodes []*DracoNodeInfo `json:"Nodes,omitnil,omitempty" name:"Nodes"`
}

type CreateEdgeNodeBatchRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 节点信息
	Nodes []*DracoNodeInfo `json:"Nodes,omitnil,omitempty" name:"Nodes"`
}

func (r *CreateEdgeNodeBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeNodeBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "Nodes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEdgeNodeBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeNodeBatchResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEdgeNodeBatchResponse struct {
	*tchttp.BaseResponse
	Response *CreateEdgeNodeBatchResponseParams `json:"Response"`
}

func (r *CreateEdgeNodeBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeNodeBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeNodeGroupRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// NodeGroup名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间，不填默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 模版ID数组
	NodeUnitTemplateIDs []*uint64 `json:"NodeUnitTemplateIDs,omitnil,omitempty" name:"NodeUnitTemplateIDs"`
}

type CreateEdgeNodeGroupRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// NodeGroup名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间，不填默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 模版ID数组
	NodeUnitTemplateIDs []*uint64 `json:"NodeUnitTemplateIDs,omitnil,omitempty" name:"NodeUnitTemplateIDs"`
}

func (r *CreateEdgeNodeGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeNodeGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "Name")
	delete(f, "Namespace")
	delete(f, "Description")
	delete(f, "NodeUnitTemplateIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEdgeNodeGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeNodeGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEdgeNodeGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateEdgeNodeGroupResponseParams `json:"Response"`
}

func (r *CreateEdgeNodeGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeNodeGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeNodeRequestParams struct {
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 节点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type CreateEdgeNodeRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 节点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *CreateEdgeNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEdgeNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeNodeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEdgeNodeResponse struct {
	*tchttp.BaseResponse
	Response *CreateEdgeNodeResponseParams `json:"Response"`
}

func (r *CreateEdgeNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeNodeUnitTemplateRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// NodeUnit模板名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 包含的节点列表
	Nodes []*string `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateEdgeNodeUnitTemplateRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// NodeUnit模板名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 包含的节点列表
	Nodes []*string `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateEdgeNodeUnitTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeNodeUnitTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "Name")
	delete(f, "Namespace")
	delete(f, "Nodes")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEdgeNodeUnitTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeNodeUnitTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEdgeNodeUnitTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateEdgeNodeUnitTemplateResponseParams `json:"Response"`
}

func (r *CreateEdgeNodeUnitTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeNodeUnitTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeUnitCloudRequestParams struct {
	// 集群名称，长度小于32
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// k8s版本，仅支持1.16.7 和 1.18.2
	K8sVersion *string `json:"K8sVersion,omitnil,omitempty" name:"K8sVersion"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 集群描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 集群pod cidr， 默认  10.1.0.0/16
	PodCIDR *string `json:"PodCIDR,omitnil,omitempty" name:"PodCIDR"`

	// 集群service cidr, 默认 10.2.0.0/16
	ServiceCIDR *string `json:"ServiceCIDR,omitnil,omitempty" name:"ServiceCIDR"`

	// 是否开启监控。目前内存中权限开启联系产品开通白名单
	OpenCloudMonitor *bool `json:"OpenCloudMonitor,omitnil,omitempty" name:"OpenCloudMonitor"`
}

type CreateEdgeUnitCloudRequest struct {
	*tchttp.BaseRequest
	
	// 集群名称，长度小于32
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// k8s版本，仅支持1.16.7 和 1.18.2
	K8sVersion *string `json:"K8sVersion,omitnil,omitempty" name:"K8sVersion"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 集群描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 集群pod cidr， 默认  10.1.0.0/16
	PodCIDR *string `json:"PodCIDR,omitnil,omitempty" name:"PodCIDR"`

	// 集群service cidr, 默认 10.2.0.0/16
	ServiceCIDR *string `json:"ServiceCIDR,omitnil,omitempty" name:"ServiceCIDR"`

	// 是否开启监控。目前内存中权限开启联系产品开通白名单
	OpenCloudMonitor *bool `json:"OpenCloudMonitor,omitnil,omitempty" name:"OpenCloudMonitor"`
}

func (r *CreateEdgeUnitCloudRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeUnitCloudRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "K8sVersion")
	delete(f, "VpcId")
	delete(f, "Description")
	delete(f, "PodCIDR")
	delete(f, "ServiceCIDR")
	delete(f, "OpenCloudMonitor")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEdgeUnitCloudRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeUnitCloudResponseParams struct {
	// tke集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// IECP集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEdgeUnitCloudResponse struct {
	*tchttp.BaseResponse
	Response *CreateEdgeUnitCloudResponseParams `json:"Response"`
}

func (r *CreateEdgeUnitCloudResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeUnitCloudResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeUnitDevicesRequestParams struct {
	// 无
	EdgeUnitId *int64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 无
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 无
	DeviceNames []*string `json:"DeviceNames,omitnil,omitempty" name:"DeviceNames"`
}

type CreateEdgeUnitDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 无
	EdgeUnitId *int64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 无
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 无
	DeviceNames []*string `json:"DeviceNames,omitnil,omitempty" name:"DeviceNames"`
}

func (r *CreateEdgeUnitDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeUnitDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "ProductId")
	delete(f, "DeviceNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEdgeUnitDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeUnitDevicesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEdgeUnitDevicesResponse struct {
	*tchttp.BaseResponse
	Response *CreateEdgeUnitDevicesResponseParams `json:"Response"`
}

func (r *CreateEdgeUnitDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeUnitDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIotDeviceRequestParams struct {
	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备所属的产品id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 无
	UnitID *int64 `json:"UnitID,omitnil,omitempty" name:"UnitID"`
}

type CreateIotDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备所属的产品id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 无
	UnitID *int64 `json:"UnitID,omitnil,omitempty" name:"UnitID"`
}

func (r *CreateIotDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIotDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceName")
	delete(f, "ProductId")
	delete(f, "Description")
	delete(f, "UnitID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIotDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIotDeviceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIotDeviceResponse struct {
	*tchttp.BaseResponse
	Response *CreateIotDeviceResponseParams `json:"Response"`
}

func (r *CreateIotDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIotDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMessageRouteRequestParams struct {
	// 路由名称
	RouteName *string `json:"RouteName,omitnil,omitempty" name:"RouteName"`

	// 路由备注
	Descript *string `json:"Descript,omitnil,omitempty" name:"Descript"`
}

type CreateMessageRouteRequest struct {
	*tchttp.BaseRequest
	
	// 路由名称
	RouteName *string `json:"RouteName,omitnil,omitempty" name:"RouteName"`

	// 路由备注
	Descript *string `json:"Descript,omitnil,omitempty" name:"Descript"`
}

func (r *CreateMessageRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMessageRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteName")
	delete(f, "Descript")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMessageRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMessageRouteResponseParams struct {
	// 路由id
	RouteID *int64 `json:"RouteID,omitnil,omitempty" name:"RouteID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMessageRouteResponse struct {
	*tchttp.BaseResponse
	Response *CreateMessageRouteResponseParams `json:"Response"`
}

func (r *CreateMessageRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMessageRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNamespaceRequestParams struct {
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitID")
	delete(f, "Namespace")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNamespaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateNamespaceResponseParams `json:"Response"`
}

func (r *CreateNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecretRequestParams struct {
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// secret名
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// 命名空间（默认:default）
	SecretNamespace *string `json:"SecretNamespace,omitnil,omitempty" name:"SecretNamespace"`

	// secret类型(取值范围:DockerConfigJson,Opaque 默认Opaque)
	SecretType *string `json:"SecretType,omitnil,omitempty" name:"SecretType"`

	// DockerConfig的序列化base64编码后的字符串
	DockerConfigJson *string `json:"DockerConfigJson,omitnil,omitempty" name:"DockerConfigJson"`

	// Opaque类型的Secret内容
	CloudData []*KeyValueObj `json:"CloudData,omitnil,omitempty" name:"CloudData"`

	// DockerConfig配置
	DockerConfig *DockerConfig `json:"DockerConfig,omitnil,omitempty" name:"DockerConfig"`
}

type CreateSecretRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// secret名
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// 命名空间（默认:default）
	SecretNamespace *string `json:"SecretNamespace,omitnil,omitempty" name:"SecretNamespace"`

	// secret类型(取值范围:DockerConfigJson,Opaque 默认Opaque)
	SecretType *string `json:"SecretType,omitnil,omitempty" name:"SecretType"`

	// DockerConfig的序列化base64编码后的字符串
	DockerConfigJson *string `json:"DockerConfigJson,omitnil,omitempty" name:"DockerConfigJson"`

	// Opaque类型的Secret内容
	CloudData []*KeyValueObj `json:"CloudData,omitnil,omitempty" name:"CloudData"`

	// DockerConfig配置
	DockerConfig *DockerConfig `json:"DockerConfig,omitnil,omitempty" name:"DockerConfig"`
}

func (r *CreateSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitID")
	delete(f, "SecretName")
	delete(f, "SecretNamespace")
	delete(f, "SecretType")
	delete(f, "DockerConfigJson")
	delete(f, "CloudData")
	delete(f, "DockerConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecretResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSecretResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecretResponseParams `json:"Response"`
}

func (r *CreateSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUpdateNodeUnitRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// NodeUnit所属的NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitnil,omitempty" name:"NodeGroupName"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// NodeUnit名称，通过模版创建可不填
	NodeUnitName *string `json:"NodeUnitName,omitnil,omitempty" name:"NodeUnitName"`

	// NodeUnit包含的节点列表，通过模版创建可不填
	Nodes []*string `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// NodeUnit模版ID列表
	NodeUnitTemplateIDs []*uint64 `json:"NodeUnitTemplateIDs,omitnil,omitempty" name:"NodeUnitTemplateIDs"`
}

type CreateUpdateNodeUnitRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// NodeUnit所属的NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitnil,omitempty" name:"NodeGroupName"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// NodeUnit名称，通过模版创建可不填
	NodeUnitName *string `json:"NodeUnitName,omitnil,omitempty" name:"NodeUnitName"`

	// NodeUnit包含的节点列表，通过模版创建可不填
	Nodes []*string `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// NodeUnit模版ID列表
	NodeUnitTemplateIDs []*uint64 `json:"NodeUnitTemplateIDs,omitnil,omitempty" name:"NodeUnitTemplateIDs"`
}

func (r *CreateUpdateNodeUnitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUpdateNodeUnitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "NodeGroupName")
	delete(f, "Namespace")
	delete(f, "NodeUnitName")
	delete(f, "Nodes")
	delete(f, "NodeUnitTemplateIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUpdateNodeUnitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUpdateNodeUnitResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUpdateNodeUnitResponse struct {
	*tchttp.BaseResponse
	Response *CreateUpdateNodeUnitResponseParams `json:"Response"`
}

func (r *CreateUpdateNodeUnitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUpdateNodeUnitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserTokenRequestParams struct {
	// token过期时间，有效值是1~300秒
	Second *int64 `json:"Second,omitnil,omitempty" name:"Second"`
}

type CreateUserTokenRequest struct {
	*tchttp.BaseRequest
	
	// token过期时间，有效值是1~300秒
	Second *int64 `json:"Second,omitnil,omitempty" name:"Second"`
}

func (r *CreateUserTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Second")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserTokenResponseParams struct {
	// 无
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserTokenResponseParams `json:"Response"`
}

func (r *CreateUserTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CronJob struct {
	// 调度配置
	Schedule *string `json:"Schedule,omitnil,omitempty" name:"Schedule"`

	// 运行时间
	StartingDeadlineSeconds *int64 `json:"StartingDeadlineSeconds,omitnil,omitempty" name:"StartingDeadlineSeconds"`

	// job并行策略(Allow|Forbid|Replace)
	ConcurrencyPolicy *string `json:"ConcurrencyPolicy,omitnil,omitempty" name:"ConcurrencyPolicy"`

	// Job配置
	Job *Job `json:"Job,omitnil,omitempty" name:"Job"`
}

// Predefined struct for user
type DeleteConfigMapRequestParams struct {
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// ConfigMap名
	ConfigMapName *string `json:"ConfigMapName,omitnil,omitempty" name:"ConfigMapName"`

	// ConfigMap命名空间，默认：default
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil,omitempty" name:"ConfigMapNamespace"`
}

type DeleteConfigMapRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// ConfigMap名
	ConfigMapName *string `json:"ConfigMapName,omitnil,omitempty" name:"ConfigMapName"`

	// ConfigMap命名空间，默认：default
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil,omitempty" name:"ConfigMapNamespace"`
}

func (r *DeleteConfigMapRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigMapRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitID")
	delete(f, "ConfigMapName")
	delete(f, "ConfigMapNamespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConfigMapRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigMapResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConfigMapResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConfigMapResponseParams `json:"Response"`
}

func (r *DeleteConfigMapResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigMapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeNodeGroupRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// NodeGroup名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type DeleteEdgeNodeGroupRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// NodeGroup名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *DeleteEdgeNodeGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeNodeGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "Name")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEdgeNodeGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeNodeGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEdgeNodeGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEdgeNodeGroupResponseParams `json:"Response"`
}

func (r *DeleteEdgeNodeGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeNodeGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeNodeUnitTemplatesRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 删除的NodeUnit模板ID列表
	NodeUnitTemplateIDs []*uint64 `json:"NodeUnitTemplateIDs,omitnil,omitempty" name:"NodeUnitTemplateIDs"`
}

type DeleteEdgeNodeUnitTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 删除的NodeUnit模板ID列表
	NodeUnitTemplateIDs []*uint64 `json:"NodeUnitTemplateIDs,omitnil,omitempty" name:"NodeUnitTemplateIDs"`
}

func (r *DeleteEdgeNodeUnitTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeNodeUnitTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "NodeUnitTemplateIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEdgeNodeUnitTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeNodeUnitTemplatesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEdgeNodeUnitTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEdgeNodeUnitTemplatesResponseParams `json:"Response"`
}

func (r *DeleteEdgeNodeUnitTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeNodeUnitTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeNodesRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// IECP边缘节点ID列表
	NodeIds []*uint64 `json:"NodeIds,omitnil,omitempty" name:"NodeIds"`
}

type DeleteEdgeNodesRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// IECP边缘节点ID列表
	NodeIds []*uint64 `json:"NodeIds,omitnil,omitempty" name:"NodeIds"`
}

func (r *DeleteEdgeNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "NodeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEdgeNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeNodesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEdgeNodesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEdgeNodesResponseParams `json:"Response"`
}

func (r *DeleteEdgeNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeUnitCloudRequestParams struct {
	// 边缘集群ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`
}

type DeleteEdgeUnitCloudRequest struct {
	*tchttp.BaseRequest
	
	// 边缘集群ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`
}

func (r *DeleteEdgeUnitCloudRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeUnitCloudRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEdgeUnitCloudRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeUnitCloudResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEdgeUnitCloudResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEdgeUnitCloudResponseParams `json:"Response"`
}

func (r *DeleteEdgeUnitCloudResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeUnitCloudResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeUnitDeployGridItemRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil,omitempty" name:"WorkloadKind"`

	// Grid部署名称
	GridItemName *string `json:"GridItemName,omitnil,omitempty" name:"GridItemName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type DeleteEdgeUnitDeployGridItemRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil,omitempty" name:"WorkloadKind"`

	// Grid部署名称
	GridItemName *string `json:"GridItemName,omitnil,omitempty" name:"GridItemName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *DeleteEdgeUnitDeployGridItemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeUnitDeployGridItemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "WorkloadKind")
	delete(f, "GridItemName")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEdgeUnitDeployGridItemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeUnitDeployGridItemResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEdgeUnitDeployGridItemResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEdgeUnitDeployGridItemResponseParams `json:"Response"`
}

func (r *DeleteEdgeUnitDeployGridItemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeUnitDeployGridItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEdgeUnitDevicesDevice struct {
	// 无
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 无
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

// Predefined struct for user
type DeleteEdgeUnitDevicesRequestParams struct {
	// 无
	EdgeUnitId *int64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 无
	Devices []*DeleteEdgeUnitDevicesDevice `json:"Devices,omitnil,omitempty" name:"Devices"`
}

type DeleteEdgeUnitDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 无
	EdgeUnitId *int64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 无
	Devices []*DeleteEdgeUnitDevicesDevice `json:"Devices,omitnil,omitempty" name:"Devices"`
}

func (r *DeleteEdgeUnitDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeUnitDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "Devices")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEdgeUnitDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeUnitDevicesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEdgeUnitDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEdgeUnitDevicesResponseParams `json:"Response"`
}

func (r *DeleteEdgeUnitDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeUnitDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeUnitPodRequestParams struct {
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// Pod名称
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type DeleteEdgeUnitPodRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// Pod名称
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *DeleteEdgeUnitPodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeUnitPodRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "PodName")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEdgeUnitPodRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeUnitPodResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEdgeUnitPodResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEdgeUnitPodResponseParams `json:"Response"`
}

func (r *DeleteEdgeUnitPodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeUnitPodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIotDeviceBatchRequestParams struct {
	// 无
	DeviceIDList []*uint64 `json:"DeviceIDList,omitnil,omitempty" name:"DeviceIDList"`
}

type DeleteIotDeviceBatchRequest struct {
	*tchttp.BaseRequest
	
	// 无
	DeviceIDList []*uint64 `json:"DeviceIDList,omitnil,omitempty" name:"DeviceIDList"`
}

func (r *DeleteIotDeviceBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIotDeviceBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceIDList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIotDeviceBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIotDeviceBatchResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIotDeviceBatchResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIotDeviceBatchResponseParams `json:"Response"`
}

func (r *DeleteIotDeviceBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIotDeviceBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIotDeviceRequestParams struct {
	// 设备id
	DeviceId *int64 `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type DeleteIotDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备id
	DeviceId *int64 `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

func (r *DeleteIotDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIotDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIotDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIotDeviceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIotDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIotDeviceResponseParams `json:"Response"`
}

func (r *DeleteIotDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIotDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMessageRouteRequestParams struct {
	// 无
	RouteID *int64 `json:"RouteID,omitnil,omitempty" name:"RouteID"`
}

type DeleteMessageRouteRequest struct {
	*tchttp.BaseRequest
	
	// 无
	RouteID *int64 `json:"RouteID,omitnil,omitempty" name:"RouteID"`
}

func (r *DeleteMessageRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMessageRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMessageRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMessageRouteResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMessageRouteResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMessageRouteResponseParams `json:"Response"`
}

func (r *DeleteMessageRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMessageRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNamespaceRequestParams struct {
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type DeleteNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *DeleteNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitID")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNamespaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNamespaceResponseParams `json:"Response"`
}

func (r *DeleteNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNodeUnitRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// NodeUnit所属的NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitnil,omitempty" name:"NodeGroupName"`

	// NodeUnit名称
	NodeUnitName *string `json:"NodeUnitName,omitnil,omitempty" name:"NodeUnitName"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// NodeUnit包含的节点列表
	Nodes []*string `json:"Nodes,omitnil,omitempty" name:"Nodes"`
}

type DeleteNodeUnitRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// NodeUnit所属的NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitnil,omitempty" name:"NodeGroupName"`

	// NodeUnit名称
	NodeUnitName *string `json:"NodeUnitName,omitnil,omitempty" name:"NodeUnitName"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// NodeUnit包含的节点列表
	Nodes []*string `json:"Nodes,omitnil,omitempty" name:"Nodes"`
}

func (r *DeleteNodeUnitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNodeUnitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "NodeGroupName")
	delete(f, "NodeUnitName")
	delete(f, "Namespace")
	delete(f, "Nodes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNodeUnitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNodeUnitResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNodeUnitResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNodeUnitResponseParams `json:"Response"`
}

func (r *DeleteNodeUnitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNodeUnitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecretRequestParams struct {
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// secret名称
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// secret命名空间（默认:default）
	SecretNamespace *string `json:"SecretNamespace,omitnil,omitempty" name:"SecretNamespace"`
}

type DeleteSecretRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// secret名称
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// secret命名空间（默认:default）
	SecretNamespace *string `json:"SecretNamespace,omitnil,omitempty" name:"SecretNamespace"`
}

func (r *DeleteSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitID")
	delete(f, "SecretName")
	delete(f, "SecretNamespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecretResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSecretResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecretResponseParams `json:"Response"`
}

func (r *DeleteSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigMapRequestParams struct {
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// ConfigMap名称
	ConfigMapName *string `json:"ConfigMapName,omitnil,omitempty" name:"ConfigMapName"`

	// ConfigMap命名空间
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil,omitempty" name:"ConfigMapNamespace"`
}

type DescribeConfigMapRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// ConfigMap名称
	ConfigMapName *string `json:"ConfigMapName,omitnil,omitempty" name:"ConfigMapName"`

	// ConfigMap命名空间
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil,omitempty" name:"ConfigMapNamespace"`
}

func (r *DescribeConfigMapRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigMapRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitID")
	delete(f, "ConfigMapName")
	delete(f, "ConfigMapNamespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigMapRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigMapResponseParams struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// yaml配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`

	// 配置项的json格式(base64编码)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Json *string `json:"Json,omitnil,omitempty" name:"Json"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigMapResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigMapResponseParams `json:"Response"`
}

func (r *DescribeConfigMapResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigMapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigMapYamlErrorRequestParams struct {
	// yaml文件
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`
}

type DescribeConfigMapYamlErrorRequest struct {
	*tchttp.BaseRequest
	
	// yaml文件
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`
}

func (r *DescribeConfigMapYamlErrorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigMapYamlErrorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Yaml")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigMapYamlErrorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigMapYamlErrorResponseParams struct {
	// 校验是通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckPass *bool `json:"CheckPass,omitnil,omitempty" name:"CheckPass"`

	// 错误类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrType *uint64 `json:"ErrType,omitnil,omitempty" name:"ErrType"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrInfo *string `json:"ErrInfo,omitnil,omitempty" name:"ErrInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigMapYamlErrorResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigMapYamlErrorResponseParams `json:"Response"`
}

func (r *DescribeConfigMapYamlErrorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigMapYamlErrorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigMapsRequestParams struct {
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// 翻页偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页大小(最大100)
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 命名空间
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil,omitempty" name:"ConfigMapNamespace"`

	// 模糊匹配的名称
	NamePattern *string `json:"NamePattern,omitnil,omitempty" name:"NamePattern"`

	// Sort.Fileld填写CreateTime Sort.Order(ASC|DESC) 默认ASC
	Sort *FieldSort `json:"Sort,omitnil,omitempty" name:"Sort"`
}

type DescribeConfigMapsRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// 翻页偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页大小(最大100)
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 命名空间
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil,omitempty" name:"ConfigMapNamespace"`

	// 模糊匹配的名称
	NamePattern *string `json:"NamePattern,omitnil,omitempty" name:"NamePattern"`

	// Sort.Fileld填写CreateTime Sort.Order(ASC|DESC) 默认ASC
	Sort *FieldSort `json:"Sort,omitnil,omitempty" name:"Sort"`
}

func (r *DescribeConfigMapsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigMapsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitID")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ConfigMapNamespace")
	delete(f, "NamePattern")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigMapsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigMapsResponseParams struct {
	// ConfigMap列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*ConfigMapBasicInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigMapsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigMapsResponseParams `json:"Response"`
}

func (r *DescribeConfigMapsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigMapsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDracoEdgeNodeInstallerRequestParams struct {
	// 设备SN
	SN *string `json:"SN,omitnil,omitempty" name:"SN"`
}

type DescribeDracoEdgeNodeInstallerRequest struct {
	*tchttp.BaseRequest
	
	// 设备SN
	SN *string `json:"SN,omitnil,omitempty" name:"SN"`
}

func (r *DescribeDracoEdgeNodeInstallerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDracoEdgeNodeInstallerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SN")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDracoEdgeNodeInstallerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDracoEdgeNodeInstallerResponseParams struct {
	// 在线安装命名
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnlineInstallationCommand *string `json:"OnlineInstallationCommand,omitnil,omitempty" name:"OnlineInstallationCommand"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDracoEdgeNodeInstallerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDracoEdgeNodeInstallerResponseParams `json:"Response"`
}

func (r *DescribeDracoEdgeNodeInstallerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDracoEdgeNodeInstallerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeAgentNodeInstallerRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// IECP边缘节点ID
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

type DescribeEdgeAgentNodeInstallerRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// IECP边缘节点ID
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

func (r *DescribeEdgeAgentNodeInstallerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeAgentNodeInstallerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "NodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeAgentNodeInstallerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeAgentNodeInstallerResponseParams struct {
	// 节点在线安装信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Online *EdgeNodeInstallerOnline `json:"Online,omitnil,omitempty" name:"Online"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeAgentNodeInstallerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeAgentNodeInstallerResponseParams `json:"Response"`
}

func (r *DescribeEdgeAgentNodeInstallerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeAgentNodeInstallerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeDefaultVpcRequestParams struct {

}

type DescribeEdgeDefaultVpcRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeEdgeDefaultVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeDefaultVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeDefaultVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeDefaultVpcResponseParams struct {
	// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 网络CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcCidrBlock *string `json:"VpcCidrBlock,omitnil,omitempty" name:"VpcCidrBlock"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 子网CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetCidrBlock *string `json:"SubnetCidrBlock,omitnil,omitempty" name:"SubnetCidrBlock"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeDefaultVpcResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeDefaultVpcResponseParams `json:"Response"`
}

func (r *DescribeEdgeDefaultVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeDefaultVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeNodePodContainersRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 节点ID
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Pod名称
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type DescribeEdgeNodePodContainersRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 节点ID
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Pod名称
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *DescribeEdgeNodePodContainersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeNodePodContainersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "NodeId")
	delete(f, "PodName")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeNodePodContainersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeNodePodContainersResponseParams struct {
	// Pod容器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerSet []*EdgeNodePodContainerInfo `json:"ContainerSet,omitnil,omitempty" name:"ContainerSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeNodePodContainersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeNodePodContainersResponseParams `json:"Response"`
}

func (r *DescribeEdgeNodePodContainersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeNodePodContainersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeNodePodsRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 节点ID
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Pod名称过滤串
	PodNamePattern *string `json:"PodNamePattern,omitnil,omitempty" name:"PodNamePattern"`
}

type DescribeEdgeNodePodsRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 节点ID
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Pod名称过滤串
	PodNamePattern *string `json:"PodNamePattern,omitnil,omitempty" name:"PodNamePattern"`
}

func (r *DescribeEdgeNodePodsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeNodePodsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "NodeId")
	delete(f, "Namespace")
	delete(f, "PodNamePattern")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeNodePodsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeNodePodsResponseParams struct {
	// Pod列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodSet []*EdgeNodePodInfo `json:"PodSet,omitnil,omitempty" name:"PodSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeNodePodsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeNodePodsResponseParams `json:"Response"`
}

func (r *DescribeEdgeNodePodsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeNodePodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeNodeRemarkListRequestParams struct {
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`
}

type DescribeEdgeNodeRemarkListRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`
}

func (r *DescribeEdgeNodeRemarkListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeNodeRemarkListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeNodeRemarkListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeNodeRemarkListResponseParams struct {
	// 边缘单元内的备注列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remarks []*string `json:"Remarks,omitnil,omitempty" name:"Remarks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeNodeRemarkListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeNodeRemarkListResponseParams `json:"Response"`
}

func (r *DescribeEdgeNodeRemarkListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeNodeRemarkListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeNodeRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// IECP边缘节点ID
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

type DescribeEdgeNodeRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// IECP边缘节点ID
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

func (r *DescribeEdgeNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "NodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeNodeResponseParams struct {
	// 节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 节点类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 节点状态 （1健康｜2异常｜3离线｜4未激活）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// CPU体系结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuArchitecture *string `json:"CpuArchitecture,omitnil,omitempty" name:"CpuArchitecture"`

	// AI处理器体系结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiChipArchitecture *string `json:"AiChipArchitecture,omitnil,omitempty" name:"AiChipArchitecture"`

	// IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 节点标签列表
	Labels []*EdgeNodeLabel `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 节点资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *EdgeNodeResourceInfo `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeNodeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeNodeResponseParams `json:"Response"`
}

func (r *DescribeEdgeNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeNodesRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 边缘节点名称模糊搜索串
	NamePattern *string `json:"NamePattern,omitnil,omitempty" name:"NamePattern"`

	// 边缘节点名称列表，支持批量查询 ，优先于模糊查询
	NameMatchedList []*string `json:"NameMatchedList,omitnil,omitempty" name:"NameMatchedList"`

	// 排序信息列表
	Sort []*Sort `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页面大小Limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 节点类型
	NodeType *int64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`
}

type DescribeEdgeNodesRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 边缘节点名称模糊搜索串
	NamePattern *string `json:"NamePattern,omitnil,omitempty" name:"NamePattern"`

	// 边缘节点名称列表，支持批量查询 ，优先于模糊查询
	NameMatchedList []*string `json:"NameMatchedList,omitnil,omitempty" name:"NameMatchedList"`

	// 排序信息列表
	Sort []*Sort `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页面大小Limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 节点类型
	NodeType *int64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`
}

func (r *DescribeEdgeNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "NamePattern")
	delete(f, "NameMatchedList")
	delete(f, "Sort")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NodeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeNodesResponseParams struct {
	// 边缘节点数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeSet []*EdgeNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeNodesResponseParams `json:"Response"`
}

func (r *DescribeEdgeNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeOperationLogsRequestParams struct {
	// 开始时间
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段
	Sort []*FieldSort `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 模块
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 过滤条件
	Condition *OperationLogsCondition `json:"Condition,omitnil,omitempty" name:"Condition"`
}

type DescribeEdgeOperationLogsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段
	Sort []*FieldSort `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 模块
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 过滤条件
	Condition *OperationLogsCondition `json:"Condition,omitnil,omitempty" name:"Condition"`
}

func (r *DescribeEdgeOperationLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeOperationLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Sort")
	delete(f, "Module")
	delete(f, "Condition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeOperationLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeOperationLogsResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 操作日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationLogSet []*OperationLog `json:"OperationLogSet,omitnil,omitempty" name:"OperationLogSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeOperationLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeOperationLogsResponseParams `json:"Response"`
}

func (r *DescribeEdgeOperationLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeOperationLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgePodRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Pod名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeEdgePodRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Pod名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DescribeEdgePodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgePodRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "Namespace")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgePodRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgePodResponseParams struct {
	// Pod详情信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pod *EdgeNodePodInfo `json:"Pod,omitnil,omitempty" name:"Pod"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgePodResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgePodResponseParams `json:"Response"`
}

func (r *DescribeEdgePodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgePodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeSnNodesRequestParams struct {
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 根据节点名称模糊匹配
	NamePattern *string `json:"NamePattern,omitnil,omitempty" name:"NamePattern"`

	// 根据设备SN模糊匹配
	SNPattern *string `json:"SNPattern,omitnil,omitempty" name:"SNPattern"`

	// 根据备注批次信息模糊匹配
	RemarkPattern *string `json:"RemarkPattern,omitnil,omitempty" name:"RemarkPattern"`

	// 默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeEdgeSnNodesRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 根据节点名称模糊匹配
	NamePattern *string `json:"NamePattern,omitnil,omitempty" name:"NamePattern"`

	// 根据设备SN模糊匹配
	SNPattern *string `json:"SNPattern,omitnil,omitempty" name:"SNPattern"`

	// 根据备注批次信息模糊匹配
	RemarkPattern *string `json:"RemarkPattern,omitnil,omitempty" name:"RemarkPattern"`

	// 默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeEdgeSnNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeSnNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "NamePattern")
	delete(f, "SNPattern")
	delete(f, "RemarkPattern")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeSnNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeSnNodesResponseParams struct {
	// 满足条件的总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 节点详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeSet []*EdgeDracoNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeSnNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeSnNodesResponseParams `json:"Response"`
}

func (r *DescribeEdgeSnNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeSnNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitApplicationsRequestParams struct {
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 翻页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 名称模糊匹配
	NamePattern *string `json:"NamePattern,omitnil,omitempty" name:"NamePattern"`

	// 字段排序 (Sort.Filed为:StartTime）
	Sort []*FieldSort `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 命名空间过滤
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type DescribeEdgeUnitApplicationsRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 翻页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 名称模糊匹配
	NamePattern *string `json:"NamePattern,omitnil,omitempty" name:"NamePattern"`

	// 字段排序 (Sort.Filed为:StartTime）
	Sort []*FieldSort `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 命名空间过滤
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *DescribeEdgeUnitApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitApplicationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NamePattern")
	delete(f, "Sort")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitApplicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitApplicationsResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 应用列表
	ApplicationSet []*ApplicationStatusInfo `json:"ApplicationSet,omitnil,omitempty" name:"ApplicationSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeUnitApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitApplicationsResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitDeployGridItemRequestParams struct {
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// Grid名称
	GridName *string `json:"GridName,omitnil,omitempty" name:"GridName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil,omitempty" name:"WorkloadKind"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 排序，默认ASC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeEdgeUnitDeployGridItemRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// Grid名称
	GridName *string `json:"GridName,omitnil,omitempty" name:"GridName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil,omitempty" name:"WorkloadKind"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 排序，默认ASC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *DescribeEdgeUnitDeployGridItemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitDeployGridItemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "GridName")
	delete(f, "WorkloadKind")
	delete(f, "Namespace")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitDeployGridItemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitDeployGridItemResponseParams struct {
	// 记录总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Grid部署列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeploySet []*GridItemInfo `json:"DeploySet,omitnil,omitempty" name:"DeploySet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeUnitDeployGridItemResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitDeployGridItemResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitDeployGridItemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitDeployGridItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitDeployGridItemYamlRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil,omitempty" name:"WorkloadKind"`

	// Grid部署项名称
	GridItemName *string `json:"GridItemName,omitnil,omitempty" name:"GridItemName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type DescribeEdgeUnitDeployGridItemYamlRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil,omitempty" name:"WorkloadKind"`

	// Grid部署项名称
	GridItemName *string `json:"GridItemName,omitnil,omitempty" name:"GridItemName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *DescribeEdgeUnitDeployGridItemYamlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitDeployGridItemYamlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "WorkloadKind")
	delete(f, "GridItemName")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitDeployGridItemYamlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitDeployGridItemYamlResponseParams struct {
	// yaml，base64编码字符串
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`

	// 对应类型的副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas []*int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeUnitDeployGridItemYamlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitDeployGridItemYamlResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitDeployGridItemYamlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitDeployGridItemYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitDeployGridRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 模糊匹配
	NamePattern *string `json:"NamePattern,omitnil,omitempty" name:"NamePattern"`

	// 分页offset，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页limit，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序，默认为ASC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeEdgeUnitDeployGridRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 模糊匹配
	NamePattern *string `json:"NamePattern,omitnil,omitempty" name:"NamePattern"`

	// 分页offset，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页limit，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序，默认为ASC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *DescribeEdgeUnitDeployGridRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitDeployGridRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "Namespace")
	delete(f, "NamePattern")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitDeployGridRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitDeployGridResponseParams struct {
	// 记录总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Grid列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	GridSet []*GridInfo `json:"GridSet,omitnil,omitempty" name:"GridSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeUnitDeployGridResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitDeployGridResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitDeployGridResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitDeployGridResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitExtraRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`
}

type DescribeEdgeUnitExtraRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`
}

func (r *DescribeEdgeUnitExtraRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitExtraRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitExtraRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitExtraResponseParams struct {
	// APIServer类型
	APIServerType *string `json:"APIServerType,omitnil,omitempty" name:"APIServerType"`

	// 域名URL
	APIServerURL *string `json:"APIServerURL,omitnil,omitempty" name:"APIServerURL"`

	// 域名URL对应的端口
	APIServerURLPort *string `json:"APIServerURLPort,omitnil,omitempty" name:"APIServerURLPort"`

	// 域名URL对应的端口
	APIServerResolveIP *string `json:"APIServerResolveIP,omitnil,omitempty" name:"APIServerResolveIP"`

	// 对外可访问的IP
	APIServerExposeAddress *string `json:"APIServerExposeAddress,omitnil,omitempty" name:"APIServerExposeAddress"`

	// 是否开启监控
	IsCreatePrometheus *bool `json:"IsCreatePrometheus,omitnil,omitempty" name:"IsCreatePrometheus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeUnitExtraResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitExtraResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitExtraResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitExtraResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitGridEventsRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// Grid名称
	GridName *string `json:"GridName,omitnil,omitempty" name:"GridName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil,omitempty" name:"WorkloadKind"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// NodeUnit名称
	NodeUnit *string `json:"NodeUnit,omitnil,omitempty" name:"NodeUnit"`

	// Pod名称
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`
}

type DescribeEdgeUnitGridEventsRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// Grid名称
	GridName *string `json:"GridName,omitnil,omitempty" name:"GridName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil,omitempty" name:"WorkloadKind"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// NodeUnit名称
	NodeUnit *string `json:"NodeUnit,omitnil,omitempty" name:"NodeUnit"`

	// Pod名称
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`
}

func (r *DescribeEdgeUnitGridEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitGridEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "GridName")
	delete(f, "WorkloadKind")
	delete(f, "Namespace")
	delete(f, "NodeUnit")
	delete(f, "PodName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitGridEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitGridEventsResponseParams struct {
	// 事件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventSet []*GridEventInfo `json:"EventSet,omitnil,omitempty" name:"EventSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeUnitGridEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitGridEventsResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitGridEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitGridEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitGridPodsRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// Grid名称
	GridName *string `json:"GridName,omitnil,omitempty" name:"GridName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil,omitempty" name:"WorkloadKind"`

	// NodeUnit名
	NodeUnit *string `json:"NodeUnit,omitnil,omitempty" name:"NodeUnit"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type DescribeEdgeUnitGridPodsRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// Grid名称
	GridName *string `json:"GridName,omitnil,omitempty" name:"GridName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil,omitempty" name:"WorkloadKind"`

	// NodeUnit名
	NodeUnit *string `json:"NodeUnit,omitnil,omitempty" name:"NodeUnit"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *DescribeEdgeUnitGridPodsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitGridPodsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "GridName")
	delete(f, "WorkloadKind")
	delete(f, "NodeUnit")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitGridPodsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitGridPodsResponseParams struct {
	// Pod列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodSet []*GridPodInfo `json:"PodSet,omitnil,omitempty" name:"PodSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeUnitGridPodsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitGridPodsResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitGridPodsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitGridPodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitMonitorStatusRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`
}

type DescribeEdgeUnitMonitorStatusRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`
}

func (r *DescribeEdgeUnitMonitorStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitMonitorStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitMonitorStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitMonitorStatusResponseParams struct {
	// 监控状态描述：
	// "running" 单元监控正常运行
	// "deploying" 单元监控部署中
	// "norsc" 单元需要可用节点以部署监控
	// "abnormal" 单元监控异常
	// "none" 单元监控不可用
	MonitorStatus *string `json:"MonitorStatus,omitnil,omitempty" name:"MonitorStatus"`

	// 监控是否就绪
	IsAvailable *bool `json:"IsAvailable,omitnil,omitempty" name:"IsAvailable"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeUnitMonitorStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitMonitorStatusResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitMonitorStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitMonitorStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitNodeGroupRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 分页offset，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页limit，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模糊匹配参数，精确匹配时失效
	NameFilter *string `json:"NameFilter,omitnil,omitempty" name:"NameFilter"`

	// 精确匹配参数
	NameMatched *string `json:"NameMatched,omitnil,omitempty" name:"NameMatched"`

	// 按时间排序，ASC/DESC，默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeEdgeUnitNodeGroupRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 分页offset，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页limit，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模糊匹配参数，精确匹配时失效
	NameFilter *string `json:"NameFilter,omitnil,omitempty" name:"NameFilter"`

	// 精确匹配参数
	NameMatched *string `json:"NameMatched,omitnil,omitempty" name:"NameMatched"`

	// 按时间排序，ASC/DESC，默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *DescribeEdgeUnitNodeGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitNodeGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "Namespace")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NameFilter")
	delete(f, "NameMatched")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitNodeGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitNodeGroupResponseParams struct {
	// 记录总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// NodeGroup数组
	NodeGroupInfo []*NodeGroupInfo `json:"NodeGroupInfo,omitnil,omitempty" name:"NodeGroupInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeUnitNodeGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitNodeGroupResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitNodeGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitNodeGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitNodeUnitTemplatesRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 分页查询offset，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询limit，默认为20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模糊匹配，精确匹配时失效
	NameFilter *string `json:"NameFilter,omitnil,omitempty" name:"NameFilter"`

	// 精确匹配
	NameMatched *string `json:"NameMatched,omitnil,omitempty" name:"NameMatched"`

	// 按时间排序顺序，默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeEdgeUnitNodeUnitTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 分页查询offset，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询limit，默认为20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模糊匹配，精确匹配时失效
	NameFilter *string `json:"NameFilter,omitnil,omitempty" name:"NameFilter"`

	// 精确匹配
	NameMatched *string `json:"NameMatched,omitnil,omitempty" name:"NameMatched"`

	// 按时间排序顺序，默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *DescribeEdgeUnitNodeUnitTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitNodeUnitTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "Namespace")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NameFilter")
	delete(f, "NameMatched")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitNodeUnitTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitNodeUnitTemplatesResponseParams struct {
	// 符合查询条件的记录总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// NodeUnit模板列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeUnitTemplates []*NodeUnitTemplate `json:"NodeUnitTemplates,omitnil,omitempty" name:"NodeUnitTemplates"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeUnitNodeUnitTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitNodeUnitTemplatesResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitNodeUnitTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitNodeUnitTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitsCloudRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// limit值
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 集群名称模糊匹配
	NamePattern *string `json:"NamePattern,omitnil,omitempty" name:"NamePattern"`

	// 排序，ASC/DESC(默认)
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeEdgeUnitsCloudRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// limit值
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 集群名称模糊匹配
	NamePattern *string `json:"NamePattern,omitnil,omitempty" name:"NamePattern"`

	// 排序，ASC/DESC(默认)
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *DescribeEdgeUnitsCloudRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitsCloudRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NamePattern")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitsCloudRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitsCloudResponseParams struct {
	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	EdgeUnitSet []*EdgeCloudCluster `json:"EdgeUnitSet,omitnil,omitempty" name:"EdgeUnitSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeUnitsCloudResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitsCloudResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitsCloudResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitsCloudResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIotDeviceRequestParams struct {
	// 设备id，传0值表示此参数无效
	DeviceId *int64 `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 无
	ProductID *string `json:"ProductID,omitnil,omitempty" name:"ProductID"`

	// 无
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

type DescribeIotDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备id，传0值表示此参数无效
	DeviceId *int64 `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 无
	ProductID *string `json:"ProductID,omitnil,omitempty" name:"ProductID"`

	// 无
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

func (r *DescribeIotDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIotDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "ProductID")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIotDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIotDeviceResponseParams struct {
	// 设备id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 设备名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 版本号
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// ssl证书
	Cert *string `json:"Cert,omitnil,omitempty" name:"Cert"`

	// ssl私钥
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// psk认证密钥
	Psk *string `json:"Psk,omitnil,omitempty" name:"Psk"`

	// 设备是否打开
	Disabled *bool `json:"Disabled,omitnil,omitempty" name:"Disabled"`

	// 设备日志
	LogSetting *int64 `json:"LogSetting,omitnil,omitempty" name:"LogSetting"`

	// 设备日志级别
	LogLevel *int64 `json:"LogLevel,omitnil,omitempty" name:"LogLevel"`

	// mqtt参数
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// mqtt参数
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// mqtt参数
	ClientID *string `json:"ClientID,omitnil,omitempty" name:"ClientID"`

	// 16进制的psk格式
	PskHex *string `json:"PskHex,omitnil,omitempty" name:"PskHex"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 设备在线状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 无
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 无
	UnitID *int64 `json:"UnitID,omitnil,omitempty" name:"UnitID"`

	// 无
	UnitName *string `json:"UnitName,omitnil,omitempty" name:"UnitName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIotDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIotDeviceResponseParams `json:"Response"`
}

func (r *DescribeIotDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIotDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIotDevicesRequestParams struct {
	// 页偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 产品id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称模糊查找
	NamePattern *string `json:"NamePattern,omitnil,omitempty" name:"NamePattern"`

	// 版本列表
	Versions []*string `json:"Versions,omitnil,omitempty" name:"Versions"`

	// ASC 或 DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeIotDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 页偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 产品id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称模糊查找
	NamePattern *string `json:"NamePattern,omitnil,omitempty" name:"NamePattern"`

	// 版本列表
	Versions []*string `json:"Versions,omitnil,omitempty" name:"Versions"`

	// ASC 或 DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *DescribeIotDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIotDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProductId")
	delete(f, "NamePattern")
	delete(f, "Versions")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIotDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIotDevicesResponseParams struct {
	// 符合查找条件的总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 设备列表
	DeviceSet []*IotDevicesInfo `json:"DeviceSet,omitnil,omitempty" name:"DeviceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIotDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIotDevicesResponseParams `json:"Response"`
}

func (r *DescribeIotDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIotDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMessageRouteListRequestParams struct {
	// 无
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 无
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 无
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 无
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 无
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 无
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeMessageRouteListRequest struct {
	*tchttp.BaseRequest
	
	// 无
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 无
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 无
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 无
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 无
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 无
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *DescribeMessageRouteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageRouteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filter")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMessageRouteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMessageRouteListResponseParams struct {
	// 无
	RouteList []*RouteInfo `json:"RouteList,omitnil,omitempty" name:"RouteList"`

	// 无
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMessageRouteListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMessageRouteListResponseParams `json:"Response"`
}

func (r *DescribeMessageRouteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageRouteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonitorMetricsRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 查询维度
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 起始时间Unix秒时间戳
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 终止时间Unix秒时间戳
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 步长（分钟）
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 节点名称，查询节点监控时必填
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 命名空间，不填则默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Pod名称，查询Pod监控时必填
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// Workload名称，查询Workload监控时必填
	WorkloadName *string `json:"WorkloadName,omitnil,omitempty" name:"WorkloadName"`
}

type DescribeMonitorMetricsRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 查询维度
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 起始时间Unix秒时间戳
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 终止时间Unix秒时间戳
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 步长（分钟）
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 节点名称，查询节点监控时必填
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 命名空间，不填则默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Pod名称，查询Pod监控时必填
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// Workload名称，查询Workload监控时必填
	WorkloadName *string `json:"WorkloadName,omitnil,omitempty" name:"WorkloadName"`
}

func (r *DescribeMonitorMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonitorMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "QueryType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Interval")
	delete(f, "NodeName")
	delete(f, "Namespace")
	delete(f, "PodName")
	delete(f, "WorkloadName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMonitorMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonitorMetricsResponseParams struct {
	// 查询监控指标结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metrics []*MonitorMetricsColumn `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMonitorMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMonitorMetricsResponseParams `json:"Response"`
}

func (r *DescribeMonitorMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonitorMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNamespaceRequestParams struct {
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// 命名空间名
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type DescribeNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// 命名空间名
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *DescribeNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitID")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNamespaceResourcesRequestParams struct {
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type DescribeNamespaceResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *DescribeNamespaceResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespaceResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitID")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNamespaceResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNamespaceResourcesResponseParams struct {
	// 资源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resources []*NamespaceResource `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNamespaceResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNamespaceResourcesResponseParams `json:"Response"`
}

func (r *DescribeNamespaceResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespaceResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNamespaceResponseParams struct {
	// 命名空间名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 状态 (Active|Terminating)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否保护-不允许删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protected *bool `json:"Protected,omitnil,omitempty" name:"Protected"`

	// Yaml文件格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNamespaceResponseParams `json:"Response"`
}

func (r *DescribeNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNamespacesRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// 边缘节点名称模糊搜索串
	NamePattern *string `json:"NamePattern,omitnil,omitempty" name:"NamePattern"`
}

type DescribeNamespacesRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// 边缘节点名称模糊搜索串
	NamePattern *string `json:"NamePattern,omitnil,omitempty" name:"NamePattern"`
}

func (r *DescribeNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitID")
	delete(f, "NamePattern")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNamespacesResponseParams struct {
	// 命名空间信息列表
	Items []*NamespaceInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNamespacesResponseParams `json:"Response"`
}

func (r *DescribeNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodeUnitRequestParams struct {
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// NodeUnit所属的NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitnil,omitempty" name:"NodeGroupName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 分页查询limit，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询offset，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 模糊匹配
	NameFilter *string `json:"NameFilter,omitnil,omitempty" name:"NameFilter"`
}

type DescribeNodeUnitRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// NodeUnit所属的NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitnil,omitempty" name:"NodeGroupName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 分页查询limit，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询offset，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 模糊匹配
	NameFilter *string `json:"NameFilter,omitnil,omitempty" name:"NameFilter"`
}

func (r *DescribeNodeUnitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeUnitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "NodeGroupName")
	delete(f, "Namespace")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "NameFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNodeUnitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodeUnitResponseParams struct {
	// 符合查询条件的记录总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// NodeUnit信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeGridInfo []*NodeUnitInfo `json:"NodeGridInfo,omitnil,omitempty" name:"NodeGridInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNodeUnitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNodeUnitResponseParams `json:"Response"`
}

func (r *DescribeNodeUnitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeUnitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodeUnitTemplateOnNodeGroupRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitnil,omitempty" name:"NodeGroupName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 名称模糊匹配
	NodeUnitNamePattern *string `json:"NodeUnitNamePattern,omitnil,omitempty" name:"NodeUnitNamePattern"`

	// 分页查询offset，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询limit，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序，默认DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeNodeUnitTemplateOnNodeGroupRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitnil,omitempty" name:"NodeGroupName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 名称模糊匹配
	NodeUnitNamePattern *string `json:"NodeUnitNamePattern,omitnil,omitempty" name:"NodeUnitNamePattern"`

	// 分页查询offset，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询limit，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序，默认DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *DescribeNodeUnitTemplateOnNodeGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeUnitTemplateOnNodeGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "NodeGroupName")
	delete(f, "Namespace")
	delete(f, "NodeUnitNamePattern")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNodeUnitTemplateOnNodeGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodeUnitTemplateOnNodeGroupResponseParams struct {
	// 记录总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// NodeUnit模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeUnitTemplates []*NodeGroupNodeUnitTemplateInfo `json:"NodeUnitTemplates,omitnil,omitempty" name:"NodeUnitTemplates"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNodeUnitTemplateOnNodeGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNodeUnitTemplateOnNodeGroupResponseParams `json:"Response"`
}

func (r *DescribeNodeUnitTemplateOnNodeGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeUnitTemplateOnNodeGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecretRequestParams struct {
	// 边缘单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// secret名
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// 命名空间(默认值:default）
	SecretNamespace *string `json:"SecretNamespace,omitnil,omitempty" name:"SecretNamespace"`
}

type DescribeSecretRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// secret名
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// 命名空间(默认值:default）
	SecretNamespace *string `json:"SecretNamespace,omitnil,omitempty" name:"SecretNamespace"`
}

func (r *DescribeSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitID")
	delete(f, "SecretName")
	delete(f, "SecretNamespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecretResponseParams struct {
	// Secret名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// secret的yaml格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`

	// secret的json格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Json *string `json:"Json,omitnil,omitempty" name:"Json"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecretResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecretResponseParams `json:"Response"`
}

func (r *DescribeSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecretYamlErrorRequestParams struct {
	// yaml文件
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`
}

type DescribeSecretYamlErrorRequest struct {
	*tchttp.BaseRequest
	
	// yaml文件
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`
}

func (r *DescribeSecretYamlErrorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecretYamlErrorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Yaml")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecretYamlErrorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecretYamlErrorResponseParams struct {
	// 校验是通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckPass *bool `json:"CheckPass,omitnil,omitempty" name:"CheckPass"`

	// 错误类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrType *uint64 `json:"ErrType,omitnil,omitempty" name:"ErrType"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrInfo *string `json:"ErrInfo,omitnil,omitempty" name:"ErrInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecretYamlErrorResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecretYamlErrorResponseParams `json:"Response"`
}

func (r *DescribeSecretYamlErrorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecretYamlErrorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecretsRequestParams struct {
	// 边缘单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// 页号
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 命名空间
	SecretNamespace *string `json:"SecretNamespace,omitnil,omitempty" name:"SecretNamespace"`

	// Secret名(模糊匹配)
	NamePattern *string `json:"NamePattern,omitnil,omitempty" name:"NamePattern"`

	// Sort.Field:CreateTime Sort.Order:ASC|DESC
	Sort *FieldSort `json:"Sort,omitnil,omitempty" name:"Sort"`

	// Secret类型(DockerConfigJson或Opaque)
	SecretType *string `json:"SecretType,omitnil,omitempty" name:"SecretType"`
}

type DescribeSecretsRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// 页号
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 命名空间
	SecretNamespace *string `json:"SecretNamespace,omitnil,omitempty" name:"SecretNamespace"`

	// Secret名(模糊匹配)
	NamePattern *string `json:"NamePattern,omitnil,omitempty" name:"NamePattern"`

	// Sort.Field:CreateTime Sort.Order:ASC|DESC
	Sort *FieldSort `json:"Sort,omitnil,omitempty" name:"Sort"`

	// Secret类型(DockerConfigJson或Opaque)
	SecretType *string `json:"SecretType,omitnil,omitempty" name:"SecretType"`
}

func (r *DescribeSecretsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecretsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitID")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SecretNamespace")
	delete(f, "NamePattern")
	delete(f, "Sort")
	delete(f, "SecretType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecretsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecretsResponseParams struct {
	// 总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Secret列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*SecretItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecretsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecretsResponseParams `json:"Response"`
}

func (r *DescribeSecretsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecretsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeYeheResourceLimitRequestParams struct {

}

type DescribeYeheResourceLimitRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeYeheResourceLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeYeheResourceLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeYeheResourceLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeYeheResourceLimitResponseParams struct {
	// 用户父账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 允许创建的节点数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateNodeLimit *uint64 `json:"CreateNodeLimit,omitnil,omitempty" name:"CreateNodeLimit"`

	// 允许创建的集群数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateClusterLimit *uint64 `json:"CreateClusterLimit,omitnil,omitempty" name:"CreateClusterLimit"`

	// 是否有监控开启权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnablePermMonitor *bool `json:"EnablePermMonitor,omitnil,omitempty" name:"EnablePermMonitor"`

	// 节点是否有admin的所有权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnablePermAdminNode *bool `json:"EnablePermAdminNode,omitnil,omitempty" name:"EnablePermAdminNode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeYeheResourceLimitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeYeheResourceLimitResponseParams `json:"Response"`
}

func (r *DescribeYeheResourceLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeYeheResourceLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DockerConfig struct {
	// 镜像仓库地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryDomain *string `json:"RegistryDomain,omitnil,omitempty" name:"RegistryDomain"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type DracoNodeInfo struct {
	// 设备SN。SN仅支持大写字母、数字，长度限制为1~32个字符
	SN *string `json:"SN,omitnil,omitempty" name:"SN"`

	// 节点名称。长度限制为1~63个字符，节点名称只支持小写英文、数字、中横线、英文句号
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 节点备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type EdgeCloudCluster struct {
	// IECP侧边缘集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EdgeId *uint64 `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 边缘集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	K8SVersion *string `json:"K8SVersion,omitnil,omitempty" name:"K8SVersion"`

	// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterDesc *string `json:"ClusterDesc,omitnil,omitempty" name:"ClusterDesc"`

	// 集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// pod cidr
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodCIDR *string `json:"PodCIDR,omitnil,omitempty" name:"PodCIDR"`

	// service cidr
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceCIDR *string `json:"ServiceCIDR,omitnil,omitempty" name:"ServiceCIDR"`

	// 边缘版本类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EdgeClusterVersion *string `json:"EdgeClusterVersion,omitnil,omitempty" name:"EdgeClusterVersion"`

	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UID *string `json:"UID,omitnil,omitempty" name:"UID"`
}

type EdgeDracoNodeInfo struct {
	// 节点ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 节点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否已激活
	IsUsed *bool `json:"IsUsed,omitnil,omitempty" name:"IsUsed"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 备注信息，如批次
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// SN 设备号
	SN *string `json:"SN,omitnil,omitempty" name:"SN"`
}

type EdgeNodeInfo struct {
	// IECP边缘节点ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 节点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 节点状态 （1健康｜2异常｜3离线｜4未激活）
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 节点资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *EdgeNodeResourceInfo `json:"Resource,omitnil,omitempty" name:"Resource"`

	// CPU体系结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuArchitecture *string `json:"CpuArchitecture,omitnil,omitempty" name:"CpuArchitecture"`

	// IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 操作系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatingSystem *string `json:"OperatingSystem,omitnil,omitempty" name:"OperatingSystem"`

	// 节点所属的NodeUnit
	// key：NodeUnit模版ID，Value：NodeUnit模版名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeUnits *KeyValueObj `json:"NodeUnits,omitnil,omitempty" name:"NodeUnits"`
}

type EdgeNodeInstallerOnline struct {
	// 节点安装脚本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptName *string `json:"ScriptName,omitnil,omitempty" name:"ScriptName"`

	// 节点安装脚本下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptDownloadUrl *string `json:"ScriptDownloadUrl,omitnil,omitempty" name:"ScriptDownloadUrl"`

	// 节点安装命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	Guide *string `json:"Guide,omitnil,omitempty" name:"Guide"`
}

type EdgeNodeLabel struct {
	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 是否受保护
	Protected *bool `json:"Protected,omitnil,omitempty" name:"Protected"`
}

type EdgeNodePodContainerInfo struct {
	// Pod名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 容器ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 镜像（含版本号）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// CPU Request
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuRequest *string `json:"CpuRequest,omitnil,omitempty" name:"CpuRequest"`

	// CPU Limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuLimit *string `json:"CpuLimit,omitnil,omitempty" name:"CpuLimit"`

	// Memory Request
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryRequest *string `json:"MemoryRequest,omitnil,omitempty" name:"MemoryRequest"`

	// Memory Limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryLimit *string `json:"MemoryLimit,omitnil,omitempty" name:"MemoryLimit"`

	// 重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *uint64 `json:"RestartCount,omitnil,omitempty" name:"RestartCount"`

	// 容器状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type EdgeNodePodInfo struct {
	// Pod名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Pod状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 所在节点IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeIp *string `json:"NodeIp,omitnil,omitempty" name:"NodeIp"`

	// 实例IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// CPU Request
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuRequest *string `json:"CpuRequest,omitnil,omitempty" name:"CpuRequest"`

	// Memory Request
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryRequest *string `json:"MemoryRequest,omitnil,omitempty" name:"MemoryRequest"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 工作负载类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkloadType *string `json:"WorkloadType,omitnil,omitempty" name:"WorkloadType"`

	// 工作负载名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkloadName *string `json:"WorkloadName,omitnil,omitempty" name:"WorkloadName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *uint64 `json:"RestartCount,omitnil,omitempty" name:"RestartCount"`

	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`
}

type EdgeNodeResourceInfo struct {
	// 可使用的CPU 单位: m核
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllocatedCPU *string `json:"AllocatedCPU,omitnil,omitempty" name:"AllocatedCPU"`

	// CPU总量 单位:m核
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCPU *string `json:"TotalCPU,omitnil,omitempty" name:"TotalCPU"`

	// 已分配的内存 单位G
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllocatedMemory *string `json:"AllocatedMemory,omitnil,omitempty" name:"AllocatedMemory"`

	// 内存总量 单位G
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalMemory *string `json:"TotalMemory,omitnil,omitempty" name:"TotalMemory"`

	// 已分配的GPU资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllocatedGPU *string `json:"AllocatedGPU,omitnil,omitempty" name:"AllocatedGPU"`

	// GPU总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalGPU *string `json:"TotalGPU,omitnil,omitempty" name:"TotalGPU"`

	// 可使用的CPU 单位: m核
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableCPU *string `json:"AvailableCPU,omitnil,omitempty" name:"AvailableCPU"`

	// 可使用的内存 单位: G
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableMemory *string `json:"AvailableMemory,omitnil,omitempty" name:"AvailableMemory"`

	// 可使用的GPU资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableGPU *string `json:"AvailableGPU,omitnil,omitempty" name:"AvailableGPU"`
}

type Env struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 值引用
	ValueFrom *EnvValueSelector `json:"ValueFrom,omitnil,omitempty" name:"ValueFrom"`
}

type EnvValueSelector struct {
	// 健名
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 对象名
	ObjectName *string `json:"ObjectName,omitnil,omitempty" name:"ObjectName"`

	// 对象值
	ObjectType *string `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`
}

type FieldSort struct {
	// 字段名
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 排序(ASC:升序 DESC:降序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

// Predefined struct for user
type GetMarketComponentListRequestParams struct {
	// 页偏移，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 名称模糊筛选
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 以名称排序，ASC、DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type GetMarketComponentListRequest struct {
	*tchttp.BaseRequest
	
	// 页偏移，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 名称模糊筛选
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 以名称排序，ASC、DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *GetMarketComponentListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMarketComponentListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filter")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetMarketComponentListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetMarketComponentListResponseParams struct {
	// 组件列表
	ComponentList []*MarketComponentInfo `json:"ComponentList,omitnil,omitempty" name:"ComponentList"`

	// 组件总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetMarketComponentListResponse struct {
	*tchttp.BaseResponse
	Response *GetMarketComponentListResponseParams `json:"Response"`
}

func (r *GetMarketComponentListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMarketComponentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetMarketComponentRequestParams struct {
	// 组件ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type GetMarketComponentRequest struct {
	*tchttp.BaseRequest
	
	// 组件ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *GetMarketComponentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMarketComponentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetMarketComponentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetMarketComponentResponseParams struct {
	// 组件ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 组件名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 发行组织
	Author *string `json:"Author,omitnil,omitempty" name:"Author"`

	// 发布时间
	ReleaseTime *string `json:"ReleaseTime,omitnil,omitempty" name:"ReleaseTime"`

	// 组件简介
	Outline *string `json:"Outline,omitnil,omitempty" name:"Outline"`

	// 详细介绍链接
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 图标连接
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`

	// 组件版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 组件可视化配置
	WorkloadVisualConfig *string `json:"WorkloadVisualConfig,omitnil,omitempty" name:"WorkloadVisualConfig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetMarketComponentResponse struct {
	*tchttp.BaseResponse
	Response *GetMarketComponentResponseParams `json:"Response"`
}

func (r *GetMarketComponentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMarketComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GridDetail struct {
	// Grid名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// GridID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type GridEventInfo struct {
	// 首次出现时间
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// 最后出现时间
	LastTime *string `json:"LastTime,omitnil,omitempty" name:"LastTime"`

	// 对象类型
	InvolvedObjectKind *string `json:"InvolvedObjectKind,omitnil,omitempty" name:"InvolvedObjectKind"`

	// 对象名称
	InvolvedObjectName *string `json:"InvolvedObjectName,omitnil,omitempty" name:"InvolvedObjectName"`

	// 事件类型(Normal,Warning)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 事件原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 事件内容
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 次数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 节点名（Pod事件类型时有值）
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 节点内部IP（Pod事件类型时有值）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`
}

type GridInfo struct {
	// DeployGridId
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Key
	GridUniqKey *string `json:"GridUniqKey,omitnil,omitempty" name:"GridUniqKey"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 工作负载类型
	WorkloadKind *string `json:"WorkloadKind,omitnil,omitempty" name:"WorkloadKind"`

	// 启动时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 创建人
	Publisher *string `json:"Publisher,omitnil,omitempty" name:"Publisher"`

	// 版本信息
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`
}

type GridItemInfo struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 期望副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 可用副本数
	AvailableReplicas *int64 `json:"AvailableReplicas,omitnil,omitempty" name:"AvailableReplicas"`

	// 启动时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 工作负载类型
	WorkloadKind *string `json:"WorkloadKind,omitnil,omitempty" name:"WorkloadKind"`
}

type GridPodInfo struct {
	// Pod名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间
	NameSpace *string `json:"NameSpace,omitnil,omitempty" name:"NameSpace"`

	// 状态(Pending｜Running｜Succeeded｜Failed｜Unknown)
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 节点名
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 节点IP
	NodeIP *string `json:"NodeIP,omitnil,omitempty" name:"NodeIP"`

	// Pod的IP
	PodIP *string `json:"PodIP,omitnil,omitempty" name:"PodIP"`

	// 启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 运行时长（秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunSec *int64 `json:"RunSec,omitnil,omitempty" name:"RunSec"`

	// 重启次数
	RestartCount *int64 `json:"RestartCount,omitnil,omitempty" name:"RestartCount"`

	// 集群名称ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`
}

type HorizontalPodAutoscaler struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 最小实例数
	MinReplicas *int64 `json:"MinReplicas,omitnil,omitempty" name:"MinReplicas"`

	// 最大实例数
	MaxReplicas *int64 `json:"MaxReplicas,omitnil,omitempty" name:"MaxReplicas"`

	// 资源目标指标
	ResourceMetricTarget []*ResourceMetricTarget `json:"ResourceMetricTarget,omitnil,omitempty" name:"ResourceMetricTarget"`
}

type HttpHeader struct {
	// HTTP头的名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// HTTP头的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type HttpProbe struct {
	// 请求路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 请求端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 请求地址，默认Pod的IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 请求模式  HTTP|HTTPS，默认HTTP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scheme *string `json:"Scheme,omitnil,omitempty" name:"Scheme"`

	// HTTP的请求头
	// 注意：此字段可能返回 null，表示取不到有效值。
	Headers []*HttpHeader `json:"Headers,omitnil,omitempty" name:"Headers"`
}

type IotDevicesInfo struct {
	// 设备id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 设备名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 设备状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 设备打开状态
	Disabled *bool `json:"Disabled,omitnil,omitempty" name:"Disabled"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 设备创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后在线时间
	LastOnlineTime *string `json:"LastOnlineTime,omitnil,omitempty" name:"LastOnlineTime"`

	// 设备是否绑定到节点
	IsBound *bool `json:"IsBound,omitnil,omitempty" name:"IsBound"`

	// 设备版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 无
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 无
	UnitID *int64 `json:"UnitID,omitnil,omitempty" name:"UnitID"`

	// 无
	UnitName *string `json:"UnitName,omitnil,omitempty" name:"UnitName"`
}

type Job struct {
	// 并发数
	Parallelism *int64 `json:"Parallelism,omitnil,omitempty" name:"Parallelism"`

	// 完成数
	Completion *int64 `json:"Completion,omitnil,omitempty" name:"Completion"`

	// 最大运行时间
	ActiveDeadlineSeconds *int64 `json:"ActiveDeadlineSeconds,omitnil,omitempty" name:"ActiveDeadlineSeconds"`

	// 失败前重试次数
	BackOffLimit *int64 `json:"BackOffLimit,omitnil,omitempty" name:"BackOffLimit"`
}

type KeyValueObj struct {
	// Key值
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Label struct {
	// 健名
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 健值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type MarketComponentInfo struct {
	// 组件ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 组件名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 发布者
	Author *string `json:"Author,omitnil,omitempty" name:"Author"`

	// 发布时间
	ReleaseTime *string `json:"ReleaseTime,omitnil,omitempty" name:"ReleaseTime"`

	// 组件简介
	Outline *string `json:"Outline,omitnil,omitempty" name:"Outline"`

	// 指向详细描述的url
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 图标链接
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`

	// 组件版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 组件可视化信息
	WorkloadVisualConfig *string `json:"WorkloadVisualConfig,omitnil,omitempty" name:"WorkloadVisualConfig"`

	// 无
	DetailUrl *string `json:"DetailUrl,omitnil,omitempty" name:"DetailUrl"`

	// 无
	Installed *bool `json:"Installed,omitnil,omitempty" name:"Installed"`
}

// Predefined struct for user
type ModifyConfigMapRequestParams struct {
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// ConfigMap名称
	ConfigMapName *string `json:"ConfigMapName,omitnil,omitempty" name:"ConfigMapName"`

	// Yaml配置, base64之后的串
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`

	// ConfigMap命名空间
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil,omitempty" name:"ConfigMapNamespace"`
}

type ModifyConfigMapRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// ConfigMap名称
	ConfigMapName *string `json:"ConfigMapName,omitnil,omitempty" name:"ConfigMapName"`

	// Yaml配置, base64之后的串
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`

	// ConfigMap命名空间
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil,omitempty" name:"ConfigMapNamespace"`
}

func (r *ModifyConfigMapRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigMapRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitID")
	delete(f, "ConfigMapName")
	delete(f, "Yaml")
	delete(f, "ConfigMapNamespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConfigMapRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConfigMapResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyConfigMapResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConfigMapResponseParams `json:"Response"`
}

func (r *ModifyConfigMapResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigMapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeDracoNodeRequestParams struct {
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 边缘节点ID
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点信息
	NodeInfo *DracoNodeInfo `json:"NodeInfo,omitnil,omitempty" name:"NodeInfo"`

	// 是否重置draco设备
	IsReset *bool `json:"IsReset,omitnil,omitempty" name:"IsReset"`
}

type ModifyEdgeDracoNodeRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 边缘节点ID
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点信息
	NodeInfo *DracoNodeInfo `json:"NodeInfo,omitnil,omitempty" name:"NodeInfo"`

	// 是否重置draco设备
	IsReset *bool `json:"IsReset,omitnil,omitempty" name:"IsReset"`
}

func (r *ModifyEdgeDracoNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeDracoNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "NodeId")
	delete(f, "NodeInfo")
	delete(f, "IsReset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEdgeDracoNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeDracoNodeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEdgeDracoNodeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEdgeDracoNodeResponseParams `json:"Response"`
}

func (r *ModifyEdgeDracoNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeDracoNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeNodeLabelsRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// IECP边缘节点ID
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 标签列表
	Labels []*KeyValueObj `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type ModifyEdgeNodeLabelsRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// IECP边缘节点ID
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 标签列表
	Labels []*KeyValueObj `json:"Labels,omitnil,omitempty" name:"Labels"`
}

func (r *ModifyEdgeNodeLabelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeNodeLabelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "NodeId")
	delete(f, "Labels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEdgeNodeLabelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeNodeLabelsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEdgeNodeLabelsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEdgeNodeLabelsResponseParams `json:"Response"`
}

func (r *ModifyEdgeNodeLabelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeNodeLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeUnitApplicationBasicInfoRequestParams struct {
	// 应用基本信息
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitnil,omitempty" name:"BasicInfo"`

	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
}

type ModifyEdgeUnitApplicationBasicInfoRequest struct {
	*tchttp.BaseRequest
	
	// 应用基本信息
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitnil,omitempty" name:"BasicInfo"`

	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
}

func (r *ModifyEdgeUnitApplicationBasicInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeUnitApplicationBasicInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BasicInfo")
	delete(f, "EdgeUnitId")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEdgeUnitApplicationBasicInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeUnitApplicationBasicInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEdgeUnitApplicationBasicInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEdgeUnitApplicationBasicInfoResponseParams `json:"Response"`
}

func (r *ModifyEdgeUnitApplicationBasicInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeUnitApplicationBasicInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeUnitApplicationVisualizationRequestParams struct {
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 应用配置
	BasicConfig *ApplicationBasicConfig `json:"BasicConfig,omitnil,omitempty" name:"BasicConfig"`

	// 卷配置
	Volumes []*Volume `json:"Volumes,omitnil,omitempty" name:"Volumes"`

	// 初始容器列表
	InitContainers []*Container `json:"InitContainers,omitnil,omitempty" name:"InitContainers"`

	// 容器配置
	Containers []*Container `json:"Containers,omitnil,omitempty" name:"Containers"`

	// 服务配置
	Service *Service `json:"Service,omitnil,omitempty" name:"Service"`

	// Job配置
	Job *Job `json:"Job,omitnil,omitempty" name:"Job"`

	// CronJob配置
	CronJob *CronJob `json:"CronJob,omitnil,omitempty" name:"CronJob"`

	// 重启策略
	RestartPolicy *string `json:"RestartPolicy,omitnil,omitempty" name:"RestartPolicy"`

	// 镜像拉取密钥
	ImagePullSecrets []*string `json:"ImagePullSecrets,omitnil,omitempty" name:"ImagePullSecrets"`

	// HPA配置
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil,omitempty" name:"HorizontalPodAutoscaler"`
}

type ModifyEdgeUnitApplicationVisualizationRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 应用配置
	BasicConfig *ApplicationBasicConfig `json:"BasicConfig,omitnil,omitempty" name:"BasicConfig"`

	// 卷配置
	Volumes []*Volume `json:"Volumes,omitnil,omitempty" name:"Volumes"`

	// 初始容器列表
	InitContainers []*Container `json:"InitContainers,omitnil,omitempty" name:"InitContainers"`

	// 容器配置
	Containers []*Container `json:"Containers,omitnil,omitempty" name:"Containers"`

	// 服务配置
	Service *Service `json:"Service,omitnil,omitempty" name:"Service"`

	// Job配置
	Job *Job `json:"Job,omitnil,omitempty" name:"Job"`

	// CronJob配置
	CronJob *CronJob `json:"CronJob,omitnil,omitempty" name:"CronJob"`

	// 重启策略
	RestartPolicy *string `json:"RestartPolicy,omitnil,omitempty" name:"RestartPolicy"`

	// 镜像拉取密钥
	ImagePullSecrets []*string `json:"ImagePullSecrets,omitnil,omitempty" name:"ImagePullSecrets"`

	// HPA配置
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil,omitempty" name:"HorizontalPodAutoscaler"`
}

func (r *ModifyEdgeUnitApplicationVisualizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeUnitApplicationVisualizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "ApplicationId")
	delete(f, "BasicConfig")
	delete(f, "Volumes")
	delete(f, "InitContainers")
	delete(f, "Containers")
	delete(f, "Service")
	delete(f, "Job")
	delete(f, "CronJob")
	delete(f, "RestartPolicy")
	delete(f, "ImagePullSecrets")
	delete(f, "HorizontalPodAutoscaler")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEdgeUnitApplicationVisualizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeUnitApplicationVisualizationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEdgeUnitApplicationVisualizationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEdgeUnitApplicationVisualizationResponseParams `json:"Response"`
}

func (r *ModifyEdgeUnitApplicationVisualizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeUnitApplicationVisualizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeUnitApplicationYamlRequestParams struct {
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Yaml配置
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`
}

type ModifyEdgeUnitApplicationYamlRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Yaml配置
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`
}

func (r *ModifyEdgeUnitApplicationYamlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeUnitApplicationYamlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "ApplicationId")
	delete(f, "Yaml")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEdgeUnitApplicationYamlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeUnitApplicationYamlResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEdgeUnitApplicationYamlResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEdgeUnitApplicationYamlResponseParams `json:"Response"`
}

func (r *ModifyEdgeUnitApplicationYamlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeUnitApplicationYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeUnitCloudApiRequestParams struct {
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 边缘单元名称，64字符内
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述，200字符内
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否开启监控
	OpenCloudMonitor *bool `json:"OpenCloudMonitor,omitnil,omitempty" name:"OpenCloudMonitor"`
}

type ModifyEdgeUnitCloudApiRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 边缘单元名称，64字符内
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述，200字符内
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否开启监控
	OpenCloudMonitor *bool `json:"OpenCloudMonitor,omitnil,omitempty" name:"OpenCloudMonitor"`
}

func (r *ModifyEdgeUnitCloudApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeUnitCloudApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "OpenCloudMonitor")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEdgeUnitCloudApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeUnitCloudApiResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEdgeUnitCloudApiResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEdgeUnitCloudApiResponseParams `json:"Response"`
}

func (r *ModifyEdgeUnitCloudApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeUnitCloudApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeUnitDeployGridItemRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// Grid名称
	GridItemName *string `json:"GridItemName,omitnil,omitempty" name:"GridItemName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil,omitempty" name:"WorkloadKind"`

	// 副本数
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type ModifyEdgeUnitDeployGridItemRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// Grid名称
	GridItemName *string `json:"GridItemName,omitnil,omitempty" name:"GridItemName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil,omitempty" name:"WorkloadKind"`

	// 副本数
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *ModifyEdgeUnitDeployGridItemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeUnitDeployGridItemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "GridItemName")
	delete(f, "WorkloadKind")
	delete(f, "Replicas")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEdgeUnitDeployGridItemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeUnitDeployGridItemResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEdgeUnitDeployGridItemResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEdgeUnitDeployGridItemResponseParams `json:"Response"`
}

func (r *ModifyEdgeUnitDeployGridItemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeUnitDeployGridItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeUnitRequestParams struct {
	// 边缘集群ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 边缘集群名称，64字符以内
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 集群描述，200字符以内
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyEdgeUnitRequest struct {
	*tchttp.BaseRequest
	
	// 边缘集群ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 边缘集群名称，64字符以内
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 集群描述，200字符以内
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyEdgeUnitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeUnitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEdgeUnitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeUnitResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEdgeUnitResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEdgeUnitResponseParams `json:"Response"`
}

func (r *ModifyEdgeUnitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeUnitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIotDeviceRequestParams struct {
	// 设备id
	DeviceId *int64 `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 设备是否开启
	Disabled *bool `json:"Disabled,omitnil,omitempty" name:"Disabled"`

	// 日志设置
	LogSetting *int64 `json:"LogSetting,omitnil,omitempty" name:"LogSetting"`

	// 日志级别
	LogLevel *int64 `json:"LogLevel,omitnil,omitempty" name:"LogLevel"`
}

type ModifyIotDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备id
	DeviceId *int64 `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 设备是否开启
	Disabled *bool `json:"Disabled,omitnil,omitempty" name:"Disabled"`

	// 日志设置
	LogSetting *int64 `json:"LogSetting,omitnil,omitempty" name:"LogSetting"`

	// 日志级别
	LogLevel *int64 `json:"LogLevel,omitnil,omitempty" name:"LogLevel"`
}

func (r *ModifyIotDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIotDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "Description")
	delete(f, "Disabled")
	delete(f, "LogSetting")
	delete(f, "LogLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIotDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIotDeviceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyIotDeviceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIotDeviceResponseParams `json:"Response"`
}

func (r *ModifyIotDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIotDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNodeUnitTemplateRequestParams struct {
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// NodeUnit模板ID
	NodeUnitTemplateID *uint64 `json:"NodeUnitTemplateID,omitnil,omitempty" name:"NodeUnitTemplateID"`

	// 包含的节点列表
	Nodes []*string `json:"Nodes,omitnil,omitempty" name:"Nodes"`
}

type ModifyNodeUnitTemplateRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// NodeUnit模板ID
	NodeUnitTemplateID *uint64 `json:"NodeUnitTemplateID,omitnil,omitempty" name:"NodeUnitTemplateID"`

	// 包含的节点列表
	Nodes []*string `json:"Nodes,omitnil,omitempty" name:"Nodes"`
}

func (r *ModifyNodeUnitTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNodeUnitTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "NodeUnitTemplateID")
	delete(f, "Nodes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNodeUnitTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNodeUnitTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNodeUnitTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNodeUnitTemplateResponseParams `json:"Response"`
}

func (r *ModifyNodeUnitTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNodeUnitTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecretRequestParams struct {
	// 边缘单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// Secret名
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Secret的Yaml格式
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`

	// Secret命名空间（默认:default）
	SecretNamespace *string `json:"SecretNamespace,omitnil,omitempty" name:"SecretNamespace"`
}

type ModifySecretRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil,omitempty" name:"EdgeUnitID"`

	// Secret名
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Secret的Yaml格式
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`

	// Secret命名空间（默认:default）
	SecretNamespace *string `json:"SecretNamespace,omitnil,omitempty" name:"SecretNamespace"`
}

func (r *ModifySecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitID")
	delete(f, "SecretName")
	delete(f, "Yaml")
	delete(f, "SecretNamespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecretResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySecretResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecretResponseParams `json:"Response"`
}

func (r *ModifySecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorMetricsColumn struct {
	// 数据名称
	ColumnName *string `json:"ColumnName,omitnil,omitempty" name:"ColumnName"`

	// 数据内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnData []*string `json:"ColumnData,omitnil,omitempty" name:"ColumnData"`

	// 数据所属，查询Workload类型时有值
	ColumnBelong *string `json:"ColumnBelong,omitnil,omitempty" name:"ColumnBelong"`

	// 最大值
	MaxValue *float64 `json:"MaxValue,omitnil,omitempty" name:"MaxValue"`

	// 最小值
	MinValue *float64 `json:"MinValue,omitnil,omitempty" name:"MinValue"`

	// 平均值
	AvgValue *float64 `json:"AvgValue,omitnil,omitempty" name:"AvgValue"`

	// 时间戳数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnTime *int64 `json:"ColumnTime,omitnil,omitempty" name:"ColumnTime"`
}

type NamespaceInfo struct {
	// 命名空间名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 状态(Active|Terminating)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否保护(不允许删除)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protected *bool `json:"Protected,omitnil,omitempty" name:"Protected"`

	// 对应的Yaml配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`
}

type NamespaceResource struct {
	// 类型(workload|grid|configmap|secret)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 名称(最多返回5个）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

type NodeGroupInfo struct {
	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitnil,omitempty" name:"NodeGroupName"`

	// DeploymentGrid数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeploymentGridList []*GridDetail `json:"DeploymentGridList,omitnil,omitempty" name:"DeploymentGridList"`

	// StatefulSetGrid数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatefulSetGridList []*GridDetail `json:"StatefulSetGridList,omitnil,omitempty" name:"StatefulSetGridList"`

	// 是否平台保护
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protect *bool `json:"Protect,omitnil,omitempty" name:"Protect"`
}

type NodeGroupNodeUnitTemplateInfo struct {
	// 模版ID
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 包含节点列表
	NodeList []*NodeSimpleInfo `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否关联
	Relation *bool `json:"Relation,omitnil,omitempty" name:"Relation"`
}

type NodeSimpleInfo struct {
	// 节点ID
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 节点名称
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`
}

type NodeUnitInfo struct {
	// NodeUnitId
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// NodeUnit名称
	NodeUnitName *string `json:"NodeUnitName,omitnil,omitempty" name:"NodeUnitName"`

	// 包含节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeList []*NodeUnitNodeInfo `json:"NodeList,omitnil,omitempty" name:"NodeList"`
}

type NodeUnitNodeInfo struct {
	// 节点ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 节点状态  NodeStatusHealthy (健康)/NodeStatusAbnormal (异常)/NodeStatusOffline (下线)/NodeStatusNotActivated (未激活
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 节点名称
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 内网节点IP
	InternalIP *string `json:"InternalIP,omitnil,omitempty" name:"InternalIP"`
}

type NodeUnitTemplate struct {
	// NodeUnit模版ID
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// NodeUnit模版名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 包含节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeList []*NodeSimpleInfo `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// NodeGroup列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeGroups []*string `json:"NodeGroups,omitnil,omitempty" name:"NodeGroups"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type OperationLog struct {
	// 操作时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateTime *string `json:"OperateTime,omitnil,omitempty" name:"OperateTime"`

	// 模块名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 状态: 1:成功 2:失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 操作用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUserID *string `json:"OperatorUserID,omitnil,omitempty" name:"OperatorUserID"`

	// 操作动作
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`
}

type OperationLogsCondition struct {
	// 状态列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type PortConfig struct {
	// 协议类型(tcp|udp)
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 源端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 目标端口
	TargetPort *int64 `json:"TargetPort,omitnil,omitempty" name:"TargetPort"`

	// 节点端口
	NodePort *int64 `json:"NodePort,omitnil,omitempty" name:"NodePort"`
}

type Probe struct {
	// 启动后，延迟探测时间 单位:秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitialDelaySeconds *int64 `json:"InitialDelaySeconds,omitnil,omitempty" name:"InitialDelaySeconds"`

	// 探测间隔，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeriodSeconds *int64 `json:"PeriodSeconds,omitnil,omitempty" name:"PeriodSeconds"`

	// 探测超时时间 单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeoutSeconds *int64 `json:"TimeoutSeconds,omitnil,omitempty" name:"TimeoutSeconds"`

	// 失败后检查成功的最小连续成功次数。默认为1.活跃度必须为1。最小值为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessThreshold *int64 `json:"SuccessThreshold,omitnil,omitempty" name:"SuccessThreshold"`

	// 当Pod成功启动且检查失败时，放弃之前尝试次数。默认为3.最小值为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureThreshold *int64 `json:"FailureThreshold,omitnil,omitempty" name:"FailureThreshold"`

	// HTTP探测配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpProbe *HttpProbe `json:"HttpProbe,omitnil,omitempty" name:"HttpProbe"`

	// TCP探测配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TcpProbe *TcpProbe `json:"TcpProbe,omitnil,omitempty" name:"TcpProbe"`
}

// Predefined struct for user
type RedeployEdgeUnitApplicationRequestParams struct {
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
}

type RedeployEdgeUnitApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil,omitempty" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
}

func (r *RedeployEdgeUnitApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RedeployEdgeUnitApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RedeployEdgeUnitApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RedeployEdgeUnitApplicationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RedeployEdgeUnitApplicationResponse struct {
	*tchttp.BaseResponse
	Response *RedeployEdgeUnitApplicationResponseParams `json:"Response"`
}

func (r *RedeployEdgeUnitApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RedeployEdgeUnitApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceMetricTarget struct {
	// 类型(cpu|memory)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 平均值
	AverageValue *int64 `json:"AverageValue,omitnil,omitempty" name:"AverageValue"`

	// 单位
	Scale *string `json:"Scale,omitnil,omitempty" name:"Scale"`

	// 平均值
	AverageUtilization *int64 `json:"AverageUtilization,omitnil,omitempty" name:"AverageUtilization"`
}

type RouteInfo struct {
	// 无
	RouteID *int64 `json:"RouteID,omitnil,omitempty" name:"RouteID"`

	// 无
	RouteName *string `json:"RouteName,omitnil,omitempty" name:"RouteName"`

	// 无
	SourceProductID *string `json:"SourceProductID,omitnil,omitempty" name:"SourceProductID"`

	// 无
	TopicFilter *string `json:"TopicFilter,omitnil,omitempty" name:"TopicFilter"`

	// 无
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 无
	TargetOptions *string `json:"TargetOptions,omitnil,omitempty" name:"TargetOptions"`

	// 无
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 无
	Descript *string `json:"Descript,omitnil,omitempty" name:"Descript"`

	// 无
	Healthy *string `json:"Healthy,omitnil,omitempty" name:"Healthy"`

	// 无
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 无
	MessageCount *int64 `json:"MessageCount,omitnil,omitempty" name:"MessageCount"`

	// 无
	MessageLastTime *string `json:"MessageLastTime,omitnil,omitempty" name:"MessageLastTime"`

	// 无
	SourceProductName *string `json:"SourceProductName,omitnil,omitempty" name:"SourceProductName"`

	// 无
	SourceUnitIDList []*string `json:"SourceUnitIDList,omitnil,omitempty" name:"SourceUnitIDList"`

	// 无
	SourceUnitNameList []*string `json:"SourceUnitNameList,omitnil,omitempty" name:"SourceUnitNameList"`

	// 无
	SourceDeviceNameList []*string `json:"SourceDeviceNameList,omitnil,omitempty" name:"SourceDeviceNameList"`
}

type SecretItem struct {
	// Secret名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Secret类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretType *string `json:"SecretType,omitnil,omitempty" name:"SecretType"`
}

type SecurityCapabilities struct {
	// 允许操作列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Add []*string `json:"Add,omitnil,omitempty" name:"Add"`

	// 禁止操作列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Drop []*string `json:"Drop,omitnil,omitempty" name:"Drop"`
}

type SecurityContext struct {
	// 是否开启特权模式
	Privilege *bool `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 目录/Proc挂载方式
	ProcMount *string `json:"ProcMount,omitnil,omitempty" name:"ProcMount"`

	// 安全配置
	Capabilities *SecurityCapabilities `json:"Capabilities,omitnil,omitempty" name:"Capabilities"`
}

type Service struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 类型 (ClusterIP|NodePort)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 端口配置
	Ports []*PortConfig `json:"Ports,omitnil,omitempty" name:"Ports"`

	// 标签
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 命名空间默认default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 服务IP
	ClusterIP *string `json:"ClusterIP,omitnil,omitempty" name:"ClusterIP"`
}

// Predefined struct for user
type SetRouteOnOffRequestParams struct {
	// 无
	RouteID *int64 `json:"RouteID,omitnil,omitempty" name:"RouteID"`

	// on 或 off
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type SetRouteOnOffRequest struct {
	*tchttp.BaseRequest
	
	// 无
	RouteID *int64 `json:"RouteID,omitnil,omitempty" name:"RouteID"`

	// on 或 off
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *SetRouteOnOffRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetRouteOnOffRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteID")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetRouteOnOffRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetRouteOnOffResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetRouteOnOffResponse struct {
	*tchttp.BaseResponse
	Response *SetRouteOnOffResponseParams `json:"Response"`
}

func (r *SetRouteOnOffResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetRouteOnOffResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Sort struct {
	// 排序字段
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 排序方式，升序ASC / 降序DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type TcpProbe struct {
	// 连接端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type Volume struct {
	// 来源(emptyDir|hostPath|configMap|secret|nfs)
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Host挂载配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostPath *VolumeHostPath `json:"HostPath,omitnil,omitempty" name:"HostPath"`

	// ConfigMap挂载配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigMap *VolumeConfigMap `json:"ConfigMap,omitnil,omitempty" name:"ConfigMap"`

	// Secret挂载配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Secret *VolumeConfigMap `json:"Secret,omitnil,omitempty" name:"Secret"`

	// NFS挂载配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	NFS *VolumeNFS `json:"NFS,omitnil,omitempty" name:"NFS"`
}

type VolumeConfigMap struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Key列表配置
	Items []*VolumeConfigMapKeyToPath `json:"Items,omitnil,omitempty" name:"Items"`
}

type VolumeConfigMapKeyToPath struct {
	// 健名
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 对应本地路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 对应权限模式
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`
}

type VolumeHostPath struct {
	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type VolumeMount struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 挂载路径
	MountPath *string `json:"MountPath,omitnil,omitempty" name:"MountPath"`

	// 子路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubPath *string `json:"SubPath,omitnil,omitempty" name:"SubPath"`

	// 是否只读
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`
}

type VolumeNFS struct {
	// 服务地址
	Server *string `json:"Server,omitnil,omitempty" name:"Server"`

	// 对应服务器路径
	ServerPath *string `json:"ServerPath,omitnil,omitempty" name:"ServerPath"`

	// 对应本地路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}