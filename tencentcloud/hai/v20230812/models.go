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
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitnil" name:"ApplicationName"`

	// 应用描述
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 应用的环境配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigEnvironment *string `json:"ConfigEnvironment,omitnil" name:"ConfigEnvironment"`

	// 系统盘大小下限
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinSystemDiskSize *int64 `json:"MinSystemDiskSize,omitnil" name:"MinSystemDiskSize"`
}

// Predefined struct for user
type DescribeApplicationsRequestParams struct {
	// 应用id列表
	ApplicationIds []*string `json:"ApplicationIds,omitnil" name:"ApplicationIds"`

	// 过滤器，跟ApplicationIds不能共用，支持的filter主要有：
	// application-id，精确匹配
	// scene-id，精确匹配
	// application-name，模糊匹配
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回量，默认为20
	// MC：1000
	// 用户：100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeApplicationsRequest struct {
	*tchttp.BaseRequest
	
	// 应用id列表
	ApplicationIds []*string `json:"ApplicationIds,omitnil" name:"ApplicationIds"`

	// 过滤器，跟ApplicationIds不能共用，支持的filter主要有：
	// application-id，精确匹配
	// scene-id，精确匹配
	// application-name，模糊匹配
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回量，默认为20
	// MC：1000
	// 用户：100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationsResponseParams struct {
	// 应用总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 分页返回的应用列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationSet []*ApplicationInfo `json:"ApplicationSet,omitnil" name:"ApplicationSet"`

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
type DescribeInstanceNetworkStatusRequestParams struct {
	// 实例ID数组，单次请求最多不超过100个实例
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type DescribeInstanceNetworkStatusRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID数组，单次请求最多不超过100个实例
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 查询结果集
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkStatusSet []*NetworkStatus `json:"NetworkStatusSet,omitnil" name:"NetworkStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 实例元组
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 描述键值对过滤器，用于条件过滤查询。目前支持的过滤器有：instance-id，实例id；instance-state，实例状态
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回量，默认为20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例元组
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 描述键值对过滤器，用于条件过滤查询。目前支持的过滤器有：instance-id，实例id；instance-state，实例状态
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回量，默认为20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
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
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 分页实例详情
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceSet []*Instance `json:"InstanceSet,omitnil" name:"InstanceSet"`

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
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionSet []*RegionInfo `json:"RegionSet,omitnil" name:"RegionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 场景id列表
	SceneIds []*string `json:"SceneIds,omitnil" name:"SceneIds"`
}

type DescribeScenesRequest struct {
	*tchttp.BaseRequest
	
	// 场景id列表
	SceneIds []*string `json:"SceneIds,omitnil" name:"SceneIds"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneSet []*SceneInfo `json:"SceneSet,omitnil" name:"SceneSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 服务名称
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`
}

type DescribeServiceLoginSettingsRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 服务名称
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoginSettings []*LoginSetting `json:"LoginSettings,omitnil" name:"LoginSettings"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitnil" name:"Values"`
}

// Predefined struct for user
type InquirePriceRunInstancesRequestParams struct {
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 算力套餐类型
	BundleType *string `json:"BundleType,omitnil" name:"BundleType"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil" name:"SystemDisk"`

	// 购买实例数量。
	InstanceCount *uint64 `json:"InstanceCount,omitnil" name:"InstanceCount"`

	// 实例显示名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 幂等请求token
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// DryRun为True就是只验接口连通性，默认为False
	DryRun *bool `json:"DryRun,omitnil" name:"DryRun"`
}

type InquirePriceRunInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 算力套餐类型
	BundleType *string `json:"BundleType,omitnil" name:"BundleType"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil" name:"SystemDisk"`

	// 购买实例数量。
	InstanceCount *uint64 `json:"InstanceCount,omitnil" name:"InstanceCount"`

	// 实例显示名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 幂等请求token
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// DryRun为True就是只验接口连通性，默认为False
	DryRun *bool `json:"DryRun,omitnil" name:"DryRun"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceRunInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRunInstancesResponseParams struct {
	// 发货参数对应的价格组合，当DryRun=True，会返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Price *Price `json:"Price,omitnil" name:"Price"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 实例状态：
	// PENDING：表示创建中
	// LAUNCH_FAILED：表示创建失败
	// RUNNING：表示运行中
	// ARREAR：表示欠费隔离
	// TERMINATING：表示销毁中。
	// TERMINATED：表示已销毁
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceState *string `json:"InstanceState,omitnil" name:"InstanceState"`

	// 应用名称
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitnil" name:"ApplicationName"`

	// 算力套餐名称
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleName *string `json:"BundleName,omitnil" name:"BundleName"`

	// 实例所包含的GPU卡数
	// 注意：此字段可能返回 null，表示取不到有效值。
	GPUCount *uint64 `json:"GPUCount,omitnil" name:"GPUCount"`

	// 算力
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	GPUPerformance *string `json:"GPUPerformance,omitnil" name:"GPUPerformance"`

	// 显存
	// 注意：此字段可能返回 null，表示取不到有效值。
	GPUMemory *string `json:"GPUMemory,omitnil" name:"GPUMemory"`

	// CPU核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CPU *string `json:"CPU,omitnil" name:"CPU"`

	// 内存
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *string `json:"Memory,omitnil" name:"Memory"`

	// 系统盘数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil" name:"SystemDisk"`

	// 内网ip地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil" name:"PrivateIpAddresses"`

	// 公网ip地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil" name:"PublicIpAddresses"`

	// 安全组ID
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// 实例最新操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperation *string `json:"LatestOperation,omitnil" name:"LatestOperation"`

	// 实例最新操作状态：
	// SUCCESS：表示操作成功
	// OPERATING：表示操作执行中
	// FAILED：表示操作失败
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperationState *string `json:"LatestOperationState,omitnil" name:"LatestOperationState"`

	// 实例创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 公网出带宽上限，默认5Mbps
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxOutBandwidth *string `json:"MaxOutBandwidth,omitnil" name:"MaxOutBandwidth"`

	// 每月免费流量，默认500G
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxFreeTraffic *string `json:"MaxFreeTraffic,omitnil" name:"MaxFreeTraffic"`

	// 应用配置环境
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigurationEnvironment *string `json:"ConfigurationEnvironment,omitnil" name:"ConfigurationEnvironment"`

	// 实例包含的登录服务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoginServices []*LoginService `json:"LoginServices,omitnil" name:"LoginServices"`
}

