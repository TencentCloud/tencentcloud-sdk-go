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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 工作负载类型
	WorkflowKind *string `json:"WorkflowKind,omitnil" name:"WorkflowKind"`

	// 标签信息
	Labels []*Label `json:"Labels,omitnil" name:"Labels"`

	// Grid唯一Key
	GridUniqKey *string `json:"GridUniqKey,omitnil" name:"GridUniqKey"`

	// NodeSelector标签
	NodeSelector []*Label `json:"NodeSelector,omitnil" name:"NodeSelector"`

	// 实例数
	Replicas *int64 `json:"Replicas,omitnil" name:"Replicas"`

	// 可用实例数
	AvailableReplicas *int64 `json:"AvailableReplicas,omitnil" name:"AvailableReplicas"`

	// 是否开启service环境变量注入pod
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableServiceLinks *bool `json:"EnableServiceLinks,omitnil" name:"EnableServiceLinks"`
}

type ApplicationBasicInfo struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 管理URL地址
	ManageUrl *string `json:"ManageUrl,omitnil" name:"ManageUrl"`

	// 描述信息
	Description *string `json:"Description,omitnil" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 是否允许可视化修改
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowVisualModify *bool `json:"AllowVisualModify,omitnil" name:"AllowVisualModify"`
}

type ApplicationDeployMode struct {
	// 1:指定节点部署 2:单元部署
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 资源ID
	ResourceID *uint64 `json:"ResourceID,omitnil" name:"ResourceID"`

	// 资源名
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`
}

type ApplicationStatusInfo struct {
	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 应用版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil" name:"Version"`

	// 应用状态(1:待部署 2:部署中 3:运行中 4:待更新 5:更新中 6:待删除 7:删除中 8:已删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 管理地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManageUrl *string `json:"ManageUrl,omitnil" name:"ManageUrl"`

	// 负载类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkloadKind *string `json:"WorkloadKind,omitnil" name:"WorkloadKind"`

	// 应用部署模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployMode *ApplicationDeployMode `json:"DeployMode,omitnil" name:"DeployMode"`

	// 期望Pod数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas *int64 `json:"Replicas,omitnil" name:"Replicas"`

	// 运行Pod数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableReplicas *int64 `json:"AvailableReplicas,omitnil" name:"AvailableReplicas"`
}

type ApplicationTemplate struct {
	// 模板ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 来源。1 自定义应用模板 ;  2 官方应用模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *uint64 `json:"Source,omitnil" name:"Source"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkloadKind *string `json:"WorkloadKind,omitnil" name:"WorkloadKind"`

	// 管理地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManageUrl *string `json:"ManageUrl,omitnil" name:"ManageUrl"`

	// 发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DistributeTime *string `json:"DistributeTime,omitnil" name:"DistributeTime"`
}

// Predefined struct for user
type ApplyMarketComponentRequestParams struct {
	// 组件ID
	ID *int64 `json:"ID,omitnil" name:"ID"`
}

type ApplyMarketComponentRequest struct {
	*tchttp.BaseRequest
	
	// 组件ID
	ID *int64 `json:"ID,omitnil" name:"ID"`
}

