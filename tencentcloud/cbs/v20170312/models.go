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

package v20170312

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AdvancedRetentionPolicy struct {
	// 保留最新快照Days天内的每天最新的一个快照，取值范围：[0, 100]
	Days *uint64 `json:"Days,omitnil,omitempty" name:"Days"`

	// 保留最新快照Weeks周内的每周最新的一个快照，取值范围：[0, 100]
	Weeks *uint64 `json:"Weeks,omitnil,omitempty" name:"Weeks"`

	// 保留最新快照Months月内的每月最新的一个快照， 取值范围：[0, 100]
	Months *uint64 `json:"Months,omitnil,omitempty" name:"Months"`

	// 保留最新快照Years年内的每年最新的一个快照，取值范围：[0, 100]
	Years *uint64 `json:"Years,omitnil,omitempty" name:"Years"`
}

type ApplyDisk struct {
	// 快照组关联的快照ID。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 快照组关联快照对应的原云硬盘ID。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`
}

// Predefined struct for user
type ApplyDiskBackupRequestParams struct {
	// 云硬盘备份点ID，可以通过[DescribeDiskBackups](/document/product/362/80278)接口查询。
	DiskBackupId *string `json:"DiskBackupId,omitnil,omitempty" name:"DiskBackupId"`

	// 云硬盘备份点原云硬盘ID，可以通过[DescribeDisks](/document/product/362/16315)接口查。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 回滚云硬盘备份点前是否自动关机，默认为FALSE，表示不自动关机
	AutoStopInstance *bool `json:"AutoStopInstance,omitnil,omitempty" name:"AutoStopInstance"`

	// 回滚云硬盘备份点完成后是否自动开机，默认为FALSE，表示不自动开机; AutoStartInstance参数需要在AutoStopInstance为true时才能为true。
	AutoStartInstance *bool `json:"AutoStartInstance,omitnil,omitempty" name:"AutoStartInstance"`
}

type ApplyDiskBackupRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘备份点ID，可以通过[DescribeDiskBackups](/document/product/362/80278)接口查询。
	DiskBackupId *string `json:"DiskBackupId,omitnil,omitempty" name:"DiskBackupId"`

	// 云硬盘备份点原云硬盘ID，可以通过[DescribeDisks](/document/product/362/16315)接口查。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 回滚云硬盘备份点前是否自动关机，默认为FALSE，表示不自动关机
	AutoStopInstance *bool `json:"AutoStopInstance,omitnil,omitempty" name:"AutoStopInstance"`

	// 回滚云硬盘备份点完成后是否自动开机，默认为FALSE，表示不自动开机; AutoStartInstance参数需要在AutoStopInstance为true时才能为true。
	AutoStartInstance *bool `json:"AutoStartInstance,omitnil,omitempty" name:"AutoStartInstance"`
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
	delete(f, "DiskBackupId")
	delete(f, "DiskId")
	delete(f, "AutoStopInstance")
	delete(f, "AutoStartInstance")
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
type ApplySnapshotGroupRequestParams struct {
	// 回滚的快照组ID。
	SnapshotGroupId *string `json:"SnapshotGroupId,omitnil,omitempty" name:"SnapshotGroupId"`

	// 回滚的快照组关联的快照ID，及快照对应的原云硬盘ID列表。
	ApplyDisks []*ApplyDisk `json:"ApplyDisks,omitnil,omitempty" name:"ApplyDisks"`

	// 回滚前是否执行自动关机。
	AutoStopInstance *bool `json:"AutoStopInstance,omitnil,omitempty" name:"AutoStopInstance"`

	// 回滚完成后是否自动开机。
	AutoStartInstance *bool `json:"AutoStartInstance,omitnil,omitempty" name:"AutoStartInstance"`
}

type ApplySnapshotGroupRequest struct {
	*tchttp.BaseRequest
	
	// 回滚的快照组ID。
	SnapshotGroupId *string `json:"SnapshotGroupId,omitnil,omitempty" name:"SnapshotGroupId"`

	// 回滚的快照组关联的快照ID，及快照对应的原云硬盘ID列表。
	ApplyDisks []*ApplyDisk `json:"ApplyDisks,omitnil,omitempty" name:"ApplyDisks"`

	// 回滚前是否执行自动关机。
	AutoStopInstance *bool `json:"AutoStopInstance,omitnil,omitempty" name:"AutoStopInstance"`

	// 回滚完成后是否自动开机。
	AutoStartInstance *bool `json:"AutoStartInstance,omitnil,omitempty" name:"AutoStartInstance"`
}

