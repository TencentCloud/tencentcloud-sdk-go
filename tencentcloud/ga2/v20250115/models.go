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

package v20250115

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AcceleratorAreas struct {
	// 加速地域。
	AccelerateRegion *string `json:"AccelerateRegion,omitnil,omitempty" name:"AccelerateRegion"`

	// 带宽。
	Bandwidth *uint64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 支持'BGP', '三网', '精品'，默认BGP。
	IspType *string `json:"IspType,omitnil,omitempty" name:"IspType"`

	// 仅支持IPv4，默认是IPv4。
	IpVersion *string `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`

	// 加速地域ID。
	AcceleratorAreaId *string `json:"AcceleratorAreaId,omitnil,omitempty" name:"AcceleratorAreaId"`

	// IP。
	IpAddress []*string `json:"IpAddress,omitnil,omitempty" name:"IpAddress"`

	// IP信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpAddressInfoSet []*IpAddressInfoSet `json:"IpAddressInfoSet,omitnil,omitempty" name:"IpAddressInfoSet"`
}

type AcceleratorRegionSet struct {
	// <p>地域中文名称。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>是否可用；0：不可用，1:可用。</p>
	IsAvailable *int64 `json:"IsAvailable,omitnil,omitempty" name:"IsAvailable"`

	// <p>地域信息。</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>地区名称。</p>
	AreaName *string `json:"AreaName,omitnil,omitempty" name:"AreaName"`

	// <p>是否中国地域。</p>
	IsChinaMainland *uint64 `json:"IsChinaMainland,omitnil,omitempty" name:"IsChinaMainland"`

	// <p>支持IspType类型。</p>
	SupportIspType []*string `json:"SupportIspType,omitnil,omitempty" name:"SupportIspType"`

	// <p>是否腾讯地域。</p>
	IsTencentRegion *uint64 `json:"IsTencentRegion,omitnil,omitempty" name:"IsTencentRegion"`
}

// Predefined struct for user
type CreateAccelerateAreasRequestParams struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 加速地域信息。
	AcceleratorAreas []*AcceleratorAreas `json:"AcceleratorAreas,omitnil,omitempty" name:"AcceleratorAreas"`
}

type CreateAccelerateAreasRequest struct {
	*tchttp.BaseRequest
	
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 加速地域信息。
	AcceleratorAreas []*AcceleratorAreas `json:"AcceleratorAreas,omitnil,omitempty" name:"AcceleratorAreas"`
}

func (r *CreateAccelerateAreasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccelerateAreasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "AcceleratorAreas")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccelerateAreasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccelerateAreasResponseParams struct {
	// 异步任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAccelerateAreasResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccelerateAreasResponseParams `json:"Response"`
}

func (r *CreateAccelerateAreasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccelerateAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEndpointGroupRequestParams struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 终端节点组类型。支持VIRTUAL，DEFAULT。
	EndpointGroupType *string `json:"EndpointGroupType,omitnil,omitempty" name:"EndpointGroupType"`

	// 终端节点组配置。
	EndpointGroupConfiguration *EndpointGroupConfiguration `json:"EndpointGroupConfiguration,omitnil,omitempty" name:"EndpointGroupConfiguration"`
}

type CreateEndpointGroupRequest struct {
	*tchttp.BaseRequest
	
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 终端节点组类型。支持VIRTUAL，DEFAULT。
	EndpointGroupType *string `json:"EndpointGroupType,omitnil,omitempty" name:"EndpointGroupType"`

	// 终端节点组配置。
	EndpointGroupConfiguration *EndpointGroupConfiguration `json:"EndpointGroupConfiguration,omitnil,omitempty" name:"EndpointGroupConfiguration"`
}

func (r *CreateEndpointGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEndpointGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "EndpointGroupType")
	delete(f, "EndpointGroupConfiguration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEndpointGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEndpointGroupResponseParams struct {
	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 终端节点组实例ID。
	EndpointGroupId *string `json:"EndpointGroupId,omitnil,omitempty" name:"EndpointGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEndpointGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateEndpointGroupResponseParams `json:"Response"`
}

func (r *CreateEndpointGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEndpointGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateForwardingRuleRequestParams struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 策略ID。
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// 七层转发规则条件信息。
	RuleConditions []*RuleCondition `json:"RuleConditions,omitnil,omitempty" name:"RuleConditions"`

	// 七层转发规则行为信息。
	RuleActions []*RuleAction `json:"RuleActions,omitnil,omitempty" name:"RuleActions"`

	// 回源Header信息。
	OriginHeaders []*OriginHeader `json:"OriginHeaders,omitnil,omitempty" name:"OriginHeaders"`

	// 是否开启回源sni。
	EnableOriginSni *bool `json:"EnableOriginSni,omitnil,omitempty" name:"EnableOriginSni"`

	// 回源sni。
	OriginSni *string `json:"OriginSni,omitnil,omitempty" name:"OriginSni"`

	// 回源host。
	OriginHost *string `json:"OriginHost,omitnil,omitempty" name:"OriginHost"`
}

type CreateForwardingRuleRequest struct {
	*tchttp.BaseRequest
	
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 策略ID。
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// 七层转发规则条件信息。
	RuleConditions []*RuleCondition `json:"RuleConditions,omitnil,omitempty" name:"RuleConditions"`

	// 七层转发规则行为信息。
	RuleActions []*RuleAction `json:"RuleActions,omitnil,omitempty" name:"RuleActions"`

	// 回源Header信息。
	OriginHeaders []*OriginHeader `json:"OriginHeaders,omitnil,omitempty" name:"OriginHeaders"`

	// 是否开启回源sni。
	EnableOriginSni *bool `json:"EnableOriginSni,omitnil,omitempty" name:"EnableOriginSni"`

	// 回源sni。
	OriginSni *string `json:"OriginSni,omitnil,omitempty" name:"OriginSni"`

	// 回源host。
	OriginHost *string `json:"OriginHost,omitnil,omitempty" name:"OriginHost"`
}

func (r *CreateForwardingRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateForwardingRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "ForwardingPolicyId")
	delete(f, "RuleConditions")
	delete(f, "RuleActions")
	delete(f, "OriginHeaders")
	delete(f, "EnableOriginSni")
	delete(f, "OriginSni")
	delete(f, "OriginHost")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateForwardingRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateForwardingRuleResponseParams struct {
	// 异步任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 七层转发规则ID。
	ForwardingRuleId *string `json:"ForwardingRuleId,omitnil,omitempty" name:"ForwardingRuleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateForwardingRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateForwardingRuleResponseParams `json:"Response"`
}

