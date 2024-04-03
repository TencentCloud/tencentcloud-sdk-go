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
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 版本ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`

	// 发布名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 发布描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 入口文件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Entrypoint *string `json:"Entrypoint,omitnil,omitempty" name:"Entrypoint"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 创建者名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorName *string `json:"CreatorName,omitnil,omitempty" name:"CreatorName"`

	// 创建者ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorId *string `json:"CreatorId,omitnil,omitempty" name:"CreatorId"`

	// Git信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	GitInfo *string `json:"GitInfo,omitnil,omitempty" name:"GitInfo"`
}

type CVMOption struct {
	// 云服务器可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 云服务器实例规格。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type CacheInfo struct {
	// 缓存清理时间(小时)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheClearDelay *uint64 `json:"CacheClearDelay,omitnil,omitempty" name:"CacheClearDelay"`

	// 缓存清理计划时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheClearTime *string `json:"CacheClearTime,omitnil,omitempty" name:"CacheClearTime"`

	// 缓存是否已被清理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheCleared *bool `json:"CacheCleared,omitnil,omitempty" name:"CacheCleared"`
}

type ClusterOption struct {
	// 计算集群可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 计算集群类型，取值范围：
	// - KUBERNETES
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 计算集群Service CIDR，不能与VPC网段重合。
	ServiceCidr *string `json:"ServiceCidr,omitnil,omitempty" name:"ServiceCidr"`

	// 资源配额。
	ResourceQuota *ResourceQuota `json:"ResourceQuota,omitnil,omitempty" name:"ResourceQuota"`

	// 限制范围。
	LimitRange *LimitRange `json:"LimitRange,omitnil,omitempty" name:"LimitRange"`
}

// Predefined struct for user
type CreateEnvironmentRequestParams struct {
	// 环境名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 环境配置信息。
	Config *EnvironmentConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 环境描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否为默认环境。
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`
}

type CreateEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// 环境名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 环境配置信息。
	Config *EnvironmentConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 环境描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否为默认环境。
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`
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
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 工作流UUID。
	WorkflowUuid *string `json:"WorkflowUuid,omitnil,omitempty" name:"WorkflowUuid"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CreateVolumeRequestParams struct {
	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 缓存卷类型，取值范围：
	// * SHARED：多点挂载共享存储
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 缓存卷规格，取值范围：
	// 
	// - SD：通用标准型
	// - HP：通用性能型
	// - TB：turbo标准型
	// - TP：turbo性能型
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// 描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 缓存卷大小（GB），Turbo系列需要指定。
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`
}

type CreateVolumeRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 缓存卷类型，取值范围：
	// * SHARED：多点挂载共享存储
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 缓存卷规格，取值范围：
	// 
	// - SD：通用标准型
	// - HP：通用性能型
	// - TB：turbo标准型
	// - TP：turbo性能型
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// 描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 缓存卷大小（GB），Turbo系列需要指定。
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`
}

func (r *CreateVolumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVolumeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "Spec")
	delete(f, "Description")
	delete(f, "Capacity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVolumeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVolumeResponseParams struct {
	// 缓存卷ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeId *string `json:"VolumeId,omitnil,omitempty" name:"VolumeId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVolumeResponse struct {
	*tchttp.BaseResponse
	Response *CreateVolumeResponseParams `json:"Response"`
}