func (r *ApplyMarketComponentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyMarketComponentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyMarketComponentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyMarketComponentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ApplyMarketComponentResponse struct {
	*tchttp.BaseResponse
	Response *ApplyMarketComponentResponseParams `json:"Response"`
}

func (r *ApplyMarketComponentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyMarketComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BuildMessageRouteRequestParams struct {
	// 路由名字
	RouteName *string `json:"RouteName,omitnil" name:"RouteName"`

	// 源产品id
	SourceProductID *string `json:"SourceProductID,omitnil" name:"SourceProductID"`

	// 源设备名列表
	SourceDeviceNameList []*string `json:"SourceDeviceNameList,omitnil" name:"SourceDeviceNameList"`

	// 第一个字符为 "0"或"1"，"1"表示自定义topic
	TopicFilter *string `json:"TopicFilter,omitnil" name:"TopicFilter"`

	// http或mqtt-broker
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 源单元id列表
	SourceUnitIDList []*string `json:"SourceUnitIDList,omitnil" name:"SourceUnitIDList"`

	// 描述
	Descript *string `json:"Descript,omitnil" name:"Descript"`

	// 无
	TargetOptions *string `json:"TargetOptions,omitnil" name:"TargetOptions"`
}

type BuildMessageRouteRequest struct {
	*tchttp.BaseRequest
	
	// 路由名字
	RouteName *string `json:"RouteName,omitnil" name:"RouteName"`

	// 源产品id
	SourceProductID *string `json:"SourceProductID,omitnil" name:"SourceProductID"`

	// 源设备名列表
	SourceDeviceNameList []*string `json:"SourceDeviceNameList,omitnil" name:"SourceDeviceNameList"`

	// 第一个字符为 "0"或"1"，"1"表示自定义topic
	TopicFilter *string `json:"TopicFilter,omitnil" name:"TopicFilter"`

	// http或mqtt-broker
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 源单元id列表
	SourceUnitIDList []*string `json:"SourceUnitIDList,omitnil" name:"SourceUnitIDList"`

	// 描述
	Descript *string `json:"Descript,omitnil" name:"Descript"`

	// 无
	TargetOptions *string `json:"TargetOptions,omitnil" name:"TargetOptions"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`
}

type Container struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 镜像名
	ImageName *string `json:"ImageName,omitnil" name:"ImageName"`

	// 镜像版本
	ImageVersion *string `json:"ImageVersion,omitnil" name:"ImageVersion"`

	// 镜像拉取策略(Always|Never|IfNotPresent)
	ImagePullPolicy *string `json:"ImagePullPolicy,omitnil" name:"ImagePullPolicy"`

	// 卷挂载配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeMounts []*VolumeMount `json:"VolumeMounts,omitnil" name:"VolumeMounts"`

	// cpu最低配置
	CpuRequest *string `json:"CpuRequest,omitnil" name:"CpuRequest"`

	// cpu最高限制
	CpuLimit *string `json:"CpuLimit,omitnil" name:"CpuLimit"`

	// 内存最低要求
	MemoryRequest *string `json:"MemoryRequest,omitnil" name:"MemoryRequest"`

	// 内存最高要求
	MemoryLimit *string `json:"MemoryLimit,omitnil" name:"MemoryLimit"`

	// 内存单位
	MemoryUnit *string `json:"MemoryUnit,omitnil" name:"MemoryUnit"`

	// gpu最高限制
	GpuLimit *string `json:"GpuLimit,omitnil" name:"GpuLimit"`

	// 资源配置
	ResourceMapCloud []*KeyValueObj `json:"ResourceMapCloud,omitnil" name:"ResourceMapCloud"`

	// 环境配置
	Envs []*Env `json:"Envs,omitnil" name:"Envs"`

	// 工作目录
	WorkingDir *string `json:"WorkingDir,omitnil" name:"WorkingDir"`

	// 命令
	Commands []*string `json:"Commands,omitnil" name:"Commands"`

	// 参数
	Args []*string `json:"Args,omitnil" name:"Args"`

	// 安全配置
	SecurityContext *SecurityContext `json:"SecurityContext,omitnil" name:"SecurityContext"`

	// 就绪探针配置
	ReadinessProbe *Probe `json:"ReadinessProbe,omitnil" name:"ReadinessProbe"`
}

type ContainerStatus struct {
	// 容器名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 容器ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *string `json:"ID,omitnil" name:"ID"`

	// 镜像
	// 注意：此字段可能返回 null，表示取不到有效值。
	Image *string `json:"Image,omitnil" name:"Image"`

	// 重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *int64 `json:"RestartCount,omitnil" name:"RestartCount"`

	// 状态
	Status *string `json:"Status,omitnil" name:"Status"`
}

// Predefined struct for user
type CreateApplicationVisualizationRequestParams struct {
	// 基本信息
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitnil" name:"BasicInfo"`

	// 基本配置
	BasicConfig *ApplicationBasicConfig `json:"BasicConfig,omitnil" name:"BasicConfig"`

	// 卷列表
	Volumes []*Volume `json:"Volumes,omitnil" name:"Volumes"`

	// 服务配置
	Service *Service `json:"Service,omitnil" name:"Service"`

	// Job配置
	Job *Job `json:"Job,omitnil" name:"Job"`

	// CronJob配置
	CronJob *CronJob `json:"CronJob,omitnil" name:"CronJob"`

	// 重新运行策略
	RestartPolicy *string `json:"RestartPolicy,omitnil" name:"RestartPolicy"`

	// 镜像拉取密钥
	ImagePullSecrets []*string `json:"ImagePullSecrets,omitnil" name:"ImagePullSecrets"`

	// HPA配置
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil" name:"HorizontalPodAutoscaler"`

	// 初始化容器列表
	InitContainers []*Container `json:"InitContainers,omitnil" name:"InitContainers"`

	// 容器列表
	Containers []*Container `json:"Containers,omitnil" name:"Containers"`
}

type CreateApplicationVisualizationRequest struct {
	*tchttp.BaseRequest
	
	// 基本信息
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitnil" name:"BasicInfo"`

	// 基本配置
	BasicConfig *ApplicationBasicConfig `json:"BasicConfig,omitnil" name:"BasicConfig"`

	// 卷列表
	Volumes []*Volume `json:"Volumes,omitnil" name:"Volumes"`

	// 服务配置
	Service *Service `json:"Service,omitnil" name:"Service"`

	// Job配置
	Job *Job `json:"Job,omitnil" name:"Job"`

	// CronJob配置
	CronJob *CronJob `json:"CronJob,omitnil" name:"CronJob"`

	// 重新运行策略
	RestartPolicy *string `json:"RestartPolicy,omitnil" name:"RestartPolicy"`

	// 镜像拉取密钥
	ImagePullSecrets []*string `json:"ImagePullSecrets,omitnil" name:"ImagePullSecrets"`

	// HPA配置
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil" name:"HorizontalPodAutoscaler"`

	// 初始化容器列表
	InitContainers []*Container `json:"InitContainers,omitnil" name:"InitContainers"`

	// 容器列表
	Containers []*Container `json:"Containers,omitnil" name:"Containers"`
}

func (r *CreateApplicationVisualizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationVisualizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BasicInfo")
	delete(f, "BasicConfig")
	delete(f, "Volumes")
	delete(f, "Service")
	delete(f, "Job")
	delete(f, "CronJob")
	delete(f, "RestartPolicy")
	delete(f, "ImagePullSecrets")
	delete(f, "HorizontalPodAutoscaler")
	delete(f, "InitContainers")
	delete(f, "Containers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationVisualizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationVisualizationResponseParams struct {
	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateApplicationVisualizationResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationVisualizationResponseParams `json:"Response"`
}

func (r *CreateApplicationVisualizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationVisualizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigMapRequestParams struct {
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// ConfigMap名称
	ConfigMapName *string `json:"ConfigMapName,omitnil" name:"ConfigMapName"`

	// ConfigMap内容
	ConfigMapData []*KeyValueObj `json:"ConfigMapData,omitnil" name:"ConfigMapData"`

	// ConfigMap命名空间,默认：default
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil" name:"ConfigMapNamespace"`
}

type CreateConfigMapRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// ConfigMap名称
	ConfigMapName *string `json:"ConfigMapName,omitnil" name:"ConfigMapName"`

	// ConfigMap内容
	ConfigMapData []*KeyValueObj `json:"ConfigMapData,omitnil" name:"ConfigMapData"`

	// ConfigMap命名空间,默认：default
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil" name:"ConfigMapNamespace"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 节点信息
	Nodes []*DracoNodeInfo `json:"Nodes,omitnil" name:"Nodes"`
}

type CreateEdgeNodeBatchRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 节点信息
	Nodes []*DracoNodeInfo `json:"Nodes,omitnil" name:"Nodes"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// NodeGroup名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命名空间，不填默认为default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 模版ID数组
	NodeUnitTemplateIDs []*uint64 `json:"NodeUnitTemplateIDs,omitnil" name:"NodeUnitTemplateIDs"`
}

type CreateEdgeNodeGroupRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// NodeGroup名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命名空间，不填默认为default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 模版ID数组
	NodeUnitTemplateIDs []*uint64 `json:"NodeUnitTemplateIDs,omitnil" name:"NodeUnitTemplateIDs"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 节点名称
	Name *string `json:"Name,omitnil" name:"Name"`
}

type CreateEdgeNodeRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 节点名称
	Name *string `json:"Name,omitnil" name:"Name"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// NodeUnit模板名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 包含的节点列表
	Nodes []*string `json:"Nodes,omitnil" name:"Nodes"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`
}

type CreateEdgeNodeUnitTemplateRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// NodeUnit模板名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 包含的节点列表
	Nodes []*string `json:"Nodes,omitnil" name:"Nodes"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type CreateEdgeUnitApplicationVisualizationRequestParams struct {
	// 基本信息
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitnil" name:"BasicInfo"`

	// 基本配置
	BasicConfig *ApplicationBasicConfig `json:"BasicConfig,omitnil" name:"BasicConfig"`

	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 卷列表
	Volumes []*Volume `json:"Volumes,omitnil" name:"Volumes"`

	// 服务配置
	Service *Service `json:"Service,omitnil" name:"Service"`

	// 模版ID
	TemplateID *uint64 `json:"TemplateID,omitnil" name:"TemplateID"`

	// Job配置
	Job *Job `json:"Job,omitnil" name:"Job"`

	// CronJob配置
	CronJob *CronJob `json:"CronJob,omitnil" name:"CronJob"`

	// 重新运行策略
	RestartPolicy *string `json:"RestartPolicy,omitnil" name:"RestartPolicy"`

	// 镜像拉取密钥
	ImagePullSecrets []*string `json:"ImagePullSecrets,omitnil" name:"ImagePullSecrets"`

	// HPA配置
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil" name:"HorizontalPodAutoscaler"`

	// 初始化容器列表
	InitContainers []*Container `json:"InitContainers,omitnil" name:"InitContainers"`

	// 容器列表
	Containers []*Container `json:"Containers,omitnil" name:"Containers"`
}

type CreateEdgeUnitApplicationVisualizationRequest struct {
	*tchttp.BaseRequest
	
	// 基本信息
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitnil" name:"BasicInfo"`

	// 基本配置
	BasicConfig *ApplicationBasicConfig `json:"BasicConfig,omitnil" name:"BasicConfig"`

	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 卷列表
	Volumes []*Volume `json:"Volumes,omitnil" name:"Volumes"`

	// 服务配置
	Service *Service `json:"Service,omitnil" name:"Service"`

	// 模版ID
	TemplateID *uint64 `json:"TemplateID,omitnil" name:"TemplateID"`

	// Job配置
	Job *Job `json:"Job,omitnil" name:"Job"`

	// CronJob配置
	CronJob *CronJob `json:"CronJob,omitnil" name:"CronJob"`

	// 重新运行策略
	RestartPolicy *string `json:"RestartPolicy,omitnil" name:"RestartPolicy"`

	// 镜像拉取密钥
	ImagePullSecrets []*string `json:"ImagePullSecrets,omitnil" name:"ImagePullSecrets"`

	// HPA配置
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil" name:"HorizontalPodAutoscaler"`

	// 初始化容器列表
	InitContainers []*Container `json:"InitContainers,omitnil" name:"InitContainers"`

	// 容器列表
	Containers []*Container `json:"Containers,omitnil" name:"Containers"`
}

func (r *CreateEdgeUnitApplicationVisualizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeUnitApplicationVisualizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BasicInfo")
	delete(f, "BasicConfig")
	delete(f, "EdgeUnitId")
	delete(f, "Volumes")
	delete(f, "Service")
	delete(f, "TemplateID")
	delete(f, "Job")
	delete(f, "CronJob")
	delete(f, "RestartPolicy")
	delete(f, "ImagePullSecrets")
	delete(f, "HorizontalPodAutoscaler")
	delete(f, "InitContainers")
	delete(f, "Containers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEdgeUnitApplicationVisualizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeUnitApplicationVisualizationResponseParams struct {
	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateEdgeUnitApplicationVisualizationResponse struct {
	*tchttp.BaseResponse
	Response *CreateEdgeUnitApplicationVisualizationResponseParams `json:"Response"`
}

func (r *CreateEdgeUnitApplicationVisualizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeUnitApplicationVisualizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeUnitApplicationYamlRequestParams struct {
	// 单元ID
	EdgeUnitId *int64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// base64后的Yaml配置
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`

	// 基本信息
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitnil" name:"BasicInfo"`
}

type CreateEdgeUnitApplicationYamlRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitId *int64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// base64后的Yaml配置
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`

	// 基本信息
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitnil" name:"BasicInfo"`
}

func (r *CreateEdgeUnitApplicationYamlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeUnitApplicationYamlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "Yaml")
	delete(f, "BasicInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEdgeUnitApplicationYamlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeUnitApplicationYamlResponseParams struct {
	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateEdgeUnitApplicationYamlResponse struct {
	*tchttp.BaseResponse
	Response *CreateEdgeUnitApplicationYamlResponseParams `json:"Response"`
}

func (r *CreateEdgeUnitApplicationYamlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeUnitApplicationYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeUnitCloudRequestParams struct {
	// 集群名称，长度小于32
	Name *string `json:"Name,omitnil" name:"Name"`

	// k8s版本，仅支持1.16.7 和 1.18.2
	K8sVersion *string `json:"K8sVersion,omitnil" name:"K8sVersion"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 集群描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 集群pod cidr， 默认  10.1.0.0/16
	PodCIDR *string `json:"PodCIDR,omitnil" name:"PodCIDR"`

	// 集群service cidr, 默认 10.2.0.0/16
	ServiceCIDR *string `json:"ServiceCIDR,omitnil" name:"ServiceCIDR"`

	// 是否开启监控。目前内存中权限开启联系产品开通白名单
	OpenCloudMonitor *bool `json:"OpenCloudMonitor,omitnil" name:"OpenCloudMonitor"`
}

type CreateEdgeUnitCloudRequest struct {
	*tchttp.BaseRequest
	
	// 集群名称，长度小于32
	Name *string `json:"Name,omitnil" name:"Name"`

	// k8s版本，仅支持1.16.7 和 1.18.2
	K8sVersion *string `json:"K8sVersion,omitnil" name:"K8sVersion"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 集群描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 集群pod cidr， 默认  10.1.0.0/16
	PodCIDR *string `json:"PodCIDR,omitnil" name:"PodCIDR"`

	// 集群service cidr, 默认 10.2.0.0/16
	ServiceCIDR *string `json:"ServiceCIDR,omitnil" name:"ServiceCIDR"`

	// 是否开启监控。目前内存中权限开启联系产品开通白名单
	OpenCloudMonitor *bool `json:"OpenCloudMonitor,omitnil" name:"OpenCloudMonitor"`
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
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// IECP集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *int64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 无
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// 无
	DeviceNames []*string `json:"DeviceNames,omitnil" name:"DeviceNames"`
}

type CreateEdgeUnitDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 无
	EdgeUnitId *int64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 无
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// 无
	DeviceNames []*string `json:"DeviceNames,omitnil" name:"DeviceNames"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 设备所属的产品id
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 无
	UnitID *int64 `json:"UnitID,omitnil" name:"UnitID"`
}

type CreateIotDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 设备所属的产品id
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 无
	UnitID *int64 `json:"UnitID,omitnil" name:"UnitID"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RouteName *string `json:"RouteName,omitnil" name:"RouteName"`

	// 路由备注
	Descript *string `json:"Descript,omitnil" name:"Descript"`
}

type CreateMessageRouteRequest struct {
	*tchttp.BaseRequest
	
	// 路由名称
	RouteName *string `json:"RouteName,omitnil" name:"RouteName"`

	// 路由备注
	Descript *string `json:"Descript,omitnil" name:"Descript"`
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
	RouteID *int64 `json:"RouteID,omitnil" name:"RouteID"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 描述信息
	Description *string `json:"Description,omitnil" name:"Description"`
}

type CreateNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 描述信息
	Description *string `json:"Description,omitnil" name:"Description"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// secret名
	SecretName *string `json:"SecretName,omitnil" name:"SecretName"`

	// 命名空间（默认:default）
	SecretNamespace *string `json:"SecretNamespace,omitnil" name:"SecretNamespace"`

	// secret类型(取值范围:DockerConfigJson,Opaque 默认Opaque)
	SecretType *string `json:"SecretType,omitnil" name:"SecretType"`

	// DockerConfig的序列化base64编码后的字符串
	DockerConfigJson *string `json:"DockerConfigJson,omitnil" name:"DockerConfigJson"`

	// Opaque类型的Secret内容
	CloudData []*KeyValueObj `json:"CloudData,omitnil" name:"CloudData"`

	// DockerConfig配置
	DockerConfig *DockerConfig `json:"DockerConfig,omitnil" name:"DockerConfig"`
}

type CreateSecretRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// secret名
	SecretName *string `json:"SecretName,omitnil" name:"SecretName"`

	// 命名空间（默认:default）
	SecretNamespace *string `json:"SecretNamespace,omitnil" name:"SecretNamespace"`

	// secret类型(取值范围:DockerConfigJson,Opaque 默认Opaque)
	SecretType *string `json:"SecretType,omitnil" name:"SecretType"`

	// DockerConfig的序列化base64编码后的字符串
	DockerConfigJson *string `json:"DockerConfigJson,omitnil" name:"DockerConfigJson"`

	// Opaque类型的Secret内容
	CloudData []*KeyValueObj `json:"CloudData,omitnil" name:"CloudData"`

	// DockerConfig配置
	DockerConfig *DockerConfig `json:"DockerConfig,omitnil" name:"DockerConfig"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// NodeUnit所属的NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitnil" name:"NodeGroupName"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// NodeUnit名称，通过模版创建可不填
	NodeUnitName *string `json:"NodeUnitName,omitnil" name:"NodeUnitName"`

	// NodeUnit包含的节点列表，通过模版创建可不填
	Nodes []*string `json:"Nodes,omitnil" name:"Nodes"`

	// NodeUnit模版ID列表
	NodeUnitTemplateIDs []*uint64 `json:"NodeUnitTemplateIDs,omitnil" name:"NodeUnitTemplateIDs"`
}

type CreateUpdateNodeUnitRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// NodeUnit所属的NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitnil" name:"NodeGroupName"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// NodeUnit名称，通过模版创建可不填
	NodeUnitName *string `json:"NodeUnitName,omitnil" name:"NodeUnitName"`

	// NodeUnit包含的节点列表，通过模版创建可不填
	Nodes []*string `json:"Nodes,omitnil" name:"Nodes"`

	// NodeUnit模版ID列表
	NodeUnitTemplateIDs []*uint64 `json:"NodeUnitTemplateIDs,omitnil" name:"NodeUnitTemplateIDs"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Second *int64 `json:"Second,omitnil" name:"Second"`
}

type CreateUserTokenRequest struct {
	*tchttp.BaseRequest
	
	// token过期时间，有效值是1~300秒
	Second *int64 `json:"Second,omitnil" name:"Second"`
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
	Token *string `json:"Token,omitnil" name:"Token"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Schedule *string `json:"Schedule,omitnil" name:"Schedule"`

	// 运行时间
	StartingDeadlineSeconds *int64 `json:"StartingDeadlineSeconds,omitnil" name:"StartingDeadlineSeconds"`

	// job并行策略(Allow|Forbid|Replace)
	ConcurrencyPolicy *string `json:"ConcurrencyPolicy,omitnil" name:"ConcurrencyPolicy"`

	// Job配置
	Job *Job `json:"Job,omitnil" name:"Job"`
}

// Predefined struct for user
type DeleteApplicationsRequestParams struct {
	// 应用模板ID列表
	ApplicationIds []*uint64 `json:"ApplicationIds,omitnil" name:"ApplicationIds"`
}

type DeleteApplicationsRequest struct {
	*tchttp.BaseRequest
	
	// 应用模板ID列表
	ApplicationIds []*uint64 `json:"ApplicationIds,omitnil" name:"ApplicationIds"`
}

func (r *DeleteApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationsResponseParams `json:"Response"`
}

func (r *DeleteApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigMapRequestParams struct {
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// ConfigMap名
	ConfigMapName *string `json:"ConfigMapName,omitnil" name:"ConfigMapName"`

	// ConfigMap命名空间，默认：default
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil" name:"ConfigMapNamespace"`
}

type DeleteConfigMapRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// ConfigMap名
	ConfigMapName *string `json:"ConfigMapName,omitnil" name:"ConfigMapName"`

	// ConfigMap命名空间，默认：default
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil" name:"ConfigMapNamespace"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// NodeGroup名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
}

type DeleteEdgeNodeGroupRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// NodeGroup名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 删除的NodeUnit模板ID列表
	NodeUnitTemplateIDs []*uint64 `json:"NodeUnitTemplateIDs,omitnil" name:"NodeUnitTemplateIDs"`
}

type DeleteEdgeNodeUnitTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 删除的NodeUnit模板ID列表
	NodeUnitTemplateIDs []*uint64 `json:"NodeUnitTemplateIDs,omitnil" name:"NodeUnitTemplateIDs"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// IECP边缘节点ID列表
	NodeIds []*uint64 `json:"NodeIds,omitnil" name:"NodeIds"`
}

type DeleteEdgeNodesRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// IECP边缘节点ID列表
	NodeIds []*uint64 `json:"NodeIds,omitnil" name:"NodeIds"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DeleteEdgeUnitApplicationsRequestParams struct {
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// 应用ID列表
	ApplicationIDs []*uint64 `json:"ApplicationIDs,omitnil" name:"ApplicationIDs"`
}

type DeleteEdgeUnitApplicationsRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// 应用ID列表
	ApplicationIDs []*uint64 `json:"ApplicationIDs,omitnil" name:"ApplicationIDs"`
}

