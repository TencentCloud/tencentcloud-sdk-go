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

package v20211228

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type ActionAlterUserRequestParams struct {
	// 用户信息
	UserInfo *UserInfo `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// api接口类型
	ApiType *string `json:"ApiType,omitnil,omitempty" name:"ApiType"`

	// 用户权限类型 0:普通用户 1:管理员
	UserPrivilege *int64 `json:"UserPrivilege,omitnil,omitempty" name:"UserPrivilege"`
}

type ActionAlterUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户信息
	UserInfo *UserInfo `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// api接口类型
	ApiType *string `json:"ApiType,omitnil,omitempty" name:"ApiType"`

	// 用户权限类型 0:普通用户 1:管理员
	UserPrivilege *int64 `json:"UserPrivilege,omitnil,omitempty" name:"UserPrivilege"`
}

func (r *ActionAlterUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActionAlterUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserInfo")
	delete(f, "ApiType")
	delete(f, "UserPrivilege")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActionAlterUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActionAlterUserResponseParams struct {
	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ActionAlterUserResponse struct {
	*tchttp.BaseResponse
	Response *ActionAlterUserResponseParams `json:"Response"`
}

func (r *ActionAlterUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActionAlterUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachCBSSpec struct {
	// 节点磁盘类型，例如“CLOUD_SSD”\"CLOUD_PREMIUM"
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 磁盘容量，单位G
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 磁盘总数
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// 描述
	DiskDesc *string `json:"DiskDesc,omitnil,omitempty" name:"DiskDesc"`
}

type BackUpJobDisplay struct {
	// 备份实例id
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 备份实例名
	Snapshot *string `json:"Snapshot,omitnil,omitempty" name:"Snapshot"`

	// 备份数据量
	BackUpSize *int64 `json:"BackUpSize,omitnil,omitempty" name:"BackUpSize"`

	// 备份单副本数据量
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackUpSingleSize *int64 `json:"BackUpSingleSize,omitnil,omitempty" name:"BackUpSingleSize"`

	// 实例创建时间
	BackUpTime *string `json:"BackUpTime,omitnil,omitempty" name:"BackUpTime"`

	// 实例过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 实例状态
	JobStatus *string `json:"JobStatus,omitnil,omitempty" name:"JobStatus"`

	// 0为默认。1时是对远端的doris进行备份，不周期，一次性
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupType *int64 `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 0为默认。1时是立即备份。2时是迁移
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupTimeType *int64 `json:"BackupTimeType,omitnil,omitempty" name:"BackupTimeType"`

	// 远端doris的连接信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DorisSourceInfo *DorisSourceInfo `json:"DorisSourceInfo,omitnil,omitempty" name:"DorisSourceInfo"`

	// 实例状态对应的数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobStatusNum *int64 `json:"JobStatusNum,omitnil,omitempty" name:"JobStatusNum"`

	// 备份实例中关于cos的信息	
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupCosInfo *BackupCosInfo `json:"BackupCosInfo,omitnil,omitempty" name:"BackupCosInfo"`
}

type BackupCosInfo struct {
	// 备份文件所在的cos桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// 备份文件所在的完整cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPath *string `json:"CosPath,omitnil,omitempty" name:"CosPath"`

	// 备份文件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapShotPath *string `json:"SnapShotPath,omitnil,omitempty" name:"SnapShotPath"`
}

type BackupStatus struct {
	// 备份任务id
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 快照名称
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 库名
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 状态
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 备份对象
	BackupObjects *string `json:"BackupObjects,omitnil,omitempty" name:"BackupObjects"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 快照结束时间
	SnapshotFinishedTime *string `json:"SnapshotFinishedTime,omitnil,omitempty" name:"SnapshotFinishedTime"`

	// 上传结束时间
	UploadFinishedTime *string `json:"UploadFinishedTime,omitnil,omitempty" name:"UploadFinishedTime"`

	// 结束时间
	FinishedTime *string `json:"FinishedTime,omitnil,omitempty" name:"FinishedTime"`

	// 未完成任务
	UnfinishedTasks *string `json:"UnfinishedTasks,omitnil,omitempty" name:"UnfinishedTasks"`

	// 进度
	Progress *string `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 错误信息
	TaskErrMsg *string `json:"TaskErrMsg,omitnil,omitempty" name:"TaskErrMsg"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 超时信息
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 备份实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupJobId *int64 `json:"BackupJobId,omitnil,omitempty" name:"BackupJobId"`

	// 实例对应snapshoit的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type BackupTableContent struct {
	// 数据库
	// 注意：此字段可能返回 null，表示取不到有效值。
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 表总字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalBytes *int64 `json:"TotalBytes,omitnil,omitempty" name:"TotalBytes"`

	// 表单个副本的大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	SingleReplicaBytes *string `json:"SingleReplicaBytes,omitnil,omitempty" name:"SingleReplicaBytes"`

	// 备份状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupStatus *int64 `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// 备份的错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupErrorMsg *string `json:"BackupErrorMsg,omitnil,omitempty" name:"BackupErrorMsg"`

	// 改库表是否绑定降冷策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsOpenCoolDown *bool `json:"IsOpenCoolDown,omitnil,omitempty" name:"IsOpenCoolDown"`
}

type BindUser struct {
	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 主机信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

// Predefined struct for user
type CancelBackupJobRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要取消的备份实例id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
}

type CancelBackupJobRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要取消的备份实例id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
}

func (r *CancelBackupJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelBackupJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackUpJobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelBackupJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelBackupJobResponseParams struct {
	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelBackupJobResponse struct {
	*tchttp.BaseResponse
	Response *CancelBackupJobResponseParams `json:"Response"`
}

func (r *CancelBackupJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelBackupJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChargeProperties struct {
	// 计费类型，“PREPAID” 预付费，“POSTPAID_BY_HOUR” 后付费
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 是否自动续费，1表示自动续费开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 计费时间长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 计费时间单位，“m”表示月等
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
}

// Predefined struct for user
type CheckCoolDownWorkingVariableConfigCorrectRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CheckCoolDownWorkingVariableConfigCorrectRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *CheckCoolDownWorkingVariableConfigCorrectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCoolDownWorkingVariableConfigCorrectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckCoolDownWorkingVariableConfigCorrectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckCoolDownWorkingVariableConfigCorrectResponseParams struct {
	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckCoolDownWorkingVariableConfigCorrectResponse struct {
	*tchttp.BaseResponse
	Response *CheckCoolDownWorkingVariableConfigCorrectResponseParams `json:"Response"`
}

func (r *CheckCoolDownWorkingVariableConfigCorrectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCoolDownWorkingVariableConfigCorrectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterConfigsHistory struct {
	// 配置文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 修改后的配置文件内容，base64编码
	NewConfValue *string `json:"NewConfValue,omitnil,omitempty" name:"NewConfValue"`

	// 修改前的配置文件内容，base64编码
	OldConfValue *string `json:"OldConfValue,omitnil,omitempty" name:"OldConfValue"`

	// 修改原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 修改子账号id
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`
}

type ClusterConfigsInfoFromEMR struct {
	// 配置文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 配置文件对应的相关属性信息
	FileConf *string `json:"FileConf,omitnil,omitempty" name:"FileConf"`

	// 配置文件对应的其他属性信息
	KeyConf *string `json:"KeyConf,omitnil,omitempty" name:"KeyConf"`

	// 配置文件的内容，base64编码
	OriParam *string `json:"OriParam,omitnil,omitempty" name:"OriParam"`

	// 用于表示当前配置文件是不是有过修改后没有重启，提醒用户需要重启
	NeedRestart *int64 `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// 配置文件路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// 配置文件kv值
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: FileKeyValues is deprecated.
	FileKeyValues *string `json:"FileKeyValues,omitnil,omitempty" name:"FileKeyValues"`

	// 配置文件kv值
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileKeyValuesNew []*ConfigKeyValue `json:"FileKeyValuesNew,omitnil,omitempty" name:"FileKeyValuesNew"`
}

type ConfigKeyValue struct {
	// key
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 1-只读，2-可修改但不可删除，3-可删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	Display *int64 `json:"Display,omitnil,omitempty" name:"Display"`

	// 0不支持 1支持热更新
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportHotUpdate *int64 `json:"SupportHotUpdate,omitnil,omitempty" name:"SupportHotUpdate"`
}

type ConfigSubmitContext struct {
	// 配置文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 配置文件新内容，base64编码
	NewConfValue *string `json:"NewConfValue,omitnil,omitempty" name:"NewConfValue"`

	// 配置文件旧内容，base64编码
	OldConfValue *string `json:"OldConfValue,omitnil,omitempty" name:"OldConfValue"`

	// 文件路径
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`
}

type CoolDownBackend struct {
	// 字段：Host
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 字段：DataUsedCapacity
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataUsedCapacity *string `json:"DataUsedCapacity,omitnil,omitempty" name:"DataUsedCapacity"`

	// 字段：TotalCapacity
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCapacity *string `json:"TotalCapacity,omitnil,omitempty" name:"TotalCapacity"`

	// 字段：RemoteUsedCapacity
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemoteUsedCapacity *string `json:"RemoteUsedCapacity,omitnil,omitempty" name:"RemoteUsedCapacity"`
}

type CoolDownPolicyInfo struct {
	// 策略名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// cooldown_ttl
	// 注意：此字段可能返回 null，表示取不到有效值。
	CooldownDatetime *string `json:"CooldownDatetime,omitnil,omitempty" name:"CooldownDatetime"`

	// cooldown_datetime
	// 注意：此字段可能返回 null，表示取不到有效值。
	CooldownTtl *string `json:"CooldownTtl,omitnil,omitempty" name:"CooldownTtl"`
}

type CoolDownTableDataInfo struct {
	// 列：DatabaseName
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 列：TableName
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 列：Size
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// 列：RemoteSize
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemoteSize *string `json:"RemoteSize,omitnil,omitempty" name:"RemoteSize"`
}

type CosSourceInfo struct {
	// cos认证中的Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretId *string `json:"SecretId,omitnil,omitempty" name:"SecretId"`

	// cos认证中的key
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// cos认证中的路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPath *string `json:"CosPath,omitnil,omitempty" name:"CosPath"`
}

// Predefined struct for user
type CreateBackUpScheduleRequestParams struct {
	// 编辑时需要传
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`

	// 选择的星期 逗号分隔
	// 废弃：使用ScheduleInfo
	WeekDays *string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// 执行小时
	// 废弃：使用ScheduleInfo
	ExecuteHour *int64 `json:"ExecuteHour,omitnil,omitempty" name:"ExecuteHour"`

	// 备份表列表
	BackUpTables []*BackupTableContent `json:"BackUpTables,omitnil,omitempty" name:"BackUpTables"`

	// 0为默认。1时是对远端的doris进行备份，不周期，一次性
	BackupType *int64 `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 远端doris集群的连接信息
	DorisSourceInfo *DorisSourceInfo `json:"DorisSourceInfo,omitnil,omitempty" name:"DorisSourceInfo"`

	// 0为默认。1时是一次性备份。2时是远端备份
	BackupTimeType *int64 `json:"BackupTimeType,omitnil,omitempty" name:"BackupTimeType"`

	// 0为默认。1时是备份完成后立即恢复
	RestoreType *int64 `json:"RestoreType,omitnil,omitempty" name:"RestoreType"`

	// 0为默认。1时是提供自定义的secret连接cos
	AuthType *int64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// cos认证的信息
	CosSourceInfo *CosSourceInfo `json:"CosSourceInfo,omitnil,omitempty" name:"CosSourceInfo"`
}

type CreateBackUpScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 编辑时需要传
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`

	// 选择的星期 逗号分隔
	// 废弃：使用ScheduleInfo
	WeekDays *string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// 执行小时
	// 废弃：使用ScheduleInfo
	ExecuteHour *int64 `json:"ExecuteHour,omitnil,omitempty" name:"ExecuteHour"`

	// 备份表列表
	BackUpTables []*BackupTableContent `json:"BackUpTables,omitnil,omitempty" name:"BackUpTables"`

	// 0为默认。1时是对远端的doris进行备份，不周期，一次性
	BackupType *int64 `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 远端doris集群的连接信息
	DorisSourceInfo *DorisSourceInfo `json:"DorisSourceInfo,omitnil,omitempty" name:"DorisSourceInfo"`

	// 0为默认。1时是一次性备份。2时是远端备份
	BackupTimeType *int64 `json:"BackupTimeType,omitnil,omitempty" name:"BackupTimeType"`

	// 0为默认。1时是备份完成后立即恢复
	RestoreType *int64 `json:"RestoreType,omitnil,omitempty" name:"RestoreType"`

	// 0为默认。1时是提供自定义的secret连接cos
	AuthType *int64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// cos认证的信息
	CosSourceInfo *CosSourceInfo `json:"CosSourceInfo,omitnil,omitempty" name:"CosSourceInfo"`
}

func (r *CreateBackUpScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackUpScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduleId")
	delete(f, "WeekDays")
	delete(f, "ExecuteHour")
	delete(f, "BackUpTables")
	delete(f, "BackupType")
	delete(f, "DorisSourceInfo")
	delete(f, "BackupTimeType")
	delete(f, "RestoreType")
	delete(f, "AuthType")
	delete(f, "CosSourceInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackUpScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackUpScheduleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBackUpScheduleResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackUpScheduleResponseParams `json:"Response"`
}

func (r *CreateBackUpScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackUpScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCoolDownPolicyRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// cooldown_ttl
	CoolDownTtl *string `json:"CoolDownTtl,omitnil,omitempty" name:"CoolDownTtl"`

	// cooldown_datetime
	CoolDownDatetime *string `json:"CoolDownDatetime,omitnil,omitempty" name:"CoolDownDatetime"`
}

type CreateCoolDownPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// cooldown_ttl
	CoolDownTtl *string `json:"CoolDownTtl,omitnil,omitempty" name:"CoolDownTtl"`

	// cooldown_datetime
	CoolDownDatetime *string `json:"CoolDownDatetime,omitnil,omitempty" name:"CoolDownDatetime"`
}

func (r *CreateCoolDownPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCoolDownPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PolicyName")
	delete(f, "CoolDownTtl")
	delete(f, "CoolDownDatetime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCoolDownPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCoolDownPolicyResponseParams struct {
	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCoolDownPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateCoolDownPolicyResponseParams `json:"Response"`
}

func (r *CreateCoolDownPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCoolDownPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceNewRequestParams struct {
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// FE规格
	FeSpec *CreateInstanceSpec `json:"FeSpec,omitnil,omitempty" name:"FeSpec"`

	// BE规格
	BeSpec *CreateInstanceSpec `json:"BeSpec,omitnil,omitempty" name:"BeSpec"`

	// 是否高可用
	HaFlag *bool `json:"HaFlag,omitnil,omitempty" name:"HaFlag"`

	// 用户VPCID
	UserVPCId *string `json:"UserVPCId,omitnil,omitempty" name:"UserVPCId"`

	// 用户子网ID
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// 产品版本号
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// 付费类型
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// 实例名字
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 数据库密码
	DorisUserPwd *string `json:"DorisUserPwd,omitnil,omitempty" name:"DorisUserPwd"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 高可用类型：
	// 0：非高可用（只有1个FE，FeSpec.CreateInstanceSpec.Count=1），
	// 1：读高可用（至少需部署3个FE，FeSpec.CreateInstanceSpec.Count>=3，且为奇数），
	// 2：读写高可用（至少需部署5个FE，FeSpec.CreateInstanceSpec.Count>=5，且为奇数）。
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`

	// 表名大小写是否敏感，0：敏感；1：不敏感，以小写进行比较；2：不敏感，表名改为以小写存储
	CaseSensitive *int64 `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// 是否开启多可用区
	EnableMultiZones *bool `json:"EnableMultiZones,omitnil,omitempty" name:"EnableMultiZones"`

	// 开启多可用区后，用户的所有可用区和子网信息
	UserMultiZoneInfos *NetworkInfo `json:"UserMultiZoneInfos,omitnil,omitempty" name:"UserMultiZoneInfos"`
}

type CreateInstanceNewRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// FE规格
	FeSpec *CreateInstanceSpec `json:"FeSpec,omitnil,omitempty" name:"FeSpec"`

	// BE规格
	BeSpec *CreateInstanceSpec `json:"BeSpec,omitnil,omitempty" name:"BeSpec"`

	// 是否高可用
	HaFlag *bool `json:"HaFlag,omitnil,omitempty" name:"HaFlag"`

	// 用户VPCID
	UserVPCId *string `json:"UserVPCId,omitnil,omitempty" name:"UserVPCId"`

	// 用户子网ID
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// 产品版本号
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// 付费类型
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// 实例名字
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 数据库密码
	DorisUserPwd *string `json:"DorisUserPwd,omitnil,omitempty" name:"DorisUserPwd"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 高可用类型：
	// 0：非高可用（只有1个FE，FeSpec.CreateInstanceSpec.Count=1），
	// 1：读高可用（至少需部署3个FE，FeSpec.CreateInstanceSpec.Count>=3，且为奇数），
	// 2：读写高可用（至少需部署5个FE，FeSpec.CreateInstanceSpec.Count>=5，且为奇数）。
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`

	// 表名大小写是否敏感，0：敏感；1：不敏感，以小写进行比较；2：不敏感，表名改为以小写存储
	CaseSensitive *int64 `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// 是否开启多可用区
	EnableMultiZones *bool `json:"EnableMultiZones,omitnil,omitempty" name:"EnableMultiZones"`

	// 开启多可用区后，用户的所有可用区和子网信息
	UserMultiZoneInfos *NetworkInfo `json:"UserMultiZoneInfos,omitnil,omitempty" name:"UserMultiZoneInfos"`
}

func (r *CreateInstanceNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "FeSpec")
	delete(f, "BeSpec")
	delete(f, "HaFlag")
	delete(f, "UserVPCId")
	delete(f, "UserSubnetId")
	delete(f, "ProductVersion")
	delete(f, "ChargeProperties")
	delete(f, "InstanceName")
	delete(f, "DorisUserPwd")
	delete(f, "Tags")
	delete(f, "HaType")
	delete(f, "CaseSensitive")
	delete(f, "EnableMultiZones")
	delete(f, "UserMultiZoneInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceNewResponseParams struct {
	// 流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceNewResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceNewResponseParams `json:"Response"`
}

func (r *CreateInstanceNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceSpec struct {
	// 规格名字
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 数量
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 云盘大小
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

// Predefined struct for user
type CreateWorkloadGroupRequestParams struct {
	// 集群id	
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 资源组配置
	WorkloadGroup *WorkloadGroupConfig `json:"WorkloadGroup,omitnil,omitempty" name:"WorkloadGroup"`
}

type CreateWorkloadGroupRequest struct {
	*tchttp.BaseRequest
	
	// 集群id	
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 资源组配置
	WorkloadGroup *WorkloadGroupConfig `json:"WorkloadGroup,omitnil,omitempty" name:"WorkloadGroup"`
}

func (r *CreateWorkloadGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkloadGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "WorkloadGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkloadGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkloadGroupResponseParams struct {
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWorkloadGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkloadGroupResponseParams `json:"Response"`
}

func (r *CreateWorkloadGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkloadGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataBaseAuditRecord struct {
	// 查询用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsUser *string `json:"OsUser,omitnil,omitempty" name:"OsUser"`

	// 查询ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitialQueryId *string `json:"InitialQueryId,omitnil,omitempty" name:"InitialQueryId"`

	// SQL语句
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryStartTime *string `json:"QueryStartTime,omitnil,omitempty" name:"QueryStartTime"`

	// 执行耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	DurationMs *int64 `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// 读取行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadRows *int64 `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// 读取字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultRows *int64 `json:"ResultRows,omitnil,omitempty" name:"ResultRows"`

	// 结果字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultBytes *uint64 `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// 内存
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryUsage *int64 `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`

	// 初始查询IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitialAddress *string `json:"InitialAddress,omitnil,omitempty" name:"InitialAddress"`

	// 数据库
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// sql类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// catalog名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`
}

// Predefined struct for user
type DeleteBackUpDataRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`

	// 是否删除所有实例
	IsDeleteAll *bool `json:"IsDeleteAll,omitnil,omitempty" name:"IsDeleteAll"`
}

type DeleteBackUpDataRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`

	// 是否删除所有实例
	IsDeleteAll *bool `json:"IsDeleteAll,omitnil,omitempty" name:"IsDeleteAll"`
}

func (r *DeleteBackUpDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackUpDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackUpJobId")
	delete(f, "IsDeleteAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBackUpDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackUpDataResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBackUpDataResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBackUpDataResponseParams `json:"Response"`
}

func (r *DeleteBackUpDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackUpDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkloadGroupRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要删除的资源组名称
	WorkloadGroupName *string `json:"WorkloadGroupName,omitnil,omitempty" name:"WorkloadGroupName"`
}

type DeleteWorkloadGroupRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要删除的资源组名称
	WorkloadGroupName *string `json:"WorkloadGroupName,omitnil,omitempty" name:"WorkloadGroupName"`
}

func (r *DeleteWorkloadGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkloadGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "WorkloadGroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkloadGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkloadGroupResponseParams struct {
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWorkloadGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWorkloadGroupResponseParams `json:"Response"`
}

func (r *DeleteWorkloadGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkloadGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAreaRegionRequestParams struct {
	// 是否是国际站
	IsInternationalSite *bool `json:"IsInternationalSite,omitnil,omitempty" name:"IsInternationalSite"`
}

type DescribeAreaRegionRequest struct {
	*tchttp.BaseRequest
	
	// 是否是国际站
	IsInternationalSite *bool `json:"IsInternationalSite,omitnil,omitempty" name:"IsInternationalSite"`
}

func (r *DescribeAreaRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAreaRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsInternationalSite")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAreaRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAreaRegionResponseParams struct {
	// 地域列表
	Items []*RegionAreaInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 前端规则描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrontEndRules []*FrontEndRule `json:"FrontEndRules,omitnil,omitempty" name:"FrontEndRules"`

	// 返回可用的白名单名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableWhiteListNames []*string `json:"AvailableWhiteListNames,omitnil,omitempty" name:"AvailableWhiteListNames"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAreaRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAreaRegionResponseParams `json:"Response"`
}

func (r *DescribeAreaRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAreaRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpJobDetailRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
}

type DescribeBackUpJobDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
}

func (r *DescribeBackUpJobDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpJobDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackUpJobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackUpJobDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpJobDetailResponseParams struct {
	// 备份表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableContents []*BackupTableContent `json:"TableContents,omitnil,omitempty" name:"TableContents"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackUpJobDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackUpJobDetailResponseParams `json:"Response"`
}

func (r *DescribeBackUpJobDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpJobDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpJobRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页号
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// jobid的string类型
	JobIdFiltersStr *string `json:"JobIdFiltersStr,omitnil,omitempty" name:"JobIdFiltersStr"`
}

type DescribeBackUpJobRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页号
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// jobid的string类型
	JobIdFiltersStr *string `json:"JobIdFiltersStr,omitnil,omitempty" name:"JobIdFiltersStr"`
}

func (r *DescribeBackUpJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "JobIdFiltersStr")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackUpJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpJobResponseParams struct {
	// 任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackUpJobs []*BackUpJobDisplay `json:"BackUpJobs,omitnil,omitempty" name:"BackUpJobs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackUpJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackUpJobResponseParams `json:"Response"`
}

func (r *DescribeBackUpJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpSchedulesRequestParams struct {

}

type DescribeBackUpSchedulesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBackUpSchedulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpSchedulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackUpSchedulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpSchedulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackUpSchedulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackUpSchedulesResponseParams `json:"Response"`
}

func (r *DescribeBackUpSchedulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpSchedulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpTablesRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 0为默认。1时是对远端的doris进行备份，不周期，一次性。2时为cos恢复，一次性
	BackupType *int64 `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 远端doris集群的连接信息
	DorisSourceInfo *DorisSourceInfo `json:"DorisSourceInfo,omitnil,omitempty" name:"DorisSourceInfo"`

	// cos信息
	CosSourceInfo *CosSourceInfo `json:"CosSourceInfo,omitnil,omitempty" name:"CosSourceInfo"`
}

type DescribeBackUpTablesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 0为默认。1时是对远端的doris进行备份，不周期，一次性。2时为cos恢复，一次性
	BackupType *int64 `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 远端doris集群的连接信息
	DorisSourceInfo *DorisSourceInfo `json:"DorisSourceInfo,omitnil,omitempty" name:"DorisSourceInfo"`

	// cos信息
	CosSourceInfo *CosSourceInfo `json:"CosSourceInfo,omitnil,omitempty" name:"CosSourceInfo"`
}

func (r *DescribeBackUpTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupType")
	delete(f, "DorisSourceInfo")
	delete(f, "CosSourceInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackUpTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpTablesResponseParams struct {
	// 可备份表列表
	AvailableTables []*BackupTableContent `json:"AvailableTables,omitnil,omitempty" name:"AvailableTables"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackUpTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackUpTablesResponseParams `json:"Response"`
}

func (r *DescribeBackUpTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpTaskDetailRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
}

type DescribeBackUpTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
}

func (r *DescribeBackUpTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackUpJobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackUpTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpTaskDetailResponseParams struct {
	// 备份任务进度详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupStatus []*BackupStatus `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackUpTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackUpTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeBackUpTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterConfigsHistoryRequestParams struct {
	// 集群id名称
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 配置修改历史的时间范围开始
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 配置修改历史的时间范围结束
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 需要查询的配置文件名称数组，如果为空则查询全部历史记录。目前支持的配置文件名称有：
	// apache_hdfs_broker.conf、be.conf、fe.conf、core-site.xml、hdfs-site.xml、odbcinst.ini
	ConfigFileNames []*string `json:"ConfigFileNames,omitnil,omitempty" name:"ConfigFileNames"`
}

type DescribeClusterConfigsHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 集群id名称
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 配置修改历史的时间范围开始
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 配置修改历史的时间范围结束
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 需要查询的配置文件名称数组，如果为空则查询全部历史记录。目前支持的配置文件名称有：
	// apache_hdfs_broker.conf、be.conf、fe.conf、core-site.xml、hdfs-site.xml、odbcinst.ini
	ConfigFileNames []*string `json:"ConfigFileNames,omitnil,omitempty" name:"ConfigFileNames"`
}

func (r *DescribeClusterConfigsHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterConfigsHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ConfigFileNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterConfigsHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterConfigsHistoryResponseParams struct {
	// 实例总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 配置文件修改历史
	ClusterConfHistory []*ClusterConfigsHistory `json:"ClusterConfHistory,omitnil,omitempty" name:"ClusterConfHistory"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterConfigsHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterConfigsHistoryResponseParams `json:"Response"`
}

func (r *DescribeClusterConfigsHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterConfigsHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterConfigsRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	//  0 公有云查询；1青鹅查询，青鹅查询显示所有需要展示的
	ConfigType *int64 `json:"ConfigType,omitnil,omitempty" name:"ConfigType"`

	// 模糊搜索关键字文件
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 0集群维度 1节点维度
	ClusterConfigType *int64 `json:"ClusterConfigType,omitnil,omitempty" name:"ClusterConfigType"`

	// eth0的ip地址
	IPAddress *string `json:"IPAddress,omitnil,omitempty" name:"IPAddress"`
}

type DescribeClusterConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	//  0 公有云查询；1青鹅查询，青鹅查询显示所有需要展示的
	ConfigType *int64 `json:"ConfigType,omitnil,omitempty" name:"ConfigType"`

	// 模糊搜索关键字文件
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 0集群维度 1节点维度
	ClusterConfigType *int64 `json:"ClusterConfigType,omitnil,omitempty" name:"ClusterConfigType"`

	// eth0的ip地址
	IPAddress *string `json:"IPAddress,omitnil,omitempty" name:"IPAddress"`
}

func (r *DescribeClusterConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConfigType")
	delete(f, "FileName")
	delete(f, "ClusterConfigType")
	delete(f, "IPAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterConfigsResponseParams struct {
	// 返回实例的配置文件相关的信息
	ClusterConfList []*ClusterConfigsInfoFromEMR `json:"ClusterConfList,omitnil,omitempty" name:"ClusterConfList"`

	// 返回当前内核版本 如果不存在则返回空字符串
	BuildVersion *string `json:"BuildVersion,omitnil,omitempty" name:"BuildVersion"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterConfigsResponseParams `json:"Response"`
}

func (r *DescribeClusterConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCoolDownBackendsRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeCoolDownBackendsRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeCoolDownBackendsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCoolDownBackendsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCoolDownBackendsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCoolDownBackendsResponseParams struct {
	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 节点信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*CoolDownBackend `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCoolDownBackendsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCoolDownBackendsResponseParams `json:"Response"`
}

func (r *DescribeCoolDownBackendsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCoolDownBackendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCoolDownPoliciesRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeCoolDownPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeCoolDownPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCoolDownPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCoolDownPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCoolDownPoliciesResponseParams struct {
	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 冷热分层策略列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*CoolDownPolicyInfo `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCoolDownPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCoolDownPoliciesResponseParams `json:"Response"`
}

func (r *DescribeCoolDownPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCoolDownPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCoolDownTableDataRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`
}

type DescribeCoolDownTableDataRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`
}

func (r *DescribeCoolDownTableDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCoolDownTableDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DatabaseName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCoolDownTableDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCoolDownTableDataResponseParams struct {
	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 冷热分层Table数据列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*CoolDownTableDataInfo `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCoolDownTableDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCoolDownTableDataResponseParams `json:"Response"`
}

func (r *DescribeCoolDownTableDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCoolDownTableDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseAuditDownloadRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 排序参数
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 数据库
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// sql类型
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// sql语句
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 用户 多选
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// 数据库 多选
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`

	// sql类型 多选
	SqlTypes []*string `json:"SqlTypes,omitnil,omitempty" name:"SqlTypes"`

	// catalog名称 （多选）
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`
}

type DescribeDatabaseAuditDownloadRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 排序参数
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 数据库
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// sql类型
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// sql语句
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 用户 多选
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// 数据库 多选
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`

	// sql类型 多选
	SqlTypes []*string `json:"SqlTypes,omitnil,omitempty" name:"SqlTypes"`

	// catalog名称 （多选）
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`
}

func (r *DescribeDatabaseAuditDownloadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseAuditDownloadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "OrderType")
	delete(f, "User")
	delete(f, "DbName")
	delete(f, "SqlType")
	delete(f, "Sql")
	delete(f, "Users")
	delete(f, "DbNames")
	delete(f, "SqlTypes")
	delete(f, "Catalogs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseAuditDownloadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseAuditDownloadResponseParams struct {
	// 日志的cos地址
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabaseAuditDownloadResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseAuditDownloadResponseParams `json:"Response"`
}

func (r *DescribeDatabaseAuditDownloadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseAuditDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseAuditRecordsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 排序参数
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 数据库
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// sql类型
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// sql语句
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 用户 （多选）
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// 数据库 （多选）
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`

	// sql类型 （多选）
	SqlTypes []*string `json:"SqlTypes,omitnil,omitempty" name:"SqlTypes"`

	// catalog名称（多选）
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`
}

type DescribeDatabaseAuditRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 排序参数
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 数据库
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// sql类型
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// sql语句
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 用户 （多选）
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// 数据库 （多选）
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`

	// sql类型 （多选）
	SqlTypes []*string `json:"SqlTypes,omitnil,omitempty" name:"SqlTypes"`

	// catalog名称（多选）
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`
}

func (r *DescribeDatabaseAuditRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseAuditRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "OrderType")
	delete(f, "User")
	delete(f, "DbName")
	delete(f, "SqlType")
	delete(f, "Sql")
	delete(f, "Users")
	delete(f, "DbNames")
	delete(f, "SqlTypes")
	delete(f, "Catalogs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseAuditRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseAuditRecordsResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 记录列表
	SlowQueryRecords *DataBaseAuditRecord `json:"SlowQueryRecords,omitnil,omitempty" name:"SlowQueryRecords"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabaseAuditRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseAuditRecordsResponseParams `json:"Response"`
}

func (r *DescribeDatabaseAuditRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseAuditRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesInfoRequestParams struct {
	// 集群id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type DescribeInstanceNodesInfoRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

func (r *DescribeInstanceNodesInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceNodesInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesInfoResponseParams struct {
	// Be节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeNodes []*string `json:"BeNodes,omitnil,omitempty" name:"BeNodes"`

	// Fe节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeNodes []*string `json:"FeNodes,omitnil,omitempty" name:"FeNodes"`

	// Fe master节点
	FeMaster *string `json:"FeMaster,omitnil,omitempty" name:"FeMaster"`

	// Be节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeNodeInfos []*NodeInfo `json:"BeNodeInfos,omitnil,omitempty" name:"BeNodeInfos"`

	// Fe节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeNodeInfos []*NodeInfo `json:"FeNodeInfos,omitnil,omitempty" name:"FeNodeInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceNodesInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceNodesInfoResponseParams `json:"Response"`
}

func (r *DescribeInstanceNodesInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群角色类型，默认为 "data"数据节点
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 展现策略，All时显示所有
	DisplayPolicy *string `json:"DisplayPolicy,omitnil,omitempty" name:"DisplayPolicy"`
}

type DescribeInstanceNodesRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群角色类型，默认为 "data"数据节点
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 展现策略，All时显示所有
	DisplayPolicy *string `json:"DisplayPolicy,omitnil,omitempty" name:"DisplayPolicy"`
}

func (r *DescribeInstanceNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeRole")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DisplayPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例节点总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNodesList []*InstanceNode `json:"InstanceNodesList,omitnil,omitempty" name:"InstanceNodesList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceNodesResponseParams `json:"Response"`
}

func (r *DescribeInstanceNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesRoleRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 过滤ip
	IpFilter *string `json:"IpFilter,omitnil,omitempty" name:"IpFilter"`
}

type DescribeInstanceNodesRoleRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 过滤ip
	IpFilter *string `json:"IpFilter,omitnil,omitempty" name:"IpFilter"`
}

func (r *DescribeInstanceNodesRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IpFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceNodesRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesRoleResponseParams struct {
	// 错误码
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 节点总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 无
	NodeInfos []*NodeInfos `json:"NodeInfos,omitnil,omitempty" name:"NodeInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceNodesRoleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceNodesRoleResponseParams `json:"Response"`
}

func (r *DescribeInstanceNodesRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceOperationsRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页参数，偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，每页数目，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeInstanceOperationsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页参数，偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，每页数目，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeInstanceOperationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceOperationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceOperationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceOperationsResponseParams struct {
	// 操作记录总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 操作记录具体数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operations []*InstanceOperation `json:"Operations,omitnil,omitempty" name:"Operations"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceOperationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceOperationsResponseParams `json:"Response"`
}

func (r *DescribeInstanceOperationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceOperationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceResponseParams struct {
	// 实例描述信息
	InstanceInfo *InstanceInfo `json:"InstanceInfo,omitnil,omitempty" name:"InstanceInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceResponseParams `json:"Response"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceStateRequestParams struct {
	// 集群实例名称
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceStateRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例名称
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceStateResponseParams struct {
	// 集群状态，例如：Serving
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// 集群操作创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowCreateTime *string `json:"FlowCreateTime,omitnil,omitempty" name:"FlowCreateTime"`

	// 集群操作名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 集群操作进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowProgress *float64 `json:"FlowProgress,omitnil,omitempty" name:"FlowProgress"`

	// 集群状态描述，例如：运行中
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStateDesc *string `json:"InstanceStateDesc,omitnil,omitempty" name:"InstanceStateDesc"`

	// 集群流程错误信息，例如：“创建失败，资源不足”
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMsg *string `json:"FlowMsg,omitnil,omitempty" name:"FlowMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceStateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceStateResponseParams `json:"Response"`
}

func (r *DescribeInstanceStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceUsedSubnetsRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceUsedSubnetsRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceUsedSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceUsedSubnetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceUsedSubnetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceUsedSubnetsResponseParams struct {
	// 集群使用的vpc信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 集群使用的subnet信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedSubnets []*string `json:"UsedSubnets,omitnil,omitempty" name:"UsedSubnets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceUsedSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceUsedSubnetsResponseParams `json:"Response"`
}

func (r *DescribeInstanceUsedSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceUsedSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesHealthStateRequestParams struct {
	// 集群Id
	//
	// Deprecated: InstanceID is deprecated.
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 为空：代表当前appId下所有集群 或者  某个集群Id
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`
}

type DescribeInstancesHealthStateRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 为空：代表当前appId下所有集群 或者  某个集群Id
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`
}

func (r *DescribeInstancesHealthStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesHealthStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "Input")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesHealthStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesHealthStateResponseParams struct {
	// base64编码后的数据，包含了集群的健康信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesHealthStateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesHealthStateResponseParams `json:"Response"`
}

func (r *DescribeInstancesHealthStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesHealthStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 搜索的集群id名称
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// 搜索的集群name
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索标签列表
	SearchTags []*SearchTags `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 搜索的集群id名称
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// 搜索的集群name
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索标签列表
	SearchTags []*SearchTags `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
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
	delete(f, "SearchInstanceId")
	delete(f, "SearchInstanceName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 实例总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例数组
	InstancesList []*InstanceInfo `json:"InstancesList,omitnil,omitempty" name:"InstancesList"`

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
type DescribeRestoreTaskDetailRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
}

type DescribeRestoreTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
}

func (r *DescribeRestoreTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRestoreTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackUpJobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRestoreTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRestoreTaskDetailResponseParams struct {
	// 恢复任务进度详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestoreStatus []*RestoreStatus `json:"RestoreStatus,omitnil,omitempty" name:"RestoreStatus"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRestoreTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRestoreTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeRestoreTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRestoreTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryRecordsDownloadRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 慢查询时间
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil,omitempty" name:"QueryDurationMs"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序参数
	DurationMs *string `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// 查询sql
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 排序参数
	ReadRows *string `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// 排序参数
	ResultBytes *string `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// 排序参数
	MemoryUsage *string `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`

	// IsQuery条件
	IsQuery *int64 `json:"IsQuery,omitnil,omitempty" name:"IsQuery"`

	// 数据库名称
	DbName []*string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// catalog名称
	CatalogName []*string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`
}

type DescribeSlowQueryRecordsDownloadRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 慢查询时间
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil,omitempty" name:"QueryDurationMs"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序参数
	DurationMs *string `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// 查询sql
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 排序参数
	ReadRows *string `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// 排序参数
	ResultBytes *string `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// 排序参数
	MemoryUsage *string `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`

	// IsQuery条件
	IsQuery *int64 `json:"IsQuery,omitnil,omitempty" name:"IsQuery"`

	// 数据库名称
	DbName []*string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// catalog名称
	CatalogName []*string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`
}

func (r *DescribeSlowQueryRecordsDownloadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryRecordsDownloadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "QueryDurationMs")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DurationMs")
	delete(f, "Sql")
	delete(f, "ReadRows")
	delete(f, "ResultBytes")
	delete(f, "MemoryUsage")
	delete(f, "IsQuery")
	delete(f, "DbName")
	delete(f, "CatalogName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowQueryRecordsDownloadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryRecordsDownloadResponseParams struct {
	// cos地址
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowQueryRecordsDownloadResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowQueryRecordsDownloadResponseParams `json:"Response"`
}

func (r *DescribeSlowQueryRecordsDownloadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryRecordsDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryRecordsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 慢查询时间
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil,omitempty" name:"QueryDurationMs"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 排序参数
	DurationMs *string `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// 数据库名称
	DbName []*string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 是否是查询，0：否， 1：是
	IsQuery *int64 `json:"IsQuery,omitnil,omitempty" name:"IsQuery"`

	// catalog名称
	CatalogName []*string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// sql名
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// ReadRows排序字段
	ReadRows *string `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// ResultBytes排序字段
	ResultBytes *string `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// MemoryUsage排序字段
	MemoryUsage *string `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`
}

type DescribeSlowQueryRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 慢查询时间
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil,omitempty" name:"QueryDurationMs"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 排序参数
	DurationMs *string `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// 数据库名称
	DbName []*string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 是否是查询，0：否， 1：是
	IsQuery *int64 `json:"IsQuery,omitnil,omitempty" name:"IsQuery"`

	// catalog名称
	CatalogName []*string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// sql名
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// ReadRows排序字段
	ReadRows *string `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// ResultBytes排序字段
	ResultBytes *string `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// MemoryUsage排序字段
	MemoryUsage *string `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`
}

func (r *DescribeSlowQueryRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "QueryDurationMs")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "DurationMs")
	delete(f, "DbName")
	delete(f, "IsQuery")
	delete(f, "CatalogName")
	delete(f, "Sql")
	delete(f, "ReadRows")
	delete(f, "ResultBytes")
	delete(f, "MemoryUsage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowQueryRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryRecordsResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 记录列表
	SlowQueryRecords []*SlowQueryRecord `json:"SlowQueryRecords,omitnil,omitempty" name:"SlowQueryRecords"`

	// 所有数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBNameList []*string `json:"DBNameList,omitnil,omitempty" name:"DBNameList"`

	// 所有catalog名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CatalogNameList []*string `json:"CatalogNameList,omitnil,omitempty" name:"CatalogNameList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowQueryRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowQueryRecordsResponseParams `json:"Response"`
}

func (r *DescribeSlowQueryRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecRequestParams struct {
	// 地域信息，例如"ap-guangzhou-1"
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 计费类型，PREPAID 包年包月，POSTPAID_BY_HOUR 按量计费
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 多可用区
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 机型名称
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`
}

type DescribeSpecRequest struct {
	*tchttp.BaseRequest
	
	// 地域信息，例如"ap-guangzhou-1"
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 计费类型，PREPAID 包年包月，POSTPAID_BY_HOUR 按量计费
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 多可用区
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 机型名称
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`
}

func (r *DescribeSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "PayMode")
	delete(f, "Zones")
	delete(f, "SpecName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecResponseParams struct {
	// zookeeper节点规格描述
	MasterSpec []*ResourceSpec `json:"MasterSpec,omitnil,omitempty" name:"MasterSpec"`

	// 数据节点规格描述
	CoreSpec []*ResourceSpec `json:"CoreSpec,omitnil,omitempty" name:"CoreSpec"`

	// 云盘列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachCBSSpec []*DiskSpec `json:"AttachCBSSpec,omitnil,omitempty" name:"AttachCBSSpec"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSpecResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpecResponseParams `json:"Response"`
}

func (r *DescribeSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSqlApisRequestParams struct {
	// 用户链接来自的 IP
	WhiteHost *string `json:"WhiteHost,omitnil,omitempty" name:"WhiteHost"`

	// catalog名称
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// catalog集合
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`

	// 数据库名
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

type DescribeSqlApisRequest struct {
	*tchttp.BaseRequest
	
	// 用户链接来自的 IP
	WhiteHost *string `json:"WhiteHost,omitnil,omitempty" name:"WhiteHost"`

	// catalog名称
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// catalog集合
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`

	// 数据库名
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

func (r *DescribeSqlApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSqlApisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WhiteHost")
	delete(f, "Catalog")
	delete(f, "Catalogs")
	delete(f, "DatabaseName")
	delete(f, "TableName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSqlApisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSqlApisResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSqlApisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSqlApisResponseParams `json:"Response"`
}

func (r *DescribeSqlApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSqlApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableListRequestParams struct {
	// 资源ID，建表所用的TCHouse-D资源ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要获取表列表的库
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 使用该用户进行操作，该用户需要有对应的权限。如果该TCHouse-D集群使用CAM用户注册内核账户，则不需要填写
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户对应的密码。如果该TCHouse-D集群使用CAM用户注册内核账户，则不需要填写
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// 查询库所在的数据源，不填则默认为内部数据源（internal）。
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`
}

type DescribeTableListRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID，建表所用的TCHouse-D资源ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要获取表列表的库
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 使用该用户进行操作，该用户需要有对应的权限。如果该TCHouse-D集群使用CAM用户注册内核账户，则不需要填写
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户对应的密码。如果该TCHouse-D集群使用CAM用户注册内核账户，则不需要填写
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// 查询库所在的数据源，不填则默认为内部数据源（internal）。
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`
}

func (r *DescribeTableListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DbName")
	delete(f, "UserName")
	delete(f, "PassWord")
	delete(f, "CatalogName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableListResponseParams struct {
	// 表名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableNames []*string `json:"TableNames,omitnil,omitempty" name:"TableNames"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTableListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableListResponseParams `json:"Response"`
}

func (r *DescribeTableListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserBindWorkloadGroupRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeUserBindWorkloadGroupRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeUserBindWorkloadGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserBindWorkloadGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserBindWorkloadGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserBindWorkloadGroupResponseParams struct {
	// 用户绑定资源组信息
	UserBindInfos []*UserWorkloadGroup `json:"UserBindInfos,omitnil,omitempty" name:"UserBindInfos"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserBindWorkloadGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserBindWorkloadGroupResponseParams `json:"Response"`
}

func (r *DescribeUserBindWorkloadGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserBindWorkloadGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkloadGroupRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeWorkloadGroupRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeWorkloadGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkloadGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkloadGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkloadGroupResponseParams struct {
	// 资源组信息
	WorkloadGroups []*WorkloadGroupConfig `json:"WorkloadGroups,omitnil,omitempty" name:"WorkloadGroups"`

	// 是否开启资源组：开启-open、关闭-close
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkloadGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkloadGroupResponseParams `json:"Response"`
}

func (r *DescribeWorkloadGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkloadGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstanceRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DestroyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DestroyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstanceResponseParams struct {
	// 流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyInstanceResponseParams `json:"Response"`
}

func (r *DestroyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskSpec struct {
	// 磁盘类型，例如“CLOUD_SSD", "LOCAL_SSD"等
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 磁盘类型说明，例如"云SSD", "本地SSD"等
	DiskDesc *string `json:"DiskDesc,omitnil,omitempty" name:"DiskDesc"`

	// 磁盘最小规格大小，单位G
	MinDiskSize *int64 `json:"MinDiskSize,omitnil,omitempty" name:"MinDiskSize"`

	// 磁盘最大规格大小，单位G
	MaxDiskSize *int64 `json:"MaxDiskSize,omitnil,omitempty" name:"MaxDiskSize"`

	// 磁盘数目
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`
}

type DorisSourceInfo struct {
	// doris集群的fe的ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// doris集群的fe的端口号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// doris集群的账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// doris集群的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type FrontEndRule struct {
	// id序列
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 详细规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`
}

type InstanceConfigItem struct {
	// key
	ConfKey *string `json:"ConfKey,omitnil,omitempty" name:"ConfKey"`

	// value
	ConfValue *string `json:"ConfValue,omitnil,omitempty" name:"ConfValue"`
}

type InstanceDetail struct {
	// 告警策略是否可用	
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableAlarmStrategy *bool `json:"EnableAlarmStrategy,omitnil,omitempty" name:"EnableAlarmStrategy"`
}

type InstanceInfo struct {
	// 集群实例ID, "cdw-xxxx" 字符串类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 状态,
	// Init 创建中; Serving 运行中； 
	// Deleted已销毁；Deleting 销毁中；
	// Modify 集群变更中；
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 地域, ap-guangzhou
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区， ap-guangzhou-3
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 私有网络名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 付费类型，"hour", "prepay"
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 数据节点描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterSummary *NodesSummary `json:"MasterSummary,omitnil,omitempty" name:"MasterSummary"`

	// zookeeper节点描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoreSummary *NodesSummary `json:"CoreSummary,omitnil,omitempty" name:"CoreSummary"`

	// 高可用，“true" "false"
	// 注意：此字段可能返回 null，表示取不到有效值。
	HA *string `json:"HA,omitnil,omitempty" name:"HA"`

	// 高可用类型：
	// 0：非高可用
	// 1：读高可用
	// 2：读写高可用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`

	// 访问地址，例如 "10.0.0.1:9000"
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessInfo *string `json:"AccessInfo,omitnil,omitempty" name:"AccessInfo"`

	// 记录ID，数值型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// regionId, 表示地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 可用区说明，例如 "广州二区"
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneDesc *string `json:"ZoneDesc,omitnil,omitempty" name:"ZoneDesc"`

	// 错误流程说明信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMsg *string `json:"FlowMsg,omitnil,omitempty" name:"FlowMsg"`

	// 状态描述，例如“运行中”等
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 自动续费标记
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *bool `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 监控信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Monitor *string `json:"Monitor,omitnil,omitempty" name:"Monitor"`

	// 是否开通日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasClsTopic *bool `json:"HasClsTopic,omitnil,omitempty" name:"HasClsTopic"`

	// 日志主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClsTopicId *string `json:"ClsTopicId,omitnil,omitempty" name:"ClsTopicId"`

	// 日志集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClsLogSetId *string `json:"ClsLogSetId,omitnil,omitempty" name:"ClsLogSetId"`

	// 是否支持xml配置管理
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableXMLConfig *int64 `json:"EnableXMLConfig,omitnil,omitempty" name:"EnableXMLConfig"`

	// 区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionDesc *string `json:"RegionDesc,omitnil,omitempty" name:"RegionDesc"`

	// 弹性网卡地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Eip *string `json:"Eip,omitnil,omitempty" name:"Eip"`

	// 冷热分层系数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosMoveFactor *int64 `json:"CosMoveFactor,omitnil,omitempty" name:"CosMoveFactor"`

	// external/local/yunti
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// cos桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosBucketName *string `json:"CosBucketName,omitnil,omitempty" name:"CosBucketName"`

	// cbs
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanAttachCbs *bool `json:"CanAttachCbs,omitnil,omitempty" name:"CanAttachCbs"`

	// 小版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuildVersion *string `json:"BuildVersion,omitnil,omitempty" name:"BuildVersion"`

	// 组件信息
	// 注：这里返回类型实际为map[string]struct类型，并非显示的string类型，可以参考“示例值”进行数据的解析。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Components *string `json:"Components,omitnil,omitempty" name:"Components"`

	// 判断审计日志表是否有catalog字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: IfExistCatalog is deprecated.
	IfExistCatalog *int64 `json:"IfExistCatalog,omitnil,omitempty" name:"IfExistCatalog"`

	// 页面特性，用于前端屏蔽一些页面入口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Characteristic []*string `json:"Characteristic,omitnil,omitempty" name:"Characteristic"`

	// 超时时间 单位s
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartTimeout *string `json:"RestartTimeout,omitnil,omitempty" name:"RestartTimeout"`

	// 内核优雅重启超时时间，如果为-1说明未设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	GraceShutdownWaitSeconds *string `json:"GraceShutdownWaitSeconds,omitnil,omitempty" name:"GraceShutdownWaitSeconds"`

	// 表名大小写是否敏感，0：敏感；1：不敏感，以小写进行比较；2：不敏感，表名改为以小写存储
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaseSensitive *int64 `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// 用户是否可以绑定安全组
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsWhiteSGs *bool `json:"IsWhiteSGs,omitnil,omitempty" name:"IsWhiteSGs"`

	// 已绑定的安全组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindSGs []*string `json:"BindSGs,omitnil,omitempty" name:"BindSGs"`

	// 是否为多可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableMultiZones *bool `json:"EnableMultiZones,omitnil,omitempty" name:"EnableMultiZones"`

	// 用户可用区和子网信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserNetworkInfos *string `json:"UserNetworkInfos,omitnil,omitempty" name:"UserNetworkInfos"`

	// 是否启用冷热分层。0：未开启 1：已开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableCoolDown *int64 `json:"EnableCoolDown,omitnil,omitempty" name:"EnableCoolDown"`

	// 冷热分层使用COS桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoolDownBucket *string `json:"CoolDownBucket,omitnil,omitempty" name:"CoolDownBucket"`

	// 实例扩展信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details *InstanceDetail `json:"Details,omitnil,omitempty" name:"Details"`

	// 是否启用DLC 0:关闭 1:开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableDlc *int64 `json:"EnableDlc,omitnil,omitempty" name:"EnableDlc"`

	// 账户类型 0:普通用户 1:CAM用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`
}

type InstanceNode struct {
	// IP地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 机型，如 S1
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// cpu核数
	Core *int64 `json:"Core,omitnil,omitempty" name:"Core"`

	// 内存大小
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 磁盘类型
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 磁盘大小
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 所属clickhouse cluster名称
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// rip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rip *string `json:"Rip,omitnil,omitempty" name:"Rip"`

	// FE节点角色
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeRole *string `json:"FeRole,omitnil,omitempty" name:"FeRole"`

	// UUID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UUID *string `json:"UUID,omitnil,omitempty" name:"UUID"`
}

type InstanceOperation struct {
	// 操作名称，例如“create_instance"、“scaleout_instance”等
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 操作结果，“Success"表示成功，”Fail"表示失败
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 操作名称描述，例如“创建”，“修改集群名称”等
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 操作级别，例如“Critical", "Normal"等
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 操作级别描述，例如“高危”，“一般”等
	LevelDesc *string `json:"LevelDesc,omitnil,omitempty" name:"LevelDesc"`

	// 操作开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 操作结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 操作结果描述，例如“成功”，“失败”
	ResultDesc *string `json:"ResultDesc,omitnil,omitempty" name:"ResultDesc"`

	// 操作用户ID
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// 操作对应的jobid
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 操作明细
	OperationDetail *string `json:"OperationDetail,omitnil,omitempty" name:"OperationDetail"`
}

// Predefined struct for user
type ModifyClusterConfigsRequestParams struct {
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件修改信息
	ModifyConfContext []*ConfigSubmitContext `json:"ModifyConfContext,omitnil,omitempty" name:"ModifyConfContext"`

	// 修改原因
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyClusterConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件修改信息
	ModifyConfContext []*ConfigSubmitContext `json:"ModifyConfContext,omitnil,omitempty" name:"ModifyConfContext"`

	// 修改原因
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyClusterConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ModifyConfContext")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterConfigsResponseParams struct {
	// 流程相关信息
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterConfigsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterConfigsResponseParams `json:"Response"`
}

func (r *ModifyClusterConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCoolDownPolicyRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// cooldown_ttl
	CoolDownTtl *string `json:"CoolDownTtl,omitnil,omitempty" name:"CoolDownTtl"`

	// cooldown_datetime
	CoolDownDatetime *string `json:"CoolDownDatetime,omitnil,omitempty" name:"CoolDownDatetime"`
}

type ModifyCoolDownPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// cooldown_ttl
	CoolDownTtl *string `json:"CoolDownTtl,omitnil,omitempty" name:"CoolDownTtl"`

	// cooldown_datetime
	CoolDownDatetime *string `json:"CoolDownDatetime,omitnil,omitempty" name:"CoolDownDatetime"`
}

func (r *ModifyCoolDownPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCoolDownPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PolicyName")
	delete(f, "CoolDownTtl")
	delete(f, "CoolDownDatetime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCoolDownPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCoolDownPolicyResponseParams struct {
	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCoolDownPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCoolDownPolicyResponseParams `json:"Response"`
}

func (r *ModifyCoolDownPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCoolDownPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceKeyValConfigsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 新增配置列表
	AddItems []*InstanceConfigItem `json:"AddItems,omitnil,omitempty" name:"AddItems"`

	// 更新配置列表
	UpdateItems []*InstanceConfigItem `json:"UpdateItems,omitnil,omitempty" name:"UpdateItems"`

	// 删除配置列表
	DelItems []*InstanceConfigItem `json:"DelItems,omitnil,omitempty" name:"DelItems"`

	// 备注（50字以内）
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 热更新列表
	HotUpdateItems []*InstanceConfigItem `json:"HotUpdateItems,omitnil,omitempty" name:"HotUpdateItems"`

	// 删除配置列表
	DeleteItems *InstanceConfigItem `json:"DeleteItems,omitnil,omitempty" name:"DeleteItems"`

	// ip地址
	IPAddress *string `json:"IPAddress,omitnil,omitempty" name:"IPAddress"`
}

type ModifyInstanceKeyValConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 新增配置列表
	AddItems []*InstanceConfigItem `json:"AddItems,omitnil,omitempty" name:"AddItems"`

	// 更新配置列表
	UpdateItems []*InstanceConfigItem `json:"UpdateItems,omitnil,omitempty" name:"UpdateItems"`

	// 删除配置列表
	DelItems []*InstanceConfigItem `json:"DelItems,omitnil,omitempty" name:"DelItems"`

	// 备注（50字以内）
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 热更新列表
	HotUpdateItems []*InstanceConfigItem `json:"HotUpdateItems,omitnil,omitempty" name:"HotUpdateItems"`

	// 删除配置列表
	DeleteItems *InstanceConfigItem `json:"DeleteItems,omitnil,omitempty" name:"DeleteItems"`

	// ip地址
	IPAddress *string `json:"IPAddress,omitnil,omitempty" name:"IPAddress"`
}

func (r *ModifyInstanceKeyValConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceKeyValConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FileName")
	delete(f, "AddItems")
	delete(f, "UpdateItems")
	delete(f, "DelItems")
	delete(f, "Message")
	delete(f, "HotUpdateItems")
	delete(f, "DeleteItems")
	delete(f, "IPAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceKeyValConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceKeyValConfigsResponseParams struct {
	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceKeyValConfigsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceKeyValConfigsResponseParams `json:"Response"`
}

func (r *ModifyInstanceKeyValConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceKeyValConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新修改的实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新修改的实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

func (r *ModifyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceResponseParams `json:"Response"`
}

func (r *ModifyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNodeStatusRequestParams struct {
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点信息
	NodeInfos []*NodeInfos `json:"NodeInfos,omitnil,omitempty" name:"NodeInfos"`

	// 节点操作
	OperationCode *string `json:"OperationCode,omitnil,omitempty" name:"OperationCode"`

	// 超时时间（秒）
	RestartTimeOut *string `json:"RestartTimeOut,omitnil,omitempty" name:"RestartTimeOut"`
}

type ModifyNodeStatusRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点信息
	NodeInfos []*NodeInfos `json:"NodeInfos,omitnil,omitempty" name:"NodeInfos"`

	// 节点操作
	OperationCode *string `json:"OperationCode,omitnil,omitempty" name:"OperationCode"`

	// 超时时间（秒）
	RestartTimeOut *string `json:"RestartTimeOut,omitnil,omitempty" name:"RestartTimeOut"`
}

func (r *ModifyNodeStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNodeStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeInfos")
	delete(f, "OperationCode")
	delete(f, "RestartTimeOut")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNodeStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNodeStatusResponseParams struct {
	// 流程相关信息
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNodeStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNodeStatusResponseParams `json:"Response"`
}

func (r *ModifyNodeStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNodeStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupsRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改前的安全组信息
	OldSecurityGroupIds []*string `json:"OldSecurityGroupIds,omitnil,omitempty" name:"OldSecurityGroupIds"`

	// 修改后的安全组信息
	ModifySecurityGroupIds []*string `json:"ModifySecurityGroupIds,omitnil,omitempty" name:"ModifySecurityGroupIds"`
}

type ModifySecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改前的安全组信息
	OldSecurityGroupIds []*string `json:"OldSecurityGroupIds,omitnil,omitempty" name:"OldSecurityGroupIds"`

	// 修改后的安全组信息
	ModifySecurityGroupIds []*string `json:"ModifySecurityGroupIds,omitnil,omitempty" name:"ModifySecurityGroupIds"`
}

func (r *ModifySecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OldSecurityGroupIds")
	delete(f, "ModifySecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupsResponseParams struct {
	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityGroupsResponseParams `json:"Response"`
}

func (r *ModifySecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserBindWorkloadGroupRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要绑定资源组的用户信息，如果一个账户拥有多个主机信息，需要将这些信息都传入
	BindUsers []*BindUser `json:"BindUsers,omitnil,omitempty" name:"BindUsers"`

	// 修改前绑定的资源组名称
	OldWorkloadGroupName *string `json:"OldWorkloadGroupName,omitnil,omitempty" name:"OldWorkloadGroupName"`

	// 修改后绑定的资源组名称
	NewWorkloadGroupName *string `json:"NewWorkloadGroupName,omitnil,omitempty" name:"NewWorkloadGroupName"`
}

type ModifyUserBindWorkloadGroupRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要绑定资源组的用户信息，如果一个账户拥有多个主机信息，需要将这些信息都传入
	BindUsers []*BindUser `json:"BindUsers,omitnil,omitempty" name:"BindUsers"`

	// 修改前绑定的资源组名称
	OldWorkloadGroupName *string `json:"OldWorkloadGroupName,omitnil,omitempty" name:"OldWorkloadGroupName"`

	// 修改后绑定的资源组名称
	NewWorkloadGroupName *string `json:"NewWorkloadGroupName,omitnil,omitempty" name:"NewWorkloadGroupName"`
}

func (r *ModifyUserBindWorkloadGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserBindWorkloadGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BindUsers")
	delete(f, "OldWorkloadGroupName")
	delete(f, "NewWorkloadGroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserBindWorkloadGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserBindWorkloadGroupResponseParams struct {
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserBindWorkloadGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserBindWorkloadGroupResponseParams `json:"Response"`
}

func (r *ModifyUserBindWorkloadGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserBindWorkloadGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserPrivilegesV3RequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户权限
	UserPrivileges *UpdateUserPrivileges `json:"UserPrivileges,omitnil,omitempty" name:"UserPrivileges"`

	// 用户链接来自的 IP	
	WhiteHost *string `json:"WhiteHost,omitnil,omitempty" name:"WhiteHost"`
}

type ModifyUserPrivilegesV3Request struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户权限
	UserPrivileges *UpdateUserPrivileges `json:"UserPrivileges,omitnil,omitempty" name:"UserPrivileges"`

	// 用户链接来自的 IP	
	WhiteHost *string `json:"WhiteHost,omitnil,omitempty" name:"WhiteHost"`
}

func (r *ModifyUserPrivilegesV3Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserPrivilegesV3Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "UserPrivileges")
	delete(f, "WhiteHost")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserPrivilegesV3Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserPrivilegesV3ResponseParams struct {
	// 错误信息，为空就是没有错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserPrivilegesV3Response struct {
	*tchttp.BaseResponse
	Response *ModifyUserPrivilegesV3ResponseParams `json:"Response"`
}

func (r *ModifyUserPrivilegesV3Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserPrivilegesV3Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkloadGroupRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改的资源组信息
	WorkloadGroup *WorkloadGroupConfig `json:"WorkloadGroup,omitnil,omitempty" name:"WorkloadGroup"`
}

type ModifyWorkloadGroupRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改的资源组信息
	WorkloadGroup *WorkloadGroupConfig `json:"WorkloadGroup,omitnil,omitempty" name:"WorkloadGroup"`
}

func (r *ModifyWorkloadGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkloadGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "WorkloadGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWorkloadGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkloadGroupResponseParams struct {
	// 错误信息	
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyWorkloadGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWorkloadGroupResponseParams `json:"Response"`
}

func (r *ModifyWorkloadGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkloadGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkloadGroupStatusRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否开启资源组：open-开启、close-关闭
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`
}

type ModifyWorkloadGroupStatusRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否开启资源组：open-开启、close-关闭
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`
}

func (r *ModifyWorkloadGroupStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkloadGroupStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OperationType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWorkloadGroupStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkloadGroupStatusResponseParams struct {
	// 错误信息	
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyWorkloadGroupStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWorkloadGroupStatusResponseParams `json:"Response"`
}

func (r *ModifyWorkloadGroupStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkloadGroupStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkInfo struct {
	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 当前子网可用ip数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetIpNum *int64 `json:"SubnetIpNum,omitnil,omitempty" name:"SubnetIpNum"`
}

type NodeInfo struct {
	// 用户IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 节点状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 节点角色名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 组件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentName *string `json:"ComponentName,omitnil,omitempty" name:"ComponentName"`

	// 节点角色
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// 节点上次重启的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastRestartTime *string `json:"LastRestartTime,omitnil,omitempty" name:"LastRestartTime"`

	// 节点所在可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type NodeInfos struct {
	// 节点名称
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 节点状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 节点ip
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 节点角色
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// 组件名
	ComponentName *string `json:"ComponentName,omitnil,omitempty" name:"ComponentName"`

	// 上一次重启时间
	LastRestartTime *string `json:"LastRestartTime,omitnil,omitempty" name:"LastRestartTime"`
}

type NodesSummary struct {
	// 机型，如 S1
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// 节点数目
	NodeSize *int64 `json:"NodeSize,omitnil,omitempty" name:"NodeSize"`

	// cpu核数，单位个
	Core *int64 `json:"Core,omitnil,omitempty" name:"Core"`

	// 内存大小，单位G
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 磁盘大小，单位G
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// 磁盘类型
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 磁盘描述
	DiskDesc *string `json:"DiskDesc,omitnil,omitempty" name:"DiskDesc"`

	// 挂载云盘信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachCBSSpec *AttachCBSSpec `json:"AttachCBSSpec,omitnil,omitempty" name:"AttachCBSSpec"`

	// 子产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductType *string `json:"SubProductType,omitnil,omitempty" name:"SubProductType"`

	// 规格核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecCore *int64 `json:"SpecCore,omitnil,omitempty" name:"SpecCore"`

	// 规格内存
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecMemory *int64 `json:"SpecMemory,omitnil,omitempty" name:"SpecMemory"`

	// 磁盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// 是否加密
	// 注意：此字段可能返回 null，表示取不到有效值。
	Encrypt *int64 `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`

	// 最大磁盘
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDiskSize *int64 `json:"MaxDiskSize,omitnil,omitempty" name:"MaxDiskSize"`
}

// Predefined struct for user
type OpenCoolDownPolicyRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// db名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// table名称
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 操作类型
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 逗号分隔 需要带上db的名字 db1.tb1,db1.tb2,db2.tb1
	BatchOpenCoolDownTables *string `json:"BatchOpenCoolDownTables,omitnil,omitempty" name:"BatchOpenCoolDownTables"`

	// 绑定的时候用 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 逗号分隔 p1,p2,p3
	BatchOpenCoolDownPartitions *string `json:"BatchOpenCoolDownPartitions,omitnil,omitempty" name:"BatchOpenCoolDownPartitions"`
}

type OpenCoolDownPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// db名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// table名称
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 操作类型
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 逗号分隔 需要带上db的名字 db1.tb1,db1.tb2,db2.tb1
	BatchOpenCoolDownTables *string `json:"BatchOpenCoolDownTables,omitnil,omitempty" name:"BatchOpenCoolDownTables"`

	// 绑定的时候用 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 逗号分隔 p1,p2,p3
	BatchOpenCoolDownPartitions *string `json:"BatchOpenCoolDownPartitions,omitnil,omitempty" name:"BatchOpenCoolDownPartitions"`
}

func (r *OpenCoolDownPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenCoolDownPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DatabaseName")
	delete(f, "TableName")
	delete(f, "OperationType")
	delete(f, "BatchOpenCoolDownTables")
	delete(f, "PolicyName")
	delete(f, "BatchOpenCoolDownPartitions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenCoolDownPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenCoolDownPolicyResponseParams struct {
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 返回信息
	QueryDocument *string `json:"QueryDocument,omitnil,omitempty" name:"QueryDocument"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenCoolDownPolicyResponse struct {
	*tchttp.BaseResponse
	Response *OpenCoolDownPolicyResponseParams `json:"Response"`
}

func (r *OpenCoolDownPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenCoolDownPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenCoolDownRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type OpenCoolDownRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *OpenCoolDownRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenCoolDownRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenCoolDownRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenCoolDownResponseParams struct {
	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenCoolDownResponse struct {
	*tchttp.BaseResponse
	Response *OpenCoolDownResponseParams `json:"Response"`
}

func (r *OpenCoolDownResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenCoolDownResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverBackUpJobRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`

	// 恢复出来的新表副本数
	ReplicationNum *int64 `json:"ReplicationNum,omitnil,omitempty" name:"ReplicationNum"`

	// 恢复是否保持源表中的配置，1时表示保留源表中的配置
	ReserveSourceConfig *int64 `json:"ReserveSourceConfig,omitnil,omitempty" name:"ReserveSourceConfig"`

	// 0默认 1cos恢复
	RecoverType *int64 `json:"RecoverType,omitnil,omitempty" name:"RecoverType"`

	// CosSourceInfo对象
	CosSourceInfo *CosSourceInfo `json:"CosSourceInfo,omitnil,omitempty" name:"CosSourceInfo"`

	// 0默认 1定期执行
	ScheduleType *int64 `json:"ScheduleType,omitnil,omitempty" name:"ScheduleType"`

	// 年-月-日 时:分:秒
	NextTime *string `json:"NextTime,omitnil,omitempty" name:"NextTime"`

	// 调度名称
	ScheduleName *string `json:"ScheduleName,omitnil,omitempty" name:"ScheduleName"`

	// create update
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 恢复粒度：All全量、Database按库、Table按表
	RecoverScope *string `json:"RecoverScope,omitnil,omitempty" name:"RecoverScope"`

	// 恢复库：如果是按库备份，则需要该字段，库之间用","分割
	RecoverDatabase *string `json:"RecoverDatabase,omitnil,omitempty" name:"RecoverDatabase"`
}

type RecoverBackUpJobRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`

	// 恢复出来的新表副本数
	ReplicationNum *int64 `json:"ReplicationNum,omitnil,omitempty" name:"ReplicationNum"`

	// 恢复是否保持源表中的配置，1时表示保留源表中的配置
	ReserveSourceConfig *int64 `json:"ReserveSourceConfig,omitnil,omitempty" name:"ReserveSourceConfig"`

	// 0默认 1cos恢复
	RecoverType *int64 `json:"RecoverType,omitnil,omitempty" name:"RecoverType"`

	// CosSourceInfo对象
	CosSourceInfo *CosSourceInfo `json:"CosSourceInfo,omitnil,omitempty" name:"CosSourceInfo"`

	// 0默认 1定期执行
	ScheduleType *int64 `json:"ScheduleType,omitnil,omitempty" name:"ScheduleType"`

	// 年-月-日 时:分:秒
	NextTime *string `json:"NextTime,omitnil,omitempty" name:"NextTime"`

	// 调度名称
	ScheduleName *string `json:"ScheduleName,omitnil,omitempty" name:"ScheduleName"`

	// create update
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 恢复粒度：All全量、Database按库、Table按表
	RecoverScope *string `json:"RecoverScope,omitnil,omitempty" name:"RecoverScope"`

	// 恢复库：如果是按库备份，则需要该字段，库之间用","分割
	RecoverDatabase *string `json:"RecoverDatabase,omitnil,omitempty" name:"RecoverDatabase"`
}

func (r *RecoverBackUpJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverBackUpJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackUpJobId")
	delete(f, "ReplicationNum")
	delete(f, "ReserveSourceConfig")
	delete(f, "RecoverType")
	delete(f, "CosSourceInfo")
	delete(f, "ScheduleType")
	delete(f, "NextTime")
	delete(f, "ScheduleName")
	delete(f, "OperationType")
	delete(f, "RecoverScope")
	delete(f, "RecoverDatabase")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverBackUpJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverBackUpJobResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecoverBackUpJobResponse struct {
	*tchttp.BaseResponse
	Response *RecoverBackUpJobResponseParams `json:"Response"`
}

func (r *RecoverBackUpJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverBackUpJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReduceInstanceRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点列表
	DelHosts []*string `json:"DelHosts,omitnil,omitempty" name:"DelHosts"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 缩容后集群高可用类型：0：非高可用，1：读高可用，2：读写高可用。
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`
}

type ReduceInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点列表
	DelHosts []*string `json:"DelHosts,omitnil,omitempty" name:"DelHosts"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 缩容后集群高可用类型：0：非高可用，1：读高可用，2：读写高可用。
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`
}

func (r *ReduceInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReduceInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DelHosts")
	delete(f, "Type")
	delete(f, "HaType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReduceInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReduceInstanceResponseParams struct {
	// 流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReduceInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ReduceInstanceResponseParams `json:"Response"`
}

func (r *ReduceInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReduceInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionAreaInfo struct {
	// 大类地域信息，如"south_china", "east_china"等
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 对应Name的描述，例如“华南地区”，“华东地区”等
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 具体地域列表1
	Regions []*RegionInfo `json:"Regions,omitnil,omitempty" name:"Regions"`
}

type RegionInfo struct {
	// 地域名称，例如“ap-guangzhou"
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 地域描述，例如"广州”
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 地域唯一标记
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域下所有可用区列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zones []*ZoneInfo `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 该地域下集群数目
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 0代表是国际站 1代表不是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsInternationalSite *uint64 `json:"IsInternationalSite,omitnil,omitempty" name:"IsInternationalSite"`

	// 桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`
}

// Predefined struct for user
type ResizeDiskRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 云盘大小
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

type ResizeDiskRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 云盘大小
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
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
	delete(f, "InstanceId")
	delete(f, "Type")
	delete(f, "DiskSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResizeDiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeDiskResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

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

type ResourceSpec struct {
	// 规格名称，例如“SCH1"
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// cpu核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存大小，单位G
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 分类标记，STANDARD/BIGDATA/HIGHIO分别表示标准型/大数据型/高IO
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 系统盘描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemDisk *DiskSpec `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 数据盘描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataDisk *DiskSpec `json:"DataDisk,omitnil,omitempty" name:"DataDisk"`

	// 最大节点数目限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxNodeSize *int64 `json:"MaxNodeSize,omitnil,omitempty" name:"MaxNodeSize"`

	// 是否可用，false代表售罄
	// 注意：此字段可能返回 null，表示取不到有效值。
	Available *bool `json:"Available,omitnil,omitempty" name:"Available"`

	// 规格描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComputeSpecDesc *string `json:"ComputeSpecDesc,omitnil,omitempty" name:"ComputeSpecDesc"`

	// cvm库存
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceQuota *int64 `json:"InstanceQuota,omitnil,omitempty" name:"InstanceQuota"`
}

// Predefined struct for user
type RestartClusterForConfigsRequestParams struct {
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件名称
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// grace_restart为优雅滚动重启 不填默认立刻重启
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`
}

type RestartClusterForConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件名称
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// grace_restart为优雅滚动重启 不填默认立刻重启
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`
}

func (r *RestartClusterForConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartClusterForConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConfigName")
	delete(f, "OperationType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartClusterForConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartClusterForConfigsResponseParams struct {
	// 流程相关信息
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartClusterForConfigsResponse struct {
	*tchttp.BaseResponse
	Response *RestartClusterForConfigsResponseParams `json:"Response"`
}

func (r *RestartClusterForConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartClusterForConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartClusterForNodeRequestParams struct {
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件名称
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// 每次重启的批次
	BatchSize *int64 `json:"BatchSize,omitnil,omitempty" name:"BatchSize"`

	// 重启节点
	NodeList []*string `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// false表示非滚动重启，true表示滚动重启
	RollingRestart *bool `json:"RollingRestart,omitnil,omitempty" name:"RollingRestart"`
}

type RestartClusterForNodeRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件名称
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// 每次重启的批次
	BatchSize *int64 `json:"BatchSize,omitnil,omitempty" name:"BatchSize"`

	// 重启节点
	NodeList []*string `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// false表示非滚动重启，true表示滚动重启
	RollingRestart *bool `json:"RollingRestart,omitnil,omitempty" name:"RollingRestart"`
}

func (r *RestartClusterForNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartClusterForNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConfigName")
	delete(f, "BatchSize")
	delete(f, "NodeList")
	delete(f, "RollingRestart")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartClusterForNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartClusterForNodeResponseParams struct {
	// 流程相关信息
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartClusterForNodeResponse struct {
	*tchttp.BaseResponse
	Response *RestartClusterForNodeResponseParams `json:"Response"`
}

func (r *RestartClusterForNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartClusterForNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreStatus struct {
	// 恢复任务id
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 恢复任务标签
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 恢复任务时间戳
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 恢复任务所在库
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 恢复任务状态
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 恢复时是否允许导入
	AllowLoad *bool `json:"AllowLoad,omitnil,omitempty" name:"AllowLoad"`

	// 副本数
	ReplicationNum *string `json:"ReplicationNum,omitnil,omitempty" name:"ReplicationNum"`

	// 副本数
	ReplicaAllocation *string `json:"ReplicaAllocation,omitnil,omitempty" name:"ReplicaAllocation"`

	// 恢复对象
	RestoreObjects *string `json:"RestoreObjects,omitnil,omitempty" name:"RestoreObjects"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 元数据准备时间
	MetaPreparedTime *string `json:"MetaPreparedTime,omitnil,omitempty" name:"MetaPreparedTime"`

	// 快照结束时间
	SnapshotFinishedTime *string `json:"SnapshotFinishedTime,omitnil,omitempty" name:"SnapshotFinishedTime"`

	// 下载结束时间
	DownloadFinishedTime *string `json:"DownloadFinishedTime,omitnil,omitempty" name:"DownloadFinishedTime"`

	// 结束时间
	FinishedTime *string `json:"FinishedTime,omitnil,omitempty" name:"FinishedTime"`

	// 未完成任务
	UnfinishedTasks *string `json:"UnfinishedTasks,omitnil,omitempty" name:"UnfinishedTasks"`

	// 进度
	Progress *string `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 错误信息
	TaskErrMsg *string `json:"TaskErrMsg,omitnil,omitempty" name:"TaskErrMsg"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 作业超时时间
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 是否保持源表中的副本数
	ReserveReplica *bool `json:"ReserveReplica,omitnil,omitempty" name:"ReserveReplica"`

	// 是否保持源表中的动态分区
	ReserveDynamicPartitionEnable *bool `json:"ReserveDynamicPartitionEnable,omitnil,omitempty" name:"ReserveDynamicPartitionEnable"`

	// 备份实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupJobId *int64 `json:"BackupJobId,omitnil,omitempty" name:"BackupJobId"`

	// 实例对应snapshot的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type ScaleOutInstanceRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 节点数量
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 扩容后集群高可用类型：0：非高可用，1：读高可用，2：读写高可用。
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`
}

type ScaleOutInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 节点数量
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 扩容后集群高可用类型：0：非高可用，1：读高可用，2：读写高可用。
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`
}

func (r *ScaleOutInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	delete(f, "NodeCount")
	delete(f, "HaType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleOutInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceResponseParams struct {
	// 流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleOutInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ScaleOutInstanceResponseParams `json:"Response"`
}

func (r *ScaleOutInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleUpInstanceRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点规格
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ScaleUpInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点规格
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *ScaleUpInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleUpInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpecName")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleUpInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleUpInstanceResponseParams struct {
	// 流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleUpInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ScaleUpInstanceResponseParams `json:"Response"`
}

func (r *ScaleUpInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleUpInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchTags struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`

	// 1表示只输入标签的键，没有输入值；0表示输入键时且输入值
	AllValue *int64 `json:"AllValue,omitnil,omitempty" name:"AllValue"`
}

type SlowQueryRecord struct {
	// 查询用户
	OsUser *string `json:"OsUser,omitnil,omitempty" name:"OsUser"`

	// 查询ID
	InitialQueryId *string `json:"InitialQueryId,omitnil,omitempty" name:"InitialQueryId"`

	// SQL语句
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 开始时间
	QueryStartTime *string `json:"QueryStartTime,omitnil,omitempty" name:"QueryStartTime"`

	// 执行耗时
	DurationMs *int64 `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// 读取行数
	ReadRows *int64 `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// 读取字节数
	ResultRows *int64 `json:"ResultRows,omitnil,omitempty" name:"ResultRows"`

	// 结果字节数
	ResultBytes *uint64 `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// 内存
	MemoryUsage *int64 `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`

	// 初始查询IP
	InitialAddress *string `json:"InitialAddress,omitnil,omitempty" name:"InitialAddress"`

	// 数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 是否是查询，0：否，1：查询语句
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsQuery *int64 `json:"IsQuery,omitnil,omitempty" name:"IsQuery"`

	// ResultBytes的MB格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultBytesMB *float64 `json:"ResultBytesMB,omitnil,omitempty" name:"ResultBytesMB"`

	// MemoryUsage的MB表示
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryUsageMB *float64 `json:"MemoryUsageMB,omitnil,omitempty" name:"MemoryUsageMB"`

	// DurationMs的秒表示
	// 注意：此字段可能返回 null，表示取不到有效值。
	DurationSec *float64 `json:"DurationSec,omitnil,omitempty" name:"DurationSec"`
}

type Tag struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type UpdateCoolDownRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否启用 0：不启用 1：启用
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 用户存放冷热分层数据Cos桶地址
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`
}

type UpdateCoolDownRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否启用 0：不启用 1：启用
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 用户存放冷热分层数据Cos桶地址
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`
}

func (r *UpdateCoolDownRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCoolDownRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Enable")
	delete(f, "Bucket")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCoolDownRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCoolDownResponseParams struct {
	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCoolDownResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCoolDownResponseParams `json:"Response"`
}

func (r *UpdateCoolDownResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCoolDownResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateUserPrivileges struct {
	// 是否设置catalog权限
	IsSetGlobalCatalog *bool `json:"IsSetGlobalCatalog,omitnil,omitempty" name:"IsSetGlobalCatalog"`
}

type UserInfo struct {
	// 集群实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 密码
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// 用户链接来自的 IP
	WhiteHost *string `json:"WhiteHost,omitnil,omitempty" name:"WhiteHost"`

	// 修改前用户链接来自的 IP
	OldWhiteHost *string `json:"OldWhiteHost,omitnil,omitempty" name:"OldWhiteHost"`

	// 描述
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// 旧密码
	OldPwd *string `json:"OldPwd,omitnil,omitempty" name:"OldPwd"`

	// 绑定的子用户uin
	CamUin *string `json:"CamUin,omitnil,omitempty" name:"CamUin"`

	// ranger group id列表
	CamRangerGroupIds []*int64 `json:"CamRangerGroupIds,omitnil,omitempty" name:"CamRangerGroupIds"`
}

type UserWorkloadGroup struct {
	// test
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// normal
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkloadGroupName *string `json:"WorkloadGroupName,omitnil,omitempty" name:"WorkloadGroupName"`
}

type WorkloadGroupConfig struct {
	// 资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkloadGroupName *string `json:"WorkloadGroupName,omitnil,omitempty" name:"WorkloadGroupName"`

	// CPU权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuShare *int64 `json:"CpuShare,omitnil,omitempty" name:"CpuShare"`

	// 内存限制，所有资源组的内存限制值之和应该小于等于100
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryLimit *int64 `json:"MemoryLimit,omitnil,omitempty" name:"MemoryLimit"`

	// 是否允许超配分配
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableMemoryOverCommit *bool `json:"EnableMemoryOverCommit,omitnil,omitempty" name:"EnableMemoryOverCommit"`

	// cpu硬限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuHardLimit *string `json:"CpuHardLimit,omitnil,omitempty" name:"CpuHardLimit"`
}

type ZoneInfo struct {
	// 可用区名称，例如"ap-guangzhou-1"
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 可用区描述信息，例如“广州一区”
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 可用区唯一标记
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Encryptid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Encrypt *int64 `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`
}