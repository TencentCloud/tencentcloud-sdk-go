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

type AddCustomRuleRequest struct {
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

	// 动作类型
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 如果动作是重定向，则表示重定向的地址；其他情况可以为空
	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`

	// "clb-waf"或者"sparta-waf"
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 放行的详情
	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCustomRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddCustomRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
		Success *ResponseCode `json:"Success,omitempty" name:"Success"`

		// 添加成功的规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// 规则的开关
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

type AddDomainWhiteRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则id
		Id *uint64 `json:"Id,omitempty" name:"Id"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// 负载均衡策略，0表示轮徇，1表示IP hash
	LoadBalance *string `json:"LoadBalance,omitempty" name:"LoadBalance"`

	// CertType=1时，需要填次参数，表示证书内容
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

	// 版本：sparta-waf、clb-waf、cdn-waf
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 是否开启长连接，仅IP回源时可以用填次参数，域名回源时这个参数无效
	IsKeepAlive *string `json:"IsKeepAlive,omitempty" name:"IsKeepAlive"`

	// 实例id，上线之后带上此字段
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// anycast IP类型开关： 0 普通IP 1 Anycast IP
	Anycast *int64 `json:"Anycast,omitempty" name:"Anycast"`

	// src权重
	Weights []*int64 `json:"Weights,omitempty" name:"Weights"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddSpartaProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddSpartaProtectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateAccessExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志导出ID。
		ExportId *string `json:"ExportId,omitempty" name:"ExportId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateAttackDownloadTaskRequest struct {
	*tchttp.BaseRequest

	// 域名，所有域名填写all
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 查询起始时间
	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`

	// 下载任务名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 风险等级
	RiskLevel *uint64 `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 拦截状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 自定义策略ID
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// 攻击者IP
	AttackIp *string `json:"AttackIp,omitempty" name:"AttackIp"`

	// 攻击类型
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`
}

func (r *CreateAttackDownloadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAttackDownloadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "FromTime")
	delete(f, "ToTime")
	delete(f, "Name")
	delete(f, "RiskLevel")
	delete(f, "Status")
	delete(f, "RuleId")
	delete(f, "AttackIp")
	delete(f, "AttackType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAttackDownloadTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAttackDownloadTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		Flow *string `json:"Flow,omitempty" name:"Flow"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAttackDownloadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAttackDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type DeleteAccessExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteAttackDownloadRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteDomainWhiteRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 出参
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteDownloadRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteIpAccessControlRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 删除的ip数组
	Items []*string `json:"Items,omitempty" name:"Items"`

	// 删除对应的域名下的所有黑/白IP名额单
	DeleteAll *bool `json:"DeleteAll,omitempty" name:"DeleteAll"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIpAccessControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIpAccessControlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除失败的条目
	// 注意：此字段可能返回 null，表示取不到有效值。
		FailedItems *string `json:"FailedItems,omitempty" name:"FailedItems"`

		// 删除失败的条目数
	// 注意：此字段可能返回 null，表示取不到有效值。
		FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAccessExportsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志导出ID。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 日志导出列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Exports []*ExportAccessInfo `json:"Exports,omitempty" name:"Exports"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAccessFastAnalysisResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAccessIndexResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否生效
		Status *bool `json:"Status,omitempty" name:"Status"`

		// 索引配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Rule *AccessRuleInfo `json:"Rule,omitempty" name:"Rule"`

		// 索引修改时间，初始值为索引创建时间。
		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAutoDenyIPResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询IP封禁状态返回结果
		Data *IpHitItemsData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeCustomRulesPagingInfo struct {

	// 当前页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 当前页的最大数据条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeCustomRulesRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 分页参数
	Paging *DescribeCustomRulesPagingInfo `json:"Paging,omitempty" name:"Paging"`

	// clb-waf或者sparta-waf
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 过滤参数：动作类型：0放行，1阻断，2人机识别，3观察，4重定向
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 过滤参数：规则名称过滤条件
	Search *string `json:"Search,omitempty" name:"Search"`
}

func (r *DescribeCustomRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Paging")
	delete(f, "Edition")
	delete(f, "ActionType")
	delete(f, "Search")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则详情
		RuleList []*DescribeCustomRulesRspRuleListItem `json:"RuleList,omitempty" name:"RuleList"`

		// 规则条数
		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomRulesResponse) FromJsonString(s string) error {
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

	// 排序方式
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

type DescribeDomainWhiteRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则列表
		RuleList []*RuleList `json:"RuleList,omitempty" name:"RuleList"`

		// 规则的数量
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeDomainsRequest struct {
	*tchttp.BaseRequest

	// 偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 容量
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

type DescribeDomainsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// domain列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Domains []*DomainInfo `json:"Domains,omitempty" name:"Domains"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeFlowTrendResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 流量趋势数据
		Data []*BotStatPointItem `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeIpAccessControlRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 计数标识
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 动作
	ActionType *uint64 `json:"ActionType,omitempty" name:"ActionType"`

	// 有效时间最小时间戳
	VtsMin *uint64 `json:"VtsMin,omitempty" name:"VtsMin"`

	// 有效时间最大时间戳
	VtsMax *uint64 `json:"VtsMax,omitempty" name:"VtsMax"`

	// 创建时间最小时间戳
	CtsMin *uint64 `json:"CtsMin,omitempty" name:"CtsMin"`

	// 创建时间最大时间戳
	CtsMax *uint64 `json:"CtsMax,omitempty" name:"CtsMax"`

	// 偏移
	OffSet *uint64 `json:"OffSet,omitempty" name:"OffSet"`

	// 限制
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

type DescribeIpAccessControlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 输出
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *IpAccessControlData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeIpHitItemsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *IpHitItemsData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeUserClbWafRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 地域（标准的ap-格式）列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeWafAutoDenyRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeWafAutoDenyStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// WAF 自动封禁详情
		WafAutoDenyDetails *AutoDenyDetail `json:"WafAutoDenyDetails,omitempty" name:"WafAutoDenyDetails"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeWafThreatenIntelligenceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// WAF 威胁情报封禁信息
		WafThreatenIntelligenceDetails *WafThreatenIntelligenceDetails `json:"WafThreatenIntelligenceDetails,omitempty" name:"WafThreatenIntelligenceDetails"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type FiltersItemNew struct {

	// 字段名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤值
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否精确查找
	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
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

type ModifyAccessPeriodResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyCustomRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
		Success *ResponseCode `json:"Success,omitempty" name:"Success"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// 规则的开关状态
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

type ModifyDomainWhiteRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyWafAutoDenyRulesRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 攻击次数阈值
	AttackThreshold *int64 `json:"AttackThreshold,omitempty" name:"AttackThreshold"`

	// 攻击时间阈值
	TimeThreshold *int64 `json:"TimeThreshold,omitempty" name:"TimeThreshold"`

	// 自动封禁时间
	DenyTimeThreshold *int64 `json:"DenyTimeThreshold,omitempty" name:"DenyTimeThreshold"`

	// 自动封禁状态
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

type ModifyWafAutoDenyRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功的状态码，需要JSON解码后再使用，返回的格式是{"域名":"状态"}，成功的状态码为Success，其它的为失败的状态码（yunapi定义的错误码）
		Success *ResponseCode `json:"Success,omitempty" name:"Success"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyWafAutoDenyStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// WAF 自动封禁配置项
		WafAutoDenyDetails *AutoDenyDetail `json:"WafAutoDenyDetails,omitempty" name:"WafAutoDenyDetails"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyWafThreatenIntelligenceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 当前WAF威胁情报封禁模块详情
		WafThreatenIntelligenceDetails *WafThreatenIntelligenceDetails `json:"WafThreatenIntelligenceDetails,omitempty" name:"WafThreatenIntelligenceDetails"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ResponseCode struct {

	// 如果成功则返回Success，失败则返回yunapi定义的错误码
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

type SearchAccessLogRequest struct {
	*tchttp.BaseRequest

	// 客户要查询的日志主题ID，每个客户都有对应的一个主题
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitempty" name:"To"`

	// 查询语句，语句长度最大为4096
	Query *string `json:"Query,omitempty" name:"Query"`

	// 单次查询返回的日志条数，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 加载更多日志时使用，透传上次返回的Context值，获取后续的日志内容
	Context *string `json:"Context,omitempty" name:"Context"`

	// 日志接口是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc
	Sort *string `json:"Sort,omitempty" name:"Sort"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchAccessLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SearchAccessLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 加载后续内容的Context
		Context *string `json:"Context,omitempty" name:"Context"`

		// 日志查询结果是否全部返回
		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

		// 返回的是否为分析结果
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
	} `json:"Response"`
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

type Strategy struct {

	// 匹配字段
	Field *string `json:"Field,omitempty" name:"Field"`

	// 逻辑符号
	CompareFunc *string `json:"CompareFunc,omitempty" name:"CompareFunc"`

	// 匹配内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 匹配参数
	Arg *string `json:"Arg,omitempty" name:"Arg"`
}

type UpsertIpAccessControlRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// ip 参数列表，json数组由ip，source，note，action，valid_ts组成。ip对应配置的ip地址，source固定为custom值，note为注释，action值42为黑名单，40为白名单，valid_ts为有效日期，值为秒级时间戳
	Items []*string `json:"Items,omitempty" name:"Items"`

	// clb-waf或者sparta-waf
	Edition *string `json:"Edition,omitempty" name:"Edition"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpsertIpAccessControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpsertIpAccessControlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 添加或修改失败的条目
	// 注意：此字段可能返回 null，表示取不到有效值。
		FailedItems *string `json:"FailedItems,omitempty" name:"FailedItems"`

		// 添加或修改失败的数目
	// 注意：此字段可能返回 null，表示取不到有效值。
		FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type WafThreatenIntelligenceDetails struct {

	// 封禁模组启用状态
	DefenseStatus *int64 `json:"DefenseStatus,omitempty" name:"DefenseStatus"`

	// 封禁属性标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 最后更新时间
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}
