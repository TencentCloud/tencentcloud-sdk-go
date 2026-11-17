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

package v20211108

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type BackupColumnItem struct {
	// 列名。
	ColumnName *string `json:"ColumnName,omitnil,omitempty" name:"ColumnName"`

	// 重命名后的列名。
	NewColumnName *string `json:"NewColumnName,omitnil,omitempty" name:"NewColumnName"`
}

type BackupEndpoint struct {
	// 数据库类型。目前支持的值["mysql", "mariadb", "percona"]。注意，该值必须和备份计划的类型一致。
	DatabaseType *string `json:"DatabaseType,omitnil,omitempty" name:"DatabaseType"`

	// 实例接入类型，支持的值包括：
	// "extranet" - 外网;
	// "cvm" - cvm自建实例;
	// "dcg" - 专线接入;
	// "vpncloud" - 云vpn接入;
	// "cdb" - 腾讯云数据库实例;
	// "ccn" - 云联网接入。
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// 用户名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 登录密码。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 接入地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 服务提供商，支持的值包括["aliyun", "aws", "others"]。
	Supplier *string `json:"Supplier,omitnil,omitempty" name:"Supplier"`

	// 实例 Ip。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 实例端口号。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 云数据库实例ID，格式如：cdb-qcloudtest。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// CVM 实例ID，格式如：ins-olgl39y8，与云服务器控制台页面显示的实例ID相同。如果是CVM自建实例，需要填写该字段。
	CvmInstanceId *string `json:"CvmInstanceId,omitnil,omitempty" name:"CvmInstanceId"`

	// 专线网关ID，格式如：dcg-0rxtqqxb。
	UniqDcgId *string `json:"UniqDcgId,omitnil,omitempty" name:"UniqDcgId"`

	// VPN网关ID，格式如：vpngw-9ghexg7q。
	UniqVpnGwId *string `json:"UniqVpnGwId,omitnil,omitempty" name:"UniqVpnGwId"`

	// 私有网络ID，格式如：vpc-92jblxto。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID，格式如：subnet-3paxmkdz。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 云联网ID，如：ccn-afp6kltc。
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// 数据库版本，当实例为 RDS 或 AWS 实例时才有效，格式如：5.6或者5.7，默认为5.6。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// mariadb三引擎版本。
	DBKernel *string `json:"DBKernel,omitnil,omitempty" name:"DBKernel"`
}

type BackupFile struct {
	// 文件名。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 备份集 ID。
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// 文件状态。取值：
	// - notReady：未就绪
	// - backuping：备份中
	// - finished：备份完成
	// - failed：备份失败
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 源库连接信息。云数据库实例为实例 ID，其余为 ip:port。
	SourceInfo *string `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`

	// 开始时间。格式为 YYYY-MM-DD hh:mm:ss。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。格式为 YYYY-MM-DD hh:mm:ss，备份集尚未结束时返回 0000-00-00 00:00:00。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 过期时间。格式为 YYYY-MM-DD hh:mm:ss。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 文件大小，单位：字节。
	BackupSize *int64 `json:"BackupSize,omitnil,omitempty" name:"BackupSize"`

	// 备份方式。取值：logical（逻辑备份）。
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 存储类型。取值：DBSStorage（内置存储）。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 存储策略类型。取值：
	// - standard：标准存储
	// - tiered：分级存储
	RetentionType *string `json:"RetentionType,omitnil,omitempty" name:"RetentionType"`

	// 备份集存储详情。
	ResourceDetail *ResourceDetail `json:"ResourceDetail,omitnil,omitempty" name:"ResourceDetail"`

	// 备份池。
	BackupPool *string `json:"BackupPool,omitnil,omitempty" name:"BackupPool"`

	// 当前状态下允许的操作。取值：download。
	AllowedActions []*string `json:"AllowedActions,omitnil,omitempty" name:"AllowedActions"`

	// 支持的全部操作。
	AllActions []*string `json:"AllActions,omitnil,omitempty" name:"AllActions"`

	// 文件最后修改时间。格式为 YYYY-MM-DD hh:mm:ss。
	LastModified *string `json:"LastModified,omitnil,omitempty" name:"LastModified"`
}

type BackupObject struct {
	// 备份对象类型，可能的取值为:
	// "all" - 整实例;
	// "partial" - 部分对象。
	ObjectMode *string `json:"ObjectMode,omitnil,omitempty" name:"ObjectMode"`

	// 备份对象详情，当 ObjectMode 为 partial, 即选择部分对象备份时，该字段不能为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectItems []*BackupObjectItem `json:"ObjectItems,omitnil,omitempty" name:"ObjectItems"`
}

