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

package v20211122

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AnalysisInstanceInfo struct {
	// <p>副本数</p>
	ReplicasNum *uint64 `json:"ReplicasNum,omitnil,omitempty" name:"ReplicasNum"`
}

type AnalysisRelationInfo struct {
	// <p>源实例Id</p>
	PrimaryInstanceId *string `json:"PrimaryInstanceId,omitnil,omitempty" name:"PrimaryInstanceId"`

	// <p>分析引擎实例Id</p>
	AnalysisInstanceId *string `json:"AnalysisInstanceId,omitnil,omitempty" name:"AnalysisInstanceId"`

	// <p>分析引擎关系状态</p><p>枚举值：</p><ul><li>creating： 创建中</li><li>running： 运行中</li><li>destroyed： 已销毁</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>创建时间</p>
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// <p>更新时间</p>
	UpdateAt *string `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`
}

type ArchiveLogInterval struct {
	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 大版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	MajorVersion *string `json:"MajorVersion,omitnil,omitempty" name:"MajorVersion"`

	// 小版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinorVersion *string `json:"MinorVersion,omitnil,omitempty" name:"MinorVersion"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type ArchiveLogModel struct {
	// 归档日志ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ArchiveLogId *int64 `json:"ArchiveLogId,omitnil,omitempty" name:"ArchiveLogId"`

	// 备份耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupDuration *int64 `json:"BackupDuration,omitnil,omitempty" name:"BackupDuration"`

	// 备份集状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupStatus *string `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// 备份结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 备份文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 备份集文件大小，单位Byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备份开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type AutoScalingConfig struct {
	// <p>ccu最小值</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RangeMin *float64 `json:"RangeMin,omitnil,omitempty" name:"RangeMin"`

	// <p>ccu最大值</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RangeMax *float64 `json:"RangeMax,omitnil,omitempty" name:"RangeMax"`
}

type BackupMethodStatisticsModel struct {
	// <p>自动备份大小，单位Byte</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoBackupSize *int64 `json:"AutoBackupSize,omitnil,omitempty" name:"AutoBackupSize"`

	// <p>总备份每天平均大小，单位Byte</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AverageSizePerDay *int64 `json:"AverageSizePerDay,omitnil,omitempty" name:"AverageSizePerDay"`

	// <p>手动备份大小，单位Byte</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManualBackupSize *int64 `json:"ManualBackupSize,omitnil,omitempty" name:"ManualBackupSize"`

	// <p>总备份大小，单位Byte</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSize *int64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`
}

type BackupMethodStatisticsOutPut struct {
	// <p>自动备份大小，单位Byte</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoBackupSize []*int64 `json:"AutoBackupSize,omitnil,omitempty" name:"AutoBackupSize"`

	// <p>手动备份大小，单位Byte</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManualBackupSize []*int64 `json:"ManualBackupSize,omitnil,omitempty" name:"ManualBackupSize"`
}

type BackupPolicyModelInput struct {
	// <p>备份结束时间</p>
	BackupEndTime *string `json:"BackupEndTime,omitnil,omitempty" name:"BackupEndTime"`

	// <p>备份方式  physical 物理备份 snapshot 快照备份</p>
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// <p>备份开始时间</p>
	BackupStartTime *string `json:"BackupStartTime,omitnil,omitempty" name:"BackupStartTime"`

	// <p>是否开启全量备份</p>
	EnableFull *int64 `json:"EnableFull,omitnil,omitempty" name:"EnableFull"`

	// <p>是否开启日志备份</p>
	EnableLog *int64 `json:"EnableLog,omitnil,omitempty" name:"EnableLog"`

	// <p>全备保留时间,目前只能设置7天</p>
	FullRetentionPeriod *int64 `json:"FullRetentionPeriod,omitnil,omitempty" name:"FullRetentionPeriod"`

	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>日志保留天数，目前只能设置保留7天</p>
	LogRetentionPeriod *int64 `json:"LogRetentionPeriod,omitnil,omitempty" name:"LogRetentionPeriod"`

	// <p>一周的哪几天进行备份</p>
	PeriodTime *string `json:"PeriodTime,omitnil,omitempty" name:"PeriodTime"`

	// <p>存储类型:COS,SNAPSHOT</p>枚举值：<ul><li> COS： COS存储</li><li> SNAPSHOT： 云盘快照</li></ul>
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`
}

type BackupPolicyModelOutPut struct {
	// <p>备份结束时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupEndTime *string `json:"BackupEndTime,omitnil,omitempty" name:"BackupEndTime"`

	// <p>备份方式</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// <p>备份开始时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupStartTime *string `json:"BackupStartTime,omitnil,omitempty" name:"BackupStartTime"`

	// <p>引擎类型</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBType *string `json:"DBType,omitnil,omitempty" name:"DBType"`

	// <p>引擎版本</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// <p>是否开启全量备份</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableFull *int64 `json:"EnableFull,omitnil,omitempty" name:"EnableFull"`

	// <p>是否开启日志备份</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableLog *int64 `json:"EnableLog,omitnil,omitempty" name:"EnableLog"`

	// <p>预计下次备份时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpectedNextBackupPeriod *string `json:"ExpectedNextBackupPeriod,omitnil,omitempty" name:"ExpectedNextBackupPeriod"`

	// <p>全备保留时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullRetentionPeriod *int64 `json:"FullRetentionPeriod,omitnil,omitempty" name:"FullRetentionPeriod"`

	// <p>策略ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>实例ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>日志保留天数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogRetentionPeriod *int64 `json:"LogRetentionPeriod,omitnil,omitempty" name:"LogRetentionPeriod"`

	// <p>一周的哪几天进行备份</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeriodTime *string `json:"PeriodTime,omitnil,omitempty" name:"PeriodTime"`

	// <p>地域</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>备份周期类型  0:Weekly</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeriodType *int64 `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`
}

type BackupSetModel struct {
	// 备份集耗时，单位 sec
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupDuration *int64 `json:"BackupDuration,omitnil,omitempty" name:"BackupDuration"`

	// 备份方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 备份备注名
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// 备份集进度百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupProgress *float64 `json:"BackupProgress,omitnil,omitempty" name:"BackupProgress"`

	// 备份集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupSetId *int64 `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// 备份集状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupStatus *string `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// 备份集类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 实例版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// 备份结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 事务commit结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTrxTime *string `json:"EndTrxTime,omitnil,omitempty" name:"EndTrxTime"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 备份过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 备份文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 备份集文件大小，单位Byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备份触发，0 - autobackup, 1 - manual
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManualBackup *int64 `json:"ManualBackup,omitnil,omitempty" name:"ManualBackup"`

	// 备份开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type BackupSetsReqFilter struct {
	// <p>备份方式,多个值用逗号隔开</p><p>枚举值：</p><ul><li>physical： 物理备份</li><li>snapshot： 快照备份</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// <p>备份状态，多个值用逗号隔开，含义说明：等待调度pending， 运行中 running, 成功success,失败 failed</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupStatus *string `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// <p>备份类型，多个值用逗号隔开，含义说明：全量备份 full</p><p>枚举值：</p><ul><li>full： 全量备份</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// <p>备份触发方式</p><p>枚举值：</p><ul><li>0： 自动备份</li><li>1： 手工备份</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManualBackup *string `json:"ManualBackup,omitnil,omitempty" name:"ManualBackup"`
}

type BackupStatisticsModel struct {
	// <p>总备份每天平均大小,单位Byte</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AverageSizePerDay *int64 `json:"AverageSizePerDay,omitnil,omitempty" name:"AverageSizePerDay"`

	// <p>数据备份大小，单位Byte</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataBackupSize *int64 `json:"DataBackupSize,omitnil,omitempty" name:"DataBackupSize"`

	// <p>日志备份大小，单位Byte</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogBackupSize *int64 `json:"LogBackupSize,omitnil,omitempty" name:"LogBackupSize"`

	// <p>总备份集个数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>总备份大小，单位Byte</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSize *int64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`
}

type BackupTypeStatisticsModel struct {
	// <p>数据备份大小，单位Byte</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataBackupSize []*int64 `json:"DataBackupSize,omitnil,omitempty" name:"DataBackupSize"`

	// <p>日志备份大小，单位Byte</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogBackupSize []*int64 `json:"LogBackupSize,omitnil,omitempty" name:"LogBackupSize"`
}

