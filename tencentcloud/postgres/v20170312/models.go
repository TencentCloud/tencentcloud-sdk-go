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

package v20170312

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccountInfo struct {
	// 实例ID，形如postgres-lnp6j617
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 帐号
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 帐号备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 帐号状态。 1-创建中，2-正常，3-修改中，4-密码重置中，-1-删除中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 帐号创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 帐号最后一次更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type AddDBInstanceToReadOnlyGroupRequestParams struct {
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
}

type AddDBInstanceToReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
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
	// 流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 慢SQL执行的用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 抽象参数之后的慢SQL
	NormalQuery *string `json:"NormalQuery,omitempty" name:"NormalQuery"`

	// 慢SQL执行的客户端地址
	ClientAddr *string `json:"ClientAddr,omitempty" name:"ClientAddr"`

	// 在选定时间范围内慢SQL语句执行的次数
	CallNum *uint64 `json:"CallNum,omitempty" name:"CallNum"`

	// 在选定时间范围内，慢SQL语句执行的次数占所有慢SQL的比例（小数返回）
	CallPercent *float64 `json:"CallPercent,omitempty" name:"CallPercent"`

	// 在选定时间范围内，慢SQL执行的总时间
	CostTime *float64 `json:"CostTime,omitempty" name:"CostTime"`

	// 在选定时间范围内，慢SQL语句执行的总时间占所有慢SQL的比例（小数返回）
	CostPercent *float64 `json:"CostPercent,omitempty" name:"CostPercent"`

	// 在选定时间范围内，慢SQL语句执行的耗时最短的时间（单位：ms）
	MinCostTime *float64 `json:"MinCostTime,omitempty" name:"MinCostTime"`

	// 在选定时间范围内，慢SQL语句执行的耗时最长的时间（单位：ms）
	MaxCostTime *float64 `json:"MaxCostTime,omitempty" name:"MaxCostTime"`

	// 在选定时间范围内，慢SQL语句执行的耗时平均时间（单位：ms）
	AvgCostTime *float64 `json:"AvgCostTime,omitempty" name:"AvgCostTime"`

	// 在选定时间范围内，慢SQL第一条开始执行的时间戳
	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`

	// 在选定时间范围内，慢SQL最后一条开始执行的时间戳
	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
}

type BackupPlan struct {
	// 备份周期
	BackupPeriod *string `json:"BackupPeriod,omitempty" name:"BackupPeriod"`

	// 基础备份保留时长
	BaseBackupRetentionPeriod *uint64 `json:"BaseBackupRetentionPeriod,omitempty" name:"BaseBackupRetentionPeriod"`

	// 开始备份的最早时间
	MinBackupStartTime *string `json:"MinBackupStartTime,omitempty" name:"MinBackupStartTime"`

	// 开始备份的最晚时间
	MaxBackupStartTime *string `json:"MaxBackupStartTime,omitempty" name:"MaxBackupStartTime"`
}

// Predefined struct for user
type CloneDBInstanceRequestParams struct {
	// 克隆的源实例ID。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 售卖规格ID。该参数可以通过调用DescribeProductConfig的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 实例容量大小，单位：GB。
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值，按量计费模式下该参数传1。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 续费标记：0-正常续费（默认）；1-自动续费。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 已配置的私有网络中的子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 新购实例的实例名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 实例计费类型。目前支持：PREPAID（预付费，即包年包月），POSTPAID_BY_HOUR（后付费，即按量计费）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 安全组ID。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例需要绑定的Tag信息，默认为空。
	TagList []*Tag `json:"TagList,omitempty" name:"TagList"`

	// 购买多可用区实例时填写。
	DBNodeSet []*DBNode `json:"DBNodeSet,omitempty" name:"DBNodeSet"`

	// 是否自动使用代金券。1（是），0（否），默认不使用。
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表。
	VoucherIds *string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 活动ID。
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 基础备份集ID。
	BackupSetId *string `json:"BackupSetId,omitempty" name:"BackupSetId"`

	// 恢复时间点。
	RecoveryTargetTime *string `json:"RecoveryTargetTime,omitempty" name:"RecoveryTargetTime"`
}

type CloneDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 克隆的源实例ID。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 售卖规格ID。该参数可以通过调用DescribeProductConfig的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 实例容量大小，单位：GB。
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值，按量计费模式下该参数传1。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 续费标记：0-正常续费（默认）；1-自动续费。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 已配置的私有网络中的子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 新购实例的实例名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 实例计费类型。目前支持：PREPAID（预付费，即包年包月），POSTPAID_BY_HOUR（后付费，即按量计费）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 安全组ID。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例需要绑定的Tag信息，默认为空。
	TagList []*Tag `json:"TagList,omitempty" name:"TagList"`

	// 购买多可用区实例时填写。
	DBNodeSet []*DBNode `json:"DBNodeSet,omitempty" name:"DBNodeSet"`

	// 是否自动使用代金券。1（是），0（否），默认不使用。
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表。
	VoucherIds *string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 活动ID。
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 基础备份集ID。
	BackupSetId *string `json:"BackupSetId,omitempty" name:"BackupSetId"`

	// 恢复时间点。
	RecoveryTargetTime *string `json:"RecoveryTargetTime,omitempty" name:"RecoveryTargetTime"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloneDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneDBInstanceResponseParams struct {
	// 订单号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 订单流水号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillId *string `json:"BillId,omitempty" name:"BillId"`

	// 克隆出的新实例ID，当前只支持后付费返回该值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CloseDBExtranetAccessRequestParams struct {
	// 实例ID，形如postgres-6r233v55
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 是否关闭Ipv6外网，1：是，0：否
	IsIpv6 *int64 `json:"IsIpv6,omitempty" name:"IsIpv6"`
}

type CloseDBExtranetAccessRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-6r233v55
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 是否关闭Ipv6外网，1：是，0：否
	IsIpv6 *int64 `json:"IsIpv6,omitempty" name:"IsIpv6"`
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
	// 异步任务流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CloseServerlessDBExtranetAccessRequestParams struct {
	// 实例唯一标识符
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例名称
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
}

type CloseServerlessDBExtranetAccessRequest struct {
	*tchttp.BaseRequest
	
	// 实例唯一标识符
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例名称
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
}

func (r *CloseServerlessDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseServerlessDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "DBInstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseServerlessDBExtranetAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseServerlessDBExtranetAccessResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CloseServerlessDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *CloseServerlessDBExtranetAccessResponseParams `json:"Response"`
}

