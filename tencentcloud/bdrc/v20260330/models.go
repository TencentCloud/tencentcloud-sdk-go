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

package v20260330

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AdvancedRetentionPolicy struct {
	// 保留设定天数中的每天最新的一个备份
	Days *uint64 `json:"Days,omitnil,omitempty" name:"Days"`

	// 保留设置周中的每周最新的一个备份
	Weeks *uint64 `json:"Weeks,omitnil,omitempty" name:"Weeks"`

	// 保留设置月内的每月最新的一个备份
	Months *uint64 `json:"Months,omitnil,omitempty" name:"Months"`

	// 保留设置年内的每年最新的一个备份
	Years *uint64 `json:"Years,omitnil,omitempty" name:"Years"`
}

// Predefined struct for user
type ApplyBackupGroupRequestParams struct {
	// 回滚的备份组ID。
	BackupGroupId *string `json:"BackupGroupId,omitnil,omitempty" name:"BackupGroupId"`

	// 回滚的备份ID、云硬盘ID列表。
	ApplyDisks []*ApplyDisk `json:"ApplyDisks,omitnil,omitempty" name:"ApplyDisks"`

	// 回滚备份前是否执行自动关机，如果回滚的盘挂载在实例上且实例处于运行状态，可传入该参数。
	AutoStopInstance *bool `json:"AutoStopInstance,omitnil,omitempty" name:"AutoStopInstance"`

	// 回滚备份完成后是否执行自动开机。
	AutoStartInstance *bool `json:"AutoStartInstance,omitnil,omitempty" name:"AutoStartInstance"`
}

type ApplyBackupGroupRequest struct {
	*tchttp.BaseRequest
	
	// 回滚的备份组ID。
	BackupGroupId *string `json:"BackupGroupId,omitnil,omitempty" name:"BackupGroupId"`

	// 回滚的备份ID、云硬盘ID列表。
	ApplyDisks []*ApplyDisk `json:"ApplyDisks,omitnil,omitempty" name:"ApplyDisks"`

	// 回滚备份前是否执行自动关机，如果回滚的盘挂载在实例上且实例处于运行状态，可传入该参数。
	AutoStopInstance *bool `json:"AutoStopInstance,omitnil,omitempty" name:"AutoStopInstance"`

	// 回滚备份完成后是否执行自动开机。
	AutoStartInstance *bool `json:"AutoStartInstance,omitnil,omitempty" name:"AutoStartInstance"`
}

