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

package v20180709

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type BaradData struct {
	// 指标名（connum表示TCP活跃连接数；
	// new_conn表示新建TCP连接数；
	// inactive_conn表示非活跃连接数;
	// intraffic表示入流量；
	// outtraffic表示出流量；
	// alltraffic表示出流量和入流量之和；
	// inpkg表示入包速率；
	// outpkg表示出包速率；）
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 值数组
	Data []*float64 `json:"Data,omitempty" name:"Data"`

	// 值数组的大小
	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type BoundIpInfo struct {
	// IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 绑定的产品分类，取值[public（CVM、CLB产品），bm（黑石产品），eni（弹性网卡），vpngw（VPN网关）， natgw（NAT网关），waf（Web应用安全产品），fpc（金融产品），gaap（GAAP产品）, other(托管IP)]
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// 产品分类下的子类型，取值[cvm（CVM），lb（负载均衡器），eni（弹性网卡），vpngw（VPN），natgw（NAT），waf（WAF），fpc（金融），gaap（GAAP），other（托管IP），eip（黑石弹性IP）]
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// IP所属的资源实例ID，当绑定新IP时必须填写此字段；例如是弹性网卡的IP，则InstanceId填写弹性网卡的ID(eni-*); 如果绑定的是托管IP没有对应的资源实例ID，请填写"none";
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 运营商，0：电信；1：联通；2：移动；5：BGP
	IspCode *uint64 `json:"IspCode,omitempty" name:"IspCode"`
}

type CCAlarmThreshold struct {
	// CC告警阈值
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`
}

type CCEventRecord struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源的IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 攻击开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 攻击结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 总请求QPS峰值
	ReqQps *uint64 `json:"ReqQps,omitempty" name:"ReqQps"`

	// 攻击QPS峰值
	DropQps *uint64 `json:"DropQps,omitempty" name:"DropQps"`

	// 攻击状态，取值[0（攻击中）, 1（攻击结束）]
	AttackStatus *uint64 `json:"AttackStatus,omitempty" name:"AttackStatus"`

	// 资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainList *string `json:"DomainList,omitempty" name:"DomainList"`

	// uri列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UriList *string `json:"UriList,omitempty" name:"UriList"`

	// 攻击源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackipList *string `json:"AttackipList,omitempty" name:"AttackipList"`
}

type CCFrequencyRule struct {
	// CC的访问频率控制规则ID
	CCFrequencyRuleId *string `json:"CCFrequencyRuleId,omitempty" name:"CCFrequencyRuleId"`

	// URI字符串，必须以/开头，例如/abc/a.php，长度不超过31；当URI=/时，匹配模式只能选择前缀匹配；
	Uri *string `json:"Uri,omitempty" name:"Uri"`

	// User-Agent字符串，长度不超过80
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// Cookie字符串，长度不超过40
	Cookie *string `json:"Cookie,omitempty" name:"Cookie"`

	// 匹配规则，取值["include"(前缀匹配)，"equal"(完全匹配)]
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 统计周期，单位秒，取值[10, 30, 60]
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 访问次数，取值[1-10000]
	ReqNumber *uint64 `json:"ReqNumber,omitempty" name:"ReqNumber"`

	// 执行动作，取值["alg"（人机识别）, "drop"（拦截）]
	Act *string `json:"Act,omitempty" name:"Act"`

	// 执行时间，单位秒，取值[1-900]
	ExeDuration *uint64 `json:"ExeDuration,omitempty" name:"ExeDuration"`
}

type CCPolicy struct {
	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 匹配模式，取值[matching(匹配模式), speedlimit(限速模式)]
	Smode *string `json:"Smode,omitempty" name:"Smode"`

	// 策略id
	SetId *string `json:"SetId,omitempty" name:"SetId"`

	// 每分钟限制的次数
	Frequency *uint64 `json:"Frequency,omitempty" name:"Frequency"`

	// 执行策略模式，拦截或者验证码，取值[alg（验证码）, drop（拦截）]
	ExeMode *string `json:"ExeMode,omitempty" name:"ExeMode"`

	// 生效开关
	Switch *uint64 `json:"Switch,omitempty" name:"Switch"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 规则列表
	RuleList []*CCRule `json:"RuleList,omitempty" name:"RuleList"`

	// IP列表，如果不填时，请传空数组但不能为null；
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// cc防护类型，取值[http，https]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 可选字段，表示HTTPS的CC防护域名对应的转发规则ID;
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// HTTPS的CC防护域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type CCRule struct {
	// 规则的key, 可以为host、cgi、ua、referer
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// 规则的条件，可以为include、not_include、equal
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 规则的值，长度小于31字节
	Value *string `json:"Value,omitempty" name:"Value"`
}

type CCRuleConfig struct {
	// 统计周期，单位秒，取值[10, 30, 60]
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 访问次数，取值[1-10000]
	ReqNumber *uint64 `json:"ReqNumber,omitempty" name:"ReqNumber"`

	// 执行动作，取值["alg"（人机识别）, "drop"（拦截）]
	Action *string `json:"Action,omitempty" name:"Action"`

	// 执行时间，单位秒，取值[1-900]
	ExeDuration *uint64 `json:"ExeDuration,omitempty" name:"ExeDuration"`
}

// Predefined struct for user
type CreateBasicDDoSAlarmThresholdRequestParams struct {
	// 大禹子产品代号（basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// =get表示读取告警阈值；=set表示设置告警阈值；
	Method *string `json:"Method,omitempty" name:"Method"`

	// 可选，告警阈值类型，1-入流量，2-清洗流量；当Method为set时必须填写；
	AlarmType *uint64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// 可选，告警阈值，当Method为set时必须填写；当设置阈值为0时表示清除告警阈值配置；
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`
}

type CreateBasicDDoSAlarmThresholdRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// =get表示读取告警阈值；=set表示设置告警阈值；
	Method *string `json:"Method,omitempty" name:"Method"`

	// 可选，告警阈值类型，1-入流量，2-清洗流量；当Method为set时必须填写；
	AlarmType *uint64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// 可选，告警阈值，当Method为set时必须填写；当设置阈值为0时表示清除告警阈值配置；
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`
}

func (r *CreateBasicDDoSAlarmThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBasicDDoSAlarmThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Method")
	delete(f, "AlarmType")
	delete(f, "AlarmThreshold")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBasicDDoSAlarmThresholdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBasicDDoSAlarmThresholdResponseParams struct {
	// 当存在告警阈值配置时，返回告警阈值大于0，当不存在告警配置时，返回告警阈值为0；
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`

	// 告警阈值类型，1-入流量，2-清洗流量；当AlarmThreshold大于0时有效；
	AlarmType *uint64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBasicDDoSAlarmThresholdResponse struct {
	*tchttp.BaseResponse
	Response *CreateBasicDDoSAlarmThresholdResponseParams `json:"Response"`
}

func (r *CreateBasicDDoSAlarmThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBasicDDoSAlarmThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBoundIPRequestParams struct {
	// 大禹子产品代号（bgp表示独享包；bgp-multip表示共享包）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 绑定到资源实例的IP数组，当资源实例为高防包(独享包)时，数组只允许填1个IP；当没有要绑定的IP时可以为空数组；但是BoundDevList和UnBoundDevList至少有一个不为空；
	BoundDevList []*BoundIpInfo `json:"BoundDevList,omitempty" name:"BoundDevList"`

	// 与资源实例解绑的IP数组，当资源实例为高防包(独享包)时，数组只允许填1个IP；当没有要解绑的IP时可以为空数组；但是BoundDevList和UnBoundDevList至少有一个不为空；
	UnBoundDevList []*BoundIpInfo `json:"UnBoundDevList,omitempty" name:"UnBoundDevList"`

	// 已弃用，不填
	CopyPolicy *string `json:"CopyPolicy,omitempty" name:"CopyPolicy"`
}

type CreateBoundIPRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgp表示独享包；bgp-multip表示共享包）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 绑定到资源实例的IP数组，当资源实例为高防包(独享包)时，数组只允许填1个IP；当没有要绑定的IP时可以为空数组；但是BoundDevList和UnBoundDevList至少有一个不为空；
	BoundDevList []*BoundIpInfo `json:"BoundDevList,omitempty" name:"BoundDevList"`

	// 与资源实例解绑的IP数组，当资源实例为高防包(独享包)时，数组只允许填1个IP；当没有要解绑的IP时可以为空数组；但是BoundDevList和UnBoundDevList至少有一个不为空；
	UnBoundDevList []*BoundIpInfo `json:"UnBoundDevList,omitempty" name:"UnBoundDevList"`

	// 已弃用，不填
	CopyPolicy *string `json:"CopyPolicy,omitempty" name:"CopyPolicy"`
}

func (r *CreateBoundIPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBoundIPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "BoundDevList")
	delete(f, "UnBoundDevList")
	delete(f, "CopyPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBoundIPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBoundIPResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBoundIPResponse struct {
	*tchttp.BaseResponse
	Response *CreateBoundIPResponseParams `json:"Response"`
}

func (r *CreateBoundIPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBoundIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCCFrequencyRulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 7层转发规则ID（通过获取7层转发规则接口可以获取规则ID）
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 匹配规则，取值["include"(前缀匹配)，"equal"(完全匹配)]
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 统计周期，单位秒，取值[10, 30, 60]
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 访问次数，取值[1-10000]
	ReqNumber *uint64 `json:"ReqNumber,omitempty" name:"ReqNumber"`

	// 执行动作，取值["alg"（人机识别）, "drop"（拦截）]
	Act *string `json:"Act,omitempty" name:"Act"`

	// 执行时间，单位秒，取值[1-900]
	ExeDuration *uint64 `json:"ExeDuration,omitempty" name:"ExeDuration"`

	// URI字符串，必须以/开头，例如/abc/a.php，长度不超过31；当URI=/时，匹配模式只能选择前缀匹配；
	Uri *string `json:"Uri,omitempty" name:"Uri"`

	// User-Agent字符串，长度不超过80
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// Cookie字符串，长度不超过40
	Cookie *string `json:"Cookie,omitempty" name:"Cookie"`
}

type CreateCCFrequencyRulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 7层转发规则ID（通过获取7层转发规则接口可以获取规则ID）
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 匹配规则，取值["include"(前缀匹配)，"equal"(完全匹配)]
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 统计周期，单位秒，取值[10, 30, 60]
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 访问次数，取值[1-10000]
	ReqNumber *uint64 `json:"ReqNumber,omitempty" name:"ReqNumber"`

	// 执行动作，取值["alg"（人机识别）, "drop"（拦截）]
	Act *string `json:"Act,omitempty" name:"Act"`

	// 执行时间，单位秒，取值[1-900]
	ExeDuration *uint64 `json:"ExeDuration,omitempty" name:"ExeDuration"`

	// URI字符串，必须以/开头，例如/abc/a.php，长度不超过31；当URI=/时，匹配模式只能选择前缀匹配；
	Uri *string `json:"Uri,omitempty" name:"Uri"`

	// User-Agent字符串，长度不超过80
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// Cookie字符串，长度不超过40
	Cookie *string `json:"Cookie,omitempty" name:"Cookie"`
}

func (r *CreateCCFrequencyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCFrequencyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleId")
	delete(f, "Mode")
	delete(f, "Period")
	delete(f, "ReqNumber")
	delete(f, "Act")
	delete(f, "ExeDuration")
	delete(f, "Uri")
	delete(f, "UserAgent")
	delete(f, "Cookie")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCCFrequencyRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCCFrequencyRulesResponseParams struct {
	// CC防护的访问频率控制规则ID
	CCFrequencyRuleId *string `json:"CCFrequencyRuleId,omitempty" name:"CCFrequencyRuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCCFrequencyRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateCCFrequencyRulesResponseParams `json:"Response"`
}

func (r *CreateCCFrequencyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCFrequencyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCCSelfDefinePolicyRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// CC策略描述
	Policy *CCPolicy `json:"Policy,omitempty" name:"Policy"`
}

type CreateCCSelfDefinePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// CC策略描述
	Policy *CCPolicy `json:"Policy,omitempty" name:"Policy"`
}

func (r *CreateCCSelfDefinePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCSelfDefinePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Policy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCCSelfDefinePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCCSelfDefinePolicyResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCCSelfDefinePolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateCCSelfDefinePolicyResponseParams `json:"Response"`
}

func (r *CreateCCSelfDefinePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCSelfDefinePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSPolicyCaseRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 策略场景名，字符串长度小于64
	CaseName *string `json:"CaseName,omitempty" name:"CaseName"`

	// 开发平台，取值[PC（PC客户端）， MOBILE（移动端）， TV（电视端）， SERVER（主机）]
	PlatformTypes []*string `json:"PlatformTypes,omitempty" name:"PlatformTypes"`

	// 细分品类，取值[WEB（网站）， GAME（游戏）， APP（应用）， OTHER（其他）]
	AppType *string `json:"AppType,omitempty" name:"AppType"`

	// 应用协议，取值[tcp（TCP协议），udp（UDP协议），icmp（ICMP协议），all（其他协议）]
	AppProtocols []*string `json:"AppProtocols,omitempty" name:"AppProtocols"`

	// TCP业务起始端口，取值(0, 65535]
	TcpSportStart *string `json:"TcpSportStart,omitempty" name:"TcpSportStart"`

	// TCP业务结束端口，取值(0, 65535]，必须大于等于TCP业务起始端口
	TcpSportEnd *string `json:"TcpSportEnd,omitempty" name:"TcpSportEnd"`

	// UDP业务起始端口，取值范围(0, 65535]
	UdpSportStart *string `json:"UdpSportStart,omitempty" name:"UdpSportStart"`

	// UDP业务结束端口，取值范围(0, 65535)，必须大于等于UDP业务起始端口
	UdpSportEnd *string `json:"UdpSportEnd,omitempty" name:"UdpSportEnd"`

	// 是否有海外客户，取值[no（没有）, yes（有）]
	HasAbroad *string `json:"HasAbroad,omitempty" name:"HasAbroad"`

	// 是否会主动对外发起TCP请求，取值[no（不会）, yes（会）]
	HasInitiateTcp *string `json:"HasInitiateTcp,omitempty" name:"HasInitiateTcp"`

	// 是否会主动对外发起UDP业务请求，取值[no（不会）, yes（会）]
	HasInitiateUdp *string `json:"HasInitiateUdp,omitempty" name:"HasInitiateUdp"`

	// 主动发起TCP请求的端口，取值范围(0, 65535]
	PeerTcpPort *string `json:"PeerTcpPort,omitempty" name:"PeerTcpPort"`

	// 主动发起UDP请求的端口，取值范围(0, 65535]
	PeerUdpPort *string `json:"PeerUdpPort,omitempty" name:"PeerUdpPort"`

	// TCP载荷的固定特征码，字符串长度小于512
	TcpFootprint *string `json:"TcpFootprint,omitempty" name:"TcpFootprint"`

	// UDP载荷的固定特征码，字符串长度小于512
	UdpFootprint *string `json:"UdpFootprint,omitempty" name:"UdpFootprint"`

	// Web业务的API的URL
	WebApiUrl []*string `json:"WebApiUrl,omitempty" name:"WebApiUrl"`

	// TCP业务报文长度最小值，取值范围(0, 1500)
	MinTcpPackageLen *string `json:"MinTcpPackageLen,omitempty" name:"MinTcpPackageLen"`

	// TCP业务报文长度最大值，取值范围(0, 1500)，必须大于等于TCP业务报文长度最小值
	MaxTcpPackageLen *string `json:"MaxTcpPackageLen,omitempty" name:"MaxTcpPackageLen"`

	// UDP业务报文长度最小值，取值范围(0, 1500)
	MinUdpPackageLen *string `json:"MinUdpPackageLen,omitempty" name:"MinUdpPackageLen"`

	// UDP业务报文长度最大值，取值范围(0, 1500)，必须大于等于UDP业务报文长度最小值
	MaxUdpPackageLen *string `json:"MaxUdpPackageLen,omitempty" name:"MaxUdpPackageLen"`

	// 是否有VPN业务，取值[no（没有）, yes（有）]
	HasVPN *string `json:"HasVPN,omitempty" name:"HasVPN"`

	// TCP业务端口列表，同时支持单个端口和端口段，字符串格式，例如：80,443,700-800,53,1000-3000
	TcpPortList *string `json:"TcpPortList,omitempty" name:"TcpPortList"`

	// UDP业务端口列表，同时支持单个端口和端口段，字符串格式，例如：80,443,700-800,53,1000-3000
	UdpPortList *string `json:"UdpPortList,omitempty" name:"UdpPortList"`
}

type CreateDDoSPolicyCaseRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 策略场景名，字符串长度小于64
	CaseName *string `json:"CaseName,omitempty" name:"CaseName"`

	// 开发平台，取值[PC（PC客户端）， MOBILE（移动端）， TV（电视端）， SERVER（主机）]
	PlatformTypes []*string `json:"PlatformTypes,omitempty" name:"PlatformTypes"`

	// 细分品类，取值[WEB（网站）， GAME（游戏）， APP（应用）， OTHER（其他）]
	AppType *string `json:"AppType,omitempty" name:"AppType"`

	// 应用协议，取值[tcp（TCP协议），udp（UDP协议），icmp（ICMP协议），all（其他协议）]
	AppProtocols []*string `json:"AppProtocols,omitempty" name:"AppProtocols"`

	// TCP业务起始端口，取值(0, 65535]
	TcpSportStart *string `json:"TcpSportStart,omitempty" name:"TcpSportStart"`

	// TCP业务结束端口，取值(0, 65535]，必须大于等于TCP业务起始端口
	TcpSportEnd *string `json:"TcpSportEnd,omitempty" name:"TcpSportEnd"`

	// UDP业务起始端口，取值范围(0, 65535]
	UdpSportStart *string `json:"UdpSportStart,omitempty" name:"UdpSportStart"`

	// UDP业务结束端口，取值范围(0, 65535)，必须大于等于UDP业务起始端口
	UdpSportEnd *string `json:"UdpSportEnd,omitempty" name:"UdpSportEnd"`

	// 是否有海外客户，取值[no（没有）, yes（有）]
	HasAbroad *string `json:"HasAbroad,omitempty" name:"HasAbroad"`

	// 是否会主动对外发起TCP请求，取值[no（不会）, yes（会）]
	HasInitiateTcp *string `json:"HasInitiateTcp,omitempty" name:"HasInitiateTcp"`

	// 是否会主动对外发起UDP业务请求，取值[no（不会）, yes（会）]
	HasInitiateUdp *string `json:"HasInitiateUdp,omitempty" name:"HasInitiateUdp"`

	// 主动发起TCP请求的端口，取值范围(0, 65535]
	PeerTcpPort *string `json:"PeerTcpPort,omitempty" name:"PeerTcpPort"`

	// 主动发起UDP请求的端口，取值范围(0, 65535]
	PeerUdpPort *string `json:"PeerUdpPort,omitempty" name:"PeerUdpPort"`

	// TCP载荷的固定特征码，字符串长度小于512
	TcpFootprint *string `json:"TcpFootprint,omitempty" name:"TcpFootprint"`

	// UDP载荷的固定特征码，字符串长度小于512
	UdpFootprint *string `json:"UdpFootprint,omitempty" name:"UdpFootprint"`

	// Web业务的API的URL
	WebApiUrl []*string `json:"WebApiUrl,omitempty" name:"WebApiUrl"`

	// TCP业务报文长度最小值，取值范围(0, 1500)
	MinTcpPackageLen *string `json:"MinTcpPackageLen,omitempty" name:"MinTcpPackageLen"`

	// TCP业务报文长度最大值，取值范围(0, 1500)，必须大于等于TCP业务报文长度最小值
	MaxTcpPackageLen *string `json:"MaxTcpPackageLen,omitempty" name:"MaxTcpPackageLen"`

	// UDP业务报文长度最小值，取值范围(0, 1500)
	MinUdpPackageLen *string `json:"MinUdpPackageLen,omitempty" name:"MinUdpPackageLen"`

	// UDP业务报文长度最大值，取值范围(0, 1500)，必须大于等于UDP业务报文长度最小值
	MaxUdpPackageLen *string `json:"MaxUdpPackageLen,omitempty" name:"MaxUdpPackageLen"`

	// 是否有VPN业务，取值[no（没有）, yes（有）]
	HasVPN *string `json:"HasVPN,omitempty" name:"HasVPN"`

	// TCP业务端口列表，同时支持单个端口和端口段，字符串格式，例如：80,443,700-800,53,1000-3000
	TcpPortList *string `json:"TcpPortList,omitempty" name:"TcpPortList"`

	// UDP业务端口列表，同时支持单个端口和端口段，字符串格式，例如：80,443,700-800,53,1000-3000
	UdpPortList *string `json:"UdpPortList,omitempty" name:"UdpPortList"`
}

func (r *CreateDDoSPolicyCaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSPolicyCaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "CaseName")
	delete(f, "PlatformTypes")
	delete(f, "AppType")
	delete(f, "AppProtocols")
	delete(f, "TcpSportStart")
	delete(f, "TcpSportEnd")
	delete(f, "UdpSportStart")
	delete(f, "UdpSportEnd")
	delete(f, "HasAbroad")
	delete(f, "HasInitiateTcp")
	delete(f, "HasInitiateUdp")
	delete(f, "PeerTcpPort")
	delete(f, "PeerUdpPort")
	delete(f, "TcpFootprint")
	delete(f, "UdpFootprint")
	delete(f, "WebApiUrl")
	delete(f, "MinTcpPackageLen")
	delete(f, "MaxTcpPackageLen")
	delete(f, "MinUdpPackageLen")
	delete(f, "MaxUdpPackageLen")
	delete(f, "HasVPN")
	delete(f, "TcpPortList")
	delete(f, "UdpPortList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDDoSPolicyCaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSPolicyCaseResponseParams struct {
	// 策略场景ID
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDDoSPolicyCaseResponse struct {
	*tchttp.BaseResponse
	Response *CreateDDoSPolicyCaseResponseParams `json:"Response"`
}

func (r *CreateDDoSPolicyCaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSPolicyCaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSPolicyRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 协议禁用，必须填写且数组长度必须为1
	DropOptions []*DDoSPolicyDropOption `json:"DropOptions,omitempty" name:"DropOptions"`

	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 端口禁用，当没有禁用端口时填空数组
	PortLimits []*DDoSPolicyPortLimit `json:"PortLimits,omitempty" name:"PortLimits"`

	// 请求源IP黑白名单，当没有IP黑白名单时填空数组
	IpAllowDenys []*IpBlackWhite `json:"IpAllowDenys,omitempty" name:"IpAllowDenys"`

	// 报文过滤，当没有报文过滤时填空数组
	PacketFilters []*DDoSPolicyPacketFilter `json:"PacketFilters,omitempty" name:"PacketFilters"`

	// 水印策略参数，当没有启用水印功能时填空数组，最多只能传一条水印策略（即数组大小不超过1）
	WaterPrint []*WaterPrintPolicy `json:"WaterPrint,omitempty" name:"WaterPrint"`
}

type CreateDDoSPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 协议禁用，必须填写且数组长度必须为1
	DropOptions []*DDoSPolicyDropOption `json:"DropOptions,omitempty" name:"DropOptions"`

	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 端口禁用，当没有禁用端口时填空数组
	PortLimits []*DDoSPolicyPortLimit `json:"PortLimits,omitempty" name:"PortLimits"`

	// 请求源IP黑白名单，当没有IP黑白名单时填空数组
	IpAllowDenys []*IpBlackWhite `json:"IpAllowDenys,omitempty" name:"IpAllowDenys"`

	// 报文过滤，当没有报文过滤时填空数组
	PacketFilters []*DDoSPolicyPacketFilter `json:"PacketFilters,omitempty" name:"PacketFilters"`

	// 水印策略参数，当没有启用水印功能时填空数组，最多只能传一条水印策略（即数组大小不超过1）
	WaterPrint []*WaterPrintPolicy `json:"WaterPrint,omitempty" name:"WaterPrint"`
}

