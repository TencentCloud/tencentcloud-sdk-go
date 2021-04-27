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

package v20180411

import (
    "encoding/json"
    "errors"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddShardConfig struct {

	// 新增分片的数量
	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`

	// 分片内存大小，单位 GB
	ShardMemory *int64 `json:"ShardMemory,omitempty" name:"ShardMemory"`

	// 分片存储大小，单位 GB
	ShardStorage *int64 `json:"ShardStorage,omitempty" name:"ShardStorage"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称，本接口取值：dcdb。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 要绑定的安全组ID，类似sg-efil73jd。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 被绑定的实例ID，类似tdsqlshard-lesecurk，支持指定多个实例。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return errors.New("AssociateSecurityGroupsRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BriefNodeInfo struct {

	// DB节点ID
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`

	// DB节点角色，取值为master或者slave
	Role *string `json:"Role,omitempty" name:"Role"`

	// 节点所属分片的分片ID
	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`
}

type CloneAccountRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 源用户账户名
	SrcUser *string `json:"SrcUser,omitempty" name:"SrcUser"`

	// 源用户HOST
	SrcHost *string `json:"SrcHost,omitempty" name:"SrcHost"`

	// 目的用户账户名
	DstUser *string `json:"DstUser,omitempty" name:"DstUser"`

	// 目的用户HOST
	DstHost *string `json:"DstHost,omitempty" name:"DstHost"`

	// 目的用户账户描述
	DstDesc *string `json:"DstDesc,omitempty" name:"DstDesc"`
}

func (r *CloneAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SrcUser")
	delete(f, "SrcHost")
	delete(f, "DstUser")
	delete(f, "DstHost")
	delete(f, "DstDesc")
	if len(f) > 0 {
		return errors.New("CloneAccountRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloneAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务流程ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloneAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseDBExtranetAccessRequest struct {
	*tchttp.BaseRequest

	// 待关闭外网访问的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 是否IPv6，默认0
	Ipv6Flag *int64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
}

func (r *CloseDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ipv6Flag")
	if len(f) > 0 {
		return errors.New("CloseDBExtranetAccessRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloseDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务ID，可通过 DescribeFlow 查询任务状态。
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConstraintRange struct {

	// 约束类型为section时的最小值
	Min *string `json:"Min,omitempty" name:"Min"`

	// 约束类型为section时的最大值
	Max *string `json:"Max,omitempty" name:"Max"`
}

type CopyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 源用户名
	SrcUserName *string `json:"SrcUserName,omitempty" name:"SrcUserName"`

	// 源用户允许的访问 host
	SrcHost *string `json:"SrcHost,omitempty" name:"SrcHost"`

	// 目的用户名
	DstUserName *string `json:"DstUserName,omitempty" name:"DstUserName"`

	// 目的用户允许的访问 host
	DstHost *string `json:"DstHost,omitempty" name:"DstHost"`

	// 源账号的 ReadOnly 属性
	SrcReadOnly *string `json:"SrcReadOnly,omitempty" name:"SrcReadOnly"`

	// 目的账号的 ReadOnly 属性
	DstReadOnly *string `json:"DstReadOnly,omitempty" name:"DstReadOnly"`
}

func (r *CopyAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SrcUserName")
	delete(f, "SrcHost")
	delete(f, "DstUserName")
	delete(f, "DstHost")
	delete(f, "SrcReadOnly")
	delete(f, "DstReadOnly")
	if len(f) > 0 {
		return errors.New("CopyAccountPrivilegesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CopyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// AccountName
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 可以登录的主机，与mysql 账号的 host 格式一致，可以支持通配符，例如 %，10.%，10.20.%。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 账号密码，由字母、数字或常见符号组成，不能包含分号、单引号和双引号，长度为6~32位。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 是否创建为只读账号，0：否， 1：该账号的sql请求优先选择备机执行，备机不可用时选择主机执行，2：优先选择备机执行，备机不可用时操作失败，3：只从备机读取。
	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// 账号备注，可以包含中文、英文字符、常见符号和数字，长度为0~256字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 如果备机延迟超过本参数设置值，系统将认为备机发生故障
	// 建议该参数值大于10。当ReadOnly选择1、2时该参数生效。
	DelayThresh *int64 `json:"DelayThresh,omitempty" name:"DelayThresh"`
}

func (r *CreateAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "Password")
	delete(f, "ReadOnly")
	delete(f, "Description")
	delete(f, "DelayThresh")
	if len(f) > 0 {
		return errors.New("CreateAccountRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID，透传入参。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 用户名，透传入参。
		UserName *string `json:"UserName,omitempty" name:"UserName"`

		// 允许访问的 host，透传入参。
		Host *string `json:"Host,omitempty" name:"Host"`

		// 透传入参。
		ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 分片节点可用区分布，最多可填两个可用区。当分片规格为一主两从时，其中两个节点在第一个可用区。
	// 注意当前可售卖的可用区需要通过DescribeDCDBSaleInfo接口拉取。
	Zones []*string `json:"Zones,omitempty" name:"Zones" list`

	// 欲购买的时长，单位：月。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 分片内存大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardMemory *int64 `json:"ShardMemory,omitempty" name:"ShardMemory"`

	// 分片存储空间大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardStorage *int64 `json:"ShardStorage,omitempty" name:"ShardStorage"`

	// 单个分片节点个数，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardNodeCount *int64 `json:"ShardNodeCount,omitempty" name:"ShardNodeCount"`

	// 实例分片个数，可选范围2-8，可以通过升级实例进行新增分片到最多64个分片。
	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`

	// 欲购买实例的数量
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 项目 ID，可以通过查看项目列表获取，不传则关联到默认项目
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 虚拟私有网络 ID，不传或传空表示创建为基础网络
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 虚拟私有网络子网 ID，VpcId不为空时必填
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 数据库引擎版本，当前可选：8.0.18，10.1.9，5.7.17。
	// 8.0.18 - MySQL 8.0.18；
	// 10.1.9 - Mariadb 10.1.9；
	// 5.7.17 - Percona 5.7.17。
	// 如果不填的话，默认为5.7.17，表示Percona 5.7.17。
	DbVersionId *string `json:"DbVersionId,omitempty" name:"DbVersionId"`

	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`

	// 安全组id
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例名称， 可以通过该字段自主的设置实例的名字
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 是否支持IPv6
	Ipv6Flag *int64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`

	// 标签键值对数组
	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags" list`
}

func (r *CreateDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zones")
	delete(f, "Period")
	delete(f, "ShardMemory")
	delete(f, "ShardStorage")
	delete(f, "ShardNodeCount")
	delete(f, "ShardCount")
	delete(f, "Count")
	delete(f, "ProjectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "DbVersionId")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "SecurityGroupId")
	delete(f, "InstanceName")
	delete(f, "Ipv6Flag")
	delete(f, "ResourceTags")
	if len(f) > 0 {
		return errors.New("CreateDCDBInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 长订单号。可以据此调用 DescribeOrders
	//  查询订单详细信息，或在支付失败时调用用户账号相关接口进行支付。
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// 订单对应的实例 ID 列表，如果此处没有返回实例 ID，可以通过订单查询接口获取。还可通过实例查询接口查询实例是否创建完成。
	// 注意：此字段可能返回 null，表示取不到有效值。
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBAccount struct {

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户可以从哪台主机登录（对应 MySQL 用户的 host 字段，UserName + Host 唯一标识一个用户，IP形式，IP段以%结尾；支持填入%；为空默认等于%）
	Host *string `json:"Host,omitempty" name:"Host"`

	// 用户备注信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最后更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 只读标记，0：否， 1：该账号的sql请求优先选择备机执行，备机不可用时选择主机执行，2：优先选择备机执行，备机不可用时操作失败。
	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// 如果备机延迟超过本参数设置值，系统将认为备机发生故障
	// 建议该参数值大于10。当ReadOnly选择1、2时该参数生效。
	DelayThresh *int64 `json:"DelayThresh,omitempty" name:"DelayThresh"`
}

type DBParamValue struct {

	// 参数名称
	Param *string `json:"Param,omitempty" name:"Param"`

	// 参数值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DCDBInstanceInfo struct {

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 应用ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// VPC数字ID
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet数字ID
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// 状态中文描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 内网IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 内网端口
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 自动续费标志
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 内存大小，单位 GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 存储大小，单位 GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 分片个数
	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`

	// 到期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// 隔离时间
	IsolatedTimestamp *string `json:"IsolatedTimestamp,omitempty" name:"IsolatedTimestamp"`

	// 账号ID
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 分片详情
	ShardDetail []*ShardInfo `json:"ShardDetail,omitempty" name:"ShardDetail" list`

	// 节点数，2 为一主一从， 3 为一主二从
	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// 临时实例标记，0 为非临时实例
	IsTmp *int64 `json:"IsTmp,omitempty" name:"IsTmp"`

	// 独享集群ID，为空表示非独享集群实例
	ExclusterId *string `json:"ExclusterId,omitempty" name:"ExclusterId"`

	// 字符串型的私有网络ID
	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`

	// 字符串型的私有网络子网ID
	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`

	// 数字实例ID（过时字段，请勿依赖该值）
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 外网访问的域名，公网可解析
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`

	// 外网 IP 地址，公网可访问
	WanVip *string `json:"WanVip,omitempty" name:"WanVip"`

	// 外网端口
	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`

	// 产品类型 ID（过时字段，请勿依赖该值）
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// 实例最后更新时间，格式为 2006-01-02 15:04:05
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 数据库引擎
	DbEngine *string `json:"DbEngine,omitempty" name:"DbEngine"`

	// 数据库引擎版本
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// 付费模式
	Paymode *string `json:"Paymode,omitempty" name:"Paymode"`

	// 实例处于异步任务状态时，表示异步任务流程ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Locker *int64 `json:"Locker,omitempty" name:"Locker"`

	// 外网状态，0-未开通；1-已开通；2-关闭；3-开通中
	WanStatus *int64 `json:"WanStatus,omitempty" name:"WanStatus"`

	// 该实例是否支持审计。1-支持；0-不支持
	IsAuditSupported *uint64 `json:"IsAuditSupported,omitempty" name:"IsAuditSupported"`

	// Cpu核数
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 实例IPv6标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6Flag *uint64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`

	// 内网IPv6
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vipv6 *string `json:"Vipv6,omitempty" name:"Vipv6"`

	// 外网IPv6
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanVipv6 *string `json:"WanVipv6,omitempty" name:"WanVipv6"`

	// 外网IPv6端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanPortIpv6 *uint64 `json:"WanPortIpv6,omitempty" name:"WanPortIpv6"`

	// 外网IPv6状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanStatusIpv6 *uint64 `json:"WanStatusIpv6,omitempty" name:"WanStatusIpv6"`

	// DCN标志，0-无，1-主实例，2-灾备实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	DcnFlag *int64 `json:"DcnFlag,omitempty" name:"DcnFlag"`

	// DCN状态，0-无，1-创建中，2-同步中，3-已断开
	// 注意：此字段可能返回 null，表示取不到有效值。
	DcnStatus *int64 `json:"DcnStatus,omitempty" name:"DcnStatus"`

	// DCN灾备实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DcnDstNum *int64 `json:"DcnDstNum,omitempty" name:"DcnDstNum"`

	// 1： 主实例（独享型）, 2: 主实例, 3： 灾备实例, 4： 灾备实例（独享型）
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`
}

type DCDBShardInfo struct {

	// 所属实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分片SQL透传Id，用于将sql透传到指定分片执行
	ShardSerialId *string `json:"ShardSerialId,omitempty" name:"ShardSerialId"`

	// 全局唯一的分片Id
	ShardInstanceId *string `json:"ShardInstanceId,omitempty" name:"ShardInstanceId"`

	// 状态：0 创建中，1 流程处理中， 2 运行中，3 分片未初始化
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 状态中文描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 字符串格式的私有网络Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 字符串格式的私有网络子网Id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 内存大小，单位 GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 存储大小，单位 GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 到期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// 节点数，2 为一主一从， 3 为一主二从
	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// 存储使用率，单位为 %
	StorageUsage *float64 `json:"StorageUsage,omitempty" name:"StorageUsage"`

	// 内存使用率，单位为 %
	MemoryUsage *float64 `json:"MemoryUsage,omitempty" name:"MemoryUsage"`

	// 数字分片Id（过时字段，请勿依赖该值）
	ShardId *int64 `json:"ShardId,omitempty" name:"ShardId"`

	// 产品ProductID
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// Proxy版本
	ProxyVersion *string `json:"ProxyVersion,omitempty" name:"ProxyVersion"`

	// 付费模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Paymode *string `json:"Paymode,omitempty" name:"Paymode"`

	// 分片的主可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShardMasterZone *string `json:"ShardMasterZone,omitempty" name:"ShardMasterZone"`

	// 分片的从可用区列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShardSlaveZones []*string `json:"ShardSlaveZones,omitempty" name:"ShardSlaveZones" list`

	// CPU核数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 分片ShardKey的范围（总共64个哈希值），例如： 0-31，32-63
	Range *string `json:"Range,omitempty" name:"Range"`
}

type Database struct {

	// 数据库名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

type DatabaseFunction struct {

	// 函数名称
	Func *string `json:"Func,omitempty" name:"Func"`
}

type DatabaseProcedure struct {

	// 存储过程名称
	Proc *string `json:"Proc,omitempty" name:"Proc"`
}

type DatabaseTable struct {

	// 表名
	Table *string `json:"Table,omitempty" name:"Table"`
}

type DatabaseView struct {

	// 视图名称
	View *string `json:"View,omitempty" name:"View"`
}

type Deal struct {

	// 订单号
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 所属账号
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 商品数量
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 关联的流程 Id，可用于查询流程执行状态
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 只有创建实例的订单会填充该字段，表示该订单创建的实例的 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 付费模式，0后付费/1预付费
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
}

type DeleteAccountRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户允许的访问 host
	Host *string `json:"Host,omitempty" name:"Host"`
}

func (r *DeleteAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	if len(f) > 0 {
		return errors.New("DeleteAccountRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 登录用户名。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 数据库名。如果为 \*，表示查询全局权限（即 \*.\*），此时忽略 Type 和 Object 参数
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 类型,可以填入 table 、 view 、 proc 、 func 和 \*。当 DbName 为具体数据库名，Type为 \* 时，表示查询该数据库权限（即db.\*），此时忽略 Object 参数
	Type *string `json:"Type,omitempty" name:"Type"`

	// 具体的 Type 的名称，例如 Type 为 table 时就是具体的表名。DbName 和 Type 都为具体名称，则 Object 表示具体对象名，不能为 \* 或者为空
	Object *string `json:"Object,omitempty" name:"Object"`

	// 当 Type=table 时，ColName 为 \* 表示查询表的权限，如果为具体字段名，表示查询对应字段的权限
	ColName *string `json:"ColName,omitempty" name:"ColName"`
}

func (r *DescribeAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "DbName")
	delete(f, "Type")
	delete(f, "Object")
	delete(f, "ColName")
	if len(f) > 0 {
		return errors.New("DescribeAccountPrivilegesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 权限列表。
		Privileges []*string `json:"Privileges,omitempty" name:"Privileges" list`

		// 数据库账号用户名
		UserName *string `json:"UserName,omitempty" name:"UserName"`

		// 数据库账号Host
		Host *string `json:"Host,omitempty" name:"Host"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return errors.New("DescribeAccountsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID，透传入参。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 实例用户列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Users []*DBAccount `json:"Users,omitempty" name:"Users" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBLogFilesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分片 ID，形如：shard-7noic7tv
	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`

	// 请求日志类型，取值只能为1、2、3或者4。1-binlog，2-冷备，3-errlog，4-slowlog。
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeDBLogFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBLogFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ShardId")
	delete(f, "Type")
	if len(f) > 0 {
		return errors.New("DescribeDBLogFilesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBLogFilesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例 ID，形如：dcdbt-ow728lmc。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 请求日志类型。1-binlog，2-冷备，3-errlog，4-slowlog。
		Type *uint64 `json:"Type,omitempty" name:"Type"`

		// 请求日志总数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 日志文件列表
		Files []*LogFileInfo `json:"Files,omitempty" name:"Files" list`

		// 如果是VPC网络的实例，做用本前缀加上URI为下载地址
		VpcPrefix *string `json:"VpcPrefix,omitempty" name:"VpcPrefix"`

		// 如果是普通网络的实例，做用本前缀加上URI为下载地址
		NormalPrefix *string `json:"NormalPrefix,omitempty" name:"NormalPrefix"`

		// 分片 ID，形如：shard-7noic7tv
		ShardId *string `json:"ShardId,omitempty" name:"ShardId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBLogFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBLogFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBParametersRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return errors.New("DescribeDBParametersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBParametersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例 ID，形如：dcdbt-ow7t8lmc。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 请求DB的当前参数值
		Params []*ParamDesc `json:"Params,omitempty" name:"Params" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称，本接口取值：dcdb。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return errors.New("DescribeDBSecurityGroupsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全组详情。
		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups" list`

		// 实例VIP
	// 注意：此字段可能返回 null，表示取不到有效值。
		VIP *string `json:"VIP,omitempty" name:"VIP"`

		// 实例端口
	// 注意：此字段可能返回 null，表示取不到有效值。
		VPort *int64 `json:"VPort,omitempty" name:"VPort"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSyncModeRequest struct {
	*tchttp.BaseRequest

	// 待修改同步模式的实例ID。形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSyncModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSyncModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return errors.New("DescribeDBSyncModeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSyncModeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 同步模式：0 异步，1 强同步， 2 强同步可退化
		SyncMode *int64 `json:"SyncMode,omitempty" name:"SyncMode"`

		// 是否有修改流程在执行中：1 是， 0 否。
		IsModifying *int64 `json:"IsModifying,omitempty" name:"IsModifying"`

		// 当前复制方式，0 异步，1 同步
		CurrentSyncMode *int64 `json:"CurrentSyncMode,omitempty" name:"CurrentSyncMode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSyncModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSyncModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstanceNodeInfoRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 单次最多返回多少条，取值范围为(0-100]，默认为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回数据的偏移值，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDCDBInstanceNodeInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBInstanceNodeInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeDCDBInstanceNodeInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstanceNodeInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 节点总个数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 节点信息
		NodesInfo []*BriefNodeInfo `json:"NodesInfo,omitempty" name:"NodesInfo" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBInstanceNodeInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBInstanceNodeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个实例 ID 查询。实例 ID 形如：dcdbt-2t4cf98d
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 搜索的字段名，当前支持的值有：instancename、vip、all。传 instancename 表示按实例名进行搜索；传 vip 表示按内网IP进行搜索；传 all 将会按实例ID、实例名和内网IP进行搜索。
	SearchName *string `json:"SearchName,omitempty" name:"SearchName"`

	// 搜索的关键字，支持模糊搜索。多个关键字使用换行符（'\n'）分割。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 按项目 ID 查询
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// 是否根据 VPC 网络来搜索
	IsFilterVpc *bool `json:"IsFilterVpc,omitempty" name:"IsFilterVpc"`

	// 私有网络 ID， IsFilterVpc 为 1 时有效
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络的子网 ID， IsFilterVpc 为 1 时有效
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 排序字段， projectId， createtime， instancename 三者之一
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型， desc 或者 asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 偏移量，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 10，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 1非独享集群，2独享集群， 0全部
	ExclusterType *int64 `json:"ExclusterType,omitempty" name:"ExclusterType"`

	// 标识是否使用ExclusterType字段, false不使用，true使用
	IsFilterExcluster *bool `json:"IsFilterExcluster,omitempty" name:"IsFilterExcluster"`

	// 独享集群ID
	ExclusterIds []*string `json:"ExclusterIds,omitempty" name:"ExclusterIds" list`

	// 按标签key查询
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys" list`

	// 实例类型过滤，1-独享实例，2-主实例，3-灾备实例，多个按逗号分隔
	FilterInstanceType *string `json:"FilterInstanceType,omitempty" name:"FilterInstanceType"`
}

func (r *DescribeDCDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "SearchName")
	delete(f, "SearchKey")
	delete(f, "ProjectIds")
	delete(f, "IsFilterVpc")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ExclusterType")
	delete(f, "IsFilterExcluster")
	delete(f, "ExclusterIds")
	delete(f, "TagKeys")
	delete(f, "FilterInstanceType")
	if len(f) > 0 {
		return errors.New("DescribeDCDBInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的实例数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例详细信息列表
		Instances []*DCDBInstanceInfo `json:"Instances,omitempty" name:"Instances" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBPriceRequest struct {
	*tchttp.BaseRequest

	// 欲新购实例的可用区ID。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 欲购买实例的数量，目前支持购买1-10个实例
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 欲购买的时长，单位：月。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 单个分片节点个数大小，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardNodeCount *int64 `json:"ShardNodeCount,omitempty" name:"ShardNodeCount"`

	// 分片内存大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardMemory *int64 `json:"ShardMemory,omitempty" name:"ShardMemory"`

	// 分片存储空间大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardStorage *int64 `json:"ShardStorage,omitempty" name:"ShardStorage"`

	// 实例分片个数，可选范围2-8，可以通过升级实例进行新增分片到最多64个分片。
	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`

	// 付费类型。postpaid：按量付费   prepaid：预付费
	Paymode *string `json:"Paymode,omitempty" name:"Paymode"`
}

func (r *DescribeDCDBPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBPriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "Count")
	delete(f, "Period")
	delete(f, "ShardNodeCount")
	delete(f, "ShardMemory")
	delete(f, "ShardStorage")
	delete(f, "ShardCount")
	delete(f, "Paymode")
	if len(f) > 0 {
		return errors.New("DescribeDCDBPriceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBPriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 原价，单位：分
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 实际价格，单位：分。受折扣等影响，可能和原价不同。
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBRenewalPriceRequest struct {
	*tchttp.BaseRequest

	// 待续费的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 续费时长，单位：月。不传则默认为1个月。
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *DescribeDCDBRenewalPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBRenewalPriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Period")
	if len(f) > 0 {
		return errors.New("DescribeDCDBRenewalPriceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBRenewalPriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 原价，单位：分
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 实际价格，单位：分。受折扣等影响，可能和原价不同。
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBRenewalPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBRenewalPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBSaleInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDCDBSaleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBSaleInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeDCDBSaleInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBSaleInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可售卖地域信息列表
		RegionList []*RegionInfo `json:"RegionList,omitempty" name:"RegionList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBSaleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBSaleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBShardsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分片ID列表。
	ShardInstanceIds []*string `json:"ShardInstanceIds,omitempty" name:"ShardInstanceIds" list`

	// 偏移量，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段， 目前仅支持 createtime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型， desc 或者 asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeDCDBShardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBShardsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ShardInstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return errors.New("DescribeDCDBShardsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBShardsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的分片数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 分片信息列表
		Shards []*DCDBShardInfo `json:"Shards,omitempty" name:"Shards" list`

		// 灾备标志，0-无，1-主实例，2-灾备实例
	// 注意：此字段可能返回 null，表示取不到有效值。
		DcnFlag *int64 `json:"DcnFlag,omitempty" name:"DcnFlag"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBShardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBShardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBUpgradePriceRequest struct {
	*tchttp.BaseRequest

	// 待升级的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 升级类型，取值范围: 
	// <li> ADD: 新增分片 </li> 
	//  <li> EXPAND: 升级实例中的已有分片 </li> 
	//  <li> SPLIT: 将已有分片中的数据切分到新增分片上</li>
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`

	// 新增分片配置，当UpgradeType为ADD时生效。
	AddShardConfig *AddShardConfig `json:"AddShardConfig,omitempty" name:"AddShardConfig"`

	// 扩容分片配置，当UpgradeType为EXPAND时生效。
	ExpandShardConfig *ExpandShardConfig `json:"ExpandShardConfig,omitempty" name:"ExpandShardConfig"`

	// 切分分片配置，当UpgradeType为SPLIT时生效。
	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig,omitempty" name:"SplitShardConfig"`
}

func (r *DescribeDCDBUpgradePriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBUpgradePriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UpgradeType")
	delete(f, "AddShardConfig")
	delete(f, "ExpandShardConfig")
	delete(f, "SplitShardConfig")
	if len(f) > 0 {
		return errors.New("DescribeDCDBUpgradePriceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBUpgradePriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 原价，单位：分
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 实际价格，单位：分。受折扣等影响，可能和原价不同。
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBUpgradePriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBUpgradePriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseObjectsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库名称，通过 DescribeDatabases 接口获取。
	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

func (r *DescribeDatabaseObjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseObjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DbName")
	if len(f) > 0 {
		return errors.New("DescribeDatabaseObjectsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseObjectsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 透传入参。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 数据库名称。
		DbName *string `json:"DbName,omitempty" name:"DbName"`

		// 表列表。
		Tables []*DatabaseTable `json:"Tables,omitempty" name:"Tables" list`

		// 视图列表。
		Views []*DatabaseView `json:"Views,omitempty" name:"Views" list`

		// 存储过程列表。
		Procs []*DatabaseProcedure `json:"Procs,omitempty" name:"Procs" list`

		// 函数列表。
		Funcs []*DatabaseFunction `json:"Funcs,omitempty" name:"Funcs" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabaseObjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseTableRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库名称，通过 DescribeDatabases 接口获取。
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 表名称，通过 DescribeDatabaseObjects 接口获取。
	Table *string `json:"Table,omitempty" name:"Table"`
}

func (r *DescribeDatabaseTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DbName")
	delete(f, "Table")
	if len(f) > 0 {
		return errors.New("DescribeDatabaseTableRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseTableResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例名称。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 数据库名称。
		DbName *string `json:"DbName,omitempty" name:"DbName"`

		// 表名称。
		Table *string `json:"Table,omitempty" name:"Table"`

		// 列信息。
		Cols []*TableColumn `json:"Cols,omitempty" name:"Cols" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabaseTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return errors.New("DescribeDatabasesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 该实例上的数据库列表。
		Databases []*Database `json:"Databases,omitempty" name:"Databases" list`

		// 透传入参。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrdersRequest struct {
	*tchttp.BaseRequest

	// 待查询的长订单号列表，创建实例、续费实例、扩容实例接口返回。
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`
}

func (r *DescribeOrdersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealNames")
	if len(f) > 0 {
		return errors.New("DescribeOrdersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrdersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的订单数量。
		TotalCount []*int64 `json:"TotalCount,omitempty" name:"TotalCount" list`

		// 订单信息列表。
		Deals []*Deal `json:"Deals,omitempty" name:"Deals" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrdersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称，本接口取值：dcdb。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeProjectSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return errors.New("DescribeProjectSecurityGroupsRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeProjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeProjectsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 项目列表
		Projects []*Project `json:"Projects,omitempty" name:"Projects" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShardSpecRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeShardSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShardSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeShardSpecRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShardSpecResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 按机型分类的可售卖规格列表
		SpecConfig []*SpecConfig `json:"SpecConfig,omitempty" name:"SpecConfig" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShardSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShardSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSqlLogsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// SQL日志偏移。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 拉取数量（0-10000，为0时拉取总数信息）。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSqlLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSqlLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return errors.New("DescribeSqlLogsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSqlLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 当前消息队列中的sql日志条目数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 消息队列中的sql日志起始偏移。
		StartOffset *uint64 `json:"StartOffset,omitempty" name:"StartOffset"`

		// 消息队列中的sql日志结束偏移。
		EndOffset *uint64 `json:"EndOffset,omitempty" name:"EndOffset"`

		// 返回的第一条sql日志的偏移。
		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

		// 返回的sql日志数量。
		Count *uint64 `json:"Count,omitempty" name:"Count"`

		// Sql日志列表。
		SqlItems []*SqlLogItem `json:"SqlItems,omitempty" name:"SqlItems" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSqlLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSqlLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserTasksRequest struct {
	*tchttp.BaseRequest

	// 任务的状态列表。0-任务启动中；1-任务运行中；2-任务成功；3-任务失败
	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses" list`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 任务类型列表，当前支持的任务类型有：0-回档任务；1-创建实例任务；2-扩容任务；3-迁移任务；4-删除实例任务；5-重启任务
	FlowTypes []*int64 `json:"FlowTypes,omitempty" name:"FlowTypes" list`

	// 任务的创建时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务的结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务ID的数组
	UTaskIds []*int64 `json:"UTaskIds,omitempty" name:"UTaskIds" list`

	// 每次最多返回多少条任务，取值范围为(0,100]，不传的话默认值为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回任务默认是按照创建时间降序排列，从偏移值Offset处开始返回
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUserTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Statuses")
	delete(f, "InstanceIds")
	delete(f, "FlowTypes")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "UTaskIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeUserTasksRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 任务列表
		FlowSet []*UserTaskInfo `json:"FlowSet,omitempty" name:"FlowSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称，本接口取值：dcdb。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 安全组Id。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例ID列表，一个或者多个实例Id组成的数组。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return errors.New("DisassociateSecurityGroupsRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpandShardConfig struct {

	// 分片ID数组
	ShardInstanceIds []*string `json:"ShardInstanceIds,omitempty" name:"ShardInstanceIds" list`

	// 分片内存大小，单位 GB
	ShardMemory *int64 `json:"ShardMemory,omitempty" name:"ShardMemory"`

	// 分片存储大小，单位 GB
	ShardStorage *int64 `json:"ShardStorage,omitempty" name:"ShardStorage"`
}

type FlushBinlogRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *FlushBinlogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlushBinlogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return errors.New("FlushBinlogRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type FlushBinlogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FlushBinlogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlushBinlogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GrantAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 登录用户名。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 数据库名。如果为 \*，表示查询全局权限（即 \*.\*），此时忽略 Type 和 Object 参数
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 全局权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE TEMPORARY TABLES，LOCK TABLES，EXECUTE，CREATE VIEW，SHOW VIEW，CREATE ROUTINE，ALTER ROUTINE，EVENT，TRIGGER，SHOW DATABASES 
	// 库权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE TEMPORARY TABLES，LOCK TABLES，EXECUTE，CREATE VIEW，SHOW VIEW，CREATE ROUTINE，ALTER ROUTINE，EVENT，TRIGGER 
	// 表/视图权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE VIEW，SHOW VIEW，TRIGGER 
	// 存储过程/函数权限： ALTER ROUTINE，EXECUTE 
	// 字段权限： INSERT，REFERENCES，SELECT，UPDATE
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges" list`

	// 类型,可以填入 table 、 view 、 proc 、 func 和 \*。当 DbName 为具体数据库名，Type为 \* 时，表示设置该数据库权限（即db.\*），此时忽略 Object 参数
	Type *string `json:"Type,omitempty" name:"Type"`

	// 具体的 Type 的名称，例如 Type 为 table 时就是具体的表名。DbName 和 Type 都为具体名称，则 Object 表示具体对象名，不能为 \* 或者为空
	Object *string `json:"Object,omitempty" name:"Object"`

	// 当 Type=table 时，ColName 为 \* 表示对表授权，如果为具体字段名，表示对字段授权
	ColName *string `json:"ColName,omitempty" name:"ColName"`
}

func (r *GrantAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "DbName")
	delete(f, "Privileges")
	delete(f, "Type")
	delete(f, "Object")
	delete(f, "ColName")
	if len(f) > 0 {
		return errors.New("GrantAccountPrivilegesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GrantAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GrantAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InitDCDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 待初始化的实例ID列表，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 参数列表。本接口的可选值为：character_set_server（字符集，必传），lower_case_table_names（表名大小写敏感，必传，0 - 敏感；1-不敏感），innodb_page_size（innodb数据页，默认16K），sync_mode（同步模式：0 - 异步； 1 - 强同步；2 - 强同步可退化。默认为强同步）。
	Params []*DBParamValue `json:"Params,omitempty" name:"Params" list`
}

func (r *InitDCDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitDCDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Params")
	if len(f) > 0 {
		return errors.New("InitDCDBInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InitDCDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务ID，可通过 DescribeFlow 查询任务状态。
		FlowIds []*uint64 `json:"FlowIds,omitempty" name:"FlowIds" list`

		// 透传入参。
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitDCDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitDCDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KillSessionRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 会话ID列表
	SessionId []*int64 `json:"SessionId,omitempty" name:"SessionId" list`

	// 分片ID，与ShardSerialId设置一个
	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`

	// 分片序列ID，与ShardId设置一个
	ShardSerialId *string `json:"ShardSerialId,omitempty" name:"ShardSerialId"`
}

func (r *KillSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SessionId")
	delete(f, "ShardId")
	delete(f, "ShardSerialId")
	if len(f) > 0 {
		return errors.New("KillSessionRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type KillSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *KillSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogFileInfo struct {

	// Log最后修改时间
	Mtime *uint64 `json:"Mtime,omitempty" name:"Mtime"`

	// 文件长度
	Length *uint64 `json:"Length,omitempty" name:"Length"`

	// 下载Log时用到的统一资源标识符
	Uri *string `json:"Uri,omitempty" name:"Uri"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 登录用户名。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 新的账号备注，长度 0~256。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAccountDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "Description")
	if len(f) > 0 {
		return errors.New("ModifyAccountDescriptionRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称，本接口取值：dcdb。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 要修改的安全组 ID 列表，一个或者多个安全组 ID 组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "InstanceId")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return errors.New("ModifyDBInstanceSecurityGroupsRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstancesProjectRequest struct {
	*tchttp.BaseRequest

	// 待修改的实例ID列表。实例 ID 形如：dcdbt-ow728lmc。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 要分配的项目 ID，可以通过 DescribeProjects 查询项目列表接口获取。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyDBInstancesProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstancesProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return errors.New("ModifyDBInstancesProjectRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstancesProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstancesProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstancesProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParametersRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 参数列表，每一个元素是Param和Value的组合
	Params []*DBParamValue `json:"Params,omitempty" name:"Params" list`
}

func (r *ModifyDBParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Params")
	if len(f) > 0 {
		return errors.New("ModifyDBParametersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParametersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例 ID，形如：dcdbt-ow728lmc。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 各参数修改结果
		Result []*ParamModifyResult `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBSyncModeRequest struct {
	*tchttp.BaseRequest

	// 待修改同步模式的实例ID。形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 同步模式：0 异步，1 强同步， 2 强同步可退化
	SyncMode *int64 `json:"SyncMode,omitempty" name:"SyncMode"`
}

func (r *ModifyDBSyncModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBSyncModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SyncMode")
	if len(f) > 0 {
		return errors.New("ModifyDBSyncModeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBSyncModeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务Id，可通过 DescribeFlow 查询任务状态。
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBSyncModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBSyncModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenDBExtranetAccessRequest struct {
	*tchttp.BaseRequest

	// 待开放外网访问的实例ID。形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 是否IPv6，默认0
	Ipv6Flag *int64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
}

func (r *OpenDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ipv6Flag")
	if len(f) > 0 {
		return errors.New("OpenDBExtranetAccessRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OpenDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务ID，可通过 DescribeFlow 查询任务状态。
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamConstraint struct {

	// 约束类型,如枚举enum，区间section
	Type *string `json:"Type,omitempty" name:"Type"`

	// 约束类型为enum时的可选值列表
	Enum *string `json:"Enum,omitempty" name:"Enum"`

	// 约束类型为section时的范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	Range *ConstraintRange `json:"Range,omitempty" name:"Range"`

	// 约束类型为string时的可选值列表
	String *string `json:"String,omitempty" name:"String"`
}

type ParamDesc struct {

	// 参数名字
	Param *string `json:"Param,omitempty" name:"Param"`

	// 当前参数值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 设置过的值，参数生效后，该值和value一样。未设置过就不返回该字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SetValue *string `json:"SetValue,omitempty" name:"SetValue"`

	// 系统默认值
	Default *string `json:"Default,omitempty" name:"Default"`

	// 参数限制
	Constraint *ParamConstraint `json:"Constraint,omitempty" name:"Constraint"`

	// 是否有设置过值，false:没有设置过值，true:有设置过值。
	HaveSetValue *bool `json:"HaveSetValue,omitempty" name:"HaveSetValue"`
}

type ParamModifyResult struct {

	// 修改参数名字
	Param *string `json:"Param,omitempty" name:"Param"`

	// 参数修改结果。0表示修改成功；-1表示修改失败；-2表示该参数值非法
	Code *int64 `json:"Code,omitempty" name:"Code"`
}

type Project struct {

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 资源拥有者（主账号）uin
	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 应用Id
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 项目名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 创建者uin
	CreatorUin *int64 `json:"CreatorUin,omitempty" name:"CreatorUin"`

	// 来源平台
	SrcPlat *string `json:"SrcPlat,omitempty" name:"SrcPlat"`

	// 来源AppId
	SrcAppId *int64 `json:"SrcAppId,omitempty" name:"SrcAppId"`

	// 项目状态,0正常，-1关闭。默认项目为3
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 是否默认项目，1 是，0 不是
	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`

	// 描述信息
	Info *string `json:"Info,omitempty" name:"Info"`
}

type RegionInfo struct {

	// 地域英文ID
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域数字ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 地域中文名
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 可用区列表
	ZoneList []*ZonesInfo `json:"ZoneList,omitempty" name:"ZoneList" list`

	// 可选择的主可用区和从可用区
	AvailableChoice []*ShardZoneChooseInfo `json:"AvailableChoice,omitempty" name:"AvailableChoice" list`
}

type RenewDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 待续费的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 续费时长，单位：月。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`
}

func (r *RenewDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Period")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	if len(f) > 0 {
		return errors.New("RenewDCDBInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RenewDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 长订单号。可以据此调用 DescribeOrders
	//  查询订单详细信息，或在支付失败时调用用户账号相关接口进行支付。
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 登录用户名。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 新密码，由字母、数字或常见符号组成，不能包含分号、单引号和双引号，长度为6~32位。
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "Password")
	if len(f) > 0 {
		return errors.New("ResetAccountPasswordRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceTag struct {

	// 标签键key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type SecurityGroup struct {

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 创建时间，时间格式：yyyy-mm-dd hh:mm:ss
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组名称
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// 安全组备注
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`

	// 入站规则
	Inbound []*SecurityGroupBound `json:"Inbound,omitempty" name:"Inbound" list`

	// 出站规则
	Outbound []*SecurityGroupBound `json:"Outbound,omitempty" name:"Outbound" list`
}

type SecurityGroupBound struct {

	// 策略，ACCEPT 或者 DROP
	Action *string `json:"Action,omitempty" name:"Action"`

	// 来源 IP 或 IP 段，例如192.168.0.0/16
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// 端口
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// 网络协议，支持 UDP、TCP 等
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`
}

type ShardInfo struct {

	// 分片ID
	ShardInstanceId *string `json:"ShardInstanceId,omitempty" name:"ShardInstanceId"`

	// 分片Set ID
	ShardSerialId *string `json:"ShardSerialId,omitempty" name:"ShardSerialId"`

	// 状态：0 创建中，1 流程处理中， 2 运行中，3 分片未初始化，-2 分片已删除
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	Createtime *string `json:"Createtime,omitempty" name:"Createtime"`

	// 内存大小，单位 GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 存储大小，单位 GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 分片数字ID
	ShardId *int64 `json:"ShardId,omitempty" name:"ShardId"`

	// 节点数，2 为一主一从， 3 为一主二从
	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// 产品类型 Id（过时字段，请勿依赖该值）
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// Cpu核数
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`
}

type ShardZoneChooseInfo struct {

	// 主可用区
	MasterZone *ZonesInfo `json:"MasterZone,omitempty" name:"MasterZone"`

	// 可选的从可用区
	SlaveZones []*ZonesInfo `json:"SlaveZones,omitempty" name:"SlaveZones" list`
}

type SpecConfig struct {

	// 规格机型
	Machine *string `json:"Machine,omitempty" name:"Machine"`

	// 规格列表
	SpecConfigInfos []*SpecConfigInfo `json:"SpecConfigInfos,omitempty" name:"SpecConfigInfos" list`
}

type SpecConfigInfo struct {

	// 节点个数，2 表示一主一从，3 表示一主二从
	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// 内存大小，单位 GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 数据盘规格最小值，单位 GB
	MinStorage *int64 `json:"MinStorage,omitempty" name:"MinStorage"`

	// 数据盘规格最大值，单位 GB
	MaxStorage *int64 `json:"MaxStorage,omitempty" name:"MaxStorage"`

	// 推荐的使用场景
	SuitInfo *string `json:"SuitInfo,omitempty" name:"SuitInfo"`

	// 产品类型 Id
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// 最大 Qps 值
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`

	// CPU核数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
}

type SplitShardConfig struct {

	// 分片ID数组
	ShardInstanceIds []*string `json:"ShardInstanceIds,omitempty" name:"ShardInstanceIds" list`

	// 数据切分比例
	SplitRate *int64 `json:"SplitRate,omitempty" name:"SplitRate"`

	// 分片内存大小，单位 GB
	ShardMemory *int64 `json:"ShardMemory,omitempty" name:"ShardMemory"`

	// 分片存储大小，单位 GB
	ShardStorage *int64 `json:"ShardStorage,omitempty" name:"ShardStorage"`
}

type SqlLogItem struct {

	// 本条日志在消息队列中的偏移量。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 执行本条sql的用户。
	User *string `json:"User,omitempty" name:"User"`

	// 执行本条sql的客户端IP+端口。
	Client *string `json:"Client,omitempty" name:"Client"`

	// 数据库名称。
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 执行的sql语句。
	Sql *string `json:"Sql,omitempty" name:"Sql"`

	// 返回的数据行数。
	SelectRowNum *uint64 `json:"SelectRowNum,omitempty" name:"SelectRowNum"`

	// 影响行数。
	AffectRowNum *uint64 `json:"AffectRowNum,omitempty" name:"AffectRowNum"`

	// Sql执行时间戳。
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// Sql耗时，单位为毫秒。
	TimeCostMs *uint64 `json:"TimeCostMs,omitempty" name:"TimeCostMs"`

	// Sql返回码，0为成功。
	ResultCode *uint64 `json:"ResultCode,omitempty" name:"ResultCode"`
}

type TableColumn struct {

	// 列名称
	Col *string `json:"Col,omitempty" name:"Col"`

	// 列类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

type UpgradeDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 待升级的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 升级类型，取值范围: 
	// <li> ADD: 新增分片 </li> 
	//  <li> EXPAND: 升级实例中的已有分片 </li> 
	//  <li> SPLIT: 将已有分片中的数据切分到新增分片上</li>
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`

	// 新增分片配置，当UpgradeType为ADD时生效。
	AddShardConfig *AddShardConfig `json:"AddShardConfig,omitempty" name:"AddShardConfig"`

	// 扩容分片配置，当UpgradeType为EXPAND时生效。
	ExpandShardConfig *ExpandShardConfig `json:"ExpandShardConfig,omitempty" name:"ExpandShardConfig"`

	// 切分分片配置，当UpgradeType为SPLIT时生效。
	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig,omitempty" name:"SplitShardConfig"`

	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`
}

func (r *UpgradeDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UpgradeType")
	delete(f, "AddShardConfig")
	delete(f, "ExpandShardConfig")
	delete(f, "SplitShardConfig")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	if len(f) > 0 {
		return errors.New("UpgradeDCDBInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 长订单号。可以据此调用 DescribeOrders
	//  查询订单详细信息，或在支付失败时调用用户账号相关接口进行支付。
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserTaskInfo struct {

	// 任务ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 用户账户ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 任务状态，0-任务初始化中；1-任务运行中；2-任务成功；3-任务失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 任务类型，0-实例回档；1-实例创建；2-实例扩容；3-实例迁移；4-实例删除；5-实例重启
	UserTaskType *int64 `json:"UserTaskType,omitempty" name:"UserTaskType"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务错误信息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 客户端参数
	InputData *string `json:"InputData,omitempty" name:"InputData"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
}

type ZonesInfo struct {

	// 可用区英文ID
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区数字ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 可用区中文名
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}
