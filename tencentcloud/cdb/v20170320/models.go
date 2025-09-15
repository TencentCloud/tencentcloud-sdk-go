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

package v20170320

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Account struct {
	// 账号名，可输入1 - 32个字符。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 账号的主机。
	// 说明：
	// 1. IP 形式，支持填入%。
	// 2. 多个主机以分隔符分隔，分隔符支持;,|换行符和空格。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type AccountInfo struct {
	// 账号备注信息
	Notes *string `json:"Notes,omitnil,omitempty" name:"Notes"`

	// 账号的域名
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 账号的名称
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 账号信息修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 修改密码的时间
	ModifyPasswordTime *string `json:"ModifyPasswordTime,omitnil,omitempty" name:"ModifyPasswordTime"`

	// 该值已废弃
	//
	// Deprecated: CreateTime is deprecated.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 用户最大可用实例连接数
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`

	// 用户账号是否开启了密码轮转
	OpenCam *bool `json:"OpenCam,omitnil,omitempty" name:"OpenCam"`
}

// Predefined struct for user
type AddTimeWindowRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 星期一的可维护时间段，其中每一个时间段的格式形如：10:00-12:00；起始时间按半个小时对齐；最短半个小时，最长三个小时；可设置多个时间段。 一周中应至少设置一天的时间窗。下同。
	Monday []*string `json:"Monday,omitnil,omitempty" name:"Monday"`

	// 星期二的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Tuesday []*string `json:"Tuesday,omitnil,omitempty" name:"Tuesday"`

	// 星期三的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Wednesday []*string `json:"Wednesday,omitnil,omitempty" name:"Wednesday"`

	// 星期四的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Thursday []*string `json:"Thursday,omitnil,omitempty" name:"Thursday"`

	// 星期五的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Friday []*string `json:"Friday,omitnil,omitempty" name:"Friday"`

	// 星期六的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Saturday []*string `json:"Saturday,omitnil,omitempty" name:"Saturday"`

	// 星期日的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Sunday []*string `json:"Sunday,omitnil,omitempty" name:"Sunday"`

	// 最大延迟阈值（秒），仅对主实例和灾备实例有效。默认值：10，取值范围：1-10的整数。
	MaxDelayTime *uint64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`
}

type AddTimeWindowRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 星期一的可维护时间段，其中每一个时间段的格式形如：10:00-12:00；起始时间按半个小时对齐；最短半个小时，最长三个小时；可设置多个时间段。 一周中应至少设置一天的时间窗。下同。
	Monday []*string `json:"Monday,omitnil,omitempty" name:"Monday"`

	// 星期二的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Tuesday []*string `json:"Tuesday,omitnil,omitempty" name:"Tuesday"`

	// 星期三的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Wednesday []*string `json:"Wednesday,omitnil,omitempty" name:"Wednesday"`

	// 星期四的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Thursday []*string `json:"Thursday,omitnil,omitempty" name:"Thursday"`

	// 星期五的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Friday []*string `json:"Friday,omitnil,omitempty" name:"Friday"`

	// 星期六的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Saturday []*string `json:"Saturday,omitnil,omitempty" name:"Saturday"`

	// 星期日的可维护时间窗口。 一周中应至少设置一天的时间窗。
	Sunday []*string `json:"Sunday,omitnil,omitempty" name:"Sunday"`

	// 最大延迟阈值（秒），仅对主实例和灾备实例有效。默认值：10，取值范围：1-10的整数。
	MaxDelayTime *uint64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type AddressInfo struct {
	// 地址的资源id标识。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 地址所在的vpc。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 地址所在的子网。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 地址的vip。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 地址的端口。
	VPort *int64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// 外网地址域名。
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// 外网地址端口。
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`
}

// Predefined struct for user
type AdjustCdbProxyAddressRequestParams struct {
	// 代理组 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 权重分配模式，
	// 系统自动分配："system"， 自定义："custom"
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// 是否开启延迟剔除，取值："true" | "false"
	IsKickOut *bool `json:"IsKickOut,omitnil,omitempty" name:"IsKickOut"`

	// 最小保留数量，最小取值：0。
	// 说明：当 IsKickOut 为 true 时才有效。
	MinCount *uint64 `json:"MinCount,omitnil,omitempty" name:"MinCount"`

	// 延迟剔除阈值，最小取值：1，取值范围：[1,10000]，整数。
	MaxDelay *uint64 `json:"MaxDelay,omitnil,omitempty" name:"MaxDelay"`

	// 是否开启故障转移，取值："true" | "false"
	FailOver *bool `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// 是否自动添加RO，取值："true" | "false"
	AutoAddRo *bool `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// 是否是只读，取值："true" | "false"
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// 代理组地址 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`

	// 是否开启事务分离，取值："true" | "false"，默认值 false。
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// 是否开启连接池。默认关闭。
	// 注意：如需使用数据库代理连接池能力，MySQL 8.0 主实例的内核小版本要大于等于 MySQL 8.0 20230630。
	ConnectionPool *bool `json:"ConnectionPool,omitnil,omitempty" name:"ConnectionPool"`

	// 读写权重分配。如果 WeightMode 传的是 system ，则传入的权重不生效，由系统分配默认权重。
	ProxyAllocation []*ProxyAllocation `json:"ProxyAllocation,omitnil,omitempty" name:"ProxyAllocation"`

	// 是否开启自适应负载均衡。默认关闭。
	AutoLoadBalance *bool `json:"AutoLoadBalance,omitnil,omitempty" name:"AutoLoadBalance"`

	// 访问模式：nearby - 就近访问，balance - 均衡分配，默认就近访问。
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// 是否将libra节点当作普通RO节点
	ApNodeAsRoNode *string `json:"ApNodeAsRoNode,omitnil,omitempty" name:"ApNodeAsRoNode"`

	// libra节点故障，是否转发给其他节点
	ApQueryToOtherNode *string `json:"ApQueryToOtherNode,omitnil,omitempty" name:"ApQueryToOtherNode"`
}

type AdjustCdbProxyAddressRequest struct {
	*tchttp.BaseRequest
	
	// 代理组 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 权重分配模式，
	// 系统自动分配："system"， 自定义："custom"
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// 是否开启延迟剔除，取值："true" | "false"
	IsKickOut *bool `json:"IsKickOut,omitnil,omitempty" name:"IsKickOut"`

	// 最小保留数量，最小取值：0。
	// 说明：当 IsKickOut 为 true 时才有效。
	MinCount *uint64 `json:"MinCount,omitnil,omitempty" name:"MinCount"`

	// 延迟剔除阈值，最小取值：1，取值范围：[1,10000]，整数。
	MaxDelay *uint64 `json:"MaxDelay,omitnil,omitempty" name:"MaxDelay"`

	// 是否开启故障转移，取值："true" | "false"
	FailOver *bool `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// 是否自动添加RO，取值："true" | "false"
	AutoAddRo *bool `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// 是否是只读，取值："true" | "false"
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// 代理组地址 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`

	// 是否开启事务分离，取值："true" | "false"，默认值 false。
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// 是否开启连接池。默认关闭。
	// 注意：如需使用数据库代理连接池能力，MySQL 8.0 主实例的内核小版本要大于等于 MySQL 8.0 20230630。
	ConnectionPool *bool `json:"ConnectionPool,omitnil,omitempty" name:"ConnectionPool"`

	// 读写权重分配。如果 WeightMode 传的是 system ，则传入的权重不生效，由系统分配默认权重。
	ProxyAllocation []*ProxyAllocation `json:"ProxyAllocation,omitnil,omitempty" name:"ProxyAllocation"`

	// 是否开启自适应负载均衡。默认关闭。
	AutoLoadBalance *bool `json:"AutoLoadBalance,omitnil,omitempty" name:"AutoLoadBalance"`

	// 访问模式：nearby - 就近访问，balance - 均衡分配，默认就近访问。
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// 是否将libra节点当作普通RO节点
	ApNodeAsRoNode *string `json:"ApNodeAsRoNode,omitnil,omitempty" name:"ApNodeAsRoNode"`

	// libra节点故障，是否转发给其他节点
	ApQueryToOtherNode *string `json:"ApQueryToOtherNode,omitnil,omitempty" name:"ApQueryToOtherNode"`
}

