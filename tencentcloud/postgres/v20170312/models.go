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

func (r *CloseDBExtranetAccessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloseDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务流程ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloseDBExtranetAccessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *CloseServerlessDBExtranetAccessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloseServerlessDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseServerlessDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloseServerlessDBExtranetAccessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 售卖规格ID。该参数可以通过调用DescribeProductConfig的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// PostgreSQL内核版本，目前支持：9.3.5、9.5.4、10.4三种版本。
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

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

	// 实例计费类型。目前支持：PREPAID（预付费，即包年包月），POSTPAID_BY_HOUR（后付费，即按量计费）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 是否自动使用代金券。1（是），0（否），默认不使用。
	AutoVoucher *uint64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`

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
}

func (r *CreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单号列表。每个实例对应一个订单号。
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`

		// 冻结流水号
		BillId *string `json:"BillId,omitempty" name:"BillId"`

		// 创建成功的实例ID集合，只在后付费情景下有返回值
		DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
}

func (r *CreateServerlessDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServerlessDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServerlessDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID，该ID全局唯一，如：postgres-xxxxx
		DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServerlessDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
	DbList []*string `json:"DbList,omitempty" name:"DbList" list`

	// 内网下载地址
	InternalAddr *string `json:"InternalAddr,omitempty" name:"InternalAddr"`

	// 外网下载地址
	ExternalAddr *string `json:"ExternalAddr,omitempty" name:"ExternalAddr"`
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

	// 实例状态，分别为：applying（申请中）、init(待初始化)、initing(初始化中)、running(运行中)、limited run（受限运行）、isolated（已隔离）、recycling（回收中）、recycled（已回收）、job running（任务执行中）、offline（下线）、migrating（迁移中）、expanding（扩容中）、readonly（只读）、restarting（重启中）
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

	// PostgreSQL内核版本
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
	DBInstanceNetInfo []*DBInstanceNetInfo `json:"DBInstanceNetInfo,omitempty" name:"DBInstanceNetInfo" list`

	// 机器类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 用户的AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 实例的Uid
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// 实例是否支持Ipv6，1：支持，0：不支持
	SupportIpv6 *uint64 `json:"SupportIpv6,omitempty" name:"SupportIpv6"`
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

	// 网络连接状态
	Status *string `json:"Status,omitempty" name:"Status"`
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

func (r *DeleteServerlessDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServerlessDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServerlessDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServerlessDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如postgres-6fego161
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 分页返回，每页最大返回数目，默认20，取值范围为1-100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回，返回第几页的用户数据。页码从0开始计数
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

func (r *DescribeAccountsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 本次调用接口共返回了多少条数据。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 帐号列表详细信息。
		Details []*AccountInfo `json:"Details,omitempty" name:"Details" list`

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

func (r *DescribeDBBackupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBBackupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回备份列表中备份文件的个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 备份列表
		BackupList []*DBBackup `json:"BackupList,omitempty" name:"BackupList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBBackupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	SearchKeys []*string `json:"SearchKeys,omitempty" name:"SearchKeys" list`

	// 分页返回，每页返回的最大数量。取值为1-100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回，返回第几页的数据，从第0页开始计数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDBErrlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBErrlogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBErrlogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 本次调用返回了多少条数据
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 错误日志列表
		Details []*ErrLogDetail `json:"Details,omitempty" name:"Details" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBErrlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBErrlogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeDBInstanceAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例详细信息。
		DBInstance *DBInstance `json:"DBInstance,omitempty" name:"DBInstance"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件，目前支持：db-instance-id、db-instance-name、db-project-id、db-pay-mode。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 每页显示数量，默认返回10条。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页序号，从0开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序指标，如实例名、创建时间等，支持DBInstanceId,CreateTime,Name,EndTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，包括升序、降序
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
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

		// 查询到的实例数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例详细信息集合。
		DBInstanceSet []*DBInstance `json:"DBInstanceSet,omitempty" name:"DBInstanceSet" list`

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

func (r *DescribeDBSlowlogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSlowlogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 本次返回多少条数据
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 慢查询日志详情
		Detail *SlowlogDetail `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSlowlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSlowlogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeDBXlogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBXlogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 表示此次返回结果有多少条数据。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Xlog列表
		XlogList []*Xlog `json:"XlogList,omitempty" name:"XlogList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBXlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBXlogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据库信息
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

type DescribeOrdersRequest struct {
	*tchttp.BaseRequest

	// 订单名集合
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`
}

func (r *DescribeOrdersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOrdersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOrdersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 订单数组
		Deals []*PgDeal `json:"Deals,omitempty" name:"Deals" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrdersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOrdersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductConfigRequest struct {
	*tchttp.BaseRequest

	// 可用区名称
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeProductConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 售卖规格列表。
		SpecInfoList []*SpecInfo `json:"SpecInfoList,omitempty" name:"SpecInfoList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的结果数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 地域信息集合。
		RegionSet []*RegionInfo `json:"RegionSet,omitempty" name:"RegionSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServerlessDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 查询条件
	Filter []*Filter `json:"Filter,omitempty" name:"Filter" list`

	// 查询个数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServerlessDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServerlessDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServerlessDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询结果数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 查询结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		DBInstanceSet []*ServerlessDBInstance `json:"DBInstanceSet,omitempty" name:"DBInstanceSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServerlessDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServerlessDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZonesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的结果数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 可用区信息集合。
		ZoneSet []*ZoneInfo `json:"ZoneSet,omitempty" name:"ZoneSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZonesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 待删除实例标识符
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

func (r *DestroyDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

type Filter struct {

	// 过滤键的名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type InitDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID集合。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

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

func (r *InitDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InitDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID集合。
		DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

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

	// 计费ID。该参数可以通过调用DescribeProductConfig接口的返回值中的Pid字段来获取。
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 实例计费类型。目前只支持：PREPAID（预付费，即包年包月）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

func (r *InquiryPriceCreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceCreateDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 原始价格，单位：分
		OriginalPrice *uint64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 折后价格，单位：分
		Price *uint64 `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceCreateDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *InquiryPriceRenewDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总费用，打折前的。比如24650表示246.5元
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 实际需要付款金额。比如24650表示246.5元
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例的磁盘大小，单位GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 实例的内存大小，单位GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例ID，形如postgres-hez4fh0v
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例计费类型，预付费或者后付费。PREPAID-预付费。目前只支持预付费。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

func (r *InquiryPriceUpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceUpgradeDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总费用，打折前的
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 实际需要付款金额
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceUpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceUpgradeDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *ModifyAccountRemarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountRemarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountRemarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

type ModifyDBInstancesProjectRequest struct {
	*tchttp.BaseRequest

	// postgresql实例ID数组
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

	// postgresql实例所属新项目的ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyDBInstancesProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstancesProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstancesProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 转移项目成功的实例个数
		Count *int64 `json:"Count,omitempty" name:"Count"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstancesProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstancesProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NormalQueryItem struct {

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 调用次数
	Calls *int64 `json:"Calls,omitempty" name:"Calls"`

	// 粒度点
	CallsGrids []*int64 `json:"CallsGrids,omitempty" name:"CallsGrids" list`

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

func (r *OpenDBExtranetAccessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务流程ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenDBExtranetAccessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *OpenServerlessDBExtranetAccessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenServerlessDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenServerlessDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenServerlessDBExtranetAccessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`
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
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`
}

func (r *RenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单名
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
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

func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *RestartDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RestartDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步流程ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestartDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

	// projectId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// VpcId
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
	DBInstanceNetInfo []*ServerlessDBInstanceNetInfo `json:"DBInstanceNetInfo,omitempty" name:"DBInstanceNetInfo" list`

	// 实例账户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBAccountSet []*ServerlessDBAccount `json:"DBAccountSet,omitempty" name:"DBAccountSet" list`

	// 实例下的db信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBDatabaseList []*string `json:"DBDatabaseList,omitempty" name:"DBDatabaseList" list`
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

type SetAutoRenewFlagRequest struct {
	*tchttp.BaseRequest

	// 实例ID数组
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

	// 续费标记。0-正常续费；1-自动续费；2-到期不续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

func (r *SetAutoRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetAutoRenewFlagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetAutoRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设置成功的实例个数
		Count *int64 `json:"Count,omitempty" name:"Count"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetAutoRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetAutoRenewFlagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SlowlogDetail struct {

	// 花费总时间
	TotalTime *float64 `json:"TotalTime,omitempty" name:"TotalTime"`

	// 调用总次数
	TotalCalls *int64 `json:"TotalCalls,omitempty" name:"TotalCalls"`

	// 脱敏后的慢SQL列表
	NormalQueries []*NormalQueryItem `json:"NormalQueries,omitempty" name:"NormalQueries" list`
}

type SpecInfo struct {

	// 地域英文编码，对应RegionSet的Region字段
	Region *string `json:"Region,omitempty" name:"Region"`

	// 区域英文编码，对应ZoneSet的Zone字段
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 规格详细信息列表
	SpecItemInfoList []*SpecItemInfo `json:"SpecItemInfoList,omitempty" name:"SpecItemInfoList" list`
}

type SpecItemInfo struct {

	// 规格ID
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// PostgreSQL的内核版本编号
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

	// 该规格对应的计费ID
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 机器类型
	Type *string `json:"Type,omitempty" name:"Type"`
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
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`

	// 活动ID
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`
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

		// 交易名字。
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// 冻结流水号
		BillId *string `json:"BillId,omitempty" name:"BillId"`

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

	// 可用状态，UNAVAILABLE表示不可用，AVAILABLE表示可用
	ZoneState *string `json:"ZoneState,omitempty" name:"ZoneState"`

	// 该可用区是否支持Ipv6
	ZoneSupportIpv6 *uint64 `json:"ZoneSupportIpv6,omitempty" name:"ZoneSupportIpv6"`
}
