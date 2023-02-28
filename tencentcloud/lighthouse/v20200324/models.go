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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type ApplyDiskBackupRequestParams struct {
	// 云硬盘ID，可通过[DescribeDisks](https://cloud.tencent.com/document/api/1207/66093)接口查询。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 云硬盘备份点ID，可通过[DescribeDiskBackups](https://cloud.tencent.com/document/api/1207/84379)接口查询。
	DiskBackupId *string `json:"DiskBackupId,omitempty" name:"DiskBackupId"`
}

type ApplyDiskBackupRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID，可通过[DescribeDisks](https://cloud.tencent.com/document/api/1207/66093)接口查询。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 云硬盘备份点ID，可通过[DescribeDiskBackups](https://cloud.tencent.com/document/api/1207/84379)接口查询。
	DiskBackupId *string `json:"DiskBackupId,omitempty" name:"DiskBackupId"`
}

func (r *ApplyDiskBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyDiskBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	delete(f, "DiskBackupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyDiskBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyDiskBackupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyDiskBackupResponse struct {
	*tchttp.BaseResponse
	Response *ApplyDiskBackupResponseParams `json:"Response"`
}

func (r *ApplyDiskBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyDiskBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyInstanceSnapshotRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 快照 ID。
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyInstanceSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SnapshotId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyInstanceSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyInstanceSnapshotResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyInstanceSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *ApplyInstanceSnapshotResponseParams `json:"Response"`
}

func (r *ApplyInstanceSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyInstanceSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateInstancesKeyPairsRequestParams struct {
	// 密钥对 ID 列表。每次请求批量密钥对的上限为 100。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type AssociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// 密钥对 ID 列表。每次请求批量密钥对的上限为 100。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *AssociateInstancesKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateInstancesKeyPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateInstancesKeyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateInstancesKeyPairsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssociateInstancesKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateInstancesKeyPairsResponseParams `json:"Response"`
}

func (r *AssociateInstancesKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateInstancesKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachCcnRequestParams struct {
	// 云联网实例ID。
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

type AttachCcnRequest struct {
	*tchttp.BaseRequest
	
	// 云联网实例ID。
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *AttachCcnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachCcnRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachCcnRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachCcnResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AttachCcnResponse struct {
	*tchttp.BaseResponse
	Response *AttachCcnResponseParams `json:"Response"`
}

func (r *AttachCcnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachCcnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachDetail struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例已挂载弹性云盘数量
	AttachedDiskCount *int64 `json:"AttachedDiskCount,omitempty" name:"AttachedDiskCount"`

	// 可挂载弹性云盘数量
	MaxAttachCount *int64 `json:"MaxAttachCount,omitempty" name:"MaxAttachCount"`
}

// Predefined struct for user
type AttachDisksRequestParams struct {
	// 云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 续费标识。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type AttachDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 续费标识。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

func (r *AttachDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "InstanceId")
	delete(f, "RenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachDisksResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AttachDisksResponse struct {
	*tchttp.BaseResponse
	Response *AttachDisksResponseParams `json:"Response"`
}

func (r *AttachDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachDisksResponse) FromJsonString(s string) error {
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
	// 注意：此字段可能返回 null，表示取不到有效值。
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

	// 镜像所需系统盘大小，单位 GB。
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

	// CVM镜像共享到轻量应用服务器轻量应用服务器后的CVM镜像ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 官方网站Url。
	CommunityUrl *string `json:"CommunityUrl,omitempty" name:"CommunityUrl"`

	// 指导文章Url。
	GuideUrl *string `json:"GuideUrl,omitempty" name:"GuideUrl"`

	// 镜像关联使用场景Id列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneIdSet []*string `json:"SceneIdSet,omitempty" name:"SceneIdSet"`

	// Docker版本号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`
}

type BlueprintInstance struct {
	// 镜像信息。
	Blueprint *Blueprint `json:"Blueprint,omitempty" name:"Blueprint"`

	// 软件列表。
	SoftwareSet []*Software `json:"SoftwareSet,omitempty" name:"SoftwareSet"`

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type BlueprintPrice struct {
	// 镜像单价，原价。单位元。
	OriginalBlueprintPrice *float64 `json:"OriginalBlueprintPrice,omitempty" name:"OriginalBlueprintPrice"`

	// 镜像总价，原价。单位元。
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 折扣。
	Discount *float64 `json:"Discount,omitempty" name:"Discount"`

	// 镜像折扣后总价。单位元。
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
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

type CcnAttachedInstance struct {
	// 云联网ID。
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// 关联实例CIDR。
	CidrBlock []*string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 关联实例状态：
	// 
	// •  PENDING：申请中
	// •  ACTIVE：已连接
	// •  EXPIRED：已过期
	// •  REJECTED：已拒绝
	// •  DELETED：已删除
	// •  FAILED：失败的（2小时后将异步强制解关联）
	// •  ATTACHING：关联中
	// •  DETACHING：解关联中
	// •  DETACHFAILED：解关联失败（2小时后将异步强制解关联）
	State *string `json:"State,omitempty" name:"State"`

	// 关联时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachedTime *string `json:"AttachedTime,omitempty" name:"AttachedTime"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ContainerEnv struct {
	// 环境变量Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 环境变量值
	Value *string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type CreateBlueprintRequestParams struct {
	// 镜像名称。最大长度60。
	BlueprintName *string `json:"BlueprintName,omitempty" name:"BlueprintName"`

	// 镜像描述。最大长度60。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 需要制作镜像的实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBlueprintRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBlueprintResponseParams struct {
	// 自定义镜像ID。
	BlueprintId *string `json:"BlueprintId,omitempty" name:"BlueprintId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBlueprintResponse struct {
	*tchttp.BaseResponse
	Response *CreateBlueprintResponseParams `json:"Response"`
}

func (r *CreateBlueprintResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBlueprintResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDiskBackupRequestParams struct {
	// 云硬盘 ID。当前只支持数据盘创建备份点。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 云硬盘备份点名称，最大长度90。
	DiskBackupName *string `json:"DiskBackupName,omitempty" name:"DiskBackupName"`
}

type CreateDiskBackupRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘 ID。当前只支持数据盘创建备份点。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 云硬盘备份点名称，最大长度90。
	DiskBackupName *string `json:"DiskBackupName,omitempty" name:"DiskBackupName"`
}

func (r *CreateDiskBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDiskBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	delete(f, "DiskBackupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDiskBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDiskBackupResponseParams struct {
	// 备份点ID。
	DiskBackupId *string `json:"DiskBackupId,omitempty" name:"DiskBackupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDiskBackupResponse struct {
	*tchttp.BaseResponse
	Response *CreateDiskBackupResponseParams `json:"Response"`
}

func (r *CreateDiskBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDiskBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFirewallRulesRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 防火墙规则列表。
	FirewallRules []*FirewallRule `json:"FirewallRules,omitempty" name:"FirewallRules"`

	// 防火墙当前版本。用户每次更新防火墙规则时版本会自动加1，防止规则已过期，不填不考虑冲突。
	FirewallVersion *uint64 `json:"FirewallVersion,omitempty" name:"FirewallVersion"`
}

type CreateFirewallRulesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 防火墙规则列表。
	FirewallRules []*FirewallRule `json:"FirewallRules,omitempty" name:"FirewallRules"`

	// 防火墙当前版本。用户每次更新防火墙规则时版本会自动加1，防止规则已过期，不填不考虑冲突。
	FirewallVersion *uint64 `json:"FirewallVersion,omitempty" name:"FirewallVersion"`
}

func (r *CreateFirewallRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFirewallRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFirewallRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFirewallRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateFirewallRulesResponseParams `json:"Response"`
}

func (r *CreateFirewallRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFirewallRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceSnapshotRequestParams struct {
	// 需要创建快照的实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 快照名称，最长为 60 个字符。
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SnapshotName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceSnapshotResponseParams struct {
	// 快照 ID。
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateInstanceSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceSnapshotResponseParams `json:"Response"`
}

func (r *CreateInstanceSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesRequestParams struct {
	// 套餐ID。
	BundleId *string `json:"BundleId,omitempty" name:"BundleId"`

	// 镜像ID。
	BlueprintId *string `json:"BlueprintId,omitempty" name:"BlueprintId"`

	// 当前实例仅支持预付费模式，即包年包月相关参数设置，单位（月）。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 实例显示名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 购买实例数量。包年包月实例取值范围：[1，30]。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 可用区列表。默认为随机可用区
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和库存。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId.
	// false（默认）：发送正常请求，通过检查后直接创建实例
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 实例登录密码信息配置。本字段目前仅支持WINDOWS实例进行密码设置。默认缺失情况下代表用户选择实例创建后设置登录密码。
	LoginConfiguration *LoginConfiguration `json:"LoginConfiguration,omitempty" name:"LoginConfiguration"`

	// 要创建的容器配置列表。
	Containers []*DockerContainerConfiguration `json:"Containers,omitempty" name:"Containers"`

	// 是否自动使用代金券。默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitempty" name:"AutoVoucher"`
}

type CreateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 套餐ID。
	BundleId *string `json:"BundleId,omitempty" name:"BundleId"`

	// 镜像ID。
	BlueprintId *string `json:"BlueprintId,omitempty" name:"BlueprintId"`

	// 当前实例仅支持预付费模式，即包年包月相关参数设置，单位（月）。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 实例显示名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 购买实例数量。包年包月实例取值范围：[1，30]。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 可用区列表。默认为随机可用区
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和库存。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId.
	// false（默认）：发送正常请求，通过检查后直接创建实例
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 实例登录密码信息配置。本字段目前仅支持WINDOWS实例进行密码设置。默认缺失情况下代表用户选择实例创建后设置登录密码。
	LoginConfiguration *LoginConfiguration `json:"LoginConfiguration,omitempty" name:"LoginConfiguration"`

	// 要创建的容器配置列表。
	Containers []*DockerContainerConfiguration `json:"Containers,omitempty" name:"Containers"`

	// 是否自动使用代金券。默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitempty" name:"AutoVoucher"`
}

func (r *CreateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BundleId")
	delete(f, "BlueprintId")
	delete(f, "InstanceChargePrepaid")
	delete(f, "InstanceName")
	delete(f, "InstanceCount")
	delete(f, "Zones")
	delete(f, "DryRun")
	delete(f, "ClientToken")
	delete(f, "LoginConfiguration")
	delete(f, "Containers")
	delete(f, "AutoVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesResponseParams struct {
	// 当通过本接口来创建实例时会返回该参数，表示一个或多个实例ID。返回实例ID列表并不代表实例创建成功。
	// 
	// 可根据 DescribeInstances 接口查询返回的InstancesSet中对应实例的ID的状态来判断创建是否完成；如果实例状态由“启动中”变为“运行中”，则为创建成功。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstancesResponseParams `json:"Response"`
}

func (r *CreateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeyPairRequestParams struct {
	// 密钥对名称，可由数字，字母和下划线组成，长度不超过 25 个字符。
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
}

type CreateKeyPairRequest struct {
	*tchttp.BaseRequest
	
	// 密钥对名称，可由数字，字母和下划线组成，长度不超过 25 个字符。
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
}

func (r *CreateKeyPairRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKeyPairRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKeyPairRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeyPairResponseParams struct {
	// 密钥对信息。
	KeyPair *KeyPair `json:"KeyPair,omitempty" name:"KeyPair"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateKeyPairResponse struct {
	*tchttp.BaseResponse
	Response *CreateKeyPairResponseParams `json:"Response"`
}

func (r *CreateKeyPairResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKeyPairResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataDiskPrice struct {
	// 云硬盘ID。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 云硬盘单价。
	OriginalDiskPrice *float64 `json:"OriginalDiskPrice,omitempty" name:"OriginalDiskPrice"`

	// 云硬盘总价。
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 折扣。
	Discount *float64 `json:"Discount,omitempty" name:"Discount"`

	// 折后总价。
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

	// 数据盘挂载的实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

// Predefined struct for user
type DeleteBlueprintsRequestParams struct {
	// 镜像ID列表。镜像ID，可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值中的BlueprintId获取。
	BlueprintIds []*string `json:"BlueprintIds,omitempty" name:"BlueprintIds"`
}

type DeleteBlueprintsRequest struct {
	*tchttp.BaseRequest
	
	// 镜像ID列表。镜像ID，可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值中的BlueprintId获取。
	BlueprintIds []*string `json:"BlueprintIds,omitempty" name:"BlueprintIds"`
}

func (r *DeleteBlueprintsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlueprintsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BlueprintIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBlueprintsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBlueprintsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteBlueprintsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBlueprintsResponseParams `json:"Response"`
}

func (r *DeleteBlueprintsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlueprintsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDiskBackupsRequestParams struct {
	// 云硬盘备份点ID列表，可通过 [DescribeDiskBackups](https://cloud.tencent.com/document/api/1207/84379)接口查询。
	DiskBackupIds []*string `json:"DiskBackupIds,omitempty" name:"DiskBackupIds"`
}

type DeleteDiskBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘备份点ID列表，可通过 [DescribeDiskBackups](https://cloud.tencent.com/document/api/1207/84379)接口查询。
	DiskBackupIds []*string `json:"DiskBackupIds,omitempty" name:"DiskBackupIds"`
}

func (r *DeleteDiskBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDiskBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskBackupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDiskBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDiskBackupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDiskBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDiskBackupsResponseParams `json:"Response"`
}

func (r *DeleteDiskBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDiskBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFirewallRulesRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 防火墙规则列表。
	FirewallRules []*FirewallRule `json:"FirewallRules,omitempty" name:"FirewallRules"`

	// 防火墙当前版本。用户每次更新防火墙规则时版本会自动加1，防止规则已过期，不填不考虑冲突。
	FirewallVersion *uint64 `json:"FirewallVersion,omitempty" name:"FirewallVersion"`
}

type DeleteFirewallRulesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 防火墙规则列表。
	FirewallRules []*FirewallRule `json:"FirewallRules,omitempty" name:"FirewallRules"`

	// 防火墙当前版本。用户每次更新防火墙规则时版本会自动加1，防止规则已过期，不填不考虑冲突。
	FirewallVersion *uint64 `json:"FirewallVersion,omitempty" name:"FirewallVersion"`
}

func (r *DeleteFirewallRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFirewallRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFirewallRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteFirewallRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFirewallRulesResponseParams `json:"Response"`
}

func (r *DeleteFirewallRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFirewallRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteKeyPairsRequestParams struct {
	// 密钥对 ID 列表，每次请求批量密钥对的上限为 10。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

type DeleteKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// 密钥对 ID 列表，每次请求批量密钥对的上限为 10。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

func (r *DeleteKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKeyPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteKeyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteKeyPairsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteKeyPairsResponseParams `json:"Response"`
}

func (r *DeleteKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotsRequestParams struct {
	// 要删除的快照 ID 列表，可通过 DescribeSnapshots 查询。
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
}

type DeleteSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的快照 ID 列表，可通过 DescribeSnapshots 查询。
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
}

func (r *DeleteSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSnapshotsResponseParams `json:"Response"`
}

func (r *DeleteSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeniedAction struct {
	// 限制操作名。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 限制操作消息码。
	Code *string `json:"Code,omitempty" name:"Code"`

	// 限制操作消息。
	Message *string `json:"Message,omitempty" name:"Message"`
}

// Predefined struct for user
type DescribeAllScenesRequestParams struct {
	// 使用场景ID列表。
	SceneIds []*string `json:"SceneIds,omitempty" name:"SceneIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeAllScenesRequest struct {
	*tchttp.BaseRequest
	
	// 使用场景ID列表。
	SceneIds []*string `json:"SceneIds,omitempty" name:"SceneIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAllScenesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllScenesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SceneIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllScenesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllScenesResponseParams struct {
	// 使用场景详细信息列表。
	SceneInfoSet []*SceneInfo `json:"SceneInfoSet,omitempty" name:"SceneInfoSet"`

	// 使用场景详细信息总数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAllScenesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllScenesResponseParams `json:"Response"`
}

func (r *DescribeAllScenesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllScenesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlueprintInstancesRequestParams struct {
	// 实例 ID 列表，当前最多支持 1 个。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DescribeBlueprintInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表，当前最多支持 1 个。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeBlueprintInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlueprintInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlueprintInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlueprintInstancesResponseParams struct {
	// 符合条件的镜像实例数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 镜像实例列表信息。
	BlueprintInstanceSet []*BlueprintInstance `json:"BlueprintInstanceSet,omitempty" name:"BlueprintInstanceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBlueprintInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBlueprintInstancesResponseParams `json:"Response"`
}

func (r *DescribeBlueprintInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlueprintInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlueprintsRequestParams struct {
	// 镜像 ID 列表。
	BlueprintIds []*string `json:"BlueprintIds,omitempty" name:"BlueprintIds"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤器列表。
	// <li>blueprint-id</li>按照【镜像 ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>blueprint-type</li>按照【镜像类型】进行过滤。
	// 取值：APP_OS（应用镜像 ）；PURE_OS（系统镜像）；DOCKER（Docker容器镜像）；PRIVATE（自定义镜像）；SHARED（共享镜像）。
	// 类型：String
	// 必选：否
	// <li>platform-type</li>按照【镜像平台类型】进行过滤。
	// 取值： LINUX_UNIX（Linux/Unix系统）；WINDOWS（Windows 系统）。
	// 类型：String
	// 必选：否
	// <li>blueprint-name</li>按照【镜像名称】进行过滤。
	// 类型：String
	// 必选：否
	// <li>blueprint-state</li>按照【镜像状态】进行过滤。
	// 类型：String
	// 必选：否
	// <li>scene-id</li>按照【使用场景Id】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 100。参数不支持同时指定 BlueprintIds 和 Filters 。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeBlueprintsRequest struct {
	*tchttp.BaseRequest
	
	// 镜像 ID 列表。
	BlueprintIds []*string `json:"BlueprintIds,omitempty" name:"BlueprintIds"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤器列表。
	// <li>blueprint-id</li>按照【镜像 ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>blueprint-type</li>按照【镜像类型】进行过滤。
	// 取值：APP_OS（应用镜像 ）；PURE_OS（系统镜像）；DOCKER（Docker容器镜像）；PRIVATE（自定义镜像）；SHARED（共享镜像）。
	// 类型：String
	// 必选：否
	// <li>platform-type</li>按照【镜像平台类型】进行过滤。
	// 取值： LINUX_UNIX（Linux/Unix系统）；WINDOWS（Windows 系统）。
	// 类型：String
	// 必选：否
	// <li>blueprint-name</li>按照【镜像名称】进行过滤。
	// 类型：String
	// 必选：否
	// <li>blueprint-state</li>按照【镜像状态】进行过滤。
	// 类型：String
	// 必选：否
	// <li>scene-id</li>按照【使用场景Id】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 100。参数不支持同时指定 BlueprintIds 和 Filters 。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeBlueprintsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlueprintsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlueprintsResponseParams struct {
	// 符合条件的镜像数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 镜像详细信息列表。
	BlueprintSet []*Blueprint `json:"BlueprintSet,omitempty" name:"BlueprintSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBlueprintsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBlueprintsResponseParams `json:"Response"`
}

func (r *DescribeBlueprintsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlueprintsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBundleDiscountRequestParams struct {
	// 套餐 ID。
	BundleId *string `json:"BundleId,omitempty" name:"BundleId"`
}

type DescribeBundleDiscountRequest struct {
	*tchttp.BaseRequest
	
	// 套餐 ID。
	BundleId *string `json:"BundleId,omitempty" name:"BundleId"`
}

func (r *DescribeBundleDiscountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBundleDiscountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BundleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBundleDiscountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBundleDiscountResponseParams struct {
	// 币种：CNY人民币，USD 美元。
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 折扣梯度详情，每个梯度包含的信息有：时长，折扣数，总价，折扣价，折扣详情（用户折扣、官网折扣、最终折扣）。
	DiscountDetail []*DiscountDetail `json:"DiscountDetail,omitempty" name:"DiscountDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBundleDiscountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBundleDiscountResponseParams `json:"Response"`
}

func (r *DescribeBundleDiscountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBundleDiscountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBundlesRequestParams struct {
	// 套餐 ID 列表。
	BundleIds []*string `json:"BundleIds,omitempty" name:"BundleIds"`

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
	// <li>bundle-type</li>按照 【套餐类型进行过滤】。
	// 取值：GENERAL_BUNDLE (通用型套餐); STORAGE_BUNDLE(存储型套餐);ENTERPRISE_BUNDLE( 企业型套餐);EXCLUSIVE_BUNDLE(专属型套餐);BEFAST_BUNDLE(蜂驰型套餐);
	// 类型：String
	// 必选：否
	// <li>bundle-state</li>按照【套餐状态】进行过滤。
	// 取值: ‘ONLINE’(在线); ‘OFFLINE’(下线);
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。参数不支持同时指定 BundleIds 和 Filters。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 可用区列表。默认为全部可用区。
	Zones []*string `json:"Zones,omitempty" name:"Zones"`
}

type DescribeBundlesRequest struct {
	*tchttp.BaseRequest
	
	// 套餐 ID 列表。
	BundleIds []*string `json:"BundleIds,omitempty" name:"BundleIds"`

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
	// <li>bundle-type</li>按照 【套餐类型进行过滤】。
	// 取值：GENERAL_BUNDLE (通用型套餐); STORAGE_BUNDLE(存储型套餐);ENTERPRISE_BUNDLE( 企业型套餐);EXCLUSIVE_BUNDLE(专属型套餐);BEFAST_BUNDLE(蜂驰型套餐);
	// 类型：String
	// 必选：否
	// <li>bundle-state</li>按照【套餐状态】进行过滤。
	// 取值: ‘ONLINE’(在线); ‘OFFLINE’(下线);
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。参数不支持同时指定 BundleIds 和 Filters。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 可用区列表。默认为全部可用区。
	Zones []*string `json:"Zones,omitempty" name:"Zones"`
}

func (r *DescribeBundlesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
	delete(f, "Zones")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBundlesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBundlesResponseParams struct {
	// 套餐详细信息列表。
	BundleSet []*Bundle `json:"BundleSet,omitempty" name:"BundleSet"`

	// 符合要求的套餐总数，用于分页展示。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBundlesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBundlesResponseParams `json:"Response"`
}

func (r *DescribeBundlesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBundlesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnAttachedInstancesRequestParams struct {

}

type DescribeCcnAttachedInstancesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCcnAttachedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnAttachedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcnAttachedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnAttachedInstancesResponseParams struct {
	// 云联网关联的实例列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnAttachedInstanceSet []*CcnAttachedInstance `json:"CcnAttachedInstanceSet,omitempty" name:"CcnAttachedInstanceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCcnAttachedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcnAttachedInstancesResponseParams `json:"Response"`
}

func (r *DescribeCcnAttachedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnAttachedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskBackupsDeniedActionsRequestParams struct {
	// 云硬盘备份点 ID 列表, 可通过 DescribeDiskBackups 接口查询。
	DiskBackupIds []*string `json:"DiskBackupIds,omitempty" name:"DiskBackupIds"`
}

type DescribeDiskBackupsDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘备份点 ID 列表, 可通过 DescribeDiskBackups 接口查询。
	DiskBackupIds []*string `json:"DiskBackupIds,omitempty" name:"DiskBackupIds"`
}

func (r *DescribeDiskBackupsDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskBackupsDeniedActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskBackupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiskBackupsDeniedActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskBackupsDeniedActionsResponseParams struct {
	// 云硬盘备份点操作限制列表详细信息。
	DiskBackupDeniedActionSet []*DiskBackupDeniedActions `json:"DiskBackupDeniedActionSet,omitempty" name:"DiskBackupDeniedActionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDiskBackupsDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiskBackupsDeniedActionsResponseParams `json:"Response"`
}

func (r *DescribeDiskBackupsDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskBackupsDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskBackupsRequestParams struct {
	// 要查询云硬盘备份点的ID列表。参数不支持同时指定 DiskBackupIds 和 Filters。
	DiskBackupIds []*string `json:"DiskBackupIds,omitempty" name:"DiskBackupIds"`

	// 过滤器列表。
	// <li>disk-backup-id</li>按照【云硬盘备份点 ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>disk-id</li>按照【云硬盘 ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>disk-backup-state</li>按照【云硬盘备份点状态】进行过滤。
	// 类型：String
	// 必选：否
	// 取值：参考数据结构[DiskBackup](https://cloud.tencent.com/document/product/1207/47576#DiskBackup)下的DiskBackupState取值。
	// <li>disk-usage</li>按照【云硬盘类型】进行过滤。
	// 类型：String
	// 必选：否
	// 取值：SYSTEM_DISK或DATA_DISK
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为5。参数不支持同时指定DiskBackupIds 和 Filters。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDiskBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 要查询云硬盘备份点的ID列表。参数不支持同时指定 DiskBackupIds 和 Filters。
	DiskBackupIds []*string `json:"DiskBackupIds,omitempty" name:"DiskBackupIds"`

	// 过滤器列表。
	// <li>disk-backup-id</li>按照【云硬盘备份点 ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>disk-id</li>按照【云硬盘 ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>disk-backup-state</li>按照【云硬盘备份点状态】进行过滤。
	// 类型：String
	// 必选：否
	// 取值：参考数据结构[DiskBackup](https://cloud.tencent.com/document/product/1207/47576#DiskBackup)下的DiskBackupState取值。
	// <li>disk-usage</li>按照【云硬盘类型】进行过滤。
	// 类型：String
	// 必选：否
	// 取值：SYSTEM_DISK或DATA_DISK
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为5。参数不支持同时指定DiskBackupIds 和 Filters。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDiskBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskBackupIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiskBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskBackupsResponseParams struct {
	// 云硬盘备份点的数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 云硬盘备份点信息列表。
	DiskBackupSet []*DiskBackup `json:"DiskBackupSet,omitempty" name:"DiskBackupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDiskBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiskBackupsResponseParams `json:"Response"`
}

func (r *DescribeDiskBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskConfigsRequestParams struct {
	// 过滤器列表。
	// <li>zone</li>按照【可用区】进行过滤。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeDiskConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器列表。
	// <li>zone</li>按照【可用区】进行过滤。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDiskConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiskConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskConfigsResponseParams struct {
	// 云硬盘配置列表。
	DiskConfigSet []*DiskConfig `json:"DiskConfigSet,omitempty" name:"DiskConfigSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDiskConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiskConfigsResponseParams `json:"Response"`
}

func (r *DescribeDiskConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskDiscountRequestParams struct {
	// 云硬盘类型, 取值: "CLOUD_PREMIUM"。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 云硬盘大小。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 指定云硬盘备份点配额，不传时默认为不带备份点配额。目前只支持不带或设置1个云硬盘备份点配额。
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitempty" name:"DiskBackupQuota"`
}

type DescribeDiskDiscountRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘类型, 取值: "CLOUD_PREMIUM"。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 云硬盘大小。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 指定云硬盘备份点配额，不传时默认为不带备份点配额。目前只支持不带或设置1个云硬盘备份点配额。
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitempty" name:"DiskBackupQuota"`
}

func (r *DescribeDiskDiscountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskDiscountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskType")
	delete(f, "DiskSize")
	delete(f, "DiskBackupQuota")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiskDiscountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskDiscountResponseParams struct {
	// 币种：CNY人民币，USD 美元。
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 折扣梯度详情，每个梯度包含的信息有：时长，折扣数，总价，折扣价，折扣详情（用户折扣、官网折扣、最终折扣）。
	DiscountDetail []*DiscountDetail `json:"DiscountDetail,omitempty" name:"DiscountDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDiskDiscountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiskDiscountResponseParams `json:"Response"`
}

func (r *DescribeDiskDiscountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskDiscountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksDeniedActionsRequestParams struct {
	// 云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

type DescribeDisksDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

func (r *DescribeDisksDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisksDeniedActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisksDeniedActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksDeniedActionsResponseParams struct {
	// 云硬盘操作限制列表详细信息。
	DiskDeniedActionSet []*DiskDeniedActions `json:"DiskDeniedActionSet,omitempty" name:"DiskDeniedActionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDisksDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDisksDeniedActionsResponseParams `json:"Response"`
}

func (r *DescribeDisksDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisksDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksRequestParams struct {
	// 云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// 过滤器列表。
	// disk-id
	// 按照【云硬盘 ID】进行过滤。
	// 类型：String
	// 必选：否
	// instance-id
	// 按照【实例ID】进行过滤。
	// 类型：String
	// 必选：否
	// disk-name
	// 按照【云硬盘名称】进行过滤。
	// 类型：String
	// 必选：否
	// zone
	// 按照【可用区】进行过滤。
	// 类型：String
	// 必选：否
	// disk-usage
	// 按照【云硬盘类型】进行过滤。
	// 类型：String
	// 必选：否
	// 取值：SYSTEM_DISK或DATA_DISK
	// disk-state
	// 按照【云硬盘状态】进行过滤。
	// 类型：String
	// 必选：否
	// 取值：参考数据结构[Disk](https://cloud.tencent.com/document/api/1207/47576#Disk)中DiskState取值。
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 100。参数不支持同时指定 DiskIds 和 Filters。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 云硬盘列表排序的依据字段。取值范围："CREATED_TIME"：依据云硬盘的创建时间排序。 "EXPIRED_TIME"：依据云硬盘的到期时间排序。"DISK_SIZE"：依据云硬盘的大小排序。默认按云硬盘创建时间排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 输出云硬盘列表的排列顺序。取值范围："ASC"：升序排列。 "DESC"：降序排列。默认按降序排列。
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// 过滤器列表。
	// disk-id
	// 按照【云硬盘 ID】进行过滤。
	// 类型：String
	// 必选：否
	// instance-id
	// 按照【实例ID】进行过滤。
	// 类型：String
	// 必选：否
	// disk-name
	// 按照【云硬盘名称】进行过滤。
	// 类型：String
	// 必选：否
	// zone
	// 按照【可用区】进行过滤。
	// 类型：String
	// 必选：否
	// disk-usage
	// 按照【云硬盘类型】进行过滤。
	// 类型：String
	// 必选：否
	// 取值：SYSTEM_DISK或DATA_DISK
	// disk-state
	// 按照【云硬盘状态】进行过滤。
	// 类型：String
	// 必选：否
	// 取值：参考数据结构[Disk](https://cloud.tencent.com/document/api/1207/47576#Disk)中DiskState取值。
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 100。参数不支持同时指定 DiskIds 和 Filters。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 云硬盘列表排序的依据字段。取值范围："CREATED_TIME"：依据云硬盘的创建时间排序。 "EXPIRED_TIME"：依据云硬盘的到期时间排序。"DISK_SIZE"：依据云硬盘的大小排序。默认按云硬盘创建时间排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 输出云硬盘列表的排列顺序。取值范围："ASC"：升序排列。 "DESC"：降序排列。默认按降序排列。
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderField")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksResponseParams struct {
	// 云硬盘信息列表。
	DiskSet []*Disk `json:"DiskSet,omitempty" name:"DiskSet"`

	// 符合条件的云硬盘信息数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDisksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDisksResponseParams `json:"Response"`
}

func (r *DescribeDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksReturnableRequestParams struct {
	// 云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeDisksReturnableRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDisksReturnableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisksReturnableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisksReturnableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksReturnableResponseParams struct {
	// 可退还云硬盘详细信息列表。
	DiskReturnableSet []*DiskReturnable `json:"DiskReturnableSet,omitempty" name:"DiskReturnableSet"`

	// 符合条件的云硬盘数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDisksReturnableResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDisksReturnableResponseParams `json:"Response"`
}

func (r *DescribeDisksReturnableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisksReturnableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirewallRulesRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFirewallRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirewallRulesResponseParams struct {
	// 符合条件的防火墙规则数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 防火墙规则详细信息列表。
	FirewallRuleSet []*FirewallRuleInfo `json:"FirewallRuleSet,omitempty" name:"FirewallRuleSet"`

	// 防火墙版本号。
	FirewallVersion *uint64 `json:"FirewallVersion,omitempty" name:"FirewallVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFirewallRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFirewallRulesResponseParams `json:"Response"`
}

func (r *DescribeFirewallRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirewallRulesTemplateRequestParams struct {

}

type DescribeFirewallRulesTemplateRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeFirewallRulesTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallRulesTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFirewallRulesTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirewallRulesTemplateResponseParams struct {
	// 符合条件的防火墙规则数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 防火墙规则详细信息列表。
	FirewallRuleSet []*FirewallRuleInfo `json:"FirewallRuleSet,omitempty" name:"FirewallRuleSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFirewallRulesTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFirewallRulesTemplateResponseParams `json:"Response"`
}

func (r *DescribeFirewallRulesTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallRulesTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralResourceQuotasRequestParams struct {
	// 资源名列表，可取值:
	// - GENERAL_BUNDLE_INSTANCE 通用型套餐实例
	// - STORAGE_BUNDLE_INSTANCE 存储型套餐实例 
	// - ENTERPRISE_BUNDLE_INSTANCE 企业型套餐实例 
	// - EXCLUSIVE_BUNDLE_INSTANCE 专属型套餐实例
	// - BEFAST_BUNDLE_INSTANCE 蜂驰型套餐实例
	// - USER_KEY_PAIR 密钥对
	// - SNAPSHOT 快照
	// - BLUEPRINT 自定义镜像
	// - FREE_BLUEPRINT 免费自定义镜像
	// - DATA_DISK 数据盘
	// - FIREWALL_RULE 防火墙规则
	ResourceNames []*string `json:"ResourceNames,omitempty" name:"ResourceNames"`
}

type DescribeGeneralResourceQuotasRequest struct {
	*tchttp.BaseRequest
	
	// 资源名列表，可取值:
	// - GENERAL_BUNDLE_INSTANCE 通用型套餐实例
	// - STORAGE_BUNDLE_INSTANCE 存储型套餐实例 
	// - ENTERPRISE_BUNDLE_INSTANCE 企业型套餐实例 
	// - EXCLUSIVE_BUNDLE_INSTANCE 专属型套餐实例
	// - BEFAST_BUNDLE_INSTANCE 蜂驰型套餐实例
	// - USER_KEY_PAIR 密钥对
	// - SNAPSHOT 快照
	// - BLUEPRINT 自定义镜像
	// - FREE_BLUEPRINT 免费自定义镜像
	// - DATA_DISK 数据盘
	// - FIREWALL_RULE 防火墙规则
	ResourceNames []*string `json:"ResourceNames,omitempty" name:"ResourceNames"`
}

func (r *DescribeGeneralResourceQuotasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralResourceQuotasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGeneralResourceQuotasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralResourceQuotasResponseParams struct {
	// 通用资源配额详细信息列表。
	GeneralResourceQuotaSet []*GeneralResourceQuota `json:"GeneralResourceQuotaSet,omitempty" name:"GeneralResourceQuotaSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGeneralResourceQuotasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGeneralResourceQuotasResponseParams `json:"Response"`
}

func (r *DescribeGeneralResourceQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralResourceQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLoginKeyPairAttributeRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeInstanceLoginKeyPairAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceLoginKeyPairAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLoginKeyPairAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceLoginKeyPairAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLoginKeyPairAttributeResponseParams struct {
	// 是否允许使用默认密钥对登录，YES：允许登录 NO：禁止登录。
	PermitLogin *string `json:"PermitLogin,omitempty" name:"PermitLogin"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceLoginKeyPairAttributeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceLoginKeyPairAttributeResponseParams `json:"Response"`
}

func (r *DescribeInstanceLoginKeyPairAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLoginKeyPairAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceVncUrlRequestParams struct {
	// 实例 ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeInstanceVncUrlRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceVncUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceVncUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceVncUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceVncUrlResponseParams struct {
	// 实例的管理终端地址。
	InstanceVncUrl *string `json:"InstanceVncUrl,omitempty" name:"InstanceVncUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceVncUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceVncUrlResponseParams `json:"Response"`
}

func (r *DescribeInstanceVncUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceVncUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesDeniedActionsRequestParams struct {
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DescribeInstancesDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeInstancesDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesDeniedActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesDeniedActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesDeniedActionsResponseParams struct {
	// 实例操作限制列表详细信息。
	InstanceDeniedActionSet []*InstanceDeniedActions `json:"InstanceDeniedActionSet,omitempty" name:"InstanceDeniedActionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstancesDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesDeniedActionsResponseParams `json:"Response"`
}

func (r *DescribeInstancesDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesDiskNumRequestParams struct {
	// 实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DescribeInstancesDiskNumRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeInstancesDiskNumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesDiskNumRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesDiskNumRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesDiskNumResponseParams struct {
	// 挂载信息列表
	AttachDetailSet []*AttachDetail `json:"AttachDetailSet,omitempty" name:"AttachDetailSet"`

	// 挂载信息数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstancesDiskNumResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesDiskNumResponseParams `json:"Response"`
}

func (r *DescribeInstancesDiskNumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesDiskNumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 实例 ID 列表。每次请求批量实例的上限为 100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

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
	// <li>zone</li>按照【可用区】进行过滤。
	// 类型：String
	// 必选：否
	// <li>instance-state</li>按照【实例状态】进行过滤。
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 100。参数不支持同时指定 InstanceIds 和 Filters。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

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
	// <li>zone</li>按照【可用区】进行过滤。
	// 类型：String
	// 必选：否
	// <li>instance-state</li>按照【实例状态】进行过滤。
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 100。参数不支持同时指定 InstanceIds 和 Filters。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	// 符合条件的实例数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例详细信息列表。
	InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeInstancesReturnableRequestParams struct {
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeInstancesReturnableRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesReturnableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesReturnableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesReturnableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesReturnableResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 可退还实例详细信息列表。
	InstanceReturnableSet []*InstanceReturnable `json:"InstanceReturnableSet,omitempty" name:"InstanceReturnableSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstancesReturnableResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesReturnableResponseParams `json:"Response"`
}

func (r *DescribeInstancesReturnableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesReturnableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesTrafficPackagesRequestParams struct {
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeInstancesTrafficPackagesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesTrafficPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesTrafficPackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesTrafficPackagesResponseParams struct {
	// 符合条件的实例流量包详情数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例流量包详情列表。
	InstanceTrafficPackageSet []*InstanceTrafficPackage `json:"InstanceTrafficPackageSet,omitempty" name:"InstanceTrafficPackageSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstancesTrafficPackagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesTrafficPackagesResponseParams `json:"Response"`
}

func (r *DescribeInstancesTrafficPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesTrafficPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeyPairsRequestParams struct {
	// 密钥对 ID 列表。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤器列表。
	// <li>key-id</li>按照【密钥对ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>key-name</li>按照【密钥对名称】进行过滤（支持模糊匹配）。
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。参数不支持同时指定 KeyIds 和 Filters。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// 密钥对 ID 列表。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤器列表。
	// <li>key-id</li>按照【密钥对ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>key-name</li>按照【密钥对名称】进行过滤（支持模糊匹配）。
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。参数不支持同时指定 KeyIds 和 Filters。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeyPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKeyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeyPairsResponseParams struct {
	// 符合条件的密钥对数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 密钥对详细信息列表。
	KeyPairSet []*KeyPair `json:"KeyPairSet,omitempty" name:"KeyPairSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKeyPairsResponseParams `json:"Response"`
}

func (r *DescribeKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModifyInstanceBundlesRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 过滤器列表。
	// <li>bundle-id</li>按照【套餐 ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>support-platform-type</li>按照【系统类型】进行过滤。
	// 取值： LINUX_UNIX（Linux/Unix系统）；WINDOWS（Windows 系统）
	// 类型：String
	// 必选：否
	// <li>bundle-type</li>按照 【套餐类型进行过滤】。
	// 取值：GENERAL_BUNDLE (通用型套餐); STORAGE_BUNDLE(存储型套餐);ENTERPRISE_BUNDLE( 企业型套餐);EXCLUSIVE_BUNDLE(专属型套餐);BEFAST_BUNDLE(蜂驰型套餐);
	// 类型：String
	// 必选：否
	// <li>bundle-state</li>按照【套餐状态】进行过滤。
	// 取值: ‘ONLINE’(在线); ‘OFFLINE’(下线);
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeModifyInstanceBundlesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 过滤器列表。
	// <li>bundle-id</li>按照【套餐 ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>support-platform-type</li>按照【系统类型】进行过滤。
	// 取值： LINUX_UNIX（Linux/Unix系统）；WINDOWS（Windows 系统）
	// 类型：String
	// 必选：否
	// <li>bundle-type</li>按照 【套餐类型进行过滤】。
	// 取值：GENERAL_BUNDLE (通用型套餐); STORAGE_BUNDLE(存储型套餐);ENTERPRISE_BUNDLE( 企业型套餐);EXCLUSIVE_BUNDLE(专属型套餐);BEFAST_BUNDLE(蜂驰型套餐);
	// 类型：String
	// 必选：否
	// <li>bundle-state</li>按照【套餐状态】进行过滤。
	// 取值: ‘ONLINE’(在线); ‘OFFLINE’(下线);
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeModifyInstanceBundlesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModifyInstanceBundlesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModifyInstanceBundlesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModifyInstanceBundlesResponseParams struct {
	// 符合条件的套餐数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 变更套餐详细信息。
	ModifyBundleSet []*ModifyBundle `json:"ModifyBundleSet,omitempty" name:"ModifyBundleSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeModifyInstanceBundlesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModifyInstanceBundlesResponseParams `json:"Response"`
}

func (r *DescribeModifyInstanceBundlesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModifyInstanceBundlesResponse) FromJsonString(s string) error {
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
	// 地域数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 地域信息列表。
	RegionSet []*RegionInfo `json:"RegionSet,omitempty" name:"RegionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeResetInstanceBlueprintsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤器列表。
	// <li>blueprint-id</li>按照【镜像 ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>blueprint-type</li>按照【镜像类型】进行过滤。
	// 取值： APP_OS（应用镜像 ）；PURE_OS（ 系统镜像）；PRIVATE（自定义镜像）。
	// 类型：String
	// 必选：否
	// <li>platform-type</li>按照【镜像平台类型】进行过滤。
	// 取值： LINUX_UNIX（Linux/Unix系统）；WINDOWS（Windows 系统）。
	// 类型：String
	// 必选：否
	// <li>blueprint-name</li>按照【镜像名称】进行过滤。
	// 类型：String
	// 必选：否
	// <li>blueprint-state</li>按照【镜像状态】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。参数不支持同时指定 BlueprintIds 和 Filters 。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeResetInstanceBlueprintsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤器列表。
	// <li>blueprint-id</li>按照【镜像 ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>blueprint-type</li>按照【镜像类型】进行过滤。
	// 取值： APP_OS（应用镜像 ）；PURE_OS（ 系统镜像）；PRIVATE（自定义镜像）。
	// 类型：String
	// 必选：否
	// <li>platform-type</li>按照【镜像平台类型】进行过滤。
	// 取值： LINUX_UNIX（Linux/Unix系统）；WINDOWS（Windows 系统）。
	// 类型：String
	// 必选：否
	// <li>blueprint-name</li>按照【镜像名称】进行过滤。
	// 类型：String
	// 必选：否
	// <li>blueprint-state</li>按照【镜像状态】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。参数不支持同时指定 BlueprintIds 和 Filters 。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeResetInstanceBlueprintsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResetInstanceBlueprintsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResetInstanceBlueprintsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResetInstanceBlueprintsResponseParams struct {
	// 符合条件的镜像数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 镜像重置信息列表
	ResetInstanceBlueprintSet []*ResetInstanceBlueprint `json:"ResetInstanceBlueprintSet,omitempty" name:"ResetInstanceBlueprintSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResetInstanceBlueprintsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResetInstanceBlueprintsResponseParams `json:"Response"`
}

func (r *DescribeResetInstanceBlueprintsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResetInstanceBlueprintsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScenesRequestParams struct {
	// 使用场景ID列表。
	SceneIds []*string `json:"SceneIds,omitempty" name:"SceneIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeScenesRequest struct {
	*tchttp.BaseRequest
	
	// 使用场景ID列表。
	SceneIds []*string `json:"SceneIds,omitempty" name:"SceneIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScenesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScenesResponseParams struct {
	// 使用场景列表。
	SceneSet []*Scene `json:"SceneSet,omitempty" name:"SceneSet"`

	// 使用场景总数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeSnapshotsDeniedActionsRequestParams struct {
	// 快照 ID 列表, 可通过 DescribeSnapshots 查询。
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
}

type DescribeSnapshotsDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// 快照 ID 列表, 可通过 DescribeSnapshots 查询。
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
}

func (r *DescribeSnapshotsDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsDeniedActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotsDeniedActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotsDeniedActionsResponseParams struct {
	// 快照操作限制列表详细信息。
	SnapshotDeniedActionSet []*SnapshotDeniedActions `json:"SnapshotDeniedActionSet,omitempty" name:"SnapshotDeniedActionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSnapshotsDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotsDeniedActionsResponseParams `json:"Response"`
}

func (r *DescribeSnapshotsDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotsRequestParams struct {
	// 要查询快照的 ID 列表。
	// 参数不支持同时指定 SnapshotIds 和 Filters。
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`

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
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// 要查询快照的 ID 列表。
	// 参数不支持同时指定 SnapshotIds 和 Filters。
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`

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
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotsResponseParams struct {
	// 快照的数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 快照的详情列表。
	SnapshotSet []*Snapshot `json:"SnapshotSet,omitempty" name:"SnapshotSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotsResponseParams `json:"Response"`
}

func (r *DescribeSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesRequestParams struct {
	// 可用区列表排序的依据字段。取值范围：
	// <li>ZONE：依据可用区排序。
	// <li>INSTANCE_DISPLAY_LABEL：依据可用区展示标签排序，可用区展示标签包括：HIDDEN（隐藏）、NORMAL（普通）、SELECTED（默认选中），默认采用的升序排列为：['HIDDEN', 'NORMAL', 'SELECTED']。
	// 默认按可用区排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 输出可用区列表的排列顺序。取值范围：
	// <li>ASC：升序排列。 
	// <li>DESC：降序排列。
	// 默认按升序排列。
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
	// 可用区列表排序的依据字段。取值范围：
	// <li>ZONE：依据可用区排序。
	// <li>INSTANCE_DISPLAY_LABEL：依据可用区展示标签排序，可用区展示标签包括：HIDDEN（隐藏）、NORMAL（普通）、SELECTED（默认选中），默认采用的升序排列为：['HIDDEN', 'NORMAL', 'SELECTED']。
	// 默认按可用区排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 输出可用区列表的排列顺序。取值范围：
	// <li>ASC：升序排列。 
	// <li>DESC：降序排列。
	// 默认按升序排列。
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderField")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// 可用区数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 可用区详细信息列表
	ZoneInfoSet []*ZoneInfo `json:"ZoneInfoSet,omitempty" name:"ZoneInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZonesResponseParams `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachCcnRequestParams struct {
	// 云联网实例ID。
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

type DetachCcnRequest struct {
	*tchttp.BaseRequest
	
	// 云联网实例ID。
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *DetachCcnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachCcnRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachCcnRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachCcnResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetachCcnResponse struct {
	*tchttp.BaseResponse
	Response *DetachCcnResponseParams `json:"Response"`
}

func (r *DetachCcnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachCcnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachDisksRequestParams struct {
	// 云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

type DetachDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

func (r *DetachDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachDisksResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetachDisksResponse struct {
	*tchttp.BaseResponse
	Response *DetachDisksResponseParams `json:"Response"`
}

func (r *DetachDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailPrice struct {
	// 描述计费项目名称，目前取值
	// <li>"DiskSpace"代表云硬盘空间收费项。</li>
	// <li>"DiskBackupQuota"代表云硬盘备份点配额收费项。</li>
	PriceName *string `json:"PriceName,omitempty" name:"PriceName"`

	// 云硬盘计费项维度单价。
	OriginUnitPrice *float64 `json:"OriginUnitPrice,omitempty" name:"OriginUnitPrice"`

	// 云硬盘计费项维度总价。
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 云硬盘在计费项维度折扣。
	Discount *float64 `json:"Discount,omitempty" name:"Discount"`

	// 云硬盘在计费项维度折后总价。
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
}

// Predefined struct for user
type DisassociateInstancesKeyPairsRequestParams struct {
	// 密钥对 ID 列表。每次请求批量密钥对的上限为 100。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DisassociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// 密钥对 ID 列表。每次请求批量密钥对的上限为 100。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DisassociateInstancesKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateInstancesKeyPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateInstancesKeyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateInstancesKeyPairsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisassociateInstancesKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateInstancesKeyPairsResponseParams `json:"Response"`
}

func (r *DisassociateInstancesKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateInstancesKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiscountDetail struct {
	// 计费时长。
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 计费单元。
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 总价。
	TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`

	// 折后总价。
	RealTotalCost *float64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 折扣。
	Discount *float64 `json:"Discount,omitempty" name:"Discount"`

	// 具体折扣详情。
	PolicyDetail *PolicyDetail `json:"PolicyDetail,omitempty" name:"PolicyDetail"`
}

type Disk struct {
	// 磁盘ID
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 磁盘名称
	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`

	// 磁盘类型
	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`

	// 磁盘介质类型
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 磁盘付费类型
	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`

	// 磁盘大小
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 续费标识
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 磁盘状态，取值范围：
	// <li>PENDING：创建中。 </li>
	// <li>UNATTACHED：未挂载。</li>
	// <li>ATTACHING：挂载中。</li>
	// <li>ATTACHED：已挂载。</li>
	// <li>DETACHING：卸载中。 </li>
	// <li> SHUTDOWN：已隔离。</li>
	// <li> CREATED_FAILED：创建失败。</li>
	// <li>TERMINATING：销毁中。</li>
	// <li> DELETING：删除中。</li>
	// <li> FREEZING：冻结中。</li>
	DiskState *string `json:"DiskState,omitempty" name:"DiskState"`

	// 磁盘挂载状态
	Attached *bool `json:"Attached,omitempty" name:"Attached"`

	// 是否随实例释放
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`

	// 上一次操作
	LatestOperation *string `json:"LatestOperation,omitempty" name:"LatestOperation"`

	// 上一次操作状态
	LatestOperationState *string `json:"LatestOperationState,omitempty" name:"LatestOperationState"`

	// 上一次请求ID
	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitempty" name:"LatestOperationRequestId"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 到期时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 隔离时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolatedTime *string `json:"IsolatedTime,omitempty" name:"IsolatedTime"`

	// 云硬盘的已有备份点数量。
	DiskBackupCount *int64 `json:"DiskBackupCount,omitempty" name:"DiskBackupCount"`

	// 云硬盘的备份点配额数量。
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitempty" name:"DiskBackupQuota"`
}

type DiskBackup struct {
	// 云硬盘备份点ID。
	DiskBackupId *string `json:"DiskBackupId,omitempty" name:"DiskBackupId"`

	// 创建此云硬盘备份点的云硬盘类型。取值：<li>DATA_DISK：数据盘</li>
	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`

	// 创建此云硬盘备份点的云硬盘 ID。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 创建此云硬盘备份点的云硬盘大小，单位 GB。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 云硬盘备份点名称，用户自定义的云硬盘备份点别名。
	DiskBackupName *string `json:"DiskBackupName,omitempty" name:"DiskBackupName"`

	// 云硬盘备份点的状态。取值范围：
	// <li>NORMAL：正常。 </li>
	// <li>CREATING：创建中。</li>
	// <li>ROLLBACKING：回滚中。</li>
	// <li>DELETING：删除中。</li>
	DiskBackupState *string `json:"DiskBackupState,omitempty" name:"DiskBackupState"`

	// 创建或回滚云硬盘备份点进度百分比，成功后此字段取值为 100。
	Percent *int64 `json:"Percent,omitempty" name:"Percent"`

	// 上一次操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperation *string `json:"LatestOperation,omitempty" name:"LatestOperation"`

	// 上一次操作状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperationState *string `json:"LatestOperationState,omitempty" name:"LatestOperationState"`

	// 上一次请求ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitempty" name:"LatestOperationRequestId"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type DiskBackupDeniedActions struct {
	// 云硬盘备份点ID。
	DiskBackupId *string `json:"DiskBackupId,omitempty" name:"DiskBackupId"`

	// 操作限制列表。
	DeniedActions []*DeniedAction `json:"DeniedActions,omitempty" name:"DeniedActions"`
}

type DiskChargePrepaid struct {
	// 新购周期。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 续费标识。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 新购单位. 默认值: "m"。
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
}

type DiskConfig struct {
	// 可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 云硬盘类型。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 云硬盘可售卖状态。
	DiskSalesState *string `json:"DiskSalesState,omitempty" name:"DiskSalesState"`

	// 最大云硬盘大小。
	MaxDiskSize *int64 `json:"MaxDiskSize,omitempty" name:"MaxDiskSize"`

	// 最小云硬盘大小。
	MinDiskSize *int64 `json:"MinDiskSize,omitempty" name:"MinDiskSize"`

	// 云硬盘步长。
	DiskStepSize *int64 `json:"DiskStepSize,omitempty" name:"DiskStepSize"`
}

type DiskDeniedActions struct {
	// 云硬盘ID。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 操作限制列表。
	DeniedActions []*DeniedAction `json:"DeniedActions,omitempty" name:"DeniedActions"`
}

type DiskPrice struct {
	// 云硬盘单价。
	OriginalDiskPrice *float64 `json:"OriginalDiskPrice,omitempty" name:"OriginalDiskPrice"`

	// 云硬盘总价。
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 折扣。
	Discount *float64 `json:"Discount,omitempty" name:"Discount"`

	// 折后总价。
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

	// 计费项目明细列表。
	DetailPrices []*DetailPrice `json:"DetailPrices,omitempty" name:"DetailPrices"`
}

type DiskReturnable struct {
	// 云硬盘ID。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 云硬盘是否可退还。
	IsReturnable *bool `json:"IsReturnable,omitempty" name:"IsReturnable"`

	// 云硬盘退还失败错误码。
	ReturnFailCode *int64 `json:"ReturnFailCode,omitempty" name:"ReturnFailCode"`

	// 云硬盘退还失败错误信息。
	ReturnFailMessage *string `json:"ReturnFailMessage,omitempty" name:"ReturnFailMessage"`
}

type DockerContainerConfiguration struct {
	// 容器镜像地址
	ContainerImage *string `json:"ContainerImage,omitempty" name:"ContainerImage"`

	// 容器名称
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 环境变量列表
	Envs []*ContainerEnv `json:"Envs,omitempty" name:"Envs"`

	// 容器端口主机端口映射列表
	PublishPorts []*DockerContainerPublishPort `json:"PublishPorts,omitempty" name:"PublishPorts"`

	// 容器加载本地卷列表
	Volumes []*DockerContainerVolume `json:"Volumes,omitempty" name:"Volumes"`

	// 运行的命令
	Command *string `json:"Command,omitempty" name:"Command"`
}

type DockerContainerPublishPort struct {
	// 主机端口
	HostPort *int64 `json:"HostPort,omitempty" name:"HostPort"`

	// 容器端口
	ContainerPort *int64 `json:"ContainerPort,omitempty" name:"ContainerPort"`

	// 对外绑定IP，默认0.0.0.0
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 协议，默认tcp，支持tcp/udp/sctp
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type DockerContainerVolume struct {
	// 容器路径
	ContainerPath *string `json:"ContainerPath,omitempty" name:"ContainerPath"`

	// 主机路径
	HostPath *string `json:"HostPath,omitempty" name:"HostPath"`
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`
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

type GeneralResourceQuota struct {
	// 资源名称。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 资源当前可用数量。
	ResourceQuotaAvailable *int64 `json:"ResourceQuotaAvailable,omitempty" name:"ResourceQuotaAvailable"`

	// 资源总数量。
	ResourceQuotaTotal *int64 `json:"ResourceQuotaTotal,omitempty" name:"ResourceQuotaTotal"`
}

// Predefined struct for user
type ImportKeyPairRequestParams struct {
	// 密钥对名称，可由数字，字母和下划线组成，长度不超过 25 个字符。
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// 密钥对的公钥内容， OpenSSH RSA 格式。
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
}

type ImportKeyPairRequest struct {
	*tchttp.BaseRequest
	
	// 密钥对名称，可由数字，字母和下划线组成，长度不超过 25 个字符。
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// 密钥对的公钥内容， OpenSSH RSA 格式。
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
}

func (r *ImportKeyPairRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportKeyPairRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyName")
	delete(f, "PublicKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportKeyPairRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportKeyPairResponseParams struct {
	// 密钥对 ID。
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImportKeyPairResponse struct {
	*tchttp.BaseResponse
	Response *ImportKeyPairResponseParams `json:"Response"`
}

func (r *ImportKeyPairResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportKeyPairResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateBlueprintRequestParams struct {
	// 自定义镜像的个数。默认值为1。
	BlueprintCount *int64 `json:"BlueprintCount,omitempty" name:"BlueprintCount"`
}

type InquirePriceCreateBlueprintRequest struct {
	*tchttp.BaseRequest
	
	// 自定义镜像的个数。默认值为1。
	BlueprintCount *int64 `json:"BlueprintCount,omitempty" name:"BlueprintCount"`
}

func (r *InquirePriceCreateBlueprintRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateBlueprintRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BlueprintCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceCreateBlueprintRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateBlueprintResponseParams struct {
	// 自定义镜像的价格参数。
	BlueprintPrice *BlueprintPrice `json:"BlueprintPrice,omitempty" name:"BlueprintPrice"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquirePriceCreateBlueprintResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceCreateBlueprintResponseParams `json:"Response"`
}

func (r *InquirePriceCreateBlueprintResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateBlueprintResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateDisksRequestParams struct {
	// 云硬盘大小, 单位: GB。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 云硬盘介质类型。取值: "CLOUD_PREMIUM"(高性能云盘), "CLOUD_SSD"(SSD云硬盘)。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 新购云硬盘包年包月相关参数设置。
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitempty" name:"DiskChargePrepaid"`

	// 云硬盘个数, 默认值: 1。
	DiskCount *int64 `json:"DiskCount,omitempty" name:"DiskCount"`

	// 指定云硬盘备份点配额，不传时默认为不带备份点配额。目前只支持不带或设置1个云硬盘备份点配额。
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitempty" name:"DiskBackupQuota"`
}

type InquirePriceCreateDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘大小, 单位: GB。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 云硬盘介质类型。取值: "CLOUD_PREMIUM"(高性能云盘), "CLOUD_SSD"(SSD云硬盘)。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 新购云硬盘包年包月相关参数设置。
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitempty" name:"DiskChargePrepaid"`

	// 云硬盘个数, 默认值: 1。
	DiskCount *int64 `json:"DiskCount,omitempty" name:"DiskCount"`

	// 指定云硬盘备份点配额，不传时默认为不带备份点配额。目前只支持不带或设置1个云硬盘备份点配额。
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitempty" name:"DiskBackupQuota"`
}

func (r *InquirePriceCreateDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskSize")
	delete(f, "DiskType")
	delete(f, "DiskChargePrepaid")
	delete(f, "DiskCount")
	delete(f, "DiskBackupQuota")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceCreateDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateDisksResponseParams struct {
	// 云硬盘价格。
	DiskPrice *DiskPrice `json:"DiskPrice,omitempty" name:"DiskPrice"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquirePriceCreateDisksResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceCreateDisksResponseParams `json:"Response"`
}

func (r *InquirePriceCreateDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateInstancesRequestParams struct {
	// 实例的套餐 ID。
	BundleId *string `json:"BundleId,omitempty" name:"BundleId"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 创建数量，默认为 1。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 应用镜像 ID，使用收费应用镜像时必填。可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值中的BlueprintId获取。
	BlueprintId *string `json:"BlueprintId,omitempty" name:"BlueprintId"`
}

type InquirePriceCreateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例的套餐 ID。
	BundleId *string `json:"BundleId,omitempty" name:"BundleId"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 创建数量，默认为 1。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 应用镜像 ID，使用收费应用镜像时必填。可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值中的BlueprintId获取。
	BlueprintId *string `json:"BlueprintId,omitempty" name:"BlueprintId"`
}

func (r *InquirePriceCreateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BundleId")
	delete(f, "InstanceChargePrepaid")
	delete(f, "InstanceCount")
	delete(f, "BlueprintId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceCreateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateInstancesResponseParams struct {
	// 询价信息。
	Price *Price `json:"Price,omitempty" name:"Price"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquirePriceCreateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceCreateInstancesResponseParams `json:"Response"`
}

func (r *InquirePriceCreateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewDisksRequestParams struct {
	// 云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// 续费云硬盘包年包月相关参数设置。
	RenewDiskChargePrepaid *RenewDiskChargePrepaid `json:"RenewDiskChargePrepaid,omitempty" name:"RenewDiskChargePrepaid"`
}

type InquirePriceRenewDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// 续费云硬盘包年包月相关参数设置。
	RenewDiskChargePrepaid *RenewDiskChargePrepaid `json:"RenewDiskChargePrepaid,omitempty" name:"RenewDiskChargePrepaid"`
}

func (r *InquirePriceRenewDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "RenewDiskChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceRenewDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewDisksResponseParams struct {
	// 云硬盘价格。
	DiskPrice *DiskPrice `json:"DiskPrice,omitempty" name:"DiskPrice"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquirePriceRenewDisksResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceRenewDisksResponseParams `json:"Response"`
}

func (r *InquirePriceRenewDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewInstancesRequestParams struct {
	// 待续费的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573 )接口返回值中的InstanceId获取。每次请求批量实例的上限为50。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 是否续费数据盘。默认值: false, 即不续费。
	RenewDataDisk *bool `json:"RenewDataDisk,omitempty" name:"RenewDataDisk"`

	// 数据盘是否对齐实例到期时间。默认值: false, 即不对齐。
	AlignInstanceExpiredTime *bool `json:"AlignInstanceExpiredTime,omitempty" name:"AlignInstanceExpiredTime"`
}

type InquirePriceRenewInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 待续费的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573 )接口返回值中的InstanceId获取。每次请求批量实例的上限为50。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 是否续费数据盘。默认值: false, 即不续费。
	RenewDataDisk *bool `json:"RenewDataDisk,omitempty" name:"RenewDataDisk"`

	// 数据盘是否对齐实例到期时间。默认值: false, 即不对齐。
	AlignInstanceExpiredTime *bool `json:"AlignInstanceExpiredTime,omitempty" name:"AlignInstanceExpiredTime"`
}

func (r *InquirePriceRenewInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceChargePrepaid")
	delete(f, "RenewDataDisk")
	delete(f, "AlignInstanceExpiredTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceRenewInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewInstancesResponseParams struct {
	// 询价信息。默认为列表中第一个实例的价格信息。
	Price *Price `json:"Price,omitempty" name:"Price"`

	// 数据盘价格信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataDiskPriceSet []*DataDiskPrice `json:"DataDiskPriceSet,omitempty" name:"DataDiskPriceSet"`

	// 待续费实例价格列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstancePriceDetailSet []*InstancePriceDetail `json:"InstancePriceDetailSet,omitempty" name:"InstancePriceDetailSet"`

	// 总计价格。
	TotalPrice *TotalPrice `json:"TotalPrice,omitempty" name:"TotalPrice"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquirePriceRenewInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceRenewInstancesResponseParams `json:"Response"`
}

func (r *InquirePriceRenewInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
	PrivateAddresses []*string `json:"PrivateAddresses,omitempty" name:"PrivateAddresses"`

	// 实例主网卡的公网 IP。 
	// 注意：此字段可能返回 空，表示取不到有效值。
	PublicAddresses []*string `json:"PublicAddresses,omitempty" name:"PublicAddresses"`

	// 实例带宽信息。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 自动续费标识。取值范围： 
	// NOTIFY_AND_MANUAL_RENEW：表示通知即将过期，但不自动续费  
	// NOTIFY_AND_AUTO_RENEW：表示通知即将过期，而且自动续费 。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 实例登录设置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 实例状态。取值范围： 
	// <li>PENDING：表示创建中</li><li>LAUNCH_FAILED：表示创建失败</li><li>RUNNING：表示运行中</li><li>STOPPED：表示关机</li><li>STARTING：表示开机中</li><li>STOPPING：表示关机中</li><li>REBOOTING：表示重启中</li><li>SHUTDOWN：表示停止待销毁</li><li>TERMINATING：表示销毁中</li><li>DELETING：表示删除中</li><li>FREEZING：表示冻结中</li><li>ENTER_RESCUE_MODE：表示进入救援模式中</li><li>RESCUE_MODE：表示救援模式</li><li>EXIT_RESCUE_MODE：表示退出救援模式中</li>
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

	// 可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例绑定的标签列表。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 实例封禁状态。取值范围：
	// <li>NORMAL实例正常。</li><li>NETWORK_RESTRICT：网络封禁。</li>
	InstanceRestrictState *string `json:"InstanceRestrictState,omitempty" name:"InstanceRestrictState"`
}

type InstanceChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费，用户需要手动续费<br><li>DISABLE_NOTIFY_AND_AUTO_RENEW：不自动续费，且不通知<br><br>默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type InstanceDeniedActions struct {
	// 实例 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 操作限制列表。
	DeniedActions []*DeniedAction `json:"DeniedActions,omitempty" name:"DeniedActions"`
}

type InstancePrice struct {
	// 套餐单价原价。
	OriginalBundlePrice *float64 `json:"OriginalBundlePrice,omitempty" name:"OriginalBundlePrice"`

	// 原价。
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 折扣。
	Discount *float64 `json:"Discount,omitempty" name:"Discount"`

	// 折后价。
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

	// 价格货币单位。取值范围CNY:人民币。USD:美元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Currency *string `json:"Currency,omitempty" name:"Currency"`
}

type InstancePriceDetail struct {
	// 实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 询价信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstancePrice *InstancePrice `json:"InstancePrice,omitempty" name:"InstancePrice"`

	// 折扣梯度详情，每个梯度包含的信息有：时长，折扣数，总价，折扣价，折扣详情（用户折扣、官网折扣、最终折扣）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountDetail []*DiscountDetail `json:"DiscountDetail,omitempty" name:"DiscountDetail"`
}

type InstanceReturnable struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例是否可退还。
	IsReturnable *bool `json:"IsReturnable,omitempty" name:"IsReturnable"`

	// 实例退还失败错误码。
	ReturnFailCode *int64 `json:"ReturnFailCode,omitempty" name:"ReturnFailCode"`

	// 实例退还失败错误信息。
	ReturnFailMessage *string `json:"ReturnFailMessage,omitempty" name:"ReturnFailMessage"`
}

type InstanceTrafficPackage struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 流量包详情列表。
	TrafficPackageSet []*TrafficPackage `json:"TrafficPackageSet,omitempty" name:"TrafficPackageSet"`
}

type InternetAccessible struct {
	// 网络计费类型，取值范围：
	// <li>按流量包付费：TRAFFIC_POSTPAID_BY_HOUR</li>
	// <li>按带宽付费： BANDWIDTH_POSTPAID_BY_HOUR</li>
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// 公网出带宽上限，单位：Mbps。
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// 是否分配公网 IP。
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty" name:"PublicIpAssigned"`
}

// Predefined struct for user
type IsolateInstancesRequestParams struct {
	// 实例ID列表。一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。每次请求退还实例和数据盘数量总计上限为20。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 是否退还挂载的数据盘。取值范围：
	// TRUE：表示退还实例同时退还其挂载的数据盘。
	// FALSE：表示退还实例同时不再退还其挂载的数据盘。
	// 默认取值：TRUE。
	IsolateDataDisk *bool `json:"IsolateDataDisk,omitempty" name:"IsolateDataDisk"`
}

type IsolateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表。一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。每次请求退还实例和数据盘数量总计上限为20。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 是否退还挂载的数据盘。取值范围：
	// TRUE：表示退还实例同时退还其挂载的数据盘。
	// FALSE：表示退还实例同时不再退还其挂载的数据盘。
	// 默认取值：TRUE。
	IsolateDataDisk *bool `json:"IsolateDataDisk,omitempty" name:"IsolateDataDisk"`
}

func (r *IsolateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "IsolateDataDisk")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type IsolateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *IsolateInstancesResponseParams `json:"Response"`
}

func (r *IsolateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyPair struct {
	// 密钥对 ID ，是密钥对的唯一标识。
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 密钥对名称。
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// 密钥对的纯文本公钥。
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// 密钥对关联的实例 ID 列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociatedInstanceIds []*string `json:"AssociatedInstanceIds,omitempty" name:"AssociatedInstanceIds"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 密钥对私钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
}

type LoginConfiguration struct {
	// <li>"YES"代表选择自动生成密码，这时不指定Password字段。</li>
	// <li>"NO"代表选择自定义密码，这时要指定Password字段。</li>
	AutoGeneratePassword *string `json:"AutoGeneratePassword,omitempty" name:"AutoGeneratePassword"`

	// 实例登录密码。具体按照操作系统的复杂度要求。
	// WINDOWS 实例密码必须 12-30 位，不能以“/”开头且不包括用户名，至少包含以下字符中的三种不同字符
	// <li>小写字母：[a-z]</li>
	// <li>大写字母：[A-Z]</li>
	// <li>数字： 0-9</li>
	// <li>特殊字符：()`~!@#$%^&*-+=_|{}[]:;' <>,.?/</li>
	Password *string `json:"Password,omitempty" name:"Password"`
}

type LoginSettings struct {
	// 密钥 ID 列表。关联密钥后，就可以通过对应的私钥来访问实例。注意：此字段可能返回 []，表示取不到有效值。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

// Predefined struct for user
type ModifyBlueprintAttributeRequestParams struct {
	// 镜像 ID。可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值中的BlueprintId获取。
	BlueprintId *string `json:"BlueprintId,omitempty" name:"BlueprintId"`

	// 设置新的镜像名称。最大长度60。
	BlueprintName *string `json:"BlueprintName,omitempty" name:"BlueprintName"`

	// 设置新的镜像描述。最大长度60。
	Description *string `json:"Description,omitempty" name:"Description"`
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

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBlueprintAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBlueprintAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBlueprintAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBlueprintAttributeResponseParams `json:"Response"`
}

func (r *ModifyBlueprintAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlueprintAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBundle struct {
	// 更改实例套餐后需要补的差价。
	ModifyPrice *Price `json:"ModifyPrice,omitempty" name:"ModifyPrice"`

	// 变更套餐状态。取值：
	// <li>SOLD_OUT：套餐售罄</li>
	// <li>AVAILABLE：支持套餐变更</li>
	// <li>UNAVAILABLE：暂不支持套餐变更</li>
	ModifyBundleState *string `json:"ModifyBundleState,omitempty" name:"ModifyBundleState"`

	// 套餐信息。
	Bundle *Bundle `json:"Bundle,omitempty" name:"Bundle"`

	// 不支持套餐变更原因信息。变更套餐状态为"AVAILABLE"时, 该信息为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotSupportModifyMessage *string `json:"NotSupportModifyMessage,omitempty" name:"NotSupportModifyMessage"`
}

// Predefined struct for user
type ModifyDiskBackupsAttributeRequestParams struct {
	// 云硬盘备份点ID列表。
	DiskBackupIds []*string `json:"DiskBackupIds,omitempty" name:"DiskBackupIds"`

	// 云硬盘备份点名称，最大长度90。
	DiskBackupName *string `json:"DiskBackupName,omitempty" name:"DiskBackupName"`
}

type ModifyDiskBackupsAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘备份点ID列表。
	DiskBackupIds []*string `json:"DiskBackupIds,omitempty" name:"DiskBackupIds"`

	// 云硬盘备份点名称，最大长度90。
	DiskBackupName *string `json:"DiskBackupName,omitempty" name:"DiskBackupName"`
}

func (r *ModifyDiskBackupsAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskBackupsAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskBackupIds")
	delete(f, "DiskBackupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDiskBackupsAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiskBackupsAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDiskBackupsAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDiskBackupsAttributeResponseParams `json:"Response"`
}

func (r *ModifyDiskBackupsAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskBackupsAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDisksAttributeRequestParams struct {
	// 云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// 云硬盘名称。
	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
}

type ModifyDisksAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// 云硬盘名称。
	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
}

func (r *ModifyDisksAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDisksAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "DiskName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDisksAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDisksAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDisksAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDisksAttributeResponseParams `json:"Response"`
}

func (r *ModifyDisksAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDisksAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDisksRenewFlagRequestParams struct {
	// 云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// 续费标识。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type ModifyDisksRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// 续费标识。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

func (r *ModifyDisksRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDisksRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "RenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDisksRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDisksRenewFlagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDisksRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDisksRenewFlagResponseParams `json:"Response"`
}

func (r *ModifyDisksRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDisksRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFirewallRuleDescriptionRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 防火墙规则。
	FirewallRule *FirewallRule `json:"FirewallRule,omitempty" name:"FirewallRule"`

	// 防火墙当前版本。用户每次更新防火墙规则时版本会自动加1，防止规则已过期，不填不考虑冲突。
	FirewallVersion *uint64 `json:"FirewallVersion,omitempty" name:"FirewallVersion"`
}

type ModifyFirewallRuleDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 防火墙规则。
	FirewallRule *FirewallRule `json:"FirewallRule,omitempty" name:"FirewallRule"`

	// 防火墙当前版本。用户每次更新防火墙规则时版本会自动加1，防止规则已过期，不填不考虑冲突。
	FirewallVersion *uint64 `json:"FirewallVersion,omitempty" name:"FirewallVersion"`
}

func (r *ModifyFirewallRuleDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFirewallRuleDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FirewallRule")
	delete(f, "FirewallVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFirewallRuleDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFirewallRuleDescriptionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyFirewallRuleDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFirewallRuleDescriptionResponseParams `json:"Response"`
}

func (r *ModifyFirewallRuleDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFirewallRuleDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFirewallRulesRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 防火墙规则列表。
	FirewallRules []*FirewallRule `json:"FirewallRules,omitempty" name:"FirewallRules"`

	// 防火墙当前版本。用户每次更新防火墙规则时版本会自动加1，防止规则已过期，不填不考虑冲突。
	FirewallVersion *uint64 `json:"FirewallVersion,omitempty" name:"FirewallVersion"`
}

type ModifyFirewallRulesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 防火墙规则列表。
	FirewallRules []*FirewallRule `json:"FirewallRules,omitempty" name:"FirewallRules"`

	// 防火墙当前版本。用户每次更新防火墙规则时版本会自动加1，防止规则已过期，不填不考虑冲突。
	FirewallVersion *uint64 `json:"FirewallVersion,omitempty" name:"FirewallVersion"`
}

func (r *ModifyFirewallRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFirewallRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FirewallRules")
	delete(f, "FirewallVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFirewallRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFirewallRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyFirewallRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFirewallRulesResponseParams `json:"Response"`
}

func (r *ModifyFirewallRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFirewallRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesAttributeRequestParams struct {
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 实例名称。可任意命名，但不得超过 60 个字符。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type ModifyInstancesAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 实例名称。可任意命名，但不得超过 60 个字符。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyInstancesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyInstancesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesAttributeResponseParams `json:"Response"`
}

func (r *ModifyInstancesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesBundleRequestParams struct {
	// 实例ID列表。一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。每次请求批量实例的上限为30。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 待变更的套餐Id。可通过[DescribeBundles](https://cloud.tencent.com/document/api/1207/47575)接口返回值中的BundleId获取。
	BundleId *string `json:"BundleId,omitempty" name:"BundleId"`

	// 是否自动抵扣代金券。取值范围：
	// true：表示自动抵扣代金券
	// false：表示不自动抵扣代金券
	// 默认取值：false。
	AutoVoucher *bool `json:"AutoVoucher,omitempty" name:"AutoVoucher"`
}

type ModifyInstancesBundleRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表。一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。每次请求批量实例的上限为30。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 待变更的套餐Id。可通过[DescribeBundles](https://cloud.tencent.com/document/api/1207/47575)接口返回值中的BundleId获取。
	BundleId *string `json:"BundleId,omitempty" name:"BundleId"`

	// 是否自动抵扣代金券。取值范围：
	// true：表示自动抵扣代金券
	// false：表示不自动抵扣代金券
	// 默认取值：false。
	AutoVoucher *bool `json:"AutoVoucher,omitempty" name:"AutoVoucher"`
}

func (r *ModifyInstancesBundleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesBundleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "BundleId")
	delete(f, "AutoVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesBundleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesBundleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyInstancesBundleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesBundleResponseParams `json:"Response"`
}

func (r *ModifyInstancesBundleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesBundleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesLoginKeyPairAttributeRequestParams struct {
	// 实例 ID 列表。每次请求批量实例的上限为 100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 是否允许使用默认密钥对登录，YES：允许登录；NO：禁止登录
	PermitLogin *string `json:"PermitLogin,omitempty" name:"PermitLogin"`
}

type ModifyInstancesLoginKeyPairAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 是否允许使用默认密钥对登录，YES：允许登录；NO：禁止登录
	PermitLogin *string `json:"PermitLogin,omitempty" name:"PermitLogin"`
}

func (r *ModifyInstancesLoginKeyPairAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesLoginKeyPairAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "PermitLogin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesLoginKeyPairAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesLoginKeyPairAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyInstancesLoginKeyPairAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesLoginKeyPairAttributeResponseParams `json:"Response"`
}

func (r *ModifyInstancesLoginKeyPairAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesLoginKeyPairAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesRenewFlagRequestParams struct {
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type ModifyInstancesRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

func (r *ModifyInstancesRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "RenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesRenewFlagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyInstancesRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesRenewFlagResponseParams `json:"Response"`
}

func (r *ModifyInstancesRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotAttributeRequestParams struct {
	// 快照 ID, 可通过 DescribeSnapshots 查询。
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 新的快照名称，最长为 60 个字符。
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotId")
	delete(f, "SnapshotName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySnapshotAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySnapshotAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifySnapshotAttributeResponseParams `json:"Response"`
}

func (r *ModifySnapshotAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PolicyDetail struct {
	// 用户折扣。
	UserDiscount *int64 `json:"UserDiscount,omitempty" name:"UserDiscount"`

	// 公共折扣。
	CommonDiscount *int64 `json:"CommonDiscount,omitempty" name:"CommonDiscount"`

	// 最终折扣。
	FinalDiscount *int64 `json:"FinalDiscount,omitempty" name:"FinalDiscount"`

	// 活动折扣。取值为null，表示无有效值，即没有折扣。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityDiscount *float64 `json:"ActivityDiscount,omitempty" name:"ActivityDiscount"`

	// 折扣类型。
	// user：用户折扣; common：官网折扣; activity：活动折扣。 取值为null，表示无有效值，即没有折扣。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountType *string `json:"DiscountType,omitempty" name:"DiscountType"`
}

type Price struct {
	// 实例价格。
	InstancePrice *InstancePrice `json:"InstancePrice,omitempty" name:"InstancePrice"`
}

// Predefined struct for user
type RebootInstancesRequestParams struct {
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *RebootInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebootInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RebootInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RebootInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RebootInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RebootInstancesResponseParams `json:"Response"`
}

func (r *RebootInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebootInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// 地域名称，例如，ap-guangzhou。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域描述，例如，华南地区(广州)。
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地域是否可用状态，取值仅为AVAILABLE。
	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`

	// 是否中国大陆地域
	IsChinaMainland *bool `json:"IsChinaMainland,omitempty" name:"IsChinaMainland"`
}

type RenewDiskChargePrepaid struct {
	// 新购周期。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 续费标识。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 周期单位. 默认值: "m"。
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 当前实例到期时间。
	CurInstanceDeadline *string `json:"CurInstanceDeadline,omitempty" name:"CurInstanceDeadline"`
}

// Predefined struct for user
type RenewInstancesRequestParams struct {
	// 实例ID列表。一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 是否续费弹性数据盘。取值范围：
	// TRUE：表示续费实例同时续费其挂载的数据盘
	// FALSE：表示续费实例同时不再续费其挂载的数据盘
	// 默认取值：TRUE。
	RenewDataDisk *bool `json:"RenewDataDisk,omitempty" name:"RenewDataDisk"`

	// 是否自动抵扣代金券。取值范围：
	// TRUE：表示自动抵扣代金券
	// FALSE：表示不自动抵扣代金券
	// 默认取值：FALSE。
	AutoVoucher *bool `json:"AutoVoucher,omitempty" name:"AutoVoucher"`
}

type RenewInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表。一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 是否续费弹性数据盘。取值范围：
	// TRUE：表示续费实例同时续费其挂载的数据盘
	// FALSE：表示续费实例同时不再续费其挂载的数据盘
	// 默认取值：TRUE。
	RenewDataDisk *bool `json:"RenewDataDisk,omitempty" name:"RenewDataDisk"`

	// 是否自动抵扣代金券。取值范围：
	// TRUE：表示自动抵扣代金券
	// FALSE：表示不自动抵扣代金券
	// 默认取值：FALSE。
	AutoVoucher *bool `json:"AutoVoucher,omitempty" name:"AutoVoucher"`
}

func (r *RenewInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceChargePrepaid")
	delete(f, "RenewDataDisk")
	delete(f, "AutoVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RenewInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RenewInstancesResponseParams `json:"Response"`
}

func (r *RenewInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAttachCcnRequestParams struct {
	// 云联网实例ID。
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

type ResetAttachCcnRequest struct {
	*tchttp.BaseRequest
	
	// 云联网实例ID。
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *ResetAttachCcnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAttachCcnRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAttachCcnRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAttachCcnResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetAttachCcnResponse struct {
	*tchttp.BaseResponse
	Response *ResetAttachCcnResponseParams `json:"Response"`
}

func (r *ResetAttachCcnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAttachCcnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetInstanceBlueprint struct {
	// 镜像详细信息
	BlueprintInfo *Blueprint `json:"BlueprintInfo,omitempty" name:"BlueprintInfo"`

	// 实例镜像是否可重置为目标镜像
	IsResettable *bool `json:"IsResettable,omitempty" name:"IsResettable"`

	// 不可重置信息.当镜像可重置时为""
	NonResettableMessage *string `json:"NonResettableMessage,omitempty" name:"NonResettableMessage"`
}

// Predefined struct for user
type ResetInstanceRequestParams struct {
	// 实例 ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 镜像 ID。可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值中的BlueprintId获取。
	BlueprintId *string `json:"BlueprintId,omitempty" name:"BlueprintId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BlueprintId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ResetInstanceResponseParams `json:"Response"`
}

func (r *ResetInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesPasswordRequestParams struct {
	// 实例 ID 列表。每次请求批量实例的上限为 100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：
	// `LINUX_UNIX` 实例密码必须 8-30 位，推荐使用 12 位以上密码，不能以“/”开头，至少包含以下字符中的三种不同字符，字符种类：<br><li>小写字母：[a-z]<br><li>大写字母：[A-Z]<br><li>数字：0-9<br><li>特殊字符： ()\`~!@#$%^&\*-+=\_|{}[]:;'<>,.?/</li>
	// `WINDOWS` 实例密码必须 12-30 位，不能以“/”开头且不包括用户名，至少包含以下字符中的三种不同字符<br><li>小写字母：[a-z]<br><li>大写字母：[A-Z]<br><li>数字： 0-9<br><li>特殊字符：()\`~!@#$%^&\*-+=\_|{}[]:;' <>,.?/<br><li>如果实例即包含 `LINUX_UNIX` 实例又包含 `WINDOWS` 实例，则密码复杂度限制按照 `WINDOWS` 实例的限制。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 待重置密码的实例操作系统用户名。不得超过 64 个字符。
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type ResetInstancesPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：
	// `LINUX_UNIX` 实例密码必须 8-30 位，推荐使用 12 位以上密码，不能以“/”开头，至少包含以下字符中的三种不同字符，字符种类：<br><li>小写字母：[a-z]<br><li>大写字母：[A-Z]<br><li>数字：0-9<br><li>特殊字符： ()\`~!@#$%^&\*-+=\_|{}[]:;'<>,.?/</li>
	// `WINDOWS` 实例密码必须 12-30 位，不能以“/”开头且不包括用户名，至少包含以下字符中的三种不同字符<br><li>小写字母：[a-z]<br><li>大写字母：[A-Z]<br><li>数字： 0-9<br><li>特殊字符：()\`~!@#$%^&\*-+=\_|{}[]:;' <>,.?/<br><li>如果实例即包含 `LINUX_UNIX` 实例又包含 `WINDOWS` 实例，则密码复杂度限制按照 `WINDOWS` 实例的限制。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 待重置密码的实例操作系统用户名。不得超过 64 个字符。
	UserName *string `json:"UserName,omitempty" name:"UserName"`
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
	delete(f, "UserName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstancesPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesPasswordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type Scene struct {
	// 使用场景Id
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`

	// 使用场景展示名称
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 使用场景描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type SceneInfo struct {
	// 使用场景Id。
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`

	// 使用场景展示名称。
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 使用场景描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type SnapshotDeniedActions struct {
	// 快照 ID。
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 操作限制列表。
	DeniedActions []*DeniedAction `json:"DeniedActions,omitempty" name:"DeniedActions"`
}

type Software struct {
	// 软件名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 软件版本。
	Version *string `json:"Version,omitempty" name:"Version"`

	// 软件图片 URL。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 软件安装目录。
	InstallDir *string `json:"InstallDir,omitempty" name:"InstallDir"`

	// 软件详情列表。
	DetailSet []*SoftwareDetail `json:"DetailSet,omitempty" name:"DetailSet"`
}

type SoftwareDetail struct {
	// 详情唯一键。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 详情标题。
	Title *string `json:"Title,omitempty" name:"Title"`

	// 详情值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type StartInstancesRequestParams struct {
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type StartInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *StartInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartInstancesResponse struct {
	*tchttp.BaseResponse
	Response *StartInstancesResponseParams `json:"Response"`
}

func (r *StartInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstancesRequestParams struct {
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type StopInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *StopInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopInstancesResponse struct {
	*tchttp.BaseResponse
	Response *StopInstancesResponseParams `json:"Response"`
}

func (r *StopInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type TerminateDisksRequestParams struct {
	// 云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

type TerminateDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

func (r *TerminateDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateDisksResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TerminateDisksResponse struct {
	*tchttp.BaseResponse
	Response *TerminateDisksResponseParams `json:"Response"`
}

func (r *TerminateDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateInstancesRequestParams struct {
	// 实例ID列表。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type TotalPrice struct {
	// 原始总计价格。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 折扣总计价格。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
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

type ZoneInfo struct {
	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区中文名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 实例购买页可用区展示标签
	InstanceDisplayLabel *string `json:"InstanceDisplayLabel,omitempty" name:"InstanceDisplayLabel"`
}