func (r *CreateDDoSPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "DropOptions")
	delete(f, "Name")
	delete(f, "PortLimits")
	delete(f, "IpAllowDenys")
	delete(f, "PacketFilters")
	delete(f, "WaterPrint")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDDoSPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSPolicyResponseParams struct {
	// 策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDDoSPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateDDoSPolicyResponseParams `json:"Response"`
}

func (r *CreateDDoSPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceNameRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源实例名称，长度不超过32个字符
	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源实例名称，长度不超过32个字符
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceNameResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceNameResponseParams `json:"Response"`
}

func (r *CreateInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL4HealthConfigRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 四层健康检查配置数组
	HealthConfig []*L4HealthConfig `json:"HealthConfig,omitempty" name:"HealthConfig"`
}

type CreateL4HealthConfigRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 四层健康检查配置数组
	HealthConfig []*L4HealthConfig `json:"HealthConfig,omitempty" name:"HealthConfig"`
}

func (r *CreateL4HealthConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL4HealthConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "HealthConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateL4HealthConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL4HealthConfigResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateL4HealthConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateL4HealthConfigResponseParams `json:"Response"`
}

func (r *CreateL4HealthConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL4HealthConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL4RulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则列表
	Rules []*L4RuleEntry `json:"Rules,omitempty" name:"Rules"`
}

type CreateL4RulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则列表
	Rules []*L4RuleEntry `json:"Rules,omitempty" name:"Rules"`
}

func (r *CreateL4RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL4RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateL4RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL4RulesResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateL4RulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateL4RulesResponseParams `json:"Response"`
}

func (r *CreateL4RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL4RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL7CCRuleRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 操作码，取值[query(表示查询)，add(表示添加)，del(表示删除)]
	Method *string `json:"Method,omitempty" name:"Method"`

	// 7层转发规则ID，例如：rule-0000001
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 7层CC自定义规则参数，当操作码为query时，可以不用填写；当操作码为add或del时，必须填写，且数组长度只能为1；
	RuleConfig []*CCRuleConfig `json:"RuleConfig,omitempty" name:"RuleConfig"`
}

type CreateL7CCRuleRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 操作码，取值[query(表示查询)，add(表示添加)，del(表示删除)]
	Method *string `json:"Method,omitempty" name:"Method"`

	// 7层转发规则ID，例如：rule-0000001
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 7层CC自定义规则参数，当操作码为query时，可以不用填写；当操作码为add或del时，必须填写，且数组长度只能为1；
	RuleConfig []*CCRuleConfig `json:"RuleConfig,omitempty" name:"RuleConfig"`
}

func (r *CreateL7CCRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7CCRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Method")
	delete(f, "RuleId")
	delete(f, "RuleConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateL7CCRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL7CCRuleResponseParams struct {
	// 7层CC自定义规则参数，当没有开启CC自定义规则时，返回空数组
	RuleConfig []*CCRuleConfig `json:"RuleConfig,omitempty" name:"RuleConfig"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateL7CCRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateL7CCRuleResponseParams `json:"Response"`
}

func (r *CreateL7CCRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7CCRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL7HealthConfigRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 七层健康检查配置数组
	HealthConfig []*L7HealthConfig `json:"HealthConfig,omitempty" name:"HealthConfig"`
}

type CreateL7HealthConfigRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 七层健康检查配置数组
	HealthConfig []*L7HealthConfig `json:"HealthConfig,omitempty" name:"HealthConfig"`
}

func (r *CreateL7HealthConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7HealthConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "HealthConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateL7HealthConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL7HealthConfigResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateL7HealthConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateL7HealthConfigResponseParams `json:"Response"`
}

func (r *CreateL7HealthConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7HealthConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL7RuleCertRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID，例如高防IP实例的ID，高防IP专业版实例的ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 证书类型，当为协议为HTTPS协议时必须填，取值[2(腾讯云托管证书)]
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// 当证书来源为腾讯云托管证书时，此字段必须填写托管证书ID
	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`

	// 当证书来源为自有证书时，此字段必须填写证书内容；(因已不再支持自有证书，此字段已弃用，请不用填写此字段)
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// 当证书来源为自有证书时，此字段必须填写证书密钥；(因已不再支持自有证书，此字段已弃用，请不用填写此字段)
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
}

type CreateL7RuleCertRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID，例如高防IP实例的ID，高防IP专业版实例的ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 证书类型，当为协议为HTTPS协议时必须填，取值[2(腾讯云托管证书)]
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// 当证书来源为腾讯云托管证书时，此字段必须填写托管证书ID
	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`

	// 当证书来源为自有证书时，此字段必须填写证书内容；(因已不再支持自有证书，此字段已弃用，请不用填写此字段)
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// 当证书来源为自有证书时，此字段必须填写证书密钥；(因已不再支持自有证书，此字段已弃用，请不用填写此字段)
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
}

func (r *CreateL7RuleCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RuleCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleId")
	delete(f, "CertType")
	delete(f, "SSLId")
	delete(f, "Cert")
	delete(f, "PrivateKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateL7RuleCertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL7RuleCertResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateL7RuleCertResponse struct {
	*tchttp.BaseResponse
	Response *CreateL7RuleCertResponseParams `json:"Response"`
}

func (r *CreateL7RuleCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RuleCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL7RulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则列表
	Rules []*L7RuleEntry `json:"Rules,omitempty" name:"Rules"`
}

type CreateL7RulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则列表
	Rules []*L7RuleEntry `json:"Rules,omitempty" name:"Rules"`
}

func (r *CreateL7RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateL7RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL7RulesResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateL7RulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateL7RulesResponseParams `json:"Response"`
}

func (r *CreateL7RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL7RulesUploadRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则列表
	Rules []*L7RuleEntry `json:"Rules,omitempty" name:"Rules"`
}

type CreateL7RulesUploadRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则列表
	Rules []*L7RuleEntry `json:"Rules,omitempty" name:"Rules"`
}

func (r *CreateL7RulesUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RulesUploadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateL7RulesUploadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL7RulesUploadResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateL7RulesUploadResponse struct {
	*tchttp.BaseResponse
	Response *CreateL7RulesUploadResponseParams `json:"Response"`
}

func (r *CreateL7RulesUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RulesUploadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetReturnRequestParams struct {
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type CreateNetReturnRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *CreateNetReturnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetReturnRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNetReturnRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetReturnResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNetReturnResponse struct {
	*tchttp.BaseResponse
	Response *CreateNetReturnResponseParams `json:"Response"`
}

func (r *CreateNetReturnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetReturnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNewL4RulesRequestParams struct {
	// 高防产品代号：bgpip
	Business *string `json:"Business,omitempty" name:"Business"`

	// 添加规则资源列表
	IdList []*string `json:"IdList,omitempty" name:"IdList"`

	// 添加规则IP列表
	VipList []*string `json:"VipList,omitempty" name:"VipList"`

	// 规则列表
	Rules []*L4RuleEntry `json:"Rules,omitempty" name:"Rules"`
}

type CreateNewL4RulesRequest struct {
	*tchttp.BaseRequest
	
	// 高防产品代号：bgpip
	Business *string `json:"Business,omitempty" name:"Business"`

	// 添加规则资源列表
	IdList []*string `json:"IdList,omitempty" name:"IdList"`

	// 添加规则IP列表
	VipList []*string `json:"VipList,omitempty" name:"VipList"`

	// 规则列表
	Rules []*L4RuleEntry `json:"Rules,omitempty" name:"Rules"`
}

func (r *CreateNewL4RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNewL4RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "IdList")
	delete(f, "VipList")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNewL4RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNewL4RulesResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNewL4RulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateNewL4RulesResponseParams `json:"Response"`
}

func (r *CreateNewL4RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNewL4RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNewL7RulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID列表
	IdList []*string `json:"IdList,omitempty" name:"IdList"`

	// 资源IP列表
	VipList []*string `json:"VipList,omitempty" name:"VipList"`

	// 规则列表
	Rules []*L7RuleEntry `json:"Rules,omitempty" name:"Rules"`
}

type CreateNewL7RulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID列表
	IdList []*string `json:"IdList,omitempty" name:"IdList"`

	// 资源IP列表
	VipList []*string `json:"VipList,omitempty" name:"VipList"`

	// 规则列表
	Rules []*L7RuleEntry `json:"Rules,omitempty" name:"Rules"`
}

func (r *CreateNewL7RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNewL7RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "IdList")
	delete(f, "VipList")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNewL7RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNewL7RulesResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNewL7RulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateNewL7RulesResponseParams `json:"Response"`
}

func (r *CreateNewL7RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNewL7RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNewL7RulesUploadRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID列表
	IdList []*string `json:"IdList,omitempty" name:"IdList"`

	// 资源IP列表
	VipList []*string `json:"VipList,omitempty" name:"VipList"`

	// 规则列表
	Rules []*L7RuleEntry `json:"Rules,omitempty" name:"Rules"`
}

type CreateNewL7RulesUploadRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID列表
	IdList []*string `json:"IdList,omitempty" name:"IdList"`

	// 资源IP列表
	VipList []*string `json:"VipList,omitempty" name:"VipList"`

	// 规则列表
	Rules []*L7RuleEntry `json:"Rules,omitempty" name:"Rules"`
}

func (r *CreateNewL7RulesUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNewL7RulesUploadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "IdList")
	delete(f, "VipList")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNewL7RulesUploadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNewL7RulesUploadResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNewL7RulesUploadResponse struct {
	*tchttp.BaseResponse
	Response *CreateNewL7RulesUploadResponseParams `json:"Response"`
}

func (r *CreateNewL7RulesUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNewL7RulesUploadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUnblockIpRequestParams struct {
	// IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 解封类型（user：自助解封；auto：自动解封； update：升级解封；bind：绑定高防包解封）
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

type CreateUnblockIpRequest struct {
	*tchttp.BaseRequest
	
	// IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 解封类型（user：自助解封；auto：自动解封； update：升级解封；bind：绑定高防包解封）
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

func (r *CreateUnblockIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUnblockIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ip")
	delete(f, "ActionType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUnblockIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUnblockIpResponseParams struct {
	// IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 解封类型（user：自助解封；auto：自动解封； update：升级解封；bind：绑定高防包解封）
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 解封时间（预计解封时间）
	UnblockTime *string `json:"UnblockTime,omitempty" name:"UnblockTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUnblockIpResponse struct {
	*tchttp.BaseResponse
	Response *CreateUnblockIpResponseParams `json:"Response"`
}

func (r *CreateUnblockIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUnblockIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DDoSAlarmThreshold struct {
	// 告警阈值类型，1-入流量，2-清洗流量
	AlarmType *uint64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// 告警阈值，大于0（目前排定的值）
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`
}

type DDoSAttackSourceRecord struct {
	// 攻击源ip
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 省份（国内有效，不包含港澳台）
	Province *string `json:"Province,omitempty" name:"Province"`

	// 国家
	Nation *string `json:"Nation,omitempty" name:"Nation"`

	// 累计攻击包量
	PacketSum *uint64 `json:"PacketSum,omitempty" name:"PacketSum"`

	// 累计攻击流量
	PacketLen *uint64 `json:"PacketLen,omitempty" name:"PacketLen"`
}

type DDoSEventRecord struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源的IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 攻击开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 攻击结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 攻击最大带宽
	Mbps *uint64 `json:"Mbps,omitempty" name:"Mbps"`

	// 攻击最大包速率
	Pps *uint64 `json:"Pps,omitempty" name:"Pps"`

	// 攻击类型
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 是否被封堵，取值[1（是），0（否），2（无效值）]
	BlockFlag *uint64 `json:"BlockFlag,omitempty" name:"BlockFlag"`

	// 是否超过弹性防护峰值，取值取值[yes(是)，no(否)，空字符串（未知值）]
	OverLoad *string `json:"OverLoad,omitempty" name:"OverLoad"`

	// 攻击状态，取值[0（攻击中）, 1（攻击结束）]
	AttackStatus *uint64 `json:"AttackStatus,omitempty" name:"AttackStatus"`

	// 资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 攻击事件Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

type DDoSPolicyDropOption struct {
	// 禁用TCP协议，取值范围[0,1]
	DropTcp *uint64 `json:"DropTcp,omitempty" name:"DropTcp"`

	// 禁用UDP协议，取值范围[0,1]
	DropUdp *uint64 `json:"DropUdp,omitempty" name:"DropUdp"`

	// 禁用ICMP协议，取值范围[0,1]
	DropIcmp *uint64 `json:"DropIcmp,omitempty" name:"DropIcmp"`

	// 禁用其他协议，取值范围[0,1]
	DropOther *uint64 `json:"DropOther,omitempty" name:"DropOther"`

	// 拒绝海外流量，取值范围[0,1]
	DropAbroad *uint64 `json:"DropAbroad,omitempty" name:"DropAbroad"`

	// 空连接防护，取值范围[0,1]
	CheckSyncConn *uint64 `json:"CheckSyncConn,omitempty" name:"CheckSyncConn"`

	// 基于来源IP及目的IP的新建连接抑制，取值范围[0,4294967295]
	SdNewLimit *uint64 `json:"SdNewLimit,omitempty" name:"SdNewLimit"`

	// 基于目的IP的新建连接抑制，取值范围[0,4294967295]
	DstNewLimit *uint64 `json:"DstNewLimit,omitempty" name:"DstNewLimit"`

	// 基于来源IP及目的IP的并发连接抑制，取值范围[0,4294967295]
	SdConnLimit *uint64 `json:"SdConnLimit,omitempty" name:"SdConnLimit"`

	// 基于目的IP的并发连接抑制，取值范围[0,4294967295]
	DstConnLimit *uint64 `json:"DstConnLimit,omitempty" name:"DstConnLimit"`

	// 基于连接抑制触发阈值，取值范围[0,4294967295]
	BadConnThreshold *uint64 `json:"BadConnThreshold,omitempty" name:"BadConnThreshold"`

	// 异常连接检测条件，空连接防护开关，，取值范围[0,1]
	NullConnEnable *uint64 `json:"NullConnEnable,omitempty" name:"NullConnEnable"`

	// 异常连接检测条件，连接超时，，取值范围[0,65535]
	ConnTimeout *uint64 `json:"ConnTimeout,omitempty" name:"ConnTimeout"`

	// 异常连接检测条件，syn占比ack百分比，，取值范围[0,100]
	SynRate *uint64 `json:"SynRate,omitempty" name:"SynRate"`

	// 异常连接检测条件，syn阈值，取值范围[0,100]
	SynLimit *uint64 `json:"SynLimit,omitempty" name:"SynLimit"`

	// tcp限速，取值范围[0,4294967295]
	DTcpMbpsLimit *uint64 `json:"DTcpMbpsLimit,omitempty" name:"DTcpMbpsLimit"`

	// udp限速，取值范围[0,4294967295]
	DUdpMbpsLimit *uint64 `json:"DUdpMbpsLimit,omitempty" name:"DUdpMbpsLimit"`

	// icmp限速，取值范围[0,4294967295]
	DIcmpMbpsLimit *uint64 `json:"DIcmpMbpsLimit,omitempty" name:"DIcmpMbpsLimit"`

	// other协议限速，取值范围[0,4294967295]
	DOtherMbpsLimit *uint64 `json:"DOtherMbpsLimit,omitempty" name:"DOtherMbpsLimit"`
}

type DDoSPolicyPacketFilter struct {
	// 协议，取值范围[tcp,udp,icmp,all]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 开始源端口，取值范围[0,65535]
	SportStart *uint64 `json:"SportStart,omitempty" name:"SportStart"`

	// 结束源端口，取值范围[0,65535]
	SportEnd *uint64 `json:"SportEnd,omitempty" name:"SportEnd"`

	// 开始目的端口，取值范围[0,65535]
	DportStart *uint64 `json:"DportStart,omitempty" name:"DportStart"`

	// 结束目的端口，取值范围[0,65535]
	DportEnd *uint64 `json:"DportEnd,omitempty" name:"DportEnd"`

	// 最小包长，取值范围[0,1500]
	PktlenMin *uint64 `json:"PktlenMin,omitempty" name:"PktlenMin"`

	// 最大包长，取值范围[0,1500]
	PktlenMax *uint64 `json:"PktlenMax,omitempty" name:"PktlenMax"`

	// 是否检测载荷，取值范围[
	// begin_l3(IP头)
	// begin_l4(TCP头)
	// begin_l5(载荷)
	// no_match(不检测)
	// ]
	MatchBegin *string `json:"MatchBegin,omitempty" name:"MatchBegin"`

	// 是否是正则表达式，取值范围[sunday(表示关键字),pcre(表示正则表达式)]
	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`

	// 关键字或正则表达式
	Str *string `json:"Str,omitempty" name:"Str"`

	// 检测深度，取值范围[0,1500]
	Depth *uint64 `json:"Depth,omitempty" name:"Depth"`

	// 检测偏移量，取值范围[0,1500]
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 是否包括，取值范围[0(表示不包含),1(表示包含)]
	IsNot *uint64 `json:"IsNot,omitempty" name:"IsNot"`

	// 策略动作，取值范围[drop，drop_black，drop_rst，drop_black_rst，transmit]
	Action *string `json:"Action,omitempty" name:"Action"`
}

type DDoSPolicyPortLimit struct {
	// 协议，取值范围[tcp,udp,all]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 开始目的端口，取值范围[0,65535]
	DPortStart *uint64 `json:"DPortStart,omitempty" name:"DPortStart"`

	// 结束目的端口，取值范围[0,65535]，要求大于等于开始目的端口
	DPortEnd *uint64 `json:"DPortEnd,omitempty" name:"DPortEnd"`

	// 开始源端口，取值范围[0,65535]
	// 注意：此字段可能返回 null，表示取不到有效值。
	SPortStart *uint64 `json:"SPortStart,omitempty" name:"SPortStart"`

	// 结束源端口，取值范围[0,65535]，要求大于等于开始源端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	SPortEnd *uint64 `json:"SPortEnd,omitempty" name:"SPortEnd"`

	// 执行动作，取值[drop(丢弃) ，transmit(转发)]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 禁用端口类型，取值[0（目的端口范围禁用）， 1（源端口范围禁用）， 2（目的和源端口范围同时禁用）]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
}

type DDosPolicy struct {
	// 策略绑定的资源
	Resources []*ResourceIp `json:"Resources,omitempty" name:"Resources"`

	// 禁用协议
	DropOptions *DDoSPolicyDropOption `json:"DropOptions,omitempty" name:"DropOptions"`

	// 禁用端口
	PortLimits []*DDoSPolicyPortLimit `json:"PortLimits,omitempty" name:"PortLimits"`

	// 报文过滤
	PacketFilters []*DDoSPolicyPacketFilter `json:"PacketFilters,omitempty" name:"PacketFilters"`

	// 黑白IP名单
	IpBlackWhiteLists []*IpBlackWhite `json:"IpBlackWhiteLists,omitempty" name:"IpBlackWhiteLists"`

	// 策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 策略创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 水印策略参数，最多只有一个，当没有水印策略时数组为空
	WaterPrint []*WaterPrintPolicy `json:"WaterPrint,omitempty" name:"WaterPrint"`

	// 水印密钥，最多只有2个，当没有水印策略时数组为空
	WaterKey []*WaterPrintKey `json:"WaterKey,omitempty" name:"WaterKey"`

	// 策略绑定的资源实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	BoundResources []*string `json:"BoundResources,omitempty" name:"BoundResources"`

	// 策略所属的策略场景
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
}

// Predefined struct for user
type DeleteCCFrequencyRulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// CC防护的访问频率控制规则ID
	CCFrequencyRuleId *string `json:"CCFrequencyRuleId,omitempty" name:"CCFrequencyRuleId"`
}

type DeleteCCFrequencyRulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// CC防护的访问频率控制规则ID
	CCFrequencyRuleId *string `json:"CCFrequencyRuleId,omitempty" name:"CCFrequencyRuleId"`
}

func (r *DeleteCCFrequencyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCFrequencyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "CCFrequencyRuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCCFrequencyRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCFrequencyRulesResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCCFrequencyRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCCFrequencyRulesResponseParams `json:"Response"`
}

func (r *DeleteCCFrequencyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCFrequencyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCSelfDefinePolicyRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 策略ID
	SetId *string `json:"SetId,omitempty" name:"SetId"`
}

type DeleteCCSelfDefinePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 策略ID
	SetId *string `json:"SetId,omitempty" name:"SetId"`
}

func (r *DeleteCCSelfDefinePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCSelfDefinePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "SetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCCSelfDefinePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCSelfDefinePolicyResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCCSelfDefinePolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCCSelfDefinePolicyResponseParams `json:"Response"`
}

func (r *DeleteCCSelfDefinePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCSelfDefinePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDDoSPolicyCaseRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 策略场景ID
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
}

type DeleteDDoSPolicyCaseRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 策略场景ID
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
}

func (r *DeleteDDoSPolicyCaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSPolicyCaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "SceneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDDoSPolicyCaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDDoSPolicyCaseResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDDoSPolicyCaseResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDDoSPolicyCaseResponseParams `json:"Response"`
}

func (r *DeleteDDoSPolicyCaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSPolicyCaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDDoSPolicyRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

type DeleteDDoSPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DeleteDDoSPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDDoSPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDDoSPolicyResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDDoSPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDDoSPolicyResponseParams `json:"Response"`
}

func (r *DeleteDDoSPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL4RulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID列表
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

type DeleteL4RulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID列表
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

func (r *DeleteL4RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL4RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteL4RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL4RulesResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteL4RulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteL4RulesResponseParams `json:"Response"`
}

func (r *DeleteL4RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL4RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL7RulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID列表
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

type DeleteL7RulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID列表
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

func (r *DeleteL7RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL7RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteL7RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL7RulesResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteL7RulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteL7RulesResponseParams `json:"Response"`
}

func (r *DeleteL7RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL7RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNewL4RulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 删除接口结构体
	Rule []*L4DelRule `json:"Rule,omitempty" name:"Rule"`
}

type DeleteNewL4RulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 删除接口结构体
	Rule []*L4DelRule `json:"Rule,omitempty" name:"Rule"`
}

func (r *DeleteNewL4RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNewL4RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNewL4RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNewL4RulesResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteNewL4RulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNewL4RulesResponseParams `json:"Response"`
}

func (r *DeleteNewL4RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNewL4RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNewL7RulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP)
	Business *string `json:"Business,omitempty" name:"Business"`

	// 删除规则列表
	Rule []*L4DelRule `json:"Rule,omitempty" name:"Rule"`
}

