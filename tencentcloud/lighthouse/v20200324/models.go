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

package v20200324

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ApplyInstanceSnapshotRequest struct {
	*tchttp.BaseRequest

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 快照 ID。
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

func (r *ApplyInstanceSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyInstanceSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SnapshotId")
	if len(f) > 0 {
		return errors.New("ApplyInstanceSnapshotRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ApplyInstanceSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyInstanceSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyInstanceSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Blueprint struct {

	// 镜像 ID  ，是 Blueprint 的唯一标识。
	BlueprintId *string `json:"BlueprintId,omitempty" name:"BlueprintId"`

	// 镜像对外展示标题。
	DisplayTitle *string `json:"DisplayTitle,omitempty" name:"DisplayTitle"`

	// 镜像对外展示版本。
	DisplayVersion *string `json:"DisplayVersion,omitempty" name:"DisplayVersion"`

	// 镜像描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 操作系统名称。
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 操作系统平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 操作系统平台类型，如 LINUX_UNIX、WINDOWS。
	PlatformType *string `json:"PlatformType,omitempty" name:"PlatformType"`

	// 镜像类型，如 APP_OS、PURE_OS、PRIVATE。
	BlueprintType *string `json:"BlueprintType,omitempty" name:"BlueprintType"`

	// 镜像图片 URL。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 镜像所需系统盘大小。
	RequiredSystemDiskSize *int64 `json:"RequiredSystemDiskSize,omitempty" name:"RequiredSystemDiskSize"`

	// 镜像状态。
	BlueprintState *string `json:"BlueprintState,omitempty" name:"BlueprintState"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 镜像名称。
	BlueprintName *string `json:"BlueprintName,omitempty" name:"BlueprintName"`

	// 镜像是否支持自动化助手。
	SupportAutomationTools *bool `json:"SupportAutomationTools,omitempty" name:"SupportAutomationTools"`

	// 镜像所需内存大小, 单位: GB
	RequiredMemorySize *int64 `json:"RequiredMemorySize,omitempty" name:"RequiredMemorySize"`
}

type Bundle struct {

	// 套餐 ID。
	BundleId *string `json:"BundleId,omitempty" name:"BundleId"`

	// 内存大小，单位 GB。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 系统盘类型。
	// 取值范围： 
	// <li> LOCAL_BASIC：本地硬盘</li><li> LOCAL_SSD：本地 SSD 硬盘</li><li> CLOUD_BASIC：普通云硬盘</li><li> CLOUD_SSD：SSD 云硬盘</li><li> CLOUD_PREMIUM：高性能云硬盘</li>
	SystemDiskType *string `json:"SystemDiskType,omitempty" name:"SystemDiskType"`

	// 系统盘大小。
	SystemDiskSize *int64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`

	// 每月网络流量，单位 Gb。
	MonthlyTraffic *int64 `json:"MonthlyTraffic,omitempty" name:"MonthlyTraffic"`

	// 是否支持 Linux/Unix 平台。
	SupportLinuxUnixPlatform *bool `json:"SupportLinuxUnixPlatform,omitempty" name:"SupportLinuxUnixPlatform"`

	// 是否支持 Windows 平台。
	SupportWindowsPlatform *bool `json:"SupportWindowsPlatform,omitempty" name:"SupportWindowsPlatform"`

	// 套餐当前单位价格信息。
	Price *Price `json:"Price,omitempty" name:"Price"`

	// CPU 核数。
	CPU *int64 `json:"CPU,omitempty" name:"CPU"`

	// 峰值带宽，单位 Mbps。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// 网络计费类型。
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// 套餐售卖状态,取值:‘AVAILABLE’(可用) , ‘SOLD_OUT’(售罄)
	BundleSalesState *string `json:"BundleSalesState,omitempty" name:"BundleSalesState"`

	// 套餐类型。
	// 取值范围：
	// <li> GENERAL_BUNDLE：通用型</li><li> STORAGE_BUNDLE：存储型 </li>
	BundleType *string `json:"BundleType,omitempty" name:"BundleType"`

	// 套餐展示标签.
	// 取值范围:
	// "ACTIVITY": 活动套餐,
	// "NORMAL": 普通套餐
	// "CAREFREE": 无忧套餐
	BundleDisplayLabel *string `json:"BundleDisplayLabel,omitempty" name:"BundleDisplayLabel"`
}

type CreateBlueprintRequest struct {
	*tchttp.BaseRequest

	// 镜像名称。最大长度60。
	BlueprintName *string `json:"BlueprintName,omitempty" name:"BlueprintName"`

	// 镜像描述。最大长度60。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 需要制作镜像的实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CreateBlueprintRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBlueprintRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BlueprintName")
	delete(f, "Description")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return errors.New("CreateBlueprintRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateBlueprintResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 自定义镜像ID。
		BlueprintId *string `json:"BlueprintId,omitempty" name:"BlueprintId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBlueprintResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBlueprintResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFirewallRulesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 防火墙规则列表。
	FirewallRules []*FirewallRule `json:"FirewallRules,omitempty" name:"FirewallRules" list`

	// 防火墙当前版本。用户每次更新防火墙规则时版本会自动加1，防止规则已过期，不填不考虑冲突。
	FirewallVersion *uint64 `json:"FirewallVersion,omitempty" name:"FirewallVersion"`
}

func (r *CreateFirewallRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFirewallRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FirewallRules")
	delete(f, "FirewallVersion")
	if len(f) > 0 {
		return errors.New("CreateFirewallRulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateFirewallRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFirewallRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFirewallRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceSnapshotRequest struct {
	*tchttp.BaseRequest

	// 需要创建快照的实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 快照名称，最长为 60 个字符。
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
}

func (r *CreateInstanceSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SnapshotName")
	if len(f) > 0 {
		return errors.New("CreateInstanceSnapshotRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 快照 ID。
		SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBlueprintsRequest struct {
	*tchttp.BaseRequest

	// 镜像ID列表。镜像ID，可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值中的BlueprintId获取。
	BlueprintIds []*string `json:"BlueprintIds,omitempty" name:"BlueprintIds" list`
}

func (r *DeleteBlueprintsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlueprintsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BlueprintIds")
	if len(f) > 0 {
		return errors.New("DeleteBlueprintsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBlueprintsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBlueprintsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlueprintsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFirewallRulesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 防火墙规则列表。
	FirewallRules []*FirewallRule `json:"FirewallRules,omitempty" name:"FirewallRules" list`

	// 防火墙当前版本。用户每次更新防火墙规则时版本会自动加1，防止规则已过期，不填不考虑冲突。
	FirewallVersion *uint64 `json:"FirewallVersion,omitempty" name:"FirewallVersion"`
}

func (r *DeleteFirewallRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFirewallRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FirewallRules")
	delete(f, "FirewallVersion")
	if len(f) > 0 {
		return errors.New("DeleteFirewallRulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFirewallRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFirewallRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFirewallRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotsRequest struct {
	*tchttp.BaseRequest

	// 要删除的快照 ID 列表，可通过 DescribeSnapshots 查询。
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds" list`
}

func (r *DeleteSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotIds")
	if len(f) > 0 {
		return errors.New("DeleteSnapshotsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlueprintsRequest struct {
	*tchttp.BaseRequest

	// 镜像 ID 列表。
	BlueprintIds []*string `json:"BlueprintIds,omitempty" name:"BlueprintIds" list`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤器列表。
	// <li>blueprint-id</li>按照【镜像 ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>blueprint-type</li>按照【镜像类型】进行过滤。
	// 取值：APP_OS（预置应用的系统 ）；PURE_OS（纯净的 OS 系统）；PRIVATE（自定义镜像）。
	// 类型：String
	// 必选：否
	// <li>platform-type</li>按照【镜像平台类型】进行过滤。
	// 取值： LINUX_UNIX（Linux/Unix系统）；WINDOWS（Windows 系统）。
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。参数不支持同时指定 BlueprintIds 和 Filters 。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeBlueprintsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlueprintsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BlueprintIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return errors.New("DescribeBlueprintsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlueprintsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的镜像数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 镜像详细信息列表。
		BlueprintSet []*Blueprint `json:"BlueprintSet,omitempty" name:"BlueprintSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlueprintsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlueprintsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBundlesRequest struct {
	*tchttp.BaseRequest

	// 套餐 ID 列表。
	BundleIds []*string `json:"BundleIds,omitempty" name:"BundleIds" list`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤器列表。
	// <li>bundle-id</li>按照【套餐 ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>support-platform-type</li>按照【系统类型】进行过滤。
	// 取值： LINUX_UNIX（Linux/Unix系统）；WINDOWS（Windows 系统）
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。参数不支持同时指定 BundleIds 和 Filters。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeBundlesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBundlesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BundleIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return errors.New("DescribeBundlesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBundlesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 套餐详细信息列表。
		BundleSet []*Bundle `json:"BundleSet,omitempty" name:"BundleSet" list`

		// 符合要求的套餐总数，用于分页展示。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBundlesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBundlesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFirewallRulesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeFirewallRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return errors.New("DescribeFirewallRulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFirewallRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的防火墙规则数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 防火墙规则详细信息列表。
		FirewallRuleSet []*FirewallRuleInfo `json:"FirewallRuleSet,omitempty" name:"FirewallRuleSet" list`

		// 防火墙版本号。
		FirewallVersion *uint64 `json:"FirewallVersion,omitempty" name:"FirewallVersion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFirewallRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 列表。每次请求批量实例的上限为 100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 过滤器列表。
	// <li>instance-name</li>按照【实例名称】进行过滤。
	// 类型：String
	// 必选：否
	// <li>private-ip-address</li>按照【实例主网卡的内网 IP】进行过滤。
	// 类型：String
	// 必选：否
	// <li>public-ip-address</li>按照【实例主网卡的公网 IP】进行过滤。
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。参数不支持同时指定 InstanceIds 和 Filters。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的实例数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例详细信息列表。
		InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesTrafficPackagesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesTrafficPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesTrafficPackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return errors.New("DescribeInstancesTrafficPackagesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesTrafficPackagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的实例流量包详情数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例流量包详情列表。
		InstanceTrafficPackageSet []*InstanceTrafficPackage `json:"InstanceTrafficPackageSet,omitempty" name:"InstanceTrafficPackageSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesTrafficPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesTrafficPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotsRequest struct {
	*tchttp.BaseRequest

	// 要查询快照的 ID 列表。
	// 参数不支持同时指定 SnapshotIds 和 Filters。
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds" list`

	// 过滤器列表。
	// <li>snapshot-id</li>按照【快照 ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>disk-id</li>按照【磁盘 ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>snapshot-name</li>按照【快照名称】进行过滤。
	// 类型：String
	// 必选：否
	// <li>instance-id</li>按照【实例 ID 】进行过滤。
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。参数不支持同时指定 SnapshotIds 和 Filters。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return errors.New("DescribeSnapshotsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 快照的数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 快照的详情列表。
		SnapshotSet []*Snapshot `json:"SnapshotSet,omitempty" name:"SnapshotSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type FirewallRule struct {

	// 协议，取值：TCP，UDP，ICMP，ALL。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 端口，取值：ALL，单独的端口，逗号分隔的离散端口，减号分隔的端口范围。
	Port *string `json:"Port,omitempty" name:"Port"`

	// 网段或 IP (互斥)。默认为 0.0.0.0/0，表示所有来源。
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 取值：ACCEPT，DROP。默认为 ACCEPT。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 防火墙规则描述。
	FirewallRuleDescription *string `json:"FirewallRuleDescription,omitempty" name:"FirewallRuleDescription"`
}

type FirewallRuleInfo struct {

	// 应用类型，取值：自定义，HTTP(80)，HTTPS(443)，Linux登录(22)，Windows登录(3389)，MySQL(3306)，SQL Server(1433)，全部TCP，全部UDP，Ping-ICMP，ALL。
	AppType *string `json:"AppType,omitempty" name:"AppType"`

	// 协议，取值：TCP，UDP，ICMP，ALL。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 端口，取值：ALL，单独的端口，逗号分隔的离散端口，减号分隔的端口范围。
	Port *string `json:"Port,omitempty" name:"Port"`

	// 网段或 IP (互斥)。默认为 0.0.0.0/0，表示所有来源。
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 取值：ACCEPT，DROP。默认为 ACCEPT。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 防火墙规则描述。
	FirewallRuleDescription *string `json:"FirewallRuleDescription,omitempty" name:"FirewallRuleDescription"`
}

type Instance struct {

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 套餐 ID。
	BundleId *string `json:"BundleId,omitempty" name:"BundleId"`

	// 镜像 ID。
	BlueprintId *string `json:"BlueprintId,omitempty" name:"BlueprintId"`

	// 实例的 CPU 核数，单位：核。
	CPU *int64 `json:"CPU,omitempty" name:"CPU"`

	// 实例内存容量，单位：GB 。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例计费模式。取值范围： 
	// PREPAID：表示预付费，即包年包月。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 实例系统盘信息。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 实例主网卡的内网 IP。 
	// 注意：此字段可能返回 空，表示取不到有效值。
	PrivateAddresses []*string `json:"PrivateAddresses,omitempty" name:"PrivateAddresses" list`

	// 实例主网卡的公网 IP。 
	// 注意：此字段可能返回 空，表示取不到有效值。
	PublicAddresses []*string `json:"PublicAddresses,omitempty" name:"PublicAddresses" list`

	// 实例带宽信息。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 自动续费标识。取值范围： 
	// NOTIFY_AND_MANUAL_RENEW：表示通知即将过期，但不自动续费  
	// NOTIFY_AND_AUTO_RENEW：表示通知即将过期，而且自动续费 。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 实例登录设置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 实例状态。取值范围： 
	// <li>PENDING：表示创建中</li><li>LAUNCH_FAILED：表示创建失败</li><li>RUNNING：表示运行中</li><li>STOPPED：表示关机</li><li>STARTING：表示开机中</li><li>STOPPING：表示关机中</li><li>REBOOTING：表示重启中</li><li>SHUTDOWN：表示停止待销毁</li><li>TERMINATING：表示销毁中</li>
	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`

	// 实例全局唯一 ID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 实例的最新操作。例：StopInstances、ResetInstance。注意：此字段可能返回 空值，表示取不到有效值。
	LatestOperation *string `json:"LatestOperation,omitempty" name:"LatestOperation"`

	// 实例的最新操作状态。取值范围： 
	// SUCCESS：表示操作成功 
	// OPERATING：表示操作执行中 
	// FAILED：表示操作失败 
	// 注意：此字段可能返回 空值，表示取不到有效值。
	LatestOperationState *string `json:"LatestOperationState,omitempty" name:"LatestOperationState"`

	// 实例最新操作的唯一请求 ID。 
	// 注意：此字段可能返回 空值，表示取不到有效值。
	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitempty" name:"LatestOperationRequestId"`

	// 隔离时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolatedTime *string `json:"IsolatedTime,omitempty" name:"IsolatedTime"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 到期时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 操作系统平台类型，如 LINUX_UNIX、WINDOWS。
	PlatformType *string `json:"PlatformType,omitempty" name:"PlatformType"`

	// 操作系统平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 操作系统名称。
	OsName *string `json:"OsName,omitempty" name:"OsName"`
}

type InstancePrice struct {

	// 套餐单价原价。
	OriginalBundlePrice *float64 `json:"OriginalBundlePrice,omitempty" name:"OriginalBundlePrice"`

	// 原价。
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 折扣。
	Discount *int64 `json:"Discount,omitempty" name:"Discount"`

	// 折后价。
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
}

type InstanceTrafficPackage struct {

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 流量包详情列表。
	TrafficPackageSet []*TrafficPackage `json:"TrafficPackageSet,omitempty" name:"TrafficPackageSet" list`
}

type InternetAccessible struct {

	// 网络计费类型。
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// 公网出带宽上限，单位：Mbps。
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// 是否分配公网 IP。
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty" name:"PublicIpAssigned"`
}

type LoginSettings struct {

	// 密钥 ID 列表。关联密钥后，就可以通过对应的私钥来访问实例。注意：此字段可能返回 []，表示取不到有效值。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`
}

type ModifyBlueprintAttributeRequest struct {
	*tchttp.BaseRequest

	// 镜像 ID。可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值中的BlueprintId获取。
	BlueprintId *string `json:"BlueprintId,omitempty" name:"BlueprintId"`

	// 设置新的镜像名称。最大长度60。
	BlueprintName *string `json:"BlueprintName,omitempty" name:"BlueprintName"`

	// 设置新的镜像描述。最大长度60。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyBlueprintAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlueprintAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BlueprintId")
	delete(f, "BlueprintName")
	delete(f, "Description")
	if len(f) > 0 {
		return errors.New("ModifyBlueprintAttributeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBlueprintAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBlueprintAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlueprintAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySnapshotAttributeRequest struct {
	*tchttp.BaseRequest

	// 快照 ID, 可通过 DescribeSnapshots 查询。
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 新的快照名称，最长为 60 个字符。
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
}

func (r *ModifySnapshotAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotId")
	delete(f, "SnapshotName")
	if len(f) > 0 {
		return errors.New("ModifySnapshotAttributeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySnapshotAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySnapshotAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Price struct {

	// 实例价格。
	InstancePrice *InstancePrice `json:"InstancePrice,omitempty" name:"InstancePrice"`
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *RebootInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebootInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return errors.New("RebootInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RebootInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RebootInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebootInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例 ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 镜像 ID。可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值中的BlueprintId获取。
	BlueprintId *string `json:"BlueprintId,omitempty" name:"BlueprintId"`
}

func (r *ResetInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BlueprintId")
	if len(f) > 0 {
		return errors.New("ResetInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Snapshot struct {

	// 快照 ID。
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 创建此快照的磁盘类型。取值：<li>SYSTEM_DISK：系统盘</li>
	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`

	// 创建此快照的磁盘 ID。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 创建此快照的磁盘大小，单位 GB。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 快照名称，用户自定义的快照别名。
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// 快照的状态。取值范围：
	// <li>NORMAL：正常 </li>
	// <li>CREATING：创建中</li>
	// <li>ROLLBACKING：回滚中。</li>
	SnapshotState *string `json:"SnapshotState,omitempty" name:"SnapshotState"`

	// 创建或回滚快照进度百分比，成功后此字段取值为 100。
	Percent *int64 `json:"Percent,omitempty" name:"Percent"`

	// 快照的最新操作，只有创建、回滚快照时记录。
	// 取值如 CreateInstanceSnapshot，RollbackInstanceSnapshot。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperation *string `json:"LatestOperation,omitempty" name:"LatestOperation"`

	// 快照的最新操作状态，只有创建、回滚快照时记录。
	// 取值范围：
	// <li>SUCCESS：表示操作成功</li>
	// <li>OPERATING：表示操作执行中</li>
	// <li>FAILED：表示操作失败</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperationState *string `json:"LatestOperationState,omitempty" name:"LatestOperationState"`

	// 快照最新操作的唯一请求 ID，只有创建、回滚快照时记录。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitempty" name:"LatestOperationRequestId"`

	// 快照的创建时间。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type StartInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *StartInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return errors.New("StartInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *StopInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return errors.New("StopInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StopInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemDisk struct {

	// 系统盘类型。
	// 取值范围： 
	// <li> LOCAL_BASIC：本地硬盘</li><li> LOCAL_SSD：本地 SSD 硬盘</li><li> CLOUD_BASIC：普通云硬盘</li><li> CLOUD_SSD：SSD 云硬盘</li><li> CLOUD_PREMIUM：高性能云硬盘</li>
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 系统盘大小，单位：GB。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 系统盘ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

type TrafficPackage struct {

	// 流量包ID。
	TrafficPackageId *string `json:"TrafficPackageId,omitempty" name:"TrafficPackageId"`

	// 流量包生效周期内已使用流量，单位字节。
	TrafficUsed *int64 `json:"TrafficUsed,omitempty" name:"TrafficUsed"`

	// 流量包生效周期内的总流量，单位字节。
	TrafficPackageTotal *int64 `json:"TrafficPackageTotal,omitempty" name:"TrafficPackageTotal"`

	// 流量包生效周期内的剩余流量，单位字节。
	TrafficPackageRemaining *int64 `json:"TrafficPackageRemaining,omitempty" name:"TrafficPackageRemaining"`

	// 流量包生效周期内超出流量包额度的流量，单位字节。
	TrafficOverflow *int64 `json:"TrafficOverflow,omitempty" name:"TrafficOverflow"`

	// 流量包生效周期开始时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 流量包生效周期结束时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 流量包到期时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`

	// 流量包状态：
	// <li>NETWORK_NORMAL：正常</li>
	// <li>OVERDUE_NETWORK_DISABLED：欠费断网</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}
