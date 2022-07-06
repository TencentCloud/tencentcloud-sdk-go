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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CheckSavepointRequestParams struct {
	// 作业 id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 快照资源 id
	SerialId *string `json:"SerialId,omitempty" name:"SerialId"`

	// 快照类型 1: savepoint；2: checkpoint；3: cancelWithSavepoint
	RecordType *int64 `json:"RecordType,omitempty" name:"RecordType"`

	// 快照路径，目前只支持 cos 路径
	SavepointPath *string `json:"SavepointPath,omitempty" name:"SavepointPath"`

	// 工作空间 id
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
}

type CheckSavepointRequest struct {
	*tchttp.BaseRequest
	
	// 作业 id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 快照资源 id
	SerialId *string `json:"SerialId,omitempty" name:"SerialId"`

	// 快照类型 1: savepoint；2: checkpoint；3: cancelWithSavepoint
	RecordType *int64 `json:"RecordType,omitempty" name:"RecordType"`

	// 快照路径，目前只支持 cos 路径
	SavepointPath *string `json:"SavepointPath,omitempty" name:"SavepointPath"`

	// 工作空间 id
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
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
	SerialId *string `json:"SerialId,omitempty" name:"SerialId"`

	// 1=可用，2=不可用
	SavepointStatus *int64 `json:"SavepointStatus,omitempty" name:"SavepointStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CreateJobConfigRequestParams struct {
	// 作业Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 主类
	EntrypointClass *string `json:"EntrypointClass,omitempty" name:"EntrypointClass"`

	// 主类入参
	ProgramArgs *string `json:"ProgramArgs,omitempty" name:"ProgramArgs"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 资源引用数组
	ResourceRefs []*ResourceRef `json:"ResourceRefs,omitempty" name:"ResourceRefs"`

	// 作业默认并行度
	DefaultParallelism *uint64 `json:"DefaultParallelism,omitempty" name:"DefaultParallelism"`

	// 系统参数
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`

	// 1: 作业配置达到上限之后，自动删除可删除的最早版本
	AutoDelete *int64 `json:"AutoDelete,omitempty" name:"AutoDelete"`

	// 作业使用的 COS 存储桶名
	COSBucket *string `json:"COSBucket,omitempty" name:"COSBucket"`

	// 是否采集作业日志
	LogCollect *bool `json:"LogCollect,omitempty" name:"LogCollect"`

	// JobManager规格
	JobManagerSpec *float64 `json:"JobManagerSpec,omitempty" name:"JobManagerSpec"`

	// TaskManager规格
	TaskManagerSpec *float64 `json:"TaskManagerSpec,omitempty" name:"TaskManagerSpec"`

	// CLS日志集ID
	ClsLogsetId *string `json:"ClsLogsetId,omitempty" name:"ClsLogsetId"`

	// CLS日志主题ID
	ClsTopicId *string `json:"ClsTopicId,omitempty" name:"ClsTopicId"`

	// 日志采集类型 2：CLS；3：COS
	LogCollectType *int64 `json:"LogCollectType,omitempty" name:"LogCollectType"`

	// pyflink作业运行时使用的python版本
	PythonVersion *string `json:"PythonVersion,omitempty" name:"PythonVersion"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`

	// 日志级别
	LogLevel *string `json:"LogLevel,omitempty" name:"LogLevel"`
}

type CreateJobConfigRequest struct {
	*tchttp.BaseRequest
	
	// 作业Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 主类
	EntrypointClass *string `json:"EntrypointClass,omitempty" name:"EntrypointClass"`

	// 主类入参
	ProgramArgs *string `json:"ProgramArgs,omitempty" name:"ProgramArgs"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 资源引用数组
	ResourceRefs []*ResourceRef `json:"ResourceRefs,omitempty" name:"ResourceRefs"`

	// 作业默认并行度
	DefaultParallelism *uint64 `json:"DefaultParallelism,omitempty" name:"DefaultParallelism"`

	// 系统参数
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`

	// 1: 作业配置达到上限之后，自动删除可删除的最早版本
	AutoDelete *int64 `json:"AutoDelete,omitempty" name:"AutoDelete"`

	// 作业使用的 COS 存储桶名
	COSBucket *string `json:"COSBucket,omitempty" name:"COSBucket"`

	// 是否采集作业日志
	LogCollect *bool `json:"LogCollect,omitempty" name:"LogCollect"`

	// JobManager规格
	JobManagerSpec *float64 `json:"JobManagerSpec,omitempty" name:"JobManagerSpec"`

	// TaskManager规格
	TaskManagerSpec *float64 `json:"TaskManagerSpec,omitempty" name:"TaskManagerSpec"`

	// CLS日志集ID
	ClsLogsetId *string `json:"ClsLogsetId,omitempty" name:"ClsLogsetId"`

	// CLS日志主题ID
	ClsTopicId *string `json:"ClsTopicId,omitempty" name:"ClsTopicId"`

	// 日志采集类型 2：CLS；3：COS
	LogCollectType *int64 `json:"LogCollectType,omitempty" name:"LogCollectType"`

	// pyflink作业运行时使用的python版本
	PythonVersion *string `json:"PythonVersion,omitempty" name:"PythonVersion"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`

	// 日志级别
	LogLevel *string `json:"LogLevel,omitempty" name:"LogLevel"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateJobConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateJobConfigResponseParams struct {
	// 作业配置版本号
	Version *uint64 `json:"Version,omitempty" name:"Version"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 作业名称，允许输入长度小于等于50个字符的中文、英文、数字、-（横线）、_（下划线）、.（点），且符号必须半角字符。注意作业名不能和现有作业同名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 作业的类型，1 表示 SQL 作业，2 表示 JAR 作业
	JobType *int64 `json:"JobType,omitempty" name:"JobType"`

	// 集群的类型，1 表示共享集群，2 表示独享集群
	ClusterType *int64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// 当 ClusterType=2 时，必选，用来指定该作业提交的独享集群 ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 设置每 CU 的内存规格，单位为 GB，支持 2、4、8、16（需申请开通白名单后使用）。默认为 4，即 1 CU 对应 4 GB 的运行内存
	CuMem *uint64 `json:"CuMem,omitempty" name:"CuMem"`

	// 作业的备注信息，可以随意设置
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 作业名所属文件夹ID，根目录为"root"
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 作业运行的Flink版本
	FlinkVersion *string `json:"FlinkVersion,omitempty" name:"FlinkVersion"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
}

type CreateJobRequest struct {
	*tchttp.BaseRequest
	
	// 作业名称，允许输入长度小于等于50个字符的中文、英文、数字、-（横线）、_（下划线）、.（点），且符号必须半角字符。注意作业名不能和现有作业同名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 作业的类型，1 表示 SQL 作业，2 表示 JAR 作业
	JobType *int64 `json:"JobType,omitempty" name:"JobType"`

	// 集群的类型，1 表示共享集群，2 表示独享集群
	ClusterType *int64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// 当 ClusterType=2 时，必选，用来指定该作业提交的独享集群 ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 设置每 CU 的内存规格，单位为 GB，支持 2、4、8、16（需申请开通白名单后使用）。默认为 4，即 1 CU 对应 4 GB 的运行内存
	CuMem *uint64 `json:"CuMem,omitempty" name:"CuMem"`

	// 作业的备注信息，可以随意设置
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 作业名所属文件夹ID，根目录为"root"
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 作业运行的Flink版本
	FlinkVersion *string `json:"FlinkVersion,omitempty" name:"FlinkVersion"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateJobResponseParams struct {
	// 作业Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

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
type CreateResourceConfigRequestParams struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 位置信息
	ResourceLoc *ResourceLoc `json:"ResourceLoc,omitempty" name:"ResourceLoc"`

	// 资源描述信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 1： 资源版本达到上限，自动删除最早可删除的版本
	AutoDelete *int64 `json:"AutoDelete,omitempty" name:"AutoDelete"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
}

type CreateResourceConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 位置信息
	ResourceLoc *ResourceLoc `json:"ResourceLoc,omitempty" name:"ResourceLoc"`

	// 资源描述信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 1： 资源版本达到上限，自动删除最早可删除的版本
	AutoDelete *int64 `json:"AutoDelete,omitempty" name:"AutoDelete"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
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
	Version *int64 `json:"Version,omitempty" name:"Version"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ResourceLoc *ResourceLoc `json:"ResourceLoc,omitempty" name:"ResourceLoc"`

	// 资源类型。目前只支持 JAR，取值为 1
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源描述
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 资源名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 资源版本描述
	ResourceConfigRemark *string `json:"ResourceConfigRemark,omitempty" name:"ResourceConfigRemark"`

	// 目录ID
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
}

