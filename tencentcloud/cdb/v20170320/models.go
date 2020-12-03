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

package v20170320

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Account struct {

	// 新账户的名称
	User *string `json:"User,omitempty" name:"User"`

	// 新账户的域名
	Host *string `json:"Host,omitempty" name:"Host"`
}

type AccountInfo struct {

	// 账号备注信息
	Notes *string `json:"Notes,omitempty" name:"Notes"`

	// 账号的域名
	Host *string `json:"Host,omitempty" name:"Host"`

	// 账号的名称
	User *string `json:"User,omitempty" name:"User"`

	// 账号信息修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 修改密码的时间
	ModifyPasswordTime *string `json:"ModifyPasswordTime,omitempty" name:"ModifyPasswordTime"`

	// 账号的创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type AddTimeWindowRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 星期一的可维护时间段，其中每一个时间段的格式形如：10:00-12:00；起始时间按半个小时对齐；最短半个小时，最长三个小时；可设置多个时间段。 一周中应至少设置一天的时间窗。下同。
	Monday []*string `json:"Monday,omitempty" name:"Monday" list`

	// 星期二的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Tuesday []*string `json:"Tuesday,omitempty" name:"Tuesday" list`

	// 星期三的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Wednesday []*string `json:"Wednesday,omitempty" name:"Wednesday" list`

	// 星期四的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Thursday []*string `json:"Thursday,omitempty" name:"Thursday" list`

	// 星期五的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Friday []*string `json:"Friday,omitempty" name:"Friday" list`

	// 星期六的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Saturday []*string `json:"Saturday,omitempty" name:"Saturday" list`

	// 星期日的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Sunday []*string `json:"Sunday,omitempty" name:"Sunday" list`
}