func (r *DeleteEdgeUnitApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeUnitApplicationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitID")
	delete(f, "ApplicationIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEdgeUnitApplicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeUnitApplicationsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteEdgeUnitApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEdgeUnitApplicationsResponseParams `json:"Response"`
}

func (r *DeleteEdgeUnitApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeUnitApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeUnitCloudRequestParams struct {
	// 边缘集群ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`
}

type DeleteEdgeUnitCloudRequest struct {
	*tchttp.BaseRequest
	
	// 边缘集群ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil" name:"WorkloadKind"`

	// Grid部署名称
	GridItemName *string `json:"GridItemName,omitnil" name:"GridItemName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
}

type DeleteEdgeUnitDeployGridItemRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil" name:"WorkloadKind"`

	// Grid部署名称
	GridItemName *string `json:"GridItemName,omitnil" name:"GridItemName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// 无
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`
}

// Predefined struct for user
type DeleteEdgeUnitDevicesRequestParams struct {
	// 无
	EdgeUnitId *int64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 无
	Devices []*DeleteEdgeUnitDevicesDevice `json:"Devices,omitnil" name:"Devices"`
}

type DeleteEdgeUnitDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 无
	EdgeUnitId *int64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 无
	Devices []*DeleteEdgeUnitDevicesDevice `json:"Devices,omitnil" name:"Devices"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ClusterID *string `json:"ClusterID,omitnil" name:"ClusterID"`

	// Pod名称
	PodName *string `json:"PodName,omitnil" name:"PodName"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
}

type DeleteEdgeUnitPodRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil" name:"ClusterID"`

	// Pod名称
	PodName *string `json:"PodName,omitnil" name:"PodName"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	DeviceIDList []*uint64 `json:"DeviceIDList,omitnil" name:"DeviceIDList"`
}

type DeleteIotDeviceBatchRequest struct {
	*tchttp.BaseRequest
	
	// 无
	DeviceIDList []*uint64 `json:"DeviceIDList,omitnil" name:"DeviceIDList"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	DeviceId *int64 `json:"DeviceId,omitnil" name:"DeviceId"`
}

type DeleteIotDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备id
	DeviceId *int64 `json:"DeviceId,omitnil" name:"DeviceId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RouteID *int64 `json:"RouteID,omitnil" name:"RouteID"`
}

type DeleteMessageRouteRequest struct {
	*tchttp.BaseRequest
	
	// 无
	RouteID *int64 `json:"RouteID,omitnil" name:"RouteID"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
}

type DeleteNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// NodeUnit所属的NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitnil" name:"NodeGroupName"`

	// NodeUnit名称
	NodeUnitName *string `json:"NodeUnitName,omitnil" name:"NodeUnitName"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// NodeUnit包含的节点列表
	Nodes []*string `json:"Nodes,omitnil" name:"Nodes"`
}

type DeleteNodeUnitRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// NodeUnit所属的NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitnil" name:"NodeGroupName"`

	// NodeUnit名称
	NodeUnitName *string `json:"NodeUnitName,omitnil" name:"NodeUnitName"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// NodeUnit包含的节点列表
	Nodes []*string `json:"Nodes,omitnil" name:"Nodes"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// secret名称
	SecretName *string `json:"SecretName,omitnil" name:"SecretName"`

	// secret命名空间（默认:default）
	SecretNamespace *string `json:"SecretNamespace,omitnil" name:"SecretNamespace"`
}

type DeleteSecretRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// secret名称
	SecretName *string `json:"SecretName,omitnil" name:"SecretName"`

	// secret命名空间（默认:default）
	SecretNamespace *string `json:"SecretNamespace,omitnil" name:"SecretNamespace"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeApplicationVisualizationRequestParams struct {
	// 应用模板ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`
}

type DescribeApplicationVisualizationRequest struct {
	*tchttp.BaseRequest
	
	// 应用模板ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`
}

func (r *DescribeApplicationVisualizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationVisualizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationVisualizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationVisualizationResponseParams struct {
	// 基本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitnil" name:"BasicInfo"`

	// 基本配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	BasicConfig *ApplicationBasicConfig `json:"BasicConfig,omitnil" name:"BasicConfig"`

	// 卷配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Volumes []*Volume `json:"Volumes,omitnil" name:"Volumes"`

	// 初始化容器配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitContainers []*Container `json:"InitContainers,omitnil" name:"InitContainers"`

	// 容器配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Containers []*Container `json:"Containers,omitnil" name:"Containers"`

	// 服务配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Service *Service `json:"Service,omitnil" name:"Service"`

	// Job配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Job *Job `json:"Job,omitnil" name:"Job"`

	// CronJob配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CronJob *CronJob `json:"CronJob,omitnil" name:"CronJob"`

	// 重启策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartPolicy *string `json:"RestartPolicy,omitnil" name:"RestartPolicy"`

	// HPA
	// 注意：此字段可能返回 null，表示取不到有效值。
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil" name:"HorizontalPodAutoscaler"`

	// 镜像拉取Secret
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImagePullSecrets []*string `json:"ImagePullSecrets,omitnil" name:"ImagePullSecrets"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeApplicationVisualizationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationVisualizationResponseParams `json:"Response"`
}

func (r *DescribeApplicationVisualizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationVisualizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationYamlErrorRequestParams struct {
	// Yaml配置
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`
}

type DescribeApplicationYamlErrorRequest struct {
	*tchttp.BaseRequest
	
	// Yaml配置
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`
}

func (r *DescribeApplicationYamlErrorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationYamlErrorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Yaml")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationYamlErrorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationYamlErrorResponseParams struct {
	// 是否通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckPass *bool `json:"CheckPass,omitnil" name:"CheckPass"`

	// 错误类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrType *int64 `json:"ErrType,omitnil" name:"ErrType"`

	// 错误信息
	ErrInfo *string `json:"ErrInfo,omitnil" name:"ErrInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeApplicationYamlErrorResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationYamlErrorResponseParams `json:"Response"`
}

func (r *DescribeApplicationYamlErrorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationYamlErrorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationYamlRequestParams struct {
	// 应用模板ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`
}

type DescribeApplicationYamlRequest struct {
	*tchttp.BaseRequest
	
	// 应用模板ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`
}

func (r *DescribeApplicationYamlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationYamlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationYamlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationYamlResponseParams struct {
	// base64 后的yaml
	// 注意：此字段可能返回 null，表示取不到有效值。
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeApplicationYamlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationYamlResponseParams `json:"Response"`
}

func (r *DescribeApplicationYamlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationsRequestParams struct {
	// 模糊搜索字符串
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`

	// 默认 0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 默认 20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 仅支持对 DistributeTime 字段排序，ASC/DESC
	Sort []*FieldSort `json:"Sort,omitnil" name:"Sort"`
}

type DescribeApplicationsRequest struct {
	*tchttp.BaseRequest
	
	// 模糊搜索字符串
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`

	// 默认 0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 默认 20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 仅支持对 DistributeTime 字段排序，ASC/DESC
	Sort []*FieldSort `json:"Sort,omitnil" name:"Sort"`
}

func (r *DescribeApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NamePattern")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationsResponseParams struct {
	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 详细列表
	ApplicationSet []*ApplicationTemplate `json:"ApplicationSet,omitnil" name:"ApplicationSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationsResponseParams `json:"Response"`
}

func (r *DescribeApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigMapRequestParams struct {
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// ConfigMap名称
	ConfigMapName *string `json:"ConfigMapName,omitnil" name:"ConfigMapName"`

	// ConfigMap命名空间
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil" name:"ConfigMapNamespace"`
}

type DescribeConfigMapRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// ConfigMap名称
	ConfigMapName *string `json:"ConfigMapName,omitnil" name:"ConfigMapName"`

	// ConfigMap命名空间
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil" name:"ConfigMapNamespace"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// yaml配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`

	// 配置项的json格式(base64编码)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Json *string `json:"Json,omitnil" name:"Json"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`
}

type DescribeConfigMapYamlErrorRequest struct {
	*tchttp.BaseRequest
	
	// yaml文件
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`
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
	CheckPass *bool `json:"CheckPass,omitnil" name:"CheckPass"`

	// 错误类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrType *uint64 `json:"ErrType,omitnil" name:"ErrType"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrInfo *string `json:"ErrInfo,omitnil" name:"ErrInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// 翻页偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页大小(最大100)
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 命名空间
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil" name:"ConfigMapNamespace"`

	// 模糊匹配的名称
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`

	// Sort.Fileld填写CreateTime Sort.Order(ASC|DESC) 默认ASC
	Sort *FieldSort `json:"Sort,omitnil" name:"Sort"`
}

type DescribeConfigMapsRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// 翻页偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页大小(最大100)
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 命名空间
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil" name:"ConfigMapNamespace"`

	// 模糊匹配的名称
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`

	// Sort.Fileld填写CreateTime Sort.Order(ASC|DESC) 默认ASC
	Sort *FieldSort `json:"Sort,omitnil" name:"Sort"`
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
	Items []*ConfigMapBasicInfo `json:"Items,omitnil" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	SN *string `json:"SN,omitnil" name:"SN"`
}

type DescribeDracoEdgeNodeInstallerRequest struct {
	*tchttp.BaseRequest
	
	// 设备SN
	SN *string `json:"SN,omitnil" name:"SN"`
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
	OnlineInstallationCommand *string `json:"OnlineInstallationCommand,omitnil" name:"OnlineInstallationCommand"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// IECP边缘节点ID
	NodeId *uint64 `json:"NodeId,omitnil" name:"NodeId"`
}

type DescribeEdgeAgentNodeInstallerRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// IECP边缘节点ID
	NodeId *uint64 `json:"NodeId,omitnil" name:"NodeId"`
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
	Online *EdgeNodeInstallerOnline `json:"Online,omitnil" name:"Online"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 网络CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcCidrBlock *string `json:"VpcCidrBlock,omitnil" name:"VpcCidrBlock"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 子网CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetCidrBlock *string `json:"SubnetCidrBlock,omitnil" name:"SubnetCidrBlock"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 节点ID
	NodeId *uint64 `json:"NodeId,omitnil" name:"NodeId"`

	// Pod名称
	PodName *string `json:"PodName,omitnil" name:"PodName"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
}

type DescribeEdgeNodePodContainersRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 节点ID
	NodeId *uint64 `json:"NodeId,omitnil" name:"NodeId"`

	// Pod名称
	PodName *string `json:"PodName,omitnil" name:"PodName"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
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
	ContainerSet []*EdgeNodePodContainerInfo `json:"ContainerSet,omitnil" name:"ContainerSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 节点ID
	NodeId *uint64 `json:"NodeId,omitnil" name:"NodeId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// Pod名称过滤串
	PodNamePattern *string `json:"PodNamePattern,omitnil" name:"PodNamePattern"`
}

type DescribeEdgeNodePodsRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 节点ID
	NodeId *uint64 `json:"NodeId,omitnil" name:"NodeId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// Pod名称过滤串
	PodNamePattern *string `json:"PodNamePattern,omitnil" name:"PodNamePattern"`
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
	PodSet []*EdgeNodePodInfo `json:"PodSet,omitnil" name:"PodSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`
}

type DescribeEdgeNodeRemarkListRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`
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
	Remarks []*string `json:"Remarks,omitnil" name:"Remarks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// IECP边缘节点ID
	NodeId *uint64 `json:"NodeId,omitnil" name:"NodeId"`
}

type DescribeEdgeNodeRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// IECP边缘节点ID
	NodeId *uint64 `json:"NodeId,omitnil" name:"NodeId"`
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
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 节点类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// 节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 节点状态 （1健康｜2异常｜3离线｜4未激活）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// CPU体系结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuArchitecture *string `json:"CpuArchitecture,omitnil" name:"CpuArchitecture"`

	// AI处理器体系结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiChipArchitecture *string `json:"AiChipArchitecture,omitnil" name:"AiChipArchitecture"`

	// IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 节点标签列表
	Labels []*EdgeNodeLabel `json:"Labels,omitnil" name:"Labels"`

	// 节点资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *EdgeNodeResourceInfo `json:"Resource,omitnil" name:"Resource"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 边缘节点名称模糊搜索串
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`

	// 边缘节点名称列表，支持批量查询 ，优先于模糊查询
	NameMatchedList []*string `json:"NameMatchedList,omitnil" name:"NameMatchedList"`

	// 排序信息列表
	Sort []*Sort `json:"Sort,omitnil" name:"Sort"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 页面大小Limit
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 节点类型
	NodeType *int64 `json:"NodeType,omitnil" name:"NodeType"`
}

type DescribeEdgeNodesRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 边缘节点名称模糊搜索串
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`

	// 边缘节点名称列表，支持批量查询 ，优先于模糊查询
	NameMatchedList []*string `json:"NameMatchedList,omitnil" name:"NameMatchedList"`

	// 排序信息列表
	Sort []*Sort `json:"Sort,omitnil" name:"Sort"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 页面大小Limit
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 节点类型
	NodeType *int64 `json:"NodeType,omitnil" name:"NodeType"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeSet []*EdgeNodeInfo `json:"NodeSet,omitnil" name:"NodeSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段
	Sort []*FieldSort `json:"Sort,omitnil" name:"Sort"`

	// 模块
	Module *string `json:"Module,omitnil" name:"Module"`

	// 过滤条件
	Condition *OperationLogsCondition `json:"Condition,omitnil" name:"Condition"`
}

type DescribeEdgeOperationLogsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段
	Sort []*FieldSort `json:"Sort,omitnil" name:"Sort"`

	// 模块
	Module *string `json:"Module,omitnil" name:"Module"`

	// 过滤条件
	Condition *OperationLogsCondition `json:"Condition,omitnil" name:"Condition"`
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
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 操作日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationLogSet []*OperationLog `json:"OperationLogSet,omitnil" name:"OperationLogSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// Pod名称
	Name *string `json:"Name,omitnil" name:"Name"`
}

type DescribeEdgePodRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// Pod名称
	Name *string `json:"Name,omitnil" name:"Name"`
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
	Pod *EdgeNodePodInfo `json:"Pod,omitnil" name:"Pod"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 根据节点名称模糊匹配
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`

	// 根据设备SN模糊匹配
	SNPattern *string `json:"SNPattern,omitnil" name:"SNPattern"`

	// 根据备注批次信息模糊匹配
	RemarkPattern *string `json:"RemarkPattern,omitnil" name:"RemarkPattern"`

	// 默认0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 默认20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeEdgeSnNodesRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 根据节点名称模糊匹配
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`

	// 根据设备SN模糊匹配
	SNPattern *string `json:"SNPattern,omitnil" name:"SNPattern"`

	// 根据备注批次信息模糊匹配
	RemarkPattern *string `json:"RemarkPattern,omitnil" name:"RemarkPattern"`

	// 默认0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 默认20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 节点详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeSet []*EdgeDracoNodeInfo `json:"NodeSet,omitnil" name:"NodeSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeEdgeUnitApplicationEventsRequestParams struct {
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`
}

type DescribeEdgeUnitApplicationEventsRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`
}

func (r *DescribeEdgeUnitApplicationEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitApplicationEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitApplicationEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitApplicationEventsResponseParams struct {
	// 事件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventSet []*Event `json:"EventSet,omitnil" name:"EventSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEdgeUnitApplicationEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitApplicationEventsResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitApplicationEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitApplicationEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitApplicationLogsRequestParams struct {
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 最大条数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Pod名
	PodName *string `json:"PodName,omitnil" name:"PodName"`

	// 容器名
	ContainerName *string `json:"ContainerName,omitnil" name:"ContainerName"`
}

type DescribeEdgeUnitApplicationLogsRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 最大条数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Pod名
	PodName *string `json:"PodName,omitnil" name:"PodName"`

	// 容器名
	ContainerName *string `json:"ContainerName,omitnil" name:"ContainerName"`
}

func (r *DescribeEdgeUnitApplicationLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitApplicationLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "ApplicationId")
	delete(f, "Limit")
	delete(f, "PodName")
	delete(f, "ContainerName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitApplicationLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitApplicationLogsResponseParams struct {
	// 日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogSet []*string `json:"LogSet,omitnil" name:"LogSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEdgeUnitApplicationLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitApplicationLogsResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitApplicationLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitApplicationLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitApplicationPodContainersRequestParams struct {
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// Pod名
	PodName *string `json:"PodName,omitnil" name:"PodName"`
}

type DescribeEdgeUnitApplicationPodContainersRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// Pod名
	PodName *string `json:"PodName,omitnil" name:"PodName"`
}

func (r *DescribeEdgeUnitApplicationPodContainersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitApplicationPodContainersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "ApplicationId")
	delete(f, "PodName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitApplicationPodContainersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitApplicationPodContainersResponseParams struct {
	// 容器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerSet []*ContainerStatus `json:"ContainerSet,omitnil" name:"ContainerSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEdgeUnitApplicationPodContainersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitApplicationPodContainersResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitApplicationPodContainersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitApplicationPodContainersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitApplicationPodsRequestParams struct {
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *int64 `json:"ApplicationId,omitnil" name:"ApplicationId"`
}

type DescribeEdgeUnitApplicationPodsRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *int64 `json:"ApplicationId,omitnil" name:"ApplicationId"`
}

func (r *DescribeEdgeUnitApplicationPodsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitApplicationPodsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitApplicationPodsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitApplicationPodsResponseParams struct {
	// Pod列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodSet []*PodStatus `json:"PodSet,omitnil" name:"PodSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEdgeUnitApplicationPodsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitApplicationPodsResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitApplicationPodsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitApplicationPodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitApplicationVisualizationRequestParams struct {
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`
}

type DescribeEdgeUnitApplicationVisualizationRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`
}

func (r *DescribeEdgeUnitApplicationVisualizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitApplicationVisualizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitApplicationVisualizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitApplicationVisualizationResponseParams struct {
	// 基本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitnil" name:"BasicInfo"`

	// 基本配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	BasicConfig *ApplicationBasicConfig `json:"BasicConfig,omitnil" name:"BasicConfig"`

	// 卷配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Volumes []*Volume `json:"Volumes,omitnil" name:"Volumes"`

	// 初始化容器配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitContainers []*Container `json:"InitContainers,omitnil" name:"InitContainers"`

	// 容器配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Containers []*Container `json:"Containers,omitnil" name:"Containers"`

	// 服务配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Service *Service `json:"Service,omitnil" name:"Service"`

	// Job配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Job *Job `json:"Job,omitnil" name:"Job"`

	// CronJob配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CronJob *CronJob `json:"CronJob,omitnil" name:"CronJob"`

	// 重启策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartPolicy *string `json:"RestartPolicy,omitnil" name:"RestartPolicy"`

	// HPA
	// 注意：此字段可能返回 null，表示取不到有效值。
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil" name:"HorizontalPodAutoscaler"`

	// 镜像拉取Secret
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImagePullSecrets []*string `json:"ImagePullSecrets,omitnil" name:"ImagePullSecrets"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEdgeUnitApplicationVisualizationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitApplicationVisualizationResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitApplicationVisualizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitApplicationVisualizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitApplicationYamlErrorRequestParams struct {
	// Yaml配置
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`
}

type DescribeEdgeUnitApplicationYamlErrorRequest struct {
	*tchttp.BaseRequest
	
	// Yaml配置
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`
}

func (r *DescribeEdgeUnitApplicationYamlErrorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitApplicationYamlErrorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Yaml")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitApplicationYamlErrorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitApplicationYamlErrorResponseParams struct {
	// 是否通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckPass *bool `json:"CheckPass,omitnil" name:"CheckPass"`

	// 错误类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrType *int64 `json:"ErrType,omitnil" name:"ErrType"`

	// 错误信息
	ErrInfo *string `json:"ErrInfo,omitnil" name:"ErrInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEdgeUnitApplicationYamlErrorResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitApplicationYamlErrorResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitApplicationYamlErrorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitApplicationYamlErrorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitApplicationYamlRequestParams struct {
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`
}

type DescribeEdgeUnitApplicationYamlRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`
}

func (r *DescribeEdgeUnitApplicationYamlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitApplicationYamlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitApplicationYamlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitApplicationYamlResponseParams struct {
	// Yaml配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEdgeUnitApplicationYamlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitApplicationYamlResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitApplicationYamlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitApplicationYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitApplicationsRequestParams struct {
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 翻页偏移
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 名称模糊匹配
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`

	// 字段排序 (Sort.Filed为:StartTime）
	Sort []*FieldSort `json:"Sort,omitnil" name:"Sort"`

	// 命名空间过滤
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
}

type DescribeEdgeUnitApplicationsRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 翻页偏移
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 名称模糊匹配
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`

	// 字段排序 (Sort.Filed为:StartTime）
	Sort []*FieldSort `json:"Sort,omitnil" name:"Sort"`

	// 命名空间过滤
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 应用列表
	ApplicationSet []*ApplicationStatusInfo `json:"ApplicationSet,omitnil" name:"ApplicationSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeEdgeUnitCloudRequestParams struct {
	// 边缘集群ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`
}

type DescribeEdgeUnitCloudRequest struct {
	*tchttp.BaseRequest
	
	// 边缘集群ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`
}

func (r *DescribeEdgeUnitCloudRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitCloudRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeUnitId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeUnitCloudRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitCloudResponseParams struct {
	// 边缘集群名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 集群最后探活时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveTime *string `json:"LiveTime,omitnil" name:"LiveTime"`

	// 集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterStatus *string `json:"MasterStatus,omitnil" name:"MasterStatus"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	K8sVersion *string `json:"K8sVersion,omitnil" name:"K8sVersion"`

	// pod cidr
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodCIDR *string `json:"PodCIDR,omitnil" name:"PodCIDR"`

	// service cidr
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceCIDR *string `json:"ServiceCIDR,omitnil" name:"ServiceCIDR"`

	// 集群内网访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	APIServerAddress *string `json:"APIServerAddress,omitnil" name:"APIServerAddress"`

	// 集群外网访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	APIServerExposeAddress *string `json:"APIServerExposeAddress,omitnil" name:"APIServerExposeAddress"`

	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UID *string `json:"UID,omitnil" name:"UID"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitID *uint64 `json:"UnitID,omitnil" name:"UnitID"`

	// 集群标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cluster *string `json:"Cluster,omitnil" name:"Cluster"`

	// 节点统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	Node *EdgeUnitStatisticItem `json:"Node,omitnil" name:"Node"`

	// 工作负载统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	Workload *EdgeUnitStatisticItem `json:"Workload,omitnil" name:"Workload"`

	// Grid应用统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grid *EdgeUnitStatisticItem `json:"Grid,omitnil" name:"Grid"`

	// 设备统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubDevice *EdgeUnitStatisticItem `json:"SubDevice,omitnil" name:"SubDevice"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEdgeUnitCloudResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeUnitCloudResponseParams `json:"Response"`
}

func (r *DescribeEdgeUnitCloudResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeUnitCloudResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeUnitDeployGridItemRequestParams struct {
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// Grid名称
	GridName *string `json:"GridName,omitnil" name:"GridName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil" name:"WorkloadKind"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 排序，默认ASC
	Order *string `json:"Order,omitnil" name:"Order"`
}

type DescribeEdgeUnitDeployGridItemRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// Grid名称
	GridName *string `json:"GridName,omitnil" name:"GridName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil" name:"WorkloadKind"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 排序，默认ASC
	Order *string `json:"Order,omitnil" name:"Order"`
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
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Grid部署列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeploySet []*GridItemInfo `json:"DeploySet,omitnil" name:"DeploySet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil" name:"WorkloadKind"`

	// Grid部署项名称
	GridItemName *string `json:"GridItemName,omitnil" name:"GridItemName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
}

type DescribeEdgeUnitDeployGridItemYamlRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil" name:"WorkloadKind"`

	// Grid部署项名称
	GridItemName *string `json:"GridItemName,omitnil" name:"GridItemName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
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
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`

	// 对应类型的副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas []*int64 `json:"Replicas,omitnil" name:"Replicas"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 模糊匹配
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`

	// 分页offset，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页limit，默认为20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 排序，默认为ASC
	Order *string `json:"Order,omitnil" name:"Order"`
}

type DescribeEdgeUnitDeployGridRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 模糊匹配
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`

	// 分页offset，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页limit，默认为20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 排序，默认为ASC
	Order *string `json:"Order,omitnil" name:"Order"`
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
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Grid列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	GridSet []*GridInfo `json:"GridSet,omitnil" name:"GridSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`
}

type DescribeEdgeUnitExtraRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`
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
	APIServerType *string `json:"APIServerType,omitnil" name:"APIServerType"`

	// 域名URL
	APIServerURL *string `json:"APIServerURL,omitnil" name:"APIServerURL"`

	// 域名URL对应的端口
	APIServerURLPort *string `json:"APIServerURLPort,omitnil" name:"APIServerURLPort"`

	// 域名URL对应的端口
	APIServerResolveIP *string `json:"APIServerResolveIP,omitnil" name:"APIServerResolveIP"`

	// 对外可访问的IP
	APIServerExposeAddress *string `json:"APIServerExposeAddress,omitnil" name:"APIServerExposeAddress"`

	// 是否开启监控
	IsCreatePrometheus *bool `json:"IsCreatePrometheus,omitnil" name:"IsCreatePrometheus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// Grid名称
	GridName *string `json:"GridName,omitnil" name:"GridName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil" name:"WorkloadKind"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// NodeUnit名称
	NodeUnit *string `json:"NodeUnit,omitnil" name:"NodeUnit"`

	// Pod名称
	PodName *string `json:"PodName,omitnil" name:"PodName"`
}

type DescribeEdgeUnitGridEventsRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// Grid名称
	GridName *string `json:"GridName,omitnil" name:"GridName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil" name:"WorkloadKind"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// NodeUnit名称
	NodeUnit *string `json:"NodeUnit,omitnil" name:"NodeUnit"`

	// Pod名称
	PodName *string `json:"PodName,omitnil" name:"PodName"`
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
	EventSet []*GridEventInfo `json:"EventSet,omitnil" name:"EventSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// Grid名称
	GridName *string `json:"GridName,omitnil" name:"GridName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil" name:"WorkloadKind"`

	// NodeUnit名
	NodeUnit *string `json:"NodeUnit,omitnil" name:"NodeUnit"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
}

type DescribeEdgeUnitGridPodsRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// Grid名称
	GridName *string `json:"GridName,omitnil" name:"GridName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil" name:"WorkloadKind"`

	// NodeUnit名
	NodeUnit *string `json:"NodeUnit,omitnil" name:"NodeUnit"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
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
	PodSet []*GridPodInfo `json:"PodSet,omitnil" name:"PodSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`
}

type DescribeEdgeUnitMonitorStatusRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`
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
	MonitorStatus *string `json:"MonitorStatus,omitnil" name:"MonitorStatus"`

	// 监控是否就绪
	IsAvailable *bool `json:"IsAvailable,omitnil" name:"IsAvailable"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 分页offset，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页limit，默认为20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 模糊匹配参数，精确匹配时失效
	NameFilter *string `json:"NameFilter,omitnil" name:"NameFilter"`

	// 精确匹配参数
	NameMatched *string `json:"NameMatched,omitnil" name:"NameMatched"`

	// 按时间排序，ASC/DESC，默认为DESC
	Order *string `json:"Order,omitnil" name:"Order"`
}

type DescribeEdgeUnitNodeGroupRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 分页offset，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页limit，默认为20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 模糊匹配参数，精确匹配时失效
	NameFilter *string `json:"NameFilter,omitnil" name:"NameFilter"`

	// 精确匹配参数
	NameMatched *string `json:"NameMatched,omitnil" name:"NameMatched"`

	// 按时间排序，ASC/DESC，默认为DESC
	Order *string `json:"Order,omitnil" name:"Order"`
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
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// NodeGroup数组
	NodeGroupInfo []*NodeGroupInfo `json:"NodeGroupInfo,omitnil" name:"NodeGroupInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 分页查询offset，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询limit，默认为20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模糊匹配，精确匹配时失效
	NameFilter *string `json:"NameFilter,omitnil" name:"NameFilter"`

	// 精确匹配
	NameMatched *string `json:"NameMatched,omitnil" name:"NameMatched"`

	// 按时间排序顺序，默认为DESC
	Order *string `json:"Order,omitnil" name:"Order"`
}

type DescribeEdgeUnitNodeUnitTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 分页查询offset，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询limit，默认为20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 模糊匹配，精确匹配时失效
	NameFilter *string `json:"NameFilter,omitnil" name:"NameFilter"`

	// 精确匹配
	NameMatched *string `json:"NameMatched,omitnil" name:"NameMatched"`

	// 按时间排序顺序，默认为DESC
	Order *string `json:"Order,omitnil" name:"Order"`
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
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// NodeUnit模板列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeUnitTemplates []*NodeUnitTemplate `json:"NodeUnitTemplates,omitnil" name:"NodeUnitTemplates"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// limit值
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 集群名称模糊匹配
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`

	// 排序，ASC/DESC(默认)
	Order *string `json:"Order,omitnil" name:"Order"`
}

type DescribeEdgeUnitsCloudRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// limit值
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 集群名称模糊匹配
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`

	// 排序，ASC/DESC(默认)
	Order *string `json:"Order,omitnil" name:"Order"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 集群详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	EdgeUnitSet []*EdgeCloudCluster `json:"EdgeUnitSet,omitnil" name:"EdgeUnitSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	DeviceId *int64 `json:"DeviceId,omitnil" name:"DeviceId"`

	// 无
	ProductID *string `json:"ProductID,omitnil" name:"ProductID"`

	// 无
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`
}

type DescribeIotDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备id，传0值表示此参数无效
	DeviceId *int64 `json:"DeviceId,omitnil" name:"DeviceId"`

	// 无
	ProductID *string `json:"ProductID,omitnil" name:"ProductID"`

	// 无
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`
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
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 设备名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 版本号
	Version *string `json:"Version,omitnil" name:"Version"`

	// ssl证书
	Cert *string `json:"Cert,omitnil" name:"Cert"`

	// ssl私钥
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`

	// psk认证密钥
	Psk *string `json:"Psk,omitnil" name:"Psk"`

	// 设备是否打开
	Disabled *bool `json:"Disabled,omitnil" name:"Disabled"`

	// 设备日志
	LogSetting *int64 `json:"LogSetting,omitnil" name:"LogSetting"`

	// 设备日志级别
	LogLevel *int64 `json:"LogLevel,omitnil" name:"LogLevel"`

	// mqtt参数
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// mqtt参数
	Password *string `json:"Password,omitnil" name:"Password"`

	// mqtt参数
	ClientID *string `json:"ClientID,omitnil" name:"ClientID"`

	// 16进制的psk格式
	PskHex *string `json:"PskHex,omitnil" name:"PskHex"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 设备在线状态
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 无
	Region *string `json:"Region,omitnil" name:"Region"`

	// 无
	UnitID *int64 `json:"UnitID,omitnil" name:"UnitID"`

	// 无
	UnitName *string `json:"UnitName,omitnil" name:"UnitName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 每页数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 产品id
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// 设备名称模糊查找
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`

	// 版本列表
	Versions []*string `json:"Versions,omitnil" name:"Versions"`

	// ASC 或 DESC
	Order *string `json:"Order,omitnil" name:"Order"`
}

type DescribeIotDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 页偏移
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 每页数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 产品id
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// 设备名称模糊查找
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`

	// 版本列表
	Versions []*string `json:"Versions,omitnil" name:"Versions"`

	// ASC 或 DESC
	Order *string `json:"Order,omitnil" name:"Order"`
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
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 设备列表
	DeviceSet []*IotDevicesInfo `json:"DeviceSet,omitnil" name:"DeviceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 无
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 无
	Filter *string `json:"Filter,omitnil" name:"Filter"`

	// 无
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 无
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 无
	Order *string `json:"Order,omitnil" name:"Order"`
}

