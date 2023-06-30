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

package v20190107

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Ability struct {
	// 是否支持从可用区
	IsSupportSlaveZone *string `json:"IsSupportSlaveZone,omitempty" name:"IsSupportSlaveZone"`

	// 不支持从可用区的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	NonsupportSlaveZoneReason *string `json:"NonsupportSlaveZoneReason,omitempty" name:"NonsupportSlaveZoneReason"`

	// 是否支持RO实例
	IsSupportRo *string `json:"IsSupportRo,omitempty" name:"IsSupportRo"`

	// 不支持RO实例的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	NonsupportRoReason *string `json:"NonsupportRoReason,omitempty" name:"NonsupportRoReason"`
}

type Account struct {
	// 数据库账号名
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 数据库账号描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 主机
	Host *string `json:"Host,omitempty" name:"Host"`

	// 用户最大连接数
	MaxUserConnections *int64 `json:"MaxUserConnections,omitempty" name:"MaxUserConnections"`
}

type AccountParam struct {
	// 参数名称，当前仅支持参数：max_user_connections
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 参数值
	ParamValue *string `json:"ParamValue,omitempty" name:"ParamValue"`
}

// Predefined struct for user
type ActivateInstanceRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例 ID 列表，单个实例 ID 格式如：cynosdbmysql-ins-n7ocdslw，与TDSQL-C MySQL数据库控制台页面中显示的实例 ID 相同，可使用 查询实例列表 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
}

type ActivateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例 ID 列表，单个实例 ID 格式如：cynosdbmysql-ins-n7ocdslw，与TDSQL-C MySQL数据库控制台页面中显示的实例 ID 相同，可使用 查询实例列表 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
}

