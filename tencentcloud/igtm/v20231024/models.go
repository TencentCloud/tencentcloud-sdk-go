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

package v20231024

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Address struct {
	// 地址值：只支持ipv4、ipv6和域名格式；
	// 不支持回环地址、保留地址、内网地址与腾讯保留网段
	Addr *string `json:"Addr,omitnil,omitempty" name:"Addr"`

	// 是否启用:DISABLED不启用；ENABLED启用
	IsEnable *string `json:"IsEnable,omitnil,omitempty" name:"IsEnable"`

	// 地址id
	AddressId *uint64 `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	// 地址名称
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// OK正常，DOWN故障，WARN风险，UNKNOWN探测中，UNMONITORED未知
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 权重，流量策略为WEIGHT时，必填；范围1-100
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 创建时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 修改时间
	UpdatedOn *string `json:"UpdatedOn,omitnil,omitempty" name:"UpdatedOn"`
}

type AddressLocation struct {
	// ip地址
	Addr *string `json:"Addr,omitnil,omitempty" name:"Addr"`

	// 所属地域
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`
}

type AddressPool struct {
	// 地址池 id
	PoolId *uint64 `json:"PoolId,omitnil,omitempty" name:"PoolId"`

	// 地址池名
	PoolName *string `json:"PoolName,omitnil,omitempty" name:"PoolName"`

	// 地址池地址类型：IPV4、IPV6、DOMAIN
	AddrType *string `json:"AddrType,omitnil,omitempty" name:"AddrType"`

	// 流量策略: WEIGHT负载均衡，ALL解析全部
	TrafficStrategy *string `json:"TrafficStrategy,omitnil,omitempty" name:"TrafficStrategy"`

	// 监控器id
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorId *uint64 `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`

	// OK正常，DOWN故障，WARN风险，UNKNOWN未知
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 地址数
	AddressNum *int64 `json:"AddressNum,omitnil,omitempty" name:"AddressNum"`

	// 探点数
	MonitorGroupNum *int64 `json:"MonitorGroupNum,omitnil,omitempty" name:"MonitorGroupNum"`

	// 探测任务数
	MonitorTaskNum *int64 `json:"MonitorTaskNum,omitnil,omitempty" name:"MonitorTaskNum"`

	// 实例相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceInfo []*InstanceInfo `json:"InstanceInfo,omitnil,omitempty" name:"InstanceInfo"`

	// 地址池地址信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressSet []*Address `json:"AddressSet,omitnil,omitempty" name:"AddressSet"`

	// 创建时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 更新时间
	UpdatedOn *string `json:"UpdatedOn,omitnil,omitempty" name:"UpdatedOn"`
}

type AddressPoolDetail struct {
	// 地址池 id
	PoolId *uint64 `json:"PoolId,omitnil,omitempty" name:"PoolId"`

	// 地址池名
	PoolName *string `json:"PoolName,omitnil,omitempty" name:"PoolName"`

	// 地址池地址类型：IPV4、IPV6、DOMAIN
	AddrType *string `json:"AddrType,omitnil,omitempty" name:"AddrType"`

	// 流量策略: WEIGHT负载均衡，ALL解析全部
	TrafficStrategy *string `json:"TrafficStrategy,omitnil,omitempty" name:"TrafficStrategy"`

	// 监控器id
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorId *uint64 `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`

	// 创建时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 更新时间
	UpdatedOn *string `json:"UpdatedOn,omitnil,omitempty" name:"UpdatedOn"`
}

type CostItem struct {
	// 计费项名称
	CostName *string `json:"CostName,omitnil,omitempty" name:"CostName"`

	// 计费项值
	CostValue *uint64 `json:"CostValue,omitnil,omitempty" name:"CostValue"`
}

// Predefined struct for user
type CreateAddressPoolRequestParams struct {
	// 地址池名称，不允许重复
	PoolName *string `json:"PoolName,omitnil,omitempty" name:"PoolName"`

	// 流量策略：WEIGHT负载均衡，ALL解析所有健康地址
	TrafficStrategy *string `json:"TrafficStrategy,omitnil,omitempty" name:"TrafficStrategy"`

	// 地址列表
	AddressSet []*Address `json:"AddressSet,omitnil,omitempty" name:"AddressSet"`

	// 监控器id
	MonitorId *uint64 `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`
}

type CreateAddressPoolRequest struct {
	*tchttp.BaseRequest
	
	// 地址池名称，不允许重复
	PoolName *string `json:"PoolName,omitnil,omitempty" name:"PoolName"`

	// 流量策略：WEIGHT负载均衡，ALL解析所有健康地址
	TrafficStrategy *string `json:"TrafficStrategy,omitnil,omitempty" name:"TrafficStrategy"`

	// 地址列表
	AddressSet []*Address `json:"AddressSet,omitnil,omitempty" name:"AddressSet"`

	// 监控器id
	MonitorId *uint64 `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`
}