type BinlogInfo struct {
	// 日志服务的唯一 ID
	Sid *string `json:"Sid,omitnil,omitempty" name:"Sid"`

	// 日志服务的类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 实例的唯一 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

// Predefined struct for user
type CancelIsolateDBInstancesRequestParams struct {
	// 需要隔离的实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type CancelIsolateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 需要隔离的实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *CancelIsolateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelIsolateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelIsolateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelIsolateDBInstancesResponseParams struct {
	// 解除隔离成功实例Id列表
	SuccessInstanceIds []*string `json:"SuccessInstanceIds,omitnil,omitempty" name:"SuccessInstanceIds"`

	// 解除隔离失败实例Id列表
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitnil,omitempty" name:"FailedInstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelIsolateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CancelIsolateDBInstancesResponseParams `json:"Response"`
}

func (r *CancelIsolateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelIsolateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloneInstanceModel struct {
	// 克隆任务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloneEndTime *string `json:"CloneEndTime,omitnil,omitempty" name:"CloneEndTime"`

	// 克隆记录ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloneId *int64 `json:"CloneId,omitnil,omitempty" name:"CloneId"`

	// 克隆实例类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloneInsType *string `json:"CloneInsType,omitnil,omitempty" name:"CloneInsType"`

	// 克隆实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloneInstanceId *string `json:"CloneInstanceId,omitnil,omitempty" name:"CloneInstanceId"`

	// 克隆实例是否已经删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloneInstanceIsDeleted *bool `json:"CloneInstanceIsDeleted,omitnil,omitempty" name:"CloneInstanceIsDeleted"`

	// 克隆任务进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloneProgress *float64 `json:"CloneProgress,omitnil,omitempty" name:"CloneProgress"`

	// 克隆任务开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloneStartTime *string `json:"CloneStartTime,omitnil,omitempty" name:"CloneStartTime"`

	// 克隆任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloneStatus *string `json:"CloneStatus,omitnil,omitempty" name:"CloneStatus"`

	// 克隆时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloneTime *string `json:"CloneTime,omitnil,omitempty" name:"CloneTime"`

	// 克隆类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloneType *string `json:"CloneType,omitnil,omitempty" name:"CloneType"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 引擎类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBType *string `json:"DBType,omitnil,omitempty" name:"DBType"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 源实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceInstanceId *string `json:"SourceInstanceId,omitnil,omitempty" name:"SourceInstanceId"`
}

type ConstraintRange struct {
	// 约束类型为section时的最小值
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// 约束类型为section时的最大值
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`
}

// Predefined struct for user
type CreateCloneInstanceRequestParams struct {
	// <p>创建实例区域</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>字符型vpcid</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>字符型subnetid</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>购买规格</p>
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// <p>存储节点磁盘容量，单位GB</p>
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// <p>存储节点数量</p>
	StorageNodeNum *int64 `json:"StorageNodeNum,omitnil,omitempty" name:"StorageNodeNum"`

	// <p>源实例id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>实例名称，要求长度1-60，允许包含中文、英文大小写、数字、-、_</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>标签键值对数组</p>
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>备份回档名称</p>
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// <p>存储节点CPU核数</p>
	StorageNodeCpu *int64 `json:"StorageNodeCpu,omitnil,omitempty" name:"StorageNodeCpu"`

	// <p>存储节点内存大小</p>
	StorageNodeMem *int64 `json:"StorageNodeMem,omitnil,omitempty" name:"StorageNodeMem"`

	// <p>创建版本</p>
	CreateVersion *string `json:"CreateVersion,omitnil,omitempty" name:"CreateVersion"`

	// <p>创建端口号</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>恢复时间点</p>
	RecoverTime *string `json:"RecoverTime,omitnil,omitempty" name:"RecoverTime"`

	// <p>实例架构类型，separate:分离架构；hybrid:对等架构</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>磁盘类型，CLOUD_HSSD增强型SSD,CLOUD_TCS本地SSD盘</p>
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>多可用区列表，Zones[0]表示主可用区</p>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>全能型副本数</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`

	// <p>实例模式，normal：标准型；enhanced:加强型</p>
	InstanceMode *string `json:"InstanceMode,omitnil,omitempty" name:"InstanceMode"`

	// <p>安全组id列表</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type CreateCloneInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>创建实例区域</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>字符型vpcid</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>字符型subnetid</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>购买规格</p>
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// <p>存储节点磁盘容量，单位GB</p>
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// <p>存储节点数量</p>
	StorageNodeNum *int64 `json:"StorageNodeNum,omitnil,omitempty" name:"StorageNodeNum"`

	// <p>源实例id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>实例名称，要求长度1-60，允许包含中文、英文大小写、数字、-、_</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>标签键值对数组</p>
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>备份回档名称</p>
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// <p>存储节点CPU核数</p>
	StorageNodeCpu *int64 `json:"StorageNodeCpu,omitnil,omitempty" name:"StorageNodeCpu"`

	// <p>存储节点内存大小</p>
	StorageNodeMem *int64 `json:"StorageNodeMem,omitnil,omitempty" name:"StorageNodeMem"`

	// <p>创建版本</p>
	CreateVersion *string `json:"CreateVersion,omitnil,omitempty" name:"CreateVersion"`

	// <p>创建端口号</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>恢复时间点</p>
	RecoverTime *string `json:"RecoverTime,omitnil,omitempty" name:"RecoverTime"`

	// <p>实例架构类型，separate:分离架构；hybrid:对等架构</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>磁盘类型，CLOUD_HSSD增强型SSD,CLOUD_TCS本地SSD盘</p>
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>多可用区列表，Zones[0]表示主可用区</p>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>全能型副本数</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`

	// <p>实例模式，normal：标准型；enhanced:加强型</p>
	InstanceMode *string `json:"InstanceMode,omitnil,omitempty" name:"InstanceMode"`

	// <p>安全组id列表</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *CreateCloneInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloneInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "SpecCode")
	delete(f, "Disk")
	delete(f, "StorageNodeNum")
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "ResourceTags")
	delete(f, "BackupName")
	delete(f, "StorageNodeCpu")
	delete(f, "StorageNodeMem")
	delete(f, "CreateVersion")
	delete(f, "Vport")
	delete(f, "RecoverTime")
	delete(f, "InstanceType")
	delete(f, "StorageType")
	delete(f, "Zones")
	delete(f, "FullReplications")
	delete(f, "InstanceMode")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloneInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloneInstanceResponseParams struct {
	// <p>克隆实例ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>任务ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloneInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloneInstanceResponseParams `json:"Response"`
}

func (r *CreateCloneInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloneInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstancesRequestParams struct {
	// <p>创建实例区域</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>字符型vpcid</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>字符型subnetid</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>购买规格</p>
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// <p>存储节点磁盘容量，单位GB</p>
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// <p>存储节点数量</p>
	StorageNodeNum *int64 `json:"StorageNodeNum,omitnil,omitempty" name:"StorageNodeNum"`

	// <p>存储节点副本数量，最大为5，要求为奇数</p>
	Replications *int64 `json:"Replications,omitnil,omitempty" name:"Replications"`

	// <p>创建实例个数，上限为10</p>
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// <p>全能型副本数量</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`

	// <p>创建实例版本，默认使用当前最新版本</p>
	CreateVersion *string `json:"CreateVersion,omitnil,omitempty" name:"CreateVersion"`

	// <p>实例名称，要求长度1-60，允许包含中文、英文大小写、数字、-、_</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>标签键值对数组</p>
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>初始化实例参数。比如：<br>character_set_server（字符集，默认为utf8），<br>lower_case_table_names（表名大小写敏感，0 - 敏感；1-不敏感，默认为0）</p>
	InitParams []*InstanceParam `json:"InitParams,omitnil,omitempty" name:"InitParams"`

	// <p>时间单位，m：月</p>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// <p>商品的时间大小</p>
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// <p>存储节点CPU核数</p>
	StorageNodeCpu *int64 `json:"StorageNodeCpu,omitnil,omitempty" name:"StorageNodeCpu"`

	// <p>存储节点内存大小</p>
	StorageNodeMem *int64 `json:"StorageNodeMem,omitnil,omitempty" name:"StorageNodeMem"`

	// <p>付费模式，0表示按需计费/后付费，1表示预付费</p>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>管控节点数量</p>
	MCNum *int64 `json:"MCNum,omitnil,omitempty" name:"MCNum"`

	// <p>自定义端口</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>多AZ可用区列表</p>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>是否使用优惠卷</p>
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>优惠卷列表</p>
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// <p>实例架构类型，separate:分离架构；hybrid:对等架构</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>磁盘类型,CLOUD_HSSD增强型SSD,CLOUD_TCS本地SSD盘</p>
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>AZ模式，1:单AZ，2:多AZ非主AZ，3:多AZ主AZ</p>
	AZMode *int64 `json:"AZMode,omitnil,omitempty" name:"AZMode"`

	// <p>实例模式</p>
	InstanceMode *string `json:"InstanceMode,omitnil,omitempty" name:"InstanceMode"`

	// <p>参数模板id</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>兼容模式，enum:MySQL,HBase</p>
	SQLMode *string `json:"SQLMode,omitnil,omitempty" name:"SQLMode"`

	// <p>svls实例的ccu变配配置</p>
	AutoScaleConfig *AutoScalingConfig `json:"AutoScaleConfig,omitnil,omitempty" name:"AutoScaleConfig"`

	// <p>绑定安全组列表</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// <p>root用户名,当前版本默认为dbaadmin，传值也会重置为dbaadmin</p>
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// <p>dbaadmin密码</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>是否开启透明加密，0：不开启，1：开启</p>
	EncryptionEnable *int64 `json:"EncryptionEnable,omitnil,omitempty" name:"EncryptionEnable"`
}

type CreateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>创建实例区域</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>字符型vpcid</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>字符型subnetid</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>购买规格</p>
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// <p>存储节点磁盘容量，单位GB</p>
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// <p>存储节点数量</p>
	StorageNodeNum *int64 `json:"StorageNodeNum,omitnil,omitempty" name:"StorageNodeNum"`

	// <p>存储节点副本数量，最大为5，要求为奇数</p>
	Replications *int64 `json:"Replications,omitnil,omitempty" name:"Replications"`

	// <p>创建实例个数，上限为10</p>
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// <p>全能型副本数量</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`

	// <p>创建实例版本，默认使用当前最新版本</p>
	CreateVersion *string `json:"CreateVersion,omitnil,omitempty" name:"CreateVersion"`

	// <p>实例名称，要求长度1-60，允许包含中文、英文大小写、数字、-、_</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>标签键值对数组</p>
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>初始化实例参数。比如：<br>character_set_server（字符集，默认为utf8），<br>lower_case_table_names（表名大小写敏感，0 - 敏感；1-不敏感，默认为0）</p>
	InitParams []*InstanceParam `json:"InitParams,omitnil,omitempty" name:"InitParams"`

	// <p>时间单位，m：月</p>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// <p>商品的时间大小</p>
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// <p>存储节点CPU核数</p>
	StorageNodeCpu *int64 `json:"StorageNodeCpu,omitnil,omitempty" name:"StorageNodeCpu"`

	// <p>存储节点内存大小</p>
	StorageNodeMem *int64 `json:"StorageNodeMem,omitnil,omitempty" name:"StorageNodeMem"`

	// <p>付费模式，0表示按需计费/后付费，1表示预付费</p>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>管控节点数量</p>
	MCNum *int64 `json:"MCNum,omitnil,omitempty" name:"MCNum"`

	// <p>自定义端口</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>多AZ可用区列表</p>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>是否使用优惠卷</p>
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>优惠卷列表</p>
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// <p>实例架构类型，separate:分离架构；hybrid:对等架构</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>磁盘类型,CLOUD_HSSD增强型SSD,CLOUD_TCS本地SSD盘</p>
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>AZ模式，1:单AZ，2:多AZ非主AZ，3:多AZ主AZ</p>
	AZMode *int64 `json:"AZMode,omitnil,omitempty" name:"AZMode"`

	// <p>实例模式</p>
	InstanceMode *string `json:"InstanceMode,omitnil,omitempty" name:"InstanceMode"`

	// <p>参数模板id</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>兼容模式，enum:MySQL,HBase</p>
	SQLMode *string `json:"SQLMode,omitnil,omitempty" name:"SQLMode"`

	// <p>svls实例的ccu变配配置</p>
	AutoScaleConfig *AutoScalingConfig `json:"AutoScaleConfig,omitnil,omitempty" name:"AutoScaleConfig"`

	// <p>绑定安全组列表</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// <p>root用户名,当前版本默认为dbaadmin，传值也会重置为dbaadmin</p>
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// <p>dbaadmin密码</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>是否开启透明加密，0：不开启，1：开启</p>
	EncryptionEnable *int64 `json:"EncryptionEnable,omitnil,omitempty" name:"EncryptionEnable"`
}

func (r *CreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "SpecCode")
	delete(f, "Disk")
	delete(f, "StorageNodeNum")
	delete(f, "Replications")
	delete(f, "InstanceCount")
	delete(f, "FullReplications")
	delete(f, "CreateVersion")
	delete(f, "InstanceName")
	delete(f, "ResourceTags")
	delete(f, "InitParams")
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	delete(f, "StorageNodeCpu")
	delete(f, "StorageNodeMem")
	delete(f, "PayMode")
	delete(f, "MCNum")
	delete(f, "Vport")
	delete(f, "Zones")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "InstanceType")
	delete(f, "StorageType")
	delete(f, "AZMode")
	delete(f, "InstanceMode")
	delete(f, "TemplateId")
	delete(f, "SQLMode")
	delete(f, "AutoScaleConfig")
	delete(f, "SecurityGroupIds")
	delete(f, "UserName")
	delete(f, "Password")
	delete(f, "EncryptionEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstancesResponseParams struct {
	// <p>实例ID</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>任务ID</p>
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstancesResponseParams `json:"Response"`
}