func (r *ActivateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActivateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActivateInstanceResponseParams struct {
	// 任务流id
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ActivateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ActivateInstanceResponseParams `json:"Response"`
}

func (r *ActivateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddClusterSlaveZoneRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 从可用区
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`
}

type AddClusterSlaveZoneRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 从可用区
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`
}

func (r *AddClusterSlaveZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddClusterSlaveZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SlaveZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddClusterSlaveZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddClusterSlaveZoneResponseParams struct {
	// 异步FlowId
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddClusterSlaveZoneResponse struct {
	*tchttp.BaseResponse
	Response *AddClusterSlaveZoneResponseParams `json:"Response"`
}

func (r *AddClusterSlaveZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddClusterSlaveZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddInstancesRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cpu核数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存，单位为GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 新增只读实例数，取值范围为[0,4]
	ReadOnlyCount *int64 `json:"ReadOnlyCount,omitempty" name:"ReadOnlyCount"`

	// 实例组ID，在已有RO组中新增实例时使用，不传则新增RO组。当前版本不建议传输该值。当前版本已废弃。
	InstanceGrpId *string `json:"InstanceGrpId,omitempty" name:"InstanceGrpId"`

	// 所属VPC网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 所属子网ID，如果设置了VpcId，则SubnetId必填。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 新增RO组时使用的Port，取值范围为[0,65535)
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 实例名称，字符串长度范围为[0,64)，取值范围为大小写字母，0-9数字，'_','-','.'
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 订单来源，字符串长度范围为[0,64)
	OrderSource *string `json:"OrderSource,omitempty" name:"OrderSource"`

	// 交易模式 0-下单并支付 1-下单
	DealMode *int64 `json:"DealMode,omitempty" name:"DealMode"`

	// 参数模版ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`

	// 参数列表，ParamTemplateId 传入时InstanceParams才有效
	InstanceParams []*ModifyParamItem `json:"InstanceParams,omitempty" name:"InstanceParams"`

	// 安全组ID，新建只读实例时可以指定安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

type AddInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cpu核数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存，单位为GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 新增只读实例数，取值范围为[0,4]
	ReadOnlyCount *int64 `json:"ReadOnlyCount,omitempty" name:"ReadOnlyCount"`

	// 实例组ID，在已有RO组中新增实例时使用，不传则新增RO组。当前版本不建议传输该值。当前版本已废弃。
	InstanceGrpId *string `json:"InstanceGrpId,omitempty" name:"InstanceGrpId"`

	// 所属VPC网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 所属子网ID，如果设置了VpcId，则SubnetId必填。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 新增RO组时使用的Port，取值范围为[0,65535)
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 实例名称，字符串长度范围为[0,64)，取值范围为大小写字母，0-9数字，'_','-','.'
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 订单来源，字符串长度范围为[0,64)
	OrderSource *string `json:"OrderSource,omitempty" name:"OrderSource"`

	// 交易模式 0-下单并支付 1-下单
	DealMode *int64 `json:"DealMode,omitempty" name:"DealMode"`

	// 参数模版ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`

	// 参数列表，ParamTemplateId 传入时InstanceParams才有效
	InstanceParams []*ModifyParamItem `json:"InstanceParams,omitempty" name:"InstanceParams"`

	// 安全组ID，新建只读实例时可以指定安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

func (r *AddInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "ReadOnlyCount")
	delete(f, "InstanceGrpId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Port")
	delete(f, "InstanceName")
	delete(f, "AutoVoucher")
	delete(f, "DbType")
	delete(f, "OrderSource")
	delete(f, "DealMode")
	delete(f, "ParamTemplateId")
	delete(f, "InstanceParams")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddInstancesResponseParams struct {
	// 冻结流水，一次开通一个冻结流水。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranId *string `json:"TranId,omitempty" name:"TranId"`

	// 后付费订单号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 发货资源id列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 大订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AddInstancesResponseParams `json:"Response"`
}

func (r *AddInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Addr struct {
	// IP
	IP *string `json:"IP,omitempty" name:"IP"`

	// 端口
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

// Predefined struct for user
type AssociateSecurityGroupsRequestParams struct {
	// 实例组ID数组
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 要修改的安全组ID列表，一个或者多个安全组Id组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例组ID数组
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 要修改的安全组ID列表，一个或者多个安全组Id组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`
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
	delete(f, "InstanceIds")
	delete(f, "SecurityGroupIds")
	delete(f, "Zone")
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

type AuditLog struct {
	// 影响行数。
	AffectRows *int64 `json:"AffectRows,omitempty" name:"AffectRows"`

	// 错误码。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// SQL类型。
	SqlType *string `json:"SqlType,omitempty" name:"SqlType"`

	// 表名称。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 审计策略名称。
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 数据库名称。
	DBName *string `json:"DBName,omitempty" name:"DBName"`

	// SQL语句。
	Sql *string `json:"Sql,omitempty" name:"Sql"`

	// 客户端地址。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 用户名。
	User *string `json:"User,omitempty" name:"User"`

	// 执行时间。
	ExecTime *int64 `json:"ExecTime,omitempty" name:"ExecTime"`

	// 时间戳。
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`

	// 发送行数。
	SentRows *int64 `json:"SentRows,omitempty" name:"SentRows"`

	// 执行线程ID。
	ThreadId *int64 `json:"ThreadId,omitempty" name:"ThreadId"`
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

	// 返回行数。
	SentRows *uint64 `json:"SentRows,omitempty" name:"SentRows"`

	// 线程ID。
	ThreadId []*string `json:"ThreadId,omitempty" name:"ThreadId"`
}

type AuditRuleFilters struct {
	// 单条审计规则。
	RuleFilters []*RuleFilters `json:"RuleFilters,omitempty" name:"RuleFilters"`
}

type AuditRuleTemplateInfo struct {
	// 规则模版ID。
	RuleTemplateId *string `json:"RuleTemplateId,omitempty" name:"RuleTemplateId"`

	// 规则模版名称。
	RuleTemplateName *string `json:"RuleTemplateName,omitempty" name:"RuleTemplateName"`

	// 规则模版的过滤条件
	RuleFilters []*RuleFilters `json:"RuleFilters,omitempty" name:"RuleFilters"`

	// 规则模版描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 规则模版创建时间。
	CreateAt *string `json:"CreateAt,omitempty" name:"CreateAt"`
}

type BackupFileInfo struct {
	// 快照文件ID，已废弃，请使用BackupId
	SnapshotId *uint64 `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 备份文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 备份文件大小
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// 备份开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 备份完成时间
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// 备份类型：snapshot，快照备份；logic，逻辑备份
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// 备份方式：auto，自动备份；manual，手动备份
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// 备份文件状态：success：备份成功；fail：备份失败；creating：备份文件创建中；deleting：备份文件删除中
	BackupStatus *string `json:"BackupStatus,omitempty" name:"BackupStatus"`

	// 备份文件时间
	SnapshotTime *string `json:"SnapshotTime,omitempty" name:"SnapshotTime"`

	// 备份ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`

	// 快照类型，可选值：full，全量；increment，增量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapShotType *string `json:"SnapShotType,omitempty" name:"SnapShotType"`

	// 备份文件备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
}

type BillingResourceInfo struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 订单ID
	DealName *string `json:"DealName,omitempty" name:"DealName"`
}

// Predefined struct for user
type BindClusterResourcePackagesRequestParams struct {
	// 资源包唯一ID
	PackageIds []*string `json:"PackageIds,omitempty" name:"PackageIds"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type BindClusterResourcePackagesRequest struct {
	*tchttp.BaseRequest
	
	// 资源包唯一ID
	PackageIds []*string `json:"PackageIds,omitempty" name:"PackageIds"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *BindClusterResourcePackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindClusterResourcePackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageIds")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindClusterResourcePackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindClusterResourcePackagesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindClusterResourcePackagesResponse struct {
	*tchttp.BaseResponse
	Response *BindClusterResourcePackagesResponseParams `json:"Response"`
}

func (r *BindClusterResourcePackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindClusterResourcePackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindInstanceInfo struct {
	// 绑定的实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 绑定的实例所在的地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceRegion *string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`

	// 绑定的实例类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

type BinlogItem struct {
	// Binlog文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件大小，单位：字节
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// 事务最早时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 事务最晚时间
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// Binlog文件ID
	BinlogId *int64 `json:"BinlogId,omitempty" name:"BinlogId"`
}

// Predefined struct for user
type CloseAuditServiceRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type CloseAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CloseAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseAuditServiceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CloseAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *CloseAuditServiceResponseParams `json:"Response"`
}

func (r *CloseAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseClusterPasswordComplexityRequestParams struct {
	// 集群ID数组
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

type CloseClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID数组
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

func (r *CloseClusterPasswordComplexityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseClusterPasswordComplexityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseClusterPasswordComplexityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseClusterPasswordComplexityResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CloseClusterPasswordComplexityResponse struct {
	*tchttp.BaseResponse
	Response *CloseClusterPasswordComplexityResponseParams `json:"Response"`
}

func (r *CloseClusterPasswordComplexityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseClusterPasswordComplexityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseProxyRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 是否只关闭读写分离，取值：是 "true","false"
	OnlyCloseRW *bool `json:"OnlyCloseRW,omitempty" name:"OnlyCloseRW"`
}

type CloseProxyRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 是否只关闭读写分离，取值：是 "true","false"
	OnlyCloseRW *bool `json:"OnlyCloseRW,omitempty" name:"OnlyCloseRW"`
}

func (r *CloseProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProxyGroupId")
	delete(f, "OnlyCloseRW")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseProxyResponseParams struct {
	// 异步流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 异步任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CloseProxyResponse struct {
	*tchttp.BaseResponse
	Response *CloseProxyResponseParams `json:"Response"`
}

func (r *CloseProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseWanRequestParams struct {
	// 实例组id
	InstanceGrpId *string `json:"InstanceGrpId,omitempty" name:"InstanceGrpId"`
}

type CloseWanRequest struct {
	*tchttp.BaseRequest
	
	// 实例组id
	InstanceGrpId *string `json:"InstanceGrpId,omitempty" name:"InstanceGrpId"`
}

func (r *CloseWanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseWanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceGrpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseWanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseWanResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CloseWanResponse struct {
	*tchttp.BaseResponse
	Response *CloseWanResponseParams `json:"Response"`
}

func (r *CloseWanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseWanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterInstanceDetail struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 引擎类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例状态
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// 实例状态描述
	InstanceStatusDesc *string `json:"InstanceStatusDesc,omitempty" name:"InstanceStatusDesc"`

	// cpu核数
	InstanceCpu *int64 `json:"InstanceCpu,omitempty" name:"InstanceCpu"`

	// 内存
	InstanceMemory *int64 `json:"InstanceMemory,omitempty" name:"InstanceMemory"`

	// 硬盘
	InstanceStorage *int64 `json:"InstanceStorage,omitempty" name:"InstanceStorage"`

	// 实例角色
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
}

type ClusterParamModifyLog struct {
	// 参数名称
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 当前值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 修改后的值
	UpdateValue *string `json:"UpdateValue,omitempty" name:"UpdateValue"`

	// 修改状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

// Predefined struct for user
type CopyClusterPasswordComplexityRequestParams struct {
	// 复制集群ID数组
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// 集群id
	SourceClusterId *string `json:"SourceClusterId,omitempty" name:"SourceClusterId"`
}

type CopyClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// 复制集群ID数组
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// 集群id
	SourceClusterId *string `json:"SourceClusterId,omitempty" name:"SourceClusterId"`
}

func (r *CopyClusterPasswordComplexityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyClusterPasswordComplexityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "SourceClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyClusterPasswordComplexityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyClusterPasswordComplexityResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CopyClusterPasswordComplexityResponse struct {
	*tchttp.BaseResponse
	Response *CopyClusterPasswordComplexityResponseParams `json:"Response"`
}

func (r *CopyClusterPasswordComplexityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyClusterPasswordComplexityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountsRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 新账户列表
	Accounts []*NewAccount `json:"Accounts,omitempty" name:"Accounts"`
}

type CreateAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 新账户列表
	Accounts []*NewAccount `json:"Accounts,omitempty" name:"Accounts"`
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
	delete(f, "ClusterId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountsResponseParams struct {
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
	// 实例ID
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
	
	// 实例ID
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
type CreateAuditRuleTemplateRequestParams struct {
	// 审计规则。
	RuleFilters []*RuleFilters `json:"RuleFilters,omitempty" name:"RuleFilters"`

	// 规则模版名称。
	RuleTemplateName *string `json:"RuleTemplateName,omitempty" name:"RuleTemplateName"`

	// 规则模版描述。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateAuditRuleTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 审计规则。
	RuleFilters []*RuleFilters `json:"RuleFilters,omitempty" name:"RuleFilters"`

	// 规则模版名称。
	RuleTemplateName *string `json:"RuleTemplateName,omitempty" name:"RuleTemplateName"`

	// 规则模版描述。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateAuditRuleTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditRuleTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleFilters")
	delete(f, "RuleTemplateName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditRuleTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditRuleTemplateResponseParams struct {
	// 生成的规则模版ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleTemplateId *string `json:"RuleTemplateId,omitempty" name:"RuleTemplateId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAuditRuleTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuditRuleTemplateResponseParams `json:"Response"`
}

func (r *CreateAuditRuleTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditRuleTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 备份类型, 可选值：logic，逻辑备份；snapshot，物理备份
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// 备份的库, 只在 BackupType 为 logic 时有效
	BackupDatabases []*string `json:"BackupDatabases,omitempty" name:"BackupDatabases"`

	// 备份的表, 只在 BackupType 为 logic 时有效
	BackupTables []*DatabaseTables `json:"BackupTables,omitempty" name:"BackupTables"`

	// 备注名
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 备份类型, 可选值：logic，逻辑备份；snapshot，物理备份
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// 备份的库, 只在 BackupType 为 logic 时有效
	BackupDatabases []*string `json:"BackupDatabases,omitempty" name:"BackupDatabases"`

	// 备份的表, 只在 BackupType 为 logic 时有效
	BackupTables []*DatabaseTables `json:"BackupTables,omitempty" name:"BackupTables"`

	// 备注名
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
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
	delete(f, "ClusterId")
	delete(f, "BackupType")
	delete(f, "BackupDatabases")
	delete(f, "BackupTables")
	delete(f, "BackupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupResponseParams struct {
	// 异步任务流ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

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
type CreateClusterDatabaseRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库名
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 字符集类型
	CharacterSet *string `json:"CharacterSet,omitempty" name:"CharacterSet"`

	// 排序规则
	CollateRule *string `json:"CollateRule,omitempty" name:"CollateRule"`

	// 授权用户主机权限
	UserHostPrivileges []*UserHostPrivilege `json:"UserHostPrivileges,omitempty" name:"UserHostPrivileges"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateClusterDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库名
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 字符集类型
	CharacterSet *string `json:"CharacterSet,omitempty" name:"CharacterSet"`

	// 排序规则
	CollateRule *string `json:"CollateRule,omitempty" name:"CollateRule"`

	// 授权用户主机权限
	UserHostPrivileges []*UserHostPrivilege `json:"UserHostPrivileges,omitempty" name:"UserHostPrivileges"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateClusterDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DbName")
	delete(f, "CharacterSet")
	delete(f, "CollateRule")
	delete(f, "UserHostPrivileges")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterDatabaseResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClusterDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterDatabaseResponseParams `json:"Response"`
}

func (r *CreateClusterDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClustersRequestParams struct {
	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 所属VPC网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 所属子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 数据库版本，取值范围: 
	// <li> MYSQL可选值：5.7，8.0 </li>
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// 所属项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 当DbMode为NORMAL或不填时必选
	// 普通实例Cpu核数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 当DbMode为NORMAL或不填时必选
	// 普通实例内存,单位G
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 该参数无实际意义，已废弃。
	// 存储大小，单位G。
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 集群名称，长度小于64个字符，每个字符取值范围：大/小写字母，数字，特殊符号（'-','_','.'）
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 账号密码(8-64个字符，包含大小写英文字母、数字和符号~!@#$%^&*_-+=`|\(){}[]:;'<>,.?/中的任意三种)
	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`

	// 端口，默认3306，取值范围[0, 65535)
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 计费模式，按量计费：0，包年包月：1。默认按量计费。
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 购买集群数，可选值范围[1,50]，默认为1
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 回档类型：
	// noneRollback：不回档；
	// snapRollback，快照回档；
	// timeRollback，时间点回档
	RollbackStrategy *string `json:"RollbackStrategy,omitempty" name:"RollbackStrategy"`

	// 快照回档，表示snapshotId；时间点回档，表示queryId，为0，表示需要判断时间点是否有效
	RollbackId *uint64 `json:"RollbackId,omitempty" name:"RollbackId"`

	// 回档时，传入源集群ID，用于查找源poolId
	OriginalClusterId *string `json:"OriginalClusterId,omitempty" name:"OriginalClusterId"`

	// 时间点回档，指定时间；快照回档，快照时间
	ExpectTime *string `json:"ExpectTime,omitempty" name:"ExpectTime"`

	// 该参数无实际意义，已废弃。
	// 时间点回档，指定时间允许范围
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitempty" name:"ExpectTimeThresh"`

	// 普通实例存储上限，单位GB
	// 当DbType为MYSQL，且存储计费模式为预付费时，该参数需不大于cpu与memory对应存储规格上限
	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// 实例数量，数量范围为(0,16]
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 包年包月购买时长
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 包年包月购买时长单位，['s','d','m','y']
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 包年包月购买是否自动续费，默认为0
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 实例数量（该参数已不再使用，只做存量兼容处理）
	HaCount *int64 `json:"HaCount,omitempty" name:"HaCount"`

	// 订单来源
	OrderSource *string `json:"OrderSource,omitempty" name:"OrderSource"`

	// 集群创建需要绑定的tag数组信息
	ResourceTags []*Tag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// Db类型
	// 当DbType为MYSQL时可选(默认NORMAL)：
	// <li>NORMAL</li>
	// <li>SERVERLESS</li>
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// 当DbMode为SEVERLESS时必填
	// cpu最小值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MinCpu *float64 `json:"MinCpu,omitempty" name:"MinCpu"`

	// 当DbMode为SEVERLESS时必填：
	// cpu最大值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MaxCpu *float64 `json:"MaxCpu,omitempty" name:"MaxCpu"`

	// 当DbMode为SEVERLESS时，指定集群是否自动暂停，可选范围
	// <li>yes</li>
	// <li>no</li>
	// 默认值:yes
	AutoPause *string `json:"AutoPause,omitempty" name:"AutoPause"`

	// 当DbMode为SEVERLESS时，指定集群自动暂停的延迟，单位秒，可选范围[600,691200]
	// 默认值:600
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitempty" name:"AutoPauseDelay"`

	// 集群存储计费模式，按量计费：0，包年包月：1。默认按量计费
	// 当DbType为MYSQL时，在集群计算计费模式为后付费（包括DbMode为SERVERLESS）时，存储计费模式仅可为按量计费
	// 回档与克隆均不支持包年包月存储
	StoragePayMode *int64 `json:"StoragePayMode,omitempty" name:"StoragePayMode"`

	// 安全组id数组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 告警策略Id数组
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitempty" name:"AlarmPolicyIds"`

	// 参数数组，暂时支持character_set_server （utf8｜latin1｜gbk｜utf8mb4） ，lower_case_table_names，1-大小写不敏感，0-大小写敏感
	ClusterParams []*ParamItem `json:"ClusterParams,omitempty" name:"ClusterParams"`

	// 交易模式，0-下单且支付，1-下单
	DealMode *int64 `json:"DealMode,omitempty" name:"DealMode"`

	// 参数模版ID，可以通过查询参数模板信息DescribeParamTemplates获得参数模板ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`

	// 多可用区地址
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// 实例初始化配置信息，主要用于购买集群时选不同规格实例
	InstanceInitInfos []*InstanceInitInfo `json:"InstanceInitInfos,omitempty" name:"InstanceInitInfos"`
}

type CreateClustersRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 所属VPC网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 所属子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 数据库版本，取值范围: 
	// <li> MYSQL可选值：5.7，8.0 </li>
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// 所属项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 当DbMode为NORMAL或不填时必选
	// 普通实例Cpu核数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 当DbMode为NORMAL或不填时必选
	// 普通实例内存,单位G
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 该参数无实际意义，已废弃。
	// 存储大小，单位G。
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 集群名称，长度小于64个字符，每个字符取值范围：大/小写字母，数字，特殊符号（'-','_','.'）
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 账号密码(8-64个字符，包含大小写英文字母、数字和符号~!@#$%^&*_-+=`|\(){}[]:;'<>,.?/中的任意三种)
	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`

	// 端口，默认3306，取值范围[0, 65535)
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 计费模式，按量计费：0，包年包月：1。默认按量计费。
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 购买集群数，可选值范围[1,50]，默认为1
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 回档类型：
	// noneRollback：不回档；
	// snapRollback，快照回档；
	// timeRollback，时间点回档
	RollbackStrategy *string `json:"RollbackStrategy,omitempty" name:"RollbackStrategy"`

	// 快照回档，表示snapshotId；时间点回档，表示queryId，为0，表示需要判断时间点是否有效
	RollbackId *uint64 `json:"RollbackId,omitempty" name:"RollbackId"`

	// 回档时，传入源集群ID，用于查找源poolId
	OriginalClusterId *string `json:"OriginalClusterId,omitempty" name:"OriginalClusterId"`

	// 时间点回档，指定时间；快照回档，快照时间
	ExpectTime *string `json:"ExpectTime,omitempty" name:"ExpectTime"`

	// 该参数无实际意义，已废弃。
	// 时间点回档，指定时间允许范围
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitempty" name:"ExpectTimeThresh"`

	// 普通实例存储上限，单位GB
	// 当DbType为MYSQL，且存储计费模式为预付费时，该参数需不大于cpu与memory对应存储规格上限
	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// 实例数量，数量范围为(0,16]
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 包年包月购买时长
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 包年包月购买时长单位，['s','d','m','y']
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 包年包月购买是否自动续费，默认为0
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 实例数量（该参数已不再使用，只做存量兼容处理）
	HaCount *int64 `json:"HaCount,omitempty" name:"HaCount"`

	// 订单来源
	OrderSource *string `json:"OrderSource,omitempty" name:"OrderSource"`

	// 集群创建需要绑定的tag数组信息
	ResourceTags []*Tag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// Db类型
	// 当DbType为MYSQL时可选(默认NORMAL)：
	// <li>NORMAL</li>
	// <li>SERVERLESS</li>
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// 当DbMode为SEVERLESS时必填
	// cpu最小值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MinCpu *float64 `json:"MinCpu,omitempty" name:"MinCpu"`

	// 当DbMode为SEVERLESS时必填：
	// cpu最大值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MaxCpu *float64 `json:"MaxCpu,omitempty" name:"MaxCpu"`

	// 当DbMode为SEVERLESS时，指定集群是否自动暂停，可选范围
	// <li>yes</li>
	// <li>no</li>
	// 默认值:yes
	AutoPause *string `json:"AutoPause,omitempty" name:"AutoPause"`

	// 当DbMode为SEVERLESS时，指定集群自动暂停的延迟，单位秒，可选范围[600,691200]
	// 默认值:600
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitempty" name:"AutoPauseDelay"`

	// 集群存储计费模式，按量计费：0，包年包月：1。默认按量计费
	// 当DbType为MYSQL时，在集群计算计费模式为后付费（包括DbMode为SERVERLESS）时，存储计费模式仅可为按量计费
	// 回档与克隆均不支持包年包月存储
	StoragePayMode *int64 `json:"StoragePayMode,omitempty" name:"StoragePayMode"`

	// 安全组id数组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 告警策略Id数组
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitempty" name:"AlarmPolicyIds"`

	// 参数数组，暂时支持character_set_server （utf8｜latin1｜gbk｜utf8mb4） ，lower_case_table_names，1-大小写不敏感，0-大小写敏感
	ClusterParams []*ParamItem `json:"ClusterParams,omitempty" name:"ClusterParams"`

	// 交易模式，0-下单且支付，1-下单
	DealMode *int64 `json:"DealMode,omitempty" name:"DealMode"`

	// 参数模版ID，可以通过查询参数模板信息DescribeParamTemplates获得参数模板ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`

	// 多可用区地址
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// 实例初始化配置信息，主要用于购买集群时选不同规格实例
	InstanceInitInfos []*InstanceInitInfo `json:"InstanceInitInfos,omitempty" name:"InstanceInitInfos"`
}

func (r *CreateClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "DbType")
	delete(f, "DbVersion")
	delete(f, "ProjectId")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "ClusterName")
	delete(f, "AdminPassword")
	delete(f, "Port")
	delete(f, "PayMode")
	delete(f, "Count")
	delete(f, "RollbackStrategy")
	delete(f, "RollbackId")
	delete(f, "OriginalClusterId")
	delete(f, "ExpectTime")
	delete(f, "ExpectTimeThresh")
	delete(f, "StorageLimit")
	delete(f, "InstanceCount")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "AutoRenewFlag")
	delete(f, "AutoVoucher")
	delete(f, "HaCount")
	delete(f, "OrderSource")
	delete(f, "ResourceTags")
	delete(f, "DbMode")
	delete(f, "MinCpu")
	delete(f, "MaxCpu")
	delete(f, "AutoPause")
	delete(f, "AutoPauseDelay")
	delete(f, "StoragePayMode")
	delete(f, "SecurityGroupIds")
	delete(f, "AlarmPolicyIds")
	delete(f, "ClusterParams")
	delete(f, "DealMode")
	delete(f, "ParamTemplateId")
	delete(f, "SlaveZone")
	delete(f, "InstanceInitInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClustersResponseParams struct {
	// 冻结流水ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranId *string `json:"TranId,omitempty" name:"TranId"`

	// 订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 资源ID列表（该字段已不再维护，请使用dealNames字段查询接口DescribeResourcesByDealName获取资源ID）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 集群ID列表（该字段已不再维护，请使用dealNames字段查询接口DescribeResourcesByDealName获取集群ID）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// 大订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClustersResponse struct {
	*tchttp.BaseResponse
	Response *CreateClustersResponseParams `json:"Response"`
}

func (r *CreateClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParamTemplateRequestParams struct {
	// 模版名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// mysql版本号
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 模版描述
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`

	// 可选参数，需要复制的模版ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 数据库类型，可选值：NORMAL（默认值），SERVERLESS
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// 参数列表
	ParamList []*ParamItem `json:"ParamList,omitempty" name:"ParamList"`
}

type CreateParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模版名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// mysql版本号
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 模版描述
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`

	// 可选参数，需要复制的模版ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 数据库类型，可选值：NORMAL（默认值），SERVERLESS
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// 参数列表
	ParamList []*ParamItem `json:"ParamList,omitempty" name:"ParamList"`
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
	delete(f, "TemplateName")
	delete(f, "EngineVersion")
	delete(f, "TemplateDescription")
	delete(f, "TemplateId")
	delete(f, "DbMode")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParamTemplateResponseParams struct {
	// 模版ID
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
type CreateProxyEndPointRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 私有网络ID，默认与集群私有网络ID保持一致
	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`

	// 私有网络子网ID，默认与集群子网ID保持一致
	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`

	// 连接池类型：SessionConnectionPool(会话级别连接池 )
	ConnectionPoolType *string `json:"ConnectionPoolType,omitempty" name:"ConnectionPoolType"`

	// 是否开启连接池,yes-开启，no-不开启
	OpenConnectionPool *string `json:"OpenConnectionPool,omitempty" name:"OpenConnectionPool"`

	// 连接池阀值：单位（秒）
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitempty" name:"ConnectionPoolTimeOut"`

	// 安全组ID数组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 描述说明
	Description *string `json:"Description,omitempty" name:"Description"`

	// vip信息
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 权重模式：
	// system-系统分配，custom-自定义
	WeightMode *string `json:"WeightMode,omitempty" name:"WeightMode"`

	// 是否自动添加只读实例，yes-是，no-不自动添加
	AutoAddRo *string `json:"AutoAddRo,omitempty" name:"AutoAddRo"`

	// 是否开启故障转移
	FailOver *string `json:"FailOver,omitempty" name:"FailOver"`

	// 一致性类型：
	// eventual,global,session
	ConsistencyType *string `json:"ConsistencyType,omitempty" name:"ConsistencyType"`

	// 读写属性：
	// READWRITE,READONLY
	RwType *string `json:"RwType,omitempty" name:"RwType"`

	// 一致性超时时间
	ConsistencyTimeOut *int64 `json:"ConsistencyTimeOut,omitempty" name:"ConsistencyTimeOut"`

	// 事务拆分
	TransSplit *bool `json:"TransSplit,omitempty" name:"TransSplit"`

	// 连接模式：
	// nearby,balance
	AccessMode *string `json:"AccessMode,omitempty" name:"AccessMode"`

	// 实例权重
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitempty" name:"InstanceWeights"`
}

type CreateProxyEndPointRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 私有网络ID，默认与集群私有网络ID保持一致
	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`

	// 私有网络子网ID，默认与集群子网ID保持一致
	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`

	// 连接池类型：SessionConnectionPool(会话级别连接池 )
	ConnectionPoolType *string `json:"ConnectionPoolType,omitempty" name:"ConnectionPoolType"`

	// 是否开启连接池,yes-开启，no-不开启
	OpenConnectionPool *string `json:"OpenConnectionPool,omitempty" name:"OpenConnectionPool"`

	// 连接池阀值：单位（秒）
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitempty" name:"ConnectionPoolTimeOut"`

	// 安全组ID数组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 描述说明
	Description *string `json:"Description,omitempty" name:"Description"`

	// vip信息
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 权重模式：
	// system-系统分配，custom-自定义
	WeightMode *string `json:"WeightMode,omitempty" name:"WeightMode"`

	// 是否自动添加只读实例，yes-是，no-不自动添加
	AutoAddRo *string `json:"AutoAddRo,omitempty" name:"AutoAddRo"`

	// 是否开启故障转移
	FailOver *string `json:"FailOver,omitempty" name:"FailOver"`

	// 一致性类型：
	// eventual,global,session
	ConsistencyType *string `json:"ConsistencyType,omitempty" name:"ConsistencyType"`

	// 读写属性：
	// READWRITE,READONLY
	RwType *string `json:"RwType,omitempty" name:"RwType"`

	// 一致性超时时间
	ConsistencyTimeOut *int64 `json:"ConsistencyTimeOut,omitempty" name:"ConsistencyTimeOut"`

	// 事务拆分
	TransSplit *bool `json:"TransSplit,omitempty" name:"TransSplit"`

	// 连接模式：
	// nearby,balance
	AccessMode *string `json:"AccessMode,omitempty" name:"AccessMode"`

	// 实例权重
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitempty" name:"InstanceWeights"`
}

func (r *CreateProxyEndPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyEndPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "UniqueVpcId")
	delete(f, "UniqueSubnetId")
	delete(f, "ConnectionPoolType")
	delete(f, "OpenConnectionPool")
	delete(f, "ConnectionPoolTimeOut")
	delete(f, "SecurityGroupIds")
	delete(f, "Description")
	delete(f, "Vip")
	delete(f, "WeightMode")
	delete(f, "AutoAddRo")
	delete(f, "FailOver")
	delete(f, "ConsistencyType")
	delete(f, "RwType")
	delete(f, "ConsistencyTimeOut")
	delete(f, "TransSplit")
	delete(f, "AccessMode")
	delete(f, "InstanceWeights")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProxyEndPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProxyEndPointResponseParams struct {
	// 异步流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 异步任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateProxyEndPointResponse struct {
	*tchttp.BaseResponse
	Response *CreateProxyEndPointResponseParams `json:"Response"`
}

func (r *CreateProxyEndPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyEndPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProxyRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// cpu核数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存
	Mem *int64 `json:"Mem,omitempty" name:"Mem"`

	// 私有网络ID，默认与集群私有网络ID保持一致
	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`

	// 私有网络子网ID，默认与集群子网ID保持一致
	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`

	// 数据库代理组节点个数
	ProxyCount *int64 `json:"ProxyCount,omitempty" name:"ProxyCount"`

	// 连接池类型：SessionConnectionPool(会话级别连接池 )
	ConnectionPoolType *string `json:"ConnectionPoolType,omitempty" name:"ConnectionPoolType"`

	// 是否开启连接池,yes-开启，no-不开启
	OpenConnectionPool *string `json:"OpenConnectionPool,omitempty" name:"OpenConnectionPool"`

	// 连接池阀值：单位（秒）
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitempty" name:"ConnectionPoolTimeOut"`

	// 安全组ID数组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 描述说明
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据库节点信息
	ProxyZones []*ProxyZone `json:"ProxyZones,omitempty" name:"ProxyZones"`
}

type CreateProxyRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// cpu核数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存
	Mem *int64 `json:"Mem,omitempty" name:"Mem"`

	// 私有网络ID，默认与集群私有网络ID保持一致
	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`

	// 私有网络子网ID，默认与集群子网ID保持一致
	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`

	// 数据库代理组节点个数
	ProxyCount *int64 `json:"ProxyCount,omitempty" name:"ProxyCount"`

	// 连接池类型：SessionConnectionPool(会话级别连接池 )
	ConnectionPoolType *string `json:"ConnectionPoolType,omitempty" name:"ConnectionPoolType"`

	// 是否开启连接池,yes-开启，no-不开启
	OpenConnectionPool *string `json:"OpenConnectionPool,omitempty" name:"OpenConnectionPool"`

	// 连接池阀值：单位（秒）
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitempty" name:"ConnectionPoolTimeOut"`

	// 安全组ID数组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 描述说明
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据库节点信息
	ProxyZones []*ProxyZone `json:"ProxyZones,omitempty" name:"ProxyZones"`
}

func (r *CreateProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Cpu")
	delete(f, "Mem")
	delete(f, "UniqueVpcId")
	delete(f, "UniqueSubnetId")
	delete(f, "ProxyCount")
	delete(f, "ConnectionPoolType")
	delete(f, "OpenConnectionPool")
	delete(f, "ConnectionPoolTimeOut")
	delete(f, "SecurityGroupIds")
	delete(f, "Description")
	delete(f, "ProxyZones")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProxyResponseParams struct {
	// 异步流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 异步任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateProxyResponse struct {
	*tchttp.BaseResponse
	Response *CreateProxyResponseParams `json:"Response"`
}

func (r *CreateProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourcePackageRequestParams struct {
	// 实例类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 资源包使用地域
	// china-中国内地通用，overseas-港澳台及海外通用
	PackageRegion *string `json:"PackageRegion,omitempty" name:"PackageRegion"`

	// 资源包类型：CCU-计算资源包，DISK-存储资源包
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// 资源包版本
	// base-基础版本，common-通用版本，enterprise-企业版本
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

	// 资源包大小，计算资源单位：万个；存储资源：GB
	PackageSpec *float64 `json:"PackageSpec,omitempty" name:"PackageSpec"`

	// 资源包有效期，单位:天
	ExpireDay *int64 `json:"ExpireDay,omitempty" name:"ExpireDay"`

	// 购买资源包个数
	PackageCount *int64 `json:"PackageCount,omitempty" name:"PackageCount"`

	// 资源包名称
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`
}

type CreateResourcePackageRequest struct {
	*tchttp.BaseRequest
	
	// 实例类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 资源包使用地域
	// china-中国内地通用，overseas-港澳台及海外通用
	PackageRegion *string `json:"PackageRegion,omitempty" name:"PackageRegion"`

	// 资源包类型：CCU-计算资源包，DISK-存储资源包
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// 资源包版本
	// base-基础版本，common-通用版本，enterprise-企业版本
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

	// 资源包大小，计算资源单位：万个；存储资源：GB
	PackageSpec *float64 `json:"PackageSpec,omitempty" name:"PackageSpec"`

	// 资源包有效期，单位:天
	ExpireDay *int64 `json:"ExpireDay,omitempty" name:"ExpireDay"`

	// 购买资源包个数
	PackageCount *int64 `json:"PackageCount,omitempty" name:"PackageCount"`

	// 资源包名称
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`
}

func (r *CreateResourcePackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourcePackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "PackageRegion")
	delete(f, "PackageType")
	delete(f, "PackageVersion")
	delete(f, "PackageSpec")
	delete(f, "ExpireDay")
	delete(f, "PackageCount")
	delete(f, "PackageName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResourcePackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourcePackageResponseParams struct {
	// 付费总订单号
	BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`

	// 每个物品对应一个dealName，业务需要根据dealName保证发货接口幂等
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateResourcePackageResponse struct {
	*tchttp.BaseResponse
	Response *CreateResourcePackageResponseParams `json:"Response"`
}

func (r *CreateResourcePackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourcePackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CynosdbCluster struct {
	// 集群状态， 可选值如下:
	// creating: 创建中
	// running:运行中
	// isolating:隔离中
	// isolated:已隔离
	// activating:解隔离中
	// offlining:下线中
	// offlined:已下线
	// deleting:删除中
	// deleted:已删除
	Status *string `json:"Status,omitempty" name:"Status"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 数据库版本
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例数
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 引擎类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 用户appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 集群状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 集群创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 付费模式。0-按量计费，1-包年包月
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 截止时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// 集群读写vip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 集群读写vport
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// cynos内核版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	CynosVersion *string `json:"CynosVersion,omitempty" name:"CynosVersion"`

	// 存储容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// 续费标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 正在处理的任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessingTask *string `json:"ProcessingTask,omitempty" name:"ProcessingTask"`

	// 集群的任务数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*ObjectTask `json:"Tasks,omitempty" name:"Tasks"`

	// 集群绑定的tag数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceTags []*Tag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// Db类型(NORMAL, SERVERLESS)
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// 当Db类型为SERVERLESS时，serverless集群状态，可选值:
	// resume
	// pause
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerlessStatus *string `json:"ServerlessStatus,omitempty" name:"ServerlessStatus"`

	// 集群预付费存储值大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 集群存储为预付费时的存储ID，用于预付费存储变配
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageId *string `json:"StorageId,omitempty" name:"StorageId"`

	// 集群存储付费模式。0-按量计费，1-包年包月
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoragePayMode *int64 `json:"StoragePayMode,omitempty" name:"StoragePayMode"`

	// 集群计算规格对应的最小存储值
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinStorageSize *int64 `json:"MinStorageSize,omitempty" name:"MinStorageSize"`

	// 集群计算规格对应的最大存储值
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxStorageSize *int64 `json:"MaxStorageSize,omitempty" name:"MaxStorageSize"`

	// 集群网络信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetAddrs []*NetAddr `json:"NetAddrs,omitempty" name:"NetAddrs"`

	// 物理可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhysicalZone *string `json:"PhysicalZone,omitempty" name:"PhysicalZone"`

	// 主可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`

	// 是否有从可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasSlaveZone *string `json:"HasSlaveZone,omitempty" name:"HasSlaveZone"`

	// 从可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaveZones []*string `json:"SlaveZones,omitempty" name:"SlaveZones"`

	// 商业类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// 是否冻结
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsFreeze *string `json:"IsFreeze,omitempty" name:"IsFreeze"`

	// 订单来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderSource *string `json:"OrderSource,omitempty" name:"OrderSource"`

	// 能力
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ability *Ability `json:"Ability,omitempty" name:"Ability"`

	// 实例绑定资源包信息（此处只返回存储资源包，即packageType=DISK）	
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourcePackages []*ResourcePackage `json:"ResourcePackages,omitempty" name:"ResourcePackages"`
}

type CynosdbClusterDetail struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 物理可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhysicalZone *string `json:"PhysicalZone,omitempty" name:"PhysicalZone"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 当Db类型为SERVERLESS时，serverless集群状态，可选值:
	// resume
	// resuming
	// pause
	// pausing
	ServerlessStatus *string `json:"ServerlessStatus,omitempty" name:"ServerlessStatus"`

	// 存储Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageId *string `json:"StorageId,omitempty" name:"StorageId"`

	// 存储大小，单位为G
	// 注意：此字段可能返回 null，表示取不到有效值。
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 最大存储规格，单位为G
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxStorageSize *int64 `json:"MaxStorageSize,omitempty" name:"MaxStorageSize"`

	// 最小存储规格，单位为G
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinStorageSize *int64 `json:"MinStorageSize,omitempty" name:"MinStorageSize"`

	// 存储付费类型，1为包年包月，0为按量计费
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoragePayMode *int64 `json:"StoragePayMode,omitempty" name:"StoragePayMode"`

	// VPC名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// vpc唯一id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网名称
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 字符集
	Charset *string `json:"Charset,omitempty" name:"Charset"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 数据库类型
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 数据库类型，normal，serverless
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// 数据库版本
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// 存储空间上限
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// 使用容量
	UsedStorage *int64 `json:"UsedStorage,omitempty" name:"UsedStorage"`

	// vip地址
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// vport端口
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 读写分离Vport
	RoAddr []*Addr `json:"RoAddr,omitempty" name:"RoAddr"`

	// 集群支持的功能
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ability *Ability `json:"Ability,omitempty" name:"Ability"`

	// cynos版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	CynosVersion *string `json:"CynosVersion,omitempty" name:"CynosVersion"`

	// 商业类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// 是否有从可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasSlaveZone *string `json:"HasSlaveZone,omitempty" name:"HasSlaveZone"`

	// 是否冻结
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsFreeze *string `json:"IsFreeze,omitempty" name:"IsFreeze"`

	// 任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*ObjectTask `json:"Tasks,omitempty" name:"Tasks"`

	// 主可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`

	// 从可用区列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaveZones []*string `json:"SlaveZones,omitempty" name:"SlaveZones"`

	// 实例信息
	InstanceSet []*ClusterInstanceDetail `json:"InstanceSet,omitempty" name:"InstanceSet"`

	// 付费模式
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 到期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// 项目id
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// 实例绑定的tag数组信息
	ResourceTags []*Tag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// Proxy状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyStatus *string `json:"ProxyStatus,omitempty" name:"ProxyStatus"`

	// binlog开关，可选值：ON, OFF
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogBin *string `json:"LogBin,omitempty" name:"LogBin"`

	// 是否跳过交易
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSkipTrade *string `json:"IsSkipTrade,omitempty" name:"IsSkipTrade"`

	// pitr类型，可选值：normal, redo_pitr
	// 注意：此字段可能返回 null，表示取不到有效值。
	PitrType *string `json:"PitrType,omitempty" name:"PitrType"`

	// 是否打开密码复杂度
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsOpenPasswordComplexity *string `json:"IsOpenPasswordComplexity,omitempty" name:"IsOpenPasswordComplexity"`

	// 网络类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkStatus *string `json:"NetworkStatus,omitempty" name:"NetworkStatus"`

	// 集群绑定的资源包信息	
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourcePackages []*ResourcePackage `json:"ResourcePackages,omitempty" name:"ResourcePackages"`

	// 自动续费标识，1为自动续费，0为到期不续
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type CynosdbErrorLogItem struct {
	// 日志时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 日志等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *string `json:"Level,omitempty" name:"Level"`

	// 日志内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`
}

type CynosdbInstance struct {
	// 用户Uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 用户AppId
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例状态中文描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 实例形态，是否为serverless实例
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// 数据库类型
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 数据库版本
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// Cpu，单位：核
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存，单位：GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 存储量，单位：GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 实例类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例当前角色
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// VPC网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例内网IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 实例内网端口
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 付费模式
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 实例过期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// 销毁期限
	DestroyDeadlineText *string `json:"DestroyDeadlineText,omitempty" name:"DestroyDeadlineText"`

	// 隔离时间
	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`

	// 网络类型
	NetType *int64 `json:"NetType,omitempty" name:"NetType"`

	// 外网域名
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`

	// 外网IP
	WanIP *string `json:"WanIP,omitempty" name:"WanIP"`

	// 外网端口
	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`

	// 外网状态
	WanStatus *string `json:"WanStatus,omitempty" name:"WanStatus"`

	// 实例销毁时间
	DestroyTime *string `json:"DestroyTime,omitempty" name:"DestroyTime"`

	// Cynos内核版本
	CynosVersion *string `json:"CynosVersion,omitempty" name:"CynosVersion"`

	// 正在处理的任务
	ProcessingTask *string `json:"ProcessingTask,omitempty" name:"ProcessingTask"`

	// 续费标志
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// serverless实例cpu下限
	MinCpu *float64 `json:"MinCpu,omitempty" name:"MinCpu"`

	// serverless实例cpu上限
	MaxCpu *float64 `json:"MaxCpu,omitempty" name:"MaxCpu"`

	// serverless实例状态, 可选值：
	// resume
	// pause
	ServerlessStatus *string `json:"ServerlessStatus,omitempty" name:"ServerlessStatus"`

	// 预付费存储Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageId *string `json:"StorageId,omitempty" name:"StorageId"`

	// 存储付费类型
	StoragePayMode *int64 `json:"StoragePayMode,omitempty" name:"StoragePayMode"`

	// 物理区
	PhysicalZone *string `json:"PhysicalZone,omitempty" name:"PhysicalZone"`

	// 商业类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// 任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*ObjectTask `json:"Tasks,omitempty" name:"Tasks"`

	// 是否冻结
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsFreeze *string `json:"IsFreeze,omitempty" name:"IsFreeze"`

	// 资源标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceTags []*Tag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 主可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`

	// 备可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaveZones []*string `json:"SlaveZones,omitempty" name:"SlaveZones"`

	// 实例网络信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNetInfo []*InstanceNetInfo `json:"InstanceNetInfo,omitempty" name:"InstanceNetInfo"`

	// 实例绑定资源包信息（此处只返回计算资源包，即packageType=CCU）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourcePackages []*ResourcePackage `json:"ResourcePackages,omitempty" name:"ResourcePackages"`
}

type CynosdbInstanceDetail struct {
	// 用户Uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 用户AppId
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例状态中文描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 数据库类型
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 数据库版本
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// Cpu，单位：核
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存，单位：GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 存储量，单位：GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 实例类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例当前角色
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 付费模式
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 实例过期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// 网络类型
	NetType *int64 `json:"NetType,omitempty" name:"NetType"`

	// VPC网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例内网IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 实例内网端口
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 实例外网域名
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`

	// 字符集
	Charset *string `json:"Charset,omitempty" name:"Charset"`

	// Cynos内核版本
	CynosVersion *string `json:"CynosVersion,omitempty" name:"CynosVersion"`

	// 续费标志
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// serverless实例cpu下限
	MinCpu *float64 `json:"MinCpu,omitempty" name:"MinCpu"`

	// serverless实例cpu上限
	MaxCpu *float64 `json:"MaxCpu,omitempty" name:"MaxCpu"`

	// serverless实例状态, 可能值：
	// resume
	// pause
	ServerlessStatus *string `json:"ServerlessStatus,omitempty" name:"ServerlessStatus"`
}

type CynosdbInstanceGrp struct {
	// 用户appId
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 删除时间
	DeletedTime *string `json:"DeletedTime,omitempty" name:"DeletedTime"`

	// 实例组ID
	InstanceGrpId *string `json:"InstanceGrpId,omitempty" name:"InstanceGrpId"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例组类型。ha-ha组；ro-只读组
	Type *string `json:"Type,omitempty" name:"Type"`

	// 更新时间
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 内网IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 内网端口
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 外网域名
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`

	// 外网ip
	WanIP *string `json:"WanIP,omitempty" name:"WanIP"`

	// 外网端口
	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`

	// 外网状态
	WanStatus *string `json:"WanStatus,omitempty" name:"WanStatus"`

	// 实例组包含实例信息
	InstanceSet []*CynosdbInstance `json:"InstanceSet,omitempty" name:"InstanceSet"`

	// VPC的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 正在回收IP信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldAddrInfo *OldAddrInfo `json:"OldAddrInfo,omitempty" name:"OldAddrInfo"`

	// 正在进行的任务
	ProcessingTasks []*string `json:"ProcessingTasks,omitempty" name:"ProcessingTasks"`

	// 任务列表
	Tasks []*ObjectTask `json:"Tasks,omitempty" name:"Tasks"`

	// biz_net_service表id
	NetServiceId *int64 `json:"NetServiceId,omitempty" name:"NetServiceId"`
}

type DatabasePrivileges struct {
	// 数据库
	Db *string `json:"Db,omitempty" name:"Db"`

	// 权限列表
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`
}

type DatabaseTables struct {
	// 数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Database *string `json:"Database,omitempty" name:"Database"`

	// 表名称列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tables []*string `json:"Tables,omitempty" name:"Tables"`
}

type DbInfo struct {
	// 数据库名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 字符集类型
	CharacterSet *string `json:"CharacterSet,omitempty" name:"CharacterSet"`

	// 数据库状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 排序规则
	CollateRule *string `json:"CollateRule,omitempty" name:"CollateRule"`

	// 数据库备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 用户权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserHostPrivileges []*UserHostPrivilege `json:"UserHostPrivileges,omitempty" name:"UserHostPrivileges"`

	// 数据库ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbId *int64 `json:"DbId,omitempty" name:"DbId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 用户appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 用户Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 集群Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DbTable struct {
	// 数据库名称
	Db *string `json:"Db,omitempty" name:"Db"`

	// 数据库表名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`
}

// Predefined struct for user
type DeleteAccountsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 账号数组，包含account和host
	Accounts []*InputAccount `json:"Accounts,omitempty" name:"Accounts"`
}

type DeleteAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 账号数组，包含account和host
	Accounts []*InputAccount `json:"Accounts,omitempty" name:"Accounts"`
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
	delete(f, "ClusterId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountsResponseParams struct {
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
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 审计日志文件名称。
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

type DeleteAuditLogFileRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 审计日志文件名称。
	FileName *string `json:"FileName,omitempty" name:"FileName"`
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
	delete(f, "InstanceId")
	delete(f, "FileName")
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
type DeleteAuditRuleTemplatesRequestParams struct {
	// 审计规则模版ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`
}

type DeleteAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 审计规则模版ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`
}

func (r *DeleteAuditRuleTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditRuleTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuditRuleTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditRuleTemplatesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAuditRuleTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuditRuleTemplatesResponseParams `json:"Response"`
}

func (r *DeleteAuditRuleTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditRuleTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 备份文件ID，旧版本使用的字段，不推荐使用
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitempty" name:"SnapshotIdList"`

	// 备份文件ID，推荐使用
	BackupIds []*int64 `json:"BackupIds,omitempty" name:"BackupIds"`
}

type DeleteBackupRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 备份文件ID，旧版本使用的字段，不推荐使用
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitempty" name:"SnapshotIdList"`

	// 备份文件ID，推荐使用
	BackupIds []*int64 `json:"BackupIds,omitempty" name:"BackupIds"`
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
	delete(f, "ClusterId")
	delete(f, "SnapshotIdList")
	delete(f, "BackupIds")
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
type DeleteClusterDatabaseRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库名
	DbNames []*string `json:"DbNames,omitempty" name:"DbNames"`
}

type DeleteClusterDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库名
	DbNames []*string `json:"DbNames,omitempty" name:"DbNames"`
}

func (r *DeleteClusterDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DbNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterDatabaseResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteClusterDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterDatabaseResponseParams `json:"Response"`
}

func (r *DeleteClusterDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteParamTemplateRequestParams struct {
	// 参数模版ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DeleteParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 参数模版ID
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
type DescribeAccountAllGrantPrivilegesRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitempty" name:"Account"`
}

type DescribeAccountAllGrantPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitempty" name:"Account"`
}

func (r *DescribeAccountAllGrantPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountAllGrantPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Account")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountAllGrantPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountAllGrantPrivilegesResponseParams struct {
	// 权限语句
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivilegeStatements []*string `json:"PrivilegeStatements,omitempty" name:"PrivilegeStatements"`

	// 全局权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitempty" name:"GlobalPrivileges"`

	// 数据库权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabasePrivileges []*DatabasePrivileges `json:"DatabasePrivileges,omitempty" name:"DatabasePrivileges"`

	// 数据库表权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	TablePrivileges []*TablePrivileges `json:"TablePrivileges,omitempty" name:"TablePrivileges"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccountAllGrantPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountAllGrantPrivilegesResponseParams `json:"Response"`
}

func (r *DescribeAccountAllGrantPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountAllGrantPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountPrivilegesRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 账户名
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 主机
	Host *string `json:"Host,omitempty" name:"Host"`

	// 数据库名，为*时，忽略Type/TableName, 表示修改用户全局权限；
	Db *string `json:"Db,omitempty" name:"Db"`

	// 指定数据库下的对象类型，可选"table"，"*"
	Type *string `json:"Type,omitempty" name:"Type"`

	// 当Type="table"时，用来指定表名
	TableName *string `json:"TableName,omitempty" name:"TableName"`
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 账户名
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 主机
	Host *string `json:"Host,omitempty" name:"Host"`

	// 数据库名，为*时，忽略Type/TableName, 表示修改用户全局权限；
	Db *string `json:"Db,omitempty" name:"Db"`

	// 指定数据库下的对象类型，可选"table"，"*"
	Type *string `json:"Type,omitempty" name:"Type"`

	// 当Type="table"时，用来指定表名
	TableName *string `json:"TableName,omitempty" name:"TableName"`
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
	delete(f, "ClusterId")
	delete(f, "AccountName")
	delete(f, "Host")
	delete(f, "Db")
	delete(f, "Type")
	delete(f, "TableName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountPrivilegesResponseParams struct {
	// 权限列表，示例值为：["select","update","delete","create","drop","references","index","alter","show_db","create_tmp_table","lock_tables","execute","create_view","show_view","create_routine","alter_routine","event","trigger"]
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`

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
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 需要过滤的账户列表
	AccountNames []*string `json:"AccountNames,omitempty" name:"AccountNames"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	// 该参数已废用
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 需要过滤的账户列表
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`

	// 限制量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 需要过滤的账户列表
	AccountNames []*string `json:"AccountNames,omitempty" name:"AccountNames"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	// 该参数已废用
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 需要过滤的账户列表
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`

	// 限制量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	delete(f, "ClusterId")
	delete(f, "AccountNames")
	delete(f, "DbType")
	delete(f, "Hosts")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsResponseParams struct {
	// 数据库账号列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountSet []*Account `json:"AccountSet,omitempty" name:"AccountSet"`

	// 账号总数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

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
type DescribeAuditLogFilesRequestParams struct {
	// 实例ID
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
	
	// 实例ID
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
type DescribeAuditLogsRequestParams struct {
	// 实例ID
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

	// 分页参数，单次返回的数据条数。默认值为100，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeAuditLogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
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

	// 分页参数，单次返回的数据条数。默认值为100，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAuditLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogsRequest) FromJsonString(s string) error {
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
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogsResponseParams struct {
	// 符合条件的审计日志条数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 审计日志详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*AuditLog `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAuditLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditLogsResponseParams `json:"Response"`
}

func (r *DescribeAuditLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRuleTemplatesRequestParams struct {
	// 规则模版ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`

	// 规则模版名称
	RuleTemplateNames []*string `json:"RuleTemplateNames,omitempty" name:"RuleTemplateNames"`

	// 单次请求返回的数量。默认值20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认值为 0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 规则模版ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`

	// 规则模版名称
	RuleTemplateNames []*string `json:"RuleTemplateNames,omitempty" name:"RuleTemplateNames"`

	// 单次请求返回的数量。默认值20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认值为 0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAuditRuleTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRuleTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleTemplateIds")
	delete(f, "RuleTemplateNames")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditRuleTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRuleTemplatesResponseParams struct {
	// 符合查询条件的实例总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 规则模版详细信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*AuditRuleTemplateInfo `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAuditRuleTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditRuleTemplatesResponseParams `json:"Response"`
}

func (r *DescribeAuditRuleTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRuleTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRuleWithInstanceIdsRequestParams struct {
	// 实例ID。目前仅支持单个实例的查询。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DescribeAuditRuleWithInstanceIdsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。目前仅支持单个实例的查询。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeAuditRuleWithInstanceIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRuleWithInstanceIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditRuleWithInstanceIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRuleWithInstanceIdsResponseParams struct {
	// 无
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例审计规则信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*InstanceAuditRule `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAuditRuleWithInstanceIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditRuleWithInstanceIdsResponseParams `json:"Response"`
}

func (r *DescribeAuditRuleWithInstanceIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRuleWithInstanceIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupConfigRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
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
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupConfigResponseParams struct {
	// 表示全备开始时间，[0-24*3600]， 如0:00, 1:00, 2:00 分别为 0，3600， 7200
	BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitempty" name:"BackupTimeBeg"`

	// 表示全备开始时间，[0-24*3600]， 如0:00, 1:00, 2:00 分别为 0，3600， 7200
	BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitempty" name:"BackupTimeEnd"`

	// 表示保留备份时长, 单位秒，超过该时间将被清理, 七天表示为3600*24*7=604800
	ReserveDuration *uint64 `json:"ReserveDuration,omitempty" name:"ReserveDuration"`

	// 备份频率，长度为7的数组，分别对应周一到周日的备份方式，full-全量备份，increment-增量备份
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupFreq []*string `json:"BackupFreq,omitempty" name:"BackupFreq"`

	// 备份方式，logic-逻辑备份，snapshot-快照备份
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

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
type DescribeBackupDownloadUrlRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 备份ID
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`
}

type DescribeBackupDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 备份ID
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`
}

func (r *DescribeBackupDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "BackupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadUrlResponseParams struct {
	// 备份下载地址
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDownloadUrlResponseParams `json:"Response"`
}

func (r *DescribeBackupDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupListRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 备份文件列表大小，取值范围(0,100]
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 备份文件列表偏移，取值范围[0,INF)
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 备份ID
	BackupIds []*int64 `json:"BackupIds,omitempty" name:"BackupIds"`

	// 备份类型，可选值：snapshot，快照备份； logic，逻辑备份
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// 备份方式，可选值：auto，自动备份；manual，手动备
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// 快照类型，可选值：full，全量；increment，增量
	SnapShotType *string `json:"SnapShotType,omitempty" name:"SnapShotType"`

	// 备份开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 备份结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 备份文件名，模糊查询
	FileNames []*string `json:"FileNames,omitempty" name:"FileNames"`

	// 备份备注名，模糊查询
	BackupNames []*string `json:"BackupNames,omitempty" name:"BackupNames"`

	// 快照备份Id列表
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitempty" name:"SnapshotIdList"`
}

type DescribeBackupListRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 备份文件列表大小，取值范围(0,100]
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 备份文件列表偏移，取值范围[0,INF)
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 备份ID
	BackupIds []*int64 `json:"BackupIds,omitempty" name:"BackupIds"`

	// 备份类型，可选值：snapshot，快照备份； logic，逻辑备份
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// 备份方式，可选值：auto，自动备份；manual，手动备
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// 快照类型，可选值：full，全量；increment，增量
	SnapShotType *string `json:"SnapShotType,omitempty" name:"SnapShotType"`

	// 备份开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 备份结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 备份文件名，模糊查询
	FileNames []*string `json:"FileNames,omitempty" name:"FileNames"`

	// 备份备注名，模糊查询
	BackupNames []*string `json:"BackupNames,omitempty" name:"BackupNames"`

	// 快照备份Id列表
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitempty" name:"SnapshotIdList"`
}

func (r *DescribeBackupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "DbType")
	delete(f, "BackupIds")
	delete(f, "BackupType")
	delete(f, "BackupMethod")
	delete(f, "SnapShotType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "FileNames")
	delete(f, "BackupNames")
	delete(f, "SnapshotIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupListResponseParams struct {
	// 总共备份文件个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 备份文件列表
	BackupList []*BackupFileInfo `json:"BackupList,omitempty" name:"BackupList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupListResponseParams `json:"Response"`
}

func (r *DescribeBackupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogDownloadUrlRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Binlog文件ID
	BinlogId *int64 `json:"BinlogId,omitempty" name:"BinlogId"`
}

type DescribeBinlogDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Binlog文件ID
	BinlogId *int64 `json:"BinlogId,omitempty" name:"BinlogId"`
}

func (r *DescribeBinlogDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "BinlogId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogDownloadUrlResponseParams struct {
	// 下载地址
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBinlogDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBinlogDownloadUrlResponseParams `json:"Response"`
}

func (r *DescribeBinlogDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogSaveDaysRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeBinlogSaveDaysRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeBinlogSaveDaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogSaveDaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogSaveDaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogSaveDaysResponseParams struct {
	// Binlog保留天数
	BinlogSaveDays *int64 `json:"BinlogSaveDays,omitempty" name:"BinlogSaveDays"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBinlogSaveDaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBinlogSaveDaysResponseParams `json:"Response"`
}

func (r *DescribeBinlogSaveDaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogSaveDaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeBinlogsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制条数
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
	delete(f, "ClusterId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogsResponseParams struct {
	// 记录总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Binlog列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Binlogs []*BinlogItem `json:"Binlogs,omitempty" name:"Binlogs"`

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
type DescribeClusterDetailDatabasesRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认20,最大100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 数据库名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

type DescribeClusterDetailDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认20,最大100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 数据库名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

func (r *DescribeClusterDetailDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDetailDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DbName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterDetailDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDetailDatabasesResponseParams struct {
	// 数据库信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbInfos []*DbInfo `json:"DbInfos,omitempty" name:"DbInfos"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterDetailDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterDetailDatabasesResponseParams `json:"Response"`
}

func (r *DescribeClusterDetailDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDetailDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDetailRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeClusterDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDetailResponseParams struct {
	// 集群详细信息
	Detail *CynosdbClusterDetail `json:"Detail,omitempty" name:"Detail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterDetailResponseParams `json:"Response"`
}

func (r *DescribeClusterDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterInstanceGrpsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeClusterInstanceGrpsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterInstanceGrpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInstanceGrpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterInstanceGrpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterInstanceGrpsResponseParams struct {
	// 实例组个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例组列表
	InstanceGrpInfoList []*CynosdbInstanceGrp `json:"InstanceGrpInfoList,omitempty" name:"InstanceGrpInfoList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterInstanceGrpsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterInstanceGrpsResponseParams `json:"Response"`
}

func (r *DescribeClusterInstanceGrpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInstanceGrpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterParamLogsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID列表，用来记录具体操作哪些实例
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 排序字段，定义在回返结果的基于哪个字段进行排序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 定义具体的排序规则，限定为desc,asc,DESC,ASC其中之一
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 返回数量，默认为 20，取值范围为(0,100]
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0，取值范围为[0,INF)
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeClusterParamLogsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID列表，用来记录具体操作哪些实例
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 排序字段，定义在回返结果的基于哪个字段进行排序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 定义具体的排序规则，限定为desc,asc,DESC,ASC其中之一
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 返回数量，默认为 20，取值范围为(0,100]
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0，取值范围为[0,INF)
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeClusterParamLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterParamLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIds")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterParamLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterParamLogsResponseParams struct {
	// 记录总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 参数修改记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterParamLogs []*ClusterParamModifyLog `json:"ClusterParamLogs,omitempty" name:"ClusterParamLogs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterParamLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterParamLogsResponseParams `json:"Response"`
}

func (r *DescribeClusterParamLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterParamLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterParamsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 参数名字
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`
}

type DescribeClusterParamsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 参数名字
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`
}

func (r *DescribeClusterParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ParamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterParamsResponseParams struct {
	// 参数个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例参数列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*ParamInfo `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterParamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterParamsResponseParams `json:"Response"`
}

func (r *DescribeClusterParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterPasswordComplexityRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterPasswordComplexityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterPasswordComplexityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterPasswordComplexityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterPasswordComplexityResponseParams struct {
	// 数据字典参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidatePasswordDictionary *ParamInfo `json:"ValidatePasswordDictionary,omitempty" name:"ValidatePasswordDictionary"`

	// 密码长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidatePasswordLength *ParamInfo `json:"ValidatePasswordLength,omitempty" name:"ValidatePasswordLength"`

	// 大小写敏感字符个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidatePasswordMixedCaseCount *ParamInfo `json:"ValidatePasswordMixedCaseCount,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// 数字个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidatePasswordNumberCount *ParamInfo `json:"ValidatePasswordNumberCount,omitempty" name:"ValidatePasswordNumberCount"`

	// 密码等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidatePasswordPolicy *ParamInfo `json:"ValidatePasswordPolicy,omitempty" name:"ValidatePasswordPolicy"`

	// 特殊字符个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidatePasswordSpecialCharCount *ParamInfo `json:"ValidatePasswordSpecialCharCount,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterPasswordComplexityResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterPasswordComplexityResponseParams `json:"Response"`
}

func (r *DescribeClusterPasswordComplexityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterPasswordComplexityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersRequestParams struct {
	// 引擎类型：目前支持“MYSQL”， “POSTGRESQL”
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，取值范围：
	// <li> CREATETIME：创建时间</li>
	// <li> PERIODENDTIME：过期时间</li>
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型，取值范围：
	// <li> ASC：升序排序 </li>
	// <li> DESC：降序排序 </li>
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 搜索条件，若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Filters []*QueryFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
	
	// 引擎类型：目前支持“MYSQL”， “POSTGRESQL”
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，取值范围：
	// <li> CREATETIME：创建时间</li>
	// <li> PERIODENDTIME：过期时间</li>
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型，取值范围：
	// <li> ASC：升序排序 </li>
	// <li> DESC：降序排序 </li>
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 搜索条件，若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Filters []*QueryFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DbType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersResponseParams struct {
	// 集群数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 集群列表
	ClusterSet []*CynosdbCluster `json:"ClusterSet,omitempty" name:"ClusterSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClustersResponseParams `json:"Response"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsRequestParams struct {
	// 实例组ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例组ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	// 安全组信息
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
type DescribeFlowRequestParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

type DescribeFlowRequest struct {
	*tchttp.BaseRequest
	
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
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
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowResponseParams struct {
	// 任务流状态。0-成功，1-失败，2-处理中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeInstanceDetailRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDetailResponseParams struct {
	// 实例详情
	Detail *CynosdbInstanceDetail `json:"Detail,omitempty" name:"Detail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceDetailResponseParams `json:"Response"`
}

func (r *DescribeInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceErrorLogsRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 日志条数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 日志条数偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 排序字段，有Timestamp枚举值
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型，有ASC,DESC枚举值
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 日志等级，有error、warning、note三种，支持多个等级同时搜索
	LogLevels []*string `json:"LogLevels,omitempty" name:"LogLevels"`

	// 关键字，支持模糊搜索
	KeyWords []*string `json:"KeyWords,omitempty" name:"KeyWords"`
}

type DescribeInstanceErrorLogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 日志条数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 日志条数偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 排序字段，有Timestamp枚举值
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型，有ASC,DESC枚举值
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 日志等级，有error、warning、note三种，支持多个等级同时搜索
	LogLevels []*string `json:"LogLevels,omitempty" name:"LogLevels"`

	// 关键字，支持模糊搜索
	KeyWords []*string `json:"KeyWords,omitempty" name:"KeyWords"`
}

func (r *DescribeInstanceErrorLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceErrorLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "LogLevels")
	delete(f, "KeyWords")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceErrorLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceErrorLogsResponseParams struct {
	// 日志条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 错误日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorLogs []*CynosdbErrorLogItem `json:"ErrorLogs,omitempty" name:"ErrorLogs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceErrorLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceErrorLogsResponseParams `json:"Response"`
}

func (r *DescribeInstanceErrorLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceErrorLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID，支持批量查询
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 参数名搜索条件，支持模糊匹配
	ParamKeyword *string `json:"ParamKeyword,omitempty" name:"ParamKeyword"`
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID，支持批量查询
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 参数名搜索条件，支持模糊匹配
	ParamKeyword *string `json:"ParamKeyword,omitempty" name:"ParamKeyword"`
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
	delete(f, "ClusterId")
	delete(f, "InstanceIds")
	delete(f, "ParamKeyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsResponseParams struct {
	// 实例参数列表
	Items []*InstanceParamItem `json:"Items,omitempty" name:"Items"`

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
type DescribeInstanceSlowQueriesRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 事务开始最早时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 事务开始最晚时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 客户端host
	Host *string `json:"Host,omitempty" name:"Host"`

	// 数据库名
	Database *string `json:"Database,omitempty" name:"Database"`

	// 排序字段，可选值：QueryTime,LockTime,RowsExamined,RowsSent
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型，可选值：asc,desc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

type DescribeInstanceSlowQueriesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 事务开始最早时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 事务开始最晚时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 客户端host
	Host *string `json:"Host,omitempty" name:"Host"`

	// 数据库名
	Database *string `json:"Database,omitempty" name:"Database"`

	// 排序字段，可选值：QueryTime,LockTime,RowsExamined,RowsSent
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型，可选值：asc,desc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeInstanceSlowQueriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSlowQueriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Username")
	delete(f, "Host")
	delete(f, "Database")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSlowQueriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSlowQueriesResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 慢查询记录
	SlowQueries []*SlowQueriesItem `json:"SlowQueries,omitempty" name:"SlowQueries"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceSlowQueriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceSlowQueriesResponseParams `json:"Response"`
}

func (r *DescribeInstanceSlowQueriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSlowQueriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSpecsRequestParams struct {
	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 是否需要返回可用区信息
	IncludeZoneStocks *bool `json:"IncludeZoneStocks,omitempty" name:"IncludeZoneStocks"`
}

type DescribeInstanceSpecsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 是否需要返回可用区信息
	IncludeZoneStocks *bool `json:"IncludeZoneStocks,omitempty" name:"IncludeZoneStocks"`
}

func (r *DescribeInstanceSpecsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSpecsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DbType")
	delete(f, "IncludeZoneStocks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSpecsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSpecsResponseParams struct {
	// 规格信息
	InstanceSpecSet []*InstanceSpec `json:"InstanceSpecSet,omitempty" name:"InstanceSpecSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceSpecsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceSpecsResponseParams `json:"Response"`
}

func (r *DescribeInstanceSpecsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSpecsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，取值范围：
	// <li> CREATETIME：创建时间</li>
	// <li> PERIODENDTIME：过期时间</li>
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型，取值范围：
	// <li> ASC：升序排序 </li>
	// <li> DESC：降序排序 </li>
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 搜索条件，若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Filters []*QueryFilter `json:"Filters,omitempty" name:"Filters"`

	// 引擎类型：目前支持“MYSQL”， “POSTGRESQL”
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 实例状态, 可选值:
	// creating 创建中
	// running 运行中
	// isolating 隔离中
	// isolated 已隔离
	// activating 恢复中
	// offlining 下线中
	// offlined 已下线
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例id列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，取值范围：
	// <li> CREATETIME：创建时间</li>
	// <li> PERIODENDTIME：过期时间</li>
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型，取值范围：
	// <li> ASC：升序排序 </li>
	// <li> DESC：降序排序 </li>
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 搜索条件，若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Filters []*QueryFilter `json:"Filters,omitempty" name:"Filters"`

	// 引擎类型：目前支持“MYSQL”， “POSTGRESQL”
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 实例状态, 可选值:
	// creating 创建中
	// running 运行中
	// isolating 隔离中
	// isolated 已隔离
	// activating 恢复中
	// offlining 下线中
	// offlined 已下线
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例id列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
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
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Filters")
	delete(f, "DbType")
	delete(f, "Status")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 实例个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例列表
	InstanceSet []*CynosdbInstance `json:"InstanceSet,omitempty" name:"InstanceSet"`

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
type DescribeMaintainPeriodRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeMaintainPeriodRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeMaintainPeriodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintainPeriodRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMaintainPeriodRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaintainPeriodResponseParams struct {
	// 维护week days
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitempty" name:"MaintainWeekDays"`

	// 维护开始时间，单位秒
	MaintainStartTime *int64 `json:"MaintainStartTime,omitempty" name:"MaintainStartTime"`

	// 维护时长，单位秒
	MaintainDuration *int64 `json:"MaintainDuration,omitempty" name:"MaintainDuration"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMaintainPeriodResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMaintainPeriodResponseParams `json:"Response"`
}

func (r *DescribeMaintainPeriodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintainPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplateDetailRequestParams struct {
	// 参数模板ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DescribeParamTemplateDetailRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeParamTemplateDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplateDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplateDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplateDetailResponseParams struct {
	// 参数模板ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 参数模板名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 参数模板描述
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`

	// 引擎版本
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 参数总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 参数列表
	Items []*ParamDetail `json:"Items,omitempty" name:"Items"`

	// 数据库类型，可选值：NORMAL，SERVERLESS
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeParamTemplateDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParamTemplateDetailResponseParams `json:"Response"`
}

func (r *DescribeParamTemplateDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplateDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplatesRequestParams struct {
	// 数据库引擎版本号
	EngineVersions []*string `json:"EngineVersions,omitempty" name:"EngineVersions"`

	// 模版名称
	TemplateNames []*string `json:"TemplateNames,omitempty" name:"TemplateNames"`

	// 模版ID
	TemplateIds []*int64 `json:"TemplateIds,omitempty" name:"TemplateIds"`

	// 数据库类型，可选值：NORMAL，SERVERLESS
	DbModes []*string `json:"DbModes,omitempty" name:"DbModes"`

	// 查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询的模板对应的产品类型
	Products []*string `json:"Products,omitempty" name:"Products"`

	// 模版类型
	TemplateTypes []*string `json:"TemplateTypes,omitempty" name:"TemplateTypes"`

	// 版本类型
	EngineTypes []*string `json:"EngineTypes,omitempty" name:"EngineTypes"`

	// 返回结果的排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式（asc、desc）
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeParamTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 数据库引擎版本号
	EngineVersions []*string `json:"EngineVersions,omitempty" name:"EngineVersions"`

	// 模版名称
	TemplateNames []*string `json:"TemplateNames,omitempty" name:"TemplateNames"`

	// 模版ID
	TemplateIds []*int64 `json:"TemplateIds,omitempty" name:"TemplateIds"`

	// 数据库类型，可选值：NORMAL，SERVERLESS
	DbModes []*string `json:"DbModes,omitempty" name:"DbModes"`

	// 查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询的模板对应的产品类型
	Products []*string `json:"Products,omitempty" name:"Products"`

	// 模版类型
	TemplateTypes []*string `json:"TemplateTypes,omitempty" name:"TemplateTypes"`

	// 版本类型
	EngineTypes []*string `json:"EngineTypes,omitempty" name:"EngineTypes"`

	// 返回结果的排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式（asc、desc）
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
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
	delete(f, "TemplateNames")
	delete(f, "TemplateIds")
	delete(f, "DbModes")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Products")
	delete(f, "TemplateTypes")
	delete(f, "EngineTypes")
	delete(f, "OrderBy")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplatesResponseParams struct {
	// 参数模板数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 参数模板信息
	Items []*ParamTemplateListInfo `json:"Items,omitempty" name:"Items"`

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
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 限制量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 搜索关键字
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 限制量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 搜索关键字
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
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
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsResponseParams struct {
	// 安全组详情
	Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`

	// 总数量
	Total *int64 `json:"Total,omitempty" name:"Total"`

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
type DescribeProxiesRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，取值范围：
	// <li> CREATETIME：创建时间</li>
	// <li> PERIODENDTIME：过期时间</li>
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型，取值范围：
	// <li> ASC：升序排序 </li>
	// <li> DESC：降序排序 </li>
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 搜索条件，若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Filters []*QueryParamFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeProxiesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，取值范围：
	// <li> CREATETIME：创建时间</li>
	// <li> PERIODENDTIME：过期时间</li>
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型，取值范围：
	// <li> ASC：升序排序 </li>
	// <li> DESC：降序排序 </li>
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 搜索条件，若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Filters []*QueryParamFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeProxiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxiesResponseParams struct {
	// 数据库代理组数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据库代理组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyGroupInfos []*ProxyGroupInfo `json:"ProxyGroupInfos,omitempty" name:"ProxyGroupInfos"`

	// 数据库代理节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyNodeInfos []*ProxyNodeInfo `json:"ProxyNodeInfos,omitempty" name:"ProxyNodeInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProxiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxiesResponseParams `json:"Response"`
}

func (r *DescribeProxiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxyNodesRequestParams struct {
	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，取值范围：
	// <li> CREATETIME：创建时间</li>
	// <li> PERIODENDTIME：过期时间</li>
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型，取值范围：
	// <li> ASC：升序排序 </li>
	// <li> DESC：降序排序 </li>
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 搜索条件，若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Filters []*QueryFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeProxyNodesRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，取值范围：
	// <li> CREATETIME：创建时间</li>
	// <li> PERIODENDTIME：过期时间</li>
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型，取值范围：
	// <li> ASC：升序排序 </li>
	// <li> DESC：降序排序 </li>
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 搜索条件，若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Filters []*QueryFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeProxyNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxyNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxyNodesResponseParams struct {
	// 数据库代理节点总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据库代理节点列表
	ProxyNodeInfos []*ProxyNodeInfo `json:"ProxyNodeInfos,omitempty" name:"ProxyNodeInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProxyNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxyNodesResponseParams `json:"Response"`
}

func (r *DescribeProxyNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySpecsRequestParams struct {

}

type DescribeProxySpecsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeProxySpecsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySpecsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxySpecsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySpecsResponseParams struct {
	// 数据库代理规格列表
	ProxySpecs []*ProxySpec `json:"ProxySpecs,omitempty" name:"ProxySpecs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProxySpecsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxySpecsResponseParams `json:"Response"`
}

func (r *DescribeProxySpecsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySpecsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageDetailRequestParams struct {
	// 资源包唯一ID
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 实例ID
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 偏移量
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// 限制
	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

type DescribeResourcePackageDetailRequest struct {
	*tchttp.BaseRequest
	
	// 资源包唯一ID
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 实例ID
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 偏移量
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// 限制
	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeResourcePackageDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePackageDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageId")
	delete(f, "ClusterIds")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcePackageDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageDetailResponseParams struct {
	// 总使用明细数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 资源包明细说明
	Detail []*PackageDetail `json:"Detail,omitempty" name:"Detail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourcePackageDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcePackageDetailResponseParams `json:"Response"`
}

func (r *DescribeResourcePackageDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePackageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageListRequestParams struct {
	// 资源包唯一ID
	PackageId []*string `json:"PackageId,omitempty" name:"PackageId"`

	// 资源包名称
	PackageName []*string `json:"PackageName,omitempty" name:"PackageName"`

	// 资源包类型
	// CCU-计算资源包，DISK-存储资源包
	PackageType []*string `json:"PackageType,omitempty" name:"PackageType"`

	// 资源包使用地域
	// china-中国内地通用，overseas-港澳台及海外通用
	PackageRegion []*string `json:"PackageRegion,omitempty" name:"PackageRegion"`

	// 资源包状态
	// creating-创建中；
	// using-使用中；
	// expired-已过期；
	// normal_finish-使用完；
	// apply_refund-申请退费中；
	// refund-已退费。
	Status []*string `json:"Status,omitempty" name:"Status"`

	// 排序条件，支持排序条件:startTime-生效时间，
	// expireTime-过期时间，packageUsedSpec-使用容量，packageTotalSpec-总存储量。
	// 按照数组顺序排列；
	OrderBy []*string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，DESC-降序，ASC-升序
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeResourcePackageListRequest struct {
	*tchttp.BaseRequest
	
	// 资源包唯一ID
	PackageId []*string `json:"PackageId,omitempty" name:"PackageId"`

	// 资源包名称
	PackageName []*string `json:"PackageName,omitempty" name:"PackageName"`

	// 资源包类型
	// CCU-计算资源包，DISK-存储资源包
	PackageType []*string `json:"PackageType,omitempty" name:"PackageType"`

	// 资源包使用地域
	// china-中国内地通用，overseas-港澳台及海外通用
	PackageRegion []*string `json:"PackageRegion,omitempty" name:"PackageRegion"`

	// 资源包状态
	// creating-创建中；
	// using-使用中；
	// expired-已过期；
	// normal_finish-使用完；
	// apply_refund-申请退费中；
	// refund-已退费。
	Status []*string `json:"Status,omitempty" name:"Status"`

	// 排序条件，支持排序条件:startTime-生效时间，
	// expireTime-过期时间，packageUsedSpec-使用容量，packageTotalSpec-总存储量。
	// 按照数组顺序排列；
	OrderBy []*string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，DESC-降序，ASC-升序
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeResourcePackageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePackageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageId")
	delete(f, "PackageName")
	delete(f, "PackageType")
	delete(f, "PackageRegion")
	delete(f, "Status")
	delete(f, "OrderBy")
	delete(f, "OrderDirection")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcePackageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageListResponseParams struct {
	// 总配置数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 资源包明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*Package `json:"Detail,omitempty" name:"Detail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourcePackageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcePackageListResponseParams `json:"Response"`
}

func (r *DescribeResourcePackageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePackageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageSaleSpecRequestParams struct {
	// 实例类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 资源包使用地域
	// china-中国内地通用，overseas-港澳台及海外通用
	PackageRegion *string `json:"PackageRegion,omitempty" name:"PackageRegion"`

	// 资源包类型
	// CCU-计算资源包
	// DISK-存储资源包
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeResourcePackageSaleSpecRequest struct {
	*tchttp.BaseRequest
	
	// 实例类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 资源包使用地域
	// china-中国内地通用，overseas-港澳台及海外通用
	PackageRegion *string `json:"PackageRegion,omitempty" name:"PackageRegion"`

	// 资源包类型
	// CCU-计算资源包
	// DISK-存储资源包
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeResourcePackageSaleSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePackageSaleSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "PackageRegion")
	delete(f, "PackageType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcePackageSaleSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageSaleSpecResponseParams struct {
	// 可售卖资源包规格总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 资源包明细说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*SalePackageSpec `json:"Detail,omitempty" name:"Detail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourcePackageSaleSpecResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcePackageSaleSpecResponseParams `json:"Response"`
}

func (r *DescribeResourcePackageSaleSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePackageSaleSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesByDealNameRequestParams struct {
	// 计费订单ID（如果计费还没回调业务发货，可能出现错误码InvalidParameterValue.DealNameNotFound，这种情况需要业务重试DescribeResourcesByDealName接口直到成功）
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 计费订单ID列表，可以一次查询若干条订单ID对应资源信息（如果计费还没回调业务发货，可能出现错误码InvalidParameterValue.DealNameNotFound，这种情况需要业务重试DescribeResourcesByDealName接口直到成功）
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
}

type DescribeResourcesByDealNameRequest struct {
	*tchttp.BaseRequest
	
	// 计费订单ID（如果计费还没回调业务发货，可能出现错误码InvalidParameterValue.DealNameNotFound，这种情况需要业务重试DescribeResourcesByDealName接口直到成功）
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 计费订单ID列表，可以一次查询若干条订单ID对应资源信息（如果计费还没回调业务发货，可能出现错误码InvalidParameterValue.DealNameNotFound，这种情况需要业务重试DescribeResourcesByDealName接口直到成功）
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
}

func (r *DescribeResourcesByDealNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByDealNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealName")
	delete(f, "DealNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcesByDealNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesByDealNameResponseParams struct {
	// 计费资源id信息数组
	BillingResourceInfos []*BillingResourceInfo `json:"BillingResourceInfos,omitempty" name:"BillingResourceInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourcesByDealNameResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcesByDealNameResponseParams `json:"Response"`
}

func (r *DescribeResourcesByDealNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByDealNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollbackTimeRangeRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeRollbackTimeRangeRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeRollbackTimeRangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTimeRangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRollbackTimeRangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollbackTimeRangeResponseParams struct {
	// 有效回归时间范围开始时间点（已废弃）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeRangeStart *string `json:"TimeRangeStart,omitempty" name:"TimeRangeStart"`

	// 有效回归时间范围结束时间点（已废弃）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeRangeEnd *string `json:"TimeRangeEnd,omitempty" name:"TimeRangeEnd"`

	// 可回档时间范围
	RollbackTimeRanges []*RollbackTimeRange `json:"RollbackTimeRanges,omitempty" name:"RollbackTimeRanges"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRollbackTimeRangeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRollbackTimeRangeResponseParams `json:"Response"`
}

func (r *DescribeRollbackTimeRangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTimeRangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollbackTimeValidityRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 期望回滚的时间点
	ExpectTime *string `json:"ExpectTime,omitempty" name:"ExpectTime"`

	// 回滚时间点的允许误差范围
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitempty" name:"ExpectTimeThresh"`
}

type DescribeRollbackTimeValidityRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 期望回滚的时间点
	ExpectTime *string `json:"ExpectTime,omitempty" name:"ExpectTime"`

	// 回滚时间点的允许误差范围
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitempty" name:"ExpectTimeThresh"`
}

func (r *DescribeRollbackTimeValidityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTimeValidityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ExpectTime")
	delete(f, "ExpectTimeThresh")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRollbackTimeValidityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollbackTimeValidityResponseParams struct {
	// 存储poolID
	PoolId *uint64 `json:"PoolId,omitempty" name:"PoolId"`

	// 回滚任务ID，后续按该时间点回滚时，需要传入
	QueryId *uint64 `json:"QueryId,omitempty" name:"QueryId"`

	// 时间点是否有效：pass，检测通过；fail，检测失败
	Status *string `json:"Status,omitempty" name:"Status"`

	// 建议时间点，在Status为fail时，该值才有效
	SuggestTime *string `json:"SuggestTime,omitempty" name:"SuggestTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRollbackTimeValidityResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRollbackTimeValidityResponseParams `json:"Response"`
}

func (r *DescribeRollbackTimeValidityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTimeValidityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportProxyVersionRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`
}

type DescribeSupportProxyVersionRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`
}

func (r *DescribeSupportProxyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportProxyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProxyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSupportProxyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportProxyVersionResponseParams struct {
	// 支持的数据库代理版本集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportProxyVersions []*string `json:"SupportProxyVersions,omitempty" name:"SupportProxyVersions"`

	// 当前proxy版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitempty" name:"CurrentProxyVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSupportProxyVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSupportProxyVersionResponseParams `json:"Response"`
}

func (r *DescribeSupportProxyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportProxyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesRequestParams struct {
	// 是否包含虚拟区
	IncludeVirtualZones *bool `json:"IncludeVirtualZones,omitempty" name:"IncludeVirtualZones"`

	// 是否展示地域下所有可用区，并显示用户每个可用区权限
	ShowPermission *bool `json:"ShowPermission,omitempty" name:"ShowPermission"`
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
	// 是否包含虚拟区
	IncludeVirtualZones *bool `json:"IncludeVirtualZones,omitempty" name:"IncludeVirtualZones"`

	// 是否展示地域下所有可用区，并显示用户每个可用区权限
	ShowPermission *bool `json:"ShowPermission,omitempty" name:"ShowPermission"`
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
	delete(f, "IncludeVirtualZones")
	delete(f, "ShowPermission")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// 地域信息
	RegionSet []*SaleRegion `json:"RegionSet,omitempty" name:"RegionSet"`

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
type DisassociateSecurityGroupsRequestParams struct {
	// 实例组ID数组
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 要修改的安全组ID列表，一个或者多个安全组ID组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例组ID数组
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 要修改的安全组ID列表，一个或者多个安全组ID组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`
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
	delete(f, "InstanceIds")
	delete(f, "SecurityGroupIds")
	delete(f, "Zone")
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

type ErrorLogItemExport struct {
	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`

	// 日志等级，可选值note, warning，error
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *string `json:"Level,omitempty" name:"Level"`

	// 日志内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`
}

// Predefined struct for user
type ExportInstanceErrorLogsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 日志最早时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 日志最晚时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 日志等级
	LogLevels []*string `json:"LogLevels,omitempty" name:"LogLevels"`

	// 关键字
	KeyWords []*string `json:"KeyWords,omitempty" name:"KeyWords"`

	// 文件类型，可选值：csv, original
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 可选值Timestamp
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// ASC或DESC
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

type ExportInstanceErrorLogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 日志最早时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 日志最晚时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 日志等级
	LogLevels []*string `json:"LogLevels,omitempty" name:"LogLevels"`

	// 关键字
	KeyWords []*string `json:"KeyWords,omitempty" name:"KeyWords"`

	// 文件类型，可选值：csv, original
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 可选值Timestamp
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// ASC或DESC
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *ExportInstanceErrorLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportInstanceErrorLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "LogLevels")
	delete(f, "KeyWords")
	delete(f, "FileType")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportInstanceErrorLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportInstanceErrorLogsResponseParams struct {
	// 错误日志导出内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorLogItems []*ErrorLogItemExport `json:"ErrorLogItems,omitempty" name:"ErrorLogItems"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportInstanceErrorLogsResponse struct {
	*tchttp.BaseResponse
	Response *ExportInstanceErrorLogsResponseParams `json:"Response"`
}

func (r *ExportInstanceErrorLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportInstanceErrorLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportInstanceSlowQueriesRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 事务开始最早时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 事务开始最晚时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 客户端host
	Host *string `json:"Host,omitempty" name:"Host"`

	// 数据库名
	Database *string `json:"Database,omitempty" name:"Database"`

	// 文件类型，可选值：csv, original
	FileType *string `json:"FileType,omitempty" name:"FileType"`
}

type ExportInstanceSlowQueriesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 事务开始最早时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 事务开始最晚时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 客户端host
	Host *string `json:"Host,omitempty" name:"Host"`

	// 数据库名
	Database *string `json:"Database,omitempty" name:"Database"`

	// 文件类型，可选值：csv, original
	FileType *string `json:"FileType,omitempty" name:"FileType"`
}

func (r *ExportInstanceSlowQueriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportInstanceSlowQueriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Username")
	delete(f, "Host")
	delete(f, "Database")
	delete(f, "FileType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportInstanceSlowQueriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportInstanceSlowQueriesResponseParams struct {
	// 慢查询导出内容
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportInstanceSlowQueriesResponse struct {
	*tchttp.BaseResponse
	Response *ExportInstanceSlowQueriesResponseParams `json:"Response"`
}

func (r *ExportInstanceSlowQueriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportInstanceSlowQueriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GrantAccountPrivilegesRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitempty" name:"Account"`

	// 数据库表权限码数组
	DbTablePrivileges []*string `json:"DbTablePrivileges,omitempty" name:"DbTablePrivileges"`

	// 数据库表信息
	DbTables []*DbTable `json:"DbTables,omitempty" name:"DbTables"`
}

type GrantAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitempty" name:"Account"`

	// 数据库表权限码数组
	DbTablePrivileges []*string `json:"DbTablePrivileges,omitempty" name:"DbTablePrivileges"`

	// 数据库表信息
	DbTables []*DbTable `json:"DbTables,omitempty" name:"DbTables"`
}

func (r *GrantAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Account")
	delete(f, "DbTablePrivileges")
	delete(f, "DbTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GrantAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GrantAccountPrivilegesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GrantAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *GrantAccountPrivilegesResponseParams `json:"Response"`
}

func (r *GrantAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InputAccount struct {
	// 账号
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 主机，默认‘%’
	Host *string `json:"Host,omitempty" name:"Host"`
}

// Predefined struct for user
type InquirePriceCreateRequestParams struct {
	// 可用区,每个地域提供最佳实践
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 购买计算节点个数
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 实例购买类型，可选值为：PREPAID, POSTPAID, SERVERLESS
	InstancePayMode *string `json:"InstancePayMode,omitempty" name:"InstancePayMode"`

	// 存储购买类型，可选值为：PREPAID, POSTPAID
	StoragePayMode *string `json:"StoragePayMode,omitempty" name:"StoragePayMode"`

	// CPU核数，PREPAID与POSTPAID实例类型必传
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存大小，单位G，PREPAID与POSTPAID实例类型必传
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Ccu大小，serverless类型必传
	Ccu *float64 `json:"Ccu,omitempty" name:"Ccu"`

	// 存储大小，PREPAID存储类型必传
	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// 购买时长，PREPAID购买类型必传
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 时长单位，可选值为：m,d。PREPAID购买类型必传
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
}

type InquirePriceCreateRequest struct {
	*tchttp.BaseRequest
	
	// 可用区,每个地域提供最佳实践
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 购买计算节点个数
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 实例购买类型，可选值为：PREPAID, POSTPAID, SERVERLESS
	InstancePayMode *string `json:"InstancePayMode,omitempty" name:"InstancePayMode"`

	// 存储购买类型，可选值为：PREPAID, POSTPAID
	StoragePayMode *string `json:"StoragePayMode,omitempty" name:"StoragePayMode"`

	// CPU核数，PREPAID与POSTPAID实例类型必传
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存大小，单位G，PREPAID与POSTPAID实例类型必传
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Ccu大小，serverless类型必传
	Ccu *float64 `json:"Ccu,omitempty" name:"Ccu"`

	// 存储大小，PREPAID存储类型必传
	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// 购买时长，PREPAID购买类型必传
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 时长单位，可选值为：m,d。PREPAID购买类型必传
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
}

func (r *InquirePriceCreateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "GoodsNum")
	delete(f, "InstancePayMode")
	delete(f, "StoragePayMode")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "Ccu")
	delete(f, "StorageLimit")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceCreateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateResponseParams struct {
	// 实例价格
	InstancePrice *TradePrice `json:"InstancePrice,omitempty" name:"InstancePrice"`

	// 存储价格
	StoragePrice *TradePrice `json:"StoragePrice,omitempty" name:"StoragePrice"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquirePriceCreateResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceCreateResponseParams `json:"Response"`
}

func (r *InquirePriceCreateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 购买时长,与TimeUnit组合才能生效
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 购买时长单位, 与TimeSpan组合生效，可选:日:d,月:m
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
}

type InquirePriceRenewRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 购买时长,与TimeUnit组合才能生效
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 购买时长单位, 与TimeSpan组合生效，可选:日:d,月:m
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
}

func (r *InquirePriceRenewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceRenewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewResponseParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 对应的询价结果数组
	Prices []*TradePrice `json:"Prices,omitempty" name:"Prices"`

	// 续费计算节点的总价格
	InstanceRealTotalPrice *int64 `json:"InstanceRealTotalPrice,omitempty" name:"InstanceRealTotalPrice"`

	// 续费存储节点的总价格
	StorageRealTotalPrice *int64 `json:"StorageRealTotalPrice,omitempty" name:"StorageRealTotalPrice"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquirePriceRenewResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceRenewResponseParams `json:"Response"`
}

func (r *InquirePriceRenewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceAuditRule struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 是否是规则审计。true-规则审计，false-全审计。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuditRule *bool `json:"AuditRule,omitempty" name:"AuditRule"`

	// 审计规则详情。仅当AuditRule=true时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitempty" name:"AuditRuleFilters"`
}

type InstanceInitInfo struct {
	// 实例cpu
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 实例内存
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例类型 rw/ro
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例个数,范围[1,15]
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
}

type InstanceNetInfo struct {
	// 网络类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceGroupType *string `json:"InstanceGroupType,omitempty" name:"InstanceGroupType"`

	// 实例组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 网络类型, 0-基础网络, 1-vpc网络, 2-黑石网络
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetType *int64 `json:"NetType,omitempty" name:"NetType"`

	// 私有网络IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 私有网络端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 外网域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`

	// 外网Ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanIP *string `json:"WanIP,omitempty" name:"WanIP"`

	// 外网端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`

	// 外网开启状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanStatus *string `json:"WanStatus,omitempty" name:"WanStatus"`
}

type InstanceParamItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例参数列表
	ParamsItems []*ParamItemDetail `json:"ParamsItems,omitempty" name:"ParamsItems"`
}

type InstanceSpec struct {
	// 实例CPU，单位：核
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 实例内存，单位：GB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 实例最大可用存储，单位：GB
	MaxStorageSize *uint64 `json:"MaxStorageSize,omitempty" name:"MaxStorageSize"`

	// 实例最小可用存储，单位：GB
	MinStorageSize *uint64 `json:"MinStorageSize,omitempty" name:"MinStorageSize"`

	// 是否有库存
	HasStock *bool `json:"HasStock,omitempty" name:"HasStock"`

	// 机器类型
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 最大IOPS
	MaxIops *int64 `json:"MaxIops,omitempty" name:"MaxIops"`

	// 最大IO带宽
	MaxIoBandWidth *int64 `json:"MaxIoBandWidth,omitempty" name:"MaxIoBandWidth"`

	// 地域库存信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneStockInfos []*ZoneStockInfo `json:"ZoneStockInfos,omitempty" name:"ZoneStockInfos"`

	// 库存数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	StockCount *int64 `json:"StockCount,omitempty" name:"StockCount"`
}

// Predefined struct for user
type IsolateClusterRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 该参数已废用
	DbType *string `json:"DbType,omitempty" name:"DbType"`
}

type IsolateClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 该参数已废用
	DbType *string `json:"DbType,omitempty" name:"DbType"`
}

func (r *IsolateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DbType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateClusterResponseParams struct {
	// 任务流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 退款订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type IsolateClusterResponse struct {
	*tchttp.BaseResponse
	Response *IsolateClusterResponseParams `json:"Response"`
}

func (r *IsolateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateInstanceRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID数组
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// 该参数已废弃
	DbType *string `json:"DbType,omitempty" name:"DbType"`
}

type IsolateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID数组
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// 该参数已废弃
	DbType *string `json:"DbType,omitempty" name:"DbType"`
}

func (r *IsolateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdList")
	delete(f, "DbType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateInstanceResponseParams struct {
	// 任务流id
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 隔离实例的订单id（预付费实例）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type IsolateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateInstanceResponseParams `json:"Response"`
}

func (r *IsolateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifiableInfo struct {

}

// Predefined struct for user
type ModifyAccountDescriptionRequestParams struct {
	// 数据库账号名
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 数据库账号描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 主机，默认为"%"
	Host *string `json:"Host,omitempty" name:"Host"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// 数据库账号名
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 数据库账号描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 主机，默认为"%"
	Host *string `json:"Host,omitempty" name:"Host"`
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
	delete(f, "AccountName")
	delete(f, "Description")
	delete(f, "ClusterId")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountDescriptionResponseParams struct {
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
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 新主机
	NewHost *string `json:"NewHost,omitempty" name:"NewHost"`

	// 账号信息
	Account *InputAccount `json:"Account,omitempty" name:"Account"`
}

type ModifyAccountHostRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 新主机
	NewHost *string `json:"NewHost,omitempty" name:"NewHost"`

	// 账号信息
	Account *InputAccount `json:"Account,omitempty" name:"Account"`
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
	delete(f, "ClusterId")
	delete(f, "NewHost")
	delete(f, "Account")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountHostResponseParams struct {
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
type ModifyAccountParamsRequestParams struct {
	// 集群id，不超过32个字符
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitempty" name:"Account"`

	// 数据库表权限数组,当前仅支持参数：max_user_connections，max_user_connections不能大于10240
	AccountParams []*AccountParam `json:"AccountParams,omitempty" name:"AccountParams"`
}

type ModifyAccountParamsRequest struct {
	*tchttp.BaseRequest
	
	// 集群id，不超过32个字符
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitempty" name:"Account"`

	// 数据库表权限数组,当前仅支持参数：max_user_connections，max_user_connections不能大于10240
	AccountParams []*AccountParam `json:"AccountParams,omitempty" name:"AccountParams"`
}

func (r *ModifyAccountParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Account")
	delete(f, "AccountParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountParamsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAccountParamsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountParamsResponseParams `json:"Response"`
}

func (r *ModifyAccountParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPrivilegesRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitempty" name:"Account"`

	// 全局权限数组
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitempty" name:"GlobalPrivileges"`

	// 数据库权限数组
	DatabasePrivileges []*DatabasePrivileges `json:"DatabasePrivileges,omitempty" name:"DatabasePrivileges"`

	// 表权限数组
	TablePrivileges []*TablePrivileges `json:"TablePrivileges,omitempty" name:"TablePrivileges"`
}

type ModifyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitempty" name:"Account"`

	// 全局权限数组
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitempty" name:"GlobalPrivileges"`

	// 数据库权限数组
	DatabasePrivileges []*DatabasePrivileges `json:"DatabasePrivileges,omitempty" name:"DatabasePrivileges"`

	// 表权限数组
	TablePrivileges []*TablePrivileges `json:"TablePrivileges,omitempty" name:"TablePrivileges"`
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
	delete(f, "ClusterId")
	delete(f, "Account")
	delete(f, "GlobalPrivileges")
	delete(f, "DatabasePrivileges")
	delete(f, "TablePrivileges")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPrivilegesResponseParams struct {
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
type ModifyAuditRuleTemplatesRequestParams struct {
	// 审计规则模版ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`

	// 修改后的审计规则。
	RuleFilters []*RuleFilters `json:"RuleFilters,omitempty" name:"RuleFilters"`

	// 修改后的规则模版名称。
	RuleTemplateName *string `json:"RuleTemplateName,omitempty" name:"RuleTemplateName"`

	// 修改后的规则模版描述。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ModifyAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 审计规则模版ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`

	// 修改后的审计规则。
	RuleFilters []*RuleFilters `json:"RuleFilters,omitempty" name:"RuleFilters"`

	// 修改后的规则模版名称。
	RuleTemplateName *string `json:"RuleTemplateName,omitempty" name:"RuleTemplateName"`

	// 修改后的规则模版描述。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAuditRuleTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditRuleTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleTemplateIds")
	delete(f, "RuleFilters")
	delete(f, "RuleTemplateName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuditRuleTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditRuleTemplatesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAuditRuleTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAuditRuleTemplatesResponseParams `json:"Response"`
}

func (r *ModifyAuditRuleTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditRuleTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditServiceRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 日志保留时长。
	LogExpireDay *uint64 `json:"LogExpireDay,omitempty" name:"LogExpireDay"`

	// 高频日志保留时长。
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitempty" name:"HighLogExpireDay"`

	// 修改实例审计规则为全审计。
	AuditAll *bool `json:"AuditAll,omitempty" name:"AuditAll"`

	// 规则审计。
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitempty" name:"AuditRuleFilters"`

	// 规则模版ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`
}

type ModifyAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 日志保留时长。
	LogExpireDay *uint64 `json:"LogExpireDay,omitempty" name:"LogExpireDay"`

	// 高频日志保留时长。
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitempty" name:"HighLogExpireDay"`

	// 修改实例审计规则为全审计。
	AuditAll *bool `json:"AuditAll,omitempty" name:"AuditAll"`

	// 规则审计。
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitempty" name:"AuditRuleFilters"`

	// 规则模版ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`
}

func (r *ModifyAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	delete(f, "HighLogExpireDay")
	delete(f, "AuditAll")
	delete(f, "AuditRuleFilters")
	delete(f, "RuleTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditServiceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAuditServiceResponseParams `json:"Response"`
}

func (r *ModifyAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupConfigRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 表示保留备份时长, 单位秒，超过该时间将被清理, 七天表示为3600*24*7=604800，最大为158112000
	ReserveDuration *uint64 `json:"ReserveDuration,omitempty" name:"ReserveDuration"`

	// 表示全备开始时间，[0-24*3600]， 如0:00, 1:00, 2:00 分别为 0，3600， 7200
	BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitempty" name:"BackupTimeBeg"`

	// 表示全备结束时间，[0-24*3600]， 如0:00, 1:00, 2:00 分别为 0，3600， 7200
	BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitempty" name:"BackupTimeEnd"`

	// 该参数目前不支持修改，无需填写。备份频率，长度为7的数组，分别对应周一到周日的备份方式，full-全量备份，increment-增量备份
	BackupFreq []*string `json:"BackupFreq,omitempty" name:"BackupFreq"`

	// 该参数目前不支持修改，无需填写。备份方式，logic-逻辑备份，snapshot-快照备份
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`
}

type ModifyBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 表示保留备份时长, 单位秒，超过该时间将被清理, 七天表示为3600*24*7=604800，最大为158112000
	ReserveDuration *uint64 `json:"ReserveDuration,omitempty" name:"ReserveDuration"`

	// 表示全备开始时间，[0-24*3600]， 如0:00, 1:00, 2:00 分别为 0，3600， 7200
	BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitempty" name:"BackupTimeBeg"`

	// 表示全备结束时间，[0-24*3600]， 如0:00, 1:00, 2:00 分别为 0，3600， 7200
	BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitempty" name:"BackupTimeEnd"`

	// 该参数目前不支持修改，无需填写。备份频率，长度为7的数组，分别对应周一到周日的备份方式，full-全量备份，increment-增量备份
	BackupFreq []*string `json:"BackupFreq,omitempty" name:"BackupFreq"`

	// 该参数目前不支持修改，无需填写。备份方式，logic-逻辑备份，snapshot-快照备份
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`
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
	delete(f, "ClusterId")
	delete(f, "ReserveDuration")
	delete(f, "BackupTimeBeg")
	delete(f, "BackupTimeEnd")
	delete(f, "BackupFreq")
	delete(f, "BackupType")
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
type ModifyBackupNameRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 备份文件ID
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`

	// 备注名，长度不能超过60个字符
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
}

type ModifyBackupNameRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 备份文件ID
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`

	// 备注名，长度不能超过60个字符
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
}

func (r *ModifyBackupNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "BackupId")
	delete(f, "BackupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupNameResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBackupNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupNameResponseParams `json:"Response"`
}

func (r *ModifyBackupNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBinlogSaveDaysRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Binlog保留天数
	BinlogSaveDays *int64 `json:"BinlogSaveDays,omitempty" name:"BinlogSaveDays"`
}

type ModifyBinlogSaveDaysRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Binlog保留天数
	BinlogSaveDays *int64 `json:"BinlogSaveDays,omitempty" name:"BinlogSaveDays"`
}

func (r *ModifyBinlogSaveDaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBinlogSaveDaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "BinlogSaveDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBinlogSaveDaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBinlogSaveDaysResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBinlogSaveDaysResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBinlogSaveDaysResponseParams `json:"Response"`
}

func (r *ModifyBinlogSaveDaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBinlogSaveDaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterDatabaseRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库名
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 新授权用户主机权限
	NewUserHostPrivileges []*UserHostPrivilege `json:"NewUserHostPrivileges,omitempty" name:"NewUserHostPrivileges"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`

	// 历史授权用户主机权限
	OldUserHostPrivileges []*UserHostPrivilege `json:"OldUserHostPrivileges,omitempty" name:"OldUserHostPrivileges"`
}

type ModifyClusterDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库名
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 新授权用户主机权限
	NewUserHostPrivileges []*UserHostPrivilege `json:"NewUserHostPrivileges,omitempty" name:"NewUserHostPrivileges"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`

	// 历史授权用户主机权限
	OldUserHostPrivileges []*UserHostPrivilege `json:"OldUserHostPrivileges,omitempty" name:"OldUserHostPrivileges"`
}

func (r *ModifyClusterDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DbName")
	delete(f, "NewUserHostPrivileges")
	delete(f, "Description")
	delete(f, "OldUserHostPrivileges")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterDatabaseResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterDatabaseResponseParams `json:"Response"`
}

func (r *ModifyClusterDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterNameRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type ModifyClusterNameRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

func (r *ModifyClusterNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterNameResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterNameResponseParams `json:"Response"`
}

func (r *ModifyClusterNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterParamRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 要修改的参数列表。每一个元素是ParamName、CurrentValue和OldValue的组合。ParamName是参数名称，CurrentValue是当前值，OldValue是之前值且不做校验
	ParamList []*ParamItem `json:"ParamList,omitempty" name:"ParamList"`

	// 维护期间执行-yes,立即执行-no
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitempty" name:"IsInMaintainPeriod"`
}

type ModifyClusterParamRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 要修改的参数列表。每一个元素是ParamName、CurrentValue和OldValue的组合。ParamName是参数名称，CurrentValue是当前值，OldValue是之前值且不做校验
	ParamList []*ParamItem `json:"ParamList,omitempty" name:"ParamList"`

	// 维护期间执行-yes,立即执行-no
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitempty" name:"IsInMaintainPeriod"`
}

func (r *ModifyClusterParamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterParamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ParamList")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterParamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterParamResponseParams struct {
	// 异步请求Id，用于查询结果
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterParamResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterParamResponseParams `json:"Response"`
}

func (r *ModifyClusterParamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterParamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterPasswordComplexityRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 密码长度
	ValidatePasswordLength *int64 `json:"ValidatePasswordLength,omitempty" name:"ValidatePasswordLength"`

	// 大小写字符个数
	ValidatePasswordMixedCaseCount *int64 `json:"ValidatePasswordMixedCaseCount,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// 特殊字符个数
	ValidatePasswordSpecialCharCount *int64 `json:"ValidatePasswordSpecialCharCount,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// 数字个数
	ValidatePasswordNumberCount *int64 `json:"ValidatePasswordNumberCount,omitempty" name:"ValidatePasswordNumberCount"`

	// 密码强度（"MEDIUM", "STRONG"）
	ValidatePasswordPolicy *string `json:"ValidatePasswordPolicy,omitempty" name:"ValidatePasswordPolicy"`

	// 数据字典
	ValidatePasswordDictionary []*string `json:"ValidatePasswordDictionary,omitempty" name:"ValidatePasswordDictionary"`
}

type ModifyClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 密码长度
	ValidatePasswordLength *int64 `json:"ValidatePasswordLength,omitempty" name:"ValidatePasswordLength"`

	// 大小写字符个数
	ValidatePasswordMixedCaseCount *int64 `json:"ValidatePasswordMixedCaseCount,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// 特殊字符个数
	ValidatePasswordSpecialCharCount *int64 `json:"ValidatePasswordSpecialCharCount,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// 数字个数
	ValidatePasswordNumberCount *int64 `json:"ValidatePasswordNumberCount,omitempty" name:"ValidatePasswordNumberCount"`

	// 密码强度（"MEDIUM", "STRONG"）
	ValidatePasswordPolicy *string `json:"ValidatePasswordPolicy,omitempty" name:"ValidatePasswordPolicy"`

	// 数据字典
	ValidatePasswordDictionary []*string `json:"ValidatePasswordDictionary,omitempty" name:"ValidatePasswordDictionary"`
}

func (r *ModifyClusterPasswordComplexityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterPasswordComplexityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ValidatePasswordLength")
	delete(f, "ValidatePasswordMixedCaseCount")
	delete(f, "ValidatePasswordSpecialCharCount")
	delete(f, "ValidatePasswordNumberCount")
	delete(f, "ValidatePasswordPolicy")
	delete(f, "ValidatePasswordDictionary")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterPasswordComplexityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterPasswordComplexityResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterPasswordComplexityResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterPasswordComplexityResponseParams `json:"Response"`
}

func (r *ModifyClusterPasswordComplexityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterPasswordComplexityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterSlaveZoneRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 旧从可用区
	OldSlaveZone *string `json:"OldSlaveZone,omitempty" name:"OldSlaveZone"`

	// 新从可用区
	NewSlaveZone *string `json:"NewSlaveZone,omitempty" name:"NewSlaveZone"`
}

type ModifyClusterSlaveZoneRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 旧从可用区
	OldSlaveZone *string `json:"OldSlaveZone,omitempty" name:"OldSlaveZone"`

	// 新从可用区
	NewSlaveZone *string `json:"NewSlaveZone,omitempty" name:"NewSlaveZone"`
}

func (r *ModifyClusterSlaveZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterSlaveZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "OldSlaveZone")
	delete(f, "NewSlaveZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterSlaveZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterSlaveZoneResponseParams struct {
	// 异步FlowId
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterSlaveZoneResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterSlaveZoneResponseParams `json:"Response"`
}

func (r *ModifyClusterSlaveZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterSlaveZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterStorageRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群新存储大小（单位G）
	NewStorageLimit *int64 `json:"NewStorageLimit,omitempty" name:"NewStorageLimit"`

	// 集群原存储大小（单位G）
	OldStorageLimit *int64 `json:"OldStorageLimit,omitempty" name:"OldStorageLimit"`

	// 交易模式 0-下单并支付 1-下单
	DealMode *int64 `json:"DealMode,omitempty" name:"DealMode"`
}

type ModifyClusterStorageRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群新存储大小（单位G）
	NewStorageLimit *int64 `json:"NewStorageLimit,omitempty" name:"NewStorageLimit"`

	// 集群原存储大小（单位G）
	OldStorageLimit *int64 `json:"OldStorageLimit,omitempty" name:"OldStorageLimit"`

	// 交易模式 0-下单并支付 1-下单
	DealMode *int64 `json:"DealMode,omitempty" name:"DealMode"`
}

func (r *ModifyClusterStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterStorageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NewStorageLimit")
	delete(f, "OldStorageLimit")
	delete(f, "DealMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterStorageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterStorageResponseParams struct {
	// 冻结流水ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranId *string `json:"TranId,omitempty" name:"TranId"`

	// 大订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`

	// 订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterStorageResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterStorageResponseParams `json:"Response"`
}

func (r *ModifyClusterStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// 实例组ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 要修改的安全组ID列表，一个或者多个安全组ID组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例组ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 要修改的安全组ID列表，一个或者多个安全组ID组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`
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
	delete(f, "Zone")
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
type ModifyInstanceNameRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type ModifyInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyInstanceParamRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 集群参数列表
	ClusterParamList []*ModifyParamItem `json:"ClusterParamList,omitempty" name:"ClusterParamList"`

	// 实例参数列表
	InstanceParamList []*ModifyParamItem `json:"InstanceParamList,omitempty" name:"InstanceParamList"`

	// yes：在运维时间窗内修改，no：立即执行（默认值）
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitempty" name:"IsInMaintainPeriod"`
}

type ModifyInstanceParamRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 集群参数列表
	ClusterParamList []*ModifyParamItem `json:"ClusterParamList,omitempty" name:"ClusterParamList"`

	// 实例参数列表
	InstanceParamList []*ModifyParamItem `json:"InstanceParamList,omitempty" name:"InstanceParamList"`

	// yes：在运维时间窗内修改，no：立即执行（默认值）
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitempty" name:"IsInMaintainPeriod"`
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
	delete(f, "ClusterId")
	delete(f, "InstanceIds")
	delete(f, "ClusterParamList")
	delete(f, "InstanceParamList")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceParamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamResponseParams struct {
	// 任务ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

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
type ModifyMaintainPeriodConfigRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 维护开始时间，单位为秒，如3:00为10800
	MaintainStartTime *int64 `json:"MaintainStartTime,omitempty" name:"MaintainStartTime"`

	// 维护持续时间，单位为秒，如1小时为3600
	MaintainDuration *int64 `json:"MaintainDuration,omitempty" name:"MaintainDuration"`

	// 每周维护日期，日期取值范围[Mon, Tue, Wed, Thu, Fri, Sat, Sun]
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitempty" name:"MaintainWeekDays"`
}

type ModifyMaintainPeriodConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 维护开始时间，单位为秒，如3:00为10800
	MaintainStartTime *int64 `json:"MaintainStartTime,omitempty" name:"MaintainStartTime"`

	// 维护持续时间，单位为秒，如1小时为3600
	MaintainDuration *int64 `json:"MaintainDuration,omitempty" name:"MaintainDuration"`

	// 每周维护日期，日期取值范围[Mon, Tue, Wed, Thu, Fri, Sat, Sun]
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitempty" name:"MaintainWeekDays"`
}

func (r *ModifyMaintainPeriodConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaintainPeriodConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MaintainStartTime")
	delete(f, "MaintainDuration")
	delete(f, "MaintainWeekDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMaintainPeriodConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaintainPeriodConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyMaintainPeriodConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMaintainPeriodConfigResponseParams `json:"Response"`
}

func (r *ModifyMaintainPeriodConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaintainPeriodConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyParamItem struct {
	// 参数名
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 参数当前值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 参数旧值（只在出参时有用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldValue *string `json:"OldValue,omitempty" name:"OldValue"`
}

// Predefined struct for user
type ModifyParamTemplateRequestParams struct {
	// 模版ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模版名
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 模版描述
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`

	// 参数列表
	ParamList []*ModifyParamItem `json:"ParamList,omitempty" name:"ParamList"`
}

type ModifyParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模版ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模版名
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 模版描述
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`

	// 参数列表
	ParamList []*ModifyParamItem `json:"ParamList,omitempty" name:"ParamList"`
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
	delete(f, "TemplateName")
	delete(f, "TemplateDescription")
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
type ModifyProxyDescRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 数据库代理描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ModifyProxyDescRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 数据库代理描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyProxyDescRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyDescRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProxyGroupId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProxyDescRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProxyDescResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyProxyDescResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProxyDescResponseParams `json:"Response"`
}

func (r *ModifyProxyDescResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyDescResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProxyRwSplitRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 一致性类型；“eventual"-最终一致性, "session"-会话一致性, "global"-全局一致性
	ConsistencyType *string `json:"ConsistencyType,omitempty" name:"ConsistencyType"`

	// 一致性超时时间
	ConsistencyTimeOut *string `json:"ConsistencyTimeOut,omitempty" name:"ConsistencyTimeOut"`

	// 读写权重分配模式；系统自动分配："system"， 自定义："custom"
	WeightMode *string `json:"WeightMode,omitempty" name:"WeightMode"`

	// 实例只读权重
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitempty" name:"InstanceWeights"`

	// 是否开启故障转移，代理出现故障后，连接地址将路由到主实例，取值："yes" , "no"
	FailOver *string `json:"FailOver,omitempty" name:"FailOver"`

	// 是否自动添加只读实例，取值："yes" , "no"
	AutoAddRo *string `json:"AutoAddRo,omitempty" name:"AutoAddRo"`

	// 是否打开读写分离
	OpenRw *string `json:"OpenRw,omitempty" name:"OpenRw"`

	// 读写类型：
	// READWRITE,READONLY
	RwType *string `json:"RwType,omitempty" name:"RwType"`

	// 事务拆分
	TransSplit *bool `json:"TransSplit,omitempty" name:"TransSplit"`

	// 连接模式：
	// nearby,balance
	AccessMode *string `json:"AccessMode,omitempty" name:"AccessMode"`

	// 是否打开连接池：
	// yes,no
	OpenConnectionPool *string `json:"OpenConnectionPool,omitempty" name:"OpenConnectionPool"`

	// 连接池类型：
	// SessionConnectionPool
	ConnectionPoolType *string `json:"ConnectionPoolType,omitempty" name:"ConnectionPoolType"`

	// 连接池时间
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitempty" name:"ConnectionPoolTimeOut"`
}

type ModifyProxyRwSplitRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 一致性类型；“eventual"-最终一致性, "session"-会话一致性, "global"-全局一致性
	ConsistencyType *string `json:"ConsistencyType,omitempty" name:"ConsistencyType"`

	// 一致性超时时间
	ConsistencyTimeOut *string `json:"ConsistencyTimeOut,omitempty" name:"ConsistencyTimeOut"`

	// 读写权重分配模式；系统自动分配："system"， 自定义："custom"
	WeightMode *string `json:"WeightMode,omitempty" name:"WeightMode"`

	// 实例只读权重
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitempty" name:"InstanceWeights"`

	// 是否开启故障转移，代理出现故障后，连接地址将路由到主实例，取值："yes" , "no"
	FailOver *string `json:"FailOver,omitempty" name:"FailOver"`

	// 是否自动添加只读实例，取值："yes" , "no"
	AutoAddRo *string `json:"AutoAddRo,omitempty" name:"AutoAddRo"`

	// 是否打开读写分离
	OpenRw *string `json:"OpenRw,omitempty" name:"OpenRw"`

	// 读写类型：
	// READWRITE,READONLY
	RwType *string `json:"RwType,omitempty" name:"RwType"`

	// 事务拆分
	TransSplit *bool `json:"TransSplit,omitempty" name:"TransSplit"`

	// 连接模式：
	// nearby,balance
	AccessMode *string `json:"AccessMode,omitempty" name:"AccessMode"`

	// 是否打开连接池：
	// yes,no
	OpenConnectionPool *string `json:"OpenConnectionPool,omitempty" name:"OpenConnectionPool"`

	// 连接池类型：
	// SessionConnectionPool
	ConnectionPoolType *string `json:"ConnectionPoolType,omitempty" name:"ConnectionPoolType"`

	// 连接池时间
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitempty" name:"ConnectionPoolTimeOut"`
}

func (r *ModifyProxyRwSplitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyRwSplitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProxyGroupId")
	delete(f, "ConsistencyType")
	delete(f, "ConsistencyTimeOut")
	delete(f, "WeightMode")
	delete(f, "InstanceWeights")
	delete(f, "FailOver")
	delete(f, "AutoAddRo")
	delete(f, "OpenRw")
	delete(f, "RwType")
	delete(f, "TransSplit")
	delete(f, "AccessMode")
	delete(f, "OpenConnectionPool")
	delete(f, "ConnectionPoolType")
	delete(f, "ConnectionPoolTimeOut")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProxyRwSplitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProxyRwSplitResponseParams struct {
	// 异步FlowId
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 异步任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyProxyRwSplitResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProxyRwSplitResponseParams `json:"Response"`
}

func (r *ModifyProxyRwSplitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyRwSplitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePackageClustersRequestParams struct {
	// 资源包唯一ID
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 需要建立绑定关系的集群ID
	BindClusterIds []*string `json:"BindClusterIds,omitempty" name:"BindClusterIds"`

	// 需要解除绑定关系的集群ID
	UnbindClusterIds []*string `json:"UnbindClusterIds,omitempty" name:"UnbindClusterIds"`
}

type ModifyResourcePackageClustersRequest struct {
	*tchttp.BaseRequest
	
	// 资源包唯一ID
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 需要建立绑定关系的集群ID
	BindClusterIds []*string `json:"BindClusterIds,omitempty" name:"BindClusterIds"`

	// 需要解除绑定关系的集群ID
	UnbindClusterIds []*string `json:"UnbindClusterIds,omitempty" name:"UnbindClusterIds"`
}

func (r *ModifyResourcePackageClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePackageClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageId")
	delete(f, "BindClusterIds")
	delete(f, "UnbindClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourcePackageClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePackageClustersResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyResourcePackageClustersResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourcePackageClustersResponseParams `json:"Response"`
}

func (r *ModifyResourcePackageClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePackageClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePackageNameRequestParams struct {
	// 资源包唯一ID
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 自定义的资源包名称，最长支持120个字符
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`
}

type ModifyResourcePackageNameRequest struct {
	*tchttp.BaseRequest
	
	// 资源包唯一ID
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 自定义的资源包名称，最长支持120个字符
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`
}

func (r *ModifyResourcePackageNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePackageNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageId")
	delete(f, "PackageName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourcePackageNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePackageNameResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyResourcePackageNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourcePackageNameResponseParams `json:"Response"`
}

func (r *ModifyResourcePackageNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePackageNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVipVportRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例组id
	InstanceGrpId *string `json:"InstanceGrpId,omitempty" name:"InstanceGrpId"`

	// 需要修改的目的ip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 需要修改的目的端口
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 旧ip回收前的保留时间，单位小时，0表示立即回收
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitempty" name:"OldIpReserveHours"`
}

type ModifyVipVportRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例组id
	InstanceGrpId *string `json:"InstanceGrpId,omitempty" name:"InstanceGrpId"`

	// 需要修改的目的ip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 需要修改的目的端口
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 旧ip回收前的保留时间，单位小时，0表示立即回收
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitempty" name:"OldIpReserveHours"`
}

func (r *ModifyVipVportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVipVportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceGrpId")
	delete(f, "Vip")
	delete(f, "Vport")
	delete(f, "DbType")
	delete(f, "OldIpReserveHours")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVipVportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVipVportResponseParams struct {
	// 异步任务id
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVipVportResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVipVportResponseParams `json:"Response"`
}

func (r *ModifyVipVportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVipVportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Module struct {
	// 是否支持，可选值:yes,no
	IsDisable *string `json:"IsDisable,omitempty" name:"IsDisable"`

	// 模块名
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

type NetAddr struct {
	// 内网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 内网端口号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 外网域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`

	// 外网端口号
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`

	// 网络类型（ro-只读,rw/ha-读写）
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 外网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanIP *string `json:"WanIP,omitempty" name:"WanIP"`

	// 外网状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanStatus *string `json:"WanStatus,omitempty" name:"WanStatus"`

	// 实例组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
}

type NewAccount struct {
	// 账户名，包含字母数字_,以字母开头，字母或数字结尾，长度1-16
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 密码，密码长度范围为8到64个字符
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`

	// 主机
	Host *string `json:"Host,omitempty" name:"Host"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 用户最大连接数，不能大于10240
	MaxUserConnections *int64 `json:"MaxUserConnections,omitempty" name:"MaxUserConnections"`
}

type ObjectTask struct {
	// 任务自增ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 任务ID（集群ID|实例组ID|实例ID）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectId *string `json:"ObjectId,omitempty" name:"ObjectId"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectType *string `json:"ObjectType,omitempty" name:"ObjectType"`
}

// Predefined struct for user
type OfflineClusterRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type OfflineClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *OfflineClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OfflineClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineClusterResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OfflineClusterResponse struct {
	*tchttp.BaseResponse
	Response *OfflineClusterResponseParams `json:"Response"`
}

func (r *OfflineClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineInstanceRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID数组
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
}

type OfflineInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID数组
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
}

func (r *OfflineInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OfflineInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineInstanceResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OfflineInstanceResponse struct {
	*tchttp.BaseResponse
	Response *OfflineInstanceResponseParams `json:"Response"`
}

func (r *OfflineInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OldAddrInfo struct {
	// IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 期望执行回收时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnTime *string `json:"ReturnTime,omitempty" name:"ReturnTime"`
}

// Predefined struct for user
type OpenAuditServiceRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 日志保留时长。
	LogExpireDay *uint64 `json:"LogExpireDay,omitempty" name:"LogExpireDay"`

	// 高频日志保留时长。
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitempty" name:"HighLogExpireDay"`

	// 审计规则。同RuleTemplateIds都不填是全审计。
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitempty" name:"AuditRuleFilters"`

	// 规则模版ID。同AuditRuleFilters都不填是全审计。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`
}

type OpenAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 日志保留时长。
	LogExpireDay *uint64 `json:"LogExpireDay,omitempty" name:"LogExpireDay"`

	// 高频日志保留时长。
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitempty" name:"HighLogExpireDay"`

	// 审计规则。同RuleTemplateIds都不填是全审计。
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitempty" name:"AuditRuleFilters"`

	// 规则模版ID。同AuditRuleFilters都不填是全审计。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`
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
	delete(f, "AuditRuleFilters")
	delete(f, "RuleTemplateIds")
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
type OpenClusterPasswordComplexityRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 密码长度
	ValidatePasswordLength *int64 `json:"ValidatePasswordLength,omitempty" name:"ValidatePasswordLength"`

	// 大小写字符个数
	ValidatePasswordMixedCaseCount *int64 `json:"ValidatePasswordMixedCaseCount,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// 特殊字符个数
	ValidatePasswordSpecialCharCount *int64 `json:"ValidatePasswordSpecialCharCount,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// 数字个数
	ValidatePasswordNumberCount *int64 `json:"ValidatePasswordNumberCount,omitempty" name:"ValidatePasswordNumberCount"`

	// 密码强度（"MEDIUM", "STRONG"）
	ValidatePasswordPolicy *string `json:"ValidatePasswordPolicy,omitempty" name:"ValidatePasswordPolicy"`

	// 数据字典
	ValidatePasswordDictionary []*string `json:"ValidatePasswordDictionary,omitempty" name:"ValidatePasswordDictionary"`
}

type OpenClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 密码长度
	ValidatePasswordLength *int64 `json:"ValidatePasswordLength,omitempty" name:"ValidatePasswordLength"`

	// 大小写字符个数
	ValidatePasswordMixedCaseCount *int64 `json:"ValidatePasswordMixedCaseCount,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// 特殊字符个数
	ValidatePasswordSpecialCharCount *int64 `json:"ValidatePasswordSpecialCharCount,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// 数字个数
	ValidatePasswordNumberCount *int64 `json:"ValidatePasswordNumberCount,omitempty" name:"ValidatePasswordNumberCount"`

	// 密码强度（"MEDIUM", "STRONG"）
	ValidatePasswordPolicy *string `json:"ValidatePasswordPolicy,omitempty" name:"ValidatePasswordPolicy"`

	// 数据字典
	ValidatePasswordDictionary []*string `json:"ValidatePasswordDictionary,omitempty" name:"ValidatePasswordDictionary"`
}

func (r *OpenClusterPasswordComplexityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenClusterPasswordComplexityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ValidatePasswordLength")
	delete(f, "ValidatePasswordMixedCaseCount")
	delete(f, "ValidatePasswordSpecialCharCount")
	delete(f, "ValidatePasswordNumberCount")
	delete(f, "ValidatePasswordPolicy")
	delete(f, "ValidatePasswordDictionary")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenClusterPasswordComplexityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenClusterPasswordComplexityResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OpenClusterPasswordComplexityResponse struct {
	*tchttp.BaseResponse
	Response *OpenClusterPasswordComplexityResponseParams `json:"Response"`
}

func (r *OpenClusterPasswordComplexityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenClusterPasswordComplexityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenReadOnlyInstanceExclusiveAccessRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 需要开通独有访问的只读实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 指定的vpc ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 指定的子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 安全组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

type OpenReadOnlyInstanceExclusiveAccessRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 需要开通独有访问的只读实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 指定的vpc ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 指定的子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 安全组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

func (r *OpenReadOnlyInstanceExclusiveAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenReadOnlyInstanceExclusiveAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Port")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenReadOnlyInstanceExclusiveAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenReadOnlyInstanceExclusiveAccessResponseParams struct {
	// 开通流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OpenReadOnlyInstanceExclusiveAccessResponse struct {
	*tchttp.BaseResponse
	Response *OpenReadOnlyInstanceExclusiveAccessResponseParams `json:"Response"`
}

func (r *OpenReadOnlyInstanceExclusiveAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenReadOnlyInstanceExclusiveAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenWanRequestParams struct {
	// 实例组id
	InstanceGrpId *string `json:"InstanceGrpId,omitempty" name:"InstanceGrpId"`
}

type OpenWanRequest struct {
	*tchttp.BaseRequest
	
	// 实例组id
	InstanceGrpId *string `json:"InstanceGrpId,omitempty" name:"InstanceGrpId"`
}

func (r *OpenWanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenWanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceGrpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenWanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenWanResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OpenWanResponse struct {
	*tchttp.BaseResponse
	Response *OpenWanResponseParams `json:"Response"`
}

func (r *OpenWanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenWanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Package struct {
	// AppID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 资源包唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 资源包名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 资源包类型
	// CCU-计算资源包，DISK-存储资源包
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// 资源包使用地域
	// china-中国内地通用，overseas-港澳台及海外通用
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageRegion *string `json:"PackageRegion,omitempty" name:"PackageRegion"`

	// 资源包状态
	// creating-创建中；
	// using-使用中；
	// expired-已过期；
	// normal_finish-使用完；
	// apply_refund-申请退费中；
	// refund-已退费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 资源包总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageTotalSpec *float64 `json:"PackageTotalSpec,omitempty" name:"PackageTotalSpec"`

	// 资源包已使用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageUsedSpec *float64 `json:"PackageUsedSpec,omitempty" name:"PackageUsedSpec"`

	// 资源包已使用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasQuota *bool `json:"HasQuota,omitempty" name:"HasQuota"`

	// 绑定实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindInstanceInfos []*BindInstanceInfo `json:"BindInstanceInfos,omitempty" name:"BindInstanceInfos"`

	// 生效时间：2022-07-01 00:00:00
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 失效时间：2022-08-01 00:00:00
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type PackageDetail struct {
	// AppId账户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 资源包唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 成功抵扣容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessDeductSpec *float64 `json:"SuccessDeductSpec,omitempty" name:"SuccessDeductSpec"`

	// 截止当前，资源包已使用的容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageTotalUsedSpec *float64 `json:"PackageTotalUsedSpec,omitempty" name:"PackageTotalUsedSpec"`

	// 抵扣开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 抵扣结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 扩展信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtendInfo *string `json:"ExtendInfo,omitempty" name:"ExtendInfo"`
}

type ParamDetail struct {
	// 参数名称
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 参数类型：integer，enum，float，string，func
	ParamType *string `json:"ParamType,omitempty" name:"ParamType"`

	// true-支持"func"，false-不支持公式
	SupportFunc *bool `json:"SupportFunc,omitempty" name:"SupportFunc"`

	// 默认值
	Default *string `json:"Default,omitempty" name:"Default"`

	// 参数描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 参数当前值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 修改参数后，是否需要重启数据库以使参数生效。0-不需要重启，1-需要重启。
	NeedReboot *int64 `json:"NeedReboot,omitempty" name:"NeedReboot"`

	// 参数容许的最大值
	Max *string `json:"Max,omitempty" name:"Max"`

	// 参数容许的最小值
	Min *string `json:"Min,omitempty" name:"Min"`

	// 参数的可选枚举值。如果为非枚举值，则为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`

	// 1：全局参数，0：非全局参数
	IsGlobal *int64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 匹配类型，multiVal
	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`

	// 匹配目标值，当multiVal时，各个key用，分割
	MatchValue *string `json:"MatchValue,omitempty" name:"MatchValue"`

	// true-为公式，false-非公式
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsFunc *bool `json:"IsFunc,omitempty" name:"IsFunc"`

	// 参数设置为公式时，Func返回设置的公式内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Func *string `json:"Func,omitempty" name:"Func"`

	// 参数是否可修改
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiableInfo *ModifiableInfo `json:"ModifiableInfo,omitempty" name:"ModifiableInfo"`
}

type ParamInfo struct {
	// 当前值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 默认值
	Default *string `json:"Default,omitempty" name:"Default"`

	// 参数为enum/string/bool时，可选值列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`

	// 参数类型为float/integer时的最大值
	Max *string `json:"Max,omitempty" name:"Max"`

	// 参数类型为float/integer时的最小值
	Min *string `json:"Min,omitempty" name:"Min"`

	// 参数名称
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 是否需要重启生效
	NeedReboot *int64 `json:"NeedReboot,omitempty" name:"NeedReboot"`

	// 参数类型：integer/float/string/enum/bool
	ParamType *string `json:"ParamType,omitempty" name:"ParamType"`

	// 匹配类型，multiVal, regex在参数类型是string时使用
	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`

	// 匹配目标值，当multiVal时，各个key用;分割
	MatchValue *string `json:"MatchValue,omitempty" name:"MatchValue"`

	// 参数描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 是否为全局参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsGlobal *int64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 参数是否可修改
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiableInfo *ModifiableInfo `json:"ModifiableInfo,omitempty" name:"ModifiableInfo"`

	// 是否为函数
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsFunc *bool `json:"IsFunc,omitempty" name:"IsFunc"`

	// 函数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Func *string `json:"Func,omitempty" name:"Func"`
}

type ParamItem struct {
	// 参数名称
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 当前值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 原有值
	OldValue *string `json:"OldValue,omitempty" name:"OldValue"`
}

type ParamItemDetail struct {
	// 当前值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 默认值
	Default *string `json:"Default,omitempty" name:"Default"`

	// 参数的可选枚举值。如果为非枚举值，则为空
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`

	// 1：全局参数，0：非全局参数
	IsGlobal *int64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 最大值
	Max *string `json:"Max,omitempty" name:"Max"`

	// 最小值
	Min *string `json:"Min,omitempty" name:"Min"`

	// 修改参数后，是否需要重启数据库以使参数生效。0-不需要重启，1-需要重启。
	NeedReboot *int64 `json:"NeedReboot,omitempty" name:"NeedReboot"`

	// 参数名称
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 参数类型：integer，enum，float，string，func
	ParamType *string `json:"ParamType,omitempty" name:"ParamType"`

	// 参数描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 类型是否为公式
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsFunc *bool `json:"IsFunc,omitempty" name:"IsFunc"`

	// 参数配置公式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Func *string `json:"Func,omitempty" name:"Func"`
}

type ParamTemplateListInfo struct {
	// 参数模板ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 参数模板名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 参数模板描述
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`

	// 引擎版本
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 数据库类型，可选值：NORMAL，SERVERLESS
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// 参数模板详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamInfoSet []*TemplateParamInfo `json:"ParamInfoSet,omitempty" name:"ParamInfoSet"`
}

// Predefined struct for user
type PauseServerlessRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 是否强制暂停，忽略当前的用户链接  0:不强制  1:强制， 默认为1
	ForcePause *int64 `json:"ForcePause,omitempty" name:"ForcePause"`
}

type PauseServerlessRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 是否强制暂停，忽略当前的用户链接  0:不强制  1:强制， 默认为1
	ForcePause *int64 `json:"ForcePause,omitempty" name:"ForcePause"`
}

func (r *PauseServerlessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseServerlessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ForcePause")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PauseServerlessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseServerlessResponseParams struct {
	// 异步流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PauseServerlessResponse struct {
	*tchttp.BaseResponse
	Response *PauseServerlessResponseParams `json:"Response"`
}

func (r *PauseServerlessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseServerlessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PolicyRule struct {
	// 策略，ACCEPT或者DROP
	Action *string `json:"Action,omitempty" name:"Action"`

	// 来源Ip或Ip段，例如192.168.0.0/16
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// 端口
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// 网络协议，支持udp、tcp等
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// 协议端口ID或者协议端口组ID。
	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`

	// IP地址ID或者ID地址组ID。
	AddressModule *string `json:"AddressModule,omitempty" name:"AddressModule"`

	// id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type ProxyConnectionPoolInfo struct {
	// 连接池保持阈值：单位（秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitempty" name:"ConnectionPoolTimeOut"`

	// 是否开启了连接池
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenConnectionPool *string `json:"OpenConnectionPool,omitempty" name:"OpenConnectionPool"`

	// 连接池类型：SessionConnectionPool（会话级别连接池
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectionPoolType *string `json:"ConnectionPoolType,omitempty" name:"ConnectionPoolType"`
}

type ProxyGroup struct {
	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 数据库代理组节点个数
	ProxyNodeCount *int64 `json:"ProxyNodeCount,omitempty" name:"ProxyNodeCount"`

	// 数据库代理组状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 当前代理版本
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitempty" name:"CurrentProxyVersion"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 用户AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 读写节点开通数据库代理
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenRw *string `json:"OpenRw,omitempty" name:"OpenRw"`
}

type ProxyGroupInfo struct {
	// 数据库代理组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyGroup *ProxyGroup `json:"ProxyGroup,omitempty" name:"ProxyGroup"`

	// 数据库代理组读写分离信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyGroupRwInfo *ProxyGroupRwInfo `json:"ProxyGroupRwInfo,omitempty" name:"ProxyGroupRwInfo"`

	// 数据库代理节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyNodes []*ProxyNodeInfo `json:"ProxyNodes,omitempty" name:"ProxyNodes"`

	// 数据库代理连接池信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectionPool *ProxyConnectionPoolInfo `json:"ConnectionPool,omitempty" name:"ConnectionPool"`

	// 数据库代理网络信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetAddrInfos []*NetAddr `json:"NetAddrInfos,omitempty" name:"NetAddrInfos"`

	// 数据库代理任务集
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*ObjectTask `json:"Tasks,omitempty" name:"Tasks"`
}

type ProxyGroupRwInfo struct {
	// 一致性类型 eventual-最终一致性,global-全局一致性,session-会话一致性
	ConsistencyType *string `json:"ConsistencyType,omitempty" name:"ConsistencyType"`

	// 一致性超时时间
	ConsistencyTimeOut *int64 `json:"ConsistencyTimeOut,omitempty" name:"ConsistencyTimeOut"`

	// 权重模式 system-系统分配，custom-自定义
	WeightMode *string `json:"WeightMode,omitempty" name:"WeightMode"`

	// 是否开启故障转移
	FailOver *string `json:"FailOver,omitempty" name:"FailOver"`

	// 是否自动添加只读实例，yes-是，no-不自动添加
	AutoAddRo *string `json:"AutoAddRo,omitempty" name:"AutoAddRo"`

	// 实例权重数组
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitempty" name:"InstanceWeights"`

	// 是否开通读写节点，yse-是，no-否
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenRw *string `json:"OpenRw,omitempty" name:"OpenRw"`

	// 读写属性，可选值：READWRITE,READONLY
	RwType *string `json:"RwType,omitempty" name:"RwType"`

	// 事务拆分
	TransSplit *bool `json:"TransSplit,omitempty" name:"TransSplit"`

	// 连接模式，可选值：balance，nearby
	AccessMode *string `json:"AccessMode,omitempty" name:"AccessMode"`
}

type ProxyInstanceWeight struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例权重
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type ProxyNodeInfo struct {
	// 数据库代理节点ID
	ProxyNodeId *string `json:"ProxyNodeId,omitempty" name:"ProxyNodeId"`

	// 节点当前连接数, DescribeProxyNodes接口此字段值不返回
	ProxyNodeConnections *int64 `json:"ProxyNodeConnections,omitempty" name:"ProxyNodeConnections"`

	// 数据库代理节点cpu
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 数据库代理节点内存
	Mem *int64 `json:"Mem,omitempty" name:"Mem"`

	// 数据库代理节点状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 用户AppID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type ProxySpec struct {
	// 数据库代理cpu核数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 数据库代理内存
	Mem *int64 `json:"Mem,omitempty" name:"Mem"`
}

type ProxyZone struct {
	// proxy节点可用区
	ProxyNodeZone *string `json:"ProxyNodeZone,omitempty" name:"ProxyNodeZone"`

	// proxy节点数量
	ProxyNodeCount *int64 `json:"ProxyNodeCount,omitempty" name:"ProxyNodeCount"`
}

type QueryFilter struct {
	// 搜索字段，目前支持："InstanceId", "ProjectId", "InstanceName", "Vip"
	Names []*string `json:"Names,omitempty" name:"Names"`

	// 搜索字符串
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否精确匹配
	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`

	// 搜索字段
	Name *string `json:"Name,omitempty" name:"Name"`

	// 操作符
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type QueryParamFilter struct {
	// 搜索字段，目前支持："InstanceId", "ProjectId", "InstanceName", "Vip"
	Names []*string `json:"Names,omitempty" name:"Names"`

	// 搜索字符串
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否精确匹配
	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

// Predefined struct for user
type RefundResourcePackageRequestParams struct {
	// 资源包唯一ID
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`
}

type RefundResourcePackageRequest struct {
	*tchttp.BaseRequest
	
	// 资源包唯一ID
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`
}

func (r *RefundResourcePackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundResourcePackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefundResourcePackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefundResourcePackageResponseParams struct {
	// 每个物品对应一个dealName，业务需要根据dealName保证发货接口幂等
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RefundResourcePackageResponse struct {
	*tchttp.BaseResponse
	Response *RefundResourcePackageResponseParams `json:"Response"`
}

func (r *RefundResourcePackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundResourcePackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReloadBalanceProxyNodeRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`
}

type ReloadBalanceProxyNodeRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`
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
	delete(f, "ClusterId")
	delete(f, "ProxyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReloadBalanceProxyNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReloadBalanceProxyNodeResponseParams struct {
	// 异步流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 异步任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

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

// Predefined struct for user
type RemoveClusterSlaveZoneRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 从可用区
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`
}

type RemoveClusterSlaveZoneRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 从可用区
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`
}

func (r *RemoveClusterSlaveZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveClusterSlaveZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SlaveZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveClusterSlaveZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveClusterSlaveZoneResponseParams struct {
	// 异步FlowId
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveClusterSlaveZoneResponse struct {
	*tchttp.BaseResponse
	Response *RemoveClusterSlaveZoneResponseParams `json:"Response"`
}

func (r *RemoveClusterSlaveZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveClusterSlaveZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAccountPasswordRequestParams struct {
	// 数据库账号名
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 数据库账号新密码
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 主机，不填默认为"%"
	Host *string `json:"Host,omitempty" name:"Host"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 数据库账号名
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 数据库账号新密码
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 主机，不填默认为"%"
	Host *string `json:"Host,omitempty" name:"Host"`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountName")
	delete(f, "AccountPassword")
	delete(f, "ClusterId")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAccountPasswordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetAccountPasswordResponseParams `json:"Response"`
}

func (r *ResetAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourcePackage struct {
	// 资源包的唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 资源包类型：CCU：计算资源包
	// DISK：存储资源包
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`
}

// Predefined struct for user
type RestartInstanceRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type RestartInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *RestartInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartInstanceResponseParams struct {
	// 异步任务id
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RestartInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RestartInstanceResponseParams `json:"Response"`
}

func (r *RestartInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeServerlessRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type ResumeServerlessRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ResumeServerlessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeServerlessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeServerlessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeServerlessResponseParams struct {
	// 异步流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResumeServerlessResponse struct {
	*tchttp.BaseResponse
	Response *ResumeServerlessResponseParams `json:"Response"`
}

func (r *ResumeServerlessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeServerlessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RevokeAccountPrivilegesRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitempty" name:"Account"`

	// 数据库表权限数组
	DbTablePrivileges []*string `json:"DbTablePrivileges,omitempty" name:"DbTablePrivileges"`

	// 数据库表信息
	DbTables []*DbTable `json:"DbTables,omitempty" name:"DbTables"`
}

type RevokeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitempty" name:"Account"`

	// 数据库表权限数组
	DbTablePrivileges []*string `json:"DbTablePrivileges,omitempty" name:"DbTablePrivileges"`

	// 数据库表信息
	DbTables []*DbTable `json:"DbTables,omitempty" name:"DbTables"`
}

func (r *RevokeAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevokeAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Account")
	delete(f, "DbTablePrivileges")
	delete(f, "DbTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RevokeAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RevokeAccountPrivilegesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RevokeAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *RevokeAccountPrivilegesResponseParams `json:"Response"`
}

func (r *RevokeAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevokeAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollBackClusterRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 回档策略 timeRollback-按时间点回档 snapRollback-按备份文件回档
	RollbackStrategy *string `json:"RollbackStrategy,omitempty" name:"RollbackStrategy"`

	// 回档ID
	RollbackId *uint64 `json:"RollbackId,omitempty" name:"RollbackId"`

	// 期望回档时间
	ExpectTime *string `json:"ExpectTime,omitempty" name:"ExpectTime"`

	// 期望阈值（已废弃）
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitempty" name:"ExpectTimeThresh"`

	// 回档数据库列表
	RollbackDatabases []*RollbackDatabase `json:"RollbackDatabases,omitempty" name:"RollbackDatabases"`

	// 回档数据库表列表
	RollbackTables []*RollbackTable `json:"RollbackTables,omitempty" name:"RollbackTables"`

	// 按时间点回档模式，full: 普通; db: 快速; table: 极速  （默认是普通）
	RollbackMode *string `json:"RollbackMode,omitempty" name:"RollbackMode"`
}

type RollBackClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 回档策略 timeRollback-按时间点回档 snapRollback-按备份文件回档
	RollbackStrategy *string `json:"RollbackStrategy,omitempty" name:"RollbackStrategy"`

	// 回档ID
	RollbackId *uint64 `json:"RollbackId,omitempty" name:"RollbackId"`

	// 期望回档时间
	ExpectTime *string `json:"ExpectTime,omitempty" name:"ExpectTime"`

	// 期望阈值（已废弃）
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitempty" name:"ExpectTimeThresh"`

	// 回档数据库列表
	RollbackDatabases []*RollbackDatabase `json:"RollbackDatabases,omitempty" name:"RollbackDatabases"`

	// 回档数据库表列表
	RollbackTables []*RollbackTable `json:"RollbackTables,omitempty" name:"RollbackTables"`

	// 按时间点回档模式，full: 普通; db: 快速; table: 极速  （默认是普通）
	RollbackMode *string `json:"RollbackMode,omitempty" name:"RollbackMode"`
}

func (r *RollBackClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollBackClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "RollbackStrategy")
	delete(f, "RollbackId")
	delete(f, "ExpectTime")
	delete(f, "ExpectTimeThresh")
	delete(f, "RollbackDatabases")
	delete(f, "RollbackTables")
	delete(f, "RollbackMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollBackClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollBackClusterResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RollBackClusterResponse struct {
	*tchttp.BaseResponse
	Response *RollBackClusterResponseParams `json:"Response"`
}

func (r *RollBackClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollBackClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollbackDatabase struct {
	// 旧数据库名称
	OldDatabase *string `json:"OldDatabase,omitempty" name:"OldDatabase"`

	// 新数据库名称
	NewDatabase *string `json:"NewDatabase,omitempty" name:"NewDatabase"`
}

type RollbackTable struct {
	// 数据库名称
	Database *string `json:"Database,omitempty" name:"Database"`

	// 数据库表
	Tables []*RollbackTableInfo `json:"Tables,omitempty" name:"Tables"`
}

type RollbackTableInfo struct {
	// 旧表名称
	OldTable *string `json:"OldTable,omitempty" name:"OldTable"`

	// 新表名称
	NewTable *string `json:"NewTable,omitempty" name:"NewTable"`
}

type RollbackTimeRange struct {
	// 开始时间
	TimeRangeStart *string `json:"TimeRangeStart,omitempty" name:"TimeRangeStart"`

	// 结束时间
	TimeRangeEnd *string `json:"TimeRangeEnd,omitempty" name:"TimeRangeEnd"`
}

type RuleFilters struct {
	// 审计规则过滤条件的参数名称。可选值：host – 客户端 IP；user – 数据库账户；dbName – 数据库名称；sqlType-SQL类型；sql-sql语句。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 审计规则过滤条件的匹配类型。可选值：INC – 包含；EXC – 不包含；EQS – 等于；NEQ – 不等于。
	Compare *string `json:"Compare,omitempty" name:"Compare"`

	// 审计规则过滤条件的匹配值。
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type SalePackageSpec struct {
	// 资源包使用地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageRegion *string `json:"PackageRegion,omitempty" name:"PackageRegion"`

	// 资源包类型
	// CCU-计算资源包
	// DISK-存储资源包
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// 资源包版本
	// base-基础版本，common-通用版本，enterprise-企业版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

	// 当前版本资源包最小资源数，计算资源单位：个；存储资源：GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinPackageSpec *float64 `json:"MinPackageSpec,omitempty" name:"MinPackageSpec"`

	// 当前版本资源包最大资源数，计算资源单位：个；存储资源：GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxPackageSpec *float64 `json:"MaxPackageSpec,omitempty" name:"MaxPackageSpec"`

	// 资源包有效期，单位:天
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireDay *int64 `json:"ExpireDay,omitempty" name:"ExpireDay"`
}

type SaleRegion struct {
	// 地域英文名
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域数字ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 地域中文名
	RegionZh *string `json:"RegionZh,omitempty" name:"RegionZh"`

	// 可售卖可用区列表
	ZoneSet []*SaleZone `json:"ZoneSet,omitempty" name:"ZoneSet"`

	// 引擎类型
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 地域模块支持情况
	Modules []*Module `json:"Modules,omitempty" name:"Modules"`
}

type SaleZone struct {
	// 可用区英文名
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区数字ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 可用区中文名
	ZoneZh *string `json:"ZoneZh,omitempty" name:"ZoneZh"`

	// 是否支持serverless集群<br>
	// 0:不支持<br>
	// 1:支持
	IsSupportServerless *int64 `json:"IsSupportServerless,omitempty" name:"IsSupportServerless"`

	// 是否支持普通集群<br>
	// 0:不支持<br>
	// 1:支持
	IsSupportNormal *int64 `json:"IsSupportNormal,omitempty" name:"IsSupportNormal"`

	// 物理区
	PhysicalZone *string `json:"PhysicalZone,omitempty" name:"PhysicalZone"`

	// 用户是否有可用区权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasPermission *bool `json:"HasPermission,omitempty" name:"HasPermission"`

	// 是否为全链路RDMA可用区
	IsWholeRdmaZone *string `json:"IsWholeRdmaZone,omitempty" name:"IsWholeRdmaZone"`
}

// Predefined struct for user
type SearchClusterDatabasesRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库名
	Database *string `json:"Database,omitempty" name:"Database"`

	// 是否精确搜索。
	// 0: 模糊搜索 1:精确搜索 
	// 默认为0
	MatchType *int64 `json:"MatchType,omitempty" name:"MatchType"`
}

type SearchClusterDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库名
	Database *string `json:"Database,omitempty" name:"Database"`

	// 是否精确搜索。
	// 0: 模糊搜索 1:精确搜索 
	// 默认为0
	MatchType *int64 `json:"MatchType,omitempty" name:"MatchType"`
}

func (r *SearchClusterDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClusterDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Database")
	delete(f, "MatchType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchClusterDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchClusterDatabasesResponseParams struct {
	// 数据库列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Databases []*string `json:"Databases,omitempty" name:"Databases"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchClusterDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *SearchClusterDatabasesResponseParams `json:"Response"`
}

func (r *SearchClusterDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClusterDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchClusterTablesRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库名
	Database *string `json:"Database,omitempty" name:"Database"`

	// 数据表名
	Table *string `json:"Table,omitempty" name:"Table"`

	// 数据表类型：
	// view：只返回 view，
	// base_table： 只返回基本表，
	// all：返回 view 和表
	TableType *string `json:"TableType,omitempty" name:"TableType"`
}

type SearchClusterTablesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库名
	Database *string `json:"Database,omitempty" name:"Database"`

	// 数据表名
	Table *string `json:"Table,omitempty" name:"Table"`

	// 数据表类型：
	// view：只返回 view，
	// base_table： 只返回基本表，
	// all：返回 view 和表
	TableType *string `json:"TableType,omitempty" name:"TableType"`
}

func (r *SearchClusterTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClusterTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Database")
	delete(f, "Table")
	delete(f, "TableType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchClusterTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchClusterTablesResponseParams struct {
	// 数据表列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tables []*DatabaseTables `json:"Tables,omitempty" name:"Tables"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchClusterTablesResponse struct {
	*tchttp.BaseResponse
	Response *SearchClusterTablesResponseParams `json:"Response"`
}

func (r *SearchClusterTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClusterTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroup struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 创建时间，时间格式：yyyy-mm-dd hh:mm:ss
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 入站规则
	Inbound []*PolicyRule `json:"Inbound,omitempty" name:"Inbound"`

	// 出站规则
	Outbound []*PolicyRule `json:"Outbound,omitempty" name:"Outbound"`

	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组名称
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// 安全组备注
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`
}

// Predefined struct for user
type SetRenewFlagRequestParams struct {
	// 需操作的实例ID
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 自动续费标志位，续费标记 0:正常续费  1:自动续费 2:到期不续
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

type SetRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 需操作的实例ID
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 自动续费标志位，续费标记 0:正常续费  1:自动续费 2:到期不续
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

func (r *SetRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceIds")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetRenewFlagResponseParams struct {
	// 操作成功实例数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *SetRenewFlagResponseParams `json:"Response"`
}

func (r *SetRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SlowQueriesItem struct {
	// 执行时间戳
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 执行时长，单位秒
	QueryTime *float64 `json:"QueryTime,omitempty" name:"QueryTime"`

	// sql语句
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// 客户端host
	UserHost *string `json:"UserHost,omitempty" name:"UserHost"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 数据库名
	Database *string `json:"Database,omitempty" name:"Database"`

	// 锁时长，单位秒
	LockTime *float64 `json:"LockTime,omitempty" name:"LockTime"`

	// 扫描行数
	RowsExamined *int64 `json:"RowsExamined,omitempty" name:"RowsExamined"`

	// 返回行数
	RowsSent *int64 `json:"RowsSent,omitempty" name:"RowsSent"`

	// sql模版
	SqlTemplate *string `json:"SqlTemplate,omitempty" name:"SqlTemplate"`

	// sql语句md5
	SqlMd5 *string `json:"SqlMd5,omitempty" name:"SqlMd5"`
}

// Predefined struct for user
type SwitchClusterVpcRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 字符串vpc id
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 字符串子网id
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 旧地址回收时间
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitempty" name:"OldIpReserveHours"`
}

type SwitchClusterVpcRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 字符串vpc id
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 字符串子网id
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 旧地址回收时间
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitempty" name:"OldIpReserveHours"`
}

func (r *SwitchClusterVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchClusterVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "OldIpReserveHours")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchClusterVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchClusterVpcResponseParams struct {
	// 异步任务id。
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwitchClusterVpcResponse struct {
	*tchttp.BaseResponse
	Response *SwitchClusterVpcResponseParams `json:"Response"`
}

func (r *SwitchClusterVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchClusterVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchClusterZoneRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 当前可用区
	OldZone *string `json:"OldZone,omitempty" name:"OldZone"`

	// 要切换到的可用区
	NewZone *string `json:"NewZone,omitempty" name:"NewZone"`

	// 维护期间执行-yes,立即执行-no
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitempty" name:"IsInMaintainPeriod"`
}

type SwitchClusterZoneRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 当前可用区
	OldZone *string `json:"OldZone,omitempty" name:"OldZone"`

	// 要切换到的可用区
	NewZone *string `json:"NewZone,omitempty" name:"NewZone"`

	// 维护期间执行-yes,立即执行-no
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitempty" name:"IsInMaintainPeriod"`
}

func (r *SwitchClusterZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchClusterZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "OldZone")
	delete(f, "NewZone")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchClusterZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchClusterZoneResponseParams struct {
	// 异步FlowId
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwitchClusterZoneResponse struct {
	*tchttp.BaseResponse
	Response *SwitchClusterZoneResponseParams `json:"Response"`
}

func (r *SwitchClusterZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchClusterZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchProxyVpcRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 字符串vpc id
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 字符串子网id
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 旧地址回收时间
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitempty" name:"OldIpReserveHours"`

	// 数据库代理组Id（该参数为必填项，可以通过DescribeProxies接口获得）
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`
}

type SwitchProxyVpcRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 字符串vpc id
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 字符串子网id
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 旧地址回收时间
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitempty" name:"OldIpReserveHours"`

	// 数据库代理组Id（该参数为必填项，可以通过DescribeProxies接口获得）
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`
}

func (r *SwitchProxyVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchProxyVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "OldIpReserveHours")
	delete(f, "ProxyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchProxyVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchProxyVpcResponseParams struct {
	// 异步任务id。
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwitchProxyVpcResponse struct {
	*tchttp.BaseResponse
	Response *SwitchProxyVpcResponseParams `json:"Response"`
}

func (r *SwitchProxyVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchProxyVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TablePrivileges struct {
	// 数据库名
	Db *string `json:"Db,omitempty" name:"Db"`

	// 表名
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 权限列表
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`
}

type Tag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TemplateParamInfo struct {
	// 当前值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 默认值
	Default *string `json:"Default,omitempty" name:"Default"`

	// 参数类型为enum时可选的值类型集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`

	// 参数类型为float/integer时的最大值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Max *string `json:"Max,omitempty" name:"Max"`

	// 参数类型为float/integer时的最小值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Min *string `json:"Min,omitempty" name:"Min"`

	// 参数名称
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 是否需要重启
	NeedReboot *int64 `json:"NeedReboot,omitempty" name:"NeedReboot"`

	// 参数描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 参数类型，integer/float/string/enum
	ParamType *string `json:"ParamType,omitempty" name:"ParamType"`
}

type TradePrice struct {
	// 预付费模式下资源总价，不包含优惠，单位:分
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPrice *int64 `json:"TotalPrice,omitempty" name:"TotalPrice"`

	// 总的折扣，100表示100%不打折
	Discount *float64 `json:"Discount,omitempty" name:"Discount"`

	// 预付费模式下的优惠后总价, 单位: 分,例如用户享有折扣 =TotalPrice × Discount
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPriceDiscount *int64 `json:"TotalPriceDiscount,omitempty" name:"TotalPriceDiscount"`

	// 后付费模式下的单位资源价格，不包含优惠，单位:分
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPrice *int64 `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// 优惠后后付费模式下的单位资源价格, 单位: 分,例如用户享有折扣=UnitPricet × Discount
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceDiscount *int64 `json:"UnitPriceDiscount,omitempty" name:"UnitPriceDiscount"`

	// 计费价格单位
	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`
}

// Predefined struct for user
type UnbindClusterResourcePackagesRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 资源包唯一ID,如果不传，解绑该实例绑定的所有资源包
	PackageIds []*string `json:"PackageIds,omitempty" name:"PackageIds"`
}

type UnbindClusterResourcePackagesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 资源包唯一ID,如果不传，解绑该实例绑定的所有资源包
	PackageIds []*string `json:"PackageIds,omitempty" name:"PackageIds"`
}

func (r *UnbindClusterResourcePackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindClusterResourcePackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "PackageIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindClusterResourcePackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindClusterResourcePackagesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindClusterResourcePackagesResponse struct {
	*tchttp.BaseResponse
	Response *UnbindClusterResourcePackagesResponseParams `json:"Response"`
}

func (r *UnbindClusterResourcePackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindClusterResourcePackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeClusterVersionRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 内核版本
	CynosVersion *string `json:"CynosVersion,omitempty" name:"CynosVersion"`

	// 升级时间类型，可选：upgradeImmediate,upgradeInMaintain
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`
}

type UpgradeClusterVersionRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 内核版本
	CynosVersion *string `json:"CynosVersion,omitempty" name:"CynosVersion"`

	// 升级时间类型，可选：upgradeImmediate,upgradeInMaintain
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`
}

func (r *UpgradeClusterVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeClusterVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "CynosVersion")
	delete(f, "UpgradeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeClusterVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeClusterVersionResponseParams struct {
	// 异步任务id
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpgradeClusterVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeClusterVersionResponseParams `json:"Response"`
}

func (r *UpgradeClusterVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeClusterVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库CPU
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 数据库内存，单位GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 升级类型：upgradeImmediate，upgradeInMaintain
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`

	// 该参数已废弃
	StorageLimit *uint64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 该参数已废弃
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 交易模式 0-下单并支付 1-下单
	DealMode *int64 `json:"DealMode,omitempty" name:"DealMode"`

	// NormalUpgrade：普通变配，FastUpgrade：极速变配，若变配过程判断会造成闪断，变配流程会终止。
	UpgradeMode *string `json:"UpgradeMode,omitempty" name:"UpgradeMode"`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库CPU
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 数据库内存，单位GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 升级类型：upgradeImmediate，upgradeInMaintain
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`

	// 该参数已废弃
	StorageLimit *uint64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 该参数已废弃
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 交易模式 0-下单并支付 1-下单
	DealMode *int64 `json:"DealMode,omitempty" name:"DealMode"`

	// NormalUpgrade：普通变配，FastUpgrade：极速变配，若变配过程判断会造成闪断，变配流程会终止。
	UpgradeMode *string `json:"UpgradeMode,omitempty" name:"UpgradeMode"`
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
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "UpgradeType")
	delete(f, "StorageLimit")
	delete(f, "AutoVoucher")
	delete(f, "DbType")
	delete(f, "DealMode")
	delete(f, "UpgradeMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceResponseParams struct {
	// 冻结流水ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranId *string `json:"TranId,omitempty" name:"TranId"`

	// 大订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`

	// 订单号
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type UpgradeProxyRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// cpu核数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存
	Mem *int64 `json:"Mem,omitempty" name:"Mem"`

	// 数据库代理组节点个数
	ProxyCount *int64 `json:"ProxyCount,omitempty" name:"ProxyCount"`

	// 数据库代理组ID（已废弃）
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 重新负载均衡：auto（自动），manual（手动）
	ReloadBalance *string `json:"ReloadBalance,omitempty" name:"ReloadBalance"`

	// 升级时间 ：no（升级完成时）yes（实例维护时间）
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitempty" name:"IsInMaintainPeriod"`

	// 数据库代理节点信息
	ProxyZones []*ProxyZone `json:"ProxyZones,omitempty" name:"ProxyZones"`
}

type UpgradeProxyRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// cpu核数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存
	Mem *int64 `json:"Mem,omitempty" name:"Mem"`

	// 数据库代理组节点个数
	ProxyCount *int64 `json:"ProxyCount,omitempty" name:"ProxyCount"`

	// 数据库代理组ID（已废弃）
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 重新负载均衡：auto（自动），manual（手动）
	ReloadBalance *string `json:"ReloadBalance,omitempty" name:"ReloadBalance"`

	// 升级时间 ：no（升级完成时）yes（实例维护时间）
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitempty" name:"IsInMaintainPeriod"`

	// 数据库代理节点信息
	ProxyZones []*ProxyZone `json:"ProxyZones,omitempty" name:"ProxyZones"`
}

func (r *UpgradeProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Cpu")
	delete(f, "Mem")
	delete(f, "ProxyCount")
	delete(f, "ProxyGroupId")
	delete(f, "ReloadBalance")
	delete(f, "IsInMaintainPeriod")
	delete(f, "ProxyZones")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeProxyResponseParams struct {
	// 异步流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 异步任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpgradeProxyResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeProxyResponseParams `json:"Response"`
}

func (r *UpgradeProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeProxyVersionRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库代理当前版本
	SrcProxyVersion *string `json:"SrcProxyVersion,omitempty" name:"SrcProxyVersion"`

	// 数据库代理升级版本
	DstProxyVersion *string `json:"DstProxyVersion,omitempty" name:"DstProxyVersion"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 升级时间 ：no（升级完成时）yes（实例维护时间）
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitempty" name:"IsInMaintainPeriod"`
}

type UpgradeProxyVersionRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库代理当前版本
	SrcProxyVersion *string `json:"SrcProxyVersion,omitempty" name:"SrcProxyVersion"`

	// 数据库代理升级版本
	DstProxyVersion *string `json:"DstProxyVersion,omitempty" name:"DstProxyVersion"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`

	// 升级时间 ：no（升级完成时）yes（实例维护时间）
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitempty" name:"IsInMaintainPeriod"`
}

func (r *UpgradeProxyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeProxyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SrcProxyVersion")
	delete(f, "DstProxyVersion")
	delete(f, "ProxyGroupId")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeProxyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeProxyVersionResponseParams struct {
	// 异步流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 异步任务id
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpgradeProxyVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeProxyVersionResponseParams `json:"Response"`
}

func (r *UpgradeProxyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeProxyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserHostPrivilege struct {
	// 授权用户
	DbUserName *string `json:"DbUserName,omitempty" name:"DbUserName"`

	// 客户端ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbHost *string `json:"DbHost,omitempty" name:"DbHost"`

	// 用户权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbPrivilege *string `json:"DbPrivilege,omitempty" name:"DbPrivilege"`
}

type ZoneStockInfo struct {
	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 是否有库存
	HasStock *bool `json:"HasStock,omitempty" name:"HasStock"`

	// 库存数量
	StockCount *int64 `json:"StockCount,omitempty" name:"StockCount"`
}