type CreateResourceRequest struct {
	*tchttp.BaseRequest
	
	// 资源位置
	ResourceLoc *ResourceLoc `json:"ResourceLoc,omitempty" name:"ResourceLoc"`

	// 资源类型。目前只支持 JAR，取值为 1
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源描述
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 资源名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 资源版本描述
	ResourceConfigRemark *string `json:"ResourceConfigRemark,omitempty" name:"ResourceConfigRemark"`

	// 目录ID
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
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
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源版本
	Version *int64 `json:"Version,omitempty" name:"Version"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeleteResourceConfigsRequestParams struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源版本数组
	ResourceConfigVersions []*int64 `json:"ResourceConfigVersions,omitempty" name:"ResourceConfigVersions"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
}

type DeleteResourceConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源版本数组
	ResourceConfigVersions []*int64 `json:"ResourceConfigVersions,omitempty" name:"ResourceConfigVersions"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
}

type DeleteResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 待删除资源ID列表
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 调试作业ID
	DebugId *int64 `json:"DebugId,omitempty" name:"DebugId"`

	// 表名
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
}

type DeleteTableConfigRequest struct {
	*tchttp.BaseRequest
	
	// 作业ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 调试作业ID
	DebugId *int64 `json:"DebugId,omitempty" name:"DebugId"`

	// 表名
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeJobConfigsRequestParams struct {
	// 作业Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 作业配置版本
	JobConfigVersions []*uint64 `json:"JobConfigVersions,omitempty" name:"JobConfigVersions"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认20，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// true 表示只展示草稿
	OnlyDraft *bool `json:"OnlyDraft,omitempty" name:"OnlyDraft"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
}

type DescribeJobConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 作业Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 作业配置版本
	JobConfigVersions []*uint64 `json:"JobConfigVersions,omitempty" name:"JobConfigVersions"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认20，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// true 表示只展示草稿
	OnlyDraft *bool `json:"OnlyDraft,omitempty" name:"OnlyDraft"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 作业配置列表
	JobConfigSet []*JobConfig `json:"JobConfigSet,omitempty" name:"JobConfigSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeJobSavepointRequestParams struct {
	// 作业 SerialId
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 分页参数，单页总数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
}

