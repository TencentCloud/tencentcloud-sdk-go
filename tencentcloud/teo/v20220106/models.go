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

package v20220106

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ApplicationProxy struct {

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 实例名称
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// 调度模式：
	// ip表示Anycast IP
	// domain表示CNAME
	PlatType *string `json:"PlatType,omitempty" name:"PlatType"`

	// 0关闭安全，1开启安全
	SecurityType *int64 `json:"SecurityType,omitempty" name:"SecurityType"`

	// 0关闭加速，1开启加速
	AccelerateType *int64 `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// 字段已经移至Rule.ForwardClientIp
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// 字段已经移至Rule.SessionPersist
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// 规则列表
	Rule []*ApplicationProxyRule `json:"Rule,omitempty" name:"Rule"`

	// 状态：
	// online：启用
	// offline：停用
	// progress：部署中
	// stopping：停用中
	// fail：部署失败/停用失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 调度信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleValue []*string `json:"ScheduleValue,omitempty" name:"ScheduleValue"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 站点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 会话保持时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// 服务类型
	// hostname：子域名
	// instance：实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// 七层实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

type ApplicationProxyRule struct {

	// 协议，取值为TCP或者UDP
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// 端口，支持格式：
	// 80：80端口
	// 81-90：81至90端口
	Port []*string `json:"Port,omitempty" name:"Port"`

	// 源站类型，取值：
	// custom：手动添加
	// origins：源站组
	// load_balancing：负载均衡
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 源站信息：
	// 当OriginType=custom时，表示多个：
	// IP:端口
	// 域名:端口
	// 当OriginType=origins时，包含一个元素，表示源站组ID
	// 当OriginType=load_balancing时，包含一个元素，表示负载均衡ID
	OriginValue []*string `json:"OriginValue,omitempty" name:"OriginValue"`

	// 规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 状态：
	// online：启用
	// offline：停用
	// progress：部署中
	// stopping：停用中
	// fail：部署失败/停用失败
	Status *string `json:"Status,omitempty" name:"Status"`

	// 传递客户端IP，当Proto=TCP时，取值：
	// TOA：TOA
	// PPV1: Proxy Protocol传递，协议版本V1
	// PPV2: Proxy Protocol传递，协议版本V2
	// OFF：不传递
	// 当Proto=UDP时，取值：
	// PPV2: Proxy Protocol传递，协议版本V2
	// OFF：不传递
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// 是否开启会话保持
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`
}

type CacheConfig struct {

	// 缓存配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cache *CacheConfigCache `json:"Cache,omitempty" name:"Cache"`

	// 不缓存配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoCache *CacheConfigNoCache `json:"NoCache,omitempty" name:"NoCache"`

	// 遵循源站配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowOrigin *CacheConfigFollowOrigin `json:"FollowOrigin,omitempty" name:"FollowOrigin"`
}

type CacheConfigCache struct {

	// 缓存配置开关
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 缓存过期时间设置
	// 单位为秒，最大可设置为 365 天
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheTime *int64 `json:"CacheTime,omitempty" name:"CacheTime"`

	// 是否开启强制缓存
	// 开启：on
	// 关闭：off
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreCacheControl *string `json:"IgnoreCacheControl,omitempty" name:"IgnoreCacheControl"`
}

type CacheConfigFollowOrigin struct {

	// 遵循源站配置开关
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type CacheConfigNoCache struct {

	// 不缓存配置开关
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type CacheKey struct {

	// 是否开启全路径缓存
	// on：开启全路径缓存（即关闭参数忽略）
	// off：关闭全路径缓存（即开启参数忽略）
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullUrlCache *string `json:"FullUrlCache,omitempty" name:"FullUrlCache"`

	// 是否忽略大小写缓存
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreCase *string `json:"IgnoreCase,omitempty" name:"IgnoreCase"`

	// CacheKey中包含请求参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryString *QueryString `json:"QueryString,omitempty" name:"QueryString"`
}

type CertFilter struct {

	// 过滤字段名，支持的列表如下:
	//  - host：域名。
	//  - certId: 证书ID
	//  - certAlias: 证书备用名
	//  - certType: default: 默认证书, upload: 上传证书, managed:腾讯云证书
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤字段值
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否启用模糊查询，仅支持过滤字段名host。
	// 模糊查询时，Value长度最大为1。
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}

type CertSort struct {

	// 排序字段，当前支持：
	// createTime，域名创建时间
	// certExpireTime，证书过期时间
	// certDeployTime,  证书部署时间
	Key *string `json:"Key,omitempty" name:"Key"`

	// asc/desc，默认desc。
	Sequence *string `json:"Sequence,omitempty" name:"Sequence"`
}

type CheckCertificateRequest struct {
	*tchttp.BaseRequest

	// 证书
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// 私钥
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
}

func (r *CheckCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Certificate")
	delete(f, "PrivateKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CheckCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClientIp struct {

	// 客户端IP头部配置开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 回源客户端IP请求头名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`
}

type CnameStatus struct {

	// 记录名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// CNAME 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// 状态
	// 生效：active
	// 不生效：moved
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type Compression struct {

	// 智能压缩配置开关
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type CreateApplicationProxyRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 四层代理名称
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// 调度模式：
	// ip表示Anycast IP
	// domain表示CNAME
	PlatType *string `json:"PlatType,omitempty" name:"PlatType"`

	// 0关闭安全，1开启安全
	SecurityType *int64 `json:"SecurityType,omitempty" name:"SecurityType"`

	// 0关闭加速，1开启加速
	AccelerateType *int64 `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// 字段已经移至Rule.ForwardClientIp
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// 字段已经移至Rule.SessionPersist
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// 规则详细信息
	Rule []*ApplicationProxyRule `json:"Rule,omitempty" name:"Rule"`

	// 会话保持时间，取值范围：30-3600，单位：秒
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// 服务类型
	// hostname：子域名
	// instance：实例
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`
}

func (r *CreateApplicationProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ZoneName")
	delete(f, "ProxyName")
	delete(f, "PlatType")
	delete(f, "SecurityType")
	delete(f, "AccelerateType")
	delete(f, "ForwardClientIp")
	delete(f, "SessionPersist")
	delete(f, "Rule")
	delete(f, "SessionPersistTime")
	delete(f, "ProxyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationProxyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 新增的四层代理应用ID
		ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 协议，取值为TCP或者UDP
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// 端口，支持格式：
	// 80：80端口
	// 81-90：81至90端口
	Port []*string `json:"Port,omitempty" name:"Port"`

	// 源站类型，取值：
	// custom：手动添加
	// origins：源站组
	// load_balancing：负载均衡
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 源站信息：
	// 当OriginType=custom时，表示多个：
	// IP:端口
	// 域名:端口
	// 当OriginType=origins时，包含一个元素，表示源站组ID
	// 当OriginType=load_balancing时，包含一个元素，表示负载均衡ID
	OriginValue []*string `json:"OriginValue,omitempty" name:"OriginValue"`

	// 传递客户端IP，当Proto=TCP时，取值：
	// TOA：TOA
	// PPV1: Proxy Protocol传递，协议版本V1
	// PPV2: Proxy Protocol传递，协议版本V2
	// OFF：不传递
	// 当Proto=UDP时，取值：
	// PPV2: Proxy Protocol传递，协议版本V2
	// OFF：不传递
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// 是否开启会话保持
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`
}

func (r *CreateApplicationProxyRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationProxyRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "Proto")
	delete(f, "Port")
	delete(f, "OriginType")
	delete(f, "OriginValue")
	delete(f, "ForwardClientIp")
	delete(f, "SessionPersist")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationProxyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationProxyRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则ID
		RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationProxyRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationProxyRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationProxyRulesRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 规则列表
	Rule []*ApplicationProxyRule `json:"Rule,omitempty" name:"Rule"`
}

func (r *CreateApplicationProxyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationProxyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationProxyRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationProxyRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 新增的规则ID数组
		RuleId []*string `json:"RuleId,omitempty" name:"RuleId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationProxyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationProxyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDnsRecordRequest struct {
	*tchttp.BaseRequest

	// 站点 ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 记录类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 记录名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 记录内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 代理模式，可选值：dns_only, cdn_only, secure_cdn
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 生存时间值
	Ttl *int64 `json:"Ttl,omitempty" name:"Ttl"`

	// 优先级
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
}

func (r *CreateDnsRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDnsRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Type")
	delete(f, "Name")
	delete(f, "Content")
	delete(f, "Mode")
	delete(f, "Ttl")
	delete(f, "Priority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDnsRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDnsRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录 ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// 记录类型
		Type *string `json:"Type,omitempty" name:"Type"`

		// 记录名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 记录内容
		Content *string `json:"Content,omitempty" name:"Content"`

		// 生存时间值
		Ttl *int64 `json:"Ttl,omitempty" name:"Ttl"`

		// 优先级
		Priority *int64 `json:"Priority,omitempty" name:"Priority"`

		// 代理模式
		Mode *string `json:"Mode,omitempty" name:"Mode"`

		// 解析状态
	// active: 生效
	// pending: 不生效
		Status *string `json:"Status,omitempty" name:"Status"`

		// 已锁定
		Locked *bool `json:"Locked,omitempty" name:"Locked"`

		// 创建时间
		CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

		// 修改时间
		ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

		// 站点 ID
		ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

		// 站点名称
		ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

		// CNAME 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		Cname *string `json:"Cname,omitempty" name:"Cname"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDnsRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDnsRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLoadBalancingRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名，填写@表示根域
	Host *string `json:"Host,omitempty" name:"Host"`

	// 代理模式：
	// dns_only: 仅DNS
	// proxied: 开启代理
	Type *string `json:"Type,omitempty" name:"Type"`

	// 使用的源站组ID
	OriginId []*string `json:"OriginId,omitempty" name:"OriginId"`

	// 当Type=dns_only表示DNS的TTL时间
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`
}

func (r *CreateLoadBalancingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Host")
	delete(f, "Type")
	delete(f, "OriginId")
	delete(f, "TTL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoadBalancingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateLoadBalancingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 负载均衡ID
		LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLoadBalancingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrefetchTaskRequest struct {
	*tchttp.BaseRequest

	// Zone ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 要预热的资源列表，每个元素格式类似如下:
	// http://www.example.com/example.txt
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// 是否对url进行encode
	// 若内容含有非 ASCII 字符集的字符，请开启此开关，编码转换（编码规则遵循 RFC3986）
	EncodeUrl *bool `json:"EncodeUrl,omitempty" name:"EncodeUrl"`

	// 附带的http头部信息
	Headers []*Header `json:"Headers,omitempty" name:"Headers"`
}

func (r *CreatePrefetchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrefetchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Targets")
	delete(f, "EncodeUrl")
	delete(f, "Headers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrefetchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrefetchTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 失败的任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		FailedList []*FailReason `json:"FailedList,omitempty" name:"FailedList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePrefetchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrefetchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePurgeTaskRequest struct {
	*tchttp.BaseRequest

	// Zone ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 类型，当前支持的类型：
	// - purge_url：URL
	// - purge_prefix：前缀
	// - purge_host：Hostname
	// - purge_all：全部缓存
	Type *string `json:"Type,omitempty" name:"Type"`

	// 要刷新的资源列表，每个元素格式依据Type而定
	// 1) Type = purge_host 时
	// 形如：www.example.com 或 foo.bar.example.com
	// 2) Type = purge_prefix 时
	// 形如：http://www.example.com/example
	// 3) Type = purge_url 时
	// 形如：https://www.example.com/example.jpg
	// 4）Type = purge_all 时
	// Targets可为空，不需要填写
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// 若有编码转换，仅清除编码转换后匹配的资源
	// 若内容含有非 ASCII 字符集的字符，请开启此开关，编码转换（编码规则遵循 RFC3986）
	EncodeUrl *bool `json:"EncodeUrl,omitempty" name:"EncodeUrl"`
}

func (r *CreatePurgeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePurgeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Type")
	delete(f, "Targets")
	delete(f, "EncodeUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePurgeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePurgeTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 失败的任务列表及原因
	// 注意：此字段可能返回 null，表示取不到有效值。
		FailedList []*FailReason `json:"FailedList,omitempty" name:"FailedList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePurgeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePurgeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateZoneRequest struct {
	*tchttp.BaseRequest

	// 站点名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 接入方式，默认full
	// - full NS接入
	// - partial CNAME接入
	Type *string `json:"Type,omitempty" name:"Type"`

	// 是否跳过站点历史解析记录扫描
	JumpStart *bool `json:"JumpStart,omitempty" name:"JumpStart"`
}

func (r *CreateZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "JumpStart")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 站点ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// 站点名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 站点接入方式
		Type *string `json:"Type,omitempty" name:"Type"`

		// 站点状态
	// - pending 未切换NS
	// - active NS 已切换到分配的 NS
	// - moved NS 从腾讯云切走
		Status *string `json:"Status,omitempty" name:"Status"`

		// 当前使用的 NS 列表
		OriginalNameServers []*string `json:"OriginalNameServers,omitempty" name:"OriginalNameServers"`

		// 给用户分配的 NS 列表
		NameServers []*string `json:"NameServers,omitempty" name:"NameServers"`

		// 站点创建时间
		CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

		// 站点更新时间
		ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataItem struct {

	// 时间
	Time *string `json:"Time,omitempty" name:"Time"`

	// 数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type DefaultServerCertInfo struct {

	// 服务器证书 ID, 默认证书ID, 或在 SSL 证书管理进行证书托管时自动生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 证书备注名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 证书类型:
	// default: 默认证书
	// upload:用户上传
	// managed:腾讯云托管
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 证书过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 证书生效时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EffectiveTime *string `json:"EffectiveTime,omitempty" name:"EffectiveTime"`

	// 证书公用名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommonName *string `json:"CommonName,omitempty" name:"CommonName"`

	// 证书SAN域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubjectAltName []*string `json:"SubjectAltName,omitempty" name:"SubjectAltName"`

	// 证书状态:
	// applying: 证书申请中
	// failed: 证书(申请)失败
	// processing: 证书部署中
	// deployed: 证书已部署
	// disabled: 证书被禁用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// Status为失败时,此字段返回失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type DeleteApplicationProxyRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

func (r *DeleteApplicationProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationProxyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 代理ID
		ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApplicationProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteApplicationProxyRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationProxyRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationProxyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationProxyRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则ID
		RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApplicationProxyRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationProxyRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDnsRecordsRequest struct {
	*tchttp.BaseRequest

	// 站点 ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 记录 ID
	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteDnsRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDnsRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDnsRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDnsRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录 ID
		Ids []*string `json:"Ids,omitempty" name:"Ids"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDnsRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDnsRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoadBalancingRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 负载均衡ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`
}

func (r *DeleteLoadBalancingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "LoadBalancingId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoadBalancingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoadBalancingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 负载均衡ID
		LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLoadBalancingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteZoneRequest struct {
	*tchttp.BaseRequest

	// 站点 ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 站点 ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationProxyDetailRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 实例ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

func (r *DescribeApplicationProxyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationProxyDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationProxyDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationProxyDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID
		ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

		// 实例名称
		ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

		// 调度模式：
	// ip表示Anycast IP
	// domain表示CNAME
		PlatType *string `json:"PlatType,omitempty" name:"PlatType"`

		// 0关闭安全，1开启安全
		SecurityType *int64 `json:"SecurityType,omitempty" name:"SecurityType"`

		// 0关闭加速，1开启加速
		AccelerateType *int64 `json:"AccelerateType,omitempty" name:"AccelerateType"`

		// 字段已经移至Rule.ForwardClientIp
		ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

		// 字段已经移至Rule.SessionPersist
		SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

		// 规则列表
		Rule []*ApplicationProxyRule `json:"Rule,omitempty" name:"Rule"`

		// 状态：
	// online：启用
	// offline：停用
	// progress：部署中
		Status *string `json:"Status,omitempty" name:"Status"`

		// 调度信息
		ScheduleValue []*string `json:"ScheduleValue,omitempty" name:"ScheduleValue"`

		// 更新时间
		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// 站点ID
		ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

		// 站点名称
		ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

		// 会话保持时间
		SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

		// 服务类型
	// hostname：子域名
	// instance：实例
		ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

		// 七层实例ID
		HostId *string `json:"HostId,omitempty" name:"HostId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationProxyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationProxyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationProxyRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 分页参数Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页参数Limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeApplicationProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationProxyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*ApplicationProxy `json:"Data,omitempty" name:"Data"`

		// 记录总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 当ZoneId不为空时，表示当前站点允许创建的实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		Quota *int64 `json:"Quota,omitempty" name:"Quota"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCnameStatusRequest struct {
	*tchttp.BaseRequest

	// 站点 ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 域名列表
	Names []*string `json:"Names,omitempty" name:"Names"`
}

func (r *DescribeCnameStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCnameStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCnameStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCnameStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态列表
		Status []*CnameStatus `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCnameStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCnameStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefaultCertificatesRequest struct {
	*tchttp.BaseRequest

	// Zone ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeDefaultCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultCertificatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDefaultCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefaultCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 默认证书列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		CertInfo []*DefaultServerCertInfo `json:"CertInfo,omitempty" name:"CertInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDefaultCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDnsDataRequest struct {
	*tchttp.BaseRequest

	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 过滤参数
	Filters []*DnsDataFilter `json:"Filters,omitempty" name:"Filters"`

	// 时间粒度，默认为1分钟粒度，服务端根据时间范围自适应。
	// 支持指定以下几种粒度：
	// min：1分钟粒度
	// 5min：5分钟粒度
	// hour：1小时粒度
	// day：天粒度
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

func (r *DescribeDnsDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDnsDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	delete(f, "Interval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDnsDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDnsDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 统计曲线数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*DataItem `json:"Data,omitempty" name:"Data"`

		// 时间粒度
	// 注意：此字段可能返回 null，表示取不到有效值。
		Interval *string `json:"Interval,omitempty" name:"Interval"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDnsDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDnsDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDnsRecordsRequest struct {
	*tchttp.BaseRequest

	// 查询条件过滤器
	Filters []*DnsRecordFilter `json:"Filters,omitempty" name:"Filters"`

	// 排序
	Order *string `json:"Order,omitempty" name:"Order"`

	// 可选值 asc, desc
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// 可选值 all, any
	Match *string `json:"Match,omitempty" name:"Match"`

	// 分页查询限制数目，默认为 100，最大可设置为 1000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 站点 ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeDnsRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDnsRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "Direction")
	delete(f, "Match")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDnsRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDnsRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数，用于分页查询
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// DNS 记录列表
		Records []*DnsRecord `json:"Records,omitempty" name:"Records"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDnsRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDnsRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDnssecRequest struct {
	*tchttp.BaseRequest

	// 站点 ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeDnssecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDnssecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDnssecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDnssecResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 站点 ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// 站点名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// DNSSEC 状态
	// - enabled 开启
	// - disabled 关闭
		Status *string `json:"Status,omitempty" name:"Status"`

		// DNSSEC 相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Dnssec *DnssecInfo `json:"Dnssec,omitempty" name:"Dnssec"`

		// 修改时间
		ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDnssecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDnssecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostsCertificateRequest struct {
	*tchttp.BaseRequest

	// Zone ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为 100，最大可设置为 1000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件过滤器
	Filters []*CertFilter `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Sort *CertSort `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeHostsCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostsCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostsCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostsCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数，用于分页查询
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 域名证书配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Hosts []*HostCertSetting `json:"Hosts,omitempty" name:"Hosts"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostsCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostsCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostsSettingRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为 100，最大可设置为 1000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 指定域名查询
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`
}

func (r *DescribeHostsSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostsSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Hosts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostsSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostsSettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名列表
		Hosts []*DetailHost `json:"Hosts,omitempty" name:"Hosts"`

		// 域名数量
		TotalNumber *int64 `json:"TotalNumber,omitempty" name:"TotalNumber"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostsSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostsSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdentificationRequest struct {
	*tchttp.BaseRequest

	// 站点名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeIdentificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdentificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIdentificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdentificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 站点名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 验证状态
	// - pending 验证中
	// - finished 验证完成
		Status *string `json:"Status,omitempty" name:"Status"`

		// 子域
		Subdomain *string `json:"Subdomain,omitempty" name:"Subdomain"`

		// 记录类型
		RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

		// 记录值
		RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`

		// 域名当前的 NS 记录
	// 注意：此字段可能返回 null，表示取不到有效值。
		OriginalNameServers []*string `json:"OriginalNameServers,omitempty" name:"OriginalNameServers"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdentificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdentificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancingDetailRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 负载均衡ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`
}

func (r *DescribeLoadBalancingDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancingDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "LoadBalancingId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancingDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancingDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 负载均衡ID
		LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

		// 站点ID
		ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

		// 子域名，填写@表示根域
		Host *string `json:"Host,omitempty" name:"Host"`

		// 代理模式：
	// dns_only: 仅DNS
	// proxied: 开启代理
		Type *string `json:"Type,omitempty" name:"Type"`

		// 当Type=dns_only表示DNS的TTL时间
		TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

		// 使用的源站组ID
		OriginId []*string `json:"OriginId,omitempty" name:"OriginId"`

		// 使用的源站信息
		Origin []*OriginGroup `json:"Origin,omitempty" name:"Origin"`

		// 更新时间
		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// 状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 调度域名
	// 注意：此字段可能返回 null，表示取不到有效值。
		Cname *string `json:"Cname,omitempty" name:"Cname"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoadBalancingDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancingDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancingRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 分页参数Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页参数Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤参数Host
	Host *string `json:"Host,omitempty" name:"Host"`

	// 过滤参数Host是否支持模糊匹配
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}

func (r *DescribeLoadBalancingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Host")
	delete(f, "Fuzzy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 负载均衡信息
		Data []*LoadBalancing `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoadBalancingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrefetchTasksRequest struct {
	*tchttp.BaseRequest

	// 任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询最大返回的结果条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询的状态
	// 允许的值为：processing、success、failed、timeout、invalid
	Statuses []*string `json:"Statuses,omitempty" name:"Statuses"`

	// zone id
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 查询的域名列表
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询的资源
	Target *string `json:"Target,omitempty" name:"Target"`
}

func (r *DescribePrefetchTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrefetchTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Statuses")
	delete(f, "ZoneId")
	delete(f, "Domains")
	delete(f, "Target")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrefetchTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrefetchTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 该查询条件总共条目数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 任务结果列表
		Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrefetchTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrefetchTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePurgeTasksRequest struct {
	*tchttp.BaseRequest

	// 任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询最大返回的结果条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询的状态
	// 允许的值为：processing、success、failed、timeout、invalid
	Statuses []*string `json:"Statuses,omitempty" name:"Statuses"`

	// zone id
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 查询的域名列表
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询内容
	Target *string `json:"Target,omitempty" name:"Target"`
}

func (r *DescribePurgeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "Type")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Statuses")
	delete(f, "ZoneId")
	delete(f, "Domains")
	delete(f, "Target")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePurgeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePurgeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 该查询条件总共条目数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 任务结果列表
		Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePurgeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneDetailsRequest struct {
	*tchttp.BaseRequest

	// 站点 ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeZoneDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 站点 ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// 站点名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 用户当前使用的 NS 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		OriginalNameServers []*string `json:"OriginalNameServers,omitempty" name:"OriginalNameServers"`

		// 腾讯云分配给用户的 NS 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		NameServers []*string `json:"NameServers,omitempty" name:"NameServers"`

		// 站点状态
	// - active：NS 已切换
	// - pending：NS 未切换
	// - moved：NS 已切走
	// - deactivated：被封禁
		Status *string `json:"Status,omitempty" name:"Status"`

		// 站点接入方式
	// - full：NS 接入
	// - partial：CNAME 接入
		Type *string `json:"Type,omitempty" name:"Type"`

		// 站点是否关闭
		Paused *bool `json:"Paused,omitempty" name:"Paused"`

		// 站点创建时间
		CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

		// 站点修改时间
		ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

		// 用户自定义 NS 信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		VanityNameServers *VanityNameServers `json:"VanityNameServers,omitempty" name:"VanityNameServers"`

		// 用户自定义 NS IP 信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		VanityNameServersIps []*VanityNameServersIps `json:"VanityNameServersIps,omitempty" name:"VanityNameServersIps"`

		// 是否开启 CNAME 加速
	// - enabled：开启
	// - disabled：关闭
		CnameSpeedUp *string `json:"CnameSpeedUp,omitempty" name:"CnameSpeedUp"`

		// cname切换验证状态
	// - finished 切换完成
	// - pending 切换验证中
	// 注意：此字段可能返回 null，表示取不到有效值。
		CnameStatus *string `json:"CnameStatus,omitempty" name:"CnameStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneSettingRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeZoneSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneSettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 缓存过期时间配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		Cache *CacheConfig `json:"Cache,omitempty" name:"Cache"`

		// 节点缓存键配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

		// 浏览器缓存配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

		// 离线缓存
	// 注意：此字段可能返回 null，表示取不到有效值。
		OfflineCache *OfflineCache `json:"OfflineCache,omitempty" name:"OfflineCache"`

		// Quic访问
	// 注意：此字段可能返回 null，表示取不到有效值。
		Quic *Quic `json:"Quic,omitempty" name:"Quic"`

		// POST请求传输配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		PostMaxSize *PostMaxSize `json:"PostMaxSize,omitempty" name:"PostMaxSize"`

		// 智能压缩配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		Compression *Compression `json:"Compression,omitempty" name:"Compression"`

		// http2回源配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitempty" name:"UpstreamHttp2"`

		// 访问协议强制https跳转配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

		// Https 加速配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		Https *Https `json:"Https,omitempty" name:"Https"`

		// 源站配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		Origin *Origin `json:"Origin,omitempty" name:"Origin"`

		// 动态加速配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		SmartRouting *SmartRouting `json:"SmartRouting,omitempty" name:"SmartRouting"`

		// 站点ID
		ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

		// 站点域名
		Zone *string `json:"Zone,omitempty" name:"Zone"`

		// WebSocket配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`

		// 客户端IP回源请求头配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClientIpHeader *ClientIp `json:"ClientIpHeader,omitempty" name:"ClientIpHeader"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest

	// 分页参数，页偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页参数，每页返回的站点个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件过滤器，复杂类型
	Filters []*ZoneFilter `json:"Filters,omitempty" name:"Filters"`
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
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的站点数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 站点详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Zones []*Zone `json:"Zones,omitempty" name:"Zones"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DetailHost struct {

	// 腾讯云账号ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 加速服务状态
	// process：部署中
	// online：已启动
	// offline：已关闭
	Status *string `json:"Status,omitempty" name:"Status"`

	// 域名
	Host *string `json:"Host,omitempty" name:"Host"`
}

type DnsDataFilter struct {

	// 参数名称，取值范围：
	// zone：站点名
	// host：域名
	// type：dns解析类型
	// code：dns返回状态码
	// area：解析服务器所在区域
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数值
	// 当Name=area时，Value取值范围：
	// 亚洲：Asia
	// 欧洲：Europe
	// 非洲：Africa
	// 大洋洲：Oceania
	// 美洲：Americas
	// 
	// 当Name=code时，Value取值范围：
	// NoError：成功的响应
	// NXDomain：只在权威域名服务器的响应消息中有效，标示请求中请求的域不存在
	// NotImp：域名服务器不支持请求的类型
	// Refused：域名服务器因为策略的原因拒绝执行请求的操作
	Value *string `json:"Value,omitempty" name:"Value"`

	// 参数值
	// 当Name=area时，Value取值范围：
	// 亚洲：Asia
	// 欧洲：Europe
	// 非洲：Africa
	// 大洋洲：Oceania
	// 美洲：Americas
	// 
	// 当Name=code时，Value取值范围：
	// NoError：成功的响应
	// NXDomain：只在权威域名服务器的响应消息中有效，标示请求中请求的域不存在
	// NotImp：域名服务器不支持请求的类型
	// Refused：域名服务器因为策略的原因拒绝执行请求的操作
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DnsRecord struct {

	// 记录 ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 记录类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 主机记录
	Name *string `json:"Name,omitempty" name:"Name"`

	// 记录值
	Content *string `json:"Content,omitempty" name:"Content"`

	// 代理模式
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// TTL 值
	Ttl *int64 `json:"Ttl,omitempty" name:"Ttl"`

	// 优先级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 创建时间
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 修改时间
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// 域名锁
	Locked *bool `json:"Locked,omitempty" name:"Locked"`

	// 站点 ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 解析状态
	// active: 生效
	// pending: 不生效
	Status *string `json:"Status,omitempty" name:"Status"`

	// CNAME 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// 域名是否开启了lb，四层，安全
	// - lb 负载均衡
	// - security 安全
	// - l4 四层
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainStatus []*string `json:"DomainStatus,omitempty" name:"DomainStatus"`
}

type DnsRecordFilter struct {

	// 过滤字段名，支持的列表如下：
	// - name: 站点名。
	// - status: 站点状态
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤字段值
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否启用模糊查询，仅支持过滤字段名为name。模糊查询时，Values长度最大为1
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}

type DnssecInfo struct {

	// 标志
	Flags *int64 `json:"Flags,omitempty" name:"Flags"`

	// 加密算法
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// 加密类型
	KeyType *string `json:"KeyType,omitempty" name:"KeyType"`

	// 摘要类型
	DigestType *string `json:"DigestType,omitempty" name:"DigestType"`

	// 摘要算法
	DigestAlgorithm *string `json:"DigestAlgorithm,omitempty" name:"DigestAlgorithm"`

	// 摘要信息
	Digest *string `json:"Digest,omitempty" name:"Digest"`

	// DS 记录值
	DS *string `json:"DS,omitempty" name:"DS"`

	// 密钥标签
	KeyTag *int64 `json:"KeyTag,omitempty" name:"KeyTag"`

	// 公钥
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
}

type DownloadL7LogsRequest struct {
	*tchttp.BaseRequest

	// 起始时间(需严格按照RFC3339标准传参)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间(需严格按照RFC3339标准传参)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每页展示条数
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 当前页
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 站点集合
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// 域名集合
	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

func (r *DownloadL7LogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadL7LogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "Zones")
	delete(f, "Domains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadL7LogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DownloadL7LogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 七层离线日志data
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*L7OfflineLog `json:"Data,omitempty" name:"Data"`

		// 页面大小
	// 注意：此字段可能返回 null，表示取不到有效值。
		PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

		// 页号
	// 注意：此字段可能返回 null，表示取不到有效值。
		PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

		// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Pages *int64 `json:"Pages,omitempty" name:"Pages"`

		// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadL7LogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadL7LogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FailReason struct {

	// 失败原因
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 处理失败的资源列表。
	// 该列表元素来源于输入参数中的Targets，因此格式和入参中的Targets保持一致
	Targets []*string `json:"Targets,omitempty" name:"Targets"`
}

type ForceRedirect struct {

	// 访问强制跳转配置开关
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 重定向状态码
	// 301
	// 302
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitempty" name:"RedirectStatusCode"`
}

type Header struct {

	// HTTP头部
	Name *string `json:"Name,omitempty" name:"Name"`

	// HTTP头部值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type HostCertSetting struct {

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 服务端证书配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertInfo []*ServerCertInfo `json:"CertInfo,omitempty" name:"CertInfo"`
}

type Hsts struct {

	// 是否开启，on或off。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// MaxAge数值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAge *int64 `json:"MaxAge,omitempty" name:"MaxAge"`

	// 是否包含子域名，on或off。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncludeSubDomains *string `json:"IncludeSubDomains,omitempty" name:"IncludeSubDomains"`

	// 是否预加载，on或off。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Preload *string `json:"Preload,omitempty" name:"Preload"`
}

type Https struct {

	// http2 配置开关
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Http2 *string `json:"Http2,omitempty" name:"Http2"`

	// OCSP 配置开关
	// on：开启
	// off：关闭
	// 默认为关闭状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcspStapling *string `json:"OcspStapling,omitempty" name:"OcspStapling"`

	// Tls版本设置，支持设置 TLSv1, TLSV1.1, TLSV1.2, TLSv1.3，修改时必须开启连续的版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	TlsVersion []*string `json:"TlsVersion,omitempty" name:"TlsVersion"`

	// HSTS 配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hsts *Hsts `json:"Hsts,omitempty" name:"Hsts"`
}

type IdentifyZoneRequest struct {
	*tchttp.BaseRequest

	// 站点名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *IdentifyZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IdentifyZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IdentifyZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type IdentifyZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 站点名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 子域
		Subdomain *string `json:"Subdomain,omitempty" name:"Subdomain"`

		// 记录类型
		RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

		// 记录值
		RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IdentifyZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IdentifyZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportDnsRecordsRequest struct {
	*tchttp.BaseRequest

	// 站点 ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 文件内容
	File *string `json:"File,omitempty" name:"File"`
}

func (r *ImportDnsRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportDnsRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "File")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportDnsRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ImportDnsRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录 ID
		Ids []*string `json:"Ids,omitempty" name:"Ids"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportDnsRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportDnsRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type L7OfflineLog struct {

	// 日志打包开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogTime *int64 `json:"LogTime,omitempty" name:"LogTime"`

	// 站点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 原始大小 单位byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 日志数据包名
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogPacketName *string `json:"LogPacketName,omitempty" name:"LogPacketName"`
}

type LoadBalancing struct {

	// 负载均衡ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名，填写@表示根域
	Host *string `json:"Host,omitempty" name:"Host"`

	// 代理模式：
	// dns_only: 仅DNS
	// proxied: 开启代理
	Type *string `json:"Type,omitempty" name:"Type"`

	// 当Type=dns_only表示DNS的TTL时间
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// 使用的源站组ID
	OriginId []*string `json:"OriginId,omitempty" name:"OriginId"`

	// 使用的源站信息
	Origin []*OriginGroup `json:"Origin,omitempty" name:"Origin"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 调度域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cname *string `json:"Cname,omitempty" name:"Cname"`
}

type MaxAge struct {

	// MaxAge 时间设置，单位秒，最大365天
	// 注意：时间为0，即不缓存。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAgeTime *int64 `json:"MaxAgeTime,omitempty" name:"MaxAgeTime"`

	// 是否遵循源站，on或off，开启时忽略时间设置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowOrigin *string `json:"FollowOrigin,omitempty" name:"FollowOrigin"`
}

type ModifyApplicationProxyRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 四层代理ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 四层代理名称
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// 参数已经废弃
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// 参数已经废弃
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// 会话保持时间，取值范围：30-3600，单位：秒
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// 服务类型
	// hostname：子域名
	// instance：实例
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`
}

func (r *ModifyApplicationProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "ProxyName")
	delete(f, "ForwardClientIp")
	delete(f, "SessionPersist")
	delete(f, "SessionPersistTime")
	delete(f, "ProxyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApplicationProxyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 四层代理ID
		ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApplicationProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 协议，取值为TCP或者UDP
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// 端口，支持格式：
	// 80：80端口
	// 81-90：81至90端口
	Port []*string `json:"Port,omitempty" name:"Port"`

	// 源站类型，取值：
	// custom：手动添加
	// origins：源站组
	// load_balancing：负载均衡
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 源站信息：
	// 当OriginType=custom时，表示多个：
	// IP:端口
	// 域名:端口
	// 当OriginType=origins时，包含一个元素，表示源站组ID
	// 当OriginType=load_balancing时，包含一个元素，表示负载均衡ID
	OriginValue []*string `json:"OriginValue,omitempty" name:"OriginValue"`

	// 传递客户端IP，当Proto=TCP时，取值：
	// TOA：TOA
	// PPV1: Proxy Protocol传递，协议版本V1
	// PPV2: Proxy Protocol传递，协议版本V2
	// OFF：不传递
	// 当Proto=UDP时，取值：
	// PPV2: Proxy Protocol传递，协议版本V2
	// OFF：不传递
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// 是否开启会话保持
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`
}

func (r *ModifyApplicationProxyRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "RuleId")
	delete(f, "Proto")
	delete(f, "Port")
	delete(f, "OriginType")
	delete(f, "OriginValue")
	delete(f, "ForwardClientIp")
	delete(f, "SessionPersist")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApplicationProxyRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则ID
		RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApplicationProxyRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApplicationProxyRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 四层代理ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 状态
	// offline: 停用
	// online: 启用
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyApplicationProxyRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "RuleId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApplicationProxyRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则ID
		RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApplicationProxyRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApplicationProxyStatusRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 四层代理ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 状态
	// offline: 停用
	// online: 启用
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyApplicationProxyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApplicationProxyStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 四层代理ID
		ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApplicationProxyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDefaultCertificateRequest struct {
	*tchttp.BaseRequest

	// Zone ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 默认证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 证书状态
	// deployed: 部署证书
	// disabled:禁用证书
	// 失败状态下重新deployed即可重试失败
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDefaultCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDefaultCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "CertId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDefaultCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDefaultCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDefaultCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDefaultCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDnsRecordRequest struct {
	*tchttp.BaseRequest

	// 记录 ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 站点 ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 记录类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 记录名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 记录内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 生存时间值
	Ttl *int64 `json:"Ttl,omitempty" name:"Ttl"`

	// 优先级
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 代理模式
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

func (r *ModifyDnsRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDnsRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "ZoneId")
	delete(f, "Type")
	delete(f, "Name")
	delete(f, "Content")
	delete(f, "Ttl")
	delete(f, "Priority")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDnsRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDnsRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录 ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// 记录类型
		Type *string `json:"Type,omitempty" name:"Type"`

		// 记录名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 记录内容
		Content *string `json:"Content,omitempty" name:"Content"`

		// 生存时间值
		Ttl *int64 `json:"Ttl,omitempty" name:"Ttl"`

		// 优先级
		Priority *int64 `json:"Priority,omitempty" name:"Priority"`

		// 代理模式
		Mode *string `json:"Mode,omitempty" name:"Mode"`

		// 解析状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// CNAME 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		Cname *string `json:"Cname,omitempty" name:"Cname"`

		// 锁定状态
		Locked *bool `json:"Locked,omitempty" name:"Locked"`

		// 创建时间
		CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

		// 修改时间
		ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

		// 站点 ID
		ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

		// 站点名称
		ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDnsRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDnsRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDnssecRequest struct {
	*tchttp.BaseRequest

	// 站点 ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// DNSSEC 状态
	// - enabled 开启
	// - disabled 关闭
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDnssecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDnssecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDnssecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDnssecResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 站点 ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// 站点名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// DNSSEC 状态
	// - enabled 开启
	// - disabled 关闭
		Status *string `json:"Status,omitempty" name:"Status"`

		// DNSSEC 相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Dnssec *DnssecInfo `json:"Dnssec,omitempty" name:"Dnssec"`

		// 修改时间
		ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDnssecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDnssecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostsCertificateRequest struct {
	*tchttp.BaseRequest

	// Zone ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 本次变更的域名
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`

	// 证书信息, 只需要传入 CertId 即可, 如果为空, 则使用默认证书
	CertInfo []*ServerCertInfo `json:"CertInfo,omitempty" name:"CertInfo"`
}

func (r *ModifyHostsCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostsCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Hosts")
	delete(f, "CertInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHostsCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostsCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostsCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostsCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancingRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 负载均衡ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// 代理模式：
	// dns_only: 仅DNS
	// proxied: 开启代理
	Type *string `json:"Type,omitempty" name:"Type"`

	// 使用的源站组ID
	OriginId []*string `json:"OriginId,omitempty" name:"OriginId"`

	// 当Type=dns_only表示DNS的TTL时间
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`
}

func (r *ModifyLoadBalancingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "LoadBalancingId")
	delete(f, "Type")
	delete(f, "OriginId")
	delete(f, "TTL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 负载均衡ID
		LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLoadBalancingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancingStatusRequest struct {
	*tchttp.BaseRequest

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 负载均衡ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// 状态
	// online: 启用
	// offline: 停用
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyLoadBalancingStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancingStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "LoadBalancingId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancingStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancingStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 负载均衡ID
		LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLoadBalancingStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancingStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneCnameSpeedUpRequest struct {
	*tchttp.BaseRequest

	// 站点 ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// CNAME 加速状态
	// - enabled 开启
	// - disabled 关闭
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyZoneCnameSpeedUpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneCnameSpeedUpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneCnameSpeedUpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneCnameSpeedUpResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 站点 ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// 站点名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// CNAME 加速状态
	// - enabled 开启
	// - disabled 关闭
		Status *string `json:"Status,omitempty" name:"Status"`

		// 更新时间
		ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyZoneCnameSpeedUpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneCnameSpeedUpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneRequest struct {
	*tchttp.BaseRequest

	// 站点 ID，用于唯一标识站点信息
	Id *string `json:"Id,omitempty" name:"Id"`

	// 站点接入方式
	// - full NS 接入
	// - partial CNAME 接入
	Type *string `json:"Type,omitempty" name:"Type"`

	// 自定义站点信息
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitempty" name:"VanityNameServers"`
}

func (r *ModifyZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Type")
	delete(f, "VanityNameServers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 站点 ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// 站点名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 站点当前使用的 NS
		OriginalNameServers []*string `json:"OriginalNameServers,omitempty" name:"OriginalNameServers"`

		// 站点状态
	// - pending 未接入 NS
	// - active 已接入 NS
	// - moved NS 已切走
		Status *string `json:"Status,omitempty" name:"Status"`

		// 站点接入方式
	// - full NS 接入
	// - partial CNAME 接入
		Type *string `json:"Type,omitempty" name:"Type"`

		// 腾讯云分配的 NS 列表
		NameServers []*string `json:"NameServers,omitempty" name:"NameServers"`

		// 创建时间
		CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

		// 修改时间
		ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

		// cname 接入状态
	// - finished 站点验证完成
	// - pending 站点验证中
	// 注意：此字段可能返回 null，表示取不到有效值。
		CnameStatus *string `json:"CnameStatus,omitempty" name:"CnameStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneSettingRequest struct {
	*tchttp.BaseRequest

	// 待变更的站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 缓存过期时间配置
	Cache *CacheConfig `json:"Cache,omitempty" name:"Cache"`

	// 节点缓存键配置
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// 浏览器缓存配置
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// 离线缓存
	OfflineCache *OfflineCache `json:"OfflineCache,omitempty" name:"OfflineCache"`

	// Quic访问
	Quic *Quic `json:"Quic,omitempty" name:"Quic"`

	// POST请求传输配置
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitempty" name:"PostMaxSize"`

	// 智能压缩配置
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// http2回源配置
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitempty" name:"UpstreamHttp2"`

	// 访问协议强制https跳转配置
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// Https 加速配置
	Https *Https `json:"Https,omitempty" name:"Https"`

	// 源站配置
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// 智能加速配置
	SmartRouting *SmartRouting `json:"SmartRouting,omitempty" name:"SmartRouting"`

	// WebSocket配置
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`

	// 客户端IP回源请求头配置
	ClientIpHeader *ClientIp `json:"ClientIpHeader,omitempty" name:"ClientIpHeader"`
}

func (r *ModifyZoneSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Cache")
	delete(f, "CacheKey")
	delete(f, "MaxAge")
	delete(f, "OfflineCache")
	delete(f, "Quic")
	delete(f, "PostMaxSize")
	delete(f, "Compression")
	delete(f, "UpstreamHttp2")
	delete(f, "ForceRedirect")
	delete(f, "Https")
	delete(f, "Origin")
	delete(f, "SmartRouting")
	delete(f, "WebSocket")
	delete(f, "ClientIpHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneSettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 站点ID
		ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyZoneSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneStatusRequest struct {
	*tchttp.BaseRequest

	// 站点 ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 站点状态
	// - false 开启站点
	// - true 关闭站点
	Paused *bool `json:"Paused,omitempty" name:"Paused"`
}

func (r *ModifyZoneStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Paused")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 站点 ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// 站点名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 站点状态
	// - false 开启站点
	// - true 关闭站点
		Paused *bool `json:"Paused,omitempty" name:"Paused"`

		// 更新时间
		ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyZoneStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OfflineCache struct {

	// on | off, 离线缓存是否开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Origin struct {

	// 回源协议配置
	// http：强制 http 回源
	// follow：协议跟随回源
	// https：强制 https 回源，https 回源时仅支持源站 443 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginPullProtocol *string `json:"OriginPullProtocol,omitempty" name:"OriginPullProtocol"`
}

type OriginGroup struct {

	// 源站组ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// 源站组名称
	OriginName *string `json:"OriginName,omitempty" name:"OriginName"`

	// 配置类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 记录
	Record []*OriginRecord `json:"Record,omitempty" name:"Record"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 站点ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 源站类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 是否为四层代理使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationProxyUsed *bool `json:"ApplicationProxyUsed,omitempty" name:"ApplicationProxyUsed"`

	// 是否为负载均衡使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancingUsed *bool `json:"LoadBalancingUsed,omitempty" name:"LoadBalancingUsed"`
}

type OriginRecord struct {

	// 记录值
	Record *string `json:"Record,omitempty" name:"Record"`

	// 当源站配置类型Type=area时，表示区域
	// 当源站类型Type=area时，为空表示默认区域
	Area []*string `json:"Area,omitempty" name:"Area"`

	// 当源站配置类型Type=weight时，表示权重
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// 端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 记录ID
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// 是否私有鉴权
	// 当源站类型OriginType=third_part时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Private *bool `json:"Private,omitempty" name:"Private"`

	// 私有鉴权参数
	// 当源站类型Private=true时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateParameter []*OriginRecordPrivateParameter `json:"PrivateParameter,omitempty" name:"PrivateParameter"`
}

type OriginRecordPrivateParameter struct {

	// 私有鉴权参数名称：
	// "AccessKeyId"：Access Key ID
	// "SecretAccessKey"：Secret Access Key
	Name *string `json:"Name,omitempty" name:"Name"`

	// 私有鉴权参数数值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type PostMaxSize struct {

	// 是调整POST请求限制，平台默认为32MB。
	// 关闭：off，
	// 开启：on。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 最大限制，取值在1MB和500MB之间。单位字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

type QueryString struct {

	// on | off CacheKey是否由QueryString组成
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// includeCustom:使用部分url参数
	// excludeCustom:排除部分url参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 使用/排除的url参数数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type Quic struct {

	// 是否启动Quic配置
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ReclaimZoneRequest struct {
	*tchttp.BaseRequest

	// 站点名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ReclaimZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReclaimZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReclaimZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ReclaimZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 站点名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReclaimZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReclaimZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanDnsRecordsRequest struct {
	*tchttp.BaseRequest

	// 站点 ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *ScanDnsRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanDnsRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanDnsRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ScanDnsRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 扫描状态
	// - doing 扫描中
	// - done 扫描完成
		Status *string `json:"Status,omitempty" name:"Status"`

		// 扫描后添加的记录数
		RecordsAdded *int64 `json:"RecordsAdded,omitempty" name:"RecordsAdded"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScanDnsRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanDnsRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerCertInfo struct {

	// 服务器证书 ID, 默认证书ID, 或在 SSL 证书管理进行证书托管时自动生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 证书备注名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 证书类型:
	// default: 默认证书
	// upload:用户上传
	// managed:腾讯云托管
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 证书过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 证书部署时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployTime *string `json:"DeployTime,omitempty" name:"DeployTime"`

	// 部署状态:
	// processing: 部署中
	// deployed: 已部署
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type SmartRouting struct {

	// 智能加速配置开关
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Task struct {

	// 任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 资源
	Target *string `json:"Target,omitempty" name:"Target"`

	// 任务类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务完成时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type UpstreamHttp2 struct {

	// http2回源配置开关
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type VanityNameServers struct {

	// 自定义 ns 开关
	// - on 开启
	// - off 关闭
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 自定义 ns 列表
	Servers []*string `json:"Servers,omitempty" name:"Servers"`
}

type VanityNameServersIps struct {

	// 自定义名字服务器名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 自定义名字服务器 IPv4 地址
	IPv4 *string `json:"IPv4,omitempty" name:"IPv4"`
}

type WebSocket struct {

	// WebSocket 超时配置开关, 开关为off时，平台仍支持WebSocket连接，此时超时时间默认为15秒，若需要调整超时时间，将开关置为on.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 设置超时时间，单位为秒，最大超时时间120秒。
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
}

type Zone struct {

	// 站点ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 站点名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 站点当前使用的 NS 列表
	OriginalNameServers []*string `json:"OriginalNameServers,omitempty" name:"OriginalNameServers"`

	// 腾讯云分配的 NS 列表
	NameServers []*string `json:"NameServers,omitempty" name:"NameServers"`

	// 站点状态
	// - active：NS 已切换
	// - pending：NS 未切换
	// - moved：NS 已切走
	// - deactivated：被封禁
	Status *string `json:"Status,omitempty" name:"Status"`

	// 站点接入方式
	// - full：NS 接入
	// - partial：CNAME 接入
	Type *string `json:"Type,omitempty" name:"Type"`

	// 站点是否关闭
	Paused *bool `json:"Paused,omitempty" name:"Paused"`

	// 站点创建时间
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 站点修改时间
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// cname 接入状态
	// - finished 站点已验证
	// - pending 站点验证中
	// 注意：此字段可能返回 null，表示取不到有效值。
	CnameStatus *string `json:"CnameStatus,omitempty" name:"CnameStatus"`
}

type ZoneFilter struct {

	// 过滤字段名，支持的列表如下：
	// - name: 站点名。
	// - status: 站点状态
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤字段值
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否启用模糊查询，仅支持过滤字段名为name。模糊查询时，Values长度最大为1
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}
