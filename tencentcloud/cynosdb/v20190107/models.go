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
}

type AddInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cpu核数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 新增只读实例数
	ReadOnlyCount *int64 `json:"ReadOnlyCount,omitempty" name:"ReadOnlyCount"`

	// 实例组ID，在已有RO组中新增实例时使用，不传则新增RO组
	InstanceGrpId *string `json:"InstanceGrpId,omitempty" name:"InstanceGrpId"`

	// 所属VPC网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 所属子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 新增RO组时使用的Port
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 订单来源
	OrderSource *string `json:"OrderSource,omitempty" name:"OrderSource"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type BackupFileInfo struct {

	// 快照文件ID，用于回滚
	SnapshotId *uint64 `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 快照文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 快照文件大小
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// 快照备份开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 快照备份完成时间
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// 备份类型：snapshot，快照备份；timepoint，时间点备份
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// 备份方式：auto，自动备份；manual，手动备份
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// 备份文件状态：success：备份成功；fail：备份失败；creating：备份文件创建中；deleting：备份文件删除中
	BackupStatus *string `json:"BackupStatus,omitempty" name:"BackupStatus"`

	// 备份文件时间
	SnapshotTime *string `json:"SnapshotTime,omitempty" name:"SnapshotTime"`
}

type BillingResourceInfo struct {

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
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
	// <li> MYSQL可选值：5.7 </li>
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// 所属项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 普通实例Cpu核数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 普通实例内存
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 存储
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 账号密码(8-64个字符，至少包含字母、数字、字符（支持的字符：_+-&=!@#$%^*()~）中的两种)
	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`

	// 端口，默认5432
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 计费模式，按量计费：0，包年包月：1。默认按量计费。
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 购买个数，目前只支持传1（不传默认为1）
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

	// 时间点回档，指定时间允许范围
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitempty" name:"ExpectTimeThresh"`

	// 普通实例存储上限，单位GB
	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// 实例数量
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 包年包月购买时长
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 包年包月购买时长单位
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 包年包月购买是否自动续费
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateClustersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 冻结流水ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		TranId *string `json:"TranId,omitempty" name:"TranId"`

		// 订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

		// 资源ID列表（异步发货可能无法返回该字段, 强烈建议使用dealNames字段查询接口DescribeResourcesByDealName获取异步发货的资源ID）
	// 注意：此字段可能返回 null，表示取不到有效值。
		ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

		// 集群ID列表（异步发货可能不返回该字段, 强烈建议使用dealNames查询接口DescribeResourcesByDealName获取异步发货的集群ID）
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

		// 大订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
		BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CynosdbCluster struct {

	// 集群状态
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
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 引擎类型
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 用户appid
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 集群状态描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 集群创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 付费模式。0-按量计费，1-包年包月
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 截止时间
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// 集群读写vip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 集群读写vport
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 项目id
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// cynos内核版本
	CynosVersion *string `json:"CynosVersion,omitempty" name:"CynosVersion"`

	// 存储容量
	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// 续费标志
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 正在处理的任务
	ProcessingTask *string `json:"ProcessingTask,omitempty" name:"ProcessingTask"`

	// 集群的任务数组
	Tasks []*ObjectTask `json:"Tasks,omitempty" name:"Tasks"`

	// 集群绑定的tag数组
	ResourceTags []*Tag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// Db类型(NORMAL, SERVERLESS)
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// 当Db类型为SERVERLESS时，serverless集群状态，可选值:
	// resume
	// pause
	ServerlessStatus *string `json:"ServerlessStatus,omitempty" name:"ServerlessStatus"`
}

type CynosdbClusterDetail struct {

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

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

	// 数据库版本
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// 使用容量
	UsedStorage *int64 `json:"UsedStorage,omitempty" name:"UsedStorage"`

	// 读写分离Vport
	RoAddr []*Addr `json:"RoAddr,omitempty" name:"RoAddr"`

	// 实例信息
	InstanceSet []*ClusterInstanceDetail `json:"InstanceSet,omitempty" name:"InstanceSet"`

	// 付费模式
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 到期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// vip地址
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// vport端口
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 项目id
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例绑定的tag数组信息
	ResourceTags []*Tag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 当Db类型为SERVERLESS时，serverless集群状态，可选值:
	// resume
	// resuming
	// pause
	// pausing
	ServerlessStatus *string `json:"ServerlessStatus,omitempty" name:"ServerlessStatus"`
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
}

type CynosdbInstanceGrp struct {

	// appId
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
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 需要过滤的账户列表
	AccountNames []*string `json:"AccountNames,omitempty" name:"AccountNames"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据库账号列表
		AccountSet []*Account `json:"AccountSet,omitempty" name:"AccountSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeBackupListRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 备份文件列表偏移
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 备份文件列表起始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总共备份文件个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 备份文件列表
		BackupList []*BackupFileInfo `json:"BackupList,omitempty" name:"BackupList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeClusterDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群详细信息
		Detail *CynosdbClusterDetail `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeClusterInstanceGrpsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例组个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例组列表
		InstanceGrpInfoList []*CynosdbInstanceGrp `json:"InstanceGrpInfoList,omitempty" name:"InstanceGrpInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 集群列表
		ClusterSet []*CynosdbCluster `json:"ClusterSet,omitempty" name:"ClusterSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全组信息
		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例详情
		Detail *CynosdbInstanceDetail `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeInstanceSpecsRequest struct {
	*tchttp.BaseRequest

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSpecsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceSpecsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规格信息
		InstanceSpecSet []*InstanceSpec `json:"InstanceSpecSet,omitempty" name:"InstanceSpecSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// 实例状态
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

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例列表
		InstanceSet []*CynosdbInstance `json:"InstanceSet,omitempty" name:"InstanceSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeMaintainPeriodResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 维护week days
		MaintainWeekDays []*string `json:"MaintainWeekDays,omitempty" name:"MaintainWeekDays"`

		// 维护开始时间，单位秒
		MaintainStartTime *int64 `json:"MaintainStartTime,omitempty" name:"MaintainStartTime"`

		// 维护时长，单位秒
		MaintainDuration *int64 `json:"MaintainDuration,omitempty" name:"MaintainDuration"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 项目ID
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

type DescribeProjectSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全组详情
		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeResourcesByDealNameRequest struct {
	*tchttp.BaseRequest

	// 计费订单id
	DealName *string `json:"DealName,omitempty" name:"DealName"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcesByDealNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesByDealNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 计费资源id信息数组
		BillingResourceInfos []*BillingResourceInfo `json:"BillingResourceInfos,omitempty" name:"BillingResourceInfos"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeRollbackTimeRangeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 有效回归时间范围开始时间点
		TimeRangeStart *string `json:"TimeRangeStart,omitempty" name:"TimeRangeStart"`

		// 有效回归时间范围结束时间点
		TimeRangeEnd *string `json:"TimeRangeEnd,omitempty" name:"TimeRangeEnd"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeRollbackTimeValidityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type InstanceSpec struct {

	// 实例CPU，单位：核
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 实例内存，单位：GB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 实例最大可用存储，单位：GB
	MaxStorageSize *uint64 `json:"MaxStorageSize,omitempty" name:"MaxStorageSize"`

	// 实例最小可用存储，单位：GB
	MinStorageSize *uint64 `json:"MinStorageSize,omitempty" name:"MinStorageSize"`
}

type IsolateClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
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

type IsolateClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 退款订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type IsolateInstanceRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID数组
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
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

type IsolateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务流id
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 隔离实例的订单id（预付费实例）
	// 注意：此字段可能返回 null，表示取不到有效值。
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyBackupConfigRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 表示全备开始时间，[0-24*3600]， 如0:00, 1:00, 2:00 分别为 0，3600， 7200
	BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitempty" name:"BackupTimeBeg"`

	// 表示全备开始时间，[0-24*3600]， 如0:00, 1:00, 2:00 分别为 0，3600， 7200
	BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitempty" name:"BackupTimeEnd"`

	// 表示保留备份时长, 单位秒，超过该时间将被清理, 七天表示为3600*24*7=604800
	ReserveDuration *uint64 `json:"ReserveDuration,omitempty" name:"ReserveDuration"`

	// 备份频率，长度为7的数组，分别对应周一到周日的备份方式，full-全量备份，increment-增量备份
	BackupFreq []*string `json:"BackupFreq,omitempty" name:"BackupFreq"`

	// 备份方式，logic-逻辑备份，snapshot-快照备份
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
	delete(f, "BackupTimeBeg")
	delete(f, "BackupTimeEnd")
	delete(f, "ReserveDuration")
	delete(f, "BackupFreq")
	delete(f, "BackupType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupConfigRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 实例组ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 要修改的安全组ID列表，一个或者多个安全组Id组成的数组。
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMaintainPeriodConfigRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 维护开始时间，单位为秒，如3:00为10800
	MaintainStartTime *int64 `json:"MaintainStartTime,omitempty" name:"MaintainStartTime"`

	// 维护持续时间，单位为秒，如1小时为3600
	MaintainDuration *int64 `json:"MaintainDuration,omitempty" name:"MaintainDuration"`

	// 每周维护日期
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

type ModifyMaintainPeriodConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type OfflineClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务流ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type OfflineInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务流ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type QueryFilter struct {

	// 搜索字段，目前支持："InstanceId", "ProjectId", "InstanceName", "Vip"
	Names []*string `json:"Names,omitempty" name:"Names"`

	// 搜索字符串
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否精确匹配
	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`

	// 搜索字段
	Name *string `json:"Name,omitempty" name:"Name"`
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

type SetRenewFlagRequest struct {
	*tchttp.BaseRequest

	// 需操作的实例ID
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 自动续费标志位
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

type SetRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作成功实例数
		Count *int64 `json:"Count,omitempty" name:"Count"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type Tag struct {

	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库CPU
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 数据库内存
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 升级类型：upgradeImmediate，upgradeInMaintain
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`

	// 存储上限，为0表示使用标准配置
	StorageLimit *uint64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 数据库类型，取值范围: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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
