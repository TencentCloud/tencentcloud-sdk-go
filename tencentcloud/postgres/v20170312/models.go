// Copyright 1999-2018 Tencent Ltd.
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

type CreateDBInstancesRequest struct {
	*tchttp.BaseRequest
	// 项目ID。
	ProjectId *int64 `json:"ProjectId" name:"ProjectId"`
	// 售卖规格ID。该参数可以通过调用DescribeProductConfig的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode" name:"SpecCode"`
	// PostgreSQL内核版本，目前只支持：9.3.5、9.5.4两种版本。
	DBVersion *string `json:"DBVersion" name:"DBVersion"`
	// 实例容量大小，单位：GB。
	Storage *uint64 `json:"Storage" name:"Storage"`
	// 一次性购买的实例数量。
	InstanceCount *uint64 `json:"InstanceCount" name:"InstanceCount"`
	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值。
	Period *uint64 `json:"Period" name:"Period"`
	// 实例计费类型。目前只支持：PREPAID（预付费，即包年包月）。
	InstanceChargeType *string `json:"InstanceChargeType" name:"InstanceChargeType"`
	// 是否自动使用代金券。1（是），0（否），默认不使用。
	AutoVoucher *uint64 `json:"AutoVoucher" name:"AutoVoucher"`
	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds" name:"VoucherIds" list`
	// 私有网络ID。
	VpcId *string `json:"VpcId" name:"VpcId"`
	// 私有网络子网ID。
	SubnetId *string `json:"SubnetId" name:"SubnetId"`
	// 可用区ID。该参数可以通过调用 DescribeZones 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone" name:"Zone"`
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
		DealNames []*string `json:"DealNames" name:"DealNames" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DBInstance struct {
	// 实例所属地域，如: ap-guangzhou，对应RegionSet的Region字段
	Region *string `json:"Region" name:"Region"`
	// 实例所属可用区， 如：ap-guangzhou-3，对应ZoneSet的Zone字段
	Zone *string `json:"Zone" name:"Zone"`
	// 项目ID
	ProjectId *uint64 `json:"ProjectId" name:"ProjectId"`
	// 私有网络ID
	VpcId *string `json:"VpcId" name:"VpcId"`
	// SubnetId
	SubnetId *string `json:"SubnetId" name:"SubnetId"`
	// 实例ID
	DBInstanceId *string `json:"DBInstanceId" name:"DBInstanceId"`
	// 实例名称
	DBInstanceName *string `json:"DBInstanceName" name:"DBInstanceName"`
	// 实例状态
	DBInstanceStatus *string `json:"DBInstanceStatus" name:"DBInstanceStatus"`
	// 实例分配的内存大小，单位：GB
	DBInstanceMemory *uint64 `json:"DBInstanceMemory" name:"DBInstanceMemory"`
	// 实例分配的存储空间大小，单位：GB
	DBInstanceStorage *uint64 `json:"DBInstanceStorage" name:"DBInstanceStorage"`
	// 实例分配的CPU数量，单位：个
	DBInstanceCpu *uint64 `json:"DBInstanceCpu" name:"DBInstanceCpu"`
	// 售卖规格ID
	DBInstanceClass *string `json:"DBInstanceClass" name:"DBInstanceClass"`
	// 实例类型，类型有：1、primary（主实例）；2、readonly（只读实例）；3、guard（灾备实例）；4、temp（临时实例）
	DBInstanceType *string `json:"DBInstanceType" name:"DBInstanceType"`
	// 实例版本，目前只支持standard（双机高可用版, 一主一从）
	DBInstanceVersion *string `json:"DBInstanceVersion" name:"DBInstanceVersion"`
	// 实例DB字符集
	DBCharset *string `json:"DBCharset" name:"DBCharset"`
	// PostgreSQL内核版本
	DBVersion *string `json:"DBVersion" name:"DBVersion"`
	// 实例创建时间
	CreateTime *string `json:"CreateTime" name:"CreateTime"`
	// 实例执行最后一次更新的时间
	UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
	// 实例到期时间
	ExpireTime *string `json:"ExpireTime" name:"ExpireTime"`
	// 实例隔离时间
	IsolatedTime *string `json:"IsolatedTime" name:"IsolatedTime"`
	// 计费模式，1、prepaid（包年包月,预付费）；2、postpaid（按量计费，后付费）
	PayType *string `json:"PayType" name:"PayType"`
	// 是否自动续费，1：自动续费，0：不自动续费
	AutoRenew *uint64 `json:"AutoRenew" name:"AutoRenew"`
	// 实例网络连接信息
	DBInstanceNetInfo []*DBInstanceNetInfo `json:"DBInstanceNetInfo" name:"DBInstanceNetInfo" list`
}

type DBInstanceNetInfo struct {
	// DNS域名
	Address *string `json:"Address" name:"Address"`
	// Ip
	Ip *string `json:"Ip" name:"Ip"`
	// 连接Port地址
	Port *uint64 `json:"Port" name:"Port"`
	// 网络类型，1、inner（内网地址）；2、public（外网地址）
	NetType *string `json:"NetType" name:"NetType"`
	// 网络连接状态
	Status *string `json:"Status" name:"Status"`
}

type DescribeDBInstanceAttributeRequest struct {
	*tchttp.BaseRequest
	// 实例ID。
	DBInstanceId *string `json:"DBInstanceId" name:"DBInstanceId"`
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
		DBInstance *DBInstance `json:"DBInstance" name:"DBInstance"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
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
	// 每页显示数量，默认返回10条。
	PageSize *uint64 `json:"PageSize" name:"PageSize"`
	// 分页序号，从1开始。
	PageNumber *uint64 `json:"PageNumber" name:"PageNumber"`
	// 过滤条件，目前支持：db-instance-id、db-instance-name两种。
	Filters []*Filter `json:"Filters" name:"Filters" list`
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
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 实例详细信息集合。
		DBInstanceSet []*DBInstance `json:"DBInstanceSet" name:"DBInstanceSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductConfigRequest struct {
	*tchttp.BaseRequest
	// 可用区名称
	Zone *string `json:"Zone" name:"Zone"`
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
		SpecInfoList []*SpecInfo `json:"SpecInfoList" name:"SpecInfoList" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
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
		TotalCount *int64 `json:"TotalCount" name:"TotalCount"`
		// 地域信息集合。
		RegionSet []*RegionInfo `json:"RegionSet" name:"RegionSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionsResponse) FromJsonString(s string) error {
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
		TotalCount *int64 `json:"TotalCount" name:"TotalCount"`
		// 可用区信息集合。
		ZoneSet []*ZoneInfo `json:"ZoneSet" name:"ZoneSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZonesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 过滤键的名称。
	Name *string `json:"Name" name:"Name"`
	// 一个或者多个过滤值。
	Values []*string `json:"Values" name:"Values" list`
}