type ItemPrice struct {
	// 原单价
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPrice *float64 `json:"UnitPrice,omitnil" name:"UnitPrice"`

	// 折扣后单价
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountUnitPrice *float64 `json:"DiscountUnitPrice,omitnil" name:"DiscountUnitPrice"`

	// 折扣
	// 注意：此字段可能返回 null，表示取不到有效值。
	Discount *float64 `json:"Discount,omitnil" name:"Discount"`

	// 单位：时
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeUnit *string `json:"ChargeUnit,omitnil" name:"ChargeUnit"`

	// 商品数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Amount *uint64 `json:"Amount,omitnil" name:"Amount"`
}

type LoginService struct {
	// 登录方式名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`
}

type LoginSetting struct {
	// 服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`

	// 服务登录url
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil" name:"Url"`
}

type NetworkStatus struct {
	// HAI 的实例 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 公网 IP 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressIp *string `json:"AddressIp,omitnil" name:"AddressIp"`

	// 出带宽上限，单位Mbps
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bandwidth *uint64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// 流量包总量，单位GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalTrafficAmount *float64 `json:"TotalTrafficAmount,omitnil" name:"TotalTrafficAmount"`

	// 流量包剩余量，单位GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemainingTrafficAmount *float64 `json:"RemainingTrafficAmount,omitnil" name:"RemainingTrafficAmount"`
}

type Price struct {
	// 实例价格信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstancePrice *ItemPrice `json:"InstancePrice,omitnil" name:"InstancePrice"`

	// 云盘价格信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloudDiskPrice *ItemPrice `json:"CloudDiskPrice,omitnil" name:"CloudDiskPrice"`
}

type RegionInfo struct {
	// ap-guangzhou
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 华南地区(广州)
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`

	// 地域是否可用状态
	// AVAILABLE：可用
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionState *string `json:"RegionState,omitnil" name:"RegionState"`

	// 学术加速是否支持：
	// NO_NEED_SUPPORT表示不需支持；NOT_SUPPORT_YET表示暂未支持；ALREADY_SUPPORT表示已经支持。对于ALREADY_SUPPORT的地域才需进一步调用DescribeScholarRocketStatus查看学术加速是开启还是关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScholarRocketSupportState *string `json:"ScholarRocketSupportState,omitnil" name:"ScholarRocketSupportState"`
}

// Predefined struct for user
type RunInstancesRequestParams struct {
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 算力套餐类型
	BundleType *string `json:"BundleType,omitnil" name:"BundleType"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil" name:"SystemDisk"`

	// 购买实例数量。
	InstanceCount *uint64 `json:"InstanceCount,omitnil" name:"InstanceCount"`

	// 实例显示名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 幂等请求的token
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// DryRun为True就是只验接口连通性，默认为False
	DryRun *bool `json:"DryRun,omitnil" name:"DryRun"`
}

type RunInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 算力套餐类型
	BundleType *string `json:"BundleType,omitnil" name:"BundleType"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil" name:"SystemDisk"`

	// 购买实例数量。
	InstanceCount *uint64 `json:"InstanceCount,omitnil" name:"InstanceCount"`

	// 实例显示名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 幂等请求的token
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// DryRun为True就是只验接口连通性，默认为False
	DryRun *bool `json:"DryRun,omitnil" name:"DryRun"`
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
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil" name:"InstanceIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneId *string `json:"SceneId,omitnil" name:"SceneId"`

	// 场景名
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneName *string `json:"SceneName,omitnil" name:"SceneName"`
}

type SystemDisk struct {
	// 系统盘类型。系统盘类型限制详见[存储概述](https://cloud.tencent.com/document/product/213/4952)。取值范围：<br><li>LOCAL_BASIC：本地硬盘<br><li>LOCAL_SSD：本地SSD硬盘<br><li>CLOUD_BASIC：普通云硬盘<br><li>CLOUD_SSD：SSD云硬盘<br><li>CLOUD_PREMIUM：高性能云硬盘<br><li>CLOUD_BSSD：通用性SSD云硬盘<br><br>默认取值：当前有库存的硬盘类型。	
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// 系统盘大小，单位：GB。默认值为 80
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// 系统盘分区盘符
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskName *string `json:"DiskName,omitnil" name:"DiskName"`
}

// Predefined struct for user
type TerminateInstancesRequestParams struct {
	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 默认为False，True代表只验证接口连通性
	DryRun *bool `json:"DryRun,omitnil" name:"DryRun"`
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 默认为False，True代表只验证接口连通性
	DryRun *bool `json:"DryRun,omitnil" name:"DryRun"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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