func (r *ApplySnapshotGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplySnapshotGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotGroupId")
	delete(f, "ApplyDisks")
	delete(f, "AutoStopInstance")
	delete(f, "AutoStartInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplySnapshotGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplySnapshotGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplySnapshotGroupResponse struct {
	*tchttp.BaseResponse
	Response *ApplySnapshotGroupResponseParams `json:"Response"`
}

func (r *ApplySnapshotGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplySnapshotGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplySnapshotRequestParams struct {
	// 快照ID, 可通过[DescribeSnapshots](/document/product/362/15647)查询。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 快照原云硬盘ID，可通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 回滚前是否执行自动关机，仅支持回滚快照至已挂载的云硬盘时传入。
	// 此参数为true时，AutoStartInstance才能为true。
	AutoStopInstance *bool `json:"AutoStopInstance,omitnil,omitempty" name:"AutoStopInstance"`

	// 回滚完成后是否自动开机，仅支持回滚快照至已挂载的云硬盘时传入。该参数传入时，需要同时传入AutoStopInstance参数。
	AutoStartInstance *bool `json:"AutoStartInstance,omitnil,omitempty" name:"AutoStartInstance"`
}

type ApplySnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 快照ID, 可通过[DescribeSnapshots](/document/product/362/15647)查询。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 快照原云硬盘ID，可通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 回滚前是否执行自动关机，仅支持回滚快照至已挂载的云硬盘时传入。
	// 此参数为true时，AutoStartInstance才能为true。
	AutoStopInstance *bool `json:"AutoStopInstance,omitnil,omitempty" name:"AutoStopInstance"`

	// 回滚完成后是否自动开机，仅支持回滚快照至已挂载的云硬盘时传入。该参数传入时，需要同时传入AutoStopInstance参数。
	AutoStartInstance *bool `json:"AutoStartInstance,omitnil,omitempty" name:"AutoStartInstance"`
}

func (r *ApplySnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplySnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotId")
	delete(f, "DiskId")
	delete(f, "AutoStopInstance")
	delete(f, "AutoStartInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplySnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplySnapshotResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplySnapshotResponse struct {
	*tchttp.BaseResponse
	Response *ApplySnapshotResponseParams `json:"Response"`
}

func (r *ApplySnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplySnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachDetail struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例已挂载数据盘的数量。
	AttachedDiskCount *uint64 `json:"AttachedDiskCount,omitnil,omitempty" name:"AttachedDiskCount"`

	// 实例最大可挂载数据盘的数量。
	MaxAttachCount *uint64 `json:"MaxAttachCount,omitnil,omitempty" name:"MaxAttachCount"`
}

// Predefined struct for user
type AttachDisksRequestParams struct {
	// 云服务器实例ID。云盘将被挂载到此云服务器上，通过[DescribeInstances](/document/product/213/15728)接口查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 将要被挂载的弹性云盘ID。通过[DescribeDisks](/document/product/362/16315)接口查询。单次最多可挂载10块弹性云盘。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 可选参数，不传该参数则仅执行挂载操作。传入`True`时，会在挂载成功后将云硬盘设置为随云主机销毁模式，仅对按量计费云硬盘有效。
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitnil,omitempty" name:"DeleteWithInstance"`

	// 可选参数，用于控制云盘挂载时使用的挂载模式，目前仅对黑石裸金属机型有效。取值范围：<br><li>PF</li><br><li>VF</li>
	AttachMode *string `json:"AttachMode,omitnil,omitempty" name:"AttachMode"`
}

type AttachDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云服务器实例ID。云盘将被挂载到此云服务器上，通过[DescribeInstances](/document/product/213/15728)接口查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 将要被挂载的弹性云盘ID。通过[DescribeDisks](/document/product/362/16315)接口查询。单次最多可挂载10块弹性云盘。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 可选参数，不传该参数则仅执行挂载操作。传入`True`时，会在挂载成功后将云硬盘设置为随云主机销毁模式，仅对按量计费云硬盘有效。
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitnil,omitempty" name:"DeleteWithInstance"`

	// 可选参数，用于控制云盘挂载时使用的挂载模式，目前仅对黑石裸金属机型有效。取值范围：<br><li>PF</li><br><li>VF</li>
	AttachMode *string `json:"AttachMode,omitnil,omitempty" name:"AttachMode"`
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
	delete(f, "InstanceId")
	delete(f, "DiskIds")
	delete(f, "DeleteWithInstance")
	delete(f, "AttachMode")
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
	// 要挂载到的实例ID。
	InstanceId []*string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 子机内的挂载点。
	MountPoint []*string `json:"MountPoint,omitnil,omitempty" name:"MountPoint"`

	// 文件系统类型，支持的有 ext4、xfs。
	FileSystemType *string `json:"FileSystemType,omitnil,omitempty" name:"FileSystemType"`
}

type AutoSnapshotPolicy struct {
	// 已绑定当前定期快照策略的云盘ID列表。
	// DescribeDiskAssociatedAutoSnapshotPolicy场景下该字段返回为空。
	DiskIdSet []*string `json:"DiskIdSet,omitnil,omitempty" name:"DiskIdSet"`

	// 定期快照策略是否激活。
	IsActivated *bool `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// 定期快照策略的状态。取值范围：
	// <ul>
	//   <li>NORMAL：正常</li>
	//   <li>ISOLATED：已隔离</li>
	// </ul>
	AutoSnapshotPolicyState *string `json:"AutoSnapshotPolicyState,omitnil,omitempty" name:"AutoSnapshotPolicyState"`

	// 是否是跨账号复制快照, 1：是, 0: 不是
	IsCopyToRemote *uint64 `json:"IsCopyToRemote,omitnil,omitempty" name:"IsCopyToRemote"`

	// 使用该定期快照策略创建出来的快照是否永久保留。
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 定期快照下次触发的时间。
	NextTriggerTime *string `json:"NextTriggerTime,omitnil,omitempty" name:"NextTriggerTime"`

	// 定期快照策略名称。
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitnil,omitempty" name:"AutoSnapshotPolicyName"`

	// 定期快照策略ID。
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 定期快照的执行策略。
	Policy []*Policy `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 定期快照策略的创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 使用该定期快照策略创建出来的快照保留天数。
	RetentionDays *uint64 `json:"RetentionDays,omitnil,omitempty" name:"RetentionDays"`

	// 复制的目标账户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CopyToAccountUin *string `json:"CopyToAccountUin,omitnil,omitempty" name:"CopyToAccountUin"`

	// 已绑定当前定期快照策略的实例ID列表。
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`

	// 该定期快照创建的快照可以保留的月数。
	RetentionMonths *uint64 `json:"RetentionMonths,omitnil,omitempty" name:"RetentionMonths"`

	// 该定期快照创建的快照最大保留数量。
	RetentionAmount *uint64 `json:"RetentionAmount,omitnil,omitempty" name:"RetentionAmount"`

	// 定期快照高级保留策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedRetentionPolicy *AdvancedRetentionPolicy `json:"AdvancedRetentionPolicy,omitnil,omitempty" name:"AdvancedRetentionPolicy"`

	// 该复制快照策略的源端账户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CopyFromAccountUin *string `json:"CopyFromAccountUin,omitnil,omitempty" name:"CopyFromAccountUin"`

	// 标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

// Predefined struct for user
type BindAutoSnapshotPolicyRequestParams struct {
	// 要绑定的定期快照策略ID，通过[ DescribeAutoSnapshotPolicies](https://cloud.tencent.com/document/api/362/33556)接口查询。
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 要绑定的云硬盘ID列表，一次请求最多绑定80块云盘。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`
}

type BindAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 要绑定的定期快照策略ID，通过[ DescribeAutoSnapshotPolicies](https://cloud.tencent.com/document/api/362/33556)接口查询。
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 要绑定的云硬盘ID列表，一次请求最多绑定80块云盘。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`
}

func (r *BindAutoSnapshotPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoSnapshotPolicyId")
	delete(f, "DiskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindAutoSnapshotPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *BindAutoSnapshotPolicyResponseParams `json:"Response"`
}

func (r *BindAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Cdc struct {
	// 独享集群围笼ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// 独享集群状态。取值范围：<br><li>NORMAL：正常；</li><br><li>CLOSED：关闭，此时将不可使用该独享集群创建新的云硬盘；</li><br><li>FAULT：独享集群状态异常，此时独享集群将不可操作，腾讯云运维团队将会及时修复该集群；</li><br><li>ISOLATED：因未及时续费导致独享集群被隔离，此时将不可使用该独享集群创建新的云硬盘，对应的云硬盘也将不可操作。</li>
	CdcState *string `json:"CdcState,omitnil,omitempty" name:"CdcState"`

	// 独享集群所属的[可用区](/document/product/213/15753#ZoneInfo)ID。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 独享集群实例名称。
	CdcName *string `json:"CdcName,omitnil,omitempty" name:"CdcName"`

	// 独享集群的资源容量大小。
	CdcResource *CdcSize `json:"CdcResource,omitnil,omitempty" name:"CdcResource"`

	// 独享集群实例id。
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// 独享集群类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬盘集群</li><br><li>CLOUD_PREMIUM：表示高性能云硬盘集群</li><br><li>CLOUD_SSD：SSD表示SSD云硬盘集群。</li>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 独享集群到期时间。
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 存储池创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 当前集群中已创建的云盘数量。
	DiskNumber *uint64 `json:"DiskNumber,omitnil,omitempty" name:"DiskNumber"`
}

type CdcSize struct {
	// 独享集群的总容量大小，单位GiB
	DiskTotal *uint64 `json:"DiskTotal,omitnil,omitempty" name:"DiskTotal"`

	// 独享集群的可用容量大小，单位GiB
	DiskAvailable *uint64 `json:"DiskAvailable,omitnil,omitempty" name:"DiskAvailable"`
}

// Predefined struct for user
type CopySnapshotCrossRegionsRequestParams struct {
	// 快照需要复制到的目标地域，各地域的标准取值可通过接口[DescribeRegions](https://cloud.tencent.com/document/product/213/9456)查询，且只能传入支持快照的地域。
	DestinationRegions []*string `json:"DestinationRegions,omitnil,omitempty" name:"DestinationRegions"`

	// 需要跨地域复制的源快照ID，可通过[DescribeSnapshots](/document/product/362/15647)查询。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 新复制快照的名称，如果不传，则默认取值为“Copied 源快照ID from 地域名”。
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`
}

type CopySnapshotCrossRegionsRequest struct {
	*tchttp.BaseRequest
	
	// 快照需要复制到的目标地域，各地域的标准取值可通过接口[DescribeRegions](https://cloud.tencent.com/document/product/213/9456)查询，且只能传入支持快照的地域。
	DestinationRegions []*string `json:"DestinationRegions,omitnil,omitempty" name:"DestinationRegions"`

	// 需要跨地域复制的源快照ID，可通过[DescribeSnapshots](/document/product/362/15647)查询。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 新复制快照的名称，如果不传，则默认取值为“Copied 源快照ID from 地域名”。
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`
}

func (r *CopySnapshotCrossRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopySnapshotCrossRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DestinationRegions")
	delete(f, "SnapshotId")
	delete(f, "SnapshotName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopySnapshotCrossRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopySnapshotCrossRegionsResponseParams struct {
	// 快照跨地域复制的结果，如果请求下发成功，则返回相应地地域的新快照ID，否则返回Error。
	SnapshotCopyResultSet []*SnapshotCopyResult `json:"SnapshotCopyResultSet,omitnil,omitempty" name:"SnapshotCopyResultSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CopySnapshotCrossRegionsResponse struct {
	*tchttp.BaseResponse
	Response *CopySnapshotCrossRegionsResponseParams `json:"Response"`
}

func (r *CopySnapshotCrossRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopySnapshotCrossRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoSnapshotPolicyRequestParams struct {
	// 定期快照的执行策略。
	Policy []*Policy `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 是否创建定期快照的执行策略。TRUE表示只需获取首次开始备份的时间，不实际创建定期快照策略，FALSE表示创建，默认为FALSE。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 是否激活定期快照策略，FALSE表示未激活，TRUE表示激活，默认为TRUE。
	IsActivated *bool `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// 要创建的定期快照策略名。不传则默认为“未命名”。最大长度不能超60个字节。
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitnil,omitempty" name:"AutoSnapshotPolicyName"`

	// 通过该定期快照策略创建的快照是否永久保留。FALSE表示非永久保留，TRUE表示永久保留，默认为FALSE。
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 通过该定期快照策略创建的快照保留天数，默认保留7天。如果指定本参数，则IsPermanent入参不可指定为TRUE，否则会产生冲突。
	RetentionDays *uint64 `json:"RetentionDays,omitnil,omitempty" name:"RetentionDays"`
}

type CreateAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 定期快照的执行策略。
	Policy []*Policy `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 是否创建定期快照的执行策略。TRUE表示只需获取首次开始备份的时间，不实际创建定期快照策略，FALSE表示创建，默认为FALSE。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 是否激活定期快照策略，FALSE表示未激活，TRUE表示激活，默认为TRUE。
	IsActivated *bool `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// 要创建的定期快照策略名。不传则默认为“未命名”。最大长度不能超60个字节。
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitnil,omitempty" name:"AutoSnapshotPolicyName"`

	// 通过该定期快照策略创建的快照是否永久保留。FALSE表示非永久保留，TRUE表示永久保留，默认为FALSE。
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 通过该定期快照策略创建的快照保留天数，默认保留7天。如果指定本参数，则IsPermanent入参不可指定为TRUE，否则会产生冲突。
	RetentionDays *uint64 `json:"RetentionDays,omitnil,omitempty" name:"RetentionDays"`
}

func (r *CreateAutoSnapshotPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Policy")
	delete(f, "DryRun")
	delete(f, "IsActivated")
	delete(f, "AutoSnapshotPolicyName")
	delete(f, "IsPermanent")
	delete(f, "RetentionDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoSnapshotPolicyResponseParams struct {
	// 新创建的定期快照策略ID。
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 首次开始备份的时间。
	NextTriggerTime *string `json:"NextTriggerTime,omitnil,omitempty" name:"NextTriggerTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateAutoSnapshotPolicyResponseParams `json:"Response"`
}

func (r *CreateAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDiskBackupRequestParams struct {
	// 要创建备份点的云硬盘名称。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 云硬盘备份点名称。长度不能超过100个字符。
	DiskBackupName *string `json:"DiskBackupName,omitnil,omitempty" name:"DiskBackupName"`
}

type CreateDiskBackupRequest struct {
	*tchttp.BaseRequest
	
	// 要创建备份点的云硬盘名称。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 云硬盘备份点名称。长度不能超过100个字符。
	DiskBackupName *string `json:"DiskBackupName,omitnil,omitempty" name:"DiskBackupName"`
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
	// 云硬盘备份点的ID。
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
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目。若不指定项目，将在默认项目下进行创建。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 云硬盘计费类型。<br><li>PREPAID：预付费，即包年包月</li><br><li>POSTPAID_BY_HOUR：按小时后付费</li><br><li>CDCPAID：独享集群付费<br>各类型价格请参考云硬盘[价格总览](/document/product/362/2413)。</li>
	DiskChargeType *string `json:"DiskChargeType,omitnil,omitempty" name:"DiskChargeType"`

	// 硬盘介质类型。取值范围：<br><li>CLOUD_PREMIUM：表示高性能云硬盘</li><br><li>CLOUD_BSSD：表示通用型SSD云硬盘</li><br><li>CLOUD_SSD：表示SSD云硬盘</li><br><li>CLOUD_HSSD：表示增强型SSD云硬盘</li><br><li>CLOUD_TSSD：表示极速型SSD云硬盘。</li>极速型SSD云硬盘（CLOUD_TSSD）仅支持随部分实例类型一同购买，暂不支持单独创建。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 云盘显示名称。不传则默认为“未命名”。最大长度不能超60个字节。
	DiskName *string `json:"DiskName,omitnil,omitempty" name:"DiskName"`

	// 云盘绑定的标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 快照ID，如果传入则根据此快照创建云硬盘，快照类型必须为数据盘快照，可通过[DescribeSnapshots](/document/product/362/15647)接口查询快照，见输出参数DiskUsage解释。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 创建云硬盘数量，不传则默认为1。单次请求最多可创建的云盘数有限制，具体参见[云硬盘使用限制](https://cloud.tencent.com/doc/product/362/5145)。
	DiskCount *uint64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// 使用此参数可给云硬盘购买额外的性能，单位MB/s。<br>当前仅支持极速型云盘（CLOUD_TSSD）和增强型SSD云硬盘（CLOUD_HSSD）。
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitnil,omitempty" name:"ThroughputPerformance"`

	// 购买加密盘时自定义密钥，当传入该参数时，Encrypt参数不得为空。
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`

	// 云硬盘大小，单位为GiB。<br><li>如果传入`SnapshotId`则可不传`DiskSize`，此时新建云盘的大小为快照大小</li><br><li>如果传入`SnapshotId`同时传入`DiskSize`，则云盘大小必须大于或等于快照大小</li><br><li>云盘大小取值范围参见云硬盘[产品分类](/document/product/362/2353)的说明。</li>
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 传入True时，云盘将创建为共享型云盘，默认为False。因共享型云盘不支持加密，此参数与Encrypt参数不可同时传入。
	Shareable *bool `json:"Shareable,omitnil,omitempty" name:"Shareable"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 传入该参数用于创建加密云盘，取值固定为ENCRYPT。因共享型云盘不支持加密，此参数与Shareable参数不可同时传入。
	Encrypt *string `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`

	// 预付费模式，即包年包月相关参数设置。通过该参数指定包年包月云盘的购买时长、是否设置自动续费等属性。<br>创建预付费云盘该参数必传，创建按小时后付费云盘无需传该参数。
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitnil,omitempty" name:"DiskChargePrepaid"`

	// 销毁云盘时删除关联的非永久保留快照。0 表示非永久快照不随云盘销毁而销毁，1表示非永久快照随云盘销毁而销毁，默认取0。快照是否永久保留可以通过[DescribeSnapshots](/document/api/362/15647)接口返回的快照详情的IsPermanent字段来判断，True表示永久快照，False表示非永久快照。
	DeleteSnapshot *int64 `json:"DeleteSnapshot,omitnil,omitempty" name:"DeleteSnapshot"`

	// 创建云盘时指定自动挂载并初始化该数据盘。因加密盘不支持自动挂载及初始化，此参数与Encrypt参数不可同时传入。
	AutoMountConfiguration *AutoMountConfiguration `json:"AutoMountConfiguration,omitnil,omitempty" name:"AutoMountConfiguration"`

	// 指定云硬盘备份点配额。
	DiskBackupQuota *uint64 `json:"DiskBackupQuota,omitnil,omitempty" name:"DiskBackupQuota"`

	// 创建云盘时是否开启性能突发。当前仅支持极速型云盘（CLOUD_TSSD）和增强型SSD云硬盘（CLOUD_HSSD）且云盘大小不小于460GiB。
	BurstPerformance *bool `json:"BurstPerformance,omitnil,omitempty" name:"BurstPerformance"`

	// 指定云硬盘加密类型，取值为ENCRYPT_V1和ENCRYPT_V2，分别表示第一代和第二代加密技术，两种加密技术互不兼容。推荐优先使用第二代加密技术ENCRYPT_V2，第一代加密技术仅支持在部分老旧机型使用。该参数仅当创建加密云硬盘时有效。
	EncryptType *string `json:"EncryptType,omitnil,omitempty" name:"EncryptType"`
}

type CreateDisksRequest struct {
	*tchttp.BaseRequest
	
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目。若不指定项目，将在默认项目下进行创建。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 云硬盘计费类型。<br><li>PREPAID：预付费，即包年包月</li><br><li>POSTPAID_BY_HOUR：按小时后付费</li><br><li>CDCPAID：独享集群付费<br>各类型价格请参考云硬盘[价格总览](/document/product/362/2413)。</li>
	DiskChargeType *string `json:"DiskChargeType,omitnil,omitempty" name:"DiskChargeType"`

	// 硬盘介质类型。取值范围：<br><li>CLOUD_PREMIUM：表示高性能云硬盘</li><br><li>CLOUD_BSSD：表示通用型SSD云硬盘</li><br><li>CLOUD_SSD：表示SSD云硬盘</li><br><li>CLOUD_HSSD：表示增强型SSD云硬盘</li><br><li>CLOUD_TSSD：表示极速型SSD云硬盘。</li>极速型SSD云硬盘（CLOUD_TSSD）仅支持随部分实例类型一同购买，暂不支持单独创建。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 云盘显示名称。不传则默认为“未命名”。最大长度不能超60个字节。
	DiskName *string `json:"DiskName,omitnil,omitempty" name:"DiskName"`

	// 云盘绑定的标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 快照ID，如果传入则根据此快照创建云硬盘，快照类型必须为数据盘快照，可通过[DescribeSnapshots](/document/product/362/15647)接口查询快照，见输出参数DiskUsage解释。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 创建云硬盘数量，不传则默认为1。单次请求最多可创建的云盘数有限制，具体参见[云硬盘使用限制](https://cloud.tencent.com/doc/product/362/5145)。
	DiskCount *uint64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// 使用此参数可给云硬盘购买额外的性能，单位MB/s。<br>当前仅支持极速型云盘（CLOUD_TSSD）和增强型SSD云硬盘（CLOUD_HSSD）。
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitnil,omitempty" name:"ThroughputPerformance"`

	// 购买加密盘时自定义密钥，当传入该参数时，Encrypt参数不得为空。
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`

	// 云硬盘大小，单位为GiB。<br><li>如果传入`SnapshotId`则可不传`DiskSize`，此时新建云盘的大小为快照大小</li><br><li>如果传入`SnapshotId`同时传入`DiskSize`，则云盘大小必须大于或等于快照大小</li><br><li>云盘大小取值范围参见云硬盘[产品分类](/document/product/362/2353)的说明。</li>
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 传入True时，云盘将创建为共享型云盘，默认为False。因共享型云盘不支持加密，此参数与Encrypt参数不可同时传入。
	Shareable *bool `json:"Shareable,omitnil,omitempty" name:"Shareable"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 传入该参数用于创建加密云盘，取值固定为ENCRYPT。因共享型云盘不支持加密，此参数与Shareable参数不可同时传入。
	Encrypt *string `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`

	// 预付费模式，即包年包月相关参数设置。通过该参数指定包年包月云盘的购买时长、是否设置自动续费等属性。<br>创建预付费云盘该参数必传，创建按小时后付费云盘无需传该参数。
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitnil,omitempty" name:"DiskChargePrepaid"`

	// 销毁云盘时删除关联的非永久保留快照。0 表示非永久快照不随云盘销毁而销毁，1表示非永久快照随云盘销毁而销毁，默认取0。快照是否永久保留可以通过[DescribeSnapshots](/document/api/362/15647)接口返回的快照详情的IsPermanent字段来判断，True表示永久快照，False表示非永久快照。
	DeleteSnapshot *int64 `json:"DeleteSnapshot,omitnil,omitempty" name:"DeleteSnapshot"`

	// 创建云盘时指定自动挂载并初始化该数据盘。因加密盘不支持自动挂载及初始化，此参数与Encrypt参数不可同时传入。
	AutoMountConfiguration *AutoMountConfiguration `json:"AutoMountConfiguration,omitnil,omitempty" name:"AutoMountConfiguration"`

	// 指定云硬盘备份点配额。
	DiskBackupQuota *uint64 `json:"DiskBackupQuota,omitnil,omitempty" name:"DiskBackupQuota"`

	// 创建云盘时是否开启性能突发。当前仅支持极速型云盘（CLOUD_TSSD）和增强型SSD云硬盘（CLOUD_HSSD）且云盘大小不小于460GiB。
	BurstPerformance *bool `json:"BurstPerformance,omitnil,omitempty" name:"BurstPerformance"`

	// 指定云硬盘加密类型，取值为ENCRYPT_V1和ENCRYPT_V2，分别表示第一代和第二代加密技术，两种加密技术互不兼容。推荐优先使用第二代加密技术ENCRYPT_V2，第一代加密技术仅支持在部分老旧机型使用。该参数仅当创建加密云硬盘时有效。
	EncryptType *string `json:"EncryptType,omitnil,omitempty" name:"EncryptType"`
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
	delete(f, "Placement")
	delete(f, "DiskChargeType")
	delete(f, "DiskType")
	delete(f, "DiskName")
	delete(f, "Tags")
	delete(f, "SnapshotId")
	delete(f, "DiskCount")
	delete(f, "ThroughputPerformance")
	delete(f, "KmsKeyId")
	delete(f, "DiskSize")
	delete(f, "Shareable")
	delete(f, "ClientToken")
	delete(f, "Encrypt")
	delete(f, "DiskChargePrepaid")
	delete(f, "DeleteSnapshot")
	delete(f, "AutoMountConfiguration")
	delete(f, "DiskBackupQuota")
	delete(f, "BurstPerformance")
	delete(f, "EncryptType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDisksResponseParams struct {
	// 创建的云硬盘ID列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
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
type CreateSnapshotGroupRequestParams struct {
	// 需要创建快照组的云硬盘ID列表，必须选择挂载在同一实例上的盘列表。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 快照组名称，快照组关联的快照也会继承快照组的名称。例如：快照组名称为testSnapshotGroup，快照组关联两个快照，则两个快照的名称分别为testSnapshotGroup_0，testSnapshotGroup_1。
	SnapshotGroupName *string `json:"SnapshotGroupName,omitnil,omitempty" name:"SnapshotGroupName"`

	// 快照组需要绑定的标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateSnapshotGroupRequest struct {
	*tchttp.BaseRequest
	
	// 需要创建快照组的云硬盘ID列表，必须选择挂载在同一实例上的盘列表。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 快照组名称，快照组关联的快照也会继承快照组的名称。例如：快照组名称为testSnapshotGroup，快照组关联两个快照，则两个快照的名称分别为testSnapshotGroup_0，testSnapshotGroup_1。
	SnapshotGroupName *string `json:"SnapshotGroupName,omitnil,omitempty" name:"SnapshotGroupName"`

	// 快照组需要绑定的标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateSnapshotGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "SnapshotGroupName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSnapshotGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotGroupResponseParams struct {
	// 创建成功的快照组ID。
	SnapshotGroupId *string `json:"SnapshotGroupId,omitnil,omitempty" name:"SnapshotGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSnapshotGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateSnapshotGroupResponseParams `json:"Response"`
}

func (r *CreateSnapshotGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotRequestParams struct {
	// 需要创建快照的云硬盘ID，可通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 快照名称，不传则新快照名称默认为“未命名”。
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 快照的到期时间，到期后该快照将会自动删除,需要传入UTC时间下的ISO-8601标准时间格式,例如:2022-01-08T09:47:55+00:00,。到期时间最小可设置为一天后的当前时间。
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 云硬盘备份点ID。传入此参数时，将通过备份点创建快照。备份点 ID 可以通过[DescribeDiskBackups](/document/product/362/80278)接口查询。
	DiskBackupId *string `json:"DiskBackupId,omitnil,omitempty" name:"DiskBackupId"`

	// 快照绑定的标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 快照关联云硬盘类型, SYSTEM_DISK: 系统盘, DATA_DISK: 数据盘,非必填参数，不填时快照类型与云盘类型保持一致， 该参数基于某些场景用户需要将系统盘创建出数据盘快照共享使用。
	DiskUsage *string `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`
}

type CreateSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 需要创建快照的云硬盘ID，可通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 快照名称，不传则新快照名称默认为“未命名”。
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 快照的到期时间，到期后该快照将会自动删除,需要传入UTC时间下的ISO-8601标准时间格式,例如:2022-01-08T09:47:55+00:00,。到期时间最小可设置为一天后的当前时间。
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 云硬盘备份点ID。传入此参数时，将通过备份点创建快照。备份点 ID 可以通过[DescribeDiskBackups](/document/product/362/80278)接口查询。
	DiskBackupId *string `json:"DiskBackupId,omitnil,omitempty" name:"DiskBackupId"`

	// 快照绑定的标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 快照关联云硬盘类型, SYSTEM_DISK: 系统盘, DATA_DISK: 数据盘,非必填参数，不填时快照类型与云盘类型保持一致， 该参数基于某些场景用户需要将系统盘创建出数据盘快照共享使用。
	DiskUsage *string `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`
}

func (r *CreateSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	delete(f, "SnapshotName")
	delete(f, "Deadline")
	delete(f, "DiskBackupId")
	delete(f, "Tags")
	delete(f, "DiskUsage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotResponseParams struct {
	// 新创建的快照ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *CreateSnapshotResponseParams `json:"Response"`
}

func (r *CreateSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAutoSnapshotPoliciesRequestParams struct {
	// 要删除的定期快照策略ID列表，通过[ DescribeAutoSnapshotPolicies](https://cloud.tencent.com/document/api/362/33556)接口查询。
	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitnil,omitempty" name:"AutoSnapshotPolicyIds"`
}

type DeleteAutoSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的定期快照策略ID列表，通过[ DescribeAutoSnapshotPolicies](https://cloud.tencent.com/document/api/362/33556)接口查询。
	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitnil,omitempty" name:"AutoSnapshotPolicyIds"`
}

func (r *DeleteAutoSnapshotPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoSnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoSnapshotPolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAutoSnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAutoSnapshotPoliciesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAutoSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAutoSnapshotPoliciesResponseParams `json:"Response"`
}

func (r *DeleteAutoSnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDiskBackupsRequestParams struct {
	// 待删除的云硬盘备份点ID，可以通过[DescribeDiskBackups](/document/product/362/80278)接口查询。
	DiskBackupIds []*string `json:"DiskBackupIds,omitnil,omitempty" name:"DiskBackupIds"`
}

type DeleteDiskBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的云硬盘备份点ID，可以通过[DescribeDiskBackups](/document/product/362/80278)接口查询。
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
type DeleteSnapshotGroupRequestParams struct {
	// 快照组ID。
	SnapshotGroupId *string `json:"SnapshotGroupId,omitnil,omitempty" name:"SnapshotGroupId"`

	// 快照组ID 列表。此参数与快照组 ID 至少传 1 个，同时传会与快照组 ID 合并。
	SnapshotGroupIds []*string `json:"SnapshotGroupIds,omitnil,omitempty" name:"SnapshotGroupIds"`

	// 是否同时删除快照组关联的镜像；取值为false，表示不删除快照组绑定的镜像，此时，如果快照组有绑定的镜像，删除会失败；取值为true，表示同时删除快照组绑定的镜像；默认值为false。
	DeleteBindImages *bool `json:"DeleteBindImages,omitnil,omitempty" name:"DeleteBindImages"`
}

type DeleteSnapshotGroupRequest struct {
	*tchttp.BaseRequest
	
	// 快照组ID。
	SnapshotGroupId *string `json:"SnapshotGroupId,omitnil,omitempty" name:"SnapshotGroupId"`

	// 快照组ID 列表。此参数与快照组 ID 至少传 1 个，同时传会与快照组 ID 合并。
	SnapshotGroupIds []*string `json:"SnapshotGroupIds,omitnil,omitempty" name:"SnapshotGroupIds"`

	// 是否同时删除快照组关联的镜像；取值为false，表示不删除快照组绑定的镜像，此时，如果快照组有绑定的镜像，删除会失败；取值为true，表示同时删除快照组绑定的镜像；默认值为false。
	DeleteBindImages *bool `json:"DeleteBindImages,omitnil,omitempty" name:"DeleteBindImages"`
}

func (r *DeleteSnapshotGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotGroupId")
	delete(f, "SnapshotGroupIds")
	delete(f, "DeleteBindImages")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSnapshotGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSnapshotGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSnapshotGroupResponseParams `json:"Response"`
}

func (r *DeleteSnapshotGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotsRequestParams struct {
	// 要删除的快照ID列表，可通过[DescribeSnapshots](/document/product/362/15647)查询。
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`

	// 是否强制删除快照关联的镜像
	DeleteBindImages *bool `json:"DeleteBindImages,omitnil,omitempty" name:"DeleteBindImages"`
}

type DeleteSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的快照ID列表，可通过[DescribeSnapshots](/document/product/362/15647)查询。
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`

	// 是否强制删除快照关联的镜像
	DeleteBindImages *bool `json:"DeleteBindImages,omitnil,omitempty" name:"DeleteBindImages"`
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
	delete(f, "DeleteBindImages")
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

// Predefined struct for user
type DescribeAutoSnapshotPoliciesRequestParams struct {
	// 要查询的定期快照策略ID列表。参数不支持同时指定`AutoSnapshotPolicyIds`和`Filters`。
	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitnil,omitempty" name:"AutoSnapshotPolicyIds"`

	// 过滤条件。参数不支持同时指定`AutoSnapshotPolicyIds`和`Filters`。<br>
	// <li>auto-snapshot-policy-id - Array of String - 是否必填：否 -（过滤条件）按定期快照策略ID进行过滤。定期快照策略ID形如：`asp-3stvwfxx`。</li>
	// <li>auto-snapshot-policy-state - Array of String - 是否必填：否 -（过滤条件）按定期快照策略的状态进行过滤。定期快照策略ID形如：`asp-3stvwfxx`。(NORMAL：正常 | ISOLATED：已隔离。)</li>
	// <li>auto-snapshot-policy-name - Array of String - 是否必填：否 -（过滤条件）按定期快照策略名称进行过滤。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 输出定期快照列表的排列顺序。取值范围：<br><li>ASC：升序排列<br></li><li>DESC：降序排列。</li>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 定期快照列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据定期快照的创建时间排序，默认按创建时间排序。</li>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeAutoSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的定期快照策略ID列表。参数不支持同时指定`AutoSnapshotPolicyIds`和`Filters`。
	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitnil,omitempty" name:"AutoSnapshotPolicyIds"`

	// 过滤条件。参数不支持同时指定`AutoSnapshotPolicyIds`和`Filters`。<br>
	// <li>auto-snapshot-policy-id - Array of String - 是否必填：否 -（过滤条件）按定期快照策略ID进行过滤。定期快照策略ID形如：`asp-3stvwfxx`。</li>
	// <li>auto-snapshot-policy-state - Array of String - 是否必填：否 -（过滤条件）按定期快照策略的状态进行过滤。定期快照策略ID形如：`asp-3stvwfxx`。(NORMAL：正常 | ISOLATED：已隔离。)</li>
	// <li>auto-snapshot-policy-name - Array of String - 是否必填：否 -（过滤条件）按定期快照策略名称进行过滤。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 输出定期快照列表的排列顺序。取值范围：<br><li>ASC：升序排列<br></li><li>DESC：降序排列。</li>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 定期快照列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据定期快照的创建时间排序，默认按创建时间排序。</li>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeAutoSnapshotPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoSnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoSnapshotPolicyIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoSnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoSnapshotPoliciesResponseParams struct {
	// 有效的定期快照策略数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 定期快照策略列表。
	AutoSnapshotPolicySet []*AutoSnapshotPolicy `json:"AutoSnapshotPolicySet,omitnil,omitempty" name:"AutoSnapshotPolicySet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoSnapshotPoliciesResponseParams `json:"Response"`
}

func (r *DescribeAutoSnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskAssociatedAutoSnapshotPolicyRequestParams struct {
	// 要查询的云硬盘ID，通过[DescribeDisks](https://cloud.tencent.com/document/api/362/16315)接口查询。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`
}

type DescribeDiskAssociatedAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的云硬盘ID，通过[DescribeDisks](https://cloud.tencent.com/document/api/362/16315)接口查询。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`
}

func (r *DescribeDiskAssociatedAutoSnapshotPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskAssociatedAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiskAssociatedAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskAssociatedAutoSnapshotPolicyResponseParams struct {
	// 云盘绑定的定期快照数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 云盘绑定的定期快照列表。
	AutoSnapshotPolicySet []*AutoSnapshotPolicy `json:"AutoSnapshotPolicySet,omitnil,omitempty" name:"AutoSnapshotPolicySet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDiskAssociatedAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiskAssociatedAutoSnapshotPolicyResponseParams `json:"Response"`
}

func (r *DescribeDiskAssociatedAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskAssociatedAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskBackupsRequestParams struct {
	// 要查询备份点的ID列表。参数不支持同时指定 DiskBackupIds 和 Filters。
	DiskBackupIds []*string `json:"DiskBackupIds,omitnil,omitempty" name:"DiskBackupIds"`

	// 过滤条件，参数不支持同时指定 DiskBackupIds 和 Filters。过滤条件：<br><li>disk-backup-id - Array of String - 是否必填：否 -（过滤条件）按照备份点的ID过滤。备份点ID形如：dbp-11112222。</li><br><li>disk-id - Array of String - 是否必填：否 -（过滤条件）按照创建备份点的云硬盘ID过滤。云硬盘ID形如：disk-srftydert。</li><br><li>disk-usage - Array of String - 是否必填：否 -（过滤条件）按创建备份点的云硬盘类型过滤。 (SYSTEM_DISK：代表系统盘 | DATA_DISK：代表数据盘。)</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出云硬盘备份点列表的排列顺序，默认排序：ASC。取值范围：<br><li>ASC：升序排列</li><br><li>DESC：降序排列。</li>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 云硬盘备份点列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据云硬盘备份点的创建时间排序</li><br>默认按创建时间排序。
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeDiskBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 要查询备份点的ID列表。参数不支持同时指定 DiskBackupIds 和 Filters。
	DiskBackupIds []*string `json:"DiskBackupIds,omitnil,omitempty" name:"DiskBackupIds"`

	// 过滤条件，参数不支持同时指定 DiskBackupIds 和 Filters。过滤条件：<br><li>disk-backup-id - Array of String - 是否必填：否 -（过滤条件）按照备份点的ID过滤。备份点ID形如：dbp-11112222。</li><br><li>disk-id - Array of String - 是否必填：否 -（过滤条件）按照创建备份点的云硬盘ID过滤。云硬盘ID形如：disk-srftydert。</li><br><li>disk-usage - Array of String - 是否必填：否 -（过滤条件）按创建备份点的云硬盘类型过滤。 (SYSTEM_DISK：代表系统盘 | DATA_DISK：代表数据盘。)</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出云硬盘备份点列表的排列顺序，默认排序：ASC。取值范围：<br><li>ASC：升序排列</li><br><li>DESC：降序排列。</li>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 云硬盘备份点列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据云硬盘备份点的创建时间排序</li><br>默认按创建时间排序。
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
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
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiskBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskBackupsResponseParams struct {
	// 符合条件的云硬盘备份点数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 云硬盘备份点的详细信息列表。
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
type DescribeDiskConfigQuotaRequestParams struct {
	// 查询类别，取值范围。<br> INQUIRY_CBS_CONFIG：查询云盘配置列表<br> INQUIRY_CVM_CONFIG：查询云盘与实例搭配的配置列表。
	InquiryType *string `json:"InquiryType,omitnil,omitempty" name:"InquiryType"`

	// 付费模式。取值范围：<br> PREPAID：预付费<br> POSTPAID_BY_HOUR：后付费。
	DiskChargeType *string `json:"DiskChargeType,omitnil,omitempty" name:"DiskChargeType"`

	// 按照实例机型系列过滤。实例机型系列形如：S1、I1、M1等。详见[实例类型](https://cloud.tencent.com/document/product/213/11518)
	InstanceFamilies []*string `json:"InstanceFamilies,omitnil,omitempty" name:"InstanceFamilies"`

	// 硬盘介质类型。取值范围：<br> CLOUD_BASIC：表示普通云硬盘<br> CLOUD_PREMIUM：表示高性能云硬盘<br> CLOUD_SSD：表示SSD云硬盘<br> CLOUD_HSSD：表示增强型SSD云硬盘。
	DiskTypes []*string `json:"DiskTypes,omitnil,omitempty" name:"DiskTypes"`

	// 查询一个或多个[可用区](/document/product/213/15753#ZoneInfo)下的配置。
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 实例内存大小,单位GB。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 系统盘或数据盘。取值范围：<br> SYSTEM_DISK：表示系统盘<br> DATA_DISK：表示数据盘。
	DiskUsage *string `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// 实例CPU核数。
	CPU *uint64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// 专用集群ID。
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

type DescribeDiskConfigQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 查询类别，取值范围。<br> INQUIRY_CBS_CONFIG：查询云盘配置列表<br> INQUIRY_CVM_CONFIG：查询云盘与实例搭配的配置列表。
	InquiryType *string `json:"InquiryType,omitnil,omitempty" name:"InquiryType"`

	// 付费模式。取值范围：<br> PREPAID：预付费<br> POSTPAID_BY_HOUR：后付费。
	DiskChargeType *string `json:"DiskChargeType,omitnil,omitempty" name:"DiskChargeType"`

	// 按照实例机型系列过滤。实例机型系列形如：S1、I1、M1等。详见[实例类型](https://cloud.tencent.com/document/product/213/11518)
	InstanceFamilies []*string `json:"InstanceFamilies,omitnil,omitempty" name:"InstanceFamilies"`

	// 硬盘介质类型。取值范围：<br> CLOUD_BASIC：表示普通云硬盘<br> CLOUD_PREMIUM：表示高性能云硬盘<br> CLOUD_SSD：表示SSD云硬盘<br> CLOUD_HSSD：表示增强型SSD云硬盘。
	DiskTypes []*string `json:"DiskTypes,omitnil,omitempty" name:"DiskTypes"`

	// 查询一个或多个[可用区](/document/product/213/15753#ZoneInfo)下的配置。
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 实例内存大小,单位GB。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 系统盘或数据盘。取值范围：<br> SYSTEM_DISK：表示系统盘<br> DATA_DISK：表示数据盘。
	DiskUsage *string `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// 实例CPU核数。
	CPU *uint64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// 专用集群ID。
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

func (r *DescribeDiskConfigQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskConfigQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InquiryType")
	delete(f, "DiskChargeType")
	delete(f, "InstanceFamilies")
	delete(f, "DiskTypes")
	delete(f, "Zones")
	delete(f, "Memory")
	delete(f, "DiskUsage")
	delete(f, "CPU")
	delete(f, "DedicatedClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiskConfigQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskConfigQuotaResponseParams struct {
	// 云盘配置列表。
	DiskConfigSet []*DiskConfig `json:"DiskConfigSet,omitnil,omitempty" name:"DiskConfigSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDiskConfigQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiskConfigQuotaResponseParams `json:"Response"`
}

func (r *DescribeDiskConfigQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskConfigQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskStoragePoolRequestParams struct {
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 指定需要查询的独享集群ID列表，该入参不能与Filters一起使用。
	CdcIds []*string `json:"CdcIds,omitnil,omitempty" name:"CdcIds"`

	// 过滤条件。参数不支持同时指定`CdcIds`和`Filters`。<br><li>cdc-id - Array of String - 是否必填：否 -（过滤条件）按独享集群ID过滤。<br><li>zone - Array of String - 是否必填：否 -（过滤条件）按独享集群所在[可用区](/document/product/213/15753#ZoneInfo)过滤。<br><li>cage-id - Array of String - 是否必填：否 -（过滤条件）按独享集群所在围笼的ID过滤。<br><li>disk-type - Array of String - 是否必填：否 -（过滤条件）按照云盘介质类型过滤。(CLOUD_BASIC：表示普通云硬盘 | CLOUD_PREMIUM：表示高性能云硬盘。| CLOUD_SSD：SSD表示SSD云硬盘。)
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeDiskStoragePoolRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 指定需要查询的独享集群ID列表，该入参不能与Filters一起使用。
	CdcIds []*string `json:"CdcIds,omitnil,omitempty" name:"CdcIds"`

	// 过滤条件。参数不支持同时指定`CdcIds`和`Filters`。<br><li>cdc-id - Array of String - 是否必填：否 -（过滤条件）按独享集群ID过滤。<br><li>zone - Array of String - 是否必填：否 -（过滤条件）按独享集群所在[可用区](/document/product/213/15753#ZoneInfo)过滤。<br><li>cage-id - Array of String - 是否必填：否 -（过滤条件）按独享集群所在围笼的ID过滤。<br><li>disk-type - Array of String - 是否必填：否 -（过滤条件）按照云盘介质类型过滤。(CLOUD_BASIC：表示普通云硬盘 | CLOUD_PREMIUM：表示高性能云硬盘。| CLOUD_SSD：SSD表示SSD云硬盘。)
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeDiskStoragePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskStoragePoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "CdcIds")
	delete(f, "Filters")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiskStoragePoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskStoragePoolResponseParams struct {
	// 符合条件的独享集群的数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 独享集群的详细信息列表
	CdcSet []*Cdc `json:"CdcSet,omitnil,omitempty" name:"CdcSet"`

	// 独享集群的详细信息列表
	DiskStoragePoolSet []*Cdc `json:"DiskStoragePoolSet,omitnil,omitempty" name:"DiskStoragePoolSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDiskStoragePoolResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiskStoragePoolResponseParams `json:"Response"`
}

func (r *DescribeDiskStoragePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskStoragePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksRequestParams struct {
	// 过滤条件。参数不支持同时指定`DiskIds`和`Filters`。<br> <li>disk-usage - Array of String - 是否必填：否 -（过滤条件）按云盘类型过滤。 (SYSTEM_DISK：表示系统盘 | DATA_DISK：表示数据盘)<br></li> <li>disk-charge-type - Array of String - 是否必填：否 -（过滤条件）按照云硬盘计费模式过滤。 (PREPAID：表示预付费，即包年包月 | POSTPAID_BY_HOUR：表示后付费，即按量计费。)<br></li> <li>portable - Array of String - 是否必填：否 -（过滤条件）按是否为弹性云盘过滤。 (TRUE：表示弹性云盘 | FALSE：表示非弹性云盘。)<br></li> <li>project-id - Array of String - 是否必填：否 -（过滤条件）按云硬盘所属项目ID过滤。<br></li> <li>disk-id - Array of String - 是否必填：否 -（过滤条件）按照云硬盘ID过滤。云盘ID形如：`disk-11112222`。<br></li> <li>disk-name - Array of String - 是否必填：否 -（过滤条件）按照云盘名称过滤。<br></li> <li>disk-type - Array of String - 是否必填：否 -（过滤条件）按照云盘介质类型过滤。(CLOUD_BASIC：表示普通云硬盘 | CLOUD_PREMIUM：表示高性能云硬盘。| CLOUD_SSD：表示SSD云硬盘 | CLOUD_HSSD：表示增强型SSD云硬盘。| CLOUD_TSSD：表示极速型云硬盘。)<br></li> <li>disk-state - Array of String - 是否必填：否 -（过滤条件）按照云盘状态过滤。(UNATTACHED：未挂载 | ATTACHING：挂载中 | ATTACHED：已挂载 | DETACHING：解挂中 | EXPANDING：扩容中 | ROLLBACKING：回滚中 | TORECYCLE：待回收 | DUMPING：拷贝硬盘中。)<br></li> <li>instance-id - Array of String - 是否必填：否 -（过滤条件）按照云盘挂载的云主机实例ID过滤。可根据此参数查询挂载在指定云主机下的云硬盘。<br></li> <li>zone - Array of String - 是否必填：否 -（过滤条件）按照[可用区](/document/product/213/15753#ZoneInfo)过滤。<br></li> <li>instance-ip-address - Array of String - 是否必填：否 -（过滤条件）按云盘所挂载云主机的内网或外网IP过滤。<br></li> <li>instance-name - Array of String - 是否必填：否 -（过滤条件）按云盘所挂载的实例名称过滤。<br></li> <li>tag-key - Array of String - 是否必填：否 -（过滤条件）按照标签键进行过滤。<br></li> <li>tag-value - Array of String - 是否必填：否 -（过滤条件）照标签值进行过滤。<br></li> <li>tag:tag-key - Array of String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。<br></li> <li>dedicated-cluster-id - Array of String - 是否必填：否 -（过滤条件）按照 CDC 独享集群 ID 进行过滤。<br></li> <li>cluster-group-id - String - 是否必填：否 -（过滤条件）按照 集群群组 ID 进行过滤。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 云盘列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据云盘的创建时间排序<br></li><li>DEADLINE：依据云盘的到期时间排序<br>默认按云盘创建时间排序。</li>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 云盘详情中是否需要返回云盘绑定的定期快照策略ID，TRUE表示需要返回，FALSE表示不返回。
	ReturnBindAutoSnapshotPolicy *bool `json:"ReturnBindAutoSnapshotPolicy,omitnil,omitempty" name:"ReturnBindAutoSnapshotPolicy"`

	// 按照一个或者多个云硬盘ID查询。云硬盘ID形如：`disk-11112222`，此参数的具体格式可参考API[简介](/document/product/362/15633)的ids.N一节）。参数不支持同时指定`DiskIds`和`Filters`。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 输出云盘列表的排列顺序。取值范围：<br><li>ASC：升序排列<br></li><li>DESC：降序排列。</li>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeDisksRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。参数不支持同时指定`DiskIds`和`Filters`。<br> <li>disk-usage - Array of String - 是否必填：否 -（过滤条件）按云盘类型过滤。 (SYSTEM_DISK：表示系统盘 | DATA_DISK：表示数据盘)<br></li> <li>disk-charge-type - Array of String - 是否必填：否 -（过滤条件）按照云硬盘计费模式过滤。 (PREPAID：表示预付费，即包年包月 | POSTPAID_BY_HOUR：表示后付费，即按量计费。)<br></li> <li>portable - Array of String - 是否必填：否 -（过滤条件）按是否为弹性云盘过滤。 (TRUE：表示弹性云盘 | FALSE：表示非弹性云盘。)<br></li> <li>project-id - Array of String - 是否必填：否 -（过滤条件）按云硬盘所属项目ID过滤。<br></li> <li>disk-id - Array of String - 是否必填：否 -（过滤条件）按照云硬盘ID过滤。云盘ID形如：`disk-11112222`。<br></li> <li>disk-name - Array of String - 是否必填：否 -（过滤条件）按照云盘名称过滤。<br></li> <li>disk-type - Array of String - 是否必填：否 -（过滤条件）按照云盘介质类型过滤。(CLOUD_BASIC：表示普通云硬盘 | CLOUD_PREMIUM：表示高性能云硬盘。| CLOUD_SSD：表示SSD云硬盘 | CLOUD_HSSD：表示增强型SSD云硬盘。| CLOUD_TSSD：表示极速型云硬盘。)<br></li> <li>disk-state - Array of String - 是否必填：否 -（过滤条件）按照云盘状态过滤。(UNATTACHED：未挂载 | ATTACHING：挂载中 | ATTACHED：已挂载 | DETACHING：解挂中 | EXPANDING：扩容中 | ROLLBACKING：回滚中 | TORECYCLE：待回收 | DUMPING：拷贝硬盘中。)<br></li> <li>instance-id - Array of String - 是否必填：否 -（过滤条件）按照云盘挂载的云主机实例ID过滤。可根据此参数查询挂载在指定云主机下的云硬盘。<br></li> <li>zone - Array of String - 是否必填：否 -（过滤条件）按照[可用区](/document/product/213/15753#ZoneInfo)过滤。<br></li> <li>instance-ip-address - Array of String - 是否必填：否 -（过滤条件）按云盘所挂载云主机的内网或外网IP过滤。<br></li> <li>instance-name - Array of String - 是否必填：否 -（过滤条件）按云盘所挂载的实例名称过滤。<br></li> <li>tag-key - Array of String - 是否必填：否 -（过滤条件）按照标签键进行过滤。<br></li> <li>tag-value - Array of String - 是否必填：否 -（过滤条件）照标签值进行过滤。<br></li> <li>tag:tag-key - Array of String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。<br></li> <li>dedicated-cluster-id - Array of String - 是否必填：否 -（过滤条件）按照 CDC 独享集群 ID 进行过滤。<br></li> <li>cluster-group-id - String - 是否必填：否 -（过滤条件）按照 集群群组 ID 进行过滤。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 云盘列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据云盘的创建时间排序<br></li><li>DEADLINE：依据云盘的到期时间排序<br>默认按云盘创建时间排序。</li>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 云盘详情中是否需要返回云盘绑定的定期快照策略ID，TRUE表示需要返回，FALSE表示不返回。
	ReturnBindAutoSnapshotPolicy *bool `json:"ReturnBindAutoSnapshotPolicy,omitnil,omitempty" name:"ReturnBindAutoSnapshotPolicy"`

	// 按照一个或者多个云硬盘ID查询。云硬盘ID形如：`disk-11112222`，此参数的具体格式可参考API[简介](/document/product/362/15633)的ids.N一节）。参数不支持同时指定`DiskIds`和`Filters`。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 输出云盘列表的排列顺序。取值范围：<br><li>ASC：升序排列<br></li><li>DESC：降序排列。</li>
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
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Offset")
	delete(f, "ReturnBindAutoSnapshotPolicy")
	delete(f, "DiskIds")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksResponseParams struct {
	// 符合条件的云硬盘数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 云硬盘的详细信息列表。
	DiskSet []*Disk `json:"DiskSet,omitnil,omitempty" name:"DiskSet"`

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
type DescribeInstancesDiskNumRequestParams struct {
	// 云服务器实例ID，通过[DescribeInstances](/document/product/213/15728)接口查询。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeInstancesDiskNumRequest struct {
	*tchttp.BaseRequest
	
	// 云服务器实例ID，通过[DescribeInstances](/document/product/213/15728)接口查询。
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
	// 各个云服务器已挂载和可挂载弹性云盘的数量。
	AttachDetail []*AttachDetail `json:"AttachDetail,omitnil,omitempty" name:"AttachDetail"`

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
type DescribeSnapshotGroupsRequestParams struct {
	// 过滤条件。<br><li>snapshot-group-id - Array of String - 是否必填：否 -（过滤条件）按快照组ID过滤 <br><li>snapshot-group-state - Array of String - 是否必填：否 -（过滤条件）按快照组状态过滤。(NORMAL: 正常 | CREATING:创建中 | ROLLBACKING:回滚中) <br><li>snapshot-group-name - Array of String - 是否必填：否 -（过滤条件）按快照组名称过滤 <br><li>snapshot-id - Array of String - 是否必填：否 -（过滤条件）按快照组内的快照ID过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSnapshotGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。<br><li>snapshot-group-id - Array of String - 是否必填：否 -（过滤条件）按快照组ID过滤 <br><li>snapshot-group-state - Array of String - 是否必填：否 -（过滤条件）按快照组状态过滤。(NORMAL: 正常 | CREATING:创建中 | ROLLBACKING:回滚中) <br><li>snapshot-group-name - Array of String - 是否必填：否 -（过滤条件）按快照组名称过滤 <br><li>snapshot-id - Array of String - 是否必填：否 -（过滤条件）按快照组内的快照ID过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeSnapshotGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotGroupsResponseParams struct {
	// 符合条件的总数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 快照组列表详情。
	SnapshotGroupSet []*SnapshotGroup `json:"SnapshotGroupSet,omitnil,omitempty" name:"SnapshotGroupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSnapshotGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotGroupsResponseParams `json:"Response"`
}

func (r *DescribeSnapshotGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotOverviewRequestParams struct {

}

type DescribeSnapshotOverviewRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSnapshotOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotOverviewResponseParams struct {
	// 当前总有效快照数量
	TotalNums *uint64 `json:"TotalNums,omitnil,omitempty" name:"TotalNums"`

	// 已使用快照总容量大小，容量单位为GiB
	TotalSize *float64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`

	// 快照免费额度大小，额度单位为GiB
	FreeQuota *float64 `json:"FreeQuota,omitnil,omitempty" name:"FreeQuota"`

	// 快照真实产生计费的总容量大小，单位为GiB
	RealTradeSize *float64 `json:"RealTradeSize,omitnil,omitempty" name:"RealTradeSize"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSnapshotOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotOverviewResponseParams `json:"Response"`
}

func (r *DescribeSnapshotOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotSharePermissionRequestParams struct {
	// 要查询快照的ID。可通过[DescribeSnapshots](https://cloud.tencent.com/document/api/362/15647)查询获取。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`
}

type DescribeSnapshotSharePermissionRequest struct {
	*tchttp.BaseRequest
	
	// 要查询快照的ID。可通过[DescribeSnapshots](https://cloud.tencent.com/document/api/362/15647)查询获取。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`
}

func (r *DescribeSnapshotSharePermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotSharePermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotSharePermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotSharePermissionResponseParams struct {
	// 快照的分享信息的集合
	SharePermissionSet []*SharePermission `json:"SharePermissionSet,omitnil,omitempty" name:"SharePermissionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSnapshotSharePermissionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotSharePermissionResponseParams `json:"Response"`
}

func (r *DescribeSnapshotSharePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotSharePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotsRequestParams struct {
	// 要查询快照的ID列表。参数不支持同时指定`SnapshotIds`和`Filters`。
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`

	// 过滤条件。参数不支持同时指定SnapshotIds和Filters。<br><ul><li>snapshot-id<ul><li>按照云硬盘快照ID进行过滤</li><li>类型：String</li><li>必选：否</li></ul></li><li>snapshot-name<ul><li>按照云硬盘快照名称进行过滤</li><li>类型：String</li><li>必选：否</li></ul></li><li>snapshot-state<ul><li>按照云硬盘快照状态进行过滤</li><li>类型：String</li><li>必选：否</li><li>取值范围：<ul><li><code>NORMAL</code>：正常 </li><li><code>CREATING</code>：创建中 </li><li><code>ROLLBACKING</code>：回滚中 </li><li><code>COPYING_FROM_REMOTE</code>：跨地域复制中 </li><li><code>CHECKING_COPIED</code>：复制校验中</li><li><code>TORECYCLE</code>：待回收</li></ul></li></ul></li><li>disk-usage<ul><li>按照云硬盘使用用途进行过滤</li><li>类型：String</li><li>必选：否</li><li>取值范围：<ul><li><code>SYSTEM_DISK</code>：代表系统盘</li><li><code>DATA_DISK</code>：代表数据盘</li></ul></li></ul></li><li>project-id<ul><li>按云硬盘所属项目ID过滤</li><li>类型：String</li><li>必选：否</li></ul></li><li>disk-id<ul><li>按照云硬盘ID进行过滤，一次最多只能传入10个值</li><li>类型：String</li><li>必选：否</li></ul></li><li>encrypt<ul><li>按是否加密进行过滤</li><li>类型：String</li><li>必选：否</li></ul></li><li>snapshot-type<ul><li>按快照归属类型查询</li><li>类型：String</li><li>必选：否</li><li>取值范围：<ul><li><code>SHARED_SNAPSHOT</code>：表示共享过来的快照</li><li><code>PRIVATE_SNAPSHOT</code>：表示自己的私有快照</li></ul></li></ul></li></ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 快照列表排序的依据字段。取值范围：
	// <ul>
	// <li>CREATE_TIME：依据快照的创建时间排序</li>
	// <li>默认按创建时间排序</li>
	// </ul>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 输出云盘列表的排列顺序。取值范围：
	// <ul>
	//     <li>ASC：升序排列</li>
	//     <li>DESC：降序排列。</li>
	// </ul>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// 要查询快照的ID列表。参数不支持同时指定`SnapshotIds`和`Filters`。
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`

	// 过滤条件。参数不支持同时指定SnapshotIds和Filters。<br><ul><li>snapshot-id<ul><li>按照云硬盘快照ID进行过滤</li><li>类型：String</li><li>必选：否</li></ul></li><li>snapshot-name<ul><li>按照云硬盘快照名称进行过滤</li><li>类型：String</li><li>必选：否</li></ul></li><li>snapshot-state<ul><li>按照云硬盘快照状态进行过滤</li><li>类型：String</li><li>必选：否</li><li>取值范围：<ul><li><code>NORMAL</code>：正常 </li><li><code>CREATING</code>：创建中 </li><li><code>ROLLBACKING</code>：回滚中 </li><li><code>COPYING_FROM_REMOTE</code>：跨地域复制中 </li><li><code>CHECKING_COPIED</code>：复制校验中</li><li><code>TORECYCLE</code>：待回收</li></ul></li></ul></li><li>disk-usage<ul><li>按照云硬盘使用用途进行过滤</li><li>类型：String</li><li>必选：否</li><li>取值范围：<ul><li><code>SYSTEM_DISK</code>：代表系统盘</li><li><code>DATA_DISK</code>：代表数据盘</li></ul></li></ul></li><li>project-id<ul><li>按云硬盘所属项目ID过滤</li><li>类型：String</li><li>必选：否</li></ul></li><li>disk-id<ul><li>按照云硬盘ID进行过滤，一次最多只能传入10个值</li><li>类型：String</li><li>必选：否</li></ul></li><li>encrypt<ul><li>按是否加密进行过滤</li><li>类型：String</li><li>必选：否</li></ul></li><li>snapshot-type<ul><li>按快照归属类型查询</li><li>类型：String</li><li>必选：否</li><li>取值范围：<ul><li><code>SHARED_SNAPSHOT</code>：表示共享过来的快照</li><li><code>PRIVATE_SNAPSHOT</code>：表示自己的私有快照</li></ul></li></ul></li></ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 快照列表排序的依据字段。取值范围：
	// <ul>
	// <li>CREATE_TIME：依据快照的创建时间排序</li>
	// <li>默认按创建时间排序</li>
	// </ul>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 输出云盘列表的排列顺序。取值范围：
	// <ul>
	//     <li>ASC：升序排列</li>
	//     <li>DESC：降序排列。</li>
	// </ul>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
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
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Offset")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotsResponseParams struct {
	// 快照的数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

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
type DetachDisksRequestParams struct {
	// 将要卸载的云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询，单次请求最多可卸载10块弹性云盘。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 对于非共享型云盘，会根据该参数校验是否与实际挂载的实例一致；对于共享型云盘，该参数表示要从哪个CVM实例上卸载云盘。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DetachDisksRequest struct {
	*tchttp.BaseRequest
	
	// 将要卸载的云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询，单次请求最多可卸载10块弹性云盘。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 对于非共享型云盘，会根据该参数校验是否与实际挂载的实例一致；对于共享型云盘，该参数表示要从哪个CVM实例上卸载云盘。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	delete(f, "InstanceId")
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
	// 描述计费项目名称。
	PriceTitle *string `json:"PriceTitle,omitnil,omitempty" name:"PriceTitle"`

	// 描述计费项目显示名称，用户控制台展示。
	PriceName *string `json:"PriceName,omitnil,omitempty" name:"PriceName"`

	// 预付费云盘预支费用的原价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 预付费云盘预支费用的折扣价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// 后付费云盘原单价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// 后付费云盘折扣单价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`

	// 后付费云盘的计价单元，取值范围：HOUR：表示后付费云盘的计价单元是按小时计算。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeUnit *string `json:"ChargeUnit,omitnil,omitempty" name:"ChargeUnit"`

	// 高精度预付费云盘预支费用的原价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalPriceHigh *string `json:"OriginalPriceHigh,omitnil,omitempty" name:"OriginalPriceHigh"`

	// 高精度预付费云盘预支费用的折扣价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountPriceHigh *string `json:"DiscountPriceHigh,omitnil,omitempty" name:"DiscountPriceHigh"`

	// 高精度后付费云盘原单价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceHigh *string `json:"UnitPriceHigh,omitnil,omitempty" name:"UnitPriceHigh"`

	// 高精度后付费云盘折扣单价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceDiscountHigh *string `json:"UnitPriceDiscountHigh,omitnil,omitempty" name:"UnitPriceDiscountHigh"`
}

type Disk struct {
	// 云盘是否与挂载的实例一起销毁。<br><li>true:销毁实例时会同时销毁云盘，只支持按小时后付费云盘。</li><li>false：销毁实例时不销毁云盘。</li>
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitnil,omitempty" name:"DeleteWithInstance"`

	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费。</li>
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 硬盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬盘</li><li>CLOUD_PREMIUM：表示高性能云硬盘</li><li>CLOUD_BSSD：表示通用型SSD云硬盘</li><li>CLOUD_SSD：表示SSD云硬盘</li><li>CLOUD_HSSD：表示增强型SSD云硬盘</li><li>CLOUD_TSSD：表示极速型SSD云硬盘。</li>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 云盘状态。取值范围：<br><li>UNATTACHED：未挂载</li><li>ATTACHING：挂载中</li><li>ATTACHED：已挂载</li><li>DETACHING：解挂中</li><li>EXPANDING：扩容中</li><li>ROLLBACKING：回滚中</li><li>TORECYCLE：待回收</li><li>DUMPING：拷贝硬盘中。</li>
	DiskState *string `json:"DiskState,omitnil,omitempty" name:"DiskState"`

	// 云盘拥有的快照总数。
	SnapshotCount *int64 `json:"SnapshotCount,omitnil,omitempty" name:"SnapshotCount"`

	// 云盘已挂载到子机，且子机与云盘都是包年包月。<br><li>true：子机设置了自动续费标识，但云盘未设置</li><li>false：云盘自动续费标识正常。</li>
	AutoRenewFlagError *bool `json:"AutoRenewFlagError,omitnil,omitempty" name:"AutoRenewFlagError"`

	// 云盘是否处于快照回滚状态。取值范围：<br><li>false:表示不处于快照回滚状态</li><li>true:表示处于快照回滚状态。</li>
	Rollbacking *bool `json:"Rollbacking,omitnil,omitempty" name:"Rollbacking"`

	// 对于非共享型云盘，该参数为空数组。对于共享型云盘，则表示该云盘当前被挂载到的CVM实例InstanceId
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// 云盘是否为加密盘。取值范围：<br><li>false:表示非加密盘</li><li>true:表示加密盘。</li>
	Encrypt *bool `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`

	// 云硬盘名称。
	DiskName *string `json:"DiskName,omitnil,omitempty" name:"DiskName"`

	// 云硬盘因欠费销毁或者到期销毁时， 是否使用快照备份数据的标识。true表示销毁时创建快照进行数据备份。false表示直接销毁，不进行数据备份。
	BackupDisk *bool `json:"BackupDisk,omitnil,omitempty" name:"BackupDisk"`

	// 与云盘绑定的标签，云盘未绑定标签则取值为空。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 云硬盘挂载的云主机ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 云盘的挂载类型。取值范围：<br><li>PF: PF挂载</li><li>VF: VF挂载</li>
	AttachMode *string `json:"AttachMode,omitnil,omitempty" name:"AttachMode"`

	// 云盘关联的定期快照ID。只有在调用[DescribeDisks](/document/product/362/16315)接口时，入参ReturnBindAutoSnapshotPolicy取值为TRUE才会返回该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitnil,omitempty" name:"AutoSnapshotPolicyIds"`

	// 云硬盘额外性能值，单位MiB/s。
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitnil,omitempty" name:"ThroughputPerformance"`

	// 云盘是否处于类型变更中。取值范围：<br><li>false:表示云盘不处于类型变更中</li><li>true:表示云盘已发起类型变更，正处于迁移中。</li>
	Migrating *bool `json:"Migrating,omitnil,omitempty" name:"Migrating"`

	// 云硬盘ID。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 云盘拥有的快照总容量，单位为MiB。
	SnapshotSize *uint64 `json:"SnapshotSize,omitnil,omitempty" name:"SnapshotSize"`

	// 云硬盘所在的位置。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 判断预付费的云盘是否支持主动退还。<br><li>true:支持主动退还</li><li>false:不支持主动退还。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsReturnable *bool `json:"IsReturnable,omitnil,omitempty" name:"IsReturnable"`

	// 云硬盘的到期时间。
	DeadlineTime *string `json:"DeadlineTime,omitnil,omitempty" name:"DeadlineTime"`

	// 云盘是否挂载到云主机上。取值范围：<br><li>false:表示未挂载</li><li>true:表示已挂载。</li>
	Attached *bool `json:"Attached,omitnil,omitempty" name:"Attached"`

	// 云硬盘大小，单位GiB。
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 云盘类型变更的迁移进度，取值0到100。
	MigratePercent *uint64 `json:"MigratePercent,omitnil,omitempty" name:"MigratePercent"`

	// 云硬盘类型。取值范围：<br><li>SYSTEM_DISK：系统盘</li><li>DATA_DISK：数据盘。</li>
	DiskUsage *string `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// 付费模式。取值范围：<br><li>PREPAID：预付费，即包年包月</li><li>POSTPAID_BY_HOUR：后付费，即按量计费。</li>
	DiskChargeType *string `json:"DiskChargeType,omitnil,omitempty" name:"DiskChargeType"`

	// 是否为弹性云盘，false表示非弹性云盘，true表示弹性云盘。
	Portable *bool `json:"Portable,omitnil,omitempty" name:"Portable"`

	// 云盘是否具备创建快照的能力。取值范围：<br><li>false表示不具备</li><li>true表示具备。</li>
	SnapshotAbility *bool `json:"SnapshotAbility,omitnil,omitempty" name:"SnapshotAbility"`

	// 在云盘已挂载到实例，且实例与云盘都是包年包月的条件下，此字段才有意义。<br><li>true:云盘到期时间早于实例。</li><li>false：云盘到期时间晚于实例。</li>
	DeadlineError *bool `json:"DeadlineError,omitnil,omitempty" name:"DeadlineError"`

	// 云盘快照回滚的进度。
	RollbackPercent *uint64 `json:"RollbackPercent,omitnil,omitempty" name:"RollbackPercent"`

	// 当前时间距离云硬盘到期的天数（仅对预付费云硬盘有意义）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DifferDaysOfDeadline *int64 `json:"DifferDaysOfDeadline,omitnil,omitempty" name:"DifferDaysOfDeadline"`

	// 预付费云盘在不支持主动退还的情况下，该参数表明不支持主动退还的具体原因。取值范围：<br><li>1：云硬盘已经退还</li><li>2：云硬盘已过期</li><li>3：云盘不支持退还</li><li>8：超过可退还数量的限制。</li><li>10：非弹性云硬盘、系统盘、后付费云硬盘等不支持退还</li>
	ReturnFailCode *int64 `json:"ReturnFailCode,omitnil,omitempty" name:"ReturnFailCode"`

	// 云盘是否为共享型云盘。
	Shareable *bool `json:"Shareable,omitnil,omitempty" name:"Shareable"`

	// 云硬盘的创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 销毁云盘时删除关联的非永久保留快照。0 表示非永久快照不随云盘销毁而销毁，1表示非永久快照随云盘销毁而销毁，默认取0。快照是否永久保留可以通过[DescribeSnapshots](https://cloud.tencent.com/document/product/362/15647)接口返回的快照详情的IsPermanent字段来判断，true表示永久快照，false表示非永久快照。
	DeleteSnapshot *int64 `json:"DeleteSnapshot,omitnil,omitempty" name:"DeleteSnapshot"`

	// 云硬盘备份点配额。表示最大可以保留的备份点数量。
	DiskBackupQuota *uint64 `json:"DiskBackupQuota,omitnil,omitempty" name:"DiskBackupQuota"`

	// 云硬盘备份点已使用的数量。
	DiskBackupCount *uint64 `json:"DiskBackupCount,omitnil,omitempty" name:"DiskBackupCount"`

	// 云硬盘挂载实例的类型。取值范围：<br><li>CVM</li><li>EKS</li>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 云硬盘最后一次挂载的实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastAttachInsId *string `json:"LastAttachInsId,omitnil,omitempty" name:"LastAttachInsId"`

	// 云硬盘最后一次操作错误提示
	ErrorPrompt *string `json:"ErrorPrompt,omitnil,omitempty" name:"ErrorPrompt"`

	// 云盘是否开启性能突发
	BurstPerformance *bool `json:"BurstPerformance,omitnil,omitempty" name:"BurstPerformance"`

	// 云硬盘加密类型，值为ENCRYPT_V1和ENCRYPT_V2，分别表示第一代和第二代加密技术，两种加密技术互不兼容
	EncryptType *string `json:"EncryptType,omitnil,omitempty" name:"EncryptType"`

	// 加密盘密钥ID
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`
}

type DiskBackup struct {
	// 云硬盘备份点的ID。
	DiskBackupId *string `json:"DiskBackupId,omitnil,omitempty" name:"DiskBackupId"`

	// 云硬盘备份点关联的云硬盘ID。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 云硬盘大小，单位GiB。
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 云硬盘类型。取值范围：<br>
	// <li>SYSTEM_DISK：系统盘</li>
	// <li>DATA_DISK：数据盘。</li>
	DiskUsage *string `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// 备份点名称。
	DiskBackupName *string `json:"DiskBackupName,omitnil,omitempty" name:"DiskBackupName"`

	// <p>云硬盘备份点状态。取值范围：</p>
	// <ul>
	//   <li>NORMAL：正常</li>
	//   <li>CREATING：创建中</li>
	//   <li>ROLLBACKING：回滚中</li>
	// </ul>
	DiskBackupState *string `json:"DiskBackupState,omitnil,omitempty" name:"DiskBackupState"`

	// 云硬盘备份点创建百分比。
	Percent *uint64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 云硬盘备份点的创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 云盘是否为加密盘。取值范围：<br><li>false:表示非加密盘<br></li>true:表示加密盘。
	Encrypt *bool `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`
}

type DiskChargePrepaid struct {
	// 购买云硬盘的时长，默认单位为月，取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 自动续费标识。取值范围：
	// <ul>
	//   <li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li>
	//   <li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li>
	//   <li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li>
	// </ul>
	// 默认取值：NOTIFY_AND_MANUAL_RENEW。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 需要将云硬盘的到期时间与挂载的子机对齐时，可传入该参数。该参数表示子机当前的到期时间，此时Period如果传入，则表示子机需要续费的时长，云盘会自动按对齐到子机续费后的到期时间续费.
	CurInstanceDeadline *string `json:"CurInstanceDeadline,omitnil,omitempty" name:"CurInstanceDeadline"`
}

type DiskConfig struct {
	// 配置是否可用。
	Available *bool `json:"Available,omitnil,omitempty" name:"Available"`

	// 付费模式。取值范围：<br><li>PREPAID：表示预付费，即包年包月</li><br><li>POSTPAID_BY_HOUR：表示后付费，即按量计费。</li>
	DiskChargeType *string `json:"DiskChargeType,omitnil,omitempty" name:"DiskChargeType"`

	// 云硬盘所属的[可用区](/document/product/213/15753#ZoneInfo)。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例机型系列。详见[实例类型](https://cloud.tencent.com/document/product/213/11518)
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceFamily *string `json:"InstanceFamily,omitnil,omitempty" name:"InstanceFamily"`

	// 云盘介质类型。取值范围：<br>
	// CLOUD_BASIC：表示普通云硬盘<br>
	// CLOUD_PREMIUM：表示高性能云硬盘<br>
	// CLOUD_BSSD：表示通用型SSD云硬盘<br>
	// CLOUD_SSD：表示SSD云硬盘<br>
	// CLOUD_HSSD：表示增强型SSD云硬盘<br>
	// CLOUD_TSSD：表示极速型SSD云硬盘。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 云盘大小变化的最小步长，单位GiB。
	StepSize *uint64 `json:"StepSize,omitnil,omitempty" name:"StepSize"`

	// 额外的性能区间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraPerformanceRange []*int64 `json:"ExtraPerformanceRange,omitnil,omitempty" name:"ExtraPerformanceRange"`

	// 实例机型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceClass *string `json:"DeviceClass,omitnil,omitempty" name:"DeviceClass"`

	// 云盘类型。取值范围：<br><li>SYSTEM_DISK：表示系统盘</li><br><li>DATA_DISK：表示数据盘。</li>
	DiskUsage *string `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// 最小可配置云盘大小，单位GiB。
	MinDiskSize *uint64 `json:"MinDiskSize,omitnil,omitempty" name:"MinDiskSize"`

	// 最大可配置云盘大小，单位GiB。
	MaxDiskSize *uint64 `json:"MaxDiskSize,omitnil,omitempty" name:"MaxDiskSize"`

	// 描述预付费或后付费云盘的价格。
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`
}

type Filter struct {
	// 过滤键的名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 一个或者多个过滤值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type GetSnapOverviewRequestParams struct {

}

type GetSnapOverviewRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetSnapOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSnapOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSnapOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSnapOverviewResponseParams struct {
	// 用户快照总大小
	TotalSize *float64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`

	// 用户快照总大小（用于计费）
	RealTradeSize *float64 `json:"RealTradeSize,omitnil,omitempty" name:"RealTradeSize"`

	// 快照免费额度
	FreeQuota *float64 `json:"FreeQuota,omitnil,omitempty" name:"FreeQuota"`

	// 快照总个数
	TotalNums *int64 `json:"TotalNums,omitnil,omitempty" name:"TotalNums"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSnapOverviewResponse struct {
	*tchttp.BaseResponse
	Response *GetSnapOverviewResponseParams `json:"Response"`
}

func (r *GetSnapOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSnapOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Image struct {
	// 镜像名称。
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 镜像实例ID。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`
}

// Predefined struct for user
type InitializeDisksRequestParams struct {
	// 待重新初始化的云硬盘ID列表，可以通过[DescribeDisks](/document/product/362/16315)接口查询， 单次初始化限制20块以内
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`
}

type InitializeDisksRequest struct {
	*tchttp.BaseRequest
	
	// 待重新初始化的云硬盘ID列表，可以通过[DescribeDisks](/document/product/362/16315)接口查询， 单次初始化限制20块以内
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`
}

func (r *InitializeDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitializeDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InitializeDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InitializeDisksResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InitializeDisksResponse struct {
	*tchttp.BaseResponse
	Response *InitializeDisksResponseParams `json:"Response"`
}

func (r *InitializeDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitializeDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceModifyDiskBackupQuotaRequestParams struct {
	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 修改后的云硬盘备份点配额，即云盘可以拥有的备份点数量，单位为个。
	DiskBackupQuota *uint64 `json:"DiskBackupQuota,omitnil,omitempty" name:"DiskBackupQuota"`
}

type InquirePriceModifyDiskBackupQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 修改后的云硬盘备份点配额，即云盘可以拥有的备份点数量，单位为个。
	DiskBackupQuota *uint64 `json:"DiskBackupQuota,omitnil,omitempty" name:"DiskBackupQuota"`
}

func (r *InquirePriceModifyDiskBackupQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyDiskBackupQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	delete(f, "DiskBackupQuota")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceModifyDiskBackupQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceModifyDiskBackupQuotaResponseParams struct {
	// 描述了修改云硬盘备份点之后的云盘价格。
	DiskPrice *Price `json:"DiskPrice,omitnil,omitempty" name:"DiskPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceModifyDiskBackupQuotaResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceModifyDiskBackupQuotaResponseParams `json:"Response"`
}

func (r *InquirePriceModifyDiskBackupQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyDiskBackupQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceModifyDiskExtraPerformanceRequestParams struct {
	// 额外购买的云硬盘性能值，单位MiB/s。仅大小超过460GiB的增强型SSD（CLOUD_HSSD）和极速型SSD（CLOUD_TSSD）云硬盘才支持购买额外性能。
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitnil,omitempty" name:"ThroughputPerformance"`

	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`
}

type InquirePriceModifyDiskExtraPerformanceRequest struct {
	*tchttp.BaseRequest
	
	// 额外购买的云硬盘性能值，单位MiB/s。仅大小超过460GiB的增强型SSD（CLOUD_HSSD）和极速型SSD（CLOUD_TSSD）云硬盘才支持购买额外性能。
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitnil,omitempty" name:"ThroughputPerformance"`

	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`
}

func (r *InquirePriceModifyDiskExtraPerformanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyDiskExtraPerformanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ThroughputPerformance")
	delete(f, "DiskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceModifyDiskExtraPerformanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceModifyDiskExtraPerformanceResponseParams struct {
	// 描述了调整云盘额外性能时对应的价格。
	DiskPrice *Price `json:"DiskPrice,omitnil,omitempty" name:"DiskPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceModifyDiskExtraPerformanceResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceModifyDiskExtraPerformanceResponseParams `json:"Response"`
}

func (r *InquirePriceModifyDiskExtraPerformanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyDiskExtraPerformanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateDisksRequestParams struct {
	// 云硬盘计费类型： <ul>   <li>PREPAID：预付费，即包年包月</li>   <li>POSTPAID_BY_HOUR：按小时后付费</li> </ul>
	DiskChargeType *string `json:"DiskChargeType,omitnil,omitempty" name:"DiskChargeType"`

	// 硬盘介质类型。取值范围： <ul>   <li>CLOUD_PREMIUM：表示高性能云硬盘</li>   <li>CLOUD_SSD：表示SSD云硬盘</li>   <li>CLOUD_HSSD：表示增强型SSD云硬盘</li>   <li>CLOUD_TSSD：表示极速型SSD云硬盘</li> </ul>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 云硬盘大小，单位为GiB。云盘大小取值范围参见云硬盘[产品分类](/document/product/362/2353)的说明。
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 云硬盘所属项目ID。该参数可以通过调用[DescribeProject](https://cloud.tencent.com/document/api/651/78725) 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 购买云硬盘的数量。不填则默认为1。
	DiskCount *uint64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// 额外购买的云硬盘性能值，单位MiB/s。仅大小超过460GiB的增强型SSD（CLOUD_HSSD）和极速型SSD（CLOUD_TSSD）云硬盘才支持购买额外性能。
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitnil,omitempty" name:"ThroughputPerformance"`

	// 预付费模式，即包年包月相关参数设置。通过该参数指定包年包月云盘的购买时长、是否设置自动续费等属性。<br>创建预付费云盘该参数必传，创建按小时后付费云盘无需传该参数。
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitnil,omitempty" name:"DiskChargePrepaid"`

	// 指定云硬盘备份点配额。
	DiskBackupQuota *uint64 `json:"DiskBackupQuota,omitnil,omitempty" name:"DiskBackupQuota"`
}

type InquiryPriceCreateDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘计费类型： <ul>   <li>PREPAID：预付费，即包年包月</li>   <li>POSTPAID_BY_HOUR：按小时后付费</li> </ul>
	DiskChargeType *string `json:"DiskChargeType,omitnil,omitempty" name:"DiskChargeType"`

	// 硬盘介质类型。取值范围： <ul>   <li>CLOUD_PREMIUM：表示高性能云硬盘</li>   <li>CLOUD_SSD：表示SSD云硬盘</li>   <li>CLOUD_HSSD：表示增强型SSD云硬盘</li>   <li>CLOUD_TSSD：表示极速型SSD云硬盘</li> </ul>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 云硬盘大小，单位为GiB。云盘大小取值范围参见云硬盘[产品分类](/document/product/362/2353)的说明。
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 云硬盘所属项目ID。该参数可以通过调用[DescribeProject](https://cloud.tencent.com/document/api/651/78725) 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 购买云硬盘的数量。不填则默认为1。
	DiskCount *uint64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// 额外购买的云硬盘性能值，单位MiB/s。仅大小超过460GiB的增强型SSD（CLOUD_HSSD）和极速型SSD（CLOUD_TSSD）云硬盘才支持购买额外性能。
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitnil,omitempty" name:"ThroughputPerformance"`

	// 预付费模式，即包年包月相关参数设置。通过该参数指定包年包月云盘的购买时长、是否设置自动续费等属性。<br>创建预付费云盘该参数必传，创建按小时后付费云盘无需传该参数。
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitnil,omitempty" name:"DiskChargePrepaid"`

	// 指定云硬盘备份点配额。
	DiskBackupQuota *uint64 `json:"DiskBackupQuota,omitnil,omitempty" name:"DiskBackupQuota"`
}

func (r *InquiryPriceCreateDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskChargeType")
	delete(f, "DiskType")
	delete(f, "DiskSize")
	delete(f, "ProjectId")
	delete(f, "DiskCount")
	delete(f, "ThroughputPerformance")
	delete(f, "DiskChargePrepaid")
	delete(f, "DiskBackupQuota")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateDisksResponseParams struct {
	// 描述了新购云盘的价格。
	DiskPrice *Price `json:"DiskPrice,omitnil,omitempty" name:"DiskPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceCreateDisksResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceCreateDisksResponseParams `json:"Response"`
}

func (r *InquiryPriceCreateDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewDisksRequestParams struct {
	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月云盘的购买时长。如果在该参数中指定CurInstanceDeadline，则会按对齐到子机到期时间来续费。如果是批量续费询价，该参数与Disks参数一一对应，元素数量需保持一致。
	DiskChargePrepaids []*DiskChargePrepaid `json:"DiskChargePrepaids,omitnil,omitempty" name:"DiskChargePrepaids"`

	// 指定云硬盘新的到期时间，形式如：2017-12-17 00:00:00。参数`NewDeadline`和`DiskChargePrepaids`是两种指定询价时长的方式，两者必传一个。
	NewDeadline *string `json:"NewDeadline,omitnil,omitempty" name:"NewDeadline"`

	// 云硬盘所属项目ID。该参数可以通过调用[DescribeProject](https://cloud.tencent.com/document/api/651/78725) 的返回值中的 projectId 字段来获取。 如传入则仅用于鉴权。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type InquiryPriceRenewDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月云盘的购买时长。如果在该参数中指定CurInstanceDeadline，则会按对齐到子机到期时间来续费。如果是批量续费询价，该参数与Disks参数一一对应，元素数量需保持一致。
	DiskChargePrepaids []*DiskChargePrepaid `json:"DiskChargePrepaids,omitnil,omitempty" name:"DiskChargePrepaids"`

	// 指定云硬盘新的到期时间，形式如：2017-12-17 00:00:00。参数`NewDeadline`和`DiskChargePrepaids`是两种指定询价时长的方式，两者必传一个。
	NewDeadline *string `json:"NewDeadline,omitnil,omitempty" name:"NewDeadline"`

	// 云硬盘所属项目ID。该参数可以通过调用[DescribeProject](https://cloud.tencent.com/document/api/651/78725) 的返回值中的 projectId 字段来获取。 如传入则仅用于鉴权。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *InquiryPriceRenewDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "DiskChargePrepaids")
	delete(f, "NewDeadline")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceRenewDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewDisksResponseParams struct {
	// 描述了续费云盘的价格。
	DiskPrice *PrepayPrice `json:"DiskPrice,omitnil,omitempty" name:"DiskPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceRenewDisksResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceRenewDisksResponseParams `json:"Response"`
}

func (r *InquiryPriceRenewDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResizeDiskRequestParams struct {
	// 云硬盘扩容后的大小，单位为GiB，不得小于当前云硬盘大小。云盘大小取值范围参见云硬盘[产品分类](/document/product/362/2353)的说明。
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 云硬盘所属项目ID。该参数可以通过调用[DescribeProject](https://cloud.tencent.com/document/api/651/78725) 的返回值中的 projectId 字段来获取。 如传入则仅用于鉴权。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type InquiryPriceResizeDiskRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘扩容后的大小，单位为GiB，不得小于当前云硬盘大小。云盘大小取值范围参见云硬盘[产品分类](/document/product/362/2353)的说明。
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 云硬盘所属项目ID。该参数可以通过调用[DescribeProject](https://cloud.tencent.com/document/api/651/78725) 的返回值中的 projectId 字段来获取。 如传入则仅用于鉴权。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *InquiryPriceResizeDiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResizeDiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskSize")
	delete(f, "DiskId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceResizeDiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResizeDiskResponseParams struct {
	// 描述了扩容云盘的价格。
	DiskPrice *PrepayPrice `json:"DiskPrice,omitnil,omitempty" name:"DiskPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceResizeDiskResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceResizeDiskResponseParams `json:"Response"`
}

func (r *InquiryPriceResizeDiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResizeDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoSnapshotPolicyAttributeRequestParams struct {
	// 定期快照策略ID。可以通过[查询定期快照策略](https://cloud.tencent.com/document/product/362/33556)API查询。
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 是否激活定期快照策略，`false`表示未激活，`true`表示激活；默认为`true`。
	IsActivated *bool `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// 通过该定期快照策略创建的快照是否永久保留。`false`表示非永久保留，`true`表示永久保留，默认为`false`。
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 要创建的定期快照策略名。不传则默认为“未命名”。最大长度不能超60个字节。
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitnil,omitempty" name:"AutoSnapshotPolicyName"`

	// 定期快照的执行策略。
	Policy []*Policy `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 通过该定期快照策略创建的快照保留天数。如果指定本参数，则IsPermanent入参不可指定为TRUE，否则会产生冲突。
	RetentionDays *uint64 `json:"RetentionDays,omitnil,omitempty" name:"RetentionDays"`
}

type ModifyAutoSnapshotPolicyAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 定期快照策略ID。可以通过[查询定期快照策略](https://cloud.tencent.com/document/product/362/33556)API查询。
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 是否激活定期快照策略，`false`表示未激活，`true`表示激活；默认为`true`。
	IsActivated *bool `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// 通过该定期快照策略创建的快照是否永久保留。`false`表示非永久保留，`true`表示永久保留，默认为`false`。
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 要创建的定期快照策略名。不传则默认为“未命名”。最大长度不能超60个字节。
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitnil,omitempty" name:"AutoSnapshotPolicyName"`

	// 定期快照的执行策略。
	Policy []*Policy `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 通过该定期快照策略创建的快照保留天数。如果指定本参数，则IsPermanent入参不可指定为TRUE，否则会产生冲突。
	RetentionDays *uint64 `json:"RetentionDays,omitnil,omitempty" name:"RetentionDays"`
}

func (r *ModifyAutoSnapshotPolicyAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoSnapshotPolicyAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoSnapshotPolicyId")
	delete(f, "IsActivated")
	delete(f, "IsPermanent")
	delete(f, "AutoSnapshotPolicyName")
	delete(f, "Policy")
	delete(f, "RetentionDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoSnapshotPolicyAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoSnapshotPolicyAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAutoSnapshotPolicyAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoSnapshotPolicyAttributeResponseParams `json:"Response"`
}

func (r *ModifyAutoSnapshotPolicyAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoSnapshotPolicyAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiskAttributesRequestParams struct {
	// 一个或多个待操作的云硬盘ID，可以通过[DescribeDisks](/document/product/362/16315)接口查询。如果传入多个云盘ID，仅支持将所有云盘修改为同一属性。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 新的云硬盘名称。
	DiskName *string `json:"DiskName,omitnil,omitempty" name:"DiskName"`

	// 是否为弹性云盘，FALSE表示非弹性云盘，TRUE表示弹性云盘。仅支持非弹性云盘修改为弹性云盘。
	Portable *bool `json:"Portable,omitnil,omitempty" name:"Portable"`

	// 新的云硬盘项目ID，只支持修改弹性云盘的项目ID。通过[DescribeProject](/document/api/378/4400)接口查询可用项目及其ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 成功挂载到云主机后该云硬盘是否随云主机销毁，TRUE表示随云主机销毁，FALSE表示不随云主机销毁。仅支持按量计费云硬盘数据盘。
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitnil,omitempty" name:"DeleteWithInstance"`

	// 变更云盘类型时，可传入该参数，表示变更的目标类型，取值范围：<br><li>CLOUD_PREMIUM：表示高性能云硬盘</li><li>CLOUD_SSD：表示SSD云硬盘。</li>当前不支持批量变更类型，即传入DiskType时，DiskIds仅支持传入一块云盘；<br>变更云盘类型时不支持同时变更其他属性。
	// 具体说明请参考[调整云硬盘类型](https://cloud.tencent.com/document/product/362/32540)
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 开启/关闭云盘性能突发功能，取值范围： 
	// CREATE：开启
	// CANCEL：关闭
	BurstPerformanceOperation *string `json:"BurstPerformanceOperation,omitnil,omitempty" name:"BurstPerformanceOperation"`
}

type ModifyDiskAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的云硬盘ID，可以通过[DescribeDisks](/document/product/362/16315)接口查询。如果传入多个云盘ID，仅支持将所有云盘修改为同一属性。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 新的云硬盘名称。
	DiskName *string `json:"DiskName,omitnil,omitempty" name:"DiskName"`

	// 是否为弹性云盘，FALSE表示非弹性云盘，TRUE表示弹性云盘。仅支持非弹性云盘修改为弹性云盘。
	Portable *bool `json:"Portable,omitnil,omitempty" name:"Portable"`

	// 新的云硬盘项目ID，只支持修改弹性云盘的项目ID。通过[DescribeProject](/document/api/378/4400)接口查询可用项目及其ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 成功挂载到云主机后该云硬盘是否随云主机销毁，TRUE表示随云主机销毁，FALSE表示不随云主机销毁。仅支持按量计费云硬盘数据盘。
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitnil,omitempty" name:"DeleteWithInstance"`

	// 变更云盘类型时，可传入该参数，表示变更的目标类型，取值范围：<br><li>CLOUD_PREMIUM：表示高性能云硬盘</li><li>CLOUD_SSD：表示SSD云硬盘。</li>当前不支持批量变更类型，即传入DiskType时，DiskIds仅支持传入一块云盘；<br>变更云盘类型时不支持同时变更其他属性。
	// 具体说明请参考[调整云硬盘类型](https://cloud.tencent.com/document/product/362/32540)
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 开启/关闭云盘性能突发功能，取值范围： 
	// CREATE：开启
	// CANCEL：关闭
	BurstPerformanceOperation *string `json:"BurstPerformanceOperation,omitnil,omitempty" name:"BurstPerformanceOperation"`
}

func (r *ModifyDiskAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "DiskName")
	delete(f, "Portable")
	delete(f, "ProjectId")
	delete(f, "DeleteWithInstance")
	delete(f, "DiskType")
	delete(f, "BurstPerformanceOperation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDiskAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiskAttributesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDiskAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDiskAttributesResponseParams `json:"Response"`
}

func (r *ModifyDiskAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiskBackupQuotaRequestParams struct {
	// 云硬盘ID。可通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 调整之后的云硬盘备份点配额。取值范围为1 ~ 1024。
	DiskBackupQuota *uint64 `json:"DiskBackupQuota,omitnil,omitempty" name:"DiskBackupQuota"`
}

type ModifyDiskBackupQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘ID。可通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 调整之后的云硬盘备份点配额。取值范围为1 ~ 1024。
	DiskBackupQuota *uint64 `json:"DiskBackupQuota,omitnil,omitempty" name:"DiskBackupQuota"`
}

func (r *ModifyDiskBackupQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskBackupQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	delete(f, "DiskBackupQuota")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDiskBackupQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiskBackupQuotaResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDiskBackupQuotaResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDiskBackupQuotaResponseParams `json:"Response"`
}

func (r *ModifyDiskBackupQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskBackupQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiskExtraPerformanceRequestParams struct {
	// 额外购买的云硬盘性能值，单位MiB/s。
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitnil,omitempty" name:"ThroughputPerformance"`

	// 需要购买额外性能值的云硬盘ID，可通过[DescribeDisks](/document/product/362/16315)接口查询。仅大小超过460GiB的增强型SSD（CLOUD_HSSD）和极速型SSD（CLOUD_TSSD）云硬盘才支持购买额外性能。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`
}

type ModifyDiskExtraPerformanceRequest struct {
	*tchttp.BaseRequest
	
	// 额外购买的云硬盘性能值，单位MiB/s。
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitnil,omitempty" name:"ThroughputPerformance"`

	// 需要购买额外性能值的云硬盘ID，可通过[DescribeDisks](/document/product/362/16315)接口查询。仅大小超过460GiB的增强型SSD（CLOUD_HSSD）和极速型SSD（CLOUD_TSSD）云硬盘才支持购买额外性能。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`
}

func (r *ModifyDiskExtraPerformanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskExtraPerformanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ThroughputPerformance")
	delete(f, "DiskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDiskExtraPerformanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiskExtraPerformanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDiskExtraPerformanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDiskExtraPerformanceResponseParams `json:"Response"`
}

func (r *ModifyDiskExtraPerformanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskExtraPerformanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDisksChargeTypeRequestParams struct {
	// 一个或多个待操作的云硬盘ID,可以通过[DescribeDisks](/document/product/362/16315)接口查询。每次请求批量云硬盘上限为100。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 设置为预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitnil,omitempty" name:"DiskChargePrepaid"`

	// 设置为后付费模式
	DiskChargePostpaid *bool `json:"DiskChargePostpaid,omitnil,omitempty" name:"DiskChargePostpaid"`
}

type ModifyDisksChargeTypeRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的云硬盘ID,可以通过[DescribeDisks](/document/product/362/16315)接口查询。每次请求批量云硬盘上限为100。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 设置为预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitnil,omitempty" name:"DiskChargePrepaid"`

	// 设置为后付费模式
	DiskChargePostpaid *bool `json:"DiskChargePostpaid,omitnil,omitempty" name:"DiskChargePostpaid"`
}

func (r *ModifyDisksChargeTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDisksChargeTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "DiskChargePrepaid")
	delete(f, "DiskChargePostpaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDisksChargeTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDisksChargeTypeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDisksChargeTypeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDisksChargeTypeResponseParams `json:"Response"`
}

func (r *ModifyDisksChargeTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDisksChargeTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDisksRenewFlagRequestParams struct {
	// 一个或多个待操作的云硬盘ID，该参数可以通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 云硬盘的自动续费标识。取值范围：<ul><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li></ul>
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 该参数支持设置云硬盘的自动续费周期，单位为月。取值范围：[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36]
	AutoRenewPeriod *uint64 `json:"AutoRenewPeriod,omitnil,omitempty" name:"AutoRenewPeriod"`
}

type ModifyDisksRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的云硬盘ID，该参数可以通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 云硬盘的自动续费标识。取值范围：<ul><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li></ul>
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 该参数支持设置云硬盘的自动续费周期，单位为月。取值范围：[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36]
	AutoRenewPeriod *uint64 `json:"AutoRenewPeriod,omitnil,omitempty" name:"AutoRenewPeriod"`
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
	delete(f, "AutoRenewPeriod")
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
type ModifySnapshotAttributeRequestParams struct {
	// 快照ID, 可通过[DescribeSnapshots](https://cloud.tencent.com/document/api/362/15647)查询。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 快照的保留方式，FALSE表示非永久保留，TRUE表示永久保留。
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 新的快照名称。最长为60个字符。
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 快照的到期时间；设置好快照将会被同时设置为非永久保留方式；超过到期时间后快照将会被自动删除。注：该参数仅在参数IsPermanent为False时生效。
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`
}

type ModifySnapshotAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 快照ID, 可通过[DescribeSnapshots](https://cloud.tencent.com/document/api/362/15647)查询。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 快照的保留方式，FALSE表示非永久保留，TRUE表示永久保留。
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 新的快照名称。最长为60个字符。
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 快照的到期时间；设置好快照将会被同时设置为非永久保留方式；超过到期时间后快照将会被自动删除。注：该参数仅在参数IsPermanent为False时生效。
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`
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
	delete(f, "IsPermanent")
	delete(f, "SnapshotName")
	delete(f, "Deadline")
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

// Predefined struct for user
type ModifySnapshotsSharePermissionRequestParams struct {
	// 快照ID, 可通过[DescribeSnapshots](https://cloud.tencent.com/document/api/362/15647)查询获取。
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`

	// 接收分享快照的账号Id列表，array型参数的格式可以参考[API简介](https://cloud.tencent.com/document/api/213/568)。账号ID不同于QQ号，查询用户账号ID请查看[账号信息](https://console.cloud.tencent.com/developer)中的账号ID栏。
	AccountIds []*string `json:"AccountIds,omitnil,omitempty" name:"AccountIds"`

	// 操作，包括 SHARE，CANCEL。其中SHARE代表分享操作，CANCEL代表取消分享操作。
	Permission *string `json:"Permission,omitnil,omitempty" name:"Permission"`
}

type ModifySnapshotsSharePermissionRequest struct {
	*tchttp.BaseRequest
	
	// 快照ID, 可通过[DescribeSnapshots](https://cloud.tencent.com/document/api/362/15647)查询获取。
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`

	// 接收分享快照的账号Id列表，array型参数的格式可以参考[API简介](https://cloud.tencent.com/document/api/213/568)。账号ID不同于QQ号，查询用户账号ID请查看[账号信息](https://console.cloud.tencent.com/developer)中的账号ID栏。
	AccountIds []*string `json:"AccountIds,omitnil,omitempty" name:"AccountIds"`

	// 操作，包括 SHARE，CANCEL。其中SHARE代表分享操作，CANCEL代表取消分享操作。
	Permission *string `json:"Permission,omitnil,omitempty" name:"Permission"`
}

func (r *ModifySnapshotsSharePermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotsSharePermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotIds")
	delete(f, "AccountIds")
	delete(f, "Permission")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySnapshotsSharePermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotsSharePermissionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySnapshotsSharePermissionResponse struct {
	*tchttp.BaseResponse
	Response *ModifySnapshotsSharePermissionResponseParams `json:"Response"`
}

func (r *ModifySnapshotsSharePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotsSharePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Placement struct {
	// 云硬盘所属的[可用区](/document/product/213/15753#ZoneInfo)。该参数也可以通过调用  [DescribeZones](/document/product/213/15707) 的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 围笼Id，可通过 [DescribeDiskStoragePool](https://cloud.tencent.com/document/api/362/62143) 获取。作为入参时，表示对指定的CageId的资源进行操作，可为空。 作为出参时，表示资源所属围笼ID，可为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// 实例所属项目ID，可通过[DescribeProject](/document/api/651/78725)获取。不填默认为0，表示默认项目。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例所属项目名称，可通过[DescribeProject](/document/api/651/78725)获取。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 独享集群名字。作为入参时，忽略。作为出参时，表示云硬盘所属的独享集群名，可为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdcName *string `json:"CdcName,omitnil,omitempty" name:"CdcName"`

	// 实例所属的独享集群ID。可通过 [DescribeDiskStoragePool](https://cloud.tencent.com/document/api/362/62143) 获取。作为入参时，表示对指定的CdcId独享集群的资源进行操作，可为空。 作为出参时，表示资源所属的独享集群的ID，可为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// 独享集群id。
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

type Policy struct {
	// 指定定期快照策略的触发时间。单位为小时，取值范围：[0, 23]。00:00 ~ 23:00 共 24 个时间点可选，1表示 01:00，依此类推。
	Hour []*uint64 `json:"Hour,omitnil,omitempty" name:"Hour"`

	// 指定每周从周一到周日需要触发定期快照的日期，取值范围：[0, 6]。0表示周日触发，1-6分别表示周一至周六。
	DayOfWeek []*uint64 `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 指定每月从月初到月底需要触发定期快照的日期,取值范围：[1, 31]，1-31分别表示每月的具体日期，比如5表示每月的5号。注：若设置29、30、31等部分月份不存在的日期，则对应不存在日期的月份会跳过不打定期快照。
	DayOfMonth []*uint64 `json:"DayOfMonth,omitnil,omitempty" name:"DayOfMonth"`

	// 指定创建定期快照的间隔天数，取值范围：[1, 365]，例如设置为5，则间隔5天即触发定期快照创建。注：当选择按天备份时，理论上第一次备份的时间为备份策略创建当天。如果当天备份策略创建的时间已经晚于设置的备份时间，那么将会等到第二个备份周期再进行第一次备份。
	IntervalDays *uint64 `json:"IntervalDays,omitnil,omitempty" name:"IntervalDays"`
}

type PrepayPrice struct {
	// 预付费云盘或快照预支费用的折扣价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// 后付费云盘的计价单元，取值范围：<br><li>HOUR：表示后付费云盘的计价单元是按小时计算。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeUnit *string `json:"ChargeUnit,omitnil,omitempty" name:"ChargeUnit"`

	// 高精度后付费云盘原单价, 单位：元
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceHigh *string `json:"UnitPriceHigh,omitnil,omitempty" name:"UnitPriceHigh"`

	// 高精度预付费云盘或快照预支费用的原价，单位：元
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalPriceHigh *string `json:"OriginalPriceHigh,omitnil,omitempty" name:"OriginalPriceHigh"`

	// 预付费云盘或快照预支费用的原价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 后付费云盘折扣单价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`

	// 高精度后付费云盘折扣单价, 单位：元
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceDiscountHigh *string `json:"UnitPriceDiscountHigh,omitnil,omitempty" name:"UnitPriceDiscountHigh"`

	// 高精度预付费云盘或快照预支费用的折扣价，单位：元
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountPriceHigh *string `json:"DiscountPriceHigh,omitnil,omitempty" name:"DiscountPriceHigh"`

	// 后付费云盘原单价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// 计费项目明细列表。
	DetailPrices []*DetailPrice `json:"DetailPrices,omitnil,omitempty" name:"DetailPrices"`
}

type Price struct {
	// 后付费云盘折扣单价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`

	// 预付费云盘预支费用的折扣价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// 后付费云盘原单价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// 高精度后付费云盘原单价, 单位：元
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceHigh *string `json:"UnitPriceHigh,omitnil,omitempty" name:"UnitPriceHigh"`

	// 高精度预付费云盘预支费用的原价, 单位：元	。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalPriceHigh *string `json:"OriginalPriceHigh,omitnil,omitempty" name:"OriginalPriceHigh"`

	// 预付费云盘预支费用的原价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 高精度预付费云盘预支费用的折扣价, 单位：元
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountPriceHigh *string `json:"DiscountPriceHigh,omitnil,omitempty" name:"DiscountPriceHigh"`

	// 高精度后付费云盘折扣单价, 单位：元
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceDiscountHigh *string `json:"UnitPriceDiscountHigh,omitnil,omitempty" name:"UnitPriceDiscountHigh"`

	// 后付费云盘的计价单元，取值范围：<br><li>HOUR：表示后付费云盘的计价单元是按小时计算。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeUnit *string `json:"ChargeUnit,omitnil,omitempty" name:"ChargeUnit"`
}

// Predefined struct for user
type RenewDiskRequestParams struct {
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月云硬盘的续费时长。<br>在云硬盘与挂载的实例一起续费的场景下，可以指定参数CurInstanceDeadline，此时云硬盘会按对齐到实例续费后的到期时间来续费。
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitnil,omitempty" name:"DiskChargePrepaid"`

	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`
}

type RenewDiskRequest struct {
	*tchttp.BaseRequest
	
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月云硬盘的续费时长。<br>在云硬盘与挂载的实例一起续费的场景下，可以指定参数CurInstanceDeadline，此时云硬盘会按对齐到实例续费后的到期时间来续费。
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitnil,omitempty" name:"DiskChargePrepaid"`

	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`
}

func (r *RenewDiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskChargePrepaid")
	delete(f, "DiskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewDiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDiskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewDiskResponse struct {
	*tchttp.BaseResponse
	Response *RenewDiskResponseParams `json:"Response"`
}

func (r *RenewDiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeDiskRequestParams struct {
	// 云硬盘扩容后的大小，单位为GB，必须大于当前云硬盘大小。云盘大小取值范围参见云硬盘[产品分类](/document/product/362/2353)的说明。
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。该字段仅供单块云硬盘扩容时传入。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`
}

type ResizeDiskRequest struct {
	*tchttp.BaseRequest
	
	// 云硬盘扩容后的大小，单位为GB，必须大于当前云硬盘大小。云盘大小取值范围参见云硬盘[产品分类](/document/product/362/2353)的说明。
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。该字段仅供单块云硬盘扩容时传入。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`
}

func (r *ResizeDiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeDiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskSize")
	delete(f, "DiskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResizeDiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeDiskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResizeDiskResponse struct {
	*tchttp.BaseResponse
	Response *ResizeDiskResponseParams `json:"Response"`
}

func (r *ResizeDiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SharePermission struct {
	// 快照分享的时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 分享的账号Id
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`
}

type Snapshot struct {
	// 快照所在的位置。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 是否为跨地域复制的快照。取值范围：
	// <ul>
	//     <li>true：表示为跨地域复制的快照。</li>
	//     <li>false：本地域的快照。</li>
	// </ul>
	CopyFromRemote *bool `json:"CopyFromRemote,omitnil,omitempty" name:"CopyFromRemote"`

	// 快照的状态。取值范围：
	// <ul>
	//     <li>NORMAL：正常</li>
	//     <li>CREATING：创建中</li>
	//     <li>ROLLBACKING：回滚中</li>
	//     <li>COPYING_FROM_REMOTE：跨地域复制中</li>
	//     <li>CHECKING_COPIED：复制校验中</li>
	//     <li>TORECYCLE：待回收</li>
	// </ul>
	SnapshotState *string `json:"SnapshotState,omitnil,omitempty" name:"SnapshotState"`

	// 是否为永久快照。取值范围：
	// <ul>
	//     <li>true：永久快照</li>
	//     <li>false：非永久快照</li>
	// </ul>
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 快照名称，用户自定义的快照别名。调用[ModifySnapshotAttribute](/document/product/362/15650)可修改此字段。
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 快照到期时间。如果快照为永久保留，此字段为空。
	DeadlineTime *string `json:"DeadlineTime,omitnil,omitempty" name:"DeadlineTime"`

	// 快照创建进度百分比，快照创建成功后此字段恒为100。
	Percent *uint64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 快照关联的镜像列表。
	Images []*Image `json:"Images,omitnil,omitempty" name:"Images"`

	// 快照当前被共享数。
	ShareReference *uint64 `json:"ShareReference,omitnil,omitempty" name:"ShareReference"`

	// 快照类型，目前该项取值可以为`PRIVATE_SNAPSHOT`（私有快照）或者`SHARED_SNAPSHOT`（共享快照）
	SnapshotType *string `json:"SnapshotType,omitnil,omitempty" name:"SnapshotType"`

	// 创建此快照的云硬盘大小，单位GiB。
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 创建此快照的云硬盘ID。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 快照正在跨地域复制的目的地域，若没有则返回`[]`。
	CopyingToRegions []*string `json:"CopyingToRegions,omitnil,omitempty" name:"CopyingToRegions"`

	// 是否为加密盘创建的快照。取值范围：
	// <ul>
	//     <li>true：该快照为加密盘创建的</li>
	//     <li>false：非加密盘创建的快照</li>
	// </ul>
	Encrypt *bool `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`

	// 快照的创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 快照关联的镜像个数。
	ImageCount *uint64 `json:"ImageCount,omitnil,omitempty" name:"ImageCount"`

	// 创建此快照的云硬盘类型。取值范围：
	// <ul>
	//     <li>SYSTEM_DISK：系统盘</li>
	//     <li>DATA_DISK：数据盘</li>
	// </ul>
	DiskUsage *string `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// 快照ID。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 快照开始共享的时间。
	TimeStartShare *string `json:"TimeStartShare,omitnil,omitempty" name:"TimeStartShare"`

	// 快照绑定的标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type SnapshotCopyResult struct {
	// 复制到目标地域的新快照ID。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 指示具体错误信息，成功时为空字符串。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 错误码，成功时取值为“Success”。
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 跨地复制的目标地域。
	DestinationRegion *string `json:"DestinationRegion,omitnil,omitempty" name:"DestinationRegion"`
}

type SnapshotGroup struct {
	// 快照组ID。
	SnapshotGroupId *string `json:"SnapshotGroupId,omitnil,omitempty" name:"SnapshotGroupId"`

	// 快照组类型。NORMAL: 普通快照组，非一致性快照。
	SnapshotGroupType *string `json:"SnapshotGroupType,omitnil,omitempty" name:"SnapshotGroupType"`

	// 快照组是否包含系统盘快照。
	ContainRootSnapshot *bool `json:"ContainRootSnapshot,omitnil,omitempty" name:"ContainRootSnapshot"`

	// 快照组包含的快照ID列表。
	SnapshotIdSet []*string `json:"SnapshotIdSet,omitnil,omitempty" name:"SnapshotIdSet"`

	// <ul>
	//     <li>NORMAL: 正常</li>
	//     <li>CREATING: 创建中</li>
	//     <li>ROLLBACKING: 回滚中</li>
	// </ul>
	SnapshotGroupState *string `json:"SnapshotGroupState,omitnil,omitempty" name:"SnapshotGroupState"`

	// 快照组创建进度。
	Percent *uint64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 快照组创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 快照组最新修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 快照组关联的镜像列表。
	Images []*Image `json:"Images,omitnil,omitempty" name:"Images"`

	// 快照组名称。
	SnapshotGroupName *string `json:"SnapshotGroupName,omitnil,omitempty" name:"SnapshotGroupName"`

	// 快照组关联的镜像数量。
	ImageCount *uint64 `json:"ImageCount,omitnil,omitempty" name:"ImageCount"`

	// 快照组是否永久保留
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 快照组到期时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadlineTime *string `json:"DeadlineTime,omitnil,omitempty" name:"DeadlineTime"`

	// 来源自动快照策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`
}

type Tag struct {
	// 标签健。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type TerminateDisksRequestParams struct {
	// 需退还的云盘ID列表，通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 销毁云盘时删除关联的非永久保留快照。0 表示非永久快照不随云盘销毁而销毁，1表示非永久快照随云盘销毁而销毁，默认取0。快照是否永久保留可以通过DescribeSnapshots接口返回的快照详情的IsPermanent字段来判断，true表示永久快照，false表示非永久快照。
	DeleteSnapshot *int64 `json:"DeleteSnapshot,omitnil,omitempty" name:"DeleteSnapshot"`
}

type TerminateDisksRequest struct {
	*tchttp.BaseRequest
	
	// 需退还的云盘ID列表，通过[DescribeDisks](/document/product/362/16315)接口查询。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 销毁云盘时删除关联的非永久保留快照。0 表示非永久快照不随云盘销毁而销毁，1表示非永久快照随云盘销毁而销毁，默认取0。快照是否永久保留可以通过DescribeSnapshots接口返回的快照详情的IsPermanent字段来判断，true表示永久快照，false表示非永久快照。
	DeleteSnapshot *int64 `json:"DeleteSnapshot,omitnil,omitempty" name:"DeleteSnapshot"`
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
	delete(f, "DeleteSnapshot")
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
type UnbindAutoSnapshotPolicyRequestParams struct {
	// 要解绑的定期快照策略ID。
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 要解绑定期快照策略的云盘ID列表。此参数与 InstanceIds 参数至少需要传入一个。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 要解绑定期快照策略的实例ID列表。此参数与 DiskIds 参数至少需要传入一个。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type UnbindAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 要解绑的定期快照策略ID。
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 要解绑定期快照策略的云盘ID列表。此参数与 InstanceIds 参数至少需要传入一个。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 要解绑定期快照策略的实例ID列表。此参数与 DiskIds 参数至少需要传入一个。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *UnbindAutoSnapshotPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoSnapshotPolicyId")
	delete(f, "DiskIds")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindAutoSnapshotPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *UnbindAutoSnapshotPolicyResponseParams `json:"Response"`
}

func (r *UnbindAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}