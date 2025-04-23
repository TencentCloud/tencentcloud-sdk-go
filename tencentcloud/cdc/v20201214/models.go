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

package v20201214

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type CbsInfo struct {
	// cbs存储大小，单位TB
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// cbs存储类型，默认为SSD
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type CosCapacity struct {
	// 已购cos的总容量大小，单位GB
	TotalCapacity *float64 `json:"TotalCapacity,omitnil,omitempty" name:"TotalCapacity"`

	// 剩余可用cos的容量大小，单位GB
	TotalFreeCapacity *float64 `json:"TotalFreeCapacity,omitnil,omitempty" name:"TotalFreeCapacity"`

	// 已用cos的容量大小，单位GB
	TotalUsedCapacity *float64 `json:"TotalUsedCapacity,omitnil,omitempty" name:"TotalUsedCapacity"`
}

type CosInfo struct {
	// COS存储大小，单位TB
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// COS存储类型，默认为cos
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type CreateDedicatedClusterImageCacheRequestParams struct {
	// 集群ID
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 镜像ID
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`
}

type CreateDedicatedClusterImageCacheRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 镜像ID
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`
}

func (r *CreateDedicatedClusterImageCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDedicatedClusterImageCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	delete(f, "ImageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDedicatedClusterImageCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDedicatedClusterImageCacheResponseParams struct {
	// 任务id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDedicatedClusterImageCacheResponse struct {
	*tchttp.BaseResponse
	Response *CreateDedicatedClusterImageCacheResponseParams `json:"Response"`
}

func (r *CreateDedicatedClusterImageCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDedicatedClusterImageCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDedicatedClusterOrderRequestParams struct {
	// 专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// order关联的专用集群类型数组
	DedicatedClusterTypes []*DedicatedClusterTypeInfo `json:"DedicatedClusterTypes,omitnil,omitempty" name:"DedicatedClusterTypes"`

	// order关联的cos存储信息
	CosInfo *CosInfo `json:"CosInfo,omitnil,omitempty" name:"CosInfo"`

	// order关联的cbs存储信息
	CbsInfo *CbsInfo `json:"CbsInfo,omitnil,omitempty" name:"CbsInfo"`

	// 购买来源，默认为cloudApi
	PurchaseSource *string `json:"PurchaseSource,omitnil,omitempty" name:"PurchaseSource"`

	// 当调用API接口提交订单时，需要提交DedicatedClusterOrderId，此处DedicatedClusterOrderId是之前创建的订单，可通过DescribeDedicatedClusterOrders接口查询，这里传入DedicatedClusterOrderId用于调整订单和支付。
	DedicatedClusterOrderId *string `json:"DedicatedClusterOrderId,omitnil,omitempty" name:"DedicatedClusterOrderId"`
}

type CreateDedicatedClusterOrderRequest struct {
	*tchttp.BaseRequest
	
	// 专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// order关联的专用集群类型数组
	DedicatedClusterTypes []*DedicatedClusterTypeInfo `json:"DedicatedClusterTypes,omitnil,omitempty" name:"DedicatedClusterTypes"`

	// order关联的cos存储信息
	CosInfo *CosInfo `json:"CosInfo,omitnil,omitempty" name:"CosInfo"`

	// order关联的cbs存储信息
	CbsInfo *CbsInfo `json:"CbsInfo,omitnil,omitempty" name:"CbsInfo"`

	// 购买来源，默认为cloudApi
	PurchaseSource *string `json:"PurchaseSource,omitnil,omitempty" name:"PurchaseSource"`

	// 当调用API接口提交订单时，需要提交DedicatedClusterOrderId，此处DedicatedClusterOrderId是之前创建的订单，可通过DescribeDedicatedClusterOrders接口查询，这里传入DedicatedClusterOrderId用于调整订单和支付。
	DedicatedClusterOrderId *string `json:"DedicatedClusterOrderId,omitnil,omitempty" name:"DedicatedClusterOrderId"`
}

func (r *CreateDedicatedClusterOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDedicatedClusterOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	delete(f, "DedicatedClusterTypes")
	delete(f, "CosInfo")
	delete(f, "CbsInfo")
	delete(f, "PurchaseSource")
	delete(f, "DedicatedClusterOrderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDedicatedClusterOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDedicatedClusterOrderResponseParams struct {
	// 专用集群订单id
	DedicatedClusterOrderId *string `json:"DedicatedClusterOrderId,omitnil,omitempty" name:"DedicatedClusterOrderId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDedicatedClusterOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateDedicatedClusterOrderResponseParams `json:"Response"`
}

func (r *CreateDedicatedClusterOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDedicatedClusterOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDedicatedClusterRequestParams struct {
	// 专用集群所属的SiteId
	SiteId *string `json:"SiteId,omitnil,omitempty" name:"SiteId"`

	// 专用集群的名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 专用集群所属的可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 专用集群的描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateDedicatedClusterRequest struct {
	*tchttp.BaseRequest
	
	// 专用集群所属的SiteId
	SiteId *string `json:"SiteId,omitnil,omitempty" name:"SiteId"`

	// 专用集群的名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 专用集群所属的可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 专用集群的描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateDedicatedClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDedicatedClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SiteId")
	delete(f, "Name")
	delete(f, "Zone")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDedicatedClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDedicatedClusterResponseParams struct {
	// 创建的专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDedicatedClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateDedicatedClusterResponseParams `json:"Response"`
}

func (r *CreateDedicatedClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDedicatedClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSiteRequestParams struct {
	// 站点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 站点所在国家
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 站点所在省份
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// 站点所在城市
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// 站点所在地区的详细地址信息
	AddressLine *string `json:"AddressLine,omitnil,omitempty" name:"AddressLine"`

	// 站点描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 注意事项
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 您将使用光纤类型将CDC设备连接到网络。有单模和多模两种选项。取值范围："MM","SM"
	FiberType *string `json:"FiberType,omitnil,omitempty" name:"FiberType"`

	// 您将CDC连接到网络时采用的光学标准。此字段取决于上行链路速度、光纤类型和到上游设备的距离。
	OpticalStandard *string `json:"OpticalStandard,omitnil,omitempty" name:"OpticalStandard"`

	// 电源连接器类型
	PowerConnectors *string `json:"PowerConnectors,omitnil,omitempty" name:"PowerConnectors"`

	// 从机架上方还是下方供电。取值范围：["UP","DOWN"]
	PowerFeedDrop *string `json:"PowerFeedDrop,omitnil,omitempty" name:"PowerFeedDrop"`

	// 最大承重(KG)
	MaxWeight *int64 `json:"MaxWeight,omitnil,omitempty" name:"MaxWeight"`

	// 功耗(KW)
	PowerDrawKva *int64 `json:"PowerDrawKva,omitnil,omitempty" name:"PowerDrawKva"`

	// 网络到腾讯云Region区域的上行链路速度(Gbps)
	UplinkSpeedGbps *int64 `json:"UplinkSpeedGbps,omitnil,omitempty" name:"UplinkSpeedGbps"`

	// 将CDC连接到网络时，每台CDC网络设备(每个机架 2 台设备)使用的上行链路数量。
	UplinkCount *int64 `json:"UplinkCount,omitnil,omitempty" name:"UplinkCount"`

	// 是否满足下面环境条件：
	// 1、场地没有材料要求或验收标准会影响 CDC 设备配送和安装。
	// 2、确定的机架位置包含:
	// 温度范围为 41 到 104°F (5 到 40°C)。
	// 湿度范围为 10°F (-12°C)和 8% RH (相对湿度)到 70°F(21°C)和 80% RH。
	// 机架位置的气流方向为从前向后，且应具有足够的 CFM (每分钟立方英尺)。CFM 必须是 CDC 配置的 kVA 功耗值的 145.8 倍。
	ConditionRequirement *bool `json:"ConditionRequirement,omitnil,omitempty" name:"ConditionRequirement"`

	// 是否满足下面的尺寸条件：
	// 您的装货站台可以容纳一个机架箱(高 x 宽 x 深 = 94" x 54" x 48")。
	// 您可以提供从机架(高 x 宽 x 深 = 80" x 24" x 48")交货地点到机架最终安置位置的明确通道。测量深度时，应包括站台、走廊通道、门、转弯、坡道、货梯，并将其他通道限制考虑在内。
	// 在最终的 CDC安置位置，前部间隙可以为 48" 或更大，后部间隙可以为 24" 或更大。
	DimensionRequirement *bool `json:"DimensionRequirement,omitnil,omitempty" name:"DimensionRequirement"`

	// 是否提供冗余的上游设备(交换机或路由器)，以便两台  网络设备都能连接到网络设备。
	RedundantNetworking *bool `json:"RedundantNetworking,omitnil,omitempty" name:"RedundantNetworking"`

	// 站点所在地区的邮编
	PostalCode *int64 `json:"PostalCode,omitnil,omitempty" name:"PostalCode"`

	// 站点所在地区的详细地址信息（补充）
	OptionalAddressLine *string `json:"OptionalAddressLine,omitnil,omitempty" name:"OptionalAddressLine"`

	// 是否需要腾讯云团队协助完成机架支撑工作
	NeedHelp *bool `json:"NeedHelp,omitnil,omitempty" name:"NeedHelp"`

	// 是否电源冗余
	RedundantPower *bool `json:"RedundantPower,omitnil,omitempty" name:"RedundantPower"`

	// 上游断路器是否具备
	BreakerRequirement *bool `json:"BreakerRequirement,omitnil,omitempty" name:"BreakerRequirement"`
}

type CreateSiteRequest struct {
	*tchttp.BaseRequest
	
	// 站点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 站点所在国家
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 站点所在省份
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// 站点所在城市
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// 站点所在地区的详细地址信息
	AddressLine *string `json:"AddressLine,omitnil,omitempty" name:"AddressLine"`

	// 站点描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 注意事项
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 您将使用光纤类型将CDC设备连接到网络。有单模和多模两种选项。取值范围："MM","SM"
	FiberType *string `json:"FiberType,omitnil,omitempty" name:"FiberType"`

	// 您将CDC连接到网络时采用的光学标准。此字段取决于上行链路速度、光纤类型和到上游设备的距离。
	OpticalStandard *string `json:"OpticalStandard,omitnil,omitempty" name:"OpticalStandard"`

	// 电源连接器类型
	PowerConnectors *string `json:"PowerConnectors,omitnil,omitempty" name:"PowerConnectors"`

	// 从机架上方还是下方供电。取值范围：["UP","DOWN"]
	PowerFeedDrop *string `json:"PowerFeedDrop,omitnil,omitempty" name:"PowerFeedDrop"`

	// 最大承重(KG)
	MaxWeight *int64 `json:"MaxWeight,omitnil,omitempty" name:"MaxWeight"`

	// 功耗(KW)
	PowerDrawKva *int64 `json:"PowerDrawKva,omitnil,omitempty" name:"PowerDrawKva"`

	// 网络到腾讯云Region区域的上行链路速度(Gbps)
	UplinkSpeedGbps *int64 `json:"UplinkSpeedGbps,omitnil,omitempty" name:"UplinkSpeedGbps"`

	// 将CDC连接到网络时，每台CDC网络设备(每个机架 2 台设备)使用的上行链路数量。
	UplinkCount *int64 `json:"UplinkCount,omitnil,omitempty" name:"UplinkCount"`

	// 是否满足下面环境条件：
	// 1、场地没有材料要求或验收标准会影响 CDC 设备配送和安装。
	// 2、确定的机架位置包含:
	// 温度范围为 41 到 104°F (5 到 40°C)。
	// 湿度范围为 10°F (-12°C)和 8% RH (相对湿度)到 70°F(21°C)和 80% RH。
	// 机架位置的气流方向为从前向后，且应具有足够的 CFM (每分钟立方英尺)。CFM 必须是 CDC 配置的 kVA 功耗值的 145.8 倍。
	ConditionRequirement *bool `json:"ConditionRequirement,omitnil,omitempty" name:"ConditionRequirement"`

	// 是否满足下面的尺寸条件：
	// 您的装货站台可以容纳一个机架箱(高 x 宽 x 深 = 94" x 54" x 48")。
	// 您可以提供从机架(高 x 宽 x 深 = 80" x 24" x 48")交货地点到机架最终安置位置的明确通道。测量深度时，应包括站台、走廊通道、门、转弯、坡道、货梯，并将其他通道限制考虑在内。
	// 在最终的 CDC安置位置，前部间隙可以为 48" 或更大，后部间隙可以为 24" 或更大。
	DimensionRequirement *bool `json:"DimensionRequirement,omitnil,omitempty" name:"DimensionRequirement"`

	// 是否提供冗余的上游设备(交换机或路由器)，以便两台  网络设备都能连接到网络设备。
	RedundantNetworking *bool `json:"RedundantNetworking,omitnil,omitempty" name:"RedundantNetworking"`

	// 站点所在地区的邮编
	PostalCode *int64 `json:"PostalCode,omitnil,omitempty" name:"PostalCode"`

	// 站点所在地区的详细地址信息（补充）
	OptionalAddressLine *string `json:"OptionalAddressLine,omitnil,omitempty" name:"OptionalAddressLine"`

	// 是否需要腾讯云团队协助完成机架支撑工作
	NeedHelp *bool `json:"NeedHelp,omitnil,omitempty" name:"NeedHelp"`

	// 是否电源冗余
	RedundantPower *bool `json:"RedundantPower,omitnil,omitempty" name:"RedundantPower"`

	// 上游断路器是否具备
	BreakerRequirement *bool `json:"BreakerRequirement,omitnil,omitempty" name:"BreakerRequirement"`
}

func (r *CreateSiteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSiteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Country")
	delete(f, "Province")
	delete(f, "City")
	delete(f, "AddressLine")
	delete(f, "Description")
	delete(f, "Note")
	delete(f, "FiberType")
	delete(f, "OpticalStandard")
	delete(f, "PowerConnectors")
	delete(f, "PowerFeedDrop")
	delete(f, "MaxWeight")
	delete(f, "PowerDrawKva")
	delete(f, "UplinkSpeedGbps")
	delete(f, "UplinkCount")
	delete(f, "ConditionRequirement")
	delete(f, "DimensionRequirement")
	delete(f, "RedundantNetworking")
	delete(f, "PostalCode")
	delete(f, "OptionalAddressLine")
	delete(f, "NeedHelp")
	delete(f, "RedundantPower")
	delete(f, "BreakerRequirement")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSiteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSiteResponseParams struct {
	// 创建Site生成的id
	SiteId *string `json:"SiteId,omitnil,omitempty" name:"SiteId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSiteResponse struct {
	*tchttp.BaseResponse
	Response *CreateSiteResponseParams `json:"Response"`
}

func (r *CreateSiteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DedicatedCluster struct {
	// 专用集群id。如"cluster-xxxxx"。
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 专用集群所属可用区名称。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 专用集群的描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 专用集群的名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 专用集群的生命周期。如"PENDING"。
	LifecycleStatus *string `json:"LifecycleStatus,omitnil,omitempty" name:"LifecycleStatus"`

	// 专用集群的创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 专用集群所属的站点id。
	SiteId *string `json:"SiteId,omitnil,omitempty" name:"SiteId"`

	// 专用集群的运营状态
	RunningStatus *string `json:"RunningStatus,omitnil,omitempty" name:"RunningStatus"`
}

type DedicatedClusterInstanceType struct {
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 规格名称
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 网卡类型，例如：25代表25G网卡
	NetworkCard *int64 `json:"NetworkCard,omitnil,omitempty" name:"NetworkCard"`

	// 实例的CPU核数，单位：核。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例内存容量，单位：`GB`。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例机型系列。
	InstanceFamily *string `json:"InstanceFamily,omitnil,omitempty" name:"InstanceFamily"`

	// 机型名称。
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 本地存储块数量。
	StorageBlockAmount *int64 `json:"StorageBlockAmount,omitnil,omitempty" name:"StorageBlockAmount"`

	// 内网带宽，单位Gbps。
	InstanceBandwidth *float64 `json:"InstanceBandwidth,omitnil,omitempty" name:"InstanceBandwidth"`

	// 网络收发包能力，单位万PPS。
	InstancePps *int64 `json:"InstancePps,omitnil,omitempty" name:"InstancePps"`

	// 处理器型号。
	CpuType *string `json:"CpuType,omitnil,omitempty" name:"CpuType"`

	// 实例的GPU数量。
	Gpu *int64 `json:"Gpu,omitnil,omitempty" name:"Gpu"`

	// 实例的FPGA数量。
	Fpga *int64 `json:"Fpga,omitnil,omitempty" name:"Fpga"`

	// 机型描述
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 实例是否售卖。取值范围： <br><li>SELL：表示实例可购买<br></li><li>SOLD_OUT：表示实例已售罄。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type DedicatedClusterOrder struct {
	// 专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 专用集群类型id（移到下一层级，已经废弃，后续将删除）
	DedicatedClusterTypeId *string `json:"DedicatedClusterTypeId,omitnil,omitempty" name:"DedicatedClusterTypeId"`

	// 支持的存储类型列表（移到下一层级，已经废弃，后续将删除）
	SupportedStorageType []*string `json:"SupportedStorageType,omitnil,omitempty" name:"SupportedStorageType"`

	// 支持的上连交换机的链路传输速率(GiB)（移到下一层级，已经废弃，后续将删除）
	SupportedUplinkSpeed []*int64 `json:"SupportedUplinkSpeed,omitnil,omitempty" name:"SupportedUplinkSpeed"`

	// 支持的实例族列表（移到下一层级，已经废弃，后续将删除）
	SupportedInstanceFamily []*string `json:"SupportedInstanceFamily,omitnil,omitempty" name:"SupportedInstanceFamily"`

	// 地板承重要求(KG)
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 功率要求(KW)
	PowerDraw *float64 `json:"PowerDraw,omitnil,omitempty" name:"PowerDraw"`

	// 订单状态
	OrderStatus *string `json:"OrderStatus,omitnil,omitempty" name:"OrderStatus"`

	// 订单创建的时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 大订单ID
	DedicatedClusterOrderId *string `json:"DedicatedClusterOrderId,omitnil,omitempty" name:"DedicatedClusterOrderId"`

	// 订单类型，创建CREATE或扩容EXTEND
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 子订单详情列表
	DedicatedClusterOrderItems []*DedicatedClusterOrderItem `json:"DedicatedClusterOrderItems,omitnil,omitempty" name:"DedicatedClusterOrderItems"`

	// cpu值
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// mem值
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// gpu值
	Gpu *int64 `json:"Gpu,omitnil,omitempty" name:"Gpu"`

	// 0代表未支付，1代表已支付
	PayStatus *int64 `json:"PayStatus,omitnil,omitempty" name:"PayStatus"`

	// 支付方式，一次性、按月、按年
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// 购买时长的单位
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 购买时长
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 订单类型
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 验收状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckStatus *string `json:"CheckStatus,omitnil,omitempty" name:"CheckStatus"`

	// 交付预期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeliverExpectTime *string `json:"DeliverExpectTime,omitnil,omitempty" name:"DeliverExpectTime"`

	// 交付实际完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeliverFinishTime *string `json:"DeliverFinishTime,omitnil,omitempty" name:"DeliverFinishTime"`

	// 验收预期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckExpectTime *string `json:"CheckExpectTime,omitnil,omitempty" name:"CheckExpectTime"`

	// 验收实际完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckFinishTime *string `json:"CheckFinishTime,omitnil,omitempty" name:"CheckFinishTime"`

	// 订单SLA
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderSLA *string `json:"OrderSLA,omitnil,omitempty" name:"OrderSLA"`

	// 订单支付计划
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderPayPlan *string `json:"OrderPayPlan,omitnil,omitempty" name:"OrderPayPlan"`
}

type DedicatedClusterOrderItem struct {
	// 专用集群类型id
	DedicatedClusterTypeId *string `json:"DedicatedClusterTypeId,omitnil,omitempty" name:"DedicatedClusterTypeId"`

	// 支持的存储类型列表
	SupportedStorageType []*string `json:"SupportedStorageType,omitnil,omitempty" name:"SupportedStorageType"`

	// 支持的上连交换机的链路传输速率(GiB)
	SupportedUplinkSpeed []*int64 `json:"SupportedUplinkSpeed,omitnil,omitempty" name:"SupportedUplinkSpeed"`

	// 支持的实例族列表
	SupportedInstanceFamily []*string `json:"SupportedInstanceFamily,omitnil,omitempty" name:"SupportedInstanceFamily"`

	// 地板承重要求(KG)
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 功率要求(KW)
	PowerDraw *float64 `json:"PowerDraw,omitnil,omitempty" name:"PowerDraw"`

	// 订单状态
	SubOrderStatus *string `json:"SubOrderStatus,omitnil,omitempty" name:"SubOrderStatus"`

	// 订单创建的时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 子订单ID
	SubOrderId *string `json:"SubOrderId,omitnil,omitempty" name:"SubOrderId"`

	// 关联的集群规格数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 规格简单描述
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 规格详细描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// CPU数
	TotalCpu *int64 `json:"TotalCpu,omitnil,omitempty" name:"TotalCpu"`

	// 内存数
	TotalMem *int64 `json:"TotalMem,omitnil,omitempty" name:"TotalMem"`

	// GPU数
	TotalGpu *int64 `json:"TotalGpu,omitnil,omitempty" name:"TotalGpu"`

	// 规格英文名
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 规格展示
	ComputeFormat *string `json:"ComputeFormat,omitnil,omitempty" name:"ComputeFormat"`

	// 规格类型
	TypeFamily *string `json:"TypeFamily,omitnil,omitempty" name:"TypeFamily"`

	// 0未支付，1已支付
	SubOrderPayStatus *int64 `json:"SubOrderPayStatus,omitnil,omitempty" name:"SubOrderPayStatus"`
}

type DedicatedClusterType struct {
	// 配置id
	DedicatedClusterTypeId *string `json:"DedicatedClusterTypeId,omitnil,omitempty" name:"DedicatedClusterTypeId"`

	// 配置描述，对应描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 配置名称，对应计算资源类型
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 创建配置的时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 支持的存储类型列表
	SupportedStorageType []*string `json:"SupportedStorageType,omitnil,omitempty" name:"SupportedStorageType"`

	// 支持的上连交换机的链路传输速率
	SupportedUplinkGiB []*int64 `json:"SupportedUplinkGiB,omitnil,omitempty" name:"SupportedUplinkGiB"`

	// 支持的实例族列表
	SupportedInstanceFamily []*string `json:"SupportedInstanceFamily,omitnil,omitempty" name:"SupportedInstanceFamily"`

	// 地板承重要求(KG)
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 功率要求(KW)
	PowerDrawKva *float64 `json:"PowerDrawKva,omitnil,omitempty" name:"PowerDrawKva"`

	// 显示计算资源规格详情，存储等资源不显示
	ComputeFormatDesc *string `json:"ComputeFormatDesc,omitnil,omitempty" name:"ComputeFormatDesc"`
}

type DedicatedClusterTypeInfo struct {
	// 集群类型Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 集群类型个数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

// Predefined struct for user
type DeleteDedicatedClusterImageCacheRequestParams struct {
	// 集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 镜像id
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`
}

type DeleteDedicatedClusterImageCacheRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 镜像id
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`
}

func (r *DeleteDedicatedClusterImageCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDedicatedClusterImageCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	delete(f, "ImageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDedicatedClusterImageCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDedicatedClusterImageCacheResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDedicatedClusterImageCacheResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDedicatedClusterImageCacheResponseParams `json:"Response"`
}

func (r *DeleteDedicatedClusterImageCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDedicatedClusterImageCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDedicatedClustersRequestParams struct {
	// 要删除的专用集群id
	DedicatedClusterIds []*string `json:"DedicatedClusterIds,omitnil,omitempty" name:"DedicatedClusterIds"`
}

type DeleteDedicatedClustersRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的专用集群id
	DedicatedClusterIds []*string `json:"DedicatedClusterIds,omitnil,omitempty" name:"DedicatedClusterIds"`
}

func (r *DeleteDedicatedClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDedicatedClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDedicatedClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDedicatedClustersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDedicatedClustersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDedicatedClustersResponseParams `json:"Response"`
}

func (r *DeleteDedicatedClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDedicatedClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSitesRequestParams struct {
	// 要删除的站点id列表
	SiteIds []*string `json:"SiteIds,omitnil,omitempty" name:"SiteIds"`
}

type DeleteSitesRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的站点id列表
	SiteIds []*string `json:"SiteIds,omitnil,omitempty" name:"SiteIds"`
}

func (r *DeleteSitesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSitesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SiteIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSitesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSitesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSitesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSitesResponseParams `json:"Response"`
}

func (r *DeleteSitesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSitesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterCbsStatisticsRequestParams struct {
	// 查询的专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 云硬盘仓库id
	SetId *string `json:"SetId,omitnil,omitempty" name:"SetId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 时间范围精度，1分钟(ONE_MINUTE)/5分钟(FIVE_MINUTE)
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDedicatedClusterCbsStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 查询的专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 云硬盘仓库id
	SetId *string `json:"SetId,omitnil,omitempty" name:"SetId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 时间范围精度，1分钟(ONE_MINUTE)/5分钟(FIVE_MINUTE)
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDedicatedClusterCbsStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterCbsStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	delete(f, "SetId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Period")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedClusterCbsStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterCbsStatisticsResponseParams struct {
	// 云硬盘仓库信息
	SetList []*SetInfo `json:"SetList,omitnil,omitempty" name:"SetList"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDedicatedClusterCbsStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedClusterCbsStatisticsResponseParams `json:"Response"`
}

func (r *DescribeDedicatedClusterCbsStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterCbsStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterCosCapacityRequestParams struct {
	// 查询的专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

type DescribeDedicatedClusterCosCapacityRequest struct {
	*tchttp.BaseRequest
	
	// 查询的专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

func (r *DescribeDedicatedClusterCosCapacityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterCosCapacityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedClusterCosCapacityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterCosCapacityResponseParams struct {
	// 本集群内cos容量信息，单位：‘GB’
	CosCapacity *CosCapacity `json:"CosCapacity,omitnil,omitempty" name:"CosCapacity"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDedicatedClusterCosCapacityResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedClusterCosCapacityResponseParams `json:"Response"`
}

func (r *DescribeDedicatedClusterCosCapacityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterCosCapacityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterHostStatisticsRequestParams struct {
	// 查询的专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 宿主机id
	HostId *string `json:"HostId,omitnil,omitempty" name:"HostId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 时间范围精度，1分钟(ONE_MINUTE)/5分钟(FIVE_MINUTE)
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`
}

type DescribeDedicatedClusterHostStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 查询的专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 宿主机id
	HostId *string `json:"HostId,omitnil,omitempty" name:"HostId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 时间范围精度，1分钟(ONE_MINUTE)/5分钟(FIVE_MINUTE)
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *DescribeDedicatedClusterHostStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterHostStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	delete(f, "HostId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedClusterHostStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterHostStatisticsResponseParams struct {
	// 该集群内宿主机的统计信息列表
	HostStatisticSet []*HostStatistic `json:"HostStatisticSet,omitnil,omitempty" name:"HostStatisticSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDedicatedClusterHostStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedClusterHostStatisticsResponseParams `json:"Response"`
}

func (r *DescribeDedicatedClusterHostStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterHostStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterHostsRequestParams struct {
	// 集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDedicatedClusterHostsRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDedicatedClusterHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedClusterHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterHostsResponseParams struct {
	// 宿主机信息
	HostInfoSet []*HostInfo `json:"HostInfoSet,omitnil,omitempty" name:"HostInfoSet"`

	// 宿主机总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDedicatedClusterHostsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedClusterHostsResponseParams `json:"Response"`
}

func (r *DescribeDedicatedClusterHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterInstanceTypesRequestParams struct {
	// 查询的专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

type DescribeDedicatedClusterInstanceTypesRequest struct {
	*tchttp.BaseRequest
	
	// 查询的专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

func (r *DescribeDedicatedClusterInstanceTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterInstanceTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedClusterInstanceTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterInstanceTypesResponseParams struct {
	// 支持的实例规格列表
	DedicatedClusterInstanceTypeSet []*DedicatedClusterInstanceType `json:"DedicatedClusterInstanceTypeSet,omitnil,omitempty" name:"DedicatedClusterInstanceTypeSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDedicatedClusterInstanceTypesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedClusterInstanceTypesResponseParams `json:"Response"`
}

func (r *DescribeDedicatedClusterInstanceTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterInstanceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterOrdersRequestParams struct {
	// 按照专用集群id过滤
	DedicatedClusterIds []*string `json:"DedicatedClusterIds,omitnil,omitempty" name:"DedicatedClusterIds"`

	// 按照专用集群订单id过滤
	DedicatedClusterOrderIds *string `json:"DedicatedClusterOrderIds,omitnil,omitempty" name:"DedicatedClusterOrderIds"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 订单状态为过滤条件：PENDING INCONSTRUCTION DELIVERING DELIVERED EXPIRED CANCELLED  OFFLINE
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 订单类型为过滤条件：CREATE  EXTEND
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 订单类型列表
	OrderTypes []*string `json:"OrderTypes,omitnil,omitempty" name:"OrderTypes"`
}

type DescribeDedicatedClusterOrdersRequest struct {
	*tchttp.BaseRequest
	
	// 按照专用集群id过滤
	DedicatedClusterIds []*string `json:"DedicatedClusterIds,omitnil,omitempty" name:"DedicatedClusterIds"`

	// 按照专用集群订单id过滤
	DedicatedClusterOrderIds *string `json:"DedicatedClusterOrderIds,omitnil,omitempty" name:"DedicatedClusterOrderIds"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 订单状态为过滤条件：PENDING INCONSTRUCTION DELIVERING DELIVERED EXPIRED CANCELLED  OFFLINE
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 订单类型为过滤条件：CREATE  EXTEND
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 订单类型列表
	OrderTypes []*string `json:"OrderTypes,omitnil,omitempty" name:"OrderTypes"`
}

func (r *DescribeDedicatedClusterOrdersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterOrdersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterIds")
	delete(f, "DedicatedClusterOrderIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Status")
	delete(f, "ActionType")
	delete(f, "OrderTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedClusterOrdersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterOrdersResponseParams struct {
	// 专用集群订单列表
	DedicatedClusterOrderSet []*DedicatedClusterOrder `json:"DedicatedClusterOrderSet,omitnil,omitempty" name:"DedicatedClusterOrderSet"`

	// 符合条件的专用集群订单总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDedicatedClusterOrdersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedClusterOrdersResponseParams `json:"Response"`
}

func (r *DescribeDedicatedClusterOrdersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterOrdersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterOverviewRequestParams struct {
	// 集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

type DescribeDedicatedClusterOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

func (r *DescribeDedicatedClusterOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedClusterOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterOverviewResponseParams struct {
	// 云服务器数量
	CvmCount *uint64 `json:"CvmCount,omitnil,omitempty" name:"CvmCount"`

	// 宿主机数量
	HostCount *uint64 `json:"HostCount,omitnil,omitempty" name:"HostCount"`

	// vpn通道状态
	VpnConnectionState *string `json:"VpnConnectionState,omitnil,omitempty" name:"VpnConnectionState"`

	// vpn网关监控数据
	VpngwBandwidthData *VpngwBandwidthData `json:"VpngwBandwidthData,omitnil,omitempty" name:"VpngwBandwidthData"`

	// 本地网关信息
	LocalNetInfo *LocalNetInfo `json:"LocalNetInfo,omitnil,omitempty" name:"LocalNetInfo"`

	// vpn网关通道监控数据
	VpnConnectionBandwidthData []*VpngwBandwidthData `json:"VpnConnectionBandwidthData,omitnil,omitempty" name:"VpnConnectionBandwidthData"`

	// 宿主机资源概览信息
	HostDetailInfo []*HostDetailInfo `json:"HostDetailInfo,omitnil,omitempty" name:"HostDetailInfo"`

	// 热备宿主机数量
	HostStandbyCount *uint64 `json:"HostStandbyCount,omitnil,omitempty" name:"HostStandbyCount"`

	// 普通宿主机数量
	HostNormalCount *uint64 `json:"HostNormalCount,omitnil,omitempty" name:"HostNormalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDedicatedClusterOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedClusterOverviewResponseParams `json:"Response"`
}

func (r *DescribeDedicatedClusterOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterTypesRequestParams struct {
	// 模糊匹配专用集群配置名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 待查询的专用集群配置id列表
	DedicatedClusterTypeIds []*string `json:"DedicatedClusterTypeIds,omitnil,omitempty" name:"DedicatedClusterTypeIds"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否只查询计算规格类型
	IsCompute *bool `json:"IsCompute,omitnil,omitempty" name:"IsCompute"`
}

type DescribeDedicatedClusterTypesRequest struct {
	*tchttp.BaseRequest
	
	// 模糊匹配专用集群配置名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 待查询的专用集群配置id列表
	DedicatedClusterTypeIds []*string `json:"DedicatedClusterTypeIds,omitnil,omitempty" name:"DedicatedClusterTypeIds"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否只查询计算规格类型
	IsCompute *bool `json:"IsCompute,omitnil,omitempty" name:"IsCompute"`
}

func (r *DescribeDedicatedClusterTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "DedicatedClusterTypeIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "IsCompute")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedClusterTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClusterTypesResponseParams struct {
	// 专用集群配置列表
	DedicatedClusterTypeSet []*DedicatedClusterType `json:"DedicatedClusterTypeSet,omitnil,omitempty" name:"DedicatedClusterTypeSet"`

	// 符合条件的个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDedicatedClusterTypesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedClusterTypesResponseParams `json:"Response"`
}

func (r *DescribeDedicatedClusterTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClusterTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClustersRequestParams struct {
	// 按照一个或者多个实例ID查询。实例ID形如：`cluster-xxxxxxxx`
	DedicatedClusterIds []*string `json:"DedicatedClusterIds,omitnil,omitempty" name:"DedicatedClusterIds"`

	// 按照可用区名称过滤
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 按照站点id过滤
	SiteIds []*string `json:"SiteIds,omitnil,omitempty" name:"SiteIds"`

	// 按照专用集群生命周期过滤
	LifecycleStatuses []*string `json:"LifecycleStatuses,omitnil,omitempty" name:"LifecycleStatuses"`

	// 模糊匹配专用集群名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDedicatedClustersRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个实例ID查询。实例ID形如：`cluster-xxxxxxxx`
	DedicatedClusterIds []*string `json:"DedicatedClusterIds,omitnil,omitempty" name:"DedicatedClusterIds"`

	// 按照可用区名称过滤
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 按照站点id过滤
	SiteIds []*string `json:"SiteIds,omitnil,omitempty" name:"SiteIds"`

	// 按照专用集群生命周期过滤
	LifecycleStatuses []*string `json:"LifecycleStatuses,omitnil,omitempty" name:"LifecycleStatuses"`

	// 模糊匹配专用集群名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDedicatedClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterIds")
	delete(f, "Zones")
	delete(f, "SiteIds")
	delete(f, "LifecycleStatuses")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClustersResponseParams struct {
	// 符合查询条件的专用集群列表
	DedicatedClusterSet []*DedicatedCluster `json:"DedicatedClusterSet,omitnil,omitempty" name:"DedicatedClusterSet"`

	// 符合条件的专用集群数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDedicatedClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedClustersResponseParams `json:"Response"`
}

func (r *DescribeDedicatedClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedSupportedZonesRequestParams struct {
	// 传入region列表
	Regions []*int64 `json:"Regions,omitnil,omitempty" name:"Regions"`
}

type DescribeDedicatedSupportedZonesRequest struct {
	*tchttp.BaseRequest
	
	// 传入region列表
	Regions []*int64 `json:"Regions,omitnil,omitempty" name:"Regions"`
}

func (r *DescribeDedicatedSupportedZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedSupportedZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Regions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedSupportedZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedSupportedZonesResponseParams struct {
	// 支持的可用区列表
	ZoneSet []*RegionZoneInfo `json:"ZoneSet,omitnil,omitempty" name:"ZoneSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDedicatedSupportedZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedSupportedZonesResponseParams `json:"Response"`
}

func (r *DescribeDedicatedSupportedZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedSupportedZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSitesDetailRequestParams struct {
	// 按照站点id过滤
	SiteIds []*string `json:"SiteIds,omitnil,omitempty" name:"SiteIds"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 按照站定名称模糊匹配
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeSitesDetailRequest struct {
	*tchttp.BaseRequest
	
	// 按照站点id过滤
	SiteIds []*string `json:"SiteIds,omitnil,omitempty" name:"SiteIds"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 按照站定名称模糊匹配
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DescribeSitesDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSitesDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SiteIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSitesDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSitesDetailResponseParams struct {
	// 站点详情
	SiteDetailSet []*SiteDetail `json:"SiteDetailSet,omitnil,omitempty" name:"SiteDetailSet"`

	// 符合条件的站点总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSitesDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSitesDetailResponseParams `json:"Response"`
}

func (r *DescribeSitesDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSitesDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSitesRequestParams struct {
	// 按照站点id过滤
	SiteIds []*string `json:"SiteIds,omitnil,omitempty" name:"SiteIds"`

	// 模糊匹配站点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSitesRequest struct {
	*tchttp.BaseRequest
	
	// 按照站点id过滤
	SiteIds []*string `json:"SiteIds,omitnil,omitempty" name:"SiteIds"`

	// 模糊匹配站点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeSitesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSitesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SiteIds")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSitesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSitesResponseParams struct {
	// 符合查询条件的站点列表
	SiteSet []*Site `json:"SiteSet,omitnil,omitempty" name:"SiteSet"`

	// 符合条件的站点数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSitesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSitesResponseParams `json:"Response"`
}

func (r *DescribeSitesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSitesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailData struct {
	// 时间戳
	Timestamps []*float64 `json:"Timestamps,omitnil,omitempty" name:"Timestamps"`

	// 对应的具体值
	Values []*float64 `json:"Values,omitnil,omitempty" name:"Values"`
}

type HostDetailInfo struct {
	// 类型族
	HostTypeFamily *string `json:"HostTypeFamily,omitnil,omitempty" name:"HostTypeFamily"`

	// 总CPU
	CpuTotal *float64 `json:"CpuTotal,omitnil,omitempty" name:"CpuTotal"`

	// 可用CPU
	CpuAvailable *float64 `json:"CpuAvailable,omitnil,omitempty" name:"CpuAvailable"`

	// 总内存
	MemTotal *float64 `json:"MemTotal,omitnil,omitempty" name:"MemTotal"`

	// 可用内存
	MemAvailable *float64 `json:"MemAvailable,omitnil,omitempty" name:"MemAvailable"`
}

type HostInfo struct {
	// 宿主机IP（废弃）
	HostIp *string `json:"HostIp,omitnil,omitempty" name:"HostIp"`

	// 云服务类型
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 宿主机运行状态
	HostStatus *string `json:"HostStatus,omitnil,omitempty" name:"HostStatus"`

	// 宿主机类型
	HostType *string `json:"HostType,omitnil,omitempty" name:"HostType"`

	// cpu可用数
	CpuAvailable *uint64 `json:"CpuAvailable,omitnil,omitempty" name:"CpuAvailable"`

	// cpu总数
	CpuTotal *uint64 `json:"CpuTotal,omitnil,omitempty" name:"CpuTotal"`

	// 内存可用数
	MemAvailable *uint64 `json:"MemAvailable,omitnil,omitempty" name:"MemAvailable"`

	// 内存总数
	MemTotal *uint64 `json:"MemTotal,omitnil,omitempty" name:"MemTotal"`

	// 运行时间
	RunTime *string `json:"RunTime,omitnil,omitempty" name:"RunTime"`

	// 到期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 宿主机id
	HostId *string `json:"HostId,omitnil,omitempty" name:"HostId"`
}

type HostStatistic struct {
	// 宿主机规格
	HostType *string `json:"HostType,omitnil,omitempty" name:"HostType"`

	// 宿主机机型系列
	HostFamily *string `json:"HostFamily,omitnil,omitempty" name:"HostFamily"`

	// 宿主机的CPU核数，单位：核
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 宿主机内存大小，单位：GB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 该规格宿主机的数量
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 平均cpu负载百分比
	CpuAverage *float64 `json:"CpuAverage,omitnil,omitempty" name:"CpuAverage"`

	// 平均内存使用率百分比
	MemAverage *float64 `json:"MemAverage,omitnil,omitempty" name:"MemAverage"`

	// 平均网络流量
	NetAverage *float64 `json:"NetAverage,omitnil,omitempty" name:"NetAverage"`

	// cpu详细监控数据
	CpuDetailData *DetailData `json:"CpuDetailData,omitnil,omitempty" name:"CpuDetailData"`

	// 内存详细数据
	MemDetailData *DetailData `json:"MemDetailData,omitnil,omitempty" name:"MemDetailData"`

	// 网络速率详细数据
	NetRateDetailData *DetailData `json:"NetRateDetailData,omitnil,omitempty" name:"NetRateDetailData"`

	// 网速包详细数据
	NetPacketDetailData *DetailData `json:"NetPacketDetailData,omitnil,omitempty" name:"NetPacketDetailData"`
}

type InBandwidth struct {
	// 时间戳
	Timestamps []*float64 `json:"Timestamps,omitnil,omitempty" name:"Timestamps"`

	// 时间对应的值
	Values []*float64 `json:"Values,omitnil,omitempty" name:"Values"`
}

type LocalNetInfo struct {
	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 网络id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 路由信息
	BGPRoute *string `json:"BGPRoute,omitnil,omitempty" name:"BGPRoute"`

	// 本地IP
	LocalIp *string `json:"LocalIp,omitnil,omitempty" name:"LocalIp"`
}

// Predefined struct for user
type ModifyDedicatedClusterInfoRequestParams struct {
	// 本地专用集群ID
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 集群的新名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 集群的新可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 集群的新描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 集群所在站点
	SiteId *string `json:"SiteId,omitnil,omitempty" name:"SiteId"`
}

type ModifyDedicatedClusterInfoRequest struct {
	*tchttp.BaseRequest
	
	// 本地专用集群ID
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 集群的新名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 集群的新可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 集群的新描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 集群所在站点
	SiteId *string `json:"SiteId,omitnil,omitempty" name:"SiteId"`
}

func (r *ModifyDedicatedClusterInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDedicatedClusterInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	delete(f, "Name")
	delete(f, "Zone")
	delete(f, "Description")
	delete(f, "SiteId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDedicatedClusterInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDedicatedClusterInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDedicatedClusterInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDedicatedClusterInfoResponseParams `json:"Response"`
}

func (r *ModifyDedicatedClusterInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDedicatedClusterInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOrderStatusRequestParams struct {
	// 要更新成的状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 大订单ID，可以在本地专用集群的基础信息页获取大订单ID
	DedicatedClusterOrderId *string `json:"DedicatedClusterOrderId,omitnil,omitempty" name:"DedicatedClusterOrderId"`

	// 小订单ID，进入大订单的详情页，可以看到小订单ID
	SubOrderIds []*string `json:"SubOrderIds,omitnil,omitempty" name:"SubOrderIds"`
}

type ModifyOrderStatusRequest struct {
	*tchttp.BaseRequest
	
	// 要更新成的状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 大订单ID，可以在本地专用集群的基础信息页获取大订单ID
	DedicatedClusterOrderId *string `json:"DedicatedClusterOrderId,omitnil,omitempty" name:"DedicatedClusterOrderId"`

	// 小订单ID，进入大订单的详情页，可以看到小订单ID
	SubOrderIds []*string `json:"SubOrderIds,omitnil,omitempty" name:"SubOrderIds"`
}

func (r *ModifyOrderStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOrderStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "DedicatedClusterOrderId")
	delete(f, "SubOrderIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOrderStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOrderStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyOrderStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyOrderStatusResponseParams `json:"Response"`
}

func (r *ModifyOrderStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOrderStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySiteDeviceInfoRequestParams struct {
	// 机房ID
	SiteId *string `json:"SiteId,omitnil,omitempty" name:"SiteId"`

	// 您将使用光纤类型将CDC设备连接到网络。有单模和多模两种选项。
	FiberType *string `json:"FiberType,omitnil,omitempty" name:"FiberType"`

	// 您将CDC连接到网络时采用的光学标准。此字段取决于上行链路速度、光纤类型和到上游设备的距离。
	OpticalStandard *string `json:"OpticalStandard,omitnil,omitempty" name:"OpticalStandard"`

	// 电源连接器类型
	PowerConnectors *string `json:"PowerConnectors,omitnil,omitempty" name:"PowerConnectors"`

	// 从机架上方还是下方供电。取值范围：["UP","DOWN"]
	PowerFeedDrop *string `json:"PowerFeedDrop,omitnil,omitempty" name:"PowerFeedDrop"`

	// 最大承重(KG)
	MaxWeight *int64 `json:"MaxWeight,omitnil,omitempty" name:"MaxWeight"`

	// 功耗(KW)
	PowerDrawKva *int64 `json:"PowerDrawKva,omitnil,omitempty" name:"PowerDrawKva"`

	// 网络到腾讯云Region区域的上行链路速度(Gbps)
	UplinkSpeedGbps *int64 `json:"UplinkSpeedGbps,omitnil,omitempty" name:"UplinkSpeedGbps"`

	// 将CDC连接到网络时，每台CDC网络设备(每个机架 2 台设备)使用的上行链路数量。
	UplinkCount *int64 `json:"UplinkCount,omitnil,omitempty" name:"UplinkCount"`

	// 是否满足下面环境条件：
	// 1、场地没有材料要求或验收标准会影响 CDC 设备配送和安装。
	// 2、确定的机架位置包含:
	// 温度范围为 41 到 104°F (5 到 40°C)。
	// 湿度范围为 10°F (-12°C)和 8% RH (相对湿度)到 70°F(21°C)和 80% RH。
	// 机架位置的气流方向为从前向后，且应具有足够的 CFM (每分钟立方英尺)，散热功率须达到CDC运行功率值的 145.8 倍以上。
	ConditionRequirement *bool `json:"ConditionRequirement,omitnil,omitempty" name:"ConditionRequirement"`

	// 是否满足下面的尺寸条件：
	// 您的装货站台可以容纳一个机架箱(高 x 宽 x 深 = 94" x 54" x 48")。
	// 您可以提供从机架(高 x 宽 x 深 = 80" x 24" x 48")交货地点到机架最终安置位置的明确通道。测量深度时，应包括站台、走廊通道、门、转弯、坡道、货梯，并将其他通道限制考虑在内。
	// 在最终的 CDC安置位置，前部间隙可以为 48" 或更大，后部间隙可以为 24" 或更大。
	DimensionRequirement *bool `json:"DimensionRequirement,omitnil,omitempty" name:"DimensionRequirement"`

	// 是否提供冗余的上游设备(交换机或路由器)，以便实现网络出口的高可用。
	RedundantNetworking *bool `json:"RedundantNetworking,omitnil,omitempty" name:"RedundantNetworking"`

	// 是否需要腾讯云团队协助完成机架支撑工作
	NeedHelp *bool `json:"NeedHelp,omitnil,omitempty" name:"NeedHelp"`

	// 是否电源冗余
	RedundantPower *bool `json:"RedundantPower,omitnil,omitempty" name:"RedundantPower"`

	// 上游断路器是否具备
	BreakerRequirement *bool `json:"BreakerRequirement,omitnil,omitempty" name:"BreakerRequirement"`
}

type ModifySiteDeviceInfoRequest struct {
	*tchttp.BaseRequest
	
	// 机房ID
	SiteId *string `json:"SiteId,omitnil,omitempty" name:"SiteId"`

	// 您将使用光纤类型将CDC设备连接到网络。有单模和多模两种选项。
	FiberType *string `json:"FiberType,omitnil,omitempty" name:"FiberType"`

	// 您将CDC连接到网络时采用的光学标准。此字段取决于上行链路速度、光纤类型和到上游设备的距离。
	OpticalStandard *string `json:"OpticalStandard,omitnil,omitempty" name:"OpticalStandard"`

	// 电源连接器类型
	PowerConnectors *string `json:"PowerConnectors,omitnil,omitempty" name:"PowerConnectors"`

	// 从机架上方还是下方供电。取值范围：["UP","DOWN"]
	PowerFeedDrop *string `json:"PowerFeedDrop,omitnil,omitempty" name:"PowerFeedDrop"`

	// 最大承重(KG)
	MaxWeight *int64 `json:"MaxWeight,omitnil,omitempty" name:"MaxWeight"`

	// 功耗(KW)
	PowerDrawKva *int64 `json:"PowerDrawKva,omitnil,omitempty" name:"PowerDrawKva"`

	// 网络到腾讯云Region区域的上行链路速度(Gbps)
	UplinkSpeedGbps *int64 `json:"UplinkSpeedGbps,omitnil,omitempty" name:"UplinkSpeedGbps"`

	// 将CDC连接到网络时，每台CDC网络设备(每个机架 2 台设备)使用的上行链路数量。
	UplinkCount *int64 `json:"UplinkCount,omitnil,omitempty" name:"UplinkCount"`

	// 是否满足下面环境条件：
	// 1、场地没有材料要求或验收标准会影响 CDC 设备配送和安装。
	// 2、确定的机架位置包含:
	// 温度范围为 41 到 104°F (5 到 40°C)。
	// 湿度范围为 10°F (-12°C)和 8% RH (相对湿度)到 70°F(21°C)和 80% RH。
	// 机架位置的气流方向为从前向后，且应具有足够的 CFM (每分钟立方英尺)，散热功率须达到CDC运行功率值的 145.8 倍以上。
	ConditionRequirement *bool `json:"ConditionRequirement,omitnil,omitempty" name:"ConditionRequirement"`

	// 是否满足下面的尺寸条件：
	// 您的装货站台可以容纳一个机架箱(高 x 宽 x 深 = 94" x 54" x 48")。
	// 您可以提供从机架(高 x 宽 x 深 = 80" x 24" x 48")交货地点到机架最终安置位置的明确通道。测量深度时，应包括站台、走廊通道、门、转弯、坡道、货梯，并将其他通道限制考虑在内。
	// 在最终的 CDC安置位置，前部间隙可以为 48" 或更大，后部间隙可以为 24" 或更大。
	DimensionRequirement *bool `json:"DimensionRequirement,omitnil,omitempty" name:"DimensionRequirement"`

	// 是否提供冗余的上游设备(交换机或路由器)，以便实现网络出口的高可用。
	RedundantNetworking *bool `json:"RedundantNetworking,omitnil,omitempty" name:"RedundantNetworking"`

	// 是否需要腾讯云团队协助完成机架支撑工作
	NeedHelp *bool `json:"NeedHelp,omitnil,omitempty" name:"NeedHelp"`

	// 是否电源冗余
	RedundantPower *bool `json:"RedundantPower,omitnil,omitempty" name:"RedundantPower"`

	// 上游断路器是否具备
	BreakerRequirement *bool `json:"BreakerRequirement,omitnil,omitempty" name:"BreakerRequirement"`
}

func (r *ModifySiteDeviceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySiteDeviceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SiteId")
	delete(f, "FiberType")
	delete(f, "OpticalStandard")
	delete(f, "PowerConnectors")
	delete(f, "PowerFeedDrop")
	delete(f, "MaxWeight")
	delete(f, "PowerDrawKva")
	delete(f, "UplinkSpeedGbps")
	delete(f, "UplinkCount")
	delete(f, "ConditionRequirement")
	delete(f, "DimensionRequirement")
	delete(f, "RedundantNetworking")
	delete(f, "NeedHelp")
	delete(f, "RedundantPower")
	delete(f, "BreakerRequirement")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySiteDeviceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySiteDeviceInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySiteDeviceInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifySiteDeviceInfoResponseParams `json:"Response"`
}

func (r *ModifySiteDeviceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySiteDeviceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySiteInfoRequestParams struct {
	// 机房ID
	SiteId *string `json:"SiteId,omitnil,omitempty" name:"SiteId"`

	// 站点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 站点描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 注意事项
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 站点所在国家
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 站点所在省份
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// 站点所在城市
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// 站点所在地区的邮编
	PostalCode *string `json:"PostalCode,omitnil,omitempty" name:"PostalCode"`

	// 站点所在地区的详细地址信息
	AddressLine *string `json:"AddressLine,omitnil,omitempty" name:"AddressLine"`
}

type ModifySiteInfoRequest struct {
	*tchttp.BaseRequest
	
	// 机房ID
	SiteId *string `json:"SiteId,omitnil,omitempty" name:"SiteId"`

	// 站点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 站点描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 注意事项
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 站点所在国家
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 站点所在省份
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// 站点所在城市
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// 站点所在地区的邮编
	PostalCode *string `json:"PostalCode,omitnil,omitempty" name:"PostalCode"`

	// 站点所在地区的详细地址信息
	AddressLine *string `json:"AddressLine,omitnil,omitempty" name:"AddressLine"`
}

func (r *ModifySiteInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySiteInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SiteId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Note")
	delete(f, "Country")
	delete(f, "Province")
	delete(f, "City")
	delete(f, "PostalCode")
	delete(f, "AddressLine")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySiteInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySiteInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySiteInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifySiteInfoResponseParams `json:"Response"`
}

func (r *ModifySiteInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySiteInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OutBandwidth struct {
	// 时间戳
	Timestamps []*float64 `json:"Timestamps,omitnil,omitempty" name:"Timestamps"`

	// 对应时间的值
	Values []*float64 `json:"Values,omitnil,omitempty" name:"Values"`
}

type RegionZoneInfo struct {
	// Region id
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// ZoneInfo数组
	Zones []*ZoneInfo `json:"Zones,omitnil,omitempty" name:"Zones"`
}

type SetInfo struct {
	// 云硬盘仓库id
	SetId *string `json:"SetId,omitnil,omitempty" name:"SetId"`

	// 云硬盘仓库名称
	SetName *string `json:"SetName,omitnil,omitempty" name:"SetName"`

	// 云硬盘仓库类型
	SetType *string `json:"SetType,omitnil,omitempty" name:"SetType"`

	// 云硬盘仓库容量
	SetSize *float64 `json:"SetSize,omitnil,omitempty" name:"SetSize"`

	// 云硬盘仓库状态
	SetStatus *string `json:"SetStatus,omitnil,omitempty" name:"SetStatus"`

	// 云硬盘仓库创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 读流量
	ReadTraffic *DetailData `json:"ReadTraffic,omitnil,omitempty" name:"ReadTraffic"`

	// 写流量
	WriteTraffic *DetailData `json:"WriteTraffic,omitnil,omitempty" name:"WriteTraffic"`

	// 读IO
	ReadIO *DetailData `json:"ReadIO,omitnil,omitempty" name:"ReadIO"`

	// 写IO
	WriteIO *DetailData `json:"WriteIO,omitnil,omitempty" name:"WriteIO"`

	// 平均等待时间
	Await *DetailData `json:"Await,omitnil,omitempty" name:"Await"`

	// 利用率
	Util *DetailData `json:"Util,omitnil,omitempty" name:"Util"`
}

type Site struct {
	// 站点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 站点id
	SiteId *string `json:"SiteId,omitnil,omitempty" name:"SiteId"`

	// 站点描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 站点创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type SiteDetail struct {
	// 站点id
	SiteId *string `json:"SiteId,omitnil,omitempty" name:"SiteId"`

	// 站点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 站点描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 站点创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 光纤类型
	FiberType *string `json:"FiberType,omitnil,omitempty" name:"FiberType"`

	// 网络到腾讯云Region区域的上行链路速度
	UplinkSpeedGbps *int64 `json:"UplinkSpeedGbps,omitnil,omitempty" name:"UplinkSpeedGbps"`

	// 将CDC连接到网络时，每台CDC网络设备(每个机架 2 台设备)使用的上行链路数量。
	UplinkCount *int64 `json:"UplinkCount,omitnil,omitempty" name:"UplinkCount"`

	// 将CDC连接到网络时采用的光学标准
	OpticalStandard *string `json:"OpticalStandard,omitnil,omitempty" name:"OpticalStandard"`

	// 是否提供冗余的上游设备(交换机或路由器)，以便两台  网络设备都能连接到网络设备。
	RedundantNetworking *bool `json:"RedundantNetworking,omitnil,omitempty" name:"RedundantNetworking"`

	// 电源连接器类型
	PowerConnectors *string `json:"PowerConnectors,omitnil,omitempty" name:"PowerConnectors"`

	// 从机架上方还是下方供电。
	PowerFeedDrop *string `json:"PowerFeedDrop,omitnil,omitempty" name:"PowerFeedDrop"`

	// 功耗(KW)
	PowerDrawKva *float64 `json:"PowerDrawKva,omitnil,omitempty" name:"PowerDrawKva"`

	// 是否满足下面环境条件：
	// 1、场地没有材料要求或验收标准会影响 CDC 设备配送和安装。
	// 2、确定的机架位置包含:
	// 温度范围为 41 到 104°F (5 到 40°C)。
	// 湿度范围为 10°F (-12°C)和 8% RH (相对湿度)到 70°F(21°C)和 80% RH。
	// 机架位置的气流方向为从前向后，且应具有足够的 CFM (每分钟立方英尺)。CFM 必须是 CDC 配置的 kVA 功耗值的 145.8 倍。
	ConditionRequirement *bool `json:"ConditionRequirement,omitnil,omitempty" name:"ConditionRequirement"`

	// 是否满足下面的尺寸条件：
	// 您的装货站台可以容纳一个机架箱(高 x 宽 x 深 = 94" x 54" x 48")。
	// 您可以提供从机架(高 x 宽 x 深 = 80" x 24" x 48")交货地点到机架最终安置位置的明确通道。测量深度时，应包括站台、走廊通道、门、转弯、坡道、货梯，并将其他通道限制考虑在内。
	// 在最终的 CDC安置位置，前部间隙可以为 48" 或更大，后部间隙可以为 24" 或更大。
	DimensionRequirement *bool `json:"DimensionRequirement,omitnil,omitempty" name:"DimensionRequirement"`

	// 最大承重(KG)
	MaxWeight *int64 `json:"MaxWeight,omitnil,omitempty" name:"MaxWeight"`

	// 站点地址
	AddressLine *string `json:"AddressLine,omitnil,omitempty" name:"AddressLine"`

	// 站点所在地区的详细地址信息（补充）
	OptionalAddressLine *string `json:"OptionalAddressLine,omitnil,omitempty" name:"OptionalAddressLine"`

	// 是否需要腾讯云团队协助完成机架支撑工作
	NeedHelp *bool `json:"NeedHelp,omitnil,omitempty" name:"NeedHelp"`

	// 上游断路器是否具备
	BreakerRequirement *bool `json:"BreakerRequirement,omitnil,omitempty" name:"BreakerRequirement"`

	// 是否电源冗余
	RedundantPower *bool `json:"RedundantPower,omitnil,omitempty" name:"RedundantPower"`

	// 站点所在国家
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 站点所在省份
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// 站点所在城市
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// 站点所在地区的邮编
	PostalCode *int64 `json:"PostalCode,omitnil,omitempty" name:"PostalCode"`
}

type VpngwBandwidthData struct {
	// 出带宽流量
	OutBandwidth *OutBandwidth `json:"OutBandwidth,omitnil,omitempty" name:"OutBandwidth"`

	// 入带宽流量
	InBandwidth *InBandwidth `json:"InBandwidth,omitnil,omitempty" name:"InBandwidth"`
}

type ZoneInfo struct {
	// 可用区名称
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 可用区描述
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 可用区ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 可用区状态，包含AVAILABLE和UNAVAILABLE。AVAILABLE代表可用，UNAVAILABLE代表不可用。
	ZoneState *string `json:"ZoneState,omitnil,omitempty" name:"ZoneState"`
}