type DescribeMessageRouteListRequest struct {
	*tchttp.BaseRequest
	
	// 无
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 无
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 无
	Filter *string `json:"Filter,omitnil" name:"Filter"`

	// 无
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 无
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 无
	Order *string `json:"Order,omitnil" name:"Order"`
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
	RouteList []*RouteInfo `json:"RouteList,omitnil" name:"RouteList"`

	// 无
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 查询维度
	QueryType *string `json:"QueryType,omitnil" name:"QueryType"`

	// 起始时间Unix秒时间戳
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 终止时间Unix秒时间戳
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 步长（分钟）
	Interval *int64 `json:"Interval,omitnil" name:"Interval"`

	// 节点名称，查询节点监控时必填
	NodeName *string `json:"NodeName,omitnil" name:"NodeName"`

	// 命名空间，不填则默认为default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// Pod名称，查询Pod监控时必填
	PodName *string `json:"PodName,omitnil" name:"PodName"`

	// Workload名称，查询Workload监控时必填
	WorkloadName *string `json:"WorkloadName,omitnil" name:"WorkloadName"`
}

type DescribeMonitorMetricsRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 查询维度
	QueryType *string `json:"QueryType,omitnil" name:"QueryType"`

	// 起始时间Unix秒时间戳
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 终止时间Unix秒时间戳
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 步长（分钟）
	Interval *int64 `json:"Interval,omitnil" name:"Interval"`

	// 节点名称，查询节点监控时必填
	NodeName *string `json:"NodeName,omitnil" name:"NodeName"`

	// 命名空间，不填则默认为default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// Pod名称，查询Pod监控时必填
	PodName *string `json:"PodName,omitnil" name:"PodName"`

	// Workload名称，查询Workload监控时必填
	WorkloadName *string `json:"WorkloadName,omitnil" name:"WorkloadName"`
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
	Metrics []*MonitorMetricsColumn `json:"Metrics,omitnil" name:"Metrics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// 命名空间名
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
}

type DescribeNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// 命名空间名
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
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
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
}

type DescribeNamespaceResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
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
	Resources []*NamespaceResource `json:"Resources,omitnil" name:"Resources"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 状态 (Active|Terminating)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 是否保护-不允许删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protected *bool `json:"Protected,omitnil" name:"Protected"`

	// Yaml文件格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// 边缘节点名称模糊搜索串
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`
}

type DescribeNamespacesRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// 边缘节点名称模糊搜索串
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`
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
	Items []*NamespaceInfo `json:"Items,omitnil" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// NodeUnit所属的NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitnil" name:"NodeGroupName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 分页查询limit，默认20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页查询offset，默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 模糊匹配
	NameFilter *string `json:"NameFilter,omitnil" name:"NameFilter"`
}

type DescribeNodeUnitRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// NodeUnit所属的NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitnil" name:"NodeGroupName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 分页查询limit，默认20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页查询offset，默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 模糊匹配
	NameFilter *string `json:"NameFilter,omitnil" name:"NameFilter"`
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
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// NodeUnit信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeGridInfo []*NodeUnitInfo `json:"NodeGridInfo,omitnil" name:"NodeGridInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitnil" name:"NodeGroupName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 名称模糊匹配
	NodeUnitNamePattern *string `json:"NodeUnitNamePattern,omitnil" name:"NodeUnitNamePattern"`

	// 分页查询offset，默认0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询limit，默认20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序，默认DESC
	Order *string `json:"Order,omitnil" name:"Order"`
}

type DescribeNodeUnitTemplateOnNodeGroupRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitnil" name:"NodeGroupName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 名称模糊匹配
	NodeUnitNamePattern *string `json:"NodeUnitNamePattern,omitnil" name:"NodeUnitNamePattern"`

	// 分页查询offset，默认0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询limit，默认20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序，默认DESC
	Order *string `json:"Order,omitnil" name:"Order"`
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
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// NodeUnit模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeUnitTemplates []*NodeGroupNodeUnitTemplateInfo `json:"NodeUnitTemplates,omitnil" name:"NodeUnitTemplates"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// secret名
	SecretName *string `json:"SecretName,omitnil" name:"SecretName"`

	// 命名空间(默认值:default）
	SecretNamespace *string `json:"SecretNamespace,omitnil" name:"SecretNamespace"`
}

type DescribeSecretRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// secret名
	SecretName *string `json:"SecretName,omitnil" name:"SecretName"`

	// 命名空间(默认值:default）
	SecretNamespace *string `json:"SecretNamespace,omitnil" name:"SecretNamespace"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// secret的yaml格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`

	// secret的json格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Json *string `json:"Json,omitnil" name:"Json"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`
}

type DescribeSecretYamlErrorRequest struct {
	*tchttp.BaseRequest
	
	// yaml文件
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`
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
	CheckPass *bool `json:"CheckPass,omitnil" name:"CheckPass"`

	// 错误类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrType *uint64 `json:"ErrType,omitnil" name:"ErrType"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrInfo *string `json:"ErrInfo,omitnil" name:"ErrInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// 页号
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页数目
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 命名空间
	SecretNamespace *string `json:"SecretNamespace,omitnil" name:"SecretNamespace"`

	// Secret名(模糊匹配)
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`

	// Sort.Field:CreateTime Sort.Order:ASC|DESC
	Sort *FieldSort `json:"Sort,omitnil" name:"Sort"`

	// Secret类型(DockerConfigJson或Opaque)
	SecretType *string `json:"SecretType,omitnil" name:"SecretType"`
}

type DescribeSecretsRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// 页号
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页数目
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 命名空间
	SecretNamespace *string `json:"SecretNamespace,omitnil" name:"SecretNamespace"`

	// Secret名(模糊匹配)
	NamePattern *string `json:"NamePattern,omitnil" name:"NamePattern"`

	// Sort.Field:CreateTime Sort.Order:ASC|DESC
	Sort *FieldSort `json:"Sort,omitnil" name:"Sort"`

	// Secret类型(DockerConfigJson或Opaque)
	SecretType *string `json:"SecretType,omitnil" name:"SecretType"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Secret列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*SecretItem `json:"Items,omitnil" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 允许创建的节点数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateNodeLimit *uint64 `json:"CreateNodeLimit,omitnil" name:"CreateNodeLimit"`

	// 允许创建的集群数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateClusterLimit *uint64 `json:"CreateClusterLimit,omitnil" name:"CreateClusterLimit"`

	// 是否有监控开启权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnablePermMonitor *bool `json:"EnablePermMonitor,omitnil" name:"EnablePermMonitor"`

	// 节点是否有admin的所有权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnablePermAdminNode *bool `json:"EnablePermAdminNode,omitnil" name:"EnablePermAdminNode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RegistryDomain *string `json:"RegistryDomain,omitnil" name:"RegistryDomain"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 密码
	Password *string `json:"Password,omitnil" name:"Password"`
}

