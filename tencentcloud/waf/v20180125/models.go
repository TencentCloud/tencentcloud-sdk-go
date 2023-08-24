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

package v20180125

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccessFullTextInfo struct {
	// 是否大小写敏感
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`

	// 全文索引的分词符，字符串中每个字符代表一个分词符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tokenizer *string `json:"Tokenizer,omitempty" name:"Tokenizer"`

	// 是否包含中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainZH *bool `json:"ContainZH,omitempty" name:"ContainZH"`
}

type AccessHistogramItem struct {
	// 时间，单位ms
	BTime *int64 `json:"BTime,omitempty" name:"BTime"`

	// 日志条数
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type AccessKeyValueInfo struct {
	// 需要配置键值或者元字段索引的字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 字段的索引描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *AccessValueInfo `json:"Value,omitempty" name:"Value"`
}

type AccessLogInfo struct {
	// 日志时间，单位ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// 日志主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志主题名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 日志来源IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitempty" name:"Source"`

	// 日志文件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 日志上报请求包的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 请求包内日志的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgLogId *string `json:"PkgLogId,omitempty" name:"PkgLogId"`

	// 日志内容的Json序列化字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogJson *string `json:"LogJson,omitempty" name:"LogJson"`
}

type AccessLogItem struct {
	// 日记Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 日志Value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type AccessLogItems struct {
	// 分析结果返回的KV数据对
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*AccessLogItem `json:"Data,omitempty" name:"Data"`
}

type AccessRuleInfo struct {
	// 全文索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullText *AccessFullTextInfo `json:"FullText,omitempty" name:"FullText"`

	// 键值索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyValue *AccessRuleKeyValueInfo `json:"KeyValue,omitempty" name:"KeyValue"`

	// 元字段索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *AccessRuleTagInfo `json:"Tag,omitempty" name:"Tag"`
}

type AccessRuleKeyValueInfo struct {
	// 是否大小写敏感
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`

	// 需要建立索引的键值对信息；最大只能配置100个键值对
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyValues []*AccessKeyValueInfo `json:"KeyValues,omitempty" name:"KeyValues"`
}

type AccessRuleTagInfo struct {
	// 是否大小写敏感
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`

	// 标签索引配置中的字段信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyValues []*AccessKeyValueInfo `json:"KeyValues,omitempty" name:"KeyValues"`
}

type AccessValueInfo struct {
	// 字段类型，目前支持的类型有：long、text、double
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 字段的分词符，只有当字段类型为text时才有意义；输入字符串中的每个字符代表一个分词符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tokenizer *string `json:"Tokenizer,omitempty" name:"Tokenizer"`

	// 字段是否开启分析功能
	// 注意：此字段可能返回 null，表示取不到有效值。
	SqlFlag *bool `json:"SqlFlag,omitempty" name:"SqlFlag"`

	// 是否包含中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainZH *bool `json:"ContainZH,omitempty" name:"ContainZH"`
}

// Predefined struct for user
type AddCustomRuleRequestParams struct {
	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 优先级
	SortId *string `json:"SortId,omitempty" name:"SortId"`

	// 过期时间，单位为秒级时间戳，例如1677254399表示过期时间为2023-02-24 23:59:59. 0表示永不过期
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies"`

	// 需要添加策略的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 动作类型，1代表阻断，2代表人机识别，3代表观察，4代表重定向
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 如果动作是重定向，则表示重定向的地址；其他情况可以为空
	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 放行的详情
	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`

	// 添加规则的来源，默认为空
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

type AddCustomRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 优先级
	SortId *string `json:"SortId,omitempty" name:"SortId"`

	// 过期时间，单位为秒级时间戳，例如1677254399表示过期时间为2023-02-24 23:59:59. 0表示永不过期
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies"`

	// 需要添加策略的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 动作类型，1代表阻断，2代表人机识别，3代表观察，4代表重定向
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 如果动作是重定向，则表示重定向的地址；其他情况可以为空
	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 放行的详情
	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`

	// 添加规则的来源，默认为空
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *AddCustomRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCustomRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SortId")
	delete(f, "ExpireTime")
	delete(f, "Strategies")
	delete(f, "Domain")
	delete(f, "ActionType")
	delete(f, "Redirect")
	delete(f, "Edition")
	delete(f, "Bypass")
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCustomRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCustomRuleResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitempty" name:"Success"`

	// 添加成功的规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddCustomRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddCustomRuleResponseParams `json:"Response"`
}

