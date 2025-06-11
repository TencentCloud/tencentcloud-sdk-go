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

package v20230812

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ApplicationInfo struct {
	// 应用id
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 应用名称
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// 应用描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 应用的环境配置
	ConfigEnvironment *string `json:"ConfigEnvironment,omitnil,omitempty" name:"ConfigEnvironment"`

	// 系统盘大小下限，单位GB
	MinSystemDiskSize *int64 `json:"MinSystemDiskSize,omitnil,omitempty" name:"MinSystemDiskSize"`

	// 应用类型，目前该项取值可以为PUBLIC_APPLICATION（公共应用）；PRIVATE_APPLICATION（自定义应用）；COMMUNITY_APPLICATION（社区应用）
	ApplicationType *string `json:"ApplicationType,omitnil,omitempty" name:"ApplicationType"`

	// 应用状态：CREATING-创建中；ONLINE -正常在线；DELETING -删除中；ARREARS - 欠费隔离
	// 示例值：ONLINE
	ApplicationState *string `json:"ApplicationState,omitnil,omitempty" name:"ApplicationState"`

	// 应用创建时间，格式：%Y-%m-%d %H:%M:%S
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 应用大小，单位GB
	ApplicationSize *int64 `json:"ApplicationSize,omitnil,omitempty" name:"ApplicationSize"`
}

// Predefined struct for user
type CreateApplicationRequestParams struct {
	// 需要制作自定义应用的HAI实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自定义应用的应用名称
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// 自定义应用的描述
	ApplicationDescription *string `json:"ApplicationDescription,omitnil,omitempty" name:"ApplicationDescription"`
}

type CreateApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 需要制作自定义应用的HAI实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自定义应用的应用名称
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// 自定义应用的描述
	ApplicationDescription *string `json:"ApplicationDescription,omitnil,omitempty" name:"ApplicationDescription"`
}