type DracoNodeInfo struct {
	// 设备SN。SN仅支持大写字母、数字，长度限制为1~32个字符
	SN *string `json:"SN,omitnil" name:"SN"`

	// 节点名称。长度限制为1~63个字符，节点名称只支持小写英文、数字、中横线、英文句号
	Name *string `json:"Name,omitnil" name:"Name"`

	// 节点备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type EdgeCloudCluster struct {
	// IECP侧边缘集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EdgeId *uint64 `json:"EdgeId,omitnil" name:"EdgeId"`

	// 边缘集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// 集群版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	K8SVersion *string `json:"K8SVersion,omitnil" name:"K8SVersion"`

	// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterDesc *string `json:"ClusterDesc,omitnil" name:"ClusterDesc"`

	// 集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// pod cidr
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodCIDR *string `json:"PodCIDR,omitnil" name:"PodCIDR"`

	// service cidr
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceCIDR *string `json:"ServiceCIDR,omitnil" name:"ServiceCIDR"`

	// 边缘版本类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EdgeClusterVersion *string `json:"EdgeClusterVersion,omitnil" name:"EdgeClusterVersion"`

	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UID *string `json:"UID,omitnil" name:"UID"`
}

type EdgeDracoNodeInfo struct {
	// 节点ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 节点名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 是否已激活
	IsUsed *bool `json:"IsUsed,omitnil" name:"IsUsed"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 备注信息，如批次
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// SN 设备号
	SN *string `json:"SN,omitnil" name:"SN"`
}

type EdgeNodeInfo struct {
	// IECP边缘节点ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 节点名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 节点状态 （1健康｜2异常｜3离线｜4未激活）
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 节点资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *EdgeNodeResourceInfo `json:"Resource,omitnil" name:"Resource"`

	// CPU体系结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuArchitecture *string `json:"CpuArchitecture,omitnil" name:"CpuArchitecture"`

	// IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 操作系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatingSystem *string `json:"OperatingSystem,omitnil" name:"OperatingSystem"`

	// 节点所属的NodeUnit
	// key：NodeUnit模版ID，Value：NodeUnit模版名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeUnits *KeyValueObj `json:"NodeUnits,omitnil" name:"NodeUnits"`
}

type EdgeNodeInstallerOnline struct {
	// 节点安装脚本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptName *string `json:"ScriptName,omitnil" name:"ScriptName"`

	// 节点安装脚本下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptDownloadUrl *string `json:"ScriptDownloadUrl,omitnil" name:"ScriptDownloadUrl"`

	// 节点安装命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	Guide *string `json:"Guide,omitnil" name:"Guide"`
}

type EdgeNodeLabel struct {
	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 是否受保护
	Protected *bool `json:"Protected,omitnil" name:"Protected"`
}

type EdgeNodePodContainerInfo struct {
	// Pod名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 容器ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil" name:"Id"`

	// 镜像（含版本号）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Image *string `json:"Image,omitnil" name:"Image"`

	// CPU Request
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuRequest *string `json:"CpuRequest,omitnil" name:"CpuRequest"`

	// CPU Limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuLimit *string `json:"CpuLimit,omitnil" name:"CpuLimit"`

	// Memory Request
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryRequest *string `json:"MemoryRequest,omitnil" name:"MemoryRequest"`

	// Memory Limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryLimit *string `json:"MemoryLimit,omitnil" name:"MemoryLimit"`

	// 重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *uint64 `json:"RestartCount,omitnil" name:"RestartCount"`

	// 容器状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`
}

type EdgeNodePodInfo struct {
	// Pod名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// Pod状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 所在节点IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeIp *string `json:"NodeIp,omitnil" name:"NodeIp"`

	// 实例IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// CPU Request
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuRequest *string `json:"CpuRequest,omitnil" name:"CpuRequest"`

	// Memory Request
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryRequest *string `json:"MemoryRequest,omitnil" name:"MemoryRequest"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 工作负载类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkloadType *string `json:"WorkloadType,omitnil" name:"WorkloadType"`

	// 工作负载名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkloadName *string `json:"WorkloadName,omitnil" name:"WorkloadName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *uint64 `json:"RestartCount,omitnil" name:"RestartCount"`

	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil" name:"ClusterID"`
}

type EdgeNodeResourceInfo struct {
	// 可使用的CPU 单位: m核
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllocatedCPU *string `json:"AllocatedCPU,omitnil" name:"AllocatedCPU"`

	// CPU总量 单位:m核
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCPU *string `json:"TotalCPU,omitnil" name:"TotalCPU"`

	// 已分配的内存 单位G
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllocatedMemory *string `json:"AllocatedMemory,omitnil" name:"AllocatedMemory"`

	// 内存总量 单位G
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalMemory *string `json:"TotalMemory,omitnil" name:"TotalMemory"`

	// 已分配的GPU资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllocatedGPU *string `json:"AllocatedGPU,omitnil" name:"AllocatedGPU"`

	// GPU总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalGPU *string `json:"TotalGPU,omitnil" name:"TotalGPU"`

	// 可使用的CPU 单位: m核
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableCPU *string `json:"AvailableCPU,omitnil" name:"AvailableCPU"`

	// 可使用的内存 单位: G
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableMemory *string `json:"AvailableMemory,omitnil" name:"AvailableMemory"`

	// 可使用的GPU资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableGPU *string `json:"AvailableGPU,omitnil" name:"AvailableGPU"`
}

type EdgeUnitStatisticItem struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 在线数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Online *uint64 `json:"Online,omitnil" name:"Online"`

	// 异常数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Abnormal *uint64 `json:"Abnormal,omitnil" name:"Abnormal"`

	// 离线数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offline *uint64 `json:"Offline,omitnil" name:"Offline"`

	// 未激活
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotActive *uint64 `json:"NotActive,omitnil" name:"NotActive"`
}

type Env struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 值
	Value *string `json:"Value,omitnil" name:"Value"`

	// 值引用
	ValueFrom *EnvValueSelector `json:"ValueFrom,omitnil" name:"ValueFrom"`
}

type EnvValueSelector struct {
	// 健名
	Key *string `json:"Key,omitnil" name:"Key"`

	// 对象名
	ObjectName *string `json:"ObjectName,omitnil" name:"ObjectName"`

	// 对象值
	ObjectType *string `json:"ObjectType,omitnil" name:"ObjectType"`
}

type Event struct {
	// 第一次出现时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 最后一次出现时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTime *string `json:"LastTime,omitnil" name:"LastTime"`

	// 事件关联对象类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvolvedObjectKind *string `json:"InvolvedObjectKind,omitnil" name:"InvolvedObjectKind"`

	// 事件关联对象名
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvolvedObjectName *string `json:"InvolvedObjectName,omitnil" name:"InvolvedObjectName"`

	// 事件类型(Normal|Warning)
	Type *string `json:"Type,omitnil" name:"Type"`

	// 原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// 内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 出现次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitnil" name:"Count"`
}

type FieldSort struct {
	// 字段名
	Field *string `json:"Field,omitnil" name:"Field"`

	// 排序(ASC:升序 DESC:降序
	Order *string `json:"Order,omitnil" name:"Order"`
}

// Predefined struct for user
type GetMarketComponentListRequestParams struct {
	// 页偏移，从0开始
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 名称模糊筛选
	Filter *string `json:"Filter,omitnil" name:"Filter"`

	// 以名称排序，ASC、DESC
	Order *string `json:"Order,omitnil" name:"Order"`
}

type GetMarketComponentListRequest struct {
	*tchttp.BaseRequest
	
	// 页偏移，从0开始
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 名称模糊筛选
	Filter *string `json:"Filter,omitnil" name:"Filter"`

	// 以名称排序，ASC、DESC
	Order *string `json:"Order,omitnil" name:"Order"`
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
	ComponentList []*MarketComponentInfo `json:"ComponentList,omitnil" name:"ComponentList"`

	// 组件总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ID *int64 `json:"ID,omitnil" name:"ID"`
}

type GetMarketComponentRequest struct {
	*tchttp.BaseRequest
	
	// 组件ID
	ID *int64 `json:"ID,omitnil" name:"ID"`
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
	ID *int64 `json:"ID,omitnil" name:"ID"`

	// 组件名称
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// 发行组织
	Author *string `json:"Author,omitnil" name:"Author"`

	// 发布时间
	ReleaseTime *string `json:"ReleaseTime,omitnil" name:"ReleaseTime"`

	// 组件简介
	Outline *string `json:"Outline,omitnil" name:"Outline"`

	// 详细介绍链接
	Detail *string `json:"Detail,omitnil" name:"Detail"`

	// 图标连接
	Icon *string `json:"Icon,omitnil" name:"Icon"`

	// 组件版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 组件可视化配置
	WorkloadVisualConfig *string `json:"WorkloadVisualConfig,omitnil" name:"WorkloadVisualConfig"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// GridID
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

type GridEventInfo struct {
	// 首次出现时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 最后出现时间
	LastTime *string `json:"LastTime,omitnil" name:"LastTime"`

	// 对象类型
	InvolvedObjectKind *string `json:"InvolvedObjectKind,omitnil" name:"InvolvedObjectKind"`

	// 对象名称
	InvolvedObjectName *string `json:"InvolvedObjectName,omitnil" name:"InvolvedObjectName"`

	// 事件类型(Normal,Warning)
	Type *string `json:"Type,omitnil" name:"Type"`

	// 事件原因
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// 事件内容
	Message *string `json:"Message,omitnil" name:"Message"`

	// 次数
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 节点名（Pod事件类型时有值）
	NodeName *string `json:"NodeName,omitnil" name:"NodeName"`

	// 节点内部IP（Pod事件类型时有值）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IP *string `json:"IP,omitnil" name:"IP"`
}

type GridInfo struct {
	// DeployGridId
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// Key
	GridUniqKey *string `json:"GridUniqKey,omitnil" name:"GridUniqKey"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 工作负载类型
	WorkloadKind *string `json:"WorkloadKind,omitnil" name:"WorkloadKind"`

	// 启动时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas *int64 `json:"Replicas,omitnil" name:"Replicas"`

	// 创建人
	Publisher *string `json:"Publisher,omitnil" name:"Publisher"`

	// 版本信息
	Version *string `json:"Version,omitnil" name:"Version"`
}

type GridItemInfo struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 期望副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas *int64 `json:"Replicas,omitnil" name:"Replicas"`

	// 可用副本数
	AvailableReplicas *int64 `json:"AvailableReplicas,omitnil" name:"AvailableReplicas"`

	// 启动时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 工作负载类型
	WorkloadKind *string `json:"WorkloadKind,omitnil" name:"WorkloadKind"`
}

type GridPodInfo struct {
	// Pod名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命名空间
	NameSpace *string `json:"NameSpace,omitnil" name:"NameSpace"`

	// 状态(Pending｜Running｜Succeeded｜Failed｜Unknown)
	Status *string `json:"Status,omitnil" name:"Status"`

	// 节点名
	NodeName *string `json:"NodeName,omitnil" name:"NodeName"`

	// 节点IP
	NodeIP *string `json:"NodeIP,omitnil" name:"NodeIP"`

	// Pod的IP
	PodIP *string `json:"PodIP,omitnil" name:"PodIP"`

	// 启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 运行时长（秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunSec *int64 `json:"RunSec,omitnil" name:"RunSec"`

	// 重启次数
	RestartCount *int64 `json:"RestartCount,omitnil" name:"RestartCount"`

	// 集群名称ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterID *string `json:"ClusterID,omitnil" name:"ClusterID"`
}

type HorizontalPodAutoscaler struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 最小实例数
	MinReplicas *int64 `json:"MinReplicas,omitnil" name:"MinReplicas"`

	// 最大实例数
	MaxReplicas *int64 `json:"MaxReplicas,omitnil" name:"MaxReplicas"`

	// 资源目标指标
	ResourceMetricTarget []*ResourceMetricTarget `json:"ResourceMetricTarget,omitnil" name:"ResourceMetricTarget"`
}

type HttpHeader struct {
	// HTTP头的名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// HTTP头的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type HttpProbe struct {
	// 请求路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil" name:"Path"`

	// 请求端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// 请求地址，默认Pod的IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitnil" name:"Host"`

	// 请求模式  HTTP|HTTPS，默认HTTP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scheme *string `json:"Scheme,omitnil" name:"Scheme"`

	// HTTP的请求头
	// 注意：此字段可能返回 null，表示取不到有效值。
	Headers []*HttpHeader `json:"Headers,omitnil" name:"Headers"`
}

