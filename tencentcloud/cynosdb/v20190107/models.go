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

package v20190107

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Ability struct {
	// 是否支持从可用区
	IsSupportSlaveZone *string `json:"IsSupportSlaveZone,omitnil,omitempty" name:"IsSupportSlaveZone"`

	// 不支持从可用区的原因
	NonsupportSlaveZoneReason *string `json:"NonsupportSlaveZoneReason,omitnil,omitempty" name:"NonsupportSlaveZoneReason"`

	// 是否支持RO实例
	IsSupportRo *string `json:"IsSupportRo,omitnil,omitempty" name:"IsSupportRo"`

	// 不支持RO实例的原因
	NonsupportRoReason *string `json:"NonsupportRoReason,omitnil,omitempty" name:"NonsupportRoReason"`

	// 是否支持手动发起快照备份
	IsSupportManualSnapshot *string `json:"IsSupportManualSnapshot,omitnil,omitempty" name:"IsSupportManualSnapshot"`

	// 是否支持透明数据加密
	IsSupportTransparentDataEncryption *string `json:"IsSupportTransparentDataEncryption,omitnil,omitempty" name:"IsSupportTransparentDataEncryption"`

	// 不支持透明数据加密原因
	NoSupportTransparentDataEncryptionReason *string `json:"NoSupportTransparentDataEncryptionReason,omitnil,omitempty" name:"NoSupportTransparentDataEncryptionReason"`

	// 是否支持手动发起逻辑备份
	IsSupportManualLogic *string `json:"IsSupportManualLogic,omitnil,omitempty" name:"IsSupportManualLogic"`

	// 是否支持开启全局加密
	IsSupportGlobalEncryption *string `json:"IsSupportGlobalEncryption,omitnil,omitempty" name:"IsSupportGlobalEncryption"`

	// 不支持全局加密的原因
	NoSupportGlobalEncryptionReason *string `json:"NoSupportGlobalEncryptionReason,omitnil,omitempty" name:"NoSupportGlobalEncryptionReason"`

	// 不支持透明加密原因状态码
	NoSupportTransparentDataEncryptionReasonCode *string `json:"NoSupportTransparentDataEncryptionReasonCode,omitnil,omitempty" name:"NoSupportTransparentDataEncryptionReasonCode"`

	// 不支持全局加密原因状态码
	NoSupportGlobalEncryptionReasonCode *string `json:"NoSupportGlobalEncryptionReasonCode,omitnil,omitempty" name:"NoSupportGlobalEncryptionReasonCode"`
}

type Account struct {
	// 数据库账号名
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 数据库账号描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 主机
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 用户最大连接数
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`
}

type AccountParam struct {
	// 参数名称，当前仅支持参数：max_user_connections
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 参数值
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`
}

// Predefined struct for user
type ActivateInstanceRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例 ID 列表，单个实例 ID 格式如：cynosdbmysql-ins-n7ocdslw，与TDSQL-C MySQL数据库控制台页面中显示的实例 ID 相同，可使用 查询实例列表 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
}

type ActivateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例 ID 列表，单个实例 ID 格式如：cynosdbmysql-ins-n7ocdslw，与TDSQL-C MySQL数据库控制台页面中显示的实例 ID 相同，可使用 查询实例列表 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 从可用区
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// binlog同步方式。默认值：async。可选值：sync、semisync、async
	BinlogSyncWay *string `json:"BinlogSyncWay,omitnil,omitempty" name:"BinlogSyncWay"`

	// 半同步超时时间，单位ms。为保证业务稳定性，半同步复制存在退化逻辑，当主可用区集群在等待备可用区集群确认事务时若超过该超时时间，复制方式将降为异步复制。最低设置为1000ms，最高支持4294967295ms，默认10000ms。
	SemiSyncTimeout *int64 `json:"SemiSyncTimeout,omitnil,omitempty" name:"SemiSyncTimeout"`
}

type AddClusterSlaveZoneRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 从可用区
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// binlog同步方式。默认值：async。可选值：sync、semisync、async
	BinlogSyncWay *string `json:"BinlogSyncWay,omitnil,omitempty" name:"BinlogSyncWay"`

	// 半同步超时时间，单位ms。为保证业务稳定性，半同步复制存在退化逻辑，当主可用区集群在等待备可用区集群确认事务时若超过该超时时间，复制方式将降为异步复制。最低设置为1000ms，最高支持4294967295ms，默认10000ms。
	SemiSyncTimeout *int64 `json:"SemiSyncTimeout,omitnil,omitempty" name:"SemiSyncTimeout"`
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
	delete(f, "BinlogSyncWay")
	delete(f, "SemiSyncTimeout")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddClusterSlaveZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddClusterSlaveZoneResponseParams struct {
	// 异步FlowId
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cpu核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存，单位为GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 新增只读实例数，取值范围为(0,15]
	ReadOnlyCount *int64 `json:"ReadOnlyCount,omitnil,omitempty" name:"ReadOnlyCount"`

	// 实例机器类型，支持值如下：
	// - common：表示通用型
	// - exclusive：表示独享型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 实例组ID，在已有RO组中新增实例时使用，不传则新增RO组。当前版本不建议传输该值。
	//
	// Deprecated: InstanceGrpId is deprecated.
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// 所属VPC网络ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 所属子网ID，如果设置了VpcId，则SubnetId必填。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 新增RO组时使用的Port，取值范围为[0,65535)
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 实例名称，字符串长度范围为[0,64)，取值范围为大小写字母，0-9数字，'_','-','.'
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 订单来源，字符串长度范围为[0,64)
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// 交易模式 0-下单并支付 1-下单
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// 参数模板ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 参数列表，ParamTemplateId 传入时InstanceParams才有效
	InstanceParams []*ModifyParamItem `json:"InstanceParams,omitnil,omitempty" name:"InstanceParams"`

	// 安全组ID，新建只读实例时可以指定安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// proxy同步升级
	UpgradeProxy *UpgradeProxy `json:"UpgradeProxy,omitnil,omitempty" name:"UpgradeProxy"`
}

type AddInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cpu核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存，单位为GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 新增只读实例数，取值范围为(0,15]
	ReadOnlyCount *int64 `json:"ReadOnlyCount,omitnil,omitempty" name:"ReadOnlyCount"`

	// 实例机器类型，支持值如下：
	// - common：表示通用型
	// - exclusive：表示独享型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 实例组ID，在已有RO组中新增实例时使用，不传则新增RO组。当前版本不建议传输该值。
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// 所属VPC网络ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 所属子网ID，如果设置了VpcId，则SubnetId必填。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 新增RO组时使用的Port，取值范围为[0,65535)
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 实例名称，字符串长度范围为[0,64)，取值范围为大小写字母，0-9数字，'_','-','.'
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 订单来源，字符串长度范围为[0,64)
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// 交易模式 0-下单并支付 1-下单
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// 参数模板ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 参数列表，ParamTemplateId 传入时InstanceParams才有效
	InstanceParams []*ModifyParamItem `json:"InstanceParams,omitnil,omitempty" name:"InstanceParams"`

	// 安全组ID，新建只读实例时可以指定安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// proxy同步升级
	UpgradeProxy *UpgradeProxy `json:"UpgradeProxy,omitnil,omitempty" name:"UpgradeProxy"`
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
	delete(f, "DeviceType")
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
	delete(f, "UpgradeProxy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddInstancesResponseParams struct {
	// 冻结流水，一次开通一个冻结流水。
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// 后付费订单号。
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 发货资源id列表。
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 大订单号
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// IP地址
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// 端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

// Predefined struct for user
type AssociateSecurityGroupsRequestParams struct {
	// 实例组 ID 数组，cynosdbmysql-grp-前缀开头或集群 ID。
	// 说明：要获取集群的实例组 ID，可通过 [查询集群实例组](https://cloud.tencent.com/document/product/1003/103934) 进行。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 要修改的安全组ID列表，一个或者多个安全组Id组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例组 ID 数组，cynosdbmysql-grp-前缀开头或集群 ID。
	// 说明：要获取集群的实例组 ID，可通过 [查询集群实例组](https://cloud.tencent.com/document/product/1003/103934) 进行。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 要修改的安全组ID列表，一个或者多个安全组Id组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type AuditInstanceFilters struct {
	// 过滤条件值。支持InstanceId-实例ID，InstanceName-实例名称，ProjectId-项目ID，TagKey-标签键，Tag-标签（以竖线分割，例：Tagkey|Tagvalue）。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// true表示精确查找，false表示模糊匹配。
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`

	// 筛选值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type AuditInstanceInfo struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 标签信息
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type AuditLog struct {
	// 影响行数。
	AffectRows *int64 `json:"AffectRows,omitnil,omitempty" name:"AffectRows"`

	// 错误码。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// SQL类型。
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// 表名称。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 审计策略名称。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 数据库名称。
	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`

	// SQL语句。
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 客户端地址。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 用户名。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 执行时间，微秒。
	ExecTime *int64 `json:"ExecTime,omitnil,omitempty" name:"ExecTime"`

	// 时间。
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 返回行数。
	SentRows *int64 `json:"SentRows,omitnil,omitempty" name:"SentRows"`

	// 执行线程ID。
	ThreadId *int64 `json:"ThreadId,omitnil,omitempty" name:"ThreadId"`

	// 扫描行数。
	CheckRows *int64 `json:"CheckRows,omitnil,omitempty" name:"CheckRows"`

	// cpu执行时间，微秒。
	CpuTime *float64 `json:"CpuTime,omitnil,omitempty" name:"CpuTime"`

	// IO等待时间，微秒。
	IoWaitTime *int64 `json:"IoWaitTime,omitnil,omitempty" name:"IoWaitTime"`

	// 锁等待时间，微秒。
	LockWaitTime *int64 `json:"LockWaitTime,omitnil,omitempty" name:"LockWaitTime"`

	// 事务持续等待时间，微秒。
	TrxLivingTime *int64 `json:"TrxLivingTime,omitnil,omitempty" name:"TrxLivingTime"`

	// 开始时间，与timestamp构成一个精确到纳秒的时间。
	NsTime *int64 `json:"NsTime,omitnil,omitempty" name:"NsTime"`

	// 日志命中规则模板的基本信息
	TemplateInfo []*LogRuleTemplateInfo `json:"TemplateInfo,omitnil,omitempty" name:"TemplateInfo"`

	// 事务ID
	TrxId *int64 `json:"TrxId,omitnil,omitempty" name:"TrxId"`
}

type AuditLogFile struct {
	// 审计日志文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 审计日志文件创建时间。格式为 : "2019-03-20 17:09:13"。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 文件状态值。可能返回的值为：
	// "creating" - 生成中;
	// "failed" - 创建失败;
	// "success" - 已生成;
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 文件大小，单位为 KB。
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 审计日志下载地址。
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 错误信息。
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 日志下载进度。
	ProgressRate *int64 `json:"ProgressRate,omitnil,omitempty" name:"ProgressRate"`
}

type AuditLogFilter struct {
	// 客户端地址。
	Host []*string `json:"Host,omitnil,omitempty" name:"Host"`

	// 用户名。
	User []*string `json:"User,omitnil,omitempty" name:"User"`

	// 数据库名称。
	DBName []*string `json:"DBName,omitnil,omitempty" name:"DBName"`

	// 表名称。
	TableName []*string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 审计策略名称。
	PolicyName []*string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// SQL 语句。支持模糊匹配。
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// SQL 类型。目前支持："SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "ALTER", "SET", "REPLACE", "EXECUTE"。
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// 执行时间。单位为：ms。表示筛选执行时间大于该值的审计日志。
	ExecTime *int64 `json:"ExecTime,omitnil,omitempty" name:"ExecTime"`

	// 影响行数。表示筛选影响行数大于该值的审计日志。
	AffectRows *int64 `json:"AffectRows,omitnil,omitempty" name:"AffectRows"`

	// SQL 类型。支持多个类型同时查询。目前支持："SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "ALTER", "SET", "REPLACE", "EXECUTE"。
	SqlTypes []*string `json:"SqlTypes,omitnil,omitempty" name:"SqlTypes"`

	// SQL 语句。支持传递多个sql语句。
	Sqls []*string `json:"Sqls,omitnil,omitempty" name:"Sqls"`

	// 返回行数。
	SentRows *uint64 `json:"SentRows,omitnil,omitempty" name:"SentRows"`

	// 线程ID。
	ThreadId []*string `json:"ThreadId,omitnil,omitempty" name:"ThreadId"`
}

type AuditRuleFilters struct {
	// 单条审计规则。
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`
}

type AuditRuleTemplateInfo struct {
	// 规则模板ID。
	RuleTemplateId *string `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// 规则模板名称。
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// 规则模板的过滤条件
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// 规则模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 规则模板创建时间。
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// 规则模板修改时间。
	UpdateAt *string `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`

	// 告警等级。1-低风险，2-中风险，3-高风险。
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警策略。0-不告警，1-告警。
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`

	// 模板状态。0-无任务 ，1-修改中。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则模板应用在哪些在实例。
	AffectedInstances []*string `json:"AffectedInstances,omitnil,omitempty" name:"AffectedInstances"`
}

type BackupFileInfo struct {
	// 快照文件ID，已废弃，请使用BackupId
	SnapshotId *uint64 `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 备份文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 备份文件大小
	FileSize *uint64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 备份开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 备份完成时间
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 备份类型：snapshot，快照备份；logic，逻辑备份
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 备份方式：auto，自动备份；manual，手动备份
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 备份文件状态：success：备份成功；fail：备份失败；creating：备份文件创建中；deleting：备份文件删除中
	BackupStatus *string `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// 备份文件时间
	SnapshotTime *string `json:"SnapshotTime,omitnil,omitempty" name:"SnapshotTime"`

	// 备份ID
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 快照类型，可选值：full，全量；increment，增量
	SnapShotType *string `json:"SnapShotType,omitnil,omitempty" name:"SnapShotType"`

	// 备份文件备注
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
}

type BackupLimitClusterRestriction struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 下载限制配置
	BackupLimitRestriction *BackupLimitRestriction `json:"BackupLimitRestriction,omitnil,omitempty" name:"BackupLimitRestriction"`
}

type BackupLimitRestriction struct {
	// 限制类型
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// 该参数仅支持 In， 表示 LimitVpc 指定的vpc可以下载。默认为In
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// In: 指定的ip可以下载； NotIn: 指定的ip不可以下载
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// 限制下载的vpc设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	LimitVpcs []*BackupLimitVpcItem `json:"LimitVpcs,omitnil,omitempty" name:"LimitVpcs"`

	// 限制下载的ip设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	LimitIps []*string `json:"LimitIps,omitnil,omitempty" name:"LimitIps"`
}

type BackupLimitVpcItem struct {
	// 限制下载来源的地域。目前仅支持当前地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 限制下载的vpc列表
	VpcList []*string `json:"VpcList,omitnil,omitempty" name:"VpcList"`
}

type BillingResourceInfo struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 订单ID
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`
}

// Predefined struct for user
type BindClusterResourcePackagesRequestParams struct {
	// 资源包唯一ID
	PackageIds []*string `json:"PackageIds,omitnil,omitempty" name:"PackageIds"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type BindClusterResourcePackagesRequest struct {
	*tchttp.BaseRequest
	
	// 资源包唯一ID
	PackageIds []*string `json:"PackageIds,omitnil,omitempty" name:"PackageIds"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 绑定的集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 绑定的实例所在的地域
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`

	// 绑定的实例类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 绑定集群下的实例ID
	ExtendIds []*string `json:"ExtendIds,omitnil,omitempty" name:"ExtendIds"`
}

type BinlogConfigInfo struct {
	// binlog保留时间
	BinlogSaveDays *int64 `json:"BinlogSaveDays,omitnil,omitempty" name:"BinlogSaveDays"`

	// binlog异地地域备份是否开启
	BinlogCrossRegionsEnable *string `json:"BinlogCrossRegionsEnable,omitnil,omitempty" name:"BinlogCrossRegionsEnable"`

	// binlog异地地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	BinlogCrossRegions []*string `json:"BinlogCrossRegions,omitnil,omitempty" name:"BinlogCrossRegions"`
}

type BinlogItem struct {
	// Binlog文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件大小，单位：字节
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 事务最早时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 事务最晚时间
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// Binlog文件ID
	BinlogId *int64 `json:"BinlogId,omitnil,omitempty" name:"BinlogId"`

	// binlog所跨地域
	CrossRegions []*string `json:"CrossRegions,omitnil,omitempty" name:"CrossRegions"`
}

type BizTaskInfo struct {
	// 任务id
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 用户appid
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 延迟执行时间
	DelayTime *string `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 任务失败信息
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 异步任务流id
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 任务输入信息
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// 实例组id
	//
	// Deprecated: InstanceGrpId is deprecated.
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// 实例组id
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务操作对象id
	ObjectId *string `json:"ObjectId,omitnil,omitempty" name:"ObjectId"`

	// 任务操作对象类型
	ObjectType *string `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`

	// 操作者uin
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 任务输出信息
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// 任务状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 触发本任务的父任务ID
	TriggerTaskId *int64 `json:"TriggerTaskId,omitnil,omitempty" name:"TriggerTaskId"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 任务开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 任务进度
	Process *int64 `json:"Process,omitnil,omitempty" name:"Process"`

	// 修改参数任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ModifyParamsData is deprecated.
	ModifyParamsData []*ModifyParamsData `json:"ModifyParamsData,omitnil,omitempty" name:"ModifyParamsData"`

	// 创建集群任务信息
	CreateClustersData *CreateClustersData `json:"CreateClustersData,omitnil,omitempty" name:"CreateClustersData"`

	// 集群回档任务信息
	RollbackData *RollbackData `json:"RollbackData,omitnil,omitempty" name:"RollbackData"`

	// 实例变配任务信息
	ModifyInstanceData *ModifyInstanceData `json:"ModifyInstanceData,omitnil,omitempty" name:"ModifyInstanceData"`

	// 手动备份任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManualBackupData *ManualBackupData `json:"ManualBackupData,omitnil,omitempty" name:"ManualBackupData"`

	// 修改内核版本任务信息
	ModifyDbVersionData *ModifyDbVersionData `json:"ModifyDbVersionData,omitnil,omitempty" name:"ModifyDbVersionData"`

	// 集群可用区信息
	ClusterSlaveData *ClusterSlaveData `json:"ClusterSlaveData,omitnil,omitempty" name:"ClusterSlaveData"`

	// 转换集群日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	SwitchClusterLogBin *SwitchClusterLogBin `json:"SwitchClusterLogBin,omitnil,omitempty" name:"SwitchClusterLogBin"`

	// 修改实例参数数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyInstanceParamsData *BizTaskModifyParamsData `json:"ModifyInstanceParamsData,omitnil,omitempty" name:"ModifyInstanceParamsData"`

	// 维护时间
	TaskMaintainInfo *TaskMaintainInfo `json:"TaskMaintainInfo,omitnil,omitempty" name:"TaskMaintainInfo"`

	// 实例日志投递信息
	InstanceCLSDeliveryInfos []*InstanceCLSDeliveryInfo `json:"InstanceCLSDeliveryInfos,omitnil,omitempty" name:"InstanceCLSDeliveryInfos"`

	// 任务进度信息
	TaskProgressInfo *TaskProgressInfo `json:"TaskProgressInfo,omitnil,omitempty" name:"TaskProgressInfo"`

	// 全球数据库网络任务
	GdnTaskInfo *GdnTaskInfo `json:"GdnTaskInfo,omitnil,omitempty" name:"GdnTaskInfo"`
}

type BizTaskModifyInstanceParam struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例参数修改任务详情
	ModifyInstanceParamList []*ModifyParamItem `json:"ModifyInstanceParamList,omitnil,omitempty" name:"ModifyInstanceParamList"`
}

type BizTaskModifyParamsData struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群参数修改数据
	ClusterParamList []*ModifyParamItem `json:"ClusterParamList,omitnil,omitempty" name:"ClusterParamList"`

	// 实例参数修改数据
	ModifyInstanceParams []*BizTaskModifyInstanceParam `json:"ModifyInstanceParams,omitnil,omitempty" name:"ModifyInstanceParams"`
}

type CLSInfo struct {
	// 日志主题操作：可选create,reuse。
	// create:新增日志主题，使用TopicName创建日志主题。
	// reuse:使用已有日志主题，使用TopicId指定日志主题。
	// 不允许使用已有日志主题且新建日志集的组合。
	TopicOperation *string `json:"TopicOperation,omitnil,omitempty" name:"TopicOperation"`

	// 日志集操作：可选create,reuse。
	// create:新增日志集，使用GroupName创建日志集。
	// reuse:使用已有日志集，使用GroupId指定日志集。
	// 不允许使用已有日志主题且新建日志集的组合。
	GroupOperation *string `json:"GroupOperation,omitnil,omitempty" name:"GroupOperation"`

	// 日志投递地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 日志主题id
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志主题name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 日志集id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 日志集name
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
}

// Predefined struct for user
type CloseAuditServiceRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CloseAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
}

type CloseClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID数组
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CloseProxyEndPointRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

type CloseProxyEndPointRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

func (r *CloseProxyEndPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProxyEndPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProxyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseProxyEndPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseProxyEndPointResponseParams struct {
	// 异步流程ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 异步任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseProxyEndPointResponse struct {
	*tchttp.BaseResponse
	Response *CloseProxyEndPointResponseParams `json:"Response"`
}

func (r *CloseProxyEndPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProxyEndPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseProxyRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 是否只关闭读写分离，取值：是 "true","false"
	OnlyCloseRW *bool `json:"OnlyCloseRW,omitnil,omitempty" name:"OnlyCloseRW"`
}

type CloseProxyRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 是否只关闭读写分离，取值：是 "true","false"
	OnlyCloseRW *bool `json:"OnlyCloseRW,omitnil,omitempty" name:"OnlyCloseRW"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 异步任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CloseSSLRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CloseSSLRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *CloseSSLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseSSLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseSSLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseSSLResponseParams struct {
	// 流程ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 任务id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseSSLResponse struct {
	*tchttp.BaseResponse
	Response *CloseSSLResponseParams `json:"Response"`
}

func (r *CloseSSLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseSSLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseWanRequestParams struct {
	// 实例组id
	//
	// Deprecated: InstanceGrpId is deprecated.
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// 实例组id
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CloseWanRequest struct {
	*tchttp.BaseRequest
	
	// 实例组id
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// 实例组id
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	delete(f, "InstanceGroupId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseWanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseWanResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 引擎类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例状态
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 实例状态描述
	InstanceStatusDesc *string `json:"InstanceStatusDesc,omitnil,omitempty" name:"InstanceStatusDesc"`

	// cpu核数
	InstanceCpu *int64 `json:"InstanceCpu,omitnil,omitempty" name:"InstanceCpu"`

	// 内存
	InstanceMemory *int64 `json:"InstanceMemory,omitnil,omitempty" name:"InstanceMemory"`

	// 硬盘
	InstanceStorage *int64 `json:"InstanceStorage,omitnil,omitempty" name:"InstanceStorage"`

	// 实例角色
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 执行开始时间(距离0点的秒数)	
	MaintainStartTime *int64 `json:"MaintainStartTime,omitnil,omitempty" name:"MaintainStartTime"`

	// 持续的时间(单位：秒)	
	MaintainDuration *int64 `json:"MaintainDuration,omitnil,omitempty" name:"MaintainDuration"`

	// 可以执行的时间，枚举值：["Mon","Tue","Wed","Thu","Fri", "Sat", "Sun"]
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitnil,omitempty" name:"MaintainWeekDays"`

	// serverless实例子状态
	ServerlessStatus *string `json:"ServerlessStatus,omitnil,omitempty" name:"ServerlessStatus"`

	// 实例任务信息
	InstanceTasks []*ObjectTask `json:"InstanceTasks,omitnil,omitempty" name:"InstanceTasks"`

	// 实例机器类型
	InstanceDeviceType *string `json:"InstanceDeviceType,omitnil,omitempty" name:"InstanceDeviceType"`

	// 实例存储类型
	// 说明：仅当要查询的资源为 LibraDB 时，此参数才会返回值。
	InstanceStorageType *string `json:"InstanceStorageType,omitnil,omitempty" name:"InstanceStorageType"`

	// 数据库类型
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// 节点列表
	// 说明：仅当要查询的资源为 LibraDB 时，此参数才会返回值。
	NodeList []*string `json:"NodeList,omitnil,omitempty" name:"NodeList"`
}

type ClusterParamModifyLog struct {
	// 参数名称
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 当前值
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 修改后的值
	UpdateValue *string `json:"UpdateValue,omitnil,omitempty" name:"UpdateValue"`

	// 修改状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ClusterReadOnlyValue struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 只读开关值
	ReadOnlyValue *string `json:"ReadOnlyValue,omitnil,omitempty" name:"ReadOnlyValue"`
}

type ClusterSlaveData struct {
	// 旧主可用区
	OldMasterZone *string `json:"OldMasterZone,omitnil,omitempty" name:"OldMasterZone"`

	// 旧从可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldSlaveZone []*string `json:"OldSlaveZone,omitnil,omitempty" name:"OldSlaveZone"`

	// 新主可用区
	NewMasterZone *string `json:"NewMasterZone,omitnil,omitempty" name:"NewMasterZone"`

	// 新从可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewSlaveZone []*string `json:"NewSlaveZone,omitnil,omitempty" name:"NewSlaveZone"`

	// 新从可用区属性
	NewSlaveZoneAttr []*SlaveZoneAttrItem `json:"NewSlaveZoneAttr,omitnil,omitempty" name:"NewSlaveZoneAttr"`

	// 旧可用区属性
	OldSlaveZoneAttr []*SlaveZoneAttrItem `json:"OldSlaveZoneAttr,omitnil,omitempty" name:"OldSlaveZoneAttr"`
}

type ClusterTaskId struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type CopyClusterPasswordComplexityRequestParams struct {
	// 复制集群ID数组，例如["cynosdbmysql-bzxxrmtq","cynosdbmysql-qwer"]
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 集群id，例如"cynosdbmysql-bzxxrmtq"
	SourceClusterId *string `json:"SourceClusterId,omitnil,omitempty" name:"SourceClusterId"`
}

type CopyClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// 复制集群ID数组，例如["cynosdbmysql-bzxxrmtq","cynosdbmysql-qwer"]
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 集群id，例如"cynosdbmysql-bzxxrmtq"
	SourceClusterId *string `json:"SourceClusterId,omitnil,omitempty" name:"SourceClusterId"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 新账户列表
	Accounts []*NewAccount `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

type CreateAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 新账户列表
	Accounts []*NewAccount `json:"Accounts,omitnil,omitempty" name:"Accounts"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，格式为："2017-07-12 10:29:20"。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，格式为："2017-07-12 10:29:20"。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序方式。支持值包括："ASC" - 升序，"DESC" - 降序。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段。支持值包括：
	// "timestamp" - 时间戳；
	// "affectRows" - 影响行数；
	// "execTime" - 执行时间。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 已废弃。
	//
	// Deprecated: Filter is deprecated.
	Filter *AuditLogFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 审计日志过滤条件
	LogFilter []*InstanceAuditLogFilter `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`

	// 审计日志列
	ColumnFilter []*string `json:"ColumnFilter,omitnil,omitempty" name:"ColumnFilter"`
}

type CreateAuditLogFileRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，格式为："2017-07-12 10:29:20"。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，格式为："2017-07-12 10:29:20"。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序方式。支持值包括："ASC" - 升序，"DESC" - 降序。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段。支持值包括：
	// "timestamp" - 时间戳；
	// "affectRows" - 影响行数；
	// "execTime" - 执行时间。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 已废弃。
	Filter *AuditLogFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 审计日志过滤条件
	LogFilter []*InstanceAuditLogFilter `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`

	// 审计日志列
	ColumnFilter []*string `json:"ColumnFilter,omitnil,omitempty" name:"ColumnFilter"`
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
	delete(f, "LogFilter")
	delete(f, "ColumnFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditLogFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditLogFileResponseParams struct {
	// 审计日志文件名称。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// 规则模板名称。
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// 规则模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 告警等级。1-低风险，2-中风险，3-高风险
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警策略。0-不告警，1-告警。
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`
}

type CreateAuditRuleTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 审计规则。
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// 规则模板名称。
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// 规则模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 告警等级。1-低风险，2-中风险，3-高风险
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警策略。0-不告警，1-告警。
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`
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
	delete(f, "AlarmLevel")
	delete(f, "AlarmPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditRuleTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditRuleTemplateResponseParams struct {
	// 生成的规则模板ID。
	RuleTemplateId *string `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备份类型, 可选值：logic，逻辑备份；snapshot，物理备份
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 备份的库, 只在 BackupType 为 logic 时有效
	BackupDatabases []*string `json:"BackupDatabases,omitnil,omitempty" name:"BackupDatabases"`

	// 备份的表, 只在 BackupType 为 logic 时有效
	BackupTables []*DatabaseTables `json:"BackupTables,omitnil,omitempty" name:"BackupTables"`

	// 备注名
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备份类型, 可选值：logic，逻辑备份；snapshot，物理备份
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 备份的库, 只在 BackupType 为 logic 时有效
	BackupDatabases []*string `json:"BackupDatabases,omitnil,omitempty" name:"BackupDatabases"`

	// 备份的表, 只在 BackupType 为 logic 时有效
	BackupTables []*DatabaseTables `json:"BackupTables,omitnil,omitempty" name:"BackupTables"`

	// 备注名
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateCLSDeliveryRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志投递配置
	CLSInfoList []*CLSInfo `json:"CLSInfoList,omitnil,omitempty" name:"CLSInfoList"`

	// 日志类型
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 是否维护时间运行
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type CreateCLSDeliveryRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志投递配置
	CLSInfoList []*CLSInfo `json:"CLSInfoList,omitnil,omitempty" name:"CLSInfoList"`

	// 日志类型
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 是否维护时间运行
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

func (r *CreateCLSDeliveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCLSDeliveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CLSInfoList")
	delete(f, "LogType")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCLSDeliveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCLSDeliveryResponseParams struct {
	// 异步任务id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCLSDeliveryResponse struct {
	*tchttp.BaseResponse
	Response *CreateCLSDeliveryResponseParams `json:"Response"`
}

func (r *CreateCLSDeliveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCLSDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterDatabaseRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库名
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 字符集类型
	CharacterSet *string `json:"CharacterSet,omitnil,omitempty" name:"CharacterSet"`

	// 排序规则
	CollateRule *string `json:"CollateRule,omitnil,omitempty" name:"CollateRule"`

	// 授权用户主机权限
	UserHostPrivileges []*UserHostPrivilege `json:"UserHostPrivileges,omitnil,omitempty" name:"UserHostPrivileges"`

	// 备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateClusterDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库名
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 字符集类型
	CharacterSet *string `json:"CharacterSet,omitnil,omitempty" name:"CharacterSet"`

	// 排序规则
	CollateRule *string `json:"CollateRule,omitnil,omitempty" name:"CollateRule"`

	// 授权用户主机权限
	UserHostPrivileges []*UserHostPrivilege `json:"UserHostPrivileges,omitnil,omitempty" name:"UserHostPrivileges"`

	// 备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type CreateClustersData struct {
	// 实例CPU
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例内存
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 集群存储上限
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`
}

// Predefined struct for user
type CreateClustersRequestParams struct {
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 所属VPC网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 所属子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 数据库版本，取值范围: 
	// <li> MYSQL可选值：5.7，8.0 </li>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// 所属项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 当DbMode为NORMAL或不填时必选
	// 普通实例Cpu核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 当DbMode为NORMAL或不填时必选
	// 普通实例内存,单位GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例数量，数量范围为(0,16]，默认值为2（即一个rw实例+一个ro实例），传递的n表示1个rw实例+n-1个ro实例（规格相同），如需要更精确的集群组成搭配，请使用InstanceInitInfos
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 该参数无实际意义，已废弃。
	// 存储大小，单位GB。
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 集群名称，长度小于64个字符，每个字符取值范围：大/小写字母，数字，特殊符号（'-','_','.'）
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 账号密码(8-64个字符，包含大小写英文字母、数字和符号~!@#$%^&*_-+=`|\(){}[]:;'<>,.?/中的任意三种)
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// 端口，默认3306，取值范围[0, 65535)
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 计费模式，支持值为0和1，默认值为0。
	// 取值为0，表示按量计费。
	// 取值为1，表示包年包月。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 购买集群数，可选值范围[1,50]，默认为1
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 回档类型：
	// noneRollback：不回档；
	// snapRollback，快照回档；
	// timeRollback，时间点回档
	RollbackStrategy *string `json:"RollbackStrategy,omitnil,omitempty" name:"RollbackStrategy"`

	// 快照回档，表示snapshotId；时间点回档，表示queryId，为0，表示需要判断时间点是否有效
	RollbackId *uint64 `json:"RollbackId,omitnil,omitempty" name:"RollbackId"`

	// 回档时，传入源集群ID，用于查找源poolId
	OriginalClusterId *string `json:"OriginalClusterId,omitnil,omitempty" name:"OriginalClusterId"`

	// 时间点回档，指定时间；快照回档，快照时间
	ExpectTime *string `json:"ExpectTime,omitnil,omitempty" name:"ExpectTime"`

	// 该参数无实际意义，已废弃。
	// 时间点回档，指定时间允许范围
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitnil,omitempty" name:"ExpectTimeThresh"`

	// 普通实例存储上限，单位GB
	// 当DbType为MYSQL，且存储计费模式为预付费时，该参数需不大于cpu与memory对应存储规格上限
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// 包年包月购买时长
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 包年包月购买时长单位，['s','d','m','y']
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 包年包月购买是否自动续费，默认为0。
	// 0标识默认续费方式，1表示自动续费，2表示不自动续费。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 实例数量（该参数已不再使用，只做存量兼容处理）
	HaCount *int64 `json:"HaCount,omitnil,omitempty" name:"HaCount"`

	// 订单来源
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// 集群创建需要绑定的tag数组信息
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Db类型
	// 当DbType为MYSQL时可选(默认NORMAL)：
	// <li>NORMAL</li>
	// <li>SERVERLESS</li>
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// 当DbMode为SERVERLESS时必填
	// cpu最小值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// 当DbMode为SERVERLESS时必填：
	// cpu最大值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// 当DbMode为SERVERLESS时，指定集群是否自动暂停，可选范围
	// <li>yes</li>
	// <li>no</li>
	// 默认值:yes
	AutoPause *string `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// 当DbMode为SERVERLESS时，指定集群自动暂停的延迟，单位秒，可选范围[600,691200]
	// 默认值:600
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitnil,omitempty" name:"AutoPauseDelay"`

	// 集群存储计费模式，按量计费：0，包年包月：1。默认按量计费
	// 当DbType为MYSQL时，在集群计算计费模式为后付费（包括DbMode为SERVERLESS）时，存储计费模式仅可为按量计费
	// 回档与克隆均不支持包年包月存储
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// 安全组id数组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 告警策略Id数组
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitnil,omitempty" name:"AlarmPolicyIds"`

	// 参数数组，暂时支持character_set_server （utf8｜latin1｜gbk｜utf8mb4） ，lower_case_table_names，1-大小写不敏感，0-大小写敏感
	ClusterParams []*ParamItem `json:"ClusterParams,omitnil,omitempty" name:"ClusterParams"`

	// 交易模式，0-下单且支付，1-下单
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// 参数模板ID，可以通过查询参数模板信息DescribeParamTemplates获得参数模板ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 多可用区地址
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// 实例初始化配置信息，主要用于购买集群时选不同规格实例
	InstanceInitInfos []*InstanceInitInfo `json:"InstanceInitInfos,omitnil,omitempty" name:"InstanceInitInfos"`

	// 全球数据库唯一标识
	GdnId *string `json:"GdnId,omitnil,omitempty" name:"GdnId"`

	// 数据库代理配置
	ProxyConfig *ProxyConfig `json:"ProxyConfig,omitnil,omitempty" name:"ProxyConfig"`

	// 是否自动归档
	AutoArchive *string `json:"AutoArchive,omitnil,omitempty" name:"AutoArchive"`

	// 暂停后的归档处理时间
	AutoArchiveDelayHours *int64 `json:"AutoArchiveDelayHours,omitnil,omitempty" name:"AutoArchiveDelayHours"`

	// 内核小版本号
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`
}

type CreateClustersRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 所属VPC网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 所属子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 数据库版本，取值范围: 
	// <li> MYSQL可选值：5.7，8.0 </li>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// 所属项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 当DbMode为NORMAL或不填时必选
	// 普通实例Cpu核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 当DbMode为NORMAL或不填时必选
	// 普通实例内存,单位GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例数量，数量范围为(0,16]，默认值为2（即一个rw实例+一个ro实例），传递的n表示1个rw实例+n-1个ro实例（规格相同），如需要更精确的集群组成搭配，请使用InstanceInitInfos
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 该参数无实际意义，已废弃。
	// 存储大小，单位GB。
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 集群名称，长度小于64个字符，每个字符取值范围：大/小写字母，数字，特殊符号（'-','_','.'）
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 账号密码(8-64个字符，包含大小写英文字母、数字和符号~!@#$%^&*_-+=`|\(){}[]:;'<>,.?/中的任意三种)
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// 端口，默认3306，取值范围[0, 65535)
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 计费模式，支持值为0和1，默认值为0。
	// 取值为0，表示按量计费。
	// 取值为1，表示包年包月。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 购买集群数，可选值范围[1,50]，默认为1
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 回档类型：
	// noneRollback：不回档；
	// snapRollback，快照回档；
	// timeRollback，时间点回档
	RollbackStrategy *string `json:"RollbackStrategy,omitnil,omitempty" name:"RollbackStrategy"`

	// 快照回档，表示snapshotId；时间点回档，表示queryId，为0，表示需要判断时间点是否有效
	RollbackId *uint64 `json:"RollbackId,omitnil,omitempty" name:"RollbackId"`

	// 回档时，传入源集群ID，用于查找源poolId
	OriginalClusterId *string `json:"OriginalClusterId,omitnil,omitempty" name:"OriginalClusterId"`

	// 时间点回档，指定时间；快照回档，快照时间
	ExpectTime *string `json:"ExpectTime,omitnil,omitempty" name:"ExpectTime"`

	// 该参数无实际意义，已废弃。
	// 时间点回档，指定时间允许范围
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitnil,omitempty" name:"ExpectTimeThresh"`

	// 普通实例存储上限，单位GB
	// 当DbType为MYSQL，且存储计费模式为预付费时，该参数需不大于cpu与memory对应存储规格上限
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// 包年包月购买时长
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 包年包月购买时长单位，['s','d','m','y']
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 包年包月购买是否自动续费，默认为0。
	// 0标识默认续费方式，1表示自动续费，2表示不自动续费。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 实例数量（该参数已不再使用，只做存量兼容处理）
	HaCount *int64 `json:"HaCount,omitnil,omitempty" name:"HaCount"`

	// 订单来源
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// 集群创建需要绑定的tag数组信息
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Db类型
	// 当DbType为MYSQL时可选(默认NORMAL)：
	// <li>NORMAL</li>
	// <li>SERVERLESS</li>
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// 当DbMode为SERVERLESS时必填
	// cpu最小值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// 当DbMode为SERVERLESS时必填：
	// cpu最大值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// 当DbMode为SERVERLESS时，指定集群是否自动暂停，可选范围
	// <li>yes</li>
	// <li>no</li>
	// 默认值:yes
	AutoPause *string `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// 当DbMode为SERVERLESS时，指定集群自动暂停的延迟，单位秒，可选范围[600,691200]
	// 默认值:600
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitnil,omitempty" name:"AutoPauseDelay"`

	// 集群存储计费模式，按量计费：0，包年包月：1。默认按量计费
	// 当DbType为MYSQL时，在集群计算计费模式为后付费（包括DbMode为SERVERLESS）时，存储计费模式仅可为按量计费
	// 回档与克隆均不支持包年包月存储
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// 安全组id数组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 告警策略Id数组
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitnil,omitempty" name:"AlarmPolicyIds"`

	// 参数数组，暂时支持character_set_server （utf8｜latin1｜gbk｜utf8mb4） ，lower_case_table_names，1-大小写不敏感，0-大小写敏感
	ClusterParams []*ParamItem `json:"ClusterParams,omitnil,omitempty" name:"ClusterParams"`

	// 交易模式，0-下单且支付，1-下单
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// 参数模板ID，可以通过查询参数模板信息DescribeParamTemplates获得参数模板ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 多可用区地址
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// 实例初始化配置信息，主要用于购买集群时选不同规格实例
	InstanceInitInfos []*InstanceInitInfo `json:"InstanceInitInfos,omitnil,omitempty" name:"InstanceInitInfos"`

	// 全球数据库唯一标识
	GdnId *string `json:"GdnId,omitnil,omitempty" name:"GdnId"`

	// 数据库代理配置
	ProxyConfig *ProxyConfig `json:"ProxyConfig,omitnil,omitempty" name:"ProxyConfig"`

	// 是否自动归档
	AutoArchive *string `json:"AutoArchive,omitnil,omitempty" name:"AutoArchive"`

	// 暂停后的归档处理时间
	AutoArchiveDelayHours *int64 `json:"AutoArchiveDelayHours,omitnil,omitempty" name:"AutoArchiveDelayHours"`

	// 内核小版本号
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`
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
	delete(f, "InstanceCount")
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
	delete(f, "GdnId")
	delete(f, "ProxyConfig")
	delete(f, "AutoArchive")
	delete(f, "AutoArchiveDelayHours")
	delete(f, "CynosVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClustersResponseParams struct {
	// 冻结流水ID
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// 订单号
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 资源ID列表（该字段已不再维护，请使用dealNames字段查询接口DescribeResourcesByDealName获取资源ID）
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 集群ID列表（该字段已不再维护，请使用dealNames字段查询接口DescribeResourcesByDealName获取集群ID）
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 大订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateIntegrateClusterRequestParams struct {
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 所属VPC网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 所属子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 数据库版本，取值范围: 
	// <li> MYSQL可选值：5.7，8.0 </li>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// 所属项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 集群名称，长度小于64个字符，每个字符取值范围：大/小写字母，数字，特殊符号（'-','_','.'）
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 账号密码(8-64个字符，包含大小写英文字母、数字和符号~!@#$%^&*_-+=`|\(){}[]:;'<>,.?/中的任意三种)
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// 端口，默认3306，取值范围[0, 65535)
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 计费模式，按量计费：0，包年包月：1。默认按量计费。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 购买集群数，可选值范围[1,3]，默认为1
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 普通实例存储上限，单位GB
	// 当DbType为MYSQL，且存储计费模式为预付费时，该参数需不大于cpu与memory对应存储规格上限
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// 包年包月购买时长
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 包年包月购买时长单位，['s','d','m','y']
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 包年包月购买是否自动续费，默认为0。
	// 0标识默认续费方式，1表示自动续费，2表示不自动续费。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 集群创建需要绑定的tag数组信息
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 集群存储计费模式，按量计费：0，包年包月：1。默认按量计费
	// 当DbType为MYSQL时，在集群计算计费模式为后付费（包括DbMode为SERVERLESS）时，存储计费模式仅可为按量计费
	// 回档与克隆均不支持包年包月存储
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// 安全组id数组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 告警策略Id数组
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitnil,omitempty" name:"AlarmPolicyIds"`

	// 参数数组，暂时支持character_set_server （utf8｜latin1｜gbk｜utf8mb4） ，lower_case_table_names，1-大小写不敏感，0-大小写敏感
	ClusterParams []*ParamItem `json:"ClusterParams,omitnil,omitempty" name:"ClusterParams"`

	// 交易模式，0-下单且支付，1-下单
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// 参数模板ID，可以通过查询参数模板信息DescribeParamTemplates获得参数模板ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 多可用区地址
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// 实例初始化配置信息，主要用于购买集群时选不同规格实例
	InstanceInitInfos []*IntegrateInstanceInfo `json:"InstanceInitInfos,omitnil,omitempty" name:"InstanceInitInfos"`

	// 全球数据库唯一标识
	GdnId *string `json:"GdnId,omitnil,omitempty" name:"GdnId"`

	// 数据库代理配置
	ProxyConfig *ProxyConfigInfo `json:"ProxyConfig,omitnil,omitempty" name:"ProxyConfig"`

	// 是否自动归档
	AutoArchive *string `json:"AutoArchive,omitnil,omitempty" name:"AutoArchive"`

	// 暂停后的归档处理时间
	AutoArchiveDelayHours *int64 `json:"AutoArchiveDelayHours,omitnil,omitempty" name:"AutoArchiveDelayHours"`

	// 加密方法（由加密算法和密钥对版本组成）
	EncryptMethod *string `json:"EncryptMethod,omitnil,omitempty" name:"EncryptMethod"`

	// 集成集群配置信息
	IntegrateCreateClusterConfig *IntegrateCreateClusterConfig `json:"IntegrateCreateClusterConfig,omitnil,omitempty" name:"IntegrateCreateClusterConfig"`

	// 存储架构类型。 枚举值：1.0/2.0 默认值：1.0
	StorageVersion *string `json:"StorageVersion,omitnil,omitempty" name:"StorageVersion"`
}

type CreateIntegrateClusterRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 所属VPC网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 所属子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 数据库版本，取值范围: 
	// <li> MYSQL可选值：5.7，8.0 </li>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// 所属项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 集群名称，长度小于64个字符，每个字符取值范围：大/小写字母，数字，特殊符号（'-','_','.'）
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 账号密码(8-64个字符，包含大小写英文字母、数字和符号~!@#$%^&*_-+=`|\(){}[]:;'<>,.?/中的任意三种)
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// 端口，默认3306，取值范围[0, 65535)
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 计费模式，按量计费：0，包年包月：1。默认按量计费。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 购买集群数，可选值范围[1,3]，默认为1
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 普通实例存储上限，单位GB
	// 当DbType为MYSQL，且存储计费模式为预付费时，该参数需不大于cpu与memory对应存储规格上限
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// 包年包月购买时长
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 包年包月购买时长单位，['s','d','m','y']
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 包年包月购买是否自动续费，默认为0。
	// 0标识默认续费方式，1表示自动续费，2表示不自动续费。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 集群创建需要绑定的tag数组信息
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 集群存储计费模式，按量计费：0，包年包月：1。默认按量计费
	// 当DbType为MYSQL时，在集群计算计费模式为后付费（包括DbMode为SERVERLESS）时，存储计费模式仅可为按量计费
	// 回档与克隆均不支持包年包月存储
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// 安全组id数组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 告警策略Id数组
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitnil,omitempty" name:"AlarmPolicyIds"`

	// 参数数组，暂时支持character_set_server （utf8｜latin1｜gbk｜utf8mb4） ，lower_case_table_names，1-大小写不敏感，0-大小写敏感
	ClusterParams []*ParamItem `json:"ClusterParams,omitnil,omitempty" name:"ClusterParams"`

	// 交易模式，0-下单且支付，1-下单
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// 参数模板ID，可以通过查询参数模板信息DescribeParamTemplates获得参数模板ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 多可用区地址
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// 实例初始化配置信息，主要用于购买集群时选不同规格实例
	InstanceInitInfos []*IntegrateInstanceInfo `json:"InstanceInitInfos,omitnil,omitempty" name:"InstanceInitInfos"`

	// 全球数据库唯一标识
	GdnId *string `json:"GdnId,omitnil,omitempty" name:"GdnId"`

	// 数据库代理配置
	ProxyConfig *ProxyConfigInfo `json:"ProxyConfig,omitnil,omitempty" name:"ProxyConfig"`

	// 是否自动归档
	AutoArchive *string `json:"AutoArchive,omitnil,omitempty" name:"AutoArchive"`

	// 暂停后的归档处理时间
	AutoArchiveDelayHours *int64 `json:"AutoArchiveDelayHours,omitnil,omitempty" name:"AutoArchiveDelayHours"`

	// 加密方法（由加密算法和密钥对版本组成）
	EncryptMethod *string `json:"EncryptMethod,omitnil,omitempty" name:"EncryptMethod"`

	// 集成集群配置信息
	IntegrateCreateClusterConfig *IntegrateCreateClusterConfig `json:"IntegrateCreateClusterConfig,omitnil,omitempty" name:"IntegrateCreateClusterConfig"`

	// 存储架构类型。 枚举值：1.0/2.0 默认值：1.0
	StorageVersion *string `json:"StorageVersion,omitnil,omitempty" name:"StorageVersion"`
}

func (r *CreateIntegrateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "DbVersion")
	delete(f, "ProjectId")
	delete(f, "ClusterName")
	delete(f, "AdminPassword")
	delete(f, "Port")
	delete(f, "PayMode")
	delete(f, "Count")
	delete(f, "StorageLimit")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "AutoRenewFlag")
	delete(f, "AutoVoucher")
	delete(f, "ResourceTags")
	delete(f, "StoragePayMode")
	delete(f, "SecurityGroupIds")
	delete(f, "AlarmPolicyIds")
	delete(f, "ClusterParams")
	delete(f, "DealMode")
	delete(f, "ParamTemplateId")
	delete(f, "SlaveZone")
	delete(f, "InstanceInitInfos")
	delete(f, "GdnId")
	delete(f, "ProxyConfig")
	delete(f, "AutoArchive")
	delete(f, "AutoArchiveDelayHours")
	delete(f, "EncryptMethod")
	delete(f, "IntegrateCreateClusterConfig")
	delete(f, "StorageVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIntegrateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrateClusterResponseParams struct {
	// 冻结流水ID
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// 订单号
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 资源ID列表（该字段已不再维护，请使用dealNames字段查询接口DescribeResourcesByDealName获取资源ID）
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 集群ID列表（该字段已不再维护，请使用dealNames字段查询接口DescribeResourcesByDealName获取集群ID）
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 大订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIntegrateClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateIntegrateClusterResponseParams `json:"Response"`
}

func (r *CreateIntegrateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParamTemplateRequestParams struct {
	// 模板名称
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// mysql版本号
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 模板描述
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// 可选参数，需要复制的模板ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 数据库类型，可选值：NORMAL（默认值），SERVERLESS
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// 参数列表
	ParamList []*ParamItem `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type CreateParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// mysql版本号
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 模板描述
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// 可选参数，需要复制的模板ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 数据库类型，可选值：NORMAL（默认值），SERVERLESS
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// 参数列表
	ParamList []*ParamItem `json:"ParamList,omitnil,omitempty" name:"ParamList"`
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
	// 模板ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 集群 ID。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 私有网络 ID，默认与集群私有网络 ID 保持一致。
	UniqueVpcId *string `json:"UniqueVpcId,omitnil,omitempty" name:"UniqueVpcId"`

	// 私有网络子网 ID，默认与集群子网 ID 保持一致。
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil,omitempty" name:"UniqueSubnetId"`

	// 连接池类型：SessionConnectionPool（会话级别连接池）。
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// 是否开启连接池。
	// yes：表示开启。
	// no：表示不开启。
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// 连接池阈值：单位（秒），可选范围：0 - 300秒。
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// 绑定的安全组 ID 数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 描述说明。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 想要绑定的 vip 信息，需与 UniqueVpcId 对应。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 权重模式：
	// system：系统分配。
	// custom：自定义。
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// 是否自动添加只读实例。
	// yes：表示自动添加只读实例。
	// no：表示不自动添加只读实例。
	AutoAddRo *string `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// 是否开启故障转移。
	// yes：表示开启，开启后，当数据库代理出现故障时，连接地址将会路由到主实例。
	// no：表示不开启。
	// 说明：
	// 仅当 RwType 参数值为 READWRITE 时，才支持设置此项。
	FailOver *string `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// 一致性类型：
	// eventual：最终一致性。
	// global：全局一致性。
	// session：会话一致性。
	// 说明：
	// 仅当 RwType 参数值为 READWRITE 时，才支持设置此项。
	ConsistencyType *string `json:"ConsistencyType,omitnil,omitempty" name:"ConsistencyType"`

	// 读写属性：
	// READWRITE：表示读写分离。当此参数值为 READWRITE 时，才支持设置 FailOver、ConsistencyType 参数。
	// READONLY：表示只读。
	RwType *string `json:"RwType,omitnil,omitempty" name:"RwType"`

	// 一致性超时时间。取值范围：0 ~ 1000000（微秒）。设置为0时，表示若只读实例出现延迟导致一致性策略不满足时，请求将一直等待。
	ConsistencyTimeOut *int64 `json:"ConsistencyTimeOut,omitnil,omitempty" name:"ConsistencyTimeOut"`

	// 是否开启事务拆分。开启后，在一个事务中拆分读和写到不同的实例上去执行。
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// 接入模式：
	// nearby：就近访问。
	// balance：均衡分配。
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// 实例权重。
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitnil,omitempty" name:"InstanceWeights"`
}

type CreateProxyEndPointRequest struct {
	*tchttp.BaseRequest
	
	// 集群 ID。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 私有网络 ID，默认与集群私有网络 ID 保持一致。
	UniqueVpcId *string `json:"UniqueVpcId,omitnil,omitempty" name:"UniqueVpcId"`

	// 私有网络子网 ID，默认与集群子网 ID 保持一致。
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil,omitempty" name:"UniqueSubnetId"`

	// 连接池类型：SessionConnectionPool（会话级别连接池）。
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// 是否开启连接池。
	// yes：表示开启。
	// no：表示不开启。
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// 连接池阈值：单位（秒），可选范围：0 - 300秒。
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// 绑定的安全组 ID 数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 描述说明。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 想要绑定的 vip 信息，需与 UniqueVpcId 对应。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 权重模式：
	// system：系统分配。
	// custom：自定义。
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// 是否自动添加只读实例。
	// yes：表示自动添加只读实例。
	// no：表示不自动添加只读实例。
	AutoAddRo *string `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// 是否开启故障转移。
	// yes：表示开启，开启后，当数据库代理出现故障时，连接地址将会路由到主实例。
	// no：表示不开启。
	// 说明：
	// 仅当 RwType 参数值为 READWRITE 时，才支持设置此项。
	FailOver *string `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// 一致性类型：
	// eventual：最终一致性。
	// global：全局一致性。
	// session：会话一致性。
	// 说明：
	// 仅当 RwType 参数值为 READWRITE 时，才支持设置此项。
	ConsistencyType *string `json:"ConsistencyType,omitnil,omitempty" name:"ConsistencyType"`

	// 读写属性：
	// READWRITE：表示读写分离。当此参数值为 READWRITE 时，才支持设置 FailOver、ConsistencyType 参数。
	// READONLY：表示只读。
	RwType *string `json:"RwType,omitnil,omitempty" name:"RwType"`

	// 一致性超时时间。取值范围：0 ~ 1000000（微秒）。设置为0时，表示若只读实例出现延迟导致一致性策略不满足时，请求将一直等待。
	ConsistencyTimeOut *int64 `json:"ConsistencyTimeOut,omitnil,omitempty" name:"ConsistencyTimeOut"`

	// 是否开启事务拆分。开启后，在一个事务中拆分读和写到不同的实例上去执行。
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// 接入模式：
	// nearby：就近访问。
	// balance：均衡分配。
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// 实例权重。
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitnil,omitempty" name:"InstanceWeights"`
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
	// 异步流程 ID。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 异步任务 ID。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 数据库代理组 ID。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// cpu核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 私有网络ID，默认与集群私有网络ID保持一致
	UniqueVpcId *string `json:"UniqueVpcId,omitnil,omitempty" name:"UniqueVpcId"`

	// 私有网络子网ID，默认与集群子网ID保持一致
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil,omitempty" name:"UniqueSubnetId"`

	// 数据库代理组节点个数（该参数不再建议使用，建议使用ProxyZones)
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// 连接池类型：SessionConnectionPool(会话级别连接池 )
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// 是否开启连接池,yes-开启，no-不开启
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// 连接池阈值：单位（秒）
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// 安全组ID数组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 描述说明
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据库节点信息（该参数与ProxyCount需要任选一个输入）
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`
}

type CreateProxyRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// cpu核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 私有网络ID，默认与集群私有网络ID保持一致
	UniqueVpcId *string `json:"UniqueVpcId,omitnil,omitempty" name:"UniqueVpcId"`

	// 私有网络子网ID，默认与集群子网ID保持一致
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil,omitempty" name:"UniqueSubnetId"`

	// 数据库代理组节点个数（该参数不再建议使用，建议使用ProxyZones)
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// 连接池类型：SessionConnectionPool(会话级别连接池 )
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// 是否开启连接池,yes-开启，no-不开启
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// 连接池阈值：单位（秒）
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// 安全组ID数组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 描述说明
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据库节点信息（该参数与ProxyCount需要任选一个输入）
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 异步任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例类型，目前固定传cynosdb-serverless
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 资源包使用地域chineseMainland-中国内地通用，overseas-港澳台及海外通用
	PackageRegion *string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// 资源包类型：CCU-计算资源包，DISK-存储资源包
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 资源包版本
	// base-基础版本，common-通用版本，enterprise-企业版本
	PackageVersion *string `json:"PackageVersion,omitnil,omitempty" name:"PackageVersion"`

	// 资源包大小，计算资源单位：个；存储资源：GB
	PackageSpec *float64 `json:"PackageSpec,omitnil,omitempty" name:"PackageSpec"`

	// 资源包有效期，单位:天
	ExpireDay *int64 `json:"ExpireDay,omitnil,omitempty" name:"ExpireDay"`

	// 购买资源包个数
	PackageCount *int64 `json:"PackageCount,omitnil,omitempty" name:"PackageCount"`

	// 资源包名称
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type CreateResourcePackageRequest struct {
	*tchttp.BaseRequest
	
	// 实例类型，目前固定传cynosdb-serverless
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 资源包使用地域chineseMainland-中国内地通用，overseas-港澳台及海外通用
	PackageRegion *string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// 资源包类型：CCU-计算资源包，DISK-存储资源包
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 资源包版本
	// base-基础版本，common-通用版本，enterprise-企业版本
	PackageVersion *string `json:"PackageVersion,omitnil,omitempty" name:"PackageVersion"`

	// 资源包大小，计算资源单位：个；存储资源：GB
	PackageSpec *float64 `json:"PackageSpec,omitnil,omitempty" name:"PackageSpec"`

	// 资源包有效期，单位:天
	ExpireDay *int64 `json:"ExpireDay,omitnil,omitempty" name:"ExpireDay"`

	// 购买资源包个数
	PackageCount *int64 `json:"PackageCount,omitnil,omitempty" name:"PackageCount"`

	// 资源包名称
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
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
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// 每个物品对应一个dealName，业务需要根据dealName保证发货接口幂等
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type CrossRegionBackupItem struct {
	// 备份的目标地域
	CrossRegion *string `json:"CrossRegion,omitnil,omitempty" name:"CrossRegion"`

	// 目标地域的备份任务ID
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 目标地域的备份状态
	BackupStatus *string `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`
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
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 数据库版本
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例数
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// 用户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 引擎类型
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 用户appid
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 集群状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 集群创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 付费模式。0-按量计费，1-包年包月
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 截止时间
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// 集群读写vip
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 集群读写vport
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 项目id
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// cynos内核版本
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// cynos版本标签
	CynosVersionTag *string `json:"CynosVersionTag,omitnil,omitempty" name:"CynosVersionTag"`

	// 存储容量
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// 续费标志
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 正在处理的任务
	ProcessingTask *string `json:"ProcessingTask,omitnil,omitempty" name:"ProcessingTask"`

	// 集群的任务数组
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 集群绑定的tag数组
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Db类型(NORMAL, SERVERLESS)
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// 当Db类型为SERVERLESS时，serverless集群状态，可选值:
	// resume
	// pause
	ServerlessStatus *string `json:"ServerlessStatus,omitnil,omitempty" name:"ServerlessStatus"`

	// 集群预付费存储值大小
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 集群存储为预付费时的存储ID，用于预付费存储变配
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// 集群存储付费模式。0-按量计费，1-包年包月
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// 集群计算规格对应的最小存储值
	MinStorageSize *int64 `json:"MinStorageSize,omitnil,omitempty" name:"MinStorageSize"`

	// 集群计算规格对应的最大存储值
	MaxStorageSize *int64 `json:"MaxStorageSize,omitnil,omitempty" name:"MaxStorageSize"`

	// 集群网络信息
	NetAddrs []*NetAddr `json:"NetAddrs,omitnil,omitempty" name:"NetAddrs"`

	// 物理可用区
	PhysicalZone *string `json:"PhysicalZone,omitnil,omitempty" name:"PhysicalZone"`

	// 主可用区
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// 是否有从可用区
	HasSlaveZone *string `json:"HasSlaveZone,omitnil,omitempty" name:"HasSlaveZone"`

	// 从可用区
	SlaveZones []*string `json:"SlaveZones,omitnil,omitempty" name:"SlaveZones"`

	// 商业类型
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// 是否冻结
	IsFreeze *string `json:"IsFreeze,omitnil,omitempty" name:"IsFreeze"`

	// 订单来源
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// 能力
	Ability *Ability `json:"Ability,omitnil,omitempty" name:"Ability"`

	// 实例绑定资源包信息（此处只返回存储资源包，即packageType=DISK）	
	ResourcePackages []*ResourcePackage `json:"ResourcePackages,omitnil,omitempty" name:"ResourcePackages"`

	// 全球数据库唯一标识
	GdnId *string `json:"GdnId,omitnil,omitempty" name:"GdnId"`

	// 集群角色。主集群- primary，从集群 - standby，如果 GdnId为空，该字段无效。
	GdnRole *string `json:"GdnRole,omitnil,omitempty" name:"GdnRole"`
}

type CynosdbClusterDetail struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 物理可用区
	PhysicalZone *string `json:"PhysicalZone,omitnil,omitempty" name:"PhysicalZone"`

	// 状态，支持的值如下：
	// - creating：创建中
	// - running：运行中
	// - isolating：隔离中
	// - isolated：已隔离
	// - activating：从回收站重新恢复
	// - offlining：下线中
	// - offlined：已下线
	// - deleting：删除中
	// - deleted：已删除
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 当Db类型为SERVERLESS时，serverless集群状态，可选值:
	// resume
	// resuming
	// pause
	// pausing
	ServerlessStatus *string `json:"ServerlessStatus,omitnil,omitempty" name:"ServerlessStatus"`

	// 存储Id
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// 存储大小，单位为G
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 最大存储规格，单位为G
	MaxStorageSize *int64 `json:"MaxStorageSize,omitnil,omitempty" name:"MaxStorageSize"`

	// 最小存储规格，单位为G
	MinStorageSize *int64 `json:"MinStorageSize,omitnil,omitempty" name:"MinStorageSize"`

	// 存储付费类型，1为包年包月，0为按量计费
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// VPC名称
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// vpc唯一id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网名称
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 字符集
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 数据库类型
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Db类型：<li>NORMAL</li><li>SERVERLESS</li>
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// 数据库版本
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// 存储空间上限
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// 使用容量
	UsedStorage *int64 `json:"UsedStorage,omitnil,omitempty" name:"UsedStorage"`

	// vip地址
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// vport端口
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 集群只读实例的vip地址和vport端口
	RoAddr []*Addr `json:"RoAddr,omitnil,omitempty" name:"RoAddr"`

	// 集群支持的功能
	Ability *Ability `json:"Ability,omitnil,omitempty" name:"Ability"`

	// cynos版本
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// 商业类型
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// 是否有从可用区
	HasSlaveZone *string `json:"HasSlaveZone,omitnil,omitempty" name:"HasSlaveZone"`

	// 是否冻结
	IsFreeze *string `json:"IsFreeze,omitnil,omitempty" name:"IsFreeze"`

	// 任务列表
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 主可用区
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// 从可用区列表
	SlaveZones []*string `json:"SlaveZones,omitnil,omitempty" name:"SlaveZones"`

	// 实例信息
	InstanceSet []*ClusterInstanceDetail `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// 付费模式
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 到期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// 项目id
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 实例绑定的tag数组信息
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Proxy状态
	ProxyStatus *string `json:"ProxyStatus,omitnil,omitempty" name:"ProxyStatus"`

	// binlog开关，可选值：ON, OFF
	LogBin *string `json:"LogBin,omitnil,omitempty" name:"LogBin"`

	// 是否跳过交易
	IsSkipTrade *string `json:"IsSkipTrade,omitnil,omitempty" name:"IsSkipTrade"`

	// pitr类型，可选值：normal, redo_pitr
	PitrType *string `json:"PitrType,omitnil,omitempty" name:"PitrType"`

	// 是否打开密码复杂度
	IsOpenPasswordComplexity *string `json:"IsOpenPasswordComplexity,omitnil,omitempty" name:"IsOpenPasswordComplexity"`

	// 网络类型
	NetworkStatus *string `json:"NetworkStatus,omitnil,omitempty" name:"NetworkStatus"`

	// 集群绑定的资源包信息	
	ResourcePackages []*ResourcePackage `json:"ResourcePackages,omitnil,omitempty" name:"ResourcePackages"`

	// 自动续费标识，1为自动续费，0为到期不续
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 节点网络类型
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// 备可用区属性
	SlaveZoneAttr []*SlaveZoneAttrItem `json:"SlaveZoneAttr,omitnil,omitempty" name:"SlaveZoneAttr"`

	// 版本标签
	CynosVersionTag *string `json:"CynosVersionTag,omitnil,omitempty" name:"CynosVersionTag"`

	// 全球数据库网络唯一标识
	GdnId *string `json:"GdnId,omitnil,omitempty" name:"GdnId"`

	// 集群在全球数据网络中的角色。
	// 主集群- primary
	// 从集群 - standby
	// 如为空，该字段无效
	GdnRole *string `json:"GdnRole,omitnil,omitempty" name:"GdnRole"`

	// 二级存储使用量，单位：G
	UsedArchiveStorage *int64 `json:"UsedArchiveStorage,omitnil,omitempty" name:"UsedArchiveStorage"`

	// 归档状态，枚举值<li>normal:正常</li><li>archiving:归档中</li><li>resuming:恢复中</li><li>archived :已归档</li>
	ArchiveStatus *string `json:"ArchiveStatus,omitnil,omitempty" name:"ArchiveStatus"`

	// 归档进度，百分比。
	ArchiveProgress *int64 `json:"ArchiveProgress,omitnil,omitempty" name:"ArchiveProgress"`
}

type CynosdbErrorLogItem struct {
	// 日志时间戳
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 日志内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type CynosdbInstance struct {
	// 用户Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 用户AppId
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例状态中文描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 实例形态，是否为serverless实例
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// 数据库类型
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 数据库版本
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// Cpu，单位：核
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存，单位：GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 存储量，单位：GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 实例类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例当前角色
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// VPC网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例内网IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 实例内网端口
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 付费模式
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 实例过期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// 销毁期限
	DestroyDeadlineText *string `json:"DestroyDeadlineText,omitnil,omitempty" name:"DestroyDeadlineText"`

	// 隔离时间
	IsolateTime *string `json:"IsolateTime,omitnil,omitempty" name:"IsolateTime"`

	// 网络类型
	NetType *int64 `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 外网域名
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// 外网IP
	WanIP *string `json:"WanIP,omitnil,omitempty" name:"WanIP"`

	// 外网端口
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// 外网状态
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// 实例销毁时间
	DestroyTime *string `json:"DestroyTime,omitnil,omitempty" name:"DestroyTime"`

	// Cynos内核版本
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// 正在处理的任务
	ProcessingTask *string `json:"ProcessingTask,omitnil,omitempty" name:"ProcessingTask"`

	// 续费标志
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// serverless实例cpu下限
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// serverless实例cpu上限
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// serverless实例状态, 可选值：
	// resume
	// pause
	ServerlessStatus *string `json:"ServerlessStatus,omitnil,omitempty" name:"ServerlessStatus"`

	// 预付费存储Id
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// 存储付费类型
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// 物理区
	PhysicalZone *string `json:"PhysicalZone,omitnil,omitempty" name:"PhysicalZone"`

	// 商业类型
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// 任务
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 是否冻结
	IsFreeze *string `json:"IsFreeze,omitnil,omitempty" name:"IsFreeze"`

	// 资源标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 主可用区
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// 备可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaveZones []*string `json:"SlaveZones,omitnil,omitempty" name:"SlaveZones"`

	// 实例网络信息
	InstanceNetInfo []*InstanceNetInfo `json:"InstanceNetInfo,omitnil,omitempty" name:"InstanceNetInfo"`

	// 实例绑定资源包信息（此处只返回计算资源包，即packageType=CCU）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourcePackages []*ResourcePackage `json:"ResourcePackages,omitnil,omitempty" name:"ResourcePackages"`

	// 实例索引形态,可选值【mixedRowColumn（行列混存），onlyRowIndex（仅行存）】
	InstanceIndexMode *string `json:"InstanceIndexMode,omitnil,omitempty" name:"InstanceIndexMode"`

	// 当前实例支持的能力
	InstanceAbility *InstanceAbility `json:"InstanceAbility,omitnil,omitempty" name:"InstanceAbility"`

	// 实例机器类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 实例存储类型
	InstanceStorageType *string `json:"InstanceStorageType,omitnil,omitempty" name:"InstanceStorageType"`

	// 未知字段
	CynosVersionTag *string `json:"CynosVersionTag,omitnil,omitempty" name:"CynosVersionTag"`

	// libradb 节点信息
	NodeList []*string `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// 全球数据库唯一标识
	GdnId *string `json:"GdnId,omitnil,omitempty" name:"GdnId"`
}

type CynosdbInstanceDetail struct {
	// 用户Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 用户AppId
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例状态
	// creating：创建中
	// running：运行中
	// isolating：隔离中
	// isolated：已隔离
	// activating：恢复中
	// offlining：下线中
	// offlined：已下线
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例状态中文描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// serverless实例状态, 可能值：
	// resume
	// pause
	ServerlessStatus *string `json:"ServerlessStatus,omitnil,omitempty" name:"ServerlessStatus"`

	// 数据库类型
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 数据库版本
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// Cpu，单位：核
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存，单位：GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 存储量，单位：GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 实例类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例当前角色
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 付费模式
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 实例过期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// 网络类型
	NetType *int64 `json:"NetType,omitnil,omitempty" name:"NetType"`

	// VPC网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例内网IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 实例内网端口
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 实例外网域名
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// 字符集
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// Cynos内核版本
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// 续费标志
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// serverless实例cpu下限
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// serverless实例cpu上限
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// Db类型:<li>NORMAL</li><li>SERVERLESS</li>
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`
}

type CynosdbInstanceGroup struct {
	// 用户appId
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 删除时间
	DeletedTime *string `json:"DeletedTime,omitnil,omitempty" name:"DeletedTime"`

	// 实例组ID
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例组（网络）类型。ha-ha组；ro-只读组；proxy-代理；singleRo-只读实例独占
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 更新时间
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// 内网IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 内网端口
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 外网域名
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// 外网ip
	WanIP *string `json:"WanIP,omitnil,omitempty" name:"WanIP"`

	// 外网端口
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// 外网状态
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// 实例组包含实例信息
	InstanceSet []*CynosdbInstance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// VPC的ID
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 子网ID
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 正在回收IP信息
	OldAddrInfo *OldAddrInfo `json:"OldAddrInfo,omitnil,omitempty" name:"OldAddrInfo"`

	// 正在进行的任务
	ProcessingTasks []*string `json:"ProcessingTasks,omitnil,omitempty" name:"ProcessingTasks"`

	// 任务列表
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// biz_net_service表id
	NetServiceId *int64 `json:"NetServiceId,omitnil,omitempty" name:"NetServiceId"`
}

type CynosdbInstanceGrp struct {
	// 用户appId
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 删除时间
	DeletedTime *string `json:"DeletedTime,omitnil,omitempty" name:"DeletedTime"`

	// 实例组ID
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例组类型。ha-ha组；ro-只读组
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 更新时间
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// 内网IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 内网端口
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 外网域名
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// 外网ip
	WanIP *string `json:"WanIP,omitnil,omitempty" name:"WanIP"`

	// 外网端口
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// 外网状态
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// 实例组包含实例信息
	InstanceSet []*CynosdbInstance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// VPC的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 正在回收IP信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldAddrInfo *OldAddrInfo `json:"OldAddrInfo,omitnil,omitempty" name:"OldAddrInfo"`

	// 正在进行的任务
	ProcessingTasks []*string `json:"ProcessingTasks,omitnil,omitempty" name:"ProcessingTasks"`

	// 任务列表
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// biz_net_service表id
	NetServiceId *int64 `json:"NetServiceId,omitnil,omitempty" name:"NetServiceId"`
}

type DatabasePrivileges struct {
	// 数据库
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// 权限列表
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

type DatabaseTables struct {
	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 表名称列表
	Tables []*string `json:"Tables,omitnil,omitempty" name:"Tables"`
}

type DbInfo struct {
	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 字符集类型
	CharacterSet *string `json:"CharacterSet,omitnil,omitempty" name:"CharacterSet"`

	// 数据库状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 排序规则
	CollateRule *string `json:"CollateRule,omitnil,omitempty" name:"CollateRule"`

	// 数据库备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户权限
	UserHostPrivileges []*UserHostPrivilege `json:"UserHostPrivileges,omitnil,omitempty" name:"UserHostPrivileges"`

	// 数据库ID
	DbId *int64 `json:"DbId,omitnil,omitempty" name:"DbId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 用户appid
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DbTable struct {
	// 数据库名称
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// 数据库表名称
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

// Predefined struct for user
type DeleteAccountsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 账号数组，包含account和host
	Accounts []*InputAccount `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

type DeleteAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 账号数组，包含account和host
	Accounts []*InputAccount `json:"Accounts,omitnil,omitempty" name:"Accounts"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计日志文件名称。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

type DeleteAuditLogFileRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计日志文件名称。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 审计规则模板ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
}

type DeleteAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 审计规则模板ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备份文件ID，旧版本使用的字段，不推荐使用
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitnil,omitempty" name:"SnapshotIdList"`

	// 备份文件ID，推荐使用
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`
}

type DeleteBackupRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备份文件ID，旧版本使用的字段，不推荐使用
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitnil,omitempty" name:"SnapshotIdList"`

	// 备份文件ID，推荐使用
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteCLSDeliveryRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志主题id
	CLSTopicIds []*string `json:"CLSTopicIds,omitnil,omitempty" name:"CLSTopicIds"`

	// 日志类型
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 是否维护时间运行
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type DeleteCLSDeliveryRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志主题id
	CLSTopicIds []*string `json:"CLSTopicIds,omitnil,omitempty" name:"CLSTopicIds"`

	// 日志类型
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 是否维护时间运行
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

func (r *DeleteCLSDeliveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCLSDeliveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CLSTopicIds")
	delete(f, "LogType")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCLSDeliveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCLSDeliveryResponseParams struct {
	// 异步任务id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCLSDeliveryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCLSDeliveryResponseParams `json:"Response"`
}

func (r *DeleteCLSDeliveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCLSDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterDatabaseRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库名
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`
}

type DeleteClusterDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库名
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 参数模板ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type DeliverSummary struct {
	// 投递类型，store（存储类），mq（消息通道）
	DeliverType *string `json:"DeliverType,omitnil,omitempty" name:"DeliverType"`

	// 投递子类型：cls，ckafka。
	DeliverSubType *string `json:"DeliverSubType,omitnil,omitempty" name:"DeliverSubType"`

	// 投递者
	DeliverConsumer *string `json:"DeliverConsumer,omitnil,omitempty" name:"DeliverConsumer"`

	// 投递者名称
	DeliverConsumerName *string `json:"DeliverConsumerName,omitnil,omitempty" name:"DeliverConsumerName"`
}

// Predefined struct for user
type DescribeAccountAllGrantPrivilegesRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`
}

type DescribeAccountAllGrantPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`
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
	PrivilegeStatements []*string `json:"PrivilegeStatements,omitnil,omitempty" name:"PrivilegeStatements"`

	// 全局权限
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// 数据库权限
	DatabasePrivileges []*DatabasePrivileges `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// 数据库表权限
	TablePrivileges []*TablePrivileges `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 账户名
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 主机
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 数据库名，为*时，忽略Type/TableName, 表示修改用户全局权限；
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// 指定数据库下的对象类型，可选"table"，"*"
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 当Type="table"时，用来指定表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 账户名
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 主机
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 数据库名，为*时，忽略Type/TableName, 表示修改用户全局权限；
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// 指定数据库下的对象类型，可选"table"，"*"
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 当Type="table"时，用来指定表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
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
	// 权限列表，示例值为：["","select","update","delete","create","drop","references","index","alter","show_db","create_tmp_table","lock_tables","execute","create_view","show_view","create_routine","alter_routine","event","trigger"]
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 需要过滤的账户列表
	AccountNames []*string `json:"AccountNames,omitnil,omitempty" name:"AccountNames"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	// 该参数已废用
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 需要过滤的账户列表
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// 限制量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 模糊匹配关键字(同时匹配AccountName和AccountHost，返回并集结果，支持正则)
	AccountRegular *string `json:"AccountRegular,omitnil,omitempty" name:"AccountRegular"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 需要过滤的账户列表
	AccountNames []*string `json:"AccountNames,omitnil,omitempty" name:"AccountNames"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	// 该参数已废用
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 需要过滤的账户列表
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// 限制量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 模糊匹配关键字(同时匹配AccountName和AccountHost，返回并集结果，支持正则)
	AccountRegular *string `json:"AccountRegular,omitnil,omitempty" name:"AccountRegular"`
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
	delete(f, "AccountRegular")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsResponseParams struct {
	// 数据库账号列表
	AccountSet []*Account `json:"AccountSet,omitnil,omitempty" name:"AccountSet"`

	// 账号总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAuditInstanceListRequestParams struct {
	// 实例审计开启的状态。1-已开启审计；0-未开启审计。
	AuditSwitch *int64 `json:"AuditSwitch,omitnil,omitempty" name:"AuditSwitch"`

	// 查询实例列表的过滤条件。
	Filters []*AuditInstanceFilters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 实例的审计规则模式。1-规则审计；0-全审计。
	AuditMode *int64 `json:"AuditMode,omitnil,omitempty" name:"AuditMode"`

	// 单次请求返回的数量。默认值为30，最大值为 100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认值为 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeAuditInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 实例审计开启的状态。1-已开启审计；0-未开启审计。
	AuditSwitch *int64 `json:"AuditSwitch,omitnil,omitempty" name:"AuditSwitch"`

	// 查询实例列表的过滤条件。
	Filters []*AuditInstanceFilters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 实例的审计规则模式。1-规则审计；0-全审计。
	AuditMode *int64 `json:"AuditMode,omitnil,omitempty" name:"AuditMode"`

	// 单次请求返回的数量。默认值为30，最大值为 100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认值为 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeAuditInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditSwitch")
	delete(f, "Filters")
	delete(f, "AuditMode")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditInstanceListResponseParams struct {
	// 符合查询条件的实例总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 审计实例详细信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*InstanceAuditStatus `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditInstanceListResponseParams `json:"Response"`
}

func (r *DescribeAuditInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogFilesRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页大小参数。默认值为 20，最小值为 1，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 审计日志文件名。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

type DescribeAuditLogFilesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页大小参数。默认值为 20，最小值为 1，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 审计日志文件名。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 审计日志文件详情。
	Items []*AuditLogFile `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，格式为："2017-07-12 10:29:20"。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，格式为："2017-07-12 10:29:20"。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序方式。支持值包括："ASC" - 升序，"DESC" - 降序。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段。支持值包括：
	// "timestamp" - 时间戳；
	// "affectRows" - 影响行数；
	// "execTime" - 执行时间。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 已废弃。
	//
	// Deprecated: Filter is deprecated.
	Filter *AuditLogFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 分页参数，单次返回的数据条数。默认值为100，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量。
	// 说明：Limit 和 Offset 的取值之和需小于等于65536。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件。多个值之前是且的关系。
	LogFilter []*InstanceAuditLogFilter `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`
}

type DescribeAuditLogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，格式为："2017-07-12 10:29:20"。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，格式为："2017-07-12 10:29:20"。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序方式。支持值包括："ASC" - 升序，"DESC" - 降序。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段。支持值包括：
	// "timestamp" - 时间戳；
	// "affectRows" - 影响行数；
	// "execTime" - 执行时间。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 已废弃。
	Filter *AuditLogFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 分页参数，单次返回的数据条数。默认值为100，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量。
	// 说明：Limit 和 Offset 的取值之和需小于等于65536。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件。多个值之前是且的关系。
	LogFilter []*InstanceAuditLogFilter `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`
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
	delete(f, "LogFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogsResponseParams struct {
	// 符合条件的审计日志条数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 审计日志详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*AuditLog `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 规则模板ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// 规则模板名称
	RuleTemplateNames []*string `json:"RuleTemplateNames,omitnil,omitempty" name:"RuleTemplateNames"`

	// 单次请求返回的数量。默认值20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认值为 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 告警等级。1-低风险，2-中风险，3-高风险。
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警策略。0-不告警，1-告警。
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`
}

type DescribeAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 规则模板ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// 规则模板名称
	RuleTemplateNames []*string `json:"RuleTemplateNames,omitnil,omitempty" name:"RuleTemplateNames"`

	// 单次请求返回的数量。默认值20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认值为 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 告警等级。1-低风险，2-中风险，3-高风险。
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警策略。0-不告警，1-告警。
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`
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
	delete(f, "AlarmLevel")
	delete(f, "AlarmPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditRuleTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRuleTemplatesResponseParams struct {
	// 符合查询条件的实例总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 规则模板详细信息列表。
	Items []*AuditRuleTemplateInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeAuditRuleWithInstanceIdsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。目前仅支持单个实例的查询。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例审计规则信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*InstanceAuditRule `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitnil,omitempty" name:"BackupTimeBeg"`

	// 表示全备开始时间，[0-24*3600]， 如0:00, 1:00, 2:00 分别为 0，3600， 7200
	BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitnil,omitempty" name:"BackupTimeEnd"`

	// 表示保留备份时长, 单位秒，超过该时间将被清理, 七天表示为3600*24*7=604800
	ReserveDuration *uint64 `json:"ReserveDuration,omitnil,omitempty" name:"ReserveDuration"`

	// 备份频率，长度为7的数组，分别对应周一到周日的备份方式，full-全量备份，increment-增量备份
	BackupFreq []*string `json:"BackupFreq,omitnil,omitempty" name:"BackupFreq"`

	// 备份方式，logic-逻辑备份，snapshot-快照备份
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 跨地域逻辑备份配置修改时间
	LogicCrossRegionsConfigUpdateTime *string `json:"LogicCrossRegionsConfigUpdateTime,omitnil,omitempty" name:"LogicCrossRegionsConfigUpdateTime"`

	// 自动逻辑备份配置
	LogicBackupConfig *LogicBackupConfigInfo `json:"LogicBackupConfig,omitnil,omitempty" name:"LogicBackupConfig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeBackupDownloadRestrictionRequestParams struct {
	// 集群ID
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
}

type DescribeBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
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
	delete(f, "ClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadRestrictionResponseParams struct {
	// 集群备份下载限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupLimitClusterRestrictions []*BackupLimitClusterRestriction `json:"BackupLimitClusterRestrictions,omitnil,omitempty" name:"BackupLimitClusterRestrictions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeBackupDownloadUrlRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备份ID
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 备份下载来源限制条件
	DownloadRestriction *BackupLimitRestriction `json:"DownloadRestriction,omitnil,omitempty" name:"DownloadRestriction"`
}

type DescribeBackupDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备份ID
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 备份下载来源限制条件
	DownloadRestriction *BackupLimitRestriction `json:"DownloadRestriction,omitnil,omitempty" name:"DownloadRestriction"`
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
	delete(f, "DownloadRestriction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadUrlResponseParams struct {
	// 备份下载地址
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeBackupDownloadUserRestrictionRequestParams struct {
	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 是否只查询用户级别下载限制，true-是，false-否
	OnlyUserRestriction *bool `json:"OnlyUserRestriction,omitnil,omitempty" name:"OnlyUserRestriction"`
}

type DescribeBackupDownloadUserRestrictionRequest struct {
	*tchttp.BaseRequest
	
	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 是否只查询用户级别下载限制，true-是，false-否
	OnlyUserRestriction *bool `json:"OnlyUserRestriction,omitnil,omitempty" name:"OnlyUserRestriction"`
}

func (r *DescribeBackupDownloadUserRestrictionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadUserRestrictionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OnlyUserRestriction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadUserRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadUserRestrictionResponseParams struct {
	// 集群备份下载限制信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupLimitClusterRestrictions []*BackupLimitClusterRestriction `json:"BackupLimitClusterRestrictions,omitnil,omitempty" name:"BackupLimitClusterRestrictions"`

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupDownloadUserRestrictionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDownloadUserRestrictionResponseParams `json:"Response"`
}

func (r *DescribeBackupDownloadUserRestrictionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadUserRestrictionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupListRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 返回数量，取值范围(0,100]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 记录偏移量，取值范围[0,INF)
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 备份ID
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`

	// 备份类型，可选值：snapshot，快照备份； logic，逻辑备份
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 备份方式，可选值：auto，自动备份；manual，手动备份
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 快照类型，可选值：full，全量；increment，增量
	SnapShotType *string `json:"SnapShotType,omitnil,omitempty" name:"SnapShotType"`

	// 备份开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 备份结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 备份文件名，模糊查询
	FileNames []*string `json:"FileNames,omitnil,omitempty" name:"FileNames"`

	// 备份备注名，模糊查询
	BackupNames []*string `json:"BackupNames,omitnil,omitempty" name:"BackupNames"`

	// 快照备份Id列表
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitnil,omitempty" name:"SnapshotIdList"`

	// 备份地域
	BackupRegion *string `json:"BackupRegion,omitnil,omitempty" name:"BackupRegion"`

	// 是否跨地域备份
	IsCrossRegionsBackup *string `json:"IsCrossRegionsBackup,omitnil,omitempty" name:"IsCrossRegionsBackup"`
}

type DescribeBackupListRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 返回数量，取值范围(0,100]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 记录偏移量，取值范围[0,INF)
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 备份ID
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`

	// 备份类型，可选值：snapshot，快照备份； logic，逻辑备份
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 备份方式，可选值：auto，自动备份；manual，手动备份
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 快照类型，可选值：full，全量；increment，增量
	SnapShotType *string `json:"SnapShotType,omitnil,omitempty" name:"SnapShotType"`

	// 备份开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 备份结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 备份文件名，模糊查询
	FileNames []*string `json:"FileNames,omitnil,omitempty" name:"FileNames"`

	// 备份备注名，模糊查询
	BackupNames []*string `json:"BackupNames,omitnil,omitempty" name:"BackupNames"`

	// 快照备份Id列表
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitnil,omitempty" name:"SnapshotIdList"`

	// 备份地域
	BackupRegion *string `json:"BackupRegion,omitnil,omitempty" name:"BackupRegion"`

	// 是否跨地域备份
	IsCrossRegionsBackup *string `json:"IsCrossRegionsBackup,omitnil,omitempty" name:"IsCrossRegionsBackup"`
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
	delete(f, "BackupRegion")
	delete(f, "IsCrossRegionsBackup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupListResponseParams struct {
	// 总共备份文件个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 备份文件列表
	BackupList []*BackupFileInfo `json:"BackupList,omitnil,omitempty" name:"BackupList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeBinlogConfigRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeBinlogConfigRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeBinlogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogConfigResponseParams struct {
	// Binlog跨地域配置更新时间
	BinlogCrossRegionsConfigUpdateTime *string `json:"BinlogCrossRegionsConfigUpdateTime,omitnil,omitempty" name:"BinlogCrossRegionsConfigUpdateTime"`

	// Binlog配置信息
	BinlogConfig *BinlogConfigInfo `json:"BinlogConfig,omitnil,omitempty" name:"BinlogConfig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBinlogConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBinlogConfigResponseParams `json:"Response"`
}

func (r *DescribeBinlogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogDownloadUrlRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Binlog文件ID
	BinlogId *int64 `json:"BinlogId,omitnil,omitempty" name:"BinlogId"`

	// 备份下载来源限制条件
	DownloadRestriction *BackupLimitRestriction `json:"DownloadRestriction,omitnil,omitempty" name:"DownloadRestriction"`
}

type DescribeBinlogDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Binlog文件ID
	BinlogId *int64 `json:"BinlogId,omitnil,omitempty" name:"BinlogId"`

	// 备份下载来源限制条件
	DownloadRestriction *BackupLimitRestriction `json:"DownloadRestriction,omitnil,omitempty" name:"DownloadRestriction"`
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
	delete(f, "DownloadRestriction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogDownloadUrlResponseParams struct {
	// 下载地址
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeBinlogSaveDaysRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	BinlogSaveDays *int64 `json:"BinlogSaveDays,omitnil,omitempty" name:"BinlogSaveDays"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制条数，默认值为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeBinlogsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制条数，默认值为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Binlog列表
	Binlogs []*BinlogItem `json:"Binlogs,omitnil,omitempty" name:"Binlogs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeChangedParamsAfterUpgradeRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 变配后的CPU
	DstCpu *int64 `json:"DstCpu,omitnil,omitempty" name:"DstCpu"`

	// 变配后的MEM，单位G
	DstMem *int64 `json:"DstMem,omitnil,omitempty" name:"DstMem"`
}

type DescribeChangedParamsAfterUpgradeRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 变配后的CPU
	DstCpu *int64 `json:"DstCpu,omitnil,omitempty" name:"DstCpu"`

	// 变配后的MEM，单位G
	DstMem *int64 `json:"DstMem,omitnil,omitempty" name:"DstMem"`
}

func (r *DescribeChangedParamsAfterUpgradeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChangedParamsAfterUpgradeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DstCpu")
	delete(f, "DstMem")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChangedParamsAfterUpgradeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChangedParamsAfterUpgradeResponseParams struct {
	// 参数个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例参数列表
	Items []*ParamItemInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeChangedParamsAfterUpgradeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChangedParamsAfterUpgradeResponseParams `json:"Response"`
}

func (r *DescribeChangedParamsAfterUpgradeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChangedParamsAfterUpgradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDatabaseTablesRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库名
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// 偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 个数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据表类型。
	// "view"表示只返回 view，"base_table" 表示只返回基本表，"all" 表示返回 view 和表。默认为 all。
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`
}

type DescribeClusterDatabaseTablesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库名
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// 偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 个数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据表类型。
	// "view"表示只返回 view，"base_table" 表示只返回基本表，"all" 表示返回 view 和表。默认为 all。
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`
}

func (r *DescribeClusterDatabaseTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDatabaseTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Db")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TableType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterDatabaseTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDatabaseTablesResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据库表列表
	Tables []*string `json:"Tables,omitnil,omitempty" name:"Tables"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterDatabaseTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterDatabaseTablesResponseParams `json:"Response"`
}

func (r *DescribeClusterDatabaseTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDatabaseTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDatabasesRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeClusterDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeClusterDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDatabasesResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数据库列表
	Databases []*string `json:"Databases,omitnil,omitempty" name:"Databases"`

	// 分页限制数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterDatabasesResponseParams `json:"Response"`
}

func (r *DescribeClusterDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDetailDatabasesRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20,最大100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`
}

type DescribeClusterDetailDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20,最大100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`
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
	DbInfos []*DbInfo `json:"DbInfos,omitnil,omitempty" name:"DbInfos"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	Detail *CynosdbClusterDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeClusterInstanceGroupsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterInstanceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterInstanceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInstanceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterInstanceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterInstanceGroupsResponseParams struct {
	// 实例组个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例组列表
	InstanceGroupInfoList []*CynosdbInstanceGroup `json:"InstanceGroupInfoList,omitnil,omitempty" name:"InstanceGroupInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterInstanceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterInstanceGroupsResponseParams `json:"Response"`
}

func (r *DescribeClusterInstanceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInstanceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterInstanceGrpsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterInstanceGrpsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例组列表
	//
	// Deprecated: InstanceGrpInfoList is deprecated.
	InstanceGrpInfoList []*CynosdbInstanceGrp `json:"InstanceGrpInfoList,omitnil,omitempty" name:"InstanceGrpInfoList"`

	// 实例组列表
	InstanceGroupInfoList []*CynosdbInstanceGroup `json:"InstanceGroupInfoList,omitnil,omitempty" name:"InstanceGroupInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID列表，用来记录具体操作哪些实例
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 排序字段，定义在回返结果的基于哪个字段进行排序
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 定义具体的排序规则，限定为desc,asc,DESC,ASC其中之一
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 返回数量，默认为 20，取值范围为(0,100]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0，取值范围为[0,INF)
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeClusterParamLogsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID列表，用来记录具体操作哪些实例
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 排序字段，定义在回返结果的基于哪个字段进行排序
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 定义具体的排序规则，限定为desc,asc,DESC,ASC其中之一
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 返回数量，默认为 20，取值范围为(0,100]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0，取值范围为[0,INF)
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 参数修改记录
	ClusterParamLogs []*ClusterParamModifyLog `json:"ClusterParamLogs,omitnil,omitempty" name:"ClusterParamLogs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 参数名字
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 是否为全局参数
	IsGlobal *string `json:"IsGlobal,omitnil,omitempty" name:"IsGlobal"`
}

type DescribeClusterParamsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 参数名字
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 是否为全局参数
	IsGlobal *string `json:"IsGlobal,omitnil,omitempty" name:"IsGlobal"`
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
	delete(f, "IsGlobal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterParamsResponseParams struct {
	// 参数个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例参数列表
	Items []*ParamInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	ValidatePasswordDictionary *ParamInfo `json:"ValidatePasswordDictionary,omitnil,omitempty" name:"ValidatePasswordDictionary"`

	// 密码长度
	ValidatePasswordLength *ParamInfo `json:"ValidatePasswordLength,omitnil,omitempty" name:"ValidatePasswordLength"`

	// 大小写敏感字符个数
	ValidatePasswordMixedCaseCount *ParamInfo `json:"ValidatePasswordMixedCaseCount,omitnil,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// 数字个数
	ValidatePasswordNumberCount *ParamInfo `json:"ValidatePasswordNumberCount,omitnil,omitempty" name:"ValidatePasswordNumberCount"`

	// 密码等级
	ValidatePasswordPolicy *ParamInfo `json:"ValidatePasswordPolicy,omitnil,omitempty" name:"ValidatePasswordPolicy"`

	// 特殊字符个数
	ValidatePasswordSpecialCharCount *ParamInfo `json:"ValidatePasswordSpecialCharCount,omitnil,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeClusterReadOnlyRequestParams struct {
	// 集群ID列表
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
}

type DescribeClusterReadOnlyRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID列表
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
}

func (r *DescribeClusterReadOnlyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterReadOnlyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterReadOnlyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterReadOnlyResponseParams struct {
	// 集群只读开关列表
	ClusterReadOnlyValues []*ClusterReadOnlyValue `json:"ClusterReadOnlyValues,omitnil,omitempty" name:"ClusterReadOnlyValues"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterReadOnlyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterReadOnlyResponseParams `json:"Response"`
}

func (r *DescribeClusterReadOnlyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterReadOnlyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterTransparentEncryptInfoRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterTransparentEncryptInfoRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterTransparentEncryptInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterTransparentEncryptInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterTransparentEncryptInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterTransparentEncryptInfoResponseParams struct {
	// 加密秘钥id
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 加密秘钥地域
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`

	// 秘钥类型
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`

	// 是否已经开启全局加密
	IsOpenGlobalEncryption *bool `json:"IsOpenGlobalEncryption,omitnil,omitempty" name:"IsOpenGlobalEncryption"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterTransparentEncryptInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterTransparentEncryptInfoResponseParams `json:"Response"`
}

func (r *DescribeClusterTransparentEncryptInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterTransparentEncryptInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersRequestParams struct {
	// 引擎类型：目前支持“MYSQL”， “POSTGRESQL”
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，取值范围：
	// <li> CREATETIME：创建时间</li>
	// <li> PERIODENDTIME：过期时间</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，取值范围：
	// <li> ASC：升序排序 </li>
	// <li> DESC：降序排序 </li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 搜索条件，若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
	
	// 引擎类型：目前支持“MYSQL”， “POSTGRESQL”
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，取值范围：
	// <li> CREATETIME：创建时间</li>
	// <li> PERIODENDTIME：过期时间</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，取值范围：
	// <li> ASC：升序排序 </li>
	// <li> DESC：降序排序 </li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 搜索条件，若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群列表
	ClusterSet []*CynosdbCluster `json:"ClusterSet,omitnil,omitempty" name:"ClusterSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例ID（InstanceId与InstanceGroupId必须任选一个传入）
	//
	// Deprecated: InstanceId is deprecated.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例组 ID，可通过 [DescribeClusterInstanceGroups](https://cloud.tencent.com/document/product/1003/103934) 接口查询。
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID（InstanceId与InstanceGroupId必须任选一个传入）
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例组 ID，可通过 [DescribeClusterInstanceGroups](https://cloud.tencent.com/document/product/1003/103934) 接口查询。
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`
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
	delete(f, "InstanceGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsResponseParams struct {
	// 安全组信息
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

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
type DescribeFlowRequestParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type DescribeFlowRequest struct {
	*tchttp.BaseRequest
	
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

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
type DescribeInstanceCLSLogDeliveryRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志类型
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`
}

type DescribeInstanceCLSLogDeliveryRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志类型
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`
}

func (r *DescribeInstanceCLSLogDeliveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceCLSLogDeliveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceCLSLogDeliveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceCLSLogDeliveryResponseParams struct {
	// 总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例投递信息
	InstanceCLSDeliveryInfos []*InstanceCLSDeliveryInfo `json:"InstanceCLSDeliveryInfos,omitnil,omitempty" name:"InstanceCLSDeliveryInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceCLSLogDeliveryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceCLSLogDeliveryResponseParams `json:"Response"`
}

func (r *DescribeInstanceCLSLogDeliveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceCLSLogDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDetailRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	Detail *CynosdbInstanceDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志条数限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 日志条数偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序字段，有Timestamp枚举值
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，有ASC,DESC枚举值
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 日志等级，有error、warning、note三种，支持多个等级同时搜索
	LogLevels []*string `json:"LogLevels,omitnil,omitempty" name:"LogLevels"`

	// 关键字，支持模糊搜索
	KeyWords []*string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`
}

type DescribeInstanceErrorLogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志条数限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 日志条数偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序字段，有Timestamp枚举值
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，有ASC,DESC枚举值
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 日志等级，有error、warning、note三种，支持多个等级同时搜索
	LogLevels []*string `json:"LogLevels,omitnil,omitempty" name:"LogLevels"`

	// 关键字，支持模糊搜索
	KeyWords []*string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 错误日志列表
	ErrorLogs []*CynosdbErrorLogItem `json:"ErrorLogs,omitnil,omitempty" name:"ErrorLogs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID，支持批量查询
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 参数名搜索条件，支持模糊匹配
	ParamKeyword *string `json:"ParamKeyword,omitnil,omitempty" name:"ParamKeyword"`

	// 是否为全局参数
	IsGlobal *string `json:"IsGlobal,omitnil,omitempty" name:"IsGlobal"`
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID，支持批量查询
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 参数名搜索条件，支持模糊匹配
	ParamKeyword *string `json:"ParamKeyword,omitnil,omitempty" name:"ParamKeyword"`

	// 是否为全局参数
	IsGlobal *string `json:"IsGlobal,omitnil,omitempty" name:"IsGlobal"`
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
	delete(f, "IsGlobal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsResponseParams struct {
	// 实例参数列表
	Items []*InstanceParamItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 事务开始最早时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 事务开始最晚时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 限制条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 客户端host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 排序字段，可选值：QueryTime,LockTime,RowsExamined,RowsSent
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，可选值：asc,desc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// sql语句
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`
}

type DescribeInstanceSlowQueriesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 事务开始最早时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 事务开始最晚时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 限制条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 客户端host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 排序字段，可选值：QueryTime,LockTime,RowsExamined,RowsSent
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，可选值：asc,desc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// sql语句
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`
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
	delete(f, "SqlText")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSlowQueriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSlowQueriesResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 慢查询记录
	SlowQueries []*SlowQueriesItem `json:"SlowQueries,omitnil,omitempty" name:"SlowQueries"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 是否需要返回可用区信息
	IncludeZoneStocks *bool `json:"IncludeZoneStocks,omitnil,omitempty" name:"IncludeZoneStocks"`

	// 实例机器类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

type DescribeInstanceSpecsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 是否需要返回可用区信息
	IncludeZoneStocks *bool `json:"IncludeZoneStocks,omitnil,omitempty" name:"IncludeZoneStocks"`

	// 实例机器类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
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
	delete(f, "DeviceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSpecsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSpecsResponseParams struct {
	// 规格信息
	InstanceSpecSet []*InstanceSpec `json:"InstanceSpecSet,omitnil,omitempty" name:"InstanceSpecSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 返回数量，默认为 20，取值范围为(0,100]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，取值范围：
	// <li> CREATETIME：创建时间</li>
	// <li> PERIODENDTIME：过期时间</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，取值范围：
	// <li> ASC：升序排序 </li>
	// <li> DESC：降序排序 </li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 搜索条件，若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 引擎类型：目前支持“MYSQL”
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 实例状态, 可选值:
	// creating 创建中
	// running 运行中
	// isolating 隔离中
	// isolated 已隔离
	// activating 恢复中
	// offlining 下线中
	// offlined 已下线
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例id列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 集群类型，取值范围<li> CYNOSDB：事务集群 </li><li> LIBRADB：分析集群 </li><li> ALL：全部 </li>，缺省为 ALL
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为 20，取值范围为(0,100]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，取值范围：
	// <li> CREATETIME：创建时间</li>
	// <li> PERIODENDTIME：过期时间</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，取值范围：
	// <li> ASC：升序排序 </li>
	// <li> DESC：降序排序 </li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 搜索条件，若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 引擎类型：目前支持“MYSQL”
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 实例状态, 可选值:
	// creating 创建中
	// running 运行中
	// isolating 隔离中
	// isolated 已隔离
	// activating 恢复中
	// offlining 下线中
	// offlined 已下线
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例id列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 集群类型，取值范围<li> CYNOSDB：事务集群 </li><li> LIBRADB：分析集群 </li><li> ALL：全部 </li>，缺省为 ALL
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
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
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 实例个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例列表
	InstanceSet []*CynosdbInstance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

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
type DescribeInstancesWithinSameClusterRequestParams struct {
	// vpcId
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// vip
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

type DescribeInstancesWithinSameClusterRequest struct {
	*tchttp.BaseRequest
	
	// vpcId
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// vip
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

func (r *DescribeInstancesWithinSameClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesWithinSameClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UniqVpcId")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesWithinSameClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesWithinSameClusterResponseParams struct {
	// 实例个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesWithinSameClusterResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesWithinSameClusterResponseParams `json:"Response"`
}

func (r *DescribeInstancesWithinSameClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesWithinSameClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrateTaskRequestParams struct {
	// 大订单id，大订单id和子订单id必须二选一
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`

	// 订单列表
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

type DescribeIntegrateTaskRequest struct {
	*tchttp.BaseRequest
	
	// 大订单id，大订单id和子订单id必须二选一
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`

	// 订单列表
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

func (r *DescribeIntegrateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BigDealId")
	delete(f, "DealNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrateTaskResponseParams struct {
	// 当前步骤
	CurrentStep *string `json:"CurrentStep,omitnil,omitempty" name:"CurrentStep"`

	// 当前进度
	CurrentProgress *string `json:"CurrentProgress,omitnil,omitempty" name:"CurrentProgress"`

	// 任务状态
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIntegrateTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrateTaskResponseParams `json:"Response"`
}

func (r *DescribeIntegrateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIsolatedInstancesRequestParams struct {
	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，取值范围：
	// <li> CREATETIME：创建时间</li>
	// <li> PERIODENDTIME：过期时间</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，取值范围：
	// <li> ASC：升序排序 </li>
	// <li> DESC：降序排序 </li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 搜索条件，若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 引擎类型：目前支持“MYSQL”， “POSTGRESQL”
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`
}

type DescribeIsolatedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，取值范围：
	// <li> CREATETIME：创建时间</li>
	// <li> PERIODENDTIME：过期时间</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，取值范围：
	// <li> ASC：升序排序 </li>
	// <li> DESC：降序排序 </li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 搜索条件，若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 引擎类型：目前支持“MYSQL”， “POSTGRESQL”
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`
}

func (r *DescribeIsolatedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIsolatedInstancesRequest) FromJsonString(s string) error {
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIsolatedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIsolatedInstancesResponseParams struct {
	// 实例个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例列表
	InstanceSet []*CynosdbInstance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIsolatedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIsolatedInstancesResponseParams `json:"Response"`
}

func (r *DescribeIsolatedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIsolatedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaintainPeriodRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeMaintainPeriodRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitnil,omitempty" name:"MaintainWeekDays"`

	// 维护开始时间，单位秒
	MaintainStartTime *int64 `json:"MaintainStartTime,omitnil,omitempty" name:"MaintainStartTime"`

	// 维护时长，单位秒
	MaintainDuration *int64 `json:"MaintainDuration,omitnil,omitempty" name:"MaintainDuration"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeParamTemplateDetailRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 参数模板名称
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 参数模板描述
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// 引擎版本
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 参数总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 参数列表
	Items []*ParamDetail `json:"Items,omitnil,omitempty" name:"Items"`

	// 数据库类型，可选值：NORMAL，SERVERLESS
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// 模板名称
	TemplateNames []*string `json:"TemplateNames,omitnil,omitempty" name:"TemplateNames"`

	// 模板ID
	TemplateIds []*int64 `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`

	// 数据库类型，可选值：NORMAL，SERVERLESS
	DbModes []*string `json:"DbModes,omitnil,omitempty" name:"DbModes"`

	// 查询偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询限制条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询的模板对应的产品类型
	Products []*string `json:"Products,omitnil,omitempty" name:"Products"`

	// 模板类型
	TemplateTypes []*string `json:"TemplateTypes,omitnil,omitempty" name:"TemplateTypes"`

	// 版本类型
	EngineTypes []*string `json:"EngineTypes,omitnil,omitempty" name:"EngineTypes"`

	// 返回结果的排序字段
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式（asc、desc）
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

type DescribeParamTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 数据库引擎版本号
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// 模板名称
	TemplateNames []*string `json:"TemplateNames,omitnil,omitempty" name:"TemplateNames"`

	// 模板ID
	TemplateIds []*int64 `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`

	// 数据库类型，可选值：NORMAL，SERVERLESS
	DbModes []*string `json:"DbModes,omitnil,omitempty" name:"DbModes"`

	// 查询偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询限制条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询的模板对应的产品类型
	Products []*string `json:"Products,omitnil,omitempty" name:"Products"`

	// 模板类型
	TemplateTypes []*string `json:"TemplateTypes,omitnil,omitempty" name:"TemplateTypes"`

	// 版本类型
	EngineTypes []*string `json:"EngineTypes,omitnil,omitempty" name:"EngineTypes"`

	// 返回结果的排序字段
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式（asc、desc）
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 参数模板信息
	Items []*ParamTemplateListInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 限制量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 搜索关键字
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 限制量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 搜索关键字
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
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
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 总数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 集群 ID（该参数必传，例如 cynosdbmysql-2u2mh111）。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，取值范围：
	// <li> CREATETIME：创建时间</li>
	// <li> PERIODENDTIME：过期时间</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，取值范围：
	// <li> ASC：升序排序 </li>
	// <li> DESC：降序排序 </li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 搜索条件，若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Filters []*QueryParamFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeProxiesRequest struct {
	*tchttp.BaseRequest
	
	// 集群 ID（该参数必传，例如 cynosdbmysql-2u2mh111）。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，取值范围：
	// <li> CREATETIME：创建时间</li>
	// <li> PERIODENDTIME：过期时间</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，取值范围：
	// <li> ASC：升序排序 </li>
	// <li> DESC：降序排序 </li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 搜索条件，若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Filters []*QueryParamFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 数据库代理组列表
	ProxyGroupInfos []*ProxyGroupInfo `json:"ProxyGroupInfos,omitnil,omitempty" name:"ProxyGroupInfos"`

	// 数据库代理节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyNodeInfos []*ProxyNodeInfo `json:"ProxyNodeInfos,omitnil,omitempty" name:"ProxyNodeInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，取值范围：
	// <li> CREATETIME：创建时间</li>
	// <li> PERIODENDTIME：过期时间</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，取值范围：
	// <li> ASC：升序排序 </li>
	// <li> DESC：降序排序 </li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 搜索条件，若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeProxyNodesRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 记录偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，取值范围：
	// <li> CREATETIME：创建时间</li>
	// <li> PERIODENDTIME：过期时间</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，取值范围：
	// <li> ASC：升序排序 </li>
	// <li> DESC：降序排序 </li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 搜索条件，若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 数据库代理节点列表
	ProxyNodeInfos []*ProxyNodeInfo `json:"ProxyNodeInfos,omitnil,omitempty" name:"ProxyNodeInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeProxySpecsRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxySpecsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySpecsResponseParams struct {
	// 数据库代理规格列表
	ProxySpecs []*ProxySpec `json:"ProxySpecs,omitnil,omitempty" name:"ProxySpecs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 集群ID
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 偏移量
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 实例D
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeResourcePackageDetailRequest struct {
	*tchttp.BaseRequest
	
	// 资源包唯一ID
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 集群ID
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 偏移量
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 实例D
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcePackageDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageDetailResponseParams struct {
	// 资源包抵扣总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 资源包明细说明
	Detail []*PackageDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	PackageId []*string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 资源包名称
	PackageName []*string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 资源包类型
	// CCU-计算资源包，DISK-存储资源包
	PackageType []*string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 资源包使用地域
	// china-中国内地通用，overseas-港澳台及海外通用
	PackageRegion []*string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// 资源包状态
	// creating-创建中；
	// using-使用中；
	// expired-已过期；
	// normal_finish-使用完；
	// apply_refund-申请退费中；
	// refund-已退费。
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// 排序条件，支持排序条件:startTime-生效时间，
	// expireTime-过期时间，packageUsedSpec-使用容量，packageTotalSpec-总存储量。
	// 按照数组顺序排列；
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式，DESC-降序，ASC-升序
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeResourcePackageListRequest struct {
	*tchttp.BaseRequest
	
	// 资源包唯一ID
	PackageId []*string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 资源包名称
	PackageName []*string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 资源包类型
	// CCU-计算资源包，DISK-存储资源包
	PackageType []*string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 资源包使用地域
	// china-中国内地通用，overseas-港澳台及海外通用
	PackageRegion []*string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// 资源包状态
	// creating-创建中；
	// using-使用中；
	// expired-已过期；
	// normal_finish-使用完；
	// apply_refund-申请退费中；
	// refund-已退费。
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// 排序条件，支持排序条件:startTime-生效时间，
	// expireTime-过期时间，packageUsedSpec-使用容量，packageTotalSpec-总存储量。
	// 按照数组顺序排列；
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式，DESC-降序，ASC-升序
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	// 资源包总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 资源包明细
	Detail []*Package `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 资源包使用地域
	// china-中国内地通用，overseas-港澳台及海外通用
	PackageRegion *string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// 资源包类型
	// CCU-计算资源包
	// DISK-存储资源包
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeResourcePackageSaleSpecRequest struct {
	*tchttp.BaseRequest
	
	// 实例类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 资源包使用地域
	// china-中国内地通用，overseas-港澳台及海外通用
	PackageRegion *string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// 资源包类型
	// CCU-计算资源包
	// DISK-存储资源包
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 资源包明细说明
	Detail []*SalePackageSpec `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 计费订单ID（如果计费还没回调业务发货，可能出现错误码InvalidParameterValue.DealNameNotFound，这种情况需要业务重试DescribeResourcesByDealName接口直到成功）。
	// DealName与DealNames至少应输入一项，两者都传时以DealName为准。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 计费订单ID列表，可以一次查询若干条订单ID对应资源信息（如果计费还没回调业务发货，可能出现错误码InvalidParameterValue.DealNameNotFound，这种情况需要业务重试DescribeResourcesByDealName接口直到成功）。
	// DealName与DealNames至少应输入一项，两者都传时以DealName为准。
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

type DescribeResourcesByDealNameRequest struct {
	*tchttp.BaseRequest
	
	// 计费订单ID（如果计费还没回调业务发货，可能出现错误码InvalidParameterValue.DealNameNotFound，这种情况需要业务重试DescribeResourcesByDealName接口直到成功）。
	// DealName与DealNames至少应输入一项，两者都传时以DealName为准。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 计费订单ID列表，可以一次查询若干条订单ID对应资源信息（如果计费还没回调业务发货，可能出现错误码InvalidParameterValue.DealNameNotFound，这种情况需要业务重试DescribeResourcesByDealName接口直到成功）。
	// DealName与DealNames至少应输入一项，两者都传时以DealName为准。
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
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
	BillingResourceInfos []*BillingResourceInfo `json:"BillingResourceInfos,omitnil,omitempty" name:"BillingResourceInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeRollbackTimeRangeRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	TimeRangeStart *string `json:"TimeRangeStart,omitnil,omitempty" name:"TimeRangeStart"`

	// 有效回归时间范围结束时间点（已废弃）
	TimeRangeEnd *string `json:"TimeRangeEnd,omitnil,omitempty" name:"TimeRangeEnd"`

	// 可回档时间范围
	RollbackTimeRanges []*RollbackTimeRange `json:"RollbackTimeRanges,omitnil,omitempty" name:"RollbackTimeRanges"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeSSLStatusRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeSSLStatusRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeSSLStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSSLStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSSLStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSSLStatusResponseParams struct {
	// yes-开启，no-关闭
	IsOpenSSL *string `json:"IsOpenSSL,omitnil,omitempty" name:"IsOpenSSL"`

	// 证书下载地址
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSSLStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSSLStatusResponseParams `json:"Response"`
}

func (r *DescribeSSLStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSSLStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessInstanceSpecsRequestParams struct {
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type DescribeServerlessInstanceSpecsRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

func (r *DescribeServerlessInstanceSpecsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessInstanceSpecsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServerlessInstanceSpecsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessInstanceSpecsResponseParams struct {
	// Serverless实例可选规格
	Specs []*ServerlessSpec `json:"Specs,omitnil,omitempty" name:"Specs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServerlessInstanceSpecsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServerlessInstanceSpecsResponseParams `json:"Response"`
}

func (r *DescribeServerlessInstanceSpecsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessInstanceSpecsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessStrategyRequestParams struct {
	// serverless集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeServerlessStrategyRequest struct {
	*tchttp.BaseRequest
	
	// serverless集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeServerlessStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServerlessStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessStrategyResponseParams struct {
	// cpu负载为 0 时持续多久（秒）发起自动暂停
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitnil,omitempty" name:"AutoPauseDelay"`

	// cpu负载超过当前规格核数时，持续多久（秒）发起自动扩容
	AutoScaleUpDelay *int64 `json:"AutoScaleUpDelay,omitnil,omitempty" name:"AutoScaleUpDelay"`

	// cpu 负载低于低一级规格核数时，持续多久（秒）发起自动缩容
	AutoScaleDownDelay *int64 `json:"AutoScaleDownDelay,omitnil,omitempty" name:"AutoScaleDownDelay"`

	// 是否自动暂停，可能值：
	// yes
	// no
	AutoPause *string `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// 集群是否允许向上扩容，可选范围<li>yes</li><li>no</li>
	AutoScaleUp *string `json:"AutoScaleUp,omitnil,omitempty" name:"AutoScaleUp"`

	// 集群是否允许向下缩容，可选范围<li>yes</li><li>no</li>
	AutoScaleDown *string `json:"AutoScaleDown,omitnil,omitempty" name:"AutoScaleDown"`

	// 是否开启归档，可选范围<li>yes</li><li>no</li>默认值:yes
	AutoArchive *string `json:"AutoArchive,omitnil,omitempty" name:"AutoArchive"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServerlessStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServerlessStrategyResponseParams `json:"Response"`
}

func (r *DescribeServerlessStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlaveZonesRequestParams struct {
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 云架集群ID
	OssClusterId *int64 `json:"OssClusterId,omitnil,omitempty" name:"OssClusterId"`
}

type DescribeSlaveZonesRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 云架集群ID
	OssClusterId *int64 `json:"OssClusterId,omitnil,omitempty" name:"OssClusterId"`
}

func (r *DescribeSlaveZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlaveZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "OssClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlaveZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlaveZonesResponseParams struct {
	// 从可用区
	SlaveZones []*string `json:"SlaveZones,omitnil,omitempty" name:"SlaveZones"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlaveZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlaveZonesResponseParams `json:"Response"`
}

func (r *DescribeSlaveZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlaveZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportProxyVersionRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

type DescribeSupportProxyVersionRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
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
	SupportProxyVersions []*string `json:"SupportProxyVersions,omitnil,omitempty" name:"SupportProxyVersions"`

	// 当前proxy版本号
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitnil,omitempty" name:"CurrentProxyVersion"`

	// 代理版本详情
	SupportProxyVersionDetail []*ProxyVersionInfo `json:"SupportProxyVersionDetail,omitnil,omitempty" name:"SupportProxyVersionDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeTasksRequestParams struct {
	// 任务开始时间起始值
	StartTimeBegin *string `json:"StartTimeBegin,omitnil,omitempty" name:"StartTimeBegin"`

	// 任务开始时间结束值
	StartTimeEnd *string `json:"StartTimeEnd,omitnil,omitempty" name:"StartTimeEnd"`

	// 过滤条件，支持的搜索字段："ClusterId"、"ClusterName"、"InstanceId"、"InstanceName"、"Status"、"TaskId"、"TaskType"
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询列表长度
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询列表偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务开始时间起始值
	StartTimeBegin *string `json:"StartTimeBegin,omitnil,omitempty" name:"StartTimeBegin"`

	// 任务开始时间结束值
	StartTimeEnd *string `json:"StartTimeEnd,omitnil,omitempty" name:"StartTimeEnd"`

	// 过滤条件，支持的搜索字段："ClusterId"、"ClusterName"、"InstanceId"、"InstanceName"、"Status"、"TaskId"、"TaskType"
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询列表长度
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询列表偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	delete(f, "StartTimeBegin")
	delete(f, "StartTimeEnd")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// 任务列表总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务列表
	TaskList []*BizTaskInfo `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeZonesRequestParams struct {
	// 是否包含虚拟区
	IncludeVirtualZones *bool `json:"IncludeVirtualZones,omitnil,omitempty" name:"IncludeVirtualZones"`

	// 是否展示地域下所有可用区，并显示用户每个可用区权限
	ShowPermission *bool `json:"ShowPermission,omitnil,omitempty" name:"ShowPermission"`
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
	// 是否包含虚拟区
	IncludeVirtualZones *bool `json:"IncludeVirtualZones,omitnil,omitempty" name:"IncludeVirtualZones"`

	// 是否展示地域下所有可用区，并显示用户每个可用区权限
	ShowPermission *bool `json:"ShowPermission,omitnil,omitempty" name:"ShowPermission"`
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
	RegionSet []*SaleRegion `json:"RegionSet,omitnil,omitempty" name:"RegionSet"`

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

// Predefined struct for user
type DisassociateSecurityGroupsRequestParams struct {
	// 实例组 ID 数组，cynosdbmysql-grp-前缀开头或集群 ID。
	// 说明：要获取集群的实例组 ID，可通过 [查询集群实例组](https://cloud.tencent.com/document/product/1003/103934) 进行。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 要修改的安全组ID列表，一个或者多个安全组ID组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 可用区。
	// 说明：请正确输入集群所在的主可用区，若输入非集群所在的主可用区可能显示调用成功，但实际执行会失败。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例组 ID 数组，cynosdbmysql-grp-前缀开头或集群 ID。
	// 说明：要获取集群的实例组 ID，可通过 [查询集群实例组](https://cloud.tencent.com/document/product/1003/103934) 进行。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 要修改的安全组ID列表，一个或者多个安全组ID组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 可用区。
	// 说明：请正确输入集群所在的主可用区，若输入非集群所在的主可用区可能显示调用成功，但实际执行会失败。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 日志等级，可选值note, warning，error
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 日志内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ExchangeInstanceInfo struct {
	// 源实例信息
	SrcInstanceInfo *RollbackInstanceInfo `json:"SrcInstanceInfo,omitnil,omitempty" name:"SrcInstanceInfo"`

	// 目标实例信息
	DstInstanceInfo *RollbackInstanceInfo `json:"DstInstanceInfo,omitnil,omitempty" name:"DstInstanceInfo"`
}

type ExchangeRoGroupInfo struct {
	// 源RO组信息
	SrcRoGroupInfo *RollbackRoGroupInfo `json:"SrcRoGroupInfo,omitnil,omitempty" name:"SrcRoGroupInfo"`

	// 目标RO组信息
	DstRoGroupInfo *RollbackRoGroupInfo `json:"DstRoGroupInfo,omitnil,omitempty" name:"DstRoGroupInfo"`
}

// Predefined struct for user
type ExportInstanceErrorLogsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志最早时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 日志最晚时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 限制条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 日志等级
	LogLevels []*string `json:"LogLevels,omitnil,omitempty" name:"LogLevels"`

	// 关键字
	KeyWords []*string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`

	// 文件类型，可选值：csv, original
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 可选值Timestamp
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，ASC 或 DESC。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type ExportInstanceErrorLogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志最早时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 日志最晚时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 限制条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 日志等级
	LogLevels []*string `json:"LogLevels,omitnil,omitempty" name:"LogLevels"`

	// 关键字
	KeyWords []*string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`

	// 文件类型，可选值：csv, original
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 可选值Timestamp
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，ASC 或 DESC。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
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
	ErrorLogItems []*ErrorLogItemExport `json:"ErrorLogItems,omitnil,omitempty" name:"ErrorLogItems"`

	// 错误日志字符串
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 事务开始最早时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 事务开始最晚时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 限制条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 客户端host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 文件类型，可选值：csv, original
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 排序字段，可选值： QueryTime,LockTime,RowsExamined,RowsSent
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，可选值：asc,desc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type ExportInstanceSlowQueriesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 事务开始最早时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 事务开始最晚时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 限制条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 客户端host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 文件类型，可选值：csv, original
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 排序字段，可选值： QueryTime,LockTime,RowsExamined,RowsSent
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，可选值：asc,desc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
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
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportInstanceSlowQueriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportInstanceSlowQueriesResponseParams struct {
	// 慢查询导出内容
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ExportResourcePackageDeductDetailsRequestParams struct {
	// 需要导出的资源包ID
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 使用资源包容量的cynos集群ID
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 排序字段，目前支持：createTime（资源包被抵扣时间），successDeductSpec（资源包抵扣量）
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，支持ASC、DESC、asc、desc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 单次最大导出数据行数，目前最大支持2000行
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量页数
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 导出数据格式，目前仅支持csv格式，留作扩展
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`
}

type ExportResourcePackageDeductDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 需要导出的资源包ID
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 使用资源包容量的cynos集群ID
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 排序字段，目前支持：createTime（资源包被抵扣时间），successDeductSpec（资源包抵扣量）
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，支持ASC、DESC、asc、desc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 单次最大导出数据行数，目前最大支持2000行
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量页数
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 导出数据格式，目前仅支持csv格式，留作扩展
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`
}

func (r *ExportResourcePackageDeductDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportResourcePackageDeductDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageId")
	delete(f, "ClusterIds")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "FileType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportResourcePackageDeductDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportResourcePackageDeductDetailsResponseParams struct {
	// 文件详情
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportResourcePackageDeductDetailsResponse struct {
	*tchttp.BaseResponse
	Response *ExportResourcePackageDeductDetailsResponseParams `json:"Response"`
}

func (r *ExportResourcePackageDeductDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportResourcePackageDeductDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GdnTaskInfo struct {
	// 全球数据库唯一标识
	GdnId *string `json:"GdnId,omitnil,omitempty" name:"GdnId"`

	// 全球数据库唯一别名
	GdnName *string `json:"GdnName,omitnil,omitempty" name:"GdnName"`

	// 主集群ID
	PrimaryClusterId *string `json:"PrimaryClusterId,omitnil,omitempty" name:"PrimaryClusterId"`

	// 主集群所在地域
	PrimaryClusterRegion *string `json:"PrimaryClusterRegion,omitnil,omitempty" name:"PrimaryClusterRegion"`

	// 从集群所在地域
	StandbyClusterRegion *string `json:"StandbyClusterRegion,omitnil,omitempty" name:"StandbyClusterRegion"`

	// 从集群ID
	StandbyClusterId *string `json:"StandbyClusterId,omitnil,omitempty" name:"StandbyClusterId"`

	// 从集群名称
	StandbyClusterName *string `json:"StandbyClusterName,omitnil,omitempty" name:"StandbyClusterName"`

	// 是否已强切
	ForceSwitchGdn *string `json:"ForceSwitchGdn,omitnil,omitempty" name:"ForceSwitchGdn"`

	// 返回码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 提示信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 是否支持强切
	IsSupportForce *string `json:"IsSupportForce,omitnil,omitempty" name:"IsSupportForce"`
}

type GoodsPrice struct {
	// 实例价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstancePrice *TradePrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// 存储价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoragePrice *TradePrice `json:"StoragePrice,omitnil,omitempty" name:"StoragePrice"`

	// 商品规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsSpec *GoodsSpec `json:"GoodsSpec,omitnil,omitempty" name:"GoodsSpec"`
}

type GoodsSpec struct {
	// 商品数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// CPU核数，PREPAID与POSTPAID实例类型必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存大小，单位G，PREPAID与POSTPAID实例类型必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Ccu大小，serverless类型必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ccu *float64 `json:"Ccu,omitnil,omitempty" name:"Ccu"`

	// 存储大小，PREPAID存储类型必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// 购买时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 时长单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 机器类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

// Predefined struct for user
type GrantAccountPrivilegesRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`

	// 数据库表权限码数组
	DbTablePrivileges []*string `json:"DbTablePrivileges,omitnil,omitempty" name:"DbTablePrivileges"`

	// 数据库表信息
	DbTables []*DbTable `json:"DbTables,omitnil,omitempty" name:"DbTables"`
}

type GrantAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`

	// 数据库表权限码数组
	DbTablePrivileges []*string `json:"DbTablePrivileges,omitnil,omitempty" name:"DbTablePrivileges"`

	// 数据库表信息
	DbTables []*DbTable `json:"DbTables,omitnil,omitempty" name:"DbTables"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 主机，默认‘%’
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

// Predefined struct for user
type InquirePriceCreateRequestParams struct {
	// 可用区,每个地域提供最佳实践
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 购买计算节点个数
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 实例购买类型，可选值为：PREPAID, POSTPAID, SERVERLESS
	InstancePayMode *string `json:"InstancePayMode,omitnil,omitempty" name:"InstancePayMode"`

	// 存储购买类型，可选值为：PREPAID, POSTPAID
	StoragePayMode *string `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// 实例设备类型，支持值如下：
	// - common：表示通用型
	// - exclusive：表示独享型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// CPU核数，PREPAID与POSTPAID实例类型必传
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存大小，单位G，PREPAID与POSTPAID实例类型必传
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Ccu大小，serverless类型必传
	Ccu *float64 `json:"Ccu,omitnil,omitempty" name:"Ccu"`

	// 存储大小，PREPAID存储类型必传
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// 购买时长，PREPAID购买类型必传
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 时长单位，可选值为：m,d。PREPAID购买类型必传
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
}

type InquirePriceCreateRequest struct {
	*tchttp.BaseRequest
	
	// 可用区,每个地域提供最佳实践
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 购买计算节点个数
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 实例购买类型，可选值为：PREPAID, POSTPAID, SERVERLESS
	InstancePayMode *string `json:"InstancePayMode,omitnil,omitempty" name:"InstancePayMode"`

	// 存储购买类型，可选值为：PREPAID, POSTPAID
	StoragePayMode *string `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// 实例设备类型，支持值如下：
	// - common：表示通用型
	// - exclusive：表示独享型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// CPU核数，PREPAID与POSTPAID实例类型必传
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存大小，单位G，PREPAID与POSTPAID实例类型必传
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Ccu大小，serverless类型必传
	Ccu *float64 `json:"Ccu,omitnil,omitempty" name:"Ccu"`

	// 存储大小，PREPAID存储类型必传
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// 购买时长，PREPAID购买类型必传
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 时长单位，可选值为：m,d。PREPAID购买类型必传
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
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
	delete(f, "DeviceType")
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
	InstancePrice *TradePrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// 存储价格
	StoragePrice *TradePrice `json:"StoragePrice,omitnil,omitempty" name:"StoragePrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type InquirePriceModifyRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// CPU核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存大小
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 存储大小，存储资源变配：ClusterId,StorageLimit
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// 实例ID，计算资源变配必传: ClusterId,InstanceId,Cpu,Memory 
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例设备类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// serverless实例ccu大小
	Ccu *float64 `json:"Ccu,omitnil,omitempty" name:"Ccu"`
}

type InquirePriceModifyRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// CPU核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存大小
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 存储大小，存储资源变配：ClusterId,StorageLimit
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// 实例ID，计算资源变配必传: ClusterId,InstanceId,Cpu,Memory 
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例设备类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// serverless实例ccu大小
	Ccu *float64 `json:"Ccu,omitnil,omitempty" name:"Ccu"`
}

func (r *InquirePriceModifyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "StorageLimit")
	delete(f, "InstanceId")
	delete(f, "DeviceType")
	delete(f, "Ccu")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceModifyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceModifyResponseParams struct {
	// 实例价格
	InstancePrice *TradePrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// 存储价格
	StoragePrice *TradePrice `json:"StoragePrice,omitnil,omitempty" name:"StoragePrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceModifyResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceModifyResponseParams `json:"Response"`
}

func (r *InquirePriceModifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceMultiSpecRequestParams struct {
	// 可用区,每个地域提供最佳实践
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例购买类型，可选值为：PREPAID, POSTPAID, SERVERLESS
	InstancePayMode *string `json:"InstancePayMode,omitnil,omitempty" name:"InstancePayMode"`

	// 存储购买类型，可选值为：PREPAID, POSTPAID
	StoragePayMode *string `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// 商品规格
	GoodsSpecs []*GoodsSpec `json:"GoodsSpecs,omitnil,omitempty" name:"GoodsSpecs"`
}

type InquirePriceMultiSpecRequest struct {
	*tchttp.BaseRequest
	
	// 可用区,每个地域提供最佳实践
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例购买类型，可选值为：PREPAID, POSTPAID, SERVERLESS
	InstancePayMode *string `json:"InstancePayMode,omitnil,omitempty" name:"InstancePayMode"`

	// 存储购买类型，可选值为：PREPAID, POSTPAID
	StoragePayMode *string `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// 商品规格
	GoodsSpecs []*GoodsSpec `json:"GoodsSpecs,omitnil,omitempty" name:"GoodsSpecs"`
}

func (r *InquirePriceMultiSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceMultiSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "InstancePayMode")
	delete(f, "StoragePayMode")
	delete(f, "GoodsSpecs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceMultiSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceMultiSpecResponseParams struct {
	// 商品价格
	GoodsPrice []*GoodsPrice `json:"GoodsPrice,omitnil,omitempty" name:"GoodsPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceMultiSpecResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceMultiSpecResponseParams `json:"Response"`
}

func (r *InquirePriceMultiSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceMultiSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 购买时长,与TimeUnit组合才能生效
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 购买时长单位, 与TimeSpan组合生效，可选:日:d,月:m
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
}

type InquirePriceRenewRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 购买时长,与TimeUnit组合才能生效
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 购买时长单位, 与TimeSpan组合生效，可选:日:d,月:m
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 对应的询价结果数组
	Prices []*TradePrice `json:"Prices,omitnil,omitempty" name:"Prices"`

	// 续费计算节点的总价格
	InstanceRealTotalPrice *int64 `json:"InstanceRealTotalPrice,omitnil,omitempty" name:"InstanceRealTotalPrice"`

	// 续费存储节点的总价格
	StorageRealTotalPrice *int64 `json:"StorageRealTotalPrice,omitnil,omitempty" name:"StorageRealTotalPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type InstanceAbility struct {
	// 实例是否支持强制重启，可选值：yes：支持，no：不支持
	IsSupportForceRestart *string `json:"IsSupportForceRestart,omitnil,omitempty" name:"IsSupportForceRestart"`

	// 不支持强制重启的原因
	NonsupportForceRestartReason *string `json:"NonsupportForceRestartReason,omitnil,omitempty" name:"NonsupportForceRestartReason"`
}

type InstanceAuditLogFilter struct {
	// 过滤项。目前支持以下搜索条件：
	// 
	// 包含、不包含、包含（分词维度）、不包含（分词维度）: sql - SQL详情；alarmLevel - 告警等级；ruleTemplateId - 规则模板Id
	// 
	// 等于、不等于、包含、不包含： host - 客户端地址； user - 用户名； dbName - 数据库名称；
	// 
	// 等于、不等于： sqlType - SQL类型； errCode - 错误码； threadId - 线程ID；
	// 
	// 范围搜索（时间类型统一为微秒）： execTime - 执行时间； lockWaitTime - 执行时间； ioWaitTime - IO等待时间； trxLivingTime - 事务持续时间； cpuTime - cpu时间； checkRows - 扫描行数； affectRows - 影响行数； sentRows - 返回行数。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 过滤条件。支持以下条件：
	// WINC-包含（分词维度），
	// WEXC-不包含（分词维度）,
	// INC - 包含,
	// EXC - 不包含,
	// EQS - 等于,
	// NEQ - 不等于,
	// RA - 范围。
	Compare *string `json:"Compare,omitnil,omitempty" name:"Compare"`

	// 过滤的值。反向查询时，多个值之前是且的关系，正向查询多个值是或的关系。
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type InstanceAuditRule struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否是规则审计。true-规则审计，false-全审计。
	AuditRule *bool `json:"AuditRule,omitnil,omitempty" name:"AuditRule"`

	// 审计规则详情。仅当AuditRule=true时有效。
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// 是否是审计策略
	OldRule *bool `json:"OldRule,omitnil,omitempty" name:"OldRule"`

	// 实例应用的规则模板详情
	RuleTemplates []*RuleTemplateInfo `json:"RuleTemplates,omitnil,omitempty" name:"RuleTemplates"`
}

type InstanceAuditStatus struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计状态。ON-表示审计已开启，OFF-表示审计关闭。
	AuditStatus *string `json:"AuditStatus,omitnil,omitempty" name:"AuditStatus"`

	// 日志保留时长。
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 高频存储时长。
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// 低频存储时长。
	LowLogExpireDay *uint64 `json:"LowLogExpireDay,omitnil,omitempty" name:"LowLogExpireDay"`

	// 日志存储量。
	BillingAmount *float64 `json:"BillingAmount,omitnil,omitempty" name:"BillingAmount"`

	// 高频存储量。
	HighRealStorage *float64 `json:"HighRealStorage,omitnil,omitempty" name:"HighRealStorage"`

	// 低频存储量。
	LowRealStorage *float64 `json:"LowRealStorage,omitnil,omitempty" name:"LowRealStorage"`

	// 是否为全审计。true-表示全审计。
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// 审计开通时间。
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// 实例相关信息。
	InstanceInfo *AuditInstanceInfo `json:"InstanceInfo,omitnil,omitempty" name:"InstanceInfo"`

	// 总存储量。
	RealStorage *float64 `json:"RealStorage,omitnil,omitempty" name:"RealStorage"`

	// 实例所应用的规则模板。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// 是否开启日志投递：ON，OFF
	Deliver *string `json:"Deliver,omitnil,omitempty" name:"Deliver"`

	// 日志投递类型
	DeliverSummary []*DeliverSummary `json:"DeliverSummary,omitnil,omitempty" name:"DeliverSummary"`
}

type InstanceCLSDeliveryInfo struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 日志主题id
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志主题name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 日志集id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 日志集name
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 日志投递地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 投递状态creating,running,offlining,offlined
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 日志类型
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`
}

type InstanceInitInfo struct {
	// 实例cpu
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例内存
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例类型 rw/ro
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例个数,范围[1,15]
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// Serverless实例个数最小值，范围[1,15]
	MinRoCount *int64 `json:"MinRoCount,omitnil,omitempty" name:"MinRoCount"`

	// Serverless实例个数最大值，范围[1,15]
	MaxRoCount *int64 `json:"MaxRoCount,omitnil,omitempty" name:"MaxRoCount"`

	// Serverless实例最小规格
	MinRoCpu *float64 `json:"MinRoCpu,omitnil,omitempty" name:"MinRoCpu"`

	// Serverless实例最大规格
	MaxRoCpu *float64 `json:"MaxRoCpu,omitnil,omitempty" name:"MaxRoCpu"`

	// 实例机器类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

type InstanceNameWeight struct {
	// 实例名称，创建集群中InstanceInitInfo.InstanceName所指定名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 权重
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type InstanceNetInfo struct {
	// 网络类型
	InstanceGroupType *string `json:"InstanceGroupType,omitnil,omitempty" name:"InstanceGroupType"`

	// 实例组ID
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 网络类型, 0-基础网络, 1-vpc网络, 2-黑石网络
	NetType *int64 `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 私有网络IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 私有网络端口
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 外网域名
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// 外网IP
	WanIP *string `json:"WanIP,omitnil,omitempty" name:"WanIP"`

	// 外网端口
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// 外网开启状态
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`
}

type InstanceParamItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例参数列表
	ParamsItems []*ParamItemDetail `json:"ParamsItems,omitnil,omitempty" name:"ParamsItems"`
}

type InstanceSpec struct {
	// 实例CPU，单位：核
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例内存，单位：GB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例最大可用存储，单位：GB
	MaxStorageSize *uint64 `json:"MaxStorageSize,omitnil,omitempty" name:"MaxStorageSize"`

	// 实例最小可用存储，单位：GB
	MinStorageSize *uint64 `json:"MinStorageSize,omitnil,omitempty" name:"MinStorageSize"`

	// 是否有库存
	HasStock *bool `json:"HasStock,omitnil,omitempty" name:"HasStock"`

	// 机器类型
	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`

	// 最大IOPS
	MaxIops *int64 `json:"MaxIops,omitnil,omitempty" name:"MaxIops"`

	// 最大IO带宽
	MaxIoBandWidth *int64 `json:"MaxIoBandWidth,omitnil,omitempty" name:"MaxIoBandWidth"`

	// 地域库存信息
	ZoneStockInfos []*ZoneStockInfo `json:"ZoneStockInfos,omitnil,omitempty" name:"ZoneStockInfos"`

	// 库存数量
	StockCount *int64 `json:"StockCount,omitnil,omitempty" name:"StockCount"`

	// 最大cpu
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// 最小cpu
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`
}

type IntegrateCreateClusterConfig struct {
	// binlog保留天数[7,1830]
	BinlogSaveDays *int64 `json:"BinlogSaveDays,omitnil,omitempty" name:"BinlogSaveDays"`

	// 备份保留天数[7,1830]
	BackupSaveDays *int64 `json:"BackupSaveDays,omitnil,omitempty" name:"BackupSaveDays"`

	// 半同步超时时间[1000,4294967295]
	SemiSyncTimeout *int64 `json:"SemiSyncTimeout,omitnil,omitempty" name:"SemiSyncTimeout"`

	// proxy连接地址配置信息
	ProxyEndPointConfigs []*ProxyEndPointConfigInfo `json:"ProxyEndPointConfigs,omitnil,omitempty" name:"ProxyEndPointConfigs"`
}

type IntegrateInstanceInfo struct {
	// 实例cpu
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例内存
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例类型 rw/ro
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例个数,范围[1,15]
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 实例机器类型 common-公通用型,exclusive-独享型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

// Predefined struct for user
type IsolateClusterRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 该参数已废用
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 实例退还原因类型
	IsolateReasonTypes []*int64 `json:"IsolateReasonTypes,omitnil,omitempty" name:"IsolateReasonTypes"`

	// 实例退还原因补充
	IsolateReason *string `json:"IsolateReason,omitnil,omitempty" name:"IsolateReason"`
}

type IsolateClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 该参数已废用
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 实例退还原因类型
	IsolateReasonTypes []*int64 `json:"IsolateReasonTypes,omitnil,omitempty" name:"IsolateReasonTypes"`

	// 实例退还原因补充
	IsolateReason *string `json:"IsolateReason,omitnil,omitempty" name:"IsolateReason"`
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
	delete(f, "IsolateReasonTypes")
	delete(f, "IsolateReason")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateClusterResponseParams struct {
	// 任务流ID(后付费或者serverless资源返回，如果需要同步任务状态，请使用DescribeFlow接口)
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 退款订单号(预付费资源返回，如果需要同步订单状态，请使用计费产品的DescribeDealsByCond同步订单状态)
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID数组，例如["cynosdbbmysql-ins-asd","cynosdbmysql-ins-zxc"]
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// 该参数已废弃
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 实例退还原因类型
	IsolateReasonTypes []*int64 `json:"IsolateReasonTypes,omitnil,omitempty" name:"IsolateReasonTypes"`

	// 实例退还原因补充
	IsolateReason *string `json:"IsolateReason,omitnil,omitempty" name:"IsolateReason"`
}

type IsolateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID数组，例如["cynosdbbmysql-ins-asd","cynosdbmysql-ins-zxc"]
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// 该参数已废弃
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 实例退还原因类型
	IsolateReasonTypes []*int64 `json:"IsolateReasonTypes,omitnil,omitempty" name:"IsolateReasonTypes"`

	// 实例退还原因补充
	IsolateReason *string `json:"IsolateReason,omitnil,omitempty" name:"IsolateReason"`
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
	delete(f, "IsolateReasonTypes")
	delete(f, "IsolateReason")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateInstanceResponseParams struct {
	// 任务流id
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 隔离实例的订单id（预付费实例）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type LogRuleTemplateInfo struct {
	// 模板ID
	RuleTemplateId *string `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// 规则模板名
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// 告警等级。1-低风险，2-中风险，3-高风险。
	AlarmLevel *string `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 规则模板变更状态：0-未变更；1-已变更；2-已删除
	RuleTemplateStatus *int64 `json:"RuleTemplateStatus,omitnil,omitempty" name:"RuleTemplateStatus"`
}

type LogicBackupConfigInfo struct {
	// 是否开启自动逻辑备份
	LogicBackupEnable *string `json:"LogicBackupEnable,omitnil,omitempty" name:"LogicBackupEnable"`

	// 自动逻辑备份开始时间
	LogicBackupTimeBeg *uint64 `json:"LogicBackupTimeBeg,omitnil,omitempty" name:"LogicBackupTimeBeg"`

	// 自动逻辑备份结束时间
	LogicBackupTimeEnd *uint64 `json:"LogicBackupTimeEnd,omitnil,omitempty" name:"LogicBackupTimeEnd"`

	// 自动逻辑备份保留时间
	// 单位：秒
	LogicReserveDuration *uint64 `json:"LogicReserveDuration,omitnil,omitempty" name:"LogicReserveDuration"`

	// 是否开启跨地域逻辑备份
	// 可选值：ON/OFF
	LogicCrossRegionsEnable *string `json:"LogicCrossRegionsEnable,omitnil,omitempty" name:"LogicCrossRegionsEnable"`

	// 逻辑备份所跨地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogicCrossRegions []*string `json:"LogicCrossRegions,omitnil,omitempty" name:"LogicCrossRegions"`
}

type ManualBackupData struct {
	// 备份类型。snapshot-快照备份
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 备份方式。auto-自动备份，manual-手动
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 备份时间
	SnapshotTime *string `json:"SnapshotTime,omitnil,omitempty" name:"SnapshotTime"`

	// 跨地域备份项详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrossRegionBackupInfos []*CrossRegionBackupItem `json:"CrossRegionBackupInfos,omitnil,omitempty" name:"CrossRegionBackupInfos"`
}

type ModifiableInfo struct {
	// 参数是否可被修改, 1:可以 0:不可以
	IsModifiable *int64 `json:"IsModifiable,omitnil,omitempty" name:"IsModifiable"`
}

// Predefined struct for user
type ModifyAccountDescriptionRequestParams struct {
	// 数据库账号名
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 数据库账号描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 主机，默认为"%"
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// 数据库账号名
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 数据库账号描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 主机，默认为"%"
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 新主机
	NewHost *string `json:"NewHost,omitnil,omitempty" name:"NewHost"`

	// 账号信息
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`
}

type ModifyAccountHostRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 新主机
	NewHost *string `json:"NewHost,omitnil,omitempty" name:"NewHost"`

	// 账号信息
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`

	// 数据库表权限数组,当前仅支持参数：max_user_connections，max_user_connections不能大于10240
	AccountParams []*AccountParam `json:"AccountParams,omitnil,omitempty" name:"AccountParams"`
}

type ModifyAccountParamsRequest struct {
	*tchttp.BaseRequest
	
	// 集群id，不超过32个字符
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`

	// 数据库表权限数组,当前仅支持参数：max_user_connections，max_user_connections不能大于10240
	AccountParams []*AccountParam `json:"AccountParams,omitnil,omitempty" name:"AccountParams"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`

	// 全局权限数组
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// 数据库权限数组
	DatabasePrivileges []*DatabasePrivileges `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// 表权限数组
	TablePrivileges []*TablePrivileges `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`
}

type ModifyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`

	// 全局权限数组
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// 数据库权限数组
	DatabasePrivileges []*DatabasePrivileges `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// 表权限数组
	TablePrivileges []*TablePrivileges `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 审计规则模板ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// 修改后的审计规则。
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// 修改后的规则模板名称。
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// 修改后的规则模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 告警等级。1-低风险，2-中风险，3-高风险。
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警策略。0-不告警，1-告警。
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`
}

type ModifyAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 审计规则模板ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// 修改后的审计规则。
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// 修改后的规则模板名称。
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// 修改后的规则模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 告警等级。1-低风险，2-中风险，3-高风险。
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警策略。0-不告警，1-告警。
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`
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
	delete(f, "AlarmLevel")
	delete(f, "AlarmPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuditRuleTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditRuleTemplatesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志保留时长。
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 高频日志保留时长。
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// 修改实例审计规则为全审计。
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// 规则审计。
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// 规则模板ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
}

type ModifyAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志保留时长。
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 高频日志保留时长。
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// 修改实例审计规则为全审计。
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// 规则审计。
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// 规则模板ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 表示全备开始时间，[0-24*3600]， 如0:00, 1:00, 2:00 分别为 0，3600， 7200
	BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitnil,omitempty" name:"BackupTimeBeg"`

	// 表示全备结束时间，[0-24*3600]， 如0:00, 1:00, 2:00 分别为 0，3600， 7200
	BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitnil,omitempty" name:"BackupTimeEnd"`

	// 表示保留备份时长, 单位秒，超过该时间将被清理, 七天表示为3600*24*7=604800，最大为158112000
	ReserveDuration *uint64 `json:"ReserveDuration,omitnil,omitempty" name:"ReserveDuration"`

	// 该参数目前不支持修改，无需填写。备份频率，长度为7的数组，分别对应周一到周日的备份方式，full-全量备份，increment-增量备份
	BackupFreq []*string `json:"BackupFreq,omitnil,omitempty" name:"BackupFreq"`

	// 该参数目前不支持修改，无需填写。
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 逻辑备份配置
	LogicBackupConfig *LogicBackupConfigInfo `json:"LogicBackupConfig,omitnil,omitempty" name:"LogicBackupConfig"`

	// 是否删除自动逻辑备份
	DeleteAutoLogicBackup *bool `json:"DeleteAutoLogicBackup,omitnil,omitempty" name:"DeleteAutoLogicBackup"`
}

type ModifyBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 表示全备开始时间，[0-24*3600]， 如0:00, 1:00, 2:00 分别为 0，3600， 7200
	BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitnil,omitempty" name:"BackupTimeBeg"`

	// 表示全备结束时间，[0-24*3600]， 如0:00, 1:00, 2:00 分别为 0，3600， 7200
	BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitnil,omitempty" name:"BackupTimeEnd"`

	// 表示保留备份时长, 单位秒，超过该时间将被清理, 七天表示为3600*24*7=604800，最大为158112000
	ReserveDuration *uint64 `json:"ReserveDuration,omitnil,omitempty" name:"ReserveDuration"`

	// 该参数目前不支持修改，无需填写。备份频率，长度为7的数组，分别对应周一到周日的备份方式，full-全量备份，increment-增量备份
	BackupFreq []*string `json:"BackupFreq,omitnil,omitempty" name:"BackupFreq"`

	// 该参数目前不支持修改，无需填写。
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 逻辑备份配置
	LogicBackupConfig *LogicBackupConfigInfo `json:"LogicBackupConfig,omitnil,omitempty" name:"LogicBackupConfig"`

	// 是否删除自动逻辑备份
	DeleteAutoLogicBackup *bool `json:"DeleteAutoLogicBackup,omitnil,omitempty" name:"DeleteAutoLogicBackup"`
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
	delete(f, "BackupTimeBeg")
	delete(f, "BackupTimeEnd")
	delete(f, "ReserveDuration")
	delete(f, "BackupFreq")
	delete(f, "BackupType")
	delete(f, "LogicBackupConfig")
	delete(f, "DeleteAutoLogicBackup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 集群ID
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 下载限制类型，NoLimit-不限制,LimitOnlyIntranet-限制内网 ,Customize-自定义
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// 该参数仅支持 In， 表示 LimitVpc 指定的vpc可以下载。默认为In
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// In: 指定的ip可以下载； NotIn: 指定的ip不可以下载
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// 限制下载的vpc设置
	LimitVpcs []*BackupLimitVpcItem `json:"LimitVpcs,omitnil,omitempty" name:"LimitVpcs"`

	// 限制下载的ip设置
	LimitIps []*string `json:"LimitIps,omitnil,omitempty" name:"LimitIps"`
}

type ModifyBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 下载限制类型，NoLimit-不限制,LimitOnlyIntranet-限制内网 ,Customize-自定义
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// 该参数仅支持 In， 表示 LimitVpc 指定的vpc可以下载。默认为In
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// In: 指定的ip可以下载； NotIn: 指定的ip不可以下载
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// 限制下载的vpc设置
	LimitVpcs []*BackupLimitVpcItem `json:"LimitVpcs,omitnil,omitempty" name:"LimitVpcs"`

	// 限制下载的ip设置
	LimitIps []*string `json:"LimitIps,omitnil,omitempty" name:"LimitIps"`
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
	delete(f, "ClusterIds")
	delete(f, "LimitType")
	delete(f, "VpcComparisonSymbol")
	delete(f, "IpComparisonSymbol")
	delete(f, "LimitVpcs")
	delete(f, "LimitIps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupDownloadRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadRestrictionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyBackupDownloadUserRestrictionRequestParams struct {
	// 下载限制类型，NoLimit-不限制,LimitOnlyIntranet-限制内网 ,Customize-自定义
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// 该参数仅支持 In， 表示 LimitVpc 指定的vpc可以下载。默认为In
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// In: 指定的ip可以下载； NotIn: 指定的ip不可以下载
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// 限制下载的vpc设置
	LimitVpcs []*BackupLimitVpcItem `json:"LimitVpcs,omitnil,omitempty" name:"LimitVpcs"`

	// 限制下载的ip设置
	LimitIps []*string `json:"LimitIps,omitnil,omitempty" name:"LimitIps"`
}

type ModifyBackupDownloadUserRestrictionRequest struct {
	*tchttp.BaseRequest
	
	// 下载限制类型，NoLimit-不限制,LimitOnlyIntranet-限制内网 ,Customize-自定义
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// 该参数仅支持 In， 表示 LimitVpc 指定的vpc可以下载。默认为In
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// In: 指定的ip可以下载； NotIn: 指定的ip不可以下载
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// 限制下载的vpc设置
	LimitVpcs []*BackupLimitVpcItem `json:"LimitVpcs,omitnil,omitempty" name:"LimitVpcs"`

	// 限制下载的ip设置
	LimitIps []*string `json:"LimitIps,omitnil,omitempty" name:"LimitIps"`
}

func (r *ModifyBackupDownloadUserRestrictionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupDownloadUserRestrictionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LimitType")
	delete(f, "VpcComparisonSymbol")
	delete(f, "IpComparisonSymbol")
	delete(f, "LimitVpcs")
	delete(f, "LimitIps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupDownloadUserRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadUserRestrictionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBackupDownloadUserRestrictionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupDownloadUserRestrictionResponseParams `json:"Response"`
}

func (r *ModifyBackupDownloadUserRestrictionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupDownloadUserRestrictionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupNameRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备份文件ID
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 备注名，长度不能超过60个字符
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
}

type ModifyBackupNameRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备份文件ID
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 备注名，长度不能超过60个字符
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyBinlogConfigRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Binlog配置信息
	BinlogConfig *BinlogConfigInfo `json:"BinlogConfig,omitnil,omitempty" name:"BinlogConfig"`
}

type ModifyBinlogConfigRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Binlog配置信息
	BinlogConfig *BinlogConfigInfo `json:"BinlogConfig,omitnil,omitempty" name:"BinlogConfig"`
}

func (r *ModifyBinlogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBinlogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "BinlogConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBinlogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBinlogConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBinlogConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBinlogConfigResponseParams `json:"Response"`
}

func (r *ModifyBinlogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBinlogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBinlogSaveDaysRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Binlog保留天数
	BinlogSaveDays *int64 `json:"BinlogSaveDays,omitnil,omitempty" name:"BinlogSaveDays"`
}

type ModifyBinlogSaveDaysRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Binlog保留天数
	BinlogSaveDays *int64 `json:"BinlogSaveDays,omitnil,omitempty" name:"BinlogSaveDays"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库名
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 新授权用户主机权限
	NewUserHostPrivileges []*UserHostPrivilege `json:"NewUserHostPrivileges,omitnil,omitempty" name:"NewUserHostPrivileges"`

	// 备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 历史授权用户主机权限
	OldUserHostPrivileges []*UserHostPrivilege `json:"OldUserHostPrivileges,omitnil,omitempty" name:"OldUserHostPrivileges"`
}

type ModifyClusterDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库名
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 新授权用户主机权限
	NewUserHostPrivileges []*UserHostPrivilege `json:"NewUserHostPrivileges,omitnil,omitempty" name:"NewUserHostPrivileges"`

	// 备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 历史授权用户主机权限
	OldUserHostPrivileges []*UserHostPrivilege `json:"OldUserHostPrivileges,omitnil,omitempty" name:"OldUserHostPrivileges"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyClusterGlobalEncryptionRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 开启或关闭全局加密
	IsOpenGlobalEncryption *bool `json:"IsOpenGlobalEncryption,omitnil,omitempty" name:"IsOpenGlobalEncryption"`
}

type ModifyClusterGlobalEncryptionRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 开启或关闭全局加密
	IsOpenGlobalEncryption *bool `json:"IsOpenGlobalEncryption,omitnil,omitempty" name:"IsOpenGlobalEncryption"`
}

func (r *ModifyClusterGlobalEncryptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterGlobalEncryptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "IsOpenGlobalEncryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterGlobalEncryptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterGlobalEncryptionResponseParams struct {
	// 异步任务id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterGlobalEncryptionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterGlobalEncryptionResponseParams `json:"Response"`
}

func (r *ModifyClusterGlobalEncryptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterGlobalEncryptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterNameRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`
}

type ModifyClusterNameRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 要修改的参数列表。每一个元素是ParamName、CurrentValue和OldValue的组合。ParamName是参数名称，CurrentValue是当前值，OldValue是之前值且不做校验
	ParamList []*ParamItem `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// 维护期间执行-yes,立即执行-no
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type ModifyClusterParamRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 要修改的参数列表。每一个元素是ParamName、CurrentValue和OldValue的组合。ParamName是参数名称，CurrentValue是当前值，OldValue是之前值且不做校验
	ParamList []*ParamItem `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// 维护期间执行-yes,立即执行-no
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 密码长度
	ValidatePasswordLength *int64 `json:"ValidatePasswordLength,omitnil,omitempty" name:"ValidatePasswordLength"`

	// 大小写字符个数
	ValidatePasswordMixedCaseCount *int64 `json:"ValidatePasswordMixedCaseCount,omitnil,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// 特殊字符个数
	ValidatePasswordSpecialCharCount *int64 `json:"ValidatePasswordSpecialCharCount,omitnil,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// 数字个数
	ValidatePasswordNumberCount *int64 `json:"ValidatePasswordNumberCount,omitnil,omitempty" name:"ValidatePasswordNumberCount"`

	// 密码强度（"MEDIUM", "STRONG"）
	ValidatePasswordPolicy *string `json:"ValidatePasswordPolicy,omitnil,omitempty" name:"ValidatePasswordPolicy"`

	// 数据字典
	ValidatePasswordDictionary []*string `json:"ValidatePasswordDictionary,omitnil,omitempty" name:"ValidatePasswordDictionary"`
}

type ModifyClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 密码长度
	ValidatePasswordLength *int64 `json:"ValidatePasswordLength,omitnil,omitempty" name:"ValidatePasswordLength"`

	// 大小写字符个数
	ValidatePasswordMixedCaseCount *int64 `json:"ValidatePasswordMixedCaseCount,omitnil,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// 特殊字符个数
	ValidatePasswordSpecialCharCount *int64 `json:"ValidatePasswordSpecialCharCount,omitnil,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// 数字个数
	ValidatePasswordNumberCount *int64 `json:"ValidatePasswordNumberCount,omitnil,omitempty" name:"ValidatePasswordNumberCount"`

	// 密码强度（"MEDIUM", "STRONG"）
	ValidatePasswordPolicy *string `json:"ValidatePasswordPolicy,omitnil,omitempty" name:"ValidatePasswordPolicy"`

	// 数据字典
	ValidatePasswordDictionary []*string `json:"ValidatePasswordDictionary,omitnil,omitempty" name:"ValidatePasswordDictionary"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyClusterReadOnlyRequestParams struct {
	// 集群ID列表
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 集群只读开关，可选值：ON，OFF
	ReadOnlyOperation *string `json:"ReadOnlyOperation,omitnil,omitempty" name:"ReadOnlyOperation"`

	// yes：在运维时间窗内修改，no：立即执行（默认值）
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type ModifyClusterReadOnlyRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID列表
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 集群只读开关，可选值：ON，OFF
	ReadOnlyOperation *string `json:"ReadOnlyOperation,omitnil,omitempty" name:"ReadOnlyOperation"`

	// yes：在运维时间窗内修改，no：立即执行（默认值）
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

func (r *ModifyClusterReadOnlyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterReadOnlyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "ReadOnlyOperation")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterReadOnlyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterReadOnlyResponseParams struct {
	// 集群任务ID列表
	ClusterTaskIds []*ClusterTaskId `json:"ClusterTaskIds,omitnil,omitempty" name:"ClusterTaskIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterReadOnlyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterReadOnlyResponseParams `json:"Response"`
}

func (r *ModifyClusterReadOnlyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterReadOnlyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterSlaveZoneRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 旧从可用区
	OldSlaveZone *string `json:"OldSlaveZone,omitnil,omitempty" name:"OldSlaveZone"`

	// 新从可用区
	NewSlaveZone *string `json:"NewSlaveZone,omitnil,omitempty" name:"NewSlaveZone"`

	// binlog同步方式。默认值：async。可选值：sync、semisync、async
	BinlogSyncWay *string `json:"BinlogSyncWay,omitnil,omitempty" name:"BinlogSyncWay"`

	// 半同步超时时间，单位ms。为保证业务稳定性，半同步复制存在退化逻辑，当主可用区集群在等待备可用区集群确认事务时若超过该超时时间，复制方式将降为异步复制。最低设置为1000ms，最高支持4294967295ms，默认10000ms。
	SemiSyncTimeout *int64 `json:"SemiSyncTimeout,omitnil,omitempty" name:"SemiSyncTimeout"`
}

type ModifyClusterSlaveZoneRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 旧从可用区
	OldSlaveZone *string `json:"OldSlaveZone,omitnil,omitempty" name:"OldSlaveZone"`

	// 新从可用区
	NewSlaveZone *string `json:"NewSlaveZone,omitnil,omitempty" name:"NewSlaveZone"`

	// binlog同步方式。默认值：async。可选值：sync、semisync、async
	BinlogSyncWay *string `json:"BinlogSyncWay,omitnil,omitempty" name:"BinlogSyncWay"`

	// 半同步超时时间，单位ms。为保证业务稳定性，半同步复制存在退化逻辑，当主可用区集群在等待备可用区集群确认事务时若超过该超时时间，复制方式将降为异步复制。最低设置为1000ms，最高支持4294967295ms，默认10000ms。
	SemiSyncTimeout *int64 `json:"SemiSyncTimeout,omitnil,omitempty" name:"SemiSyncTimeout"`
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
	delete(f, "BinlogSyncWay")
	delete(f, "SemiSyncTimeout")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterSlaveZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterSlaveZoneResponseParams struct {
	// 异步FlowId
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群新存储大小（单位G）
	NewStorageLimit *int64 `json:"NewStorageLimit,omitnil,omitempty" name:"NewStorageLimit"`

	// 集群原存储大小（单位G）
	OldStorageLimit *int64 `json:"OldStorageLimit,omitnil,omitempty" name:"OldStorageLimit"`

	// 交易模式 0-下单并支付 1-下单
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`
}

type ModifyClusterStorageRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群新存储大小（单位G）
	NewStorageLimit *int64 `json:"NewStorageLimit,omitnil,omitempty" name:"NewStorageLimit"`

	// 集群原存储大小（单位G）
	OldStorageLimit *int64 `json:"OldStorageLimit,omitnil,omitempty" name:"OldStorageLimit"`

	// 交易模式 0-下单并支付 1-下单
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`
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
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// 大订单号
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// 订单号
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 网络组id(cynosdbmysql-grp-前缀开头)或集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要修改的安全组ID列表，一个或者多个安全组ID组成的数组。
	// 注意：该入参会全量替换存量已有集合，非增量更新。修改需传入预期的全量集合。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 网络组id(cynosdbmysql-grp-前缀开头)或集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要修改的安全组ID列表，一个或者多个安全组ID组成的数组。
	// 注意：该入参会全量替换存量已有集合，非增量更新。修改需传入预期的全量集合。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
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

type ModifyDbVersionData struct {
	// 修改前版本
	OldVersion *string `json:"OldVersion,omitnil,omitempty" name:"OldVersion"`

	// 修改后版本
	NewVersion *string `json:"NewVersion,omitnil,omitempty" name:"NewVersion"`

	// 升级方式
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`
}

type ModifyInstanceData struct {
	// 变配后CPU
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 变配后内存
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 变配后存储上限
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// 变配前CPU
	OldCpu *int64 `json:"OldCpu,omitnil,omitempty" name:"OldCpu"`

	// 变配前内存
	OldMemory *int64 `json:"OldMemory,omitnil,omitempty" name:"OldMemory"`

	// 变配前存储上限
	OldStorageLimit *int64 `json:"OldStorageLimit,omitnil,omitempty" name:"OldStorageLimit"`

	// 变配前实例机器类型
	OldDeviceType *string `json:"OldDeviceType,omitnil,omitempty" name:"OldDeviceType"`

	// 变配后实例机器类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 升级方式。升级完成后切换或维护时间内切换
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// libra节点数量
	LibraNodeCount *int64 `json:"LibraNodeCount,omitnil,omitempty" name:"LibraNodeCount"`

	// 原libra节点数量
	OldLibraNodeCount *int64 `json:"OldLibraNodeCount,omitnil,omitempty" name:"OldLibraNodeCount"`
}

// Predefined struct for user
type ModifyInstanceNameRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
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
type ModifyInstanceParamRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 集群参数列表，例如 [{           "CurrentValue":"2",        "ParamName":"auto_increment_increment"}]
	ClusterParamList []*ModifyParamItem `json:"ClusterParamList,omitnil,omitempty" name:"ClusterParamList"`

	// 实例参数列表，例如[{           "CurrentValue":"2",        "ParamName":"innodb_stats_transient_sample_pages"}]
	InstanceParamList []*ModifyParamItem `json:"InstanceParamList,omitnil,omitempty" name:"InstanceParamList"`

	// yes：在运维时间窗内修改，no：立即执行（默认值）
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type ModifyInstanceParamRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 集群参数列表，例如 [{           "CurrentValue":"2",        "ParamName":"auto_increment_increment"}]
	ClusterParamList []*ModifyParamItem `json:"ClusterParamList,omitnil,omitempty" name:"ClusterParamList"`

	// 实例参数列表，例如[{           "CurrentValue":"2",        "ParamName":"innodb_stats_transient_sample_pages"}]
	InstanceParamList []*ModifyParamItem `json:"InstanceParamList,omitnil,omitempty" name:"InstanceParamList"`

	// yes：在运维时间窗内修改，no：立即执行（默认值）
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyInstanceUpgradeLimitDaysRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升级限制时间
	UpgradeLimitDays *int64 `json:"UpgradeLimitDays,omitnil,omitempty" name:"UpgradeLimitDays"`
}

type ModifyInstanceUpgradeLimitDaysRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升级限制时间
	UpgradeLimitDays *int64 `json:"UpgradeLimitDays,omitnil,omitempty" name:"UpgradeLimitDays"`
}

func (r *ModifyInstanceUpgradeLimitDaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceUpgradeLimitDaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceId")
	delete(f, "UpgradeLimitDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceUpgradeLimitDaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceUpgradeLimitDaysResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceUpgradeLimitDaysResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceUpgradeLimitDaysResponseParams `json:"Response"`
}

func (r *ModifyInstanceUpgradeLimitDaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceUpgradeLimitDaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaintainPeriodConfigRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 维护开始时间，单位为秒，如3:00为10800
	MaintainStartTime *int64 `json:"MaintainStartTime,omitnil,omitempty" name:"MaintainStartTime"`

	// 维护持续时间，单位为秒，如1小时为3600
	MaintainDuration *int64 `json:"MaintainDuration,omitnil,omitempty" name:"MaintainDuration"`

	// 每周维护日期，日期取值范围[Mon, Tue, Wed, Thu, Fri, Sat, Sun]
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitnil,omitempty" name:"MaintainWeekDays"`
}

type ModifyMaintainPeriodConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 维护开始时间，单位为秒，如3:00为10800
	MaintainStartTime *int64 `json:"MaintainStartTime,omitnil,omitempty" name:"MaintainStartTime"`

	// 维护持续时间，单位为秒，如1小时为3600
	MaintainDuration *int64 `json:"MaintainDuration,omitnil,omitempty" name:"MaintainDuration"`

	// 每周维护日期，日期取值范围[Mon, Tue, Wed, Thu, Fri, Sat, Sun]
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitnil,omitempty" name:"MaintainWeekDays"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 参数当前值
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 参数旧值（只在出参时有用）
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`

	// libra组件类型
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`
}

// Predefined struct for user
type ModifyParamTemplateRequestParams struct {
	// 模板ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 模板描述
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// 参数列表
	ParamList []*ModifyParamItem `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type ModifyParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 模板描述
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// 参数列表
	ParamList []*ModifyParamItem `json:"ParamList,omitnil,omitempty" name:"ParamList"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type ModifyParamsData struct {
	// 参数名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 修改前参数值
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`

	// 修改后参数值
	CurValue *string `json:"CurValue,omitnil,omitempty" name:"CurValue"`
}

// Predefined struct for user
type ModifyProxyDescRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 数据库代理描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyProxyDescRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 数据库代理描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 集群ID，例如cynosdbmysql-asd123
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库代理组ID，例如cynosdbmysql-proxy-qwe123
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 一致性类型；“eventual"-最终一致性, "session"-会话一致性, "global"-全局一致性
	ConsistencyType *string `json:"ConsistencyType,omitnil,omitempty" name:"ConsistencyType"`

	// 一致性超时时间。
	// 取值范围：0~1000000（微秒）,设置0则表示若只读实例出现延迟, 导致一致性策略不满足, 请求将一直等待。
	ConsistencyTimeOut *string `json:"ConsistencyTimeOut,omitnil,omitempty" name:"ConsistencyTimeOut"`

	// 读写权重分配模式；系统自动分配："system"， 自定义："custom"
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// 实例只读权重。
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitnil,omitempty" name:"InstanceWeights"`

	// 是否开启故障转移，代理出现故障后，连接地址将路由到主实例，取值："yes" , "no"
	FailOver *string `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// 是否自动添加只读实例，取值："yes" , "no"
	AutoAddRo *string `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// 是否打开读写分离。
	// 该参数已废弃，请通过RwType设置读写属性。
	OpenRw *string `json:"OpenRw,omitnil,omitempty" name:"OpenRw"`

	// 读写类型：
	// READWRITE,READONLY
	RwType *string `json:"RwType,omitnil,omitempty" name:"RwType"`

	// 事务拆分。
	// 在一个事务中拆分读和写到不同的实例上去执行。
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// 连接模式：
	// nearby,balance
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// 是否打开连接池：
	// yes,no
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// 连接池类型：
	// SessionConnectionPool
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// 连接池时间。
	// 可选范围:0~300（秒）
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// 是否将libra节点当作普通RO节点
	ApNodeAsRoNode *bool `json:"ApNodeAsRoNode,omitnil,omitempty" name:"ApNodeAsRoNode"`

	// libra节点故障，是否转发给其他节点
	ApQueryToOtherNode *bool `json:"ApQueryToOtherNode,omitnil,omitempty" name:"ApQueryToOtherNode"`
}

type ModifyProxyRwSplitRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，例如cynosdbmysql-asd123
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库代理组ID，例如cynosdbmysql-proxy-qwe123
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 一致性类型；“eventual"-最终一致性, "session"-会话一致性, "global"-全局一致性
	ConsistencyType *string `json:"ConsistencyType,omitnil,omitempty" name:"ConsistencyType"`

	// 一致性超时时间。
	// 取值范围：0~1000000（微秒）,设置0则表示若只读实例出现延迟, 导致一致性策略不满足, 请求将一直等待。
	ConsistencyTimeOut *string `json:"ConsistencyTimeOut,omitnil,omitempty" name:"ConsistencyTimeOut"`

	// 读写权重分配模式；系统自动分配："system"， 自定义："custom"
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// 实例只读权重。
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitnil,omitempty" name:"InstanceWeights"`

	// 是否开启故障转移，代理出现故障后，连接地址将路由到主实例，取值："yes" , "no"
	FailOver *string `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// 是否自动添加只读实例，取值："yes" , "no"
	AutoAddRo *string `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// 是否打开读写分离。
	// 该参数已废弃，请通过RwType设置读写属性。
	OpenRw *string `json:"OpenRw,omitnil,omitempty" name:"OpenRw"`

	// 读写类型：
	// READWRITE,READONLY
	RwType *string `json:"RwType,omitnil,omitempty" name:"RwType"`

	// 事务拆分。
	// 在一个事务中拆分读和写到不同的实例上去执行。
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// 连接模式：
	// nearby,balance
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// 是否打开连接池：
	// yes,no
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// 连接池类型：
	// SessionConnectionPool
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// 连接池时间。
	// 可选范围:0~300（秒）
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// 是否将libra节点当作普通RO节点
	ApNodeAsRoNode *bool `json:"ApNodeAsRoNode,omitnil,omitempty" name:"ApNodeAsRoNode"`

	// libra节点故障，是否转发给其他节点
	ApQueryToOtherNode *bool `json:"ApQueryToOtherNode,omitnil,omitempty" name:"ApQueryToOtherNode"`
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
	delete(f, "ApNodeAsRoNode")
	delete(f, "ApQueryToOtherNode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProxyRwSplitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProxyRwSplitResponseParams struct {
	// 异步FlowId
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 异步任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 需要建立绑定关系的集群ID
	BindClusterIds []*string `json:"BindClusterIds,omitnil,omitempty" name:"BindClusterIds"`

	// 需要解除绑定关系的集群ID
	UnbindClusterIds []*string `json:"UnbindClusterIds,omitnil,omitempty" name:"UnbindClusterIds"`
}

type ModifyResourcePackageClustersRequest struct {
	*tchttp.BaseRequest
	
	// 资源包唯一ID
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 需要建立绑定关系的集群ID
	BindClusterIds []*string `json:"BindClusterIds,omitnil,omitempty" name:"BindClusterIds"`

	// 需要解除绑定关系的集群ID
	UnbindClusterIds []*string `json:"UnbindClusterIds,omitnil,omitempty" name:"UnbindClusterIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 自定义的资源包名称，最长支持120个字符
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type ModifyResourcePackageNameRequest struct {
	*tchttp.BaseRequest
	
	// 资源包唯一ID
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 自定义的资源包名称，最长支持120个字符
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyResourcePackagesDeductionPriorityRequestParams struct {
	// 需要修改优先级的资源包类型，CCU：计算资源包，DISK：存储资源包
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 修改后的抵扣优先级对于哪个cynos资源生效
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 资源包抵扣优先级
	DeductionPriorities []*PackagePriority `json:"DeductionPriorities,omitnil,omitempty" name:"DeductionPriorities"`
}

type ModifyResourcePackagesDeductionPriorityRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改优先级的资源包类型，CCU：计算资源包，DISK：存储资源包
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 修改后的抵扣优先级对于哪个cynos资源生效
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 资源包抵扣优先级
	DeductionPriorities []*PackagePriority `json:"DeductionPriorities,omitnil,omitempty" name:"DeductionPriorities"`
}

func (r *ModifyResourcePackagesDeductionPriorityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePackagesDeductionPriorityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageType")
	delete(f, "ClusterId")
	delete(f, "DeductionPriorities")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourcePackagesDeductionPriorityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePackagesDeductionPriorityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourcePackagesDeductionPriorityResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourcePackagesDeductionPriorityResponseParams `json:"Response"`
}

func (r *ModifyResourcePackagesDeductionPriorityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePackagesDeductionPriorityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServerlessStrategyRequestParams struct {
	// serverless集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群是否自动暂停，可选范围
	// <li>yes</li>
	// <li>no</li>
	AutoPause *string `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// 集群自动暂停的延迟，单位秒，可选范围[600,691200]，默认600
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitnil,omitempty" name:"AutoPauseDelay"`

	// 该参数暂时无效
	AutoScaleUpDelay *int64 `json:"AutoScaleUpDelay,omitnil,omitempty" name:"AutoScaleUpDelay"`

	// 该参数暂时无效
	AutoScaleDownDelay *int64 `json:"AutoScaleDownDelay,omitnil,omitempty" name:"AutoScaleDownDelay"`

	// cpu最小值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// cpu最大值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// 只读实例cpu最小值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MinRoCpu *float64 `json:"MinRoCpu,omitnil,omitempty" name:"MinRoCpu"`

	// 只读cpu最大值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MaxRoCpu *float64 `json:"MaxRoCpu,omitnil,omitempty" name:"MaxRoCpu"`

	// 只读节点最小个数
	MinRoCount *int64 `json:"MinRoCount,omitnil,omitempty" name:"MinRoCount"`

	// 只读节点最大个数
	MaxRoCount *int64 `json:"MaxRoCount,omitnil,omitempty" name:"MaxRoCount"`

	// 是否开启归档，可选范围<li>yes</li><li>no</li>默认值:yes
	AutoArchive *string `json:"AutoArchive,omitnil,omitempty" name:"AutoArchive"`

	// 升级类型。 默认值：upgradeImmediate。 可选值： upgradeImmediate：立即完成修改 upgradeInMaintain：在维护时间窗口内完成修改
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`
}

type ModifyServerlessStrategyRequest struct {
	*tchttp.BaseRequest
	
	// serverless集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群是否自动暂停，可选范围
	// <li>yes</li>
	// <li>no</li>
	AutoPause *string `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// 集群自动暂停的延迟，单位秒，可选范围[600,691200]，默认600
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitnil,omitempty" name:"AutoPauseDelay"`

	// 该参数暂时无效
	AutoScaleUpDelay *int64 `json:"AutoScaleUpDelay,omitnil,omitempty" name:"AutoScaleUpDelay"`

	// 该参数暂时无效
	AutoScaleDownDelay *int64 `json:"AutoScaleDownDelay,omitnil,omitempty" name:"AutoScaleDownDelay"`

	// cpu最小值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// cpu最大值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// 只读实例cpu最小值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MinRoCpu *float64 `json:"MinRoCpu,omitnil,omitempty" name:"MinRoCpu"`

	// 只读cpu最大值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MaxRoCpu *float64 `json:"MaxRoCpu,omitnil,omitempty" name:"MaxRoCpu"`

	// 只读节点最小个数
	MinRoCount *int64 `json:"MinRoCount,omitnil,omitempty" name:"MinRoCount"`

	// 只读节点最大个数
	MaxRoCount *int64 `json:"MaxRoCount,omitnil,omitempty" name:"MaxRoCount"`

	// 是否开启归档，可选范围<li>yes</li><li>no</li>默认值:yes
	AutoArchive *string `json:"AutoArchive,omitnil,omitempty" name:"AutoArchive"`

	// 升级类型。 默认值：upgradeImmediate。 可选值： upgradeImmediate：立即完成修改 upgradeInMaintain：在维护时间窗口内完成修改
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`
}

func (r *ModifyServerlessStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServerlessStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AutoPause")
	delete(f, "AutoPauseDelay")
	delete(f, "AutoScaleUpDelay")
	delete(f, "AutoScaleDownDelay")
	delete(f, "MinCpu")
	delete(f, "MaxCpu")
	delete(f, "MinRoCpu")
	delete(f, "MaxRoCpu")
	delete(f, "MinRoCount")
	delete(f, "MaxRoCount")
	delete(f, "AutoArchive")
	delete(f, "UpgradeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyServerlessStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServerlessStrategyResponseParams struct {
	// 异步流程id
	//
	// Deprecated: FlowId is deprecated.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 任务id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyServerlessStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServerlessStrategyResponseParams `json:"Response"`
}

func (r *ModifyServerlessStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServerlessStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVipVportRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例组id
	//
	// Deprecated: InstanceGrpId is deprecated.
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// 实例组id
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// 需要修改的目的ip
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 需要修改的目的端口
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 旧ip回收前的保留时间，单位小时，0表示立即回收
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitnil,omitempty" name:"OldIpReserveHours"`
}

type ModifyVipVportRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例组id
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// 实例组id
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// 需要修改的目的ip
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 需要修改的目的端口
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 旧ip回收前的保留时间，单位小时，0表示立即回收
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitnil,omitempty" name:"OldIpReserveHours"`
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
	delete(f, "InstanceGroupId")
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	IsDisable *string `json:"IsDisable,omitnil,omitempty" name:"IsDisable"`

	// 模块名
	ModuleName *string `json:"ModuleName,omitnil,omitempty" name:"ModuleName"`
}

type NetAddr struct {
	// 内网ip
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 内网端口号
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 外网域名
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// 外网端口号
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// 网络类型（ro-只读,rw/ha-读写）
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 子网ID
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 私有网络ID
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 外网IP
	WanIP *string `json:"WanIP,omitnil,omitempty" name:"WanIP"`

	// 外网状态
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// 实例组ID
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`
}

type NewAccount struct {
	// 账户名，包含字母数字_,以字母开头，字母或数字结尾，长度1-30
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 密码，密码长度范围为8到64个字符
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// 主机(%或ipv4地址)
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户最大连接数，不能大于10240
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`
}

type ObjectTask struct {
	// 任务自增ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务状态
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 任务ID（集群ID|实例组ID|实例ID）
	ObjectId *string `json:"ObjectId,omitnil,omitempty" name:"ObjectId"`

	// 任务类型
	ObjectType *string `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`
}

// Predefined struct for user
type OfflineClusterRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type OfflineClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID数组
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
}

type OfflineInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID数组
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 端口
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 期望执行回收时间
	ReturnTime *string `json:"ReturnTime,omitnil,omitempty" name:"ReturnTime"`
}

// Predefined struct for user
type OpenAuditServiceRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志保留时长。
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 高频日志保留时长。
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// 审计规则(废弃)。
	//
	// Deprecated: AuditRuleFilters is deprecated.
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// 规则模板ID。同AuditRuleFilters都不填是全审计。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// 审计类型。true-全审计；默认false-规则审计。
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`
}

type OpenAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志保留时长。
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 高频日志保留时长。
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// 审计规则(废弃)。
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// 规则模板ID。同AuditRuleFilters都不填是全审计。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// 审计类型。true-全审计；默认false-规则审计。
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`
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
	delete(f, "AuditAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenAuditServiceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 密码长度
	ValidatePasswordLength *int64 `json:"ValidatePasswordLength,omitnil,omitempty" name:"ValidatePasswordLength"`

	// 大小写字符个数
	ValidatePasswordMixedCaseCount *int64 `json:"ValidatePasswordMixedCaseCount,omitnil,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// 特殊字符个数
	ValidatePasswordSpecialCharCount *int64 `json:"ValidatePasswordSpecialCharCount,omitnil,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// 数字个数
	ValidatePasswordNumberCount *int64 `json:"ValidatePasswordNumberCount,omitnil,omitempty" name:"ValidatePasswordNumberCount"`

	// 密码强度（"MEDIUM", "STRONG"）
	ValidatePasswordPolicy *string `json:"ValidatePasswordPolicy,omitnil,omitempty" name:"ValidatePasswordPolicy"`

	// 数据字典
	ValidatePasswordDictionary []*string `json:"ValidatePasswordDictionary,omitnil,omitempty" name:"ValidatePasswordDictionary"`
}

type OpenClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 密码长度
	ValidatePasswordLength *int64 `json:"ValidatePasswordLength,omitnil,omitempty" name:"ValidatePasswordLength"`

	// 大小写字符个数
	ValidatePasswordMixedCaseCount *int64 `json:"ValidatePasswordMixedCaseCount,omitnil,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// 特殊字符个数
	ValidatePasswordSpecialCharCount *int64 `json:"ValidatePasswordSpecialCharCount,omitnil,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// 数字个数
	ValidatePasswordNumberCount *int64 `json:"ValidatePasswordNumberCount,omitnil,omitempty" name:"ValidatePasswordNumberCount"`

	// 密码强度（"MEDIUM", "STRONG"）
	ValidatePasswordPolicy *string `json:"ValidatePasswordPolicy,omitnil,omitempty" name:"ValidatePasswordPolicy"`

	// 数据字典
	ValidatePasswordDictionary []*string `json:"ValidatePasswordDictionary,omitnil,omitempty" name:"ValidatePasswordDictionary"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type OpenClusterReadOnlyInstanceGroupAccessRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 安全组ID 
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type OpenClusterReadOnlyInstanceGroupAccessRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 安全组ID 
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *OpenClusterReadOnlyInstanceGroupAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenClusterReadOnlyInstanceGroupAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Port")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenClusterReadOnlyInstanceGroupAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenClusterReadOnlyInstanceGroupAccessResponseParams struct {
	// 开启流程ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenClusterReadOnlyInstanceGroupAccessResponse struct {
	*tchttp.BaseResponse
	Response *OpenClusterReadOnlyInstanceGroupAccessResponseParams `json:"Response"`
}

func (r *OpenClusterReadOnlyInstanceGroupAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenClusterReadOnlyInstanceGroupAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenClusterTransparentEncryptRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 秘钥类型(cloud,custom)
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`

	// 秘钥Id
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 秘钥地域
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`

	// 是否开启全局加密
	IsOpenGlobalEncryption *bool `json:"IsOpenGlobalEncryption,omitnil,omitempty" name:"IsOpenGlobalEncryption"`
}

type OpenClusterTransparentEncryptRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 秘钥类型(cloud,custom)
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`

	// 秘钥Id
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 秘钥地域
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`

	// 是否开启全局加密
	IsOpenGlobalEncryption *bool `json:"IsOpenGlobalEncryption,omitnil,omitempty" name:"IsOpenGlobalEncryption"`
}

func (r *OpenClusterTransparentEncryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenClusterTransparentEncryptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "KeyType")
	delete(f, "KeyId")
	delete(f, "KeyRegion")
	delete(f, "IsOpenGlobalEncryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenClusterTransparentEncryptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenClusterTransparentEncryptResponseParams struct {
	// 异步任务id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenClusterTransparentEncryptResponse struct {
	*tchttp.BaseResponse
	Response *OpenClusterTransparentEncryptResponseParams `json:"Response"`
}

func (r *OpenClusterTransparentEncryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenClusterTransparentEncryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenReadOnlyInstanceExclusiveAccessRequestParams struct {
	// 请使用 [集群信息描述](https://cloud.tencent.com/document/api/1003/48086) 获取 clusterId。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 请使用 [集群信息描述](https://cloud.tencent.com/document/api/1003/48086) 获取 instanceId。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定的 vpc ID，请使用 [查询私有网络列表](https://cloud.tencent.com/document/api/215/15778) 获取 vpc ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 指定的子网 ID，如果设置了 vpc ID，则 SubnetId 必填，请使用 [查询子网列表](https://cloud.tencent.com/document/api/215/15784) 获取 SubnetId。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 用户自定义的端口。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 安全组 ID，请使用 [查看安全组](https://cloud.tencent.com/document/api/215/15808) 获取 SecurityGroupId。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type OpenReadOnlyInstanceExclusiveAccessRequest struct {
	*tchttp.BaseRequest
	
	// 请使用 [集群信息描述](https://cloud.tencent.com/document/api/1003/48086) 获取 clusterId。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 请使用 [集群信息描述](https://cloud.tencent.com/document/api/1003/48086) 获取 instanceId。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定的 vpc ID，请使用 [查询私有网络列表](https://cloud.tencent.com/document/api/215/15778) 获取 vpc ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 指定的子网 ID，如果设置了 vpc ID，则 SubnetId 必填，请使用 [查询子网列表](https://cloud.tencent.com/document/api/215/15784) 获取 SubnetId。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 用户自定义的端口。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 安全组 ID，请使用 [查看安全组](https://cloud.tencent.com/document/api/215/15808) 获取 SecurityGroupId。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type OpenSSLRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type OpenSSLRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *OpenSSLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenSSLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenSSLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenSSLResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 任务id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenSSLResponse struct {
	*tchttp.BaseResponse
	Response *OpenSSLResponseParams `json:"Response"`
}

func (r *OpenSSLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenSSLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenWanRequestParams struct {
	// 实例组id
	//
	// Deprecated: InstanceGrpId is deprecated.
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例组id
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`
}

type OpenWanRequest struct {
	*tchttp.BaseRequest
	
	// 实例组id
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例组id
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`
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
	delete(f, "InstanceId")
	delete(f, "InstanceGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenWanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenWanResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 资源包唯一ID
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 资源包名称
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 资源包类型
	// CCU-计算资源包，DISK-存储资源包
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 资源包使用地域
	// china-中国内地通用，overseas-港澳台及海外通用
	PackageRegion *string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// 资源包状态
	// creating-创建中；
	// using-使用中；
	// expired-已过期；
	// normal_finish-使用完；
	// apply_refund-申请退费中；
	// refund-已退费。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 资源包总量
	PackageTotalSpec *float64 `json:"PackageTotalSpec,omitnil,omitempty" name:"PackageTotalSpec"`

	// 资源包已使用量
	PackageUsedSpec *float64 `json:"PackageUsedSpec,omitnil,omitempty" name:"PackageUsedSpec"`

	// 是否还有库存余量
	HasQuota *bool `json:"HasQuota,omitnil,omitempty" name:"HasQuota"`

	// 绑定实例信息
	BindInstanceInfos []*BindInstanceInfo `json:"BindInstanceInfos,omitnil,omitempty" name:"BindInstanceInfos"`

	// 生效时间：2022-07-01 00:00:00
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 失效时间：2022-08-01 00:00:00
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 资源包历史绑定（已解绑）实例信息
	HistoryBindResourceInfos []*BindInstanceInfo `json:"HistoryBindResourceInfos,omitnil,omitempty" name:"HistoryBindResourceInfos"`
}

type PackageDetail struct {
	// AppId账户ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 资源包唯一ID
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 成功抵扣容量
	SuccessDeductSpec *float64 `json:"SuccessDeductSpec,omitnil,omitempty" name:"SuccessDeductSpec"`

	// 截止当前，资源包已使用的容量
	PackageTotalUsedSpec *float64 `json:"PackageTotalUsedSpec,omitnil,omitempty" name:"PackageTotalUsedSpec"`

	// 抵扣开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 抵扣结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 扩展信息
	ExtendInfo *string `json:"ExtendInfo,omitnil,omitempty" name:"ExtendInfo"`
}

type PackagePriority struct {
	// 需要自定义抵扣优先级的资源包
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 自定义的抵扣优先级
	DeductionPriority *int64 `json:"DeductionPriority,omitnil,omitempty" name:"DeductionPriority"`
}

type ParamDetail struct {
	// 参数名称
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 参数类型：integer，enum，float，string，func
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// true-支持"func"，false-不支持公式
	SupportFunc *bool `json:"SupportFunc,omitnil,omitempty" name:"SupportFunc"`

	// 默认值
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// 参数描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 参数当前值
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 修改参数后，是否需要重启数据库以使参数生效。0-不需要重启，1-需要重启。
	NeedReboot *int64 `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// 参数容许的最大值
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// 参数容许的最小值
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// 参数的可选枚举值。如果为非枚举值，则为空
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// 1：全局参数，0：非全局参数
	IsGlobal *int64 `json:"IsGlobal,omitnil,omitempty" name:"IsGlobal"`

	// 匹配类型，multiVal
	MatchType *string `json:"MatchType,omitnil,omitempty" name:"MatchType"`

	// 匹配目标值，当multiVal时，各个key用，分割
	MatchValue *string `json:"MatchValue,omitnil,omitempty" name:"MatchValue"`

	// true-为公式，false-非公式
	IsFunc *bool `json:"IsFunc,omitnil,omitempty" name:"IsFunc"`

	// 参数设置为公式时，Func返回设置的公式内容
	Func *string `json:"Func,omitnil,omitempty" name:"Func"`

	// 参数是否可修改
	ModifiableInfo *ModifiableInfo `json:"ModifiableInfo,omitnil,omitempty" name:"ModifiableInfo"`

	// 支持公式的参数的默认公式样式
	FuncPattern *string `json:"FuncPattern,omitnil,omitempty" name:"FuncPattern"`
}

type ParamInfo struct {
	// 当前值
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 默认值
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// 参数为enum/string/bool时，可选值列表
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// 参数类型为float/integer时的最大值
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// 参数类型为float/integer时的最小值
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// 参数名称
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 是否需要重启生效
	NeedReboot *int64 `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// 参数类型：integer/float/string/enum/bool
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// 匹配类型，multiVal, regex在参数类型是string时使用
	MatchType *string `json:"MatchType,omitnil,omitempty" name:"MatchType"`

	// 匹配目标值，当multiVal时，各个key用;分割
	MatchValue *string `json:"MatchValue,omitnil,omitempty" name:"MatchValue"`

	// 参数描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否为全局参数
	IsGlobal *int64 `json:"IsGlobal,omitnil,omitempty" name:"IsGlobal"`

	// 参数是否可修改
	ModifiableInfo *ModifiableInfo `json:"ModifiableInfo,omitnil,omitempty" name:"ModifiableInfo"`

	// 是否为函数
	IsFunc *bool `json:"IsFunc,omitnil,omitempty" name:"IsFunc"`

	// 函数
	Func *string `json:"Func,omitnil,omitempty" name:"Func"`

	// 支持公式的参数的默认公式样式
	FuncPattern *string `json:"FuncPattern,omitnil,omitempty" name:"FuncPattern"`
}

type ParamItem struct {
	// 参数名称
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 当前值
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 原有值
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`
}

type ParamItemDetail struct {
	// 当前值
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 默认值
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// 参数的可选枚举值。如果为非枚举值，则为空
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// 1：全局参数，0：非全局参数
	IsGlobal *int64 `json:"IsGlobal,omitnil,omitempty" name:"IsGlobal"`

	// 最大值
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// 最小值
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// 修改参数后，是否需要重启数据库以使参数生效。0-不需要重启，1-需要重启。
	NeedReboot *int64 `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// 参数名称
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 参数类型：integer，enum，float，string，func
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// 参数描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 类型是否为公式
	IsFunc *bool `json:"IsFunc,omitnil,omitempty" name:"IsFunc"`

	// 参数配置公式
	Func *string `json:"Func,omitnil,omitempty" name:"Func"`

	// 支持公式的参数的默认公式样式
	FuncPattern *string `json:"FuncPattern,omitnil,omitempty" name:"FuncPattern"`
}

type ParamItemInfo struct {
	// 参数名字
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 参数新值
	NewValue *string `json:"NewValue,omitnil,omitempty" name:"NewValue"`

	// 参数旧值
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`

	// 参数公式
	ValueFunction *string `json:"ValueFunction,omitnil,omitempty" name:"ValueFunction"`
}

type ParamTemplateListInfo struct {
	// 参数模板ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 参数模板名称
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 参数模板描述
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// 引擎版本
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 数据库类型，可选值：NORMAL，SERVERLESS
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// 参数模板详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamInfoSet []*TemplateParamInfo `json:"ParamInfoSet,omitnil,omitempty" name:"ParamInfoSet"`
}

// Predefined struct for user
type PauseServerlessRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 是否强制暂停，忽略当前的用户链接  0:不强制  1:强制， 默认为1
	ForcePause *int64 `json:"ForcePause,omitnil,omitempty" name:"ForcePause"`
}

type PauseServerlessRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 是否强制暂停，忽略当前的用户链接  0:不强制  1:强制， 默认为1
	ForcePause *int64 `json:"ForcePause,omitnil,omitempty" name:"ForcePause"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 来源Ip或Ip段，例如192.168.0.0/16
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// 端口
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// 网络协议，支持udp、tcp等
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// 协议端口ID或者协议端口组ID。
	ServiceModule *string `json:"ServiceModule,omitnil,omitempty" name:"ServiceModule"`

	// IP地址ID或者ID地址组ID。
	AddressModule *string `json:"AddressModule,omitnil,omitempty" name:"AddressModule"`

	// id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type ProxyConfig struct {
	// 数据库代理组节点个数。该参数不再建议使用,建议使用ProxyZones
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// cpu核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 连接池类型:SessionConnectionPool(会话级别连接池 )
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// 是否开启连接池,yes-开启，no-不开启
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// 连接池阈值:单位（秒）
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// 描述说明
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据库节点信息（该参数与ProxyCount需要任选一个输入）
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`
}

type ProxyConfigInfo struct {
	// 数据库代理组节点个数。该参数不再建议使用,建议使用ProxyZones
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// cpu核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 描述说明
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据库节点信息（该参数与ProxyCount需要任选一个输入）
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`
}

type ProxyConnectionPoolInfo struct {
	// 连接池保持阈值：单位（秒）
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// 是否开启了连接池
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// 连接池类型：SessionConnectionPool（会话级别连接池）
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`
}

type ProxyEndPointConfigInfo struct {
	// 所属VPC网络ID
	UniqueVpcId *string `json:"UniqueVpcId,omitnil,omitempty" name:"UniqueVpcId"`

	// 所属子网ID
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil,omitempty" name:"UniqueSubnetId"`

	// 安全组id数组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 权重模式： system-系统分配，custom-自定义
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// 是否自动添加只读实例，yes-是，no-不自动添加
	AutoAddRo *string `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// 读写属性： READWRITE,READONLY
	RwType *string `json:"RwType,omitnil,omitempty" name:"RwType"`

	// 权重信息
	InstanceNameWeights []*InstanceNameWeight `json:"InstanceNameWeights,omitnil,omitempty" name:"InstanceNameWeights"`
}

type ProxyGroup struct {
	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 数据库代理组节点个数
	ProxyNodeCount *int64 `json:"ProxyNodeCount,omitnil,omitempty" name:"ProxyNodeCount"`

	// 数据库代理组状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 当前代理版本
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitnil,omitempty" name:"CurrentProxyVersion"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 用户AppId
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 读写节点开通数据库代理
	OpenRw *string `json:"OpenRw,omitnil,omitempty" name:"OpenRw"`
}

type ProxyGroupInfo struct {
	// 数据库代理组
	ProxyGroup *ProxyGroup `json:"ProxyGroup,omitnil,omitempty" name:"ProxyGroup"`

	// 数据库代理组读写分离信息
	ProxyGroupRwInfo *ProxyGroupRwInfo `json:"ProxyGroupRwInfo,omitnil,omitempty" name:"ProxyGroupRwInfo"`

	// 数据库代理节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyNodes []*ProxyNodeInfo `json:"ProxyNodes,omitnil,omitempty" name:"ProxyNodes"`

	// 数据库代理连接池信息
	ConnectionPool *ProxyConnectionPoolInfo `json:"ConnectionPool,omitnil,omitempty" name:"ConnectionPool"`

	// 数据库代理网络信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetAddrInfos []*NetAddr `json:"NetAddrInfos,omitnil,omitempty" name:"NetAddrInfos"`

	// 数据库代理任务集
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`
}

type ProxyGroupRwInfo struct {
	// 一致性类型 eventual-最终一致性,global-全局一致性,session-会话一致性
	ConsistencyType *string `json:"ConsistencyType,omitnil,omitempty" name:"ConsistencyType"`

	// 一致性超时时间
	ConsistencyTimeOut *int64 `json:"ConsistencyTimeOut,omitnil,omitempty" name:"ConsistencyTimeOut"`

	// 权重模式 system-系统分配，custom-自定义
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// 是否开启故障转移
	FailOver *string `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// 是否自动添加只读实例，yes-是，no-不自动添加
	AutoAddRo *string `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// 实例权重数组
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitnil,omitempty" name:"InstanceWeights"`

	// 是否开通读写节点，yse-是，no-否
	OpenRw *string `json:"OpenRw,omitnil,omitempty" name:"OpenRw"`

	// 读写属性，可选值：READWRITE,READONLY
	RwType *string `json:"RwType,omitnil,omitempty" name:"RwType"`

	// 事务拆分
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// 连接模式，可选值：balance，nearby
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// 是否将libra节点当作普通RO节点
	ApNodeAsRoNode *bool `json:"ApNodeAsRoNode,omitnil,omitempty" name:"ApNodeAsRoNode"`

	// libra节点故障，是否转发给其他节点
	ApQueryToOtherNode *bool `json:"ApQueryToOtherNode,omitnil,omitempty" name:"ApQueryToOtherNode"`
}

type ProxyInstanceWeight struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例权重
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type ProxyNodeInfo struct {
	// 数据库代理节点ID
	ProxyNodeId *string `json:"ProxyNodeId,omitnil,omitempty" name:"ProxyNodeId"`

	// 节点当前连接数, DescribeProxyNodes接口此字段值不返回
	ProxyNodeConnections *int64 `json:"ProxyNodeConnections,omitnil,omitempty" name:"ProxyNodeConnections"`

	// 数据库代理节点cpu
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 数据库代理节点内存
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 数据库代理节点状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 用户AppID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 数据库代理节点名字
	OssProxyNodeName *string `json:"OssProxyNodeName,omitnil,omitempty" name:"OssProxyNodeName"`
}

type ProxySpec struct {
	// 数据库代理cpu核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 数据库代理内存
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`
}

type ProxyVersionInfo struct {
	// proxy版本号
	ProxyVersion *string `json:"ProxyVersion,omitnil,omitempty" name:"ProxyVersion"`

	// 版本描述：GA:稳定版  BETA:尝鲜版，DEPRECATED:过旧，
	ProxyVersionType *string `json:"ProxyVersionType,omitnil,omitempty" name:"ProxyVersionType"`
}

type ProxyZone struct {
	// proxy节点可用区
	ProxyNodeZone *string `json:"ProxyNodeZone,omitnil,omitempty" name:"ProxyNodeZone"`

	// proxy节点数量
	ProxyNodeCount *int64 `json:"ProxyNodeCount,omitnil,omitempty" name:"ProxyNodeCount"`
}

type QueryFilter struct {
	// 搜索字符串
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 搜索字段，目前支持："InstanceId", "ProjectId", "InstanceName", "Vip"
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// 是否精确匹配
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`

	// 搜索字段
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 操作符
	//
	// Deprecated: Operator is deprecated.
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type QueryParamFilter struct {
	// 搜索字段，目前支持：ProxyGroupId
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// 搜索字符串
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 是否精确匹配
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`
}

// Predefined struct for user
type RefundResourcePackageRequestParams struct {
	// 资源包唯一ID
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`
}

type RefundResourcePackageRequest struct {
	*tchttp.BaseRequest
	
	// 资源包唯一ID
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`
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
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

type ReloadBalanceProxyNodeRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 异步任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 从可用区
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`
}

type RemoveClusterSlaveZoneRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 从可用区
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type RenewClustersRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 续费时长
	TimeSpan *float64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 时间单位 y,m,d,h,i,s
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 交易模式 0-下单并支付 1-下单
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`
}

type RenewClustersRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 续费时长
	TimeSpan *float64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 时间单位 y,m,d,h,i,s
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 交易模式 0-下单并支付 1-下单
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`
}

func (r *RenewClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "DealMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewClustersResponseParams struct {
	// 预付费总订单号
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// 退款订单号
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 冻结流水，一次开通一个冻结流水
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// 每个订单号对应的发货资源id列表
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 集群id列表
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewClustersResponse struct {
	*tchttp.BaseResponse
	Response *RenewClustersResponseParams `json:"Response"`
}

func (r *RenewClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplayInstanceAuditLogRequestParams struct {
	// 源集群id
	SourceClusterId *string `json:"SourceClusterId,omitnil,omitempty" name:"SourceClusterId"`

	// 源实例id
	SourceInstanceId *string `json:"SourceInstanceId,omitnil,omitempty" name:"SourceInstanceId"`

	// 目标集群id
	// 目标集群必须为原始集群三天内克隆出的集群。
	TargetClusterId *string `json:"TargetClusterId,omitnil,omitempty" name:"TargetClusterId"`

	// 目标实例id
	TargetInstanceId *string `json:"TargetInstanceId,omitnil,omitempty" name:"TargetInstanceId"`

	// 用户名.需要host为%的用户名
	TargetUserName *string `json:"TargetUserName,omitnil,omitempty" name:"TargetUserName"`

	// 密码
	TargetPassword *string `json:"TargetPassword,omitnil,omitempty" name:"TargetPassword"`

	// 开始时间。时间格式为：yyyy-DD-mm hh:mm:ss
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。时间格式为：yyyy-DD-mm hh:mm:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ReplayInstanceAuditLogRequest struct {
	*tchttp.BaseRequest
	
	// 源集群id
	SourceClusterId *string `json:"SourceClusterId,omitnil,omitempty" name:"SourceClusterId"`

	// 源实例id
	SourceInstanceId *string `json:"SourceInstanceId,omitnil,omitempty" name:"SourceInstanceId"`

	// 目标集群id
	// 目标集群必须为原始集群三天内克隆出的集群。
	TargetClusterId *string `json:"TargetClusterId,omitnil,omitempty" name:"TargetClusterId"`

	// 目标实例id
	TargetInstanceId *string `json:"TargetInstanceId,omitnil,omitempty" name:"TargetInstanceId"`

	// 用户名.需要host为%的用户名
	TargetUserName *string `json:"TargetUserName,omitnil,omitempty" name:"TargetUserName"`

	// 密码
	TargetPassword *string `json:"TargetPassword,omitnil,omitempty" name:"TargetPassword"`

	// 开始时间。时间格式为：yyyy-DD-mm hh:mm:ss
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。时间格式为：yyyy-DD-mm hh:mm:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *ReplayInstanceAuditLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplayInstanceAuditLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceClusterId")
	delete(f, "SourceInstanceId")
	delete(f, "TargetClusterId")
	delete(f, "TargetInstanceId")
	delete(f, "TargetUserName")
	delete(f, "TargetPassword")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplayInstanceAuditLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplayInstanceAuditLogResponseParams struct {
	// 任务id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReplayInstanceAuditLogResponse struct {
	*tchttp.BaseResponse
	Response *ReplayInstanceAuditLogResponseParams `json:"Response"`
}

func (r *ReplayInstanceAuditLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplayInstanceAuditLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAccountPasswordRequestParams struct {
	// 数据库账号名
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 数据库账号新密码
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 主机，不填默认为"%"
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 数据库账号名
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 数据库账号新密码
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 主机，不填默认为"%"
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 资源包类型：CCU：计算资源包
	// DISK：存储资源包
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 当前资源包绑定在当前实例下的抵扣优先级
	DeductionPriority *int64 `json:"DeductionPriority,omitnil,omitempty" name:"DeductionPriority"`
}

// Predefined struct for user
type RestartInstanceRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type RestartInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type ResumeServerlessRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`

	// 数据库表权限数组
	DbTablePrivileges []*string `json:"DbTablePrivileges,omitnil,omitempty" name:"DbTablePrivileges"`

	// 数据库表信息
	DbTables []*DbTable `json:"DbTables,omitnil,omitempty" name:"DbTables"`
}

type RevokeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 账号信息
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`

	// 数据库表权限数组
	DbTablePrivileges []*string `json:"DbTablePrivileges,omitnil,omitempty" name:"DbTablePrivileges"`

	// 数据库表信息
	DbTables []*DbTable `json:"DbTables,omitnil,omitempty" name:"DbTables"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 回档策略 timeRollback-按时间点回档 snapRollback-按备份文件回档
	RollbackStrategy *string `json:"RollbackStrategy,omitnil,omitempty" name:"RollbackStrategy"`

	// 备份文件ID。
	// 回档策略为按备份文件回档时必填。
	RollbackId *uint64 `json:"RollbackId,omitnil,omitempty" name:"RollbackId"`

	// 期望回档时间。
	// 回档策略为timeRollback按时间点回档时必填。
	ExpectTime *string `json:"ExpectTime,omitnil,omitempty" name:"ExpectTime"`

	// 期望阈值（已废弃）
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitnil,omitempty" name:"ExpectTimeThresh"`

	// 回档数据库列表
	RollbackDatabases []*RollbackDatabase `json:"RollbackDatabases,omitnil,omitempty" name:"RollbackDatabases"`

	// 回档数据库表列表
	RollbackTables []*RollbackTable `json:"RollbackTables,omitnil,omitempty" name:"RollbackTables"`

	// 按时间点回档模式，full: 普通; db: 快速; table: 极速  （默认是普通）
	RollbackMode *string `json:"RollbackMode,omitnil,omitempty" name:"RollbackMode"`
}

type RollBackClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 回档策略 timeRollback-按时间点回档 snapRollback-按备份文件回档
	RollbackStrategy *string `json:"RollbackStrategy,omitnil,omitempty" name:"RollbackStrategy"`

	// 备份文件ID。
	// 回档策略为按备份文件回档时必填。
	RollbackId *uint64 `json:"RollbackId,omitnil,omitempty" name:"RollbackId"`

	// 期望回档时间。
	// 回档策略为timeRollback按时间点回档时必填。
	ExpectTime *string `json:"ExpectTime,omitnil,omitempty" name:"ExpectTime"`

	// 期望阈值（已废弃）
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitnil,omitempty" name:"ExpectTimeThresh"`

	// 回档数据库列表
	RollbackDatabases []*RollbackDatabase `json:"RollbackDatabases,omitnil,omitempty" name:"RollbackDatabases"`

	// 回档数据库表列表
	RollbackTables []*RollbackTable `json:"RollbackTables,omitnil,omitempty" name:"RollbackTables"`

	// 按时间点回档模式，full: 普通; db: 快速; table: 极速  （默认是普通）
	RollbackMode *string `json:"RollbackMode,omitnil,omitempty" name:"RollbackMode"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type RollbackData struct {
	// 实例CPU
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例内存
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 集群存储上限
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// 原集群id
	OriginalClusterId *string `json:"OriginalClusterId,omitnil,omitempty" name:"OriginalClusterId"`

	// 原集群名
	OriginalClusterName *string `json:"OriginalClusterName,omitnil,omitempty" name:"OriginalClusterName"`

	// 回档方式
	RollbackStrategy *string `json:"RollbackStrategy,omitnil,omitempty" name:"RollbackStrategy"`

	// 快照时间
	SnapshotTime *string `json:"SnapshotTime,omitnil,omitempty" name:"SnapshotTime"`

	// 回档到 Serverless 集群时最小 CPU
	MinCpu *int64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// 回档到 Serverless 集群时最大 CPU
	MaxCpu *int64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// 快照ID
	SnapShotId *uint64 `json:"SnapShotId,omitnil,omitempty" name:"SnapShotId"`

	// 回档数据库
	// 注意：此字段可能返回 null，表示取不到有效值。
	RollbackDatabases []*RollbackDatabase `json:"RollbackDatabases,omitnil,omitempty" name:"RollbackDatabases"`

	// 回档数据表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RollbackTables []*RollbackTable `json:"RollbackTables,omitnil,omitempty" name:"RollbackTables"`

	// 备份文件名称
	BackupFileName *string `json:"BackupFileName,omitnil,omitempty" name:"BackupFileName"`

	// 回档进程
	RollbackProcess *RollbackProcessInfo `json:"RollbackProcess,omitnil,omitempty" name:"RollbackProcess"`
}

type RollbackDatabase struct {
	// 旧数据库名称
	OldDatabase *string `json:"OldDatabase,omitnil,omitempty" name:"OldDatabase"`

	// 新数据库名称
	NewDatabase *string `json:"NewDatabase,omitnil,omitempty" name:"NewDatabase"`
}

type RollbackInstanceInfo struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// vpc信息
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 子网信息
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// vip信息
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// vport信息
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// cpu大小
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存大小
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 存储大小
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`
}

type RollbackProcessInfo struct {
	// 是否可以交换vip
	IsVipSwitchable *bool `json:"IsVipSwitchable,omitnil,omitempty" name:"IsVipSwitchable"`

	// vip可交换时间
	VipSwitchableTime *string `json:"VipSwitchableTime,omitnil,omitempty" name:"VipSwitchableTime"`

	// 交换实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExchangeInstanceInfoList []*ExchangeInstanceInfo `json:"ExchangeInstanceInfoList,omitnil,omitempty" name:"ExchangeInstanceInfoList"`

	// 交换RO组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExchangeRoGroupInfoList []*ExchangeRoGroupInfo `json:"ExchangeRoGroupInfoList,omitnil,omitempty" name:"ExchangeRoGroupInfoList"`

	// 当前步骤
	CurrentStep *string `json:"CurrentStep,omitnil,omitempty" name:"CurrentStep"`

	// 当前步骤进度
	CurrentStepProgress *int64 `json:"CurrentStepProgress,omitnil,omitempty" name:"CurrentStepProgress"`

	// 当前步骤剩余时间
	CurrentStepRemainingTime *string `json:"CurrentStepRemainingTime,omitnil,omitempty" name:"CurrentStepRemainingTime"`
}

type RollbackRoGroupInfo struct {
	// 实例组ID
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// vpc信息
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 子网信息
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// vip信息
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// vport信息
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`
}

type RollbackTable struct {
	// 数据库名称
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 数据库表
	Tables []*RollbackTableInfo `json:"Tables,omitnil,omitempty" name:"Tables"`
}

type RollbackTableInfo struct {
	// 旧表名称
	OldTable *string `json:"OldTable,omitnil,omitempty" name:"OldTable"`

	// 新表名称
	NewTable *string `json:"NewTable,omitnil,omitempty" name:"NewTable"`
}

type RollbackTimeRange struct {
	// 开始时间
	TimeRangeStart *string `json:"TimeRangeStart,omitnil,omitempty" name:"TimeRangeStart"`

	// 结束时间
	TimeRangeEnd *string `json:"TimeRangeEnd,omitnil,omitempty" name:"TimeRangeEnd"`
}

// Predefined struct for user
type RollbackToNewClusterRequestParams struct {
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 回档时，传入源集群ID，用于查找源poolId
	OriginalClusterId *string `json:"OriginalClusterId,omitnil,omitempty" name:"OriginalClusterId"`

	// 所属VPC网络ID
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 所属子网ID
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 集群名称，长度小于64个字符，每个字符取值范围：大/小写字母，数字，特殊符号（'-','_','.'）
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 快照回档，表示snapshotId；时间点回档，表示queryId，为0，表示需要判断时间点是否有效
	RollbackId *uint64 `json:"RollbackId,omitnil,omitempty" name:"RollbackId"`

	// 时间点回档，指定时间；快照回档，快照时间
	ExpectTime *string `json:"ExpectTime,omitnil,omitempty" name:"ExpectTime"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 集群创建需要绑定的tag数组信息
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Db类型
	// 当DbType为MYSQL时可选(默认NORMAL)：
	// <li>NORMAL</li>
	// <li>SERVERLESS</li>
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// 当DbMode为SEVERLESS时必填
	// cpu最小值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// 当DbMode为SEVERLESS时必填：
	// cpu最大值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// 当DbMode为SEVERLESS时，指定集群是否自动暂停，可选范围
	// <li>yes</li>
	// <li>no</li>
	// 默认值:yes
	AutoPause *string `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// 当DbMode为SEVERLESS时，指定集群自动暂停的延迟，单位秒，可选范围[600,691200]
	// 默认值:600
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitnil,omitempty" name:"AutoPauseDelay"`

	// 安全组id数组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 告警策略Id数组
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitnil,omitempty" name:"AlarmPolicyIds"`

	// 参数数组，暂时支持character_set_server （utf8｜latin1｜gbk｜utf8mb4） ，lower_case_table_names，1-大小写不敏感，0-大小写敏感
	ClusterParams []*ParamItem `json:"ClusterParams,omitnil,omitempty" name:"ClusterParams"`

	// 参数模板ID，可以通过查询参数模板信息DescribeParamTemplates获得参数模板ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 实例初始化配置信息，主要用于购买集群时选不同规格实例
	InstanceInitInfos []*InstanceInitInfo `json:"InstanceInitInfos,omitnil,omitempty" name:"InstanceInitInfos"`

	// 0-下单并支付 1-下单
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// 计算节点付费模式：0-按量计费，1-预付费
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 时间
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 单位
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 回档库信息
	RollbackDatabases []*RollbackDatabase `json:"RollbackDatabases,omitnil,omitempty" name:"RollbackDatabases"`

	// 回档表信息
	RollbackTables []*RollbackTable `json:"RollbackTables,omitnil,omitempty" name:"RollbackTables"`

	// 原ro实例信息
	OriginalROInstanceList []*string `json:"OriginalROInstanceList,omitnil,omitempty" name:"OriginalROInstanceList"`

	// 项目id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否开启归档，可选范围<li>yes</li><li>no</li>默认值:yes
	AutoArchive *string `json:"AutoArchive,omitnil,omitempty" name:"AutoArchive"`
}

type RollbackToNewClusterRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 回档时，传入源集群ID，用于查找源poolId
	OriginalClusterId *string `json:"OriginalClusterId,omitnil,omitempty" name:"OriginalClusterId"`

	// 所属VPC网络ID
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 所属子网ID
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 集群名称，长度小于64个字符，每个字符取值范围：大/小写字母，数字，特殊符号（'-','_','.'）
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 快照回档，表示snapshotId；时间点回档，表示queryId，为0，表示需要判断时间点是否有效
	RollbackId *uint64 `json:"RollbackId,omitnil,omitempty" name:"RollbackId"`

	// 时间点回档，指定时间；快照回档，快照时间
	ExpectTime *string `json:"ExpectTime,omitnil,omitempty" name:"ExpectTime"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 集群创建需要绑定的tag数组信息
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Db类型
	// 当DbType为MYSQL时可选(默认NORMAL)：
	// <li>NORMAL</li>
	// <li>SERVERLESS</li>
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// 当DbMode为SEVERLESS时必填
	// cpu最小值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// 当DbMode为SEVERLESS时必填：
	// cpu最大值，可选范围参考DescribeServerlessInstanceSpecs接口返回
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// 当DbMode为SEVERLESS时，指定集群是否自动暂停，可选范围
	// <li>yes</li>
	// <li>no</li>
	// 默认值:yes
	AutoPause *string `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// 当DbMode为SEVERLESS时，指定集群自动暂停的延迟，单位秒，可选范围[600,691200]
	// 默认值:600
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitnil,omitempty" name:"AutoPauseDelay"`

	// 安全组id数组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 告警策略Id数组
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitnil,omitempty" name:"AlarmPolicyIds"`

	// 参数数组，暂时支持character_set_server （utf8｜latin1｜gbk｜utf8mb4） ，lower_case_table_names，1-大小写不敏感，0-大小写敏感
	ClusterParams []*ParamItem `json:"ClusterParams,omitnil,omitempty" name:"ClusterParams"`

	// 参数模板ID，可以通过查询参数模板信息DescribeParamTemplates获得参数模板ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 实例初始化配置信息，主要用于购买集群时选不同规格实例
	InstanceInitInfos []*InstanceInitInfo `json:"InstanceInitInfos,omitnil,omitempty" name:"InstanceInitInfos"`

	// 0-下单并支付 1-下单
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// 计算节点付费模式：0-按量计费，1-预付费
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 时间
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 单位
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 回档库信息
	RollbackDatabases []*RollbackDatabase `json:"RollbackDatabases,omitnil,omitempty" name:"RollbackDatabases"`

	// 回档表信息
	RollbackTables []*RollbackTable `json:"RollbackTables,omitnil,omitempty" name:"RollbackTables"`

	// 原ro实例信息
	OriginalROInstanceList []*string `json:"OriginalROInstanceList,omitnil,omitempty" name:"OriginalROInstanceList"`

	// 项目id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否开启归档，可选范围<li>yes</li><li>no</li>默认值:yes
	AutoArchive *string `json:"AutoArchive,omitnil,omitempty" name:"AutoArchive"`
}

func (r *RollbackToNewClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackToNewClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "OriginalClusterId")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "ClusterName")
	delete(f, "RollbackId")
	delete(f, "ExpectTime")
	delete(f, "AutoVoucher")
	delete(f, "ResourceTags")
	delete(f, "DbMode")
	delete(f, "MinCpu")
	delete(f, "MaxCpu")
	delete(f, "AutoPause")
	delete(f, "AutoPauseDelay")
	delete(f, "SecurityGroupIds")
	delete(f, "AlarmPolicyIds")
	delete(f, "ClusterParams")
	delete(f, "ParamTemplateId")
	delete(f, "InstanceInitInfos")
	delete(f, "DealMode")
	delete(f, "PayMode")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "RollbackDatabases")
	delete(f, "RollbackTables")
	delete(f, "OriginalROInstanceList")
	delete(f, "ProjectId")
	delete(f, "AutoArchive")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackToNewClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackToNewClusterResponseParams struct {
	// 冻结流水ID
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// 订单号
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 资源ID列表（该字段已不再维护，请使用dealNames字段查询接口DescribeResourcesByDealName获取资源ID）
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 集群ID列表（该字段已不再维护，请使用dealNames字段查询接口DescribeResourcesByDealName获取集群ID）
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 大订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RollbackToNewClusterResponse struct {
	*tchttp.BaseResponse
	Response *RollbackToNewClusterResponseParams `json:"Response"`
}

func (r *RollbackToNewClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackToNewClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleFilters struct {
	// 审计规则过滤条件的参数名称。可选值：host – 客户端 IP；user – 数据库账户；dbName – 数据库名称；sqlType-SQL类型；sql-sql语句；affectRows -影响行数；sentRows-返回行数；checkRows-扫描行数；execTime-执行时间。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 审计规则过滤条件的匹配类型。可选值：INC – 包含；EXC – 不包含；EQS – 等于；NEQ – 不等于；REG-正则；GT-大于；LT-小于。
	Compare *string `json:"Compare,omitnil,omitempty" name:"Compare"`

	// 审计规则过滤条件的匹配值。
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type RuleTemplateInfo struct {
	// 规则模板ID。
	RuleTemplateId *string `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// 规则模板名称。
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// 规则内容。
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// 告警等级。1-低风险，2-中风险，3-高风险。
	AlarmLevel *int64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警策略。0-不告警，1-告警。
	AlarmPolicy *int64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`

	// 规则描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type SalePackageSpec struct {
	// 资源包使用地域
	PackageRegion *string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// 资源包类型
	// CCU-计算资源包
	// DISK-存储资源包
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 资源包版本
	// base-基础版本，common-通用版本，enterprise-企业版本
	PackageVersion *string `json:"PackageVersion,omitnil,omitempty" name:"PackageVersion"`

	// 当前版本资源包最小资源数，计算资源单位：个；存储资源：GB
	MinPackageSpec *float64 `json:"MinPackageSpec,omitnil,omitempty" name:"MinPackageSpec"`

	// 当前版本资源包最大资源数，计算资源单位：个；存储资源：GB
	MaxPackageSpec *float64 `json:"MaxPackageSpec,omitnil,omitempty" name:"MaxPackageSpec"`

	// 资源包有效期，单位:天
	ExpireDay *int64 `json:"ExpireDay,omitnil,omitempty" name:"ExpireDay"`
}

type SaleRegion struct {
	// 地域英文名
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 地域数字ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域中文名
	RegionZh *string `json:"RegionZh,omitnil,omitempty" name:"RegionZh"`

	// 可售卖可用区列表
	ZoneSet []*SaleZone `json:"ZoneSet,omitnil,omitempty" name:"ZoneSet"`

	// 引擎类型
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 地域模块支持情况
	Modules []*Module `json:"Modules,omitnil,omitempty" name:"Modules"`
}

type SaleZone struct {
	// 可用区英文名
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 可用区数字ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 可用区中文名
	ZoneZh *string `json:"ZoneZh,omitnil,omitempty" name:"ZoneZh"`

	// 是否支持serverless集群<br>
	// 0:不支持<br>
	// 1:支持
	IsSupportServerless *int64 `json:"IsSupportServerless,omitnil,omitempty" name:"IsSupportServerless"`

	// 是否支持普通集群<br>
	// 0:不支持<br>
	// 1:支持
	IsSupportNormal *int64 `json:"IsSupportNormal,omitnil,omitempty" name:"IsSupportNormal"`

	// 物理区
	PhysicalZone *string `json:"PhysicalZone,omitnil,omitempty" name:"PhysicalZone"`

	// 用户是否有可用区权限
	HasPermission *bool `json:"HasPermission,omitnil,omitempty" name:"HasPermission"`

	// 是否为全链路RDMA可用区
	IsWholeRdmaZone *string `json:"IsWholeRdmaZone,omitnil,omitempty" name:"IsWholeRdmaZone"`

	// 当前可用区是否允许新购集群，1:允许，0:不允许
	IsSupportCreateCluster *int64 `json:"IsSupportCreateCluster,omitnil,omitempty" name:"IsSupportCreateCluster"`
}

// Predefined struct for user
type SearchClusterDatabasesRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 是否精确搜索。
	// 0: 模糊搜索 1:精确搜索 
	// 默认为0
	MatchType *int64 `json:"MatchType,omitnil,omitempty" name:"MatchType"`
}

type SearchClusterDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 是否精确搜索。
	// 0: 模糊搜索 1:精确搜索 
	// 默认为0
	MatchType *int64 `json:"MatchType,omitnil,omitempty" name:"MatchType"`
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
	Databases []*string `json:"Databases,omitnil,omitempty" name:"Databases"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 数据表名
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 数据表类型：
	// view：只返回 view，
	// base_table： 只返回基本表，
	// all：返回 view 和表
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`
}

type SearchClusterTablesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 数据表名
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 数据表类型：
	// view：只返回 view，
	// base_table： 只返回基本表，
	// all：返回 view 和表
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`
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
	Tables []*DatabaseTables `json:"Tables,omitnil,omitempty" name:"Tables"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 创建时间，时间格式：yyyy-mm-dd hh:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 入站规则
	Inbound []*PolicyRule `json:"Inbound,omitnil,omitempty" name:"Inbound"`

	// 出站规则
	Outbound []*PolicyRule `json:"Outbound,omitnil,omitempty" name:"Outbound"`

	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 安全组名称
	SecurityGroupName *string `json:"SecurityGroupName,omitnil,omitempty" name:"SecurityGroupName"`

	// 安全组备注
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil,omitempty" name:"SecurityGroupRemark"`
}

type ServerlessSpec struct {
	// cpu最小值
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// cpu最大值
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// 最大存储空间
	MaxStorageSize *int64 `json:"MaxStorageSize,omitnil,omitempty" name:"MaxStorageSize"`

	// 是否为默认规格
	IsDefault *int64 `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 是否有库存
	HasStock *bool `json:"HasStock,omitnil,omitempty" name:"HasStock"`

	// 库存数量
	StockCount *int64 `json:"StockCount,omitnil,omitempty" name:"StockCount"`

	// 可用区库存信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneStockInfos []*ServerlessZoneStockInfo `json:"ZoneStockInfos,omitnil,omitempty" name:"ZoneStockInfos"`
}

type ServerlessZoneStockInfo struct {
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 存储量
	StockCount *int64 `json:"StockCount,omitnil,omitempty" name:"StockCount"`

	// 是否包含库存
	HasStock *bool `json:"HasStock,omitnil,omitempty" name:"HasStock"`

	// 从可用区库存信息
	SlaveZoneStockInfos []*SlaveZoneStockInfo `json:"SlaveZoneStockInfos,omitnil,omitempty" name:"SlaveZoneStockInfos"`
}

// Predefined struct for user
type SetRenewFlagRequestParams struct {
	// 需操作的集群ID
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 自动续费标志位，续费标记 0:正常续费  1:自动续费 2:到期不续
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

type SetRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 需操作的集群ID
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 自动续费标志位，续费标记 0:正常续费  1:自动续费 2:到期不续
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
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
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type SlaveZoneAttrItem struct {
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// binlog同步方式
	BinlogSyncWay *string `json:"BinlogSyncWay,omitnil,omitempty" name:"BinlogSyncWay"`

	// 半同步超时时间，单位ms
	SemiSyncTimeout *int64 `json:"SemiSyncTimeout,omitnil,omitempty" name:"SemiSyncTimeout"`
}

type SlaveZoneStockInfo struct {
	// 备可用区
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// 备可用区的库存数量	
	StockCount *uint64 `json:"StockCount,omitnil,omitempty" name:"StockCount"`

	// 备可用区是否有库存	
	HasStock *bool `json:"HasStock,omitnil,omitempty" name:"HasStock"`
}

type SlowQueriesItem struct {
	// 执行时间戳
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 执行时长，单位秒
	QueryTime *float64 `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// sql语句
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// 客户端host
	UserHost *string `json:"UserHost,omitnil,omitempty" name:"UserHost"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 锁时长，单位秒
	LockTime *float64 `json:"LockTime,omitnil,omitempty" name:"LockTime"`

	// 扫描行数
	RowsExamined *int64 `json:"RowsExamined,omitnil,omitempty" name:"RowsExamined"`

	// 返回行数
	RowsSent *int64 `json:"RowsSent,omitnil,omitempty" name:"RowsSent"`

	// sql模板
	SqlTemplate *string `json:"SqlTemplate,omitnil,omitempty" name:"SqlTemplate"`

	// sql语句md5
	SqlMd5 *string `json:"SqlMd5,omitnil,omitempty" name:"SqlMd5"`

	// 远程读取次数
	// 数据库内核版本大于3.1.12
	SyncReadCountRemote *int64 `json:"SyncReadCountRemote,omitnil,omitempty" name:"SyncReadCountRemote"`

	// 远程读取的字节数
	// 数据库内核版本大于3.1.12
	SyncReadBytesRemote *int64 `json:"SyncReadBytesRemote,omitnil,omitempty" name:"SyncReadBytesRemote"`

	// 远程读取所花费的时间（微秒）
	// 数据库内核版本大于3.1.12
	SyncReadTimeRemote *int64 `json:"SyncReadTimeRemote,omitnil,omitempty" name:"SyncReadTimeRemote"`

	// 远程写入次数
	// 数据库内核版本大于3.1.12
	SyncWriteCountRemote *int64 `json:"SyncWriteCountRemote,omitnil,omitempty" name:"SyncWriteCountRemote"`

	// 远程写入的字节数。
	// 数据库内核版本大于3.1.12
	SyncWriteBytesRemote *int64 `json:"SyncWriteBytesRemote,omitnil,omitempty" name:"SyncWriteBytesRemote"`

	// 远程写入所花费的时间（微秒）。
	// 数据库内核版本大于3.1.12
	SyncWriteTimeRemote *int64 `json:"SyncWriteTimeRemote,omitnil,omitempty" name:"SyncWriteTimeRemote"`

	// 事务提交延迟（微秒）
	// 数据库内核版本大于3.1.12
	TrxCommitDelay *int64 `json:"TrxCommitDelay,omitnil,omitempty" name:"TrxCommitDelay"`
}

// Predefined struct for user
type StartCLSDeliveryRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开通的日志主题id
	CLSTopicIds []*string `json:"CLSTopicIds,omitnil,omitempty" name:"CLSTopicIds"`

	// 日志类型
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 是否维护时间运行
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type StartCLSDeliveryRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开通的日志主题id
	CLSTopicIds []*string `json:"CLSTopicIds,omitnil,omitempty" name:"CLSTopicIds"`

	// 日志类型
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 是否维护时间运行
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

func (r *StartCLSDeliveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartCLSDeliveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CLSTopicIds")
	delete(f, "LogType")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartCLSDeliveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartCLSDeliveryResponseParams struct {
	// 异步任务id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartCLSDeliveryResponse struct {
	*tchttp.BaseResponse
	Response *StartCLSDeliveryResponseParams `json:"Response"`
}

func (r *StartCLSDeliveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartCLSDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCLSDeliveryRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志主题id
	CLSTopicIds []*string `json:"CLSTopicIds,omitnil,omitempty" name:"CLSTopicIds"`

	// 日志类型
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 是否维护时间运行
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type StopCLSDeliveryRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志主题id
	CLSTopicIds []*string `json:"CLSTopicIds,omitnil,omitempty" name:"CLSTopicIds"`

	// 日志类型
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 是否维护时间运行
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

func (r *StopCLSDeliveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCLSDeliveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CLSTopicIds")
	delete(f, "LogType")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopCLSDeliveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCLSDeliveryResponseParams struct {
	// 异步任务id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopCLSDeliveryResponse struct {
	*tchttp.BaseResponse
	Response *StopCLSDeliveryResponseParams `json:"Response"`
}

func (r *StopCLSDeliveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCLSDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchClusterLogBin struct {
	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type SwitchClusterVpcRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 字符串vpc id
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 字符串子网id
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 旧地址回收时间
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitnil,omitempty" name:"OldIpReserveHours"`
}

type SwitchClusterVpcRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 字符串vpc id
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 字符串子网id
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 旧地址回收时间
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitnil,omitempty" name:"OldIpReserveHours"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 当前可用区
	OldZone *string `json:"OldZone,omitnil,omitempty" name:"OldZone"`

	// 要切换到的可用区
	NewZone *string `json:"NewZone,omitnil,omitempty" name:"NewZone"`

	// 维护期间执行-yes,立即执行-no
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type SwitchClusterZoneRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 当前可用区
	OldZone *string `json:"OldZone,omitnil,omitempty" name:"OldZone"`

	// 要切换到的可用区
	NewZone *string `json:"NewZone,omitnil,omitempty" name:"NewZone"`

	// 维护期间执行-yes,立即执行-no
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 字符串vpc id
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 字符串子网id
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 旧地址回收时间
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitnil,omitempty" name:"OldIpReserveHours"`

	// 数据库代理组Id（该参数为必填项，可以通过DescribeProxies接口获得）
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

type SwitchProxyVpcRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 字符串vpc id
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 字符串子网id
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 旧地址回收时间
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitnil,omitempty" name:"OldIpReserveHours"`

	// 数据库代理组Id（该参数为必填项，可以通过DescribeProxies接口获得）
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 权限列表
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

type Tag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TaskMaintainInfo struct {
	// 执行开始时间(距离0点的秒数)
	MaintainStartTime *int64 `json:"MaintainStartTime,omitnil,omitempty" name:"MaintainStartTime"`

	// 持续的时间(单位：秒)
	MaintainDuration *int64 `json:"MaintainDuration,omitnil,omitempty" name:"MaintainDuration"`

	// 可以执行的时间，枚举值：["Mon","Tue","Wed","Thu","Fri", "Sat", "Sun"]
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitnil,omitempty" name:"MaintainWeekDays"`
}

type TaskProgressInfo struct {
	// 当前步骤
	CurrentStep *string `json:"CurrentStep,omitnil,omitempty" name:"CurrentStep"`

	// 当前进度
	CurrentStepProgress *int64 `json:"CurrentStepProgress,omitnil,omitempty" name:"CurrentStepProgress"`

	// 预估时间
	CurrentStepRemainingTime *string `json:"CurrentStepRemainingTime,omitnil,omitempty" name:"CurrentStepRemainingTime"`
}

type TemplateParamInfo struct {
	// 当前值
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 默认值
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// 参数类型为enum时可选的值类型集合
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// 参数类型为float/integer时的最大值
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// 参数类型为float/integer时的最小值
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// 参数名称
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 是否需要重启
	NeedReboot *int64 `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// 参数描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 参数类型，integer/float/string/enum
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`
}

type TradePrice struct {
	// 预付费模式下资源总价，不包含优惠，单位:分
	TotalPrice *int64 `json:"TotalPrice,omitnil,omitempty" name:"TotalPrice"`

	// 总的折扣，100表示100%不打折
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 预付费模式下的优惠后总价, 单位: 分,例如用户享有折扣 =TotalPrice × Discount
	TotalPriceDiscount *int64 `json:"TotalPriceDiscount,omitnil,omitempty" name:"TotalPriceDiscount"`

	// 后付费模式下的单位资源价格，不包含优惠，单位:分
	UnitPrice *int64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// 优惠后后付费模式下的单位资源价格, 单位: 分,例如用户享有折扣=UnitPricet × Discount
	UnitPriceDiscount *int64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`

	// 计费价格单位
	ChargeUnit *string `json:"ChargeUnit,omitnil,omitempty" name:"ChargeUnit"`

	// 高精度下不包含优惠价格
	UnitPriceHighPrecision *string `json:"UnitPriceHighPrecision,omitnil,omitempty" name:"UnitPriceHighPrecision"`

	// 高精度下优惠后价格
	UnitPriceDiscountHighPrecision *string `json:"UnitPriceDiscountHighPrecision,omitnil,omitempty" name:"UnitPriceDiscountHighPrecision"`

	// 货币单位
	AmountUnit *string `json:"AmountUnit,omitnil,omitempty" name:"AmountUnit"`
}

// Predefined struct for user
type UnbindClusterResourcePackagesRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 资源包唯一ID,如果不传，解绑该实例绑定的所有资源包
	PackageIds []*string `json:"PackageIds,omitnil,omitempty" name:"PackageIds"`
}

type UnbindClusterResourcePackagesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 资源包唯一ID,如果不传，解绑该实例绑定的所有资源包
	PackageIds []*string `json:"PackageIds,omitnil,omitempty" name:"PackageIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 内核版本
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// 升级时间类型，可选：upgradeImmediate,upgradeInMaintain
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`
}

type UpgradeClusterVersionRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 内核版本
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// 升级时间类型，可选：upgradeImmediate,upgradeInMaintain
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库CPU
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 数据库内存，单位GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 升级类型：upgradeImmediate，upgradeInMaintain
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// 实例机器类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 该参数已废弃
	StorageLimit *uint64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 该参数已废弃
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 交易模式 0-下单并支付 1-下单
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// NormalUpgrade：普通变配，FastUpgrade：极速变配，若变配过程判断会造成闪断，变配流程会终止。
	UpgradeMode *string `json:"UpgradeMode,omitnil,omitempty" name:"UpgradeMode"`

	// proxy同步升级
	UpgradeProxy *UpgradeProxy `json:"UpgradeProxy,omitnil,omitempty" name:"UpgradeProxy"`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库CPU
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 数据库内存，单位GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 升级类型：upgradeImmediate，upgradeInMaintain
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// 实例机器类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 该参数已废弃
	StorageLimit *uint64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 该参数已废弃
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 交易模式 0-下单并支付 1-下单
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// NormalUpgrade：普通变配，FastUpgrade：极速变配，若变配过程判断会造成闪断，变配流程会终止。
	UpgradeMode *string `json:"UpgradeMode,omitnil,omitempty" name:"UpgradeMode"`

	// proxy同步升级
	UpgradeProxy *UpgradeProxy `json:"UpgradeProxy,omitnil,omitempty" name:"UpgradeProxy"`
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
	delete(f, "DeviceType")
	delete(f, "StorageLimit")
	delete(f, "AutoVoucher")
	delete(f, "DbType")
	delete(f, "DealMode")
	delete(f, "UpgradeMode")
	delete(f, "UpgradeProxy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceResponseParams struct {
	// 冻结流水ID
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// 大订单号
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// 订单号
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

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

type UpgradeProxy struct {
	// cpu
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// memory
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 代理节点信息
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`

	// 重新负载均衡
	ReloadBalance *string `json:"ReloadBalance,omitnil,omitempty" name:"ReloadBalance"`
}

// Predefined struct for user
type UpgradeProxyRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// cpu核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 数据库代理组节点个数
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// 数据库代理组ID（已废弃）
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 重新负载均衡：auto（自动），manual（手动）
	ReloadBalance *string `json:"ReloadBalance,omitnil,omitempty" name:"ReloadBalance"`

	// 升级时间 ：no（升级完成时）yes（实例维护时间）
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`

	// 数据库代理节点信息
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`

	// 是否滚动升级
	IsRollUpgrade *string `json:"IsRollUpgrade,omitnil,omitempty" name:"IsRollUpgrade"`

	// 滚动升级等待时间，单位：秒
	RollUpgradeWaitingTime *int64 `json:"RollUpgradeWaitingTime,omitnil,omitempty" name:"RollUpgradeWaitingTime"`
}

type UpgradeProxyRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// cpu核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 数据库代理组节点个数
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// 数据库代理组ID（已废弃）
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 重新负载均衡：auto（自动），manual（手动）
	ReloadBalance *string `json:"ReloadBalance,omitnil,omitempty" name:"ReloadBalance"`

	// 升级时间 ：no（升级完成时）yes（实例维护时间）
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`

	// 数据库代理节点信息
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`

	// 是否滚动升级
	IsRollUpgrade *string `json:"IsRollUpgrade,omitnil,omitempty" name:"IsRollUpgrade"`

	// 滚动升级等待时间，单位：秒
	RollUpgradeWaitingTime *int64 `json:"RollUpgradeWaitingTime,omitnil,omitempty" name:"RollUpgradeWaitingTime"`
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
	delete(f, "IsRollUpgrade")
	delete(f, "RollUpgradeWaitingTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeProxyResponseParams struct {
	// 异步流程ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 异步任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库代理当前版本
	SrcProxyVersion *string `json:"SrcProxyVersion,omitnil,omitempty" name:"SrcProxyVersion"`

	// 数据库代理升级版本
	DstProxyVersion *string `json:"DstProxyVersion,omitnil,omitempty" name:"DstProxyVersion"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 升级时间 ：no（升级完成时）yes（实例维护时间）
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type UpgradeProxyVersionRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据库代理当前版本
	SrcProxyVersion *string `json:"SrcProxyVersion,omitnil,omitempty" name:"SrcProxyVersion"`

	// 数据库代理升级版本
	DstProxyVersion *string `json:"DstProxyVersion,omitnil,omitempty" name:"DstProxyVersion"`

	// 数据库代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 升级时间 ：no（升级完成时）yes（实例维护时间）
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 异步任务id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DbUserName *string `json:"DbUserName,omitnil,omitempty" name:"DbUserName"`

	// 客户端ip
	DbHost *string `json:"DbHost,omitnil,omitempty" name:"DbHost"`

	// 用户权限
	DbPrivilege *string `json:"DbPrivilege,omitnil,omitempty" name:"DbPrivilege"`
}

type ZoneStockInfo struct {
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 是否有库存
	HasStock *bool `json:"HasStock,omitnil,omitempty" name:"HasStock"`

	// 库存数量
	StockCount *int64 `json:"StockCount,omitnil,omitempty" name:"StockCount"`

	// 备可用区库存信息
	SlaveZoneStockInfos []*SlaveZoneStockInfo `json:"SlaveZoneStockInfos,omitnil,omitempty" name:"SlaveZoneStockInfos"`
}