func (r *CreateApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ApplicationName")
	delete(f, "ApplicationDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationResponseParams struct {
	// HAI自定义应用ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateApplicationResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationResponseParams `json:"Response"`
}

func (r *CreateApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMuskPromptRequestParams struct {
	// workgroup id
	WorkgroupId *string `json:"WorkgroupId,omitnil,omitempty" name:"WorkgroupId"`

	// workflow id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// prompt 参数
	PromptParams *string `json:"PromptParams,omitnil,omitempty" name:"PromptParams"`
}

type CreateMuskPromptRequest struct {
	*tchttp.BaseRequest
	
	// workgroup id
	WorkgroupId *string `json:"WorkgroupId,omitnil,omitempty" name:"WorkgroupId"`

	// workflow id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// prompt 参数
	PromptParams *string `json:"PromptParams,omitnil,omitempty" name:"PromptParams"`
}

func (r *CreateMuskPromptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMuskPromptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkgroupId")
	delete(f, "WorkflowId")
	delete(f, "PromptParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMuskPromptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMuskPromptResponseParams struct {
	// prompt id
	PromptId *string `json:"PromptId,omitnil,omitempty" name:"PromptId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMuskPromptResponse struct {
	*tchttp.BaseResponse
	Response *CreateMuskPromptResponseParams `json:"Response"`
}

func (r *CreateMuskPromptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMuskPromptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationsRequestParams struct {
	// 应用id列表。单次请求数量上限为100个。
	ApplicationIds []*string `json:"ApplicationIds,omitnil,omitempty" name:"ApplicationIds"`

	// 过滤器，跟ApplicationIds不能共用，支持的filter主要有：application-id: 精确匹配;scene-id: 精确匹配，通过调用接口 [DescribeScenes](https://cloud.tencent.com/document/api/1721/101608)获取;application-name: 模糊匹配;application-type: 精确匹配，枚举类型如下：PUBLIC_APPLICATION（公共应用）/ PRIVATE_APPLICATION（自定义应用）/ COMMUNITY_APPLICATION（社区应用）;
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，不得小于0，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回量，不得大于100，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 应用列表排序的依据字段。取值范围："CREATED_TIME"：依据应用的创建时间排序。 "APPLICATION_SIZE"：依据应用的大小排序。默认按应用的创建时间排序。
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 输出应用列表的排列顺序。取值范围："ASC"：升序排列。 "DESC"：降序排列。默认按降序排列。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeApplicationsRequest struct {
	*tchttp.BaseRequest
	
	// 应用id列表。单次请求数量上限为100个。
	ApplicationIds []*string `json:"ApplicationIds,omitnil,omitempty" name:"ApplicationIds"`

	// 过滤器，跟ApplicationIds不能共用，支持的filter主要有：application-id: 精确匹配;scene-id: 精确匹配，通过调用接口 [DescribeScenes](https://cloud.tencent.com/document/api/1721/101608)获取;application-name: 模糊匹配;application-type: 精确匹配，枚举类型如下：PUBLIC_APPLICATION（公共应用）/ PRIVATE_APPLICATION（自定义应用）/ COMMUNITY_APPLICATION（社区应用）;
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，不得小于0，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回量，不得大于100，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 应用列表排序的依据字段。取值范围："CREATED_TIME"：依据应用的创建时间排序。 "APPLICATION_SIZE"：依据应用的大小排序。默认按应用的创建时间排序。
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 输出应用列表的排列顺序。取值范围："ASC"：升序排列。 "DESC"：降序排列。默认按降序排列。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
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
	delete(f, "ApplicationIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationsResponseParams struct {
	// 应用总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 分页返回的应用列表
	ApplicationSet []*ApplicationInfo `json:"ApplicationSet,omitnil,omitempty" name:"ApplicationSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeInstanceNetworkStatusRequestParams struct {
	// 实例ID数组，单次请求最多不超过100个实例；实例ID通过调用接口[DescribeInstances](https://cloud.tencent.com/document/api/1721/101612)获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeInstanceNetworkStatusRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID数组，单次请求最多不超过100个实例；实例ID通过调用接口[DescribeInstances](https://cloud.tencent.com/document/api/1721/101612)获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DescribeInstanceNetworkStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNetworkStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceNetworkStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNetworkStatusResponseParams struct {
	// 查询结果集长度
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 查询结果集
	NetworkStatusSet []*NetworkStatus `json:"NetworkStatusSet,omitnil,omitempty" name:"NetworkStatusSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceNetworkStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceNetworkStatusResponseParams `json:"Response"`
}

func (r *DescribeInstanceNetworkStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNetworkStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 实例元组，数量上限100
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 描述键值对过滤器，用于条件过滤查询。目前支持的过滤器有： instance-id，实例id； instance-state，实例状态：RUNNING，PENDING，STOPPED，ARREARS，STOPPED_NO_CHARGE； charge-type，付费方式：PREPAID_BY_MONTH，POSTPAID_BY_HOUR； public-ip-address，公网IP过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0，不得大于100
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回量，默认为20，不能小于0
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例元组，数量上限100
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 描述键值对过滤器，用于条件过滤查询。目前支持的过滤器有： instance-id，实例id； instance-state，实例状态：RUNNING，PENDING，STOPPED，ARREARS，STOPPED_NO_CHARGE； charge-type，付费方式：PREPAID_BY_MONTH，POSTPAID_BY_HOUR； public-ip-address，公网IP过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0，不得大于100
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回量，默认为20，不能小于0
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	delete(f, "InstanceIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 实例总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 分页实例详情
	InstanceSet []*Instance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeMuskPromptsRequestParams struct {
	// workgroup id
	WorkgroupId *string `json:"WorkgroupId,omitnil,omitempty" name:"WorkgroupId"`

	// workflow id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// offset 
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤参数 支持过滤的键值： PromptId，Status
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeMuskPromptsRequest struct {
	*tchttp.BaseRequest
	
	// workgroup id
	WorkgroupId *string `json:"WorkgroupId,omitnil,omitempty" name:"WorkgroupId"`

	// workflow id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// offset 
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤参数 支持过滤的键值： PromptId，Status
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeMuskPromptsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMuskPromptsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkgroupId")
	delete(f, "WorkflowId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMuskPromptsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMuskPromptsResponseParams struct {
	// total count
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// prompt列表详情
	MuskPromptInfos []*MuskPromptInfo `json:"MuskPromptInfos,omitnil,omitempty" name:"MuskPromptInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMuskPromptsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMuskPromptsResponseParams `json:"Response"`
}

func (r *DescribeMuskPromptsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMuskPromptsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsRequestParams struct {

}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsResponseParams struct {
	// 地域列表
	RegionSet []*RegionInfo `json:"RegionSet,omitnil,omitempty" name:"RegionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegionsResponseParams `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScenesRequestParams struct {
	// 场景id列表，单次能查询100个场景id
	SceneIds []*string `json:"SceneIds,omitnil,omitempty" name:"SceneIds"`
}

type DescribeScenesRequest struct {
	*tchttp.BaseRequest
	
	// 场景id列表，单次能查询100个场景id
	SceneIds []*string `json:"SceneIds,omitnil,omitempty" name:"SceneIds"`
}

func (r *DescribeScenesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScenesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SceneIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScenesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScenesResponseParams struct {
	// 场景详情
	SceneSet []*SceneInfo `json:"SceneSet,omitnil,omitempty" name:"SceneSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScenesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScenesResponseParams `json:"Response"`
}

func (r *DescribeScenesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScenesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceLoginSettingsRequestParams struct {
	// 实例ID通过调用接口[DescribeInstances](https://cloud.tencent.com/document/api/1721/101612)获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务名称
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`
}

type DescribeServiceLoginSettingsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID通过调用接口[DescribeInstances](https://cloud.tencent.com/document/api/1721/101612)获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务名称
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`
}

func (r *DescribeServiceLoginSettingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceLoginSettingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ServiceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceLoginSettingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceLoginSettingsResponseParams struct {
	// 服务登录配置详情
	LoginSettings []*LoginSetting `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServiceLoginSettingsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceLoginSettingsResponseParams `json:"Response"`
}

func (r *DescribeServiceLoginSettingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceLoginSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 需要过滤的字段。	
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type InquirePriceRunInstancesRequestParams struct {
	// 应用ID通过调用接口[DescribeApplications](https://cloud.tencent.com/document/api/1721/101609)获取。
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 算力套餐类型, 枚举：XL,XL_2X, 3XL, 3XL_2X, 4XL, 24GB_A.
	BundleType *string `json:"BundleType,omitnil,omitempty" name:"BundleType"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 购买实例数量，单次请求实例数量上限为10。
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 实例显示名称，名称长度限制为128个字符。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 幂等请求token
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// DryRun为True就是只验接口连通性，默认为False
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 付费方式，POSTPAID_BY_HOUR按量后付费，PREPAID_BY_MONTH预付费按月，PREPAID_BY_DAY预付费按天
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 预付费参数
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`
}

type InquirePriceRunInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID通过调用接口[DescribeApplications](https://cloud.tencent.com/document/api/1721/101609)获取。
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 算力套餐类型, 枚举：XL,XL_2X, 3XL, 3XL_2X, 4XL, 24GB_A.
	BundleType *string `json:"BundleType,omitnil,omitempty" name:"BundleType"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 购买实例数量，单次请求实例数量上限为10。
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 实例显示名称，名称长度限制为128个字符。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 幂等请求token
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// DryRun为True就是只验接口连通性，默认为False
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 付费方式，POSTPAID_BY_HOUR按量后付费，PREPAID_BY_MONTH预付费按月，PREPAID_BY_DAY预付费按天
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 预付费参数
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`
}

func (r *InquirePriceRunInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRunInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "BundleType")
	delete(f, "SystemDisk")
	delete(f, "InstanceCount")
	delete(f, "InstanceName")
	delete(f, "ClientToken")
	delete(f, "DryRun")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceRunInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRunInstancesResponseParams struct {
	// 发货参数对应的价格组合，当DryRun=True，会返回空
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceRunInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceRunInstancesResponseParams `json:"Response"`
}

func (r *InquirePriceRunInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRunInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Instance struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例状态：
	// PENDING：表示创建中
	// LAUNCH_FAILED：表示创建失败
	// RUNNING：表示运行中
	// ARREARS：表示待回收
	// STOPPED_NO_CHARGE：表示关机不收费
	// TERMINATING：表示销毁中
	// TERMINATED：表示已销毁
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// 应用名称
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// 算力套餐名称
	BundleName *string `json:"BundleName,omitnil,omitempty" name:"BundleName"`

	// 实例所包含的GPU卡数
	GPUCount *uint64 `json:"GPUCount,omitnil,omitempty" name:"GPUCount"`

	// 算力
	GPUPerformance *string `json:"GPUPerformance,omitnil,omitempty" name:"GPUPerformance"`

	// 显存，单位：GB
	GPUMemory *string `json:"GPUMemory,omitnil,omitempty" name:"GPUMemory"`

	// CPU核数，单位：核
	CPU *string `json:"CPU,omitnil,omitempty" name:"CPU"`

	// 内存，单位：GB
	Memory *string `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 系统盘数据
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 内网ip地址
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// 公网ip地址
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// 安全组ID
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 实例最新操作
	LatestOperation *string `json:"LatestOperation,omitnil,omitempty" name:"LatestOperation"`

	// 实例最新操作状态：
	// SUCCESS：表示操作成功
	// OPERATING：表示操作执行中
	// FAILED：表示操作失败
	LatestOperationState *string `json:"LatestOperationState,omitnil,omitempty" name:"LatestOperationState"`

	// 实例创建时间，时间格式："YYYY-MM-DD HH:MM:SS"
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 公网出带宽上限，默认10Mbps，单位：Mbps
	MaxOutBandwidth *string `json:"MaxOutBandwidth,omitnil,omitempty" name:"MaxOutBandwidth"`

	// 每月免费流量，默认500G，单位：GB
	MaxFreeTraffic *string `json:"MaxFreeTraffic,omitnil,omitempty" name:"MaxFreeTraffic"`

	// 应用配置环境
	ConfigurationEnvironment *string `json:"ConfigurationEnvironment,omitnil,omitempty" name:"ConfigurationEnvironment"`

	// 实例包含的登录服务详情
	LoginServices []*LoginService `json:"LoginServices,omitnil,omitempty" name:"LoginServices"`

	// 应用服务的操作系统类型；参数：linux、windows
	OSType *string `json:"OSType,omitnil,omitempty" name:"OSType"`
}

type InstanceChargePrepaid struct {
	// 时长，默认值：1
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 续费标志可选参数：
	// NOTIFY_AND_MANUAL_RENEW：表示默认状态(用户未设置，即初始状态：若用户有预付费不停服特权，也会对该值进行自动续费)
	// NOTIFY_AND_AUTO_RENEW：表示自动续费
	// DISABLE_NOTIFY_AND_MANUAL_RENEW：表示明确不自动续费(用户设置)
	// 默认值：NOTIFY_AND_MANUAL_RENEW
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 时长单位，枚举： MONTH, DAY, HOUR；释义：月，日，小时
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
}

type ItemPrice struct {
	// 原单价，元
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// 折扣后单价，元
	DiscountUnitPrice *float64 `json:"DiscountUnitPrice,omitnil,omitempty" name:"DiscountUnitPrice"`

	// 折扣
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 单位：时/月
	ChargeUnit *string `json:"ChargeUnit,omitnil,omitempty" name:"ChargeUnit"`

	// 商品数量
	Amount *uint64 `json:"Amount,omitnil,omitempty" name:"Amount"`
}

type ItemPriceDetail struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例价格详情
	InstancePrice *ItemPrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// 磁盘价格详情
	CloudDiskPrice *ItemPrice `json:"CloudDiskPrice,omitnil,omitempty" name:"CloudDiskPrice"`

	// 该实例的总价钱
	InstanceTotalPrice *ItemPrice `json:"InstanceTotalPrice,omitnil,omitempty" name:"InstanceTotalPrice"`
}

type LoginService struct {
	// 登录方式名称
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`
}

type LoginSetting struct {
	// 服务名称
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 服务登录url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type MuskPromptInfo struct {
	// workflow id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// workgroup id
	WorkgroupId *string `json:"WorkgroupId,omitnil,omitempty" name:"WorkgroupId"`

	// prompt id
	PromptId *string `json:"PromptId,omitnil,omitempty" name:"PromptId"`

	// 生成的内容
	OutputResource []*string `json:"OutputResource,omitnil,omitempty" name:"OutputResource"`

	// prompt status 
	// 0: 执行中
	// 1: 执行成功
	// 2: 执行失败
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 任务执行耗时，单位毫秒
	Cost *uint64 `json:"Cost,omitnil,omitempty" name:"Cost"`

	// 任务执行失败错误信息
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`
}

type NetworkStatus struct {
	// HAI 的实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 公网 IP 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressIp *string `json:"AddressIp,omitnil,omitempty" name:"AddressIp"`

	// 出带宽上限，单位Mbps
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bandwidth *uint64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 流量包总量，单位GB
	TotalTrafficAmount *float64 `json:"TotalTrafficAmount,omitnil,omitempty" name:"TotalTrafficAmount"`

	// 流量包剩余量，单位GB
	RemainingTrafficAmount *float64 `json:"RemainingTrafficAmount,omitnil,omitempty" name:"RemainingTrafficAmount"`
}

type Price struct {
	// 实例价格信息
	InstancePrice *ItemPrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// 云盘价格信息
	CloudDiskPrice *ItemPrice `json:"CloudDiskPrice,omitnil,omitempty" name:"CloudDiskPrice"`

	// 分实例价格
	PriceDetailSet []*ItemPriceDetail `json:"PriceDetailSet,omitnil,omitempty" name:"PriceDetailSet"`
}

type RegionInfo struct {
	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 地域名称
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 地域是否可用状态
	// AVAILABLE：可用
	RegionState *string `json:"RegionState,omitnil,omitempty" name:"RegionState"`

	// 学术加速是否支持：
	// NO_NEED_SUPPORT表示不需支持；NOT_SUPPORT_YET表示暂未支持；ALREADY_SUPPORT表示已经支持。
	ScholarRocketSupportState *string `json:"ScholarRocketSupportState,omitnil,omitempty" name:"ScholarRocketSupportState"`
}

// Predefined struct for user
type ResetInstancesPasswordRequestParams struct {
	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例密码必须8-30位，推荐使用12位以上密码，不能以“/”开头，至少包含以下字符中的三种不同字符，字符种类：<br><li>小写字母：[a-z]</li><br><li>大写字母：[A-Z]</li><br><li>数字：0-9</li><br><li>特殊字符： ()\`\~!@#$%^&\*-+=\_|{}[]:;'<>,.?/</li>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 默认为False，True代表只验证接口连通性
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type ResetInstancesPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例密码必须8-30位，推荐使用12位以上密码，不能以“/”开头，至少包含以下字符中的三种不同字符，字符种类：<br><li>小写字母：[a-z]</li><br><li>大写字母：[A-Z]</li><br><li>数字：0-9</li><br><li>特殊字符： ()\`\~!@#$%^&\*-+=\_|{}[]:;'<>,.?/</li>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 默认为False，True代表只验证接口连通性
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *ResetInstancesPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Password")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstancesPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesPasswordResponseParams struct {
	// task任务id
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetInstancesPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetInstancesPasswordResponseParams `json:"Response"`
}

func (r *ResetInstancesPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeInstanceDiskRequestParams struct {
	// 需要扩容云盘的HAI实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 扩容云硬盘大小，单位为GB，必须大于当前云硬盘大小。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

type ResizeInstanceDiskRequest struct {
	*tchttp.BaseRequest
	
	// 需要扩容云盘的HAI实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 扩容云硬盘大小，单位为GB，必须大于当前云硬盘大小。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

func (r *ResizeInstanceDiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeInstanceDiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DiskSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResizeInstanceDiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeInstanceDiskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResizeInstanceDiskResponse struct {
	*tchttp.BaseResponse
	Response *ResizeInstanceDiskResponseParams `json:"Response"`
}

func (r *ResizeInstanceDiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeInstanceDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunInstancesRequestParams struct {
	// 应用ID通过调用接口[DescribeApplications](https://cloud.tencent.com/document/api/1721/101609)获取。
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 算力套餐类型, 枚举：XL,XL_2X, 3XL, 3XL_2X, 4XL, 24GB_A
	BundleType *string `json:"BundleType,omitnil,omitempty" name:"BundleType"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 购买实例数量，单次请求实例数量上限为10.
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 实例显示名称，名称长度限制为128个字符.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 幂等请求的token
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// DryRun为True就是只验接口连通性，默认为False
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type RunInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID通过调用接口[DescribeApplications](https://cloud.tencent.com/document/api/1721/101609)获取。
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 算力套餐类型, 枚举：XL,XL_2X, 3XL, 3XL_2X, 4XL, 24GB_A
	BundleType *string `json:"BundleType,omitnil,omitempty" name:"BundleType"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 购买实例数量，单次请求实例数量上限为10.
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 实例显示名称，名称长度限制为128个字符.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 幂等请求的token
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// DryRun为True就是只验接口连通性，默认为False
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *RunInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "BundleType")
	delete(f, "SystemDisk")
	delete(f, "InstanceCount")
	delete(f, "InstanceName")
	delete(f, "ClientToken")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunInstancesResponseParams struct {
	// 实例ID列表
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RunInstancesResponseParams `json:"Response"`
}

func (r *RunInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SceneInfo struct {
	// 场景id
	SceneId *string `json:"SceneId,omitnil,omitempty" name:"SceneId"`

	// 场景名
	SceneName *string `json:"SceneName,omitnil,omitempty" name:"SceneName"`
}

// Predefined struct for user
type StartInstanceRequestParams struct {
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1721/101612) API获取实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 默认为False，True代表只验证接口连通性
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type StartInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1721/101612) API获取实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 默认为False，True代表只验证接口连通性
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *StartInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartInstanceResponseParams struct {
	// task任务id
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartInstanceResponse struct {
	*tchttp.BaseResponse
	Response *StartInstanceResponseParams `json:"Response"`
}

func (r *StartInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstanceRequestParams struct {
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1721/101612) API获取实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// hai实例关机的模式，目前仅支持关机不收费：
	// STOP_CHARGE -- 关闭hai实例，释放计算资源，停止收取计算资源的费用。
	// 注意：默认值为STOP_CHARGE
	StopMode *string `json:"StopMode,omitnil,omitempty" name:"StopMode"`

	// 默认为False，True代表只验证接口连通性
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type StopInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1721/101612) API获取实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// hai实例关机的模式，目前仅支持关机不收费：
	// STOP_CHARGE -- 关闭hai实例，释放计算资源，停止收取计算资源的费用。
	// 注意：默认值为STOP_CHARGE
	StopMode *string `json:"StopMode,omitnil,omitempty" name:"StopMode"`

	// 默认为False，True代表只验证接口连通性
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *StopInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StopMode")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstanceResponseParams struct {
	// task任务id
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopInstanceResponse struct {
	*tchttp.BaseResponse
	Response *StopInstanceResponseParams `json:"Response"`
}

func (r *StopInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemDisk struct {
	// 系统盘类型。取值范围：<li>CLOUD_PREMIUM：高性能云硬盘</li><li>CLOUD_HSSD：增强型SSD云盘</li>默认取值：当前有库存的硬盘类型。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 系统盘大小，单位：GB。默认值为 80，取值范围：80-1000
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 系统盘分区盘符
	DiskName *string `json:"DiskName,omitnil,omitempty" name:"DiskName"`
}

// Predefined struct for user
type TerminateInstancesRequestParams struct {
	// 实例ID列表。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1721/101612) API获取实例ID列表。单次能查询100个InstanceId。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 默认为False，True代表只验证接口连通性
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1721/101612) API获取实例ID列表。单次能查询100个InstanceId。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 默认为False，True代表只验证接口连通性
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *TerminateInstancesResponseParams `json:"Response"`
}

func (r *TerminateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}