func (r *AdjustCdbProxyAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AdjustCdbProxyAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyGroupId")
	delete(f, "WeightMode")
	delete(f, "IsKickOut")
	delete(f, "MinCount")
	delete(f, "MaxDelay")
	delete(f, "FailOver")
	delete(f, "AutoAddRo")
	delete(f, "ReadOnly")
	delete(f, "ProxyAddressId")
	delete(f, "TransSplit")
	delete(f, "ConnectionPool")
	delete(f, "ProxyAllocation")
	delete(f, "AutoLoadBalance")
	delete(f, "AccessMode")
	delete(f, "ApNodeAsRoNode")
	delete(f, "ApQueryToOtherNode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AdjustCdbProxyAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AdjustCdbProxyAddressResponseParams struct {
	// 异步任务ID
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AdjustCdbProxyAddressResponse struct {
	*tchttp.BaseResponse
	Response *AdjustCdbProxyAddressResponseParams `json:"Response"`
}

func (r *AdjustCdbProxyAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AdjustCdbProxyAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AdjustCdbProxyRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 代理组 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 节点规格配置
	// 备注：数据库代理支持的节点规格为：2C4000MB、4C8000MB、8C16000MB。
	// 示例中参数说明：
	// NodeCount：节点个数
	// Region：节点地域
	// Zone：节点可用区
	// Cpu：单个代理节点核数（单位：核）
	// Mem：单个代理节点内存数（单位：MB）
	ProxyNodeCustom []*ProxyNodeCustom `json:"ProxyNodeCustom,omitnil,omitempty" name:"ProxyNodeCustom"`

	// 重新负载均衡：auto(自动),manual(手动)
	ReloadBalance *string `json:"ReloadBalance,omitnil,omitempty" name:"ReloadBalance"`

	// 升级切换时间：nowTime(升级完成时),timeWindow(维护时间内)
	UpgradeTime *string `json:"UpgradeTime,omitnil,omitempty" name:"UpgradeTime"`
}

type AdjustCdbProxyRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 代理组 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 节点规格配置
	// 备注：数据库代理支持的节点规格为：2C4000MB、4C8000MB、8C16000MB。
	// 示例中参数说明：
	// NodeCount：节点个数
	// Region：节点地域
	// Zone：节点可用区
	// Cpu：单个代理节点核数（单位：核）
	// Mem：单个代理节点内存数（单位：MB）
	ProxyNodeCustom []*ProxyNodeCustom `json:"ProxyNodeCustom,omitnil,omitempty" name:"ProxyNodeCustom"`

	// 重新负载均衡：auto(自动),manual(手动)
	ReloadBalance *string `json:"ReloadBalance,omitnil,omitempty" name:"ReloadBalance"`

	// 升级切换时间：nowTime(升级完成时),timeWindow(维护时间内)
	UpgradeTime *string `json:"UpgradeTime,omitnil,omitempty" name:"UpgradeTime"`
}

func (r *AdjustCdbProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AdjustCdbProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProxyGroupId")
	delete(f, "ProxyNodeCustom")
	delete(f, "ReloadBalance")
	delete(f, "UpgradeTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AdjustCdbProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AdjustCdbProxyResponseParams struct {
	// 异步任务ID
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AdjustCdbProxyResponse struct {
	*tchttp.BaseResponse
	Response *AdjustCdbProxyResponseParams `json:"Response"`
}

func (r *AdjustCdbProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AdjustCdbProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AggregationCondition struct {
	// 聚合字段。目前仅支持host-源IP、user-用户名、dbName-数据库名、sqlType-sql类型。
	AggregationField *string `json:"AggregationField,omitnil,omitempty" name:"AggregationField"`

	// 偏移量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 该聚合字段下要返回聚合桶的数量，最大100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type AnalysisNodeInfo struct {
	// 节点ID
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 数据加载状态
	DataStatus *string `json:"DataStatus,omitnil,omitempty" name:"DataStatus"`

	// cpu核数，单位：核
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存大小，单位: MB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 磁盘大小，单位：GB
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 节点所在可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 数据同步错误信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type AnalyzeAuditLogsRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要分析的日志开始时间，格式为："2023-02-16 00:00:20"。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 要分析的日志结束时间，格式为："2023-02-16 00:10:20"。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 聚合维度的排序条件。
	AggregationConditions []*AggregationCondition `json:"AggregationConditions,omitnil,omitempty" name:"AggregationConditions"`

	// 已废弃。
	//
	// Deprecated: AuditLogFilter is deprecated.
	AuditLogFilter *AuditLogFilter `json:"AuditLogFilter,omitnil,omitempty" name:"AuditLogFilter"`

	// 该过滤条件下的审计日志结果集作为分析日志。
	LogFilter []*InstanceAuditLogFilters `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`
}

type AnalyzeAuditLogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要分析的日志开始时间，格式为："2023-02-16 00:00:20"。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 要分析的日志结束时间，格式为："2023-02-16 00:10:20"。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 聚合维度的排序条件。
	AggregationConditions []*AggregationCondition `json:"AggregationConditions,omitnil,omitempty" name:"AggregationConditions"`

	// 已废弃。
	AuditLogFilter *AuditLogFilter `json:"AuditLogFilter,omitnil,omitempty" name:"AuditLogFilter"`

	// 该过滤条件下的审计日志结果集作为分析日志。
	LogFilter []*InstanceAuditLogFilters `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`
}

func (r *AnalyzeAuditLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AnalyzeAuditLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AggregationConditions")
	delete(f, "AuditLogFilter")
	delete(f, "LogFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AnalyzeAuditLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AnalyzeAuditLogsResponseParams struct {
	// 返回的聚合桶信息集
	Items []*AuditLogAggregationResult `json:"Items,omitnil,omitempty" name:"Items"`

	// 扫描的日志条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AnalyzeAuditLogsResponse struct {
	*tchttp.BaseResponse
	Response *AnalyzeAuditLogsResponseParams `json:"Response"`
}

func (r *AnalyzeAuditLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AnalyzeAuditLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsRequestParams struct {
	// 安全组 ID。可通过 [DescribeDBSecurityGroups](https://cloud.tencent.com/document/api/236/15854) 接口获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 实例 ID 列表，一个或者多个实例 ID 组成的数组。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 当传入只读实例ID时，默认操作的是对应只读组的安全组。如果需要操作只读实例ID的安全组， 需要将该入参置为True
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 安全组 ID。可通过 [DescribeDBSecurityGroups](https://cloud.tencent.com/document/api/236/15854) 接口获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 实例 ID 列表，一个或者多个实例 ID 组成的数组。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 当传入只读实例ID时，默认操作的是对应只读组的安全组。如果需要操作只读实例ID的安全组， 需要将该入参置为True
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`
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

type AuditFilter struct {
	// 过滤条件参数名称。目前支持：
	// SrcIp – 客户端 IP；
	// User – 数据库账户；
	// DB – 数据库名称；
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 过滤条件匹配类型。目前支持：
	// INC – 包含；
	// EXC – 不包含；
	// EQ – 等于；
	// NEQ – 不等于；
	Compare *string `json:"Compare,omitnil,omitempty" name:"Compare"`

	// 过滤条件匹配值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AuditInstanceFilters struct {
	// 过滤条件名。支持InstanceId-实例ID，InstanceName-实例名称，ProjectId-项目ID，TagKey-标签键，Tag-标签（以竖线分割，例：Tagkey|Tagvalue）。
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
	TagList []*TagInfoUnit `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 数据库内核类型
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 数据库内核版本
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
}

type AuditLog struct {
	// 影响行数。
	AffectRows *int64 `json:"AffectRows,omitnil,omitempty" name:"AffectRows"`

	// 错误码。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// SQL 类型。
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// 审计策略名称，逐步下线。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 数据库名称。
	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`

	// SQL 语句。
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

	// 线程ID。
	ThreadId *int64 `json:"ThreadId,omitnil,omitempty" name:"ThreadId"`

	// 扫描行数。
	CheckRows *int64 `json:"CheckRows,omitnil,omitempty" name:"CheckRows"`

	// cpu执行时间，微秒。
	CpuTime *float64 `json:"CpuTime,omitnil,omitempty" name:"CpuTime"`

	// IO等待时间，微秒。
	IoWaitTime *uint64 `json:"IoWaitTime,omitnil,omitempty" name:"IoWaitTime"`

	// 锁等待时间，微秒。
	LockWaitTime *uint64 `json:"LockWaitTime,omitnil,omitempty" name:"LockWaitTime"`

	// 开始时间，与timestamp构成一个精确到纳秒的时间。
	NsTime *uint64 `json:"NsTime,omitnil,omitempty" name:"NsTime"`

	// 事物持续时间，微秒。
	TrxLivingTime *uint64 `json:"TrxLivingTime,omitnil,omitempty" name:"TrxLivingTime"`

	// 日志命中规则模板的基本信息
	TemplateInfo []*LogRuleTemplateInfo `json:"TemplateInfo,omitnil,omitempty" name:"TemplateInfo"`

	//  事务ID
	TrxId *int64 `json:"TrxId,omitnil,omitempty" name:"TrxId"`
}

type AuditLogAggregationResult struct {
	// 聚合维度
	AggregationField *string `json:"AggregationField,omitnil,omitempty" name:"AggregationField"`

	// 聚合桶的结果集
	Buckets []*Bucket `json:"Buckets,omitnil,omitempty" name:"Buckets"`
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

	// 影响行数，格式为M-N，例如：10-200
	AffectRowsSection *string `json:"AffectRowsSection,omitnil,omitempty" name:"AffectRowsSection"`

	// 返回行数，格式为M-N，例如：10-200
	SentRowsSection *string `json:"SentRowsSection,omitnil,omitempty" name:"SentRowsSection"`

	// 执行时间，格式为M-N，例如：10-200
	ExecTimeSection *string `json:"ExecTimeSection,omitnil,omitempty" name:"ExecTimeSection"`

	// 锁等待时间，格式为M-N，例如：10-200
	LockWaitTimeSection *string `json:"LockWaitTimeSection,omitnil,omitempty" name:"LockWaitTimeSection"`

	// IO等待时间，格式为M-N，例如：10-200
	IoWaitTimeSection *string `json:"IoWaitTimeSection,omitnil,omitempty" name:"IoWaitTimeSection"`

	// 事务持续时间，格式为M-N，例如：10-200
	TransactionLivingTimeSection *string `json:"TransactionLivingTimeSection,omitnil,omitempty" name:"TransactionLivingTimeSection"`

	// 线程ID
	ThreadId []*string `json:"ThreadId,omitnil,omitempty" name:"ThreadId"`

	// 返回行数。表示筛选返回行数大于该值的审计日志。
	SentRows *int64 `json:"SentRows,omitnil,omitempty" name:"SentRows"`

	// mysql错误码
	ErrCode []*int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`
}

type AuditPolicy struct {
	// 审计策略 ID。
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 审计策略的状态。可能返回的值为：
	// "creating" - 创建中;
	// "running" - 运行中;
	// "paused" - 暂停中;
	// "failed" - 创建失败。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 数据库实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计策略创建时间。格式为 : "2019-03-20 17:09:13"。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 审计策略最后修改时间。格式为 : "2019-03-20 17:09:13"。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 审计策略名称。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 审计规则 ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 审计规则名称。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 数据库实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type AuditRule struct {
	// 审计规则 Id。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 审计规则创建时间。格式为 : "2019-03-20 17:09:13"。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 审计规则最后修改时间。格式为 : "2019-03-20 17:09:13"。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 审计规则名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 审计规则描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 审计规则过滤条件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleFilters []*AuditFilter `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// 是否开启全审计。
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`
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

	// 规则模板的过滤条件。
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// 规则模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 规则模板创建时间。
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// 告警等级。1-低风险，2-中风险，3-高风险。
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警策略。0-不告警，1-告警。
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`

	// 规则模板应用在哪些在实例。
	AffectedInstances []*string `json:"AffectedInstances,omitnil,omitempty" name:"AffectedInstances"`

	// 模板状态。0-无任务 ，1-修改中。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 模板更新时间。
	UpdateAt *string `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`
}

type AutoStrategy struct {
	// 自动扩容阈值，可选值40、50、60、70、80、90，代表 CPU 利用率达到40%、50%、60%、70%、80%、90%时后台进行自动扩容。
	ExpandThreshold *int64 `json:"ExpandThreshold,omitnil,omitempty" name:"ExpandThreshold"`

	// 自动缩容阈值，可选值10、20、30，代表CPU利用率达到10%、20%、30%时后台进行自动缩容
	ShrinkThreshold *int64 `json:"ShrinkThreshold,omitnil,omitempty" name:"ShrinkThreshold"`

	// 自动扩容观测周期，单位是分钟，可选值1、3、5、10、15、30。后台会按照配置的周期进行扩容判断。
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ExpandPeriod is deprecated.
	ExpandPeriod *int64 `json:"ExpandPeriod,omitnil,omitempty" name:"ExpandPeriod"`

	// 自动缩容观测周期，单位是分钟，可选值5、10、15、30。后台会按照配置的周期进行缩容判断。
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ShrinkPeriod is deprecated.
	ShrinkPeriod *int64 `json:"ShrinkPeriod,omitnil,omitempty" name:"ShrinkPeriod"`

	// 弹性扩容观测周期（秒级），可取值为：15，30，45，60，180，300，600，900，1800。
	ExpandSecondPeriod *int64 `json:"ExpandSecondPeriod,omitnil,omitempty" name:"ExpandSecondPeriod"`

	// 缩容观测周期（秒级），可取值为：300、600、900、1800。
	ShrinkSecondPeriod *int64 `json:"ShrinkSecondPeriod,omitnil,omitempty" name:"ShrinkSecondPeriod"`
}

type BackupConfig struct {
	// 第二个从库复制方式，可能的返回值：async-异步，semisync-半同步
	ReplicationMode *string `json:"ReplicationMode,omitnil,omitempty" name:"ReplicationMode"`

	// 第二个从库可用区的正式名称，如ap-shanghai-1
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 第二个从库内网IP地址
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 第二个从库访问端口
	Vport *uint64 `json:"Vport,omitnil,omitempty" name:"Vport"`
}

type BackupInfo struct {
	// 备份文件名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备份文件大小，单位：Byte
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 备份快照时间，时间格式：2016-03-17 02:10:37
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 下载地址
	IntranetUrl *string `json:"IntranetUrl,omitnil,omitempty" name:"IntranetUrl"`

	// 下载地址
	InternetUrl *string `json:"InternetUrl,omitnil,omitempty" name:"InternetUrl"`

	// 日志具体类型。可能的值有 "logical": 逻辑冷备， "physical": 物理冷备。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 备份子任务的ID，删除备份文件时使用
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 备份任务状态。可能的值有 "SUCCESS": 备份成功， "FAILED": 备份失败， "RUNNING": 备份进行中。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 备份任务的完成时间
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// （该值将废弃，不建议使用）备份的创建者，可能的值：SYSTEM - 系统创建，Uin - 发起者Uin值。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 备份任务的开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 备份方法。可能的值有 "full": 全量备份， "partial": 部分备份。
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 备份方式。可能的值有 "manual": 手动备份， "automatic": 自动备份。
	Way *string `json:"Way,omitnil,omitempty" name:"Way"`

	// 手动备份别名
	ManualBackupName *string `json:"ManualBackupName,omitnil,omitempty" name:"ManualBackupName"`

	// 备份保留类型，save_mode_regular - 常规保存备份，save_mode_period - 定期保存备份
	SaveMode *string `json:"SaveMode,omitnil,omitempty" name:"SaveMode"`

	// 本地备份所在地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 异地备份详细信息
	RemoteInfo []*RemoteBackupInfo `json:"RemoteInfo,omitnil,omitempty" name:"RemoteInfo"`

	// 存储方式，0-常规存储，1-归档存储，2-标准存储，默认为0
	CosStorageType *int64 `json:"CosStorageType,omitnil,omitempty" name:"CosStorageType"`

	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备份文件是否加密， on-加密， off-未加密
	EncryptionFlag *string `json:"EncryptionFlag,omitnil,omitempty" name:"EncryptionFlag"`

	// 备份GTID点位
	ExecutedGTIDSet *string `json:"ExecutedGTIDSet,omitnil,omitempty" name:"ExecutedGTIDSet"`

	// 备份文件MD5值
	MD5 *string `json:"MD5,omitnil,omitempty" name:"MD5"`
}

type BackupItem struct {
	// 需要备份的库名
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// 需要备份的表名。 如果传该参数，表示备份该库中的指定表。如果不传该参数则备份该db库
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`
}

type BackupLimitVpcItem struct {
	// 限制下载来源的地域。目前仅支持当前地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 限制下载的vpc列表。
	VpcList []*string `json:"VpcList,omitnil,omitempty" name:"VpcList"`
}

type BackupSummaryItem struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 该实例自动数据备份的个数。
	AutoBackupCount *int64 `json:"AutoBackupCount,omitnil,omitempty" name:"AutoBackupCount"`

	// 该实例自动数据备份的容量。
	AutoBackupVolume *int64 `json:"AutoBackupVolume,omitnil,omitempty" name:"AutoBackupVolume"`

	// 该实例手动数据备份的个数。
	ManualBackupCount *int64 `json:"ManualBackupCount,omitnil,omitempty" name:"ManualBackupCount"`

	// 该实例手动数据备份的容量。
	ManualBackupVolume *int64 `json:"ManualBackupVolume,omitnil,omitempty" name:"ManualBackupVolume"`

	// 该实例总的数据备份（包含自动备份和手动备份）个数。
	DataBackupCount *int64 `json:"DataBackupCount,omitnil,omitempty" name:"DataBackupCount"`

	// 该实例总的数据备份容量。
	DataBackupVolume *int64 `json:"DataBackupVolume,omitnil,omitempty" name:"DataBackupVolume"`

	// 该实例日志备份的个数。
	BinlogBackupCount *int64 `json:"BinlogBackupCount,omitnil,omitempty" name:"BinlogBackupCount"`

	// 该实例日志备份的容量。
	BinlogBackupVolume *int64 `json:"BinlogBackupVolume,omitnil,omitempty" name:"BinlogBackupVolume"`

	// 该实例的总备份（包含数据备份和日志备份）占用容量。
	BackupVolume *int64 `json:"BackupVolume,omitnil,omitempty" name:"BackupVolume"`
}

// Predefined struct for user
type BalanceRoGroupLoadRequestParams struct {
	// RO 组的 ID，格式如：cdbrg-c1nl9rpv。可通过 [DescribeRoGroups](https://cloud.tencent.com/document/api/236/40939) 获取。
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`
}

type BalanceRoGroupLoadRequest struct {
	*tchttp.BaseRequest
	
	// RO 组的 ID，格式如：cdbrg-c1nl9rpv。可通过 [DescribeRoGroups](https://cloud.tencent.com/document/api/236/40939) 获取。
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type BinlogInfo struct {
	// binlog 日志备份文件名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备份文件大小，单位：Byte
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 文件存储时间，时间格式：2016-03-17 02:10:37
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 下载地址
	// 说明：此下载地址和参数 InternetUrl 的下载地址一样。
	IntranetUrl *string `json:"IntranetUrl,omitnil,omitempty" name:"IntranetUrl"`

	// 下载地址
	// 说明：此下载地址和参数 IntranetUrl 的下载地址一样。
	InternetUrl *string `json:"InternetUrl,omitnil,omitempty" name:"InternetUrl"`

	// 日志具体类型，可能的值有：binlog - 二进制日志
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// binlog 文件起始时间
	BinlogStartTime *string `json:"BinlogStartTime,omitnil,omitempty" name:"BinlogStartTime"`

	// binlog 文件截止时间
	BinlogFinishTime *string `json:"BinlogFinishTime,omitnil,omitempty" name:"BinlogFinishTime"`

	// 本地binlog文件所在地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 备份任务状态。可能的值有 "SUCCESS": 备份成功， "FAILED": 备份失败， "RUNNING": 备份进行中。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// binlog异地备份详细信息
	RemoteInfo []*RemoteBackupInfo `json:"RemoteInfo,omitnil,omitempty" name:"RemoteInfo"`

	// 存储方式，0-常规存储，1-归档存储，2-标准存储，默认为0
	CosStorageType *int64 `json:"CosStorageType,omitnil,omitempty" name:"CosStorageType"`

	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	//
	// Deprecated: InstanceId is deprecated.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type Bucket struct {
	// 无
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// key值出现的次数。
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type CdbRegionSellConf struct {
	// 地域中文名称
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 所属大区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 是否为默认地域
	IsDefaultRegion *int64 `json:"IsDefaultRegion,omitnil,omitempty" name:"IsDefaultRegion"`

	// 地域名称
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 地域的可用区售卖配置
	RegionConfig []*CdbZoneSellConf `json:"RegionConfig,omitnil,omitempty" name:"RegionConfig"`
}

type CdbSellConfig struct {
	// 内存大小，单位为MB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// CPU核心数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 磁盘最小规格，单位为GB
	VolumeMin *int64 `json:"VolumeMin,omitnil,omitempty" name:"VolumeMin"`

	// 磁盘最大规格，单位为GB
	VolumeMax *int64 `json:"VolumeMax,omitnil,omitempty" name:"VolumeMax"`

	// 磁盘步长，单位为GB
	VolumeStep *int64 `json:"VolumeStep,omitnil,omitempty" name:"VolumeStep"`

	// 每秒IO数量
	Iops *int64 `json:"Iops,omitnil,omitempty" name:"Iops"`

	// 应用场景描述
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// 状态值，0 表示该规格对外售卖
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例类型，可能的取值范围有：UNIVERSAL (通用型), EXCLUSIVE (独享型), BASIC (基础型), BASIC_V2 (基础型v2)
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 引擎类型描述，可能的取值范围有：Innodb，RocksDB
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 售卖规格Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type CdbSellType struct {
	// 售卖实例名称。Z3是高可用类型对应规格中的DeviceType包含UNIVERSAL,EXCLUSIVE；CVM是基础版类型对应规格中的DeviceType是BASIC；TKE是基础型v2类型对应规格中的DeviceType是BASIC_V2。
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 引擎版本号
	EngineVersion []*string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 售卖规格Id
	ConfigIds []*int64 `json:"ConfigIds,omitnil,omitempty" name:"ConfigIds"`
}

type CdbZoneDataResult struct {
	// 售卖规格所有集合
	Configs []*CdbSellConfig `json:"Configs,omitnil,omitempty" name:"Configs"`

	// 售卖地域可用区集合
	Regions []*CdbRegionSellConf `json:"Regions,omitnil,omitempty" name:"Regions"`
}

type CdbZoneSellConf struct {
	// 可用区状态。可能的返回值为：1-上线；3-停售；4-不展示
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 可用区中文名称
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 实例类型是否为自定义类型
	IsCustom *bool `json:"IsCustom,omitnil,omitempty" name:"IsCustom"`

	// 是否支持灾备
	IsSupportDr *bool `json:"IsSupportDr,omitnil,omitempty" name:"IsSupportDr"`

	// 是否支持私有网络
	IsSupportVpc *bool `json:"IsSupportVpc,omitnil,omitempty" name:"IsSupportVpc"`

	// 小时计费实例最大售卖数量
	HourInstanceSaleMaxNum *int64 `json:"HourInstanceSaleMaxNum,omitnil,omitempty" name:"HourInstanceSaleMaxNum"`

	// 是否为默认可用区
	IsDefaultZone *bool `json:"IsDefaultZone,omitnil,omitempty" name:"IsDefaultZone"`

	// 是否为黑石区
	IsBm *bool `json:"IsBm,omitnil,omitempty" name:"IsBm"`

	// 支持的付费类型。可能的返回值为：0-包年包月；1-小时计费；2-后付费
	PayType []*string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// 数据复制类型。0-异步复制；1-半同步复制；2-强同步复制
	ProtectMode []*string `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// 可用区名称
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 多可用区信息
	ZoneConf *ZoneConf `json:"ZoneConf,omitnil,omitempty" name:"ZoneConf"`

	// 可支持的灾备可用区信息
	DrZone []*string `json:"DrZone,omitnil,omitempty" name:"DrZone"`

	// 是否支持跨可用区只读
	IsSupportRemoteRo *bool `json:"IsSupportRemoteRo,omitnil,omitempty" name:"IsSupportRemoteRo"`

	// 可支持的跨可用区只读区信息
	RemoteRoZone []*string `json:"RemoteRoZone,omitnil,omitempty" name:"RemoteRoZone"`

	// 独享型可用区状态。可能的返回值为：1-上线；3-停售；4-不展示
	ExClusterStatus *int64 `json:"ExClusterStatus,omitnil,omitempty" name:"ExClusterStatus"`

	// 独享型可支持的跨可用区只读区信息
	ExClusterRemoteRoZone []*string `json:"ExClusterRemoteRoZone,omitnil,omitempty" name:"ExClusterRemoteRoZone"`

	// 独享型多可用区信息
	ExClusterZoneConf *ZoneConf `json:"ExClusterZoneConf,omitnil,omitempty" name:"ExClusterZoneConf"`

	// 售卖实例类型数组，其中configIds的值与configs结构体中的id一一对应。
	SellType []*CdbSellType `json:"SellType,omitnil,omitempty" name:"SellType"`

	// 可用区id
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 是否支持ipv6
	IsSupportIpv6 *bool `json:"IsSupportIpv6,omitnil,omitempty" name:"IsSupportIpv6"`

	// 可支持的售卖数据库引擎类型
	EngineType []*string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 集群版实例在当前可用区的售卖状态。可能的返回值为：1-上线；3-停售；4-不展示
	CloudNativeClusterStatus *int64 `json:"CloudNativeClusterStatus,omitnil,omitempty" name:"CloudNativeClusterStatus"`

	// 集群版或者单节点基础型支持的磁盘类型。
	DiskTypeConf []*DiskTypeConfigItem `json:"DiskTypeConf,omitnil,omitempty" name:"DiskTypeConf"`
}

// Predefined struct for user
type CheckMigrateClusterRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例 CPU 核数。当 InstanceId 为主实例时必传。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例内存大小，单位：MB。当 InstanceId 为主实例时必传。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 磁盘类型。 CLOUD_SSD: SSD 云硬盘; CLOUD_HSSD: 增强型 SSD 云硬盘。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 云盘版节点拓扑配置。当 InstanceId 为主实例时必传。
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// 迁移实例类型。支持值包括： "CLOUD_NATIVE_CLUSTER" - 标准型云盘版实例， "CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - 加强型云盘版实例。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 只读实例信息。
	RoInfo []*MigrateClusterRoInfo `json:"RoInfo,omitnil,omitempty" name:"RoInfo"`
}

type CheckMigrateClusterRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例 CPU 核数。当 InstanceId 为主实例时必传。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例内存大小，单位：MB。当 InstanceId 为主实例时必传。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 磁盘类型。 CLOUD_SSD: SSD 云硬盘; CLOUD_HSSD: 增强型 SSD 云硬盘。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 云盘版节点拓扑配置。当 InstanceId 为主实例时必传。
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// 迁移实例类型。支持值包括： "CLOUD_NATIVE_CLUSTER" - 标准型云盘版实例， "CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - 加强型云盘版实例。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 只读实例信息。
	RoInfo []*MigrateClusterRoInfo `json:"RoInfo,omitnil,omitempty" name:"RoInfo"`
}

func (r *CheckMigrateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckMigrateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "DiskType")
	delete(f, "ClusterTopology")
	delete(f, "DeviceType")
	delete(f, "RoInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckMigrateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckMigrateClusterResponseParams struct {
	// 校验是否通过，通过为pass，失败为fail
	CheckResult *string `json:"CheckResult,omitnil,omitempty" name:"CheckResult"`

	// 校验项
	Items []*CheckMigrateResult `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckMigrateClusterResponse struct {
	*tchttp.BaseResponse
	Response *CheckMigrateClusterResponseParams `json:"Response"`
}

func (r *CheckMigrateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckMigrateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckMigrateResult struct {
	// 校验名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 校验结果，通过为pass，失败为fail
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 校验结果描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type CloneItem struct {
	// 克隆任务的源实例Id。
	SrcInstanceId *string `json:"SrcInstanceId,omitnil,omitempty" name:"SrcInstanceId"`

	// 克隆任务的新产生实例Id。
	DstInstanceId *string `json:"DstInstanceId,omitnil,omitempty" name:"DstInstanceId"`

	// 克隆任务对应的任务列表Id。
	CloneJobId *int64 `json:"CloneJobId,omitnil,omitempty" name:"CloneJobId"`

	// 克隆实例使用的策略， 包括以下类型：  timepoint:指定时间点回档，  backupset: 指定备份文件回档。
	RollbackStrategy *string `json:"RollbackStrategy,omitnil,omitempty" name:"RollbackStrategy"`

	// 克隆实例回档的时间点。
	RollbackTargetTime *string `json:"RollbackTargetTime,omitnil,omitempty" name:"RollbackTargetTime"`

	// 任务开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务状态，包括以下状态：initial,running,wait_complete,success,failed
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 克隆实例所在地域Id
	NewRegionId *int64 `json:"NewRegionId,omitnil,omitempty" name:"NewRegionId"`

	// 源实例所在地域Id
	SrcRegionId *int64 `json:"SrcRegionId,omitnil,omitempty" name:"SrcRegionId"`
}

// Predefined struct for user
type CloseAuditServiceRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CloseAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
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
type CloseCDBProxyRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 代理组 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 是否只关闭读写分离，取值："true" | "false"，默认为"false"
	OnlyCloseRW *bool `json:"OnlyCloseRW,omitnil,omitempty" name:"OnlyCloseRW"`
}

type CloseCDBProxyRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 代理组 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 是否只关闭读写分离，取值："true" | "false"，默认为"false"
	OnlyCloseRW *bool `json:"OnlyCloseRW,omitnil,omitempty" name:"OnlyCloseRW"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CloseCdbProxyAddressRequestParams struct {
	// 代理组 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 代理组地址 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`
}

type CloseCdbProxyAddressRequest struct {
	*tchttp.BaseRequest
	
	// 代理组 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 代理组地址 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`
}

func (r *CloseCdbProxyAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseCdbProxyAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyGroupId")
	delete(f, "ProxyAddressId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseCdbProxyAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseCdbProxyAddressResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseCdbProxyAddressResponse struct {
	*tchttp.BaseResponse
	Response *CloseCdbProxyAddressResponseParams `json:"Response"`
}

func (r *CloseCdbProxyAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseCdbProxyAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseSSLRequestParams struct {
	// 实例 ID。只读组 ID 为空时必填。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 只读组 ID。实例 ID 为空时必填。可通过 [DescribeRoGroups](https://cloud.tencent.com/document/api/236/40939) 接口获取。
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`
}

type CloseSSLRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。只读组 ID 为空时必填。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 只读组 ID。实例 ID 为空时必填。可通过 [DescribeRoGroups](https://cloud.tencent.com/document/api/236/40939) 接口获取。
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`
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
	delete(f, "InstanceId")
	delete(f, "RoGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseSSLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseSSLResponseParams struct {
	// 异步请求 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

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
type CloseWanServiceRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。可传入只读组 ID 关闭只读组外网访问。 
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 变更云盘版实例只读组时，InstanceId 传实例 ID，需要额外指定该参数表示操作只读组。如果操作读写节点则不需指定该参数。
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
}

type CloseWanServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。可传入只读组 ID 关闭只读组外网访问。 
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 变更云盘版实例只读组时，InstanceId 传实例 ID，需要额外指定该参数表示操作只读组。如果操作读写节点则不需指定该参数。
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
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
	delete(f, "OpResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseWanServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseWanServiceResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type ClusterInfo struct {
	// 节点id
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点类型：主节点，从节点
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 地域
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type ClusterNodeInfo struct {
	// 节点id。
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点的角色。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 节点所在可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 节点的权重
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 节点状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ClusterTopology struct {
	// RW 节点拓扑。
	// 说明：NodeId 可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 获取。
	ReadWriteNode *ReadWriteNode `json:"ReadWriteNode,omitnil,omitempty" name:"ReadWriteNode"`

	// RO 节点拓扑。
	// 说明：NodeId 可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 获取。
	ReadOnlyNodes []*ReadonlyNode `json:"ReadOnlyNodes,omitnil,omitempty" name:"ReadOnlyNodes"`
}

type ColumnPrivilege struct {
	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 数据库表名
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 数据库列名
	Column *string `json:"Column,omitnil,omitempty" name:"Column"`

	// 权限信息
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

type CommonTimeWindow struct {
	// 周一的时间窗，格式如： 02:00-06:00
	Monday *string `json:"Monday,omitnil,omitempty" name:"Monday"`

	// 周二的时间窗，格式如： 02:00-06:00
	Tuesday *string `json:"Tuesday,omitnil,omitempty" name:"Tuesday"`

	// 周三的时间窗，格式如： 02:00-06:00
	Wednesday *string `json:"Wednesday,omitnil,omitempty" name:"Wednesday"`

	// 周四的时间窗，格式如： 02:00-06:00
	Thursday *string `json:"Thursday,omitnil,omitempty" name:"Thursday"`

	// 周五的时间窗，格式如： 02:00-06:00
	Friday *string `json:"Friday,omitnil,omitempty" name:"Friday"`

	// 周六的时间窗，格式如： 02:00-06:00
	Saturday *string `json:"Saturday,omitnil,omitempty" name:"Saturday"`

	// 周日的时间窗，格式如： 02:00-06:00
	Sunday *string `json:"Sunday,omitnil,omitempty" name:"Sunday"`

	// 常规备份保留策略，weekly-按周备份，monthly-按月备份，默认为weekly
	BackupPeriodStrategy *string `json:"BackupPeriodStrategy,omitnil,omitempty" name:"BackupPeriodStrategy"`

	// 如果设置为按月备份，需填入每月具体备份日期，相邻备份天数不得超过两天。例[1,4,7,9,11,14,17,19,22,25,28,30,31]
	Days []*int64 `json:"Days,omitnil,omitempty" name:"Days"`

	// 月度备份时间窗，BackupPeriodStrategy为monthly时必填。格式如： 02:00-06:00
	BackupPeriodTime *string `json:"BackupPeriodTime,omitnil,omitempty" name:"BackupPeriodTime"`
}

// Predefined struct for user
type CreateAccountsRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 云数据库账号。
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// 新账户的密码。
	// 说明：
	// 1. 在8 ～ 64位字符数以内（推荐12位以上）。
	// 2. 至少包含其中两项：小写字母 a ~ z 或 大写字母 A ～ Z。数字0 ～ 9。_+-,&=!@#$%^*().|。
	// 3. 不能包含非法字符。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 备注信息。最多支持输入255个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 新账户最大可用连接数，默认值为10240，最大可设置值为10240。
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`
}

type CreateAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 云数据库账号。
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// 新账户的密码。
	// 说明：
	// 1. 在8 ～ 64位字符数以内（推荐12位以上）。
	// 2. 至少包含其中两项：小写字母 a ~ z 或 大写字母 A ～ Z。数字0 ～ 9。_+-,&=!@#$%^*().|。
	// 3. 不能包含非法字符。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 备注信息。最多支持输入255个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 新账户最大可用连接数，默认值为10240，最大可设置值为10240。
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

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
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间(建议开始到结束时间区间最大7天)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间(建议开始到结束时间区间最大7天）。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序方式。支持值包括："ASC" - 升序，"DESC" - 降序，默认降序排序。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段。支持值包括(默认按照时间戳排序)： "timestamp" - 时间戳； "affectRows" - 影响行数； "execTime" - 执行时间。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 已废弃。
	//
	// Deprecated: Filter is deprecated.
	Filter *AuditLogFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 过滤条件。可按设置的过滤条件过滤日志。
	LogFilter []*InstanceAuditLogFilters `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`

	// 下载筛选列
	ColumnFilter []*string `json:"ColumnFilter,omitnil,omitempty" name:"ColumnFilter"`
}

type CreateAuditLogFileRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间(建议开始到结束时间区间最大7天)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间(建议开始到结束时间区间最大7天）。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序方式。支持值包括："ASC" - 升序，"DESC" - 降序，默认降序排序。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段。支持值包括(默认按照时间戳排序)： "timestamp" - 时间戳； "affectRows" - 影响行数； "execTime" - 执行时间。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 已废弃。
	Filter *AuditLogFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 过滤条件。可按设置的过滤条件过滤日志。
	LogFilter []*InstanceAuditLogFilters `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`

	// 下载筛选列
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
type CreateAuditPolicyRequestParams struct {
	// 审计策略名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 审计规则 ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计日志保存时长。支持值包括：
	// 7 - 一周
	// 30 - 一个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年；
	// 实例首次开通审计策略时，可传该值，用于设置存储日志保存天数，默认为 30 天。若实例已存在审计策略，则此参数无效，可使用 更改审计服务配置 接口修改日志存储时长。
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`
}

type CreateAuditPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 审计策略名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 审计规则 ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计日志保存时长。支持值包括：
	// 7 - 一周
	// 30 - 一个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年；
	// 实例首次开通审计策略时，可传该值，用于设置存储日志保存天数，默认为 30 天。若实例已存在审计策略，则此参数无效，可使用 更改审计服务配置 接口修改日志存储时长。
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`
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
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 审计规则描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 审计规则过滤条件。若设置了过滤条件，则不会开启全审计。
	RuleFilters []*AuditFilter `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// 是否开启全审计。支持值包括：false – 不开启全审计，true – 开启全审计。用户未设置审计规则过滤条件时，默认开启全审计。
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`
}

type CreateAuditRuleRequest struct {
	*tchttp.BaseRequest
	
	// 审计规则名称。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 审计规则描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 审计规则过滤条件。若设置了过滤条件，则不会开启全审计。
	RuleFilters []*AuditFilter `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// 是否开启全审计。支持值包括：false – 不开启全审计，true – 开启全审计。用户未设置审计规则过滤条件时，默认开启全审计。
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`
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
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateAuditRuleTemplateRequestParams struct {
	// 审计规则。
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// 规则模板名称。最多支持输入30个字符。
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// 规则模板描述。最多支持输入200个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 告警等级。1 - 低风险，2 - 中风险，3 - 高风险。默认值为1。
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警策略。0 - 不告警，1 - 告警。默认值为0。
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`
}

type CreateAuditRuleTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 审计规则。
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// 规则模板名称。最多支持输入30个字符。
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// 规则模板描述。最多支持输入200个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 告警等级。1 - 低风险，2 - 中风险，3 - 高风险。默认值为1。
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警策略。0 - 不告警，1 - 告警。默认值为0。
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
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 目标备份方法，可选的值：logical - 逻辑冷备，physical - 物理冷备，snapshot - 快照备份。基础版实例仅支持快照备份。
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 需要备份的库表信息，如果不设置该参数，则默认整实例备份。在 BackupMethod=logical 逻辑备份中才可设置该参数。指定的库表必须存在，否则可能导致备份失败。
	// 例：如果需要备份 db1 库的 tb1、tb2 表 和 db2 库。则该参数设置为 [{"Db": "db1", "Table": "tb1"}, {"Db": "db1", "Table": "tb2"}, {"Db": "db2"}]。
	BackupDBTableList []*BackupItem `json:"BackupDBTableList,omitnil,omitempty" name:"BackupDBTableList"`

	// 手动备份别名，输入长度请在60个字符内。
	ManualBackupName *string `json:"ManualBackupName,omitnil,omitempty" name:"ManualBackupName"`

	// 是否需要加密物理备份，可选值为：on - 是，off - 否。当 BackupMethod 为 physical 时，该值才有意义。不指定则使用实例备份默认加密策略，这里的默认加密策略指通过 [DescribeBackupEncryptionStatus](https://cloud.tencent.com/document/product/236/86508) 接口查询出的实例当前加密策略。
	EncryptionFlag *string `json:"EncryptionFlag,omitnil,omitempty" name:"EncryptionFlag"`
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 目标备份方法，可选的值：logical - 逻辑冷备，physical - 物理冷备，snapshot - 快照备份。基础版实例仅支持快照备份。
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 需要备份的库表信息，如果不设置该参数，则默认整实例备份。在 BackupMethod=logical 逻辑备份中才可设置该参数。指定的库表必须存在，否则可能导致备份失败。
	// 例：如果需要备份 db1 库的 tb1、tb2 表 和 db2 库。则该参数设置为 [{"Db": "db1", "Table": "tb1"}, {"Db": "db1", "Table": "tb2"}, {"Db": "db2"}]。
	BackupDBTableList []*BackupItem `json:"BackupDBTableList,omitnil,omitempty" name:"BackupDBTableList"`

	// 手动备份别名，输入长度请在60个字符内。
	ManualBackupName *string `json:"ManualBackupName,omitnil,omitempty" name:"ManualBackupName"`

	// 是否需要加密物理备份，可选值为：on - 是，off - 否。当 BackupMethod 为 physical 时，该值才有意义。不指定则使用实例备份默认加密策略，这里的默认加密策略指通过 [DescribeBackupEncryptionStatus](https://cloud.tencent.com/document/product/236/86508) 接口查询出的实例当前加密策略。
	EncryptionFlag *string `json:"EncryptionFlag,omitnil,omitempty" name:"EncryptionFlag"`
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
	delete(f, "EncryptionFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupResponseParams struct {
	// 备份任务 ID。
	BackupId *uint64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

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
type CreateCdbProxyAddressRequestParams struct {
	// 代理组 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 权重分配模式，
	// 系统自动分配："system"， 自定义："custom"
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// 是否开启延迟剔除，取值："true" | "false"
	IsKickOut *bool `json:"IsKickOut,omitnil,omitempty" name:"IsKickOut"`

	// 最小保留数量，最小取值：0
	MinCount *uint64 `json:"MinCount,omitnil,omitempty" name:"MinCount"`

	// 延迟剔除阈值，最小取值：1，范围：1 - 10000，整数。
	MaxDelay *uint64 `json:"MaxDelay,omitnil,omitempty" name:"MaxDelay"`

	// 是否开启故障转移，取值："true" | "false"
	FailOver *bool `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// 是否自动添加RO，取值："true" | "false"
	AutoAddRo *bool `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// 是否是只读，取值："true" | "false"
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// 是否开启事务分离，取值："true" | "false"
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// 读写权重分配
	ProxyAllocation []*ProxyAllocation `json:"ProxyAllocation,omitnil,omitempty" name:"ProxyAllocation"`

	// 私有网络 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 私有子网 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 是否开启连接池。默认关闭。
	// 注意：如需使用数据库代理连接池能力，MySQL 8.0 主实例的内核小版本要大于等于 MySQL 8.0 20230630。
	ConnectionPool *bool `json:"ConnectionPool,omitnil,omitempty" name:"ConnectionPool"`

	// 描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// IP 地址。不填则默认为所选 VPC 下支持的随机一个 IP。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 端口。默认值3306。
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// 安全组
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 连接池类型。可选值 transaction（事务级别连接池），connection（会话级别连接池），ConnectionPool 为 true 时生效。默认值：connection。
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// 是否开启自适应负载均衡。默认关闭。
	AutoLoadBalance *bool `json:"AutoLoadBalance,omitnil,omitempty" name:"AutoLoadBalance"`

	// 接入模式。nearBy - 就近访问，balance - 均衡分配，默认值：nearBy。
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`
}

type CreateCdbProxyAddressRequest struct {
	*tchttp.BaseRequest
	
	// 代理组 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 权重分配模式，
	// 系统自动分配："system"， 自定义："custom"
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// 是否开启延迟剔除，取值："true" | "false"
	IsKickOut *bool `json:"IsKickOut,omitnil,omitempty" name:"IsKickOut"`

	// 最小保留数量，最小取值：0
	MinCount *uint64 `json:"MinCount,omitnil,omitempty" name:"MinCount"`

	// 延迟剔除阈值，最小取值：1，范围：1 - 10000，整数。
	MaxDelay *uint64 `json:"MaxDelay,omitnil,omitempty" name:"MaxDelay"`

	// 是否开启故障转移，取值："true" | "false"
	FailOver *bool `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// 是否自动添加RO，取值："true" | "false"
	AutoAddRo *bool `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// 是否是只读，取值："true" | "false"
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// 是否开启事务分离，取值："true" | "false"
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// 读写权重分配
	ProxyAllocation []*ProxyAllocation `json:"ProxyAllocation,omitnil,omitempty" name:"ProxyAllocation"`

	// 私有网络 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 私有子网 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 是否开启连接池。默认关闭。
	// 注意：如需使用数据库代理连接池能力，MySQL 8.0 主实例的内核小版本要大于等于 MySQL 8.0 20230630。
	ConnectionPool *bool `json:"ConnectionPool,omitnil,omitempty" name:"ConnectionPool"`

	// 描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// IP 地址。不填则默认为所选 VPC 下支持的随机一个 IP。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 端口。默认值3306。
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// 安全组
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 连接池类型。可选值 transaction（事务级别连接池），connection（会话级别连接池），ConnectionPool 为 true 时生效。默认值：connection。
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// 是否开启自适应负载均衡。默认关闭。
	AutoLoadBalance *bool `json:"AutoLoadBalance,omitnil,omitempty" name:"AutoLoadBalance"`

	// 接入模式。nearBy - 就近访问，balance - 均衡分配，默认值：nearBy。
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`
}

func (r *CreateCdbProxyAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCdbProxyAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyGroupId")
	delete(f, "WeightMode")
	delete(f, "IsKickOut")
	delete(f, "MinCount")
	delete(f, "MaxDelay")
	delete(f, "FailOver")
	delete(f, "AutoAddRo")
	delete(f, "ReadOnly")
	delete(f, "TransSplit")
	delete(f, "ProxyAllocation")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "ConnectionPool")
	delete(f, "Desc")
	delete(f, "Vip")
	delete(f, "VPort")
	delete(f, "SecurityGroup")
	delete(f, "ConnectionPoolType")
	delete(f, "AutoLoadBalance")
	delete(f, "AccessMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCdbProxyAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCdbProxyAddressResponseParams struct {
	// 异步任务ID
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCdbProxyAddressResponse struct {
	*tchttp.BaseResponse
	Response *CreateCdbProxyAddressResponseParams `json:"Response"`
}

func (r *CreateCdbProxyAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCdbProxyAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCdbProxyRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 私有网络 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 私有子网 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 节点规格配置。
	// 示例中参数说明：
	// NodeCount：节点个数。
	// Region：节点地域。
	// Zone：节点可用区。
	// Cpu：单个代理节点核数（单位：核）。
	// Mem：单个代理节点内存数（单位：MB）。
	// 备注：
	// 1. 数据库代理支持的节点规格为：2C4000MB、4C8000MB、8C16000MB。
	// 2. 上述参数项（如节点个数、可用区等）均为必填，在调用接口时如未填写完整，可能会创建失败。
	ProxyNodeCustom []*ProxyNodeCustom `json:"ProxyNodeCustom,omitnil,omitempty" name:"ProxyNodeCustom"`

	// 安全组
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 连接池阈值
	// 注意：如需使用数据库代理连接池能力，MySQL 8.0 主实例的内核小版本要大于等于 MySQL 8.0 20230630。
	ConnectionPoolLimit *uint64 `json:"ConnectionPoolLimit,omitnil,omitempty" name:"ConnectionPoolLimit"`

	// 指定要购买的 proxy 内核版本。不填则默认发货最新版本的 proxy。
	ProxyVersion *string `json:"ProxyVersion,omitnil,omitempty" name:"ProxyVersion"`
}

type CreateCdbProxyRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 私有网络 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 私有子网 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 节点规格配置。
	// 示例中参数说明：
	// NodeCount：节点个数。
	// Region：节点地域。
	// Zone：节点可用区。
	// Cpu：单个代理节点核数（单位：核）。
	// Mem：单个代理节点内存数（单位：MB）。
	// 备注：
	// 1. 数据库代理支持的节点规格为：2C4000MB、4C8000MB、8C16000MB。
	// 2. 上述参数项（如节点个数、可用区等）均为必填，在调用接口时如未填写完整，可能会创建失败。
	ProxyNodeCustom []*ProxyNodeCustom `json:"ProxyNodeCustom,omitnil,omitempty" name:"ProxyNodeCustom"`

	// 安全组
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 连接池阈值
	// 注意：如需使用数据库代理连接池能力，MySQL 8.0 主实例的内核小版本要大于等于 MySQL 8.0 20230630。
	ConnectionPoolLimit *uint64 `json:"ConnectionPoolLimit,omitnil,omitempty" name:"ConnectionPoolLimit"`

	// 指定要购买的 proxy 内核版本。不填则默认发货最新版本的 proxy。
	ProxyVersion *string `json:"ProxyVersion,omitnil,omitempty" name:"ProxyVersion"`
}

func (r *CreateCdbProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCdbProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "ProxyNodeCustom")
	delete(f, "SecurityGroup")
	delete(f, "Desc")
	delete(f, "ConnectionPoolLimit")
	delete(f, "ProxyVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCdbProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCdbProxyResponseParams struct {
	// 异步任务ID
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCdbProxyResponse struct {
	*tchttp.BaseResponse
	Response *CreateCdbProxyResponseParams `json:"Response"`
}

func (r *CreateCdbProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCdbProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloneInstanceRequestParams struct {
	// 克隆源实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/api/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 如果需要克隆实例回档到指定时间，则指定该值。时间格式为：yyyy-mm-dd hh:mm:ss。
	// 说明：此参数和 SpecifiedBackupId 参数需要2选1进行设置。
	SpecifiedRollbackTime *string `json:"SpecifiedRollbackTime,omitnil,omitempty" name:"SpecifiedRollbackTime"`

	// 如果需要克隆实例回档到指定备份集，则指定该值为备份文件的 Id。请使用 [查询数据备份文件列表](/document/api/236/15842)。
	// 说明：如果是克隆双节点、三节点实例，备份文件为物理备份，如果是克隆单节点、集群版实例，备份文件为快照备份。
	SpecifiedBackupId *int64 `json:"SpecifiedBackupId,omitnil,omitempty" name:"SpecifiedBackupId"`

	// 私有网络 ID，请使用 [查询私有网络列表](/document/api/215/15778)。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 私有网络下的子网 ID，如果设置了 UniqVpcId，则 UniqSubnetId 必填，请使用 [查询子网列表](/document/api/215/15784)。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 实例内存大小，单位：MB，需要不低于克隆源实例，默认和源实例相同。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB，需要不低于克隆源实例，默认和源实例相同。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 新产生的克隆实例名称。支持输入最大60个字符。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 安全组参数，可使用 [查询项目安全组信息](https://cloud.tencent.com/document/api/236/15850) 接口查询某个项目的安全组详情。
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 实例标签信息。
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 实例Cpu核数，需要不低于克隆源实例，默认和源实例相同。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制。
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// 多可用区域，默认为 0，支持值包括：0 - 表示单可用区，1 - 表示多可用区。
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 新产生的克隆实例备库 1 的可用区信息，默认同源实例 Zone 的值。
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// 备库 2 的可用区信息，默认为空，克隆强同步主实例时可指定该参数。
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// 克隆实例类型。支持值包括："UNIVERSAL" - 通用型实例，"EXCLUSIVE" - 独享型实例，"CLOUD_NATIVE_CLUSTER" - 集群版标准型，"CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - 集群版加强型。不指定则默认为通用型。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 新克隆实例节点数。如果需要克隆出三节点实例， 请将该值设置为3 或指定 BackupZone 参数。如果需要克隆出两节点实例，请将该值设置为2。默认克隆出两节点实例。
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// 是否只预检此次请求。true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制等。如果检查不通过，则返回对应错误码；如果检查通过，则返回RequestId.默认为false：发送正常请求，通过检查后直接创建实例。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 金融围拢 ID 。
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// 项目ID，默认项目ID0
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 付费类型，PRE_PAID：包年包月，USED_PAID：按量计费。默认为按量计费
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// 实例时长，PayType为PRE_PAID时必传，单位：月，可选值包括 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 集群版节点拓扑配置。
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// 原实例所在地域名，当传入异地备份时为必选项，例：ap-guangzhou
	SrcRegion *string `json:"SrcRegion,omitnil,omitempty" name:"SrcRegion"`

	// 异地数据备份id
	SpecifiedSubBackupId *int64 `json:"SpecifiedSubBackupId,omitnil,omitempty" name:"SpecifiedSubBackupId"`
}

type CreateCloneInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 克隆源实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/api/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 如果需要克隆实例回档到指定时间，则指定该值。时间格式为：yyyy-mm-dd hh:mm:ss。
	// 说明：此参数和 SpecifiedBackupId 参数需要2选1进行设置。
	SpecifiedRollbackTime *string `json:"SpecifiedRollbackTime,omitnil,omitempty" name:"SpecifiedRollbackTime"`

	// 如果需要克隆实例回档到指定备份集，则指定该值为备份文件的 Id。请使用 [查询数据备份文件列表](/document/api/236/15842)。
	// 说明：如果是克隆双节点、三节点实例，备份文件为物理备份，如果是克隆单节点、集群版实例，备份文件为快照备份。
	SpecifiedBackupId *int64 `json:"SpecifiedBackupId,omitnil,omitempty" name:"SpecifiedBackupId"`

	// 私有网络 ID，请使用 [查询私有网络列表](/document/api/215/15778)。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 私有网络下的子网 ID，如果设置了 UniqVpcId，则 UniqSubnetId 必填，请使用 [查询子网列表](/document/api/215/15784)。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 实例内存大小，单位：MB，需要不低于克隆源实例，默认和源实例相同。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB，需要不低于克隆源实例，默认和源实例相同。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 新产生的克隆实例名称。支持输入最大60个字符。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 安全组参数，可使用 [查询项目安全组信息](https://cloud.tencent.com/document/api/236/15850) 接口查询某个项目的安全组详情。
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 实例标签信息。
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 实例Cpu核数，需要不低于克隆源实例，默认和源实例相同。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制。
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// 多可用区域，默认为 0，支持值包括：0 - 表示单可用区，1 - 表示多可用区。
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 新产生的克隆实例备库 1 的可用区信息，默认同源实例 Zone 的值。
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// 备库 2 的可用区信息，默认为空，克隆强同步主实例时可指定该参数。
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// 克隆实例类型。支持值包括："UNIVERSAL" - 通用型实例，"EXCLUSIVE" - 独享型实例，"CLOUD_NATIVE_CLUSTER" - 集群版标准型，"CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - 集群版加强型。不指定则默认为通用型。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 新克隆实例节点数。如果需要克隆出三节点实例， 请将该值设置为3 或指定 BackupZone 参数。如果需要克隆出两节点实例，请将该值设置为2。默认克隆出两节点实例。
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// 是否只预检此次请求。true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制等。如果检查不通过，则返回对应错误码；如果检查通过，则返回RequestId.默认为false：发送正常请求，通过检查后直接创建实例。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 金融围拢 ID 。
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// 项目ID，默认项目ID0
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 付费类型，PRE_PAID：包年包月，USED_PAID：按量计费。默认为按量计费
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// 实例时长，PayType为PRE_PAID时必传，单位：月，可选值包括 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 集群版节点拓扑配置。
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// 原实例所在地域名，当传入异地备份时为必选项，例：ap-guangzhou
	SrcRegion *string `json:"SrcRegion,omitnil,omitempty" name:"SrcRegion"`

	// 异地数据备份id
	SpecifiedSubBackupId *int64 `json:"SpecifiedSubBackupId,omitnil,omitempty" name:"SpecifiedSubBackupId"`
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
	delete(f, "PayType")
	delete(f, "Period")
	delete(f, "ClusterTopology")
	delete(f, "SrcRegion")
	delete(f, "SpecifiedSubBackupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloneInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloneInstanceResponseParams struct {
	// 异步任务的请求ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 云数据库的用户名。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 文件名称。该文件是指用户已上传到腾讯云的文件，仅支持.sql文件。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 云数据库实例 User 账号的密码。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 导入的目标数据库名，不传表示不指定数据库。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 腾讯云COS文件链接。 用户需要指定 FileName 或者 CosUrl 其中一个。 COS文件需要是 .sql 文件。
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`
}

type CreateDBImportJobRequest struct {
	*tchttp.BaseRequest
	
	// 实例的 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 云数据库的用户名。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 文件名称。该文件是指用户已上传到腾讯云的文件，仅支持.sql文件。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 云数据库实例 User 账号的密码。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 导入的目标数据库名，不传表示不指定数据库。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 腾讯云COS文件链接。 用户需要指定 FileName 或者 CosUrl 其中一个。 COS文件需要是 .sql 文件。
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 实例内存大小，单位：MB，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的内存规格。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的硬盘范围。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// MySQL 版本，值包括：5.5、5.6、5.7和8.0，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的实例版本。
	// 说明：创建非集群版实例时，请根据需要指定实例版本（推荐5.7或8.0），若此参数不填，则默认值为5.6；若创建的是集群版实例，则此参数仅能指定为5.7或8.0。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 私有网络 ID，请使用 [查询私有网络列表](/document/api/215/15778)。
	// 说明：如果创建的是集群版实例，此参数为必填且为私有网络类型。若此项不填，则系统会选择默认的 VPC。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 私有网络下的子网 ID，如果设置了 UniqVpcId，则 UniqSubnetId 必填，请使用 [查询子网列表](/document/api/215/15784)。
	// 说明：若此项不填，则系统会选择默认 VPC 下的默认子网。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 项目 ID，不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 可用区信息，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的可用区。
	// 说明：若您创建单节点、双节点、三节点实例，此参数为必填项，请指定可用区，若不指定可用区，则系统会自动选择一个可用区（可能不是您希望部署的可用区）；若您创建集群版实例，此参数不填，请通过参数 ClusterTopology 进行读写节点和只读节点的可用区配置。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例 ID，购买只读实例或者灾备实例时必填，该字段表示只读实例或者灾备实例的主实例 ID，请使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口查询云数据库实例 ID。
	MasterInstanceId *string `json:"MasterInstanceId,omitnil,omitempty" name:"MasterInstanceId"`

	// 实例类型，支持值包括：master - 表示主实例，dr - 表示灾备实例，ro - 表示只读实例。
	// 说明：请选择实例类型，不填会默认选择 master。
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 主实例地域信息，购买灾备、RO实例时，该字段必填。
	MasterRegion *string `json:"MasterRegion,omitnil,omitempty" name:"MasterRegion"`

	// 自定义端口，端口支持范围：[1024 - 65535]。
	// 说明：不填则默认为3306。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 设置 root 账号密码，密码规则：8 - 64 个字符，至少包含字母、数字、字符（支持的字符：_+-&=!@#$%^*()）中的两种，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 参数列表，参数格式如 ParamList.0.Name=auto_increment&ParamList.0.Value=1。可通过 [查询默认的可设置参数列表](https://cloud.tencent.com/document/api/236/32662) 查询支持设置的参数。
	// 说明：表名大小写敏感的开启和关闭可通过参数 lower_case_table_names 进行设置，参数值为0表示开启，参数值为1表示关闭，若不设置则此参数默认值为0。若您创建的是 MySQL 8.0 版本的实例，则需要在创建实例时通过设置 lower_case_table_names 参数来开启或关闭表名大小写敏感，创建实例后无法修改参数，即创建后无法修改表名大小写敏感。其他数据库版本的实例支持在创建实例后修改 lower_case_table_names 参数。创建实例时设置表名大小写敏感的 API 调用方法请参见本文中的示例2。
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// 多可用区域，默认为 0，支持值包括：0 - 表示单可用区，1 - 表示多可用区，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 备库 1 的可用区信息。
	// 说明：双节点、三节点实例请指定此参数值，若不指定，则默认为 Zone 的值；集群版实例此参数可不填，请通过参数 ClusterTopology 进行读写节点和只读节点的可用区配置；单节点实例为单可用区，无需指定此参数。
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// 备库 2 的可用区信息，默认为空，购买三节点主实例时可指定该参数。
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// 安全组参数，可使用 [查询项目安全组信息](https://cloud.tencent.com/document/api/236/15850) 接口查询某个项目的安全组详情。
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 只读实例信息。购买只读实例时，该参数必传。
	RoGroup *RoGroup `json:"RoGroup,omitnil,omitempty" name:"RoGroup"`

	// 购买按量计费实例该字段无意义。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 实例名称。一次购买多个实例命名会用后缀数字区分，例instanceName=db，goodsNum=3，实例命名分别为db1，db2，db3。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例标签信息。
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间在48小时内唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 实例隔离类型。支持值包括："UNIVERSAL" - 通用型实例，"EXCLUSIVE" - 独享型实例，"BASIC_V2" - ONTKE 单节点实例，"CLOUD_NATIVE_CLUSTER" - 集群版标准型，"CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - 集群版加强型。不指定则默认为通用型实例。
	// 说明：如果创建的是集群版实例，此参数为必填。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 参数模板 id。
	// 备注：如您使用自定义参数模板 id，可传入自定义参数模板 id；如您计划使用默认参数模板，该参数模板 id 传入 id 无效，需设置 ParamTemplateType。
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 告警策略id数组。腾讯云可观测平台DescribeAlarmPolicy接口返回的OriginId。
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// 实例节点数。对于 RO 和 基础版实例， 该值默认为1。 如果需要购买三节点实例， 请将该值设置为3 或指定 BackupZone 参数。当购买主实例，且未指定该参数和 BackupZone 参数时，该值默认是 2， 即购买两节点实例。
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// 实例cpu核数， 如果不传将根据memory指定的内存值自动填充对应的cpu值。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 是否自动发起灾备同步功能。该参数仅对购买灾备实例生效。 可选值为：0 - 不自动发起灾备同步；1 - 自动发起灾备同步。该值默认为0。
	AutoSyncFlag *int64 `json:"AutoSyncFlag,omitnil,omitempty" name:"AutoSyncFlag"`

	// 金融围拢 ID 。
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// 默认参数模板类型。支持值包括："HIGH_STABILITY" - 高稳定模板，"HIGH_PERFORMANCE" - 高性能模板，默认值是："HIGH_STABILITY"。
	// 备注：如您需使用云数据库 MySQL 默认参数模板，请设置 ParamTemplateType。
	ParamTemplateType *string `json:"ParamTemplateType,omitnil,omitempty" name:"ParamTemplateType"`

	// 告警策略名数组，例如:["policy-uyoee9wg"]，AlarmPolicyList不为空时该参数无效。
	AlarmPolicyIdList []*string `json:"AlarmPolicyIdList,omitnil,omitempty" name:"AlarmPolicyIdList"`

	// 是否只预检此次请求。true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制等。如果检查不通过，则返回对应错误码；如果检查通过，则返回RequestId.默认为false：发送正常请求，通过检查后直接创建实例。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 实例引擎类型，默认为"InnoDB"，支持值包括："InnoDB"，"RocksDB"。
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 指定实例的IP列表。仅支持主实例指定，按实例顺序，不足则按未指定处理。
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// 集群版实例的数据保护空间大小，单位 GB，设置范围1 - 10。
	DataProtectVolume *int64 `json:"DataProtectVolume,omitnil,omitempty" name:"DataProtectVolume"`

	// 集群版节点拓扑配置。
	// 说明：若购买的是集群版实例，此参数为必填，需设置集群版实例的 RW 和 RO 节点拓扑，RO 节点范围是1 - 5个，请至少设置1个 RO 节点。
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// 硬盘类型，单节点（云盘版）或者集群版实例可以指定此参数。CLOUD_SSD 表示 SSD 云硬盘，CLOUD_HSSD 表示增强型 SSD 云硬盘。
	// 说明：单节点（云盘版）、集群版实例硬盘类型所支持的地域略有不同，具体支持情况请参考 [地域和可用区](https://cloud.tencent.com/document/product/236/8458)。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 集群类型:cage——金融围拢，cdc——CDB ON CDC；dedicate——独享集群
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type CreateDBInstanceHourRequest struct {
	*tchttp.BaseRequest
	
	// 实例数量，默认值为 1，最小值 1，最大值为 100。
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 实例内存大小，单位：MB，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的内存规格。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的硬盘范围。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// MySQL 版本，值包括：5.5、5.6、5.7和8.0，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的实例版本。
	// 说明：创建非集群版实例时，请根据需要指定实例版本（推荐5.7或8.0），若此参数不填，则默认值为5.6；若创建的是集群版实例，则此参数仅能指定为5.7或8.0。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 私有网络 ID，请使用 [查询私有网络列表](/document/api/215/15778)。
	// 说明：如果创建的是集群版实例，此参数为必填且为私有网络类型。若此项不填，则系统会选择默认的 VPC。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 私有网络下的子网 ID，如果设置了 UniqVpcId，则 UniqSubnetId 必填，请使用 [查询子网列表](/document/api/215/15784)。
	// 说明：若此项不填，则系统会选择默认 VPC 下的默认子网。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 项目 ID，不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 可用区信息，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的可用区。
	// 说明：若您创建单节点、双节点、三节点实例，此参数为必填项，请指定可用区，若不指定可用区，则系统会自动选择一个可用区（可能不是您希望部署的可用区）；若您创建集群版实例，此参数不填，请通过参数 ClusterTopology 进行读写节点和只读节点的可用区配置。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例 ID，购买只读实例或者灾备实例时必填，该字段表示只读实例或者灾备实例的主实例 ID，请使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口查询云数据库实例 ID。
	MasterInstanceId *string `json:"MasterInstanceId,omitnil,omitempty" name:"MasterInstanceId"`

	// 实例类型，支持值包括：master - 表示主实例，dr - 表示灾备实例，ro - 表示只读实例。
	// 说明：请选择实例类型，不填会默认选择 master。
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 主实例地域信息，购买灾备、RO实例时，该字段必填。
	MasterRegion *string `json:"MasterRegion,omitnil,omitempty" name:"MasterRegion"`

	// 自定义端口，端口支持范围：[1024 - 65535]。
	// 说明：不填则默认为3306。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 设置 root 账号密码，密码规则：8 - 64 个字符，至少包含字母、数字、字符（支持的字符：_+-&=!@#$%^*()）中的两种，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 参数列表，参数格式如 ParamList.0.Name=auto_increment&ParamList.0.Value=1。可通过 [查询默认的可设置参数列表](https://cloud.tencent.com/document/api/236/32662) 查询支持设置的参数。
	// 说明：表名大小写敏感的开启和关闭可通过参数 lower_case_table_names 进行设置，参数值为0表示开启，参数值为1表示关闭，若不设置则此参数默认值为0。若您创建的是 MySQL 8.0 版本的实例，则需要在创建实例时通过设置 lower_case_table_names 参数来开启或关闭表名大小写敏感，创建实例后无法修改参数，即创建后无法修改表名大小写敏感。其他数据库版本的实例支持在创建实例后修改 lower_case_table_names 参数。创建实例时设置表名大小写敏感的 API 调用方法请参见本文中的示例2。
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// 多可用区域，默认为 0，支持值包括：0 - 表示单可用区，1 - 表示多可用区，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 备库 1 的可用区信息。
	// 说明：双节点、三节点实例请指定此参数值，若不指定，则默认为 Zone 的值；集群版实例此参数可不填，请通过参数 ClusterTopology 进行读写节点和只读节点的可用区配置；单节点实例为单可用区，无需指定此参数。
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// 备库 2 的可用区信息，默认为空，购买三节点主实例时可指定该参数。
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// 安全组参数，可使用 [查询项目安全组信息](https://cloud.tencent.com/document/api/236/15850) 接口查询某个项目的安全组详情。
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 只读实例信息。购买只读实例时，该参数必传。
	RoGroup *RoGroup `json:"RoGroup,omitnil,omitempty" name:"RoGroup"`

	// 购买按量计费实例该字段无意义。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 实例名称。一次购买多个实例命名会用后缀数字区分，例instanceName=db，goodsNum=3，实例命名分别为db1，db2，db3。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例标签信息。
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间在48小时内唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 实例隔离类型。支持值包括："UNIVERSAL" - 通用型实例，"EXCLUSIVE" - 独享型实例，"BASIC_V2" - ONTKE 单节点实例，"CLOUD_NATIVE_CLUSTER" - 集群版标准型，"CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - 集群版加强型。不指定则默认为通用型实例。
	// 说明：如果创建的是集群版实例，此参数为必填。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 参数模板 id。
	// 备注：如您使用自定义参数模板 id，可传入自定义参数模板 id；如您计划使用默认参数模板，该参数模板 id 传入 id 无效，需设置 ParamTemplateType。
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 告警策略id数组。腾讯云可观测平台DescribeAlarmPolicy接口返回的OriginId。
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// 实例节点数。对于 RO 和 基础版实例， 该值默认为1。 如果需要购买三节点实例， 请将该值设置为3 或指定 BackupZone 参数。当购买主实例，且未指定该参数和 BackupZone 参数时，该值默认是 2， 即购买两节点实例。
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// 实例cpu核数， 如果不传将根据memory指定的内存值自动填充对应的cpu值。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 是否自动发起灾备同步功能。该参数仅对购买灾备实例生效。 可选值为：0 - 不自动发起灾备同步；1 - 自动发起灾备同步。该值默认为0。
	AutoSyncFlag *int64 `json:"AutoSyncFlag,omitnil,omitempty" name:"AutoSyncFlag"`

	// 金融围拢 ID 。
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// 默认参数模板类型。支持值包括："HIGH_STABILITY" - 高稳定模板，"HIGH_PERFORMANCE" - 高性能模板，默认值是："HIGH_STABILITY"。
	// 备注：如您需使用云数据库 MySQL 默认参数模板，请设置 ParamTemplateType。
	ParamTemplateType *string `json:"ParamTemplateType,omitnil,omitempty" name:"ParamTemplateType"`

	// 告警策略名数组，例如:["policy-uyoee9wg"]，AlarmPolicyList不为空时该参数无效。
	AlarmPolicyIdList []*string `json:"AlarmPolicyIdList,omitnil,omitempty" name:"AlarmPolicyIdList"`

	// 是否只预检此次请求。true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制等。如果检查不通过，则返回对应错误码；如果检查通过，则返回RequestId.默认为false：发送正常请求，通过检查后直接创建实例。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 实例引擎类型，默认为"InnoDB"，支持值包括："InnoDB"，"RocksDB"。
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 指定实例的IP列表。仅支持主实例指定，按实例顺序，不足则按未指定处理。
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// 集群版实例的数据保护空间大小，单位 GB，设置范围1 - 10。
	DataProtectVolume *int64 `json:"DataProtectVolume,omitnil,omitempty" name:"DataProtectVolume"`

	// 集群版节点拓扑配置。
	// 说明：若购买的是集群版实例，此参数为必填，需设置集群版实例的 RW 和 RO 节点拓扑，RO 节点范围是1 - 5个，请至少设置1个 RO 节点。
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// 硬盘类型，单节点（云盘版）或者集群版实例可以指定此参数。CLOUD_SSD 表示 SSD 云硬盘，CLOUD_HSSD 表示增强型 SSD 云硬盘。
	// 说明：单节点（云盘版）、集群版实例硬盘类型所支持的地域略有不同，具体支持情况请参考 [地域和可用区](https://cloud.tencent.com/document/product/236/8458)。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 集群类型:cage——金融围拢，cdc——CDB ON CDC；dedicate——独享集群
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
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
	delete(f, "DataProtectVolume")
	delete(f, "ClusterTopology")
	delete(f, "DiskType")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceHourRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceHourResponseParams struct {
	// 短订单 ID。
	DealIds []*string `json:"DealIds,omitnil,omitempty" name:"DealIds"`

	// 实例 ID 列表。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的硬盘范围。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 实例时长，单位：月，可选值包括 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 实例数量，默认值为1, 最小值1，最大值为100。
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 可用区信息，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的可用区。
	// 说明：若您创建单节点、双节点、三节点实例，此参数为必填项，请指定可用区，若不指定可用区，则系统会自动选择一个可用区（可能不是您希望部署的可用区）；若您创建集群版实例，此参数不填，请通过参数 ClusterTopology 进行读写节点和只读节点的可用区配置。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 私有网络 ID，请使用 [查询私有网络列表](/document/api/215/15778)。
	// 说明：如果创建的是集群版实例，此参数为必填且为私有网络类型。若此项不填，则系统会选择默认的 VPC。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 私有网络下的子网 ID，如果设置了 UniqVpcId，则 UniqSubnetId 必填，请使用 [查询子网列表](/document/api/215/15784)。
	// 说明：若此项不填，则系统会选择默认 VPC 下的默认子网。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 项目 ID，不填为默认项目。购买只读实例和灾备实例时，项目 ID 默认和主实例保持一致。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 自定义端口，端口支持范围：[ 1024-65535 ]。
	// 说明：若此项不填，则默认为3306。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 实例类型，支持值包括：master - 表示主实例，dr - 表示灾备实例，ro - 表示只读实例。
	// 说明：请选择实例类型，不填会默认选择 master。
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 实例 ID，购买只读实例或灾备实例时必填，该字段表示只读实例或灾备实例的主实例 ID，请使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口查询云数据库实例 ID。
	MasterInstanceId *string `json:"MasterInstanceId,omitnil,omitempty" name:"MasterInstanceId"`

	// MySQL 版本，值包括：5.5、5.6、5.7和8.0，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的实例版本。
	// 说明：创建非集群版实例时，请根据需要指定实例版本（推荐5.7或8.0），若此参数不填，则默认值为5.6；若创建的是集群版实例，则此参数仅能指定为5.7或8.0。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 设置 root 账号密码，密码规则：8 - 64 个字符，至少包含字母、数字、字符（支持的字符：_+-&=!@#$%^*()）中的两种，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制。
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// 多可用区域，默认为 0，支持值包括：0 - 表示单可用区，1 - 表示多可用区。
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 备库 1 的可用区信息。
	// 说明：双节点、三节点实例请指定此参数值，若不指定，则默认为 Zone 的值；集群版实例此参数可不填，请通过参数 ClusterTopology 进行读写节点和只读节点的可用区配置；单节点实例为单可用区，无需指定此参数。
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// 参数列表，参数格式如 ParamList.0.Name=auto_increment&ParamList.0.Value=1。可通过 [查询默认的可设置参数列表](https://cloud.tencent.com/document/api/236/32662) 查询支持设置的参数。
	// 说明：表名大小写敏感的开启和关闭可通过参数 lower_case_table_names 进行设置，参数值为0表示开启，参数值为1表示关闭，若不设置则此参数默认值为0。若您创建的是 MySQL 8.0 版本的实例，则需要在创建实例时通过设置 lower_case_table_names 参数来开启或关闭表名大小写敏感，创建实例后无法修改参数，即创建后无法修改表名大小写敏感。其他数据库版本的实例支持在创建实例后修改 lower_case_table_names 参数。创建实例时设置表名大小写敏感的 API 调用方法请参见本文中的示例3。
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// 备库 2 的可用区信息，默认为空，购买三节点主实例时可指定该参数。
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// 自动续费标记，可选值为：0 - 不自动续费；1 - 自动续费。默认为0。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 主实例地域信息，购买灾备、RO实例时，该字段必填。
	MasterRegion *string `json:"MasterRegion,omitnil,omitempty" name:"MasterRegion"`

	// 安全组参数，可使用 [查询项目安全组信息](https://cloud.tencent.com/document/api/236/15850) 接口查询某个项目的安全组详情。
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 只读实例参数。购买只读实例时，该参数必传。
	RoGroup *RoGroup `json:"RoGroup,omitnil,omitempty" name:"RoGroup"`

	// 实例名称。一次购买多个实例命名会用后缀数字区分，例instnaceName=db，goodsNum=3，实例命名分别为db1，db2，db3。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例标签信息。
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间在48小时内唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 实例隔离类型。支持值包括："UNIVERSAL" - 通用型实例，"EXCLUSIVE" - 独享型实例，"BASIC_V2" - ONTKE 单节点实例，"CLOUD_NATIVE_CLUSTER" - 集群版标准型，"CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - 集群版加强型。不指定则默认为通用型实例。
	// 说明：如果创建的是集群版实例，此参数为必填。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 参数模板 id。
	// 备注：如您使用自定义参数模板 id，可传入自定义参数模板 id；如您计划使用默认参数模板，该参数模板 id 传入 id 无效，需设置 ParamTemplateType。
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 告警策略id数组。腾讯云可观测平台DescribeAlarmPolicy接口返回的OriginId。
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// 实例节点数。对于 RO 和 基础版实例， 该值默认为1。 如果需要购买三节点实例， 请将该值设置为3 或指定 BackupZone 参数。当购买主实例，且未指定该参数和 BackupZone 参数时，该值默认是 2， 即购买两节点实例。
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// 实例cpu核数， 如果不传将根据memory指定的内存值自动填充对应的cpu值。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 是否自动发起灾备同步功能。该参数仅对购买灾备实例生效。 可选值为：0 - 不自动发起灾备同步；1 - 自动发起灾备同步。该值默认为0。
	AutoSyncFlag *int64 `json:"AutoSyncFlag,omitnil,omitempty" name:"AutoSyncFlag"`

	// 金融围拢 ID。
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// 默认参数模板类型。支持值包括："HIGH_STABILITY" - 高稳定模板，"HIGH_PERFORMANCE" - 高性能模板。
	// 备注：如您需使用云数据库 MySQL 默认参数模板，请设置 ParamTemplateType。
	ParamTemplateType *string `json:"ParamTemplateType,omitnil,omitempty" name:"ParamTemplateType"`

	// 告警策略名数组，例如:["policy-uyoee9wg"]，AlarmPolicyList不为空时该参数无效。
	AlarmPolicyIdList []*string `json:"AlarmPolicyIdList,omitnil,omitempty" name:"AlarmPolicyIdList"`

	// 是否只预检此次请求。true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制等。如果检查不通过，则返回对应错误码；如果检查通过，则返回 RequestId。false：发送正常请求，通过检查后直接创建实例。
	// 默认为 false。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 实例引擎类型，默认为"InnoDB"，支持值包括："InnoDB"，"RocksDB"。
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 指定实例的IP列表。仅支持主实例指定，按实例顺序，不足则按未指定处理。
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// 集群版实例的数据保护空间大小，单位 GB，设置范围1 - 10。
	DataProtectVolume *int64 `json:"DataProtectVolume,omitnil,omitempty" name:"DataProtectVolume"`

	// 集群版节点拓扑配置。
	// 说明：若购买的是集群版实例，此参数为必填，需设置集群版实例的 RW 和 RO 节点拓扑，RO 节点范围是1 - 5个，请至少设置1个 RO 节点。
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// 磁盘类型，单节点（云盘版）或者集群版实例可以指定此参数。CLOUD_SSD 表示 SSD 云硬盘，CLOUD_HSSD 表示增强型 SSD 云硬盘。
	// 说明：单节点（云盘版）、集群版实例硬盘类型所支持的地域略有不同，具体支持情况请参考 [地域和可用区](https://cloud.tencent.com/document/product/236/8458)。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`
}

type CreateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例内存大小，单位：MB，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的内存规格。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的硬盘范围。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 实例时长，单位：月，可选值包括 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 实例数量，默认值为1, 最小值1，最大值为100。
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 可用区信息，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的可用区。
	// 说明：若您创建单节点、双节点、三节点实例，此参数为必填项，请指定可用区，若不指定可用区，则系统会自动选择一个可用区（可能不是您希望部署的可用区）；若您创建集群版实例，此参数不填，请通过参数 ClusterTopology 进行读写节点和只读节点的可用区配置。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 私有网络 ID，请使用 [查询私有网络列表](/document/api/215/15778)。
	// 说明：如果创建的是集群版实例，此参数为必填且为私有网络类型。若此项不填，则系统会选择默认的 VPC。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 私有网络下的子网 ID，如果设置了 UniqVpcId，则 UniqSubnetId 必填，请使用 [查询子网列表](/document/api/215/15784)。
	// 说明：若此项不填，则系统会选择默认 VPC 下的默认子网。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 项目 ID，不填为默认项目。购买只读实例和灾备实例时，项目 ID 默认和主实例保持一致。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 自定义端口，端口支持范围：[ 1024-65535 ]。
	// 说明：若此项不填，则默认为3306。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 实例类型，支持值包括：master - 表示主实例，dr - 表示灾备实例，ro - 表示只读实例。
	// 说明：请选择实例类型，不填会默认选择 master。
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 实例 ID，购买只读实例或灾备实例时必填，该字段表示只读实例或灾备实例的主实例 ID，请使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口查询云数据库实例 ID。
	MasterInstanceId *string `json:"MasterInstanceId,omitnil,omitempty" name:"MasterInstanceId"`

	// MySQL 版本，值包括：5.5、5.6、5.7和8.0，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口获取可创建的实例版本。
	// 说明：创建非集群版实例时，请根据需要指定实例版本（推荐5.7或8.0），若此参数不填，则默认值为5.6；若创建的是集群版实例，则此参数仅能指定为5.7或8.0。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 设置 root 账号密码，密码规则：8 - 64 个字符，至少包含字母、数字、字符（支持的字符：_+-&=!@#$%^*()）中的两种，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制。
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// 多可用区域，默认为 0，支持值包括：0 - 表示单可用区，1 - 表示多可用区。
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 备库 1 的可用区信息。
	// 说明：双节点、三节点实例请指定此参数值，若不指定，则默认为 Zone 的值；集群版实例此参数可不填，请通过参数 ClusterTopology 进行读写节点和只读节点的可用区配置；单节点实例为单可用区，无需指定此参数。
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// 参数列表，参数格式如 ParamList.0.Name=auto_increment&ParamList.0.Value=1。可通过 [查询默认的可设置参数列表](https://cloud.tencent.com/document/api/236/32662) 查询支持设置的参数。
	// 说明：表名大小写敏感的开启和关闭可通过参数 lower_case_table_names 进行设置，参数值为0表示开启，参数值为1表示关闭，若不设置则此参数默认值为0。若您创建的是 MySQL 8.0 版本的实例，则需要在创建实例时通过设置 lower_case_table_names 参数来开启或关闭表名大小写敏感，创建实例后无法修改参数，即创建后无法修改表名大小写敏感。其他数据库版本的实例支持在创建实例后修改 lower_case_table_names 参数。创建实例时设置表名大小写敏感的 API 调用方法请参见本文中的示例3。
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// 备库 2 的可用区信息，默认为空，购买三节点主实例时可指定该参数。
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// 自动续费标记，可选值为：0 - 不自动续费；1 - 自动续费。默认为0。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 主实例地域信息，购买灾备、RO实例时，该字段必填。
	MasterRegion *string `json:"MasterRegion,omitnil,omitempty" name:"MasterRegion"`

	// 安全组参数，可使用 [查询项目安全组信息](https://cloud.tencent.com/document/api/236/15850) 接口查询某个项目的安全组详情。
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 只读实例参数。购买只读实例时，该参数必传。
	RoGroup *RoGroup `json:"RoGroup,omitnil,omitempty" name:"RoGroup"`

	// 实例名称。一次购买多个实例命名会用后缀数字区分，例instnaceName=db，goodsNum=3，实例命名分别为db1，db2，db3。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例标签信息。
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间在48小时内唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 实例隔离类型。支持值包括："UNIVERSAL" - 通用型实例，"EXCLUSIVE" - 独享型实例，"BASIC_V2" - ONTKE 单节点实例，"CLOUD_NATIVE_CLUSTER" - 集群版标准型，"CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - 集群版加强型。不指定则默认为通用型实例。
	// 说明：如果创建的是集群版实例，此参数为必填。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 参数模板 id。
	// 备注：如您使用自定义参数模板 id，可传入自定义参数模板 id；如您计划使用默认参数模板，该参数模板 id 传入 id 无效，需设置 ParamTemplateType。
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 告警策略id数组。腾讯云可观测平台DescribeAlarmPolicy接口返回的OriginId。
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// 实例节点数。对于 RO 和 基础版实例， 该值默认为1。 如果需要购买三节点实例， 请将该值设置为3 或指定 BackupZone 参数。当购买主实例，且未指定该参数和 BackupZone 参数时，该值默认是 2， 即购买两节点实例。
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// 实例cpu核数， 如果不传将根据memory指定的内存值自动填充对应的cpu值。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 是否自动发起灾备同步功能。该参数仅对购买灾备实例生效。 可选值为：0 - 不自动发起灾备同步；1 - 自动发起灾备同步。该值默认为0。
	AutoSyncFlag *int64 `json:"AutoSyncFlag,omitnil,omitempty" name:"AutoSyncFlag"`

	// 金融围拢 ID。
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// 默认参数模板类型。支持值包括："HIGH_STABILITY" - 高稳定模板，"HIGH_PERFORMANCE" - 高性能模板。
	// 备注：如您需使用云数据库 MySQL 默认参数模板，请设置 ParamTemplateType。
	ParamTemplateType *string `json:"ParamTemplateType,omitnil,omitempty" name:"ParamTemplateType"`

	// 告警策略名数组，例如:["policy-uyoee9wg"]，AlarmPolicyList不为空时该参数无效。
	AlarmPolicyIdList []*string `json:"AlarmPolicyIdList,omitnil,omitempty" name:"AlarmPolicyIdList"`

	// 是否只预检此次请求。true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制等。如果检查不通过，则返回对应错误码；如果检查通过，则返回 RequestId。false：发送正常请求，通过检查后直接创建实例。
	// 默认为 false。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 实例引擎类型，默认为"InnoDB"，支持值包括："InnoDB"，"RocksDB"。
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 指定实例的IP列表。仅支持主实例指定，按实例顺序，不足则按未指定处理。
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// 集群版实例的数据保护空间大小，单位 GB，设置范围1 - 10。
	DataProtectVolume *int64 `json:"DataProtectVolume,omitnil,omitempty" name:"DataProtectVolume"`

	// 集群版节点拓扑配置。
	// 说明：若购买的是集群版实例，此参数为必填，需设置集群版实例的 RW 和 RO 节点拓扑，RO 节点范围是1 - 5个，请至少设置1个 RO 节点。
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// 磁盘类型，单节点（云盘版）或者集群版实例可以指定此参数。CLOUD_SSD 表示 SSD 云硬盘，CLOUD_HSSD 表示增强型 SSD 云硬盘。
	// 说明：单节点（云盘版）、集群版实例硬盘类型所支持的地域略有不同，具体支持情况请参考 [地域和可用区](https://cloud.tencent.com/document/product/236/8458)。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`
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
	delete(f, "DataProtectVolume")
	delete(f, "ClusterTopology")
	delete(f, "DiskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceResponseParams struct {
	// 计费子订单 ID。
	DealIds []*string `json:"DealIds,omitnil,omitempty" name:"DealIds"`

	// 实例 ID 列表。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateDatabaseRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名称，长度不超过64。
	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`

	// 字符集，可选值：utf8，gbk，latin1，utf8mb4。
	CharacterSetName *string `json:"CharacterSetName,omitnil,omitempty" name:"CharacterSetName"`
}

type CreateDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名称，长度不超过64。
	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`

	// 字符集，可选值：utf8，gbk，latin1，utf8mb4。
	CharacterSetName *string `json:"CharacterSetName,omitnil,omitempty" name:"CharacterSetName"`
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
	delete(f, "InstanceId")
	delete(f, "DBName")
	delete(f, "CharacterSetName")
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
type CreateDeployGroupRequestParams struct {
	// 置放群组名称，最长不能超过60个字符。
	DeployGroupName *string `json:"DeployGroupName,omitnil,omitempty" name:"DeployGroupName"`

	// 置放群组描述，最长不能超过200个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 置放群组的亲和性策略，目前仅支持取值为1，策略1表示同台物理机上限制实例的个数。
	Affinity []*int64 `json:"Affinity,omitnil,omitempty" name:"Affinity"`

	// 置放群组亲和性策略1中同台物理机上实例的限制个数。
	LimitNum *int64 `json:"LimitNum,omitnil,omitempty" name:"LimitNum"`

	// 置放群组机型属性，可选参数：SH12+SH02、TS85。
	DevClass []*string `json:"DevClass,omitnil,omitempty" name:"DevClass"`
}

type CreateDeployGroupRequest struct {
	*tchttp.BaseRequest
	
	// 置放群组名称，最长不能超过60个字符。
	DeployGroupName *string `json:"DeployGroupName,omitnil,omitempty" name:"DeployGroupName"`

	// 置放群组描述，最长不能超过200个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 置放群组的亲和性策略，目前仅支持取值为1，策略1表示同台物理机上限制实例的个数。
	Affinity []*int64 `json:"Affinity,omitnil,omitempty" name:"Affinity"`

	// 置放群组亲和性策略1中同台物理机上实例的限制个数。
	LimitNum *int64 `json:"LimitNum,omitnil,omitempty" name:"LimitNum"`

	// 置放群组机型属性，可选参数：SH12+SH02、TS85。
	DevClass []*string `json:"DevClass,omitnil,omitempty" name:"DevClass"`
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
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 参数模板名称。支持输入最大60个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// MySQL 版本号。可选值：5.6、5.7、8.0。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 源参数模板 ID。可通过 [DescribeParamTemplates](https://cloud.tencent.com/document/api/236/32659) 接口获取。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 参数列表。
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// 默认参数模板类型。支持值包括："HIGH_STABILITY" - 高稳定模板，"HIGH_PERFORMANCE" - 高性能模板。
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// 实例引擎类型，默认为"InnoDB"，支持值包括："InnoDB"，"RocksDB"。
	// 说明：数据库版本 MySQL 5.7、MySQL 8.0才支持 RocksDB。
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
}

type CreateParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板名称。支持输入最大60个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// MySQL 版本号。可选值：5.6、5.7、8.0。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 源参数模板 ID。可通过 [DescribeParamTemplates](https://cloud.tencent.com/document/api/236/32659) 接口获取。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 参数列表。
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// 默认参数模板类型。支持值包括："HIGH_STABILITY" - 高稳定模板，"HIGH_PERFORMANCE" - 高性能模板。
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// 实例引擎类型，默认为"InnoDB"，支持值包括："InnoDB"，"RocksDB"。
	// 说明：数据库版本 MySQL 5.7、MySQL 8.0才支持 RocksDB。
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
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
type CreateRoInstanceIpRequestParams struct {
	// 只读实例ID，格式如：cdbro-3i70uj0k，与云数据库控制台页面中显示的只读实例ID相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 子网描述符，例如：subnet-1typ0s7d。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// vpc描述符，例如：vpc-a23yt67j,如果传了该字段则UniqSubnetId必传
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`
}

type CreateRoInstanceIpRequest struct {
	*tchttp.BaseRequest
	
	// 只读实例ID，格式如：cdbro-3i70uj0k，与云数据库控制台页面中显示的只读实例ID相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 子网描述符，例如：subnet-1typ0s7d。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// vpc描述符，例如：vpc-a23yt67j,如果传了该字段则UniqSubnetId必传
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`
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
	RoVpcId *int64 `json:"RoVpcId,omitnil,omitempty" name:"RoVpcId"`

	// 只读实例的子网ID。
	RoSubnetId *int64 `json:"RoSubnetId,omitnil,omitempty" name:"RoSubnetId"`

	// 只读实例的内网IP地址。
	RoVip *string `json:"RoVip,omitnil,omitempty" name:"RoVip"`

	// 只读实例的内网端口号。
	RoVport *int64 `json:"RoVport,omitnil,omitempty" name:"RoVport"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CreateRotationPasswordRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 当前需开启密码轮转的账号信息，包含账户名与主机名
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

type CreateRotationPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 当前需开启密码轮转的账号信息，包含账户名与主机名
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

func (r *CreateRotationPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRotationPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRotationPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRotationPasswordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRotationPasswordResponse struct {
	*tchttp.BaseResponse
	Response *CreateRotationPasswordResponseParams `json:"Response"`
}

func (r *CreateRotationPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRotationPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomConfig struct {
	// 设备
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 设备类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 内存，单位为MB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 核数
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`
}

type DBSwitchInfo struct {
	// 切换时间，格式为：2017-09-03 01:34:31
	SwitchTime *string `json:"SwitchTime,omitnil,omitempty" name:"SwitchTime"`

	// 切换类型，可能的返回值为：TRANSFER - 数据迁移；MASTER2SLAVE - 主备切换；RECOVERY - 主从恢复
	SwitchType *string `json:"SwitchType,omitnil,omitempty" name:"SwitchType"`
}

type DatabasePrivilege struct {
	// 权限信息
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`

	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`
}

type DatabasesWithCharacterLists struct {
	// 数据库名
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 字符集类型
	CharacterSet *string `json:"CharacterSet,omitnil,omitempty" name:"CharacterSet"`
}

// Predefined struct for user
type DeleteAccountsRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 云数据库账号。
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

type DeleteAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 云数据库账号。
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

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
	// 审计日志文件名称。可通过 [DescribeAuditLogFiles](https://cloud.tencent.com/document/api/236/45454) 接口获取。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteAuditLogFileRequest struct {
	*tchttp.BaseRequest
	
	// 审计日志文件名称。可通过 [DescribeAuditLogFiles](https://cloud.tencent.com/document/api/236/45454) 接口获取。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
type DeleteAuditPolicyRequestParams struct {
	// 审计策略 ID。
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteAuditPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 审计策略 ID。
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DeleteAuditRuleRequest struct {
	*tchttp.BaseRequest
	
	// 审计规则 ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteAuditRuleTemplatesRequestParams struct {
	// 审计规则模板 ID，可通过 [DescribeAuditRuleTemplates](https://cloud.tencent.com/document/api/236/101811) 接口获取，单次允许最多删除5个规则模板。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
}

type DeleteAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 审计规则模板 ID，可通过 [DescribeAuditRuleTemplates](https://cloud.tencent.com/document/api/236/101811) 接口获取，单次允许最多删除5个规则模板。
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
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备份任务 ID。您可通过 [查询数据备份文件列表](https://cloud.tencent.com/document/api/236/15842)  来获取目标备份任务 ID。
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`
}

type DeleteBackupRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备份任务 ID。您可通过 [查询数据备份文件列表](https://cloud.tencent.com/document/api/236/15842)  来获取目标备份任务 ID。
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`
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
type DeleteDatabaseRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名称，长度不超过64。
	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`
}

type DeleteDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名称，长度不超过64。
	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`
}

func (r *DeleteDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DBName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDatabaseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDatabaseResponseParams `json:"Response"`
}

func (r *DeleteDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeployGroupsRequestParams struct {
	// 要删除的置放群组 ID 列表。
	DeployGroupIds []*string `json:"DeployGroupIds,omitnil,omitempty" name:"DeployGroupIds"`
}

type DeleteDeployGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的置放群组 ID 列表。
	DeployGroupIds []*string `json:"DeployGroupIds,omitnil,omitempty" name:"DeployGroupIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 参数模板 ID。可通过 [DescribeParamTemplates](https://cloud.tencent.com/document/api/236/32659) 接口获取。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板 ID。可通过 [DescribeParamTemplates](https://cloud.tencent.com/document/api/236/32659) 接口获取。
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

// Predefined struct for user
type DeleteRotationPasswordRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 关闭密码轮转的实例账户名,例如root
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 关闭密码轮转的实例账户域名，例如%
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 关闭密码轮转后实例账户的最新密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 传入不为空则对密码进行了加密处理
	EncryptMethod *string `json:"EncryptMethod,omitnil,omitempty" name:"EncryptMethod"`
}

type DeleteRotationPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 关闭密码轮转的实例账户名,例如root
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 关闭密码轮转的实例账户域名，例如%
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 关闭密码轮转后实例账户的最新密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 传入不为空则对密码进行了加密处理
	EncryptMethod *string `json:"EncryptMethod,omitnil,omitempty" name:"EncryptMethod"`
}

func (r *DeleteRotationPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRotationPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "Host")
	delete(f, "Password")
	delete(f, "EncryptMethod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRotationPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRotationPasswordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRotationPasswordResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRotationPasswordResponseParams `json:"Response"`
}

func (r *DeleteRotationPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRotationPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTimeWindowRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteTimeWindowRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// 置放群组名称。
	DeployGroupName *string `json:"DeployGroupName,omitnil,omitempty" name:"DeployGroupName"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 置放群组实例配额，表示一个置放群组中可容纳的最大实例数目。
	Quota *int64 `json:"Quota,omitnil,omitempty" name:"Quota"`

	// 置放群组亲和性策略，目前仅支持策略1，即在物理机纬度打散实例的分布。
	Affinity *string `json:"Affinity,omitnil,omitempty" name:"Affinity"`

	// 置放群组亲和性策略1中，同台物理机上同个置放群组实例的限制个数。
	LimitNum *int64 `json:"LimitNum,omitnil,omitempty" name:"LimitNum"`

	// 置放群组详细信息。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 置放群组物理机型属性。
	DevClass *string `json:"DevClass,omitnil,omitempty" name:"DevClass"`
}

// Predefined struct for user
type DescribeAccountPrivilegesRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库的账号名称。可通过 [DescribeAccounts](https://cloud.tencent.com/document/api/236/17499) 接口获取。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 数据库的账号域名。可通过 [DescribeAccounts](https://cloud.tencent.com/document/api/236/17499) 接口获取。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库的账号名称。可通过 [DescribeAccounts](https://cloud.tencent.com/document/api/236/17499) 接口获取。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 数据库的账号域名。可通过 [DescribeAccounts](https://cloud.tencent.com/document/api/236/17499) 接口获取。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
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
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// 数据库权限数组。
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// 数据库中的表权限数组。
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`

	// 数据库表中的列权限数组。
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitnil,omitempty" name:"ColumnPrivileges"`

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
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 记录偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为20，最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 匹配账号名的正则表达式，规则同 MySQL 官网。
	AccountRegexp *string `json:"AccountRegexp,omitnil,omitempty" name:"AccountRegexp"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 记录偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为20，最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 匹配账号名的正则表达式，规则同 MySQL 官网。
	AccountRegexp *string `json:"AccountRegexp,omitnil,omitempty" name:"AccountRegexp"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合查询条件的账号详细信息。
	Items []*AccountInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 用户可设置实例最大连接数。
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`

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
type DescribeAsyncRequestInfoRequestParams struct {
	// 异步任务的请求 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`
}

type DescribeAsyncRequestInfoRequest struct {
	*tchttp.BaseRequest
	
	// 异步任务的请求 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`
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
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务执行信息描述。
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeAuditConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 审计日志存储类型。目前支持的值包括："storage" - 存储型。
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 是否正在关闭审计。目前支持的值包括："false"-否，"true"-是
	IsClosing *string `json:"IsClosing,omitnil,omitempty" name:"IsClosing"`

	// 审计服务开通时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAuditInstanceListRequestParams struct {
	// 实例审计开启的状态。1-已开启审计；0-未开启审计。
	AuditSwitch *int64 `json:"AuditSwitch,omitnil,omitempty" name:"AuditSwitch"`

	// 查询实例列表的过滤条件。
	Filters []*AuditInstanceFilters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 实例的审计规则模式。1-规则审计；0-全审计。
	AuditMode *int64 `json:"AuditMode,omitnil,omitempty" name:"AuditMode"`

	// 单次请求返回的数量。默认值为30，最大值为 20000。
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

	// 单次请求返回的数量。默认值为30，最大值为 20000。
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
	Items []*InstanceDbAuditStatus `json:"Items,omitnil,omitempty" name:"Items"`

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
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页大小参数。默认值为20，最小值为1，最大值为300。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 审计日志文件名。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

type DescribeAuditLogFilesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页大小参数。默认值为20，最小值为1，最大值为300。
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
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间(建议开始到结束时间区间最大7天)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间(建议开始到结束时间区间最大7天）。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页参数，单次返回的数据条数。默认值为100，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 日志偏移量，最多支持偏移查询65535条日志。可填写范围：0 - 65535。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序方式。支持值包括："ASC" - 升序，"DESC" - 降序，默认降序排序。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段。支持值包括(默认按照时间戳排序)：
	// "timestamp" - 时间戳；
	// "affectRows" - 影响行数；
	// "execTime" - 执行时间。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 过滤条件。多个值之前是且的关系。
	LogFilter []*InstanceAuditLogFilters `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`
}

type DescribeAuditLogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间(建议开始到结束时间区间最大7天)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间(建议开始到结束时间区间最大7天）。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页参数，单次返回的数据条数。默认值为100，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 日志偏移量，最多支持偏移查询65535条日志。可填写范围：0 - 65535。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序方式。支持值包括："ASC" - 升序，"DESC" - 降序，默认降序排序。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段。支持值包括(默认按照时间戳排序)：
	// "timestamp" - 时间戳；
	// "affectRows" - 影响行数；
	// "execTime" - 执行时间。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 过滤条件。多个值之前是且的关系。
	LogFilter []*InstanceAuditLogFilters `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`
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
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "OrderBy")
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
type DescribeAuditPoliciesRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计策略 ID。
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 审计策略名称。支持按审计策略名称进行模糊匹配查询。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 分页大小参数。默认值为 20，最小值为 1，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 审计规则 ID。可使用该审计规则 ID 查询到其关联的审计策略。
	// 注意，参数 RuleId，InstanceId，PolicyId，PolicyName 必须至少传一个。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type DescribeAuditPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计策略 ID。
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 审计策略名称。支持按审计策略名称进行模糊匹配查询。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 分页大小参数。默认值为 20，最小值为 1，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 审计规则 ID。可使用该审计规则 ID 查询到其关联的审计策略。
	// 注意，参数 RuleId，InstanceId，PolicyId，PolicyName 必须至少传一个。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 审计策略详情。
	Items []*AuditPolicy `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAuditRuleTemplateModifyHistoryRequestParams struct {
	// 审计规则模板ID,可通过[DescribeAuditRuleTemplates](https://cloud.tencent.com/document/api/236/101811)接口获取。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// 查询范围的开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询范围的结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 返回条数,默认值-20，最大值-1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序方式，DESC-按修改时间倒排，ASC-正序，默认：DESC。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeAuditRuleTemplateModifyHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 审计规则模板ID,可通过[DescribeAuditRuleTemplates](https://cloud.tencent.com/document/api/236/101811)接口获取。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// 查询范围的开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询范围的结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 返回条数,默认值-20，最大值-1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序方式，DESC-按修改时间倒排，ASC-正序，默认：DESC。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *DescribeAuditRuleTemplateModifyHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRuleTemplateModifyHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleTemplateIds")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditRuleTemplateModifyHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRuleTemplateModifyHistoryResponseParams struct {
	// 总的条数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 变更详情。
	Items []*RuleTemplateRecordInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditRuleTemplateModifyHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditRuleTemplateModifyHistoryResponseParams `json:"Response"`
}

func (r *DescribeAuditRuleTemplateModifyHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRuleTemplateModifyHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRuleTemplatesRequestParams struct {
	// 规则模板ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// 规则模板名称。
	RuleTemplateNames []*string `json:"RuleTemplateNames,omitnil,omitempty" name:"RuleTemplateNames"`

	// 单次请求返回的数量。默认值20，最大值为1000。
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

	// 规则模板名称。
	RuleTemplateNames []*string `json:"RuleTemplateNames,omitnil,omitempty" name:"RuleTemplateNames"`

	// 单次请求返回的数量。默认值20，最大值为1000。
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
type DescribeAuditRulesRequestParams struct {
	// 审计规则 ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 审计规则名称。支持按审计规则名称进行模糊匹配查询。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 分页大小参数。默认值为 20，最小值为 1，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量。默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeAuditRulesRequest struct {
	*tchttp.BaseRequest
	
	// 审计规则 ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 审计规则名称。支持按审计规则名称进行模糊匹配查询。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 分页大小参数。默认值为 20，最小值为 1，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量。默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 审计规则详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*AuditRule `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	//
	// Deprecated: StartTimeMin is deprecated.
	StartTimeMin *int64 `json:"StartTimeMin,omitnil,omitempty" name:"StartTimeMin"`

	// 自动备份开始的最晚时间点，单位为时刻。例如，6 - 凌晨 6:00。（该字段已废弃，建议使用 BackupTimeWindow 字段）
	//
	// Deprecated: StartTimeMax is deprecated.
	StartTimeMax *int64 `json:"StartTimeMax,omitnil,omitempty" name:"StartTimeMax"`

	// 备份文件保留时间，单位为天。
	BackupExpireDays *int64 `json:"BackupExpireDays,omitnil,omitempty" name:"BackupExpireDays"`

	// 备份方式，可能的值为：physical - 物理备份，logical - 逻辑备份。
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// Binlog 文件保留时间，单位为天。
	BinlogExpireDays *int64 `json:"BinlogExpireDays,omitnil,omitempty" name:"BinlogExpireDays"`

	// 实例自动备份的时间窗。
	BackupTimeWindow *CommonTimeWindow `json:"BackupTimeWindow,omitnil,omitempty" name:"BackupTimeWindow"`

	// 定期保留开关，off - 不开启定期保留策略，on - 开启定期保留策略，默认为off
	EnableBackupPeriodSave *string `json:"EnableBackupPeriodSave,omitnil,omitempty" name:"EnableBackupPeriodSave"`

	// 定期保留最长天数，最小值：90，最大值：3650，默认值：1080
	BackupPeriodSaveDays *int64 `json:"BackupPeriodSaveDays,omitnil,omitempty" name:"BackupPeriodSaveDays"`

	// 定期保留策略周期，可取值：weekly - 周，monthly - 月， quarterly - 季度，yearly - 年，默认为monthly
	BackupPeriodSaveInterval *string `json:"BackupPeriodSaveInterval,omitnil,omitempty" name:"BackupPeriodSaveInterval"`

	// 定期保留的备份数量，最小值为1，最大值不超过定期保留策略周期内常规备份个数，默认值为1
	BackupPeriodSaveCount *int64 `json:"BackupPeriodSaveCount,omitnil,omitempty" name:"BackupPeriodSaveCount"`

	// 定期保留策略周期起始日期，格式：YYYY-MM-dd HH:mm:ss
	StartBackupPeriodSaveDate *string `json:"StartBackupPeriodSaveDate,omitnil,omitempty" name:"StartBackupPeriodSaveDate"`

	// 是否开启数据备份归档策略，off-关闭，on-打开，默认为off
	EnableBackupArchive *string `json:"EnableBackupArchive,omitnil,omitempty" name:"EnableBackupArchive"`

	// 数据备份归档起始天数，数据备份达到归档起始天数时进行归档，最小为180天，不得大于数据备份保留天数
	BackupArchiveDays *int64 `json:"BackupArchiveDays,omitnil,omitempty" name:"BackupArchiveDays"`

	// 是否开启日志备份归档策略，off-关闭，on-打开，默认为off
	EnableBinlogArchive *string `json:"EnableBinlogArchive,omitnil,omitempty" name:"EnableBinlogArchive"`

	// 日志备份归档起始天数，日志备份达到归档起始天数时进行归档，最小为180天，不得大于日志备份保留天数
	BinlogArchiveDays *int64 `json:"BinlogArchiveDays,omitnil,omitempty" name:"BinlogArchiveDays"`

	// 是否开启数据备份标准存储策略，off-关闭，on-打开，默认为off
	EnableBackupStandby *string `json:"EnableBackupStandby,omitnil,omitempty" name:"EnableBackupStandby"`

	// 数据备份标准存储起始天数，数据备份达到标准存储起始天数时进行转换，最小为30天，不得大于数据备份保留天数。如果开启备份归档，不得大于等于备份归档天数
	BackupStandbyDays *int64 `json:"BackupStandbyDays,omitnil,omitempty" name:"BackupStandbyDays"`

	// 是否开启日志备份标准存储策略，off-关闭，on-打开，默认为off
	EnableBinlogStandby *string `json:"EnableBinlogStandby,omitnil,omitempty" name:"EnableBinlogStandby"`

	// 日志备份标准存储起始天数，日志备份达到标准存储起始天数时进行转换，最小为30天，不得大于日志备份保留天数。如果开启备份归档，不得大于等于备份归档天数
	BinlogStandbyDays *int64 `json:"BinlogStandbyDays,omitnil,omitempty" name:"BinlogStandbyDays"`

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
type DescribeBackupDecryptionKeyRequestParams struct {
	// 实例 ID，格式如：cdb-fybaegd8。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例的备份 ID，可通过 [DescribeBackups](https://cloud.tencent.com/document/api/236/15842) 接口查询备份的 ID。
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 备份类型。data-数据备份，binlog-日志备份，默认为 data。
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`
}

type DescribeBackupDecryptionKeyRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-fybaegd8。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例的备份 ID，可通过 [DescribeBackups](https://cloud.tencent.com/document/api/236/15842) 接口查询备份的 ID。
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 备份类型。data-数据备份，binlog-日志备份，默认为 data。
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`
}

func (r *DescribeBackupDecryptionKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDecryptionKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupId")
	delete(f, "BackupType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDecryptionKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDecryptionKeyResponseParams struct {
	// 备份文件解密密钥。
	DecryptionKey *string `json:"DecryptionKey,omitnil,omitempty" name:"DecryptionKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupDecryptionKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDecryptionKeyResponseParams `json:"Response"`
}

func (r *DescribeBackupDecryptionKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDecryptionKeyResponse) FromJsonString(s string) error {
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
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// 该参数仅支持 In， 表示 LimitVpc 指定的vpc可以下载。
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// In: 指定的ip可以下载； NotIn: 指定的ip不可以下载。
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// 限制下载的vpc设置。
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil,omitempty" name:"LimitVpc"`

	// 限制下载的ip设置。
	LimitIp []*string `json:"LimitIp,omitnil,omitempty" name:"LimitIp"`

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
type DescribeBackupEncryptionStatusRequestParams struct {
	// 实例ID，格式如：cdb-XXXX。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeBackupEncryptionStatusRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cdb-XXXX。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	EncryptionStatus *string `json:"EncryptionStatus,omitnil,omitempty" name:"EncryptionStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 需要查询备份概览的云数据库产品类型。可取值为：mysql 指双节点/三节点的高可用实例，mysql-basic 指单节点云盘版实例，mysql-cluster 指云盘版（原集群版）实例。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeBackupOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询备份概览的云数据库产品类型。可取值为：mysql 指双节点/三节点的高可用实例，mysql-basic 指单节点云盘版实例，mysql-cluster 指云盘版（原集群版）实例。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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
	BackupCount *int64 `json:"BackupCount,omitnil,omitempty" name:"BackupCount"`

	// 用户在当前地域备份的总容量
	BackupVolume *int64 `json:"BackupVolume,omitnil,omitempty" name:"BackupVolume"`

	// 用户在当前地域备份的计费容量，即超出赠送容量的部分。
	BillingVolume *int64 `json:"BillingVolume,omitnil,omitempty" name:"BillingVolume"`

	// 用户在当前地域获得的赠送备份容量。
	FreeVolume *int64 `json:"FreeVolume,omitnil,omitempty" name:"FreeVolume"`

	// 用户在当前地域的异地备份总容量。
	RemoteBackupVolume *int64 `json:"RemoteBackupVolume,omitnil,omitempty" name:"RemoteBackupVolume"`

	// 归档备份容量，包含数据备份以及日志备份。
	BackupArchiveVolume *int64 `json:"BackupArchiveVolume,omitnil,omitempty" name:"BackupArchiveVolume"`

	// 标准存储备份容量，包含数据备份以及日志备份。
	BackupStandbyVolume *int64 `json:"BackupStandbyVolume,omitnil,omitempty" name:"BackupStandbyVolume"`

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
type DescribeBackupSummariesRequestParams struct {
	// 需要查询备份实时统计的云数据库产品类型。可取值为：mysql 指双节点/三节点的高可用实例，mysql-basic 指单节点云盘版实例，mysql-cluster 指云盘版（原集群版）实例。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 分页查询数据的偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询数据的条目限制，默认值为20。最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 指定按某一项排序，可选值包括： BackupVolume: 备份容量， DataBackupVolume: 数据备份容量， BinlogBackupVolume: 日志备份容量， AutoBackupVolume: 自动备份容量， ManualBackupVolume: 手动备份容量。默认按照BackupVolume排序。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 指定排序方向，可选值包括： ASC: 正序， DESC: 逆序。默认值为 ASC。
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

type DescribeBackupSummariesRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询备份实时统计的云数据库产品类型。可取值为：mysql 指双节点/三节点的高可用实例，mysql-basic 指单节点云盘版实例，mysql-cluster 指云盘版（原集群版）实例。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 分页查询数据的偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询数据的条目限制，默认值为20。最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 指定按某一项排序，可选值包括： BackupVolume: 备份容量， DataBackupVolume: 数据备份容量， BinlogBackupVolume: 日志备份容量， AutoBackupVolume: 自动备份容量， ManualBackupVolume: 手动备份容量。默认按照BackupVolume排序。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 指定排序方向，可选值包括： ASC: 正序， DESC: 逆序。默认值为 ASC。
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
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
	Items []*BackupSummaryItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 实例备份统计总条目数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

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
type DescribeBackupsRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 偏移量，最小值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值为20，最小值为1，最大值为1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 偏移量，最小值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值为20，最小值为1，最大值为1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合查询条件的备份信息详情。
	Items []*BackupInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 需要查询日志备份概览的云数据库产品类型。可取值为：mysql 指双节点/三节点的高可用实例，mysql-basic 指单节点云盘版实例，mysql-cluster 指云盘版（原集群版）实例。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeBinlogBackupOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询日志备份概览的云数据库产品类型。可取值为：mysql 指双节点/三节点的高可用实例，mysql-basic 指单节点云盘版实例，mysql-cluster 指云盘版（原集群版）实例。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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
	BinlogBackupVolume *int64 `json:"BinlogBackupVolume,omitnil,omitempty" name:"BinlogBackupVolume"`

	// 总的日志备份个数，包含异地日志备份。
	BinlogBackupCount *int64 `json:"BinlogBackupCount,omitnil,omitempty" name:"BinlogBackupCount"`

	// 异地日志备份容量（单位为字节）。
	RemoteBinlogVolume *int64 `json:"RemoteBinlogVolume,omitnil,omitempty" name:"RemoteBinlogVolume"`

	// 异地日志备份个数。
	RemoteBinlogCount *int64 `json:"RemoteBinlogCount,omitnil,omitempty" name:"RemoteBinlogCount"`

	// 归档日志备份容量（单位为字节）。
	BinlogArchiveVolume *int64 `json:"BinlogArchiveVolume,omitnil,omitempty" name:"BinlogArchiveVolume"`

	// 归档日志备份个数。
	BinlogArchiveCount *int64 `json:"BinlogArchiveCount,omitnil,omitempty" name:"BinlogArchiveCount"`

	// 标准存储日志备份容量（单位为字节）。
	BinlogStandbyVolume *int64 `json:"BinlogStandbyVolume,omitnil,omitempty" name:"BinlogStandbyVolume"`

	// 标准存储日志备份个数。
	BinlogStandbyCount *int64 `json:"BinlogStandbyCount,omitnil,omitempty" name:"BinlogStandbyCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 偏移量，最小值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值为20，最小值为1，最大值为1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// binlog最早开始时间，时间格式：2016-03-17 02:10:37
	MinStartTime *string `json:"MinStartTime,omitnil,omitempty" name:"MinStartTime"`

	// binlog最晚开始时间，时间格式：2016-03-17 02:10:37
	MaxStartTime *string `json:"MaxStartTime,omitnil,omitempty" name:"MaxStartTime"`

	// 返回binlog列表是否包含MinStartTime起始节点，默认为否
	ContainsMinStartTime *bool `json:"ContainsMinStartTime,omitnil,omitempty" name:"ContainsMinStartTime"`
}

type DescribeBinlogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 偏移量，最小值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值为20，最小值为1，最大值为1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// binlog最早开始时间，时间格式：2016-03-17 02:10:37
	MinStartTime *string `json:"MinStartTime,omitnil,omitempty" name:"MinStartTime"`

	// binlog最晚开始时间，时间格式：2016-03-17 02:10:37
	MaxStartTime *string `json:"MaxStartTime,omitnil,omitempty" name:"MaxStartTime"`

	// 返回binlog列表是否包含MinStartTime起始节点，默认为否
	ContainsMinStartTime *bool `json:"ContainsMinStartTime,omitnil,omitempty" name:"ContainsMinStartTime"`
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
	delete(f, "MinStartTime")
	delete(f, "MaxStartTime")
	delete(f, "ContainsMinStartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogsResponseParams struct {
	// 符合查询条件的日志文件总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合查询条件的二进制日志文件详情。
	Items []*BinlogInfo `json:"Items,omitnil,omitempty" name:"Items"`

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
type DescribeCPUExpandStrategyInfoRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeCPUExpandStrategyInfoRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeCPUExpandStrategyInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCPUExpandStrategyInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCPUExpandStrategyInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCPUExpandStrategyInfoResponseParams struct {
	// 策略类型。输出值：auto、manual、timeInterval、period。
	// 说明：1. auto 表示自动扩容。2. manual 表示自定义扩容，扩容时间为立即生效。3. timeInterval 表示自定义扩容，扩容时间为按时间段。4. period 表示自定义扩容，扩容时间为按周期。5. 如果返回为 NULL 说明尚未开通弹性扩容策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 自定义扩容，且扩容时间为立即生效时的 CPU。Type 为 manual 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpandCpu *int64 `json:"ExpandCpu,omitnil,omitempty" name:"ExpandCpu"`

	// 自动扩容策略。Type 为 auto 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoStrategy *AutoStrategy `json:"AutoStrategy,omitnil,omitempty" name:"AutoStrategy"`

	// 按周期扩容策略。当 Type 为 period 时有效。
	PeriodStrategy *PeriodStrategy `json:"PeriodStrategy,omitnil,omitempty" name:"PeriodStrategy"`

	// 按时间段扩容策略。当 Type 为 timeInterval 时有效。
	TimeIntervalStrategy *TimeIntervalStrategy `json:"TimeIntervalStrategy,omitnil,omitempty" name:"TimeIntervalStrategy"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCPUExpandStrategyInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCPUExpandStrategyInfoResponseParams `json:"Response"`
}

func (r *DescribeCPUExpandStrategyInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCPUExpandStrategyInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdbProxyInfoRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 代理组 ID。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

type DescribeCdbProxyInfoRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 代理组 ID。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

func (r *DescribeCdbProxyInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdbProxyInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProxyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCdbProxyInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdbProxyInfoResponseParams struct {
	// 代理组数量
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 代理组信息
	ProxyInfos []*ProxyGroupInfo `json:"ProxyInfos,omitnil,omitempty" name:"ProxyInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCdbProxyInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCdbProxyInfoResponseParams `json:"Response"`
}

func (r *DescribeCdbProxyInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdbProxyInfoResponse) FromJsonString(s string) error {
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
	DataResult *CdbZoneDataResult `json:"DataResult,omitnil,omitempty" name:"DataResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 查询指定源实例的克隆任务列表。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/api/236/15872) 接口获取实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页查询时的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询时的每页条目数，默认值为20，建议最大取值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeCloneListRequest struct {
	*tchttp.BaseRequest
	
	// 查询指定源实例的克隆任务列表。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/api/236/15872) 接口获取实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页查询时的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询时的每页条目数，默认值为20，建议最大取值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 克隆任务列表。
	Items []*CloneItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeClusterInfoRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	// 说明：仅能输入实例架构为云盘版的实例 ID，对应控制台实例配置显示为“云盘版（云盘）”的实例。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeClusterInfoRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	// 说明：仅能输入实例架构为云盘版的实例 ID，对应控制台实例配置显示为“云盘版（云盘）”的实例。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeClusterInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterInfoResponseParams struct {
	// 实例名称。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 云盘版实例的读写地址信息。
	ReadWriteAddress *AddressInfo `json:"ReadWriteAddress,omitnil,omitempty" name:"ReadWriteAddress"`

	// 云盘版实例的只读地址信息。
	ReadOnlyAddress []*AddressInfo `json:"ReadOnlyAddress,omitnil,omitempty" name:"ReadOnlyAddress"`

	// 云盘版实例的节点列表信息。
	NodeList []*ClusterNodeInfo `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// 只读空间保护阈值,单位GB
	ReadonlyLimit *int64 `json:"ReadonlyLimit,omitnil,omitempty" name:"ReadonlyLimit"`

	// 实例节点数。
	NodeCount *int64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterInfoResponseParams `json:"Response"`
}

func (r *DescribeClusterInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCpuExpandHistoryRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 扩容策略，值包括：all，manual，auto
	ExpandStrategy *string `json:"ExpandStrategy,omitnil,omitempty" name:"ExpandStrategy"`

	// 扩容状态，值包括：all，extend，reduce，extend_failed
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 查询的开始时间。只能查看30天内的扩容历史，格式为 Integer 的时间戳（秒级）。
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询的结束时间。只能查看30天内的扩容历史，格式为 Integer 的时间戳（秒级）。
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页入参
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页入参
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeCpuExpandHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 扩容策略，值包括：all，manual，auto
	ExpandStrategy *string `json:"ExpandStrategy,omitnil,omitempty" name:"ExpandStrategy"`

	// 扩容状态，值包括：all，extend，reduce，extend_failed
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 查询的开始时间。只能查看30天内的扩容历史，格式为 Integer 的时间戳（秒级）。
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询的结束时间。只能查看30天内的扩容历史，格式为 Integer 的时间戳（秒级）。
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页入参
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页入参
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeCpuExpandHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCpuExpandHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ExpandStrategy")
	delete(f, "Status")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCpuExpandHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCpuExpandHistoryResponseParams struct {
	// 满足查询要求的扩容历史
	Items []*HistoryJob `json:"Items,omitnil,omitempty" name:"Items"`

	// 总数出参
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCpuExpandHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCpuExpandHistoryResponseParams `json:"Response"`
}

func (r *DescribeCpuExpandHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCpuExpandHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBFeaturesRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBFeaturesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	IsSupportAudit *bool `json:"IsSupportAudit,omitnil,omitempty" name:"IsSupportAudit"`

	// 开启审计是否需要升级内核版本。
	AuditNeedUpgrade *bool `json:"AuditNeedUpgrade,omitnil,omitempty" name:"AuditNeedUpgrade"`

	// 是否支持数据库加密功能。
	IsSupportEncryption *bool `json:"IsSupportEncryption,omitnil,omitempty" name:"IsSupportEncryption"`

	// 开启加密是否需要升级内核版本。
	EncryptionNeedUpgrade *bool `json:"EncryptionNeedUpgrade,omitnil,omitempty" name:"EncryptionNeedUpgrade"`

	// 是否为异地只读实例。
	IsRemoteRo *bool `json:"IsRemoteRo,omitnil,omitempty" name:"IsRemoteRo"`

	// 主实例所在地域。
	// 说明：此参数可能返回空值，您可忽略此出参返回值。如需获取实例所在地域详情，您可调用 [查询实例列表](https://cloud.tencent.com/document/product/236/15872) 接口查询。
	MasterRegion *string `json:"MasterRegion,omitnil,omitempty" name:"MasterRegion"`

	// 是否支持小版本升级。
	IsSupportUpdateSubVersion *bool `json:"IsSupportUpdateSubVersion,omitnil,omitempty" name:"IsSupportUpdateSubVersion"`

	// 当前内核版本。
	CurrentSubVersion *string `json:"CurrentSubVersion,omitnil,omitempty" name:"CurrentSubVersion"`

	// 可供升级的内核版本。
	TargetSubVersion *string `json:"TargetSubVersion,omitnil,omitempty" name:"TargetSubVersion"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，时间格式如：2016-01-01 00:00:01。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，时间格式如：2016-01-01 23:59:59。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页参数，偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，单次请求返回的数量，默认值为20，最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDBImportRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，时间格式如：2016-01-01 00:00:01。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，时间格式如：2016-01-01 23:59:59。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页参数，偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，单次请求返回的数量，默认值为20，最小值为1，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 返回的导入操作记录列表。
	Items []*ImportRecord `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBInstanceCharsetRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBInstanceConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// 主实例部署方式，可能的返回值：0 - 单可用部署，1 - 多可用区部署。
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 实例可用区信息，格式如 "ap-shanghai-1"。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 备库的配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaveConfig *SlaveConfig `json:"SlaveConfig,omitnil,omitempty" name:"SlaveConfig"`

	// 强同步实例第二备库的配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupConfig *BackupConfig `json:"BackupConfig,omitnil,omitempty" name:"BackupConfig"`

	// 是否切换备库。
	Switched *bool `json:"Switched,omitnil,omitempty" name:"Switched"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBInstanceGTIDRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	IsGTIDOpen *int64 `json:"IsGTIDOpen,omitnil,omitempty" name:"IsGTIDOpen"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	// 说明：仅主实例支持查询，此项仅支持输入主实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBInstanceInfoRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	// 说明：仅主实例支持查询，此项仅支持输入主实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 是否开通加密，YES 已开通，NO 未开通。
	Encryption *string `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// 加密使用的密钥 ID 。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 密钥所在地域。
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`

	// 当前 CDB 后端服务使用的 KMS 服务的默认地域。
	DefaultKmsRegion *string `json:"DefaultKmsRegion,omitnil,omitempty" name:"DefaultKmsRegion"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeDBInstanceLogToCLSRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// CLS服务所在地域
	ClsRegion *string `json:"ClsRegion,omitnil,omitempty" name:"ClsRegion"`
}

type DescribeDBInstanceLogToCLSRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// CLS服务所在地域
	ClsRegion *string `json:"ClsRegion,omitnil,omitempty" name:"ClsRegion"`
}

func (r *DescribeDBInstanceLogToCLSRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceLogToCLSRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClsRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceLogToCLSRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceLogToCLSResponseParams struct {
	// 错误日志投递CLS配置
	ErrorLog *LogToCLSConfig `json:"ErrorLog,omitnil,omitempty" name:"ErrorLog"`

	// 慢日志投递CLS配置
	SlowLog *LogToCLSConfig `json:"SlowLog,omitnil,omitempty" name:"SlowLog"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceLogToCLSResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceLogToCLSResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceLogToCLSResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceLogToCLSResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceRebootTimeRequestParams struct {
	// 实例的 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	// 说明：可输入多个实例 ID 进行查询，json 格式如下。
	// [
	//     "cdb-30z11v8s",
	//     "cdb-93h11efg"
	//   ]
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeDBInstanceRebootTimeRequest struct {
	*tchttp.BaseRequest
	
	// 实例的 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	// 说明：可输入多个实例 ID 进行查询，json 格式如下。
	// [
	//     "cdb-30z11v8s",
	//     "cdb-93h11efg"
	//   ]
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 返回的参数信息。
	Items []*InstanceRebootTime `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例类型，可取值：1 - 主实例，2 - 灾备实例，3 - 只读实例。
	InstanceTypes []*uint64 `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 实例的内网 IP 地址。
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// 实例状态，可取值：<br>0 - 创建中<br>1 - 运行中<br>4 - 正在进行隔离操作<br>5 - 已隔离（可在回收站恢复开机）
	Status []*uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 偏移量，默认值为 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为 20，最大值为 2000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 安全组 ID。当使用安全组 ID 为过滤条件时，需指定 WithSecurityGroup 参数为 1。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 付费类型，可取值：0 - 包年包月，1 - 小时计费。
	PayTypes []*uint64 `json:"PayTypes,omitnil,omitempty" name:"PayTypes"`

	// 实例名称。
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// 实例任务状态，可能取值：<br>0 - 没有任务<br>1 - 升级中<br>2 - 数据导入中<br>3 - 开放Slave中<br>4 - 外网访问开通中<br>5 - 批量操作执行中<br>6 - 回档中<br>7 - 外网访问关闭中<br>8 - 密码修改中<br>9 - 实例名修改中<br>10 - 重启中<br>12 - 自建迁移中<br>13 - 删除库表中<br>14 - 灾备实例创建同步中<br>15 - 升级待切换<br>16 - 升级切换中<br>17 - 升级切换完成<br>19 - 参数设置待执行<br>34 - 原地升级待执行
	TaskStatus []*uint64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 实例数据库引擎版本，可能取值：5.1、5.5、5.6 和 5.7。
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// 私有网络的 ID。
	VpcIds []*uint64 `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// 可用区的 ID。
	ZoneIds []*uint64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 子网 ID。
	SubnetIds []*uint64 `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 是否锁定标记，可选值：0 - 不锁定，1 - 锁定，默认为0。
	CdbErrors []*int64 `json:"CdbErrors,omitnil,omitempty" name:"CdbErrors"`

	// 返回结果集排序的字段，目前支持："InstanceId"，"InstanceName"，"CreateTime"，"DeadlineTime"。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回结果集排序方式。目前支持值："ASC" - 表示升序，"DESC" - 表示降序，默认为 "DESC"。
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// 是否以安全组 ID 为过滤条件。
	// 说明：0表示否，1表示是。
	WithSecurityGroup *int64 `json:"WithSecurityGroup,omitnil,omitempty" name:"WithSecurityGroup"`

	// 是否包含独享集群详细信息，可取值：0 - 不包含，1 - 包含。
	WithExCluster *int64 `json:"WithExCluster,omitnil,omitempty" name:"WithExCluster"`

	// 独享集群 ID。
	ExClusterId *string `json:"ExClusterId,omitnil,omitempty" name:"ExClusterId"`

	// 实例 ID。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 初始化标记，可取值：0 - 未初始化，1 - 初始化。
	InitFlag *int64 `json:"InitFlag,omitnil,omitempty" name:"InitFlag"`

	// 是否包含灾备关系对应的实例，可取值：0 - 不包含，1 - 包含。默认取值为1。如果拉取主实例，则灾备关系的数据在DrInfo字段中， 如果拉取灾备实例， 则灾备关系的数据在MasterInfo字段中。灾备关系中只包含部分基本的数据，详细的数据需要自行调接口拉取。
	WithDr *int64 `json:"WithDr,omitnil,omitempty" name:"WithDr"`

	// 是否包含只读实例，可取值：0 - 不包含，1 - 包含。默认取值为1。
	WithRo *int64 `json:"WithRo,omitnil,omitempty" name:"WithRo"`

	// 是否包含主实例，可取值：0 - 不包含，1 - 包含。默认取值为1。
	WithMaster *int64 `json:"WithMaster,omitnil,omitempty" name:"WithMaster"`

	// 置放群组ID列表。
	DeployGroupIds []*string `json:"DeployGroupIds,omitnil,omitempty" name:"DeployGroupIds"`

	// 是否以标签键为过滤条件。
	TagKeysForSearch []*string `json:"TagKeysForSearch,omitnil,omitempty" name:"TagKeysForSearch"`

	// 金融围拢 ID 。
	CageIds []*string `json:"CageIds,omitnil,omitempty" name:"CageIds"`

	// 标签值
	TagValues []*string `json:"TagValues,omitnil,omitempty" name:"TagValues"`

	// 私有网络字符型vpcId
	UniqueVpcIds []*string `json:"UniqueVpcIds,omitnil,omitempty" name:"UniqueVpcIds"`

	// 私有网络字符型subnetId
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitnil,omitempty" name:"UniqSubnetIds"`

	// 标签键值
	// 请注意，创建中的实例无法查询到标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 数据库代理 IP 。
	ProxyVips []*string `json:"ProxyVips,omitnil,omitempty" name:"ProxyVips"`

	// 数据库代理 ID 。
	ProxyIds []*string `json:"ProxyIds,omitnil,omitempty" name:"ProxyIds"`

	// 数据库引擎类型。可选值为：InnoDB、RocksDB。
	EngineTypes []*string `json:"EngineTypes,omitnil,omitempty" name:"EngineTypes"`

	// 是否获取集群版实例节点信息，可填：true 或 false。默认为 false。
	QueryClusterInfo *bool `json:"QueryClusterInfo,omitnil,omitempty" name:"QueryClusterInfo"`
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 项目 ID。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例类型，可取值：1 - 主实例，2 - 灾备实例，3 - 只读实例。
	InstanceTypes []*uint64 `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 实例的内网 IP 地址。
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// 实例状态，可取值：<br>0 - 创建中<br>1 - 运行中<br>4 - 正在进行隔离操作<br>5 - 已隔离（可在回收站恢复开机）
	Status []*uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 偏移量，默认值为 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为 20，最大值为 2000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 安全组 ID。当使用安全组 ID 为过滤条件时，需指定 WithSecurityGroup 参数为 1。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 付费类型，可取值：0 - 包年包月，1 - 小时计费。
	PayTypes []*uint64 `json:"PayTypes,omitnil,omitempty" name:"PayTypes"`

	// 实例名称。
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// 实例任务状态，可能取值：<br>0 - 没有任务<br>1 - 升级中<br>2 - 数据导入中<br>3 - 开放Slave中<br>4 - 外网访问开通中<br>5 - 批量操作执行中<br>6 - 回档中<br>7 - 外网访问关闭中<br>8 - 密码修改中<br>9 - 实例名修改中<br>10 - 重启中<br>12 - 自建迁移中<br>13 - 删除库表中<br>14 - 灾备实例创建同步中<br>15 - 升级待切换<br>16 - 升级切换中<br>17 - 升级切换完成<br>19 - 参数设置待执行<br>34 - 原地升级待执行
	TaskStatus []*uint64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 实例数据库引擎版本，可能取值：5.1、5.5、5.6 和 5.7。
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// 私有网络的 ID。
	VpcIds []*uint64 `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// 可用区的 ID。
	ZoneIds []*uint64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 子网 ID。
	SubnetIds []*uint64 `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 是否锁定标记，可选值：0 - 不锁定，1 - 锁定，默认为0。
	CdbErrors []*int64 `json:"CdbErrors,omitnil,omitempty" name:"CdbErrors"`

	// 返回结果集排序的字段，目前支持："InstanceId"，"InstanceName"，"CreateTime"，"DeadlineTime"。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回结果集排序方式。目前支持值："ASC" - 表示升序，"DESC" - 表示降序，默认为 "DESC"。
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// 是否以安全组 ID 为过滤条件。
	// 说明：0表示否，1表示是。
	WithSecurityGroup *int64 `json:"WithSecurityGroup,omitnil,omitempty" name:"WithSecurityGroup"`

	// 是否包含独享集群详细信息，可取值：0 - 不包含，1 - 包含。
	WithExCluster *int64 `json:"WithExCluster,omitnil,omitempty" name:"WithExCluster"`

	// 独享集群 ID。
	ExClusterId *string `json:"ExClusterId,omitnil,omitempty" name:"ExClusterId"`

	// 实例 ID。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 初始化标记，可取值：0 - 未初始化，1 - 初始化。
	InitFlag *int64 `json:"InitFlag,omitnil,omitempty" name:"InitFlag"`

	// 是否包含灾备关系对应的实例，可取值：0 - 不包含，1 - 包含。默认取值为1。如果拉取主实例，则灾备关系的数据在DrInfo字段中， 如果拉取灾备实例， 则灾备关系的数据在MasterInfo字段中。灾备关系中只包含部分基本的数据，详细的数据需要自行调接口拉取。
	WithDr *int64 `json:"WithDr,omitnil,omitempty" name:"WithDr"`

	// 是否包含只读实例，可取值：0 - 不包含，1 - 包含。默认取值为1。
	WithRo *int64 `json:"WithRo,omitnil,omitempty" name:"WithRo"`

	// 是否包含主实例，可取值：0 - 不包含，1 - 包含。默认取值为1。
	WithMaster *int64 `json:"WithMaster,omitnil,omitempty" name:"WithMaster"`

	// 置放群组ID列表。
	DeployGroupIds []*string `json:"DeployGroupIds,omitnil,omitempty" name:"DeployGroupIds"`

	// 是否以标签键为过滤条件。
	TagKeysForSearch []*string `json:"TagKeysForSearch,omitnil,omitempty" name:"TagKeysForSearch"`

	// 金融围拢 ID 。
	CageIds []*string `json:"CageIds,omitnil,omitempty" name:"CageIds"`

	// 标签值
	TagValues []*string `json:"TagValues,omitnil,omitempty" name:"TagValues"`

	// 私有网络字符型vpcId
	UniqueVpcIds []*string `json:"UniqueVpcIds,omitnil,omitempty" name:"UniqueVpcIds"`

	// 私有网络字符型subnetId
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitnil,omitempty" name:"UniqSubnetIds"`

	// 标签键值
	// 请注意，创建中的实例无法查询到标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 数据库代理 IP 。
	ProxyVips []*string `json:"ProxyVips,omitnil,omitempty" name:"ProxyVips"`

	// 数据库代理 ID 。
	ProxyIds []*string `json:"ProxyIds,omitnil,omitempty" name:"ProxyIds"`

	// 数据库引擎类型。可选值为：InnoDB、RocksDB。
	EngineTypes []*string `json:"EngineTypes,omitnil,omitempty" name:"EngineTypes"`

	// 是否获取集群版实例节点信息，可填：true 或 false。默认为 false。
	QueryClusterInfo *bool `json:"QueryClusterInfo,omitnil,omitempty" name:"QueryClusterInfo"`
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
	delete(f, "EngineTypes")
	delete(f, "QueryClusterInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesResponseParams struct {
	// 符合查询条件的实例总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例详细信息列表。
	Items []*InstanceInfo `json:"Items,omitnil,omitempty" name:"Items"`

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
type DescribeDBPriceRequestParams struct {
	// 实例时长，单位：月，最小值 1，最大值为 36；查询按量计费价格时，该字段无效。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 可用区信息，格式如 "ap-guangzhou-2"。具体能设置的值请通过 <a href="https://cloud.tencent.com/document/api/236/17229">DescribeDBZoneConfig</a> 接口查询。InstanceId为空时该参数为必填项。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例数量，默认值为 1，最小值 1，最大值为 100。InstanceId为空时该参数为必填项。
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 实例内存大小，单位：MB。InstanceId 为空时该参数为必填项。为保证传入值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可售卖的实例内存大小范围。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB。InstanceId 为空时该参数为必填项。为保证传入值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可售卖的硬盘大小范围。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 实例类型，默认为 master，支持值包括：master - 表示主实例，ro - 表示只读实例，dr - 表示灾备实例。InstanceId为空时该参数为必填项。
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 付费类型，支持值包括：PRE_PAID - 包年包月，HOUR_PAID - 按量计费。InstanceId为空时该参数为必填项。
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制。
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// 实例隔离类型。支持值包括： "UNIVERSAL" - 通用型实例， "EXCLUSIVE" - 独享型实例， "BASIC_V2" - 单节点云盘版实例。 "CLOUD_NATIVE_CLUSTER" - 集群版标准型， "CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - 集群版加强型。   不指定则默认为通用型实例。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 实例节点数。对于 RO 和 基础版实例， 该值默认为1。 如果需要询价三节点实例， 请将该值设置为3。其余主实例该值默认为2。
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// 询价实例的CPU核心数目，单位：核，为保证传入 CPU 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可售卖的核心数目，当未指定该值时，将按照 Memory 大小补全一个默认值。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 询价续费实例ID。如需查询实例续费价格，填写InstanceId和Period即可。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 按量计费阶梯。仅PayType=HOUR_PAID有效，支持值包括：1，2，3。阶梯时长见https://cloud.tencent.com/document/product/236/18335。
	Ladder *uint64 `json:"Ladder,omitnil,omitempty" name:"Ladder"`

	// 磁盘类型，查询集群版、单节点云盘版实例价格可以指定该参数。支持值包括： "CLOUD_SSD" - SSD云硬盘， "CLOUD_HSSD" - 增强型SSD云硬盘。  默认为 SSD云硬盘。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`
}

type DescribeDBPriceRequest struct {
	*tchttp.BaseRequest
	
	// 实例时长，单位：月，最小值 1，最大值为 36；查询按量计费价格时，该字段无效。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 可用区信息，格式如 "ap-guangzhou-2"。具体能设置的值请通过 <a href="https://cloud.tencent.com/document/api/236/17229">DescribeDBZoneConfig</a> 接口查询。InstanceId为空时该参数为必填项。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例数量，默认值为 1，最小值 1，最大值为 100。InstanceId为空时该参数为必填项。
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 实例内存大小，单位：MB。InstanceId 为空时该参数为必填项。为保证传入值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可售卖的实例内存大小范围。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB。InstanceId 为空时该参数为必填项。为保证传入值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可售卖的硬盘大小范围。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 实例类型，默认为 master，支持值包括：master - 表示主实例，ro - 表示只读实例，dr - 表示灾备实例。InstanceId为空时该参数为必填项。
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 付费类型，支持值包括：PRE_PAID - 包年包月，HOUR_PAID - 按量计费。InstanceId为空时该参数为必填项。
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制。
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// 实例隔离类型。支持值包括： "UNIVERSAL" - 通用型实例， "EXCLUSIVE" - 独享型实例， "BASIC_V2" - 单节点云盘版实例。 "CLOUD_NATIVE_CLUSTER" - 集群版标准型， "CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - 集群版加强型。   不指定则默认为通用型实例。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 实例节点数。对于 RO 和 基础版实例， 该值默认为1。 如果需要询价三节点实例， 请将该值设置为3。其余主实例该值默认为2。
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// 询价实例的CPU核心数目，单位：核，为保证传入 CPU 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可售卖的核心数目，当未指定该值时，将按照 Memory 大小补全一个默认值。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 询价续费实例ID。如需查询实例续费价格，填写InstanceId和Period即可。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 按量计费阶梯。仅PayType=HOUR_PAID有效，支持值包括：1，2，3。阶梯时长见https://cloud.tencent.com/document/product/236/18335。
	Ladder *uint64 `json:"Ladder,omitnil,omitempty" name:"Ladder"`

	// 磁盘类型，查询集群版、单节点云盘版实例价格可以指定该参数。支持值包括： "CLOUD_SSD" - SSD云硬盘， "CLOUD_HSSD" - 增强型SSD云硬盘。  默认为 SSD云硬盘。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`
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
	delete(f, "DiskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBPriceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBPriceResponseParams struct {
	// 实例价格，单位：分。
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 实例原价，单位：分。
	OriginalPrice *int64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 货币单位。CNY-人民币，USD-美元。
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 该值默认为False，表示当传入只读实例ID时，查询操作的是对应只读组的安全组。如果需要操作只读实例ID的安全组， 需要将该入参置为True。
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`

	// 变更集群版实例只读组时，InstanceId传实例id，需要额外指定该参数表示操作只读组。 如果操作读写节点则不需指定该参数。
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 该值默认为False，表示当传入只读实例ID时，查询操作的是对应只读组的安全组。如果需要操作只读实例ID的安全组， 需要将该入参置为True。
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`

	// 变更集群版实例只读组时，InstanceId传实例id，需要额外指定该参数表示操作只读组。 如果操作读写节点则不需指定该参数。
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
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
	delete(f, "OpResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsResponseParams struct {
	// 安全组详情。
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
type DescribeDBSwitchRecordsRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值为50，最小值为1，最大值为1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDBSwitchRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值为50，最小值为1，最大值为1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例切换记录详情。
	Items []*DBSwitchInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 需要查询数据备份概览的云数据库产品类型。可取值为：mysql 指双节点/三节点的高可用实例，mysql-basic 指单节点云盘版实例，mysql-cluster 指云盘版（原集群版）实例。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeDataBackupOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询数据备份概览的云数据库产品类型。可取值为：mysql 指双节点/三节点的高可用实例，mysql-basic 指单节点云盘版实例，mysql-cluster 指云盘版（原集群版）实例。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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
	DataBackupVolume *int64 `json:"DataBackupVolume,omitnil,omitempty" name:"DataBackupVolume"`

	// 当前地域的数据备份总个数。
	DataBackupCount *int64 `json:"DataBackupCount,omitnil,omitempty" name:"DataBackupCount"`

	// 当前地域的自动备份总容量。
	AutoBackupVolume *int64 `json:"AutoBackupVolume,omitnil,omitempty" name:"AutoBackupVolume"`

	// 当前地域的自动备份总个数。
	AutoBackupCount *int64 `json:"AutoBackupCount,omitnil,omitempty" name:"AutoBackupCount"`

	// 当前地域的手动备份总容量。
	ManualBackupVolume *int64 `json:"ManualBackupVolume,omitnil,omitempty" name:"ManualBackupVolume"`

	// 当前地域的手动备份总个数。
	ManualBackupCount *int64 `json:"ManualBackupCount,omitnil,omitempty" name:"ManualBackupCount"`

	// 异地备份总容量。
	RemoteBackupVolume *int64 `json:"RemoteBackupVolume,omitnil,omitempty" name:"RemoteBackupVolume"`

	// 异地备份总个数。
	RemoteBackupCount *int64 `json:"RemoteBackupCount,omitnil,omitempty" name:"RemoteBackupCount"`

	// 当前地域归档备份总容量。
	DataBackupArchiveVolume *int64 `json:"DataBackupArchiveVolume,omitnil,omitempty" name:"DataBackupArchiveVolume"`

	// 当前地域归档备份总个数。
	DataBackupArchiveCount *int64 `json:"DataBackupArchiveCount,omitnil,omitempty" name:"DataBackupArchiveCount"`

	// 当前地域标准存储备份总容量。
	DataBackupStandbyVolume *int64 `json:"DataBackupStandbyVolume,omitnil,omitempty" name:"DataBackupStandbyVolume"`

	// 当前地域标准存储备份总个数。
	DataBackupStandbyCount *int64 `json:"DataBackupStandbyCount,omitnil,omitempty" name:"DataBackupStandbyCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 偏移量，最小值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 单次请求数量，默认值为20，最小值为1，最大值为5000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 匹配数据库库名的正则表达式。
	DatabaseRegexp *string `json:"DatabaseRegexp,omitnil,omitempty" name:"DatabaseRegexp"`
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 偏移量，最小值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 单次请求数量，默认值为20，最小值为1，最大值为5000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 匹配数据库库名的正则表达式。
	DatabaseRegexp *string `json:"DatabaseRegexp,omitnil,omitempty" name:"DatabaseRegexp"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例中的数据库名称列表。
	Items []*string `json:"Items,omitnil,omitempty" name:"Items"`

	// 数据库名以及字符集
	DatabaseList []*DatabasesWithCharacterLists `json:"DatabaseList,omitnil,omitempty" name:"DatabaseList"`

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
type DescribeDefaultParamsRequestParams struct {
	// 引擎版本，目前支持 ["5.1", "5.5", "5.6", "5.7", "8.0"]。
	// 说明：引擎版本为必填。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 默认参数模板类型。支持值包括："HIGH_STABILITY" - 高稳定模板，"HIGH_PERFORMANCE" - 高性能模板。默认值为：HIGH_STABILITY。
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// 参数模板引擎，默认值：InnoDB，可取值：InnoDB、RocksDB。
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
}

type DescribeDefaultParamsRequest struct {
	*tchttp.BaseRequest
	
	// 引擎版本，目前支持 ["5.1", "5.5", "5.6", "5.7", "8.0"]。
	// 说明：引擎版本为必填。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 默认参数模板类型。支持值包括："HIGH_STABILITY" - 高稳定模板，"HIGH_PERFORMANCE" - 高性能模板。默认值为：HIGH_STABILITY。
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// 参数模板引擎，默认值：InnoDB，可取值：InnoDB、RocksDB。
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
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
	delete(f, "EngineType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDefaultParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultParamsResponseParams struct {
	// 参数个数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 参数详情。
	Items []*ParameterDetail `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// 置放群组名称。
	DeployGroupName *string `json:"DeployGroupName,omitnil,omitempty" name:"DeployGroupName"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeDeployGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// 置放群组名称。
	DeployGroupName *string `json:"DeployGroupName,omitnil,omitempty" name:"DeployGroupName"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 返回列表。
	Items []*DeployGroupInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回当天最近Count个5分钟粒度的监控数据。最小值1，最大值288，不传该参数默认返回当天所有5分钟粒度监控数据。
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type DescribeDeviceMonitorInfoRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回当天最近Count个5分钟粒度的监控数据。最小值1，最大值288，不传该参数默认返回当天所有5分钟粒度监控数据。
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
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
	Cpu *DeviceCpuInfo `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例内存监控数据
	Mem *DeviceMemInfo `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 实例网络监控数据
	Net *DeviceNetInfo `json:"Net,omitnil,omitempty" name:"Net"`

	// 实例磁盘监控数据
	Disk *DeviceDiskInfo `json:"Disk,omitnil,omitempty" name:"Disk"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间戳。例如1585142640，秒级。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳。例如1585142640，秒级。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 要匹配的关键字列表，最多支持15个关键字，支持模糊匹配。
	KeyWords []*string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`

	// 分页的返回数量，默认为100，最大为400。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 仅在实例为主实例或者灾备实例时生效，可选值：slave，代表拉取从机的日志。
	InstType *string `json:"InstType,omitnil,omitempty" name:"InstType"`
}

type DescribeErrorLogDataRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间戳。例如1585142640，秒级。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳。例如1585142640，秒级。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 要匹配的关键字列表，最多支持15个关键字，支持模糊匹配。
	KeyWords []*string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`

	// 分页的返回数量，默认为100，最大为400。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 仅在实例为主实例或者灾备实例时生效，可选值：slave，代表拉取从机的日志。
	InstType *string `json:"InstType,omitnil,omitempty" name:"InstType"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 返回的记录。
	Items []*ErrlogItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeInstanceAlarmEventsRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 事件查询范围开始时间，闭区间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 事件查询范围截止时间，闭区间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 事件名称。 Outofmemory - 内存OOM（有状态事件）; Switch - 主从切换（有状态事件）; Roremove - 只读实例剔除（有状态事件）; MemoryUsedHigh - 内存使用率过高（有状态事件）; CPUExpansion - CPU性能扩容（无状态事件）; CPUExpansionFailed - CPU性能扩容失败（无状态事件）; CPUContraction - CPU性能回缩（无状态事件）; Restart - 实例重启（有状态事件）; ServerFailureNodeMigration - ServerFailureNodeMigration（有状态事件）; PlannedSwitch - 计划内主备切换（无状态事件）; OverusedReadonlySet - 实例将被锁定（无状态事件）; OverusedReadWriteSet - 实例解除锁定（无状态事件）。
	EventName []*string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 事件状态。"1" - 发生事件；"0" - 恢复事件；"-" - 无状态事件。
	EventStatus *string `json:"EventStatus,omitnil,omitempty" name:"EventStatus"`

	// 排序方式。按事件发生事件进行排序，"DESC"-倒排；”ASC“-正序，默认倒排。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 事件展示数量。默认为100，最大为200。
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 节点 ID。
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

type DescribeInstanceAlarmEventsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 事件查询范围开始时间，闭区间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 事件查询范围截止时间，闭区间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 事件名称。 Outofmemory - 内存OOM（有状态事件）; Switch - 主从切换（有状态事件）; Roremove - 只读实例剔除（有状态事件）; MemoryUsedHigh - 内存使用率过高（有状态事件）; CPUExpansion - CPU性能扩容（无状态事件）; CPUExpansionFailed - CPU性能扩容失败（无状态事件）; CPUContraction - CPU性能回缩（无状态事件）; Restart - 实例重启（有状态事件）; ServerFailureNodeMigration - ServerFailureNodeMigration（有状态事件）; PlannedSwitch - 计划内主备切换（无状态事件）; OverusedReadonlySet - 实例将被锁定（无状态事件）; OverusedReadWriteSet - 实例解除锁定（无状态事件）。
	EventName []*string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 事件状态。"1" - 发生事件；"0" - 恢复事件；"-" - 无状态事件。
	EventStatus *string `json:"EventStatus,omitnil,omitempty" name:"EventStatus"`

	// 排序方式。按事件发生事件进行排序，"DESC"-倒排；”ASC“-正序，默认倒排。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 事件展示数量。默认为100，最大为200。
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 节点 ID。
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

func (r *DescribeInstanceAlarmEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAlarmEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "EventName")
	delete(f, "EventStatus")
	delete(f, "Order")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "NodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceAlarmEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceAlarmEventsResponseParams struct {
	// 事件数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 事件信息。查询不到信息时，Items为null。
	Items []*InstEventInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceAlarmEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceAlarmEventsResponseParams `json:"Response"`
}

func (r *DescribeInstanceAlarmEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAlarmEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamRecordsRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值：20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeInstanceParamRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值：20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 参数修改记录。
	Items []*ParamRecord `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 参数详情。
	Items []*ParameterDetail `json:"Items,omitnil,omitempty" name:"Items"`

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
type DescribeInstanceUpgradeCheckJobRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 目标数据库版本。
	// 说明：可选值5.6、5.7、8.0，不支持跨版本升级，升级后不支持版本降级。
	DstMysqlVersion *string `json:"DstMysqlVersion,omitnil,omitempty" name:"DstMysqlVersion"`
}

type DescribeInstanceUpgradeCheckJobRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 目标数据库版本。
	// 说明：可选值5.6、5.7、8.0，不支持跨版本升级，升级后不支持版本降级。
	DstMysqlVersion *string `json:"DstMysqlVersion,omitnil,omitempty" name:"DstMysqlVersion"`
}

func (r *DescribeInstanceUpgradeCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceUpgradeCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DstMysqlVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceUpgradeCheckJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceUpgradeCheckJobResponseParams struct {
	// 24小时内是否存在历史升级校验任务
	ExistUpgradeCheckJob *bool `json:"ExistUpgradeCheckJob,omitnil,omitempty" name:"ExistUpgradeCheckJob"`

	// 任务id
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceUpgradeCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceUpgradeCheckJobResponseParams `json:"Response"`
}

func (r *DescribeInstanceUpgradeCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceUpgradeCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceUpgradeTypeRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 目标实例 CPU 的核数。为保证传入值有效，请使用 [DescribeCdbZoneConfig](https://cloud.tencent.com/document/product/236/80281) 获取实例可售卖的 CPU 值范围。
	DstCpu *float64 `json:"DstCpu,omitnil,omitempty" name:"DstCpu"`

	// 目标实例内存大小，单位：MB。为保证传入值有效，请使用 [DescribeCdbZoneConfig](https://cloud.tencent.com/document/product/236/80281) 获取实例可售卖的内存大小范围。
	DstMemory *uint64 `json:"DstMemory,omitnil,omitempty" name:"DstMemory"`

	// 目标实例磁盘大小，单位：GB。为保证传入值有效，请使用 [DescribeCdbZoneConfig](https://cloud.tencent.com/document/product/236/80281) 获取实例可售卖的磁盘大小范围。
	DstDisk *uint64 `json:"DstDisk,omitnil,omitempty" name:"DstDisk"`

	// 目标实例数据库版本。可选值：5.6，5.7，8.0。
	DstVersion *string `json:"DstVersion,omitnil,omitempty" name:"DstVersion"`

	// 目标实例部署模型。默认为0，支持值包括：0 - 表示单可用区，1 - 表示多可用区。
	DstDeployMode *int64 `json:"DstDeployMode,omitnil,omitempty" name:"DstDeployMode"`

	// 目标实例复制类型，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制。
	DstProtectMode *int64 `json:"DstProtectMode,omitnil,omitempty" name:"DstProtectMode"`

	// 目标实例备机1可用区 ID。可使用 [DescribeCdbZoneConfig](https://cloud.tencent.com/document/product/236/80281) 获取可用区 ID。
	DstSlaveZone *int64 `json:"DstSlaveZone,omitnil,omitempty" name:"DstSlaveZone"`

	// 目标实例备机2可用区 ID。可使用 [DescribeCdbZoneConfig](https://cloud.tencent.com/document/product/236/80281) 获取可用区 ID。
	DstBackupZone *int64 `json:"DstBackupZone,omitnil,omitempty" name:"DstBackupZone"`

	// 目标实例类型。支持值包括："CUSTOM" - 通用型实例，"EXCLUSIVE" - 独享型实例，"ONTKE" - ONTKE 单节点实例，"CLOUD_NATIVE_CLUSTER" - 云盘版标准型，"CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - 云盘版加强型。
	DstCdbType *string `json:"DstCdbType,omitnil,omitempty" name:"DstCdbType"`

	// 目标实例主可用区 ID。可使用 [DescribeCdbZoneConfig](https://cloud.tencent.com/document/product/236/80281) 获取可用区 ID。
	DstZoneId *int64 `json:"DstZoneId,omitnil,omitempty" name:"DstZoneId"`

	// 独享集群 CDB 实例的节点分布情况。
	NodeDistribution *NodeDistribution `json:"NodeDistribution,omitnil,omitempty" name:"NodeDistribution"`

	// 集群版的节点拓扑配置。Nodeld信息可通过 [DescribeClusterInfo](https://cloud.tencent.com/document/api/236/105116) 接口获取。
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`
}

type DescribeInstanceUpgradeTypeRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 目标实例 CPU 的核数。为保证传入值有效，请使用 [DescribeCdbZoneConfig](https://cloud.tencent.com/document/product/236/80281) 获取实例可售卖的 CPU 值范围。
	DstCpu *float64 `json:"DstCpu,omitnil,omitempty" name:"DstCpu"`

	// 目标实例内存大小，单位：MB。为保证传入值有效，请使用 [DescribeCdbZoneConfig](https://cloud.tencent.com/document/product/236/80281) 获取实例可售卖的内存大小范围。
	DstMemory *uint64 `json:"DstMemory,omitnil,omitempty" name:"DstMemory"`

	// 目标实例磁盘大小，单位：GB。为保证传入值有效，请使用 [DescribeCdbZoneConfig](https://cloud.tencent.com/document/product/236/80281) 获取实例可售卖的磁盘大小范围。
	DstDisk *uint64 `json:"DstDisk,omitnil,omitempty" name:"DstDisk"`

	// 目标实例数据库版本。可选值：5.6，5.7，8.0。
	DstVersion *string `json:"DstVersion,omitnil,omitempty" name:"DstVersion"`

	// 目标实例部署模型。默认为0，支持值包括：0 - 表示单可用区，1 - 表示多可用区。
	DstDeployMode *int64 `json:"DstDeployMode,omitnil,omitempty" name:"DstDeployMode"`

	// 目标实例复制类型，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制。
	DstProtectMode *int64 `json:"DstProtectMode,omitnil,omitempty" name:"DstProtectMode"`

	// 目标实例备机1可用区 ID。可使用 [DescribeCdbZoneConfig](https://cloud.tencent.com/document/product/236/80281) 获取可用区 ID。
	DstSlaveZone *int64 `json:"DstSlaveZone,omitnil,omitempty" name:"DstSlaveZone"`

	// 目标实例备机2可用区 ID。可使用 [DescribeCdbZoneConfig](https://cloud.tencent.com/document/product/236/80281) 获取可用区 ID。
	DstBackupZone *int64 `json:"DstBackupZone,omitnil,omitempty" name:"DstBackupZone"`

	// 目标实例类型。支持值包括："CUSTOM" - 通用型实例，"EXCLUSIVE" - 独享型实例，"ONTKE" - ONTKE 单节点实例，"CLOUD_NATIVE_CLUSTER" - 云盘版标准型，"CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - 云盘版加强型。
	DstCdbType *string `json:"DstCdbType,omitnil,omitempty" name:"DstCdbType"`

	// 目标实例主可用区 ID。可使用 [DescribeCdbZoneConfig](https://cloud.tencent.com/document/product/236/80281) 获取可用区 ID。
	DstZoneId *int64 `json:"DstZoneId,omitnil,omitempty" name:"DstZoneId"`

	// 独享集群 CDB 实例的节点分布情况。
	NodeDistribution *NodeDistribution `json:"NodeDistribution,omitnil,omitempty" name:"NodeDistribution"`

	// 集群版的节点拓扑配置。Nodeld信息可通过 [DescribeClusterInfo](https://cloud.tencent.com/document/api/236/105116) 接口获取。
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`
}

func (r *DescribeInstanceUpgradeTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceUpgradeTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DstCpu")
	delete(f, "DstMemory")
	delete(f, "DstDisk")
	delete(f, "DstVersion")
	delete(f, "DstDeployMode")
	delete(f, "DstProtectMode")
	delete(f, "DstSlaveZone")
	delete(f, "DstBackupZone")
	delete(f, "DstCdbType")
	delete(f, "DstZoneId")
	delete(f, "NodeDistribution")
	delete(f, "ClusterTopology")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceUpgradeTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceUpgradeTypeResponseParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例升级类型。Trsf - 迁移升级，InPlace - 原地升级，Topology - 架构升级。
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceUpgradeTypeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceUpgradeTypeResponseParams `json:"Response"`
}

func (r *DescribeInstanceUpgradeTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceUpgradeTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLocalBinlogConfigRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeLocalBinlogConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	LocalBinlogConfig *LocalBinlogConfig `json:"LocalBinlogConfig,omitnil,omitempty" name:"LocalBinlogConfig"`

	// 该地域默认binlog保留策略。
	LocalBinlogConfigDefault *LocalBinlogConfigDefault `json:"LocalBinlogConfigDefault,omitnil,omitempty" name:"LocalBinlogConfigDefault"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 参数模板 ID。可通过 [DescribeParamTemplates](https://cloud.tencent.com/document/api/236/32659) 接口获取。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeParamTemplateInfoRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板 ID。可通过 [DescribeParamTemplates](https://cloud.tencent.com/document/api/236/32659) 接口获取。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 参数模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数模板对应实例版本，可取值：5.5、5.6、5.7、8.0。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 参数模板中的参数数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 参数详情
	Items []*ParameterDetail `json:"Items,omitnil,omitempty" name:"Items"`

	// 参数模板描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 参数模板类型。支持值包括："HIGH_STABILITY" - 高稳定模板，"HIGH_PERFORMANCE" - 高性能模板。
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// 参数模板引擎。支持值包括："InnoDB"，"RocksDB"。
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 引擎版本，缺省则查询所有。可取值为：5.5、5.6、5.7、8.0。
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// 引擎类型，缺省则查询所有。可取值为：InnoDB、RocksDB，不区分大小写。
	EngineTypes []*string `json:"EngineTypes,omitnil,omitempty" name:"EngineTypes"`

	// 模板名称，缺省则查询所有。支持模糊匹配。
	TemplateNames []*string `json:"TemplateNames,omitnil,omitempty" name:"TemplateNames"`

	// 模板 ID，缺省则查询所有。
	TemplateIds []*int64 `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`
}

type DescribeParamTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 引擎版本，缺省则查询所有。可取值为：5.5、5.6、5.7、8.0。
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// 引擎类型，缺省则查询所有。可取值为：InnoDB、RocksDB，不区分大小写。
	EngineTypes []*string `json:"EngineTypes,omitnil,omitempty" name:"EngineTypes"`

	// 模板名称，缺省则查询所有。支持模糊匹配。
	TemplateNames []*string `json:"TemplateNames,omitnil,omitempty" name:"TemplateNames"`

	// 模板 ID，缺省则查询所有。
	TemplateIds []*int64 `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 参数模板详情。
	Items []*ParamTemplateInfo `json:"Items,omitnil,omitempty" name:"Items"`

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
	// 项目 ID。可通过 [DescribeProjects](https://cloud.tencent.com/document/api/651/78725) 接口获取。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 项目 ID。可通过 [DescribeProjects](https://cloud.tencent.com/document/api/651/78725) 接口获取。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 安全组规则数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

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
type DescribeProxyCustomConfRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeProxyCustomConfRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 代理配置
	//
	// Deprecated: CustomConf is deprecated.
	CustomConf *CustomConfig `json:"CustomConf,omitnil,omitempty" name:"CustomConf"`

	// 权重限制
	WeightRule *Rule `json:"WeightRule,omitnil,omitempty" name:"WeightRule"`

	// 代理配置
	CustomConfInfo []*CustomConfig `json:"CustomConfInfo,omitnil,omitempty" name:"CustomConfInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeProxySupportParamRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeProxySupportParamRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeProxySupportParamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySupportParamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxySupportParamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySupportParamResponseParams struct {
	// 支持最大代理版本
	ProxyVersion *string `json:"ProxyVersion,omitnil,omitempty" name:"ProxyVersion"`

	// 是否支持连接池
	SupportPool *bool `json:"SupportPool,omitnil,omitempty" name:"SupportPool"`

	// 连接池最小值
	PoolMin *uint64 `json:"PoolMin,omitnil,omitempty" name:"PoolMin"`

	// 连接池最大值
	PoolMax *uint64 `json:"PoolMax,omitnil,omitempty" name:"PoolMax"`

	// 是否支持事务拆分
	SupportTransSplit *bool `json:"SupportTransSplit,omitnil,omitempty" name:"SupportTransSplit"`

	// 支持连接池的最小代理版本
	SupportPoolMinVersion *string `json:"SupportPoolMinVersion,omitnil,omitempty" name:"SupportPoolMinVersion"`

	// 支持事务拆分的最小代理版本
	SupportTransSplitMinVersion *string `json:"SupportTransSplitMinVersion,omitnil,omitempty" name:"SupportTransSplitMinVersion"`

	// 是否支持设置只读
	SupportReadOnly *bool `json:"SupportReadOnly,omitnil,omitempty" name:"SupportReadOnly"`

	// 是否自动均衡负载
	SupportAutoLoadBalance *bool `json:"SupportAutoLoadBalance,omitnil,omitempty" name:"SupportAutoLoadBalance"`

	// 是否支持接入模式
	SupportAccessMode *bool `json:"SupportAccessMode,omitnil,omitempty" name:"SupportAccessMode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProxySupportParamResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxySupportParamResponseParams `json:"Response"`
}

func (r *DescribeProxySupportParamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySupportParamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRemoteBackupConfigRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeRemoteBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	// 异地备份保留时间，单位为天
	ExpireDays *int64 `json:"ExpireDays,omitnil,omitempty" name:"ExpireDays"`

	// 异地数据备份开关，off - 关闭异地备份，on-开启异地备份
	RemoteBackupSave *string `json:"RemoteBackupSave,omitnil,omitempty" name:"RemoteBackupSave"`

	// 异地日志备份开关，off - 关闭异地备份，on-开启异地备份，只有在参数RemoteBackupSave为on时，RemoteBinlogSave参数才可设置为on
	RemoteBinlogSave *string `json:"RemoteBinlogSave,omitnil,omitempty" name:"RemoteBinlogSave"`

	// 用户已设置异地备份地域列表
	RemoteRegion []*string `json:"RemoteRegion,omitnil,omitempty" name:"RemoteRegion"`

	// 用户可设置异地备份地域列表
	RegionList []*string `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeRoGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	// RO 组信息数组，一个实例可关联多个 RO 组。
	RoGroups []*RoGroup `json:"RoGroups,omitnil,omitempty" name:"RoGroups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RoInstanceId *string `json:"RoInstanceId,omitnil,omitempty" name:"RoInstanceId"`

	// 主实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，该参数与RoInstanceId参数不能同时为空。注意，当传入参数包含RoInstanceId时，返回值为只读实例升级时的最小规格；当传入参数只包含MasterInstanceId时，返回值为只读实例购买时的最小规格。
	MasterInstanceId *string `json:"MasterInstanceId,omitnil,omitempty" name:"MasterInstanceId"`
}

type DescribeRoMinScaleRequest struct {
	*tchttp.BaseRequest
	
	// 只读实例ID，格式如：cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，该参数与MasterInstanceId参数不能同时为空。
	RoInstanceId *string `json:"RoInstanceId,omitnil,omitempty" name:"RoInstanceId"`

	// 主实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，该参数与RoInstanceId参数不能同时为空。注意，当传入参数包含RoInstanceId时，返回值为只读实例升级时的最小规格；当传入参数只包含MasterInstanceId时，返回值为只读实例购买时的最小规格。
	MasterInstanceId *string `json:"MasterInstanceId,omitnil,omitempty" name:"MasterInstanceId"`
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
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 磁盘规格大小, 单位为：GB。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 克隆实例与源实例是否在异地可用区，是:"true"，否:"false"
	IsRemoteZone *string `json:"IsRemoteZone,omitnil,omitempty" name:"IsRemoteZone"`

	// 克隆实例与源实例不在同一地域时需填写克隆实例所在地域，例："ap-guangzhou"
	BackupRegion *string `json:"BackupRegion,omitnil,omitempty" name:"BackupRegion"`
}

type DescribeRollbackRangeTimeRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表，单个实例 ID 的格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 克隆实例与源实例是否在异地可用区，是:"true"，否:"false"
	IsRemoteZone *string `json:"IsRemoteZone,omitnil,omitempty" name:"IsRemoteZone"`

	// 克隆实例与源实例不在同一地域时需填写克隆实例所在地域，例："ap-guangzhou"
	BackupRegion *string `json:"BackupRegion,omitnil,omitempty" name:"BackupRegion"`
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
	delete(f, "IsRemoteZone")
	delete(f, "BackupRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRollbackRangeTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollbackRangeTimeResponseParams struct {
	// 符合查询条件的实例总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 返回的参数信息。
	Items []*InstanceRollbackRangeTime `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 异步任务 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 分页参数，每次请求返回的记录数。默认值为20，建议最大取值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量。默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeRollbackTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872)。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 异步任务 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 分页参数，每次请求返回的记录数。默认值为20，建议最大取值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量。默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 回档任务详情。
	Items []*RollbackTask `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeSSLStatusRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	// 说明：实例 ID 和只读组 ID 两个参数选其一填写即可。若要查询双节点、三节点实例 SSL 开通情况，请填写实例 ID 参数进行查询。单节点（云盘）、云盘版实例不支持开启 SSL，因此不支持查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 只读组 ID。可通过 [DescribeRoGroups](https://cloud.tencent.com/document/api/236/40939) 接口获取。
	// 说明：实例 ID 和只读组 ID 两个参数选其一填写即可。若要查询只读实例或只读组 SSL 开通情况，请填写 RoGroupId 参数，并注意填写的都是只读组 ID。单节点（云盘）、云盘版实例不支持开启 SSL，因此不支持查询。
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`
}

type DescribeSSLStatusRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	// 说明：实例 ID 和只读组 ID 两个参数选其一填写即可。若要查询双节点、三节点实例 SSL 开通情况，请填写实例 ID 参数进行查询。单节点（云盘）、云盘版实例不支持开启 SSL，因此不支持查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 只读组 ID。可通过 [DescribeRoGroups](https://cloud.tencent.com/document/api/236/40939) 接口获取。
	// 说明：实例 ID 和只读组 ID 两个参数选其一填写即可。若要查询只读实例或只读组 SSL 开通情况，请填写 RoGroupId 参数，并注意填写的都是只读组 ID。单节点（云盘）、云盘版实例不支持开启 SSL，因此不支持查询。
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`
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
	delete(f, "InstanceId")
	delete(f, "RoGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSSLStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSSLStatusResponseParams struct {
	// 是否开通 SSL 。ON 代表开通 ，OFF 代表未开通。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 证书下载链接。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

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
type DescribeSlowLogDataRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间戳。例如 1585142640。
	// 说明：此参数单位为秒的时间戳。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳。例如 1585142640。
	// 说明：此参数单位为秒的时间戳。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 客户端 Host 列表。
	UserHosts []*string `json:"UserHosts,omitnil,omitempty" name:"UserHosts"`

	// 客户端 用户名 列表。
	UserNames []*string `json:"UserNames,omitnil,omitempty" name:"UserNames"`

	// 访问的 数据库 列表。
	DataBases []*string `json:"DataBases,omitnil,omitempty" name:"DataBases"`

	// 排序字段，当前支持字段及含义如下，默认值为 Timestamp。
	// 1. Timestamp：SQL 的执行时间
	// 2. QueryTime：SQL 的执行时长（秒）
	// 3. LockTime：锁时长（秒）
	// 4. RowsExamined：扫描行数
	// 5. RowsSent：结果集行数
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 升序还是降序排列。当前支持值为 ASC - 升序，DESC - 降序 ，默认值为 ASC。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 偏移量，默认为0，最大为9999。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 一次性返回的记录数量，默认为100，最大为800。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 仅在实例为主实例或者灾备实例时生效，可选值：slave，代表拉取从机的日志。
	InstType *string `json:"InstType,omitnil,omitempty" name:"InstType"`

	// 节点ID
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
}

type DescribeSlowLogDataRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间戳。例如 1585142640。
	// 说明：此参数单位为秒的时间戳。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳。例如 1585142640。
	// 说明：此参数单位为秒的时间戳。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 客户端 Host 列表。
	UserHosts []*string `json:"UserHosts,omitnil,omitempty" name:"UserHosts"`

	// 客户端 用户名 列表。
	UserNames []*string `json:"UserNames,omitnil,omitempty" name:"UserNames"`

	// 访问的 数据库 列表。
	DataBases []*string `json:"DataBases,omitnil,omitempty" name:"DataBases"`

	// 排序字段，当前支持字段及含义如下，默认值为 Timestamp。
	// 1. Timestamp：SQL 的执行时间
	// 2. QueryTime：SQL 的执行时长（秒）
	// 3. LockTime：锁时长（秒）
	// 4. RowsExamined：扫描行数
	// 5. RowsSent：结果集行数
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 升序还是降序排列。当前支持值为 ASC - 升序，DESC - 降序 ，默认值为 ASC。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 偏移量，默认为0，最大为9999。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 一次性返回的记录数量，默认为100，最大为800。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 仅在实例为主实例或者灾备实例时生效，可选值：slave，代表拉取从机的日志。
	InstType *string `json:"InstType,omitnil,omitempty" name:"InstType"`

	// 节点ID
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
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
	delete(f, "OpResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogDataResponseParams struct {
	// 符合条件的记录总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 查询到的记录。
	Items []*SlowLogItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 偏移量，默认值为0，最小值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值为20，最小值为1，最大值为1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 偏移量，默认值为0，最小值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值为20，最小值为1，最大值为1000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合查询条件的慢查询日志详情。
	Items []*SlowLogInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeSupportedPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	GlobalSupportedPrivileges []*string `json:"GlobalSupportedPrivileges,omitnil,omitempty" name:"GlobalSupportedPrivileges"`

	// 实例支持的数据库权限。
	DatabaseSupportedPrivileges []*string `json:"DatabaseSupportedPrivileges,omitnil,omitempty" name:"DatabaseSupportedPrivileges"`

	// 实例支持的数据库表权限。
	TableSupportedPrivileges []*string `json:"TableSupportedPrivileges,omitnil,omitempty" name:"TableSupportedPrivileges"`

	// 实例支持的数据库列权限。
	ColumnSupportedPrivileges []*string `json:"ColumnSupportedPrivileges,omitnil,omitempty" name:"ColumnSupportedPrivileges"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeTableColumnsRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，可使用[查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名称，可使用 [查询数据库](https://cloud.tencent.com/document/api/236/17493) 接口获得。
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 数据库中的表的名称。
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`
}

type DescribeTableColumnsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，可使用[查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名称，可使用 [查询数据库](https://cloud.tencent.com/document/api/236/17493) 接口获得。
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 数据库中的表的名称。
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`
}

func (r *DescribeTableColumnsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableColumnsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Database")
	delete(f, "Table")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableColumnsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableColumnsResponseParams struct {
	// 符合查询条件的实例总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 返回的数据库列信息。
	Items []*string `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTableColumnsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableColumnsResponseParams `json:"Response"`
}

func (r *DescribeTableColumnsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableColumnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库的名称。
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 记录偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为20，最大值为5000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 匹配数据库表名的正则表达式，规则同 MySQL 官网
	TableRegexp *string `json:"TableRegexp,omitnil,omitempty" name:"TableRegexp"`
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库的名称。
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 记录偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为20，最大值为5000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 匹配数据库表名的正则表达式，规则同 MySQL 官网
	TableRegexp *string `json:"TableRegexp,omitnil,omitempty" name:"TableRegexp"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 返回的数据库表信息。
	Items []*string `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例列表。实例 ID 可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。可传入的数组长度暂无限制。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小。默认为15。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTagsOfInstanceIdsRequest struct {
	*tchttp.BaseRequest
	
	// 实例列表。实例 ID 可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。可传入的数组长度暂无限制。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小。默认为15。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 实例标签信息。
	Rows []*TagsInfoOfInstance `json:"Rows,omitnil,omitempty" name:"Rows"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 异步任务请求 ID，执行云数据库相关操作返回的 AsyncRequestId。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

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
	TaskTypes []*int64 `json:"TaskTypes,omitnil,omitempty" name:"TaskTypes"`

	// 任务状态，不传值则查询所有任务状态，支持的值包括：
	// -1 - 未定义；
	// 0 - 初始化；
	// 1 - 运行中；
	// 2 - 执行成功；
	// 3 - 执行失败；
	// 4 - 已终止；
	// 5 - 已删除；
	// 6 - 已暂停。
	TaskStatus []*int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 第一个任务的开始时间，用于范围查询，时间格式如：2017-12-31 10:40:01。
	StartTimeBegin *string `json:"StartTimeBegin,omitnil,omitempty" name:"StartTimeBegin"`

	// 最后一个任务的开始时间，用于范围查询，时间格式如：2017-12-31 10:40:01。
	StartTimeEnd *string `json:"StartTimeEnd,omitnil,omitempty" name:"StartTimeEnd"`

	// 记录偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 异步任务请求 ID，执行云数据库相关操作返回的 AsyncRequestId。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

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
	TaskTypes []*int64 `json:"TaskTypes,omitnil,omitempty" name:"TaskTypes"`

	// 任务状态，不传值则查询所有任务状态，支持的值包括：
	// -1 - 未定义；
	// 0 - 初始化；
	// 1 - 运行中；
	// 2 - 执行成功；
	// 3 - 执行失败；
	// 4 - 已终止；
	// 5 - 已删除；
	// 6 - 已暂停。
	TaskStatus []*int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 第一个任务的开始时间，用于范围查询，时间格式如：2017-12-31 10:40:01。
	StartTimeBegin *string `json:"StartTimeBegin,omitnil,omitempty" name:"StartTimeBegin"`

	// 最后一个任务的开始时间，用于范围查询，时间格式如：2017-12-31 10:40:01。
	StartTimeEnd *string `json:"StartTimeEnd,omitnil,omitempty" name:"StartTimeEnd"`

	// 记录偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 返回的实例任务信息。
	Items []*TaskDetail `json:"Items,omitnil,omitempty" name:"Items"`

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
type DescribeTimeWindowRequestParams struct {
	// 实例ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeTimeWindowRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	Monday []*string `json:"Monday,omitnil,omitempty" name:"Monday"`

	// 星期二的可维护时间列表。
	Tuesday []*string `json:"Tuesday,omitnil,omitempty" name:"Tuesday"`

	// 星期三的可维护时间列表。
	Wednesday []*string `json:"Wednesday,omitnil,omitempty" name:"Wednesday"`

	// 星期四的可维护时间列表。
	Thursday []*string `json:"Thursday,omitnil,omitempty" name:"Thursday"`

	// 星期五的可维护时间列表。
	Friday []*string `json:"Friday,omitnil,omitempty" name:"Friday"`

	// 星期六的可维护时间列表。
	Saturday []*string `json:"Saturday,omitnil,omitempty" name:"Saturday"`

	// 星期日的可维护时间列表。
	Sunday []*string `json:"Sunday,omitnil,omitempty" name:"Sunday"`

	// 最大数据延迟阈值
	MaxDelayTime *uint64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 记录偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeUploadedFilesRequest struct {
	*tchttp.BaseRequest
	
	// 文件路径。该字段应填用户主账号的OwnerUin信息。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 记录偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 单次请求返回的数量，默认值为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 返回的SQL文件列表。
	Items []*SqlFileInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Rate []*DeviceCpuRateInfo `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 实例CPU监控数据
	Load []*int64 `json:"Load,omitnil,omitempty" name:"Load"`
}

type DeviceCpuRateInfo struct {
	// Cpu核编号
	CpuCore *int64 `json:"CpuCore,omitnil,omitempty" name:"CpuCore"`

	// Cpu使用率
	Rate []*int64 `json:"Rate,omitnil,omitempty" name:"Rate"`
}

type DeviceDiskInfo struct {
	// 平均每秒有百分之几的时间用于IO操作
	IoRatioPerSec []*int64 `json:"IoRatioPerSec,omitnil,omitempty" name:"IoRatioPerSec"`

	// 平均每次设备I/O操作的等待时间*100，单位为毫秒。例如：该值为201，表示平均每次I/O操作等待时间为：201/100=2.1毫秒
	IoWaitTime []*int64 `json:"IoWaitTime,omitnil,omitempty" name:"IoWaitTime"`

	// 磁盘平均每秒完成的读操作次数总和*100。例如：该值为2002，表示磁盘平均每秒完成读操作为：2002/100=20.2次
	Read []*int64 `json:"Read,omitnil,omitempty" name:"Read"`

	// 磁盘平均每秒完成的写操作次数总和*100。例如：该值为30001，表示磁盘平均每秒完成写操作为：30001/100=300.01次
	Write []*int64 `json:"Write,omitnil,omitempty" name:"Write"`

	// 磁盘空间容量，每两个一组，第一个为已使用容量，第二个为磁盘总容量
	CapacityRatio []*int64 `json:"CapacityRatio,omitnil,omitempty" name:"CapacityRatio"`
}

type DeviceMemInfo struct {
	// 总内存大小。free命令中Mem:一行total的值,单位：KB
	Total []*int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 已使用内存。free命令中Mem:一行used的值,单位：KB
	Used []*int64 `json:"Used,omitnil,omitempty" name:"Used"`
}

type DeviceNetInfo struct {
	// tcp连接数
	Conn []*int64 `json:"Conn,omitnil,omitempty" name:"Conn"`

	// 网卡入包量，单位：个/秒
	PackageIn []*int64 `json:"PackageIn,omitnil,omitempty" name:"PackageIn"`

	// 网卡出包量，单位：个/秒
	PackageOut []*int64 `json:"PackageOut,omitnil,omitempty" name:"PackageOut"`

	// 入流量，单位：kbps
	FlowIn []*int64 `json:"FlowIn,omitnil,omitempty" name:"FlowIn"`

	// 出流量，单位：kbps
	FlowOut []*int64 `json:"FlowOut,omitnil,omitempty" name:"FlowOut"`
}

// Predefined struct for user
type DisassociateSecurityGroupsRequestParams struct {
	// 安全组 ID。可通过 [DescribeDBSecurityGroups](https://cloud.tencent.com/document/api/236/15854) 接口获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 实例 ID 列表，一个或者多个实例 ID 组成的数组。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 当传入只读实例 ID 时，默认操作的是对应只读组的安全组。如果需要操作只读实例 ID 的安全组，需要将该入参置为 True，默认为 False。
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 安全组 ID。可通过 [DescribeDBSecurityGroups](https://cloud.tencent.com/document/api/236/15854) 接口获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 实例 ID 列表，一个或者多个实例 ID 组成的数组。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 当传入只读实例 ID 时，默认操作的是对应只读组的安全组。如果需要操作只读实例 ID 的安全组，需要将该入参置为 True，默认为 False。
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`
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

type DiskTypeConfigItem struct {
	// 磁盘对应的实例类型。仅支持单节点基础型和集群版。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 可以选择的磁盘类型列表。
	DiskType []*string `json:"DiskType,omitnil,omitempty" name:"DiskType"`
}

type DrInfo struct {
	// 灾备实例状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 可用区信息
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 地域信息
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 实例同步状态。可能的返回值为：
	// 0 - 灾备未同步；
	// 1 - 灾备同步中；
	// 2 - 灾备同步成功；
	// 3 - 灾备同步失败；
	// 4 - 灾备同步修复中。
	SyncStatus *int64 `json:"SyncStatus,omitnil,omitempty" name:"SyncStatus"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例类型
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type ErrlogItem struct {
	// 错误发生时间。时间戳，秒级
	Timestamp *uint64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 错误详情
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type HistoryJob struct {
	// 操作类型
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 扩容类型
	ExpandType *string `json:"ExpandType,omitnil,omitempty" name:"ExpandType"`

	// 扩容开始时间
	// 说明：此项显示的格式是 int 类型的 unix 时间戳
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 扩容结束时间
	// 说明：此项显示的格式是 int 类型的 unix 时间戳
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 扩容前核数
	OldCpu *int64 `json:"OldCpu,omitnil,omitempty" name:"OldCpu"`

	// 扩容后核数
	NewCpu *int64 `json:"NewCpu,omitnil,omitempty" name:"NewCpu"`

	// 增减的cpu数
	ExtendCPUNum *int64 `json:"ExtendCPUNum,omitnil,omitempty" name:"ExtendCPUNum"`

	// extend_failed操作上报
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

type ImportRecord struct {
	// 状态值。0 - 初始化中，1 - 运行中，2 - 运行成功，3 - 运行失败。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态值，为负数时任务异常。
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 执行时间，单位：秒。
	CostTime *int64 `json:"CostTime,omitnil,omitempty" name:"CostTime"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 后端任务ID
	WorkId *string `json:"WorkId,omitnil,omitempty" name:"WorkId"`

	// 导入文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 执行进度，单位：百分比。
	Process *int64 `json:"Process,omitnil,omitempty" name:"Process"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 文件大小，单位：byte。
	FileSize *string `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 任务执行信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 任务ID
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 导入库表名
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 异步任务的请求ID
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`
}

type Inbound struct {
	// 策略，ACCEPT 或者 DROP
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 来源 IP 或 IP 段，例如192.168.0.0/16
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// 端口
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// 网络协议，支持 UDP、TCP 等
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// 规则限定的方向，进站规则为 INPUT
	Dir *string `json:"Dir,omitnil,omitempty" name:"Dir"`

	// 地址模块
	AddressModule *string `json:"AddressModule,omitnil,omitempty" name:"AddressModule"`

	// 规则ID，嵌套安全组的规则ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 规则描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

// Predefined struct for user
type InquiryPriceUpgradeInstancesRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升级后的内存大小，单位：MB，为保证传入 Memory 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可升级的内存规格。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 升级后的硬盘大小，单位：GB，为保证传入 Volume 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可升级的硬盘范围。
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 升级后的核心数目，单位：核，为保证传入 CPU 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可升级的核心数目，当未指定该值时，将按照 Memory 大小补全一个默认值。
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 数据复制方式，支持值包括：0 - 异步复制，1 - 半同步复制，2 - 强同步复制，升级主实例时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义。
	ProtectMode *uint64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// 实例隔离类型。支持值包括： "UNIVERSAL" - 通用型实例， "EXCLUSIVE" - 独享型实例， "BASIC" - 基础版实例。 不指定则默认为通用型实例。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 实例节点数。对于 RO 和 基础版实例， 该值默认为1。 如果需要询价三节点实例， 请将该值设置为3。其余主实例该值默认为2。
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`
}

type InquiryPriceUpgradeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升级后的内存大小，单位：MB，为保证传入 Memory 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可升级的内存规格。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 升级后的硬盘大小，单位：GB，为保证传入 Volume 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可升级的硬盘范围。
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 升级后的核心数目，单位：核，为保证传入 CPU 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可升级的核心数目，当未指定该值时，将按照 Memory 大小补全一个默认值。
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 数据复制方式，支持值包括：0 - 异步复制，1 - 半同步复制，2 - 强同步复制，升级主实例时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义。
	ProtectMode *uint64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// 实例隔离类型。支持值包括： "UNIVERSAL" - 通用型实例， "EXCLUSIVE" - 独享型实例， "BASIC" - 基础版实例。 不指定则默认为通用型实例。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 实例节点数。对于 RO 和 基础版实例， 该值默认为1。 如果需要询价三节点实例， 请将该值设置为3。其余主实例该值默认为2。
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`
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
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 实例原价，单位：分（人民币）。
	OriginalPrice *int64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type InstEventInfo struct {
	// 事件名称。
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 事件状态。
	EventStatus *string `json:"EventStatus,omitnil,omitempty" name:"EventStatus"`

	// 事件发生时间。
	OccurTime *string `json:"OccurTime,omitnil,omitempty" name:"OccurTime"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点ID
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

type InstanceAuditLogFilters struct {
	// 过滤项。目前支持以下搜索条件：
	// 
	// 包含、不包含、包含（分词维度）、不包含（分词维度）:
	// sql - SQL详情；alarmLevel - 告警等级；ruleTemplateId - 规则模板Id
	// 
	// 等于、不等于、包含、不包含：
	// host - 客户端地址；
	// user - 用户名；
	// dbName - 数据库名称；
	// 
	// 等于、不等于：
	// sqlType - SQL类型；
	// errCode - 错误码；
	// threadId - 线程ID；
	// 
	// 范围搜索（时间类型统一为微秒）：
	// execTime - 执行时间；
	// lockWaitTime - 执行时间；
	// ioWaitTime - IO等待时间；
	// trxLivingTime - 事物持续时间；
	// cpuTime - cpu时间；
	// checkRows - 扫描行数；
	// affectRows - 影响行数；
	// sentRows - 返回行数。
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

	// 过滤的值。反向查询时，多个值之前是且的关系，正向查询多个值是或的关系
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type InstanceDbAuditStatus struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计状态。ON-表示审计已开启，OFF-表示审计关闭
	AuditStatus *string `json:"AuditStatus,omitnil,omitempty" name:"AuditStatus"`

	// 任务状态。0-无任务；1-审计开启中，2-审计关闭中。
	AuditTask *uint64 `json:"AuditTask,omitnil,omitempty" name:"AuditTask"`

	// 日志保留时长。支持值包括：
	// 7 - 一周；
	// 30 - 一个月；
	// 90 - 三个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年。
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 高频存储时长。支持值包括：
	// 3 - 3天；
	// 7 - 一周；
	// 30 - 一个月；
	// 90 - 三个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年。
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// 低频存储时长。单位：天，等于日志保存时长减去高频存储时长。
	LowLogExpireDay *uint64 `json:"LowLogExpireDay,omitnil,omitempty" name:"LowLogExpireDay"`

	// 日志存储量(单位：GB)。
	BillingAmount *float64 `json:"BillingAmount,omitnil,omitempty" name:"BillingAmount"`

	// 高频存储量(单位：GB)。
	HighRealStorage *float64 `json:"HighRealStorage,omitnil,omitempty" name:"HighRealStorage"`

	// 低频存储量(单位：GB)。
	LowRealStorage *float64 `json:"LowRealStorage,omitnil,omitempty" name:"LowRealStorage"`

	// 是否为全审计。true-表示全审计。
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// 审计开通时间。
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// 实例相关信息
	InstanceInfo *AuditInstanceInfo `json:"InstanceInfo,omitnil,omitempty" name:"InstanceInfo"`

	// 总存储量(单位：GB)。
	RealStorage *float64 `json:"RealStorage,omitnil,omitempty" name:"RealStorage"`

	// 是否包含审计策略
	OldRule *bool `json:"OldRule,omitnil,omitempty" name:"OldRule"`

	// 实例所应用的规则模板。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
}

type InstanceInfo struct {
	// 外网状态，可能的返回值为：0-未开通外网；1-已开通外网；2-已关闭外网
	WanStatus *int64 `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// 可用区信息
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 初始化标志，可能的返回值为：0-未初始化；1-已初始化
	InitFlag *int64 `json:"InitFlag,omitnil,omitempty" name:"InitFlag"`

	// 只读vip信息。单独开通只读实例访问的只读实例才有该字段
	RoVipInfo *RoVipInfo `json:"RoVipInfo,omitnil,omitempty" name:"RoVipInfo"`

	// 内存容量，单位为 MB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例状态，可能的返回值：0-创建中；1-运行中；4-正在进行隔离操作；5-已隔离
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 私有网络 ID，例如：51102
	VpcId *int64 `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 备机信息
	SlaveInfo *SlaveInfo `json:"SlaveInfo,omitnil,omitempty" name:"SlaveInfo"`

	// 实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 硬盘容量，单位为 GB
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 自动续费标志，可能的返回值：0-未开通自动续费；1-已开通自动续费；2-已关闭自动续费
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 数据复制方式。0 - 异步复制；1 - 半同步复制；2 - 强同步复制
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// 只读组详细信息
	RoGroups []*RoGroup `json:"RoGroups,omitnil,omitempty" name:"RoGroups"`

	// 子网 ID，例如：2333
	SubnetId *int64 `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例类型，可能的返回值：1-主实例；2-灾备实例；3-只读实例
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 项目 ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 地域信息
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 实例到期时间
	DeadlineTime *string `json:"DeadlineTime,omitnil,omitempty" name:"DeadlineTime"`

	// 可用区部署方式。可能的值为：0 - 单可用区；1 - 多可用区
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 实例任务状态。0 - 没有任务 ,1 - 升级中,2 - 数据导入中,3 - 开放Slave中,4 - 外网访问开通中,5 - 批量操作执行中,6 - 回档中,7 - 外网访问关闭中,8 - 密码修改中,9 - 实例名修改中,10 - 重启中,12 - 自建迁移中,13 - 删除库表中,14 - 灾备实例创建同步中,15 - 升级待切换,16 - 升级切换中,17 - 升级切换完成
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 主实例详细信息
	MasterInfo *MasterInfo `json:"MasterInfo,omitnil,omitempty" name:"MasterInfo"`

	// 实例类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 内核版本
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 灾备实例详细信息
	DrInfo []*DrInfo `json:"DrInfo,omitnil,omitempty" name:"DrInfo"`

	// 外网域名
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// 外网端口号
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// 付费类型，可能的返回值：0-包年包月；1-按量计费
	PayType *int64 `json:"PayType,omitnil,omitempty" name:"PayType"`

	// 实例创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例 IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 端口号
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 磁盘写入是否被锁定（实例数据写入量已经超过磁盘配额）。0 -未被锁定 1 -已被锁定
	CdbError *int64 `json:"CdbError,omitnil,omitempty" name:"CdbError"`

	// 私有网络描述符，例如：“vpc-5v8wn9mg”
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 子网描述符，例如：“subnet-1typ0s7d”
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 物理 ID
	PhysicalId *string `json:"PhysicalId,omitnil,omitempty" name:"PhysicalId"`

	// 核心数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 每秒查询数量
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 可用区中文名称
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 物理机型
	DeviceClass *string `json:"DeviceClass,omitnil,omitempty" name:"DeviceClass"`

	// 置放群组 ID
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// 可用区 ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 节点数
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// 标签列表
	TagList []*TagInfoItem `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 引擎类型
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 最大延迟阈值
	MaxDelayTime *int64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`

	// 实例磁盘类型，仅云盘版实例才返回该值。可能的值为 CLOUD_SSD：SSD云硬盘， CLOUD_HSSD：增强型SSD云硬盘
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 当前扩容的CPU核心数。
	ExpandCpu *int64 `json:"ExpandCpu,omitnil,omitempty" name:"ExpandCpu"`

	// 实例集群版节点信息
	ClusterInfo []*ClusterInfo `json:"ClusterInfo,omitnil,omitempty" name:"ClusterInfo"`

	// 分析引擎节点列表
	AnalysisNodeInfos []*AnalysisNodeInfo `json:"AnalysisNodeInfos,omitnil,omitempty" name:"AnalysisNodeInfos"`

	// 设备带宽，单位G。当DeviceClass不为空时此参数才有效。例：25-表示当前设备带宽为25G；10-表示当前设备带宽为10G。
	DeviceBandwidth *uint64 `json:"DeviceBandwidth,omitnil,omitempty" name:"DeviceBandwidth"`
}

type InstanceRebootTime struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 预期重启时间，单位：秒。
	TimeInSeconds *int64 `json:"TimeInSeconds,omitnil,omitempty" name:"TimeInSeconds"`
}

type InstanceRollbackRangeTime struct {
	// 查询数据库错误码。0 - 正常，1600001 - 内部错误，1600003 - 入参异常，1600009 - 实例不存在，1624001 - DB 访问异常。
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 查询数据库错误信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 实例ID列表，单个实例Id的格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 可回档时间范围
	Times []*RollbackTimeRange `json:"Times,omitnil,omitempty" name:"Times"`
}

// Predefined struct for user
type IsolateDBInstanceRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type IsolateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	//
	// Deprecated: AsyncRequestId is deprecated.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 本地 binlog 保留时长，可取值范围：[6,168]。
	SaveHours *int64 `json:"SaveHours,omitnil,omitempty" name:"SaveHours"`

	// 本地 binlog 空间使用率，可取值范围：[30,50]。
	MaxUsage *int64 `json:"MaxUsage,omitnil,omitempty" name:"MaxUsage"`
}

type LocalBinlogConfigDefault struct {
	// 本地 binlog 保留时长，可取值范围：[6,168]。
	SaveHours *int64 `json:"SaveHours,omitnil,omitempty" name:"SaveHours"`

	// 本地 binlog 空间使用率，可取值范围：[30,50]。
	MaxUsage *int64 `json:"MaxUsage,omitnil,omitempty" name:"MaxUsage"`
}

type LogRuleTemplateInfo struct {
	// 模板ID。
	RuleTemplateId *string `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// 规则模板名
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// 告警等级。1-低风险，2-中风险，3-高风险。
	AlarmLevel *string `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 规则模板变更状态：0-未变更；1-已变更；2-已删除
	RuleTemplateStatus *int64 `json:"RuleTemplateStatus,omitnil,omitempty" name:"RuleTemplateStatus"`
}

type LogToCLSConfig struct {
	// 投递状态打开或者关闭
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// CLS日志集ID
	LogSetId *string `json:"LogSetId,omitnil,omitempty" name:"LogSetId"`

	// 日志主题ID
	LogTopicId *string `json:"LogTopicId,omitnil,omitempty" name:"LogTopicId"`

	// CLS服务所在地域
	ClsRegion *string `json:"ClsRegion,omitnil,omitempty" name:"ClsRegion"`
}

type MasterInfo struct {
	// 地域信息
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 可用区ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 可用区信息
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例长ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 实例状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例类型
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 任务状态
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 内存容量
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 硬盘容量
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 实例机型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 每秒查询数
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 私有网络ID
	VpcId *int64 `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *int64 `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 独享集群ID
	ExClusterId *string `json:"ExClusterId,omitnil,omitempty" name:"ExClusterId"`

	// 独享集群名称
	ExClusterName *string `json:"ExClusterName,omitnil,omitempty" name:"ExClusterName"`
}

type MigrateClusterRoInfo struct {
	// 只读实例名称
	RoInstanceId *string `json:"RoInstanceId,omitnil,omitempty" name:"RoInstanceId"`

	// 只读实例CPU核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 只读实例内存大小，单位：MB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 只读实例硬盘大小，单位：GB
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 磁盘类型。 CLOUD_SSD: SSD云硬盘; CLOUD_HSSD: 增强型SSD云硬盘
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 迁移实例类型。支持值包括： "CLOUD_NATIVE_CLUSTER" - 标准型集群版实例， "CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - 加强型集群版实例。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 只读实例所在ro组，例：cdbrg-xxx
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`

	// 实例当前告警策略id数组
	SrcAlarmPolicyList []*int64 `json:"SrcAlarmPolicyList,omitnil,omitempty" name:"SrcAlarmPolicyList"`
}

// Predefined struct for user
type ModifyAccountDescriptionRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 云数据库账号。可通过 [DescribeAccounts](https://cloud.tencent.com/document/api/236/17499) 接口获取。
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// 数据库账号的备注信息。最多支持输入255个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 云数据库账号。可通过 [DescribeAccounts](https://cloud.tencent.com/document/api/236/17499) 接口获取。
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// 数据库账号的备注信息。最多支持输入255个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

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
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 账户的名称。可通过 [DescribeAccounts](https://cloud.tencent.com/document/api/236/17499) 接口获取。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 账户的旧主机。格式：IP 形式，支持单个 IP 地址或者%。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 账户的新主机。格式：IP 形式，支持单个 IP 地址或者%。
	NewHost *string `json:"NewHost,omitnil,omitempty" name:"NewHost"`
}

type ModifyAccountHostRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 账户的名称。可通过 [DescribeAccounts](https://cloud.tencent.com/document/api/236/17499) 接口获取。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 账户的旧主机。格式：IP 形式，支持单个 IP 地址或者%。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 账户的新主机。格式：IP 形式，支持单个 IP 地址或者%。
	NewHost *string `json:"NewHost,omitnil,omitempty" name:"NewHost"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

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
type ModifyAccountMaxUserConnectionsRequestParams struct {
	// 云数据库账号。可通过 [DescribeAccounts](https://cloud.tencent.com/document/api/236/17499) 接口获取。
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设置账户最大可用连接数，最大可设置值为10240。
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`
}

type ModifyAccountMaxUserConnectionsRequest struct {
	*tchttp.BaseRequest
	
	// 云数据库账号。可通过 [DescribeAccounts](https://cloud.tencent.com/document/api/236/17499) 接口获取。
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设置账户最大可用连接数，最大可设置值为10240。
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库账号的新密码。密码应至少包含字母、数字和字符（_+-&=!@#$%^*()）中的两种，长度为8-64个字符。
	NewPassword *string `json:"NewPassword,omitnil,omitempty" name:"NewPassword"`

	// 云数据库账号。可通过 [DescribeAccounts](https://cloud.tencent.com/document/api/236/17499) 接口获取。
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

type ModifyAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库账号的新密码。密码应至少包含字母、数字和字符（_+-&=!@#$%^*()）中的两种，长度为8-64个字符。
	NewPassword *string `json:"NewPassword,omitnil,omitempty" name:"NewPassword"`

	// 云数据库账号。可通过 [DescribeAccounts](https://cloud.tencent.com/document/api/236/17499) 接口获取。
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库的账号，包括用户名和域名。可通过 [DescribeAccounts](https://cloud.tencent.com/document/api/236/17499) 接口获取。
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// 全局权限。其中，GlobalPrivileges 中权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE", "PROCESS", "DROP","REFERENCES","INDEX","ALTER","SHOW DATABASES","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER","CREATE USER","RELOAD","REPLICATION CLIENT","REPLICATION SLAVE"。
	// 注意，ModifyAction为空时，不传该参数表示清除该权限。
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// 数据库的权限。Privileges 权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE",	"DROP","REFERENCES","INDEX","ALTER","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER"。
	// 注意，ModifyAction为空时，不传该参数表示清除该权限。
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// 数据库中表的权限。Privileges 权限的可选值为：权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE",	"DROP","REFERENCES","INDEX","ALTER","CREATE VIEW","SHOW VIEW", "TRIGGER"。
	// 注意，ModifyAction为空时，不传该参数表示清除该权限。
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`

	// 数据库表中列的权限。Privileges 权限的可选值为："SELECT","INSERT","UPDATE","REFERENCES"。
	// 注意，ModifyAction为空时，不传该参数表示清除该权限。
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitnil,omitempty" name:"ColumnPrivileges"`

	// 该参数不为空时，为批量修改权限。可选值为：grant - 授予权限，revoke - 回收权限。
	ModifyAction *string `json:"ModifyAction,omitnil,omitempty" name:"ModifyAction"`
}

type ModifyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库的账号，包括用户名和域名。可通过 [DescribeAccounts](https://cloud.tencent.com/document/api/236/17499) 接口获取。
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// 全局权限。其中，GlobalPrivileges 中权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE", "PROCESS", "DROP","REFERENCES","INDEX","ALTER","SHOW DATABASES","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER","CREATE USER","RELOAD","REPLICATION CLIENT","REPLICATION SLAVE"。
	// 注意，ModifyAction为空时，不传该参数表示清除该权限。
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// 数据库的权限。Privileges 权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE",	"DROP","REFERENCES","INDEX","ALTER","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER"。
	// 注意，ModifyAction为空时，不传该参数表示清除该权限。
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// 数据库中表的权限。Privileges 权限的可选值为：权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE",	"DROP","REFERENCES","INDEX","ALTER","CREATE VIEW","SHOW VIEW", "TRIGGER"。
	// 注意，ModifyAction为空时，不传该参数表示清除该权限。
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`

	// 数据库表中列的权限。Privileges 权限的可选值为："SELECT","INSERT","UPDATE","REFERENCES"。
	// 注意，ModifyAction为空时，不传该参数表示清除该权限。
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitnil,omitempty" name:"ColumnPrivileges"`

	// 该参数不为空时，为批量修改权限。可选值为：grant - 授予权限，revoke - 回收权限。
	ModifyAction *string `json:"ModifyAction,omitnil,omitempty" name:"ModifyAction"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

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
type ModifyAuditConfigRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计日志保存时长。支持值包括：
	// 7 - 一周
	// 30 - 一个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年；
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 是否关闭审计服务。可选值：true - 关闭审计服务；false - 不关闭审计服务。默认值为 false。
	// 说明：
	// 1. 当关闭审计服务时，会删除用户的审计日志和文件，并删除该实例的所有审计策略。
	// 2. CloseAudit、LogExpireDay 必须至少提供一个，如果两个都提供则按照 CloseAudit 优先的逻辑处理。
	// 3. 可通过设置此参数来关闭审计服务，已关闭后不能通过此接口来开启审计服务。
	CloseAudit *bool `json:"CloseAudit,omitnil,omitempty" name:"CloseAudit"`

	// 高频审计日志保存时长。支持值包括：
	// 7 - 一周
	// 30 - 一个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年；
	HighLogExpireDay *int64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`
}

type ModifyAuditConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计日志保存时长。支持值包括：
	// 7 - 一周
	// 30 - 一个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年；
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 是否关闭审计服务。可选值：true - 关闭审计服务；false - 不关闭审计服务。默认值为 false。
	// 说明：
	// 1. 当关闭审计服务时，会删除用户的审计日志和文件，并删除该实例的所有审计策略。
	// 2. CloseAudit、LogExpireDay 必须至少提供一个，如果两个都提供则按照 CloseAudit 优先的逻辑处理。
	// 3. 可通过设置此参数来关闭审计服务，已关闭后不能通过此接口来开启审计服务。
	CloseAudit *bool `json:"CloseAudit,omitnil,omitempty" name:"CloseAudit"`

	// 高频审计日志保存时长。支持值包括：
	// 7 - 一周
	// 30 - 一个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年；
	HighLogExpireDay *int64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 审计规则名称。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 审计规则描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 审计规则过滤条件。若设置了过滤条件，则不会开启全审计。
	RuleFilters []*AuditFilter `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// 是否开启全审计。支持值包括：false – 不开启全审计，true – 开启全审计。用户未设置审计规则过滤条件时，默认开启全审计。
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`
}

type ModifyAuditRuleRequest struct {
	*tchttp.BaseRequest
	
	// 审计规则 ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 审计规则名称。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 审计规则描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 审计规则过滤条件。若设置了过滤条件，则不会开启全审计。
	RuleFilters []*AuditFilter `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// 是否开启全审计。支持值包括：false – 不开启全审计，true – 开启全审计。用户未设置审计规则过滤条件时，默认开启全审计。
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyAuditRuleTemplatesRequestParams struct {
	// 审计规则模板 ID。可通过 [DescribeAuditRuleTemplates](https://cloud.tencent.com/document/api/236/101811) 接口获取。
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
	
	// 审计规则模板 ID。可通过 [DescribeAuditRuleTemplates](https://cloud.tencent.com/document/api/236/101811) 接口获取。
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
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志保留时长。支持值包括：
	// 7 - 一周；
	// 30 - 一个月；
	// 90 - 三个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年。
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 高频日志保留时长。默认值为7，此项取值需小于等于 LogExpireDay，支持值包括：
	// 3 - 3天；
	// 7 - 一周；
	// 30 - 一个月；
	// 90 - 三个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年。
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// 修改实例审计规则为全审计。
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// 废弃。
	//
	// Deprecated: AuditRuleFilters is deprecated.
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// 规则模板 ID。可通过 [DescribeAuditRuleTemplates](https://cloud.tencent.com/document/api/236/101811) 接口获取。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
}

type ModifyAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志保留时长。支持值包括：
	// 7 - 一周；
	// 30 - 一个月；
	// 90 - 三个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年。
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 高频日志保留时长。默认值为7，此项取值需小于等于 LogExpireDay，支持值包括：
	// 3 - 3天；
	// 7 - 一周；
	// 30 - 一个月；
	// 90 - 三个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年。
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// 修改实例审计规则为全审计。
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// 废弃。
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// 规则模板 ID。可通过 [DescribeAuditRuleTemplates](https://cloud.tencent.com/document/api/236/101811) 接口获取。
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
type ModifyAutoRenewFlagRequestParams struct {
	// 实例的 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	// 说明：可输入多个实例 ID 进行修改，json 格式如下。
	// [
	//     "cdb-30z11v8s",
	//     "cdb-93h11efg"
	//   ]
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 自动续费标记，可取值的有：0 - 不自动续费，1 - 自动续费。
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`
}

type ModifyAutoRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 实例的 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	// 说明：可输入多个实例 ID 进行修改，json 格式如下。
	// [
	//     "cdb-30z11v8s",
	//     "cdb-93h11efg"
	//   ]
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 自动续费标记，可取值的有：0 - 不自动续费，1 - 自动续费。
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据备份文件的保留时间，单位为天。
	// 1. MySQL 双节点、三节点、云盘版数据备份文件可以保留7天 - 1830天。
	// 2. MySQL 单节点（云盘）数据备份文件可以保留7天 - 30天。
	ExpireDays *int64 `json:"ExpireDays,omitnil,omitempty" name:"ExpireDays"`

	// (将废弃，建议使用 BackupTimeWindow 参数) 备份时间范围，格式为：02:00-06:00，起点和终点时间目前限制为整点，目前可以选择的范围为： 00:00-12:00，02:00-06:00，06：00-10：00，10:00-14:00，14:00-18:00，18:00-22:00，22:00-02:00。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 自动备份方式，仅支持：physical - 物理冷备
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// binlog 的保留时间，单位为天。
	// 1. MySQL 双节点、三节点、云盘版日志备份文件可以保留7天 - 3650天。
	// 2. MySQL 单节点（云盘）日志备份文件可以保留7天 - 30天。
	BinlogExpireDays *int64 `json:"BinlogExpireDays,omitnil,omitempty" name:"BinlogExpireDays"`

	// 备份时间窗，比如要设置每周二和周日 10:00-14:00之间备份，该参数如下：{"Monday": "", "Tuesday": "10:00-14:00", "Wednesday": "", "Thursday": "", "Friday": "", "Saturday": "", "Sunday": "10:00-14:00"}    （注：可以设置一周的某几天备份，但是每天的备份时间需要设置为相同的时间段。 如果设置了该字段，将忽略StartTime字段的设置）
	BackupTimeWindow *CommonTimeWindow `json:"BackupTimeWindow,omitnil,omitempty" name:"BackupTimeWindow"`

	// 定期保留开关，off - 不开启定期保留策略，on - 开启定期保留策略，默认为off。
	EnableBackupPeriodSave *string `json:"EnableBackupPeriodSave,omitnil,omitempty" name:"EnableBackupPeriodSave"`

	// 长期保留开关,该字段功能暂未上线，可忽略。off - 不开启长期保留策略，on - 开启长期保留策略，默认为off，如果开启，则BackupPeriodSaveDays，BackupPeriodSaveInterval，BackupPeriodSaveCount参数无效
	EnableBackupPeriodLongTermSave *string `json:"EnableBackupPeriodLongTermSave,omitnil,omitempty" name:"EnableBackupPeriodLongTermSave"`

	// 定期保留最长天数，最小值：90，最大值：3650，默认值：1080
	BackupPeriodSaveDays *int64 `json:"BackupPeriodSaveDays,omitnil,omitempty" name:"BackupPeriodSaveDays"`

	// 定期保留策略周期，可取值：weekly - 周，monthly - 月， quarterly - 季度，yearly - 年，默认为monthly
	BackupPeriodSaveInterval *string `json:"BackupPeriodSaveInterval,omitnil,omitempty" name:"BackupPeriodSaveInterval"`

	// 定期保留的备份数量，最小值为1，最大值不超过定期保留策略周期内常规备份个数，默认值为1
	BackupPeriodSaveCount *int64 `json:"BackupPeriodSaveCount,omitnil,omitempty" name:"BackupPeriodSaveCount"`

	// 定期保留策略周期起始日期，格式：YYYY-MM-dd HH:mm:ss
	StartBackupPeriodSaveDate *string `json:"StartBackupPeriodSaveDate,omitnil,omitempty" name:"StartBackupPeriodSaveDate"`

	// 是否开启数据备份归档策略，off-关闭，on-打开，如果不指定该入参， 则保持不变。
	EnableBackupArchive *string `json:"EnableBackupArchive,omitnil,omitempty" name:"EnableBackupArchive"`

	// 数据备份归档起始天数，数据备份达到归档起始天数时进行归档，最小为180天，不得大于数据备份保留天数
	BackupArchiveDays *int64 `json:"BackupArchiveDays,omitnil,omitempty" name:"BackupArchiveDays"`

	// 日志备份归档起始天数，日志备份达到归档起始天数时进行归档，最小为180天，不得大于日志备份保留天数
	BinlogArchiveDays *int64 `json:"BinlogArchiveDays,omitnil,omitempty" name:"BinlogArchiveDays"`

	// 是否开启日志备份归档策略，off-关闭，on-打开，如果不指定该入参， 则保持不变。
	EnableBinlogArchive *string `json:"EnableBinlogArchive,omitnil,omitempty" name:"EnableBinlogArchive"`

	// 是否开启数据备份标准存储策略，off-关闭，on-打开，如果不指定该入参， 则保持不变。
	EnableBackupStandby *string `json:"EnableBackupStandby,omitnil,omitempty" name:"EnableBackupStandby"`

	// 数据备份标准存储起始天数，数据备份达到标准存储起始天数时进行转换，最小为30天，不得大于数据备份保留天数。如果开启备份归档，不得大于等于备份归档天数
	BackupStandbyDays *int64 `json:"BackupStandbyDays,omitnil,omitempty" name:"BackupStandbyDays"`

	// 是否开启日志备份标准存储策略，off-关闭，on-打开，如果不指定该入参， 则保持不变。
	EnableBinlogStandby *string `json:"EnableBinlogStandby,omitnil,omitempty" name:"EnableBinlogStandby"`

	// 日志备份标准存储起始天数，日志备份达到标准存储起始天数时进行转换，最小为30天，不得大于日志备份保留天数。如果开启备份归档，不得大于等于备份归档天数
	BinlogStandbyDays *int64 `json:"BinlogStandbyDays,omitnil,omitempty" name:"BinlogStandbyDays"`
}

type ModifyBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据备份文件的保留时间，单位为天。
	// 1. MySQL 双节点、三节点、云盘版数据备份文件可以保留7天 - 1830天。
	// 2. MySQL 单节点（云盘）数据备份文件可以保留7天 - 30天。
	ExpireDays *int64 `json:"ExpireDays,omitnil,omitempty" name:"ExpireDays"`

	// (将废弃，建议使用 BackupTimeWindow 参数) 备份时间范围，格式为：02:00-06:00，起点和终点时间目前限制为整点，目前可以选择的范围为： 00:00-12:00，02:00-06:00，06：00-10：00，10:00-14:00，14:00-18:00，18:00-22:00，22:00-02:00。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 自动备份方式，仅支持：physical - 物理冷备
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// binlog 的保留时间，单位为天。
	// 1. MySQL 双节点、三节点、云盘版日志备份文件可以保留7天 - 3650天。
	// 2. MySQL 单节点（云盘）日志备份文件可以保留7天 - 30天。
	BinlogExpireDays *int64 `json:"BinlogExpireDays,omitnil,omitempty" name:"BinlogExpireDays"`

	// 备份时间窗，比如要设置每周二和周日 10:00-14:00之间备份，该参数如下：{"Monday": "", "Tuesday": "10:00-14:00", "Wednesday": "", "Thursday": "", "Friday": "", "Saturday": "", "Sunday": "10:00-14:00"}    （注：可以设置一周的某几天备份，但是每天的备份时间需要设置为相同的时间段。 如果设置了该字段，将忽略StartTime字段的设置）
	BackupTimeWindow *CommonTimeWindow `json:"BackupTimeWindow,omitnil,omitempty" name:"BackupTimeWindow"`

	// 定期保留开关，off - 不开启定期保留策略，on - 开启定期保留策略，默认为off。
	EnableBackupPeriodSave *string `json:"EnableBackupPeriodSave,omitnil,omitempty" name:"EnableBackupPeriodSave"`

	// 长期保留开关,该字段功能暂未上线，可忽略。off - 不开启长期保留策略，on - 开启长期保留策略，默认为off，如果开启，则BackupPeriodSaveDays，BackupPeriodSaveInterval，BackupPeriodSaveCount参数无效
	EnableBackupPeriodLongTermSave *string `json:"EnableBackupPeriodLongTermSave,omitnil,omitempty" name:"EnableBackupPeriodLongTermSave"`

	// 定期保留最长天数，最小值：90，最大值：3650，默认值：1080
	BackupPeriodSaveDays *int64 `json:"BackupPeriodSaveDays,omitnil,omitempty" name:"BackupPeriodSaveDays"`

	// 定期保留策略周期，可取值：weekly - 周，monthly - 月， quarterly - 季度，yearly - 年，默认为monthly
	BackupPeriodSaveInterval *string `json:"BackupPeriodSaveInterval,omitnil,omitempty" name:"BackupPeriodSaveInterval"`

	// 定期保留的备份数量，最小值为1，最大值不超过定期保留策略周期内常规备份个数，默认值为1
	BackupPeriodSaveCount *int64 `json:"BackupPeriodSaveCount,omitnil,omitempty" name:"BackupPeriodSaveCount"`

	// 定期保留策略周期起始日期，格式：YYYY-MM-dd HH:mm:ss
	StartBackupPeriodSaveDate *string `json:"StartBackupPeriodSaveDate,omitnil,omitempty" name:"StartBackupPeriodSaveDate"`

	// 是否开启数据备份归档策略，off-关闭，on-打开，如果不指定该入参， 则保持不变。
	EnableBackupArchive *string `json:"EnableBackupArchive,omitnil,omitempty" name:"EnableBackupArchive"`

	// 数据备份归档起始天数，数据备份达到归档起始天数时进行归档，最小为180天，不得大于数据备份保留天数
	BackupArchiveDays *int64 `json:"BackupArchiveDays,omitnil,omitempty" name:"BackupArchiveDays"`

	// 日志备份归档起始天数，日志备份达到归档起始天数时进行归档，最小为180天，不得大于日志备份保留天数
	BinlogArchiveDays *int64 `json:"BinlogArchiveDays,omitnil,omitempty" name:"BinlogArchiveDays"`

	// 是否开启日志备份归档策略，off-关闭，on-打开，如果不指定该入参， 则保持不变。
	EnableBinlogArchive *string `json:"EnableBinlogArchive,omitnil,omitempty" name:"EnableBinlogArchive"`

	// 是否开启数据备份标准存储策略，off-关闭，on-打开，如果不指定该入参， 则保持不变。
	EnableBackupStandby *string `json:"EnableBackupStandby,omitnil,omitempty" name:"EnableBackupStandby"`

	// 数据备份标准存储起始天数，数据备份达到标准存储起始天数时进行转换，最小为30天，不得大于数据备份保留天数。如果开启备份归档，不得大于等于备份归档天数
	BackupStandbyDays *int64 `json:"BackupStandbyDays,omitnil,omitempty" name:"BackupStandbyDays"`

	// 是否开启日志备份标准存储策略，off-关闭，on-打开，如果不指定该入参， 则保持不变。
	EnableBinlogStandby *string `json:"EnableBinlogStandby,omitnil,omitempty" name:"EnableBinlogStandby"`

	// 日志备份标准存储起始天数，日志备份达到标准存储起始天数时进行转换，最小为30天，不得大于日志备份保留天数。如果开启备份归档，不得大于等于备份归档天数
	BinlogStandbyDays *int64 `json:"BinlogStandbyDays,omitnil,omitempty" name:"BinlogStandbyDays"`
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
	// NoLimit 不限制,内外网都可以下载； LimitOnlyIntranet 仅内网可下载； Customize 用户自定义vpc:ip可下载。 只有该值为 Customize 时，才可以设置 LimitVpc 和 LimitIp 。
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// 该参数仅支持 In， 表示 LimitVpc 指定的vpc可以下载。默认为In。
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// In: 指定的ip可以下载； NotIn: 指定的ip不可以下载。 默认为In。
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// 限制下载的vpc设置。
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil,omitempty" name:"LimitVpc"`

	// 限制下载的ip设置
	LimitIp []*string `json:"LimitIp,omitnil,omitempty" name:"LimitIp"`
}

type ModifyBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
	
	// NoLimit 不限制,内外网都可以下载； LimitOnlyIntranet 仅内网可下载； Customize 用户自定义vpc:ip可下载。 只有该值为 Customize 时，才可以设置 LimitVpc 和 LimitIp 。
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// 该参数仅支持 In， 表示 LimitVpc 指定的vpc可以下载。默认为In。
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// In: 指定的ip可以下载； NotIn: 指定的ip不可以下载。 默认为In。
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// 限制下载的vpc设置。
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil,omitempty" name:"LimitVpc"`

	// 限制下载的ip设置
	LimitIp []*string `json:"LimitIp,omitnil,omitempty" name:"LimitIp"`
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
type ModifyBackupEncryptionStatusRequestParams struct {
	// 实例ID，格式如：cdb-XXXX。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设置实例新增的自动物理备份文件默认加密状态。可选值为 on或者off。
	EncryptionStatus *string `json:"EncryptionStatus,omitnil,omitempty" name:"EncryptionStatus"`
}

type ModifyBackupEncryptionStatusRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cdb-XXXX。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设置实例新增的自动物理备份文件默认加密状态。可选值为 on或者off。
	EncryptionStatus *string `json:"EncryptionStatus,omitnil,omitempty" name:"EncryptionStatus"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyCdbProxyAddressDescRequestParams struct {
	// 代理组 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 代理组地址 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`

	// 描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type ModifyCdbProxyAddressDescRequest struct {
	*tchttp.BaseRequest
	
	// 代理组 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 代理组地址 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`

	// 描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

func (r *ModifyCdbProxyAddressDescRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCdbProxyAddressDescRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyGroupId")
	delete(f, "ProxyAddressId")
	delete(f, "Desc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCdbProxyAddressDescRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCdbProxyAddressDescResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCdbProxyAddressDescResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCdbProxyAddressDescResponseParams `json:"Response"`
}

func (r *ModifyCdbProxyAddressDescResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCdbProxyAddressDescResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCdbProxyAddressVipAndVPortRequestParams struct {
	// 代理组 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 代理组地址 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`

	// 私有网络 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 私有子网 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// IP 地址。若不填写则自动分配子网下的可用 IP。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 端口。默认值3306，取值范围：1024 - 65535。
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// 旧 IP 地址回收时间。单位：小时，默认值：24，取值范围：0 - 168。
	ReleaseDuration *uint64 `json:"ReleaseDuration,omitnil,omitempty" name:"ReleaseDuration"`
}

type ModifyCdbProxyAddressVipAndVPortRequest struct {
	*tchttp.BaseRequest
	
	// 代理组 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 代理组地址 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`

	// 私有网络 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 私有子网 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// IP 地址。若不填写则自动分配子网下的可用 IP。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 端口。默认值3306，取值范围：1024 - 65535。
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// 旧 IP 地址回收时间。单位：小时，默认值：24，取值范围：0 - 168。
	ReleaseDuration *uint64 `json:"ReleaseDuration,omitnil,omitempty" name:"ReleaseDuration"`
}

func (r *ModifyCdbProxyAddressVipAndVPortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCdbProxyAddressVipAndVPortRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyGroupId")
	delete(f, "ProxyAddressId")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "Vip")
	delete(f, "VPort")
	delete(f, "ReleaseDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCdbProxyAddressVipAndVPortRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCdbProxyAddressVipAndVPortResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCdbProxyAddressVipAndVPortResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCdbProxyAddressVipAndVPortResponseParams `json:"Response"`
}

func (r *ModifyCdbProxyAddressVipAndVPortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCdbProxyAddressVipAndVPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCdbProxyParamRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 代理组 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 连接池阈值。取值范围：大于0，小于等于300。
	// 注意：如需使用数据库代理连接池能力，MySQL 8.0 主实例的内核小版本要大于等于 MySQL 8.0 20230630。
	ConnectionPoolLimit *uint64 `json:"ConnectionPoolLimit,omitnil,omitempty" name:"ConnectionPoolLimit"`
}

type ModifyCdbProxyParamRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 代理组 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 连接池阈值。取值范围：大于0，小于等于300。
	// 注意：如需使用数据库代理连接池能力，MySQL 8.0 主实例的内核小版本要大于等于 MySQL 8.0 20230630。
	ConnectionPoolLimit *uint64 `json:"ConnectionPoolLimit,omitnil,omitempty" name:"ConnectionPoolLimit"`
}

func (r *ModifyCdbProxyParamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCdbProxyParamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProxyGroupId")
	delete(f, "ConnectionPoolLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCdbProxyParamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCdbProxyParamResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCdbProxyParamResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCdbProxyParamResponseParams `json:"Response"`
}

func (r *ModifyCdbProxyParamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCdbProxyParamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceLogToCLSRequestParams struct {
	// 实例 ID，可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志类型。error：错误日志，slowlog：慢日志。
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 投递状态。ON：开启，OFF：关闭。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否需要创建日志集。默认为 false。
	CreateLogset *bool `json:"CreateLogset,omitnil,omitempty" name:"CreateLogset"`

	// 需要创建日志集时为日志集名称；选择已有日志集时，为日志集 ID。默认为空。
	// 说明：当参数 Status 的值为 ON 时，Logset 和 LogTopic 参数必须填一个。
	Logset *string `json:"Logset,omitnil,omitempty" name:"Logset"`

	// 是否需要创建日志主题。默认为 false。
	CreateLogTopic *bool `json:"CreateLogTopic,omitnil,omitempty" name:"CreateLogTopic"`

	// 需要创建日志主题时为日志主题名称；选择已有日志主题时，为日志主题 ID。默认为空。
	// 说明：当参数 Status 的值为 ON 时，Logset 和 LogTopic 参数必须填一个。
	LogTopic *string `json:"LogTopic,omitnil,omitempty" name:"LogTopic"`

	// 日志主题有效期，不填写时，默认30天，最大值3600。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 创建日志主题时，是否创建索引，默认为 false。
	CreateIndex *bool `json:"CreateIndex,omitnil,omitempty" name:"CreateIndex"`

	// CLS 所在地域，不填择默认为 Region 的参数值。
	ClsRegion *string `json:"ClsRegion,omitnil,omitempty" name:"ClsRegion"`
}

type ModifyDBInstanceLogToCLSRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志类型。error：错误日志，slowlog：慢日志。
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 投递状态。ON：开启，OFF：关闭。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否需要创建日志集。默认为 false。
	CreateLogset *bool `json:"CreateLogset,omitnil,omitempty" name:"CreateLogset"`

	// 需要创建日志集时为日志集名称；选择已有日志集时，为日志集 ID。默认为空。
	// 说明：当参数 Status 的值为 ON 时，Logset 和 LogTopic 参数必须填一个。
	Logset *string `json:"Logset,omitnil,omitempty" name:"Logset"`

	// 是否需要创建日志主题。默认为 false。
	CreateLogTopic *bool `json:"CreateLogTopic,omitnil,omitempty" name:"CreateLogTopic"`

	// 需要创建日志主题时为日志主题名称；选择已有日志主题时，为日志主题 ID。默认为空。
	// 说明：当参数 Status 的值为 ON 时，Logset 和 LogTopic 参数必须填一个。
	LogTopic *string `json:"LogTopic,omitnil,omitempty" name:"LogTopic"`

	// 日志主题有效期，不填写时，默认30天，最大值3600。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 创建日志主题时，是否创建索引，默认为 false。
	CreateIndex *bool `json:"CreateIndex,omitnil,omitempty" name:"CreateIndex"`

	// CLS 所在地域，不填择默认为 Region 的参数值。
	ClsRegion *string `json:"ClsRegion,omitnil,omitempty" name:"ClsRegion"`
}

func (r *ModifyDBInstanceLogToCLSRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceLogToCLSRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogType")
	delete(f, "Status")
	delete(f, "CreateLogset")
	delete(f, "Logset")
	delete(f, "CreateLogTopic")
	delete(f, "LogTopic")
	delete(f, "Period")
	delete(f, "CreateIndex")
	delete(f, "ClsRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceLogToCLSRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceLogToCLSResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceLogToCLSResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceLogToCLSResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceLogToCLSResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceLogToCLSResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNameRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改后的实例名称，仅支持数字,英文大小写字母、中文以及特殊字符-_./()[]（）+=:：@ 且长度不能超过60。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改后的实例名称，仅支持数字,英文大小写字母、中文以及特殊字符-_./()[]（）+=:：@ 且长度不能超过60。
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
	delete(f, "InstanceId")
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
type ModifyDBInstanceProjectRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	// 说明：可输入多个实例 ID 进行修改，json 格式如下。
	// [
	//     "cdb-30z11v8s",
	//     "cdb-93h11efg"
	//   ]
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例所属项目的 ID，可在账号中心下的项目管理页面查询。
	// 说明：此项为必填。
	NewProjectId *int64 `json:"NewProjectId,omitnil,omitempty" name:"NewProjectId"`
}

type ModifyDBInstanceProjectRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	// 说明：可输入多个实例 ID 进行修改，json 格式如下。
	// [
	//     "cdb-30z11v8s",
	//     "cdb-93h11efg"
	//   ]
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例所属项目的 ID，可在账号中心下的项目管理页面查询。
	// 说明：此项为必填。
	NewProjectId *int64 `json:"NewProjectId,omitnil,omitempty" name:"NewProjectId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyDBInstanceReadOnlyStatusRequestParams struct {
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，可使用[查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否设置为只读。其中：1表示设置实例为只读，0表示解除只读状态
	ReadOnly *int64 `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`
}

type ModifyDBInstanceReadOnlyStatusRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，可使用[查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否设置为只读。其中：1表示设置实例为只读，0表示解除只读状态
	ReadOnly *int64 `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`
}

func (r *ModifyDBInstanceReadOnlyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceReadOnlyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ReadOnly")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceReadOnlyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceReadOnlyStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceReadOnlyStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceReadOnlyStatusResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceReadOnlyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceReadOnlyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要修改的安全组 ID 列表，一个或者多个安全组 ID 组成的数组。可通过 [DescribeDBSecurityGroups](https://cloud.tencent.com/document/product/236/15854) 接口获取。输入的安全组 ID 数组无长度限制。
	// 注意：该入参会全量替换存量已有集合，非增量更新。修改需传入预期的全量集合。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 当传入只读实例 ID 时，默认操作的是对应只读组的安全组。如果需要操作只读实例 ID 的安全组， 需要将该入参置为 True。默认为 False。
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`

	// 变更集群版实例只读组时，InstanceId 传实例 ID，需要额外指定该参数表示操作只读组。 如果操作读写节点则不需指定该参数。
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要修改的安全组 ID 列表，一个或者多个安全组 ID 组成的数组。可通过 [DescribeDBSecurityGroups](https://cloud.tencent.com/document/product/236/15854) 接口获取。输入的安全组 ID 数组无长度限制。
	// 注意：该入参会全量替换存量已有集合，非增量更新。修改需传入预期的全量集合。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 当传入只读实例 ID 时，默认操作的是对应只读组的安全组。如果需要操作只读实例 ID 的安全组， 需要将该入参置为 True。默认为 False。
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`

	// 变更集群版实例只读组时，InstanceId 传实例 ID，需要额外指定该参数表示操作只读组。 如果操作读写节点则不需指定该参数。
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
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
	delete(f, "OpResourceId")
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
type ModifyDBInstanceVipVportRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c2nl9rpv 或者 cdbrg-c3nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 目标 IP。
	DstIp *string `json:"DstIp,omitnil,omitempty" name:"DstIp"`

	// 目标端口，支持范围为：[1024-65535]。
	DstPort *int64 `json:"DstPort,omitnil,omitempty" name:"DstPort"`

	// 私有网络统一 ID。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 子网统一 ID。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 进行基础网络转 VPC 网络和 VPC 网络下的子网变更时，原网络中旧IP的回收时间，单位为小时，取值范围为0-168，默认值为24小时。
	ReleaseDuration *int64 `json:"ReleaseDuration,omitnil,omitempty" name:"ReleaseDuration"`

	// 变更集群版实例只读组时，InstanceId传实例id，需要额外指定该参数表示操作只读组。 如果操作读写节点则不需指定该参数。
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
}

type ModifyDBInstanceVipVportRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c2nl9rpv 或者 cdbrg-c3nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 目标 IP。
	DstIp *string `json:"DstIp,omitnil,omitempty" name:"DstIp"`

	// 目标端口，支持范围为：[1024-65535]。
	DstPort *int64 `json:"DstPort,omitnil,omitempty" name:"DstPort"`

	// 私有网络统一 ID。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 子网统一 ID。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 进行基础网络转 VPC 网络和 VPC 网络下的子网变更时，原网络中旧IP的回收时间，单位为小时，取值范围为0-168，默认值为24小时。
	ReleaseDuration *int64 `json:"ReleaseDuration,omitnil,omitempty" name:"ReleaseDuration"`

	// 变更集群版实例只读组时，InstanceId传实例id，需要额外指定该参数表示操作只读组。 如果操作读写节点则不需指定该参数。
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
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
	delete(f, "OpResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceVipVportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceVipVportResponseParams struct {
	// 异步任务ID。(该返回字段目前已废弃)
	//
	// Deprecated: AsyncRequestId is deprecated.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例 ID 列表。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 要修改的参数列表。每一个元素是 Name 和 CurrentValue 的组合。Name 是参数名，CurrentValue 是要修改成的值。
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// 模板 ID，ParamList 和 TemplateId 必须至少传其中之一。可通过 [DescribeParamTemplates](https://cloud.tencent.com/document/api/236/32659) 接口获取。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 执行参数调整任务的方式，默认为 0。支持值包括：0 - 立刻执行，1 - 时间窗执行；当该值为 1 时，每次只能传一个实例（InstanceIds数量为1）
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// 参数是否同步到主实例下的只读实例。true 为不同步，false 为同步。默认为 false。
	NotSyncRo *bool `json:"NotSyncRo,omitnil,omitempty" name:"NotSyncRo"`

	// 参数是否同步到主实例下的灾备实例。true 为不同步，false 为同步。默认为 false。
	NotSyncDr *bool `json:"NotSyncDr,omitnil,omitempty" name:"NotSyncDr"`
}

type ModifyInstanceParamRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 要修改的参数列表。每一个元素是 Name 和 CurrentValue 的组合。Name 是参数名，CurrentValue 是要修改成的值。
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// 模板 ID，ParamList 和 TemplateId 必须至少传其中之一。可通过 [DescribeParamTemplates](https://cloud.tencent.com/document/api/236/32659) 接口获取。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 执行参数调整任务的方式，默认为 0。支持值包括：0 - 立刻执行，1 - 时间窗执行；当该值为 1 时，每次只能传一个实例（InstanceIds数量为1）
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// 参数是否同步到主实例下的只读实例。true 为不同步，false 为同步。默认为 false。
	NotSyncRo *bool `json:"NotSyncRo,omitnil,omitempty" name:"NotSyncRo"`

	// 参数是否同步到主实例下的灾备实例。true 为不同步，false 为同步。默认为 false。
	NotSyncDr *bool `json:"NotSyncDr,omitnil,omitempty" name:"NotSyncDr"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

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
type ModifyInstancePasswordComplexityRequestParams struct {
	// 要修改密码复杂度的实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	// 说明：支持输入多个实例 ID 进行修改。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 要修改的密码复杂度的选项。每一个选项是以组合形式写入的，一个组合包括 Name 和 CurrentValue，其中 Name 表示对应选项的参数名，CurrentValue 表示参数值。例如：[{"Name": "validate_password.length", "CurrentValue": "10"}]，表示将密码的最小字符数修改为10。
	// 说明：不同数据库版本的实例，支持修改的密码复杂度的选项如下。
	// 1. MySQL 8.0：
	// 选项 validate_password.policy，表示密码复杂度的开关，值为 LOW 时表示关闭；值为 MEDIUM 时表示开启。
	// 选项 validate_password.length，表示密码总长度的最小字符数。
	// 选项 validate_password.mixed_case_count，表示小写和大写字母的最小字符数。
	// 选项 validate_password.number_count，表示数字的最小字符数。
	// 选项 validate_password.special_char_count，表示特殊字符的最小字符数。
	// 2. MySQL 5.6、MySQL 5.7：
	// 选项 validate_password_policy，表示密码复杂度的开关，值为 LOW 时表示关闭；值为 MEDIUM 时表示开启。
	// 选项 validate_password_length，表示密码总长度的最小字符数。
	// 选项 validate_password_mixed_case_count，表示小写和大写字母的最小字符数。
	// 选项 validate_password_number_count，表示数字的最小字符数。
	// 选项 validate_password_special_char_count，表示特殊字符的最小字符数。
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type ModifyInstancePasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// 要修改密码复杂度的实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	// 说明：支持输入多个实例 ID 进行修改。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 要修改的密码复杂度的选项。每一个选项是以组合形式写入的，一个组合包括 Name 和 CurrentValue，其中 Name 表示对应选项的参数名，CurrentValue 表示参数值。例如：[{"Name": "validate_password.length", "CurrentValue": "10"}]，表示将密码的最小字符数修改为10。
	// 说明：不同数据库版本的实例，支持修改的密码复杂度的选项如下。
	// 1. MySQL 8.0：
	// 选项 validate_password.policy，表示密码复杂度的开关，值为 LOW 时表示关闭；值为 MEDIUM 时表示开启。
	// 选项 validate_password.length，表示密码总长度的最小字符数。
	// 选项 validate_password.mixed_case_count，表示小写和大写字母的最小字符数。
	// 选项 validate_password.number_count，表示数字的最小字符数。
	// 选项 validate_password.special_char_count，表示特殊字符的最小字符数。
	// 2. MySQL 5.6、MySQL 5.7：
	// 选项 validate_password_policy，表示密码复杂度的开关，值为 LOW 时表示关闭；值为 MEDIUM 时表示开启。
	// 选项 validate_password_length，表示密码总长度的最小字符数。
	// 选项 validate_password_mixed_case_count，表示小写和大写字母的最小字符数。
	// 选项 validate_password_number_count，表示数字的最小字符数。
	// 选项 validate_password_special_char_count，表示特殊字符的最小字符数。
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要增加或修改的标签。ReplaceTags 和 DeleteTags 必填一个。
	ReplaceTags []*TagInfo `json:"ReplaceTags,omitnil,omitempty" name:"ReplaceTags"`

	// 要删除的标签。ReplaceTags 和 DeleteTags 必填一个。
	DeleteTags []*TagInfo `json:"DeleteTags,omitnil,omitempty" name:"DeleteTags"`
}

type ModifyInstanceTagRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要增加或修改的标签。ReplaceTags 和 DeleteTags 必填一个。
	ReplaceTags []*TagInfo `json:"ReplaceTags,omitnil,omitempty" name:"ReplaceTags"`

	// 要删除的标签。ReplaceTags 和 DeleteTags 必填一个。
	DeleteTags []*TagInfo `json:"DeleteTags,omitnil,omitempty" name:"DeleteTags"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 本地 binlog 保留时长。不同实例的可取值如下：
	// 1. 云盘版实例、双节点实例、三节点实例的本地 binlog 保留时长（小时）默认为120，范围：6 - 168。
	// 2. 灾备实例的本地 binlog 保留时长（小时）默认为120，范围：120 - 168。
	// 3. 单节点云盘实例的本地 binlog 保留时长（小时）默认为120，范围：0 - 168。
	// 4. 若双节点实例、三节点实例下无灾备实例，则该主实例的本地 binlog 保留时长（小时）范围是：6 - 168；若双节点实例、三节点实例下有灾备实例，或者要为双节点实例、三节点实例添加灾备实例，为避免同步异常，该主实例的本地 binlog 保留时长（小时）不能设置低于120小时，范围是：120 - 168。
	SaveHours *int64 `json:"SaveHours,omitnil,omitempty" name:"SaveHours"`

	// 本地 binlog 空间使用率，可取值范围：[30,50]。
	MaxUsage *int64 `json:"MaxUsage,omitnil,omitempty" name:"MaxUsage"`
}

type ModifyLocalBinlogConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 本地 binlog 保留时长。不同实例的可取值如下：
	// 1. 云盘版实例、双节点实例、三节点实例的本地 binlog 保留时长（小时）默认为120，范围：6 - 168。
	// 2. 灾备实例的本地 binlog 保留时长（小时）默认为120，范围：120 - 168。
	// 3. 单节点云盘实例的本地 binlog 保留时长（小时）默认为120，范围：0 - 168。
	// 4. 若双节点实例、三节点实例下无灾备实例，则该主实例的本地 binlog 保留时长（小时）范围是：6 - 168；若双节点实例、三节点实例下有灾备实例，或者要为双节点实例、三节点实例添加灾备实例，为避免同步异常，该主实例的本地 binlog 保留时长（小时）不能设置低于120小时，范围是：120 - 168。
	SaveHours *int64 `json:"SaveHours,omitnil,omitempty" name:"SaveHours"`

	// 本地 binlog 空间使用率，可取值范围：[30,50]。
	MaxUsage *int64 `json:"MaxUsage,omitnil,omitempty" name:"MaxUsage"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// 置放群组名称，最长不能超过60个字符。置放群组名和置放群组描述不能都为空。
	DeployGroupName *string `json:"DeployGroupName,omitnil,omitempty" name:"DeployGroupName"`

	// 置放群组描述，最长不能超过200个字符。置放群组名和置放群组描述不能都为空。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyNameOrDescByDpIdRequest struct {
	*tchttp.BaseRequest
	
	// 置放群组 ID。
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// 置放群组名称，最长不能超过60个字符。置放群组名和置放群组描述不能都为空。
	DeployGroupName *string `json:"DeployGroupName,omitnil,omitempty" name:"DeployGroupName"`

	// 置放群组描述，最长不能超过200个字符。置放群组名和置放群组描述不能都为空。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 模板 ID。可通过 [DescribeParamTemplates](https://cloud.tencent.com/document/api/236/32659) 接口获取。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名称，仅支持数字、英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@,且长度不能超过60。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述，长度不超过255。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 参数列表。
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type ModifyParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板 ID。可通过 [DescribeParamTemplates](https://cloud.tencent.com/document/api/236/32659) 接口获取。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名称，仅支持数字、英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@,且长度不能超过60。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述，长度不超过255。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 参数列表。
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`
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

// Predefined struct for user
type ModifyProtectModeRequestParams struct {
	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制。
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ModifyProtectModeRequest struct {
	*tchttp.BaseRequest
	
	// 数据复制方式，默认为 0，支持值包括：0 - 表示异步复制，1 - 表示半同步复制，2 - 表示强同步复制。
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *ModifyProtectModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProtectModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProtectMode")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProtectModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProtectModeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProtectModeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProtectModeResponseParams `json:"Response"`
}

func (r *ModifyProtectModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProtectModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRemoteBackupConfigRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 异地数据备份开关，off - 关闭异地备份，on-开启异地备份
	RemoteBackupSave *string `json:"RemoteBackupSave,omitnil,omitempty" name:"RemoteBackupSave"`

	// 异地日志备份开关，off - 关闭异地备份，on-开启异地备份，只有在参数RemoteBackupSave为on时，RemoteBinlogSave参数才可设置为on
	RemoteBinlogSave *string `json:"RemoteBinlogSave,omitnil,omitempty" name:"RemoteBinlogSave"`

	// 用户设置异地备份地域列表
	RemoteRegion []*string `json:"RemoteRegion,omitnil,omitempty" name:"RemoteRegion"`

	// 异地备份保留时间，单位为天
	ExpireDays *int64 `json:"ExpireDays,omitnil,omitempty" name:"ExpireDays"`
}

type ModifyRemoteBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 异地数据备份开关，off - 关闭异地备份，on-开启异地备份
	RemoteBackupSave *string `json:"RemoteBackupSave,omitnil,omitempty" name:"RemoteBackupSave"`

	// 异地日志备份开关，off - 关闭异地备份，on-开启异地备份，只有在参数RemoteBackupSave为on时，RemoteBinlogSave参数才可设置为on
	RemoteBinlogSave *string `json:"RemoteBinlogSave,omitnil,omitempty" name:"RemoteBinlogSave"`

	// 用户设置异地备份地域列表
	RemoteRegion []*string `json:"RemoteRegion,omitnil,omitempty" name:"RemoteRegion"`

	// 异地备份保留时间，单位为天
	ExpireDays *int64 `json:"ExpireDays,omitnil,omitempty" name:"ExpireDays"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// RO 组的 ID。可通过 [DescribeRoGroups](https://cloud.tencent.com/document/api/236/40939) 接口获取。
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`

	// RO 组的详细信息。
	RoGroupInfo *RoGroupAttr `json:"RoGroupInfo,omitnil,omitempty" name:"RoGroupInfo"`

	// RO 组内实例的权重。若修改 RO 组的权重模式为用户自定义模式（custom），则必须设置该参数，且需要设置每个 RO 实例的权重值。RO 实例 ID 可通过 [DescribeRoGroups](https://cloud.tencent.com/document/api/236/40939) 接口获取。
	RoWeightValues []*RoWeightValue `json:"RoWeightValues,omitnil,omitempty" name:"RoWeightValues"`

	// 是否重新均衡 RO 组内的 RO 实例的负载。支持值包括：1 - 重新均衡负载；0 - 不重新均衡负载。默认值为 0。注意，设置为重新均衡负载时，RO 组内 RO 实例会有一次数据库连接瞬断，请确保应用程序能重连数据库。
	IsBalanceRoLoad *int64 `json:"IsBalanceRoLoad,omitnil,omitempty" name:"IsBalanceRoLoad"`

	// 废弃参数，无意义。
	//
	// Deprecated: ReplicationDelayTime is deprecated.
	ReplicationDelayTime *int64 `json:"ReplicationDelayTime,omitnil,omitempty" name:"ReplicationDelayTime"`
}

type ModifyRoGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// RO 组的 ID。可通过 [DescribeRoGroups](https://cloud.tencent.com/document/api/236/40939) 接口获取。
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`

	// RO 组的详细信息。
	RoGroupInfo *RoGroupAttr `json:"RoGroupInfo,omitnil,omitempty" name:"RoGroupInfo"`

	// RO 组内实例的权重。若修改 RO 组的权重模式为用户自定义模式（custom），则必须设置该参数，且需要设置每个 RO 实例的权重值。RO 实例 ID 可通过 [DescribeRoGroups](https://cloud.tencent.com/document/api/236/40939) 接口获取。
	RoWeightValues []*RoWeightValue `json:"RoWeightValues,omitnil,omitempty" name:"RoWeightValues"`

	// 是否重新均衡 RO 组内的 RO 实例的负载。支持值包括：1 - 重新均衡负载；0 - 不重新均衡负载。默认值为 0。注意，设置为重新均衡负载时，RO 组内 RO 实例会有一次数据库连接瞬断，请确保应用程序能重连数据库。
	IsBalanceRoLoad *int64 `json:"IsBalanceRoLoad,omitnil,omitempty" name:"IsBalanceRoLoad"`

	// 废弃参数，无意义。
	ReplicationDelayTime *int64 `json:"ReplicationDelayTime,omitnil,omitempty" name:"ReplicationDelayTime"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyRoGroupVipVportRequestParams struct {
	// RO组的ID。
	UGroupId *string `json:"UGroupId,omitnil,omitempty" name:"UGroupId"`

	// 目标IP。
	DstIp *string `json:"DstIp,omitnil,omitempty" name:"DstIp"`

	// 目标Port。
	DstPort *int64 `json:"DstPort,omitnil,omitempty" name:"DstPort"`
}

type ModifyRoGroupVipVportRequest struct {
	*tchttp.BaseRequest
	
	// RO组的ID。
	UGroupId *string `json:"UGroupId,omitnil,omitempty" name:"UGroupId"`

	// 目标IP。
	DstIp *string `json:"DstIp,omitnil,omitempty" name:"DstIp"`

	// 目标Port。
	DstPort *int64 `json:"DstPort,omitnil,omitempty" name:"DstPort"`
}

func (r *ModifyRoGroupVipVportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoGroupVipVportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UGroupId")
	delete(f, "DstIp")
	delete(f, "DstPort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoGroupVipVportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoGroupVipVportResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRoGroupVipVportResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRoGroupVipVportResponseParams `json:"Response"`
}

func (r *ModifyRoGroupVipVportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoGroupVipVportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTimeWindowRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改后的可维护时间段，其中每一个时间段的格式形如：10:00-12:00；起止时间按半个小时对齐；最短半个小时，最长三个小时；最多设置两个时间段；起止时间范围为：[00:00, 24:00]。
	// 说明：设置两个时间段的 json 示例如下。
	// [
	//     "01:00-01:30",
	//     "02:00-02:30"
	//   ]
	TimeRanges []*string `json:"TimeRanges,omitnil,omitempty" name:"TimeRanges"`

	// 指定修改哪一天的可维护时间段，可能的取值为：monday，tuesday，wednesday，thursday，friday，saturday，sunday。如果不指定该值或者为空，则默认一周七天都修改。
	// 说明：指定修改多天的 json 示例如下。
	// [
	//     "monday",
	//     "tuesday"
	//   ]
	Weekdays []*string `json:"Weekdays,omitnil,omitempty" name:"Weekdays"`

	// 数据延迟阈值（秒），仅对主实例和灾备实例有效。不传默认不修改，保持原来的阈值，取值范围：1-10的整数。
	MaxDelayTime *uint64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`
}

type ModifyTimeWindowRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改后的可维护时间段，其中每一个时间段的格式形如：10:00-12:00；起止时间按半个小时对齐；最短半个小时，最长三个小时；最多设置两个时间段；起止时间范围为：[00:00, 24:00]。
	// 说明：设置两个时间段的 json 示例如下。
	// [
	//     "01:00-01:30",
	//     "02:00-02:30"
	//   ]
	TimeRanges []*string `json:"TimeRanges,omitnil,omitempty" name:"TimeRanges"`

	// 指定修改哪一天的可维护时间段，可能的取值为：monday，tuesday，wednesday，thursday，friday，saturday，sunday。如果不指定该值或者为空，则默认一周七天都修改。
	// 说明：指定修改多天的 json 示例如下。
	// [
	//     "monday",
	//     "tuesday"
	//   ]
	Weekdays []*string `json:"Weekdays,omitnil,omitempty" name:"Weekdays"`

	// 数据延迟阈值（秒），仅对主实例和灾备实例有效。不传默认不修改，保持原来的阈值，取值范围：1-10的整数。
	MaxDelayTime *uint64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type NodeDistribution struct {
	// 主实例Master节点所在主机ID或者只读实例所在主机ID
	Node *string `json:"Node,omitnil,omitempty" name:"Node"`

	// 主实例第一Slave节点所在主机ID
	SlaveNodeOne *string `json:"SlaveNodeOne,omitnil,omitempty" name:"SlaveNodeOne"`

	// 主实例第二Slave节点所在主机ID
	SlaveNodeTwo *string `json:"SlaveNodeTwo,omitnil,omitempty" name:"SlaveNodeTwo"`
}

// Predefined struct for user
type OfflineIsolatedInstancesRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type OfflineIsolatedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// CDB 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计日志保存时长。支持值包括：
	// 7 - 一周；
	// 30 - 一个月；
	// 90 - 三个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年。
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 高频审计日志保存时长。默认值为7，此项取值需小于等于 LogExpireDay，支持值包括：
	// 3 - 3天；
	// 7 - 一周；
	// 30 - 一个月；
	// 90 - 三个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年。
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// 审计规则（该参数已废弃，不再生效）。
	//
	// Deprecated: AuditRuleFilters is deprecated.
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// 规则模板 ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// 审计类型。true - 全审计；默认 false - 规则审计。
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`
}

type OpenAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// CDB 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计日志保存时长。支持值包括：
	// 7 - 一周；
	// 30 - 一个月；
	// 90 - 三个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年。
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 高频审计日志保存时长。默认值为7，此项取值需小于等于 LogExpireDay，支持值包括：
	// 3 - 3天；
	// 7 - 一周；
	// 30 - 一个月；
	// 90 - 三个月；
	// 180 - 六个月；
	// 365 - 一年；
	// 1095 - 三年；
	// 1825 - 五年。
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// 审计规则（该参数已废弃，不再生效）。
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// 规则模板 ID。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// 审计类型。true - 全审计；默认 false - 规则审计。
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
type OpenDBInstanceEncryptionRequestParams struct {
	// 云数据库实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户自定义密钥 ID，CMK 唯一标识符。该值为空时，将使用腾讯云自动生成的密钥 KMS-CDB。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 用户自定义密钥的存储地域。如：ap-guangzhou 。KeyId 不为空时，该参数必填。
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`
}

type OpenDBInstanceEncryptionRequest struct {
	*tchttp.BaseRequest
	
	// 云数据库实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户自定义密钥 ID，CMK 唯一标识符。该值为空时，将使用腾讯云自动生成的密钥 KMS-CDB。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 用户自定义密钥的存储地域。如：ap-guangzhou 。KeyId 不为空时，该参数必填。
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type OpenDBInstanceGTIDRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type OpenSSLRequestParams struct {
	// 实例 ID。只读组 ID 为空时必填。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 只读组 ID。实例 ID 为空时必填。可通过 [DescribeRoGroups](https://cloud.tencent.com/document/api/236/40939) 接口获取。
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`
}

type OpenSSLRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。只读组 ID 为空时必填。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 只读组 ID。实例 ID 为空时必填。可通过 [DescribeRoGroups](https://cloud.tencent.com/document/api/236/40939) 接口获取。
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`
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
	delete(f, "InstanceId")
	delete(f, "RoGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenSSLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenSSLResponseParams struct {
	// 异步请求 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

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
type OpenWanServiceRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。可以传入只读组 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 变更集群版实例只读组时，InstanceId传实例id，需要额外指定该参数表示操作只读组。 如果操作读写节点则不需指定该参数。
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
}

type OpenWanServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。可以传入只读组 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 变更集群版实例只读组时，InstanceId传实例id，需要额外指定该参数表示操作只读组。 如果操作读写节点则不需指定该参数。
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
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
	delete(f, "OpResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenWanServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenWanServiceResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 目的 IP 或 IP 段，例如172.16.0.0/12
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// 端口或者端口范围
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// 网络协议，支持 UDP、TCP等
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// 规则限定的方向，进站规则为 OUTPUT
	Dir *string `json:"Dir,omitnil,omitempty" name:"Dir"`

	// 地址模块
	AddressModule *string `json:"AddressModule,omitnil,omitempty" name:"AddressModule"`

	// 规则描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type ParamInfo struct {
	// 参数名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ParamRecord struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 参数名称
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 参数修改前的值
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`

	// 参数修改后的值
	NewValue *string `json:"NewValue,omitnil,omitempty" name:"NewValue"`

	// 参数是否修改成功
	//
	// Deprecated: IsSucess is deprecated.
	IsSucess *bool `json:"IsSucess,omitnil,omitempty" name:"IsSucess"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 参数是否修改成功
	IsSuccess *bool `json:"IsSuccess,omitnil,omitempty" name:"IsSuccess"`
}

type ParamTemplateInfo struct {
	// 参数模板 ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 参数模板名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数模板描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 实例引擎版本，值为：5.5、5.6、5.7、8.0。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 参数模板类型，值为：HIGH_STABILITY、HIGH_PERFORMANCE。
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// 参数模板引擎，值为：InnoDB、RocksDB。
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
}

type Parameter struct {
	// 参数名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数值
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`
}

type ParameterDetail struct {
	// 参数名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数类型：integer，enum，float，string，func
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// 参数默认值
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// 参数描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 参数当前值
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 修改参数后，是否需要重启数据库以使参数生效。可能的值包括：0-不需要重启；1-需要重启
	NeedReboot *int64 `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// 参数允许的最大值
	Max *int64 `json:"Max,omitnil,omitempty" name:"Max"`

	// 参数允许的最小值
	Min *int64 `json:"Min,omitnil,omitempty" name:"Min"`

	// 参数的可选枚举值。如果为非枚举参数，则为空
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// 参数是公式类型时，该字段有效，表示公式类型最大值
	MaxFunc *string `json:"MaxFunc,omitnil,omitempty" name:"MaxFunc"`

	// 参数是公式类型时，该字段有效，表示公式类型最小值
	MinFunc *string `json:"MinFunc,omitnil,omitempty" name:"MinFunc"`

	// 参数是否不支持修改
	IsNotSupportEdit *bool `json:"IsNotSupportEdit,omitnil,omitempty" name:"IsNotSupportEdit"`
}

type PeriodStrategy struct {
	// 扩容周期
	TimeCycle *TImeCycle `json:"TimeCycle,omitnil,omitempty" name:"TimeCycle"`

	// 时间间隔
	TimeInterval *TimeInterval `json:"TimeInterval,omitnil,omitempty" name:"TimeInterval"`
}

type ProxyAddress struct {
	// 代理组地址ID
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`

	// 私有网络ID
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 私有子网ID
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// IP地址
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 端口
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// 权重分配模式；
	// 系统自动分配："system"， 自定义："custom"
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// 是否开启延迟剔除，取值："true" | "false"
	IsKickOut *bool `json:"IsKickOut,omitnil,omitempty" name:"IsKickOut"`

	// 最小保留数量，最小取值：0
	MinCount *uint64 `json:"MinCount,omitnil,omitempty" name:"MinCount"`

	// 延迟剔除阈值，最小取值：0
	MaxDelay *uint64 `json:"MaxDelay,omitnil,omitempty" name:"MaxDelay"`

	// 是否自动添加RO，取值："true" | "false"
	AutoAddRo *bool `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// 是否是只读，取值："true" | "false"
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// 是否开启事务分离
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// 是否开启故障转移
	FailOver *bool `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// 是否开启连接池
	ConnectionPool *bool `json:"ConnectionPool,omitnil,omitempty" name:"ConnectionPool"`

	// 描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 实例读权重分配
	ProxyAllocation []*ProxyAllocation `json:"ProxyAllocation,omitnil,omitempty" name:"ProxyAllocation"`

	// 接入模式
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// 是否开启自动负载均衡
	AutoLoadBalance *bool `json:"AutoLoadBalance,omitnil,omitempty" name:"AutoLoadBalance"`

	// 是否把libra当作ro节点
	ApNodeAsRoNode *bool `json:"ApNodeAsRoNode,omitnil,omitempty" name:"ApNodeAsRoNode"`

	// libra节点故障，是否转发给其他节点
	ApQueryToOtherNode *bool `json:"ApQueryToOtherNode,omitnil,omitempty" name:"ApQueryToOtherNode"`
}

type ProxyAllocation struct {
	// 代理节点所属地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 代理节点所属可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 代理实例分布
	ProxyInstance []*ProxyInst `json:"ProxyInstance,omitnil,omitempty" name:"ProxyInstance"`
}

type ProxyGroupInfo struct {
	// 代理组ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 代理版本
	ProxyVersion *string `json:"ProxyVersion,omitnil,omitempty" name:"ProxyVersion"`

	// 代理支持升级版本
	SupportUpgradeProxyVersion *string `json:"SupportUpgradeProxyVersion,omitnil,omitempty" name:"SupportUpgradeProxyVersion"`

	// 代理状态。0 - 初始化中，1 - 在线中，2 - 在线中-读写分离中，3 - 下线，4 - 销毁。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 代理任务状态，Upgrading - 升级中，UpgradeTo - 升级待切换，UpgradeSwitching - 升级切换中，ProxyCreateAddress - 配置地址中，ProxyModifyAddress - 修改地址中，ProxyCloseAddress - 关闭地址中。
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 代理组节点信息
	ProxyNode []*ProxyNode `json:"ProxyNode,omitnil,omitempty" name:"ProxyNode"`

	// 代理组地址信息
	ProxyAddress []*ProxyAddress `json:"ProxyAddress,omitnil,omitempty" name:"ProxyAddress"`

	// 连接池阈值
	ConnectionPoolLimit *uint64 `json:"ConnectionPoolLimit,omitnil,omitempty" name:"ConnectionPoolLimit"`

	// 支持创建地址
	SupportCreateProxyAddress *bool `json:"SupportCreateProxyAddress,omitnil,omitempty" name:"SupportCreateProxyAddress"`

	// 支持升级代理版本所需的cdb版本
	SupportUpgradeProxyMysqlVersion *string `json:"SupportUpgradeProxyMysqlVersion,omitnil,omitempty" name:"SupportUpgradeProxyMysqlVersion"`
}

type ProxyInst struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例类型：1 master 主实例; 2 ro 只读实例; 3 dr 灾备实例; 4 sdr 小灾备实例
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例状态，可能的返回值：0-创建中；1-运行中；4-隔离中；5-已隔离
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 只读权重,如果权重为系统自动分配，改值不生效，只代表是否启用该实例
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 实例所属地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 实例所属可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例节点ID
	InstNodeId *string `json:"InstNodeId,omitnil,omitempty" name:"InstNodeId"`

	// 节点角色
	InstNodeRole *string `json:"InstNodeRole,omitnil,omitempty" name:"InstNodeRole"`
}

type ProxyNode struct {
	// 代理节点ID
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// CPU核数
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存大小，单位为 MB。
	Mem *uint64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 节点状态，0 - 初始化中，1 - 在线中，2 - 下线中，3 - 销毁中，4 - 故障恢复中，5 - 节点故障，6 - 切换中。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 代理节点可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 代理节点地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 连接数
	Connection *uint64 `json:"Connection,omitnil,omitempty" name:"Connection"`
}

type ProxyNodeCustom struct {
	// 节点个数
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// CPU核数
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存大小
	Mem *uint64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type ReadWriteNode struct {
	// RW 节点所在可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 升级集群版实例时，如果要调整只读节点可用区，需要指定节点id。
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

type ReadonlyNode struct {
	// 是否分布在随机可用区。传入YES表示随机可用区。否则使用Zone指定的可用区。
	IsRandomZone *string `json:"IsRandomZone,omitnil,omitempty" name:"IsRandomZone"`

	// 指定该节点分布在哪个可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 升级集群版实例时，如果要调整只读节点可用区，需要指定节点id。
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

// Predefined struct for user
type ReleaseIsolatedDBInstancesRequestParams struct {
	// 实例 ID，单个实例 ID 格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	// 说明：可输入多个实例 ID 进行操作，json 格式如下。
	// [
	//     "cdb-30z11v8s",
	//     "cdb-93h11efg"
	//   ]
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type ReleaseIsolatedDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，单个实例 ID 格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	// 说明：可输入多个实例 ID 进行操作，json 格式如下。
	// [
	//     "cdb-30z11v8s",
	//     "cdb-93h11efg"
	//   ]
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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
	Items []*ReleaseResult `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例解隔离操作的结果值。返回值为0表示成功。
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 实例解隔离操作的错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type ReloadBalanceProxyNodeRequestParams struct {
	// 代理组 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 代理组地址 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。如果不传则会对所有代理组地址进行负载均衡。
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`
}

type ReloadBalanceProxyNodeRequest struct {
	*tchttp.BaseRequest
	
	// 代理组 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 代理组地址 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。如果不传则会对所有代理组地址进行负载均衡。
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`
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

type RemoteBackupInfo struct {
	// 异地备份子任务的ID
	SubBackupId *int64 `json:"SubBackupId,omitnil,omitempty" name:"SubBackupId"`

	// 异地备份所在地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 备份任务状态。可能的值有 "SUCCESS": 备份成功， "FAILED": 备份失败， "RUNNING": 备份进行中。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 异地备份任务的开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 异地备份任务的结束时间
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 下载地址
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

// Predefined struct for user
type RenewDBInstanceRequestParams struct {
	// 待续费的实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872)。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 续费时长，单位：月，可选值包括 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]。
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 如果需要将按量计费实例续费为包年包月的实例，该入参的值需要指定为 "PREPAID" 。
	ModifyPayType *string `json:"ModifyPayType,omitnil,omitempty" name:"ModifyPayType"`

	// 自动续费标记，0表示不自动续费，1表示进行自动续费
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`
}

type RenewDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 待续费的实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872)。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 续费时长，单位：月，可选值包括 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]。
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 如果需要将按量计费实例续费为包年包月的实例，该入参的值需要指定为 "PREPAID" 。
	ModifyPayType *string `json:"ModifyPayType,omitnil,omitempty" name:"ModifyPayType"`

	// 自动续费标记，0表示不自动续费，1表示进行自动续费
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`
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
	delete(f, "AutoRenew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDBInstanceResponseParams struct {
	// 订单 ID。
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ResetPasswordRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 手动刷新轮转密码的实例账户名，例如root
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 手动刷新轮转密码的实例账户域名，例如%
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type ResetPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 手动刷新轮转密码的实例账户名，例如root
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 手动刷新轮转密码的实例账户域名，例如%
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

func (r *ResetPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetPasswordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetPasswordResponseParams `json:"Response"`
}

func (r *ResetPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetRootAccountRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ResetRootAccountRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type RestartDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 数组，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RoGroupMode *string `json:"RoGroupMode,omitnil,omitempty" name:"RoGroupMode"`

	// 只读组 ID。
	// 说明：若此数据结构在购买实例操作中被使用，则当只读组模式选择 join 时，此项为必填。
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`

	// 只读组名称。
	RoGroupName *string `json:"RoGroupName,omitnil,omitempty" name:"RoGroupName"`

	// 是否启用延迟超限剔除功能，启用该功能后，只读实例与主实例的延迟超过延迟阈值，只读实例将被隔离。可选值：1-启用；0-不启用。
	RoOfflineDelay *int64 `json:"RoOfflineDelay,omitnil,omitempty" name:"RoOfflineDelay"`

	// 延迟阈值。单位：秒。值范围：1-10000，整数。
	RoMaxDelayTime *int64 `json:"RoMaxDelayTime,omitnil,omitempty" name:"RoMaxDelayTime"`

	// 最少实例保留个数，若购买只读实例数量小于设置数量将不做剔除。
	MinRoInGroup *int64 `json:"MinRoInGroup,omitnil,omitempty" name:"MinRoInGroup"`

	// 读写权重分配模式，可选值：system-系统自动分配；custom-自定义。
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// 该字段已经废弃，无意义。查看只读实例的权重，请查看 RoInstances 字段里的 Weight 值。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 只读组中的只读实例详情。
	RoInstances []*RoInstanceInfo `json:"RoInstances,omitnil,omitempty" name:"RoInstances"`

	// 只读组的内网 IP。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 只读组的内网端口号。
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 私有网络 ID。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 子网 ID。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 只读组所在的地域。
	RoGroupRegion *string `json:"RoGroupRegion,omitnil,omitempty" name:"RoGroupRegion"`

	// 只读组所在的可用区。
	RoGroupZone *string `json:"RoGroupZone,omitnil,omitempty" name:"RoGroupZone"`

	// 延迟复制时间。单位：秒。值范围：1-259200，整数。
	DelayReplicationTime *int64 `json:"DelayReplicationTime,omitnil,omitempty" name:"DelayReplicationTime"`
}

type RoGroupAttr struct {
	// RO 组名称。
	RoGroupName *string `json:"RoGroupName,omitnil,omitempty" name:"RoGroupName"`

	// RO 实例最大延迟阈值。单位为秒，最小值为 1。范围：[1,10000]，整数。
	// 注意：RO 组必须设置了开启实例延迟剔除策略，该值才有效。
	RoMaxDelayTime *int64 `json:"RoMaxDelayTime,omitnil,omitempty" name:"RoMaxDelayTime"`

	// 是否开启实例延迟剔除。支持的值包括：1 - 开启；0 - 不开启。注意，若设置开启实例延迟剔除，则必须设置延迟阈值（RoMaxDelayTime）参数。
	RoOfflineDelay *int64 `json:"RoOfflineDelay,omitnil,omitempty" name:"RoOfflineDelay"`

	// 最少保留实例数。可设置为小于或等于该 RO 组下 RO 实例个数的任意值。默认值为1。
	// 注意：若设置值大于 RO 实例数量将不做剔除；若设置为 0，所有实例延迟超限都会被剔除。
	MinRoInGroup *int64 `json:"MinRoInGroup,omitnil,omitempty" name:"MinRoInGroup"`

	// 权重模式。支持值包括："system" - 系统自动分配； "custom" - 用户自定义设置。注意，若设置 "custom" 模式，则必须设置 RO 实例权重配置（RoWeightValues）参数。
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// 延迟复制时间。单位：秒，范围：1 - 259200秒，不传此参数表示不开启实例延迟复制。
	ReplicationDelayTime *int64 `json:"ReplicationDelayTime,omitnil,omitempty" name:"ReplicationDelayTime"`
}

type RoInstanceInfo struct {
	// RO组对应的主实例的ID
	MasterInstanceId *string `json:"MasterInstanceId,omitnil,omitempty" name:"MasterInstanceId"`

	// RO实例在RO组内的状态，可能的值：online-在线，offline-下线
	RoStatus *string `json:"RoStatus,omitnil,omitempty" name:"RoStatus"`

	// RO实例在RO组内上一次下线的时间
	OfflineTime *string `json:"OfflineTime,omitnil,omitempty" name:"OfflineTime"`

	// RO实例在RO组内的权重
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// RO实例所在区域名称，如ap-shanghai
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// RO可用区的正式名称，如ap-shanghai-1
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// RO实例ID，格式如：cdbro-c1nl9rpv
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// RO实例状态，可能返回值：0-创建中，1-运行中，3-异地RO（仅在使用DescribeDBInstances查询主实例信息时，返回值中异地RO的状态恒等于3，其他场景下无此值），4-删除中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例类型，可能返回值：1-主实例，2-灾备实例，3-只读实例
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// RO实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 按量计费状态，可能的取值：1-正常，2-欠费
	HourFeeStatus *int64 `json:"HourFeeStatus,omitnil,omitempty" name:"HourFeeStatus"`

	// RO实例任务状态，可能返回值：<br>0-没有任务<br>1-升级中<br>2-数据导入中<br>3-开放Slave中<br>4-外网访问开通中<br>5-批量操作执行中<br>6-回档中<br>7-外网访问关闭中<br>8-密码修改中<br>9-实例名修改中<br>10-重启中<br>12-自建迁移中<br>13-删除库表中<br>14-灾备实例创建同步中
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// RO实例内存大小，单位：MB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// RO实例硬盘大小，单位：GB
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 每次查询数量
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// RO实例的内网IP地址
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// RO实例访问端口
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// RO实例所在私有网络ID
	VpcId *int64 `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// RO实例所在私有网络子网ID
	SubnetId *int64 `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// RO实例规格描述，目前可取值 CUSTOM
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// RO实例数据库引擎版本，可能返回值：5.1、5.5、5.6、5.7、8.0
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// RO实例到期时间，时间格式：yyyy-mm-dd hh:mm:ss，如实例为按量计费模式，则此字段值为0000-00-00 00:00:00
	DeadlineTime *string `json:"DeadlineTime,omitnil,omitempty" name:"DeadlineTime"`

	// RO实例计费类型，可能返回值：0-包年包月，1-按量计费，2-后付费月结
	PayType *int64 `json:"PayType,omitnil,omitempty" name:"PayType"`

	// RO延迟复制状态。
	ReplicationStatus *string `json:"ReplicationStatus,omitnil,omitempty" name:"ReplicationStatus"`
}

type RoVipInfo struct {
	// 只读vip状态
	RoVipStatus *int64 `json:"RoVipStatus,omitnil,omitempty" name:"RoVipStatus"`

	// 只读vip的子网
	RoSubnetId *int64 `json:"RoSubnetId,omitnil,omitempty" name:"RoSubnetId"`

	// 只读vip的私有网络
	RoVpcId *int64 `json:"RoVpcId,omitnil,omitempty" name:"RoVpcId"`

	// 只读vip的端口号
	RoVport *int64 `json:"RoVport,omitnil,omitempty" name:"RoVport"`

	// 只读vip
	RoVip *string `json:"RoVip,omitnil,omitempty" name:"RoVip"`
}

type RoWeightValue struct {
	// RO 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 权重值。取值范围为 [0, 100]。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type RollbackDBName struct {
	// 回档前的原数据库名
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 回档后的新数据库名
	NewDatabaseName *string `json:"NewDatabaseName,omitnil,omitempty" name:"NewDatabaseName"`
}

type RollbackInstancesInfo struct {
	// 云数据库实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 回档策略。可选值为：table、db、full。table - 极速回档模式，仅导入所选中表级别的备份和binlog，如有跨表操作，且关联表未被同时选中，将会导致回档失败，该模式下参数Databases必须为空；db - 快速模式，仅导入所选中库级别的备份和binlog，如有跨库操作，且关联库未被同时选中，将会导致回档失败；full - 普通回档模式，将导入整个实例的备份和 binlog，速度较慢。
	Strategy *string `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// 数据库回档时间，时间格式为：yyyy-mm-dd hh:mm:ss。
	RollbackTime *string `json:"RollbackTime,omitnil,omitempty" name:"RollbackTime"`

	// 待回档的数据库信息，表示整库回档。
	Databases []*RollbackDBName `json:"Databases,omitnil,omitempty" name:"Databases"`

	// 待回档的数据库表信息，表示按表回档。
	Tables []*RollbackTables `json:"Tables,omitnil,omitempty" name:"Tables"`
}

type RollbackTableName struct {
	// 回档前的原数据库表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 回档后的新数据库表名
	NewTableName *string `json:"NewTableName,omitnil,omitempty" name:"NewTableName"`
}

type RollbackTables struct {
	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 数据库表详情
	Table []*RollbackTableName `json:"Table,omitnil,omitempty" name:"Table"`
}

type RollbackTask struct {
	// 任务执行信息描述。
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// 任务执行结果。可能的取值：INITIAL - 初始化，RUNNING - 运行中，SUCCESS - 执行成功，FAILED - 执行失败，KILLED - 已终止，REMOVED - 已删除，PAUSED - 终止中。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务执行进度。取值范围为[0, 100]。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 任务开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 回档任务详情。
	Detail []*RollbackInstancesInfo `json:"Detail,omitnil,omitempty" name:"Detail"`
}

type RollbackTimeRange struct {
	// 实例可回档开始时间，时间格式：2016-10-29 01:06:04
	Begin *string `json:"Begin,omitnil,omitempty" name:"Begin"`

	// 实例可回档结束时间，时间格式：2016-11-02 11:44:47
	End *string `json:"End,omitnil,omitempty" name:"End"`
}

type Rule struct {
	// 划分上限
	LessThan *uint64 `json:"LessThan,omitnil,omitempty" name:"LessThan"`

	// 权重
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type RuleFilters struct {
	// 审计规则过滤条件的参数名称。可选值：host – 客户端 IP；user – 数据库账户；dbName – 数据库名称；sqlType-SQL类型；sql-sql语句；affectRows -影响行数；sentRows-返回行数；checkRows-扫描行数；execTime-执行时间。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 审计规则过滤条件的匹配类型。可选值：INC – 包含；EXC – 不包含；EQS – 等于；NEQ – 不等于；REG-正则；GT-大于；LT-小于。
	Compare *string `json:"Compare,omitnil,omitempty" name:"Compare"`

	// 审计规则过滤条件的匹配值。sqlType条件的Value需在以下选择"alter", "changeuser", "create", "delete", "drop", "execute", "insert", "login", "logout", "other", "replace", "select", "set", "update"。
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

type RuleTemplateRecordInfo struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 修改前规则模板的详情。
	ModifyBeforeInfo *RuleTemplateInfo `json:"ModifyBeforeInfo,omitnil,omitempty" name:"ModifyBeforeInfo"`

	// 修改后规则模板的详情。
	ModifyAfterInfo *RuleTemplateInfo `json:"ModifyAfterInfo,omitnil,omitempty" name:"ModifyAfterInfo"`

	// 影响的实例。
	AffectedInstances []*string `json:"AffectedInstances,omitnil,omitempty" name:"AffectedInstances"`

	// 操作人，账号uin。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 变更的时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type SecurityGroup struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 创建时间，时间格式：yyyy-mm-dd hh:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 入站规则
	Inbound []*Inbound `json:"Inbound,omitnil,omitempty" name:"Inbound"`

	// 出站规则
	Outbound []*Outbound `json:"Outbound,omitnil,omitempty" name:"Outbound"`

	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 安全组名称
	SecurityGroupName *string `json:"SecurityGroupName,omitnil,omitempty" name:"SecurityGroupName"`

	// 安全组备注
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil,omitempty" name:"SecurityGroupRemark"`
}

type SlaveConfig struct {
	// 从库复制方式，可能的返回值：aysnc-异步，semisync-半同步
	ReplicationMode *string `json:"ReplicationMode,omitnil,omitempty" name:"ReplicationMode"`

	// 从库可用区的正式名称，如ap-shanghai-1
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type SlaveInfo struct {
	// 第一备机信息
	First *SlaveInstanceInfo `json:"First,omitnil,omitempty" name:"First"`

	// 第二备机信息
	Second *SlaveInstanceInfo `json:"Second,omitnil,omitempty" name:"Second"`
}

type SlaveInstanceInfo struct {
	// 端口号
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 地域信息
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 虚拟 IP 信息
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 可用区信息
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type SlowLogInfo struct {
	// 备份文件名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备份文件大小，单位：Byte
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 备份快照时间，时间格式：2016-03-17
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 内网下载地址
	IntranetUrl *string `json:"IntranetUrl,omitnil,omitempty" name:"IntranetUrl"`

	// 外网下载地址
	InternetUrl *string `json:"InternetUrl,omitnil,omitempty" name:"InternetUrl"`

	// 日志具体类型，可能的值：slowlog - 慢日志
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type SlowLogItem struct {
	// Sql的执行时间。秒级时间戳。
	Timestamp *uint64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Sql的执行时长（秒）。
	QueryTime *float64 `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// Sql语句。
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// 客户端地址。
	UserHost *string `json:"UserHost,omitnil,omitempty" name:"UserHost"`

	// 用户名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 数据库名。
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 锁时长（秒）。
	LockTime *float64 `json:"LockTime,omitnil,omitempty" name:"LockTime"`

	// 扫描行数。
	RowsExamined *int64 `json:"RowsExamined,omitnil,omitempty" name:"RowsExamined"`

	// 结果集行数。
	RowsSent *int64 `json:"RowsSent,omitnil,omitempty" name:"RowsSent"`

	// Sql模板。
	SqlTemplate *string `json:"SqlTemplate,omitnil,omitempty" name:"SqlTemplate"`

	// Sql语句的md5。
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`
}

type SqlFileInfo struct {
	// 上传时间
	UploadTime *string `json:"UploadTime,omitnil,omitempty" name:"UploadTime"`

	// 上传进度
	UploadInfo *UploadInfo `json:"UploadInfo,omitnil,omitempty" name:"UploadInfo"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件大小，单位为Bytes
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 上传是否完成标志，可选值：0 - 未完成，1 - 已完成
	IsUploadFinished *int64 `json:"IsUploadFinished,omitnil,omitempty" name:"IsUploadFinished"`

	// 文件ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`
}

// Predefined struct for user
type StartBatchRollbackRequestParams struct {
	// 用于回档的实例详情信息。
	Instances []*RollbackInstancesInfo `json:"Instances,omitnil,omitempty" name:"Instances"`
}

type StartBatchRollbackRequest struct {
	*tchttp.BaseRequest
	
	// 用于回档的实例详情信息。
	Instances []*RollbackInstancesInfo `json:"Instances,omitnil,omitempty" name:"Instances"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type StartCpuExpandRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 扩容类型，支持自动扩容和自定义扩容。
	// 说明：1. auto 表示自动扩容。2. manual 表示自定义扩容，扩容时间为立即生效。3. timeInterval 表示自定义扩容，扩容时间为按时间段。4. period 表示自定义扩容，扩容时间为按周期。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 自定义扩容时，扩容的 CPU 核心数。
	// 说明：1. Type 为 manual、timeInterval、period 时必传。2. 扩容的 CPU 核心数上限为当前实例 CPU 核心数，比如8核16G最大可手动扩容的 CPU 核心数为8，即范围为1 - 8。
	ExpandCpu *int64 `json:"ExpandCpu,omitnil,omitempty" name:"ExpandCpu"`

	// 自动扩容策略。Type 为 auto 时必传。
	AutoStrategy *AutoStrategy `json:"AutoStrategy,omitnil,omitempty" name:"AutoStrategy"`

	// 按时间段扩容策略。
	// 说明：当 Type 为 timeInterval 时，TimeIntervalStrategy 必填。
	TimeIntervalStrategy *TimeIntervalStrategy `json:"TimeIntervalStrategy,omitnil,omitempty" name:"TimeIntervalStrategy"`

	// 按周期扩容策略。
	// 说明：当 Type 为 period 时，PeriodStrategy 必填。
	PeriodStrategy *PeriodStrategy `json:"PeriodStrategy,omitnil,omitempty" name:"PeriodStrategy"`
}

type StartCpuExpandRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 扩容类型，支持自动扩容和自定义扩容。
	// 说明：1. auto 表示自动扩容。2. manual 表示自定义扩容，扩容时间为立即生效。3. timeInterval 表示自定义扩容，扩容时间为按时间段。4. period 表示自定义扩容，扩容时间为按周期。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 自定义扩容时，扩容的 CPU 核心数。
	// 说明：1. Type 为 manual、timeInterval、period 时必传。2. 扩容的 CPU 核心数上限为当前实例 CPU 核心数，比如8核16G最大可手动扩容的 CPU 核心数为8，即范围为1 - 8。
	ExpandCpu *int64 `json:"ExpandCpu,omitnil,omitempty" name:"ExpandCpu"`

	// 自动扩容策略。Type 为 auto 时必传。
	AutoStrategy *AutoStrategy `json:"AutoStrategy,omitnil,omitempty" name:"AutoStrategy"`

	// 按时间段扩容策略。
	// 说明：当 Type 为 timeInterval 时，TimeIntervalStrategy 必填。
	TimeIntervalStrategy *TimeIntervalStrategy `json:"TimeIntervalStrategy,omitnil,omitempty" name:"TimeIntervalStrategy"`

	// 按周期扩容策略。
	// 说明：当 Type 为 period 时，PeriodStrategy 必填。
	PeriodStrategy *PeriodStrategy `json:"PeriodStrategy,omitnil,omitempty" name:"PeriodStrategy"`
}

func (r *StartCpuExpandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartCpuExpandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	delete(f, "ExpandCpu")
	delete(f, "AutoStrategy")
	delete(f, "TimeIntervalStrategy")
	delete(f, "PeriodStrategy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartCpuExpandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartCpuExpandResponseParams struct {
	// 异步任务 ID 。可以调用 DescribeAsyncRequest 传入该 ID，进行任务执行进度的查询。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartCpuExpandResponse struct {
	*tchttp.BaseResponse
	Response *StartCpuExpandResponseParams `json:"Response"`
}

func (r *StartCpuExpandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartCpuExpandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartReplicationRequestParams struct {
	// 实例 ID。仅支持只读实例。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type StartReplicationRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。仅支持只读实例。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type StopCpuExpandRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type StopCpuExpandRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *StopCpuExpandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCpuExpandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopCpuExpandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCpuExpandResponseParams struct {
	// 异步任务 ID。在调用 [DescribeAsyncRequestInfo](https://cloud.tencent.com/document/api/236/20410) 进行任务执行进度的查询时，可以传入该 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopCpuExpandResponse struct {
	*tchttp.BaseResponse
	Response *StopCpuExpandResponseParams `json:"Response"`
}

func (r *StopCpuExpandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCpuExpandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopDBImportJobRequestParams struct {
	// 异步任务的请求 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`
}

type StopDBImportJobRequest struct {
	*tchttp.BaseRequest
	
	// 异步任务的请求 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 实例 ID。仅支持只读实例。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type StopReplicationRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。仅支持只读实例。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 撤销回档任务对应的实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/api/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type StopRollbackRequest struct {
	*tchttp.BaseRequest
	
	// 撤销回档任务对应的实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/api/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	// 执行请求的异步任务 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type SubmitInstanceUpgradeCheckJobRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 目标数据库版本。
	// 说明：可选值5.6、5.7、8.0，不支持跨版本升级，升级后不支持版本降级。
	DstMysqlVersion *string `json:"DstMysqlVersion,omitnil,omitempty" name:"DstMysqlVersion"`
}

type SubmitInstanceUpgradeCheckJobRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 目标数据库版本。
	// 说明：可选值5.6、5.7、8.0，不支持跨版本升级，升级后不支持版本降级。
	DstMysqlVersion *string `json:"DstMysqlVersion,omitnil,omitempty" name:"DstMysqlVersion"`
}

func (r *SubmitInstanceUpgradeCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitInstanceUpgradeCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DstMysqlVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitInstanceUpgradeCheckJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitInstanceUpgradeCheckJobResponseParams struct {
	// 任务 ID
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitInstanceUpgradeCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitInstanceUpgradeCheckJobResponseParams `json:"Response"`
}

func (r *SubmitInstanceUpgradeCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitInstanceUpgradeCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchCDBProxyRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库代理 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

type SwitchCDBProxyRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库代理 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 目标从实例。可选值："first" - 第一备机；"second" - 第二备机。默认值为 "first"，仅多可用区实例支持设置为 "second"。
	DstSlave *string `json:"DstSlave,omitnil,omitempty" name:"DstSlave"`

	// 是否强制切换。默认为 False。注意，若设置强制切换为 True，实例存在丢失数据的风险，请谨慎使用。
	ForceSwitch *bool `json:"ForceSwitch,omitnil,omitempty" name:"ForceSwitch"`

	// 是否时间窗内切换。默认为 False，即不在时间窗内切换。注意，如果设置了 ForceSwitch 参数为 True，则该参数不生效。
	WaitSwitch *bool `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// 集群版实例指定节点 ID 发起主从切换。
	DstNodeId *string `json:"DstNodeId,omitnil,omitempty" name:"DstNodeId"`
}

type SwitchDBInstanceMasterSlaveRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 目标从实例。可选值："first" - 第一备机；"second" - 第二备机。默认值为 "first"，仅多可用区实例支持设置为 "second"。
	DstSlave *string `json:"DstSlave,omitnil,omitempty" name:"DstSlave"`

	// 是否强制切换。默认为 False。注意，若设置强制切换为 True，实例存在丢失数据的风险，请谨慎使用。
	ForceSwitch *bool `json:"ForceSwitch,omitnil,omitempty" name:"ForceSwitch"`

	// 是否时间窗内切换。默认为 False，即不在时间窗内切换。注意，如果设置了 ForceSwitch 参数为 True，则该参数不生效。
	WaitSwitch *bool `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// 集群版实例指定节点 ID 发起主从切换。
	DstNodeId *string `json:"DstNodeId,omitnil,omitempty" name:"DstNodeId"`
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
	delete(f, "DstNodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchDBInstanceMasterSlaveRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDBInstanceMasterSlaveResponseParams struct {
	// 异步任务 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 灾备实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type SwitchDrInstanceToMasterRequest struct {
	*tchttp.BaseRequest
	
	// 灾备实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type SwitchForUpgradeRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type TImeCycle struct {
	// 按周期扩容时，是否选择周一扩容。
	// 说明：取值 true，表示选择，取值 false，表示不选择。
	Monday *bool `json:"Monday,omitnil,omitempty" name:"Monday"`

	// 按周期扩容时，是否选择周二扩容。
	// 说明：取值 true，表示选择，取值 false，表示不选择。
	Tuesday *bool `json:"Tuesday,omitnil,omitempty" name:"Tuesday"`

	// 按周期扩容时，是否选择周三扩容。
	// 说明：取值 true，表示选择，取值 false，表示不选择。
	Wednesday *bool `json:"Wednesday,omitnil,omitempty" name:"Wednesday"`

	// 按周期扩容时，是否选择周四扩容。
	// 说明：取值 true，表示选择，取值 false，表示不选择。
	Thursday *bool `json:"Thursday,omitnil,omitempty" name:"Thursday"`

	// 按周期扩容时，是否选择周五扩容。
	// 说明：取值 true，表示选择，取值 false，表示不选择。
	Friday *bool `json:"Friday,omitnil,omitempty" name:"Friday"`

	// 按周期扩容时，是否选择周六扩容。
	// 说明：取值 true，表示选择，取值 false，表示不选择。
	Saturday *bool `json:"Saturday,omitnil,omitempty" name:"Saturday"`

	// 按周期扩容时，是否选择周日扩容。
	// 说明：取值 true，表示选择，取值 false，表示不选择。
	Sunday *bool `json:"Sunday,omitnil,omitempty" name:"Sunday"`
}

type TablePrivilege struct {
	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 数据库表名
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 权限信息
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TagInfo struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue []*string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagInfoItem struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagInfoUnit struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagsInfoOfInstance struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 标签信息
	Tags []*TagInfoUnit `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type TaskAttachInfo struct {
	// 升级任务：
	// ”FastUpgradeStatus“：表示升级类型。1-原地升级；0-普通升级。
	AttachKey *string `json:"AttachKey,omitnil,omitempty" name:"AttachKey"`

	// 升级任务：
	// ”FastUpgradeStatus“：表示升级类型。1-原地升级；0-普通升级。
	AttachValue *string `json:"AttachValue,omitnil,omitempty" name:"AttachValue"`
}

type TaskDetail struct {
	// 错误码。0代表成功，其他对应不同的报错场景。
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 实例任务 ID。
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 实例任务进度。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

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
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

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
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 实例任务开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 实例任务结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务关联的实例 ID。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 异步任务的请求 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 任务的附加信息。
	TaskAttachInfo []*TaskAttachInfo `json:"TaskAttachInfo,omitnil,omitempty" name:"TaskAttachInfo"`
}

type TimeInterval struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type TimeIntervalStrategy struct {
	// 开始扩容时间。
	// 说明：此值的格式为 Integer 的时间戳（秒级）。
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束扩容时间。
	// 说明：此值的格式为 Integer 的时间戳（秒级）。
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

// Predefined struct for user
type UpgradeCDBProxyVersionRequestParams struct {
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库代理 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 数据库代理当前版本
	SrcProxyVersion *string `json:"SrcProxyVersion,omitnil,omitempty" name:"SrcProxyVersion"`

	// 数据库代理升级版本
	DstProxyVersion *string `json:"DstProxyVersion,omitnil,omitempty" name:"DstProxyVersion"`

	// 升级时间 ：nowTime（升级完成时）timeWindow（实例维护时间）
	UpgradeTime *string `json:"UpgradeTime,omitnil,omitempty" name:"UpgradeTime"`
}

type UpgradeCDBProxyVersionRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDBInstances](https://cloud.tencent.com/document/product/236/15872) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库代理 ID。可通过 [DescribeCdbProxyInfo](https://cloud.tencent.com/document/api/236/90585) 接口获取。
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// 数据库代理当前版本
	SrcProxyVersion *string `json:"SrcProxyVersion,omitnil,omitempty" name:"SrcProxyVersion"`

	// 数据库代理升级版本
	DstProxyVersion *string `json:"DstProxyVersion,omitnil,omitempty" name:"DstProxyVersion"`

	// 升级时间 ：nowTime（升级完成时）timeWindow（实例维护时间）
	UpgradeTime *string `json:"UpgradeTime,omitnil,omitempty" name:"UpgradeTime"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主实例数据库引擎版本，支持值包括：5.6、5.7、8.0。
	// 说明：不支持越级升级，升级后不支持降级。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 切换访问新实例的方式，默认为 0。支持值包括：0 - 立刻切换，1 - 时间窗切换；当该值为 1 时，升级过程中，切换访问新实例的流程将会在时间窗内进行，或者用户主动调用接口 [切换访问新实例](https://cloud.tencent.com/document/product/236/15864) 触发该流程。
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// 是否是内核子版本升级，支持的值：1 - 升级内核子版本；0 - 升级数据库引擎版本。无默认值，请指定要升级的版本类型。
	UpgradeSubversion *int64 `json:"UpgradeSubversion,omitnil,omitempty" name:"UpgradeSubversion"`

	// 延迟阈值。取值范围：1 - 10。无默认值，不传此参数时，延迟阈值为0，表示延迟阈值不做设置。
	MaxDelayTime *int64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`

	// 5.7升级8.0是否忽略关键字错误，取值范围[0,1]，1表示忽略，0表示不忽略。无默认值，不传此参数表示不做处理。
	IgnoreErrKeyword *int64 `json:"IgnoreErrKeyword,omitnil,omitempty" name:"IgnoreErrKeyword"`

	// 版本升级支持指定参数
	ParamList []*UpgradeEngineVersionParams `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type UpgradeDBInstanceEngineVersionRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主实例数据库引擎版本，支持值包括：5.6、5.7、8.0。
	// 说明：不支持越级升级，升级后不支持降级。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 切换访问新实例的方式，默认为 0。支持值包括：0 - 立刻切换，1 - 时间窗切换；当该值为 1 时，升级过程中，切换访问新实例的流程将会在时间窗内进行，或者用户主动调用接口 [切换访问新实例](https://cloud.tencent.com/document/product/236/15864) 触发该流程。
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// 是否是内核子版本升级，支持的值：1 - 升级内核子版本；0 - 升级数据库引擎版本。无默认值，请指定要升级的版本类型。
	UpgradeSubversion *int64 `json:"UpgradeSubversion,omitnil,omitempty" name:"UpgradeSubversion"`

	// 延迟阈值。取值范围：1 - 10。无默认值，不传此参数时，延迟阈值为0，表示延迟阈值不做设置。
	MaxDelayTime *int64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`

	// 5.7升级8.0是否忽略关键字错误，取值范围[0,1]，1表示忽略，0表示不忽略。无默认值，不传此参数表示不做处理。
	IgnoreErrKeyword *int64 `json:"IgnoreErrKeyword,omitnil,omitempty" name:"IgnoreErrKeyword"`

	// 版本升级支持指定参数
	ParamList []*UpgradeEngineVersionParams `json:"ParamList,omitnil,omitempty" name:"ParamList"`
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
	delete(f, "IgnoreErrKeyword")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceEngineVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceEngineVersionResponseParams struct {
	// 异步任务 ID，可使用 [查询异步任务的执行结果](https://cloud.tencent.com/document/api/236/20410) 获取其执行情况。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升级后的内存大小，单位：MB，为保证传入 Memory 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可升级的内存规格。
	// 说明：如果进行迁移业务，请一定填写实例规格（CPU、内存），不然系统会默认以最小允许规格传参。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 升级后的硬盘大小，单位：GB，为保证传入 Volume 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可升级的硬盘范围。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 数据复制方式，支持值包括：0 - 异步复制，1 - 半同步复制，2 - 强同步复制，升级主实例时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义。
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// 部署模式，默认为 0，支持值包括：0 - 单可用区部署，1 - 多可用区部署，升级主实例时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义。
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 备库1的可用区信息，默认和实例的 Zone 参数一致，升级主实例为多可用区部署时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义。可通过 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口查询支持的可用区。
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// 主实例数据库引擎版本，支持值包括：5.5、5.6、5.7、8.0。
	// 说明：升级数据库版本请使用 [UpgradeDBInstanceEngineVersion](https://cloud.tencent.com/document/api/236/15870) 接口。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 切换访问新实例的方式，默认为 0。支持值包括：0 - 立刻切换，1 - 时间窗切换；当该值为 1 时，升级过程中，切换访问新实例的流程将会在时间窗内进行，或者用户主动调用接口 [切换访问新实例](https://cloud.tencent.com/document/product/236/15864) 触发该流程。
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// 备库2的可用区信息，默认为空，升级主实例时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义。可通过 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口查询支持的可用区。
	// 备注：如您要将三节点降级至双节点，将该参数设置为空值即可实现。
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// 实例类型，默认为 master，支持值包括：master - 表示主实例，dr - 表示灾备实例，ro - 表示只读实例。
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 实例隔离类型。支持值包括： "UNIVERSAL" - 通用型实例， "EXCLUSIVE" - 独享型实例， "BASIC" - 基础版实例。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 升级后的实例cpu核数，如果不传将根据 Memory 指定的内存值自动填充最小允许规格的cpu值。
	// 说明：如果进行迁移业务，请一定填写实例规格（CPU、内存），不然系统会默认以最小允许规格传参。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 是否极速变配。0-普通升级，1-极速变配，2 极速优先。选择极速变配会根据资源状况校验是否可以进行极速变配，满足条件则进行极速变配，不满足条件会返回报错信息。
	FastUpgrade *int64 `json:"FastUpgrade,omitnil,omitempty" name:"FastUpgrade"`

	// 延迟阈值。取值范围1~10，默认值为10。
	MaxDelayTime *int64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`

	// 是否跨区迁移。0-普通迁移，1-跨区迁移，默认值为0。该值为1时支持变更实例主节点可用区。
	CrossCluster *int64 `json:"CrossCluster,omitnil,omitempty" name:"CrossCluster"`

	// 主节点可用区，该值仅在跨区迁移时生效。仅支持同地域下的可用区进行迁移。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 针对跨集群搬迁场景，选择同可用区RO的处理逻辑。together-同可用区RO跟随主实例迁移至目标可用区（默认选项），severally-同可用区RO保持原部署模式、不迁移至目标可用区。
	RoTransType *string `json:"RoTransType,omitnil,omitempty" name:"RoTransType"`

	// 集群版节点拓扑配置。
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// 检查原地升级是否需要重启，1 检查， 0 不检查。如果值为1，检查为原地升级需要重启，则会停止升级并进行返回提示，如果为原地升级不重启，则正常执行升级流程。
	CheckFastUpgradeReboot *int64 `json:"CheckFastUpgradeReboot,omitnil,omitempty" name:"CheckFastUpgradeReboot"`

	// 数据校验敏感度，非极速变配时使用此参数，敏感度根据当前实例规格计算迁移过程中的数据对比使用的cpu资源
	// 对应的选项为: "high"、"normal"、"low"，默认为空
	// 参数详解，：
	// "high": 对应控制台中的高，数据库负载过高不建议使用
	// "normal"：对应控制台中的标准
	// "low"：对应控制台中的低
	DataCheckSensitive *string `json:"DataCheckSensitive,omitnil,omitempty" name:"DataCheckSensitive"`
}

type UpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv 或者 cdbro-c1nl9rpv。与云数据库控制台页面中显示的实例 ID 相同，可使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升级后的内存大小，单位：MB，为保证传入 Memory 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可升级的内存规格。
	// 说明：如果进行迁移业务，请一定填写实例规格（CPU、内存），不然系统会默认以最小允许规格传参。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 升级后的硬盘大小，单位：GB，为保证传入 Volume 值有效，请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口获取可升级的硬盘范围。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 数据复制方式，支持值包括：0 - 异步复制，1 - 半同步复制，2 - 强同步复制，升级主实例时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义。
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// 部署模式，默认为 0，支持值包括：0 - 单可用区部署，1 - 多可用区部署，升级主实例时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义。
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 备库1的可用区信息，默认和实例的 Zone 参数一致，升级主实例为多可用区部署时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义。可通过 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口查询支持的可用区。
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// 主实例数据库引擎版本，支持值包括：5.5、5.6、5.7、8.0。
	// 说明：升级数据库版本请使用 [UpgradeDBInstanceEngineVersion](https://cloud.tencent.com/document/api/236/15870) 接口。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 切换访问新实例的方式，默认为 0。支持值包括：0 - 立刻切换，1 - 时间窗切换；当该值为 1 时，升级过程中，切换访问新实例的流程将会在时间窗内进行，或者用户主动调用接口 [切换访问新实例](https://cloud.tencent.com/document/product/236/15864) 触发该流程。
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// 备库2的可用区信息，默认为空，升级主实例时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义。可通过 [获取云数据库可售卖规格](https://cloud.tencent.com/document/product/236/17229) 接口查询支持的可用区。
	// 备注：如您要将三节点降级至双节点，将该参数设置为空值即可实现。
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// 实例类型，默认为 master，支持值包括：master - 表示主实例，dr - 表示灾备实例，ro - 表示只读实例。
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 实例隔离类型。支持值包括： "UNIVERSAL" - 通用型实例， "EXCLUSIVE" - 独享型实例， "BASIC" - 基础版实例。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 升级后的实例cpu核数，如果不传将根据 Memory 指定的内存值自动填充最小允许规格的cpu值。
	// 说明：如果进行迁移业务，请一定填写实例规格（CPU、内存），不然系统会默认以最小允许规格传参。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 是否极速变配。0-普通升级，1-极速变配，2 极速优先。选择极速变配会根据资源状况校验是否可以进行极速变配，满足条件则进行极速变配，不满足条件会返回报错信息。
	FastUpgrade *int64 `json:"FastUpgrade,omitnil,omitempty" name:"FastUpgrade"`

	// 延迟阈值。取值范围1~10，默认值为10。
	MaxDelayTime *int64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`

	// 是否跨区迁移。0-普通迁移，1-跨区迁移，默认值为0。该值为1时支持变更实例主节点可用区。
	CrossCluster *int64 `json:"CrossCluster,omitnil,omitempty" name:"CrossCluster"`

	// 主节点可用区，该值仅在跨区迁移时生效。仅支持同地域下的可用区进行迁移。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 针对跨集群搬迁场景，选择同可用区RO的处理逻辑。together-同可用区RO跟随主实例迁移至目标可用区（默认选项），severally-同可用区RO保持原部署模式、不迁移至目标可用区。
	RoTransType *string `json:"RoTransType,omitnil,omitempty" name:"RoTransType"`

	// 集群版节点拓扑配置。
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// 检查原地升级是否需要重启，1 检查， 0 不检查。如果值为1，检查为原地升级需要重启，则会停止升级并进行返回提示，如果为原地升级不重启，则正常执行升级流程。
	CheckFastUpgradeReboot *int64 `json:"CheckFastUpgradeReboot,omitnil,omitempty" name:"CheckFastUpgradeReboot"`

	// 数据校验敏感度，非极速变配时使用此参数，敏感度根据当前实例规格计算迁移过程中的数据对比使用的cpu资源
	// 对应的选项为: "high"、"normal"、"low"，默认为空
	// 参数详解，：
	// "high": 对应控制台中的高，数据库负载过高不建议使用
	// "normal"：对应控制台中的标准
	// "low"：对应控制台中的低
	DataCheckSensitive *string `json:"DataCheckSensitive,omitnil,omitempty" name:"DataCheckSensitive"`
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
	delete(f, "RoTransType")
	delete(f, "ClusterTopology")
	delete(f, "CheckFastUpgradeReboot")
	delete(f, "DataCheckSensitive")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceResponseParams struct {
	// 订单 ID。
	DealIds []*string `json:"DealIds,omitnil,omitempty" name:"DealIds"`

	// 异步任务的请求 ID，可使用此 ID [查询异步任务的执行结果](https://cloud.tencent.com/document/product/236/20410)。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type UpgradeEngineVersionParams struct {
	// 参数名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type UploadInfo struct {
	// 文件所有分片数
	AllSliceNum *int64 `json:"AllSliceNum,omitnil,omitempty" name:"AllSliceNum"`

	// 已完成分片数
	CompleteNum *int64 `json:"CompleteNum,omitnil,omitempty" name:"CompleteNum"`
}

// Predefined struct for user
type VerifyRootAccountRequestParams struct {
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例 ROOT 账号的密码。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type VerifyRootAccountRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例 ROOT 账号的密码。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DeployMode []*int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 主实例所在的可用区
	MasterZone []*string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// 实例为多可用区部署时，备库1所在的可用区
	SlaveZone []*string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// 实例为多可用区部署时，备库2所在的可用区
	BackupZone []*string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`
}