func (r *ApplyBackupGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyBackupGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupGroupId")
	delete(f, "ApplyDisks")
	delete(f, "AutoStopInstance")
	delete(f, "AutoStartInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyBackupGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyBackupGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyBackupGroupResponse struct {
	*tchttp.BaseResponse
	Response *ApplyBackupGroupResponseParams `json:"Response"`
}

func (r *ApplyBackupGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyBackupGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyDisk struct {
	// 备份ID
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 云盘ID
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`
}

type AspInfo struct {
	// 备份策略ID
	AspId *string `json:"AspId,omitnil,omitempty" name:"AspId"`

	// 备份策略名称
	AspName *string `json:"AspName,omitnil,omitempty" name:"AspName"`

	// 备份策略状态
	AspState *string `json:"AspState,omitnil,omitempty" name:"AspState"`

	// 备份策略执行详情
	Policy []*Policy `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 备份策略是否使能
	IsActivated *bool `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// 是否永久保留
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 保留时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetentionDays *uint64 `json:"RetentionDays,omitnil,omitempty" name:"RetentionDays"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type AutoBackupPolicy struct {
	// 定期备份策略是否激活。
	IsActivated *bool `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// 使用该定期备份策略创建出来的备份是否永久保留。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 使用该定期备份策略创建出来的备份是否永久保留。
	NextTriggerTime *string `json:"NextTriggerTime,omitnil,omitempty" name:"NextTriggerTime"`

	// NORMAL
	AutoBackupPolicyState *string `json:"AutoBackupPolicyState,omitnil,omitempty" name:"AutoBackupPolicyState"`

	// 备份策略的名称。
	AutoBackupPolicyName *string `json:"AutoBackupPolicyName,omitnil,omitempty" name:"AutoBackupPolicyName"`

	// 定期备份的执行策略。
	Policy []*Policy `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 备份策略ID。
	AutoBackupPolicyId *string `json:"AutoBackupPolicyId,omitnil,omitempty" name:"AutoBackupPolicyId"`

	// 备份策略的创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 使用该定期备份策略创建出来的备份保留天数。
	RetentionDays *uint64 `json:"RetentionDays,omitnil,omitempty" name:"RetentionDays"`

	// 用户AppId。
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 定期备份策略绑定的实例ID列表。
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`

	// 该定期快照创建的快照最大保留月数
	RetentionMonths *uint64 `json:"RetentionMonths,omitnil,omitempty" name:"RetentionMonths"`

	// 该定期快照创建的快照最大保留数量
	RetentionAmount *uint64 `json:"RetentionAmount,omitnil,omitempty" name:"RetentionAmount"`

	// 创建人。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 主账号uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountUin *string `json:"AccountUin,omitnil,omitempty" name:"AccountUin"`

	// 子账号uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitnil,omitempty" name:"SubAccountUin"`

	// 策略存储类型
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 备份库ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// 高级保留策略
	AdvancedRetentionPolicy *AdvancedRetentionPolicy `json:"AdvancedRetentionPolicy,omitnil,omitempty" name:"AdvancedRetentionPolicy"`
}

type AutomationServiceEnabled struct {
	// 是否开启该服务。取值范围：TRUE（开启）/FALSE（不开启）。默认取值：TRUE。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type BackupDeniedAction struct {
	// 备份ID。
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 具体的备份操作掩码列表。
	DeniedActions []*DeniedAction `json:"DeniedActions,omitnil,omitempty" name:"DeniedActions"`
}

type BackupDetail struct {
	// 备份组ID
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 备份策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoBackupPolicyId *string `json:"AutoBackupPolicyId,omitnil,omitempty" name:"AutoBackupPolicyId"`

	// 备份和云盘绑定关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupBindDisk []*ApplyDisk `json:"BackupBindDisk,omitnil,omitempty" name:"BackupBindDisk"`
}

type BackupGroup struct {
	// 备份组ID。
	BackupGroupId *string `json:"BackupGroupId,omitnil,omitempty" name:"BackupGroupId"`

	// 备份组创建进度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent *uint64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 备份和云盘绑定关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupBindDisk []*ApplyDisk `json:"BackupBindDisk,omitnil,omitempty" name:"BackupBindDisk"`

	// 备份组名称。
	BackupGroupName *string `json:"BackupGroupName,omitnil,omitempty" name:"BackupGroupName"`

	// 备份组状态。NORMAL: 正常；CREATING: 创建中；ROLLBACKING: 回滚中
	BackupGroupState *string `json:"BackupGroupState,omitnil,omitempty" name:"BackupGroupState"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 用户AppId。
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 是否为永久备份组。
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 备份组的到期时间。如果为永久备份组，则取值为null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadlineTime *string `json:"DeadlineTime,omitnil,omitempty" name:"DeadlineTime"`

	// 创建备份组的实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 创建备份组时刻实例的详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceDetails *string `json:"InstanceDetails,omitnil,omitempty" name:"InstanceDetails"`

	// 创建人名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 主账号uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountUin *string `json:"AccountUin,omitnil,omitempty" name:"AccountUin"`

	// 创建备份的子账号uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitnil,omitempty" name:"SubAccountUin"`

	// 创建当前备份的定期备份策略ID，为null则为手动创建的备份。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoBackupPolicyId *string `json:"AutoBackupPolicyId,omitnil,omitempty" name:"AutoBackupPolicyId"`
}

type BackupGroupDeniedAction struct {
	// 备份组ID
	BackupGroupId *string `json:"BackupGroupId,omitnil,omitempty" name:"BackupGroupId"`

	// 拒绝的操作
	DeniedActions []*DeniedAction `json:"DeniedActions,omitnil,omitempty" name:"DeniedActions"`
}

type BackupGroupRollbackTask struct {
	// 备份组回滚任务
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 源实例ID
	SourceInstanceId *string `json:"SourceInstanceId,omitnil,omitempty" name:"SourceInstanceId"`

	// 目标实例ID
	TargetInstanceId *string `json:"TargetInstanceId,omitnil,omitempty" name:"TargetInstanceId"`

	// 备份组ID
	BackupGroupId *string `json:"BackupGroupId,omitnil,omitempty" name:"BackupGroupId"`

	// 恢复类型：ORIGINAL-原实例恢复，NEW-新实例恢复
	RollbackType *string `json:"RollbackType,omitnil,omitempty" name:"RollbackType"`

	// 任务状态。取值包括"init"、"migrating"、"done"、"failed"。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent *uint64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// APP ID
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 备份点名称
	BackupGroupName *string `json:"BackupGroupName,omitnil,omitempty" name:"BackupGroupName"`

	// 恢复失败原因
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`
}

type BackupInfo struct {
	// 备份点ID
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 备份名称
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// 所属计划ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 策略ID
	AspInstanceId *string `json:"AspInstanceId,omitnil,omitempty" name:"AspInstanceId"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 备份状态，取值如下：
	// 0 备份完成
	// 1 创建中（备份进行中）
	// 2 部分成功（指定的备份路径中部分目录不存在）
	// 3 恢复中（该备份点正在被恢复任务使用）
	// 92  已取消
	// 98 创建失败
	// 99 已删除
	// 100 删除中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 备份路径
	BackupPaths []*string `json:"BackupPaths,omitnil,omitempty" name:"BackupPaths"`

	// 包含文件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncludeFileTypes []*string `json:"IncludeFileTypes,omitnil,omitempty" name:"IncludeFileTypes"`

	// 排除路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludePatterns []*string `json:"ExcludePatterns,omitnil,omitempty" name:"ExcludePatterns"`

	// 是否排除系统目录
	ExcludeSystemDirectories *bool `json:"ExcludeSystemDirectories,omitnil,omitempty" name:"ExcludeSystemDirectories"`

	// 备份库ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// 扫描文件数
	ScannedFileCount *int64 `json:"ScannedFileCount,omitnil,omitempty" name:"ScannedFileCount"`

	// 扫描大小(字节)
	ScannedSize *int64 `json:"ScannedSize,omitnil,omitempty" name:"ScannedSize"`

	// 扫描大小(格式化)
	ScannedSizeFormatted *string `json:"ScannedSizeFormatted,omitnil,omitempty" name:"ScannedSizeFormatted"`

	// 已备份文件数量
	BackupFileCount *int64 `json:"BackupFileCount,omitnil,omitempty" name:"BackupFileCount"`

	// 已备份大小(字节)
	BackupSize *int64 `json:"BackupSize,omitnil,omitempty" name:"BackupSize"`

	// 已备份大小(格式化)
	BackupSizeFormatted *string `json:"BackupSizeFormatted,omitnil,omitempty" name:"BackupSizeFormatted"`

	// 备份进度(0-100)
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 是否为永久保留
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 到期时间
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 不存在的路径信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NonExistSourcePaths []*string `json:"NonExistSourcePaths,omitnil,omitempty" name:"NonExistSourcePaths"`

	// 备份失败原因
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`

	// 备份所属AppId
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 备份类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

type BackupInstance struct {
	// 实例绑定的定期备份策略列表。
	AutoBackupPolicyIdSet []*string `json:"AutoBackupPolicyIdSet,omitnil,omitempty" name:"AutoBackupPolicyIdSet"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户AppId。
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 实例最新备份时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestBackupTime *string `json:"LatestBackupTime,omitnil,omitempty" name:"LatestBackupTime"`

	// 实例的备份组ID列表。
	BackupGroupIdSet []*string `json:"BackupGroupIdSet,omitnil,omitempty" name:"BackupGroupIdSet"`

	// 修改时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type BackupPlan struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备份策略ID
	AutoBackupPolicyId *string `json:"AutoBackupPolicyId,omitnil,omitempty" name:"AutoBackupPolicyId"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// APP ID
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 备份数量
	BackupCount *uint64 `json:"BackupCount,omitnil,omitempty" name:"BackupCount"`

	// 上次执行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTriggerTime *string `json:"LastTriggerTime,omitnil,omitempty" name:"LastTriggerTime"`

	// 上次执行错误信息，如果为空表示上次执行成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTriggerError *string `json:"LastTriggerError,omitnil,omitempty" name:"LastTriggerError"`
}

type BackupPolicyOverview struct {
	// 自动备份策略总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 已绑定资源的策略数
	BoundCount *int64 `json:"BoundCount,omitnil,omitempty" name:"BoundCount"`

	// 未绑定任何资源的策略数
	UnboundCount *int64 `json:"UnboundCount,omitnil,omitempty" name:"UnboundCount"`
}

type BackupVault struct {
	// 备份库ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// 备份库名称
	VaultName *string `json:"VaultName,omitnil,omitempty" name:"VaultName"`

	// 备份库描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 备份库状态：READ_WRITE / READ_ONLY / UNAVAILABLE / DELETING
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 加密方式：NONE / SSE-COS / SSE-KMS
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptType *string `json:"EncryptType,omitnil,omitempty" name:"EncryptType"`

	// KMS密钥ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`

	// 备份库类型：COMMON
	VaultType *string `json:"VaultType,omitnil,omitempty" name:"VaultType"`

	// 关联的备份策略按类型统计
	BackupPolicySet []*TypeCount `json:"BackupPolicySet,omitnil,omitempty" name:"BackupPolicySet"`

	// 备份点按类型统计（不含已删除）
	BackupSet []*TypeCount `json:"BackupSet,omitnil,omitempty" name:"BackupSet"`

	// 地域信息
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 源端数据量
	SourceDataSize *uint64 `json:"SourceDataSize,omitnil,omitempty" name:"SourceDataSize"`

	// 存储库数据量
	VaultDataSize *uint64 `json:"VaultDataSize,omitnil,omitempty" name:"VaultDataSize"`
}

type BackupVaultOverview struct {
	// 备份库总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 备份库总存储量（已用容量），单位 MB
	TotalSizeMb *int64 `json:"TotalSizeMb,omitnil,omitempty" name:"TotalSizeMb"`
}

type BasicServicesSettings struct {
	// 是否开启基础服务。取值范围：TRUE（开启）/FALSE（不开启）。默认取值：TRUE。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

// Predefined struct for user
type BindAutoBackupPolicyRequestParams struct {

}

type BindAutoBackupPolicyRequest struct {
	*tchttp.BaseRequest
	
}

func (r *BindAutoBackupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindAutoBackupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindAutoBackupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindAutoBackupPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindAutoBackupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *BindAutoBackupPolicyResponseParams `json:"Response"`
}

func (r *BindAutoBackupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindAutoBackupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommonBackupPoint struct {
	// 共同时间点（精确到小时）
	BackupCommonTime *string `json:"BackupCommonTime,omitnil,omitempty" name:"BackupCommonTime"`

	// 共同备份点信息
	BackupDetailSet []*BackupDetail `json:"BackupDetailSet,omitnil,omitempty" name:"BackupDetailSet"`
}

type CopyPair struct {
	// 用户ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 复制对ID（CVM 类型为 cvmcopypair-xxxxxxxx，DISK/CFS 类型为 copypair-xxxxxxxx）
	CopyPairId *string `json:"CopyPairId,omitnil,omitempty" name:"CopyPairId"`

	// 复制对名称
	CopyPairName *string `json:"CopyPairName,omitnil,omitempty" name:"CopyPairName"`

	// 所属容灾站点对ID
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 所属容灾站点对名称
	SitePairName *string `json:"SitePairName,omitnil,omitempty" name:"SitePairName"`

	// 保护组ID
	ProtectGroupId *string `json:"ProtectGroupId,omitnil,omitempty" name:"ProtectGroupId"`

	// 保护组名称
	ProtectGroupName *string `json:"ProtectGroupName,omitnil,omitempty" name:"ProtectGroupName"`

	// 复制对状态。可选值：INIT、RUNNING、FULL_COPYING、INC_COPYING、NORMAL、DOWN、DEGRADE 等
	CopyPairState *string `json:"CopyPairState,omitnil,omitempty" name:"CopyPairState"`

	// 复制对类型。可选值：DISK、INSTANCE、CFS
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`

	// 生产地域
	SourceRegion *string `json:"SourceRegion,omitnil,omitempty" name:"SourceRegion"`

	// 生产可用区
	SourceZone *string `json:"SourceZone,omitnil,omitempty" name:"SourceZone"`

	// 生产端VPC
	SourceVpc *string `json:"SourceVpc,omitnil,omitempty" name:"SourceVpc"`

	// 容灾地域
	TargetRegion *string `json:"TargetRegion,omitnil,omitempty" name:"TargetRegion"`

	// 容灾可用区
	TargetZone *string `json:"TargetZone,omitnil,omitempty" name:"TargetZone"`

	// 容灾端VPC
	TargetVpc *string `json:"TargetVpc,omitnil,omitempty" name:"TargetVpc"`

	// 生产资源ID。CVM 类型为源 InstanceId（ins-xxx）；DISK 类型为源 DiskId（disk-xxx）；CFS 类型为源 FilesystemId（cfs-xxx）
	SourceResourceId *string `json:"SourceResourceId,omitnil,omitempty" name:"SourceResourceId"`

	// 容灾资源ID。语义同 SourceResourceId（CVM/DISK/CFS）。延迟创建模式且 CVM 未真实创建时为占位符 drp-xxx，CVM 创建后为真实 ins-xxx
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetResourceId *string `json:"TargetResourceId,omitnil,omitempty" name:"TargetResourceId"`

	// 生产站点盘挂载的实例ID（DISK 类型时为挂载的 CVM ins-xxx；INSTANCE 类型时与 SourceResourceId 一致）
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 所属CVM复制对ID（仅 DISK 类型且其 CVM 复制对存在时返回）
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCopyPairId *string `json:"InstanceCopyPairId,omitnil,omitempty" name:"InstanceCopyPairId"`

	// 复制进度。CVM 类型为所有挂载磁盘进度的平均值；DISK/CFS 类型为本盘进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent *int64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 最新保护时间点。当 CopyPairState=FULL_COPYING 时为 null（首次全量未完成）
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestProtectionTime *string `json:"LatestProtectionTime,omitnil,omitempty" name:"LatestProtectionTime"`

	// RPO（秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecoveryPointObjective *int64 `json:"RecoveryPointObjective,omitnil,omitempty" name:"RecoveryPointObjective"`

	// 数据方向。可选值：POSITIVE（正向）、REVERSE（反向，failover 后）。后端在 REVERSE 时已自动轮转 src/target 字段
	DataDirection *string `json:"DataDirection,omitnil,omitempty" name:"DataDirection"`

	// 创建来源。可选值：LOCAL（本地侧创建）、PEER（对端创建）
	CreateFrom *string `json:"CreateFrom,omitnil,omitempty" name:"CreateFrom"`

	// 容灾类型。可选值：CROSS_ZONE（跨可用区）、CROSS_REGION（跨地域）、CROSS_CLOUD（跨云）
	DisasterRecoveryType *string `json:"DisasterRecoveryType,omitnil,omitempty" name:"DisasterRecoveryType"`

	// 对端云名称（仅跨云场景）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeerCloudName *string `json:"PeerCloudName,omitnil,omitempty" name:"PeerCloudName"`

	// 是否在回滚中（0/1）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rollbacking *int64 `json:"Rollbacking,omitnil,omitempty" name:"Rollbacking"`

	// 回滚进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	RollbackPercent *int64 `json:"RollbackPercent,omitnil,omitempty" name:"RollbackPercent"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 创建账户 Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountUin *string `json:"AccountUin,omitnil,omitempty" name:"AccountUin"`

	// 创建协作者 Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitnil,omitempty" name:"SubAccountUin"`

	// 演练组ID（用于演练组内过滤存量复制对，无演练时为 null）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DrillGroupId *string `json:"DrillGroupId,omitnil,omitempty" name:"DrillGroupId"`

	// 保护时间点列表（仅当 QueryProtectionTime=true 时返回）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectionTimeSet []*string `json:"ProtectionTimeSet,omitnil,omitempty" name:"ProtectionTimeSet"`

	// CVM下挂载磁盘的复制对列表（仅 CopyPairType=INSTANCE 时返回）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskCopyPairSet []*DiskCopyPairForCvm `json:"DiskCopyPairSet,omitnil,omitempty" name:"DiskCopyPairSet"`

	// 是否为延迟创建模式（创建后固定不变）。仅 CVM 复制对返回
	DeferredCreate *bool `json:"DeferredCreate,omitnil,omitempty" name:"DeferredCreate"`

	// 目标 CVM 是否已真实创建（首次 failover 完成后置 true）。仅 CVM 复制对返回
	TargetCvmCreated *bool `json:"TargetCvmCreated,omitnil,omitempty" name:"TargetCvmCreated"`

	// CVM 创建参数（JSON 字符串）。仅当请求传 QueryCvmCreateParams=true 且复制对处于 deferred_create=1 AND target_cvm_created=0 时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	CvmCreateParams *string `json:"CvmCreateParams,omitnil,omitempty" name:"CvmCreateParams"`
}

type CopyPairDeniedAction struct {
	// 复制对ID
	CopyPairId *string `json:"CopyPairId,omitnil,omitempty" name:"CopyPairId"`

	// 被禁止的操作列表（Action名称数组）
	DeniedActions []*DeniedAction `json:"DeniedActions,omitnil,omitempty" name:"DeniedActions"`
}

type CopyPairPrice struct {
	// 后付费每小时原价，单位：元
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// 高精度后付费每小时原价，单位：元（字符串形式，避免精度丢失）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceHigh *string `json:"UnitPriceHigh,omitnil,omitempty" name:"UnitPriceHigh"`

	// 后付费每小时折扣价，单位：元
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`

	// 高精度后付费每小时折扣价，单位：元（字符串形式，避免精度丢失）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceDiscountHigh *string `json:"UnitPriceDiscountHigh,omitnil,omitempty" name:"UnitPriceDiscountHigh"`

	// 折扣，100 表示无折扣，80 表示 8 折
	// 注意：此字段可能返回 null，表示取不到有效值。
	Discount *int64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 计价单元，固定为 HOUR（按小时计费）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeUnit *string `json:"ChargeUnit,omitnil,omitempty" name:"ChargeUnit"`

	// 计费项目明细列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetailPrices []*CopyPairPriceDetail `json:"DetailPrices,omitnil,omitempty" name:"DetailPrices"`
}

type CopyPairPriceDetail struct {
	// 计费项目标识名称。取值：InstanceCount（容灾CVM实例数）、InstanceDataCapacity（容灾CVM实例数据量）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriceName *string `json:"PriceName,omitnil,omitempty" name:"PriceName"`

	// 计费项目展示名称（跟随语言环境翻译）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriceTitle *string `json:"PriceTitle,omitnil,omitempty" name:"PriceTitle"`

	// 该计费项每小时原价，单位：元
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// 该计费项每小时折扣价，单位：元
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`

	// 该计费项的折扣，100 表示无折扣
	// 注意：此字段可能返回 null，表示取不到有效值。
	Discount *int64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 计价单元，固定为 HOUR
	ChargeUnit *string `json:"ChargeUnit,omitnil,omitempty" name:"ChargeUnit"`
}

// Predefined struct for user
type CreateAutoBackupPolicyRequestParams struct {
	// 定期备份的执行策略。
	Policy []*Policy `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 通过该定期备份策略创建的备份是否永久保留。false表示非永久保留，true表示永久保留，默认为false。
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 定期备份策略的名称。
	AutoBackupPolicyName *string `json:"AutoBackupPolicyName,omitnil,omitempty" name:"AutoBackupPolicyName"`

	// 是否激活定期备份策略。
	IsActivated *bool `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// 通过定期备份策略创建出的备份保留时间。
	RetentionDays *uint64 `json:"RetentionDays,omitnil,omitempty" name:"RetentionDays"`

	// 该定期备份策略创建的备份可以保留的月数，该参数不可与IsPermanent/RetentionDays参数冲突。
	RetentionMonths *uint64 `json:"RetentionMonths,omitnil,omitempty" name:"RetentionMonths"`

	// 通过该定期备份策略最多保留的备份个数，超过该个数限制后自动删除最先创建的备份，该参数不可与IsPermanent参数冲突。
	RetentionAmount *uint64 `json:"RetentionAmount,omitnil,omitempty" name:"RetentionAmount"`

	// 备份存储类型。COMMON表示走普通模式（不需要备份库），VAULT表示走备份库（必须关联一个备份库）。默认为COMMON
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 备份库ID，创建agent备份策略时必须指定。当StorageType为VAULT时必传。
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// 定期备份高级保留策略，该参数不可与IsPermanent参数冲突。
	AdvancedRetentionPolicy *AdvancedRetentionPolicy `json:"AdvancedRetentionPolicy,omitnil,omitempty" name:"AdvancedRetentionPolicy"`
}

type CreateAutoBackupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 定期备份的执行策略。
	Policy []*Policy `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 通过该定期备份策略创建的备份是否永久保留。false表示非永久保留，true表示永久保留，默认为false。
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 定期备份策略的名称。
	AutoBackupPolicyName *string `json:"AutoBackupPolicyName,omitnil,omitempty" name:"AutoBackupPolicyName"`

	// 是否激活定期备份策略。
	IsActivated *bool `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// 通过定期备份策略创建出的备份保留时间。
	RetentionDays *uint64 `json:"RetentionDays,omitnil,omitempty" name:"RetentionDays"`

	// 该定期备份策略创建的备份可以保留的月数，该参数不可与IsPermanent/RetentionDays参数冲突。
	RetentionMonths *uint64 `json:"RetentionMonths,omitnil,omitempty" name:"RetentionMonths"`

	// 通过该定期备份策略最多保留的备份个数，超过该个数限制后自动删除最先创建的备份，该参数不可与IsPermanent参数冲突。
	RetentionAmount *uint64 `json:"RetentionAmount,omitnil,omitempty" name:"RetentionAmount"`

	// 备份存储类型。COMMON表示走普通模式（不需要备份库），VAULT表示走备份库（必须关联一个备份库）。默认为COMMON
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 备份库ID，创建agent备份策略时必须指定。当StorageType为VAULT时必传。
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// 定期备份高级保留策略，该参数不可与IsPermanent参数冲突。
	AdvancedRetentionPolicy *AdvancedRetentionPolicy `json:"AdvancedRetentionPolicy,omitnil,omitempty" name:"AdvancedRetentionPolicy"`
}

func (r *CreateAutoBackupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoBackupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Policy")
	delete(f, "IsPermanent")
	delete(f, "AutoBackupPolicyName")
	delete(f, "IsActivated")
	delete(f, "RetentionDays")
	delete(f, "RetentionMonths")
	delete(f, "RetentionAmount")
	delete(f, "StorageType")
	delete(f, "VaultId")
	delete(f, "AdvancedRetentionPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAutoBackupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoBackupPolicyResponseParams struct {
	// 定期备份策略ID。
	AutoBackupPolicyId *string `json:"AutoBackupPolicyId,omitnil,omitempty" name:"AutoBackupPolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAutoBackupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateAutoBackupPolicyResponseParams `json:"Response"`
}

func (r *CreateAutoBackupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoBackupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupGroupRequestParams struct {
	// 需要创建备份组的云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 备份组的名称
	BackupGroupName *string `json:"BackupGroupName,omitnil,omitempty" name:"BackupGroupName"`

	// 指定备份组到期时间，如果未传入该参数，默认为永久保留。
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`
}

type CreateBackupGroupRequest struct {
	*tchttp.BaseRequest
	
	// 需要创建备份组的云硬盘ID列表。
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 备份组的名称
	BackupGroupName *string `json:"BackupGroupName,omitnil,omitempty" name:"BackupGroupName"`

	// 指定备份组到期时间，如果未传入该参数，默认为永久保留。
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`
}

func (r *CreateBackupGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "BackupGroupName")
	delete(f, "Deadline")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupGroupResponseParams struct {
	// 备份组ID。
	BackupGroupId *string `json:"BackupGroupId,omitnil,omitempty" name:"BackupGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBackupGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupGroupResponseParams `json:"Response"`
}

func (r *CreateBackupGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupVaultRequestParams struct {
	// 备份库名称
	VaultName *string `json:"VaultName,omitnil,omitempty" name:"VaultName"`

	// 备份库描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 加密方式: NONE/SSE-COS/SSE-KMS
	EncryptType *string `json:"EncryptType,omitnil,omitempty" name:"EncryptType"`

	// KMS密钥ID（SSE-KMS时使用）
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`
}

type CreateBackupVaultRequest struct {
	*tchttp.BaseRequest
	
	// 备份库名称
	VaultName *string `json:"VaultName,omitnil,omitempty" name:"VaultName"`

	// 备份库描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 加密方式: NONE/SSE-COS/SSE-KMS
	EncryptType *string `json:"EncryptType,omitnil,omitempty" name:"EncryptType"`

	// KMS密钥ID（SSE-KMS时使用）
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`
}

func (r *CreateBackupVaultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupVaultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VaultName")
	delete(f, "Description")
	delete(f, "EncryptType")
	delete(f, "KmsKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupVaultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupVaultResponseParams struct {
	// 备份库唯一ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBackupVaultResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupVaultResponseParams `json:"Response"`
}

func (r *CreateBackupVaultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupVaultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDisasterRecoveryProtectGroupRequestParams struct {
	// 所属容灾站点对id
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 容灾保护组的产品类型
	ProtectGroupType *string `json:"ProtectGroupType,omitnil,omitempty" name:"ProtectGroupType"`

	// 容灾保护组预期rpo, 单位分钟（当前仅支持15分钟）
	RecoveryPointObjective *int64 `json:"RecoveryPointObjective,omitnil,omitempty" name:"RecoveryPointObjective"`

	// 容灾保护组的名称，最大长度不能超60个字符。
	ProtectGroupName *string `json:"ProtectGroupName,omitnil,omitempty" name:"ProtectGroupName"`

	// 数据复制方向， ['POSITIVE', 'REVERSE']
	DataDirection *string `json:"DataDirection,omitnil,omitempty" name:"DataDirection"`
}

type CreateDisasterRecoveryProtectGroupRequest struct {
	*tchttp.BaseRequest
	
	// 所属容灾站点对id
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 容灾保护组的产品类型
	ProtectGroupType *string `json:"ProtectGroupType,omitnil,omitempty" name:"ProtectGroupType"`

	// 容灾保护组预期rpo, 单位分钟（当前仅支持15分钟）
	RecoveryPointObjective *int64 `json:"RecoveryPointObjective,omitnil,omitempty" name:"RecoveryPointObjective"`

	// 容灾保护组的名称，最大长度不能超60个字符。
	ProtectGroupName *string `json:"ProtectGroupName,omitnil,omitempty" name:"ProtectGroupName"`

	// 数据复制方向， ['POSITIVE', 'REVERSE']
	DataDirection *string `json:"DataDirection,omitnil,omitempty" name:"DataDirection"`
}

func (r *CreateDisasterRecoveryProtectGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisasterRecoveryProtectGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SitePairId")
	delete(f, "ProtectGroupType")
	delete(f, "RecoveryPointObjective")
	delete(f, "ProtectGroupName")
	delete(f, "DataDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDisasterRecoveryProtectGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDisasterRecoveryProtectGroupResponseParams struct {
	// 创建的容灾保护组ID
	ProtectGroupId *string `json:"ProtectGroupId,omitnil,omitempty" name:"ProtectGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDisasterRecoveryProtectGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateDisasterRecoveryProtectGroupResponseParams `json:"Response"`
}

func (r *CreateDisasterRecoveryProtectGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisasterRecoveryProtectGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDisasterRecoverySitePairRequestParams struct {
	// 容灾策略的容灾类型，跨地域：CROSS_REGION，或跨可用区：CROSS_ZONE
	DisasterRecoveryType *string `json:"DisasterRecoveryType,omitnil,omitempty" name:"DisasterRecoveryType"`

	// 生产站点地域
	SourceRegion *string `json:"SourceRegion,omitnil,omitempty" name:"SourceRegion"`

	// 容灾策略生产站点可用区
	SourceZone *string `json:"SourceZone,omitnil,omitempty" name:"SourceZone"`

	// 容灾站点地域
	TargetRegion *string `json:"TargetRegion,omitnil,omitempty" name:"TargetRegion"`

	// 容灾策略容灾站点可用区
	TargetZone *string `json:"TargetZone,omitnil,omitempty" name:"TargetZone"`

	// 容灾策略生产vpc
	SourceVpc *string `json:"SourceVpc,omitnil,omitempty" name:"SourceVpc"`

	// 容灾策略容灾vpc
	TargetVpc *string `json:"TargetVpc,omitnil,omitempty" name:"TargetVpc"`

	// 容灾策略所属产品类型，包括DISK、CFS、INSTANCE
	SitePairProductType *string `json:"SitePairProductType,omitnil,omitempty" name:"SitePairProductType"`

	// 容灾策略的名称，最大长度为60个字符。
	SitePairName *string `json:"SitePairName,omitnil,omitempty" name:"SitePairName"`

	// 容灾策略复制技术SYN/ASY
	CopyType *string `json:"CopyType,omitnil,omitempty" name:"CopyType"`
}

type CreateDisasterRecoverySitePairRequest struct {
	*tchttp.BaseRequest
	
	// 容灾策略的容灾类型，跨地域：CROSS_REGION，或跨可用区：CROSS_ZONE
	DisasterRecoveryType *string `json:"DisasterRecoveryType,omitnil,omitempty" name:"DisasterRecoveryType"`

	// 生产站点地域
	SourceRegion *string `json:"SourceRegion,omitnil,omitempty" name:"SourceRegion"`

	// 容灾策略生产站点可用区
	SourceZone *string `json:"SourceZone,omitnil,omitempty" name:"SourceZone"`

	// 容灾站点地域
	TargetRegion *string `json:"TargetRegion,omitnil,omitempty" name:"TargetRegion"`

	// 容灾策略容灾站点可用区
	TargetZone *string `json:"TargetZone,omitnil,omitempty" name:"TargetZone"`

	// 容灾策略生产vpc
	SourceVpc *string `json:"SourceVpc,omitnil,omitempty" name:"SourceVpc"`

	// 容灾策略容灾vpc
	TargetVpc *string `json:"TargetVpc,omitnil,omitempty" name:"TargetVpc"`

	// 容灾策略所属产品类型，包括DISK、CFS、INSTANCE
	SitePairProductType *string `json:"SitePairProductType,omitnil,omitempty" name:"SitePairProductType"`

	// 容灾策略的名称，最大长度为60个字符。
	SitePairName *string `json:"SitePairName,omitnil,omitempty" name:"SitePairName"`

	// 容灾策略复制技术SYN/ASY
	CopyType *string `json:"CopyType,omitnil,omitempty" name:"CopyType"`
}

func (r *CreateDisasterRecoverySitePairRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisasterRecoverySitePairRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisasterRecoveryType")
	delete(f, "SourceRegion")
	delete(f, "SourceZone")
	delete(f, "TargetRegion")
	delete(f, "TargetZone")
	delete(f, "SourceVpc")
	delete(f, "TargetVpc")
	delete(f, "SitePairProductType")
	delete(f, "SitePairName")
	delete(f, "CopyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDisasterRecoverySitePairRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDisasterRecoverySitePairResponseParams struct {
	// 容灾站点对ID
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDisasterRecoverySitePairResponse struct {
	*tchttp.BaseResponse
	Response *CreateDisasterRecoverySitePairResponseParams `json:"Response"`
}

func (r *CreateDisasterRecoverySitePairResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisasterRecoverySitePairResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDisasterRecoveryVpcMappingRequestParams struct {
	// 源端VPC ID
	SourceVpcId *string `json:"SourceVpcId,omitnil,omitempty" name:"SourceVpcId"`

	// 源端子网ID
	SourceSubnetId *string `json:"SourceSubnetId,omitnil,omitempty" name:"SourceSubnetId"`

	// 目标端VPC ID
	TargetVpcId *string `json:"TargetVpcId,omitnil,omitempty" name:"TargetVpcId"`

	// 目标端子网ID
	TargetSubnetId *string `json:"TargetSubnetId,omitnil,omitempty" name:"TargetSubnetId"`

	// 站点对ID
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`
}

type CreateDisasterRecoveryVpcMappingRequest struct {
	*tchttp.BaseRequest
	
	// 源端VPC ID
	SourceVpcId *string `json:"SourceVpcId,omitnil,omitempty" name:"SourceVpcId"`

	// 源端子网ID
	SourceSubnetId *string `json:"SourceSubnetId,omitnil,omitempty" name:"SourceSubnetId"`

	// 目标端VPC ID
	TargetVpcId *string `json:"TargetVpcId,omitnil,omitempty" name:"TargetVpcId"`

	// 目标端子网ID
	TargetSubnetId *string `json:"TargetSubnetId,omitnil,omitempty" name:"TargetSubnetId"`

	// 站点对ID
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`
}

func (r *CreateDisasterRecoveryVpcMappingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisasterRecoveryVpcMappingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceVpcId")
	delete(f, "SourceSubnetId")
	delete(f, "TargetVpcId")
	delete(f, "TargetSubnetId")
	delete(f, "SitePairId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDisasterRecoveryVpcMappingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDisasterRecoveryVpcMappingResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDisasterRecoveryVpcMappingResponse struct {
	*tchttp.BaseResponse
	Response *CreateDisasterRecoveryVpcMappingResponseParams `json:"Response"`
}

func (r *CreateDisasterRecoveryVpcMappingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisasterRecoveryVpcMappingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileBackupPlanRequestParams struct {
	// 备份策略ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 备份库ID
	BackupStorageId *string `json:"BackupStorageId,omitnil,omitempty" name:"BackupStorageId"`

	// 计划名称
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// 实例配置列表，[1,20]
	Resources []*ResourcePlan `json:"Resources,omitnil,omitempty" name:"Resources"`
}

type CreateFileBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// 备份策略ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 备份库ID
	BackupStorageId *string `json:"BackupStorageId,omitnil,omitempty" name:"BackupStorageId"`

	// 计划名称
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// 实例配置列表，[1,20]
	Resources []*ResourcePlan `json:"Resources,omitnil,omitempty" name:"Resources"`
}

func (r *CreateFileBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "BackupStorageId")
	delete(f, "PlanName")
	delete(f, "Resources")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFileBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileBackupPlanResponseParams struct {
	// 备份计划 ID 列表
	PlanIds []*string `json:"PlanIds,omitnil,omitempty" name:"PlanIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFileBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *CreateFileBackupPlanResponseParams `json:"Response"`
}

func (r *CreateFileBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileBackupRequestParams struct {
	// 资源ID列表
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 计划ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 备份路径列表，1~20 个
	BackupPaths []*string `json:"BackupPaths,omitnil,omitempty" name:"BackupPaths"`

	// 包含文件类型，0~20 个
	IncludeFileTypes []*string `json:"IncludeFileTypes,omitnil,omitempty" name:"IncludeFileTypes"`

	// 排除文件路径列表，0~20 个
	ExcludePatterns []*string `json:"ExcludePatterns,omitnil,omitempty" name:"ExcludePatterns"`

	// 是否排除系统目录
	ExcludeSystemDirectories *bool `json:"ExcludeSystemDirectories,omitnil,omitempty" name:"ExcludeSystemDirectories"`

	// 备份库ID
	BackupStorageId *string `json:"BackupStorageId,omitnil,omitempty" name:"BackupStorageId"`

	// 备份到期时间
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 备份名称
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
}

type CreateFileBackupRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID列表
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 计划ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 备份路径列表，1~20 个
	BackupPaths []*string `json:"BackupPaths,omitnil,omitempty" name:"BackupPaths"`

	// 包含文件类型，0~20 个
	IncludeFileTypes []*string `json:"IncludeFileTypes,omitnil,omitempty" name:"IncludeFileTypes"`

	// 排除文件路径列表，0~20 个
	ExcludePatterns []*string `json:"ExcludePatterns,omitnil,omitempty" name:"ExcludePatterns"`

	// 是否排除系统目录
	ExcludeSystemDirectories *bool `json:"ExcludeSystemDirectories,omitnil,omitempty" name:"ExcludeSystemDirectories"`

	// 备份库ID
	BackupStorageId *string `json:"BackupStorageId,omitnil,omitempty" name:"BackupStorageId"`

	// 备份到期时间
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 备份名称
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
}

func (r *CreateFileBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "PlanId")
	delete(f, "BackupPaths")
	delete(f, "IncludeFileTypes")
	delete(f, "ExcludePatterns")
	delete(f, "ExcludeSystemDirectories")
	delete(f, "BackupStorageId")
	delete(f, "Deadline")
	delete(f, "BackupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFileBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileBackupResponseParams struct {
	// 备份Id
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFileBackupResponse struct {
	*tchttp.BaseResponse
	Response *CreateFileBackupResponseParams `json:"Response"`
}

func (r *CreateFileBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileRestoreTaskRequestParams struct {
	// 冲突处理策略：skip-跳过/"         "overwrite-覆盖/newer-保留较新的版本/"         "if_changed-内容变化时覆盖，默认overwrite
	ConflictStrategy *string `json:"ConflictStrategy,omitnil,omitempty" name:"ConflictStrategy"`
}

type CreateFileRestoreTaskRequest struct {
	*tchttp.BaseRequest
	
	// 冲突处理策略：skip-跳过/"         "overwrite-覆盖/newer-保留较新的版本/"         "if_changed-内容变化时覆盖，默认overwrite
	ConflictStrategy *string `json:"ConflictStrategy,omitnil,omitempty" name:"ConflictStrategy"`
}

func (r *CreateFileRestoreTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileRestoreTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConflictStrategy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFileRestoreTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileRestoreTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFileRestoreTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateFileRestoreTaskResponseParams `json:"Response"`
}

func (r *CreateFileRestoreTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileRestoreTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceCopyPairRequestParams struct {
	// 所属保护组
	ProtectGroupId *string `json:"ProtectGroupId,omitnil,omitempty" name:"ProtectGroupId"`

	// 目标端CVM创建参数列表（1~10 个）
	CreateTargetInstanceParameters []*CreateInstanceModel `json:"CreateTargetInstanceParameters,omitnil,omitempty" name:"CreateTargetInstanceParameters"`

	// 复制对名称，不传则新名称为"未命名"
	InstanceCopyPairName *string `json:"InstanceCopyPairName,omitnil,omitempty" name:"InstanceCopyPairName"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 用户期望的RPO，单位分钟，目前仅支持15分钟
	RecoveryPointObjective *int64 `json:"RecoveryPointObjective,omitnil,omitempty" name:"RecoveryPointObjective"`
}

type CreateInstanceCopyPairRequest struct {
	*tchttp.BaseRequest
	
	// 所属保护组
	ProtectGroupId *string `json:"ProtectGroupId,omitnil,omitempty" name:"ProtectGroupId"`

	// 目标端CVM创建参数列表（1~10 个）
	CreateTargetInstanceParameters []*CreateInstanceModel `json:"CreateTargetInstanceParameters,omitnil,omitempty" name:"CreateTargetInstanceParameters"`

	// 复制对名称，不传则新名称为"未命名"
	InstanceCopyPairName *string `json:"InstanceCopyPairName,omitnil,omitempty" name:"InstanceCopyPairName"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 用户期望的RPO，单位分钟，目前仅支持15分钟
	RecoveryPointObjective *int64 `json:"RecoveryPointObjective,omitnil,omitempty" name:"RecoveryPointObjective"`
}

func (r *CreateInstanceCopyPairRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceCopyPairRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProtectGroupId")
	delete(f, "CreateTargetInstanceParameters")
	delete(f, "InstanceCopyPairName")
	delete(f, "ClientToken")
	delete(f, "RecoveryPointObjective")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceCopyPairRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceCopyPairResponseParams struct {
	// 创建的CVM复制对ID列表
	CopyPairIds []*string `json:"CopyPairIds,omitnil,omitempty" name:"CopyPairIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceCopyPairResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceCopyPairResponseParams `json:"Response"`
}

func (r *CreateInstanceCopyPairResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceCopyPairResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceDrillPairsRequestParams struct {
	// 所属容灾保护组
	ProtectGroupId *string `json:"ProtectGroupId,omitnil,omitempty" name:"ProtectGroupId"`

	// 演练组vpc
	DrillPairGroupVpc *string `json:"DrillPairGroupVpc,omitnil,omitempty" name:"DrillPairGroupVpc"`

	// 文件系统复制对名称,不传则新名称为“未命名”
	DrillPairGroupName *string `json:"DrillPairGroupName,omitnil,omitempty" name:"DrillPairGroupName"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性
	CreationToken *string `json:"CreationToken,omitnil,omitempty" name:"CreationToken"`

	// 指定创建入哪个演练组
	DrillPairGroupId *string `json:"DrillPairGroupId,omitnil,omitempty" name:"DrillPairGroupId"`

	// 创建目标演练实例的参数列表
	CreateTargetInstanceParameters []*CreateInstanceModel `json:"CreateTargetInstanceParameters,omitnil,omitempty" name:"CreateTargetInstanceParameters"`
}

type CreateInstanceDrillPairsRequest struct {
	*tchttp.BaseRequest
	
	// 所属容灾保护组
	ProtectGroupId *string `json:"ProtectGroupId,omitnil,omitempty" name:"ProtectGroupId"`

	// 演练组vpc
	DrillPairGroupVpc *string `json:"DrillPairGroupVpc,omitnil,omitempty" name:"DrillPairGroupVpc"`

	// 文件系统复制对名称,不传则新名称为“未命名”
	DrillPairGroupName *string `json:"DrillPairGroupName,omitnil,omitempty" name:"DrillPairGroupName"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性
	CreationToken *string `json:"CreationToken,omitnil,omitempty" name:"CreationToken"`

	// 指定创建入哪个演练组
	DrillPairGroupId *string `json:"DrillPairGroupId,omitnil,omitempty" name:"DrillPairGroupId"`

	// 创建目标演练实例的参数列表
	CreateTargetInstanceParameters []*CreateInstanceModel `json:"CreateTargetInstanceParameters,omitnil,omitempty" name:"CreateTargetInstanceParameters"`
}

func (r *CreateInstanceDrillPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceDrillPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProtectGroupId")
	delete(f, "DrillPairGroupVpc")
	delete(f, "DrillPairGroupName")
	delete(f, "CreationToken")
	delete(f, "DrillPairGroupId")
	delete(f, "CreateTargetInstanceParameters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceDrillPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceDrillPairsResponseParams struct {
	// 演练对ID列表
	DrillPairIds []*string `json:"DrillPairIds,omitnil,omitempty" name:"DrillPairIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceDrillPairsResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceDrillPairsResponseParams `json:"Response"`
}

func (r *CreateInstanceDrillPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceDrillPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceModel struct {
	// 源CVM ID
	SourceInstanceId *string `json:"SourceInstanceId,omitnil,omitempty" name:"SourceInstanceId"`

	// 实例计费模式
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 镜像ID
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 指定系统盘规格
	SystemDisk *DiskModel `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 实例类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 指定数据盘规格列表
	DataDisks []*DiskModel `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 私有网络相关信息配置
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil,omitempty" name:"VirtualPrivateCloud"`

	// 公网带宽相关信息设置
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 实例显示名称。不传则新实例名为"未命名"。最大长度不能超60个字节。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例登录设置
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 增强服务配置
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 竞价实例最高出价
	SpotPrice *string `json:"SpotPrice,omitnil,omitempty" name:"SpotPrice"`

	// 实例主机名
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 提供给实例使用的用户数据
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 放置群组ID
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 关机计费模式，默认关机收费（KEEP_CHARGING / STOP_CHARGING），仅 CreateInstanceCopyPair 场景生效
	StoppedMode *string `json:"StoppedMode,omitnil,omitempty" name:"StoppedMode"`

	// 容灾演练使用的复制对ID，仅 CreateInstanceDrillPairs 场景生效
	CopyPairId *string `json:"CopyPairId,omitnil,omitempty" name:"CopyPairId"`

	// 容灾演练的恢复时间点，仅 CreateInstanceDrillPairs 场景生效
	RecoveryTime *string `json:"RecoveryTime,omitnil,omitempty" name:"RecoveryTime"`
}

// Predefined struct for user
type CreateSecurityGroupMappingRequestParams struct {
	// 生产端实例绑定的安全组ID
	SrcSecurityGroupId *string `json:"SrcSecurityGroupId,omitnil,omitempty" name:"SrcSecurityGroupId"`

	// 容灾端实例绑定的安全组ID
	TargetSecurityGroupId *string `json:"TargetSecurityGroupId,omitnil,omitempty" name:"TargetSecurityGroupId"`

	// 安全组映射所属的站点对ID。
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`
}

type CreateSecurityGroupMappingRequest struct {
	*tchttp.BaseRequest
	
	// 生产端实例绑定的安全组ID
	SrcSecurityGroupId *string `json:"SrcSecurityGroupId,omitnil,omitempty" name:"SrcSecurityGroupId"`

	// 容灾端实例绑定的安全组ID
	TargetSecurityGroupId *string `json:"TargetSecurityGroupId,omitnil,omitempty" name:"TargetSecurityGroupId"`

	// 安全组映射所属的站点对ID。
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`
}

func (r *CreateSecurityGroupMappingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupMappingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SrcSecurityGroupId")
	delete(f, "TargetSecurityGroupId")
	delete(f, "SitePairId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityGroupMappingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupMappingResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSecurityGroupMappingResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityGroupMappingResponseParams `json:"Response"`
}

func (r *CreateSecurityGroupMappingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupMappingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CrossCloudDetails struct {
	// 源端云名称（跨云对端云名称）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceCloudName *string `json:"SourceCloudName,omitnil,omitempty" name:"SourceCloudName"`

	// 目标端云名称（跨云本端云名称）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetCloudName *string `json:"TargetCloudName,omitnil,omitempty" name:"TargetCloudName"`

	// 源端云AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceAppId *int64 `json:"SourceAppId,omitnil,omitempty" name:"SourceAppId"`

	// 源端云主账号 Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceUin *string `json:"SourceUin,omitnil,omitempty" name:"SourceUin"`

	// 源端云子账号 Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceSubAccountUin *string `json:"SourceSubAccountUin,omitnil,omitempty" name:"SourceSubAccountUin"`

	// 源端云用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceUserName *string `json:"SourceUserName,omitnil,omitempty" name:"SourceUserName"`

	// 目标端云AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetAppId *int64 `json:"TargetAppId,omitnil,omitempty" name:"TargetAppId"`

	// 目标端云主账号 Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetUin *string `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 目标端云子账号 Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetSubAccountUin *string `json:"TargetSubAccountUin,omitnil,omitempty" name:"TargetSubAccountUin"`

	// 对端云的地域显示名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeerRegionName *string `json:"PeerRegionName,omitnil,omitempty" name:"PeerRegionName"`

	// 对端云的可用区显示名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeerZoneName *string `json:"PeerZoneName,omitnil,omitempty" name:"PeerZoneName"`

	// 对端云的VPC显示名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeerVpcName *string `json:"PeerVpcName,omitnil,omitempty" name:"PeerVpcName"`
}

// Predefined struct for user
type DeleteAutoBackupPoliciesRequestParams struct {
	// 备份策略 ID 列表
	AutoBackupPolicyIds []*string `json:"AutoBackupPolicyIds,omitnil,omitempty" name:"AutoBackupPolicyIds"`
}

type DeleteAutoBackupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 备份策略 ID 列表
	AutoBackupPolicyIds []*string `json:"AutoBackupPolicyIds,omitnil,omitempty" name:"AutoBackupPolicyIds"`
}

func (r *DeleteAutoBackupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoBackupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoBackupPolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAutoBackupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAutoBackupPoliciesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAutoBackupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAutoBackupPoliciesResponseParams `json:"Response"`
}

func (r *DeleteAutoBackupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoBackupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupGroupsRequestParams struct {
	// 备份组ID列表。
	BackupGroupIds []*string `json:"BackupGroupIds,omitnil,omitempty" name:"BackupGroupIds"`
}

type DeleteBackupGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 备份组ID列表。
	BackupGroupIds []*string `json:"BackupGroupIds,omitnil,omitempty" name:"BackupGroupIds"`
}

func (r *DeleteBackupGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBackupGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBackupGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBackupGroupsResponseParams `json:"Response"`
}

func (r *DeleteBackupGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupVaultsRequestParams struct {
	// 备份库 ID 列表
	VaultIds []*string `json:"VaultIds,omitnil,omitempty" name:"VaultIds"`
}

type DeleteBackupVaultsRequest struct {
	*tchttp.BaseRequest
	
	// 备份库 ID 列表
	VaultIds []*string `json:"VaultIds,omitnil,omitempty" name:"VaultIds"`
}

func (r *DeleteBackupVaultsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupVaultsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VaultIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBackupVaultsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupVaultsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBackupVaultsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBackupVaultsResponseParams `json:"Response"`
}

func (r *DeleteBackupVaultsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupVaultsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCopyPairsRequestParams struct {
	// 要删除的复制对ID列表（长度 1~10）
	CopyPairIds []*string `json:"CopyPairIds,omitnil,omitempty" name:"CopyPairIds"`

	// 要删除复制对的类型，可选值：DISK、INSTANCE、CFS
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`

	// 是否一并删除容灾站点云盘，默认 true（容灾盘数据可能处于中间状态，保留也无法正常使用）
	DeleteTargetResource *bool `json:"DeleteTargetResource,omitnil,omitempty" name:"DeleteTargetResource"`
}

type DeleteCopyPairsRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的复制对ID列表（长度 1~10）
	CopyPairIds []*string `json:"CopyPairIds,omitnil,omitempty" name:"CopyPairIds"`

	// 要删除复制对的类型，可选值：DISK、INSTANCE、CFS
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`

	// 是否一并删除容灾站点云盘，默认 true（容灾盘数据可能处于中间状态，保留也无法正常使用）
	DeleteTargetResource *bool `json:"DeleteTargetResource,omitnil,omitempty" name:"DeleteTargetResource"`
}

func (r *DeleteCopyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCopyPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CopyPairIds")
	delete(f, "CopyPairType")
	delete(f, "DeleteTargetResource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCopyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCopyPairsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCopyPairsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCopyPairsResponseParams `json:"Response"`
}

func (r *DeleteCopyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCopyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDisasterRecoveryProtectGroupsRequestParams struct {
	// 删除容灾保护组ID列表，最多10个
	ProtectGroups []*string `json:"ProtectGroups,omitnil,omitempty" name:"ProtectGroups"`
}

type DeleteDisasterRecoveryProtectGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 删除容灾保护组ID列表，最多10个
	ProtectGroups []*string `json:"ProtectGroups,omitnil,omitempty" name:"ProtectGroups"`
}

func (r *DeleteDisasterRecoveryProtectGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDisasterRecoveryProtectGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProtectGroups")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDisasterRecoveryProtectGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDisasterRecoveryProtectGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDisasterRecoveryProtectGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDisasterRecoveryProtectGroupsResponseParams `json:"Response"`
}

func (r *DeleteDisasterRecoveryProtectGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDisasterRecoveryProtectGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDisasterRecoverySitePairsRequestParams struct {
	// 删除容灾策略ID列表
	SitePairIds []*string `json:"SitePairIds,omitnil,omitempty" name:"SitePairIds"`
}

type DeleteDisasterRecoverySitePairsRequest struct {
	*tchttp.BaseRequest
	
	// 删除容灾策略ID列表
	SitePairIds []*string `json:"SitePairIds,omitnil,omitempty" name:"SitePairIds"`
}

func (r *DeleteDisasterRecoverySitePairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDisasterRecoverySitePairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SitePairIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDisasterRecoverySitePairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDisasterRecoverySitePairsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDisasterRecoverySitePairsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDisasterRecoverySitePairsResponseParams `json:"Response"`
}

func (r *DeleteDisasterRecoverySitePairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDisasterRecoverySitePairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDisasterRecoveryVpcMappingRequestParams struct {
	// 删除容灾vpc映射主键id列表
	VpcMappingIds []*uint64 `json:"VpcMappingIds,omitnil,omitempty" name:"VpcMappingIds"`
}

type DeleteDisasterRecoveryVpcMappingRequest struct {
	*tchttp.BaseRequest
	
	// 删除容灾vpc映射主键id列表
	VpcMappingIds []*uint64 `json:"VpcMappingIds,omitnil,omitempty" name:"VpcMappingIds"`
}

func (r *DeleteDisasterRecoveryVpcMappingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDisasterRecoveryVpcMappingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcMappingIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDisasterRecoveryVpcMappingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDisasterRecoveryVpcMappingResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDisasterRecoveryVpcMappingResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDisasterRecoveryVpcMappingResponseParams `json:"Response"`
}

func (r *DeleteDisasterRecoveryVpcMappingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDisasterRecoveryVpcMappingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDrillPairResult struct {
	// 演练对ID。
	DrillPairId *string `json:"DrillPairId,omitnil,omitempty" name:"DrillPairId"`

	// 删除结果码。成功为 Success，失败为对应错误码（如 InternalError.ComponentError）。
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 删除结果描述信息，成功时为空串。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type DeleteDrillPairsRequestParams struct {
	// 要删除演练对的类型，其类型枚举跟复制对保持一致。枚举值：DISK / INSTANCE / CFS。
	DrillPairType *string `json:"DrillPairType,omitnil,omitempty" name:"DrillPairType"`

	// 要删除的演练对列表。长度范围 [1, 10]。
	DrillPairIds []*string `json:"DrillPairIds,omitnil,omitempty" name:"DrillPairIds"`

	// 要删除的演练组id列表。
	DrillGroupIds []*string `json:"DrillGroupIds,omitnil,omitempty" name:"DrillGroupIds"`

	// 是否一并删除演练CFS/CVM/DISK演练资源。
	DeleteDrillResource *bool `json:"DeleteDrillResource,omitnil,omitempty" name:"DeleteDrillResource"`
}

type DeleteDrillPairsRequest struct {
	*tchttp.BaseRequest
	
	// 要删除演练对的类型，其类型枚举跟复制对保持一致。枚举值：DISK / INSTANCE / CFS。
	DrillPairType *string `json:"DrillPairType,omitnil,omitempty" name:"DrillPairType"`

	// 要删除的演练对列表。长度范围 [1, 10]。
	DrillPairIds []*string `json:"DrillPairIds,omitnil,omitempty" name:"DrillPairIds"`

	// 要删除的演练组id列表。
	DrillGroupIds []*string `json:"DrillGroupIds,omitnil,omitempty" name:"DrillGroupIds"`

	// 是否一并删除演练CFS/CVM/DISK演练资源。
	DeleteDrillResource *bool `json:"DeleteDrillResource,omitnil,omitempty" name:"DeleteDrillResource"`
}

func (r *DeleteDrillPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDrillPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DrillPairType")
	delete(f, "DrillPairIds")
	delete(f, "DrillGroupIds")
	delete(f, "DeleteDrillResource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDrillPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDrillPairsResponseParams struct {
	// 删除演练对的逐条结果列表。
	DeleteDrillPairResultSet []*DeleteDrillPairResult `json:"DeleteDrillPairResultSet,omitnil,omitempty" name:"DeleteDrillPairResultSet"`

	// 成功标记为删除的演练组ID列表。
	DeleteDrillPairGroupSet []*string `json:"DeleteDrillPairGroupSet,omitnil,omitempty" name:"DeleteDrillPairGroupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDrillPairsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDrillPairsResponseParams `json:"Response"`
}

func (r *DeleteDrillPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDrillPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFileBackupPlansRequestParams struct {
	// 备份点 ID 列表
	PlanIds []*string `json:"PlanIds,omitnil,omitempty" name:"PlanIds"`
}

type DeleteFileBackupPlansRequest struct {
	*tchttp.BaseRequest
	
	// 备份点 ID 列表
	PlanIds []*string `json:"PlanIds,omitnil,omitempty" name:"PlanIds"`
}

func (r *DeleteFileBackupPlansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileBackupPlansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFileBackupPlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFileBackupPlansResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFileBackupPlansResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFileBackupPlansResponseParams `json:"Response"`
}

func (r *DeleteFileBackupPlansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileBackupPlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFileBackupsRequestParams struct {

}

type DeleteFileBackupsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DeleteFileBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFileBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFileBackupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFileBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFileBackupsResponseParams `json:"Response"`
}

func (r *DeleteFileBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupMappingRequestParams struct {
	// 要删除安全组映射所属的站点对ID
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 要删除的安全组映射ID列表
	SecurityGroupMappingIds []*string `json:"SecurityGroupMappingIds,omitnil,omitempty" name:"SecurityGroupMappingIds"`
}

type DeleteSecurityGroupMappingRequest struct {
	*tchttp.BaseRequest
	
	// 要删除安全组映射所属的站点对ID
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 要删除的安全组映射ID列表
	SecurityGroupMappingIds []*string `json:"SecurityGroupMappingIds,omitnil,omitempty" name:"SecurityGroupMappingIds"`
}

func (r *DeleteSecurityGroupMappingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupMappingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SitePairId")
	delete(f, "SecurityGroupMappingIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityGroupMappingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupMappingResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSecurityGroupMappingResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityGroupMappingResponseParams `json:"Response"`
}

func (r *DeleteSecurityGroupMappingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupMappingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeniedAction struct {
	// 不能操作的接口名。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 接口不能操作的原因。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 接口不能操作对应提示的错误码。
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`
}

// Predefined struct for user
type DescribeAutoBackupPoliciesRequestParams struct {
	// 过滤条件。支持以下过滤条件：\n"              "auto-backup-policy-id - 定期快照策略ID，如asp-xxx。\n"              "auto-backup-policy-state - 定期快照策略状态。\n"              "auto-backup-policy-name - 定期快照策略名称，支持模糊匹配。\n"              "tag - 按标签键值对过滤，需包含Key和/或Value。\n"              "tag-key - 按标签键过滤。\n"              "tag-value - 按标签值过滤。\n"              "tag:tag-key - 按指定标签键的标签值过滤。\n"              "vault-id - 备份库ID过滤。\n"              "storage-type - 存储类型过滤"              "（COMMON：普通模式，VAULT：备份库模式）。
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeAutoBackupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。支持以下过滤条件：\n"              "auto-backup-policy-id - 定期快照策略ID，如asp-xxx。\n"              "auto-backup-policy-state - 定期快照策略状态。\n"              "auto-backup-policy-name - 定期快照策略名称，支持模糊匹配。\n"              "tag - 按标签键值对过滤，需包含Key和/或Value。\n"              "tag-key - 按标签键过滤。\n"              "tag-value - 按标签值过滤。\n"              "tag:tag-key - 按指定标签键的标签值过滤。\n"              "vault-id - 备份库ID过滤。\n"              "storage-type - 存储类型过滤"              "（COMMON：普通模式，VAULT：备份库模式）。
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeAutoBackupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoBackupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoBackupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoBackupPoliciesResponseParams struct {
	// 符合条件的总数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 备份策略列表详情。
	AutoBackupPolicySet []*AutoBackupPolicy `json:"AutoBackupPolicySet,omitnil,omitempty" name:"AutoBackupPolicySet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoBackupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoBackupPoliciesResponseParams `json:"Response"`
}

func (r *DescribeAutoBackupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoBackupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupGroupRollbackTasksRequestParams struct {
	// 过滤条件，支持恢复任务ID（task-id）、备份组ID（backup-group-id）、源实例ID（source-instance-id）、目标实例ID（target-instance-id）、恢复状态（status）和回滚类型（rollback-type）过滤
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeBackupGroupRollbackTasksRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，支持恢复任务ID（task-id）、备份组ID（backup-group-id）、源实例ID（source-instance-id）、目标实例ID（target-instance-id）、恢复状态（status）和回滚类型（rollback-type）过滤
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeBackupGroupRollbackTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupGroupRollbackTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupGroupRollbackTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupGroupRollbackTasksResponseParams struct {
	// 符合条件的总数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 备份组恢复详情。
	RollbackTaskSet []*BackupGroupRollbackTask `json:"RollbackTaskSet,omitnil,omitempty" name:"RollbackTaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupGroupRollbackTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupGroupRollbackTasksResponseParams `json:"Response"`
}

func (r *DescribeBackupGroupRollbackTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupGroupRollbackTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupGroupsDeniedActionsRequestParams struct {
	// 备份组列表
	BackupGroupIds []*string `json:"BackupGroupIds,omitnil,omitempty" name:"BackupGroupIds"`
}

type DescribeBackupGroupsDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// 备份组列表
	BackupGroupIds []*string `json:"BackupGroupIds,omitnil,omitempty" name:"BackupGroupIds"`
}

func (r *DescribeBackupGroupsDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupGroupsDeniedActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupGroupsDeniedActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupGroupsDeniedActionsResponseParams struct {
	// 备份组不允许操作信息
	BackupGroupDeniedActionSet []*BackupGroupDeniedAction `json:"BackupGroupDeniedActionSet,omitnil,omitempty" name:"BackupGroupDeniedActionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupGroupsDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupGroupsDeniedActionsResponseParams `json:"Response"`
}

func (r *DescribeBackupGroupsDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupGroupsDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupGroupsRequestParams struct {
	// 过滤条件。backup-group-id - Array of String - 是否必填：否 -（过滤条件）按备份组ID过滤 ;backup-group-state - Array of String - 是否必填：否 -（过滤条件）按备份组状态过滤。(NORMAL: 正常 | CREATING:创建中 | ROLLBACKING:回滚中) ;backup-group-name - Array of String - 是否必填：否 -（过滤条件）按备份组名称过滤 ;backup-id - Array of String - 是否必填：否 -（过滤条件）按备份组内的备份ID过滤
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段。目前支持CREATE_TIME。
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeBackupGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。backup-group-id - Array of String - 是否必填：否 -（过滤条件）按备份组ID过滤 ;backup-group-state - Array of String - 是否必填：否 -（过滤条件）按备份组状态过滤。(NORMAL: 正常 | CREATING:创建中 | ROLLBACKING:回滚中) ;backup-group-name - Array of String - 是否必填：否 -（过滤条件）按备份组名称过滤 ;backup-id - Array of String - 是否必填：否 -（过滤条件）按备份组内的备份ID过滤
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段。目前支持CREATE_TIME。
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeBackupGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupGroupsResponseParams struct {
	// 符合条件的总数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 备份列表详情。
	BackupGroupSet []*BackupGroup `json:"BackupGroupSet,omitnil,omitempty" name:"BackupGroupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupGroupsResponseParams `json:"Response"`
}

func (r *DescribeBackupGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupInstancesRequestParams struct {
	// 过滤条件。;instance-id - Array of String - 是否必填：否 -（过滤条件）按实例ID过滤。;auto-backup-policy-id - Array of String - 是否必填：否 -（过滤条件）按照实例绑定的定期备份策略过滤。;auto-backup-policy-name - Array of String - 是否必填：否 -（过滤条件）按照云硬盘绑定的定期备份策略名称过滤。
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeBackupInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。;instance-id - Array of String - 是否必填：否 -（过滤条件）按实例ID过滤。;auto-backup-policy-id - Array of String - 是否必填：否 -（过滤条件）按照实例绑定的定期备份策略过滤。;auto-backup-policy-name - Array of String - 是否必填：否 -（过滤条件）按照云硬盘绑定的定期备份策略名称过滤。
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeBackupInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupInstancesResponseParams struct {
	// 符合条件的受保护实例总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合条件的受保护实例详情
	BackupInstanceSet []*BackupInstance `json:"BackupInstanceSet,omitnil,omitempty" name:"BackupInstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupInstancesResponseParams `json:"Response"`
}

func (r *DescribeBackupInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupOverviewGeneralRequestParams struct {
	// <p>是否查询全部地域。false-仅当前地域（默认），true-全部地域汇总</p>
	AllRegions *bool `json:"AllRegions,omitnil,omitempty" name:"AllRegions"`
}

type DescribeBackupOverviewGeneralRequest struct {
	*tchttp.BaseRequest
	
	// <p>是否查询全部地域。false-仅当前地域（默认），true-全部地域汇总</p>
	AllRegions *bool `json:"AllRegions,omitnil,omitempty" name:"AllRegions"`
}

func (r *DescribeBackupOverviewGeneralRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupOverviewGeneralRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AllRegions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupOverviewGeneralRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupOverviewGeneralResponseParams struct {
	// <p>整机备份（CVM 备份组）概览数据</p>
	InstanceBackupOverview *InstanceBackupOverview `json:"InstanceBackupOverview,omitnil,omitempty" name:"InstanceBackupOverview"`

	// <p>文件备份概览数据</p>
	FileBackupOverview *FileBackupOverview `json:"FileBackupOverview,omitnil,omitempty" name:"FileBackupOverview"`

	// <p>备份策略概览</p>
	BackupPolicyOverview *BackupPolicyOverview `json:"BackupPolicyOverview,omitnil,omitempty" name:"BackupPolicyOverview"`

	// <p>备份库概览</p>
	BackupVaultOverview *BackupVaultOverview `json:"BackupVaultOverview,omitnil,omitempty" name:"BackupVaultOverview"`

	// <p>受保护资源概览</p>
	ProtectedResourceOverview *ProtectedResourceOverview `json:"ProtectedResourceOverview,omitnil,omitempty" name:"ProtectedResourceOverview"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupOverviewGeneralResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupOverviewGeneralResponseParams `json:"Response"`
}

func (r *DescribeBackupOverviewGeneralResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupOverviewGeneralResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupPlansRequestParams struct {
	// 过滤条件，支持instance-id和auto-backup-policy-id
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeBackupPlansRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，支持instance-id和auto-backup-policy-id
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeBackupPlansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupPlansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupPlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupPlansResponseParams struct {
	// 符合条件的总数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 备份列表详情。
	BackupPlanSet []*BackupPlan `json:"BackupPlanSet,omitnil,omitempty" name:"BackupPlanSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupPlansResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupPlansResponseParams `json:"Response"`
}

func (r *DescribeBackupPlansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupPlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupVaultsDeniedActionsRequestParams struct {
	// 备份库ID列表
	VaultIds []*string `json:"VaultIds,omitnil,omitempty" name:"VaultIds"`
}

type DescribeBackupVaultsDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// 备份库ID列表
	VaultIds []*string `json:"VaultIds,omitnil,omitempty" name:"VaultIds"`
}

func (r *DescribeBackupVaultsDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupVaultsDeniedActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VaultIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupVaultsDeniedActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupVaultsDeniedActionsResponseParams struct {
	// 备份库不允许操作信息
	BackupVaultDeniedActionSet []*VaultDeniedAction `json:"BackupVaultDeniedActionSet,omitnil,omitempty" name:"BackupVaultDeniedActionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupVaultsDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupVaultsDeniedActionsResponseParams `json:"Response"`
}

func (r *DescribeBackupVaultsDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupVaultsDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupVaultsRequestParams struct {
	// 备份库ID列表
	VaultIds []*string `json:"VaultIds,omitnil,omitempty" name:"VaultIds"`

	// 过滤条件，支持instance-id和auto-backup-policy-id
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeBackupVaultsRequest struct {
	*tchttp.BaseRequest
	
	// 备份库ID列表
	VaultIds []*string `json:"VaultIds,omitnil,omitempty" name:"VaultIds"`

	// 过滤条件，支持instance-id和auto-backup-policy-id
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeBackupVaultsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupVaultsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VaultIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupVaultsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupVaultsResponseParams struct {
	// 符合条件的总数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 备份库列表详情。
	BackupVaultSet []*BackupVault `json:"BackupVaultSet,omitnil,omitempty" name:"BackupVaultSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupVaultsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupVaultsResponseParams `json:"Response"`
}

func (r *DescribeBackupVaultsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupVaultsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCommonBackupPointsRequestParams struct {
	// 实例列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeCommonBackupPointsRequest struct {
	*tchttp.BaseRequest
	
	// 实例列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DescribeCommonBackupPointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCommonBackupPointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCommonBackupPointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCommonBackupPointsResponseParams struct {
	// 符合条件的总数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 共同备份点详情。
	CommonBackupPointSet []*CommonBackupPoint `json:"CommonBackupPointSet,omitnil,omitempty" name:"CommonBackupPointSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCommonBackupPointsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCommonBackupPointsResponseParams `json:"Response"`
}

func (r *DescribeCommonBackupPointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCommonBackupPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCopyPairsDeniedActionsRequestParams struct {
	// 复制对ID列表
	CopyPairIds []*string `json:"CopyPairIds,omitnil,omitempty" name:"CopyPairIds"`

	// 要查询复制对的类型，枚举值：DISK（云硬盘）、INSTANCE（云服务器）、CFS（文件存储）
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`
}

type DescribeCopyPairsDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// 复制对ID列表
	CopyPairIds []*string `json:"CopyPairIds,omitnil,omitempty" name:"CopyPairIds"`

	// 要查询复制对的类型，枚举值：DISK（云硬盘）、INSTANCE（云服务器）、CFS（文件存储）
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`
}

func (r *DescribeCopyPairsDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCopyPairsDeniedActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CopyPairIds")
	delete(f, "CopyPairType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCopyPairsDeniedActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCopyPairsDeniedActionsResponseParams struct {
	// 复制对操作掩码列表，返回每个复制对被禁止执行的操作
	CopyPairDeniedActionSet []*CopyPairDeniedAction `json:"CopyPairDeniedActionSet,omitnil,omitempty" name:"CopyPairDeniedActionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCopyPairsDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCopyPairsDeniedActionsResponseParams `json:"Response"`
}

func (r *DescribeCopyPairsDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCopyPairsDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCopyPairsRequestParams struct {
	// <p>要查询复制对的类型，可选值：DISK、INSTANCE、CFS</p>
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`

	// <p>要查询复制对ID列表</p>
	CopyPairIds []*string `json:"CopyPairIds,omitnil,omitempty" name:"CopyPairIds"`

	// <p>过滤条件，详见过滤条件表。支持的Name：disaster-recovery-site-pair-id、target-resource-id、source-resource-id、copy-pair-id、copy-pair-name</p>
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量，默认为20，最大值为100。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>输出结果按升序还是降序，可选值：ASC、DESC</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>输出结果的排序字段，可选值：CREATE_TIME</p>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// <p>是否要查询保护时间点列表，默认 false。当设置为 true 时，必须同时传入 CopyPairIds 参数。</p>
	QueryProtectionTime *bool `json:"QueryProtectionTime,omitnil,omitempty" name:"QueryProtectionTime"`

	// <p>是否查询跨云+非跨云全部复制对，默认 false</p>
	GetAllCopyPair *bool `json:"GetAllCopyPair,omitnil,omitempty" name:"GetAllCopyPair"`

	// <p>是否要查询 CVM 创建参数（仅对延迟创建模式且目标 CVM 未创建的复制对生效），默认为true。为 true 时，每条 deferred_create=1 AND target_cvm_created=0 的 CVM 复制对出参会附带 CvmCreateParams 字段</p>
	QueryCvmCreateParams *bool `json:"QueryCvmCreateParams,omitnil,omitempty" name:"QueryCvmCreateParams"`

	// <p>复制对创建来源过滤。不传则查询所有；传 LOCAL 仅查本端创建的复制对，传 PEER 仅查对端创建的复制对。</p><p>枚举值：</p><ul><li>LOCAL： 仅查本端创建的复制对</li><li>PEER： 仅查对端创建的复制对</li></ul>
	CreateFrom *string `json:"CreateFrom,omitnil,omitempty" name:"CreateFrom"`
}

type DescribeCopyPairsRequest struct {
	*tchttp.BaseRequest
	
	// <p>要查询复制对的类型，可选值：DISK、INSTANCE、CFS</p>
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`

	// <p>要查询复制对ID列表</p>
	CopyPairIds []*string `json:"CopyPairIds,omitnil,omitempty" name:"CopyPairIds"`

	// <p>过滤条件，详见过滤条件表。支持的Name：disaster-recovery-site-pair-id、target-resource-id、source-resource-id、copy-pair-id、copy-pair-name</p>
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量，默认为20，最大值为100。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>输出结果按升序还是降序，可选值：ASC、DESC</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>输出结果的排序字段，可选值：CREATE_TIME</p>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// <p>是否要查询保护时间点列表，默认 false。当设置为 true 时，必须同时传入 CopyPairIds 参数。</p>
	QueryProtectionTime *bool `json:"QueryProtectionTime,omitnil,omitempty" name:"QueryProtectionTime"`

	// <p>是否查询跨云+非跨云全部复制对，默认 false</p>
	GetAllCopyPair *bool `json:"GetAllCopyPair,omitnil,omitempty" name:"GetAllCopyPair"`

	// <p>是否要查询 CVM 创建参数（仅对延迟创建模式且目标 CVM 未创建的复制对生效），默认为true。为 true 时，每条 deferred_create=1 AND target_cvm_created=0 的 CVM 复制对出参会附带 CvmCreateParams 字段</p>
	QueryCvmCreateParams *bool `json:"QueryCvmCreateParams,omitnil,omitempty" name:"QueryCvmCreateParams"`

	// <p>复制对创建来源过滤。不传则查询所有；传 LOCAL 仅查本端创建的复制对，传 PEER 仅查对端创建的复制对。</p><p>枚举值：</p><ul><li>LOCAL： 仅查本端创建的复制对</li><li>PEER： 仅查对端创建的复制对</li></ul>
	CreateFrom *string `json:"CreateFrom,omitnil,omitempty" name:"CreateFrom"`
}

func (r *DescribeCopyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCopyPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CopyPairType")
	delete(f, "CopyPairIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "QueryProtectionTime")
	delete(f, "GetAllCopyPair")
	delete(f, "QueryCvmCreateParams")
	delete(f, "CreateFrom")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCopyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCopyPairsResponseParams struct {
	// <p>符合条件的复制对总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>复制对列表。</p>
	CopyPairSet []*CopyPair `json:"CopyPairSet,omitnil,omitempty" name:"CopyPairSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCopyPairsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCopyPairsResponseParams `json:"Response"`
}

func (r *DescribeCopyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCopyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoveryDrillGroupsRequestParams struct {
	// 要查询的容灾演练组产品类型。枚举值：DISK / INSTANCE / CFS。
	DrillGroupType *string `json:"DrillGroupType,omitnil,omitempty" name:"DrillGroupType"`

	// 要查询的容灾演练组ID列表。
	DrillGroupIds []*string `json:"DrillGroupIds,omitnil,omitempty" name:"DrillGroupIds"`

	// 过滤条件，详见定期快照过滤条件表。
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出结果按升序还是降序。枚举值：ASC / DESC。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 输出结果的排序字段。枚举值：CREATE_TIME。
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeDisasterRecoveryDrillGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的容灾演练组产品类型。枚举值：DISK / INSTANCE / CFS。
	DrillGroupType *string `json:"DrillGroupType,omitnil,omitempty" name:"DrillGroupType"`

	// 要查询的容灾演练组ID列表。
	DrillGroupIds []*string `json:"DrillGroupIds,omitnil,omitempty" name:"DrillGroupIds"`

	// 过滤条件，详见定期快照过滤条件表。
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出结果按升序还是降序。枚举值：ASC / DESC。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 输出结果的排序字段。枚举值：CREATE_TIME。
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeDisasterRecoveryDrillGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoveryDrillGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DrillGroupType")
	delete(f, "DrillGroupIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisasterRecoveryDrillGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoveryDrillGroupsResponseParams struct {
	// 有效的容灾演练组数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 容灾演练组列表。
	DrillGroupSet []*DisasterRecoveryDrillGroup `json:"DrillGroupSet,omitnil,omitempty" name:"DrillGroupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDisasterRecoveryDrillGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDisasterRecoveryDrillGroupsResponseParams `json:"Response"`
}

func (r *DescribeDisasterRecoveryDrillGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoveryDrillGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoveryOverviewRequestParams struct {
	// 要查询的产品/复制对的类型，枚举值：• DISK：云硬盘类型复制对• INSTANCE：CVM 实例复制对• CFS：文件存储复制对• ALL：聚合当前支持的类型；默认为CFS
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`
}

type DescribeDisasterRecoveryOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的产品/复制对的类型，枚举值：• DISK：云硬盘类型复制对• INSTANCE：CVM 实例复制对• CFS：文件存储复制对• ALL：聚合当前支持的类型；默认为CFS
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`
}

func (r *DescribeDisasterRecoveryOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoveryOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CopyPairType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisasterRecoveryOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoveryOverviewResponseParams struct {
	// 跨所有地域聚合后的容灾总览数据
	DisasterRecoveryOverview *DisasterRecoveryOverview `json:"DisasterRecoveryOverview,omitnil,omitempty" name:"DisasterRecoveryOverview"`

	// 按地域拆分的容灾总览列表
	OverviewInRegionSet []*DisasterRecoveryOverview `json:"OverviewInRegionSet,omitnil,omitempty" name:"OverviewInRegionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDisasterRecoveryOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDisasterRecoveryOverviewResponseParams `json:"Response"`
}

func (r *DescribeDisasterRecoveryOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoveryOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoveryProtectGroupsRequestParams struct {
	// 要查询的容灾保护组产品类型，枚举值：DISK / INSTANCE / CFS。
	ProtectGroupType *string `json:"ProtectGroupType,omitnil,omitempty" name:"ProtectGroupType"`

	// 要查询的容灾保护组ID列表。
	ProtectGroupIds []*string `json:"ProtectGroupIds,omitnil,omitempty" name:"ProtectGroupIds"`

	// 过滤条件（过滤项由 core handler 定义，如 disaster-recovery-protect-group-id 等）。
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出结果按升序还是降序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 输出结果的排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeDisasterRecoveryProtectGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的容灾保护组产品类型，枚举值：DISK / INSTANCE / CFS。
	ProtectGroupType *string `json:"ProtectGroupType,omitnil,omitempty" name:"ProtectGroupType"`

	// 要查询的容灾保护组ID列表。
	ProtectGroupIds []*string `json:"ProtectGroupIds,omitnil,omitempty" name:"ProtectGroupIds"`

	// 过滤条件（过滤项由 core handler 定义，如 disaster-recovery-protect-group-id 等）。
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出结果按升序还是降序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 输出结果的排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeDisasterRecoveryProtectGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoveryProtectGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProtectGroupType")
	delete(f, "ProtectGroupIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisasterRecoveryProtectGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoveryProtectGroupsResponseParams struct {
	// 符合条件的容灾保护组总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 容灾保护组列表
	ProtectGroupSet []*ProtectGroup `json:"ProtectGroupSet,omitnil,omitempty" name:"ProtectGroupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDisasterRecoveryProtectGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDisasterRecoveryProtectGroupsResponseParams `json:"Response"`
}

func (r *DescribeDisasterRecoveryProtectGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoveryProtectGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoverySitePairsDeniedActionsRequestParams struct {
	// 要查询的容灾策略ID列表，单个ID格式为 sitepair-xxxxxxxx
	SitePairIds []*string `json:"SitePairIds,omitnil,omitempty" name:"SitePairIds"`
}

type DescribeDisasterRecoverySitePairsDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的容灾策略ID列表，单个ID格式为 sitepair-xxxxxxxx
	SitePairIds []*string `json:"SitePairIds,omitnil,omitempty" name:"SitePairIds"`
}

func (r *DescribeDisasterRecoverySitePairsDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoverySitePairsDeniedActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SitePairIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisasterRecoverySitePairsDeniedActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoverySitePairsDeniedActionsResponseParams struct {
	// 每个容灾策略对应的禁止操作集合，返回顺序与入参 SitePairIds 一致
	SitePairDeniedActionSet []*SitePairDeniedAction `json:"SitePairDeniedActionSet,omitnil,omitempty" name:"SitePairDeniedActionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDisasterRecoverySitePairsDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDisasterRecoverySitePairsDeniedActionsResponseParams `json:"Response"`
}

func (r *DescribeDisasterRecoverySitePairsDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoverySitePairsDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoverySitePairsRequestParams struct {
	// 要查询的容灾策略产品类型。取值范围：DISK / INSTANCE / CFS。
	SitePairType *string `json:"SitePairType,omitnil,omitempty" name:"SitePairType"`

	// 要查询的容灾策略ID列表。
	SitePairIds []*string `json:"SitePairIds,omitnil,omitempty" name:"SitePairIds"`

	// 过滤条件，详见定期快照过滤条件表。
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出结果按升序还是降序，DESC表示降序，ASC表示升序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 输出结果的排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeDisasterRecoverySitePairsRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的容灾策略产品类型。取值范围：DISK / INSTANCE / CFS。
	SitePairType *string `json:"SitePairType,omitnil,omitempty" name:"SitePairType"`

	// 要查询的容灾策略ID列表。
	SitePairIds []*string `json:"SitePairIds,omitnil,omitempty" name:"SitePairIds"`

	// 过滤条件，详见定期快照过滤条件表。
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出结果按升序还是降序，DESC表示降序，ASC表示升序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 输出结果的排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeDisasterRecoverySitePairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoverySitePairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SitePairType")
	delete(f, "SitePairIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisasterRecoverySitePairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoverySitePairsResponseParams struct {
	// 有效的容灾策略数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 容灾策略列表。
	SitePairSet []*SitePair `json:"SitePairSet,omitnil,omitempty" name:"SitePairSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDisasterRecoverySitePairsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDisasterRecoverySitePairsResponseParams `json:"Response"`
}

func (r *DescribeDisasterRecoverySitePairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoverySitePairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoverySupportRegionRequestParams struct {
	// <p>状态过滤：valid（生效）/ invalid（停用）；为空则同时返回生效与停用的全部记录。</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type DescribeDisasterRecoverySupportRegionRequest struct {
	*tchttp.BaseRequest
	
	// <p>状态过滤：valid（生效）/ invalid（停用）；为空则同时返回生效与停用的全部记录。</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *DescribeDisasterRecoverySupportRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoverySupportRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisasterRecoverySupportRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoverySupportRegionResponseParams struct {
	// <p>符合条件的支持的生产地域配置总数。</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>支持的生产地域配置详情列表。</p>
	SupportRegionSet []*SupportRegionInfo `json:"SupportRegionSet,omitnil,omitempty" name:"SupportRegionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDisasterRecoverySupportRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDisasterRecoverySupportRegionResponseParams `json:"Response"`
}

func (r *DescribeDisasterRecoverySupportRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoverySupportRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksRequestParams struct {
	// 要查询信息的云盘ID列表
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 云盘所在地域
	DiskRegion *string `json:"DiskRegion,omitnil,omitempty" name:"DiskRegion"`
}

type DescribeDisksRequest struct {
	*tchttp.BaseRequest
	
	// 要查询信息的云盘ID列表
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 云盘所在地域
	DiskRegion *string `json:"DiskRegion,omitnil,omitempty" name:"DiskRegion"`
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
	delete(f, "DiskRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksResponseParams struct {
	// 符合条件的云盘总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 云盘详情列表
	DiskInfoSet []*DiskInfo `json:"DiskInfoSet,omitnil,omitempty" name:"DiskInfoSet"`

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
type DescribeDrillPairsDeniedActionsRequestParams struct {
	// 要查询演练对的类型，枚举值：DISK（云硬盘）、INSTANCE（云服务器）、CFS（文件存储）
	DrillPairType *string `json:"DrillPairType,omitnil,omitempty" name:"DrillPairType"`

	// 演练对ID列表
	DrillPairIds []*string `json:"DrillPairIds,omitnil,omitempty" name:"DrillPairIds"`
}

type DescribeDrillPairsDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// 要查询演练对的类型，枚举值：DISK（云硬盘）、INSTANCE（云服务器）、CFS（文件存储）
	DrillPairType *string `json:"DrillPairType,omitnil,omitempty" name:"DrillPairType"`

	// 演练对ID列表
	DrillPairIds []*string `json:"DrillPairIds,omitnil,omitempty" name:"DrillPairIds"`
}

func (r *DescribeDrillPairsDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDrillPairsDeniedActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DrillPairType")
	delete(f, "DrillPairIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDrillPairsDeniedActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDrillPairsDeniedActionsResponseParams struct {
	// 演练对操作掩码列表，返回每个演练对被禁止执行的操作
	DrillPairDeniedActionSet []*DrillPairDeniedAction `json:"DrillPairDeniedActionSet,omitnil,omitempty" name:"DrillPairDeniedActionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDrillPairsDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDrillPairsDeniedActionsResponseParams `json:"Response"`
}

func (r *DescribeDrillPairsDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDrillPairsDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDrillPairsRequestParams struct {
	// 要查询演练对的类型。枚举值：DISK / INSTANCE / CFS。
	DrillPairType *string `json:"DrillPairType,omitnil,omitempty" name:"DrillPairType"`

	// 要查询演练对ID列表。
	DrillPairIds []*string `json:"DrillPairIds,omitnil,omitempty" name:"DrillPairIds"`

	// 过滤条件，详见定期快照过滤条件表。
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出结果按升序还是降序。枚举值：ASC / DESC。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 输出结果的排序字段。枚举值：CREATE_TIME / END_TIME。
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeDrillPairsRequest struct {
	*tchttp.BaseRequest
	
	// 要查询演练对的类型。枚举值：DISK / INSTANCE / CFS。
	DrillPairType *string `json:"DrillPairType,omitnil,omitempty" name:"DrillPairType"`

	// 要查询演练对ID列表。
	DrillPairIds []*string `json:"DrillPairIds,omitnil,omitempty" name:"DrillPairIds"`

	// 过滤条件，详见定期快照过滤条件表。
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出结果按升序还是降序。枚举值：ASC / DESC。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 输出结果的排序字段。枚举值：CREATE_TIME / END_TIME。
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeDrillPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDrillPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DrillPairType")
	delete(f, "DrillPairIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDrillPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDrillPairsResponseParams struct {
	// 有效的容灾演练对数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 容灾演练对列表。
	DrillPairSet []*DrillPair `json:"DrillPairSet,omitnil,omitempty" name:"DrillPairSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDrillPairsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDrillPairsResponseParams `json:"Response"`
}

func (r *DescribeDrillPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDrillPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileBackupObjectsRequestParams struct {

}

type DescribeFileBackupObjectsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeFileBackupObjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileBackupObjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileBackupObjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileBackupObjectsResponseParams struct {
	// 当前路径下包含的目录及文件总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFileBackupObjectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileBackupObjectsResponseParams `json:"Response"`
}

func (r *DescribeFileBackupObjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileBackupObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileBackupPlansRequestParams struct {
	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 过滤条件。支持: instance-id, plan-id, plan-name, status, auto-backup-policy-id
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeFileBackupPlansRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 过滤条件。支持: instance-id, plan-id, plan-name, status, auto-backup-policy-id
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeFileBackupPlansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileBackupPlansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Order")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileBackupPlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileBackupPlansResponseParams struct {
	// 符合条件的计划总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合条件的计划详情
	PlanSet []*PlanInfo `json:"PlanSet,omitnil,omitempty" name:"PlanSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFileBackupPlansResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileBackupPlansResponseParams `json:"Response"`
}

func (r *DescribeFileBackupPlansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileBackupPlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileBackupsDeniedActionsRequestParams struct {
	// 要查询的文件备份ID列表
	BackupIds []*string `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`
}

type DescribeFileBackupsDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的文件备份ID列表
	BackupIds []*string `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`
}

func (r *DescribeFileBackupsDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileBackupsDeniedActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileBackupsDeniedActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileBackupsDeniedActionsResponseParams struct {
	// 备份的操作掩码。
	BackupDeniedActionSet []*BackupDeniedAction `json:"BackupDeniedActionSet,omitnil,omitempty" name:"BackupDeniedActionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFileBackupsDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileBackupsDeniedActionsResponseParams `json:"Response"`
}

func (r *DescribeFileBackupsDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileBackupsDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileBackupsRequestParams struct {
	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 过滤条件。支持: backup-id, plan-id, instance-id, status, backup-type, auto-backup-policy-id
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeFileBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 过滤条件。支持: backup-id, plan-id, instance-id, status, backup-type, auto-backup-policy-id
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeFileBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Order")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileBackupsResponseParams struct {
	// 符合条件的备份点总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合条件的备份点详情
	BackupSet []*BackupInfo `json:"BackupSet,omitnil,omitempty" name:"BackupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFileBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileBackupsResponseParams `json:"Response"`
}

func (r *DescribeFileBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileRestoreTasksRequestParams struct {
	// 过滤条件。支持: backup-id, task-id, instance-id, "         "target-instance-id, status
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeFileRestoreTasksRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。支持: backup-id, task-id, instance-id, "         "target-instance-id, status
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeFileRestoreTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileRestoreTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileRestoreTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileRestoreTasksResponseParams struct {
	// 符合条件的总数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 恢复任务列表详情。
	RestoreTaskSet []*RestoreTask `json:"RestoreTaskSet,omitnil,omitempty" name:"RestoreTaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFileRestoreTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileRestoreTasksResponseParams `json:"Response"`
}

func (r *DescribeFileRestoreTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileRestoreTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobsRequestParams struct {

}

type DescribeJobsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobsResponseParams `json:"Response"`
}

func (r *DescribeJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePriceCreateCopyPairsRequestParams struct {
	// 每个复制对的容量列表，长度 1~10。数组长度即为询价的复制对个数，每个元素对应一个复制对的容量
	DataCapacities []*int64 `json:"DataCapacities,omitnil,omitempty" name:"DataCapacities"`
}

type DescribePriceCreateCopyPairsRequest struct {
	*tchttp.BaseRequest
	
	// 每个复制对的容量列表，长度 1~10。数组长度即为询价的复制对个数，每个元素对应一个复制对的容量
	DataCapacities []*int64 `json:"DataCapacities,omitnil,omitempty" name:"DataCapacities"`
}

func (r *DescribePriceCreateCopyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePriceCreateCopyPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataCapacities")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePriceCreateCopyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePriceCreateCopyPairsResponseParams struct {
	// 复制对价格列表，与入参一一对应
	CopyPairPrices []*CopyPairPrice `json:"CopyPairPrices,omitnil,omitempty" name:"CopyPairPrices"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePriceCreateCopyPairsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePriceCreateCopyPairsResponseParams `json:"Response"`
}

func (r *DescribePriceCreateCopyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePriceCreateCopyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProtectGroupsDeniedActionsRequestParams struct {
	// 保护组ID列表
	ProtectGroupIds []*string `json:"ProtectGroupIds,omitnil,omitempty" name:"ProtectGroupIds"`
}

type DescribeProtectGroupsDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// 保护组ID列表
	ProtectGroupIds []*string `json:"ProtectGroupIds,omitnil,omitempty" name:"ProtectGroupIds"`
}

func (r *DescribeProtectGroupsDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProtectGroupsDeniedActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProtectGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProtectGroupsDeniedActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProtectGroupsDeniedActionsResponseParams struct {
	// 保护组操作掩码列表，返回每个保护组被禁止执行的操作
	ProtectGroupDeniedActionSet []*ProtectGroupDeniedAction `json:"ProtectGroupDeniedActionSet,omitnil,omitempty" name:"ProtectGroupDeniedActionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProtectGroupsDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProtectGroupsDeniedActionsResponseParams `json:"Response"`
}

func (r *DescribeProtectGroupsDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProtectGroupsDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProtectedInstancesRequestParams struct {
	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 过滤条件。支持: instance-id, agent-status
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeProtectedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 过滤条件。支持: instance-id, agent-status
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeProtectedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProtectedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Order")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProtectedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProtectedInstancesResponseParams struct {
	// 符合条件的受保护实例总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合条件的受保护实例详情
	InstanceSet []*ProtectInstance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProtectedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProtectedInstancesResponseParams `json:"Response"`
}

func (r *DescribeProtectedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProtectedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupMappingsRequestParams struct {
	// 安全组映射所属的站点对ID。
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 过滤条件，详见过滤条件表。支持的Name：src-security-group-id、target-security-group-id
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为500。关于Limit的更进一步介绍请参考 API 简介中的相关小节
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出结果按升序还是降序，可选值：ASC、DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 输出结果的排序字段，可选值：CREATE_TIME
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeSecurityGroupMappingsRequest struct {
	*tchttp.BaseRequest
	
	// 安全组映射所属的站点对ID。
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 过滤条件，详见过滤条件表。支持的Name：src-security-group-id、target-security-group-id
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为500。关于Limit的更进一步介绍请参考 API 简介中的相关小节
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出结果按升序还是降序，可选值：ASC、DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 输出结果的排序字段，可选值：CREATE_TIME
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeSecurityGroupMappingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupMappingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SitePairId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupMappingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupMappingsResponseParams struct {
	// 符合条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 安全组映射详情。
	SecurityGroupMappingSet []*SecurityGroupMapping `json:"SecurityGroupMappingSet,omitnil,omitempty" name:"SecurityGroupMappingSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupMappingsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupMappingsResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupMappingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupMappingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcMappingsRequestParams struct {
	// 要查询的站点对id
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 过滤条件。支持: source-vpc-id, target-vpc-id, source-subnet-id, target-subnet-id
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeVpcMappingsRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的站点对id
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 过滤条件。支持: source-vpc-id, target-vpc-id, source-subnet-id, target-subnet-id
	Filters []*FilterModel `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeVpcMappingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcMappingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SitePairId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcMappingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcMappingsResponseParams struct {
	// 符合条件的VPC映射规则总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// VPC映射规则列表
	VpcMappingSet []*VpcMapping `json:"VpcMappingSet,omitnil,omitempty" name:"VpcMappingSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpcMappingsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcMappingsResponseParams `json:"Response"`
}

func (r *DescribeVpcMappingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcMappingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisasterRecoveryDrillGroup struct {
	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 账户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountUin *string `json:"AccountUin,omitnil,omitempty" name:"AccountUin"`

	// 子账户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitnil,omitempty" name:"SubAccountUin"`

	// 容灾站点对ID
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 保护组ID
	ProtectGroupId *string `json:"ProtectGroupId,omitnil,omitempty" name:"ProtectGroupId"`

	// 演练组ID
	DrillGroupId *string `json:"DrillGroupId,omitnil,omitempty" name:"DrillGroupId"`

	// 演练组名称
	DrillGroupName *string `json:"DrillGroupName,omitnil,omitempty" name:"DrillGroupName"`

	// 演练组类型。枚举值：DISK / INSTANCE / CFS。
	DrillGroupType *string `json:"DrillGroupType,omitnil,omitempty" name:"DrillGroupType"`

	// 恢复时间点
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecoveryTime *string `json:"RecoveryTime,omitnil,omitempty" name:"RecoveryTime"`

	// 演练VPC
	DrillVpc *string `json:"DrillVpc,omitnil,omitempty" name:"DrillVpc"`

	// 演练安全组
	// 注意：此字段可能返回 null，表示取不到有效值。
	DrillSecurityGroup *string `json:"DrillSecurityGroup,omitnil,omitempty" name:"DrillSecurityGroup"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 生命周期状态。枚举值：NORMAL / DELETED。
	LifeState *string `json:"LifeState,omitnil,omitempty" name:"LifeState"`

	// 容灾类型。枚举值：CROSS_ZONE / CROSS_REGION 等。
	DisasterRecoveryType *string `json:"DisasterRecoveryType,omitnil,omitempty" name:"DisasterRecoveryType"`

	// 复制技术。枚举值：SYN（同步）/ ASYN（异步）。
	CopyType *string `json:"CopyType,omitnil,omitempty" name:"CopyType"`

	// 对端云名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeerCloudName *string `json:"PeerCloudName,omitnil,omitempty" name:"PeerCloudName"`

	// 本地云名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalCloudName *string `json:"LocalCloudName,omitnil,omitempty" name:"LocalCloudName"`

	// 生产地域
	SourceRegion *string `json:"SourceRegion,omitnil,omitempty" name:"SourceRegion"`

	// 生产可用区
	SourceZone *string `json:"SourceZone,omitnil,omitempty" name:"SourceZone"`

	// 生产端VPC
	SourceVpc *string `json:"SourceVpc,omitnil,omitempty" name:"SourceVpc"`

	// 演练地域
	DrillRegion *string `json:"DrillRegion,omitnil,omitempty" name:"DrillRegion"`

	// 演练可用区
	DrillZone *string `json:"DrillZone,omitnil,omitempty" name:"DrillZone"`

	// 数据方向。枚举值：POSITIVE（正向）/ REVERSE（反向）。
	DataDirection *string `json:"DataDirection,omitnil,omitempty" name:"DataDirection"`

	// 绑定的演练资源数量。
	BindDrilledResourceCount *int64 `json:"BindDrilledResourceCount,omitnil,omitempty" name:"BindDrilledResourceCount"`

	// 演练资源状态分布（key 为状态名如 FAILED / SUCCESS，value 为该状态数量）。
	DrilledResourceStatusSet []*DrilledResourceStatus `json:"DrilledResourceStatusSet,omitnil,omitempty" name:"DrilledResourceStatusSet"`
}

type DisasterRecoveryOverview struct {
	// 地域 ID
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 站点对总数
	SitePairCount *int64 `json:"SitePairCount,omitnil,omitempty" name:"SitePairCount"`

	// 跨地域站点对数
	SitePairCrossRegionCount *int64 `json:"SitePairCrossRegionCount,omitnil,omitempty" name:"SitePairCrossRegionCount"`

	// 跨可用区站点对数
	SitePairCrossZoneCount *int64 `json:"SitePairCrossZoneCount,omitnil,omitempty" name:"SitePairCrossZoneCount"`

	// 跨云站点对数
	SitePairCrossCloudCount *int64 `json:"SitePairCrossCloudCount,omitnil,omitempty" name:"SitePairCrossCloudCount"`

	// 保护组总数
	ProtectGroupCount *int64 `json:"ProtectGroupCount,omitnil,omitempty" name:"ProtectGroupCount"`

	// 跨地域保护组数
	ProtectGroupCrossRegionCount *int64 `json:"ProtectGroupCrossRegionCount,omitnil,omitempty" name:"ProtectGroupCrossRegionCount"`

	// 跨可用区保护组数
	ProtectGroupCrossZoneCount *int64 `json:"ProtectGroupCrossZoneCount,omitnil,omitempty" name:"ProtectGroupCrossZoneCount"`

	// 跨云保护组数
	ProtectGroupCrossCloudCount *int64 `json:"ProtectGroupCrossCloudCount,omitnil,omitempty" name:"ProtectGroupCrossCloudCount"`

	// 复制对总数
	CopyPairCount *int64 `json:"CopyPairCount,omitnil,omitempty" name:"CopyPairCount"`

	// RPO 正常的复制对数
	CopyPairSuccessRPOCount *int64 `json:"CopyPairSuccessRPOCount,omitnil,omitempty" name:"CopyPairSuccessRPOCount"`

	// RPO 异常的复制对数
	CopyPairErrorRPOCount *int64 `json:"CopyPairErrorRPOCount,omitnil,omitempty" name:"CopyPairErrorRPOCount"`

	// 演练对总数
	DrillPairCount *int64 `json:"DrillPairCount,omitnil,omitempty" name:"DrillPairCount"`

	// 演练中
	DrillPairDrillingCount *int64 `json:"DrillPairDrillingCount,omitnil,omitempty" name:"DrillPairDrillingCount"`

	// 演练失败
	DrillPairFailedCount *int64 `json:"DrillPairFailedCount,omitnil,omitempty" name:"DrillPairFailedCount"`

	// 演练成功
	DrillPairSuccessCount *int64 `json:"DrillPairSuccessCount,omitnil,omitempty" name:"DrillPairSuccessCount"`

	// 受保护资源总数
	ProtectedResourceCount *int64 `json:"ProtectedResourceCount,omitnil,omitempty" name:"ProtectedResourceCount"`

	// 受保护资源-复制中
	ProtectedResourceCopyingCount *int64 `json:"ProtectedResourceCopyingCount,omitnil,omitempty" name:"ProtectedResourceCopyingCount"`

	// 受保护资源-已停止/初始化
	ProtectedResourceStoppedCount *int64 `json:"ProtectedResourceStoppedCount,omitnil,omitempty" name:"ProtectedResourceStoppedCount"`

	// 切换失败
	FailoverFailedCount *int64 `json:"FailoverFailedCount,omitnil,omitempty" name:"FailoverFailedCount"`
}

type DiskCopyPairForCvm struct {
	// 云硬盘复制对ID
	CopyPairId *string `json:"CopyPairId,omitnil,omitempty" name:"CopyPairId"`

	// 云硬盘复制对名称
	CopyPairName *string `json:"CopyPairName,omitnil,omitempty" name:"CopyPairName"`

	// 生产端云硬盘ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceResourceId *string `json:"SourceResourceId,omitnil,omitempty" name:"SourceResourceId"`

	// 容灾端云硬盘ID（延迟创建模式且 CVM 未真实创建时被脱敏为空字符串）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetResourceId *string `json:"TargetResourceId,omitnil,omitempty" name:"TargetResourceId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type DiskInfo struct {
	// 云硬盘ID
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 云盘的镜像格式。QCOW2:  qcow2格式，这种格式的云盘不能用于容灾；RAW：raw格式，可以用于容灾。
	ImageFormat *string `json:"ImageFormat,omitnil,omitempty" name:"ImageFormat"`
}

type DiskModel struct {
	// 云盘类型
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 云盘大小（单位GB，范围 (0, 32000]）
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 是否随实例删除（仅 DataDisks 元素能传）
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitnil,omitempty" name:"DeleteWithInstance"`
}

type DrillPair struct {
	// 用户ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 演练对ID
	DrillPairId *string `json:"DrillPairId,omitnil,omitempty" name:"DrillPairId"`

	// 演练对名称
	DrillPairName *string `json:"DrillPairName,omitnil,omitempty" name:"DrillPairName"`

	// 演练对状态。枚举值：RUNNING / SUCCESS / FAILED 等。
	DrillPairState *string `json:"DrillPairState,omitnil,omitempty" name:"DrillPairState"`

	// 容灾站点对ID
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 云硬盘复制对ID
	CopyPairId *string `json:"CopyPairId,omitnil,omitempty" name:"CopyPairId"`

	// 生产地域
	SourceRegion *string `json:"SourceRegion,omitnil,omitempty" name:"SourceRegion"`

	// 生产可用区
	SourceZone *string `json:"SourceZone,omitnil,omitempty" name:"SourceZone"`

	// 容灾地域
	TargetRegion *string `json:"TargetRegion,omitnil,omitempty" name:"TargetRegion"`

	// 容灾可用区
	TargetZone *string `json:"TargetZone,omitnil,omitempty" name:"TargetZone"`

	// 生产站点盘ID
	SourceResourceId *string `json:"SourceResourceId,omitnil,omitempty" name:"SourceResourceId"`

	// 演练资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetResourceId *string `json:"TargetResourceId,omitnil,omitempty" name:"TargetResourceId"`

	// 演练对的类型。枚举值：DISK / INSTANCE / CFS。
	DrillPairType *string `json:"DrillPairType,omitnil,omitempty" name:"DrillPairType"`

	// 演练资源容量（GB）。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 演练的容灾点
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecoveryTime *string `json:"RecoveryTime,omitnil,omitempty" name:"RecoveryTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 演练结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 是否正在回滚。0 - 未回滚，1 - 回滚中。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rollbacking *int64 `json:"Rollbacking,omitnil,omitempty" name:"Rollbacking"`

	// 回滚进度百分比（0-100）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RollbackPercent *int64 `json:"RollbackPercent,omitnil,omitempty" name:"RollbackPercent"`

	// 创建定期备份策略的账户uin ID信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountUin *string `json:"AccountUin,omitnil,omitempty" name:"AccountUin"`

	// 创建定期备份策略的子账户uin ID信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitnil,omitempty" name:"SubAccountUin"`

	// 保护组ID
	ProtectGroupId *string `json:"ProtectGroupId,omitnil,omitempty" name:"ProtectGroupId"`

	// 演练组ID
	DrillGroupId *string `json:"DrillGroupId,omitnil,omitempty" name:"DrillGroupId"`

	// 复制对名称。
	CopyPairName *string `json:"CopyPairName,omitnil,omitempty" name:"CopyPairName"`

	// 演练组名称。
	DrillGroupName *string `json:"DrillGroupName,omitnil,omitempty" name:"DrillGroupName"`
}

type DrillPairDeniedAction struct {
	// 演练对ID
	DrillPairId *string `json:"DrillPairId,omitnil,omitempty" name:"DrillPairId"`

	// 被禁止的操作列表（Action名称数组）
	DeniedActions []*DeniedAction `json:"DeniedActions,omitnil,omitempty" name:"DeniedActions"`
}

type DrilledResourceStatus struct {
	// 演练组关联的演练资源的状态
	ResourceStatus *string `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`

	// 演练组关联演练资源处于某个状态的数量
	ResourceCount *uint64 `json:"ResourceCount,omitnil,omitempty" name:"ResourceCount"`
}

type EnhancedService struct {
	// 开启云安全服务。若不指定该参数，则默认开启云安全服务。
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitnil,omitempty" name:"SecurityService"`

	// 开启云监控服务。若不指定该参数，则默认开启云监控服务。
	MonitorService *RunSecurityServiceEnabled `json:"MonitorService,omitnil,omitempty" name:"MonitorService"`

	// 安装 tat-agent。若不指定该参数，则默认逻辑与 CVM 控制台一致：境外地域不安装、境内非 GPU 机型默认安装、境内 GPU 机型默认不安装。
	AutomationService *AutomationServiceEnabled `json:"AutomationService,omitnil,omitempty" name:"AutomationService"`

	// 开启基础服务。
	BasicService *BasicServicesSettings `json:"BasicService,omitnil,omitempty" name:"BasicService"`
}

type FileBackupOverview struct {
	// 整机备份点总数
	BackupCount *int64 `json:"BackupCount,omitnil,omitempty" name:"BackupCount"`

	// 创建中数量
	CreatingBackupCount *int64 `json:"CreatingBackupCount,omitnil,omitempty" name:"CreatingBackupCount"`

	// 失败数量
	FailedBackupCount *int64 `json:"FailedBackupCount,omitnil,omitempty" name:"FailedBackupCount"`

	// 已完成数量
	SuccessBackupCount *int64 `json:"SuccessBackupCount,omitnil,omitempty" name:"SuccessBackupCount"`

	// 恢复中的总数量
	RestoringBackupCount *int64 `json:"RestoringBackupCount,omitnil,omitempty" name:"RestoringBackupCount"`

	// 整机备份总容量
	BackupSizeMb *int64 `json:"BackupSizeMb,omitnil,omitempty" name:"BackupSizeMb"`

	// 受保护 CVM 资源数
	BackupResourceCount *int64 `json:"BackupResourceCount,omitnil,omitempty" name:"BackupResourceCount"`
}

type FilterModel struct {
	// 过滤器名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤器值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type FinishFailoverCopyPairsRequestParams struct {
	// <p>复制对ID列表。长度范围 [1, 50]。当 CopyPairType=INSTANCE 时传 CVM 复制对ID，否则传云盘/CFS 复制对ID。</p>
	CopyPairIds []*string `json:"CopyPairIds,omitnil,omitempty" name:"CopyPairIds"`

	// <p>要完成切换的复制对类型。枚举值：DISK / INSTANCE / CFS。</p>
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`
}

type FinishFailoverCopyPairsRequest struct {
	*tchttp.BaseRequest
	
	// <p>复制对ID列表。长度范围 [1, 50]。当 CopyPairType=INSTANCE 时传 CVM 复制对ID，否则传云盘/CFS 复制对ID。</p>
	CopyPairIds []*string `json:"CopyPairIds,omitnil,omitempty" name:"CopyPairIds"`

	// <p>要完成切换的复制对类型。枚举值：DISK / INSTANCE / CFS。</p>
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`
}

func (r *FinishFailoverCopyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FinishFailoverCopyPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CopyPairIds")
	delete(f, "CopyPairType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FinishFailoverCopyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FinishFailoverCopyPairsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FinishFailoverCopyPairsResponse struct {
	*tchttp.BaseResponse
	Response *FinishFailoverCopyPairsResponseParams `json:"Response"`
}

func (r *FinishFailoverCopyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FinishFailoverCopyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlowControlRule struct {
	// 流控开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 流控结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 流控规则最大带宽，单位MB/s
	MaxBandwidthMBps *uint64 `json:"MaxBandwidthMBps,omitnil,omitempty" name:"MaxBandwidthMBps"`
}

type InstanceBackupOverview struct {
	// 整机备份点总数
	BackupCount *int64 `json:"BackupCount,omitnil,omitempty" name:"BackupCount"`

	// 创建中数量
	CreatingBackupCount *int64 `json:"CreatingBackupCount,omitnil,omitempty" name:"CreatingBackupCount"`

	// 失败数量
	FailedBackupCount *int64 `json:"FailedBackupCount,omitnil,omitempty" name:"FailedBackupCount"`

	// 已完成数量
	SuccessBackupCount *int64 `json:"SuccessBackupCount,omitnil,omitempty" name:"SuccessBackupCount"`

	// 恢复中的总数量
	RestoringBackupCount *int64 `json:"RestoringBackupCount,omitnil,omitempty" name:"RestoringBackupCount"`

	// 整机备份总容量
	BackupSizeMb *int64 `json:"BackupSizeMb,omitnil,omitempty" name:"BackupSizeMb"`

	// 受保护 CVM 资源数
	BackupResourceCount *int64 `json:"BackupResourceCount,omitnil,omitempty" name:"BackupResourceCount"`
}

type InstanceChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 自动续费标识。取值范围：NOTIFY_AND_AUTO_RENEW（通知过期且自动续费）、NOTIFY_AND_MANUAL_RENEW（通知过期不自动续费）、DISABLE_NOTIFY_AND_MANUAL_RENEW（不通知过期不自动续费）。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type InternetAccessible struct {
	// 网络计费类型。取值范围：BANDWIDTH_PREPAID（预付费按带宽结算）、TRAFFIC_POSTPAID_BY_HOUR（流量按小时后付费）、BANDWIDTH_POSTPAID_BY_HOUR（带宽按小时后付费）、BANDWIDTH_PACKAGE（带宽包用户）。默认取值：非带宽包用户默认与子机付费类型保持一致。
	InternetChargeType *string `json:"InternetChargeType,omitnil,omitempty" name:"InternetChargeType"`

	// 公网出带宽上限，单位：Mbps。默认值：0Mbps。不同机型带宽上限范围不一致，具体限制详见购买网络带宽。
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 是否分配公网IP。取值范围：true（表示分配公网IP）/false（表示不分配公网IP）。当公网带宽大于0Mbps时，可自由选择开通与否，默认开通公网IP；当公网带宽为0，则不允许分配公网IP。该参数仅在 RunInstances 接口中作为入参使用。
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitnil,omitempty" name:"PublicIpAssigned"`

	// 网络模式：移动:"CMCC"、电信:"CTCC"、联通:"CUCC"。
	InternetServiceProvider *string `json:"InternetServiceProvider,omitnil,omitempty" name:"InternetServiceProvider"`
}

type LoginSettings struct {
	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：Linux 实例密码必须 8-30 位，推荐使用 12 位以上密码，不能以"/"开头，至少包含以下字符中的三种不同字符，字符种类：小写字母 a-z、大写字母 A-Z、数字 0-9、特殊字符 ()`~!@#$%^&*-+=_|{}[]:;'<>,.?/。Windows 实例密码必须 12-30 位，不能以"/"开头且不包括用户名，至少包含以下字符中的三种不同字符，字符种类：小写字母 a-z、大写字母 A-Z、数字 0-9、特殊字符 ()`~!@#$%^&*-+=_|{}[]:;' <>,.?/。若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口 [DescribeKeyPairs](https://cloud.tencent.com/document/api/213/15699) 获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 保持镜像的原始设置。该参数与 Password 或 KeyIds.N 不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为 TRUE。取值范围：TRUE（表示保持镜像的登录设置）/FALSE（表示不保持镜像的登录设置）。默认取值：FALSE。
	KeepImageLogin *string `json:"KeepImageLogin,omitnil,omitempty" name:"KeepImageLogin"`
}

// Predefined struct for user
type ModifyAutoBackupPolicyAttributeRequestParams struct {
	// 备份策略id
	AutoBackupPolicyId *string `json:"AutoBackupPolicyId,omitnil,omitempty" name:"AutoBackupPolicyId"`

	// 定期备份的执行策略。
	Policy []*Policy `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 通过该定期备份策略创建的备份是否永久保留。false表示非永久保留，true表示永久保留，默认为false。
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 定期备份策略的名称。
	AutoBackupPolicyName *string `json:"AutoBackupPolicyName,omitnil,omitempty" name:"AutoBackupPolicyName"`

	// 是否激活定期备份策略。
	IsActivated *bool `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// 通过定期备份策略创建出的备份保留时间。
	RetentionDays *uint64 `json:"RetentionDays,omitnil,omitempty" name:"RetentionDays"`

	// 该定期备份策略创建的备份可以保留的月数，该参数不可与IsPermanent/RetentionDays参数冲突。
	RetentionMonths *uint64 `json:"RetentionMonths,omitnil,omitempty" name:"RetentionMonths"`

	// 通过该定期备份策略最多保留的备份个数，超过该个数限制后自动删除最先创建的备份，该参数不可与IsPermanent参数冲突。
	RetentionAmount *uint64 `json:"RetentionAmount,omitnil,omitempty" name:"RetentionAmount"`

	// 备份存储类型。SNAPSHOT表示走快照（不需要备份库），VAULT表示走备份库（必须关联一个备份库）。默认为SNAPSHOT
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 备份库ID，创建agent备份策略时必须指定。当StorageType为VAULT时必传。
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// 定期备份高级保留策略，该参数不可与IsPermanent参数冲突。
	AdvancedRetentionPolicy *AdvancedRetentionPolicy `json:"AdvancedRetentionPolicy,omitnil,omitempty" name:"AdvancedRetentionPolicy"`
}

type ModifyAutoBackupPolicyAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 备份策略id
	AutoBackupPolicyId *string `json:"AutoBackupPolicyId,omitnil,omitempty" name:"AutoBackupPolicyId"`

	// 定期备份的执行策略。
	Policy []*Policy `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 通过该定期备份策略创建的备份是否永久保留。false表示非永久保留，true表示永久保留，默认为false。
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 定期备份策略的名称。
	AutoBackupPolicyName *string `json:"AutoBackupPolicyName,omitnil,omitempty" name:"AutoBackupPolicyName"`

	// 是否激活定期备份策略。
	IsActivated *bool `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// 通过定期备份策略创建出的备份保留时间。
	RetentionDays *uint64 `json:"RetentionDays,omitnil,omitempty" name:"RetentionDays"`

	// 该定期备份策略创建的备份可以保留的月数，该参数不可与IsPermanent/RetentionDays参数冲突。
	RetentionMonths *uint64 `json:"RetentionMonths,omitnil,omitempty" name:"RetentionMonths"`

	// 通过该定期备份策略最多保留的备份个数，超过该个数限制后自动删除最先创建的备份，该参数不可与IsPermanent参数冲突。
	RetentionAmount *uint64 `json:"RetentionAmount,omitnil,omitempty" name:"RetentionAmount"`

	// 备份存储类型。SNAPSHOT表示走快照（不需要备份库），VAULT表示走备份库（必须关联一个备份库）。默认为SNAPSHOT
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 备份库ID，创建agent备份策略时必须指定。当StorageType为VAULT时必传。
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// 定期备份高级保留策略，该参数不可与IsPermanent参数冲突。
	AdvancedRetentionPolicy *AdvancedRetentionPolicy `json:"AdvancedRetentionPolicy,omitnil,omitempty" name:"AdvancedRetentionPolicy"`
}

func (r *ModifyAutoBackupPolicyAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoBackupPolicyAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoBackupPolicyId")
	delete(f, "Policy")
	delete(f, "IsPermanent")
	delete(f, "AutoBackupPolicyName")
	delete(f, "IsActivated")
	delete(f, "RetentionDays")
	delete(f, "RetentionMonths")
	delete(f, "RetentionAmount")
	delete(f, "StorageType")
	delete(f, "VaultId")
	delete(f, "AdvancedRetentionPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoBackupPolicyAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoBackupPolicyAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAutoBackupPolicyAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoBackupPolicyAttributeResponseParams `json:"Response"`
}

func (r *ModifyAutoBackupPolicyAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoBackupPolicyAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupAttributeRequestParams struct {
	// 备份ID。该字段的取值取决于ResourceType：当ResourceType=CVM（默认）时，需传入备份组ID（BackupGroupId），可通过DescribeBackupGroups（查询备份组列表）查询
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 备份的名称。
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// 是否为永久保留的备份。
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 备份到期时间。
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`
}

type ModifyBackupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 备份ID。该字段的取值取决于ResourceType：当ResourceType=CVM（默认）时，需传入备份组ID（BackupGroupId），可通过DescribeBackupGroups（查询备份组列表）查询
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 备份的名称。
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// 是否为永久保留的备份。
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 备份到期时间。
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`
}

func (r *ModifyBackupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupId")
	delete(f, "BackupName")
	delete(f, "IsPermanent")
	delete(f, "Deadline")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBackupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupAttributeResponseParams `json:"Response"`
}

func (r *ModifyBackupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupVaultAttributeRequestParams struct {
	// 备份库ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// 备份库名称
	VaultName *string `json:"VaultName,omitnil,omitempty" name:"VaultName"`

	// 备份库描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyBackupVaultAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 备份库ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// 备份库名称
	VaultName *string `json:"VaultName,omitnil,omitempty" name:"VaultName"`

	// 备份库描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyBackupVaultAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupVaultAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VaultId")
	delete(f, "VaultName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupVaultAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupVaultAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBackupVaultAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupVaultAttributeResponseParams `json:"Response"`
}

func (r *ModifyBackupVaultAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupVaultAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCopyPairAttributeRequestParams struct {
	// 要修改属性的复制对id
	CopyPairId *string `json:"CopyPairId,omitnil,omitempty" name:"CopyPairId"`

	// 要修改的复制对类型，可选值：DISK、INSTANCE、CFS，默认 INSTANCE
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`

	// 修改复制对名称（长度最大支持 64 个字符）
	CopyPairName *string `json:"CopyPairName,omitnil,omitempty" name:"CopyPairName"`
}

type ModifyCopyPairAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 要修改属性的复制对id
	CopyPairId *string `json:"CopyPairId,omitnil,omitempty" name:"CopyPairId"`

	// 要修改的复制对类型，可选值：DISK、INSTANCE、CFS，默认 INSTANCE
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`

	// 修改复制对名称（长度最大支持 64 个字符）
	CopyPairName *string `json:"CopyPairName,omitnil,omitempty" name:"CopyPairName"`
}

func (r *ModifyCopyPairAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCopyPairAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CopyPairId")
	delete(f, "CopyPairType")
	delete(f, "CopyPairName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCopyPairAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCopyPairAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCopyPairAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCopyPairAttributeResponseParams `json:"Response"`
}

func (r *ModifyCopyPairAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCopyPairAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDrillGroupAttributeRequestParams struct {
	// 要修改属性的容灾演练组id。
	DrillGroupId *string `json:"DrillGroupId,omitnil,omitempty" name:"DrillGroupId"`

	// 修改容灾演练组名称（长度最大支持 64 个字符）
	DrillGroupName *string `json:"DrillGroupName,omitnil,omitempty" name:"DrillGroupName"`
}

type ModifyDrillGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 要修改属性的容灾演练组id。
	DrillGroupId *string `json:"DrillGroupId,omitnil,omitempty" name:"DrillGroupId"`

	// 修改容灾演练组名称（长度最大支持 64 个字符）
	DrillGroupName *string `json:"DrillGroupName,omitnil,omitempty" name:"DrillGroupName"`
}

func (r *ModifyDrillGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDrillGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DrillGroupId")
	delete(f, "DrillGroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDrillGroupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDrillGroupAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDrillGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDrillGroupAttributeResponseParams `json:"Response"`
}

func (r *ModifyDrillGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDrillGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDrillPairAttributeRequestParams struct {
	// 要修改属性的容灾演练对id
	DrillPairId *string `json:"DrillPairId,omitnil,omitempty" name:"DrillPairId"`

	// 修改容灾演练对名称（长度最大支持 64 个字符）
	DrillPairName *string `json:"DrillPairName,omitnil,omitempty" name:"DrillPairName"`
}

type ModifyDrillPairAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 要修改属性的容灾演练对id
	DrillPairId *string `json:"DrillPairId,omitnil,omitempty" name:"DrillPairId"`

	// 修改容灾演练对名称（长度最大支持 64 个字符）
	DrillPairName *string `json:"DrillPairName,omitnil,omitempty" name:"DrillPairName"`
}

func (r *ModifyDrillPairAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDrillPairAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DrillPairId")
	delete(f, "DrillPairName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDrillPairAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDrillPairAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDrillPairAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDrillPairAttributeResponseParams `json:"Response"`
}

func (r *ModifyDrillPairAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDrillPairAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFileBackupAttributeRequestParams struct {
	// 备份ID
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 备份的名称。
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// 是否为永久保留的备份。
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 备份到期时间。
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`
}

type ModifyFileBackupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 备份ID
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 备份的名称。
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// 是否为永久保留的备份。
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// 备份到期时间。
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`
}

func (r *ModifyFileBackupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFileBackupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupId")
	delete(f, "BackupName")
	delete(f, "IsPermanent")
	delete(f, "Deadline")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFileBackupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFileBackupAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyFileBackupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFileBackupAttributeResponseParams `json:"Response"`
}

func (r *ModifyFileBackupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFileBackupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFileBackupPlanRequestParams struct {
	// 备份计划ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 备份策略ID
	//
	// Deprecated: PolicyId is deprecated.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 计划名称
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// 备份路径列表，1~20 个
	BackupPaths []*string `json:"BackupPaths,omitnil,omitempty" name:"BackupPaths"`

	// 包含文件类型，0~20 个
	IncludeFileTypes []*string `json:"IncludeFileTypes,omitnil,omitempty" name:"IncludeFileTypes"`

	// 排除文件路径列表，0~20 个
	ExcludePatterns []*string `json:"ExcludePatterns,omitnil,omitempty" name:"ExcludePatterns"`

	// 是否排除系统目录
	ExcludeSystemDirectories *bool `json:"ExcludeSystemDirectories,omitnil,omitempty" name:"ExcludeSystemDirectories"`

	// 备份库ID
	//
	// Deprecated: BackupStorageId is deprecated.
	BackupStorageId *string `json:"BackupStorageId,omitnil,omitempty" name:"BackupStorageId"`

	// 计划状态，可选值：normal（正常）、paused（暂停）
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyFileBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// 备份计划ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 备份策略ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 计划名称
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// 备份路径列表，1~20 个
	BackupPaths []*string `json:"BackupPaths,omitnil,omitempty" name:"BackupPaths"`

	// 包含文件类型，0~20 个
	IncludeFileTypes []*string `json:"IncludeFileTypes,omitnil,omitempty" name:"IncludeFileTypes"`

	// 排除文件路径列表，0~20 个
	ExcludePatterns []*string `json:"ExcludePatterns,omitnil,omitempty" name:"ExcludePatterns"`

	// 是否排除系统目录
	ExcludeSystemDirectories *bool `json:"ExcludeSystemDirectories,omitnil,omitempty" name:"ExcludeSystemDirectories"`

	// 备份库ID
	BackupStorageId *string `json:"BackupStorageId,omitnil,omitempty" name:"BackupStorageId"`

	// 计划状态，可选值：normal（正常）、paused（暂停）
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyFileBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFileBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "PolicyId")
	delete(f, "PlanName")
	delete(f, "BackupPaths")
	delete(f, "IncludeFileTypes")
	delete(f, "ExcludePatterns")
	delete(f, "ExcludeSystemDirectories")
	delete(f, "BackupStorageId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFileBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFileBackupPlanResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyFileBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFileBackupPlanResponseParams `json:"Response"`
}

func (r *ModifyFileBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFileBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProtectGroupAttributeRequestParams struct {
	// 要修改属性的保护组id
	ProtectGroupId *string `json:"ProtectGroupId,omitnil,omitempty" name:"ProtectGroupId"`

	// 保护组名称
	ProtectGroupName *string `json:"ProtectGroupName,omitnil,omitempty" name:"ProtectGroupName"`
}

type ModifyProtectGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 要修改属性的保护组id
	ProtectGroupId *string `json:"ProtectGroupId,omitnil,omitempty" name:"ProtectGroupId"`

	// 保护组名称
	ProtectGroupName *string `json:"ProtectGroupName,omitnil,omitempty" name:"ProtectGroupName"`
}

func (r *ModifyProtectGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProtectGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProtectGroupId")
	delete(f, "ProtectGroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProtectGroupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProtectGroupAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProtectGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProtectGroupAttributeResponseParams `json:"Response"`
}

func (r *ModifyProtectGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProtectGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySitePairAttributeRequestParams struct {
	// 要修改属性的容灾站点id
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 容灾站点名称
	SitePairName *string `json:"SitePairName,omitnil,omitempty" name:"SitePairName"`
}

type ModifySitePairAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 要修改属性的容灾站点id
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 容灾站点名称
	SitePairName *string `json:"SitePairName,omitnil,omitempty" name:"SitePairName"`
}

func (r *ModifySitePairAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySitePairAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SitePairId")
	delete(f, "SitePairName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySitePairAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySitePairAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySitePairAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifySitePairAttributeResponseParams `json:"Response"`
}

func (r *ModifySitePairAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySitePairAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Placement struct {
	// 实例所属的可用区 ID。该参数也可以通过调用 [DescribeZones]的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例所属项目ID。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例所属的专用宿主机ID列表。如果您有购买专用宿主机并且指定了该参数，则您购买的实例就会随机的部署在这些专用宿主机上。仅用于出参，当前暂不支持。
	HostId *string `json:"HostId,omitnil,omitempty" name:"HostId"`

	// 实例所属的专用宿主机ID列表，仅用于入参。如果您有购买专用宿主机并且指定了该参数，则您购买的实例就会随机的部署在这些专用宿主机上。
	HostIds []*string `json:"HostIds,omitnil,omitempty" name:"HostIds"`

	// 实例所属项目名。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`
}

type PlanInfo struct {
	// 备份计划ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 计划关联的实例ID
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 计划名称
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// 备份路径列表，1~20 个
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupPaths []*string `json:"BackupPaths,omitnil,omitempty" name:"BackupPaths"`

	// 包含文件类型，0~20 个
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncludeFileTypes []*string `json:"IncludeFileTypes,omitnil,omitempty" name:"IncludeFileTypes"`

	// 排除文件路径列表，0~20 个
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludePatterns []*string `json:"ExcludePatterns,omitnil,omitempty" name:"ExcludePatterns"`

	// 是否排除系统目录
	ExcludeSystemDirectories *bool `json:"ExcludeSystemDirectories,omitnil,omitempty" name:"ExcludeSystemDirectories"`

	// 备份库ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// 备份计划状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 策略ID
	AspId *string `json:"AspId,omitnil,omitempty" name:"AspId"`

	// 策略名称
	AspName *string `json:"AspName,omitnil,omitempty" name:"AspName"`

	// 策略详情
	AspPolicy *AspInfo `json:"AspPolicy,omitnil,omitempty" name:"AspPolicy"`

	// 最近一次执行时间
	LastExecuteTime *string `json:"LastExecuteTime,omitnil,omitempty" name:"LastExecuteTime"`

	// 下次触发时间
	NextTriggerTime *string `json:"NextTriggerTime,omitnil,omitempty" name:"NextTriggerTime"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 最近一次执行错误信息
	LastTriggerError *string `json:"LastTriggerError,omitnil,omitempty" name:"LastTriggerError"`

	// 备份数量
	BackupCount *int64 `json:"BackupCount,omitnil,omitempty" name:"BackupCount"`

	// 流控信息
	FlowControlSettings []*FlowControlRule `json:"FlowControlSettings,omitnil,omitempty" name:"FlowControlSettings"`
}

type Policy struct {
	// 选定周一到周日中需要创建备份的日期，取值范围：[0, 6]。0表示周日触发，1表示周一触发，依次类推。
	DayOfWeek []*uint64 `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 指定定期备份策略的触发时间。单位为小时，取值范围：[0, 23]。00:00 ~ 23:00 共 24 个时间点可选，1表示 01:00，依此类推。
	Hour []*uint64 `json:"Hour,omitnil,omitempty" name:"Hour"`

	// 指定每月从月初到月底需要触发定期备份的日期,取值范围：[1, 31]，1-31分别表示每月的具体日期，比如5表示每月的5号。注：若设置29、30、31等部分月份不存在的日期，则对应不存在日期的月份会跳过不打定期备份。
	DayOfMonth []*uint64 `json:"DayOfMonth,omitnil,omitempty" name:"DayOfMonth"`

	// 指定创建定期备份的间隔天数，取值范围：[1, 365]，例如设置为5，则间隔5天即触发定期备份创建。注：当选择按天备份时，理论上第一次备份的时间为备份策略创建当天。如果当天备份策略创建的时间已经晚于设置的备份时间，那么将会等到第二个备份周期再进行第一次备份。
	IntervalDays *uint64 `json:"IntervalDays,omitnil,omitempty" name:"IntervalDays"`
}

type ProtectGroup struct {
	// 用户AppId
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 保护组ID
	ProtectGroupId *string `json:"ProtectGroupId,omitnil,omitempty" name:"ProtectGroupId"`

	// 保护组名称
	ProtectGroupName *string `json:"ProtectGroupName,omitnil,omitempty" name:"ProtectGroupName"`

	// 保护组类型（产品类型，如 DISK/CFS/INSTANCE）
	ProtectGroupType *string `json:"ProtectGroupType,omitnil,omitempty" name:"ProtectGroupType"`

	// 所属容灾策略ID
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 所属容灾策略名称
	SitePairName *string `json:"SitePairName,omitnil,omitempty" name:"SitePairName"`

	// RPO时间（单位秒）
	RecoveryPointObjective *int64 `json:"RecoveryPointObjective,omitnil,omitempty" name:"RecoveryPointObjective"`

	// 生产地域（当 DataDirection=REVERSE 时会与 TargetRegion 自动轮转，保持用户视角一致）
	SourceRegion *string `json:"SourceRegion,omitnil,omitempty" name:"SourceRegion"`

	// 生产可用区（REVERSE 时与 TargetZone 自动轮转）
	SourceZone *string `json:"SourceZone,omitnil,omitempty" name:"SourceZone"`

	// 生产端VPC（REVERSE 时与 TargetVpc 自动轮转）
	SourceVpc *string `json:"SourceVpc,omitnil,omitempty" name:"SourceVpc"`

	// 容灾地域（REVERSE 时与 SourceRegion 自动轮转）
	TargetRegion *string `json:"TargetRegion,omitnil,omitempty" name:"TargetRegion"`

	// 容灾可用区
	TargetZone *string `json:"TargetZone,omitnil,omitempty" name:"TargetZone"`

	// 容灾端VPC
	TargetVpc *string `json:"TargetVpc,omitnil,omitempty" name:"TargetVpc"`

	// 复制技术（SYN 同步 / ASY 异步）
	CopyType *string `json:"CopyType,omitnil,omitempty" name:"CopyType"`

	// 容灾类型（CROSS_ZONE 跨可用区 / CROSS_REGION 跨地域 / CROSS_CLOUD 跨云）
	DisasterRecoveryType *string `json:"DisasterRecoveryType,omitnil,omitempty" name:"DisasterRecoveryType"`

	// 数据复制方向（POSITIVE 正向 / REVERSE 反向）
	DataDirection *string `json:"DataDirection,omitnil,omitempty" name:"DataDirection"`

	// 跨云场景对端云名称（仅 DisasterRecoveryType=CROSS_CLOUD 时返回）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeerCloudName *string `json:"PeerCloudName,omitnil,omitempty" name:"PeerCloudName"`

	// 创建来源（LOCAL 本端创建 / PEER 对端创建）
	CreateFrom *string `json:"CreateFrom,omitnil,omitempty" name:"CreateFrom"`

	// 生命周期状态
	LifeState *string `json:"LifeState,omitnil,omitempty" name:"LifeState"`

	// 创建保护组的账户主账号 Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountUin *string `json:"AccountUin,omitnil,omitempty" name:"AccountUin"`

	// 创建保护组的子账号 Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitnil,omitempty" name:"SubAccountUin"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 绑定的已保护资源数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindProtectedResourceCount *int64 `json:"BindProtectedResourceCount,omitnil,omitempty" name:"BindProtectedResourceCount"`

	// RPO 异常（超过 15 分钟未同步）的复制对数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorRecoveryPointObjectiveCount *int64 `json:"ErrorRecoveryPointObjectiveCount,omitnil,omitempty" name:"ErrorRecoveryPointObjectiveCount"`

	// 已保护资源状态统计，key 为复制对状态，value 为该状态下的资源数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectedResourceStatusSet []*ProtectedResourceStatus `json:"ProtectedResourceStatusSet,omitnil,omitempty" name:"ProtectedResourceStatusSet"`
}

type ProtectGroupDeniedAction struct {
	// 保护组ID
	ProtectGroupId *string `json:"ProtectGroupId,omitnil,omitempty" name:"ProtectGroupId"`

	// 被禁止的操作列表（Action名称数组）
	DeniedActions []*DeniedAction `json:"DeniedActions,omitnil,omitempty" name:"DeniedActions"`
}

type ProtectInstance struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 客户端ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 客户端版本
	AgentVersion *string `json:"AgentVersion,omitnil,omitempty" name:"AgentVersion"`

	// 客户端状态
	AgentStatus *string `json:"AgentStatus,omitnil,omitempty" name:"AgentStatus"`

	// 最后心跳时间
	LastHeartbeatTime *string `json:"LastHeartbeatTime,omitnil,omitempty" name:"LastHeartbeatTime"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 最新备份点中记录的 CVM 基础信息
	ExtraInfo *string `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 该实例可用备份点数量
	BackupCount *uint64 `json:"BackupCount,omitnil,omitempty" name:"BackupCount"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 最近一次备份时间
	LatestBackupTime *string `json:"LatestBackupTime,omitnil,omitempty" name:"LatestBackupTime"`

	// 离线原因
	OfflineReason *string `json:"OfflineReason,omitnil,omitempty" name:"OfflineReason"`
}

type ProtectedResource struct {
	// 资源类型（与请求 SitePairType 一致，如 DISK/CFS/INSTANCE）
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 该类型下被保护的源端资源ID列表（DISK:disk-xxx / CFS:cfs-xxx / INSTANCE:ins-xxx）
	ResourceIdSet []*string `json:"ResourceIdSet,omitnil,omitempty" name:"ResourceIdSet"`
}

type ProtectedResourceOverview struct {
	// 受保护资源总数
	TotalProtectedCount *int64 `json:"TotalProtectedCount,omitnil,omitempty" name:"TotalProtectedCount"`

	// 总资源数
	TotalResourceCount *int64 `json:"TotalResourceCount,omitnil,omitempty" name:"TotalResourceCount"`

	// CVM 受保护统计
	Cvm *ResourceProtectStat `json:"Cvm,omitnil,omitempty" name:"Cvm"`

	// CFS 受保护统计
	CFS *ResourceProtectStat `json:"CFS,omitnil,omitempty" name:"CFS"`
}

type ProtectedResourceStatus struct {
	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 数量
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

// Predefined struct for user
type ReportAgentMetricsRequestParams struct {

}

type ReportAgentMetricsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ReportAgentMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportAgentMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportAgentMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportAgentMetricsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReportAgentMetricsResponse struct {
	*tchttp.BaseResponse
	Response *ReportAgentMetricsResponseParams `json:"Response"`
}

func (r *ReportAgentMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportAgentMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportGatewayHeartbeatRequestParams struct {

}

type ReportGatewayHeartbeatRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ReportGatewayHeartbeatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportGatewayHeartbeatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportGatewayHeartbeatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportGatewayHeartbeatResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReportGatewayHeartbeatResponse struct {
	*tchttp.BaseResponse
	Response *ReportGatewayHeartbeatResponseParams `json:"Response"`
}

func (r *ReportGatewayHeartbeatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportGatewayHeartbeatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportJobProgressRequestParams struct {

}

type ReportJobProgressRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ReportJobProgressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportJobProgressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportJobProgressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportJobProgressResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReportJobProgressResponse struct {
	*tchttp.BaseResponse
	Response *ReportJobProgressResponseParams `json:"Response"`
}

func (r *ReportJobProgressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportJobProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourcePlan struct {
	// 云服务器实例 ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 备份路径，[0,20]
	BackupPaths []*string `json:"BackupPaths,omitnil,omitempty" name:"BackupPaths"`

	// 包含文件类型，[0,20]
	IncludeFileTypes []*string `json:"IncludeFileTypes,omitnil,omitempty" name:"IncludeFileTypes"`

	// 排除路径，[0,20]
	ExcludePatterns []*string `json:"ExcludePatterns,omitnil,omitempty" name:"ExcludePatterns"`

	// 是否排除系统目录
	ExcludeSystemDirectories *bool `json:"ExcludeSystemDirectories,omitnil,omitempty" name:"ExcludeSystemDirectories"`

	// 是否立即触发全量备份
	ExecuteImmediately *bool `json:"ExecuteImmediately,omitnil,omitempty" name:"ExecuteImmediately"`
}

type ResourceProtectStat struct {
	// 受保护资源数
	ProtectedCount *int64 `json:"ProtectedCount,omitnil,omitempty" name:"ProtectedCount"`

	// 资源总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type RestoreTask struct {
	// 恢复任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 关联备份点 ID
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 源实例 ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 目标实例 ID
	TargetResourceId *string `json:"TargetResourceId,omitnil,omitempty" name:"TargetResourceId"`

	// 恢复路径列表
	RestorePaths []*string `json:"RestorePaths,omitnil,omitempty" name:"RestorePaths"`

	// 目标恢复位置
	TargetLocation *string `json:"TargetLocation,omitnil,omitempty" name:"TargetLocation"`

	// 任务状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 需恢复文件总数
	TotalFileCount *int64 `json:"TotalFileCount,omitnil,omitempty" name:"TotalFileCount"`

	// 需恢复数据总量（字节）
	TotalSize *int64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`

	// 需恢复数据总量（格式化，如 "1.5 GB"）
	TotalSizeFormatted *string `json:"TotalSizeFormatted,omitnil,omitempty" name:"TotalSizeFormatted"`

	// 已恢复文件数
	RestoreFileCount *int64 `json:"RestoreFileCount,omitnil,omitempty" name:"RestoreFileCount"`

	// 已恢复数据量（字节）
	RestoreSize *int64 `json:"RestoreSize,omitnil,omitempty" name:"RestoreSize"`

	// 已恢复数据量（格式化）
	RestoreSizeFormatted *string `json:"RestoreSizeFormatted,omitnil,omitempty" name:"RestoreSizeFormatted"`

	// 恢复进度（0-100）
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 关联 Job ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 任务开始时间（ISO 格式）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 恢复任务失败原因
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`

	// 冲突处理策略：skip-跳过/overwrite-覆盖/newer-保留最新版本/if_changed-内容变化时覆盖
	ConflictStrategy *string `json:"ConflictStrategy,omitnil,omitempty" name:"ConflictStrategy"`
}

// Predefined struct for user
type RunCopyPairTasksRequestParams struct {
	// 复制对ID列表
	CopyPairIds []*string `json:"CopyPairIds,omitnil,omitempty" name:"CopyPairIds"`

	// 要启动复制对的类型（DISK/INSTANCE/CFS）
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`
}

type RunCopyPairTasksRequest struct {
	*tchttp.BaseRequest
	
	// 复制对ID列表
	CopyPairIds []*string `json:"CopyPairIds,omitnil,omitempty" name:"CopyPairIds"`

	// 要启动复制对的类型（DISK/INSTANCE/CFS）
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`
}

func (r *RunCopyPairTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunCopyPairTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CopyPairIds")
	delete(f, "CopyPairType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunCopyPairTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunCopyPairTasksResponseParams struct {
	// 已启动复制任务的复制对ID列表
	CopyPairIds []*string `json:"CopyPairIds,omitnil,omitempty" name:"CopyPairIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunCopyPairTasksResponse struct {
	*tchttp.BaseResponse
	Response *RunCopyPairTasksResponseParams `json:"Response"`
}

func (r *RunCopyPairTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunCopyPairTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunFailoverCopyPairsRequestParams struct {
	// 复制对ID列表
	CopyPairIds []*string `json:"CopyPairIds,omitnil,omitempty" name:"CopyPairIds"`

	// 复制对类型，枚举值：DISK / INSTANCE / CFS。
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`

	// 切换类型，支持WAIT和NOW
	FailoverType *string `json:"FailoverType,omitnil,omitempty" name:"FailoverType"`
}

type RunFailoverCopyPairsRequest struct {
	*tchttp.BaseRequest
	
	// 复制对ID列表
	CopyPairIds []*string `json:"CopyPairIds,omitnil,omitempty" name:"CopyPairIds"`

	// 复制对类型，枚举值：DISK / INSTANCE / CFS。
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`

	// 切换类型，支持WAIT和NOW
	FailoverType *string `json:"FailoverType,omitnil,omitempty" name:"FailoverType"`
}

func (r *RunFailoverCopyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunFailoverCopyPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CopyPairIds")
	delete(f, "CopyPairType")
	delete(f, "FailoverType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunFailoverCopyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunFailoverCopyPairsResponseParams struct {
	// 故障切换任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunFailoverCopyPairsResponse struct {
	*tchttp.BaseResponse
	Response *RunFailoverCopyPairsResponseParams `json:"Response"`
}

func (r *RunFailoverCopyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunFailoverCopyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunInstancesWithBackupGroupRequestParams struct {
	// 备份组ID
	BackupGroupId *string `json:"BackupGroupId,omitnil,omitempty" name:"BackupGroupId"`
}

type RunInstancesWithBackupGroupRequest struct {
	*tchttp.BaseRequest
	
	// 备份组ID
	BackupGroupId *string `json:"BackupGroupId,omitnil,omitempty" name:"BackupGroupId"`
}

func (r *RunInstancesWithBackupGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunInstancesWithBackupGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunInstancesWithBackupGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunInstancesWithBackupGroupResponseParams struct {
	// 创建的实例ID
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunInstancesWithBackupGroupResponse struct {
	*tchttp.BaseResponse
	Response *RunInstancesWithBackupGroupResponseParams `json:"Response"`
}

func (r *RunInstancesWithBackupGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunInstancesWithBackupGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunSecurityServiceEnabled struct {
	// 是否开启该服务。取值范围：TRUE（开启）/FALSE（不开启）。默认取值：TRUE。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type SecurityGroupMapping struct {
	// 安全组映射ID
	SecurityGroupMappingId *string `json:"SecurityGroupMappingId,omitnil,omitempty" name:"SecurityGroupMappingId"`

	// 安全组映射所属的站点对ID
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 生产端安全组ID
	SourceSecurityGroupId *string `json:"SourceSecurityGroupId,omitnil,omitempty" name:"SourceSecurityGroupId"`

	// 容灾端安全组ID
	TargetSecurityGroupId *string `json:"TargetSecurityGroupId,omitnil,omitempty" name:"TargetSecurityGroupId"`

	// 安全组映射的生命状态；NORMAL:正常。
	LifeState *string `json:"LifeState,omitnil,omitempty" name:"LifeState"`
}

type SitePair struct {
	// 用户AppId
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 容灾策略ID
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 容灾策略名称
	SitePairName *string `json:"SitePairName,omitnil,omitempty" name:"SitePairName"`

	// 容灾策略类型（产品类型，如 DISK/CFS/INSTANCE 等）
	SitePairType *string `json:"SitePairType,omitnil,omitempty" name:"SitePairType"`

	// 容灾策略状态
	SitePairState *string `json:"SitePairState,omitnil,omitempty" name:"SitePairState"`

	// 生产地域
	SourceRegion *string `json:"SourceRegion,omitnil,omitempty" name:"SourceRegion"`

	// 生产可用区
	SourceZone *string `json:"SourceZone,omitnil,omitempty" name:"SourceZone"`

	// 容灾地域
	TargetRegion *string `json:"TargetRegion,omitnil,omitempty" name:"TargetRegion"`

	// 容灾可用区
	TargetZone *string `json:"TargetZone,omitnil,omitempty" name:"TargetZone"`

	// 生产端VPC
	SourceVpc *string `json:"SourceVpc,omitnil,omitempty" name:"SourceVpc"`

	// 容灾端VPC
	TargetVpc *string `json:"TargetVpc,omitnil,omitempty" name:"TargetVpc"`

	// 复制技术（SYN 同步 / ASY 异步）
	CopyType *string `json:"CopyType,omitnil,omitempty" name:"CopyType"`

	// 容灾类型（CROSS_ZONE 跨可用区 / CROSS_REGION 跨地域 / CROSS_CLOUD 跨云）
	DisasterRecoveryType *string `json:"DisasterRecoveryType,omitnil,omitempty" name:"DisasterRecoveryType"`

	// 创建来源（LOCAL 本端创建 / PEER 对端创建）
	CreateFrom *string `json:"CreateFrom,omitnil,omitempty" name:"CreateFrom"`

	// 创建容灾策略的账户主账号 Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountUin *string `json:"AccountUin,omitnil,omitempty" name:"AccountUin"`

	// 创建容灾策略的子账户 Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitnil,omitempty" name:"SubAccountUin"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 已绑定的保护组数量
	BindProtectGroupCount *int64 `json:"BindProtectGroupCount,omitnil,omitempty" name:"BindProtectGroupCount"`

	// RPO 异常的复制对ID列表（最近一次保护点距今超过15分钟的复制对）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorRecoveryPointObjectiveCopyPairSet []*string `json:"ErrorRecoveryPointObjectiveCopyPairSet,omitnil,omitempty" name:"ErrorRecoveryPointObjectiveCopyPairSet"`

	// 已保护的资源列表（按资源类型分组）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectedResourceSet []*ProtectedResource `json:"ProtectedResourceSet,omitnil,omitempty" name:"ProtectedResourceSet"`

	// 已保护资源的状态统计，key 为复制对状态，value 为该状态下的资源数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectedResourceStatusSet []*ProtectedResourceStatus `json:"ProtectedResourceStatusSet,omitnil,omitempty" name:"ProtectedResourceStatusSet"`

	// 跨云场景下的额外信息（仅 IsCrossCloud=true 时返回，非跨云为 null）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrossCloudDetails *CrossCloudDetails `json:"CrossCloudDetails,omitnil,omitempty" name:"CrossCloudDetails"`
}

type SitePairDeniedAction struct {
	// 容灾策略ID
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 该容灾策略当前被禁止执行的操作列表
	DeniedActions []*DeniedAction `json:"DeniedActions,omitnil,omitempty" name:"DeniedActions"`
}

// Predefined struct for user
type StopCopyPairTasksRequestParams struct {
	// 复制对ID列表
	CopyPairIds []*string `json:"CopyPairIds,omitnil,omitempty" name:"CopyPairIds"`

	// 复制对类型（DISK/INSTANCE/CFS）
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`
}

type StopCopyPairTasksRequest struct {
	*tchttp.BaseRequest
	
	// 复制对ID列表
	CopyPairIds []*string `json:"CopyPairIds,omitnil,omitempty" name:"CopyPairIds"`

	// 复制对类型（DISK/INSTANCE/CFS）
	CopyPairType *string `json:"CopyPairType,omitnil,omitempty" name:"CopyPairType"`
}

func (r *StopCopyPairTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCopyPairTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CopyPairIds")
	delete(f, "CopyPairType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopCopyPairTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCopyPairTasksResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopCopyPairTasksResponse struct {
	*tchttp.BaseResponse
	Response *StopCopyPairTasksResponseParams `json:"Response"`
}

func (r *StopCopyPairTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCopyPairTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SupportRegionInfo struct {
	// 生产地域。
	SourceRegion *string `json:"SourceRegion,omitnil,omitempty" name:"SourceRegion"`

	// 支持类型：REGION（地域级，整个生产地域均支持容灾）；ZONE（可用区级，按 SupportZoneRules 控制粒度）。
	SupportType *string `json:"SupportType,omitnil,omitempty" name:"SupportType"`

	// 配置状态：valid（生效）/ invalid（停用）。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 可用区级容灾规则列表。仅当 SupportType=ZONE 时有效；REGION 类型时该字段返回空数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportZoneRules []*SupportZoneRule `json:"SupportZoneRules,omitnil,omitempty" name:"SupportZoneRules"`
}

type SupportZoneRule struct {
	// 生产可用区。
	SourceZone *string `json:"SourceZone,omitnil,omitempty" name:"SourceZone"`

	// 是否支持容灾到生产地域内的全部可用区。true 时 TargetZones 可忽略。
	IsAllZoneSupport *bool `json:"IsAllZoneSupport,omitnil,omitempty" name:"IsAllZoneSupport"`

	// 目标可用区列表。当 IsAllZoneSupport=false 时枚举具体可容灾到的可用区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetZones []*string `json:"TargetZones,omitnil,omitempty" name:"TargetZones"`
}

type TypeCount struct {
	// 备份库类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 备份库数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

// Predefined struct for user
type UnbindAutoBackupPolicyRequestParams struct {

}

type UnbindAutoBackupPolicyRequest struct {
	*tchttp.BaseRequest
	
}

func (r *UnbindAutoBackupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindAutoBackupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindAutoBackupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindAutoBackupPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindAutoBackupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *UnbindAutoBackupPolicyResponseParams `json:"Response"`
}

func (r *UnbindAutoBackupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindAutoBackupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VaultDeniedAction struct {
	// 备份库实例ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// 被禁止的操作列表
	DeniedActions []*DeniedAction `json:"DeniedActions,omitnil,omitempty" name:"DeniedActions"`
}

type VirtualPrivateCloud struct {
	// 私有网络ID，形如 vpc-xxxxxxxx。私有网络ID可通过登录控制台查询，也可通过调用接口 [DescribeVpcEx]的返回值中的unVpcId字段获取。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络子网ID，形如 subnet-xxxxxxxx。私有网络子网ID可通过登录控制台查询，也可通过调用接口 [DescribeSubnets](https://cloud.tencent.com/document/api/215/15784) 的返回值中的 unSubnetId 字段获取。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 私有网络子网名称。
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// 是否用作公网网关。公网网关只有在实例拥有公网IP以及处于私有网络下时才能正常使用。取值范围：true（用作公网网关）/false（不作为公网网关），默认取值：false。
	AsVpcGateway *bool `json:"AsVpcGateway,omitnil,omitempty" name:"AsVpcGateway"`

	// 私有网络子网 IP 数组，在创建实例、修改实例 vpc 属性操作中可使用此参数。当前仅批量创建多台实例时支持传入相同子网的多个 IP。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// 私有网络名称，仅做展示用。
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 为弹性网卡指定随机生成的 IPv6 地址数量。
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitnil,omitempty" name:"Ipv6AddressCount"`
}

type VpcMapping struct {
	// 映射规则主键ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 所属容灾策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SitePairId *string `json:"SitePairId,omitnil,omitempty" name:"SitePairId"`

	// 源端VPC ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceVpc *string `json:"SourceVpc,omitnil,omitempty" name:"SourceVpc"`

	// 源端子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceSubnet *string `json:"SourceSubnet,omitnil,omitempty" name:"SourceSubnet"`

	// 目标端VPC ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetVpc *string `json:"TargetVpc,omitnil,omitempty" name:"TargetVpc"`

	// 目标端子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetSubnet *string `json:"TargetSubnet,omitnil,omitempty" name:"TargetSubnet"`

	// 映射状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 生命周期状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifeState *string `json:"LifeState,omitnil,omitempty" name:"LifeState"`
}