type DescribeJobSavepointRequest struct {
	*tchttp.BaseRequest
	
	// 作业 SerialId
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 分页参数，单页总数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobSavepointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobSavepointResponseParams struct {
	// 快照列表总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalNumber *int64 `json:"TotalNumber,omitempty" name:"TotalNumber"`

	// 快照列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Savepoint []*Savepoint `json:"Savepoint,omitempty" name:"Savepoint"`

	// 进行中的快照列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningSavepoint []*Savepoint `json:"RunningSavepoint,omitempty" name:"RunningSavepoint"`

	// 进行中的快照列表总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningTotalNumber *int64 `json:"RunningTotalNumber,omitempty" name:"RunningTotalNumber"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeJobsRequestParams struct {
	// 按照一个或者多个作业ID查询。作业ID形如：cql-11112222，每次请求的作业上限为100。参数不支持同时指定JobIds和Filters。
	JobIds []*string `json:"JobIds,omitempty" name:"JobIds"`

	// 过滤条件，支持的 Filter.Name 为：作业名 Name、作业状态 Status、所属集群 ClusterId、作业id JobId、集群名称 ClusterName。 每次请求的 Filters 个数的上限为 5，Filter.Values 的个数上限为 5。参数不支持同时指定 JobIds 和 Filters。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
}

type DescribeJobsRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个作业ID查询。作业ID形如：cql-11112222，每次请求的作业上限为100。参数不支持同时指定JobIds和Filters。
	JobIds []*string `json:"JobIds,omitempty" name:"JobIds"`

	// 过滤条件，支持的 Filter.Name 为：作业名 Name、作业状态 Status、所属集群 ClusterId、作业id JobId、集群名称 ClusterName。 每次请求的 Filters 个数的上限为 5，Filter.Values 的个数上限为 5。参数不支持同时指定 JobIds 和 Filters。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobsResponseParams struct {
	// 作业总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 作业列表
	JobSet []*JobV1 `json:"JobSet,omitempty" name:"JobSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 偏移量，仅当设置 Limit 时该参数有效
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回值大小，不填则返回全量数据
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源配置Versions集合
	ResourceConfigVersions []*int64 `json:"ResourceConfigVersions,omitempty" name:"ResourceConfigVersions"`

	// 作业配置版本
	JobConfigVersion *int64 `json:"JobConfigVersion,omitempty" name:"JobConfigVersion"`

	// 作业ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
}

type DescribeResourceConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 偏移量，仅当设置 Limit 时该参数有效
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回值大小，不填则返回全量数据
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源配置Versions集合
	ResourceConfigVersions []*int64 `json:"ResourceConfigVersions,omitempty" name:"ResourceConfigVersions"`

	// 作业配置版本
	JobConfigVersion *int64 `json:"JobConfigVersion,omitempty" name:"JobConfigVersion"`

	// 作业ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
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
	ResourceConfigSet []*ResourceConfigItem `json:"ResourceConfigSet,omitempty" name:"ResourceConfigSet"`

	// 资源配置数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 默认0;   1： 按照作业版本创建时间降序
	DESCByJobConfigCreateTime *int64 `json:"DESCByJobConfigCreateTime,omitempty" name:"DESCByJobConfigCreateTime"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源版本号
	ResourceConfigVersion *int64 `json:"ResourceConfigVersion,omitempty" name:"ResourceConfigVersion"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
}

type DescribeResourceRelatedJobsRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 默认0;   1： 按照作业版本创建时间降序
	DESCByJobConfigCreateTime *int64 `json:"DESCByJobConfigCreateTime,omitempty" name:"DESCByJobConfigCreateTime"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源版本号
	ResourceConfigVersion *int64 `json:"ResourceConfigVersion,omitempty" name:"ResourceConfigVersion"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 关联作业信息
	RefJobInfos []*ResourceRefJobInfo `json:"RefJobInfos,omitempty" name:"RefJobInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 偏移量，仅当设置 Limit 参数时有效
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 条数限制。如果不填，默认返回 20 条
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// <li><strong>ResourceName</strong></li>
	// <p style="padding-left: 30px;">按照资源名字过滤，支持模糊过滤。传入的过滤名字不超过5个</p><p style="padding-left: 30px;">类型: String</p><p style="padding-left: 30px;">必选: 否</p>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
}

type DescribeResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的资源ID数组，数量不超过100个。如果填写了该参数则忽略Filters参数。
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 偏移量，仅当设置 Limit 参数时有效
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 条数限制。如果不填，默认返回 20 条
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// <li><strong>ResourceName</strong></li>
	// <p style="padding-left: 30px;">按照资源名字过滤，支持模糊过滤。传入的过滤名字不超过5个</p><p style="padding-left: 30px;">类型: String</p><p style="padding-left: 30px;">必选: 否</p>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
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
	ResourceSet []*ResourceItem `json:"ResourceSet,omitempty" name:"ResourceSet"`

	// 总数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 偏移量，仅当设置 Limit 参数时有效
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 条数限制，默认返回 20 条
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询资源配置列表， 如果不填写，返回该 ResourceIds.N 下所有作业配置列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 查询对应Flink版本的内置connector
	FlinkVersion *string `json:"FlinkVersion,omitempty" name:"FlinkVersion"`
}

type DescribeSystemResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的资源ID数组
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 偏移量，仅当设置 Limit 参数时有效
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 条数限制，默认返回 20 条
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询资源配置列表， 如果不填写，返回该 ResourceIds.N 下所有作业配置列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 查询对应Flink版本的内置connector
	FlinkVersion *string `json:"FlinkVersion,omitempty" name:"FlinkVersion"`
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
	ResourceSet []*SystemResourceItem `json:"ResourceSet,omitempty" name:"ResourceSet"`

	// 总数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type Filter struct {
	// 要过滤的字段
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type JobConfig struct {
	// 作业Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 主类
	// 注意：此字段可能返回 null，表示取不到有效值。
	EntrypointClass *string `json:"EntrypointClass,omitempty" name:"EntrypointClass"`

	// 主类入参
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramArgs *string `json:"ProgramArgs,omitempty" name:"ProgramArgs"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 作业配置创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 作业配置的版本号
	Version *int64 `json:"Version,omitempty" name:"Version"`

	// 作业默认并行度
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultParallelism *uint64 `json:"DefaultParallelism,omitempty" name:"DefaultParallelism"`

	// 系统参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`

	// 引用资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceRefDetails []*ResourceRefDetail `json:"ResourceRefDetails,omitempty" name:"ResourceRefDetails"`

	// 创建者uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`

	// 作业配置上次启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 作业绑定的存储桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	COSBucket *string `json:"COSBucket,omitempty" name:"COSBucket"`

	// 是否启用日志收集，0-未启用，1-已启用，2-历史集群未设置日志集，3-历史集群已开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogCollect *int64 `json:"LogCollect,omitempty" name:"LogCollect"`

	// 作业的最大并行度
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxParallelism *uint64 `json:"MaxParallelism,omitempty" name:"MaxParallelism"`

	// JobManager规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobManagerSpec *float64 `json:"JobManagerSpec,omitempty" name:"JobManagerSpec"`

	// TaskManager规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskManagerSpec *float64 `json:"TaskManagerSpec,omitempty" name:"TaskManagerSpec"`

	// CLS日志集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClsLogsetId *string `json:"ClsLogsetId,omitempty" name:"ClsLogsetId"`

	// CLS日志主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClsTopicId *string `json:"ClsTopicId,omitempty" name:"ClsTopicId"`

	// pyflink作业运行的python版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PythonVersion *string `json:"PythonVersion,omitempty" name:"PythonVersion"`
}

