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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ApplicationVersion struct {
	// 版本类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 版本ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil" name:"ApplicationVersionId"`

	// 发布名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 发布描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 入口文件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Entrypoint *string `json:"Entrypoint,omitnil" name:"Entrypoint"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 创建者名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorName *string `json:"CreatorName,omitnil" name:"CreatorName"`

	// 创建者ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorId *string `json:"CreatorId,omitnil" name:"CreatorId"`

	// Git信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	GitInfo *string `json:"GitInfo,omitnil" name:"GitInfo"`
}

type CVMOption struct {
	// 云服务器可用区。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 云服务器实例规格。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`
}

type CacheInfo struct {
	// 缓存清理时间(小时)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheClearDelay *uint64 `json:"CacheClearDelay,omitnil" name:"CacheClearDelay"`

	// 缓存清理计划时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheClearTime *string `json:"CacheClearTime,omitnil" name:"CacheClearTime"`

	// 缓存是否已被清理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheCleared *bool `json:"CacheCleared,omitnil" name:"CacheCleared"`
}

type ClusterOption struct {
	// 计算集群可用区。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 计算集群类型，取值范围：
	// - KUBERNETES
	Type *string `json:"Type,omitnil" name:"Type"`

	// 资源配额。
	ResourceQuota *ResourceQuota `json:"ResourceQuota,omitnil" name:"ResourceQuota"`

	// 限制范围。
	LimitRange *LimitRange `json:"LimitRange,omitnil" name:"LimitRange"`
}

// Predefined struct for user
type CreateEnvironmentRequestParams struct {
	// 环境名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 环境配置信息。
	Config *EnvironmentConfig `json:"Config,omitnil" name:"Config"`

	// 环境描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 是否为默认环境。
	IsDefault *bool `json:"IsDefault,omitnil" name:"IsDefault"`
}

type CreateEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// 环境名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 环境配置信息。
	Config *EnvironmentConfig `json:"Config,omitnil" name:"Config"`

	// 环境描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 是否为默认环境。
	IsDefault *bool `json:"IsDefault,omitnil" name:"IsDefault"`
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
	delete(f, "IsDefault")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnvironmentResponseParams struct {
	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// 工作流UUID。
	WorkflowUuid *string `json:"WorkflowUuid,omitnil" name:"WorkflowUuid"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Zone *string `json:"Zone,omitnil" name:"Zone"`
}

// Predefined struct for user
type DeleteEnvironmentRequestParams struct {
	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`
}

type DeleteEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`
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
	WorkflowUuid *string `json:"WorkflowUuid,omitnil" name:"WorkflowUuid"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤器，支持过滤字段：
	// - EnvironmentId：环境ID
	// - Name：名称
	// - Status：环境状态
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤器，支持过滤字段：
	// - EnvironmentId：环境ID
	// - Name：名称
	// - Status：环境状态
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 环境详情列表。
	Environments []*Environment `json:"Environments,omitnil" name:"Environments"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - Name：任务批次名称
	// - RunGroupId：任务批次ID
	// - Status：任务批次状态
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeRunGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - Name：任务批次名称
	// - RunGroupId：任务批次ID
	// - Status：任务批次状态
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 任务批次列表。
	RunGroups []*RunGroup `json:"RunGroups,omitnil" name:"RunGroups"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - RunGroupId：任务批次ID
	// - Status：任务状态
	// - RunUuid：任务UUID
	// - UserDefinedId：用户定义ID
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeRunsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - RunGroupId：任务批次ID
	// - Status：任务状态
	// - RunUuid：任务UUID
	// - UserDefinedId：用户定义ID
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 任务列表。
	Runs []*Run `json:"Runs,omitnil" name:"Runs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

// Predefined struct for user
type DescribeTablesRequestParams struct {
	// 项目ID。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - Name：表格名称
	// - TableId：表格ID
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - Name：表格名称
	// - TableId：表格ID
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesResponseParams struct {
	// 结果总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 表格列表。
	Tables []*Table `json:"Tables,omitnil" name:"Tables"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTablesResponseParams `json:"Response"`
}

