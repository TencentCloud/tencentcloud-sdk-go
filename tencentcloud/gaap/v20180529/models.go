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

package v20180529

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccessConfiguration struct {

	// 加速地域。
	AccessRegion *string `json:"AccessRegion,omitempty" name:"AccessRegion"`

	// 通道带宽上限，单位：Mbps。
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 通道并发量上限，表示同时在线的连接数，单位：万。
	Concurrent *uint64 `json:"Concurrent,omitempty" name:"Concurrent"`

	// 网络类型，可取值：normal、cn2，默认值为normal
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
}

type AccessRegionDetial struct {

	// 区域ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 区域的中文或英文名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 可选的并发量取值数组
	ConcurrentList []*int64 `json:"ConcurrentList,omitempty" name:"ConcurrentList"`

	// 可选的带宽取值数组
	BandwidthList []*int64 `json:"BandwidthList,omitempty" name:"BandwidthList"`

	// 机房所属大区
	RegionArea *string `json:"RegionArea,omitempty" name:"RegionArea"`

	// 机房所属大区名
	RegionAreaName *string `json:"RegionAreaName,omitempty" name:"RegionAreaName"`

	// 机房类型, dc表示DataCenter数据中心, ec表示EdgeComputing边缘节点
	IDCType *string `json:"IDCType,omitempty" name:"IDCType"`

	// 特性位图，每个bit位代表一种特性，其中：
	// 0，表示不支持该特性；
	// 1，表示支持该特性。
	// 特性位图含义如下（从右往左）：
	// 第1个bit，支持4层加速；
	// 第2个bit，支持7层加速；
	// 第3个bit，支持Http3接入；
	// 第4个bit，支持IPv6；
	// 第5个bit，支持精品BGP接入；
	// 第6个bit，支持三网接入；
	// 第7个bit，支持接入段Qos加速。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeatureBitmap *int64 `json:"FeatureBitmap,omitempty" name:"FeatureBitmap"`
}

type AccessRegionDomainConf struct {

	// 地域ID。
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 就近接入区域国家内部编码，编码列表可通过DescribeCountryAreaMapping接口获取。
	NationCountryInnerList []*string `json:"NationCountryInnerList,omitempty" name:"NationCountryInnerList"`
}

type AddRealServersRequest struct {
	*tchttp.BaseRequest

	// 源站对应的项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 源站对应的IP或域名
	RealServerIP []*string `json:"RealServerIP,omitempty" name:"RealServerIP"`

	// 源站名称
	RealServerName *string `json:"RealServerName,omitempty" name:"RealServerName"`

	// 标签列表
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet"`
}

func (r *AddRealServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRealServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "RealServerIP")
	delete(f, "RealServerName")
	delete(f, "TagSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddRealServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddRealServersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 源站信息列表
		RealServerSet []*NewRealServer `json:"RealServerSet,omitempty" name:"RealServerSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddRealServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRealServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BandwidthPriceGradient struct {

	// 带宽范围。
	BandwidthRange []*int64 `json:"BandwidthRange,omitempty" name:"BandwidthRange"`

	// 在对应带宽范围内的单宽单价，单位：元/Mbps/天。
	BandwidthUnitPrice *float64 `json:"BandwidthUnitPrice,omitempty" name:"BandwidthUnitPrice"`

	// 带宽折扣价，单位：元/Mbps/天。
	DiscountBandwidthUnitPrice *float64 `json:"DiscountBandwidthUnitPrice,omitempty" name:"DiscountBandwidthUnitPrice"`
}

type BindListenerRealServersRequest struct {
	*tchttp.BaseRequest

	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 待绑定源站列表。如果该监听器的源站调度策略是加权轮询，需要填写源站权重 RealServerWeight, 不填或者其他调度类型默认源站权重为1。
	RealServerBindSet []*RealServerBindSetReq `json:"RealServerBindSet,omitempty" name:"RealServerBindSet"`
}

func (r *BindListenerRealServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindListenerRealServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "RealServerBindSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindListenerRealServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BindListenerRealServersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindListenerRealServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindListenerRealServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindRealServer struct {

	// 源站ID
	RealServerId *string `json:"RealServerId,omitempty" name:"RealServerId"`

	// 源站IP或者域名
	RealServerIP *string `json:"RealServerIP,omitempty" name:"RealServerIP"`

	// 该源站所占权重
	RealServerWeight *int64 `json:"RealServerWeight,omitempty" name:"RealServerWeight"`

	// 源站健康检查状态，其中：
	// 0表示正常；
	// 1表示异常。
	// 未开启健康检查状态时，该状态始终为正常。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealServerStatus *int64 `json:"RealServerStatus,omitempty" name:"RealServerStatus"`

	// 源站的端口号
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealServerPort *int64 `json:"RealServerPort,omitempty" name:"RealServerPort"`

	// 当源站为域名时，域名被解析成一个或者多个IP，该字段表示其中异常的IP列表。状态异常，但该字段为空时，表示域名解析异常。
	DownIPList []*string `json:"DownIPList,omitempty" name:"DownIPList"`
}

type BindRealServerInfo struct {

	// 源站的IP或域名
	RealServerIP *string `json:"RealServerIP,omitempty" name:"RealServerIP"`

	// 源站ID
	RealServerId *string `json:"RealServerId,omitempty" name:"RealServerId"`

	// 源站名称
	RealServerName *string `json:"RealServerName,omitempty" name:"RealServerName"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet"`
}

type BindRuleRealServersRequest struct {
	*tchttp.BaseRequest

	// 转发规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 需要绑定的源站信息列表。
	// 如果已经存在绑定的源站，则会覆盖更新成这个源站列表。
	// 当不带该字段时，表示解绑该规则上的所有源站。
	// 如果该规则的源站调度策略是加权轮询，需要填写源站权重 RealServerWeight, 不填或者其他调度类型默认源站权重为1。
	RealServerBindSet []*RealServerBindSetReq `json:"RealServerBindSet,omitempty" name:"RealServerBindSet"`
}

func (r *BindRuleRealServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindRuleRealServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "RealServerBindSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindRuleRealServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BindRuleRealServersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindRuleRealServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindRuleRealServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Certificate struct {

	// 证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 证书名称（旧参数，请使用CertificateAlias）。
	CertificateName *string `json:"CertificateName,omitempty" name:"CertificateName"`

	// 证书类型。
	CertificateType *int64 `json:"CertificateType,omitempty" name:"CertificateType"`

	// 证书名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateAlias *string `json:"CertificateAlias,omitempty" name:"CertificateAlias"`

	// 证书创建时间，采用Unix时间戳的方式，表示从1970年1月1日（UTC/GMT的午夜）开始所经过的秒数。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 证书生效起始时间，采用Unix时间戳的方式，表示从1970年1月1日（UTC/GMT的午夜）开始所经过的秒数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *uint64 `json:"BeginTime,omitempty" name:"BeginTime"`

	// 证书过期时间，采用Unix时间戳的方式，表示从1970年1月1日（UTC/GMT的午夜）开始所经过的秒数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 证书签发者通用名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IssuerCN *string `json:"IssuerCN,omitempty" name:"IssuerCN"`

	// 证书主题通用名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubjectCN *string `json:"SubjectCN,omitempty" name:"SubjectCN"`
}

type CertificateAliasInfo struct {

	// 证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 证书别名
	CertificateAlias *string `json:"CertificateAlias,omitempty" name:"CertificateAlias"`
}

type CertificateDetail struct {

	// 证书ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 证书类型。
	CertificateType *int64 `json:"CertificateType,omitempty" name:"CertificateType"`

	// 证书名字。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateAlias *string `json:"CertificateAlias,omitempty" name:"CertificateAlias"`

	// 证书内容。
	CertificateContent *string `json:"CertificateContent,omitempty" name:"CertificateContent"`

	// 密钥内容。仅当证书类型为SSL证书时，返回该字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateKey *string `json:"CertificateKey,omitempty" name:"CertificateKey"`

	// 创建时间，采用Unix时间戳的方式，表示从1970年1月1日（UTC/GMT的午夜）开始所经过的秒数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 证书生效起始时间，采用Unix时间戳的方式，表示从1970年1月1日（UTC/GMT的午夜）开始所经过的秒数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *uint64 `json:"BeginTime,omitempty" name:"BeginTime"`

	// 证书过期时间，采用Unix时间戳的方式，表示从1970年1月1日（UTC/GMT的午夜）开始所经过的秒数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 证书签发者通用名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IssuerCN *string `json:"IssuerCN,omitempty" name:"IssuerCN"`

	// 证书主题通用名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubjectCN *string `json:"SubjectCN,omitempty" name:"SubjectCN"`
}

type CheckProxyCreateRequest struct {
	*tchttp.BaseRequest

	// 通道的接入(加速)区域。取值可通过接口DescribeAccessRegionsByDestRegion获取到
	AccessRegion *string `json:"AccessRegion,omitempty" name:"AccessRegion"`

	// 通道的源站区域。取值可通过接口DescribeDestRegions获取到
	RealServerRegion *string `json:"RealServerRegion,omitempty" name:"RealServerRegion"`

	// 通道带宽上限，单位：Mbps。
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 通道并发量上限，表示同时在线的连接数，单位：万。
	Concurrent *uint64 `json:"Concurrent,omitempty" name:"Concurrent"`

	// 如果在通道组下创建通道，需要填写通道组的ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// IP版本，可取值：IPv4、IPv6，默认值IPv4
	IPAddressVersion *string `json:"IPAddressVersion,omitempty" name:"IPAddressVersion"`

	// 网络类型，可取值：normal、cn2，默认值normal
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// 通道套餐类型。Thunder表示标准通道组，Accelerator表示游戏加速器通道，CrossBorder表示跨境通道。
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// 支持Http3的开关，其中：0，表示不需要支持Http3接入；1，表示需要支持Http3接入。注意：如果开启了Http3的功能，那么该通道就不再支持TCP/UDP接入的功能。该功能的启停无法在通道创建完毕后再修改。
	Http3Supported *int64 `json:"Http3Supported,omitempty" name:"Http3Supported"`
}

func (r *CheckProxyCreateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckProxyCreateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessRegion")
	delete(f, "RealServerRegion")
	delete(f, "Bandwidth")
	delete(f, "Concurrent")
	delete(f, "GroupId")
	delete(f, "IPAddressVersion")
	delete(f, "NetworkType")
	delete(f, "PackageType")
	delete(f, "Http3Supported")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckProxyCreateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CheckProxyCreateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询能否创建给定配置的通道，1可以创建，0不可创建。
		CheckFlag *uint64 `json:"CheckFlag,omitempty" name:"CheckFlag"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckProxyCreateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckProxyCreateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseProxiesRequest struct {
	*tchttp.BaseRequest

	// （旧参数，请切换到ProxyIds）通道的实例ID。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	// 更多详细信息请参阅：如何保证幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// （新参数）通道的实例ID。
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`
}

