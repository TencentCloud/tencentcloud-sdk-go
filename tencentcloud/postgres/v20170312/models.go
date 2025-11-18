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

package v20170312

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AccountInfo struct {
	// 实例ID，形如postgres-lnp6j617
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 账号
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 账号备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 账号状态。 1-创建中，2-正常，3-修改中，4-密码重置中，5-锁定中，-1-删除中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 账号创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 账号最后一次更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 账号密码最近一次修改时间。
	// 
	// 此字段只在2025-10-31后才生效，之前无论是否修改密码，该值统一为默认值：0000-00-00 00:00:00
	// 同时仅通过云API或者管控控制台修改密码，才会更新该字段。
	PasswordUpdateTime *string `json:"PasswordUpdateTime,omitnil,omitempty" name:"PasswordUpdateTime"`

	// 账号类型。支持normal、tencentDBSuper。normal指代普通用户，tencentDBSuper为拥有pg_tencentdb_superuser角色的账号。
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 用户账号是否启用CAM验证
	OpenCam *bool `json:"OpenCam,omitnil,omitempty" name:"OpenCam"`
}

// Predefined struct for user
type AddDBInstanceToReadOnlyGroupRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 只读组ID。可通过[DescribeReadOnlyGroups](https://cloud.tencent.com/document/api/409/52599)接口获取
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

type AddDBInstanceToReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 只读组ID。可通过[DescribeReadOnlyGroups](https://cloud.tencent.com/document/api/409/52599)接口获取
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

func (r *AddDBInstanceToReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDBInstanceToReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddDBInstanceToReadOnlyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddDBInstanceToReadOnlyGroupResponseParams struct {
	// 流程ID，FlowId等同于TaskId
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddDBInstanceToReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *AddDBInstanceToReadOnlyGroupResponseParams `json:"Response"`
}

func (r *AddDBInstanceToReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDBInstanceToReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalysisItems struct {
	// 慢SQL查询的数据库名
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 慢SQL执行的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 抽象参数之后的慢SQL
	NormalQuery *string `json:"NormalQuery,omitnil,omitempty" name:"NormalQuery"`

	// 慢SQL执行的客户端地址
	ClientAddr *string `json:"ClientAddr,omitnil,omitempty" name:"ClientAddr"`

	// 在选定时间范围内慢SQL语句执行的次数
	CallNum *uint64 `json:"CallNum,omitnil,omitempty" name:"CallNum"`

	// 在选定时间范围内，慢SQL语句执行的次数占所有慢SQL的百分比。
	CallPercent *float64 `json:"CallPercent,omitnil,omitempty" name:"CallPercent"`

	// 在选定时间范围内，慢SQL执行的总时间
	CostTime *float64 `json:"CostTime,omitnil,omitempty" name:"CostTime"`

	// 在选定时间范围内，慢SQL语句执行的总时间占所有慢SQL的比例（小数返回）
	CostPercent *float64 `json:"CostPercent,omitnil,omitempty" name:"CostPercent"`

	// 在选定时间范围内，慢SQL语句执行的耗时最短的时间（单位：ms）
	MinCostTime *float64 `json:"MinCostTime,omitnil,omitempty" name:"MinCostTime"`

	// 在选定时间范围内，慢SQL语句执行的耗时最长的时间（单位：ms）
	MaxCostTime *float64 `json:"MaxCostTime,omitnil,omitempty" name:"MaxCostTime"`

	// 在选定时间范围内，慢SQL语句执行的耗时平均时间（单位：ms）
	AvgCostTime *float64 `json:"AvgCostTime,omitnil,omitempty" name:"AvgCostTime"`

	// 在选定时间范围内，慢SQL第一条开始执行的时间
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// 在选定时间范围内，慢SQL最后一条开始执行的时间
	LastTime *string `json:"LastTime,omitnil,omitempty" name:"LastTime"`
}

type BackupDownloadRestriction struct {
	// 备份文件下载限制类型，NONE 无限制，内外网都可以下载；INTRANET 只允许内网下载；CUSTOMIZE 自定义限制下载的vpc或ip。当该参数取值为CUSTOMIZE 时，vpc或ip信息至少填写一项
	RestrictionType *string `json:"RestrictionType,omitnil,omitempty" name:"RestrictionType"`

	// vpc限制效力，ALLOW 允许；DENY 拒绝。
	VpcRestrictionEffect *string `json:"VpcRestrictionEffect,omitnil,omitempty" name:"VpcRestrictionEffect"`

	// 允许或拒绝下载备份文件的vpcId列表。
	VpcIdSet []*string `json:"VpcIdSet,omitnil,omitempty" name:"VpcIdSet"`

	// ip限制效力，ALLOW 允许；DENY 拒绝。
	IpRestrictionEffect *string `json:"IpRestrictionEffect,omitnil,omitempty" name:"IpRestrictionEffect"`

	// 允许或拒绝下载备份文件的ip列表。
	IpSet []*string `json:"IpSet,omitnil,omitempty" name:"IpSet"`
}

type BackupPlan struct {
	// 备份周期
	BackupPeriod *string `json:"BackupPeriod,omitnil,omitempty" name:"BackupPeriod"`

	// 数据备份保留时长。单位：天
	BaseBackupRetentionPeriod *uint64 `json:"BaseBackupRetentionPeriod,omitnil,omitempty" name:"BaseBackupRetentionPeriod"`

	// 开始备份的最早时间
	MinBackupStartTime *string `json:"MinBackupStartTime,omitnil,omitempty" name:"MinBackupStartTime"`

	// 开始备份的最晚时间
	MaxBackupStartTime *string `json:"MaxBackupStartTime,omitnil,omitempty" name:"MaxBackupStartTime"`

	// 备份计划ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 备份计划自定义名称。
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// 日志备份保留时长。单位：天
	LogBackupRetentionPeriod *uint64 `json:"LogBackupRetentionPeriod,omitnil,omitempty" name:"LogBackupRetentionPeriod"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 最近一次的修改时间。
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// 备份计划类型。系统默认创建的为default，自定义的为custom。
	PlanType *string `json:"PlanType,omitnil,omitempty" name:"PlanType"`

	// 备份周期类型。当前支持week、month。
	BackupPeriodType *string `json:"BackupPeriodType,omitnil,omitempty" name:"BackupPeriodType"`
}

type BackupSummary struct {
	// 实例ID。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 实例日志备份数量。
	LogBackupCount *uint64 `json:"LogBackupCount,omitnil,omitempty" name:"LogBackupCount"`

	// 实例日志备份大小。
	LogBackupSize *uint64 `json:"LogBackupSize,omitnil,omitempty" name:"LogBackupSize"`

	// 手动创建的实例数据备份数量。
	ManualBaseBackupCount *uint64 `json:"ManualBaseBackupCount,omitnil,omitempty" name:"ManualBaseBackupCount"`

	// 手动创建的实例数据备份大小。
	ManualBaseBackupSize *uint64 `json:"ManualBaseBackupSize,omitnil,omitempty" name:"ManualBaseBackupSize"`

	// 自动创建的实例数据备份数量。
	AutoBaseBackupCount *uint64 `json:"AutoBaseBackupCount,omitnil,omitempty" name:"AutoBaseBackupCount"`

	// 自动创建的实例数据备份大小。
	AutoBaseBackupSize *uint64 `json:"AutoBaseBackupSize,omitnil,omitempty" name:"AutoBaseBackupSize"`

	// 总备份数量
	TotalBackupCount *uint64 `json:"TotalBackupCount,omitnil,omitempty" name:"TotalBackupCount"`

	// 总备份大小
	TotalBackupSize *uint64 `json:"TotalBackupSize,omitnil,omitempty" name:"TotalBackupSize"`
}

type BaseBackup struct {
	// 实例ID。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 备份文件唯一标识。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 备份文件名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备份方式：physical - 物理备份、logical - 逻辑备份。
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 备份模式：automatic - 自动备份、manual - 手动备份。
	BackupMode *string `json:"BackupMode,omitnil,omitempty" name:"BackupMode"`

	// 备份任务状态。枚举值：init、running、finished、failed、canceled
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 备份集大小，单位bytes。
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 备份的开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 备份的结束时间。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 备份的过期时间。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type ClassInfo struct {
	// 规格ID
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// CPU核数
	CPU *uint64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// 内存大小，单位：MB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 该规格所支持最大存储容量，单位：GB
	MaxStorage *uint64 `json:"MaxStorage,omitnil,omitempty" name:"MaxStorage"`

	// 该规格所支持最小存储容量，单位：GB
	MinStorage *uint64 `json:"MinStorage,omitnil,omitempty" name:"MinStorage"`

	// 该规格的预估QPS
	QPS *uint64 `json:"QPS,omitnil,omitempty" name:"QPS"`
}

// Predefined struct for user
type CloneDBInstanceRequestParams struct {
	// 克隆的源实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 售卖规格码。该参数可以通过调用[DescribeClasses](https://cloud.tencent.com/document/api/409/89019)的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// 实例磁盘容量大小，设置步长限制为10。单位：GB。
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 购买时长，单位：月。
	// 
	// - 预付费：支持1,2,3,4,5,6,7,8,9,10,11,12,24,36
	// - 后付费：只支持1
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 续费标记。仅当计费模式为预付费时生效。
	// 枚举值：
	// 
	// - 0：手动续费
	// - 1：自动续费
	// 
	// 默认值：0
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 私有网络ID，形如vpc-xxxxxxxx。有效的VpcId可通过登录控制台查询；也可以调用接口 [DescribeVpcEx](https://cloud.tencent.com/document/api/215/1372) ，从接口返回中的unVpcId字段获取。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络子网ID，形如subnet-xxxxxxxx。有效的私有网络子网ID可通过登录控制台查询；也可以调用接口 [DescribeSubnets ](https://cloud.tencent.com/document/api/215/15784)，从接口返回中的unSubnetId字段获取。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 新购的实例名称，仅支持长度小于60的中文/英文/数字/"_"/"-"，不指定实例名称则默认显示"源实例名-Copy"。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例计费类型，目前支持：
	// 
	// - PREPAID：预付费，即包年包月
	// - POSTPAID_BY_HOUR：后付费，即按量计费
	// 
	// 默认值：PREPAID
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例所属安全组。该参数可以通过调用[DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808)的返回值中的SecurityGroupId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 项目ID。默认值为0，表示所属默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例需要绑定的Tag信息，默认为空；可以通过调用 [DescribeTags](https://cloud.tencent.com/document/api/651/35316) 返回值中的 Tags 字段来获取。
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 实例节点部署信息，必须填写主备节点可用区。支持多可用区部署时需要指定每个节点的部署可用区信息。
	// 可用区信息可以通过调用 [DescribeZones](https://cloud.tencent.com/document/api/409/16769) 接口的返回值中的Zone字段来获取。
	DBNodeSet []*DBNode `json:"DBNodeSet,omitnil,omitempty" name:"DBNodeSet"`

	// 是否自动使用代金券：
	// 
	// - 0：否
	// - 1：是
	// 
	// 默认值：0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表。
	VoucherIds *string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 活动ID。
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 基础备份集ID。参数BackupSetId、RecoveryTargetTime两者必须填写一项，且不能同时填写。
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// 恢复时间点。参数BackupSetId、RecoveryTargetTime两者必须填写一项，且不能同时填写。
	RecoveryTargetTime *string `json:"RecoveryTargetTime,omitnil,omitempty" name:"RecoveryTargetTime"`

	// 主从同步方式，支持： 
	// <li>Semi-sync：半同步</li>
	// <li>Async：异步</li>
	// 主实例默认值：Semi-sync
	// 只读实例默认值：Async
	SyncMode *string `json:"SyncMode,omitnil,omitempty" name:"SyncMode"`

	// 实例是否开启删除保护: true-开启删除保护；false-关闭删除保护。
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

type CloneDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 克隆的源实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 售卖规格码。该参数可以通过调用[DescribeClasses](https://cloud.tencent.com/document/api/409/89019)的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// 实例磁盘容量大小，设置步长限制为10。单位：GB。
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 购买时长，单位：月。
	// 
	// - 预付费：支持1,2,3,4,5,6,7,8,9,10,11,12,24,36
	// - 后付费：只支持1
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 续费标记。仅当计费模式为预付费时生效。
	// 枚举值：
	// 
	// - 0：手动续费
	// - 1：自动续费
	// 
	// 默认值：0
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 私有网络ID，形如vpc-xxxxxxxx。有效的VpcId可通过登录控制台查询；也可以调用接口 [DescribeVpcEx](https://cloud.tencent.com/document/api/215/1372) ，从接口返回中的unVpcId字段获取。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络子网ID，形如subnet-xxxxxxxx。有效的私有网络子网ID可通过登录控制台查询；也可以调用接口 [DescribeSubnets ](https://cloud.tencent.com/document/api/215/15784)，从接口返回中的unSubnetId字段获取。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 新购的实例名称，仅支持长度小于60的中文/英文/数字/"_"/"-"，不指定实例名称则默认显示"源实例名-Copy"。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例计费类型，目前支持：
	// 
	// - PREPAID：预付费，即包年包月
	// - POSTPAID_BY_HOUR：后付费，即按量计费
	// 
	// 默认值：PREPAID
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例所属安全组。该参数可以通过调用[DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808)的返回值中的SecurityGroupId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 项目ID。默认值为0，表示所属默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例需要绑定的Tag信息，默认为空；可以通过调用 [DescribeTags](https://cloud.tencent.com/document/api/651/35316) 返回值中的 Tags 字段来获取。
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 实例节点部署信息，必须填写主备节点可用区。支持多可用区部署时需要指定每个节点的部署可用区信息。
	// 可用区信息可以通过调用 [DescribeZones](https://cloud.tencent.com/document/api/409/16769) 接口的返回值中的Zone字段来获取。
	DBNodeSet []*DBNode `json:"DBNodeSet,omitnil,omitempty" name:"DBNodeSet"`

	// 是否自动使用代金券：
	// 
	// - 0：否
	// - 1：是
	// 
	// 默认值：0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表。
	VoucherIds *string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 活动ID。
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 基础备份集ID。参数BackupSetId、RecoveryTargetTime两者必须填写一项，且不能同时填写。
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// 恢复时间点。参数BackupSetId、RecoveryTargetTime两者必须填写一项，且不能同时填写。
	RecoveryTargetTime *string `json:"RecoveryTargetTime,omitnil,omitempty" name:"RecoveryTargetTime"`

	// 主从同步方式，支持： 
	// <li>Semi-sync：半同步</li>
	// <li>Async：异步</li>
	// 主实例默认值：Semi-sync
	// 只读实例默认值：Async
	SyncMode *string `json:"SyncMode,omitnil,omitempty" name:"SyncMode"`

	// 实例是否开启删除保护: true-开启删除保护；false-关闭删除保护。
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

func (r *CloneDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "SpecCode")
	delete(f, "Storage")
	delete(f, "Period")
	delete(f, "AutoRenewFlag")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Name")
	delete(f, "InstanceChargeType")
	delete(f, "SecurityGroupIds")
	delete(f, "ProjectId")
	delete(f, "TagList")
	delete(f, "DBNodeSet")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "ActivityId")
	delete(f, "BackupSetId")
	delete(f, "RecoveryTargetTime")
	delete(f, "SyncMode")
	delete(f, "DeletionProtection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloneDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneDBInstanceResponseParams struct {
	// 订单号。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 订单流水号。
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// 克隆出的新实例ID，当前只支持后付费返回该值。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloneDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CloneDBInstanceResponseParams `json:"Response"`
}

func (r *CloneDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseAccountCAMRequestParams struct {
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 需要关闭CAM服务的账号名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 关闭CAM后，登录该账号所需要的新密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 密码是否加密
	PasswordEncrypt *bool `json:"PasswordEncrypt,omitnil,omitempty" name:"PasswordEncrypt"`
}

type CloseAccountCAMRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 需要关闭CAM服务的账号名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 关闭CAM后，登录该账号所需要的新密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 密码是否加密
	PasswordEncrypt *bool `json:"PasswordEncrypt,omitnil,omitempty" name:"PasswordEncrypt"`
}

func (r *CloseAccountCAMRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseAccountCAMRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "UserName")
	delete(f, "Password")
	delete(f, "PasswordEncrypt")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseAccountCAMRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseAccountCAMResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseAccountCAMResponse struct {
	*tchttp.BaseResponse
	Response *CloseAccountCAMResponseParams `json:"Response"`
}

func (r *CloseAccountCAMResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseAccountCAMResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseDBExtranetAccessRequestParams struct {
	// 实例ID，可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取。形如postgres-6r233v55
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 是否关闭Ipv6外网，1：是，0：否。默认值：0。
	IsIpv6 *int64 `json:"IsIpv6,omitnil,omitempty" name:"IsIpv6"`
}

type CloseDBExtranetAccessRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取。形如postgres-6r233v55
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 是否关闭Ipv6外网，1：是，0：否。默认值：0。
	IsIpv6 *int64 `json:"IsIpv6,omitnil,omitempty" name:"IsIpv6"`
}

func (r *CloseDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "IsIpv6")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseDBExtranetAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseDBExtranetAccessResponseParams struct {
	// 流程ID，FlowId等同于TaskId
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *CloseDBExtranetAccessResponseParams `json:"Response"`
}

func (r *CloseDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 创建的账号名称。由字母（a-z, A-Z）、数字（0-9）、下划线（_）组成，以字母或（_）开头，最多63个字符。不能使用系统保留关键字，不能为postgres，且不能由pg_或tencentdb_开头
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 账号类型。当前支持normal、tencentDBSuper两个输入。normal指代普通用户，tencentDBSuper为拥有pg_tencentdb_superuser角色的账号。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 账号对应的密码。密码规则如下：
	// <li>长度8 ~ 32位，推荐使用12位以上的密码</li>
	// <li>不能以" / "开头</li>
	// <li>必须包含以下四项:</li>
	// 
	// 小写字母 a ~ z           
	// 大写字母 A ～ Z
	// 数字 0 ～ 9
	// 特殊字符 ()`~!@#$%^&*-+=_|{}[]:<>,.?/
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 账号备注。只允许英文字母、数字、下划线、中划线，以及全体汉字，限60个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 账号是否开启CAM验证
	OpenCam *bool `json:"OpenCam,omitnil,omitempty" name:"OpenCam"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 创建的账号名称。由字母（a-z, A-Z）、数字（0-9）、下划线（_）组成，以字母或（_）开头，最多63个字符。不能使用系统保留关键字，不能为postgres，且不能由pg_或tencentdb_开头
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 账号类型。当前支持normal、tencentDBSuper两个输入。normal指代普通用户，tencentDBSuper为拥有pg_tencentdb_superuser角色的账号。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 账号对应的密码。密码规则如下：
	// <li>长度8 ~ 32位，推荐使用12位以上的密码</li>
	// <li>不能以" / "开头</li>
	// <li>必须包含以下四项:</li>
	// 
	// 小写字母 a ~ z           
	// 大写字母 A ～ Z
	// 数字 0 ～ 9
	// 特殊字符 ()`~!@#$%^&*-+=_|{}[]:<>,.?/
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 账号备注。只允许英文字母、数字、下划线、中划线，以及全体汉字，限60个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 账号是否开启CAM验证
	OpenCam *bool `json:"OpenCam,omitnil,omitempty" name:"OpenCam"`
}

func (r *CreateAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "UserName")
	delete(f, "Type")
	delete(f, "Password")
	delete(f, "Remark")
	delete(f, "OpenCam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccountResponseParams `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupPlanRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 备份计划名称。
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// 创建的备份计划类型，当前仅支持month创建。
	BackupPeriodType *string `json:"BackupPeriodType,omitnil,omitempty" name:"BackupPeriodType"`

	// 备份的日期，示例是每个月的2号开启备份。
	BackupPeriod []*string `json:"BackupPeriod,omitnil,omitempty" name:"BackupPeriod"`

	// 备份开始时间，不传跟随默认备份计划。
	MinBackupStartTime *string `json:"MinBackupStartTime,omitnil,omitempty" name:"MinBackupStartTime"`

	// 备份结束时间，不传跟随默认计划。
	MaxBackupStartTime *string `json:"MaxBackupStartTime,omitnil,omitempty" name:"MaxBackupStartTime"`

	// 数据备份保留时长，单位：天。取值范围为：[0,30000)
	// BackupPeriodType为week时默认是7,为month时默认为31。
	BaseBackupRetentionPeriod *uint64 `json:"BaseBackupRetentionPeriod,omitnil,omitempty" name:"BaseBackupRetentionPeriod"`
}

type CreateBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 备份计划名称。
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// 创建的备份计划类型，当前仅支持month创建。
	BackupPeriodType *string `json:"BackupPeriodType,omitnil,omitempty" name:"BackupPeriodType"`

	// 备份的日期，示例是每个月的2号开启备份。
	BackupPeriod []*string `json:"BackupPeriod,omitnil,omitempty" name:"BackupPeriod"`

	// 备份开始时间，不传跟随默认备份计划。
	MinBackupStartTime *string `json:"MinBackupStartTime,omitnil,omitempty" name:"MinBackupStartTime"`

	// 备份结束时间，不传跟随默认计划。
	MaxBackupStartTime *string `json:"MaxBackupStartTime,omitnil,omitempty" name:"MaxBackupStartTime"`

	// 数据备份保留时长，单位：天。取值范围为：[0,30000)
	// BackupPeriodType为week时默认是7,为month时默认为31。
	BaseBackupRetentionPeriod *uint64 `json:"BaseBackupRetentionPeriod,omitnil,omitempty" name:"BaseBackupRetentionPeriod"`
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
	delete(f, "DBInstanceId")
	delete(f, "PlanName")
	delete(f, "BackupPeriodType")
	delete(f, "BackupPeriod")
	delete(f, "MinBackupStartTime")
	delete(f, "MaxBackupStartTime")
	delete(f, "BaseBackupRetentionPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupPlanResponseParams struct {
	// 备份策略的ID.
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

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
type CreateBaseBackupRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type CreateBaseBackupRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *CreateBaseBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBaseBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBaseBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBaseBackupResponseParams struct {
	// 数据备份集ID
	BaseBackupId *string `json:"BaseBackupId,omitnil,omitempty" name:"BaseBackupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBaseBackupResponse struct {
	*tchttp.BaseResponse
	Response *CreateBaseBackupResponseParams `json:"Response"`
}

func (r *CreateBaseBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBaseBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceNetworkAccessRequestParams struct {
	// 实例ID，形如：postgres-6bwgamo3。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 私有网络统一 ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 是否指定分配vip true-指定分配  false-自动分配。
	IsAssignVip *bool `json:"IsAssignVip,omitnil,omitempty" name:"IsAssignVip"`

	// 目标VIP地址。当不指定该参数，且IsAssignVip为true时，默认自动分配Vip。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

type CreateDBInstanceNetworkAccessRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如：postgres-6bwgamo3。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 私有网络统一 ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 是否指定分配vip true-指定分配  false-自动分配。
	IsAssignVip *bool `json:"IsAssignVip,omitnil,omitempty" name:"IsAssignVip"`

	// 目标VIP地址。当不指定该参数，且IsAssignVip为true时，默认自动分配Vip。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

func (r *CreateDBInstanceNetworkAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceNetworkAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "IsAssignVip")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceNetworkAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceNetworkAccessResponseParams struct {
	// 流程ID，FlowId等同于TaskId
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBInstanceNetworkAccessResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstanceNetworkAccessResponseParams `json:"Response"`
}

func (r *CreateDBInstanceNetworkAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceNetworkAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatabaseRequestParams struct {
	// 实例ID，形如postgres-6fego161。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 创建的数据库名。
	// 名称规范：由字母（a-z, A-Z）、数字（0-9）、下划线（_）组成，以字母或（_）开头，最多63个字符。不能使用系统保留关键字，不能为postgres。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据库的所有者。可通过[DescribeAccounts](https://cloud.tencent.com/document/api/409/18109)接口获取
	DatabaseOwner *string `json:"DatabaseOwner,omitnil,omitempty" name:"DatabaseOwner"`

	// 数据库的字符编码。
	// 支持的常用字符集包括：UTF8、LATIN1、LATIN2、WIN1250、WIN1251、WIN1252、KOI8R、EUC_JP、EUC_KR
	// 默认值：UTF8
	Encoding *string `json:"Encoding,omitnil,omitempty" name:"Encoding"`

	// 数据库的排序规则
	Collate *string `json:"Collate,omitnil,omitempty" name:"Collate"`

	// 数据库的字符分类
	Ctype *string `json:"Ctype,omitnil,omitempty" name:"Ctype"`
}

type CreateDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-6fego161。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 创建的数据库名。
	// 名称规范：由字母（a-z, A-Z）、数字（0-9）、下划线（_）组成，以字母或（_）开头，最多63个字符。不能使用系统保留关键字，不能为postgres。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据库的所有者。可通过[DescribeAccounts](https://cloud.tencent.com/document/api/409/18109)接口获取
	DatabaseOwner *string `json:"DatabaseOwner,omitnil,omitempty" name:"DatabaseOwner"`

	// 数据库的字符编码。
	// 支持的常用字符集包括：UTF8、LATIN1、LATIN2、WIN1250、WIN1251、WIN1252、KOI8R、EUC_JP、EUC_KR
	// 默认值：UTF8
	Encoding *string `json:"Encoding,omitnil,omitempty" name:"Encoding"`

	// 数据库的排序规则
	Collate *string `json:"Collate,omitnil,omitempty" name:"Collate"`

	// 数据库的字符分类
	Ctype *string `json:"Ctype,omitnil,omitempty" name:"Ctype"`
}

func (r *CreateDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "DatabaseName")
	delete(f, "DatabaseOwner")
	delete(f, "Encoding")
	delete(f, "Collate")
	delete(f, "Ctype")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatabaseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *CreateDatabaseResponseParams `json:"Response"`
}

func (r *CreateDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesRequestParams struct {
	// 实例所属主可用区， 如：ap-guangzhou-3；若需要支持多可用区，在DBNodeSet.N字段中进行添加主可用区和备可用区信息；
	// 可用区信息可以通过调用 [DescribeZones](https://cloud.tencent.com/document/api/409/16769) 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 售卖规格码。该参数可以通过调用[DescribeClasses](https://cloud.tencent.com/document/api/409/89019)的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// 实例磁盘容量大小，单位：GB。该参数的设置步长为10。
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 购买实例数量，取值范围：[1-10]。一次性购买支持最大数量10个，若超过该数量，可进行多次调用进行购买。
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。
	// <li>预付费：支持1,2,3,4,5,6,7,8,9,10,11,12,24,36</li>
	// <li>后付费：只支持1</li>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 实例字符集，目前只支持：
	// <li> UTF8</li>
	// <li> LATIN1</li>
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// 实例根账号用户名，具体规范如下：
	// <li>用户名需要1-16个字符，只能由字母、数字或下划线组成</li>
	// <li>不能为postgres</li>
	// <li>不能由数字和pg_开头</li>
	// <li>所有规则均不区分大小写</li>
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// 实例根账号用户名对应的密码，长度8 ~ 32位，推荐使用12位以上的密码;不能以" / "开头;
	// 必须包含以下四项，字符种类:
	// <li>小写字母： [a ~ z]</li>
	// <li>大写字母：[A ～ Z]</li>
	// <li>数字：0 - 9</li>
	// <li>特殊字符：()`~!@#$%^&*-+=_|{}[]:;'<>,.?/</li>
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// PostgreSQL大版本号（该参数当前必传），版本信息可从[DescribeDBVersions](https://cloud.tencent.com/document/api/409/89018)获取。目前支持10，11，12，13，14，15这几个大版本，详情见[内核版本概述](https://cloud.tencent.com/document/product/409/67018)。
	// 输入该参数时，会基于此大版本号创建对应的最新小版本的最新内核版本号实例。
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// PostgreSQL社区大版本+小版本号。
	// 一般场景不推荐传入该参数。如需指定，只能传当前大版本号下最新小版本号。
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// PostgreSQL内核版本号。
	// 一般场景不推荐传入该参数。如需指定，只能传当前大版本号下最新内核版本号。
	DBKernelVersion *string `json:"DBKernelVersion,omitnil,omitempty" name:"DBKernelVersion"`

	// 实例计费类型，目前支持：
	// <li>PREPAID：预付费，即包年包月</li>
	// <li>POSTPAID_BY_HOUR：后付费，即按量计费</li>
	// 默认值：PREPAID
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 私有网络ID，形如vpc-xxxxxxxx（该参数当前必传）。有效的VpcId可通过登录控制台查询；也可以调用接口 [DescribeVpcEx](https://cloud.tencent.com/document/api/215/1372) ，从接口返回中的unVpcId字段获取。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络子网ID，形如subnet-xxxxxxxx（该参数当前必传）。有效的私有网络子网ID可通过登录控制台查询；也可以调用接口 [DescribeSubnets ](https://cloud.tencent.com/document/api/215/15784)，从接口返回中的unSubnetId字段获取。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例节点部署信息，支持多可用区部署时需要指定每个节点的部署可用区信息。
	// 可用区信息可以通过调用 [DescribeZones](https://cloud.tencent.com/document/api/409/16769) 接口的返回值中的Zone字段来获取。
	DBNodeSet []*DBNode `json:"DBNodeSet,omitnil,omitempty" name:"DBNodeSet"`

	// 续费标记：
	// <li>0：手动续费</li>
	// <li>1：自动续费</li>
	// 默认值：0
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 是否自动使用代金券：
	// <li>0：否</li>
	// <li>1：是</li>
	// 默认值：0
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 项目ID。默认取之为0，表示归属默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 活动ID。
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 实例名称，仅支持长度小于60的中文/英文/数字/"_"/"-"，不指定实例名称则默认显示"未命名"。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例需要绑定的Tag信息，默认为空；可以通过调用 [DescribeTags](https://cloud.tencent.com/document/api/651/35316) 返回值中的 Tags 字段来获取。
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 实例所属安全组，该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 是否需要支持数据透明加密：
	// <li>0：否</li>
	// <li>1：是</li>
	// 默认值：0
	// 参考[数据透明加密概述](https://cloud.tencent.com/document/product/409/71748)
	NeedSupportTDE *uint64 `json:"NeedSupportTDE,omitnil,omitempty" name:"NeedSupportTDE"`

	// 自定义密钥的KeyId，若选择自定义密匙加密，则需要传入自定义密匙的KeyId，KeyId是CMK的唯一标识。
	// KeyId创建获取相关参考[开启透明数据加密](https://cloud.tencent.com/document/product/409/71749)
	KMSKeyId *string `json:"KMSKeyId,omitnil,omitempty" name:"KMSKeyId"`

	// 使用KMS服务的地域，KMSRegion为空默认使用本地域的KMS，本地域不支持的情况下需自选其他KMS支持的地域。
	// KMSRegion相关介绍参考[开启透明数据加密](https://cloud.tencent.com/document/product/409/71749)
	KMSRegion *string `json:"KMSRegion,omitnil,omitempty" name:"KMSRegion"`

	// 指定KMS服务的集群，KMSClusterId为空使用默认集群的KMS，若选择指定KMS集群，则需要传入KMSClusterId。 KMSClusterId相关介绍参考开启透明数据加密
	KMSClusterId *string `json:"KMSClusterId,omitnil,omitempty" name:"KMSClusterId"`

	// 数据库引擎，支持：
	// <li>postgresql：云数据库PostgreSQL</li>
	// <li>mssql_compatible：MSSQL兼容-云数据库PostgreSQL</li>
	// 默认值：postgresql
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// 数据库引擎的配置信息，配置格式如下：
	// {"$key1":"$value1", "$key2":"$value2"}
	// 各引擎支持如下：
	// mssql_compatible引擎：
	// <li>migrationMode：数据库模式，可选参数，可取值：single-db（单数据库模式），multi-db（多数据库模式）。默认为single-db。</li>
	// <li>defaultLocale：排序区域规则，可选参数，在初始化后不可修改，默认为en_US，可选值如下：
	// "af_ZA", "sq_AL", "ar_DZ", "ar_BH", "ar_EG", "ar_IQ", "ar_JO", "ar_KW", "ar_LB", "ar_LY", "ar_MA", "ar_OM", "ar_QA", "ar_SA", "ar_SY", "ar_TN", "ar_AE", "ar_YE", "hy_AM", "az_Cyrl_AZ", "az_Latn_AZ", "eu_ES", "be_BY", "bg_BG", "ca_ES", "zh_HK", "zh_MO", "zh_CN", "zh_SG", "zh_TW", "hr_HR", "cs_CZ", "da_DK", "nl_BE", "nl_NL", "en_AU", "en_BZ", "en_CA", "en_IE", "en_JM", "en_NZ", "en_PH", "en_ZA", "en_TT", "en_GB", "en_US", "en_ZW", "et_EE", "fo_FO", "fa_IR", "fi_FI", "fr_BE", "fr_CA", "fr_FR", "fr_LU", "fr_MC", "fr_CH", "mk_MK", "ka_GE", "de_AT", "de_DE", "de_LI", "de_LU", "de_CH", "el_GR", "gu_IN", "he_IL", "hi_IN", "hu_HU", "is_IS", "id_ID", "it_IT", "it_CH", "ja_JP", "kn_IN", "kok_IN", "ko_KR", "ky_KG", "lv_LV", "lt_LT", "ms_BN", "ms_MY", "mr_IN", "mn_MN", "nb_NO", "nn_NO", "pl_PL", "pt_BR", "pt_PT", "pa_IN", "ro_RO", "ru_RU", "sa_IN", "sr_Cyrl_RS", "sr_Latn_RS", "sk_SK", "sl_SI", "es_AR", "es_BO", "es_CL", "es_CO", "es_CR", "es_DO", "es_EC", "es_SV", "es_GT", "es_HN", "es_MX", "es_NI", "es_PA", "es_PY","es_PE", "es_PR", "es_ES", "es_TRADITIONAL", "es_UY", "es_VE", "sw_KE", "sv_FI", "sv_SE", "tt_RU", "te_IN", "th_TH", "tr_TR", "uk_UA", "ur_IN", "ur_PK", "uz_Cyrl_UZ", "uz_Latn_UZ", "vi_VN"。</li>
	// <li>serverCollationName：排序规则名称，可选参数，在初始化后不可修改，默认为sql_latin1_general_cp1_ci_as，可选值如下："bbf_unicode_general_ci_as", "bbf_unicode_cp1_ci_as", "bbf_unicode_CP1250_ci_as", "bbf_unicode_CP1251_ci_as", "bbf_unicode_cp1253_ci_as", "bbf_unicode_cp1254_ci_as", "bbf_unicode_cp1255_ci_as", "bbf_unicode_cp1256_ci_as", "bbf_unicode_cp1257_ci_as", "bbf_unicode_cp1258_ci_as", "bbf_unicode_cp874_ci_as", "sql_latin1_general_cp1250_ci_as", "sql_latin1_general_cp1251_ci_as", "sql_latin1_general_cp1_ci_as", "sql_latin1_general_cp1253_ci_as", "sql_latin1_general_cp1254_ci_as", "sql_latin1_general_cp1255_ci_as","sql_latin1_general_cp1256_ci_as", "sql_latin1_general_cp1257_ci_as", "sql_latin1_general_cp1258_ci_as", "chinese_prc_ci_as", "cyrillic_general_ci_as", "finnish_swedish_ci_as", "french_ci_as", "japanese_ci_as", "korean_wansung_ci_as", "latin1_general_ci_as", "modern_spanish_ci_as", "polish_ci_as", "thai_ci_as", "traditional_spanish_ci_as", "turkish_ci_as", "ukrainian_ci_as", "vietnamese_ci_as"。</li>
	DBEngineConfig *string `json:"DBEngineConfig,omitnil,omitempty" name:"DBEngineConfig"`

	// 主从同步方式，支持： 
	// <li>Semi-sync：半同步</li>
	// <li>Async：异步</li>
	// 主实例默认值：Semi-sync
	// 只读实例默认值：Async
	SyncMode *string `json:"SyncMode,omitnil,omitempty" name:"SyncMode"`

	// 是否需要支持Ipv6：
	// <li>0：否</li>
	// <li>1：是</li>
	// 默认值：0
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitnil,omitempty" name:"NeedSupportIpv6"`

	// 实例是否开启删除保护: true-开启删除保护；false-关闭删除保护。
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

type CreateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例所属主可用区， 如：ap-guangzhou-3；若需要支持多可用区，在DBNodeSet.N字段中进行添加主可用区和备可用区信息；
	// 可用区信息可以通过调用 [DescribeZones](https://cloud.tencent.com/document/api/409/16769) 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 售卖规格码。该参数可以通过调用[DescribeClasses](https://cloud.tencent.com/document/api/409/89019)的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// 实例磁盘容量大小，单位：GB。该参数的设置步长为10。
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 购买实例数量，取值范围：[1-10]。一次性购买支持最大数量10个，若超过该数量，可进行多次调用进行购买。
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。
	// <li>预付费：支持1,2,3,4,5,6,7,8,9,10,11,12,24,36</li>
	// <li>后付费：只支持1</li>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 实例字符集，目前只支持：
	// <li> UTF8</li>
	// <li> LATIN1</li>
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// 实例根账号用户名，具体规范如下：
	// <li>用户名需要1-16个字符，只能由字母、数字或下划线组成</li>
	// <li>不能为postgres</li>
	// <li>不能由数字和pg_开头</li>
	// <li>所有规则均不区分大小写</li>
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// 实例根账号用户名对应的密码，长度8 ~ 32位，推荐使用12位以上的密码;不能以" / "开头;
	// 必须包含以下四项，字符种类:
	// <li>小写字母： [a ~ z]</li>
	// <li>大写字母：[A ～ Z]</li>
	// <li>数字：0 - 9</li>
	// <li>特殊字符：()`~!@#$%^&*-+=_|{}[]:;'<>,.?/</li>
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// PostgreSQL大版本号（该参数当前必传），版本信息可从[DescribeDBVersions](https://cloud.tencent.com/document/api/409/89018)获取。目前支持10，11，12，13，14，15这几个大版本，详情见[内核版本概述](https://cloud.tencent.com/document/product/409/67018)。
	// 输入该参数时，会基于此大版本号创建对应的最新小版本的最新内核版本号实例。
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// PostgreSQL社区大版本+小版本号。
	// 一般场景不推荐传入该参数。如需指定，只能传当前大版本号下最新小版本号。
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// PostgreSQL内核版本号。
	// 一般场景不推荐传入该参数。如需指定，只能传当前大版本号下最新内核版本号。
	DBKernelVersion *string `json:"DBKernelVersion,omitnil,omitempty" name:"DBKernelVersion"`

	// 实例计费类型，目前支持：
	// <li>PREPAID：预付费，即包年包月</li>
	// <li>POSTPAID_BY_HOUR：后付费，即按量计费</li>
	// 默认值：PREPAID
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 私有网络ID，形如vpc-xxxxxxxx（该参数当前必传）。有效的VpcId可通过登录控制台查询；也可以调用接口 [DescribeVpcEx](https://cloud.tencent.com/document/api/215/1372) ，从接口返回中的unVpcId字段获取。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络子网ID，形如subnet-xxxxxxxx（该参数当前必传）。有效的私有网络子网ID可通过登录控制台查询；也可以调用接口 [DescribeSubnets ](https://cloud.tencent.com/document/api/215/15784)，从接口返回中的unSubnetId字段获取。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例节点部署信息，支持多可用区部署时需要指定每个节点的部署可用区信息。
	// 可用区信息可以通过调用 [DescribeZones](https://cloud.tencent.com/document/api/409/16769) 接口的返回值中的Zone字段来获取。
	DBNodeSet []*DBNode `json:"DBNodeSet,omitnil,omitempty" name:"DBNodeSet"`

	// 续费标记：
	// <li>0：手动续费</li>
	// <li>1：自动续费</li>
	// 默认值：0
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 是否自动使用代金券：
	// <li>0：否</li>
	// <li>1：是</li>
	// 默认值：0
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 项目ID。默认取之为0，表示归属默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 活动ID。
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 实例名称，仅支持长度小于60的中文/英文/数字/"_"/"-"，不指定实例名称则默认显示"未命名"。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例需要绑定的Tag信息，默认为空；可以通过调用 [DescribeTags](https://cloud.tencent.com/document/api/651/35316) 返回值中的 Tags 字段来获取。
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 实例所属安全组，该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 是否需要支持数据透明加密：
	// <li>0：否</li>
	// <li>1：是</li>
	// 默认值：0
	// 参考[数据透明加密概述](https://cloud.tencent.com/document/product/409/71748)
	NeedSupportTDE *uint64 `json:"NeedSupportTDE,omitnil,omitempty" name:"NeedSupportTDE"`

	// 自定义密钥的KeyId，若选择自定义密匙加密，则需要传入自定义密匙的KeyId，KeyId是CMK的唯一标识。
	// KeyId创建获取相关参考[开启透明数据加密](https://cloud.tencent.com/document/product/409/71749)
	KMSKeyId *string `json:"KMSKeyId,omitnil,omitempty" name:"KMSKeyId"`

	// 使用KMS服务的地域，KMSRegion为空默认使用本地域的KMS，本地域不支持的情况下需自选其他KMS支持的地域。
	// KMSRegion相关介绍参考[开启透明数据加密](https://cloud.tencent.com/document/product/409/71749)
	KMSRegion *string `json:"KMSRegion,omitnil,omitempty" name:"KMSRegion"`

	// 指定KMS服务的集群，KMSClusterId为空使用默认集群的KMS，若选择指定KMS集群，则需要传入KMSClusterId。 KMSClusterId相关介绍参考开启透明数据加密
	KMSClusterId *string `json:"KMSClusterId,omitnil,omitempty" name:"KMSClusterId"`

	// 数据库引擎，支持：
	// <li>postgresql：云数据库PostgreSQL</li>
	// <li>mssql_compatible：MSSQL兼容-云数据库PostgreSQL</li>
	// 默认值：postgresql
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// 数据库引擎的配置信息，配置格式如下：
	// {"$key1":"$value1", "$key2":"$value2"}
	// 各引擎支持如下：
	// mssql_compatible引擎：
	// <li>migrationMode：数据库模式，可选参数，可取值：single-db（单数据库模式），multi-db（多数据库模式）。默认为single-db。</li>
	// <li>defaultLocale：排序区域规则，可选参数，在初始化后不可修改，默认为en_US，可选值如下：
	// "af_ZA", "sq_AL", "ar_DZ", "ar_BH", "ar_EG", "ar_IQ", "ar_JO", "ar_KW", "ar_LB", "ar_LY", "ar_MA", "ar_OM", "ar_QA", "ar_SA", "ar_SY", "ar_TN", "ar_AE", "ar_YE", "hy_AM", "az_Cyrl_AZ", "az_Latn_AZ", "eu_ES", "be_BY", "bg_BG", "ca_ES", "zh_HK", "zh_MO", "zh_CN", "zh_SG", "zh_TW", "hr_HR", "cs_CZ", "da_DK", "nl_BE", "nl_NL", "en_AU", "en_BZ", "en_CA", "en_IE", "en_JM", "en_NZ", "en_PH", "en_ZA", "en_TT", "en_GB", "en_US", "en_ZW", "et_EE", "fo_FO", "fa_IR", "fi_FI", "fr_BE", "fr_CA", "fr_FR", "fr_LU", "fr_MC", "fr_CH", "mk_MK", "ka_GE", "de_AT", "de_DE", "de_LI", "de_LU", "de_CH", "el_GR", "gu_IN", "he_IL", "hi_IN", "hu_HU", "is_IS", "id_ID", "it_IT", "it_CH", "ja_JP", "kn_IN", "kok_IN", "ko_KR", "ky_KG", "lv_LV", "lt_LT", "ms_BN", "ms_MY", "mr_IN", "mn_MN", "nb_NO", "nn_NO", "pl_PL", "pt_BR", "pt_PT", "pa_IN", "ro_RO", "ru_RU", "sa_IN", "sr_Cyrl_RS", "sr_Latn_RS", "sk_SK", "sl_SI", "es_AR", "es_BO", "es_CL", "es_CO", "es_CR", "es_DO", "es_EC", "es_SV", "es_GT", "es_HN", "es_MX", "es_NI", "es_PA", "es_PY","es_PE", "es_PR", "es_ES", "es_TRADITIONAL", "es_UY", "es_VE", "sw_KE", "sv_FI", "sv_SE", "tt_RU", "te_IN", "th_TH", "tr_TR", "uk_UA", "ur_IN", "ur_PK", "uz_Cyrl_UZ", "uz_Latn_UZ", "vi_VN"。</li>
	// <li>serverCollationName：排序规则名称，可选参数，在初始化后不可修改，默认为sql_latin1_general_cp1_ci_as，可选值如下："bbf_unicode_general_ci_as", "bbf_unicode_cp1_ci_as", "bbf_unicode_CP1250_ci_as", "bbf_unicode_CP1251_ci_as", "bbf_unicode_cp1253_ci_as", "bbf_unicode_cp1254_ci_as", "bbf_unicode_cp1255_ci_as", "bbf_unicode_cp1256_ci_as", "bbf_unicode_cp1257_ci_as", "bbf_unicode_cp1258_ci_as", "bbf_unicode_cp874_ci_as", "sql_latin1_general_cp1250_ci_as", "sql_latin1_general_cp1251_ci_as", "sql_latin1_general_cp1_ci_as", "sql_latin1_general_cp1253_ci_as", "sql_latin1_general_cp1254_ci_as", "sql_latin1_general_cp1255_ci_as","sql_latin1_general_cp1256_ci_as", "sql_latin1_general_cp1257_ci_as", "sql_latin1_general_cp1258_ci_as", "chinese_prc_ci_as", "cyrillic_general_ci_as", "finnish_swedish_ci_as", "french_ci_as", "japanese_ci_as", "korean_wansung_ci_as", "latin1_general_ci_as", "modern_spanish_ci_as", "polish_ci_as", "thai_ci_as", "traditional_spanish_ci_as", "turkish_ci_as", "ukrainian_ci_as", "vietnamese_ci_as"。</li>
	DBEngineConfig *string `json:"DBEngineConfig,omitnil,omitempty" name:"DBEngineConfig"`

	// 主从同步方式，支持： 
	// <li>Semi-sync：半同步</li>
	// <li>Async：异步</li>
	// 主实例默认值：Semi-sync
	// 只读实例默认值：Async
	SyncMode *string `json:"SyncMode,omitnil,omitempty" name:"SyncMode"`

	// 是否需要支持Ipv6：
	// <li>0：否</li>
	// <li>1：是</li>
	// 默认值：0
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitnil,omitempty" name:"NeedSupportIpv6"`

	// 实例是否开启删除保护: true-开启删除保护；false-关闭删除保护。
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

func (r *CreateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "SpecCode")
	delete(f, "Storage")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "Charset")
	delete(f, "AdminName")
	delete(f, "AdminPassword")
	delete(f, "DBMajorVersion")
	delete(f, "DBVersion")
	delete(f, "DBKernelVersion")
	delete(f, "InstanceChargeType")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "DBNodeSet")
	delete(f, "AutoRenewFlag")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "ProjectId")
	delete(f, "ActivityId")
	delete(f, "Name")
	delete(f, "TagList")
	delete(f, "SecurityGroupIds")
	delete(f, "NeedSupportTDE")
	delete(f, "KMSKeyId")
	delete(f, "KMSRegion")
	delete(f, "KMSClusterId")
	delete(f, "DBEngine")
	delete(f, "DBEngineConfig")
	delete(f, "SyncMode")
	delete(f, "NeedSupportIpv6")
	delete(f, "DeletionProtection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesResponseParams struct {
	// 订单号列表。每个实例对应一个订单号。
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 冻结流水号。
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// 创建成功的实例ID集合，只在后付费情景下有返回值。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstancesResponseParams `json:"Response"`
}

func (r *CreateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParameterTemplateRequestParams struct {
	// 模板名称，长度为1～60个字符，仅支持数字,英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 数据库大版本号，例如：11，12，13。可通过[DescribeDBVersions](https://cloud.tencent.com/document/api/409/89018)接口获取
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// 数据库引擎，例如：postgresql，mssql_compatible
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// 参数模板描述，长度为0～60个字符，仅支持数字,英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`
}

type CreateParameterTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称，长度为1～60个字符，仅支持数字,英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 数据库大版本号，例如：11，12，13。可通过[DescribeDBVersions](https://cloud.tencent.com/document/api/409/89018)接口获取
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// 数据库引擎，例如：postgresql，mssql_compatible
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// 参数模板描述，长度为0～60个字符，仅支持数字,英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`
}

func (r *CreateParameterTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateParameterTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "DBMajorVersion")
	delete(f, "DBEngine")
	delete(f, "TemplateDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateParameterTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParameterTemplateResponseParams struct {
	// 参数模板ID，用于唯一确认参数模板
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateParameterTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateParameterTemplateResponseParams `json:"Response"`
}

func (r *CreateParameterTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateParameterTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReadOnlyDBInstanceRequestParams struct {
	// 实例所属主可用区， 如：ap-guangzhou-3；
	// 可用区信息可以通过调用 [DescribeZones](https://cloud.tencent.com/document/api/409/16769) 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 只读实例的主实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitnil,omitempty" name:"MasterDBInstanceId"`

	// 售卖规格码。该参数可以通过调用[DescribeClasses](https://cloud.tencent.com/document/api/409/89019)的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// 实例硬盘容量大小，单位：GB。该参数的设置步长为10。
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 购买实例数量，取值范围：[1-6]。购买支持最大数量6个。
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。
	// <li>预付费：支持1,2,3,4,5,6,7,8,9,10,11,12,24,36</li>
	// <li>后付费：只支持1</li>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 私有网络ID，形如vpc-xxxxxxxx（该参数当前必传）。有效的VpcId可通过登录控制台查询；也可以调用接口 [DescribeVpcEx](https://cloud.tencent.com/document/api/215/1372) ，从接口返回中的unVpcId字段获取。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络子网ID，形如subnet-xxxxxxxx（该参数当前必传）。有效的私有网络子网ID可通过登录控制台查询；也可以调用接口 [DescribeSubnets ](https://cloud.tencent.com/document/api/215/15784)，从接口返回中的unSubnetId字段获取。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例计费类型，目前支持：
	// <li>PREPAID：预付费，即包年包月。</li>
	// <li>POSTPAID_BY_HOUR：后付费，即按量计费。</li>
	// 默认值：PREPAID。如果主实例为后付费，只读实例必须也为后付费。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 是否自动使用代金券：
	// <li>0：否</li>
	// <li>1：是</li>
	// 默认值：0
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 续费标记：
	// <li>0：手动续费</li>
	// <li>1：自动续费</li>
	// 默认值：0
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 项目ID。默认值为0，表示归属默认项目。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 优惠活动ID
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 只读组ID。
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// 实例需要绑定的Tag信息，默认为空；可以通过调用 [DescribeTags](https://cloud.tencent.com/document/api/651/35316) 返回值中的 Tags 字段来获取。
	TagList *Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 实例所属安全组，该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 是否需要支持Ipv6：
	// <li>0：否</li>
	// <li>1：是</li>
	// 默认值：0
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitnil,omitempty" name:"NeedSupportIpv6"`

	// 实例名。仅支持长度小于60的中文/英文/数字/"_"/"-"
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 不再需要指定，内核版本号与主实例保持一致
	//
	// Deprecated: DBVersion is deprecated.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// 专属集群ID
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 实例是否开启删除保护: true-开启删除保护；false-关闭删除保护。
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

type CreateReadOnlyDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例所属主可用区， 如：ap-guangzhou-3；
	// 可用区信息可以通过调用 [DescribeZones](https://cloud.tencent.com/document/api/409/16769) 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 只读实例的主实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitnil,omitempty" name:"MasterDBInstanceId"`

	// 售卖规格码。该参数可以通过调用[DescribeClasses](https://cloud.tencent.com/document/api/409/89019)的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// 实例硬盘容量大小，单位：GB。该参数的设置步长为10。
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 购买实例数量，取值范围：[1-6]。购买支持最大数量6个。
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。
	// <li>预付费：支持1,2,3,4,5,6,7,8,9,10,11,12,24,36</li>
	// <li>后付费：只支持1</li>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 私有网络ID，形如vpc-xxxxxxxx（该参数当前必传）。有效的VpcId可通过登录控制台查询；也可以调用接口 [DescribeVpcEx](https://cloud.tencent.com/document/api/215/1372) ，从接口返回中的unVpcId字段获取。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络子网ID，形如subnet-xxxxxxxx（该参数当前必传）。有效的私有网络子网ID可通过登录控制台查询；也可以调用接口 [DescribeSubnets ](https://cloud.tencent.com/document/api/215/15784)，从接口返回中的unSubnetId字段获取。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例计费类型，目前支持：
	// <li>PREPAID：预付费，即包年包月。</li>
	// <li>POSTPAID_BY_HOUR：后付费，即按量计费。</li>
	// 默认值：PREPAID。如果主实例为后付费，只读实例必须也为后付费。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 是否自动使用代金券：
	// <li>0：否</li>
	// <li>1：是</li>
	// 默认值：0
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 续费标记：
	// <li>0：手动续费</li>
	// <li>1：自动续费</li>
	// 默认值：0
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 项目ID。默认值为0，表示归属默认项目。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 优惠活动ID
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 只读组ID。
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// 实例需要绑定的Tag信息，默认为空；可以通过调用 [DescribeTags](https://cloud.tencent.com/document/api/651/35316) 返回值中的 Tags 字段来获取。
	TagList *Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 实例所属安全组，该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 是否需要支持Ipv6：
	// <li>0：否</li>
	// <li>1：是</li>
	// 默认值：0
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitnil,omitempty" name:"NeedSupportIpv6"`

	// 实例名。仅支持长度小于60的中文/英文/数字/"_"/"-"
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 不再需要指定，内核版本号与主实例保持一致
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// 专属集群ID
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 实例是否开启删除保护: true-开启删除保护；false-关闭删除保护。
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

func (r *CreateReadOnlyDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "MasterDBInstanceId")
	delete(f, "SpecCode")
	delete(f, "Storage")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "InstanceChargeType")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "AutoRenewFlag")
	delete(f, "ProjectId")
	delete(f, "ActivityId")
	delete(f, "ReadOnlyGroupId")
	delete(f, "TagList")
	delete(f, "SecurityGroupIds")
	delete(f, "NeedSupportIpv6")
	delete(f, "Name")
	delete(f, "DBVersion")
	delete(f, "DedicatedClusterId")
	delete(f, "DeletionProtection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReadOnlyDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReadOnlyDBInstanceResponseParams struct {
	// 订单号列表。每个实例对应一个订单号
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 冻结流水号
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// 创建成功的实例ID集合，只在后付费情景下有返回值
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// 入参有BillingParameters值时，出参才有值，值为商品下单的参数。
	BillingParameters *string `json:"BillingParameters,omitnil,omitempty" name:"BillingParameters"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReadOnlyDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateReadOnlyDBInstanceResponseParams `json:"Response"`
}

func (r *CreateReadOnlyDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReadOnlyGroupNetworkAccessRequestParams struct {
	// RO组ID，形如：pgrogrp-4t9c6g7k。可通过[DescribeReadOnlyGroups](https://cloud.tencent.com/document/api/409/52599)接口获取
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// 私有网络统一 ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 是否指定分配vip true-指定分配  false-自动分配。
	IsAssignVip *bool `json:"IsAssignVip,omitnil,omitempty" name:"IsAssignVip"`

	// 目标VIP地址。当不指定该参数，且IsAssignVip为true时，默认自动分配Vip。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

type CreateReadOnlyGroupNetworkAccessRequest struct {
	*tchttp.BaseRequest
	
	// RO组ID，形如：pgrogrp-4t9c6g7k。可通过[DescribeReadOnlyGroups](https://cloud.tencent.com/document/api/409/52599)接口获取
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// 私有网络统一 ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 是否指定分配vip true-指定分配  false-自动分配。
	IsAssignVip *bool `json:"IsAssignVip,omitnil,omitempty" name:"IsAssignVip"`

	// 目标VIP地址。当不指定该参数，且IsAssignVip为true时，默认自动分配Vip。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

func (r *CreateReadOnlyGroupNetworkAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyGroupNetworkAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReadOnlyGroupId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "IsAssignVip")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReadOnlyGroupNetworkAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReadOnlyGroupNetworkAccessResponseParams struct {
	// 流程ID，FlowId等同于TaskId
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReadOnlyGroupNetworkAccessResponse struct {
	*tchttp.BaseResponse
	Response *CreateReadOnlyGroupNetworkAccessResponseParams `json:"Response"`
}

func (r *CreateReadOnlyGroupNetworkAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyGroupNetworkAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReadOnlyGroupRequestParams struct {
	// 主实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitnil,omitempty" name:"MasterDBInstanceId"`

	// 只读组名称。仅支持长度小于60的中文/英文/数字/"_"/"-"。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 项目ID。默认值为0，表示归属于默认项目。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 私有网络ID。注：默认使用基础网络，当前不支持基础网络，故该参数必填。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID。注：默认使用基础网络，当前不支持基础网络，故该参数必填。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 延迟时间大小开关：0关、1开。该参数必填。
	ReplayLagEliminate *uint64 `json:"ReplayLagEliminate,omitnil,omitempty" name:"ReplayLagEliminate"`

	// 延迟空间大小开关： 0关、1开。该参数的填写需要与ReplayLagEliminate一致。
	ReplayLatencyEliminate *uint64 `json:"ReplayLatencyEliminate,omitnil,omitempty" name:"ReplayLatencyEliminate"`

	// 延迟时间大小阈值，取值为正整数，单位s。当ReplayLagEliminate为1时，该参数必填；当ReplayLagEliminate为0时，该参数需填0。
	MaxReplayLag *uint64 `json:"MaxReplayLag,omitnil,omitempty" name:"MaxReplayLag"`

	// 延迟空间大小阈值，取值为正整数，单位MB。当ReplayLatencyEliminate为1时，该参数必填；当ReplayLatencyEliminate为0时，该参数需填0。
	MaxReplayLatency *uint64 `json:"MaxReplayLatency,omitnil,omitempty" name:"MaxReplayLatency"`

	// 延迟剔除最小保留实例数。取值范围[0,100]。当ReplayLatencyEliminate为1时，该参数必填；当ReplayLagEliminate为0时，该参数无效。
	MinDelayEliminateReserve *uint64 `json:"MinDelayEliminateReserve,omitnil,omitempty" name:"MinDelayEliminateReserve"`

	// 安全组id
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type CreateReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 主实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitnil,omitempty" name:"MasterDBInstanceId"`

	// 只读组名称。仅支持长度小于60的中文/英文/数字/"_"/"-"。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 项目ID。默认值为0，表示归属于默认项目。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 私有网络ID。注：默认使用基础网络，当前不支持基础网络，故该参数必填。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID。注：默认使用基础网络，当前不支持基础网络，故该参数必填。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 延迟时间大小开关：0关、1开。该参数必填。
	ReplayLagEliminate *uint64 `json:"ReplayLagEliminate,omitnil,omitempty" name:"ReplayLagEliminate"`

	// 延迟空间大小开关： 0关、1开。该参数的填写需要与ReplayLagEliminate一致。
	ReplayLatencyEliminate *uint64 `json:"ReplayLatencyEliminate,omitnil,omitempty" name:"ReplayLatencyEliminate"`

	// 延迟时间大小阈值，取值为正整数，单位s。当ReplayLagEliminate为1时，该参数必填；当ReplayLagEliminate为0时，该参数需填0。
	MaxReplayLag *uint64 `json:"MaxReplayLag,omitnil,omitempty" name:"MaxReplayLag"`

	// 延迟空间大小阈值，取值为正整数，单位MB。当ReplayLatencyEliminate为1时，该参数必填；当ReplayLatencyEliminate为0时，该参数需填0。
	MaxReplayLatency *uint64 `json:"MaxReplayLatency,omitnil,omitempty" name:"MaxReplayLatency"`

	// 延迟剔除最小保留实例数。取值范围[0,100]。当ReplayLatencyEliminate为1时，该参数必填；当ReplayLagEliminate为0时，该参数无效。
	MinDelayEliminateReserve *uint64 `json:"MinDelayEliminateReserve,omitnil,omitempty" name:"MinDelayEliminateReserve"`

	// 安全组id
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *CreateReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MasterDBInstanceId")
	delete(f, "Name")
	delete(f, "ProjectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ReplayLagEliminate")
	delete(f, "ReplayLatencyEliminate")
	delete(f, "MaxReplayLag")
	delete(f, "MaxReplayLatency")
	delete(f, "MinDelayEliminateReserve")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReadOnlyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReadOnlyGroupResponseParams struct {
	// 只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// 流程ID，FlowId等同于TaskId
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateReadOnlyGroupResponseParams `json:"Response"`
}

func (r *CreateReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBBackup struct {
	// 备份文件唯一标识
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 文件生成的开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 文件生成的结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 文件大小(K)
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 策略（0-实例备份；1-多库备份）
	Strategy *int64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// 类型（0-定时）
	Way *int64 `json:"Way,omitnil,omitempty" name:"Way"`

	// 备份方式（1-完整）
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 状态（1-创建中；2-成功；3-失败）
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// DB列表
	DbList []*string `json:"DbList,omitnil,omitempty" name:"DbList"`

	// 内网下载地址
	InternalAddr *string `json:"InternalAddr,omitnil,omitempty" name:"InternalAddr"`

	// 外网下载地址
	ExternalAddr *string `json:"ExternalAddr,omitnil,omitempty" name:"ExternalAddr"`

	// 备份集ID
	SetId *string `json:"SetId,omitnil,omitempty" name:"SetId"`
}

type DBInstance struct {
	// 实例所属地域，如: ap-guangzhou，对应RegionSet的Region字段。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 实例所属可用区， 如：ap-guangzhou-3，对应ZoneSet的Zone字段。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 私有网络ID，形如vpc-e6w23k31。有效的VpcId可通过登录控制台查询；也可以调用接口 [DescribeVpcs](https://cloud.tencent.com/document/api/215/15778) ，从接口返回中的unVpcId字段获取。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络子网ID，形如subnet-51lcif9y。有效的私有网络子网ID可通过登录控制台查询；也可以调用接口 [DescribeSubnets ](https://cloud.tencent.com/document/api/215/15784)，从接口返回中的unSubnetId字段获取。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例ID。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 实例名称。
	DBInstanceName *string `json:"DBInstanceName,omitnil,omitempty" name:"DBInstanceName"`

	// 实例状态，分别为：applying（申请中）、init(待初始化)、initing(初始化中)、running(运行中)、limited run（受限运行）、isolating（隔离中）、isolated（已隔离）、disisolating（解隔离中）、recycling（回收中）、recycled（已回收）、job running（任务执行中）、offline（下线）、migrating（迁移中）、expanding（扩容中）、waitSwitch（等待切换）、switching（切换中）、readonly（只读）、restarting（重启中）、network changing（网络变更中）、upgrading（内核版本升级中）、audit-switching（审计状态变更中）、primary-switching（主备切换中）、offlining(下线中)、deployment changing（可用区变更中）、cloning（恢复数据中）、parameter modifying（参数修改中）、log-switching（日志状态变更中）、restoring（恢复中）、expanding（变配中）
	DBInstanceStatus *string `json:"DBInstanceStatus,omitnil,omitempty" name:"DBInstanceStatus"`

	// 实例分配的内存大小，单位：GB
	DBInstanceMemory *uint64 `json:"DBInstanceMemory,omitnil,omitempty" name:"DBInstanceMemory"`

	// 实例分配的存储空间大小，单位：GB
	DBInstanceStorage *uint64 `json:"DBInstanceStorage,omitnil,omitempty" name:"DBInstanceStorage"`

	// 实例分配的CPU数量，单位：个
	DBInstanceCpu *uint64 `json:"DBInstanceCpu,omitnil,omitempty" name:"DBInstanceCpu"`

	// 售卖规格ID
	DBInstanceClass *string `json:"DBInstanceClass,omitnil,omitempty" name:"DBInstanceClass"`

	// PostgreSQL大版本号，版本信息可从[DescribeDBVersions](https://cloud.tencent.com/document/api/409/89018)获取，目前支持10，11，12，13，14，15这几个大版本。
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// PostgreSQL社区大版本+小版本号，如12.4，版本信息可从[DescribeDBVersions](https://cloud.tencent.com/document/api/409/89018)获取。
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// PostgreSQL内核版本号，如v12.7_r1.8，版本信息可从[DescribeDBVersions](https://cloud.tencent.com/document/api/409/89018)获取。
	DBKernelVersion *string `json:"DBKernelVersion,omitnil,omitempty" name:"DBKernelVersion"`

	// 实例类型，类型有：
	// <li>primary：主实例</li>
	// <li>readonly：只读实例</li>
	// <li>guard：灾备实例</li>
	// <li>temp：临时实例</li>
	DBInstanceType *string `json:"DBInstanceType,omitnil,omitempty" name:"DBInstanceType"`

	// 实例版本，目前只支持standard（双机高可用版, 一主一从）。
	DBInstanceVersion *string `json:"DBInstanceVersion,omitnil,omitempty" name:"DBInstanceVersion"`

	// 实例字符集，目前只支持：
	// <li> UTF8</li>
	// <li> LATIN1</li>
	DBCharset *string `json:"DBCharset,omitnil,omitempty" name:"DBCharset"`

	// 实例创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例执行最后一次更新的时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 实例到期时间。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 实例隔离时间。
	IsolatedTime *string `json:"IsolatedTime,omitnil,omitempty" name:"IsolatedTime"`

	// 计费模式：
	// <li>prepaid：包年包月,预付费</li>
	// <li>postpaid：按量计费，后付费</li>
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// 是否自动续费：
	// <li>0：手动续费</li>
	// <li>1：自动续费</li>
	// 默认值：0
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 实例网络连接信息。
	DBInstanceNetInfo []*DBInstanceNetInfo `json:"DBInstanceNetInfo,omitnil,omitempty" name:"DBInstanceNetInfo"`

	// 机器类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 用户的AppId。
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 实例的Uid。
	Uid *uint64 `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 项目ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例绑定的标签信息。
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 主实例信息，仅在实例为只读实例时返回。
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitnil,omitempty" name:"MasterDBInstanceId"`

	// 只读实例数量。
	ReadOnlyInstanceNum *int64 `json:"ReadOnlyInstanceNum,omitnil,omitempty" name:"ReadOnlyInstanceNum"`

	// 只读实例在只读组中的状态。
	StatusInReadonlyGroup *string `json:"StatusInReadonlyGroup,omitnil,omitempty" name:"StatusInReadonlyGroup"`

	// 下线时间。
	OfflineTime *string `json:"OfflineTime,omitnil,omitempty" name:"OfflineTime"`

	// 实例的节点信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBNodeSet []*DBNode `json:"DBNodeSet,omitnil,omitempty" name:"DBNodeSet"`

	// 实例是否支持TDE数据加密：
	// <li>0：不支持</li>
	// <li>1：支持</li>
	// 默认值：0
	// TDE数据加密可参考[数据透明加密概述](https://cloud.tencent.com/document/product/409/71748)
	IsSupportTDE *int64 `json:"IsSupportTDE,omitnil,omitempty" name:"IsSupportTDE"`

	// 数据库引擎，支持：
	// <li>postgresql：云数据库PostgreSQL</li>
	// <li>mssql_compatible：MSSQL兼容-云数据库PostgreSQL</li>
	// 默认值：postgresql
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// 数据库引擎的配置信息，配置格式如下：
	// {"$key1":"$value1", "$key2":"$value2"}
	// 各引擎支持如下：
	// mssql_compatible引擎：
	// <li>migrationMode：数据库模式，可选参数，可取值：single-db（单数据库模式），multi-db（多数据库模式）。默认为single-db。</li>
	// <li>defaultLocale：排序区域规则，可选参数，在初始化后不可修改，默认为en_US，可选值如下：
	// "af_ZA", "sq_AL", "ar_DZ", "ar_BH", "ar_EG", "ar_IQ", "ar_JO", "ar_KW", "ar_LB", "ar_LY", "ar_MA", "ar_OM", "ar_QA", "ar_SA", "ar_SY", "ar_TN", "ar_AE", "ar_YE", "hy_AM", "az_Cyrl_AZ", "az_Latn_AZ", "eu_ES", "be_BY", "bg_BG", "ca_ES", "zh_HK", "zh_MO", "zh_CN", "zh_SG", "zh_TW", "hr_HR", "cs_CZ", "da_DK", "nl_BE", "nl_NL", "en_AU", "en_BZ", "en_CA", "en_IE", "en_JM", "en_NZ", "en_PH", "en_ZA", "en_TT", "en_GB", "en_US", "en_ZW", "et_EE", "fo_FO", "fa_IR", "fi_FI", "fr_BE", "fr_CA", "fr_FR", "fr_LU", "fr_MC", "fr_CH", "mk_MK", "ka_GE", "de_AT", "de_DE", "de_LI", "de_LU", "de_CH", "el_GR", "gu_IN", "he_IL", "hi_IN", "hu_HU", "is_IS", "id_ID", "it_IT", "it_CH", "ja_JP", "kn_IN", "kok_IN", "ko_KR", "ky_KG", "lv_LV", "lt_LT", "ms_BN", "ms_MY", "mr_IN", "mn_MN", "nb_NO", "nn_NO", "pl_PL", "pt_BR", "pt_PT", "pa_IN", "ro_RO", "ru_RU", "sa_IN", "sr_Cyrl_RS", "sr_Latn_RS", "sk_SK", "sl_SI", "es_AR", "es_BO", "es_CL", "es_CO", "es_CR", "es_DO", "es_EC", "es_SV", "es_GT", "es_HN", "es_MX", "es_NI", "es_PA", "es_PY","es_PE", "es_PR", "es_ES", "es_TRADITIONAL", "es_UY", "es_VE", "sw_KE", "sv_FI", "sv_SE", "tt_RU", "te_IN", "th_TH", "tr_TR", "uk_UA", "ur_IN", "ur_PK", "uz_Cyrl_UZ", "uz_Latn_UZ", "vi_VN"。</li>
	// <li>serverCollationName：排序规则名称，可选参数，在初始化后不可修改，默认为sql_latin1_general_cp1_ci_as，可选值如下："bbf_unicode_general_ci_as", "bbf_unicode_cp1_ci_as", "bbf_unicode_CP1250_ci_as", "bbf_unicode_CP1251_ci_as", "bbf_unicode_cp1253_ci_as", "bbf_unicode_cp1254_ci_as", "bbf_unicode_cp1255_ci_as", "bbf_unicode_cp1256_ci_as", "bbf_unicode_cp1257_ci_as", "bbf_unicode_cp1258_ci_as", "bbf_unicode_cp874_ci_as", "sql_latin1_general_cp1250_ci_as", "sql_latin1_general_cp1251_ci_as", "sql_latin1_general_cp1_ci_as", "sql_latin1_general_cp1253_ci_as", "sql_latin1_general_cp1254_ci_as", "sql_latin1_general_cp1255_ci_as","sql_latin1_general_cp1256_ci_as", "sql_latin1_general_cp1257_ci_as", "sql_latin1_general_cp1258_ci_as", "chinese_prc_ci_as", "cyrillic_general_ci_as", "finnish_swedish_ci_as", "french_ci_as", "japanese_ci_as", "korean_wansung_ci_as", "latin1_general_ci_as", "modern_spanish_ci_as", "polish_ci_as", "thai_ci_as", "traditional_spanish_ci_as", "turkish_ci_as", "ukrainian_ci_as", "vietnamese_ci_as"。</li>
	DBEngineConfig *string `json:"DBEngineConfig,omitnil,omitempty" name:"DBEngineConfig"`

	// 实例网络信息列表（此字段已废弃）
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkAccessList []*NetworkAccess `json:"NetworkAccessList,omitnil,omitempty" name:"NetworkAccessList"`

	// 实例是否支持Ipv6：
	// <li>0：否</li>
	// <li>1：是</li>
	// 默认值：0
	SupportIpv6 *uint64 `json:"SupportIpv6,omitnil,omitempty" name:"SupportIpv6"`

	// 实例已经弹性扩容的cpu核数
	ExpandedCpu *uint64 `json:"ExpandedCpu,omitnil,omitempty" name:"ExpandedCpu"`

	// 实例是否开启删除保护，取值如下：
	// - true：开启删除保护
	// - false：关闭删除保护
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

type DBInstanceNetInfo struct {
	// DNS域名
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// IP地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 连接Port地址
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 网络类型，1、inner（基础网络内网地址）；2、private（私有网络内网地址）；3、public（基础网络或私有网络的外网地址）；
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 网络连接状态，1、initing（未开通）；2、opened（已开通）；3、closed（已关闭）；4、opening（开通中）；5、closing（关闭中）；
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 连接数据库的协议类型，当前支持：postgresql、mssql（MSSQL兼容语法）
	ProtocolType *string `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`
}

type DBNode struct {
	// 节点类型，值可以为：
	// Primary，代表主节点；
	// Standby，代表备节点。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 节点所在可用区，例如 ap-guangzhou-1。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 专属集群ID
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

type Database struct {
	// 数据库名
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据库所有者
	DatabaseOwner *string `json:"DatabaseOwner,omitnil,omitempty" name:"DatabaseOwner"`

	// 数据库字符编码
	Encoding *string `json:"Encoding,omitnil,omitempty" name:"Encoding"`

	// 数据库排序规则
	Collate *string `json:"Collate,omitnil,omitempty" name:"Collate"`

	// 数据库字符分类
	Ctype *string `json:"Ctype,omitnil,omitempty" name:"Ctype"`

	// 数据库是否允许连接
	AllowConn *bool `json:"AllowConn,omitnil,omitempty" name:"AllowConn"`

	// 数据库最大连接数，-1表示无限制
	ConnLimit *int64 `json:"ConnLimit,omitnil,omitempty" name:"ConnLimit"`

	// 数据库权限列表
	Privileges *string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

type DatabaseObject struct {
	// 支持使用的数据库对象类型有：account,database,schema,sequence,procedure,type,function,table,view,matview,column。
	ObjectType *string `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`

	// 所描述的数据库对象名称
	ObjectName *string `json:"ObjectName,omitnil,omitempty" name:"ObjectName"`

	// 所要描述的数据库对象，所属的数据库名称。当描述对象类型不为database时，此参数必选。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 所要描述的数据库对象，所属的模式名称。当描述对象不为database、schema时，此参数必选。
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 所要描述的数据库对象，所属的表名称。当描述的对象类型为column时，此参数必填。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

type DatabasePrivilege struct {
	// 数据库对象，当ObjectType为database时，DatabaseName/SchemaName/TableName可为空；当ObjectType为schema时，SchemaName/TableName可为空；当ObjectType为column时，TableName不可为空，其余情况均可为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Object *DatabaseObject `json:"Object,omitnil,omitempty" name:"Object"`

	// 指定账号对数据库对象拥有的权限列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivilegeSet []*string `json:"PrivilegeSet,omitnil,omitempty" name:"PrivilegeSet"`
}

type DedicatedCluster struct {
	// 专属集群ID
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 专属集群名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 专属集群所在可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 灾备集群
	StandbyDedicatedClusterSet []*string `json:"StandbyDedicatedClusterSet,omitnil,omitempty" name:"StandbyDedicatedClusterSet"`

	// 实例数量
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// Cpu总数量
	CpuTotal *int64 `json:"CpuTotal,omitnil,omitempty" name:"CpuTotal"`

	// Cpu可用数量
	CpuAvailable *int64 `json:"CpuAvailable,omitnil,omitempty" name:"CpuAvailable"`

	// 内存总量，单位GB
	MemTotal *int64 `json:"MemTotal,omitnil,omitempty" name:"MemTotal"`

	// 内存可用量，单位GB
	MemAvailable *int64 `json:"MemAvailable,omitnil,omitempty" name:"MemAvailable"`

	// 磁盘总量，单位GB
	DiskTotal *int64 `json:"DiskTotal,omitnil,omitempty" name:"DiskTotal"`

	// 磁盘可用量，单位GB
	DiskAvailable *int64 `json:"DiskAvailable,omitnil,omitempty" name:"DiskAvailable"`
}

// Predefined struct for user
type DeleteAccountRequestParams struct {
	// 实例ID。	可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 删除的账号名称。	可通过[DescribeAccounts](https://cloud.tencent.com/document/api/409/18109)接口获取
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

type DeleteAccountRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。	可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 删除的账号名称。	可通过[DescribeAccounts](https://cloud.tencent.com/document/api/409/18109)接口获取
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

func (r *DeleteAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "UserName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccountResponseParams `json:"Response"`
}

func (r *DeleteAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupPlanRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 备份计划的ID。可通过[DescribeBackupPlans](https://cloud.tencent.com/document/api/409/68069)接口获取
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
}

type DeleteBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 备份计划的ID。可通过[DescribeBackupPlans](https://cloud.tencent.com/document/api/409/68069)接口获取
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
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
	delete(f, "DBInstanceId")
	delete(f, "PlanId")
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
type DeleteBaseBackupRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 数据备份ID。可通过[DescribeBaseBackups](https://cloud.tencent.com/document/api/409/89022)接口获取。7天内自动备份集不允许删除。
	BaseBackupId *string `json:"BaseBackupId,omitnil,omitempty" name:"BaseBackupId"`
}

type DeleteBaseBackupRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 数据备份ID。可通过[DescribeBaseBackups](https://cloud.tencent.com/document/api/409/89022)接口获取。7天内自动备份集不允许删除。
	BaseBackupId *string `json:"BaseBackupId,omitnil,omitempty" name:"BaseBackupId"`
}

func (r *DeleteBaseBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBaseBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "BaseBackupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBaseBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBaseBackupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBaseBackupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBaseBackupResponseParams `json:"Response"`
}

func (r *DeleteBaseBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBaseBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDBInstanceNetworkAccessRequestParams struct {
	// 实例ID，形如：postgres-6bwgamo3。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 私有网络统一 ID，若是基础网络则传"0"。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID，若是基础网络则传"0"。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 目标VIP地址。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

type DeleteDBInstanceNetworkAccessRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如：postgres-6bwgamo3。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 私有网络统一 ID，若是基础网络则传"0"。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID，若是基础网络则传"0"。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 目标VIP地址。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

func (r *DeleteDBInstanceNetworkAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDBInstanceNetworkAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDBInstanceNetworkAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDBInstanceNetworkAccessResponseParams struct {
	// 流程ID，FlowId等同于TaskId
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDBInstanceNetworkAccessResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDBInstanceNetworkAccessResponseParams `json:"Response"`
}

func (r *DeleteDBInstanceNetworkAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDBInstanceNetworkAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogBackupRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 日志备份ID。可通过[DescribeLogBackups](https://cloud.tencent.com/document/api/409/89021)接口获取。注：7天内自动备份集不允许删除。
	LogBackupId *string `json:"LogBackupId,omitnil,omitempty" name:"LogBackupId"`
}

type DeleteLogBackupRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 日志备份ID。可通过[DescribeLogBackups](https://cloud.tencent.com/document/api/409/89021)接口获取。注：7天内自动备份集不允许删除。
	LogBackupId *string `json:"LogBackupId,omitnil,omitempty" name:"LogBackupId"`
}

func (r *DeleteLogBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "LogBackupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLogBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogBackupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLogBackupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLogBackupResponseParams `json:"Response"`
}

func (r *DeleteLogBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteParameterTemplateRequestParams struct {
	// 参数模板ID，用于唯一确认待操作的参数模板。可通过[DescribeParameterTemplates](https://cloud.tencent.com/document/api/409/84067)接口获取
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteParameterTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板ID，用于唯一确认待操作的参数模板。可通过[DescribeParameterTemplates](https://cloud.tencent.com/document/api/409/84067)接口获取
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeleteParameterTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteParameterTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteParameterTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteParameterTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteParameterTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteParameterTemplateResponseParams `json:"Response"`
}

func (r *DeleteParameterTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteParameterTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReadOnlyGroupNetworkAccessRequestParams struct {
	// RO组ID，形如：pgrogrp-4t9c6g7k。可通过[DescribeReadOnlyGroups](https://cloud.tencent.com/document/api/409/52599)接口获取
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// 私有网络统一 ID，若是基础网络则传"0"。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID，若是基础网络则传"0"。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 目标VIP地址。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

type DeleteReadOnlyGroupNetworkAccessRequest struct {
	*tchttp.BaseRequest
	
	// RO组ID，形如：pgrogrp-4t9c6g7k。可通过[DescribeReadOnlyGroups](https://cloud.tencent.com/document/api/409/52599)接口获取
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// 私有网络统一 ID，若是基础网络则传"0"。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID，若是基础网络则传"0"。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 目标VIP地址。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

func (r *DeleteReadOnlyGroupNetworkAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReadOnlyGroupNetworkAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReadOnlyGroupId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReadOnlyGroupNetworkAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReadOnlyGroupNetworkAccessResponseParams struct {
	// 流程ID，FlowId等同于TaskId
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteReadOnlyGroupNetworkAccessResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReadOnlyGroupNetworkAccessResponseParams `json:"Response"`
}

func (r *DeleteReadOnlyGroupNetworkAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReadOnlyGroupNetworkAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReadOnlyGroupRequestParams struct {
	// 待删除只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

type DeleteReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 待删除只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

func (r *DeleteReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReadOnlyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReadOnlyGroupResponseParams struct {
	// 流程ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReadOnlyGroupResponseParams `json:"Response"`
}

func (r *DeleteReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountPrivilegesRequestParams struct {
	// 实例ID。	可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 查询此账号对某数据库对象所拥有的权限信息。账号名可通过[DescribeAccounts](https://cloud.tencent.com/document/api/409/18109)接口获取
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 要查询的数据库对象信息
	DatabaseObjectSet []*DatabaseObject `json:"DatabaseObjectSet,omitnil,omitempty" name:"DatabaseObjectSet"`
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。	可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 查询此账号对某数据库对象所拥有的权限信息。账号名可通过[DescribeAccounts](https://cloud.tencent.com/document/api/409/18109)接口获取
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 要查询的数据库对象信息
	DatabaseObjectSet []*DatabaseObject `json:"DatabaseObjectSet,omitnil,omitempty" name:"DatabaseObjectSet"`
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
	delete(f, "DBInstanceId")
	delete(f, "UserName")
	delete(f, "DatabaseObjectSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountPrivilegesResponseParams struct {
	// 用户拥有数据库user_database的CREATE、CONNECT、TEMPORARY权限
	PrivilegeSet []*DatabasePrivilege `json:"PrivilegeSet,omitnil,omitempty" name:"PrivilegeSet"`

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
	// 实例ID，形如postgres-6fego161。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 分页返回，每页最大返回数目，默认20，取值范围为1-100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数据按照创建时间或者用户名排序。取值支持createTime、name、updateTime。createTime-按照创建时间排序；name-按照用户名排序; updateTime-按照更新时间排序。
	// 默认值：createTime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回结果是升序还是降序。取值只能为desc或者asc。desc-降序；asc-升序
	// 默认值：desc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-6fego161。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 分页返回，每页最大返回数目，默认20，取值范围为1-100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数据按照创建时间或者用户名排序。取值支持createTime、name、updateTime。createTime-按照创建时间排序；name-按照用户名排序; updateTime-按照更新时间排序。
	// 默认值：createTime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回结果是升序还是降序。取值只能为desc或者asc。desc-降序；asc-升序
	// 默认值：desc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
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
	delete(f, "DBInstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsResponseParams struct {
	// 本次调用接口共返回了多少条数据。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 账号列表详细信息。当CreateTime项为0000-00-00 00:00:00时，意味着对应账号是直连数据库创建的，并非通过CreateAccount接口创建。
	Details []*AccountInfo `json:"Details,omitnil,omitempty" name:"Details"`

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
type DescribeAvailableRecoveryTimeRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type DescribeAvailableRecoveryTimeRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *DescribeAvailableRecoveryTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableRecoveryTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAvailableRecoveryTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableRecoveryTimeResponseParams struct {
	// 可恢复的最早时间，时区为东八区（UTC+8）。
	RecoveryBeginTime *string `json:"RecoveryBeginTime,omitnil,omitempty" name:"RecoveryBeginTime"`

	// 可恢复的最晚时间，时区为东八区（UTC+8）。
	RecoveryEndTime *string `json:"RecoveryEndTime,omitnil,omitempty" name:"RecoveryEndTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAvailableRecoveryTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAvailableRecoveryTimeResponseParams `json:"Response"`
}

func (r *DescribeAvailableRecoveryTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableRecoveryTimeResponse) FromJsonString(s string) error {
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
	// 备份文件下载限制类型，NONE 无限制，内外网都可以下载；INTRANET 只允许内网下载；CUSTOMIZE 自定义限制下载的vpc或ip。
	RestrictionType *string `json:"RestrictionType,omitnil,omitempty" name:"RestrictionType"`

	// vpc限制效力，ALLOW 允许；DENY 拒绝。
	VpcRestrictionEffect *string `json:"VpcRestrictionEffect,omitnil,omitempty" name:"VpcRestrictionEffect"`

	// 允许或拒绝下载备份文件的vpcId列表。
	VpcIdSet []*string `json:"VpcIdSet,omitnil,omitempty" name:"VpcIdSet"`

	// ip限制效力，ALLOW 允许；DENY 拒绝。
	IpRestrictionEffect *string `json:"IpRestrictionEffect,omitnil,omitempty" name:"IpRestrictionEffect"`

	// 允许或拒绝下载备份文件的ip列表。
	IpSet []*string `json:"IpSet,omitnil,omitempty" name:"IpSet"`

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
type DescribeBackupDownloadURLRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 备份类型，目前支持：LogBackup，BaseBackup。
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 备份的唯一ID。
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 链接的有效时间，取值为[0,36]，默认为12小时。
	URLExpireTime *uint64 `json:"URLExpireTime,omitnil,omitempty" name:"URLExpireTime"`

	// 备份下载限制
	BackupDownloadRestriction *BackupDownloadRestriction `json:"BackupDownloadRestriction,omitnil,omitempty" name:"BackupDownloadRestriction"`
}

type DescribeBackupDownloadURLRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 备份类型，目前支持：LogBackup，BaseBackup。
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 备份的唯一ID。
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 链接的有效时间，取值为[0,36]，默认为12小时。
	URLExpireTime *uint64 `json:"URLExpireTime,omitnil,omitempty" name:"URLExpireTime"`

	// 备份下载限制
	BackupDownloadRestriction *BackupDownloadRestriction `json:"BackupDownloadRestriction,omitnil,omitempty" name:"BackupDownloadRestriction"`
}

func (r *DescribeBackupDownloadURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "BackupType")
	delete(f, "BackupId")
	delete(f, "URLExpireTime")
	delete(f, "BackupDownloadRestriction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadURLResponseParams struct {
	// 备份的下载地址。
	BackupDownloadURL *string `json:"BackupDownloadURL,omitnil,omitempty" name:"BackupDownloadURL"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupDownloadURLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDownloadURLResponseParams `json:"Response"`
}

func (r *DescribeBackupDownloadURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupOverviewRequestParams struct {

}

type DescribeBackupOverviewRequest struct {
	*tchttp.BaseRequest
	
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
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupOverviewResponseParams struct {
	// 总免费空间大小，单位byte。
	TotalFreeSize *uint64 `json:"TotalFreeSize,omitnil,omitempty" name:"TotalFreeSize"`

	// 已使用免费空间大小，单位byte。
	UsedFreeSize *uint64 `json:"UsedFreeSize,omitnil,omitempty" name:"UsedFreeSize"`

	// 已使用收费空间大小，单位byte。
	UsedBillingSize *uint64 `json:"UsedBillingSize,omitnil,omitempty" name:"UsedBillingSize"`

	// 日志备份数量。
	LogBackupCount *uint64 `json:"LogBackupCount,omitnil,omitempty" name:"LogBackupCount"`

	// 日志备份大小，单位byte。
	LogBackupSize *uint64 `json:"LogBackupSize,omitnil,omitempty" name:"LogBackupSize"`

	// 手动创建的基础备份数量。
	ManualBaseBackupCount *uint64 `json:"ManualBaseBackupCount,omitnil,omitempty" name:"ManualBaseBackupCount"`

	// 手动创建的基础备份大小，单位byte。
	ManualBaseBackupSize *uint64 `json:"ManualBaseBackupSize,omitnil,omitempty" name:"ManualBaseBackupSize"`

	// 自动创建的基础备份数量。
	AutoBaseBackupCount *uint64 `json:"AutoBaseBackupCount,omitnil,omitempty" name:"AutoBaseBackupCount"`

	// 自动创建的基础备份大小，单位byte。
	AutoBaseBackupSize *uint64 `json:"AutoBaseBackupSize,omitnil,omitempty" name:"AutoBaseBackupSize"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeBackupPlansRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type DescribeBackupPlansRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
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
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupPlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupPlansResponseParams struct {
	// 实例的备份计划集
	Plans []*BackupPlan `json:"Plans,omitnil,omitempty" name:"Plans"`

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
type DescribeBackupSummariesRequestParams struct {
	// 每页显示数量，取值范围为1-100，默认为返回10条。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 按照一个或者多个过滤条件进行查询，目前支持的过滤条件有：
	// db-instance-id：按照实例ID过滤，类型为string。
	// db-instance-name：按照实例名过滤，支持模糊匹配，类型为string。
	// db-instance-ip：按照实例私有网络IP地址过滤，类型为string。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段，支持TotalBackupSize - 备份总大小、LogBackupSize - 备份日志的大小、ManualBaseBackupSize - 手动备份数据大小、AutoBaseBackupSize - 自动备份数据大小。当不传入该参数时，默认不进行排序。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式，包括升序：asc，降序：desc。默认值：asc。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeBackupSummariesRequest struct {
	*tchttp.BaseRequest
	
	// 每页显示数量，取值范围为1-100，默认为返回10条。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 按照一个或者多个过滤条件进行查询，目前支持的过滤条件有：
	// db-instance-id：按照实例ID过滤，类型为string。
	// db-instance-name：按照实例名过滤，支持模糊匹配，类型为string。
	// db-instance-ip：按照实例私有网络IP地址过滤，类型为string。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段，支持TotalBackupSize - 备份总大小、LogBackupSize - 备份日志的大小、ManualBaseBackupSize - 手动备份数据大小、AutoBaseBackupSize - 自动备份数据大小。当不传入该参数时，默认不进行排序。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式，包括升序：asc，降序：desc。默认值：asc。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
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
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupSummariesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupSummariesResponseParams struct {
	// 备份统计信息列表。
	BackupSummarySet []*BackupSummary `json:"BackupSummarySet,omitnil,omitempty" name:"BackupSummarySet"`

	// 查询到的所有备份信息数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeBaseBackupsRequestParams struct {
	// 备份的最小结束时间，形如2018-01-01 00:00:00。默认为7天前。
	MinFinishTime *string `json:"MinFinishTime,omitnil,omitempty" name:"MinFinishTime"`

	// 备份的最大结束时间，形如2018-01-01 00:00:00。默认为当前时间。
	MaxFinishTime *string `json:"MaxFinishTime,omitnil,omitempty" name:"MaxFinishTime"`

	// 按照一个或者多个过滤条件进行查询，目前支持的过滤条件有：
	// db-instance-id：按照实例ID过滤，类型为string。
	// db-instance-name：按照实例名过滤，支持模糊匹配，类型为string。
	// db-instance-ip：按照实例私有网络IP地址过滤，类型为string。
	// base-backup-id：按照备份集ID过滤，类型为string。
	// db-instance-status：按实例状态过滤，类型为string。取值参考[DBInstance](https://cloud.tencent.com/document/api/409/16778#DBInstance)结构的DBInstanceStatus字段。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 每页显示数量，取值范围为1-100，默认为返回10条。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，支持StartTime,FinishTime,Size。默认值：StartTime。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式，包括升序：asc，降序：desc。默认值：desc。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeBaseBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 备份的最小结束时间，形如2018-01-01 00:00:00。默认为7天前。
	MinFinishTime *string `json:"MinFinishTime,omitnil,omitempty" name:"MinFinishTime"`

	// 备份的最大结束时间，形如2018-01-01 00:00:00。默认为当前时间。
	MaxFinishTime *string `json:"MaxFinishTime,omitnil,omitempty" name:"MaxFinishTime"`

	// 按照一个或者多个过滤条件进行查询，目前支持的过滤条件有：
	// db-instance-id：按照实例ID过滤，类型为string。
	// db-instance-name：按照实例名过滤，支持模糊匹配，类型为string。
	// db-instance-ip：按照实例私有网络IP地址过滤，类型为string。
	// base-backup-id：按照备份集ID过滤，类型为string。
	// db-instance-status：按实例状态过滤，类型为string。取值参考[DBInstance](https://cloud.tencent.com/document/api/409/16778#DBInstance)结构的DBInstanceStatus字段。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 每页显示数量，取值范围为1-100，默认为返回10条。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，支持StartTime,FinishTime,Size。默认值：StartTime。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式，包括升序：asc，降序：desc。默认值：desc。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeBaseBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaseBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MinFinishTime")
	delete(f, "MaxFinishTime")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaseBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaseBackupsResponseParams struct {
	// 查询到的数据备份数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 数据备份详细信息列表。
	BaseBackupSet []*BaseBackup `json:"BaseBackupSet,omitnil,omitempty" name:"BaseBackupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBaseBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaseBackupsResponseParams `json:"Response"`
}

func (r *DescribeBaseBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaseBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassesRequestParams struct {
	// 可用区名称。可以通过接口[DescribeZones](https://cloud.tencent.com/document/product/409/16769)获取。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 数据库引擎，支持：
	// 1、postgresql（云数据库PostgreSQL）；
	// 2、mssql_compatible（MSSQL兼容-云数据库PostgreSQL）；
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// 数据库主版本号。例如12，13，可以通过接口[DescribeDBVersions](https://cloud.tencent.com/document/product/409/89018)获取。
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`
}

type DescribeClassesRequest struct {
	*tchttp.BaseRequest
	
	// 可用区名称。可以通过接口[DescribeZones](https://cloud.tencent.com/document/product/409/16769)获取。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 数据库引擎，支持：
	// 1、postgresql（云数据库PostgreSQL）；
	// 2、mssql_compatible（MSSQL兼容-云数据库PostgreSQL）；
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// 数据库主版本号。例如12，13，可以通过接口[DescribeDBVersions](https://cloud.tencent.com/document/product/409/89018)获取。
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`
}

func (r *DescribeClassesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "DBEngine")
	delete(f, "DBMajorVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClassesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassesResponseParams struct {
	// 数据库规格列表
	ClassInfoSet []*ClassInfo `json:"ClassInfoSet,omitnil,omitempty" name:"ClassInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClassesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClassesResponseParams `json:"Response"`
}

func (r *DescribeClassesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloneDBInstanceSpecRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 基础备份集ID，可通过[DescribeBaseBackups](https://cloud.tencent.com/document/api/409/89022)接口获取。此入参和RecoveryTargetTime必须选择一个传入。如与RecoveryTargetTime参数同时设置，则以此参数为准。
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// 恢复目标时间，此入参和BackupSetId必须选择一个传入。时区以东八区（UTC+8）为准。
	RecoveryTargetTime *string `json:"RecoveryTargetTime,omitnil,omitempty" name:"RecoveryTargetTime"`
}

type DescribeCloneDBInstanceSpecRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 基础备份集ID，可通过[DescribeBaseBackups](https://cloud.tencent.com/document/api/409/89022)接口获取。此入参和RecoveryTargetTime必须选择一个传入。如与RecoveryTargetTime参数同时设置，则以此参数为准。
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// 恢复目标时间，此入参和BackupSetId必须选择一个传入。时区以东八区（UTC+8）为准。
	RecoveryTargetTime *string `json:"RecoveryTargetTime,omitnil,omitempty" name:"RecoveryTargetTime"`
}

func (r *DescribeCloneDBInstanceSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloneDBInstanceSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "BackupSetId")
	delete(f, "RecoveryTargetTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloneDBInstanceSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloneDBInstanceSpecResponseParams struct {
	// 可购买的最小规格码。
	MinSpecCode *string `json:"MinSpecCode,omitnil,omitempty" name:"MinSpecCode"`

	// 可购买的最小磁盘容量，单位GB。
	MinStorage *int64 `json:"MinStorage,omitnil,omitempty" name:"MinStorage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloneDBInstanceSpecResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloneDBInstanceSpecResponseParams `json:"Response"`
}

func (r *DescribeCloneDBInstanceSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloneDBInstanceSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBBackupsRequestParams struct {
	// 实例ID，形如postgres-4wdeb0zv。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 备份方式（1-全量）。目前只支持全量，取值为1。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 查询开始时间，形如2018-06-10 17:06:38，起始时间不得小于7天以前
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 备份列表分页返回，每页返回数量，默认为 20，最小为1，最大值为 100。（当该参数不传或者传0时按默认值处理）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回结果中的第几页，从第0页开始。默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeDBBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-4wdeb0zv。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 备份方式（1-全量）。目前只支持全量，取值为1。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 查询开始时间，形如2018-06-10 17:06:38，起始时间不得小于7天以前
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 备份列表分页返回，每页返回数量，默认为 20，最小为1，最大值为 100。（当该参数不传或者传0时按默认值处理）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回结果中的第几页，从第0页开始。默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeDBBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "Type")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBBackupsResponseParams struct {
	// 返回备份列表中备份文件的个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 备份列表
	BackupList []*DBBackup `json:"BackupList,omitnil,omitempty" name:"BackupList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBBackupsResponseParams `json:"Response"`
}

func (r *DescribeDBBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBErrlogsRequestParams struct {
	// 实例ID。	可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 查询起始时间，形如2018-01-01 00:00:00。日志保留时间默认为7天，起始时间不能超出保留时间范围。	
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-01-01 00:00:00。	
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 数据库名字。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 搜索关键字。
	SearchKeys []*string `json:"SearchKeys,omitnil,omitempty" name:"SearchKeys"`

	// 每页显示数量，取值范围为1-100。默认值为50。	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。默认值为0。	
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeDBErrlogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。	可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 查询起始时间，形如2018-01-01 00:00:00。日志保留时间默认为7天，起始时间不能超出保留时间范围。	
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-01-01 00:00:00。	
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 数据库名字。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 搜索关键字。
	SearchKeys []*string `json:"SearchKeys,omitnil,omitempty" name:"SearchKeys"`

	// 每页显示数量，取值范围为1-100。默认值为50。	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。默认值为0。	
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeDBErrlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBErrlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DatabaseName")
	delete(f, "SearchKeys")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBErrlogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBErrlogsResponseParams struct {
	// 查询到的日志数量，最大值为10000条。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 错误日志详细信息集合。
	Details []*ErrLogDetail `json:"Details,omitnil,omitempty" name:"Details"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBErrlogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBErrlogsResponseParams `json:"Response"`
}

func (r *DescribeDBErrlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBErrlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceAttributeRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type DescribeDBInstanceAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *DescribeDBInstanceAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceAttributeResponseParams struct {
	// 实例详细信息。
	DBInstance *DBInstance `json:"DBInstance,omitnil,omitempty" name:"DBInstance"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceAttributeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceAttributeResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceHAConfigRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type DescribeDBInstanceHAConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *DescribeDBInstanceHAConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceHAConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceHAConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceHAConfigResponseParams struct {
	// 主从同步方式：
	// <li>Semi-sync：半同步</li>
	// <li>Async：异步</li>
	SyncMode *string `json:"SyncMode,omitnil,omitempty" name:"SyncMode"`

	// 高可用备机最大延迟数据量。备节点延迟数据量小于等于该值，且备节点延迟时间小于等于MaxStandbyLag时，可以切换为主节点。
	// <li>单位：byte</li>
	// <li>参数范围：[1073741824, 322122547200]</li>
	MaxStandbyLatency *uint64 `json:"MaxStandbyLatency,omitnil,omitempty" name:"MaxStandbyLatency"`

	// 高可用备机最大延迟时间。备节点延迟时间小于等于该值，且备节点延迟数据量小于等于MaxStandbyLatency时，可以切换为主节点。
	// <li>单位：s</li>
	// <li>参数范围：[5, 10]</li>
	MaxStandbyLag *uint64 `json:"MaxStandbyLag,omitnil,omitempty" name:"MaxStandbyLag"`

	// 同步备机最大延迟数据量。备机延迟数据量小于等于该值，且该备机延迟时间小于等于MaxSyncStandbyLag时，则该备机采用同步复制；否则，采用异步复制。
	// 该参数值针对SyncMode设置为Semi-sync的实例有效。
	// 异步实例该字段返回null。
	// 半同步实例禁止退化为异步复制时，该字段返回null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxSyncStandbyLatency *uint64 `json:"MaxSyncStandbyLatency,omitnil,omitempty" name:"MaxSyncStandbyLatency"`

	// 同步备机最大延迟时间。备机延迟时间小于等于该值，且该备机延迟数据量小于等于MaxSyncStandbyLatency时，则该备机采用同步复制；否则，采用异步复制。
	// 该参数值针对SyncMode设置为Semi-sync的实例有效。
	// 异步实例不返回该字段。
	// 半同步实例禁止退化为异步复制时，不返回该字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxSyncStandbyLag *uint64 `json:"MaxSyncStandbyLag,omitnil,omitempty" name:"MaxSyncStandbyLag"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceHAConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceHAConfigResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceHAConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceHAConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceParametersRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 查询指定参数详情。ParamName为空或不传，默认返回全部参数列表
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`
}

type DescribeDBInstanceParametersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 查询指定参数详情。ParamName为空或不传，默认返回全部参数列表
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`
}

func (r *DescribeDBInstanceParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ParamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceParametersResponseParams struct {
	// 参数列表总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 参数列表返回详情
	Detail []*ParamInfo `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceParametersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceParametersResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceSSLConfigRequestParams struct {
	// 实例ID，形如postgres-6bwgamo3。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type DescribeDBInstanceSSLConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-6bwgamo3。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *DescribeDBInstanceSSLConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceSSLConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceSSLConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceSSLConfigResponseParams struct {
	// true 代表开通 ，false 代表未开通
	SSLEnabled *bool `json:"SSLEnabled,omitnil,omitempty" name:"SSLEnabled"`

	// 云端根证书下载链接
	CAUrl *string `json:"CAUrl,omitnil,omitempty" name:"CAUrl"`

	// 服务器证书中配置的内网或外网连接地址
	ConnectAddress *string `json:"ConnectAddress,omitnil,omitempty" name:"ConnectAddress"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceSSLConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceSSLConfigResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceSSLConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceSSLConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceSecurityGroupsRequestParams struct {
	// 实例ID，可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取。DBInstanceId和ReadOnlyGroupId至少传一个；如果都传，忽略ReadOnlyGroupId。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 只读组ID，可通过[DescribeReadOnlyGroups](https://cloud.tencent.com/document/api/409/52599)接口获取。DBInstanceId和ReadOnlyGroupId至少传一个；如果要查询只读组关联的安全组，只传ReadOnlyGroupId。
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

type DescribeDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取。DBInstanceId和ReadOnlyGroupId至少传一个；如果都传，忽略ReadOnlyGroupId。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 只读组ID，可通过[DescribeReadOnlyGroups](https://cloud.tencent.com/document/api/409/52599)接口获取。DBInstanceId和ReadOnlyGroupId至少传一个；如果要查询只读组关联的安全组，只传ReadOnlyGroupId。
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

func (r *DescribeDBInstanceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceSecurityGroupsResponseParams struct {
	// 安全组信息数组
	SecurityGroupSet []*SecurityGroup `json:"SecurityGroupSet,omitnil,omitempty" name:"SecurityGroupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesRequestParams struct {
	// 按照一个或者多个过滤条件进行查询，目前支持的过滤条件有：
	// db-instance-id：按照实例ID过滤，类型为string
	// db-instance-name：按照实例名过滤，支持模糊匹配，类型为string
	// db-project-id：按照项目ID过滤，类型为integer
	// db-pay-mode：按照实例付费模式过滤，prepaid - 预付费；postpaid - 后付费。类型为string
	// db-tag-key：按照标签键过滤，类型为string
	// db-private-ip： 按照实例私有网络IP过滤，类型为string
	// db-public-address： 按照实例外网地址过滤，类型为string
	// db-dedicated-cluster-id: 按照私有集群Id过滤，类型为string
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 每页显示数量，取值范围为0-100，传入0时，取默认配置。默认为返回10条。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序指标，如实例名、创建时间等，支持DBInstanceId,CreateTime,Name,EndTime。默认值：CreateTime。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式，包括升序：asc、降序：desc。默认值：asc。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个过滤条件进行查询，目前支持的过滤条件有：
	// db-instance-id：按照实例ID过滤，类型为string
	// db-instance-name：按照实例名过滤，支持模糊匹配，类型为string
	// db-project-id：按照项目ID过滤，类型为integer
	// db-pay-mode：按照实例付费模式过滤，prepaid - 预付费；postpaid - 后付费。类型为string
	// db-tag-key：按照标签键过滤，类型为string
	// db-private-ip： 按照实例私有网络IP过滤，类型为string
	// db-public-address： 按照实例外网地址过滤，类型为string
	// db-dedicated-cluster-id: 按照私有集群Id过滤，类型为string
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 每页显示数量，取值范围为0-100，传入0时，取默认配置。默认为返回10条。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序指标，如实例名、创建时间等，支持DBInstanceId,CreateTime,Name,EndTime。默认值：CreateTime。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式，包括升序：asc、降序：desc。默认值：asc。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
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
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesResponseParams struct {
	// 查询到的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例详细信息集合。
	DBInstanceSet []*DBInstance `json:"DBInstanceSet,omitnil,omitempty" name:"DBInstanceSet"`

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
type DescribeDBVersionsRequestParams struct {

}

type DescribeDBVersionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDBVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBVersionsResponseParams struct {
	// 数据库版本号信息列表
	VersionSet []*Version `json:"VersionSet,omitnil,omitempty" name:"VersionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBVersionsResponseParams `json:"Response"`
}

func (r *DescribeDBVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBXlogsRequestParams struct {
	// 实例ID，形如postgres-4wdeb0zv。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 查询开始时间，形如2018-06-10 17:06:38，起始时间不得小于7天以前
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页返回，表示返回第几页的条目。从第0页开始计数。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页返回，表示每页有多少条目。取值为1-100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDBXlogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-4wdeb0zv。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 查询开始时间，形如2018-06-10 17:06:38，起始时间不得小于7天以前
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页返回，表示返回第几页的条目。从第0页开始计数。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页返回，表示每页有多少条目。取值为1-100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDBXlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBXlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBXlogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBXlogsResponseParams struct {
	// 表示此次返回结果有多少条数据。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Xlog列表
	XlogList []*Xlog `json:"XlogList,omitnil,omitempty" name:"XlogList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBXlogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBXlogsResponseParams `json:"Response"`
}

func (r *DescribeDBXlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBXlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseObjectsRequestParams struct {
	// 实例ID。	可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 查询的对象类型。支持查询的数据对象有：database,schema,sequence,procedure,type,function,table,view,matview,column。
	ObjectType *string `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`

	// 单次显示数量，默认20。可选范围为[0,100]。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。	
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询对象所属的数据库。当查询对象类型不为database时，此参数必填。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 查询对象所属的模式。当查询对象类型不为database、schema时，此参数必填。
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 查询对象所属的表。当查询对象类型为column时，此参数必填。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

type DescribeDatabaseObjectsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。	可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 查询的对象类型。支持查询的数据对象有：database,schema,sequence,procedure,type,function,table,view,matview,column。
	ObjectType *string `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`

	// 单次显示数量，默认20。可选范围为[0,100]。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。	
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询对象所属的数据库。当查询对象类型不为database时，此参数必填。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 查询对象所属的模式。当查询对象类型不为database、schema时，此参数必填。
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 查询对象所属的表。当查询对象类型为column时，此参数必填。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
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
	delete(f, "DBInstanceId")
	delete(f, "ObjectType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "DatabaseName")
	delete(f, "SchemaName")
	delete(f, "TableName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseObjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseObjectsResponseParams struct {
	// 查询对象列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectSet []*string `json:"ObjectSet,omitnil,omitempty" name:"ObjectSet"`

	// 查询对象总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

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
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/product/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 按照一个或者多个过滤条件进行查询，目前支持的过滤条件有：database-name：按照数据库名称过滤，类型为string。此处使用模糊匹配搜索符合条件的数据库。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 数据偏移量，从0开始。	
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 单次显示数量。建议最大取值100。
	// 默认值：20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/product/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 按照一个或者多个过滤条件进行查询，目前支持的过滤条件有：database-name：按照数据库名称过滤，类型为string。此处使用模糊匹配搜索符合条件的数据库。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 数据偏移量，从0开始。	
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 单次显示数量。建议最大取值100。
	// 默认值：20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	delete(f, "DBInstanceId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesResponseParams struct {
	// 数据库信息
	Items []*string `json:"Items,omitnil,omitempty" name:"Items"`

	// 数据库总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 数据库详情列表
	Databases []*Database `json:"Databases,omitnil,omitempty" name:"Databases"`

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
type DescribeDedicatedClustersRequestParams struct {
	// 按照一个或者多个过滤条件进行查询，目前支持的过滤条件有：
	// dedicated-cluster-id: 按照专属集群ID筛选，类型为string
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeDedicatedClustersRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个过滤条件进行查询，目前支持的过滤条件有：
	// dedicated-cluster-id: 按照专属集群ID筛选，类型为string
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeDedicatedClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDedicatedClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDedicatedClustersResponseParams struct {
	// 专属集群信息
	DedicatedClusterSet []*DedicatedCluster `json:"DedicatedClusterSet,omitnil,omitempty" name:"DedicatedClusterSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDedicatedClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDedicatedClustersResponseParams `json:"Response"`
}

func (r *DescribeDedicatedClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDedicatedClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultParametersRequestParams struct {
	// 数据库版本，大版本号，例如11，12，13。可从[DescribeDBVersions](https://cloud.tencent.com/document/api/409/89018)接口获取
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// 数据库引擎，例如：postgresql,mssql_compatible
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`
}

type DescribeDefaultParametersRequest struct {
	*tchttp.BaseRequest
	
	// 数据库版本，大版本号，例如11，12，13。可从[DescribeDBVersions](https://cloud.tencent.com/document/api/409/89018)接口获取
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// 数据库引擎，例如：postgresql,mssql_compatible
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`
}

func (r *DescribeDefaultParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBMajorVersion")
	delete(f, "DBEngine")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDefaultParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultParametersResponseParams struct {
	// 参数个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 参数信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamInfoSet []*ParamInfo `json:"ParamInfoSet,omitnil,omitempty" name:"ParamInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDefaultParametersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDefaultParametersResponseParams `json:"Response"`
}

func (r *DescribeDefaultParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEncryptionKeysRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type DescribeEncryptionKeysRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *DescribeEncryptionKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEncryptionKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEncryptionKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEncryptionKeysResponseParams struct {
	// 实例密钥信息列表。
	EncryptionKeys []*EncryptionKey `json:"EncryptionKeys,omitnil,omitempty" name:"EncryptionKeys"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEncryptionKeysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEncryptionKeysResponseParams `json:"Response"`
}

func (r *DescribeEncryptionKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEncryptionKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogBackupsRequestParams struct {
	// 备份的最小结束时间，形如2018-01-01 00:00:00。默认为7天前。
	MinFinishTime *string `json:"MinFinishTime,omitnil,omitempty" name:"MinFinishTime"`

	// 备份的最大结束时间，形如2018-01-01 00:00:00。默认为当前时间。
	MaxFinishTime *string `json:"MaxFinishTime,omitnil,omitempty" name:"MaxFinishTime"`

	// 按照一个或者多个过滤条件进行查询，目前支持的过滤条件有：
	// db-instance-id：按照实例ID过滤，类型为string。
	// db-instance-name：按照实例名过滤，支持模糊匹配，类型为string。
	// db-instance-ip：按照实例私有网络IP地址过滤，类型为string。
	// db-instance-status：按实例状态过滤，类型为string。取值参考[DBInstance](https://cloud.tencent.com/document/api/409/16778#DBInstance)结构的DBInstanceStatus字段。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 每页显示数量，取值范围为1-100，默认为返回10条。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，支持StartTime,FinishTime,Size。默认值：StartTime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式，包括升序：asc，降序：desc。默认值：desc。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeLogBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 备份的最小结束时间，形如2018-01-01 00:00:00。默认为7天前。
	MinFinishTime *string `json:"MinFinishTime,omitnil,omitempty" name:"MinFinishTime"`

	// 备份的最大结束时间，形如2018-01-01 00:00:00。默认为当前时间。
	MaxFinishTime *string `json:"MaxFinishTime,omitnil,omitempty" name:"MaxFinishTime"`

	// 按照一个或者多个过滤条件进行查询，目前支持的过滤条件有：
	// db-instance-id：按照实例ID过滤，类型为string。
	// db-instance-name：按照实例名过滤，支持模糊匹配，类型为string。
	// db-instance-ip：按照实例私有网络IP地址过滤，类型为string。
	// db-instance-status：按实例状态过滤，类型为string。取值参考[DBInstance](https://cloud.tencent.com/document/api/409/16778#DBInstance)结构的DBInstanceStatus字段。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 每页显示数量，取值范围为1-100，默认为返回10条。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，支持StartTime,FinishTime,Size。默认值：StartTime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式，包括升序：asc，降序：desc。默认值：desc。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeLogBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MinFinishTime")
	delete(f, "MaxFinishTime")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogBackupsResponseParams struct {
	// 查询到的日志备份数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 日志备份详细信息列表。
	LogBackupSet []*LogBackup `json:"LogBackupSet,omitnil,omitempty" name:"LogBackupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogBackupsResponseParams `json:"Response"`
}

func (r *DescribeLogBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaintainTimeWindowRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type DescribeMaintainTimeWindowRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *DescribeMaintainTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintainTimeWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMaintainTimeWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaintainTimeWindowResponseParams struct {
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 维护开始时间。时区为东八区（UTC+8）
	MaintainStartTime *string `json:"MaintainStartTime,omitnil,omitempty" name:"MaintainStartTime"`

	// 维护持续时间。单位：小时
	MaintainDuration *uint64 `json:"MaintainDuration,omitnil,omitempty" name:"MaintainDuration"`

	// 维护周期
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitnil,omitempty" name:"MaintainWeekDays"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMaintainTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMaintainTimeWindowResponseParams `json:"Response"`
}

func (r *DescribeMaintainTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintainTimeWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrdersRequestParams struct {
	// 订单名集合
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

type DescribeOrdersRequest struct {
	*tchttp.BaseRequest
	
	// 订单名集合
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

func (r *DescribeOrdersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrdersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrdersResponseParams struct {
	// 订单数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 订单数组
	Deals []*PgDeal `json:"Deals,omitnil,omitempty" name:"Deals"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrdersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrdersResponseParams `json:"Response"`
}

func (r *DescribeOrdersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParameterTemplateAttributesRequestParams struct {
	// 参数模板ID。可通过[DescribeParameterTemplates](https://cloud.tencent.com/document/api/409/84067)接口获取
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeParameterTemplateAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板ID。可通过[DescribeParameterTemplates](https://cloud.tencent.com/document/api/409/84067)接口获取
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeParameterTemplateAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParameterTemplateAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParameterTemplateAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParameterTemplateAttributesResponseParams struct {
	// 参数模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 参数模板包含的参数个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 参数模板包含的参数信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamInfoSet []*ParamInfo `json:"ParamInfoSet,omitnil,omitempty" name:"ParamInfoSet"`

	// 参数模板名称
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 参数模板适用的数据库版本
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// 参数模板适用的数据库引擎
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// 参数模板描述
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeParameterTemplateAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParameterTemplateAttributesResponseParams `json:"Response"`
}

func (r *DescribeParameterTemplateAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParameterTemplateAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParameterTemplatesRequestParams struct {
	// 过滤条件，目前支持的过滤条件有：TemplateName, TemplateId，DBMajorVersion，DBEngine。TemplateName不支持模糊匹配。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 每页显示数量，[0，100]，默认 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序指标，枚举值，支持：CreateTime，TemplateName，DBMajorVersion。如果不指定该参数，默认将按照参数模板的编号倒序排列，也就是说最新添加的参数模板会排在最前面。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式，枚举值，支持：asc（升序） ，desc（降序）。默认值为asc。当未指定OrderBy时，该参数失效，此时排序方式为OrderBy参数描述中给出的默认排序方式。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeParameterTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，目前支持的过滤条件有：TemplateName, TemplateId，DBMajorVersion，DBEngine。TemplateName不支持模糊匹配。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 每页显示数量，[0，100]，默认 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序指标，枚举值，支持：CreateTime，TemplateName，DBMajorVersion。如果不指定该参数，默认将按照参数模板的编号倒序排列，也就是说最新添加的参数模板会排在最前面。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式，枚举值，支持：asc（升序） ，desc（降序）。默认值为asc。当未指定OrderBy时，该参数失效，此时排序方式为OrderBy参数描述中给出的默认排序方式。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeParameterTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParameterTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParameterTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParameterTemplatesResponseParams struct {
	// 符合查询条件的参数模板总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 参数模板列表
	ParameterTemplateSet []*ParameterTemplate `json:"ParameterTemplateSet,omitnil,omitempty" name:"ParameterTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeParameterTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParameterTemplatesResponseParams `json:"Response"`
}

func (r *DescribeParameterTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParameterTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamsEventRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type DescribeParamsEventRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *DescribeParamsEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamsEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamsEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamsEventResponseParams struct {
	// 参数修改事件总数，以参数为统计粒度
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例参数修改事件详情
	EventItems []*EventItem `json:"EventItems,omitnil,omitempty" name:"EventItems"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeParamsEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParamsEventResponseParams `json:"Response"`
}

func (r *DescribeParamsEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamsEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductConfigRequestParams struct {
	// 可用区名称
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 数据库引擎，支持：
	// 1、postgresql（云数据库PostgreSQL）；
	// 2、mssql_compatible（MSSQL兼容-云数据库PostgreSQL）；
	// 如不指定默认使用postgresql。
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`
}

type DescribeProductConfigRequest struct {
	*tchttp.BaseRequest
	
	// 可用区名称
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 数据库引擎，支持：
	// 1、postgresql（云数据库PostgreSQL）；
	// 2、mssql_compatible（MSSQL兼容-云数据库PostgreSQL）；
	// 如不指定默认使用postgresql。
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`
}

func (r *DescribeProductConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "DBEngine")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductConfigResponseParams struct {
	// 售卖规格列表。
	SpecInfoList []*SpecInfo `json:"SpecInfoList,omitnil,omitempty" name:"SpecInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProductConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductConfigResponseParams `json:"Response"`
}

func (r *DescribeProductConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReadOnlyGroupsRequestParams struct {
	// 按照一个或者多个过滤条件进行查询，目前支持的过滤条件有：
	// db-master-instance-id：按照主实例过滤，类型为string。
	// read-only-group-id：按照只读组ID过滤，类型为string。
	// 注：该参数的过滤条件中，db-master-instance-id为必须指定项。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询每一页的条数，默认为10，最大值99。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询的页码，默认为1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 查询排序依据，目前支持:ROGroupId,CreateTime,Name。默认值CreateTime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 查询排序依据类型，目前支持:desc,asc。默认值asc。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeReadOnlyGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个过滤条件进行查询，目前支持的过滤条件有：
	// db-master-instance-id：按照主实例过滤，类型为string。
	// read-only-group-id：按照只读组ID过滤，类型为string。
	// 注：该参数的过滤条件中，db-master-instance-id为必须指定项。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询每一页的条数，默认为10，最大值99。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询的页码，默认为1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 查询排序依据，目前支持:ROGroupId,CreateTime,Name。默认值CreateTime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 查询排序依据类型，目前支持:desc,asc。默认值asc。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeReadOnlyGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReadOnlyGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReadOnlyGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReadOnlyGroupsResponseParams struct {
	// 只读组列表
	ReadOnlyGroupList []*ReadOnlyGroup `json:"ReadOnlyGroupList,omitnil,omitempty" name:"ReadOnlyGroupList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReadOnlyGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReadOnlyGroupsResponseParams `json:"Response"`
}

func (r *DescribeReadOnlyGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReadOnlyGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsRequestParams struct {

}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsResponseParams struct {
	// 返回的结果数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 地域信息集合。
	RegionSet []*RegionInfo `json:"RegionSet,omitnil,omitempty" name:"RegionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegionsResponseParams `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryAnalysisRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 查询起始时间，形如2018-01-01 00:00:00。日志保留时间默认为7天，起始时间不能超出保留时间范围。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-01-01 00:00:00。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 数据库名字。	
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 排序字段，取值范围[CallNum,CostTime,AvgCostTime]。默认值为CallNum。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式，包括升序：asc 降序：desc。默认值为desc。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 每页显示数量，取值范围为1-100。默认值为50。	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeSlowQueryAnalysisRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 查询起始时间，形如2018-01-01 00:00:00。日志保留时间默认为7天，起始时间不能超出保留时间范围。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-01-01 00:00:00。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 数据库名字。	
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 排序字段，取值范围[CallNum,CostTime,AvgCostTime]。默认值为CallNum。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式，包括升序：asc 降序：desc。默认值为desc。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 每页显示数量，取值范围为1-100。默认值为50。	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeSlowQueryAnalysisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryAnalysisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DatabaseName")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowQueryAnalysisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryAnalysisResponseParams struct {
	// 查询到的总条数，最大值为10000条。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 查询到的慢SQL统计分析详细信息集合。
	Detail *Detail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowQueryAnalysisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowQueryAnalysisResponseParams `json:"Response"`
}

func (r *DescribeSlowQueryAnalysisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryAnalysisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryListRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 查询起始时间，形如2018-01-01 00:00:00。日志保留时间默认为7天，起始时间不能超出保留时间范围。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-01-01 00:00:00。	
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 数据库名字。	
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 排序方式，包括升序：asc 降序：desc。默认值为desc。	
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 排序字段，取值范围[SessionStartTime,Duration]。默认值为SessionStartTime。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 每页显示数量，取值范围为1-100。默认值为50。	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。默认值为0。	
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeSlowQueryListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 查询起始时间，形如2018-01-01 00:00:00。日志保留时间默认为7天，起始时间不能超出保留时间范围。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-01-01 00:00:00。	
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 数据库名字。	
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 排序方式，包括升序：asc 降序：desc。默认值为desc。	
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 排序字段，取值范围[SessionStartTime,Duration]。默认值为SessionStartTime。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 每页显示数量，取值范围为1-100。默认值为50。	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。默认值为0。	
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeSlowQueryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DatabaseName")
	delete(f, "OrderByType")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowQueryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryListResponseParams struct {
	// 查询到的慢日志数量，最大值为10000条。	
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 查询到的慢日志耗时分段分析结果。
	DurationAnalysis []*DurationAnalysis `json:"DurationAnalysis,omitnil,omitempty" name:"DurationAnalysis"`

	// 查询到的慢日志详细信息集合。
	RawSlowQueryList []*RawSlowQuery `json:"RawSlowQueryList,omitnil,omitempty" name:"RawSlowQueryList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowQueryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowQueryListResponseParams `json:"Response"`
}

func (r *DescribeSlowQueryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksRequestParams struct {
	// 按照任务ID进行查询。其余云API中返回的FlowId和TaskId等价。
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 按照数据库实例ID进行查询。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 任务的最早开始时间，形如2024-08-23 00:00:00,默认只展示180天内的数据。
	MinStartTime *string `json:"MinStartTime,omitnil,omitempty" name:"MinStartTime"`

	// 任务的最晚开始时间，形如2024-08-23 00:00:00，默认为当前时间。
	MaxStartTime *string `json:"MaxStartTime,omitnil,omitempty" name:"MaxStartTime"`

	// 每页显示数量，取值范围为1-100，默认为返回20条。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，支持StartTime,EndTime，默认为StartTime。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式，包括升序：asc，降序：desc，默认为desc。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 按照任务ID进行查询。其余云API中返回的FlowId和TaskId等价。
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 按照数据库实例ID进行查询。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 任务的最早开始时间，形如2024-08-23 00:00:00,默认只展示180天内的数据。
	MinStartTime *string `json:"MinStartTime,omitnil,omitempty" name:"MinStartTime"`

	// 任务的最晚开始时间，形如2024-08-23 00:00:00，默认为当前时间。
	MaxStartTime *string `json:"MaxStartTime,omitnil,omitempty" name:"MaxStartTime"`

	// 每页显示数量，取值范围为1-100，默认为返回20条。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，支持StartTime,EndTime，默认为StartTime。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式，包括升序：asc，降序：desc，默认为desc。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
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
	delete(f, "TaskId")
	delete(f, "DBInstanceId")
	delete(f, "MinStartTime")
	delete(f, "MaxStartTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// 查询到的任务数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务信息列表
	TaskSet []*TaskSet `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

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

}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
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
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// 返回的结果数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 可用区信息集合。
	ZoneSet []*ZoneInfo `json:"ZoneSet,omitnil,omitempty" name:"ZoneSet"`

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
type DestroyDBInstanceRequestParams struct {
	// 待下线实例ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type DestroyDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 待下线实例ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *DestroyDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyDBInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyDBInstanceResponseParams `json:"Response"`
}

func (r *DestroyDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Detail struct {
	// 输入时间范围内所有慢sql执行的总时间，单位毫秒（ms）
	TotalTime *float64 `json:"TotalTime,omitnil,omitempty" name:"TotalTime"`

	// 输入时间范围内所有慢sql总条数
	TotalCallNum *uint64 `json:"TotalCallNum,omitnil,omitempty" name:"TotalCallNum"`

	// 慢SQL统计分析列表
	AnalysisItems []*AnalysisItems `json:"AnalysisItems,omitnil,omitempty" name:"AnalysisItems"`
}

// Predefined struct for user
type DisIsolateDBInstancesRequestParams struct {
	// 实例ID列表。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取。支持同时解隔离多个实例。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// 购买时长，单位：月。
	// <li>预付费：支持1,2,3,4,5,6,7,8,9,10,11,12,24,36</li>
	// <li>后付费：该参数不生效</li>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 是否使用代金券：
	// <li>true：使用</li>
	// <li>false：不使用</li>
	// 默认值：false
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券id列表。
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`
}

type DisIsolateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取。支持同时解隔离多个实例。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// 购买时长，单位：月。
	// <li>预付费：支持1,2,3,4,5,6,7,8,9,10,11,12,24,36</li>
	// <li>后付费：该参数不生效</li>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 是否使用代金券：
	// <li>true：使用</li>
	// <li>false：不使用</li>
	// 默认值：false
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券id列表。
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`
}

func (r *DisIsolateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisIsolateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceIdSet")
	delete(f, "Period")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisIsolateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisIsolateDBInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisIsolateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DisIsolateDBInstancesResponseParams `json:"Response"`
}

func (r *DisIsolateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisIsolateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DurationAnalysis struct {
	// 慢SQL耗时，时段
	TimeSegment *string `json:"TimeSegment,omitnil,omitempty" name:"TimeSegment"`

	// 对应时段区间慢SQL 条数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type EncryptionKey struct {
	// KMS实例加密的KeyId。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// KMS实例加密Key的别名。
	KeyAlias *string `json:"KeyAlias,omitnil,omitempty" name:"KeyAlias"`

	// 实例加密密钥DEK的密文。
	DEKCipherTextBlob *string `json:"DEKCipherTextBlob,omitnil,omitempty" name:"DEKCipherTextBlob"`

	// 密钥是否启用，1-启用， 0-未启用。
	IsEnabled *int64 `json:"IsEnabled,omitnil,omitempty" name:"IsEnabled"`

	// KMS密钥所在地域。
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`

	// DEK密钥创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 密钥所在的KMS服务集群Id，为空表示密钥在默认的KMS集群中，不为空表示在指定的KMS服务集群中
	KMSClusterId *string `json:"KMSClusterId,omitnil,omitempty" name:"KMSClusterId"`
}

type ErrLogDetail struct {
	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 数据库名字
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 错误发生时间
	ErrTime *string `json:"ErrTime,omitnil,omitempty" name:"ErrTime"`

	// 错误消息
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`
}

type EventInfo struct {
	// 参数名
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 原参数值
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`

	// 本次修改期望参数值
	NewValue *string `json:"NewValue,omitnil,omitempty" name:"NewValue"`

	// 后台参数修改开始时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 后台参数生效开始时间
	EffectiveTime *string `json:"EffectiveTime,omitnil,omitempty" name:"EffectiveTime"`

	// 修改状态。枚举值：in progress、success、paused
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 操作者（一般为用户sub UIN）
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 时间日志。
	EventLog *string `json:"EventLog,omitnil,omitempty" name:"EventLog"`
}

type EventItem struct {
	// 参数名
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 修改事件数
	EventCount *int64 `json:"EventCount,omitnil,omitempty" name:"EventCount"`

	// 修改时间详情
	EventDetail []*EventInfo `json:"EventDetail,omitnil,omitempty" name:"EventDetail"`
}

type Filter struct {
	// 过滤键的名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type InquiryPriceCreateDBInstancesRequestParams struct {
	// 可用区名称。该参数可以通过调用[ DescribeZones](https://cloud.tencent.com/document/product/409/16769) 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 规格ID。该参数可以通过调用[DescribeClasses](https://cloud.tencent.com/document/product/409/89019)接口的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// 存储容量大小，单位：GB。该参数的设置步长为10。
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 实例数量。目前最大数量不超过100，如需一次性创建更多实例，请联系客服支持。
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 【弃字段，不再生效】，计费ID。该参数可以通过调用DescribeProductConfig接口的返回值中的Pid字段来获取。
	Pid *uint64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// 实例计费类型。目前支持：PREPAID（预付费，即包年包月）和 POSTPAID（按量计费）。
	// 默认值：PREPAID
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例类型，默认primary，支持如下：
	// primary（双机高可用（一主一从））
	// readonly（只读实例）
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// DB引擎，默认postgresql，支持如下：
	// postgresql（云数据库PostgreSQL）
	// mssql_compatible（MSSQL兼容-云数据库PostgreSQL）
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`
}

type InquiryPriceCreateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 可用区名称。该参数可以通过调用[ DescribeZones](https://cloud.tencent.com/document/product/409/16769) 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 规格ID。该参数可以通过调用[DescribeClasses](https://cloud.tencent.com/document/product/409/89019)接口的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// 存储容量大小，单位：GB。该参数的设置步长为10。
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 实例数量。目前最大数量不超过100，如需一次性创建更多实例，请联系客服支持。
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 【弃字段，不再生效】，计费ID。该参数可以通过调用DescribeProductConfig接口的返回值中的Pid字段来获取。
	Pid *uint64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// 实例计费类型。目前支持：PREPAID（预付费，即包年包月）和 POSTPAID（按量计费）。
	// 默认值：PREPAID
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例类型，默认primary，支持如下：
	// primary（双机高可用（一主一从））
	// readonly（只读实例）
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// DB引擎，默认postgresql，支持如下：
	// postgresql（云数据库PostgreSQL）
	// mssql_compatible（MSSQL兼容-云数据库PostgreSQL）
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`
}

func (r *InquiryPriceCreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "SpecCode")
	delete(f, "Storage")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "Pid")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceType")
	delete(f, "DBEngine")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateDBInstancesResponseParams struct {
	// 刊例价，单位：分
	OriginalPrice *uint64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 折后实际付款金额，单位：分
	Price *uint64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 币种。例如，CNY：人民币。
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceCreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceCreateDBInstancesResponseParams `json:"Response"`
}

func (r *InquiryPriceCreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewDBInstanceRequestParams struct {
	// 实例ID，可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)获取。
	// （此接口仅支持预付费实例的查询）
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 续费周期，按月计算
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`
}

type InquiryPriceRenewDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)获取。
	// （此接口仅支持预付费实例的查询）
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 续费周期，按月计算
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *InquiryPriceRenewDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceRenewDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewDBInstanceResponseParams struct {
	// 刊例价，单位为分。如24650表示246.5元
	OriginalPrice *int64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 折后实际付款金额，单位为分。如24650表示246.5元
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 币种。例如，CNY：人民币。
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceRenewDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceRenewDBInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceRenewDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpgradeDBInstanceRequestParams struct {
	// 实例的磁盘大小，单位GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 实例的内存大小，单位GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例ID，形如postgres-hez4fh0v。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 实例计费类型。
	//
	// Deprecated: InstanceChargeType is deprecated.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例的Cpu大小，单位Core。
	// 不传入此参数时，默认根据Memory确定的售卖规格所对应的Cpu进行设置。如Memory为2，支持的售卖规格有1核2GiB，则不传入Cpu时，Cpu默认为1。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`
}

type InquiryPriceUpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例的磁盘大小，单位GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 实例的内存大小，单位GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例ID，形如postgres-hez4fh0v。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 实例计费类型。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例的Cpu大小，单位Core。
	// 不传入此参数时，默认根据Memory确定的售卖规格所对应的Cpu进行设置。如Memory为2，支持的售卖规格有1核2GiB，则不传入Cpu时，Cpu默认为1。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`
}

func (r *InquiryPriceUpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Storage")
	delete(f, "Memory")
	delete(f, "DBInstanceId")
	delete(f, "InstanceChargeType")
	delete(f, "Cpu")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceUpgradeDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpgradeDBInstanceResponseParams struct {
	// 刊例价费用
	OriginalPrice *int64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 折后实际付款金额
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 币种。例如，CNY：人民币。
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceUpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceUpgradeDBInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceUpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDBInstancesRequestParams struct {
	// 实例ID集合。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取。注意：不推荐同时隔离多个实例。建议每次操作仅传入单个实例ID。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`
}

type IsolateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID集合。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取。注意：不推荐同时隔离多个实例。建议每次操作仅传入单个实例ID。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`
}

func (r *IsolateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDBInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *IsolateDBInstancesResponseParams `json:"Response"`
}

func (r *IsolateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LockAccountRequestParams struct {
	// 实例ID。	
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 账号名称。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

type LockAccountRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。	
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 账号名称。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

func (r *LockAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LockAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "UserName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LockAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LockAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type LockAccountResponse struct {
	*tchttp.BaseResponse
	Response *LockAccountResponseParams `json:"Response"`
}

func (r *LockAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LockAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogBackup struct {
	// 实例ID。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 备份文件唯一标识。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 备份文件名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备份方式。枚举值，physical - 物理备份；logical - 逻辑备份。
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 备份模式。枚举值，manual - 手动备份；automatic - 自动备份 。
	BackupMode *string `json:"BackupMode,omitnil,omitempty" name:"BackupMode"`

	// 备份任务状态。枚举值：init、running、finished、failed、canceled
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 备份集大小，单位bytes。
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 备份的开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 备份的结束时间。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 备份的过期时间。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

// Predefined struct for user
type ModifyAccountPrivilegesRequestParams struct {
	// 实例ID。	可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 修改此账号对某数据库对象的权限。可通过[DescribeAccounts](https://cloud.tencent.com/document/api/409/18109)接口获取
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 修改的权限信息，支持批量修改，一次最高修改50条。
	ModifyPrivilegeSet []*ModifyPrivilege `json:"ModifyPrivilegeSet,omitnil,omitempty" name:"ModifyPrivilegeSet"`
}

type ModifyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。	可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 修改此账号对某数据库对象的权限。可通过[DescribeAccounts](https://cloud.tencent.com/document/api/409/18109)接口获取
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 修改的权限信息，支持批量修改，一次最高修改50条。
	ModifyPrivilegeSet []*ModifyPrivilege `json:"ModifyPrivilegeSet,omitnil,omitempty" name:"ModifyPrivilegeSet"`
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
	delete(f, "DBInstanceId")
	delete(f, "UserName")
	delete(f, "ModifyPrivilegeSet")
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
type ModifyAccountRemarkRequestParams struct {
	// 实例ID，形如postgres-4wdeb0zv
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 实例用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户UserName对应的新备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyAccountRemarkRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-4wdeb0zv
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 实例用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户UserName对应的新备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyAccountRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountRemarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "UserName")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountRemarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountRemarkResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccountRemarkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountRemarkResponseParams `json:"Response"`
}

func (r *ModifyAccountRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadRestrictionRequestParams struct {
	// 备份文件下载限制类型，NONE 无限制，内外网都可以下载；INTRANET 只允许内网下载；CUSTOMIZE 自定义限制下载的vpc或ip。当该参数取值为CUSTOMIZE时，Vpc限制和Ip限制需要至少填写一项。
	RestrictionType *string `json:"RestrictionType,omitnil,omitempty" name:"RestrictionType"`

	// vpc限制效力，ALLOW 允许；DENY 拒绝。
	VpcRestrictionEffect *string `json:"VpcRestrictionEffect,omitnil,omitempty" name:"VpcRestrictionEffect"`

	// 允许或拒绝下载备份文件的vpcId列表。
	// **注意：**该入参会全量替换存量已有集合，非增量更新。修改需传入预期的全量集合。
	VpcIdSet []*string `json:"VpcIdSet,omitnil,omitempty" name:"VpcIdSet"`

	// ip限制效力，ALLOW 允许；DENY 拒绝。
	IpRestrictionEffect *string `json:"IpRestrictionEffect,omitnil,omitempty" name:"IpRestrictionEffect"`

	// 允许或拒绝下载备份文件的ip列表。
	// **注意：**该入参会全量替换存量已有集合，非增量更新。修改需传入预期的全量集合。
	IpSet []*string `json:"IpSet,omitnil,omitempty" name:"IpSet"`
}

type ModifyBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
	
	// 备份文件下载限制类型，NONE 无限制，内外网都可以下载；INTRANET 只允许内网下载；CUSTOMIZE 自定义限制下载的vpc或ip。当该参数取值为CUSTOMIZE时，Vpc限制和Ip限制需要至少填写一项。
	RestrictionType *string `json:"RestrictionType,omitnil,omitempty" name:"RestrictionType"`

	// vpc限制效力，ALLOW 允许；DENY 拒绝。
	VpcRestrictionEffect *string `json:"VpcRestrictionEffect,omitnil,omitempty" name:"VpcRestrictionEffect"`

	// 允许或拒绝下载备份文件的vpcId列表。
	// **注意：**该入参会全量替换存量已有集合，非增量更新。修改需传入预期的全量集合。
	VpcIdSet []*string `json:"VpcIdSet,omitnil,omitempty" name:"VpcIdSet"`

	// ip限制效力，ALLOW 允许；DENY 拒绝。
	IpRestrictionEffect *string `json:"IpRestrictionEffect,omitnil,omitempty" name:"IpRestrictionEffect"`

	// 允许或拒绝下载备份文件的ip列表。
	// **注意：**该入参会全量替换存量已有集合，非增量更新。修改需传入预期的全量集合。
	IpSet []*string `json:"IpSet,omitnil,omitempty" name:"IpSet"`
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
	delete(f, "RestrictionType")
	delete(f, "VpcRestrictionEffect")
	delete(f, "VpcIdSet")
	delete(f, "IpRestrictionEffect")
	delete(f, "IpSet")
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
type ModifyBackupPlanRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 实例最早开始备份时间
	MinBackupStartTime *string `json:"MinBackupStartTime,omitnil,omitempty" name:"MinBackupStartTime"`

	// 实例最晚开始备份时间
	MaxBackupStartTime *string `json:"MaxBackupStartTime,omitnil,omitempty" name:"MaxBackupStartTime"`

	// 实例备份保留时长，取值范围为7-1830，单位是天
	BaseBackupRetentionPeriod *uint64 `json:"BaseBackupRetentionPeriod,omitnil,omitempty" name:"BaseBackupRetentionPeriod"`

	// 实例备份周期，若是星期维度，格式为小写星期英文单词，且至少设置两天备份；若是按月维度，格式为数字字符，如["1","2"]。
	BackupPeriod []*string `json:"BackupPeriod,omitnil,omitempty" name:"BackupPeriod"`

	// 实例日志备份保留时长，取值范围为7-1830，单位是天
	LogBackupRetentionPeriod *uint64 `json:"LogBackupRetentionPeriod,omitnil,omitempty" name:"LogBackupRetentionPeriod"`

	// 备份计划ID，用于指明要修改哪个备份计划，不传则是修改默认备份计划。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 要修改的备份计划名称。
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`
}

type ModifyBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 实例最早开始备份时间
	MinBackupStartTime *string `json:"MinBackupStartTime,omitnil,omitempty" name:"MinBackupStartTime"`

	// 实例最晚开始备份时间
	MaxBackupStartTime *string `json:"MaxBackupStartTime,omitnil,omitempty" name:"MaxBackupStartTime"`

	// 实例备份保留时长，取值范围为7-1830，单位是天
	BaseBackupRetentionPeriod *uint64 `json:"BaseBackupRetentionPeriod,omitnil,omitempty" name:"BaseBackupRetentionPeriod"`

	// 实例备份周期，若是星期维度，格式为小写星期英文单词，且至少设置两天备份；若是按月维度，格式为数字字符，如["1","2"]。
	BackupPeriod []*string `json:"BackupPeriod,omitnil,omitempty" name:"BackupPeriod"`

	// 实例日志备份保留时长，取值范围为7-1830，单位是天
	LogBackupRetentionPeriod *uint64 `json:"LogBackupRetentionPeriod,omitnil,omitempty" name:"LogBackupRetentionPeriod"`

	// 备份计划ID，用于指明要修改哪个备份计划，不传则是修改默认备份计划。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 要修改的备份计划名称。
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`
}

func (r *ModifyBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "MinBackupStartTime")
	delete(f, "MaxBackupStartTime")
	delete(f, "BaseBackupRetentionPeriod")
	delete(f, "BackupPeriod")
	delete(f, "LogBackupRetentionPeriod")
	delete(f, "PlanId")
	delete(f, "PlanName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupPlanResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupPlanResponseParams `json:"Response"`
}

func (r *ModifyBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBaseBackupExpireTimeRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 数据备份ID。可通过[DescribeBaseBackups](https://cloud.tencent.com/document/api/409/89022)接口获取
	BaseBackupId *string `json:"BaseBackupId,omitnil,omitempty" name:"BaseBackupId"`

	// 新过期时间。
	NewExpireTime *string `json:"NewExpireTime,omitnil,omitempty" name:"NewExpireTime"`
}

type ModifyBaseBackupExpireTimeRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 数据备份ID。可通过[DescribeBaseBackups](https://cloud.tencent.com/document/api/409/89022)接口获取
	BaseBackupId *string `json:"BaseBackupId,omitnil,omitempty" name:"BaseBackupId"`

	// 新过期时间。
	NewExpireTime *string `json:"NewExpireTime,omitnil,omitempty" name:"NewExpireTime"`
}

func (r *ModifyBaseBackupExpireTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBaseBackupExpireTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "BaseBackupId")
	delete(f, "NewExpireTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBaseBackupExpireTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBaseBackupExpireTimeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBaseBackupExpireTimeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBaseBackupExpireTimeResponseParams `json:"Response"`
}

func (r *ModifyBaseBackupExpireTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBaseBackupExpireTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceChargeTypeRequestParams struct {
	// 实例ID，形如postgres-6fego161。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 实例计费类型，目前支持：
	// <li>PREPAID：预付费，即包年包月</li>
	// <li>POSTPAID_BY_HOUR：后付费，即按量计费</li>
	// 默认值：PREPAID
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 购买时长，单位：月。
	// <li>预付费：支持1,2,3,4,5,6,7,8,9,10,11,12,24,36</li>
	// <li>后付费：只支持1</li>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 续费标记：
	// <li>0：手动续费</li>
	// <li>1：自动续费</li>
	// 默认值：0
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 是否自动使用代金券：
	// <li>0：否</li>
	// <li>1：是</li>
	// 默认值：0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

type ModifyDBInstanceChargeTypeRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-6fego161。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 实例计费类型，目前支持：
	// <li>PREPAID：预付费，即包年包月</li>
	// <li>POSTPAID_BY_HOUR：后付费，即按量计费</li>
	// 默认值：PREPAID
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 购买时长，单位：月。
	// <li>预付费：支持1,2,3,4,5,6,7,8,9,10,11,12,24,36</li>
	// <li>后付费：只支持1</li>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 续费标记：
	// <li>0：手动续费</li>
	// <li>1：自动续费</li>
	// 默认值：0
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 是否自动使用代金券：
	// <li>0：否</li>
	// <li>1：是</li>
	// 默认值：0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

func (r *ModifyDBInstanceChargeTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceChargeTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "InstanceChargeType")
	delete(f, "Period")
	delete(f, "AutoRenewFlag")
	delete(f, "AutoVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceChargeTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceChargeTypeResponseParams struct {
	// 订单名
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceChargeTypeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceChargeTypeResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceChargeTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceChargeTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceDeletionProtectionRequestParams struct {
	// 实例 ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 开启或关闭实例删除保护。true - 开启 ；false - 关闭。
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

type ModifyDBInstanceDeletionProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 开启或关闭实例删除保护。true - 开启 ；false - 关闭。
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

func (r *ModifyDBInstanceDeletionProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceDeletionProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "DeletionProtection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceDeletionProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceDeletionProtectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceDeletionProtectionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceDeletionProtectionResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceDeletionProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceDeletionProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceDeploymentRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 实例节点部署信息，支持多可用区部署时需要指定每个节点的部署可用区信息。
	// 可用区信息可以通过调用 [DescribeZones](https://cloud.tencent.com/document/api/409/16769) 接口的返回值中的Zone字段来获取。
	DBNodeSet []*DBNode `json:"DBNodeSet,omitnil,omitempty" name:"DBNodeSet"`

	// 指定实例配置完成变更后的切换时间。
	// <li>0：立即切换 </li>
	// <li>1：指定时间切换</li>
	// <li>2：维护时间窗口内切换</li>
	SwitchTag *int64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// 切换开始时间，时间格式：HH:MM:SS，例如：01:00:00。当SwitchTag为0或2时，该参数失效。
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// 切换截止时间，时间格式：HH:MM:SS，例如：01:30:00。当SwitchTag为0或2时，该参数失效。
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`
}

type ModifyDBInstanceDeploymentRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 实例节点部署信息，支持多可用区部署时需要指定每个节点的部署可用区信息。
	// 可用区信息可以通过调用 [DescribeZones](https://cloud.tencent.com/document/api/409/16769) 接口的返回值中的Zone字段来获取。
	DBNodeSet []*DBNode `json:"DBNodeSet,omitnil,omitempty" name:"DBNodeSet"`

	// 指定实例配置完成变更后的切换时间。
	// <li>0：立即切换 </li>
	// <li>1：指定时间切换</li>
	// <li>2：维护时间窗口内切换</li>
	SwitchTag *int64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// 切换开始时间，时间格式：HH:MM:SS，例如：01:00:00。当SwitchTag为0或2时，该参数失效。
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// 切换截止时间，时间格式：HH:MM:SS，例如：01:30:00。当SwitchTag为0或2时，该参数失效。
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`
}

func (r *ModifyDBInstanceDeploymentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceDeploymentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "DBNodeSet")
	delete(f, "SwitchTag")
	delete(f, "SwitchStartTime")
	delete(f, "SwitchEndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceDeploymentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceDeploymentResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceDeploymentResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceDeploymentResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceDeploymentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceDeploymentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceHAConfigRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 主从同步方式：
	// <li>Semi-sync：半同步</li>
	// <li>Async：异步</li>
	SyncMode *string `json:"SyncMode,omitnil,omitempty" name:"SyncMode"`

	// 高可用备机最大延迟数据量。备节点延迟数据量小于等于该值，且备节点延迟时间小于等于MaxStandbyLag时，可以切换为主节点。
	// <li>单位：byte</li>
	// <li>参数范围：[1073741824, 322122547200]</li>
	MaxStandbyLatency *uint64 `json:"MaxStandbyLatency,omitnil,omitempty" name:"MaxStandbyLatency"`

	// 高可用备机最大延迟时间。备节点延迟时间小于等于该值，且备节点延迟数据量小于等于MaxStandbyLatency时，可以切换为主节点。
	// <li>单位：s</li>
	// <li>参数范围：[5, 10]</li>
	MaxStandbyLag *uint64 `json:"MaxStandbyLag,omitnil,omitempty" name:"MaxStandbyLag"`

	// 同步备机最大延迟数据量。备机延迟数据量小于等于该值，且该备机延迟时间小于等于MaxSyncStandbyLag时，则该备机采用同步复制；否则，采用异步复制。
	// 该参数值针对SyncMode设置为Semi-sync的实例有效。
	// 半同步实例禁止退化为异步复制时，不设置MaxSyncStandbyLatency、MaxSyncStandbyLag。
	// 半同步实例允许退化异步复制时，PostgreSQL 9版本的实例须设置MaxSyncStandbyLatency且不设置MaxSyncStandbyLag，PostgreSQL 10及以上版本的实例须设置MaxSyncStandbyLatency、MaxSyncStandbyLag。
	MaxSyncStandbyLatency *uint64 `json:"MaxSyncStandbyLatency,omitnil,omitempty" name:"MaxSyncStandbyLatency"`

	// 同步备机最大延迟时间。备机延迟时间小于等于该值，且该备机延迟数据量小于等于MaxSyncStandbyLatency时，则该备机采用同步复制；否则，采用异步复制。
	// 该参数值针对SyncMode设置为Semi-sync的实例有效。
	// 半同步实例禁止退化为异步复制时，不设置MaxSyncStandbyLatency、MaxSyncStandbyLag。
	// 半同步实例允许退化异步复制时，PostgreSQL 9版本的实例须设置MaxSyncStandbyLatency且不设置MaxSyncStandbyLag，PostgreSQL 10及以上版本的实例须设置MaxSyncStandbyLatency、MaxSyncStandbyLag，
	MaxSyncStandbyLag *uint64 `json:"MaxSyncStandbyLag,omitnil,omitempty" name:"MaxSyncStandbyLag"`
}

type ModifyDBInstanceHAConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 主从同步方式：
	// <li>Semi-sync：半同步</li>
	// <li>Async：异步</li>
	SyncMode *string `json:"SyncMode,omitnil,omitempty" name:"SyncMode"`

	// 高可用备机最大延迟数据量。备节点延迟数据量小于等于该值，且备节点延迟时间小于等于MaxStandbyLag时，可以切换为主节点。
	// <li>单位：byte</li>
	// <li>参数范围：[1073741824, 322122547200]</li>
	MaxStandbyLatency *uint64 `json:"MaxStandbyLatency,omitnil,omitempty" name:"MaxStandbyLatency"`

	// 高可用备机最大延迟时间。备节点延迟时间小于等于该值，且备节点延迟数据量小于等于MaxStandbyLatency时，可以切换为主节点。
	// <li>单位：s</li>
	// <li>参数范围：[5, 10]</li>
	MaxStandbyLag *uint64 `json:"MaxStandbyLag,omitnil,omitempty" name:"MaxStandbyLag"`

	// 同步备机最大延迟数据量。备机延迟数据量小于等于该值，且该备机延迟时间小于等于MaxSyncStandbyLag时，则该备机采用同步复制；否则，采用异步复制。
	// 该参数值针对SyncMode设置为Semi-sync的实例有效。
	// 半同步实例禁止退化为异步复制时，不设置MaxSyncStandbyLatency、MaxSyncStandbyLag。
	// 半同步实例允许退化异步复制时，PostgreSQL 9版本的实例须设置MaxSyncStandbyLatency且不设置MaxSyncStandbyLag，PostgreSQL 10及以上版本的实例须设置MaxSyncStandbyLatency、MaxSyncStandbyLag。
	MaxSyncStandbyLatency *uint64 `json:"MaxSyncStandbyLatency,omitnil,omitempty" name:"MaxSyncStandbyLatency"`

	// 同步备机最大延迟时间。备机延迟时间小于等于该值，且该备机延迟数据量小于等于MaxSyncStandbyLatency时，则该备机采用同步复制；否则，采用异步复制。
	// 该参数值针对SyncMode设置为Semi-sync的实例有效。
	// 半同步实例禁止退化为异步复制时，不设置MaxSyncStandbyLatency、MaxSyncStandbyLag。
	// 半同步实例允许退化异步复制时，PostgreSQL 9版本的实例须设置MaxSyncStandbyLatency且不设置MaxSyncStandbyLag，PostgreSQL 10及以上版本的实例须设置MaxSyncStandbyLatency、MaxSyncStandbyLag，
	MaxSyncStandbyLag *uint64 `json:"MaxSyncStandbyLag,omitnil,omitempty" name:"MaxSyncStandbyLag"`
}

func (r *ModifyDBInstanceHAConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceHAConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "SyncMode")
	delete(f, "MaxStandbyLatency")
	delete(f, "MaxStandbyLag")
	delete(f, "MaxSyncStandbyLatency")
	delete(f, "MaxSyncStandbyLag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceHAConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceHAConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceHAConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceHAConfigResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceHAConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceHAConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNameRequestParams struct {
	// 数据库实例ID，形如postgres-6fego161。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 实例名称，仅支持长度小于60的中文/英文/数字/"_"/"-"。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// 数据库实例ID，形如postgres-6fego161。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 实例名称，仅支持长度小于60的中文/英文/数字/"_"/"-"。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
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
	delete(f, "DBInstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNameResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyDBInstanceParametersRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 待修改参数及期望值。
	ParamList []*ParamEntry `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type ModifyDBInstanceParametersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 待修改参数及期望值。
	ParamList []*ParamEntry `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

func (r *ModifyDBInstanceParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceParametersResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceParametersResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceParametersResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceReadOnlyGroupRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 当前实例所在只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// 实例修改的目标只读组ID
	NewReadOnlyGroupId *string `json:"NewReadOnlyGroupId,omitnil,omitempty" name:"NewReadOnlyGroupId"`
}

type ModifyDBInstanceReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 当前实例所在只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// 实例修改的目标只读组ID
	NewReadOnlyGroupId *string `json:"NewReadOnlyGroupId,omitnil,omitempty" name:"NewReadOnlyGroupId"`
}

func (r *ModifyDBInstanceReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ReadOnlyGroupId")
	delete(f, "NewReadOnlyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceReadOnlyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceReadOnlyGroupResponseParams struct {
	// 流程ID，FlowId等同于TaskId
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceReadOnlyGroupResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSSLConfigRequestParams struct {
	// 实例 ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 开启或关闭SSL。true - 开启 ；false - 关闭。
	SSLEnabled *bool `json:"SSLEnabled,omitnil,omitempty" name:"SSLEnabled"`

	// SSL证书保护的唯一连接地址，若为主实例，可设置为内外网IP地址；若为只读实例，可设置为实例IP或只读组IP。在开启SSL或修改SSL保护的连接地址时，该参数为必传项；在关闭SSL时，该参数将被忽略。
	ConnectAddress *string `json:"ConnectAddress,omitnil,omitempty" name:"ConnectAddress"`
}

type ModifyDBInstanceSSLConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 开启或关闭SSL。true - 开启 ；false - 关闭。
	SSLEnabled *bool `json:"SSLEnabled,omitnil,omitempty" name:"SSLEnabled"`

	// SSL证书保护的唯一连接地址，若为主实例，可设置为内外网IP地址；若为只读实例，可设置为实例IP或只读组IP。在开启SSL或修改SSL保护的连接地址时，该参数为必传项；在关闭SSL时，该参数将被忽略。
	ConnectAddress *string `json:"ConnectAddress,omitnil,omitempty" name:"ConnectAddress"`
}

func (r *ModifyDBInstanceSSLConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSSLConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "SSLEnabled")
	delete(f, "ConnectAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSSLConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSSLConfigResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceSSLConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSSLConfigResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceSSLConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSSLConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// 实例或只读组要绑定的安全组列表。
	// 安全组信息可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来查询。
	// **注意：**该入参会全量替换存量已有集合，非增量更新。修改需传入预期的全量集合。
	SecurityGroupIdSet []*string `json:"SecurityGroupIdSet,omitnil,omitempty" name:"SecurityGroupIdSet"`

	// 实例ID，可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取。DBInstanceId和ReadOnlyGroupId至少传一个；如果都传，忽略ReadOnlyGroupId。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 只读组ID，可通过[DescribeReadOnlyGroups](https://cloud.tencent.com/document/api/409/52599)接口获取。DBInstanceId和ReadOnlyGroupId至少传一个；如果要修改只读组关联的安全组，只传ReadOnlyGroupId
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例或只读组要绑定的安全组列表。
	// 安全组信息可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来查询。
	// **注意：**该入参会全量替换存量已有集合，非增量更新。修改需传入预期的全量集合。
	SecurityGroupIdSet []*string `json:"SecurityGroupIdSet,omitnil,omitempty" name:"SecurityGroupIdSet"`

	// 实例ID，可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取。DBInstanceId和ReadOnlyGroupId至少传一个；如果都传，忽略ReadOnlyGroupId。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 只读组ID，可通过[DescribeReadOnlyGroups](https://cloud.tencent.com/document/api/409/52599)接口获取。DBInstanceId和ReadOnlyGroupId至少传一个；如果要修改只读组关联的安全组，只传ReadOnlyGroupId
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
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
	delete(f, "SecurityGroupIdSet")
	delete(f, "DBInstanceId")
	delete(f, "ReadOnlyGroupId")
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
type ModifyDBInstanceSpecRequestParams struct {
	// 实例ID，形如：postgres-6bwgamo3。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 修改后的实例内存大小，单位GiB。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 修改后的实例磁盘大小，单位GiB。该参数的设置步长为10。
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 是否自动使用代金券：
	// <li>0：否</li>
	// <li>1：是</li>
	// 默认值：0
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 活动ID。
	ActivityId *uint64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 指定实例配置完成变更后的切换时间。
	// <li>0：立即切换 </li>
	// <li>1：指定时间切换</li>
	// <li>2：维护时间窗口内</li>切换
	// 默认值：0 
	SwitchTag *uint64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// 切换开始时间，时间格式：HH:MM:SS，例如：01:00:00。当SwitchTag为0或2时，该参数失效。
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// 切换截止时间，时间格式：HH:MM:SS，例如：01:30:00。当SwitchTag为0或2时，该参数失效。
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`

	// 修改后的实例CPU大小，单位Core。不填写该参数时，默认根据Memory确定Cpu大小。如Memory为2，支持的规格有1核2GiB，则不传入Cpu时，Cpu默认为1。
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`
}

type ModifyDBInstanceSpecRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如：postgres-6bwgamo3。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 修改后的实例内存大小，单位GiB。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 修改后的实例磁盘大小，单位GiB。该参数的设置步长为10。
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 是否自动使用代金券：
	// <li>0：否</li>
	// <li>1：是</li>
	// 默认值：0
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 活动ID。
	ActivityId *uint64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 指定实例配置完成变更后的切换时间。
	// <li>0：立即切换 </li>
	// <li>1：指定时间切换</li>
	// <li>2：维护时间窗口内</li>切换
	// 默认值：0 
	SwitchTag *uint64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// 切换开始时间，时间格式：HH:MM:SS，例如：01:00:00。当SwitchTag为0或2时，该参数失效。
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// 切换截止时间，时间格式：HH:MM:SS，例如：01:30:00。当SwitchTag为0或2时，该参数失效。
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`

	// 修改后的实例CPU大小，单位Core。不填写该参数时，默认根据Memory确定Cpu大小。如Memory为2，支持的规格有1核2GiB，则不传入Cpu时，Cpu默认为1。
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`
}

func (r *ModifyDBInstanceSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "ActivityId")
	delete(f, "SwitchTag")
	delete(f, "SwitchStartTime")
	delete(f, "SwitchEndTime")
	delete(f, "Cpu")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSpecResponseParams struct {
	// 订单号。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 冻结流水号。
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceSpecResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSpecResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstancesProjectRequestParams struct {
	// 实例ID集合。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取。支持同时操作多个实例。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// 所属新项目的ID。可通过[DescribeProjects](https://cloud.tencent.com/document/api/651/78725)获取
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type ModifyDBInstancesProjectRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID集合。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取。支持同时操作多个实例。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// 所属新项目的ID。可通过[DescribeProjects](https://cloud.tencent.com/document/api/651/78725)获取
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *ModifyDBInstancesProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstancesProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceIdSet")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstancesProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstancesProjectResponseParams struct {
	// 转移项目成功的实例个数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstancesProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstancesProjectResponseParams `json:"Response"`
}

func (r *ModifyDBInstancesProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstancesProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseOwnerRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 数据库名称。可通过[DescribeDatabases](https://cloud.tencent.com/document/api/409/43353)接口获取
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据库新所有者。可通过[DescribeAccounts](https://cloud.tencent.com/document/api/409/18109)接口获取
	DatabaseOwner *string `json:"DatabaseOwner,omitnil,omitempty" name:"DatabaseOwner"`
}

type ModifyDatabaseOwnerRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 数据库名称。可通过[DescribeDatabases](https://cloud.tencent.com/document/api/409/43353)接口获取
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据库新所有者。可通过[DescribeAccounts](https://cloud.tencent.com/document/api/409/18109)接口获取
	DatabaseOwner *string `json:"DatabaseOwner,omitnil,omitempty" name:"DatabaseOwner"`
}

func (r *ModifyDatabaseOwnerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseOwnerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "DatabaseName")
	delete(f, "DatabaseOwner")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatabaseOwnerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseOwnerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDatabaseOwnerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDatabaseOwnerResponseParams `json:"Response"`
}

func (r *ModifyDatabaseOwnerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseOwnerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaintainTimeWindowRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 维护开始时间。时区为东八区（UTC+8）
	MaintainStartTime *string `json:"MaintainStartTime,omitnil,omitempty" name:"MaintainStartTime"`

	// 维护持续时间。单位：小时。取值范围：[1,4]
	MaintainDuration *uint64 `json:"MaintainDuration,omitnil,omitempty" name:"MaintainDuration"`

	// 维护周期
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitnil,omitempty" name:"MaintainWeekDays"`
}

type ModifyMaintainTimeWindowRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 维护开始时间。时区为东八区（UTC+8）
	MaintainStartTime *string `json:"MaintainStartTime,omitnil,omitempty" name:"MaintainStartTime"`

	// 维护持续时间。单位：小时。取值范围：[1,4]
	MaintainDuration *uint64 `json:"MaintainDuration,omitnil,omitempty" name:"MaintainDuration"`

	// 维护周期
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitnil,omitempty" name:"MaintainWeekDays"`
}

func (r *ModifyMaintainTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaintainTimeWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "MaintainStartTime")
	delete(f, "MaintainDuration")
	delete(f, "MaintainWeekDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMaintainTimeWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaintainTimeWindowResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMaintainTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMaintainTimeWindowResponseParams `json:"Response"`
}

func (r *ModifyMaintainTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaintainTimeWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyParameterTemplateRequestParams struct {
	// 参数模板ID，用于唯一确认参数模板，不可修改。可通过[DescribeParameterTemplates](https://cloud.tencent.com/document/api/409/84067)接口获取
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 参数模板名称，长度为1～60个字符，仅支持数字,英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@  注：若该字段为空    ，则保持原参数模板名称
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 参数模板描述，长度为0～60个字符，仅支持数字,英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@  注：若不传入该参数，则保持原参数模板描述
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// 需要修改或添加的参数集合，注：同一参数不能同时出现在修改添加集合和删除集合中
	ModifyParamEntrySet []*ParamEntry `json:"ModifyParamEntrySet,omitnil,omitempty" name:"ModifyParamEntrySet"`

	// 需要从模板中删除的参数集合，注：同一参数不能同时出现在修改添加集合和删除集合中
	DeleteParamSet []*string `json:"DeleteParamSet,omitnil,omitempty" name:"DeleteParamSet"`
}

type ModifyParameterTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板ID，用于唯一确认参数模板，不可修改。可通过[DescribeParameterTemplates](https://cloud.tencent.com/document/api/409/84067)接口获取
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 参数模板名称，长度为1～60个字符，仅支持数字,英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@  注：若该字段为空    ，则保持原参数模板名称
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 参数模板描述，长度为0～60个字符，仅支持数字,英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@  注：若不传入该参数，则保持原参数模板描述
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// 需要修改或添加的参数集合，注：同一参数不能同时出现在修改添加集合和删除集合中
	ModifyParamEntrySet []*ParamEntry `json:"ModifyParamEntrySet,omitnil,omitempty" name:"ModifyParamEntrySet"`

	// 需要从模板中删除的参数集合，注：同一参数不能同时出现在修改添加集合和删除集合中
	DeleteParamSet []*string `json:"DeleteParamSet,omitnil,omitempty" name:"DeleteParamSet"`
}

func (r *ModifyParameterTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyParameterTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TemplateName")
	delete(f, "TemplateDescription")
	delete(f, "ModifyParamEntrySet")
	delete(f, "DeleteParamSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyParameterTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyParameterTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyParameterTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyParameterTemplateResponseParams `json:"Response"`
}

func (r *ModifyParameterTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyParameterTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivilege struct {
	// 要修改的数据库对象及权限列表
	DatabasePrivilege *DatabasePrivilege `json:"DatabasePrivilege,omitnil,omitempty" name:"DatabasePrivilege"`

	// 修改的方式，当前仅支持grantObject、revokeObject、alterRole。grantObject代表授权、revokeObject代表收回权、alterRole代表修改账号类型。
	ModifyType *string `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`

	// 当ModifyType为revokeObject才需要此参数，参数为true时，撤销权限会级联撤销。默认为false。
	IsCascade *bool `json:"IsCascade,omitnil,omitempty" name:"IsCascade"`
}

// Predefined struct for user
type ModifyReadOnlyDBInstanceWeightRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 只读组ID。可通过[DescribeReadOnlyGroups](https://cloud.tencent.com/document/api/409/52599)接口获取
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// 只读实例在只读组中的流量权重(1-50)
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type ModifyReadOnlyDBInstanceWeightRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 只读组ID。可通过[DescribeReadOnlyGroups](https://cloud.tencent.com/document/api/409/52599)接口获取
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// 只读实例在只读组中的流量权重(1-50)
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

func (r *ModifyReadOnlyDBInstanceWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReadOnlyDBInstanceWeightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ReadOnlyGroupId")
	delete(f, "Weight")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyReadOnlyDBInstanceWeightRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyReadOnlyDBInstanceWeightResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyReadOnlyDBInstanceWeightResponse struct {
	*tchttp.BaseResponse
	Response *ModifyReadOnlyDBInstanceWeightResponseParams `json:"Response"`
}

func (r *ModifyReadOnlyDBInstanceWeightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReadOnlyDBInstanceWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyReadOnlyGroupConfigRequestParams struct {
	// 只读组ID。
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// 只读组名称。仅支持长度小于60的中文/英文/数字/"_"/"-"
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitnil,omitempty" name:"ReadOnlyGroupName"`

	// 延迟时间配置开关：0关、1开
	ReplayLagEliminate *uint64 `json:"ReplayLagEliminate,omitnil,omitempty" name:"ReplayLagEliminate"`

	// 延迟日志大小配置开关：0关、1开
	ReplayLatencyEliminate *uint64 `json:"ReplayLatencyEliminate,omitnil,omitempty" name:"ReplayLatencyEliminate"`

	// 延迟日志大小阈值，单位MB。当开启延迟日志大小配置，应输入正整数
	MaxReplayLatency *uint64 `json:"MaxReplayLatency,omitnil,omitempty" name:"MaxReplayLatency"`

	// 延迟时间大小阈值，单位s。当开启延迟时间配置时，应输入正整数。
	MaxReplayLag *uint64 `json:"MaxReplayLag,omitnil,omitempty" name:"MaxReplayLag"`

	// 自动负载均衡开关：0关、1开
	Rebalance *uint64 `json:"Rebalance,omitnil,omitempty" name:"Rebalance"`

	// 延迟剔除最小保留实例数。取值范围[0,100]
	MinDelayEliminateReserve *uint64 `json:"MinDelayEliminateReserve,omitnil,omitempty" name:"MinDelayEliminateReserve"`
}

type ModifyReadOnlyGroupConfigRequest struct {
	*tchttp.BaseRequest
	
	// 只读组ID。
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// 只读组名称。仅支持长度小于60的中文/英文/数字/"_"/"-"
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitnil,omitempty" name:"ReadOnlyGroupName"`

	// 延迟时间配置开关：0关、1开
	ReplayLagEliminate *uint64 `json:"ReplayLagEliminate,omitnil,omitempty" name:"ReplayLagEliminate"`

	// 延迟日志大小配置开关：0关、1开
	ReplayLatencyEliminate *uint64 `json:"ReplayLatencyEliminate,omitnil,omitempty" name:"ReplayLatencyEliminate"`

	// 延迟日志大小阈值，单位MB。当开启延迟日志大小配置，应输入正整数
	MaxReplayLatency *uint64 `json:"MaxReplayLatency,omitnil,omitempty" name:"MaxReplayLatency"`

	// 延迟时间大小阈值，单位s。当开启延迟时间配置时，应输入正整数。
	MaxReplayLag *uint64 `json:"MaxReplayLag,omitnil,omitempty" name:"MaxReplayLag"`

	// 自动负载均衡开关：0关、1开
	Rebalance *uint64 `json:"Rebalance,omitnil,omitempty" name:"Rebalance"`

	// 延迟剔除最小保留实例数。取值范围[0,100]
	MinDelayEliminateReserve *uint64 `json:"MinDelayEliminateReserve,omitnil,omitempty" name:"MinDelayEliminateReserve"`
}

func (r *ModifyReadOnlyGroupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReadOnlyGroupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReadOnlyGroupId")
	delete(f, "ReadOnlyGroupName")
	delete(f, "ReplayLagEliminate")
	delete(f, "ReplayLatencyEliminate")
	delete(f, "MaxReplayLatency")
	delete(f, "MaxReplayLag")
	delete(f, "Rebalance")
	delete(f, "MinDelayEliminateReserve")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyReadOnlyGroupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyReadOnlyGroupConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyReadOnlyGroupConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyReadOnlyGroupConfigResponseParams `json:"Response"`
}

func (r *ModifyReadOnlyGroupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReadOnlyGroupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySwitchTimePeriodRequestParams struct {
	// 处于等待切换状态中的实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 入参取值为 0 ，代表立即切换。
	SwitchTag *uint64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`
}

type ModifySwitchTimePeriodRequest struct {
	*tchttp.BaseRequest
	
	// 处于等待切换状态中的实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 入参取值为 0 ，代表立即切换。
	SwitchTag *uint64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`
}

func (r *ModifySwitchTimePeriodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySwitchTimePeriodRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "SwitchTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySwitchTimePeriodRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySwitchTimePeriodResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySwitchTimePeriodResponse struct {
	*tchttp.BaseResponse
	Response *ModifySwitchTimePeriodResponseParams `json:"Response"`
}

func (r *ModifySwitchTimePeriodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySwitchTimePeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkAccess struct {
	// 网络资源id，实例id或RO组id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源类型，1-实例 2-RO组
	ResourceType *uint64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// IPV4地址
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// IPV6地址
	Vip6 *string `json:"Vip6,omitnil,omitempty" name:"Vip6"`

	// 访问端口
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 网络状态，1-申请中，2-使用中，3-删除中，4-已删除
	VpcStatus *int64 `json:"VpcStatus,omitnil,omitempty" name:"VpcStatus"`
}

// Predefined struct for user
type OpenAccountCAMRequestParams struct {
	// 数据库实例ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 需要开启CAM服务的账号名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

type OpenAccountCAMRequest struct {
	*tchttp.BaseRequest
	
	// 数据库实例ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 需要开启CAM服务的账号名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

func (r *OpenAccountCAMRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAccountCAMRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "UserName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenAccountCAMRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenAccountCAMResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenAccountCAMResponse struct {
	*tchttp.BaseResponse
	Response *OpenAccountCAMResponseParams `json:"Response"`
}

func (r *OpenAccountCAMResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAccountCAMResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenDBExtranetAccessRequestParams struct {
	// 实例ID，形如postgres-hez4fh0v。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 是否开通Ipv6外网，1：是，0：否
	// 默认值：0
	IsIpv6 *int64 `json:"IsIpv6,omitnil,omitempty" name:"IsIpv6"`
}

type OpenDBExtranetAccessRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-hez4fh0v。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 是否开通Ipv6外网，1：是，0：否
	// 默认值：0
	IsIpv6 *int64 `json:"IsIpv6,omitnil,omitempty" name:"IsIpv6"`
}

func (r *OpenDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "IsIpv6")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenDBExtranetAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenDBExtranetAccessResponseParams struct {
	// 流程ID，FlowId等同于TaskId
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *OpenDBExtranetAccessResponseParams `json:"Response"`
}

func (r *OpenDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamEntry struct {
	// 参数名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 修改参数值。入参均以字符串形式传递，例如：小数”0.1“、整数”1000“、枚举”replica“
	ExpectedValue *string `json:"ExpectedValue,omitnil,omitempty" name:"ExpectedValue"`
}

type ParamInfo struct {
	// 参数ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 参数名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数值类型：integer（整型）、real（浮点型）、bool（布尔型）、enum（枚举类型）、mutil_enum（枚举类型、支持多选）。
	// 当参数类型为integer（整型）、real（浮点型）时，参数的取值范围根据返回值的Max、Min确定； 
	// 当参数类型为bool（布尔型）时，参数设置值取值范围是true | false； 
	// 当参数类型为enum（枚举类型）、mutil_enum（多枚举类型）时，参数的取值范围由返回值中的EnumValue确定。
	ParamValueType *string `json:"ParamValueType,omitnil,omitempty" name:"ParamValueType"`

	// 参数值 单位。参数没有单位时，该字段返回空
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 参数默认值。以字符串形式返回
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 参数当前运行值。以字符串形式返回
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 数值类型（integer、real）参数，取值下界
	Max *float64 `json:"Max,omitnil,omitempty" name:"Max"`

	// 枚举类型参数，取值范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// 数值类型（integer、real）参数，取值上界
	Min *float64 `json:"Min,omitnil,omitempty" name:"Min"`

	// 参数中文描述
	ParamDescriptionCH *string `json:"ParamDescriptionCH,omitnil,omitempty" name:"ParamDescriptionCH"`

	// 参数英文描述
	ParamDescriptionEN *string `json:"ParamDescriptionEN,omitnil,omitempty" name:"ParamDescriptionEN"`

	// 参数修改，是否重启生效。（true为需要，false为不需要）
	NeedReboot *bool `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// 参数中文分类
	ClassificationCN *string `json:"ClassificationCN,omitnil,omitempty" name:"ClassificationCN"`

	// 参数英文分类
	ClassificationEN *string `json:"ClassificationEN,omitnil,omitempty" name:"ClassificationEN"`

	// 是否和规格相关。（true为相关，false为不想关）
	SpecRelated *bool `json:"SpecRelated,omitnil,omitempty" name:"SpecRelated"`

	// 是否为重点参数。（true为重点参数，修改是需要重点关注，可能会影响实例性能）
	Advanced *bool `json:"Advanced,omitnil,omitempty" name:"Advanced"`

	// 参数最后一次修改时间
	LastModifyTime *string `json:"LastModifyTime,omitnil,omitempty" name:"LastModifyTime"`

	// 参数主备制约，0：无主备制约关系，1:备机参数值需比主机大，2:主机参数值需比备机大
	StandbyRelated *int64 `json:"StandbyRelated,omitnil,omitempty" name:"StandbyRelated"`

	// 参数版本关联信息，内容为相应内核版本下的参数详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionRelationSet []*ParamVersionRelation `json:"VersionRelationSet,omitnil,omitempty" name:"VersionRelationSet"`

	// 参数规格关联信息，内容为相应规格下的参数详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecRelationSet []*ParamSpecRelation `json:"SpecRelationSet,omitnil,omitempty" name:"SpecRelationSet"`
}

type ParamSpecRelation struct {
	// 参数名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数信息所属规格
	Memory *string `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 参数在该规格下的默认值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 参数值单位。参数没有单位时，该字段返回空
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 数值类型（integer、real）参数，取值上界
	Max *float64 `json:"Max,omitnil,omitempty" name:"Max"`

	// 数值类型（integer、real）参数，取值下界
	Min *float64 `json:"Min,omitnil,omitempty" name:"Min"`

	// 枚举类型参数，取值范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`
}

type ParamVersionRelation struct {
	// 参数名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数信息所属内核版本
	DBKernelVersion *string `json:"DBKernelVersion,omitnil,omitempty" name:"DBKernelVersion"`

	// 参数在该版本该规格下的默认值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 参数值单位。参数没有单位时，该字段返回空
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 数值类型（integer、real）参数，取值上界
	Max *float64 `json:"Max,omitnil,omitempty" name:"Max"`

	// 数值类型（integer、real）参数，取值下界
	Min *float64 `json:"Min,omitnil,omitempty" name:"Min"`

	// 枚举类型参数，取值范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`
}

type ParameterTemplate struct {
	// 参数模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 参数模板名称
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 参数模板适用的数据库版本
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// 参数模板适用的数据库引擎
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// 参数模板描述
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`
}

type PgDeal struct {
	// 订单名
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 所属用户
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 订单涉及多少个实例
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 付费模式。1-预付费；0-后付费
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 异步任务流程ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 实例ID数组
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`
}

type PolicyRule struct {
	// 策略，ACCEPT 或者 DROP
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 来源或目的 IP 或 IP 段，例如172.16.0.0/12
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// 端口
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// 网络协议，支持 UDP、TCP 等
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// 规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type RawSlowQuery struct {
	// 慢SQL 语句
	RawQuery *string `json:"RawQuery,omitnil,omitempty" name:"RawQuery"`

	// 慢SQL 查询的数据库
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 慢SQL执行 耗时
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 执行慢SQL的客户端
	ClientAddr *string `json:"ClientAddr,omitnil,omitempty" name:"ClientAddr"`

	// 执行慢SQL的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 慢SQL执行的开始时间
	SessionStartTime *string `json:"SessionStartTime,omitnil,omitempty" name:"SessionStartTime"`
}

type ReadOnlyGroup struct {
	// 只读组标识
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// 只读组名字
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitnil,omitempty" name:"ReadOnlyGroupName"`

	// 项目id
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 主实例id
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitnil,omitempty" name:"MasterDBInstanceId"`

	// 最小保留实例数
	MinDelayEliminateReserve *int64 `json:"MinDelayEliminateReserve,omitnil,omitempty" name:"MinDelayEliminateReserve"`

	// 延迟空间大小阈值。单位MB。
	MaxReplayLatency *int64 `json:"MaxReplayLatency,omitnil,omitempty" name:"MaxReplayLatency"`

	// 延迟大小开关。0 - 关闭； 1 - 开启。
	ReplayLatencyEliminate *int64 `json:"ReplayLatencyEliminate,omitnil,omitempty" name:"ReplayLatencyEliminate"`

	// 延迟时间大小阈值，单位：秒。
	MaxReplayLag *float64 `json:"MaxReplayLag,omitnil,omitempty" name:"MaxReplayLag"`

	// 延迟时间开关。0 - 关闭； 1 - 开启。
	ReplayLagEliminate *int64 `json:"ReplayLagEliminate,omitnil,omitempty" name:"ReplayLagEliminate"`

	// 虚拟网络id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 地域id
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 地区id
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 状态。枚举值：creating、ok、modifying、deleting、deleted
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例详细信息
	ReadOnlyDBInstanceList []*DBInstance `json:"ReadOnlyDBInstanceList,omitnil,omitempty" name:"ReadOnlyDBInstanceList"`

	// 自动负载均衡开关
	Rebalance *int64 `json:"Rebalance,omitnil,omitempty" name:"Rebalance"`

	// 网络信息
	DBInstanceNetInfo []*DBInstanceNetInfo `json:"DBInstanceNetInfo,omitnil,omitempty" name:"DBInstanceNetInfo"`

	// 只读组网络信息列表（此字段已废弃）
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkAccessList []*NetworkAccess `json:"NetworkAccessList,omitnil,omitempty" name:"NetworkAccessList"`
}

// Predefined struct for user
type RebalanceReadOnlyGroupRequestParams struct {
	// 只读组ID。可通过[DescribeReadOnlyGroups](https://cloud.tencent.com/document/api/409/52599)接口获取
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

type RebalanceReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 只读组ID。可通过[DescribeReadOnlyGroups](https://cloud.tencent.com/document/api/409/52599)接口获取
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

func (r *RebalanceReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebalanceReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RebalanceReadOnlyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RebalanceReadOnlyGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RebalanceReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *RebalanceReadOnlyGroupResponseParams `json:"Response"`
}

func (r *RebalanceReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebalanceReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefreshAccountPasswordRequestParams struct {
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 账号名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

type RefreshAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 账号名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

func (r *RefreshAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "UserName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefreshAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefreshAccountPasswordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RefreshAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *RefreshAccountPasswordResponseParams `json:"Response"`
}

func (r *RefreshAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// 该地域对应的英文名称
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 该地域对应的中文名称
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 该地域对应的数字编号
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 可用状态，UNAVAILABLE表示不可用，AVAILABLE表示可用
	RegionState *string `json:"RegionState,omitnil,omitempty" name:"RegionState"`

	// 该地域是否支持国际站售卖，0：不支持，1：支持
	SupportInternational *uint64 `json:"SupportInternational,omitnil,omitempty" name:"SupportInternational"`
}

// Predefined struct for user
type RemoveDBInstanceFromReadOnlyGroupRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 只读组ID。可通过[DescribeReadOnlyGroups](https://cloud.tencent.com/document/api/409/52599)接口获取
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

type RemoveDBInstanceFromReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 只读组ID。可通过[DescribeReadOnlyGroups](https://cloud.tencent.com/document/api/409/52599)接口获取
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

func (r *RemoveDBInstanceFromReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveDBInstanceFromReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveDBInstanceFromReadOnlyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveDBInstanceFromReadOnlyGroupResponseParams struct {
	// 流程ID，FlowId等同于TaskId
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveDBInstanceFromReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *RemoveDBInstanceFromReadOnlyGroupResponseParams `json:"Response"`
}

func (r *RemoveDBInstanceFromReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveDBInstanceFromReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstanceRequestParams struct {
	// 实例ID，形如postgres-6fego161。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取。仅支持预付费（包年包月）实例。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 购买时长，单位：月。
	// <li>预付费：支持1,2,3,4,5,6,7,8,9,10,11,12,24,36</li>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 是否自动使用代金券：
	// <li>0：否</li>
	// <li>1：是</li>
	// 默认值：0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`
}

type RenewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-6fego161。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取。仅支持预付费（包年包月）实例。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 购买时长，单位：月。
	// <li>预付费：支持1,2,3,4,5,6,7,8,9,10,11,12,24,36</li>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 是否自动使用代金券：
	// <li>0：否</li>
	// <li>1：是</li>
	// 默认值：0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`
}

func (r *RenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "Period")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstanceResponseParams struct {
	// 订单名
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RenewInstanceResponseParams `json:"Response"`
}

func (r *RenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAccountPasswordRequestParams struct {
	// 实例ID，形如postgres-4wdeb0zv。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 实例账户名。可通过[DescribeAccounts](https://cloud.tencent.com/document/api/409/18109)接口获取
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// UserName账户对应的新密码。
	// 密码设置规则如下：
	// - 长度8~ 32位，推荐使用12位以上的密码
	// - 不能以" / "开头
	// - 必须包含以下四项:
	//   1.    小写字母a ~ z
	//   2.    大写字母 A ～ Z
	//   3.    数字 0 ～ 9
	//   4.    特殊字符 ()`~!@#$%^&*-+=_|{}[]:<>,.?/
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-4wdeb0zv。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 实例账户名。可通过[DescribeAccounts](https://cloud.tencent.com/document/api/409/18109)接口获取
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// UserName账户对应的新密码。
	// 密码设置规则如下：
	// - 长度8~ 32位，推荐使用12位以上的密码
	// - 不能以" / "开头
	// - 必须包含以下四项:
	//   1.    小写字母a ~ z
	//   2.    大写字母 A ～ Z
	//   3.    数字 0 ～ 9
	//   4.    特殊字符 ()`~!@#$%^&*-+=_|{}[]:<>,.?/
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
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
	delete(f, "DBInstanceId")
	delete(f, "UserName")
	delete(f, "Password")
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

// Predefined struct for user
type RestartDBInstanceRequestParams struct {
	// 实例ID，形如postgres-6r233v55。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type RestartDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-6r233v55。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *RestartDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartDBInstanceResponseParams struct {
	// 流程ID，FlowId等同于TaskId
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RestartDBInstanceResponseParams `json:"Response"`
}

func (r *RestartDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreDBInstanceObjectsRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 需要恢复的对象列表。假设需要恢复的对象名为user，则恢复后的名称为user_bak_${LinuxTime}。${LinuxTime}无法指定，由系统根据任务发起的linux时间设定。
	RestoreObjects []*string `json:"RestoreObjects,omitnil,omitempty" name:"RestoreObjects"`

	// 恢复所用备份集。BackupSetId与RestoreTargetTime有且只能传一个。可通过[DescribeBaseBackups](https://cloud.tencent.com/document/api/409/89022)接口获取
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// 恢复目标时间，北京时间。BackupSetId与RestoreTargetTime有且只能传一个。
	RestoreTargetTime *string `json:"RestoreTargetTime,omitnil,omitempty" name:"RestoreTargetTime"`
}

type RestoreDBInstanceObjectsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 需要恢复的对象列表。假设需要恢复的对象名为user，则恢复后的名称为user_bak_${LinuxTime}。${LinuxTime}无法指定，由系统根据任务发起的linux时间设定。
	RestoreObjects []*string `json:"RestoreObjects,omitnil,omitempty" name:"RestoreObjects"`

	// 恢复所用备份集。BackupSetId与RestoreTargetTime有且只能传一个。可通过[DescribeBaseBackups](https://cloud.tencent.com/document/api/409/89022)接口获取
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// 恢复目标时间，北京时间。BackupSetId与RestoreTargetTime有且只能传一个。
	RestoreTargetTime *string `json:"RestoreTargetTime,omitnil,omitempty" name:"RestoreTargetTime"`
}

func (r *RestoreDBInstanceObjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreDBInstanceObjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "RestoreObjects")
	delete(f, "BackupSetId")
	delete(f, "RestoreTargetTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestoreDBInstanceObjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreDBInstanceObjectsResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestoreDBInstanceObjectsResponse struct {
	*tchttp.BaseResponse
	Response *RestoreDBInstanceObjectsResponseParams `json:"Response"`
}

func (r *RestoreDBInstanceObjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreDBInstanceObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroup struct {
	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 创建时间
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
	SecurityGroupDescription *string `json:"SecurityGroupDescription,omitnil,omitempty" name:"SecurityGroupDescription"`
}

// Predefined struct for user
type SetAutoRenewFlagRequestParams struct {
	// 实例ID集合。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取。仅支持预付费（包年包月）的实例。支持同时操作多个实例。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// 续费标记。0-正常续费；1-自动续费；2-到期不续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

type SetAutoRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID集合。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取。仅支持预付费（包年包月）的实例。支持同时操作多个实例。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// 续费标记。0-正常续费；1-自动续费；2-到期不续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

func (r *SetAutoRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAutoRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceIdSet")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetAutoRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetAutoRenewFlagResponseParams struct {
	// 设置成功的实例个数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetAutoRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *SetAutoRenewFlagResponseParams `json:"Response"`
}

func (r *SetAutoRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAutoRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpecInfo struct {
	// 地域英文编码，对应RegionSet的Region字段
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 区域英文编码，对应ZoneSet的Zone字段
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 规格详细信息列表
	SpecItemInfoList []*SpecItemInfo `json:"SpecItemInfoList,omitnil,omitempty" name:"SpecItemInfoList"`

	// 支持KMS的地域
	SupportKMSRegions []*string `json:"SupportKMSRegions,omitnil,omitempty" name:"SupportKMSRegions"`
}

type SpecItemInfo struct {
	// 规格ID
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// PostgreSQL的版本编号
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 内核编号对应的完整版本名称
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// CPU核数
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存大小，单位：MB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 该规格所支持最大存储容量，单位：GB
	MaxStorage *uint64 `json:"MaxStorage,omitnil,omitempty" name:"MaxStorage"`

	// 该规格所支持最小存储容量，单位：GB
	MinStorage *uint64 `json:"MinStorage,omitnil,omitempty" name:"MinStorage"`

	// 该规格的预估QPS
	Qps *uint64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 【该字段废弃】
	Pid *uint64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// 机器类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// PostgreSQL的主要版本编号
	MajorVersion *string `json:"MajorVersion,omitnil,omitempty" name:"MajorVersion"`

	// PostgreSQL的内核版本编号
	KernelVersion *string `json:"KernelVersion,omitnil,omitempty" name:"KernelVersion"`

	// 是否支持TDE数据加密功能，0-不支持，1-支持
	IsSupportTDE *int64 `json:"IsSupportTDE,omitnil,omitempty" name:"IsSupportTDE"`
}

// Predefined struct for user
type SwitchDBInstancePrimaryRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 是否强制切换。强制切换时只要备节点可访问，无论主备延迟多大都会发起切换。只有SwitchTag为0时，才可使用立即切换。
	// <li>默认：false</li>
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`

	// 指定实例配置完成变更后的切换时间。
	// <li>0：立即切换 </li>
	// <li>1：指定时间切换</li>
	// <li>2：维护时间窗口内切换</li>
	// 默认值：0 
	SwitchTag *uint64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// 切换开始时间，时间格式：HH:MM:SS，例如：01:00:00。当SwitchTag为0或2时，该参数失效。
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// 切换截止时间，时间格式：HH:MM:SS，例如：01:30:00。当SwitchTag为0或2时，该参数失效。SwitchStartTime和SwitchEndTime时间窗口不能小于30分钟。
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`
}

type SwitchDBInstancePrimaryRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 是否强制切换。强制切换时只要备节点可访问，无论主备延迟多大都会发起切换。只有SwitchTag为0时，才可使用立即切换。
	// <li>默认：false</li>
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`

	// 指定实例配置完成变更后的切换时间。
	// <li>0：立即切换 </li>
	// <li>1：指定时间切换</li>
	// <li>2：维护时间窗口内切换</li>
	// 默认值：0 
	SwitchTag *uint64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// 切换开始时间，时间格式：HH:MM:SS，例如：01:00:00。当SwitchTag为0或2时，该参数失效。
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// 切换截止时间，时间格式：HH:MM:SS，例如：01:30:00。当SwitchTag为0或2时，该参数失效。SwitchStartTime和SwitchEndTime时间窗口不能小于30分钟。
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`
}

func (r *SwitchDBInstancePrimaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDBInstancePrimaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "Force")
	delete(f, "SwitchTag")
	delete(f, "SwitchStartTime")
	delete(f, "SwitchEndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchDBInstancePrimaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDBInstancePrimaryResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchDBInstancePrimaryResponse struct {
	*tchttp.BaseResponse
	Response *SwitchDBInstancePrimaryResponseParams `json:"Response"`
}

func (r *SwitchDBInstancePrimaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDBInstancePrimaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TaskDetail struct {
	// 当前执行的子任务步骤名称。
	CurrentStep *string `json:"CurrentStep,omitnil,omitempty" name:"CurrentStep"`

	// 当前任务所拥有的子步骤描述。
	AllSteps *string `json:"AllSteps,omitnil,omitempty" name:"AllSteps"`

	// 任务的输入参数。
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// 任务的输出参数。
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// 指定实例配置完成变更后的切换时间，默认值：0
	// 0:   此任务不需要切换
	// 1：立即切换
	// 2：指定时间切换
	// 3：维护时间窗口内切换。
	SwitchTag *uint64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// 指定的切换时间。
	SwitchTime *string `json:"SwitchTime,omitnil,omitempty" name:"SwitchTime"`

	// 任务的提示信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type TaskSet struct {
	// 任务ID。
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务的类型。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务实例的实例ID。
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 任务的开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务的结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务的运行状态，包括Running,Success,WaitSwitch,Fail,Pause。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务的执行进度，取值范围0-100。
	Progress *uint64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 任务的详情信息
	TaskDetail *TaskDetail `json:"TaskDetail,omitnil,omitempty" name:"TaskDetail"`
}

// Predefined struct for user
type UnlockAccountRequestParams struct {
	// 实例ID。	
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 账号名称。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

type UnlockAccountRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。	
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 账号名称。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

func (r *UnlockAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnlockAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "UserName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnlockAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnlockAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnlockAccountResponse struct {
	*tchttp.BaseResponse
	Response *UnlockAccountResponseParams `json:"Response"`
}

func (r *UnlockAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnlockAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceKernelVersionRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 升级的目标内核版本号。可以通过接口[DescribeDBVersions](https://cloud.tencent.com/document/api/409/89018)的返回字段AvailableUpgradeTarget获取。
	TargetDBKernelVersion *string `json:"TargetDBKernelVersion,omitnil,omitempty" name:"TargetDBKernelVersion"`

	// 指定实例升级内核版本号完成后的切换时间。可选值:
	// <li>0：立即切换</li>
	// <li>1：指定时间切换</li>
	// <li>2：维护时间窗口内切换</li>
	// 默认值：0 
	SwitchTag *uint64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// 切换开始时间，时间格式：HH:MM:SS，例如：01:00:00。当SwitchTag为0或2时，该参数失效。
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// 切换截止时间，时间格式：HH:MM:SS，例如：01:30:00。当SwitchTag为0或2时，该参数失效。SwitchStartTime和SwitchEndTime时间窗口不能小于30分钟。
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`

	// 是否对本次升级实例内核版本号操作执行预检查。
	// <li>true：执行预检查操作，不升级内核版本号。检查项目包含请求参数、内核版本号兼容性、实例参数等。</li>
	// <li>false：发送正常请求（默认值），通过检查后直接升级内核版本号。</li>
	// 默认值：false
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type UpgradeDBInstanceKernelVersionRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 升级的目标内核版本号。可以通过接口[DescribeDBVersions](https://cloud.tencent.com/document/api/409/89018)的返回字段AvailableUpgradeTarget获取。
	TargetDBKernelVersion *string `json:"TargetDBKernelVersion,omitnil,omitempty" name:"TargetDBKernelVersion"`

	// 指定实例升级内核版本号完成后的切换时间。可选值:
	// <li>0：立即切换</li>
	// <li>1：指定时间切换</li>
	// <li>2：维护时间窗口内切换</li>
	// 默认值：0 
	SwitchTag *uint64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// 切换开始时间，时间格式：HH:MM:SS，例如：01:00:00。当SwitchTag为0或2时，该参数失效。
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// 切换截止时间，时间格式：HH:MM:SS，例如：01:30:00。当SwitchTag为0或2时，该参数失效。SwitchStartTime和SwitchEndTime时间窗口不能小于30分钟。
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`

	// 是否对本次升级实例内核版本号操作执行预检查。
	// <li>true：执行预检查操作，不升级内核版本号。检查项目包含请求参数、内核版本号兼容性、实例参数等。</li>
	// <li>false：发送正常请求（默认值），通过检查后直接升级内核版本号。</li>
	// 默认值：false
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *UpgradeDBInstanceKernelVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceKernelVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "TargetDBKernelVersion")
	delete(f, "SwitchTag")
	delete(f, "SwitchStartTime")
	delete(f, "SwitchEndTime")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceKernelVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceKernelVersionResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeDBInstanceKernelVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDBInstanceKernelVersionResponseParams `json:"Response"`
}

func (r *UpgradeDBInstanceKernelVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceKernelVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceMajorVersionRequestParams struct {
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 目标内核版本号，可以通过API [DescribeDBVersions](https://cloud.tencent.com/document/product/409/89018)获取可以升级的目标内核版本号。
	TargetDBKernelVersion *string `json:"TargetDBKernelVersion,omitnil,omitempty" name:"TargetDBKernelVersion"`

	// 是否为校验模式，若UpgradeCheck为True，表示仅进行内核版本兼容性检查，不会进行实质性的升级操作，对原实例无影响。检查结果可以通过升级日志查看。
	UpgradeCheck *bool `json:"UpgradeCheck,omitnil,omitempty" name:"UpgradeCheck"`

	// 升级前备份选项。True，表示升级前需要创建全量备份，False，表示升级前不需要创建全量备份。当实例已有备份集可以恢复到升级前的状态时，可选择False，否则需要指定为True。UpgradeCheck为True时，此参数无效。
	BackupBeforeUpgrade *bool `json:"BackupBeforeUpgrade,omitnil,omitempty" name:"BackupBeforeUpgrade"`

	// 统计信息收集选项，对主例运行 ANALYZE 以在升级后更新系统统计信息。可选值包括，
	// 0：不需要收集统计信息；
	// 1：实例恢复写之前收集统计信息；
	// 3：实例恢复写之后收集统计信息。
	// UpgradeCheck为True时，此参数无效。
	StatisticsRefreshOption *int64 `json:"StatisticsRefreshOption,omitnil,omitempty" name:"StatisticsRefreshOption"`

	// 插件升级选项，pg_upgrade不会升级任何插件，需要在升级完成后在创建过插件的库上执行"ALTER EXTENSION UPDATE"。发起升级实例大版本时可以指定在实例恢复写前/后是否需要升级任务自动升级插件版本。可选值包括：
	// 0：不需要自动升级插件；
	// 1：恢复写之前升级插件；
	// 2：恢复写之后升级插件。
	// UpgradeCheck为True时，此参数无效。
	ExtensionUpgradeOption *int64 `json:"ExtensionUpgradeOption,omitnil,omitempty" name:"ExtensionUpgradeOption"`

	// 升级时间选项，升级过程中会有一段时间实例只读，并会有一次秒级闪断，发起升级时需要选择这段影响的时间窗。可选值包括：
	// 0：自动执行，不需要指定时间窗；
	// 1：指定本次升级任务的时间窗，通过参数UpgradeTimeBegin和UpgradeTimeEnd设置；
	// 2：在实例运维时间窗内执行。
	// UpgradeCheck为True时，此参数无效。
	UpgradeTimeOption *int64 `json:"UpgradeTimeOption,omitnil,omitempty" name:"UpgradeTimeOption"`

	// 升级时间窗开始时间，时间格式：HH:MM:SS，例如：01:00:00。当UpgradeTimeOption为1时，该参数有效。
	// UpgradeCheck为True时，此参数无效。
	UpgradeTimeBegin *string `json:"UpgradeTimeBegin,omitnil,omitempty" name:"UpgradeTimeBegin"`

	// 升级时间窗截止时间，时间格式：HH:MM:SS，例如：02:00:00。当UpgradeTimeOption为1时，该参数有效。
	// UpgradeCheck为True时，此参数无效。
	UpgradeTimeEnd *string `json:"UpgradeTimeEnd,omitnil,omitempty" name:"UpgradeTimeEnd"`
}

type UpgradeDBInstanceMajorVersionRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过[DescribeDBInstances](https://cloud.tencent.com/document/api/409/16773)接口获取
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// 目标内核版本号，可以通过API [DescribeDBVersions](https://cloud.tencent.com/document/product/409/89018)获取可以升级的目标内核版本号。
	TargetDBKernelVersion *string `json:"TargetDBKernelVersion,omitnil,omitempty" name:"TargetDBKernelVersion"`

	// 是否为校验模式，若UpgradeCheck为True，表示仅进行内核版本兼容性检查，不会进行实质性的升级操作，对原实例无影响。检查结果可以通过升级日志查看。
	UpgradeCheck *bool `json:"UpgradeCheck,omitnil,omitempty" name:"UpgradeCheck"`

	// 升级前备份选项。True，表示升级前需要创建全量备份，False，表示升级前不需要创建全量备份。当实例已有备份集可以恢复到升级前的状态时，可选择False，否则需要指定为True。UpgradeCheck为True时，此参数无效。
	BackupBeforeUpgrade *bool `json:"BackupBeforeUpgrade,omitnil,omitempty" name:"BackupBeforeUpgrade"`

	// 统计信息收集选项，对主例运行 ANALYZE 以在升级后更新系统统计信息。可选值包括，
	// 0：不需要收集统计信息；
	// 1：实例恢复写之前收集统计信息；
	// 3：实例恢复写之后收集统计信息。
	// UpgradeCheck为True时，此参数无效。
	StatisticsRefreshOption *int64 `json:"StatisticsRefreshOption,omitnil,omitempty" name:"StatisticsRefreshOption"`

	// 插件升级选项，pg_upgrade不会升级任何插件，需要在升级完成后在创建过插件的库上执行"ALTER EXTENSION UPDATE"。发起升级实例大版本时可以指定在实例恢复写前/后是否需要升级任务自动升级插件版本。可选值包括：
	// 0：不需要自动升级插件；
	// 1：恢复写之前升级插件；
	// 2：恢复写之后升级插件。
	// UpgradeCheck为True时，此参数无效。
	ExtensionUpgradeOption *int64 `json:"ExtensionUpgradeOption,omitnil,omitempty" name:"ExtensionUpgradeOption"`

	// 升级时间选项，升级过程中会有一段时间实例只读，并会有一次秒级闪断，发起升级时需要选择这段影响的时间窗。可选值包括：
	// 0：自动执行，不需要指定时间窗；
	// 1：指定本次升级任务的时间窗，通过参数UpgradeTimeBegin和UpgradeTimeEnd设置；
	// 2：在实例运维时间窗内执行。
	// UpgradeCheck为True时，此参数无效。
	UpgradeTimeOption *int64 `json:"UpgradeTimeOption,omitnil,omitempty" name:"UpgradeTimeOption"`

	// 升级时间窗开始时间，时间格式：HH:MM:SS，例如：01:00:00。当UpgradeTimeOption为1时，该参数有效。
	// UpgradeCheck为True时，此参数无效。
	UpgradeTimeBegin *string `json:"UpgradeTimeBegin,omitnil,omitempty" name:"UpgradeTimeBegin"`

	// 升级时间窗截止时间，时间格式：HH:MM:SS，例如：02:00:00。当UpgradeTimeOption为1时，该参数有效。
	// UpgradeCheck为True时，此参数无效。
	UpgradeTimeEnd *string `json:"UpgradeTimeEnd,omitnil,omitempty" name:"UpgradeTimeEnd"`
}

func (r *UpgradeDBInstanceMajorVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceMajorVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "TargetDBKernelVersion")
	delete(f, "UpgradeCheck")
	delete(f, "BackupBeforeUpgrade")
	delete(f, "StatisticsRefreshOption")
	delete(f, "ExtensionUpgradeOption")
	delete(f, "UpgradeTimeOption")
	delete(f, "UpgradeTimeBegin")
	delete(f, "UpgradeTimeEnd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceMajorVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceMajorVersionResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeDBInstanceMajorVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDBInstanceMajorVersionResponseParams `json:"Response"`
}

func (r *UpgradeDBInstanceMajorVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceMajorVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Version struct {
	// 数据库引擎，支持：
	// 1、postgresql（云数据库PostgreSQL）；
	// 2、mssql_compatible（MSSQL兼容-云数据库PostgreSQL）；
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// 数据库版本，例如：12.4
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// 数据库主要版本，例如：12
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// 数据库内核版本，例如：v12.4_r1.3
	DBKernelVersion *string `json:"DBKernelVersion,omitnil,omitempty" name:"DBKernelVersion"`

	// 数据库内核支持的特性列表。例如，
	// TDE：支持数据加密。
	SupportedFeatureNames []*string `json:"SupportedFeatureNames,omitnil,omitempty" name:"SupportedFeatureNames"`

	// 数据库版本状态，包括：
	// AVAILABLE：可用；
	// UPGRADE_ONLY：不可创建，此版本仅可升级至高版本；
	// DEPRECATED：已弃用。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 该数据库版本（DBKernelVersion）可以升级到的版本号列表。其中包含可升级的小版本号和可升级的大版本号（完整内核版本格式示例：v15.1_v1.6）。
	AvailableUpgradeTarget []*string `json:"AvailableUpgradeTarget,omitnil,omitempty" name:"AvailableUpgradeTarget"`
}

type Xlog struct {
	// 备份文件唯一标识
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 文件生成的开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 文件生成的结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 内网下载地址
	InternalAddr *string `json:"InternalAddr,omitnil,omitempty" name:"InternalAddr"`

	// 外网下载地址
	ExternalAddr *string `json:"ExternalAddr,omitnil,omitempty" name:"ExternalAddr"`

	// 备份文件大小
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`
}

type ZoneInfo struct {
	// 该可用区的英文名称
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 该可用区的中文名称
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 该可用区对应的数字编号
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 可用状态包含，
	// UNAVAILABLE：不可用。
	// AVAILABLE：可用。
	// SELLOUT：售罄。
	// SUPPORTMODIFYONLY：支持变配。
	ZoneState *string `json:"ZoneState,omitnil,omitempty" name:"ZoneState"`

	// 该可用区是否支持Ipv6
	ZoneSupportIpv6 *uint64 `json:"ZoneSupportIpv6,omitnil,omitempty" name:"ZoneSupportIpv6"`

	// 该可用区对应的备可用区集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	StandbyZoneSet []*string `json:"StandbyZoneSet,omitnil,omitempty" name:"StandbyZoneSet"`
}