type IotDevicesInfo struct {
	// 设备id
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 设备名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 设备状态
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 设备打开状态
	Disabled *bool `json:"Disabled,omitnil" name:"Disabled"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 设备创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 最后在线时间
	LastOnlineTime *string `json:"LastOnlineTime,omitnil" name:"LastOnlineTime"`

	// 设备是否绑定到节点
	IsBound *bool `json:"IsBound,omitnil" name:"IsBound"`

	// 设备版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 无
	Region *string `json:"Region,omitnil" name:"Region"`

	// 无
	UnitID *int64 `json:"UnitID,omitnil" name:"UnitID"`

	// 无
	UnitName *string `json:"UnitName,omitnil" name:"UnitName"`
}

type Job struct {
	// 并发数
	Parallelism *int64 `json:"Parallelism,omitnil" name:"Parallelism"`

	// 完成数
	Completion *int64 `json:"Completion,omitnil" name:"Completion"`

	// 最大运行时间
	ActiveDeadlineSeconds *int64 `json:"ActiveDeadlineSeconds,omitnil" name:"ActiveDeadlineSeconds"`

	// 失败前重试次数
	BackOffLimit *int64 `json:"BackOffLimit,omitnil" name:"BackOffLimit"`
}

type KeyValueObj struct {
	// Key值
	Key *string `json:"Key,omitnil" name:"Key"`

	// Value值
	Value *string `json:"Value,omitnil" name:"Value"`
}

type Label struct {
	// 健名
	Key *string `json:"Key,omitnil" name:"Key"`

	// 健值
	Value *string `json:"Value,omitnil" name:"Value"`
}

type MarketComponentInfo struct {
	// 组件ID
	ID *int64 `json:"ID,omitnil" name:"ID"`

	// 组件名称
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// 发布者
	Author *string `json:"Author,omitnil" name:"Author"`

	// 发布时间
	ReleaseTime *string `json:"ReleaseTime,omitnil" name:"ReleaseTime"`

	// 组件简介
	Outline *string `json:"Outline,omitnil" name:"Outline"`

	// 指向详细描述的url
	Detail *string `json:"Detail,omitnil" name:"Detail"`

	// 图标链接
	Icon *string `json:"Icon,omitnil" name:"Icon"`

	// 组件版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 组件可视化信息
	WorkloadVisualConfig *string `json:"WorkloadVisualConfig,omitnil" name:"WorkloadVisualConfig"`

	// 无
	DetailUrl *string `json:"DetailUrl,omitnil" name:"DetailUrl"`

	// 无
	Installed *bool `json:"Installed,omitnil" name:"Installed"`
}

// Predefined struct for user
type ModifyApplicationBasicInfoRequestParams struct {
	// 应用模板ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 应用模板基本信息
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitnil" name:"BasicInfo"`
}

type ModifyApplicationBasicInfoRequest struct {
	*tchttp.BaseRequest
	
	// 应用模板ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 应用模板基本信息
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitnil" name:"BasicInfo"`
}

func (r *ModifyApplicationBasicInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationBasicInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "BasicInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationBasicInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationBasicInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyApplicationBasicInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationBasicInfoResponseParams `json:"Response"`
}

func (r *ModifyApplicationBasicInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationBasicInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationVisualizationRequestParams struct {
	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 应用配置
	BasicConfig *ApplicationBasicConfig `json:"BasicConfig,omitnil" name:"BasicConfig"`

	// 卷配置
	Volumes []*Volume `json:"Volumes,omitnil" name:"Volumes"`

	// 初始容器
	InitContainers []*Container `json:"InitContainers,omitnil" name:"InitContainers"`

	// 容器配置
	Containers []*Container `json:"Containers,omitnil" name:"Containers"`

	// 服务配置
	Service *Service `json:"Service,omitnil" name:"Service"`

	// Job配置
	Job *Job `json:"Job,omitnil" name:"Job"`

	// CronJob配置
	CronJob *CronJob `json:"CronJob,omitnil" name:"CronJob"`

	// 重启策略
	RestartPolicy *string `json:"RestartPolicy,omitnil" name:"RestartPolicy"`

	// 镜像拉取密钥
	ImagePullSecrets []*string `json:"ImagePullSecrets,omitnil" name:"ImagePullSecrets"`

	// HPA配置
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil" name:"HorizontalPodAutoscaler"`

	// 单个初始化容器
	InitContainer *Container `json:"InitContainer,omitnil" name:"InitContainer"`
}

type ModifyApplicationVisualizationRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 应用配置
	BasicConfig *ApplicationBasicConfig `json:"BasicConfig,omitnil" name:"BasicConfig"`

	// 卷配置
	Volumes []*Volume `json:"Volumes,omitnil" name:"Volumes"`

	// 初始容器
	InitContainers []*Container `json:"InitContainers,omitnil" name:"InitContainers"`

	// 容器配置
	Containers []*Container `json:"Containers,omitnil" name:"Containers"`

	// 服务配置
	Service *Service `json:"Service,omitnil" name:"Service"`

	// Job配置
	Job *Job `json:"Job,omitnil" name:"Job"`

	// CronJob配置
	CronJob *CronJob `json:"CronJob,omitnil" name:"CronJob"`

	// 重启策略
	RestartPolicy *string `json:"RestartPolicy,omitnil" name:"RestartPolicy"`

	// 镜像拉取密钥
	ImagePullSecrets []*string `json:"ImagePullSecrets,omitnil" name:"ImagePullSecrets"`

	// HPA配置
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil" name:"HorizontalPodAutoscaler"`

	// 单个初始化容器
	InitContainer *Container `json:"InitContainer,omitnil" name:"InitContainer"`
}

func (r *ModifyApplicationVisualizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationVisualizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
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
	delete(f, "InitContainer")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationVisualizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationVisualizationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyApplicationVisualizationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationVisualizationResponseParams `json:"Response"`
}

func (r *ModifyApplicationVisualizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationVisualizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConfigMapRequestParams struct {
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// ConfigMap名称
	ConfigMapName *string `json:"ConfigMapName,omitnil" name:"ConfigMapName"`

	// Yaml配置, base64之后的串
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`

	// ConfigMap命名空间
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil" name:"ConfigMapNamespace"`
}

type ModifyConfigMapRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// ConfigMap名称
	ConfigMapName *string `json:"ConfigMapName,omitnil" name:"ConfigMapName"`

	// Yaml配置, base64之后的串
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`

	// ConfigMap命名空间
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitnil" name:"ConfigMapNamespace"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 边缘节点ID
	NodeId *uint64 `json:"NodeId,omitnil" name:"NodeId"`

	// 节点信息
	NodeInfo *DracoNodeInfo `json:"NodeInfo,omitnil" name:"NodeInfo"`

	// 是否重置draco设备
	IsReset *bool `json:"IsReset,omitnil" name:"IsReset"`
}

type ModifyEdgeDracoNodeRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 边缘节点ID
	NodeId *uint64 `json:"NodeId,omitnil" name:"NodeId"`

	// 节点信息
	NodeInfo *DracoNodeInfo `json:"NodeInfo,omitnil" name:"NodeInfo"`

	// 是否重置draco设备
	IsReset *bool `json:"IsReset,omitnil" name:"IsReset"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// IECP边缘节点ID
	NodeId *uint64 `json:"NodeId,omitnil" name:"NodeId"`

	// 标签列表
	Labels []*KeyValueObj `json:"Labels,omitnil" name:"Labels"`
}

type ModifyEdgeNodeLabelsRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// IECP边缘节点ID
	NodeId *uint64 `json:"NodeId,omitnil" name:"NodeId"`

	// 标签列表
	Labels []*KeyValueObj `json:"Labels,omitnil" name:"Labels"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitnil" name:"BasicInfo"`

	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`
}

type ModifyEdgeUnitApplicationBasicInfoRequest struct {
	*tchttp.BaseRequest
	
	// 应用基本信息
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitnil" name:"BasicInfo"`

	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 应用配置
	BasicConfig *ApplicationBasicConfig `json:"BasicConfig,omitnil" name:"BasicConfig"`

	// 卷配置
	Volumes []*Volume `json:"Volumes,omitnil" name:"Volumes"`

	// 初始容器列表
	InitContainers []*Container `json:"InitContainers,omitnil" name:"InitContainers"`

	// 容器配置
	Containers []*Container `json:"Containers,omitnil" name:"Containers"`

	// 服务配置
	Service *Service `json:"Service,omitnil" name:"Service"`

	// Job配置
	Job *Job `json:"Job,omitnil" name:"Job"`

	// CronJob配置
	CronJob *CronJob `json:"CronJob,omitnil" name:"CronJob"`

	// 重启策略
	RestartPolicy *string `json:"RestartPolicy,omitnil" name:"RestartPolicy"`

	// 镜像拉取密钥
	ImagePullSecrets []*string `json:"ImagePullSecrets,omitnil" name:"ImagePullSecrets"`

	// HPA配置
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil" name:"HorizontalPodAutoscaler"`
}

type ModifyEdgeUnitApplicationVisualizationRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 应用配置
	BasicConfig *ApplicationBasicConfig `json:"BasicConfig,omitnil" name:"BasicConfig"`

	// 卷配置
	Volumes []*Volume `json:"Volumes,omitnil" name:"Volumes"`

	// 初始容器列表
	InitContainers []*Container `json:"InitContainers,omitnil" name:"InitContainers"`

	// 容器配置
	Containers []*Container `json:"Containers,omitnil" name:"Containers"`

	// 服务配置
	Service *Service `json:"Service,omitnil" name:"Service"`

	// Job配置
	Job *Job `json:"Job,omitnil" name:"Job"`

	// CronJob配置
	CronJob *CronJob `json:"CronJob,omitnil" name:"CronJob"`

	// 重启策略
	RestartPolicy *string `json:"RestartPolicy,omitnil" name:"RestartPolicy"`

	// 镜像拉取密钥
	ImagePullSecrets []*string `json:"ImagePullSecrets,omitnil" name:"ImagePullSecrets"`

	// HPA配置
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil" name:"HorizontalPodAutoscaler"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// Yaml配置
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`
}

type ModifyEdgeUnitApplicationYamlRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// Yaml配置
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 边缘单元名称，64字符内
	Name *string `json:"Name,omitnil" name:"Name"`

	// 描述，200字符内
	Description *string `json:"Description,omitnil" name:"Description"`

	// 是否开启监控
	OpenCloudMonitor *bool `json:"OpenCloudMonitor,omitnil" name:"OpenCloudMonitor"`
}

type ModifyEdgeUnitCloudApiRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 边缘单元名称，64字符内
	Name *string `json:"Name,omitnil" name:"Name"`

	// 描述，200字符内
	Description *string `json:"Description,omitnil" name:"Description"`

	// 是否开启监控
	OpenCloudMonitor *bool `json:"OpenCloudMonitor,omitnil" name:"OpenCloudMonitor"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// Grid名称
	GridItemName *string `json:"GridItemName,omitnil" name:"GridItemName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil" name:"WorkloadKind"`

	// 副本数
	Replicas *int64 `json:"Replicas,omitnil" name:"Replicas"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
}

type ModifyEdgeUnitDeployGridItemRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// Grid名称
	GridItemName *string `json:"GridItemName,omitnil" name:"GridItemName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitnil" name:"WorkloadKind"`

	// 副本数
	Replicas *int64 `json:"Replicas,omitnil" name:"Replicas"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 边缘集群名称，64字符以内
	Name *string `json:"Name,omitnil" name:"Name"`

	// 集群描述，200字符以内
	Description *string `json:"Description,omitnil" name:"Description"`
}

type ModifyEdgeUnitRequest struct {
	*tchttp.BaseRequest
	
	// 边缘集群ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 边缘集群名称，64字符以内
	Name *string `json:"Name,omitnil" name:"Name"`

	// 集群描述，200字符以内
	Description *string `json:"Description,omitnil" name:"Description"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	DeviceId *int64 `json:"DeviceId,omitnil" name:"DeviceId"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 设备是否开启
	Disabled *bool `json:"Disabled,omitnil" name:"Disabled"`

	// 日志设置
	LogSetting *int64 `json:"LogSetting,omitnil" name:"LogSetting"`

	// 日志级别
	LogLevel *int64 `json:"LogLevel,omitnil" name:"LogLevel"`
}

type ModifyIotDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备id
	DeviceId *int64 `json:"DeviceId,omitnil" name:"DeviceId"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 设备是否开启
	Disabled *bool `json:"Disabled,omitnil" name:"Disabled"`

	// 日志设置
	LogSetting *int64 `json:"LogSetting,omitnil" name:"LogSetting"`

	// 日志级别
	LogLevel *int64 `json:"LogLevel,omitnil" name:"LogLevel"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// NodeUnit模板ID
	NodeUnitTemplateID *uint64 `json:"NodeUnitTemplateID,omitnil" name:"NodeUnitTemplateID"`

	// 包含的节点列表
	Nodes []*string `json:"Nodes,omitnil" name:"Nodes"`
}