func (r *CreateVolumeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVolumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabaseOption struct {
	// 数据库可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

// Predefined struct for user
type DeleteEnvironmentRequestParams struct {
	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`
}

type DeleteEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`
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
	WorkflowUuid *string `json:"WorkflowUuid,omitnil,omitempty" name:"WorkflowUuid"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteVolumeDataRequestParams struct {
	// 缓存卷ID。
	VolumeId *string `json:"VolumeId,omitnil,omitempty" name:"VolumeId"`

	// 需要删除的路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type DeleteVolumeDataRequest struct {
	*tchttp.BaseRequest
	
	// 缓存卷ID。
	VolumeId *string `json:"VolumeId,omitnil,omitempty" name:"VolumeId"`

	// 需要删除的路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

func (r *DeleteVolumeDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVolumeDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VolumeId")
	delete(f, "Path")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVolumeDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVolumeDataResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVolumeDataResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVolumeDataResponseParams `json:"Response"`
}

func (r *DeleteVolumeDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVolumeDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVolumeRequestParams struct {
	// 缓存卷ID。
	VolumeId *string `json:"VolumeId,omitnil,omitempty" name:"VolumeId"`
}

type DeleteVolumeRequest struct {
	*tchttp.BaseRequest
	
	// 缓存卷ID。
	VolumeId *string `json:"VolumeId,omitnil,omitempty" name:"VolumeId"`
}

func (r *DeleteVolumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVolumeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VolumeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVolumeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVolumeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVolumeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVolumeResponseParams `json:"Response"`
}

func (r *DeleteVolumeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVolumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsRequestParams struct {
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤器，支持过滤字段：
	// - EnvironmentId：环境ID
	// - Name：名称
	// - Status：环境状态
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤器，支持过滤字段：
	// - EnvironmentId：环境ID
	// - Name：名称
	// - Status：环境状态
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 环境详情列表。
	Environments []*Environment `json:"Environments,omitnil,omitempty" name:"Environments"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - Name：任务批次名称
	// - RunGroupId：任务批次ID
	// - Status：任务批次状态
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRunGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - Name：任务批次名称
	// - RunGroupId：任务批次ID
	// - Status：任务批次状态
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务批次列表。
	RunGroups []*RunGroup `json:"RunGroups,omitnil,omitempty" name:"RunGroups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - RunGroupId：任务批次ID
	// - Status：任务状态
	// - RunUuid：任务UUID
	// - UserDefinedId：用户定义ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRunsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - RunGroupId：任务批次ID
	// - Status：任务状态
	// - RunUuid：任务UUID
	// - UserDefinedId：用户定义ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务列表。
	Runs []*Run `json:"Runs,omitnil,omitempty" name:"Runs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - Name：表格名称
	// - TableId：表格ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - Name：表格名称
	// - TableId：表格ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 表格列表。
	Tables []*Table `json:"Tables,omitnil,omitempty" name:"Tables"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 表格ID。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - Tr：表格数据，支持模糊查询
	// - TableRowUuid：表格行UUID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTablesRowsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 表格ID。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - Tr：表格数据，支持模糊查询
	// - TableRowUuid：表格行UUID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 表格行列表。
	Rows []*TableRow `json:"Rows,omitnil,omitempty" name:"Rows"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DescribeVolumesRequestParams struct {
	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - Name：名称
	// - IsDefault：是否为默认
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeVolumesRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤器，支持过滤字段：
	// - Name：名称
	// - IsDefault：是否为默认
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeVolumesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVolumesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVolumesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVolumesResponseParams struct {
	// 缓存卷。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Volumes []*Volume `json:"Volumes,omitnil,omitempty" name:"Volumes"`

	// 符合条件的数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVolumesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVolumesResponseParams `json:"Response"`
}

func (r *DescribeVolumesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVolumesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Environment struct {
	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 环境名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 环境描述信息。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 环境地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 环境类型，取值范围：
	// - KUBERNETES：Kubernetes容器集群
	// - HPC：HPC高性能计算集群
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 环境状态，取值范围：
	// - INITIALIZING：创建中
	// - INITIALIZATION_ERROR：创建失败
	// - RUNNING：运行中
	// - ERROR：异常
	// - DELETING：正在删除
	// - DELETE_ERROR：删除失败
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 环境是否可用。环境需要可用才能投递计算任务。
	Available *bool `json:"Available,omitnil,omitempty" name:"Available"`

	// 环境是否为默认环境。
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 环境是否为托管环境。
	IsManaged *bool `json:"IsManaged,omitnil,omitempty" name:"IsManaged"`

	// 环境信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 云资源ID。
	ResourceIds *ResourceIds `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 上个工作流UUID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastWorkflowUuid *string `json:"LastWorkflowUuid,omitnil,omitempty" name:"LastWorkflowUuid"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`
}

type EnvironmentConfig struct {
	// 私有网络配置。
	VPCOption *VPCOption `json:"VPCOption,omitnil,omitempty" name:"VPCOption"`

	// 计算集群配置。
	ClusterOption *ClusterOption `json:"ClusterOption,omitnil,omitempty" name:"ClusterOption"`

	// 数据库配置。
	DatabaseOption *DatabaseOption `json:"DatabaseOption,omitnil,omitempty" name:"DatabaseOption"`

	// 存储配置。
	StorageOption *StorageOption `json:"StorageOption,omitnil,omitempty" name:"StorageOption"`

	// 云服务器配置。
	CVMOption *CVMOption `json:"CVMOption,omitnil,omitempty" name:"CVMOption"`

	// 安全组配置。
	SecurityGroupOption *SecurityGroupOption `json:"SecurityGroupOption,omitnil,omitempty" name:"SecurityGroupOption"`
}

type ExecutionTime struct {
	// 提交时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitTime *string `json:"SubmitTime,omitnil,omitempty" name:"SubmitTime"`

	// 开始时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type Filter struct {
	// 过滤字段。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤字段值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type GetRunCallsRequestParams struct {
	// 任务Uuid。
	RunUuid *string `json:"RunUuid,omitnil,omitempty" name:"RunUuid"`

	// 作业路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type GetRunCallsRequest struct {
	*tchttp.BaseRequest
	
	// 任务Uuid。
	RunUuid *string `json:"RunUuid,omitnil,omitempty" name:"RunUuid"`

	// 作业路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	Calls []*RunMetadata `json:"Calls,omitnil,omitempty" name:"Calls"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RunUuid *string `json:"RunUuid,omitnil,omitempty" name:"RunUuid"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

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
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 需要批量获取的文件名。
	// 
	// 默认支持以下文件：
	// - nextflow.log
	// 
	// 提交时NFOption中report指定为true时，额外支持以下文件：
	// - execution_report.html
	// - execution_timeline.html
	// - execution_trace.txt
	// - pipeline_dag.html
	Keys []*string `json:"Keys,omitnil,omitempty" name:"Keys"`
}

type GetRunMetadataFileRequest struct {
	*tchttp.BaseRequest
	
	// 任务Uuid。
	RunUuid *string `json:"RunUuid,omitnil,omitempty" name:"RunUuid"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

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
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 需要批量获取的文件名。
	// 
	// 默认支持以下文件：
	// - nextflow.log
	// 
	// 提交时NFOption中report指定为true时，额外支持以下文件：
	// - execution_report.html
	// - execution_timeline.html
	// - execution_trace.txt
	// - pipeline_dag.html
	Keys []*string `json:"Keys,omitnil,omitempty" name:"Keys"`
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
	delete(f, "ProjectId")
	delete(f, "Key")
	delete(f, "Keys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRunMetadataFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRunMetadataFileResponseParams struct {
	// 文件预签名链接，一分钟内有效。
	CosSignedUrl *string `json:"CosSignedUrl,omitnil,omitempty" name:"CosSignedUrl"`

	// 批量文件预签名链接，一分钟内有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosSignedUrls []*string `json:"CosSignedUrls,omitnil,omitempty" name:"CosSignedUrls"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RunUuid *string `json:"RunUuid,omitnil,omitempty" name:"RunUuid"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type GetRunStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务Uuid。
	RunUuid *string `json:"RunUuid,omitnil,omitempty" name:"RunUuid"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	Metadata *RunMetadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GitHttpPath *string `json:"GitHttpPath,omitnil,omitempty" name:"GitHttpPath"`

	// Git用户名。
	GitUserName *string `json:"GitUserName,omitnil,omitempty" name:"GitUserName"`

	// Git密码或者Token。
	GitTokenOrPassword *string `json:"GitTokenOrPassword,omitnil,omitempty" name:"GitTokenOrPassword"`

	// 分支。
	Branch *string `json:"Branch,omitnil,omitempty" name:"Branch"`

	// 标签。
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`
}

// Predefined struct for user
type ImportTableFileRequestParams struct {
	// 表格关联的项目ID。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 表格名称。最多支持200个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 表格文件Cos对象路径。
	CosUri *string `json:"CosUri,omitnil,omitempty" name:"CosUri"`

	// 表格文件中每列的数据类型，支持的类型包括：Int、Float、String、File、Boolean、Array[Int]、Array[Float]、Array[String]、Array[File]、Array[Boolean]
	DataType []*string `json:"DataType,omitnil,omitempty" name:"DataType"`

	// 表格描述。最多支持500个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ImportTableFileRequest struct {
	*tchttp.BaseRequest
	
	// 表格关联的项目ID。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 表格名称。最多支持200个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 表格文件Cos对象路径。
	CosUri *string `json:"CosUri,omitnil,omitempty" name:"CosUri"`

	// 表格文件中每列的数据类型，支持的类型包括：Int、Float、String、File、Boolean、Array[Int]、Array[Float]、Array[String]、Array[File]、Array[Boolean]
	DataType []*string `json:"DataType,omitnil,omitempty" name:"DataType"`

	// 表格描述。最多支持500个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
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
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	MaxCPU *string `json:"MaxCPU,omitnil,omitempty" name:"MaxCPU"`

	// 最大内存设置（单位：Mi，Gi，Ti，M，G，T）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMemory *string `json:"MaxMemory,omitnil,omitempty" name:"MaxMemory"`
}

// Predefined struct for user
type ModifyVolumeRequestParams struct {
	// 缓存卷ID。
	VolumeId *string `json:"VolumeId,omitnil,omitempty" name:"VolumeId"`

	// 名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyVolumeRequest struct {
	*tchttp.BaseRequest
	
	// 缓存卷ID。
	VolumeId *string `json:"VolumeId,omitnil,omitempty" name:"VolumeId"`

	// 名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyVolumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVolumeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VolumeId")
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVolumeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVolumeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVolumeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVolumeResponseParams `json:"Response"`
}

func (r *ModifyVolumeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVolumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NFOption struct {
	// Config。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// Profile。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Profile *string `json:"Profile,omitnil,omitempty" name:"Profile"`

	// Report。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Report *bool `json:"Report,omitnil,omitempty" name:"Report"`

	// Resume。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resume *bool `json:"Resume,omitnil,omitempty" name:"Resume"`

	// Nextflow引擎版本，取值范围：
	// - 22.10.4
	// - 22.10.8 
	// - 23.10.1
	// 注意：此字段可能返回 null，表示取不到有效值。
	NFVersion *string `json:"NFVersion,omitnil,omitempty" name:"NFVersion"`
}

type ResourceIds struct {
	// 私有网络ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VPCId *string `json:"VPCId,omitnil,omitempty" name:"VPCId"`

	// 子网ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 安全组ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// TDSQL-C Mysql版数据库ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TDSQLCId *string `json:"TDSQLCId,omitnil,omitempty" name:"TDSQLCId"`

	// 文件存储ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFSId *string `json:"CFSId,omitnil,omitempty" name:"CFSId"`

	// 文件存储类型：取值范围：
	// - SD：通用标准型
	// - HP：通用性能型
	// - TB：turbo标准型
	// - TP：turbo性能型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFSStorageType *string `json:"CFSStorageType,omitnil,omitempty" name:"CFSStorageType"`

	// 云服务器ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVMId *string `json:"CVMId,omitnil,omitempty" name:"CVMId"`

	// 弹性容器集群ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EKSId *string `json:"EKSId,omitnil,omitempty" name:"EKSId"`
}

type ResourceQuota struct {
	// CPU Limit设置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CPULimit *string `json:"CPULimit,omitnil,omitempty" name:"CPULimit"`

	// 内存Limit设置（单位：Mi，Gi，Ti，M，G，T）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryLimit *string `json:"MemoryLimit,omitnil,omitempty" name:"MemoryLimit"`

	// Pods数量设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pods *string `json:"Pods,omitnil,omitempty" name:"Pods"`
}

// Predefined struct for user
type RetryRunsRequestParams struct {
	// 项目ID。（不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 需要重试的任务批次ID。
	RunGroupId *string `json:"RunGroupId,omitnil,omitempty" name:"RunGroupId"`

	// 需要重试的任务UUID。
	RunUuids []*string `json:"RunUuids,omitnil,omitempty" name:"RunUuids"`

	// WDL运行选项，不填使用被重试的任务批次运行选项。
	WDLOption *RunOption `json:"WDLOption,omitnil,omitempty" name:"WDLOption"`

	// Nextflow运行选项，不填使用被重试的任务批次运行选项。
	NFOption *NFOption `json:"NFOption,omitnil,omitempty" name:"NFOption"`
}

type RetryRunsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID。（不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 需要重试的任务批次ID。
	RunGroupId *string `json:"RunGroupId,omitnil,omitempty" name:"RunGroupId"`

	// 需要重试的任务UUID。
	RunUuids []*string `json:"RunUuids,omitnil,omitempty" name:"RunUuids"`

	// WDL运行选项，不填使用被重试的任务批次运行选项。
	WDLOption *RunOption `json:"WDLOption,omitnil,omitempty" name:"WDLOption"`

	// Nextflow运行选项，不填使用被重试的任务批次运行选项。
	NFOption *NFOption `json:"NFOption,omitnil,omitempty" name:"NFOption"`
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
	delete(f, "RunGroupId")
	delete(f, "RunUuids")
	delete(f, "WDLOption")
	delete(f, "NFOption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryRunsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryRunsResponseParams struct {
	// 新的任务批次ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunGroupId *string `json:"RunGroupId,omitnil,omitempty" name:"RunGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RunUuid *string `json:"RunUuid,omitnil,omitempty" name:"RunUuid"`

	// 项目ID。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 应用ID。
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 任务批次ID。
	RunGroupId *string `json:"RunGroupId,omitnil,omitempty" name:"RunGroupId"`

	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 用户定义ID，单例运行为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDefinedId *string `json:"UserDefinedId,omitnil,omitempty" name:"UserDefinedId"`

	// 表格ID，单例运行为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 表格行UUID，单例运行为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableRowUuid *string `json:"TableRowUuid,omitnil,omitempty" name:"TableRowUuid"`

	// 任务状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务输入。
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// 运行选项。
	//
	// Deprecated: Option is deprecated.
	Option *RunOption `json:"Option,omitnil,omitempty" name:"Option"`

	// 执行时间。
	ExecutionTime *ExecutionTime `json:"ExecutionTime,omitnil,omitempty" name:"ExecutionTime"`

	// 缓存信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cache *CacheInfo `json:"Cache,omitnil,omitempty" name:"Cache"`

	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type RunApplicationRequestParams struct {
	// 应用ID。
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 任务批次名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 投递环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 项目ID。（不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务批次描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 任务输入COS地址。（InputBase64和InputCosUri必选其一）
	InputCosUri *string `json:"InputCosUri,omitnil,omitempty" name:"InputCosUri"`

	// 任务输入JSON。需要进行base64编码。（InputBase64和InputCosUri必选其一）
	InputBase64 *string `json:"InputBase64,omitnil,omitempty" name:"InputBase64"`

	// 批量投递表格ID，不填表示单例投递。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 批量投递表格行UUID。不填表示表格全部行。
	TableRowUuids []*string `json:"TableRowUuids,omitnil,omitempty" name:"TableRowUuids"`

	// 任务缓存清理时间（小时）。不填或0表示不清理。
	CacheClearDelay *uint64 `json:"CacheClearDelay,omitnil,omitempty" name:"CacheClearDelay"`

	// 应用版本ID。不填表示使用当前最新版本。
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`

	// WDL运行选项。
	Option *RunOption `json:"Option,omitnil,omitempty" name:"Option"`

	// Nextflow运行选项。
	NFOption *NFOption `json:"NFOption,omitnil,omitempty" name:"NFOption"`

	// 工作目录，使用缓存卷内的相对路径 (暂时仅支持Nextflow)
	WorkDir *string `json:"WorkDir,omitnil,omitempty" name:"WorkDir"`

	// 访问模式，不填默认私有。取值范围
	// - PRIVATE：私有应用
	// - PUBLIC：公共应用
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`
}

type RunApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID。
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 任务批次名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 投递环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 项目ID。（不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务批次描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 任务输入COS地址。（InputBase64和InputCosUri必选其一）
	InputCosUri *string `json:"InputCosUri,omitnil,omitempty" name:"InputCosUri"`

	// 任务输入JSON。需要进行base64编码。（InputBase64和InputCosUri必选其一）
	InputBase64 *string `json:"InputBase64,omitnil,omitempty" name:"InputBase64"`

	// 批量投递表格ID，不填表示单例投递。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 批量投递表格行UUID。不填表示表格全部行。
	TableRowUuids []*string `json:"TableRowUuids,omitnil,omitempty" name:"TableRowUuids"`

	// 任务缓存清理时间（小时）。不填或0表示不清理。
	CacheClearDelay *uint64 `json:"CacheClearDelay,omitnil,omitempty" name:"CacheClearDelay"`

	// 应用版本ID。不填表示使用当前最新版本。
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`

	// WDL运行选项。
	Option *RunOption `json:"Option,omitnil,omitempty" name:"Option"`

	// Nextflow运行选项。
	NFOption *NFOption `json:"NFOption,omitnil,omitempty" name:"NFOption"`

	// 工作目录，使用缓存卷内的相对路径 (暂时仅支持Nextflow)
	WorkDir *string `json:"WorkDir,omitnil,omitempty" name:"WorkDir"`

	// 访问模式，不填默认私有。取值范围
	// - PRIVATE：私有应用
	// - PUBLIC：公共应用
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`
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
	delete(f, "Name")
	delete(f, "EnvironmentId")
	delete(f, "ProjectId")
	delete(f, "Description")
	delete(f, "InputCosUri")
	delete(f, "InputBase64")
	delete(f, "TableId")
	delete(f, "TableRowUuids")
	delete(f, "CacheClearDelay")
	delete(f, "ApplicationVersionId")
	delete(f, "Option")
	delete(f, "NFOption")
	delete(f, "WorkDir")
	delete(f, "AccessMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunApplicationResponseParams struct {
	// 任务批次ID。
	RunGroupId *string `json:"RunGroupId,omitnil,omitempty" name:"RunGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RunGroupId *string `json:"RunGroupId,omitnil,omitempty" name:"RunGroupId"`

	// 项目ID。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 应用ID。
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 应用名称。
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// 应用类型。
	ApplicationType *string `json:"ApplicationType,omitnil,omitempty" name:"ApplicationType"`

	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 环境名称。
	EnvironmentName *string `json:"EnvironmentName,omitnil,omitempty" name:"EnvironmentName"`

	// 表格ID，单例运行为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 任务名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 任务状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务输入。
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// WDL运行选项。
	Option *RunOption `json:"Option,omitnil,omitempty" name:"Option"`

	// Nextflow运行选项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NFOption *NFOption `json:"NFOption,omitnil,omitempty" name:"NFOption"`

	// 任务总数量。
	TotalRun *uint64 `json:"TotalRun,omitnil,omitempty" name:"TotalRun"`

	// 各状态任务的数量。
	RunStatusCounts []*RunStatusCount `json:"RunStatusCounts,omitnil,omitempty" name:"RunStatusCounts"`

	// 执行时间。
	ExecutionTime *ExecutionTime `json:"ExecutionTime,omitnil,omitempty" name:"ExecutionTime"`

	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 创建者。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 创建者ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorId *string `json:"CreatorId,omitnil,omitempty" name:"CreatorId"`

	// 运行结果通知方式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultNotify *string `json:"ResultNotify,omitnil,omitempty" name:"ResultNotify"`

	// 应用版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationVersion *ApplicationVersion `json:"ApplicationVersion,omitnil,omitempty" name:"ApplicationVersion"`
}

type RunMetadata struct {
	// 任务类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunType *string `json:"RunType,omitnil,omitempty" name:"RunType"`

	// 任务ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunId *string `json:"RunId,omitnil,omitempty" name:"RunId"`

	// 父层ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 作业ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 作业名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallName *string `json:"CallName,omitnil,omitempty" name:"CallName"`

	// Scatter索引。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScatterIndex *string `json:"ScatterIndex,omitnil,omitempty" name:"ScatterIndex"`

	// 输入。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// 输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 提交时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitTime *string `json:"SubmitTime,omitnil,omitempty" name:"SubmitTime"`

	// 结束时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 命令行。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 运行时。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Runtime *string `json:"Runtime,omitnil,omitempty" name:"Runtime"`

	// 预处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Preprocess *bool `json:"Preprocess,omitnil,omitempty" name:"Preprocess"`

	// 后处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostProcess *bool `json:"PostProcess,omitnil,omitempty" name:"PostProcess"`

	// Cache命中
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallCached *bool `json:"CallCached,omitnil,omitempty" name:"CallCached"`

	// 标准输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Stdout *string `json:"Stdout,omitnil,omitempty" name:"Stdout"`

	// 错误输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Stderr *string `json:"Stderr,omitnil,omitempty" name:"Stderr"`

	// 其他信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Meta *string `json:"Meta,omitnil,omitempty" name:"Meta"`
}

type RunOption struct {
	// 运行失败模式，取值范围：
	// - ContinueWhilePossible
	// - NoNewCalls
	FailureMode *string `json:"FailureMode,omitnil,omitempty" name:"FailureMode"`

	// 是否使用Call-Caching功能。
	UseCallCache *bool `json:"UseCallCache,omitnil,omitempty" name:"UseCallCache"`

	// 是否使用错误挂起功能。
	UseErrorOnHold *bool `json:"UseErrorOnHold,omitnil,omitempty" name:"UseErrorOnHold"`

	// 输出归档COS路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinalWorkflowOutputsDir *string `json:"FinalWorkflowOutputsDir,omitnil,omitempty" name:"FinalWorkflowOutputsDir"`

	// 是否使用相对目录归档输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseRelativeOutputPaths *bool `json:"UseRelativeOutputPaths,omitnil,omitempty" name:"UseRelativeOutputPaths"`
}

type RunStatusCount struct {
	// 状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 数量。
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

// Predefined struct for user
type RunWorkflowRequestParams struct {
	// 任务批次名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 投递环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 工作流Git仓库信息。
	GitSource *GitInfo `json:"GitSource,omitnil,omitempty" name:"GitSource"`

	// 工作流类型。
	// 
	// 支持类型：
	// - NEXTFLOW
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Nextflow选项。
	NFOption *NFOption `json:"NFOption,omitnil,omitempty" name:"NFOption"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务批次描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 任务输入JSON。需要进行base64编码。
	// （InputBase64和InputCosUri必选其一）
	InputBase64 *string `json:"InputBase64,omitnil,omitempty" name:"InputBase64"`

	// 任务输入COS地址。
	// （InputBase64和InputCosUri必选其一）
	InputCosUri *string `json:"InputCosUri,omitnil,omitempty" name:"InputCosUri"`

	// 任务缓存清理时间（小时）。不填或0表示不清理。
	CacheClearDelay *uint64 `json:"CacheClearDelay,omitnil,omitempty" name:"CacheClearDelay"`

	// 工作目录，使用缓存卷内的相对路径 (暂时仅支持Nextflow)
	WorkDir *string `json:"WorkDir,omitnil,omitempty" name:"WorkDir"`
}

type RunWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 任务批次名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 投递环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 工作流Git仓库信息。
	GitSource *GitInfo `json:"GitSource,omitnil,omitempty" name:"GitSource"`

	// 工作流类型。
	// 
	// 支持类型：
	// - NEXTFLOW
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Nextflow选项。
	NFOption *NFOption `json:"NFOption,omitnil,omitempty" name:"NFOption"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务批次描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 任务输入JSON。需要进行base64编码。
	// （InputBase64和InputCosUri必选其一）
	InputBase64 *string `json:"InputBase64,omitnil,omitempty" name:"InputBase64"`

	// 任务输入COS地址。
	// （InputBase64和InputCosUri必选其一）
	InputCosUri *string `json:"InputCosUri,omitnil,omitempty" name:"InputCosUri"`

	// 任务缓存清理时间（小时）。不填或0表示不清理。
	CacheClearDelay *uint64 `json:"CacheClearDelay,omitnil,omitempty" name:"CacheClearDelay"`

	// 工作目录，使用缓存卷内的相对路径 (暂时仅支持Nextflow)
	WorkDir *string `json:"WorkDir,omitnil,omitempty" name:"WorkDir"`
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
	delete(f, "WorkDir")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunWorkflowResponseParams struct {
	// 任务批次ID。
	RunGroupId *string `json:"RunGroupId,omitnil,omitempty" name:"RunGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`
}

type StorageOption struct {
	// 文件存储类型，取值范围：
	// - SD：通用标准型
	// - HP：通用性能型
	// - TB：turbo标准型
	// - TP：turbo性能型
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 文件存储可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 文件系统容量，turbo系列必填，单位为GiB。 
	// - turbo标准型起售40TiB，即40960GiB；扩容步长20TiB，即20480 GiB。
	// - turbo性能型起售20TiB，即20480 GiB；扩容步长10TiB，即10240 GiB。
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`
}

type Table struct {
	// 表格ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 关联项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 表格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 表格描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 表格列
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*TableColumn `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`
}

type TableColumn struct {
	// 列名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Header *string `json:"Header,omitnil,omitempty" name:"Header"`

	// 列数据类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataType *string `json:"DataType,omitnil,omitempty" name:"DataType"`
}

type TableRow struct {
	// 表格行UUID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableRowUuid *string `json:"TableRowUuid,omitnil,omitempty" name:"TableRowUuid"`

	// 表格行内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*string `json:"Content,omitnil,omitempty" name:"Content"`
}

// Predefined struct for user
type TerminateRunGroupRequestParams struct {
	// 任务批次ID。
	RunGroupId *string `json:"RunGroupId,omitnil,omitempty" name:"RunGroupId"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type TerminateRunGroupRequest struct {
	*tchttp.BaseRequest
	
	// 任务批次ID。
	RunGroupId *string `json:"RunGroupId,omitnil,omitempty" name:"RunGroupId"`

	// 项目ID。
	// （不填使用指定地域下的默认项目）
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	VPCId *string `json:"VPCId,omitnil,omitempty" name:"VPCId"`

	// 子网ID（SubnetId和SubnetZone&SubnetCIDRBlock必选其一。若使用SubnetId，则使用现用子网；若使用SubnetZone&SubnetCIDRBlock，则创建新的子网）
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 子网可用区。
	SubnetZone *string `json:"SubnetZone,omitnil,omitempty" name:"SubnetZone"`

	// 私有网络CIDR。
	VPCCIDRBlock *string `json:"VPCCIDRBlock,omitnil,omitempty" name:"VPCCIDRBlock"`

	// 子网CIDR。
	SubnetCIDRBlock *string `json:"SubnetCIDRBlock,omitnil,omitempty" name:"SubnetCIDRBlock"`
}

type Volume struct {
	// 缓存卷ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeId *string `json:"VolumeId,omitnil,omitempty" name:"VolumeId"`

	// 名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 环境ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 缓存卷类型，取值范围：
	// * SHARED：多点挂载共享存储
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 缓存卷规格，取值范围：
	// 
	// - SD：通用标准型
	// - HP：通用性能型
	// - TB：turbo标准型
	// - TP：turbo性能型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// 缓存卷大小（GB）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`

	// 缓存卷使用量（Byte）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Usage *uint64 `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 缓存卷吞吐上限（MiB/s）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthLimit *float64 `json:"BandwidthLimit,omitnil,omitempty" name:"BandwidthLimit"`

	// 默认挂载路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultMountPath *string `json:"DefaultMountPath,omitnil,omitempty" name:"DefaultMountPath"`

	// 是否为默认缓存卷。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}