func (r *CreateAddressPoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAddressPoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PoolName")
	delete(f, "TrafficStrategy")
	delete(f, "AddressSet")
	delete(f, "MonitorId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAddressPoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAddressPoolResponseParams struct {
	// 地址池id
	AddressPoolId *uint64 `json:"AddressPoolId,omitnil,omitempty" name:"AddressPoolId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAddressPoolResponse struct {
	*tchttp.BaseResponse
	Response *CreateAddressPoolResponseParams `json:"Response"`
}

func (r *CreateAddressPoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAddressPoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceRequestParams struct {
	// 业务域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// CUSTOM: 自定义接入域名
	// SYSTEM: 系统接入域名
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// 解析生效时间	
	GlobalTtl *uint64 `json:"GlobalTtl,omitnil,omitempty" name:"GlobalTtl"`

	// 套餐类型
	// FREE: 免费版
	// STANDARD：标准版
	// ULTIMATE：旗舰版
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 接入主域名
	AccessDomain *string `json:"AccessDomain,omitnil,omitempty" name:"AccessDomain"`

	// 接入子域名
	AccessSubDomain *string `json:"AccessSubDomain,omitnil,omitempty" name:"AccessSubDomain"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 套餐资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 业务域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// CUSTOM: 自定义接入域名
	// SYSTEM: 系统接入域名
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// 解析生效时间	
	GlobalTtl *uint64 `json:"GlobalTtl,omitnil,omitempty" name:"GlobalTtl"`

	// 套餐类型
	// FREE: 免费版
	// STANDARD：标准版
	// ULTIMATE：旗舰版
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 接入主域名
	AccessDomain *string `json:"AccessDomain,omitnil,omitempty" name:"AccessDomain"`

	// 接入子域名
	AccessSubDomain *string `json:"AccessSubDomain,omitnil,omitempty" name:"AccessSubDomain"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 套餐资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "AccessType")
	delete(f, "GlobalTtl")
	delete(f, "PackageType")
	delete(f, "InstanceName")
	delete(f, "AccessDomain")
	delete(f, "AccessSubDomain")
	delete(f, "Remark")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceResponseParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceResponseParams `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMonitorRequestParams struct {
	// 监控器名称
	MonitorName *string `json:"MonitorName,omitnil,omitempty" name:"MonitorName"`

	// 探测协议，可选值 PING TCP HTTP HTTPS
	CheckProtocol *string `json:"CheckProtocol,omitnil,omitempty" name:"CheckProtocol"`

	// 检查间隔（秒），可选值有15 60 120 300
	CheckInterval *uint64 `json:"CheckInterval,omitnil,omitempty" name:"CheckInterval"`

	// 超时时间，单位秒，可选值为2 3 5 10
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 重试次数，可选值为 0，1，2
	FailTimes *uint64 `json:"FailTimes,omitnil,omitempty" name:"FailTimes"`

	// 失败比例，取值为 20 30 40 50 60 70 80 100，默认值50
	FailRate *uint64 `json:"FailRate,omitnil,omitempty" name:"FailRate"`

	// 监控节点类型，可选值有AUTO INTERNAL OVERSEAS IPV6 ALL
	DetectorStyle *string `json:"DetectorStyle,omitnil,omitempty" name:"DetectorStyle"`

	// 探测器组id列表以,分隔
	DetectorGroupIds []*uint64 `json:"DetectorGroupIds,omitnil,omitempty" name:"DetectorGroupIds"`

	// PING 包数目，当CheckProtocol=ping时必填，可选值有20 50 100
	PingNum *uint64 `json:"PingNum,omitnil,omitempty" name:"PingNum"`

	// 检查端口，可选值在1-65535之间
	TcpPort *uint64 `json:"TcpPort,omitnil,omitempty" name:"TcpPort"`

	// Host 设置，默认为业务域名
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// URL 路径，默认为“/”
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 返回错误码阈值, 可选值为 400 和 500, 默认值500
	ReturnCodeThreshold *uint64 `json:"ReturnCodeThreshold,omitnil,omitempty" name:"ReturnCodeThreshold"`

	// 跟随 3XX 重定向 ，不开启为 DISABLED， 开启为 ENABLED，默认不开启
	EnableRedirect *string `json:"EnableRedirect,omitnil,omitempty" name:"EnableRedirect"`

	// 启用 SNI，不开启为 DISABLED， 开启为 ENABLED，默认不开启
	EnableSni *string `json:"EnableSni,omitnil,omitempty" name:"EnableSni"`

	// 丢包率告警阈值，当CheckProtocol=ping时必填取值为10 30 50 80 90 100
	PacketLossRate *uint64 `json:"PacketLossRate,omitnil,omitempty" name:"PacketLossRate"`

	// 持续周期数，可选值1-5
	ContinuePeriod *uint64 `json:"ContinuePeriod,omitnil,omitempty" name:"ContinuePeriod"`
}

type CreateMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 监控器名称
	MonitorName *string `json:"MonitorName,omitnil,omitempty" name:"MonitorName"`

	// 探测协议，可选值 PING TCP HTTP HTTPS
	CheckProtocol *string `json:"CheckProtocol,omitnil,omitempty" name:"CheckProtocol"`

	// 检查间隔（秒），可选值有15 60 120 300
	CheckInterval *uint64 `json:"CheckInterval,omitnil,omitempty" name:"CheckInterval"`

	// 超时时间，单位秒，可选值为2 3 5 10
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 重试次数，可选值为 0，1，2
	FailTimes *uint64 `json:"FailTimes,omitnil,omitempty" name:"FailTimes"`

	// 失败比例，取值为 20 30 40 50 60 70 80 100，默认值50
	FailRate *uint64 `json:"FailRate,omitnil,omitempty" name:"FailRate"`

	// 监控节点类型，可选值有AUTO INTERNAL OVERSEAS IPV6 ALL
	DetectorStyle *string `json:"DetectorStyle,omitnil,omitempty" name:"DetectorStyle"`

	// 探测器组id列表以,分隔
	DetectorGroupIds []*uint64 `json:"DetectorGroupIds,omitnil,omitempty" name:"DetectorGroupIds"`

	// PING 包数目，当CheckProtocol=ping时必填，可选值有20 50 100
	PingNum *uint64 `json:"PingNum,omitnil,omitempty" name:"PingNum"`

	// 检查端口，可选值在1-65535之间
	TcpPort *uint64 `json:"TcpPort,omitnil,omitempty" name:"TcpPort"`

	// Host 设置，默认为业务域名
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// URL 路径，默认为“/”
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 返回错误码阈值, 可选值为 400 和 500, 默认值500
	ReturnCodeThreshold *uint64 `json:"ReturnCodeThreshold,omitnil,omitempty" name:"ReturnCodeThreshold"`

	// 跟随 3XX 重定向 ，不开启为 DISABLED， 开启为 ENABLED，默认不开启
	EnableRedirect *string `json:"EnableRedirect,omitnil,omitempty" name:"EnableRedirect"`

	// 启用 SNI，不开启为 DISABLED， 开启为 ENABLED，默认不开启
	EnableSni *string `json:"EnableSni,omitnil,omitempty" name:"EnableSni"`

	// 丢包率告警阈值，当CheckProtocol=ping时必填取值为10 30 50 80 90 100
	PacketLossRate *uint64 `json:"PacketLossRate,omitnil,omitempty" name:"PacketLossRate"`

	// 持续周期数，可选值1-5
	ContinuePeriod *uint64 `json:"ContinuePeriod,omitnil,omitempty" name:"ContinuePeriod"`
}

func (r *CreateMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MonitorName")
	delete(f, "CheckProtocol")
	delete(f, "CheckInterval")
	delete(f, "Timeout")
	delete(f, "FailTimes")
	delete(f, "FailRate")
	delete(f, "DetectorStyle")
	delete(f, "DetectorGroupIds")
	delete(f, "PingNum")
	delete(f, "TcpPort")
	delete(f, "Host")
	delete(f, "Path")
	delete(f, "ReturnCodeThreshold")
	delete(f, "EnableRedirect")
	delete(f, "EnableSni")
	delete(f, "PacketLossRate")
	delete(f, "ContinuePeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMonitorResponseParams struct {
	// 监控器id
	MonitorId *uint64 `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMonitorResponse struct {
	*tchttp.BaseResponse
	Response *CreateMonitorResponseParams `json:"Response"`
}

func (r *CreateMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePackageAndPayRequestParams struct {
	// 下单类型：CREATE 新购；RENEW 续费；MODIFY 变配
	DealType *string `json:"DealType,omitnil,omitempty" name:"DealType"`

	// 套餐类型：STANDARD 标准版；ULTIMATE 旗舰版；TASK 任务探测
	GoodsType *string `json:"GoodsType,omitnil,omitempty" name:"GoodsType"`

	// 商品数量：STANDARD和ULTIMATE固定为1，TASK为任务探测数量。取值范围：1～10000
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 自动续费：1 开启自动续费；2 关闭自动续费
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 资源ID，续费和变配的时候需要传
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 套餐时长，以月为单位，创建和续费的时候需要传。取值范围：1～120
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 升级的套餐类型，暂时只支持传ULTIMATE，不支持降配
	NewPackageType *string `json:"NewPackageType,omitnil,omitempty" name:"NewPackageType"`

	// 是否自动选择代金券，1 是；0否，默认为0
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

type CreatePackageAndPayRequest struct {
	*tchttp.BaseRequest
	
	// 下单类型：CREATE 新购；RENEW 续费；MODIFY 变配
	DealType *string `json:"DealType,omitnil,omitempty" name:"DealType"`

	// 套餐类型：STANDARD 标准版；ULTIMATE 旗舰版；TASK 任务探测
	GoodsType *string `json:"GoodsType,omitnil,omitempty" name:"GoodsType"`

	// 商品数量：STANDARD和ULTIMATE固定为1，TASK为任务探测数量。取值范围：1～10000
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 自动续费：1 开启自动续费；2 关闭自动续费
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 资源ID，续费和变配的时候需要传
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 套餐时长，以月为单位，创建和续费的时候需要传。取值范围：1～120
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 升级的套餐类型，暂时只支持传ULTIMATE，不支持降配
	NewPackageType *string `json:"NewPackageType,omitnil,omitempty" name:"NewPackageType"`

	// 是否自动选择代金券，1 是；0否，默认为0
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

func (r *CreatePackageAndPayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePackageAndPayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealType")
	delete(f, "GoodsType")
	delete(f, "GoodsNum")
	delete(f, "AutoRenew")
	delete(f, "ResourceId")
	delete(f, "TimeSpan")
	delete(f, "NewPackageType")
	delete(f, "AutoVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePackageAndPayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePackageAndPayResponseParams struct {
	// 资源id列表，目前只会返回一个资源，取第一个值即可
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePackageAndPayResponse struct {
	*tchttp.BaseResponse
	Response *CreatePackageAndPayResponseParams `json:"Response"`
}

func (r *CreatePackageAndPayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePackageAndPayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStrategyRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略名称，不允许重复
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 解析线路
	Source []*Source `json:"Source,omitnil,omitempty" name:"Source"`

	// 主力地址池集合，最多允许配置四级
	MainAddressPoolSet []*MainAddressPool `json:"MainAddressPoolSet,omitnil,omitempty" name:"MainAddressPoolSet"`

	// 兜底地址池集合，只允许配置一级，且地址池数量为1
	FallbackAddressPoolSet []*MainAddressPool `json:"FallbackAddressPoolSet,omitnil,omitempty" name:"FallbackAddressPoolSet"`

	// 是否开启策略强制保留默认线路 disabled, enabled，默认不开启且只有一个策略能开启
	KeepDomainRecords *string `json:"KeepDomainRecords,omitnil,omitempty" name:"KeepDomainRecords"`

	// 策略调度模式：AUTO默认切换；STOP仅暂停不切换
	SwitchPoolType *string `json:"SwitchPoolType,omitnil,omitempty" name:"SwitchPoolType"`
}

type CreateStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略名称，不允许重复
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 解析线路
	Source []*Source `json:"Source,omitnil,omitempty" name:"Source"`

	// 主力地址池集合，最多允许配置四级
	MainAddressPoolSet []*MainAddressPool `json:"MainAddressPoolSet,omitnil,omitempty" name:"MainAddressPoolSet"`

	// 兜底地址池集合，只允许配置一级，且地址池数量为1
	FallbackAddressPoolSet []*MainAddressPool `json:"FallbackAddressPoolSet,omitnil,omitempty" name:"FallbackAddressPoolSet"`

	// 是否开启策略强制保留默认线路 disabled, enabled，默认不开启且只有一个策略能开启
	KeepDomainRecords *string `json:"KeepDomainRecords,omitnil,omitempty" name:"KeepDomainRecords"`

	// 策略调度模式：AUTO默认切换；STOP仅暂停不切换
	SwitchPoolType *string `json:"SwitchPoolType,omitnil,omitempty" name:"SwitchPoolType"`
}

func (r *CreateStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StrategyName")
	delete(f, "Source")
	delete(f, "MainAddressPoolSet")
	delete(f, "FallbackAddressPoolSet")
	delete(f, "KeepDomainRecords")
	delete(f, "SwitchPoolType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStrategyResponseParams struct {
	// 新增策略id
	StrategyId *int64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStrategyResponse struct {
	*tchttp.BaseResponse
	Response *CreateStrategyResponseParams `json:"Response"`
}

func (r *CreateStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAddressPoolRequestParams struct {
	// 地址池id
	PoolId *uint64 `json:"PoolId,omitnil,omitempty" name:"PoolId"`
}

type DeleteAddressPoolRequest struct {
	*tchttp.BaseRequest
	
	// 地址池id
	PoolId *uint64 `json:"PoolId,omitnil,omitempty" name:"PoolId"`
}

func (r *DeleteAddressPoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAddressPoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PoolId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAddressPoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAddressPoolResponseParams struct {
	// 是否成功
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAddressPoolResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAddressPoolResponseParams `json:"Response"`
}

func (r *DeleteAddressPoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAddressPoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMonitorRequestParams struct {
	// 监控器id
	MonitorId *uint64 `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`
}

type DeleteMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 监控器id
	MonitorId *uint64 `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`
}

func (r *DeleteMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MonitorId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMonitorResponseParams struct {
	// 成功返回
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMonitorResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMonitorResponseParams `json:"Response"`
}

func (r *DeleteMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStrategyRequestParams struct {
	// 策略id
	StrategyId *uint64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 策略id
	StrategyId *uint64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StrategyId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStrategyResponseParams struct {
	// 是否成功
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStrategyResponseParams `json:"Response"`
}

func (r *DeleteStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressLocationRequestParams struct {
	// 地址数组
	Address []*string `json:"Address,omitnil,omitempty" name:"Address"`
}

type DescribeAddressLocationRequest struct {
	*tchttp.BaseRequest
	
	// 地址数组
	Address []*string `json:"Address,omitnil,omitempty" name:"Address"`
}

func (r *DescribeAddressLocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressLocationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Address")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddressLocationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressLocationResponseParams struct {
	// 所属地域
	AddressLocation []*AddressLocation `json:"AddressLocation,omitnil,omitempty" name:"AddressLocation"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAddressLocationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAddressLocationResponseParams `json:"Response"`
}

func (r *DescribeAddressLocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressLocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressPoolDetailRequestParams struct {
	// 地址池id
	PoolId *int64 `json:"PoolId,omitnil,omitempty" name:"PoolId"`
}

type DescribeAddressPoolDetailRequest struct {
	*tchttp.BaseRequest
	
	// 地址池id
	PoolId *int64 `json:"PoolId,omitnil,omitempty" name:"PoolId"`
}

func (r *DescribeAddressPoolDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressPoolDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PoolId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddressPoolDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressPoolDetailResponseParams struct {
	// 资源组详情描述
	AddressPool *AddressPoolDetail `json:"AddressPool,omitnil,omitempty" name:"AddressPool"`

	// 资源组中的资源列表
	AddressSet []*Address `json:"AddressSet,omitnil,omitempty" name:"AddressSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAddressPoolDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAddressPoolDetailResponseParams `json:"Response"`
}

func (r *DescribeAddressPoolDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressPoolDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressPoolListRequestParams struct {
	// 告警过滤条件：PoolName：地址池名称；MonitorId：监控器id
	Filters []*ResourceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 页数
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAddressPoolListRequest struct {
	*tchttp.BaseRequest
	
	// 告警过滤条件：PoolName：地址池名称；MonitorId：监控器id
	Filters []*ResourceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 页数
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAddressPoolListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressPoolListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddressPoolListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressPoolListResponseParams struct {
	// 资源组列表
	AddressPoolSet []*AddressPool `json:"AddressPoolSet,omitnil,omitempty" name:"AddressPoolSet"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAddressPoolListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAddressPoolListResponseParams `json:"Response"`
}

func (r *DescribeAddressPoolListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressPoolListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDetectPackageDetailRequestParams struct {
	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type DescribeDetectPackageDetailRequest struct {
	*tchttp.BaseRequest
	
	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *DescribeDetectPackageDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDetectPackageDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDetectPackageDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDetectPackageDetailResponseParams struct {
	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源类型 TASK 探测任务
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 额度
	Quota *uint64 `json:"Quota,omitnil,omitempty" name:"Quota"`

	// 过期时间
	CurrentDeadline *string `json:"CurrentDeadline,omitnil,omitempty" name:"CurrentDeadline"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否过期
	IsExpire *uint64 `json:"IsExpire,omitnil,omitempty" name:"IsExpire"`

	// 状态 ENABLED: 正常 ISOLATED: 隔离 DESTROYED：销毁 REFUNDED：已退款
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否自动续费：0否1是
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 计费项
	CostItemList []*CostItem `json:"CostItemList,omitnil,omitempty" name:"CostItemList"`

	// 使用数量
	UsedNum *uint64 `json:"UsedNum,omitnil,omitempty" name:"UsedNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDetectPackageDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDetectPackageDetailResponseParams `json:"Response"`
}

func (r *DescribeDetectPackageDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDetectPackageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDetectTaskPackageListRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 探测任务过滤条件：ResourceId 探测任务的资源id，PeriodStart 最小过期时间,PeriodEnd 最大过期时间
	Filters []*ResourceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeDetectTaskPackageListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 探测任务过滤条件：ResourceId 探测任务的资源id，PeriodStart 最小过期时间,PeriodEnd 最大过期时间
	Filters []*ResourceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeDetectTaskPackageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDetectTaskPackageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDetectTaskPackageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDetectTaskPackageListResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 探测任务套餐列表
	TaskPackageSet []*DetectTaskPackage `json:"TaskPackageSet,omitnil,omitempty" name:"TaskPackageSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDetectTaskPackageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDetectTaskPackageListResponseParams `json:"Response"`
}

func (r *DescribeDetectTaskPackageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDetectTaskPackageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDetectorsRequestParams struct {

}

type DescribeDetectorsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDetectorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDetectorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDetectorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDetectorsResponseParams struct {
	// 探测器组列表
	DetectorGroupSet []*DetectorGroup `json:"DetectorGroupSet,omitnil,omitempty" name:"DetectorGroupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDetectorsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDetectorsResponseParams `json:"Response"`
}

func (r *DescribeDetectorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDetectorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDnsLineListRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDnsLineListRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDnsLineListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDnsLineListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDnsLineListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDnsLineListResponseParams struct {
	// 地址池列表
	DnsLineSet []*GroupLine `json:"DnsLineSet,omitnil,omitempty" name:"DnsLineSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDnsLineListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDnsLineListResponseParams `json:"Response"`
}

func (r *DescribeDnsLineListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDnsLineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDetailRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
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
	Instance *InstanceDetail `json:"Instance,omitnil,omitempty" name:"Instance"`

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
type DescribeInstanceListRequestParams struct {
	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Name: "实例名称" 模糊查询Domain: "域名" 模糊查询MonitorId: "监控器 id" PoolId: "地址池id", AccessDomain接入主域名
	Filters []*ResourceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Name: "实例名称" 模糊查询Domain: "域名" 模糊查询MonitorId: "监控器 id" PoolId: "地址池id", AccessDomain接入主域名
	Filters []*ResourceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceListResponseParams struct {
	// 实例列表
	InstanceSet []*Instance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// 列表总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 是否支持系统域名接入：true支持；false不支持
	SystemAccessEnabled *bool `json:"SystemAccessEnabled,omitnil,omitempty" name:"SystemAccessEnabled"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceListResponseParams `json:"Response"`
}

func (r *DescribeInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancePackageListRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// InstanceId实例Id，InstanceName实例名称，ResourceId套餐Id，PackageType套餐类型 
	Filters []*ResourceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否使用：0未使用1已使用
	IsUsed *uint64 `json:"IsUsed,omitnil,omitempty" name:"IsUsed"`
}

type DescribeInstancePackageListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// InstanceId实例Id，InstanceName实例名称，ResourceId套餐Id，PackageType套餐类型 
	Filters []*ResourceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否使用：0未使用1已使用
	IsUsed *uint64 `json:"IsUsed,omitnil,omitempty" name:"IsUsed"`
}

func (r *DescribeInstancePackageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancePackageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "IsUsed")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancePackageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancePackageListResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例套餐列表
	InstanceSet []*InstancePackage `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancePackageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancePackageListResponseParams `json:"Response"`
}

func (r *DescribeInstancePackageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancePackageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonitorDetailRequestParams struct {
	// 监控器id
	MonitorId *uint64 `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`
}

type DescribeMonitorDetailRequest struct {
	*tchttp.BaseRequest
	
	// 监控器id
	MonitorId *uint64 `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`
}

func (r *DescribeMonitorDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonitorDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MonitorId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMonitorDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonitorDetailResponseParams struct {
	// 探测规则
	MonitorDetail *MonitorDetail `json:"MonitorDetail,omitnil,omitempty" name:"MonitorDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMonitorDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMonitorDetailResponseParams `json:"Response"`
}

func (r *DescribeMonitorDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonitorDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonitorsRequestParams struct {
	// 分页，偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页，当前分页记录数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询过滤条件：MonitorName：监控器名称；MonitorId：监控器id
	Filters []*ResourceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否查探测次数0否1是
	IsDetectNum *uint64 `json:"IsDetectNum,omitnil,omitempty" name:"IsDetectNum"`
}

type DescribeMonitorsRequest struct {
	*tchttp.BaseRequest
	
	// 分页，偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页，当前分页记录数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询过滤条件：MonitorName：监控器名称；MonitorId：监控器id
	Filters []*ResourceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否查探测次数0否1是
	IsDetectNum *uint64 `json:"IsDetectNum,omitnil,omitempty" name:"IsDetectNum"`
}

func (r *DescribeMonitorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonitorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "IsDetectNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMonitorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonitorsResponseParams struct {
	// 监控器列表
	MonitorDataSet []*MonitorDetail `json:"MonitorDataSet,omitnil,omitempty" name:"MonitorDataSet"`

	// 数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMonitorsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMonitorsResponseParams `json:"Response"`
}

func (r *DescribeMonitorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonitorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuotasRequestParams struct {
	// 接入域名
	AccessDomain *string `json:"AccessDomain,omitnil,omitempty" name:"AccessDomain"`
}

type DescribeQuotasRequest struct {
	*tchttp.BaseRequest
	
	// 接入域名
	AccessDomain *string `json:"AccessDomain,omitnil,omitempty" name:"AccessDomain"`
}

func (r *DescribeQuotasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQuotasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuotasResponseParams struct {
	// 配额id
	Quotas *Quota `json:"Quotas,omitnil,omitempty" name:"Quotas"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQuotasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQuotasResponseParams `json:"Response"`
}

func (r *DescribeQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStrategyDetailRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略 id
	StrategyId *int64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

type DescribeStrategyDetailRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略 id
	StrategyId *int64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

func (r *DescribeStrategyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStrategyDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StrategyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStrategyDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStrategyDetailResponseParams struct {
	// 策略详情
	StrategyDetail *StrategyDetail `json:"StrategyDetail,omitnil,omitempty" name:"StrategyDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStrategyDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStrategyDetailResponseParams `json:"Response"`
}

func (r *DescribeStrategyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStrategyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStrategyListRequestParams struct {
	// 实例 id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 策略过滤条件：StrategyName：策略名称
	Filters []*ResourceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeStrategyListRequest struct {
	*tchttp.BaseRequest
	
	// 实例 id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 策略过滤条件：StrategyName：策略名称
	Filters []*ResourceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeStrategyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStrategyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStrategyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStrategyListResponseParams struct {
	// 策略列表
	StrategySet []*Strategy `json:"StrategySet,omitnil,omitempty" name:"StrategySet"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStrategyListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStrategyListResponseParams `json:"Response"`
}

func (r *DescribeStrategyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStrategyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetectTaskPackage struct {
	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源类型
	// TASK 探测任务
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 额度
	Quota *uint64 `json:"Quota,omitnil,omitempty" name:"Quota"`

	// 套餐过期时间
	CurrentDeadline *string `json:"CurrentDeadline,omitnil,omitempty" name:"CurrentDeadline"`

	// 套餐创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否过期0否1是
	IsExpire *uint64 `json:"IsExpire,omitnil,omitempty" name:"IsExpire"`

	// 状态
	// ENABLED: 正常
	// ISOLATED: 隔离
	// DESTROYED：销毁
	// REFUNDED：已退款
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否自动续费0不1是
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 计费项
	CostItemList []*CostItem `json:"CostItemList,omitnil,omitempty" name:"CostItemList"`

	// 探测任务类型：100系统设定；200计费；300管理系统；110D监控迁移的免费任务；120容灾切换任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	Group *uint64 `json:"Group,omitnil,omitempty" name:"Group"`
}

type DetectorGroup struct {
	// 线路组id GroupLineId
	Gid *uint64 `json:"Gid,omitnil,omitempty" name:"Gid"`

	// bgp, international, isp
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 组名
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// ipv4, ipv6
	InternetFamily *string `json:"InternetFamily,omitnil,omitempty" name:"InternetFamily"`

	// 支持的套餐类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageSet []*string `json:"PackageSet,omitnil,omitempty" name:"PackageSet"`
}

type GroupLine struct {
	// 分组线路id
	DnsLineId *uint64 `json:"DnsLineId,omitnil,omitempty" name:"DnsLineId"`

	// 父节点 0为根节点
	Parent *uint64 `json:"Parent,omitnil,omitempty" name:"Parent"`

	// 线路名
	LineName *string `json:"LineName,omitnil,omitempty" name:"LineName"`

	// 10=9 DNSPod 线路 id
	LineId *string `json:"LineId,omitnil,omitempty" name:"LineId"`

	// 是否已使用过
	Useful *bool `json:"Useful,omitnil,omitempty" name:"Useful"`

	// 0为未使用
	SubGroup *uint64 `json:"SubGroup,omitnil,omitempty" name:"SubGroup"`

	// 权限标识
	LinePackage *uint64 `json:"LinePackage,omitnil,omitempty" name:"LinePackage"`

	// 1
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type Instance struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 资源 id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 业务域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Cname域名接入方式
	// CUSTOM: 自定义接入域名
	// SYSTEM: 系统接入域名
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// 接入域名
	AccessDomain *string `json:"AccessDomain,omitnil,omitempty" name:"AccessDomain"`

	// 接入子域名
	AccessSubDomain *string `json:"AccessSubDomain,omitnil,omitempty" name:"AccessSubDomain"`

	// 全局记录过期时间
	GlobalTtl *int64 `json:"GlobalTtl,omitnil,omitempty" name:"GlobalTtl"`

	// 套餐类型
	// FREE: 免费版
	// STANDARD：标准版
	// ULTIMATE：旗舰版
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 实例运行状态
	// NORMAL: 健康
	// FAULTY: 有风险
	// DOWN: 宕机
	// UNKNOWN: 未知
	WorkingStatus *string `json:"WorkingStatus,omitnil,omitempty" name:"WorkingStatus"`

	// 实例状态，ENABLED: 正常，DISABLED: 禁用
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否cname接入：true已接入；false未接入
	IsCnameConfigured *bool `json:"IsCnameConfigured,omitnil,omitempty" name:"IsCnameConfigured"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 策略数量
	StrategyNum *int64 `json:"StrategyNum,omitnil,omitempty" name:"StrategyNum"`

	// 绑定地址池个数
	AddressPoolNum *int64 `json:"AddressPoolNum,omitnil,omitempty" name:"AddressPoolNum"`

	// 绑定监控器数量
	MonitorNum *int64 `json:"MonitorNum,omitnil,omitempty" name:"MonitorNum"`

	// 地址池id
	PoolId *uint64 `json:"PoolId,omitnil,omitempty" name:"PoolId"`

	// 地址池名称
	PoolName *string `json:"PoolName,omitnil,omitempty" name:"PoolName"`

	// 实例创建时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 实例更新时间
	UpdatedOn *string `json:"UpdatedOn,omitnil,omitempty" name:"UpdatedOn"`
}

type InstanceConfig struct {
	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 业务域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// CUSTOM: 自定义接入域名，SYSTEM: 系统接入域名
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 全局记录过期时间	
	GlobalTtl *int64 `json:"GlobalTtl,omitnil,omitempty" name:"GlobalTtl"`

	// 接入主域名
	AccessDomain *string `json:"AccessDomain,omitnil,omitempty" name:"AccessDomain"`

	// 接入子域名
	AccessSubDomain *string `json:"AccessSubDomain,omitnil,omitempty" name:"AccessSubDomain"`
}

type InstanceDetail struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 业务域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Cname域名接入方式
	// CUSTOM: 自定义接入域名
	// SYSTEM: 系统接入域名
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// 接入子域名	
	AccessSubDomain *string `json:"AccessSubDomain,omitnil,omitempty" name:"AccessSubDomain"`

	// 接入域名	
	AccessDomain *string `json:"AccessDomain,omitnil,omitempty" name:"AccessDomain"`

	// 解析生效时间
	GlobalTtl *uint64 `json:"GlobalTtl,omitnil,omitempty" name:"GlobalTtl"`

	// 套餐类型
	// FREE: 免费版
	// STANDARD：标准版
	// ULTIMATE：旗舰版
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 实例运行状态
	// NORMAL: 健康
	// FAULTY: 有风险
	// DOWN: 宕机
	// UNKNOWN: 未知
	WorkingStatus *string `json:"WorkingStatus,omitnil,omitempty" name:"WorkingStatus"`

	// 实例状态，ENABLED: 正常；DISABLED: 禁用
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// cname是否接入：true已接入；false未接入
	IsCnameConfigured *bool `json:"IsCnameConfigured,omitnil,omitempty" name:"IsCnameConfigured"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 策略数量
	StrategyNum *int64 `json:"StrategyNum,omitnil,omitempty" name:"StrategyNum"`

	// 绑定地址池个数
	AddressPoolNum *int64 `json:"AddressPoolNum,omitnil,omitempty" name:"AddressPoolNum"`

	// 绑定监控器数量
	MonitorNum *int64 `json:"MonitorNum,omitnil,omitempty" name:"MonitorNum"`

	// 实例绑定套餐资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 订阅事件列表
	NotifyEventSet []*string `json:"NotifyEventSet,omitnil,omitempty" name:"NotifyEventSet"`

	// 实例创建时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 实例更新时间
	UpdatedOn *string `json:"UpdatedOn,omitnil,omitempty" name:"UpdatedOn"`
}

type InstanceInfo struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type InstancePackage struct {
	// 实例套餐资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 套餐类型
	// FREE: 免费版
	// STANDARD：标准版
	// ULTIMATE：旗舰版
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 套餐过期时间
	CurrentDeadline *string `json:"CurrentDeadline,omitnil,omitempty" name:"CurrentDeadline"`

	// 套餐创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否过期0否1是
	IsExpire *uint64 `json:"IsExpire,omitnil,omitempty" name:"IsExpire"`

	// 实例状态
	// ENABLED: 正常
	// DISABLED: 禁用
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否自动续费0不1是
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 计费项
	CostItemList []*CostItem `json:"CostItemList,omitnil,omitempty" name:"CostItemList"`

	// 最小检查间隔时间s
	MinCheckInterval *uint64 `json:"MinCheckInterval,omitnil,omitempty" name:"MinCheckInterval"`

	// 最小TTL s
	MinGlobalTtl *uint64 `json:"MinGlobalTtl,omitnil,omitempty" name:"MinGlobalTtl"`

	// 流量策略类型：ALL返回全部，WEIGHT权重
	TrafficStrategy []*string `json:"TrafficStrategy,omitnil,omitempty" name:"TrafficStrategy"`

	// 策略类型：LOCATION按地理位置调度，DELAY按延迟调度
	ScheduleStrategy []*string `json:"ScheduleStrategy,omitnil,omitempty" name:"ScheduleStrategy"`
}

type MainAddressPool struct {
	// 集合中的地址池id与权重，数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressPools []*MainPoolWeight `json:"AddressPools,omitnil,omitempty" name:"AddressPools"`

	// 地址池集合id
	MainAddressPoolId *uint64 `json:"MainAddressPoolId,omitnil,omitempty" name:"MainAddressPoolId"`

	// 切换阀值，不能大于主力集合内地址总数
	MinSurviveNum *uint64 `json:"MinSurviveNum,omitnil,omitempty" name:"MinSurviveNum"`

	// 切换策略:ALL解析所有地址；WEIGHT：负载均衡。当为ALL时，解析地址的权重值为1；当为WEIGHT时；权重为地址池权重*地址权重
	TrafficStrategy *string `json:"TrafficStrategy,omitnil,omitempty" name:"TrafficStrategy"`
}

type MainPoolWeight struct {
	// 地址池id
	PoolId *uint64 `json:"PoolId,omitnil,omitempty" name:"PoolId"`

	// 权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

// Predefined struct for user
type ModifyAddressPoolRequestParams struct {
	// 地址池id
	PoolId *uint64 `json:"PoolId,omitnil,omitempty" name:"PoolId"`

	// 地址池名称，不允许重复
	PoolName *string `json:"PoolName,omitnil,omitempty" name:"PoolName"`

	// 流量策略: WEIGHT负载均衡，ALL解析全部
	TrafficStrategy *string `json:"TrafficStrategy,omitnil,omitempty" name:"TrafficStrategy"`

	// 监控器id，当监控器已关联策略时，此字段必传
	MonitorId *uint64 `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`

	// 地址列表，全量更新逻辑，对于存量不需要修改的地址信息也需要带上(其中参数里的AddressId需传入正确的值)，否则会被删除。
	AddressSet []*Address `json:"AddressSet,omitnil,omitempty" name:"AddressSet"`
}

type ModifyAddressPoolRequest struct {
	*tchttp.BaseRequest
	
	// 地址池id
	PoolId *uint64 `json:"PoolId,omitnil,omitempty" name:"PoolId"`

	// 地址池名称，不允许重复
	PoolName *string `json:"PoolName,omitnil,omitempty" name:"PoolName"`

	// 流量策略: WEIGHT负载均衡，ALL解析全部
	TrafficStrategy *string `json:"TrafficStrategy,omitnil,omitempty" name:"TrafficStrategy"`

	// 监控器id，当监控器已关联策略时，此字段必传
	MonitorId *uint64 `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`

	// 地址列表，全量更新逻辑，对于存量不需要修改的地址信息也需要带上(其中参数里的AddressId需传入正确的值)，否则会被删除。
	AddressSet []*Address `json:"AddressSet,omitnil,omitempty" name:"AddressSet"`
}

func (r *ModifyAddressPoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressPoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PoolId")
	delete(f, "PoolName")
	delete(f, "TrafficStrategy")
	delete(f, "MonitorId")
	delete(f, "AddressSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAddressPoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressPoolResponseParams struct {
	// 是否修改成功
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAddressPoolResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAddressPoolResponseParams `json:"Response"`
}

func (r *ModifyAddressPoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressPoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceConfigRequestParams struct {
	// 实例id	
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例配置详情
	InstanceConfig *InstanceConfig `json:"InstanceConfig,omitnil,omitempty" name:"InstanceConfig"`
}

type ModifyInstanceConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例id	
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例配置详情
	InstanceConfig *InstanceConfig `json:"InstanceConfig,omitnil,omitempty" name:"InstanceConfig"`
}

func (r *ModifyInstanceConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceConfigResponseParams struct {
	// 实例详情
	Instance *InstanceDetail `json:"Instance,omitnil,omitempty" name:"Instance"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceConfigResponseParams `json:"Response"`
}

func (r *ModifyInstanceConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMonitorRequestParams struct {
	// 监控器id
	MonitorId *uint64 `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`

	// 监控器名称
	MonitorName *string `json:"MonitorName,omitnil,omitempty" name:"MonitorName"`

	// 检查协议，可选值 PING TCP HTTP HTTPS
	CheckProtocol *string `json:"CheckProtocol,omitnil,omitempty" name:"CheckProtocol"`

	// 检查间隔（秒），可选值 15 60 120 300
	CheckInterval *uint64 `json:"CheckInterval,omitnil,omitempty" name:"CheckInterval"`

	// 超时时间，单位:秒，可选值 2  3  5  10
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 重试次数，可选值为 0，1，2
	FailTimes *uint64 `json:"FailTimes,omitnil,omitempty" name:"FailTimes"`

	// 失败比例，可选值为 20 30 40 50 60 70 80 100，默认值为50
	FailRate *uint64 `json:"FailRate,omitnil,omitempty" name:"FailRate"`

	// 监控节点类型，可选值有AUTO INTERNAL OVERSEAS IPV6 ALL
	DetectorStyle *string `json:"DetectorStyle,omitnil,omitempty" name:"DetectorStyle"`

	// 探测器组id列表
	DetectorGroupIds []*uint64 `json:"DetectorGroupIds,omitnil,omitempty" name:"DetectorGroupIds"`

	// PING 包数目， 当CheckProtocol=ping时必填，可选值 20 50 100
	PingNum *uint64 `json:"PingNum,omitnil,omitempty" name:"PingNum"`

	// 检查端口，可选值为1-65535之间的整数
	TcpPort *uint64 `json:"TcpPort,omitnil,omitempty" name:"TcpPort"`

	// Host 设置，默认为业务域名
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// URL 路径，默认为“/”
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 返回错误码阈值, 可选值为 400 和 500, 默认值500
	ReturnCodeThreshold *uint64 `json:"ReturnCodeThreshold,omitnil,omitempty" name:"ReturnCodeThreshold"`

	// 跟随 3XX 重定向 ，不开启为 DISABLED， 开启为 ENABLED，默认不开启
	EnableRedirect *string `json:"EnableRedirect,omitnil,omitempty" name:"EnableRedirect"`

	// 启用 SNI，不开启为 DISABLED， 开启为 ENABLED，默认不开启
	EnableSni *string `json:"EnableSni,omitnil,omitempty" name:"EnableSni"`

	// 丢包率告警阈值，当CheckProtocol=ping时必填，取值在10 30 50 80 90 100
	PacketLossRate *uint64 `json:"PacketLossRate,omitnil,omitempty" name:"PacketLossRate"`

	// 持续周期数，可选值1-5
	ContinuePeriod *uint64 `json:"ContinuePeriod,omitnil,omitempty" name:"ContinuePeriod"`
}

type ModifyMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 监控器id
	MonitorId *uint64 `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`

	// 监控器名称
	MonitorName *string `json:"MonitorName,omitnil,omitempty" name:"MonitorName"`

	// 检查协议，可选值 PING TCP HTTP HTTPS
	CheckProtocol *string `json:"CheckProtocol,omitnil,omitempty" name:"CheckProtocol"`

	// 检查间隔（秒），可选值 15 60 120 300
	CheckInterval *uint64 `json:"CheckInterval,omitnil,omitempty" name:"CheckInterval"`

	// 超时时间，单位:秒，可选值 2  3  5  10
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 重试次数，可选值为 0，1，2
	FailTimes *uint64 `json:"FailTimes,omitnil,omitempty" name:"FailTimes"`

	// 失败比例，可选值为 20 30 40 50 60 70 80 100，默认值为50
	FailRate *uint64 `json:"FailRate,omitnil,omitempty" name:"FailRate"`

	// 监控节点类型，可选值有AUTO INTERNAL OVERSEAS IPV6 ALL
	DetectorStyle *string `json:"DetectorStyle,omitnil,omitempty" name:"DetectorStyle"`

	// 探测器组id列表
	DetectorGroupIds []*uint64 `json:"DetectorGroupIds,omitnil,omitempty" name:"DetectorGroupIds"`

	// PING 包数目， 当CheckProtocol=ping时必填，可选值 20 50 100
	PingNum *uint64 `json:"PingNum,omitnil,omitempty" name:"PingNum"`

	// 检查端口，可选值为1-65535之间的整数
	TcpPort *uint64 `json:"TcpPort,omitnil,omitempty" name:"TcpPort"`

	// Host 设置，默认为业务域名
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// URL 路径，默认为“/”
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 返回错误码阈值, 可选值为 400 和 500, 默认值500
	ReturnCodeThreshold *uint64 `json:"ReturnCodeThreshold,omitnil,omitempty" name:"ReturnCodeThreshold"`

	// 跟随 3XX 重定向 ，不开启为 DISABLED， 开启为 ENABLED，默认不开启
	EnableRedirect *string `json:"EnableRedirect,omitnil,omitempty" name:"EnableRedirect"`

	// 启用 SNI，不开启为 DISABLED， 开启为 ENABLED，默认不开启
	EnableSni *string `json:"EnableSni,omitnil,omitempty" name:"EnableSni"`

	// 丢包率告警阈值，当CheckProtocol=ping时必填，取值在10 30 50 80 90 100
	PacketLossRate *uint64 `json:"PacketLossRate,omitnil,omitempty" name:"PacketLossRate"`

	// 持续周期数，可选值1-5
	ContinuePeriod *uint64 `json:"ContinuePeriod,omitnil,omitempty" name:"ContinuePeriod"`
}

func (r *ModifyMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MonitorId")
	delete(f, "MonitorName")
	delete(f, "CheckProtocol")
	delete(f, "CheckInterval")
	delete(f, "Timeout")
	delete(f, "FailTimes")
	delete(f, "FailRate")
	delete(f, "DetectorStyle")
	delete(f, "DetectorGroupIds")
	delete(f, "PingNum")
	delete(f, "TcpPort")
	delete(f, "Host")
	delete(f, "Path")
	delete(f, "ReturnCodeThreshold")
	delete(f, "EnableRedirect")
	delete(f, "EnableSni")
	delete(f, "PacketLossRate")
	delete(f, "ContinuePeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMonitorResponseParams struct {
	// success 为修改成功
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMonitorResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMonitorResponseParams `json:"Response"`
}

func (r *ModifyMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPackageAutoRenewRequestParams struct {
	// 资源ID，续费和变配的时候需要传
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 自动续费：1 开启自动续费；2 关闭自动续费
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`
}

type ModifyPackageAutoRenewRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID，续费和变配的时候需要传
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 自动续费：1 开启自动续费；2 关闭自动续费
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`
}

func (r *ModifyPackageAutoRenewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPackageAutoRenewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "AutoRenew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPackageAutoRenewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPackageAutoRenewResponseParams struct {
	// 资源id列表，目前只会返回一个资源，取第一个值即可
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPackageAutoRenewResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPackageAutoRenewResponseParams `json:"Response"`
}

func (r *ModifyPackageAutoRenewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPackageAutoRenewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStrategyRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略id
	StrategyId *uint64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 解析线路，需要全量传参
	Source []*Source `json:"Source,omitnil,omitempty" name:"Source"`

	// 主力地址池集合，需要全量传参
	MainAddressPoolSet []*MainAddressPool `json:"MainAddressPoolSet,omitnil,omitempty" name:"MainAddressPoolSet"`

	// 兜底地址池集合，需要全量传参
	FallbackAddressPoolSet []*MainAddressPool `json:"FallbackAddressPoolSet,omitnil,omitempty" name:"FallbackAddressPoolSet"`

	// 策略名称，不允许重复
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 策略开启状态：ENABLED开启；DISABLED关闭
	IsEnabled *string `json:"IsEnabled,omitnil,omitempty" name:"IsEnabled"`

	// 是否开启策略强制保留默认线路 disabled, enabled，默认不开启且只有一个策略能开启
	KeepDomainRecords *string `json:"KeepDomainRecords,omitnil,omitempty" name:"KeepDomainRecords"`

	// 调度模式：AUTO默认；STOP仅暂停不切换
	SwitchPoolType *string `json:"SwitchPoolType,omitnil,omitempty" name:"SwitchPoolType"`
}

type ModifyStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略id
	StrategyId *uint64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 解析线路，需要全量传参
	Source []*Source `json:"Source,omitnil,omitempty" name:"Source"`

	// 主力地址池集合，需要全量传参
	MainAddressPoolSet []*MainAddressPool `json:"MainAddressPoolSet,omitnil,omitempty" name:"MainAddressPoolSet"`

	// 兜底地址池集合，需要全量传参
	FallbackAddressPoolSet []*MainAddressPool `json:"FallbackAddressPoolSet,omitnil,omitempty" name:"FallbackAddressPoolSet"`

	// 策略名称，不允许重复
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 策略开启状态：ENABLED开启；DISABLED关闭
	IsEnabled *string `json:"IsEnabled,omitnil,omitempty" name:"IsEnabled"`

	// 是否开启策略强制保留默认线路 disabled, enabled，默认不开启且只有一个策略能开启
	KeepDomainRecords *string `json:"KeepDomainRecords,omitnil,omitempty" name:"KeepDomainRecords"`

	// 调度模式：AUTO默认；STOP仅暂停不切换
	SwitchPoolType *string `json:"SwitchPoolType,omitnil,omitempty" name:"SwitchPoolType"`
}

func (r *ModifyStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StrategyId")
	delete(f, "Source")
	delete(f, "MainAddressPoolSet")
	delete(f, "FallbackAddressPoolSet")
	delete(f, "StrategyName")
	delete(f, "IsEnabled")
	delete(f, "KeepDomainRecords")
	delete(f, "SwitchPoolType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStrategyResponseParams struct {
	// 是否成功
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStrategyResponseParams `json:"Response"`
}

func (r *ModifyStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorDetail struct {
	// 探测规则id
	MonitorId *uint64 `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`

	// 监控器名称
	MonitorName *string `json:"MonitorName,omitnil,omitempty" name:"MonitorName"`

	// 所属用户
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 监控节点id组
	DetectorGroupIds []*uint64 `json:"DetectorGroupIds,omitnil,omitempty" name:"DetectorGroupIds"`

	// 探测协议 PING TCP HTTP HTTPS
	CheckProtocol *string `json:"CheckProtocol,omitnil,omitempty" name:"CheckProtocol"`

	// 探测周期
	CheckInterval *uint64 `json:"CheckInterval,omitnil,omitempty" name:"CheckInterval"`

	// 发包数量
	PingNum *uint64 `json:"PingNum,omitnil,omitempty" name:"PingNum"`

	// tcp端口
	TcpPort *uint64 `json:"TcpPort,omitnil,omitempty" name:"TcpPort"`

	// 探测 host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 探测路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 返回值阈值
	ReturnCodeThreshold *uint64 `json:"ReturnCodeThreshold,omitnil,omitempty" name:"ReturnCodeThreshold"`

	// 是否开启3xx重定向跟随 ENABLED DISABLED
	EnableRedirect *string `json:"EnableRedirect,omitnil,omitempty" name:"EnableRedirect"`

	// 是否启用 sni
	// ENABLED DISABLED
	EnableSni *string `json:"EnableSni,omitnil,omitempty" name:"EnableSni"`

	// 丢包率上限
	PacketLossRate *uint64 `json:"PacketLossRate,omitnil,omitempty" name:"PacketLossRate"`

	// 探测超时
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 失败次数
	FailTimes *uint64 `json:"FailTimes,omitnil,omitempty" name:"FailTimes"`

	// 失败率上限100
	FailRate *uint64 `json:"FailRate,omitnil,omitempty" name:"FailRate"`

	// 创建时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 更新时间
	UpdatedOn *string `json:"UpdatedOn,omitnil,omitempty" name:"UpdatedOn"`

	// 监控节点类型
	// AUTO INTERNAL OVERSEAS IPV6 ALL
	DetectorStyle *string `json:"DetectorStyle,omitnil,omitempty" name:"DetectorStyle"`

	// 探测次数
	DetectNum *uint64 `json:"DetectNum,omitnil,omitempty" name:"DetectNum"`

	// 持续周期数
	ContinuePeriod *uint64 `json:"ContinuePeriod,omitnil,omitempty" name:"ContinuePeriod"`
}

type Quota struct {
	// 探测任务配额
	TaskQuota *uint64 `json:"TaskQuota,omitnil,omitempty" name:"TaskQuota"`

	// 地址池配额
	PoolQuota *uint64 `json:"PoolQuota,omitnil,omitempty" name:"PoolQuota"`

	// 地址配额
	AddressQuota *uint64 `json:"AddressQuota,omitnil,omitempty" name:"AddressQuota"`

	// 探点资源数
	MonitorQuota *uint64 `json:"MonitorQuota,omitnil,omitempty" name:"MonitorQuota"`

	// 消息资源数
	MessageQuota *uint64 `json:"MessageQuota,omitnil,omitempty" name:"MessageQuota"`

	// 已使用探测任务数
	UsedTaskQuota *uint64 `json:"UsedTaskQuota,omitnil,omitempty" name:"UsedTaskQuota"`

	// 已使用体验实例数
	UsedFreeInstanceNum *uint64 `json:"UsedFreeInstanceNum,omitnil,omitempty" name:"UsedFreeInstanceNum"`

	// 已使用付费实例
	UsedBillInstanceNum *uint64 `json:"UsedBillInstanceNum,omitnil,omitempty" name:"UsedBillInstanceNum"`

	// 体验套餐总数
	FreePackageNum *uint64 `json:"FreePackageNum,omitnil,omitempty" name:"FreePackageNum"`

	// 已使用付费套餐数
	UsedBillPackageNum *uint64 `json:"UsedBillPackageNum,omitnil,omitempty" name:"UsedBillPackageNum"`

	// 付费套餐总数
	BillPackageNum *uint64 `json:"BillPackageNum,omitnil,omitempty" name:"BillPackageNum"`
}

type ResourceFilter struct {
	// 过滤字段名，支持的列表如下：
	// - type：主资源类型，CDN。
	// - instanceId：IGTM实例ID。此为必传参数，未传将导致接口查询失败。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤字段值。
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`

	// 是否启用模糊查询，仅支持过滤字段名为domain。
	// 模糊查询时，Value长度最大为1，否则Value长度最大为5。(预留字段，暂未使用)
	Fuzzy *bool `json:"Fuzzy,omitnil,omitempty" name:"Fuzzy"`
}

type Source struct {
	// 解析请求来源线路id
	DnsLineId *uint64 `json:"DnsLineId,omitnil,omitempty" name:"DnsLineId"`

	// 解析请求来源线路名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type Strategy struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 地址来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source []*Source `json:"Source,omitnil,omitempty" name:"Source"`

	// 策略id
	StrategyId *uint64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 健康状态：ok健康、warn风险、down故障
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 生效的主力池id，null则为未知
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivateMainPoolId *int64 `json:"ActivateMainPoolId,omitnil,omitempty" name:"ActivateMainPoolId"`

	// 当前生效地址池所在级数，为0则代表兜底生效，null则为未知
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivateLevel *int64 `json:"ActivateLevel,omitnil,omitempty" name:"ActivateLevel"`

	// 当前生效地址池集合类型：main主力；fallback兜底
	ActivePoolType *string `json:"ActivePoolType,omitnil,omitempty" name:"ActivePoolType"`

	// 当前生效地址池流量策略：all解析所有；weight负载均衡
	ActiveTrafficStrategy *string `json:"ActiveTrafficStrategy,omitnil,omitempty" name:"ActiveTrafficStrategy"`

	// 监控器数量
	MonitorNum *uint64 `json:"MonitorNum,omitnil,omitempty" name:"MonitorNum"`

	// 是否开启：ENABLED开启；DISABLED关闭
	IsEnabled *string `json:"IsEnabled,omitnil,omitempty" name:"IsEnabled"`

	// 是否保留线路：enabled保留，disabled不保留，只保留默认线路
	KeepDomainRecords *string `json:"KeepDomainRecords,omitnil,omitempty" name:"KeepDomainRecords"`

	// 调度模式：AUTO默认；PAUSE仅暂停不切换
	SwitchPoolType *string `json:"SwitchPoolType,omitnil,omitempty" name:"SwitchPoolType"`

	// 创建时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 更新时间
	UpdatedOn *string `json:"UpdatedOn,omitnil,omitempty" name:"UpdatedOn"`
}

type StrategyDetail struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略id
	StrategyId *uint64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 策略名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 线路
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source []*Source `json:"Source,omitnil,omitempty" name:"Source"`

	// 主力地址池集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainAddressPoolSet []*MainAddressPool `json:"MainAddressPoolSet,omitnil,omitempty" name:"MainAddressPoolSet"`

	// 兜底地址池id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FallbackAddressPoolSet []*MainAddressPool `json:"FallbackAddressPoolSet,omitnil,omitempty" name:"FallbackAddressPoolSet"`

	// 是否保留线路：enabled保留，disabled不保留，只保留默认线路
	KeepDomainRecords *string `json:"KeepDomainRecords,omitnil,omitempty" name:"KeepDomainRecords"`

	// 生效主力地址池id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivateMainPoolId *uint64 `json:"ActivateMainPoolId,omitnil,omitempty" name:"ActivateMainPoolId"`

	// 创建时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 更新时间
	UpdatedOn *string `json:"UpdatedOn,omitnil,omitempty" name:"UpdatedOn"`

	// 调度模式：AUTO默认；PAUSE仅暂停不切换
	SwitchPoolType *string `json:"SwitchPoolType,omitnil,omitempty" name:"SwitchPoolType"`
}