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

package v20190416

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Conditions struct {
	// 原因
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// 具有相同原因的副本个数
	Count *uint64 `json:"Count,omitnil" name:"Count"`
}

type Config struct {
	// Id
	Id *string `json:"Id,omitnil" name:"Id"`

	// 配置名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模型地址
	ModelUri *string `json:"ModelUri,omitnil" name:"ModelUri"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 运行环境
	Runtime *string `json:"Runtime,omitnil" name:"Runtime"`

	// 配置版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 配置描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`
}

// Predefined struct for user
type CreateJobRequestParams struct {
	// 任务名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 使用的资源组 Id，默认使用共享资源组
	ResourceGroupId *string `json:"ResourceGroupId,omitnil" name:"ResourceGroupId"`

	// 处理器配置, 单位为1/1000核；范围[100, 256000]
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存配置, 单位为1M；范围[100, 256000]
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 运行集群
	Cluster *string `json:"Cluster,omitnil" name:"Cluster"`

	// 预测输入
	PredictInput *PredictInput `json:"PredictInput,omitnil" name:"PredictInput"`

	// 任务描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 同时处理任务的 Worker 个数
	WorkerCount *uint64 `json:"WorkerCount,omitnil" name:"WorkerCount"`

	// 使用的配置 Id
	ConfigId *string `json:"ConfigId,omitnil" name:"ConfigId"`

	// GPU算力配置，单位为1/1000 卡，范围 [0, 256000]
	Gpu *uint64 `json:"Gpu,omitnil" name:"Gpu"`

	// 显存配置, 单位为1M，范围 [0, 256000]
	GpuMemory *uint64 `json:"GpuMemory,omitnil" name:"GpuMemory"`

	// GPU类型
	GpuType *string `json:"GpuType,omitnil" name:"GpuType"`

	// 量化输入
	QuantizationInput *QuantizationInput `json:"QuantizationInput,omitnil" name:"QuantizationInput"`

	// Cls日志主题ID
	LogTopicId *string `json:"LogTopicId,omitnil" name:"LogTopicId"`
}

type CreateJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 使用的资源组 Id，默认使用共享资源组
	ResourceGroupId *string `json:"ResourceGroupId,omitnil" name:"ResourceGroupId"`

	// 处理器配置, 单位为1/1000核；范围[100, 256000]
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存配置, 单位为1M；范围[100, 256000]
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 运行集群
	Cluster *string `json:"Cluster,omitnil" name:"Cluster"`

	// 预测输入
	PredictInput *PredictInput `json:"PredictInput,omitnil" name:"PredictInput"`

	// 任务描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 同时处理任务的 Worker 个数
	WorkerCount *uint64 `json:"WorkerCount,omitnil" name:"WorkerCount"`

	// 使用的配置 Id
	ConfigId *string `json:"ConfigId,omitnil" name:"ConfigId"`

	// GPU算力配置，单位为1/1000 卡，范围 [0, 256000]
	Gpu *uint64 `json:"Gpu,omitnil" name:"Gpu"`

	// 显存配置, 单位为1M，范围 [0, 256000]
	GpuMemory *uint64 `json:"GpuMemory,omitnil" name:"GpuMemory"`

	// GPU类型
	GpuType *string `json:"GpuType,omitnil" name:"GpuType"`

	// 量化输入
	QuantizationInput *QuantizationInput `json:"QuantizationInput,omitnil" name:"QuantizationInput"`

	// Cls日志主题ID
	LogTopicId *string `json:"LogTopicId,omitnil" name:"LogTopicId"`
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
	delete(f, "ResourceGroupId")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "Cluster")
	delete(f, "PredictInput")
	delete(f, "Description")
	delete(f, "WorkerCount")
	delete(f, "ConfigId")
	delete(f, "Gpu")
	delete(f, "GpuMemory")
	delete(f, "GpuType")
	delete(f, "QuantizationInput")
	delete(f, "LogTopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateJobResponseParams struct {
	// 任务
	Job *Job `json:"Job,omitnil" name:"Job"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type CreateRsgAsGroupRequestParams struct {
	// 资源组 ID
	RsgId *string `json:"RsgId,omitnil" name:"RsgId"`

	// 伸缩组允许的最大节点数
	MaxSize *uint64 `json:"MaxSize,omitnil" name:"MaxSize"`

	// 伸缩组允许的最小节点数
	MinSize *uint64 `json:"MinSize,omitnil" name:"MinSize"`

	// 伸缩组的节点规格
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 资源组所在的集群名
	Cluster *string `json:"Cluster,omitnil" name:"Cluster"`

	// 伸缩组名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 伸缩组期望的节点数
	DesiredSize *uint64 `json:"DesiredSize,omitnil" name:"DesiredSize"`
}

type CreateRsgAsGroupRequest struct {
	*tchttp.BaseRequest
	
	// 资源组 ID
	RsgId *string `json:"RsgId,omitnil" name:"RsgId"`

	// 伸缩组允许的最大节点数
	MaxSize *uint64 `json:"MaxSize,omitnil" name:"MaxSize"`

	// 伸缩组允许的最小节点数
	MinSize *uint64 `json:"MinSize,omitnil" name:"MinSize"`

	// 伸缩组的节点规格
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 资源组所在的集群名
	Cluster *string `json:"Cluster,omitnil" name:"Cluster"`

	// 伸缩组名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 伸缩组期望的节点数
	DesiredSize *uint64 `json:"DesiredSize,omitnil" name:"DesiredSize"`
}

func (r *CreateRsgAsGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRsgAsGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RsgId")
	delete(f, "MaxSize")
	delete(f, "MinSize")
	delete(f, "InstanceType")
	delete(f, "Cluster")
	delete(f, "Name")
	delete(f, "DesiredSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRsgAsGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRsgAsGroupResponseParams struct {
	// 所创建的资源组的伸缩组
	RsgAsGroup *RsgAsGroup `json:"RsgAsGroup,omitnil" name:"RsgAsGroup"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateRsgAsGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateRsgAsGroupResponseParams `json:"Response"`
}