type ModifyNodeUnitTemplateRequest struct {
	*tchttp.BaseRequest
	
	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// NodeUnit模板ID
	NodeUnitTemplateID *uint64 `json:"NodeUnitTemplateID,omitnil" name:"NodeUnitTemplateID"`

	// 包含的节点列表
	Nodes []*string `json:"Nodes,omitnil" name:"Nodes"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// Secret名
	SecretName *string `json:"SecretName,omitnil" name:"SecretName"`

	// Secret的Yaml格式
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`

	// Secret命名空间（默认:default）
	SecretNamespace *string `json:"SecretNamespace,omitnil" name:"SecretNamespace"`
}

type ModifySecretRequest struct {
	*tchttp.BaseRequest
	
	// 边缘单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitnil" name:"EdgeUnitID"`

	// Secret名
	SecretName *string `json:"SecretName,omitnil" name:"SecretName"`

	// Secret的Yaml格式
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`

	// Secret命名空间（默认:default）
	SecretNamespace *string `json:"SecretNamespace,omitnil" name:"SecretNamespace"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ColumnName *string `json:"ColumnName,omitnil" name:"ColumnName"`

	// 数据内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnData []*string `json:"ColumnData,omitnil" name:"ColumnData"`

	// 数据所属，查询Workload类型时有值
	ColumnBelong *string `json:"ColumnBelong,omitnil" name:"ColumnBelong"`

	// 最大值
	MaxValue *float64 `json:"MaxValue,omitnil" name:"MaxValue"`

	// 最小值
	MinValue *float64 `json:"MinValue,omitnil" name:"MinValue"`

	// 平均值
	AvgValue *float64 `json:"AvgValue,omitnil" name:"AvgValue"`

	// 时间戳数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnTime *int64 `json:"ColumnTime,omitnil" name:"ColumnTime"`
}

type NamespaceInfo struct {
	// 命名空间名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 状态(Active|Terminating)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 是否保护(不允许删除)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protected *bool `json:"Protected,omitnil" name:"Protected"`

	// 对应的Yaml配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Yaml *string `json:"Yaml,omitnil" name:"Yaml"`
}

type NamespaceResource struct {
	// 类型(workload|grid|configmap|secret)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 名称(最多返回5个）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Names []*string `json:"Names,omitnil" name:"Names"`
}

type NodeGroupInfo struct {
	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitnil" name:"NodeGroupName"`

	// DeploymentGrid数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeploymentGridList []*GridDetail `json:"DeploymentGridList,omitnil" name:"DeploymentGridList"`

	// StatefulSetGrid数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatefulSetGridList []*GridDetail `json:"StatefulSetGridList,omitnil" name:"StatefulSetGridList"`

	// 是否平台保护
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protect *bool `json:"Protect,omitnil" name:"Protect"`
}

type NodeGroupNodeUnitTemplateInfo struct {
	// 模版ID
	ID *uint64 `json:"ID,omitnil" name:"ID"`

	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 包含节点列表
	NodeList []*NodeSimpleInfo `json:"NodeList,omitnil" name:"NodeList"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 是否关联
	Relation *bool `json:"Relation,omitnil" name:"Relation"`
}

type NodeSimpleInfo struct {
	// 节点ID
	ID *uint64 `json:"ID,omitnil" name:"ID"`

	// 节点名称
	NodeName *string `json:"NodeName,omitnil" name:"NodeName"`
}

type NodeUnitInfo struct {
	// NodeUnitId
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// NodeUnit名称
	NodeUnitName *string `json:"NodeUnitName,omitnil" name:"NodeUnitName"`

	// 包含节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeList []*NodeUnitNodeInfo `json:"NodeList,omitnil" name:"NodeList"`
}

type NodeUnitNodeInfo struct {
	// 节点ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 节点状态  NodeStatusHealthy (健康)/NodeStatusAbnormal (异常)/NodeStatusOffline (下线)/NodeStatusNotActivated (未激活
	Status *string `json:"Status,omitnil" name:"Status"`

	// 节点名称
	NodeName *string `json:"NodeName,omitnil" name:"NodeName"`

	// 内网节点IP
	InternalIP *string `json:"InternalIP,omitnil" name:"InternalIP"`
}

type NodeUnitTemplate struct {
	// NodeUnit模版ID
	ID *uint64 `json:"ID,omitnil" name:"ID"`

	// NodeUnit模版名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 包含节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeList []*NodeSimpleInfo `json:"NodeList,omitnil" name:"NodeList"`

	// NodeGroup列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeGroups []*string `json:"NodeGroups,omitnil" name:"NodeGroups"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`
}

type OperationLog struct {
	// 操作时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateTime *string `json:"OperateTime,omitnil" name:"OperateTime"`

	// 模块名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Module *string `json:"Module,omitnil" name:"Module"`

	// 操作信息
	Description *string `json:"Description,omitnil" name:"Description"`

	// 用户ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 状态: 1:成功 2:失败
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 操作用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUserID *string `json:"OperatorUserID,omitnil" name:"OperatorUserID"`

	// 操作动作
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitnil" name:"Action"`
}

type OperationLogsCondition struct {
	// 状态列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status []*int64 `json:"Status,omitnil" name:"Status"`
}

type PodStatus struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameSpace *string `json:"NameSpace,omitnil" name:"NameSpace"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	IP *string `json:"IP,omitnil" name:"IP"`

	// 启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunSec *int64 `json:"RunSec,omitnil" name:"RunSec"`

	// 重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *int64 `json:"RestartCount,omitnil" name:"RestartCount"`
}

type PortConfig struct {
	// 协议类型(tcp|udp)
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 源端口
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// 目标端口
	TargetPort *int64 `json:"TargetPort,omitnil" name:"TargetPort"`

	// 节点端口
	NodePort *int64 `json:"NodePort,omitnil" name:"NodePort"`
}

type Probe struct {
	// 启动后，延迟探测时间 单位:秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitialDelaySeconds *int64 `json:"InitialDelaySeconds,omitnil" name:"InitialDelaySeconds"`

	// 探测间隔，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeriodSeconds *int64 `json:"PeriodSeconds,omitnil" name:"PeriodSeconds"`

	// 探测超时时间 单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeoutSeconds *int64 `json:"TimeoutSeconds,omitnil" name:"TimeoutSeconds"`

	// 失败后检查成功的最小连续成功次数。默认为1.活跃度必须为1。最小值为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessThreshold *int64 `json:"SuccessThreshold,omitnil" name:"SuccessThreshold"`

	// 当Pod成功启动且检查失败时，放弃之前尝试次数。默认为3.最小值为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureThreshold *int64 `json:"FailureThreshold,omitnil" name:"FailureThreshold"`

	// HTTP探测配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpProbe *HttpProbe `json:"HttpProbe,omitnil" name:"HttpProbe"`

	// TCP探测配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TcpProbe *TcpProbe `json:"TcpProbe,omitnil" name:"TcpProbe"`
}

// Predefined struct for user
type RedeployEdgeUnitApplicationRequestParams struct {
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`
}

type RedeployEdgeUnitApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitnil" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil" name:"ApplicationId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Type *string `json:"Type,omitnil" name:"Type"`

	// 平均值
	AverageValue *int64 `json:"AverageValue,omitnil" name:"AverageValue"`

	// 单位
	Scale *string `json:"Scale,omitnil" name:"Scale"`

	// 平均值
	AverageUtilization *int64 `json:"AverageUtilization,omitnil" name:"AverageUtilization"`
}

type RouteInfo struct {
	// 无
	RouteID *int64 `json:"RouteID,omitnil" name:"RouteID"`

	// 无
	RouteName *string `json:"RouteName,omitnil" name:"RouteName"`

	// 无
	SourceProductID *string `json:"SourceProductID,omitnil" name:"SourceProductID"`

	// 无
	TopicFilter *string `json:"TopicFilter,omitnil" name:"TopicFilter"`

	// 无
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 无
	TargetOptions *string `json:"TargetOptions,omitnil" name:"TargetOptions"`

	// 无
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 无
	Descript *string `json:"Descript,omitnil" name:"Descript"`

	// 无
	Healthy *string `json:"Healthy,omitnil" name:"Healthy"`

	// 无
	Status *string `json:"Status,omitnil" name:"Status"`

	// 无
	MessageCount *int64 `json:"MessageCount,omitnil" name:"MessageCount"`

	// 无
	MessageLastTime *string `json:"MessageLastTime,omitnil" name:"MessageLastTime"`

	// 无
	SourceProductName *string `json:"SourceProductName,omitnil" name:"SourceProductName"`

	// 无
	SourceUnitIDList []*string `json:"SourceUnitIDList,omitnil" name:"SourceUnitIDList"`

	// 无
	SourceUnitNameList []*string `json:"SourceUnitNameList,omitnil" name:"SourceUnitNameList"`

	// 无
	SourceDeviceNameList []*string `json:"SourceDeviceNameList,omitnil" name:"SourceDeviceNameList"`
}

type SecretItem struct {
	// Secret名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Secret类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretType *string `json:"SecretType,omitnil" name:"SecretType"`
}

type SecurityCapabilities struct {
	// 允许操作列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Add []*string `json:"Add,omitnil" name:"Add"`

	// 禁止操作列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Drop []*string `json:"Drop,omitnil" name:"Drop"`
}

type SecurityContext struct {
	// 是否开启特权模式
	Privilege *bool `json:"Privilege,omitnil" name:"Privilege"`

	// 目录/Proc挂载方式
	ProcMount *string `json:"ProcMount,omitnil" name:"ProcMount"`

	// 安全配置
	Capabilities *SecurityCapabilities `json:"Capabilities,omitnil" name:"Capabilities"`
}

type Service struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 类型 (ClusterIP|NodePort)
	Type *string `json:"Type,omitnil" name:"Type"`

	// 端口配置
	Ports []*PortConfig `json:"Ports,omitnil" name:"Ports"`

	// 标签
	Labels []*Label `json:"Labels,omitnil" name:"Labels"`

	// 命名空间默认default
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 服务IP
	ClusterIP *string `json:"ClusterIP,omitnil" name:"ClusterIP"`
}

// Predefined struct for user
type SetRouteOnOffRequestParams struct {
	// 无
	RouteID *int64 `json:"RouteID,omitnil" name:"RouteID"`

	// on 或 off
	Status *string `json:"Status,omitnil" name:"Status"`
}

type SetRouteOnOffRequest struct {
	*tchttp.BaseRequest
	
	// 无
	RouteID *int64 `json:"RouteID,omitnil" name:"RouteID"`

	// on 或 off
	Status *string `json:"Status,omitnil" name:"Status"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Field *string `json:"Field,omitnil" name:"Field"`

	// 排序方式，升序ASC / 降序DESC
	Order *string `json:"Order,omitnil" name:"Order"`
}

type TcpProbe struct {
	// 连接端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil" name:"Port"`
}

type Volume struct {
	// 来源(emptyDir|hostPath|configMap|secret|nfs)
	Source *string `json:"Source,omitnil" name:"Source"`

	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// Host挂载配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostPath *VolumeHostPath `json:"HostPath,omitnil" name:"HostPath"`

	// ConfigMap挂载配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigMap *VolumeConfigMap `json:"ConfigMap,omitnil" name:"ConfigMap"`

	// Secret挂载配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Secret *VolumeConfigMap `json:"Secret,omitnil" name:"Secret"`

	// NFS挂载配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	NFS *VolumeNFS `json:"NFS,omitnil" name:"NFS"`
}

type VolumeConfigMap struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// Key列表配置
	Items []*VolumeConfigMapKeyToPath `json:"Items,omitnil" name:"Items"`
}

type VolumeConfigMapKeyToPath struct {
	// 健名
	Key *string `json:"Key,omitnil" name:"Key"`

	// 对应本地路径
	Path *string `json:"Path,omitnil" name:"Path"`

	// 对应权限模式
	Mode *string `json:"Mode,omitnil" name:"Mode"`
}

type VolumeHostPath struct {
	// 类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 路径
	Path *string `json:"Path,omitnil" name:"Path"`
}

type VolumeMount struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 挂载路径
	MountPath *string `json:"MountPath,omitnil" name:"MountPath"`

	// 子路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubPath *string `json:"SubPath,omitnil" name:"SubPath"`

	// 是否只读
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadOnly *bool `json:"ReadOnly,omitnil" name:"ReadOnly"`
}

type VolumeNFS struct {
	// 服务地址
	Server *string `json:"Server,omitnil" name:"Server"`

	// 对应服务器路径
	ServerPath *string `json:"ServerPath,omitnil" name:"ServerPath"`

	// 对应本地路径
	Path *string `json:"Path,omitnil" name:"Path"`
}