type DeleteNewL7RulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP)
	Business *string `json:"Business,omitempty" name:"Business"`

	// 删除规则列表
	Rule []*L4DelRule `json:"Rule,omitempty" name:"Rule"`
}

func (r *DeleteNewL7RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNewL7RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNewL7RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNewL7RulesResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteNewL7RulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNewL7RulesResponseParams `json:"Response"`
}

func (r *DeleteNewL7RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNewL7RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActionLogRequestParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 搜索值，只支持资源ID或用户UIN
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeActionLogRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 搜索值，只支持资源ID或用户UIN
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeActionLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActionLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Business")
	delete(f, "Filter")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeActionLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActionLogResponseParams struct {
	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 记录数组
	Data []*KeyValueRecord `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeActionLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeActionLogResponseParams `json:"Response"`
}

func (r *DescribeActionLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActionLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBGPIPL7RuleMaxCntRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DescribeBGPIPL7RuleMaxCntRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeBGPIPL7RuleMaxCntRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBGPIPL7RuleMaxCntRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBGPIPL7RuleMaxCntRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBGPIPL7RuleMaxCntResponseParams struct {
	// 高防IP最多可添加的7层规则数量
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBGPIPL7RuleMaxCntResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBGPIPL7RuleMaxCntResponseParams `json:"Response"`
}

func (r *DescribeBGPIPL7RuleMaxCntResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBGPIPL7RuleMaxCntResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaradDataRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 指标名，取值：
	// connum表示TCP活跃连接数；
	// new_conn表示新建TCP连接数；
	// inactive_conn表示非活跃连接数;
	// intraffic表示入流量；
	// outtraffic表示出流量；
	// alltraffic表示出流量和入流量之和；
	// inpkg表示入包速率；
	// outpkg表示出包速率；
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 统计时间粒度，单位秒（300表示5分钟；3600表示小时；86400表示天）
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间，秒部分保持为0，分钟部分为5的倍数
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间，秒部分保持为0，分钟部分为5的倍数
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计方式，取值：
	// max表示最大值；
	// min表示最小值；
	// avg表示均值；
	Statistics *string `json:"Statistics,omitempty" name:"Statistics"`

	// 协议端口数组
	ProtocolPort []*ProtocolPort `json:"ProtocolPort,omitempty" name:"ProtocolPort"`

	// 资源实例下的IP，只有当Business=net(高防IP专业版)时才必须填写资源的一个IP（因为高防IP专业版资源实例有多个IP，才需要指定）；
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type DescribeBaradDataRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 指标名，取值：
	// connum表示TCP活跃连接数；
	// new_conn表示新建TCP连接数；
	// inactive_conn表示非活跃连接数;
	// intraffic表示入流量；
	// outtraffic表示出流量；
	// alltraffic表示出流量和入流量之和；
	// inpkg表示入包速率；
	// outpkg表示出包速率；
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 统计时间粒度，单位秒（300表示5分钟；3600表示小时；86400表示天）
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间，秒部分保持为0，分钟部分为5的倍数
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间，秒部分保持为0，分钟部分为5的倍数
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计方式，取值：
	// max表示最大值；
	// min表示最小值；
	// avg表示均值；
	Statistics *string `json:"Statistics,omitempty" name:"Statistics"`

	// 协议端口数组
	ProtocolPort []*ProtocolPort `json:"ProtocolPort,omitempty" name:"ProtocolPort"`

	// 资源实例下的IP，只有当Business=net(高防IP专业版)时才必须填写资源的一个IP（因为高防IP专业版资源实例有多个IP，才需要指定）；
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

func (r *DescribeBaradDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaradDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "MetricName")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Statistics")
	delete(f, "ProtocolPort")
	delete(f, "Ip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaradDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaradDataResponseParams struct {
	// 返回指标的值
	DataList []*BaradData `json:"DataList,omitempty" name:"DataList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBaradDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaradDataResponseParams `json:"Response"`
}

func (r *DescribeBaradDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaradDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBasicCCThresholdRequestParams struct {
	// 查询的IP地址，取值如：1.1.1.1
	BasicIp *string `json:"BasicIp,omitempty" name:"BasicIp"`

	// 查询IP所属地域，取值如：gz、bj、sh、hk等地域缩写
	BasicRegion *string `json:"BasicRegion,omitempty" name:"BasicRegion"`

	// 专区类型，取值如：公有云专区：public，黑石专区：bm, NAT服务器专区：nat，互联网通道：channel。
	BasicBizType *string `json:"BasicBizType,omitempty" name:"BasicBizType"`

	// 设备类型，取值如：服务器：cvm，公有云负载均衡：clb，黑石负载均衡：lb，NAT服务器：nat，互联网通道：channel.
	BasicDeviceType *string `json:"BasicDeviceType,omitempty" name:"BasicDeviceType"`

	// 可选，IPInstance Nat 网关（如果查询的设备类型是NAT服务器，需要传此参数，通过nat资源查询接口获取）
	BasicIpInstance *string `json:"BasicIpInstance,omitempty" name:"BasicIpInstance"`

	// 可选，运营商线路（如果查询的设备类型是NAT服务器，需要传此参数为5）
	BasicIspCode *uint64 `json:"BasicIspCode,omitempty" name:"BasicIspCode"`
}

type DescribeBasicCCThresholdRequest struct {
	*tchttp.BaseRequest
	
	// 查询的IP地址，取值如：1.1.1.1
	BasicIp *string `json:"BasicIp,omitempty" name:"BasicIp"`

	// 查询IP所属地域，取值如：gz、bj、sh、hk等地域缩写
	BasicRegion *string `json:"BasicRegion,omitempty" name:"BasicRegion"`

	// 专区类型，取值如：公有云专区：public，黑石专区：bm, NAT服务器专区：nat，互联网通道：channel。
	BasicBizType *string `json:"BasicBizType,omitempty" name:"BasicBizType"`

	// 设备类型，取值如：服务器：cvm，公有云负载均衡：clb，黑石负载均衡：lb，NAT服务器：nat，互联网通道：channel.
	BasicDeviceType *string `json:"BasicDeviceType,omitempty" name:"BasicDeviceType"`

	// 可选，IPInstance Nat 网关（如果查询的设备类型是NAT服务器，需要传此参数，通过nat资源查询接口获取）
	BasicIpInstance *string `json:"BasicIpInstance,omitempty" name:"BasicIpInstance"`

	// 可选，运营商线路（如果查询的设备类型是NAT服务器，需要传此参数为5）
	BasicIspCode *uint64 `json:"BasicIspCode,omitempty" name:"BasicIspCode"`
}

func (r *DescribeBasicCCThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicCCThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BasicIp")
	delete(f, "BasicRegion")
	delete(f, "BasicBizType")
	delete(f, "BasicDeviceType")
	delete(f, "BasicIpInstance")
	delete(f, "BasicIspCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBasicCCThresholdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBasicCCThresholdResponseParams struct {
	// CC启动开关（0:关闭；1:开启）
	CCEnable *uint64 `json:"CCEnable,omitempty" name:"CCEnable"`

	// CC防护阈值
	CCThreshold *uint64 `json:"CCThreshold,omitempty" name:"CCThreshold"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBasicCCThresholdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBasicCCThresholdResponseParams `json:"Response"`
}

func (r *DescribeBasicCCThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicCCThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBasicDeviceThresholdRequestParams struct {
	// 查询的IP地址，取值如：1.1.1.1
	BasicIp *string `json:"BasicIp,omitempty" name:"BasicIp"`

	// 查询IP所属地域，取值如：gz、bj、sh、hk等地域缩写
	BasicRegion *string `json:"BasicRegion,omitempty" name:"BasicRegion"`

	// 专区类型，取值如：公有云专区：public，黑石专区：bm, NAT服务器专区：nat，互联网通道：channel。
	BasicBizType *string `json:"BasicBizType,omitempty" name:"BasicBizType"`

	// 设备类型，取值如：服务器：cvm，公有云负载均衡：clb，黑石负载均衡：lb，NAT服务器：nat，互联网通道：channel.
	BasicDeviceType *string `json:"BasicDeviceType,omitempty" name:"BasicDeviceType"`

	// 有效性检查，取值为1
	BasicCheckFlag *uint64 `json:"BasicCheckFlag,omitempty" name:"BasicCheckFlag"`

	// 可选，IPInstance Nat 网关（如果查询的设备类型是NAT服务器，需要传此参数，通过nat资源查询接口获取）
	BasicIpInstance *string `json:"BasicIpInstance,omitempty" name:"BasicIpInstance"`

	// 可选，运营商线路（如果查询的设备类型是NAT服务器，需要传此参数为5）
	BasicIspCode *uint64 `json:"BasicIspCode,omitempty" name:"BasicIspCode"`
}

type DescribeBasicDeviceThresholdRequest struct {
	*tchttp.BaseRequest
	
	// 查询的IP地址，取值如：1.1.1.1
	BasicIp *string `json:"BasicIp,omitempty" name:"BasicIp"`

	// 查询IP所属地域，取值如：gz、bj、sh、hk等地域缩写
	BasicRegion *string `json:"BasicRegion,omitempty" name:"BasicRegion"`

	// 专区类型，取值如：公有云专区：public，黑石专区：bm, NAT服务器专区：nat，互联网通道：channel。
	BasicBizType *string `json:"BasicBizType,omitempty" name:"BasicBizType"`

	// 设备类型，取值如：服务器：cvm，公有云负载均衡：clb，黑石负载均衡：lb，NAT服务器：nat，互联网通道：channel.
	BasicDeviceType *string `json:"BasicDeviceType,omitempty" name:"BasicDeviceType"`

	// 有效性检查，取值为1
	BasicCheckFlag *uint64 `json:"BasicCheckFlag,omitempty" name:"BasicCheckFlag"`

	// 可选，IPInstance Nat 网关（如果查询的设备类型是NAT服务器，需要传此参数，通过nat资源查询接口获取）
	BasicIpInstance *string `json:"BasicIpInstance,omitempty" name:"BasicIpInstance"`

	// 可选，运营商线路（如果查询的设备类型是NAT服务器，需要传此参数为5）
	BasicIspCode *uint64 `json:"BasicIspCode,omitempty" name:"BasicIspCode"`
}

func (r *DescribeBasicDeviceThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicDeviceThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BasicIp")
	delete(f, "BasicRegion")
	delete(f, "BasicBizType")
	delete(f, "BasicDeviceType")
	delete(f, "BasicCheckFlag")
	delete(f, "BasicIpInstance")
	delete(f, "BasicIspCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBasicDeviceThresholdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBasicDeviceThresholdResponseParams struct {
	// 返回黑洞封堵值
	Threshold *uint64 `json:"Threshold,omitempty" name:"Threshold"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBasicDeviceThresholdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBasicDeviceThresholdResponseParams `json:"Response"`
}

func (r *DescribeBasicDeviceThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicDeviceThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBizHttpStatusRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源Id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 统计周期，可取值300，1800，3600， 21600，86400，单位秒
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计方式，仅支持sum
	Statistics *string `json:"Statistics,omitempty" name:"Statistics"`

	// 协议及端口列表，协议可取值TCP, UDP, HTTP, HTTPS，仅统计纬度为连接数时有效
	ProtoInfo []*ProtocolPort `json:"ProtoInfo,omitempty" name:"ProtoInfo"`

	// 特定域名查询
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type DescribeBizHttpStatusRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源Id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 统计周期，可取值300，1800，3600， 21600，86400，单位秒
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计方式，仅支持sum
	Statistics *string `json:"Statistics,omitempty" name:"Statistics"`

	// 协议及端口列表，协议可取值TCP, UDP, HTTP, HTTPS，仅统计纬度为连接数时有效
	ProtoInfo []*ProtocolPort `json:"ProtoInfo,omitempty" name:"ProtoInfo"`

	// 特定域名查询
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBizHttpStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBizHttpStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Statistics")
	delete(f, "ProtoInfo")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBizHttpStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBizHttpStatusResponseParams struct {
	// 业务流量http状态码统计数据
	HttpStatusMap *HttpStatusMap `json:"HttpStatusMap,omitempty" name:"HttpStatusMap"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBizHttpStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBizHttpStatusResponseParams `json:"Response"`
}

func (r *DescribeBizHttpStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBizHttpStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBizTrendRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 统计周期，可取值300，1800，3600，21600，86400，单位秒
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计方式，可取值max, min, avg, sum, 如统计纬度是流量速率或包量速率，仅可取值max
	Statistics *string `json:"Statistics,omitempty" name:"Statistics"`

	// 统计纬度，可取值connum, new_conn, inactive_conn, intraffic, outtraffic, inpkg, outpkg, qps
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 协议及端口列表，协议可取值TCP, UDP, HTTP, HTTPS，仅统计纬度为连接数时有效
	ProtoInfo []*ProtocolPort `json:"ProtoInfo,omitempty" name:"ProtoInfo"`

	// 统计纬度为qps时，可选特定域名查询
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type DescribeBizTrendRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 统计周期，可取值300，1800，3600，21600，86400，单位秒
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计方式，可取值max, min, avg, sum, 如统计纬度是流量速率或包量速率，仅可取值max
	Statistics *string `json:"Statistics,omitempty" name:"Statistics"`

	// 统计纬度，可取值connum, new_conn, inactive_conn, intraffic, outtraffic, inpkg, outpkg, qps
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 协议及端口列表，协议可取值TCP, UDP, HTTP, HTTPS，仅统计纬度为连接数时有效
	ProtoInfo []*ProtocolPort `json:"ProtoInfo,omitempty" name:"ProtoInfo"`

	// 统计纬度为qps时，可选特定域名查询
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBizTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBizTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Statistics")
	delete(f, "MetricName")
	delete(f, "ProtoInfo")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBizTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBizTrendResponseParams struct {
	// 曲线图各个时间点的值
	DataList []*float64 `json:"DataList,omitempty" name:"DataList"`

	// 统计纬度
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBizTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBizTrendResponseParams `json:"Response"`
}

func (r *DescribeBizTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBizTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCAlarmThresholdRequestParams struct {
	// 大禹子产品代号（shield表示棋牌；bgpip表示高防IP；bgp表示高防包；bgp-multip表示多ip高防包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID,字符串类型
	RsId *string `json:"RsId,omitempty" name:"RsId"`
}

type DescribeCCAlarmThresholdRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（shield表示棋牌；bgpip表示高防IP；bgp表示高防包；bgp-multip表示多ip高防包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID,字符串类型
	RsId *string `json:"RsId,omitempty" name:"RsId"`
}

func (r *DescribeCCAlarmThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCAlarmThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "RsId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCAlarmThresholdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCAlarmThresholdResponseParams struct {
	// CC告警阈值
	CCAlarmThreshold *CCAlarmThreshold `json:"CCAlarmThreshold,omitempty" name:"CCAlarmThreshold"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCCAlarmThresholdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCAlarmThresholdResponseParams `json:"Response"`
}

func (r *DescribeCCAlarmThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCAlarmThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCEvListRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源实例的IP，当business不为basic时，如果IpList不为空则Id也必须不能为空；
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeCCEvListRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源实例的IP，当business不为basic时，如果IpList不为空则Id也必须不能为空；
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCCEvListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCEvListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Id")
	delete(f, "IpList")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCEvListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCEvListResponseParams struct {
	// 大禹子产品代号（shield表示棋牌盾；bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源实例的IP列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// CC攻击事件列表
	Data []*CCEventRecord `json:"Data,omitempty" name:"Data"`

	// 总记录数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCCEvListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCEvListResponseParams `json:"Response"`
}

func (r *DescribeCCEvListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCEvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCFrequencyRulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 7层转发规则ID（通过获取7层转发规则接口可以获取规则ID）；当填写时表示获取转发规则的访问频率控制规则；
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type DescribeCCFrequencyRulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 7层转发规则ID（通过获取7层转发规则接口可以获取规则ID）；当填写时表示获取转发规则的访问频率控制规则；
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DescribeCCFrequencyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCFrequencyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCFrequencyRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCFrequencyRulesResponseParams struct {
	// 访问频率控制规则列表
	CCFrequencyRuleList []*CCFrequencyRule `json:"CCFrequencyRuleList,omitempty" name:"CCFrequencyRuleList"`

	// 访问频率控制规则开关状态，取值[on(开启)，off(关闭)]
	CCFrequencyRuleStatus *string `json:"CCFrequencyRuleStatus,omitempty" name:"CCFrequencyRuleStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCCFrequencyRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCFrequencyRulesResponseParams `json:"Response"`
}

func (r *DescribeCCFrequencyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCFrequencyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCIpAllowDenyRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 黑或白名单，取值[white(白名单)，black(黑名单)]
	// 注意：此数组只能有一个值，不能同时获取黑名单和白名单
	Type []*string `json:"Type,omitempty" name:"Type"`

	// 分页参数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 可选，代表HTTP协议或HTTPS协议的CC防护，取值[http（HTTP协议的CC防护），https（HTTPS协议的CC防护）]；
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type DescribeCCIpAllowDenyRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 黑或白名单，取值[white(白名单)，black(黑名单)]
	// 注意：此数组只能有一个值，不能同时获取黑名单和白名单
	Type []*string `json:"Type,omitempty" name:"Type"`

	// 分页参数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 可选，代表HTTP协议或HTTPS协议的CC防护，取值[http（HTTP协议的CC防护），https（HTTPS协议的CC防护）]；
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *DescribeCCIpAllowDenyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCIpAllowDenyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Type")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCIpAllowDenyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCIpAllowDenyResponseParams struct {
	// 该字段被RecordList字段替代了，请不要使用
	Data []*KeyValue `json:"Data,omitempty" name:"Data"`

	// 记录数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 返回黑/白名单的记录，
	// "Key":"ip"时，"Value":值表示ip;
	// "Key":"domain"时， "Value":值表示域名;
	// "Key":"type"时，"Value":值表示黑白名单类型(white为白名单，block为黑名单);
	// "Key":"protocol"时，"Value":值表示CC防护的协议(http或https);
	RecordList []*KeyValueRecord `json:"RecordList,omitempty" name:"RecordList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCCIpAllowDenyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCIpAllowDenyResponseParams `json:"Response"`
}

func (r *DescribeCCIpAllowDenyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCIpAllowDenyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCSelfDefinePolicyRequestParams struct {
	// 大禹子产品代号（bgp高防包；bgp-multip共享包）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 拉取的条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeCCSelfDefinePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgp高防包；bgp-multip共享包）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 拉取的条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCCSelfDefinePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCSelfDefinePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCSelfDefinePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCSelfDefinePolicyResponseParams struct {
	// 自定义规则总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 策略列表
	Policys []*CCPolicy `json:"Policys,omitempty" name:"Policys"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCCSelfDefinePolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCSelfDefinePolicyResponseParams `json:"Response"`
}

func (r *DescribeCCSelfDefinePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCSelfDefinePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCTrendRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 指标，取值[inqps(总请求峰值，dropqps(攻击请求峰值))]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 资源实例ID，当Business为basic时，此字段不用填写（因为基础防护没有资源实例）
	Id *string `json:"Id,omitempty" name:"Id"`

	// 域名，可选
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type DescribeCCTrendRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 指标，取值[inqps(总请求峰值，dropqps(攻击请求峰值))]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 资源实例ID，当Business为basic时，此字段不用填写（因为基础防护没有资源实例）
	Id *string `json:"Id,omitempty" name:"Id"`

	// 域名，可选
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeCCTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Ip")
	delete(f, "MetricName")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Id")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCTrendResponseParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 指标，取值[inqps(总请求峰值，dropqps(攻击请求峰值))]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 值数组
	Data []*uint64 `json:"Data,omitempty" name:"Data"`

	// 值个数
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCCTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCTrendResponseParams `json:"Response"`
}

func (r *DescribeCCTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCUrlAllowRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 黑或白名单，取值[white(白名单)]，目前只支持白名单
	// 注意：此数组只能有一个值，且只能为white
	Type []*string `json:"Type,omitempty" name:"Type"`

	// 分页参数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 可选，代表HTTP协议或HTTPS协议的CC防护，取值[http（HTTP协议的CC防护），https（HTTPS协议的CC防护）]；
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type DescribeCCUrlAllowRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 黑或白名单，取值[white(白名单)]，目前只支持白名单
	// 注意：此数组只能有一个值，且只能为white
	Type []*string `json:"Type,omitempty" name:"Type"`

	// 分页参数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 可选，代表HTTP协议或HTTPS协议的CC防护，取值[http（HTTP协议的CC防护），https（HTTPS协议的CC防护）]；
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *DescribeCCUrlAllowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCUrlAllowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Type")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCUrlAllowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCUrlAllowResponseParams struct {
	// 该字段被RecordList字段替代了，请不要使用
	Data []*KeyValue `json:"Data,omitempty" name:"Data"`

	// 记录总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 返回黑/白名单的记录，
	// "Key":"url"时，"Value":值表示URL;
	// "Key":"domain"时， "Value":值表示域名;
	// "Key":"type"时，"Value":值表示黑白名单类型(white为白名单，block为黑名单);
	// "Key":"protocol"时，"Value":值表示CC的防护类型(HTTP防护或HTTPS域名防护);
	RecordList []*KeyValueRecord `json:"RecordList,omitempty" name:"RecordList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCCUrlAllowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCUrlAllowResponseParams `json:"Response"`
}

func (r *DescribeCCUrlAllowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCUrlAllowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAlarmThresholdRequestParams struct {
	// 大禹子产品代号（shield表示棋牌；bgpip表示高防IP；bgp表示高防包；bgp-multip表示多ip高防包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID,字符串类型
	RsId *string `json:"RsId,omitempty" name:"RsId"`
}

type DescribeDDoSAlarmThresholdRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（shield表示棋牌；bgpip表示高防IP；bgp表示高防包；bgp-multip表示多ip高防包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID,字符串类型
	RsId *string `json:"RsId,omitempty" name:"RsId"`
}

func (r *DescribeDDoSAlarmThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAlarmThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "RsId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAlarmThresholdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAlarmThresholdResponseParams struct {
	// DDoS告警阈值
	DDoSAlarmThreshold *DDoSAlarmThreshold `json:"DDoSAlarmThreshold,omitempty" name:"DDoSAlarmThreshold"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSAlarmThresholdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSAlarmThresholdResponseParams `json:"Response"`
}

func (r *DescribeDDoSAlarmThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAlarmThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackIPRegionMapRequestParams struct {
	// 大禹子产品代号（shield表示棋牌；bgpip表示高防IP；bgp表示高防包；bgp-multip表示多ip高防包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间，最大可统计的时间范围是半年；
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指定资源的特定IP的攻击源，可选
	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

type DescribeDDoSAttackIPRegionMapRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（shield表示棋牌；bgpip表示高防IP；bgp表示高防包；bgp-multip表示多ip高防包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间，最大可统计的时间范围是半年；
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指定资源的特定IP的攻击源，可选
	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

func (r *DescribeDDoSAttackIPRegionMapRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackIPRegionMapRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "IpList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackIPRegionMapRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackIPRegionMapResponseParams struct {
	// 全球地域分布数据
	NationCount []*KeyValueRecord `json:"NationCount,omitempty" name:"NationCount"`

	// 国内省份地域分布数据
	ProvinceCount []*KeyValueRecord `json:"ProvinceCount,omitempty" name:"ProvinceCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSAttackIPRegionMapResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSAttackIPRegionMapResponseParams `json:"Response"`
}

func (r *DescribeDDoSAttackIPRegionMapResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackIPRegionMapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackSourceRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 获取指定资源的特定ip的攻击源，可选
	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

type DescribeDDoSAttackSourceRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 获取指定资源的特定ip的攻击源，可选
	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

func (r *DescribeDDoSAttackSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "IpList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackSourceResponseParams struct {
	// 总攻击源条数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 攻击源列表
	AttackSourceList []*DDoSAttackSourceRecord `json:"AttackSourceList,omitempty" name:"AttackSourceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSAttackSourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSAttackSourceResponseParams `json:"Response"`
}

func (r *DescribeDDoSAttackSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSCountRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标，取值[traffic（攻击协议流量, 单位KB）, pkg（攻击协议报文数）, classnum（攻击事件次数）]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

type DescribeDDoSCountRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标，取值[traffic（攻击协议流量, 单位KB）, pkg（攻击协议报文数）, classnum（攻击事件次数）]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

func (r *DescribeDDoSCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Ip")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSCountResponseParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标，取值[traffic（攻击协议流量, 单位KB）, pkg（攻击协议报文数）, classnum（攻击事件次数）]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Key-Value值数组，Key说明如下，
	// 当MetricName为traffic时：
	// key为"TcpKBSum"，表示TCP报文流量，单位KB
	// key为"UdpKBSum"，表示UDP报文流量，单位KB
	// key为"IcmpKBSum"，表示ICMP报文流量，单位KB
	// key为"OtherKBSum"，表示其他报文流量，单位KB
	// 
	// 当MetricName为pkg时：
	// key为"TcpPacketSum"，表示TCP报文个数，单位个
	// key为"UdpPacketSum"，表示UDP报文个数，单位个
	// key为"IcmpPacketSum"，表示ICMP报文个数，单位个
	// key为"OtherPacketSum"，表示其他报文个数，单位个
	// 
	// 当MetricName为classnum时：
	// key的值表示攻击事件类型，其中Key为"UNKNOWNFLOOD"，表示未知的攻击事件
	Data []*KeyValue `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSCountResponseParams `json:"Response"`
}

func (r *DescribeDDoSCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSDefendStatusRequestParams struct {
	// 大禹子产品代号（basic表示基础防护；bgp表示独享包；bgp-multip表示共享包；bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID，只有当Business不是基础防护时才需要填写此字段；
	Id *string `json:"Id,omitempty" name:"Id"`

	// 基础防护的IP，只有当Business为基础防护时才需要填写此字段；
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 只有当Business为基础防护时才需要填写此字段，IP所属的产品类型，取值[public（CVM产品），bm（黑石产品），eni（弹性网卡），vpngw（VPN网关）， natgw（NAT网关），waf（Web应用安全产品），fpc（金融产品），gaap（GAAP产品）, other(托管IP)]
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// 只有当Business为基础防护时才需要填写此字段，IP所属的产品子类，取值[cvm（CVM），lb（负载均衡器），eni（弹性网卡），vpngw（VPN），natgw（NAT），waf（WAF），fpc（金融），gaap（GAAP），other（托管IP），eip（黑石弹性IP）]
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 只有当Business为基础防护时才需要填写此字段，IP所属的资源实例ID，当绑定新IP时必须填写此字段；例如是弹性网卡的IP，则InstanceId填写弹性网卡的ID(eni-*);
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 只有当Business为基础防护时才需要填写此字段，表示IP所属的地域，取值：
	// "bj":     华北地区(北京)
	// "cd":     西南地区(成都)
	// "cq":     西南地区(重庆)
	// "gz":     华南地区(广州)
	// "gzopen": 华南地区(广州Open)
	// "hk":     中国香港
	// "kr":     东南亚地区(首尔)
	// "sh":     华东地区(上海)
	// "shjr":   华东地区(上海金融)
	// "szjr":   华南地区(深圳金融)
	// "sg":     东南亚地区(新加坡)
	// "th":     东南亚地区(泰国)
	// "de":     欧洲地区(德国)
	// "usw":    美国西部（硅谷）
	// "ca":     北美地区(多伦多)
	// "jp":     日本
	// "hzec":   杭州
	// "in":     印度
	// "use":    美东地区（弗吉尼亚）
	// "ru":     俄罗斯
	// "tpe":    中国台湾
	// "nj":     南京
	IPRegion *string `json:"IPRegion,omitempty" name:"IPRegion"`
}

type DescribeDDoSDefendStatusRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（basic表示基础防护；bgp表示独享包；bgp-multip表示共享包；bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID，只有当Business不是基础防护时才需要填写此字段；
	Id *string `json:"Id,omitempty" name:"Id"`

	// 基础防护的IP，只有当Business为基础防护时才需要填写此字段；
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 只有当Business为基础防护时才需要填写此字段，IP所属的产品类型，取值[public（CVM产品），bm（黑石产品），eni（弹性网卡），vpngw（VPN网关）， natgw（NAT网关），waf（Web应用安全产品），fpc（金融产品），gaap（GAAP产品）, other(托管IP)]
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// 只有当Business为基础防护时才需要填写此字段，IP所属的产品子类，取值[cvm（CVM），lb（负载均衡器），eni（弹性网卡），vpngw（VPN），natgw（NAT），waf（WAF），fpc（金融），gaap（GAAP），other（托管IP），eip（黑石弹性IP）]
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 只有当Business为基础防护时才需要填写此字段，IP所属的资源实例ID，当绑定新IP时必须填写此字段；例如是弹性网卡的IP，则InstanceId填写弹性网卡的ID(eni-*);
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 只有当Business为基础防护时才需要填写此字段，表示IP所属的地域，取值：
	// "bj":     华北地区(北京)
	// "cd":     西南地区(成都)
	// "cq":     西南地区(重庆)
	// "gz":     华南地区(广州)
	// "gzopen": 华南地区(广州Open)
	// "hk":     中国香港
	// "kr":     东南亚地区(首尔)
	// "sh":     华东地区(上海)
	// "shjr":   华东地区(上海金融)
	// "szjr":   华南地区(深圳金融)
	// "sg":     东南亚地区(新加坡)
	// "th":     东南亚地区(泰国)
	// "de":     欧洲地区(德国)
	// "usw":    美国西部（硅谷）
	// "ca":     北美地区(多伦多)
	// "jp":     日本
	// "hzec":   杭州
	// "in":     印度
	// "use":    美东地区（弗吉尼亚）
	// "ru":     俄罗斯
	// "tpe":    中国台湾
	// "nj":     南京
	IPRegion *string `json:"IPRegion,omitempty" name:"IPRegion"`
}

func (r *DescribeDDoSDefendStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSDefendStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Ip")
	delete(f, "BizType")
	delete(f, "DeviceType")
	delete(f, "InstanceId")
	delete(f, "IPRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSDefendStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSDefendStatusResponseParams struct {
	// 防护状态，为0表示防护处于关闭状态，为1表示防护处于开启状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefendStatus *uint64 `json:"DefendStatus,omitempty" name:"DefendStatus"`

	// 防护临时关闭的过期时间，当防护状态为开启时此字段为空；
	// 注意：此字段可能返回 null，表示取不到有效值。
	UndefendExpire *string `json:"UndefendExpire,omitempty" name:"UndefendExpire"`

	// 控制台功能展示字段，为1表示控制台功能展示，为0表示控制台功能隐藏
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowFlag *uint64 `json:"ShowFlag,omitempty" name:"ShowFlag"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSDefendStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSDefendStatusResponseParams `json:"Response"`
}

func (r *DescribeDDoSDefendStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSDefendStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSEvInfoRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 攻击开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 攻击结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeDDoSEvInfoRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 攻击开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 攻击结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDDoSEvInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSEvInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Ip")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSEvInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSEvInfoResponseParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 攻击开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 攻击结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// TCP报文攻击包数
	TcpPacketSum *uint64 `json:"TcpPacketSum,omitempty" name:"TcpPacketSum"`

	// TCP报文攻击流量，单位KB
	TcpKBSum *uint64 `json:"TcpKBSum,omitempty" name:"TcpKBSum"`

	// UDP报文攻击包数
	UdpPacketSum *uint64 `json:"UdpPacketSum,omitempty" name:"UdpPacketSum"`

	// UDP报文攻击流量，单位KB
	UdpKBSum *uint64 `json:"UdpKBSum,omitempty" name:"UdpKBSum"`

	// ICMP报文攻击包数
	IcmpPacketSum *uint64 `json:"IcmpPacketSum,omitempty" name:"IcmpPacketSum"`

	// ICMP报文攻击流量，单位KB
	IcmpKBSum *uint64 `json:"IcmpKBSum,omitempty" name:"IcmpKBSum"`

	// 其他报文攻击包数
	OtherPacketSum *uint64 `json:"OtherPacketSum,omitempty" name:"OtherPacketSum"`

	// 其他报文攻击流量，单位KB
	OtherKBSum *uint64 `json:"OtherKBSum,omitempty" name:"OtherKBSum"`

	// 累计攻击流量，单位KB
	TotalTraffic *uint64 `json:"TotalTraffic,omitempty" name:"TotalTraffic"`

	// 攻击流量带宽峰值
	Mbps *uint64 `json:"Mbps,omitempty" name:"Mbps"`

	// 攻击包速率峰值
	Pps *uint64 `json:"Pps,omitempty" name:"Pps"`

	// PCAP文件下载链接
	PcapUrl []*string `json:"PcapUrl,omitempty" name:"PcapUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSEvInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSEvInfoResponseParams `json:"Response"`
}

func (r *DescribeDDoSEvInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSEvInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSEvListRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 资源实例ID，当Business为basic时，此字段不用填写（因为基础防护没有资源实例）
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源的IP
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// 是否超过弹性防护峰值，取值[yes(是)，no(否)]，填写空字符串时表示不进行过滤
	OverLoad *string `json:"OverLoad,omitempty" name:"OverLoad"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeDDoSEvListRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 资源实例ID，当Business为basic时，此字段不用填写（因为基础防护没有资源实例）
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源的IP
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// 是否超过弹性防护峰值，取值[yes(是)，no(否)]，填写空字符串时表示不进行过滤
	OverLoad *string `json:"OverLoad,omitempty" name:"OverLoad"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDDoSEvListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSEvListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Id")
	delete(f, "IpList")
	delete(f, "OverLoad")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSEvListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSEvListResponseParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源的IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// DDoS攻击事件列表
	Data []*DDoSEventRecord `json:"Data,omitempty" name:"Data"`

	// 总记录数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSEvListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSEvListResponseParams `json:"Response"`
}

func (r *DescribeDDoSEvListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSEvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSIpLogRequestParams struct {
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 攻击开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 攻击结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeDDoSIpLogRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 攻击开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 攻击结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDDoSIpLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSIpLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Ip")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSIpLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSIpLogResponseParams struct {
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 攻击开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 攻击结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// IP攻击日志，KeyValue数组，Key-Value取值说明：
	// Key为"LogTime"时，Value值为IP日志时间
	// Key为"LogMessage"时，Value值为Ip日志内容
	Data []*KeyValueRecord `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSIpLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSIpLogResponseParams `json:"Response"`
}

func (r *DescribeDDoSIpLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSIpLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSNetCountRequestParams struct {
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标，取值[traffic（攻击协议流量, 单位KB）, pkg（攻击协议报文数）, classnum（攻击事件次数）]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

type DescribeDDoSNetCountRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标，取值[traffic（攻击协议流量, 单位KB）, pkg（攻击协议报文数）, classnum（攻击事件次数）]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

func (r *DescribeDDoSNetCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSNetCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSNetCountResponseParams struct {
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标，取值[traffic（攻击协议流量, 单位KB）, pkg（攻击协议报文数）, classnum（攻击事件次数）]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Key-Value值数组，Key说明如下，
	// 当MetricName为traffic时：
	// key为"TcpKBSum"，表示TCP报文流量，单位KB
	// key为"UdpKBSum"，表示UDP报文流量，单位KB
	// key为"IcmpKBSum"，表示ICMP报文流量，单位KB
	// key为"OtherKBSum"，表示其他报文流量，单位KB
	// 
	// 当MetricName为pkg时：
	// key为"TcpPacketSum"，表示TCP报文个数，单位个
	// key为"UdpPacketSum"，表示UDP报文个数，单位个
	// key为"IcmpPacketSum"，表示ICMP报文个数，单位个
	// key为"OtherPacketSum"，表示其他报文个数，单位个
	// 
	// 当MetricName为classnum时：
	// key的值表示攻击事件类型，其中Key为"UNKNOWNFLOOD"，表示未知的攻击事件
	Data []*KeyValue `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSNetCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSNetCountResponseParams `json:"Response"`
}

func (r *DescribeDDoSNetCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSNetEvInfoRequestParams struct {
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 攻击开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 攻击结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeDDoSNetEvInfoRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 攻击开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 攻击结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDDoSNetEvInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetEvInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSNetEvInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSNetEvInfoResponseParams struct {
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 攻击开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 攻击结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// TCP报文攻击包数
	TcpPacketSum *uint64 `json:"TcpPacketSum,omitempty" name:"TcpPacketSum"`

	// TCP报文攻击流量，单位KB
	TcpKBSum *uint64 `json:"TcpKBSum,omitempty" name:"TcpKBSum"`

	// UDP报文攻击包数
	UdpPacketSum *uint64 `json:"UdpPacketSum,omitempty" name:"UdpPacketSum"`

	// UDP报文攻击流量，单位KB
	UdpKBSum *uint64 `json:"UdpKBSum,omitempty" name:"UdpKBSum"`

	// ICMP报文攻击包数
	IcmpPacketSum *uint64 `json:"IcmpPacketSum,omitempty" name:"IcmpPacketSum"`

	// ICMP报文攻击流量，单位KB
	IcmpKBSum *uint64 `json:"IcmpKBSum,omitempty" name:"IcmpKBSum"`

	// 其他报文攻击包数
	OtherPacketSum *uint64 `json:"OtherPacketSum,omitempty" name:"OtherPacketSum"`

	// 其他报文攻击流量，单位KB
	OtherKBSum *uint64 `json:"OtherKBSum,omitempty" name:"OtherKBSum"`

	// 累计攻击流量，单位KB
	TotalTraffic *uint64 `json:"TotalTraffic,omitempty" name:"TotalTraffic"`

	// 攻击流量带宽峰值
	Mbps *uint64 `json:"Mbps,omitempty" name:"Mbps"`

	// 攻击包速率峰值
	Pps *uint64 `json:"Pps,omitempty" name:"Pps"`

	// PCAP文件下载链接
	PcapUrl []*string `json:"PcapUrl,omitempty" name:"PcapUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSNetEvInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSNetEvInfoResponseParams `json:"Response"`
}

func (r *DescribeDDoSNetEvInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetEvInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSNetEvListRequestParams struct {
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeDDoSNetEvListRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDDoSNetEvListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetEvListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSNetEvListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSNetEvListResponseParams struct {
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// DDoS攻击事件列表
	Data []*DDoSEventRecord `json:"Data,omitempty" name:"Data"`

	// 总记录数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSNetEvListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSNetEvListResponseParams `json:"Response"`
}

func (r *DescribeDDoSNetEvListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetEvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSNetIpLogRequestParams struct {
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 攻击开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 攻击结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeDDoSNetIpLogRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 攻击开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 攻击结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDDoSNetIpLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetIpLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSNetIpLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSNetIpLogResponseParams struct {
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 攻击开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 攻击结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// IP攻击日志，KeyValue数组，Key-Value取值说明：
	// Key为"LogTime"时，Value值为IP日志时间
	// Key为"LogMessage"时，Value值为Ip日志内容
	Data []*KeyValueRecord `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSNetIpLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSNetIpLogResponseParams `json:"Response"`
}

func (r *DescribeDDoSNetIpLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetIpLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSNetTrendRequestParams struct {
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 指标，取值[bps(攻击流量带宽，pps(攻击包速率))]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeDDoSNetTrendRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 指标，取值[bps(攻击流量带宽，pps(攻击包速率))]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDDoSNetTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "MetricName")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSNetTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSNetTrendResponseParams struct {
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 指标，取值[bps(攻击流量带宽，pps(攻击包速率))]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 值数组
	Data []*uint64 `json:"Data,omitempty" name:"Data"`

	// 值个数
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSNetTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSNetTrendResponseParams `json:"Response"`
}

func (r *DescribeDDoSNetTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSPolicyRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 可选字段，资源ID，如果填写则表示该资源绑定的DDoS高级策略
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DescribeDDoSPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 可选字段，资源ID，如果填写则表示该资源绑定的DDoS高级策略
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeDDoSPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSPolicyResponseParams struct {
	// DDoS高级策略列表
	DDosPolicyList []*DDosPolicy `json:"DDosPolicyList,omitempty" name:"DDosPolicyList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSPolicyResponseParams `json:"Response"`
}

func (r *DescribeDDoSPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSTrendRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 指标，取值[bps(攻击流量带宽，pps(攻击包速率))]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 资源实例ID，当Business为basic时，此字段不用填写（因为基础防护没有资源实例）
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DescribeDDoSTrendRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 指标，取值[bps(攻击流量带宽，pps(攻击包速率))]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 资源实例ID，当Business为basic时，此字段不用填写（因为基础防护没有资源实例）
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeDDoSTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Ip")
	delete(f, "MetricName")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSTrendResponseParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 指标，取值[bps(攻击流量带宽，pps(攻击包速率))]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 值数组，攻击流量带宽单位为Mbps，包速率单位为pps
	Data []*uint64 `json:"Data,omitempty" name:"Data"`

	// 值个数
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSTrendResponseParams `json:"Response"`
}

func (r *DescribeDDoSTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSUsedStatisRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`
}

type DescribeDDoSUsedStatisRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`
}

func (r *DescribeDDoSUsedStatisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSUsedStatisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSUsedStatisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSUsedStatisResponseParams struct {
	// 字段值，如下：
	// Days：高防资源使用天数
	// Attacks：DDoS防护次数
	Data []*KeyValue `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSUsedStatisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSUsedStatisResponseParams `json:"Response"`
}

func (r *DescribeDDoSUsedStatisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSUsedStatisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPProductInfoRequestParams struct {
	// 大禹子产品代号（bgp表示独享包；bgp-multip表示共享包）
	Business *string `json:"Business,omitempty" name:"Business"`

	// IP列表
	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

type DescribeIPProductInfoRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgp表示独享包；bgp-multip表示共享包）
	Business *string `json:"Business,omitempty" name:"Business"`

	// IP列表
	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

func (r *DescribeIPProductInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPProductInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "IpList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIPProductInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPProductInfoResponseParams struct {
	// 云产品信息列表，如果没有查询到则返回空数组，值说明如下：
	// Key为ProductName时，value表示云产品实例的名称；
	// Key为ProductInstanceId时，value表示云产品实例的ID；
	// Key为ProductType时，value表示的是云产品的类型（cvm表示云主机、clb表示负载均衡）;
	// Key为IP时，value表示云产品实例的IP；
	Data []*KeyValueRecord `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIPProductInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIPProductInfoResponseParams `json:"Response"`
}

func (r *DescribeIPProductInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPProductInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInsurePacksRequestParams struct {
	// 可选字段，保险包套餐ID，当要获取指定ID（例如insure-000000xe）的保险包套餐时请填写此字段；
	IdList []*string `json:"IdList,omitempty" name:"IdList"`
}

type DescribeInsurePacksRequest struct {
	*tchttp.BaseRequest
	
	// 可选字段，保险包套餐ID，当要获取指定ID（例如insure-000000xe）的保险包套餐时请填写此字段；
	IdList []*string `json:"IdList,omitempty" name:"IdList"`
}

func (r *DescribeInsurePacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInsurePacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInsurePacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInsurePacksResponseParams struct {
	// 保险包套餐列表
	InsurePacks []*KeyValueRecord `json:"InsurePacks,omitempty" name:"InsurePacks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInsurePacksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInsurePacksResponseParams `json:"Response"`
}

func (r *DescribeInsurePacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInsurePacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpBlockListRequestParams struct {

}

type DescribeIpBlockListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeIpBlockListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpBlockListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpBlockListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpBlockListResponseParams struct {
	// IP封堵列表
	List []*IpBlockData `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIpBlockListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpBlockListResponseParams `json:"Response"`
}

func (r *DescribeIpBlockListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpBlockListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpUnBlockListRequestParams struct {
	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// IP（不为空时，进行IP过滤）
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 分页参数（不为空时，进行分页查询），此字段后面会弃用，请用Limit和Offset字段代替；
	Paging *Paging `json:"Paging,omitempty" name:"Paging"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeIpUnBlockListRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// IP（不为空时，进行IP过滤）
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 分页参数（不为空时，进行分页查询），此字段后面会弃用，请用Limit和Offset字段代替；
	Paging *Paging `json:"Paging,omitempty" name:"Paging"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeIpUnBlockListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpUnBlockListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Ip")
	delete(f, "Paging")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpUnBlockListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpUnBlockListResponseParams struct {
	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// IP解封记录
	List []*IpUnBlockData `json:"List,omitempty" name:"List"`

	// 总记录数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIpUnBlockListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpUnBlockListResponseParams `json:"Response"`
}

func (r *DescribeIpUnBlockListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpUnBlockListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL4HealthConfigRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID数组，当导出所有规则的健康检查配置则不填或填空数组；
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

type DescribeL4HealthConfigRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID数组，当导出所有规则的健康检查配置则不填或填空数组；
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

func (r *DescribeL4HealthConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4HealthConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL4HealthConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL4HealthConfigResponseParams struct {
	// 四层健康检查配置数组
	HealthConfig []*L4HealthConfig `json:"HealthConfig,omitempty" name:"HealthConfig"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeL4HealthConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL4HealthConfigResponseParams `json:"Response"`
}

func (r *DescribeL4HealthConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4HealthConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL4RulesErrHealthRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DescribeL4RulesErrHealthRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeL4RulesErrHealthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4RulesErrHealthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL4RulesErrHealthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL4RulesErrHealthResponseParams struct {
	// 异常规则的总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 异常规则列表，返回值说明: Key值为规则ID，Value值为异常IP，多个IP用","分割
	ErrHealths []*KeyValue `json:"ErrHealths,omitempty" name:"ErrHealths"`

	// 异常规则列表(提供更多的错误相关信息)，返回值说明:
	// Key值为RuleId时，Value值为规则ID；
	// Key值为Protocol时，Value值为规则的转发协议；
	// Key值为VirtualPort时，Value值为规则的转发端口；
	// Key值为ErrMessage时，Value值为健康检查异常信息；
	// 健康检查异常信息的格式为"SourceIp:1.1.1.1|SourcePort:1234|AbnormalStatTime:1570689065|AbnormalReason:connection time out|Interval:20|CheckNum:6|FailNum:6" 多个源IP的错误信息用，分割,
	// SourceIp表示源站IP，SourcePort表示源站端口，AbnormalStatTime表示异常时间，AbnormalReason表示异常原因，Interval表示检查周期，CheckNum表示检查次数，FailNum表示失败次数；
	ExtErrHealths []*KeyValueRecord `json:"ExtErrHealths,omitempty" name:"ExtErrHealths"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeL4RulesErrHealthResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL4RulesErrHealthResponseParams `json:"Response"`
}

func (r *DescribeL4RulesErrHealthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4RulesErrHealthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7HealthConfigRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID数组，当导出所有规则的健康检查配置则不填或填空数组；
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

type DescribeL7HealthConfigRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID数组，当导出所有规则的健康检查配置则不填或填空数组；
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

func (r *DescribeL7HealthConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7HealthConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL7HealthConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7HealthConfigResponseParams struct {
	// 七层健康检查配置数组
	HealthConfig []*L7HealthConfig `json:"HealthConfig,omitempty" name:"HealthConfig"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeL7HealthConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL7HealthConfigResponseParams `json:"Response"`
}

func (r *DescribeL7HealthConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7HealthConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewL4RulesErrHealthRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 规则ID列表
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

type DescribeNewL4RulesErrHealthRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 规则ID列表
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

func (r *DescribeNewL4RulesErrHealthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewL4RulesErrHealthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "RuleIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNewL4RulesErrHealthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewL4RulesErrHealthResponseParams struct {
	// 异常规则的总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 异常规则列表，返回值说明: Key值为规则ID，Value值为异常IP，多个IP用","分割
	ErrHealths []*KeyValue `json:"ErrHealths,omitempty" name:"ErrHealths"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNewL4RulesErrHealthResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNewL4RulesErrHealthResponseParams `json:"Response"`
}

func (r *DescribeNewL4RulesErrHealthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewL4RulesErrHealthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewL4RulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 指定IP查询
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 指定高防IP端口查询
	VirtualPort *uint64 `json:"VirtualPort,omitempty" name:"VirtualPort"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeNewL4RulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 指定IP查询
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 指定高防IP端口查询
	VirtualPort *uint64 `json:"VirtualPort,omitempty" name:"VirtualPort"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeNewL4RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewL4RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Ip")
	delete(f, "VirtualPort")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNewL4RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewL4RulesResponseParams struct {
	// 转发规则列表
	Rules []*NewL4RuleEntry `json:"Rules,omitempty" name:"Rules"`

	// 总规则数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 四层健康检查配置列表
	Healths []*L4RuleHealth `json:"Healths,omitempty" name:"Healths"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNewL4RulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNewL4RulesResponseParams `json:"Response"`
}

func (r *DescribeNewL4RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewL4RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewL7RulesErrHealthRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP)
	Business *string `json:"Business,omitempty" name:"Business"`

	// 规则Id列表
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

type DescribeNewL7RulesErrHealthRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP)
	Business *string `json:"Business,omitempty" name:"Business"`

	// 规则Id列表
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

func (r *DescribeNewL7RulesErrHealthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewL7RulesErrHealthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "RuleIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNewL7RulesErrHealthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewL7RulesErrHealthResponseParams struct {
	// 异常规则的总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 异常规则列表，返回值说明: Key值为规则ID，Value值为异常IP及错误信息，多个IP用","分割
	ErrHealths []*KeyValue `json:"ErrHealths,omitempty" name:"ErrHealths"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNewL7RulesErrHealthResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNewL7RulesErrHealthResponseParams `json:"Response"`
}

func (r *DescribeNewL7RulesErrHealthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewL7RulesErrHealthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePackIndexRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示高防包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`
}

type DescribePackIndexRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示高防包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`
}

func (r *DescribePackIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePackIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePackIndexResponseParams struct {
	// 字段值，如下：
	// TotalPackCount：资源数
	// AttackPackCount：清洗中的资源数
	// BlockPackCount：封堵中的资源数
	// ExpiredPackCount：过期的资源数
	// ExpireingPackCount：即将过期的资源数
	// IsolatePackCount：隔离中的资源数
	Data []*KeyValue `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePackIndexResponse struct {
	*tchttp.BaseResponse
	Response *DescribePackIndexResponseParams `json:"Response"`
}

func (r *DescribePackIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePcapRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 攻击事件的开始时间，格式为"2018-08-28 07:00:00"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 攻击事件的结束时间，格式为"2018-08-28 07:02:00"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 资源的IP，只有当Business为net时才需要填写资源实例下的IP；
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type DescribePcapRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 攻击事件的开始时间，格式为"2018-08-28 07:00:00"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 攻击事件的结束时间，格式为"2018-08-28 07:02:00"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 资源的IP，只有当Business为net时才需要填写资源实例下的IP；
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

func (r *DescribePcapRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePcapRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Ip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePcapRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePcapResponseParams struct {
	// pcap包的下载链接列表，无pcap包时为空数组；
	PcapUrlList []*string `json:"PcapUrlList,omitempty" name:"PcapUrlList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePcapResponse struct {
	*tchttp.BaseResponse
	Response *DescribePcapResponseParams `json:"Response"`
}

func (r *DescribePcapResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePcapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyCaseRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 策略场景ID
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
}

type DescribePolicyCaseRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 策略场景ID
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
}

func (r *DescribePolicyCaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyCaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "SceneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePolicyCaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyCaseResponseParams struct {
	// 策略场景列表
	CaseList []*KeyValueRecord `json:"CaseList,omitempty" name:"CaseList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePolicyCaseResponse struct {
	*tchttp.BaseResponse
	Response *DescribePolicyCaseResponseParams `json:"Response"`
}

func (r *DescribePolicyCaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyCaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResIpListRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID, 如果不填，则获取用户所有资源的IP
	IdList []*string `json:"IdList,omitempty" name:"IdList"`
}

type DescribeResIpListRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID, 如果不填，则获取用户所有资源的IP
	IdList []*string `json:"IdList,omitempty" name:"IdList"`
}

func (r *DescribeResIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "IdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResIpListResponseParams struct {
	// 资源的IP列表
	Resource []*ResourceIp `json:"Resource,omitempty" name:"Resource"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResIpListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResIpListResponseParams `json:"Response"`
}

func (r *DescribeResIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceListRequestParams struct {
	// 大禹子产品代号（bgp表示独享包；bgp-multip表示共享包；bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 地域码搜索，可选，当不指定地域时空数组，当指定地域时，填地域码。例如：["gz", "sh"]
	RegionList []*string `json:"RegionList,omitempty" name:"RegionList"`

	// 线路搜索，可选，只有当获取高防IP资源列表是可以选填，取值为[1（BGP线路），2（南京电信），3（南京联通），99（第三方合作线路）]，当获取其他产品时请填空数组；
	Line []*uint64 `json:"Line,omitempty" name:"Line"`

	// 资源ID搜索，可选，当不为空数组时表示获取指定资源的资源列表；
	IdList []*string `json:"IdList,omitempty" name:"IdList"`

	// 资源名称搜索，可选，当不为空字符串时表示按名称搜索资源；
	Name *string `json:"Name,omitempty" name:"Name"`

	// IP搜索列表，可选，当不为空时表示按照IP搜索资源；
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// 资源状态搜索列表，可选，取值为[0（运行中）, 1（清洗中）, 2（封堵中）]，当填空数组时不进行状态搜索；
	Status []*uint64 `json:"Status,omitempty" name:"Status"`

	// 即将到期搜索；可选，取值为[0（不搜索），1（搜索即将到期的资源）]
	Expire *uint64 `json:"Expire,omitempty" name:"Expire"`

	// 排序字段，可选
	OderBy []*OrderBy `json:"OderBy,omitempty" name:"OderBy"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 高防IP专业版资源的CNAME，可选，只对高防IP专业版资源列表有效；
	CName *string `json:"CName,omitempty" name:"CName"`

	// 高防IP专业版资源的域名，可选，只对高防IP专业版资源列表有效；
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type DescribeResourceListRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgp表示独享包；bgp-multip表示共享包；bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 地域码搜索，可选，当不指定地域时空数组，当指定地域时，填地域码。例如：["gz", "sh"]
	RegionList []*string `json:"RegionList,omitempty" name:"RegionList"`

	// 线路搜索，可选，只有当获取高防IP资源列表是可以选填，取值为[1（BGP线路），2（南京电信），3（南京联通），99（第三方合作线路）]，当获取其他产品时请填空数组；
	Line []*uint64 `json:"Line,omitempty" name:"Line"`

	// 资源ID搜索，可选，当不为空数组时表示获取指定资源的资源列表；
	IdList []*string `json:"IdList,omitempty" name:"IdList"`

	// 资源名称搜索，可选，当不为空字符串时表示按名称搜索资源；
	Name *string `json:"Name,omitempty" name:"Name"`

	// IP搜索列表，可选，当不为空时表示按照IP搜索资源；
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// 资源状态搜索列表，可选，取值为[0（运行中）, 1（清洗中）, 2（封堵中）]，当填空数组时不进行状态搜索；
	Status []*uint64 `json:"Status,omitempty" name:"Status"`

	// 即将到期搜索；可选，取值为[0（不搜索），1（搜索即将到期的资源）]
	Expire *uint64 `json:"Expire,omitempty" name:"Expire"`

	// 排序字段，可选
	OderBy []*OrderBy `json:"OderBy,omitempty" name:"OderBy"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 高防IP专业版资源的CNAME，可选，只对高防IP专业版资源列表有效；
	CName *string `json:"CName,omitempty" name:"CName"`

	// 高防IP专业版资源的域名，可选，只对高防IP专业版资源列表有效；
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeResourceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "RegionList")
	delete(f, "Line")
	delete(f, "IdList")
	delete(f, "Name")
	delete(f, "IpList")
	delete(f, "Status")
	delete(f, "Expire")
	delete(f, "OderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "CName")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceListResponseParams struct {
	// 总记录数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 资源记录列表，返回Key值说明：
	// "Key": "CreateTime" 表示资源实例购买时间
	// "Key": "Region" 表示资源实例的地域
	// "Key": "BoundIP" 表示独享包实例绑定的IP
	// "Key": "Id" 表示资源实例的ID
	// "Key": "CCEnabled" 表示资源实例的CC防护开关状态
	// "Key": "DDoSThreshold" 表示资源实例的DDoS的清洗阈值	
	// "Key": "BoundStatus" 表示独享包或共享包实例的绑定IP操作状态(绑定中或绑定完成)
	// "Key": "Type" 此字段弃用
	// "Key": "ElasticLimit" 表示资源实例的弹性防护值
	// "Key": "DDoSAI" 表示资源实例的DDoS AI防护开关
	// "Key": "OverloadCount" 表示资源实例受到超过弹性防护值的次数
	// "Key": "Status" 表示资源实例的状态(idle:运行中, attacking:攻击中, blocking:封堵中, isolate:隔离中)
	// "Key": "Lbid" 此字段弃用
	// "Key": "ShowFlag" 此字段弃用
	// "Key": "Expire" 表示资源实例的过期时间
	// "Key": "CCThreshold" 表示资源实例的CC防护触发阈值
	// "Key": "AutoRenewFlag" 表示资源实例的自动续费是否开启
	// "Key": "IspCode" 表示独享包或共享包的线路(0-电信, 1-联通, 2-移动, 5-BGP)
	// "Key": "PackType" 表示套餐包类型
	// "Key": "PackId" 表示套餐包ID
	// "Key": "Name" 表示资源实例的名称
	// "Key": "Locked" 此字段弃用
	// "Key": "IpDDoSLevel" 表示资源实例的防护等级(low-宽松, middle-正常, high-严格)
	// "Key": "DefendStatus" 表示资源实例的DDoS防护状态(防护开启或临时关闭)
	// "Key": "UndefendExpire" 表示资源实例的DDoS防护临时关闭结束时间
	// "Key": "Tgw" 表示资源实例是否是新资源
	// "Key": "Bandwidth" 表示资源实例的保底防护值，只针对高防包和高防IP
	// "Key": "DdosMax" 表示资源实例的保底防护值，只针对高防IP专业版
	// "Key": "GFBandwidth" 表示资源实例的保底业务带宽，只针对高防IP
	// "Key": "ServiceBandwidth" 表示资源实例的保底业务带宽，只针对高防IP专业版
	ServicePacks []*KeyValueRecord `json:"ServicePacks,omitempty" name:"ServicePacks"`

	// 大禹子产品代号（bgp表示独享包；bgp-multip表示共享包；bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceListResponseParams `json:"Response"`
}

func (r *DescribeResourceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleSetsRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID列表
	IdList []*string `json:"IdList,omitempty" name:"IdList"`
}

type DescribeRuleSetsRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID列表
	IdList []*string `json:"IdList,omitempty" name:"IdList"`
}

func (r *DescribeRuleSetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleSetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "IdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleSetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleSetsResponseParams struct {
	// 规则记录数数组，取值说明:
	// Key值为"Id"时，Value表示资源ID
	// Key值为"RuleIdList"时，Value值表示资源的规则ID，多个规则ID用","分割
	// Key值为"RuleNameList"时，Value值表示资源的规则名，多个规则名用","分割
	// Key值为"RuleNum"时，Value值表示资源的规则数
	L4RuleSets []*KeyValueRecord `json:"L4RuleSets,omitempty" name:"L4RuleSets"`

	// 规则记录数数组，取值说明:
	// Key值为"Id"时，Value表示资源ID
	// Key值为"RuleIdList"时，Value值表示资源的规则ID，多个规则ID用","分割
	// Key值为"RuleNameList"时，Value值表示资源的规则名，多个规则名用","分割
	// Key值为"RuleNum"时，Value值表示资源的规则数
	L7RuleSets []*KeyValueRecord `json:"L7RuleSets,omitempty" name:"L7RuleSets"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRuleSetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleSetsResponseParams `json:"Response"`
}

func (r *DescribeRuleSetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSchedulingDomainListRequestParams struct {
	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 可选，筛选特定的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type DescribeSchedulingDomainListRequest struct {
	*tchttp.BaseRequest
	
	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 可选，筛选特定的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeSchedulingDomainListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSchedulingDomainListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSchedulingDomainListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSchedulingDomainListResponseParams struct {
	// 调度域名总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 调度域名列表信息
	DomainList []*SchedulingDomain `json:"DomainList,omitempty" name:"DomainList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSchedulingDomainListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSchedulingDomainListResponseParams `json:"Response"`
}

func (r *DescribeSchedulingDomainListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSchedulingDomainListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecIndexRequestParams struct {

}

type DescribeSecIndexRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSecIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecIndexResponseParams struct {
	// 字段值，如下：
	// AttackIpCount：受攻击的IP数
	// AttackCount：攻击次数
	// BlockCount：封堵次数
	// MaxMbps：攻击峰值Mbps
	// IpNum：统计的IP数据
	Data []*KeyValue `json:"Data,omitempty" name:"Data"`

	// 本月开始时间
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 本月结束时间
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecIndexResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecIndexResponseParams `json:"Response"`
}

func (r *DescribeSecIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSourceIpSegmentRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DescribeSourceIpSegmentRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeSourceIpSegmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceIpSegmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSourceIpSegmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSourceIpSegmentResponseParams struct {
	// 回源IP段，多个用"；"分隔
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSourceIpSegmentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSourceIpSegmentResponseParams `json:"Response"`
}

func (r *DescribeSourceIpSegmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceIpSegmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTransmitStatisRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版；bgp表示独享包；bgp-multip表示共享包）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 指标名，取值：
	// traffic表示流量带宽；
	// pkg表示包速率；
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 统计时间粒度（300表示5分钟；3600表示小时；86400表示天）
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间，秒部分保持为0，分钟部分为5的倍数
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间，秒部分保持为0，分钟部分为5的倍数
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 资源的IP（当Business为bgp-multip时必填，且仅支持一个IP）；当不填写时，默认统计资源实例的所有IP；资源实例有多个IP（比如高防IP专业版）时，统计方式是求和；
	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

type DescribeTransmitStatisRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版；bgp表示独享包；bgp-multip表示共享包）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 指标名，取值：
	// traffic表示流量带宽；
	// pkg表示包速率；
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 统计时间粒度（300表示5分钟；3600表示小时；86400表示天）
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间，秒部分保持为0，分钟部分为5的倍数
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间，秒部分保持为0，分钟部分为5的倍数
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 资源的IP（当Business为bgp-multip时必填，且仅支持一个IP）；当不填写时，默认统计资源实例的所有IP；资源实例有多个IP（比如高防IP专业版）时，统计方式是求和；
	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

func (r *DescribeTransmitStatisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTransmitStatisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "MetricName")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "IpList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTransmitStatisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTransmitStatisResponseParams struct {
	// 当MetricName=traffic时，表示入流量带宽，单位bps；
	// 当MetricName=pkg时，表示入包速率，单位pps；
	InDataList []*float64 `json:"InDataList,omitempty" name:"InDataList"`

	// 当MetricName=traffic时，表示出流量带宽，单位bps；
	// 当MetricName=pkg时，表示出包速率，单位pps；
	OutDataList []*float64 `json:"OutDataList,omitempty" name:"OutDataList"`

	// 指标名：
	// traffic表示流量带宽；
	// pkg表示包速率；
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTransmitStatisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTransmitStatisResponseParams `json:"Response"`
}

func (r *DescribeTransmitStatisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTransmitStatisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnBlockStatisRequestParams struct {

}

type DescribeUnBlockStatisRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUnBlockStatisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnBlockStatisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnBlockStatisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnBlockStatisResponseParams struct {
	// 解封总配额数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 已使用次数
	Used *uint64 `json:"Used,omitempty" name:"Used"`

	// 统计起始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUnBlockStatisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUnBlockStatisResponseParams `json:"Response"`
}

func (r *DescribeUnBlockStatisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnBlockStatisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribleL4RulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID，可选参数，填写后获取指定的规则
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribleL4RulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID，可选参数，填写后获取指定的规则
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribleL4RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribleL4RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleIdList")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribleL4RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribleL4RulesResponseParams struct {
	// 转发规则列表
	Rules []*L4RuleEntry `json:"Rules,omitempty" name:"Rules"`

	// 总规则数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 健康检查配置列表
	Healths []*L4RuleHealth `json:"Healths,omitempty" name:"Healths"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribleL4RulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribleL4RulesResponseParams `json:"Response"`
}

func (r *DescribleL4RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribleL4RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribleL7RulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID，可选参数，填写后获取指定的规则
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 域名搜索，选填，当需要搜索域名请填写
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 转发协议搜索，选填，取值[http, https, http/https]
	ProtocolList []*string `json:"ProtocolList,omitempty" name:"ProtocolList"`

	// 状态搜索，选填，取值[0(规则配置成功)，1(规则配置生效中)，2(规则配置失败)，3(规则删除生效中)，5(规则删除失败)，6(规则等待配置)，7(规则等待删除)，8(规则待配置证书)]
	StatusList []*uint64 `json:"StatusList,omitempty" name:"StatusList"`
}

type DescribleL7RulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID，可选参数，填写后获取指定的规则
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 域名搜索，选填，当需要搜索域名请填写
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 转发协议搜索，选填，取值[http, https, http/https]
	ProtocolList []*string `json:"ProtocolList,omitempty" name:"ProtocolList"`

	// 状态搜索，选填，取值[0(规则配置成功)，1(规则配置生效中)，2(规则配置失败)，3(规则删除生效中)，5(规则删除失败)，6(规则等待配置)，7(规则等待删除)，8(规则待配置证书)]
	StatusList []*uint64 `json:"StatusList,omitempty" name:"StatusList"`
}

func (r *DescribleL7RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribleL7RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleIdList")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Domain")
	delete(f, "ProtocolList")
	delete(f, "StatusList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribleL7RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribleL7RulesResponseParams struct {
	// 转发规则列表
	Rules []*L7RuleEntry `json:"Rules,omitempty" name:"Rules"`

	// 总规则数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 健康检查配置列表
	Healths []*L7RuleHealth `json:"Healths,omitempty" name:"Healths"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribleL7RulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribleL7RulesResponseParams `json:"Response"`
}

func (r *DescribleL7RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribleL7RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribleNewL7RulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 域名搜索，选填，当需要搜索域名请填写
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 转发协议搜索，选填，取值[http, https, http/https]
	ProtocolList []*string `json:"ProtocolList,omitempty" name:"ProtocolList"`

	// 状态搜索，选填，取值[0(规则配置成功)，1(规则配置生效中)，2(规则配置失败)，3(规则删除生效中)，5(规则删除失败)，6(规则等待配置)，7(规则等待删除)，8(规则待配置证书)]
	StatusList []*uint64 `json:"StatusList,omitempty" name:"StatusList"`

	// IP搜索，选填，当需要搜索IP请填写
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type DescribleNewL7RulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 域名搜索，选填，当需要搜索域名请填写
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 转发协议搜索，选填，取值[http, https, http/https]
	ProtocolList []*string `json:"ProtocolList,omitempty" name:"ProtocolList"`

	// 状态搜索，选填，取值[0(规则配置成功)，1(规则配置生效中)，2(规则配置失败)，3(规则删除生效中)，5(规则删除失败)，6(规则等待配置)，7(规则等待删除)，8(规则待配置证书)]
	StatusList []*uint64 `json:"StatusList,omitempty" name:"StatusList"`

	// IP搜索，选填，当需要搜索IP请填写
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

func (r *DescribleNewL7RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribleNewL7RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Domain")
	delete(f, "ProtocolList")
	delete(f, "StatusList")
	delete(f, "Ip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribleNewL7RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribleNewL7RulesResponseParams struct {
	// 转发规则列表
	Rules []*NewL7RuleEntry `json:"Rules,omitempty" name:"Rules"`

	// 总规则数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 健康检查配置列表
	Healths []*L7RuleHealth `json:"Healths,omitempty" name:"Healths"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribleNewL7RulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribleNewL7RulesResponseParams `json:"Response"`
}

func (r *DescribleNewL7RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribleNewL7RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribleRegionCountRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 根据线路统计，取值为[1（BGP线路），2（南京电信），3（南京联通），99（第三方合作线路）]；只对高防IP产品有效，其他产品此字段忽略
	LineList []*uint64 `json:"LineList,omitempty" name:"LineList"`
}

type DescribleRegionCountRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 根据线路统计，取值为[1（BGP线路），2（南京电信），3（南京联通），99（第三方合作线路）]；只对高防IP产品有效，其他产品此字段忽略
	LineList []*uint64 `json:"LineList,omitempty" name:"LineList"`
}

func (r *DescribleRegionCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribleRegionCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "LineList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribleRegionCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribleRegionCountResponseParams struct {
	// 地域资源实例数
	RegionList []*RegionInstanceCount `json:"RegionList,omitempty" name:"RegionList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribleRegionCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribleRegionCountResponseParams `json:"Response"`
}

func (r *DescribleRegionCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribleRegionCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HttpStatusMap struct {
	// http2xx状态码
	Http2xx []*float64 `json:"Http2xx,omitempty" name:"Http2xx"`

	// http3xx状态码
	Http3xx []*float64 `json:"Http3xx,omitempty" name:"Http3xx"`

	// http404状态码
	Http404 []*float64 `json:"Http404,omitempty" name:"Http404"`

	// http4xx状态码
	Http4xx []*float64 `json:"Http4xx,omitempty" name:"Http4xx"`

	// http5xx状态码
	Http5xx []*float64 `json:"Http5xx,omitempty" name:"Http5xx"`

	// http2xx回源状态码
	SourceHttp2xx []*float64 `json:"SourceHttp2xx,omitempty" name:"SourceHttp2xx"`

	// http3xx回源状态码
	SourceHttp3xx []*float64 `json:"SourceHttp3xx,omitempty" name:"SourceHttp3xx"`

	// http404回源状态码
	SourceHttp404 []*float64 `json:"SourceHttp404,omitempty" name:"SourceHttp404"`

	// http4xx回源状态码
	SourceHttp4xx []*float64 `json:"SourceHttp4xx,omitempty" name:"SourceHttp4xx"`

	// http5xx回源状态码
	SourceHttp5xx []*float64 `json:"SourceHttp5xx,omitempty" name:"SourceHttp5xx"`
}

type IpBlackWhite struct {
	// IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 黑白类型，取值范围[black，white]
	Type *string `json:"Type,omitempty" name:"Type"`
}

type IpBlockData struct {
	// IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 状态（Blocked：被封堵；UnBlocking：解封中；UnBlockFailed：解封失败）
	Status *string `json:"Status,omitempty" name:"Status"`

	// 封堵时间
	BlockTime *string `json:"BlockTime,omitempty" name:"BlockTime"`

	// 解封时间（预计解封时间）
	UnBlockTime *string `json:"UnBlockTime,omitempty" name:"UnBlockTime"`

	// 解封类型（user：自助解封；auto：自动解封； update：升级解封；bind：绑定高防包解封）
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 高防标记，0：非高防，1：高防
	ProtectFlag *uint64 `json:"ProtectFlag,omitempty" name:"ProtectFlag"`
}

type IpUnBlockData struct {
	// IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 封堵时间
	BlockTime *string `json:"BlockTime,omitempty" name:"BlockTime"`

	// 解封时间（实际解封时间）
	UnBlockTime *string `json:"UnBlockTime,omitempty" name:"UnBlockTime"`

	// 解封类型（user：自助解封；auto：自动解封； update：升级解封；bind：绑定高防包解封）
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

type KeyValue struct {
	// 字段名称
	Key *string `json:"Key,omitempty" name:"Key"`

	// 字段取值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type KeyValueRecord struct {
	// 一条记录的Key-Value数组
	Record []*KeyValue `json:"Record,omitempty" name:"Record"`
}

type L4DelRule struct {
	// 资源Id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 规则Id
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

type L4HealthConfig struct {
	// 转发协议，取值[TCP, UDP]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 转发端口
	VirtualPort *uint64 `json:"VirtualPort,omitempty" name:"VirtualPort"`

	// =1表示开启；=0表示关闭
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

	// 响应超时时间，单位秒
	TimeOut *uint64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// 检测间隔时间，单位秒
	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`

	// 不健康阈值，单位次
	KickNum *uint64 `json:"KickNum,omitempty" name:"KickNum"`

	// 健康阈值，单位次
	AliveNum *uint64 `json:"AliveNum,omitempty" name:"AliveNum"`

	// 会话保持时间，单位秒
	KeepTime *uint64 `json:"KeepTime,omitempty" name:"KeepTime"`
}

type L4RuleEntry struct {
	// 转发协议，取值[TCP, UDP]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 转发端口
	VirtualPort *uint64 `json:"VirtualPort,omitempty" name:"VirtualPort"`

	// 源站端口
	SourcePort *uint64 `json:"SourcePort,omitempty" name:"SourcePort"`

	// 回源方式，取值[1(域名回源)，2(IP回源)]
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 会话保持时间，单位秒
	KeepTime *uint64 `json:"KeepTime,omitempty" name:"KeepTime"`

	// 回源列表
	SourceList []*L4RuleSource `json:"SourceList,omitempty" name:"SourceList"`

	// 负载均衡方式，取值[1(加权轮询)，2(源IP hash)]
	LbType *uint64 `json:"LbType,omitempty" name:"LbType"`

	// 会话保持开关，取值[0(会话保持关闭)，1(会话保持开启)]；
	KeepEnable *uint64 `json:"KeepEnable,omitempty" name:"KeepEnable"`

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 规则描述
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 移除水印状态，取值[0(关闭)，1(开启)]
	RemoveSwitch *uint64 `json:"RemoveSwitch,omitempty" name:"RemoveSwitch"`
}

type L4RuleHealth struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// =1表示开启；=0表示关闭
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

	// 响应超时时间，单位秒
	TimeOut *uint64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// 检测间隔时间，单位秒，必须要大于响应超时时间
	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`

	// 不健康阈值，单位次
	KickNum *uint64 `json:"KickNum,omitempty" name:"KickNum"`

	// 健康阈值，单位次
	AliveNum *uint64 `json:"AliveNum,omitempty" name:"AliveNum"`
}

type L4RuleSource struct {
	// 回源IP或域名
	Source *string `json:"Source,omitempty" name:"Source"`

	// 权重值，取值[0,100]
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
}

type L7HealthConfig struct {
	// 转发协议，取值[http, https, http/https]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 转发域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// =1表示开启；=0表示关闭
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

	// 检测间隔时间，单位秒
	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`

	// 异常判定次数，单位次
	KickNum *uint64 `json:"KickNum,omitempty" name:"KickNum"`

	// 健康判定次数，单位次
	AliveNum *uint64 `json:"AliveNum,omitempty" name:"AliveNum"`

	// 健康检查探测方法，可选HEAD或GET，默认为HEAD
	Method *string `json:"Method,omitempty" name:"Method"`

	// 健康检查判定正常状态码，1xx =1, 2xx=2, 3xx=4, 4xx=8,5xx=16，多个状态码值加和
	StatusCode *uint64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 检查目录的URL，默认为/
	Url *string `json:"Url,omitempty" name:"Url"`
}

type L7RuleEntry struct {
	// 转发协议，取值[http, https]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 转发域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 回源方式，取值[1(域名回源)，2(IP回源)]
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 会话保持时间，单位秒
	KeepTime *uint64 `json:"KeepTime,omitempty" name:"KeepTime"`

	// 回源列表
	SourceList []*L4RuleSource `json:"SourceList,omitempty" name:"SourceList"`

	// 负载均衡方式，取值[1(加权轮询)]
	LbType *uint64 `json:"LbType,omitempty" name:"LbType"`

	// 会话保持开关，取值[0(会话保持关闭)，1(会话保持开启)]
	KeepEnable *uint64 `json:"KeepEnable,omitempty" name:"KeepEnable"`

	// 规则ID，当添加新规则时可以不用填写此字段；当修改或者删除规则时需要填写此字段；
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 证书来源，当转发协议为https时必须填，取值[2(腾讯云托管证书)]，当转发协议为http时也可以填0
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// 当证书来源为腾讯云托管证书时，此字段必须填写托管证书ID
	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`

	// 当证书来源为自有证书时，此字段必须填写证书内容；(因已不再支持自有证书，此字段已弃用，请不用填写此字段)
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// 当证书来源为自有证书时，此字段必须填写证书密钥；(因已不再支持自有证书，此字段已弃用，请不用填写此字段)
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// 规则描述
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则状态，取值[0(规则配置成功)，1(规则配置生效中)，2(规则配置失败)，3(规则删除生效中)，5(规则删除失败)，6(规则等待配置)，7(规则等待删除)，8(规则待配置证书)]
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// cc防护状态，取值[0(关闭), 1(开启)]
	CCStatus *uint64 `json:"CCStatus,omitempty" name:"CCStatus"`

	// HTTPS协议的CC防护状态，取值[0(关闭), 1(开启)]
	CCEnable *uint64 `json:"CCEnable,omitempty" name:"CCEnable"`

	// HTTPS协议的CC防护阈值
	CCThreshold *uint64 `json:"CCThreshold,omitempty" name:"CCThreshold"`

	// HTTPS协议的CC防护等级
	CCLevel *string `json:"CCLevel,omitempty" name:"CCLevel"`

	// 是否开启Https协议使用Http回源，取值[0(关闭), 1(开启)]，不填写默认是关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpsToHttpEnable *uint64 `json:"HttpsToHttpEnable,omitempty" name:"HttpsToHttpEnable"`

	// 接入端口值
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualPort *uint64 `json:"VirtualPort,omitempty" name:"VirtualPort"`
}

type L7RuleHealth struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// =1表示开启；=0表示关闭
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

	// 检测间隔时间，单位秒
	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`

	// 不健康阈值，单位次
	KickNum *uint64 `json:"KickNum,omitempty" name:"KickNum"`

	// 健康阈值，单位次
	AliveNum *uint64 `json:"AliveNum,omitempty" name:"AliveNum"`

	// HTTP请求方式，取值[HEAD,GET]
	Method *string `json:"Method,omitempty" name:"Method"`

	// 健康检查判定正常状态码，1xx =1, 2xx=2, 3xx=4, 4xx=8,5xx=16，多个状态码值加和
	StatusCode *uint64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 检查目录的URL，默认为/
	Url *string `json:"Url,omitempty" name:"Url"`

	// 配置状态，0： 正常，1：配置中，2：配置失败
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type ModifyCCAlarmThresholdRequestParams struct {
	// 大禹子产品代号（shield表示棋牌；bgpip表示高防IP；bgp表示高防包；bgp-multip表示多ip高防包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID,字符串类型
	RsId *string `json:"RsId,omitempty" name:"RsId"`

	// 告警阈值，大于0（目前排定的值），后台设置默认值为1000
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`

	// 资源关联的IP列表，高防包未绑定时，传空数组，高防IP专业版传多个IP的数据
	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

type ModifyCCAlarmThresholdRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（shield表示棋牌；bgpip表示高防IP；bgp表示高防包；bgp-multip表示多ip高防包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID,字符串类型
	RsId *string `json:"RsId,omitempty" name:"RsId"`

	// 告警阈值，大于0（目前排定的值），后台设置默认值为1000
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`

	// 资源关联的IP列表，高防包未绑定时，传空数组，高防IP专业版传多个IP的数据
	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

func (r *ModifyCCAlarmThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCAlarmThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "RsId")
	delete(f, "AlarmThreshold")
	delete(f, "IpList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCCAlarmThresholdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCAlarmThresholdResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCCAlarmThresholdResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCCAlarmThresholdResponseParams `json:"Response"`
}

func (r *ModifyCCAlarmThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCAlarmThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCFrequencyRulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// CC的访问频率控制规则ID
	CCFrequencyRuleId *string `json:"CCFrequencyRuleId,omitempty" name:"CCFrequencyRuleId"`

	// 匹配规则，取值["include"(前缀匹配)，"equal"(完全匹配)]
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 统计周期，单位秒，取值[10, 30, 60]
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 访问次数，取值[1-10000]
	ReqNumber *uint64 `json:"ReqNumber,omitempty" name:"ReqNumber"`

	// 执行动作，取值["alg"（人机识别）, "drop"（拦截）]
	Act *string `json:"Act,omitempty" name:"Act"`

	// 执行时间，单位秒，取值[1-900]
	ExeDuration *uint64 `json:"ExeDuration,omitempty" name:"ExeDuration"`

	// URI字符串，必须以/开头，例如/abc/a.php，长度不超过31；当URI=/时，匹配模式只能选择前缀匹配；
	Uri *string `json:"Uri,omitempty" name:"Uri"`

	// User-Agent字符串，长度不超过80
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// Cookie字符串，长度不超过40
	Cookie *string `json:"Cookie,omitempty" name:"Cookie"`
}

type ModifyCCFrequencyRulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// CC的访问频率控制规则ID
	CCFrequencyRuleId *string `json:"CCFrequencyRuleId,omitempty" name:"CCFrequencyRuleId"`

	// 匹配规则，取值["include"(前缀匹配)，"equal"(完全匹配)]
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 统计周期，单位秒，取值[10, 30, 60]
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 访问次数，取值[1-10000]
	ReqNumber *uint64 `json:"ReqNumber,omitempty" name:"ReqNumber"`

	// 执行动作，取值["alg"（人机识别）, "drop"（拦截）]
	Act *string `json:"Act,omitempty" name:"Act"`

	// 执行时间，单位秒，取值[1-900]
	ExeDuration *uint64 `json:"ExeDuration,omitempty" name:"ExeDuration"`

	// URI字符串，必须以/开头，例如/abc/a.php，长度不超过31；当URI=/时，匹配模式只能选择前缀匹配；
	Uri *string `json:"Uri,omitempty" name:"Uri"`

	// User-Agent字符串，长度不超过80
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// Cookie字符串，长度不超过40
	Cookie *string `json:"Cookie,omitempty" name:"Cookie"`
}

func (r *ModifyCCFrequencyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCFrequencyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "CCFrequencyRuleId")
	delete(f, "Mode")
	delete(f, "Period")
	delete(f, "ReqNumber")
	delete(f, "Act")
	delete(f, "ExeDuration")
	delete(f, "Uri")
	delete(f, "UserAgent")
	delete(f, "Cookie")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCCFrequencyRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCFrequencyRulesResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCCFrequencyRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCCFrequencyRulesResponseParams `json:"Response"`
}

func (r *ModifyCCFrequencyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCFrequencyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCFrequencyRulesStatusRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 7层转发规则ID（通过获取7层转发规则接口可以获取规则ID）
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 开启或关闭，取值["on"(开启)，"off"(关闭)]
	Method *string `json:"Method,omitempty" name:"Method"`
}

type ModifyCCFrequencyRulesStatusRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 7层转发规则ID（通过获取7层转发规则接口可以获取规则ID）
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 开启或关闭，取值["on"(开启)，"off"(关闭)]
	Method *string `json:"Method,omitempty" name:"Method"`
}

func (r *ModifyCCFrequencyRulesStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCFrequencyRulesStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleId")
	delete(f, "Method")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCCFrequencyRulesStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCFrequencyRulesStatusResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCCFrequencyRulesStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCCFrequencyRulesStatusResponseParams `json:"Response"`
}

func (r *ModifyCCFrequencyRulesStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCFrequencyRulesStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCHostProtectionRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 开启/关闭CC域名防护，取值[open(表示开启)，close(表示关闭)]
	Method *string `json:"Method,omitempty" name:"Method"`
}

type ModifyCCHostProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 开启/关闭CC域名防护，取值[open(表示开启)，close(表示关闭)]
	Method *string `json:"Method,omitempty" name:"Method"`
}

func (r *ModifyCCHostProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCHostProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleId")
	delete(f, "Method")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCCHostProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCHostProtectionResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCCHostProtectionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCCHostProtectionResponseParams `json:"Response"`
}

func (r *ModifyCCHostProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCHostProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCIpAllowDenyRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// add表示添加，delete表示删除
	Method *string `json:"Method,omitempty" name:"Method"`

	// 黑/白名单类型；取值[white(白名单)，black(黑名单)]
	Type *string `json:"Type,omitempty" name:"Type"`

	// 黑/白名单的IP数组
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// 可选字段，代表CC防护类型，取值[http（HTTP协议的CC防护），https（HTTPS协议的CC防护）]；当不填时，默认为HTTP协议的CC防护；当填写https时还需要填写Domain和RuleId字段；
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 可选字段，表示HTTPS协议的7层转发规则域名（通过获取7层转发规则接口可以获取域名），只有当Protocol字段为https时才必须填写此字段；
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 可选字段，表示HTTPS协议的7层转发规则ID（通过获取7层转发规则接口可以获取规则ID），
	// 当Method为delete时，不用填写此字段；
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type ModifyCCIpAllowDenyRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// add表示添加，delete表示删除
	Method *string `json:"Method,omitempty" name:"Method"`

	// 黑/白名单类型；取值[white(白名单)，black(黑名单)]
	Type *string `json:"Type,omitempty" name:"Type"`

	// 黑/白名单的IP数组
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// 可选字段，代表CC防护类型，取值[http（HTTP协议的CC防护），https（HTTPS协议的CC防护）]；当不填时，默认为HTTP协议的CC防护；当填写https时还需要填写Domain和RuleId字段；
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 可选字段，表示HTTPS协议的7层转发规则域名（通过获取7层转发规则接口可以获取域名），只有当Protocol字段为https时才必须填写此字段；
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 可选字段，表示HTTPS协议的7层转发规则ID（通过获取7层转发规则接口可以获取规则ID），
	// 当Method为delete时，不用填写此字段；
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *ModifyCCIpAllowDenyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCIpAllowDenyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Method")
	delete(f, "Type")
	delete(f, "IpList")
	delete(f, "Protocol")
	delete(f, "Domain")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCCIpAllowDenyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCIpAllowDenyResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCCIpAllowDenyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCCIpAllowDenyResponseParams `json:"Response"`
}

func (r *ModifyCCIpAllowDenyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCIpAllowDenyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCLevelRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// CC防护等级，取值[default(正常), loose(宽松), strict(严格)];
	Level *string `json:"Level,omitempty" name:"Level"`

	// 可选字段，代表CC防护类型，取值[http（HTTP协议的CC防护），https（HTTPS协议的CC防护）]；当不填时，默认为HTTP协议的CC防护；当填写https时还需要填写RuleId字段；
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 表示7层转发规则ID（通过获取7层转发规则接口可以获取规则ID）；
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type ModifyCCLevelRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// CC防护等级，取值[default(正常), loose(宽松), strict(严格)];
	Level *string `json:"Level,omitempty" name:"Level"`

	// 可选字段，代表CC防护类型，取值[http（HTTP协议的CC防护），https（HTTPS协议的CC防护）]；当不填时，默认为HTTP协议的CC防护；当填写https时还需要填写RuleId字段；
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 表示7层转发规则ID（通过获取7层转发规则接口可以获取规则ID）；
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *ModifyCCLevelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCLevelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Level")
	delete(f, "Protocol")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCCLevelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCLevelResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCCLevelResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCCLevelResponseParams `json:"Response"`
}

func (r *ModifyCCLevelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCPolicySwitchRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 策略ID
	SetId *string `json:"SetId,omitempty" name:"SetId"`

	// 开关状态
	Switch *uint64 `json:"Switch,omitempty" name:"Switch"`
}

type ModifyCCPolicySwitchRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 策略ID
	SetId *string `json:"SetId,omitempty" name:"SetId"`

	// 开关状态
	Switch *uint64 `json:"Switch,omitempty" name:"Switch"`
}

func (r *ModifyCCPolicySwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCPolicySwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "SetId")
	delete(f, "Switch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCCPolicySwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCPolicySwitchResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCCPolicySwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCCPolicySwitchResponseParams `json:"Response"`
}

func (r *ModifyCCPolicySwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCPolicySwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCSelfDefinePolicyRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 策略ID
	SetId *string `json:"SetId,omitempty" name:"SetId"`

	// CC策略描述
	Policy *CCPolicy `json:"Policy,omitempty" name:"Policy"`
}

type ModifyCCSelfDefinePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 策略ID
	SetId *string `json:"SetId,omitempty" name:"SetId"`

	// CC策略描述
	Policy *CCPolicy `json:"Policy,omitempty" name:"Policy"`
}

func (r *ModifyCCSelfDefinePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCSelfDefinePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "SetId")
	delete(f, "Policy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCCSelfDefinePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCSelfDefinePolicyResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCCSelfDefinePolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCCSelfDefinePolicyResponseParams `json:"Response"`
}

func (r *ModifyCCSelfDefinePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCSelfDefinePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCThresholdRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// CC防护阈值，取值(0 100 150 240 350 480 550 700 850 1000 1500 2000 3000 5000 10000 20000);
	// 当Business为高防IP、高防IP专业版时，其CC防护最大阈值跟资源的保底防护带宽有关，对应关系如下：
	//   保底带宽: 最大C防护阈值
	//   10:  20000,
	//   20:  40000,
	//   30:  70000,
	//   40:  100000,
	//   50:  150000,
	//   60:  200000,
	//   80:  250000,
	//   100: 300000,
	Threshold *uint64 `json:"Threshold,omitempty" name:"Threshold"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 可选字段，代表CC防护类型，取值[http（HTTP协议的CC防护），https（HTTPS协议的CC防护）]；当不填时，默认为HTTP协议的CC防护；当填写https时还需要填写RuleId字段；
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 可选字段，表示HTTPS协议的7层转发规则ID（通过获取7层转发规则接口可以获取规则ID）；
	// 当Protocol=https时必须填写；
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 查询的IP地址（仅基础防护提供），取值如：1.1.1.1
	BasicIp *string `json:"BasicIp,omitempty" name:"BasicIp"`

	// 查询IP所属地域（仅基础防护提供），取值如：gz、bj、sh、hk等地域缩写
	BasicRegion *string `json:"BasicRegion,omitempty" name:"BasicRegion"`

	// 专区类型（仅基础防护提供），取值如：公有云专区：public，黑石专区：bm, NAT服务器专区：nat，互联网通道：channel。
	BasicBizType *string `json:"BasicBizType,omitempty" name:"BasicBizType"`

	// 设备类型（仅基础防护提供），取值如：服务器：cvm，公有云负载均衡：clb，黑石负载均衡：lb，NAT服务器：nat，互联网通道：channel.
	BasicDeviceType *string `json:"BasicDeviceType,omitempty" name:"BasicDeviceType"`

	// 仅基础防护提供。可选，IPInstance Nat 网关（如果查询的设备类型是NAT服务器，需要传此参数，通过nat资源查询接口获取）
	BasicIpInstance *string `json:"BasicIpInstance,omitempty" name:"BasicIpInstance"`

	// 仅基础防护提供。可选，运营商线路（如果查询的设备类型是NAT服务器，需要传此参数为5）
	BasicIspCode *uint64 `json:"BasicIspCode,omitempty" name:"BasicIspCode"`

	// 可选字段，当协议取值HTTPS时，必填
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type ModifyCCThresholdRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// CC防护阈值，取值(0 100 150 240 350 480 550 700 850 1000 1500 2000 3000 5000 10000 20000);
	// 当Business为高防IP、高防IP专业版时，其CC防护最大阈值跟资源的保底防护带宽有关，对应关系如下：
	//   保底带宽: 最大C防护阈值
	//   10:  20000,
	//   20:  40000,
	//   30:  70000,
	//   40:  100000,
	//   50:  150000,
	//   60:  200000,
	//   80:  250000,
	//   100: 300000,
	Threshold *uint64 `json:"Threshold,omitempty" name:"Threshold"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 可选字段，代表CC防护类型，取值[http（HTTP协议的CC防护），https（HTTPS协议的CC防护）]；当不填时，默认为HTTP协议的CC防护；当填写https时还需要填写RuleId字段；
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 可选字段，表示HTTPS协议的7层转发规则ID（通过获取7层转发规则接口可以获取规则ID）；
	// 当Protocol=https时必须填写；
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 查询的IP地址（仅基础防护提供），取值如：1.1.1.1
	BasicIp *string `json:"BasicIp,omitempty" name:"BasicIp"`

	// 查询IP所属地域（仅基础防护提供），取值如：gz、bj、sh、hk等地域缩写
	BasicRegion *string `json:"BasicRegion,omitempty" name:"BasicRegion"`

	// 专区类型（仅基础防护提供），取值如：公有云专区：public，黑石专区：bm, NAT服务器专区：nat，互联网通道：channel。
	BasicBizType *string `json:"BasicBizType,omitempty" name:"BasicBizType"`

	// 设备类型（仅基础防护提供），取值如：服务器：cvm，公有云负载均衡：clb，黑石负载均衡：lb，NAT服务器：nat，互联网通道：channel.
	BasicDeviceType *string `json:"BasicDeviceType,omitempty" name:"BasicDeviceType"`

	// 仅基础防护提供。可选，IPInstance Nat 网关（如果查询的设备类型是NAT服务器，需要传此参数，通过nat资源查询接口获取）
	BasicIpInstance *string `json:"BasicIpInstance,omitempty" name:"BasicIpInstance"`

	// 仅基础防护提供。可选，运营商线路（如果查询的设备类型是NAT服务器，需要传此参数为5）
	BasicIspCode *uint64 `json:"BasicIspCode,omitempty" name:"BasicIspCode"`

	// 可选字段，当协议取值HTTPS时，必填
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *ModifyCCThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Threshold")
	delete(f, "Id")
	delete(f, "Protocol")
	delete(f, "RuleId")
	delete(f, "BasicIp")
	delete(f, "BasicRegion")
	delete(f, "BasicBizType")
	delete(f, "BasicDeviceType")
	delete(f, "BasicIpInstance")
	delete(f, "BasicIspCode")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCCThresholdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCThresholdResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCCThresholdResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCCThresholdResponseParams `json:"Response"`
}

func (r *ModifyCCThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCUrlAllowRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// =add表示添加，=delete表示删除
	Method *string `json:"Method,omitempty" name:"Method"`

	// 黑/白名单类型；取值[white(白名单)]
	Type *string `json:"Type,omitempty" name:"Type"`

	// URL数组，URL格式如下：
	// http://域名/cgi
	// https://域名/cgi
	UrlList []*string `json:"UrlList,omitempty" name:"UrlList"`

	// 可选字段，代表CC防护类型，取值[http（HTTP协议的CC防护），https（HTTPS协议的CC防护）]；当不填时，默认为HTTP协议的CC防护；当填写https时还需要填写Domain和RuleId字段；
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 可选字段，表示HTTPS协议的7层转发规则域名（通过获取7层转发规则接口可以获取域名），只有当Protocol字段为https时才必须填写此字段；
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 可选字段，表示HTTPS协议的7层转发规则ID（通过获取7层转发规则接口可以获取规则ID），当添加并且Protocol=https时必须填写；
	// 当Method为delete时，可以不用填写此字段；
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type ModifyCCUrlAllowRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// =add表示添加，=delete表示删除
	Method *string `json:"Method,omitempty" name:"Method"`

	// 黑/白名单类型；取值[white(白名单)]
	Type *string `json:"Type,omitempty" name:"Type"`

	// URL数组，URL格式如下：
	// http://域名/cgi
	// https://域名/cgi
	UrlList []*string `json:"UrlList,omitempty" name:"UrlList"`

	// 可选字段，代表CC防护类型，取值[http（HTTP协议的CC防护），https（HTTPS协议的CC防护）]；当不填时，默认为HTTP协议的CC防护；当填写https时还需要填写Domain和RuleId字段；
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 可选字段，表示HTTPS协议的7层转发规则域名（通过获取7层转发规则接口可以获取域名），只有当Protocol字段为https时才必须填写此字段；
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 可选字段，表示HTTPS协议的7层转发规则ID（通过获取7层转发规则接口可以获取规则ID），当添加并且Protocol=https时必须填写；
	// 当Method为delete时，可以不用填写此字段；
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *ModifyCCUrlAllowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCUrlAllowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Method")
	delete(f, "Type")
	delete(f, "UrlList")
	delete(f, "Protocol")
	delete(f, "Domain")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCCUrlAllowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCUrlAllowResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCCUrlAllowResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCCUrlAllowResponseParams `json:"Response"`
}

func (r *ModifyCCUrlAllowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCUrlAllowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSAIStatusRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// =get表示读取AI防护状态；=set表示修改AI防护状态；
	Method *string `json:"Method,omitempty" name:"Method"`

	// AI防护状态，取值[on，off]；当Method=set时必填；
	DDoSAI *string `json:"DDoSAI,omitempty" name:"DDoSAI"`
}

type ModifyDDoSAIStatusRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// =get表示读取AI防护状态；=set表示修改AI防护状态；
	Method *string `json:"Method,omitempty" name:"Method"`

	// AI防护状态，取值[on，off]；当Method=set时必填；
	DDoSAI *string `json:"DDoSAI,omitempty" name:"DDoSAI"`
}

func (r *ModifyDDoSAIStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSAIStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Method")
	delete(f, "DDoSAI")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSAIStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSAIStatusResponseParams struct {
	// AI防护状态，取值[on，off]
	DDoSAI *string `json:"DDoSAI,omitempty" name:"DDoSAI"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDDoSAIStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSAIStatusResponseParams `json:"Response"`
}

func (r *ModifyDDoSAIStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSAIStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSAlarmThresholdRequestParams struct {
	// 大禹子产品代号（shield表示棋牌；bgpip表示高防IP；bgp表示高防包；bgp-multip表示多ip高防包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID,字符串类型
	RsId *string `json:"RsId,omitempty" name:"RsId"`

	// 告警阈值类型，0-未设置，1-入流量，2-清洗流量
	AlarmType *uint64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// 告警阈值，大于0（目前暂定的值）
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`

	// 资源关联的IP列表，高防包未绑定时，传空数组，高防IP专业版传多个IP的数据
	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

type ModifyDDoSAlarmThresholdRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（shield表示棋牌；bgpip表示高防IP；bgp表示高防包；bgp-multip表示多ip高防包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID,字符串类型
	RsId *string `json:"RsId,omitempty" name:"RsId"`

	// 告警阈值类型，0-未设置，1-入流量，2-清洗流量
	AlarmType *uint64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// 告警阈值，大于0（目前暂定的值）
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`

	// 资源关联的IP列表，高防包未绑定时，传空数组，高防IP专业版传多个IP的数据
	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

func (r *ModifyDDoSAlarmThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSAlarmThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "RsId")
	delete(f, "AlarmType")
	delete(f, "AlarmThreshold")
	delete(f, "IpList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSAlarmThresholdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSAlarmThresholdResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDDoSAlarmThresholdResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSAlarmThresholdResponseParams `json:"Response"`
}

func (r *ModifyDDoSAlarmThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSAlarmThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSDefendStatusRequestParams struct {
	// 大禹子产品代号（bgp表示独享包；bgp-multip表示共享包；bgpip表示高防IP；net表示高防IP专业版；basic表示基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 防护状态值，取值[0（关闭），1（开启）]
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 关闭时长，单位小时，取值[0，1，2，3，4，5，6]；当Status=0表示关闭时，Hour必须大于0；
	Hour *int64 `json:"Hour,omitempty" name:"Hour"`

	// 资源ID；当Business不是基础防护时必须填写此字段；
	Id *string `json:"Id,omitempty" name:"Id"`

	// 基础防护的IP，只有当Business为基础防护时才需要填写此字段；
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 只有当Business为基础防护时才需要填写此字段，IP所属的产品类型，取值[public（CVM产品），bm（黑石产品），eni（弹性网卡），vpngw（VPN网关）， natgw（NAT网关），waf（Web应用安全产品），fpc（金融产品），gaap（GAAP产品）, other(托管IP)]
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// 只有当Business为基础防护时才需要填写此字段，IP所属的产品子类，取值[cvm（CVM），lb（负载均衡器），eni（弹性网卡），vpngw（VPN），natgw（NAT），waf（WAF），fpc（金融），gaap（GAAP），other（托管IP），eip（黑石弹性IP）]
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 只有当Business为基础防护时才需要填写此字段，IP所属的资源实例ID，当绑定新IP时必须填写此字段；例如是弹性网卡的IP，则InstanceId填写弹性网卡的ID(eni-*);
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 只有当Business为基础防护时才需要填写此字段，表示IP所属的地域，取值：
	// "bj":     华北地区(北京)
	// "cd":     西南地区(成都)
	// "cq":     西南地区(重庆)
	// "gz":     华南地区(广州)
	// "gzopen": 华南地区(广州Open)
	// "hk":     中国香港
	// "kr":     东南亚地区(首尔)
	// "sh":     华东地区(上海)
	// "shjr":   华东地区(上海金融)
	// "szjr":   华南地区(深圳金融)
	// "sg":     东南亚地区(新加坡)
	// "th":     东南亚地区(泰国)
	// "de":     欧洲地区(德国)
	// "usw":    美国西部（硅谷）
	// "ca":     北美地区(多伦多)
	// "jp":     日本
	// "hzec":   杭州
	// "in":     印度
	// "use":    美东地区（弗吉尼亚）
	// "ru":     俄罗斯
	// "tpe":    中国台湾
	// "nj":     南京
	IPRegion *string `json:"IPRegion,omitempty" name:"IPRegion"`
}

type ModifyDDoSDefendStatusRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgp表示独享包；bgp-multip表示共享包；bgpip表示高防IP；net表示高防IP专业版；basic表示基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 防护状态值，取值[0（关闭），1（开启）]
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 关闭时长，单位小时，取值[0，1，2，3，4，5，6]；当Status=0表示关闭时，Hour必须大于0；
	Hour *int64 `json:"Hour,omitempty" name:"Hour"`

	// 资源ID；当Business不是基础防护时必须填写此字段；
	Id *string `json:"Id,omitempty" name:"Id"`

	// 基础防护的IP，只有当Business为基础防护时才需要填写此字段；
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 只有当Business为基础防护时才需要填写此字段，IP所属的产品类型，取值[public（CVM产品），bm（黑石产品），eni（弹性网卡），vpngw（VPN网关）， natgw（NAT网关），waf（Web应用安全产品），fpc（金融产品），gaap（GAAP产品）, other(托管IP)]
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// 只有当Business为基础防护时才需要填写此字段，IP所属的产品子类，取值[cvm（CVM），lb（负载均衡器），eni（弹性网卡），vpngw（VPN），natgw（NAT），waf（WAF），fpc（金融），gaap（GAAP），other（托管IP），eip（黑石弹性IP）]
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 只有当Business为基础防护时才需要填写此字段，IP所属的资源实例ID，当绑定新IP时必须填写此字段；例如是弹性网卡的IP，则InstanceId填写弹性网卡的ID(eni-*);
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 只有当Business为基础防护时才需要填写此字段，表示IP所属的地域，取值：
	// "bj":     华北地区(北京)
	// "cd":     西南地区(成都)
	// "cq":     西南地区(重庆)
	// "gz":     华南地区(广州)
	// "gzopen": 华南地区(广州Open)
	// "hk":     中国香港
	// "kr":     东南亚地区(首尔)
	// "sh":     华东地区(上海)
	// "shjr":   华东地区(上海金融)
	// "szjr":   华南地区(深圳金融)
	// "sg":     东南亚地区(新加坡)
	// "th":     东南亚地区(泰国)
	// "de":     欧洲地区(德国)
	// "usw":    美国西部（硅谷）
	// "ca":     北美地区(多伦多)
	// "jp":     日本
	// "hzec":   杭州
	// "in":     印度
	// "use":    美东地区（弗吉尼亚）
	// "ru":     俄罗斯
	// "tpe":    中国台湾
	// "nj":     南京
	IPRegion *string `json:"IPRegion,omitempty" name:"IPRegion"`
}

func (r *ModifyDDoSDefendStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSDefendStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Status")
	delete(f, "Hour")
	delete(f, "Id")
	delete(f, "Ip")
	delete(f, "BizType")
	delete(f, "DeviceType")
	delete(f, "InstanceId")
	delete(f, "IPRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSDefendStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSDefendStatusResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDDoSDefendStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSDefendStatusResponseParams `json:"Response"`
}

func (r *ModifyDDoSDefendStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSDefendStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSLevelRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// =get表示读取防护等级；=set表示修改防护等级
	Method *string `json:"Method,omitempty" name:"Method"`

	// 防护等级，取值[low,middle,high]；当Method=set时必填
	DDoSLevel *string `json:"DDoSLevel,omitempty" name:"DDoSLevel"`
}

type ModifyDDoSLevelRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// =get表示读取防护等级；=set表示修改防护等级
	Method *string `json:"Method,omitempty" name:"Method"`

	// 防护等级，取值[low,middle,high]；当Method=set时必填
	DDoSLevel *string `json:"DDoSLevel,omitempty" name:"DDoSLevel"`
}

func (r *ModifyDDoSLevelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSLevelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Method")
	delete(f, "DDoSLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSLevelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSLevelResponseParams struct {
	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 防护等级，取值[low,middle,high]
	DDoSLevel *string `json:"DDoSLevel,omitempty" name:"DDoSLevel"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDDoSLevelResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSLevelResponseParams `json:"Response"`
}

func (r *ModifyDDoSLevelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSPolicyCaseRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 策略场景ID
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`

	// 开发平台，取值[PC（PC客户端）， MOBILE（移动端）， TV（电视端）， SERVER（主机）]
	PlatformTypes []*string `json:"PlatformTypes,omitempty" name:"PlatformTypes"`

	// 细分品类，取值[WEB（网站）， GAME（游戏）， APP（应用）， OTHER（其他）]
	AppType *string `json:"AppType,omitempty" name:"AppType"`

	// 应用协议，取值[tcp（TCP协议），udp（UDP协议），icmp（ICMP协议），all（其他协议）]
	AppProtocols []*string `json:"AppProtocols,omitempty" name:"AppProtocols"`

	// TCP业务起始端口，取值(0, 65535]
	TcpSportStart *string `json:"TcpSportStart,omitempty" name:"TcpSportStart"`

	// TCP业务结束端口，取值(0, 65535]，必须大于等于TCP业务起始端口
	TcpSportEnd *string `json:"TcpSportEnd,omitempty" name:"TcpSportEnd"`

	// UDP业务起始端口，取值范围(0, 65535]
	UdpSportStart *string `json:"UdpSportStart,omitempty" name:"UdpSportStart"`

	// UDP业务结束端口，取值范围(0, 65535)，必须大于等于UDP业务起始端口
	UdpSportEnd *string `json:"UdpSportEnd,omitempty" name:"UdpSportEnd"`

	// 是否有海外客户，取值[no（没有）, yes（有）]
	HasAbroad *string `json:"HasAbroad,omitempty" name:"HasAbroad"`

	// 是否会主动对外发起TCP请求，取值[no（不会）, yes（会）]
	HasInitiateTcp *string `json:"HasInitiateTcp,omitempty" name:"HasInitiateTcp"`

	// 是否会主动对外发起UDP业务请求，取值[no（不会）, yes（会）]
	HasInitiateUdp *string `json:"HasInitiateUdp,omitempty" name:"HasInitiateUdp"`

	// 主动发起TCP请求的端口，取值范围(0, 65535]
	PeerTcpPort *string `json:"PeerTcpPort,omitempty" name:"PeerTcpPort"`

	// 主动发起UDP请求的端口，取值范围(0, 65535]
	PeerUdpPort *string `json:"PeerUdpPort,omitempty" name:"PeerUdpPort"`

	// TCP载荷的固定特征码，字符串长度小于512
	TcpFootprint *string `json:"TcpFootprint,omitempty" name:"TcpFootprint"`

	// UDP载荷的固定特征码，字符串长度小于512
	UdpFootprint *string `json:"UdpFootprint,omitempty" name:"UdpFootprint"`

	// Web业务的API的URL
	WebApiUrl []*string `json:"WebApiUrl,omitempty" name:"WebApiUrl"`

	// TCP业务报文长度最小值，取值范围(0, 1500)
	MinTcpPackageLen *string `json:"MinTcpPackageLen,omitempty" name:"MinTcpPackageLen"`

	// TCP业务报文长度最大值，取值范围(0, 1500)，必须大于等于TCP业务报文长度最小值
	MaxTcpPackageLen *string `json:"MaxTcpPackageLen,omitempty" name:"MaxTcpPackageLen"`

	// UDP业务报文长度最小值，取值范围(0, 1500)
	MinUdpPackageLen *string `json:"MinUdpPackageLen,omitempty" name:"MinUdpPackageLen"`

	// UDP业务报文长度最大值，取值范围(0, 1500)，必须大于等于UDP业务报文长度最小值
	MaxUdpPackageLen *string `json:"MaxUdpPackageLen,omitempty" name:"MaxUdpPackageLen"`

	// 是否有VPN业务，取值[no（没有）, yes（有）]
	HasVPN *string `json:"HasVPN,omitempty" name:"HasVPN"`

	// TCP业务端口列表，同时支持单个端口和端口段，字符串格式，例如：80,443,700-800,53,1000-3000
	TcpPortList *string `json:"TcpPortList,omitempty" name:"TcpPortList"`

	// UDP业务端口列表，同时支持单个端口和端口段，字符串格式，例如：80,443,700-800,53,1000-3000
	UdpPortList *string `json:"UdpPortList,omitempty" name:"UdpPortList"`
}

type ModifyDDoSPolicyCaseRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 策略场景ID
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`

	// 开发平台，取值[PC（PC客户端）， MOBILE（移动端）， TV（电视端）， SERVER（主机）]
	PlatformTypes []*string `json:"PlatformTypes,omitempty" name:"PlatformTypes"`

	// 细分品类，取值[WEB（网站）， GAME（游戏）， APP（应用）， OTHER（其他）]
	AppType *string `json:"AppType,omitempty" name:"AppType"`

	// 应用协议，取值[tcp（TCP协议），udp（UDP协议），icmp（ICMP协议），all（其他协议）]
	AppProtocols []*string `json:"AppProtocols,omitempty" name:"AppProtocols"`

	// TCP业务起始端口，取值(0, 65535]
	TcpSportStart *string `json:"TcpSportStart,omitempty" name:"TcpSportStart"`

	// TCP业务结束端口，取值(0, 65535]，必须大于等于TCP业务起始端口
	TcpSportEnd *string `json:"TcpSportEnd,omitempty" name:"TcpSportEnd"`

	// UDP业务起始端口，取值范围(0, 65535]
	UdpSportStart *string `json:"UdpSportStart,omitempty" name:"UdpSportStart"`

	// UDP业务结束端口，取值范围(0, 65535)，必须大于等于UDP业务起始端口
	UdpSportEnd *string `json:"UdpSportEnd,omitempty" name:"UdpSportEnd"`

	// 是否有海外客户，取值[no（没有）, yes（有）]
	HasAbroad *string `json:"HasAbroad,omitempty" name:"HasAbroad"`

	// 是否会主动对外发起TCP请求，取值[no（不会）, yes（会）]
	HasInitiateTcp *string `json:"HasInitiateTcp,omitempty" name:"HasInitiateTcp"`

	// 是否会主动对外发起UDP业务请求，取值[no（不会）, yes（会）]
	HasInitiateUdp *string `json:"HasInitiateUdp,omitempty" name:"HasInitiateUdp"`

	// 主动发起TCP请求的端口，取值范围(0, 65535]
	PeerTcpPort *string `json:"PeerTcpPort,omitempty" name:"PeerTcpPort"`

	// 主动发起UDP请求的端口，取值范围(0, 65535]
	PeerUdpPort *string `json:"PeerUdpPort,omitempty" name:"PeerUdpPort"`

	// TCP载荷的固定特征码，字符串长度小于512
	TcpFootprint *string `json:"TcpFootprint,omitempty" name:"TcpFootprint"`

	// UDP载荷的固定特征码，字符串长度小于512
	UdpFootprint *string `json:"UdpFootprint,omitempty" name:"UdpFootprint"`

	// Web业务的API的URL
	WebApiUrl []*string `json:"WebApiUrl,omitempty" name:"WebApiUrl"`

	// TCP业务报文长度最小值，取值范围(0, 1500)
	MinTcpPackageLen *string `json:"MinTcpPackageLen,omitempty" name:"MinTcpPackageLen"`

	// TCP业务报文长度最大值，取值范围(0, 1500)，必须大于等于TCP业务报文长度最小值
	MaxTcpPackageLen *string `json:"MaxTcpPackageLen,omitempty" name:"MaxTcpPackageLen"`

	// UDP业务报文长度最小值，取值范围(0, 1500)
	MinUdpPackageLen *string `json:"MinUdpPackageLen,omitempty" name:"MinUdpPackageLen"`

	// UDP业务报文长度最大值，取值范围(0, 1500)，必须大于等于UDP业务报文长度最小值
	MaxUdpPackageLen *string `json:"MaxUdpPackageLen,omitempty" name:"MaxUdpPackageLen"`

	// 是否有VPN业务，取值[no（没有）, yes（有）]
	HasVPN *string `json:"HasVPN,omitempty" name:"HasVPN"`

	// TCP业务端口列表，同时支持单个端口和端口段，字符串格式，例如：80,443,700-800,53,1000-3000
	TcpPortList *string `json:"TcpPortList,omitempty" name:"TcpPortList"`

	// UDP业务端口列表，同时支持单个端口和端口段，字符串格式，例如：80,443,700-800,53,1000-3000
	UdpPortList *string `json:"UdpPortList,omitempty" name:"UdpPortList"`
}

func (r *ModifyDDoSPolicyCaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSPolicyCaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "SceneId")
	delete(f, "PlatformTypes")
	delete(f, "AppType")
	delete(f, "AppProtocols")
	delete(f, "TcpSportStart")
	delete(f, "TcpSportEnd")
	delete(f, "UdpSportStart")
	delete(f, "UdpSportEnd")
	delete(f, "HasAbroad")
	delete(f, "HasInitiateTcp")
	delete(f, "HasInitiateUdp")
	delete(f, "PeerTcpPort")
	delete(f, "PeerUdpPort")
	delete(f, "TcpFootprint")
	delete(f, "UdpFootprint")
	delete(f, "WebApiUrl")
	delete(f, "MinTcpPackageLen")
	delete(f, "MaxTcpPackageLen")
	delete(f, "MinUdpPackageLen")
	delete(f, "MaxUdpPackageLen")
	delete(f, "HasVPN")
	delete(f, "TcpPortList")
	delete(f, "UdpPortList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSPolicyCaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSPolicyCaseResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDDoSPolicyCaseResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSPolicyCaseResponseParams `json:"Response"`
}

func (r *ModifyDDoSPolicyCaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSPolicyCaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSPolicyNameRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type ModifyDDoSPolicyNameRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ModifyDDoSPolicyNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSPolicyNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "PolicyId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSPolicyNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSPolicyNameResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDDoSPolicyNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSPolicyNameResponseParams `json:"Response"`
}

func (r *ModifyDDoSPolicyNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSPolicyNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSPolicyRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 协议禁用，必须填写且数组长度必须为1
	DropOptions []*DDoSPolicyDropOption `json:"DropOptions,omitempty" name:"DropOptions"`

	// 端口禁用，当没有禁用端口时填空数组
	PortLimits []*DDoSPolicyPortLimit `json:"PortLimits,omitempty" name:"PortLimits"`

	// IP黑白名单，当没有IP黑白名单时填空数组
	IpAllowDenys []*IpBlackWhite `json:"IpAllowDenys,omitempty" name:"IpAllowDenys"`

	// 报文过滤，当没有报文过滤时填空数组
	PacketFilters []*DDoSPolicyPacketFilter `json:"PacketFilters,omitempty" name:"PacketFilters"`

	// 水印策略参数，当没有启用水印功能时填空数组，最多只能传一条水印策略（即数组大小不超过1）
	WaterPrint []*WaterPrintPolicy `json:"WaterPrint,omitempty" name:"WaterPrint"`
}

type ModifyDDoSPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 协议禁用，必须填写且数组长度必须为1
	DropOptions []*DDoSPolicyDropOption `json:"DropOptions,omitempty" name:"DropOptions"`

	// 端口禁用，当没有禁用端口时填空数组
	PortLimits []*DDoSPolicyPortLimit `json:"PortLimits,omitempty" name:"PortLimits"`

	// IP黑白名单，当没有IP黑白名单时填空数组
	IpAllowDenys []*IpBlackWhite `json:"IpAllowDenys,omitempty" name:"IpAllowDenys"`

	// 报文过滤，当没有报文过滤时填空数组
	PacketFilters []*DDoSPolicyPacketFilter `json:"PacketFilters,omitempty" name:"PacketFilters"`

	// 水印策略参数，当没有启用水印功能时填空数组，最多只能传一条水印策略（即数组大小不超过1）
	WaterPrint []*WaterPrintPolicy `json:"WaterPrint,omitempty" name:"WaterPrint"`
}

func (r *ModifyDDoSPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "PolicyId")
	delete(f, "DropOptions")
	delete(f, "PortLimits")
	delete(f, "IpAllowDenys")
	delete(f, "PacketFilters")
	delete(f, "WaterPrint")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSPolicyResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDDoSPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSPolicyResponseParams `json:"Response"`
}

func (r *ModifyDDoSPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSSwitchRequestParams struct {
	// 大禹子产品代号（basic表示基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// =get表示读取DDoS防护状态；=set表示修改DDoS防护状态；
	Method *string `json:"Method,omitempty" name:"Method"`

	// 基础防护的IP，只有当Business为基础防护时才需要填写此字段；
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 只有当Business为基础防护时才需要填写此字段，IP所属的产品类型，取值[public（CVM产品），bm（黑石产品），eni（弹性网卡），vpngw（VPN网关）， natgw（NAT网关），waf（Web应用安全产品），fpc（金融产品），gaap（GAAP产品）, other(托管IP)]
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// 只有当Business为基础防护时才需要填写此字段，IP所属的产品子类，取值[cvm（CVM），lb（负载均衡器），eni（弹性网卡），vpngw（VPN），natgw（NAT），waf（WAF），fpc（金融），gaap（GAAP），other（托管IP），eip（黑石弹性IP）]
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 只有当Business为基础防护时才需要填写此字段，IP所属的资源实例ID，当绑定新IP时必须填写此字段；例如是弹性网卡的IP，则InstanceId填写弹性网卡的ID(eni-*);
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 只有当Business为基础防护时才需要填写此字段，表示IP所属的地域，取值：
	// "bj":     华北地区(北京)
	// "cd":     西南地区(成都)
	// "cq":     西南地区(重庆)
	// "gz":     华南地区(广州)
	// "gzopen": 华南地区(广州Open)
	// "hk":     中国香港
	// "kr":     东南亚地区(首尔)
	// "sh":     华东地区(上海)
	// "shjr":   华东地区(上海金融)
	// "szjr":   华南地区(深圳金融)
	// "sg":     东南亚地区(新加坡)
	// "th":     东南亚地区(泰国)
	// "de":     欧洲地区(德国)
	// "usw":    美国西部（硅谷）
	// "ca":     北美地区(多伦多)
	// "jp":     日本
	// "hzec":   杭州
	// "in":     印度
	// "use":    美东地区（弗吉尼亚）
	// "ru":     俄罗斯
	// "tpe":    中国台湾
	// "nj":     南京
	IPRegion *string `json:"IPRegion,omitempty" name:"IPRegion"`

	// 可选字段，防护状态值，取值[0（关闭），1（开启）]；当Method为get时可以不填写此字段；
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type ModifyDDoSSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（basic表示基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// =get表示读取DDoS防护状态；=set表示修改DDoS防护状态；
	Method *string `json:"Method,omitempty" name:"Method"`

	// 基础防护的IP，只有当Business为基础防护时才需要填写此字段；
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 只有当Business为基础防护时才需要填写此字段，IP所属的产品类型，取值[public（CVM产品），bm（黑石产品），eni（弹性网卡），vpngw（VPN网关）， natgw（NAT网关），waf（Web应用安全产品），fpc（金融产品），gaap（GAAP产品）, other(托管IP)]
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// 只有当Business为基础防护时才需要填写此字段，IP所属的产品子类，取值[cvm（CVM），lb（负载均衡器），eni（弹性网卡），vpngw（VPN），natgw（NAT），waf（WAF），fpc（金融），gaap（GAAP），other（托管IP），eip（黑石弹性IP）]
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 只有当Business为基础防护时才需要填写此字段，IP所属的资源实例ID，当绑定新IP时必须填写此字段；例如是弹性网卡的IP，则InstanceId填写弹性网卡的ID(eni-*);
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 只有当Business为基础防护时才需要填写此字段，表示IP所属的地域，取值：
	// "bj":     华北地区(北京)
	// "cd":     西南地区(成都)
	// "cq":     西南地区(重庆)
	// "gz":     华南地区(广州)
	// "gzopen": 华南地区(广州Open)
	// "hk":     中国香港
	// "kr":     东南亚地区(首尔)
	// "sh":     华东地区(上海)
	// "shjr":   华东地区(上海金融)
	// "szjr":   华南地区(深圳金融)
	// "sg":     东南亚地区(新加坡)
	// "th":     东南亚地区(泰国)
	// "de":     欧洲地区(德国)
	// "usw":    美国西部（硅谷）
	// "ca":     北美地区(多伦多)
	// "jp":     日本
	// "hzec":   杭州
	// "in":     印度
	// "use":    美东地区（弗吉尼亚）
	// "ru":     俄罗斯
	// "tpe":    中国台湾
	// "nj":     南京
	IPRegion *string `json:"IPRegion,omitempty" name:"IPRegion"`

	// 可选字段，防护状态值，取值[0（关闭），1（开启）]；当Method为get时可以不填写此字段；
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDDoSSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Method")
	delete(f, "Ip")
	delete(f, "BizType")
	delete(f, "DeviceType")
	delete(f, "InstanceId")
	delete(f, "IPRegion")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSSwitchResponseParams struct {
	// 当前防护状态值，取值[0（关闭），1（开启）]
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDDoSSwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSSwitchResponseParams `json:"Response"`
}

func (r *ModifyDDoSSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSThresholdRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// DDoS清洗阈值，取值[0, 60, 80, 100, 150, 200, 250, 300, 400, 500, 700, 1000];
	// 当设置值为0时，表示采用默认值；
	Threshold *uint64 `json:"Threshold,omitempty" name:"Threshold"`
}

type ModifyDDoSThresholdRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// DDoS清洗阈值，取值[0, 60, 80, 100, 150, 200, 250, 300, 400, 500, 700, 1000];
	// 当设置值为0时，表示采用默认值；
	Threshold *uint64 `json:"Threshold,omitempty" name:"Threshold"`
}

func (r *ModifyDDoSThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Threshold")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSThresholdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSThresholdResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDDoSThresholdResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSThresholdResponseParams `json:"Response"`
}

func (r *ModifyDDoSThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSWaterKeyRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 密钥操作，取值：[add（添加），delete（删除），open（开启），close（关闭），get（获取密钥）]
	Method *string `json:"Method,omitempty" name:"Method"`

	// 密钥ID，当添加密钥操作时可以不填或填0，其他操作时必须填写；
	KeyId *uint64 `json:"KeyId,omitempty" name:"KeyId"`
}

type ModifyDDoSWaterKeyRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 密钥操作，取值：[add（添加），delete（删除），open（开启），close（关闭），get（获取密钥）]
	Method *string `json:"Method,omitempty" name:"Method"`

	// 密钥ID，当添加密钥操作时可以不填或填0，其他操作时必须填写；
	KeyId *uint64 `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *ModifyDDoSWaterKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSWaterKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "PolicyId")
	delete(f, "Method")
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSWaterKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSWaterKeyResponseParams struct {
	// 水印密钥列表
	KeyList []*WaterPrintKey `json:"KeyList,omitempty" name:"KeyList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDDoSWaterKeyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSWaterKeyResponseParams `json:"Response"`
}

func (r *ModifyDDoSWaterKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSWaterKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyElasticLimitRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 弹性防护阈值，取值[0 10000 20000 30000 40000 50000 60000 70000 80000 90000 100000 120000 150000 200000 250000 300000 400000 600000 800000 220000 310000 110000 270000 610000]
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type ModifyElasticLimitRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 弹性防护阈值，取值[0 10000 20000 30000 40000 50000 60000 70000 80000 90000 100000 120000 150000 200000 250000 300000 400000 600000 800000 220000 310000 110000 270000 610000]
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ModifyElasticLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyElasticLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyElasticLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyElasticLimitResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyElasticLimitResponse struct {
	*tchttp.BaseResponse
	Response *ModifyElasticLimitResponseParams `json:"Response"`
}

func (r *ModifyElasticLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyElasticLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4HealthRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 健康检查参数数组
	Healths []*L4RuleHealth `json:"Healths,omitempty" name:"Healths"`
}

type ModifyL4HealthRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 健康检查参数数组
	Healths []*L4RuleHealth `json:"Healths,omitempty" name:"Healths"`
}

func (r *ModifyL4HealthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4HealthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Healths")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL4HealthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4HealthResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyL4HealthResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL4HealthResponseParams `json:"Response"`
}

func (r *ModifyL4HealthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4HealthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4KeepTimeRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 会话保持开关，取值[0(会话保持关闭)，1(会话保持开启)]
	KeepEnable *uint64 `json:"KeepEnable,omitempty" name:"KeepEnable"`

	// 会话保持时间，单位秒
	KeepTime *uint64 `json:"KeepTime,omitempty" name:"KeepTime"`
}

type ModifyL4KeepTimeRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 会话保持开关，取值[0(会话保持关闭)，1(会话保持开启)]
	KeepEnable *uint64 `json:"KeepEnable,omitempty" name:"KeepEnable"`

	// 会话保持时间，单位秒
	KeepTime *uint64 `json:"KeepTime,omitempty" name:"KeepTime"`
}

func (r *ModifyL4KeepTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4KeepTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleId")
	delete(f, "KeepEnable")
	delete(f, "KeepTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL4KeepTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4KeepTimeResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyL4KeepTimeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL4KeepTimeResponseParams `json:"Response"`
}

func (r *ModifyL4KeepTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4KeepTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4RulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则
	Rule *L4RuleEntry `json:"Rule,omitempty" name:"Rule"`
}

type ModifyL4RulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则
	Rule *L4RuleEntry `json:"Rule,omitempty" name:"Rule"`
}

func (r *ModifyL4RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL4RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4RulesResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyL4RulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL4RulesResponseParams `json:"Response"`
}

func (r *ModifyL4RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL7RulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则
	Rule *L7RuleEntry `json:"Rule,omitempty" name:"Rule"`
}

type ModifyL7RulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则
	Rule *L7RuleEntry `json:"Rule,omitempty" name:"Rule"`
}

func (r *ModifyL7RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL7RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL7RulesResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyL7RulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL7RulesResponseParams `json:"Response"`
}

func (r *ModifyL7RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetReturnSwitchRequestParams struct {
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Status 表示回切开关，0: 关闭， 1:打开
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 回切时长，单位：小时，取值[0,1,2,3,4,5,6;]当status=1时必选填写Hour>0
	Hour *uint64 `json:"Hour,omitempty" name:"Hour"`
}

type ModifyNetReturnSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Status 表示回切开关，0: 关闭， 1:打开
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 回切时长，单位：小时，取值[0,1,2,3,4,5,6;]当status=1时必选填写Hour>0
	Hour *uint64 `json:"Hour,omitempty" name:"Hour"`
}

func (r *ModifyNetReturnSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetReturnSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Status")
	delete(f, "Hour")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNetReturnSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetReturnSwitchResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyNetReturnSwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNetReturnSwitchResponseParams `json:"Response"`
}

func (r *ModifyNetReturnSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetReturnSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNewDomainRulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 域名转发规则
	Rule *NewL7RuleEntry `json:"Rule,omitempty" name:"Rule"`
}

type ModifyNewDomainRulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 域名转发规则
	Rule *NewL7RuleEntry `json:"Rule,omitempty" name:"Rule"`
}

func (r *ModifyNewDomainRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNewDomainRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNewDomainRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNewDomainRulesResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyNewDomainRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNewDomainRulesResponseParams `json:"Response"`
}

func (r *ModifyNewDomainRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNewDomainRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNewL4RuleRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 转发规则
	Rule *L4RuleEntry `json:"Rule,omitempty" name:"Rule"`
}

type ModifyNewL4RuleRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 转发规则
	Rule *L4RuleEntry `json:"Rule,omitempty" name:"Rule"`
}

func (r *ModifyNewL4RuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNewL4RuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNewL4RuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNewL4RuleResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyNewL4RuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNewL4RuleResponseParams `json:"Response"`
}

func (r *ModifyNewL4RuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNewL4RuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResBindDDoSPolicyRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 绑定或解绑，bind表示绑定策略，unbind表示解绑策略
	Method *string `json:"Method,omitempty" name:"Method"`
}

type ModifyResBindDDoSPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 绑定或解绑，bind表示绑定策略，unbind表示解绑策略
	Method *string `json:"Method,omitempty" name:"Method"`
}

func (r *ModifyResBindDDoSPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResBindDDoSPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "PolicyId")
	delete(f, "Method")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResBindDDoSPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResBindDDoSPolicyResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyResBindDDoSPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResBindDDoSPolicyResponseParams `json:"Response"`
}

func (r *ModifyResBindDDoSPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResBindDDoSPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceRenewFlagRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版；shield表示棋牌盾；bgp表示独享包；bgp-multip表示共享包；insurance表示保险包；staticpack表示三网套餐包）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源Id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 自动续费标记（0手动续费；1自动续费；2到期不续费）
	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type ModifyResourceRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；net表示高防IP专业版；shield表示棋牌盾；bgp表示独享包；bgp-multip表示共享包；insurance表示保险包；staticpack表示三网套餐包）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源Id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 自动续费标记（0手动续费；1自动续费；2到期不续费）
	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

func (r *ModifyResourceRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceRenewFlagResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyResourceRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceRenewFlagResponseParams `json:"Response"`
}

func (r *ModifyResourceRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewL4RuleEntry struct {
	// 转发协议，取值[TCP, UDP]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 转发端口
	VirtualPort *uint64 `json:"VirtualPort,omitempty" name:"VirtualPort"`

	// 源站端口
	SourcePort *uint64 `json:"SourcePort,omitempty" name:"SourcePort"`

	// 会话保持时间，单位秒
	KeepTime *uint64 `json:"KeepTime,omitempty" name:"KeepTime"`

	// 回源列表
	SourceList []*L4RuleSource `json:"SourceList,omitempty" name:"SourceList"`

	// 负载均衡方式，取值[1(加权轮询)，2(源IP hash)]
	LbType *uint64 `json:"LbType,omitempty" name:"LbType"`

	// 会话保持开关，取值[0(会话保持关闭)，1(会话保持开启)]；
	KeepEnable *uint64 `json:"KeepEnable,omitempty" name:"KeepEnable"`

	// 回源方式，取值[1(域名回源)，2(IP回源)]
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 规则描述
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 移除水印状态，取值[0(关闭)，1(开启)]
	RemoveSwitch *uint64 `json:"RemoveSwitch,omitempty" name:"RemoveSwitch"`

	// 规则修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 对应地区信息
	Region *uint64 `json:"Region,omitempty" name:"Region"`

	// 绑定资源IP信息
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 绑定资源Id信息
	Id *string `json:"Id,omitempty" name:"Id"`
}

type NewL7RuleEntry struct {
	// 转发协议，取值[http, https]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 转发域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 回源方式，取值[1(域名回源)，2(IP回源)]
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 会话保持时间，单位秒
	KeepTime *uint64 `json:"KeepTime,omitempty" name:"KeepTime"`

	// 回源列表
	SourceList []*L4RuleSource `json:"SourceList,omitempty" name:"SourceList"`

	// 负载均衡方式，取值[1(加权轮询)]
	LbType *uint64 `json:"LbType,omitempty" name:"LbType"`

	// 会话保持开关，取值[0(会话保持关闭)，1(会话保持开启)]
	KeepEnable *uint64 `json:"KeepEnable,omitempty" name:"KeepEnable"`

	// 规则ID，当添加新规则时可以不用填写此字段；当修改或者删除规则时需要填写此字段；
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 证书来源，当转发协议为https时必须填，取值[2(腾讯云托管证书)]，当转发协议为http时也可以填0
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// 当证书来源为腾讯云托管证书时，此字段必须填写托管证书ID
	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`

	// 当证书来源为自有证书时，此字段必须填写证书内容；(因已不再支持自有证书，此字段已弃用，请不用填写此字段)
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// 当证书来源为自有证书时，此字段必须填写证书密钥；(因已不再支持自有证书，此字段已弃用，请不用填写此字段)
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// 规则描述
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则状态，取值[0(规则配置成功)，1(规则配置生效中)，2(规则配置失败)，3(规则删除生效中)，5(规则删除失败)，6(规则等待配置)，7(规则等待删除)，8(规则待配置证书)]
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// cc防护状态，取值[0(关闭), 1(开启)]
	CCStatus *uint64 `json:"CCStatus,omitempty" name:"CCStatus"`

	// HTTPS协议的CC防护状态，取值[0(关闭), 1(开启)]
	CCEnable *uint64 `json:"CCEnable,omitempty" name:"CCEnable"`

	// HTTPS协议的CC防护阈值
	CCThreshold *uint64 `json:"CCThreshold,omitempty" name:"CCThreshold"`

	// HTTPS协议的CC防护等级
	CCLevel *string `json:"CCLevel,omitempty" name:"CCLevel"`

	// 区域码
	Region *uint64 `json:"Region,omitempty" name:"Region"`

	// 资源Id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源Ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 是否开启Https协议使用Http回源，取值[0(关闭), 1(开启)]，不填写默认是关闭
	HttpsToHttpEnable *uint64 `json:"HttpsToHttpEnable,omitempty" name:"HttpsToHttpEnable"`

	// 接入端口值
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualPort *uint64 `json:"VirtualPort,omitempty" name:"VirtualPort"`
}

type OrderBy struct {
	// 排序字段名称，取值[
	// bandwidth（带宽），
	// overloadCount（超峰值次数）
	// ]
	Field *string `json:"Field,omitempty" name:"Field"`

	// 升降序，取值为[asc（升序），（升序），desc（降序）， DESC（降序）]
	Order *string `json:"Order,omitempty" name:"Order"`
}

type Paging struct {
	// 起始位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type ProtocolPort struct {
	// 协议（tcp；udp）
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`
}

type RegionInstanceCount struct {
	// 地域码
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域码（新规范）
	RegionV3 *string `json:"RegionV3,omitempty" name:"RegionV3"`

	// 资源实例数
	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type ResourceIp struct {
	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源的IP数组
	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

type SchedulingDomain struct {
	// 调度域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// BGP线路IP列表
	BGPIpList []*string `json:"BGPIpList,omitempty" name:"BGPIpList"`

	// 电信线路IP列表
	CTCCIpList []*string `json:"CTCCIpList,omitempty" name:"CTCCIpList"`

	// 联通线路IP列表
	CUCCIpList []*string `json:"CUCCIpList,omitempty" name:"CUCCIpList"`

	// 移动线路IP列表
	CMCCIpList []*string `json:"CMCCIpList,omitempty" name:"CMCCIpList"`

	// 海外线路IP列表
	OverseaIpList []*string `json:"OverseaIpList,omitempty" name:"OverseaIpList"`

	// 调度方式，当前仅支持优先级, 取值为priority
	Method *string `json:"Method,omitempty" name:"Method"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// ttl
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type SuccessCode struct {
	// 成功/错误码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 描述
	Message *string `json:"Message,omitempty" name:"Message"`
}

type WaterPrintKey struct {
	// 水印KeyID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 水印Key值
	KeyContent *string `json:"KeyContent,omitempty" name:"KeyContent"`

	// 水印Key的版本号
	KeyVersion *string `json:"KeyVersion,omitempty" name:"KeyVersion"`

	// 是否开启，取值[0（没有开启），1（已开启）]
	OpenStatus *uint64 `json:"OpenStatus,omitempty" name:"OpenStatus"`

	// 密钥生成时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type WaterPrintPolicy struct {
	// TCP端口段，例如["2000-3000","3500-4000"]
	TcpPortList []*string `json:"TcpPortList,omitempty" name:"TcpPortList"`

	// UDP端口段，例如["2000-3000","3500-4000"]
	UdpPortList []*string `json:"UdpPortList,omitempty" name:"UdpPortList"`

	// 水印偏移量，取值范围[0, 100)
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 是否自动剥离，取值[0（不自动剥离），1（自动剥离）]
	RemoveSwitch *uint64 `json:"RemoveSwitch,omitempty" name:"RemoveSwitch"`

	// 是否开启，取值[0（没有开启），1（已开启）]
	OpenStatus *uint64 `json:"OpenStatus,omitempty" name:"OpenStatus"`
}