func (r *CloseProxiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProxiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ClientToken")
	delete(f, "ProxyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseProxiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloseProxiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 非运行状态下的通道实例ID列表，不可开启。
		InvalidStatusInstanceSet []*string `json:"InvalidStatusInstanceSet,omitempty" name:"InvalidStatusInstanceSet"`

		// 开启操作失败的通道实例ID列表。
		OperationFailedInstanceSet []*string `json:"OperationFailedInstanceSet,omitempty" name:"OperationFailedInstanceSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseProxiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProxiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseProxyGroupRequest struct {
	*tchttp.BaseRequest

	// 通道组的实例 ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CloseProxyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProxyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseProxyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloseProxyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 非运行状态下的通道实例ID列表，不可开启。
		InvalidStatusInstanceSet []*string `json:"InvalidStatusInstanceSet,omitempty" name:"InvalidStatusInstanceSet"`

		// 开启操作失败的通道实例ID列表。
		OperationFailedInstanceSet []*string `json:"OperationFailedInstanceSet,omitempty" name:"OperationFailedInstanceSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseProxyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProxyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseSecurityPolicyRequest struct {
	*tchttp.BaseRequest

	// 通道ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 安全组策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *CloseSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyId")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloseSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步流程ID，可以通过DescribeAsyncTaskStatus 查询流程执行进展和状态
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CountryAreaMap struct {

	// 国家名称。
	NationCountryName *string `json:"NationCountryName,omitempty" name:"NationCountryName"`

	// 国家编码。
	NationCountryInnerCode *string `json:"NationCountryInnerCode,omitempty" name:"NationCountryInnerCode"`

	// 地区名称。
	GeographicalZoneName *string `json:"GeographicalZoneName,omitempty" name:"GeographicalZoneName"`

	// 地区编码。
	GeographicalZoneInnerCode *string `json:"GeographicalZoneInnerCode,omitempty" name:"GeographicalZoneInnerCode"`

	// 大洲名称。
	ContinentName *string `json:"ContinentName,omitempty" name:"ContinentName"`

	// 大洲编码。
	ContinentInnerCode *string `json:"ContinentInnerCode,omitempty" name:"ContinentInnerCode"`
}

type CreateCertificateRequest struct {
	*tchttp.BaseRequest

	// 证书类型。其中：
	// 0，表示基础认证配置；
	// 1，表示客户端CA证书；
	// 2，服务器SSL证书；
	// 3，表示源站CA证书；
	// 4，表示通道SSL证书。
	CertificateType *int64 `json:"CertificateType,omitempty" name:"CertificateType"`

	// 证书内容。采用url编码。其中：
	// 当证书类型为基础认证配置时，该参数填写用户名/密码对。格式：“用户名：密码”，例如：root:FSGdT。其中密码使用htpasswd或者openssl，例如：openssl passwd -crypt 123456。
	// 当证书类型为CA/SSL证书时，该参数填写证书内容，格式为pem。
	CertificateContent *string `json:"CertificateContent,omitempty" name:"CertificateContent"`

	// 证书名称
	CertificateAlias *string `json:"CertificateAlias,omitempty" name:"CertificateAlias"`

	// 密钥内容。采用url编码。仅当证书类型为SSL证书时，需要填写该参数。格式为pem。
	CertificateKey *string `json:"CertificateKey,omitempty" name:"CertificateKey"`
}

func (r *CreateCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateType")
	delete(f, "CertificateContent")
	delete(f, "CertificateAlias")
	delete(f, "CertificateKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书ID
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCustomHeaderRequest struct {
	*tchttp.BaseRequest

	// 规则id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 新增的header名称和内容列表， ‘’$remote_addr‘’会被解析替换成客户端ip，其他值原样透传到源站。
	Headers []*HttpHeaderParam `json:"Headers,omitempty" name:"Headers"`
}

func (r *CreateCustomHeaderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomHeaderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "Headers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomHeaderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCustomHeaderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCustomHeaderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomHeaderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainErrorPageInfoRequest struct {
	*tchttp.BaseRequest

	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 原始错误码
	ErrorNos []*int64 `json:"ErrorNos,omitempty" name:"ErrorNos"`

	// 新的响应包体
	Body *string `json:"Body,omitempty" name:"Body"`

	// 新错误码
	NewErrorNo *int64 `json:"NewErrorNo,omitempty" name:"NewErrorNo"`

	// 需要删除的响应头
	ClearHeaders []*string `json:"ClearHeaders,omitempty" name:"ClearHeaders"`

	// 需要设置的响应头
	SetHeaders []*HttpHeaderParam `json:"SetHeaders,omitempty" name:"SetHeaders"`
}

func (r *CreateDomainErrorPageInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainErrorPageInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "Domain")
	delete(f, "ErrorNos")
	delete(f, "Body")
	delete(f, "NewErrorNo")
	delete(f, "ClearHeaders")
	delete(f, "SetHeaders")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainErrorPageInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainErrorPageInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 错误定制响应的配置ID
		ErrorPageId *string `json:"ErrorPageId,omitempty" name:"ErrorPageId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDomainErrorPageInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainErrorPageInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainRequest struct {
	*tchttp.BaseRequest

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 需要创建的域名，一个监听器下最大支持100个域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 服务器证书，用于客户端与GAAP的HTTPS的交互。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 客户端CA证书，用于客户端与GAAP的HTTPS的交互。
	// 仅当采用双向认证的方式时，需要设置该字段或PolyClientCertificateIds字段。
	ClientCertificateId *string `json:"ClientCertificateId,omitempty" name:"ClientCertificateId"`

	// 客户端CA证书，用于客户端与GAAP的HTTPS的交互。
	// 仅当采用双向认证的方式时，需要设置该字段或ClientCertificateId字段。
	PolyClientCertificateIds []*string `json:"PolyClientCertificateIds,omitempty" name:"PolyClientCertificateIds"`

	// 是否开启Http3特性的标识，其中：
	// 0，表示不开启Http3；
	// 1，表示开启Http3。
	// 默认不开启Http3。可以通过SetDomainHttp3开启。
	Http3Supported *int64 `json:"Http3Supported,omitempty" name:"Http3Supported"`
}

func (r *CreateDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "Domain")
	delete(f, "CertificateId")
	delete(f, "ClientCertificateId")
	delete(f, "PolyClientCertificateIds")
	delete(f, "Http3Supported")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHTTPListenerRequest struct {
	*tchttp.BaseRequest

	// 监听器名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器端口，基于同种传输层协议（TCP 或 UDP）的监听器，端口不可重复
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 通道ID，与GroupId不能同时设置，对应为通道创建监听器
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 通道组ID，与ProxyId不能同时设置，对应为通道组创建监听器
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CreateHTTPListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHTTPListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerName")
	delete(f, "Port")
	delete(f, "ProxyId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHTTPListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateHTTPListenerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建的监听器ID
		ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHTTPListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHTTPListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHTTPSListenerRequest struct {
	*tchttp.BaseRequest

	// 监听器名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器端口，基于同种传输层协议（TCP 或 UDP）的监听器，端口不可重复
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 服务器证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 加速通道转发到源站的协议类型：HTTP | HTTPS
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`

	// 通道ID，与GroupId之间只能设置一个。表示创建通道的监听器。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 认证类型，其中：
	// 0，单向认证；
	// 1，双向认证。
	// 默认使用单向认证。
	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`

	// 客户端CA单证书ID，仅当双向认证时设置该参数或PolyClientCertificateIds参数
	ClientCertificateId *string `json:"ClientCertificateId,omitempty" name:"ClientCertificateId"`

	// 新的客户端多CA证书ID，仅当双向认证时设置该参数或设置ClientCertificateId参数
	PolyClientCertificateIds []*string `json:"PolyClientCertificateIds,omitempty" name:"PolyClientCertificateIds"`

	// 通道组ID，与ProxyId之间只能设置一个。表示创建通道组的监听器。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 支持Http3的开关，其中：
	// 0，表示不需要支持Http3接入；
	// 1，表示需要支持Http3接入。
	// 注意：如果支持了Http3的功能，那么该监听器会占用对应的UDP接入端口，不可再创建相同端口的UDP监听器。
	// 该功能的启停无法在监听器创建完毕后再修改。
	Http3Supported *int64 `json:"Http3Supported,omitempty" name:"Http3Supported"`
}

func (r *CreateHTTPSListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHTTPSListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerName")
	delete(f, "Port")
	delete(f, "CertificateId")
	delete(f, "ForwardProtocol")
	delete(f, "ProxyId")
	delete(f, "AuthType")
	delete(f, "ClientCertificateId")
	delete(f, "PolyClientCertificateIds")
	delete(f, "GroupId")
	delete(f, "Http3Supported")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHTTPSListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateHTTPSListenerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建的监听器ID
		ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHTTPSListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHTTPSListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProxyGroupDomainRequest struct {
	*tchttp.BaseRequest

	// 需要开启域名的通道组ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CreateProxyGroupDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyGroupDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProxyGroupDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateProxyGroupDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道组ID。
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProxyGroupDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyGroupDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProxyGroupRequest struct {
	*tchttp.BaseRequest

	// 通道组所属项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 通道组别名
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 源站地域，参考接口DescribeDestRegions 返回参数RegionDetail中的RegionId
	RealServerRegion *string `json:"RealServerRegion,omitempty" name:"RealServerRegion"`

	// 标签列表
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet"`

	// 加速地域列表，包括加速地域名，及该地域对应的带宽和并发配置。
	AccessRegionSet []*AccessConfiguration `json:"AccessRegionSet,omitempty" name:"AccessRegionSet"`

	// IP版本，可取值：IPv4、IPv6，默认值IPv4
	IPAddressVersion *string `json:"IPAddressVersion,omitempty" name:"IPAddressVersion"`

	// 通道组套餐类型，可取值：Thunder、Accelerator，默认值Thunder
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// 支持Http3的开关，其中：
	// 0，表示不需要支持Http3接入；
	// 1，表示需要支持Http3接入。
	// 注意：如果开启了Http3的功能，那么该通道组就不再支持TCP/UDP接入的功能。
	// 该功能的启停无法在通道组创建完毕后再修改。
	Http3Supported *int64 `json:"Http3Supported,omitempty" name:"Http3Supported"`
}

func (r *CreateProxyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "GroupName")
	delete(f, "RealServerRegion")
	delete(f, "TagSet")
	delete(f, "AccessRegionSet")
	delete(f, "IPAddressVersion")
	delete(f, "PackageType")
	delete(f, "Http3Supported")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProxyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateProxyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道组ID
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProxyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProxyRequest struct {
	*tchttp.BaseRequest

	// 通道的项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 通道名称。
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// 接入地域。
	AccessRegion *string `json:"AccessRegion,omitempty" name:"AccessRegion"`

	// 通道带宽上限，单位：Mbps。
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 通道并发量上限，表示同时在线的连接数，单位：万。
	Concurrent *uint64 `json:"Concurrent,omitempty" name:"Concurrent"`

	// 源站地域。当GroupId存在时，源站地域为通道组的源站地域,此时可不填该字段。当GroupId不存在时，需要填写该字段
	RealServerRegion *string `json:"RealServerRegion,omitempty" name:"RealServerRegion"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	// 更多详细信息请参阅：如何保证幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 通道所在的通道组ID，当在通道组中创建通道时必带，否则忽略该字段。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 通道需要添加的标签列表。
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet"`

	// 被复制的通道ID。只有处于运行中状态的通道可以被复制。
	// 当设置该参数时，表示复制该通道。
	ClonedProxyId *string `json:"ClonedProxyId,omitempty" name:"ClonedProxyId"`

	// 计费方式 (0:按带宽计费，1:按流量计费 默认按带宽计费）
	BillingType *int64 `json:"BillingType,omitempty" name:"BillingType"`

	// IP版本，可取值：IPv4、IPv6，默认值IPv4
	IPAddressVersion *string `json:"IPAddressVersion,omitempty" name:"IPAddressVersion"`

	// 网络类型，normal表示常规BGP，cn2表示精品BGP，triple表示三网
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// 通道套餐类型，Thunder表示标准通道组，Accelerator表示游戏加速器通道，CrossBorder表示跨境通道。
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// 支持Http3的开关，其中：0，表示不需要支持Http3接入；1，表示需要支持Http3接入。注意：如果开启了Http3的功能，那么该通道就不再支持TCP/UDP接入的功能。该功能的启停无法在通道创建完毕后再修改。
	Http3Supported *int64 `json:"Http3Supported,omitempty" name:"Http3Supported"`
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
	delete(f, "ProjectId")
	delete(f, "ProxyName")
	delete(f, "AccessRegion")
	delete(f, "Bandwidth")
	delete(f, "Concurrent")
	delete(f, "RealServerRegion")
	delete(f, "ClientToken")
	delete(f, "GroupId")
	delete(f, "TagSet")
	delete(f, "ClonedProxyId")
	delete(f, "BillingType")
	delete(f, "IPAddressVersion")
	delete(f, "NetworkType")
	delete(f, "PackageType")
	delete(f, "Http3Supported")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateProxyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道的实例ID。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateRuleRequest struct {
	*tchttp.BaseRequest

	// 7层监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发规则的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 转发规则的路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 转发规则对应源站的类型，支持IP和DOMAIN类型。
	RealServerType *string `json:"RealServerType,omitempty" name:"RealServerType"`

	// 规则转发源站调度策略，支持轮询（rr），加权轮询（wrr），最小连接数（lc）。
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 规则是否开启健康检查，1开启，0关闭。
	HealthCheck *uint64 `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 源站健康检查相关参数
	CheckParams *RuleCheckParams `json:"CheckParams,omitempty" name:"CheckParams"`

	// 加速通道转发到源站的协议类型：支持HTTP或HTTPS。
	// 不传递该字段时表示使用对应监听器的ForwardProtocol。
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`

	// 加速通道转发到远照的host，不设置该参数时，使用默认的host设置，即客户端发起的http请求的host。
	ForwardHost *string `json:"ForwardHost,omitempty" name:"ForwardHost"`
}

func (r *CreateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "Domain")
	delete(f, "Path")
	delete(f, "RealServerType")
	delete(f, "Scheduler")
	delete(f, "HealthCheck")
	delete(f, "CheckParams")
	delete(f, "ForwardProtocol")
	delete(f, "ForwardHost")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建转发规则成功返回规则ID
		RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityPolicyRequest struct {
	*tchttp.BaseRequest

	// 默认策略：ACCEPT或DROP
	DefaultAction *string `json:"DefaultAction,omitempty" name:"DefaultAction"`

	// 加速通道ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 通道组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CreateSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DefaultAction")
	delete(f, "ProxyId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全策略ID
		PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityRulesRequest struct {
	*tchttp.BaseRequest

	// 安全策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 访问规则列表
	RuleList []*SecurityPolicyRuleIn `json:"RuleList,omitempty" name:"RuleList"`
}

func (r *CreateSecurityRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "RuleList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则ID列表
		RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTCPListenersRequest struct {
	*tchttp.BaseRequest

	// 监听器名称。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器端口列表。
	Ports []*uint64 `json:"Ports,omitempty" name:"Ports"`

	// 监听器源站调度策略，支持轮询（rr），加权轮询（wrr），最小连接数（lc）。
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 源站是否开启健康检查：1开启，0关闭，UDP监听器不支持健康检查
	HealthCheck *uint64 `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 监听器对应源站类型，支持IP或者DOMAIN类型。DOMAIN源站类型不支持wrr的源站调度策略。
	RealServerType *string `json:"RealServerType,omitempty" name:"RealServerType"`

	// 通道ID，ProxyId和GroupId必须设置一个，但不能同时设置。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 通道组ID，ProxyId和GroupId必须设置一个，但不能同时设置。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 源站健康检查时间间隔，单位：秒。时间间隔取值在[5，300]之间。
	DelayLoop *uint64 `json:"DelayLoop,omitempty" name:"DelayLoop"`

	// 源站健康检查响应超时时间，单位：秒。超时时间取值在[2，60]之间。超时时间应小于健康检查时间间隔DelayLoop。
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitempty" name:"ConnectTimeout"`

	// 源站端口列表，该参数仅支持v1版本监听器和通道组监听器。
	RealServerPorts []*uint64 `json:"RealServerPorts,omitempty" name:"RealServerPorts"`

	// 监听器获取客户端 IP 的方式，0表示 TOA, 1表示Proxy Protocol
	ClientIPMethod *int64 `json:"ClientIPMethod,omitempty" name:"ClientIPMethod"`

	// 源站是否开启主备模式：1开启，0关闭，DOMAIN类型源站不支持开启
	FailoverSwitch *int64 `json:"FailoverSwitch,omitempty" name:"FailoverSwitch"`

	// 健康阈值，表示连续检查成功多少次后认定源站健康。范围为1到10
	HealthyThreshold *uint64 `json:"HealthyThreshold,omitempty" name:"HealthyThreshold"`

	// 不健康阈值，表示连续检查失败多少次数后认为源站不健康。范围为1到10
	UnhealthyThreshold *uint64 `json:"UnhealthyThreshold,omitempty" name:"UnhealthyThreshold"`
}

func (r *CreateTCPListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTCPListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerName")
	delete(f, "Ports")
	delete(f, "Scheduler")
	delete(f, "HealthCheck")
	delete(f, "RealServerType")
	delete(f, "ProxyId")
	delete(f, "GroupId")
	delete(f, "DelayLoop")
	delete(f, "ConnectTimeout")
	delete(f, "RealServerPorts")
	delete(f, "ClientIPMethod")
	delete(f, "FailoverSwitch")
	delete(f, "HealthyThreshold")
	delete(f, "UnhealthyThreshold")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTCPListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTCPListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回监听器ID
		ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTCPListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTCPListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUDPListenersRequest struct {
	*tchttp.BaseRequest

	// 监听器名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器端口列表
	Ports []*uint64 `json:"Ports,omitempty" name:"Ports"`

	// 监听器源站调度策略，支持轮询（rr），加权轮询（wrr），最小连接数（lc）
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 监听器对应源站类型，支持IP或者DOMAIN类型
	RealServerType *string `json:"RealServerType,omitempty" name:"RealServerType"`

	// 通道ID，ProxyId和GroupId必须设置一个，但不能同时设置。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 通道组ID，ProxyId和GroupId必须设置一个，但不能同时设置。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 源站端口列表，该参数仅支持v1版本监听器和通道组监听器
	RealServerPorts []*uint64 `json:"RealServerPorts,omitempty" name:"RealServerPorts"`
}

func (r *CreateUDPListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUDPListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerName")
	delete(f, "Ports")
	delete(f, "Scheduler")
	delete(f, "RealServerType")
	delete(f, "ProxyId")
	delete(f, "GroupId")
	delete(f, "RealServerPorts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUDPListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateUDPListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回监听器ID
		ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUDPListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUDPListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCertificateRequest struct {
	*tchttp.BaseRequest

	// 需要删除的证书ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *DeleteCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDomainErrorPageInfoRequest struct {
	*tchttp.BaseRequest

	// 定制错误响应页的唯一ID，请参考CreateDomainErrorPageInfo的响应
	ErrorPageId *string `json:"ErrorPageId,omitempty" name:"ErrorPageId"`
}

func (r *DeleteDomainErrorPageInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainErrorPageInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ErrorPageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDomainErrorPageInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDomainErrorPageInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDomainErrorPageInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainErrorPageInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDomainRequest struct {
	*tchttp.BaseRequest

	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 需要删除的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 是否强制删除已绑定源站的转发规则，0非强制，1强制。
	// 当采用非强制删除时，如果域名下已有规则绑定了源站，则无法删除。
	Force *uint64 `json:"Force,omitempty" name:"Force"`
}

func (r *DeleteDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "Domain")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteListenersRequest struct {
	*tchttp.BaseRequest

	// 待删除的监听器ID列表
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// 已绑定源站的监听器是否允许强制删除，1：允许， 0：不允许
	Force *uint64 `json:"Force,omitempty" name:"Force"`

	// 通道组ID，该参数和GroupId必须设置一个，但不能同时设置。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 通道ID，该参数和GroupId必须设置一个，但不能同时设置。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

func (r *DeleteListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerIds")
	delete(f, "Force")
	delete(f, "GroupId")
	delete(f, "ProxyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除操作失败的监听器ID列表
		OperationFailedListenerSet []*string `json:"OperationFailedListenerSet,omitempty" name:"OperationFailedListenerSet"`

		// 删除操作成功的监听器ID列表
		OperationSucceedListenerSet []*string `json:"OperationSucceedListenerSet,omitempty" name:"OperationSucceedListenerSet"`

		// 无效的监听器ID列表，如：监听器不存在，监听器对应实例不匹配
		InvalidStatusListenerSet []*string `json:"InvalidStatusListenerSet,omitempty" name:"InvalidStatusListenerSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProxyGroupRequest struct {
	*tchttp.BaseRequest

	// 需要删除的通道组ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 强制删除标识。其中：
	// 0，不强制删除，
	// 1，强制删除。
	// 默认为0，当通道组中存在通道或通道组中存在监听器/规则绑定了源站时，且Force为0时，该操作会返回失败。
	Force *uint64 `json:"Force,omitempty" name:"Force"`
}

func (r *DeleteProxyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProxyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProxyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProxyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProxyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProxyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleRequest struct {
	*tchttp.BaseRequest

	// 7层监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 是否可以强制删除已绑定源站的转发规则，0非强制，1强制
	Force *uint64 `json:"Force,omitempty" name:"Force"`
}

func (r *DeleteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "RuleId")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DeleteSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityRulesRequest struct {
	*tchttp.BaseRequest

	// 安全策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 访问规则ID列表
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

func (r *DeleteSecurityRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "RuleIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecurityRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessRegionsByDestRegionRequest struct {
	*tchttp.BaseRequest

	// 源站区域：接口DescribeDestRegions返回DestRegionSet中的RegionId字段值
	DestRegion *string `json:"DestRegion,omitempty" name:"DestRegion"`

	// IP版本，可取值：IPv4、IPv6，默认值IPv4
	IPAddressVersion *string `json:"IPAddressVersion,omitempty" name:"IPAddressVersion"`

	// 通道套餐类型，Thunder表示标准通道组，Accelerator表示游戏加速器通道，CrossBorder表示跨境通道。
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`
}

func (r *DescribeAccessRegionsByDestRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessRegionsByDestRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DestRegion")
	delete(f, "IPAddressVersion")
	delete(f, "PackageType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessRegionsByDestRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessRegionsByDestRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可用加速区域数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 可用加速区域信息列表
		AccessRegionSet []*AccessRegionDetial `json:"AccessRegionSet,omitempty" name:"AccessRegionSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessRegionsByDestRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessRegionsByDestRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccessRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 加速区域总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 加速区域详情列表
		AccessRegionSet []*RegionDetail `json:"AccessRegionSet,omitempty" name:"AccessRegionSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlackHeaderRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBlackHeaderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlackHeaderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlackHeaderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlackHeaderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 禁用的自定义header列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		BlackHeaders []*string `json:"BlackHeaders,omitempty" name:"BlackHeaders"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlackHeaderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlackHeaderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificateDetailRequest struct {
	*tchttp.BaseRequest

	// 证书ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *DescribeCertificateDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificateDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificateDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书详情。
		CertificateDetail *CertificateDetail `json:"CertificateDetail,omitempty" name:"CertificateDetail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCertificateDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificatesRequest struct {
	*tchttp.BaseRequest

	// 证书类型。其中：
	// 0，表示基础认证配置；
	// 1，表示客户端CA证书；
	// 2，表示服务器SSL证书；
	// 3，表示源站CA证书；
	// 4，表示通道SSL证书。
	// -1，所有类型。
	// 默认为-1。
	CertificateType *int64 `json:"CertificateType,omitempty" name:"CertificateType"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数量，默认为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务器证书列表，包括证书ID 和证书名称。
		CertificateSet []*Certificate `json:"CertificateSet,omitempty" name:"CertificateSet"`

		// 满足查询条件的服务器证书总数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCountryAreaMappingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCountryAreaMappingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCountryAreaMappingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCountryAreaMappingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCountryAreaMappingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 国家地区编码映射表。
		CountryAreaMappingList []*CountryAreaMap `json:"CountryAreaMappingList,omitempty" name:"CountryAreaMappingList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCountryAreaMappingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCountryAreaMappingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomHeaderRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCustomHeaderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomHeaderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomHeaderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomHeaderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则id
	// 注意：此字段可能返回 null，表示取不到有效值。
		RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

		// 自定义header列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Headers []*HttpHeaderParam `json:"Headers,omitempty" name:"Headers"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomHeaderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomHeaderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDestRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDestRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDestRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDestRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDestRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 源站区域总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 源站区域详情列表
		DestRegionSet []*RegionDetail `json:"DestRegionSet,omitempty" name:"DestRegionSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDestRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDestRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainErrorPageInfoByIdsRequest struct {
	*tchttp.BaseRequest

	// 定制错误ID列表,最多支持10个
	ErrorPageIds []*string `json:"ErrorPageIds,omitempty" name:"ErrorPageIds"`
}

func (r *DescribeDomainErrorPageInfoByIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainErrorPageInfoByIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ErrorPageIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainErrorPageInfoByIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainErrorPageInfoByIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 定制错误响应配置集
	// 注意：此字段可能返回 null，表示取不到有效值。
		ErrorPageSet []*DomainErrorPageInfo `json:"ErrorPageSet,omitempty" name:"ErrorPageSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainErrorPageInfoByIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainErrorPageInfoByIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainErrorPageInfoRequest struct {
	*tchttp.BaseRequest

	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeDomainErrorPageInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainErrorPageInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainErrorPageInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainErrorPageInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 定制错误响应配置集
	// 注意：此字段可能返回 null，表示取不到有效值。
		ErrorPageSet []*DomainErrorPageInfo `json:"ErrorPageSet,omitempty" name:"ErrorPageSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainErrorPageInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainErrorPageInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupAndStatisticsProxyRequest struct {
	*tchttp.BaseRequest

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeGroupAndStatisticsProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupAndStatisticsProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupAndStatisticsProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupAndStatisticsProxyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可以统计的通道组信息
		GroupSet []*GroupStatisticsInfo `json:"GroupSet,omitempty" name:"GroupSet"`

		// 通道组数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupAndStatisticsProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupAndStatisticsProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupDomainConfigRequest struct {
	*tchttp.BaseRequest

	// 通道组ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeGroupDomainConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupDomainConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupDomainConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupDomainConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名解析就近接入配置列表。
		AccessRegionList []*DomainAccessRegionDict `json:"AccessRegionList,omitempty" name:"AccessRegionList"`

		// 默认访问Ip。
		DefaultDnsIp *string `json:"DefaultDnsIp,omitempty" name:"DefaultDnsIp"`

		// 通道组ID。
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// 接入地域的配置的总数。
		AccessRegionCount *int64 `json:"AccessRegionCount,omitempty" name:"AccessRegionCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupDomainConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupDomainConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHTTPListenersRequest struct {
	*tchttp.BaseRequest

	// 通道ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 过滤条件，按照监听器ID进行精确查询
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 过滤条件，按照监听器名称进行精确查询
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 过滤条件，按照监听器端口进行精确查询
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数量，默认为20个
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，支持按照端口或监听器名称进行模糊查询，该参数不能与ListenerName和Port同时使用
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// 通道组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeHTTPListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHTTPListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyId")
	delete(f, "ListenerId")
	delete(f, "ListenerName")
	delete(f, "Port")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchValue")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHTTPListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHTTPListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 监听器数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// HTTP监听器列表
		ListenerSet []*HTTPListener `json:"ListenerSet,omitempty" name:"ListenerSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHTTPListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHTTPListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHTTPSListenersRequest struct {
	*tchttp.BaseRequest

	// 过滤条件，通道ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 过滤条件，根据监听器ID进行精确查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 过滤条件，根据监听器名称进行精确查询。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 过滤条件，根据监听器端口进行精确查询。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 偏移量， 默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数量，默认为20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，支持按照端口或监听器名称进行模糊查询
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// 过滤条件，通道组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 支持Http3的开关，其中：
	// 0，表示不需要支持Http3接入；
	// 1，表示需要支持Http3接入。
	// 注意：如果支持了Http3的功能，那么该监听器会占用对应的UDP接入端口，不可再创建相同端口的UDP监听器。
	// 该功能的启停无法在监听器创建完毕后再修改。
	Http3Supported *int64 `json:"Http3Supported,omitempty" name:"Http3Supported"`
}

func (r *DescribeHTTPSListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHTTPSListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyId")
	delete(f, "ListenerId")
	delete(f, "ListenerName")
	delete(f, "Port")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchValue")
	delete(f, "GroupId")
	delete(f, "Http3Supported")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHTTPSListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHTTPSListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 监听器数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// HTTPS监听器列表
		ListenerSet []*HTTPSListener `json:"ListenerSet,omitempty" name:"ListenerSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHTTPSListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHTTPSListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListenerRealServersRequest struct {
	*tchttp.BaseRequest

	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

func (r *DescribeListenerRealServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerRealServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListenerRealServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListenerRealServersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可绑定源站的个数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 源站信息列表
		RealServerSet []*RealServer `json:"RealServerSet,omitempty" name:"RealServerSet"`

		// 已绑定源站的个数
		BindRealServerTotalCount *uint64 `json:"BindRealServerTotalCount,omitempty" name:"BindRealServerTotalCount"`

		// 已绑定源站信息列表
		BindRealServerSet []*BindRealServer `json:"BindRealServerSet,omitempty" name:"BindRealServerSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListenerRealServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerRealServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListenerStatisticsRequest struct {
	*tchttp.BaseRequest

	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标名称列表，支持: 入带宽:InBandwidth, 出带宽:OutBandwidth, 并发:Concurrent, 入包量:InPackets, 出包量:OutPackets。
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 监控粒度，目前支持300，3600，86400，单位：秒。
	// 查询时间范围不超过1天，支持最小粒度300秒；
	// 查询间范围不超过7天，支持最小粒度3600秒；
	// 查询间范围超过7天，支持最小粒度86400秒。
	Granularity *uint64 `json:"Granularity,omitempty" name:"Granularity"`
}

func (r *DescribeListenerStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "Granularity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListenerStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListenerStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道组统计数据
		StatisticsData []*MetricStatisticsInfo `json:"StatisticsData,omitempty" name:"StatisticsData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListenerStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxiesRequest struct {
	*tchttp.BaseRequest

	// （旧参数，请切换到ProxyIds）按照一个或者多个实例ID查询。每次请求的实例的上限为100。参数不支持同时指定InstanceIds和Filters。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件。   
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定InstanceIds和Filters。 
	// ProjectId - String - 是否必填：否 -（过滤条件）按照项目ID过滤。   
	// AccessRegion - String - 是否必填：否 - （过滤条件）按照接入地域过滤。    
	// RealServerRegion - String - 是否必填：否 - （过滤条件）按照源站地域过滤。
	// GroupId - String - 是否必填：否 - （过滤条件）按照通道组ID过滤。
	// IPAddressVersion - String - 是否必填：否 - （过滤条件）按照IP版本过滤。
	// PackageType - String - 是否必填：否 - （过滤条件）按照通道套餐类型过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// （新参数，替代InstanceIds）按照一个或者多个实例ID查询。每次请求的实例的上限为100。参数不支持同时指定InstanceIds和Filters。
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`

	// 标签列表，当存在该字段时，拉取对应标签下的资源列表。
	// 最多支持5个标签，当存在两个或两个以上的标签时，满足其中任意一个标签时，通道会被拉取出来。
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet"`

	// 当该字段为1时，仅拉取非通道组的通道，
	// 当该字段为0时，仅拉取通道组的通道，
	// 不存在该字段时，拉取所有通道，包括独立通道和通道组通道。
	Independent *int64 `json:"Independent,omitempty" name:"Independent"`

	// 输出通道列表的排列顺序。取值范围：
	// asc：升序排列；
	// desc：降序排列。
	// 默认为降序。
	Order *string `json:"Order,omitempty" name:"Order"`

	// 通道列表排序的依据字段。取值范围：
	// create_time：依据通道的创建时间排序；
	// proxy_id：依据通道的ID排序；
	// bandwidth：依据通道带宽上限排序；
	// concurrent_connections：依据通道并发排序；
	// 默认按通道创建时间排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
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
	delete(f, "InstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "ProxyIds")
	delete(f, "TagSet")
	delete(f, "Independent")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道个数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// （旧参数，请切换到ProxySet）通道实例信息列表。
		InstanceSet []*ProxyInfo `json:"InstanceSet,omitempty" name:"InstanceSet"`

		// （新参数）通道实例信息列表。
		ProxySet []*ProxyInfo `json:"ProxySet,omitempty" name:"ProxySet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeProxiesStatusRequest struct {
	*tchttp.BaseRequest

	// （旧参数，请切换到ProxyIds）通道ID列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// （新参数）通道ID列表。
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`
}

func (r *DescribeProxiesStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxiesStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ProxyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxiesStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxiesStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道状态列表。
		InstanceStatusSet []*ProxyStatus `json:"InstanceStatusSet,omitempty" name:"InstanceStatusSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProxiesStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxiesStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyAndStatisticsListenersRequest struct {
	*tchttp.BaseRequest

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeProxyAndStatisticsListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyAndStatisticsListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxyAndStatisticsListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyAndStatisticsListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可以统计的通道信息
		ProxySet []*ProxySimpleInfo `json:"ProxySet,omitempty" name:"ProxySet"`

		// 通道数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProxyAndStatisticsListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyAndStatisticsListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyDetailRequest struct {
	*tchttp.BaseRequest

	// 需查询的通道ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

func (r *DescribeProxyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxyDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道详情信息。
		ProxyDetail *ProxyInfo `json:"ProxyDetail,omitempty" name:"ProxyDetail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProxyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyGroupDetailsRequest struct {
	*tchttp.BaseRequest

	// 通道组ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeProxyGroupDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyGroupDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxyGroupDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyGroupDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道组详细信息。
		ProxyGroupDetail *ProxyGroupDetail `json:"ProxyGroupDetail,omitempty" name:"ProxyGroupDetail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProxyGroupDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyGroupDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyGroupListRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 项目ID。取值范围：
	// -1，该用户下所有项目
	// 0，默认项目
	// 其他值，指定的项目
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 过滤条件。   
	// 每次请求的Filter.Values的上限为5。
	// RealServerRegion - String - 是否必填：否 -（过滤条件）按照源站地域过滤，可参考DescribeDestRegions接口返回结果中的RegionId。
	// PackageType - String - 是否必填：否 - （过滤条件）通道组类型，Thunder表示标准通道组，Accelerator表示游戏加速器通道。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 标签列表，当存在该字段时，拉取对应标签下的资源列表。
	// 最多支持5个标签，当存在两个或两个以上的标签时，满足其中任意一个标签时，该通道组会被拉取出来。
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet"`
}

func (r *DescribeProxyGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProjectId")
	delete(f, "Filters")
	delete(f, "TagSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxyGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyGroupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道组总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 通道组列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ProxyGroupList []*ProxyGroupInfo `json:"ProxyGroupList,omitempty" name:"ProxyGroupList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProxyGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyGroupStatisticsRequest struct {
	*tchttp.BaseRequest

	// 通道组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标名称列表，支持: 入带宽:InBandwidth, 出带宽:OutBandwidth, 并发:Concurrent, 入包量:InPackets, 出包量:OutPackets
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 监控粒度，目前支持60，300，3600，86400，单位：秒。
	// 当时间范围不超过1天，支持最小粒度60秒；
	// 当时间范围不超过7天，支持最小粒度3600秒；
	// 当时间范围不超过30天，支持最小粒度86400秒。
	Granularity *uint64 `json:"Granularity,omitempty" name:"Granularity"`
}

func (r *DescribeProxyGroupStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyGroupStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "Granularity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxyGroupStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyGroupStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道组统计数据
		StatisticsData []*MetricStatisticsInfo `json:"StatisticsData,omitempty" name:"StatisticsData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProxyGroupStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyGroupStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyStatisticsRequest struct {
	*tchttp.BaseRequest

	// 通道ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 起始时间(2019-03-25 12:00:00)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间(2019-03-25 12:00:00)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标名称列表，支持: 入带宽:InBandwidth, 出带宽:OutBandwidth, 并发:Concurrent, 入包量:InPackets, 出包量:OutPackets, 丢包率:PacketLoss, 延迟:Latency，http请求量：HttpQPS, Https请求量：HttpsQPS
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 监控粒度，目前支持60，300，3600，86400，单位：秒。
	// 当时间范围不超过3天，支持最小粒度60秒；
	// 当时间范围不超过7天，支持最小粒度300秒；
	// 当时间范围不超过30天，支持最小粒度3600秒。
	Granularity *uint64 `json:"Granularity,omitempty" name:"Granularity"`

	// 运营商（通道为三网通道时有效），支持CMCC，CUCC，CTCC，传空值或不传则合并三个运营商数据
	Isp *string `json:"Isp,omitempty" name:"Isp"`
}

func (r *DescribeProxyStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "Granularity")
	delete(f, "Isp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxyStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道统计数据
		StatisticsData []*MetricStatisticsInfo `json:"StatisticsData,omitempty" name:"StatisticsData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProxyStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealServerStatisticsRequest struct {
	*tchttp.BaseRequest

	// 源站ID
	RealServerId *string `json:"RealServerId,omitempty" name:"RealServerId"`

	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// L7层规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 统计时长，单位：小时。仅支持最近1,3,6,12,24小时的统计查询
	WithinTime *uint64 `json:"WithinTime,omitempty" name:"WithinTime"`

	// 统计开始时间(2020-08-19 00:00:00)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间(2020-08-19 23:59:59)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计的数据粒度，单位：秒，仅支持1分钟-60和5分钟-300粒度
	Granularity *uint64 `json:"Granularity,omitempty" name:"Granularity"`
}

func (r *DescribeRealServerStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealServerStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RealServerId")
	delete(f, "ListenerId")
	delete(f, "RuleId")
	delete(f, "WithinTime")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Granularity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRealServerStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealServerStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 指定监听器的源站状态统计数据
		StatisticsData []*StatisticsDataInfo `json:"StatisticsData,omitempty" name:"StatisticsData"`

		// 多个源站状态统计数据
		RsStatisticsData []*MetricStatisticsInfo `json:"RsStatisticsData,omitempty" name:"RsStatisticsData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealServerStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealServerStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealServersRequest struct {
	*tchttp.BaseRequest

	// 查询源站的所属项目ID，-1表示所有项目
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 需要查询的源站IP或域名，支持模糊匹配
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// 偏移量，默认值是0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20个，最大值为50个
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 标签列表，当存在该字段时，拉取对应标签下的资源列表。
	// 最多支持5个标签，当存在两个或两个以上的标签时，满足其中任意一个标签时，源站会被拉取出来。
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet"`

	// 过滤条件。filter的name取值(RealServerName,RealServerIP)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRealServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "SearchValue")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TagSet")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRealServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealServersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 源站信息列表
		RealServerSet []*BindRealServerInfo `json:"RealServerSet,omitempty" name:"RealServerSet"`

		// 查询得到的源站数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealServersStatusRequest struct {
	*tchttp.BaseRequest

	// 源站ID列表
	RealServerIds []*string `json:"RealServerIds,omitempty" name:"RealServerIds"`
}

func (r *DescribeRealServersStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealServersStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RealServerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRealServersStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealServersStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回源站查询结果的个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 源站被绑定状态列表
		RealServerStatusSet []*RealServerStatus `json:"RealServerStatusSet,omitempty" name:"RealServerStatusSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealServersStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealServersStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionAndPriceRequest struct {
	*tchttp.BaseRequest

	// IP版本，可取值：IPv4、IPv6，默认值IPv4
	IPAddressVersion *string `json:"IPAddressVersion,omitempty" name:"IPAddressVersion"`

	// 通道套餐类型，Thunder表示标准通道组，Accelerator表示游戏加速器通道，CrossBorder表示跨境通道。
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`
}

func (r *DescribeRegionAndPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionAndPriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IPAddressVersion")
	delete(f, "PackageType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionAndPriceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionAndPriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 源站区域总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 源站区域详情列表
		DestRegionSet []*RegionDetail `json:"DestRegionSet,omitempty" name:"DestRegionSet"`

		// 通道带宽费用梯度价格
		BandwidthUnitPrice []*BandwidthPriceGradient `json:"BandwidthUnitPrice,omitempty" name:"BandwidthUnitPrice"`

		// 带宽价格货币类型：
	// CNY 人民币
	// USD 美元
		Currency *string `json:"Currency,omitempty" name:"Currency"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionAndPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionAndPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesByTagRequest struct {
	*tchttp.BaseRequest

	// 标签键。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// 资源类型，其中：
	// Proxy表示通道；
	// ProxyGroup表示通道组；
	// RealServer表示源站。
	// 不指定该字段则查询该标签下所有资源。
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *DescribeResourcesByTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "TagValue")
	delete(f, "ResourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcesByTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesByTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 资源总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 标签对应的资源列表
		ResourceSet []*TagResourceInfo `json:"ResourceSet,omitempty" name:"ResourceSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourcesByTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleRealServersRequest struct {
	*tchttp.BaseRequest

	// 转发规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRuleRealServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleRealServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleRealServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleRealServersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可绑定的源站个数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 可绑定的源站信息列表
		RealServerSet []*RealServer `json:"RealServerSet,omitempty" name:"RealServerSet"`

		// 已绑定的源站个数
		BindRealServerTotalCount *uint64 `json:"BindRealServerTotalCount,omitempty" name:"BindRealServerTotalCount"`

		// 已绑定的源站信息列表
		BindRealServerSet []*BindRealServer `json:"BindRealServerSet,omitempty" name:"BindRealServerSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRuleRealServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleRealServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRulesByRuleIdsRequest struct {
	*tchttp.BaseRequest

	// 规则ID列表。最多支持10个规则。
	RuleIds []*string `json:"RuleIds,omitempty" name:"RuleIds"`
}

func (r *DescribeRulesByRuleIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesByRuleIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRulesByRuleIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRulesByRuleIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的规则总个数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回的规则列表。
		RuleSet []*RuleInfo `json:"RuleSet,omitempty" name:"RuleSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRulesByRuleIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesByRuleIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRulesRequest struct {
	*tchttp.BaseRequest

	// 7层监听器Id。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

func (r *DescribeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 按照域名分类的规则信息列表
		DomainRuleSet []*DomainRuleSet `json:"DomainRuleSet,omitempty" name:"DomainRuleSet"`

		// 该监听器下的域名总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityPolicyDetailRequest struct {
	*tchttp.BaseRequest

	// 安全策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DescribeSecurityPolicyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPolicyDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityPolicyDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

		// 安全策略状态：
	// BOUND，已开启安全策略
	// UNBIND，已关闭安全策略
	// BINDING，安全策略开启中
	// UNBINDING，安全策略关闭中。
		Status *string `json:"Status,omitempty" name:"Status"`

		// 默认策略：ACCEPT或DROP。
		DefaultAction *string `json:"DefaultAction,omitempty" name:"DefaultAction"`

		// 策略ID
		PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

		// 规则列表
		RuleList []*SecurityPolicyRuleOut `json:"RuleList,omitempty" name:"RuleList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityPolicyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityRulesRequest struct {
	*tchttp.BaseRequest

	// 安全规则ID列表。总数不能超过20个。
	SecurityRuleIds []*string `json:"SecurityRuleIds,omitempty" name:"SecurityRuleIds"`
}

func (r *DescribeSecurityRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityRuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的安全规则详情总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回的安全规则详情列表。
		SecurityRuleSet []*SecurityPolicyRuleOut `json:"SecurityRuleSet,omitempty" name:"SecurityRuleSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTCPListenersRequest struct {
	*tchttp.BaseRequest

	// 过滤条件，根据通道ID进行拉取，ProxyId/GroupId/ListenerId必须设置一个，但ProxyId和GroupId不能同时设置。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 过滤条件，根据监听器ID精确查询。
	// 当设置了ProxyId时，会检查该监听器是否归属于该通道。
	// 当设置了GroupId时，会检查该监听器是否归属于该通道组。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 过滤条件，根据监听器名称精确查询
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 过滤条件，根据监听器端口精确查询
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数量，默认为20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，根据通道组ID进行拉取，ProxyId/GroupId/ListenerId必须设置一个，但ProxyId和GroupId不能同时设置。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 过滤条件，支持按照端口或监听器名称进行模糊查询，该参数不能与ListenerName和Port同时使用
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

func (r *DescribeTCPListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTCPListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyId")
	delete(f, "ListenerId")
	delete(f, "ListenerName")
	delete(f, "Port")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "GroupId")
	delete(f, "SearchValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTCPListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTCPListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 满足条件的监听器总个数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// TCP监听器列表
		ListenerSet []*TCPListener `json:"ListenerSet,omitempty" name:"ListenerSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTCPListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTCPListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUDPListenersRequest struct {
	*tchttp.BaseRequest

	// 过滤条件，根据通道ID进行拉取，ProxyId/GroupId/ListenerId必须设置一个，但ProxyId和GroupId不能同时设置。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 过滤条件，根据监听器ID精确查询。
	// 当设置了ProxyId时，会检查该监听器是否归属于该通道。
	// 当设置了GroupId时，会检查该监听器是否归属于该通道组。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 过滤条件，根据监听器名称精确查询
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 过滤条件，根据监听器端口精确查询
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数量，默认为20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，根据通道组ID进行拉取，ProxyId/GroupId/ListenerId必须设置一个，但ProxyId和GroupId不能同时设置。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 过滤条件，支持按照端口或监听器名称进行模糊查询，该参数不能与ListenerName和Port同时使用
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

func (r *DescribeUDPListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUDPListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyId")
	delete(f, "ListenerId")
	delete(f, "ListenerName")
	delete(f, "Port")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "GroupId")
	delete(f, "SearchValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUDPListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUDPListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 监听器个数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// UDP监听器列表
		ListenerSet []*UDPListener `json:"ListenerSet,omitempty" name:"ListenerSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUDPListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUDPListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyProxiesRequest struct {
	*tchttp.BaseRequest

	// 强制删除标识。
	// 1，强制删除该通道列表，无论是否已经绑定了源站；
	// 0，如果已绑定了源站，则无法删除。
	// 删除多通道时，如果该标识为0，只有所有的通道都没有绑定源站，才允许删除。
	Force *int64 `json:"Force,omitempty" name:"Force"`

	// （旧参数，请切换到ProxyIds）通道实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	// 更多详细信息请参阅：如何保证幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// （新参数）通道实例ID列表。
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`
}

func (r *DestroyProxiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyProxiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Force")
	delete(f, "InstanceIds")
	delete(f, "ClientToken")
	delete(f, "ProxyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyProxiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DestroyProxiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 处于不可销毁状态下的通道实例ID列表。
		InvalidStatusInstanceSet []*string `json:"InvalidStatusInstanceSet,omitempty" name:"InvalidStatusInstanceSet"`

		// 销毁操作失败的通道实例ID列表。
		OperationFailedInstanceSet []*string `json:"OperationFailedInstanceSet,omitempty" name:"OperationFailedInstanceSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyProxiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyProxiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainAccessRegionDict struct {

	// 就近接入区域
	NationCountryInnerList []*NationCountryInnerInfo `json:"NationCountryInnerList,omitempty" name:"NationCountryInnerList"`

	// 加速区域通道列表
	ProxyList []*ProxyIdDict `json:"ProxyList,omitempty" name:"ProxyList"`

	// 加速区域ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 加速区域内部编码
	GeographicalZoneInnerCode *string `json:"GeographicalZoneInnerCode,omitempty" name:"GeographicalZoneInnerCode"`

	// 加速区域所属大洲内部编码
	ContinentInnerCode *string `json:"ContinentInnerCode,omitempty" name:"ContinentInnerCode"`

	// 加速区域别名
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

type DomainErrorPageInfo struct {

	// 错误定制响应的配置ID
	ErrorPageId *string `json:"ErrorPageId,omitempty" name:"ErrorPageId"`

	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 原始错误码
	ErrorNos []*int64 `json:"ErrorNos,omitempty" name:"ErrorNos"`

	// 新的错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewErrorNo *int64 `json:"NewErrorNo,omitempty" name:"NewErrorNo"`

	// 需要清理的响应头
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClearHeaders []*string `json:"ClearHeaders,omitempty" name:"ClearHeaders"`

	// 需要设置的响应头
	// 注意：此字段可能返回 null，表示取不到有效值。
	SetHeaders []*HttpHeaderParam `json:"SetHeaders,omitempty" name:"SetHeaders"`

	// 设置的响应体(不包括 HTTP头)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Body *string `json:"Body,omitempty" name:"Body"`

	// 规则状态,0为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DomainRuleSet struct {

	// 转发规则域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 该域名对应的转发规则列表。
	RuleSet []*RuleInfo `json:"RuleSet,omitempty" name:"RuleSet"`

	// 该域名对应的服务器证书ID，值为default时，表示使用默认证书（监听器配置的证书）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 该域名对应服务器证书名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateAlias *string `json:"CertificateAlias,omitempty" name:"CertificateAlias"`

	// 该域名对应的客户端证书ID，值为default时，表示使用默认证书（监听器配置的证书）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientCertificateId *string `json:"ClientCertificateId,omitempty" name:"ClientCertificateId"`

	// 该域名对应客户端证书名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientCertificateAlias *string `json:"ClientCertificateAlias,omitempty" name:"ClientCertificateAlias"`

	// 该域名对应基础认证配置ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BasicAuthConfId *string `json:"BasicAuthConfId,omitempty" name:"BasicAuthConfId"`

	// 基础认证开关，其中：
	// 0，表示未开启；
	// 1，表示已开启。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BasicAuth *int64 `json:"BasicAuth,omitempty" name:"BasicAuth"`

	// 该域名对应基础认证配置名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BasicAuthConfAlias *string `json:"BasicAuthConfAlias,omitempty" name:"BasicAuthConfAlias"`

	// 该域名对应源站认证证书ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealServerCertificateId *string `json:"RealServerCertificateId,omitempty" name:"RealServerCertificateId"`

	// 源站认证开关，其中：
	// 0，表示未开启；
	// 1，表示已开启。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealServerAuth *int64 `json:"RealServerAuth,omitempty" name:"RealServerAuth"`

	// 该域名对应源站认证证书名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealServerCertificateAlias *string `json:"RealServerCertificateAlias,omitempty" name:"RealServerCertificateAlias"`

	// 该域名对应通道认证证书ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	GaapCertificateId *string `json:"GaapCertificateId,omitempty" name:"GaapCertificateId"`

	// 通道认证开关，其中：
	// 0，表示未开启；
	// 1，表示已开启。
	// 注意：此字段可能返回 null，表示取不到有效值。
	GaapAuth *int64 `json:"GaapAuth,omitempty" name:"GaapAuth"`

	// 该域名对应通道认证证书名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	GaapCertificateAlias *string `json:"GaapCertificateAlias,omitempty" name:"GaapCertificateAlias"`

	// 源站认证域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealServerCertificateDomain *string `json:"RealServerCertificateDomain,omitempty" name:"RealServerCertificateDomain"`

	// 多客户端证书时，返回多个证书的id和别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolyClientCertificateAliasInfo []*CertificateAliasInfo `json:"PolyClientCertificateAliasInfo,omitempty" name:"PolyClientCertificateAliasInfo"`

	// 多源站证书时，返回多个证书的id和别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolyRealServerCertificateAliasInfo []*CertificateAliasInfo `json:"PolyRealServerCertificateAliasInfo,omitempty" name:"PolyRealServerCertificateAliasInfo"`

	// 域名的状态。
	// 0表示运行中，
	// 1表示变更中，
	// 2表示删除中。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainStatus *uint64 `json:"DomainStatus,omitempty" name:"DomainStatus"`

	// 封禁解封状态：BANNED表示已封禁，RECOVER表示已解封或未封禁，BANNING表示封禁中，RECOVERING表示解封中，BAN_FAILED表示封禁失败，RECOVER_FAILED表示解封失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BanStatus *string `json:"BanStatus,omitempty" name:"BanStatus"`

	// Http3特性标识，其中：
	// 0表示关闭；
	// 1表示启用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Http3Supported *int64 `json:"Http3Supported,omitempty" name:"Http3Supported"`
}

type Filter struct {

	// 过滤条件
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤值
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type GroupStatisticsInfo struct {

	// 通道组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 通道组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 通道组下通道列表
	ProxySet []*ProxySimpleInfo `json:"ProxySet,omitempty" name:"ProxySet"`
}

type HTTPListener struct {

	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 监听器创建时间，Unix时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 监听器协议， HTTP表示HTTP，HTTPS表示HTTPS，此结构取值HTTP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器状态，其中：
	// 0表示运行中；
	// 1表示创建中；
	// 2表示销毁中；
	// 3表示源站调整中；
	// 4表示配置变更中。
	ListenerStatus *uint64 `json:"ListenerStatus,omitempty" name:"ListenerStatus"`
}

type HTTPSListener struct {

	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 监听器协议， HTTP表示HTTP，HTTPS表示HTTPS，此结构取值HTTPS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器状态，其中：
	// 0表示运行中；
	// 1表示创建中；
	// 2表示销毁中；
	// 3表示源站调整中；
	// 4表示配置变更中。
	ListenerStatus *uint64 `json:"ListenerStatus,omitempty" name:"ListenerStatus"`

	// 监听器服务器SSL证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 监听器后端转发源站协议
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`

	// 监听器创建时间，Unix时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 服务器SSL证书的别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateAlias *string `json:"CertificateAlias,omitempty" name:"CertificateAlias"`

	// 监听器客户端CA证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientCertificateId *string `json:"ClientCertificateId,omitempty" name:"ClientCertificateId"`

	// 监听器认证方式。其中，
	// 0表示单向认证；
	// 1表示双向认证。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthType *int64 `json:"AuthType,omitempty" name:"AuthType"`

	// 客户端CA证书别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientCertificateAlias *string `json:"ClientCertificateAlias,omitempty" name:"ClientCertificateAlias"`

	// 多客户端CA证书别名信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolyClientCertificateAliasInfo []*CertificateAliasInfo `json:"PolyClientCertificateAliasInfo,omitempty" name:"PolyClientCertificateAliasInfo"`
}

type HttpHeaderParam struct {

	// HTTP头名
	HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`

	// HTTP头值
	HeaderValue *string `json:"HeaderValue,omitempty" name:"HeaderValue"`
}

type IPDetail struct {

	// IP字符串
	IP *string `json:"IP,omitempty" name:"IP"`

	// 供应商，BGP表示默认，CMCC表示中国移动，CUCC表示中国联通，CTCC表示中国电信
	Provider *string `json:"Provider,omitempty" name:"Provider"`

	// 带宽
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
}

type InquiryPriceCreateProxyRequest struct {
	*tchttp.BaseRequest

	// 加速区域名称。
	AccessRegion *string `json:"AccessRegion,omitempty" name:"AccessRegion"`

	// 通道带宽上限，单位：Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// （旧参数，请切换到RealServerRegion）源站区域名称。
	DestRegion *string `json:"DestRegion,omitempty" name:"DestRegion"`

	// （旧参数，请切换到Concurrent）通道并发量上限，表示同时在线的连接数，单位：万。
	Concurrency *int64 `json:"Concurrency,omitempty" name:"Concurrency"`

	// （新参数）源站区域名称。
	RealServerRegion *string `json:"RealServerRegion,omitempty" name:"RealServerRegion"`

	// （新参数）通道并发量上限，表示同时在线的连接数，单位：万。
	Concurrent *int64 `json:"Concurrent,omitempty" name:"Concurrent"`

	// 计费方式，0表示按带宽计费，1表示按流量计费。默认按带宽计费
	BillingType *int64 `json:"BillingType,omitempty" name:"BillingType"`

	// IP版本，可取值：IPv4、IPv6，默认值IPv4
	IPAddressVersion *string `json:"IPAddressVersion,omitempty" name:"IPAddressVersion"`

	// 网络类型，可取值：normal、cn2，默认值normal
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// 通道套餐类型，Thunder表示标准通道组，Accelerator表示游戏加速器通道，CrossBorder表示跨境通道。
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// 支持Http3的开关，其中：0，表示不需要支持Http3接入；1，表示需要支持Http3接入。注意：如果开启了Http3的功能，那么该通道就不再支持TCP/UDP接入的功能。该功能的启停无法在通道创建完毕后再修改。
	Http3Supported *int64 `json:"Http3Supported,omitempty" name:"Http3Supported"`
}

func (r *InquiryPriceCreateProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessRegion")
	delete(f, "Bandwidth")
	delete(f, "DestRegion")
	delete(f, "Concurrency")
	delete(f, "RealServerRegion")
	delete(f, "Concurrent")
	delete(f, "BillingType")
	delete(f, "IPAddressVersion")
	delete(f, "NetworkType")
	delete(f, "PackageType")
	delete(f, "Http3Supported")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateProxyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道基础费用价格，单位：元/天。
		ProxyDailyPrice *float64 `json:"ProxyDailyPrice,omitempty" name:"ProxyDailyPrice"`

		// 通道带宽费用梯度价格。
	// 注意：此字段可能返回 null，表示取不到有效值。
		BandwidthUnitPrice []*BandwidthPriceGradient `json:"BandwidthUnitPrice,omitempty" name:"BandwidthUnitPrice"`

		// 通道基础费用折扣价格，单位：元/天。
		DiscountProxyDailyPrice *float64 `json:"DiscountProxyDailyPrice,omitempty" name:"DiscountProxyDailyPrice"`

		// 价格使用的货币，支持人民币，美元等。
		Currency *string `json:"Currency,omitempty" name:"Currency"`

		// 通道的流量费用价格，单位: 元/GB
	// 注意：此字段可能返回 null，表示取不到有效值。
		FlowUnitPrice *float64 `json:"FlowUnitPrice,omitempty" name:"FlowUnitPrice"`

		// 通道的流量费用折扣价格，单位:元/GB
	// 注意：此字段可能返回 null，表示取不到有效值。
		DiscountFlowUnitPrice *float64 `json:"DiscountFlowUnitPrice,omitempty" name:"DiscountFlowUnitPrice"`

		// 精品BGP的带宽费用价格，单位: 元/Mbps/天
	// 注意：此字段可能返回 null，表示取不到有效值。
		Cn2BandwidthPrice *float64 `json:"Cn2BandwidthPrice,omitempty" name:"Cn2BandwidthPrice"`

		// 精品BGP的折后带宽费用价格，单位: 元/Mbps/天
	// 注意：此字段可能返回 null，表示取不到有效值。
		Cn2BandwidthPriceWithDiscount *float64 `json:"Cn2BandwidthPriceWithDiscount,omitempty" name:"Cn2BandwidthPriceWithDiscount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListenerInfo struct {

	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器监听端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 监听器协议类型
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type MetricStatisticsInfo struct {

	// 指标名称
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 指标统计数据
	MetricData []*StatisticsDataInfo `json:"MetricData,omitempty" name:"MetricData"`
}

type ModifyCertificateAttributesRequest struct {
	*tchttp.BaseRequest

	// 证书ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 证书名字。长度不超过50个字符。
	CertificateAlias *string `json:"CertificateAlias,omitempty" name:"CertificateAlias"`
}

func (r *ModifyCertificateAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "CertificateAlias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCertificateAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCertificateAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCertificateAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCertificateRequest struct {
	*tchttp.BaseRequest

	// 监听器实例ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 需要修改证书的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 新的服务器证书ID。其中：
	// 当CertificateId=default时，表示使用监听器的证书。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 新的客户端证书ID。其中：
	// 当ClientCertificateId=default时，表示使用监听器的证书。
	// 仅当采用双向认证方式时，需要设置该参数或者PolyClientCertificateIds。
	ClientCertificateId *string `json:"ClientCertificateId,omitempty" name:"ClientCertificateId"`

	// 新的多客户端证书ID列表。其中：
	// 仅当采用双向认证方式时，需要设置该参数或ClientCertificateId参数。
	PolyClientCertificateIds []*string `json:"PolyClientCertificateIds,omitempty" name:"PolyClientCertificateIds"`
}

func (r *ModifyCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "Domain")
	delete(f, "CertificateId")
	delete(f, "ClientCertificateId")
	delete(f, "PolyClientCertificateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainRequest struct {
	*tchttp.BaseRequest

	// 7层监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 修改前的域名信息
	OldDomain *string `json:"OldDomain,omitempty" name:"OldDomain"`

	// 修改后的域名信息
	NewDomain *string `json:"NewDomain,omitempty" name:"NewDomain"`

	// 服务器SSL证书ID，仅适用于version3.0的通道。其中：
	// 不带该字段时，表示使用原证书；
	// 携带该字段时并且CertificateId=default，表示使用监听器证书；
	// 其他情况，使用该CertificateId指定的证书。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 客户端CA证书ID，仅适用于version3.0的通道。其中：
	// 不带该字段和PolyClientCertificateIds时，表示使用原证书；
	// 携带该字段时并且ClientCertificateId=default，表示使用监听器证书；
	// 其他情况，使用该ClientCertificateId或PolyClientCertificateIds指定的证书。
	ClientCertificateId *string `json:"ClientCertificateId,omitempty" name:"ClientCertificateId"`

	// 客户端CA证书ID，仅适用于version3.0的通道。其中：
	// 不带该字段和ClientCertificateId时，表示使用原证书；
	// 携带该字段时并且ClientCertificateId=default，表示使用监听器证书；
	// 其他情况，使用该ClientCertificateId或PolyClientCertificateIds指定的证书。
	PolyClientCertificateIds []*string `json:"PolyClientCertificateIds,omitempty" name:"PolyClientCertificateIds"`
}

func (r *ModifyDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "OldDomain")
	delete(f, "NewDomain")
	delete(f, "CertificateId")
	delete(f, "ClientCertificateId")
	delete(f, "PolyClientCertificateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGroupDomainConfigRequest struct {
	*tchttp.BaseRequest

	// 通道组ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 域名解析默认访问IP或域名。
	DefaultDnsIp *string `json:"DefaultDnsIp,omitempty" name:"DefaultDnsIp"`

	// 就近接入区域配置。
	AccessRegionList []*AccessRegionDomainConf `json:"AccessRegionList,omitempty" name:"AccessRegionList"`
}

func (r *ModifyGroupDomainConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupDomainConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "DefaultDnsIp")
	delete(f, "AccessRegionList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGroupDomainConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGroupDomainConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGroupDomainConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupDomainConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHTTPListenerAttributeRequest struct {
	*tchttp.BaseRequest

	// 需要修改的监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 新的监听器名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 通道ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

func (r *ModifyHTTPListenerAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHTTPListenerAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "ListenerName")
	delete(f, "ProxyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHTTPListenerAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHTTPListenerAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHTTPListenerAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHTTPListenerAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHTTPSListenerAttributeRequest struct {
	*tchttp.BaseRequest

	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 通道ID， 若为单通道监听器，此项必须填写
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 修改后的监听器名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器后端转发与源站之间的协议类型
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`

	// 修改后的监听器服务器证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 修改后的监听器客户端证书ID，不支持多客户端证书，多客户端证书新采用PolyClientCertificateIds字段
	ClientCertificateId *string `json:"ClientCertificateId,omitempty" name:"ClientCertificateId"`

	// 新字段,修改后的监听器客户端证书ID
	PolyClientCertificateIds []*string `json:"PolyClientCertificateIds,omitempty" name:"PolyClientCertificateIds"`
}

func (r *ModifyHTTPSListenerAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHTTPSListenerAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "ProxyId")
	delete(f, "ListenerName")
	delete(f, "ForwardProtocol")
	delete(f, "CertificateId")
	delete(f, "ClientCertificateId")
	delete(f, "PolyClientCertificateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHTTPSListenerAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHTTPSListenerAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHTTPSListenerAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHTTPSListenerAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProxiesAttributeRequest struct {
	*tchttp.BaseRequest

	// （旧参数，请切换到ProxyIds）一个或多个待操作的通道ID。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 通道名称。可任意命名，但不得超过30个字符。
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	// 更多详细信息请参阅：如何保证幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// （新参数）一个或多个待操作的通道ID。
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`
}

func (r *ModifyProxiesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxiesAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ProxyName")
	delete(f, "ClientToken")
	delete(f, "ProxyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProxiesAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProxiesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProxiesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxiesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProxiesProjectRequest struct {
	*tchttp.BaseRequest

	// 需要修改到的项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// （旧参数，请切换到ProxyIds）一个或多个待操作的通道ID。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	// 更多详细信息请参阅：如何保证幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// （新参数）一个或多个待操作的通道ID。
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`
}

func (r *ModifyProxiesProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxiesProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceIds")
	delete(f, "ClientToken")
	delete(f, "ProxyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProxiesProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProxiesProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProxiesProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxiesProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProxyConfigurationRequest struct {
	*tchttp.BaseRequest

	// （旧参数，请切换到ProxyId）通道的实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 需要调整到的目标带宽，单位：Mbps。
	// Bandwidth与Concurrent必须至少设置一个。取值范围根据DescribeAccessRegionsByDestRegion接口获取得到
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 需要调整到的目标并发值，单位：万。
	// Bandwidth与Concurrent必须至少设置一个。取值范围根据DescribeAccessRegionsByDestRegion接口获取得到
	Concurrent *uint64 `json:"Concurrent,omitempty" name:"Concurrent"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	// 更多详细信息请参阅：如何保证幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// （新参数）通道的实例ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 计费方式 (0:按带宽计费，1:按流量计费 默认按带宽计费）
	BillingType *int64 `json:"BillingType,omitempty" name:"BillingType"`
}

func (r *ModifyProxyConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Bandwidth")
	delete(f, "Concurrent")
	delete(f, "ClientToken")
	delete(f, "ProxyId")
	delete(f, "BillingType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProxyConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProxyConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProxyConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProxyGroupAttributeRequest struct {
	*tchttp.BaseRequest

	// 需要修改的通道组ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 修改后的通道组名称：不超过30个字符，超过部分会被截断。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyProxyGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProxyGroupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProxyGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProxyGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRealServerNameRequest struct {
	*tchttp.BaseRequest

	// 源站名称
	RealServerName *string `json:"RealServerName,omitempty" name:"RealServerName"`

	// 源站ID
	RealServerId *string `json:"RealServerId,omitempty" name:"RealServerId"`
}

func (r *ModifyRealServerNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRealServerNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RealServerName")
	delete(f, "RealServerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRealServerNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRealServerNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRealServerNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRealServerNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRuleAttributeRequest struct {
	*tchttp.BaseRequest

	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 调度策略，其中：
	// rr，轮询；
	// wrr，加权轮询；
	// lc，最小连接数。
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 源站健康检查开关，其中：
	// 1，开启；
	// 0，关闭。
	HealthCheck *uint64 `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 健康检查配置参数
	CheckParams *RuleCheckParams `json:"CheckParams,omitempty" name:"CheckParams"`

	// 转发规则路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 加速通道转发到源站的协议类型，支持：default, HTTP和HTTPS。
	// 当ForwardProtocol=default时，表示使用对应监听器的ForwardProtocol。
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`

	// 加速通道转发到源站的请求中携带的host。
	// 当ForwardHost=default时，使用规则的域名，其他情况为该字段所设置的值。
	ForwardHost *string `json:"ForwardHost,omitempty" name:"ForwardHost"`
}

func (r *ModifyRuleAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "RuleId")
	delete(f, "Scheduler")
	delete(f, "HealthCheck")
	delete(f, "CheckParams")
	delete(f, "Path")
	delete(f, "ForwardProtocol")
	delete(f, "ForwardHost")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRuleAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRuleAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRuleAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityRuleRequest struct {
	*tchttp.BaseRequest

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 规则名：不得超过30个字符，超过部分会被截断。
	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`

	// 安全策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 安全规则动作
	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`

	// 规则关联地址，格式需要满足CIDR网络地址规范
	SourceCidr *string `json:"SourceCidr,omitempty" name:"SourceCidr"`

	// 协议类型
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 端口范围，支持以下格式
	// 单个端口: 80
	// 多个端口: 80,443
	// 连续端口: 3306-20000
	// 所有端口: ALL
	DestPortRange *string `json:"DestPortRange,omitempty" name:"DestPortRange"`
}

func (r *ModifySecurityRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "AliasName")
	delete(f, "PolicyId")
	delete(f, "RuleAction")
	delete(f, "SourceCidr")
	delete(f, "Protocol")
	delete(f, "DestPortRange")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecurityRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTCPListenerAttributeRequest struct {
	*tchttp.BaseRequest

	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 通道组ID，ProxyId和GroupId必须设置一个，但不能同时设置。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 通道ID，ProxyId和GroupId必须设置一个，但不能同时设置。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 监听器名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器源站调度策略，支持轮询（rr），加权轮询（wrr），最小连接数（lc）。
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 源站健康检查时间间隔，单位：秒。时间间隔取值在[5，300]之间。
	DelayLoop *uint64 `json:"DelayLoop,omitempty" name:"DelayLoop"`

	// 源站健康检查响应超时时间，单位：秒。超时时间取值在[2，60]之间。超时时间应小于健康检查时间间隔DelayLoop。
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitempty" name:"ConnectTimeout"`

	// 是否开启健康检查，1开启，0关闭。
	HealthCheck *uint64 `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 源站是否开启主备模式：1开启，0关闭，DOMAIN类型源站不支持开启
	FailoverSwitch *uint64 `json:"FailoverSwitch,omitempty" name:"FailoverSwitch"`

	// 健康阈值，表示连续检查成功多少次数后认定源站健康。范围为1到10
	HealthyThreshold *uint64 `json:"HealthyThreshold,omitempty" name:"HealthyThreshold"`

	// 不健康阈值，表示连续检查失败次数后认定源站不健康。范围为1到10
	UnhealthyThreshold *uint64 `json:"UnhealthyThreshold,omitempty" name:"UnhealthyThreshold"`
}

func (r *ModifyTCPListenerAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTCPListenerAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "GroupId")
	delete(f, "ProxyId")
	delete(f, "ListenerName")
	delete(f, "Scheduler")
	delete(f, "DelayLoop")
	delete(f, "ConnectTimeout")
	delete(f, "HealthCheck")
	delete(f, "FailoverSwitch")
	delete(f, "HealthyThreshold")
	delete(f, "UnhealthyThreshold")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTCPListenerAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTCPListenerAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTCPListenerAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTCPListenerAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUDPListenerAttributeRequest struct {
	*tchttp.BaseRequest

	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 通道组ID，ProxyId和GroupId必须设置一个，但不能同时设置。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 通道ID，ProxyId和GroupId必须设置一个，但不能同时设置。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 监听器名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器源站调度策略
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`
}

func (r *ModifyUDPListenerAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUDPListenerAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "GroupId")
	delete(f, "ProxyId")
	delete(f, "ListenerName")
	delete(f, "Scheduler")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUDPListenerAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUDPListenerAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUDPListenerAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUDPListenerAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NationCountryInnerInfo struct {

	// 国家名
	NationCountryName *string `json:"NationCountryName,omitempty" name:"NationCountryName"`

	// 国家内部编码
	NationCountryInnerCode *string `json:"NationCountryInnerCode,omitempty" name:"NationCountryInnerCode"`
}

type NewRealServer struct {

	// 源站ID
	RealServerId *string `json:"RealServerId,omitempty" name:"RealServerId"`

	// 源站ip或域名
	RealServerIP *string `json:"RealServerIP,omitempty" name:"RealServerIP"`
}

type OpenProxiesRequest struct {
	*tchttp.BaseRequest

	// （旧参数，请切换到ProxyIds）通道的实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	// 更多详细信息请参阅：如何保证幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// （新参数）通道的实例ID列表。
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`
}

func (r *OpenProxiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProxiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ClientToken")
	delete(f, "ProxyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenProxiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OpenProxiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 非关闭状态下的通道实例ID列表，不可开启。
		InvalidStatusInstanceSet []*string `json:"InvalidStatusInstanceSet,omitempty" name:"InvalidStatusInstanceSet"`

		// 开启操作失败的通道实例ID列表。
		OperationFailedInstanceSet []*string `json:"OperationFailedInstanceSet,omitempty" name:"OperationFailedInstanceSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenProxiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProxiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenProxyGroupRequest struct {
	*tchttp.BaseRequest

	// 通道组实例 ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *OpenProxyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProxyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenProxyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OpenProxyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 非关闭状态下的通道实例ID列表，不可开启。
		InvalidStatusInstanceSet []*string `json:"InvalidStatusInstanceSet,omitempty" name:"InvalidStatusInstanceSet"`

		// 开启操作失败的通道实例ID列表。
		OperationFailedInstanceSet []*string `json:"OperationFailedInstanceSet,omitempty" name:"OperationFailedInstanceSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenProxyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProxyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenSecurityPolicyRequest struct {
	*tchttp.BaseRequest

	// 需开启安全策略的通道ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 安全策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *OpenSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyId")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OpenSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步流程ID，可以通过DescribeAsyncTaskStatus接口查询流程运行状态
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProxyGroupDetail struct {

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 通道组中通道数量
	ProxyNum *int64 `json:"ProxyNum,omitempty" name:"ProxyNum"`

	// 通道组状态：
	// 0表示正常运行；
	// 1表示创建中；
	// 4表示销毁中；
	// 11表示迁移中；
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 归属Uin
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 创建Uin
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`

	// 通道名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 通道组域名解析默认IP
	DnsDefaultIp *string `json:"DnsDefaultIp,omitempty" name:"DnsDefaultIp"`

	// 通道组域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 目标地域
	RealServerRegionInfo *RegionDetail `json:"RealServerRegionInfo,omitempty" name:"RealServerRegionInfo"`

	// 是否老通道组，2018-08-03之前创建的通道组为老通道组
	IsOldGroup *bool `json:"IsOldGroup,omitempty" name:"IsOldGroup"`

	// 通道组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet"`

	// 安全策略ID，当设置了安全策略时，存在该字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 通道组版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`

	// 通道获取客户端IP的方式，0表示TOA，1表示Proxy Protocol
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIPMethod []*int64 `json:"ClientIPMethod,omitempty" name:"ClientIPMethod"`

	// IP版本，可取值：IPv4、IPv6，默认值IPv4
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPAddressVersion *string `json:"IPAddressVersion,omitempty" name:"IPAddressVersion"`

	// 通道组套餐类型：Thunder表示标准通道组，Accelerator表示游戏加速器通道组，CrossBorder表示跨境通道组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// 支持Http3特性的标识，其中：
	// 0表示关闭；
	// 1表示启用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Http3Supported *int64 `json:"Http3Supported,omitempty" name:"Http3Supported"`
}

type ProxyGroupInfo struct {

	// 通道组id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 通道组域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 通道组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 目标地域
	RealServerRegionInfo *RegionDetail `json:"RealServerRegionInfo,omitempty" name:"RealServerRegionInfo"`

	// 通道组状态。
	// 其中，
	// RUNNING表示运行中；
	// CREATING表示创建中；
	// DESTROYING表示销毁中；
	// MOVING表示通道迁移中。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 标签列表。
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet"`

	// 通道组版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 通道组是否包含微软通道
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyType *uint64 `json:"ProxyType,omitempty" name:"ProxyType"`

	// 支持Http3特性的标识，其中：
	// 0，表示不支持Http3；
	// 1，表示支持Http3。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Http3Supported *int64 `json:"Http3Supported,omitempty" name:"Http3Supported"`
}

type ProxyIdDict struct {

	// 通道ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

type ProxyInfo struct {

	// （旧参数，请使用ProxyId）通道实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 创建时间，采用Unix时间戳的方式，表示从1970年1月1日（UTC/GMT的午夜）开始所经过的秒数。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 通道名称。
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// 接入地域。
	AccessRegion *string `json:"AccessRegion,omitempty" name:"AccessRegion"`

	// 源站地域。
	RealServerRegion *string `json:"RealServerRegion,omitempty" name:"RealServerRegion"`

	// 带宽，单位：Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 并发，单位：个/秒。
	Concurrent *int64 `json:"Concurrent,omitempty" name:"Concurrent"`

	// 通道状态。其中：
	// RUNNING表示运行中；
	// CREATING表示创建中；
	// DESTROYING表示销毁中；
	// OPENING表示开启中；
	// CLOSING表示关闭中；
	// CLOSED表示已关闭；
	// ADJUSTING表示配置变更中；
	// ISOLATING表示隔离中；
	// ISOLATED表示已隔离；
	// CLONING表示复制中。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 接入域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 接入IP。
	IP *string `json:"IP,omitempty" name:"IP"`

	// 通道版本号：1.0，2.0，3.0。
	Version *string `json:"Version,omitempty" name:"Version"`

	// （新参数）通道实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 1，该通道可缩扩容；0，该通道无法缩扩容。
	Scalarable *int64 `json:"Scalarable,omitempty" name:"Scalarable"`

	// 支持的协议类型。
	SupportProtocols []*string `json:"SupportProtocols,omitempty" name:"SupportProtocols"`

	// 通道组ID，当通道归属于某一通道组时，存在该字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 安全策略ID，当设置了安全策略时，存在该字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 接入地域详细信息，包括地域ID和地域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessRegionInfo *RegionDetail `json:"AccessRegionInfo,omitempty" name:"AccessRegionInfo"`

	// 源站地域详细信息，包括地域ID和地域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealServerRegionInfo *RegionDetail `json:"RealServerRegionInfo,omitempty" name:"RealServerRegionInfo"`

	// 通道转发IP
	ForwardIP *string `json:"ForwardIP,omitempty" name:"ForwardIP"`

	// 标签列表，不存在标签时，该字段为空列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet"`

	// 是否支持安全组配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportSecurity *int64 `json:"SupportSecurity,omitempty" name:"SupportSecurity"`

	// 计费类型: 0表示按带宽计费  1表示按流量计费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingType *int64 `json:"BillingType,omitempty" name:"BillingType"`

	// 关联了解析的域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelatedGlobalDomains []*string `json:"RelatedGlobalDomains,omitempty" name:"RelatedGlobalDomains"`

	// 配置变更时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyConfigTime *uint64 `json:"ModifyConfigTime,omitempty" name:"ModifyConfigTime"`

	// 通道类型，100表示THUNDER通道，103表示微软合作通道
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyType *uint64 `json:"ProxyType,omitempty" name:"ProxyType"`

	// 通道获取客户端IP的方式，0表示TOA，1表示Proxy Protocol
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIPMethod []*int64 `json:"ClientIPMethod,omitempty" name:"ClientIPMethod"`

	// IP版本：IPv4、IPv6
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPAddressVersion *string `json:"IPAddressVersion,omitempty" name:"IPAddressVersion"`

	// 网络类型：normal表示常规BGP，cn2表示精品BGP，triple表示三网
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// 通道套餐类型：Thunder表示标准通道，Accelerator表示游戏加速器通道，
	// CrossBorder表示跨境通道。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// 封禁解封状态：BANNED表示已封禁，RECOVER表示已解封或未封禁，BANNING表示封禁中，RECOVERING表示解封中，BAN_FAILED表示封禁失败，RECOVER_FAILED表示解封失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BanStatus *string `json:"BanStatus,omitempty" name:"BanStatus"`

	// IP列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPList []*IPDetail `json:"IPList,omitempty" name:"IPList"`

	// 支持Http3协议的标识，其中：
	// 0表示关闭；
	// 1表示启用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Http3Supported *int64 `json:"Http3Supported,omitempty" name:"Http3Supported"`
}

type ProxySimpleInfo struct {

	// 通道ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 通道名称
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// 监听器列表
	ListenerList []*ListenerInfo `json:"ListenerList,omitempty" name:"ListenerList"`
}

type ProxyStatus struct {

	// 通道实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 通道状态。
	// 其中：
	// RUNNING表示运行中；
	// CREATING表示创建中；
	// DESTROYING表示销毁中；
	// OPENING表示开启中；
	// CLOSING表示关闭中；
	// CLOSED表示已关闭；
	// ADJUSTING表示配置变更中；
	// ISOLATING表示隔离中；
	// ISOLATED表示已隔离。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type RealServer struct {

	// 源站的IP或域名
	RealServerIP *string `json:"RealServerIP,omitempty" name:"RealServerIP"`

	// 源站ID
	RealServerId *string `json:"RealServerId,omitempty" name:"RealServerId"`

	// 源站名称
	RealServerName *string `json:"RealServerName,omitempty" name:"RealServerName"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type RealServerBindSetReq struct {

	// 源站id
	RealServerId *string `json:"RealServerId,omitempty" name:"RealServerId"`

	// 源站端口
	RealServerPort *uint64 `json:"RealServerPort,omitempty" name:"RealServerPort"`

	// 源站IP
	RealServerIP *string `json:"RealServerIP,omitempty" name:"RealServerIP"`

	// 源站权重
	RealServerWeight *uint64 `json:"RealServerWeight,omitempty" name:"RealServerWeight"`

	// 源站主备角色：master主，slave备，该参数必须在监听器打开了源站主备模式，且监听器类型为TCP监听器
	RealServerFailoverRole *string `json:"RealServerFailoverRole,omitempty" name:"RealServerFailoverRole"`
}

type RealServerStatus struct {

	// 源站ID。
	RealServerId *string `json:"RealServerId,omitempty" name:"RealServerId"`

	// 0表示未被绑定 1表示被规则或者监听器绑定。
	BindStatus *int64 `json:"BindStatus,omitempty" name:"BindStatus"`

	// 绑定此源站的通道ID，没有绑定时为空字符串。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 绑定此源站的通道组ID，没有绑定时为空字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type RegionDetail struct {

	// 区域ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 区域英文名或中文名
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 机房所属大区
	RegionArea *string `json:"RegionArea,omitempty" name:"RegionArea"`

	// 机房所属大区名
	RegionAreaName *string `json:"RegionAreaName,omitempty" name:"RegionAreaName"`

	// 机房类型, dc表示DataCenter数据中心, ec表示EdgeComputing边缘节点
	IDCType *string `json:"IDCType,omitempty" name:"IDCType"`

	// 特性位图，每个bit位代表一种特性，其中：
	// 0，表示不支持该特性；
	// 1，表示支持该特性。
	// 特性位图含义如下（从右往左）：
	// 第1个bit，支持4层加速；
	// 第2个bit，支持7层加速；
	// 第3个bit，支持Http3接入；
	// 第4个bit，支持IPv6；
	// 第5个bit，支持精品BGP接入；
	// 第6个bit，支持三网接入；
	// 第7个bit，支持接入段Qos加速。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeatureBitmap *uint64 `json:"FeatureBitmap,omitempty" name:"FeatureBitmap"`
}

type RemoveRealServersRequest struct {
	*tchttp.BaseRequest

	// 源站Id列表
	RealServerIds []*string `json:"RealServerIds,omitempty" name:"RealServerIds"`
}

func (r *RemoveRealServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveRealServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RealServerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveRealServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RemoveRealServersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveRealServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveRealServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleCheckParams struct {

	// 健康检查的时间间隔
	DelayLoop *uint64 `json:"DelayLoop,omitempty" name:"DelayLoop"`

	// 健康检查的响应超时时间
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitempty" name:"ConnectTimeout"`

	// 健康检查的检查路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 健康检查的方法，GET/HEAD
	Method *string `json:"Method,omitempty" name:"Method"`

	// 确认源站正常的返回码，可选范围[100, 200, 300, 400, 500]
	StatusCode []*uint64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 健康检查的检查域名。
	// 当调用ModifyRuleAttribute时，不支持修改该参数。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 源站服务失败统计频率
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedCountInter *uint64 `json:"FailedCountInter,omitempty" name:"FailedCountInter"`

	// 源站健康性检查阀值，超过该阀值会屏蔽服务
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedThreshold *uint64 `json:"FailedThreshold,omitempty" name:"FailedThreshold"`

	// 源站健康性检测超出阀值后，屏蔽的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlockInter *uint64 `json:"BlockInter,omitempty" name:"BlockInter"`
}

type RuleInfo struct {

	// 规则信息
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 监听器信息
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 规则域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 规则路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 源站类型
	RealServerType *string `json:"RealServerType,omitempty" name:"RealServerType"`

	// 转发源站策略
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 是否开启健康检查标志，1表示开启，0表示关闭
	HealthCheck *uint64 `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 规则状态，0表示运行中，1表示创建中，2表示销毁中，3表示绑定解绑源站中，4表示配置更新中
	RuleStatus *uint64 `json:"RuleStatus,omitempty" name:"RuleStatus"`

	// 健康检查相关参数
	CheckParams *RuleCheckParams `json:"CheckParams,omitempty" name:"CheckParams"`

	// 已绑定的源站相关信息
	RealServerSet []*BindRealServer `json:"RealServerSet,omitempty" name:"RealServerSet"`

	// 源站的服务状态，0表示异常，1表示正常。
	// 未开启健康检查时，该状态始终未正常。
	// 只要有一个源站健康状态为异常时，该状态为异常，具体源站的状态请查看RealServerSet。
	BindStatus *uint64 `json:"BindStatus,omitempty" name:"BindStatus"`

	// 通道转发到源站的请求所携带的host，其中default表示直接转发接收到的host。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForwardHost *string `json:"ForwardHost,omitempty" name:"ForwardHost"`
}

type SecurityPolicyRuleIn struct {

	// 请求来源IP或IP段。
	SourceCidr *string `json:"SourceCidr,omitempty" name:"SourceCidr"`

	// 策略：允许（ACCEPT）或拒绝（DROP）
	Action *string `json:"Action,omitempty" name:"Action"`

	// 规则别名
	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`

	// 协议：TCP或UDP，ALL表示所有协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 目标端口，填写格式举例：
	// 单个端口: 80
	// 多个端口: 80,443
	// 连续端口: 3306-20000
	// 所有端口: ALL
	DestPortRange *string `json:"DestPortRange,omitempty" name:"DestPortRange"`
}

type SecurityPolicyRuleOut struct {

	// 策略：允许（ACCEPT）或拒绝（DROP）
	Action *string `json:"Action,omitempty" name:"Action"`

	// 请求来源Ip或Ip段
	SourceCidr *string `json:"SourceCidr,omitempty" name:"SourceCidr"`

	// 规则别名
	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`

	// 目标端口范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestPortRange *string `json:"DestPortRange,omitempty" name:"DestPortRange"`

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 要匹配的协议类型（TCP/UDP）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 安全策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

type SetAuthenticationRequest struct {
	*tchttp.BaseRequest

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 需要进行高级配置的域名，该域名为监听器下的转发规则的域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 基础认证开关，其中：
	// 0，关闭基础认证；
	// 1，开启基础认证。
	// 默认为0。
	BasicAuth *int64 `json:"BasicAuth,omitempty" name:"BasicAuth"`

	// 通道认证开关，用于源站对Gaap的认证，其中：
	// 0，关闭通道认证；
	// 1，开启通道认证。
	// 默认为0。
	GaapAuth *int64 `json:"GaapAuth,omitempty" name:"GaapAuth"`

	// 源站认证开关，用于Gaap对服务器的认证，其中：
	// 0，关闭源站认证；
	// 1，开启源站认证。
	// 默认为0。
	RealServerAuth *int64 `json:"RealServerAuth,omitempty" name:"RealServerAuth"`

	// 基础认证配置ID，从证书管理页获取。
	BasicAuthConfId *string `json:"BasicAuthConfId,omitempty" name:"BasicAuthConfId"`

	// 通道SSL证书ID，从证书管理页获取。
	GaapCertificateId *string `json:"GaapCertificateId,omitempty" name:"GaapCertificateId"`

	// 源站CA证书ID，从证书管理页获取。源站认证时，填写该参数或RealServerCertificateId参数
	RealServerCertificateId *string `json:"RealServerCertificateId,omitempty" name:"RealServerCertificateId"`

	// 源站证书域名。
	RealServerCertificateDomain *string `json:"RealServerCertificateDomain,omitempty" name:"RealServerCertificateDomain"`

	// 多源站CA证书ID，从证书管理页获取。源站认证时，填写该参数或RealServerCertificateId参数
	PolyRealServerCertificateIds []*string `json:"PolyRealServerCertificateIds,omitempty" name:"PolyRealServerCertificateIds"`
}

func (r *SetAuthenticationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAuthenticationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "Domain")
	delete(f, "BasicAuth")
	delete(f, "GaapAuth")
	delete(f, "RealServerAuth")
	delete(f, "BasicAuthConfId")
	delete(f, "GaapCertificateId")
	delete(f, "RealServerCertificateId")
	delete(f, "RealServerCertificateDomain")
	delete(f, "PolyRealServerCertificateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetAuthenticationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetAuthenticationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetAuthenticationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAuthenticationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StatisticsDataInfo struct {

	// 对应的时间点
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// 统计数据值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *float64 `json:"Data,omitempty" name:"Data"`
}

type TCPListener struct {

	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 监听器转发源站端口，仅对版本为1.0的通道有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealServerPort *uint64 `json:"RealServerPort,omitempty" name:"RealServerPort"`

	// 监听器绑定源站类型
	RealServerType *string `json:"RealServerType,omitempty" name:"RealServerType"`

	// 监听器协议， TCP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器状态，其中：
	// 0表示运行中；
	// 1表示创建中；
	// 2表示销毁中；
	// 3表示源站调整中；
	// 4表示配置变更中。
	ListenerStatus *uint64 `json:"ListenerStatus,omitempty" name:"ListenerStatus"`

	// 监听器源站访问策略，其中：
	// rr表示轮询；
	// wrr表示加权轮询；
	// lc表示最小连接数。
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 源站健康检查响应超时时间，单位：秒
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitempty" name:"ConnectTimeout"`

	// 源站健康检查时间间隔，单位：秒
	DelayLoop *uint64 `json:"DelayLoop,omitempty" name:"DelayLoop"`

	// 监听器是否开启健康检查，其中：
	// 0表示关闭；
	// 1表示开启
	HealthCheck *uint64 `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 监听器绑定的源站状态， 其中：
	// 0表示异常；
	// 1表示正常。
	BindStatus *uint64 `json:"BindStatus,omitempty" name:"BindStatus"`

	// 监听器绑定的源站信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealServerSet []*BindRealServer `json:"RealServerSet,omitempty" name:"RealServerSet"`

	// 监听器创建时间，Unix时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 监听器获取客户端 IP 的方式，0表示TOA, 1表示Proxy Protocol
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIPMethod *uint64 `json:"ClientIPMethod,omitempty" name:"ClientIPMethod"`

	// 健康阈值，表示连续检查成功多少次后认定源站健康。范围为1到10
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthyThreshold *uint64 `json:"HealthyThreshold,omitempty" name:"HealthyThreshold"`

	// 不健康阈值，表示连续检查失败多少次数后认为源站不健康。范围为1到10
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnhealthyThreshold *uint64 `json:"UnhealthyThreshold,omitempty" name:"UnhealthyThreshold"`

	// 源站是否开启主备模式：1开启，0关闭，DOMAIN类型源站不支持开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailoverSwitch *uint64 `json:"FailoverSwitch,omitempty" name:"FailoverSwitch"`

	// 是否开启会话保持选项：0关闭， 非0开启，非0值为会话保持时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionPersist *uint64 `json:"SessionPersist,omitempty" name:"SessionPersist"`
}

type TagPair struct {

	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagResourceInfo struct {

	// 资源类型，其中：
	// Proxy表示通道，
	// ProxyGroup表示通道组，
	// RealServer表示源站
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type UDPListener struct {

	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 监听器转发源站端口，仅V1版本通道或通道组监听器有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealServerPort *uint64 `json:"RealServerPort,omitempty" name:"RealServerPort"`

	// 监听器绑定源站类型
	RealServerType *string `json:"RealServerType,omitempty" name:"RealServerType"`

	// 监听器协议， UDP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器状态，其中：
	// 0表示运行中；
	// 1表示创建中；
	// 2表示销毁中；
	// 3表示源站调整中；
	// 4表示配置变更中。
	ListenerStatus *uint64 `json:"ListenerStatus,omitempty" name:"ListenerStatus"`

	// 监听器源站访问策略
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 监听器绑定源站状态， 0表示正常，1表示IP异常，2表示域名解析异常
	BindStatus *uint64 `json:"BindStatus,omitempty" name:"BindStatus"`

	// 监听器绑定的源站信息
	RealServerSet []*BindRealServer `json:"RealServerSet,omitempty" name:"RealServerSet"`

	// 监听器创建时间，Unix时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 是否开启会话保持选项：0关闭， 非0开启，非0值为会话保持时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionPersist *uint64 `json:"SessionPersist,omitempty" name:"SessionPersist"`
}