func (r *AddCustomRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCustomRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCustomWhiteRuleRequestParams struct {
	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 优先级
	SortId *string `json:"SortId,omitempty" name:"SortId"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies"`

	// 需要添加策略的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 放行的详情
	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`
}

type AddCustomWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 优先级
	SortId *string `json:"SortId,omitempty" name:"SortId"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies"`

	// 需要添加策略的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 放行的详情
	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`
}

func (r *AddCustomWhiteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCustomWhiteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SortId")
	delete(f, "ExpireTime")
	delete(f, "Strategies")
	delete(f, "Domain")
	delete(f, "Bypass")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCustomWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCustomWhiteRuleResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitempty" name:"Success"`

	// 添加成功的规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddCustomWhiteRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddCustomWhiteRuleResponseParams `json:"Response"`
}

func (r *AddCustomWhiteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCustomWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddDomainWhiteRuleRequestParams struct {
	// 需要添加的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 需要添加的规则
	Rules []*uint64 `json:"Rules,omitempty" name:"Rules"`

	// 需要添加的规则url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 规则的方法
	Function *string `json:"Function,omitempty" name:"Function"`

	// 规则的开关，0表示规则关闭，1表示规则打开
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type AddDomainWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要添加的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 需要添加的规则
	Rules []*uint64 `json:"Rules,omitempty" name:"Rules"`

	// 需要添加的规则url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 规则的方法
	Function *string `json:"Function,omitempty" name:"Function"`

	// 规则的开关，0表示规则关闭，1表示规则打开
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *AddDomainWhiteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDomainWhiteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Rules")
	delete(f, "Url")
	delete(f, "Function")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddDomainWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddDomainWhiteRuleResponseParams struct {
	// 规则id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddDomainWhiteRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddDomainWhiteRuleResponseParams `json:"Response"`
}

func (r *AddDomainWhiteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDomainWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddSpartaProtectionAutoRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type AddSpartaProtectionAutoRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *AddSpartaProtectionAutoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSpartaProtectionAutoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddSpartaProtectionAutoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddSpartaProtectionAutoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddSpartaProtectionAutoResponse struct {
	*tchttp.BaseResponse
	Response *AddSpartaProtectionAutoResponseParams `json:"Response"`
}

func (r *AddSpartaProtectionAutoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSpartaProtectionAutoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddSpartaProtectionRequestParams struct {
	// 需要防御的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 证书类型，0表示没有证书，CertType=1表示自有证书,2 为托管证书
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// 表示是否开启了CDN代理，1：有部署CDN，0：未部署CDN
	IsCdn *int64 `json:"IsCdn,omitempty" name:"IsCdn"`

	// 回源类型，0表示通过IP回源,1 表示通过域名回源
	UpstreamType *int64 `json:"UpstreamType,omitempty" name:"UpstreamType"`

	// 是否开启WebSocket支持，1表示开启，0不开启
	IsWebsocket *int64 `json:"IsWebsocket,omitempty" name:"IsWebsocket"`

	// 负载均衡策略，0表示轮询，1表示IP hash
	LoadBalance *string `json:"LoadBalance,omitempty" name:"LoadBalance"`

	// 值为1时，需要填次参数，表示证书内容
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// CertType=1时，需要填次参数，表示证书的私钥
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// CertType=2时，需要填次参数，表示证书的ID
	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`

	// Waf的资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// HTTPS回源协议，填http或者https
	UpstreamScheme *string `json:"UpstreamScheme,omitempty" name:"UpstreamScheme"`

	// HTTPS回源端口,仅UpstreamScheme为http时需要填当前字段
	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitempty" name:"HttpsUpstreamPort"`

	// 是否开启灰度，0表示不开启灰度
	IsGray *int64 `json:"IsGray,omitempty" name:"IsGray"`

	// 灰度的地区
	GrayAreas []*string `json:"GrayAreas,omitempty" name:"GrayAreas"`

	// UpstreamType=1时，填次字段表示回源域名
	UpstreamDomain *string `json:"UpstreamDomain,omitempty" name:"UpstreamDomain"`

	// UpstreamType=0时，填次字段表示回源IP
	SrcList []*string `json:"SrcList,omitempty" name:"SrcList"`

	// 是否开启HTTP2,开启HTTP2需要HTTPS支持
	IsHttp2 *int64 `json:"IsHttp2,omitempty" name:"IsHttp2"`

	// 表示是否强制跳转到HTTPS，1强制跳转Https，0不强制跳转
	HttpsRewrite *int64 `json:"HttpsRewrite,omitempty" name:"HttpsRewrite"`

	// 服务有多端口需要设置此字段
	Ports []*PortItem `json:"Ports,omitempty" name:"Ports"`

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF，cdn-waf表示CDN上的Web防护能力
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 是否开启长连接，仅IP回源时可以用填次参数，域名回源时这个参数无效
	IsKeepAlive *string `json:"IsKeepAlive,omitempty" name:"IsKeepAlive"`

	// 实例id，上线之后带上此字段
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// anycast IP类型开关： 0 普通IP 1 Anycast IP
	Anycast *int64 `json:"Anycast,omitempty" name:"Anycast"`

	// src权重
	Weights []*int64 `json:"Weights,omitempty" name:"Weights"`

	// 是否开启主动健康检测，1表示开启，0表示不开启
	ActiveCheck *int64 `json:"ActiveCheck,omitempty" name:"ActiveCheck"`

	// TLS版本信息
	TLSVersion *int64 `json:"TLSVersion,omitempty" name:"TLSVersion"`

	// 加密套件信息
	Ciphers []*int64 `json:"Ciphers,omitempty" name:"Ciphers"`

	// 0:不支持选择：默认模版  1:通用型模版 2:安全型模版 3:自定义模版
	CipherTemplate *int64 `json:"CipherTemplate,omitempty" name:"CipherTemplate"`

	// 300s
	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitempty" name:"ProxyReadTimeout"`

	// 300s
	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitempty" name:"ProxySendTimeout"`

	// 0:关闭SNI；1:开启SNI，SNI=源请求host；2:开启SNI，SNI=修改为源站host；3：开启SNI，自定义host，SNI=SniHost；
	SniType *int64 `json:"SniType,omitempty" name:"SniType"`

	// SniType=3时，需要填此参数，表示自定义的host；
	SniHost *string `json:"SniHost,omitempty" name:"SniHost"`

	// is_cdn=3时，需要填此参数，表示自定义header
	IpHeaders []*string `json:"IpHeaders,omitempty" name:"IpHeaders"`

	// 0:关闭xff重置；1:开启xff重置
	XFFReset *int64 `json:"XFFReset,omitempty" name:"XFFReset"`
}

type AddSpartaProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 需要防御的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 证书类型，0表示没有证书，CertType=1表示自有证书,2 为托管证书
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// 表示是否开启了CDN代理，1：有部署CDN，0：未部署CDN
	IsCdn *int64 `json:"IsCdn,omitempty" name:"IsCdn"`

	// 回源类型，0表示通过IP回源,1 表示通过域名回源
	UpstreamType *int64 `json:"UpstreamType,omitempty" name:"UpstreamType"`

	// 是否开启WebSocket支持，1表示开启，0不开启
	IsWebsocket *int64 `json:"IsWebsocket,omitempty" name:"IsWebsocket"`

	// 负载均衡策略，0表示轮询，1表示IP hash
	LoadBalance *string `json:"LoadBalance,omitempty" name:"LoadBalance"`

	// 值为1时，需要填次参数，表示证书内容
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// CertType=1时，需要填次参数，表示证书的私钥
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// CertType=2时，需要填次参数，表示证书的ID
	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`

	// Waf的资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// HTTPS回源协议，填http或者https
	UpstreamScheme *string `json:"UpstreamScheme,omitempty" name:"UpstreamScheme"`

	// HTTPS回源端口,仅UpstreamScheme为http时需要填当前字段
	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitempty" name:"HttpsUpstreamPort"`

	// 是否开启灰度，0表示不开启灰度
	IsGray *int64 `json:"IsGray,omitempty" name:"IsGray"`

	// 灰度的地区
	GrayAreas []*string `json:"GrayAreas,omitempty" name:"GrayAreas"`

	// UpstreamType=1时，填次字段表示回源域名
	UpstreamDomain *string `json:"UpstreamDomain,omitempty" name:"UpstreamDomain"`

	// UpstreamType=0时，填次字段表示回源IP
	SrcList []*string `json:"SrcList,omitempty" name:"SrcList"`

	// 是否开启HTTP2,开启HTTP2需要HTTPS支持
	IsHttp2 *int64 `json:"IsHttp2,omitempty" name:"IsHttp2"`

	// 表示是否强制跳转到HTTPS，1强制跳转Https，0不强制跳转
	HttpsRewrite *int64 `json:"HttpsRewrite,omitempty" name:"HttpsRewrite"`

	// 服务有多端口需要设置此字段
	Ports []*PortItem `json:"Ports,omitempty" name:"Ports"`

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF，cdn-waf表示CDN上的Web防护能力
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 是否开启长连接，仅IP回源时可以用填次参数，域名回源时这个参数无效
	IsKeepAlive *string `json:"IsKeepAlive,omitempty" name:"IsKeepAlive"`

	// 实例id，上线之后带上此字段
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// anycast IP类型开关： 0 普通IP 1 Anycast IP
	Anycast *int64 `json:"Anycast,omitempty" name:"Anycast"`

	// src权重
	Weights []*int64 `json:"Weights,omitempty" name:"Weights"`

	// 是否开启主动健康检测，1表示开启，0表示不开启
	ActiveCheck *int64 `json:"ActiveCheck,omitempty" name:"ActiveCheck"`

	// TLS版本信息
	TLSVersion *int64 `json:"TLSVersion,omitempty" name:"TLSVersion"`

	// 加密套件信息
	Ciphers []*int64 `json:"Ciphers,omitempty" name:"Ciphers"`

	// 0:不支持选择：默认模版  1:通用型模版 2:安全型模版 3:自定义模版
	CipherTemplate *int64 `json:"CipherTemplate,omitempty" name:"CipherTemplate"`

	// 300s
	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitempty" name:"ProxyReadTimeout"`

	// 300s
	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitempty" name:"ProxySendTimeout"`

	// 0:关闭SNI；1:开启SNI，SNI=源请求host；2:开启SNI，SNI=修改为源站host；3：开启SNI，自定义host，SNI=SniHost；
	SniType *int64 `json:"SniType,omitempty" name:"SniType"`

	// SniType=3时，需要填此参数，表示自定义的host；
	SniHost *string `json:"SniHost,omitempty" name:"SniHost"`

	// is_cdn=3时，需要填此参数，表示自定义header
	IpHeaders []*string `json:"IpHeaders,omitempty" name:"IpHeaders"`

	// 0:关闭xff重置；1:开启xff重置
	XFFReset *int64 `json:"XFFReset,omitempty" name:"XFFReset"`
}

func (r *AddSpartaProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSpartaProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "CertType")
	delete(f, "IsCdn")
	delete(f, "UpstreamType")
	delete(f, "IsWebsocket")
	delete(f, "LoadBalance")
	delete(f, "Cert")
	delete(f, "PrivateKey")
	delete(f, "SSLId")
	delete(f, "ResourceId")
	delete(f, "UpstreamScheme")
	delete(f, "HttpsUpstreamPort")
	delete(f, "IsGray")
	delete(f, "GrayAreas")
	delete(f, "UpstreamDomain")
	delete(f, "SrcList")
	delete(f, "IsHttp2")
	delete(f, "HttpsRewrite")
	delete(f, "Ports")
	delete(f, "Edition")
	delete(f, "IsKeepAlive")
	delete(f, "InstanceID")
	delete(f, "Anycast")
	delete(f, "Weights")
	delete(f, "ActiveCheck")
	delete(f, "TLSVersion")
	delete(f, "Ciphers")
	delete(f, "CipherTemplate")
	delete(f, "ProxyReadTimeout")
	delete(f, "ProxySendTimeout")
	delete(f, "SniType")
	delete(f, "SniHost")
	delete(f, "IpHeaders")
	delete(f, "XFFReset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddSpartaProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddSpartaProtectionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddSpartaProtectionResponse struct {
	*tchttp.BaseResponse
	Response *AddSpartaProtectionResponseParams `json:"Response"`
}

func (r *AddSpartaProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSpartaProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddSpartaProtectionsAutoRequestParams struct {
	// 多域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type AddSpartaProtectionsAutoRequest struct {
	*tchttp.BaseRequest
	
	// 多域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *AddSpartaProtectionsAutoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSpartaProtectionsAutoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddSpartaProtectionsAutoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddSpartaProtectionsAutoResponseParams struct {
	// 失败原因
	FailedInfos []*FailedInfo `json:"FailedInfos,omitempty" name:"FailedInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddSpartaProtectionsAutoResponse struct {
	*tchttp.BaseResponse
	Response *AddSpartaProtectionsAutoResponseParams `json:"Response"`
}

func (r *AddSpartaProtectionsAutoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSpartaProtectionsAutoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttackLogInfo struct {
	// 攻击日志的详情内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// CLS返回内容
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// CLS返回内容
	Source *string `json:"Source,omitempty" name:"Source"`

	// CLS返回内容
	TimeStamp *string `json:"TimeStamp,omitempty" name:"TimeStamp"`
}

type AutoDenyDetail struct {
	// 攻击封禁类型标签
	AttackTags []*string `json:"AttackTags,omitempty" name:"AttackTags"`

	// 攻击次数阈值
	AttackThreshold *int64 `json:"AttackThreshold,omitempty" name:"AttackThreshold"`

	// 自动封禁状态
	DefenseStatus *int64 `json:"DefenseStatus,omitempty" name:"DefenseStatus"`

	// 攻击时间阈值
	TimeThreshold *int64 `json:"TimeThreshold,omitempty" name:"TimeThreshold"`

	// 自动封禁时间
	DenyTimeThreshold *int64 `json:"DenyTimeThreshold,omitempty" name:"DenyTimeThreshold"`

	// 最后更新时间
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

type BotPkg struct {
	// 资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *int64 `json:"Region,omitempty" name:"Region"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 申请数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InquireNum *int64 `json:"InquireNum,omitempty" name:"InquireNum"`

	// 使用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedNum *int64 `json:"UsedNum,omitempty" name:"UsedNum"`

	// 子产品code
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 续费标志	
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type BotQPS struct {
	// 资源id
	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 有效时间
	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`

	// 资源数量
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 资源所在地区
	Region *string `json:"Region,omitempty" name:"Region"`

	// 使用qps的最大值
	MaxBotQPS *uint64 `json:"MaxBotQPS,omitempty" name:"MaxBotQPS"`

	// 续费标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type BotStatPointItem struct {
	// 横坐标
	TimeStamp *string `json:"TimeStamp,omitempty" name:"TimeStamp"`

	// value的所属对象
	Key *string `json:"Key,omitempty" name:"Key"`

	// 纵列表
	Value *int64 `json:"Value,omitempty" name:"Value"`

	// Key对应的页面展示内容
	Label *string `json:"Label,omitempty" name:"Label"`
}

type CdcCluster struct {
	// cdc的集群id
	Id *string `json:"Id,omitempty" name:"Id"`

	// cdc的集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type CdcRegion struct {
	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 该地域对应的集群信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Clusters []*CdcCluster `json:"Clusters,omitempty" name:"Clusters"`
}

type ClbDomainsInfo struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名id
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// waf类型
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 是否是cdn
	IsCdn *uint64 `json:"IsCdn,omitempty" name:"IsCdn"`

	// 负载均衡算法
	LoadBalancerSet []*LoadBalancerPackageNew `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`

	// 镜像模式
	FlowMode *uint64 `json:"FlowMode,omitempty" name:"FlowMode"`

	// 绑定clb状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *int64 `json:"State,omitempty" name:"State"`

	// 负载均衡类型，clb或者apisix
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlbType *string `json:"AlbType,omitempty" name:"AlbType"`

	// IsCdn=3时，表示自定义header
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpHeaders []*string `json:"IpHeaders,omitempty" name:"IpHeaders"`

	// cdc类型会增加集群信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdcClusters *string `json:"CdcClusters,omitempty" name:"CdcClusters"`
}

type ClbHostResult struct {
	// WAF绑定的监听器实例
	LoadBalancer *LoadBalancer `json:"LoadBalancer,omitempty" name:"LoadBalancer"`

	// WAF绑定的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// WAF绑定的实例ID
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 是否有绑定WAF，1：绑定了WAF，0：没有绑定WAF
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 绑定了WAF的情况下，WAF流量模式，1：清洗模式，0：镜像模式（默认）
	FlowMode *uint64 `json:"FlowMode,omitempty" name:"FlowMode"`
}

type ClbHostsParams struct {
	// 负载均衡实例ID，如果不传次参数则默认认为操作的是整个AppId的监听器，如果此参数不为空则认为操作的是对应负载均衡实例。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID，，如果不传次参数则默认认为操作的是整个负载均衡实例，如果此参数不为空则认为操作的是对应负载均衡监听器。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// WAF实例ID，，如果不传次参数则默认认为操作的是整个负载均衡监听器实例，如果此参数不为空则认为操作的是对应负载均衡监听器的某一个具体的域名。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
}

// Predefined struct for user
type CreateAccessExportRequestParams struct {
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitempty" name:"To"`

	// 日志导出检索语句
	Query *string `json:"Query,omitempty" name:"Query"`

	// 日志导出数量，最大值100w
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 日志导出数据格式。json，csv，默认为json
	Format *string `json:"Format,omitempty" name:"Format"`

	// 日志导出时间排序。desc，asc，默认为desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type CreateAccessExportRequest struct {
	*tchttp.BaseRequest
	
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitempty" name:"To"`

	// 日志导出检索语句
	Query *string `json:"Query,omitempty" name:"Query"`

	// 日志导出数量，最大值100w
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 日志导出数据格式。json，csv，默认为json
	Format *string `json:"Format,omitempty" name:"Format"`

	// 日志导出时间排序。desc，asc，默认为desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *CreateAccessExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "Count")
	delete(f, "Format")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccessExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessExportResponseParams struct {
	// 日志导出ID。
	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAccessExportResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccessExportResponseParams `json:"Response"`
}

func (r *CreateAccessExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHostRequestParams struct {
	// 防护域名配置信息
	Host *HostRecord `json:"Host,omitempty" name:"Host"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type CreateHostRequest struct {
	*tchttp.BaseRequest
	
	// 防护域名配置信息
	Host *HostRecord `json:"Host,omitempty" name:"Host"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *CreateHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Host")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHostResponseParams struct {
	// 新增防护域名ID
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateHostResponse struct {
	*tchttp.BaseResponse
	Response *CreateHostResponseParams `json:"Response"`
}

func (r *CreateHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccessExportRequestParams struct {
	// 日志导出ID
	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`

	// 日志主题
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

type DeleteAccessExportRequest struct {
	*tchttp.BaseRequest
	
	// 日志导出ID
	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`

	// 日志主题
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteAccessExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccessExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportId")
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccessExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccessExportResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAccessExportResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccessExportResponseParams `json:"Response"`
}

func (r *DeleteAccessExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccessExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAttackDownloadRecordRequestParams struct {
	// 下载任务记录唯一标记
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DeleteAttackDownloadRecordRequest struct {
	*tchttp.BaseRequest
	
	// 下载任务记录唯一标记
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteAttackDownloadRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttackDownloadRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAttackDownloadRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAttackDownloadRecordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAttackDownloadRecordResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAttackDownloadRecordResponseParams `json:"Response"`
}

func (r *DeleteAttackDownloadRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttackDownloadRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomWhiteRuleRequestParams struct {
	// 删除的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 删除的规则ID
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
}

type DeleteCustomWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 删除的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 删除的规则ID
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteCustomWhiteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomWhiteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomWhiteRuleResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCustomWhiteRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomWhiteRuleResponseParams `json:"Response"`
}

func (r *DeleteCustomWhiteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainWhiteRulesRequestParams struct {
	// 需要删除的规则域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 需要删除的白名单规则
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type DeleteDomainWhiteRulesRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的规则域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 需要删除的白名单规则
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteDomainWhiteRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainWhiteRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDomainWhiteRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainWhiteRulesResponseParams struct {
	// 出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDomainWhiteRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDomainWhiteRulesResponseParams `json:"Response"`
}

func (r *DeleteDomainWhiteRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainWhiteRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDownloadRecordRequestParams struct {
	// 记录id
	Flow *string `json:"Flow,omitempty" name:"Flow"`
}

type DeleteDownloadRecordRequest struct {
	*tchttp.BaseRequest
	
	// 记录id
	Flow *string `json:"Flow,omitempty" name:"Flow"`
}

func (r *DeleteDownloadRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDownloadRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Flow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDownloadRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDownloadRecordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDownloadRecordResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDownloadRecordResponseParams `json:"Response"`
}

func (r *DeleteDownloadRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDownloadRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHostRequestParams struct {
	// 删除的域名列表
	HostsDel []*HostDel `json:"HostsDel,omitempty" name:"HostsDel"`
}

type DeleteHostRequest struct {
	*tchttp.BaseRequest
	
	// 删除的域名列表
	HostsDel []*HostDel `json:"HostsDel,omitempty" name:"HostsDel"`
}

func (r *DeleteHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HostsDel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHostResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteHostResponse struct {
	*tchttp.BaseResponse
	Response *DeleteHostResponseParams `json:"Response"`
}

func (r *DeleteHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIpAccessControlRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 删除的ip数组
	Items []*string `json:"Items,omitempty" name:"Items"`

	// 是否删除对应的域名下的所有黑/白IP名单，true表示全部删除，false表示只删除指定ip名单
	DeleteAll *bool `json:"DeleteAll,omitempty" name:"DeleteAll"`

	// 是否为多域名黑白名单
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
}

type DeleteIpAccessControlRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 删除的ip数组
	Items []*string `json:"Items,omitempty" name:"Items"`

	// 是否删除对应的域名下的所有黑/白IP名单，true表示全部删除，false表示只删除指定ip名单
	DeleteAll *bool `json:"DeleteAll,omitempty" name:"DeleteAll"`

	// 是否为多域名黑白名单
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
}

func (r *DeleteIpAccessControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIpAccessControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Items")
	delete(f, "DeleteAll")
	delete(f, "SourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIpAccessControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIpAccessControlResponseParams struct {
	// 删除失败的条目
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedItems *string `json:"FailedItems,omitempty" name:"FailedItems"`

	// 删除失败的条目数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteIpAccessControlResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIpAccessControlResponseParams `json:"Response"`
}

func (r *DeleteIpAccessControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIpAccessControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSessionRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// clb-waf 或者 sprta-waf
	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

type DeleteSessionRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// clb-waf 或者 sprta-waf
	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *DeleteSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Edition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSessionResponseParams struct {
	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSessionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSessionResponseParams `json:"Response"`
}

func (r *DeleteSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSpartaProtectionRequestParams struct {
	// 域名列表
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 版本
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type DeleteSpartaProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 域名列表
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 版本
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DeleteSpartaProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSpartaProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "Edition")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSpartaProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSpartaProtectionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSpartaProtectionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSpartaProtectionResponseParams `json:"Response"`
}

func (r *DeleteSpartaProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSpartaProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessExportsRequestParams struct {
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeAccessExportsRequest struct {
	*tchttp.BaseRequest
	
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAccessExportsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessExportsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessExportsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessExportsResponseParams struct {
	// 日志导出ID。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 日志导出列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Exports []*ExportAccessInfo `json:"Exports,omitempty" name:"Exports"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccessExportsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessExportsResponseParams `json:"Response"`
}

func (r *DescribeAccessExportsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessExportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessFastAnalysisRequestParams struct {
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitempty" name:"To"`

	// 查询语句，语句长度最大为4096，由于本接口是分析接口，如果无过滤条件，必须传 * 表示匹配所有，参考CLS的分析统计语句的文档
	Query *string `json:"Query,omitempty" name:"Query"`

	// 需要分析统计的字段名
	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`
}

type DescribeAccessFastAnalysisRequest struct {
	*tchttp.BaseRequest
	
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitempty" name:"To"`

	// 查询语句，语句长度最大为4096，由于本接口是分析接口，如果无过滤条件，必须传 * 表示匹配所有，参考CLS的分析统计语句的文档
	Query *string `json:"Query,omitempty" name:"Query"`

	// 需要分析统计的字段名
	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`
}

func (r *DescribeAccessFastAnalysisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessFastAnalysisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "FieldName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessFastAnalysisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessFastAnalysisResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccessFastAnalysisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessFastAnalysisResponseParams `json:"Response"`
}

func (r *DescribeAccessFastAnalysisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessFastAnalysisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessHistogramRequestParams struct {
	// 老版本查询的日志主题ID，新版本传空字符串即可
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitempty" name:"To"`

	// 查询语句，语句长度最大为4096
	Query *string `json:"Query,omitempty" name:"Query"`

	// 柱状图间隔时间差，单位ms
	Interval *int64 `json:"Interval,omitempty" name:"Interval"`
}

type DescribeAccessHistogramRequest struct {
	*tchttp.BaseRequest
	
	// 老版本查询的日志主题ID，新版本传空字符串即可
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitempty" name:"To"`

	// 查询语句，语句长度最大为4096
	Query *string `json:"Query,omitempty" name:"Query"`

	// 柱状图间隔时间差，单位ms
	Interval *int64 `json:"Interval,omitempty" name:"Interval"`
}

func (r *DescribeAccessHistogramRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessHistogramRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "Interval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessHistogramRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessHistogramResponseParams struct {
	// 柱状图间隔时间差，单位ms
	Interval *int64 `json:"Interval,omitempty" name:"Interval"`

	// 满足条件的日志条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 注意：此字段可能返回 null，表示取不到有效值
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistogramInfos []*AccessHistogramItem `json:"HistogramInfos,omitempty" name:"HistogramInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccessHistogramResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessHistogramResponseParams `json:"Response"`
}

func (r *DescribeAccessHistogramResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessHistogramResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessIndexRequestParams struct {

}

type DescribeAccessIndexRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAccessIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessIndexResponseParams struct {
	// 是否生效，true表示生效，false表示未生效
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 索引配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rule *AccessRuleInfo `json:"Rule,omitempty" name:"Rule"`

	// 索引修改时间，初始值为索引创建时间。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccessIndexResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessIndexResponseParams `json:"Response"`
}

func (r *DescribeAccessIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackOverviewRequestParams struct {
	// 查询开始时间
	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`

	// 客户的Appid
	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`

	// 被查询的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type DescribeAttackOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间
	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`

	// 客户的Appid
	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`

	// 被查询的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribeAttackOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromTime")
	delete(f, "ToTime")
	delete(f, "Appid")
	delete(f, "Domain")
	delete(f, "Edition")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAttackOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackOverviewResponseParams struct {
	// 访问请求总数
	AccessCount *uint64 `json:"AccessCount,omitempty" name:"AccessCount"`

	// Web攻击总数
	AttackCount *uint64 `json:"AttackCount,omitempty" name:"AttackCount"`

	// 访问控制总数
	ACLCount *uint64 `json:"ACLCount,omitempty" name:"ACLCount"`

	// CC攻击总数
	CCCount *uint64 `json:"CCCount,omitempty" name:"CCCount"`

	// Bot攻击总数
	BotCount *uint64 `json:"BotCount,omitempty" name:"BotCount"`

	// api资产总数
	ApiAssetsCount *uint64 `json:"ApiAssetsCount,omitempty" name:"ApiAssetsCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAttackOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAttackOverviewResponseParams `json:"Response"`
}

func (r *DescribeAttackOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoDenyIPRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 查询IP自动封禁状态
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 计数标识
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 类别
	Category *string `json:"Category,omitempty" name:"Category"`

	// 有效时间最小时间戳
	VtsMin *uint64 `json:"VtsMin,omitempty" name:"VtsMin"`

	// 有效时间最大时间戳
	VtsMax *uint64 `json:"VtsMax,omitempty" name:"VtsMax"`

	// 创建时间最小时间戳
	CtsMin *uint64 `json:"CtsMin,omitempty" name:"CtsMin"`

	// 创建时间最大时间戳
	CtsMax *uint64 `json:"CtsMax,omitempty" name:"CtsMax"`

	// 偏移量
	Skip *uint64 `json:"Skip,omitempty" name:"Skip"`

	// 限制条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 策略名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 排序参数
	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

type DescribeAutoDenyIPRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 查询IP自动封禁状态
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 计数标识
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 类别
	Category *string `json:"Category,omitempty" name:"Category"`

	// 有效时间最小时间戳
	VtsMin *uint64 `json:"VtsMin,omitempty" name:"VtsMin"`

	// 有效时间最大时间戳
	VtsMax *uint64 `json:"VtsMax,omitempty" name:"VtsMax"`

	// 创建时间最小时间戳
	CtsMin *uint64 `json:"CtsMin,omitempty" name:"CtsMin"`

	// 创建时间最大时间戳
	CtsMax *uint64 `json:"CtsMax,omitempty" name:"CtsMax"`

	// 偏移量
	Skip *uint64 `json:"Skip,omitempty" name:"Skip"`

	// 限制条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 策略名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 排序参数
	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeAutoDenyIPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoDenyIPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Ip")
	delete(f, "Count")
	delete(f, "Category")
	delete(f, "VtsMin")
	delete(f, "VtsMax")
	delete(f, "CtsMin")
	delete(f, "CtsMax")
	delete(f, "Skip")
	delete(f, "Limit")
	delete(f, "Name")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoDenyIPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoDenyIPResponseParams struct {
	// 查询IP封禁状态返回结果
	Data *IpHitItemsData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAutoDenyIPResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoDenyIPResponseParams `json:"Response"`
}

func (r *DescribeAutoDenyIPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoDenyIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCiphersDetailRequestParams struct {

}

type DescribeCiphersDetailRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCiphersDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCiphersDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCiphersDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCiphersDetailResponseParams struct {
	// 加密套件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ciphers []*TLSCiphers `json:"Ciphers,omitempty" name:"Ciphers"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCiphersDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCiphersDetailResponseParams `json:"Response"`
}

func (r *DescribeCiphersDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCiphersDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomRulesRspRuleListItem struct {
	// 动作类型
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 跳过的策略
	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 重定向地址
	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`

	// 策略ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 优先级
	SortId *string `json:"SortId,omitempty" name:"SortId"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies"`

	// 事件id
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

// Predefined struct for user
type DescribeCustomWhiteRuleRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤数组,name可以是如下的值： RuleID,RuleName,Match
	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`

	// asc或者desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// exp_ts或者mod_ts
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeCustomWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤数组,name可以是如下的值： RuleID,RuleName,Match
	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`

	// asc或者desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// exp_ts或者mod_ts
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeCustomWhiteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomWhiteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomWhiteRuleResponseParams struct {
	// 规则详情
	RuleList []*DescribeCustomRulesRspRuleListItem `json:"RuleList,omitempty" name:"RuleList"`

	// 规则条数
	TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCustomWhiteRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomWhiteRuleResponseParams `json:"Response"`
}

func (r *DescribeCustomWhiteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainCountInfoRequestParams struct {

}

type DescribeDomainCountInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDomainCountInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainCountInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainCountInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainCountInfoResponseParams struct {
	// 域名总数
	AllDomain *uint64 `json:"AllDomain,omitempty" name:"AllDomain"`

	// 最近发现时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 接入域名总数
	WafDomainCount *uint64 `json:"WafDomainCount,omitempty" name:"WafDomainCount"`

	// 剩下配额
	LeftDomainCount *uint64 `json:"LeftDomainCount,omitempty" name:"LeftDomainCount"`

	// 开启防护域名数
	OpenWafDomain *uint64 `json:"OpenWafDomain,omitempty" name:"OpenWafDomain"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDomainCountInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainCountInfoResponseParams `json:"Response"`
}

func (r *DescribeDomainCountInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainCountInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainDetailsClbRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名id
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeDomainDetailsClbRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名id
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDomainDetailsClbRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainDetailsClbRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainDetailsClbRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainDetailsClbResponseParams struct {
	// clb域名详情
	DomainsClbPartInfo *ClbDomainsInfo `json:"DomainsClbPartInfo,omitempty" name:"DomainsClbPartInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDomainDetailsClbResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainDetailsClbResponseParams `json:"Response"`
}

func (r *DescribeDomainDetailsClbResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainDetailsClbResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainDetailsSaasRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名id
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeDomainDetailsSaasRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名id
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDomainDetailsSaasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainDetailsSaasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainDetailsSaasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainDetailsSaasResponseParams struct {
	// 域名详情
	DomainsPartInfo *DomainsPartInfo `json:"DomainsPartInfo,omitempty" name:"DomainsPartInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDomainDetailsSaasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainDetailsSaasResponseParams `json:"Response"`
}

func (r *DescribeDomainDetailsSaasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainDetailsSaasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainWhiteRulesRequestParams struct {
	// 需要查询的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 请求的白名单匹配路径
	Url *string `json:"Url,omitempty" name:"Url"`

	// 翻到多少页
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页展示的条数
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 排序方式,desc表示降序，asc表示升序
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type DescribeDomainWhiteRulesRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 请求的白名单匹配路径
	Url *string `json:"Url,omitempty" name:"Url"`

	// 翻到多少页
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页展示的条数
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 排序方式,desc表示降序，asc表示升序
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DescribeDomainWhiteRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainWhiteRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Url")
	delete(f, "Page")
	delete(f, "Count")
	delete(f, "Sort")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainWhiteRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainWhiteRulesResponseParams struct {
	// 规则列表
	RuleList []*RuleList `json:"RuleList,omitempty" name:"RuleList"`

	// 规则的数量
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDomainWhiteRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainWhiteRulesResponseParams `json:"Response"`
}

func (r *DescribeDomainWhiteRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainWhiteRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsRequestParams struct {
	// 数据偏移量，从1开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回域名的数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤数组
	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

type DescribeDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 数据偏移量，从1开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回域名的数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤数组
	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// domain列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domains []*DomainInfo `json:"Domains,omitempty" name:"Domains"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainsResponseParams `json:"Response"`
}

func (r *DescribeDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFindDomainListRequestParams struct {
	// 分页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页容量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件
	Key *string `json:"Key,omitempty" name:"Key"`

	// 是否接入waf
	IsWafDomain *string `json:"IsWafDomain,omitempty" name:"IsWafDomain"`

	// 排序参数
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeFindDomainListRequest struct {
	*tchttp.BaseRequest
	
	// 分页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页容量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件
	Key *string `json:"Key,omitempty" name:"Key"`

	// 是否接入waf
	IsWafDomain *string `json:"IsWafDomain,omitempty" name:"IsWafDomain"`

	// 排序参数
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeFindDomainListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFindDomainListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Key")
	delete(f, "IsWafDomain")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFindDomainListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFindDomainListResponseParams struct {
	// 域名总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 域名信息列表
	List []*FindAllDomainDetail `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFindDomainListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFindDomainListResponseParams `json:"Response"`
}

func (r *DescribeFindDomainListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFindDomainListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowTrendRequestParams struct {
	// 需要获取流量趋势的域名, all表示所有域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 起始时间戳，精度秒
	StartTs *int64 `json:"StartTs,omitempty" name:"StartTs"`

	// 结束时间戳，精度秒
	EndTs *int64 `json:"EndTs,omitempty" name:"EndTs"`
}

type DescribeFlowTrendRequest struct {
	*tchttp.BaseRequest
	
	// 需要获取流量趋势的域名, all表示所有域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 起始时间戳，精度秒
	StartTs *int64 `json:"StartTs,omitempty" name:"StartTs"`

	// 结束时间戳，精度秒
	EndTs *int64 `json:"EndTs,omitempty" name:"EndTs"`
}

func (r *DescribeFlowTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "StartTs")
	delete(f, "EndTs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowTrendResponseParams struct {
	// 流量趋势数据
	Data []*BotStatPointItem `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFlowTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowTrendResponseParams `json:"Response"`
}

func (r *DescribeFlowTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostLimitRequestParams struct {
	// 添加的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 流量来源
	AlbType *string `json:"AlbType,omitempty" name:"AlbType"`
}

type DescribeHostLimitRequest struct {
	*tchttp.BaseRequest
	
	// 添加的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 流量来源
	AlbType *string `json:"AlbType,omitempty" name:"AlbType"`
}

func (r *DescribeHostLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "InstanceID")
	delete(f, "AlbType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostLimitResponseParams struct {
	// 成功返回的状态码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Success *ResponseCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostLimitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostLimitResponseParams `json:"Response"`
}

func (r *DescribeHostLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type DescribeHostRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribeHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostResponseParams struct {
	// 域名详情
	Host *HostRecord `json:"Host,omitempty" name:"Host"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostResponseParams `json:"Response"`
}

func (r *DescribeHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsRequestParams struct {
	// 防护域名，如果是要查询某一具体的防护域名则传入此参数，要求是准确的域名，此参数不支持模糊搜索
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 防护域名ID，如果是要查询某一具体的防护域名则传入此参数，要求是准确的域名ID，此参数不支持模糊搜索
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 搜索条件，根据此参数对域名做模糊搜索
	Search *string `json:"Search,omitempty" name:"Search"`

	// 复杂的搜索条件
	Item *SearchItem `json:"Item,omitempty" name:"Item"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type DescribeHostsRequest struct {
	*tchttp.BaseRequest
	
	// 防护域名，如果是要查询某一具体的防护域名则传入此参数，要求是准确的域名，此参数不支持模糊搜索
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 防护域名ID，如果是要查询某一具体的防护域名则传入此参数，要求是准确的域名ID，此参数不支持模糊搜索
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 搜索条件，根据此参数对域名做模糊搜索
	Search *string `json:"Search,omitempty" name:"Search"`

	// 复杂的搜索条件
	Item *SearchItem `json:"Item,omitempty" name:"Item"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribeHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "Search")
	delete(f, "Item")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsResponseParams struct {
	// 防护域名列表的长度
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 防护域名的列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostList []*HostRecord `json:"HostList,omitempty" name:"HostList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostsResponseParams `json:"Response"`
}

func (r *DescribeHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤数组
	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤数组
	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// instance列表
	Instances []*InstanceInfo `json:"Instances,omitempty" name:"Instances"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesResponseParams `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpAccessControlRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 计数标识
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 动作，40表示查询白名单，42表示查询黑名单
	ActionType *uint64 `json:"ActionType,omitempty" name:"ActionType"`

	// 最小有效时间的时间戳
	VtsMin *uint64 `json:"VtsMin,omitempty" name:"VtsMin"`

	// 最大有效时间的时间戳
	VtsMax *uint64 `json:"VtsMax,omitempty" name:"VtsMax"`

	// 最小创建时间的时间戳
	CtsMin *uint64 `json:"CtsMin,omitempty" name:"CtsMin"`

	// 最大创建时间的时间戳
	CtsMax *uint64 `json:"CtsMax,omitempty" name:"CtsMax"`

	// 分页开始条数
	OffSet *uint64 `json:"OffSet,omitempty" name:"OffSet"`

	// 每页的条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 来源
	Source *string `json:"Source,omitempty" name:"Source"`

	// 排序参数
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type DescribeIpAccessControlRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 计数标识
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 动作，40表示查询白名单，42表示查询黑名单
	ActionType *uint64 `json:"ActionType,omitempty" name:"ActionType"`

	// 最小有效时间的时间戳
	VtsMin *uint64 `json:"VtsMin,omitempty" name:"VtsMin"`

	// 最大有效时间的时间戳
	VtsMax *uint64 `json:"VtsMax,omitempty" name:"VtsMax"`

	// 最小创建时间的时间戳
	CtsMin *uint64 `json:"CtsMin,omitempty" name:"CtsMin"`

	// 最大创建时间的时间戳
	CtsMax *uint64 `json:"CtsMax,omitempty" name:"CtsMax"`

	// 分页开始条数
	OffSet *uint64 `json:"OffSet,omitempty" name:"OffSet"`

	// 每页的条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 来源
	Source *string `json:"Source,omitempty" name:"Source"`

	// 排序参数
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

func (r *DescribeIpAccessControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpAccessControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Count")
	delete(f, "ActionType")
	delete(f, "VtsMin")
	delete(f, "VtsMax")
	delete(f, "CtsMin")
	delete(f, "CtsMax")
	delete(f, "OffSet")
	delete(f, "Limit")
	delete(f, "Source")
	delete(f, "Sort")
	delete(f, "Ip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpAccessControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpAccessControlResponseParams struct {
	// 输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *IpAccessControlData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIpAccessControlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpAccessControlResponseParams `json:"Response"`
}

func (r *DescribeIpAccessControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpAccessControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpHitItemsRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 计数标识
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 类别
	Category *string `json:"Category,omitempty" name:"Category"`

	// 有效时间最小时间戳
	VtsMin *uint64 `json:"VtsMin,omitempty" name:"VtsMin"`

	// 有效时间最大时间戳
	VtsMax *uint64 `json:"VtsMax,omitempty" name:"VtsMax"`

	// 创建时间最小时间戳
	CtsMin *uint64 `json:"CtsMin,omitempty" name:"CtsMin"`

	// 创建时间最大时间戳
	CtsMax *uint64 `json:"CtsMax,omitempty" name:"CtsMax"`

	// 偏移参数
	Skip *uint64 `json:"Skip,omitempty" name:"Skip"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 排序参数
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type DescribeIpHitItemsRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 计数标识
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 类别
	Category *string `json:"Category,omitempty" name:"Category"`

	// 有效时间最小时间戳
	VtsMin *uint64 `json:"VtsMin,omitempty" name:"VtsMin"`

	// 有效时间最大时间戳
	VtsMax *uint64 `json:"VtsMax,omitempty" name:"VtsMax"`

	// 创建时间最小时间戳
	CtsMin *uint64 `json:"CtsMin,omitempty" name:"CtsMin"`

	// 创建时间最大时间戳
	CtsMax *uint64 `json:"CtsMax,omitempty" name:"CtsMax"`

	// 偏移参数
	Skip *uint64 `json:"Skip,omitempty" name:"Skip"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 排序参数
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

func (r *DescribeIpHitItemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpHitItemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Count")
	delete(f, "Category")
	delete(f, "VtsMin")
	delete(f, "VtsMax")
	delete(f, "CtsMin")
	delete(f, "CtsMax")
	delete(f, "Skip")
	delete(f, "Limit")
	delete(f, "Name")
	delete(f, "Sort")
	delete(f, "Ip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpHitItemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpHitItemsResponseParams struct {
	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *IpHitItemsData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIpHitItemsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpHitItemsResponseParams `json:"Response"`
}

func (r *DescribeIpHitItemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpHitItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePeakPointsRequestParams struct {
	// 查询起始时间
	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`

	// 查询终止时间
	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`

	// 查询的域名，如果查询所有域名数据，该参数不填写
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 十三个值可选：
	// access-峰值qps趋势图
	// botAccess- bot峰值qps趋势图
	// down-下行峰值带宽趋势图
	// up-上行峰值带宽趋势图
	// attack-Web攻击总数趋势图
	// cc-CC攻击总数趋势图
	// StatusServerError-WAF返回给客户端状态码次数趋势图
	// StatusClientError-WAF返回给客户端状态码次数趋势图
	// StatusRedirect-WAF返回给客户端状态码次数趋势图
	// StatusOk-WAF返回给客户端状态码次数趋势图
	// UpstreamServerError-源站返回给WAF状态码次数趋势图
	// UpstreamClientError-源站返回给WAF状态码次数趋势图
	// UpstreamRedirect-源站返回给WAF状态码次数趋势图
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

type DescribePeakPointsRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间
	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`

	// 查询终止时间
	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`

	// 查询的域名，如果查询所有域名数据，该参数不填写
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 十三个值可选：
	// access-峰值qps趋势图
	// botAccess- bot峰值qps趋势图
	// down-下行峰值带宽趋势图
	// up-上行峰值带宽趋势图
	// attack-Web攻击总数趋势图
	// cc-CC攻击总数趋势图
	// StatusServerError-WAF返回给客户端状态码次数趋势图
	// StatusClientError-WAF返回给客户端状态码次数趋势图
	// StatusRedirect-WAF返回给客户端状态码次数趋势图
	// StatusOk-WAF返回给客户端状态码次数趋势图
	// UpstreamServerError-源站返回给WAF状态码次数趋势图
	// UpstreamClientError-源站返回给WAF状态码次数趋势图
	// UpstreamRedirect-源站返回给WAF状态码次数趋势图
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

func (r *DescribePeakPointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePeakPointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromTime")
	delete(f, "ToTime")
	delete(f, "Domain")
	delete(f, "Edition")
	delete(f, "InstanceID")
	delete(f, "MetricName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePeakPointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePeakPointsResponseParams struct {
	// 数据点
	Points []*PeakPointsItem `json:"Points,omitempty" name:"Points"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePeakPointsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePeakPointsResponseParams `json:"Response"`
}

func (r *DescribePeakPointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePeakPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePeakValueRequestParams struct {
	// 查询起始时间
	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`

	// 需要查询的域名，当前用户所有域名可以不传
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 五个值可选：
	// access-峰值qps
	// down-下行峰值带宽
	// up-上行峰值带宽
	// attack-Web攻击总数
	// cc-CC攻击总数趋势图
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

type DescribePeakValueRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间
	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`

	// 需要查询的域名，当前用户所有域名可以不传
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 五个值可选：
	// access-峰值qps
	// down-下行峰值带宽
	// up-上行峰值带宽
	// attack-Web攻击总数
	// cc-CC攻击总数趋势图
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

func (r *DescribePeakValueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePeakValueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromTime")
	delete(f, "ToTime")
	delete(f, "Domain")
	delete(f, "Edition")
	delete(f, "InstanceID")
	delete(f, "MetricName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePeakValueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePeakValueResponseParams struct {
	// QPS峰值
	Access *uint64 `json:"Access,omitempty" name:"Access"`

	// 上行带宽峰值，单位B
	Up *uint64 `json:"Up,omitempty" name:"Up"`

	// 下行带宽峰值，单位B
	Down *uint64 `json:"Down,omitempty" name:"Down"`

	// Web攻击总数
	Attack *uint64 `json:"Attack,omitempty" name:"Attack"`

	// CC攻击总数
	Cc *uint64 `json:"Cc,omitempty" name:"Cc"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePeakValueResponse struct {
	*tchttp.BaseResponse
	Response *DescribePeakValueResponseParams `json:"Response"`
}

func (r *DescribePeakValueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePeakValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyStatusRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// clb-waf或者saas-waf
	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

type DescribePolicyStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// clb-waf或者saas-waf
	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *DescribePolicyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Edition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePolicyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyStatusResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 防护状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePolicyStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribePolicyStatusResponseParams `json:"Response"`
}

func (r *DescribePolicyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleLimitRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type DescribeRuleLimitRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeRuleLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleLimitResponseParams struct {
	// waf模块的规格
	Res *WafRuleLimit `json:"Res,omitempty" name:"Res"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRuleLimitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleLimitResponseParams `json:"Response"`
}

func (r *DescribeRuleLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTlsVersionRequestParams struct {

}

type DescribeTlsVersionRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeTlsVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTlsVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTlsVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTlsVersionResponseParams struct {
	// TLS key value
	TLS []*TLSVersion `json:"TLS,omitempty" name:"TLS"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTlsVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTlsVersionResponseParams `json:"Response"`
}

func (r *DescribeTlsVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTlsVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserCdcClbWafRegionsRequestParams struct {

}

type DescribeUserCdcClbWafRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserCdcClbWafRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserCdcClbWafRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserCdcClbWafRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserCdcClbWafRegionsResponseParams struct {
	// CdcRegion的类型描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*CdcRegion `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserCdcClbWafRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserCdcClbWafRegionsResponseParams `json:"Response"`
}

func (r *DescribeUserCdcClbWafRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserCdcClbWafRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserClbWafRegionsRequestParams struct {

}

type DescribeUserClbWafRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserClbWafRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserClbWafRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserClbWafRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserClbWafRegionsResponseParams struct {
	// 地域（标准的ap-格式）列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserClbWafRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserClbWafRegionsResponseParams `json:"Response"`
}

func (r *DescribeUserClbWafRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserClbWafRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserDomainInfoRequestParams struct {

}

type DescribeUserDomainInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserDomainInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDomainInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserDomainInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserDomainInfoResponseParams struct {
	// saas和clb域名信息
	UsersInfo []*UserDomainInfo `json:"UsersInfo,omitempty" name:"UsersInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserDomainInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserDomainInfoResponseParams `json:"Response"`
}

func (r *DescribeUserDomainInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDomainInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVipInfoRequestParams struct {
	// waf实例id列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DescribeVipInfoRequest struct {
	*tchttp.BaseRequest
	
	// waf实例id列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeVipInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVipInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVipInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVipInfoResponseParams struct {
	// VIP信息
	VipInfo []*VipInfo `json:"VipInfo,omitempty" name:"VipInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVipInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVipInfoResponseParams `json:"Response"`
}

func (r *DescribeVipInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVipInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafAutoDenyRulesRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type DescribeWafAutoDenyRulesRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeWafAutoDenyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafAutoDenyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWafAutoDenyRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafAutoDenyRulesResponseParams struct {
	// 攻击次数阈值
	AttackThreshold *int64 `json:"AttackThreshold,omitempty" name:"AttackThreshold"`

	// 攻击时间阈值
	TimeThreshold *int64 `json:"TimeThreshold,omitempty" name:"TimeThreshold"`

	// 自动封禁时间
	DenyTimeThreshold *int64 `json:"DenyTimeThreshold,omitempty" name:"DenyTimeThreshold"`

	// 自动封禁状态
	DefenseStatus *int64 `json:"DefenseStatus,omitempty" name:"DefenseStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWafAutoDenyRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWafAutoDenyRulesResponseParams `json:"Response"`
}

func (r *DescribeWafAutoDenyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafAutoDenyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafAutoDenyStatusRequestParams struct {

}

type DescribeWafAutoDenyStatusRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeWafAutoDenyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafAutoDenyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWafAutoDenyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafAutoDenyStatusResponseParams struct {
	// WAF 自动封禁详情
	WafAutoDenyDetails *AutoDenyDetail `json:"WafAutoDenyDetails,omitempty" name:"WafAutoDenyDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWafAutoDenyStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWafAutoDenyStatusResponseParams `json:"Response"`
}

func (r *DescribeWafAutoDenyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafAutoDenyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafInfoRequestParams struct {
	// CLB回调WAF接口（获取、删除）的参数
	Params []*ClbHostsParams `json:"Params,omitempty" name:"Params"`
}

type DescribeWafInfoRequest struct {
	*tchttp.BaseRequest
	
	// CLB回调WAF接口（获取、删除）的参数
	Params []*ClbHostsParams `json:"Params,omitempty" name:"Params"`
}

func (r *DescribeWafInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Params")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWafInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafInfoResponseParams struct {
	// 返回的WAF信息数组的长度，为0则表示没有查询到对应的信息
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 对应的WAF信息的数组。
	HostList []*ClbHostResult `json:"HostList,omitempty" name:"HostList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWafInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWafInfoResponseParams `json:"Response"`
}

func (r *DescribeWafInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafThreatenIntelligenceRequestParams struct {

}

type DescribeWafThreatenIntelligenceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeWafThreatenIntelligenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafThreatenIntelligenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWafThreatenIntelligenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafThreatenIntelligenceResponseParams struct {
	// WAF 威胁情报封禁信息
	WafThreatenIntelligenceDetails *WafThreatenIntelligenceDetails `json:"WafThreatenIntelligenceDetails,omitempty" name:"WafThreatenIntelligenceDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWafThreatenIntelligenceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWafThreatenIntelligenceResponseParams `json:"Response"`
}

func (r *DescribeWafThreatenIntelligenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafThreatenIntelligenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainInfo struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// cname地址
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// 实例类型,sparta-waf表示saaswaf实例域名,clb-waf表示clbwaf实例域名,cdc-clb-waf表示CDC环境下clbwaf实例域名,cdn-waf表示cdnwaf实例域名
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 日志包
	ClsStatus *uint64 `json:"ClsStatus,omitempty" name:"ClsStatus"`

	// clbwaf使用模式,0镜像模式 1清洗模式
	FlowMode *uint64 `json:"FlowMode,omitempty" name:"FlowMode"`

	// waf开关,0关闭 1开启
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 规则防御模式,0观察模式 1拦截模式
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// AI防御模式,10规则引擎观察&&AI引擎关闭模式 11规则引擎观察&&AI引擎观察模式 12规则引擎观察&&AI引擎拦截模式 20规则引擎拦截&&AI引擎关闭模式 21规则引擎拦截&&AI引擎观察模式 22规则引擎拦截&&AI引擎拦截模式
	Engine *uint64 `json:"Engine,omitempty" name:"Engine"`

	// CC列表
	CCList []*string `json:"CCList,omitempty" name:"CCList"`

	// 回源ip
	RsList []*string `json:"RsList,omitempty" name:"RsList"`

	// 服务端口配置
	Ports []*PortInfo `json:"Ports,omitempty" name:"Ports"`

	// 负载均衡器
	LoadBalancerSet []*LoadBalancerPackageNew `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`

	// 用户id
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// clbwaf域名监听器状态,0操作成功 4正在绑定LB 6正在解绑LB 7解绑LB失败 8绑定LB失败 10内部错误
	State *int64 `json:"State,omitempty" name:"State"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Ipv6开关状态,0关闭 1开启
	Ipv6Status *int64 `json:"Ipv6Status,omitempty" name:"Ipv6Status"`

	// BOT开关状态,0关闭 1开启
	BotStatus *int64 `json:"BotStatus,omitempty" name:"BotStatus"`

	// 版本信息
	Level *int64 `json:"Level,omitempty" name:"Level"`

	// 是否开启投递CLS功能,0关闭 1开启
	PostCLSStatus *int64 `json:"PostCLSStatus,omitempty" name:"PostCLSStatus"`

	// 是否开启投递CKafka功能,0关闭 1开启
	PostCKafkaStatus *int64 `json:"PostCKafkaStatus,omitempty" name:"PostCKafkaStatus"`

	// cdc实例域名接入的集群信息,非cdc实例忽略
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdcClusters *string `json:"CdcClusters,omitempty" name:"CdcClusters"`

	// api安全开关状态,0关闭 1开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiStatus *int64 `json:"ApiStatus,omitempty" name:"ApiStatus"`

	// 应用型负载均衡类型,clb或者apisix，默认clb
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlbType *string `json:"AlbType,omitempty" name:"AlbType"`

	// 安全组状态,0不展示 1非腾讯云源站 2安全组绑定失败 3安全组发生变更
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgState *int64 `json:"SgState,omitempty" name:"SgState"`

	// 安全组状态的详细解释
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgDetail *string `json:"SgDetail,omitempty" name:"SgDetail"`
}

type DomainPackageNew struct {
	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`

	// 是否自动续费，1：自动续费，0：不自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 套餐购买个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 套餐购买地域，clb-waf暂时没有用到
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`
}

type DomainURI struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 版本
	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

type DomainsPartInfo struct {
	// 是否开启httpRewrite
	HttpsRewrite *uint64 `json:"HttpsRewrite,omitempty" name:"HttpsRewrite"`

	// https回源端口
	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitempty" name:"HttpsUpstreamPort"`

	// 是否是cdn
	IsCdn *uint64 `json:"IsCdn,omitempty" name:"IsCdn"`

	// 是否开启gray
	IsGray *uint64 `json:"IsGray,omitempty" name:"IsGray"`

	// 是否是http2
	IsHttp2 *uint64 `json:"IsHttp2,omitempty" name:"IsHttp2"`

	// 是否开启websocket
	IsWebsocket *uint64 `json:"IsWebsocket,omitempty" name:"IsWebsocket"`

	// 负载均衡
	LoadBalance *uint64 `json:"LoadBalance,omitempty" name:"LoadBalance"`

	// 防御模式
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// 私钥
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// ssl id
	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`

	// 回源域名
	UpstreamDomain *string `json:"UpstreamDomain,omitempty" name:"UpstreamDomain"`

	// 回源类型
	UpstreamType *uint64 `json:"UpstreamType,omitempty" name:"UpstreamType"`

	// 回源ip
	SrcList []*string `json:"SrcList,omitempty" name:"SrcList"`

	// 服务端口配置
	Ports []*PortInfo `json:"Ports,omitempty" name:"Ports"`

	// 证书类型
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// 回源方式
	UpstreamScheme *string `json:"UpstreamScheme,omitempty" name:"UpstreamScheme"`

	// 日志包
	Cls *uint64 `json:"Cls,omitempty" name:"Cls"`

	// 一级cname
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// 是否长连接
	IsKeepAlive *uint64 `json:"IsKeepAlive,omitempty" name:"IsKeepAlive"`

	// 是否开启主动健康检测，1表示开启，0表示不开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveCheck *uint64 `json:"ActiveCheck,omitempty" name:"ActiveCheck"`

	// TLS版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TLSVersion *int64 `json:"TLSVersion,omitempty" name:"TLSVersion"`

	// 加密套件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ciphers []*int64 `json:"Ciphers,omitempty" name:"Ciphers"`

	// 模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	CipherTemplate *int64 `json:"CipherTemplate,omitempty" name:"CipherTemplate"`

	// 300s
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitempty" name:"ProxyReadTimeout"`

	// 300s
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitempty" name:"ProxySendTimeout"`

	// 0:关闭SNI；1:开启SNI，SNI=源请求host；2:开启SNI，SNI=修改为源站host；3：开启SNI，自定义host，SNI=SniHost；
	// 注意：此字段可能返回 null，表示取不到有效值。
	SniType *int64 `json:"SniType,omitempty" name:"SniType"`

	// SniType=3时，需要填此参数，表示自定义的host；
	// 注意：此字段可能返回 null，表示取不到有效值。
	SniHost *string `json:"SniHost,omitempty" name:"SniHost"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weights []*string `json:"Weights,omitempty" name:"Weights"`

	// IsCdn=3时，表示自定义header
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpHeaders []*string `json:"IpHeaders,omitempty" name:"IpHeaders"`

	// 0:关闭xff重置；1:开启xff重置
	// 注意：此字段可能返回 null，表示取不到有效值。
	XFFReset *int64 `json:"XFFReset,omitempty" name:"XFFReset"`
}

type DownloadAttackRecordInfo struct {
	// 记录ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 下载任务名
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 域名
	Host *string `json:"Host,omitempty" name:"Host"`

	// 当前下载任务的日志条数
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 下载任务运行状态：-1-下载超时，0-下载等待，1-下载完成，2-下载失败，4-正在下载
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 下载文件URL
	Url *string `json:"Url,omitempty" name:"Url"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最后更新修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 下载任务需下载的日志总条数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type ExportAccessInfo struct {
	// 日志导出任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`

	// 日志导出查询语句
	// 注意：此字段可能返回 null，表示取不到有效值。
	Query *string `json:"Query,omitempty" name:"Query"`

	// 日志导出文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 日志文件大小
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// 日志导出时间排序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Order *string `json:"Order,omitempty" name:"Order"`

	// 日志导出格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Format *string `json:"Format,omitempty" name:"Format"`

	// 日志导出数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 日志下载状态。Processing:导出正在进行中，Complete:导出完成，Failed:导出失败，Expired:日志导出已过期（三天有效期）
	Status *string `json:"Status,omitempty" name:"Status"`

	// 日志导出起始时间
	From *int64 `json:"From,omitempty" name:"From"`

	// 日志导出结束时间
	To *int64 `json:"To,omitempty" name:"To"`

	// 日志导出路径
	CosPath *string `json:"CosPath,omitempty" name:"CosPath"`

	// 日志导出创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type FailedInfo struct {
	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 失败信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type FiltersItemNew struct {
	// 字段名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤值
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否精确查找
	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type FindAllDomainDetail struct {
	// 用户id
	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名ip
	Ips []*string `json:"Ips,omitempty" name:"Ips"`

	// 发现时间
	FindTime *string `json:"FindTime,omitempty" name:"FindTime"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 域名id
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// waf类型
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 是否接入waf
	IsWafDomain *uint64 `json:"IsWafDomain,omitempty" name:"IsWafDomain"`
}

type FraudPkg struct {
	// 资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *int64 `json:"Region,omitempty" name:"Region"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 申请数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InquireNum *int64 `json:"InquireNum,omitempty" name:"InquireNum"`

	// 使用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedNum *int64 `json:"UsedNum,omitempty" name:"UsedNum"`

	// 续费标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

// Predefined struct for user
type GetAttackDownloadRecordsRequestParams struct {

}

type GetAttackDownloadRecordsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetAttackDownloadRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAttackDownloadRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAttackDownloadRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAttackDownloadRecordsResponseParams struct {
	// 下载攻击日志记录数组
	Records []*DownloadAttackRecordInfo `json:"Records,omitempty" name:"Records"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetAttackDownloadRecordsResponse struct {
	*tchttp.BaseResponse
	Response *GetAttackDownloadRecordsResponseParams `json:"Response"`
}

func (r *GetAttackDownloadRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAttackDownloadRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAttackHistogramRequestParams struct {
	// 查询的域名，所有域名使用all
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Lucene语法
	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
}

type GetAttackHistogramRequest struct {
	*tchttp.BaseRequest
	
	// 查询的域名，所有域名使用all
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Lucene语法
	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
}

func (r *GetAttackHistogramRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAttackHistogramRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QueryString")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAttackHistogramRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAttackHistogramResponseParams struct {
	// 统计详情
	Data []*LogHistogramInfo `json:"Data,omitempty" name:"Data"`

	// 时间段大小
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 统计的条目数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetAttackHistogramResponse struct {
	*tchttp.BaseResponse
	Response *GetAttackHistogramResponseParams `json:"Response"`
}

func (r *GetAttackHistogramResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAttackHistogramResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAttackTotalCountRequestParams struct {
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询的域名，全部域名不指定
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 查询条件，默认为""
	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
}

type GetAttackTotalCountRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询的域名，全部域名不指定
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 查询条件，默认为""
	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
}

func (r *GetAttackTotalCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAttackTotalCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Domain")
	delete(f, "QueryString")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAttackTotalCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAttackTotalCountResponseParams struct {
	// 攻击总次数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetAttackTotalCountResponse struct {
	*tchttp.BaseResponse
	Response *GetAttackTotalCountResponseParams `json:"Response"`
}

func (r *GetAttackTotalCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAttackTotalCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostDel struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 实例类型
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type HostRecord struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 主域名，入参时为空
	MainDomain *string `json:"MainDomain,omitempty" name:"MainDomain"`

	// waf模式，同saas waf保持一致
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// waf和LD的绑定，0：没有绑定，1：绑定
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 域名状态，0：正常，1：未检测到流量，2：即将过期，3：过期
	State *uint64 `json:"State,omitempty" name:"State"`

	// 使用的规则，同saas waf保持一致
	Engine *uint64 `json:"Engine,omitempty" name:"Engine"`

	// 是否开启代理，0：不开启，1：开启
	IsCdn *uint64 `json:"IsCdn,omitempty" name:"IsCdn"`

	// 绑定的LB列表
	LoadBalancerSet []*LoadBalancer `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`

	// 域名绑定的LB的地域，以,分割多个地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 产品分类，取值为：sparta-waf、clb-waf、cdn-waf
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// WAF的流量模式，1：清洗模式，0：镜像模式
	FlowMode *uint64 `json:"FlowMode,omitempty" name:"FlowMode"`

	// 是否开启访问日志，1：开启，0：关闭
	ClsStatus *uint64 `json:"ClsStatus,omitempty" name:"ClsStatus"`

	// 防护等级，可选值100,200,300
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// 域名需要下发到的cdc集群列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdcClusters []*string `json:"CdcClusters,omitempty" name:"CdcClusters"`

	// 应用型负载均衡类型: clb或者apisix，默认clb
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlbType *string `json:"AlbType,omitempty" name:"AlbType"`

	// IsCdn=3时，需要填此参数，表示自定义header
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpHeaders []*string `json:"IpHeaders,omitempty" name:"IpHeaders"`

	// 规则引擎类型， 1: menshen,   2:tiga
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineType *int64 `json:"EngineType,omitempty" name:"EngineType"`
}

type HostStatus struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// WAF的开关，1：开，0：关
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type InstanceInfo struct {
	// id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 资源id
	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 付费模式
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// 自动续费
	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 弹性计费
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// 套餐版本
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// 过期时间
	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 已用
	DomainCount *uint64 `json:"DomainCount,omitempty" name:"DomainCount"`

	// 上限
	SubDomainLimit *uint64 `json:"SubDomainLimit,omitempty" name:"SubDomainLimit"`

	// 已用
	MainDomainCount *uint64 `json:"MainDomainCount,omitempty" name:"MainDomainCount"`

	// 上限
	MainDomainLimit *uint64 `json:"MainDomainLimit,omitempty" name:"MainDomainLimit"`

	// 峰值
	MaxQPS *uint64 `json:"MaxQPS,omitempty" name:"MaxQPS"`

	// qps套餐
	QPS *QPSPackageNew `json:"QPS,omitempty" name:"QPS"`

	// 域名套餐
	DomainPkg *DomainPackageNew `json:"DomainPkg,omitempty" name:"DomainPkg"`

	// 用户appid
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// clb或saas
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 业务安全包
	// 注意：此字段可能返回 null，表示取不到有效值。
	FraudPkg *FraudPkg `json:"FraudPkg,omitempty" name:"FraudPkg"`

	// Bot资源包
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotPkg *BotPkg `json:"BotPkg,omitempty" name:"BotPkg"`

	// bot的qps详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotQPS *BotQPS `json:"BotQPS,omitempty" name:"BotQPS"`

	// qps弹性计费上限
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElasticBilling *uint64 `json:"ElasticBilling,omitempty" name:"ElasticBilling"`

	// 攻击日志投递开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackLogPost *int64 `json:"AttackLogPost,omitempty" name:"AttackLogPost"`

	// 带宽峰值，单位为B/s(字节每秒)
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxBandwidth *uint64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// api安全是否购买
	APISecurity *uint64 `json:"APISecurity,omitempty" name:"APISecurity"`

	// 购买的qps规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	QpsStandard *uint64 `json:"QpsStandard,omitempty" name:"QpsStandard"`

	// 购买的带宽规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthStandard *uint64 `json:"BandwidthStandard,omitempty" name:"BandwidthStandard"`

	// 实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 实例沙箱值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SandboxQps *uint64 `json:"SandboxQps,omitempty" name:"SandboxQps"`

	// 是否api 安全试用
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAPISecurityTrial *uint64 `json:"IsAPISecurityTrial,omitempty" name:"IsAPISecurityTrial"`
}

type IpAccessControlData struct {
	// ip黑白名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	Res []*IpAccessControlItem `json:"Res,omitempty" name:"Res"`

	// 计数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type IpAccessControlItem struct {
	// 动作
	ActionType *uint64 `json:"ActionType,omitempty" name:"ActionType"`

	// ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 备注
	Note *string `json:"Note,omitempty" name:"Note"`

	// 来源
	Source *string `json:"Source,omitempty" name:"Source"`

	// 更新时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsVersion *uint64 `json:"TsVersion,omitempty" name:"TsVersion"`

	// 有效截止时间戳
	ValidTs *uint64 `json:"ValidTs,omitempty" name:"ValidTs"`
}

type IpHitItem struct {
	// 动作
	Action *uint64 `json:"Action,omitempty" name:"Action"`

	// 类别
	Category *string `json:"Category,omitempty" name:"Category"`

	// ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 时间戳
	TsVersion *uint64 `json:"TsVersion,omitempty" name:"TsVersion"`

	// 有效截止时间戳
	ValidTs *uint64 `json:"ValidTs,omitempty" name:"ValidTs"`
}

type IpHitItemsData struct {
	// 数组封装
	Res []*IpHitItem `json:"Res,omitempty" name:"Res"`

	// 总数目
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type LoadBalancer struct {
	// 负载均衡LD的ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡LD的名称
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡监听器的ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 负载均衡监听器的名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 负载均衡实例的IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 负载均衡实例的端口
	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`

	// 负载均衡LD的地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 监听器协议，http、https
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 负载均衡监听器所在的zone
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 负载均衡的VPCID，公网为-1，内网按实际填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumericalVpcId *int64 `json:"NumericalVpcId,omitempty" name:"NumericalVpcId"`

	// 负载均衡的网络类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
}

type LoadBalancerPackageNew struct {
	// 监听id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 负载均衡id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡名
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 地区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 接入IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 接入端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// VPCID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumericalVpcId *int64 `json:"NumericalVpcId,omitempty" name:"NumericalVpcId"`

	// CLB类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
}

type LogHistogramInfo struct {
	// 日志条数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 时间戳
	TimeStamp *int64 `json:"TimeStamp,omitempty" name:"TimeStamp"`
}

// Predefined struct for user
type ModifyAccessPeriodRequestParams struct {
	// 访问日志保存期限，范围为[1, 30]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 日志主题
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

type ModifyAccessPeriodRequest struct {
	*tchttp.BaseRequest
	
	// 访问日志保存期限，范围为[1, 30]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 日志主题
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *ModifyAccessPeriodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessPeriodRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Period")
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccessPeriodRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccessPeriodResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAccessPeriodResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccessPeriodResponseParams `json:"Response"`
}

func (r *ModifyAccessPeriodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAreaBanStatusRequestParams struct {
	// 需要修改的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 状态值，0表示关闭，1表示开启
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ModifyAreaBanStatusRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 状态值，0表示关闭，1表示开启
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAreaBanStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAreaBanStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAreaBanStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAreaBanStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAreaBanStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAreaBanStatusResponseParams `json:"Response"`
}

func (r *ModifyAreaBanStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAreaBanStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomRuleStatusRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// 开关的状态，1是开启、0是关闭
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

type ModifyCustomRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// 开关的状态，1是开启、0是关闭
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *ModifyCustomRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RuleId")
	delete(f, "Status")
	delete(f, "Edition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomRuleStatusResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCustomRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomRuleStatusResponseParams `json:"Response"`
}

func (r *ModifyCustomRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomWhiteRuleRequestParams struct {
	// 编辑的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 编辑的规则ID
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// 编辑的规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 放行时是否继续执行其它检查逻辑，继续执行地域封禁防护：geoip、继续执行CC策略防护：cc、继续执行WEB应用防护：owasp、继续执行AI引擎防护：ai、继续执行信息防泄漏防护：antileakage。如果勾选多个，则以“，”串接。
	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`

	// 优先级，1~100的整数，数字越小，代表这条规则的执行优先级越高。
	SortId *uint64 `json:"SortId,omitempty" name:"SortId"`

	// 规则生效截止时间，0：永久生效，其它值为对应时间的时间戳。
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 匹配条件数组
	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies"`
}

type ModifyCustomWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 编辑的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 编辑的规则ID
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// 编辑的规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 放行时是否继续执行其它检查逻辑，继续执行地域封禁防护：geoip、继续执行CC策略防护：cc、继续执行WEB应用防护：owasp、继续执行AI引擎防护：ai、继续执行信息防泄漏防护：antileakage。如果勾选多个，则以“，”串接。
	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`

	// 优先级，1~100的整数，数字越小，代表这条规则的执行优先级越高。
	SortId *uint64 `json:"SortId,omitempty" name:"SortId"`

	// 规则生效截止时间，0：永久生效，其它值为对应时间的时间戳。
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 匹配条件数组
	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies"`
}

func (r *ModifyCustomWhiteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomWhiteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RuleId")
	delete(f, "RuleName")
	delete(f, "Bypass")
	delete(f, "SortId")
	delete(f, "ExpireTime")
	delete(f, "Strategies")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomWhiteRuleResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCustomWhiteRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomWhiteRuleResponseParams `json:"Response"`
}

func (r *ModifyCustomWhiteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainIpv6StatusRequestParams struct {
	// 需要修改的域名所属的实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 需要修改的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 需要修改的域名ID
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 修改域名的Ipv6开关为Status （0:关闭 1:开启）
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ModifyDomainIpv6StatusRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的域名所属的实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 需要修改的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 需要修改的域名ID
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 修改域名的Ipv6开关为Status （0:关闭 1:开启）
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDomainIpv6StatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainIpv6StatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainIpv6StatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainIpv6StatusResponseParams struct {
	// 返回的状态 （0: 操作失败 1:操作成功 2:企业版以上不支持 3:企业版以下不支持 ）
	Ipv6Status *int64 `json:"Ipv6Status,omitempty" name:"Ipv6Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDomainIpv6StatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainIpv6StatusResponseParams `json:"Response"`
}

func (r *ModifyDomainIpv6StatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainIpv6StatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainWhiteRuleRequestParams struct {
	// 需要更改的规则的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 白名单id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 规则的id列表
	Rules []*uint64 `json:"Rules,omitempty" name:"Rules"`

	// 规则匹配路径
	Url *string `json:"Url,omitempty" name:"Url"`

	// 规则匹配方法
	Function *string `json:"Function,omitempty" name:"Function"`

	// 规则的开关状态，0表示关闭开关，1表示打开开关
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type ModifyDomainWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要更改的规则的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 白名单id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 规则的id列表
	Rules []*uint64 `json:"Rules,omitempty" name:"Rules"`

	// 规则匹配路径
	Url *string `json:"Url,omitempty" name:"Url"`

	// 规则匹配方法
	Function *string `json:"Function,omitempty" name:"Function"`

	// 规则的开关状态，0表示关闭开关，1表示打开开关
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDomainWhiteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainWhiteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Id")
	delete(f, "Rules")
	delete(f, "Url")
	delete(f, "Function")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainWhiteRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDomainWhiteRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainWhiteRuleResponseParams `json:"Response"`
}

func (r *ModifyDomainWhiteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainsCLSStatusRequestParams struct {
	// 需要修改的域名列表
	Domains []*DomainURI `json:"Domains,omitempty" name:"Domains"`

	// 修改域名的访问日志开关为Status
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type ModifyDomainsCLSStatusRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的域名列表
	Domains []*DomainURI `json:"Domains,omitempty" name:"Domains"`

	// 修改域名的访问日志开关为Status
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDomainsCLSStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainsCLSStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainsCLSStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainsCLSStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDomainsCLSStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainsCLSStatusResponseParams `json:"Response"`
}

func (r *ModifyDomainsCLSStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainsCLSStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostFlowModeRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// WAF流量模式，1：清洗模式，0：镜像模式（默认）
	FlowMode *uint64 `json:"FlowMode,omitempty" name:"FlowMode"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type ModifyHostFlowModeRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// WAF流量模式，1：清洗模式，0：镜像模式（默认）
	FlowMode *uint64 `json:"FlowMode,omitempty" name:"FlowMode"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *ModifyHostFlowModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostFlowModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "FlowMode")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHostFlowModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostFlowModeResponseParams struct {
	// 成功的状态码
	Success *ResponseCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyHostFlowModeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHostFlowModeResponseParams `json:"Response"`
}

func (r *ModifyHostFlowModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostFlowModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostModeRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 防护状态：
	// 10：规则观察&&AI关闭模式，11：规则观察&&AI观察模式，12：规则观察&&AI拦截模式
	// 20：规则拦截&&AI关闭模式，21：规则拦截&&AI观察模式，22：规则拦截&&AI拦截模式
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// 0:修改防护模式，1:修改AI
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 实例类型
	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

type ModifyHostModeRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 防护状态：
	// 10：规则观察&&AI关闭模式，11：规则观察&&AI观察模式，12：规则观察&&AI拦截模式
	// 20：规则拦截&&AI关闭模式，21：规则拦截&&AI观察模式，22：规则拦截&&AI拦截模式
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// 0:修改防护模式，1:修改AI
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 实例类型
	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *ModifyHostModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "Mode")
	delete(f, "Type")
	delete(f, "InstanceID")
	delete(f, "Edition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHostModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostModeResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Success *ResponseCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyHostModeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHostModeResponseParams `json:"Response"`
}

func (r *ModifyHostModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostRequestParams struct {
	// 编辑的域名配置信息
	Host *HostRecord `json:"Host,omitempty" name:"Host"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type ModifyHostRequest struct {
	*tchttp.BaseRequest
	
	// 编辑的域名配置信息
	Host *HostRecord `json:"Host,omitempty" name:"Host"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *ModifyHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Host")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostResponseParams struct {
	// 编辑的域名ID
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyHostResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHostResponseParams `json:"Response"`
}

func (r *ModifyHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostStatusRequestParams struct {
	// 域名状态列表
	HostsStatus []*HostStatus `json:"HostsStatus,omitempty" name:"HostsStatus"`
}

type ModifyHostStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名状态列表
	HostsStatus []*HostStatus `json:"HostsStatus,omitempty" name:"HostsStatus"`
}

func (r *ModifyHostStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HostsStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHostStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostStatusResponseParams struct {
	// 成功的状态码，需要JSON解码后再使用，返回的格式是{"域名":"状态"}，成功的状态码为Success，其它的为失败的状态码（yunapi定义的错误码）
	Success *ResponseCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyHostStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHostStatusResponseParams `json:"Response"`
}

func (r *ModifyHostStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProtectionStatusRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

type ModifyProtectionStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *ModifyProtectionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProtectionStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Status")
	delete(f, "Edition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProtectionStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProtectionStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyProtectionStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProtectionStatusResponseParams `json:"Response"`
}

func (r *ModifyProtectionStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProtectionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySpartaProtectionModeRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 防护状态：
	// 10：规则观察&&AI关闭模式，11：规则观察&&AI观察模式，12：规则观察&&AI拦截模式
	// 20：规则拦截&&AI关闭模式，21：规则拦截&&AI观察模式，22：规则拦截&&AI拦截模式
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 0是修改规则引擎状态，1是修改AI的状态
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type ModifySpartaProtectionModeRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 防护状态：
	// 10：规则观察&&AI关闭模式，11：规则观察&&AI观察模式，12：规则观察&&AI拦截模式
	// 20：规则拦截&&AI关闭模式，21：规则拦截&&AI观察模式，22：规则拦截&&AI拦截模式
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 0是修改规则引擎状态，1是修改AI的状态
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *ModifySpartaProtectionModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySpartaProtectionModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Mode")
	delete(f, "Edition")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySpartaProtectionModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySpartaProtectionModeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySpartaProtectionModeResponse struct {
	*tchttp.BaseResponse
	Response *ModifySpartaProtectionModeResponseParams `json:"Response"`
}

func (r *ModifySpartaProtectionModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySpartaProtectionModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySpartaProtectionRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 证书类型，0表示没有证书，CertType=1表示自有证书,2 为托管证书
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// CertType=1时，需要填次参数，表示证书内容
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// CertType=1时，需要填次参数，表示证书的私钥
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// CertType=2时，需要填次参数，表示证书的ID
	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`

	// 表示是否开启了CDN代理
	IsCdn *int64 `json:"IsCdn,omitempty" name:"IsCdn"`

	// HTTPS回源协议
	UpstreamScheme *string `json:"UpstreamScheme,omitempty" name:"UpstreamScheme"`

	// HTTPS回源端口,仅UpstreamScheme为http时需要填当前字段
	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitempty" name:"HttpsUpstreamPort"`

	// 表示是否强制跳转到HTTPS，1表示开启，0表示不开启
	HttpsRewrite *uint64 `json:"HttpsRewrite,omitempty" name:"HttpsRewrite"`

	// 回源类型，0表示通过IP回源,1 表示通过域名回源
	UpstreamType *int64 `json:"UpstreamType,omitempty" name:"UpstreamType"`

	// UpstreamType=1时，填次字段表示回源域名
	UpstreamDomain *string `json:"UpstreamDomain,omitempty" name:"UpstreamDomain"`

	// UpstreamType=0时，填次字段表示回源ip
	SrcList []*string `json:"SrcList,omitempty" name:"SrcList"`

	// 是否开启HTTP2，1表示开启，0表示不开启http2。开启HTTP2需要HTTPS支持
	IsHttp2 *int64 `json:"IsHttp2,omitempty" name:"IsHttp2"`

	// 是否开启WebSocket， 1：开启WebSocket，0：不开启WebSocket
	IsWebsocket *int64 `json:"IsWebsocket,omitempty" name:"IsWebsocket"`

	// 负载均衡策略，0表示轮徇，1表示IP hash
	LoadBalance *int64 `json:"LoadBalance,omitempty" name:"LoadBalance"`

	// 是否灰度
	IsGray *int64 `json:"IsGray,omitempty" name:"IsGray"`

	// WAF版本
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 端口信息
	Ports []*SpartaProtectionPort `json:"Ports,omitempty" name:"Ports"`

	// 长短连接标志，仅IP回源时有效
	IsKeepAlive *string `json:"IsKeepAlive,omitempty" name:"IsKeepAlive"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 是否为Anycast ip类型：1 Anycast 0 普通ip
	Anycast *int64 `json:"Anycast,omitempty" name:"Anycast"`

	// src的权重
	Weights []*int64 `json:"Weights,omitempty" name:"Weights"`

	// 是否开启源站的主动健康检测，1表示开启，0表示不开启
	ActiveCheck *int64 `json:"ActiveCheck,omitempty" name:"ActiveCheck"`

	// TLS版本信息
	TLSVersion *int64 `json:"TLSVersion,omitempty" name:"TLSVersion"`

	// 加密套件信息
	Ciphers []*int64 `json:"Ciphers,omitempty" name:"Ciphers"`

	// 0:不支持选择：默认模板  1:通用型模板 2:安全型模板 3:自定义模板
	CipherTemplate *int64 `json:"CipherTemplate,omitempty" name:"CipherTemplate"`

	// 300s
	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitempty" name:"ProxyReadTimeout"`

	// 300s
	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitempty" name:"ProxySendTimeout"`

	// 0:关闭SNI；1:开启SNI，SNI=源请求host；2:开启SNI，SNI=修改为源站host；3：开启SNI，自定义host，SNI=SniHost；
	SniType *int64 `json:"SniType,omitempty" name:"SniType"`

	// SniType=3时，需要填此参数，表示自定义的host；
	SniHost *string `json:"SniHost,omitempty" name:"SniHost"`

	// IsCdn=3时，需要填此参数，表示自定义header
	IpHeaders []*string `json:"IpHeaders,omitempty" name:"IpHeaders"`

	// 0:关闭xff重置；1:开启xff重置
	XFFReset *int64 `json:"XFFReset,omitempty" name:"XFFReset"`
}

type ModifySpartaProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 证书类型，0表示没有证书，CertType=1表示自有证书,2 为托管证书
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// CertType=1时，需要填次参数，表示证书内容
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// CertType=1时，需要填次参数，表示证书的私钥
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// CertType=2时，需要填次参数，表示证书的ID
	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`

	// 表示是否开启了CDN代理
	IsCdn *int64 `json:"IsCdn,omitempty" name:"IsCdn"`

	// HTTPS回源协议
	UpstreamScheme *string `json:"UpstreamScheme,omitempty" name:"UpstreamScheme"`

	// HTTPS回源端口,仅UpstreamScheme为http时需要填当前字段
	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitempty" name:"HttpsUpstreamPort"`

	// 表示是否强制跳转到HTTPS，1表示开启，0表示不开启
	HttpsRewrite *uint64 `json:"HttpsRewrite,omitempty" name:"HttpsRewrite"`

	// 回源类型，0表示通过IP回源,1 表示通过域名回源
	UpstreamType *int64 `json:"UpstreamType,omitempty" name:"UpstreamType"`

	// UpstreamType=1时，填次字段表示回源域名
	UpstreamDomain *string `json:"UpstreamDomain,omitempty" name:"UpstreamDomain"`

	// UpstreamType=0时，填次字段表示回源ip
	SrcList []*string `json:"SrcList,omitempty" name:"SrcList"`

	// 是否开启HTTP2，1表示开启，0表示不开启http2。开启HTTP2需要HTTPS支持
	IsHttp2 *int64 `json:"IsHttp2,omitempty" name:"IsHttp2"`

	// 是否开启WebSocket， 1：开启WebSocket，0：不开启WebSocket
	IsWebsocket *int64 `json:"IsWebsocket,omitempty" name:"IsWebsocket"`

	// 负载均衡策略，0表示轮徇，1表示IP hash
	LoadBalance *int64 `json:"LoadBalance,omitempty" name:"LoadBalance"`

	// 是否灰度
	IsGray *int64 `json:"IsGray,omitempty" name:"IsGray"`

	// WAF版本
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 端口信息
	Ports []*SpartaProtectionPort `json:"Ports,omitempty" name:"Ports"`

	// 长短连接标志，仅IP回源时有效
	IsKeepAlive *string `json:"IsKeepAlive,omitempty" name:"IsKeepAlive"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 是否为Anycast ip类型：1 Anycast 0 普通ip
	Anycast *int64 `json:"Anycast,omitempty" name:"Anycast"`

	// src的权重
	Weights []*int64 `json:"Weights,omitempty" name:"Weights"`

	// 是否开启源站的主动健康检测，1表示开启，0表示不开启
	ActiveCheck *int64 `json:"ActiveCheck,omitempty" name:"ActiveCheck"`

	// TLS版本信息
	TLSVersion *int64 `json:"TLSVersion,omitempty" name:"TLSVersion"`

	// 加密套件信息
	Ciphers []*int64 `json:"Ciphers,omitempty" name:"Ciphers"`

	// 0:不支持选择：默认模板  1:通用型模板 2:安全型模板 3:自定义模板
	CipherTemplate *int64 `json:"CipherTemplate,omitempty" name:"CipherTemplate"`

	// 300s
	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitempty" name:"ProxyReadTimeout"`

	// 300s
	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitempty" name:"ProxySendTimeout"`

	// 0:关闭SNI；1:开启SNI，SNI=源请求host；2:开启SNI，SNI=修改为源站host；3：开启SNI，自定义host，SNI=SniHost；
	SniType *int64 `json:"SniType,omitempty" name:"SniType"`

	// SniType=3时，需要填此参数，表示自定义的host；
	SniHost *string `json:"SniHost,omitempty" name:"SniHost"`

	// IsCdn=3时，需要填此参数，表示自定义header
	IpHeaders []*string `json:"IpHeaders,omitempty" name:"IpHeaders"`

	// 0:关闭xff重置；1:开启xff重置
	XFFReset *int64 `json:"XFFReset,omitempty" name:"XFFReset"`
}

func (r *ModifySpartaProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySpartaProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "CertType")
	delete(f, "Cert")
	delete(f, "PrivateKey")
	delete(f, "SSLId")
	delete(f, "IsCdn")
	delete(f, "UpstreamScheme")
	delete(f, "HttpsUpstreamPort")
	delete(f, "HttpsRewrite")
	delete(f, "UpstreamType")
	delete(f, "UpstreamDomain")
	delete(f, "SrcList")
	delete(f, "IsHttp2")
	delete(f, "IsWebsocket")
	delete(f, "LoadBalance")
	delete(f, "IsGray")
	delete(f, "Edition")
	delete(f, "Ports")
	delete(f, "IsKeepAlive")
	delete(f, "InstanceID")
	delete(f, "Anycast")
	delete(f, "Weights")
	delete(f, "ActiveCheck")
	delete(f, "TLSVersion")
	delete(f, "Ciphers")
	delete(f, "CipherTemplate")
	delete(f, "ProxyReadTimeout")
	delete(f, "ProxySendTimeout")
	delete(f, "SniType")
	delete(f, "SniHost")
	delete(f, "IpHeaders")
	delete(f, "XFFReset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySpartaProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySpartaProtectionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySpartaProtectionResponse struct {
	*tchttp.BaseResponse
	Response *ModifySpartaProtectionResponseParams `json:"Response"`
}

func (r *ModifySpartaProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySpartaProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWafAutoDenyRulesRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 触发IP封禁的攻击次数阈值，范围为2~100次
	AttackThreshold *int64 `json:"AttackThreshold,omitempty" name:"AttackThreshold"`

	// IP封禁统计时间，范围为1-60分钟
	TimeThreshold *int64 `json:"TimeThreshold,omitempty" name:"TimeThreshold"`

	// 触发IP封禁后的封禁时间，范围为5~360分钟
	DenyTimeThreshold *int64 `json:"DenyTimeThreshold,omitempty" name:"DenyTimeThreshold"`

	// 自动封禁状态，0表示关闭，1表示打开
	DefenseStatus *int64 `json:"DefenseStatus,omitempty" name:"DefenseStatus"`
}

type ModifyWafAutoDenyRulesRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 触发IP封禁的攻击次数阈值，范围为2~100次
	AttackThreshold *int64 `json:"AttackThreshold,omitempty" name:"AttackThreshold"`

	// IP封禁统计时间，范围为1-60分钟
	TimeThreshold *int64 `json:"TimeThreshold,omitempty" name:"TimeThreshold"`

	// 触发IP封禁后的封禁时间，范围为5~360分钟
	DenyTimeThreshold *int64 `json:"DenyTimeThreshold,omitempty" name:"DenyTimeThreshold"`

	// 自动封禁状态，0表示关闭，1表示打开
	DefenseStatus *int64 `json:"DefenseStatus,omitempty" name:"DefenseStatus"`
}

func (r *ModifyWafAutoDenyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWafAutoDenyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "AttackThreshold")
	delete(f, "TimeThreshold")
	delete(f, "DenyTimeThreshold")
	delete(f, "DefenseStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWafAutoDenyRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWafAutoDenyRulesResponseParams struct {
	// 成功的状态码，需要JSON解码后再使用，返回的格式是{"域名":"状态"}，成功的状态码为Success，其它的为失败的状态码（yunapi定义的错误码）
	Success *ResponseCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyWafAutoDenyRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWafAutoDenyRulesResponseParams `json:"Response"`
}

func (r *ModifyWafAutoDenyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWafAutoDenyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWafAutoDenyStatusRequestParams struct {
	// WAF 自动封禁配置项
	WafAutoDenyDetails *AutoDenyDetail `json:"WafAutoDenyDetails,omitempty" name:"WafAutoDenyDetails"`
}

type ModifyWafAutoDenyStatusRequest struct {
	*tchttp.BaseRequest
	
	// WAF 自动封禁配置项
	WafAutoDenyDetails *AutoDenyDetail `json:"WafAutoDenyDetails,omitempty" name:"WafAutoDenyDetails"`
}

func (r *ModifyWafAutoDenyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWafAutoDenyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WafAutoDenyDetails")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWafAutoDenyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWafAutoDenyStatusResponseParams struct {
	// WAF 自动封禁配置项
	WafAutoDenyDetails *AutoDenyDetail `json:"WafAutoDenyDetails,omitempty" name:"WafAutoDenyDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyWafAutoDenyStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWafAutoDenyStatusResponseParams `json:"Response"`
}

func (r *ModifyWafAutoDenyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWafAutoDenyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWafThreatenIntelligenceRequestParams struct {
	// 配置WAF威胁情报封禁模块详情
	WafThreatenIntelligenceDetails *WafThreatenIntelligenceDetails `json:"WafThreatenIntelligenceDetails,omitempty" name:"WafThreatenIntelligenceDetails"`
}

type ModifyWafThreatenIntelligenceRequest struct {
	*tchttp.BaseRequest
	
	// 配置WAF威胁情报封禁模块详情
	WafThreatenIntelligenceDetails *WafThreatenIntelligenceDetails `json:"WafThreatenIntelligenceDetails,omitempty" name:"WafThreatenIntelligenceDetails"`
}

func (r *ModifyWafThreatenIntelligenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWafThreatenIntelligenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WafThreatenIntelligenceDetails")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWafThreatenIntelligenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWafThreatenIntelligenceResponseParams struct {
	// 当前WAF威胁情报封禁模块详情
	WafThreatenIntelligenceDetails *WafThreatenIntelligenceDetails `json:"WafThreatenIntelligenceDetails,omitempty" name:"WafThreatenIntelligenceDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyWafThreatenIntelligenceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWafThreatenIntelligenceResponseParams `json:"Response"`
}

func (r *ModifyWafThreatenIntelligenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWafThreatenIntelligenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PeakPointsItem struct {
	// 秒级别时间戳
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// QPS
	Access *uint64 `json:"Access,omitempty" name:"Access"`

	// 上行带宽峰值，单位B
	Up *uint64 `json:"Up,omitempty" name:"Up"`

	// 下行带宽峰值，单位B
	Down *uint64 `json:"Down,omitempty" name:"Down"`

	// Web攻击次数
	Attack *uint64 `json:"Attack,omitempty" name:"Attack"`

	// CC攻击次数
	Cc *uint64 `json:"Cc,omitempty" name:"Cc"`

	// Bot qps
	BotAccess *uint64 `json:"BotAccess,omitempty" name:"BotAccess"`

	// WAF返回给客户端状态码次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusServerError *uint64 `json:"StatusServerError,omitempty" name:"StatusServerError"`

	// WAF返回给客户端状态码次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusClientError *uint64 `json:"StatusClientError,omitempty" name:"StatusClientError"`

	// WAF返回给客户端状态码次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusRedirect *uint64 `json:"StatusRedirect,omitempty" name:"StatusRedirect"`

	// WAF返回给客户端状态码次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusOk *uint64 `json:"StatusOk,omitempty" name:"StatusOk"`

	// 源站返回给WAF状态码次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamServerError *uint64 `json:"UpstreamServerError,omitempty" name:"UpstreamServerError"`

	// 源站返回给WAF状态码次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamClientError *uint64 `json:"UpstreamClientError,omitempty" name:"UpstreamClientError"`

	// 源站返回给WAF状态码次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamRedirect *uint64 `json:"UpstreamRedirect,omitempty" name:"UpstreamRedirect"`
}

type PortInfo struct {
	// Nginx的服务器id
	NginxServerId *uint64 `json:"NginxServerId,omitempty" name:"NginxServerId"`

	// 监听端口配置
	Port *string `json:"Port,omitempty" name:"Port"`

	// 与端口对应的协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 回源端口
	UpstreamPort *string `json:"UpstreamPort,omitempty" name:"UpstreamPort"`

	// 回源协议
	UpstreamProtocol *string `json:"UpstreamProtocol,omitempty" name:"UpstreamProtocol"`
}

type PortItem struct {
	// 监听端口配置
	Port *string `json:"Port,omitempty" name:"Port"`

	// 与Port一一对应，表示端口对应的协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 与Port一一对应,  表示回源端口
	UpstreamPort *string `json:"UpstreamPort,omitempty" name:"UpstreamPort"`

	// 与Port一一对应,  表示回源协议
	UpstreamProtocol *string `json:"UpstreamProtocol,omitempty" name:"UpstreamProtocol"`

	// Nginx的服务器ID
	NginxServerId *string `json:"NginxServerId,omitempty" name:"NginxServerId"`
}

// Predefined struct for user
type PostAttackDownloadTaskRequestParams struct {
	// 查询的域名，所有域名使用all
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Lucene语法
	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 默认为desc，可以取值desc和asc
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 下载的日志条数
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type PostAttackDownloadTaskRequest struct {
	*tchttp.BaseRequest
	
	// 查询的域名，所有域名使用all
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Lucene语法
	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 默认为desc，可以取值desc和asc
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 下载的日志条数
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

func (r *PostAttackDownloadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PostAttackDownloadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QueryString")
	delete(f, "TaskName")
	delete(f, "Sort")
	delete(f, "Count")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PostAttackDownloadTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PostAttackDownloadTaskResponseParams struct {
	// 任务task id
	Flow *string `json:"Flow,omitempty" name:"Flow"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PostAttackDownloadTaskResponse struct {
	*tchttp.BaseResponse
	Response *PostAttackDownloadTaskResponseParams `json:"Response"`
}

func (r *PostAttackDownloadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PostAttackDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QPSPackageNew struct {
	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`

	// 是否自动续费，1：自动续费，0：不自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 套餐购买个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 套餐购买地域，clb-waf暂时没有用到
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 计费项
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingItem *string `json:"BillingItem,omitempty" name:"BillingItem"`
}

// Predefined struct for user
type RefreshAccessCheckResultRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type RefreshAccessCheckResultRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *RefreshAccessCheckResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshAccessCheckResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefreshAccessCheckResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefreshAccessCheckResultResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RefreshAccessCheckResultResponse struct {
	*tchttp.BaseResponse
	Response *RefreshAccessCheckResultResponseParams `json:"Response"`
}

func (r *RefreshAccessCheckResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshAccessCheckResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResponseCode struct {
	// 如果成功则返回Success，失败则返回云api定义的错误码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 如果成功则返回Success，失败则返回WAF定义的二级错误码
	Message *string `json:"Message,omitempty" name:"Message"`
}

type RuleList struct {
	// 规则Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 规则列表的id
	Rules []*uint64 `json:"Rules,omitempty" name:"Rules"`

	// 请求url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 请求的方法
	Function *string `json:"Function,omitempty" name:"Function"`

	// 时间戳
	Time *string `json:"Time,omitempty" name:"Time"`

	// 开关状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type SearchAccessLogRequestParams struct {
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题，新版本此字段填空字符串
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitempty" name:"To"`

	// 查询语句，语句长度最大为4096
	Query *string `json:"Query,omitempty" name:"Query"`

	// 单次查询返回的日志条数，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 新版本此字段失效，填空字符串，翻页使用Page
	Context *string `json:"Context,omitempty" name:"Context"`

	// 日志接口是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 第几页，从0开始。新版本接口字段
	Page *int64 `json:"Page,omitempty" name:"Page"`
}

type SearchAccessLogRequest struct {
	*tchttp.BaseRequest
	
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题，新版本此字段填空字符串
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitempty" name:"To"`

	// 查询语句，语句长度最大为4096
	Query *string `json:"Query,omitempty" name:"Query"`

	// 单次查询返回的日志条数，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 新版本此字段失效，填空字符串，翻页使用Page
	Context *string `json:"Context,omitempty" name:"Context"`

	// 日志接口是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 第几页，从0开始。新版本接口字段
	Page *int64 `json:"Page,omitempty" name:"Page"`
}

func (r *SearchAccessLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchAccessLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Sort")
	delete(f, "Page")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchAccessLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchAccessLogResponseParams struct {
	// 新接口此字段失效，默认返回空字符串
	Context *string `json:"Context,omitempty" name:"Context"`

	// 日志查询结果是否全部返回，其中，“true”表示结果返回，“false”表示结果为返回
	ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

	// 返回的是否为分析结果，其中，“true”表示返回分析结果，“false”表示未返回分析结果
	Analysis *bool `json:"Analysis,omitempty" name:"Analysis"`

	// 如果Analysis为True，则返回分析结果的列名，否则为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColNames []*string `json:"ColNames,omitempty" name:"ColNames"`

	// 日志查询结果；当Analysis为True时，可能返回为null
	// 注意：此字段可能返回 null，表示取不到有效值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*AccessLogInfo `json:"Results,omitempty" name:"Results"`

	// 日志分析结果；当Analysis为False时，可能返回为null
	// 注意：此字段可能返回 null，表示取不到有效值
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnalysisResults []*AccessLogItems `json:"AnalysisResults,omitempty" name:"AnalysisResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchAccessLogResponse struct {
	*tchttp.BaseResponse
	Response *SearchAccessLogResponseParams `json:"Response"`
}

func (r *SearchAccessLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchAccessLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchAttackLogRequestParams struct {
	// 查询的域名，所有域名使用all
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 接口升级，这个字段传空字符串,翻页使用Page字段
	Context *string `json:"Context,omitempty" name:"Context"`

	// Lucene语法
	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`

	// 查询的数量，默认10条，最多100条
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 默认为desc，可以取值desc和asc
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 第几页，从0开始
	Page *int64 `json:"Page,omitempty" name:"Page"`
}

type SearchAttackLogRequest struct {
	*tchttp.BaseRequest
	
	// 查询的域名，所有域名使用all
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 接口升级，这个字段传空字符串,翻页使用Page字段
	Context *string `json:"Context,omitempty" name:"Context"`

	// Lucene语法
	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`

	// 查询的数量，默认10条，最多100条
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 默认为desc，可以取值desc和asc
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 第几页，从0开始
	Page *int64 `json:"Page,omitempty" name:"Page"`
}

func (r *SearchAttackLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchAttackLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Context")
	delete(f, "QueryString")
	delete(f, "Count")
	delete(f, "Sort")
	delete(f, "Page")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchAttackLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchAttackLogResponseParams struct {
	// 当前返回的攻击日志条数
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 接口升级，此字段无效，默认返回空字符串
	Context *string `json:"Context,omitempty" name:"Context"`

	// 攻击日志数组条目内容
	Data []*AttackLogInfo `json:"Data,omitempty" name:"Data"`

	// CLS接口返回内容
	ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

	// CLS接口返回内容，标志是否启动新版本索引
	SqlFlag *bool `json:"SqlFlag,omitempty" name:"SqlFlag"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchAttackLogResponse struct {
	*tchttp.BaseResponse
	Response *SearchAttackLogResponseParams `json:"Response"`
}

func (r *SearchAttackLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchAttackLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchItem struct {
	// 日志开关
	ClsStatus *string `json:"ClsStatus,omitempty" name:"ClsStatus"`

	// waf开关
	Status *string `json:"Status,omitempty" name:"Status"`

	// 流量模式
	FlowMode *string `json:"FlowMode,omitempty" name:"FlowMode"`
}

type SpartaProtectionPort struct {
	// nginx Id
	NginxServerId *uint64 `json:"NginxServerId,omitempty" name:"NginxServerId"`

	// 端口
	Port *string `json:"Port,omitempty" name:"Port"`

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 后端端口
	UpstreamPort *string `json:"UpstreamPort,omitempty" name:"UpstreamPort"`

	// 后端协议
	UpstreamProtocol *string `json:"UpstreamProtocol,omitempty" name:"UpstreamProtocol"`
}

type Strategy struct {
	// 匹配字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Field *string `json:"Field,omitempty" name:"Field"`

	// 逻辑符号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareFunc *string `json:"CompareFunc,omitempty" name:"CompareFunc"`

	// 匹配内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 匹配参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Arg *string `json:"Arg,omitempty" name:"Arg"`
}

// Predefined struct for user
type SwitchDomainRulesRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 规则列表
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`

	// 开关状态，0表示关闭，1表示开启，2表示只观察
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 设置为观察模式原因，
	// 1表示业务自身原因观察，2表示系统规则误报上报，3表示核心业务灰度观察，4表示其他
	Reason *uint64 `json:"Reason,omitempty" name:"Reason"`
}

type SwitchDomainRulesRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 规则列表
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`

	// 开关状态，0表示关闭，1表示开启，2表示只观察
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 设置为观察模式原因，
	// 1表示业务自身原因观察，2表示系统规则误报上报，3表示核心业务灰度观察，4表示其他
	Reason *uint64 `json:"Reason,omitempty" name:"Reason"`
}

func (r *SwitchDomainRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDomainRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Ids")
	delete(f, "Status")
	delete(f, "Reason")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchDomainRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDomainRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwitchDomainRulesResponse struct {
	*tchttp.BaseResponse
	Response *SwitchDomainRulesResponseParams `json:"Response"`
}

func (r *SwitchDomainRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDomainRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TLSCiphers struct {
	// TLS版本ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`

	// 加密套件ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CipherId *int64 `json:"CipherId,omitempty" name:"CipherId"`

	// 加密套件
	// 注意：此字段可能返回 null，表示取不到有效值。
	CipherName *string `json:"CipherName,omitempty" name:"CipherName"`
}

type TLSVersion struct {
	// TLSVERSION的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`

	// TLSVERSION的NAME
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
}

// Predefined struct for user
type UpsertIpAccessControlRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// ip 参数列表，json数组由ip，source，note，action，valid_ts组成。ip对应配置的ip地址，source固定为custom值，note为注释，action值42为黑名单，40为白名单，valid_ts为有效日期，值为秒级时间戳（（如1680570420代表2023-04-04 09:07:00））
	Items []*string `json:"Items,omitempty" name:"Items"`

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 是否为多域名黑白名单，当为多域名的黑白名单时，取值为batch，否则为空
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
}

type UpsertIpAccessControlRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// ip 参数列表，json数组由ip，source，note，action，valid_ts组成。ip对应配置的ip地址，source固定为custom值，note为注释，action值42为黑名单，40为白名单，valid_ts为有效日期，值为秒级时间戳（（如1680570420代表2023-04-04 09:07:00））
	Items []*string `json:"Items,omitempty" name:"Items"`

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 是否为多域名黑白名单，当为多域名的黑白名单时，取值为batch，否则为空
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
}

func (r *UpsertIpAccessControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpsertIpAccessControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Items")
	delete(f, "Edition")
	delete(f, "SourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpsertIpAccessControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpsertIpAccessControlResponseParams struct {
	// 添加或修改失败的条目
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedItems *string `json:"FailedItems,omitempty" name:"FailedItems"`

	// 添加或修改失败的数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpsertIpAccessControlResponse struct {
	*tchttp.BaseResponse
	Response *UpsertIpAccessControlResponseParams `json:"Response"`
}

func (r *UpsertIpAccessControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpsertIpAccessControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserDomainInfo struct {
	// 用户id
	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名id
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// waf类型
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *string `json:"Level,omitempty" name:"Level"`

	// 指定域名访问日志字段的开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	WriteConfig *string `json:"WriteConfig,omitempty" name:"WriteConfig"`

	// 指定域名是否写cls的开关 1:写 0:不写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cls *uint64 `json:"Cls,omitempty" name:"Cls"`
}

type VipInfo struct {
	// Virtual IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// waf实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type WafRuleLimit struct {
	// 自定义CC的规格
	CC *uint64 `json:"CC,omitempty" name:"CC"`

	// 自定义策略的规格
	CustomRule *uint64 `json:"CustomRule,omitempty" name:"CustomRule"`

	// 黑白名单的规格
	IPControl *uint64 `json:"IPControl,omitempty" name:"IPControl"`

	// 信息防泄漏的规格
	AntiLeak *uint64 `json:"AntiLeak,omitempty" name:"AntiLeak"`

	// 防篡改的规格
	AntiTamper *uint64 `json:"AntiTamper,omitempty" name:"AntiTamper"`

	// 紧急CC的规格
	AutoCC *uint64 `json:"AutoCC,omitempty" name:"AutoCC"`

	// 地域封禁的规格
	AreaBan *uint64 `json:"AreaBan,omitempty" name:"AreaBan"`

	// 自定义CC中配置session
	CCSession *uint64 `json:"CCSession,omitempty" name:"CCSession"`

	// AI的规格
	AI *uint64 `json:"AI,omitempty" name:"AI"`

	// 精准白名单的规格
	CustomWhite *uint64 `json:"CustomWhite,omitempty" name:"CustomWhite"`

	// api安全的规格
	ApiSecurity *uint64 `json:"ApiSecurity,omitempty" name:"ApiSecurity"`

	// 客户端流量标记的规格
	ClientMsg *uint64 `json:"ClientMsg,omitempty" name:"ClientMsg"`

	// 流量标记的规格
	TrafficMarking *uint64 `json:"TrafficMarking,omitempty" name:"TrafficMarking"`
}

type WafThreatenIntelligenceDetails struct {
	// 封禁属性标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 封禁模组启用状态
	DefenseStatus *int64 `json:"DefenseStatus,omitempty" name:"DefenseStatus"`

	// 最后更新时间
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}