func (r *CloseServerlessDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseServerlessDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceNetworkAccessRequestParams struct {
	// 实例ID，形如：postgres-6bwgamo3。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 私有网络统一 ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 是否指定分配vip true-指定分配  false-自动分配。
	IsAssignVip *bool `json:"IsAssignVip,omitempty" name:"IsAssignVip"`

	// 目标VIP地址。
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

type CreateDBInstanceNetworkAccessRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如：postgres-6bwgamo3。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 私有网络统一 ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 是否指定分配vip true-指定分配  false-自动分配。
	IsAssignVip *bool `json:"IsAssignVip,omitempty" name:"IsAssignVip"`

	// 目标VIP地址。
	Vip *string `json:"Vip,omitempty" name:"Vip"`
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
	// 流程ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateDBInstancesRequestParams struct {
	// 售卖规格ID。该参数可以通过调用DescribeProductConfig的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 实例容量大小，单位：GB。
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// 一次性购买的实例数量。取值1-100
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值，按量计费模式下该参数传1。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 可用区ID。该参数可以通过调用 DescribeZones 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// PostgreSQL版本。当输入该参数时，会基于此版本创建对应的最新内核版本号实例。该参数和DBMajorVersion、DBKernelVersion至少需要传递一个。
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 实例计费类型。目前支持：PREPAID（预付费，即包年包月），POSTPAID_BY_HOUR（后付费，即按量计费）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 是否自动使用代金券。1（是），0（否），默认不使用。
	AutoVoucher *uint64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 续费标记：0-正常续费（默认）；1-自动续费；
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 活动ID
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 实例名(后续支持)
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否需要支持Ipv6，1：是，0：否
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitempty" name:"NeedSupportIpv6"`

	// 实例需要绑定的Tag信息，默认为空
	TagList []*Tag `json:"TagList,omitempty" name:"TagList"`

	// 安全组id
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// PostgreSQL主要版本。当输入该参数时，会基于此版本创建对应的最新内核版本号实例。该参数和DBVersion、DBKernelVersion至少需要传递一个。
	DBMajorVersion *string `json:"DBMajorVersion,omitempty" name:"DBMajorVersion"`

	// PostgreSQL内核版本。当输入该参数时，会创建该内核版本号实例。该参数和DBVersion、DBMajorVersion至少需要传递一个。
	DBKernelVersion *string `json:"DBKernelVersion,omitempty" name:"DBKernelVersion"`
}

type CreateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 售卖规格ID。该参数可以通过调用DescribeProductConfig的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 实例容量大小，单位：GB。
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// 一次性购买的实例数量。取值1-100
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值，按量计费模式下该参数传1。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 可用区ID。该参数可以通过调用 DescribeZones 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// PostgreSQL版本。当输入该参数时，会基于此版本创建对应的最新内核版本号实例。该参数和DBMajorVersion、DBKernelVersion至少需要传递一个。
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 实例计费类型。目前支持：PREPAID（预付费，即包年包月），POSTPAID_BY_HOUR（后付费，即按量计费）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 是否自动使用代金券。1（是），0（否），默认不使用。
	AutoVoucher *uint64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 续费标记：0-正常续费（默认）；1-自动续费；
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 活动ID
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 实例名(后续支持)
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否需要支持Ipv6，1：是，0：否
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitempty" name:"NeedSupportIpv6"`

	// 实例需要绑定的Tag信息，默认为空
	TagList []*Tag `json:"TagList,omitempty" name:"TagList"`

	// 安全组id
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// PostgreSQL主要版本。当输入该参数时，会基于此版本创建对应的最新内核版本号实例。该参数和DBVersion、DBKernelVersion至少需要传递一个。
	DBMajorVersion *string `json:"DBMajorVersion,omitempty" name:"DBMajorVersion"`

	// PostgreSQL内核版本。当输入该参数时，会创建该内核版本号实例。该参数和DBVersion、DBMajorVersion至少需要传递一个。
	DBKernelVersion *string `json:"DBKernelVersion,omitempty" name:"DBKernelVersion"`
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
	delete(f, "SpecCode")
	delete(f, "Storage")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "Zone")
	delete(f, "ProjectId")
	delete(f, "DBVersion")
	delete(f, "InstanceChargeType")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "AutoRenewFlag")
	delete(f, "ActivityId")
	delete(f, "Name")
	delete(f, "NeedSupportIpv6")
	delete(f, "TagList")
	delete(f, "SecurityGroupIds")
	delete(f, "DBMajorVersion")
	delete(f, "DBKernelVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstancesResponseParams struct {
	// 订单号列表。每个实例对应一个订单号。
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 冻结流水号
	BillId *string `json:"BillId,omitempty" name:"BillId"`

	// 创建成功的实例ID集合，只在后付费情景下有返回值
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateInstancesRequestParams struct {
	// 售卖规格ID。该参数可以通过调用DescribeProductConfig的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 实例容量大小，单位：GB。
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// 一次性购买的实例数量。取值1-10。
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值，按量计费模式下该参数传1。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 可用区ID。该参数可以通过调用 DescribeZones 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例字符集，目前只支持：UTF8、LATIN1。
	Charset *string `json:"Charset,omitempty" name:"Charset"`

	// 实例根账号用户名。
	AdminName *string `json:"AdminName,omitempty" name:"AdminName"`

	// 实例根账号用户名对应的密码。
	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`

	// 项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// PostgreSQL版本。当输入该参数时，会基于此版本创建对应的最新内核版本号实例。该参数和DBMajorVersion、DBKernelVersion至少需要传递一个。
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 实例计费类型。目前支持：PREPAID（预付费，即包年包月），POSTPAID_BY_HOUR（后付费，即按量计费）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 是否自动使用代金券。1（是），0（否），默认不使用。
	AutoVoucher *uint64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 已配置的私有网络中的子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 续费标记：0-正常续费（默认）；1-自动续费。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 活动ID。
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 实例名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否需要支持Ipv6，1：是，0：否（默认）。
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitempty" name:"NeedSupportIpv6"`

	// 实例需要绑定的Tag信息，默认为空。
	TagList []*Tag `json:"TagList,omitempty" name:"TagList"`

	// 安全组ID。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// PostgreSQL主要版本。目前支持10，11，12，13这几个版本。当输入该参数时，会基于此版本创建对应的最新内核版本号实例。该参数和DBVersion、DBKernelVersion至少需要传递一个。
	DBMajorVersion *string `json:"DBMajorVersion,omitempty" name:"DBMajorVersion"`

	// PostgreSQL内核版本。当输入该参数时，会创建该内核版本号实例。该参数和DBVersion、DBMajorVersion至少需要传递一个。
	DBKernelVersion *string `json:"DBKernelVersion,omitempty" name:"DBKernelVersion"`

	// 实例节点信息，购买跨可用区实例时填写。
	DBNodeSet []*DBNode `json:"DBNodeSet,omitempty" name:"DBNodeSet"`

	// 是否需要支持数据透明加密，1：是，0：否（默认）。
	NeedSupportTDE *uint64 `json:"NeedSupportTDE,omitempty" name:"NeedSupportTDE"`

	// 自定义密钥的KeyId，若选择自定义密匙加密，则需要传入自定义密匙的KeyId，KeyId是CMK的唯一标识。
	KMSKeyId *string `json:"KMSKeyId,omitempty" name:"KMSKeyId"`

	// 使用KMS服务的地域，KMSRegion为空默认使用本地域的KMS，本地域不支持的情况下需自选其他KMS支持的地域。
	KMSRegion *string `json:"KMSRegion,omitempty" name:"KMSRegion"`

	// 数据库引擎，支持：
	// 1、postgresql（云数据库PostgreSQL）；
	// 2、mssql_compatible（MSSQL兼容-云数据库PostgreSQL）；
	// 如不指定默认使用postgresql。
	DBEngine *string `json:"DBEngine,omitempty" name:"DBEngine"`

	// 数据库引擎的配置信息，配置格式如下：
	// {"$key1":"$value1", "$key2":"$value2"}
	// 
	// 各引擎支持如下：
	// 1、mssql_compatible引擎：
	// migrationMode：数据库模式，可选参数，可取值：single-db（单数据库模式），multi-db（多数据库模式）。默认为single-db。
	// defaultLocale：排序区域规则，可选参数，在初始化后不可修改，默认为en_US，可选值如下：
	// "af_ZA", "sq_AL", "ar_DZ", "ar_BH", "ar_EG", "ar_IQ", "ar_JO", "ar_KW", "ar_LB", "ar_LY", "ar_MA", "ar_OM", "ar_QA", "ar_SA", "ar_SY", "ar_TN", "ar_AE", "ar_YE", "hy_AM", "az_Cyrl_AZ", "az_Latn_AZ", "eu_ES", "be_BY", "bg_BG", "ca_ES", "zh_HK", "zh_MO", "zh_CN", "zh_SG", "zh_TW", "hr_HR", "cs_CZ", "da_DK", "nl_BE", "nl_NL", "en_AU", "en_BZ", "en_CA", "en_IE", "en_JM", "en_NZ", "en_PH", "en_ZA", "en_TT", "en_GB", "en_US", "en_ZW", "et_EE", "fo_FO", "fa_IR", "fi_FI", "fr_BE", "fr_CA", "fr_FR", "fr_LU", "fr_MC", "fr_CH", "mk_MK", "ka_GE", "de_AT", "de_DE", "de_LI", "de_LU", "de_CH", "el_GR", "gu_IN", "he_IL", "hi_IN", "hu_HU", "is_IS", "id_ID", "it_IT", "it_CH", "ja_JP", "kn_IN", "kok_IN", "ko_KR", "ky_KG", "lv_LV", "lt_LT", "ms_BN", "ms_MY", "mr_IN", "mn_MN", "nb_NO", "nn_NO", "pl_PL", "pt_BR", "pt_PT", "pa_IN", "ro_RO", "ru_RU", "sa_IN", "sr_Cyrl_RS", "sr_Latn_RS", "sk_SK", "sl_SI", "es_AR", "es_BO", "es_CL", "es_CO", "es_CR", "es_DO", "es_EC", "es_SV", "es_GT", "es_HN", "es_MX", "es_NI", "es_PA", "es_PY","es_PE", "es_PR", "es_ES", "es_TRADITIONAL", "es_UY", "es_VE", "sw_KE", "sv_FI", "sv_SE", "tt_RU", "te_IN", "th_TH", "tr_TR", "uk_UA", "ur_IN", "ur_PK", "uz_Cyrl_UZ", "uz_Latn_UZ", "vi_VN"。
	// serverCollationName：排序规则名称，可选参数，在初始化后不可修改，默认为sql_latin1_general_cp1_ci_as，可选值如下：
	// "bbf_unicode_general_ci_as", "bbf_unicode_cp1_ci_as", "bbf_unicode_CP1250_ci_as", "bbf_unicode_CP1251_ci_as", "bbf_unicode_cp1253_ci_as", "bbf_unicode_cp1254_ci_as", "bbf_unicode_cp1255_ci_as", "bbf_unicode_cp1256_ci_as", "bbf_unicode_cp1257_ci_as", "bbf_unicode_cp1258_ci_as", "bbf_unicode_cp874_ci_as", "sql_latin1_general_cp1250_ci_as", "sql_latin1_general_cp1251_ci_as", "sql_latin1_general_cp1_ci_as", "sql_latin1_general_cp1253_ci_as", "sql_latin1_general_cp1254_ci_as", "sql_latin1_general_cp1255_ci_as","sql_latin1_general_cp1256_ci_as", "sql_latin1_general_cp1257_ci_as", "sql_latin1_general_cp1258_ci_as", "chinese_prc_ci_as", "cyrillic_general_ci_as", "finnish_swedish_ci_as", "french_ci_as", "japanese_ci_as", "korean_wansung_ci_as", "latin1_general_ci_as", "modern_spanish_ci_as", "polish_ci_as", "thai_ci_as", "traditional_spanish_ci_as", "turkish_ci_as", "ukrainian_ci_as", "vietnamese_ci_as"。
	DBEngineConfig *string `json:"DBEngineConfig,omitempty" name:"DBEngineConfig"`
}

type CreateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 售卖规格ID。该参数可以通过调用DescribeProductConfig的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 实例容量大小，单位：GB。
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// 一次性购买的实例数量。取值1-10。
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值，按量计费模式下该参数传1。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 可用区ID。该参数可以通过调用 DescribeZones 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例字符集，目前只支持：UTF8、LATIN1。
	Charset *string `json:"Charset,omitempty" name:"Charset"`

	// 实例根账号用户名。
	AdminName *string `json:"AdminName,omitempty" name:"AdminName"`

	// 实例根账号用户名对应的密码。
	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`

	// 项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// PostgreSQL版本。当输入该参数时，会基于此版本创建对应的最新内核版本号实例。该参数和DBMajorVersion、DBKernelVersion至少需要传递一个。
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 实例计费类型。目前支持：PREPAID（预付费，即包年包月），POSTPAID_BY_HOUR（后付费，即按量计费）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 是否自动使用代金券。1（是），0（否），默认不使用。
	AutoVoucher *uint64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 已配置的私有网络中的子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 续费标记：0-正常续费（默认）；1-自动续费。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 活动ID。
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 实例名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否需要支持Ipv6，1：是，0：否（默认）。
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitempty" name:"NeedSupportIpv6"`

	// 实例需要绑定的Tag信息，默认为空。
	TagList []*Tag `json:"TagList,omitempty" name:"TagList"`

	// 安全组ID。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// PostgreSQL主要版本。目前支持10，11，12，13这几个版本。当输入该参数时，会基于此版本创建对应的最新内核版本号实例。该参数和DBVersion、DBKernelVersion至少需要传递一个。
	DBMajorVersion *string `json:"DBMajorVersion,omitempty" name:"DBMajorVersion"`

	// PostgreSQL内核版本。当输入该参数时，会创建该内核版本号实例。该参数和DBVersion、DBMajorVersion至少需要传递一个。
	DBKernelVersion *string `json:"DBKernelVersion,omitempty" name:"DBKernelVersion"`

	// 实例节点信息，购买跨可用区实例时填写。
	DBNodeSet []*DBNode `json:"DBNodeSet,omitempty" name:"DBNodeSet"`

	// 是否需要支持数据透明加密，1：是，0：否（默认）。
	NeedSupportTDE *uint64 `json:"NeedSupportTDE,omitempty" name:"NeedSupportTDE"`

	// 自定义密钥的KeyId，若选择自定义密匙加密，则需要传入自定义密匙的KeyId，KeyId是CMK的唯一标识。
	KMSKeyId *string `json:"KMSKeyId,omitempty" name:"KMSKeyId"`

	// 使用KMS服务的地域，KMSRegion为空默认使用本地域的KMS，本地域不支持的情况下需自选其他KMS支持的地域。
	KMSRegion *string `json:"KMSRegion,omitempty" name:"KMSRegion"`

	// 数据库引擎，支持：
	// 1、postgresql（云数据库PostgreSQL）；
	// 2、mssql_compatible（MSSQL兼容-云数据库PostgreSQL）；
	// 如不指定默认使用postgresql。
	DBEngine *string `json:"DBEngine,omitempty" name:"DBEngine"`

	// 数据库引擎的配置信息，配置格式如下：
	// {"$key1":"$value1", "$key2":"$value2"}
	// 
	// 各引擎支持如下：
	// 1、mssql_compatible引擎：
	// migrationMode：数据库模式，可选参数，可取值：single-db（单数据库模式），multi-db（多数据库模式）。默认为single-db。
	// defaultLocale：排序区域规则，可选参数，在初始化后不可修改，默认为en_US，可选值如下：
	// "af_ZA", "sq_AL", "ar_DZ", "ar_BH", "ar_EG", "ar_IQ", "ar_JO", "ar_KW", "ar_LB", "ar_LY", "ar_MA", "ar_OM", "ar_QA", "ar_SA", "ar_SY", "ar_TN", "ar_AE", "ar_YE", "hy_AM", "az_Cyrl_AZ", "az_Latn_AZ", "eu_ES", "be_BY", "bg_BG", "ca_ES", "zh_HK", "zh_MO", "zh_CN", "zh_SG", "zh_TW", "hr_HR", "cs_CZ", "da_DK", "nl_BE", "nl_NL", "en_AU", "en_BZ", "en_CA", "en_IE", "en_JM", "en_NZ", "en_PH", "en_ZA", "en_TT", "en_GB", "en_US", "en_ZW", "et_EE", "fo_FO", "fa_IR", "fi_FI", "fr_BE", "fr_CA", "fr_FR", "fr_LU", "fr_MC", "fr_CH", "mk_MK", "ka_GE", "de_AT", "de_DE", "de_LI", "de_LU", "de_CH", "el_GR", "gu_IN", "he_IL", "hi_IN", "hu_HU", "is_IS", "id_ID", "it_IT", "it_CH", "ja_JP", "kn_IN", "kok_IN", "ko_KR", "ky_KG", "lv_LV", "lt_LT", "ms_BN", "ms_MY", "mr_IN", "mn_MN", "nb_NO", "nn_NO", "pl_PL", "pt_BR", "pt_PT", "pa_IN", "ro_RO", "ru_RU", "sa_IN", "sr_Cyrl_RS", "sr_Latn_RS", "sk_SK", "sl_SI", "es_AR", "es_BO", "es_CL", "es_CO", "es_CR", "es_DO", "es_EC", "es_SV", "es_GT", "es_HN", "es_MX", "es_NI", "es_PA", "es_PY","es_PE", "es_PR", "es_ES", "es_TRADITIONAL", "es_UY", "es_VE", "sw_KE", "sv_FI", "sv_SE", "tt_RU", "te_IN", "th_TH", "tr_TR", "uk_UA", "ur_IN", "ur_PK", "uz_Cyrl_UZ", "uz_Latn_UZ", "vi_VN"。
	// serverCollationName：排序规则名称，可选参数，在初始化后不可修改，默认为sql_latin1_general_cp1_ci_as，可选值如下：
	// "bbf_unicode_general_ci_as", "bbf_unicode_cp1_ci_as", "bbf_unicode_CP1250_ci_as", "bbf_unicode_CP1251_ci_as", "bbf_unicode_cp1253_ci_as", "bbf_unicode_cp1254_ci_as", "bbf_unicode_cp1255_ci_as", "bbf_unicode_cp1256_ci_as", "bbf_unicode_cp1257_ci_as", "bbf_unicode_cp1258_ci_as", "bbf_unicode_cp874_ci_as", "sql_latin1_general_cp1250_ci_as", "sql_latin1_general_cp1251_ci_as", "sql_latin1_general_cp1_ci_as", "sql_latin1_general_cp1253_ci_as", "sql_latin1_general_cp1254_ci_as", "sql_latin1_general_cp1255_ci_as","sql_latin1_general_cp1256_ci_as", "sql_latin1_general_cp1257_ci_as", "sql_latin1_general_cp1258_ci_as", "chinese_prc_ci_as", "cyrillic_general_ci_as", "finnish_swedish_ci_as", "french_ci_as", "japanese_ci_as", "korean_wansung_ci_as", "latin1_general_ci_as", "modern_spanish_ci_as", "polish_ci_as", "thai_ci_as", "traditional_spanish_ci_as", "turkish_ci_as", "ukrainian_ci_as", "vietnamese_ci_as"。
	DBEngineConfig *string `json:"DBEngineConfig,omitempty" name:"DBEngineConfig"`
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
	delete(f, "SpecCode")
	delete(f, "Storage")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "Zone")
	delete(f, "Charset")
	delete(f, "AdminName")
	delete(f, "AdminPassword")
	delete(f, "ProjectId")
	delete(f, "DBVersion")
	delete(f, "InstanceChargeType")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "AutoRenewFlag")
	delete(f, "ActivityId")
	delete(f, "Name")
	delete(f, "NeedSupportIpv6")
	delete(f, "TagList")
	delete(f, "SecurityGroupIds")
	delete(f, "DBMajorVersion")
	delete(f, "DBKernelVersion")
	delete(f, "DBNodeSet")
	delete(f, "NeedSupportTDE")
	delete(f, "KMSKeyId")
	delete(f, "KMSRegion")
	delete(f, "DBEngine")
	delete(f, "DBEngineConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesResponseParams struct {
	// 订单号列表。每个实例对应一个订单号。
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 冻结流水号。
	BillId *string `json:"BillId,omitempty" name:"BillId"`

	// 创建成功的实例ID集合，只在后付费情景下有返回值。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 数据库大版本号，例如：11，12，13
	DBMajorVersion *string `json:"DBMajorVersion,omitempty" name:"DBMajorVersion"`

	// 数据库引擎，例如：postgresql，mssql_compatible
	DBEngine *string `json:"DBEngine,omitempty" name:"DBEngine"`

	// 参数模板描述，长度为0～60个字符，仅支持数字,英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`
}

type CreateParameterTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称，长度为1～60个字符，仅支持数字,英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 数据库大版本号，例如：11，12，13
	DBMajorVersion *string `json:"DBMajorVersion,omitempty" name:"DBMajorVersion"`

	// 数据库引擎，例如：postgresql，mssql_compatible
	DBEngine *string `json:"DBEngine,omitempty" name:"DBEngine"`

	// 参数模板描述，长度为0～60个字符，仅支持数字,英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`
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
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 售卖规格ID。该参数可以通过调用DescribeProductConfig的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 实例容量大小，单位：GB。
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// 一次性购买的实例数量。取值1-100
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值，按量计费模式下该参数传1。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 只读实例的主实例ID
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitempty" name:"MasterDBInstanceId"`

	// 可用区ID。该参数可以通过调用 DescribeZones 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 项目ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 【废弃】不再需要指定，内核版本号与主实例保持一致
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 实例计费类型。目前支持：PREPAID（预付费，即包年包月），POSTPAID_BY_HOUR（后付费，即按量计费）。如果主实例为后付费，只读实例必须也为后付费。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 是否自动使用代金券。1（是），0（否），默认不使用。
	AutoVoucher *uint64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 续费标记：0-正常续费（默认）；1-自动续费；
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 优惠活动ID
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 实例名(后续支持)
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否需要支持Ipv6，1：是，0：否
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitempty" name:"NeedSupportIpv6"`

	// 只读组ID。
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 实例需要绑定的Tag信息，默认为空（该类型为Tag数组类型）
	TagList *Tag `json:"TagList,omitempty" name:"TagList"`

	// 安全组id
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

type CreateReadOnlyDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 售卖规格ID。该参数可以通过调用DescribeProductConfig的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 实例容量大小，单位：GB。
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// 一次性购买的实例数量。取值1-100
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值，按量计费模式下该参数传1。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 只读实例的主实例ID
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitempty" name:"MasterDBInstanceId"`

	// 可用区ID。该参数可以通过调用 DescribeZones 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 项目ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 【废弃】不再需要指定，内核版本号与主实例保持一致
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 实例计费类型。目前支持：PREPAID（预付费，即包年包月），POSTPAID_BY_HOUR（后付费，即按量计费）。如果主实例为后付费，只读实例必须也为后付费。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 是否自动使用代金券。1（是），0（否），默认不使用。
	AutoVoucher *uint64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 续费标记：0-正常续费（默认）；1-自动续费；
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 优惠活动ID
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 实例名(后续支持)
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否需要支持Ipv6，1：是，0：否
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitempty" name:"NeedSupportIpv6"`

	// 只读组ID。
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 实例需要绑定的Tag信息，默认为空（该类型为Tag数组类型）
	TagList *Tag `json:"TagList,omitempty" name:"TagList"`

	// 安全组id
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
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
	delete(f, "SpecCode")
	delete(f, "Storage")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "MasterDBInstanceId")
	delete(f, "Zone")
	delete(f, "ProjectId")
	delete(f, "DBVersion")
	delete(f, "InstanceChargeType")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "AutoRenewFlag")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ActivityId")
	delete(f, "Name")
	delete(f, "NeedSupportIpv6")
	delete(f, "ReadOnlyGroupId")
	delete(f, "TagList")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReadOnlyDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReadOnlyDBInstanceResponseParams struct {
	// 订单号列表。每个实例对应一个订单号
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 冻结流水号
	BillId *string `json:"BillId,omitempty" name:"BillId"`

	// 创建成功的实例ID集合，只在后付费情景下有返回值
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// RO组ID，形如：pgro-4t9c6g7k。
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 私有网络统一 ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 是否指定分配vip true-指定分配  false-自动分配。
	IsAssignVip *bool `json:"IsAssignVip,omitempty" name:"IsAssignVip"`

	// 目标VIP地址。
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

type CreateReadOnlyGroupNetworkAccessRequest struct {
	*tchttp.BaseRequest
	
	// RO组ID，形如：pgro-4t9c6g7k。
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 私有网络统一 ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 是否指定分配vip true-指定分配  false-自动分配。
	IsAssignVip *bool `json:"IsAssignVip,omitempty" name:"IsAssignVip"`

	// 目标VIP地址。
	Vip *string `json:"Vip,omitempty" name:"Vip"`
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
	// 流程ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 主实例ID
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitempty" name:"MasterDBInstanceId"`

	// 只读组名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 延迟时间大小开关：0关、1开
	ReplayLagEliminate *uint64 `json:"ReplayLagEliminate,omitempty" name:"ReplayLagEliminate"`

	// 延迟空间大小开关： 0关、1开
	ReplayLatencyEliminate *uint64 `json:"ReplayLatencyEliminate,omitempty" name:"ReplayLatencyEliminate"`

	// 延迟时间大小阈值，单位ms
	MaxReplayLag *uint64 `json:"MaxReplayLag,omitempty" name:"MaxReplayLag"`

	// 延迟空间大小阈值，单位MB
	MaxReplayLatency *uint64 `json:"MaxReplayLatency,omitempty" name:"MaxReplayLatency"`

	// 延迟剔除最小保留实例数
	MinDelayEliminateReserve *uint64 `json:"MinDelayEliminateReserve,omitempty" name:"MinDelayEliminateReserve"`

	// 安全组id
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

type CreateReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 主实例ID
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitempty" name:"MasterDBInstanceId"`

	// 只读组名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 延迟时间大小开关：0关、1开
	ReplayLagEliminate *uint64 `json:"ReplayLagEliminate,omitempty" name:"ReplayLagEliminate"`

	// 延迟空间大小开关： 0关、1开
	ReplayLatencyEliminate *uint64 `json:"ReplayLatencyEliminate,omitempty" name:"ReplayLatencyEliminate"`

	// 延迟时间大小阈值，单位ms
	MaxReplayLag *uint64 `json:"MaxReplayLag,omitempty" name:"MaxReplayLag"`

	// 延迟空间大小阈值，单位MB
	MaxReplayLatency *uint64 `json:"MaxReplayLatency,omitempty" name:"MaxReplayLatency"`

	// 延迟剔除最小保留实例数
	MinDelayEliminateReserve *uint64 `json:"MinDelayEliminateReserve,omitempty" name:"MinDelayEliminateReserve"`

	// 安全组id
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
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
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 流程ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CreateServerlessDBInstanceRequestParams struct {
	// 可用区ID。公测阶段仅支持ap-shanghai-2、ap-beijing-1,ap-guangzhou-2.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// DB实例名称，同一个账号下该值必须唯一。
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`

	// PostgreSQL内核版本，目前只支持：10.4。
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// PostgreSQL数据库字符集，目前支持UTF8。
	DBCharset *string `json:"DBCharset,omitempty" name:"DBCharset"`

	// 项目ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例需要绑定的标签数组信息
	TagList []*Tag `json:"TagList,omitempty" name:"TagList"`
}

type CreateServerlessDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 可用区ID。公测阶段仅支持ap-shanghai-2、ap-beijing-1,ap-guangzhou-2.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// DB实例名称，同一个账号下该值必须唯一。
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`

	// PostgreSQL内核版本，目前只支持：10.4。
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// PostgreSQL数据库字符集，目前支持UTF8。
	DBCharset *string `json:"DBCharset,omitempty" name:"DBCharset"`

	// 项目ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例需要绑定的标签数组信息
	TagList []*Tag `json:"TagList,omitempty" name:"TagList"`
}

func (r *CreateServerlessDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServerlessDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "DBInstanceName")
	delete(f, "DBVersion")
	delete(f, "DBCharset")
	delete(f, "ProjectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "TagList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServerlessDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServerlessDBInstanceResponseParams struct {
	// 实例ID，该ID全局唯一，如：postgres-xxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateServerlessDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateServerlessDBInstanceResponseParams `json:"Response"`
}

func (r *CreateServerlessDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServerlessDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBBackup struct {
	// 备份文件唯一标识
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 文件生成的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 文件生成的结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 文件大小(K)
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 策略（0-实例备份；1-多库备份）
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// 类型（0-定时）
	Way *int64 `json:"Way,omitempty" name:"Way"`

	// 备份方式（1-完整）
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 状态（1-创建中；2-成功；3-失败）
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// DB列表
	DbList []*string `json:"DbList,omitempty" name:"DbList"`

	// 内网下载地址
	InternalAddr *string `json:"InternalAddr,omitempty" name:"InternalAddr"`

	// 外网下载地址
	ExternalAddr *string `json:"ExternalAddr,omitempty" name:"ExternalAddr"`

	// 备份集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SetId *string `json:"SetId,omitempty" name:"SetId"`
}

type DBInstance struct {
	// 实例所属地域，如: ap-guangzhou，对应RegionSet的Region字段
	Region *string `json:"Region,omitempty" name:"Region"`

	// 实例所属可用区， 如：ap-guangzhou-3，对应ZoneSet的Zone字段
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例名称
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`

	// 实例状态，分别为：applying（申请中）、init(待初始化)、initing(初始化中)、running(运行中)、limited run（受限运行）、isolated（已隔离）、recycling（回收中）、recycled（已回收）、job running（任务执行中）、offline（下线）、migrating（迁移中）、expanding（扩容中）、waitSwitch（等待切换）、switching（切换中）、readonly（只读）、restarting（重启中）、network changing（网络变更中）、upgrading（内核版本升级中）
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" name:"DBInstanceStatus"`

	// 实例分配的内存大小，单位：GB
	DBInstanceMemory *uint64 `json:"DBInstanceMemory,omitempty" name:"DBInstanceMemory"`

	// 实例分配的存储空间大小，单位：GB
	DBInstanceStorage *uint64 `json:"DBInstanceStorage,omitempty" name:"DBInstanceStorage"`

	// 实例分配的CPU数量，单位：个
	DBInstanceCpu *uint64 `json:"DBInstanceCpu,omitempty" name:"DBInstanceCpu"`

	// 售卖规格ID
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" name:"DBInstanceClass"`

	// 实例类型，类型有：1、primary（主实例）；2、readonly（只读实例）；3、guard（灾备实例）；4、temp（临时实例）
	DBInstanceType *string `json:"DBInstanceType,omitempty" name:"DBInstanceType"`

	// 实例版本，目前只支持standard（双机高可用版, 一主一从）
	DBInstanceVersion *string `json:"DBInstanceVersion,omitempty" name:"DBInstanceVersion"`

	// 实例DB字符集
	DBCharset *string `json:"DBCharset,omitempty" name:"DBCharset"`

	// PostgreSQL版本
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 实例创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例执行最后一次更新的时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 实例到期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 实例隔离时间
	IsolatedTime *string `json:"IsolatedTime,omitempty" name:"IsolatedTime"`

	// 计费模式，1、prepaid（包年包月,预付费）；2、postpaid（按量计费，后付费）
	PayType *string `json:"PayType,omitempty" name:"PayType"`

	// 是否自动续费，1：自动续费，0：不自动续费
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 实例网络连接信息
	DBInstanceNetInfo []*DBInstanceNetInfo `json:"DBInstanceNetInfo,omitempty" name:"DBInstanceNetInfo"`

	// 机器类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 用户的AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 实例的Uid
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// 实例是否支持Ipv6，1：支持，0：不支持
	SupportIpv6 *uint64 `json:"SupportIpv6,omitempty" name:"SupportIpv6"`

	// 实例绑定的标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*Tag `json:"TagList,omitempty" name:"TagList"`

	// 主实例信息，仅在实例为只读实例时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitempty" name:"MasterDBInstanceId"`

	// 只读实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadOnlyInstanceNum *int64 `json:"ReadOnlyInstanceNum,omitempty" name:"ReadOnlyInstanceNum"`

	// 只读实例在只读组中的状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusInReadonlyGroup *string `json:"StatusInReadonlyGroup,omitempty" name:"StatusInReadonlyGroup"`

	// 下线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`

	// 数据库内核版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBKernelVersion *string `json:"DBKernelVersion,omitempty" name:"DBKernelVersion"`

	// 实例网络信息列表（此字段已废弃）
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkAccessList []*NetworkAccess `json:"NetworkAccessList,omitempty" name:"NetworkAccessList"`

	// PostgreSQL主要版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBMajorVersion *string `json:"DBMajorVersion,omitempty" name:"DBMajorVersion"`

	// 实例的节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBNodeSet []*DBNode `json:"DBNodeSet,omitempty" name:"DBNodeSet"`

	// 实例是否支持TDE数据加密  0：不支持，1：支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSupportTDE *int64 `json:"IsSupportTDE,omitempty" name:"IsSupportTDE"`

	// 数据库引擎，支持：
	// 1、postgresql（云数据库PostgreSQL）；
	// 2、mssql_compatible（MSSQL兼容-云数据库PostgreSQL）；
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBEngine *string `json:"DBEngine,omitempty" name:"DBEngine"`

	// 数据库引擎的配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBEngineConfig *string `json:"DBEngineConfig,omitempty" name:"DBEngineConfig"`
}

type DBInstanceNetInfo struct {
	// DNS域名
	Address *string `json:"Address,omitempty" name:"Address"`

	// IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 连接Port地址
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 网络类型，1、inner（基础网络内网地址）；2、private（私有网络内网地址）；3、public（基础网络或私有网络的外网地址）；
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 网络连接状态，1、initing（未开通）；2、opened（已开通）；3、closed（已关闭）；4、opening（开通中）；5、closing（关闭中）；
	Status *string `json:"Status,omitempty" name:"Status"`

	// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 连接数据库的协议类型，当前支持：postgresql、mssql（MSSQL兼容语法）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`
}

type DBNode struct {
	// 节点类型，值可以为：
	// Primary，代表主节点；
	// Standby，代表备节点。
	Role *string `json:"Role,omitempty" name:"Role"`

	// 节点所在可用区，例如 ap-guangzhou-1。
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

// Predefined struct for user
type DeleteDBInstanceNetworkAccessRequestParams struct {
	// 实例ID，形如：postgres-6bwgamo3。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 私有网络统一 ID，若是基础网络则传"0"。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID，若是基础网络则传"0"。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 目标VIP地址。
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

type DeleteDBInstanceNetworkAccessRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如：postgres-6bwgamo3。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 私有网络统一 ID，若是基础网络则传"0"。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID，若是基础网络则传"0"。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 目标VIP地址。
	Vip *string `json:"Vip,omitempty" name:"Vip"`
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
	// 流程ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeleteParameterTemplateRequestParams struct {
	// 参数模板ID，用于唯一确认待操作的参数模板
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DeleteParameterTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板ID，用于唯一确认待操作的参数模板
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// RO组ID，形如：pgro-4t9c6g7k。
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 私有网络统一 ID，若是基础网络则传"0"。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID，若是基础网络则传"0"。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 目标VIP地址。
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

type DeleteReadOnlyGroupNetworkAccessRequest struct {
	*tchttp.BaseRequest
	
	// RO组ID，形如：pgro-4t9c6g7k。
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 私有网络统一 ID，若是基础网络则传"0"。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID，若是基础网络则传"0"。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 目标VIP地址。
	Vip *string `json:"Vip,omitempty" name:"Vip"`
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
	// 流程ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
}

type DeleteReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 待删除只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeleteServerlessDBInstanceRequestParams struct {
	// DB实例名称，实例名和实例ID必须至少传一个，如果同时存在，将只以实例ID为准。
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`

	// DB实例ID，实例名和实例ID必须至少传一个，如果同时存在，将只以实例ID为准。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

type DeleteServerlessDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// DB实例名称，实例名和实例ID必须至少传一个，如果同时存在，将只以实例ID为准。
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`

	// DB实例ID，实例名和实例ID必须至少传一个，如果同时存在，将只以实例ID为准。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

func (r *DeleteServerlessDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServerlessDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceName")
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServerlessDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServerlessDBInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteServerlessDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServerlessDBInstanceResponseParams `json:"Response"`
}

func (r *DeleteServerlessDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServerlessDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsRequestParams struct {
	// 实例ID，形如postgres-6fego161
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 分页返回，每页最大返回数目，默认10，取值范围为1-100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数据按照创建时间或者用户名排序。取值只能为createTime或者name。createTime-按照创建时间排序；name-按照用户名排序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 返回结果是升序还是降序。取值只能为desc或者asc。desc-降序；asc-升序
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-6fego161
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 分页返回，每页最大返回数目，默认10，取值范围为1-100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数据按照创建时间或者用户名排序。取值只能为createTime或者name。createTime-按照创建时间排序；name-按照用户名排序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 返回结果是升序还是降序。取值只能为desc或者asc。desc-降序；asc-升序
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 帐号列表详细信息。
	Details []*AccountInfo `json:"Details,omitempty" name:"Details"`

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
type DescribeAvailableRecoveryTimeRequestParams struct {
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

type DescribeAvailableRecoveryTimeRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
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
	RecoveryBeginTime *string `json:"RecoveryBeginTime,omitempty" name:"RecoveryBeginTime"`

	// 可恢复的最晚时间，时区为东八区（UTC+8）。
	RecoveryEndTime *string `json:"RecoveryEndTime,omitempty" name:"RecoveryEndTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeBackupPlansRequestParams struct {
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

type DescribeBackupPlansRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
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
	Plans []*BackupPlan `json:"Plans,omitempty" name:"Plans"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeCloneDBInstanceSpecRequestParams struct {
	// 实例ID。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 基础备份集ID，此入参和RecoveryTargetTime必须选择一个传入。如与RecoveryTargetTime参数同时设置，则以此参数为准。
	BackupSetId *string `json:"BackupSetId,omitempty" name:"BackupSetId"`

	// 恢复目标时间，此入参和BackupSetId必须选择一个传入。时区以东八区（UTC+8）为准。
	RecoveryTargetTime *string `json:"RecoveryTargetTime,omitempty" name:"RecoveryTargetTime"`
}

type DescribeCloneDBInstanceSpecRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 基础备份集ID，此入参和RecoveryTargetTime必须选择一个传入。如与RecoveryTargetTime参数同时设置，则以此参数为准。
	BackupSetId *string `json:"BackupSetId,omitempty" name:"BackupSetId"`

	// 恢复目标时间，此入参和BackupSetId必须选择一个传入。时区以东八区（UTC+8）为准。
	RecoveryTargetTime *string `json:"RecoveryTargetTime,omitempty" name:"RecoveryTargetTime"`
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
	MinSpecCode *string `json:"MinSpecCode,omitempty" name:"MinSpecCode"`

	// 可购买的最小磁盘容量，单位GB。
	MinStorage *int64 `json:"MinStorage,omitempty" name:"MinStorage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 备份方式（1-全量）。目前只支持全量，取值为1。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 查询开始时间，形如2018-06-10 17:06:38，起始时间不得小于7天以前
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 备份列表分页返回，每页返回数量，默认为 20，最小为1，最大值为 100。（当该参数不传或者传0时按默认值处理）
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果中的第几页，从第0页开始。默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeDBBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-4wdeb0zv。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 备份方式（1-全量）。目前只支持全量，取值为1。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 查询开始时间，形如2018-06-10 17:06:38，起始时间不得小于7天以前
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 备份列表分页返回，每页返回数量，默认为 20，最小为1，最大值为 100。（当该参数不传或者传0时按默认值处理）
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果中的第几页，从第0页开始。默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 备份列表
	BackupList []*DBBackup `json:"BackupList,omitempty" name:"BackupList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例ID，形如postgres-5bq3wfjd
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 查询起始时间，形如2018-01-01 00:00:00，起始时间不得小于7天以前
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 数据库名字
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 搜索关键字
	SearchKeys []*string `json:"SearchKeys,omitempty" name:"SearchKeys"`

	// 分页返回，每页返回的最大数量。取值为1-100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回，返回第几页的数据，从第0页开始计数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeDBErrlogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-5bq3wfjd
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 查询起始时间，形如2018-01-01 00:00:00，起始时间不得小于7天以前
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 数据库名字
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 搜索关键字
	SearchKeys []*string `json:"SearchKeys,omitempty" name:"SearchKeys"`

	// 分页返回，每页返回的最大数量。取值为1-100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回，返回第几页的数据，从第0页开始计数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	// 本次调用返回了多少条数据
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 错误日志列表
	Details []*ErrLogDetail `json:"Details,omitempty" name:"Details"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

type DescribeDBInstanceAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
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
	DBInstance *DBInstance `json:"DBInstance,omitempty" name:"DBInstance"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeDBInstanceParametersRequestParams struct {
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 查询指定参数详情。ParamName为空或不传，默认返回全部参数列表
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`
}

type DescribeDBInstanceParametersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 查询指定参数详情。ParamName为空或不传，默认返回全部参数列表
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 参数列表返回详情
	Detail []*ParamInfo `json:"Detail,omitempty" name:"Detail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeDBInstanceSecurityGroupsRequestParams struct {
	// 实例ID，DBInstanceId和ReadOnlyGroupId至少传一个；如果都传，忽略ReadOnlyGroupId
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 只读组ID，DBInstanceId和ReadOnlyGroupId至少传一个；如果要查询只读组关联的安全组，只传ReadOnlyGroupId
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
}

type DescribeDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，DBInstanceId和ReadOnlyGroupId至少传一个；如果都传，忽略ReadOnlyGroupId
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 只读组ID，DBInstanceId和ReadOnlyGroupId至少传一个；如果要查询只读组关联的安全组，只传ReadOnlyGroupId
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
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
	SecurityGroupSet []*SecurityGroup `json:"SecurityGroupSet,omitempty" name:"SecurityGroupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// db-instance-name：按照实例名过滤，类型为string
	// db-project-id：按照项目ID过滤，类型为integer
	// db-pay-mode：按照付费模式过滤，类型为string
	// db-tag-key：按照标签键过滤，类型为string
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 每页显示数量，取值范围为1-100，默认为返回10条。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序指标，如实例名、创建时间等，支持DBInstanceId,CreateTime,Name,EndTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，包括升序：asc、降序：desc。
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个过滤条件进行查询，目前支持的过滤条件有：
	// db-instance-id：按照实例ID过滤，类型为string
	// db-instance-name：按照实例名过滤，类型为string
	// db-project-id：按照项目ID过滤，类型为integer
	// db-pay-mode：按照付费模式过滤，类型为string
	// db-tag-key：按照标签键过滤，类型为string
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 每页显示数量，取值范围为1-100，默认为返回10条。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序指标，如实例名、创建时间等，支持DBInstanceId,CreateTime,Name,EndTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，包括升序：asc、降序：desc。
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例详细信息集合。
	DBInstanceSet []*DBInstance `json:"DBInstanceSet,omitempty" name:"DBInstanceSet"`

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
type DescribeDBSlowlogsRequestParams struct {
	// 实例ID，形如postgres-lnp6j617
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 查询起始时间，形如2018-06-10 17:06:38，起始时间不得小于7天以前
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 数据库名字
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 按照何种指标排序，取值为sum_calls或者sum_cost_time。sum_calls-总调用次数；sum_cost_time-总的花费时间
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序规则。desc-降序；asc-升序
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 分页返回结果，每页最大返回数量，取值为1-100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回结果，返回结果的第几页，从0开始计数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeDBSlowlogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-lnp6j617
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 查询起始时间，形如2018-06-10 17:06:38，起始时间不得小于7天以前
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 数据库名字
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 按照何种指标排序，取值为sum_calls或者sum_cost_time。sum_calls-总调用次数；sum_cost_time-总的花费时间
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序规则。desc-降序；asc-升序
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 分页返回结果，每页最大返回数量，取值为1-100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回结果，返回结果的第几页，从0开始计数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDBSlowlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSlowlogsRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSlowlogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSlowlogsResponseParams struct {
	// 本次返回多少条数据
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 慢查询日志详情
	Detail *SlowlogDetail `json:"Detail,omitempty" name:"Detail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBSlowlogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSlowlogsResponseParams `json:"Response"`
}

func (r *DescribeDBSlowlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSlowlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBXlogsRequestParams struct {
	// 实例ID，形如postgres-4wdeb0zv。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 查询开始时间，形如2018-06-10 17:06:38，起始时间不得小于7天以前
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页返回，表示返回第几页的条目。从第0页开始计数。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回，表示每页有多少条目。取值为1-100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDBXlogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-4wdeb0zv。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 查询开始时间，形如2018-06-10 17:06:38，起始时间不得小于7天以前
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页返回，表示返回第几页的条目。从第0页开始计数。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回，表示每页有多少条目。取值为1-100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Xlog列表
	XlogList []*Xlog `json:"XlogList,omitempty" name:"XlogList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeDatabasesRequestParams struct {
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesResponseParams struct {
	// 数据库信息
	Items []*string `json:"Items,omitempty" name:"Items"`

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
type DescribeDefaultParametersRequestParams struct {
	// 数据库版本，大版本号，例如11，12，13
	DBMajorVersion *string `json:"DBMajorVersion,omitempty" name:"DBMajorVersion"`

	// 数据库引擎，例如：postgresql,mssql_compatible
	DBEngine *string `json:"DBEngine,omitempty" name:"DBEngine"`
}

type DescribeDefaultParametersRequest struct {
	*tchttp.BaseRequest
	
	// 数据库版本，大版本号，例如11，12，13
	DBMajorVersion *string `json:"DBMajorVersion,omitempty" name:"DBMajorVersion"`

	// 数据库引擎，例如：postgresql,mssql_compatible
	DBEngine *string `json:"DBEngine,omitempty" name:"DBEngine"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 参数信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamInfoSet []*ParamInfo `json:"ParamInfoSet,omitempty" name:"ParamInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例ID。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

type DescribeEncryptionKeysRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptionKeys []*EncryptionKey `json:"EncryptionKeys,omitempty" name:"EncryptionKeys"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeOrdersRequestParams struct {
	// 订单名集合
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
}

type DescribeOrdersRequest struct {
	*tchttp.BaseRequest
	
	// 订单名集合
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 订单数组
	Deals []*PgDeal `json:"Deals,omitempty" name:"Deals"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 参数模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DescribeParameterTemplateAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 参数模板包含的参数个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 参数模板包含的参数信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamInfoSet []*ParamInfo `json:"ParamInfoSet,omitempty" name:"ParamInfoSet"`

	// 参数模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 参数模板适用的数据库版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBMajorVersion *string `json:"DBMajorVersion,omitempty" name:"DBMajorVersion"`

	// 参数模板适用的数据库引擎
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBEngine *string `json:"DBEngine,omitempty" name:"DBEngine"`

	// 参数模板描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 过滤条件，目前支持的过滤条件有：TemplateName, TemplateId，DBMajorVersion，DBEngine
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 每页显示数量，[0，100]，默认 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序指标，枚举值，支持：CreateTime，TemplateName，DBMajorVersion
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，枚举值，支持：asc（升序） ，desc（降序）
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

type DescribeParameterTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，目前支持的过滤条件有：TemplateName, TemplateId，DBMajorVersion，DBEngine
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 每页显示数量，[0，100]，默认 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序指标，枚举值，支持：CreateTime，TemplateName，DBMajorVersion
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，枚举值，支持：asc（升序） ，desc（降序）
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 参数模板列表
	ParameterTemplateSet []*ParameterTemplate `json:"ParameterTemplateSet,omitempty" name:"ParameterTemplateSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例DB ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

type DescribeParamsEventRequest struct {
	*tchttp.BaseRequest
	
	// 实例DB ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例参数修改事件详情
	EventItems []*EventItem `json:"EventItems,omitempty" name:"EventItems"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 数据库引擎，支持：
	// 1、postgresql（云数据库PostgreSQL）；
	// 2、mssql_compatible（MSSQL兼容-云数据库PostgreSQL）；
	// 如不指定默认使用postgresql。
	DBEngine *string `json:"DBEngine,omitempty" name:"DBEngine"`
}

type DescribeProductConfigRequest struct {
	*tchttp.BaseRequest
	
	// 可用区名称
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 数据库引擎，支持：
	// 1、postgresql（云数据库PostgreSQL）；
	// 2、mssql_compatible（MSSQL兼容-云数据库PostgreSQL）；
	// 如不指定默认使用postgresql。
	DBEngine *string `json:"DBEngine,omitempty" name:"DBEngine"`
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
	SpecInfoList []*SpecInfo `json:"SpecInfoList,omitempty" name:"SpecInfoList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 过滤条件，必须传入主实例ID进行过滤，否则返回值将为空，过滤参数为：db-master-instance-id
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 查询每一页的条数，默认为10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 查询的页码，默认为1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 查询排序依据，目前支持:ROGroupId,CreateTime,Name
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 查询排序依据类型，目前支持:desc,asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

type DescribeReadOnlyGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，必须传入主实例ID进行过滤，否则返回值将为空，过滤参数为：db-master-instance-id
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 查询每一页的条数，默认为10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 查询的页码，默认为1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 查询排序依据，目前支持:ROGroupId,CreateTime,Name
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 查询排序依据类型，目前支持:desc,asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
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
	ReadOnlyGroupList []*ReadOnlyGroup `json:"ReadOnlyGroupList,omitempty" name:"ReadOnlyGroupList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 地域信息集合。
	RegionSet []*RegionInfo `json:"RegionSet,omitempty" name:"RegionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeServerlessDBInstancesRequestParams struct {
	// 查询条件
	Filter []*Filter `json:"Filter,omitempty" name:"Filter"`

	// 查询个数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序指标，目前支持实例创建时间CreateTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，包括升序、降序
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

type DescribeServerlessDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 查询条件
	Filter []*Filter `json:"Filter,omitempty" name:"Filter"`

	// 查询个数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序指标，目前支持实例创建时间CreateTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，包括升序、降序
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeServerlessDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServerlessDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessDBInstancesResponseParams struct {
	// 查询结果数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 查询结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBInstanceSet []*ServerlessDBInstance `json:"DBInstanceSet,omitempty" name:"DBInstanceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServerlessDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServerlessDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeServerlessDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryAnalysisRequestParams struct {
	// 实例ID。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 查询起始时间戳，格式 “YYYY-MM-DD HH:mm:ss” ，日志保留时间默认为7天，起始时间不能超出保留时间范围。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间戳，格式 “YYYY-MM-DD HH:mm:ss”。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 根据数据库名进行筛选，可以为空。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 排序维度。 可选参数，取值范围[CallNum,CostTime,AvgCostTime]。默认CallNum。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型。升序asc、降序desc。默认desc。
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 分页大小。取值范围[1,100]。默认50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移。取值范围[0,INF)。默认0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeSlowQueryAnalysisRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 查询起始时间戳，格式 “YYYY-MM-DD HH:mm:ss” ，日志保留时间默认为7天，起始时间不能超出保留时间范围。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间戳，格式 “YYYY-MM-DD HH:mm:ss”。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 根据数据库名进行筛选，可以为空。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 排序维度。 可选参数，取值范围[CallNum,CostTime,AvgCostTime]。默认CallNum。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型。升序asc、降序desc。默认desc。
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 分页大小。取值范围[1,100]。默认50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移。取值范围[0,INF)。默认0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	// 查询总条数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 慢SQL统计分析接口返回详情。
	Detail *Detail `json:"Detail,omitempty" name:"Detail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例ID。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 查询起始时间戳，格式 “YYYY-MM-DD HH:mm:ss” ，日志保留时间默认为7天，起始时间不能超出保留时间范围。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间戳，格式 “YYYY-MM-DD HH:mm:ss”。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 根据数据库名进行筛选，可以为空。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 排序类型。升序asc、降序desc。默认为desc。
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 排序维度。 可选参数，取值范围[SessionStartTime,Duration]，默认为SessionStartTime。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 分页大小。取值范围[1,100],默认为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移。取值范围[0,INF)，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeSlowQueryListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 查询起始时间戳，格式 “YYYY-MM-DD HH:mm:ss” ，日志保留时间默认为7天，起始时间不能超出保留时间范围。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间戳，格式 “YYYY-MM-DD HH:mm:ss”。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 根据数据库名进行筛选，可以为空。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 排序类型。升序asc、降序desc。默认为desc。
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 排序维度。 可选参数，取值范围[SessionStartTime,Duration]，默认为SessionStartTime。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 分页大小。取值范围[1,100],默认为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移。取值范围[0,INF)，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	// 选定时间范围内慢SQL总条数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 指定时间范围内，慢SQL耗时分段分析。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DurationAnalysis []*DurationAnalysis `json:"DurationAnalysis,omitempty" name:"DurationAnalysis"`

	// 指定时间范围内 慢SQL流水。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RawSlowQueryList []*RawSlowQuery `json:"RawSlowQueryList,omitempty" name:"RawSlowQueryList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 可用区信息集合。
	ZoneSet []*ZoneInfo `json:"ZoneSet,omitempty" name:"ZoneSet"`

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
type DestroyDBInstanceRequestParams struct {
	// 待下线实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

type DestroyDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 待下线实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TotalTime *float64 `json:"TotalTime,omitempty" name:"TotalTime"`

	// 输入时间范围内所有慢sql总条数
	TotalCallNum *uint64 `json:"TotalCallNum,omitempty" name:"TotalCallNum"`

	// 慢SQL统计分析列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnalysisItems []*AnalysisItems `json:"AnalysisItems,omitempty" name:"AnalysisItems"`
}

// Predefined struct for user
type DisIsolateDBInstancesRequestParams struct {
	// 资源ID列表。注意：当前已不支持同时解隔离多个实例，这里只能传入单个实例ID。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet"`

	// 包年包月实例解隔离时购买时常 以月为单位
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 是否使用代金券：true-使用,false-不使用，默认不使用
	AutoVoucher *bool `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券id列表
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`
}

type DisIsolateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID列表。注意：当前已不支持同时解隔离多个实例，这里只能传入单个实例ID。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet"`

	// 包年包月实例解隔离时购买时常 以月为单位
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 是否使用代金券：true-使用,false-不使用，默认不使用
	AutoVoucher *bool `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券id列表
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TimeSegment *string `json:"TimeSegment,omitempty" name:"TimeSegment"`

	// 对应时段区间慢SQL 条数
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type EncryptionKey struct {
	// KMS实例加密的KeyId。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// KMS实例加密Key的别名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyAlias *string `json:"KeyAlias,omitempty" name:"KeyAlias"`

	// 实例加密密钥DEK的密文。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DEKCipherTextBlob *string `json:"DEKCipherTextBlob,omitempty" name:"DEKCipherTextBlob"`

	// 密钥是否启用，1-启用， 0-未启用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`

	// KMS密钥所在地域。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyRegion *string `json:"KeyRegion,omitempty" name:"KeyRegion"`

	// DEK密钥创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ErrLogDetail struct {
	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 数据库名字
	Database *string `json:"Database,omitempty" name:"Database"`

	// 错误发生时间
	ErrTime *string `json:"ErrTime,omitempty" name:"ErrTime"`

	// 错误消息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type EventInfo struct {
	// 参数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 原参数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldValue *string `json:"OldValue,omitempty" name:"OldValue"`

	// 本次修改期望参数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewValue *string `json:"NewValue,omitempty" name:"NewValue"`

	// 后台参数修改开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 后台参数生效开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EffectiveTime *string `json:"EffectiveTime,omitempty" name:"EffectiveTime"`

	// 修改状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitempty" name:"State"`

	// 操作者（一般为用户sub UIN）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 时间日志。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventLog *string `json:"EventLog,omitempty" name:"EventLog"`
}

type EventItem struct {
	// 参数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 修改事件数
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`

	// 修改时间详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventDetail []*EventInfo `json:"EventDetail,omitempty" name:"EventDetail"`
}

type Filter struct {
	// 过滤键的名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

// Predefined struct for user
type InitDBInstancesRequestParams struct {
	// 实例ID集合。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet"`

	// 实例根账号用户名。
	AdminName *string `json:"AdminName,omitempty" name:"AdminName"`

	// 实例根账号用户名对应的密码。
	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`

	// 实例字符集，目前只支持：UTF8、LATIN1。
	Charset *string `json:"Charset,omitempty" name:"Charset"`
}

type InitDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID集合。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet"`

	// 实例根账号用户名。
	AdminName *string `json:"AdminName,omitempty" name:"AdminName"`

	// 实例根账号用户名对应的密码。
	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`

	// 实例字符集，目前只支持：UTF8、LATIN1。
	Charset *string `json:"Charset,omitempty" name:"Charset"`
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
	delete(f, "DBInstanceIdSet")
	delete(f, "AdminName")
	delete(f, "AdminPassword")
	delete(f, "Charset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InitDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InitDBInstancesResponseParams struct {
	// 实例ID集合。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet"`

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
type InquiryPriceCreateDBInstancesRequestParams struct {
	// 可用区ID。该参数可以通过调用 DescribeZones 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 规格ID。该参数可以通过调用DescribeProductConfig接口的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 存储容量大小，单位：GB。
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// 实例数量。目前最大数量不超过100，如需一次性创建更多实例，请联系客服支持。
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 【弃字段，不再生效】，计费ID。该参数可以通过调用DescribeProductConfig接口的返回值中的Pid字段来获取。
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 实例计费类型。目前只支持：PREPAID（预付费，即包年包月）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 实例类型，默认primary，支持如下：
	// primary（双机高可用（一主一从））
	// readonly（只读实例）
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// DB引擎，默认postgresql，支持如下：
	// postgresql（云数据库PostgreSQL）
	// mssql_compatible（MSSQL兼容-云数据库PostgreSQL）
	DBEngine *string `json:"DBEngine,omitempty" name:"DBEngine"`
}

type InquiryPriceCreateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 可用区ID。该参数可以通过调用 DescribeZones 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 规格ID。该参数可以通过调用DescribeProductConfig接口的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 存储容量大小，单位：GB。
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// 实例数量。目前最大数量不超过100，如需一次性创建更多实例，请联系客服支持。
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 【弃字段，不再生效】，计费ID。该参数可以通过调用DescribeProductConfig接口的返回值中的Pid字段来获取。
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 实例计费类型。目前只支持：PREPAID（预付费，即包年包月）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 实例类型，默认primary，支持如下：
	// primary（双机高可用（一主一从））
	// readonly（只读实例）
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// DB引擎，默认postgresql，支持如下：
	// postgresql（云数据库PostgreSQL）
	// mssql_compatible（MSSQL兼容-云数据库PostgreSQL）
	DBEngine *string `json:"DBEngine,omitempty" name:"DBEngine"`
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
	OriginalPrice *uint64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 折后实际付款金额，单位：分
	Price *uint64 `json:"Price,omitempty" name:"Price"`

	// 币种。例如，CNY：人民币。
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 续费周期，按月计算，最大不超过48
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

type InquiryPriceRenewDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 续费周期，按月计算，最大不超过48
	Period *int64 `json:"Period,omitempty" name:"Period"`
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
	OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 折后实际付款金额，单位为分。如24650表示246.5元
	Price *int64 `json:"Price,omitempty" name:"Price"`

	// 币种。例如，CNY：人民币。
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 实例的内存大小，单位GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例ID，形如postgres-hez4fh0v
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 【废弃参数，不再生效】，实例计费类型。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

type InquiryPriceUpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例的磁盘大小，单位GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 实例的内存大小，单位GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例ID，形如postgres-hez4fh0v
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 【废弃参数，不再生效】，实例计费类型。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceUpgradeDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpgradeDBInstanceResponseParams struct {
	// 刊例价费用
	OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 折后实际付款金额
	Price *int64 `json:"Price,omitempty" name:"Price"`

	// 币种。例如，CNY：人民币。
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例ID集合。注意：当前已不支持同时隔离多个实例，这里只能传入单个实例ID。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet"`
}

type IsolateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID集合。注意：当前已不支持同时隔离多个实例，这里只能传入单个实例ID。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyAccountRemarkRequestParams struct {
	// 实例ID，形如postgres-4wdeb0zv
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户UserName对应的新备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type ModifyAccountRemarkRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-4wdeb0zv
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户UserName对应的新备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyBackupPlanRequestParams struct {
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例最早开始备份时间
	MinBackupStartTime *string `json:"MinBackupStartTime,omitempty" name:"MinBackupStartTime"`

	// 实例最晚开始备份时间
	MaxBackupStartTime *string `json:"MaxBackupStartTime,omitempty" name:"MaxBackupStartTime"`

	// 实例备份保留时长，取值范围为3-7，单位是天
	BaseBackupRetentionPeriod *uint64 `json:"BaseBackupRetentionPeriod,omitempty" name:"BaseBackupRetentionPeriod"`

	// 实例备份周期，按照星期维度，格式为小写星期英文单词
	BackupPeriod []*string `json:"BackupPeriod,omitempty" name:"BackupPeriod"`
}

type ModifyBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例最早开始备份时间
	MinBackupStartTime *string `json:"MinBackupStartTime,omitempty" name:"MinBackupStartTime"`

	// 实例最晚开始备份时间
	MaxBackupStartTime *string `json:"MaxBackupStartTime,omitempty" name:"MaxBackupStartTime"`

	// 实例备份保留时长，取值范围为3-7，单位是天
	BaseBackupRetentionPeriod *uint64 `json:"BaseBackupRetentionPeriod,omitempty" name:"BaseBackupRetentionPeriod"`

	// 实例备份周期，按照星期维度，格式为小写星期英文单词
	BackupPeriod []*string `json:"BackupPeriod,omitempty" name:"BackupPeriod"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupPlanResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyDBInstanceDeploymentRequestParams struct {
	// 实例ID。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例节点信息。
	DBNodeSet []*DBNode `json:"DBNodeSet,omitempty" name:"DBNodeSet"`

	// 切换时间。默认为 立即切换，入参为 0 ：立即切换 。1：指定时间切换。2：维护时间窗口内切换
	SwitchTag *int64 `json:"SwitchTag,omitempty" name:"SwitchTag"`

	// 切换开始时间，时间格式：HH:MM:SS，例如：01:00:00。当SwitchTag为0或2时，该参数失效。
	SwitchStartTime *string `json:"SwitchStartTime,omitempty" name:"SwitchStartTime"`

	// 切换截止时间，时间格式：HH:MM:SS，例如：01:30:00。当SwitchTag为0或2时，该参数失效。
	SwitchEndTime *string `json:"SwitchEndTime,omitempty" name:"SwitchEndTime"`
}

type ModifyDBInstanceDeploymentRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例节点信息。
	DBNodeSet []*DBNode `json:"DBNodeSet,omitempty" name:"DBNodeSet"`

	// 切换时间。默认为 立即切换，入参为 0 ：立即切换 。1：指定时间切换。2：维护时间窗口内切换
	SwitchTag *int64 `json:"SwitchTag,omitempty" name:"SwitchTag"`

	// 切换开始时间，时间格式：HH:MM:SS，例如：01:00:00。当SwitchTag为0或2时，该参数失效。
	SwitchStartTime *string `json:"SwitchStartTime,omitempty" name:"SwitchStartTime"`

	// 切换截止时间，时间格式：HH:MM:SS，例如：01:30:00。当SwitchTag为0或2时，该参数失效。
	SwitchEndTime *string `json:"SwitchEndTime,omitempty" name:"SwitchEndTime"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyDBInstanceNameRequestParams struct {
	// 数据库实例ID，形如postgres-6fego161
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 新的数据库实例名字
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// 数据库实例ID，形如postgres-6fego161
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 新的数据库实例名字
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
	delete(f, "DBInstanceId")
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
type ModifyDBInstanceParametersRequestParams struct {
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 待修改参数及期望值
	ParamList []*ParamEntry `json:"ParamList,omitempty" name:"ParamList"`
}

type ModifyDBInstanceParametersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 待修改参数及期望值
	ParamList []*ParamEntry `json:"ParamList,omitempty" name:"ParamList"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 当前实例所在只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 实例修改的目标只读组ID
	NewReadOnlyGroupId *string `json:"NewReadOnlyGroupId,omitempty" name:"NewReadOnlyGroupId"`
}

type ModifyDBInstanceReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 当前实例所在只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 实例修改的目标只读组ID
	NewReadOnlyGroupId *string `json:"NewReadOnlyGroupId,omitempty" name:"NewReadOnlyGroupId"`
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
	// 流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// 实例或只读组要绑定的安全组列表
	SecurityGroupIdSet []*string `json:"SecurityGroupIdSet,omitempty" name:"SecurityGroupIdSet"`

	// 实例ID，DBInstanceId和ReadOnlyGroupId至少传一个；如果都传，忽略ReadOnlyGroupId
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 只读组ID，DBInstanceId和ReadOnlyGroupId至少传一个；如果要修改只读组关联的安全组，只传ReadOnlyGroupId
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例或只读组要绑定的安全组列表
	SecurityGroupIdSet []*string `json:"SecurityGroupIdSet,omitempty" name:"SecurityGroupIdSet"`

	// 实例ID，DBInstanceId和ReadOnlyGroupId至少传一个；如果都传，忽略ReadOnlyGroupId
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 只读组ID，DBInstanceId和ReadOnlyGroupId至少传一个；如果要修改只读组关联的安全组，只传ReadOnlyGroupId
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
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
type ModifyDBInstanceSpecRequestParams struct {
	// 实例ID，形如：postgres-6bwgamo3。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 修改后的实例内存大小，单位GiB。
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 修改后的实例磁盘大小，单位GiB。
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// 是否自动使用代金券,1是,0否，默认不使用。
	AutoVoucher *uint64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 活动ID。
	ActivityId *uint64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 指定实例配置完成变更后的切换时间，默认为 立即切换，入参为 0 ：立即切换 。1：指定时间切换。2：维护时间窗口内切换。
	SwitchTag *uint64 `json:"SwitchTag,omitempty" name:"SwitchTag"`

	// 切换开始时间，时间格式：HH:MM:SS，例如：01:00:00。当SwitchTag为0或2时，该参数失效。
	SwitchStartTime *string `json:"SwitchStartTime,omitempty" name:"SwitchStartTime"`

	// 切换截止时间，时间格式：HH:MM:SS，例如：01:30:00。当SwitchTag为0或2时，该参数失效。
	SwitchEndTime *string `json:"SwitchEndTime,omitempty" name:"SwitchEndTime"`
}

type ModifyDBInstanceSpecRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如：postgres-6bwgamo3。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 修改后的实例内存大小，单位GiB。
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 修改后的实例磁盘大小，单位GiB。
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// 是否自动使用代金券,1是,0否，默认不使用。
	AutoVoucher *uint64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 活动ID。
	ActivityId *uint64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 指定实例配置完成变更后的切换时间，默认为 立即切换，入参为 0 ：立即切换 。1：指定时间切换。2：维护时间窗口内切换。
	SwitchTag *uint64 `json:"SwitchTag,omitempty" name:"SwitchTag"`

	// 切换开始时间，时间格式：HH:MM:SS，例如：01:00:00。当SwitchTag为0或2时，该参数失效。
	SwitchStartTime *string `json:"SwitchStartTime,omitempty" name:"SwitchStartTime"`

	// 切换截止时间，时间格式：HH:MM:SS，例如：01:30:00。当SwitchTag为0或2时，该参数失效。
	SwitchEndTime *string `json:"SwitchEndTime,omitempty" name:"SwitchEndTime"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSpecResponseParams struct {
	// 订单号。
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 冻结流水号。
	BillId *string `json:"BillId,omitempty" name:"BillId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例ID集合。注意：当前已不支持同时操作多个实例，这里只能传入单个实例ID。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet"`

	// 所属新项目的ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type ModifyDBInstancesProjectRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID集合。注意：当前已不支持同时操作多个实例，这里只能传入单个实例ID。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet"`

	// 所属新项目的ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
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
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyParameterTemplateRequestParams struct {
	// 参数模板ID，用于唯一确认参数模板，不可修改
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 参数模板名称，长度为1～60个字符，仅支持数字,英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@  注：若该字段为空    ，则保持原参数模板名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 参数模板描述，长度为0～60个字符，仅支持数字,英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@  注：若不传入该参数，则保持原参数模板描述
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`

	// 需要修改或添加的参数集合，注：同一参数不能同时出现在修改添加集合和删除集合中
	ModifyParamEntrySet []*ParamEntry `json:"ModifyParamEntrySet,omitempty" name:"ModifyParamEntrySet"`

	// 需要从模板中删除的参数集合，注：同一参数不能同时出现在修改添加集合和删除集合中
	DeleteParamSet []*string `json:"DeleteParamSet,omitempty" name:"DeleteParamSet"`
}

type ModifyParameterTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板ID，用于唯一确认参数模板，不可修改
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 参数模板名称，长度为1～60个字符，仅支持数字,英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@  注：若该字段为空    ，则保持原参数模板名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 参数模板描述，长度为0～60个字符，仅支持数字,英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@  注：若不传入该参数，则保持原参数模板描述
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`

	// 需要修改或添加的参数集合，注：同一参数不能同时出现在修改添加集合和删除集合中
	ModifyParamEntrySet []*ParamEntry `json:"ModifyParamEntrySet,omitempty" name:"ModifyParamEntrySet"`

	// 需要从模板中删除的参数集合，注：同一参数不能同时出现在修改添加集合和删除集合中
	DeleteParamSet []*string `json:"DeleteParamSet,omitempty" name:"DeleteParamSet"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ModifyReadOnlyGroupConfigRequestParams struct {
	// 只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 只读组名称
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitempty" name:"ReadOnlyGroupName"`

	// 延迟时间配置开关：0关、1开
	ReplayLagEliminate *uint64 `json:"ReplayLagEliminate,omitempty" name:"ReplayLagEliminate"`

	// 延迟日志大小配置开关：0关、1开
	ReplayLatencyEliminate *uint64 `json:"ReplayLatencyEliminate,omitempty" name:"ReplayLatencyEliminate"`

	// 延迟日志大小阈值，单位MB
	MaxReplayLatency *uint64 `json:"MaxReplayLatency,omitempty" name:"MaxReplayLatency"`

	// 延迟时间大小阈值，单位ms
	MaxReplayLag *uint64 `json:"MaxReplayLag,omitempty" name:"MaxReplayLag"`

	// 自动负载均衡开关：0关、1开
	Rebalance *uint64 `json:"Rebalance,omitempty" name:"Rebalance"`

	// 延迟剔除最小保留实例数
	MinDelayEliminateReserve *uint64 `json:"MinDelayEliminateReserve,omitempty" name:"MinDelayEliminateReserve"`
}

type ModifyReadOnlyGroupConfigRequest struct {
	*tchttp.BaseRequest
	
	// 只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 只读组名称
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitempty" name:"ReadOnlyGroupName"`

	// 延迟时间配置开关：0关、1开
	ReplayLagEliminate *uint64 `json:"ReplayLagEliminate,omitempty" name:"ReplayLagEliminate"`

	// 延迟日志大小配置开关：0关、1开
	ReplayLatencyEliminate *uint64 `json:"ReplayLatencyEliminate,omitempty" name:"ReplayLatencyEliminate"`

	// 延迟日志大小阈值，单位MB
	MaxReplayLatency *uint64 `json:"MaxReplayLatency,omitempty" name:"MaxReplayLatency"`

	// 延迟时间大小阈值，单位ms
	MaxReplayLag *uint64 `json:"MaxReplayLag,omitempty" name:"MaxReplayLag"`

	// 自动负载均衡开关：0关、1开
	Rebalance *uint64 `json:"Rebalance,omitempty" name:"Rebalance"`

	// 延迟剔除最小保留实例数
	MinDelayEliminateReserve *uint64 `json:"MinDelayEliminateReserve,omitempty" name:"MinDelayEliminateReserve"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 处于等待切换状态中的实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 入参取值为 0 ，代表立即切换。
	SwitchTag *uint64 `json:"SwitchTag,omitempty" name:"SwitchTag"`
}

type ModifySwitchTimePeriodRequest struct {
	*tchttp.BaseRequest
	
	// 处于等待切换状态中的实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 入参取值为 0 ，代表立即切换。
	SwitchTag *uint64 `json:"SwitchTag,omitempty" name:"SwitchTag"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源类型，1-实例 2-RO组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *uint64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// IPV4地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// IPV6地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip6 *string `json:"Vip6,omitempty" name:"Vip6"`

	// 访问端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 网络状态，1-申请中，2-使用中，3-删除中，4-已删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcStatus *int64 `json:"VpcStatus,omitempty" name:"VpcStatus"`
}

type NormalQueryItem struct {
	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 调用次数
	Calls *int64 `json:"Calls,omitempty" name:"Calls"`

	// 粒度点
	CallsGrids []*int64 `json:"CallsGrids,omitempty" name:"CallsGrids"`

	// 花费总时间
	CostTime *float64 `json:"CostTime,omitempty" name:"CostTime"`

	// 影响的行数
	Rows *int64 `json:"Rows,omitempty" name:"Rows"`

	// 花费最小时间
	MinCostTime *float64 `json:"MinCostTime,omitempty" name:"MinCostTime"`

	// 花费最大时间
	MaxCostTime *float64 `json:"MaxCostTime,omitempty" name:"MaxCostTime"`

	// 最早一条慢SQL时间
	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`

	// 最晚一条慢SQL时间
	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`

	// 读共享内存块数
	SharedReadBlks *int64 `json:"SharedReadBlks,omitempty" name:"SharedReadBlks"`

	// 写共享内存块数
	SharedWriteBlks *int64 `json:"SharedWriteBlks,omitempty" name:"SharedWriteBlks"`

	// 读io总耗时
	ReadCostTime *int64 `json:"ReadCostTime,omitempty" name:"ReadCostTime"`

	// 写io总耗时
	WriteCostTime *int64 `json:"WriteCostTime,omitempty" name:"WriteCostTime"`

	// 数据库名字
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 脱敏后的慢SQL
	NormalQuery *string `json:"NormalQuery,omitempty" name:"NormalQuery"`
}

// Predefined struct for user
type OpenDBExtranetAccessRequestParams struct {
	// 实例ID，形如postgres-hez4fh0v
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 是否开通Ipv6外网，1：是，0：否
	IsIpv6 *int64 `json:"IsIpv6,omitempty" name:"IsIpv6"`
}

type OpenDBExtranetAccessRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-hez4fh0v
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 是否开通Ipv6外网，1：是，0：否
	IsIpv6 *int64 `json:"IsIpv6,omitempty" name:"IsIpv6"`
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
	// 异步任务流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type OpenServerlessDBExtranetAccessRequestParams struct {
	// 实例的唯一标识符
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例名称
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
}

type OpenServerlessDBExtranetAccessRequest struct {
	*tchttp.BaseRequest
	
	// 实例的唯一标识符
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例名称
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
}

func (r *OpenServerlessDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenServerlessDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "DBInstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenServerlessDBExtranetAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenServerlessDBExtranetAccessResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OpenServerlessDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *OpenServerlessDBExtranetAccessResponseParams `json:"Response"`
}

func (r *OpenServerlessDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenServerlessDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamEntry struct {
	// 参数名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 修改参数值。入参均以字符串形式传递，例如：小数”0.1“、整数”1000“、枚举”replica“
	ExpectedValue *string `json:"ExpectedValue,omitempty" name:"ExpectedValue"`
}

type ParamInfo struct {
	// 参数ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 参数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数值类型：integer（整型）、real（浮点型）、bool（布尔型）、enum（枚举类型）、mutil_enum（枚举类型、支持多选）。
	// 当参数类型为integer（整型）、real（浮点型）时，参数的取值范围根据返回值的Max、Min确定； 
	// 当参数类型为bool（布尔型）时，参数设置值取值范围是true | false； 
	// 当参数类型为enum（枚举类型）、mutil_enum（多枚举类型）时，参数的取值范围由返回值中的EnumValue确定。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamValueType *string `json:"ParamValueType,omitempty" name:"ParamValueType"`

	// 参数值 单位。参数没有单位时，该字段返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 参数默认值。以字符串形式返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// 参数当前运行值。以字符串形式返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 数值类型（integer、real）参数，取值下界
	// 注意：此字段可能返回 null，表示取不到有效值。
	Max *float64 `json:"Max,omitempty" name:"Max"`

	// 枚举类型参数，取值范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`

	// 数值类型（integer、real）参数，取值上界
	// 注意：此字段可能返回 null，表示取不到有效值。
	Min *float64 `json:"Min,omitempty" name:"Min"`

	// 参数中文描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamDescriptionCH *string `json:"ParamDescriptionCH,omitempty" name:"ParamDescriptionCH"`

	// 参数英文描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamDescriptionEN *string `json:"ParamDescriptionEN,omitempty" name:"ParamDescriptionEN"`

	// 参数修改，是否重启生效。（true为需要，false为不需要）
	// 注意：此字段可能返回 null，表示取不到有效值。
	NeedReboot *bool `json:"NeedReboot,omitempty" name:"NeedReboot"`

	// 参数中文分类
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassificationCN *string `json:"ClassificationCN,omitempty" name:"ClassificationCN"`

	// 参数英文分类
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassificationEN *string `json:"ClassificationEN,omitempty" name:"ClassificationEN"`

	// 是否和规格相关。（true为相关，false为不想关）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecRelated *bool `json:"SpecRelated,omitempty" name:"SpecRelated"`

	// 是否为重点参数。（true为重点参数，修改是需要重点关注，可能会影响实例性能）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Advanced *bool `json:"Advanced,omitempty" name:"Advanced"`

	// 参数最后一次修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyTime *string `json:"LastModifyTime,omitempty" name:"LastModifyTime"`

	// 参数存在主备制约，0：无主备制约关系，1:备机参数值需比主机大，2:主机参数值需比备机大
	// 注意：此字段可能返回 null，表示取不到有效值。
	StandbyRelated *int64 `json:"StandbyRelated,omitempty" name:"StandbyRelated"`

	// 参数版本关联信息，存储具体内核版本下的具体参数信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionRelationSet []*ParamVersionRelation `json:"VersionRelationSet,omitempty" name:"VersionRelationSet"`

	// 参数规格关联信息，存储具体规格下具体的参数信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecRelationSet []*ParamSpecRelation `json:"SpecRelationSet,omitempty" name:"SpecRelationSet"`
}

type ParamSpecRelation struct {
	// 参数名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数信息所属规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *string `json:"Memory,omitempty" name:"Memory"`

	// 参数在该规格下的默认值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 参数值单位。参数没有单位时，该字段返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 数值类型（integer、real）参数，取值上界
	// 注意：此字段可能返回 null，表示取不到有效值。
	Max *float64 `json:"Max,omitempty" name:"Max"`

	// 数值类型（integer、real）参数，取值下界
	// 注意：此字段可能返回 null，表示取不到有效值。
	Min *float64 `json:"Min,omitempty" name:"Min"`

	// 枚举类型参数，取值范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`
}

type ParamVersionRelation struct {
	// 参数名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数信息所属内核版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBKernelVersion *string `json:"DBKernelVersion,omitempty" name:"DBKernelVersion"`

	// 参数在该版本该规格下的默认值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 参数值单位。参数没有单位时，该字段返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 数值类型（integer、real）参数，取值上界
	// 注意：此字段可能返回 null，表示取不到有效值。
	Max *float64 `json:"Max,omitempty" name:"Max"`

	// 数值类型（integer、real）参数，取值下界
	// 注意：此字段可能返回 null，表示取不到有效值。
	Min *float64 `json:"Min,omitempty" name:"Min"`

	// 枚举类型参数，取值范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`
}

type ParameterTemplate struct {
	// 参数模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 参数模板名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 参数模板适用的数据库版本
	DBMajorVersion *string `json:"DBMajorVersion,omitempty" name:"DBMajorVersion"`

	// 参数模板适用的数据库引擎
	DBEngine *string `json:"DBEngine,omitempty" name:"DBEngine"`

	// 参数模板描述
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`
}

type PgDeal struct {
	// 订单名
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 所属用户
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 订单涉及多少个实例
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 付费模式。1-预付费；0-后付费
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 异步任务流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 实例ID数组
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet"`
}

type PolicyRule struct {
	// 策略，ACCEPT 或者 DROP
	Action *string `json:"Action,omitempty" name:"Action"`

	// 来源或目的 IP 或 IP 段，例如172.16.0.0/12
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// 端口
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// 网络协议，支持 UDP、TCP 等
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// 规则描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type RawSlowQuery struct {
	// 慢SQL 语句
	RawQuery *string `json:"RawQuery,omitempty" name:"RawQuery"`

	// 慢SQL 查询的数据库
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 慢SQL执行 耗时
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 执行慢SQL的客户端
	ClientAddr *string `json:"ClientAddr,omitempty" name:"ClientAddr"`

	// 执行慢SQL的用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 慢SQL执行的开始时间
	SessionStartTime *string `json:"SessionStartTime,omitempty" name:"SessionStartTime"`
}

type ReadOnlyGroup struct {
	// 只读组标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 只读组名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitempty" name:"ReadOnlyGroupName"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 主实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitempty" name:"MasterDBInstanceId"`

	// 最小保留实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinDelayEliminateReserve *int64 `json:"MinDelayEliminateReserve,omitempty" name:"MinDelayEliminateReserve"`

	// 延迟空间大小阈值
	MaxReplayLatency *int64 `json:"MaxReplayLatency,omitempty" name:"MaxReplayLatency"`

	// 延迟大小开关
	ReplayLatencyEliminate *int64 `json:"ReplayLatencyEliminate,omitempty" name:"ReplayLatencyEliminate"`

	// 延迟时间大小阈值
	MaxReplayLag *float64 `json:"MaxReplayLag,omitempty" name:"MaxReplayLag"`

	// 延迟时间开关
	ReplayLagEliminate *int64 `json:"ReplayLagEliminate,omitempty" name:"ReplayLagEliminate"`

	// 虚拟网络id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 地域id
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地区id
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例详细信息
	ReadOnlyDBInstanceList []*DBInstance `json:"ReadOnlyDBInstanceList,omitempty" name:"ReadOnlyDBInstanceList"`

	// 自动负载均衡开关
	Rebalance *int64 `json:"Rebalance,omitempty" name:"Rebalance"`

	// 网络信息
	DBInstanceNetInfo []*DBInstanceNetInfo `json:"DBInstanceNetInfo,omitempty" name:"DBInstanceNetInfo"`

	// 只读组网络信息列表（此字段已废弃）
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkAccessList []*NetworkAccess `json:"NetworkAccessList,omitempty" name:"NetworkAccessList"`
}

// Predefined struct for user
type RebalanceReadOnlyGroupRequestParams struct {
	// 只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
}

type RebalanceReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type RegionInfo struct {
	// 该地域对应的英文名称
	Region *string `json:"Region,omitempty" name:"Region"`

	// 该地域对应的中文名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 该地域对应的数字编号
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 可用状态，UNAVAILABLE表示不可用，AVAILABLE表示可用
	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`

	// 该地域是否支持国际站售卖，0：不支持，1：支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportInternational *uint64 `json:"SupportInternational,omitempty" name:"SupportInternational"`
}

// Predefined struct for user
type RemoveDBInstanceFromReadOnlyGroupRequestParams struct {
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
}

type RemoveDBInstanceFromReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
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
	// 流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例ID，形如postgres-6fego161
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 续费多少个月
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 是否自动使用代金券,1是,0否，默认不使用
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`
}

type RenewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-6fego161
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 续费多少个月
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 是否自动使用代金券,1是,0否，默认不使用
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`
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
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例ID，形如postgres-4wdeb0zv
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例账户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// UserName账户对应的新密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-4wdeb0zv
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例账户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// UserName账户对应的新密码
	Password *string `json:"Password,omitempty" name:"Password"`
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

// Predefined struct for user
type RestartDBInstanceRequestParams struct {
	// 实例ID，形如postgres-6r233v55
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

type RestartDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如postgres-6r233v55
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
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
	// 异步流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type SecurityGroup struct {
	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 创建时间
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
	SecurityGroupDescription *string `json:"SecurityGroupDescription,omitempty" name:"SecurityGroupDescription"`
}

type ServerlessDBAccount struct {
	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBUser *string `json:"DBUser,omitempty" name:"DBUser"`

	// 密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBPassword *string `json:"DBPassword,omitempty" name:"DBPassword"`

	// 连接数限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBConnLimit *int64 `json:"DBConnLimit,omitempty" name:"DBConnLimit"`
}

type ServerlessDBInstance struct {
	// 实例id，唯一标识符
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`

	// 实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" name:"DBInstanceStatus"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 私有网络Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 字符集
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBCharset *string `json:"DBCharset,omitempty" name:"DBCharset"`

	// 数据库版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例网络信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBInstanceNetInfo []*ServerlessDBInstanceNetInfo `json:"DBInstanceNetInfo,omitempty" name:"DBInstanceNetInfo"`

	// 实例账户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBAccountSet []*ServerlessDBAccount `json:"DBAccountSet,omitempty" name:"DBAccountSet"`

	// 实例下的db信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBDatabaseList []*string `json:"DBDatabaseList,omitempty" name:"DBDatabaseList"`

	// 实例绑定的标签数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*Tag `json:"TagList,omitempty" name:"TagList"`

	// 数据库内核版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBKernelVersion *string `json:"DBKernelVersion,omitempty" name:"DBKernelVersion"`

	// 数据库主要版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBMajorVersion *string `json:"DBMajorVersion,omitempty" name:"DBMajorVersion"`
}

type ServerlessDBInstanceNetInfo struct {
	// 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitempty" name:"Address"`

	// ip地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 端口号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 网络类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetType *string `json:"NetType,omitempty" name:"NetType"`
}

// Predefined struct for user
type SetAutoRenewFlagRequestParams struct {
	// 实例ID集合。注意：当前已不支持同时操作多个实例，这里只能传入单个实例ID。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet"`

	// 续费标记。0-正常续费；1-自动续费；2-到期不续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

type SetAutoRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID集合。注意：当前已不支持同时操作多个实例，这里只能传入单个实例ID。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet"`

	// 续费标记。0-正常续费；1-自动续费；2-到期不续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
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
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type SlowlogDetail struct {
	// 花费总时间
	TotalTime *float64 `json:"TotalTime,omitempty" name:"TotalTime"`

	// 调用总次数
	TotalCalls *int64 `json:"TotalCalls,omitempty" name:"TotalCalls"`

	// 脱敏后的慢SQL列表
	NormalQueries []*NormalQueryItem `json:"NormalQueries,omitempty" name:"NormalQueries"`
}

type SpecInfo struct {
	// 地域英文编码，对应RegionSet的Region字段
	Region *string `json:"Region,omitempty" name:"Region"`

	// 区域英文编码，对应ZoneSet的Zone字段
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 规格详细信息列表
	SpecItemInfoList []*SpecItemInfo `json:"SpecItemInfoList,omitempty" name:"SpecItemInfoList"`

	// 支持KMS的地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportKMSRegions []*string `json:"SupportKMSRegions,omitempty" name:"SupportKMSRegions"`
}

type SpecItemInfo struct {
	// 规格ID
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// PostgreSQL的版本编号
	Version *string `json:"Version,omitempty" name:"Version"`

	// 内核编号对应的完整版本名称
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// CPU核数
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存大小，单位：MB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 该规格所支持最大存储容量，单位：GB
	MaxStorage *uint64 `json:"MaxStorage,omitempty" name:"MaxStorage"`

	// 该规格所支持最小存储容量，单位：GB
	MinStorage *uint64 `json:"MinStorage,omitempty" name:"MinStorage"`

	// 该规格的预估QPS
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// 【该字段废弃】
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 机器类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// PostgreSQL的主要版本编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MajorVersion *string `json:"MajorVersion,omitempty" name:"MajorVersion"`

	// PostgreSQL的内核版本编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	KernelVersion *string `json:"KernelVersion,omitempty" name:"KernelVersion"`

	// 是否支持TDE数据加密功能，0-不支持，1-支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSupportTDE *int64 `json:"IsSupportTDE,omitempty" name:"IsSupportTDE"`
}

type Tag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

// Predefined struct for user
type UpgradeDBInstanceRequestParams struct {
	// 升级后的实例内存大小，单位GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 升级后的实例磁盘大小，单位GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 实例ID，形如postgres-lnp6j617
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 是否自动使用代金券,1是,0否，默认不使用
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 活动ID
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 指定实例配置完成变更后的切换时间，默认为 立即切换，入参为 0 ：立即切换 。1：指定时间切换。
	SwitchTag *int64 `json:"SwitchTag,omitempty" name:"SwitchTag"`

	// 切换开始时间
	SwitchStartTime *string `json:"SwitchStartTime,omitempty" name:"SwitchStartTime"`

	// 切换截止时间
	SwitchEndTime *string `json:"SwitchEndTime,omitempty" name:"SwitchEndTime"`
}

type UpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 升级后的实例内存大小，单位GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 升级后的实例磁盘大小，单位GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 实例ID，形如postgres-lnp6j617
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 是否自动使用代金券,1是,0否，默认不使用
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 活动ID
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 指定实例配置完成变更后的切换时间，默认为 立即切换，入参为 0 ：立即切换 。1：指定时间切换。
	SwitchTag *int64 `json:"SwitchTag,omitempty" name:"SwitchTag"`

	// 切换开始时间
	SwitchStartTime *string `json:"SwitchStartTime,omitempty" name:"SwitchStartTime"`

	// 切换截止时间
	SwitchEndTime *string `json:"SwitchEndTime,omitempty" name:"SwitchEndTime"`
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
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "DBInstanceId")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "ActivityId")
	delete(f, "SwitchTag")
	delete(f, "SwitchStartTime")
	delete(f, "SwitchEndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceResponseParams struct {
	// 交易名字。
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 冻结流水号
	BillId *string `json:"BillId,omitempty" name:"BillId"`

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

type Xlog struct {
	// 备份文件唯一标识
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 文件生成的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 文件生成的结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 内网下载地址
	InternalAddr *string `json:"InternalAddr,omitempty" name:"InternalAddr"`

	// 外网下载地址
	ExternalAddr *string `json:"ExternalAddr,omitempty" name:"ExternalAddr"`

	// 备份文件大小
	Size *int64 `json:"Size,omitempty" name:"Size"`
}

type ZoneInfo struct {
	// 该可用区的英文名称
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 该可用区的中文名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 该可用区对应的数字编号
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 可用状态包含，
	// UNAVAILABLE：不可用。
	// AVAILABLE：可用。
	// SELLOUT：售罄。
	// SUPPORTMODIFYONLY：支持变配。
	ZoneState *string `json:"ZoneState,omitempty" name:"ZoneState"`

	// 该可用区是否支持Ipv6
	ZoneSupportIpv6 *uint64 `json:"ZoneSupportIpv6,omitempty" name:"ZoneSupportIpv6"`

	// 该可用区对应的备可用区集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	StandbyZoneSet []*string `json:"StandbyZoneSet,omitempty" name:"StandbyZoneSet"`
}