func (r *AddTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddTimeWindowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddTimeWindowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 安全组 ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例 ID 列表，一个或者多个实例 ID 组成的数组。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AuditFilter struct {

	// 过滤条件参数名称。目前支持：
	// SrcIp – 客户端 IP；
	// User – 数据库账户；
	// DB – 数据库名称；
	Type *string `json:"Type,omitempty" name:"Type"`

	// 过滤条件匹配类型。目前支持：
	// INC – 包含；
	// EXC – 不包含；
	// EQ – 等于；
	// NEQ – 不等于；
	Compare *string `json:"Compare,omitempty" name:"Compare"`

	// 过滤条件匹配值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type AuditLogFile struct {

	// 审计日志文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 审计日志文件创建时间。格式为 : "2019-03-20 17:09:13"。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 文件状态值。可能返回的值为：
	// "creating" - 生成中;
	// "failed" - 创建失败;
	// "success" - 已生成;
	Status *string `json:"Status,omitempty" name:"Status"`

	// 文件大小，单位为 KB。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// 审计日志下载地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type AuditLogFilter struct {

	// 客户端地址。
	Host []*string `json:"Host,omitempty" name:"Host" list`

	// 用户名。
	User []*string `json:"User,omitempty" name:"User" list`

	// 数据库名称。
	DBName []*string `json:"DBName,omitempty" name:"DBName" list`

	// 表名称。
	TableName []*string `json:"TableName,omitempty" name:"TableName" list`

	// 审计策略名称。
	PolicyName []*string `json:"PolicyName,omitempty" name:"PolicyName" list`

	// SQL 语句。支持模糊匹配。
	Sql *string `json:"Sql,omitempty" name:"Sql"`

	// SQL 类型。目前支持："SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "ALTER", "SET", "REPLACE", "EXECUTE"。
	SqlType *string `json:"SqlType,omitempty" name:"SqlType"`

	// 执行时间。单位为：ms。表示筛选执行时间大于该值的审计日志。
	ExecTime *int64 `json:"ExecTime,omitempty" name:"ExecTime"`

	// 影响行数。表示筛选影响行数大于该值的审计日志。
	AffectRows *int64 `json:"AffectRows,omitempty" name:"AffectRows"`
}

type AuditPolicy struct {

	// 审计策略 ID。
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 审计策略的状态。可能返回的值为：
	// "creating" - 创建中;
	// "running" - 运行中;
	// "paused" - 暂停中;
	// "failed" - 创建失败。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 数据库实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 审计策略创建时间。格式为 : "2019-03-20 17:09:13"。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 审计策略最后修改时间。格式为 : "2019-03-20 17:09:13"。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 审计策略名称。
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 审计规则 ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 审计规则名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

type AuditRule struct {

	// 审计规则 Id。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 审计规则创建时间。格式为 : "2019-03-20 17:09:13"。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 审计规则最后修改时间。格式为 : "2019-03-20 17:09:13"。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 审计规则名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 审计规则描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 审计规则过滤条件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleFilters []*AuditFilter `json:"RuleFilters,omitempty" name:"RuleFilters" list`

	// 是否开启全审计。
	AuditAll *bool `json:"AuditAll,omitempty" name:"AuditAll"`
}

type BackupConfig struct {

	// 第二个从库复制方式，可能的返回值：async-异步，semisync-半同步
	ReplicationMode *string `json:"ReplicationMode,omitempty" name:"ReplicationMode"`

	// 第二个从库可用区的正式名称，如ap-shanghai-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 第二个从库内网IP地址
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 第二个从库访问端口
	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
}

type BackupInfo struct {

	// 备份文件名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备份文件大小，单位：Byte
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 备份快照时间，时间格式：2016-03-17 02:10:37
	Date *string `json:"Date,omitempty" name:"Date"`

	// 内网下载地址
	IntranetUrl *string `json:"IntranetUrl,omitempty" name:"IntranetUrl"`

	// 外网下载地址
	InternetUrl *string `json:"InternetUrl,omitempty" name:"InternetUrl"`

	// 日志具体类型。可能的值有 "logical": 逻辑冷备， "physical": 物理冷备。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 备份子任务的ID，删除备份文件时使用
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`

	// 备份任务状态。可能的值有 "SUCCESS": 备份成功， "FAILED": 备份失败， "RUNNING": 备份进行中。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 备份任务的完成时间
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// （该值将废弃，不建议使用）备份的创建者，可能的值：SYSTEM - 系统创建，Uin - 发起者Uin值。
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 备份任务的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 备份方法。可能的值有 "full": 全量备份， "partial": 部分备份。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 备份方式。可能的值有 "manual": 手动备份， "automatic": 自动备份。
	Way *string `json:"Way,omitempty" name:"Way"`
}

type BackupItem struct {

	// 需要备份的库名
	Db *string `json:"Db,omitempty" name:"Db"`

	// 需要备份的表名。 如果传该参数，表示备份该库中的指定表。如果不传该参数则备份该db库
	Table *string `json:"Table,omitempty" name:"Table"`
}

type BackupSummaryItem struct {

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 该实例自动数据备份的个数。
	AutoBackupCount *int64 `json:"AutoBackupCount,omitempty" name:"AutoBackupCount"`

	// 该实例自动数据备份的容量。
	AutoBackupVolume *int64 `json:"AutoBackupVolume,omitempty" name:"AutoBackupVolume"`

	// 该实例手动数据备份的个数。
	ManualBackupCount *int64 `json:"ManualBackupCount,omitempty" name:"ManualBackupCount"`

	// 该实例手动数据备份的容量。
	ManualBackupVolume *int64 `json:"ManualBackupVolume,omitempty" name:"ManualBackupVolume"`

	// 该实例总的数据备份（包含自动备份和手动备份）个数。
	DataBackupCount *int64 `json:"DataBackupCount,omitempty" name:"DataBackupCount"`

	// 该实例总的数据备份容量。
	DataBackupVolume *int64 `json:"DataBackupVolume,omitempty" name:"DataBackupVolume"`

	// 该实例日志备份的个数。
	BinlogBackupCount *int64 `json:"BinlogBackupCount,omitempty" name:"BinlogBackupCount"`

	// 该实例日志备份的容量。
	BinlogBackupVolume *int64 `json:"BinlogBackupVolume,omitempty" name:"BinlogBackupVolume"`

	// 该实例的总备份（包含数据备份和日志备份）占用容量。
	BackupVolume *int64 `json:"BackupVolume,omitempty" name:"BackupVolume"`
}

type BalanceRoGroupLoadRequest struct {
	*tchttp.BaseRequest

	// RO 组的 ID，格式如：cdbrg-c1nl9rpv。
	RoGroupId *string `json:"RoGroupId,omitempty" name:"RoGroupId"`
}

func (r *BalanceRoGroupLoadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BalanceRoGroupLoadRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BalanceRoGroupLoadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BalanceRoGroupLoadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BalanceRoGroupLoadResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BinlogInfo struct {

	// binlog 日志备份文件名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备份文件大小，单位：Byte
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 文件存储时间，时间格式：2016-03-17 02:10:37
	Date *string `json:"Date,omitempty" name:"Date"`

	// 内网下载地址
	IntranetUrl *string `json:"IntranetUrl,omitempty" name:"IntranetUrl"`

	// 外网下载地址
	InternetUrl *string `json:"InternetUrl,omitempty" name:"InternetUrl"`

	// 日志具体类型，可能的值有：binlog - 二进制日志
	Type *string `json:"Type,omitempty" name:"Type"`

	// binlog 文件起始时间
	BinlogStartTime *string `json:"BinlogStartTime,omitempty" name:"BinlogStartTime"`

	// binlog 文件截止时间
	BinlogFinishTime *string `json:"BinlogFinishTime,omitempty" name:"BinlogFinishTime"`
}

type CloneItem struct {

	// 克隆任务的源实例Id。
	SrcInstanceId *string `json:"SrcInstanceId,omitempty" name:"SrcInstanceId"`

	// 克隆任务的新产生实例Id。
	DstInstanceId *string `json:"DstInstanceId,omitempty" name:"DstInstanceId"`

	// 克隆任务对应的任务列表Id。
	CloneJobId *int64 `json:"CloneJobId,omitempty" name:"CloneJobId"`

	// 克隆实例使用的策略， 包括以下类型：  timepoint:指定时间点回档，  backupset: 指定备份文件回档。
	RollbackStrategy *string `json:"RollbackStrategy,omitempty" name:"RollbackStrategy"`

	// 克隆实例回档的时间点。
	RollbackTargetTime *string `json:"RollbackTargetTime,omitempty" name:"RollbackTargetTime"`

	// 任务开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务状态，包括以下状态：initial,running,wait_complete,success,failed
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
}

type CloseWanServiceRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CloseWanServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloseWanServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloseWanServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseWanServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloseWanServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ColumnPrivilege struct {

	// 数据库名
	Database *string `json:"Database,omitempty" name:"Database"`

	// 数据库表名
	Table *string `json:"Table,omitempty" name:"Table"`

	// 数据库列名
	Column *string `json:"Column,omitempty" name:"Column"`

	// 权限信息
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges" list`
}

type CommonTimeWindow struct {

	// 周一的时间窗，格式如： 02:00-06:00
	Monday *string `json:"Monday,omitempty" name:"Monday"`

	// 周二的时间窗，格式如： 02:00-06:00
	Tuesday *string `json:"Tuesday,omitempty" name:"Tuesday"`

	// 周三的时间窗，格式如： 02:00-06:00
	Wednesday *string `json:"Wednesday,omitempty" name:"Wednesday"`

	// 周四的时间窗，格式如： 02:00-06:00
	Thursday *string `json:"Thursday,omitempty" name:"Thursday"`

	// 周五的时间窗，格式如： 02:00-06:00
	Friday *string `json:"Friday,omitempty" name:"Friday"`

	// 周六的时间窗，格式如： 02:00-06:00
	Saturday *string `json:"Saturday,omitempty" name:"Saturday"`

	// 周日的时间窗，格式如： 02:00-06:00
	Sunday *string `json:"Sunday,omitempty" name:"Sunday"`
}

type CreateAccountsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 云数据库账号。
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts" list`

	// 新账户的密码。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 备注信息。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAccountsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAccountsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAuditLogFileRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间，格式为："2017-07-12 10:29:20"。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，格式为："2017-07-12 10:29:20"。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 排序方式。支持值包括："ASC" - 升序，"DESC" - 降序。
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段。支持值包括：
	// "timestamp" - 时间戳；
	// "affectRows" - 影响行数；
	// "execTime" - 执行时间。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 过滤条件。可按设置的过滤条件过滤日志。
	Filter *AuditLogFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *CreateAuditLogFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAuditLogFileRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAuditLogFileResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 审计日志文件名称。
		FileName *string `json:"FileName,omitempty" name:"FileName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAuditLogFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAuditLogFileResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAuditPolicyRequest struct {
	*tchttp.BaseRequest

	// 审计策略名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 审计规则 ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 审计日志保存时长。支持值包括：
	// 30 - 一个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年；
	// 实例首次开通审计策略时，可传该值，用于设置存储日志保存天数，默认为 30 天。若实例已存在审计策略，则此参数无效，可使用 更改审计服务配置 接口修改日志存储时长。
	LogExpireDay *int64 `json:"LogExpireDay,omitempty" name:"LogExpireDay"`
}

func (r *CreateAuditPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAuditPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAuditPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 审计策略 ID。
		PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAuditPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAuditPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAuditRuleRequest struct {
	*tchttp.BaseRequest

	// 审计规则名称。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 审计规则描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 审计规则过滤条件。若设置了过滤条件，则不会开启全审计。
	RuleFilters []*AuditFilter `json:"RuleFilters,omitempty" name:"RuleFilters" list`

	// 是否开启全审计。支持值包括：false – 不开启全审计，true – 开启全审计。用户未设置审计规则过滤条件时，默认开启全审计。
	AuditAll *bool `json:"AuditAll,omitempty" name:"AuditAll"`
}

func (r *CreateAuditRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAuditRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAuditRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 审计规则 ID。
		RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAuditRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAuditRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 目标备份方法，可选的值：logical - 逻辑冷备，physical - 物理冷备。
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// 需要备份的库表信息，如果不设置该参数，则默认整实例备份。在 BackupMethod=logical 逻辑备份中才可设置该参数。指定的库表必须存在，否则可能导致备份失败。
	// 例：如果需要备份 db1 库的 tb1、tb2 表 和 db2 库。则该参数设置为 [{"Db": "db1", "Table": "tb1"}, {"Db": "db1", "Table": "tb2"}, {"Db": "db2"} ]。
	BackupDBTableList []*BackupItem `json:"BackupDBTableList,omitempty" name:"BackupDBTableList" list`
}

func (r *CreateBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBackupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBackupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 备份任务 ID。
		BackupId *uint64 `json:"BackupId,omitempty" name:"BackupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBackupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCloneInstanceRequest struct {
	*tchttp.BaseRequest

	// 克隆源实例Id。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 如果需要克隆实例回档到指定时间，则指定该值。时间格式为： yyyy-mm-dd hh:mm:ss 。
	SpecifiedRollbackTime *string `json:"SpecifiedRollbackTime,omitempty" name:"SpecifiedRollbackTime"`

	// 如果需要克隆实例回档到指定备份的时间点，则指定该值为物理备份的Id。请使用 [查询数据备份文件列表](/document/api/236/15842) 。
	SpecifiedBackupId *int64 `json:"SpecifiedBackupId,omitempty" name:"SpecifiedBackupId"`

	// 私有网络 ID，如果不传则默认选择基础网络，请使用 [查询私有网络列表](/document/api/215/15778) 。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 私有网络下的子网 ID，如果设置了 UniqVpcId，则 UniqSubnetId 必填，请使用 [查询子网列表](/document/api/215/15784)。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 实例内存大小，单位：MB，需要不低于克隆源实例，默认和源实例相同。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB，需要不低于克隆源实例，默认和源实例相同。
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// 新产生的克隆实例名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 安全组参数，可使用 [查询项目安全组信息](https://cloud.tencent.com/document/api/236/15850) 接口查询某个项目的安全组详情。
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup" list`

	// 实例标签信息。
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags" list`

	// 实例Cpu核数，需要不低于克隆源实例。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制。
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// 多可用区域，默认为 0，支持值包括：0 - 表示单可用区，1 - 表示多可用区。
	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// 新产生的克隆实例备库 1 的可用区信息，默认同源实例 Zone 的值。
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// 备库 2 的可用区信息，默认为空，克隆强同步主实例时可指定该参数。
	BackupZone *string `json:"BackupZone,omitempty" name:"BackupZone"`

	// 克隆实例类型。支持值包括： "HA" - 高可用版实例， "EXCLUSIVE" - 独享型实例。 不指定则默认为高可用版。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
}

func (r *CreateCloneInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCloneInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCloneInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的请求ID，可使用此 ID 查询异步任务的执行结果。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCloneInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCloneInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBImportJobRequest struct {
	*tchttp.BaseRequest

	// 实例的 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 文件名称。该文件是指用户已上传到腾讯云的文件。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 云数据库的用户名。
	User *string `json:"User,omitempty" name:"User"`

	// 云数据库实例 User 账号的密码。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 导入的目标数据库名，不传表示不指定数据库。
	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

func (r *CreateDBImportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBImportJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBImportJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBImportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBImportJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceHourRequest struct {
	*tchttp.BaseRequest

	// 实例数量，默认值为 1，最小值 1，最大值为 100。
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 实例内存大小，单位：MB，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的内存规格。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的硬盘范围。
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// MySQL 版本，值包括：5.5、5.6 和 5.7，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的实例版本。
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 私有网络 ID，如果不传则默认选择基础网络，请使用 [查询私有网络列表](/document/api/215/15778) 。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 私有网络下的子网 ID，如果设置了 UniqVpcId，则 UniqSubnetId 必填，请使用[查询子网列表](/document/api/215/15784)。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 项目 ID，不填为默认项目。请使用 [查询项目列表](https://cloud.tencent.com/document/product/378/4400) 接口获取项目 ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 可用区信息，该参数缺省时，系统会自动选择一个可用区，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例 ID，购买只读实例或者灾备实例时必填，该字段表示只读实例或者灾备实例的主实例 ID，请使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口查询云数据库实例 ID。
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" name:"MasterInstanceId"`

	// 实例类型，默认为 master，支持值包括：master - 表示主实例，dr - 表示灾备实例，ro - 表示只读实例。
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// 主实例的可用区信息，购买灾备实例时必填。
	MasterRegion *string `json:"MasterRegion,omitempty" name:"MasterRegion"`

	// 自定义端口，端口支持范围：[ 1024-65535 ] 。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 设置 root 帐号密码，密码规则：8 - 64 个字符，至少包含字母、数字、字符（支持的字符：_+-&=!@#$%^*()）中的两种，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 参数列表，参数格式如 ParamList.0.Name=auto_increment&ParamList.0.Value=1。可通过 [查询默认的可设置参数列表](https://cloud.tencent.com/document/api/236/32662) 查询支持设置的参数。
	ParamList []*ParamInfo `json:"ParamList,omitempty" name:"ParamList" list`

	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// 多可用区域，默认为 0，支持值包括：0 - 表示单可用区，1 - 表示多可用区，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// 备库 1 的可用区信息，默认为 Zone 的值，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// 备库 2 的可用区信息，默认为空，购买强同步主实例时可指定该参数，购买其他类型实例时指定该参数无意义。
	BackupZone *string `json:"BackupZone,omitempty" name:"BackupZone"`

	// 安全组参数，可使用 [查询项目安全组信息](https://cloud.tencent.com/document/api/236/15850) 接口查询某个项目的安全组详情。
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup" list`

	// 只读实例信息。购买只读实例时，该参数必传。
	RoGroup *RoGroup `json:"RoGroup,omitempty" name:"RoGroup"`

	// 购买按量计费实例该字段无意义。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例标签信息。
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags" list`

	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间在当天内唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 实例类型。支持值包括： "HA" - 高可用版实例， "BASIC" - 基础版实例。 不指定则默认为高可用版。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 参数模板id。
	ParamTemplateId *int64 `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`

	// 告警策略id数组。
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitempty" name:"AlarmPolicyList" list`
}

func (r *CreateDBInstanceHourRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstanceHourRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceHourResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 短订单 ID。
		DealIds []*string `json:"DealIds,omitempty" name:"DealIds" list`

		// 实例 ID 列表。
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBInstanceHourResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstanceHourResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例内存大小，单位：MB，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的内存规格。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的硬盘范围。
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// 实例时长，单位：月，可选值包括 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 实例数量，默认值为1, 最小值1，最大值为100。
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 可用区信息，该参数缺省时，系统会自动选择一个可用区，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 私有网络 ID，如果不传则默认选择基础网络，请使用 [查询私有网络列表](/document/api/215/15778) 。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 私有网络下的子网 ID，如果设置了 UniqVpcId，则 UniqSubnetId 必填，请使用 [查询子网列表](/document/api/215/15784)。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 项目 ID，不填为默认项目。请使用 [查询项目列表](https://cloud.tencent.com/document/product/378/4400) 接口获取项目 ID。购买只读实例和灾备实例时，项目 ID 默认和主实例保持一致。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 自定义端口，端口支持范围：[ 1024-65535 ]。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 实例类型，默认为 master，支持值包括：master - 表示主实例，dr - 表示灾备实例，ro - 表示只读实例。
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// 实例 ID，购买只读实例时必填，该字段表示只读实例的主实例ID，请使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口查询云数据库实例 ID。
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" name:"MasterInstanceId"`

	// MySQL 版本，值包括：5.5、5.6 和 5.7，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的实例版本。
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 设置 root 帐号密码，密码规则：8 - 64 个字符，至少包含字母、数字、字符（支持的字符：_+-&=!@#$%^*()）中的两种，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制。
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// 多可用区域，默认为 0，支持值包括：0 - 表示单可用区，1 - 表示多可用区。
	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// 备库 1 的可用区信息，默认为 Zone 的值。
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// 参数列表，参数格式如 ParamList.0.Name=auto_increment&ParamList.0.Value=1。可通过 [查询默认的可设置参数列表](https://cloud.tencent.com/document/api/236/32662) 查询支持设置的参数。
	ParamList []*ParamInfo `json:"ParamList,omitempty" name:"ParamList" list`

	// 备库 2 的可用区信息，默认为空，购买强同步主实例时可指定该参数，购买其他类型实例时指定该参数无意义。
	BackupZone *string `json:"BackupZone,omitempty" name:"BackupZone"`

	// 自动续费标记，可选值为：0 - 不自动续费；1 - 自动续费。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 主实例地域信息，购买灾备实例时，该字段必填。
	MasterRegion *string `json:"MasterRegion,omitempty" name:"MasterRegion"`

	// 安全组参数，可使用 [查询项目安全组信息](https://cloud.tencent.com/document/api/236/15850) 接口查询某个项目的安全组详情。
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup" list`

	// 只读实例参数。购买只读实例时，该参数必传。
	RoGroup *RoGroup `json:"RoGroup,omitempty" name:"RoGroup"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例标签信息。
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags" list`

	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间在当天内唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 实例类型。支持值包括： "HA" - 高可用版实例， "BASIC" - 基础版实例。 不指定则默认为高可用版。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 参数模板id。
	ParamTemplateId *int64 `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`

	// 告警策略id数组。
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitempty" name:"AlarmPolicyList" list`
}

func (r *CreateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 短订单 ID。
		DealIds []*string `json:"DealIds,omitempty" name:"DealIds" list`

		// 实例 ID 列表。
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDeployGroupRequest struct {
	*tchttp.BaseRequest

	// 置放群组名称，最长不能超过60个字符。
	DeployGroupName *string `json:"DeployGroupName,omitempty" name:"DeployGroupName"`

	// 置放群组描述，最长不能超过200个字符。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 置放群组的亲和性策略，目前仅支持取值为1，策略1表示同台物理机上限制实例的个数。
	Affinity []*int64 `json:"Affinity,omitempty" name:"Affinity" list`

	// 置放群组亲和性策略1中同台物理机上实例的限制个数。
	LimitNum *int64 `json:"LimitNum,omitempty" name:"LimitNum"`

	// 置放群组机型属性，可选参数：SH12+SH02、TS85。
	DevClass []*string `json:"DevClass,omitempty" name:"DevClass" list`
}

func (r *CreateDeployGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDeployGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDeployGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 置放群组ID。
		DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDeployGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDeployGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateParamTemplateRequest struct {
	*tchttp.BaseRequest

	// 参数模板名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数模板描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// MySQL 版本号。
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 源参数模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 参数列表。
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList" list`
}

func (r *CreateParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateParamTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 参数模板 ID。
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateParamTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRoInstanceIpRequest struct {
	*tchttp.BaseRequest

	// 只读实例ID，格式如：cdbro-3i70uj0k，与云数据库控制台页面中显示的只读实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 子网描述符，例如：subnet-1typ0s7d。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// vpc描述符，例如：vpc-xxx,如果传了该字段则UniqSubnetId必传
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
}

func (r *CreateRoInstanceIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRoInstanceIpRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRoInstanceIpResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 只读实例的私有网络的ID。
		RoVpcId *int64 `json:"RoVpcId,omitempty" name:"RoVpcId"`

		// 只读实例的子网ID。
		RoSubnetId *int64 `json:"RoSubnetId,omitempty" name:"RoSubnetId"`

		// 只读实例的内网IP地址。
		RoVip *string `json:"RoVip,omitempty" name:"RoVip"`

		// 只读实例的内网端口号。
		RoVport *int64 `json:"RoVport,omitempty" name:"RoVport"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRoInstanceIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRoInstanceIpResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DBSwitchInfo struct {

	// 切换时间，格式为：2017-09-03 01:34:31
	SwitchTime *string `json:"SwitchTime,omitempty" name:"SwitchTime"`

	// 切换类型，可能的返回值为：TRANSFER - 数据迁移；MASTER2SLAVE - 主备切换；RECOVERY - 主从恢复
	SwitchType *string `json:"SwitchType,omitempty" name:"SwitchType"`
}

type DatabaseName struct {

	// 数据库表名
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
}

type DatabasePrivilege struct {

	// 权限信息
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges" list`

	// 数据库名
	Database *string `json:"Database,omitempty" name:"Database"`
}

type DeleteAccountsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 云数据库账号。
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts" list`
}

func (r *DeleteAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAccountsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAccountsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAuditLogFileRequest struct {
	*tchttp.BaseRequest

	// 审计日志文件名称。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteAuditLogFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAuditLogFileRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAuditLogFileResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAuditLogFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAuditLogFileResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAuditPolicyRequest struct {
	*tchttp.BaseRequest

	// 审计策略 ID。
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteAuditPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAuditPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAuditPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAuditPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAuditPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAuditRuleRequest struct {
	*tchttp.BaseRequest

	// 审计规则 ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteAuditRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAuditRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAuditRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAuditRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAuditRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBackupRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份任务 ID。该任务 ID 为 [创建云数据库备份](https://cloud.tencent.com/document/api/236/15844) 接口返回的任务 ID。
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`
}

func (r *DeleteBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBackupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBackupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBackupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDeployGroupsRequest struct {
	*tchttp.BaseRequest

	// 要删除的置放群组 ID 列表。
	DeployGroupIds []*string `json:"DeployGroupIds,omitempty" name:"DeployGroupIds" list`
}

func (r *DeleteDeployGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDeployGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDeployGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDeployGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDeployGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteParamTemplateRequest struct {
	*tchttp.BaseRequest

	// 参数模板ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteParamTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteParamTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTimeWindowRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTimeWindowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTimeWindowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployGroupInfo struct {

	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// 置放群组名称。
	DeployGroupName *string `json:"DeployGroupName,omitempty" name:"DeployGroupName"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 置放群组实例配额，表示一个置放群组中可容纳的最大实例数目。
	Quota *int64 `json:"Quota,omitempty" name:"Quota"`

	// 置放群组亲和性策略，目前仅支持策略1，即在物理机纬度打散实例的分布。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Affinity *string `json:"Affinity,omitempty" name:"Affinity"`

	// 置放群组亲和性策略1中，同台物理机上同个置放群组实例的限制个数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LimitNum *int64 `json:"LimitNum,omitempty" name:"LimitNum"`

	// 置放群组详细信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 置放群组物理机型属性。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevClass *string `json:"DevClass,omitempty" name:"DevClass"`
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库的账号名称。
	User *string `json:"User,omitempty" name:"User"`

	// 数据库的账号域名。
	Host *string `json:"Host,omitempty" name:"Host"`
}

func (r *DescribeAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountPrivilegesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 全局权限数组。
		GlobalPrivileges []*string `json:"GlobalPrivileges,omitempty" name:"GlobalPrivileges" list`

		// 数据库权限数组。
		DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitempty" name:"DatabasePrivileges" list`

		// 数据库中的表权限数组。
		TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitempty" name:"TablePrivileges" list`

		// 数据库表中的列权限数组。
		ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitempty" name:"ColumnPrivileges" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountPrivilegesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 记录偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为20，最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 匹配账号名的正则表达式，规则同 MySQL 官网。
	AccountRegexp *string `json:"AccountRegexp,omitempty" name:"AccountRegexp"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件的账号数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 符合查询条件的账号详细信息。
		Items []*AccountInfo `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncRequestInfoRequest struct {
	*tchttp.BaseRequest

	// 异步任务的请求 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

func (r *DescribeAsyncRequestInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAsyncRequestInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncRequestInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务执行结果。可能的取值：INITIAL - 初始化，RUNNING - 运行中，SUCCESS - 执行成功，FAILED - 执行失败，KILLED - 已终止，REMOVED - 已删除，PAUSED - 终止中。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *string `json:"Status,omitempty" name:"Status"`

		// 任务执行信息描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Info *string `json:"Info,omitempty" name:"Info"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAsyncRequestInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAsyncRequestInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditConfigRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeAuditConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuditConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 审计日志保存时长。目前支持的值包括：[0，30，180，365，1095，1825]。
	// 注意：此字段可能返回 null，表示取不到有效值。
		LogExpireDay *int64 `json:"LogExpireDay,omitempty" name:"LogExpireDay"`

		// 审计日志存储类型。目前支持的值包括："storage" - 存储型。
		LogType *string `json:"LogType,omitempty" name:"LogType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuditConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuditConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditLogFilesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页大小参数。默认值为 20，最小值为 1，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 审计日志文件名。
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

func (r *DescribeAuditLogFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuditLogFilesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditLogFilesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的审计日志文件个数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 审计日志文件详情。
		Items []*AuditLogFile `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuditLogFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuditLogFilesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditPoliciesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 审计策略 ID。
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 审计策略名称。支持按审计策略名称进行模糊匹配查询。
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 分页大小参数。默认值为 20，最小值为 1，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 审计规则 ID。可使用该审计规则 ID 查询到其关联的审计策略。
	// 注意，参数 RuleId，InstanceId，PolicyId，PolicyName 必须至少传一个。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DescribeAuditPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuditPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的审计策略个数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 审计策略详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Items []*AuditPolicy `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuditPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuditPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditRulesRequest struct {
	*tchttp.BaseRequest

	// 审计规则 ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 审计规则名称。支持按审计规则名称进行模糊匹配查询。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 分页大小参数。默认值为 20，最小值为 1，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAuditRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuditRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的审计规则个数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 审计规则详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Items []*AuditRule `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuditRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuditRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupConfigRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 自动备份开始的最早时间点，单位为时刻。例如，2 - 凌晨 2:00。（该字段已废弃，建议使用 BackupTimeWindow 字段）
		StartTimeMin *int64 `json:"StartTimeMin,omitempty" name:"StartTimeMin"`

		// 自动备份开始的最晚时间点，单位为时刻。例如，6 - 凌晨 6:00。（该字段已废弃，建议使用 BackupTimeWindow 字段）
		StartTimeMax *int64 `json:"StartTimeMax,omitempty" name:"StartTimeMax"`

		// 备份文件保留时间，单位为天。
		BackupExpireDays *int64 `json:"BackupExpireDays,omitempty" name:"BackupExpireDays"`

		// 备份方式，可能的值为：physical - 物理备份，logical - 逻辑备份。
		BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

		// Binlog 文件保留时间，单位为天。
		BinlogExpireDays *int64 `json:"BinlogExpireDays,omitempty" name:"BinlogExpireDays"`

		// 实例自动备份的时间窗。
		BackupTimeWindow *CommonTimeWindow `json:"BackupTimeWindow,omitempty" name:"BackupTimeWindow"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupDatabasesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间，格式为：2017-07-12 10:29:20。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 要查询的数据库名前缀。
	SearchDatabase *string `json:"SearchDatabase,omitempty" name:"SearchDatabase"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，最小值为1，最大值为2000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBackupDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupDatabasesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的数据个数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 符合查询条件的数据库数组。
		Items []*DatabaseName `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupDatabasesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupOverviewRequest struct {
	*tchttp.BaseRequest

	// 需要查询的云数据库产品类型，目前仅支持 "mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeBackupOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户在当前地域备份的总个数（包含数据备份和日志备份）。
		BackupCount *int64 `json:"BackupCount,omitempty" name:"BackupCount"`

		// 用户在当前地域备份的总容量
		BackupVolume *int64 `json:"BackupVolume,omitempty" name:"BackupVolume"`

		// 用户在当前地域备份的计费容量，即超出赠送容量的部分。
		BillingVolume *int64 `json:"BillingVolume,omitempty" name:"BillingVolume"`

		// 用户在当前地域获得的赠送备份容量。
		FreeVolume *int64 `json:"FreeVolume,omitempty" name:"FreeVolume"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupSummariesRequest struct {
	*tchttp.BaseRequest

	// 需要查询的云数据库产品类型，目前仅支持 "mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 分页查询数据的偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询数据的条目限制，默认值为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 指定按某一项排序，可选值包括： BackupVolume: 备份容量， DataBackupVolume: 数据备份容量， BinlogBackupVolume: 日志备份容量， AutoBackupVolume: 自动备份容量， ManualBackupVolume: 手动备份容量。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 指定排序方向，可选值包括： ASC: 正序， DESC: 逆序。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeBackupSummariesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupSummariesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupSummariesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例备份统计条目。
		Items []*BackupSummaryItem `json:"Items,omitempty" name:"Items" list`

		// 实例备份统计总条目数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupSummariesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupSummariesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupTablesRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间，格式为：2017-07-12 10:29:20。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 指定的数据库名。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 要查询的数据表名前缀。
	SearchTable *string `json:"SearchTable,omitempty" name:"SearchTable"`

	// 分页偏移。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，最小值为1，最大值为2000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBackupTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的数据个数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 符合条件的数据表数组。
		Items []*TableName `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量，最小值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认值为20，最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件的实例总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 符合查询条件的备份信息详情。
		Items []*BackupInfo `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBinlogBackupOverviewRequest struct {
	*tchttp.BaseRequest

	// 需要查询的云数据库产品类型，目前仅支持 "mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeBinlogBackupOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBinlogBackupOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBinlogBackupOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总的日志备份容量（单位为字节）。
		BinlogBackupVolume *int64 `json:"BinlogBackupVolume,omitempty" name:"BinlogBackupVolume"`

		// 总的日志备份个数。
		BinlogBackupCount *int64 `json:"BinlogBackupCount,omitempty" name:"BinlogBackupCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBinlogBackupOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBinlogBackupOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBinlogsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量，最小值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认值为20，最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBinlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBinlogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBinlogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件的日志文件总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 符合查询条件的二进制日志文件详情。
		Items []*BinlogInfo `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBinlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBinlogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloneListRequest struct {
	*tchttp.BaseRequest

	// 查询指定源实例的克隆任务列表。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页查询时的偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询时的每页条目数。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCloneListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloneListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloneListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 满足条件的条目数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 克隆任务列表。
		Items []*CloneItem `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloneListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloneListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBImportRecordsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间，时间格式如：2016-01-01 00:00:01。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，时间格式如：2016-01-01 23:59:59。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页参数，偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页参数，单次请求返回的数量，默认值为20，最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDBImportRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBImportRecordsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBImportRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件的导入任务操作日志总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回的导入操作记录列表。
		Items []*ImportRecord `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBImportRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBImportRecordsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceCharsetRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceCharsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceCharsetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceCharsetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例的默认字符集，如 "latin1"，"utf8" 等。
		Charset *string `json:"Charset,omitempty" name:"Charset"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceCharsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceCharsetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceConfigRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 主实例数据保护方式，可能的返回值：0 - 异步复制方式，1 - 半同步复制方式，2 - 强同步复制方式。
		ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

		// 主实例部署方式，可能的返回值：0 - 单可用部署，1 - 多可用区部署。
		DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

		// 实例可用区信息，格式如 "ap-shanghai-1"。
		Zone *string `json:"Zone,omitempty" name:"Zone"`

		// 备库的配置信息。
		SlaveConfig *SlaveConfig `json:"SlaveConfig,omitempty" name:"SlaveConfig"`

		// 强同步实例第二备库的配置信息。
		BackupConfig *BackupConfig `json:"BackupConfig,omitempty" name:"BackupConfig"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceGTIDRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceGTIDRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceGTIDRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceGTIDResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// GTID 是否开通的标记，可能的取值为：0 - 未开通，1 - 已开通。
		IsGTIDOpen *int64 `json:"IsGTIDOpen,omitempty" name:"IsGTIDOpen"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceGTIDResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceGTIDResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceInfoRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例 ID 。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 实例名称。
		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

		// 是否开通加密，YES 已开通，NO 未开通。
		Encryption *string `json:"Encryption,omitempty" name:"Encryption"`

		// 加密使用的密钥 ID 。
	// 注意：此字段可能返回 null，表示取不到有效值。
		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

		// 密钥所在地域。
	// 注意：此字段可能返回 null，表示取不到有效值。
		KeyRegion *string `json:"KeyRegion,omitempty" name:"KeyRegion"`

		// 当前 CDB 后端服务使用的 KMS 服务的默认地域。
	// 注意：此字段可能返回 null，表示取不到有效值。
		DefaultKmsRegion *string `json:"DefaultKmsRegion,omitempty" name:"DefaultKmsRegion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceRebootTimeRequest struct {
	*tchttp.BaseRequest

	// 实例的 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DescribeDBInstanceRebootTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceRebootTimeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceRebootTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件的实例总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回的参数信息。
		Items []*InstanceRebootTime `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceRebootTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceRebootTimeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 项目 ID，可使用 [查询项目列表](https://cloud.tencent.com/document/product/378/4400) 接口查询项目 ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例类型，可取值：1 - 主实例，2 - 灾备实例，3 - 只读实例。
	InstanceTypes []*uint64 `json:"InstanceTypes,omitempty" name:"InstanceTypes" list`

	// 实例的内网 IP 地址。
	Vips []*string `json:"Vips,omitempty" name:"Vips" list`

	// 实例状态，可取值：<br>0 - 创建中<br>1 - 运行中<br>4 - 正在进行隔离操作<br>5 - 隔离中（可在回收站恢复开机）
	Status []*uint64 `json:"Status,omitempty" name:"Status" list`

	// 偏移量，默认值为 0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为 20，最大值为 2000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 安全组 ID。当使用安全组 ID 为过滤条件时，需指定 WithSecurityGroup 参数为 1。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 付费类型，可取值：0 - 包年包月，1 - 小时计费。
	PayTypes []*uint64 `json:"PayTypes,omitempty" name:"PayTypes" list`

	// 实例名称。
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames" list`

	// 实例任务状态，可能取值：<br>0 - 没有任务<br>1 - 升级中<br>2 - 数据导入中<br>3 - 开放Slave中<br>4 - 外网访问开通中<br>5 - 批量操作执行中<br>6 - 回档中<br>7 - 外网访问关闭中<br>8 - 密码修改中<br>9 - 实例名修改中<br>10 - 重启中<br>12 - 自建迁移中<br>13 - 删除库表中<br>14 - 灾备实例创建同步中<br>15 - 升级待切换<br>16 - 升级切换中<br>17 - 升级切换完成<br>19 - 参数设置待执行
	TaskStatus []*uint64 `json:"TaskStatus,omitempty" name:"TaskStatus" list`

	// 实例数据库引擎版本，可能取值：5.1、5.5、5.6 和 5.7。
	EngineVersions []*string `json:"EngineVersions,omitempty" name:"EngineVersions" list`

	// 私有网络的 ID。
	VpcIds []*uint64 `json:"VpcIds,omitempty" name:"VpcIds" list`

	// 可用区的 ID。
	ZoneIds []*uint64 `json:"ZoneIds,omitempty" name:"ZoneIds" list`

	// 子网 ID。
	SubnetIds []*uint64 `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// 是否锁定标记。
	CdbErrors []*int64 `json:"CdbErrors,omitempty" name:"CdbErrors" list`

	// 返回结果集排序的字段，目前支持："InstanceId"，"InstanceName"，"CreateTime"，"DeadlineTime"。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 返回结果集排序方式，目前支持："ASC" 或者 "DESC"。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 是否以安全组 ID 为过滤条件。
	WithSecurityGroup *int64 `json:"WithSecurityGroup,omitempty" name:"WithSecurityGroup"`

	// 是否包含独享集群详细信息，可取值：0 - 不包含，1 - 包含。
	WithExCluster *int64 `json:"WithExCluster,omitempty" name:"WithExCluster"`

	// 独享集群 ID。
	ExClusterId *string `json:"ExClusterId,omitempty" name:"ExClusterId"`

	// 实例 ID。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 初始化标记，可取值：0 - 未初始化，1 - 初始化。
	InitFlag *int64 `json:"InitFlag,omitempty" name:"InitFlag"`

	// 是否包含灾备关系对应的实例，可取值：0 - 不包含，1 - 包含。默认取值为1。如果拉取主实例，则灾备关系的数据在DrInfo字段中， 如果拉取灾备实例， 则灾备关系的数据在MasterInfo字段中。灾备关系中只包含部分基本的数据，详细的数据需要自行调接口拉取。
	WithDr *int64 `json:"WithDr,omitempty" name:"WithDr"`

	// 是否包含只读实例，可取值：0 - 不包含，1 - 包含。默认取值为1。
	WithRo *int64 `json:"WithRo,omitempty" name:"WithRo"`

	// 是否包含主实例，可取值：0 - 不包含，1 - 包含。默认取值为1。
	WithMaster *int64 `json:"WithMaster,omitempty" name:"WithMaster"`

	// 置放群组ID列表。
	DeployGroupIds []*string `json:"DeployGroupIds,omitempty" name:"DeployGroupIds" list`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件的实例总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例详细信息。
		Items []*InstanceInfo `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBPriceRequest struct {
	*tchttp.BaseRequest

	// 可用区信息，格式如 "ap-guangzhou-2"。具体能设置的值请通过 <a href="https://cloud.tencent.com/document/api/236/17229">DescribeDBZoneConfig</a> 接口查询。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例数量，默认值为 1，最小值 1，最大值为 100。
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 实例内存大小，单位：MB。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB。
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// 付费类型，支持值包括：PRE_PAID - 包年包月，HOUR_PAID - 按量计费。
	PayType *string `json:"PayType,omitempty" name:"PayType"`

	// 实例时长，单位：月，最小值 1，最大值为 36；查询按量计费价格时，该字段无效。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 实例类型，默认为 master，支持值包括：master - 表示主实例，ro - 表示只读实例，dr - 表示灾备实例。
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制。
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// 部署策略，取值范围：HA-高可用版两节点，FE-金融版三节点；默认值根据数据复制方式决定，数据复制方式为强同步复制时默认值为FE，数据复制方式非强同步时默认值为HA。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
}

func (r *DescribeDBPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBPriceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBPriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例价格，单位：分（人民币）。
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// 实例原价，单位：分（人民币）。
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBPriceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全组详情。
		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSwitchRecordsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认值为 50，最小值为 1，最大值为 2000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDBSwitchRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSwitchRecordsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSwitchRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例切换记录的总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例切换记录详情。
		Items []*DBSwitchInfo `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSwitchRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSwitchRecordsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBZoneConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDBZoneConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBZoneConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBZoneConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可售卖地域配置数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 可售卖地域配置详情
		Items []*RegionSellConf `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBZoneConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBZoneConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDataBackupOverviewRequest struct {
	*tchttp.BaseRequest

	// 需要查询的云数据库产品类型，目前仅支持 "mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeDataBackupOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDataBackupOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDataBackupOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 当前地域的数据备份总容量（包含自动备份和手动备份，单位为字节）。
		DataBackupVolume *int64 `json:"DataBackupVolume,omitempty" name:"DataBackupVolume"`

		// 当前地域的数据备份总个数。
		DataBackupCount *int64 `json:"DataBackupCount,omitempty" name:"DataBackupCount"`

		// 当前地域的自动备份总容量。
		AutoBackupVolume *int64 `json:"AutoBackupVolume,omitempty" name:"AutoBackupVolume"`

		// 当前地域的自动备份总个数。
		AutoBackupCount *int64 `json:"AutoBackupCount,omitempty" name:"AutoBackupCount"`

		// 当前地域的手动备份总容量。
		ManualBackupVolume *int64 `json:"ManualBackupVolume,omitempty" name:"ManualBackupVolume"`

		// 当前地域的手动备份总个数。
		ManualBackupCount *int64 `json:"ManualBackupCount,omitempty" name:"ManualBackupCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataBackupOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDataBackupOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量，最小值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单次请求数量，默认值为20，最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 匹配数据库库名的正则表达式。
	DatabaseRegexp *string `json:"DatabaseRegexp,omitempty" name:"DatabaseRegexp"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件的实例总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回的实例信息。
		Items []*string `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDefaultParamsRequest struct {
	*tchttp.BaseRequest

	// mysql版本，目前支持 ["5.1", "5.5", "5.6", "5.7"]。
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
}

func (r *DescribeDefaultParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDefaultParamsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDefaultParamsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 参数个数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 参数详情。
		Items []*ParameterDetail `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDefaultParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDefaultParamsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeployGroupListRequest struct {
	*tchttp.BaseRequest

	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// 置放群组名称。
	DeployGroupName *string `json:"DeployGroupName,omitempty" name:"DeployGroupName"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDeployGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeployGroupListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeployGroupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的记录总数。
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 返回列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Items []*DeployGroupInfo `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeployGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeployGroupListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceMonitorInfoRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 返回当天最近Count个5分钟粒度的监控数据。最小值1，最大值288，不传该参数默认返回当天所有5分钟粒度监控数据。
	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

func (r *DescribeDeviceMonitorInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceMonitorInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceMonitorInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例CPU监控数据
		Cpu *DeviceCpuInfo `json:"Cpu,omitempty" name:"Cpu"`

		// 实例内存监控数据
		Mem *DeviceMemInfo `json:"Mem,omitempty" name:"Mem"`

		// 实例网络监控数据
		Net *DeviceNetInfo `json:"Net,omitempty" name:"Net"`

		// 实例磁盘监控数据
		Disk *DeviceDiskInfo `json:"Disk,omitempty" name:"Disk"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceMonitorInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceMonitorInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeErrorLogDataRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间戳。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 要匹配的关键字列表，最多支持15个关键字。
	KeyWords []*string `json:"KeyWords,omitempty" name:"KeyWords" list`

	// 分页的返回数量，默认为100，最大为400。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeErrorLogDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeErrorLogDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeErrorLogDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的记录总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回的记录。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Items []*ErrlogItem `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeErrorLogDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeErrorLogDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceParamRecordsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstanceParamRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceParamRecordsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceParamRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的记录数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 参数修改记录。
		Items []*ParamRecord `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceParamRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceParamRecordsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceParamsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例的参数总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 参数详情。
		Items []*ParameterDetail `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceParamsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeParamTemplateInfoRequest struct {
	*tchttp.BaseRequest

	// 参数模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeParamTemplateInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeParamTemplateInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeParamTemplateInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 参数模板 ID。
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// 参数模板名称。
		Name *string `json:"Name,omitempty" name:"Name"`

		// 参数模板描述
		EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

		// 参数模板中的参数数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 参数详情
		Items []*ParameterDetail `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeParamTemplateInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeParamTemplateInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeParamTemplatesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeParamTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeParamTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeParamTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 该用户的参数模板数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 参数模板详情。
		Items []*ParamTemplateInfo `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeParamTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeParamTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeProjectSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全组详情。
		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoGroupsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cdb-c1nl9rpv或者cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeRoGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// RO组信息数组，一个实例可关联多个RO组。
		RoGroups []*RoGroup `json:"RoGroups,omitempty" name:"RoGroups" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoMinScaleRequest struct {
	*tchttp.BaseRequest

	// 只读实例ID，格式如：cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，该参数与MasterInstanceId参数不能同时为空。
	RoInstanceId *string `json:"RoInstanceId,omitempty" name:"RoInstanceId"`

	// 主实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，该参数与RoInstanceId参数不能同时为空。注意，当传入参数包含RoInstanceId时，返回值为只读实例升级时的最小规格；当传入参数只包含MasterInstanceId时，返回值为只读实例购买时的最小规格。
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" name:"MasterInstanceId"`
}

func (r *DescribeRoMinScaleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoMinScaleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoMinScaleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 内存规格大小, 单位为：MB。
		Memory *int64 `json:"Memory,omitempty" name:"Memory"`

		// 磁盘规格大小, 单位为：GB。
		Volume *int64 `json:"Volume,omitempty" name:"Volume"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoMinScaleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoMinScaleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRollbackRangeTimeRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 列表，单个实例 ID 的格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DescribeRollbackRangeTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRollbackRangeTimeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRollbackRangeTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件的实例总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回的参数信息。
		Items []*InstanceRollbackRangeTime `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRollbackRangeTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRollbackRangeTimeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRollbackTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 实例 ID。与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872)。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 异步任务 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 分页参数，每次请求返回的记录数。默认值为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量。默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeRollbackTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRollbackTaskDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRollbackTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的记录总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 回档任务详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Items []*RollbackTask `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRollbackTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRollbackTaskDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogDataRequest struct {
	*tchttp.BaseRequest

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间戳。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 客户端 Host 列表。
	UserHosts []*string `json:"UserHosts,omitempty" name:"UserHosts" list`

	// 客户端 用户名 列表。
	UserNames []*string `json:"UserNames,omitempty" name:"UserNames" list`

	// 访问的 数据库 列表。
	DataBases []*string `json:"DataBases,omitempty" name:"DataBases" list`

	// 排序字段。当前支持：Timestamp,QueryTime,LockTime,RowsExamined,RowsSent 。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 升序还是降序排列。当前支持：ASC,DESC 。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 一次性返回的记录数量，默认为100，最大为400。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSlowLogDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的记录总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 查询到的记录。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Items []*SlowLogItem `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量，最小值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认值为20，最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件的慢查询日志总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 符合查询条件的慢查询日志详情。
		Items []*SlowLogInfo `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSupportedPrivilegesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeSupportedPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSupportedPrivilegesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSupportedPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例支持的全局权限。
		GlobalSupportedPrivileges []*string `json:"GlobalSupportedPrivileges,omitempty" name:"GlobalSupportedPrivileges" list`

		// 实例支持的数据库权限。
		DatabaseSupportedPrivileges []*string `json:"DatabaseSupportedPrivileges,omitempty" name:"DatabaseSupportedPrivileges" list`

		// 实例支持的数据库表权限。
		TableSupportedPrivileges []*string `json:"TableSupportedPrivileges,omitempty" name:"TableSupportedPrivileges" list`

		// 实例支持的数据库列权限。
		ColumnSupportedPrivileges []*string `json:"ColumnSupportedPrivileges,omitempty" name:"ColumnSupportedPrivileges" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSupportedPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSupportedPrivilegesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库的名称。
	Database *string `json:"Database,omitempty" name:"Database"`

	// 记录偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为20，最大值为2000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 匹配数据库表名的正则表达式，规则同 MySQL 官网
	TableRegexp *string `json:"TableRegexp,omitempty" name:"TableRegexp"`
}

func (r *DescribeTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件的数据库表总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回的数据库表信息。
		Items []*string `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsOfInstanceIdsRequest struct {
	*tchttp.BaseRequest

	// 实例列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTagsOfInstanceIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTagsOfInstanceIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsOfInstanceIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页偏移量。
		Offset *int64 `json:"Offset,omitempty" name:"Offset"`

		// 分页大小。
		Limit *int64 `json:"Limit,omitempty" name:"Limit"`

		// 实例标签信息。
		Rows []*TagsInfoOfInstance `json:"Rows,omitempty" name:"Rows" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagsOfInstanceIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTagsOfInstanceIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 异步任务请求 ID，执行云数据库相关操作返回的 AsyncRequestId。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 任务类型，不传值则查询所有任务类型，支持的值包括：
	// 1 - 数据库回档；
	// 2 - SQL操作；
	// 3 - 数据导入；
	// 5 - 参数设置；
	// 6 - 初始化云数据库实例；
	// 7 - 重启云数据库实例；
	// 8 - 开启云数据库实例GTID；
	// 9 - 只读实例升级；
	// 10 - 数据库批量回档；
	// 11 - 主实例升级；
	// 12 - 删除云数据库库表；
	// 13 - 灾备实例提升为主。
	TaskTypes []*int64 `json:"TaskTypes,omitempty" name:"TaskTypes" list`

	// 任务状态，不传值则查询所有任务状态，支持的值包括：
	// -1 - 未定义；
	// 0 - 初始化；
	// 1 - 运行中；
	// 2 - 执行成功；
	// 3 - 执行失败；
	// 4 - 已终止；
	// 5 - 已删除；
	// 6 - 已暂停。
	TaskStatus []*int64 `json:"TaskStatus,omitempty" name:"TaskStatus" list`

	// 第一个任务的开始时间，用于范围查询，时间格式如：2017-12-31 10:40:01。
	StartTimeBegin *string `json:"StartTimeBegin,omitempty" name:"StartTimeBegin"`

	// 最后一个任务的开始时间，用于范围查询，时间格式如：2017-12-31 10:40:01。
	StartTimeEnd *string `json:"StartTimeEnd,omitempty" name:"StartTimeEnd"`

	// 记录偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件的实例总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回的实例任务信息。
		Items []*TaskDetail `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTimeWindowRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTimeWindowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 星期一的可维护时间列表。
		Monday []*string `json:"Monday,omitempty" name:"Monday" list`

		// 星期二的可维护时间列表。
		Tuesday []*string `json:"Tuesday,omitempty" name:"Tuesday" list`

		// 星期三的可维护时间列表。
		Wednesday []*string `json:"Wednesday,omitempty" name:"Wednesday" list`

		// 星期四的可维护时间列表。
		Thursday []*string `json:"Thursday,omitempty" name:"Thursday" list`

		// 星期五的可维护时间列表。
		Friday []*string `json:"Friday,omitempty" name:"Friday" list`

		// 星期六的可维护时间列表。
		Saturday []*string `json:"Saturday,omitempty" name:"Saturday" list`

		// 星期日的可维护时间列表。
		Sunday []*string `json:"Sunday,omitempty" name:"Sunday" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTimeWindowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUploadedFilesRequest struct {
	*tchttp.BaseRequest

	// 文件路径。该字段应填用户主账号的OwnerUin信息。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 记录偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeUploadedFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUploadedFilesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUploadedFilesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件的SQL文件总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回的SQL文件列表。
		Items []*SqlFileInfo `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUploadedFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUploadedFilesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeviceCpuInfo struct {

	// 实例CPU平均使用率
	Rate []*DeviceCpuRateInfo `json:"Rate,omitempty" name:"Rate" list`

	// 实例CPU监控数据
	Load []*int64 `json:"Load,omitempty" name:"Load" list`
}

type DeviceCpuRateInfo struct {

	// Cpu核编号
	CpuCore *int64 `json:"CpuCore,omitempty" name:"CpuCore"`

	// Cpu使用率
	Rate []*int64 `json:"Rate,omitempty" name:"Rate" list`
}

type DeviceDiskInfo struct {

	// 平均每秒有百分之几的时间用于IO操作
	IoRatioPerSec []*int64 `json:"IoRatioPerSec,omitempty" name:"IoRatioPerSec" list`

	// 平均每次设备I/O操作的等待时间*100，单位为毫秒。例如：该值为201，表示平均每次I/O操作等待时间为：201/100=2.1毫秒
	IoWaitTime []*int64 `json:"IoWaitTime,omitempty" name:"IoWaitTime" list`

	// 磁盘平均每秒完成的读操作次数总和*100。例如：该值为2002，表示磁盘平均每秒完成读操作为：2002/100=20.2次
	Read []*int64 `json:"Read,omitempty" name:"Read" list`

	// 磁盘平均每秒完成的写操作次数总和*100。例如：该值为30001，表示磁盘平均每秒完成写操作为：30001/100=300.01次
	Write []*int64 `json:"Write,omitempty" name:"Write" list`

	// 磁盘空间容量，每两个一组，第一个为已使用容量，第二个为磁盘总容量
	CapacityRatio []*int64 `json:"CapacityRatio,omitempty" name:"CapacityRatio" list`
}

type DeviceMemInfo struct {

	// 总内存大小。free命令中Mem:一行total的值,单位：KB
	Total []*int64 `json:"Total,omitempty" name:"Total" list`

	// 已使用内存。free命令中Mem:一行used的值,单位：KB
	Used []*int64 `json:"Used,omitempty" name:"Used" list`
}

type DeviceNetInfo struct {

	// tcp连接数
	Conn []*int64 `json:"Conn,omitempty" name:"Conn" list`

	// 网卡入包量，单位：个/秒
	PackageIn []*int64 `json:"PackageIn,omitempty" name:"PackageIn" list`

	// 网卡出包量，单位：个/秒
	PackageOut []*int64 `json:"PackageOut,omitempty" name:"PackageOut" list`

	// 入流量，单位：kbps
	FlowIn []*int64 `json:"FlowIn,omitempty" name:"FlowIn" list`

	// 出流量，单位：kbps
	FlowOut []*int64 `json:"FlowOut,omitempty" name:"FlowOut" list`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 安全组 ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例 ID 列表，一个或者多个实例 ID 组成的数组。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DrInfo struct {

	// 灾备实例状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 可用区信息
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 地域信息
	Region *string `json:"Region,omitempty" name:"Region"`

	// 实例同步状态。可能的返回值为：
	// 0 - 灾备未同步；
	// 1 - 灾备同步中；
	// 2 - 灾备同步成功；
	// 3 - 灾备同步失败；
	// 4 - 灾备同步修复中。
	SyncStatus *int64 `json:"SyncStatus,omitempty" name:"SyncStatus"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例类型
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`
}

type ErrlogItem struct {

	// 错误发生时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 错误详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`
}

type ImportRecord struct {

	// 状态值
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 状态值
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 执行时间
	CostTime *int64 `json:"CostTime,omitempty" name:"CostTime"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 后端任务ID
	WorkId *string `json:"WorkId,omitempty" name:"WorkId"`

	// 导入文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 执行进度
	Process *int64 `json:"Process,omitempty" name:"Process"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 文件大小
	FileSize *string `json:"FileSize,omitempty" name:"FileSize"`

	// 任务执行信息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 任务ID
	JobId *int64 `json:"JobId,omitempty" name:"JobId"`

	// 导入库表名
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 异步任务的请求ID
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

type Inbound struct {

	// 策略，ACCEPT 或者 DROP
	Action *string `json:"Action,omitempty" name:"Action"`

	// 来源 IP 或 IP 段，例如192.168.0.0/16
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// 端口
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// 网络协议，支持 UDP、TCP 等
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// 规则限定的方向，进站规则为 INPUT
	Dir *string `json:"Dir,omitempty" name:"Dir"`
}

type InitDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，可使用[查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 实例新的密码，密码规则：8-64个字符，至少包含字母、数字、字符（支持的字符：!@#$%^*()）中的两种。
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// 实例的参数列表，目前支持设置“character_set_server”、“lower_case_table_names”参数。其中，“character_set_server”参数可选值为["utf8","latin1","gbk","utf8mb4"]；“lower_case_table_names”可选值为[“0”,“1”]。
	Parameters []*ParamInfo `json:"Parameters,omitempty" name:"Parameters" list`

	// 实例的端口，取值范围为[1024, 65535]
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
}

func (r *InitDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InitDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的请求ID数组，可使用此ID查询异步任务的执行结果
		AsyncRequestIds []*string `json:"AsyncRequestIds,omitempty" name:"AsyncRequestIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpgradeInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 升级后的内存大小，单位：MB，为保证传入 Memory 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可升级的内存规格。
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 升级后的硬盘大小，单位：GB，为保证传入 Volume 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可升级的硬盘范围。
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// 升级后的核心数目，单位：核，为保证传入 CPU 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可升级的核心数目，当未指定该值时，将按照 Memory 大小补全一个默认值。
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 数据复制方式，支持值包括：0 - 异步复制，1 - 半同步复制，2 - 强同步复制，升级主实例时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义。
	ProtectMode *uint64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// 部署策略，取值范围：HA-高可用版两节点，FE-金融版三节点；默认值根据数据复制方式决定，数据复制方式为强同步复制时默认值为FE，数据复制方式非强同步时默认值为HA。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
}

func (r *InquiryPriceUpgradeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceUpgradeInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpgradeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例价格，单位：分（人民币）。
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// 实例原价，单位：分（人民币）。
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceUpgradeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceUpgradeInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstanceInfo struct {

	// 外网状态，可能的返回值为：0-未开通外网；1-已开通外网；2-已关闭外网
	WanStatus *int64 `json:"WanStatus,omitempty" name:"WanStatus"`

	// 可用区信息
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 初始化标志，可能的返回值为：0-未初始化；1-已初始化
	InitFlag *int64 `json:"InitFlag,omitempty" name:"InitFlag"`

	// 只读vip信息。单独开通只读实例访问的只读实例才有该字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoVipInfo *RoVipInfo `json:"RoVipInfo,omitempty" name:"RoVipInfo"`

	// 内存容量，单位为 MB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例状态，可能的返回值：0-创建中；1-运行中；4-隔离中；5-已隔离
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 私有网络 ID，例如：51102
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// 备机信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaveInfo *SlaveInfo `json:"SlaveInfo,omitempty" name:"SlaveInfo"`

	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 硬盘容量，单位为 GB
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// 自动续费标志，可能的返回值：0-未开通自动续费；1-已开通自动续费；2-已关闭自动续费
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 数据复制方式。0 - 异步复制；1 - 半同步复制；2 - 强同步复制
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// 只读组详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoGroups []*RoGroup `json:"RoGroups,omitempty" name:"RoGroups" list`

	// 子网 ID，例如：2333
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例类型，可能的返回值：1-主实例；2-灾备实例；3-只读实例
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// 项目 ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 地域信息
	Region *string `json:"Region,omitempty" name:"Region"`

	// 实例到期时间
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// 可用区部署方式。可能的值为：0 - 单可用区；1 - 多可用区
	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// 实例任务状态。0 - 没有任务 ,1 - 升级中,2 - 数据导入中,3 - 开放Slave中,4 - 外网访问开通中,5 - 批量操作执行中,6 - 回档中,7 - 外网访问关闭中,8 - 密码修改中,9 - 实例名修改中,10 - 重启中,12 - 自建迁移中,13 - 删除库表中,14 - 灾备实例创建同步中,15 - 升级待切换,16 - 升级切换中,17 - 升级切换完成
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 主实例详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterInfo *MasterInfo `json:"MasterInfo,omitempty" name:"MasterInfo"`

	// 实例类型，可能的返回值：“HA”-高可用版；“FE”-金融版；“BASIC”-基础版
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 内核版本
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 灾备实例详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DrInfo []*DrInfo `json:"DrInfo,omitempty" name:"DrInfo" list`

	// 外网域名
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`

	// 外网端口号
	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`

	// 付费类型，可能的返回值：0-包年包月；1-按量计费
	PayType *int64 `json:"PayType,omitempty" name:"PayType"`

	// 实例创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例 IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 端口号
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 是否锁定标记
	CdbError *int64 `json:"CdbError,omitempty" name:"CdbError"`

	// 私有网络描述符，例如：“vpc-5v8wn9mg”
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 子网描述符，例如：“subnet-1typ0s7d”
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 物理 ID
	PhysicalId *string `json:"PhysicalId,omitempty" name:"PhysicalId"`

	// 核心数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 每秒查询数量
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`

	// 可用区中文名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 物理机型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceClass *string `json:"DeviceClass,omitempty" name:"DeviceClass"`

	// 置放群组 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// 可用区 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type InstanceRebootTime struct {

	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 预期重启时间
	TimeInSeconds *int64 `json:"TimeInSeconds,omitempty" name:"TimeInSeconds"`
}

type InstanceRollbackRangeTime struct {

	// 查询数据库错误码
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 查询数据库错误信息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 实例ID列表，单个实例Id的格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 可回档时间范围
	Times []*RollbackTimeRange `json:"Times,omitempty" name:"Times" list`
}

type IsolateDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *IsolateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IsolateDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IsolateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。(该返回字段目前已废弃，可以通过 DescribeDBInstances 接口查询实例的隔离状态)
	// 注意：此字段可能返回 null，表示取不到有效值。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IsolateDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MasterInfo struct {

	// 地域信息
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 可用区ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 可用区信息
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例长ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 实例状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例类型
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// 任务状态
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 内存容量
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 硬盘容量
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// 实例机型
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 每秒查询数
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`

	// 私有网络ID
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// 独享集群ID
	ExClusterId *string `json:"ExClusterId,omitempty" name:"ExClusterId"`

	// 独享集群名称
	ExClusterName *string `json:"ExClusterName,omitempty" name:"ExClusterName"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 云数据库账号。
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts" list`

	// 数据库账号的备注信息。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAccountDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountDescriptionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountDescriptionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库账号的新密码。密码应至少包含字母、数字和字符（_+-&=!@#$%^*()）中的两种，长度为8-64个字符。
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// 云数据库账号。
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts" list`
}

func (r *ModifyAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库的账号，包括用户名和域名。
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts" list`

	// 全局权限。其中，GlobalPrivileges 中权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE", "PROCESS", "DROP","REFERENCES","INDEX","ALTER","SHOW DATABASES","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER"。
	// 注意，不传该参数表示清除该权限。
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitempty" name:"GlobalPrivileges" list`

	// 数据库的权限。Privileges 权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE",	"DROP","REFERENCES","INDEX","ALTER","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER"。
	// 注意，不传该参数表示清除该权限。
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitempty" name:"DatabasePrivileges" list`

	// 数据库中表的权限。Privileges 权限的可选值为：权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE",	"DROP","REFERENCES","INDEX","ALTER","CREATE VIEW","SHOW VIEW", "TRIGGER"。
	// 注意，不传该参数表示清除该权限。
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitempty" name:"TablePrivileges" list`

	// 数据库表中列的权限。Privileges 权限的可选值为："SELECT","INSERT","UPDATE","REFERENCES"。
	// 注意，不传该参数表示清除该权限。
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitempty" name:"ColumnPrivileges" list`
}

func (r *ModifyAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountPrivilegesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountPrivilegesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAuditConfigRequest struct {
	*tchttp.BaseRequest

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 审计日志保存时长。支持值包括：
	// 30 - 一个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年；
	LogExpireDay *int64 `json:"LogExpireDay,omitempty" name:"LogExpireDay"`

	// 是否关闭审计服务。可选值：true - 关闭审计服务；false - 不关闭审计服务。默认值为 false。
	// 当关闭审计服务时，会删除用户的审计日志和文件，并删除该实例的所有审计策略。
	// CloseAudit、LogExpireDay必须至少提供一个，如果两个都提供则按照CloseAudit优先的逻辑处理。
	CloseAudit *bool `json:"CloseAudit,omitempty" name:"CloseAudit"`
}

func (r *ModifyAuditConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAuditConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAuditConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAuditConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAuditConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAuditRuleRequest struct {
	*tchttp.BaseRequest

	// 审计规则 ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 审计规则名称。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 审计规则描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 审计规则过滤条件。若设置了过滤条件，则不会开启全审计。
	RuleFilters []*AuditFilter `json:"RuleFilters,omitempty" name:"RuleFilters" list`

	// 是否开启全审计。支持值包括：false – 不开启全审计，true – 开启全审计。用户未设置审计规则过滤条件时，默认开启全审计。
	AuditAll *bool `json:"AuditAll,omitempty" name:"AuditAll"`
}

func (r *ModifyAuditRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAuditRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAuditRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAuditRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAuditRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoRenewFlagRequest struct {
	*tchttp.BaseRequest

	// 实例的 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 自动续费标记，可取值的有：0 - 不自动续费，1 - 自动续费。
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`
}

func (r *ModifyAutoRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAutoRenewFlagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAutoRenewFlagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupConfigRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份文件的保留时间，单位为天。最小值为7天，最大值为732天。
	ExpireDays *int64 `json:"ExpireDays,omitempty" name:"ExpireDays"`

	// (将废弃，建议使用 BackupTimeWindow 参数) 备份时间范围，格式为：02:00-06:00，起点和终点时间目前限制为整点，目前可以选择的范围为： 00:00-12:00，02:00-06:00，06：00-10：00，10:00-14:00，14:00-18:00，18:00-22:00，22:00-02:00。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 自动备份方式，仅支持：physical - 物理冷备
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// binlog的保留时间，单位为天。最小值为7天，最大值为732天。该值的设置不能大于备份文件的保留时间。
	BinlogExpireDays *int64 `json:"BinlogExpireDays,omitempty" name:"BinlogExpireDays"`

	// 备份时间窗，比如要设置每周二和周日 10:00-14:00之间备份，该参数如下：{"Monday": "", "Tuesday": "10:00-14:00", "Wednesday": "", "Thursday": "", "Friday": "", "Saturday": "", "Sunday": "10:00-14:00"}    （注：可以设置一周的某几天备份，但是每天的备份时间需要设置为相同的时间段。 如果设置了该字段，将忽略StartTime字段的设置）
	BackupTimeWindow *CommonTimeWindow `json:"BackupTimeWindow,omitempty" name:"BackupTimeWindow"`
}

func (r *ModifyBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyBackupConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyBackupConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyDBInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceProjectRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 数组，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 项目的 ID。
	NewProjectId *int64 `json:"NewProjectId,omitempty" name:"NewProjectId"`
}

func (r *ModifyDBInstanceProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 要修改的安全组 ID 列表，一个或者多个安全组 ID 组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceVipVportRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 目标 IP。该参数和 DstPort 参数，两者必传一个。
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// 目标端口，支持范围为：[1024-65535]。该参数和 DstIp 参数，两者必传一个。
	DstPort *int64 `json:"DstPort,omitempty" name:"DstPort"`

	// 私有网络统一 ID。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 子网统一 ID。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 进行基础网络转 VPC 网络和 VPC 网络下的子网变更时，原网络中旧IP的回收时间，单位为小时，取值范围为0-168，默认值为24小时。
	ReleaseDuration *int64 `json:"ReleaseDuration,omitempty" name:"ReleaseDuration"`
}

func (r *ModifyDBInstanceVipVportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceVipVportRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceVipVportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务ID。(该返回字段目前已废弃)
	// 注意：此字段可能返回 null，表示取不到有效值。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceVipVportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceVipVportResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceParamRequest struct {
	*tchttp.BaseRequest

	// 实例短 ID 列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 要修改的参数列表。每一个元素是 Name 和 CurrentValue 的组合。Name 是参数名，CurrentValue 是要修改成的值。
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList" list`

	// 模板id，ParamList和TemplateId必须至少传其中之一
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 执行参数调整任务的方式，默认为 0。支持值包括：0 - 立刻执行，1 - 时间窗执行；当该值为 1 时，每次只能传一个实例（InstanceIds数量为1）
	WaitSwitch *int64 `json:"WaitSwitch,omitempty" name:"WaitSwitch"`
}

func (r *ModifyInstanceParamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceParamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceParamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务 ID，可用于查询任务进度。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceParamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceParamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceTagRequest struct {
	*tchttp.BaseRequest

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 要增加或修改的标签。
	ReplaceTags []*TagInfo `json:"ReplaceTags,omitempty" name:"ReplaceTags" list`

	// 要删除的标签。
	DeleteTags []*TagInfo `json:"DeleteTags,omitempty" name:"DeleteTags" list`
}

func (r *ModifyInstanceTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceTagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceTagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNameOrDescByDpIdRequest struct {
	*tchttp.BaseRequest

	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// 置放群组名称，最长不能超过60个字符。置放群组名和置放群组描述不能都为空。
	DeployGroupName *string `json:"DeployGroupName,omitempty" name:"DeployGroupName"`

	// 置放群组描述，最长不能超过200个字符。置放群组名和置放群组描述不能都为空。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyNameOrDescByDpIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNameOrDescByDpIdRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNameOrDescByDpIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNameOrDescByDpIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNameOrDescByDpIdResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyParamTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 参数列表。
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList" list`
}

func (r *ModifyParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyParamTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyParamTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRoGroupInfoRequest struct {
	*tchttp.BaseRequest

	// RO 组的 ID。
	RoGroupId *string `json:"RoGroupId,omitempty" name:"RoGroupId"`

	// RO 组的详细信息。
	RoGroupInfo *RoGroupAttr `json:"RoGroupInfo,omitempty" name:"RoGroupInfo"`

	// RO 组内实例的权重。若修改 RO 组的权重模式为用户自定义模式（custom），则必须设置该参数，且需要设置每个 RO 实例的权重值。
	RoWeightValues []*RoWeightValue `json:"RoWeightValues,omitempty" name:"RoWeightValues" list`

	// 是否重新均衡 RO 组内的 RO 实例的负载。支持值包括：1 - 重新均衡负载；0 - 不重新均衡负载。默认值为 0。注意，设置为重新均衡负载是，RO 组内 RO 实例会有一次数据库连接瞬断，请确保应用程序能重连数据库。
	IsBalanceRoLoad *int64 `json:"IsBalanceRoLoad,omitempty" name:"IsBalanceRoLoad"`
}

func (r *ModifyRoGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRoGroupInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRoGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRoGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRoGroupInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRoReplicationDelayRequest struct {
	*tchttp.BaseRequest

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 延迟时间（s）。最小值1，最大值259200。
	ReplicationDelay *int64 `json:"ReplicationDelay,omitempty" name:"ReplicationDelay"`
}

func (r *ModifyRoReplicationDelayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRoReplicationDelayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRoReplicationDelayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRoReplicationDelayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRoReplicationDelayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRoTypeRequest struct {
	*tchttp.BaseRequest

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 只读实例源类型，取值 NORMAL（普通只读实例）、DELAY_REPLICATION（延迟只读实例）。
	SrcRoInstType *string `json:"SrcRoInstType,omitempty" name:"SrcRoInstType"`

	// 只读实例目标类型，取值 NORMAL（普通只读实例）、DELAY_REPLICATION（延迟只读实例）。
	DstRoInstType *string `json:"DstRoInstType,omitempty" name:"DstRoInstType"`

	// 延迟时间（s），将实例修改为延迟只读实例时必传。最小值1，最大值259200。
	ReplicationDelay *int64 `json:"ReplicationDelay,omitempty" name:"ReplicationDelay"`
}

func (r *ModifyRoTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRoTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRoTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRoTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRoTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTimeWindowRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 修改后的可维护时间段，其中每一个时间段的格式形如：10:00-12:00；起止时间按半个小时对齐；最短半个小时，最长三个小时；最多设置两个时间段；起止时间范围为：[00:00, 24:00]。
	TimeRanges []*string `json:"TimeRanges,omitempty" name:"TimeRanges" list`

	// 指定修改哪一天的客户时间段，可能的取值为：monday，tuesday，wednesday，thursday，friday，saturday，sunday。如果不指定该值或者为空，则默认一周七天都修改。
	Weekdays []*string `json:"Weekdays,omitempty" name:"Weekdays" list`
}

func (r *ModifyTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTimeWindowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTimeWindowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OfflineIsolatedInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *OfflineIsolatedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OfflineIsolatedInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OfflineIsolatedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OfflineIsolatedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OfflineIsolatedInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenDBInstanceGTIDRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *OpenDBInstanceGTIDRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenDBInstanceGTIDRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenDBInstanceGTIDResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenDBInstanceGTIDResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenDBInstanceGTIDResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenWanServiceRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *OpenWanServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenWanServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenWanServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenWanServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenWanServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Outbound struct {

	// 策略，ACCEPT 或者 DROP
	Action *string `json:"Action,omitempty" name:"Action"`

	// 目的 IP 或 IP 段，例如172.16.0.0/12
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// 端口或者端口范围
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// 网络协议，支持 UDP、TCP等
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// 规则限定的方向，进站规则为 OUTPUT
	Dir *string `json:"Dir,omitempty" name:"Dir"`
}

type ParamInfo struct {

	// 参数名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ParamRecord struct {

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 参数名称
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 参数修改前的值
	OldValue *string `json:"OldValue,omitempty" name:"OldValue"`

	// 参数修改后的值
	NewValue *string `json:"NewValue,omitempty" name:"NewValue"`

	// 参数是否修改成功
	IsSucess *bool `json:"IsSucess,omitempty" name:"IsSucess"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type ParamTemplateInfo struct {

	// 参数模板ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 参数模板名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数模板描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 实例引擎版本
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
}

type Parameter struct {

	// 参数名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`
}

type ParameterDetail struct {

	// 参数名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数类型
	ParamType *string `json:"ParamType,omitempty" name:"ParamType"`

	// 参数默认值
	Default *string `json:"Default,omitempty" name:"Default"`

	// 参数描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 参数当前值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 修改参数后，是否需要重启数据库以使参数生效。可能的值包括：0-不需要重启；1-需要重启
	NeedReboot *int64 `json:"NeedReboot,omitempty" name:"NeedReboot"`

	// 参数允许的最大值
	Max *int64 `json:"Max,omitempty" name:"Max"`

	// 参数允许的最小值
	Min *int64 `json:"Min,omitempty" name:"Min"`

	// 参数的可选枚举值。如果为非枚举参数，则为空
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue" list`
}

type RegionSellConf struct {

	// 地域中文名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 所属大区
	Area *string `json:"Area,omitempty" name:"Area"`

	// 是否为默认地域
	IsDefaultRegion *int64 `json:"IsDefaultRegion,omitempty" name:"IsDefaultRegion"`

	// 地域名称
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区售卖配置
	ZonesConf []*ZoneSellConf `json:"ZonesConf,omitempty" name:"ZonesConf" list`
}

type ReleaseIsolatedDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 数组，单个实例 ID 格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *ReleaseIsolatedDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseIsolatedDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseIsolatedDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 解隔离操作的结果集。
		Items []*ReleaseResult `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseIsolatedDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseIsolatedDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseResult struct {

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例解隔离操作的结果值。返回值为0表示成功。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 实例解隔离操作的错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type RenewDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 待续费的实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872)。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 续费时长，单位：月，可选值包括 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]。
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 如果需要将按量计费实例续费为包年包月的实例，该入参的值需要指定为 "PREPAID" 。
	ModifyPayType *string `json:"ModifyPayType,omitempty" name:"ModifyPayType"`
}

func (r *RenewDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单 ID。
		DealId *string `json:"DealId,omitempty" name:"DealId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RestartDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 数组，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *RestartDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestartDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RestartDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestartDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RoGroup struct {

	// 只读组模式，可选值为：alone-系统自动分配只读组；allinone-新建只读组；join-使用现有只读组。
	RoGroupMode *string `json:"RoGroupMode,omitempty" name:"RoGroupMode"`

	// 只读组 ID。
	RoGroupId *string `json:"RoGroupId,omitempty" name:"RoGroupId"`

	// 只读组名称。
	RoGroupName *string `json:"RoGroupName,omitempty" name:"RoGroupName"`

	// 是否启用延迟超限剔除功能，启用该功能后，只读实例与主实例的延迟超过延迟阈值，只读实例将被隔离。可选值：1-启用；0-不启用。
	RoOfflineDelay *int64 `json:"RoOfflineDelay,omitempty" name:"RoOfflineDelay"`

	// 延迟阈值。
	RoMaxDelayTime *int64 `json:"RoMaxDelayTime,omitempty" name:"RoMaxDelayTime"`

	// 最少实例保留个数，若购买只读实例数量小于设置数量将不做剔除。
	MinRoInGroup *int64 `json:"MinRoInGroup,omitempty" name:"MinRoInGroup"`

	// 读写权重分配模式，可选值：system-系统自动分配；custom-自定义。
	WeightMode *string `json:"WeightMode,omitempty" name:"WeightMode"`

	// 权重值。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 只读组中的只读实例详情。
	RoInstances []*RoInstanceInfo `json:"RoInstances,omitempty" name:"RoInstances" list`

	// 只读组的内网 IP。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 只读组的内网端口号。
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 私有网络 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 子网 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 只读组所在的地域。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoGroupRegion *string `json:"RoGroupRegion,omitempty" name:"RoGroupRegion"`

	// 只读组所在的可用区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoGroupZone *string `json:"RoGroupZone,omitempty" name:"RoGroupZone"`
}

type RoGroupAttr struct {

	// RO 组名称。
	RoGroupName *string `json:"RoGroupName,omitempty" name:"RoGroupName"`

	// RO 实例最大延迟阈值。单位为秒，最小值为 1。注意，RO 组必须设置了开启实例延迟剔除策略，该值才有效。
	RoMaxDelayTime *int64 `json:"RoMaxDelayTime,omitempty" name:"RoMaxDelayTime"`

	// 是否开启实例延迟剔除。支持的值包括：1 - 开启；0 - 不开启。注意，若设置开启实例延迟剔除，则必须设置延迟阈值（RoMaxDelayTime）参数。
	RoOfflineDelay *int64 `json:"RoOfflineDelay,omitempty" name:"RoOfflineDelay"`

	// 最少保留实例数。可设置为小于或等于该 RO 组下 RO 实例个数的任意值。注意，若设置值大于 RO 实例数量将不做剔除；若设置为 0，所有实例延迟超限都会被剔除。
	MinRoInGroup *int64 `json:"MinRoInGroup,omitempty" name:"MinRoInGroup"`

	// 权重模式。支持值包括："system" - 系统自动分配； "custom" - 用户自定义设置。注意，若设置 "custom" 模式，则必须设置 RO 实例权重配置（RoWeightValues）参数。
	WeightMode *string `json:"WeightMode,omitempty" name:"WeightMode"`
}

type RoInstanceInfo struct {

	// RO组对应的主实例的ID
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" name:"MasterInstanceId"`

	// RO实例在RO组内的状态，可能的值：online-在线，offline-下线
	RoStatus *string `json:"RoStatus,omitempty" name:"RoStatus"`

	// RO实例在RO组内上一次下线的时间
	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`

	// RO实例在RO组内的权重
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// RO实例所在区域名称，如ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`

	// RO可用区的正式名称，如ap-shanghai-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// RO实例ID，格式如：cdbro-c1nl9rpv
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// RO实例状态，可能返回值：0-创建中，1-运行中，4-删除中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例类型，可能返回值：1-主实例，2-灾备实例，3-只读实例
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// RO实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 按量计费状态，可能的取值：1-正常，2-欠费
	HourFeeStatus *int64 `json:"HourFeeStatus,omitempty" name:"HourFeeStatus"`

	// RO实例任务状态，可能返回值：<br>0-没有任务<br>1-升级中<br>2-数据导入中<br>3-开放Slave中<br>4-外网访问开通中<br>5-批量操作执行中<br>6-回档中<br>7-外网访问关闭中<br>8-密码修改中<br>9-实例名修改中<br>10-重启中<br>12-自建迁移中<br>13-删除库表中<br>14-灾备实例创建同步中
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// RO实例内存大小，单位：MB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// RO实例硬盘大小，单位：GB
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// 每次查询数量
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`

	// RO实例的内网IP地址
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// RO实例访问端口
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// RO实例所在私有网络ID
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// RO实例所在私有网络子网ID
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// RO实例规格描述，目前可取值 CUSTOM
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// RO实例数据库引擎版本，可能返回值：5.1、5.5、5.6和5.7
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// RO实例到期时间，时间格式：yyyy-mm-dd hh:mm:ss，如实例为按量计费模式，则此字段值为0000-00-00 00:00:00
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// RO实例计费类型，可能返回值：0-包年包月，1-按量计费，2-后付费月结
	PayType *int64 `json:"PayType,omitempty" name:"PayType"`
}

type RoVipInfo struct {

	// 只读vip状态
	RoVipStatus *int64 `json:"RoVipStatus,omitempty" name:"RoVipStatus"`

	// 只读vip的子网
	RoSubnetId *int64 `json:"RoSubnetId,omitempty" name:"RoSubnetId"`

	// 只读vip的私有网络
	RoVpcId *int64 `json:"RoVpcId,omitempty" name:"RoVpcId"`

	// 只读vip的端口号
	RoVport *int64 `json:"RoVport,omitempty" name:"RoVport"`

	// 只读vip
	RoVip *string `json:"RoVip,omitempty" name:"RoVip"`
}

type RoWeightValue struct {

	// RO 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 权重值。取值范围为 [0, 100]。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type RollbackDBName struct {

	// 回档前的原数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 回档后的新数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewDatabaseName *string `json:"NewDatabaseName,omitempty" name:"NewDatabaseName"`
}

type RollbackInstancesInfo struct {

	// 云数据库实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 回档策略。可选值为：table、db、full；默认值为full。table - 急速回档模式，仅导入所选中表级别的备份和binlog，如有跨表操作，且关联表未被同时选中，将会导致回档失败，该模式下参数Databases必须为空；db - 快速模式，仅导入所选中库级别的备份和binlog，如有跨库操作，且关联库未被同时选中，将会导致回档失败；full - 普通回档模式，将导入整个实例的备份和binlog，速度较慢。
	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`

	// 数据库回档时间，时间格式为：yyyy-mm-dd hh:mm:ss
	RollbackTime *string `json:"RollbackTime,omitempty" name:"RollbackTime"`

	// 待回档的数据库信息，表示整库回档
	// 注意：此字段可能返回 null，表示取不到有效值。
	Databases []*RollbackDBName `json:"Databases,omitempty" name:"Databases" list`

	// 待回档的数据库表信息，表示按表回档
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tables []*RollbackTables `json:"Tables,omitempty" name:"Tables" list`
}

type RollbackTableName struct {

	// 回档前的原数据库表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 回档后的新数据库表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewTableName *string `json:"NewTableName,omitempty" name:"NewTableName"`
}

type RollbackTables struct {

	// 数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Database *string `json:"Database,omitempty" name:"Database"`

	// 数据库表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Table []*RollbackTableName `json:"Table,omitempty" name:"Table" list`
}

type RollbackTask struct {

	// 任务执行信息描述。
	Info *string `json:"Info,omitempty" name:"Info"`

	// 任务执行结果。可能的取值：INITIAL - 初始化，RUNNING - 运行中，SUCCESS - 执行成功，FAILED - 执行失败，KILLED - 已终止，REMOVED - 已删除，PAUSED - 终止中。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务执行进度。取值范围为[0, 100]。
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// 任务开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 回档任务详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*RollbackInstancesInfo `json:"Detail,omitempty" name:"Detail" list`
}

type RollbackTimeRange struct {

	// 实例可回档开始时间，时间格式：2016-10-29 01:06:04
	Begin *string `json:"Begin,omitempty" name:"Begin"`

	// 实例可回档结束时间，时间格式：2016-11-02 11:44:47
	End *string `json:"End,omitempty" name:"End"`
}

type SecurityGroup struct {

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 创建时间，时间格式：yyyy-mm-dd hh:mm:ss
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 入站规则
	Inbound []*Inbound `json:"Inbound,omitempty" name:"Inbound" list`

	// 出站规则
	Outbound []*Outbound `json:"Outbound,omitempty" name:"Outbound" list`

	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组名称
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// 安全组备注
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`
}

type SellConfig struct {

	// 设备类型
	Device *string `json:"Device,omitempty" name:"Device"`

	// 售卖规格描述
	Type *string `json:"Type,omitempty" name:"Type"`

	// 实例类型
	CdbType *string `json:"CdbType,omitempty" name:"CdbType"`

	// 内存大小，单位为MB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// CPU核心数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 磁盘最小规格，单位为GB
	VolumeMin *int64 `json:"VolumeMin,omitempty" name:"VolumeMin"`

	// 磁盘最大规格，单位为GB
	VolumeMax *int64 `json:"VolumeMax,omitempty" name:"VolumeMax"`

	// 磁盘步长，单位为GB
	VolumeStep *int64 `json:"VolumeStep,omitempty" name:"VolumeStep"`

	// 链接数
	Connection *int64 `json:"Connection,omitempty" name:"Connection"`

	// 每秒查询数量
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`

	// 每秒IO数量
	Iops *int64 `json:"Iops,omitempty" name:"Iops"`

	// 应用场景描述
	Info *string `json:"Info,omitempty" name:"Info"`

	// 状态值
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 标签值
	Tag *int64 `json:"Tag,omitempty" name:"Tag"`
}

type SellType struct {

	// 售卖实例名称
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// 内核版本号
	EngineVersion []*string `json:"EngineVersion,omitempty" name:"EngineVersion" list`

	// 售卖规格详细配置
	Configs []*SellConfig `json:"Configs,omitempty" name:"Configs" list`
}

type SlaveConfig struct {

	// 从库复制方式，可能的返回值：aysnc-异步，semisync-半同步
	ReplicationMode *string `json:"ReplicationMode,omitempty" name:"ReplicationMode"`

	// 从库可用区的正式名称，如ap-shanghai-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type SlaveInfo struct {

	// 第一备机信息
	First *SlaveInstanceInfo `json:"First,omitempty" name:"First"`

	// 第二备机信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Second *SlaveInstanceInfo `json:"Second,omitempty" name:"Second"`
}

type SlaveInstanceInfo struct {

	// 端口号
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 地域信息
	Region *string `json:"Region,omitempty" name:"Region"`

	// 虚拟 IP 信息
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 可用区信息
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type SlowLogInfo struct {

	// 备份文件名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备份文件大小，单位：Byte
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 备份快照时间，时间格式：2016-03-17 02:10:37
	Date *string `json:"Date,omitempty" name:"Date"`

	// 内网下载地址
	IntranetUrl *string `json:"IntranetUrl,omitempty" name:"IntranetUrl"`

	// 外网下载地址
	InternetUrl *string `json:"InternetUrl,omitempty" name:"InternetUrl"`

	// 日志具体类型，可能的值：slowlog - 慢日志
	Type *string `json:"Type,omitempty" name:"Type"`
}

type SlowLogItem struct {

	// Sql的执行时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// Sql的执行时长。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryTime *float64 `json:"QueryTime,omitempty" name:"QueryTime"`

	// Sql语句。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// 客户端地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserHost *string `json:"UserHost,omitempty" name:"UserHost"`

	// 用户名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 数据库名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Database *string `json:"Database,omitempty" name:"Database"`

	// 锁时长。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LockTime *float64 `json:"LockTime,omitempty" name:"LockTime"`

	// 扫描行数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RowsExamined *int64 `json:"RowsExamined,omitempty" name:"RowsExamined"`

	// 结果集行数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RowsSent *int64 `json:"RowsSent,omitempty" name:"RowsSent"`

	// Sql模板。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SqlTemplate *string `json:"SqlTemplate,omitempty" name:"SqlTemplate"`

	// Sql语句的md5。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

type SqlFileInfo struct {

	// 上传时间
	UploadTime *string `json:"UploadTime,omitempty" name:"UploadTime"`

	// 上传进度
	UploadInfo *UploadInfo `json:"UploadInfo,omitempty" name:"UploadInfo"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件大小，单位为Bytes
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// 上传是否完成标志，可选值：0 - 未完成，1 - 已完成
	IsUploadFinished *int64 `json:"IsUploadFinished,omitempty" name:"IsUploadFinished"`

	// 文件ID
	FileId *string `json:"FileId,omitempty" name:"FileId"`
}

type StartBatchRollbackRequest struct {
	*tchttp.BaseRequest

	// 用于回档的实例详情信息。
	Instances []*RollbackInstancesInfo `json:"Instances,omitempty" name:"Instances" list`
}

func (r *StartBatchRollbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartBatchRollbackRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartBatchRollbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartBatchRollbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartBatchRollbackResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartDelayReplicationRequest struct {
	*tchttp.BaseRequest

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 延迟复制类型。可选值 DEFAULT（按照延迟复制时间进行复制）、GTID（回放到指定GTID）、DUE_TIME（回放到指定时间点）。
	DelayReplicationType *string `json:"DelayReplicationType,omitempty" name:"DelayReplicationType"`

	// 指定时间点，默认为0，最大值不能超过当前时间。
	DueTime *int64 `json:"DueTime,omitempty" name:"DueTime"`

	// 指定GITD。回放到指定GTID时必传。
	Gtid *string `json:"Gtid,omitempty" name:"Gtid"`
}

func (r *StartDelayReplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartDelayReplicationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartDelayReplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 延迟复制任务 ID。DelayReplicationType不为DEFAULT时返回，可用来查询回放任务状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartDelayReplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartDelayReplicationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopDBImportJobRequest struct {
	*tchttp.BaseRequest

	// 异步任务的请求 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

func (r *StopDBImportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopDBImportJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopDBImportJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopDBImportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopDBImportJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopDelayReplicationRequest struct {
	*tchttp.BaseRequest

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *StopDelayReplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopDelayReplicationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopDelayReplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopDelayReplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopDelayReplicationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopRollbackRequest struct {
	*tchttp.BaseRequest

	// 撤销回档任务对应的实例Id。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *StopRollbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopRollbackRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopRollbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 执行请求的异步任务ID
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopRollbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopRollbackResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchForUpgradeRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *SwitchForUpgradeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchForUpgradeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchForUpgradeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchForUpgradeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchForUpgradeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TableName struct {

	// 表名
	TableName *string `json:"TableName,omitempty" name:"TableName"`
}

type TablePrivilege struct {

	// 数据库名
	Database *string `json:"Database,omitempty" name:"Database"`

	// 数据库表名
	Table *string `json:"Table,omitempty" name:"Table"`

	// 权限信息
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges" list`
}

type TagInfo struct {

	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue []*string `json:"TagValue,omitempty" name:"TagValue" list`
}

type TagInfoUnit struct {

	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagsInfoOfInstance struct {

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 标签信息
	Tags []*TagInfoUnit `json:"Tags,omitempty" name:"Tags" list`
}

type TaskDetail struct {

	// 错误码。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 实例任务 ID。
	JobId *int64 `json:"JobId,omitempty" name:"JobId"`

	// 实例任务进度。
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// 实例任务状态，可能的值包括：
	// "UNDEFINED" - 未定义；
	// "INITIAL" - 初始化；
	// "RUNNING" - 运行中；
	// "SUCCEED" - 执行成功；
	// "FAILED" - 执行失败；
	// "KILLED" - 已终止；
	// "REMOVED" - 已删除；
	// "PAUSED" - 已暂停。
	// "WAITING" - 等待中（可撤销）
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 实例任务类型，可能的值包括：
	// "ROLLBACK" - 数据库回档；
	// "SQL OPERATION" - SQL操作；
	// "IMPORT DATA" - 数据导入；
	// "MODIFY PARAM" - 参数设置；
	// "INITIAL" - 初始化云数据库实例；
	// "REBOOT" - 重启云数据库实例；
	// "OPEN GTID" - 开启云数据库实例GTID；
	// "UPGRADE RO" - 只读实例升级；
	// "BATCH ROLLBACK" - 数据库批量回档；
	// "UPGRADE MASTER" - 主实例升级；
	// "DROP TABLES" - 删除云数据库库表；
	// "SWITCH DR TO MASTER" - 灾备实例提升为主。
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 实例任务开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 实例任务结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务关联的实例 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 异步任务的请求 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

type UpgradeDBInstanceEngineVersionRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主实例数据库引擎版本，支持值包括：5.6 和 5.7。
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 切换访问新实例的方式，默认为 0。支持值包括：0 - 立刻切换，1 - 时间窗切换；当该值为 1 时，升级中过程中，切换访问新实例的流程将会在时间窗内进行，或者用户主动调用接口 [切换访问新实例](https://cloud.tencent.com/document/product/236/15864) 触发该流程。
	WaitSwitch *int64 `json:"WaitSwitch,omitempty" name:"WaitSwitch"`

	// 是否是内核子版本升级，支持的值：1 - 升级内核子版本；0 - 升级数据库引擎版本。
	UpgradeSubversion *int64 `json:"UpgradeSubversion,omitempty" name:"UpgradeSubversion"`
}

func (r *UpgradeDBInstanceEngineVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDBInstanceEngineVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceEngineVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务 ID，可使用 [查询异步任务的执行结果](https://cloud.tencent.com/document/api/236/20410) 获取其执行情况。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeDBInstanceEngineVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDBInstanceEngineVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 升级后的内存大小，单位：MB，为保证传入 Memory 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可升级的内存规格。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 升级后的硬盘大小，单位：GB，为保证传入 Volume 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可升级的硬盘范围。
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// 数据复制方式，支持值包括：0 - 异步复制，1 - 半同步复制，2 - 强同步复制，升级主实例时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义。
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// 部署模式，默认为 0，支持值包括：0 - 单可用区部署，1 - 多可用区部署，升级主实例时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义。
	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// 备库1的可用区信息，默认和实例的 Zone 参数一致，升级主实例为多可用区部署时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义。可通过 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口查询支持的可用区。
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// 主实例数据库引擎版本，支持值包括：5.5、5.6 和 5.7。
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 切换访问新实例的方式，默认为 0。支持值包括：0 - 立刻切换，1 - 时间窗切换；当该值为 1 时，升级中过程中，切换访问新实例的流程将会在时间窗内进行，或者用户主动调用接口 [切换访问新实例](https://cloud.tencent.com/document/product/236/15864) 触发该流程。
	WaitSwitch *int64 `json:"WaitSwitch,omitempty" name:"WaitSwitch"`

	// 备库 2 的可用区信息，默认为空，升级主实例时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义。
	BackupZone *string `json:"BackupZone,omitempty" name:"BackupZone"`

	// 实例类型，默认为 master，支持值包括：master - 表示主实例，dr - 表示灾备实例，ro - 表示只读实例。
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
}

func (r *UpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单 ID。
		DealIds []*string `json:"DealIds,omitempty" name:"DealIds" list`

		// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UploadInfo struct {

	// 文件所有分片数
	AllSliceNum *int64 `json:"AllSliceNum,omitempty" name:"AllSliceNum"`

	// 已完成分片数
	CompleteNum *int64 `json:"CompleteNum,omitempty" name:"CompleteNum"`
}

type VerifyRootAccountRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例 ROOT 账号的密码。
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *VerifyRootAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VerifyRootAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VerifyRootAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyRootAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VerifyRootAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ZoneConf struct {

	// 可用区部署方式，可能的值为：0-单可用区；1-多可用区
	DeployMode []*int64 `json:"DeployMode,omitempty" name:"DeployMode" list`

	// 主实例所在的可用区
	MasterZone []*string `json:"MasterZone,omitempty" name:"MasterZone" list`

	// 实例为多可用区部署时，备库1所在的可用区
	SlaveZone []*string `json:"SlaveZone,omitempty" name:"SlaveZone" list`

	// 实例为多可用区部署时，备库2所在的可用区
	BackupZone []*string `json:"BackupZone,omitempty" name:"BackupZone" list`
}

type ZoneSellConf struct {

	// 可用区状态。可能的返回值为：0-未上线；1-上线；2-开放；3-停售；4-不展示
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 可用区中文名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 实例类型是否为自定义类型
	IsCustom *bool `json:"IsCustom,omitempty" name:"IsCustom"`

	// 是否支持灾备
	IsSupportDr *bool `json:"IsSupportDr,omitempty" name:"IsSupportDr"`

	// 是否支持私有网络
	IsSupportVpc *bool `json:"IsSupportVpc,omitempty" name:"IsSupportVpc"`

	// 小时计费实例最大售卖数量
	HourInstanceSaleMaxNum *int64 `json:"HourInstanceSaleMaxNum,omitempty" name:"HourInstanceSaleMaxNum"`

	// 是否为默认可用区
	IsDefaultZone *bool `json:"IsDefaultZone,omitempty" name:"IsDefaultZone"`

	// 是否为黑石区
	IsBm *bool `json:"IsBm,omitempty" name:"IsBm"`

	// 支持的付费类型。可能的返回值为：0-包年包月；1-小时计费；2-后付费
	PayType []*string `json:"PayType,omitempty" name:"PayType" list`

	// 数据复制类型。0-异步复制；1-半同步复制；2-强同步复制
	ProtectMode []*string `json:"ProtectMode,omitempty" name:"ProtectMode" list`

	// 可用区名称
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 售卖实例类型数组
	SellType []*SellType `json:"SellType,omitempty" name:"SellType" list`

	// 多可用区信息
	ZoneConf *ZoneConf `json:"ZoneConf,omitempty" name:"ZoneConf"`

	// 可支持的灾备可用区信息
	DrZone []*string `json:"DrZone,omitempty" name:"DrZone" list`

	// 是否支持跨可用区只读
	IsSupportRemoteRo *bool `json:"IsSupportRemoteRo,omitempty" name:"IsSupportRemoteRo"`

	// 可支持的跨可用区只读区信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemoteRoZone []*string `json:"RemoteRoZone,omitempty" name:"RemoteRoZone" list`
}