type JobV1 struct {
	// 作业ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 用户AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 用户UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 创建者UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`

	// 作业名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 作业类型，1：sql作业，2：Jar作业
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobType *int64 `json:"JobType,omitempty" name:"JobType"`

	// 作业状态，1：未初始化，2：未发布，3：操作中，4：运行中，5：停止，6：暂停，-1：故障
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 作业创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 作业启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 作业停止时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StopTime *string `json:"StopTime,omitempty" name:"StopTime"`

	// 作业更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 作业累计运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalRunMillis *int64 `json:"TotalRunMillis,omitempty" name:"TotalRunMillis"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 操作错误提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOpResult *string `json:"LastOpResult,omitempty" name:"LastOpResult"`

	// 集群名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 最新配置版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestJobConfigVersion *int64 `json:"LatestJobConfigVersion,omitempty" name:"LatestJobConfigVersion"`

	// 已发布的配置版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublishedJobConfigVersion *int64 `json:"PublishedJobConfigVersion,omitempty" name:"PublishedJobConfigVersion"`

	// 运行的CU数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningCuNum *int64 `json:"RunningCuNum,omitempty" name:"RunningCuNum"`

	// 作业内存规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	CuMem *int64 `json:"CuMem,omitempty" name:"CuMem"`

	// 作业状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 运行状态时表示单次运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentRunMillis *int64 `json:"CurrentRunMillis,omitempty" name:"CurrentRunMillis"`

	// 作业所在的集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 作业管理WEB UI 入口
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebUIUrl *string `json:"WebUIUrl,omitempty" name:"WebUIUrl"`

	// 作业所在集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerType *int64 `json:"SchedulerType,omitempty" name:"SchedulerType"`

	// 作业所在集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterStatus *int64 `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// 细粒度下的运行的CU数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningCu *float64 `json:"RunningCu,omitempty" name:"RunningCu"`

	// 作业运行的 Flink 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlinkVersion *string `json:"FlinkVersion,omitempty" name:"FlinkVersion"`

	// 工作空间 SerialId
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`

	// 工作空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkSpaceName *string `json:"WorkSpaceName,omitempty" name:"WorkSpaceName"`
}

// Predefined struct for user
type ModifyJobRequestParams struct {
	// 作业Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 作业名称，支持长度小于50的中文/英文/数字/”-”/”_”/”.”，不能重名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 拖拽文件需传入此参数
	TargetFolderId *string `json:"TargetFolderId,omitempty" name:"TargetFolderId"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
}

