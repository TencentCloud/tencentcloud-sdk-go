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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
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

	// 该值已废弃
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 用户最大可用实例连接数
	MaxUserConnections *int64 `json:"MaxUserConnections,omitempty" name:"MaxUserConnections"`
}

// Predefined struct for user
type AddTimeWindowRequestParams struct {
	// 实例ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 星期一的可维护时间段，其中每一个时间段的格式形如：10:00-12:00；起始时间按半个小时对齐；最短半个小时，最长三个小时；可设置多个时间段。 一周中应至少设置一天的时间窗。下同。
	Monday []*string `json:"Monday,omitempty" name:"Monday"`

	// 星期二的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Tuesday []*string `json:"Tuesday,omitempty" name:"Tuesday"`

	// 星期三的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Wednesday []*string `json:"Wednesday,omitempty" name:"Wednesday"`

	// 星期四的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Thursday []*string `json:"Thursday,omitempty" name:"Thursday"`

	// 星期五的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Friday []*string `json:"Friday,omitempty" name:"Friday"`

	// 星期六的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Saturday []*string `json:"Saturday,omitempty" name:"Saturday"`

	// 星期日的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Sunday []*string `json:"Sunday,omitempty" name:"Sunday"`

	// 最大延迟阈值，仅对主实例和灾备实例有效
	MaxDelayTime *uint64 `json:"MaxDelayTime,omitempty" name:"MaxDelayTime"`
}

type AddTimeWindowRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 星期一的可维护时间段，其中每一个时间段的格式形如：10:00-12:00；起始时间按半个小时对齐；最短半个小时，最长三个小时；可设置多个时间段。 一周中应至少设置一天的时间窗。下同。
	Monday []*string `json:"Monday,omitempty" name:"Monday"`

	// 星期二的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Tuesday []*string `json:"Tuesday,omitempty" name:"Tuesday"`

	// 星期三的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Wednesday []*string `json:"Wednesday,omitempty" name:"Wednesday"`

	// 星期四的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Thursday []*string `json:"Thursday,omitempty" name:"Thursday"`

	// 星期五的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Friday []*string `json:"Friday,omitempty" name:"Friday"`

	// 星期六的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Saturday []*string `json:"Saturday,omitempty" name:"Saturday"`

	// 星期日的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Sunday []*string `json:"Sunday,omitempty" name:"Sunday"`

	// 最大延迟阈值，仅对主实例和灾备实例有效
	MaxDelayTime *uint64 `json:"MaxDelayTime,omitempty" name:"MaxDelayTime"`
}

