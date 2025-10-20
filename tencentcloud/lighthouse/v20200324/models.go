// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type ApplyDiskBackupRequestParams struct {
	// 云硬盘ID，可通过[DescribeDisks](https://cloud.tencent.com/document/api/1207/66093)接口查询。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 云硬盘备份点ID，可通过[DescribeDiskBackups](https://cloud.tencent.com/document/api/1207/84379)接口查询。
	DiskBackupId *string `json:"DiskBackupId,omitnil,omitempty" name:"DiskBackupId"`
}

type ApplyDiskBackupRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID，可通过[DescribeDisks](https://cloud.tencent.com/document/api/1207/66093)接口查询。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 云硬盘备份点ID，可通过[DescribeDiskBackups](https://cloud.tencent.com/document/api/1207/84379)接口查询。
	DiskBackupId *string `json:"DiskBackupId,omitnil,omitempty" name:"DiskBackupId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ApplyFirewallTemplateRequestParams struct {
	// 防火墙模板ID。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 应用防火墙模板的实例列表。列表长度最大值是100。
	ApplyInstances []*InstanceIdentifier `json:"ApplyInstances,omitnil,omitempty" name:"ApplyInstances"`
}

type ApplyFirewallTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙模板ID。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 应用防火墙模板的实例列表。列表长度最大值是100。
	ApplyInstances []*InstanceIdentifier `json:"ApplyInstances,omitnil,omitempty" name:"ApplyInstances"`
}

func (r *ApplyFirewallTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyFirewallTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "ApplyInstances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyFirewallTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyFirewallTemplateResponseParams struct {
	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyFirewallTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ApplyFirewallTemplateResponseParams `json:"Response"`
}

func (r *ApplyFirewallTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyFirewallTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyInstanceSnapshotRequestParams struct {
	// 实例 ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/product/1207/47573) 接口返回值中的 InstanceId	获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 快照 ID。可通过 [DescribeSnapshots](https://cloud.tencent.com/document/product/1207/54388) 接口返回值中的 SnapshotId		获取。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`
}

type ApplyInstanceSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/product/1207/47573) 接口返回值中的 InstanceId	获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 快照 ID。可通过 [DescribeSnapshots](https://cloud.tencent.com/document/product/1207/54388) 接口返回值中的 SnapshotId		获取。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 密钥对 ID 列表，每次请求批量密钥对的上限为 100。可通过[DescribeKeyPairs](https://cloud.tencent.com/document/api/1207/55540)接口返回值中的KeyId获取。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type AssociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// 密钥对 ID 列表，每次请求批量密钥对的上限为 100。可通过[DescribeKeyPairs](https://cloud.tencent.com/document/api/1207/55540)接口返回值中的KeyId获取。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 云联网实例ID。可通过[DescribeCcns](https://cloud.tencent.com/document/product/215/19199)接口返回值中的CcnId获取。
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`
}

type AttachCcnRequest struct {
	*tchttp.BaseRequest
	
	// 云联网实例ID。可通过[DescribeCcns](https://cloud.tencent.com/document/product/215/19199)接口返回值中的CcnId获取。
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例已挂载弹性云盘数量
	AttachedDiskCount *int64 `json:"AttachedDiskCount,omitnil,omitempty" name:"AttachedDiskCount"`

	// 可挂载弹性云盘数量
	MaxAttachCount *int64 `json:"MaxAttachCount,omitnil,omitempty" name:"MaxAttachCount"`
}

// Predefined struct for user
type AttachDisksRequestParams struct {
	// 云硬盘ID列表。每次批量请求云硬盘的上限为 100。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自动续费标识。取值范围：
	// 
	// NOTIFY_AND_AUTO_RENEW：通知过期且自动续费。 NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费，用户需要手动续费。 DISABLE_NOTIFY_AND_MANUAL_RENEW：不自动续费，且不通知。
	// 
	// 默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，云盘到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type AttachDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。每次批量请求云硬盘的上限为 100。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自动续费标识。取值范围：
	// 
	// NOTIFY_AND_AUTO_RENEW：通知过期且自动续费。 NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费，用户需要手动续费。 DISABLE_NOTIFY_AND_MANUAL_RENEW：不自动续费，且不通知。
	// 
	// 默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，云盘到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type AutoMountConfiguration struct {
	// 待挂载的实例ID。指定的实例必须与指定的数据盘处于同一可用区，实例状态必须处于“运行中”状态，且实例必须支持[自动化助手](https://cloud.tencent.com/document/product/1340/50752)。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例内的挂载点。仅Linux操作系统的实例可传入该参数, 不传则默认挂载在“/data/disk”路径下。
	MountPoint *string `json:"MountPoint,omitnil,omitempty" name:"MountPoint"`

	// 文件系统类型。取值: “ext4”、“xfs”。仅Linux操作系统的实例可传入该参数, 不传则默认为“ext4”。
	FileSystemType *string `json:"FileSystemType,omitnil,omitempty" name:"FileSystemType"`
}

type Blueprint struct {
	// 镜像 ID  ，是 Blueprint 的唯一标识。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`

	// 镜像对外展示标题。
	DisplayTitle *string `json:"DisplayTitle,omitnil,omitempty" name:"DisplayTitle"`

	// 镜像对外展示版本。
	DisplayVersion *string `json:"DisplayVersion,omitnil,omitempty" name:"DisplayVersion"`

	// 镜像描述信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 操作系统名称。
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// 操作系统平台。
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 操作系统平台类型，如 LINUX_UNIX、WINDOWS。
	PlatformType *string `json:"PlatformType,omitnil,omitempty" name:"PlatformType"`

	// 镜像类型，如 APP_OS（应用镜像）, PURE_OS（系统镜像）, DOCKER（容器）, PRIVATE（私有镜像）, SHARED（共享镜像）, GAME_PORTAL（游戏专区镜像）。
	BlueprintType *string `json:"BlueprintType,omitnil,omitempty" name:"BlueprintType"`

	// 镜像图片 URL。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 镜像所需系统盘大小，单位 GB。
	RequiredSystemDiskSize *int64 `json:"RequiredSystemDiskSize,omitnil,omitempty" name:"RequiredSystemDiskSize"`

	// 镜像状态。
	// 可选值：
	// NORMAL（正常）、SYNCING（同步中）、OFFLINE（下线）、ISOLATED（已隔离）、CREATEFAILED（创建失败）、SYNCING_FAILED（目的地域同步失败）、ISOLATING（隔离中）、ISOLATED（已隔离）、DELETING（删除中）、DESTROYING（销毁中）。
	BlueprintState *string `json:"BlueprintState,omitnil,omitempty" name:"BlueprintState"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 镜像名称。
	BlueprintName *string `json:"BlueprintName,omitnil,omitempty" name:"BlueprintName"`

	// 镜像是否支持自动化助手。
	SupportAutomationTools *bool `json:"SupportAutomationTools,omitnil,omitempty" name:"SupportAutomationTools"`

	// 镜像所需内存大小, 单位: GB
	RequiredMemorySize *int64 `json:"RequiredMemorySize,omitnil,omitempty" name:"RequiredMemorySize"`

	// CVM镜像共享到轻量应用服务器轻量应用服务器后的CVM镜像ID。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 官方网站Url。
	CommunityUrl *string `json:"CommunityUrl,omitnil,omitempty" name:"CommunityUrl"`

	// 指导文章Url。
	GuideUrl *string `json:"GuideUrl,omitnil,omitempty" name:"GuideUrl"`

	// 镜像关联使用场景Id列表。
	SceneIdSet []*string `json:"SceneIdSet,omitnil,omitempty" name:"SceneIdSet"`

	// Docker版本号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DockerVersion *string `json:"DockerVersion,omitnil,omitempty" name:"DockerVersion"`

	// 镜像是否已共享。
	BlueprintShared *bool `json:"BlueprintShared,omitnil,omitempty" name:"BlueprintShared"`

	// 镜像绑定的标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type BlueprintInstance struct {
	// 镜像信息。
	Blueprint *Blueprint `json:"Blueprint,omitnil,omitempty" name:"Blueprint"`

	// 软件列表。
	SoftwareSet []*Software `json:"SoftwareSet,omitnil,omitempty" name:"SoftwareSet"`

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type BlueprintPrice struct {
	// 镜像单价，原价。单位元。
	OriginalBlueprintPrice *float64 `json:"OriginalBlueprintPrice,omitnil,omitempty" name:"OriginalBlueprintPrice"`

	// 镜像总价，原价。单位元。
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 折扣。
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 镜像折扣后总价。单位元。
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`
}

type Bundle struct {
	// 套餐 ID。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 内存大小，单位 GB。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 系统盘类型。
	// 取值范围： 
	// <li> CLOUD_SSD：SSD 云硬盘</li><li> CLOUD_PREMIUM：高性能云硬盘</li>
	SystemDiskType *string `json:"SystemDiskType,omitnil,omitempty" name:"SystemDiskType"`

	// 系统盘大小。单位GB。
	SystemDiskSize *int64 `json:"SystemDiskSize,omitnil,omitempty" name:"SystemDiskSize"`

	// 每月网络流量，单位 GB。
	MonthlyTraffic *int64 `json:"MonthlyTraffic,omitnil,omitempty" name:"MonthlyTraffic"`

	// 是否支持 Linux/Unix 平台。
	SupportLinuxUnixPlatform *bool `json:"SupportLinuxUnixPlatform,omitnil,omitempty" name:"SupportLinuxUnixPlatform"`

	// 是否支持 Windows 平台。
	SupportWindowsPlatform *bool `json:"SupportWindowsPlatform,omitnil,omitempty" name:"SupportWindowsPlatform"`

	// 套餐当前单位价格信息。
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// CPU 核数。
	CPU *int64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// 峰值带宽，单位 Mbps。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 网络计费类型。
	InternetChargeType *string `json:"InternetChargeType,omitnil,omitempty" name:"InternetChargeType"`

	// 套餐售卖状态,取值:‘AVAILABLE’(可用) , ‘SOLD_OUT’(售罄)
	BundleSalesState *string `json:"BundleSalesState,omitnil,omitempty" name:"BundleSalesState"`

	// 套餐类型。
	// 取值范围：
	// <li>GENERAL_BUNDLE：通用型</li>
	// <li>STORAGE_BUNDLE：存储型</li>
	// <li>ENTERPRISE_BUNDLE：企业型</li>
	// <li>EXCLUSIVE_BUNDLE：专属型</li>
	// <li>BEFAST_BUNDLE：蜂驰型 </li>
	// <li>STARTER_BUNDLE：入门型</li>
	// <li>CAREFREE_BUNDLE：无忧型</li>
	// <li>RAZOR_SPEED_BUNDLE：锐驰型</li>
	BundleType *string `json:"BundleType,omitnil,omitempty" name:"BundleType"`

	// 套餐类型描述信息。
	BundleTypeDescription *string `json:"BundleTypeDescription,omitnil,omitempty" name:"BundleTypeDescription"`

	// 套餐展示标签.
	// 取值范围:
	// "ACTIVITY": 活动套餐,
	// "NORMAL": 普通套餐
	// "CAREFREE": 无忧套餐
	BundleDisplayLabel *string `json:"BundleDisplayLabel,omitnil,omitempty" name:"BundleDisplayLabel"`

	// 流量是否无上限。
	TrafficUnlimited *bool `json:"TrafficUnlimited,omitnil,omitempty" name:"TrafficUnlimited"`
}

// Predefined struct for user
type CancelShareBlueprintAcrossAccountsRequestParams struct {
	// 镜像ID, 可以通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回的BlueprintId获取。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`

	// 接收共享镜像的[账号ID](https://cloud.tencent.com/document/product/213/4944#.E8.8E.B7.E5.8F.96.E4.B8.BB.E8.B4.A6.E5.8F.B7.E7.9A.84.E8.B4.A6.E5.8F.B7-id)列表。账号ID不同于QQ号，查询用户账号ID请查看账号信息中的账号ID栏。账号个数取值最大为10。
	AccountIds []*string `json:"AccountIds,omitnil,omitempty" name:"AccountIds"`
}

type CancelShareBlueprintAcrossAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 镜像ID, 可以通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回的BlueprintId获取。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`

	// 接收共享镜像的[账号ID](https://cloud.tencent.com/document/product/213/4944#.E8.8E.B7.E5.8F.96.E4.B8.BB.E8.B4.A6.E5.8F.B7.E7.9A.84.E8.B4.A6.E5.8F.B7-id)列表。账号ID不同于QQ号，查询用户账号ID请查看账号信息中的账号ID栏。账号个数取值最大为10。
	AccountIds []*string `json:"AccountIds,omitnil,omitempty" name:"AccountIds"`
}

func (r *CancelShareBlueprintAcrossAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelShareBlueprintAcrossAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BlueprintId")
	delete(f, "AccountIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelShareBlueprintAcrossAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelShareBlueprintAcrossAccountsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelShareBlueprintAcrossAccountsResponse struct {
	*tchttp.BaseResponse
	Response *CancelShareBlueprintAcrossAccountsResponseParams `json:"Response"`
}

func (r *CancelShareBlueprintAcrossAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelShareBlueprintAcrossAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnAttachedInstance struct {
	// 云联网ID。
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// 关联实例CIDR。
	CidrBlock []*string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

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
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 关联时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachedTime *string `json:"AttachedTime,omitnil,omitempty" name:"AttachedTime"`

	// 备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type Command struct {
	// Base64编码后的命令内容，长度不可超过64KB。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 命令超时时间，默认60秒。取值范围[1, 86400]。
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 命令执行路径，对于 SHELL 命令默认为 /root，对于 POWERSHELL 命令默认为 C:\Program Files\qcloud\tat_agent\workdir。
	WorkingDirectory *string `json:"WorkingDirectory,omitnil,omitempty" name:"WorkingDirectory"`

	// 在 Lighthouse 实例中执行命令的用户名称。
	// 默认情况下，在 Linux 实例中以 root 用户执行命令；在Windows 实例中以 System 用户执行命令。
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`
}