type ModifyJobRequest struct {
	*tchttp.BaseRequest
	
	// 作业Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 作业名称，支持长度小于50的中文/英文/数字/”-”/”_”/”.”，不能重名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 拖拽文件需传入此参数
	TargetFolderId *string `json:"TargetFolderId,omitempty" name:"TargetFolderId"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type Property struct {
	// 系统配置的Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 系统配置的Value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ResourceConfigItem struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源类型
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源所属地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 资源所属AppId
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 主账号Uin
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 子账号Uin
	CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`

	// 资源位置描述
	ResourceLoc *ResourceLoc `json:"ResourceLoc,omitempty" name:"ResourceLoc"`

	// 资源创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 资源版本
	Version *int64 `json:"Version,omitempty" name:"Version"`

	// 资源描述
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 资源状态：0: 资源同步中，1:资源已就绪
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 关联作业个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefJobCount *int64 `json:"RefJobCount,omitempty" name:"RefJobCount"`
}

type ResourceItem struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 资源类型
	ResourceType *uint64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源位置
	ResourceLoc *ResourceLoc `json:"ResourceLoc,omitempty" name:"ResourceLoc"`

	// 资源地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 应用ID
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 主账号Uin
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 子账号Uin
	CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`

	// 资源创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 资源最后更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 资源的资源版本ID
	LatestResourceConfigVersion *int64 `json:"LatestResourceConfigVersion,omitempty" name:"LatestResourceConfigVersion"`

	// 资源备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 版本个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionCount *int64 `json:"VersionCount,omitempty" name:"VersionCount"`

	// 关联作业数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefJobCount *int64 `json:"RefJobCount,omitempty" name:"RefJobCount"`
}

type ResourceLoc struct {
	// 资源位置的存储类型，目前只支持1:COS
	StorageType *int64 `json:"StorageType,omitempty" name:"StorageType"`

	// 描述资源位置的json
	Param *ResourceLocParam `json:"Param,omitempty" name:"Param"`
}

type ResourceLocParam struct {
	// 资源bucket
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 资源路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 资源所在地域，如果不填，则使用Resource的Region
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`
}

type ResourceRef struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源版本ID，-1表示使用最新版本
	Version *int64 `json:"Version,omitempty" name:"Version"`

	// 引用资源类型，例如主资源设置为1，代表main class所在的jar包
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type ResourceRefDetail struct {
	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源版本，-1表示使用最新版本
	Version *int64 `json:"Version,omitempty" name:"Version"`

	// 资源名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 1: 主资源
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 1: 系统内置资源
	SystemProvide *int64 `json:"SystemProvide,omitempty" name:"SystemProvide"`
}