func (r *CreateRsgAsGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRsgAsGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuntimeRequestParams struct {
	// 全局唯一的运行环境名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 运行环境镜像地址
	Image *string `json:"Image,omitnil" name:"Image"`

	// 运行环境框架
	Framework *string `json:"Framework,omitnil" name:"Framework"`

	// 运行环境描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 是否支持健康检查，默认为False
	HealthCheckOn *bool `json:"HealthCheckOn,omitnil" name:"HealthCheckOn"`
}

type CreateRuntimeRequest struct {
	*tchttp.BaseRequest
	
	// 全局唯一的运行环境名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 运行环境镜像地址
	Image *string `json:"Image,omitnil" name:"Image"`

	// 运行环境框架
	Framework *string `json:"Framework,omitnil" name:"Framework"`

	// 运行环境描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 是否支持健康检查，默认为False
	HealthCheckOn *bool `json:"HealthCheckOn,omitnil" name:"HealthCheckOn"`
}

func (r *CreateRuntimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuntimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Image")
	delete(f, "Framework")
	delete(f, "Description")
	delete(f, "HealthCheckOn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRuntimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuntimeResponseParams struct {
	// 运行环境
	Runtime *Runtime `json:"Runtime,omitnil" name:"Runtime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateRuntimeResponse struct {
	*tchttp.BaseResponse
	Response *CreateRuntimeResponseParams `json:"Response"`
}

func (r *CreateRuntimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuntimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceConfigRequestParams struct {
	// 配置名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 运行环境
	Runtime *string `json:"Runtime,omitnil" name:"Runtime"`

	// 模型地址，支持cos路径，格式为 cos://bucket名-appid.cos.region名.myqcloud.com/模型文件夹路径。为模型文件的上一层文件夹地址。
	ModelUri *string `json:"ModelUri,omitnil" name:"ModelUri"`

	// 配置描述
	Description *string `json:"Description,omitnil" name:"Description"`
}

type CreateServiceConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 运行环境
	Runtime *string `json:"Runtime,omitnil" name:"Runtime"`

	// 模型地址，支持cos路径，格式为 cos://bucket名-appid.cos.region名.myqcloud.com/模型文件夹路径。为模型文件的上一层文件夹地址。
	ModelUri *string `json:"ModelUri,omitnil" name:"ModelUri"`

	// 配置描述
	Description *string `json:"Description,omitnil" name:"Description"`
}

func (r *CreateServiceConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Runtime")
	delete(f, "ModelUri")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServiceConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceConfigResponseParams struct {
	// 服务配置
	ServiceConfig *Config `json:"ServiceConfig,omitnil" name:"ServiceConfig"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateServiceConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateServiceConfigResponseParams `json:"Response"`
}

func (r *CreateServiceConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceRequestParams struct {
	// 扩缩容配置
	Scaler *Scaler `json:"Scaler,omitnil" name:"Scaler"`

	// 服务配置Id
	ServiceConfigId *string `json:"ServiceConfigId,omitnil" name:"ServiceConfigId"`

	// 服务名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 扩缩容方式，支持AUTO, MANUAL，分别表示自动扩缩容和手动扩缩容
	ScaleMode *string `json:"ScaleMode,omitnil" name:"ScaleMode"`

	// 部署要使用的资源组Id，默认为共享资源组
	ResourceGroupId *string `json:"ResourceGroupId,omitnil" name:"ResourceGroupId"`

	// 处理器配置, 单位为1/1000核；范围[100, 256000]
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存配置, 单位为1M；范围[100, 256000]
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 集群，不填则使用默认集群
	Cluster *string `json:"Cluster,omitnil" name:"Cluster"`

	// 默认为空，表示不需要鉴权，TOKEN 表示选择 Token 鉴权方式
	Authentication *string `json:"Authentication,omitnil" name:"Authentication"`

	// GPU算力配置，单位为1/1000 卡，范围 [0, 256000]
	Gpu *uint64 `json:"Gpu,omitnil" name:"Gpu"`

	// 显存配置, 单位为1M，范围 [0, 256000]
	GpuMemory *uint64 `json:"GpuMemory,omitnil" name:"GpuMemory"`

	// 备注
	Description *string `json:"Description,omitnil" name:"Description"`

	// GPU类型
	GpuType *string `json:"GpuType,omitnil" name:"GpuType"`

	// Cls日志主题ID
	LogTopicId *string `json:"LogTopicId,omitnil" name:"LogTopicId"`
}

type CreateServiceRequest struct {
	*tchttp.BaseRequest
	
	// 扩缩容配置
	Scaler *Scaler `json:"Scaler,omitnil" name:"Scaler"`

	// 服务配置Id
	ServiceConfigId *string `json:"ServiceConfigId,omitnil" name:"ServiceConfigId"`

	// 服务名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 扩缩容方式，支持AUTO, MANUAL，分别表示自动扩缩容和手动扩缩容
	ScaleMode *string `json:"ScaleMode,omitnil" name:"ScaleMode"`

	// 部署要使用的资源组Id，默认为共享资源组
	ResourceGroupId *string `json:"ResourceGroupId,omitnil" name:"ResourceGroupId"`

	// 处理器配置, 单位为1/1000核；范围[100, 256000]
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存配置, 单位为1M；范围[100, 256000]
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 集群，不填则使用默认集群
	Cluster *string `json:"Cluster,omitnil" name:"Cluster"`

	// 默认为空，表示不需要鉴权，TOKEN 表示选择 Token 鉴权方式
	Authentication *string `json:"Authentication,omitnil" name:"Authentication"`

	// GPU算力配置，单位为1/1000 卡，范围 [0, 256000]
	Gpu *uint64 `json:"Gpu,omitnil" name:"Gpu"`

	// 显存配置, 单位为1M，范围 [0, 256000]
	GpuMemory *uint64 `json:"GpuMemory,omitnil" name:"GpuMemory"`

	// 备注
	Description *string `json:"Description,omitnil" name:"Description"`

	// GPU类型
	GpuType *string `json:"GpuType,omitnil" name:"GpuType"`

	// Cls日志主题ID
	LogTopicId *string `json:"LogTopicId,omitnil" name:"LogTopicId"`
}

func (r *CreateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Scaler")
	delete(f, "ServiceConfigId")
	delete(f, "Name")
	delete(f, "ScaleMode")
	delete(f, "ResourceGroupId")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "Cluster")
	delete(f, "Authentication")
	delete(f, "Gpu")
	delete(f, "GpuMemory")
	delete(f, "Description")
	delete(f, "GpuType")
	delete(f, "LogTopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceResponseParams struct {
	// 服务
	Service *ModelService `json:"Service,omitnil" name:"Service"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateServiceResponseParams `json:"Response"`
}

func (r *CreateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceRequestParams struct {
	// 要删除的节点 ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的节点 ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceResponseParams `json:"Response"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteJobRequestParams struct {
	// 任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type DeleteJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *DeleteJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteJobResponse struct {
	*tchttp.BaseResponse
	Response *DeleteJobResponseParams `json:"Response"`
}

func (r *DeleteJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceGroupRequestParams struct {
	// 要删除的资源组 ID
	ResourceGroupId *string `json:"ResourceGroupId,omitnil" name:"ResourceGroupId"`
}

type DeleteResourceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的资源组 ID
	ResourceGroupId *string `json:"ResourceGroupId,omitnil" name:"ResourceGroupId"`
}

func (r *DeleteResourceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteResourceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteResourceGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteResourceGroupResponseParams `json:"Response"`
}

func (r *DeleteResourceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRsgAsGroupRequestParams struct {
	// 伸缩组 ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DeleteRsgAsGroupRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组 ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DeleteRsgAsGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRsgAsGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRsgAsGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRsgAsGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRsgAsGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRsgAsGroupResponseParams `json:"Response"`
}

func (r *DeleteRsgAsGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRsgAsGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRuntimeRequestParams struct {
	// 要删除的Runtime名
	Runtime *string `json:"Runtime,omitnil" name:"Runtime"`
}

type DeleteRuntimeRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的Runtime名
	Runtime *string `json:"Runtime,omitnil" name:"Runtime"`
}

func (r *DeleteRuntimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuntimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Runtime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRuntimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRuntimeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRuntimeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRuntimeResponseParams `json:"Response"`
}

func (r *DeleteRuntimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuntimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceConfigRequestParams struct {
	// 服务配置Id
	ServiceConfigId *string `json:"ServiceConfigId,omitnil" name:"ServiceConfigId"`

	// 服务配置名称
	ServiceConfigName *string `json:"ServiceConfigName,omitnil" name:"ServiceConfigName"`
}

type DeleteServiceConfigRequest struct {
	*tchttp.BaseRequest
	
	// 服务配置Id
	ServiceConfigId *string `json:"ServiceConfigId,omitnil" name:"ServiceConfigId"`

	// 服务配置名称
	ServiceConfigName *string `json:"ServiceConfigName,omitnil" name:"ServiceConfigName"`
}

func (r *DeleteServiceConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceConfigId")
	delete(f, "ServiceConfigName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServiceConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteServiceConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServiceConfigResponseParams `json:"Response"`
}

func (r *DeleteServiceConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceRequestParams struct {
	// 服务Id
	ServiceId *string `json:"ServiceId,omitnil" name:"ServiceId"`
}

type DeleteServiceRequest struct {
	*tchttp.BaseRequest
	
	// 服务Id
	ServiceId *string `json:"ServiceId,omitnil" name:"ServiceId"`
}

func (r *DeleteServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteServiceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServiceResponseParams `json:"Response"`
}

func (r *DeleteServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 筛选选项
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序的依据字段， 取值范围 "CREATE_TIME", "UPDATE_TIME", "NAME"
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 要查询的资源组 ID
	ResourceGroupId *string `json:"ResourceGroupId,omitnil" name:"ResourceGroupId"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 筛选选项
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序的依据字段， 取值范围 "CREATE_TIME", "UPDATE_TIME", "NAME"
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 要查询的资源组 ID
	ResourceGroupId *string `json:"ResourceGroupId,omitnil" name:"ResourceGroupId"`
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
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "ResourceGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 资源组下节点总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 资源组下节点列表
	Instances []*Instance `json:"Instances,omitnil" name:"Instances"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeResourceGroupsRequestParams struct {
	// 筛选选项
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序的依据字段， 取值范围 "CREATE_TIME", "UPDATE_TIME", "NAME"
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`
}

type DescribeResourceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 筛选选项
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序的依据字段， 取值范围 "CREATE_TIME", "UPDATE_TIME", "NAME"
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`
}

func (r *DescribeResourceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceGroupsResponseParams struct {
	// 资源组总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 资源组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroups []*ResourceGroup `json:"ResourceGroups,omitnil" name:"ResourceGroups"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeResourceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceGroupsResponseParams `json:"Response"`
}

func (r *DescribeResourceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRsgAsGroupActivitiesRequestParams struct {
	// 伸缩组 ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 查询活动的开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询互动的结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 筛选选项
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为 0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为 20，最大值为 200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 输出列表的排列顺序。取值范围："ASC", "DESC"
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序的依据字段， 取值范围 "CREATE_TIME", "UPDATE_TIME", "NAME"
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`
}

type DescribeRsgAsGroupActivitiesRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组 ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 查询活动的开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询互动的结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 筛选选项
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为 0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为 20，最大值为 200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 输出列表的排列顺序。取值范围："ASC", "DESC"
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序的依据字段， 取值范围 "CREATE_TIME", "UPDATE_TIME", "NAME"
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`
}

func (r *DescribeRsgAsGroupActivitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRsgAsGroupActivitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRsgAsGroupActivitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRsgAsGroupActivitiesResponseParams struct {
	// 伸缩组活动数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	RsgAsGroupActivitySet []*RsgAsGroupActivity `json:"RsgAsGroupActivitySet,omitnil" name:"RsgAsGroupActivitySet"`

	// 所查询的伸缩组活动总数目
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRsgAsGroupActivitiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRsgAsGroupActivitiesResponseParams `json:"Response"`
}

func (r *DescribeRsgAsGroupActivitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRsgAsGroupActivitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRsgAsGroupsRequestParams struct {
	// 筛选选项
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为 0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为 20，最大值为 200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 输出列表的排列顺序。取值范围："ASC", "DESC"
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序的依据字段， 取值范围 "CREATE_TIME", "UPDATE_TIME", "NAME"
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`
}

type DescribeRsgAsGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 筛选选项
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为 0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为 20，最大值为 200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 输出列表的排列顺序。取值范围："ASC", "DESC"
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序的依据字段， 取值范围 "CREATE_TIME", "UPDATE_TIME", "NAME"
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`
}

func (r *DescribeRsgAsGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRsgAsGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRsgAsGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRsgAsGroupsResponseParams struct {
	// 所查询的伸缩组数组
	RsgAsGroupSet []*RsgAsGroup `json:"RsgAsGroupSet,omitnil" name:"RsgAsGroupSet"`

	// 伸缩组数组总数目
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRsgAsGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRsgAsGroupsResponseParams `json:"Response"`
}

func (r *DescribeRsgAsGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRsgAsGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuntimesRequestParams struct {

}

type DescribeRuntimesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRuntimesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuntimesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuntimesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuntimesResponseParams struct {
	// TIEMS支持的运行环境列表
	Runtimes []*Runtime `json:"Runtimes,omitnil" name:"Runtimes"`

	// 用户对runtime对权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAccess *uint64 `json:"UserAccess,omitnil" name:"UserAccess"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRuntimesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuntimesResponseParams `json:"Response"`
}

func (r *DescribeRuntimesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuntimesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceConfigsRequestParams struct {
	// 筛选选项，支持按照name等进行筛选
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为1000
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序的依据字段， 取值范围 "CREATE_TIME", "UPDATE_TIME", "NAME"
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 是否按照配置名分页
	PageByName *bool `json:"PageByName,omitnil" name:"PageByName"`
}

type DescribeServiceConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 筛选选项，支持按照name等进行筛选
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为1000
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序的依据字段， 取值范围 "CREATE_TIME", "UPDATE_TIME", "NAME"
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 是否按照配置名分页
	PageByName *bool `json:"PageByName,omitnil" name:"PageByName"`
}

func (r *DescribeServiceConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "PageByName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceConfigsResponseParams struct {
	// 服务配置
	ServiceConfigs []*Config `json:"ServiceConfigs,omitnil" name:"ServiceConfigs"`

	// 服务配置总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeServiceConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceConfigsResponseParams `json:"Response"`
}

func (r *DescribeServiceConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServicesRequestParams struct {
	// 筛选选项，支持筛选的字段：id, region, zone, cluster, status, runtime, rsg_id
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序的依据字段， 取值范围 "CREATE_TIME" "UPDATE_TIME"
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`
}

type DescribeServicesRequest struct {
	*tchttp.BaseRequest
	
	// 筛选选项，支持筛选的字段：id, region, zone, cluster, status, runtime, rsg_id
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序的依据字段， 取值范围 "CREATE_TIME" "UPDATE_TIME"
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`
}

func (r *DescribeServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServicesResponseParams struct {
	// 服务列表
	Services []*ModelService `json:"Services,omitnil" name:"Services"`

	// 服务总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeServicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServicesResponseParams `json:"Response"`
}

func (r *DescribeServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableRsgAsGroupRequestParams struct {
	// 伸缩组 ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DisableRsgAsGroupRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组 ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DisableRsgAsGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableRsgAsGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableRsgAsGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableRsgAsGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisableRsgAsGroupResponse struct {
	*tchttp.BaseResponse
	Response *DisableRsgAsGroupResponseParams `json:"Response"`
}

func (r *DisableRsgAsGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableRsgAsGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableRsgAsGroupRequestParams struct {
	// 伸缩组 ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

type EnableRsgAsGroupRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组 ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *EnableRsgAsGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableRsgAsGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableRsgAsGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableRsgAsGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnableRsgAsGroupResponse struct {
	*tchttp.BaseResponse
	Response *EnableRsgAsGroupResponseParams `json:"Response"`
}

func (r *EnableRsgAsGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableRsgAsGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExposeInfo struct {
	// 暴露方式，支持 EXTERNAL（外网暴露），VPC （VPC内网打通）
	ExposeType *string `json:"ExposeType,omitnil" name:"ExposeType"`

	// 暴露Ip。暴露方式为 EXTERNAL 为外网 Ip，暴露方式为 VPC 时为指定 Vpc 下的Vip
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 暴露方式为 VPC 时，打通的私有网络Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 暴露方式为 VPC 时，打通的子网Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// GATEWAY 服务id，ExposeType = GATEWAY 时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	GateWayServiceId *string `json:"GateWayServiceId,omitnil" name:"GateWayServiceId"`

	// GATEWAY api id，ExposeType = GATEWAY 时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	GateWayAPIId *string `json:"GateWayAPIId,omitnil" name:"GateWayAPIId"`

	// GATEWAY domain，ExposeType = GATEWAY 时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	GateWayDomain *string `json:"GateWayDomain,omitnil" name:"GateWayDomain"`
}

// Predefined struct for user
type ExposeServiceRequestParams struct {
	// 服务Id
	ServiceId *string `json:"ServiceId,omitnil" name:"ServiceId"`

	// 暴露方式，支持 EXTERNAL（外网暴露），VPC （VPC内网打通）
	ExposeType *string `json:"ExposeType,omitnil" name:"ExposeType"`

	// 暴露方式为 VPC 时，填写需要打通的私有网络Id
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 暴露方式为 VPC 时，填写需要打通的子网Id
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}

type ExposeServiceRequest struct {
	*tchttp.BaseRequest
	
	// 服务Id
	ServiceId *string `json:"ServiceId,omitnil" name:"ServiceId"`

	// 暴露方式，支持 EXTERNAL（外网暴露），VPC （VPC内网打通）
	ExposeType *string `json:"ExposeType,omitnil" name:"ExposeType"`

	// 暴露方式为 VPC 时，填写需要打通的私有网络Id
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 暴露方式为 VPC 时，填写需要打通的子网Id
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}

func (r *ExposeServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExposeServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ExposeType")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExposeServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExposeServiceResponseParams struct {
	// 暴露方式
	Expose *ExposeInfo `json:"Expose,omitnil" name:"Expose"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExposeServiceResponse struct {
	*tchttp.BaseResponse
	Response *ExposeServiceResponseParams `json:"Response"`
}

func (r *ExposeServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExposeServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 取值
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type Instance struct {
	// 节点 ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 节点所在地区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 节点类型
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 节点充值类型
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// Cpu 核数
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// Gpu 核数
	Gpu *uint64 `json:"Gpu,omitnil" name:"Gpu"`

	// 节点状态
	State *string `json:"State,omitnil" name:"State"`

	// 节点故障信息
	AbnormalReason *string `json:"AbnormalReason,omitnil" name:"AbnormalReason"`

	// 创建时间
	Created *string `json:"Created,omitnil" name:"Created"`

	// 更新时间
	Updated *string `json:"Updated,omitnil" name:"Updated"`

	// 到期时间
	DeadlineTime *string `json:"DeadlineTime,omitnil" name:"DeadlineTime"`

	// 所属资源组 ID
	ResourceGroupId *string `json:"ResourceGroupId,omitnil" name:"ResourceGroupId"`

	// 自动续费标签
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// 节点所在地域
	Region *string `json:"Region,omitnil" name:"Region"`

	// 当前 Cpu 申请使用量
	CpuRequested *uint64 `json:"CpuRequested,omitnil" name:"CpuRequested"`

	// 当前 Memory 申请使用量
	MemoryRequested *uint64 `json:"MemoryRequested,omitnil" name:"MemoryRequested"`

	// 当前 Gpu 申请使用量
	GpuRequested *uint64 `json:"GpuRequested,omitnil" name:"GpuRequested"`

	// 节点所在伸缩组 ID
	RsgAsGroupId *string `json:"RsgAsGroupId,omitnil" name:"RsgAsGroupId"`
}

type Job struct {
	// 任务 Id
	Id *string `json:"Id,omitnil" name:"Id"`

	// 集群名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cluster *string `json:"Cluster,omitnil" name:"Cluster"`

	// Region 名
	Region *string `json:"Region,omitnil" name:"Region"`

	// 任务名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// Worker 使用的运行环境
	// 注意：此字段可能返回 null，表示取不到有效值。
	Runtime *string `json:"Runtime,omitnil" name:"Runtime"`

	// 任务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 配置 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitnil" name:"ConfigId"`

	// 预测输入
	// 注意：此字段可能返回 null，表示取不到有效值。
	PredictInput *PredictInput `json:"PredictInput,omitnil" name:"PredictInput"`

	// 任务状态
	Status *JobStatus `json:"Status,omitnil" name:"Status"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 任务开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 任务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 任务取消时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CancelTime *string `json:"CancelTime,omitnil" name:"CancelTime"`

	// 任务使用资源组 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil" name:"ResourceGroupId"`

	// 处理器配置, 单位为1/1000核；范围[100, 256000]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存配置, 单位为1M；范围[100, 256000]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// GPU算力配置，单位为1/1000 卡，范围 [0, 256000]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gpu *uint64 `json:"Gpu,omitnil" name:"Gpu"`

	// 显存配置, 单位为1M，范围 [0, 256000]
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuMemory *uint64 `json:"GpuMemory,omitnil" name:"GpuMemory"`

	// 任务使用资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitnil" name:"ResourceGroupName"`

	// GPU类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuType *string `json:"GpuType,omitnil" name:"GpuType"`

	// 配置名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitnil" name:"ConfigName"`

	// 配置版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersion *string `json:"ConfigVersion,omitnil" name:"ConfigVersion"`

	// Job类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobType *string `json:"JobType,omitnil" name:"JobType"`

	// 量化输入
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuantizationInput *QuantizationInput `json:"QuantizationInput,omitnil" name:"QuantizationInput"`

	// Cls日志主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogTopicId *string `json:"LogTopicId,omitnil" name:"LogTopicId"`
}

type JobStatus struct {
	// 任务状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误时为错误描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 预期Worker数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DesiredWorkers *uint64 `json:"DesiredWorkers,omitnil" name:"DesiredWorkers"`

	// 当前Worker数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentWorkers *uint64 `json:"CurrentWorkers,omitnil" name:"CurrentWorkers"`

	// 副本名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas []*string `json:"Replicas,omitnil" name:"Replicas"`

	// 副本实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplicaInfos []*ReplicaInfo `json:"ReplicaInfos,omitnil" name:"ReplicaInfos"`
}

type ModelService struct {
	// 服务ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 运行集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cluster *string `json:"Cluster,omitnil" name:"Cluster"`

	// 服务名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 运行环境
	Runtime *string `json:"Runtime,omitnil" name:"Runtime"`

	// 模型地址
	ModelUri *string `json:"ModelUri,omitnil" name:"ModelUri"`

	// 处理器配置, 单位为1/1000核
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存配置, 单位为1M
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// GPU 配置, 单位为1/1000 卡
	Gpu *uint64 `json:"Gpu,omitnil" name:"Gpu"`

	// 显存配置, 单位为1M
	GpuMemory *uint64 `json:"GpuMemory,omitnil" name:"GpuMemory"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 支持AUTO, MANUAL
	ScaleMode *string `json:"ScaleMode,omitnil" name:"ScaleMode"`

	// 弹性伸缩配置
	Scaler *Scaler `json:"Scaler,omitnil" name:"Scaler"`

	// 服务状态
	Status *ServiceStatus `json:"Status,omitnil" name:"Status"`

	// 访问密钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessToken *string `json:"AccessToken,omitnil" name:"AccessToken"`

	// 服务配置Id
	ConfigId *string `json:"ConfigId,omitnil" name:"ConfigId"`

	// 服务配置名
	ConfigName *string `json:"ConfigName,omitnil" name:"ConfigName"`

	// 服务运行时长
	ServeSeconds *uint64 `json:"ServeSeconds,omitnil" name:"ServeSeconds"`

	// 配置版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersion *string `json:"ConfigVersion,omitnil" name:"ConfigVersion"`

	// 服务使用资源组 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil" name:"ResourceGroupId"`

	// 暴露方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Exposes []*ExposeInfo `json:"Exposes,omitnil" name:"Exposes"`

	// Region 名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 服务使用资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitnil" name:"ResourceGroupName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// GPU类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuType *string `json:"GpuType,omitnil" name:"GpuType"`

	// Cls日志主题Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogTopicId *string `json:"LogTopicId,omitnil" name:"LogTopicId"`
}

type Option struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 取值
	Value *int64 `json:"Value,omitnil" name:"Value"`
}

type PredictInput struct {
	// 输入路径，支持 cos 格式路径文件夹或文件
	InputPath *string `json:"InputPath,omitnil" name:"InputPath"`

	// 输出路径，支持 cos 格式路径
	OutputPath *string `json:"OutputPath,omitnil" name:"OutputPath"`

	// 输入数据格式，目前支持：JSON
	InputDataFormat *string `json:"InputDataFormat,omitnil" name:"InputDataFormat"`

	// 输出数据格式，目前支持：JSON
	OutputDataFormat *string `json:"OutputDataFormat,omitnil" name:"OutputDataFormat"`

	// 预测批大小，默认为 64
	BatchSize *uint64 `json:"BatchSize,omitnil" name:"BatchSize"`

	// 模型签名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignatureName *string `json:"SignatureName,omitnil" name:"SignatureName"`
}

type QuantizationInput struct {
	// 量化输入路径
	InputPath *string `json:"InputPath,omitnil" name:"InputPath"`

	// 量化输出路径
	OutputPath *string `json:"OutputPath,omitnil" name:"OutputPath"`

	// 量化批大小
	BatchSize *uint64 `json:"BatchSize,omitnil" name:"BatchSize"`

	// 量化精度，支持：FP32，FP16，INT8
	Precision *string `json:"Precision,omitnil" name:"Precision"`

	// 转换类型
	ConvertType *string `json:"ConvertType,omitnil" name:"ConvertType"`
}

type ReplicaInfo struct {
	// 实例名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 弹性网卡模式时，弹性网卡Ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	EniIp *string `json:"EniIp,omitnil" name:"EniIp"`

	// Normal: 正常运行中; Abnormal: 异常；Waiting：等待中
	Status *string `json:"Status,omitnil" name:"Status"`

	// 当 status为 Abnormal 的时候，一些额外的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 重启次数
	Restarted *uint64 `json:"Restarted,omitnil" name:"Restarted"`
}

type ResourceGroup struct {
	// 资源组 Id
	Id *string `json:"Id,omitnil" name:"Id"`

	// 地域
	Region *string `json:"Region,omitnil" name:"Region"`

	// 集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cluster *string `json:"Cluster,omitnil" name:"Cluster"`

	// 资源组名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 资源组描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 创建时间
	Created *string `json:"Created,omitnil" name:"Created"`

	// 更新时间
	Updated *string `json:"Updated,omitnil" name:"Updated"`

	// 资源组主机数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *uint64 `json:"InstanceCount,omitnil" name:"InstanceCount"`

	// 使用资源组的服务数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceCount *uint64 `json:"ServiceCount,omitnil" name:"ServiceCount"`

	// 使用资源组的任务数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobCount *uint64 `json:"JobCount,omitnil" name:"JobCount"`

	// 资源组是否为公共资源组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Public *bool `json:"Public,omitnil" name:"Public"`

	// 机器类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 资源组状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 显卡总张数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gpu *uint64 `json:"Gpu,omitnil" name:"Gpu"`

	// 处理器总核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存总量，单位为G
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Gpu类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuType []*string `json:"GpuType,omitnil" name:"GpuType"`

	// 该资源组下是否有预付费资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasPrepaid *bool `json:"HasPrepaid,omitnil" name:"HasPrepaid"`

	// 资源组是否允许预付费或后付费模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`
}

type RsgAsActivityRelatedInstance struct {
	// 节点 ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 节点状态
	InstanceStatus *string `json:"InstanceStatus,omitnil" name:"InstanceStatus"`
}

type RsgAsGroup struct {
	// 伸缩组 ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 伸缩组所在地域
	Region *string `json:"Region,omitnil" name:"Region"`

	// 伸缩组所在可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 伸缩组所在集群
	Cluster *string `json:"Cluster,omitnil" name:"Cluster"`

	// 伸缩组所在资源组 ID
	RsgId *string `json:"RsgId,omitnil" name:"RsgId"`

	// 伸缩组名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 伸缩组允许的最大节点个数
	MaxSize *uint64 `json:"MaxSize,omitnil" name:"MaxSize"`

	// 伸缩组允许的最小节点个数
	MinSize *uint64 `json:"MinSize,omitnil" name:"MinSize"`

	// 伸缩组创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 伸缩组更新时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 伸缩组状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 伸缩组节点类型
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 伸缩组内节点个数
	InstanceCount *uint64 `json:"InstanceCount,omitnil" name:"InstanceCount"`

	// 伸缩组起始节点数
	DesiredSize *uint64 `json:"DesiredSize,omitnil" name:"DesiredSize"`
}

type RsgAsGroupActivity struct {
	// 伸缩组活动 ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 关联的伸缩组 ID
	RsgAsGroupId *string `json:"RsgAsGroupId,omitnil" name:"RsgAsGroupId"`

	// 活动类型
	ActivityType *string `json:"ActivityType,omitnil" name:"ActivityType"`

	// 状态的编码
	StatusCode *string `json:"StatusCode,omitnil" name:"StatusCode"`

	// 状态的消息
	StatusMessage *string `json:"StatusMessage,omitnil" name:"StatusMessage"`

	// 活动原因
	Cause *string `json:"Cause,omitnil" name:"Cause"`

	// 活动描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 活动开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 活动结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 活动创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 活动相关联的节点
	RsgAsActivityRelatedInstance []*RsgAsActivityRelatedInstance `json:"RsgAsActivityRelatedInstance,omitnil" name:"RsgAsActivityRelatedInstance"`

	// 简略的状态消息
	StatusMessageSimplified *string `json:"StatusMessageSimplified,omitnil" name:"StatusMessageSimplified"`
}

type Runtime struct {
	// 运行环境名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 运行环境框架
	Framework *string `json:"Framework,omitnil" name:"Framework"`

	// 运行环境描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 是否为公开运行环境
	// 注意：此字段可能返回 null，表示取不到有效值。
	Public *bool `json:"Public,omitnil" name:"Public"`

	// 是否打开健康检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheckOn *bool `json:"HealthCheckOn,omitnil" name:"HealthCheckOn"`

	// 镜像地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Image *string `json:"Image,omitnil" name:"Image"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`
}

type Scaler struct {
	// 最大副本数，ScaleMode 为 MANUAL 时辞会此值会被置为 StartReplicas 取值
	MaxReplicas *uint64 `json:"MaxReplicas,omitnil" name:"MaxReplicas"`

	// 最小副本数，ScaleMode 为 MANUAL 时辞会此值会被置为 StartReplicas 取值
	MinReplicas *uint64 `json:"MinReplicas,omitnil" name:"MinReplicas"`

	// 起始副本数
	StartReplicas *uint64 `json:"StartReplicas,omitnil" name:"StartReplicas"`

	// 扩缩容指标，选择自动扩缩容时至少需要选择一个指标，支持CPU-UTIL、MEMORY-UTIL
	HpaMetrics []*Option `json:"HpaMetrics,omitnil" name:"HpaMetrics"`
}

type ServiceStatus struct {
	// 预期副本数
	DesiredReplicas *uint64 `json:"DesiredReplicas,omitnil" name:"DesiredReplicas"`

	// 当前副本数
	CurrentReplicas *uint64 `json:"CurrentReplicas,omitnil" name:"CurrentReplicas"`

	// Normal：正常运行中；Abnormal：服务异常，例如容器启动失败等；Waiting：服务等待中，例如容器下载镜像过程等；Stopped：已停止 Stopping 停止中；Resuming：重启中；Updating：服务更新中
	Status *string `json:"Status,omitnil" name:"Status"`

	// 服务处于当前状态的原因集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conditions []*Conditions `json:"Conditions,omitnil" name:"Conditions"`

	// 副本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas []*string `json:"Replicas,omitnil" name:"Replicas"`

	// 运行状态对额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 副本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplicaInfos []*ReplicaInfo `json:"ReplicaInfos,omitnil" name:"ReplicaInfos"`
}

// Predefined struct for user
type UpdateJobRequestParams struct {
	// 任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 任务更新动作，支持：Cancel
	JobAction *string `json:"JobAction,omitnil" name:"JobAction"`

	// 备注
	Description *string `json:"Description,omitnil" name:"Description"`
}

type UpdateJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 任务更新动作，支持：Cancel
	JobAction *string `json:"JobAction,omitnil" name:"JobAction"`

	// 备注
	Description *string `json:"Description,omitnil" name:"Description"`
}

func (r *UpdateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobAction")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateJobResponseParams struct {
	// 任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	Job *Job `json:"Job,omitnil" name:"Job"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateJobResponse struct {
	*tchttp.BaseResponse
	Response *UpdateJobResponseParams `json:"Response"`
}

func (r *UpdateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRsgAsGroupRequestParams struct {
	// 伸缩组 ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 重命名名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 伸缩组最大节点数
	MaxSize *uint64 `json:"MaxSize,omitnil" name:"MaxSize"`

	// 伸缩组最小节点数
	MinSize *uint64 `json:"MinSize,omitnil" name:"MinSize"`

	// 伸缩组期望的节点数
	DesiredSize *uint64 `json:"DesiredSize,omitnil" name:"DesiredSize"`
}

type UpdateRsgAsGroupRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组 ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 重命名名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 伸缩组最大节点数
	MaxSize *uint64 `json:"MaxSize,omitnil" name:"MaxSize"`

	// 伸缩组最小节点数
	MinSize *uint64 `json:"MinSize,omitnil" name:"MinSize"`

	// 伸缩组期望的节点数
	DesiredSize *uint64 `json:"DesiredSize,omitnil" name:"DesiredSize"`
}

func (r *UpdateRsgAsGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRsgAsGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "MaxSize")
	delete(f, "MinSize")
	delete(f, "DesiredSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRsgAsGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRsgAsGroupResponseParams struct {
	// 资源组的伸缩组
	RsgAsGroup *RsgAsGroup `json:"RsgAsGroup,omitnil" name:"RsgAsGroup"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateRsgAsGroupResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRsgAsGroupResponseParams `json:"Response"`
}

func (r *UpdateRsgAsGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRsgAsGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateServiceRequestParams struct {
	// 服务Id
	ServiceId *string `json:"ServiceId,omitnil" name:"ServiceId"`

	// 扩缩容配置
	Scaler *Scaler `json:"Scaler,omitnil" name:"Scaler"`

	// 服务配置Id
	ServiceConfigId *string `json:"ServiceConfigId,omitnil" name:"ServiceConfigId"`

	// 支持AUTO, MANUAL，分别表示自动扩缩容，手动扩缩容
	ScaleMode *string `json:"ScaleMode,omitnil" name:"ScaleMode"`

	// 支持STOP(停止) RESUME(重启)
	ServiceAction *string `json:"ServiceAction,omitnil" name:"ServiceAction"`

	// 备注
	Description *string `json:"Description,omitnil" name:"Description"`

	// GPU卡类型
	GpuType *string `json:"GpuType,omitnil" name:"GpuType"`

	// 处理器配置，单位为 1/1000 核
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存配置，单位为1M
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 显卡配置，单位为 1/1000 卡
	Gpu *uint64 `json:"Gpu,omitnil" name:"Gpu"`

	// Cls日志主题ID
	LogTopicId *string `json:"LogTopicId,omitnil" name:"LogTopicId"`
}

type UpdateServiceRequest struct {
	*tchttp.BaseRequest
	
	// 服务Id
	ServiceId *string `json:"ServiceId,omitnil" name:"ServiceId"`

	// 扩缩容配置
	Scaler *Scaler `json:"Scaler,omitnil" name:"Scaler"`

	// 服务配置Id
	ServiceConfigId *string `json:"ServiceConfigId,omitnil" name:"ServiceConfigId"`

	// 支持AUTO, MANUAL，分别表示自动扩缩容，手动扩缩容
	ScaleMode *string `json:"ScaleMode,omitnil" name:"ScaleMode"`

	// 支持STOP(停止) RESUME(重启)
	ServiceAction *string `json:"ServiceAction,omitnil" name:"ServiceAction"`

	// 备注
	Description *string `json:"Description,omitnil" name:"Description"`

	// GPU卡类型
	GpuType *string `json:"GpuType,omitnil" name:"GpuType"`

	// 处理器配置，单位为 1/1000 核
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存配置，单位为1M
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 显卡配置，单位为 1/1000 卡
	Gpu *uint64 `json:"Gpu,omitnil" name:"Gpu"`

	// Cls日志主题ID
	LogTopicId *string `json:"LogTopicId,omitnil" name:"LogTopicId"`
}

func (r *UpdateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Scaler")
	delete(f, "ServiceConfigId")
	delete(f, "ScaleMode")
	delete(f, "ServiceAction")
	delete(f, "Description")
	delete(f, "GpuType")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "Gpu")
	delete(f, "LogTopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateServiceResponseParams struct {
	// 服务
	Service *ModelService `json:"Service,omitnil" name:"Service"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateServiceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateServiceResponseParams `json:"Response"`
}

func (r *UpdateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}