func (r *AddTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddTimeWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Monday")
	delete(f, "Tuesday")
	delete(f, "Wednesday")
	delete(f, "Thursday")
	delete(f, "Friday")
	delete(f, "Saturday")
	delete(f, "Sunday")
	delete(f, "MaxDelayTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddTimeWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddTimeWindowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *AddTimeWindowResponseParams `json:"Response"`
}

func (r *AddTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddTimeWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Address struct {
	// 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	VPort *uint64 `json:"VPort,omitempty" name:"VPort"`

	// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 私有网络子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqSubnet *string `json:"UniqSubnet,omitempty" name:"UniqSubnet"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

// Predefined struct for user
type AssociateSecurityGroupsRequestParams struct {
	// 安全组 ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例 ID 列表，一个或者多个实例 ID 组成的数组。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 当传入只读实例ID时，默认操作的是对应只读组的安全组。如果需要操作只读实例ID的安全组， 需要将该入参置为True
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitempty" name:"ForReadonlyInstance"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 安全组 ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例 ID 列表，一个或者多个实例 ID 组成的数组。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 当传入只读实例ID时，默认操作的是对应只读组的安全组。如果需要操作只读实例ID的安全组， 需要将该入参置为True
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitempty" name:"ForReadonlyInstance"`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIds")
	delete(f, "ForReadonlyInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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
	Host []*string `json:"Host,omitempty" name:"Host"`

	// 用户名。
	User []*string `json:"User,omitempty" name:"User"`

	// 数据库名称。
	DBName []*string `json:"DBName,omitempty" name:"DBName"`

	// 表名称。
	TableName []*string `json:"TableName,omitempty" name:"TableName"`

	// 审计策略名称。
	PolicyName []*string `json:"PolicyName,omitempty" name:"PolicyName"`

	// SQL 语句。支持模糊匹配。
	Sql *string `json:"Sql,omitempty" name:"Sql"`

	// SQL 类型。目前支持："SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "ALTER", "SET", "REPLACE", "EXECUTE"。
	SqlType *string `json:"SqlType,omitempty" name:"SqlType"`

	// 执行时间。单位为：ms。表示筛选执行时间大于该值的审计日志。
	ExecTime *int64 `json:"ExecTime,omitempty" name:"ExecTime"`

	// 影响行数。表示筛选影响行数大于该值的审计日志。
	AffectRows *int64 `json:"AffectRows,omitempty" name:"AffectRows"`

	// SQL 类型。支持多个类型同时查询。目前支持："SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "ALTER", "SET", "REPLACE", "EXECUTE"。
	SqlTypes []*string `json:"SqlTypes,omitempty" name:"SqlTypes"`

	// SQL 语句。支持传递多个sql语句。
	Sqls []*string `json:"Sqls,omitempty" name:"Sqls"`
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

	// 数据库实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
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
	RuleFilters []*AuditFilter `json:"RuleFilters,omitempty" name:"RuleFilters"`

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

	// 下载地址
	IntranetUrl *string `json:"IntranetUrl,omitempty" name:"IntranetUrl"`

	// 下载地址
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

	// 手动备份别名
	ManualBackupName *string `json:"ManualBackupName,omitempty" name:"ManualBackupName"`

	// 备份保留类型，save_mode_regular - 常规保存备份，save_mode_period - 定期保存备份
	SaveMode *string `json:"SaveMode,omitempty" name:"SaveMode"`

	// 本地备份所在地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 异地备份详细信息
	RemoteInfo []*RemoteBackupInfo `json:"RemoteInfo,omitempty" name:"RemoteInfo"`

	// 存储方式，0-常规存储，1-归档存储，默认为0
	CosStorageType *int64 `json:"CosStorageType,omitempty" name:"CosStorageType"`

	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份文件是否加密， on-加密， off-未加密
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptionFlag *string `json:"EncryptionFlag,omitempty" name:"EncryptionFlag"`
}

type BackupItem struct {
	// 需要备份的库名
	Db *string `json:"Db,omitempty" name:"Db"`

	// 需要备份的表名。 如果传该参数，表示备份该库中的指定表。如果不传该参数则备份该db库
	Table *string `json:"Table,omitempty" name:"Table"`
}

type BackupLimitVpcItem struct {
	// 限制下载来源的地域。目前仅支持当前地域。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 限制下载的vpc列表。
	VpcList []*string `json:"VpcList,omitempty" name:"VpcList"`
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

// Predefined struct for user
type BalanceRoGroupLoadRequestParams struct {
	// RO 组的 ID，格式如：cdbrg-c1nl9rpv。
	RoGroupId *string `json:"RoGroupId,omitempty" name:"RoGroupId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BalanceRoGroupLoadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BalanceRoGroupLoadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BalanceRoGroupLoadResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BalanceRoGroupLoadResponse struct {
	*tchttp.BaseResponse
	Response *BalanceRoGroupLoadResponseParams `json:"Response"`
}

func (r *BalanceRoGroupLoadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BalanceRoGroupLoadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaseGroupInfo struct {
	// 代理组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 代理节点数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// 状态：发货中（init）运行中（online）下线中（offline）销毁中（destroy）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 是否开启读写分离
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenRW *bool `json:"OpenRW,omitempty" name:"OpenRW"`

	// 当前代理版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitempty" name:"CurrentProxyVersion"`

	// 支持升级版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportUpgradeProxyVersion *string `json:"SupportUpgradeProxyVersion,omitempty" name:"SupportUpgradeProxyVersion"`
}

type BinlogInfo struct {
	// binlog 日志备份文件名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备份文件大小，单位：Byte
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 文件存储时间，时间格式：2016-03-17 02:10:37
	Date *string `json:"Date,omitempty" name:"Date"`

	// 下载地址
	IntranetUrl *string `json:"IntranetUrl,omitempty" name:"IntranetUrl"`

	// 下载地址
	InternetUrl *string `json:"InternetUrl,omitempty" name:"InternetUrl"`

	// 日志具体类型，可能的值有：binlog - 二进制日志
	Type *string `json:"Type,omitempty" name:"Type"`

	// binlog 文件起始时间
	BinlogStartTime *string `json:"BinlogStartTime,omitempty" name:"BinlogStartTime"`

	// binlog 文件截止时间
	BinlogFinishTime *string `json:"BinlogFinishTime,omitempty" name:"BinlogFinishTime"`

	// 本地binlog文件所在地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 备份任务状态。可能的值有 "SUCCESS": 备份成功， "FAILED": 备份失败， "RUNNING": 备份进行中。
	Status *string `json:"Status,omitempty" name:"Status"`

	// binlog异地备份详细信息
	RemoteInfo []*RemoteBackupInfo `json:"RemoteInfo,omitempty" name:"RemoteInfo"`

	// 存储方式，0-常规存储，1-归档存储，默认为0
	CosStorageType *int64 `json:"CosStorageType,omitempty" name:"CosStorageType"`

	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type CdbRegionSellConf struct {
	// 地域中文名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 所属大区
	Area *string `json:"Area,omitempty" name:"Area"`

	// 是否为默认地域
	IsDefaultRegion *int64 `json:"IsDefaultRegion,omitempty" name:"IsDefaultRegion"`

	// 地域名称
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域的可用区售卖配置
	RegionConfig []*CdbZoneSellConf `json:"RegionConfig,omitempty" name:"RegionConfig"`
}

type CdbSellConfig struct {
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

	// 每秒IO数量
	Iops *int64 `json:"Iops,omitempty" name:"Iops"`

	// 应用场景描述
	Info *string `json:"Info,omitempty" name:"Info"`

	// 状态值，0 表示该规格对外售卖
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例类型，可能的取值范围有：UNIVERSAL (通用型), EXCLUSIVE (独享型), BASIC (基础型), BASIC_V2 (基础型v2)
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 引擎类型描述，可能的取值范围有：Innodb，RocksDB
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`

	// 售卖规格Id
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type CdbSellType struct {
	// 售卖实例名称。Z3是高可用类型对应规格中的DeviceType包含UNIVERSAL,EXCLUSIVE；CVM是基础版类型对应规格中的DeviceType是BASIC；TKE是基础型v2类型对应规格中的DeviceType是BASIC_V2。
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// 引擎版本号
	EngineVersion []*string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 售卖规格Id
	ConfigIds []*int64 `json:"ConfigIds,omitempty" name:"ConfigIds"`
}

type CdbZoneDataResult struct {
	// 售卖规格所有集合
	Configs []*CdbSellConfig `json:"Configs,omitempty" name:"Configs"`

	// 售卖地域可用区集合
	Regions []*CdbRegionSellConf `json:"Regions,omitempty" name:"Regions"`
}

type CdbZoneSellConf struct {
	// 可用区状态。可能的返回值为：1-上线；3-停售；4-不展示
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
	PayType []*string `json:"PayType,omitempty" name:"PayType"`

	// 数据复制类型。0-异步复制；1-半同步复制；2-强同步复制
	ProtectMode []*string `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// 可用区名称
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 多可用区信息
	ZoneConf *ZoneConf `json:"ZoneConf,omitempty" name:"ZoneConf"`

	// 可支持的灾备可用区信息
	DrZone []*string `json:"DrZone,omitempty" name:"DrZone"`

	// 是否支持跨可用区只读
	IsSupportRemoteRo *bool `json:"IsSupportRemoteRo,omitempty" name:"IsSupportRemoteRo"`

	// 可支持的跨可用区只读区信息
	RemoteRoZone []*string `json:"RemoteRoZone,omitempty" name:"RemoteRoZone"`

	// 独享型可用区状态。可能的返回值为：1-上线；3-停售；4-不展示
	ExClusterStatus *int64 `json:"ExClusterStatus,omitempty" name:"ExClusterStatus"`

	// 独享型可支持的跨可用区只读区信息
	ExClusterRemoteRoZone []*string `json:"ExClusterRemoteRoZone,omitempty" name:"ExClusterRemoteRoZone"`

	// 独享型多可用区信息
	ExClusterZoneConf *ZoneConf `json:"ExClusterZoneConf,omitempty" name:"ExClusterZoneConf"`

	// 售卖实例类型数组，其中configIds的值与configs结构体中的id一一对应。
	SellType []*CdbSellType `json:"SellType,omitempty" name:"SellType"`

	// 可用区id
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 是否支持ipv6
	IsSupportIpv6 *bool `json:"IsSupportIpv6,omitempty" name:"IsSupportIpv6"`

	// 可支持的售卖数据库引擎类型
	EngineType []*string `json:"EngineType,omitempty" name:"EngineType"`
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

	// 克隆实例所在地域Id
	NewRegionId *int64 `json:"NewRegionId,omitempty" name:"NewRegionId"`

	// 源实例所在地域Id
	SrcRegionId *int64 `json:"SrcRegionId,omitempty" name:"SrcRegionId"`
}

// Predefined struct for user
type CloseCDBProxyRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 是否只关闭读写分离，取值："true" | "false"，默认为"false"
	OnlyCloseRW *bool `json:"OnlyCloseRW,omitempty" name:"OnlyCloseRW"`
}

type CloseCDBProxyRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 是否只关闭读写分离，取值："true" | "false"，默认为"false"
	OnlyCloseRW *bool `json:"OnlyCloseRW,omitempty" name:"OnlyCloseRW"`
}

func (r *CloseCDBProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseCDBProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProxyGroupId")
	delete(f, "OnlyCloseRW")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseCDBProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseCDBProxyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CloseCDBProxyResponse struct {
	*tchttp.BaseResponse
	Response *CloseCDBProxyResponseParams `json:"Response"`
}

func (r *CloseCDBProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseCDBProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseWanServiceRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseWanServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseWanServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseWanServiceResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CloseWanServiceResponse struct {
	*tchttp.BaseResponse
	Response *CloseWanServiceResponseParams `json:"Response"`
}

func (r *CloseWanServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`
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

	// 常规备份保留策略，weekly-按周备份，monthly-按月备份，默认为weekly
	BackupPeriodStrategy *string `json:"BackupPeriodStrategy,omitempty" name:"BackupPeriodStrategy"`

	// 如果设置为按月备份，需填入每月具体备份日期，相邻备份天数不得超过两天。例[1,4,7,9,11,14,17,19,22,25,28,30,31]
	Days []*int64 `json:"Days,omitempty" name:"Days"`

	// 月度备份时间窗，BackupPeriodStrategy为monthly时必填。格式如： 02:00-06:00
	BackupPeriodTime *string `json:"BackupPeriodTime,omitempty" name:"BackupPeriodTime"`
}

type ConnectionPoolInfo struct {
	// 是否开启了连接池
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectionPool *bool `json:"ConnectionPool,omitempty" name:"ConnectionPool"`

	// 连接池类型：SessionConnectionPool（会话级别连接池）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectionPoolType *string `json:"ConnectionPoolType,omitempty" name:"ConnectionPoolType"`

	// 连接池保持阈值：单位（秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoolConnectionTimeOut *int64 `json:"PoolConnectionTimeOut,omitempty" name:"PoolConnectionTimeOut"`
}

// Predefined struct for user
type CreateAccountsRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 云数据库账号。
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`

	// 新账户的密码。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 备注信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 新账户最大可用连接数，默认值为10240，最大可设置值为10240。
	MaxUserConnections *int64 `json:"MaxUserConnections,omitempty" name:"MaxUserConnections"`
}

type CreateAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 云数据库账号。
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`

	// 新账户的密码。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 备注信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 新账户最大可用连接数，默认值为10240，最大可设置值为10240。
	MaxUserConnections *int64 `json:"MaxUserConnections,omitempty" name:"MaxUserConnections"`
}

func (r *CreateAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	delete(f, "Password")
	delete(f, "Description")
	delete(f, "MaxUserConnections")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountsResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAccountsResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccountsResponseParams `json:"Response"`
}

func (r *CreateAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditLogFileRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditLogFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Order")
	delete(f, "OrderBy")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditLogFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditLogFileResponseParams struct {
	// 审计日志文件名称。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAuditLogFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuditLogFileResponseParams `json:"Response"`
}

func (r *CreateAuditLogFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditLogFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditPolicyRequestParams struct {
	// 审计策略名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 审计规则 ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 审计日志保存时长。支持值包括：
	// 7 - 一周
	// 30 - 一个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年；
	// 实例首次开通审计策略时，可传该值，用于设置存储日志保存天数，默认为 30 天。若实例已存在审计策略，则此参数无效，可使用 更改审计服务配置 接口修改日志存储时长。
	LogExpireDay *int64 `json:"LogExpireDay,omitempty" name:"LogExpireDay"`
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
	// 7 - 一周
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "RuleId")
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditPolicyResponseParams struct {
	// 审计策略 ID。
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAuditPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuditPolicyResponseParams `json:"Response"`
}

func (r *CreateAuditPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditRuleRequestParams struct {
	// 审计规则名称。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 审计规则描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 审计规则过滤条件。若设置了过滤条件，则不会开启全审计。
	RuleFilters []*AuditFilter `json:"RuleFilters,omitempty" name:"RuleFilters"`

	// 是否开启全审计。支持值包括：false – 不开启全审计，true – 开启全审计。用户未设置审计规则过滤条件时，默认开启全审计。
	AuditAll *bool `json:"AuditAll,omitempty" name:"AuditAll"`
}

type CreateAuditRuleRequest struct {
	*tchttp.BaseRequest
	
	// 审计规则名称。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 审计规则描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 审计规则过滤条件。若设置了过滤条件，则不会开启全审计。
	RuleFilters []*AuditFilter `json:"RuleFilters,omitempty" name:"RuleFilters"`

	// 是否开启全审计。支持值包括：false – 不开启全审计，true – 开启全审计。用户未设置审计规则过滤条件时，默认开启全审计。
	AuditAll *bool `json:"AuditAll,omitempty" name:"AuditAll"`
}

func (r *CreateAuditRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	delete(f, "Description")
	delete(f, "RuleFilters")
	delete(f, "AuditAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditRuleResponseParams struct {
	// 审计规则 ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAuditRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuditRuleResponseParams `json:"Response"`
}

func (r *CreateAuditRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 目标备份方法，可选的值：logical - 逻辑冷备，physical - 物理冷备，snapshot - 快照备份。基础版实例仅支持快照备份。
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// 需要备份的库表信息，如果不设置该参数，则默认整实例备份。在 BackupMethod=logical 逻辑备份中才可设置该参数。指定的库表必须存在，否则可能导致备份失败。
	// 例：如果需要备份 db1 库的 tb1、tb2 表 和 db2 库。则该参数设置为 [{"Db": "db1", "Table": "tb1"}, {"Db": "db1", "Table": "tb2"}, {"Db": "db2"} ]。
	BackupDBTableList []*BackupItem `json:"BackupDBTableList,omitempty" name:"BackupDBTableList"`

	// 手动备份别名
	ManualBackupName *string `json:"ManualBackupName,omitempty" name:"ManualBackupName"`
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 目标备份方法，可选的值：logical - 逻辑冷备，physical - 物理冷备，snapshot - 快照备份。基础版实例仅支持快照备份。
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// 需要备份的库表信息，如果不设置该参数，则默认整实例备份。在 BackupMethod=logical 逻辑备份中才可设置该参数。指定的库表必须存在，否则可能导致备份失败。
	// 例：如果需要备份 db1 库的 tb1、tb2 表 和 db2 库。则该参数设置为 [{"Db": "db1", "Table": "tb1"}, {"Db": "db1", "Table": "tb2"}, {"Db": "db2"} ]。
	BackupDBTableList []*BackupItem `json:"BackupDBTableList,omitempty" name:"BackupDBTableList"`

	// 手动备份别名
	ManualBackupName *string `json:"ManualBackupName,omitempty" name:"ManualBackupName"`
}

func (r *CreateBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMethod")
	delete(f, "BackupDBTableList")
	delete(f, "ManualBackupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupResponseParams struct {
	// 备份任务 ID。
	BackupId *uint64 `json:"BackupId,omitempty" name:"BackupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBackupResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupResponseParams `json:"Response"`
}

func (r *CreateBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloneInstanceRequestParams struct {
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
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// 实例标签信息。
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 实例Cpu核数，需要不低于克隆源实例，默认和源实例相同。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制。
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// 多可用区域，默认为 0，支持值包括：0 - 表示单可用区，1 - 表示多可用区。
	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// 新产生的克隆实例备库 1 的可用区信息，默认同源实例 Zone 的值。
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// 备库 2 的可用区信息，默认为空，克隆强同步主实例时可指定该参数。
	BackupZone *string `json:"BackupZone,omitempty" name:"BackupZone"`

	// 克隆实例类型。支持值包括： "UNIVERSAL" - 通用型实例， "EXCLUSIVE" - 独享型实例。 不指定则默认为通用型。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 新克隆实例节点数。如果需要克隆出三节点实例， 请将该值设置为3 或指定 BackupZone 参数。如果需要克隆出两节点实例，请将该值设置为2。默认克隆出两节点实例。
	InstanceNodes *int64 `json:"InstanceNodes,omitempty" name:"InstanceNodes"`

	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// 是否只预检此次请求。true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制等。如果检查不通过，则返回对应错误码；如果检查通过，则返回RequestId.默认为false：发送正常请求，通过检查后直接创建实例。
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 金融围拢 ID 。
	CageId *string `json:"CageId,omitempty" name:"CageId"`

	// 项目ID，默认项目ID0
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
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
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// 实例标签信息。
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 实例Cpu核数，需要不低于克隆源实例，默认和源实例相同。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制。
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// 多可用区域，默认为 0，支持值包括：0 - 表示单可用区，1 - 表示多可用区。
	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// 新产生的克隆实例备库 1 的可用区信息，默认同源实例 Zone 的值。
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// 备库 2 的可用区信息，默认为空，克隆强同步主实例时可指定该参数。
	BackupZone *string `json:"BackupZone,omitempty" name:"BackupZone"`

	// 克隆实例类型。支持值包括： "UNIVERSAL" - 通用型实例， "EXCLUSIVE" - 独享型实例。 不指定则默认为通用型。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 新克隆实例节点数。如果需要克隆出三节点实例， 请将该值设置为3 或指定 BackupZone 参数。如果需要克隆出两节点实例，请将该值设置为2。默认克隆出两节点实例。
	InstanceNodes *int64 `json:"InstanceNodes,omitempty" name:"InstanceNodes"`

	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// 是否只预检此次请求。true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制等。如果检查不通过，则返回对应错误码；如果检查通过，则返回RequestId.默认为false：发送正常请求，通过检查后直接创建实例。
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 金融围拢 ID 。
	CageId *string `json:"CageId,omitempty" name:"CageId"`

	// 项目ID，默认项目ID0
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
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
	delete(f, "InstanceId")
	delete(f, "SpecifiedRollbackTime")
	delete(f, "SpecifiedBackupId")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "InstanceName")
	delete(f, "SecurityGroup")
	delete(f, "ResourceTags")
	delete(f, "Cpu")
	delete(f, "ProtectMode")
	delete(f, "DeployMode")
	delete(f, "SlaveZone")
	delete(f, "BackupZone")
	delete(f, "DeviceType")
	delete(f, "InstanceNodes")
	delete(f, "DeployGroupId")
	delete(f, "DryRun")
	delete(f, "CageId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloneInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloneInstanceResponseParams struct {
	// 异步任务的请求ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateDBImportJobRequestParams struct {
	// 实例的 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 云数据库的用户名。
	User *string `json:"User,omitempty" name:"User"`

	// 文件名称。该文件是指用户已上传到腾讯云的文件，仅支持.sql文件。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 云数据库实例 User 账号的密码。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 导入的目标数据库名，不传表示不指定数据库。
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 腾讯云COS文件链接。 用户需要指定 FileName 或者 CosUrl 其中一个。 COS文件需要是 .sql 文件。
	CosUrl *string `json:"CosUrl,omitempty" name:"CosUrl"`
}

type CreateDBImportJobRequest struct {
	*tchttp.BaseRequest
	
	// 实例的 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 云数据库的用户名。
	User *string `json:"User,omitempty" name:"User"`

	// 文件名称。该文件是指用户已上传到腾讯云的文件，仅支持.sql文件。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 云数据库实例 User 账号的密码。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 导入的目标数据库名，不传表示不指定数据库。
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 腾讯云COS文件链接。 用户需要指定 FileName 或者 CosUrl 其中一个。 COS文件需要是 .sql 文件。
	CosUrl *string `json:"CosUrl,omitempty" name:"CosUrl"`
}

func (r *CreateDBImportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBImportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "FileName")
	delete(f, "Password")
	delete(f, "DbName")
	delete(f, "CosUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBImportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBImportJobResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDBImportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBImportJobResponseParams `json:"Response"`
}

func (r *CreateDBImportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBImportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceHourRequestParams struct {
	// 实例数量，默认值为 1，最小值 1，最大值为 100。
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 实例内存大小，单位：MB，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的内存规格。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的硬盘范围。
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// MySQL 版本，值包括：5.5、5.6 、5.7 、8.0，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的实例版本。
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 私有网络 ID，如果不传则默认选择基础网络，请使用 [查询私有网络列表](/document/api/215/15778) 。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 私有网络下的子网 ID，如果设置了 UniqVpcId，则 UniqSubnetId 必填，请使用[查询子网列表](/document/api/215/15784)。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 项目 ID，不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 可用区信息，该参数缺省时，系统会自动选择一个可用区，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例 ID，购买只读实例或者灾备实例时必填，该字段表示只读实例或者灾备实例的主实例 ID，请使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口查询云数据库实例 ID。
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" name:"MasterInstanceId"`

	// 实例类型，默认为 master，支持值包括：master - 表示主实例，dr - 表示灾备实例，ro - 表示只读实例。
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// 主实例的可用区信息，购买灾备、RO实例时必填。
	MasterRegion *string `json:"MasterRegion,omitempty" name:"MasterRegion"`

	// 自定义端口，端口支持范围：[ 1024-65535 ] 。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 设置 root 帐号密码，密码规则：8 - 64 个字符，至少包含字母、数字、字符（支持的字符：_+-&=!@#$%^*()）中的两种，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 参数列表，参数格式如 ParamList.0.Name=auto_increment&ParamList.0.Value=1。可通过 [查询默认的可设置参数列表](https://cloud.tencent.com/document/api/236/32662) 查询支持设置的参数。
	ParamList []*ParamInfo `json:"ParamList,omitempty" name:"ParamList"`

	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// 多可用区域，默认为 0，支持值包括：0 - 表示单可用区，1 - 表示多可用区，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// 备库 1 的可用区信息，默认为 Zone 的值，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// 备库 2 的可用区信息，默认为空，购买三节点主实例时可指定该参数。
	BackupZone *string `json:"BackupZone,omitempty" name:"BackupZone"`

	// 安全组参数，可使用 [查询项目安全组信息](https://cloud.tencent.com/document/api/236/15850) 接口查询某个项目的安全组详情。
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// 只读实例信息。购买只读实例时，该参数必传。
	RoGroup *RoGroup `json:"RoGroup,omitempty" name:"RoGroup"`

	// 购买按量计费实例该字段无意义。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 实例名称。一次购买多个实例命名会用后缀数字区分，例instnaceName=db，goodsNum=3，实例命名分别为db1，db2，db3。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例标签信息。
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间在48小时内唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 实例隔离类型。支持值包括： "UNIVERSAL" - 通用型实例， "EXCLUSIVE" - 独享型实例， "BASIC" - 基础版实例。 不指定则默认为通用型实例。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 参数模板id。
	ParamTemplateId *int64 `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`

	// 告警策略id数组。云监控DescribeAlarmPolicy接口返回的OriginId。
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitempty" name:"AlarmPolicyList"`

	// 实例节点数。对于 RO 和 基础版实例， 该值默认为1。 如果需要购买三节点实例， 请将该值设置为3 或指定 BackupZone 参数。当购买主实例，且未指定该参数和 BackupZone 参数时，该值默认是 2， 即购买两节点实例。
	InstanceNodes *int64 `json:"InstanceNodes,omitempty" name:"InstanceNodes"`

	// 实例cpu核数， 如果不传将根据memory指定的内存值自动填充对应的cpu值。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 是否自动发起灾备同步功能。该参数仅对购买灾备实例生效。 可选值为：0 - 不自动发起灾备同步；1 - 自动发起灾备同步。该值默认为0。
	AutoSyncFlag *int64 `json:"AutoSyncFlag,omitempty" name:"AutoSyncFlag"`

	// 金融围拢 ID 。
	CageId *string `json:"CageId,omitempty" name:"CageId"`

	// 默认参数模板类型。支持值包括："HIGH_STABILITY" - 高稳定模板，"HIGH_PERFORMANCE" - 高性能模板，默认值是："HIGH_STABILITY"。
	ParamTemplateType *string `json:"ParamTemplateType,omitempty" name:"ParamTemplateType"`

	// 告警策略名数组，例如:["policy-uyoee9wg"]，AlarmPolicyList不为空时该参数无效。
	AlarmPolicyIdList []*string `json:"AlarmPolicyIdList,omitempty" name:"AlarmPolicyIdList"`

	// 是否只预检此次请求。true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制等。如果检查不通过，则返回对应错误码；如果检查通过，则返回RequestId.默认为false：发送正常请求，通过检查后直接创建实例。
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 实例引擎类型，默认为"InnoDB"，支持值包括："InnoDB"，"RocksDB"。
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`

	// 指定实例的IP列表。仅支持主实例指定，按实例顺序，不足则按未指定处理。
	Vips []*string `json:"Vips,omitempty" name:"Vips"`
}

type CreateDBInstanceHourRequest struct {
	*tchttp.BaseRequest
	
	// 实例数量，默认值为 1，最小值 1，最大值为 100。
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 实例内存大小，单位：MB，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的内存规格。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的硬盘范围。
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// MySQL 版本，值包括：5.5、5.6 、5.7 、8.0，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的实例版本。
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 私有网络 ID，如果不传则默认选择基础网络，请使用 [查询私有网络列表](/document/api/215/15778) 。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 私有网络下的子网 ID，如果设置了 UniqVpcId，则 UniqSubnetId 必填，请使用[查询子网列表](/document/api/215/15784)。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 项目 ID，不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 可用区信息，该参数缺省时，系统会自动选择一个可用区，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例 ID，购买只读实例或者灾备实例时必填，该字段表示只读实例或者灾备实例的主实例 ID，请使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口查询云数据库实例 ID。
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" name:"MasterInstanceId"`

	// 实例类型，默认为 master，支持值包括：master - 表示主实例，dr - 表示灾备实例，ro - 表示只读实例。
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// 主实例的可用区信息，购买灾备、RO实例时必填。
	MasterRegion *string `json:"MasterRegion,omitempty" name:"MasterRegion"`

	// 自定义端口，端口支持范围：[ 1024-65535 ] 。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 设置 root 帐号密码，密码规则：8 - 64 个字符，至少包含字母、数字、字符（支持的字符：_+-&=!@#$%^*()）中的两种，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 参数列表，参数格式如 ParamList.0.Name=auto_increment&ParamList.0.Value=1。可通过 [查询默认的可设置参数列表](https://cloud.tencent.com/document/api/236/32662) 查询支持设置的参数。
	ParamList []*ParamInfo `json:"ParamList,omitempty" name:"ParamList"`

	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// 多可用区域，默认为 0，支持值包括：0 - 表示单可用区，1 - 表示多可用区，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// 备库 1 的可用区信息，默认为 Zone 的值，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// 备库 2 的可用区信息，默认为空，购买三节点主实例时可指定该参数。
	BackupZone *string `json:"BackupZone,omitempty" name:"BackupZone"`

	// 安全组参数，可使用 [查询项目安全组信息](https://cloud.tencent.com/document/api/236/15850) 接口查询某个项目的安全组详情。
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// 只读实例信息。购买只读实例时，该参数必传。
	RoGroup *RoGroup `json:"RoGroup,omitempty" name:"RoGroup"`

	// 购买按量计费实例该字段无意义。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 实例名称。一次购买多个实例命名会用后缀数字区分，例instnaceName=db，goodsNum=3，实例命名分别为db1，db2，db3。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例标签信息。
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间在48小时内唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 实例隔离类型。支持值包括： "UNIVERSAL" - 通用型实例， "EXCLUSIVE" - 独享型实例， "BASIC" - 基础版实例。 不指定则默认为通用型实例。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 参数模板id。
	ParamTemplateId *int64 `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`

	// 告警策略id数组。云监控DescribeAlarmPolicy接口返回的OriginId。
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitempty" name:"AlarmPolicyList"`

	// 实例节点数。对于 RO 和 基础版实例， 该值默认为1。 如果需要购买三节点实例， 请将该值设置为3 或指定 BackupZone 参数。当购买主实例，且未指定该参数和 BackupZone 参数时，该值默认是 2， 即购买两节点实例。
	InstanceNodes *int64 `json:"InstanceNodes,omitempty" name:"InstanceNodes"`

	// 实例cpu核数， 如果不传将根据memory指定的内存值自动填充对应的cpu值。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 是否自动发起灾备同步功能。该参数仅对购买灾备实例生效。 可选值为：0 - 不自动发起灾备同步；1 - 自动发起灾备同步。该值默认为0。
	AutoSyncFlag *int64 `json:"AutoSyncFlag,omitempty" name:"AutoSyncFlag"`

	// 金融围拢 ID 。
	CageId *string `json:"CageId,omitempty" name:"CageId"`

	// 默认参数模板类型。支持值包括："HIGH_STABILITY" - 高稳定模板，"HIGH_PERFORMANCE" - 高性能模板，默认值是："HIGH_STABILITY"。
	ParamTemplateType *string `json:"ParamTemplateType,omitempty" name:"ParamTemplateType"`

	// 告警策略名数组，例如:["policy-uyoee9wg"]，AlarmPolicyList不为空时该参数无效。
	AlarmPolicyIdList []*string `json:"AlarmPolicyIdList,omitempty" name:"AlarmPolicyIdList"`

	// 是否只预检此次请求。true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制等。如果检查不通过，则返回对应错误码；如果检查通过，则返回RequestId.默认为false：发送正常请求，通过检查后直接创建实例。
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 实例引擎类型，默认为"InnoDB"，支持值包括："InnoDB"，"RocksDB"。
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`

	// 指定实例的IP列表。仅支持主实例指定，按实例顺序，不足则按未指定处理。
	Vips []*string `json:"Vips,omitempty" name:"Vips"`
}

func (r *CreateDBInstanceHourRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceHourRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GoodsNum")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "EngineVersion")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "ProjectId")
	delete(f, "Zone")
	delete(f, "MasterInstanceId")
	delete(f, "InstanceRole")
	delete(f, "MasterRegion")
	delete(f, "Port")
	delete(f, "Password")
	delete(f, "ParamList")
	delete(f, "ProtectMode")
	delete(f, "DeployMode")
	delete(f, "SlaveZone")
	delete(f, "BackupZone")
	delete(f, "SecurityGroup")
	delete(f, "RoGroup")
	delete(f, "AutoRenewFlag")
	delete(f, "InstanceName")
	delete(f, "ResourceTags")
	delete(f, "DeployGroupId")
	delete(f, "ClientToken")
	delete(f, "DeviceType")
	delete(f, "ParamTemplateId")
	delete(f, "AlarmPolicyList")
	delete(f, "InstanceNodes")
	delete(f, "Cpu")
	delete(f, "AutoSyncFlag")
	delete(f, "CageId")
	delete(f, "ParamTemplateType")
	delete(f, "AlarmPolicyIdList")
	delete(f, "DryRun")
	delete(f, "EngineType")
	delete(f, "Vips")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceHourRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceHourResponseParams struct {
	// 短订单 ID。
	DealIds []*string `json:"DealIds,omitempty" name:"DealIds"`

	// 实例 ID 列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDBInstanceHourResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstanceHourResponseParams `json:"Response"`
}

func (r *CreateDBInstanceHourResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceHourResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceRequestParams struct {
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

	// 项目 ID，不填为默认项目。购买只读实例和灾备实例时，项目 ID 默认和主实例保持一致。
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
	ParamList []*ParamInfo `json:"ParamList,omitempty" name:"ParamList"`

	// 备库 2 的可用区信息，默认为空，购买三节点主实例时可指定该参数。
	BackupZone *string `json:"BackupZone,omitempty" name:"BackupZone"`

	// 自动续费标记，可选值为：0 - 不自动续费；1 - 自动续费。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 主实例地域信息，购买灾备、RO实例时，该字段必填。
	MasterRegion *string `json:"MasterRegion,omitempty" name:"MasterRegion"`

	// 安全组参数，可使用 [查询项目安全组信息](https://cloud.tencent.com/document/api/236/15850) 接口查询某个项目的安全组详情。
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// 只读实例参数。购买只读实例时，该参数必传。
	RoGroup *RoGroup `json:"RoGroup,omitempty" name:"RoGroup"`

	// 实例名称。一次购买多个实例命名会用后缀数字区分，例instnaceName=db，goodsNum=3，实例命名分别为db1，db2，db3。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例标签信息。
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间在48小时内唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 实例隔离类型。支持值包括： "UNIVERSAL" - 通用型实例， "EXCLUSIVE" - 独享型实例， "BASIC" - 基础版实例。 不指定则默认为通用型实例。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 参数模板id。
	ParamTemplateId *int64 `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`

	// 告警策略id数组。云监控DescribeAlarmPolicy接口返回的OriginId。
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitempty" name:"AlarmPolicyList"`

	// 实例节点数。对于 RO 和 基础版实例， 该值默认为1。 如果需要购买三节点实例， 请将该值设置为3 或指定 BackupZone 参数。当购买主实例，且未指定该参数和 BackupZone 参数时，该值默认是 2， 即购买两节点实例。
	InstanceNodes *int64 `json:"InstanceNodes,omitempty" name:"InstanceNodes"`

	// 实例cpu核数， 如果不传将根据memory指定的内存值自动填充对应的cpu值。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 是否自动发起灾备同步功能。该参数仅对购买灾备实例生效。 可选值为：0 - 不自动发起灾备同步；1 - 自动发起灾备同步。该值默认为0。
	AutoSyncFlag *int64 `json:"AutoSyncFlag,omitempty" name:"AutoSyncFlag"`

	// 金融围拢 ID。
	CageId *string `json:"CageId,omitempty" name:"CageId"`

	// 默认参数模板类型。支持值包括："HIGH_STABILITY" - 高稳定模板，"HIGH_PERFORMANCE" - 高性能模板。
	ParamTemplateType *string `json:"ParamTemplateType,omitempty" name:"ParamTemplateType"`

	// 告警策略名数组，例如:["policy-uyoee9wg"]，AlarmPolicyList不为空时该参数无效。
	AlarmPolicyIdList []*string `json:"AlarmPolicyIdList,omitempty" name:"AlarmPolicyIdList"`

	// 是否只预检此次请求。true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制等。如果检查不通过，则返回对应错误码；如果检查通过，则返回RequestId.默认为false：发送正常请求，通过检查后直接创建实例。
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 实例引擎类型，默认为"InnoDB"，支持值包括："InnoDB"，"RocksDB"。
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`

	// 指定实例的IP列表。仅支持主实例指定，按实例顺序，不足则按未指定处理。
	Vips []*string `json:"Vips,omitempty" name:"Vips"`
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

	// 项目 ID，不填为默认项目。购买只读实例和灾备实例时，项目 ID 默认和主实例保持一致。
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
	ParamList []*ParamInfo `json:"ParamList,omitempty" name:"ParamList"`

	// 备库 2 的可用区信息，默认为空，购买三节点主实例时可指定该参数。
	BackupZone *string `json:"BackupZone,omitempty" name:"BackupZone"`

	// 自动续费标记，可选值为：0 - 不自动续费；1 - 自动续费。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 主实例地域信息，购买灾备、RO实例时，该字段必填。
	MasterRegion *string `json:"MasterRegion,omitempty" name:"MasterRegion"`

	// 安全组参数，可使用 [查询项目安全组信息](https://cloud.tencent.com/document/api/236/15850) 接口查询某个项目的安全组详情。
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// 只读实例参数。购买只读实例时，该参数必传。
	RoGroup *RoGroup `json:"RoGroup,omitempty" name:"RoGroup"`

	// 实例名称。一次购买多个实例命名会用后缀数字区分，例instnaceName=db，goodsNum=3，实例命名分别为db1，db2，db3。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例标签信息。
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间在48小时内唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 实例隔离类型。支持值包括： "UNIVERSAL" - 通用型实例， "EXCLUSIVE" - 独享型实例， "BASIC" - 基础版实例。 不指定则默认为通用型实例。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 参数模板id。
	ParamTemplateId *int64 `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`

	// 告警策略id数组。云监控DescribeAlarmPolicy接口返回的OriginId。
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitempty" name:"AlarmPolicyList"`

	// 实例节点数。对于 RO 和 基础版实例， 该值默认为1。 如果需要购买三节点实例， 请将该值设置为3 或指定 BackupZone 参数。当购买主实例，且未指定该参数和 BackupZone 参数时，该值默认是 2， 即购买两节点实例。
	InstanceNodes *int64 `json:"InstanceNodes,omitempty" name:"InstanceNodes"`

	// 实例cpu核数， 如果不传将根据memory指定的内存值自动填充对应的cpu值。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 是否自动发起灾备同步功能。该参数仅对购买灾备实例生效。 可选值为：0 - 不自动发起灾备同步；1 - 自动发起灾备同步。该值默认为0。
	AutoSyncFlag *int64 `json:"AutoSyncFlag,omitempty" name:"AutoSyncFlag"`

	// 金融围拢 ID。
	CageId *string `json:"CageId,omitempty" name:"CageId"`

	// 默认参数模板类型。支持值包括："HIGH_STABILITY" - 高稳定模板，"HIGH_PERFORMANCE" - 高性能模板。
	ParamTemplateType *string `json:"ParamTemplateType,omitempty" name:"ParamTemplateType"`

	// 告警策略名数组，例如:["policy-uyoee9wg"]，AlarmPolicyList不为空时该参数无效。
	AlarmPolicyIdList []*string `json:"AlarmPolicyIdList,omitempty" name:"AlarmPolicyIdList"`

	// 是否只预检此次请求。true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制等。如果检查不通过，则返回对应错误码；如果检查通过，则返回RequestId.默认为false：发送正常请求，通过检查后直接创建实例。
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 实例引擎类型，默认为"InnoDB"，支持值包括："InnoDB"，"RocksDB"。
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`

	// 指定实例的IP列表。仅支持主实例指定，按实例顺序，不足则按未指定处理。
	Vips []*string `json:"Vips,omitempty" name:"Vips"`
}

func (r *CreateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "Period")
	delete(f, "GoodsNum")
	delete(f, "Zone")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "ProjectId")
	delete(f, "Port")
	delete(f, "InstanceRole")
	delete(f, "MasterInstanceId")
	delete(f, "EngineVersion")
	delete(f, "Password")
	delete(f, "ProtectMode")
	delete(f, "DeployMode")
	delete(f, "SlaveZone")
	delete(f, "ParamList")
	delete(f, "BackupZone")
	delete(f, "AutoRenewFlag")
	delete(f, "MasterRegion")
	delete(f, "SecurityGroup")
	delete(f, "RoGroup")
	delete(f, "InstanceName")
	delete(f, "ResourceTags")
	delete(f, "DeployGroupId")
	delete(f, "ClientToken")
	delete(f, "DeviceType")
	delete(f, "ParamTemplateId")
	delete(f, "AlarmPolicyList")
	delete(f, "InstanceNodes")
	delete(f, "Cpu")
	delete(f, "AutoSyncFlag")
	delete(f, "CageId")
	delete(f, "ParamTemplateType")
	delete(f, "AlarmPolicyIdList")
	delete(f, "DryRun")
	delete(f, "EngineType")
	delete(f, "Vips")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceResponseParams struct {
	// 短订单 ID。
	DealIds []*string `json:"DealIds,omitempty" name:"DealIds"`

	// 实例 ID 列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstanceResponseParams `json:"Response"`
}

func (r *CreateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeployGroupRequestParams struct {
	// 置放群组名称，最长不能超过60个字符。
	DeployGroupName *string `json:"DeployGroupName,omitempty" name:"DeployGroupName"`

	// 置放群组描述，最长不能超过200个字符。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 置放群组的亲和性策略，目前仅支持取值为1，策略1表示同台物理机上限制实例的个数。
	Affinity []*int64 `json:"Affinity,omitempty" name:"Affinity"`

	// 置放群组亲和性策略1中同台物理机上实例的限制个数。
	LimitNum *int64 `json:"LimitNum,omitempty" name:"LimitNum"`

	// 置放群组机型属性，可选参数：SH12+SH02、TS85。
	DevClass []*string `json:"DevClass,omitempty" name:"DevClass"`
}

type CreateDeployGroupRequest struct {
	*tchttp.BaseRequest
	
	// 置放群组名称，最长不能超过60个字符。
	DeployGroupName *string `json:"DeployGroupName,omitempty" name:"DeployGroupName"`

	// 置放群组描述，最长不能超过200个字符。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 置放群组的亲和性策略，目前仅支持取值为1，策略1表示同台物理机上限制实例的个数。
	Affinity []*int64 `json:"Affinity,omitempty" name:"Affinity"`

	// 置放群组亲和性策略1中同台物理机上实例的限制个数。
	LimitNum *int64 `json:"LimitNum,omitempty" name:"LimitNum"`

	// 置放群组机型属性，可选参数：SH12+SH02、TS85。
	DevClass []*string `json:"DevClass,omitempty" name:"DevClass"`
}

func (r *CreateDeployGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeployGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployGroupName")
	delete(f, "Description")
	delete(f, "Affinity")
	delete(f, "LimitNum")
	delete(f, "DevClass")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeployGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeployGroupResponseParams struct {
	// 置放群组ID。
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDeployGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateDeployGroupResponseParams `json:"Response"`
}

func (r *CreateDeployGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeployGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParamTemplateRequestParams struct {
	// 参数模板名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数模板描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// MySQL 版本号。
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 源参数模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 参数列表。
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList"`

	// 默认参数模板类型。支持值包括："HIGH_STABILITY" - 高稳定模板，"HIGH_PERFORMANCE" - 高性能模板。
	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`

	// 实例引擎类型，默认为"InnoDB"，支持值包括："InnoDB"，"RocksDB"。
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
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
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList"`

	// 默认参数模板类型。支持值包括："HIGH_STABILITY" - 高稳定模板，"HIGH_PERFORMANCE" - 高性能模板。
	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`

	// 实例引擎类型，默认为"InnoDB"，支持值包括："InnoDB"，"RocksDB"。
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
}

func (r *CreateParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "EngineVersion")
	delete(f, "TemplateId")
	delete(f, "ParamList")
	delete(f, "TemplateType")
	delete(f, "EngineType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParamTemplateResponseParams struct {
	// 参数模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateParamTemplateResponseParams `json:"Response"`
}

func (r *CreateParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoInstanceIpRequestParams struct {
	// 只读实例ID，格式如：cdbro-3i70uj0k，与云数据库控制台页面中显示的只读实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 子网描述符，例如：subnet-1typ0s7d。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// vpc描述符，例如：vpc-a23yt67j,如果传了该字段则UniqSubnetId必传
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
}

type CreateRoInstanceIpRequest struct {
	*tchttp.BaseRequest
	
	// 只读实例ID，格式如：cdbro-3i70uj0k，与云数据库控制台页面中显示的只读实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 子网描述符，例如：subnet-1typ0s7d。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// vpc描述符，例如：vpc-a23yt67j,如果传了该字段则UniqSubnetId必传
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
}

func (r *CreateRoInstanceIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoInstanceIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UniqSubnetId")
	delete(f, "UniqVpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoInstanceIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoInstanceIpResponseParams struct {
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
}

type CreateRoInstanceIpResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoInstanceIpResponseParams `json:"Response"`
}

func (r *CreateRoInstanceIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoInstanceIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomConfig struct {
	// 设备
	// 注意：此字段可能返回 null，表示取不到有效值。
	Device *string `json:"Device,omitempty" name:"Device"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 设备类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 内存
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`
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
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`

	// 数据库名
	Database *string `json:"Database,omitempty" name:"Database"`
}

type DatabasesWithCharacterLists struct {
	// 数据库名
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 字符集类型
	CharacterSet *string `json:"CharacterSet,omitempty" name:"CharacterSet"`
}

// Predefined struct for user
type DeleteAccountsRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 云数据库账号。
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`
}

type DeleteAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 云数据库账号。
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`
}

func (r *DeleteAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountsResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccountsResponseParams `json:"Response"`
}

func (r *DeleteAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditLogFileRequestParams struct {
	// 审计日志文件名称。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditLogFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileName")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuditLogFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditLogFileResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAuditLogFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuditLogFileResponseParams `json:"Response"`
}

func (r *DeleteAuditLogFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditLogFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditPolicyRequestParams struct {
	// 审计策略 ID。
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuditPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAuditPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuditPolicyResponseParams `json:"Response"`
}

func (r *DeleteAuditPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditRuleRequestParams struct {
	// 审计规则 ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuditRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAuditRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuditRuleResponseParams `json:"Response"`
}

func (r *DeleteAuditRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份任务 ID。该任务 ID 为 [创建云数据库备份](https://cloud.tencent.com/document/api/236/15844) 接口返回的任务 ID。
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteBackupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBackupResponseParams `json:"Response"`
}

func (r *DeleteBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeployGroupsRequestParams struct {
	// 要删除的置放群组 ID 列表。
	DeployGroupIds []*string `json:"DeployGroupIds,omitempty" name:"DeployGroupIds"`
}

type DeleteDeployGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的置放群组 ID 列表。
	DeployGroupIds []*string `json:"DeployGroupIds,omitempty" name:"DeployGroupIds"`
}

func (r *DeleteDeployGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeployGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeployGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeployGroupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDeployGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDeployGroupsResponseParams `json:"Response"`
}

func (r *DeleteDeployGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeployGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteParamTemplateRequestParams struct {
	// 参数模板ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteParamTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteParamTemplateResponseParams `json:"Response"`
}

func (r *DeleteParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTimeWindowRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTimeWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTimeWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTimeWindowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTimeWindowResponseParams `json:"Response"`
}

func (r *DeleteTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// Predefined struct for user
type DescribeAccountPrivilegesRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库的账号名称。
	User *string `json:"User,omitempty" name:"User"`

	// 数据库的账号域名。
	Host *string `json:"Host,omitempty" name:"Host"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountPrivilegesResponseParams struct {
	// 全局权限数组。
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitempty" name:"GlobalPrivileges"`

	// 数据库权限数组。
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitempty" name:"DatabasePrivileges"`

	// 数据库中的表权限数组。
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitempty" name:"TablePrivileges"`

	// 数据库表中的列权限数组。
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitempty" name:"ColumnPrivileges"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountPrivilegesResponseParams `json:"Response"`
}

func (r *DescribeAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 记录偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为20，最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 匹配账号名的正则表达式，规则同 MySQL 官网。
	AccountRegexp *string `json:"AccountRegexp,omitempty" name:"AccountRegexp"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AccountRegexp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsResponseParams struct {
	// 符合查询条件的账号数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 符合查询条件的账号详细信息。
	Items []*AccountInfo `json:"Items,omitempty" name:"Items"`

	// 用户可设置实例最大连接数。
	MaxUserConnections *int64 `json:"MaxUserConnections,omitempty" name:"MaxUserConnections"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountsResponseParams `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncRequestInfoRequestParams struct {
	// 异步任务的请求 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AsyncRequestId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAsyncRequestInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncRequestInfoResponseParams struct {
	// 任务执行结果。可能的取值：INITIAL - 初始化，RUNNING - 运行中，SUCCESS - 执行成功，FAILED - 执行失败，KILLED - 已终止，REMOVED - 已删除，PAUSED - 终止中。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务执行信息描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Info *string `json:"Info,omitempty" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAsyncRequestInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAsyncRequestInfoResponseParams `json:"Response"`
}

func (r *DescribeAsyncRequestInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditConfigRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditConfigResponseParams struct {
	// 审计日志保存时长。目前支持的值包括：[0，7，30，180，365，1095，1825]。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogExpireDay *int64 `json:"LogExpireDay,omitempty" name:"LogExpireDay"`

	// 审计日志存储类型。目前支持的值包括："storage" - 存储型。
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 是否正在关闭审计。目前支持的值包括："false"-否，"true"-是
	IsClosing *string `json:"IsClosing,omitempty" name:"IsClosing"`

	// 审计服务开通时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAuditConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditConfigResponseParams `json:"Response"`
}

func (r *DescribeAuditConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogFilesRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页大小参数。默认值为 20，最小值为 1，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 审计日志文件名。
	FileName *string `json:"FileName,omitempty" name:"FileName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "FileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditLogFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogFilesResponseParams struct {
	// 符合条件的审计日志文件个数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 审计日志文件详情。
	Items []*AuditLogFile `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAuditLogFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditLogFilesResponseParams `json:"Response"`
}

func (r *DescribeAuditLogFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditPoliciesRequestParams struct {
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

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
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

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *DescribeAuditPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PolicyId")
	delete(f, "PolicyName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "RuleId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditPoliciesResponseParams struct {
	// 符合条件的审计策略个数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 审计策略详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*AuditPolicy `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAuditPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditPoliciesResponseParams `json:"Response"`
}

func (r *DescribeAuditPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRulesRequestParams struct {
	// 审计规则 ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 审计规则名称。支持按审计规则名称进行模糊匹配查询。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 分页大小参数。默认值为 20，最小值为 1，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量。默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeAuditRulesRequest struct {
	*tchttp.BaseRequest
	
	// 审计规则 ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 审计规则名称。支持按审计规则名称进行模糊匹配查询。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 分页大小参数。默认值为 20，最小值为 1，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量。默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAuditRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "RuleName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRulesResponseParams struct {
	// 符合条件的审计规则个数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 审计规则详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*AuditRule `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAuditRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditRulesResponseParams `json:"Response"`
}

func (r *DescribeAuditRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupConfigRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupConfigResponseParams struct {
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

	// 定期保留开关，off - 不开启定期保留策略，on - 开启定期保留策略，默认为off
	EnableBackupPeriodSave *string `json:"EnableBackupPeriodSave,omitempty" name:"EnableBackupPeriodSave"`

	// 定期保留最长天数，最小值：90，最大值：3650，默认值：1080
	BackupPeriodSaveDays *int64 `json:"BackupPeriodSaveDays,omitempty" name:"BackupPeriodSaveDays"`

	// 定期保留策略周期，可取值：weekly - 周，monthly - 月， quarterly - 季度，yearly - 年，默认为monthly
	BackupPeriodSaveInterval *string `json:"BackupPeriodSaveInterval,omitempty" name:"BackupPeriodSaveInterval"`

	// 定期保留的备份数量，最小值为1，最大值不超过定期保留策略周期内常规备份个数，默认值为1
	BackupPeriodSaveCount *int64 `json:"BackupPeriodSaveCount,omitempty" name:"BackupPeriodSaveCount"`

	// 定期保留策略周期起始日期，格式：YYYY-MM-dd HH:mm:ss
	StartBackupPeriodSaveDate *string `json:"StartBackupPeriodSaveDate,omitempty" name:"StartBackupPeriodSaveDate"`

	// 是否开启数据备份归档策略，off-关闭，on-打开，默认为off
	EnableBackupArchive *string `json:"EnableBackupArchive,omitempty" name:"EnableBackupArchive"`

	// 数据备份归档起始天数，数据备份达到归档起始天数时进行归档，最小为180天，不得大于数据备份保留天数
	BackupArchiveDays *int64 `json:"BackupArchiveDays,omitempty" name:"BackupArchiveDays"`

	// 是否开启日志备份归档策略，off-关闭，on-打开，默认为off
	EnableBinlogArchive *string `json:"EnableBinlogArchive,omitempty" name:"EnableBinlogArchive"`

	// 日志备份归档起始天数，日志备份达到归档起始天数时进行归档，最小为180天，不得大于日志备份保留天数
	BinlogArchiveDays *int64 `json:"BinlogArchiveDays,omitempty" name:"BinlogArchiveDays"`

	// 是否开启数据备份标准存储策略，off-关闭，on-打开，默认为off
	EnableBackupStandby *string `json:"EnableBackupStandby,omitempty" name:"EnableBackupStandby"`

	// 数据备份标准存储起始天数，数据备份达到标准存储起始天数时进行转换，最小为30天，不得大于数据备份保留天数。如果开启备份归档，不得大于等于备份归档天数
	BackupStandbyDays *int64 `json:"BackupStandbyDays,omitempty" name:"BackupStandbyDays"`

	// 是否开启日志备份标准存储策略，off-关闭，on-打开，默认为off
	EnableBinlogStandby *string `json:"EnableBinlogStandby,omitempty" name:"EnableBinlogStandby"`

	// 日志备份标准存储起始天数，日志备份达到标准存储起始天数时进行转换，最小为30天，不得大于日志备份保留天数。如果开启备份归档，不得大于等于备份归档天数
	BinlogStandbyDays *int64 `json:"BinlogStandbyDays,omitempty" name:"BinlogStandbyDays"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupConfigResponseParams `json:"Response"`
}

func (r *DescribeBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDatabasesRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "SearchDatabase")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDatabasesResponseParams struct {
	// 返回的数据个数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 符合查询条件的数据库数组。
	Items []*DatabaseName `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDatabasesResponseParams `json:"Response"`
}

func (r *DescribeBackupDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadRestrictionRequestParams struct {

}

type DescribeBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBackupDownloadRestrictionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadRestrictionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadRestrictionResponseParams struct {
	// NoLimit 不限制,内外网都可以下载； LimitOnlyIntranet 仅内网可下载； Customize 用户自定义vpc:ip可下载。 只有该值为 Customize 时， LimitVpc 和 LimitIp 才有意义。
	LimitType *string `json:"LimitType,omitempty" name:"LimitType"`

	// 该参数仅支持 In， 表示 LimitVpc 指定的vpc可以下载。
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitempty" name:"VpcComparisonSymbol"`

	// In: 指定的ip可以下载； NotIn: 指定的ip不可以下载。
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitempty" name:"IpComparisonSymbol"`

	// 限制下载的vpc设置。
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitempty" name:"LimitVpc"`

	// 限制下载的ip设置。
	LimitIp []*string `json:"LimitIp,omitempty" name:"LimitIp"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupDownloadRestrictionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDownloadRestrictionResponseParams `json:"Response"`
}

func (r *DescribeBackupDownloadRestrictionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadRestrictionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupEncryptionStatusRequestParams struct {
	// 实例ID，格式如：cdb-XXXX。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeBackupEncryptionStatusRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cdb-XXXX。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeBackupEncryptionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupEncryptionStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupEncryptionStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupEncryptionStatusResponseParams struct {
	// 实例是否开启了物理备份加密。可能的值有 on, off 。
	EncryptionStatus *string `json:"EncryptionStatus,omitempty" name:"EncryptionStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupEncryptionStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupEncryptionStatusResponseParams `json:"Response"`
}

func (r *DescribeBackupEncryptionStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupEncryptionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupOverviewRequestParams struct {
	// 需要查询的云数据库产品类型，目前仅支持 "mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupOverviewResponseParams struct {
	// 用户在当前地域备份的总个数（包含数据备份和日志备份）。
	BackupCount *int64 `json:"BackupCount,omitempty" name:"BackupCount"`

	// 用户在当前地域备份的总容量
	BackupVolume *int64 `json:"BackupVolume,omitempty" name:"BackupVolume"`

	// 用户在当前地域备份的计费容量，即超出赠送容量的部分。
	BillingVolume *int64 `json:"BillingVolume,omitempty" name:"BillingVolume"`

	// 用户在当前地域获得的赠送备份容量。
	FreeVolume *int64 `json:"FreeVolume,omitempty" name:"FreeVolume"`

	// 用户在当前地域的异地备份总容量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemoteBackupVolume *int64 `json:"RemoteBackupVolume,omitempty" name:"RemoteBackupVolume"`

	// 归档备份容量，包含数据备份以及日志备份。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupArchiveVolume *int64 `json:"BackupArchiveVolume,omitempty" name:"BackupArchiveVolume"`

	// 标准存储备份容量，包含数据备份以及日志备份。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupStandbyVolume *int64 `json:"BackupStandbyVolume,omitempty" name:"BackupStandbyVolume"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupOverviewResponseParams `json:"Response"`
}

func (r *DescribeBackupOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupSummariesRequestParams struct {
	// 需要查询的云数据库产品类型，目前仅支持 "mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 分页查询数据的偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询数据的条目限制，默认值为20。最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 指定按某一项排序，可选值包括： BackupVolume: 备份容量， DataBackupVolume: 数据备份容量， BinlogBackupVolume: 日志备份容量， AutoBackupVolume: 自动备份容量， ManualBackupVolume: 手动备份容量。默认按照BackupVolume排序。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 指定排序方向，可选值包括： ASC: 正序， DESC: 逆序。默认值为 ASC。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeBackupSummariesRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的云数据库产品类型，目前仅支持 "mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 分页查询数据的偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询数据的条目限制，默认值为20。最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 指定按某一项排序，可选值包括： BackupVolume: 备份容量， DataBackupVolume: 数据备份容量， BinlogBackupVolume: 日志备份容量， AutoBackupVolume: 自动备份容量， ManualBackupVolume: 手动备份容量。默认按照BackupVolume排序。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 指定排序方向，可选值包括： ASC: 正序， DESC: 逆序。默认值为 ASC。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeBackupSummariesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupSummariesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupSummariesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupSummariesResponseParams struct {
	// 实例备份统计条目。
	Items []*BackupSummaryItem `json:"Items,omitempty" name:"Items"`

	// 实例备份统计总条目数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupSummariesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupSummariesResponseParams `json:"Response"`
}

func (r *DescribeBackupSummariesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupSummariesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupTablesRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "DatabaseName")
	delete(f, "SearchTable")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupTablesResponseParams struct {
	// 返回的数据个数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 符合条件的数据表数组。
	Items []*TableName `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupTablesResponseParams `json:"Response"`
}

func (r *DescribeBackupTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupsRequestParams struct {
	// 实例ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量，最小值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认值为20，最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupsResponseParams struct {
	// 符合查询条件的实例总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 符合查询条件的备份信息详情。
	Items []*BackupInfo `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupsResponseParams `json:"Response"`
}

func (r *DescribeBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogBackupOverviewRequestParams struct {
	// 需要查询的云数据库产品类型，目前仅支持 "mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogBackupOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogBackupOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogBackupOverviewResponseParams struct {
	// 总的日志备份容量，包含异地日志备份（单位为字节）。
	BinlogBackupVolume *int64 `json:"BinlogBackupVolume,omitempty" name:"BinlogBackupVolume"`

	// 总的日志备份个数，包含异地日志备份。
	BinlogBackupCount *int64 `json:"BinlogBackupCount,omitempty" name:"BinlogBackupCount"`

	// 异地日志备份容量（单位为字节）。
	RemoteBinlogVolume *int64 `json:"RemoteBinlogVolume,omitempty" name:"RemoteBinlogVolume"`

	// 异地日志备份个数。
	RemoteBinlogCount *int64 `json:"RemoteBinlogCount,omitempty" name:"RemoteBinlogCount"`

	// 归档日志备份容量（单位为字节）。
	BinlogArchiveVolume *int64 `json:"BinlogArchiveVolume,omitempty" name:"BinlogArchiveVolume"`

	// 归档日志备份个数。
	BinlogArchiveCount *int64 `json:"BinlogArchiveCount,omitempty" name:"BinlogArchiveCount"`

	// 标准存储日志备份容量（单位为字节）。
	BinlogStandbyVolume *int64 `json:"BinlogStandbyVolume,omitempty" name:"BinlogStandbyVolume"`

	// 标准存储日志备份个数。
	BinlogStandbyCount *int64 `json:"BinlogStandbyCount,omitempty" name:"BinlogStandbyCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBinlogBackupOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBinlogBackupOverviewResponseParams `json:"Response"`
}

func (r *DescribeBinlogBackupOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogBackupOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogsRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量，最小值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认值为20，最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogsResponseParams struct {
	// 符合查询条件的日志文件总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 符合查询条件的二进制日志文件详情。
	Items []*BinlogInfo `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBinlogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBinlogsResponseParams `json:"Response"`
}

func (r *DescribeBinlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCDBProxyRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`
}

type DescribeCDBProxyRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`
}

func (r *DescribeCDBProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCDBProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProxyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCDBProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCDBProxyResponseParams struct {
	// 代理组基本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaseGroup *BaseGroupInfo `json:"BaseGroup,omitempty" name:"BaseGroup"`

	// 代理组地址信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *Address `json:"Address,omitempty" name:"Address"`

	// 代理组节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyNode *ProxyNodeInfo `json:"ProxyNode,omitempty" name:"ProxyNode"`

	// 读写分析信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RWInstInfo *RWInfo `json:"RWInstInfo,omitempty" name:"RWInstInfo"`

	// 连接池信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectionPoolInfo *ConnectionPoolInfo `json:"ConnectionPoolInfo,omitempty" name:"ConnectionPoolInfo"`

	// 代理数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 代理信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyGroup []*ProxyGroup `json:"ProxyGroup,omitempty" name:"ProxyGroup"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCDBProxyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCDBProxyResponseParams `json:"Response"`
}

func (r *DescribeCDBProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCDBProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdbZoneConfigRequestParams struct {

}

type DescribeCdbZoneConfigRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCdbZoneConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdbZoneConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCdbZoneConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdbZoneConfigResponseParams struct {
	// 售卖规格和地域信息集合
	DataResult *CdbZoneDataResult `json:"DataResult,omitempty" name:"DataResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCdbZoneConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCdbZoneConfigResponseParams `json:"Response"`
}

func (r *DescribeCdbZoneConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdbZoneConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloneListRequestParams struct {
	// 查询指定源实例的克隆任务列表。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页查询时的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询时的每页条目数，默认值为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeCloneListRequest struct {
	*tchttp.BaseRequest
	
	// 查询指定源实例的克隆任务列表。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页查询时的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询时的每页条目数，默认值为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCloneListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloneListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloneListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloneListResponseParams struct {
	// 满足条件的条目数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 克隆任务列表。
	Items []*CloneItem `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloneListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloneListResponseParams `json:"Response"`
}

func (r *DescribeCloneListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloneListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBFeaturesRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeDBFeaturesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBFeaturesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBFeaturesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBFeaturesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBFeaturesResponseParams struct {
	// 是否支持数据库审计功能。
	IsSupportAudit *bool `json:"IsSupportAudit,omitempty" name:"IsSupportAudit"`

	// 开启审计是否需要升级内核版本。
	AuditNeedUpgrade *bool `json:"AuditNeedUpgrade,omitempty" name:"AuditNeedUpgrade"`

	// 是否支持数据库加密功能。
	IsSupportEncryption *bool `json:"IsSupportEncryption,omitempty" name:"IsSupportEncryption"`

	// 开启加密是否需要升级内核版本。
	EncryptionNeedUpgrade *bool `json:"EncryptionNeedUpgrade,omitempty" name:"EncryptionNeedUpgrade"`

	// 是否为异地只读实例。
	IsRemoteRo *bool `json:"IsRemoteRo,omitempty" name:"IsRemoteRo"`

	// 主实例所在地域。
	MasterRegion *string `json:"MasterRegion,omitempty" name:"MasterRegion"`

	// 是否支持小版本升级。
	IsSupportUpdateSubVersion *bool `json:"IsSupportUpdateSubVersion,omitempty" name:"IsSupportUpdateSubVersion"`

	// 当前内核版本。
	CurrentSubVersion *string `json:"CurrentSubVersion,omitempty" name:"CurrentSubVersion"`

	// 可供升级的内核版本。
	TargetSubVersion *string `json:"TargetSubVersion,omitempty" name:"TargetSubVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBFeaturesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBFeaturesResponseParams `json:"Response"`
}

func (r *DescribeDBFeaturesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBFeaturesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBImportRecordsRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBImportRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBImportRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBImportRecordsResponseParams struct {
	// 符合查询条件的导入任务操作日志总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回的导入操作记录列表。
	Items []*ImportRecord `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBImportRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBImportRecordsResponseParams `json:"Response"`
}

func (r *DescribeDBImportRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBImportRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceCharsetRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceCharsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceCharsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceCharsetResponseParams struct {
	// 实例的默认字符集，如 "latin1"，"utf8" 等。
	Charset *string `json:"Charset,omitempty" name:"Charset"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBInstanceCharsetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceCharsetResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceCharsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceCharsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceConfigRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceConfigResponseParams struct {
	// 主实例数据保护方式，可能的返回值：0 - 异步复制方式，1 - 半同步复制方式，2 - 强同步复制方式。
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// 主实例部署方式，可能的返回值：0 - 单可用部署，1 - 多可用区部署。
	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// 实例可用区信息，格式如 "ap-shanghai-1"。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 备库的配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaveConfig *SlaveConfig `json:"SlaveConfig,omitempty" name:"SlaveConfig"`

	// 强同步实例第二备库的配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupConfig *BackupConfig `json:"BackupConfig,omitempty" name:"BackupConfig"`

	// 是否切换备库。
	Switched *bool `json:"Switched,omitempty" name:"Switched"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBInstanceConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceConfigResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceGTIDRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceGTIDRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceGTIDRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceGTIDResponseParams struct {
	// GTID 是否开通的标记，可能的取值为：0 - 未开通，1 - 已开通。
	IsGTIDOpen *int64 `json:"IsGTIDOpen,omitempty" name:"IsGTIDOpen"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBInstanceGTIDResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceGTIDResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceGTIDResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceGTIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceInfoRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceInfoResponseParams struct {
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
}

type DescribeDBInstanceInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceInfoResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceRebootTimeRequestParams struct {
	// 实例的 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DescribeDBInstanceRebootTimeRequest struct {
	*tchttp.BaseRequest
	
	// 实例的 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeDBInstanceRebootTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceRebootTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceRebootTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceRebootTimeResponseParams struct {
	// 符合查询条件的实例总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回的参数信息。
	Items []*InstanceRebootTime `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBInstanceRebootTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceRebootTimeResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceRebootTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceRebootTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesRequestParams struct {
	// 项目 ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例类型，可取值：1 - 主实例，2 - 灾备实例，3 - 只读实例。
	InstanceTypes []*uint64 `json:"InstanceTypes,omitempty" name:"InstanceTypes"`

	// 实例的内网 IP 地址。
	Vips []*string `json:"Vips,omitempty" name:"Vips"`

	// 实例状态，可取值：<br>0 - 创建中<br>1 - 运行中<br>4 - 正在进行隔离操作<br>5 - 已隔离（可在回收站恢复开机）
	Status []*uint64 `json:"Status,omitempty" name:"Status"`

	// 偏移量，默认值为 0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为 20，最大值为 2000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 安全组 ID。当使用安全组 ID 为过滤条件时，需指定 WithSecurityGroup 参数为 1。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 付费类型，可取值：0 - 包年包月，1 - 小时计费。
	PayTypes []*uint64 `json:"PayTypes,omitempty" name:"PayTypes"`

	// 实例名称。
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`

	// 实例任务状态，可能取值：<br>0 - 没有任务<br>1 - 升级中<br>2 - 数据导入中<br>3 - 开放Slave中<br>4 - 外网访问开通中<br>5 - 批量操作执行中<br>6 - 回档中<br>7 - 外网访问关闭中<br>8 - 密码修改中<br>9 - 实例名修改中<br>10 - 重启中<br>12 - 自建迁移中<br>13 - 删除库表中<br>14 - 灾备实例创建同步中<br>15 - 升级待切换<br>16 - 升级切换中<br>17 - 升级切换完成<br>19 - 参数设置待执行
	TaskStatus []*uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 实例数据库引擎版本，可能取值：5.1、5.5、5.6 和 5.7。
	EngineVersions []*string `json:"EngineVersions,omitempty" name:"EngineVersions"`

	// 私有网络的 ID。
	VpcIds []*uint64 `json:"VpcIds,omitempty" name:"VpcIds"`

	// 可用区的 ID。
	ZoneIds []*uint64 `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 子网 ID。
	SubnetIds []*uint64 `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 是否锁定标记，可选值：0 - 不锁定，1 - 锁定，默认为0。
	CdbErrors []*int64 `json:"CdbErrors,omitempty" name:"CdbErrors"`

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
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 初始化标记，可取值：0 - 未初始化，1 - 初始化。
	InitFlag *int64 `json:"InitFlag,omitempty" name:"InitFlag"`

	// 是否包含灾备关系对应的实例，可取值：0 - 不包含，1 - 包含。默认取值为1。如果拉取主实例，则灾备关系的数据在DrInfo字段中， 如果拉取灾备实例， 则灾备关系的数据在MasterInfo字段中。灾备关系中只包含部分基本的数据，详细的数据需要自行调接口拉取。
	WithDr *int64 `json:"WithDr,omitempty" name:"WithDr"`

	// 是否包含只读实例，可取值：0 - 不包含，1 - 包含。默认取值为1。
	WithRo *int64 `json:"WithRo,omitempty" name:"WithRo"`

	// 是否包含主实例，可取值：0 - 不包含，1 - 包含。默认取值为1。
	WithMaster *int64 `json:"WithMaster,omitempty" name:"WithMaster"`

	// 置放群组ID列表。
	DeployGroupIds []*string `json:"DeployGroupIds,omitempty" name:"DeployGroupIds"`

	// 是否以标签键为过滤条件。
	TagKeysForSearch []*string `json:"TagKeysForSearch,omitempty" name:"TagKeysForSearch"`

	// 金融围拢 ID 。
	CageIds []*string `json:"CageIds,omitempty" name:"CageIds"`

	// 标签值
	TagValues []*string `json:"TagValues,omitempty" name:"TagValues"`

	// 私有网络字符型vpcId
	UniqueVpcIds []*string `json:"UniqueVpcIds,omitempty" name:"UniqueVpcIds"`

	// 私有网络字符型subnetId
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitempty" name:"UniqSubnetIds"`

	// 标签键值
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 数据库代理 IP 。
	ProxyVips []*string `json:"ProxyVips,omitempty" name:"ProxyVips"`

	// 数据库代理 ID 。
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 项目 ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例类型，可取值：1 - 主实例，2 - 灾备实例，3 - 只读实例。
	InstanceTypes []*uint64 `json:"InstanceTypes,omitempty" name:"InstanceTypes"`

	// 实例的内网 IP 地址。
	Vips []*string `json:"Vips,omitempty" name:"Vips"`

	// 实例状态，可取值：<br>0 - 创建中<br>1 - 运行中<br>4 - 正在进行隔离操作<br>5 - 已隔离（可在回收站恢复开机）
	Status []*uint64 `json:"Status,omitempty" name:"Status"`

	// 偏移量，默认值为 0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为 20，最大值为 2000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 安全组 ID。当使用安全组 ID 为过滤条件时，需指定 WithSecurityGroup 参数为 1。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 付费类型，可取值：0 - 包年包月，1 - 小时计费。
	PayTypes []*uint64 `json:"PayTypes,omitempty" name:"PayTypes"`

	// 实例名称。
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`

	// 实例任务状态，可能取值：<br>0 - 没有任务<br>1 - 升级中<br>2 - 数据导入中<br>3 - 开放Slave中<br>4 - 外网访问开通中<br>5 - 批量操作执行中<br>6 - 回档中<br>7 - 外网访问关闭中<br>8 - 密码修改中<br>9 - 实例名修改中<br>10 - 重启中<br>12 - 自建迁移中<br>13 - 删除库表中<br>14 - 灾备实例创建同步中<br>15 - 升级待切换<br>16 - 升级切换中<br>17 - 升级切换完成<br>19 - 参数设置待执行
	TaskStatus []*uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 实例数据库引擎版本，可能取值：5.1、5.5、5.6 和 5.7。
	EngineVersions []*string `json:"EngineVersions,omitempty" name:"EngineVersions"`

	// 私有网络的 ID。
	VpcIds []*uint64 `json:"VpcIds,omitempty" name:"VpcIds"`

	// 可用区的 ID。
	ZoneIds []*uint64 `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 子网 ID。
	SubnetIds []*uint64 `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 是否锁定标记，可选值：0 - 不锁定，1 - 锁定，默认为0。
	CdbErrors []*int64 `json:"CdbErrors,omitempty" name:"CdbErrors"`

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
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 初始化标记，可取值：0 - 未初始化，1 - 初始化。
	InitFlag *int64 `json:"InitFlag,omitempty" name:"InitFlag"`

	// 是否包含灾备关系对应的实例，可取值：0 - 不包含，1 - 包含。默认取值为1。如果拉取主实例，则灾备关系的数据在DrInfo字段中， 如果拉取灾备实例， 则灾备关系的数据在MasterInfo字段中。灾备关系中只包含部分基本的数据，详细的数据需要自行调接口拉取。
	WithDr *int64 `json:"WithDr,omitempty" name:"WithDr"`

	// 是否包含只读实例，可取值：0 - 不包含，1 - 包含。默认取值为1。
	WithRo *int64 `json:"WithRo,omitempty" name:"WithRo"`

	// 是否包含主实例，可取值：0 - 不包含，1 - 包含。默认取值为1。
	WithMaster *int64 `json:"WithMaster,omitempty" name:"WithMaster"`

	// 置放群组ID列表。
	DeployGroupIds []*string `json:"DeployGroupIds,omitempty" name:"DeployGroupIds"`

	// 是否以标签键为过滤条件。
	TagKeysForSearch []*string `json:"TagKeysForSearch,omitempty" name:"TagKeysForSearch"`

	// 金融围拢 ID 。
	CageIds []*string `json:"CageIds,omitempty" name:"CageIds"`

	// 标签值
	TagValues []*string `json:"TagValues,omitempty" name:"TagValues"`

	// 私有网络字符型vpcId
	UniqueVpcIds []*string `json:"UniqueVpcIds,omitempty" name:"UniqueVpcIds"`

	// 私有网络字符型subnetId
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitempty" name:"UniqSubnetIds"`

	// 标签键值
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 数据库代理 IP 。
	ProxyVips []*string `json:"ProxyVips,omitempty" name:"ProxyVips"`

	// 数据库代理 ID 。
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`
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
	delete(f, "ProjectId")
	delete(f, "InstanceTypes")
	delete(f, "Vips")
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SecurityGroupId")
	delete(f, "PayTypes")
	delete(f, "InstanceNames")
	delete(f, "TaskStatus")
	delete(f, "EngineVersions")
	delete(f, "VpcIds")
	delete(f, "ZoneIds")
	delete(f, "SubnetIds")
	delete(f, "CdbErrors")
	delete(f, "OrderBy")
	delete(f, "OrderDirection")
	delete(f, "WithSecurityGroup")
	delete(f, "WithExCluster")
	delete(f, "ExClusterId")
	delete(f, "InstanceIds")
	delete(f, "InitFlag")
	delete(f, "WithDr")
	delete(f, "WithRo")
	delete(f, "WithMaster")
	delete(f, "DeployGroupIds")
	delete(f, "TagKeysForSearch")
	delete(f, "CageIds")
	delete(f, "TagValues")
	delete(f, "UniqueVpcIds")
	delete(f, "UniqSubnetIds")
	delete(f, "Tags")
	delete(f, "ProxyVips")
	delete(f, "ProxyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesResponseParams struct {
	// 符合查询条件的实例总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例详细信息列表。
	Items []*InstanceInfo `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeDBPriceRequestParams struct {
	// 实例时长，单位：月，最小值 1，最大值为 36；查询按量计费价格时，该字段无效。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 可用区信息，格式如 "ap-guangzhou-2"。具体能设置的值请通过 <a href="https://cloud.tencent.com/document/api/236/17229">DescribeDBZoneConfig</a> 接口查询。InstanceId为空时该参数为必填项。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例数量，默认值为 1，最小值 1，最大值为 100。InstanceId为空时该参数为必填项。
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 实例内存大小，单位：MB。InstanceId为空时该参数为必填项。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB。InstanceId为空时该参数为必填项。
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// 实例类型，默认为 master，支持值包括：master - 表示主实例，ro - 表示只读实例，dr - 表示灾备实例。InstanceId为空时该参数为必填项。
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// 付费类型，支持值包括：PRE_PAID - 包年包月，HOUR_PAID - 按量计费。InstanceId为空时该参数为必填项。
	PayType *string `json:"PayType,omitempty" name:"PayType"`

	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制。
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// 实例隔离类型。支持值包括： "UNIVERSAL" - 通用型实例， "EXCLUSIVE" - 独享型实例， "BASIC" - 基础版实例。 不指定则默认为通用型实例。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 实例节点数。对于 RO 和 基础版实例， 该值默认为1。 如果需要询价三节点实例， 请将该值设置为3。其余主实例该值默认为2。
	InstanceNodes *int64 `json:"InstanceNodes,omitempty" name:"InstanceNodes"`

	// 询价实例的CPU核心数目，单位：核，为保证传入 CPU 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可售卖的核心数目，当未指定该值时，将按照 Memory 大小补全一个默认值。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 询价续费实例ID。如需查询实例续费价格，填写InstanceId和Period即可。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 按量计费阶梯。仅PayType=HOUR_PAID有效，支持值包括：1，2，3。阶梯时长见https://cloud.tencent.com/document/product/236/18335。
	Ladder *uint64 `json:"Ladder,omitempty" name:"Ladder"`
}

type DescribeDBPriceRequest struct {
	*tchttp.BaseRequest
	
	// 实例时长，单位：月，最小值 1，最大值为 36；查询按量计费价格时，该字段无效。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 可用区信息，格式如 "ap-guangzhou-2"。具体能设置的值请通过 <a href="https://cloud.tencent.com/document/api/236/17229">DescribeDBZoneConfig</a> 接口查询。InstanceId为空时该参数为必填项。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例数量，默认值为 1，最小值 1，最大值为 100。InstanceId为空时该参数为必填项。
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 实例内存大小，单位：MB。InstanceId为空时该参数为必填项。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB。InstanceId为空时该参数为必填项。
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// 实例类型，默认为 master，支持值包括：master - 表示主实例，ro - 表示只读实例，dr - 表示灾备实例。InstanceId为空时该参数为必填项。
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// 付费类型，支持值包括：PRE_PAID - 包年包月，HOUR_PAID - 按量计费。InstanceId为空时该参数为必填项。
	PayType *string `json:"PayType,omitempty" name:"PayType"`

	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制。
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// 实例隔离类型。支持值包括： "UNIVERSAL" - 通用型实例， "EXCLUSIVE" - 独享型实例， "BASIC" - 基础版实例。 不指定则默认为通用型实例。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 实例节点数。对于 RO 和 基础版实例， 该值默认为1。 如果需要询价三节点实例， 请将该值设置为3。其余主实例该值默认为2。
	InstanceNodes *int64 `json:"InstanceNodes,omitempty" name:"InstanceNodes"`

	// 询价实例的CPU核心数目，单位：核，为保证传入 CPU 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可售卖的核心数目，当未指定该值时，将按照 Memory 大小补全一个默认值。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 询价续费实例ID。如需查询实例续费价格，填写InstanceId和Period即可。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 按量计费阶梯。仅PayType=HOUR_PAID有效，支持值包括：1，2，3。阶梯时长见https://cloud.tencent.com/document/product/236/18335。
	Ladder *uint64 `json:"Ladder,omitempty" name:"Ladder"`
}

func (r *DescribeDBPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBPriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Period")
	delete(f, "Zone")
	delete(f, "GoodsNum")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "InstanceRole")
	delete(f, "PayType")
	delete(f, "ProtectMode")
	delete(f, "DeviceType")
	delete(f, "InstanceNodes")
	delete(f, "Cpu")
	delete(f, "InstanceId")
	delete(f, "Ladder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBPriceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBPriceResponseParams struct {
	// 实例价格，单位：分。
	Price *int64 `json:"Price,omitempty" name:"Price"`

	// 实例原价，单位：分。
	OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 货币单位。CNY-人民币，USD-美元。
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBPriceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBPriceResponseParams `json:"Response"`
}

func (r *DescribeDBPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsRequestParams struct {
	// 实例ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 该值默认为False，表示当传入只读实例ID时，查询操作的是对应只读组的安全组。如果需要操作只读实例ID的安全组， 需要将该入参置为True。
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitempty" name:"ForReadonlyInstance"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 该值默认为False，表示当传入只读实例ID时，查询操作的是对应只读组的安全组。如果需要操作只读实例ID的安全组， 需要将该入参置为True。
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitempty" name:"ForReadonlyInstance"`
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
	delete(f, "ForReadonlyInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsResponseParams struct {
	// 安全组详情。
	Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeDBSwitchRecordsRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认值为 50，最小值为 1，最大值为 2000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSwitchRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSwitchRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSwitchRecordsResponseParams struct {
	// 实例切换记录的总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例切换记录详情。
	Items []*DBSwitchInfo `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBSwitchRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSwitchRecordsResponseParams `json:"Response"`
}

func (r *DescribeDBSwitchRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSwitchRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataBackupOverviewRequestParams struct {
	// 需要查询的云数据库产品类型，目前仅支持 "mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataBackupOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataBackupOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataBackupOverviewResponseParams struct {
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

	// 异地备份总容量。
	RemoteBackupVolume *int64 `json:"RemoteBackupVolume,omitempty" name:"RemoteBackupVolume"`

	// 异地备份总个数。
	RemoteBackupCount *int64 `json:"RemoteBackupCount,omitempty" name:"RemoteBackupCount"`

	// 当前地域归档备份总容量。
	DataBackupArchiveVolume *int64 `json:"DataBackupArchiveVolume,omitempty" name:"DataBackupArchiveVolume"`

	// 当前地域归档备份总个数。
	DataBackupArchiveCount *int64 `json:"DataBackupArchiveCount,omitempty" name:"DataBackupArchiveCount"`

	// 当前地域标准存储备份总容量。
	DataBackupStandbyVolume *int64 `json:"DataBackupStandbyVolume,omitempty" name:"DataBackupStandbyVolume"`

	// 当前地域标准存储备份总个数。
	DataBackupStandbyCount *int64 `json:"DataBackupStandbyCount,omitempty" name:"DataBackupStandbyCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataBackupOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataBackupOverviewResponseParams `json:"Response"`
}

func (r *DescribeDataBackupOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataBackupOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量，最小值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单次请求数量，默认值为20，最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 匹配数据库库名的正则表达式。
	DatabaseRegexp *string `json:"DatabaseRegexp,omitempty" name:"DatabaseRegexp"`
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
	// 符合查询条件的实例总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回的实例信息。
	Items []*string `json:"Items,omitempty" name:"Items"`

	// 数据库名以及字符集
	DatabaseList []*DatabasesWithCharacterLists `json:"DatabaseList,omitempty" name:"DatabaseList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeDefaultParamsRequestParams struct {
	// mysql版本，目前支持 ["5.1", "5.5", "5.6", "5.7"]。
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 默认参数模板类型。支持值包括："HIGH_STABILITY" - 高稳定模板，"HIGH_PERFORMANCE" - 高性能模板。
	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
}

type DescribeDefaultParamsRequest struct {
	*tchttp.BaseRequest
	
	// mysql版本，目前支持 ["5.1", "5.5", "5.6", "5.7"]。
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 默认参数模板类型。支持值包括："HIGH_STABILITY" - 高稳定模板，"HIGH_PERFORMANCE" - 高性能模板。
	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
}

func (r *DescribeDefaultParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineVersion")
	delete(f, "TemplateType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDefaultParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultParamsResponseParams struct {
	// 参数个数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 参数详情。
	Items []*ParameterDetail `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDefaultParamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDefaultParamsResponseParams `json:"Response"`
}

func (r *DescribeDefaultParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeployGroupListRequestParams struct {
	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// 置放群组名称。
	DeployGroupName *string `json:"DeployGroupName,omitempty" name:"DeployGroupName"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeployGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployGroupId")
	delete(f, "DeployGroupName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeployGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeployGroupListResponseParams struct {
	// 符合条件的记录总数。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 返回列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DeployGroupInfo `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeployGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeployGroupListResponseParams `json:"Response"`
}

func (r *DescribeDeployGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeployGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceMonitorInfoRequestParams struct {
	// 实例ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 返回当天最近Count个5分钟粒度的监控数据。最小值1，最大值288，不传该参数默认返回当天所有5分钟粒度监控数据。
	Count *uint64 `json:"Count,omitempty" name:"Count"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceMonitorInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Count")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceMonitorInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceMonitorInfoResponseParams struct {
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
}

type DescribeDeviceMonitorInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceMonitorInfoResponseParams `json:"Response"`
}

func (r *DescribeDeviceMonitorInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceMonitorInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeErrorLogDataRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间戳。例如 1585142640 。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳。例如 1585142640 。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 要匹配的关键字列表，最多支持15个关键字。
	KeyWords []*string `json:"KeyWords,omitempty" name:"KeyWords"`

	// 分页的返回数量，默认为100，最大为400。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 仅在实例为主实例或者灾备实例时生效，可选值：slave，代表拉取从机的日志。
	InstType *string `json:"InstType,omitempty" name:"InstType"`
}

type DescribeErrorLogDataRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间戳。例如 1585142640 。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳。例如 1585142640 。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 要匹配的关键字列表，最多支持15个关键字。
	KeyWords []*string `json:"KeyWords,omitempty" name:"KeyWords"`

	// 分页的返回数量，默认为100，最大为400。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 仅在实例为主实例或者灾备实例时生效，可选值：slave，代表拉取从机的日志。
	InstType *string `json:"InstType,omitempty" name:"InstType"`
}

func (r *DescribeErrorLogDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeErrorLogDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "KeyWords")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InstType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeErrorLogDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeErrorLogDataResponseParams struct {
	// 符合条件的记录总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回的记录。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*ErrlogItem `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeErrorLogDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeErrorLogDataResponseParams `json:"Response"`
}

func (r *DescribeErrorLogDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeErrorLogDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamRecordsRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认值：20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeInstanceParamRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认值：20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstanceParamRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamRecordsResponseParams struct {
	// 符合条件的记录数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 参数修改记录。
	Items []*ParamRecord `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceParamRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceParamRecordsResponseParams `json:"Response"`
}

func (r *DescribeInstanceParamRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsResponseParams struct {
	// 实例的参数总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 参数详情。
	Items []*ParameterDetail `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceParamsResponseParams `json:"Response"`
}

func (r *DescribeInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLocalBinlogConfigRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeLocalBinlogConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeLocalBinlogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLocalBinlogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLocalBinlogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLocalBinlogConfigResponseParams struct {
	// 实例binlog保留策略。
	LocalBinlogConfig *LocalBinlogConfig `json:"LocalBinlogConfig,omitempty" name:"LocalBinlogConfig"`

	// 该地域默认binlog保留策略。
	LocalBinlogConfigDefault *LocalBinlogConfigDefault `json:"LocalBinlogConfigDefault,omitempty" name:"LocalBinlogConfigDefault"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLocalBinlogConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLocalBinlogConfigResponseParams `json:"Response"`
}

func (r *DescribeLocalBinlogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLocalBinlogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplateInfoRequestParams struct {
	// 参数模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplateInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplateInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplateInfoResponseParams struct {
	// 参数模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 参数模板名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数模板对应实例版本
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 参数模板中的参数数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 参数详情
	Items []*ParameterDetail `json:"Items,omitempty" name:"Items"`

	// 参数模板描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 参数模板类型。支持值包括："HIGH_STABILITY" - 高稳定模板，"HIGH_PERFORMANCE" - 高性能模板。
	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeParamTemplateInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParamTemplateInfoResponseParams `json:"Response"`
}

func (r *DescribeParamTemplateInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplateInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplatesRequestParams struct {
	// 引擎版本，缺省则查询所有
	EngineVersions []*string `json:"EngineVersions,omitempty" name:"EngineVersions"`

	// 引擎类型，缺省则查询所有
	EngineTypes []*string `json:"EngineTypes,omitempty" name:"EngineTypes"`

	// 模板名称，缺省则查询所有
	TemplateNames []*string `json:"TemplateNames,omitempty" name:"TemplateNames"`

	// 模板id，缺省则查询所有
	TemplateIds []*int64 `json:"TemplateIds,omitempty" name:"TemplateIds"`
}

type DescribeParamTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 引擎版本，缺省则查询所有
	EngineVersions []*string `json:"EngineVersions,omitempty" name:"EngineVersions"`

	// 引擎类型，缺省则查询所有
	EngineTypes []*string `json:"EngineTypes,omitempty" name:"EngineTypes"`

	// 模板名称，缺省则查询所有
	TemplateNames []*string `json:"TemplateNames,omitempty" name:"TemplateNames"`

	// 模板id，缺省则查询所有
	TemplateIds []*int64 `json:"TemplateIds,omitempty" name:"TemplateIds"`
}

func (r *DescribeParamTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineVersions")
	delete(f, "EngineTypes")
	delete(f, "TemplateNames")
	delete(f, "TemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplatesResponseParams struct {
	// 该用户的参数模板数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 参数模板详情。
	Items []*ParamTemplateInfo `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeParamTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParamTemplatesResponseParams `json:"Response"`
}

func (r *DescribeParamTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsRequestParams struct {
	// 项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsResponseParams struct {
	// 安全组详情。
	Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`

	// 安全组规则数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProjectSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeProjectSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxyConnectionPoolConfRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeProxyConnectionPoolConfRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeProxyConnectionPoolConfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyConnectionPoolConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxyConnectionPoolConfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxyConnectionPoolConfResponseParams struct {
	// 配置规格数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 连接池配置规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoolConf *PoolConf `json:"PoolConf,omitempty" name:"PoolConf"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProxyConnectionPoolConfResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxyConnectionPoolConfResponseParams `json:"Response"`
}

func (r *DescribeProxyConnectionPoolConfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyConnectionPoolConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxyCustomConfRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeProxyCustomConfRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeProxyCustomConfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyCustomConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxyCustomConfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxyCustomConfResponseParams struct {
	// 代理配置数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 代理配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomConf *CustomConfig `json:"CustomConf,omitempty" name:"CustomConf"`

	// 权重限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeightRule *Rule `json:"WeightRule,omitempty" name:"WeightRule"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProxyCustomConfResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxyCustomConfResponseParams `json:"Response"`
}

func (r *DescribeProxyCustomConfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyCustomConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRemoteBackupConfigRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeRemoteBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeRemoteBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRemoteBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRemoteBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRemoteBackupConfigResponseParams struct {
	// 异地备份保留天时间，单位为天
	ExpireDays *int64 `json:"ExpireDays,omitempty" name:"ExpireDays"`

	// 异地数据备份开关，off - 关闭异地备份，on-开启异地备份
	RemoteBackupSave *string `json:"RemoteBackupSave,omitempty" name:"RemoteBackupSave"`

	// 异地日志备份开关，off - 关闭异地备份，on-开启异地备份，只有在参数RemoteBackupSave为on时，RemoteBinlogSave参数才可设置为on
	RemoteBinlogSave *string `json:"RemoteBinlogSave,omitempty" name:"RemoteBinlogSave"`

	// 用户已设置异地备份地域列表
	RemoteRegion []*string `json:"RemoteRegion,omitempty" name:"RemoteRegion"`

	// 用户可设置异地备份地域列表
	RegionList []*string `json:"RegionList,omitempty" name:"RegionList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRemoteBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRemoteBackupConfigResponseParams `json:"Response"`
}

func (r *DescribeRemoteBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRemoteBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoGroupsRequestParams struct {
	// 实例ID，格式如：cdb-c1nl9rpv或者cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoGroupsResponseParams struct {
	// RO组信息数组，一个实例可关联多个RO组。
	RoGroups []*RoGroup `json:"RoGroups,omitempty" name:"RoGroups"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRoGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoGroupsResponseParams `json:"Response"`
}

func (r *DescribeRoGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoMinScaleRequestParams struct {
	// 只读实例ID，格式如：cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，该参数与MasterInstanceId参数不能同时为空。
	RoInstanceId *string `json:"RoInstanceId,omitempty" name:"RoInstanceId"`

	// 主实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，该参数与RoInstanceId参数不能同时为空。注意，当传入参数包含RoInstanceId时，返回值为只读实例升级时的最小规格；当传入参数只包含MasterInstanceId时，返回值为只读实例购买时的最小规格。
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" name:"MasterInstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoMinScaleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoInstanceId")
	delete(f, "MasterInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoMinScaleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoMinScaleResponseParams struct {
	// 内存规格大小, 单位为：MB。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 磁盘规格大小, 单位为：GB。
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRoMinScaleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoMinScaleResponseParams `json:"Response"`
}

func (r *DescribeRoMinScaleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoMinScaleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollbackRangeTimeRequestParams struct {
	// 实例 ID 列表，单个实例 ID 的格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DescribeRollbackRangeTimeRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表，单个实例 ID 的格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeRollbackRangeTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackRangeTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRollbackRangeTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollbackRangeTimeResponseParams struct {
	// 符合查询条件的实例总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回的参数信息。
	Items []*InstanceRollbackRangeTime `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRollbackRangeTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRollbackRangeTimeResponseParams `json:"Response"`
}

func (r *DescribeRollbackRangeTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackRangeTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollbackTaskDetailRequestParams struct {
	// 实例 ID。与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872)。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 异步任务 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 分页参数，每次请求返回的记录数。默认值为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量。默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AsyncRequestId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRollbackTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollbackTaskDetailResponseParams struct {
	// 符合条件的记录总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 回档任务详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*RollbackTask `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRollbackTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRollbackTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeRollbackTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogDataRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间戳。例如 1585142640 。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳。例如 1585142640 。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 客户端 Host 列表。
	UserHosts []*string `json:"UserHosts,omitempty" name:"UserHosts"`

	// 客户端 用户名 列表。
	UserNames []*string `json:"UserNames,omitempty" name:"UserNames"`

	// 访问的 数据库 列表。
	DataBases []*string `json:"DataBases,omitempty" name:"DataBases"`

	// 排序字段。当前支持：Timestamp,QueryTime,LockTime,RowsExamined,RowsSent 。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 升序还是降序排列。当前支持：ASC,DESC 。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 一次性返回的记录数量，默认为100，最大为400。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 仅在实例为主实例或者灾备实例时生效，可选值：slave，代表拉取从机的日志。
	InstType *string `json:"InstType,omitempty" name:"InstType"`
}

type DescribeSlowLogDataRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间戳。例如 1585142640 。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳。例如 1585142640 。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 客户端 Host 列表。
	UserHosts []*string `json:"UserHosts,omitempty" name:"UserHosts"`

	// 客户端 用户名 列表。
	UserNames []*string `json:"UserNames,omitempty" name:"UserNames"`

	// 访问的 数据库 列表。
	DataBases []*string `json:"DataBases,omitempty" name:"DataBases"`

	// 排序字段。当前支持：Timestamp,QueryTime,LockTime,RowsExamined,RowsSent 。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 升序还是降序排列。当前支持：ASC,DESC 。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 一次性返回的记录数量，默认为100，最大为400。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 仅在实例为主实例或者灾备实例时生效，可选值：slave，代表拉取从机的日志。
	InstType *string `json:"InstType,omitempty" name:"InstType"`
}

func (r *DescribeSlowLogDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "UserHosts")
	delete(f, "UserNames")
	delete(f, "DataBases")
	delete(f, "SortBy")
	delete(f, "OrderBy")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogDataResponseParams struct {
	// 符合条件的记录总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 查询到的记录。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*SlowLogItem `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSlowLogDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogDataResponseParams `json:"Response"`
}

func (r *DescribeSlowLogDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogsRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量，默认值为0，最小值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认值为20，最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量，默认值为0，最小值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认值为20，最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogsResponseParams struct {
	// 符合查询条件的慢查询日志总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 符合查询条件的慢查询日志详情。
	Items []*SlowLogInfo `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportedPrivilegesRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportedPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSupportedPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportedPrivilegesResponseParams struct {
	// 实例支持的全局权限。
	GlobalSupportedPrivileges []*string `json:"GlobalSupportedPrivileges,omitempty" name:"GlobalSupportedPrivileges"`

	// 实例支持的数据库权限。
	DatabaseSupportedPrivileges []*string `json:"DatabaseSupportedPrivileges,omitempty" name:"DatabaseSupportedPrivileges"`

	// 实例支持的数据库表权限。
	TableSupportedPrivileges []*string `json:"TableSupportedPrivileges,omitempty" name:"TableSupportedPrivileges"`

	// 实例支持的数据库列权限。
	ColumnSupportedPrivileges []*string `json:"ColumnSupportedPrivileges,omitempty" name:"ColumnSupportedPrivileges"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSupportedPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSupportedPrivilegesResponseParams `json:"Response"`
}

func (r *DescribeSupportedPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportedPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Database")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TableRegexp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesResponseParams struct {
	// 符合查询条件的数据库表总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回的数据库表信息。
	Items []*string `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTablesResponseParams `json:"Response"`
}

func (r *DescribeTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsOfInstanceIdsRequestParams struct {
	// 实例列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTagsOfInstanceIdsRequest struct {
	*tchttp.BaseRequest
	
	// 实例列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTagsOfInstanceIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsOfInstanceIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagsOfInstanceIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsOfInstanceIdsResponseParams struct {
	// 分页偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 实例标签信息。
	Rows []*TagsInfoOfInstance `json:"Rows,omitempty" name:"Rows"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTagsOfInstanceIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagsOfInstanceIdsResponseParams `json:"Response"`
}

func (r *DescribeTagsOfInstanceIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsOfInstanceIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksRequestParams struct {
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
	TaskTypes []*int64 `json:"TaskTypes,omitempty" name:"TaskTypes"`

	// 任务状态，不传值则查询所有任务状态，支持的值包括：
	// -1 - 未定义；
	// 0 - 初始化；
	// 1 - 运行中；
	// 2 - 执行成功；
	// 3 - 执行失败；
	// 4 - 已终止；
	// 5 - 已删除；
	// 6 - 已暂停。
	TaskStatus []*int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 第一个任务的开始时间，用于范围查询，时间格式如：2017-12-31 10:40:01。
	StartTimeBegin *string `json:"StartTimeBegin,omitempty" name:"StartTimeBegin"`

	// 最后一个任务的开始时间，用于范围查询，时间格式如：2017-12-31 10:40:01。
	StartTimeEnd *string `json:"StartTimeEnd,omitempty" name:"StartTimeEnd"`

	// 记录偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	TaskTypes []*int64 `json:"TaskTypes,omitempty" name:"TaskTypes"`

	// 任务状态，不传值则查询所有任务状态，支持的值包括：
	// -1 - 未定义；
	// 0 - 初始化；
	// 1 - 运行中；
	// 2 - 执行成功；
	// 3 - 执行失败；
	// 4 - 已终止；
	// 5 - 已删除；
	// 6 - 已暂停。
	TaskStatus []*int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AsyncRequestId")
	delete(f, "TaskTypes")
	delete(f, "TaskStatus")
	delete(f, "StartTimeBegin")
	delete(f, "StartTimeEnd")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// 符合查询条件的实例总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回的实例任务信息。
	Items []*TaskDetail `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksResponseParams `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimeWindowRequestParams struct {
	// 实例ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimeWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimeWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimeWindowResponseParams struct {
	// 星期一的可维护时间列表。
	Monday []*string `json:"Monday,omitempty" name:"Monday"`

	// 星期二的可维护时间列表。
	Tuesday []*string `json:"Tuesday,omitempty" name:"Tuesday"`

	// 星期三的可维护时间列表。
	Wednesday []*string `json:"Wednesday,omitempty" name:"Wednesday"`

	// 星期四的可维护时间列表。
	Thursday []*string `json:"Thursday,omitempty" name:"Thursday"`

	// 星期五的可维护时间列表。
	Friday []*string `json:"Friday,omitempty" name:"Friday"`

	// 星期六的可维护时间列表。
	Saturday []*string `json:"Saturday,omitempty" name:"Saturday"`

	// 星期日的可维护时间列表。
	Sunday []*string `json:"Sunday,omitempty" name:"Sunday"`

	// 最大数据延迟阈值
	MaxDelayTime *uint64 `json:"MaxDelayTime,omitempty" name:"MaxDelayTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTimeWindowResponseParams `json:"Response"`
}

func (r *DescribeTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimeWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUploadedFilesRequestParams struct {
	// 文件路径。该字段应填用户主账号的OwnerUin信息。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 记录偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUploadedFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Path")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUploadedFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUploadedFilesResponseParams struct {
	// 符合查询条件的SQL文件总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回的SQL文件列表。
	Items []*SqlFileInfo `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUploadedFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUploadedFilesResponseParams `json:"Response"`
}

func (r *DescribeUploadedFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUploadedFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceCpuInfo struct {
	// 实例CPU平均使用率
	Rate []*DeviceCpuRateInfo `json:"Rate,omitempty" name:"Rate"`

	// 实例CPU监控数据
	Load []*int64 `json:"Load,omitempty" name:"Load"`
}

type DeviceCpuRateInfo struct {
	// Cpu核编号
	CpuCore *int64 `json:"CpuCore,omitempty" name:"CpuCore"`

	// Cpu使用率
	Rate []*int64 `json:"Rate,omitempty" name:"Rate"`
}

type DeviceDiskInfo struct {
	// 平均每秒有百分之几的时间用于IO操作
	IoRatioPerSec []*int64 `json:"IoRatioPerSec,omitempty" name:"IoRatioPerSec"`

	// 平均每次设备I/O操作的等待时间*100，单位为毫秒。例如：该值为201，表示平均每次I/O操作等待时间为：201/100=2.1毫秒
	IoWaitTime []*int64 `json:"IoWaitTime,omitempty" name:"IoWaitTime"`

	// 磁盘平均每秒完成的读操作次数总和*100。例如：该值为2002，表示磁盘平均每秒完成读操作为：2002/100=20.2次
	Read []*int64 `json:"Read,omitempty" name:"Read"`

	// 磁盘平均每秒完成的写操作次数总和*100。例如：该值为30001，表示磁盘平均每秒完成写操作为：30001/100=300.01次
	Write []*int64 `json:"Write,omitempty" name:"Write"`

	// 磁盘空间容量，每两个一组，第一个为已使用容量，第二个为磁盘总容量
	CapacityRatio []*int64 `json:"CapacityRatio,omitempty" name:"CapacityRatio"`
}

type DeviceMemInfo struct {
	// 总内存大小。free命令中Mem:一行total的值,单位：KB
	Total []*int64 `json:"Total,omitempty" name:"Total"`

	// 已使用内存。free命令中Mem:一行used的值,单位：KB
	Used []*int64 `json:"Used,omitempty" name:"Used"`
}

type DeviceNetInfo struct {
	// tcp连接数
	Conn []*int64 `json:"Conn,omitempty" name:"Conn"`

	// 网卡入包量，单位：个/秒
	PackageIn []*int64 `json:"PackageIn,omitempty" name:"PackageIn"`

	// 网卡出包量，单位：个/秒
	PackageOut []*int64 `json:"PackageOut,omitempty" name:"PackageOut"`

	// 入流量，单位：kbps
	FlowIn []*int64 `json:"FlowIn,omitempty" name:"FlowIn"`

	// 出流量，单位：kbps
	FlowOut []*int64 `json:"FlowOut,omitempty" name:"FlowOut"`
}

// Predefined struct for user
type DisassociateSecurityGroupsRequestParams struct {
	// 安全组 ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例 ID 列表，一个或者多个实例 ID 组成的数组。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 当传入只读实例ID时，默认操作的是对应只读组的安全组。如果需要操作只读实例ID的安全组， 需要将该入参置为True
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitempty" name:"ForReadonlyInstance"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 安全组 ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例 ID 列表，一个或者多个实例 ID 组成的数组。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 当传入只读实例ID时，默认操作的是对应只读组的安全组。如果需要操作只读实例ID的安全组， 需要将该入参置为True
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitempty" name:"ForReadonlyInstance"`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIds")
	delete(f, "ForReadonlyInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// 规则描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

// Predefined struct for user
type InitDBInstancesRequestParams struct {
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，可使用[查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 实例新的密码，密码规则：8-64个字符，至少包含字母、数字、字符（支持的字符：!@#$%^*()）中的两种。
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// 实例的参数列表，目前支持设置“character_set_server”、“lower_case_table_names”参数。其中，“character_set_server”参数可选值为["utf8","latin1","gbk","utf8mb4"]；“lower_case_table_names”可选值为[“0”,“1”]。
	Parameters []*ParamInfo `json:"Parameters,omitempty" name:"Parameters"`

	// 实例的端口，取值范围为[1024, 65535]
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
}

type InitDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，可使用[查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 实例新的密码，密码规则：8-64个字符，至少包含字母、数字、字符（支持的字符：!@#$%^*()）中的两种。
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// 实例的参数列表，目前支持设置“character_set_server”、“lower_case_table_names”参数。其中，“character_set_server”参数可选值为["utf8","latin1","gbk","utf8mb4"]；“lower_case_table_names”可选值为[“0”,“1”]。
	Parameters []*ParamInfo `json:"Parameters,omitempty" name:"Parameters"`

	// 实例的端口，取值范围为[1024, 65535]
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
}

func (r *InitDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "NewPassword")
	delete(f, "Parameters")
	delete(f, "Vport")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InitDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InitDBInstancesResponseParams struct {
	// 异步任务的请求ID数组，可使用此ID查询异步任务的执行结果
	AsyncRequestIds []*string `json:"AsyncRequestIds,omitempty" name:"AsyncRequestIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InitDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InitDBInstancesResponseParams `json:"Response"`
}

func (r *InitDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpgradeInstancesRequestParams struct {
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

	// 实例隔离类型。支持值包括： "UNIVERSAL" - 通用型实例， "EXCLUSIVE" - 独享型实例， "BASIC" - 基础版实例。 不指定则默认为通用型实例。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 实例节点数。对于 RO 和 基础版实例， 该值默认为1。 如果需要询价三节点实例， 请将该值设置为3。其余主实例该值默认为2。
	InstanceNodes *int64 `json:"InstanceNodes,omitempty" name:"InstanceNodes"`
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

	// 实例隔离类型。支持值包括： "UNIVERSAL" - 通用型实例， "EXCLUSIVE" - 独享型实例， "BASIC" - 基础版实例。 不指定则默认为通用型实例。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 实例节点数。对于 RO 和 基础版实例， 该值默认为1。 如果需要询价三节点实例， 请将该值设置为3。其余主实例该值默认为2。
	InstanceNodes *int64 `json:"InstanceNodes,omitempty" name:"InstanceNodes"`
}

func (r *InquiryPriceUpgradeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "Cpu")
	delete(f, "ProtectMode")
	delete(f, "DeviceType")
	delete(f, "InstanceNodes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceUpgradeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpgradeInstancesResponseParams struct {
	// 实例价格，单位：分（人民币）。
	Price *int64 `json:"Price,omitempty" name:"Price"`

	// 实例原价，单位：分（人民币）。
	OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquiryPriceUpgradeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceUpgradeInstancesResponseParams `json:"Response"`
}

func (r *InquiryPriceUpgradeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// 实例状态，可能的返回值：0-创建中；1-运行中；4-正在进行隔离操作；5-已隔离
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
	RoGroups []*RoGroup `json:"RoGroups,omitempty" name:"RoGroups"`

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

	// 实例类型
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 内核版本
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 灾备实例详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DrInfo []*DrInfo `json:"DrInfo,omitempty" name:"DrInfo"`

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

	// 磁盘写入是否被锁定（实例数据写入量已经超过磁盘配额）。0 -未被锁定 1 -已被锁定
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

	// 节点数
	InstanceNodes *int64 `json:"InstanceNodes,omitempty" name:"InstanceNodes"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*TagInfoItem `json:"TagList,omitempty" name:"TagList"`

	// 引擎类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`

	// 最大延迟阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDelayTime *int64 `json:"MaxDelayTime,omitempty" name:"MaxDelayTime"`
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
	Times []*RollbackTimeRange `json:"Times,omitempty" name:"Times"`
}

// Predefined struct for user
type IsolateDBInstanceRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDBInstanceResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。(该返回字段目前已废弃，可以通过 DescribeDBInstances 接口查询实例的隔离状态)
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type LocalBinlogConfig struct {
	// 本地binlog保留时长，可取值范围：[72,168]。
	SaveHours *int64 `json:"SaveHours,omitempty" name:"SaveHours"`

	// 本地binlog空间使用率，可取值范围：[30,50]。
	MaxUsage *int64 `json:"MaxUsage,omitempty" name:"MaxUsage"`
}

type LocalBinlogConfigDefault struct {
	// 本地binlog保留时长，可取值范围：[72,168]。
	SaveHours *int64 `json:"SaveHours,omitempty" name:"SaveHours"`

	// 本地binlog空间使用率，可取值范围：[30,50]。
	MaxUsage *int64 `json:"MaxUsage,omitempty" name:"MaxUsage"`
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

// Predefined struct for user
type ModifyAccountDescriptionRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 云数据库账号。
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`

	// 数据库账号的备注信息。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 云数据库账号。
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`

	// 数据库账号的备注信息。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAccountDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountDescriptionResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAccountDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountDescriptionResponseParams `json:"Response"`
}

func (r *ModifyAccountDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountHostRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 账户的名称
	User *string `json:"User,omitempty" name:"User"`

	// 账户的旧主机
	Host *string `json:"Host,omitempty" name:"Host"`

	// 账户的新主机
	NewHost *string `json:"NewHost,omitempty" name:"NewHost"`
}

type ModifyAccountHostRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 账户的名称
	User *string `json:"User,omitempty" name:"User"`

	// 账户的旧主机
	Host *string `json:"Host,omitempty" name:"Host"`

	// 账户的新主机
	NewHost *string `json:"NewHost,omitempty" name:"NewHost"`
}

func (r *ModifyAccountHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "Host")
	delete(f, "NewHost")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountHostResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAccountHostResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountHostResponseParams `json:"Response"`
}

func (r *ModifyAccountHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountMaxUserConnectionsRequestParams struct {
	// 云数据库账号。
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 设置账户最大可用连接数，最大可设置值为10240。
	MaxUserConnections *int64 `json:"MaxUserConnections,omitempty" name:"MaxUserConnections"`
}

type ModifyAccountMaxUserConnectionsRequest struct {
	*tchttp.BaseRequest
	
	// 云数据库账号。
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 设置账户最大可用连接数，最大可设置值为10240。
	MaxUserConnections *int64 `json:"MaxUserConnections,omitempty" name:"MaxUserConnections"`
}

func (r *ModifyAccountMaxUserConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountMaxUserConnectionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Accounts")
	delete(f, "InstanceId")
	delete(f, "MaxUserConnections")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountMaxUserConnectionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountMaxUserConnectionsResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAccountMaxUserConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountMaxUserConnectionsResponseParams `json:"Response"`
}

func (r *ModifyAccountMaxUserConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountMaxUserConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPasswordRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库账号的新密码。密码应至少包含字母、数字和字符（_+-&=!@#$%^*()）中的两种，长度为8-64个字符。
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// 云数据库账号。
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`
}

type ModifyAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库账号的新密码。密码应至少包含字母、数字和字符（_+-&=!@#$%^*()）中的两种，长度为8-64个字符。
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// 云数据库账号。
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`
}

func (r *ModifyAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NewPassword")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPasswordResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountPasswordResponseParams `json:"Response"`
}

func (r *ModifyAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPrivilegesRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库的账号，包括用户名和域名。
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`

	// 全局权限。其中，GlobalPrivileges 中权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE", "PROCESS", "DROP","REFERENCES","INDEX","ALTER","SHOW DATABASES","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER","CREATE USER","RELOAD","REPLICATION CLIENT","REPLICATION SLAVE"。
	// 注意，ModifyAction为空时，不传该参数表示清除该权限。
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitempty" name:"GlobalPrivileges"`

	// 数据库的权限。Privileges 权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE",	"DROP","REFERENCES","INDEX","ALTER","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER"。
	// 注意，ModifyAction为空时，不传该参数表示清除该权限。
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitempty" name:"DatabasePrivileges"`

	// 数据库中表的权限。Privileges 权限的可选值为：权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE",	"DROP","REFERENCES","INDEX","ALTER","CREATE VIEW","SHOW VIEW", "TRIGGER"。
	// 注意，ModifyAction为空时，不传该参数表示清除该权限。
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitempty" name:"TablePrivileges"`

	// 数据库表中列的权限。Privileges 权限的可选值为："SELECT","INSERT","UPDATE","REFERENCES"。
	// 注意，ModifyAction为空时，不传该参数表示清除该权限。
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitempty" name:"ColumnPrivileges"`

	// 该参数不为空时，为批量修改权限。可选值为：grant - 授予权限，revoke - 回收权限。
	ModifyAction *string `json:"ModifyAction,omitempty" name:"ModifyAction"`
}

type ModifyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库的账号，包括用户名和域名。
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`

	// 全局权限。其中，GlobalPrivileges 中权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE", "PROCESS", "DROP","REFERENCES","INDEX","ALTER","SHOW DATABASES","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER","CREATE USER","RELOAD","REPLICATION CLIENT","REPLICATION SLAVE"。
	// 注意，ModifyAction为空时，不传该参数表示清除该权限。
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitempty" name:"GlobalPrivileges"`

	// 数据库的权限。Privileges 权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE",	"DROP","REFERENCES","INDEX","ALTER","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER"。
	// 注意，ModifyAction为空时，不传该参数表示清除该权限。
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitempty" name:"DatabasePrivileges"`

	// 数据库中表的权限。Privileges 权限的可选值为：权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE",	"DROP","REFERENCES","INDEX","ALTER","CREATE VIEW","SHOW VIEW", "TRIGGER"。
	// 注意，ModifyAction为空时，不传该参数表示清除该权限。
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitempty" name:"TablePrivileges"`

	// 数据库表中列的权限。Privileges 权限的可选值为："SELECT","INSERT","UPDATE","REFERENCES"。
	// 注意，ModifyAction为空时，不传该参数表示清除该权限。
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitempty" name:"ColumnPrivileges"`

	// 该参数不为空时，为批量修改权限。可选值为：grant - 授予权限，revoke - 回收权限。
	ModifyAction *string `json:"ModifyAction,omitempty" name:"ModifyAction"`
}

func (r *ModifyAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	delete(f, "GlobalPrivileges")
	delete(f, "DatabasePrivileges")
	delete(f, "TablePrivileges")
	delete(f, "ColumnPrivileges")
	delete(f, "ModifyAction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPrivilegesResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountPrivilegesResponseParams `json:"Response"`
}

func (r *ModifyAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditConfigRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 审计日志保存时长。支持值包括：
	// 7 - 一周
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

	// 高频审计日志保存时长。支持值包括：
	// 7 - 一周
	// 30 - 一个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年；
	HighLogExpireDay *int64 `json:"HighLogExpireDay,omitempty" name:"HighLogExpireDay"`
}

type ModifyAuditConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 审计日志保存时长。支持值包括：
	// 7 - 一周
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

	// 高频审计日志保存时长。支持值包括：
	// 7 - 一周
	// 30 - 一个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年；
	HighLogExpireDay *int64 `json:"HighLogExpireDay,omitempty" name:"HighLogExpireDay"`
}

func (r *ModifyAuditConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	delete(f, "CloseAudit")
	delete(f, "HighLogExpireDay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuditConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAuditConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAuditConfigResponseParams `json:"Response"`
}

func (r *ModifyAuditConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditRuleRequestParams struct {
	// 审计规则 ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 审计规则名称。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 审计规则描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 审计规则过滤条件。若设置了过滤条件，则不会开启全审计。
	RuleFilters []*AuditFilter `json:"RuleFilters,omitempty" name:"RuleFilters"`

	// 是否开启全审计。支持值包括：false – 不开启全审计，true – 开启全审计。用户未设置审计规则过滤条件时，默认开启全审计。
	AuditAll *bool `json:"AuditAll,omitempty" name:"AuditAll"`
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
	RuleFilters []*AuditFilter `json:"RuleFilters,omitempty" name:"RuleFilters"`

	// 是否开启全审计。支持值包括：false – 不开启全审计，true – 开启全审计。用户未设置审计规则过滤条件时，默认开启全审计。
	AuditAll *bool `json:"AuditAll,omitempty" name:"AuditAll"`
}

func (r *ModifyAuditRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "RuleName")
	delete(f, "Description")
	delete(f, "RuleFilters")
	delete(f, "AuditAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuditRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAuditRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAuditRuleResponseParams `json:"Response"`
}

func (r *ModifyAuditRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoRenewFlagRequestParams struct {
	// 实例的 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 自动续费标记，可取值的有：0 - 不自动续费，1 - 自动续费。
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`
}

type ModifyAutoRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 实例的 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 自动续费标记，可取值的有：0 - 不自动续费，1 - 自动续费。
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`
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
	delete(f, "AutoRenew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoRenewFlagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyBackupConfigRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份文件的保留时间，单位为天。最小值为7天，最大值为1830天。
	ExpireDays *int64 `json:"ExpireDays,omitempty" name:"ExpireDays"`

	// (将废弃，建议使用 BackupTimeWindow 参数) 备份时间范围，格式为：02:00-06:00，起点和终点时间目前限制为整点，目前可以选择的范围为： 00:00-12:00，02:00-06:00，06：00-10：00，10:00-14:00，14:00-18:00，18:00-22:00，22:00-02:00。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 自动备份方式，仅支持：physical - 物理冷备
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// binlog的保留时间，单位为天。最小值为7天，最大值为1830天。该值的设置不能大于备份文件的保留时间。
	BinlogExpireDays *int64 `json:"BinlogExpireDays,omitempty" name:"BinlogExpireDays"`

	// 备份时间窗，比如要设置每周二和周日 10:00-14:00之间备份，该参数如下：{"Monday": "", "Tuesday": "10:00-14:00", "Wednesday": "", "Thursday": "", "Friday": "", "Saturday": "", "Sunday": "10:00-14:00"}    （注：可以设置一周的某几天备份，但是每天的备份时间需要设置为相同的时间段。 如果设置了该字段，将忽略StartTime字段的设置）
	BackupTimeWindow *CommonTimeWindow `json:"BackupTimeWindow,omitempty" name:"BackupTimeWindow"`

	// 定期保留开关，off - 不开启定期保留策略，on - 开启定期保留策略，默认为off。首次开启定期保留策略时，BackupPeriodSaveDays，BackupPeriodSaveInterval，BackupPeriodSaveCount，StartBackupPeriodSaveDate参数为必填项，否则定期保留策略不会生效
	EnableBackupPeriodSave *string `json:"EnableBackupPeriodSave,omitempty" name:"EnableBackupPeriodSave"`

	// 长期保留开关,该字段功能暂未上线，可忽略。off - 不开启长期保留策略，on - 开启长期保留策略，默认为off，如果开启，则BackupPeriodSaveDays，BackupPeriodSaveInterval，BackupPeriodSaveCount参数无效
	EnableBackupPeriodLongTermSave *string `json:"EnableBackupPeriodLongTermSave,omitempty" name:"EnableBackupPeriodLongTermSave"`

	// 定期保留最长天数，最小值：90，最大值：3650，默认值：1080
	BackupPeriodSaveDays *int64 `json:"BackupPeriodSaveDays,omitempty" name:"BackupPeriodSaveDays"`

	// 定期保留策略周期，可取值：weekly - 周，monthly - 月， quarterly - 季度，yearly - 年，默认为monthly
	BackupPeriodSaveInterval *string `json:"BackupPeriodSaveInterval,omitempty" name:"BackupPeriodSaveInterval"`

	// 定期保留的备份数量，最小值为1，最大值不超过定期保留策略周期内常规备份个数，默认值为1
	BackupPeriodSaveCount *int64 `json:"BackupPeriodSaveCount,omitempty" name:"BackupPeriodSaveCount"`

	// 定期保留策略周期起始日期，格式：YYYY-MM-dd HH:mm:ss
	StartBackupPeriodSaveDate *string `json:"StartBackupPeriodSaveDate,omitempty" name:"StartBackupPeriodSaveDate"`

	// 是否开启数据备份归档策略，off-关闭，on-打开，默认为off
	EnableBackupArchive *string `json:"EnableBackupArchive,omitempty" name:"EnableBackupArchive"`

	// 数据备份归档起始天数，数据备份达到归档起始天数时进行归档，最小为180天，不得大于数据备份保留天数
	BackupArchiveDays *int64 `json:"BackupArchiveDays,omitempty" name:"BackupArchiveDays"`

	// 日志备份归档起始天数，日志备份达到归档起始天数时进行归档，最小为180天，不得大于日志备份保留天数
	BinlogArchiveDays *int64 `json:"BinlogArchiveDays,omitempty" name:"BinlogArchiveDays"`

	// 是否开启日志备份归档策略，off-关闭，on-打开，默认为off
	EnableBinlogArchive *string `json:"EnableBinlogArchive,omitempty" name:"EnableBinlogArchive"`

	// 是否开启数据备份标准存储策略，off-关闭，on-打开，默认为off
	EnableBackupStandby *string `json:"EnableBackupStandby,omitempty" name:"EnableBackupStandby"`

	// 数据备份标准存储起始天数，数据备份达到标准存储起始天数时进行转换，最小为30天，不得大于数据备份保留天数。如果开启备份归档，不得大于等于备份归档天数
	BackupStandbyDays *int64 `json:"BackupStandbyDays,omitempty" name:"BackupStandbyDays"`

	// 是否开启日志备份标准存储策略，off-关闭，on-打开，默认为off
	EnableBinlogStandby *string `json:"EnableBinlogStandby,omitempty" name:"EnableBinlogStandby"`

	// 日志备份标准存储起始天数，日志备份达到标准存储起始天数时进行转换，最小为30天，不得大于日志备份保留天数。如果开启备份归档，不得大于等于备份归档天数
	BinlogStandbyDays *int64 `json:"BinlogStandbyDays,omitempty" name:"BinlogStandbyDays"`
}

type ModifyBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份文件的保留时间，单位为天。最小值为7天，最大值为1830天。
	ExpireDays *int64 `json:"ExpireDays,omitempty" name:"ExpireDays"`

	// (将废弃，建议使用 BackupTimeWindow 参数) 备份时间范围，格式为：02:00-06:00，起点和终点时间目前限制为整点，目前可以选择的范围为： 00:00-12:00，02:00-06:00，06：00-10：00，10:00-14:00，14:00-18:00，18:00-22:00，22:00-02:00。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 自动备份方式，仅支持：physical - 物理冷备
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// binlog的保留时间，单位为天。最小值为7天，最大值为1830天。该值的设置不能大于备份文件的保留时间。
	BinlogExpireDays *int64 `json:"BinlogExpireDays,omitempty" name:"BinlogExpireDays"`

	// 备份时间窗，比如要设置每周二和周日 10:00-14:00之间备份，该参数如下：{"Monday": "", "Tuesday": "10:00-14:00", "Wednesday": "", "Thursday": "", "Friday": "", "Saturday": "", "Sunday": "10:00-14:00"}    （注：可以设置一周的某几天备份，但是每天的备份时间需要设置为相同的时间段。 如果设置了该字段，将忽略StartTime字段的设置）
	BackupTimeWindow *CommonTimeWindow `json:"BackupTimeWindow,omitempty" name:"BackupTimeWindow"`

	// 定期保留开关，off - 不开启定期保留策略，on - 开启定期保留策略，默认为off。首次开启定期保留策略时，BackupPeriodSaveDays，BackupPeriodSaveInterval，BackupPeriodSaveCount，StartBackupPeriodSaveDate参数为必填项，否则定期保留策略不会生效
	EnableBackupPeriodSave *string `json:"EnableBackupPeriodSave,omitempty" name:"EnableBackupPeriodSave"`

	// 长期保留开关,该字段功能暂未上线，可忽略。off - 不开启长期保留策略，on - 开启长期保留策略，默认为off，如果开启，则BackupPeriodSaveDays，BackupPeriodSaveInterval，BackupPeriodSaveCount参数无效
	EnableBackupPeriodLongTermSave *string `json:"EnableBackupPeriodLongTermSave,omitempty" name:"EnableBackupPeriodLongTermSave"`

	// 定期保留最长天数，最小值：90，最大值：3650，默认值：1080
	BackupPeriodSaveDays *int64 `json:"BackupPeriodSaveDays,omitempty" name:"BackupPeriodSaveDays"`

	// 定期保留策略周期，可取值：weekly - 周，monthly - 月， quarterly - 季度，yearly - 年，默认为monthly
	BackupPeriodSaveInterval *string `json:"BackupPeriodSaveInterval,omitempty" name:"BackupPeriodSaveInterval"`

	// 定期保留的备份数量，最小值为1，最大值不超过定期保留策略周期内常规备份个数，默认值为1
	BackupPeriodSaveCount *int64 `json:"BackupPeriodSaveCount,omitempty" name:"BackupPeriodSaveCount"`

	// 定期保留策略周期起始日期，格式：YYYY-MM-dd HH:mm:ss
	StartBackupPeriodSaveDate *string `json:"StartBackupPeriodSaveDate,omitempty" name:"StartBackupPeriodSaveDate"`

	// 是否开启数据备份归档策略，off-关闭，on-打开，默认为off
	EnableBackupArchive *string `json:"EnableBackupArchive,omitempty" name:"EnableBackupArchive"`

	// 数据备份归档起始天数，数据备份达到归档起始天数时进行归档，最小为180天，不得大于数据备份保留天数
	BackupArchiveDays *int64 `json:"BackupArchiveDays,omitempty" name:"BackupArchiveDays"`

	// 日志备份归档起始天数，日志备份达到归档起始天数时进行归档，最小为180天，不得大于日志备份保留天数
	BinlogArchiveDays *int64 `json:"BinlogArchiveDays,omitempty" name:"BinlogArchiveDays"`

	// 是否开启日志备份归档策略，off-关闭，on-打开，默认为off
	EnableBinlogArchive *string `json:"EnableBinlogArchive,omitempty" name:"EnableBinlogArchive"`

	// 是否开启数据备份标准存储策略，off-关闭，on-打开，默认为off
	EnableBackupStandby *string `json:"EnableBackupStandby,omitempty" name:"EnableBackupStandby"`

	// 数据备份标准存储起始天数，数据备份达到标准存储起始天数时进行转换，最小为30天，不得大于数据备份保留天数。如果开启备份归档，不得大于等于备份归档天数
	BackupStandbyDays *int64 `json:"BackupStandbyDays,omitempty" name:"BackupStandbyDays"`

	// 是否开启日志备份标准存储策略，off-关闭，on-打开，默认为off
	EnableBinlogStandby *string `json:"EnableBinlogStandby,omitempty" name:"EnableBinlogStandby"`

	// 日志备份标准存储起始天数，日志备份达到标准存储起始天数时进行转换，最小为30天，不得大于日志备份保留天数。如果开启备份归档，不得大于等于备份归档天数
	BinlogStandbyDays *int64 `json:"BinlogStandbyDays,omitempty" name:"BinlogStandbyDays"`
}

func (r *ModifyBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ExpireDays")
	delete(f, "StartTime")
	delete(f, "BackupMethod")
	delete(f, "BinlogExpireDays")
	delete(f, "BackupTimeWindow")
	delete(f, "EnableBackupPeriodSave")
	delete(f, "EnableBackupPeriodLongTermSave")
	delete(f, "BackupPeriodSaveDays")
	delete(f, "BackupPeriodSaveInterval")
	delete(f, "BackupPeriodSaveCount")
	delete(f, "StartBackupPeriodSaveDate")
	delete(f, "EnableBackupArchive")
	delete(f, "BackupArchiveDays")
	delete(f, "BinlogArchiveDays")
	delete(f, "EnableBinlogArchive")
	delete(f, "EnableBackupStandby")
	delete(f, "BackupStandbyDays")
	delete(f, "EnableBinlogStandby")
	delete(f, "BinlogStandbyDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupConfigResponseParams `json:"Response"`
}

func (r *ModifyBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadRestrictionRequestParams struct {
	// NoLimit 不限制,内外网都可以下载； LimitOnlyIntranet 仅内网可下载； Customize 用户自定义vpc:ip可下载。 只有该值为 Customize 时，才可以设置 LimitVpc 和 LimitIp 。
	LimitType *string `json:"LimitType,omitempty" name:"LimitType"`

	// 该参数仅支持 In， 表示 LimitVpc 指定的vpc可以下载。默认为In。
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitempty" name:"VpcComparisonSymbol"`

	// In: 指定的ip可以下载； NotIn: 指定的ip不可以下载。 默认为In。
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitempty" name:"IpComparisonSymbol"`

	// 限制下载的vpc设置。
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitempty" name:"LimitVpc"`

	// 限制下载的ip设置
	LimitIp []*string `json:"LimitIp,omitempty" name:"LimitIp"`
}

type ModifyBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
	
	// NoLimit 不限制,内外网都可以下载； LimitOnlyIntranet 仅内网可下载； Customize 用户自定义vpc:ip可下载。 只有该值为 Customize 时，才可以设置 LimitVpc 和 LimitIp 。
	LimitType *string `json:"LimitType,omitempty" name:"LimitType"`

	// 该参数仅支持 In， 表示 LimitVpc 指定的vpc可以下载。默认为In。
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitempty" name:"VpcComparisonSymbol"`

	// In: 指定的ip可以下载； NotIn: 指定的ip不可以下载。 默认为In。
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitempty" name:"IpComparisonSymbol"`

	// 限制下载的vpc设置。
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitempty" name:"LimitVpc"`

	// 限制下载的ip设置
	LimitIp []*string `json:"LimitIp,omitempty" name:"LimitIp"`
}

func (r *ModifyBackupDownloadRestrictionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupDownloadRestrictionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LimitType")
	delete(f, "VpcComparisonSymbol")
	delete(f, "IpComparisonSymbol")
	delete(f, "LimitVpc")
	delete(f, "LimitIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupDownloadRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadRestrictionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBackupDownloadRestrictionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupDownloadRestrictionResponseParams `json:"Response"`
}

func (r *ModifyBackupDownloadRestrictionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupDownloadRestrictionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupEncryptionStatusRequestParams struct {
	// 实例ID，格式如：cdb-XXXX。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 设置实例新增的自动物理备份文件默认加密状态。可选值为 on或者off。
	EncryptionStatus *string `json:"EncryptionStatus,omitempty" name:"EncryptionStatus"`
}

type ModifyBackupEncryptionStatusRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cdb-XXXX。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 设置实例新增的自动物理备份文件默认加密状态。可选值为 on或者off。
	EncryptionStatus *string `json:"EncryptionStatus,omitempty" name:"EncryptionStatus"`
}

func (r *ModifyBackupEncryptionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupEncryptionStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EncryptionStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupEncryptionStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupEncryptionStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBackupEncryptionStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupEncryptionStatusResponseParams `json:"Response"`
}

func (r *ModifyBackupEncryptionStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupEncryptionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCDBProxyConnectionPoolRequestParams struct {
	// 数据库代理ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 是否开启连接池，true：开启连接池；
	//                              false：关闭连接池。
	OpenConnectionPool *bool `json:"OpenConnectionPool,omitempty" name:"OpenConnectionPool"`

	// 连接池类型，
	// 通过DescribeProxyConnectionPoolConf获取连接池类型值
	ConnectionPoolType *string `json:"ConnectionPoolType,omitempty" name:"ConnectionPoolType"`

	// 连接保留阈值：单位（秒）
	PoolConnectionTimeOut *int64 `json:"PoolConnectionTimeOut,omitempty" name:"PoolConnectionTimeOut"`
}

type ModifyCDBProxyConnectionPoolRequest struct {
	*tchttp.BaseRequest
	
	// 数据库代理ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 是否开启连接池，true：开启连接池；
	//                              false：关闭连接池。
	OpenConnectionPool *bool `json:"OpenConnectionPool,omitempty" name:"OpenConnectionPool"`

	// 连接池类型，
	// 通过DescribeProxyConnectionPoolConf获取连接池类型值
	ConnectionPoolType *string `json:"ConnectionPoolType,omitempty" name:"ConnectionPoolType"`

	// 连接保留阈值：单位（秒）
	PoolConnectionTimeOut *int64 `json:"PoolConnectionTimeOut,omitempty" name:"PoolConnectionTimeOut"`
}

func (r *ModifyCDBProxyConnectionPoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCDBProxyConnectionPoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyGroupId")
	delete(f, "OpenConnectionPool")
	delete(f, "ConnectionPoolType")
	delete(f, "PoolConnectionTimeOut")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCDBProxyConnectionPoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCDBProxyConnectionPoolResponseParams struct {
	// 异步处理ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCDBProxyConnectionPoolResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCDBProxyConnectionPoolResponseParams `json:"Response"`
}

func (r *ModifyCDBProxyConnectionPoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCDBProxyConnectionPoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCDBProxyDescRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库代理ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 数据库代理描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type ModifyCDBProxyDescRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库代理ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 数据库代理描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *ModifyCDBProxyDescRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCDBProxyDescRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProxyGroupId")
	delete(f, "Desc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCDBProxyDescRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCDBProxyDescResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCDBProxyDescResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCDBProxyDescResponseParams `json:"Response"`
}

func (r *ModifyCDBProxyDescResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCDBProxyDescResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCDBProxyVipVPortRequestParams struct {
	// 代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 私有网络ID
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 私有网络子网ID
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 目标IP
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// 目标端口
	DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`

	// 旧IP回收时间 单位小时
	ReleaseDuration *uint64 `json:"ReleaseDuration,omitempty" name:"ReleaseDuration"`
}

type ModifyCDBProxyVipVPortRequest struct {
	*tchttp.BaseRequest
	
	// 代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 私有网络ID
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 私有网络子网ID
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 目标IP
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// 目标端口
	DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`

	// 旧IP回收时间 单位小时
	ReleaseDuration *uint64 `json:"ReleaseDuration,omitempty" name:"ReleaseDuration"`
}

func (r *ModifyCDBProxyVipVPortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCDBProxyVipVPortRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyGroupId")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "DstIp")
	delete(f, "DstPort")
	delete(f, "ReleaseDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCDBProxyVipVPortRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCDBProxyVipVPortResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCDBProxyVipVPortResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCDBProxyVipVPortResponseParams `json:"Response"`
}

func (r *ModifyCDBProxyVipVPortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCDBProxyVipVPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNameRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 修改后的实例名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 修改后的实例名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyDBInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNameResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDBInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceNameResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceProjectRequestParams struct {
	// 实例 ID 数组，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 项目的 ID。
	NewProjectId *int64 `json:"NewProjectId,omitempty" name:"NewProjectId"`
}

type ModifyDBInstanceProjectRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 数组，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 项目的 ID。
	NewProjectId *int64 `json:"NewProjectId,omitempty" name:"NewProjectId"`
}

func (r *ModifyDBInstanceProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "NewProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceProjectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDBInstanceProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceProjectResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 要修改的安全组 ID 列表，一个或者多个安全组 ID 组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 当传入只读实例ID时，默认操作的是对应只读组的安全组。如果需要操作只读实例ID的安全组， 需要将该入参置为True
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitempty" name:"ForReadonlyInstance"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 要修改的安全组 ID 列表，一个或者多个安全组 ID 组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 当传入只读实例ID时，默认操作的是对应只读组的安全组。如果需要操作只读实例ID的安全组， 需要将该入参置为True
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitempty" name:"ForReadonlyInstance"`
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
	delete(f, "ForReadonlyInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyDBInstanceVipVportRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c2nl9rpv 或者 cdbrg-c3nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
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

type ModifyDBInstanceVipVportRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c2nl9rpv 或者 cdbrg-c3nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceVipVportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DstIp")
	delete(f, "DstPort")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "ReleaseDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceVipVportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceVipVportResponseParams struct {
	// 异步任务ID。(该返回字段目前已废弃)
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDBInstanceVipVportResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceVipVportResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceVipVportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceVipVportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamRequestParams struct {
	// 实例短 ID 列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 要修改的参数列表。每一个元素是 Name 和 CurrentValue 的组合。Name 是参数名，CurrentValue 是要修改成的值。
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList"`

	// 模板id，ParamList和TemplateId必须至少传其中之一
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 执行参数调整任务的方式，默认为 0。支持值包括：0 - 立刻执行，1 - 时间窗执行；当该值为 1 时，每次只能传一个实例（InstanceIds数量为1）
	WaitSwitch *int64 `json:"WaitSwitch,omitempty" name:"WaitSwitch"`

	// 参数是否同步到主实例下的只读实例。true 为不同步，false 为同步。默认为 false。
	NotSyncRo *bool `json:"NotSyncRo,omitempty" name:"NotSyncRo"`

	// 参数是否同步到主实例下的灾备实例。true 为不同步，false 为同步。默认为 false。
	NotSyncDr *bool `json:"NotSyncDr,omitempty" name:"NotSyncDr"`
}

type ModifyInstanceParamRequest struct {
	*tchttp.BaseRequest
	
	// 实例短 ID 列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 要修改的参数列表。每一个元素是 Name 和 CurrentValue 的组合。Name 是参数名，CurrentValue 是要修改成的值。
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList"`

	// 模板id，ParamList和TemplateId必须至少传其中之一
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 执行参数调整任务的方式，默认为 0。支持值包括：0 - 立刻执行，1 - 时间窗执行；当该值为 1 时，每次只能传一个实例（InstanceIds数量为1）
	WaitSwitch *int64 `json:"WaitSwitch,omitempty" name:"WaitSwitch"`

	// 参数是否同步到主实例下的只读实例。true 为不同步，false 为同步。默认为 false。
	NotSyncRo *bool `json:"NotSyncRo,omitempty" name:"NotSyncRo"`

	// 参数是否同步到主实例下的灾备实例。true 为不同步，false 为同步。默认为 false。
	NotSyncDr *bool `json:"NotSyncDr,omitempty" name:"NotSyncDr"`
}

func (r *ModifyInstanceParamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ParamList")
	delete(f, "TemplateId")
	delete(f, "WaitSwitch")
	delete(f, "NotSyncRo")
	delete(f, "NotSyncDr")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceParamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamResponseParams struct {
	// 异步任务 ID，可用于查询任务进度。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyInstanceParamResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceParamResponseParams `json:"Response"`
}

func (r *ModifyInstanceParamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancePasswordComplexityRequestParams struct {
	// 实例短 ID 列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 要修改的参数列表。每一个元素是Name和CurrentValue的组合。Name是参数名，CurrentValue是要修改成的值。8.0版本Name支持范围：["validate_password.policy","validate_password.length","validate_password.mixed_case_count","validate_password.number_count","validate_password.special_char_count"],5.6和5.7版本支持范围：["validate_password_policy","validate_password_length","validate_password_mixed_case_count","validate_password_number_count","validate_password_special_char_count"]
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList"`
}

type ModifyInstancePasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// 实例短 ID 列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 要修改的参数列表。每一个元素是Name和CurrentValue的组合。Name是参数名，CurrentValue是要修改成的值。8.0版本Name支持范围：["validate_password.policy","validate_password.length","validate_password.mixed_case_count","validate_password.number_count","validate_password.special_char_count"],5.6和5.7版本支持范围：["validate_password_policy","validate_password_length","validate_password_mixed_case_count","validate_password_number_count","validate_password_special_char_count"]
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList"`
}

func (r *ModifyInstancePasswordComplexityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancePasswordComplexityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancePasswordComplexityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancePasswordComplexityResponseParams struct {
	// 异步任务 ID，可用于查询任务进度。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyInstancePasswordComplexityResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancePasswordComplexityResponseParams `json:"Response"`
}

func (r *ModifyInstancePasswordComplexityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancePasswordComplexityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceTagRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 要增加或修改的标签。
	ReplaceTags []*TagInfo `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// 要删除的标签。
	DeleteTags []*TagInfo `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

type ModifyInstanceTagRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 要增加或修改的标签。
	ReplaceTags []*TagInfo `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// 要删除的标签。
	DeleteTags []*TagInfo `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

func (r *ModifyInstanceTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ReplaceTags")
	delete(f, "DeleteTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyInstanceTagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceTagResponseParams `json:"Response"`
}

func (r *ModifyInstanceTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLocalBinlogConfigRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 本地binlog保留时长，可取值范围：[72,168]，当实例存在灾备实例时，可取值范围：[120,168]。
	SaveHours *int64 `json:"SaveHours,omitempty" name:"SaveHours"`

	// 本地binlog空间使用率，可取值范围：[30,50]。
	MaxUsage *int64 `json:"MaxUsage,omitempty" name:"MaxUsage"`
}

type ModifyLocalBinlogConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 本地binlog保留时长，可取值范围：[72,168]，当实例存在灾备实例时，可取值范围：[120,168]。
	SaveHours *int64 `json:"SaveHours,omitempty" name:"SaveHours"`

	// 本地binlog空间使用率，可取值范围：[30,50]。
	MaxUsage *int64 `json:"MaxUsage,omitempty" name:"MaxUsage"`
}

func (r *ModifyLocalBinlogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLocalBinlogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SaveHours")
	delete(f, "MaxUsage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLocalBinlogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLocalBinlogConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLocalBinlogConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLocalBinlogConfigResponseParams `json:"Response"`
}

func (r *ModifyLocalBinlogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLocalBinlogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNameOrDescByDpIdRequestParams struct {
	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// 置放群组名称，最长不能超过60个字符。置放群组名和置放群组描述不能都为空。
	DeployGroupName *string `json:"DeployGroupName,omitempty" name:"DeployGroupName"`

	// 置放群组描述，最长不能超过200个字符。置放群组名和置放群组描述不能都为空。
	Description *string `json:"Description,omitempty" name:"Description"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNameOrDescByDpIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployGroupId")
	delete(f, "DeployGroupName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNameOrDescByDpIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNameOrDescByDpIdResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyNameOrDescByDpIdResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNameOrDescByDpIdResponseParams `json:"Response"`
}

func (r *ModifyNameOrDescByDpIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNameOrDescByDpIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyParamTemplateRequestParams struct {
	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名称，长度不超过64。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述，长度不超过255。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 参数列表。
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList"`
}

type ModifyParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名称，长度不超过64。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述，长度不超过255。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 参数列表。
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList"`
}

func (r *ModifyParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyParamTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyParamTemplateResponseParams `json:"Response"`
}

func (r *ModifyParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRemoteBackupConfigRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 异地数据备份开关，off - 关闭异地备份，on-开启异地备份
	RemoteBackupSave *string `json:"RemoteBackupSave,omitempty" name:"RemoteBackupSave"`

	// 异地日志备份开关，off - 关闭异地备份，on-开启异地备份，只有在参数RemoteBackupSave为on时，RemoteBinlogSave参数才可设置为on
	RemoteBinlogSave *string `json:"RemoteBinlogSave,omitempty" name:"RemoteBinlogSave"`

	// 用户设置异地备份地域列表
	RemoteRegion []*string `json:"RemoteRegion,omitempty" name:"RemoteRegion"`

	// 异地备份保留时间，单位为天
	ExpireDays *int64 `json:"ExpireDays,omitempty" name:"ExpireDays"`
}

type ModifyRemoteBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 异地数据备份开关，off - 关闭异地备份，on-开启异地备份
	RemoteBackupSave *string `json:"RemoteBackupSave,omitempty" name:"RemoteBackupSave"`

	// 异地日志备份开关，off - 关闭异地备份，on-开启异地备份，只有在参数RemoteBackupSave为on时，RemoteBinlogSave参数才可设置为on
	RemoteBinlogSave *string `json:"RemoteBinlogSave,omitempty" name:"RemoteBinlogSave"`

	// 用户设置异地备份地域列表
	RemoteRegion []*string `json:"RemoteRegion,omitempty" name:"RemoteRegion"`

	// 异地备份保留时间，单位为天
	ExpireDays *int64 `json:"ExpireDays,omitempty" name:"ExpireDays"`
}

func (r *ModifyRemoteBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRemoteBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RemoteBackupSave")
	delete(f, "RemoteBinlogSave")
	delete(f, "RemoteRegion")
	delete(f, "ExpireDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRemoteBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRemoteBackupConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRemoteBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRemoteBackupConfigResponseParams `json:"Response"`
}

func (r *ModifyRemoteBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRemoteBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoGroupInfoRequestParams struct {
	// RO 组的 ID。
	RoGroupId *string `json:"RoGroupId,omitempty" name:"RoGroupId"`

	// RO 组的详细信息。
	RoGroupInfo *RoGroupAttr `json:"RoGroupInfo,omitempty" name:"RoGroupInfo"`

	// RO 组内实例的权重。若修改 RO 组的权重模式为用户自定义模式（custom），则必须设置该参数，且需要设置每个 RO 实例的权重值。
	RoWeightValues []*RoWeightValue `json:"RoWeightValues,omitempty" name:"RoWeightValues"`

	// 是否重新均衡 RO 组内的 RO 实例的负载。支持值包括：1 - 重新均衡负载；0 - 不重新均衡负载。默认值为 0。注意，设置为重新均衡负载时，RO 组内 RO 实例会有一次数据库连接瞬断，请确保应用程序能重连数据库。
	IsBalanceRoLoad *int64 `json:"IsBalanceRoLoad,omitempty" name:"IsBalanceRoLoad"`

	// 废弃参数，无意义。
	ReplicationDelayTime *int64 `json:"ReplicationDelayTime,omitempty" name:"ReplicationDelayTime"`
}

type ModifyRoGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// RO 组的 ID。
	RoGroupId *string `json:"RoGroupId,omitempty" name:"RoGroupId"`

	// RO 组的详细信息。
	RoGroupInfo *RoGroupAttr `json:"RoGroupInfo,omitempty" name:"RoGroupInfo"`

	// RO 组内实例的权重。若修改 RO 组的权重模式为用户自定义模式（custom），则必须设置该参数，且需要设置每个 RO 实例的权重值。
	RoWeightValues []*RoWeightValue `json:"RoWeightValues,omitempty" name:"RoWeightValues"`

	// 是否重新均衡 RO 组内的 RO 实例的负载。支持值包括：1 - 重新均衡负载；0 - 不重新均衡负载。默认值为 0。注意，设置为重新均衡负载时，RO 组内 RO 实例会有一次数据库连接瞬断，请确保应用程序能重连数据库。
	IsBalanceRoLoad *int64 `json:"IsBalanceRoLoad,omitempty" name:"IsBalanceRoLoad"`

	// 废弃参数，无意义。
	ReplicationDelayTime *int64 `json:"ReplicationDelayTime,omitempty" name:"ReplicationDelayTime"`
}

func (r *ModifyRoGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoGroupId")
	delete(f, "RoGroupInfo")
	delete(f, "RoWeightValues")
	delete(f, "IsBalanceRoLoad")
	delete(f, "ReplicationDelayTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoGroupInfoResponseParams struct {
	// 异步任务 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRoGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRoGroupInfoResponseParams `json:"Response"`
}

func (r *ModifyRoGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTimeWindowRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 修改后的可维护时间段，其中每一个时间段的格式形如：10:00-12:00；起止时间按半个小时对齐；最短半个小时，最长三个小时；最多设置两个时间段；起止时间范围为：[00:00, 24:00]。
	TimeRanges []*string `json:"TimeRanges,omitempty" name:"TimeRanges"`

	// 指定修改哪一天的客户时间段，可能的取值为：monday，tuesday，wednesday，thursday，friday，saturday，sunday。如果不指定该值或者为空，则默认一周七天都修改。
	Weekdays []*string `json:"Weekdays,omitempty" name:"Weekdays"`

	// 数据延迟阈值，仅对主实例和灾备实例有效，不传默认修改为10
	MaxDelayTime *uint64 `json:"MaxDelayTime,omitempty" name:"MaxDelayTime"`
}

type ModifyTimeWindowRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 修改后的可维护时间段，其中每一个时间段的格式形如：10:00-12:00；起止时间按半个小时对齐；最短半个小时，最长三个小时；最多设置两个时间段；起止时间范围为：[00:00, 24:00]。
	TimeRanges []*string `json:"TimeRanges,omitempty" name:"TimeRanges"`

	// 指定修改哪一天的客户时间段，可能的取值为：monday，tuesday，wednesday，thursday，friday，saturday，sunday。如果不指定该值或者为空，则默认一周七天都修改。
	Weekdays []*string `json:"Weekdays,omitempty" name:"Weekdays"`

	// 数据延迟阈值，仅对主实例和灾备实例有效，不传默认修改为10
	MaxDelayTime *uint64 `json:"MaxDelayTime,omitempty" name:"MaxDelayTime"`
}

func (r *ModifyTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTimeWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TimeRanges")
	delete(f, "Weekdays")
	delete(f, "MaxDelayTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTimeWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTimeWindowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTimeWindowResponseParams `json:"Response"`
}

func (r *ModifyTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTimeWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineIsolatedInstancesRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type OfflineIsolatedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *OfflineIsolatedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineIsolatedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OfflineIsolatedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineIsolatedInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OfflineIsolatedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *OfflineIsolatedInstancesResponseParams `json:"Response"`
}

func (r *OfflineIsolatedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineIsolatedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenAuditServiceRequestParams struct {
	// CDB实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 审计日志保存时长。支持值包括：
	// 7 - 一周
	// 30 - 一个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年；
	LogExpireDay *uint64 `json:"LogExpireDay,omitempty" name:"LogExpireDay"`

	// 高频审计日志保存时长。支持值包括：
	// 7 - 一周
	// 30 - 一个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年；
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitempty" name:"HighLogExpireDay"`
}

type OpenAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// CDB实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 审计日志保存时长。支持值包括：
	// 7 - 一周
	// 30 - 一个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年；
	LogExpireDay *uint64 `json:"LogExpireDay,omitempty" name:"LogExpireDay"`

	// 高频审计日志保存时长。支持值包括：
	// 7 - 一周
	// 30 - 一个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年；
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitempty" name:"HighLogExpireDay"`
}

func (r *OpenAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	delete(f, "HighLogExpireDay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenAuditServiceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OpenAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *OpenAuditServiceResponseParams `json:"Response"`
}

func (r *OpenAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenDBInstanceEncryptionRequestParams struct {
	// 云数据库实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户自定义密钥ID，CMK唯一标识符。该值为空时，将使用腾讯云自动生成的密钥KMS-CDB。
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 用户自定义密钥的存储地域。如：ap-guangzhou 。KeyId不为空时，该参数必填。
	KeyRegion *string `json:"KeyRegion,omitempty" name:"KeyRegion"`
}

type OpenDBInstanceEncryptionRequest struct {
	*tchttp.BaseRequest
	
	// 云数据库实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户自定义密钥ID，CMK唯一标识符。该值为空时，将使用腾讯云自动生成的密钥KMS-CDB。
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 用户自定义密钥的存储地域。如：ap-guangzhou 。KeyId不为空时，该参数必填。
	KeyRegion *string `json:"KeyRegion,omitempty" name:"KeyRegion"`
}

func (r *OpenDBInstanceEncryptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBInstanceEncryptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KeyId")
	delete(f, "KeyRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenDBInstanceEncryptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenDBInstanceEncryptionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OpenDBInstanceEncryptionResponse struct {
	*tchttp.BaseResponse
	Response *OpenDBInstanceEncryptionResponseParams `json:"Response"`
}

func (r *OpenDBInstanceEncryptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBInstanceEncryptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenDBInstanceGTIDRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBInstanceGTIDRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenDBInstanceGTIDRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenDBInstanceGTIDResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OpenDBInstanceGTIDResponse struct {
	*tchttp.BaseResponse
	Response *OpenDBInstanceGTIDResponseParams `json:"Response"`
}

func (r *OpenDBInstanceGTIDResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBInstanceGTIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenWanServiceRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenWanServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenWanServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenWanServiceResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OpenWanServiceResponse struct {
	*tchttp.BaseResponse
	Response *OpenWanServiceResponseParams `json:"Response"`
}

func (r *OpenWanServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// 规则描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`
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

	// 参数模板类型
	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
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

	// 参数类型：integer，enum，float，string，func
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
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`

	// 参数是公式类型时，该字段有效，表示公式类型最大值
	MaxFunc *string `json:"MaxFunc,omitempty" name:"MaxFunc"`

	// 参数是公式类型时，该字段有效，表示公式类型最小值
	MinFunc *string `json:"MinFunc,omitempty" name:"MinFunc"`
}

type PoolConf struct {
	// 连接池类型：SessionConnectionPool（会话级别连接池
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectionPoolType *string `json:"ConnectionPoolType,omitempty" name:"ConnectionPoolType"`

	// 最大可保持连接阈值：单位（秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxPoolConnectionTimeOut *int64 `json:"MaxPoolConnectionTimeOut,omitempty" name:"MaxPoolConnectionTimeOut"`

	// 最小可保持连接阈值：单位（秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinPoolConnectionTimeOut *int64 `json:"MinPoolConnectionTimeOut,omitempty" name:"MinPoolConnectionTimeOut"`
}

type ProxyGroup struct {
	// 代理基本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaseGroup *BaseGroupInfo `json:"BaseGroup,omitempty" name:"BaseGroup"`

	// 代理地址信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address []*Address `json:"Address,omitempty" name:"Address"`

	// 代理连接池信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectionPoolInfo *ConnectionPoolInfo `json:"ConnectionPoolInfo,omitempty" name:"ConnectionPoolInfo"`

	// 代理节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyNode []*ProxyNodeInfo `json:"ProxyNode,omitempty" name:"ProxyNode"`

	// 代理路由信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RWInstInfo *RWInfo `json:"RWInstInfo,omitempty" name:"RWInstInfo"`
}

type ProxyGroups struct {
	// 代理基本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaseGroup *BaseGroupInfo `json:"BaseGroup,omitempty" name:"BaseGroup"`

	// 代理地址信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address []*Address `json:"Address,omitempty" name:"Address"`

	// 代理连接池信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectionPoolInfo *ConnectionPoolInfo `json:"ConnectionPoolInfo,omitempty" name:"ConnectionPoolInfo"`

	// 代理节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyNode []*ProxyNodeInfo `json:"ProxyNode,omitempty" name:"ProxyNode"`

	// 代理路由信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RWInstInfo *RWInfos `json:"RWInstInfo,omitempty" name:"RWInstInfo"`
}

type ProxyNodeInfo struct {
	// 代理节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyNodeId *string `json:"ProxyNodeId,omitempty" name:"ProxyNodeId"`

	// 节点当前连接数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyNodeConnections *uint64 `json:"ProxyNodeConnections,omitempty" name:"ProxyNodeConnections"`

	// cup
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyNodeCpu *uint64 `json:"ProxyNodeCpu,omitempty" name:"ProxyNodeCpu"`

	// 内存
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyNodeMem *uint64 `json:"ProxyNodeMem,omitempty" name:"ProxyNodeMem"`

	// 节点状态：
	// init（申请中）
	// online（运行中）
	// offline（离线中）
	// destroy（已销毁）
	// recovering（故障恢复中）
	// error（节点故障）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyStatus *string `json:"ProxyStatus,omitempty" name:"ProxyStatus"`
}

// Predefined struct for user
type QueryCDBProxyRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 代理ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`
}

type QueryCDBProxyRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 代理ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`
}

func (r *QueryCDBProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCDBProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProxyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCDBProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCDBProxyResponseParams struct {
	// 代理数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 代理信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyGroup []*ProxyGroups `json:"ProxyGroup,omitempty" name:"ProxyGroup"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryCDBProxyResponse struct {
	*tchttp.BaseResponse
	Response *QueryCDBProxyResponseParams `json:"Response"`
}

func (r *QueryCDBProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCDBProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RWInfo struct {
	// 代理实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstCount *uint64 `json:"InstCount,omitempty" name:"InstCount"`

	// 权重分配模式；
	// 系统自动分配："system"， 自定义："custom"
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeightMode *string `json:"WeightMode,omitempty" name:"WeightMode"`

	// 是否开启延迟剔除
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsKickOut *bool `json:"IsKickOut,omitempty" name:"IsKickOut"`

	// 最小保留数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinCount *uint64 `json:"MinCount,omitempty" name:"MinCount"`

	// 延迟剔除阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDelay *uint64 `json:"MaxDelay,omitempty" name:"MaxDelay"`

	// 是否开启故障转移
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailOver *bool `json:"FailOver,omitempty" name:"FailOver"`

	// 是否自动添加RO
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoAddRo *bool `json:"AutoAddRo,omitempty" name:"AutoAddRo"`

	// 代理实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RWInstInfo *RWInstanceInfo `json:"RWInstInfo,omitempty" name:"RWInstInfo"`
}

type RWInfos struct {
	// 代理实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstCount *uint64 `json:"InstCount,omitempty" name:"InstCount"`

	// 权重分配模式；
	// 系统自动分配："system"， 自定义："custom"
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeightMode *string `json:"WeightMode,omitempty" name:"WeightMode"`

	// 是否开启延迟剔除
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsKickOut *bool `json:"IsKickOut,omitempty" name:"IsKickOut"`

	// 最小保留数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinCount *uint64 `json:"MinCount,omitempty" name:"MinCount"`

	// 延迟剔除阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDelay *uint64 `json:"MaxDelay,omitempty" name:"MaxDelay"`

	// 是否开启故障转移
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailOver *bool `json:"FailOver,omitempty" name:"FailOver"`

	// 是否自动添加RO
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoAddRo *bool `json:"AutoAddRo,omitempty" name:"AutoAddRo"`

	// 代理实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RWInstInfo []*RWInstanceInfo `json:"RWInstInfo,omitempty" name:"RWInstInfo"`
}

type RWInstanceInfo struct {

}

// Predefined struct for user
type ReleaseIsolatedDBInstancesRequestParams struct {
	// 实例 ID 数组，单个实例 ID 格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type ReleaseIsolatedDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 数组，单个实例 ID 格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *ReleaseIsolatedDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseIsolatedDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseIsolatedDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseIsolatedDBInstancesResponseParams struct {
	// 解隔离操作的结果集。
	Items []*ReleaseResult `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReleaseIsolatedDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseIsolatedDBInstancesResponseParams `json:"Response"`
}

func (r *ReleaseIsolatedDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// Predefined struct for user
type ReloadBalanceProxyNodeRequestParams struct {
	// 代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 代理组地址ID
	ProxyAddressId *string `json:"ProxyAddressId,omitempty" name:"ProxyAddressId"`
}

type ReloadBalanceProxyNodeRequest struct {
	*tchttp.BaseRequest
	
	// 代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 代理组地址ID
	ProxyAddressId *string `json:"ProxyAddressId,omitempty" name:"ProxyAddressId"`
}

func (r *ReloadBalanceProxyNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReloadBalanceProxyNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyGroupId")
	delete(f, "ProxyAddressId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReloadBalanceProxyNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReloadBalanceProxyNodeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReloadBalanceProxyNodeResponse struct {
	*tchttp.BaseResponse
	Response *ReloadBalanceProxyNodeResponseParams `json:"Response"`
}

func (r *ReloadBalanceProxyNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReloadBalanceProxyNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoteBackupInfo struct {
	// 异地备份子任务的ID
	SubBackupId []*int64 `json:"SubBackupId,omitempty" name:"SubBackupId"`

	// 异地备份所在地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 备份任务状态。可能的值有 "SUCCESS": 备份成功， "FAILED": 备份失败， "RUNNING": 备份进行中。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 异地备份任务的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 异地备份任务的结束时间
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// 下载地址
	Url *string `json:"Url,omitempty" name:"Url"`
}

// Predefined struct for user
type RenewDBInstanceRequestParams struct {
	// 待续费的实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872)。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 续费时长，单位：月，可选值包括 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]。
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 如果需要将按量计费实例续费为包年包月的实例，该入参的值需要指定为 "PREPAID" 。
	ModifyPayType *string `json:"ModifyPayType,omitempty" name:"ModifyPayType"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TimeSpan")
	delete(f, "ModifyPayType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDBInstanceResponseParams struct {
	// 订单 ID。
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RenewDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RenewDBInstanceResponseParams `json:"Response"`
}

func (r *RenewDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetRootAccountRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ResetRootAccountRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ResetRootAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetRootAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetRootAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetRootAccountResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetRootAccountResponse struct {
	*tchttp.BaseResponse
	Response *ResetRootAccountResponseParams `json:"Response"`
}

func (r *ResetRootAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetRootAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartDBInstancesRequestParams struct {
	// 实例 ID 数组，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type RestartDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 数组，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartDBInstancesResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	RoInstances []*RoInstanceInfo `json:"RoInstances,omitempty" name:"RoInstances"`

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

	// 延迟复制时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelayReplicationTime *int64 `json:"DelayReplicationTime,omitempty" name:"DelayReplicationTime"`
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

	// 延迟复制时间。
	ReplicationDelayTime *int64 `json:"ReplicationDelayTime,omitempty" name:"ReplicationDelayTime"`
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

	// RO实例状态，可能返回值：0-创建中，1-运行中，3-异地RO（仅在使用DescribeDBInstances查询主实例信息时，返回值中异地RO的状态恒等于3，其他场景下无此值），4-删除中
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

	// RO实例数据库引擎版本，可能返回值：5.1、5.5、5.6、5.7、8.0
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

	// 回档策略。可选值为：table、db、full；默认值为full。table - 极速回档模式，仅导入所选中表级别的备份和binlog，如有跨表操作，且关联表未被同时选中，将会导致回档失败，该模式下参数Databases必须为空；db - 快速模式，仅导入所选中库级别的备份和binlog，如有跨库操作，且关联库未被同时选中，将会导致回档失败；full - 普通回档模式，将导入整个实例的备份和binlog，速度较慢。
	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`

	// 数据库回档时间，时间格式为：yyyy-mm-dd hh:mm:ss
	RollbackTime *string `json:"RollbackTime,omitempty" name:"RollbackTime"`

	// 待回档的数据库信息，表示整库回档
	// 注意：此字段可能返回 null，表示取不到有效值。
	Databases []*RollbackDBName `json:"Databases,omitempty" name:"Databases"`

	// 待回档的数据库表信息，表示按表回档
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tables []*RollbackTables `json:"Tables,omitempty" name:"Tables"`
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
	Table []*RollbackTableName `json:"Table,omitempty" name:"Table"`
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
	Detail []*RollbackInstancesInfo `json:"Detail,omitempty" name:"Detail"`
}

type RollbackTimeRange struct {
	// 实例可回档开始时间，时间格式：2016-10-29 01:06:04
	Begin *string `json:"Begin,omitempty" name:"Begin"`

	// 实例可回档结束时间，时间格式：2016-11-02 11:44:47
	End *string `json:"End,omitempty" name:"End"`
}

type Rule struct {
	// 划分上限
	// 注意：此字段可能返回 null，表示取不到有效值。
	LessThan *uint64 `json:"LessThan,omitempty" name:"LessThan"`

	// 权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
}

type SecurityGroup struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 创建时间，时间格式：yyyy-mm-dd hh:mm:ss
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 入站规则
	Inbound []*Inbound `json:"Inbound,omitempty" name:"Inbound"`

	// 出站规则
	Outbound []*Outbound `json:"Outbound,omitempty" name:"Outbound"`

	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组名称
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// 安全组备注
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`
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

	// Sql的执行时长（秒）。
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

	// 锁时长（秒）。
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

// Predefined struct for user
type StartBatchRollbackRequestParams struct {
	// 用于回档的实例详情信息。
	Instances []*RollbackInstancesInfo `json:"Instances,omitempty" name:"Instances"`
}

type StartBatchRollbackRequest struct {
	*tchttp.BaseRequest
	
	// 用于回档的实例详情信息。
	Instances []*RollbackInstancesInfo `json:"Instances,omitempty" name:"Instances"`
}

func (r *StartBatchRollbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBatchRollbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartBatchRollbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartBatchRollbackResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartBatchRollbackResponse struct {
	*tchttp.BaseResponse
	Response *StartBatchRollbackResponseParams `json:"Response"`
}

func (r *StartBatchRollbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBatchRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartReplicationRequestParams struct {
	// 实例 ID。仅支持只读实例。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type StartReplicationRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。仅支持只读实例。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *StartReplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartReplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartReplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartReplicationResponseParams struct {
	// 异步任务 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartReplicationResponse struct {
	*tchttp.BaseResponse
	Response *StartReplicationResponseParams `json:"Response"`
}

func (r *StartReplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartReplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopDBImportJobRequestParams struct {
	// 异步任务的请求 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopDBImportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AsyncRequestId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopDBImportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopDBImportJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopDBImportJobResponse struct {
	*tchttp.BaseResponse
	Response *StopDBImportJobResponseParams `json:"Response"`
}

func (r *StopDBImportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopDBImportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopReplicationRequestParams struct {
	// 实例 ID。仅支持只读实例。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type StopReplicationRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。仅支持只读实例。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *StopReplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopReplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopReplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopReplicationResponseParams struct {
	// 异步任务 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopReplicationResponse struct {
	*tchttp.BaseResponse
	Response *StopReplicationResponseParams `json:"Response"`
}

func (r *StopReplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopReplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopRollbackRequestParams struct {
	// 撤销回档任务对应的实例Id。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRollbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopRollbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopRollbackResponseParams struct {
	// 执行请求的异步任务ID
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopRollbackResponse struct {
	*tchttp.BaseResponse
	Response *StopRollbackResponseParams `json:"Response"`
}

func (r *StopRollbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchCDBProxyRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库代理ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`
}

type SwitchCDBProxyRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库代理ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`
}

func (r *SwitchCDBProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchCDBProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProxyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchCDBProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchCDBProxyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwitchCDBProxyResponse struct {
	*tchttp.BaseResponse
	Response *SwitchCDBProxyResponseParams `json:"Response"`
}

func (r *SwitchCDBProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchCDBProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDBInstanceMasterSlaveRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 目标从实例。可选值："first" - 第一备机；"second" - 第二备机。默认值为 "first"，仅多可用区实例支持设置为 "second"。
	DstSlave *string `json:"DstSlave,omitempty" name:"DstSlave"`

	// 是否强制切换。默认为 False。注意，若设置强制切换为 True，实例存在丢失数据的风险，请谨慎使用。
	ForceSwitch *bool `json:"ForceSwitch,omitempty" name:"ForceSwitch"`

	// 是否时间窗内切换。默认为 False，即不在时间窗内切换。注意，如果设置了 ForceSwitch 参数为 True，则该参数不生效。
	WaitSwitch *bool `json:"WaitSwitch,omitempty" name:"WaitSwitch"`
}

type SwitchDBInstanceMasterSlaveRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 目标从实例。可选值："first" - 第一备机；"second" - 第二备机。默认值为 "first"，仅多可用区实例支持设置为 "second"。
	DstSlave *string `json:"DstSlave,omitempty" name:"DstSlave"`

	// 是否强制切换。默认为 False。注意，若设置强制切换为 True，实例存在丢失数据的风险，请谨慎使用。
	ForceSwitch *bool `json:"ForceSwitch,omitempty" name:"ForceSwitch"`

	// 是否时间窗内切换。默认为 False，即不在时间窗内切换。注意，如果设置了 ForceSwitch 参数为 True，则该参数不生效。
	WaitSwitch *bool `json:"WaitSwitch,omitempty" name:"WaitSwitch"`
}

func (r *SwitchDBInstanceMasterSlaveRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDBInstanceMasterSlaveRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DstSlave")
	delete(f, "ForceSwitch")
	delete(f, "WaitSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchDBInstanceMasterSlaveRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDBInstanceMasterSlaveResponseParams struct {
	// 异步任务 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwitchDBInstanceMasterSlaveResponse struct {
	*tchttp.BaseResponse
	Response *SwitchDBInstanceMasterSlaveResponseParams `json:"Response"`
}

func (r *SwitchDBInstanceMasterSlaveResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDBInstanceMasterSlaveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDrInstanceToMasterRequestParams struct {
	// 灾备实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type SwitchDrInstanceToMasterRequest struct {
	*tchttp.BaseRequest
	
	// 灾备实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *SwitchDrInstanceToMasterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDrInstanceToMasterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchDrInstanceToMasterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDrInstanceToMasterResponseParams struct {
	// 异步任务的请求ID，可使用此ID查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwitchDrInstanceToMasterResponse struct {
	*tchttp.BaseResponse
	Response *SwitchDrInstanceToMasterResponseParams `json:"Response"`
}

func (r *SwitchDrInstanceToMasterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDrInstanceToMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchForUpgradeRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchForUpgradeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchForUpgradeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchForUpgradeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwitchForUpgradeResponse struct {
	*tchttp.BaseResponse
	Response *SwitchForUpgradeResponseParams `json:"Response"`
}

func (r *SwitchForUpgradeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TagInfo struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue []*string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagInfoItem struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
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
	Tags []*TagInfoUnit `json:"Tags,omitempty" name:"Tags"`
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
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 异步任务的请求 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

// Predefined struct for user
type UpgradeCDBProxyVersionRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库代理ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 数据库代理当前版本
	SrcProxyVersion *string `json:"SrcProxyVersion,omitempty" name:"SrcProxyVersion"`

	// 数据库代理升级版本
	DstProxyVersion *string `json:"DstProxyVersion,omitempty" name:"DstProxyVersion"`

	// 升级时间 ：nowTime（升级完成时）timeWindow（实例维护时间）
	UpgradeTime *string `json:"UpgradeTime,omitempty" name:"UpgradeTime"`
}

type UpgradeCDBProxyVersionRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库代理ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 数据库代理当前版本
	SrcProxyVersion *string `json:"SrcProxyVersion,omitempty" name:"SrcProxyVersion"`

	// 数据库代理升级版本
	DstProxyVersion *string `json:"DstProxyVersion,omitempty" name:"DstProxyVersion"`

	// 升级时间 ：nowTime（升级完成时）timeWindow（实例维护时间）
	UpgradeTime *string `json:"UpgradeTime,omitempty" name:"UpgradeTime"`
}

func (r *UpgradeCDBProxyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeCDBProxyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProxyGroupId")
	delete(f, "SrcProxyVersion")
	delete(f, "DstProxyVersion")
	delete(f, "UpgradeTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeCDBProxyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeCDBProxyVersionResponseParams struct {
	// 异步处理ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpgradeCDBProxyVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeCDBProxyVersionResponseParams `json:"Response"`
}

func (r *UpgradeCDBProxyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeCDBProxyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceEngineVersionRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主实例数据库引擎版本，支持值包括：5.6 和 5.7。
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 切换访问新实例的方式，默认为 0。支持值包括：0 - 立刻切换，1 - 时间窗切换；当该值为 1 时，升级中过程中，切换访问新实例的流程将会在时间窗内进行，或者用户主动调用接口 [切换访问新实例](https://cloud.tencent.com/document/product/236/15864) 触发该流程。
	WaitSwitch *int64 `json:"WaitSwitch,omitempty" name:"WaitSwitch"`

	// 是否是内核子版本升级，支持的值：1 - 升级内核子版本；0 - 升级数据库引擎版本。
	UpgradeSubversion *int64 `json:"UpgradeSubversion,omitempty" name:"UpgradeSubversion"`

	// 延迟阈值。取值范围1~10
	MaxDelayTime *int64 `json:"MaxDelayTime,omitempty" name:"MaxDelayTime"`
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

	// 延迟阈值。取值范围1~10
	MaxDelayTime *int64 `json:"MaxDelayTime,omitempty" name:"MaxDelayTime"`
}

func (r *UpgradeDBInstanceEngineVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceEngineVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EngineVersion")
	delete(f, "WaitSwitch")
	delete(f, "UpgradeSubversion")
	delete(f, "MaxDelayTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceEngineVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceEngineVersionResponseParams struct {
	// 异步任务 ID，可使用 [查询异步任务的执行结果](https://cloud.tencent.com/document/api/236/20410) 获取其执行情况。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpgradeDBInstanceEngineVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDBInstanceEngineVersionResponseParams `json:"Response"`
}

func (r *UpgradeDBInstanceEngineVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceEngineVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceRequestParams struct {
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

	// 实例隔离类型。支持值包括： "UNIVERSAL" - 通用型实例， "EXCLUSIVE" - 独享型实例， "BASIC" - 基础版实例。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 升级后的实例cpu核数， 如果不传将根据 Memory 指定的内存值自动填充对应的cpu值。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 是否极速变配。0-普通升级，1-极速变配,，2 极速优先。选择极速变配会根据资源状况校验是否可以进行极速变配，满足条件则进行极速变配，不满足条件会返回报错信息。
	FastUpgrade *int64 `json:"FastUpgrade,omitempty" name:"FastUpgrade"`

	// 延迟阈值。取值范围1~10，默认值为10。
	MaxDelayTime *int64 `json:"MaxDelayTime,omitempty" name:"MaxDelayTime"`

	// 是否跨区迁移。0-普通迁移，1-跨区迁移，默认值为0。该值为1时支持变更实例主节点可用区。
	CrossCluster *int64 `json:"CrossCluster,omitempty" name:"CrossCluster"`

	// 主节点可用区，该值仅在跨区迁移时生效。仅支持同地域下的可用区进行迁移。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
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

	// 实例隔离类型。支持值包括： "UNIVERSAL" - 通用型实例， "EXCLUSIVE" - 独享型实例， "BASIC" - 基础版实例。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 升级后的实例cpu核数， 如果不传将根据 Memory 指定的内存值自动填充对应的cpu值。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 是否极速变配。0-普通升级，1-极速变配,，2 极速优先。选择极速变配会根据资源状况校验是否可以进行极速变配，满足条件则进行极速变配，不满足条件会返回报错信息。
	FastUpgrade *int64 `json:"FastUpgrade,omitempty" name:"FastUpgrade"`

	// 延迟阈值。取值范围1~10，默认值为10。
	MaxDelayTime *int64 `json:"MaxDelayTime,omitempty" name:"MaxDelayTime"`

	// 是否跨区迁移。0-普通迁移，1-跨区迁移，默认值为0。该值为1时支持变更实例主节点可用区。
	CrossCluster *int64 `json:"CrossCluster,omitempty" name:"CrossCluster"`

	// 主节点可用区，该值仅在跨区迁移时生效。仅支持同地域下的可用区进行迁移。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *UpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "ProtectMode")
	delete(f, "DeployMode")
	delete(f, "SlaveZone")
	delete(f, "EngineVersion")
	delete(f, "WaitSwitch")
	delete(f, "BackupZone")
	delete(f, "InstanceRole")
	delete(f, "DeviceType")
	delete(f, "Cpu")
	delete(f, "FastUpgrade")
	delete(f, "MaxDelayTime")
	delete(f, "CrossCluster")
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceResponseParams struct {
	// 订单 ID。
	DealIds []*string `json:"DealIds,omitempty" name:"DealIds"`

	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDBInstanceResponseParams `json:"Response"`
}

func (r *UpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadInfo struct {
	// 文件所有分片数
	AllSliceNum *int64 `json:"AllSliceNum,omitempty" name:"AllSliceNum"`

	// 已完成分片数
	CompleteNum *int64 `json:"CompleteNum,omitempty" name:"CompleteNum"`
}

// Predefined struct for user
type VerifyRootAccountRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例 ROOT 账号的密码。
	Password *string `json:"Password,omitempty" name:"Password"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyRootAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyRootAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyRootAccountResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VerifyRootAccountResponse struct {
	*tchttp.BaseResponse
	Response *VerifyRootAccountResponseParams `json:"Response"`
}

func (r *VerifyRootAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyRootAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZoneConf struct {
	// 可用区部署方式，可能的值为：0-单可用区；1-多可用区
	DeployMode []*int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// 主实例所在的可用区
	MasterZone []*string `json:"MasterZone,omitempty" name:"MasterZone"`

	// 实例为多可用区部署时，备库1所在的可用区
	SlaveZone []*string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// 实例为多可用区部署时，备库2所在的可用区
	BackupZone []*string `json:"BackupZone,omitempty" name:"BackupZone"`
}