type ResourceRefJobInfo struct {
	// Job id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Job配置版本
	JobConfigVersion *int64 `json:"JobConfigVersion,omitempty" name:"JobConfigVersion"`

	// 资源版本
	ResourceVersion *int64 `json:"ResourceVersion,omitempty" name:"ResourceVersion"`
}

type RunJobDescription struct {
	// 作业Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 运行类型，1：启动，2：恢复
	RunType *int64 `json:"RunType,omitempty" name:"RunType"`

	// 已废弃。旧版 SQL 类型作业启动参数：指定数据源消费起始时间点
	StartMode *string `json:"StartMode,omitempty" name:"StartMode"`

	// 当前作业的某个版本
	JobConfigVersion *uint64 `json:"JobConfigVersion,omitempty" name:"JobConfigVersion"`

	// Savepoint路径
	SavepointPath *string `json:"SavepointPath,omitempty" name:"SavepointPath"`

	// Savepoint的Id
	SavepointId *string `json:"SavepointId,omitempty" name:"SavepointId"`
}

// Predefined struct for user
type RunJobsRequestParams struct {
	// 批量启动作业的描述信息
	RunJobDescriptions []*RunJobDescription `json:"RunJobDescriptions,omitempty" name:"RunJobDescriptions"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
}

type RunJobsRequest struct {
	*tchttp.BaseRequest
	
	// 批量启动作业的描述信息
	RunJobDescriptions []*RunJobDescription `json:"RunJobDescriptions,omitempty" name:"RunJobDescriptions"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type Savepoint struct {
	// 主键
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`

	// 状态 1: Active; 2: Expired; 3: InProgress; 4: Failed; 5: Timeout
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 快照类型 1: savepoint；2: checkpoint；3: cancelWithSavepoint
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordType *int64 `json:"RecordType,omitempty" name:"RecordType"`

	// 运行作业实例的顺序 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobRuntimeId *int64 `json:"JobRuntimeId,omitempty" name:"JobRuntimeId"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 固定超时时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`

	// 快照 serialId
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerialId *string `json:"SerialId,omitempty" name:"SerialId"`
}

type StopJobDescription struct {
	// 作业Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 停止类型，1 停止 2 暂停
	StopType *int64 `json:"StopType,omitempty" name:"StopType"`
}

// Predefined struct for user
type StopJobsRequestParams struct {
	// 批量停止作业的描述信息
	StopJobDescriptions []*StopJobDescription `json:"StopJobDescriptions,omitempty" name:"StopJobDescriptions"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
}

type StopJobsRequest struct {
	*tchttp.BaseRequest
	
	// 批量停止作业的描述信息
	StopJobDescriptions []*StopJobDescription `json:"StopJobDescriptions,omitempty" name:"StopJobDescriptions"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type SystemResourceItem struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 资源类型。1 表示 JAR 包，目前只支持该值。
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 资源所属地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 资源的最新版本
	LatestResourceConfigVersion *int64 `json:"LatestResourceConfigVersion,omitempty" name:"LatestResourceConfigVersion"`
}

// Predefined struct for user
type TriggerJobSavepointRequestParams struct {
	// 作业 SerialId
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
}

type TriggerJobSavepointRequest struct {
	*tchttp.BaseRequest
	
	// 作业 SerialId
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 工作空间 SerialId
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" name:"WorkSpaceId"`
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
	SavepointTrigger *bool `json:"SavepointTrigger,omitempty" name:"SavepointTrigger"`

	// 错误消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// 快照路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinalSavepointPath *string `json:"FinalSavepointPath,omitempty" name:"FinalSavepointPath"`

	// 快照 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SavepointId *string `json:"SavepointId,omitempty" name:"SavepointId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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