type BackupObjectItem struct {
	// 库名。
	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`

	// 重命名后的库名。
	NewDBName *string `json:"NewDBName,omitnil,omitempty" name:"NewDBName"`

	// schema 名。
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 重命名后的 schema 名。
	NewSchemaName *string `json:"NewSchemaName,omitnil,omitempty" name:"NewSchemaName"`

	// 库选择模式，可能的取值为：
	// "all" - 当前对象下的所有对象;
	// "partial" - 当前对象下的部分对象。
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// 表选择模式，可能的取值为:
	// "all" - 当前对象下的所有对象;
	// "partial" - 当前对象下的部分对象。
	TableMode *string `json:"TableMode,omitnil,omitempty" name:"TableMode"`

	// 表对象详情。当 TableMode 为 partial，即选择部分表备份时，此参数需要填写。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tables []*BackupTableItem `json:"Tables,omitnil,omitempty" name:"Tables"`
}

type BackupPeriod struct {
	// 全量备份频率。目前仅支持"Weekly" - 每星期
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 全量备份周期。取值范围为：["Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"]。
	Day []*string `json:"Day,omitnil,omitempty" name:"Day"`
}

type BackupPeriodItem struct {
	// 全量备份频率。目前仅支持"Weekly" - 每星期
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 全量备份周期。取值范围为：["Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"]。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Day []*string `json:"Day,omitnil,omitempty" name:"Day"`
}

type BackupPlanInfo struct {
	// 地域信息。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`

	// 备份计划名称。
	BackupPlanName *string `json:"BackupPlanName,omitnil,omitempty" name:"BackupPlanName"`

	// 备份计划状态。可能的取值为：
	// "notStarted" - 未启动;
	// "checking" - 校验中;
	// "checkPass" - 校验通过;
	// "checkNotPass" - 校验未通过;
	// "running" - 运行中;
	// "fullBacking" - 全量备份中;
	// "isolating" - 隔离中;
	// "isolated" - 已隔离;
	// "offlining" - 下线中;
	// "offlined" - 已下线;
	// "paused" - 已暂停。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 数据库类型。
	DatabaseType *string `json:"DatabaseType,omitnil,omitempty" name:"DatabaseType"`

	// 访问类型。可能的取值为：
	// "extranet" - 外网;
	// "cvm" - cvm 自建实例;
	// "dcg" - 专线接入;
	// "vpncloud" - 云vpn接入;
	// "cdb" - 腾讯云数据库实例;
	// "ccn" - 云联网。
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// 源实例信息。
	SourceInfo []*string `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 到期时间。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 下线时间。
	OfflineTime *string `json:"OfflineTime,omitnil,omitempty" name:"OfflineTime"`

	// 实例规格类型。可能的取值为：["micro", "small", "medium", "large", "xlarge"]。
	InstanceClass *string `json:"InstanceClass,omitnil,omitempty" name:"InstanceClass"`

	// 备份方式。可能的取值为：
	// "logical" - 逻辑备份;
	// "physical" - 物理备份。
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 标签信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 自动续费标记。可能的取值为：
	// 0 - 未开启自动续费;
	// 1 - 已开启自动续费;
	// 2 - 已关闭自动续费。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 是否开启增量备份标记。
	EnableIncrement *bool `json:"EnableIncrement,omitnil,omitempty" name:"EnableIncrement"`

	// 付费类型。可能的取值为：
	// "prePay" - 预付费类型;
	// "postPay" - 后付费类型。
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// 源端信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SetSourceInfo []*string `json:"SetSourceInfo,omitnil,omitempty" name:"SetSourceInfo"`
}

type BackupPlanPayInfo struct {
	// 实例规格。
	InstanceClass *string `json:"InstanceClass,omitnil,omitempty" name:"InstanceClass"`

	// 付费方式。取值：
	// - postPay：后付费
	// - prePay：预付费
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// 免费额度，单位：字节。取值为 -1 表示不限量。
	FreeSpace *int64 `json:"FreeSpace,omitnil,omitempty" name:"FreeSpace"`

	// 本月计费空间，单位：字节。
	TotalChargeSpace *int64 `json:"TotalChargeSpace,omitnil,omitempty" name:"TotalChargeSpace"`

	// 全量计费空间，单位：字节。
	FullChargeSpace *int64 `json:"FullChargeSpace,omitnil,omitempty" name:"FullChargeSpace"`

	// 增量计费空间，单位：字节。
	IncrementChargeSpace *int64 `json:"IncrementChargeSpace,omitnil,omitempty" name:"IncrementChargeSpace"`

	// 本月全量备份集大小，单位：字节。
	FullBackupSetSize *int64 `json:"FullBackupSetSize,omitnil,omitempty" name:"FullBackupSetSize"`

	// 本月增量备份集大小，单位：字节。
	IncrementBackupSetSize *int64 `json:"IncrementBackupSetSize,omitnil,omitempty" name:"IncrementBackupSetSize"`

	// 创建时间。格式为 YYYY-MM-DD HH:mm:ss，如 2024-01-01 00:00:00。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 统计周期，格式为“开始时间 - 结束时间”。
	TimeRange *string `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`
}

type BackupProgress struct {
	// 全量备份状态。
	FullBackupStatus *string `json:"FullBackupStatus,omitnil,omitempty" name:"FullBackupStatus"`

	// 增量备份状态。
	IncrementBackupStatus *string `json:"IncrementBackupStatus,omitnil,omitempty" name:"IncrementBackupStatus"`

	// 上次备份执行时间。格式为 YYYY-MM-DD HH:mm:ss，如 2024-01-01 00:00:00，尚未执行过备份时为空字符串。
	LastBackupTime *string `json:"LastBackupTime,omitnil,omitempty" name:"LastBackupTime"`

	// 下次备份执行时间。格式为 YYYY-MM-DD HH:mm:ss，如 2024-01-01 00:00:00，尚未确定时为空字符串。
	NextBackupTime *string `json:"NextBackupTime,omitnil,omitempty" name:"NextBackupTime"`

	// 上次恢复时间。格式为 YYYY-MM-DD HH:mm:ss，如 2024-01-01 00:00:00，尚未执行过恢复时为空字符串。
	LastRestoreTime *string `json:"LastRestoreTime,omitnil,omitempty" name:"LastRestoreTime"`

	// 上次恢复任务状态。
	LastRestoreStatus *string `json:"LastRestoreStatus,omitnil,omitempty" name:"LastRestoreStatus"`

	// 当前备份作业状态。
	CurrentBackupJobStatus *string `json:"CurrentBackupJobStatus,omitnil,omitempty" name:"CurrentBackupJobStatus"`
}

type BackupSetOption struct {
	// <p>存储加密方式。取值：<br>&quot;UnEncrypted&quot; - 非加密存储;<br>&quot;SSE-COS&quot; - 内置加密存储;<br>&quot;SSE-KMS&quot; - KMS 加密存储。</p>
	Encryption *string `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type BackupStrategy struct {
	// 全量备份开始时间。周期性的全量备份将在当天该时间开始。
	BackupStartTime *string `json:"BackupStartTime,omitnil,omitempty" name:"BackupStartTime"`

	// 存储策略。
	StorageStrategy *StorageStrategy `json:"StorageStrategy,omitnil,omitempty" name:"StorageStrategy"`

	// 备份周期。
	BackupPeriod *BackupPeriod `json:"BackupPeriod,omitnil,omitempty" name:"BackupPeriod"`

	// 备份方法。目前仅支持 "logical" - 逻辑备份。
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 备份周期。支持的值包括：
	// "period" - 周期性备份；
	// "single" - 单次备份。
	// 默认值为"period"。
	StrategyType *string `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 是否开启增量备份。可能的取值为["true", "false"]。默认值为"true"。
	EnableIncrement *bool `json:"EnableIncrement,omitnil,omitempty" name:"EnableIncrement"`
}

type BackupStrategyItem struct {
	// 全量备份开始时间。周期性的全量备份将在当天该时间开始，格式：HH:mm（24 小时制），如 02:00。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupStartTime *string `json:"BackupStartTime,omitnil,omitempty" name:"BackupStartTime"`

	// 存储策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageStrategy *StorageStrategyItem `json:"StorageStrategy,omitnil,omitempty" name:"StorageStrategy"`

	// 备份周期。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupPeriod *BackupPeriodItem `json:"BackupPeriod,omitnil,omitempty" name:"BackupPeriod"`

	// 备份方法。目前仅支持 "logical" - 逻辑备份。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 备份周期。支持的值包括：
	// "period" - 周期性备份；
	// "single" - 单次备份。
	// 默认值为"period"。
	StrategyType *string `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 是否开启增量备份。可能的取值为["true", "false"]。默认值为"true"。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableIncrement *bool `json:"EnableIncrement,omitnil,omitempty" name:"EnableIncrement"`
}

type BackupTableItem struct {
	// 表名。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 重命名后的表名。
	NewTableName *string `json:"NewTableName,omitnil,omitempty" name:"NewTableName"`

	// 列对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*BackupColumnItem `json:"Columns,omitnil,omitempty" name:"Columns"`
}

// Predefined struct for user
type ConfigureBackupPlanRequestParams struct {
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`

	// 备份计划名称。支持数字、英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@,且长度不能超过60。
	BackupPlanName *string `json:"BackupPlanName,omitnil,omitempty" name:"BackupPlanName"`

	// 全量备份并发数上限。
	UpperParallel *int64 `json:"UpperParallel,omitnil,omitempty" name:"UpperParallel"`

	// 备份源实例信息。
	SourceEndPoint *BackupEndpoint `json:"SourceEndPoint,omitnil,omitempty" name:"SourceEndPoint"`

	// 备份对象信息。
	BackupObject *BackupObject `json:"BackupObject,omitnil,omitempty" name:"BackupObject"`

	// 备份策略。
	BackupStrategy *BackupStrategy `json:"BackupStrategy,omitnil,omitempty" name:"BackupStrategy"`

	// 加密信息。当需要使用SSE-KMS需要传入该值，你可以通过 KMS 的 GenerateDataKey 接口生成。
	PlainText *string `json:"PlainText,omitnil,omitempty" name:"PlainText"`
}

type ConfigureBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`

	// 备份计划名称。支持数字、英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@,且长度不能超过60。
	BackupPlanName *string `json:"BackupPlanName,omitnil,omitempty" name:"BackupPlanName"`

	// 全量备份并发数上限。
	UpperParallel *int64 `json:"UpperParallel,omitnil,omitempty" name:"UpperParallel"`

	// 备份源实例信息。
	SourceEndPoint *BackupEndpoint `json:"SourceEndPoint,omitnil,omitempty" name:"SourceEndPoint"`

	// 备份对象信息。
	BackupObject *BackupObject `json:"BackupObject,omitnil,omitempty" name:"BackupObject"`

	// 备份策略。
	BackupStrategy *BackupStrategy `json:"BackupStrategy,omitnil,omitempty" name:"BackupStrategy"`

	// 加密信息。当需要使用SSE-KMS需要传入该值，你可以通过 KMS 的 GenerateDataKey 接口生成。
	PlainText *string `json:"PlainText,omitnil,omitempty" name:"PlainText"`
}

func (r *ConfigureBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfigureBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupPlanId")
	delete(f, "BackupPlanName")
	delete(f, "UpperParallel")
	delete(f, "SourceEndPoint")
	delete(f, "BackupObject")
	delete(f, "BackupStrategy")
	delete(f, "PlainText")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConfigureBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfigureBackupPlanResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ConfigureBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *ConfigureBackupPlanResponseParams `json:"Response"`
}

func (r *ConfigureBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfigureBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConnectTestResult struct {
	// <p>任务 ID</p>
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>任务状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>是否通过。0 表示未通过，1 表示通过。</p>
	IsPass *int64 `json:"IsPass,omitnil,omitempty" name:"IsPass"`

	// <p>源端地址</p>
	Addr *string `json:"Addr,omitnil,omitempty" name:"Addr"`

	// <p>源地址转换IP</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SNatIp *string `json:"SNatIp,omitnil,omitempty" name:"SNatIp"`

	// <p>检测结果集</p>
	TestItems []*TestItem `json:"TestItems,omitnil,omitempty" name:"TestItems"`
}

// Predefined struct for user
type CreateBackupPlanRequestParams struct {
	// 源端数据库类型。当前支持值为: ["mysql","cynosdbmysql","percona","mariadb","tdsqlmysql"]。
	DatabaseType *string `json:"DatabaseType,omitnil,omitempty" name:"DatabaseType"`

	// 备份方式。当前仅支持"logical"，即逻辑备份。
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 规格。当前支持值为: ["micro","small","medium","large","xlarge"]。默认为"small"。
	InstanceClass *string `json:"InstanceClass,omitnil,omitempty" name:"InstanceClass"`

	// 购买时长，单位为月，默认值为1。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 计费模式。当前仅支持"prepay"，即包年包月。
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// 购买数量。取值范围为[1, 10]，默认值为1。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 自动续费标识。1 - 开启自动续费；0 - 不开启自动续费。
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 标签值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// 源端数据库类型。当前支持值为: ["mysql","cynosdbmysql","percona","mariadb","tdsqlmysql"]。
	DatabaseType *string `json:"DatabaseType,omitnil,omitempty" name:"DatabaseType"`

	// 备份方式。当前仅支持"logical"，即逻辑备份。
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 规格。当前支持值为: ["micro","small","medium","large","xlarge"]。默认为"small"。
	InstanceClass *string `json:"InstanceClass,omitnil,omitempty" name:"InstanceClass"`

	// 购买时长，单位为月，默认值为1。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 计费模式。当前仅支持"prepay"，即包年包月。
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// 购买数量。取值范围为[1, 10]，默认值为1。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 自动续费标识。1 - 开启自动续费；0 - 不开启自动续费。
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 标签值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseType")
	delete(f, "BackupMethod")
	delete(f, "InstanceClass")
	delete(f, "Period")
	delete(f, "PayType")
	delete(f, "Count")
	delete(f, "AutoRenew")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupPlanResponseParams struct {
	// 订单参数。
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 备份计划 ID 列表。单次最多返回 10 个。
	BackupPlanIds []*string `json:"BackupPlanIds,omitnil,omitempty" name:"BackupPlanIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupPlanResponseParams `json:"Response"`
}

func (r *CreateBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConnectTestJobRequestParams struct {
	// 备份源实例信息。
	Endpoint *BackupEndpoint `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

type CreateConnectTestJobRequest struct {
	*tchttp.BaseRequest
	
	// 备份源实例信息。
	Endpoint *BackupEndpoint `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

func (r *CreateConnectTestJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConnectTestJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Endpoint")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConnectTestJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConnectTestJobResponseParams struct {
	// 连通性任务 ID。
	ConnTaskId *string `json:"ConnTaskId,omitnil,omitempty" name:"ConnTaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConnectTestJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateConnectTestJobResponseParams `json:"Response"`
}

func (r *CreateConnectTestJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConnectTestJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupPlanRequestParams struct {
	// 备份计划 ID。可通过 DescribeBackupPlans 接口获取。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

type DeleteBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// 备份计划 ID。可通过 DescribeBackupPlans 接口获取。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

func (r *DeleteBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupPlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupPlanResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBackupPlanResponseParams `json:"Response"`
}

func (r *DeleteBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupCheckJobRequestParams struct {
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

type DescribeBackupCheckJobRequest struct {
	*tchttp.BaseRequest
	
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

func (r *DescribeBackupCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupPlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupCheckJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupCheckJobResponseParams struct {
	// 校验任务状态。可能的取值为："finished" - 已完成; "running" - 进行中。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务进度。取值范围为[0, 100]，表示任务完成的百分比。例如：30表示任务完成30%。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 校验是否通过标记。可能的取值为："1" - 校验通过;"0" - 校验未通过。
	CheckFlag *int64 `json:"CheckFlag,omitnil,omitempty" name:"CheckFlag"`

	// 错误信息。
	ErrMessage *string `json:"ErrMessage,omitnil,omitempty" name:"ErrMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupCheckJobResponseParams `json:"Response"`
}

func (r *DescribeBackupCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupPlanDetailRequestParams struct {
	// <p>备份计划 ID。可通过 DescribeBackupPlans 接口获取。</p>
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

type DescribeBackupPlanDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>备份计划 ID。可通过 DescribeBackupPlans 接口获取。</p>
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

func (r *DescribeBackupPlanDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupPlanDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupPlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupPlanDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupPlanDetailResponseParams struct {
	// <p>地域。</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>备份计划 ID。</p>
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`

	// <p>备份计划名称。</p>
	BackupPlanName *string `json:"BackupPlanName,omitnil,omitempty" name:"BackupPlanName"`

	// <p>备份计划状态。取值：</p><ul><li>creating：创建中</li><li>notStarted：未启动</li><li>notConfigured：未配置</li><li>checkNotPass：校验未通过</li><li>checkPass：校验通过</li><li>checking：校验中</li><li>readyRun：待运行</li><li>running：运行中</li><li>fullBacking：全量备份中</li><li>paused：已暂停</li><li>isolating：隔离中</li><li>isolated：已隔离</li><li>offlining：下线中</li><li>offlined：已下线</li><li>completed：已完成</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>源库连接信息列表。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceEndPoint []*Endpoint `json:"SourceEndPoint,omitnil,omitempty" name:"SourceEndPoint"`

	// <p>备份对象信息。</p>
	BackupObject *BackupObject `json:"BackupObject,omitnil,omitempty" name:"BackupObject"`

	// <p>备份策略。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupStrategy *BackupStrategyItem `json:"BackupStrategy,omitnil,omitempty" name:"BackupStrategy"`

	// <p>备份进度信息。</p>
	Progress *BackupProgress `json:"Progress,omitnil,omitempty" name:"Progress"`

	// <p>错误异常信息。</p>
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>创建时间。格式为 YYYY-MM-DD HH:mm:ss，如 2024-01-01 00:00:00。</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>到期时间。格式为 YYYY-MM-DD HH:mm:ss，如 2024-01-01 00:00:00，后付费计划无到期时间时为空字符串。</p>
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// <p>下线时间。格式为 YYYY-MM-DD HH:mm:ss，如 2024-01-01 00:00:00，尚未下线时为空字符串。</p>
	OfflineTime *string `json:"OfflineTime,omitnil,omitempty" name:"OfflineTime"`

	// <p>自动续费标识。取值：</p><ul><li>0：不自动续费</li><li>1：自动续费</li></ul>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// <p>主节点连接信息。云数据库实例为实例 ID，云服务器自建为实例 ID:端口，其余为 ip:port。</p>
	SourceInfo []*string `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`

	// <p>从节点连接信息，格式同 SourceInfo。</p>
	SetSourceInfo []*string `json:"SetSourceInfo,omitnil,omitempty" name:"SetSourceInfo"`

	// <p>标签信息列表。未绑定标签时返回 null。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>上次备份作业时间。格式为 YYYY-MM-DD HH:mm:ss，如 2024-01-01 00:00:00，尚未执行过备份作业时为空字符串。</p>
	LastBackupJobTime *string `json:"LastBackupJobTime,omitnil,omitempty" name:"LastBackupJobTime"`

	// <p>上次全量备份作业状态。</p>
	LastFullBackupJobStatus *string `json:"LastFullBackupJobStatus,omitnil,omitempty" name:"LastFullBackupJobStatus"`

	// <p>上次增量备份作业状态。</p>
	LastIncBackupJobStatus *string `json:"LastIncBackupJobStatus,omitnil,omitempty" name:"LastIncBackupJobStatus"`

	// <p>下次备份作业时间。格式为 YYYY-MM-DD HH:mm:ss，如 2024-01-01 00:00:00，尚未确定时为空字符串。</p>
	NextBackupJobTime *string `json:"NextBackupJobTime,omitnil,omitempty" name:"NextBackupJobTime"`

	// <p>上次恢复作业时间。格式为 YYYY-MM-DD HH:mm:ss，如 2024-01-01 00:00:00，尚未执行过恢复作业时为空字符串。</p>
	LastRestoreJobTime *string `json:"LastRestoreJobTime,omitnil,omitempty" name:"LastRestoreJobTime"`

	// <p>上次恢复作业状态。</p>
	LastRestoreJobStatus *string `json:"LastRestoreJobStatus,omitnil,omitempty" name:"LastRestoreJobStatus"`

	// <p>付费信息。</p>
	PayInfo *BackupPlanPayInfo `json:"PayInfo,omitnil,omitempty" name:"PayInfo"`

	// <p>全量备份并行数上限。</p>
	UpperParallel *int64 `json:"UpperParallel,omitnil,omitempty" name:"UpperParallel"`

	// <p>带宽上限，单位：MB/s。</p>
	UpperBandwidth *int64 `json:"UpperBandwidth,omitnil,omitempty" name:"UpperBandwidth"`

	// <p>当前状态下允许的操作。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowedActions []*string `json:"AllowedActions,omitnil,omitempty" name:"AllowedActions"`

	// <p>可恢复时间范围。</p>
	RestoreRange *RestoreRange `json:"RestoreRange,omitnil,omitempty" name:"RestoreRange"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupPlanDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupPlanDetailResponseParams `json:"Response"`
}

func (r *DescribeBackupPlanDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupPlanDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupPlansRequestParams struct {
	// 过滤条件，备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`

	// 过滤条件，备份计划状态。
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// 过滤条件，数据库类型。
	DatabaseType []*string `json:"DatabaseType,omitnil,omitempty" name:"DatabaseType"`

	// 过滤条件，接入访问类型。
	AccessType []*string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// 过滤条件，备份计划名称。
	BackupPlanName *string `json:"BackupPlanName,omitnil,omitempty" name:"BackupPlanName"`

	// 过滤条件，标签键值。
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 分页参数。取值范围为(0, 100]，默认值为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数。默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeBackupPlansRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`

	// 过滤条件，备份计划状态。
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// 过滤条件，数据库类型。
	DatabaseType []*string `json:"DatabaseType,omitnil,omitempty" name:"DatabaseType"`

	// 过滤条件，接入访问类型。
	AccessType []*string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// 过滤条件，备份计划名称。
	BackupPlanName *string `json:"BackupPlanName,omitnil,omitempty" name:"BackupPlanName"`

	// 过滤条件，标签键值。
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 分页参数。取值范围为(0, 100]，默认值为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数。默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	delete(f, "BackupPlanId")
	delete(f, "Status")
	delete(f, "DatabaseType")
	delete(f, "AccessType")
	delete(f, "BackupPlanName")
	delete(f, "TagFilters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupPlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupPlansResponseParams struct {
	// 备份计划数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 备份计划详情。
	Items []*BackupPlanInfo `json:"Items,omitnil,omitempty" name:"Items"`

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
type DescribeConnectTestResultRequestParams struct {
	// <p>连通性检测任务 ID。</p>
	TaskIds []*int64 `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

type DescribeConnectTestResultRequest struct {
	*tchttp.BaseRequest
	
	// <p>连通性检测任务 ID。</p>
	TaskIds []*int64 `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

func (r *DescribeConnectTestResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConnectTestResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConnectTestResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConnectTestResultResponseParams struct {
	// <p>任务总数。</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>检测结果详情。</p>
	Items []*ConnectTestResult `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConnectTestResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConnectTestResultResponseParams `json:"Response"`
}

func (r *DescribeConnectTestResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConnectTestResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFullBackupSetDetailRequestParams struct {
	// <p>全量备份集 ID。可通过 DescribeFullBackupSets 接口获取。</p>
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`
}

type DescribeFullBackupSetDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>全量备份集 ID。可通过 DescribeFullBackupSets 接口获取。</p>
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`
}

func (r *DescribeFullBackupSetDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFullBackupSetDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupSetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFullBackupSetDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFullBackupSetDetailResponseParams struct {
	// <p>全量备份集详情。</p>
	BackupSet *FullBackupSetDetail `json:"BackupSet,omitnil,omitempty" name:"BackupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFullBackupSetDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFullBackupSetDetailResponseParams `json:"Response"`
}

func (r *DescribeFullBackupSetDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFullBackupSetDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIncrementBackupSetFilesRequestParams struct {
	// 增量备份集 ID。可通过 DescribeIncrementBackupSets 接口获取。
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 10。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeIncrementBackupSetFilesRequest struct {
	*tchttp.BaseRequest
	
	// 增量备份集 ID。可通过 DescribeIncrementBackupSets 接口获取。
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 10。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeIncrementBackupSetFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIncrementBackupSetFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupSetId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIncrementBackupSetFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIncrementBackupSetFilesResponseParams struct {
	// 文件总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 增量备份文件列表。
	BackupFiles []*BackupFile `json:"BackupFiles,omitnil,omitempty" name:"BackupFiles"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIncrementBackupSetFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIncrementBackupSetFilesResponseParams `json:"Response"`
}

func (r *DescribeIncrementBackupSetFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIncrementBackupSetFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIncrementBackupSetsRequestParams struct {
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`

	// 过滤条件，备份集 ID。
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// 分页参数。取值范围为(0, 100]，默认值为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数。默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeIncrementBackupSetsRequest struct {
	*tchttp.BaseRequest
	
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`

	// 过滤条件，备份集 ID。
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// 分页参数。取值范围为(0, 100]，默认值为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数。默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeIncrementBackupSetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIncrementBackupSetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupPlanId")
	delete(f, "BackupSetId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIncrementBackupSetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIncrementBackupSetsResponseParams struct {
	// 增量备份集数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 增量备份集信息。
	Items []*IncrementBackup `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIncrementBackupSetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIncrementBackupSetsResponseParams `json:"Response"`
}

func (r *DescribeIncrementBackupSetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIncrementBackupSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyBackupPlanByUserRequestParams struct {
	// 备份计划 ID。可通过 DescribeBackupPlans 接口获取。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

type DestroyBackupPlanByUserRequest struct {
	*tchttp.BaseRequest
	
	// 备份计划 ID。可通过 DescribeBackupPlans 接口获取。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

func (r *DestroyBackupPlanByUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyBackupPlanByUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupPlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyBackupPlanByUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyBackupPlanByUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyBackupPlanByUserResponse struct {
	*tchttp.BaseResponse
	Response *DestroyBackupPlanByUserResponseParams `json:"Response"`
}

func (r *DestroyBackupPlanByUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyBackupPlanByUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Endpoint struct {
	// 节点角色。取值：
	// "master" - 主节点;
	// "slave" - 从节点。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 数据库类型。目前支持的值["mysql", "mariadb", "percona"]。注意，该值必须和备份计划的类型一致。
	DatabaseType *string `json:"DatabaseType,omitnil,omitempty" name:"DatabaseType"`

	// 实例接入类型，支持的值包括：
	// "extranet" - 外网;
	// "cvm" - cvm自建实例;
	// "dcg" - 专线接入;
	// "vpncloud" - 云vpn接入;
	// "cdb" - 腾讯云数据库实例;
	// "ccn" - 云联网接入。
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// 实例 Ip。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 实例端口号。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 用户名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 登录密码。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 云数据库实例ID，格式如：cdb-qcloudtest。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// CVM 实例ID，格式如：ins-olgl39y8，与云服务器控制台页面显示的实例ID相同。如果是CVM自建实例，需要填写该字段。
	CvmInstanceId *string `json:"CvmInstanceId,omitnil,omitempty" name:"CvmInstanceId"`

	// 专线网关ID，格式如：dcg-0rxtqqxb。
	UniqDcgId *string `json:"UniqDcgId,omitnil,omitempty" name:"UniqDcgId"`

	// VPN网关ID，格式如：vpngw-9ghexg7q。
	UniqVpnGwId *string `json:"UniqVpnGwId,omitnil,omitempty" name:"UniqVpnGwId"`

	// 私有网络ID，格式如：vpc-92jblxto。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID，格式如：subnet-3paxmkdz。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 云联网ID，如：ccn-afp6kltc。
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// 接入地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 服务提供商，支持的值包括["aliyun", "aws", "others"]。
	Supplier *string `json:"Supplier,omitnil,omitempty" name:"Supplier"`

	// 数据库版本，当实例为 RDS 或 AWS 实例时才有效，格式如：5.6或者5.7，默认为5.6。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// mariadb三引擎版本。
	DBKernel *string `json:"DBKernel,omitnil,omitempty" name:"DBKernel"`
}

type ErrorInfo struct {
	// 错误类型。
	ErrorType *string `json:"ErrorType,omitnil,omitempty" name:"ErrorType"`

	// 错误信息。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`
}

type FullBackupErrorMessage struct {
	// 错误标题。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 错误详情。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type FullBackupProgress struct {
	// 结构备份进度，取值范围为 0 - 100，单位：百分比。
	SchemaProgress *int64 `json:"SchemaProgress,omitnil,omitempty" name:"SchemaProgress"`

	// 数据备份进度，取值范围为 0 - 100，单位：百分比。
	DataProgress *int64 `json:"DataProgress,omitnil,omitempty" name:"DataProgress"`
}

type FullBackupSetDetail struct {
	// <p>备份集 ID。</p>
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// <p>源实例信息。</p>
	SourceInfo *string `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`

	// <p>备份集状态。可能的取值为："running" - 运行中;"finished" - 已完成;"failed" - 已失败;"stopped" - 已终止。</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>创建时间。格式为 YYYY-MM-DD hh:mm:ss。</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间。格式为 YYYY-MM-DD hh:mm:ss，备份集尚未结束时返回 0000-00-00 00:00:00。</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>过期时间。格式为 YYYY-MM-DD hh:mm:ss。</p>
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// <p>备份集大小，单位为 Bytes。</p>
	BackupSize *int64 `json:"BackupSize,omitnil,omitempty" name:"BackupSize"`

	// <p>备份方式。可能的取值为："logical" - 逻辑备份；"physical" - 物理备份。</p>
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// <p>存储类型。目前仅支持 "DBSStorage"，即 DBS 内置存储。</p>
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>备份池类型。目前仅支持 "DefaultPool"，即默认存储池。</p>
	BackupPool *string `json:"BackupPool,omitnil,omitempty" name:"BackupPool"`

	// <p>备份集存储类型。支持的取值范围为："standard" - 标准存储；"tiered" - 分级存储。</p>
	RetentionType *string `json:"RetentionType,omitnil,omitempty" name:"RetentionType"`

	// <p>备份集存储状态详情。</p>
	ResourceDetail *ResourceDetail `json:"ResourceDetail,omitnil,omitempty" name:"ResourceDetail"`

	// 备份源实例信息。
	SourceEndpoint *Endpoint `json:"SourceEndpoint,omitnil,omitempty" name:"SourceEndpoint"`

	// 备份对象。
	BackupObject *BackupObject `json:"BackupObject,omitnil,omitempty" name:"BackupObject"`

	// 备份进度。
	Progress *FullBackupProgress `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 备份错误信息。备份失败时有值。
	ErrorMessage *FullBackupErrorMessage `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 备份集加密信息。
	BackupSetOption *BackupSetOption `json:"BackupSetOption,omitnil,omitempty" name:"BackupSetOption"`
}

type IncrementBackup struct {
	// 备份集 ID。
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// 源实例信息。
	SourceInfo *string `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`

	// 备份集状态。可能的取值为：
	// "running" - 运行中;
	// "finished" - 已完成;
	// "failed" - 已失败;
	// "stopped" - 已终止。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 开始时间。格式为 YYYY-MM-DD hh:mm:ss。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。格式为 YYYY-MM-DD hh:mm:ss，备份集尚未结束时返回 0000-00-00 00:00:00。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 过期时间。格式为 YYYY-MM-DD hh:mm:ss，备份集尚未生成过期时间时返回空字符串。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 备份集文件个数。
	FileCount *int64 `json:"FileCount,omitnil,omitempty" name:"FileCount"`

	// 备份方式。可能的取值为：
	// "logical" - 逻辑备份；
	// "physical" - 物理备份。
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 存储类型。目前仅支持 "DBSStorage"，即 DBS 内置存储。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 备份池类型。目前仅支持 "DefaultPool"，即默认存储池。
	BackupPool *string `json:"BackupPool,omitnil,omitempty" name:"BackupPool"`

	// 备份集存储类型。支持的取值范围为："standard" - 标准存储；"tiered" - 分级存储。
	RetentionType *string `json:"RetentionType,omitnil,omitempty" name:"RetentionType"`

	// 备份集存储状态详情。
	ResourceDetail *ResourceItemDetail `json:"ResourceDetail,omitnil,omitempty" name:"ResourceDetail"`
}

// Predefined struct for user
type ModifyBackupPlanSourceInfoRequestParams struct {
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`

	// 新的备份源实例信息。
	SourceEndPoint *BackupEndpoint `json:"SourceEndPoint,omitnil,omitempty" name:"SourceEndPoint"`

	// 新的备份名称。
	BackupPlanName *string `json:"BackupPlanName,omitnil,omitempty" name:"BackupPlanName"`

	// 新的备份对象信息。
	BackupObject *BackupObject `json:"BackupObject,omitnil,omitempty" name:"BackupObject"`

	// 是否立即启动新的备份任务。支持的取值为：1 - 立即启动；0 - 按备份策略配置的时间启动。
	BackupNow *int64 `json:"BackupNow,omitnil,omitempty" name:"BackupNow"`
}

type ModifyBackupPlanSourceInfoRequest struct {
	*tchttp.BaseRequest
	
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`

	// 新的备份源实例信息。
	SourceEndPoint *BackupEndpoint `json:"SourceEndPoint,omitnil,omitempty" name:"SourceEndPoint"`

	// 新的备份名称。
	BackupPlanName *string `json:"BackupPlanName,omitnil,omitempty" name:"BackupPlanName"`

	// 新的备份对象信息。
	BackupObject *BackupObject `json:"BackupObject,omitnil,omitempty" name:"BackupObject"`

	// 是否立即启动新的备份任务。支持的取值为：1 - 立即启动；0 - 按备份策略配置的时间启动。
	BackupNow *int64 `json:"BackupNow,omitnil,omitempty" name:"BackupNow"`
}

func (r *ModifyBackupPlanSourceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupPlanSourceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupPlanId")
	delete(f, "SourceEndPoint")
	delete(f, "BackupPlanName")
	delete(f, "BackupObject")
	delete(f, "BackupNow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupPlanSourceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupPlanSourceInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBackupPlanSourceInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupPlanSourceInfoResponseParams `json:"Response"`
}

func (r *ModifyBackupPlanSourceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupPlanSourceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceDetail struct {
	// <p>当前存储状态。可能的取值为："standard" - 标准存储；"lowFrequency" - 低频存储；"archive" - 归档存储；"deepArchive" - 深度归档存储；</p>
	TieredType *string `json:"TieredType,omitnil,omitempty" name:"TieredType"`

	// <p>存储时长。备份集已保留的时长，单位为小时。</p>
	StayTime *float64 `json:"StayTime,omitnil,omitempty" name:"StayTime"`
}

type ResourceItemDetail struct {
	// <p>当前存储状态。可能的取值为："standard" - 标准存储；"lowFrequency" - 低频存储；"archive" - 归档存储；"deepArchive" - 深度归档存储；</p>
	TieredType *string `json:"TieredType,omitnil,omitempty" name:"TieredType"`

	// <p>存储时长。备份集已保留的时长，单位为小时。</p>
	StayTime *float64 `json:"StayTime,omitnil,omitempty" name:"StayTime"`
}

type RestoreRange struct {
	// 可恢复起始时间。格式为 YYYY-MM-DD HH:mm:ss，如 2024-01-01 00:00:00，无可恢复时间范围时为空字符串。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 可恢复结束时间。格式为 YYYY-MM-DD HH:mm:ss，如 2024-01-01 00:00:00，无可恢复时间范围时为空字符串。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

// Predefined struct for user
type ResumeBackupPlanRequestParams struct {
	// 备份计划 ID。可通过 DescribeBackupPlans 接口获取。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`

	// 恢复后是否立即发起一次备份。取值：
	// - 0：不立即备份，默认值
	// - 1：立即备份
	BackupNow *int64 `json:"BackupNow,omitnil,omitempty" name:"BackupNow"`
}

type ResumeBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// 备份计划 ID。可通过 DescribeBackupPlans 接口获取。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`

	// 恢复后是否立即发起一次备份。取值：
	// - 0：不立即备份，默认值
	// - 1：立即备份
	BackupNow *int64 `json:"BackupNow,omitnil,omitempty" name:"BackupNow"`
}

func (r *ResumeBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupPlanId")
	delete(f, "BackupNow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeBackupPlanResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *ResumeBackupPlanResponseParams `json:"Response"`
}

func (r *ResumeBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartBackupCheckJobRequestParams struct {
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

type StartBackupCheckJobRequest struct {
	*tchttp.BaseRequest
	
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

func (r *StartBackupCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBackupCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupPlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartBackupCheckJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartBackupCheckJobResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartBackupCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *StartBackupCheckJobResponseParams `json:"Response"`
}

func (r *StartBackupCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBackupCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartBackupPlanRequestParams struct {
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

type StartBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

func (r *StartBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupPlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartBackupPlanResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *StartBackupPlanResponseParams `json:"Response"`
}

func (r *StartBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopBackupPlanRequestParams struct {
	// 备份计划 ID。可通过 DescribeBackupPlans 接口获取。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

type StopBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// 备份计划 ID。可通过 DescribeBackupPlans 接口获取。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

func (r *StopBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupPlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopBackupPlanResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *StopBackupPlanResponseParams `json:"Response"`
}

func (r *StopBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageStrategy struct {
	// 存储类型。目前仅支持 "system" - DBS 内置存储。默认值为 "system"。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 加密方式。可能的取值为：
	// "UnEncrypted" - 非加密存储;
	// "SSE-COS" - 内置加密存储;
	// 当该参数用作入参时，默认值为 "UnEncrypted"。
	Encryption *string `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// 日志保留时间，单位为天。取值范围为[7, 3650]，默认值为 30。
	BackupRetentionPeriod *int64 `json:"BackupRetentionPeriod,omitnil,omitempty" name:"BackupRetentionPeriod"`
}

type StorageStrategyItem struct {
	// 存储类型。目前仅支持 "system" - DBS 内置存储。默认值为 "system"。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 加密方式。可能的取值为：
	// "UnEncrypted" - 非加密存储;
	// "SSE-COS" - 内置加密存储;
	// 当该参数用作入参时，默认值为 "UnEncrypted"。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Encryption *string `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// 日志保留时间，单位为天。取值范围为[7, 3650]，默认值为 30。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupRetentionPeriod *int64 `json:"BackupRetentionPeriod,omitnil,omitempty" name:"BackupRetentionPeriod"`

	// 全量备份存储策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullStorageStrategy *StorageStrategySingle `json:"FullStorageStrategy,omitnil,omitempty" name:"FullStorageStrategy"`

	// 增量备份存储策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncStorageStrategy *StorageStrategySingle `json:"IncStorageStrategy,omitnil,omitempty" name:"IncStorageStrategy"`
}

type StorageStrategySingle struct {
	// 存储策略类型。取值：
	// - standard：标准存储
	// - tiered：分级存储
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetentionType *string `json:"RetentionType,omitnil,omitempty" name:"RetentionType"`

	// 备份保留天数，取值范围为 7 - 3650。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetentionPeriod *int64 `json:"RetentionPeriod,omitnil,omitempty" name:"RetentionPeriod"`

	// 分级存储策略。RetentionType 为 tiered 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TieredStrategy *StorageTieredStrategy `json:"TieredStrategy,omitnil,omitempty" name:"TieredStrategy"`
}

type StorageTieredStrategy struct {
	// 低频存储
	// 注意：此字段可能返回 null，表示取不到有效值。
	DayToLowFrequency *int64 `json:"DayToLowFrequency,omitnil,omitempty" name:"DayToLowFrequency"`

	// 归档存储
	// 注意：此字段可能返回 null，表示取不到有效值。
	DayToArchive *int64 `json:"DayToArchive,omitnil,omitempty" name:"DayToArchive"`

	// 深度归档存储
	// 注意：此字段可能返回 null，表示取不到有效值。
	DayToDeepArchive *int64 `json:"DayToDeepArchive,omitnil,omitempty" name:"DayToDeepArchive"`
}

type Tag struct {
	// 标签键。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagFilter struct {
	// 标签键。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值。
	TagValue []*string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagInfo struct {
	// 标签键。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TestItem struct {
	// <p>检测步骤名称</p>
	TestName *string `json:"TestName,omitnil,omitempty" name:"TestName"`

	// <p>错误码</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>错误信息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}