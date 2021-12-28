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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ApplicationBasicConfig struct {

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 工作负载类型
	WorkflowKind *string `json:"WorkflowKind,omitempty" name:"WorkflowKind"`

	// 标签信息
	Labels []*Label `json:"Labels,omitempty" name:"Labels"`

	// Grid唯一Key
	GridUniqKey *string `json:"GridUniqKey,omitempty" name:"GridUniqKey"`

	// NodeSelector标签
	NodeSelector []*Label `json:"NodeSelector,omitempty" name:"NodeSelector"`

	// 实例数
	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`

	// 可用实例数
	AvailableReplicas *int64 `json:"AvailableReplicas,omitempty" name:"AvailableReplicas"`

	// 是否开启service环境变量注入pod
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableServiceLinks *bool `json:"EnableServiceLinks,omitempty" name:"EnableServiceLinks"`
}

type ApplicationBasicInfo struct {

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 管理URL地址
	ManageUrl *string `json:"ManageUrl,omitempty" name:"ManageUrl"`

	// 描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ApplicationDeployMode struct {

	// 1:指定节点部署 2:单元部署
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 资源ID
	ResourceID *uint64 `json:"ResourceID,omitempty" name:"ResourceID"`

	// 资源名
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

type ApplicationStatusInfo struct {

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 应用版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`

	// 应用状态(1:待部署 2:部署中 3:运行中 4:待更新 5:更新中 6:待删除 7:删除中 8:已删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 管理地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManageUrl *string `json:"ManageUrl,omitempty" name:"ManageUrl"`

	// 负载类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkloadKind *string `json:"WorkloadKind,omitempty" name:"WorkloadKind"`

	// 应用部署模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployMode *ApplicationDeployMode `json:"DeployMode,omitempty" name:"DeployMode"`

	// 期望Pod数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`

	// 运行Pod数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableReplicas *int64 `json:"AvailableReplicas,omitempty" name:"AvailableReplicas"`
}

type ApplicationTemplate struct {

	// 模板ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 来源。1 自定义应用模板 ;  2 官方应用模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *uint64 `json:"Source,omitempty" name:"Source"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkloadKind *string `json:"WorkloadKind,omitempty" name:"WorkloadKind"`

	// 管理地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManageUrl *string `json:"ManageUrl,omitempty" name:"ManageUrl"`

	// 发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DistributeTime *string `json:"DistributeTime,omitempty" name:"DistributeTime"`
}