func (r *CreateForwardingRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateForwardingRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGlobalAcceleratorRequestParams struct {
	// <p>名称，最大长度不能超过60个字节。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>计费模式，PREPAID：表示预付费，即包年包月，POSTPAID：表示后付费，即按量计费。默认：POSTPAID。当前仅支持后付费。</p>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// <p>描述信息，最大长度不能超过100个字节。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>跨境类型；HighQuality：精品BGP-IP跨境；Unicom：联通专线跨境。</p>
	CrossBorderType *string `json:"CrossBorderType,omitnil,omitempty" name:"CrossBorderType"`

	// <p>此Flag代表签署跨境服务承诺书。当使用跨境服务时候，此字段必传。True：代表签署。</p>
	CrossBorderPromiseFlag *bool `json:"CrossBorderPromiseFlag,omitnil,omitempty" name:"CrossBorderPromiseFlag"`

	// <p>标签信息</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateGlobalAcceleratorRequest struct {
	*tchttp.BaseRequest
	
	// <p>名称，最大长度不能超过60个字节。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>计费模式，PREPAID：表示预付费，即包年包月，POSTPAID：表示后付费，即按量计费。默认：POSTPAID。当前仅支持后付费。</p>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// <p>描述信息，最大长度不能超过100个字节。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>跨境类型；HighQuality：精品BGP-IP跨境；Unicom：联通专线跨境。</p>
	CrossBorderType *string `json:"CrossBorderType,omitnil,omitempty" name:"CrossBorderType"`

	// <p>此Flag代表签署跨境服务承诺书。当使用跨境服务时候，此字段必传。True：代表签署。</p>
	CrossBorderPromiseFlag *bool `json:"CrossBorderPromiseFlag,omitnil,omitempty" name:"CrossBorderPromiseFlag"`

	// <p>标签信息</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateGlobalAcceleratorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlobalAcceleratorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "InstanceChargeType")
	delete(f, "Description")
	delete(f, "CrossBorderType")
	delete(f, "CrossBorderPromiseFlag")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGlobalAcceleratorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGlobalAcceleratorResponseParams struct {
	// <p>任务ID。</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>全球加速实例ID。</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGlobalAcceleratorResponse struct {
	*tchttp.BaseResponse
	Response *CreateGlobalAcceleratorResponseParams `json:"Response"`
}

func (r *CreateGlobalAcceleratorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlobalAcceleratorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateListenerRequestParams struct {
	// <p>全球加速实例ID。</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>名称，最大长度不能超过60个字节。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>端口范围。</p>
	PortRanges *PortRanges `json:"PortRanges,omitnil,omitempty" name:"PortRanges"`

	// <p>描述信息，最大长度不能超过100个字节。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>监听类型，默认为智能路由。</p>
	ListenerType *string `json:"ListenerType,omitnil,omitempty" name:"ListenerType"`

	// <p>协议，默认为TCP。</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>连接空闲等待时间。</p>
	IdleTimeout *uint64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`

	// <p>四层获取源IP方式，支持&#39;TOA&#39;, &#39;ProxyProtocol&#39;。</p>
	GetRealIpType *string `json:"GetRealIpType,omitnil,omitempty" name:"GetRealIpType"`

	// <p>是否开启会话保持。</p>
	ClientAffinity *string `json:"ClientAffinity,omitnil,omitempty" name:"ClientAffinity"`

	// <p>请求超时时间。</p>
	RequestTimeout *uint64 `json:"RequestTimeout,omitnil,omitempty" name:"RequestTimeout"`

	// <p>是否打开七层获取源IP方式。</p>
	XForwardedForRealIp *bool `json:"XForwardedForRealIp,omitnil,omitempty" name:"XForwardedForRealIp"`

	// <p>解析方式。UNIDIRECTIONAL：双向。MUTUAL：单向。</p>
	CertificationType *string `json:"CertificationType,omitnil,omitempty" name:"CertificationType"`

	// <p>加密算法套件。</p>
	CipherPolicyId *string `json:"CipherPolicyId,omitnil,omitempty" name:"CipherPolicyId"`

	// <p>服务器证书。</p>
	ServerCertificates []*string `json:"ServerCertificates,omitnil,omitempty" name:"ServerCertificates"`

	// <p>客户端证书。</p>
	ClientCaCertificates []*string `json:"ClientCaCertificates,omitnil,omitempty" name:"ClientCaCertificates"`
}

type CreateListenerRequest struct {
	*tchttp.BaseRequest
	
	// <p>全球加速实例ID。</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>名称，最大长度不能超过60个字节。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>端口范围。</p>
	PortRanges *PortRanges `json:"PortRanges,omitnil,omitempty" name:"PortRanges"`

	// <p>描述信息，最大长度不能超过100个字节。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>监听类型，默认为智能路由。</p>
	ListenerType *string `json:"ListenerType,omitnil,omitempty" name:"ListenerType"`

	// <p>协议，默认为TCP。</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>连接空闲等待时间。</p>
	IdleTimeout *uint64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`

	// <p>四层获取源IP方式，支持&#39;TOA&#39;, &#39;ProxyProtocol&#39;。</p>
	GetRealIpType *string `json:"GetRealIpType,omitnil,omitempty" name:"GetRealIpType"`

	// <p>是否开启会话保持。</p>
	ClientAffinity *string `json:"ClientAffinity,omitnil,omitempty" name:"ClientAffinity"`

	// <p>请求超时时间。</p>
	RequestTimeout *uint64 `json:"RequestTimeout,omitnil,omitempty" name:"RequestTimeout"`

	// <p>是否打开七层获取源IP方式。</p>
	XForwardedForRealIp *bool `json:"XForwardedForRealIp,omitnil,omitempty" name:"XForwardedForRealIp"`

	// <p>解析方式。UNIDIRECTIONAL：双向。MUTUAL：单向。</p>
	CertificationType *string `json:"CertificationType,omitnil,omitempty" name:"CertificationType"`

	// <p>加密算法套件。</p>
	CipherPolicyId *string `json:"CipherPolicyId,omitnil,omitempty" name:"CipherPolicyId"`

	// <p>服务器证书。</p>
	ServerCertificates []*string `json:"ServerCertificates,omitnil,omitempty" name:"ServerCertificates"`

	// <p>客户端证书。</p>
	ClientCaCertificates []*string `json:"ClientCaCertificates,omitnil,omitempty" name:"ClientCaCertificates"`
}

func (r *CreateListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "Name")
	delete(f, "PortRanges")
	delete(f, "Description")
	delete(f, "ListenerType")
	delete(f, "Protocol")
	delete(f, "IdleTimeout")
	delete(f, "GetRealIpType")
	delete(f, "ClientAffinity")
	delete(f, "RequestTimeout")
	delete(f, "XForwardedForRealIp")
	delete(f, "CertificationType")
	delete(f, "CipherPolicyId")
	delete(f, "ServerCertificates")
	delete(f, "ClientCaCertificates")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateListenerResponseParams struct {
	// <p>任务ID。</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>监听器ID。</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateListenerResponse struct {
	*tchttp.BaseResponse
	Response *CreateListenerResponseParams `json:"Response"`
}

func (r *CreateListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccelerateAreasRequestParams struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 加速地域ID。
	AcceleratorAreaIds []*string `json:"AcceleratorAreaIds,omitnil,omitempty" name:"AcceleratorAreaIds"`
}

type DeleteAccelerateAreasRequest struct {
	*tchttp.BaseRequest
	
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 加速地域ID。
	AcceleratorAreaIds []*string `json:"AcceleratorAreaIds,omitnil,omitempty" name:"AcceleratorAreaIds"`
}

func (r *DeleteAccelerateAreasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccelerateAreasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "AcceleratorAreaIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccelerateAreasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccelerateAreasResponseParams struct {
	// 异步任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAccelerateAreasResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccelerateAreasResponseParams `json:"Response"`
}

func (r *DeleteAccelerateAreasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccelerateAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEndpointGroupsRequestParams struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 终端节点组ID。
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitnil,omitempty" name:"EndpointGroupIds"`
}

type DeleteEndpointGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 终端节点组ID。
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitnil,omitempty" name:"EndpointGroupIds"`
}

func (r *DeleteEndpointGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEndpointGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "EndpointGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEndpointGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEndpointGroupsResponseParams struct {
	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEndpointGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEndpointGroupsResponseParams `json:"Response"`
}

func (r *DeleteEndpointGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEndpointGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteForwardingRuleRequestParams struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 策略ID。
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// 七层转发规则ID。
	ForwardingRuleId *string `json:"ForwardingRuleId,omitnil,omitempty" name:"ForwardingRuleId"`
}

type DeleteForwardingRuleRequest struct {
	*tchttp.BaseRequest
	
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 策略ID。
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// 七层转发规则ID。
	ForwardingRuleId *string `json:"ForwardingRuleId,omitnil,omitempty" name:"ForwardingRuleId"`
}

func (r *DeleteForwardingRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteForwardingRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "ForwardingPolicyId")
	delete(f, "ForwardingRuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteForwardingRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteForwardingRuleResponseParams struct {
	// 异步任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteForwardingRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteForwardingRuleResponseParams `json:"Response"`
}

func (r *DeleteForwardingRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteForwardingRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGlobalAcceleratorRequestParams struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`
}

type DeleteGlobalAcceleratorRequest struct {
	*tchttp.BaseRequest
	
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`
}

func (r *DeleteGlobalAcceleratorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGlobalAcceleratorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGlobalAcceleratorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGlobalAcceleratorResponseParams struct {
	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGlobalAcceleratorResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGlobalAcceleratorResponseParams `json:"Response"`
}

func (r *DeleteGlobalAcceleratorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGlobalAcceleratorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteListenerRequestParams struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`
}

type DeleteListenerRequest struct {
	*tchttp.BaseRequest
	
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`
}

func (r *DeleteListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteListenerResponseParams struct {
	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteListenerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteListenerResponseParams `json:"Response"`
}

func (r *DeleteListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccelerateAreasRequestParams struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 偏移量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 符合条件实例数量。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAccelerateAreasRequest struct {
	*tchttp.BaseRequest
	
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 偏移量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 符合条件实例数量。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAccelerateAreasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccelerateAreasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccelerateAreasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccelerateAreasResponseParams struct {
	// 加速地域信息。
	AccelerateAreaSet []*AcceleratorAreas `json:"AccelerateAreaSet,omitnil,omitempty" name:"AccelerateAreaSet"`

	// 实例个数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccelerateAreasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccelerateAreasResponseParams `json:"Response"`
}

func (r *DescribeAccelerateAreasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccelerateAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccelerateRegionsRequestParams struct {

}

type DescribeAccelerateRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAccelerateRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccelerateRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccelerateRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccelerateRegionsResponseParams struct {
	// 加速地域信息。
	AcceleratorRegionSet []*AcceleratorRegionSet `json:"AcceleratorRegionSet,omitnil,omitempty" name:"AcceleratorRegionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccelerateRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccelerateRegionsResponseParams `json:"Response"`
}

func (r *DescribeAccelerateRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccelerateRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCrossBorderSettlementRequestParams struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 加速地域。
	AccelerateRegion *string `json:"AccelerateRegion,omitnil,omitempty" name:"AccelerateRegion"`

	// 终端节点组地域。
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitnil,omitempty" name:"EndpointGroupRegion"`

	// 账单年月时间。
	SettlementMonth *uint64 `json:"SettlementMonth,omitnil,omitempty" name:"SettlementMonth"`
}

type DescribeCrossBorderSettlementRequest struct {
	*tchttp.BaseRequest
	
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 加速地域。
	AccelerateRegion *string `json:"AccelerateRegion,omitnil,omitempty" name:"AccelerateRegion"`

	// 终端节点组地域。
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitnil,omitempty" name:"EndpointGroupRegion"`

	// 账单年月时间。
	SettlementMonth *uint64 `json:"SettlementMonth,omitnil,omitempty" name:"SettlementMonth"`
}

func (r *DescribeCrossBorderSettlementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossBorderSettlementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "AccelerateRegion")
	delete(f, "EndpointGroupRegion")
	delete(f, "SettlementMonth")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCrossBorderSettlementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCrossBorderSettlementResponseParams struct {
	// 流量用量，单位是GB；精度为保留小数点6位。
	Traffic *float64 `json:"Traffic,omitnil,omitempty" name:"Traffic"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCrossBorderSettlementResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCrossBorderSettlementResponseParams `json:"Response"`
}

func (r *DescribeCrossBorderSettlementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossBorderSettlementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEndpointGroupsRequestParams struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件。  endpoint-group-id- String -（过滤条件）终端节点组实例ID。endpoint-group-type- String -（过滤条件）终端节点组实例类型。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeEndpointGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件。  endpoint-group-id- String -（过滤条件）终端节点组实例ID。endpoint-group-type- String -（过滤条件）终端节点组实例类型。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeEndpointGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEndpointGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEndpointGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEndpointGroupsResponseParams struct {
	// 符合条件的终端节点组。
	EndpointGroupConfigurationSet []*EndpointGroupConfigurationSet `json:"EndpointGroupConfigurationSet,omitnil,omitempty" name:"EndpointGroupConfigurationSet"`

	// 符合条件的实例个数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEndpointGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEndpointGroupsResponseParams `json:"Response"`
}

func (r *DescribeEndpointGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEndpointGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeForwardingRuleRequestParams struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 七层转发规则ID。
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeForwardingRuleRequest struct {
	*tchttp.BaseRequest
	
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 七层转发规则ID。
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeForwardingRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeForwardingRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "ForwardingPolicyId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeForwardingRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeForwardingRuleResponseParams struct {
	// 符合条件的规则信息。
	ForwardingRuleSet []*ForwardingRuleSet `json:"ForwardingRuleSet,omitnil,omitempty" name:"ForwardingRuleSet"`

	// 符合条件的实例个数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeForwardingRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeForwardingRuleResponseParams `json:"Response"`
}

func (r *DescribeForwardingRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeForwardingRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalAcceleratorsRequestParams struct {
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件。<li>global-accelerator-id - String -（过滤条件）全球加速实例ID。</li> <li>global-accelerator-state - String -（过滤条件）全球加速实例状态。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeGlobalAcceleratorsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件。<li>global-accelerator-id - String -（过滤条件）全球加速实例ID。</li> <li>global-accelerator-state - String -（过滤条件）全球加速实例状态。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeGlobalAcceleratorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalAcceleratorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGlobalAcceleratorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalAcceleratorsResponseParams struct {
	// 符合条件的全球加速实例。
	GlobalAcceleratorSet []*GlobalAcceleratorSet `json:"GlobalAcceleratorSet,omitnil,omitempty" name:"GlobalAcceleratorSet"`

	// 符合条件的实例个数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGlobalAcceleratorsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGlobalAcceleratorsResponseParams `json:"Response"`
}

func (r *DescribeGlobalAcceleratorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalAcceleratorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenersRequestParams struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件。  listener-id- String -（过滤条件）监听器实例ID。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeListenersRequest struct {
	*tchttp.BaseRequest
	
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件。  listener-id- String -（过滤条件）监听器实例ID。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenersResponseParams struct {
	// 符合条件的监听器实例。
	ListenerSet []*ListenerSet `json:"ListenerSet,omitnil,omitempty" name:"ListenerSet"`

	// 符合条件的实例个数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeListenersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListenersResponseParams `json:"Response"`
}

func (r *DescribeListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EndpointConfigurations struct {
	// 域名类型。可选值'Domain', 'PublicIp'。
	EndpointType *string `json:"EndpointType,omitnil,omitempty" name:"EndpointType"`

	// 域名。
	EndpointService *string `json:"EndpointService,omitnil,omitempty" name:"EndpointService"`

	// 权重。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 健康检查状态；HEALTH：健康；UNHEALTH：不健康。
	HealthCheckStatus *string `json:"HealthCheckStatus,omitnil,omitempty" name:"HealthCheckStatus"`
}

type EndpointGroupConfiguration struct {
	// <p>名称，最大长度不能超过60个字节。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>地域。</p>
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitnil,omitempty" name:"EndpointGroupRegion"`

	// <p>终端节点配置。</p>
	EndpointConfigurations []*EndpointConfigurations `json:"EndpointConfigurations,omitnil,omitempty" name:"EndpointConfigurations"`

	// <p>检查协议。支持&#39;TCP&#39;, &#39;HTTP&#39;, &#39;HTTPS&#39;, &#39;PING&#39;, &#39;CUSTOM&#39;。</p>
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// <p>描述信息，最大长度不能超过100个字节。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>检查端口。</p>
	CheckPort *string `json:"CheckPort,omitnil,omitempty" name:"CheckPort"`

	// <p>检查内容。</p>
	ContextType *string `json:"ContextType,omitnil,omitempty" name:"ContextType"`

	// <p>检查请求。</p>
	CheckSendContext *string `json:"CheckSendContext,omitnil,omitempty" name:"CheckSendContext"`

	// <p>检查返回结果。</p>
	CheckRecvContext *string `json:"CheckRecvContext,omitnil,omitempty" name:"CheckRecvContext"`

	// <p>是否开启健康检查。</p>
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// <p>响应超时时间。</p>
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitnil,omitempty" name:"ConnectTimeout"`

	// <p>健康检查间隔。</p>
	HealthCheckInterval *uint64 `json:"HealthCheckInterval,omitnil,omitempty" name:"HealthCheckInterval"`

	// <p>不健康阀值。</p>
	UnhealthyThreshold *uint64 `json:"UnhealthyThreshold,omitnil,omitempty" name:"UnhealthyThreshold"`

	// <p>健康阈值。</p>
	HealthyThreshold *uint64 `json:"HealthyThreshold,omitnil,omitempty" name:"HealthyThreshold"`

	// <p>回源协议。</p>
	ForwardProtocol *string `json:"ForwardProtocol,omitnil,omitempty" name:"ForwardProtocol"`

	// <p>检查域名。</p>
	CheckDomain *string `json:"CheckDomain,omitnil,omitempty" name:"CheckDomain"`

	// <p>检查URL。</p>
	CheckPath *string `json:"CheckPath,omitnil,omitempty" name:"CheckPath"`

	// <p>请求方式。</p>
	CheckMethod *string `json:"CheckMethod,omitnil,omitempty" name:"CheckMethod"`

	// <p>状态检测码。</p>
	StatusMask []*string `json:"StatusMask,omitnil,omitempty" name:"StatusMask"`

	// <p>端口映射。</p>
	PortOverrides []*PortOverride `json:"PortOverrides,omitnil,omitempty" name:"PortOverrides"`

	// <p>运用商类型。</p>
	IspType *string `json:"IspType,omitnil,omitempty" name:"IspType"`

	// <p>HPPTS加密算法套件</p>
	CipherPolicyId *string `json:"CipherPolicyId,omitnil,omitempty" name:"CipherPolicyId"`
}

type EndpointGroupConfigurationSet struct {
	// <p>全球加速实例ID。</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>监听器实例ID。</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>终端节点组ID。</p>
	EndpointGroupId *string `json:"EndpointGroupId,omitnil,omitempty" name:"EndpointGroupId"`

	// <p>名称。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>地域。</p>
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitnil,omitempty" name:"EndpointGroupRegion"`

	// <p>描述。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>终端节点信息。</p>
	EndpointConfigurations []*EndpointConfigurations `json:"EndpointConfigurations,omitnil,omitempty" name:"EndpointConfigurations"`

	// <p>是否开启健康检查。</p>
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// <p>响应超时时间。</p>
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitnil,omitempty" name:"ConnectTimeout"`

	// <p>健康检查间隔。</p>
	HealthCheckInterval *uint64 `json:"HealthCheckInterval,omitnil,omitempty" name:"HealthCheckInterval"`

	// <p>不健康阈值。</p>
	UnhealthyThreshold *uint64 `json:"UnhealthyThreshold,omitnil,omitempty" name:"UnhealthyThreshold"`

	// <p>健康阈值。</p>
	HealthyThreshold *uint64 `json:"HealthyThreshold,omitnil,omitempty" name:"HealthyThreshold"`

	// <p>检查协议。</p>
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// <p>检查端口。</p>
	CheckPort *uint64 `json:"CheckPort,omitnil,omitempty" name:"CheckPort"`

	// <p>检查内容。</p>
	ContextType *string `json:"ContextType,omitnil,omitempty" name:"ContextType"`

	// <p>检查请求。</p>
	CheckSendContext *string `json:"CheckSendContext,omitnil,omitempty" name:"CheckSendContext"`

	// <p>检查返回结果。</p>
	CheckRecvContext *string `json:"CheckRecvContext,omitnil,omitempty" name:"CheckRecvContext"`

	// <p>检查域名。</p>
	CheckDomain *string `json:"CheckDomain,omitnil,omitempty" name:"CheckDomain"`

	// <p>检查URL。</p>
	CheckPath *string `json:"CheckPath,omitnil,omitempty" name:"CheckPath"`

	// <p>请求方式。</p>
	CheckMethod *string `json:"CheckMethod,omitnil,omitempty" name:"CheckMethod"`

	// <p>状态检测码。</p>
	StatusMask []*string `json:"StatusMask,omitnil,omitempty" name:"StatusMask"`

	// <p>终端节点组类型。</p>
	EndpointGroupType *string `json:"EndpointGroupType,omitnil,omitempty" name:"EndpointGroupType"`

	// <p>回源协议。</p>
	ForwardProtocol *string `json:"ForwardProtocol,omitnil,omitempty" name:"ForwardProtocol"`

	// <p>端口映射信息。</p>
	PortOverrides []*PortOverride `json:"PortOverrides,omitnil,omitempty" name:"PortOverrides"`

	// <p>自定义终端节点组是否绑定七层转发规则。</p>
	VirtualExistForwardingRuleFlag *bool `json:"VirtualExistForwardingRuleFlag,omitnil,omitempty" name:"VirtualExistForwardingRuleFlag"`

	// <p>出终端节点组公网IP。</p>
	OriginPublicIps []*string `json:"OriginPublicIps,omitnil,omitempty" name:"OriginPublicIps"`

	// <p>运营商类型；中国移动(CMCC)，中国联通(CUCC)，中国电信(CTCC)。</p>
	IspType *string `json:"IspType,omitnil,omitempty" name:"IspType"`

	// <p>HPPTS加密算法套件</p>
	CipherPolicyId *string `json:"CipherPolicyId,omitnil,omitempty" name:"CipherPolicyId"`
}

type Filter struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。当值类型为布尔类型时，可直接取值为字符串"TRUE"或 "FALSE"。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type ForwardingRuleSet struct {
	// 七层转发规则条件信息。
	RuleCondition []*RuleCondition `json:"RuleCondition,omitnil,omitempty" name:"RuleCondition"`

	// 七层转发规则行为信息。
	RuleAction []*RuleAction `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// 是否开启回源Sni。
	EnableOriginSni *bool `json:"EnableOriginSni,omitnil,omitempty" name:"EnableOriginSni"`

	// 回源Sni。
	OriginSni *string `json:"OriginSni,omitnil,omitempty" name:"OriginSni"`

	// 回源Herder信息。
	OriginHeaders []*OriginHeader `json:"OriginHeaders,omitnil,omitempty" name:"OriginHeaders"`

	// 回源Host。
	OriginHost *string `json:"OriginHost,omitnil,omitempty" name:"OriginHost"`

	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 七层转发策略ID。
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// 七层转发规则ID。
	ForwardingRuleId *string `json:"ForwardingRuleId,omitnil,omitempty" name:"ForwardingRuleId"`
}

type GlobalAcceleratorSet struct {
	// <p>全球加速实例ID。</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>全球加速实例名称。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>全球加速实例描述。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>全球加速实例创建时间。</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>全球加速实例状态。</p>
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// <p>全球加速实例付费类型。</p>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// <p>全球加速实例DdosId。</p>
	DdosId *string `json:"DdosId,omitnil,omitempty" name:"DdosId"`

	// <p>所属加速实例监听器个数。</p>
	ListenerCounts *uint64 `json:"ListenerCounts,omitnil,omitempty" name:"ListenerCounts"`

	// <p>所属加速实例加速地域个数。</p>
	AcceleratorAreaCounts *uint64 `json:"AcceleratorAreaCounts,omitnil,omitempty" name:"AcceleratorAreaCounts"`

	// <p>全球加速实例状态。</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>域名。</p>
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// <p>跨境类型；HighQuality（精品跨境）、Unicom（联通跨境）、NotAvailable（未开通）。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrossBorderType *string `json:"CrossBorderType,omitnil,omitempty" name:"CrossBorderType"`

	// <p>标签信息。</p>
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`
}

type IpAddressInfoSet struct {
	// IP地址。
	IpAddress *string `json:"IpAddress,omitnil,omitempty" name:"IpAddress"`

	// IP类型。
	IspType *string `json:"IspType,omitnil,omitempty" name:"IspType"`
}

type ListenerSet struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 监听器名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 监听器描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 协议。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 端口范围。
	PortRanges *PortRanges `json:"PortRanges,omitnil,omitempty" name:"PortRanges"`

	// 是否打开七层获取源IP方式。
	XForwardedForRealIp *bool `json:"XForwardedForRealIp,omitnil,omitempty" name:"XForwardedForRealIp"`

	// 开启会话保持。
	ClientAffinity *string `json:"ClientAffinity,omitnil,omitempty" name:"ClientAffinity"`

	// 会话保持时间。
	ClientAffinityTime *uint64 `json:"ClientAffinityTime,omitnil,omitempty" name:"ClientAffinityTime"`

	// SSL解析方式。
	CertificationType *string `json:"CertificationType,omitnil,omitempty" name:"CertificationType"`

	// 服务器证书。
	ServerCertificates []*string `json:"ServerCertificates,omitnil,omitempty" name:"ServerCertificates"`

	// 客户端证书。
	ClientCaCertificates []*string `json:"ClientCaCertificates,omitnil,omitempty" name:"ClientCaCertificates"`

	// TLS密码套件包。
	CipherPolicyId *string `json:"CipherPolicyId,omitnil,omitempty" name:"CipherPolicyId"`

	// HTTP版本。
	HttpVersion *string `json:"HttpVersion,omitnil,omitempty" name:"HttpVersion"`

	// 请求超时时间。
	RequestTimeout *uint64 `json:"RequestTimeout,omitnil,omitempty" name:"RequestTimeout"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 监听路由类型。
	ListenerType *string `json:"ListenerType,omitnil,omitempty" name:"ListenerType"`

	// 监听器状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 所属监听器终端节点组个数。
	EndpointGroupCounts *uint64 `json:"EndpointGroupCounts,omitnil,omitempty" name:"EndpointGroupCounts"`

	// 四层获取源IP方式。
	GetRealIpType *string `json:"GetRealIpType,omitnil,omitempty" name:"GetRealIpType"`

	// 连接超时时间。
	IdleTimeout *uint64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`
}

// Predefined struct for user
type ModifyAccelerateAreasRequestParams struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 加速地域信息。
	AcceleratorAreas []*AcceleratorAreas `json:"AcceleratorAreas,omitnil,omitempty" name:"AcceleratorAreas"`
}

type ModifyAccelerateAreasRequest struct {
	*tchttp.BaseRequest
	
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 加速地域信息。
	AcceleratorAreas []*AcceleratorAreas `json:"AcceleratorAreas,omitnil,omitempty" name:"AcceleratorAreas"`
}

func (r *ModifyAccelerateAreasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccelerateAreasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "AcceleratorAreas")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccelerateAreasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccelerateAreasResponseParams struct {
	// 异步任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccelerateAreasResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccelerateAreasResponseParams `json:"Response"`
}

func (r *ModifyAccelerateAreasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccelerateAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEndpointGroupRequestParams struct {
	// <p>全球加速实例ID。</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>监听器ID。</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>终端节点组ID。</p>
	EndpointGroupId *string `json:"EndpointGroupId,omitnil,omitempty" name:"EndpointGroupId"`

	// <p>终端节点配置。</p>
	EndpointConfigurations []*EndpointConfigurations `json:"EndpointConfigurations,omitnil,omitempty" name:"EndpointConfigurations"`

	// <p>名称，最大长度不能超过60个字节。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>描述信息，最大长度不能超过100个字节。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>是否开启健康检查。</p>
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// <p>响应超时时间。</p>
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitnil,omitempty" name:"ConnectTimeout"`

	// <p>健康检查间隔。</p>
	HealthCheckInterval *uint64 `json:"HealthCheckInterval,omitnil,omitempty" name:"HealthCheckInterval"`

	// <p>不健康阀值。</p>
	UnhealthyThreshold *uint64 `json:"UnhealthyThreshold,omitnil,omitempty" name:"UnhealthyThreshold"`

	// <p>健康阀值。</p>
	HealthyThreshold *uint64 `json:"HealthyThreshold,omitnil,omitempty" name:"HealthyThreshold"`

	// <p>检查协议。</p>
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// <p>检查端口。</p>
	CheckPort *uint64 `json:"CheckPort,omitnil,omitempty" name:"CheckPort"`

	// <p>检查内容。</p>
	ContextType *string `json:"ContextType,omitnil,omitempty" name:"ContextType"`

	// <p>检查请求。</p>
	CheckSendContext *string `json:"CheckSendContext,omitnil,omitempty" name:"CheckSendContext"`

	// <p>检查返回结果。</p>
	CheckRecvContext *string `json:"CheckRecvContext,omitnil,omitempty" name:"CheckRecvContext"`

	// <p>检查域名。</p>
	CheckDomain *string `json:"CheckDomain,omitnil,omitempty" name:"CheckDomain"`

	// <p>检查URL。</p>
	CheckPath *string `json:"CheckPath,omitnil,omitempty" name:"CheckPath"`

	// <p>请求方式。</p>
	CheckMethod *string `json:"CheckMethod,omitnil,omitempty" name:"CheckMethod"`

	// <p>状态检测码。</p>
	StatusMask []*string `json:"StatusMask,omitnil,omitempty" name:"StatusMask"`

	// <p>回源协议。</p>
	ForwardProtocol *string `json:"ForwardProtocol,omitnil,omitempty" name:"ForwardProtocol"`

	// <p>端口映射。</p>
	PortOverrides []*PortOverride `json:"PortOverrides,omitnil,omitempty" name:"PortOverrides"`

	// <p>HPPTS加密算法套件</p>
	CipherPolicyId *string `json:"CipherPolicyId,omitnil,omitempty" name:"CipherPolicyId"`
}

type ModifyEndpointGroupRequest struct {
	*tchttp.BaseRequest
	
	// <p>全球加速实例ID。</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>监听器ID。</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>终端节点组ID。</p>
	EndpointGroupId *string `json:"EndpointGroupId,omitnil,omitempty" name:"EndpointGroupId"`

	// <p>终端节点配置。</p>
	EndpointConfigurations []*EndpointConfigurations `json:"EndpointConfigurations,omitnil,omitempty" name:"EndpointConfigurations"`

	// <p>名称，最大长度不能超过60个字节。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>描述信息，最大长度不能超过100个字节。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>是否开启健康检查。</p>
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// <p>响应超时时间。</p>
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitnil,omitempty" name:"ConnectTimeout"`

	// <p>健康检查间隔。</p>
	HealthCheckInterval *uint64 `json:"HealthCheckInterval,omitnil,omitempty" name:"HealthCheckInterval"`

	// <p>不健康阀值。</p>
	UnhealthyThreshold *uint64 `json:"UnhealthyThreshold,omitnil,omitempty" name:"UnhealthyThreshold"`

	// <p>健康阀值。</p>
	HealthyThreshold *uint64 `json:"HealthyThreshold,omitnil,omitempty" name:"HealthyThreshold"`

	// <p>检查协议。</p>
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// <p>检查端口。</p>
	CheckPort *uint64 `json:"CheckPort,omitnil,omitempty" name:"CheckPort"`

	// <p>检查内容。</p>
	ContextType *string `json:"ContextType,omitnil,omitempty" name:"ContextType"`

	// <p>检查请求。</p>
	CheckSendContext *string `json:"CheckSendContext,omitnil,omitempty" name:"CheckSendContext"`

	// <p>检查返回结果。</p>
	CheckRecvContext *string `json:"CheckRecvContext,omitnil,omitempty" name:"CheckRecvContext"`

	// <p>检查域名。</p>
	CheckDomain *string `json:"CheckDomain,omitnil,omitempty" name:"CheckDomain"`

	// <p>检查URL。</p>
	CheckPath *string `json:"CheckPath,omitnil,omitempty" name:"CheckPath"`

	// <p>请求方式。</p>
	CheckMethod *string `json:"CheckMethod,omitnil,omitempty" name:"CheckMethod"`

	// <p>状态检测码。</p>
	StatusMask []*string `json:"StatusMask,omitnil,omitempty" name:"StatusMask"`

	// <p>回源协议。</p>
	ForwardProtocol *string `json:"ForwardProtocol,omitnil,omitempty" name:"ForwardProtocol"`

	// <p>端口映射。</p>
	PortOverrides []*PortOverride `json:"PortOverrides,omitnil,omitempty" name:"PortOverrides"`

	// <p>HPPTS加密算法套件</p>
	CipherPolicyId *string `json:"CipherPolicyId,omitnil,omitempty" name:"CipherPolicyId"`
}

func (r *ModifyEndpointGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEndpointGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "EndpointGroupId")
	delete(f, "EndpointConfigurations")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "EnableHealthCheck")
	delete(f, "ConnectTimeout")
	delete(f, "HealthCheckInterval")
	delete(f, "UnhealthyThreshold")
	delete(f, "HealthyThreshold")
	delete(f, "CheckType")
	delete(f, "CheckPort")
	delete(f, "ContextType")
	delete(f, "CheckSendContext")
	delete(f, "CheckRecvContext")
	delete(f, "CheckDomain")
	delete(f, "CheckPath")
	delete(f, "CheckMethod")
	delete(f, "StatusMask")
	delete(f, "ForwardProtocol")
	delete(f, "PortOverrides")
	delete(f, "CipherPolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEndpointGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEndpointGroupResponseParams struct {
	// <p>任务ID。</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEndpointGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEndpointGroupResponseParams `json:"Response"`
}

func (r *ModifyEndpointGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEndpointGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyForwardingRuleRequestParams struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 策略ID。
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// 七层转发规则ID。
	ForwardingRuleId *string `json:"ForwardingRuleId,omitnil,omitempty" name:"ForwardingRuleId"`

	// 七层转发规则条件信息。
	RuleConditions []*RuleCondition `json:"RuleConditions,omitnil,omitempty" name:"RuleConditions"`

	// 七层转发规则行为信息。
	RuleActions []*RuleAction `json:"RuleActions,omitnil,omitempty" name:"RuleActions"`

	// 回源Header信息。
	OriginHeaders []*OriginHeader `json:"OriginHeaders,omitnil,omitempty" name:"OriginHeaders"`

	// 是否开启回源sni。
	EnableOriginSni *bool `json:"EnableOriginSni,omitnil,omitempty" name:"EnableOriginSni"`

	// 回源sni。
	OriginSni *string `json:"OriginSni,omitnil,omitempty" name:"OriginSni"`

	// 回源host。
	OriginHost *string `json:"OriginHost,omitnil,omitempty" name:"OriginHost"`
}

type ModifyForwardingRuleRequest struct {
	*tchttp.BaseRequest
	
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 策略ID。
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// 七层转发规则ID。
	ForwardingRuleId *string `json:"ForwardingRuleId,omitnil,omitempty" name:"ForwardingRuleId"`

	// 七层转发规则条件信息。
	RuleConditions []*RuleCondition `json:"RuleConditions,omitnil,omitempty" name:"RuleConditions"`

	// 七层转发规则行为信息。
	RuleActions []*RuleAction `json:"RuleActions,omitnil,omitempty" name:"RuleActions"`

	// 回源Header信息。
	OriginHeaders []*OriginHeader `json:"OriginHeaders,omitnil,omitempty" name:"OriginHeaders"`

	// 是否开启回源sni。
	EnableOriginSni *bool `json:"EnableOriginSni,omitnil,omitempty" name:"EnableOriginSni"`

	// 回源sni。
	OriginSni *string `json:"OriginSni,omitnil,omitempty" name:"OriginSni"`

	// 回源host。
	OriginHost *string `json:"OriginHost,omitnil,omitempty" name:"OriginHost"`
}

func (r *ModifyForwardingRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyForwardingRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "ForwardingPolicyId")
	delete(f, "ForwardingRuleId")
	delete(f, "RuleConditions")
	delete(f, "RuleActions")
	delete(f, "OriginHeaders")
	delete(f, "EnableOriginSni")
	delete(f, "OriginSni")
	delete(f, "OriginHost")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyForwardingRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyForwardingRuleResponseParams struct {
	// 异步任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyForwardingRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyForwardingRuleResponseParams `json:"Response"`
}

func (r *ModifyForwardingRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyForwardingRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGlobalAcceleratorRequestParams struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 名称，最大长度不能超过60个字节。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述信息，最大长度不能超过100个字节。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 跨境类型。
	CrossBorderType *string `json:"CrossBorderType,omitnil,omitempty" name:"CrossBorderType"`

	// 代表是否跨境服务承诺。
	CrossBorderPromiseFlag *bool `json:"CrossBorderPromiseFlag,omitnil,omitempty" name:"CrossBorderPromiseFlag"`
}

type ModifyGlobalAcceleratorRequest struct {
	*tchttp.BaseRequest
	
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 名称，最大长度不能超过60个字节。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述信息，最大长度不能超过100个字节。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 跨境类型。
	CrossBorderType *string `json:"CrossBorderType,omitnil,omitempty" name:"CrossBorderType"`

	// 代表是否跨境服务承诺。
	CrossBorderPromiseFlag *bool `json:"CrossBorderPromiseFlag,omitnil,omitempty" name:"CrossBorderPromiseFlag"`
}

func (r *ModifyGlobalAcceleratorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGlobalAcceleratorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "CrossBorderType")
	delete(f, "CrossBorderPromiseFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGlobalAcceleratorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGlobalAcceleratorResponseParams struct {
	// 异步任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGlobalAcceleratorResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGlobalAcceleratorResponseParams `json:"Response"`
}

func (r *ModifyGlobalAcceleratorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGlobalAcceleratorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyListenerRequestParams struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 名称，最大长度不能超过60个字节。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述信息，最大长度不能超过100个字节。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 连接空闲等待时间。
	IdleTimeout *uint64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`

	// 是否开启会话保持。
	ClientAffinity *string `json:"ClientAffinity,omitnil,omitempty" name:"ClientAffinity"`

	// 会话保持时间。
	ClientAffinityTime *uint64 `json:"ClientAffinityTime,omitnil,omitempty" name:"ClientAffinityTime"`

	// 请求超时时间。
	RequestTimeout *uint64 `json:"RequestTimeout,omitnil,omitempty" name:"RequestTimeout"`

	// 是否打开七层获取源IP方式。
	XForwardedForRealIp *bool `json:"XForwardedForRealIp,omitnil,omitempty" name:"XForwardedForRealIp"`

	// 解析方式。UNIDIRECTIONAL：双向。MUTUAL：单向。
	CertificationType *string `json:"CertificationType,omitnil,omitempty" name:"CertificationType"`

	// 加密算法套件。
	CipherPolicyId *string `json:"CipherPolicyId,omitnil,omitempty" name:"CipherPolicyId"`

	// 服务器证书。
	ServerCertificates []*string `json:"ServerCertificates,omitnil,omitempty" name:"ServerCertificates"`

	// 客户端证书。
	ClientCaCertificates []*string `json:"ClientCaCertificates,omitnil,omitempty" name:"ClientCaCertificates"`

	// 四层获取源IP方式。
	GetRealIpType *string `json:"GetRealIpType,omitnil,omitempty" name:"GetRealIpType"`
}

type ModifyListenerRequest struct {
	*tchttp.BaseRequest
	
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 名称，最大长度不能超过60个字节。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述信息，最大长度不能超过100个字节。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 连接空闲等待时间。
	IdleTimeout *uint64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`

	// 是否开启会话保持。
	ClientAffinity *string `json:"ClientAffinity,omitnil,omitempty" name:"ClientAffinity"`

	// 会话保持时间。
	ClientAffinityTime *uint64 `json:"ClientAffinityTime,omitnil,omitempty" name:"ClientAffinityTime"`

	// 请求超时时间。
	RequestTimeout *uint64 `json:"RequestTimeout,omitnil,omitempty" name:"RequestTimeout"`

	// 是否打开七层获取源IP方式。
	XForwardedForRealIp *bool `json:"XForwardedForRealIp,omitnil,omitempty" name:"XForwardedForRealIp"`

	// 解析方式。UNIDIRECTIONAL：双向。MUTUAL：单向。
	CertificationType *string `json:"CertificationType,omitnil,omitempty" name:"CertificationType"`

	// 加密算法套件。
	CipherPolicyId *string `json:"CipherPolicyId,omitnil,omitempty" name:"CipherPolicyId"`

	// 服务器证书。
	ServerCertificates []*string `json:"ServerCertificates,omitnil,omitempty" name:"ServerCertificates"`

	// 客户端证书。
	ClientCaCertificates []*string `json:"ClientCaCertificates,omitnil,omitempty" name:"ClientCaCertificates"`

	// 四层获取源IP方式。
	GetRealIpType *string `json:"GetRealIpType,omitnil,omitempty" name:"GetRealIpType"`
}

func (r *ModifyListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "IdleTimeout")
	delete(f, "ClientAffinity")
	delete(f, "ClientAffinityTime")
	delete(f, "RequestTimeout")
	delete(f, "XForwardedForRealIp")
	delete(f, "CertificationType")
	delete(f, "CipherPolicyId")
	delete(f, "ServerCertificates")
	delete(f, "ClientCaCertificates")
	delete(f, "GetRealIpType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyListenerResponseParams struct {
	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyListenerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyListenerResponseParams `json:"Response"`
}

func (r *ModifyListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OriginHeader struct {
	// 键。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type PortOverride struct {
	// 监听端口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerPort *uint64 `json:"ListenerPort,omitnil,omitempty" name:"ListenerPort"`

	// 映射端口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndpointPort *uint64 `json:"EndpointPort,omitnil,omitempty" name:"EndpointPort"`
}

type PortRanges struct {
	// 起始端口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FromPort *uint64 `json:"FromPort,omitnil,omitempty" name:"FromPort"`

	// 终点端口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToPort *uint64 `json:"ToPort,omitnil,omitempty" name:"ToPort"`
}

type RuleAction struct {
	// 七层转发规则行为类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleActionType *string `json:"RuleActionType,omitnil,omitempty" name:"RuleActionType"`

	// 七层转发规则行为值
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleActionValue *string `json:"RuleActionValue,omitnil,omitempty" name:"RuleActionValue"`
}

type RuleCondition struct {
	// 七层转发规则条件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleConditionType *string `json:"RuleConditionType,omitnil,omitempty" name:"RuleConditionType"`

	// 七层转发规则条件值
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleConditionValue []*string `json:"RuleConditionValue,omitnil,omitempty" name:"RuleConditionValue"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}