type ContainerEnv struct {
	// 环境变量Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 环境变量值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type CreateBlueprintRequestParams struct {
	// 镜像名称。最大长度60。
	BlueprintName *string `json:"BlueprintName,omitnil,omitempty" name:"BlueprintName"`

	// 镜像描述。最大长度60。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 需要制作镜像的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/1207/47573) 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否执行强制关机以制作镜像。
	// 取值范围：
	// True：表示关机之后制作镜像
	// False：表示开机状态制作镜像
	// 默认取值：True
	// 开机状态制作镜像，可能导致部分数据未备份，影响数据安全。
	ForcePowerOff *bool `json:"ForcePowerOff,omitnil,omitempty" name:"ForcePowerOff"`

	// 标签键和标签值。 如果指定多个标签，则会为指定资源同时创建并绑定该多个标签。 同一个资源上的同一个标签键只能对应一个标签值。如果您尝试添加已有标签键，则对应的标签值会更新为新值。 如果标签不存在会为您自动创建标签。 数组最多支持10个元素。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateBlueprintRequest struct {
	*tchttp.BaseRequest
	
	// 镜像名称。最大长度60。
	BlueprintName *string `json:"BlueprintName,omitnil,omitempty" name:"BlueprintName"`

	// 镜像描述。最大长度60。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 需要制作镜像的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/1207/47573) 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否执行强制关机以制作镜像。
	// 取值范围：
	// True：表示关机之后制作镜像
	// False：表示开机状态制作镜像
	// 默认取值：True
	// 开机状态制作镜像，可能导致部分数据未备份，影响数据安全。
	ForcePowerOff *bool `json:"ForcePowerOff,omitnil,omitempty" name:"ForcePowerOff"`

	// 标签键和标签值。 如果指定多个标签，则会为指定资源同时创建并绑定该多个标签。 同一个资源上的同一个标签键只能对应一个标签值。如果您尝试添加已有标签键，则对应的标签值会更新为新值。 如果标签不存在会为您自动创建标签。 数组最多支持10个元素。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "ForcePowerOff")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBlueprintRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBlueprintResponseParams struct {
	// 自定义镜像ID。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 云硬盘ID，可通过 [DescribeDisks](https://cloud.tencent.com/document/api/1207/66093) 接口返回值中的 DiskId 获取。 
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 云硬盘备份点名称，最大长度为 90 。
	DiskBackupName *string `json:"DiskBackupName,omitnil,omitempty" name:"DiskBackupName"`

	// 标签键和标签值。 如果指定多个标签，则会为指定资源同时创建并绑定该多个标签。 同一个资源上的同一个标签键只能对应一个标签值。如果您尝试添加已有标签键，则对应的标签值会更新为新值。 如果标签不存在会为您自动创建标签。 数组最多支持10个元素。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateDiskBackupRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID，可通过 [DescribeDisks](https://cloud.tencent.com/document/api/1207/66093) 接口返回值中的 DiskId 获取。 
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 云硬盘备份点名称，最大长度为 90 。
	DiskBackupName *string `json:"DiskBackupName,omitnil,omitempty" name:"DiskBackupName"`

	// 标签键和标签值。 如果指定多个标签，则会为指定资源同时创建并绑定该多个标签。 同一个资源上的同一个标签键只能对应一个标签值。如果您尝试添加已有标签键，则对应的标签值会更新为新值。 如果标签不存在会为您自动创建标签。 数组最多支持10个元素。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDiskBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDiskBackupResponseParams struct {
	// 备份点ID。
	DiskBackupId *string `json:"DiskBackupId,omitnil,omitempty" name:"DiskBackupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateDisksRequestParams struct {
	// 可用区。可通过[DescribeZones](https://cloud.tencent.com/document/product/1207/57513)返回值中的Zone获取。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 云硬盘大小, 单位: GB。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 云硬盘介质类型。取值: "CLOUD_PREMIUM"(高性能云盘), "CLOUD_SSD"(SSD云硬盘)。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 云硬盘包年包月相关参数设置。
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitnil,omitempty" name:"DiskChargePrepaid"`

	// 云硬盘名称。最大长度60。
	DiskName *string `json:"DiskName,omitnil,omitempty" name:"DiskName"`

	// 云硬盘个数。取值范围: [1, 30]。默认值: 1。
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// 指定云硬盘备份点配额，取值范围: [0, 500]。不传时默认为不带备份点配额。
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitnil,omitempty" name:"DiskBackupQuota"`

	// 是否自动使用代金券。默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 自动挂载并初始化数据盘。
	AutoMountConfiguration *AutoMountConfiguration `json:"AutoMountConfiguration,omitnil,omitempty" name:"AutoMountConfiguration"`

	// 标签键和标签值。 如果指定多个标签，则会为指定资源同时创建并绑定该多个标签。 同一个资源上的同一个标签键只能对应一个标签值。如果您尝试添加已有标签键，则对应的标签值会更新为新值。 如果标签不存在会为您自动创建标签。 数组最多支持10个元素。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateDisksRequest struct {
	*tchttp.BaseRequest
	
	// 可用区。可通过[DescribeZones](https://cloud.tencent.com/document/product/1207/57513)返回值中的Zone获取。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 云硬盘大小, 单位: GB。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 云硬盘介质类型。取值: "CLOUD_PREMIUM"(高性能云盘), "CLOUD_SSD"(SSD云硬盘)。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 云硬盘包年包月相关参数设置。
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitnil,omitempty" name:"DiskChargePrepaid"`

	// 云硬盘名称。最大长度60。
	DiskName *string `json:"DiskName,omitnil,omitempty" name:"DiskName"`

	// 云硬盘个数。取值范围: [1, 30]。默认值: 1。
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// 指定云硬盘备份点配额，取值范围: [0, 500]。不传时默认为不带备份点配额。
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitnil,omitempty" name:"DiskBackupQuota"`

	// 是否自动使用代金券。默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 自动挂载并初始化数据盘。
	AutoMountConfiguration *AutoMountConfiguration `json:"AutoMountConfiguration,omitnil,omitempty" name:"AutoMountConfiguration"`

	// 标签键和标签值。 如果指定多个标签，则会为指定资源同时创建并绑定该多个标签。 同一个资源上的同一个标签键只能对应一个标签值。如果您尝试添加已有标签键，则对应的标签值会更新为新值。 如果标签不存在会为您自动创建标签。 数组最多支持10个元素。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "DiskSize")
	delete(f, "DiskType")
	delete(f, "DiskChargePrepaid")
	delete(f, "DiskName")
	delete(f, "DiskCount")
	delete(f, "DiskBackupQuota")
	delete(f, "AutoVoucher")
	delete(f, "AutoMountConfiguration")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDisksResponseParams struct {
	// 当通过本接口来创建云硬盘时会返回该参数，表示一个或多个云硬盘ID。返回云硬盘ID列表并不代表云硬盘创建成功。
	// 
	// 可根据 [DescribeDisks](https://cloud.tencent.com/document/product/1207/66093) 接口查询返回的DiskSet中对应云硬盘的ID的状态来判断创建是否完成；如果云硬盘状态由“PENDING”变为“UNATTACHED”或“ATTACHED”，则为创建成功。
	DiskIdSet []*string `json:"DiskIdSet,omitnil,omitempty" name:"DiskIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDisksResponse struct {
	*tchttp.BaseResponse
	Response *CreateDisksResponseParams `json:"Response"`
}

func (r *CreateDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFirewallRulesRequestParams struct {
	// 实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/1207/47573) 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 防火墙规则列表。
	FirewallRules []*FirewallRule `json:"FirewallRules,omitnil,omitempty" name:"FirewallRules"`

	// 防火墙当前版本。用户每次更新防火墙规则时版本会自动加1，防止规则已过期，不填不考虑冲突。
	FirewallVersion *uint64 `json:"FirewallVersion,omitnil,omitempty" name:"FirewallVersion"`
}

type CreateFirewallRulesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/1207/47573) 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 防火墙规则列表。
	FirewallRules []*FirewallRule `json:"FirewallRules,omitnil,omitempty" name:"FirewallRules"`

	// 防火墙当前版本。用户每次更新防火墙规则时版本会自动加1，防止规则已过期，不填不考虑冲突。
	FirewallVersion *uint64 `json:"FirewallVersion,omitnil,omitempty" name:"FirewallVersion"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateFirewallTemplateRequestParams struct {
	// 模板名称。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 防火墙规则列表。
	TemplateRules []*FirewallRule `json:"TemplateRules,omitnil,omitempty" name:"TemplateRules"`
}

type CreateFirewallTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 防火墙规则列表。
	TemplateRules []*FirewallRule `json:"TemplateRules,omitnil,omitempty" name:"TemplateRules"`
}

func (r *CreateFirewallTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFirewallTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "TemplateRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFirewallTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFirewallTemplateResponseParams struct {
	// 防火墙模板ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFirewallTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateFirewallTemplateResponseParams `json:"Response"`
}

func (r *CreateFirewallTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFirewallTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFirewallTemplateRulesRequestParams struct {
	// 防火墙模板ID。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 防火墙模板规则列表。
	TemplateRules []*FirewallRule `json:"TemplateRules,omitnil,omitempty" name:"TemplateRules"`
}

type CreateFirewallTemplateRulesRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙模板ID。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 防火墙模板规则列表。
	TemplateRules []*FirewallRule `json:"TemplateRules,omitnil,omitempty" name:"TemplateRules"`
}

func (r *CreateFirewallTemplateRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFirewallTemplateRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TemplateRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFirewallTemplateRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFirewallTemplateRulesResponseParams struct {
	// 规则ID列表。
	TemplateRuleIdSet []*string `json:"TemplateRuleIdSet,omitnil,omitempty" name:"TemplateRuleIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFirewallTemplateRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateFirewallTemplateRulesResponseParams `json:"Response"`
}

func (r *CreateFirewallTemplateRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFirewallTemplateRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceSnapshotRequestParams struct {
	// 需要创建快照的实例 ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/product/1207/47573) 接口返回值中的 InstanceId	获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 快照名称，最长为 60 个字符。
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 标签键和标签值。 如果指定多个标签，则会为指定资源同时创建并绑定该多个标签。 同一个资源上的同一个标签键只能对应一个标签值。如果您尝试添加已有标签键，则对应的标签值会更新为新值。 如果标签不存在会为您自动创建标签。 数组最多支持10个元素。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateInstanceSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 需要创建快照的实例 ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/product/1207/47573) 接口返回值中的 InstanceId	获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 快照名称，最长为 60 个字符。
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 标签键和标签值。 如果指定多个标签，则会为指定资源同时创建并绑定该多个标签。 同一个资源上的同一个标签键只能对应一个标签值。如果您尝试添加已有标签键，则对应的标签值会更新为新值。 如果标签不存在会为您自动创建标签。 数组最多支持10个元素。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceSnapshotResponseParams struct {
	// 快照 ID。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 套餐ID。可以通过调用 [DescribeBundles](https://cloud.tencent.com/document/api/1207/47575) 接口获取。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 镜像ID。可以通过调用 [DescribeBlueprints](https://cloud.tencent.com/document/api/1207/47689) 接口获取。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`

	// 当前实例仅支持预付费模式，即包年包月相关参数设置，单位（月）。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 实例显示名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 购买实例数量。包年包月实例取值范围：[1，30]。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 可用区列表。
	// 不填此参数，表示为随机可用区。
	// 可通过 <a href="https://cloud.tencent.com/document/product/1207/57513" target="_blank">DescribeZones</a>接口获取指定地域下的可用区列表信息
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和库存。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId.
	// false（默认）：发送正常请求，通过检查后直接创建实例
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 实例登录密码信息配置。默认缺失情况下代表用户选择实例创建后设置登录密码。
	LoginConfiguration *LoginConfiguration `json:"LoginConfiguration,omitnil,omitempty" name:"LoginConfiguration"`

	// 要创建的容器配置列表。
	Containers []*DockerContainerConfiguration `json:"Containers,omitnil,omitempty" name:"Containers"`

	// 是否自动使用代金券。默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 防火墙模板ID。若不指定该参数，则使用默认防火墙策略。
	FirewallTemplateId *string `json:"FirewallTemplateId,omitnil,omitempty" name:"FirewallTemplateId"`

	// 标签键和标签值。
	// 如果指定多个标签，则会为指定资源同时创建并绑定该多个标签。
	// 同一个资源上的同一个标签键只能对应一个标签值。如果您尝试添加已有标签键，则对应的标签值会更新为新值。
	// 如果标签不存在会为您自动创建标签。
	// 数组最多支持10个元素。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 创建实例后自动执行的命令。
	InitCommand *Command `json:"InitCommand,omitnil,omitempty" name:"InitCommand"`

	// 主域名。
	// 注意：域名指定后，仅支持购买一台实例（参数InstanceCount=1）。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 子域名。
	// 注意：域名指定后，仅支持购买一台实例（参数InstanceCount=1）。
	Subdomain *string `json:"Subdomain,omitnil,omitempty" name:"Subdomain"`
}

type CreateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 套餐ID。可以通过调用 [DescribeBundles](https://cloud.tencent.com/document/api/1207/47575) 接口获取。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 镜像ID。可以通过调用 [DescribeBlueprints](https://cloud.tencent.com/document/api/1207/47689) 接口获取。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`

	// 当前实例仅支持预付费模式，即包年包月相关参数设置，单位（月）。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 实例显示名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 购买实例数量。包年包月实例取值范围：[1，30]。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 可用区列表。
	// 不填此参数，表示为随机可用区。
	// 可通过 <a href="https://cloud.tencent.com/document/product/1207/57513" target="_blank">DescribeZones</a>接口获取指定地域下的可用区列表信息
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和库存。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId.
	// false（默认）：发送正常请求，通过检查后直接创建实例
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 实例登录密码信息配置。默认缺失情况下代表用户选择实例创建后设置登录密码。
	LoginConfiguration *LoginConfiguration `json:"LoginConfiguration,omitnil,omitempty" name:"LoginConfiguration"`

	// 要创建的容器配置列表。
	Containers []*DockerContainerConfiguration `json:"Containers,omitnil,omitempty" name:"Containers"`

	// 是否自动使用代金券。默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 防火墙模板ID。若不指定该参数，则使用默认防火墙策略。
	FirewallTemplateId *string `json:"FirewallTemplateId,omitnil,omitempty" name:"FirewallTemplateId"`

	// 标签键和标签值。
	// 如果指定多个标签，则会为指定资源同时创建并绑定该多个标签。
	// 同一个资源上的同一个标签键只能对应一个标签值。如果您尝试添加已有标签键，则对应的标签值会更新为新值。
	// 如果标签不存在会为您自动创建标签。
	// 数组最多支持10个元素。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 创建实例后自动执行的命令。
	InitCommand *Command `json:"InitCommand,omitnil,omitempty" name:"InitCommand"`

	// 主域名。
	// 注意：域名指定后，仅支持购买一台实例（参数InstanceCount=1）。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 子域名。
	// 注意：域名指定后，仅支持购买一台实例（参数InstanceCount=1）。
	Subdomain *string `json:"Subdomain,omitnil,omitempty" name:"Subdomain"`
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
	delete(f, "FirewallTemplateId")
	delete(f, "Tags")
	delete(f, "InitCommand")
	delete(f, "DomainName")
	delete(f, "Subdomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesResponseParams struct {
	// 当通过本接口来创建实例时会返回该参数，表示一个或多个实例ID。返回实例ID列表并不代表实例创建成功。
	// 
	// 可根据<a href="https://cloud.tencent.com/document/product/1207/47573" target="_blank">DescribeInstances</a> 接口查询返回的InstancesSet中对应实例的ID的状态来判断创建是否完成；如果实例状态由“启动中”变为“运行中”，则为创建成功。
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// 标签键和标签值。 如果指定多个标签，则会为指定资源同时创建并绑定该多个标签。 同一个资源上的同一个标签键只能对应一个标签值。如果您尝试添加已有标签键，则对应的标签值会更新为新值。 如果标签不存在会为您自动创建标签。 数组最多支持10个元素。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateKeyPairRequest struct {
	*tchttp.BaseRequest
	
	// 密钥对名称，可由数字，字母和下划线组成，长度不超过 25 个字符。
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// 标签键和标签值。 如果指定多个标签，则会为指定资源同时创建并绑定该多个标签。 同一个资源上的同一个标签键只能对应一个标签值。如果您尝试添加已有标签键，则对应的标签值会更新为新值。 如果标签不存在会为您自动创建标签。 数组最多支持10个元素。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKeyPairRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeyPairResponseParams struct {
	// 密钥对信息。
	KeyPair *KeyPair `json:"KeyPair,omitnil,omitempty" name:"KeyPair"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CreateMcpServerRequestParams struct {
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// MCP Server名称。最大长度：64
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Base64编码后的MCP Server启动命令。最大长度：2048
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// MCP Server备注。最大长度：2048
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// MCP Server环境变量。最大长度：10
	Envs []*McpServerEnv `json:"Envs,omitnil,omitempty" name:"Envs"`
}

type CreateMcpServerRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// MCP Server名称。最大长度：64
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Base64编码后的MCP Server启动命令。最大长度：2048
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// MCP Server备注。最大长度：2048
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// MCP Server环境变量。最大长度：10
	Envs []*McpServerEnv `json:"Envs,omitnil,omitempty" name:"Envs"`
}

func (r *CreateMcpServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMcpServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Command")
	delete(f, "Description")
	delete(f, "Envs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMcpServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMcpServerResponseParams struct {
	// MCP Server ID。
	McpServerId *string `json:"McpServerId,omitnil,omitempty" name:"McpServerId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMcpServerResponse struct {
	*tchttp.BaseResponse
	Response *CreateMcpServerResponseParams `json:"Response"`
}

func (r *CreateMcpServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMcpServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataDiskPrice struct {
	// 云硬盘ID。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 云硬盘单价。
	OriginalDiskPrice *float64 `json:"OriginalDiskPrice,omitnil,omitempty" name:"OriginalDiskPrice"`

	// 云硬盘总价。
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 折扣。
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 折后总价。
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// 数据盘挂载的实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

// Predefined struct for user
type DeleteBlueprintsRequestParams struct {
	// 镜像ID列表。镜像ID，可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值中的BlueprintId获取。
	BlueprintIds []*string `json:"BlueprintIds,omitnil,omitempty" name:"BlueprintIds"`
}

type DeleteBlueprintsRequest struct {
	*tchttp.BaseRequest
	
	// 镜像ID列表。镜像ID，可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值中的BlueprintId获取。
	BlueprintIds []*string `json:"BlueprintIds,omitnil,omitempty" name:"BlueprintIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 云硬盘备份点ID列表，可通过 [DescribeDiskBackups](https://cloud.tencent.com/document/api/1207/84379)接口查询。列表长度最大值为100。
	DiskBackupIds []*string `json:"DiskBackupIds,omitnil,omitempty" name:"DiskBackupIds"`
}

type DeleteDiskBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘备份点ID列表，可通过 [DescribeDiskBackups](https://cloud.tencent.com/document/api/1207/84379)接口查询。列表长度最大值为100。
	DiskBackupIds []*string `json:"DiskBackupIds,omitnil,omitempty" name:"DiskBackupIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/1207/47573) 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 防火墙规则列表。
	FirewallRules []*FirewallRule `json:"FirewallRules,omitnil,omitempty" name:"FirewallRules"`

	// 防火墙当前版本。用户每次更新防火墙规则时版本会自动加1，防止规则已过期，不填不考虑冲突。
	FirewallVersion *uint64 `json:"FirewallVersion,omitnil,omitempty" name:"FirewallVersion"`
}

type DeleteFirewallRulesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/1207/47573) 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 防火墙规则列表。
	FirewallRules []*FirewallRule `json:"FirewallRules,omitnil,omitempty" name:"FirewallRules"`

	// 防火墙当前版本。用户每次更新防火墙规则时版本会自动加1，防止规则已过期，不填不考虑冲突。
	FirewallVersion *uint64 `json:"FirewallVersion,omitnil,omitempty" name:"FirewallVersion"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteFirewallTemplateRequestParams struct {
	// 防火墙模板ID。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteFirewallTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙模板ID。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeleteFirewallTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFirewallTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFirewallTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFirewallTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFirewallTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFirewallTemplateResponseParams `json:"Response"`
}

func (r *DeleteFirewallTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFirewallTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFirewallTemplateRulesRequestParams struct {
	// 防火墙模板ID。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 防火墙模板规则ID列表。可通过[DescribeFirewallTemplateRules](https://cloud.tencent.com/document/product/1207/96875)接口返回值字段TemplateRuleSet获取。
	TemplateRuleIds []*string `json:"TemplateRuleIds,omitnil,omitempty" name:"TemplateRuleIds"`
}

type DeleteFirewallTemplateRulesRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙模板ID。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 防火墙模板规则ID列表。可通过[DescribeFirewallTemplateRules](https://cloud.tencent.com/document/product/1207/96875)接口返回值字段TemplateRuleSet获取。
	TemplateRuleIds []*string `json:"TemplateRuleIds,omitnil,omitempty" name:"TemplateRuleIds"`
}

func (r *DeleteFirewallTemplateRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFirewallTemplateRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TemplateRuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFirewallTemplateRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFirewallTemplateRulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFirewallTemplateRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFirewallTemplateRulesResponseParams `json:"Response"`
}

func (r *DeleteFirewallTemplateRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFirewallTemplateRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteKeyPairsRequestParams struct {
	// 密钥对 ID 列表，每次请求批量密钥对的上限为 10。可通过[DescribeKeyPairs](https://cloud.tencent.com/document/api/1207/55540)接口返回值中的KeyId获取。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

type DeleteKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// 密钥对 ID 列表，每次请求批量密钥对的上限为 10。可通过[DescribeKeyPairs](https://cloud.tencent.com/document/api/1207/55540)接口返回值中的KeyId获取。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 要删除的快照 ID 列表，每次请求批量快照的上限为10个，可通过 <a href="https://cloud.tencent.com/document/product/1207/54388" target="_blank">DescribeSnapshots</a>查询。
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`
}

type DeleteSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的快照 ID 列表，每次请求批量快照的上限为10个，可通过 <a href="https://cloud.tencent.com/document/product/1207/54388" target="_blank">DescribeSnapshots</a>查询。
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 限制操作消息码。
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 限制操作消息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type DescribeAllScenesRequestParams struct {
	// 使用场景ID列表。可通过[DescribeAllScenes](https://cloud.tencent.com/document/product/1207/83513)接口返回值中的SceneId获取。
	SceneIds []*string `json:"SceneIds,omitnil,omitempty" name:"SceneIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAllScenesRequest struct {
	*tchttp.BaseRequest
	
	// 使用场景ID列表。可通过[DescribeAllScenes](https://cloud.tencent.com/document/product/1207/83513)接口返回值中的SceneId获取。
	SceneIds []*string `json:"SceneIds,omitnil,omitempty" name:"SceneIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	SceneInfoSet []*SceneInfo `json:"SceneInfoSet,omitnil,omitempty" name:"SceneInfoSet"`

	// 使用场景详细信息总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例 ID 列表。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。 当前最多支持1个。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeBlueprintInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。 当前最多支持1个。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 镜像实例列表信息。
	BlueprintInstanceSet []*BlueprintInstance `json:"BlueprintInstanceSet,omitnil,omitempty" name:"BlueprintInstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 镜像 ID 列表。可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值字段BlueprintSet获取。列表长度最大值为100。
	BlueprintIds []*string `json:"BlueprintIds,omitnil,omitempty" name:"BlueprintIds"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤器列表。
	// <li>blueprint-id</li>按照【镜像 ID】进行过滤。
	// 类型：String
	// 必选：否
	// 镜像 ID ，可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值字段BlueprintSet获取。
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
	// 镜像状态，可通过[数据结构Blueprint](https://cloud.tencent.com/document/api/1207/47576#Blueprint)中的BlueprintState来获取。
	// <li>scene-id</li>按照【使用场景Id】进行过滤。
	// 类型：String
	// 必选：否
	// 场景Id，可通过[查看使用场景列表](https://cloud.tencent.com/document/product/1207/83512)接口获取。
	// <li>tag-key</li>
	// 按照【标签键】进行过滤。 类型：String 必选：否
	// <li>tag-value</li>
	// 按照【标签值】进行过滤。 类型：String 必选：否
	// <li>tag:tag-key</li>
	// 按照【标签键值对】进行过滤。 tag-key使用具体的标签键进行替换。
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 100。参数不支持同时指定 BlueprintIds (可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值字段BlueprintSet获取BlueprintId)和 Filters 。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeBlueprintsRequest struct {
	*tchttp.BaseRequest
	
	// 镜像 ID 列表。可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值字段BlueprintSet获取。列表长度最大值为100。
	BlueprintIds []*string `json:"BlueprintIds,omitnil,omitempty" name:"BlueprintIds"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤器列表。
	// <li>blueprint-id</li>按照【镜像 ID】进行过滤。
	// 类型：String
	// 必选：否
	// 镜像 ID ，可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值字段BlueprintSet获取。
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
	// 镜像状态，可通过[数据结构Blueprint](https://cloud.tencent.com/document/api/1207/47576#Blueprint)中的BlueprintState来获取。
	// <li>scene-id</li>按照【使用场景Id】进行过滤。
	// 类型：String
	// 必选：否
	// 场景Id，可通过[查看使用场景列表](https://cloud.tencent.com/document/product/1207/83512)接口获取。
	// <li>tag-key</li>
	// 按照【标签键】进行过滤。 类型：String 必选：否
	// <li>tag-value</li>
	// 按照【标签值】进行过滤。 类型：String 必选：否
	// <li>tag:tag-key</li>
	// 按照【标签键值对】进行过滤。 tag-key使用具体的标签键进行替换。
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 100。参数不支持同时指定 BlueprintIds (可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值字段BlueprintSet获取BlueprintId)和 Filters 。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 镜像详细信息列表。
	BlueprintSet []*Blueprint `json:"BlueprintSet,omitnil,omitempty" name:"BlueprintSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 套餐 ID。可通过[DescribeBundles](https://cloud.tencent.com/document/product/1207/47575)接口返回值中的BundleId获取。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`
}

type DescribeBundleDiscountRequest struct {
	*tchttp.BaseRequest
	
	// 套餐 ID。可通过[DescribeBundles](https://cloud.tencent.com/document/product/1207/47575)接口返回值中的BundleId获取。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`
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
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 折扣梯度详情，每个梯度包含的信息有：时长，折扣数，总价，折扣价，折扣详情（用户折扣、官网折扣、最终折扣）。
	DiscountDetail []*DiscountDetail `json:"DiscountDetail,omitnil,omitempty" name:"DiscountDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 套餐 ID 列表。每次请求批量套餐的上限为 100。可通过[DescribeBundles](https://cloud.tencent.com/document/product/1207/47575)接口返回值中的BundleId获取。
	BundleIds []*string `json:"BundleIds,omitnil,omitempty" name:"BundleIds"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤器列表。
	// <li>bundle-id</li>按照【套餐 ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>support-platform-type</li>按照【系统类型】进行过滤。
	// 取值： LINUX_UNIX(Linux/Unix系统) ;WINDOWS(Windows 系统)
	// 类型：String
	// 必选：否
	// <li>bundle-type</li>按照 【套餐类型进行过滤】。
	// 取值：GENERAL_BUNDLE (通用型套餐); STORAGE_BUNDLE(存储型套餐);ENTERPRISE_BUNDLE( 企业型套餐);EXCLUSIVE_BUNDLE(专属型套餐);BEFAST_BUNDLE(蜂驰型套餐);STARTER_BUNDLE(入门型套餐);CAREFREE_BUNDLE(无忧型套餐);RAZOR_SPEED_BUNDLE(锐驰型套餐)
	// 类型：String
	// 必选：否
	// <li>bundle-state</li>按照【套餐状态】进行过滤。
	// 取值: ONLINE(在线); OFFLINE(下线);
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。参数不支持同时指定 BundleIds 和 Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 可用区列表。默认为全部可用区。
	// <li>可用区可通过接口 [DescribeZones](https://cloud.tencent.com/document/product/1207/57513) 查询</li>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`
}

type DescribeBundlesRequest struct {
	*tchttp.BaseRequest
	
	// 套餐 ID 列表。每次请求批量套餐的上限为 100。可通过[DescribeBundles](https://cloud.tencent.com/document/product/1207/47575)接口返回值中的BundleId获取。
	BundleIds []*string `json:"BundleIds,omitnil,omitempty" name:"BundleIds"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤器列表。
	// <li>bundle-id</li>按照【套餐 ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>support-platform-type</li>按照【系统类型】进行过滤。
	// 取值： LINUX_UNIX(Linux/Unix系统) ;WINDOWS(Windows 系统)
	// 类型：String
	// 必选：否
	// <li>bundle-type</li>按照 【套餐类型进行过滤】。
	// 取值：GENERAL_BUNDLE (通用型套餐); STORAGE_BUNDLE(存储型套餐);ENTERPRISE_BUNDLE( 企业型套餐);EXCLUSIVE_BUNDLE(专属型套餐);BEFAST_BUNDLE(蜂驰型套餐);STARTER_BUNDLE(入门型套餐);CAREFREE_BUNDLE(无忧型套餐);RAZOR_SPEED_BUNDLE(锐驰型套餐)
	// 类型：String
	// 必选：否
	// <li>bundle-state</li>按照【套餐状态】进行过滤。
	// 取值: ONLINE(在线); OFFLINE(下线);
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。参数不支持同时指定 BundleIds 和 Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 可用区列表。默认为全部可用区。
	// <li>可用区可通过接口 [DescribeZones](https://cloud.tencent.com/document/product/1207/57513) 查询</li>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`
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
	BundleSet []*Bundle `json:"BundleSet,omitnil,omitempty" name:"BundleSet"`

	// 符合要求的套餐总数，用于分页展示。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CcnAttachedInstanceSet []*CcnAttachedInstance `json:"CcnAttachedInstanceSet,omitnil,omitempty" name:"CcnAttachedInstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 云硬盘备份点 ID 列表, 可通过<a href="https://cloud.tencent.com/document/product/1207/84379" target="_blank">DescribeDiskBackups</a>接口查询。列表长度最大值为100。
	DiskBackupIds []*string `json:"DiskBackupIds,omitnil,omitempty" name:"DiskBackupIds"`
}

type DescribeDiskBackupsDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘备份点 ID 列表, 可通过<a href="https://cloud.tencent.com/document/product/1207/84379" target="_blank">DescribeDiskBackups</a>接口查询。列表长度最大值为100。
	DiskBackupIds []*string `json:"DiskBackupIds,omitnil,omitempty" name:"DiskBackupIds"`
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
	DiskBackupDeniedActionSet []*DiskBackupDeniedActions `json:"DiskBackupDeniedActionSet,omitnil,omitempty" name:"DiskBackupDeniedActionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 查询的云硬盘备份点ID列表。可通过[DescribeDiskBackups](https://cloud.tencent.com/document/product/1207/84379)接口返回值字段DiskBackupSet获取。列表长度最大值为100。参数不支持同时指定 DiskBackupIds 和 Filters。
	DiskBackupIds []*string `json:"DiskBackupIds,omitnil,omitempty" name:"DiskBackupIds"`

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
	// 取值：参考数据结构 [DiskBackup](https://cloud.tencent.com/document/product/1207/47576#DiskBackup) 下的DiskBackupState取值。
	// <li>disk-usage</li>按照【云硬盘类型】进行过滤。
	// 类型：String
	// 必选：否
	// 取值：
	// - SYSTEM_DISK - 系统盘
	// - DATA_DISK - 数据盘
	// <li>tag-key</li>
	// 按照【标签键】进行过滤。 类型：String 必选：否
	// <li>tag-value</li>
	// 按照【标签值】进行过滤。 类型：String 必选：否
	// <li>tag:tag-key</li>
	// 按照【标签键值对】进行过滤。 tag-key使用具体的标签键进行替换。
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为5。参数不支持同时指定DiskBackupIds 和 Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDiskBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 查询的云硬盘备份点ID列表。可通过[DescribeDiskBackups](https://cloud.tencent.com/document/product/1207/84379)接口返回值字段DiskBackupSet获取。列表长度最大值为100。参数不支持同时指定 DiskBackupIds 和 Filters。
	DiskBackupIds []*string `json:"DiskBackupIds,omitnil,omitempty" name:"DiskBackupIds"`

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
	// 取值：参考数据结构 [DiskBackup](https://cloud.tencent.com/document/product/1207/47576#DiskBackup) 下的DiskBackupState取值。
	// <li>disk-usage</li>按照【云硬盘类型】进行过滤。
	// 类型：String
	// 必选：否
	// 取值：
	// - SYSTEM_DISK - 系统盘
	// - DATA_DISK - 数据盘
	// <li>tag-key</li>
	// 按照【标签键】进行过滤。 类型：String 必选：否
	// <li>tag-value</li>
	// 按照【标签值】进行过滤。 类型：String 必选：否
	// <li>tag:tag-key</li>
	// 按照【标签键值对】进行过滤。 tag-key使用具体的标签键进行替换。
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为5。参数不支持同时指定DiskBackupIds 和 Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 云硬盘备份点信息列表。
	DiskBackupSet []*DiskBackup `json:"DiskBackupSet,omitnil,omitempty" name:"DiskBackupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeDiskConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器列表。
	// <li>zone</li>按照【可用区】进行过滤。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	DiskConfigSet []*DiskConfig `json:"DiskConfigSet,omitnil,omitempty" name:"DiskConfigSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 云硬盘类型, 取值范围: CLOUD_PREMIUM: 高性能云硬盘，CLOUD_SSD: SSD云硬盘
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 云硬盘大小, 单位: GB。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 指定云硬盘备份点配额，不传时默认为不带备份点配额。目前只支持不带或设置[0 - 500]个云硬盘备份点配额。
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitnil,omitempty" name:"DiskBackupQuota"`
}

type DescribeDiskDiscountRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘类型, 取值范围: CLOUD_PREMIUM: 高性能云硬盘，CLOUD_SSD: SSD云硬盘
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 云硬盘大小, 单位: GB。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 指定云硬盘备份点配额，不传时默认为不带备份点配额。目前只支持不带或设置[0 - 500]个云硬盘备份点配额。
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitnil,omitempty" name:"DiskBackupQuota"`
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
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 折扣梯度详情，每个梯度包含的信息有：时长，折扣数，总价，折扣价，折扣详情（用户折扣、官网折扣、最终折扣）。
	DiscountDetail []*DiscountDetail `json:"DiscountDetail,omitnil,omitempty" name:"DiscountDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 云硬盘ID列表。每次批量请求云硬盘的上限为 100。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`
}

type DescribeDisksDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。每次批量请求云硬盘的上限为 100。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`
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
	DiskDeniedActionSet []*DiskDeniedActions `json:"DiskDeniedActionSet,omitnil,omitempty" name:"DiskDeniedActionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 云硬盘ID列表。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值字段KeyPairSet获取。列表长度最大值为100。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

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
	// 取值：SYSTEM_DISK（系统盘）或 DATA_DISK（数据盘）
	// disk-state
	// 按照【云硬盘状态】进行过滤。
	// 类型：String
	// 必选：否
	// 取值：参考数据结构[Disk](https://cloud.tencent.com/document/api/1207/47576#Disk)中DiskState取值。
	// tag-key
	// 按照【标签键】进行过滤。 类型：String 必选：否
	// tag-value
	// 按照【标签值】进行过滤。 类型：String 必选：否
	// tag:tag-key
	// 按照【标签键值对】进行过滤。 tag-key使用具体的标签键进行替换。
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 100。参数不支持同时指定 DiskIds 和 Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 云硬盘列表排序的依据字段。取值范围："CREATED_TIME"：依据云硬盘的创建时间排序。 "EXPIRED_TIME"：依据云硬盘的到期时间排序。"DISK_SIZE"：依据云硬盘的大小排序。默认按云硬盘创建时间排序。
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 输出云硬盘列表的排列顺序。取值范围："ASC"：升序排列。 "DESC"：降序排列。默认按降序排列。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值字段KeyPairSet获取。列表长度最大值为100。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

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
	// 取值：SYSTEM_DISK（系统盘）或 DATA_DISK（数据盘）
	// disk-state
	// 按照【云硬盘状态】进行过滤。
	// 类型：String
	// 必选：否
	// 取值：参考数据结构[Disk](https://cloud.tencent.com/document/api/1207/47576#Disk)中DiskState取值。
	// tag-key
	// 按照【标签键】进行过滤。 类型：String 必选：否
	// tag-value
	// 按照【标签值】进行过滤。 类型：String 必选：否
	// tag:tag-key
	// 按照【标签键值对】进行过滤。 tag-key使用具体的标签键进行替换。
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 100。参数不支持同时指定 DiskIds 和 Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 云硬盘列表排序的依据字段。取值范围："CREATED_TIME"：依据云硬盘的创建时间排序。 "EXPIRED_TIME"：依据云硬盘的到期时间排序。"DISK_SIZE"：依据云硬盘的大小排序。默认按云硬盘创建时间排序。
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 输出云硬盘列表的排列顺序。取值范围："ASC"：升序排列。 "DESC"：降序排列。默认按降序排列。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
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
	DiskSet []*Disk `json:"DiskSet,omitnil,omitempty" name:"DiskSet"`

	// 符合条件的云硬盘信息数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 云硬盘ID列表。每次批量请求云硬盘的上限为 10。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeDisksReturnableRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。每次批量请求云硬盘的上限为 10。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	DiskReturnableSet []*DiskReturnable `json:"DiskReturnableSet,omitnil,omitempty" name:"DiskReturnableSet"`

	// 符合条件的云硬盘数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeDockerActivitiesRequestParams struct {
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Docker活动ID列表。可通过[DescribeDockerActivities](https://cloud.tencent.com/document/product/1207/95476)接口返回值中的ActivityId获取。
	ActivityIds []*string `json:"ActivityIds,omitnil,omitempty" name:"ActivityIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 活动创建时间的起始值，时间戳秒数。
	CreatedTimeBegin *int64 `json:"CreatedTimeBegin,omitnil,omitempty" name:"CreatedTimeBegin"`

	// 活动创建时间的结束值，时间戳秒数。
	CreatedTimeEnd *int64 `json:"CreatedTimeEnd,omitnil,omitempty" name:"CreatedTimeEnd"`
}

type DescribeDockerActivitiesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Docker活动ID列表。可通过[DescribeDockerActivities](https://cloud.tencent.com/document/product/1207/95476)接口返回值中的ActivityId获取。
	ActivityIds []*string `json:"ActivityIds,omitnil,omitempty" name:"ActivityIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 活动创建时间的起始值，时间戳秒数。
	CreatedTimeBegin *int64 `json:"CreatedTimeBegin,omitnil,omitempty" name:"CreatedTimeBegin"`

	// 活动创建时间的结束值，时间戳秒数。
	CreatedTimeEnd *int64 `json:"CreatedTimeEnd,omitnil,omitempty" name:"CreatedTimeEnd"`
}

func (r *DescribeDockerActivitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDockerActivitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ActivityIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CreatedTimeBegin")
	delete(f, "CreatedTimeEnd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDockerActivitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDockerActivitiesResponseParams struct {
	// 总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Docker活动列表。
	DockerActivitySet []*DockerActivity `json:"DockerActivitySet,omitnil,omitempty" name:"DockerActivitySet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDockerActivitiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDockerActivitiesResponseParams `json:"Response"`
}

func (r *DescribeDockerActivitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDockerActivitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDockerContainerConfigurationRequestParams struct {
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 容器ID。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。
	ContainerId *string `json:"ContainerId,omitnil,omitempty" name:"ContainerId"`
}

type DescribeDockerContainerConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 容器ID。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。
	ContainerId *string `json:"ContainerId,omitnil,omitempty" name:"ContainerId"`
}

func (r *DescribeDockerContainerConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDockerContainerConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ContainerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDockerContainerConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDockerContainerConfigurationResponseParams struct {
	// Docker容器配置信息。
	ContainerConfiguration *DockerContainerConfiguration `json:"ContainerConfiguration,omitnil,omitempty" name:"ContainerConfiguration"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDockerContainerConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDockerContainerConfigurationResponseParams `json:"Response"`
}

func (r *DescribeDockerContainerConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDockerContainerConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDockerContainerDetailRequestParams struct {
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 容器ID。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。
	ContainerId *string `json:"ContainerId,omitnil,omitempty" name:"ContainerId"`
}

type DescribeDockerContainerDetailRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 容器ID。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。
	ContainerId *string `json:"ContainerId,omitnil,omitempty" name:"ContainerId"`
}

func (r *DescribeDockerContainerDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDockerContainerDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ContainerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDockerContainerDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDockerContainerDetailResponseParams struct {
	// Docker容器详情，json字符串base64编码。
	ContainerDetail *string `json:"ContainerDetail,omitnil,omitempty" name:"ContainerDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDockerContainerDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDockerContainerDetailResponseParams `json:"Response"`
}

func (r *DescribeDockerContainerDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDockerContainerDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDockerContainersRequestParams struct {
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 容器ID列表。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。参数不支持同时指定 ContainerIds 和 Filters。
	ContainerIds []*string `json:"ContainerIds,omitnil,omitempty" name:"ContainerIds"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤器列表。
	// <li>container-id</li>按照【容器ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>container-name</li>按照【容器名称】进行过滤。
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。参数不支持同时指定 ContainerIds 和 Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeDockerContainersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 容器ID列表。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。参数不支持同时指定 ContainerIds 和 Filters。
	ContainerIds []*string `json:"ContainerIds,omitnil,omitempty" name:"ContainerIds"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤器列表。
	// <li>container-id</li>按照【容器ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>container-name</li>按照【容器名称】进行过滤。
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。参数不支持同时指定 ContainerIds 和 Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeDockerContainersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDockerContainersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ContainerIds")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDockerContainersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDockerContainersResponseParams struct {
	// 总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 容器列表。
	DockerContainerSet []*DockerContainer `json:"DockerContainerSet,omitnil,omitempty" name:"DockerContainerSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDockerContainersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDockerContainersResponseParams `json:"Response"`
}

func (r *DescribeDockerContainersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDockerContainersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirewallRulesRequestParams struct {
	// 实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/1207/47573) 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeFirewallRulesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/1207/47573) 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 防火墙规则详细信息列表。
	FirewallRuleSet []*FirewallRuleInfo `json:"FirewallRuleSet,omitnil,omitempty" name:"FirewallRuleSet"`

	// 防火墙版本号。
	FirewallVersion *uint64 `json:"FirewallVersion,omitnil,omitempty" name:"FirewallVersion"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 防火墙规则详细信息列表。
	FirewallRuleSet []*FirewallRuleInfo `json:"FirewallRuleSet,omitnil,omitempty" name:"FirewallRuleSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeFirewallTemplateApplyRecordsRequestParams struct {
	// 防火墙模板ID。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 应用防火墙模板任务ID列表。可通过[ApplyFirewallTemplate](https://cloud.tencent.com/document/product/1207/96883)接口返回值TaskId字段获取。
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

type DescribeFirewallTemplateApplyRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙模板ID。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 应用防火墙模板任务ID列表。可通过[ApplyFirewallTemplate](https://cloud.tencent.com/document/product/1207/96883)接口返回值TaskId字段获取。
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

func (r *DescribeFirewallTemplateApplyRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallTemplateApplyRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TaskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFirewallTemplateApplyRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirewallTemplateApplyRecordsResponseParams struct {
	// 防火墙模板应用记录列表。
	ApplyRecordSet []*FirewallTemplateApplyRecord `json:"ApplyRecordSet,omitnil,omitempty" name:"ApplyRecordSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFirewallTemplateApplyRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFirewallTemplateApplyRecordsResponseParams `json:"Response"`
}

func (r *DescribeFirewallTemplateApplyRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallTemplateApplyRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirewallTemplateQuotaRequestParams struct {

}

type DescribeFirewallTemplateQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeFirewallTemplateQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallTemplateQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFirewallTemplateQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirewallTemplateQuotaResponseParams struct {
	// 当前可用配额。
	Available *int64 `json:"Available,omitnil,omitempty" name:"Available"`

	// 总配额。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFirewallTemplateQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFirewallTemplateQuotaResponseParams `json:"Response"`
}

func (r *DescribeFirewallTemplateQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallTemplateQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirewallTemplateRuleQuotaRequestParams struct {
	// 防火墙模板ID。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeFirewallTemplateRuleQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙模板ID。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeFirewallTemplateRuleQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallTemplateRuleQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFirewallTemplateRuleQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirewallTemplateRuleQuotaResponseParams struct {
	// 当前可用配额。
	Available *int64 `json:"Available,omitnil,omitempty" name:"Available"`

	// 总配额。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFirewallTemplateRuleQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFirewallTemplateRuleQuotaResponseParams `json:"Response"`
}

func (r *DescribeFirewallTemplateRuleQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallTemplateRuleQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirewallTemplateRulesRequestParams struct {
	// 防火墙模板ID列表。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。列表长度最大值为100。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 防火墙模板规则ID列表。可通过[DescribeFirewallTemplateRules](https://cloud.tencent.com/document/product/1207/96875)接口返回值字段TemplateRuleSet获取。列表长度最大值为100。
	TemplateRuleIds []*string `json:"TemplateRuleIds,omitnil,omitempty" name:"TemplateRuleIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeFirewallTemplateRulesRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙模板ID列表。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。列表长度最大值为100。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 防火墙模板规则ID列表。可通过[DescribeFirewallTemplateRules](https://cloud.tencent.com/document/product/1207/96875)接口返回值字段TemplateRuleSet获取。列表长度最大值为100。
	TemplateRuleIds []*string `json:"TemplateRuleIds,omitnil,omitempty" name:"TemplateRuleIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeFirewallTemplateRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallTemplateRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TemplateRuleIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFirewallTemplateRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirewallTemplateRulesResponseParams struct {
	// 防火墙模板规则总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 防火墙模板规则信息列表。
	TemplateRuleSet []*FirewallTemplateRuleInfo `json:"TemplateRuleSet,omitnil,omitempty" name:"TemplateRuleSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFirewallTemplateRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFirewallTemplateRulesResponseParams `json:"Response"`
}

func (r *DescribeFirewallTemplateRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallTemplateRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirewallTemplatesRequestParams struct {
	// 防火墙模板ID列表。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。列表长度最大值为100。
	TemplateIds []*string `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`

	// 过滤器列表。
	// <li>template-id</li>按照【防火墙模板所属的ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>template-name</li>按照【防火墙模板所属的名称】进行过滤。
	// 类型：String
	// 必选：否
	// <li>template-type</li>按照【防火墙模板的类型】进行过滤。
	// 类型：String
	// 必选：否
	// 取值: "PRIVATE"(个人模板)
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 100。参数不支持同时指定 TemplateIds 和 Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeFirewallTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙模板ID列表。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。列表长度最大值为100。
	TemplateIds []*string `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`

	// 过滤器列表。
	// <li>template-id</li>按照【防火墙模板所属的ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>template-name</li>按照【防火墙模板所属的名称】进行过滤。
	// 类型：String
	// 必选：否
	// <li>template-type</li>按照【防火墙模板的类型】进行过滤。
	// 类型：String
	// 必选：否
	// 取值: "PRIVATE"(个人模板)
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 100。参数不支持同时指定 TemplateIds 和 Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeFirewallTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFirewallTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirewallTemplatesResponseParams struct {
	// 模板总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 防火墙模板列表。
	TemplateSet []*FirewallTemplate `json:"TemplateSet,omitnil,omitempty" name:"TemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFirewallTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFirewallTemplatesResponseParams `json:"Response"`
}

func (r *DescribeFirewallTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallTemplatesResponse) FromJsonString(s string) error {
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
	// - STARTER_BUNDLE_INSTANCE 入门型套餐实例
	// - HK_EXCLUSIVE_BUNDLE_INSTANCE 中国香港专属型套餐实例
	// - CAREFREE_BUNDLE_INSTANCE 无忧型套餐实例
	// - EXCLUSIVE_BUNDLE_02_INSTANCE 境外专属Ⅱ型
	// - NEWCOMER_BUNDLE_INSTANCE 新客专享型
	// - GAME_PORTAL_BUNDLE_INSTANCE 游戏专区实例
	// - ECONOMY_BUNDLE_INSTANCE 经济型套餐实例
	// - BUDGET_BUNDLE_INSTANCE 特惠型套餐实例
	// - RAZOR_SPEED_BUNDLE_INSTANCE 锐驰套餐实例
	// - BANDWIDTH_BUNDLE_INSTANCE 带宽型套餐实例
	// - USER_KEY_PAIR 密钥对
	// - SNAPSHOT 快照
	// - BLUEPRINT 自定义镜像
	// - FREE_BLUEPRINT 免费自定义镜像
	// - DATA_DISK 数据盘
	ResourceNames []*string `json:"ResourceNames,omitnil,omitempty" name:"ResourceNames"`
}

type DescribeGeneralResourceQuotasRequest struct {
	*tchttp.BaseRequest
	
	// 资源名列表，可取值:
	// - GENERAL_BUNDLE_INSTANCE 通用型套餐实例
	// - STORAGE_BUNDLE_INSTANCE 存储型套餐实例 
	// - ENTERPRISE_BUNDLE_INSTANCE 企业型套餐实例 
	// - EXCLUSIVE_BUNDLE_INSTANCE 专属型套餐实例
	// - BEFAST_BUNDLE_INSTANCE 蜂驰型套餐实例
	// - STARTER_BUNDLE_INSTANCE 入门型套餐实例
	// - HK_EXCLUSIVE_BUNDLE_INSTANCE 中国香港专属型套餐实例
	// - CAREFREE_BUNDLE_INSTANCE 无忧型套餐实例
	// - EXCLUSIVE_BUNDLE_02_INSTANCE 境外专属Ⅱ型
	// - NEWCOMER_BUNDLE_INSTANCE 新客专享型
	// - GAME_PORTAL_BUNDLE_INSTANCE 游戏专区实例
	// - ECONOMY_BUNDLE_INSTANCE 经济型套餐实例
	// - BUDGET_BUNDLE_INSTANCE 特惠型套餐实例
	// - RAZOR_SPEED_BUNDLE_INSTANCE 锐驰套餐实例
	// - BANDWIDTH_BUNDLE_INSTANCE 带宽型套餐实例
	// - USER_KEY_PAIR 密钥对
	// - SNAPSHOT 快照
	// - BLUEPRINT 自定义镜像
	// - FREE_BLUEPRINT 免费自定义镜像
	// - DATA_DISK 数据盘
	ResourceNames []*string `json:"ResourceNames,omitnil,omitempty" name:"ResourceNames"`
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
	GeneralResourceQuotaSet []*GeneralResourceQuota `json:"GeneralResourceQuotaSet,omitnil,omitempty" name:"GeneralResourceQuotaSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeImagesToShareRequestParams struct {
	// CVM镜像 ID 列表。可通过[DescribeImages](https://cloud.tencent.com/document/api/213/15715)接口返回值中的ImageId获取。
	ImageIds []*string `json:"ImageIds,omitnil,omitempty" name:"ImageIds"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤器列表。
	// <li>image-id</li>按照【CVM镜像ID】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// <li>image-name</li>按照【CVM镜像名称】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// <li>image-type</li>按照【CVM镜像类型】进行过滤。
	// 类型：String
	// 必选：否
	// 取值范围：
	// PRIVATE_IMAGE: 私有镜像 (本账户创建的镜像)
	// PUBLIC_IMAGE: 公共镜像 (腾讯云官方镜像)
	// SHARED_IMAGE: 共享镜像(其他账户共享给本账户的镜像) 。
	// 
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。
	// 参数不可以同时指定ImageIds和Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeImagesToShareRequest struct {
	*tchttp.BaseRequest
	
	// CVM镜像 ID 列表。可通过[DescribeImages](https://cloud.tencent.com/document/api/213/15715)接口返回值中的ImageId获取。
	ImageIds []*string `json:"ImageIds,omitnil,omitempty" name:"ImageIds"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤器列表。
	// <li>image-id</li>按照【CVM镜像ID】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// <li>image-name</li>按照【CVM镜像名称】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// <li>image-type</li>按照【CVM镜像类型】进行过滤。
	// 类型：String
	// 必选：否
	// 取值范围：
	// PRIVATE_IMAGE: 私有镜像 (本账户创建的镜像)
	// PUBLIC_IMAGE: 公共镜像 (腾讯云官方镜像)
	// SHARED_IMAGE: 共享镜像(其他账户共享给本账户的镜像) 。
	// 
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。
	// 参数不可以同时指定ImageIds和Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeImagesToShareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesToShareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImagesToShareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesToShareResponseParams struct {
	// 符合条件的镜像数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// CVM镜像详细信息列表。
	ImageSet []*Image `json:"ImageSet,omitnil,omitempty" name:"ImageSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImagesToShareResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImagesToShareResponseParams `json:"Response"`
}

func (r *DescribeImagesToShareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesToShareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceVncUrlRequestParams struct {
	// 实例 ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceVncUrlRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	InstanceVncUrl *string `json:"InstanceVncUrl,omitnil,omitempty" name:"InstanceVncUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeInstancesDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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
	InstanceDeniedActionSet []*InstanceDeniedActions `json:"InstanceDeniedActionSet,omitnil,omitempty" name:"InstanceDeniedActionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例ID列表。每次请求批量实例的上限为 100。
	// 可通过 <a href="https://cloud.tencent.com/document/product/1207/47573">DescribeInstances</a> 接口返回值中的 InstanceId 获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeInstancesDiskNumRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表。每次请求批量实例的上限为 100。
	// 可通过 <a href="https://cloud.tencent.com/document/product/1207/47573">DescribeInstances</a> 接口返回值中的 InstanceId 获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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
	AttachDetailSet []*AttachDetail `json:"AttachDetailSet,omitnil,omitempty" name:"AttachDetailSet"`

	// 挂载信息数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

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
	// <li>tag-key</li>按照【标签键】进行过滤。
	// 类型：String
	// 必选：否
	// <li>tag-value</li>按照【标签值】进行过滤。
	// 类型：String
	// 必选：否
	// <li> tag:tag-key</li>按照【标签键值对】进行过滤。 tag-key使用具体的标签键进行替换。
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 100。参数不支持同时指定 InstanceIds 和 Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 指定排序字段 。取值范围： "EXPIRED_TIME"：依据实例的到期时间排序。 
	//  不传入此字段时, 优先返回实例状态为“待回收”的实例, 其余实例以“创建时间”倒序返回。
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 输出实例列表的排列顺序。取值范围：
	// "ASC"：升序排列。
	// "DESC"：降序排列。
	// 默认按升序排序。当传入该字段时，必须指定OrderField。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

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
	// <li>tag-key</li>按照【标签键】进行过滤。
	// 类型：String
	// 必选：否
	// <li>tag-value</li>按照【标签值】进行过滤。
	// 类型：String
	// 必选：否
	// <li> tag:tag-key</li>按照【标签键值对】进行过滤。 tag-key使用具体的标签键进行替换。
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 100。参数不支持同时指定 InstanceIds 和 Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 指定排序字段 。取值范围： "EXPIRED_TIME"：依据实例的到期时间排序。 
	//  不传入此字段时, 优先返回实例状态为“待回收”的实例, 其余实例以“创建时间”倒序返回。
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 输出实例列表的排列顺序。取值范围：
	// "ASC"：升序排列。
	// "DESC"：降序排列。
	// 默认按升序排序。当传入该字段时，必须指定OrderField。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
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
	delete(f, "OrderField")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例详细信息列表。
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
type DescribeInstancesReturnableRequestParams struct {
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeInstancesReturnableRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 可退还实例详细信息列表。
	InstanceReturnableSet []*InstanceReturnable `json:"InstanceReturnableSet,omitnil,omitempty" name:"InstanceReturnableSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeInstancesTrafficPackagesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例流量包详情列表。
	InstanceTrafficPackageSet []*InstanceTrafficPackage `json:"InstanceTrafficPackageSet,omitnil,omitempty" name:"InstanceTrafficPackageSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 密钥对 ID 列表。可通过[DescribeKeyPairs](https://cloud.tencent.com/document/product/1207/55540)接口返回值字段KeyPairSet获取。列表长度最大值为100。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤器列表。
	// <li>key-id</li>按照【密钥对ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>key-name</li>按照【密钥对名称】进行过滤（支持模糊匹配）。
	// 类型：String
	// 必选：否
	// <li>tag-key</li>
	// 按照【标签键】进行过滤。 类型：String 必选：否
	// <li>tag-value</li>
	// 按照【标签值】进行过滤。 类型：String 必选：否
	// <li>tag:tag-key</li>
	// 按照【标签键值对】进行过滤。 tag-key使用具体的标签键进行替换。
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。参数不支持同时指定 KeyIds 和 Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// 密钥对 ID 列表。可通过[DescribeKeyPairs](https://cloud.tencent.com/document/product/1207/55540)接口返回值字段KeyPairSet获取。列表长度最大值为100。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤器列表。
	// <li>key-id</li>按照【密钥对ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li>key-name</li>按照【密钥对名称】进行过滤（支持模糊匹配）。
	// 类型：String
	// 必选：否
	// <li>tag-key</li>
	// 按照【标签键】进行过滤。 类型：String 必选：否
	// <li>tag-value</li>
	// 按照【标签值】进行过滤。 类型：String 必选：否
	// <li>tag:tag-key</li>
	// 按照【标签键值对】进行过滤。 tag-key使用具体的标签键进行替换。
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。参数不支持同时指定 KeyIds 和 Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 密钥对详细信息列表。
	KeyPairSet []*KeyPair `json:"KeyPairSet,omitnil,omitempty" name:"KeyPairSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeMcpServerTemplatesRequestParams struct {
	// 过滤器列表。
	// <li>name-description</li>按照MCP Server模板名称或描述进行过滤（支持模糊匹配）。
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeMcpServerTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器列表。
	// <li>name-description</li>按照MCP Server模板名称或描述进行过滤（支持模糊匹配）。
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeMcpServerTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMcpServerTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMcpServerTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMcpServerTemplatesResponseParams struct {
	// MCP Server模板列表。
	McpServerTemplateSet []*McpServerTemplate `json:"McpServerTemplateSet,omitnil,omitempty" name:"McpServerTemplateSet"`

	// 符合条件的MCP Server模板数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMcpServerTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMcpServerTemplatesResponseParams `json:"Response"`
}

func (r *DescribeMcpServerTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMcpServerTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMcpServersRequestParams struct {
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// MCP Server ID列表。列表为空时此条件不生效。最大长度：10
	McpServerIds []*string `json:"McpServerIds,omitnil,omitempty" name:"McpServerIds"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeMcpServersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// MCP Server ID列表。列表为空时此条件不生效。最大长度：10
	McpServerIds []*string `json:"McpServerIds,omitnil,omitempty" name:"McpServerIds"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeMcpServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMcpServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "McpServerIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMcpServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMcpServersResponseParams struct {
	// MCP Server列表。
	McpServerSet []*McpServer `json:"McpServerSet,omitnil,omitempty" name:"McpServerSet"`

	// 符合条件的MCP Server数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例 ID。	
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMcpServersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMcpServersResponseParams `json:"Response"`
}

func (r *DescribeMcpServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMcpServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModifyInstanceBundlesRequestParams struct {
	// 实例 ID。可通过 <a href="https://cloud.tencent.com/document/product/1207/47573">DescribeInstances</a> 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 过滤器列表。
	// <li>bundle-id</li>按照【套餐 ID】进行过滤。
	// 类型：String
	// 必选：否
	// 可通过<a href="https://cloud.tencent.com/document/product/1207/47575"> DescribeBundles </a>接口返回值中的 BundleId 获取。
	// <li>support-platform-type</li>按照【系统类型】进行过滤。
	// 取值： LINUX_UNIX（Linux/Unix系统）；WINDOWS（Windows 系统）
	// 类型：String
	// 必选：否
	// <li>bundle-type</li>按照 【套餐类型进行过滤】。
	// 取值：GENERAL_BUNDLE (通用型套餐); STORAGE_BUNDLE(存储型套餐);ENTERPRISE_BUNDLE( 企业型套餐);EXCLUSIVE_BUNDLE(专属型套餐);BEFAST_BUNDLE(蜂驰型套餐);STARTER_BUNDLE(入门型套餐);ECONOMY_BUNDLE(经济型套餐);RAZOR_SPEED_BUNDLE(锐驰型套餐)
	// 类型：String
	// 必选：否
	// <li>bundle-state</li>按照【套餐状态】进行过滤。
	// 取值: ‘ONLINE’(在线); ‘OFFLINE’(下线);
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeModifyInstanceBundlesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 <a href="https://cloud.tencent.com/document/product/1207/47573">DescribeInstances</a> 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 过滤器列表。
	// <li>bundle-id</li>按照【套餐 ID】进行过滤。
	// 类型：String
	// 必选：否
	// 可通过<a href="https://cloud.tencent.com/document/product/1207/47575"> DescribeBundles </a>接口返回值中的 BundleId 获取。
	// <li>support-platform-type</li>按照【系统类型】进行过滤。
	// 取值： LINUX_UNIX（Linux/Unix系统）；WINDOWS（Windows 系统）
	// 类型：String
	// 必选：否
	// <li>bundle-type</li>按照 【套餐类型进行过滤】。
	// 取值：GENERAL_BUNDLE (通用型套餐); STORAGE_BUNDLE(存储型套餐);ENTERPRISE_BUNDLE( 企业型套餐);EXCLUSIVE_BUNDLE(专属型套餐);BEFAST_BUNDLE(蜂驰型套餐);STARTER_BUNDLE(入门型套餐);ECONOMY_BUNDLE(经济型套餐);RAZOR_SPEED_BUNDLE(锐驰型套餐)
	// 类型：String
	// 必选：否
	// <li>bundle-state</li>按照【套餐状态】进行过滤。
	// 取值: ‘ONLINE’(在线); ‘OFFLINE’(下线);
	// 类型：String
	// 必选：否
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 变更套餐详细信息。
	ModifyBundleSet []*ModifyBundle `json:"ModifyBundleSet,omitnil,omitempty" name:"ModifyBundleSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 地域信息列表。
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
type DescribeResetInstanceBlueprintsRequestParams struct {
	// 实例ID。可通过 <a href="https://cloud.tencent.com/document/product/1207/47573">DescribeInstances</a> 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤器列表。
	// <li>blueprint-id</li>按照【镜像 ID】进行过滤。
	// 类型：String
	// 必选：否
	// 可通过 <a href="https://cloud.tencent.com/document/product/1207/47689">DescribeBlueprints</a> 接口返回值中的 BlueprintId 获取。
	// <li>blueprint-type</li>按照【镜像类型】进行过滤。
	// 取值： APP_OS（应用镜像 ）；PURE_OS（ 系统镜像）；PRIVATE（自定义镜像）;DOCKER（Docker容器镜像）；SHARED（共享镜像）。
	// 类型：String
	// 必选：否
	// <li>platform-type</li>按照【镜像平台类型】进行过滤。
	// 取值： LINUX_UNIX（Linux/Unix系统）；WINDOWS（Windows 系统）。
	// 类型：String
	// 必选：否
	// <li>blueprint-name</li>按照【镜像名称】进行过滤。
	// 类型：String
	// 必选：否
	// 可通过 <a href="https://cloud.tencent.com/document/product/1207/47689">DescribeBlueprints</a> 接口返回值中的 BlueprintName 获取。
	// <li>blueprint-state</li>按照【镜像状态】进行过滤。
	// 类型：String
	// 必选：否
	// 可通过 <a href="https://cloud.tencent.com/document/product/1207/47689">DescribeBlueprints</a> 接口返回值中的 BlueprintState 获取。
	// 
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeResetInstanceBlueprintsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过 <a href="https://cloud.tencent.com/document/product/1207/47573">DescribeInstances</a> 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 偏移量，默认为 0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/product/1207/47578)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤器列表。
	// <li>blueprint-id</li>按照【镜像 ID】进行过滤。
	// 类型：String
	// 必选：否
	// 可通过 <a href="https://cloud.tencent.com/document/product/1207/47689">DescribeBlueprints</a> 接口返回值中的 BlueprintId 获取。
	// <li>blueprint-type</li>按照【镜像类型】进行过滤。
	// 取值： APP_OS（应用镜像 ）；PURE_OS（ 系统镜像）；PRIVATE（自定义镜像）;DOCKER（Docker容器镜像）；SHARED（共享镜像）。
	// 类型：String
	// 必选：否
	// <li>platform-type</li>按照【镜像平台类型】进行过滤。
	// 取值： LINUX_UNIX（Linux/Unix系统）；WINDOWS（Windows 系统）。
	// 类型：String
	// 必选：否
	// <li>blueprint-name</li>按照【镜像名称】进行过滤。
	// 类型：String
	// 必选：否
	// 可通过 <a href="https://cloud.tencent.com/document/product/1207/47689">DescribeBlueprints</a> 接口返回值中的 BlueprintName 获取。
	// <li>blueprint-state</li>按照【镜像状态】进行过滤。
	// 类型：String
	// 必选：否
	// 可通过 <a href="https://cloud.tencent.com/document/product/1207/47689">DescribeBlueprints</a> 接口返回值中的 BlueprintState 获取。
	// 
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 镜像重置信息列表
	ResetInstanceBlueprintSet []*ResetInstanceBlueprint `json:"ResetInstanceBlueprintSet,omitnil,omitempty" name:"ResetInstanceBlueprintSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 使用场景ID列表。可通过[DescribeScenes](https://cloud.tencent.com/document/product/1207/83512)接口返回值中的SceneId获取。
	SceneIds []*string `json:"SceneIds,omitnil,omitempty" name:"SceneIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeScenesRequest struct {
	*tchttp.BaseRequest
	
	// 使用场景ID列表。可通过[DescribeScenes](https://cloud.tencent.com/document/product/1207/83512)接口返回值中的SceneId获取。
	SceneIds []*string `json:"SceneIds,omitnil,omitempty" name:"SceneIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	SceneSet []*Scene `json:"SceneSet,omitnil,omitempty" name:"SceneSet"`

	// 使用场景总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

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
type DescribeSnapshotsDeniedActionsRequestParams struct {
	// 快照 ID 列表,每次请求批量快照的上限是100个。 可通过 <a href="https://cloud.tencent.com/document/product/1207/54388" target="_blank">DescribeSnapshots</a> 查询。
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`
}

type DescribeSnapshotsDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// 快照 ID 列表,每次请求批量快照的上限是100个。 可通过 <a href="https://cloud.tencent.com/document/product/1207/54388" target="_blank">DescribeSnapshots</a> 查询。
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`
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
	SnapshotDeniedActionSet []*SnapshotDeniedActions `json:"SnapshotDeniedActionSet,omitnil,omitempty" name:"SnapshotDeniedActionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 要查询快照的 ID 列表。每次请求批量快照的上限为 100。 
	// 可通过 [DescribeSnapshots](https://cloud.tencent.com/document/product/1207/54388) 接口返回值中的 SnapshotId		获取。
	// 参数不支持同时指定 SnapshotIds 和 Filters。
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`

	// 过滤器列表。
	// <li>snapshot-id</li>按照【快照 ID】进行过滤。
	// 类型：String
	// 必选：否
	// 可通过 <a href="https://cloud.tencent.com/document/product/1207/54388">DescribeSnapshots</a> 接口返回值中的 SnapshotId 获取。
	// <li>disk-id</li>按照【磁盘 ID】进行过滤。
	// 类型：String
	// 必选：否
	// 可通过 <a href="https://cloud.tencent.com/document/product/1207/66093">DescribeDisks</a> 接口返回值中的 DiskId 获取。
	// <li>snapshot-name</li>按照【快照名称】进行过滤。
	// 类型：String
	// 必选：否
	// 可通过 <a href="https://cloud.tencent.com/document/product/1207/54388">DescribeSnapshots</a> 接口返回值中的 SnapshotName 获取。
	// <li>instance-id</li>按照【实例 ID 】进行过滤。
	// 类型：String
	// 必选：否
	// 可通过 <a href="https://cloud.tencent.com/document/product/1207/47573">DescribeInstances</a> 接口返回值中的 InstanceId 获取。
	// <li>tag-key</li>
	// 按照【标签键】进行过滤。 类型：String 必选：否
	// <li>tag-value</li>
	// 按照【标签值】进行过滤。 类型：String 必选：否
	// <li>tag:tag-key</li>
	// 按照【标签键值对】进行过滤。 tag-key使用具体的标签键进行替换。
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。参数不支持同时指定 SnapshotIds 和 Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// 要查询快照的 ID 列表。每次请求批量快照的上限为 100。 
	// 可通过 [DescribeSnapshots](https://cloud.tencent.com/document/product/1207/54388) 接口返回值中的 SnapshotId		获取。
	// 参数不支持同时指定 SnapshotIds 和 Filters。
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`

	// 过滤器列表。
	// <li>snapshot-id</li>按照【快照 ID】进行过滤。
	// 类型：String
	// 必选：否
	// 可通过 <a href="https://cloud.tencent.com/document/product/1207/54388">DescribeSnapshots</a> 接口返回值中的 SnapshotId 获取。
	// <li>disk-id</li>按照【磁盘 ID】进行过滤。
	// 类型：String
	// 必选：否
	// 可通过 <a href="https://cloud.tencent.com/document/product/1207/66093">DescribeDisks</a> 接口返回值中的 DiskId 获取。
	// <li>snapshot-name</li>按照【快照名称】进行过滤。
	// 类型：String
	// 必选：否
	// 可通过 <a href="https://cloud.tencent.com/document/product/1207/54388">DescribeSnapshots</a> 接口返回值中的 SnapshotName 获取。
	// <li>instance-id</li>按照【实例 ID 】进行过滤。
	// 类型：String
	// 必选：否
	// 可通过 <a href="https://cloud.tencent.com/document/product/1207/47573">DescribeInstances</a> 接口返回值中的 InstanceId 获取。
	// <li>tag-key</li>
	// 按照【标签键】进行过滤。 类型：String 必选：否
	// <li>tag-value</li>
	// 按照【标签值】进行过滤。 类型：String 必选：否
	// <li>tag:tag-key</li>
	// 按照【标签键值对】进行过滤。 tag-key使用具体的标签键进行替换。
	// 每次请求的 Filters 的上限为 10，Filter.Values 的上限为 5。参数不支持同时指定 SnapshotIds 和 Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 快照的详情列表。
	SnapshotSet []*Snapshot `json:"SnapshotSet,omitnil,omitempty" name:"SnapshotSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// <li>ZONE：依据可用区排序。</li>
	// <li>INSTANCE_DISPLAY_LABEL：依据可用区展示标签排序，可用区展示标签包括：HIDDEN（隐藏）、NORMAL（普通）、SELECTED（默认选中），默认采用的升序排列为：['HIDDEN', 'NORMAL', 'SELECTED']。
	// 默认按可用区排序。</li>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 输出可用区列表的排列顺序。取值范围：
	// <li>ASC：升序排列。 </li>
	// <li>DESC：降序排列。</li>
	// 默认按升序排列。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
	// 可用区列表排序的依据字段。取值范围：
	// <li>ZONE：依据可用区排序。</li>
	// <li>INSTANCE_DISPLAY_LABEL：依据可用区展示标签排序，可用区展示标签包括：HIDDEN（隐藏）、NORMAL（普通）、SELECTED（默认选中），默认采用的升序排列为：['HIDDEN', 'NORMAL', 'SELECTED']。
	// 默认按可用区排序。</li>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 输出可用区列表的排列顺序。取值范围：
	// <li>ASC：升序排列。 </li>
	// <li>DESC：降序排列。</li>
	// 默认按升序排列。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 可用区详细信息列表
	ZoneInfoSet []*ZoneInfo `json:"ZoneInfoSet,omitnil,omitempty" name:"ZoneInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type DestinationRegionBlueprint struct {
	// 目标地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 目标地域镜像ID。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`
}

// Predefined struct for user
type DetachCcnRequestParams struct {
	// 云联网实例ID。可通过[DescribeCcnAttachedInstances](https://cloud.tencent.com/document/product/1207/58797)接口返回值中的CcnId获取。
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`
}

type DetachCcnRequest struct {
	*tchttp.BaseRequest
	
	// 云联网实例ID。可通过[DescribeCcnAttachedInstances](https://cloud.tencent.com/document/product/1207/58797)接口返回值中的CcnId获取。
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 云硬盘ID列表。每次批量请求云硬盘的上限为 100。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`
}

type DetachDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。每次批量请求云硬盘的上限为 100。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// <li>"DiskBackupQuota"代表数据盘备份点配额收费项。</li>
	// <li>"Instance"代表实例收费项。</li>
	// <li>"SystemDiskBackupQuota"代表系统盘备份点配额收费项。</li>
	PriceName *string `json:"PriceName,omitnil,omitempty" name:"PriceName"`

	// 计费项维度单价。
	OriginUnitPrice *float64 `json:"OriginUnitPrice,omitnil,omitempty" name:"OriginUnitPrice"`

	// 计费项维度总价。
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 计费项维度折扣。
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 计费项维度折后总价。
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`
}

// Predefined struct for user
type DisassociateInstancesKeyPairsRequestParams struct {
	// 密钥对 ID 列表，每次请求批量密钥对的上限为 100。可通过[DescribeKeyPairs](https://cloud.tencent.com/document/api/1207/55540)接口返回值中的KeyId获取。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DisassociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// 密钥对 ID 列表，每次请求批量密钥对的上限为 100。可通过[DescribeKeyPairs](https://cloud.tencent.com/document/api/1207/55540)接口返回值中的KeyId获取。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 时间单位。
	// 取值为：
	// - m - 月
	// - d - 日
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 总价。
	TotalCost *float64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 折后总价。
	RealTotalCost *float64 `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 折扣。
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 具体折扣详情。
	PolicyDetail *PolicyDetail `json:"PolicyDetail,omitnil,omitempty" name:"PolicyDetail"`
}

type Disk struct {
	// 云硬盘ID。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 云硬盘名称。
	DiskName *string `json:"DiskName,omitnil,omitempty" name:"DiskName"`

	// 云硬盘类型。
	// 枚举值：
	// <li> SYSTEM_DISK: 系统盘 </li>
	// <li> DATA_DISK: 数据盘 </li>
	DiskUsage *string `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// 云硬盘介质类型。
	// 枚举值:
	// <li> CLOUD_BASIC: 普通云硬盘 </li>
	// <li> CLOUD_PREMIUM: 高性能云硬盘 </li>
	// <li> CLOUD_SSD: SSD云硬盘 </li>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 云硬盘付费类型。
	// <li> PREPAID: 预付费 </li>
	// <li> POSTPAID_BY_HOUR: 按小时后付费 </li>
	DiskChargeType *string `json:"DiskChargeType,omitnil,omitempty" name:"DiskChargeType"`

	// 云硬盘大小, 单位GB。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 续费标识。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 云硬盘状态，取值范围：
	// <li>PENDING：创建中。 </li>
	// <li>UNATTACHED：待挂载。</li>
	// <li>ATTACHING：挂载中。</li>
	// <li>ATTACHED：已挂载。</li>
	// <li>DETACHING：卸载中。 </li>
	// <li> SHUTDOWN：已隔离。</li>
	// <li> CREATED_FAILED：创建失败。</li>
	// <li>TERMINATING：销毁中。</li>
	// <li> DELETING：删除中。</li>
	// <li> FREEZING：冻结中。</li>
	DiskState *string `json:"DiskState,omitnil,omitempty" name:"DiskState"`

	// 云硬盘挂载状态。
	Attached *bool `json:"Attached,omitnil,omitempty" name:"Attached"`

	// 是否随实例释放。
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitnil,omitempty" name:"DeleteWithInstance"`

	// 上一次操作。
	LatestOperation *string `json:"LatestOperation,omitnil,omitempty" name:"LatestOperation"`

	// 上一次操作状态。
	LatestOperationState *string `json:"LatestOperationState,omitnil,omitempty" name:"LatestOperationState"`

	// 上一次请求ID。
	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitnil,omitempty" name:"LatestOperationRequestId"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 到期时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 隔离时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolatedTime *string `json:"IsolatedTime,omitnil,omitempty" name:"IsolatedTime"`

	// 云硬盘的已有备份点数量。
	DiskBackupCount *int64 `json:"DiskBackupCount,omitnil,omitempty" name:"DiskBackupCount"`

	// 云硬盘的备份点配额数量。
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitnil,omitempty" name:"DiskBackupQuota"`

	// 云硬盘绑定的标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DiskBackup struct {
	// 云硬盘备份点ID。
	DiskBackupId *string `json:"DiskBackupId,omitnil,omitempty" name:"DiskBackupId"`

	// 创建此云硬盘备份点的云硬盘类型。取值：<li>DATA_DISK：数据盘</li>
	DiskUsage *string `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// 创建此云硬盘备份点的云硬盘 ID。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 创建此云硬盘备份点的云硬盘大小，单位 GB。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 云硬盘备份点名称，用户自定义的云硬盘备份点别名。
	DiskBackupName *string `json:"DiskBackupName,omitnil,omitempty" name:"DiskBackupName"`

	// 云硬盘备份点的状态。取值范围：
	// <li>NORMAL：正常。 </li>
	// <li>CREATING：创建中。</li>
	// <li>ROLLBACKING：回滚中。</li>
	// <li>DELETING：删除中。</li>
	DiskBackupState *string `json:"DiskBackupState,omitnil,omitempty" name:"DiskBackupState"`

	// 创建或回滚云硬盘备份点进度百分比，成功后此字段取值为 100。
	Percent *int64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 上一次操作
	LatestOperation *string `json:"LatestOperation,omitnil,omitempty" name:"LatestOperation"`

	// 上一次操作状态
	LatestOperationState *string `json:"LatestOperationState,omitnil,omitempty" name:"LatestOperationState"`

	// 上一次请求ID
	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitnil,omitempty" name:"LatestOperationRequestId"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 云硬盘备份点绑定的标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DiskBackupDeniedActions struct {
	// 云硬盘备份点ID。
	DiskBackupId *string `json:"DiskBackupId,omitnil,omitempty" name:"DiskBackupId"`

	// 操作限制列表。
	DeniedActions []*DeniedAction `json:"DeniedActions,omitnil,omitempty" name:"DeniedActions"`
}

type DiskChargePrepaid struct {
	// 新购周期。
	// 可选值：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 自动续费标识。取值范围：
	// - NOTIFY_AND_AUTO_RENEW：通知过期且自动续费。
	// - NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费，用户需要手动续费。
	// - DISABLE_NOTIFY_AND_MANUAL_RENEW：不自动续费，且不通知。
	// 
	// 默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，云硬盘到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 新购单位.。
	// 可选值：m - 月。
	// 默认值：m - 月。
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
}

type DiskConfig struct {
	// 可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 云硬盘类型。枚举值如下：
	// 
	// <li>CLOUD_BASIC：普通云硬盘</li>
	// <li>CLOUD_PREMIUM：高性能云硬盘</li>
	// <li>CLOUD_SSD：SSD云硬盘</li>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 云硬盘可售卖状态。
	DiskSalesState *string `json:"DiskSalesState,omitnil,omitempty" name:"DiskSalesState"`

	// 最大云硬盘大小。
	MaxDiskSize *int64 `json:"MaxDiskSize,omitnil,omitempty" name:"MaxDiskSize"`

	// 最小云硬盘大小。
	MinDiskSize *int64 `json:"MinDiskSize,omitnil,omitempty" name:"MinDiskSize"`

	// 云硬盘步长。
	DiskStepSize *int64 `json:"DiskStepSize,omitnil,omitempty" name:"DiskStepSize"`
}

type DiskDeniedActions struct {
	// 云硬盘ID。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 操作限制列表。
	DeniedActions []*DeniedAction `json:"DeniedActions,omitnil,omitempty" name:"DeniedActions"`
}

type DiskPrice struct {
	// 云硬盘单价。
	OriginalDiskPrice *float64 `json:"OriginalDiskPrice,omitnil,omitempty" name:"OriginalDiskPrice"`

	// 云硬盘总价。
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 折扣。
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 折后总价。
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// 计费项目明细列表。
	DetailPrices []*DetailPrice `json:"DetailPrices,omitnil,omitempty" name:"DetailPrices"`
}

type DiskReturnable struct {
	// 云硬盘ID。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 云硬盘是否可退还。
	IsReturnable *bool `json:"IsReturnable,omitnil,omitempty" name:"IsReturnable"`

	// 云硬盘退还失败错误码。
	ReturnFailCode *int64 `json:"ReturnFailCode,omitnil,omitempty" name:"ReturnFailCode"`

	// 云硬盘退还失败错误信息。
	ReturnFailMessage *string `json:"ReturnFailMessage,omitnil,omitempty" name:"ReturnFailMessage"`
}

type DockerActivity struct {
	// 活动ID。
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 活动名称。
	ActivityName *string `json:"ActivityName,omitnil,omitempty" name:"ActivityName"`

	// 活动状态。取值范围： 
	// <li>INIT：表示初始化，活动尚未执行</li>
	// <li>OPERATING：表示活动执行中</li>
	// <li>SUCCESS：表示活动执行成功</li>
	// <li>FAILED：表示活动执行失败</li>
	ActivityState *string `json:"ActivityState,omitnil,omitempty" name:"ActivityState"`

	// 活动执行的命令输出，以base64编码。
	ActivityCommandOutput *string `json:"ActivityCommandOutput,omitnil,omitempty" name:"ActivityCommandOutput"`

	// 容器ID列表。
	ContainerIds []*string `json:"ContainerIds,omitnil,omitempty" name:"ContainerIds"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 结束时间。按照 ISO8601 标准表示，并且使用 UTC 时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DockerContainer struct {
	// 容器ID
	ContainerId *string `json:"ContainerId,omitnil,omitempty" name:"ContainerId"`

	// 容器名称
	ContainerName *string `json:"ContainerName,omitnil,omitempty" name:"ContainerName"`

	// 容器镜像地址
	ContainerImage *string `json:"ContainerImage,omitnil,omitempty" name:"ContainerImage"`

	// 容器Command
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 容器状态描述
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 容器状态，和docker的容器状态保持一致，当前取值有：created（已创建）、restarting（重启中）、running（运行中）、removing（迁移中）、paused（暂停）、exited（停止）和dead（死亡）
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 容器端口主机端口映射列表
	PublishPortSet []*DockerContainerPublishPort `json:"PublishPortSet,omitnil,omitempty" name:"PublishPortSet"`

	// 容器挂载本地卷列表
	VolumeSet []*DockerContainerVolume `json:"VolumeSet,omitnil,omitempty" name:"VolumeSet"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`
}

type DockerContainerConfiguration struct {
	// 容器镜像地址
	ContainerImage *string `json:"ContainerImage,omitnil,omitempty" name:"ContainerImage"`

	// 容器名称
	ContainerName *string `json:"ContainerName,omitnil,omitempty" name:"ContainerName"`

	// 环境变量列表
	Envs []*ContainerEnv `json:"Envs,omitnil,omitempty" name:"Envs"`

	// 容器端口主机端口映射列表
	PublishPorts []*DockerContainerPublishPort `json:"PublishPorts,omitnil,omitempty" name:"PublishPorts"`

	// 容器加载本地卷列表
	Volumes []*DockerContainerVolume `json:"Volumes,omitnil,omitempty" name:"Volumes"`

	// 运行的命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 容器重启策略。
	// - no -默认策略，在容器退出时不重启容器
	// - on-failure -在容器非正常退出时（退出状态非0），才会重启容器
	// - on-failure:3 -在容器非正常退出时重启容器，最多重启3次
	// - always -在容器退出时总是重启容器
	RestartPolicy *string `json:"RestartPolicy,omitnil,omitempty" name:"RestartPolicy"`
}

type DockerContainerPublishPort struct {
	// 主机端口
	HostPort *int64 `json:"HostPort,omitnil,omitempty" name:"HostPort"`

	// 容器端口
	ContainerPort *int64 `json:"ContainerPort,omitnil,omitempty" name:"ContainerPort"`

	// 对外绑定IP，默认0.0.0.0
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 协议，默认tcp，支持tcp/udp/sctp
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

type DockerContainerVolume struct {
	// 容器路径
	ContainerPath *string `json:"ContainerPath,omitnil,omitempty" name:"ContainerPath"`

	// 主机路径
	HostPath *string `json:"HostPath,omitnil,omitempty" name:"HostPath"`
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FirewallRule struct {
	// 协议，取值：TCP，UDP，ICMP，ALL，ICMPv6。
	// 
	// - 使用ICMP协议时，只支持CidrBlock，不支持使用Port、Ipv6CidrBlock参数；
	// - 使用ICMPv6协议时，只支持Ipv6CidrBlock，不支持使用Port、Ipv6CidrBlock参数；
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 端口，取值：ALL，单独的端口，逗号分隔的离散端口，减号分隔的端口范围。注意：单独的端口与离散端口不能同时存在。
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// IPv4网段或 IPv4地址(互斥)。
	// 示例值：0.0.0.0/0。
	// 
	// 和Ipv6CidrBlock互斥，两者都不指定时，如果Protocol不是ICMPv6，则取默认值0.0.0.0/0。
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// IPv6网段或IPv6地址(互斥)。
	// 示例值：::/0。
	// 
	// 和CidrBlock互斥，两者都不指定时，如果Protocol是ICMPv6，则取默认值::/0。
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil,omitempty" name:"Ipv6CidrBlock"`

	// 取值：ACCEPT（允许），DROP（拒绝）。默认为 ACCEPT。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 防火墙规则描述。
	FirewallRuleDescription *string `json:"FirewallRuleDescription,omitnil,omitempty" name:"FirewallRuleDescription"`
}

type FirewallRuleInfo struct {
	// 应用类型，取值：自定义，HTTP(80)，HTTPS(443)，Linux登录(22)，Windows登录(3389)，MySQL(3306)，SQL Server(1433)，全部TCP，全部UDP，Ping-ICMP，Windows登录优化 (3389)，FTP (21)，Ping，Ping (IPv6)，ALL。
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 协议，取值：TCP，UDP，ICMP，ICMPv6，ALL。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 端口，取值：ALL，单独的端口，逗号分隔的离散端口，减号分隔的端口范围。
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// IPv4网段或 IPv4地址(互斥)。
	// 示例值：0.0.0.0/0。
	// 
	// 和Ipv6CidrBlock互斥，两者都不指定时，如果Protocol不是ICMPv6，则取默认值0.0.0.0/0。
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// IPv6网段或IPv6地址(互斥)。
	// 示例值：::/0。
	// 
	// 和CidrBlock互斥，两者都不指定时，如果Protocol是ICMPv6，则取默认值::/0。
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil,omitempty" name:"Ipv6CidrBlock"`

	// 取值：ACCEPT，DROP。默认为 ACCEPT。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 防火墙规则描述。
	FirewallRuleDescription *string `json:"FirewallRuleDescription,omitnil,omitempty" name:"FirewallRuleDescription"`
}

type FirewallTemplate struct {
	// 模板ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 模板类型。取值: "PRIVATE"(个人模板)
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// 模板状态。取值: "NORMAL"(正常)
	TemplateState *string `json:"TemplateState,omitnil,omitempty" name:"TemplateState"`

	// 模板创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`
}

type FirewallTemplateApplyRecord struct {
	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 应用模板的时间。
	ApplyTime *string `json:"ApplyTime,omitnil,omitempty" name:"ApplyTime"`

	// 模板规则列表。
	TemplateRuleSet []*FirewallTemplateRule `json:"TemplateRuleSet,omitnil,omitempty" name:"TemplateRuleSet"`

	// 应用模板的执行状态。
	// 
	// - SUCCESS：成功
	// - RUNNING：运行中
	// - FAILED：失败
	ApplyState *string `json:"ApplyState,omitnil,omitempty" name:"ApplyState"`

	// 应用成功的实例数量。
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 应用失败的实例数量。
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 正在应用中的实例数量。
	RunningCount *int64 `json:"RunningCount,omitnil,omitempty" name:"RunningCount"`

	// 应用模板的执行细节。
	ApplyDetailSet []*FirewallTemplateApplyRecordDetail `json:"ApplyDetailSet,omitnil,omitempty" name:"ApplyDetailSet"`
}

type FirewallTemplateApplyRecordDetail struct {
	// 实例标识信息。
	Instance *InstanceIdentifier `json:"Instance,omitnil,omitempty" name:"Instance"`

	// 防火墙模板应用状态。
	// 
	// - SUCCESS：成功
	// - FAILED：失败
	// - RUNNING：运行中
	ApplyState *string `json:"ApplyState,omitnil,omitempty" name:"ApplyState"`

	// 防火墙模板应用错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`
}

type FirewallTemplateRule struct {
	// 防火墙模板规则ID。
	TemplateRuleId *string `json:"TemplateRuleId,omitnil,omitempty" name:"TemplateRuleId"`

	// 防火墙规则。
	FirewallRule *FirewallRule `json:"FirewallRule,omitnil,omitempty" name:"FirewallRule"`
}

type FirewallTemplateRuleInfo struct {
	// 防火墙模板规则ID。
	TemplateRuleId *string `json:"TemplateRuleId,omitnil,omitempty" name:"TemplateRuleId"`

	// 防火墙规则信息。
	FirewallRuleInfo *FirewallRuleInfo `json:"FirewallRuleInfo,omitnil,omitempty" name:"FirewallRuleInfo"`
}

type GeneralResourceQuota struct {
	// 资源名称。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 资源当前可用数量。
	ResourceQuotaAvailable *int64 `json:"ResourceQuotaAvailable,omitnil,omitempty" name:"ResourceQuotaAvailable"`

	// 资源总数量。
	ResourceQuotaTotal *int64 `json:"ResourceQuotaTotal,omitnil,omitempty" name:"ResourceQuotaTotal"`
}

type Image struct {
	// CVM镜像 ID ，是Image的唯一标识。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 镜像名称。
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 镜像描述。
	ImageDescription *string `json:"ImageDescription,omitnil,omitempty" name:"ImageDescription"`

	// 镜像大小。单位GB。
	ImageSize *int64 `json:"ImageSize,omitnil,omitempty" name:"ImageSize"`

	// 镜像来源。
	// <li>CREATE_IMAGE：自定义镜像</li>
	// <li>EXTERNAL_IMPORT：外部导入镜像</li>
	ImageSource *string `json:"ImageSource,omitnil,omitempty" name:"ImageSource"`

	// 镜像分类
	// <li>SystemImage：系统盘镜像</li>
	// <li>InstanceImage：整机镜像</li>
	ImageClass *string `json:"ImageClass,omitnil,omitempty" name:"ImageClass"`

	// 镜像状态
	// CREATING:创建中
	// NORMAL:正常
	// CREATEFAILED:创建失败
	// USING:使用中
	// SYNCING:同步中
	// IMPORTING:导入中
	// IMPORTFAILED:导入失败
	ImageState *string `json:"ImageState,omitnil,omitempty" name:"ImageState"`

	// 镜像是否支持Cloudinit。
	IsSupportCloudinit *bool `json:"IsSupportCloudinit,omitnil,omitempty" name:"IsSupportCloudinit"`

	// 镜像架构。
	Architecture *string `json:"Architecture,omitnil,omitempty" name:"Architecture"`

	// 镜像操作系统。
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// 镜像来源平台。
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 镜像创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 镜像是否可共享到轻量应用服务器。
	IsShareable *bool `json:"IsShareable,omitnil,omitempty" name:"IsShareable"`

	// 不可共享的原因。
	UnshareableReason *string `json:"UnshareableReason,omitnil,omitempty" name:"UnshareableReason"`
}

// Predefined struct for user
type ImportKeyPairRequestParams struct {
	// 密钥对名称，可由数字，字母和下划线组成，长度不超过 25 个字符。
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// 密钥对的公钥内容， OpenSSH RSA 格式。
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`

	// 标签键和标签值。 如果指定多个标签，则会为指定资源同时创建并绑定该多个标签。 同一个资源上的同一个标签键只能对应一个标签值。如果您尝试添加已有标签键，则对应的标签值会更新为新值。 如果标签不存在会为您自动创建标签。 数组最多支持10个元素。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ImportKeyPairRequest struct {
	*tchttp.BaseRequest
	
	// 密钥对名称，可由数字，字母和下划线组成，长度不超过 25 个字符。
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// 密钥对的公钥内容， OpenSSH RSA 格式。
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`

	// 标签键和标签值。 如果指定多个标签，则会为指定资源同时创建并绑定该多个标签。 同一个资源上的同一个标签键只能对应一个标签值。如果您尝试添加已有标签键，则对应的标签值会更新为新值。 如果标签不存在会为您自动创建标签。 数组最多支持10个元素。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportKeyPairRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportKeyPairResponseParams struct {
	// 密钥对 ID。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	BlueprintCount *int64 `json:"BlueprintCount,omitnil,omitempty" name:"BlueprintCount"`
}

type InquirePriceCreateBlueprintRequest struct {
	*tchttp.BaseRequest
	
	// 自定义镜像的个数。默认值为1。
	BlueprintCount *int64 `json:"BlueprintCount,omitnil,omitempty" name:"BlueprintCount"`
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
	BlueprintPrice *BlueprintPrice `json:"BlueprintPrice,omitnil,omitempty" name:"BlueprintPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 云硬盘介质类型。取值: "CLOUD_PREMIUM"(高性能云盘), "CLOUD_SSD"(SSD云硬盘)。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 新购云硬盘包年包月相关参数设置。
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitnil,omitempty" name:"DiskChargePrepaid"`

	// 云硬盘个数, 默认值: 1。
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// 指定云硬盘备份点配额，不传时默认为不带备份点配额。
	// 取值范围：0 到 500
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitnil,omitempty" name:"DiskBackupQuota"`
}

type InquirePriceCreateDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘大小, 单位: GB。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 云硬盘介质类型。取值: "CLOUD_PREMIUM"(高性能云盘), "CLOUD_SSD"(SSD云硬盘)。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 新购云硬盘包年包月相关参数设置。
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitnil,omitempty" name:"DiskChargePrepaid"`

	// 云硬盘个数, 默认值: 1。
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// 指定云硬盘备份点配额，不传时默认为不带备份点配额。
	// 取值范围：0 到 500
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitnil,omitempty" name:"DiskBackupQuota"`
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
	DiskPrice *DiskPrice `json:"DiskPrice,omitnil,omitempty" name:"DiskPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例的套餐 ID。可以通过调用[DescribeBundles](https://cloud.tencent.com/document/api/1207/47575)接口获取。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 创建数量，默认为 1。
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 应用镜像 ID，使用收费应用镜像时必填。可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值中的BlueprintId获取。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`
}

type InquirePriceCreateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例的套餐 ID。可以通过调用[DescribeBundles](https://cloud.tencent.com/document/api/1207/47575)接口获取。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 创建数量，默认为 1。
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 应用镜像 ID，使用收费应用镜像时必填。可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值中的BlueprintId获取。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`
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
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 云硬盘ID列表。每次批量请求云硬盘的上限为 1。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 续费云硬盘包年包月相关参数设置。
	RenewDiskChargePrepaid *RenewDiskChargePrepaid `json:"RenewDiskChargePrepaid,omitnil,omitempty" name:"RenewDiskChargePrepaid"`
}

type InquirePriceRenewDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。每次批量请求云硬盘的上限为 1。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 续费云硬盘包年包月相关参数设置。
	RenewDiskChargePrepaid *RenewDiskChargePrepaid `json:"RenewDiskChargePrepaid,omitnil,omitempty" name:"RenewDiskChargePrepaid"`
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
	DiskPrice *DiskPrice `json:"DiskPrice,omitnil,omitempty" name:"DiskPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 是否续费数据盘。默认值: false, 即不续费。
	RenewDataDisk *bool `json:"RenewDataDisk,omitnil,omitempty" name:"RenewDataDisk"`

	// 数据盘是否对齐实例到期时间。默认值: false, 即不对齐。
	AlignInstanceExpiredTime *bool `json:"AlignInstanceExpiredTime,omitnil,omitempty" name:"AlignInstanceExpiredTime"`
}

type InquirePriceRenewInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 待续费的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573 )接口返回值中的InstanceId获取。每次请求批量实例的上限为50。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 是否续费数据盘。默认值: false, 即不续费。
	RenewDataDisk *bool `json:"RenewDataDisk,omitnil,omitempty" name:"RenewDataDisk"`

	// 数据盘是否对齐实例到期时间。默认值: false, 即不对齐。
	AlignInstanceExpiredTime *bool `json:"AlignInstanceExpiredTime,omitnil,omitempty" name:"AlignInstanceExpiredTime"`
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
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// 数据盘价格信息列表。
	DataDiskPriceSet []*DataDiskPrice `json:"DataDiskPriceSet,omitnil,omitempty" name:"DataDiskPriceSet"`

	// 待续费实例价格列表。
	InstancePriceDetailSet []*InstancePriceDetail `json:"InstancePriceDetailSet,omitnil,omitempty" name:"InstancePriceDetailSet"`

	// 总计价格。
	TotalPrice *TotalPrice `json:"TotalPrice,omitnil,omitempty" name:"TotalPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 套餐 ID。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 镜像 ID。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`

	// 实例的 CPU 核数，单位：核。
	CPU *int64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// 实例内存容量，单位：GB 。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例计费模式。取值范围： 
	// PREPAID：表示预付费，即包年包月。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例系统盘信息。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例主网卡的内网 IP。 
	// 注意：此字段可能返回 空，表示取不到有效值。
	PrivateAddresses []*string `json:"PrivateAddresses,omitnil,omitempty" name:"PrivateAddresses"`

	// 实例主网卡的公网 IP。 
	// 注意：此字段可能返回 空，表示取不到有效值。
	PublicAddresses []*string `json:"PublicAddresses,omitnil,omitempty" name:"PublicAddresses"`

	// 实例带宽信息。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 自动续费标识。取值范围： 
	// NOTIFY_AND_MANUAL_RENEW：表示通知即将过期，但不自动续费  
	// NOTIFY_AND_AUTO_RENEW：表示通知即将过期，而且自动续费 DISABLE_NOTIFY_AND_MANUAL_RENEW：不自动续费，且不通知。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 实例登录设置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例状态。取值范围： 
	// <li>PENDING：表示创建中</li><li>LAUNCH_FAILED：表示创建失败</li><li>RUNNING：表示运行中</li><li>STOPPED：表示关机</li><li>STARTING：表示开机中</li><li>STOPPING：表示关机中</li><li>REBOOTING：表示重启中</li><li>SHUTDOWN：表示停止待销毁</li><li>TERMINATING：表示销毁中</li><li>DELETING：表示删除中</li><li>FREEZING：表示冻结中</li><li>ENTER_RESCUE_MODE：表示进入救援模式中</li><li>RESCUE_MODE：表示救援模式</li><li>EXIT_RESCUE_MODE：表示退出救援模式中</li>
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// 实例全局唯一 ID。
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// 实例的最新操作。例：StopInstances、ResetInstance。注意：此字段可能返回 空值，表示取不到有效值。
	LatestOperation *string `json:"LatestOperation,omitnil,omitempty" name:"LatestOperation"`

	// 实例的最新操作状态。取值范围： 
	// SUCCESS：表示操作成功 
	// OPERATING：表示操作执行中 
	// FAILED：表示操作失败 
	// 注意：此字段可能返回 空值，表示取不到有效值。
	LatestOperationState *string `json:"LatestOperationState,omitnil,omitempty" name:"LatestOperationState"`

	// 实例最新操作的唯一请求 ID。 
	// 注意：此字段可能返回 空值，表示取不到有效值。
	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitnil,omitempty" name:"LatestOperationRequestId"`

	// 实例最新操作的开始时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperationStartedTime *string `json:"LatestOperationStartedTime,omitnil,omitempty" name:"LatestOperationStartedTime"`

	// 隔离时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolatedTime *string `json:"IsolatedTime,omitnil,omitempty" name:"IsolatedTime"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 到期时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 操作系统平台类型，如 LINUX_UNIX、WINDOWS。
	PlatformType *string `json:"PlatformType,omitnil,omitempty" name:"PlatformType"`

	// 操作系统平台。
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 操作系统名称。
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// 可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例绑定的标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 实例封禁状态。取值范围：
	// <li>NORMAL实例正常。</li><li>NETWORK_RESTRICT：网络封禁。</li>
	InstanceRestrictState *string `json:"InstanceRestrictState,omitnil,omitempty" name:"InstanceRestrictState"`

	// 描述实例是否支持IPv6。
	SupportIpv6Detail *SupportIpv6Detail `json:"SupportIpv6Detail,omitnil,omitempty" name:"SupportIpv6Detail"`

	// 公网IPv6地址列表。
	PublicIpv6Addresses []*string `json:"PublicIpv6Addresses,omitnil,omitempty" name:"PublicIpv6Addresses"`

	// 创建实例后自动执行TAT命令的调用ID。
	InitInvocationId *string `json:"InitInvocationId,omitnil,omitempty" name:"InitInvocationId"`

	// 实例违规详情。
	InstanceViolationDetail *InstanceViolationDetail `json:"InstanceViolationDetail,omitnil,omitempty" name:"InstanceViolationDetail"`
}

type InstanceChargePrepaid struct {
	// 购买实例的时长，单位：月。
	// - 创建实例时，取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60。
	// - 续费实例时，取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li><br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费，用户需要手动续费</li><br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不自动续费，且不通知</li><br><br>默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type InstanceDeniedActions struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 操作限制列表。
	DeniedActions []*DeniedAction `json:"DeniedActions,omitnil,omitempty" name:"DeniedActions"`
}

type InstanceIdentifier struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type InstancePrice struct {
	// 套餐单价原价。
	OriginalBundlePrice *float64 `json:"OriginalBundlePrice,omitnil,omitempty" name:"OriginalBundlePrice"`

	// 原价。
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 折扣。
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 折后价。
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// 价格货币单位。取值范围CNY:人民币。USD:美元。
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 计费项目明细。
	DetailPrices []*DetailPrice `json:"DetailPrices,omitnil,omitempty" name:"DetailPrices"`
}

type InstancePriceDetail struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 询价信息。
	InstancePrice *InstancePrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// 折扣梯度详情，每个梯度包含的信息有：时长，折扣数，总价，折扣价，折扣详情（用户折扣、官网折扣、最终折扣）。
	DiscountDetail []*DiscountDetail `json:"DiscountDetail,omitnil,omitempty" name:"DiscountDetail"`
}

type InstanceReturnable struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例是否可退还。
	IsReturnable *bool `json:"IsReturnable,omitnil,omitempty" name:"IsReturnable"`

	// 实例退还失败错误码。取值:
	// 0: 可以退还
	// 1: 资源已退货。
	// 2: 资源已到期
	// 3: 资源购买超过5天不支持退款
	// 4: 非预付费资源不支持退款
	// 8: 退货数量超出限额
	// 9: 涉及活动订单不支持退款
	// 10: 资源不支持自助退，请走工单退款
	// 11: 涉及推广奖励渠道订单，请提工单咨询
	// 12: 根据业务侧产品规定，该资源不允许退款
	ReturnFailCode *int64 `json:"ReturnFailCode,omitnil,omitempty" name:"ReturnFailCode"`

	// 实例退还失败错误信息。
	ReturnFailMessage *string `json:"ReturnFailMessage,omitnil,omitempty" name:"ReturnFailMessage"`
}

type InstanceTrafficPackage struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 流量包详情列表。
	TrafficPackageSet []*TrafficPackage `json:"TrafficPackageSet,omitnil,omitempty" name:"TrafficPackageSet"`
}

type InstanceViolationDetail struct {
	//  来源：RESTRICT：封禁、FREEZW：冻结
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 是否允许自助解封：1是，2否
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 违规类型
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 违规内容（URL、关联域名）
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type InternetAccessible struct {
	// 网络计费类型，取值范围：
	// <li>按流量包付费：TRAFFIC_POSTPAID_BY_HOUR</li>
	// <li>按带宽付费： BANDWIDTH_POSTPAID_BY_HOUR</li>
	InternetChargeType *string `json:"InternetChargeType,omitnil,omitempty" name:"InternetChargeType"`

	// 公网出带宽上限，单位：Mbps。
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 是否分配公网 IP。
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitnil,omitempty" name:"PublicIpAssigned"`
}

// Predefined struct for user
type IsolateDisksRequestParams struct {
	// 云硬盘ID列表。一个或多个待操作的云硬盘ID。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。每次请求退还数据盘数量总计上限为20。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`
}

type IsolateDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。一个或多个待操作的云硬盘ID。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。每次请求退还数据盘数量总计上限为20。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`
}

func (r *IsolateDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDisksResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateDisksResponse struct {
	*tchttp.BaseResponse
	Response *IsolateDisksResponseParams `json:"Response"`
}

func (r *IsolateDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateInstancesRequestParams struct {
	// 实例ID列表。一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。每次请求退还实例和数据盘数量总计上限为20。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 是否退还挂载的数据盘。取值范围：
	// TRUE：表示退还实例同时退还其挂载的数据盘。
	// FALSE：表示退还实例同时不再退还其挂载的数据盘。
	// 默认取值：TRUE。
	IsolateDataDisk *bool `json:"IsolateDataDisk,omitnil,omitempty" name:"IsolateDataDisk"`
}

type IsolateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表。一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。每次请求退还实例和数据盘数量总计上限为20。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 是否退还挂载的数据盘。取值范围：
	// TRUE：表示退还实例同时退还其挂载的数据盘。
	// FALSE：表示退还实例同时不再退还其挂载的数据盘。
	// 默认取值：TRUE。
	IsolateDataDisk *bool `json:"IsolateDataDisk,omitnil,omitempty" name:"IsolateDataDisk"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 密钥对名称。
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// 密钥对的纯文本公钥。
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`

	// 密钥对关联的实例 ID 列表。
	AssociatedInstanceIds []*string `json:"AssociatedInstanceIds,omitnil,omitempty" name:"AssociatedInstanceIds"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 密钥对私钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// 密钥对绑定的标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type LoginConfiguration struct {
	// <li>"YES"代表选择自动生成密码，这时不指定Password字段。</li>
	// <li>"NO"代表选择自定义密码，这时要指定Password字段。</li>
	AutoGeneratePassword *string `json:"AutoGeneratePassword,omitnil,omitempty" name:"AutoGeneratePassword"`

	// 实例登录密码。具体按照操作系统的复杂度要求。 
	// 
	// `LINUX_UNIX` 实例密码必须 8-30 位，推荐使用 12 位以上密码，不能包含空格，不能以“/”开头，至少包含以下字符中的三种不同字符，字符种类：<br>
	// <li>小写字母：[a-z]</li>
	// <li>大写字母：[A-Z]</li>
	// <li>数字：0-9</li>
	// <li>特殊字符： ()\`\~!@#$%^&\*-+=\_|{}[]:;' <>,.?/</li> 
	// 
	// `WINDOWS` 实例密码必须 12-30 位，不能包含空格，不能以“/”开头且不包括用户名，至少包含以下字符中的三种不同字符：<br>
	// <li>小写字母：[a-z]</li>
	// <li>大写字母：[A-Z]</li>
	// <li>数字：0-9</li>
	// <li>特殊字符：()\`~!@#$%^&\*-+=\_|{}[]:;' <>,.?/</li> 
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 密钥ID列表，最多同时指定5个密钥。关联密钥后，就可以通过对应的私钥来访问实例。密钥与密码不能同时指定，同时WINDOWS操作系统不支持指定密钥。密钥ID列表可以通过[DescribeKeyPairs](https://cloud.tencent.com/document/product/1207/55540)接口获取。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

type LoginSettings struct {
	// 密钥 ID 列表。关联密钥后，就可以通过对应的私钥来访问实例。注意：此字段可能返回 []，表示取不到有效值。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

type McpServer struct {
	// MCP Server ID。
	McpServerId *string `json:"McpServerId,omitnil,omitempty" name:"McpServerId"`

	// MCP Server名称。最大长度：64
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// MCP Server类型。枚举值：PUBLIC_PACKAGE，公共包安装；AGENT_GENERATED，AI生成。
	McpServerType *string `json:"McpServerType,omitnil,omitempty" name:"McpServerType"`

	// MCP Server图标地址
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// Base64编码后的MCP Server启动命令。最大长度：2048
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// MCP Server状态。枚举值如下：
	// 
	// PENDING：表示创建中
	// LAUNCH_FAILED：表示创建失败
	// RUNNING：表示运行中
	// STOPPED：表示关闭
	// STARTING：表示开启中
	// STOPPING：表示关闭中
	// RESTARTING：表示重启中
	// REMOVING：表示删除中
	// UNKNOWN：表示未知
	// ENV_ERROR：表示环境错误
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// MCP Server访问地址。
	ServerUrl *string `json:"ServerUrl,omitnil,omitempty" name:"ServerUrl"`

	// MCP Server配置
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// MCP Server备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// MCP Server创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// MCP Server修改时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// MCP Server环境变量
	EnvSet []*McpServerEnv `json:"EnvSet,omitnil,omitempty" name:"EnvSet"`
}

type McpServerEnv struct {
	// MCP Server的环境变量键。最大长度：128
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// MCP Server的环境变量值。最大长度：1024。该字段可能存储密钥，出参时将固定返回“**********”，避免明文泄露。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type McpServerTemplate struct {
	// MCP Server名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Base64编码之后的MCP Server启动命令。
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// MCP Server图标地址
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// MCP Server社区地址
	CommunityUrl *string `json:"CommunityUrl,omitnil,omitempty" name:"CommunityUrl"`

	// MCP Server关联的开发平台地址或开放平台地址
	PlatformUrl *string `json:"PlatformUrl,omitnil,omitempty" name:"PlatformUrl"`

	// MCP Server环境变量
	EnvSet []*McpServerTemplateEnv `json:"EnvSet,omitnil,omitempty" name:"EnvSet"`
}

type McpServerTemplateEnv struct {
	// MCP Server模板的环境变量键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// MCP Server模板的环境变量值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type ModifyBlueprintAttributeRequestParams struct {
	// 镜像 ID。可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值中的BlueprintId获取。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`

	// 设置新的镜像名称。最大长度60。
	BlueprintName *string `json:"BlueprintName,omitnil,omitempty" name:"BlueprintName"`

	// 设置新的镜像描述。最大长度60。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyBlueprintAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 镜像 ID。可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值中的BlueprintId获取。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`

	// 设置新的镜像名称。最大长度60。
	BlueprintName *string `json:"BlueprintName,omitnil,omitempty" name:"BlueprintName"`

	// 设置新的镜像描述。最大长度60。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ModifyPrice *Price `json:"ModifyPrice,omitnil,omitempty" name:"ModifyPrice"`

	// 变更套餐状态。取值：
	// <li>SOLD_OUT：套餐售罄</li>
	// <li>AVAILABLE：支持套餐变更</li>
	// <li>UNAVAILABLE：暂不支持套餐变更</li>
	ModifyBundleState *string `json:"ModifyBundleState,omitnil,omitempty" name:"ModifyBundleState"`

	// 套餐信息。
	Bundle *Bundle `json:"Bundle,omitnil,omitempty" name:"Bundle"`

	// 不支持套餐变更原因信息。变更套餐状态为"AVAILABLE"时, 该信息为空
	NotSupportModifyMessage *string `json:"NotSupportModifyMessage,omitnil,omitempty" name:"NotSupportModifyMessage"`
}

// Predefined struct for user
type ModifyDiskBackupsAttributeRequestParams struct {
	// 云硬盘备份点ID，可通过 [DescribeDiskBackups](https://cloud.tencent.com/document/api/1207/84379) 接口返回值中的 DiskBackupId 获取。列表长度最大值为100。
	DiskBackupIds []*string `json:"DiskBackupIds,omitnil,omitempty" name:"DiskBackupIds"`

	// 云硬盘备份点名称，最大长度 90 。
	DiskBackupName *string `json:"DiskBackupName,omitnil,omitempty" name:"DiskBackupName"`
}

type ModifyDiskBackupsAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘备份点ID，可通过 [DescribeDiskBackups](https://cloud.tencent.com/document/api/1207/84379) 接口返回值中的 DiskBackupId 获取。列表长度最大值为100。
	DiskBackupIds []*string `json:"DiskBackupIds,omitnil,omitempty" name:"DiskBackupIds"`

	// 云硬盘备份点名称，最大长度 90 。
	DiskBackupName *string `json:"DiskBackupName,omitnil,omitempty" name:"DiskBackupName"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 云硬盘ID列表。每次批量请求云硬盘的上限为 100。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 云硬盘名称。
	DiskName *string `json:"DiskName,omitnil,omitempty" name:"DiskName"`
}

type ModifyDisksAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。每次批量请求云硬盘的上限为 100。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 云硬盘名称。
	DiskName *string `json:"DiskName,omitnil,omitempty" name:"DiskName"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyDisksBackupQuotaRequestParams struct {
	// 云硬盘ID列表，可通过[DescribeDisks](https://cloud.tencent.com/document/api/1207/66093)接口查询。列表最大长度为15。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 云硬盘备份点配额。取值范围: [0, 500]。调整后的配额必须大于等于已存在的备份点数量。
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitnil,omitempty" name:"DiskBackupQuota"`
}

type ModifyDisksBackupQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表，可通过[DescribeDisks](https://cloud.tencent.com/document/api/1207/66093)接口查询。列表最大长度为15。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 云硬盘备份点配额。取值范围: [0, 500]。调整后的配额必须大于等于已存在的备份点数量。
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitnil,omitempty" name:"DiskBackupQuota"`
}

func (r *ModifyDisksBackupQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDisksBackupQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "DiskBackupQuota")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDisksBackupQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDisksBackupQuotaResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDisksBackupQuotaResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDisksBackupQuotaResponseParams `json:"Response"`
}

func (r *ModifyDisksBackupQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDisksBackupQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDisksRenewFlagRequestParams struct {
	// 云硬盘ID列表。每次批量请求云硬盘的上限为 100。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 自动续费标识。取值范围：
	// 
	// - NOTIFY_AND_AUTO_RENEW：通知过期且自动续费
	// - NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费
	// - DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费
	// 
	// 若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type ModifyDisksRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。每次批量请求云硬盘的上限为 100。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 自动续费标识。取值范围：
	// 
	// - NOTIFY_AND_AUTO_RENEW：通知过期且自动续费
	// - NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费
	// - DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费
	// 
	// 若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyDockerContainerRequestParams struct {
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 容器ID。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。
	ContainerId *string `json:"ContainerId,omitnil,omitempty" name:"ContainerId"`

	// 环境变量列表
	Envs []*ContainerEnv `json:"Envs,omitnil,omitempty" name:"Envs"`

	// 容器端口主机端口映射列表
	PublishPorts []*DockerContainerPublishPort `json:"PublishPorts,omitnil,omitempty" name:"PublishPorts"`

	// 容器加载本地卷列表
	Volumes []*DockerContainerVolume `json:"Volumes,omitnil,omitempty" name:"Volumes"`

	// 运行的命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 容器重启策略，对应docker "--restart"参数。
	// 
	// 枚举值:
	// no: 不自动重启。默认策略。
	// on-failure[:max-retries]: 当容器退出码非0时重启容器。使用max-retries限制重启次数，比如on-failure:10，限制最多重启10次。
	// always: 只要容器退出就重启。
	// unless-stopped: 始终重新启动容器，包括在守护进程启动时，除非容器在 Docker 守护进程停止之前进入停止状态。
	RestartPolicy *string `json:"RestartPolicy,omitnil,omitempty" name:"RestartPolicy"`
}

type ModifyDockerContainerRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 容器ID。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。
	ContainerId *string `json:"ContainerId,omitnil,omitempty" name:"ContainerId"`

	// 环境变量列表
	Envs []*ContainerEnv `json:"Envs,omitnil,omitempty" name:"Envs"`

	// 容器端口主机端口映射列表
	PublishPorts []*DockerContainerPublishPort `json:"PublishPorts,omitnil,omitempty" name:"PublishPorts"`

	// 容器加载本地卷列表
	Volumes []*DockerContainerVolume `json:"Volumes,omitnil,omitempty" name:"Volumes"`

	// 运行的命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 容器重启策略，对应docker "--restart"参数。
	// 
	// 枚举值:
	// no: 不自动重启。默认策略。
	// on-failure[:max-retries]: 当容器退出码非0时重启容器。使用max-retries限制重启次数，比如on-failure:10，限制最多重启10次。
	// always: 只要容器退出就重启。
	// unless-stopped: 始终重新启动容器，包括在守护进程启动时，除非容器在 Docker 守护进程停止之前进入停止状态。
	RestartPolicy *string `json:"RestartPolicy,omitnil,omitempty" name:"RestartPolicy"`
}

func (r *ModifyDockerContainerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDockerContainerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ContainerId")
	delete(f, "Envs")
	delete(f, "PublishPorts")
	delete(f, "Volumes")
	delete(f, "Command")
	delete(f, "RestartPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDockerContainerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDockerContainerResponseParams struct {
	// Docker活动ID。
	DockerActivityId *string `json:"DockerActivityId,omitnil,omitempty" name:"DockerActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDockerContainerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDockerContainerResponseParams `json:"Response"`
}

func (r *ModifyDockerContainerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDockerContainerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFirewallRuleDescriptionRequestParams struct {
	// 实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/1207/47573) 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 防火墙规则。
	FirewallRule *FirewallRule `json:"FirewallRule,omitnil,omitempty" name:"FirewallRule"`

	// 防火墙当前版本。用户每次更新防火墙规则时版本会自动加1，防止规则已过期，不填不考虑冲突。
	FirewallVersion *uint64 `json:"FirewallVersion,omitnil,omitempty" name:"FirewallVersion"`
}

type ModifyFirewallRuleDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/1207/47573) 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 防火墙规则。
	FirewallRule *FirewallRule `json:"FirewallRule,omitnil,omitempty" name:"FirewallRule"`

	// 防火墙当前版本。用户每次更新防火墙规则时版本会自动加1，防止规则已过期，不填不考虑冲突。
	FirewallVersion *uint64 `json:"FirewallVersion,omitnil,omitempty" name:"FirewallVersion"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例 ID。实例的ID可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 防火墙规则列表。列表长度最大值是100。
	FirewallRules []*FirewallRule `json:"FirewallRules,omitnil,omitempty" name:"FirewallRules"`

	// 防火墙当前版本。用户每次更新防火墙规则时版本会自动加1，防止规则已过期，不填不考虑冲突。
	FirewallVersion *uint64 `json:"FirewallVersion,omitnil,omitempty" name:"FirewallVersion"`
}

type ModifyFirewallRulesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。实例的ID可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 防火墙规则列表。列表长度最大值是100。
	FirewallRules []*FirewallRule `json:"FirewallRules,omitnil,omitempty" name:"FirewallRules"`

	// 防火墙当前版本。用户每次更新防火墙规则时版本会自动加1，防止规则已过期，不填不考虑冲突。
	FirewallVersion *uint64 `json:"FirewallVersion,omitnil,omitempty" name:"FirewallVersion"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyFirewallTemplateRequestParams struct {
	// 防火墙模板ID。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 防火墙模板名称。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`
}

type ModifyFirewallTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙模板ID。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 防火墙模板名称。可通过[DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874)接口返回值字段TemplateSet获取。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`
}

func (r *ModifyFirewallTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFirewallTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TemplateName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFirewallTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFirewallTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyFirewallTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFirewallTemplateResponseParams `json:"Response"`
}

func (r *ModifyFirewallTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFirewallTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageSharePermissionRequestParams struct {
	// 镜像 ID。可通过[DescribeImages](https://cloud.tencent.com/document/api/213/15715)接口返回值中的ImageId获取。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 共享属性，包括 SHARE，CANCEL。其中SHARE代表共享，CANCEL代表取消共享。
	Permission *string `json:"Permission,omitnil,omitempty" name:"Permission"`
}

type ModifyImageSharePermissionRequest struct {
	*tchttp.BaseRequest
	
	// 镜像 ID。可通过[DescribeImages](https://cloud.tencent.com/document/api/213/15715)接口返回值中的ImageId获取。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 共享属性，包括 SHARE，CANCEL。其中SHARE代表共享，CANCEL代表取消共享。
	Permission *string `json:"Permission,omitnil,omitempty" name:"Permission"`
}

func (r *ModifyImageSharePermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageSharePermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageId")
	delete(f, "Permission")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyImageSharePermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageSharePermissionResponseParams struct {
	// CVM自定义镜像共享到轻量应用服务器后的镜像ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyImageSharePermissionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyImageSharePermissionResponseParams `json:"Response"`
}

func (r *ModifyImageSharePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageSharePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesAttributeRequestParams struct {
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例名称。可任意命名，但不得超过 60 个字符。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyInstancesAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例名称。可任意命名，但不得超过 60 个字符。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例ID列表。一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。每次请求批量实例的上限为15。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 待变更的套餐Id。注意不可和当前要升配的实例套餐ID相同。可通过[DescribeBundles](https://cloud.tencent.com/document/api/1207/47575)接口返回值中的BundleId获取。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 是否自动抵扣代金券。取值范围：
	// true：表示自动抵扣代金券
	// false：表示不自动抵扣代金券
	// 默认取值：false。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

type ModifyInstancesBundleRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表。一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。每次请求批量实例的上限为15。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 待变更的套餐Id。注意不可和当前要升配的实例套餐ID相同。可通过[DescribeBundles](https://cloud.tencent.com/document/api/1207/47575)接口返回值中的BundleId获取。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 是否自动抵扣代金券。取值范围：
	// true：表示自动抵扣代金券
	// false：表示不自动抵扣代金券
	// 默认取值：false。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyInstancesRenewFlagRequestParams struct {
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 自动续费标识。取值范围：
	// 
	// - NOTIFY_AND_AUTO_RENEW：通知过期且自动续费
	// - NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费
	// - DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费
	// 
	// 若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type ModifyInstancesRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 自动续费标识。取值范围：
	// 
	// - NOTIFY_AND_AUTO_RENEW：通知过期且自动续费
	// - NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费
	// - DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费
	// 
	// 若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyMcpServerRequestParams struct {
	// 实例ID。可以通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// MCP Server ID。可以通过DescribeMcpServers接口返回值中的McpServerId获取。
	McpServerId *string `json:"McpServerId,omitnil,omitempty" name:"McpServerId"`

	// MCP Server名称。最大长度：64
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Base64编码后的MCP Server启动命令。最大长度：2048
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// MCP Server备注。最大长度：2048
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// MCP Server环境变量。最大长度：10。用于完整替换MCP Server的环境变量。当该字段为空时，系统将清除当前所有环境变量。若无需修改环境变量，请勿传递该字段。
	Envs []*McpServerEnv `json:"Envs,omitnil,omitempty" name:"Envs"`
}

type ModifyMcpServerRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可以通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// MCP Server ID。可以通过DescribeMcpServers接口返回值中的McpServerId获取。
	McpServerId *string `json:"McpServerId,omitnil,omitempty" name:"McpServerId"`

	// MCP Server名称。最大长度：64
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Base64编码后的MCP Server启动命令。最大长度：2048
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// MCP Server备注。最大长度：2048
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// MCP Server环境变量。最大长度：10。用于完整替换MCP Server的环境变量。当该字段为空时，系统将清除当前所有环境变量。若无需修改环境变量，请勿传递该字段。
	Envs []*McpServerEnv `json:"Envs,omitnil,omitempty" name:"Envs"`
}

func (r *ModifyMcpServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMcpServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "McpServerId")
	delete(f, "Name")
	delete(f, "Command")
	delete(f, "Description")
	delete(f, "Envs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMcpServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMcpServerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMcpServerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMcpServerResponseParams `json:"Response"`
}

func (r *ModifyMcpServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMcpServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotAttributeRequestParams struct {
	// 快照 ID, 可通过 <a href="https://cloud.tencent.com/document/product/1207/54388">DescribeSnapshots</a> 查询。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 新的快照名称，最长为 60 个字符。
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`
}

type ModifySnapshotAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 快照 ID, 可通过 <a href="https://cloud.tencent.com/document/product/1207/54388">DescribeSnapshots</a> 查询。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 新的快照名称，最长为 60 个字符。
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	UserDiscount *float64 `json:"UserDiscount,omitnil,omitempty" name:"UserDiscount"`

	// 公共折扣。
	CommonDiscount *float64 `json:"CommonDiscount,omitnil,omitempty" name:"CommonDiscount"`

	// 最终折扣。
	FinalDiscount *float64 `json:"FinalDiscount,omitnil,omitempty" name:"FinalDiscount"`

	// 活动折扣。取值为null，表示无有效值，即没有折扣。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityDiscount *float64 `json:"ActivityDiscount,omitnil,omitempty" name:"ActivityDiscount"`

	// 折扣类型。
	// user：用户折扣; common：官网折扣; activity：活动折扣。 取值为null，表示无有效值，即没有折扣。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountType *string `json:"DiscountType,omitnil,omitempty" name:"DiscountType"`
}

type Price struct {
	// 实例价格。
	InstancePrice *InstancePrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`
}

// Predefined struct for user
type RebootInstancesRequestParams struct {
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 关机类型。
	// 取值范围：
	// - SOFT：表示软关机 
	// - HARD：表示硬关机 
	// - SOFT_FIRST：表示优先软关机，失败再执行硬关机  
	// 
	// 默认取值：SOFT_FIRST。
	StopType *string `json:"StopType,omitnil,omitempty" name:"StopType"`
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 关机类型。
	// 取值范围：
	// - SOFT：表示软关机 
	// - HARD：表示硬关机 
	// - SOFT_FIRST：表示优先软关机，失败再执行硬关机  
	// 
	// 默认取值：SOFT_FIRST。
	StopType *string `json:"StopType,omitnil,omitempty" name:"StopType"`
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
	delete(f, "StopType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RebootInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RebootInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 地域描述，例如，华南地区(广州)。
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 地域是否可用状态，取值仅为AVAILABLE（表示可用状态）。
	RegionState *string `json:"RegionState,omitnil,omitempty" name:"RegionState"`

	// 是否中国大陆地域
	IsChinaMainland *bool `json:"IsChinaMainland,omitnil,omitempty" name:"IsChinaMainland"`
}

// Predefined struct for user
type RemoveDockerContainersRequestParams struct {
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 容器ID列表。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。
	ContainerIds []*string `json:"ContainerIds,omitnil,omitempty" name:"ContainerIds"`
}

type RemoveDockerContainersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 容器ID列表。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。
	ContainerIds []*string `json:"ContainerIds,omitnil,omitempty" name:"ContainerIds"`
}

func (r *RemoveDockerContainersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveDockerContainersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ContainerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveDockerContainersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveDockerContainersResponseParams struct {
	// Docker活动ID。
	DockerActivityId *string `json:"DockerActivityId,omitnil,omitempty" name:"DockerActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveDockerContainersResponse struct {
	*tchttp.BaseResponse
	Response *RemoveDockerContainersResponseParams `json:"Response"`
}

func (r *RemoveDockerContainersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveDockerContainersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveMcpServersRequestParams struct {
	// 实例ID。可以通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// MCP Server ID列表。可通过DescribeMcpServers接口返回值中的McpServerId获取。最大长度：10
	McpServerIds []*string `json:"McpServerIds,omitnil,omitempty" name:"McpServerIds"`
}

type RemoveMcpServersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可以通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// MCP Server ID列表。可通过DescribeMcpServers接口返回值中的McpServerId获取。最大长度：10
	McpServerIds []*string `json:"McpServerIds,omitnil,omitempty" name:"McpServerIds"`
}

func (r *RemoveMcpServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveMcpServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "McpServerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveMcpServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveMcpServersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveMcpServersResponse struct {
	*tchttp.BaseResponse
	Response *RemoveMcpServersResponseParams `json:"Response"`
}

func (r *RemoveMcpServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveMcpServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenameDockerContainerRequestParams struct {
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 容器ID。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。
	ContainerId *string `json:"ContainerId,omitnil,omitempty" name:"ContainerId"`

	// 容器新的名称。
	ContainerName *string `json:"ContainerName,omitnil,omitempty" name:"ContainerName"`
}

type RenameDockerContainerRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 容器ID。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。
	ContainerId *string `json:"ContainerId,omitnil,omitempty" name:"ContainerId"`

	// 容器新的名称。
	ContainerName *string `json:"ContainerName,omitnil,omitempty" name:"ContainerName"`
}

func (r *RenameDockerContainerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenameDockerContainerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ContainerId")
	delete(f, "ContainerName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenameDockerContainerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenameDockerContainerResponseParams struct {
	// Docker活动ID。
	DockerActivityId *string `json:"DockerActivityId,omitnil,omitempty" name:"DockerActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenameDockerContainerResponse struct {
	*tchttp.BaseResponse
	Response *RenameDockerContainerResponseParams `json:"Response"`
}

func (r *RenameDockerContainerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenameDockerContainerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewDiskChargePrepaid struct {
	// 续费周期。
	// 单位：月。
	// 取值范围: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36]
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 自动续费标识。
	// 取值范围：
	// <li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li>
	// <li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费，用户需要手动续费</li><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不自动续费，且不通知</li>
	// 默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，云硬盘到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 周期单位。取值范围：“m”(月)。默认值: "m"。
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 当前实例到期时间。如“2018-01-01 00:00:00”。
	// 指定该参数即可对齐云硬盘所挂载的实例到期时间。
	// 该参数与Period必须指定其一，且不支持同时指定。
	// 该参数值必须大于入参中云硬盘的过期时间。
	CurInstanceDeadline *string `json:"CurInstanceDeadline,omitnil,omitempty" name:"CurInstanceDeadline"`
}

// Predefined struct for user
type RenewDisksRequestParams struct {
	// 云硬盘ID列表。一个或多个待操作的云硬盘ID。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。每次请求续费数据盘数量总计上限为50。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 续费云硬盘包年包月相关参数设置。
	RenewDiskChargePrepaid *RenewDiskChargePrepaid `json:"RenewDiskChargePrepaid,omitnil,omitempty" name:"RenewDiskChargePrepaid"`

	// 是否自动使用代金券。默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

type RenewDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。一个或多个待操作的云硬盘ID。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。每次请求续费数据盘数量总计上限为50。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 续费云硬盘包年包月相关参数设置。
	RenewDiskChargePrepaid *RenewDiskChargePrepaid `json:"RenewDiskChargePrepaid,omitnil,omitempty" name:"RenewDiskChargePrepaid"`

	// 是否自动使用代金券。默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

func (r *RenewDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "RenewDiskChargePrepaid")
	delete(f, "AutoVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDisksResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewDisksResponse struct {
	*tchttp.BaseResponse
	Response *RenewDisksResponseParams `json:"Response"`
}

func (r *RenewDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstancesRequestParams struct {
	// 实例ID列表。一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 是否续费弹性数据盘。取值范围：
	// TRUE：表示续费实例同时续费其挂载的数据盘
	// FALSE：表示续费实例同时不再续费其挂载的数据盘
	// 默认取值：TRUE。
	RenewDataDisk *bool `json:"RenewDataDisk,omitnil,omitempty" name:"RenewDataDisk"`

	// 是否自动抵扣代金券。取值范围：
	// TRUE：表示自动抵扣代金券
	// FALSE：表示不自动抵扣代金券
	// 默认取值：FALSE。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

type RenewInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表。一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 是否续费弹性数据盘。取值范围：
	// TRUE：表示续费实例同时续费其挂载的数据盘
	// FALSE：表示续费实例同时不再续费其挂载的数据盘
	// 默认取值：TRUE。
	RenewDataDisk *bool `json:"RenewDataDisk,omitnil,omitempty" name:"RenewDataDisk"`

	// 是否自动抵扣代金券。取值范围：
	// TRUE：表示自动抵扣代金券
	// FALSE：表示不自动抵扣代金券
	// 默认取值：FALSE。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ReplaceFirewallTemplateRuleRequestParams struct {
	// 防火墙模板ID。可通过 [DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874) 接口返回值中的 TemplateId 获取。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 防火墙模板规则ID。可通过 [DescribeFirewallTemplateRules](https://cloud.tencent.com/document/product/1207/96875) 接口返回值中的 TemplateRuleId 获取。
	TemplateRuleId *string `json:"TemplateRuleId,omitnil,omitempty" name:"TemplateRuleId"`

	// 替换后的防火墙模板规则。
	TemplateRule *FirewallRule `json:"TemplateRule,omitnil,omitempty" name:"TemplateRule"`
}

type ReplaceFirewallTemplateRuleRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙模板ID。可通过 [DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874) 接口返回值中的 TemplateId 获取。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 防火墙模板规则ID。可通过 [DescribeFirewallTemplateRules](https://cloud.tencent.com/document/product/1207/96875) 接口返回值中的 TemplateRuleId 获取。
	TemplateRuleId *string `json:"TemplateRuleId,omitnil,omitempty" name:"TemplateRuleId"`

	// 替换后的防火墙模板规则。
	TemplateRule *FirewallRule `json:"TemplateRule,omitnil,omitempty" name:"TemplateRule"`
}

func (r *ReplaceFirewallTemplateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceFirewallTemplateRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TemplateRuleId")
	delete(f, "TemplateRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceFirewallTemplateRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceFirewallTemplateRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReplaceFirewallTemplateRuleResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceFirewallTemplateRuleResponseParams `json:"Response"`
}

func (r *ReplaceFirewallTemplateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceFirewallTemplateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RerunDockerContainerRequestParams struct {
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 重新创建的容器配置。
	ContainerConfiguration *DockerContainerConfiguration `json:"ContainerConfiguration,omitnil,omitempty" name:"ContainerConfiguration"`

	// 容器ID。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。
	ContainerId *string `json:"ContainerId,omitnil,omitempty" name:"ContainerId"`
}

type RerunDockerContainerRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 重新创建的容器配置。
	ContainerConfiguration *DockerContainerConfiguration `json:"ContainerConfiguration,omitnil,omitempty" name:"ContainerConfiguration"`

	// 容器ID。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。
	ContainerId *string `json:"ContainerId,omitnil,omitempty" name:"ContainerId"`
}

func (r *RerunDockerContainerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RerunDockerContainerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ContainerConfiguration")
	delete(f, "ContainerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RerunDockerContainerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RerunDockerContainerResponseParams struct {
	// Docker活动ID。
	DockerActivityId *string `json:"DockerActivityId,omitnil,omitempty" name:"DockerActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RerunDockerContainerResponse struct {
	*tchttp.BaseResponse
	Response *RerunDockerContainerResponseParams `json:"Response"`
}

func (r *RerunDockerContainerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RerunDockerContainerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAttachCcnRequestParams struct {
	// 云联网实例ID。可通过[DescribeCcns](https://cloud.tencent.com/document/product/215/19199)接口返回值中的CcnId获取。
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`
}

type ResetAttachCcnRequest struct {
	*tchttp.BaseRequest
	
	// 云联网实例ID。可通过[DescribeCcns](https://cloud.tencent.com/document/product/215/19199)接口返回值中的CcnId获取。
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ResetFirewallTemplateRulesRequestParams struct {
	// 防火墙模板ID。可通过 [DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874) 接口返回值中的 TemplateId	获取。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 重置后的防火墙模板规则列表。每次请求批量防火墙规则的上限为 100。
	TemplateRules []*FirewallRule `json:"TemplateRules,omitnil,omitempty" name:"TemplateRules"`
}

type ResetFirewallTemplateRulesRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙模板ID。可通过 [DescribeFirewallTemplates](https://cloud.tencent.com/document/product/1207/96874) 接口返回值中的 TemplateId	获取。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 重置后的防火墙模板规则列表。每次请求批量防火墙规则的上限为 100。
	TemplateRules []*FirewallRule `json:"TemplateRules,omitnil,omitempty" name:"TemplateRules"`
}

func (r *ResetFirewallTemplateRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetFirewallTemplateRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TemplateRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetFirewallTemplateRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetFirewallTemplateRulesResponseParams struct {
	// 重置后的规则ID列表。
	TemplateRuleIdSet []*string `json:"TemplateRuleIdSet,omitnil,omitempty" name:"TemplateRuleIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetFirewallTemplateRulesResponse struct {
	*tchttp.BaseResponse
	Response *ResetFirewallTemplateRulesResponseParams `json:"Response"`
}

func (r *ResetFirewallTemplateRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetFirewallTemplateRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetInstanceBlueprint struct {
	// 镜像详细信息
	BlueprintInfo *Blueprint `json:"BlueprintInfo,omitnil,omitempty" name:"BlueprintInfo"`

	// 实例镜像是否可重置为目标镜像。
	// 取值：
	// true（允许）
	// false（不允许）
	IsResettable *bool `json:"IsResettable,omitnil,omitempty" name:"IsResettable"`

	// 不可重置信息.当镜像可重置时为""
	NonResettableMessage *string `json:"NonResettableMessage,omitnil,omitempty" name:"NonResettableMessage"`
}

// Predefined struct for user
type ResetInstanceRequestParams struct {
	// 实例 ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 镜像 ID。可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值中的BlueprintId获取。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`

	// 要创建的容器配置列表。
	// 注意：仅重装的镜像类型为Docker时支持传入该参数。
	Containers []*DockerContainerConfiguration `json:"Containers,omitnil,omitempty" name:"Containers"`

	// 实例登录信息配置。默认缺失情况下代表用户选择实例创建后设置登录密码或绑定密钥。
	LoginConfiguration *LoginConfiguration `json:"LoginConfiguration,omitnil,omitempty" name:"LoginConfiguration"`
}

type ResetInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 镜像 ID。可通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回值中的BlueprintId获取。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`

	// 要创建的容器配置列表。
	// 注意：仅重装的镜像类型为Docker时支持传入该参数。
	Containers []*DockerContainerConfiguration `json:"Containers,omitnil,omitempty" name:"Containers"`

	// 实例登录信息配置。默认缺失情况下代表用户选择实例创建后设置登录密码或绑定密钥。
	LoginConfiguration *LoginConfiguration `json:"LoginConfiguration,omitnil,omitempty" name:"LoginConfiguration"`
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
	delete(f, "Containers")
	delete(f, "LoginConfiguration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过 <a href="https://cloud.tencent.com/document/product/1207/47573">DescribeInstances</a> 接口返回值中的 InstanceId 获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：</br> `LINUX_UNIX` 实例密码必须 8-30 位，推荐使用 12 位以上密码，不能以“/”开头，至少包含以下字符中的三种不同字符，字符种类：</br> <li>小写字母：[a-z]</br></li> <li>大写字母：[A-Z]</br></li> <li>数字：0-9</br></li> <li>特殊字符： ()\`\~!@#$%^&\*-+=\_|{}[]:;' <>,.?/</li></br> `WINDOWS` 实例密码必须 12-30 位，不能以“/”开头且不包括用户名，至少包含以下字符中的三种不同字符</br> <li>小写字母：[a-z]</br></li> <li>大写字母：[A-Z]</br></li> <li>数字： 0-9</br></li> <li>特殊字符：()\`~!@#$%^&\*-+=\_|{}[]:;' <>,.?/</br></li> <li>如果实例即包含 `LINUX_UNIX` 实例又包含 `WINDOWS` 实例，则密码复杂度限制按照 `WINDOWS` 实例的限制。</li>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 待重置密码的实例操作系统用户名。不得超过 64 个字符。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

type ResetInstancesPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过 <a href="https://cloud.tencent.com/document/product/1207/47573">DescribeInstances</a> 接口返回值中的 InstanceId 获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：</br> `LINUX_UNIX` 实例密码必须 8-30 位，推荐使用 12 位以上密码，不能以“/”开头，至少包含以下字符中的三种不同字符，字符种类：</br> <li>小写字母：[a-z]</br></li> <li>大写字母：[A-Z]</br></li> <li>数字：0-9</br></li> <li>特殊字符： ()\`\~!@#$%^&\*-+=\_|{}[]:;' <>,.?/</li></br> `WINDOWS` 实例密码必须 12-30 位，不能以“/”开头且不包括用户名，至少包含以下字符中的三种不同字符</br> <li>小写字母：[a-z]</br></li> <li>大写字母：[A-Z]</br></li> <li>数字： 0-9</br></li> <li>特殊字符：()\`~!@#$%^&\*-+=\_|{}[]:;' <>,.?/</br></li> <li>如果实例即包含 `LINUX_UNIX` 实例又包含 `WINDOWS` 实例，则密码复杂度限制按照 `WINDOWS` 实例的限制。</li>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 待重置密码的实例操作系统用户名。不得超过 64 个字符。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
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
type ResizeDisksRequestParams struct {
	// 云硬盘ID列表，可通过[DescribeDisks](https://cloud.tencent.com/document/api/1207/66093)接口查询。列表最大长度为15。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 扩容后的云硬盘大小。单位: GB。高性能云硬盘大小取值范围：[10, 4000] ,SSD云硬盘大小取值范围：[20, 4000]。扩容后的云硬盘大小必须大于当前云硬盘大小。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

type ResizeDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表，可通过[DescribeDisks](https://cloud.tencent.com/document/api/1207/66093)接口查询。列表最大长度为15。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 扩容后的云硬盘大小。单位: GB。高性能云硬盘大小取值范围：[10, 4000] ,SSD云硬盘大小取值范围：[20, 4000]。扩容后的云硬盘大小必须大于当前云硬盘大小。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

func (r *ResizeDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "DiskSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResizeDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeDisksResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResizeDisksResponse struct {
	*tchttp.BaseResponse
	Response *ResizeDisksResponseParams `json:"Response"`
}

func (r *ResizeDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartDockerContainersRequestParams struct {
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 容器ID列表。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。
	ContainerIds []*string `json:"ContainerIds,omitnil,omitempty" name:"ContainerIds"`
}

type RestartDockerContainersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 容器ID列表。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。
	ContainerIds []*string `json:"ContainerIds,omitnil,omitempty" name:"ContainerIds"`
}

func (r *RestartDockerContainersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDockerContainersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ContainerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartDockerContainersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartDockerContainersResponseParams struct {
	// Docker活动ID。
	DockerActivityId *string `json:"DockerActivityId,omitnil,omitempty" name:"DockerActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartDockerContainersResponse struct {
	*tchttp.BaseResponse
	Response *RestartDockerContainersResponseParams `json:"Response"`
}

func (r *RestartDockerContainersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDockerContainersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartMcpServersRequestParams struct {
	// 实例 ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// MCP Server ID列表。可通过DescribeMcpServers接口返回值中的McpServerId获取。最大长度：10
	McpServerIds []*string `json:"McpServerIds,omitnil,omitempty" name:"McpServerIds"`
}

type RestartMcpServersRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// MCP Server ID列表。可通过DescribeMcpServers接口返回值中的McpServerId获取。最大长度：10
	McpServerIds []*string `json:"McpServerIds,omitnil,omitempty" name:"McpServerIds"`
}

func (r *RestartMcpServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartMcpServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "McpServerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartMcpServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartMcpServersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartMcpServersResponse struct {
	*tchttp.BaseResponse
	Response *RestartMcpServersResponseParams `json:"Response"`
}

func (r *RestartMcpServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartMcpServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunDockerContainersRequestParams struct {
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要创建的容器列表。
	Containers []*DockerContainerConfiguration `json:"Containers,omitnil,omitempty" name:"Containers"`
}

type RunDockerContainersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要创建的容器列表。
	Containers []*DockerContainerConfiguration `json:"Containers,omitnil,omitempty" name:"Containers"`
}

func (r *RunDockerContainersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunDockerContainersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Containers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunDockerContainersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunDockerContainersResponseParams struct {
	// Docker活动ID列表。
	DockerActivitySet []*string `json:"DockerActivitySet,omitnil,omitempty" name:"DockerActivitySet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunDockerContainersResponse struct {
	*tchttp.BaseResponse
	Response *RunDockerContainersResponseParams `json:"Response"`
}

func (r *RunDockerContainersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunDockerContainersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Scene struct {
	// 使用场景Id
	SceneId *string `json:"SceneId,omitnil,omitempty" name:"SceneId"`

	// 使用场景展示名称
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 使用场景描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type SceneInfo struct {
	// 使用场景Id。
	SceneId *string `json:"SceneId,omitnil,omitempty" name:"SceneId"`

	// 使用场景展示名称。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 使用场景描述信息。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

// Predefined struct for user
type ShareBlueprintAcrossAccountsRequestParams struct {
	// 镜像ID, 可以通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回的BlueprintId获取。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`

	// 接收共享镜像的[账号ID](https://cloud.tencent.com/document/product/213/4944#.E8.8E.B7.E5.8F.96.E4.B8.BB.E8.B4.A6.E5.8F.B7.E7.9A.84.E8.B4.A6.E5.8F.B7-id)列表。账号ID不同于QQ号，查询用户账号ID请查看账号信息中的账号ID栏。账号个数取值最大为10。
	AccountIds []*string `json:"AccountIds,omitnil,omitempty" name:"AccountIds"`
}

type ShareBlueprintAcrossAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 镜像ID, 可以通过[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)接口返回的BlueprintId获取。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`

	// 接收共享镜像的[账号ID](https://cloud.tencent.com/document/product/213/4944#.E8.8E.B7.E5.8F.96.E4.B8.BB.E8.B4.A6.E5.8F.B7.E7.9A.84.E8.B4.A6.E5.8F.B7-id)列表。账号ID不同于QQ号，查询用户账号ID请查看账号信息中的账号ID栏。账号个数取值最大为10。
	AccountIds []*string `json:"AccountIds,omitnil,omitempty" name:"AccountIds"`
}

func (r *ShareBlueprintAcrossAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ShareBlueprintAcrossAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BlueprintId")
	delete(f, "AccountIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ShareBlueprintAcrossAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ShareBlueprintAcrossAccountsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ShareBlueprintAcrossAccountsResponse struct {
	*tchttp.BaseResponse
	Response *ShareBlueprintAcrossAccountsResponseParams `json:"Response"`
}

func (r *ShareBlueprintAcrossAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ShareBlueprintAcrossAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Snapshot struct {
	// 快照 ID。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 创建此快照的磁盘类型。取值：<li>SYSTEM_DISK：系统盘</li>
	DiskUsage *string `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// 创建此快照的磁盘 ID。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 创建此快照的磁盘大小，单位 GB。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 快照名称，用户自定义的快照别名。
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 快照的状态。取值范围：
	// <li>NORMAL：正常 </li>
	// <li>CREATING：创建中</li>
	// <li>ROLLBACKING：回滚中。</li>
	SnapshotState *string `json:"SnapshotState,omitnil,omitempty" name:"SnapshotState"`

	// 创建或回滚快照进度百分比，成功后此字段取值为 100。
	Percent *int64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 快照的最新操作，只有创建、回滚快照时记录。
	// 取值如 CreateInstanceSnapshot，RollbackInstanceSnapshot。
	LatestOperation *string `json:"LatestOperation,omitnil,omitempty" name:"LatestOperation"`

	// 快照的最新操作状态，只有创建、回滚快照时记录。
	// 取值范围：
	// <li>SUCCESS：表示操作成功</li>
	// <li>OPERATING：表示操作执行中</li>
	// <li>FAILED：表示操作失败</li>
	LatestOperationState *string `json:"LatestOperationState,omitnil,omitempty" name:"LatestOperationState"`

	// 快照最新操作的唯一请求 ID，只有创建、回滚快照时记录。
	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitnil,omitempty" name:"LatestOperationRequestId"`

	// 快照的创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 快照绑定的标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type SnapshotDeniedActions struct {
	// 快照 ID。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 操作限制列表。
	DeniedActions []*DeniedAction `json:"DeniedActions,omitnil,omitempty" name:"DeniedActions"`
}

type Software struct {
	// 软件名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 软件版本。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 软件图片 URL。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 软件安装目录。
	InstallDir *string `json:"InstallDir,omitnil,omitempty" name:"InstallDir"`

	// 软件详情列表。
	DetailSet []*SoftwareDetail `json:"DetailSet,omitnil,omitempty" name:"DetailSet"`
}

type SoftwareDetail struct {
	// 软件的属性标识
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 软件的属性标识描述
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 软件的属性值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type StartDockerContainersRequestParams struct {
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 容器ID列表。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。
	ContainerIds []*string `json:"ContainerIds,omitnil,omitempty" name:"ContainerIds"`
}

type StartDockerContainersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 容器ID列表。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。
	ContainerIds []*string `json:"ContainerIds,omitnil,omitempty" name:"ContainerIds"`
}

func (r *StartDockerContainersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartDockerContainersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ContainerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartDockerContainersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartDockerContainersResponseParams struct {
	// Docker活动ID。
	DockerActivityId *string `json:"DockerActivityId,omitnil,omitempty" name:"DockerActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartDockerContainersResponse struct {
	*tchttp.BaseResponse
	Response *StartDockerContainersResponseParams `json:"Response"`
}

func (r *StartDockerContainersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartDockerContainersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartInstancesRequestParams struct {
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type StartInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type StartMcpServersRequestParams struct {
	// 实例 ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// MCP Server ID列表。可通过DescribeMcpServers接口返回值中的McpServerId获取。最大长度：10
	McpServerIds []*string `json:"McpServerIds,omitnil,omitempty" name:"McpServerIds"`
}

type StartMcpServersRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// MCP Server ID列表。可通过DescribeMcpServers接口返回值中的McpServerId获取。最大长度：10
	McpServerIds []*string `json:"McpServerIds,omitnil,omitempty" name:"McpServerIds"`
}

func (r *StartMcpServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMcpServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "McpServerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartMcpServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartMcpServersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartMcpServersResponse struct {
	*tchttp.BaseResponse
	Response *StartMcpServersResponseParams `json:"Response"`
}

func (r *StartMcpServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMcpServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopDockerContainersRequestParams struct {
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 容器ID列表。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。
	ContainerIds []*string `json:"ContainerIds,omitnil,omitempty" name:"ContainerIds"`
}

type StopDockerContainersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/product/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 容器ID列表。可通过[DescribeDockerContainers](https://cloud.tencent.com/document/product/1207/95473)接口返回值中的ContainerId获取。
	ContainerIds []*string `json:"ContainerIds,omitnil,omitempty" name:"ContainerIds"`
}

func (r *StopDockerContainersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopDockerContainersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ContainerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopDockerContainersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopDockerContainersResponseParams struct {
	// Docker活动ID。
	DockerActivityId *string `json:"DockerActivityId,omitnil,omitempty" name:"DockerActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopDockerContainersResponse struct {
	*tchttp.BaseResponse
	Response *StopDockerContainersResponseParams `json:"Response"`
}

func (r *StopDockerContainersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopDockerContainersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstancesRequestParams struct {
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 关机类型。
	// 取值范围： 
	// 
	// - SOFT：表示软关机
	// - HARD：表示硬关机 
	// - SOFT_FIRST：表示优先软关机，失败再执行硬关机  
	// 
	// 默认取值：SOFT_FIRST
	StopType *string `json:"StopType,omitnil,omitempty" name:"StopType"`
}

type StopInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求批量实例的上限为 100。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 关机类型。
	// 取值范围： 
	// 
	// - SOFT：表示软关机
	// - HARD：表示硬关机 
	// - SOFT_FIRST：表示优先软关机，失败再执行硬关机  
	// 
	// 默认取值：SOFT_FIRST
	StopType *string `json:"StopType,omitnil,omitempty" name:"StopType"`
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
	delete(f, "StopType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type StopMcpServersRequestParams struct {
	// 实例 ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// MCP Server ID列表。可通过DescribeMcpServers接口返回值中的McpServerId获取。最大长度：10
	McpServerIds []*string `json:"McpServerIds,omitnil,omitempty" name:"McpServerIds"`
}

type StopMcpServersRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// MCP Server ID列表。可通过DescribeMcpServers接口返回值中的McpServerId获取。最大长度：10
	McpServerIds []*string `json:"McpServerIds,omitnil,omitempty" name:"McpServerIds"`
}

func (r *StopMcpServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMcpServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "McpServerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopMcpServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMcpServersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopMcpServersResponse struct {
	*tchttp.BaseResponse
	Response *StopMcpServersResponseParams `json:"Response"`
}

func (r *StopMcpServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMcpServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SupportIpv6Detail struct {
	// 是否支持开启IPv6。
	IsSupport *bool `json:"IsSupport,omitnil,omitempty" name:"IsSupport"`

	// 详情。
	// 
	// 当IsSupport为True，Detail枚举值为:
	// 
	// EFFECTIVE_IMMEDIATELY: 立即生效
	// 
	// EFFECTIVE_AFTER_REBOOT: 分配过程需要开关机，用户需备份数据
	// 
	//  
	// 
	// 当IsSupport为False，Detail枚举值为:
	// 
	// HAD_BEEN_ASSIGNED: 已分配IPv6地址
	// 
	// REGION_NOT_SUPPORT: 地域不支持
	// 
	// BLUEPRINT_NOT_SUPPORT: 镜像不支持
	// 
	// BUNDLE_INSTANCE_NOT_SUPPORT: 套餐实例不支持
	// 
	// BUNDLE_BANDWIDTH_NOT_SUPPORT: 套餐带宽不支持
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 提示信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type SyncBlueprintRequestParams struct {
	// 镜像ID。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`

	// 同步镜像的目的地域列表。
	DestinationRegions []*string `json:"DestinationRegions,omitnil,omitempty" name:"DestinationRegions"`
}

type SyncBlueprintRequest struct {
	*tchttp.BaseRequest
	
	// 镜像ID。
	BlueprintId *string `json:"BlueprintId,omitnil,omitempty" name:"BlueprintId"`

	// 同步镜像的目的地域列表。
	DestinationRegions []*string `json:"DestinationRegions,omitnil,omitempty" name:"DestinationRegions"`
}

func (r *SyncBlueprintRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncBlueprintRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BlueprintId")
	delete(f, "DestinationRegions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncBlueprintRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncBlueprintResponseParams struct {
	// 目标地域镜像信息。
	DestinationRegionBlueprintSet []*DestinationRegionBlueprint `json:"DestinationRegionBlueprintSet,omitnil,omitempty" name:"DestinationRegionBlueprintSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SyncBlueprintResponse struct {
	*tchttp.BaseResponse
	Response *SyncBlueprintResponseParams `json:"Response"`
}

func (r *SyncBlueprintResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncBlueprintResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemDisk struct {
	// 系统盘类型。
	// 取值范围： 
	// <li> LOCAL_BASIC：本地硬盘</li><li> LOCAL_SSD：本地 SSD 硬盘</li><li> CLOUD_BASIC：普通云硬盘</li><li> CLOUD_SSD：SSD 云硬盘</li><li> CLOUD_PREMIUM：高性能云硬盘</li>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 系统盘大小，单位：GB。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 系统盘ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type TerminateDisksRequestParams struct {
	// 云硬盘ID列表。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。
	// 每次批量请求云硬盘的上限数量为100。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`
}

type TerminateDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID列表。可通过[DescribeDisks](https://cloud.tencent.com/document/product/1207/66093)接口返回值中的DiskId获取。
	// 每次批量请求云硬盘的上限数量为100。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表。可通过[DescribeInstances](https://cloud.tencent.com/document/api/1207/47573)接口返回值中的InstanceId获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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

type TotalPrice struct {
	// 原始总计价格。
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 折扣总计价格。
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`
}

type TrafficPackage struct {
	// 流量包ID。
	TrafficPackageId *string `json:"TrafficPackageId,omitnil,omitempty" name:"TrafficPackageId"`

	// 流量包生效周期内已使用流量，单位字节。
	TrafficUsed *int64 `json:"TrafficUsed,omitnil,omitempty" name:"TrafficUsed"`

	// 流量包生效周期内的总流量，单位字节。
	TrafficPackageTotal *int64 `json:"TrafficPackageTotal,omitnil,omitempty" name:"TrafficPackageTotal"`

	// 流量包生效周期内的剩余流量，单位字节。
	TrafficPackageRemaining *int64 `json:"TrafficPackageRemaining,omitnil,omitempty" name:"TrafficPackageRemaining"`

	// 流量包生效周期内超出流量包额度的流量，单位字节。
	TrafficOverflow *int64 `json:"TrafficOverflow,omitnil,omitempty" name:"TrafficOverflow"`

	// 流量包生效周期开始时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 流量包生效周期结束时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 流量包到期时间。按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 流量包状态：
	// <li>NETWORK_NORMAL：正常</li>
	// <li>OVERDUE_NETWORK_DISABLED：欠费断网</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ZoneInfo struct {
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 可用区中文名称
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 实例购买页可用区展示标签
	InstanceDisplayLabel *string `json:"InstanceDisplayLabel,omitnil,omitempty" name:"InstanceDisplayLabel"`
}