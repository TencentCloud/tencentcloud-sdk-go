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

package v20180226

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CreateJobRequestParams struct {
	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 运行任务的集群，详见 [使用集群](https://cloud.tencent.com/document/product/851/17317)
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 运行任务的环境，详见 [运行环境](https://cloud.tencent.com/document/product/851/17320)
	RuntimeVersion *string `json:"RuntimeVersion,omitempty" name:"RuntimeVersion"`

	// 挂载的路径，支持 NFS，[CFS](https://cloud.tencent.com/product/cfs) 和 [COS](https://cloud.tencent.com/product/cos)，其中 COS 只在 [TI-A 定制环境](https://cloud.tencent.com/document/product/851/17320#ti-a-.E5.AE.9A.E5.88.B6.E7.8E.AF.E5.A2.83) 中支持
	PackageDir []*string `json:"PackageDir,omitempty" name:"PackageDir"`

	// 任务启动命令
	Command []*string `json:"Command,omitempty" name:"Command"`

	// 任务启动参数
	Args []*string `json:"Args,omitempty" name:"Args"`

	// 运行任务的配置信息，详见 [训练规模](https://cloud.tencent.com/document/product/851/17319)
	ScaleTier *string `json:"ScaleTier,omitempty" name:"ScaleTier"`

	// Master 机器类型，ScaleTier 取值为 `CUSTOM` 时必填，详见 [训练规模](https://cloud.tencent.com/document/product/851/17319)
	MasterType *string `json:"MasterType,omitempty" name:"MasterType"`

	// Worker 机器类型，ScaleTier 取值为 `CUSTOM` 时必填，详见 [训练规模](https://cloud.tencent.com/document/product/851/17319)
	WorkerType *string `json:"WorkerType,omitempty" name:"WorkerType"`

	// Parameter server 机器类型，ScaleTier 取值为 `CUSTOM` 时必填,详见 [训练规模](https://cloud.tencent.com/document/product/851/17319)
	ParameterServerType *string `json:"ParameterServerType,omitempty" name:"ParameterServerType"`

	// Worker 机器数量，ScaleTier 取值为 `CUSTOM` 时必填,详见 [训练规模](https://cloud.tencent.com/document/product/851/17319)
	WorkerCount *uint64 `json:"WorkerCount,omitempty" name:"WorkerCount"`

	// Parameter server 机器数量，ScaleTier 取值为 `CUSTOM` 时必填,详见 [训练规模](https://cloud.tencent.com/document/product/851/17319)
	ParameterServerCount *uint64 `json:"ParameterServerCount,omitempty" name:"ParameterServerCount"`

	// 启动 debug 模式，默认为 false
	Debug *bool `json:"Debug,omitempty" name:"Debug"`

	// 运行任务的其他配置信息
	RuntimeConf []*string `json:"RuntimeConf,omitempty" name:"RuntimeConf"`
}

type CreateJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 运行任务的集群，详见 [使用集群](https://cloud.tencent.com/document/product/851/17317)
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 运行任务的环境，详见 [运行环境](https://cloud.tencent.com/document/product/851/17320)
	RuntimeVersion *string `json:"RuntimeVersion,omitempty" name:"RuntimeVersion"`

	// 挂载的路径，支持 NFS，[CFS](https://cloud.tencent.com/product/cfs) 和 [COS](https://cloud.tencent.com/product/cos)，其中 COS 只在 [TI-A 定制环境](https://cloud.tencent.com/document/product/851/17320#ti-a-.E5.AE.9A.E5.88.B6.E7.8E.AF.E5.A2.83) 中支持
	PackageDir []*string `json:"PackageDir,omitempty" name:"PackageDir"`

	// 任务启动命令
	Command []*string `json:"Command,omitempty" name:"Command"`

	// 任务启动参数
	Args []*string `json:"Args,omitempty" name:"Args"`

	// 运行任务的配置信息，详见 [训练规模](https://cloud.tencent.com/document/product/851/17319)
	ScaleTier *string `json:"ScaleTier,omitempty" name:"ScaleTier"`

	// Master 机器类型，ScaleTier 取值为 `CUSTOM` 时必填，详见 [训练规模](https://cloud.tencent.com/document/product/851/17319)
	MasterType *string `json:"MasterType,omitempty" name:"MasterType"`

	// Worker 机器类型，ScaleTier 取值为 `CUSTOM` 时必填，详见 [训练规模](https://cloud.tencent.com/document/product/851/17319)
	WorkerType *string `json:"WorkerType,omitempty" name:"WorkerType"`

	// Parameter server 机器类型，ScaleTier 取值为 `CUSTOM` 时必填,详见 [训练规模](https://cloud.tencent.com/document/product/851/17319)
	ParameterServerType *string `json:"ParameterServerType,omitempty" name:"ParameterServerType"`

	// Worker 机器数量，ScaleTier 取值为 `CUSTOM` 时必填,详见 [训练规模](https://cloud.tencent.com/document/product/851/17319)
	WorkerCount *uint64 `json:"WorkerCount,omitempty" name:"WorkerCount"`

	// Parameter server 机器数量，ScaleTier 取值为 `CUSTOM` 时必填,详见 [训练规模](https://cloud.tencent.com/document/product/851/17319)
	ParameterServerCount *uint64 `json:"ParameterServerCount,omitempty" name:"ParameterServerCount"`

	// 启动 debug 模式，默认为 false
	Debug *bool `json:"Debug,omitempty" name:"Debug"`

	// 运行任务的其他配置信息
	RuntimeConf []*string `json:"RuntimeConf,omitempty" name:"RuntimeConf"`
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
	delete(f, "Cluster")
	delete(f, "RuntimeVersion")
	delete(f, "PackageDir")
	delete(f, "Command")
	delete(f, "Args")
	delete(f, "ScaleTier")
	delete(f, "MasterType")
	delete(f, "WorkerType")
	delete(f, "ParameterServerType")
	delete(f, "WorkerCount")
	delete(f, "ParameterServerCount")
	delete(f, "Debug")
	delete(f, "RuntimeConf")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateJobResponseParams struct {
	// 训练任务信息
	Job *Job `json:"Job,omitempty" name:"Job"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateModelRequestParams struct {
	// 模型名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 要部署的模型文件路径名
	Model *string `json:"Model,omitempty" name:"Model"`

	// 关于模型的描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 部署目标集群的名称，`集群模式` 必填
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 运行环境镜像的标签，详见 [Serving 环境](https://cloud.tencent.com/document/product/851/17320#serving-.E7.8E.AF.E5.A2.83)
	RuntimeVersion *string `json:"RuntimeVersion,omitempty" name:"RuntimeVersion"`

	// 要部署的模型副本数目，`集群模式` 选填
	Replicas *uint64 `json:"Replicas,omitempty" name:"Replicas"`

	// 暴露外网或内网，默认暴露外网，`集群模式` 选填
	Expose *string `json:"Expose,omitempty" name:"Expose"`

	// 部署模式，取值 `serverless` 即为 `无服务器模式`，否则为 `集群模式` 下服务的运行规模，形如 `2U4G1P`，详见 [自定义的训练规模](https://cloud.tencent.com/document/product/851/17319#.E8.87.AA.E5.AE.9A.E4.B9.89.E7.9A.84.E8.AE.AD.E7.BB.83.E8.A7.84.E6.A8.A1)
	ServType *string `json:"ServType,omitempty" name:"ServType"`

	// `无服务器模式` 可选的其他配置信息，详见 [利用无服务器函数部署](https://cloud.tencent.com/document/product/851/17049#.E5.88.A9.E7.94.A8.E6.97.A0.E6.9C.8D.E5.8A.A1.E5.99.A8.E5.87.BD.E6.95.B0.E9.83.A8.E7.BD.B2)
	RuntimeConf []*string `json:"RuntimeConf,omitempty" name:"RuntimeConf"`
}

type CreateModelRequest struct {
	*tchttp.BaseRequest
	
	// 模型名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 要部署的模型文件路径名
	Model *string `json:"Model,omitempty" name:"Model"`

	// 关于模型的描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 部署目标集群的名称，`集群模式` 必填
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 运行环境镜像的标签，详见 [Serving 环境](https://cloud.tencent.com/document/product/851/17320#serving-.E7.8E.AF.E5.A2.83)
	RuntimeVersion *string `json:"RuntimeVersion,omitempty" name:"RuntimeVersion"`

	// 要部署的模型副本数目，`集群模式` 选填
	Replicas *uint64 `json:"Replicas,omitempty" name:"Replicas"`

	// 暴露外网或内网，默认暴露外网，`集群模式` 选填
	Expose *string `json:"Expose,omitempty" name:"Expose"`

	// 部署模式，取值 `serverless` 即为 `无服务器模式`，否则为 `集群模式` 下服务的运行规模，形如 `2U4G1P`，详见 [自定义的训练规模](https://cloud.tencent.com/document/product/851/17319#.E8.87.AA.E5.AE.9A.E4.B9.89.E7.9A.84.E8.AE.AD.E7.BB.83.E8.A7.84.E6.A8.A1)
	ServType *string `json:"ServType,omitempty" name:"ServType"`

	// `无服务器模式` 可选的其他配置信息，详见 [利用无服务器函数部署](https://cloud.tencent.com/document/product/851/17049#.E5.88.A9.E7.94.A8.E6.97.A0.E6.9C.8D.E5.8A.A1.E5.99.A8.E5.87.BD.E6.95.B0.E9.83.A8.E7.BD.B2)
	RuntimeConf []*string `json:"RuntimeConf,omitempty" name:"RuntimeConf"`
}

func (r *CreateModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Model")
	delete(f, "Description")
	delete(f, "Cluster")
	delete(f, "RuntimeVersion")
	delete(f, "Replicas")
	delete(f, "Expose")
	delete(f, "ServType")
	delete(f, "RuntimeConf")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelResponseParams struct {
	// 模型的详细信息
	Model *Model `json:"Model,omitempty" name:"Model"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateModelResponse struct {
	*tchttp.BaseResponse
	Response *CreateModelResponseParams `json:"Response"`
}

func (r *CreateModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteJobRequestParams struct {
	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 运行任务的集群
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

type DeleteJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 运行任务的集群
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
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
	delete(f, "Name")
	delete(f, "Cluster")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeleteModelRequestParams struct {
	// 要删除的模型名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 要删除的模型所在的集群名称，`集群模式` 必填
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 模型类型，取值 `serverless` 即为 `无服务器模式`，否则为 `集群模式`
	ServType *string `json:"ServType,omitempty" name:"ServType"`
}

type DeleteModelRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的模型名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 要删除的模型所在的集群名称，`集群模式` 必填
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 模型类型，取值 `serverless` 即为 `无服务器模式`，否则为 `集群模式`
	ServType *string `json:"ServType,omitempty" name:"ServType"`
}

func (r *DeleteModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Cluster")
	delete(f, "ServType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteModelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteModelResponseParams `json:"Response"`
}

func (r *DeleteModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobRequestParams struct {
	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 运行任务的集群
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

type DescribeJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 运行任务的集群
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *DescribeJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Cluster")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobResponseParams struct {
	// 训练任务信息
	Job *Job `json:"Job,omitempty" name:"Job"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobResponseParams `json:"Response"`
}

func (r *DescribeJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelRequestParams struct {
	// 模型名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模型所在集群名称，`集群模式` 必填
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 模型类型，取值 `serverless` 即为 `无服务器模式`，否则为 `集群模式`
	ServType *string `json:"ServType,omitempty" name:"ServType"`
}

type DescribeModelRequest struct {
	*tchttp.BaseRequest
	
	// 模型名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模型所在集群名称，`集群模式` 必填
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 模型类型，取值 `serverless` 即为 `无服务器模式`，否则为 `集群模式`
	ServType *string `json:"ServType,omitempty" name:"ServType"`
}

func (r *DescribeModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Cluster")
	delete(f, "ServType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelResponseParams struct {
	// 模型信息
	Model *Model `json:"Model,omitempty" name:"Model"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeModelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelResponseParams `json:"Response"`
}

func (r *DescribeModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstallAgentRequestParams struct {
	// 集群名称
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// Agent版本, 用于私有集群的agent安装，默认为“private-training”
	TiaVersion *string `json:"TiaVersion,omitempty" name:"TiaVersion"`

	// 是否允许更新Agent
	Update *bool `json:"Update,omitempty" name:"Update"`
}

type InstallAgentRequest struct {
	*tchttp.BaseRequest
	
	// 集群名称
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// Agent版本, 用于私有集群的agent安装，默认为“private-training”
	TiaVersion *string `json:"TiaVersion,omitempty" name:"TiaVersion"`

	// 是否允许更新Agent
	Update *bool `json:"Update,omitempty" name:"Update"`
}

func (r *InstallAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstallAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cluster")
	delete(f, "TiaVersion")
	delete(f, "Update")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InstallAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstallAgentResponseParams struct {
	// Agent版本, 用于私有集群的agent安装
	TiaVersion *string `json:"TiaVersion,omitempty" name:"TiaVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InstallAgentResponse struct {
	*tchttp.BaseResponse
	Response *InstallAgentResponseParams `json:"Response"`
}

func (r *InstallAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstallAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Job struct {
	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务创建时间，格式为：2006-01-02 15:04:05.999999999 -0700 MST
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务开始时间，格式为：2006-01-02 15:04:05.999999999 -0700 MST
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务结束时间，格式为：2006-01-02 15:04:05.999999999 -0700 MST
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务状态，可能的状态为Created（已创建），Running（运行中），Succeeded（运行完成：成功），Failed（运行完成：失败）
	State *string `json:"State,omitempty" name:"State"`

	// 任务状态信息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 运行任务的配置信息
	ScaleTier *string `json:"ScaleTier,omitempty" name:"ScaleTier"`

	// （ScaleTier为Custom时）master机器类型
	MasterType *string `json:"MasterType,omitempty" name:"MasterType"`

	// （ScaleTier为Custom时）worker机器类型
	WorkerType *string `json:"WorkerType,omitempty" name:"WorkerType"`

	// （ScaleTier为Custom时）parameter server机器类型
	ParameterServerType *string `json:"ParameterServerType,omitempty" name:"ParameterServerType"`

	// （ScaleTier为Custom时）worker机器数量
	WorkerCount *uint64 `json:"WorkerCount,omitempty" name:"WorkerCount"`

	// （ScaleTier为Custom时）parameter server机器数量
	ParameterServerCount *uint64 `json:"ParameterServerCount,omitempty" name:"ParameterServerCount"`

	// 挂载的路径
	PackageDir []*string `json:"PackageDir,omitempty" name:"PackageDir"`

	// 任务启动命令
	Command []*string `json:"Command,omitempty" name:"Command"`

	// 任务启动参数
	Args []*string `json:"Args,omitempty" name:"Args"`

	// 运行任务的集群
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 运行任务的环境
	RuntimeVersion *string `json:"RuntimeVersion,omitempty" name:"RuntimeVersion"`

	// 任务删除时间，格式为：2006-01-02 15:04:05.999999999 -0700 MST
	DelTime *string `json:"DelTime,omitempty" name:"DelTime"`

	// 创建任务的AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 创建任务的Uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 创建任务的Debug模式
	Debug *bool `json:"Debug,omitempty" name:"Debug"`

	// Runtime的额外配置信息
	RuntimeConf []*string `json:"RuntimeConf,omitempty" name:"RuntimeConf"`

	// 任务Id
	Id *string `json:"Id,omitempty" name:"Id"`
}

// Predefined struct for user
type ListJobsRequestParams struct {
	// 运行任务的集群
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 分页参数，返回数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，起始位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type ListJobsRequest struct {
	*tchttp.BaseRequest
	
	// 运行任务的集群
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 分页参数，返回数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，起始位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cluster")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListJobsResponseParams struct {
	// 训练任务列表
	Jobs []*Job `json:"Jobs,omitempty" name:"Jobs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListJobsResponse struct {
	*tchttp.BaseResponse
	Response *ListJobsResponseParams `json:"Response"`
}

func (r *ListJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListModelsRequestParams struct {
	// 部署模型的集群， `集群模式` 必填
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 分页参数，返回数量上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，分页起始位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 部署类型，取值 `serverless` 即为 `无服务器模式`，否则为 `集群模式`。
	ServType *string `json:"ServType,omitempty" name:"ServType"`
}

type ListModelsRequest struct {
	*tchttp.BaseRequest
	
	// 部署模型的集群， `集群模式` 必填
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 分页参数，返回数量上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，分页起始位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 部署类型，取值 `serverless` 即为 `无服务器模式`，否则为 `集群模式`。
	ServType *string `json:"ServType,omitempty" name:"ServType"`
}

func (r *ListModelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListModelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cluster")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ServType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListModelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListModelsResponseParams struct {
	// Model 数组，用以显示所有模型的信息
	Models []*Model `json:"Models,omitempty" name:"Models"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListModelsResponse struct {
	*tchttp.BaseResponse
	Response *ListModelsResponseParams `json:"Response"`
}

func (r *ListModelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListModelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Log struct {
	// 容器名
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 日志内容
	Log *string `json:"Log,omitempty" name:"Log"`

	// 空间名
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Pod Id
	PodId *string `json:"PodId,omitempty" name:"PodId"`

	// Pod名
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 日志日期，格式为“2018-07-02T09:10:04.916553368Z”
	Time *string `json:"Time,omitempty" name:"Time"`
}

type Model struct {
	// 模型名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模型描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 集群名称
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 模型地址
	Model *string `json:"Model,omitempty" name:"Model"`

	// 运行环境编号
	RuntimeVersion *string `json:"RuntimeVersion,omitempty" name:"RuntimeVersion"`

	// 模型创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 模型运行状态
	State *string `json:"State,omitempty" name:"State"`

	// 提供服务的url
	ServingUrl *string `json:"ServingUrl,omitempty" name:"ServingUrl"`

	// 相关消息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 编号
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 机型
	ServType *string `json:"ServType,omitempty" name:"ServType"`

	// 模型暴露方式
	Expose *string `json:"Expose,omitempty" name:"Expose"`

	// 部署副本数量
	Replicas *uint64 `json:"Replicas,omitempty" name:"Replicas"`

	// 模型Id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 创建任务的Uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 模型删除时间，格式为：2006-01-02 15:04:05.999999999 -0700 MST
	DelTime *string `json:"DelTime,omitempty" name:"DelTime"`
}

// Predefined struct for user
type QueryLogsRequestParams struct {
	// 任务的名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 任务所在集群的名称
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 查询日志的开始时间，格式：2019-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询日志的结束时间，格式：2019-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 单次要返回的日志条数上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 加载更多日志时使用，透传上次返回的 Context 值，获取后续的日志内容；使用 Context 翻页最多能获取 10000 条日志
	Context *string `json:"Context,omitempty" name:"Context"`
}

type QueryLogsRequest struct {
	*tchttp.BaseRequest
	
	// 任务的名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 任务所在集群的名称
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 查询日志的开始时间，格式：2019-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询日志的结束时间，格式：2019-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 单次要返回的日志条数上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 加载更多日志时使用，透传上次返回的 Context 值，获取后续的日志内容；使用 Context 翻页最多能获取 10000 条日志
	Context *string `json:"Context,omitempty" name:"Context"`
}

func (r *QueryLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobName")
	delete(f, "Cluster")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Context")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryLogsResponseParams struct {
	// 日志查询上下文，用于加载更多日志
	Context *string `json:"Context,omitempty" name:"Context"`

	// 日志内容列表
	Logs []*Log `json:"Logs,omitempty" name:"Logs"`

	// 是否已经返回所有符合条件的日志
	Listover *bool `json:"Listover,omitempty" name:"Listover"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryLogsResponse struct {
	*tchttp.BaseResponse
	Response *QueryLogsResponseParams `json:"Response"`
}

func (r *QueryLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}