type ApplyMarketComponentRequest struct {
	*tchttp.BaseRequest

	// 组件ID
	ID *int64 `json:"ID,omitempty" name:"ID"`
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

type ApplyMarketComponentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ConfigMapBasicInfo struct {

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type Container struct {

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 镜像名
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像版本
	ImageVersion *string `json:"ImageVersion,omitempty" name:"ImageVersion"`

	// 镜像拉取策略(Always|Never|IfNotPresent)
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" name:"ImagePullPolicy"`

	// 卷挂载配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeMounts []*VolumeMount `json:"VolumeMounts,omitempty" name:"VolumeMounts"`

	// cpu最低配置
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// cpu最高限制
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 内存最低要求
	MemoryRequest *string `json:"MemoryRequest,omitempty" name:"MemoryRequest"`

	// 内存最高要求
	MemoryLimit *string `json:"MemoryLimit,omitempty" name:"MemoryLimit"`

	// 内存单位
	MemoryUnit *string `json:"MemoryUnit,omitempty" name:"MemoryUnit"`

	// gpu最高限制
	GpuLimit *string `json:"GpuLimit,omitempty" name:"GpuLimit"`

	// 资源配置
	ResourceMapCloud []*KeyValueObj `json:"ResourceMapCloud,omitempty" name:"ResourceMapCloud"`

	// 环境配置
	Envs []*Env `json:"Envs,omitempty" name:"Envs"`

	// 工作目录
	WorkingDir *string `json:"WorkingDir,omitempty" name:"WorkingDir"`

	// 命令
	Commands []*string `json:"Commands,omitempty" name:"Commands"`

	// 参数
	Args []*string `json:"Args,omitempty" name:"Args"`

	// 安全配置
	SecurityContext *SecurityContext `json:"SecurityContext,omitempty" name:"SecurityContext"`

	// 就绪探针配置
	ReadinessProbe *Probe `json:"ReadinessProbe,omitempty" name:"ReadinessProbe"`
}

type ContainerStatus struct {

	// 容器名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 容器ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *string `json:"ID,omitempty" name:"ID"`

	// 镜像
	// 注意：此字段可能返回 null，表示取不到有效值。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *int64 `json:"RestartCount,omitempty" name:"RestartCount"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`
}

type CreateApplicationVisualizationRequest struct {
	*tchttp.BaseRequest

	// 基本信息
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitempty" name:"BasicInfo"`

	// 基本配置
	BasicConfig *ApplicationBasicConfig `json:"BasicConfig,omitempty" name:"BasicConfig"`

	// 卷列表
	Volumes []*Volume `json:"Volumes,omitempty" name:"Volumes"`

	// 服务配置
	Service *Service `json:"Service,omitempty" name:"Service"`

	// Job配置
	Job *Job `json:"Job,omitempty" name:"Job"`

	// CronJob配置
	CronJob *CronJob `json:"CronJob,omitempty" name:"CronJob"`

	// 重新运行策略
	RestartPolicy *string `json:"RestartPolicy,omitempty" name:"RestartPolicy"`

	// 镜像拉取密钥
	ImagePullSecrets []*string `json:"ImagePullSecrets,omitempty" name:"ImagePullSecrets"`

	// HPA配置
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitempty" name:"HorizontalPodAutoscaler"`

	// 初始化容器列表
	InitContainers []*Container `json:"InitContainers,omitempty" name:"InitContainers"`

	// 容器列表
	Containers []*Container `json:"Containers,omitempty" name:"Containers"`
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

type CreateApplicationVisualizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		ApplicationId *uint64 `json:"ApplicationId,omitempty" name:"ApplicationId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateConfigMapRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitempty" name:"EdgeUnitID"`

	// ConfigMap名称
	ConfigMapName *string `json:"ConfigMapName,omitempty" name:"ConfigMapName"`

	// ConfigMap内容
	ConfigMapData []*KeyValueObj `json:"ConfigMapData,omitempty" name:"ConfigMapData"`

	// ConfigMap命名空间,默认：default
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitempty" name:"ConfigMapNamespace"`
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

type CreateConfigMapResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateEdgeNodeGroupRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// NodeGroup名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命名空间，不填默认为default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 模版ID数组
	NodeUnitTemplateIDs []*uint64 `json:"NodeUnitTemplateIDs,omitempty" name:"NodeUnitTemplateIDs"`
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

type CreateEdgeNodeGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateEdgeNodeRequest struct {
	*tchttp.BaseRequest

	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 节点名称
	Name *string `json:"Name,omitempty" name:"Name"`
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

type CreateEdgeNodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateEdgeNodeUnitTemplateRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// NodeUnit模版名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 包含的节点列表
	Nodes []*string `json:"Nodes,omitempty" name:"Nodes"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`
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

type CreateEdgeNodeUnitTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateEdgeUnitApplicationVisualizationRequest struct {
	*tchttp.BaseRequest

	// 基本信息
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitempty" name:"BasicInfo"`

	// 基本配置
	BasicConfig *ApplicationBasicConfig `json:"BasicConfig,omitempty" name:"BasicConfig"`

	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 卷列表
	Volumes []*Volume `json:"Volumes,omitempty" name:"Volumes"`

	// 服务配置
	Service *Service `json:"Service,omitempty" name:"Service"`

	// 模版ID
	TemplateID *uint64 `json:"TemplateID,omitempty" name:"TemplateID"`

	// Job配置
	Job *Job `json:"Job,omitempty" name:"Job"`

	// CronJob配置
	CronJob *CronJob `json:"CronJob,omitempty" name:"CronJob"`

	// 重新运行策略
	RestartPolicy *string `json:"RestartPolicy,omitempty" name:"RestartPolicy"`

	// 镜像拉取密钥
	ImagePullSecrets []*string `json:"ImagePullSecrets,omitempty" name:"ImagePullSecrets"`

	// HPA配置
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitempty" name:"HorizontalPodAutoscaler"`

	// 初始化容器列表
	InitContainers []*Container `json:"InitContainers,omitempty" name:"InitContainers"`

	// 容器列表
	Containers []*Container `json:"Containers,omitempty" name:"Containers"`
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

type CreateEdgeUnitApplicationVisualizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		ApplicationId *uint64 `json:"ApplicationId,omitempty" name:"ApplicationId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateEdgeUnitApplicationYamlRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitId *int64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 基本信息
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitempty" name:"BasicInfo"`

	// Yaml配置
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
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
	delete(f, "BasicInfo")
	delete(f, "Yaml")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEdgeUnitApplicationYamlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateEdgeUnitApplicationYamlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		ApplicationId *uint64 `json:"ApplicationId,omitempty" name:"ApplicationId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateEdgeUnitCloudRequest struct {
	*tchttp.BaseRequest

	// 集群名称，长度小于32
	Name *string `json:"Name,omitempty" name:"Name"`

	// k8s版本，仅支持1.16.7 和 1.18.2
	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 集群描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 集群pod cidr， 默认  10.1.0.0/16
	PodCIDR *string `json:"PodCIDR,omitempty" name:"PodCIDR"`

	// 集群service cidr, 默认 10.2.0.0/16
	ServiceCIDR *string `json:"ServiceCIDR,omitempty" name:"ServiceCIDR"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEdgeUnitCloudRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateEdgeUnitCloudResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// tke集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// IECP集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateNamespaceRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitempty" name:"EdgeUnitID"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 描述信息
	Description *string `json:"Description,omitempty" name:"Description"`
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

type CreateNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateSecretRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitempty" name:"EdgeUnitID"`

	// secret名
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 命名空间（默认:default）
	SecretNamespace *string `json:"SecretNamespace,omitempty" name:"SecretNamespace"`

	// secret类型(取值范围:DockerConfigJson,Opaque 默认Opaque)
	SecretType *string `json:"SecretType,omitempty" name:"SecretType"`

	// DockerConfig的序列化base64编码后的字符串
	DockerConfigJson *string `json:"DockerConfigJson,omitempty" name:"DockerConfigJson"`

	// Opaque类型的Secret内容
	CloudData []*KeyValueObj `json:"CloudData,omitempty" name:"CloudData"`

	// DockerConfig配置
	DockerConfig *DockerConfig `json:"DockerConfig,omitempty" name:"DockerConfig"`
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

type CreateSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateUpdateNodeUnitRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// NodeUnit所属的NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitempty" name:"NodeGroupName"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// NodeUnit名称，通过模版创建可不填
	NodeUnitName *string `json:"NodeUnitName,omitempty" name:"NodeUnitName"`

	// NodeUnit包含的节点列表，通过模版创建可不填
	Nodes []*string `json:"Nodes,omitempty" name:"Nodes"`

	// NodeUnit模版ID列表
	NodeUnitTemplateIDs []*uint64 `json:"NodeUnitTemplateIDs,omitempty" name:"NodeUnitTemplateIDs"`
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

type CreateUpdateNodeUnitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CronJob struct {

	// 调度配置
	Schedule *string `json:"Schedule,omitempty" name:"Schedule"`

	// 运行时间
	StartingDeadlineSeconds *int64 `json:"StartingDeadlineSeconds,omitempty" name:"StartingDeadlineSeconds"`

	// job并行策略(Allow|Forbid|Replace)
	ConcurrencyPolicy *string `json:"ConcurrencyPolicy,omitempty" name:"ConcurrencyPolicy"`

	// Job配置
	Job *Job `json:"Job,omitempty" name:"Job"`
}

type DeleteApplicationsRequest struct {
	*tchttp.BaseRequest

	// 应用模板ID列表
	ApplicationIds []*uint64 `json:"ApplicationIds,omitempty" name:"ApplicationIds"`
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

type DeleteApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteConfigMapRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitempty" name:"EdgeUnitID"`

	// ConfigMap名
	ConfigMapName *string `json:"ConfigMapName,omitempty" name:"ConfigMapName"`

	// ConfigMap命名空间，默认：default
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitempty" name:"ConfigMapNamespace"`
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

type DeleteConfigMapResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteEdgeNodeGroupRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// NodeGroup名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
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

type DeleteEdgeNodeGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteEdgeNodeUnitTemplatesRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 删除的NodeUnit模版ID列表
	NodeUnitTemplateIDs []*uint64 `json:"NodeUnitTemplateIDs,omitempty" name:"NodeUnitTemplateIDs"`
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

type DeleteEdgeNodeUnitTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteEdgeNodesRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// IECP边缘节点ID列表
	NodeIds []*uint64 `json:"NodeIds,omitempty" name:"NodeIds"`
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

type DeleteEdgeNodesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteEdgeUnitApplicationsRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitempty" name:"EdgeUnitID"`

	// 应用ID列表
	ApplicationIDs []*uint64 `json:"ApplicationIDs,omitempty" name:"ApplicationIDs"`
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

type DeleteEdgeUnitApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteEdgeUnitCloudRequest struct {
	*tchttp.BaseRequest

	// 边缘集群ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`
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

type DeleteEdgeUnitCloudResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteEdgeUnitDeployGridItemRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitempty" name:"WorkloadKind"`

	// Grid部署名称
	GridItemName *string `json:"GridItemName,omitempty" name:"GridItemName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
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

type DeleteEdgeUnitDeployGridItemResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteEdgeUnitPodRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// Pod名称
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
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

type DeleteEdgeUnitPodResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteNamespaceRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitempty" name:"EdgeUnitID"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
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

type DeleteNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteNodeUnitRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// NodeUnit所属的NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitempty" name:"NodeGroupName"`

	// NodeUnit名称
	NodeUnitName *string `json:"NodeUnitName,omitempty" name:"NodeUnitName"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// NodeUnit包含的节点列表
	Nodes []*string `json:"Nodes,omitempty" name:"Nodes"`
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

type DeleteNodeUnitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteSecretRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitempty" name:"EdgeUnitID"`

	// secret名称
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// secret命名空间（默认:default）
	SecretNamespace *string `json:"SecretNamespace,omitempty" name:"SecretNamespace"`
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

type DeleteSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeApplicationVisualizationRequest struct {
	*tchttp.BaseRequest

	// 应用模板ID
	ApplicationId *uint64 `json:"ApplicationId,omitempty" name:"ApplicationId"`
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

type DescribeApplicationVisualizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 基本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitempty" name:"BasicInfo"`

		// 基本配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		BasicConfig *ApplicationBasicConfig `json:"BasicConfig,omitempty" name:"BasicConfig"`

		// 卷配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		Volumes []*Volume `json:"Volumes,omitempty" name:"Volumes"`

		// 初始化容器配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		InitContainers []*Container `json:"InitContainers,omitempty" name:"InitContainers"`

		// 容器配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		Containers []*Container `json:"Containers,omitempty" name:"Containers"`

		// 服务配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		Service *Service `json:"Service,omitempty" name:"Service"`

		// Job配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		Job *Job `json:"Job,omitempty" name:"Job"`

		// CronJob配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		CronJob *CronJob `json:"CronJob,omitempty" name:"CronJob"`

		// 重启策略
	// 注意：此字段可能返回 null，表示取不到有效值。
		RestartPolicy *string `json:"RestartPolicy,omitempty" name:"RestartPolicy"`

		// HPA
	// 注意：此字段可能返回 null，表示取不到有效值。
		HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitempty" name:"HorizontalPodAutoscaler"`

		// 镜像拉取Secret
	// 注意：此字段可能返回 null，表示取不到有效值。
		ImagePullSecrets []*string `json:"ImagePullSecrets,omitempty" name:"ImagePullSecrets"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeApplicationYamlErrorRequest struct {
	*tchttp.BaseRequest

	// Yaml配置
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
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

type DescribeApplicationYamlErrorResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否通过
	// 注意：此字段可能返回 null，表示取不到有效值。
		CheckPass *bool `json:"CheckPass,omitempty" name:"CheckPass"`

		// 错误类型
	// 注意：此字段可能返回 null，表示取不到有效值。
		ErrType *int64 `json:"ErrType,omitempty" name:"ErrType"`

		// 错误信息
		ErrInfo *string `json:"ErrInfo,omitempty" name:"ErrInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeApplicationYamlRequest struct {
	*tchttp.BaseRequest

	// 应用模板ID
	ApplicationId *uint64 `json:"ApplicationId,omitempty" name:"ApplicationId"`
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

type DescribeApplicationYamlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// base64 后的yaml
	// 注意：此字段可能返回 null，表示取不到有效值。
		Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeApplicationsRequest struct {
	*tchttp.BaseRequest

	// 模糊搜索字符串
	NamePattern *string `json:"NamePattern,omitempty" name:"NamePattern"`

	// 默认 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 默认 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 仅支持对 DistributeTime 字段排序，ASC/DESC
	Sort []*FieldSort `json:"Sort,omitempty" name:"Sort"`
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

type DescribeApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 详细列表
		ApplicationSet []*ApplicationTemplate `json:"ApplicationSet,omitempty" name:"ApplicationSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeConfigMapRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitempty" name:"EdgeUnitID"`

	// ConfigMap名称
	ConfigMapName *string `json:"ConfigMapName,omitempty" name:"ConfigMapName"`

	// ConfigMap命名空间
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitempty" name:"ConfigMapNamespace"`
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

type DescribeConfigMapResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		Name *string `json:"Name,omitempty" name:"Name"`

		// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
		Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

		// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// yaml配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

		// 配置项的json格式(base64编码)
	// 注意：此字段可能返回 null，表示取不到有效值。
		Json *string `json:"Json,omitempty" name:"Json"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeConfigMapYamlErrorRequest struct {
	*tchttp.BaseRequest

	// yaml文件
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
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

type DescribeConfigMapYamlErrorResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 校验是通过
	// 注意：此字段可能返回 null，表示取不到有效值。
		CheckPass *bool `json:"CheckPass,omitempty" name:"CheckPass"`

		// 错误类型
	// 注意：此字段可能返回 null，表示取不到有效值。
		ErrType *uint64 `json:"ErrType,omitempty" name:"ErrType"`

		// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		ErrInfo *string `json:"ErrInfo,omitempty" name:"ErrInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeConfigMapsRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitempty" name:"EdgeUnitID"`

	// 翻页偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小(最大100)
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 命名空间
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitempty" name:"ConfigMapNamespace"`

	// 模糊匹配的名称
	NamePattern *string `json:"NamePattern,omitempty" name:"NamePattern"`

	// Sort.Fileld填写CreateTime Sort.Order(ASC|DESC) 默认ASC
	Sort *FieldSort `json:"Sort,omitempty" name:"Sort"`
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

type DescribeConfigMapsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ConfigMap列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Items []*ConfigMapBasicInfo `json:"Items,omitempty" name:"Items"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeAgentNodeInstallerRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// IECP边缘节点ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`
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

type DescribeEdgeAgentNodeInstallerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 节点在线安装信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Online *EdgeNodeInstallerOnline `json:"Online,omitempty" name:"Online"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeDefaultVpcResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

		// 网络CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
		VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

		// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

		// 子网CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubnetCidrBlock *string `json:"SubnetCidrBlock,omitempty" name:"SubnetCidrBlock"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeNodePodContainersRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 节点ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

	// Pod名称
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
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

type DescribeEdgeNodePodContainersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Pod容器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		ContainerSet []*EdgeNodePodContainerInfo `json:"ContainerSet,omitempty" name:"ContainerSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeNodePodsRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 节点ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Pod名称过滤串
	PodNamePattern *string `json:"PodNamePattern,omitempty" name:"PodNamePattern"`
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

type DescribeEdgeNodePodsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Pod列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		PodSet []*EdgeNodePodInfo `json:"PodSet,omitempty" name:"PodSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeNodeRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// IECP边缘节点ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`
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

type DescribeEdgeNodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		Id *uint64 `json:"Id,omitempty" name:"Id"`

		// 节点类型
	// 注意：此字段可能返回 null，表示取不到有效值。
		Kind *string `json:"Kind,omitempty" name:"Kind"`

		// 节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		Name *string `json:"Name,omitempty" name:"Name"`

		// 节点状态 （1健康｜2异常｜3离线｜4未激活）
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// CPU体系结构
	// 注意：此字段可能返回 null，表示取不到有效值。
		CpuArchitecture *string `json:"CpuArchitecture,omitempty" name:"CpuArchitecture"`

		// AI处理器体系结构
	// 注意：此字段可能返回 null，表示取不到有效值。
		AiChipArchitecture *string `json:"AiChipArchitecture,omitempty" name:"AiChipArchitecture"`

		// IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		Ip *string `json:"Ip,omitempty" name:"Ip"`

		// 节点标签列表
		Labels []*EdgeNodeLabel `json:"Labels,omitempty" name:"Labels"`

		// 节点资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Resource *EdgeNodeResourceInfo `json:"Resource,omitempty" name:"Resource"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeNodesRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 边缘节点名称模糊搜索串
	NamePattern *string `json:"NamePattern,omitempty" name:"NamePattern"`

	// 边缘节点名称列表，支持批量查询 ，优先于模糊查询
	NameMatchedList []*string `json:"NameMatchedList,omitempty" name:"NameMatchedList"`

	// 排序信息列表
	Sort []*Sort `json:"Sort,omitempty" name:"Sort"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 页面大小Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 节点类型
	NodeType *int64 `json:"NodeType,omitempty" name:"NodeType"`
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

type DescribeEdgeNodesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 边缘节点数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		NodeSet []*EdgeNodeInfo `json:"NodeSet,omitempty" name:"NodeSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeOperationLogsRequest struct {
	*tchttp.BaseRequest

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段
	Sort []*FieldSort `json:"Sort,omitempty" name:"Sort"`

	// 模块
	Module *string `json:"Module,omitempty" name:"Module"`

	// 过滤条件
	Condition *OperationLogsCondition `json:"Condition,omitempty" name:"Condition"`
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

type DescribeEdgeOperationLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 操作日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		OperationLogSet []*OperationLog `json:"OperationLogSet,omitempty" name:"OperationLogSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgePodRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Pod名称
	Name *string `json:"Name,omitempty" name:"Name"`
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

type DescribeEdgePodResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Pod详情信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Pod *EdgeNodePodInfo `json:"Pod,omitempty" name:"Pod"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitApplicationEventsRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitempty" name:"ApplicationId"`
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

type DescribeEdgeUnitApplicationEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		EventSet []*Event `json:"EventSet,omitempty" name:"EventSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitApplicationLogsRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 最大条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pod名
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 容器名
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
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

type DescribeEdgeUnitApplicationLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		LogSet []*string `json:"LogSet,omitempty" name:"LogSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitApplicationPodContainersRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Pod名
	PodName *string `json:"PodName,omitempty" name:"PodName"`
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

type DescribeEdgeUnitApplicationPodContainersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 容器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		ContainerSet []*ContainerStatus `json:"ContainerSet,omitempty" name:"ContainerSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitApplicationPodsRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *int64 `json:"ApplicationId,omitempty" name:"ApplicationId"`
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

type DescribeEdgeUnitApplicationPodsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Pod列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		PodSet []*PodStatus `json:"PodSet,omitempty" name:"PodSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitApplicationVisualizationRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitempty" name:"ApplicationId"`
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

type DescribeEdgeUnitApplicationVisualizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 基本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitempty" name:"BasicInfo"`

		// 基本配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		BasicConfig *ApplicationBasicConfig `json:"BasicConfig,omitempty" name:"BasicConfig"`

		// 卷配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		Volumes []*Volume `json:"Volumes,omitempty" name:"Volumes"`

		// 初始化容器配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		InitContainers []*Container `json:"InitContainers,omitempty" name:"InitContainers"`

		// 容器配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		Containers []*Container `json:"Containers,omitempty" name:"Containers"`

		// 服务配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		Service *Service `json:"Service,omitempty" name:"Service"`

		// Job配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		Job *Job `json:"Job,omitempty" name:"Job"`

		// CronJob配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		CronJob *CronJob `json:"CronJob,omitempty" name:"CronJob"`

		// 重启策略
	// 注意：此字段可能返回 null，表示取不到有效值。
		RestartPolicy *string `json:"RestartPolicy,omitempty" name:"RestartPolicy"`

		// HPA
	// 注意：此字段可能返回 null，表示取不到有效值。
		HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitempty" name:"HorizontalPodAutoscaler"`

		// 镜像拉取Secret
	// 注意：此字段可能返回 null，表示取不到有效值。
		ImagePullSecrets []*string `json:"ImagePullSecrets,omitempty" name:"ImagePullSecrets"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitApplicationYamlErrorRequest struct {
	*tchttp.BaseRequest

	// Yaml配置
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
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

type DescribeEdgeUnitApplicationYamlErrorResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否通过
	// 注意：此字段可能返回 null，表示取不到有效值。
		CheckPass *bool `json:"CheckPass,omitempty" name:"CheckPass"`

		// 错误类型
	// 注意：此字段可能返回 null，表示取不到有效值。
		ErrType *int64 `json:"ErrType,omitempty" name:"ErrType"`

		// 错误信息
		ErrInfo *string `json:"ErrInfo,omitempty" name:"ErrInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitApplicationYamlRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitempty" name:"ApplicationId"`
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

type DescribeEdgeUnitApplicationYamlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Yaml配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitApplicationsRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 翻页偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 名称模糊匹配
	NamePattern *string `json:"NamePattern,omitempty" name:"NamePattern"`

	// 字段排序 (Sort.Filed为:StartTime）
	Sort []*FieldSort `json:"Sort,omitempty" name:"Sort"`

	// 命名空间过滤
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
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

type DescribeEdgeUnitApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 应用列表
		ApplicationSet []*ApplicationStatusInfo `json:"ApplicationSet,omitempty" name:"ApplicationSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitCloudRequest struct {
	*tchttp.BaseRequest

	// 边缘集群ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`
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

type DescribeEdgeUnitCloudResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 边缘集群名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// 集群最后探活时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		LiveTime *string `json:"LiveTime,omitempty" name:"LiveTime"`

		// 集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
		MasterStatus *string `json:"MasterStatus,omitempty" name:"MasterStatus"`

		// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
		K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`

		// pod cidr
	// 注意：此字段可能返回 null，表示取不到有效值。
		PodCIDR *string `json:"PodCIDR,omitempty" name:"PodCIDR"`

		// service cidr
	// 注意：此字段可能返回 null，表示取不到有效值。
		ServiceCIDR *string `json:"ServiceCIDR,omitempty" name:"ServiceCIDR"`

		// 集群内网访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		APIServerAddress *string `json:"APIServerAddress,omitempty" name:"APIServerAddress"`

		// 集群外网访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		APIServerExposeAddress *string `json:"APIServerExposeAddress,omitempty" name:"APIServerExposeAddress"`

		// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		UID *string `json:"UID,omitempty" name:"UID"`

		// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		UnitID *uint64 `json:"UnitID,omitempty" name:"UnitID"`

		// 集群标识
	// 注意：此字段可能返回 null，表示取不到有效值。
		Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

		// 节点统计
	// 注意：此字段可能返回 null，表示取不到有效值。
		Node *EdgeUnitStatisticItem `json:"Node,omitempty" name:"Node"`

		// 工作负载统计
	// 注意：此字段可能返回 null，表示取不到有效值。
		Workload *EdgeUnitStatisticItem `json:"Workload,omitempty" name:"Workload"`

		// Grid应用统计
	// 注意：此字段可能返回 null，表示取不到有效值。
		Grid *EdgeUnitStatisticItem `json:"Grid,omitempty" name:"Grid"`

		// 设备统计
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubDevice *EdgeUnitStatisticItem `json:"SubDevice,omitempty" name:"SubDevice"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitDeployGridItemRequest struct {
	*tchttp.BaseRequest

	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// Grid名称
	GridName *string `json:"GridName,omitempty" name:"GridName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitempty" name:"WorkloadKind"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 排序，默认ASC
	Order *string `json:"Order,omitempty" name:"Order"`
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

type DescribeEdgeUnitDeployGridItemResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Grid部署列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		DeploySet []*GridItemInfo `json:"DeploySet,omitempty" name:"DeploySet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitDeployGridItemYamlRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitempty" name:"WorkloadKind"`

	// Grid部署项名称
	GridItemName *string `json:"GridItemName,omitempty" name:"GridItemName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
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

type DescribeEdgeUnitDeployGridItemYamlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// yaml，base64编码字符串
		Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

		// 对应类型的副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Replicas []*int64 `json:"Replicas,omitempty" name:"Replicas"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitDeployGridRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 模糊匹配
	NamePattern *string `json:"NamePattern,omitempty" name:"NamePattern"`

	// 分页offset，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页limit，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序，默认为ASC
	Order *string `json:"Order,omitempty" name:"Order"`
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

type DescribeEdgeUnitDeployGridResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Grid列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		GridSet []*GridInfo `json:"GridSet,omitempty" name:"GridSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitExtraRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`
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

type DescribeEdgeUnitExtraResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// APIServer类型
		APIServerType *string `json:"APIServerType,omitempty" name:"APIServerType"`

		// 域名URL
		APIServerURL *string `json:"APIServerURL,omitempty" name:"APIServerURL"`

		// 域名URL对应的端口
		APIServerURLPort *string `json:"APIServerURLPort,omitempty" name:"APIServerURLPort"`

		// 域名URL对应的端口
		APIServerResolveIP *string `json:"APIServerResolveIP,omitempty" name:"APIServerResolveIP"`

		// 对外可访问的IP
		APIServerExposeAddress *string `json:"APIServerExposeAddress,omitempty" name:"APIServerExposeAddress"`

		// 是否开启监控
		IsCreatePrometheus *bool `json:"IsCreatePrometheus,omitempty" name:"IsCreatePrometheus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitGridEventsRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// Grid名称
	GridName *string `json:"GridName,omitempty" name:"GridName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitempty" name:"WorkloadKind"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// NodeUnit名称
	NodeUnit *string `json:"NodeUnit,omitempty" name:"NodeUnit"`

	// Pod名称
	PodName *string `json:"PodName,omitempty" name:"PodName"`
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

type DescribeEdgeUnitGridEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		EventSet []*GridEventInfo `json:"EventSet,omitempty" name:"EventSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitGridPodsRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// Grid名称
	GridName *string `json:"GridName,omitempty" name:"GridName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitempty" name:"WorkloadKind"`

	// NodeUnit名
	NodeUnit *string `json:"NodeUnit,omitempty" name:"NodeUnit"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
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

type DescribeEdgeUnitGridPodsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Pod列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		PodSet []*GridPodInfo `json:"PodSet,omitempty" name:"PodSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitMonitorStatusRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`
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

type DescribeEdgeUnitMonitorStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 监控状态描述：
	// "running" 单元监控正常运行
	// "deploying" 单元监控部署中
	// "norsc" 单元需要可用节点以部署监控
	// "abnormal" 单元监控异常
	// "none" 单元监控不可用
		MonitorStatus *string `json:"MonitorStatus,omitempty" name:"MonitorStatus"`

		// 监控是否就绪
		IsAvailable *bool `json:"IsAvailable,omitempty" name:"IsAvailable"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitNodeGroupRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 分页offset，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页limit，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 模糊匹配参数，精确匹配时失效
	NameFilter *string `json:"NameFilter,omitempty" name:"NameFilter"`

	// 精确匹配参数
	NameMatched *string `json:"NameMatched,omitempty" name:"NameMatched"`

	// 按时间排序，ASC/DESC，默认为DESC
	Order *string `json:"Order,omitempty" name:"Order"`
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

type DescribeEdgeUnitNodeGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// NodeGroup数组
		NodeGroupInfo []*NodeGroupInfo `json:"NodeGroupInfo,omitempty" name:"NodeGroupInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitNodeUnitTemplatesRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 分页查询offset，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询limit，默认为20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 模糊匹配，精确匹配时失效
	NameFilter *string `json:"NameFilter,omitempty" name:"NameFilter"`

	// 精确匹配
	NameMatched *string `json:"NameMatched,omitempty" name:"NameMatched"`

	// 按时间排序顺序，默认为DESC
	Order *string `json:"Order,omitempty" name:"Order"`
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

type DescribeEdgeUnitNodeUnitTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件的记录总数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// NodeUnit模版列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		NodeUnitTemplates []*NodeUnitTemplate `json:"NodeUnitTemplates,omitempty" name:"NodeUnitTemplates"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEdgeUnitsCloudRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// limit值
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 集群名称模糊匹配
	NamePattern *string `json:"NamePattern,omitempty" name:"NamePattern"`

	// 排序，ASC/DESC(默认)
	Order *string `json:"Order,omitempty" name:"Order"`
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

type DescribeEdgeUnitsCloudResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 集群详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		EdgeUnitSet []*EdgeCloudCluster `json:"EdgeUnitSet,omitempty" name:"EdgeUnitSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeMonitorMetricsRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 查询维度
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// 起始时间Unix秒时间戳
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 终止时间Unix秒时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 步长（分钟）
	Interval *int64 `json:"Interval,omitempty" name:"Interval"`

	// 节点名称，查询节点监控时必填
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// 命名空间，不填则默认为default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Pod名称，查询Pod监控时必填
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// Workload名称，查询Workload监控时必填
	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
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

type DescribeMonitorMetricsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询监控指标结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Metrics []*MonitorMetricsColumn `json:"Metrics,omitempty" name:"Metrics"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeNamespaceRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitempty" name:"EdgeUnitID"`

	// 命名空间名
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
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

type DescribeNamespaceResourcesRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitempty" name:"EdgeUnitID"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
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

type DescribeNamespaceResourcesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 资源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Resources []*NamespaceResource `json:"Resources,omitempty" name:"Resources"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 命名空间名
	// 注意：此字段可能返回 null，表示取不到有效值。
		Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

		// 状态 (Active|Terminating)
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *string `json:"Status,omitempty" name:"Status"`

		// 描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 是否保护-不允许删除
	// 注意：此字段可能返回 null，表示取不到有效值。
		Protected *bool `json:"Protected,omitempty" name:"Protected"`

		// Yaml文件格式
	// 注意：此字段可能返回 null，表示取不到有效值。
		Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeNamespacesRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitempty" name:"EdgeUnitID"`

	// 边缘节点名称模糊搜索串
	NamePattern *string `json:"NamePattern,omitempty" name:"NamePattern"`
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

type DescribeNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 命名空间信息列表
		Items []*NamespaceInfo `json:"Items,omitempty" name:"Items"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeNodeUnitRequest struct {
	*tchttp.BaseRequest

	// 边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// NodeUnit所属的NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitempty" name:"NodeGroupName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 分页查询limit，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页查询offset，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 模糊匹配
	NameFilter *string `json:"NameFilter,omitempty" name:"NameFilter"`
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

type DescribeNodeUnitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件的记录总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// NodeUnit信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		NodeGridInfo []*NodeUnitInfo `json:"NodeGridInfo,omitempty" name:"NodeGridInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeNodeUnitTemplateOnNodeGroupRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitempty" name:"NodeGroupName"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 名称模糊匹配
	NodeUnitNamePattern *string `json:"NodeUnitNamePattern,omitempty" name:"NodeUnitNamePattern"`

	// 分页查询offset，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询limit，默认20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序，默认DESC
	Order *string `json:"Order,omitempty" name:"Order"`
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

type DescribeNodeUnitTemplateOnNodeGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录总数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// NodeUnit模版
	// 注意：此字段可能返回 null，表示取不到有效值。
		NodeUnitTemplates []*NodeGroupNodeUnitTemplateInfo `json:"NodeUnitTemplates,omitempty" name:"NodeUnitTemplates"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeSecretRequest struct {
	*tchttp.BaseRequest

	// 边缘单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitempty" name:"EdgeUnitID"`

	// secret名
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 命名空间(默认值:default）
	SecretNamespace *string `json:"SecretNamespace,omitempty" name:"SecretNamespace"`
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

type DescribeSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Secret名
	// 注意：此字段可能返回 null，表示取不到有效值。
		Name *string `json:"Name,omitempty" name:"Name"`

		// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
		Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

		// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// secret的yaml格式
	// 注意：此字段可能返回 null，表示取不到有效值。
		Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

		// secret的json格式
	// 注意：此字段可能返回 null，表示取不到有效值。
		Json *string `json:"Json,omitempty" name:"Json"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeSecretYamlErrorRequest struct {
	*tchttp.BaseRequest

	// yaml文件
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
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

type DescribeSecretYamlErrorResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 校验是通过
	// 注意：此字段可能返回 null，表示取不到有效值。
		CheckPass *bool `json:"CheckPass,omitempty" name:"CheckPass"`

		// 错误类型
	// 注意：此字段可能返回 null，表示取不到有效值。
		ErrType *uint64 `json:"ErrType,omitempty" name:"ErrType"`

		// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		ErrInfo *string `json:"ErrInfo,omitempty" name:"ErrInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeSecretsRequest struct {
	*tchttp.BaseRequest

	// 边缘单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitempty" name:"EdgeUnitID"`

	// 页号
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 命名空间
	SecretNamespace *string `json:"SecretNamespace,omitempty" name:"SecretNamespace"`

	// Secret名(模糊匹配)
	NamePattern *string `json:"NamePattern,omitempty" name:"NamePattern"`

	// Sort.Field:CreateTime Sort.Order:ASC|DESC
	Sort *FieldSort `json:"Sort,omitempty" name:"Sort"`

	// Secret类型(DockerConfigJson或Opaque)
	SecretType *string `json:"SecretType,omitempty" name:"SecretType"`
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

type DescribeSecretsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Secret列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Items []*SecretItem `json:"Items,omitempty" name:"Items"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DockerConfig struct {

	// 镜像仓库地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryDomain *string `json:"RegistryDomain,omitempty" name:"RegistryDomain"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

type EdgeCloudCluster struct {

	// IECP侧边缘集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EdgeId *uint64 `json:"EdgeId,omitempty" name:"EdgeId"`

	// 边缘集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	K8SVersion *string `json:"K8SVersion,omitempty" name:"K8SVersion"`

	// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// 集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// pod cidr
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodCIDR *string `json:"PodCIDR,omitempty" name:"PodCIDR"`

	// service cidr
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceCIDR *string `json:"ServiceCIDR,omitempty" name:"ServiceCIDR"`

	// 边缘版本类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EdgeClusterVersion *string `json:"EdgeClusterVersion,omitempty" name:"EdgeClusterVersion"`

	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UID *string `json:"UID,omitempty" name:"UID"`
}

type EdgeNodeInfo struct {

	// IECP边缘节点ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 节点名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 节点状态 （1健康｜2异常｜3离线｜4未激活）
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 节点资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *EdgeNodeResourceInfo `json:"Resource,omitempty" name:"Resource"`

	// CPU体系结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuArchitecture *string `json:"CpuArchitecture,omitempty" name:"CpuArchitecture"`

	// IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 操作系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`

	// 节点所属的NodeUnit
	// key：NodeUnit模版ID，Value：NodeUnit模版名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeUnits *KeyValueObj `json:"NodeUnits,omitempty" name:"NodeUnits"`
}

type EdgeNodeInstallerOnline struct {

	// 节点安装脚本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptName *string `json:"ScriptName,omitempty" name:"ScriptName"`

	// 节点安装脚本下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptDownloadUrl *string `json:"ScriptDownloadUrl,omitempty" name:"ScriptDownloadUrl"`

	// 节点安装命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	Guide *string `json:"Guide,omitempty" name:"Guide"`
}

type EdgeNodeLabel struct {

	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 是否受保护
	Protected *bool `json:"Protected,omitempty" name:"Protected"`
}

type EdgeNodePodContainerInfo struct {

	// Pod名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 容器ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 镜像（含版本号）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Image *string `json:"Image,omitempty" name:"Image"`

	// CPU Request
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// CPU Limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// Memory Request
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryRequest *string `json:"MemoryRequest,omitempty" name:"MemoryRequest"`

	// Memory Limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryLimit *string `json:"MemoryLimit,omitempty" name:"MemoryLimit"`

	// 重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *uint64 `json:"RestartCount,omitempty" name:"RestartCount"`

	// 容器状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type EdgeNodePodInfo struct {

	// Pod名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// Pod状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 所在节点IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeIp *string `json:"NodeIp,omitempty" name:"NodeIp"`

	// 实例IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// CPU Request
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// Memory Request
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryRequest *string `json:"MemoryRequest,omitempty" name:"MemoryRequest"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 工作负载类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`

	// 工作负载名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *uint64 `json:"RestartCount,omitempty" name:"RestartCount"`

	// 集群ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
}

type EdgeNodeResourceInfo struct {

	// 可使用的CPU 单位: m核
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllocatedCPU *string `json:"AllocatedCPU,omitempty" name:"AllocatedCPU"`

	// CPU总量 单位:m核
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCPU *string `json:"TotalCPU,omitempty" name:"TotalCPU"`

	// 已分配的内存 单位G
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllocatedMemory *string `json:"AllocatedMemory,omitempty" name:"AllocatedMemory"`

	// 内存总量 单位G
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalMemory *string `json:"TotalMemory,omitempty" name:"TotalMemory"`

	// 已分配的GPU资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllocatedGPU *string `json:"AllocatedGPU,omitempty" name:"AllocatedGPU"`

	// GPU总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalGPU *string `json:"TotalGPU,omitempty" name:"TotalGPU"`

	// 可使用的CPU 单位: m核
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableCPU *string `json:"AvailableCPU,omitempty" name:"AvailableCPU"`

	// 可使用的内存 单位: G
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableMemory *string `json:"AvailableMemory,omitempty" name:"AvailableMemory"`

	// 可使用的GPU资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableGPU *string `json:"AvailableGPU,omitempty" name:"AvailableGPU"`
}

type EdgeUnitStatisticItem struct {

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 在线数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Online *uint64 `json:"Online,omitempty" name:"Online"`

	// 异常数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Abnormal *uint64 `json:"Abnormal,omitempty" name:"Abnormal"`

	// 离线数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offline *uint64 `json:"Offline,omitempty" name:"Offline"`

	// 未激活
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotActive *uint64 `json:"NotActive,omitempty" name:"NotActive"`
}

type Env struct {

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 值引用
	ValueFrom *EnvValueSelector `json:"ValueFrom,omitempty" name:"ValueFrom"`
}

type EnvValueSelector struct {

	// 健名
	Key *string `json:"Key,omitempty" name:"Key"`

	// 对象名
	ObjectName *string `json:"ObjectName,omitempty" name:"ObjectName"`

	// 对象值
	ObjectType *string `json:"ObjectType,omitempty" name:"ObjectType"`
}

type Event struct {

	// 第一次出现时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`

	// 最后一次出现时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`

	// 事件关联对象类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvolvedObjectKind *string `json:"InvolvedObjectKind,omitempty" name:"InvolvedObjectKind"`

	// 事件关联对象名
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvolvedObjectName *string `json:"InvolvedObjectName,omitempty" name:"InvolvedObjectName"`

	// 事件类型(Normal|Warning)
	Type *string `json:"Type,omitempty" name:"Type"`

	// 原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 出现次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type FieldSort struct {

	// 字段名
	Field *string `json:"Field,omitempty" name:"Field"`

	// 排序(ASC:升序 DESC:降序
	Order *string `json:"Order,omitempty" name:"Order"`
}

type GetMarketComponentListRequest struct {
	*tchttp.BaseRequest

	// 页偏移，从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 名称模糊筛选
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 以名称排序，ASC、DESC
	Order *string `json:"Order,omitempty" name:"Order"`
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

type GetMarketComponentListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 组件列表
		ComponentList []*MarketComponentInfo `json:"ComponentList,omitempty" name:"ComponentList"`

		// 组件总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type GetMarketComponentRequest struct {
	*tchttp.BaseRequest

	// 组件ID
	ID *int64 `json:"ID,omitempty" name:"ID"`
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

type GetMarketComponentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 组件ID
		ID *int64 `json:"ID,omitempty" name:"ID"`

		// 组件名称
		AppName *string `json:"AppName,omitempty" name:"AppName"`

		// 发行组织
		Author *string `json:"Author,omitempty" name:"Author"`

		// 发布时间
		ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`

		// 组件简介
		Outline *string `json:"Outline,omitempty" name:"Outline"`

		// 详细介绍链接
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 图标连接
		Icon *string `json:"Icon,omitempty" name:"Icon"`

		// 组件版本
		Version *string `json:"Version,omitempty" name:"Version"`

		// 组件可视化配置
		WorkloadVisualConfig *string `json:"WorkloadVisualConfig,omitempty" name:"WorkloadVisualConfig"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	Name *string `json:"Name,omitempty" name:"Name"`

	// GridID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type GridEventInfo struct {

	// 首次出现时间
	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`

	// 最后出现时间
	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`

	// 对象类型
	InvolvedObjectKind *string `json:"InvolvedObjectKind,omitempty" name:"InvolvedObjectKind"`

	// 对象名称
	InvolvedObjectName *string `json:"InvolvedObjectName,omitempty" name:"InvolvedObjectName"`

	// 事件类型(Normal,Warning)
	Type *string `json:"Type,omitempty" name:"Type"`

	// 事件原因
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 事件内容
	Message *string `json:"Message,omitempty" name:"Message"`

	// 次数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 节点名（Pod事件类型时有值）
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// 节点内部IP（Pod事件类型时有值）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IP *string `json:"IP,omitempty" name:"IP"`
}

type GridInfo struct {

	// DeployGridId
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// Key
	GridUniqKey *string `json:"GridUniqKey,omitempty" name:"GridUniqKey"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 工作负载类型
	WorkloadKind *string `json:"WorkloadKind,omitempty" name:"WorkloadKind"`

	// 启动时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`

	// 创建人
	Publisher *string `json:"Publisher,omitempty" name:"Publisher"`

	// 版本信息
	Version *string `json:"Version,omitempty" name:"Version"`
}

type GridItemInfo struct {

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 期望副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`

	// 可用副本数
	AvailableReplicas *int64 `json:"AvailableReplicas,omitempty" name:"AvailableReplicas"`

	// 启动时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 工作负载类型
	WorkloadKind *string `json:"WorkloadKind,omitempty" name:"WorkloadKind"`
}

type GridPodInfo struct {

	// Pod名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命名空间
	NameSpace *string `json:"NameSpace,omitempty" name:"NameSpace"`

	// 状态(Pending｜Running｜Succeeded｜Failed｜Unknown)
	Status *string `json:"Status,omitempty" name:"Status"`

	// 节点名
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// 节点IP
	NodeIP *string `json:"NodeIP,omitempty" name:"NodeIP"`

	// Pod的IP
	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`

	// 启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 运行时长（秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunSec *int64 `json:"RunSec,omitempty" name:"RunSec"`

	// 重启次数
	RestartCount *int64 `json:"RestartCount,omitempty" name:"RestartCount"`

	// 集群名称ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
}

type HorizontalPodAutoscaler struct {

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 最小实例数
	MinReplicas *int64 `json:"MinReplicas,omitempty" name:"MinReplicas"`

	// 最大实例数
	MaxReplicas *int64 `json:"MaxReplicas,omitempty" name:"MaxReplicas"`

	// 资源目标指标
	ResourceMetricTarget []*ResourceMetricTarget `json:"ResourceMetricTarget,omitempty" name:"ResourceMetricTarget"`
}

type HttpHeader struct {

	// HTTP头的名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// HTTP头的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type HttpProbe struct {

	// 请求路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 请求端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 请求地址，默认Pod的IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 请求模式  HTTP|HTTPS，默认HTTP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// HTTP的请求头
	// 注意：此字段可能返回 null，表示取不到有效值。
	Headers []*HttpHeader `json:"Headers,omitempty" name:"Headers"`
}

type Job struct {

	// 并发数
	Parallelism *int64 `json:"Parallelism,omitempty" name:"Parallelism"`

	// 完成数
	Completion *int64 `json:"Completion,omitempty" name:"Completion"`

	// 最大运行时间
	ActiveDeadlineSeconds *int64 `json:"ActiveDeadlineSeconds,omitempty" name:"ActiveDeadlineSeconds"`

	// 失败前重试次数
	BackOffLimit *int64 `json:"BackOffLimit,omitempty" name:"BackOffLimit"`
}

type KeyValueObj struct {

	// Key值
	Key *string `json:"Key,omitempty" name:"Key"`

	// Value值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Label struct {

	// 健名
	Key *string `json:"Key,omitempty" name:"Key"`

	// 健值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type MarketComponentInfo struct {

	// 组件ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 组件名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 发布者
	Author *string `json:"Author,omitempty" name:"Author"`

	// 发布时间
	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`

	// 组件简介
	Outline *string `json:"Outline,omitempty" name:"Outline"`

	// 指向详细描述的url
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// 图标链接
	Icon *string `json:"Icon,omitempty" name:"Icon"`

	// 组件版本
	Version *string `json:"Version,omitempty" name:"Version"`

	// 组件可视化信息
	WorkloadVisualConfig *string `json:"WorkloadVisualConfig,omitempty" name:"WorkloadVisualConfig"`
}

type ModifyApplicationBasicInfoRequest struct {
	*tchttp.BaseRequest

	// 应用模板ID
	ApplicationId *uint64 `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用模板基本信息
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitempty" name:"BasicInfo"`
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

type ModifyApplicationBasicInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyApplicationVisualizationRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用配置
	BasicConfig *ApplicationBasicConfig `json:"BasicConfig,omitempty" name:"BasicConfig"`

	// 卷配置
	Volumes []*Volume `json:"Volumes,omitempty" name:"Volumes"`

	// 初始容器
	InitContainers []*Container `json:"InitContainers,omitempty" name:"InitContainers"`

	// 容器配置
	Containers []*Container `json:"Containers,omitempty" name:"Containers"`

	// 服务配置
	Service *Service `json:"Service,omitempty" name:"Service"`

	// Job配置
	Job *Job `json:"Job,omitempty" name:"Job"`

	// CronJob配置
	CronJob *CronJob `json:"CronJob,omitempty" name:"CronJob"`

	// 重启策略
	RestartPolicy *string `json:"RestartPolicy,omitempty" name:"RestartPolicy"`

	// 镜像拉取密钥
	ImagePullSecrets []*string `json:"ImagePullSecrets,omitempty" name:"ImagePullSecrets"`

	// HPA配置
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitempty" name:"HorizontalPodAutoscaler"`

	// 单个初始化容器
	InitContainer *Container `json:"InitContainer,omitempty" name:"InitContainer"`
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

type ModifyApplicationVisualizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyConfigMapRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitempty" name:"EdgeUnitID"`

	// ConfigMap名称
	ConfigMapName *string `json:"ConfigMapName,omitempty" name:"ConfigMapName"`

	// Yaml配置
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

	// ConfigMap命名空间
	ConfigMapNamespace *string `json:"ConfigMapNamespace,omitempty" name:"ConfigMapNamespace"`
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

type ModifyConfigMapResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyEdgeNodeLabelsRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// IECP边缘节点ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

	// 标签列表
	Labels []*KeyValueObj `json:"Labels,omitempty" name:"Labels"`
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

type ModifyEdgeNodeLabelsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyEdgeUnitApplicationBasicInfoRequest struct {
	*tchttp.BaseRequest

	// 应用基本信息
	BasicInfo *ApplicationBasicInfo `json:"BasicInfo,omitempty" name:"BasicInfo"`

	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitempty" name:"ApplicationId"`
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

type ModifyEdgeUnitApplicationBasicInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyEdgeUnitApplicationVisualizationRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用配置
	BasicConfig *ApplicationBasicConfig `json:"BasicConfig,omitempty" name:"BasicConfig"`

	// 卷配置
	Volumes []*Volume `json:"Volumes,omitempty" name:"Volumes"`

	// 初始容器列表
	InitContainers []*Container `json:"InitContainers,omitempty" name:"InitContainers"`

	// 容器配置
	Containers []*Container `json:"Containers,omitempty" name:"Containers"`

	// 服务配置
	Service *Service `json:"Service,omitempty" name:"Service"`

	// Job配置
	Job *Job `json:"Job,omitempty" name:"Job"`

	// CronJob配置
	CronJob *CronJob `json:"CronJob,omitempty" name:"CronJob"`

	// 重启策略
	RestartPolicy *string `json:"RestartPolicy,omitempty" name:"RestartPolicy"`

	// 镜像拉取密钥
	ImagePullSecrets []*string `json:"ImagePullSecrets,omitempty" name:"ImagePullSecrets"`

	// HPA配置
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitempty" name:"HorizontalPodAutoscaler"`
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

type ModifyEdgeUnitApplicationVisualizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyEdgeUnitApplicationYamlRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Yaml配置
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
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

type ModifyEdgeUnitApplicationYamlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyEdgeUnitDeployGridItemRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// Grid名称
	GridItemName *string `json:"GridItemName,omitempty" name:"GridItemName"`

	// 负载类型（StatefulSetGrid｜DeploymentGrid）
	WorkloadKind *string `json:"WorkloadKind,omitempty" name:"WorkloadKind"`

	// 副本数
	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`

	// 命名空间，默认default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
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

type ModifyEdgeUnitDeployGridItemResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyEdgeUnitRequest struct {
	*tchttp.BaseRequest

	// 边缘集群ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 边缘集群名称，64字符以内
	Name *string `json:"Name,omitempty" name:"Name"`

	// 集群描述，200字符以内
	Description *string `json:"Description,omitempty" name:"Description"`
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

type ModifyEdgeUnitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyNodeUnitTemplateRequest struct {
	*tchttp.BaseRequest

	// IECP边缘单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// NodeUnit模版ID
	NodeUnitTemplateID *uint64 `json:"NodeUnitTemplateID,omitempty" name:"NodeUnitTemplateID"`

	// 包含的节点列表
	Nodes []*string `json:"Nodes,omitempty" name:"Nodes"`
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

type ModifyNodeUnitTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifySecretRequest struct {
	*tchttp.BaseRequest

	// 边缘单元ID
	EdgeUnitID *uint64 `json:"EdgeUnitID,omitempty" name:"EdgeUnitID"`

	// Secret名
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// Secret的Yaml格式
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

	// Secret命名空间（默认:default）
	SecretNamespace *string `json:"SecretNamespace,omitempty" name:"SecretNamespace"`
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

type ModifySecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	ColumnName *string `json:"ColumnName,omitempty" name:"ColumnName"`

	// 数据内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnData []*string `json:"ColumnData,omitempty" name:"ColumnData"`

	// 数据所属，查询Workload类型时有值
	ColumnBelong *string `json:"ColumnBelong,omitempty" name:"ColumnBelong"`

	// 最大值
	MaxValue *float64 `json:"MaxValue,omitempty" name:"MaxValue"`

	// 最小值
	MinValue *float64 `json:"MinValue,omitempty" name:"MinValue"`

	// 平均值
	AvgValue *float64 `json:"AvgValue,omitempty" name:"AvgValue"`

	// 时间戳数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnTime *int64 `json:"ColumnTime,omitempty" name:"ColumnTime"`
}

type NamespaceInfo struct {

	// 命名空间名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 状态(Active|Terminating)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 是否保护(不允许删除)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protected *bool `json:"Protected,omitempty" name:"Protected"`

	// 对应的Yaml配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
}

type NamespaceResource struct {

	// 类型(workload|grid|configmap|secret)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 名称(最多返回5个）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Names []*string `json:"Names,omitempty" name:"Names"`
}

type NodeGroupInfo struct {

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// NodeGroup名称
	NodeGroupName *string `json:"NodeGroupName,omitempty" name:"NodeGroupName"`

	// DeploymentGrid数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeploymentGridList []*GridDetail `json:"DeploymentGridList,omitempty" name:"DeploymentGridList"`

	// StatefulSetGrid数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatefulSetGridList []*GridDetail `json:"StatefulSetGridList,omitempty" name:"StatefulSetGridList"`

	// 是否平台保护
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protect *bool `json:"Protect,omitempty" name:"Protect"`
}

type NodeGroupNodeUnitTemplateInfo struct {

	// 模版ID
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 包含节点列表
	NodeList []*NodeSimpleInfo `json:"NodeList,omitempty" name:"NodeList"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 是否关联
	Relation *bool `json:"Relation,omitempty" name:"Relation"`
}

type NodeSimpleInfo struct {

	// 节点ID
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// 节点名称
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
}

type NodeUnitInfo struct {

	// NodeUnitId
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// NodeUnit名称
	NodeUnitName *string `json:"NodeUnitName,omitempty" name:"NodeUnitName"`

	// 包含节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeList []*NodeUnitNodeInfo `json:"NodeList,omitempty" name:"NodeList"`
}

type NodeUnitNodeInfo struct {

	// 节点ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 节点状态  NodeStatusHealthy (健康)/NodeStatusAbnormal (异常)/NodeStatusOffline (下线)/NodeStatusNotActivated (未激活
	Status *string `json:"Status,omitempty" name:"Status"`

	// 节点名称
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// 内网节点IP
	InternalIP *string `json:"InternalIP,omitempty" name:"InternalIP"`
}

type NodeUnitTemplate struct {

	// NodeUnit模版ID
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// NodeUnit模版名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 包含节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeList []*NodeSimpleInfo `json:"NodeList,omitempty" name:"NodeList"`

	// NodeGroup列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeGroups []*string `json:"NodeGroups,omitempty" name:"NodeGroups"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type OperationLog struct {

	// 操作时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateTime *string `json:"OperateTime,omitempty" name:"OperateTime"`

	// 模块名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 状态: 1:成功 2:失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 操作用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUserID *string `json:"OperatorUserID,omitempty" name:"OperatorUserID"`

	// 操作动作
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitempty" name:"Action"`
}

type OperationLogsCondition struct {

	// 状态列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status []*int64 `json:"Status,omitempty" name:"Status"`
}

type PodStatus struct {

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameSpace *string `json:"NameSpace,omitempty" name:"NameSpace"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	IP *string `json:"IP,omitempty" name:"IP"`

	// 启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunSec *int64 `json:"RunSec,omitempty" name:"RunSec"`

	// 重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *int64 `json:"RestartCount,omitempty" name:"RestartCount"`
}

type PortConfig struct {

	// 协议类型(tcp|udp)
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 源端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 目标端口
	TargetPort *int64 `json:"TargetPort,omitempty" name:"TargetPort"`

	// 节点端口
	NodePort *int64 `json:"NodePort,omitempty" name:"NodePort"`
}

type Probe struct {

	// 启动后，延迟探测时间 单位:秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitialDelaySeconds *int64 `json:"InitialDelaySeconds,omitempty" name:"InitialDelaySeconds"`

	// 探测间隔，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeriodSeconds *int64 `json:"PeriodSeconds,omitempty" name:"PeriodSeconds"`

	// 探测超时时间 单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeoutSeconds *int64 `json:"TimeoutSeconds,omitempty" name:"TimeoutSeconds"`

	// 失败后检查成功的最小连续成功次数。默认为1.活跃度必须为1。最小值为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessThreshold *int64 `json:"SuccessThreshold,omitempty" name:"SuccessThreshold"`

	// 当Pod成功启动且检查失败时，放弃之前尝试次数。默认为3.最小值为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureThreshold *int64 `json:"FailureThreshold,omitempty" name:"FailureThreshold"`

	// HTTP探测配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpProbe *HttpProbe `json:"HttpProbe,omitempty" name:"HttpProbe"`

	// TCP探测配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TcpProbe *TcpProbe `json:"TcpProbe,omitempty" name:"TcpProbe"`
}

type RedeployEdgeUnitApplicationRequest struct {
	*tchttp.BaseRequest

	// 单元ID
	EdgeUnitId *uint64 `json:"EdgeUnitId,omitempty" name:"EdgeUnitId"`

	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitempty" name:"ApplicationId"`
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

type RedeployEdgeUnitApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	Type *string `json:"Type,omitempty" name:"Type"`

	// 平均值
	AverageValue *int64 `json:"AverageValue,omitempty" name:"AverageValue"`

	// 单位
	Scale *string `json:"Scale,omitempty" name:"Scale"`

	// 平均值
	AverageUtilization *int64 `json:"AverageUtilization,omitempty" name:"AverageUtilization"`
}

type SecretItem struct {

	// Secret名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Secret类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretType *string `json:"SecretType,omitempty" name:"SecretType"`
}

type SecurityCapabilities struct {

	// 允许操作列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Add []*string `json:"Add,omitempty" name:"Add"`

	// 禁止操作列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Drop []*string `json:"Drop,omitempty" name:"Drop"`
}

type SecurityContext struct {

	// 是否开启特权模式
	Privilege *bool `json:"Privilege,omitempty" name:"Privilege"`

	// 目录/Proc挂载方式
	ProcMount *string `json:"ProcMount,omitempty" name:"ProcMount"`

	// 安全配置
	Capabilities *SecurityCapabilities `json:"Capabilities,omitempty" name:"Capabilities"`
}

type Service struct {

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 类型 (ClusterIP|NodePort)
	Type *string `json:"Type,omitempty" name:"Type"`

	// 端口配置
	Ports []*PortConfig `json:"Ports,omitempty" name:"Ports"`

	// 标签
	Labels []*Label `json:"Labels,omitempty" name:"Labels"`

	// 命名空间默认default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 服务IP
	ClusterIP *string `json:"ClusterIP,omitempty" name:"ClusterIP"`
}

type Sort struct {

	// 排序字段
	Field *string `json:"Field,omitempty" name:"Field"`

	// 排序方式，升序ASC / 降序DESC
	Order *string `json:"Order,omitempty" name:"Order"`
}

type TcpProbe struct {

	// 连接端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type Volume struct {

	// 来源(emptyDir|hostPath|configMap|secret|nfs)
	Source *string `json:"Source,omitempty" name:"Source"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// Host挂载配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostPath *VolumeHostPath `json:"HostPath,omitempty" name:"HostPath"`

	// ConfigMap挂载配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigMap *VolumeConfigMap `json:"ConfigMap,omitempty" name:"ConfigMap"`

	// Secret挂载配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Secret *VolumeConfigMap `json:"Secret,omitempty" name:"Secret"`

	// NFS挂载配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	NFS *VolumeNFS `json:"NFS,omitempty" name:"NFS"`
}

type VolumeConfigMap struct {

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// Key列表配置
	Items []*VolumeConfigMapKeyToPath `json:"Items,omitempty" name:"Items"`
}

type VolumeConfigMapKeyToPath struct {

	// 健名
	Key *string `json:"Key,omitempty" name:"Key"`

	// 对应本地路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 对应权限模式
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type VolumeHostPath struct {

	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 路径
	Path *string `json:"Path,omitempty" name:"Path"`
}

type VolumeMount struct {

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 挂载路径
	MountPath *string `json:"MountPath,omitempty" name:"MountPath"`

	// 子路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubPath *string `json:"SubPath,omitempty" name:"SubPath"`

	// 是否只读
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadOnly *bool `json:"ReadOnly,omitempty" name:"ReadOnly"`
}

type VolumeNFS struct {

	// 服务地址
	Server *string `json:"Server,omitempty" name:"Server"`

	// 对应服务器路径
	ServerPath *string `json:"ServerPath,omitempty" name:"ServerPath"`

	// 对应本地路径
	Path *string `json:"Path,omitempty" name:"Path"`
}
