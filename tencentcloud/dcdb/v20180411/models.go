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

package cvm

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddShardConfig struct {
	// 新增分片的数量
	ShardCount *int64 `json:"ShardCount" name:"ShardCount"`
	// 分片内存大小，单位 GB
	ShardMemory *int64 `json:"ShardMemory" name:"ShardMemory"`
	// 分片存储大小，单位 GB
	ShardStorage *int64 `json:"ShardStorage" name:"ShardStorage"`
}

type CreateDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	// 分片节点可用区分布，最多可填两个可用区。当分片规格为一主两从时，其中两个节点在第一个可用区。
	Zones []*string `json:"Zones" name:"Zones" list`
	// 欲购买实例的数量，目前只支持购买1个实例
	Count *int64 `json:"Count" name:"Count"`
	// 欲购买的时长，单位：月。
	Period *int64 `json:"Period" name:"Period"`
	// 项目 ID，可以通过查看项目列表获取，不传则关联到默认项目
	ProjectId *int64 `json:"ProjectId" name:"ProjectId"`
	// 虚拟私有网络 ID，不传或传空表示创建为基础网络
	VpcId *string `json:"VpcId" name:"VpcId"`
	// 虚拟私有网络子网 ID，VpcId不为空时必填
	SubnetId *string `json:"SubnetId" name:"SubnetId"`
	// 分片内存大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardMemory *int64 `json:"ShardMemory" name:"ShardMemory"`
	// 分片存储空间大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardStorage *int64 `json:"ShardStorage" name:"ShardStorage"`
	// 单个分片节点个数，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardNodeCount *int64 `json:"ShardNodeCount" name:"ShardNodeCount"`
	// 实例分片个数，可选范围2-8，可以通过升级实例进行新增分片到最多64个分片。
	ShardCount *int64 `json:"ShardCount" name:"ShardCount"`
	// 数据库引擎版本，当前可选：10.0.10，10.1.9，5.7.17
	DbVersionId *string `json:"DbVersionId" name:"DbVersionId"`
	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `json:"AutoVoucher" name:"AutoVoucher"`
	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds" name:"VoucherIds" list`
}

