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

package v20221128

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CVMOption struct {
	// 云服务器可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 云服务器实例规格。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

type ClusterOption struct {
	// 计算集群可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 计算集群类型，取值范围：
	// - KUBERNETES
	Type *string `json:"Type,omitempty" name:"Type"`
}

// Predefined struct for user
type CreateEnvironmentRequestParams struct {
	// 环境名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 环境配置信息。
	Config *EnvironmentConfig `json:"Config,omitempty" name:"Config"`

	// 环境描述。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// 环境名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 环境配置信息。
	Config *EnvironmentConfig `json:"Config,omitempty" name:"Config"`

	// 环境描述。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Config")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnvironmentResponseParams struct {
	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 工作流UUID。
	WorkflowUuid *string `json:"WorkflowUuid,omitempty" name:"WorkflowUuid"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *CreateEnvironmentResponseParams `json:"Response"`
}

func (r *CreateEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabaseOption struct {
	// 数据库可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

// Predefined struct for user
type DeleteEnvironmentRequestParams struct {
	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

type DeleteEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

func (r *DeleteEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEnvironmentResponseParams struct {
	// 工作流UUID。
	WorkflowUuid *string `json:"WorkflowUuid,omitempty" name:"WorkflowUuid"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEnvironmentResponseParams `json:"Response"`
}

func (r *DeleteEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsRequestParams struct {
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤器，支持过滤字段：
	// - EnvironmentId：环境ID
	// - Name：名称
	// - Status：环境状态
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤器，支持过滤字段：
	// - EnvironmentId：环境ID
	// - Name：名称
	// - Status：环境状态
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsResponseParams struct {
	// 符合条件的数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 环境详情列表。
	Environments []*Environment `json:"Environments,omitempty" name:"Environments"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvironmentsResponseParams `json:"Response"`
}

func (r *DescribeEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRunGroupsRequestParams struct {
	// 项目ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - Name：任务批次名称
	// - RunGroupId：任务批次ID
	// - Status：任务批次状态
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeRunGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - Name：任务批次名称
	// - RunGroupId：任务批次ID
	// - Status：任务批次状态
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRunGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRunGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRunGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRunGroupsResponseParams struct {
	// 符合条件的数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务批次列表。
	RunGroups []*RunGroup `json:"RunGroups,omitempty" name:"RunGroups"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRunGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRunGroupsResponseParams `json:"Response"`
}

func (r *DescribeRunGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRunGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRunsRequestParams struct {
	// 项目ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - RunGroupId：任务批次ID
	// - Status：任务状态
	// - RunUuid：任务UUID
	// - UserDefinedId：用户定义ID
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeRunsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - RunGroupId：任务批次ID
	// - Status：任务状态
	// - RunUuid：任务UUID
	// - UserDefinedId：用户定义ID
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRunsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRunsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRunsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRunsResponseParams struct {
	// 符合条件的数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务列表。
	Runs []*Run `json:"Runs,omitempty" name:"Runs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRunsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRunsResponseParams `json:"Response"`
}

func (r *DescribeRunsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRunsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Environment struct {
	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 环境名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 环境描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 环境地域。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 环境类型，取值范围：
	// - KUBERNETES：Kubernetes容器集群
	// - HPC：HPC高性能计算集群
	Type *string `json:"Type,omitempty" name:"Type"`

	// 环境状态，取值范围：
	// - INITIALIZING：创建中
	// - INITIALIZATION_ERROR：创建失败
	// - RUNNING：运行中
	// - ERROR：异常
	// - DELETING：正在删除
	// - DELETE_ERROR：删除失败
	Status *string `json:"Status,omitempty" name:"Status"`

	// 环境是否可用。环境需要可用才能投递计算任务。
	Available *bool `json:"Available,omitempty" name:"Available"`

	// 环境信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 云资源ID。
	ResourceIds *ResourceIds `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 上个工作流UUID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastWorkflowUuid *string `json:"LastWorkflowUuid,omitempty" name:"LastWorkflowUuid"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
}

type EnvironmentConfig struct {
	// 私有网络配置。
	VPCOption *VPCOption `json:"VPCOption,omitempty" name:"VPCOption"`

	// 计算集群配置。
	ClusterOption *ClusterOption `json:"ClusterOption,omitempty" name:"ClusterOption"`

	// 数据库配置。
	DatabaseOption *DatabaseOption `json:"DatabaseOption,omitempty" name:"DatabaseOption"`

	// 存储配置。
	StorageOption *StorageOption `json:"StorageOption,omitempty" name:"StorageOption"`

	// 云服务器配置。
	CVMOption *CVMOption `json:"CVMOption,omitempty" name:"CVMOption"`
}

type ExecutionTime struct {
	// 提交时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`

	// 开始时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type Filter struct {
	// 过滤字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤字段值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

// Predefined struct for user
type GetRunCallsRequestParams struct {
	// 任务Uuid。
	RunUuid *string `json:"RunUuid,omitempty" name:"RunUuid"`

	// 项目ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 作业路径
	Path *string `json:"Path,omitempty" name:"Path"`
}

type GetRunCallsRequest struct {
	*tchttp.BaseRequest
	
	// 任务Uuid。
	RunUuid *string `json:"RunUuid,omitempty" name:"RunUuid"`

	// 项目ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 作业路径
	Path *string `json:"Path,omitempty" name:"Path"`
}

func (r *GetRunCallsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRunCallsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RunUuid")
	delete(f, "ProjectId")
	delete(f, "Path")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRunCallsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRunCallsResponseParams struct {
	// 作业详情。
	Calls []*RunMetadata `json:"Calls,omitempty" name:"Calls"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetRunCallsResponse struct {
	*tchttp.BaseResponse
	Response *GetRunCallsResponseParams `json:"Response"`
}

func (r *GetRunCallsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRunCallsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRunStatusRequestParams struct {
	// 任务Uuid。
	RunUuid *string `json:"RunUuid,omitempty" name:"RunUuid"`

	// 项目ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type GetRunStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务Uuid。
	RunUuid *string `json:"RunUuid,omitempty" name:"RunUuid"`

	// 项目ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *GetRunStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRunStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RunUuid")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRunStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRunStatusResponseParams struct {
	// 作业详情。
	Metadata *RunMetadata `json:"Metadata,omitempty" name:"Metadata"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetRunStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetRunStatusResponseParams `json:"Response"`
}

func (r *GetRunStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRunStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportTableFileRequestParams struct {
	// 表格关联的项目ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 表格名称，支持20个字符内的英文字符、数字和下划线。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 表格文件Cos对象路径。
	CosUri *string `json:"CosUri,omitempty" name:"CosUri"`

	// 表格文件中每列的数据类型，支持的类型包括：Int、String、File、Array[File]
	DataType []*string `json:"DataType,omitempty" name:"DataType"`

	// 表格描述。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ImportTableFileRequest struct {
	*tchttp.BaseRequest
	
	// 表格关联的项目ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 表格名称，支持20个字符内的英文字符、数字和下划线。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 表格文件Cos对象路径。
	CosUri *string `json:"CosUri,omitempty" name:"CosUri"`

	// 表格文件中每列的数据类型，支持的类型包括：Int、String、File、Array[File]
	DataType []*string `json:"DataType,omitempty" name:"DataType"`

	// 表格描述。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ImportTableFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportTableFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Name")
	delete(f, "CosUri")
	delete(f, "DataType")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportTableFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportTableFileResponseParams struct {
	// 表格ID。
	TableId *string `json:"TableId,omitempty" name:"TableId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImportTableFileResponse struct {
	*tchttp.BaseResponse
	Response *ImportTableFileResponseParams `json:"Response"`
}

func (r *ImportTableFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportTableFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceIds struct {
	// 私有网络ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VPCId *string `json:"VPCId,omitempty" name:"VPCId"`

	// 子网ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 安全组ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// TDSQL-C Mysql版数据库ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TDSQLCId *string `json:"TDSQLCId,omitempty" name:"TDSQLCId"`

	// 文件存储ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFSId *string `json:"CFSId,omitempty" name:"CFSId"`

	// 文件存储类型：取值范围：
	// - SD：通用标准型
	// - HP：通用性能型
	// - TB：turbo标准型
	// - TP：turbo性能型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFSStorageType *string `json:"CFSStorageType,omitempty" name:"CFSStorageType"`

	// 云服务器ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVMId *string `json:"CVMId,omitempty" name:"CVMId"`

	// 弹性容器集群ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EKSId *string `json:"EKSId,omitempty" name:"EKSId"`
}

type Run struct {
	// 任务UUID。
	RunUuid *string `json:"RunUuid,omitempty" name:"RunUuid"`

	// 项目ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 应用ID。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 任务批次ID。
	RunGroupId *string `json:"RunGroupId,omitempty" name:"RunGroupId"`

	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 用户定义ID，单例运行为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDefinedId *string `json:"UserDefinedId,omitempty" name:"UserDefinedId"`

	// 表格ID，单例运行为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitempty" name:"TableId"`

	// 表格行UUID，单例运行为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableRowUuid *string `json:"TableRowUuid,omitempty" name:"TableRowUuid"`

	// 任务状态。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务输入。
	Input *string `json:"Input,omitempty" name:"Input"`

	// 运行选项。
	Option *RunOption `json:"Option,omitempty" name:"Option"`

	// 执行时间。
	ExecutionTime *ExecutionTime `json:"ExecutionTime,omitempty" name:"ExecutionTime"`

	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type RunApplicationRequestParams struct {
	// 应用ID。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 项目ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务批次名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 投递环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 任务输入JSON。需要进行base64编码。
	InputBase64 *string `json:"InputBase64,omitempty" name:"InputBase64"`

	// 任务缓存清理时间。不填表示不清理。
	CacheClearDelay *uint64 `json:"CacheClearDelay,omitempty" name:"CacheClearDelay"`

	// 运行选项。
	Option *RunOption `json:"Option,omitempty" name:"Option"`

	// 任务批次描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 批量投递表格ID，不填表示单例投递。
	TableId *string `json:"TableId,omitempty" name:"TableId"`

	// 批量投递表格行UUID。不填表示表格全部行。
	TableRowUuids []*string `json:"TableRowUuids,omitempty" name:"TableRowUuids"`
}

type RunApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 项目ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务批次名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 投递环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 任务输入JSON。需要进行base64编码。
	InputBase64 *string `json:"InputBase64,omitempty" name:"InputBase64"`

	// 任务缓存清理时间。不填表示不清理。
	CacheClearDelay *uint64 `json:"CacheClearDelay,omitempty" name:"CacheClearDelay"`

	// 运行选项。
	Option *RunOption `json:"Option,omitempty" name:"Option"`

	// 任务批次描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 批量投递表格ID，不填表示单例投递。
	TableId *string `json:"TableId,omitempty" name:"TableId"`

	// 批量投递表格行UUID。不填表示表格全部行。
	TableRowUuids []*string `json:"TableRowUuids,omitempty" name:"TableRowUuids"`
}

func (r *RunApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "ProjectId")
	delete(f, "Name")
	delete(f, "EnvironmentId")
	delete(f, "InputBase64")
	delete(f, "CacheClearDelay")
	delete(f, "Option")
	delete(f, "Description")
	delete(f, "TableId")
	delete(f, "TableRowUuids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunApplicationResponseParams struct {
	// 任务批次ID。
	RunGroupId *string `json:"RunGroupId,omitempty" name:"RunGroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RunApplicationResponse struct {
	*tchttp.BaseResponse
	Response *RunApplicationResponseParams `json:"Response"`
}

func (r *RunApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunGroup struct {
	// 任务批次ID。
	RunGroupId *string `json:"RunGroupId,omitempty" name:"RunGroupId"`

	// 项目ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 应用ID。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用类型。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 环境名称。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 表格ID，单例运行为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitempty" name:"TableId"`

	// 任务名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 任务状态。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务输入。
	Input *string `json:"Input,omitempty" name:"Input"`

	// 运行选项。
	Option *RunOption `json:"Option,omitempty" name:"Option"`

	// 任务总数量。
	TotalRun *uint64 `json:"TotalRun,omitempty" name:"TotalRun"`

	// 各状态任务的数量。
	RunStatusCounts []*RunStatusCount `json:"RunStatusCounts,omitempty" name:"RunStatusCounts"`

	// 执行时间。
	ExecutionTime *ExecutionTime `json:"ExecutionTime,omitempty" name:"ExecutionTime"`

	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type RunMetadata struct {
	// 任务类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunType *string `json:"RunType,omitempty" name:"RunType"`

	// 任务ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunId *string `json:"RunId,omitempty" name:"RunId"`

	// 父层ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *string `json:"ParentId,omitempty" name:"ParentId"`

	// 作业ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 作业名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallName *string `json:"CallName,omitempty" name:"CallName"`

	// Scatter索引。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScatterIndex *string `json:"ScatterIndex,omitempty" name:"ScatterIndex"`

	// 输入。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *string `json:"Input,omitempty" name:"Input"`

	// 输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *string `json:"Output,omitempty" name:"Output"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 提交时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`

	// 结束时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 命令行。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Command *string `json:"Command,omitempty" name:"Command"`

	// 运行时。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

	// 预处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Preprocess *bool `json:"Preprocess,omitempty" name:"Preprocess"`

	// 后处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostProcess *bool `json:"PostProcess,omitempty" name:"PostProcess"`

	// Cache命中
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallCached *bool `json:"CallCached,omitempty" name:"CallCached"`

	// 标准输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Stdout *string `json:"Stdout,omitempty" name:"Stdout"`

	// 错误输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Stderr *string `json:"Stderr,omitempty" name:"Stderr"`
}

type RunOption struct {
	// 运行失败模式，取值范围：
	// - ContinueWhilePossible
	// - NoNewCalls
	FailureMode *string `json:"FailureMode,omitempty" name:"FailureMode"`

	// 是否使用Call-Caching功能。
	UseCallCache *bool `json:"UseCallCache,omitempty" name:"UseCallCache"`

	// 是否使用错误挂起功能。
	UseErrorOnHold *bool `json:"UseErrorOnHold,omitempty" name:"UseErrorOnHold"`
}

type RunStatusCount struct {
	// 状态。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 数量。
	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type StorageOption struct {
	// 文件存储类型，取值范围：
	// - SD：通用标准型
	// - HP：通用性能型
	// - TB：turbo标准型
	// - TP：turbo性能型
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// 文件存储可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 文件系统容量，turbo系列必填，单位为GiB。 
	// - turbo标准型起售40TiB，即40960GiB；扩容步长20TiB，即20480 GiB。
	// - turbo性能型起售20TiB，即20480 GiB；扩容步长10TiB，即10240 GiB。
	Capacity *uint64 `json:"Capacity,omitempty" name:"Capacity"`
}

type VPCOption struct {
	// 子网可用区。
	SubnetZone *string `json:"SubnetZone,omitempty" name:"SubnetZone"`

	// 私有网络CIDR。
	VPCCIDRBlock *string `json:"VPCCIDRBlock,omitempty" name:"VPCCIDRBlock"`

	// 子网CIDR。
	SubnetCIDRBlock *string `json:"SubnetCIDRBlock,omitempty" name:"SubnetCIDRBlock"`
}