type InitDBInstancesRequest struct {
	*tchttp.BaseRequest
	// 实例ID集合。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet" name:"DBInstanceIdSet" list`
	// 实例根账号用户名。
	AdminName *string `json:"AdminName" name:"AdminName"`
	// 实例根账号用户名对应的密码。
	AdminPassword *string `json:"AdminPassword" name:"AdminPassword"`
	// 实例字符集，目前只支持：UTF8、LATIN1。
	Charset *string `json:"Charset" name:"Charset"`
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
		DBInstanceIdSet []*string `json:"DBInstanceIdSet" name:"DBInstanceIdSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
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
	Zone *string `json:"Zone" name:"Zone"`
	// 规格ID。该参数可以通过调用DescribeProductConfig接口的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode" name:"SpecCode"`
	// 存储容量大小，单位：GB。
	Storage *uint64 `json:"Storage" name:"Storage"`
	// 实例数量。目前最大数量不超过100，如需一次性创建更多实例，请联系客服支持。
	InstanceCount *uint64 `json:"InstanceCount" name:"InstanceCount"`
	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值。
	Period *uint64 `json:"Period" name:"Period"`
	// 计费ID。该参数可以通过调用DescribeProductConfig接口的返回值中的Pid字段来获取。
	Pid *uint64 `json:"Pid" name:"Pid"`
	// 实例计费类型。目前只支持：PREPAID（预付费，即包年包月）。
	InstanceChargeType *string `json:"InstanceChargeType" name:"InstanceChargeType"`
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
		OriginalPrice *uint64 `json:"OriginalPrice" name:"OriginalPrice"`
		// 折后价格，单位：分
		Price *uint64 `json:"Price" name:"Price"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceCreateDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// 该地域对应的英文名称
	Region *string `json:"Region" name:"Region"`
	// 该地域对应的中文名称
	RegionName *string `json:"RegionName" name:"RegionName"`
	// 该地域对应的数字编号
	RegionId *uint64 `json:"RegionId" name:"RegionId"`
	// 可用状态，UNAVAILABLE表示不可用，AVAILABLE表示可用
	RegionState *string `json:"RegionState" name:"RegionState"`
}

type SpecInfo struct {
	// 地域英文编码，对应RegionSet的Region字段
	Region *string `json:"Region" name:"Region"`
	// 区域英文编码，对应ZoneSet的Zone字段
	Zone *string `json:"Zone" name:"Zone"`
	// 规格详细信息列表
	SpecItemInfoList []*SpecItemInfo `json:"SpecItemInfoList" name:"SpecItemInfoList" list`
}

type SpecItemInfo struct {
	// 规格ID
	SpecCode *string `json:"SpecCode" name:"SpecCode"`
	// PostgreSQL的内核版本编号
	Version *string `json:"Version" name:"Version"`
	// 内核编号对应的完整版本名称
	VersionName *string `json:"VersionName" name:"VersionName"`
	// CPU核数
	Cpu []*uint64 `json:"Cpu" name:"Cpu" list`
	// 内存大小，单位：MB
	Memory []*uint64 `json:"Memory" name:"Memory" list`
	// 该规格所支持最大存储容量，单位：GB
	MaxStorage *uint64 `json:"MaxStorage" name:"MaxStorage"`
	// 该规格所支持最小存储容量，单位：GB
	MinStorage *uint64 `json:"MinStorage" name:"MinStorage"`
	// 该规格的预估QPS
	Qps *uint64 `json:"Qps" name:"Qps"`
	// 该规格对应的计费ID
	Pid *uint64 `json:"Pid" name:"Pid"`
}

type ZoneInfo struct {
	// 该可用区的英文名称
	Zone *string `json:"Zone" name:"Zone"`
	// 该可用区的中文名称
	ZoneName *string `json:"ZoneName" name:"ZoneName"`
	// 该可用区对应的数字编号
	ZoneId *uint64 `json:"ZoneId" name:"ZoneId"`
	// 可用状态，UNAVAILABLE表示不可用，AVAILABLE表示可用
	ZoneState *string `json:"ZoneState" name:"ZoneState"`
}