func (r *CreateDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDCDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 长订单号。可以据此调用 DescribeOrders
	//  查询订单详细信息，或在支付失败时调用用户账号相关接口进行支付。
		DealName *string `json:"DealName" name:"DealName"`
		// 订单对应的实例 ID 列表，如果此处没有返回实例 ID，可以通过订单查询接口获取。还可通过实例查询接口查询实例是否创建完成。
		InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDCDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DCDBInstanceInfo struct {
	// 实例ID
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	// 实例名称
	InstanceName *string `json:"InstanceName" name:"InstanceName"`
	// APPID
	AppId *int64 `json:"AppId" name:"AppId"`
	// 项目ID
	ProjectId *int64 `json:"ProjectId" name:"ProjectId"`
	// 地域
	Region *string `json:"Region" name:"Region"`
	// 可用区
	Zone *string `json:"Zone" name:"Zone"`
	// VPC数字ID
	VpcId *int64 `json:"VpcId" name:"VpcId"`
	// Subnet数字ID
	SubnetId *int64 `json:"SubnetId" name:"SubnetId"`
	// 状态中文描述
	StatusDesc *string `json:"StatusDesc" name:"StatusDesc"`
	// 状态
	Status *int64 `json:"Status" name:"Status"`
	// 内网IP
	Vip *string `json:"Vip" name:"Vip"`
	// 内网端口
	Vport *int64 `json:"Vport" name:"Vport"`
	// 创建时间
	CreateTime *string `json:"CreateTime" name:"CreateTime"`
	// 自动续费标志
	AutoRenewFlag *int64 `json:"AutoRenewFlag" name:"AutoRenewFlag"`
	// 内存大小，单位 GB
	Memory *int64 `json:"Memory" name:"Memory"`
	// 存储大小，单位 GB
	Storage *int64 `json:"Storage" name:"Storage"`
	// 分片个数
	ShardCount *int64 `json:"ShardCount" name:"ShardCount"`
	// 到期时间
	PeriodEndTime *string `json:"PeriodEndTime" name:"PeriodEndTime"`
	// 隔离时间
	IsolatedTimestamp *string `json:"IsolatedTimestamp" name:"IsolatedTimestamp"`
	// UIN
	Uin *string `json:"Uin" name:"Uin"`
	// 分片详情
	ShardDetail []*ShardInfo `json:"ShardDetail" name:"ShardDetail" list`
}

type Deal struct {
	// 订单号
	DealName *string `json:"DealName" name:"DealName"`
	// 所属账号
	OwnerUin *string `json:"OwnerUin" name:"OwnerUin"`
	// 商品数量
	Count *int64 `json:"Count" name:"Count"`
	// 关联的流程 Id，可用于查询流程执行状态
	FlowId *int64 `json:"FlowId" name:"FlowId"`
	// 只有创建实例的订单会填充该字段，表示该订单创建的实例的 ID。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
}

type DescribeDBLogFilesRequest struct {
	*tchttp.BaseRequest
	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	// 分片 ID，形如：shard-7noic7tv
	ShardId *string `json:"ShardId" name:"ShardId"`
	// 请求日志类型，取值只能为1、2、3或者4。1-binlog，2-冷备，3-errlog，4-slowlog。
	Type *int64 `json:"Type" name:"Type"`
}

func (r *DescribeDBLogFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBLogFilesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBLogFilesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBLogFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBLogFilesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstancesRequest struct {
	*tchttp.BaseRequest
	// 按照一个或者多个实例 ID 查询。实例 ID 形如：dcdbt-2t4cf98d
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
	// 搜索的字段名，当前支持的值有：instancename、vip、all。传 instancename 表示按实例名进行搜索；传 vip 表示按内网IP进行搜索；传 all 将会按实例ID、实例名和内网IP进行搜索。
	SearchName *string `json:"SearchName" name:"SearchName"`
	// 搜索的关键字，支持模糊搜索。多个关键字使用换行符（'\n'）分割。
	SearchKey *string `json:"SearchKey" name:"SearchKey"`
	// 按项目 ID 查询
	ProjectIds []*int64 `json:"ProjectIds" name:"ProjectIds" list`
	// 是否根据 VPC 网络来搜索，0 为否，1 为是
	IsFilterVpc []*int64 `json:"IsFilterVpc" name:"IsFilterVpc" list`
	// 私有网络 ID， IsFilterVpc 为 1 时有效
	VpcId *string `json:"VpcId" name:"VpcId"`
	// 私有网络的子网 ID， IsFilterVpc 为 1 时有效
	SubnetId *string `json:"SubnetId" name:"SubnetId"`
	// 排序字段， projectId， createtime， instancename 三者之一
	OrderBy *string `json:"OrderBy" name:"OrderBy"`
	// 排序类型， desc 或者 asc
	OrderByType *string `json:"OrderByType" name:"OrderByType"`
	// 偏移量，默认为 0
	Offset *int64 `json:"Offset" name:"Offset"`
	// 返回数量，默认为 10，最大值为 100。
	Limit *int64 `json:"Limit" name:"Limit"`
}

func (r *DescribeDCDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合条件的实例数量
		TotalCount *int64 `json:"TotalCount" name:"TotalCount"`
		// 实例详细信息列表
		Instances []*DCDBInstanceInfo `json:"Instances" name:"Instances" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBPriceRequest struct {
	*tchttp.BaseRequest
	// 欲新购实例的可用区ID。
	Zone *string `json:"Zone" name:"Zone"`
	// 欲购买实例的数量，目前只支持购买1个实例
	Count *int64 `json:"Count" name:"Count"`
	// 欲购买的时长，单位：月。
	Period *int64 `json:"Period" name:"Period"`
	// 单个分片节点个数大小，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardNodeCount *int64 `json:"ShardNodeCount" name:"ShardNodeCount"`
	// 分片内存大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardMemory *int64 `json:"ShardMemory" name:"ShardMemory"`
	// 分片存储空间大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardStorage *int64 `json:"ShardStorage" name:"ShardStorage"`
	// 实例分片个数，可选范围2-8，可以通过升级实例进行新增分片到最多64个分片。
	ShardCount *int64 `json:"ShardCount" name:"ShardCount"`
}

func (r *DescribeDCDBPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBPriceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBPriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 原价，单位：分
		OriginalPrice *int64 `json:"OriginalPrice" name:"OriginalPrice"`
		// 实际价格，单位：分。受折扣等影响，可能和原价不同。
		Price *int64 `json:"Price" name:"Price"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBPriceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBRenewalPriceRequest struct {
	*tchttp.BaseRequest
	// 待续费的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	// 续费时长，单位：月。不传则默认为1个月。
	Period *int64 `json:"Period" name:"Period"`
}

func (r *DescribeDCDBRenewalPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBRenewalPriceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBRenewalPriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 原价，单位：分
		OriginalPrice *int64 `json:"OriginalPrice" name:"OriginalPrice"`
		// 实际价格，单位：分。受折扣等影响，可能和原价不同。
		Price *int64 `json:"Price" name:"Price"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBRenewalPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

func (r *DescribeDCDBSaleInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBSaleInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 可售卖地域信息列表
		RegionList []*RegionInfo `json:"RegionList" name:"RegionList" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBSaleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBSaleInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBUpgradePriceRequest struct {
	*tchttp.BaseRequest
	// 待升级的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	// 升级类型，取值范围: 
	// <li> ADD: 新增分片 </li> 
	//  <li> EXPAND: 升级实例中的已有分片 </li> 
	//  <li> SPLIT: 将已有分片中的数据切分到新增分片上</li>
	UpgradeType *string `json:"UpgradeType" name:"UpgradeType"`
	// 新增分片配置，当UpgradeType为ADD时生效。
	AddShardConfig *AddShardConfig `json:"AddShardConfig" name:"AddShardConfig"`
	// 扩容分片配置，当UpgradeType为EXPAND时生效。
	ExpandShardConfig *ExpandShardConfig `json:"ExpandShardConfig" name:"ExpandShardConfig"`
	// 切分分片配置，当UpgradeType为SPLIT时生效。
	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig" name:"SplitShardConfig"`
}

func (r *DescribeDCDBUpgradePriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBUpgradePriceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBUpgradePriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 原价，单位：分
		OriginalPrice *int64 `json:"OriginalPrice" name:"OriginalPrice"`
		// 实际价格，单位：分。受折扣等影响，可能和原价不同。
		Price *int64 `json:"Price" name:"Price"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBUpgradePriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBUpgradePriceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOrdersRequest struct {
	*tchttp.BaseRequest
	// 待查询的长订单号列表，创建实例、续费实例、扩容实例接口返回。
	DealNames []*string `json:"DealNames" name:"DealNames" list`
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
		// 返回的订单数量。
		TotalCount []*int64 `json:"TotalCount" name:"TotalCount" list`
		// 订单信息列表。
		Deals []*Deal `json:"Deals" name:"Deals" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrdersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOrdersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeShardSpecRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeShardSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeShardSpecRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeShardSpecResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 按机型分类的可售卖规格列表
		SpecConfig []*SpecConfig `json:"SpecConfig" name:"SpecConfig" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShardSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeShardSpecResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExpandShardConfig struct {
	// 分片ID数组
	ShardInstanceIds []*string `json:"ShardInstanceIds" name:"ShardInstanceIds" list`
	// 分片内存大小，单位 GB
	ShardMemory *int64 `json:"ShardMemory" name:"ShardMemory"`
	// 分片存储大小，单位 GB
	ShardStorage *int64 `json:"ShardStorage" name:"ShardStorage"`
}

type RegionInfo struct {
	// 地域英文ID
	Region *string `json:"Region" name:"Region"`
	// 地域数字ID
	RegionId *int64 `json:"RegionId" name:"RegionId"`
	// 地域中文名
	RegionName *string `json:"RegionName" name:"RegionName"`
	// 可用区列表
	ZoneList []*ZonesInfo `json:"ZoneList" name:"ZoneList" list`
	// 可选择的主可用区和从可用区
	AvailableChoice []*ShardZoneChooseInfo `json:"AvailableChoice" name:"AvailableChoice" list`
}

type RenewDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	// 待续费的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	// 续费时长，单位：月。
	Period *int64 `json:"Period" name:"Period"`
	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `json:"AutoVoucher" name:"AutoVoucher"`
	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds" name:"VoucherIds" list`
}

func (r *RenewDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewDCDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 长订单号。可以据此调用 DescribeOrders
	//  查询订单详细信息，或在支付失败时调用用户账号相关接口进行支付。
		DealName *string `json:"DealName" name:"DealName"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewDCDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShardInfo struct {
	// 分片ID
	ShardInstanceId *string `json:"ShardInstanceId" name:"ShardInstanceId"`
	// 分片Set ID
	ShardSerialId *string `json:"ShardSerialId" name:"ShardSerialId"`
	// 状态
	Status *int64 `json:"Status" name:"Status"`
	// 创建时间
	Createtime *string `json:"Createtime" name:"Createtime"`
	// 内存大小，单位 GB
	Memory *int64 `json:"Memory" name:"Memory"`
	// 存储大小，单位 GB
	Storage *int64 `json:"Storage" name:"Storage"`
	// 分片数字ID
	ShardId *int64 `json:"ShardId" name:"ShardId"`
}

type ShardZoneChooseInfo struct {
	// 主可用区
	MasterZone *ZonesInfo `json:"MasterZone" name:"MasterZone"`
	// 可选的从可用区
	SlaveZones []*ZonesInfo `json:"SlaveZones" name:"SlaveZones" list`
}

type SpecConfig struct {
	// 规格机型
	Machine *string `json:"Machine" name:"Machine"`
	// 规格列表
	SpecConfigInfos []*SpecConfigInfo `json:"SpecConfigInfos" name:"SpecConfigInfos" list`
}

type SpecConfigInfo struct {
	// 节点个数，2 表示一主一从，3 表示一主二从
	NodeCount *uint64 `json:"NodeCount" name:"NodeCount"`
	// 内存大小，单位 GB
	Memory *int64 `json:"Memory" name:"Memory"`
	// 数据盘规格最小值，单位 GB
	MinStorage *int64 `json:"MinStorage" name:"MinStorage"`
	// 数据盘规格最大值，单位 GB
	MaxStorage *int64 `json:"MaxStorage" name:"MaxStorage"`
	// 推荐的使用场景
	SuitInfo *string `json:"SuitInfo" name:"SuitInfo"`
	// 产品类型 Id
	Pid *int64 `json:"Pid" name:"Pid"`
	// 最大 Qps 值
	Qps *int64 `json:"Qps" name:"Qps"`
}

type SplitShardConfig struct {
	// 分片ID数组
	ShardInstanceIds []*string `json:"ShardInstanceIds" name:"ShardInstanceIds" list`
	// 数据切分比例
	SplitRate *int64 `json:"SplitRate" name:"SplitRate"`
	// 分片内存大小，单位 GB
	ShardMemory *int64 `json:"ShardMemory" name:"ShardMemory"`
	// 分片存储大小，单位 GB
	ShardStorage *int64 `json:"ShardStorage" name:"ShardStorage"`
}

type UpgradeDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	// 待升级的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	// 升级类型，取值范围: 
	// <li> ADD: 新增分片 </li> 
	//  <li> EXPAND: 升级实例中的已有分片 </li> 
	//  <li> SPLIT: 将已有分片中的数据切分到新增分片上</li>
	UpgradeType *string `json:"UpgradeType" name:"UpgradeType"`
	// 新增分片配置，当UpgradeType为ADD时生效。
	AddShardConfig *AddShardConfig `json:"AddShardConfig" name:"AddShardConfig"`
	// 扩容分片配置，当UpgradeType为EXPAND时生效。
	ExpandShardConfig *ExpandShardConfig `json:"ExpandShardConfig" name:"ExpandShardConfig"`
	// 切分分片配置，当UpgradeType为SPLIT时生效。
	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig" name:"SplitShardConfig"`
	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `json:"AutoVoucher" name:"AutoVoucher"`
	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds" name:"VoucherIds" list`
}

func (r *UpgradeDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDCDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 长订单号。可以据此调用 DescribeOrders
	//  查询订单详细信息，或在支付失败时调用用户账号相关接口进行支付。
		DealName *string `json:"DealName" name:"DealName"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDCDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ZonesInfo struct {
	// 可用区英文ID
	Zone *string `json:"Zone" name:"Zone"`
	// 可用区数字ID
	ZoneId *int64 `json:"ZoneId" name:"ZoneId"`
	// 可用区中文名
	ZoneName *string `json:"ZoneName" name:"ZoneName"`
}