func (r *CreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBSBackupRequestParams struct {
	// <p>备份方式：physical、snapshot 这个值和DescribeDBSBackupPolicy接口返回的backupMethod保持一致</p><p>枚举值：</p><ul><li>physical： 物理备份</li><li>snapshot： 快照备份</li></ul>
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// <p>备份类型：暂时只支持full</p>
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>备份备注</p>
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
}

type CreateDBSBackupRequest struct {
	*tchttp.BaseRequest
	
	// <p>备份方式：physical、snapshot 这个值和DescribeDBSBackupPolicy接口返回的backupMethod保持一致</p><p>枚举值：</p><ul><li>physical： 物理备份</li><li>snapshot： 快照备份</li></ul>
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// <p>备份类型：暂时只支持full</p>
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>备份备注</p>
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
}

func (r *CreateDBSBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBSBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupMethod")
	delete(f, "BackupType")
	delete(f, "InstanceId")
	delete(f, "BackupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBSBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBSBackupResponseParams struct {
	// <p>备份集ID</p>
	BackupSetId *int64 `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// <p>是否成功</p>
	IsSuccess *bool `json:"IsSuccess,omitnil,omitempty" name:"IsSuccess"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBSBackupResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBSBackupResponseParams `json:"Response"`
}

func (r *CreateDBSBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBSBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBParamValue struct {
	// 参数名称
	Param *string `json:"Param,omitnil,omitempty" name:"Param"`

	// 参数值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type DataBackupStatisticsModel struct {
	// 自动备份个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoBackupCount *int64 `json:"AutoBackupCount,omitnil,omitempty" name:"AutoBackupCount"`

	// 自动备份大小，单位Byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoBackupSize *int64 `json:"AutoBackupSize,omitnil,omitempty" name:"AutoBackupSize"`

	// 平均每个备份大小,单位Byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	AverageSizePerBackup *int64 `json:"AverageSizePerBackup,omitnil,omitempty" name:"AverageSizePerBackup"`

	// 平均每天备份大小,单位Byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	AverageSizePerDay *int64 `json:"AverageSizePerDay,omitnil,omitempty" name:"AverageSizePerDay"`

	// 手工备份个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManualBackupCount *int64 `json:"ManualBackupCount,omitnil,omitempty" name:"ManualBackupCount"`

	// 手工备份大小，单位Byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManualBackupSize *int64 `json:"ManualBackupSize,omitnil,omitempty" name:"ManualBackupSize"`

	// 数据备份个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 数据备份大小，单位Byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSize *int64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`
}

type Database struct {
	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`
}

type DatabaseFunction struct {
	// 函数名称
	Func *string `json:"Func,omitnil,omitempty" name:"Func"`
}

type DatabasePrivileges struct {
	// database名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 权限列表
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

type DatabaseProcedure struct {
	// 存储过程名称
	Proc *string `json:"Proc,omitnil,omitempty" name:"Proc"`
}

type DatabaseTable struct {
	// 表名
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`
}

type DatabaseView struct {
	// 视图名称
	View *string `json:"View,omitnil,omitempty" name:"View"`
}

// Predefined struct for user
type DeleteDBSBackupSetsRequestParams struct {
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>备份集列表 ,值来自 DescribeDBSBackupSets 接口返回</p>
	BackupSetIdList []*int64 `json:"BackupSetIdList,omitnil,omitempty" name:"BackupSetIdList"`
}

type DeleteDBSBackupSetsRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>备份集列表 ,值来自 DescribeDBSBackupSets 接口返回</p>
	BackupSetIdList []*int64 `json:"BackupSetIdList,omitnil,omitempty" name:"BackupSetIdList"`
}

func (r *DeleteDBSBackupSetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDBSBackupSetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupSetIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDBSBackupSetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDBSBackupSetsResponseParams struct {
	// <p>本次实际删除的备份数量</p>
	Deleted *int64 `json:"Deleted,omitnil,omitempty" name:"Deleted"`

	// <p>是否全部删除成功</p>
	IsSuccess *bool `json:"IsSuccess,omitnil,omitempty" name:"IsSuccess"`

	// <p>需要删除的备份总数</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDBSBackupSetsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDBSBackupSetsResponseParams `json:"Response"`
}

func (r *DeleteDBSBackupSetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDBSBackupSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceDetailRequestParams struct {
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceDetailResponseParams struct {
	// <p>实例名称</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>区域</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>字符型vpcid</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>字符型subnetid</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>创建实例版本</p>
	CreateVersion *string `json:"CreateVersion,omitnil,omitempty" name:"CreateVersion"`

	// <p>子网IP</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>子网端口</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>实例状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>存储节点磁盘容量，单位GB</p>
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// <p>存储节点数量</p>
	StorageNodeNum *int64 `json:"StorageNodeNum,omitnil,omitempty" name:"StorageNodeNum"`

	// <p>初始化实例参数</p>
	InitParams []*InstanceParam `json:"InitParams,omitnil,omitempty" name:"InitParams"`

	// <p>实例标签信息</p>
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>存储节点副本数量</p>
	Replications *int64 `json:"Replications,omitnil,omitempty" name:"Replications"`

	// <p>全能型副本数</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`

	// <p>字符集</p>
	CharSet *string `json:"CharSet,omitnil,omitempty" name:"CharSet"`

	// <p>节点信息</p>
	Node []*NodeInfo `json:"Node,omitnil,omitempty" name:"Node"`

	// <p>实例所属地域</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>实例规格</p>
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>实例状态中文描述</p>
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// <p>存储节点CPU核数</p>
	StorageNodeCpu *int64 `json:"StorageNodeCpu,omitnil,omitempty" name:"StorageNodeCpu"`

	// <p>存储节点内存大小</p>
	StorageNodeMem *int64 `json:"StorageNodeMem,omitnil,omitempty" name:"StorageNodeMem"`

	// <p>续费标志</p>
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// <p>付费模式，0后付费，1预付费</p>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>过期时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireAt *string `json:"ExpireAt,omitnil,omitempty" name:"ExpireAt"`

	// <p>隔离时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolatedAt *string `json:"IsolatedAt,omitnil,omitempty" name:"IsolatedAt"`

	// <p>实例架构类型，separate:分离架构；hybrid:对等架构</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>磁盘类型，CLOUD_HSSD增强型SSD,CLOUD_TCS本地SSD盘</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>实例节点可用区列表。Zones[0]表示主可用区</p>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>最大节点磁盘使用量</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskUsage *int64 `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// <p>binlog是否开启</p>
	BinlogStatus *int64 `json:"BinlogStatus,omitnil,omitempty" name:"BinlogStatus"`

	// <p>az模式，1: 单AZ 2: 多AZ非主AZ模式 3: 多AZ主AZ模式</p>
	AZMode *int64 `json:"AZMode,omitnil,omitempty" name:"AZMode"`

	// <p>灾备标志位，1: 无灾备关系；2: 灾备主实例；3: 灾备备实例</p>
	StandbyFlag *int64 `json:"StandbyFlag,omitnil,omitempty" name:"StandbyFlag"`

	// <p>cdc节点类型</p>
	BinlogType *string `json:"BinlogType,omitnil,omitempty" name:"BinlogType"`

	// <p>1表示支持，0表示不支持</p>
	TimingModifyInstanceFlag *int64 `json:"TimingModifyInstanceFlag,omitnil,omitempty" name:"TimingModifyInstanceFlag"`

	// <p>列存节点cpu核数</p>
	ColumnarNodeCpu *int64 `json:"ColumnarNodeCpu,omitnil,omitempty" name:"ColumnarNodeCpu"`

	// <p>列存节点mem内存大小</p>
	ColumnarNodeMem *int64 `json:"ColumnarNodeMem,omitnil,omitempty" name:"ColumnarNodeMem"`

	// <p>列存节点个数</p>
	ColumnarNodeNum *int64 `json:"ColumnarNodeNum,omitnil,omitempty" name:"ColumnarNodeNum"`

	// <p>列存节点磁盘大小</p>
	ColumnarNodeDisk *int64 `json:"ColumnarNodeDisk,omitnil,omitempty" name:"ColumnarNodeDisk"`

	// <p>列存节点磁盘类型</p>
	ColumnarNodeStorageType *string `json:"ColumnarNodeStorageType,omitnil,omitempty" name:"ColumnarNodeStorageType"`

	// <p>列存节点规格</p>
	ColumnarNodeSpecCode *string `json:"ColumnarNodeSpecCode,omitnil,omitempty" name:"ColumnarNodeSpecCode"`

	// <p>列存 vip</p>
	ColumnarVip *string `json:"ColumnarVip,omitnil,omitempty" name:"ColumnarVip"`

	// <p>列存 vport</p>
	ColumnarVport *int64 `json:"ColumnarVport,omitnil,omitempty" name:"ColumnarVport"`

	// <p>实例是否支持列存</p>
	IsSupportColumnar *bool `json:"IsSupportColumnar,omitnil,omitempty" name:"IsSupportColumnar"`

	// <p>实例类型</p>
	InstanceCategory *int64 `json:"InstanceCategory,omitnil,omitempty" name:"InstanceCategory"`

	// <p>兼容模式</p>
	SQLMode *string `json:"SQLMode,omitnil,omitempty" name:"SQLMode"`

	// <p>是否支持修改全能型副本数量</p>
	IsSwitchFullReplicationsEnable *bool `json:"IsSwitchFullReplicationsEnable,omitnil,omitempty" name:"IsSwitchFullReplicationsEnable"`

	// <p>实例类型</p>
	InstanceMode *string `json:"InstanceMode,omitnil,omitempty" name:"InstanceMode"`

	// <p>dumper vip</p>
	DumperVip *string `json:"DumperVip,omitnil,omitempty" name:"DumperVip"`

	// <p>dumper vport</p>
	DumperVport *int64 `json:"DumperVport,omitnil,omitempty" name:"DumperVport"`

	// <p>svls实例的ccu变配范围</p>
	AutoScaleConfig *AutoScalingConfig `json:"AutoScaleConfig,omitnil,omitempty" name:"AutoScaleConfig"`

	// <p>参数模板id</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>参数模板名称</p>
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// <p>实例分析引擎模式</p><p>枚举值：</p><ul><li>libra： LibraDB分析引擎实例</li></ul>
	AnalysisMode *string `json:"AnalysisMode,omitnil,omitempty" name:"AnalysisMode"`

	// <p>分析引擎实例关系</p>
	AnalysisRelationInfos []*AnalysisRelationInfo `json:"AnalysisRelationInfos,omitnil,omitempty" name:"AnalysisRelationInfos"`

	// <p>分析引擎实例信息</p><p>Cpu、Memory、Disk复用顶层字段</p>
	AnalysisInstanceInfo *AnalysisInstanceInfo `json:"AnalysisInstanceInfo,omitnil,omitempty" name:"AnalysisInstanceInfo"`

	// <p>维护窗口配置</p>
	MaintenanceWindow *MaintenanceWindowInfo `json:"MaintenanceWindow,omitnil,omitempty" name:"MaintenanceWindow"`

	// <p>是否开启透明加密，0：未开启；1：已开启</p>
	EncryptionEnable *int64 `json:"EncryptionEnable,omitnil,omitempty" name:"EncryptionEnable"`

	// <p>真实使用的kms地域，用于后续调用kms服务</p>
	EncryptionKmsRegion *string `json:"EncryptionKmsRegion,omitnil,omitempty" name:"EncryptionKmsRegion"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceDetailResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesRequestParams struct {
	// <p>过滤参数</p>
	Filters []*InstanceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>最大返回个数，默认为20，上限为100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量，取Limit整数倍</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>指定查询引擎类型</p><p>枚举值：</p><ul><li>libra： 列存引擎</li></ul>
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>过滤参数</p>
	Filters []*InstanceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>最大返回个数，默认为20，上限为100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量，取Limit整数倍</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>指定查询引擎类型</p><p>枚举值：</p><ul><li>libra： 列存引擎</li></ul>
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EngineType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesResponseParams struct {
	// <p>返回实例列表信息</p>
	Instances []*InstanceInfo `json:"Instances,omitnil,omitempty" name:"Instances"`

	// <p>满足条件总数量</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBParametersRequestParams struct {
	// 实例 ID，形如：tdsql3-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBParametersRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如：tdsql3-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBParametersResponseParams struct {
	// 实例 ID，形如：tdsql3-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 请求实例的当前参数值
	Params []*ParamDesc `json:"Params,omitnil,omitempty" name:"Params"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBParametersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBParametersResponseParams `json:"Response"`
}

func (r *DescribeDBParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSArchiveLogsRequestParams struct {
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>日志记录ID</p>
	ArchiveLogId *int64 `json:"ArchiveLogId,omitnil,omitempty" name:"ArchiveLogId"`

	// <p>结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>备份状态：pending,running,success,failed</p>
	FilterStatus *string `json:"FilterStatus,omitnil,omitempty" name:"FilterStatus"`

	// <p>条数限制</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>排序字段，枚举：StartTime,EndTime,ExpiredTime,FileSize,BackupDuration</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>排序方式：ASC：顺序, DESC：倒序</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>开始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type DescribeDBSArchiveLogsRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>日志记录ID</p>
	ArchiveLogId *int64 `json:"ArchiveLogId,omitnil,omitempty" name:"ArchiveLogId"`

	// <p>结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>备份状态：pending,running,success,failed</p>
	FilterStatus *string `json:"FilterStatus,omitnil,omitempty" name:"FilterStatus"`

	// <p>条数限制</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>排序字段，枚举：StartTime,EndTime,ExpiredTime,FileSize,BackupDuration</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>排序方式：ASC：顺序, DESC：倒序</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>开始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

func (r *DescribeDBSArchiveLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSArchiveLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ArchiveLogId")
	delete(f, "EndTime")
	delete(f, "FilterStatus")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "StartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSArchiveLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSArchiveLogsResponseParams struct {
	// <p>归档日志列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*ArchiveLogModel `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSArchiveLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSArchiveLogsResponseParams `json:"Response"`
}

func (r *DescribeDBSArchiveLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSArchiveLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSAvailableRecoveryTimeRequestParams struct {
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBSAvailableRecoveryTimeRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSAvailableRecoveryTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSAvailableRecoveryTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSAvailableRecoveryTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSAvailableRecoveryTimeResponseParams struct {
	// <p>结束时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>可恢复时间区间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Intervals []*ArchiveLogInterval `json:"Intervals,omitnil,omitempty" name:"Intervals"`

	// <p>开始时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSAvailableRecoveryTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSAvailableRecoveryTimeResponseParams `json:"Response"`
}

func (r *DescribeDBSAvailableRecoveryTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSAvailableRecoveryTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSBackupPolicyRequestParams struct {
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBSBackupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSBackupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSBackupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSBackupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSBackupPolicyResponseParams struct {
	// <p>备份策略列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*BackupPolicyModelOutPut `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSBackupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSBackupPolicyResponseParams `json:"Response"`
}

func (r *DescribeDBSBackupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSBackupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSBackupSetsRequestParams struct {
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>实例备份集ID</p>
	BackupSetId *int64 `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// <p>结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>过滤条件</p>
	FilterBy *BackupSetsReqFilter `json:"FilterBy,omitnil,omitempty" name:"FilterBy"`

	// <p>单次查询数量[0,200]</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>本次查询偏移[0,INF]</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>StartTime,EndTime,ExpiredTime,BackupSetId,BackupDuration</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>ASC,DESC</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>开始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type DescribeDBSBackupSetsRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>实例备份集ID</p>
	BackupSetId *int64 `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// <p>结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>过滤条件</p>
	FilterBy *BackupSetsReqFilter `json:"FilterBy,omitnil,omitempty" name:"FilterBy"`

	// <p>单次查询数量[0,200]</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>本次查询偏移[0,INF]</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>StartTime,EndTime,ExpiredTime,BackupSetId,BackupDuration</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>ASC,DESC</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>开始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

func (r *DescribeDBSBackupSetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSBackupSetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupSetId")
	delete(f, "EndTime")
	delete(f, "FilterBy")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "StartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSBackupSetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSBackupSetsResponseParams struct {
	// <p>备份集列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*BackupSetModel `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSBackupSetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSBackupSetsResponseParams `json:"Response"`
}

func (r *DescribeDBSBackupSetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSBackupSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSBackupStatisticsDetailRequestParams struct {
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>开始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type DescribeDBSBackupStatisticsDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>开始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

func (r *DescribeDBSBackupStatisticsDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSBackupStatisticsDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EndTime")
	delete(f, "StartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSBackupStatisticsDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSBackupStatisticsDetailResponseParams struct {
	// <p>按备份方式统计</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupMethodStatistics *BackupMethodStatisticsOutPut `json:"BackupMethodStatistics,omitnil,omitempty" name:"BackupMethodStatistics"`

	// <p>备份时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupTime []*string `json:"BackupTime,omitnil,omitempty" name:"BackupTime"`

	// <p>按备份数据类型统计</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupTypeStatistics *BackupTypeStatisticsModel `json:"BackupTypeStatistics,omitnil,omitempty" name:"BackupTypeStatistics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSBackupStatisticsDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSBackupStatisticsDetailResponseParams `json:"Response"`
}

func (r *DescribeDBSBackupStatisticsDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSBackupStatisticsDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSBackupStatisticsRequestParams struct {
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>开始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type DescribeDBSBackupStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>开始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

func (r *DescribeDBSBackupStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSBackupStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EndTime")
	delete(f, "StartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSBackupStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSBackupStatisticsResponseParams struct {
	// <p>备份方式统计</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupMethodStatistics *BackupMethodStatisticsModel `json:"BackupMethodStatistics,omitnil,omitempty" name:"BackupMethodStatistics"`

	// <p>总备份集统计</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupStatistics *BackupStatisticsModel `json:"BackupStatistics,omitnil,omitempty" name:"BackupStatistics"`

	// <p>数据备份统计</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataBackupStatistics *DataBackupStatisticsModel `json:"DataBackupStatistics,omitnil,omitempty" name:"DataBackupStatistics"`

	// <p>日志备份统计</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogBackupStatistics *LogBackupStatisticsModel `json:"LogBackupStatistics,omitnil,omitempty" name:"LogBackupStatistics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSBackupStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSBackupStatisticsResponseParams `json:"Response"`
}

func (r *DescribeDBSBackupStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSBackupStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSCloneInstancesRequestParams struct {
	// <p>源实例ID</p>
	SourceInstanceId *string `json:"SourceInstanceId,omitnil,omitempty" name:"SourceInstanceId"`

	// <p>引擎类型</p>
	DBType *string `json:"DBType,omitnil,omitempty" name:"DBType"`

	// <p>结束创建时间</p>
	EndCreateTime *string `json:"EndCreateTime,omitnil,omitempty" name:"EndCreateTime"`

	// <p>克隆类型: PITR 时间点 BackupSet 备份集</p>
	FilterCloneType *string `json:"FilterCloneType,omitnil,omitempty" name:"FilterCloneType"`

	// <p>查询数量[0,200]</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>查询偏移量[0,INF]</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>排序字段: CloneTime,CreateTime</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>排序类型:ASC,DESC</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>开始创建时间</p>
	StartCreateTime *string `json:"StartCreateTime,omitnil,omitempty" name:"StartCreateTime"`
}

type DescribeDBSCloneInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>源实例ID</p>
	SourceInstanceId *string `json:"SourceInstanceId,omitnil,omitempty" name:"SourceInstanceId"`

	// <p>引擎类型</p>
	DBType *string `json:"DBType,omitnil,omitempty" name:"DBType"`

	// <p>结束创建时间</p>
	EndCreateTime *string `json:"EndCreateTime,omitnil,omitempty" name:"EndCreateTime"`

	// <p>克隆类型: PITR 时间点 BackupSet 备份集</p>
	FilterCloneType *string `json:"FilterCloneType,omitnil,omitempty" name:"FilterCloneType"`

	// <p>查询数量[0,200]</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>查询偏移量[0,INF]</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>排序字段: CloneTime,CreateTime</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>排序类型:ASC,DESC</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>开始创建时间</p>
	StartCreateTime *string `json:"StartCreateTime,omitnil,omitempty" name:"StartCreateTime"`
}

func (r *DescribeDBSCloneInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSCloneInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceInstanceId")
	delete(f, "DBType")
	delete(f, "EndCreateTime")
	delete(f, "FilterCloneType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "StartCreateTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSCloneInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSCloneInstancesResponseParams struct {
	// <p>克隆列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*CloneInstanceModel `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSCloneInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSCloneInstancesResponseParams `json:"Response"`
}

func (r *DescribeDBSCloneInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSCloneInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsResponseParams struct {
	// 安全组详情。
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 实例VIP
	// 注意：此字段可能返回 null，表示取不到有效值。
	VIP *string `json:"VIP,omitnil,omitempty" name:"VIP"`

	// 实例端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	VPort *string `json:"VPort,omitnil,omitempty" name:"VPort"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeDBSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseObjectsRequestParams struct {
	// <p>实例 ID，形如：tdsql3-42f40429.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>数据库名称，通过 DescribeDatabases 接口获取。</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>分页索引</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>数据表名称匹配表达式</p>
	TableRegexp *string `json:"TableRegexp,omitnil,omitempty" name:"TableRegexp"`
}

type DescribeDatabaseObjectsRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例 ID，形如：tdsql3-42f40429.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>数据库名称，通过 DescribeDatabases 接口获取。</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>分页索引</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>数据表名称匹配表达式</p>
	TableRegexp *string `json:"TableRegexp,omitnil,omitempty" name:"TableRegexp"`
}

func (r *DescribeDatabaseObjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseObjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DbName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TableRegexp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseObjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseObjectsResponseParams struct {
	// <p>透传入参。</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>数据库名称。</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>表列表。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tables []*DatabaseTable `json:"Tables,omitnil,omitempty" name:"Tables"`

	// <p>视图列表。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Views []*DatabaseView `json:"Views,omitnil,omitempty" name:"Views"`

	// <p>存储过程列表。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Procs []*DatabaseProcedure `json:"Procs,omitnil,omitempty" name:"Procs"`

	// <p>函数列表。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Funcs []*DatabaseFunction `json:"Funcs,omitnil,omitempty" name:"Funcs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabaseObjectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseObjectsResponseParams `json:"Response"`
}

func (r *DescribeDatabaseObjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesRequestParams struct {
	// <p>实例 ID，形如：tdsql3-ow7t8lmc。</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>分页索引</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>数据库名称匹配表达式</p>
	DatabaseRegexp *string `json:"DatabaseRegexp,omitnil,omitempty" name:"DatabaseRegexp"`
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例 ID，形如：tdsql3-ow7t8lmc。</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>分页索引</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>数据库名称匹配表达式</p>
	DatabaseRegexp *string `json:"DatabaseRegexp,omitnil,omitempty" name:"DatabaseRegexp"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DatabaseRegexp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesResponseParams struct {
	// <p>透传实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>该实例上的数据库列表。</p>
	Databases []*Database `json:"Databases,omitnil,omitempty" name:"Databases"`

	// <p>总数量</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabasesResponseParams `json:"Response"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowRequestParams struct {

}

type DescribeFlowRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFlowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowResponseParams `json:"Response"`
}

func (r *DescribeFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSaleInfoRequestParams struct {
	// <p>灾备主实例地域</p>
	SrcRegion *string `json:"SrcRegion,omitnil,omitempty" name:"SrcRegion"`

	// <p>实例id</p><p>传入此参数表示返回这个实例所在的地域可用区的售卖信息</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>指定支持某种类型实例的 sale 信息</p><p>枚举值：</p><ul><li>serverless： 返回支持 serverless 型实例的所有 region</li></ul><p>默认值：无</p><p>当前仅支持指定 serverless，传入其他信息或者不传则默认返回所有售卖地域信息</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type DescribeSaleInfoRequest struct {
	*tchttp.BaseRequest
	
	// <p>灾备主实例地域</p>
	SrcRegion *string `json:"SrcRegion,omitnil,omitempty" name:"SrcRegion"`

	// <p>实例id</p><p>传入此参数表示返回这个实例所在的地域可用区的售卖信息</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>指定支持某种类型实例的 sale 信息</p><p>枚举值：</p><ul><li>serverless： 返回支持 serverless 型实例的所有 region</li></ul><p>默认值：无</p><p>当前仅支持指定 serverless，传入其他信息或者不传则默认返回所有售卖地域信息</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

func (r *DescribeSaleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSaleInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SrcRegion")
	delete(f, "InstanceId")
	delete(f, "InstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSaleInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSaleInfoResponseParams struct {
	// <p>返回可售卖region信息</p>
	RegionList []*DescribeSaleRegionInfo `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSaleInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSaleInfoResponseParams `json:"Response"`
}

func (r *DescribeSaleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSaleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSaleRegionInfo struct {
	// <p>Region英文字符串</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>售卖Zone列表</p>
	ZoneList []*DescribeSaleZonesInfo `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// <p>Region中文字符串</p>
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// <p>是否售卖。1:售卖，0不售卖</p>
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// <p>多可用可选数量</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableZoneNum *int64 `json:"AvailableZoneNum,omitnil,omitempty" name:"AvailableZoneNum"`

	// <p>是否允许使用日志副本</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSupportedLogReplica *bool `json:"IsSupportedLogReplica,omitnil,omitempty" name:"IsSupportedLogReplica"`

	// <p>可用区组合</p>
	ZoneGroup []*DescribeSaleZonesGroup `json:"ZoneGroup,omitnil,omitempty" name:"ZoneGroup"`

	// <p>是否支持serverless</p>
	IsSupportServerless *bool `json:"IsSupportServerless,omitnil,omitempty" name:"IsSupportServerless"`
}

type DescribeSaleZonesGroup struct {
	// <p>可用区数</p>
	ZoneNum *int64 `json:"ZoneNum,omitnil,omitempty" name:"ZoneNum"`

	// <p>可用区组合</p>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>支持的类型</p>
	SupportTypes []*string `json:"SupportTypes,omitnil,omitempty" name:"SupportTypes"`

	// <p>支持的磁盘类型</p><p>枚举值：</p><ul><li>CLOUD_TCS： 本地盘</li><li>CLOUD_HSSD： 增强型云盘</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableDiskTypes []*string `json:"AvailableDiskTypes,omitnil,omitempty" name:"AvailableDiskTypes"`

	// <p>是否支持serverless</p>
	IsSupportServerless *bool `json:"IsSupportServerless,omitnil,omitempty" name:"IsSupportServerless"`
}

type DescribeSaleZonesInfo struct {
	// <p>Zone英文字符串</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Zone中文字符串</p>
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// <p>是否售卖，1：售卖；0：不售卖</p>
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// <p>是否是默认主可用区,0不是，1是</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefaultMaster *int64 `json:"IsDefaultMaster,omitnil,omitempty" name:"IsDefaultMaster"`

	// <p>可用区可选磁盘类型</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableDiskTypes []*string `json:"AvailableDiskTypes,omitnil,omitempty" name:"AvailableDiskTypes"`

	// <p>是否是独享可用区</p>
	SupportTypes []*string `json:"SupportTypes,omitnil,omitempty" name:"SupportTypes"`

	// <p>是否支持serverless</p>
	IsSupportServerless *bool `json:"IsSupportServerless,omitnil,omitempty" name:"IsSupportServerless"`
}

// Predefined struct for user
type DescribeSpecsRequestParams struct {
	// <p>全能型副本数，当前支持范围：1-3，默认为3</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`

	// <p>独享实例</p>
	IsExclusiveInstance *int64 `json:"IsExclusiveInstance,omitnil,omitempty" name:"IsExclusiveInstance"`
}

type DescribeSpecsRequest struct {
	*tchttp.BaseRequest
	
	// <p>全能型副本数，当前支持范围：1-3，默认为3</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`

	// <p>独享实例</p>
	IsExclusiveInstance *int64 `json:"IsExclusiveInstance,omitnil,omitempty" name:"IsExclusiveInstance"`
}

func (r *DescribeSpecsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FullReplications")
	delete(f, "IsExclusiveInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpecsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecsResponseParams struct {
	// <p>对等节点售卖规格列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	HybridNodeSpecs []*StorageNodeSpec `json:"HybridNodeSpecs,omitnil,omitempty" name:"HybridNodeSpecs"`

	// <p>svls节点售卖规格列表</p>
	ServerlessCcuSpec []*ServerlessCcu `json:"ServerlessCcuSpec,omitnil,omitempty" name:"ServerlessCcuSpec"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSpecsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpecsResponseParams `json:"Response"`
}

func (r *DescribeSpecsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserPrivilegesRequestParams struct {
	// 实例 ID，形如：tdsql3-5baee8df。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 登录用户名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 数据库名。如果为 \*，表示查询全局权限（即 \*.\*），此时忽略 Type 和 Object 参数
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 具体的 Type 的名称，例如 Type 为 table 时就是具体的表名。DbName 和 Type 都为具体名称，则 Object 表示具体对象名，不能为 \* 或者为空
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// 类型,可以填入 table 、 view 、 proc 、 func 和 \*。当 DbName 为具体数据库名，Type为 \* 时，表示查询该数据库权限（即db.\*），此时忽略 Object 参数
	ObjectType *string `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`

	// 当 Type=table 时，ColName 为 \* 表示查询表的权限，如果为具体字段名，表示查询对应字段的权限
	ColName *string `json:"ColName,omitnil,omitempty" name:"ColName"`
}

type DescribeUserPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如：tdsql3-5baee8df。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 登录用户名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 数据库名。如果为 \*，表示查询全局权限（即 \*.\*），此时忽略 Type 和 Object 参数
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 具体的 Type 的名称，例如 Type 为 table 时就是具体的表名。DbName 和 Type 都为具体名称，则 Object 表示具体对象名，不能为 \* 或者为空
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// 类型,可以填入 table 、 view 、 proc 、 func 和 \*。当 DbName 为具体数据库名，Type为 \* 时，表示查询该数据库权限（即db.\*），此时忽略 Object 参数
	ObjectType *string `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`

	// 当 Type=table 时，ColName 为 \* 表示查询表的权限，如果为具体字段名，表示查询对应字段的权限
	ColName *string `json:"ColName,omitnil,omitempty" name:"ColName"`
}

func (r *DescribeUserPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Host")
	delete(f, "UserName")
	delete(f, "DbName")
	delete(f, "Object")
	delete(f, "ObjectType")
	delete(f, "ColName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserPrivilegesResponseParams struct {
	// 权限列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserPrivilegesResponseParams `json:"Response"`
}

func (r *DescribeUserPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeUsersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersResponseParams struct {
	// 用户列表
	Users []*UserInfo `json:"Users,omitnil,omitempty" name:"Users"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUsersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsersResponseParams `json:"Response"`
}

func (r *DescribeUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstancesRequestParams struct {
	// 需要隔离的实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DestroyInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 需要隔离的实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DestroyInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstancesResponseParams struct {
	// 解除隔离成功实例Id列表
	SuccessInstanceIds []*string `json:"SuccessInstanceIds,omitnil,omitempty" name:"SuccessInstanceIds"`

	// 解除隔离失败实例Id列表
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitnil,omitempty" name:"FailedInstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DestroyInstancesResponseParams `json:"Response"`
}

func (r *DestroyInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExpandInstanceRequestParams struct {
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>存储节点扩容至多少个节点，如果没有变化，传当前节点数</p>
	StorageNodeNum *int64 `json:"StorageNodeNum,omitnil,omitempty" name:"StorageNodeNum"`

	// <p>可用区列表</p>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>az模式，1: 单AZ 2: 多AZ非主AZ模式 3: 多AZ主AZ模式</p>
	AZMode *int64 `json:"AZMode,omitnil,omitempty" name:"AZMode"`

	// <p>AZMode传3时，表示主AZ</p>
	PrimaryAZ *string `json:"PrimaryAZ,omitnil,omitempty" name:"PrimaryAZ"`

	// <p>全能型副本数，取值范围包括1-3</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`
}

type ExpandInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>存储节点扩容至多少个节点，如果没有变化，传当前节点数</p>
	StorageNodeNum *int64 `json:"StorageNodeNum,omitnil,omitempty" name:"StorageNodeNum"`

	// <p>可用区列表</p>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>az模式，1: 单AZ 2: 多AZ非主AZ模式 3: 多AZ主AZ模式</p>
	AZMode *int64 `json:"AZMode,omitnil,omitempty" name:"AZMode"`

	// <p>AZMode传3时，表示主AZ</p>
	PrimaryAZ *string `json:"PrimaryAZ,omitnil,omitempty" name:"PrimaryAZ"`

	// <p>全能型副本数，取值范围包括1-3</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`
}

func (r *ExpandInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExpandInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StorageNodeNum")
	delete(f, "Zones")
	delete(f, "AZMode")
	delete(f, "PrimaryAZ")
	delete(f, "FullReplications")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExpandInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExpandInstanceResponseParams struct {
	// <p>异步任务ID，根据此ID可以调用查询任务接口获取任务状态</p>
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExpandInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ExpandInstanceResponseParams `json:"Response"`
}

func (r *ExpandInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExpandInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceFilter struct {
	// 过滤key，支持InstanceId、VpcId、SubnetId、Vip、Vport、Status、InstanceName、TagKey
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤value
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type InstanceInfo struct {
	// <p>计算节点数量</p>
	//
	// Deprecated: ComputeNodeNum is deprecated.
	ComputeNodeNum *int64 `json:"ComputeNodeNum,omitnil,omitempty" name:"ComputeNodeNum"`

	// <p>区域</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>创建实例版本</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateVersion *string `json:"CreateVersion,omitnil,omitempty" name:"CreateVersion"`

	// <p>初始化实例参数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitParams []*InstanceParam `json:"InitParams,omitnil,omitempty" name:"InitParams"`

	// <p>实例状态：creating、created、initializing、running、modifying、isolating、isolated、destroying、destroyed</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>实例id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>存储节点数量</p>
	StorageNodeNum *int64 `json:"StorageNodeNum,omitnil,omitempty" name:"StorageNodeNum"`

	// <p>实例标签信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>实例名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>计算节点cpu核数</p>
	//
	// Deprecated: Cpu is deprecated.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>字符型vpcid</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>计算节点mem，单位GB</p>
	//
	// Deprecated: Mem is deprecated.
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// <p>子网IP</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>字符型subnetid</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>子网端口</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>存储节点磁盘容量，单位GB</p>
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// <p>实例创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>实例所属地域</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>实例状态中文描述</p>
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// <p>管控节点CPU核数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: MCCpu is deprecated.
	MCCpu *int64 `json:"MCCpu,omitnil,omitempty" name:"MCCpu"`

	// <p>管控节点CPU大小</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: MCMem is deprecated.
	MCMem *int64 `json:"MCMem,omitnil,omitempty" name:"MCMem"`

	// <p>计算节点CPU核数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ComputerNodeCpu is deprecated.
	ComputerNodeCpu *int64 `json:"ComputerNodeCpu,omitnil,omitempty" name:"ComputerNodeCpu"`

	// <p>计算节点内存大小</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ComputerNodeMem is deprecated.
	ComputerNodeMem *int64 `json:"ComputerNodeMem,omitnil,omitempty" name:"ComputerNodeMem"`

	// <p>存储节点CPU核数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageNodeCpu *int64 `json:"StorageNodeCpu,omitnil,omitempty" name:"StorageNodeCpu"`

	// <p>存储节点内存大小</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageNodeMem *int64 `json:"StorageNodeMem,omitnil,omitempty" name:"StorageNodeMem"`

	// <p>管控节点数量</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: MCNum is deprecated.
	MCNum *int64 `json:"MCNum,omitnil,omitempty" name:"MCNum"`

	// <p>续费标志</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// <p>付费模式，0按量付费；1包年包月</p>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>用户标签，inner：内部用户；external：外部用户</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountTag *string `json:"AccountTag,omitnil,omitempty" name:"AccountTag"`

	// <p>实例架构类型，separate:分离架构；hyper:对等架构</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>磁盘类型，CLOUD_HSSD增强型SSD,CLOUD_TCS本地SSD盘</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>&quot;0000-00-00 00:00:00&quot;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestroyedAt *string `json:"DestroyedAt,omitnil,omitempty" name:"DestroyedAt"`

	// <p>&quot;0000-00-00 00:00:00&quot;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireAt *string `json:"ExpireAt,omitnil,omitempty" name:"ExpireAt"`

	// <p>&quot;0000-00-00 00:00:00&quot;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolatedAt *string `json:"IsolatedAt,omitnil,omitempty" name:"IsolatedAt"`

	// <p>&quot;0000-00-00 00:00:00&quot;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolatedFrom *string `json:"IsolatedFrom,omitnil,omitempty" name:"IsolatedFrom"`

	// <p>1</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replications *int64 `json:"Replications,omitnil,omitempty" name:"Replications"`

	// <p>全能型副本数</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`

	// <p>账号信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>账号信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitnil,omitempty" name:"SubAccountUin"`

	// <p>账号信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>AZ信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>实例节点</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nodes []*InstanceNode `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// <p>binlog是否打开</p>
	BinlogStatus *int64 `json:"BinlogStatus,omitnil,omitempty" name:"BinlogStatus"`

	// <p>cdc节点核数</p>
	//
	// Deprecated: CdcNodeCpu is deprecated.
	CdcNodeCpu *int64 `json:"CdcNodeCpu,omitnil,omitempty" name:"CdcNodeCpu"`

	// <p>cdc节点内存大小</p>
	//
	// Deprecated: CdcNodeMem is deprecated.
	CdcNodeMem *int64 `json:"CdcNodeMem,omitnil,omitempty" name:"CdcNodeMem"`

	// <p>cdc节点数</p>
	//
	// Deprecated: CdcNodeNum is deprecated.
	CdcNodeNum *int64 `json:"CdcNodeNum,omitnil,omitempty" name:"CdcNodeNum"`

	// <p>az模式，1: 单AZ 2: 多AZ非主AZ模式 3: 多AZ主AZ模式</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AZMode *int64 `json:"AZMode,omitnil,omitempty" name:"AZMode"`

	// <p>灾备标志位，1: 无灾备关系；2: 灾备主实例；3: 灾备备实例</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StandbyFlag *int64 `json:"StandbyFlag,omitnil,omitempty" name:"StandbyFlag"`

	// <p>连接的备实例数量（仅当 StandbyFlag == 2 时有效）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StandbySecondaryNum *int64 `json:"StandbySecondaryNum,omitnil,omitempty" name:"StandbySecondaryNum"`

	// <p>列存节点 cpu 核数</p>
	ColumnarNodeCpu *int64 `json:"ColumnarNodeCpu,omitnil,omitempty" name:"ColumnarNodeCpu"`

	// <p>列存节点内存大小</p>
	ColumnarNodeMem *int64 `json:"ColumnarNodeMem,omitnil,omitempty" name:"ColumnarNodeMem"`

	// <p>列存节点数</p>
	ColumnarNodeNum *int64 `json:"ColumnarNodeNum,omitnil,omitempty" name:"ColumnarNodeNum"`

	// <p>列存节点磁盘容量，单位GB</p>
	ColumnarNodeDisk *int64 `json:"ColumnarNodeDisk,omitnil,omitempty" name:"ColumnarNodeDisk"`

	// <p>列存节点磁盘类型</p>
	ColumnarNodeStorageType *string `json:"ColumnarNodeStorageType,omitnil,omitempty" name:"ColumnarNodeStorageType"`

	// <p>独享标志位，1： 主实例（独享型）, 2: 主实例, 3： 灾备实例, 4： 灾备实例（独享型）</p>
	InstanceCategory *int64 `json:"InstanceCategory,omitnil,omitempty" name:"InstanceCategory"`

	// <p>dbdc-xxxxx</p>
	ExclusiveClusterId *string `json:"ExclusiveClusterId,omitnil,omitempty" name:"ExclusiveClusterId"`

	// <p>兼容模式</p>
	SQLMode *string `json:"SQLMode,omitnil,omitempty" name:"SQLMode"`

	// <p>实例模式</p>
	InstanceMode *string `json:"InstanceMode,omitnil,omitempty" name:"InstanceMode"`

	// <p>实例发货平台</p>
	//
	// Deprecated: ClusterId is deprecated.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>自动扩容配置</p>
	AutoScaleConfig *AutoScalingConfig `json:"AutoScaleConfig,omitnil,omitempty" name:"AutoScaleConfig"`

	// <p>分析引擎模式</p><p>枚举值：</p><ul><li>libra： LibraDB分析引擎模式</li></ul>
	AnalysisMode *string `json:"AnalysisMode,omitnil,omitempty" name:"AnalysisMode"`

	// <p>分析引擎关系信息</p>
	AnalysisRelationInfos []*AnalysisRelationInfo `json:"AnalysisRelationInfos,omitnil,omitempty" name:"AnalysisRelationInfos"`

	// <p>分析引擎实例信息</p>
	AnalysisInstanceInfo *AnalysisInstanceInfo `json:"AnalysisInstanceInfo,omitnil,omitempty" name:"AnalysisInstanceInfo"`
}

type InstanceNode struct {
	// 主键
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 实例Ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 实例EniIp
	// 注意：此字段可能返回 null，表示取不到有效值。
	EniIp *string `json:"EniIp,omitnil,omitempty" name:"EniIp"`

	// 实例Port
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 实例SpecCode
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// 实例NodeName
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 实例Cpu
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例Mem
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 实例Disk
	// 注意：此字段可能返回 null，表示取不到有效值。
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// 实例类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 实例地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例LocalDNS
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalDNS *string `json:"LocalDNS,omitnil,omitempty" name:"LocalDNS"`

	// 实例Region
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 实例日志盘
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogDisk *int64 `json:"LogDisk,omitnil,omitempty" name:"LogDisk"`

	// 实例数据盘
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataDisk *int64 `json:"DataDisk,omitnil,omitempty" name:"DataDisk"`

	// 实例ZoneID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneID *string `json:"ZoneID,omitnil,omitempty" name:"ZoneID"`

	// 实例SpecName
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 实例Replicas
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 实例Shards
	// 注意：此字段可能返回 null，表示取不到有效值。
	Shards *int64 `json:"Shards,omitnil,omitempty" name:"Shards"`

	// 实例数据副本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataReplicas *int64 `json:"DataReplicas,omitnil,omitempty" name:"DataReplicas"`

	// 实例初始化参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// 存储介质, CLOUD_PREMIUM:高性能云盘，CLOUD_SSD: SSD 云盘，CLOUD_HSSD: HSSD 云盘
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`
}

type InstanceParam struct {
	// 配置参数key
	Param *string `json:"Param,omitnil,omitempty" name:"Param"`

	// 配置参数value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type IsolateDBInstanceRequestParams struct {
	// 需要隔离的实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type IsolateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 需要隔离的实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *IsolateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDBInstanceResponseParams struct {
	// 隔离成功实例Id列表
	SuccessInstanceIds []*string `json:"SuccessInstanceIds,omitnil,omitempty" name:"SuccessInstanceIds"`

	// 隔离失败实例Id列表
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitnil,omitempty" name:"FailedInstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateDBInstanceResponseParams `json:"Response"`
}

func (r *IsolateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogBackupStatisticsModel struct {
	// <p>平均每个日志备份大小,单位Byte</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AverageSizePerBackup *int64 `json:"AverageSizePerBackup,omitnil,omitempty" name:"AverageSizePerBackup"`

	// <p>平均每天日志备份大小,单位Byte</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AverageSizePerDay *int64 `json:"AverageSizePerDay,omitnil,omitempty" name:"AverageSizePerDay"`

	// <p>日志备份个数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>日志备份大小，单位Byte</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSize *int64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`
}

type MaintenanceWindowInfo struct {

	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`


	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`


	WeekDays []*string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`
}

// Predefined struct for user
type ModifyAutoRenewFlagRequestParams struct {
	// <p>需要修改的实例列表</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>1表示开启自动续费，0为关闭自动续费</p>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

type ModifyAutoRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// <p>需要修改的实例列表</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>1表示开启自动续费，0为关闭自动续费</p>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

func (r *ModifyAutoRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoRenewFlagResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAutoRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoRenewFlagResponseParams `json:"Response"`
}

func (r *ModifyAutoRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要修改的安全组 ID 列表，一个或者多个安全组 ID 组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要修改的安全组 ID 列表，一个或者多个安全组 ID 组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSecurityGroupsResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBParametersRequestParams struct {
	// 实例 ID，形如：tdsql3-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 参数列表，每一个元素是Param和Value的组合
	Params []*DBParamValue `json:"Params,omitnil,omitempty" name:"Params"`
}

type ModifyDBParametersRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如：tdsql3-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 参数列表，每一个元素是Param和Value的组合
	Params []*DBParamValue `json:"Params,omitnil,omitempty" name:"Params"`
}

func (r *ModifyDBParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Params")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBParametersResponseParams struct {
	// 123
	TaskID *int64 `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBParametersResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBParametersResponseParams `json:"Response"`
}

func (r *ModifyDBParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBSBackupPolicyRequestParams struct {
	// <p>备份策略</p>
	BackupPolicy *BackupPolicyModelInput `json:"BackupPolicy,omitnil,omitempty" name:"BackupPolicy"`

	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ModifyDBSBackupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// <p>备份策略</p>
	BackupPolicy *BackupPolicyModelInput `json:"BackupPolicy,omitnil,omitempty" name:"BackupPolicy"`

	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *ModifyDBSBackupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBSBackupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupPolicy")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBSBackupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBSBackupPolicyResponseParams struct {
	// <p>是否成功</p>
	IsSuccess *bool `json:"IsSuccess,omitnil,omitempty" name:"IsSuccess"`

	// <p>消息</p>
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBSBackupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBSBackupPolicyResponseParams `json:"Response"`
}

func (r *ModifyDBSBackupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBSBackupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBSBackupSetCommentRequestParams struct {
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>备份集ID,值来自 DescribeDBSBackupSets 接口返回</p>
	BackupSetId *int64 `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// <p>备份备注</p>
	NewBackupName *string `json:"NewBackupName,omitnil,omitempty" name:"NewBackupName"`
}

type ModifyDBSBackupSetCommentRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>备份集ID,值来自 DescribeDBSBackupSets 接口返回</p>
	BackupSetId *int64 `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// <p>备份备注</p>
	NewBackupName *string `json:"NewBackupName,omitnil,omitempty" name:"NewBackupName"`
}

func (r *ModifyDBSBackupSetCommentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBSBackupSetCommentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupSetId")
	delete(f, "NewBackupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBSBackupSetCommentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBSBackupSetCommentResponseParams struct {
	// <p>是否成功</p>
	IsSuccess *bool `json:"IsSuccess,omitnil,omitempty" name:"IsSuccess"`

	// <p>修改信息</p>
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBSBackupSetCommentResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBSBackupSetCommentResponseParams `json:"Response"`
}

func (r *ModifyDBSBackupSetCommentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBSBackupSetCommentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNameRequestParams struct {
	// 需要修改的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改的实例名称，要求长度1-60，允许包含中文、英文大小写、数字、-、_
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改的实例名称，要求长度1-60，允许包含中文、英文大小写、数字、-、_
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

func (r *ModifyInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNameResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceNameResponseParams `json:"Response"`
}

func (r *ModifyInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserPrivilegesRequestParams struct {
	// 实例 ID，形如：tdsql3-5baee8df。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 登录用户名和主机信息
	Users []*User `json:"Users,omitnil,omitempty" name:"Users"`

	// 全局权限
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// Database级别权限
	DatabasePrivileges []*DatabasePrivileges `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// Table级别权限
	TablePrivileges []*TablePrivileges `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`
}

type ModifyUserPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如：tdsql3-5baee8df。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 登录用户名和主机信息
	Users []*User `json:"Users,omitnil,omitempty" name:"Users"`

	// 全局权限
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// Database级别权限
	DatabasePrivileges []*DatabasePrivileges `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// Table级别权限
	TablePrivileges []*TablePrivileges `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`
}

func (r *ModifyUserPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Users")
	delete(f, "GlobalPrivileges")
	delete(f, "DatabasePrivileges")
	delete(f, "TablePrivileges")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserPrivilegesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserPrivilegesResponseParams `json:"Response"`
}

func (r *ModifyUserPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeInfo struct {
	// <p>节点IP信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// <p>节点类型，如sqlengine，tdstore，mc</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>节点唯一标识</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// <p>节点端口信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>节点所属可用区</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>节点所属机器ip</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// <p>节点日志服务信息</p>
	BinlogInfo []*BinlogInfo `json:"BinlogInfo,omitnil,omitempty" name:"BinlogInfo"`

	// <p>节点cpu个数</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>节点mem大小</p>
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// <p>节点磁盘大小</p>
	DataDisk *int64 `json:"DataDisk,omitnil,omitempty" name:"DataDisk"`
}

type ParamConstraint struct {
	// 约束类型,如枚举enum，区间section
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 约束类型为enum时的可选值列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enum *string `json:"Enum,omitnil,omitempty" name:"Enum"`

	// 约束类型为section时的范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	Range *ConstraintRange `json:"Range,omitnil,omitempty" name:"Range"`

	// 约束类型为string时的可选值列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	String *string `json:"String,omitnil,omitempty" name:"String"`
}

type ParamDesc struct {
	// 参数名字
	Param *string `json:"Param,omitnil,omitempty" name:"Param"`

	// 当前参数值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 设置过的值，参数生效后，该值和value一样。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SetValue *string `json:"SetValue,omitnil,omitempty" name:"SetValue"`

	// 系统默认值
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// 参数限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	Constraint *ParamConstraint `json:"Constraint,omitnil,omitempty" name:"Constraint"`

	// 是否有设置过值，false:没有设置过值，true:有设置过值。
	HaveSetValue *bool `json:"HaveSetValue,omitnil,omitempty" name:"HaveSetValue"`

	// true:需要重启
	NeedRestart *bool `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// 参数描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ResourceTag struct {
	// 标签键key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type RestartDBInstancesRequestParams struct {
	// <p>需要重启的实例ID的数组</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>重启时间，不传表示立即重启</p>
	RestartTime *string `json:"RestartTime,omitnil,omitempty" name:"RestartTime"`
}

type RestartDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>需要重启的实例ID的数组</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>重启时间，不传表示立即重启</p>
	RestartTime *string `json:"RestartTime,omitnil,omitempty" name:"RestartTime"`
}

func (r *RestartDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "RestartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartDBInstancesResponseParams struct {
	// <p>异步任务id</p>
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RestartDBInstancesResponseParams `json:"Response"`
}

func (r *RestartDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroup struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 创建时间，时间格式：yyyy-mm-dd hh:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 安全组名称
	SecurityGroupName *string `json:"SecurityGroupName,omitnil,omitempty" name:"SecurityGroupName"`

	// 安全组备注
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil,omitempty" name:"SecurityGroupRemark"`

	// 入站规则
	Inbound []*SecurityGroupBound `json:"Inbound,omitnil,omitempty" name:"Inbound"`

	// 出站规则
	Outbound []*SecurityGroupBound `json:"Outbound,omitnil,omitempty" name:"Outbound"`
}

type SecurityGroupBound struct {
	// 来源 IP 或 IP 段，例如192.168.0.0/16
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// 策略，ACCEPT 或者 DROP
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 端口
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// 网络协议，支持 UDP、TCP 等
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`
}

type ServerlessCcu struct {
	// <p>ccu最小值</p>
	MinCcu *int64 `json:"MinCcu,omitnil,omitempty" name:"MinCcu"`

	// <p>ccu最大值范围</p>
	MaxCcu []*int64 `json:"MaxCcu,omitnil,omitempty" name:"MaxCcu"`
}

type StorageNodeSpec struct {
	// <p>规格码</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// <p>存储节点CPU核数</p>
	StorageNodeCpu *int64 `json:"StorageNodeCpu,omitnil,omitempty" name:"StorageNodeCpu"`

	// <p>存储节点内存大小</p>
	StorageNodeMem *int64 `json:"StorageNodeMem,omitnil,omitempty" name:"StorageNodeMem"`

	// <p>存储节点最大数量</p>
	StorageNodeMaxNum *int64 `json:"StorageNodeMaxNum,omitnil,omitempty" name:"StorageNodeMaxNum"`

	// <p>存储节点磁盘大小上限</p>
	StorageNodeMaxDisk *int64 `json:"StorageNodeMaxDisk,omitnil,omitempty" name:"StorageNodeMaxDisk"`

	// <p>存储节点最小数量</p>
	StorageNodeMinNum *int64 `json:"StorageNodeMinNum,omitnil,omitempty" name:"StorageNodeMinNum"`

	// <p>存储节点磁盘大小下限</p>
	StorageNodeMinDisk *int64 `json:"StorageNodeMinDisk,omitnil,omitempty" name:"StorageNodeMinDisk"`

	// <p>磁盘类型，CLOUD_HSSD增强型SSD,CLOUD_TCS本地SSD盘</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>存储节点默认磁盘大小，用于前端展示</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageNodeDefaultDisk *int64 `json:"StorageNodeDefaultDisk,omitnil,omitempty" name:"StorageNodeDefaultDisk"`

	// <p>规格支持计费模式列表</p>
	InstanceMode []*string `json:"InstanceMode,omitnil,omitempty" name:"InstanceMode"`

	// <p>磁盘类型CLOUD_DISK：云盘LOCAL_DISK：本地盘</p>
	DiskTypeCategory *string `json:"DiskTypeCategory,omitnil,omitempty" name:"DiskTypeCategory"`
}

type TablePrivileges struct {
	// DATABASE名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 表名
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 权限列表
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

// Predefined struct for user
type UpgradeInstanceRequestParams struct {
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>实例规格码</p>
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// <p>存储节点磁盘容量，单位GB</p>
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// <p>存储节点CPU核数</p>
	StorageNodeCpu *int64 `json:"StorageNodeCpu,omitnil,omitempty" name:"StorageNodeCpu"`

	// <p>存储节点内存大小</p>
	StorageNodeMem *int64 `json:"StorageNodeMem,omitnil,omitempty" name:"StorageNodeMem"`

	// <p>磁盘类型，需要修改磁盘类型时传</p>
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>实例规格码</p>
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// <p>存储节点磁盘容量，单位GB</p>
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// <p>存储节点CPU核数</p>
	StorageNodeCpu *int64 `json:"StorageNodeCpu,omitnil,omitempty" name:"StorageNodeCpu"`

	// <p>存储节点内存大小</p>
	StorageNodeMem *int64 `json:"StorageNodeMem,omitnil,omitempty" name:"StorageNodeMem"`

	// <p>磁盘类型，需要修改磁盘类型时传</p>
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`
}

func (r *UpgradeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpecCode")
	delete(f, "Disk")
	delete(f, "StorageNodeCpu")
	delete(f, "StorageNodeMem")
	delete(f, "StorageType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceResponseParams struct {
	// <p>任务ID</p>
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeInstanceResponseParams `json:"Response"`
}

func (r *UpgradeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type User struct {
	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 主机
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type UserInfo struct {
	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 主机IP，IP段以%结尾，表示允许该IP段的所有IP
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}