func (r *DescribeTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesRowsRequestParams struct {
	// 项目ID。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 表格ID。
	TableId *string `json:"TableId,omitnil" name:"TableId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - Tr：表格数据，支持模糊查询
	// - TableRowUuid：表格行UUID
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeTablesRowsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 表格ID。
	TableId *string `json:"TableId,omitnil" name:"TableId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - Tr：表格数据，支持模糊查询
	// - TableRowUuid：表格行UUID
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeTablesRowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesRowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TableId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablesRowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesRowsResponseParams struct {
	// 结果总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 表格行列表。
	Rows []*TableRow `json:"Rows,omitnil" name:"Rows"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTablesRowsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTablesRowsResponseParams `json:"Response"`
}

func (r *DescribeTablesRowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesRowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Environment struct {
	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// 环境名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 环境描述信息。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 环境地域。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 环境类型，取值范围：
	// - KUBERNETES：Kubernetes容器集群
	// - HPC：HPC高性能计算集群
	Type *string `json:"Type,omitnil" name:"Type"`

	// 环境状态，取值范围：
	// - INITIALIZING：创建中
	// - INITIALIZATION_ERROR：创建失败
	// - RUNNING：运行中
	// - ERROR：异常
	// - DELETING：正在删除
	// - DELETE_ERROR：删除失败
	Status *string `json:"Status,omitnil" name:"Status"`

	// 环境是否可用。环境需要可用才能投递计算任务。
	Available *bool `json:"Available,omitnil" name:"Available"`

	// 环境信息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 云资源ID。
	ResourceIds *ResourceIds `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// 上个工作流UUID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastWorkflowUuid *string `json:"LastWorkflowUuid,omitnil" name:"LastWorkflowUuid"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *string `json:"CreationTime,omitnil" name:"CreationTime"`
}

type EnvironmentConfig struct {
	// 私有网络配置。
	VPCOption *VPCOption `json:"VPCOption,omitnil" name:"VPCOption"`

	// 计算集群配置。
	ClusterOption *ClusterOption `json:"ClusterOption,omitnil" name:"ClusterOption"`

	// 数据库配置。
	DatabaseOption *DatabaseOption `json:"DatabaseOption,omitnil" name:"DatabaseOption"`

	// 存储配置。
	StorageOption *StorageOption `json:"StorageOption,omitnil" name:"StorageOption"`

	// 云服务器配置。
	CVMOption *CVMOption `json:"CVMOption,omitnil" name:"CVMOption"`

	// 安全组配置。
	SecurityGroupOption *SecurityGroupOption `json:"SecurityGroupOption,omitnil" name:"SecurityGroupOption"`
}

type ExecutionTime struct {
	// 提交时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitTime *string `json:"SubmitTime,omitnil" name:"SubmitTime"`

	// 开始时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type Filter struct {
	// 过滤字段。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 过滤字段值。
	Values []*string `json:"Values,omitnil" name:"Values"`
}

// Predefined struct for user
type GetRunCallsRequestParams struct {
	// 任务Uuid。
	RunUuid *string `json:"RunUuid,omitnil" name:"RunUuid"`

	// 作业路径
	Path *string `json:"Path,omitnil" name:"Path"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`
}

type GetRunCallsRequest struct {
	*tchttp.BaseRequest
	
	// 任务Uuid。
	RunUuid *string `json:"RunUuid,omitnil" name:"RunUuid"`

	// 作业路径
	Path *string `json:"Path,omitnil" name:"Path"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`
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
	delete(f, "Path")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRunCallsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRunCallsResponseParams struct {
	// 作业详情。
	Calls []*RunMetadata `json:"Calls,omitnil" name:"Calls"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type GetRunMetadataFileRequestParams struct {
	// 任务Uuid。
	RunUuid *string `json:"RunUuid,omitnil" name:"RunUuid"`

	// 需要获取的文件名。
	// 
	// 默认支持以下文件：
	// - nextflow.log
	// 
	// 提交时NFOption中report指定为true时，额外支持以下文件：
	// - execution_report.html
	// - execution_timeline.html
	// - execution_trace.txt
	// - pipeline_dag.html
	Key *string `json:"Key,omitnil" name:"Key"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`
}

type GetRunMetadataFileRequest struct {
	*tchttp.BaseRequest
	
	// 任务Uuid。
	RunUuid *string `json:"RunUuid,omitnil" name:"RunUuid"`

	// 需要获取的文件名。
	// 
	// 默认支持以下文件：
	// - nextflow.log
	// 
	// 提交时NFOption中report指定为true时，额外支持以下文件：
	// - execution_report.html
	// - execution_timeline.html
	// - execution_trace.txt
	// - pipeline_dag.html
	Key *string `json:"Key,omitnil" name:"Key"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`
}

func (r *GetRunMetadataFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRunMetadataFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RunUuid")
	delete(f, "Key")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRunMetadataFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRunMetadataFileResponseParams struct {
	// 文件预签名链接，一分钟内有效。
	CosSignedUrl *string `json:"CosSignedUrl,omitnil" name:"CosSignedUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetRunMetadataFileResponse struct {
	*tchttp.BaseResponse
	Response *GetRunMetadataFileResponseParams `json:"Response"`
}

func (r *GetRunMetadataFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRunMetadataFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRunStatusRequestParams struct {
	// 任务Uuid。
	RunUuid *string `json:"RunUuid,omitnil" name:"RunUuid"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`
}

type GetRunStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务Uuid。
	RunUuid *string `json:"RunUuid,omitnil" name:"RunUuid"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`
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
	Metadata *RunMetadata `json:"Metadata,omitnil" name:"Metadata"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type GitInfo struct {
	// Git地址。
	GitHttpPath *string `json:"GitHttpPath,omitnil" name:"GitHttpPath"`

	// Git用户名。
	GitUserName *string `json:"GitUserName,omitnil" name:"GitUserName"`

	// Git密码或者Token。
	GitTokenOrPassword *string `json:"GitTokenOrPassword,omitnil" name:"GitTokenOrPassword"`

	// 分支。
	Branch *string `json:"Branch,omitnil" name:"Branch"`

	// 标签。
	Tag *string `json:"Tag,omitnil" name:"Tag"`
}

// Predefined struct for user
type ImportTableFileRequestParams struct {
	// 表格关联的项目ID。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 表格名称。最多支持200个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 表格文件Cos对象路径。
	CosUri *string `json:"CosUri,omitnil" name:"CosUri"`

	// 表格文件中每列的数据类型，支持的类型包括：Int、Float、String、File、Boolean、Array[Int]、Array[Float]、Array[String]、Array[File]、Array[Boolean]
	DataType []*string `json:"DataType,omitnil" name:"DataType"`

	// 表格描述。最多支持500个字符。
	Description *string `json:"Description,omitnil" name:"Description"`
}

type ImportTableFileRequest struct {
	*tchttp.BaseRequest
	
	// 表格关联的项目ID。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 表格名称。最多支持200个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 表格文件Cos对象路径。
	CosUri *string `json:"CosUri,omitnil" name:"CosUri"`

	// 表格文件中每列的数据类型，支持的类型包括：Int、Float、String、File、Boolean、Array[Int]、Array[Float]、Array[String]、Array[File]、Array[Boolean]
	DataType []*string `json:"DataType,omitnil" name:"DataType"`

	// 表格描述。最多支持500个字符。
	Description *string `json:"Description,omitnil" name:"Description"`
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
	TableId *string `json:"TableId,omitnil" name:"TableId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type LimitRange struct {
	// 最大CPU设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxCPU *string `json:"MaxCPU,omitnil" name:"MaxCPU"`

	// 最大内存设置（单位：Mi，Gi，Ti，M，G，T）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMemory *string `json:"MaxMemory,omitnil" name:"MaxMemory"`
}

type NFOption struct {
	// Config。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config *string `json:"Config,omitnil" name:"Config"`

	// Profile。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Profile *string `json:"Profile,omitnil" name:"Profile"`

	// Report。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Report *bool `json:"Report,omitnil" name:"Report"`

	// Resume。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resume *bool `json:"Resume,omitnil" name:"Resume"`
}

type ResourceIds struct {
	// 私有网络ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VPCId *string `json:"VPCId,omitnil" name:"VPCId"`

	// 子网ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 安全组ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// TDSQL-C Mysql版数据库ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TDSQLCId *string `json:"TDSQLCId,omitnil" name:"TDSQLCId"`

	// 文件存储ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFSId *string `json:"CFSId,omitnil" name:"CFSId"`

	// 文件存储类型：取值范围：
	// - SD：通用标准型
	// - HP：通用性能型
	// - TB：turbo标准型
	// - TP：turbo性能型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFSStorageType *string `json:"CFSStorageType,omitnil" name:"CFSStorageType"`

	// 云服务器ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVMId *string `json:"CVMId,omitnil" name:"CVMId"`

	// 弹性容器集群ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EKSId *string `json:"EKSId,omitnil" name:"EKSId"`
}

type ResourceQuota struct {
	// CPU Limit设置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CPULimit *string `json:"CPULimit,omitnil" name:"CPULimit"`

	// 内存Limit设置（单位：Mi，Gi，Ti，M，G，T）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryLimit *string `json:"MemoryLimit,omitnil" name:"MemoryLimit"`

	// Pods数量设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pods *string `json:"Pods,omitnil" name:"Pods"`
}

// Predefined struct for user
type RetryRunsRequestParams struct {
	// 关联项目ID。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 任务UUID。
	RunUuids []*string `json:"RunUuids,omitnil" name:"RunUuids"`
}

type RetryRunsRequest struct {
	*tchttp.BaseRequest
	
	// 关联项目ID。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 任务UUID。
	RunUuids []*string `json:"RunUuids,omitnil" name:"RunUuids"`
}

func (r *RetryRunsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryRunsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "RunUuids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryRunsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryRunsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RetryRunsResponse struct {
	*tchttp.BaseResponse
	Response *RetryRunsResponseParams `json:"Response"`
}

func (r *RetryRunsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryRunsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Run struct {
	// 任务UUID。
	RunUuid *string `json:"RunUuid,omitnil" name:"RunUuid"`

	// 项目ID。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 应用ID。
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 任务批次ID。
	RunGroupId *string `json:"RunGroupId,omitnil" name:"RunGroupId"`

	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// 用户定义ID，单例运行为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDefinedId *string `json:"UserDefinedId,omitnil" name:"UserDefinedId"`

	// 表格ID，单例运行为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil" name:"TableId"`

	// 表格行UUID，单例运行为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableRowUuid *string `json:"TableRowUuid,omitnil" name:"TableRowUuid"`

	// 任务状态。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 任务输入。
	Input *string `json:"Input,omitnil" name:"Input"`

	// 运行选项。
	//
	// Deprecated: Option is deprecated.
	Option *RunOption `json:"Option,omitnil" name:"Option"`

	// 执行时间。
	ExecutionTime *ExecutionTime `json:"ExecutionTime,omitnil" name:"ExecutionTime"`

	// 缓存信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cache *CacheInfo `json:"Cache,omitnil" name:"Cache"`

	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitnil" name:"ErrorMessage"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

// Predefined struct for user
type RunApplicationRequestParams struct {
	// 应用ID。
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 项目ID。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 任务批次名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 投递环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// 任务输入JSON。需要进行base64编码。
	InputBase64 *string `json:"InputBase64,omitnil" name:"InputBase64"`

	// 任务缓存清理时间。不填表示不清理。
	CacheClearDelay *uint64 `json:"CacheClearDelay,omitnil" name:"CacheClearDelay"`

	// 运行选项。
	Option *RunOption `json:"Option,omitnil" name:"Option"`

	// 任务批次描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 批量投递表格ID，不填表示单例投递。
	TableId *string `json:"TableId,omitnil" name:"TableId"`

	// 批量投递表格行UUID。不填表示表格全部行。
	TableRowUuids []*string `json:"TableRowUuids,omitnil" name:"TableRowUuids"`

	// 应用版本ID。不填表示使用当前最新版本。
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil" name:"ApplicationVersionId"`
}

type RunApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID。
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 项目ID。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 任务批次名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 投递环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// 任务输入JSON。需要进行base64编码。
	InputBase64 *string `json:"InputBase64,omitnil" name:"InputBase64"`

	// 任务缓存清理时间。不填表示不清理。
	CacheClearDelay *uint64 `json:"CacheClearDelay,omitnil" name:"CacheClearDelay"`

	// 运行选项。
	Option *RunOption `json:"Option,omitnil" name:"Option"`

	// 任务批次描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 批量投递表格ID，不填表示单例投递。
	TableId *string `json:"TableId,omitnil" name:"TableId"`

	// 批量投递表格行UUID。不填表示表格全部行。
	TableRowUuids []*string `json:"TableRowUuids,omitnil" name:"TableRowUuids"`

	// 应用版本ID。不填表示使用当前最新版本。
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil" name:"ApplicationVersionId"`
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
	delete(f, "ApplicationVersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunApplicationResponseParams struct {
	// 任务批次ID。
	RunGroupId *string `json:"RunGroupId,omitnil" name:"RunGroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RunGroupId *string `json:"RunGroupId,omitnil" name:"RunGroupId"`

	// 项目ID。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 项目名称。
	ProjectName *string `json:"ProjectName,omitnil" name:"ProjectName"`

	// 应用ID。
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 应用名称。
	ApplicationName *string `json:"ApplicationName,omitnil" name:"ApplicationName"`

	// 应用类型。
	ApplicationType *string `json:"ApplicationType,omitnil" name:"ApplicationType"`

	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// 环境名称。
	EnvironmentName *string `json:"EnvironmentName,omitnil" name:"EnvironmentName"`

	// 表格ID，单例运行为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil" name:"TableId"`

	// 任务名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 任务描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 任务状态。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 任务输入。
	Input *string `json:"Input,omitnil" name:"Input"`

	// WDL运行选项。
	Option *RunOption `json:"Option,omitnil" name:"Option"`

	// Nextflow运行选项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NFOption *NFOption `json:"NFOption,omitnil" name:"NFOption"`

	// 任务总数量。
	TotalRun *uint64 `json:"TotalRun,omitnil" name:"TotalRun"`

	// 各状态任务的数量。
	RunStatusCounts []*RunStatusCount `json:"RunStatusCounts,omitnil" name:"RunStatusCounts"`

	// 执行时间。
	ExecutionTime *ExecutionTime `json:"ExecutionTime,omitnil" name:"ExecutionTime"`

	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitnil" name:"ErrorMessage"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 创建者。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil" name:"Creator"`

	// 创建者ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorId *string `json:"CreatorId,omitnil" name:"CreatorId"`

	// 运行结果通知方式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultNotify *string `json:"ResultNotify,omitnil" name:"ResultNotify"`

	// 应用版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationVersion *ApplicationVersion `json:"ApplicationVersion,omitnil" name:"ApplicationVersion"`
}

type RunMetadata struct {
	// 任务类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunType *string `json:"RunType,omitnil" name:"RunType"`

	// 任务ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunId *string `json:"RunId,omitnil" name:"RunId"`

	// 父层ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *string `json:"ParentId,omitnil" name:"ParentId"`

	// 作业ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 作业名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallName *string `json:"CallName,omitnil" name:"CallName"`

	// Scatter索引。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScatterIndex *string `json:"ScatterIndex,omitnil" name:"ScatterIndex"`

	// 输入。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *string `json:"Input,omitnil" name:"Input"`

	// 输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *string `json:"Output,omitnil" name:"Output"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil" name:"ErrorMessage"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 提交时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitTime *string `json:"SubmitTime,omitnil" name:"SubmitTime"`

	// 结束时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 命令行。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Command *string `json:"Command,omitnil" name:"Command"`

	// 运行时。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Runtime *string `json:"Runtime,omitnil" name:"Runtime"`

	// 预处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Preprocess *bool `json:"Preprocess,omitnil" name:"Preprocess"`

	// 后处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostProcess *bool `json:"PostProcess,omitnil" name:"PostProcess"`

	// Cache命中
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallCached *bool `json:"CallCached,omitnil" name:"CallCached"`

	// 标准输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Stdout *string `json:"Stdout,omitnil" name:"Stdout"`

	// 错误输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Stderr *string `json:"Stderr,omitnil" name:"Stderr"`

	// 其他信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Meta *string `json:"Meta,omitnil" name:"Meta"`
}

type RunOption struct {
	// 运行失败模式，取值范围：
	// - ContinueWhilePossible
	// - NoNewCalls
	FailureMode *string `json:"FailureMode,omitnil" name:"FailureMode"`

	// 是否使用Call-Caching功能。
	UseCallCache *bool `json:"UseCallCache,omitnil" name:"UseCallCache"`

	// 是否使用错误挂起功能。
	UseErrorOnHold *bool `json:"UseErrorOnHold,omitnil" name:"UseErrorOnHold"`

	// 输出归档COS路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinalWorkflowOutputsDir *string `json:"FinalWorkflowOutputsDir,omitnil" name:"FinalWorkflowOutputsDir"`

	// 是否使用相对目录归档输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseRelativeOutputPaths *bool `json:"UseRelativeOutputPaths,omitnil" name:"UseRelativeOutputPaths"`
}

type RunStatusCount struct {
	// 状态。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 数量。
	Count *uint64 `json:"Count,omitnil" name:"Count"`
}

// Predefined struct for user
type RunWorkflowRequestParams struct {
	// 任务批次名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 投递环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// 工作流Git仓库信息。
	GitSource *GitInfo `json:"GitSource,omitnil" name:"GitSource"`

	// 工作流类型。
	// 
	// 支持类型：
	// - NEXTFLOW
	Type *string `json:"Type,omitnil" name:"Type"`

	// Nextflow选项。
	NFOption *NFOption `json:"NFOption,omitnil" name:"NFOption"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 任务批次描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 任务输入JSON。需要进行base64编码。
	// （InputBase64和InputCosUri必选其一）
	InputBase64 *string `json:"InputBase64,omitnil" name:"InputBase64"`

	// 任务输入COS地址。
	// （InputBase64和InputCosUri必选其一）
	InputCosUri *string `json:"InputCosUri,omitnil" name:"InputCosUri"`

	// 任务缓存清理时间。不填表示不清理。
	CacheClearDelay *uint64 `json:"CacheClearDelay,omitnil" name:"CacheClearDelay"`
}

type RunWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 任务批次名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 投递环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// 工作流Git仓库信息。
	GitSource *GitInfo `json:"GitSource,omitnil" name:"GitSource"`

	// 工作流类型。
	// 
	// 支持类型：
	// - NEXTFLOW
	Type *string `json:"Type,omitnil" name:"Type"`

	// Nextflow选项。
	NFOption *NFOption `json:"NFOption,omitnil" name:"NFOption"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 任务批次描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 任务输入JSON。需要进行base64编码。
	// （InputBase64和InputCosUri必选其一）
	InputBase64 *string `json:"InputBase64,omitnil" name:"InputBase64"`

	// 任务输入COS地址。
	// （InputBase64和InputCosUri必选其一）
	InputCosUri *string `json:"InputCosUri,omitnil" name:"InputCosUri"`

	// 任务缓存清理时间。不填表示不清理。
	CacheClearDelay *uint64 `json:"CacheClearDelay,omitnil" name:"CacheClearDelay"`
}

func (r *RunWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "EnvironmentId")
	delete(f, "GitSource")
	delete(f, "Type")
	delete(f, "NFOption")
	delete(f, "ProjectId")
	delete(f, "Description")
	delete(f, "InputBase64")
	delete(f, "InputCosUri")
	delete(f, "CacheClearDelay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunWorkflowResponseParams struct {
	// 任务批次ID。
	RunGroupId *string `json:"RunGroupId,omitnil" name:"RunGroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RunWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *RunWorkflowResponseParams `json:"Response"`
}

func (r *RunWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupOption struct {
	// 安全组ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`
}

type StorageOption struct {
	// 文件存储类型，取值范围：
	// - SD：通用标准型
	// - HP：通用性能型
	// - TB：turbo标准型
	// - TP：turbo性能型
	StorageType *string `json:"StorageType,omitnil" name:"StorageType"`

	// 文件存储可用区。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 文件系统容量，turbo系列必填，单位为GiB。 
	// - turbo标准型起售40TiB，即40960GiB；扩容步长20TiB，即20480 GiB。
	// - turbo性能型起售20TiB，即20480 GiB；扩容步长10TiB，即10240 GiB。
	Capacity *uint64 `json:"Capacity,omitnil" name:"Capacity"`
}

type Table struct {
	// 表格ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil" name:"TableId"`

	// 关联项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 表格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 表格描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 表格列
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*TableColumn `json:"Columns,omitnil" name:"Columns"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil" name:"Creator"`
}

type TableColumn struct {
	// 列名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Header *string `json:"Header,omitnil" name:"Header"`

	// 列数据类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataType *string `json:"DataType,omitnil" name:"DataType"`
}

type TableRow struct {
	// 表格行UUID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableRowUuid *string `json:"TableRowUuid,omitnil" name:"TableRowUuid"`

	// 表格行内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*string `json:"Content,omitnil" name:"Content"`
}

// Predefined struct for user
type TerminateRunGroupRequestParams struct {
	// 任务批次ID。
	RunGroupId *string `json:"RunGroupId,omitnil" name:"RunGroupId"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`
}

type TerminateRunGroupRequest struct {
	*tchttp.BaseRequest
	
	// 任务批次ID。
	RunGroupId *string `json:"RunGroupId,omitnil" name:"RunGroupId"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`
}

func (r *TerminateRunGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateRunGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RunGroupId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateRunGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateRunGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TerminateRunGroupResponse struct {
	*tchttp.BaseResponse
	Response *TerminateRunGroupResponseParams `json:"Response"`
}

func (r *TerminateRunGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateRunGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VPCOption struct {
	// 私有网络ID（VPCId和VPCCIDRBlock必选其一。若使用VPCId，则使用现用私有网络；若使用VPCCIDRBlock，则创建新的私有网络）
	VPCId *string `json:"VPCId,omitnil" name:"VPCId"`

	// 子网ID（SubnetId和SubnetZone&SubnetCIDRBlock必选其一。若使用SubnetId，则使用现用子网；若使用SubnetZone&SubnetCIDRBlock，则创建新的子网）
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 子网可用区。
	SubnetZone *string `json:"SubnetZone,omitnil" name:"SubnetZone"`

	// 私有网络CIDR。
	VPCCIDRBlock *string `json:"VPCCIDRBlock,omitnil" name:"VPCCIDRBlock"`

	// 子网CIDR。
	SubnetCIDRBlock *string `json:"SubnetCIDRBlock,omitnil" name:"SubnetCIDRBlock"`
}