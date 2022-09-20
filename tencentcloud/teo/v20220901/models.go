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

package v20220901

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccelerateType struct {
	// 加速开关。取值范围：
	// <li> on：打开;</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type AclCondition struct {
	// 匹配字段，取值有：
	// <li>host：请求域名；</li>
	// <li>sip：客户端IP；</li>
	// <li>ua：User-Agent；</li>
	// <li>cookie：会话 Cookie；</li>
	// <li>cgi：CGI 脚本；</li>
	// <li>xff：XFF 扩展头部；</li>
	// <li>url：请求 URL；</li>
	// <li>accept：请求内容类型；</li>
	// <li>method：请求方式；</li>
	// <li>header：请求头部；</li>
	// <li>sip_proto：网络层协议。</li>
	MatchFrom *string `json:"MatchFrom,omitempty" name:"MatchFrom"`

	// 匹配字符串。当 MatchFrom 为 header 时，可以填入 header 的 key 作为参数。
	MatchParam *string `json:"MatchParam,omitempty" name:"MatchParam"`

	// 匹配关系，取值有：
	// <li>equal：字符串等于；</li>
	// <li>not_equal：数值不等于；</li>
	// <li>include：字符包含；</li>
	// <li>not_include：字符不包含；</li>
	// <li>match：ip匹配；</li>
	// <li>not_match：ip不匹配；</li>
	// <li>include_area：地域包含；</li>
	// <li>is_empty：存在字段但值为空；</li>
	// <li>not_exists：不存在关键字段；</li>
	// <li>regexp：正则匹配；</li>
	// <li>len_gt：数值大于；</li>
	// <li>len_lt：数值小于；</li>
	// <li>len_eq：数值等于；</li>
	// <li>match_prefix：前缀匹配；</li>
	// <li>match_suffix：后缀匹配；</li>
	// <li>wildcard：通配符。</li>
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 匹配内容。
	MatchContent *string `json:"MatchContent,omitempty" name:"MatchContent"`
}

type AclConfig struct {
	// 开关，取值有：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 用户自定义规则。
	AclUserRules []*AclUserRule `json:"AclUserRules,omitempty" name:"AclUserRules"`
}

type AclUserRule struct {
	// 规则名。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 处罚动作，取值有：
	// <li>trans：放行；</li>
	// <li>drop：拦截；</li>
	// <li>monitor：观察；</li>
	// <li>ban：IP封禁；</li>
	// <li>redirect：重定向；</li>
	// <li>page：指定页面；</li>
	// <li>alg：Javascript挑战。</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// 规则状态，取值有：
	// <li>on：生效；</li>
	// <li>off：失效。</li>
	RuleStatus *string `json:"RuleStatus,omitempty" name:"RuleStatus"`

	// 自定义规则。
	AclConditions []*AclCondition `json:"AclConditions,omitempty" name:"AclConditions"`

	// 规则优先级，取值范围0-100。
	RulePriority *int64 `json:"RulePriority,omitempty" name:"RulePriority"`

	// 规则Id。仅出参使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// 更新时间。仅出参使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// ip封禁的惩罚时间，取值范围0-2天。默认为0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PunishTime *int64 `json:"PunishTime,omitempty" name:"PunishTime"`

	// ip封禁的惩罚时间单位，取值有：
	// <li>second：秒；</li>
	// <li>minutes：分；</li>
	// <li>hour：小时。</li>默认为second。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PunishTimeUnit *string `json:"PunishTimeUnit,omitempty" name:"PunishTimeUnit"`

	// 自定义返回页面的名称。默认为空字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 自定义返回页面的实例id。默认为0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageId *int64 `json:"PageId,omitempty" name:"PageId"`

	// 重定向时候的地址，必须为本用户接入的站点子域名。默认为空字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`

	// 重定向时候的返回码。默认为0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseCode *int64 `json:"ResponseCode,omitempty" name:"ResponseCode"`
}

type Action struct {
	// 常规功能操作，选择该类型的功能项有：
	// <li> 访问URL 重写（AccessUrlRedirect）；</li>
	// <li> 回源 URL 重写 （UpstreamUrlRedirect）；</li>
	// <li> QUIC（QUIC）；</li>
	// <li> WebSocket （WebSocket）；</li>
	// <li> 视频拖拽（VideoSeek）；</li>
	// <li> Token 鉴权（Authentication）；</li>
	// <li> 自定义CacheKey（CacheKey）；</li>
	// <li> 节点缓存 TTL （Cache）；</li>
	// <li> 浏览器缓存 TTL（MaxAge）；</li>
	// <li> 离线缓存（OfflineCache）；</li>
	// <li> 智能加速（SmartRouting）；</li>
	// <li> 分片回源（RangeOriginPull）；</li>
	// <li> HTTP/2 回源（UpstreamHttp2）；</li>
	// <li> Host Header 重写（HostHeader）；</li>
	// <li> 强制 HTTPS（ForceRedirect）；</li>
	// <li> 回源 HTTPS（OriginPullProtocol）；</li>
	// <li> 缓存预刷新（CachePrefresh）；</li>
	// <li> 智能压缩（Compression）；</li>
	// <li> Hsts；</li>
	// <li> ClientIpHeader；</li>
	// <li> TlsVersion；</li>
	// <li> OcspStapling。</li>
	// <li> HTTP/2 访问（Http2）。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormalAction *NormalAction `json:"NormalAction,omitempty" name:"NormalAction"`

	// 带有请求头/响应头的功能操作，选择该类型的功能项有：
	// <li> 修改 HTTP 请求头（RequestHeader）；</li>
	// <li> 修改HTTP响应头（ResponseHeader）。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RewriteAction *RewriteAction `json:"RewriteAction,omitempty" name:"RewriteAction"`

	// 带有状态码的功能操作，选择该类型的功能项有：
	// <li> 自定义错误页面（ErrorPage）；</li>
	// <li> 状态码缓存 TTL（StatusCodeCache）。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeAction *CodeAction `json:"CodeAction,omitempty" name:"CodeAction"`
}

type AdvancedFilter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否启用模糊查询。
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}

type AiRule struct {
	// AI规则引擎状态，取值有：
	// <li> smart_status_close：关闭；</li>
	// <li> smart_status_open：拦截处置；</li>
	// <li> smart_status_observe：观察处置。</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type ApplicationProxy struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 当ProxyType=hostname时，表示域名或子域名；
	// 当ProxyType=instance时，表示代理名称。
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// 四层代理模式，取值有：
	// <li>hostname：表示子域名模式；</li>
	// <li>instance：表示实例模式。</li>
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// 调度模式，取值有：
	// <li>ip：表示Anycast IP调度；</li>
	// <li>domain：表示CNAME调度。</li>
	PlatType *string `json:"PlatType,omitempty" name:"PlatType"`

	// 加速区域，取值有：
	// <li>mainland：中国大陆境内;</li>
	// <li>overseas：全球（不含中国大陆）。</li>
	// 默认值：overseas
	Area *string `json:"Area,omitempty" name:"Area"`

	// 是否开启安全，取值有：
	// <li>0：关闭安全；</li>
	// <li>1：开启安全。</li>
	SecurityType *int64 `json:"SecurityType,omitempty" name:"SecurityType"`

	// 是否开启加速，取值有：
	// <li>0：关闭加速；</li>
	// <li>1：开启加速。</li>
	AccelerateType *int64 `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// 会话保持时间。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// 状态，取值有：
	// <li>online：启用；</li>
	// <li>offline：停用；</li>
	// <li>progress：部署中；</li>
	// <li>stopping：停用中；</li>
	// <li>fail：部署失败/停用失败。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 封禁状态，取值有：
	// <li>banned：已封禁;</li>
	// <li>banning：封禁中；</li>
	// <li>recover：已解封；</li>
	// <li>recovering：解封禁中。</li>
	BanStatus *string `json:"BanStatus,omitempty" name:"BanStatus"`

	// 调度信息。
	ScheduleValue []*string `json:"ScheduleValue,omitempty" name:"ScheduleValue"`

	// 当ProxyType=hostname时：
	// 表示代理加速唯一标识。
	HostId *string `json:"HostId,omitempty" name:"HostId"`

	// Ipv6访问配置。
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 规则列表。
	ApplicationProxyRules []*ApplicationProxyRule `json:"ApplicationProxyRules,omitempty" name:"ApplicationProxyRules"`
}

type ApplicationProxyRule struct {
	// 协议，取值有：
	// <li>TCP：TCP协议；</li>
	// <li>UDP：UDP协议。</li>
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// 端口，支持格式：
	// 单个端口，如：80。
	// 端口段，如：81-82。表示81，82两个端口。
	// 注意：一条规则最多可填写20个端口。
	Port []*string `json:"Port,omitempty" name:"Port"`

	// 源站类型，取值有：
	// <li>custom：手动添加；</li>
	// <li>origins：源站组。</li>
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 源站信息：
	// 当OriginType=custom时，表示一个或多个源站，如：
	// OriginValue=["8.8.8.8:80","9.9.9.9:80"]
	// OriginValue=["test.com:80"]；
	// 当OriginType=origins时，要求有且仅有一个元素，表示源站组ID，如：
	// OriginValue=["origin-537f5b41-162a-11ed-abaa-525400c5da15"]。
	OriginValue []*string `json:"OriginValue,omitempty" name:"OriginValue"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 状态，取值有：
	// <li>online：启用；</li>
	// <li>offline：停用；</li>
	// <li>progress：部署中；</li>
	// <li>stopping：停用中；</li>
	// <li>fail：部署失败/停用失败。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 传递客户端IP，取值有：
	// <li>TOA：TOA（仅Proto=TCP时可选）；</li>
	// <li>PPV1：Proxy Protocol传递，协议版本V1（仅Proto=TCP时可选）；</li>
	// <li>PPV2：Proxy Protocol传递，协议版本V2；</li>
	// <li>OFF：不传递。</li>默认值：OFF。
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// 是否开启会话保持，取值有：
	// <li>true：开启；</li>
	// <li>false：关闭。</li>默认值：false。
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`
}

type AscriptionInfo struct {
	// 主机记录。
	Subdomain *string `json:"Subdomain,omitempty" name:"Subdomain"`

	// 记录类型。
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// 记录值。
	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`
}

type BillingDataFilter struct {
	// 参数名称，取值范围：
	// zone：站点名
	// host：域名
	// proxy: 四层实例
	// plan: 套餐
	Type *string `json:"Type,omitempty" name:"Type"`

	// 参数值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type BotConfig struct {
	// bot开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 通用详细基础规则。如果为null，默认使用历史配置。
	BotManagedRule *BotManagedRule `json:"BotManagedRule,omitempty" name:"BotManagedRule"`

	// 用户画像规则。如果为null，默认使用历史配置。
	BotPortraitRule *BotPortraitRule `json:"BotPortraitRule,omitempty" name:"BotPortraitRule"`

	// Bot智能分析。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntelligenceRule *IntelligenceRule `json:"IntelligenceRule,omitempty" name:"IntelligenceRule"`
}

type BotLog struct {
	// 攻击时间，采用unix秒级时间戳。
	AttackTime *uint64 `json:"AttackTime,omitempty" name:"AttackTime"`

	// 攻击源（客户端）ip。
	AttackIp *string `json:"AttackIp,omitempty" name:"AttackIp"`

	// 受攻击域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// URI。
	RequestUri *string `json:"RequestUri,omitempty" name:"RequestUri"`

	// 攻击类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 请求方法。
	RequestMethod *string `json:"RequestMethod,omitempty" name:"RequestMethod"`

	// 攻击内容。
	AttackContent *string `json:"AttackContent,omitempty" name:"AttackContent"`

	// 攻击等级。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 规则ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// IP所在国家iso-3166中alpha-2编码，编码信息请参考[ISO-3166](https://git.woa.com/edgeone/iso-3166/blob/master/all/all.json)。
	SipCountryCode *string `json:"SipCountryCode,omitempty" name:"SipCountryCode"`

	// 请求（事件）ID。
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 处置方式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisposalMethod *string `json:"DisposalMethod,omitempty" name:"DisposalMethod"`

	// HTTP日志。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpLog *string `json:"HttpLog,omitempty" name:"HttpLog"`

	// user agent。
	Ua *string `json:"Ua,omitempty" name:"Ua"`

	// 检出方法。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectionMethod *string `json:"DetectionMethod,omitempty" name:"DetectionMethod"`

	// 置信度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *string `json:"Confidence,omitempty" name:"Confidence"`

	// 恶意度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Maliciousness *string `json:"Maliciousness,omitempty" name:"Maliciousness"`

	// 规则相关信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleDetailList []*SecRuleRelatedInfo `json:"RuleDetailList,omitempty" name:"RuleDetailList"`

	// Bot标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`
}

type BotManagedRule struct {
	// 触发规则后的处置方式，取值有：
	// <li>drop：拦截；</li>
	// <li>trans：放行；</li>
	// <li>alg：Javascript挑战；</li>
	// <li>monitor：观察。</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// 本规则的ID。仅出参使用。
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// 放行的规则ID。默认所有规则不配置放行。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransManagedIds []*int64 `json:"TransManagedIds,omitempty" name:"TransManagedIds"`

	// JS挑战的规则ID。默认所有规则不配置JS挑战。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgManagedIds []*int64 `json:"AlgManagedIds,omitempty" name:"AlgManagedIds"`

	// 数字验证码的规则ID。默认所有规则不配置数字验证码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CapManagedIds []*int64 `json:"CapManagedIds,omitempty" name:"CapManagedIds"`

	// 观察的规则ID。默认所有规则不配置观察。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonManagedIds []*int64 `json:"MonManagedIds,omitempty" name:"MonManagedIds"`

	// 拦截的规则ID。默认所有规则不配置拦截。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DropManagedIds []*int64 `json:"DropManagedIds,omitempty" name:"DropManagedIds"`
}

type BotManagedRuleDetail struct {
	// 规则ID。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 规则分类。
	RuleTypeName *string `json:"RuleTypeName,omitempty" name:"RuleTypeName"`

	// 该规则开启/关闭状态。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type BotPortraitRule struct {
	// 本功能的开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 本规则的ID。仅出参使用。
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// JS挑战的规则ID。默认所有规则不配置JS挑战。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgManagedIds []*int64 `json:"AlgManagedIds,omitempty" name:"AlgManagedIds"`

	// 数字验证码的规则ID。默认所有规则不配置数字验证码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CapManagedIds []*int64 `json:"CapManagedIds,omitempty" name:"CapManagedIds"`

	// 观察的规则ID。默认所有规则不配置观察。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonManagedIds []*int64 `json:"MonManagedIds,omitempty" name:"MonManagedIds"`

	// 拦截的规则ID。默认所有规则不配置拦截。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DropManagedIds []*int64 `json:"DropManagedIds,omitempty" name:"DropManagedIds"`
}

type CC struct {
	// Waf开关，取值为：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

type CCInterceptEvent struct {
	// 客户端IP。
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 拦截次数/min。
	InterceptNum *int64 `json:"InterceptNum,omitempty" name:"InterceptNum"`

	// 速拦截时间，分钟时间/min，单位为s。
	InterceptTime *int64 `json:"InterceptTime,omitempty" name:"InterceptTime"`
}

type Cache struct {
	// 缓存配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 缓存过期时间设置。
	// 单位为秒，最大可设置为 365 天。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheTime *int64 `json:"CacheTime,omitempty" name:"CacheTime"`

	// 是否开启强制缓存，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreCacheControl *string `json:"IgnoreCacheControl,omitempty" name:"IgnoreCacheControl"`
}

type CacheConfig struct {
	// 缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// 不缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoCache *NoCache `json:"NoCache,omitempty" name:"NoCache"`

	// 遵循源站配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowOrigin *FollowOrigin `json:"FollowOrigin,omitempty" name:"FollowOrigin"`
}

type CacheKey struct {
	// 是否开启全路径缓存，取值有：
	// <li>on：开启全路径缓存（即关闭参数忽略）；</li>
	// <li>off：关闭全路径缓存（即开启参数忽略）。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullUrlCache *string `json:"FullUrlCache,omitempty" name:"FullUrlCache"`

	// 是否忽略大小写缓存，取值有：
	// <li>on：忽略；</li>
	// <li>off：不忽略。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreCase *string `json:"IgnoreCase,omitempty" name:"IgnoreCase"`

	// CacheKey中包含请求参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryString *QueryString `json:"QueryString,omitempty" name:"QueryString"`
}

type CachePrefresh struct {
	// 缓存预刷新配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 缓存预刷新百分比，取值范围：1-99。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent *int64 `json:"Percent,omitempty" name:"Percent"`
}

// Predefined struct for user
type CheckCertificateRequestParams struct {
	// 证书内容。
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// 私钥内容。
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
}

type CheckCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书内容。
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// 私钥内容。
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

// Predefined struct for user
type CheckCertificateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckCertificateResponse struct {
	*tchttp.BaseResponse
	Response *CheckCertificateResponseParams `json:"Response"`
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

type ClientIpHeader struct {
	// 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 回源时，存放客户端IP的请求头名称。
	// 为空则使用默认值：X-Forwarded-IP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`
}

type ClientRule struct {
	// 客户端ip。
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 规则类型。
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// 规则id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 封禁状态，取值有：
	// <li>block ：封禁 ；</li>
	// <li>allow ：放行 。</li>
	IpStatus *string `json:"IpStatus,omitempty" name:"IpStatus"`

	// 封禁时间，采用unix秒级时间戳。
	BlockTime *int64 `json:"BlockTime,omitempty" name:"BlockTime"`

	// 每条数据的唯一标识id。
	Id *string `json:"Id,omitempty" name:"Id"`
}

type ClsLogTopicInfo struct {
	// 任务名。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 日志集ID。
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 日志主题ID。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 任务类型。
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// 任务主题保存时间。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 任务主题是否开启。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 任务主题是否异常。
	Deleted *string `json:"Deleted,omitempty" name:"Deleted"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 推送目标地址,取值有：
	// <li>cls: 推送到cls；</li>
	// <li>custom_enpoint: 自定义推送地址。</li>
	Target *string `json:"Target,omitempty" name:"Target"`

	// 日志集所属地区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`

	// 站点id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 加速区域，取值有：
	// <li>mainland：中国大陆境内;</li>
	// <li>overseas：全球（不含中国大陆）。</li>
	Area *string `json:"Area,omitempty" name:"Area"`
}

type CodeAction struct {
	// 功能名称，功能名称填写规范可调用接口 [查询规则引擎的设置参数](https://tcloud4api.woa.com/document/product/1657/79433?!preview&!document=1) 查看。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 操作参数。
	Parameters []*RuleCodeActionParams `json:"Parameters,omitempty" name:"Parameters"`
}

type Compression struct {
	// 智能压缩配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 支持的压缩算法列表，取值有：
	// <li>brotli：brotli算法；</li>
	// <li>gzip：gzip算法。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Algorithms []*string `json:"Algorithms,omitempty" name:"Algorithms"`
}

// Predefined struct for user
type CreateApplicationProxyRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 当ProxyType=hostname时，表示域名或子域名；
	// 当ProxyType=instance时，表示代理名称。
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// 调度模式，取值有：
	// <li>ip：表示Anycast IP调度；</li>
	// <li>domain：表示CNAME调度。</li>
	PlatType *string `json:"PlatType,omitempty" name:"PlatType"`

	// 是否开启安全，取值有：
	// <li>0：关闭安全；</li>
	// <li>1：开启安全。</li>
	SecurityType *int64 `json:"SecurityType,omitempty" name:"SecurityType"`

	// 是否开启加速，取值有：
	// <li>0：关闭加速；</li>
	// <li>1：开启加速。</li>
	AccelerateType *int64 `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// 四层代理模式，取值有：
	// <li>hostname：表示子域名模式；</li>
	// <li>instance：表示实例模式。</li>不填写使用默认值instance。
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// 会话保持时间，取值范围：30-3600，单位：秒。
	// 不填写使用默认值600。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// Ipv6访问配置。
	// 不填写表示关闭Ipv6访问。
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// 规则详细信息。
	// 不填写则不创建规则。
	ApplicationProxyRules []*ApplicationProxyRule `json:"ApplicationProxyRules,omitempty" name:"ApplicationProxyRules"`
}

type CreateApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 当ProxyType=hostname时，表示域名或子域名；
	// 当ProxyType=instance时，表示代理名称。
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// 调度模式，取值有：
	// <li>ip：表示Anycast IP调度；</li>
	// <li>domain：表示CNAME调度。</li>
	PlatType *string `json:"PlatType,omitempty" name:"PlatType"`

	// 是否开启安全，取值有：
	// <li>0：关闭安全；</li>
	// <li>1：开启安全。</li>
	SecurityType *int64 `json:"SecurityType,omitempty" name:"SecurityType"`

	// 是否开启加速，取值有：
	// <li>0：关闭加速；</li>
	// <li>1：开启加速。</li>
	AccelerateType *int64 `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// 四层代理模式，取值有：
	// <li>hostname：表示子域名模式；</li>
	// <li>instance：表示实例模式。</li>不填写使用默认值instance。
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// 会话保持时间，取值范围：30-3600，单位：秒。
	// 不填写使用默认值600。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// Ipv6访问配置。
	// 不填写表示关闭Ipv6访问。
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// 规则详细信息。
	// 不填写则不创建规则。
	ApplicationProxyRules []*ApplicationProxyRule `json:"ApplicationProxyRules,omitempty" name:"ApplicationProxyRules"`
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
	delete(f, "ProxyName")
	delete(f, "PlatType")
	delete(f, "SecurityType")
	delete(f, "AccelerateType")
	delete(f, "ProxyType")
	delete(f, "SessionPersistTime")
	delete(f, "Ipv6")
	delete(f, "ApplicationProxyRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProxyResponseParams struct {
	// 新增的四层代理应用ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApplicationProxyResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationProxyResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateApplicationProxyRuleRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 协议，取值有：
	// <li>TCP：TCP协议；</li>
	// <li>UDP：UDP协议。</li>
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// 源站类型，取值有：
	// <li>custom：手动添加；</li>
	// <li>origins：源站组。</li>
	Port []*string `json:"Port,omitempty" name:"Port"`

	// 源站类型，取值：
	// custom：手动添加
	// origins：源站组
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 源站信息：
	// 当OriginType=custom时，表示一个或多个源站，如：
	// OriginValue=["8.8.8.8:80","9.9.9.9:80"]
	// OriginValue=["test.com:80"]；
	// 当OriginType=origins时，要求有且仅有一个元素，表示源站组ID，如：
	// OriginValue=["origin-537f5b41-162a-11ed-abaa-525400c5da15"]。
	OriginValue []*string `json:"OriginValue,omitempty" name:"OriginValue"`

	// 传递客户端IP，取值有：
	// <li>TOA：TOA（仅Proto=TCP时可选）；</li>
	// <li>PPV1：Proxy Protocol传递，协议版本V1（仅Proto=TCP时可选）；</li>
	// <li>PPV2：Proxy Protocol传递，协议版本V2；</li>
	// <li>OFF：不传递。</li>默认值：OFF。
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// 是否开启会话保持，取值有：
	// <li>true：开启；</li>
	// <li>false：关闭。</li>默认值：false。
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`
}

type CreateApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 协议，取值有：
	// <li>TCP：TCP协议；</li>
	// <li>UDP：UDP协议。</li>
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// 源站类型，取值有：
	// <li>custom：手动添加；</li>
	// <li>origins：源站组。</li>
	Port []*string `json:"Port,omitempty" name:"Port"`

	// 源站类型，取值：
	// custom：手动添加
	// origins：源站组
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 源站信息：
	// 当OriginType=custom时，表示一个或多个源站，如：
	// OriginValue=["8.8.8.8:80","9.9.9.9:80"]
	// OriginValue=["test.com:80"]；
	// 当OriginType=origins时，要求有且仅有一个元素，表示源站组ID，如：
	// OriginValue=["origin-537f5b41-162a-11ed-abaa-525400c5da15"]。
	OriginValue []*string `json:"OriginValue,omitempty" name:"OriginValue"`

	// 传递客户端IP，取值有：
	// <li>TOA：TOA（仅Proto=TCP时可选）；</li>
	// <li>PPV1：Proxy Protocol传递，协议版本V1（仅Proto=TCP时可选）；</li>
	// <li>PPV2：Proxy Protocol传递，协议版本V2；</li>
	// <li>OFF：不传递。</li>默认值：OFF。
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// 是否开启会话保持，取值有：
	// <li>true：开启；</li>
	// <li>false：关闭。</li>默认值：false。
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

// Predefined struct for user
type CreateApplicationProxyRuleResponseParams struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApplicationProxyRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationProxyRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateCredentialRequestParams struct {

}

type CreateCredentialRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCredentialResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCredentialResponse struct {
	*tchttp.BaseResponse
	Response *CreateCredentialResponseParams `json:"Response"`
}

func (r *CreateCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomErrorPageRequestParams struct {
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 自定义页面的文件名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 自定义页面的内容，本字段的内容需要将数据经过urlencode后传入。
	Content *string `json:"Content,omitempty" name:"Content"`
}

type CreateCustomErrorPageRequest struct {
	*tchttp.BaseRequest
	
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 自定义页面的文件名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 自定义页面的内容，本字段的内容需要将数据经过urlencode后传入。
	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *CreateCustomErrorPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomErrorPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entity")
	delete(f, "Name")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomErrorPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomErrorPageResponseParams struct {
	// 自定义页面上传后的唯一id。
	PageId *int64 `json:"PageId,omitempty" name:"PageId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCustomErrorPageResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomErrorPageResponseParams `json:"Response"`
}

func (r *CreateCustomErrorPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomErrorPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDnsRecordRequestParams struct {
	// DNS记录所属站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// DNS记录类型，取值有：
	// <li>A：将域名指向一个外网 IPv4 地址，如 8.8.8.8；</li>
	// <li>AAAA：将域名指向一个外网 IPv6 地址；</li>
	// <li>MX：用于邮箱服务器，相关记录值/优先级参数由邮件注册商提供。存在多条 MX 记录时，优先级越低越优先；</li>
	// <li>CNAME：将域名指向另一个域名，再由该域名解析出最终 IP 地址；</li>
	// <li>TXT：对域名进行标识和说明，常用于域名验证和 SPF 记录（反垃圾邮件）；</li>
	// <li>NS：如果需要将子域名交给其他 DNS 服务商解析，则需要添加 NS 记录。根域名无法添加 NS 记录；</li>
	// <li>CAA：指定可为本站点颁发证书的 CA；</li>
	// <li>SRV：标识某台服务器使用了某个服务，常见于微软系统的目录管理。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// DNS记录名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// DNS记录内容。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 代理模式，取值有：
	// <li>dns_only：仅DNS解析；</li>
	// <li>proxied：代理加速。</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 缓存时间，数值越小，修改记录各地生效时间越快，默认为300，单位：秒。
	TTL *int64 `json:"TTL,omitempty" name:"TTL"`

	// 该参数在创建MX记录时生效，值越小优先级越高，用户可指定值范围1~50，不指定默认为0。
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
}

type CreateDnsRecordRequest struct {
	*tchttp.BaseRequest
	
	// DNS记录所属站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// DNS记录类型，取值有：
	// <li>A：将域名指向一个外网 IPv4 地址，如 8.8.8.8；</li>
	// <li>AAAA：将域名指向一个外网 IPv6 地址；</li>
	// <li>MX：用于邮箱服务器，相关记录值/优先级参数由邮件注册商提供。存在多条 MX 记录时，优先级越低越优先；</li>
	// <li>CNAME：将域名指向另一个域名，再由该域名解析出最终 IP 地址；</li>
	// <li>TXT：对域名进行标识和说明，常用于域名验证和 SPF 记录（反垃圾邮件）；</li>
	// <li>NS：如果需要将子域名交给其他 DNS 服务商解析，则需要添加 NS 记录。根域名无法添加 NS 记录；</li>
	// <li>CAA：指定可为本站点颁发证书的 CA；</li>
	// <li>SRV：标识某台服务器使用了某个服务，常见于微软系统的目录管理。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// DNS记录名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// DNS记录内容。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 代理模式，取值有：
	// <li>dns_only：仅DNS解析；</li>
	// <li>proxied：代理加速。</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 缓存时间，数值越小，修改记录各地生效时间越快，默认为300，单位：秒。
	TTL *int64 `json:"TTL,omitempty" name:"TTL"`

	// 该参数在创建MX记录时生效，值越小优先级越高，用户可指定值范围1~50，不指定默认为0。
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
	delete(f, "TTL")
	delete(f, "Priority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDnsRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDnsRecordResponseParams struct {
	// DNS解析记录ID。
	DnsRecordId *string `json:"DnsRecordId,omitempty" name:"DnsRecordId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDnsRecordResponse struct {
	*tchttp.BaseResponse
	Response *CreateDnsRecordResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateIpTableListRequestParams struct {
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 基础访问管控Ip规则列表。
	IpTableRules []*IpTableRule `json:"IpTableRules,omitempty" name:"IpTableRules"`
}

type CreateIpTableListRequest struct {
	*tchttp.BaseRequest
	
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 基础访问管控Ip规则列表。
	IpTableRules []*IpTableRule `json:"IpTableRules,omitempty" name:"IpTableRules"`
}

func (r *CreateIpTableListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIpTableListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entity")
	delete(f, "IpTableRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIpTableListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIpTableListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateIpTableListResponse struct {
	*tchttp.BaseResponse
	Response *CreateIpTableListResponseParams `json:"Response"`
}

func (r *CreateIpTableListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIpTableListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadBalancingRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 负载均衡域名。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 代理模式，取值有：
	// <li>dns_only：仅DNS；</li>
	// <li>proxied：开启代理。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 主源站源站组ID。
	OriginGroupId *string `json:"OriginGroupId,omitempty" name:"OriginGroupId"`

	// 备用源站源站组ID，当Type=proxied时可以填写，为空表示不使用备用源站。
	BackupOriginGroupId *string `json:"BackupOriginGroupId,omitempty" name:"BackupOriginGroupId"`

	// 当Type=dns_only时，指解析记录在DNS服务器缓存的生存时间。
	// 取值范围60-86400，单位：秒，不填写使用默认值：600。
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`
}

type CreateLoadBalancingRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 负载均衡域名。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 代理模式，取值有：
	// <li>dns_only：仅DNS；</li>
	// <li>proxied：开启代理。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 主源站源站组ID。
	OriginGroupId *string `json:"OriginGroupId,omitempty" name:"OriginGroupId"`

	// 备用源站源站组ID，当Type=proxied时可以填写，为空表示不使用备用源站。
	BackupOriginGroupId *string `json:"BackupOriginGroupId,omitempty" name:"BackupOriginGroupId"`

	// 当Type=dns_only时，指解析记录在DNS服务器缓存的生存时间。
	// 取值范围60-86400，单位：秒，不填写使用默认值：600。
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
	delete(f, "OriginGroupId")
	delete(f, "BackupOriginGroupId")
	delete(f, "TTL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoadBalancingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadBalancingResponseParams struct {
	// 负载均衡ID。
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLoadBalancingResponse struct {
	*tchttp.BaseResponse
	Response *CreateLoadBalancingResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateLogSetRequestParams struct {
	// 日志集名称。
	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`

	// 日志集归属的地域。
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`
}

type CreateLogSetRequest struct {
	*tchttp.BaseRequest
	
	// 日志集名称。
	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`

	// 日志集归属的地域。
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`
}

func (r *CreateLogSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogSetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogSetName")
	delete(f, "LogSetRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLogSetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLogSetResponseParams struct {
	// 创建成功的日志集ID。
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLogSetResponse struct {
	*tchttp.BaseResponse
	Response *CreateLogSetResponseParams `json:"Response"`
}

func (r *CreateLogSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLogTopicTaskRequestParams struct {
	// 日志集ID。
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 日志集所属的地域。
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`

	// 日志集主题名。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 推送任务的名称。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 数据推送类型，取值有：
	// <li>domain：七层代理日志；</li>
	// <li>application：四层代理日志；</li>
	// <li>web-rateLiming：速率限制日志；</li>
	// <li>web-attack：Web攻击防护日志；</li>
	// <li>web-rule：自定义规则日志；</li>
	// <li>web-bot：Bot管理日志。</li>
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// 日志主题保存时间，单位为天，取值范围为：1-366。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 推送任务实体列表。
	EntityList []*string `json:"EntityList,omitempty" name:"EntityList"`

	// 加速区域，取值有：
	// <li>mainland：中国大陆境内；</li>
	// <li>overseas：全球（不含中国大陆）。</li>不填将根据用户的地域智能选择加速区域。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type CreateLogTopicTaskRequest struct {
	*tchttp.BaseRequest
	
	// 日志集ID。
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 日志集所属的地域。
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`

	// 日志集主题名。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 推送任务的名称。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 数据推送类型，取值有：
	// <li>domain：七层代理日志；</li>
	// <li>application：四层代理日志；</li>
	// <li>web-rateLiming：速率限制日志；</li>
	// <li>web-attack：Web攻击防护日志；</li>
	// <li>web-rule：自定义规则日志；</li>
	// <li>web-bot：Bot管理日志。</li>
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// 日志主题保存时间，单位为天，取值范围为：1-366。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 推送任务实体列表。
	EntityList []*string `json:"EntityList,omitempty" name:"EntityList"`

	// 加速区域，取值有：
	// <li>mainland：中国大陆境内；</li>
	// <li>overseas：全球（不含中国大陆）。</li>不填将根据用户的地域智能选择加速区域。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *CreateLogTopicTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogTopicTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogSetId")
	delete(f, "LogSetRegion")
	delete(f, "TopicName")
	delete(f, "TaskName")
	delete(f, "ZoneId")
	delete(f, "ZoneName")
	delete(f, "EntityType")
	delete(f, "Period")
	delete(f, "EntityList")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLogTopicTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLogTopicTaskResponseParams struct {
	// 创建成功的主题ID。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLogTopicTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateLogTopicTaskResponseParams `json:"Response"`
}

func (r *CreateLogTopicTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogTopicTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOriginGroupRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 源站类型，取值有：
	// <li>self：自有源站；</li>
	// <li>third_party：第三方源站；</li>
	// <li>cos：腾讯云COS源站。</li>
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 源站组名称。
	OriginGroupName *string `json:"OriginGroupName,omitempty" name:"OriginGroupName"`

	// 源站配置类型，当OriginType=self时，取值有：
	// <li>area：按区域配置；</li>
	// <li>weight： 按权重配置；</li>
	// <li>proto： 按HTTP协议配置。</li>当OriginType=third_party/cos时放空。
	ConfigurationType *string `json:"ConfigurationType,omitempty" name:"ConfigurationType"`

	// 源站记录信息。
	OriginRecords []*OriginRecord `json:"OriginRecords,omitempty" name:"OriginRecords"`
}

type CreateOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 源站类型，取值有：
	// <li>self：自有源站；</li>
	// <li>third_party：第三方源站；</li>
	// <li>cos：腾讯云COS源站。</li>
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 源站组名称。
	OriginGroupName *string `json:"OriginGroupName,omitempty" name:"OriginGroupName"`

	// 源站配置类型，当OriginType=self时，取值有：
	// <li>area：按区域配置；</li>
	// <li>weight： 按权重配置；</li>
	// <li>proto： 按HTTP协议配置。</li>当OriginType=third_party/cos时放空。
	ConfigurationType *string `json:"ConfigurationType,omitempty" name:"ConfigurationType"`

	// 源站记录信息。
	OriginRecords []*OriginRecord `json:"OriginRecords,omitempty" name:"OriginRecords"`
}

func (r *CreateOriginGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOriginGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "OriginType")
	delete(f, "OriginGroupName")
	delete(f, "ConfigurationType")
	delete(f, "OriginRecords")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOriginGroupResponseParams struct {
	// 源站组ID。
	OriginGroupId *string `json:"OriginGroupId,omitempty" name:"OriginGroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOriginGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateOriginGroupResponseParams `json:"Response"`
}

func (r *CreateOriginGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOriginGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePlanForZoneRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 所要购买套餐的类型，取值有：
	// <li> sta: 全球内容分发网络（不包括中国大陆）标准版套餐； </li>
	// <li> sta_with_bot: 全球内容分发网络（不包括中国大陆）标准版套餐附带bot管理；</li>
	// <li> sta_cm: 中国大陆内容分发网络标准版套餐； </li>
	// <li> sta_cm_with_bot: 中国大陆内容分发网络标准版套餐附带bot管理；</li>
	// <li> ent: 全球内容分发网络（不包括中国大陆）企业版套餐； </li>
	// <li> ent_with_bot: 全球内容分发网络（不包括中国大陆）企业版套餐附带bot管理；</li>
	// <li> ent_cm: 中国大陆内容分发网络企业版套餐； </li>
	// <li> ent_cm_with_bot: 中国大陆内容分发网络企业版套餐附带bot管理。</li>当前账户可购买套餐类型请以<a href="https://tcloud4api.woa.com/document/product/1657/80124?!preview&!document=1">DescribeAvailablePlans</a>返回为准。
	PlanType *string `json:"PlanType,omitempty" name:"PlanType"`
}

type CreatePlanForZoneRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 所要购买套餐的类型，取值有：
	// <li> sta: 全球内容分发网络（不包括中国大陆）标准版套餐； </li>
	// <li> sta_with_bot: 全球内容分发网络（不包括中国大陆）标准版套餐附带bot管理；</li>
	// <li> sta_cm: 中国大陆内容分发网络标准版套餐； </li>
	// <li> sta_cm_with_bot: 中国大陆内容分发网络标准版套餐附带bot管理；</li>
	// <li> ent: 全球内容分发网络（不包括中国大陆）企业版套餐； </li>
	// <li> ent_with_bot: 全球内容分发网络（不包括中国大陆）企业版套餐附带bot管理；</li>
	// <li> ent_cm: 中国大陆内容分发网络企业版套餐； </li>
	// <li> ent_cm_with_bot: 中国大陆内容分发网络企业版套餐附带bot管理。</li>当前账户可购买套餐类型请以<a href="https://tcloud4api.woa.com/document/product/1657/80124?!preview&!document=1">DescribeAvailablePlans</a>返回为准。
	PlanType *string `json:"PlanType,omitempty" name:"PlanType"`
}

func (r *CreatePlanForZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePlanForZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "PlanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePlanForZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePlanForZoneResponseParams struct {
	// 购买的资源名字列表。
	ResourceNames []*string `json:"ResourceNames,omitempty" name:"ResourceNames"`

	// 购买的订单号列表。
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePlanForZoneResponse struct {
	*tchttp.BaseResponse
	Response *CreatePlanForZoneResponseParams `json:"Response"`
}

func (r *CreatePlanForZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePlanForZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrefetchTaskRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 要预热的资源列表，每个元素格式类似如下:
	// http://www.example.com/example.txt。
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// 是否对url进行encode，若内容含有非 ASCII 字符集的字符，请开启此开关进行编码转换（编码规则遵循 RFC3986）。
	EncodeUrl *bool `json:"EncodeUrl,omitempty" name:"EncodeUrl"`

	// 附带的http头部信息。
	Headers []*Header `json:"Headers,omitempty" name:"Headers"`
}

type CreatePrefetchTaskRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 要预热的资源列表，每个元素格式类似如下:
	// http://www.example.com/example.txt。
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// 是否对url进行encode，若内容含有非 ASCII 字符集的字符，请开启此开关进行编码转换（编码规则遵循 RFC3986）。
	EncodeUrl *bool `json:"EncodeUrl,omitempty" name:"EncodeUrl"`

	// 附带的http头部信息。
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

// Predefined struct for user
type CreatePrefetchTaskResponseParams struct {
	// 任务 ID。
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 失败的任务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedList []*FailReason `json:"FailedList,omitempty" name:"FailedList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePrefetchTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrefetchTaskResponseParams `json:"Response"`
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

// Predefined struct for user
type CreatePurgeTaskRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 清除缓存类型，取值有：
	// <li>purge_url：URL；</li>
	// <li>purge_prefix：前缀；</li>
	// <li>purge_host：Hostname；</li>
	// <li>purge_all：全部缓存。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 要刷新的资源列表，每个元素格式依据Type而定：
	// 1) Type = purge_host 时：
	// 形如：www.example.com 或 foo.bar.example.com。
	// 2) Type = purge_prefix 时：
	// 形如：http://www.example.com/example。
	// 3) Type = purge_url 时：
	// 形如：https://www.example.com/example.jpg。
	// 4）Type = purge_all 时：
	// Targets可为空，不需要填写。
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// 若有编码转换，仅清除编码转换后匹配的资源。
	// 若内容含有非 ASCII 字符集的字符，请开启此开关进行编码转换（编码规则遵循 RFC3986）。
	EncodeUrl *bool `json:"EncodeUrl,omitempty" name:"EncodeUrl"`
}

type CreatePurgeTaskRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 清除缓存类型，取值有：
	// <li>purge_url：URL；</li>
	// <li>purge_prefix：前缀；</li>
	// <li>purge_host：Hostname；</li>
	// <li>purge_all：全部缓存。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 要刷新的资源列表，每个元素格式依据Type而定：
	// 1) Type = purge_host 时：
	// 形如：www.example.com 或 foo.bar.example.com。
	// 2) Type = purge_prefix 时：
	// 形如：http://www.example.com/example。
	// 3) Type = purge_url 时：
	// 形如：https://www.example.com/example.jpg。
	// 4）Type = purge_all 时：
	// Targets可为空，不需要填写。
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// 若有编码转换，仅清除编码转换后匹配的资源。
	// 若内容含有非 ASCII 字符集的字符，请开启此开关进行编码转换（编码规则遵循 RFC3986）。
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

// Predefined struct for user
type CreatePurgeTaskResponseParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 失败的任务列表及原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedList []*FailReason `json:"FailedList,omitempty" name:"FailedList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePurgeTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreatePurgeTaskResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateReplayTaskRequestParams struct {
	// 重放任务的 ID 列表。
	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

type CreateReplayTaskRequest struct {
	*tchttp.BaseRequest
	
	// 重放任务的 ID 列表。
	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *CreateReplayTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReplayTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReplayTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReplayTaskResponseParams struct {
	// 此次任务ID。
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 失败的任务列表及原因。
	FailedList []*FailReason `json:"FailedList,omitempty" name:"FailedList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateReplayTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateReplayTaskResponseParams `json:"Response"`
}

func (r *CreateReplayTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReplayTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 规则名称，名称字符串长度 1～255。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则状态，取值有：
	// <li> enable: 启用； </li>
	// <li> disable: 未启用。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 规则内容。
	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 规则名称，名称字符串长度 1～255。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则状态，取值有：
	// <li> enable: 启用； </li>
	// <li> disable: 未启用。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 规则内容。
	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`
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
	delete(f, "ZoneId")
	delete(f, "RuleName")
	delete(f, "Status")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleResponseParams struct {
	// 规则 ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateSecurityDropPageRequestParams struct {
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点子域名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 自定义页面的文件名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 自定义页面的内容，本字段的内容需要将数据经过urlencode后传入。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 上传的类型，取值有：
	// <li> file：将页面文件内容进行urlencode编码，填入Content字段；</li>
	// <li> url：将待上传的url地址进行urlencode编码，填入Content字段，即时下载，内容不会自动更新。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 页面所属的模块，取值有：
	// <li> waf ：托管规则模块；</li>
	// <li> rate：自定义规则模块。</li>
	Module *string `json:"Module,omitempty" name:"Module"`
}

type CreateSecurityDropPageRequest struct {
	*tchttp.BaseRequest
	
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点子域名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 自定义页面的文件名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 自定义页面的内容，本字段的内容需要将数据经过urlencode后传入。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 上传的类型，取值有：
	// <li> file：将页面文件内容进行urlencode编码，填入Content字段；</li>
	// <li> url：将待上传的url地址进行urlencode编码，填入Content字段，即时下载，内容不会自动更新。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 页面所属的模块，取值有：
	// <li> waf ：托管规则模块；</li>
	// <li> rate：自定义规则模块。</li>
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *CreateSecurityDropPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityDropPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entity")
	delete(f, "Name")
	delete(f, "Content")
	delete(f, "Type")
	delete(f, "Module")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityDropPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityDropPageResponseParams struct {
	// 自定义页面上传后的唯一id。
	PageId *int64 `json:"PageId,omitempty" name:"PageId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSecurityDropPageResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityDropPageResponseParams `json:"Response"`
}

func (r *CreateSecurityDropPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityDropPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSpeedTestingRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 拨测子域名。
	Host *string `json:"Host,omitempty" name:"Host"`
}

type CreateSpeedTestingRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 拨测子域名。
	Host *string `json:"Host,omitempty" name:"Host"`
}

func (r *CreateSpeedTestingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSpeedTestingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSpeedTestingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSpeedTestingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSpeedTestingResponse struct {
	*tchttp.BaseResponse
	Response *CreateSpeedTestingResponseParams `json:"Response"`
}

func (r *CreateSpeedTestingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSpeedTestingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateZoneRequestParams struct {
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 接入方式，取值有：
	// <li> full：NS接入；</li>
	// <li> partial：CNAME接入。</li>不填写使用默认值full。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 是否跳过站点现有的DNS记录扫描。默认值：false。
	JumpStart *bool `json:"JumpStart,omitempty" name:"JumpStart"`

	// 资源标签。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateZoneRequest struct {
	*tchttp.BaseRequest
	
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 接入方式，取值有：
	// <li> full：NS接入；</li>
	// <li> partial：CNAME接入。</li>不填写使用默认值full。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 是否跳过站点现有的DNS记录扫描。默认值：false。
	JumpStart *bool `json:"JumpStart,omitempty" name:"JumpStart"`

	// 资源标签。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
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
	delete(f, "ZoneName")
	delete(f, "Type")
	delete(f, "JumpStart")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateZoneResponseParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateZoneResponse struct {
	*tchttp.BaseResponse
	Response *CreateZoneResponseParams `json:"Response"`
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

type DDoS struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type DDoSAcl struct {
	// 端口过滤规则数组。
	DDoSAclRules []*DDoSAclRule `json:"DDoSAclRules,omitempty" name:"DDoSAclRules"`

	// 清空规则标识，取值有：
	// <li>off ：清空端口过滤规则列表，DDoSAclRules无需填写；</li>
	// <li>on ：配置端口过滤规则，需填写DDoSAclRules参数。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type DDoSAclRule struct {
	// 目的端口结束，取值范围0-65535。
	DportEnd *int64 `json:"DportEnd,omitempty" name:"DportEnd"`

	// 目的端口开始，取值范围0-65535。
	DportStart *int64 `json:"DportStart,omitempty" name:"DportStart"`

	// 源端口结束，取值范围0-65535。
	SportEnd *int64 `json:"SportEnd,omitempty" name:"SportEnd"`

	// 源端口开始，取值范围0-65535。
	SportStart *int64 `json:"SportStart,omitempty" name:"SportStart"`

	// 协议，取值有：
	// <li>tcp ：tcp协议 ；</li>
	// <li>udp ：udp协议 ；</li>
	// <li>all ：全部协议 。</li>
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 执行动作，取值为：
	// <li>drop ：丢弃 ；</li>
	// <li>transmit ：放行 ；</li>
	// <li>forward ：继续防护 。</li>
	Action *string `json:"Action,omitempty" name:"Action"`
}

type DDoSAllowBlock struct {
	// 黑白名单数组。
	DDoSAllowBlockRules []*DDoSAllowBlockRule `json:"DDoSAllowBlockRules,omitempty" name:"DDoSAllowBlockRules"`

	// 开关标识防护是否清空，取值有：
	// <li>off ：关闭黑白名单；</li>
	// <li>on ：开启黑白名单，需填写UserAllowBlockIp参数。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type DDoSAllowBlockRule struct {
	// 客户端IP，支持格式有：单IP、IP范围、网段、网段范围。例如"1.1.1.1","1.1.1.2-1.1.1.3","1.2.1.0/24-1.2.2.0/24"。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 类型，取值有：
	// <li> block ：丢弃 ；</li><li> allow ：允许。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 10位时间戳，例如1199116800。不填写系统取当前时间。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DDoSAntiPly struct {
	// tcp协议封禁，取值有：
	// <li>off ：关闭 ；</li>
	// <li>on ：开启 。</li>
	DropTcp *string `json:"DropTcp,omitempty" name:"DropTcp"`

	// udp协议封禁，取值有：
	// <li>off ：关闭 ；</li>
	// <li>on ：开启 。</li>
	DropUdp *string `json:"DropUdp,omitempty" name:"DropUdp"`

	// icmp协议封禁，取值有：
	// <li>off ：关闭 ；</li>
	// <li>on ：开启 。</li>
	DropIcmp *string `json:"DropIcmp,omitempty" name:"DropIcmp"`

	// 其他协议封禁，取值有：
	// <li>off ：关闭 ；</li>
	// <li>on ：开启 。</li>
	DropOther *string `json:"DropOther,omitempty" name:"DropOther"`

	// 源站每秒新连接限速，取值范围0-4294967295。
	SourceCreateLimit *int64 `json:"SourceCreateLimit,omitempty" name:"SourceCreateLimit"`

	// 源站并发连接控制，取值范围0-4294967295。
	SourceConnectLimit *int64 `json:"SourceConnectLimit,omitempty" name:"SourceConnectLimit"`

	// 目的端口每秒新连接限速，取值范围0-4294967295。
	DestinationCreateLimit *int64 `json:"DestinationCreateLimit,omitempty" name:"DestinationCreateLimit"`

	// 目的端口并发连接控制，取值范围0-4294967295。
	DestinationConnectLimit *int64 `json:"DestinationConnectLimit,omitempty" name:"DestinationConnectLimit"`

	// 每秒异常连接数阈值，取值范围0-4294967295。
	AbnormalConnectNum *int64 `json:"AbnormalConnectNum,omitempty" name:"AbnormalConnectNum"`

	// 异常syn报文百分比阈值，取值范围0-100。
	AbnormalSynRatio *int64 `json:"AbnormalSynRatio,omitempty" name:"AbnormalSynRatio"`

	// 异常syn报文阈值，取值范围0-65535。
	AbnormalSynNum *int64 `json:"AbnormalSynNum,omitempty" name:"AbnormalSynNum"`

	// 每秒连接超时检测，取值范围0-65535。
	ConnectTimeout *int64 `json:"ConnectTimeout,omitempty" name:"ConnectTimeout"`

	// 空连接防护开启，取值有：
	// <li>off ：关闭 ；</li>
	// <li>on ：开启 。</li>
	EmptyConnectProtect *string `json:"EmptyConnectProtect,omitempty" name:"EmptyConnectProtect"`

	// UDP分片开关，取值有：
	// <li>off ：关闭 ；</li>
	// <li>on ：开启 。</li>
	UdpShard *string `json:"UdpShard,omitempty" name:"UdpShard"`
}

type DDoSAttackEvent struct {
	// 事件ID。
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 攻击类型(对应交互事件名称)。
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 攻击状态。
	AttackStatus *int64 `json:"AttackStatus,omitempty" name:"AttackStatus"`

	// 攻击最大带宽。
	AttackMaxBandWidth *int64 `json:"AttackMaxBandWidth,omitempty" name:"AttackMaxBandWidth"`

	// 攻击包速率峰值。
	AttackPacketMaxRate *int64 `json:"AttackPacketMaxRate,omitempty" name:"AttackPacketMaxRate"`

	// 攻击开始时间，单位为s。
	AttackStartTime *int64 `json:"AttackStartTime,omitempty" name:"AttackStartTime"`

	// 攻击结束时间，单位为s。
	AttackEndTime *int64 `json:"AttackEndTime,omitempty" name:"AttackEndTime"`

	// DDoS策略组ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 站点ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DDoSAttackEventDetailData struct {
	// 攻击状态，取值有：
	// <li>1 ：观察中 ；</li>
	// <li>2 ：攻击开始 ；</li>
	// <li>3 ：攻击结束 。</li>
	AttackStatus *int64 `json:"AttackStatus,omitempty" name:"AttackStatus"`

	// 攻击类型。
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 结束时间。
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 开始时间。
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 最大带宽。
	MaxBandWidth *int64 `json:"MaxBandWidth,omitempty" name:"MaxBandWidth"`

	// 最大包速率。
	PacketMaxRate *int64 `json:"PacketMaxRate,omitempty" name:"PacketMaxRate"`

	// 事件Id。
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// DDoS策略组ID。
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

type DDoSAttackSourceEvent struct {
	// 攻击源ip。
	AttackSourceIp *string `json:"AttackSourceIp,omitempty" name:"AttackSourceIp"`

	// 地区（国家）。
	AttackRegion *string `json:"AttackRegion,omitempty" name:"AttackRegion"`

	// 累计攻击流量。
	AttackFlow *uint64 `json:"AttackFlow,omitempty" name:"AttackFlow"`

	// 累计攻击包量。
	AttackPacketNum *uint64 `json:"AttackPacketNum,omitempty" name:"AttackPacketNum"`
}

type DDoSBlockData struct {
	// 开始时间，采用unix时间戳。
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，采用unix时间戳。
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type DDoSFeaturesFilter struct {
	// 执行动作，取值有：
	// <li>drop ：丢弃 ；</li>
	// <li>transmit ：放行 ；</li>
	// <li>drop_block ：丢弃并拉黑 ；</li>
	// <li>forward ：继续防护 。</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// 协议，取值有：
	// <li>tcp ：tcp协议 ；</li>
	// <li>udp ：udp协议 ；</li>
	// <li>icmp ：icmp协议 ；</li>
	// <li>all ：全部协议 。</li>
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 目标端口开始，取值范围0-65535。
	DportStart *int64 `json:"DportStart,omitempty" name:"DportStart"`

	// 目标端口结束，取值范围0-65535。
	DportEnd *int64 `json:"DportEnd,omitempty" name:"DportEnd"`

	// 最小包长，取值范围0-1500。
	PacketMin *int64 `json:"PacketMin,omitempty" name:"PacketMin"`

	// 最大包长，取值范围0-1500。
	PacketMax *int64 `json:"PacketMax,omitempty" name:"PacketMax"`

	// 源端口开始，取值范围0-65535。
	SportStart *int64 `json:"SportStart,omitempty" name:"SportStart"`

	// 源端口结束，取值范围0-65535。
	SportEnd *int64 `json:"SportEnd,omitempty" name:"SportEnd"`

	// 匹配方式1，【特征1】，取值有：
	// <li>pcre ：正则匹配 ；</li>
	// <li>sunday ：字符串匹配 。</li>默认为空字符串。
	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`

	// 取非判断，该参数对MatchType配合使用，【特征1】，取值有：
	// <li>0 ：匹配 ；</li>
	// <li>1 ：不匹配 。</li>
	IsNot *int64 `json:"IsNot,omitempty" name:"IsNot"`

	// 偏移量1，【特征1】，取值范围0-1500。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 检测包字符深度，【特征1】，取值范围1-1500。
	Depth *int64 `json:"Depth,omitempty" name:"Depth"`

	// 匹配开始层级，层级参考计算机网络结构，取值有：
	// <li>begin_l5 ：载荷开始检测 ；</li>
	// <li>begin_l4 ：tcp/udp首部开始检测 ；</li>
	// <li>begin_l3 ：ip首部开始检测 。</li>
	MatchBegin *string `json:"MatchBegin,omitempty" name:"MatchBegin"`

	// 正则或字符串匹配的内容，【特征1】。
	Str *string `json:"Str,omitempty" name:"Str"`

	// 匹配方式2，【特征2】，取值有：
	// <li>pcre ：正则匹配 ；</li>
	// <li>sunday ：字符串匹配 。</li>默认为空字符串。
	MatchType2 *string `json:"MatchType2,omitempty" name:"MatchType2"`

	// 取非判断2，该参数对MatchType2配合使用，【特征2】，取值有：
	// <li>0 ：匹配 ；</li>
	// <li>1 ：不匹配 。</li>
	IsNot2 *int64 `json:"IsNot2,omitempty" name:"IsNot2"`

	// 偏移量2，【特征2】，取值范围0-1500。
	Offset2 *int64 `json:"Offset2,omitempty" name:"Offset2"`

	// 检测包字符深度，【特征2】，取值范围1-1500。
	Depth2 *int64 `json:"Depth2,omitempty" name:"Depth2"`

	// 匹配开始层级，层级参考计算机网络结构，取值有：
	// <li>begin_l5 ：载荷开始检测 ；</li>
	// <li>begin_l4 ：tcp/udp首部开始检测；</li>
	// <li>begin_l3 ：ip首部开始检测 。</li>
	MatchBegin2 *string `json:"MatchBegin2,omitempty" name:"MatchBegin2"`

	// 正则或字符串匹配的内容，【特征2】。
	Str2 *string `json:"Str2,omitempty" name:"Str2"`

	// 多特征关系，仅配置【特征1】时填 none，存在【特征2】配置可不填。
	MatchLogic *string `json:"MatchLogic,omitempty" name:"MatchLogic"`
}

type DDoSGeoIp struct {
	// 区域封禁清空标识，取值有：
	// <li>off ：清空地域封禁列表 ；</li>
	// <li>on ：不做处理 。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 地域信息，ID参考[DescribeSecurityPolicyRegions](https://tcloud4api.woa.com/document/product/1657/81247?!preview&!document=1)。
	RegionIds []*int64 `json:"RegionIds,omitempty" name:"RegionIds"`
}

type DDoSHost struct {
	// 二级域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 域名状态；
	// init  待切ns
	// offline 需要dns开启站点加速
	// process 在部署中，稍等一会
	// online 正常状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 加速开关；on-开启加速；off-关闭加速（AccelerateType：on，SecurityType：on，安全加速，未开防护增强；AccelerateType：off，SecurityType：on，安全加速，开启防护增强；AccelerateType：on，SecurityType：off，内容加速，未开防护增强）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccelerateType *string `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// 安全开关；on-开启安全；off-关闭安全（AccelerateType：on，SecurityType：on，安全加速，未开防护增强；AccelerateType：off，SecurityType：on，安全加速，开启防护增强；AccelerateType：on，SecurityType：off，内容加速，未开防护增强）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityType *string `json:"SecurityType,omitempty" name:"SecurityType"`
}

type DDoSMajorAttackEvent struct {
	// ddos 策略组id。
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 攻击最大带宽。
	AttackMaxBandWidth *int64 `json:"AttackMaxBandWidth,omitempty" name:"AttackMaxBandWidth"`

	// 攻击请求时间，采用unix秒级时间戳。
	AttackTime *int64 `json:"AttackTime,omitempty" name:"AttackTime"`
}

type DDoSPacketFilter struct {
	// 特征过滤规则数组。
	DDoSFeaturesFilters []*DDoSFeaturesFilter `json:"DDoSFeaturesFilters,omitempty" name:"DDoSFeaturesFilters"`

	// 特征过滤清空标识，取值有：
	// <li>off ：清空特征过滤规则，无需填写 DDoSFeaturesFilters 参数 ；</li>
	// <li>on ：配置特征过滤规则，需填写 DDoSFeaturesFilters 参数。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type DDoSRule struct {
	// DDoS防护等级。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DDoSStatusInfo *DDoSStatusInfo `json:"DDoSStatusInfo,omitempty" name:"DDoSStatusInfo"`

	// DDoS地域封禁。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DDoSGeoIp *DDoSGeoIp `json:"DDoSGeoIp,omitempty" name:"DDoSGeoIp"`

	// DDoS黑白名单。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DDoSAllowBlock *DDoSAllowBlock `json:"DDoSAllowBlock,omitempty" name:"DDoSAllowBlock"`

	// DDoS 协议封禁+连接防护。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DDoSAntiPly *DDoSAntiPly `json:"DDoSAntiPly,omitempty" name:"DDoSAntiPly"`

	// DDoS特征过滤。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DDoSPacketFilter *DDoSPacketFilter `json:"DDoSPacketFilter,omitempty" name:"DDoSPacketFilter"`

	// DDoS端口过滤。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DDoSAcl *DDoSAcl `json:"DDoSAcl,omitempty" name:"DDoSAcl"`

	// DDoS开关，取值有:
	// <li>on ：开启 ；</li>
	// <li>off ：关闭 。</li>如果为null，默认使用历史配置。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// UDP分片功能是否支持，取值有:
	// <li>on ：支持 ；</li>
	// <li>off ：不支持 。</li>仅出参字段，入参无需填写。
	UdpShardOpen *string `json:"UdpShardOpen,omitempty" name:"UdpShardOpen"`

	// DDoS源站访问速率限制。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DDoSSpeedLimit *DDoSSpeedLimit `json:"DDoSSpeedLimit,omitempty" name:"DDoSSpeedLimit"`
}

type DDoSSpeedLimit struct {
	// 源站包量限制，支持单位有pps、Kpps、Mpps、Gpps。支持范围1 pps-10000 Gpps。"0 pps"或其他单位数值为0，即不限包。""为不更新。
	PackageLimit *string `json:"PackageLimit,omitempty" name:"PackageLimit"`

	// 源站流量限制，支持单位有bps、Kbps、Mbps、Gbps，支持范围1 bps-10000 Gbps。"0 bps"或其他单位数值为0，即不限流。""为不更新。
	FluxLimit *string `json:"FluxLimit,omitempty" name:"FluxLimit"`
}

type DDoSStatusInfo struct {
	// 策略等级，取值有:
	// <li>low ：宽松 ；</li>
	// <li>middle ：适中 ；</li>
	// <li>high : 严格。 </li>
	PlyLevel *string `json:"PlyLevel,omitempty" name:"PlyLevel"`
}

type DefaultServerCertInfo struct {
	// 服务器证书 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 证书备注名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 证书类型，取值有：
	// <li>default: 默认证书;</li>
	// <li>upload:用户上传;</li>
	// <li>managed:腾讯云托管。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 证书过期时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 证书生效时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EffectiveTime *string `json:"EffectiveTime,omitempty" name:"EffectiveTime"`

	// 证书公用名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommonName *string `json:"CommonName,omitempty" name:"CommonName"`

	// 证书SAN域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubjectAltName []*string `json:"SubjectAltName,omitempty" name:"SubjectAltName"`

	// 部署状态，取值有：
	// <li>processing: 部署中;</li>
	// <li>deployed: 已部署。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// Status为失败时,此字段返回失败原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`
}

// Predefined struct for user
type DeleteApplicationProxyRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

type DeleteApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID。
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

// Predefined struct for user
type DeleteApplicationProxyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteApplicationProxyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationProxyResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteApplicationProxyRuleRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type DeleteApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 规则ID。
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

// Predefined struct for user
type DeleteApplicationProxyRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteApplicationProxyRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationProxyRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteDnsRecordsRequestParams struct {
	// 待删除记录所属站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 待删除记录 ID。
	DnsRecordIds []*string `json:"DnsRecordIds,omitempty" name:"DnsRecordIds"`
}

type DeleteDnsRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 待删除记录所属站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 待删除记录 ID。
	DnsRecordIds []*string `json:"DnsRecordIds,omitempty" name:"DnsRecordIds"`
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
	delete(f, "DnsRecordIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDnsRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDnsRecordsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDnsRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDnsRecordsResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLoadBalancingRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 负载均衡ID。
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`
}

type DeleteLoadBalancingRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 负载均衡ID。
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

// Predefined struct for user
type DeleteLoadBalancingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLoadBalancingResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoadBalancingResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLogTopicTaskRequestParams struct {
	// 待删除的推送任务ID。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 推送任务所属日志集地域。
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`
}

type DeleteLogTopicTaskRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的推送任务ID。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 推送任务所属日志集地域。
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`
}

func (r *DeleteLogTopicTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogTopicTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "LogSetRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLogTopicTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogTopicTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLogTopicTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLogTopicTaskResponseParams `json:"Response"`
}

func (r *DeleteLogTopicTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogTopicTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOriginGroupRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 源站组ID。
	OriginGroupId *string `json:"OriginGroupId,omitempty" name:"OriginGroupId"`
}

type DeleteOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 源站组ID。
	OriginGroupId *string `json:"OriginGroupId,omitempty" name:"OriginGroupId"`
}

func (r *DeleteOriginGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOriginGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "OriginGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOriginGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteOriginGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOriginGroupResponseParams `json:"Response"`
}

func (r *DeleteOriginGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOriginGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRulesRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 指定删除的规则 ID 列表。
	RuleIds []*string `json:"RuleIds,omitempty" name:"RuleIds"`
}

type DeleteRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 指定删除的规则 ID 列表。
	RuleIds []*string `json:"RuleIds,omitempty" name:"RuleIds"`
}

func (r *DeleteRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRulesResponseParams `json:"Response"`
}

func (r *DeleteRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteZoneRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DeleteZoneRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
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
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteZoneResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteZoneResponse struct {
	*tchttp.BaseResponse
	Response *DeleteZoneResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAddableEntityListRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 推送数据类型，取值有:
	// <li>domain：七层加速日志；</li>
	// <li>application：四层加速日志；</li>
	// <li>web-rateLiming：速率限制日志；</li>
	// <li>web-attack：web攻击防护日志；</li>
	// <li>web-rule：自定义规则日志；</li>
	// <li>web-bot：Bot管理日志。</li>
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`
}

type DescribeAddableEntityListRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 推送数据类型，取值有:
	// <li>domain：七层加速日志；</li>
	// <li>application：四层加速日志；</li>
	// <li>web-rateLiming：速率限制日志；</li>
	// <li>web-attack：web攻击防护日志；</li>
	// <li>web-rule：自定义规则日志；</li>
	// <li>web-bot：Bot管理日志。</li>
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`
}

func (r *DescribeAddableEntityListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddableEntityListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "EntityType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddableEntityListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddableEntityListResponseParams struct {
	// 查询结果的总条数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 可添加的实体列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EntityList []*string `json:"EntityList,omitempty" name:"EntityList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAddableEntityListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAddableEntityListResponseParams `json:"Response"`
}

func (r *DescribeAddableEntityListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddableEntityListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationProxiesRequestParams struct {
	// 分页查询偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：<li>proxy-id<br>   按照【<strong>代理ID</strong>】进行过滤。代理ID形如：proxy-ev2sawbwfd。<br>   类型：String<br>   必选：否<li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-vawer2vadg。<br>   类型：String<br>   必选：否
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeApplicationProxiesRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：<li>proxy-id<br>   按照【<strong>代理ID</strong>】进行过滤。代理ID形如：proxy-ev2sawbwfd。<br>   类型：String<br>   必选：否<li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-vawer2vadg。<br>   类型：String<br>   必选：否
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApplicationProxiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationProxiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationProxiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationProxiesResponseParams struct {
	// 应用代理列表。
	ApplicationProxies []*ApplicationProxy `json:"ApplicationProxies,omitempty" name:"ApplicationProxies"`

	// 记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationProxiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationProxiesResponseParams `json:"Response"`
}

func (r *DescribeApplicationProxiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationProxiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailablePlansRequestParams struct {

}

type DescribeAvailablePlansRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAvailablePlansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailablePlansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAvailablePlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailablePlansResponseParams struct {
	// 当前账户可购买套餐类型及相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlanInfo []*PlanInfo `json:"PlanInfo,omitempty" name:"PlanInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAvailablePlansResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAvailablePlansResponseParams `json:"Response"`
}

func (r *DescribeAvailablePlansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailablePlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingDataRequestParams struct {
	// 起始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 时间粒度, 支持指定以下几种粒度：
	// <ul>
	// <li>min：1分钟粒度；</li>
	// <li>5min：5分钟粒度；</li>
	// <li>hour：1小时粒度；</li>
	// <li>day：天粒度；</li>
	// </ul>
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 指标名,支持:
	// <ul>
	// <li>acc_flux: 内容加速流量用量；</li>
	// <li>quic_request: QUIC 请求数用量；</li>
	// <li>sec_flux: 安全流量用量；</li>
	// <li>sec_request_clean: 安全干净流量请求数；</li>
	// </ul>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 筛选条件. 支持:
	// <ul>
	// <li>zone: 站点级数据；</li>
	// <li>plan: 套餐级数据；</li>
	// <li>service: l4 / l7分别筛选四七层数据；</li>
	// <li>tagKey: 标签Key；</li>
	// <li>tagValue: 标签Value。</li>
	// </ul>
	Filters []*BillingDataFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeBillingDataRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 时间粒度, 支持指定以下几种粒度：
	// <ul>
	// <li>min：1分钟粒度；</li>
	// <li>5min：5分钟粒度；</li>
	// <li>hour：1小时粒度；</li>
	// <li>day：天粒度；</li>
	// </ul>
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 指标名,支持:
	// <ul>
	// <li>acc_flux: 内容加速流量用量；</li>
	// <li>quic_request: QUIC 请求数用量；</li>
	// <li>sec_flux: 安全流量用量；</li>
	// <li>sec_request_clean: 安全干净流量请求数；</li>
	// </ul>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 筛选条件. 支持:
	// <ul>
	// <li>zone: 站点级数据；</li>
	// <li>plan: 套餐级数据；</li>
	// <li>service: l4 / l7分别筛选四七层数据；</li>
	// <li>tagKey: 标签Key；</li>
	// <li>tagValue: 标签Value。</li>
	// </ul>
	Filters []*BillingDataFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeBillingDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Interval")
	delete(f, "MetricName")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingDataResponseParams struct {
	// 统计曲线数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DnsData `json:"Data,omitempty" name:"Data"`

	// 时间粒度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBillingDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillingDataResponseParams `json:"Response"`
}

func (r *DescribeBillingDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotClientIpListRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 子域名列表，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询时间粒度，取值有：
	// <li>min ：1分钟；</li>
	// <li>5min ：5分钟；</li>
	// <li>hour ：1小时；</li>
	// <li>day ：1天。 </li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，key可选的值有：
	// <li>action: 执行动作。 </li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据归属地区，取值有：
	// <li>overseas ：全球（除中国大陆地区）数据； </li>
	// <li>mainland ：中国大陆地区数据。 </li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeBotClientIpListRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 子域名列表，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询时间粒度，取值有：
	// <li>min ：1分钟；</li>
	// <li>5min ：5分钟；</li>
	// <li>hour ：1小时；</li>
	// <li>day ：1天。 </li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，key可选的值有：
	// <li>action: 执行动作。 </li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据归属地区，取值有：
	// <li>overseas ：全球（除中国大陆地区）数据； </li>
	// <li>mainland ：中国大陆地区数据。 </li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeBotClientIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotClientIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Interval")
	delete(f, "QueryCondition")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBotClientIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotClientIpListResponseParams struct {
	// 客户端Ip相关数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*SecClientIp `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBotClientIpListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBotClientIpListResponseParams `json:"Response"`
}

func (r *DescribeBotClientIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotClientIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>bot_interceptNum ：bot拦截次数 ；</li>
	// <li>bot_noneRequestNum ：未分类bot请求次数 ；</li>
	// <li> bot_maliciousRequestNum：恶意bot请求次数 ；</li>
	// <li>bot_suspectedRequestNum ：疑似bot请求次数 ；</li>
	// <li>bot_friendlyRequestNum ：友好bot请求次数 ；</li>
	// <li>bot_normalRequestNum ：正常bot请求次数 。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 查询的子域名列表，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 站点列表，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，key可选的值有：
	// <li>action：执行动作 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据； </li>
	// <li>mainland：中国大陆地区数据。 </li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeBotDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>bot_interceptNum ：bot拦截次数 ；</li>
	// <li>bot_noneRequestNum ：未分类bot请求次数 ；</li>
	// <li> bot_maliciousRequestNum：恶意bot请求次数 ；</li>
	// <li>bot_suspectedRequestNum ：疑似bot请求次数 ；</li>
	// <li>bot_friendlyRequestNum ：友好bot请求次数 ；</li>
	// <li>bot_normalRequestNum ：正常bot请求次数 。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 查询的子域名列表，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 站点列表，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，key可选的值有：
	// <li>action：执行动作 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据； </li>
	// <li>mainland：中国大陆地区数据。 </li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeBotDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "Domains")
	delete(f, "ZoneIds")
	delete(f, "Interval")
	delete(f, "QueryCondition")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBotDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotDataResponseParams struct {
	// Bot攻击的数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*SecEntry `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBotDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBotDataResponseParams `json:"Response"`
}

func (r *DescribeBotDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotHitRuleDetailRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 子域名列表，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，key可选的值有：
	// <li>action: 执行动作。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据； </li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeBotHitRuleDetailRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 子域名列表，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，key可选的值有：
	// <li>action: 执行动作。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据； </li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeBotHitRuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotHitRuleDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Interval")
	delete(f, "QueryCondition")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBotHitRuleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotHitRuleDetailResponseParams struct {
	// 命中规则列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*SecHitRuleInfo `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBotHitRuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBotHitRuleDetailResponseParams `json:"Response"`
}

func (r *DescribeBotHitRuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotHitRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotLogRequestParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询的站点集合，不填默认查询所有站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 查询的域名集合，不填默认查询所有子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 筛选条件，当前key的可选条件有：
	// <li>action：执行动作（处置方式）；</li>
	// <li>sipCountryCode：ip所在国家；</li>
	// <li>attackIp：攻击ip；</li>
	// <li>ruleId：规则id；</li>
	// <li>eventId：事件id；</li>
	// <li>ua：用户代理；</li>
	// <li>requestMethod：请求方法；</li>
	// <li>uri：统一资源标识符 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeBotLogRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询的站点集合，不填默认查询所有站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 查询的域名集合，不填默认查询所有子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 筛选条件，当前key的可选条件有：
	// <li>action：执行动作（处置方式）；</li>
	// <li>sipCountryCode：ip所在国家；</li>
	// <li>attackIp：攻击ip；</li>
	// <li>ruleId：规则id；</li>
	// <li>eventId：事件id；</li>
	// <li>ua：用户代理；</li>
	// <li>requestMethod：请求方法；</li>
	// <li>uri：统一资源标识符 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeBotLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "QueryCondition")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBotLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotLogResponseParams struct {
	// Bot攻击数据信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*BotLog `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBotLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBotLogResponseParams `json:"Response"`
}

func (r *DescribeBotLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotManagedRulesRequestParams struct {
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 规则类型，取值有：
	// <li> idcid；</li>
	// <li>sipbot；</li>
	// <li>uabot。</li>传空或不传，即全部类型。
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
}

type DescribeBotManagedRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 规则类型，取值有：
	// <li> idcid；</li>
	// <li>sipbot；</li>
	// <li>uabot。</li>传空或不传，即全部类型。
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
}

func (r *DescribeBotManagedRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotManagedRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entity")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "RuleType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBotManagedRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotManagedRulesResponseParams struct {
	// 本次返回的规则数。
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Bot规则。
	BotManagedRuleDetails []*BotManagedRuleDetail `json:"BotManagedRuleDetails,omitempty" name:"BotManagedRuleDetails"`

	// 总规则数。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBotManagedRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBotManagedRulesResponseParams `json:"Response"`
}

func (r *DescribeBotManagedRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotManagedRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotTopDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>bot_requestNum_labelType：请求次数标签类型分布排行；</li>
	// <li>bot_requestNum_url：请求次数url分布排行；</li>
	// <li>bot_cipRequestNum_region：请求次数区域客户端ip分布排行。</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询前多少个数据，最多值为1000，不填默认默认为：10， 表示查询前top 10的数据。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，key可选的值有：
	// <li>action：执行动作 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeBotTopDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>bot_requestNum_labelType：请求次数标签类型分布排行；</li>
	// <li>bot_requestNum_url：请求次数url分布排行；</li>
	// <li>bot_cipRequestNum_region：请求次数区域客户端ip分布排行。</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询前多少个数据，最多值为1000，不填默认默认为：10， 表示查询前top 10的数据。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，key可选的值有：
	// <li>action：执行动作 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeBotTopDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotTopDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Limit")
	delete(f, "Interval")
	delete(f, "QueryCondition")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBotTopDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotTopDataResponseParams struct {
	// Bot攻击TopN数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TopEntry `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBotTopDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBotTopDataResponseParams `json:"Response"`
}

func (r *DescribeBotTopDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotTopDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientRuleListRequestParams struct {
	// 查询的站点ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 查询的子域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 规则类型，取值有：
	// <li>acl：自定义规则；</li>
	// <li>rate：限速规则。</li>不填表示查询所有规则。
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// 规则ID。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 客户端IP。
	SourceClientIp *string `json:"SourceClientIp,omitempty" name:"SourceClientIp"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeClientRuleListRequest struct {
	*tchttp.BaseRequest
	
	// 查询的站点ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 查询的子域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 规则类型，取值有：
	// <li>acl：自定义规则；</li>
	// <li>rate：限速规则。</li>不填表示查询所有规则。
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// 规则ID。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 客户端IP。
	SourceClientIp *string `json:"SourceClientIp,omitempty" name:"SourceClientIp"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeClientRuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientRuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Domain")
	delete(f, "RuleType")
	delete(f, "RuleId")
	delete(f, "SourceClientIp")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClientRuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientRuleListResponseParams struct {
	// 封禁客户端数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ClientRule `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClientRuleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClientRuleListResponseParams `json:"Response"`
}

func (r *DescribeClientRuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContentQuotaRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeContentQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeContentQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContentQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContentQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContentQuotaResponseParams struct {
	// 刷新相关配额。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PurgeQuota []*Quota `json:"PurgeQuota,omitempty" name:"PurgeQuota"`

	// 预热相关配额。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrefetchQuota []*Quota `json:"PrefetchQuota,omitempty" name:"PrefetchQuota"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeContentQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeContentQuotaResponseParams `json:"Response"`
}

func (r *DescribeContentQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContentQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>ddos_attackMaxBandwidth：攻击带宽峰值；</li>
	// <li>ddos_attackMaxPackageRate：攻击包速率峰值 ；</li>
	// <li>ddos_attackBandwidth：攻击带宽曲线；</li>
	// <li>ddos_attackPackageRate：攻击包速率曲线。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 端口号。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 攻击类型，取值有：
	// <li>flood：泛洪攻击；</li>
	// <li>icmpFlood：icmp泛洪攻击；</li>
	// <li>all：全部攻击类型。</li>不填默认为all，表示查询全部攻击类型。
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// DDOS策略组id列表，不填默认选择全部策略id。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 协议类型，取值有：
	// <li>tcp：tcp协议；</li>
	// <li>udp：udp协议；</li>
	// <li>all：全部协议。</li>不填默认为all，表示获取全部协议类型。
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeDDoSAttackDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>ddos_attackMaxBandwidth：攻击带宽峰值；</li>
	// <li>ddos_attackMaxPackageRate：攻击包速率峰值 ；</li>
	// <li>ddos_attackBandwidth：攻击带宽曲线；</li>
	// <li>ddos_attackPackageRate：攻击包速率曲线。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 端口号。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 攻击类型，取值有：
	// <li>flood：泛洪攻击；</li>
	// <li>icmpFlood：icmp泛洪攻击；</li>
	// <li>all：全部攻击类型。</li>不填默认为all，表示查询全部攻击类型。
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// DDOS策略组id列表，不填默认选择全部策略id。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 协议类型，取值有：
	// <li>tcp：tcp协议；</li>
	// <li>udp：udp协议；</li>
	// <li>all：全部协议。</li>不填默认为all，表示获取全部协议类型。
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeDDoSAttackDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "Port")
	delete(f, "AttackType")
	delete(f, "ZoneIds")
	delete(f, "PolicyIds")
	delete(f, "ProtocolType")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackDataResponseParams struct {
	// DDoS攻击数据内容列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*SecEntry `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSAttackDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSAttackDataResponseParams `json:"Response"`
}

func (r *DescribeDDoSAttackDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackEventDetailRequestParams struct {
	// 查询的事件ID。
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeDDoSAttackEventDetailRequest struct {
	*tchttp.BaseRequest
	
	// 查询的事件ID。
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeDDoSAttackEventDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackEventDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackEventDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackEventDetailResponseParams struct {
	// DDoS攻击事件详情。
	Data *DDoSAttackEventDetailData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSAttackEventDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSAttackEventDetailResponseParams `json:"Response"`
}

func (r *DescribeDDoSAttackEventDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackEventRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 协议类型，取值有：
	// <li>tcp：tcp协议；</li>
	// <li>udp：udp协议；</li>
	// <li>all：全部协议。</li>默认为: all，表示获取全部协议类型。
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// ddos策略组集合，不填默认选择全部策略。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 是否展示详细信息。
	ShowDetail *bool `json:"ShowDetail,omitempty" name:"ShowDetail"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeDDoSAttackEventRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 协议类型，取值有：
	// <li>tcp：tcp协议；</li>
	// <li>udp：udp协议；</li>
	// <li>all：全部协议。</li>默认为: all，表示获取全部协议类型。
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// ddos策略组集合，不填默认选择全部策略。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 是否展示详细信息。
	ShowDetail *bool `json:"ShowDetail,omitempty" name:"ShowDetail"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeDDoSAttackEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ProtocolType")
	delete(f, "PolicyIds")
	delete(f, "ZoneIds")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ShowDetail")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackEventResponseParams struct {
	// DDOS攻击事件数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DDoSAttackEvent `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSAttackEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSAttackEventResponseParams `json:"Response"`
}

func (r *DescribeDDoSAttackEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackSourceEventRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 协议类型，取值有：
	// <li>tcp：tcp协议；</li>
	// <li>udp：udp协议；</li>
	// <li>all：全部协议。</li>不填默认为: all，表示获取全部协议类型。
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// DDoS策略组集合，不填默认选择全部策略。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeDDoSAttackSourceEventRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 协议类型，取值有：
	// <li>tcp：tcp协议；</li>
	// <li>udp：udp协议；</li>
	// <li>all：全部协议。</li>不填默认为: all，表示获取全部协议类型。
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// DDoS策略组集合，不填默认选择全部策略。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeDDoSAttackSourceEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackSourceEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ProtocolType")
	delete(f, "PolicyIds")
	delete(f, "ZoneIds")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackSourceEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackSourceEventResponseParams struct {
	// DDoS攻击源数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DDoSAttackSourceEvent `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSAttackSourceEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSAttackSourceEventResponseParams `json:"Response"`
}

func (r *DescribeDDoSAttackSourceEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackSourceEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackTopDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询的统计指标，取值有：
	// <li>ddos_attackFlux_protocol：攻击总流量协议类型分布排行；</li>
	// <li>ddos_attackPackageNum_protocol：攻击总包量协议类型分布排行；</li>
	// <li>ddos_attackNum_attackType：攻击总次数攻击类型分布排行；</li>
	// <li>ddos_attackNum_sregion：攻击总次数攻击源地区分布排行；</li>
	// <li>ddos_attackFlux_sip：攻击总流量攻击源ip分布排行；</li>
	// <li>ddos_attackFlux_sregion：攻击总流量攻击源地区分布排行。</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 站点ID集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// DDoS策略组ID集合，不填默认选择全部策略ID。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 攻击类型，取值有：
	// <li>flood：洪泛攻击；</li>
	// <li>icmpFlood：icmp洪泛攻击；</li>
	// <li>all：所有的攻击类型。</li>不填默认为all，表示查询全部攻击类型。
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 协议类型，取值有：
	// <li>tcp：tcp协议；</li>
	// <li>udp：udp协议；</li>
	// <li>all：所有的协议类型。</li>不填默认为all，表示查询所有协议。
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 端口号。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 查询前多少个数据，不填默认默认为10， 表示查询前top 10的数据。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeDDoSAttackTopDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询的统计指标，取值有：
	// <li>ddos_attackFlux_protocol：攻击总流量协议类型分布排行；</li>
	// <li>ddos_attackPackageNum_protocol：攻击总包量协议类型分布排行；</li>
	// <li>ddos_attackNum_attackType：攻击总次数攻击类型分布排行；</li>
	// <li>ddos_attackNum_sregion：攻击总次数攻击源地区分布排行；</li>
	// <li>ddos_attackFlux_sip：攻击总流量攻击源ip分布排行；</li>
	// <li>ddos_attackFlux_sregion：攻击总流量攻击源地区分布排行。</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 站点ID集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// DDoS策略组ID集合，不填默认选择全部策略ID。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 攻击类型，取值有：
	// <li>flood：洪泛攻击；</li>
	// <li>icmpFlood：icmp洪泛攻击；</li>
	// <li>all：所有的攻击类型。</li>不填默认为all，表示查询全部攻击类型。
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 协议类型，取值有：
	// <li>tcp：tcp协议；</li>
	// <li>udp：udp协议；</li>
	// <li>all：所有的协议类型。</li>不填默认为all，表示查询所有协议。
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 端口号。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 查询前多少个数据，不填默认默认为10， 表示查询前top 10的数据。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeDDoSAttackTopDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackTopDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "ZoneIds")
	delete(f, "PolicyIds")
	delete(f, "AttackType")
	delete(f, "ProtocolType")
	delete(f, "Port")
	delete(f, "Limit")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackTopDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackTopDataResponseParams struct {
	// DDoS攻击Top数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TopEntry `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSAttackTopDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSAttackTopDataResponseParams `json:"Response"`
}

func (r *DescribeDDoSAttackTopDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackTopDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSBlockListRequestParams struct {
	// 攻击事件起始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 攻击事件列表。
	EventIds []*string `json:"EventIds,omitempty" name:"EventIds"`

	// 站点列表，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 策略列表，不填默认选择全部策略。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeDDoSBlockListRequest struct {
	*tchttp.BaseRequest
	
	// 攻击事件起始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 攻击事件列表。
	EventIds []*string `json:"EventIds,omitempty" name:"EventIds"`

	// 站点列表，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 策略列表，不填默认选择全部策略。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeDDoSBlockListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSBlockListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EventIds")
	delete(f, "ZoneIds")
	delete(f, "PolicyIds")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSBlockListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSBlockListResponseParams struct {
	// 封禁解封信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DDoSBlockData `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSBlockListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSBlockListResponseParams `json:"Response"`
}

func (r *DescribeDDoSBlockListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSBlockListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSMajorAttackEventRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 站点id列表，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// ddos策略组id集合，不填默认选择全部策略id。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 协议类型，可选的协议类型有：
	// <li>tcp：tcp协议；</li>
	// <li>udp：udp协议；</li>
	// <li>all：全部协议。</li>不填默认为all, 表示获取全部协议类型。
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeDDoSMajorAttackEventRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 站点id列表，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// ddos策略组id集合，不填默认选择全部策略id。
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// 协议类型，可选的协议类型有：
	// <li>tcp：tcp协议；</li>
	// <li>udp：udp协议；</li>
	// <li>all：全部协议。</li>不填默认为all, 表示获取全部协议类型。
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeDDoSMajorAttackEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSMajorAttackEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ZoneIds")
	delete(f, "PolicyIds")
	delete(f, "ProtocolType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSMajorAttackEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSMajorAttackEventResponseParams struct {
	// DDos主攻击事件数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DDoSMajorAttackEvent `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSMajorAttackEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSMajorAttackEventResponseParams `json:"Response"`
}

func (r *DescribeDDoSMajorAttackEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSMajorAttackEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSPolicyRequestParams struct {
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 策略Id。
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

type DescribeDDoSPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 策略Id。
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
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
	delete(f, "ZoneId")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSPolicyResponseParams struct {
	// DDoS防护配置。
	DDoSRule *DDoSRule `json:"DDoSRule,omitempty" name:"DDoSRule"`

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
type DescribeDefaultCertificatesRequestParams struct {
	// 过滤条件，Filters.Values的上限为5。详细的过滤条件如下：
	// <li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-xxx。<br>   类型：String<br>   必选：是
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDefaultCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，Filters.Values的上限为5。详细的过滤条件如下：
	// <li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-xxx。<br>   类型：String<br>   必选：是
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDefaultCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultCertificatesResponseParams struct {
	// 证书总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 默认证书列表。
	DefaultServerCertInfo []*DefaultServerCertInfo `json:"DefaultServerCertInfo,omitempty" name:"DefaultServerCertInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDefaultCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDefaultCertificatesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDnsDataRequestParams struct {
	// 起始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone<br>   按照【<strong>站点名称</strong>】进行过滤。站点名称形如：tencent.com<br>   类型：String<br>   必选：否，仅支持填写一个站点
	// <li>host<br>   按照【<strong>域名</strong>】进行过滤。域名形如：test.tencent.com<br>   类型：String<br>   必选：否，仅支持填写一个域名
	// <li>type<br>   按照【<strong>DNS解析类型</strong>】进行过滤<br>   类型：String<br>   必选：否<br>   可选项：<br>   A：A记录<br>   AAAA：AAAA记录<br>   CNAME：CNAME记录<br>   MX：MX记录<br>   TXT：TXT记录<br>   NS：NS记录<br>   SRV：SRV记录<br>   CAA：CAA记录
	// <li>code<br>   按照【<strong>DNS解析状态码</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   NoError：成功<br>   NXDomain：请求域不存在<br>   NotImp：不支持的请求类型<br>   Refused：域名服务器因为策略的原因拒绝执行请求
	// <li>area<br>   按照【<strong>DNS解析地域</strong>】进行过滤。<br>   类型：String<br>   必选：否。<br>   可选项：<br>   亚洲：Asia<br>   欧洲：Europe<br>   非洲：Africa<br>   大洋洲：Oceania<br>   美洲：Americas
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 时间粒度，取值有：
	// <li>min：1分钟粒度；</li>
	// <li>5min：5分钟粒度；</li>
	// <li>hour：1小时粒度；</li>
	// <li>day：天粒度。</li>不填写，默认值为：min。
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

type DescribeDnsDataRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone<br>   按照【<strong>站点名称</strong>】进行过滤。站点名称形如：tencent.com<br>   类型：String<br>   必选：否，仅支持填写一个站点
	// <li>host<br>   按照【<strong>域名</strong>】进行过滤。域名形如：test.tencent.com<br>   类型：String<br>   必选：否，仅支持填写一个域名
	// <li>type<br>   按照【<strong>DNS解析类型</strong>】进行过滤<br>   类型：String<br>   必选：否<br>   可选项：<br>   A：A记录<br>   AAAA：AAAA记录<br>   CNAME：CNAME记录<br>   MX：MX记录<br>   TXT：TXT记录<br>   NS：NS记录<br>   SRV：SRV记录<br>   CAA：CAA记录
	// <li>code<br>   按照【<strong>DNS解析状态码</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   NoError：成功<br>   NXDomain：请求域不存在<br>   NotImp：不支持的请求类型<br>   Refused：域名服务器因为策略的原因拒绝执行请求
	// <li>area<br>   按照【<strong>DNS解析地域</strong>】进行过滤。<br>   类型：String<br>   必选：否。<br>   可选项：<br>   亚洲：Asia<br>   欧洲：Europe<br>   非洲：Africa<br>   大洋洲：Oceania<br>   美洲：Americas
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 时间粒度，取值有：
	// <li>min：1分钟粒度；</li>
	// <li>5min：5分钟粒度；</li>
	// <li>hour：1小时粒度；</li>
	// <li>day：天粒度。</li>不填写，默认值为：min。
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

// Predefined struct for user
type DescribeDnsDataResponseParams struct {
	// 统计数据。
	Data []*DnsData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDnsDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDnsDataResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDnsRecordsRequestParams struct {
	// DNS记录所属站点ID。不填写该参数默认返回所有站点下的记录。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>record-id<br>   按照【<strong>DNS记录id</strong>】进行过滤。DNS记录ID形如：record-1a8df68z。<br>   类型：String<br>   必选：否
	// <li>record-name<br>   按照【<strong>DNS记录名称</strong>】进行过滤。<br>   类型：String<br>   必选：否
	// <li>record-type<br>   按照【<strong>DNS记录类型</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   A：将域名指向一个外网 IPv4 地址，如 8.8.8.8<br>   AAAA：将域名指向一个外网 IPv6 地址<br>   CNAME：将域名指向另一个域名，再由该域名解析出最终 IP 地址<br>   TXT：对域名进行标识和说明，常用于域名验证和 SPF 记录（反垃圾邮件）<br>   NS：如果需要将子域名交给其他 DNS 服务商解析，则需要添加 NS 记录。根域名无法添加 NS 记录<br>   CAA：指定可为本站点颁发证书的 CA<br>   SRV：标识某台服务器使用了某个服务，常见于微软系统的目录管理<br>   MX：指定收件人邮件服务器。
	// <li>mode<br>   按照【<strong>代理模式</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   dns_only：仅DNS解析<br>   proxied：代理加速
	// <li>ttl<br>   按照【<strong>解析生效时间</strong>】进行过滤。<br>   类型：string<br>   必选：否
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`

	// 列表排序方式，取值有：
	// <li>asc：升序排列；</li>
	// <li>desc：降序排列。</li>默认值为asc。
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// 匹配方式，取值有：
	// <li>all：返回匹配所有查询条件的记录；</li>
	// <li>any：返回匹配任意一个查询条件的记录。</li>默认值为all。
	Match *string `json:"Match,omitempty" name:"Match"`

	// 分页查询限制数目，默认值：20，上限：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序依据，取值有：
	// <li>content：DNS记录内容；</li>
	// <li>created_on：DNS记录创建时间；</li>
	// <li>mode：代理模式；</li>
	// <li>record-name：DNS记录名称；</li>
	// <li>ttl：解析记录生效时间；</li>
	// <li>record-type：DNS记录类型。</li>默认根据record-type, recrod-name属性组合排序。
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeDnsRecordsRequest struct {
	*tchttp.BaseRequest
	
	// DNS记录所属站点ID。不填写该参数默认返回所有站点下的记录。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>record-id<br>   按照【<strong>DNS记录id</strong>】进行过滤。DNS记录ID形如：record-1a8df68z。<br>   类型：String<br>   必选：否
	// <li>record-name<br>   按照【<strong>DNS记录名称</strong>】进行过滤。<br>   类型：String<br>   必选：否
	// <li>record-type<br>   按照【<strong>DNS记录类型</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   A：将域名指向一个外网 IPv4 地址，如 8.8.8.8<br>   AAAA：将域名指向一个外网 IPv6 地址<br>   CNAME：将域名指向另一个域名，再由该域名解析出最终 IP 地址<br>   TXT：对域名进行标识和说明，常用于域名验证和 SPF 记录（反垃圾邮件）<br>   NS：如果需要将子域名交给其他 DNS 服务商解析，则需要添加 NS 记录。根域名无法添加 NS 记录<br>   CAA：指定可为本站点颁发证书的 CA<br>   SRV：标识某台服务器使用了某个服务，常见于微软系统的目录管理<br>   MX：指定收件人邮件服务器。
	// <li>mode<br>   按照【<strong>代理模式</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   dns_only：仅DNS解析<br>   proxied：代理加速
	// <li>ttl<br>   按照【<strong>解析生效时间</strong>】进行过滤。<br>   类型：string<br>   必选：否
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`

	// 列表排序方式，取值有：
	// <li>asc：升序排列；</li>
	// <li>desc：降序排列。</li>默认值为asc。
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// 匹配方式，取值有：
	// <li>all：返回匹配所有查询条件的记录；</li>
	// <li>any：返回匹配任意一个查询条件的记录。</li>默认值为all。
	Match *string `json:"Match,omitempty" name:"Match"`

	// 分页查询限制数目，默认值：20，上限：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序依据，取值有：
	// <li>content：DNS记录内容；</li>
	// <li>created_on：DNS记录创建时间；</li>
	// <li>mode：代理模式；</li>
	// <li>record-name：DNS记录名称；</li>
	// <li>ttl：解析记录生效时间；</li>
	// <li>record-type：DNS记录类型。</li>默认根据record-type, recrod-name属性组合排序。
	Order *string `json:"Order,omitempty" name:"Order"`
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
	delete(f, "ZoneId")
	delete(f, "Filters")
	delete(f, "Direction")
	delete(f, "Match")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDnsRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDnsRecordsResponseParams struct {
	// DNS记录总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// DNS 记录列表。
	DnsRecords []*DnsRecord `json:"DnsRecords,omitempty" name:"DnsRecords"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDnsRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDnsRecordsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDnssecRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeDnssecRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
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
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDnssecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDnssecResponseParams struct {
	// DNSSEC状态信息，取值有：
	// <li>enabled：开启；</li>
	// <li>disabled：关闭。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// DNSSEC相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DnssecInfo *DnssecInfo `json:"DnssecInfo,omitempty" name:"DnssecInfo"`

	// 站点信息更新时间。
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDnssecResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDnssecResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeHostCertificatesRequestParams struct {
	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-xxx。<br>   类型：String<br>   必选：是<li>host<br>   按照【<strong>域名名称</strong>】进行过滤。<br>   类型：String<br>   必选：否<li>cert-id<br>   按照【<strong>证书ID</strong>】进行过滤。<br>   类型：String<br>   必选：否<li>cert-alias<br>   按照【<strong>证书名称</strong>】进行过滤。<br>   类型：String<br>   必选：否<li>cert-type<br>   按照【<strong>证书类型</strong>】进行过滤。<br>   类型：String<br>   必选：否
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为 100，最大可设置为 1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序方式。详细排序条件如下：
	// <li>create-time：域名创建时间；</li>
	// <li>cert-expire-time：证书过期时间；</li>
	// <li>cert-deploy-time：证书部署时间。</li>
	Sort *Sort `json:"Sort,omitempty" name:"Sort"`
}

type DescribeHostCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-xxx。<br>   类型：String<br>   必选：是<li>host<br>   按照【<strong>域名名称</strong>】进行过滤。<br>   类型：String<br>   必选：否<li>cert-id<br>   按照【<strong>证书ID</strong>】进行过滤。<br>   类型：String<br>   必选：否<li>cert-alias<br>   按照【<strong>证书名称</strong>】进行过滤。<br>   类型：String<br>   必选：否<li>cert-type<br>   按照【<strong>证书类型</strong>】进行过滤。<br>   类型：String<br>   必选：否
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为 100，最大可设置为 1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序方式。详细排序条件如下：
	// <li>create-time：域名创建时间；</li>
	// <li>cert-expire-time：证书过期时间；</li>
	// <li>cert-deploy-time：证书部署时间。</li>
	Sort *Sort `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeHostCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostCertificatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostCertificatesResponseParams struct {
	// 总数，用于分页查询。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 域名证书配置列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostCertificates []*HostsCertificate `json:"HostCertificates,omitempty" name:"HostCertificates"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostCertificatesResponseParams `json:"Response"`
}

func (r *DescribeHostCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsSettingRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 分页查询偏移量。默认值： 0，最小值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值： 100，最大值：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>host<br>   按照【<strong>域名</strong>】进行过滤。<br>   类型：string<br>   必选：否
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeHostsSettingRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 分页查询偏移量。默认值： 0，最小值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值： 100，最大值：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>host<br>   按照【<strong>域名</strong>】进行过滤。<br>   类型：string<br>   必选：否
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostsSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsSettingResponseParams struct {
	// 域名列表。
	DetailHosts []*DetailHost `json:"DetailHosts,omitempty" name:"DetailHosts"`

	// 域名数量。
	TotalNumber *int64 `json:"TotalNumber,omitempty" name:"TotalNumber"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostsSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostsSettingResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeIdentificationsRequestParams struct {
	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-name<br>   按照【<strong>站点名称</strong>】进行过滤。<br>   类型：String<br>   必选：是
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeIdentificationsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-name<br>   按照【<strong>站点名称</strong>】进行过滤。<br>   类型：String<br>   必选：是
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeIdentificationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdentificationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIdentificationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdentificationsResponseParams struct {
	// 符合条件的站点个数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 站点验证信息列表。
	Identifications []*Identification `json:"Identifications,omitempty" name:"Identifications"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIdentificationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIdentificationsResponseParams `json:"Response"`
}

func (r *DescribeIdentificationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdentificationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancingRequestParams struct {
	// 分页查询偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为10，取值：1-1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-1a8df68z<br>   类型：String<br>   必选：否<br>   模糊查询：不支持
	// <li>load-balancing-id<br>   按照【<strong>负载均衡ID</strong>】进行过滤。负载均衡ID形如：lb-d21bfaf7-8d72-11ec-841d-00ff977fb3c8<br>   类型：String<br>   必选：否<br>   模糊查询：不支持
	// <li>host<br>   按照【<strong>负载均衡host</strong>】进行过滤。host形如：lb.tencent.com<br>   类型：String<br>   必选：否<br>   模糊查询：支持，模糊查询时仅支持一个host
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeLoadBalancingRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为10，取值：1-1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-1a8df68z<br>   类型：String<br>   必选：否<br>   模糊查询：不支持
	// <li>load-balancing-id<br>   按照【<strong>负载均衡ID</strong>】进行过滤。负载均衡ID形如：lb-d21bfaf7-8d72-11ec-841d-00ff977fb3c8<br>   类型：String<br>   必选：否<br>   模糊查询：不支持
	// <li>host<br>   按照【<strong>负载均衡host</strong>】进行过滤。host形如：lb.tencent.com<br>   类型：String<br>   必选：否<br>   模糊查询：支持，模糊查询时仅支持一个host
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
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
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancingResponseParams struct {
	// 记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 负载均衡信息。
	Data []*LoadBalancing `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLoadBalancingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancingResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLogSetsRequestParams struct {
	// 日志集所属的地域。
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`

	// 日志集ID。
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 日志集名称。
	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`
}

type DescribeLogSetsRequest struct {
	*tchttp.BaseRequest
	
	// 日志集所属的地域。
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`

	// 日志集ID。
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 日志集名称。
	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`
}

func (r *DescribeLogSetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogSetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogSetRegion")
	delete(f, "LogSetId")
	delete(f, "LogSetName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogSetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogSetsResponseParams struct {
	// 日志集列表数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogSetList []*LogSetInfo `json:"LogSetList,omitempty" name:"LogSetList"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLogSetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogSetsResponseParams `json:"Response"`
}

func (r *DescribeLogSetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogTopicTaskDetailRequestParams struct {
	// 推送任务ID。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeLogTopicTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 推送任务ID。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeLogTopicTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogTopicTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogTopicTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogTopicTaskDetailResponseParams struct {
	// 推送任务详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogTopicDetailInfo *LogTopicDetailInfo `json:"LogTopicDetailInfo,omitempty" name:"LogTopicDetailInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLogTopicTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogTopicTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeLogTopicTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogTopicTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogTopicTasksRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeLogTopicTasksRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeLogTopicTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogTopicTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogTopicTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogTopicTasksResponseParams struct {
	// 推送任务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicList []*ClsLogTopicInfo `json:"TopicList,omitempty" name:"TopicList"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLogTopicTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogTopicTasksResponseParams `json:"Response"`
}

func (r *DescribeLogTopicTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogTopicTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginGroupRequestParams struct {
	// 分页查询偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为10，取值：1-1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-20hzkd4rdmy0<br>   类型：String<br>   必选：否<br>   模糊查询：不支持<li>origin-group-id<br>   按照【<strong>源站组ID</strong>】进行过滤。源站组ID形如：origin-2ccgtb24-7dc5-46s2-9r3e-95825d53dwe3a<br>   类型：String<br>   必选：否<br>   模糊查询：不支持<li>origin-group-name<br>   按照【<strong>源站组名称</strong>】进行过滤<br>   类型：String<br>   必选：否<br>   模糊查询：支持。使用模糊查询时，仅支持填写一个源站组名称
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为10，取值：1-1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-20hzkd4rdmy0<br>   类型：String<br>   必选：否<br>   模糊查询：不支持<li>origin-group-id<br>   按照【<strong>源站组ID</strong>】进行过滤。源站组ID形如：origin-2ccgtb24-7dc5-46s2-9r3e-95825d53dwe3a<br>   类型：String<br>   必选：否<br>   模糊查询：不支持<li>origin-group-name<br>   按照【<strong>源站组名称</strong>】进行过滤<br>   类型：String<br>   必选：否<br>   模糊查询：支持。使用模糊查询时，仅支持填写一个源站组名称
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeOriginGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginGroupResponseParams struct {
	// 记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 源站组信息。
	OriginGroups []*OriginGroup `json:"OriginGroups,omitempty" name:"OriginGroups"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOriginGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOriginGroupResponseParams `json:"Response"`
}

func (r *DescribeOriginGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewL7DataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询的指标，取值有：
	// <li>l7Flow_outFlux: 访问流量；</li>
	// <li>l7Flow_request: 访问请求数；</li>
	// <li>l7Flow_outBandwidth: 访问带宽；</li>
	// <li>l7Flow_hit_outFlux: 缓存命中流量。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 查询的站点集合，不填默认查询所有站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 查询的域名集合，不填默认查询所有子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询的协议类型，取值有：
	// <li>http: http协议；</li>
	// <li>https: https协议；</li>
	// <li>http2: http2协议；</li>
	// <li>all:  所有协议。</li>不填默认为: all，表示查询所有协议。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户的地域智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`
}

type DescribeOverviewL7DataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询的指标，取值有：
	// <li>l7Flow_outFlux: 访问流量；</li>
	// <li>l7Flow_request: 访问请求数；</li>
	// <li>l7Flow_outBandwidth: 访问带宽；</li>
	// <li>l7Flow_hit_outFlux: 缓存命中流量。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 查询的站点集合，不填默认查询所有站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 查询的域名集合，不填默认查询所有子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询的协议类型，取值有：
	// <li>http: http协议；</li>
	// <li>https: https协议；</li>
	// <li>http2: http2协议；</li>
	// <li>all:  所有协议。</li>不填默认为: all，表示查询所有协议。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户的地域智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeOverviewL7DataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewL7DataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Protocol")
	delete(f, "Interval")
	delete(f, "Area")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOverviewL7DataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewL7DataResponseParams struct {
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 七层监控类时序流量数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TimingDataRecord `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOverviewL7DataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOverviewL7DataResponseParams `json:"Response"`
}

func (r *DescribeOverviewL7DataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewL7DataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrefetchTasksRequestParams struct {
	// 查询起始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认值：20，上限：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-id<br>   按照【<strong>站点 ID</strong>】进行过滤。zone-id形如：zone-1379afjk91u32h，暂不支持多值。<br>   类型：String<br>   必选：否。<br>   模糊查询：不支持。<li>job-id<br>   按照【<strong>任务ID</strong>】进行过滤。job-id形如：1379afjk91u32h，暂不支持多值。<br>   类型：String<br>   必选：否。<br>   模糊查询：不支持。<li>target<br>   按照【<strong>目标资源信息</strong>】进行过滤。target形如：http://www.qq.com/1.txt，暂不支持多值。<br>   类型：String<br>   必选：否。<br>   模糊查询：不支持。<li>domains<br>   按照【<strong>域名</strong>】进行过滤。domains形如：www.qq.com。<br>   类型：String<br>   必选：否。<br>   模糊查询：不支持。<li>statuses<br>   按照【<strong>任务状态</strong>】进行过滤。<br>   必选：否<br>   模糊查询：不支持。<br>   可选项：<br>   processing：处理中<br>   success：成功<br>   failed：失败<br>   timeout：超时
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribePrefetchTasksRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认值：20，上限：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-id<br>   按照【<strong>站点 ID</strong>】进行过滤。zone-id形如：zone-1379afjk91u32h，暂不支持多值。<br>   类型：String<br>   必选：否。<br>   模糊查询：不支持。<li>job-id<br>   按照【<strong>任务ID</strong>】进行过滤。job-id形如：1379afjk91u32h，暂不支持多值。<br>   类型：String<br>   必选：否。<br>   模糊查询：不支持。<li>target<br>   按照【<strong>目标资源信息</strong>】进行过滤。target形如：http://www.qq.com/1.txt，暂不支持多值。<br>   类型：String<br>   必选：否。<br>   模糊查询：不支持。<li>domains<br>   按照【<strong>域名</strong>】进行过滤。domains形如：www.qq.com。<br>   类型：String<br>   必选：否。<br>   模糊查询：不支持。<li>statuses<br>   按照【<strong>任务状态</strong>】进行过滤。<br>   必选：否<br>   模糊查询：不支持。<br>   可选项：<br>   processing：处理中<br>   success：成功<br>   failed：失败<br>   timeout：超时
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
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
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrefetchTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrefetchTasksResponseParams struct {
	// 该查询条件总共条目数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务结果列表。
	Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrefetchTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrefetchTasksResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePurgeTasksRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 查询起始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页查询偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查限制数目，默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>job-id<br>   按照【<strong>任务ID</strong>】进行过滤。job-id形如：1379afjk91u32h，暂不支持多值。<br>   类型：String<br>   必选：否<br>   模糊查询：不支持。<li>target<br>   按照【<strong>目标资源信息</strong>】进行过滤。target形如：http://www.qq.com/1.txt，暂不支持多值。<br>   类型：String<br>   必选：否<br>   模糊查询：不支持。<li>domains<br>   按照【<strong>域名</strong>】进行过滤。domains形如：www.qq.com。<br>   类型：String<br>   必选：否<br>   模糊查询：不支持。<li>statuses<br>   按照【<strong>任务状态</strong>】进行过滤。<br>   必选：否<br>   模糊查询：不支持。<br>   可选项：<br>   processing：处理中<br>   success：成功<br>   failed：失败<br>   timeout：超时<li>type<br>   按照【<strong>清除缓存类型</strong>】进行过滤，暂不支持多值。<br>   类型：String<br>   必选：否<br>   模糊查询：不支持。<br>   可选项：<br>   purge_url：URL<br>   purge_prefix：前缀<br>   purge_all：全部缓存内容<br>   purge_host：Hostname
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribePurgeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 查询起始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页查询偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查限制数目，默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>job-id<br>   按照【<strong>任务ID</strong>】进行过滤。job-id形如：1379afjk91u32h，暂不支持多值。<br>   类型：String<br>   必选：否<br>   模糊查询：不支持。<li>target<br>   按照【<strong>目标资源信息</strong>】进行过滤。target形如：http://www.qq.com/1.txt，暂不支持多值。<br>   类型：String<br>   必选：否<br>   模糊查询：不支持。<li>domains<br>   按照【<strong>域名</strong>】进行过滤。domains形如：www.qq.com。<br>   类型：String<br>   必选：否<br>   模糊查询：不支持。<li>statuses<br>   按照【<strong>任务状态</strong>】进行过滤。<br>   必选：否<br>   模糊查询：不支持。<br>   可选项：<br>   processing：处理中<br>   success：成功<br>   failed：失败<br>   timeout：超时<li>type<br>   按照【<strong>清除缓存类型</strong>】进行过滤，暂不支持多值。<br>   类型：String<br>   必选：否<br>   模糊查询：不支持。<br>   可选项：<br>   purge_url：URL<br>   purge_prefix：前缀<br>   purge_all：全部缓存内容<br>   purge_host：Hostname
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
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
	delete(f, "ZoneId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePurgeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeTasksResponseParams struct {
	// 该查询条件总共条目数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务结果列表。
	Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePurgeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribePurgeTasksResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRateLimitIntelligenceRuleRequestParams struct {
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`
}

type DescribeRateLimitIntelligenceRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`
}

func (r *DescribeRateLimitIntelligenceRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRateLimitIntelligenceRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRateLimitIntelligenceRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRateLimitIntelligenceRuleResponseParams struct {
	// 速率限制智能规则。
	RateLimitIntelligenceRuleDetails []*RateLimitIntelligenceRuleDetail `json:"RateLimitIntelligenceRuleDetails,omitempty" name:"RateLimitIntelligenceRuleDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRateLimitIntelligenceRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRateLimitIntelligenceRuleResponseParams `json:"Response"`
}

func (r *DescribeRateLimitIntelligenceRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRateLimitIntelligenceRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>rule-id<br>   按照【<strong>规则ID</strong>】进行过滤。<br>   类型：string<br>   必选：否
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>rule-id<br>   按照【<strong>规则ID</strong>】进行过滤。<br>   类型：string<br>   必选：否
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	delete(f, "ZoneId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesResponseParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 规则列表，按规则执行顺序从先往后排序。
	RuleItems []*RuleItem `json:"RuleItems,omitempty" name:"RuleItems"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRulesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRulesSettingRequestParams struct {

}

type DescribeRulesSettingRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRulesSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRulesSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesSettingResponseParams struct {
	// 规则引擎可应用匹配请求的设置列表及其详细建议配置信息。
	Actions []*RulesSettingAction `json:"Actions,omitempty" name:"Actions"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRulesSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRulesSettingResponseParams `json:"Response"`
}

func (r *DescribeRulesSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupManagedRulesRequestParams struct {
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSecurityGroupManagedRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSecurityGroupManagedRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupManagedRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entity")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupManagedRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupManagedRulesResponseParams struct {
	// 本次返回的规则数。
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 总规则数。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 托管规则信息。
	WafGroupInfo *WafGroupInfo `json:"WafGroupInfo,omitempty" name:"WafGroupInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupManagedRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupManagedRulesResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupManagedRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupManagedRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyListRequestParams struct {
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeSecurityPolicyListRequest struct {
	*tchttp.BaseRequest
	
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeSecurityPolicyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPolicyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyListResponseParams struct {
	// 防护资源列表。
	SecurityEntities []*SecurityEntity `json:"SecurityEntities,omitempty" name:"SecurityEntities"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityPolicyListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityPolicyListResponseParams `json:"Response"`
}

func (r *DescribeSecurityPolicyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyRegionsRequestParams struct {
	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSecurityPolicyRegionsRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSecurityPolicyRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPolicyRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyRegionsResponseParams struct {
	// 总地域信息数。
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 地域信息。
	GeoIps []*GeoIp `json:"GeoIps,omitempty" name:"GeoIps"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityPolicyRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityPolicyRegionsResponseParams `json:"Response"`
}

func (r *DescribeSecurityPolicyRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyRequestParams struct {
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`
}

type DescribeSecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`
}

func (r *DescribeSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyResponseParams struct {
	// 安全配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityConfig *SecurityConfig `json:"SecurityConfig,omitempty" name:"SecurityConfig"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityPolicyResponseParams `json:"Response"`
}

func (r *DescribeSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPortraitRulesRequestParams struct {
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`
}

type DescribeSecurityPortraitRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`
}

func (r *DescribeSecurityPortraitRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPortraitRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPortraitRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPortraitRulesResponseParams struct {
	// 本次返回的规则数。
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Bot用户画像规则。
	PortraitManagedRuleDetails []*PortraitManagedRuleDetail `json:"PortraitManagedRuleDetails,omitempty" name:"PortraitManagedRuleDetails"`

	// 总规则数。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityPortraitRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityPortraitRulesResponseParams `json:"Response"`
}

func (r *DescribeSecurityPortraitRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPortraitRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityRuleIdRequestParams struct {
	// 规则Id数组。
	RuleIdList []*int64 `json:"RuleIdList,omitempty" name:"RuleIdList"`

	// 规则类型，取值有：
	// <li>waf：web托管规则；</li>
	// <li>acl：自定义规则；</li>
	// <li>rate：速率限制规则；</li>
	// <li>bot：Bot基础规则。</li>
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// 子域名/应用名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`
}

type DescribeSecurityRuleIdRequest struct {
	*tchttp.BaseRequest
	
	// 规则Id数组。
	RuleIdList []*int64 `json:"RuleIdList,omitempty" name:"RuleIdList"`

	// 规则类型，取值有：
	// <li>waf：web托管规则；</li>
	// <li>acl：自定义规则；</li>
	// <li>rate：速率限制规则；</li>
	// <li>bot：Bot基础规则。</li>
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// 子域名/应用名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`
}

func (r *DescribeSecurityRuleIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityRuleIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleIdList")
	delete(f, "RuleType")
	delete(f, "Entity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityRuleIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityRuleIdResponseParams struct {
	// 规则列表。
	WafGroupRules []*WafGroupRule `json:"WafGroupRules,omitempty" name:"WafGroupRules"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityRuleIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityRuleIdResponseParams `json:"Response"`
}

func (r *DescribeSecurityRuleIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityRuleIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSingleL7AnalysisDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询的指标，取值有:
	// <li> l7Flow_singleIpRequest：独立IP请求数。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 查询的站点集合，不填默认查询所有站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 筛选条件, key可选的值有：
	// <li>country：国家/地区；</li>
	// <li>domain：域名；</li>
	// <li>protocol：协议类型；</li>
	// <li>tagKey：标签Key；</li>
	// <li>tagValue；标签Value。</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天;。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeSingleL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询的指标，取值有:
	// <li> l7Flow_singleIpRequest：独立IP请求数。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 查询的站点集合，不填默认查询所有站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 筛选条件, key可选的值有：
	// <li>country：国家/地区；</li>
	// <li>domain：域名；</li>
	// <li>protocol：协议类型；</li>
	// <li>tagKey：标签Key；</li>
	// <li>tagValue；标签Value。</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天;。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeSingleL7AnalysisDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSingleL7AnalysisDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "Filters")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSingleL7AnalysisDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSingleL7AnalysisDataResponseParams struct {
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 单值流量数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*SingleDataRecord `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSingleL7AnalysisDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSingleL7AnalysisDataResponseParams `json:"Response"`
}

func (r *DescribeSingleL7AnalysisDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSingleL7AnalysisDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpeedTestingDetailsRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeSpeedTestingDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeSpeedTestingDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpeedTestingDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpeedTestingDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpeedTestingDetailsResponseParams struct {
	// 分地域拨测统计数据。
	SpeedTestingDetailData *SpeedTestingDetailData `json:"SpeedTestingDetailData,omitempty" name:"SpeedTestingDetailData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSpeedTestingDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpeedTestingDetailsResponseParams `json:"Response"`
}

func (r *DescribeSpeedTestingDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpeedTestingDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpeedTestingMetricDataRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeSpeedTestingMetricDataRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeSpeedTestingMetricDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpeedTestingMetricDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpeedTestingMetricDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpeedTestingMetricDataResponseParams struct {
	// 站点拨测维度数据。
	SpeedTestingMetricData *SpeedTestingMetricData `json:"SpeedTestingMetricData,omitempty" name:"SpeedTestingMetricData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSpeedTestingMetricDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpeedTestingMetricDataResponseParams `json:"Response"`
}

func (r *DescribeSpeedTestingMetricDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpeedTestingMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpeedTestingQuotaRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeSpeedTestingQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeSpeedTestingQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpeedTestingQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpeedTestingQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpeedTestingQuotaResponseParams struct {
	// 配额数据。
	SpeedTestingQuota *SpeedTestingQuota `json:"SpeedTestingQuota,omitempty" name:"SpeedTestingQuota"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSpeedTestingQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpeedTestingQuotaResponseParams `json:"Response"`
}

func (r *DescribeSpeedTestingQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpeedTestingQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL4DataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询指标，取值有：
	// <li>l4Flow_connections: 访问连接数；</li>
	// <li>l4Flow_flux: 访问总流量；</li>
	// <li>l4Flow_inFlux: 访问入流量；</li>
	// <li>l4Flow_outFlux: 访问出流量；</li>
	// <li> l4Flow_outPkt: 访问出包量。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 四层实例列表, 不填表示选择全部实例。
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟 ；</li>
	// <li>5min: 5分钟 ；</li>
	// <li>hour: 1小时 ；</li>
	// <li>day: 1天 。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件, key可选的值有：
	// <li>ruleId: 根据规则Id进行过滤；</li>
	// <li>proxyId: 根据通道Id进行过滤。</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeTimingL4DataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询指标，取值有：
	// <li>l4Flow_connections: 访问连接数；</li>
	// <li>l4Flow_flux: 访问总流量；</li>
	// <li>l4Flow_inFlux: 访问入流量；</li>
	// <li>l4Flow_outFlux: 访问出流量；</li>
	// <li> l4Flow_outPkt: 访问出包量。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 四层实例列表, 不填表示选择全部实例。
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟 ；</li>
	// <li>5min: 5分钟 ；</li>
	// <li>hour: 1小时 ；</li>
	// <li>day: 1天 。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件, key可选的值有：
	// <li>ruleId: 根据规则Id进行过滤；</li>
	// <li>proxyId: 根据通道Id进行过滤。</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeTimingL4DataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL4DataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "ProxyIds")
	delete(f, "Interval")
	delete(f, "Filters")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimingL4DataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL4DataResponseParams struct {
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 四层时序流量数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TimingDataRecord `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTimingL4DataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTimingL4DataResponseParams `json:"Response"`
}

func (r *DescribeTimingL4DataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL4DataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL7AnalysisDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标列表，取值有:
	// <li>l7Flow_outFlux: 访问流量；</li>
	// <li>l7Flow_request: 访问请求数；</li>
	// <li>l7Flow_outBandwidth: 访问带宽。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟；</li>
	// <li>5min: 5分钟；</li>
	// <li>hour: 1小时；</li>
	// <li>day: 1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，key可选的值有：
	// <li>country：国家/地区；</li>
	// <li>domain：域名；</li>
	// <li>protocol：协议类型；</li>
	// <li>resourceType：资源类型；</li>
	// <li>statusCode：状态码；</li>
	// <li> browserType：浏览器类型；</li>
	// <li>deviceType：设备类型；</li>
	// <li>operatingSystemType：操作系统类型；</li>
	// <li>tlsVersion：tls版本；</li>
	// <li>url：url地址；</li>
	// <li>referer：refer头信息；</li>
	// <li>ipVersion：ip版本；</li>
	// <li>tagKey：标签Key；</li>
	// <li>tagValue：标签Value。</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户的地域智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeTimingL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标列表，取值有:
	// <li>l7Flow_outFlux: 访问流量；</li>
	// <li>l7Flow_request: 访问请求数；</li>
	// <li>l7Flow_outBandwidth: 访问带宽。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟；</li>
	// <li>5min: 5分钟；</li>
	// <li>hour: 1小时；</li>
	// <li>day: 1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，key可选的值有：
	// <li>country：国家/地区；</li>
	// <li>domain：域名；</li>
	// <li>protocol：协议类型；</li>
	// <li>resourceType：资源类型；</li>
	// <li>statusCode：状态码；</li>
	// <li> browserType：浏览器类型；</li>
	// <li>deviceType：设备类型；</li>
	// <li>operatingSystemType：操作系统类型；</li>
	// <li>tlsVersion：tls版本；</li>
	// <li>url：url地址；</li>
	// <li>referer：refer头信息；</li>
	// <li>ipVersion：ip版本；</li>
	// <li>tagKey：标签Key；</li>
	// <li>tagValue：标签Value。</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户的地域智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeTimingL7AnalysisDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL7AnalysisDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "Interval")
	delete(f, "Filters")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimingL7AnalysisDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL7AnalysisDataResponseParams struct {
	// 时序流量数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TimingDataRecord `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTimingL7AnalysisDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTimingL7AnalysisDataResponseParams `json:"Response"`
}

func (r *DescribeTimingL7AnalysisDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL7AnalysisDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL7CacheDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询的指标，取值有：
	// <li>l7Cache_outFlux：响应流量；</li>
	// <li>l7Cache_request：响应请求数；</li>
	// <li> l7Cache_outBandwidth：响应带宽。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 筛选条件，key可选的值有：
	// <li> cacheType：缓存类型(状态)；</li>
	// <li>domain：Host/域名；</li>
	// <li>resourceType：资源类型；</li>
	// <li>url：url地址；</li>
	// <li>tagKey：标签Key；</li>
	// <li>tagValue：标签Value。</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// 查询时间粒度，可选的值有：
	// <li>min：1分钟的时间粒度；</li>
	// <li>5min：5分钟的时间粒度；</li>
	// <li>hour：1小时的时间粒度；</li>
	// <li>day：1天的时间粒度。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeTimingL7CacheDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询的指标，取值有：
	// <li>l7Cache_outFlux：响应流量；</li>
	// <li>l7Cache_request：响应请求数；</li>
	// <li> l7Cache_outBandwidth：响应带宽。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 筛选条件，key可选的值有：
	// <li> cacheType：缓存类型(状态)；</li>
	// <li>domain：Host/域名；</li>
	// <li>resourceType：资源类型；</li>
	// <li>url：url地址；</li>
	// <li>tagKey：标签Key；</li>
	// <li>tagValue：标签Value。</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// 查询时间粒度，可选的值有：
	// <li>min：1分钟的时间粒度；</li>
	// <li>5min：5分钟的时间粒度；</li>
	// <li>hour：1小时的时间粒度；</li>
	// <li>day：1天的时间粒度。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeTimingL7CacheDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL7CacheDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "Filters")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimingL7CacheDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL7CacheDataResponseParams struct {
	// 七层缓存分析时序类流量数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TimingDataRecord `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTimingL7CacheDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTimingL7CacheDataResponseParams `json:"Response"`
}

func (r *DescribeTimingL7CacheDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL7CacheDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopL7AnalysisDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询的指标，取值有：
	// <li> l7Flow_outFlux_country：请求的国家；</li>
	// <li> l7Flow_outFlux_statusCode：请求的状态码；</li>
	// <li> l7Flow_outFlux_domain：请求域名；</li>
	// <li> l7Flow_outFlux_url：请求的URL; </li>
	// <li> l7Flow_outFlux_resourceType：请求的资源类型；</li>
	// <li> l7Flow_outFlux_sip：客户端的源IP；</li>
	// <li> l7Flow_outFlux_referers：refer信息；</li>
	// <li> l7Flow_outFlux_ua_device：设备类型; </li>
	// <li> l7Flow_outFlux_ua_browser：浏览器类型；</li>
	// <li> l7Flow_outFlux_us_os：操作系统类型。</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 查询前多少个数据，最大值为1000，不填默认默认为: 10， 表示查询前top10的数据。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选条件，key可选的值有：
	// <li>country：国家/地区；</li>
	// <li>domain：域名；</li>
	// <li>protocol：协议类型；</li>
	// <li>resourceType：资源类型；</li>
	// <li>statusCode：状态码；</li>
	// <li> browserType：浏览器类型；</li>
	// <li>deviceType：设备类型；</li>
	// <li>operatingSystemType：操作系统类型；</li>
	// <li>tlsVersion：tls版本；</li>
	// <li>url：url地址；</li>
	// <li>referer：refer头信息；</li>
	// <li>ipVersion：ip版本；</li>
	// <li>tagKey：标签Key；</li>
	// <li>tagValue：标签Value。</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeTopL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询的指标，取值有：
	// <li> l7Flow_outFlux_country：请求的国家；</li>
	// <li> l7Flow_outFlux_statusCode：请求的状态码；</li>
	// <li> l7Flow_outFlux_domain：请求域名；</li>
	// <li> l7Flow_outFlux_url：请求的URL; </li>
	// <li> l7Flow_outFlux_resourceType：请求的资源类型；</li>
	// <li> l7Flow_outFlux_sip：客户端的源IP；</li>
	// <li> l7Flow_outFlux_referers：refer信息；</li>
	// <li> l7Flow_outFlux_ua_device：设备类型; </li>
	// <li> l7Flow_outFlux_ua_browser：浏览器类型；</li>
	// <li> l7Flow_outFlux_us_os：操作系统类型。</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 查询前多少个数据，最大值为1000，不填默认默认为: 10， 表示查询前top10的数据。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选条件，key可选的值有：
	// <li>country：国家/地区；</li>
	// <li>domain：域名；</li>
	// <li>protocol：协议类型；</li>
	// <li>resourceType：资源类型；</li>
	// <li>statusCode：状态码；</li>
	// <li> browserType：浏览器类型；</li>
	// <li>deviceType：设备类型；</li>
	// <li>operatingSystemType：操作系统类型；</li>
	// <li>tlsVersion：tls版本；</li>
	// <li>url：url地址；</li>
	// <li>referer：refer头信息；</li>
	// <li>ipVersion：ip版本；</li>
	// <li>tagKey：标签Key；</li>
	// <li>tagValue：标签Value。</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeTopL7AnalysisDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopL7AnalysisDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "ZoneIds")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopL7AnalysisDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopL7AnalysisDataResponseParams struct {
	// 七层流量前topN数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TopDataRecord `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTopL7AnalysisDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopL7AnalysisDataResponseParams `json:"Response"`
}

func (r *DescribeTopL7AnalysisDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopL7AnalysisDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopL7CacheDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询的指标，取值有：
	// <li> l7Cache_outFlux_domain：host/域名；</li>
	// <li> l7Cache_outFlux_url：url地址；</li>
	// <li> l7Cache_outFlux_resourceType：资源类型；</li>
	// <li> l7Cache_outFlux_statusCode：状态码。</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 站点id集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 查询前多少个数据，不填默认默认为10， 表示查询前top 10的数据。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选条件，key可选的值有：
	// <li> cacheType：缓存类型(状态)；</li>
	// <li>domain：Host/域名；</li>
	// <li>resourceType：资源类型；</li>
	// <li>url：url地址；</li>
	// <li>tagKey：标签Key；</li>
	// <li>tagValue：标签Value。</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟；</li>
	// <li>5min: 5分钟；</li>
	// <li>hour: 1小时；</li>
	// <li>day: 1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeTopL7CacheDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询的指标，取值有：
	// <li> l7Cache_outFlux_domain：host/域名；</li>
	// <li> l7Cache_outFlux_url：url地址；</li>
	// <li> l7Cache_outFlux_resourceType：资源类型；</li>
	// <li> l7Cache_outFlux_statusCode：状态码。</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 站点id集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 查询前多少个数据，不填默认默认为10， 表示查询前top 10的数据。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选条件，key可选的值有：
	// <li> cacheType：缓存类型(状态)；</li>
	// <li>domain：Host/域名；</li>
	// <li>resourceType：资源类型；</li>
	// <li>url：url地址；</li>
	// <li>tagKey：标签Key；</li>
	// <li>tagValue：标签Value。</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟；</li>
	// <li>5min: 5分钟；</li>
	// <li>hour: 1小时；</li>
	// <li>day: 1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeTopL7CacheDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopL7CacheDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "ZoneIds")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopL7CacheDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopL7CacheDataResponseParams struct {
	// 七层缓存TopN流量数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TopDataRecord `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTopL7CacheDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopL7CacheDataResponseParams `json:"Response"`
}

func (r *DescribeTopL7CacheDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopL7CacheDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>waf_interceptNum：waf拦截次数。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 子域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 筛选条件，key可选的值有：
	// <li>action：执行动作。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeWebManagedRulesDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>waf_interceptNum：waf拦截次数。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 子域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 筛选条件，key可选的值有：
	// <li>action：执行动作。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeWebManagedRulesDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "QueryCondition")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebManagedRulesDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesDataResponseParams struct {
	// WAF攻击的时序数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*SecEntry `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebManagedRulesDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebManagedRulesDataResponseParams `json:"Response"`
}

func (r *DescribeWebManagedRulesDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesHitRuleDetailRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 子域名列表，不填默认选择全部全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天 。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，key可选的值有：
	// <li>action ：执行动作 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeWebManagedRulesHitRuleDetailRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 子域名列表，不填默认选择全部全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天 。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，key可选的值有：
	// <li>action ：执行动作 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeWebManagedRulesHitRuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesHitRuleDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Interval")
	delete(f, "QueryCondition")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebManagedRulesHitRuleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesHitRuleDetailResponseParams struct {
	// 命中规则的详细列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*SecHitRuleInfo `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebManagedRulesHitRuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebManagedRulesHitRuleDetailResponseParams `json:"Response"`
}

func (r *DescribeWebManagedRulesHitRuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesHitRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesLogRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 筛选条件，key可选的值有：
	// <li>attackType：攻击类型；</li>
	// <li>riskLevel：风险等级；</li>
	// <li>action：执行动作（处置方式）；</li>
	// <li>ruleId：规则id；</li>
	// <li>sipCountryCode：ip所在国家；</li>
	// <li>attackIp：攻击ip；</li>
	// <li>oriDomain：被攻击的子域名；</li>
	// <li>eventId：事件id；</li>
	// <li>ua：用户代理；</li>
	// <li>requestMethod：请求方法；</li>
	// <li>uri：统一资源标识符。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeWebManagedRulesLogRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 筛选条件，key可选的值有：
	// <li>attackType：攻击类型；</li>
	// <li>riskLevel：风险等级；</li>
	// <li>action：执行动作（处置方式）；</li>
	// <li>ruleId：规则id；</li>
	// <li>sipCountryCode：ip所在国家；</li>
	// <li>attackIp：攻击ip；</li>
	// <li>oriDomain：被攻击的子域名；</li>
	// <li>eventId：事件id；</li>
	// <li>ua：用户代理；</li>
	// <li>requestMethod：请求方法；</li>
	// <li>uri：统一资源标识符。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeWebManagedRulesLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "QueryCondition")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebManagedRulesLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesLogResponseParams struct {
	// Web攻击日志数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*WebLogs `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebManagedRulesLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebManagedRulesLogResponseParams `json:"Response"`
}

func (r *DescribeWebManagedRulesLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionAttackEventsRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeWebProtectionAttackEventsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWebProtectionAttackEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionAttackEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebProtectionAttackEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionAttackEventsResponseParams struct {
	// CC相关攻击事件列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*CCInterceptEvent `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebProtectionAttackEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebProtectionAttackEventsResponseParams `json:"Response"`
}

func (r *DescribeWebProtectionAttackEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionAttackEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionClientIpListRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询的时间粒度，支持的粒度有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，key可选的值有：
	// <li>action：执行动作。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeWebProtectionClientIpListRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询的时间粒度，支持的粒度有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，key可选的值有：
	// <li>action：执行动作。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeWebProtectionClientIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionClientIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Interval")
	delete(f, "QueryCondition")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebProtectionClientIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionClientIpListResponseParams struct {
	// CC防护客户端（攻击源）ip信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*SecClientIp `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebProtectionClientIpListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebProtectionClientIpListResponseParams `json:"Response"`
}

func (r *DescribeWebProtectionClientIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionClientIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标，取值有：
	// <li>ccRate_interceptNum：速率限制规则限制次数；</li>
	// <li>ccAcl_interceptNum：自定义规则拦截次数。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询时间粒度，支持的时间粒度有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，key可选的值有：
	// <li>action：执行动作。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeWebProtectionDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标，取值有：
	// <li>ccRate_interceptNum：速率限制规则限制次数；</li>
	// <li>ccAcl_interceptNum：自定义规则拦截次数。</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询时间粒度，支持的时间粒度有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 筛选条件，key可选的值有：
	// <li>action：执行动作。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeWebProtectionDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Interval")
	delete(f, "QueryCondition")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebProtectionDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionDataResponseParams struct {
	// CC防护时序数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*SecEntry `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebProtectionDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebProtectionDataResponseParams `json:"Response"`
}

func (r *DescribeWebProtectionDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionHitRuleDetailRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 所属规则数据类型，支持的规则有：
	// <li>rate：限速规则；</li>
	// <li>acl：自定义规则。</li>
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名列表，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 筛选条件，key可选的值有：
	// <li>action：执行动作。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询时间粒度，支持的时间粒度有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeWebProtectionHitRuleDetailRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 所属规则数据类型，支持的规则有：
	// <li>rate：限速规则；</li>
	// <li>acl：自定义规则。</li>
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名列表，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 筛选条件，key可选的值有：
	// <li>action：执行动作。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询时间粒度，支持的时间粒度有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeWebProtectionHitRuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionHitRuleDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "EntityType")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "QueryCondition")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebProtectionHitRuleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionHitRuleDetailResponseParams struct {
	// cc防护命中规则列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*SecHitRuleInfo `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebProtectionHitRuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebProtectionHitRuleDetailResponseParams `json:"Response"`
}

func (r *DescribeWebProtectionHitRuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionHitRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionTopDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>ccRate_requestNum_url：速率限制规则请求次数url分布排行；</li>
	// <li>ccRate_cipRequestNum_region：速率限制规则请求次数区域客户端ip分布排行；</li>
	// <li>ccAcl_requestNum_url：自定义规则请求次数url分布排行；</li>
	// <li>ccAcl_requestNum_cip：自定义规则请求次数客户端ip分布排行；</li>
	// <li>ccAcl_cipRequestNum_region：自定义规则请求次数客户端区域分布排行。</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询前多少个数据，不填默认默认为10， 表示查询前top 10的数据。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选条件，key可选的值有：
	// <li>action：执行动作 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeWebProtectionTopDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>ccRate_requestNum_url：速率限制规则请求次数url分布排行；</li>
	// <li>ccRate_cipRequestNum_region：速率限制规则请求次数区域客户端ip分布排行；</li>
	// <li>ccAcl_requestNum_url：自定义规则请求次数url分布排行；</li>
	// <li>ccAcl_requestNum_cip：自定义规则请求次数客户端ip分布排行；</li>
	// <li>ccAcl_cipRequestNum_region：自定义规则请求次数客户端区域分布排行。</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询前多少个数据，不填默认默认为10， 表示查询前top 10的数据。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选条件，key可选的值有：
	// <li>action：执行动作 。</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeWebProtectionTopDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionTopDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "Interval")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Limit")
	delete(f, "QueryCondition")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebProtectionTopDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionTopDataResponseParams struct {
	// CC防护的TopN数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TopEntry `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebProtectionTopDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebProtectionTopDataResponseParams `json:"Response"`
}

func (r *DescribeWebProtectionTopDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionTopDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneDDoSPolicyRequestParams struct {
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeZoneDDoSPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeZoneDDoSPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneDDoSPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneDDoSPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneDDoSPolicyResponseParams struct {
	// DDoS防护分区。
	ShieldAreas []*ShieldArea `json:"ShieldAreas,omitempty" name:"ShieldAreas"`

	// 所有开启代理的子域名信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DDoSHosts []*DDoSHost `json:"DDoSHosts,omitempty" name:"DDoSHosts"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZoneDDoSPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneDDoSPolicyResponseParams `json:"Response"`
}

func (r *DescribeZoneDDoSPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneDDoSPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneSettingRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeZoneSettingRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
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

// Predefined struct for user
type DescribeZoneSettingResponseParams struct {
	// 站点配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneSetting *ZoneSetting `json:"ZoneSetting,omitempty" name:"ZoneSetting"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZoneSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneSettingResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeZonesRequestParams struct {
	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-name<br>   按照【<strong>站点名称</strong>】进行过滤。<br>   类型：String<br>   必选：否<li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-xxx。<br>   类型：String<br>   必选：否<li>status<br>   按照【<strong>站点状态</strong>】进行过滤。<br>   类型：String<br>   必选：否<li>tag-key<br>   按照【<strong>标签键</strong>】进行过滤。<br>   类型：String<br>   必选：否<li>tag-value<br>   按照【<strong>标签值</strong>】进行过滤。<br>   类型：String<br>   必选：否<li>Fuzzy<br>   按照【<strong>是否模糊查询</strong>】进行过滤。仅支持过滤字段名为zone-name。模糊查询时，Values长度最小为1。<br>   类型：Boolean<br>   必选：否<br>   默认值：false
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-name<br>   按照【<strong>站点名称</strong>】进行过滤。<br>   类型：String<br>   必选：否<li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-xxx。<br>   类型：String<br>   必选：否<li>status<br>   按照【<strong>站点状态</strong>】进行过滤。<br>   类型：String<br>   必选：否<li>tag-key<br>   按照【<strong>标签键</strong>】进行过滤。<br>   类型：String<br>   必选：否<li>tag-value<br>   按照【<strong>标签值</strong>】进行过滤。<br>   类型：String<br>   必选：否<li>Fuzzy<br>   按照【<strong>是否模糊查询</strong>】进行过滤。仅支持过滤字段名为zone-name。模糊查询时，Values长度最小为1。<br>   类型：Boolean<br>   必选：否<br>   默认值：false
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
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

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// 符合条件的站点个数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 站点详细信息列表。
	Zones []*Zone `json:"Zones,omitempty" name:"Zones"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZonesResponseParams `json:"Response"`
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
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 加速服务状态，取值为：
	// <li> process：部署中；</li>
	// <li> online：已启动；</li>
	// <li> offline：已关闭。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 域名。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 分配的Cname域名
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// 资源ID。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 锁状态。
	Lock *int64 `json:"Lock,omitempty" name:"Lock"`

	// 域名状态类型。
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 域名加速地域，取值有：
	// <li> global：全球；</li>
	// <li> mainland：中国大陆；</li>
	// <li> overseas：境外区域。</li>
	Area *string `json:"Area,omitempty" name:"Area"`

	// 加速类型配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccelerateType *AccelerateType `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// Https配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Https *Https `json:"Https,omitempty" name:"Https"`

	// 缓存配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheConfig *CacheConfig `json:"CacheConfig,omitempty" name:"CacheConfig"`

	// 源站配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// 安全类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityType *SecurityType `json:"SecurityType,omitempty" name:"SecurityType"`

	// 缓存键配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// 智能压缩配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// Waf防护配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Waf *Waf `json:"Waf,omitempty" name:"Waf"`

	// CC防护配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CC *CC `json:"CC,omitempty" name:"CC"`

	// DDoS防护配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DDoS *DDoS `json:"DDoS,omitempty" name:"DDoS"`

	// 智能路由配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmartRouting *SmartRouting `json:"SmartRouting,omitempty" name:"SmartRouting"`

	// Ipv6访问配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`
}

type DistrictStatistics struct {
	// ISO 3166-2 国家/地区简写，详情请参考[ISO 3166-2](https://zh.m.wikipedia.org/zh-hans/ISO_3166-2)。
	Alpha2 *string `json:"Alpha2,omitempty" name:"Alpha2"`

	// 整体拨测用时，单位ms。
	LoadTime *int64 `json:"LoadTime,omitempty" name:"LoadTime"`
}

type DnsData struct {
	// 时间。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 数值。
	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type DnsRecord struct {
	// 记录 ID。
	DnsRecordId *string `json:"DnsRecordId,omitempty" name:"DnsRecordId"`

	// DNS记录类型，取值有：
	// <li>A：将域名指向一个外网 IPv4 地址，如 8.8.8.8；</li>
	// <li>AAAA：将域名指向一个外网 IPv6 地址；</li>
	// <li>MX：用于邮箱服务器，相关记录值/优先级参数由邮件注册商提供。存在多条 MX 记录时，优先级越低越优先；</li>
	// <li>CNAME：将域名指向另一个域名，再由该域名解析出最终 IP 地址；</li>
	// <li>TXT：对域名进行标识和说明，常用于域名验证和 SPF 记录（反垃圾邮件）；</li>
	// <li>NS：如果需要将子域名交给其他 DNS 服务商解析，则需要添加 NS 记录。根域名无法添加 NS 记录；</li>
	// <li>CAA：指定可为本站点颁发证书的 CA；</li>
	// <li>SRV：标识某台服务器使用了某个服务，常见于微软系统的目录管理。</li>
	DnsRecordType *string `json:"DnsRecordType,omitempty" name:"DnsRecordType"`

	// 记录名称。
	DnsRecordName *string `json:"DnsRecordName,omitempty" name:"DnsRecordName"`

	// 记录值。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 代理模式，取值有：
	// <li>dns_only：仅DNS解析；</li>
	// <li>proxied：代理加速。</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 缓存时间，数值越小，修改记录各地生效时间越快，单位：秒。
	TTL *int64 `json:"TTL,omitempty" name:"TTL"`

	// MX记录优先级，数值越小越优先。
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 创建时间。
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 修改时间。
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// 域名锁定状态。
	Locked *bool `json:"Locked,omitempty" name:"Locked"`

	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 记录解析状态，取值有：
	// <li>active：生效；</li>
	// <li>pending：不生效。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// CNAME 地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// 域名服务类型，取值有：
	// <li>lb：负载均衡；</li>
	// <li>security：安全；</li>
	// <li>l4：四层代理。</li>
	DomainStatus []*string `json:"DomainStatus,omitempty" name:"DomainStatus"`
}

type DnssecInfo struct {
	// 标志。
	Flags *int64 `json:"Flags,omitempty" name:"Flags"`

	// 加密算法。
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// 加密类型。
	KeyType *string `json:"KeyType,omitempty" name:"KeyType"`

	// 摘要类型。
	DigestType *string `json:"DigestType,omitempty" name:"DigestType"`

	// 摘要算法。
	DigestAlgorithm *string `json:"DigestAlgorithm,omitempty" name:"DigestAlgorithm"`

	// 摘要信息。
	Digest *string `json:"Digest,omitempty" name:"Digest"`

	// DS 记录值。
	DS *string `json:"DS,omitempty" name:"DS"`

	// 密钥标签。
	KeyTag *int64 `json:"KeyTag,omitempty" name:"KeyTag"`

	// 公钥。
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
}

// Predefined struct for user
type DownloadL4LogsRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 四层实例ID集合。
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DownloadL4LogsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 四层实例ID集合。
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DownloadL4LogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadL4LogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ZoneIds")
	delete(f, "ProxyIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadL4LogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadL4LogsResponseParams struct {
	// 四层离线日志数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*L4OfflineLog `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadL4LogsResponse struct {
	*tchttp.BaseResponse
	Response *DownloadL4LogsResponseParams `json:"Response"`
}

func (r *DownloadL4LogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadL4LogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadL7LogsRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 子域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DownloadL7LogsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 子域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadL7LogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadL7LogsResponseParams struct {
	// 七层离线日志数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*L7OfflineLog `json:"Data,omitempty" name:"Data"`

	// 查询结果的总条数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadL7LogsResponse struct {
	*tchttp.BaseResponse
	Response *DownloadL7LogsResponseParams `json:"Response"`
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

type DropPageConfig struct {
	// 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Waf(托管规则)模块的拦截页面配置。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WafDropPageDetail *DropPageDetail `json:"WafDropPageDetail,omitempty" name:"WafDropPageDetail"`

	// 自定义页面的拦截页面配置。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclDropPageDetail *DropPageDetail `json:"AclDropPageDetail,omitempty" name:"AclDropPageDetail"`
}

type DropPageDetail struct {
	// 拦截页面的唯一Id。系统默认包含一个自带拦截页面，Id值为0。
	// 该Id可通过创建拦截页面接口进行上传获取。如传入0，代表使用系统默认拦截页面。
	PageId *int64 `json:"PageId,omitempty" name:"PageId"`

	// 拦截页面的HTTP状态码。状态码范围是100-600。
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 页面文件名或url。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 页面的类型，取值有：
	// <li> file：页面文件内容；</li>
	// <li> url：上传的url地址。</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

type ExceptConfig struct {
	// 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 例外规则详情。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptUserRules []*ExceptUserRule `json:"ExceptUserRules,omitempty" name:"ExceptUserRules"`
}

type ExceptUserRule struct {
	// 规则名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则的处置方式，当前仅支持skip：跳过全部托管规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 规则生效状态，取值有：
	// <li>on：生效；</li>
	// <li>off：失效。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleStatus *string `json:"RuleStatus,omitempty" name:"RuleStatus"`

	// 规则ID。仅出参使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// 更新时间。仅出参使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 匹配条件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptUserRuleConditions []*ExceptUserRuleCondition `json:"ExceptUserRuleConditions,omitempty" name:"ExceptUserRuleConditions"`

	// 规则生效的范围。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptUserRuleScope *ExceptUserRuleScope `json:"ExceptUserRuleScope,omitempty" name:"ExceptUserRuleScope"`

	// 优先级，取值范围0-100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RulePriority *int64 `json:"RulePriority,omitempty" name:"RulePriority"`
}

type ExceptUserRuleCondition struct {
	// 匹配项，取值有：
	// <li>host：请求域名；</li>
	// <li>sip：客户端IP；</li>
	// <li>ua：User-Agent；</li>
	// <li>cookie：会话 Cookie；</li>
	// <li>cgi：CGI 脚本；</li>
	// <li>xff：XFF 扩展头部；</li>
	// <li>url：请求 URL；</li>
	// <li>accept：请求内容类型；</li>
	// <li>method：请求方式；</li>
	// <li>header：请求头部；</li>
	// <li>sip_proto：网络层协议。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchFrom *string `json:"MatchFrom,omitempty" name:"MatchFrom"`

	// 匹配项的参数。当 MatchFrom 为 header 时，可以填入 header 的 key 作为参数。
	MatchParam *string `json:"MatchParam,omitempty" name:"MatchParam"`

	// 匹配操作符，取值有：
	// <li>equal：字符串等于；</li>
	// <li>not_equal：数值不等于；</li>
	// <li>include：字符包含；</li>
	// <li>not_include：字符不包含；</li>
	// <li>match：ip匹配；</li>
	// <li>not_match：ip不匹配；</li>
	// <li>include_area：地域包含；</li>
	// <li>is_empty：存在字段但值为空；</li>
	// <li>not_exists：不存在关键字段；</li>
	// <li>regexp：正则匹配；</li>
	// <li>len_gt：数值大于；</li>
	// <li>len_lt：数值小于；</li>
	// <li>len_eq：数值等于；</li>
	// <li>match_prefix：前缀匹配；</li>
	// <li>match_suffix：后缀匹配；</li>
	// <li>wildcard：通配符。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 匹配值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchContent *string `json:"MatchContent,omitempty" name:"MatchContent"`
}

type ExceptUserRuleScope struct {
	// 生效的模块。当前仅支持waf：托管规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Modules []*string `json:"Modules,omitempty" name:"Modules"`
}

type FailReason struct {
	// 失败原因。
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 处理失败的资源列表，该列表元素来源于输入参数中的Targets，因此格式和入参中的Targets保持一致。
	Targets []*string `json:"Targets,omitempty" name:"Targets"`
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type FollowOrigin struct {
	// 遵循源站配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ForceRedirect struct {
	// 访问强制跳转配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 重定向状态码，取值有：
	// <li>301：301跳转；</li>
	// <li>302：302跳转。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitempty" name:"RedirectStatusCode"`
}

type GeoIp struct {
	// 地域ID。
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 国家名。
	Country *string `json:"Country,omitempty" name:"Country"`

	// 所属洲。
	Continent *string `json:"Continent,omitempty" name:"Continent"`

	// 城市名。
	Province *string `json:"Province,omitempty" name:"Province"`
}

type Header struct {
	// HTTP头部。
	Name *string `json:"Name,omitempty" name:"Name"`

	// HTTP头部值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type HostCertInfo struct {
	// 服务器证书 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 证书备注名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 证书类型，取值有：
	// <li>default：默认证书；</lil>
	// <li>upload：用户上传；</li>
	// <li>managed:腾讯云托管。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 证书过期时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 证书部署时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployTime *string `json:"DeployTime,omitempty" name:"DeployTime"`

	// 签名算法。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignAlgo *string `json:"SignAlgo,omitempty" name:"SignAlgo"`

	// 证书状态，取值有：
	// <li>deployed：已部署;</li>
	// <li>process：部署中。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type HostsCertificate struct {
	// 域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 服务端证书配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostCertInfo *HostCertInfo `json:"HostCertInfo,omitempty" name:"HostCertInfo"`
}

type Hsts struct {
	// 是否开启，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// MaxAge数值。单位为秒，最大值为1天。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAge *int64 `json:"MaxAge,omitempty" name:"MaxAge"`

	// 是否包含子域名，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncludeSubDomains *string `json:"IncludeSubDomains,omitempty" name:"IncludeSubDomains"`

	// 是否开启预加载，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Preload *string `json:"Preload,omitempty" name:"Preload"`
}

type Https struct {
	// http2 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Http2 *string `json:"Http2,omitempty" name:"Http2"`

	// OCSP 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcspStapling *string `json:"OcspStapling,omitempty" name:"OcspStapling"`

	// Tls版本设置，取值有：
	// <li>TLSv1：TLSv1版本；</li>
	// <li>TLSV1.1：TLSv1.1版本；</li>
	// <li>TLSV1.2：TLSv1.2版本；</li>
	// <li>TLSv1.3：TLSv1.3版本。</li>修改时必须开启连续的版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TlsVersion []*string `json:"TlsVersion,omitempty" name:"TlsVersion"`

	// HSTS 配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hsts *Hsts `json:"Hsts,omitempty" name:"Hsts"`

	// 证书配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertInfo []*ServerCertInfo `json:"CertInfo,omitempty" name:"CertInfo"`
}

type Identification struct {
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 验证状态，取值有：
	// <li> pending：验证中；</li>
	// <li> finished：验证完成。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 站点归属信息。
	Ascription *AscriptionInfo `json:"Ascription,omitempty" name:"Ascription"`

	// 域名当前的 NS 记录。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalNameServers []*string `json:"OriginalNameServers,omitempty" name:"OriginalNameServers"`
}

// Predefined struct for user
type IdentifyZoneRequestParams struct {
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type IdentifyZoneRequest struct {
	*tchttp.BaseRequest
	
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
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
	delete(f, "ZoneName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IdentifyZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IdentifyZoneResponseParams struct {
	// 站点归属信息。
	Ascription *AscriptionInfo `json:"Ascription,omitempty" name:"Ascription"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type IdentifyZoneResponse struct {
	*tchttp.BaseResponse
	Response *IdentifyZoneResponseParams `json:"Response"`
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

type IntelligenceRule struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 规则详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntelligenceRuleItems []*IntelligenceRuleItem `json:"IntelligenceRuleItems,omitempty" name:"IntelligenceRuleItems"`
}

type IntelligenceRuleItem struct {
	// 智能分析标签，取值有：
	// <li>evil_bot：恶意bot；</li>
	// <li>suspect_bot：疑似bot；</li>
	// <li>good_bot：良好bot；</li>
	// <li>normal：正常请求。</li>
	Label *string `json:"Label,omitempty" name:"Label"`

	// 触发智能分析标签对应的处置方式，取值有：
	// <li>drop：拦截；</li>
	// <li>trans：放行；</li>
	// <li>alg：Javascript挑战；</li>
	// <li>captcha：数字验证码；</li>
	// <li>monitor：观察。</li>
	Action *string `json:"Action,omitempty" name:"Action"`
}

type IpTableConfig struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭；</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 基础管控规则。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpTableRules []*IpTableRule `json:"IpTableRules,omitempty" name:"IpTableRules"`
}

type IpTableRule struct {
	// 动作，取值有：
	// <li> drop：拦截；</li>
	// <li> trans：放行；</li>
	// <li> monitor：观察。</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// 根据类型匹配，取值有：
	// <li>ip：对ip进行匹配；</li>
	// <li>area：对ip所属地区匹配。</li>
	MatchFrom *string `json:"MatchFrom,omitempty" name:"MatchFrom"`

	// 匹配内容。
	MatchContent *string `json:"MatchContent,omitempty" name:"MatchContent"`

	// 规则id。仅出参使用。
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// 更新时间。仅出参使用。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type Ipv6 struct {
	// Ipv6访问功能配置，取值有：
	// <li>on：开启Ipv6访问功能；</li>
	// <li>off：关闭Ipv6访问功能。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type L4OfflineLog struct {
	// 日志打包开始时间。
	LogTime *int64 `json:"LogTime,omitempty" name:"LogTime"`

	// 四层实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 原始大小 单位byte。
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 下载地址。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 日志数据包名。
	LogPacketName *string `json:"LogPacketName,omitempty" name:"LogPacketName"`

	// 加速区域，取值有：
	// <li>mainland：中国大陆境内;</li>
	// <li>overseas：全球（不含中国大陆）。</li>
	Area *string `json:"Area,omitempty" name:"Area"`
}

type L7OfflineLog struct {
	// 日志打包开始时间。
	LogTime *int64 `json:"LogTime,omitempty" name:"LogTime"`

	// 子域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 原始大小，单位byte。
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 下载地址。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 日志数据包名。
	LogPacketName *string `json:"LogPacketName,omitempty" name:"LogPacketName"`

	// 加速区域，取值有：
	// <li>mainland：中国大陆境内; </li>
	// <li>overseas：全球（不含中国大陆）。</li>
	Area *string `json:"Area,omitempty" name:"Area"`
}

type LoadBalancing struct {
	// 负载均衡ID。
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名，填写@表示根域。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 代理模式，取值有：
	// <li>dns_only：仅DNS；</li>
	// <li>proxied：开启代理。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 当Type=dns_only表示DNS记录的缓存时间。
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// 状态，取值有：
	// <li>online：部署成功；</li>
	// <li>process：部署中。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 调度域名。
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// 主源源站组ID。
	OriginGroupId *string `json:"OriginGroupId,omitempty" name:"OriginGroupId"`

	// 备用源站源站组ID。为空表示不适用备用源站。
	BackupOriginGroupId *string `json:"BackupOriginGroupId,omitempty" name:"BackupOriginGroupId"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type LogSetInfo struct {
	// 日志集所属地区。
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`

	// 日志集名
	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`

	// 日志集Id
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 该日志集是否已被删除, 可选的值有：
	// <li>no: 日志集没有被删除；</li>
	// <li>yes: 日志集已经被删除；</li>
	Deleted *string `json:"Deleted,omitempty" name:"Deleted"`
}

type LogTopicDetailInfo struct {
	// 推送任务的任务名称。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 日志集所属的地域。
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`

	// 推送任务的类型。
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// 任务实体列表。
	EntityList []*string `json:"EntityList,omitempty" name:"EntityList"`

	// 日志集ID。
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 日志集名称。
	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`

	// 推送任务主题ID。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 推送任务主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 推送任务主题保存时间，单位为天。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 推送任务是否开启。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 推送任务创建时间，时间格式为: YYYY-mm-dd HH:MM:SS。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 加速区域，取值有：
	// <li>mainland：中国大陆境内;</li>
	// <li>overseas：全球（不含中国大陆）。</li>
	Area *string `json:"Area,omitempty" name:"Area"`

	// 站点ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 是否被删除了，取值有：
	// <li>yes: 已经被删除；</li>
	// <li>no: 没有被删除。</li>
	Deleted *string `json:"Deleted,omitempty" name:"Deleted"`
}

type MaxAge struct {
	// 是否遵循源站，取值有：
	// <li>on：遵循源站，忽略MaxAge 时间设置；</li>
	// <li>off：不遵循源站，使用MaxAge 时间设置。</li>
	FollowOrigin *string `json:"FollowOrigin,omitempty" name:"FollowOrigin"`

	// MaxAge 时间设置，单位秒，最大365天。
	// 注意：时间为0，即不缓存。
	MaxAgeTime *int64 `json:"MaxAgeTime,omitempty" name:"MaxAgeTime"`
}

// Predefined struct for user
type ModifyAlarmConfigRequestParams struct {
	// 告警服务类型，取值有：
	// <li>ddos：ddos告警服务。</li>
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 告警维度值列表。
	EntityList []*string `json:"EntityList,omitempty" name:"EntityList"`

	// 告警阈值，不传或者传0表示不修改阈值。
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`

	// 是否使用默认值，只有在不传Threshold或者Threshold=0时该参数有效。
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

type ModifyAlarmConfigRequest struct {
	*tchttp.BaseRequest
	
	// 告警服务类型，取值有：
	// <li>ddos：ddos告警服务。</li>
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 告警维度值列表。
	EntityList []*string `json:"EntityList,omitempty" name:"EntityList"`

	// 告警阈值，不传或者传0表示不修改阈值。
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`

	// 是否使用默认值，只有在不传Threshold或者Threshold=0时该参数有效。
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

func (r *ModifyAlarmConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "ZoneId")
	delete(f, "EntityList")
	delete(f, "Threshold")
	delete(f, "IsDefault")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAlarmConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmConfigResponseParams `json:"Response"`
}

func (r *ModifyAlarmConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmDefaultThresholdRequestParams struct {
	// 告警服务类型，取值有：
	// <li>ddos：ddos告警服务。</li>
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 新的阈值，单位为Mbps，最小阈值为10。
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`

	// 防护实体，如果是四层防护，防护实体为通道ID。如果是七层防护，防护实体为站点名称。
	Entity *string `json:"Entity,omitempty" name:"Entity"`
}

type ModifyAlarmDefaultThresholdRequest struct {
	*tchttp.BaseRequest
	
	// 告警服务类型，取值有：
	// <li>ddos：ddos告警服务。</li>
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 新的阈值，单位为Mbps，最小阈值为10。
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`

	// 防护实体，如果是四层防护，防护实体为通道ID。如果是七层防护，防护实体为站点名称。
	Entity *string `json:"Entity,omitempty" name:"Entity"`
}

func (r *ModifyAlarmDefaultThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmDefaultThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "ZoneId")
	delete(f, "Threshold")
	delete(f, "Entity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmDefaultThresholdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmDefaultThresholdResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAlarmDefaultThresholdResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmDefaultThresholdResponseParams `json:"Response"`
}

func (r *ModifyAlarmDefaultThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmDefaultThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 当ProxyType=hostname时，表示域名或子域名；
	// 当ProxyType=instance时，表示代理名称。
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// 会话保持时间，取值范围：30-3600，单位：秒。
	// 不填写保持原有配置。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// 四层代理模式，取值有：
	// <li>hostname：表示子域名模式；</li>
	// <li>instance：表示实例模式。</li>不填写保持原有配置。
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// Ipv6访问配置，不填写保持原有配置。
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`
}

type ModifyApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 当ProxyType=hostname时，表示域名或子域名；
	// 当ProxyType=instance时，表示代理名称。
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// 会话保持时间，取值范围：30-3600，单位：秒。
	// 不填写保持原有配置。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// 四层代理模式，取值有：
	// <li>hostname：表示子域名模式；</li>
	// <li>instance：表示实例模式。</li>不填写保持原有配置。
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// Ipv6访问配置，不填写保持原有配置。
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`
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
	delete(f, "SessionPersistTime")
	delete(f, "ProxyType")
	delete(f, "Ipv6")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApplicationProxyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationProxyResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyApplicationProxyRuleRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 源站类型，取值有：
	// <li>custom：手动添加；</li>
	// <li>origins：源站组。</li>不填保持原有值。
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 端口，支持格式：
	// 80：80端口
	// 81-90：81至90端口。不填保持原有值。
	Port []*string `json:"Port,omitempty" name:"Port"`

	// 协议，取值有：
	// <li>TCP：TCP协议；</li>
	// <li>UDP：UDP协议。</li>不填保持原有值。
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// 源站信息：
	// 当OriginType=custom时，表示一个或多个源站，如：
	// OriginValue=["8.8.8.8:80","9.9.9.9:80"]
	// OriginValue=["test.com:80"]；
	// 当OriginType=origins时，要求有且仅有一个元素，表示源站组ID，如：
	// OriginValue=["origin-537f5b41-162a-11ed-abaa-525400c5da15"]。
	// 不填保持原有值。
	OriginValue []*string `json:"OriginValue,omitempty" name:"OriginValue"`

	// 传递客户端IP，取值有：
	// <li>TOA：TOA（仅Proto=TCP时可选）；</li>
	// <li>PPV1：Proxy Protocol传递，协议版本V1（仅Proto=TCP时可选）；</li>
	// <li>PPV2：Proxy Protocol传递，协议版本V2；</li>
	// <li>OFF：不传递。</li>不填保持原有值。
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// 是否开启会话保持，取值有：
	// <li>true：开启；</li>
	// <li>false：关闭。</li>不填保持原有值。
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`
}

type ModifyApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 源站类型，取值有：
	// <li>custom：手动添加；</li>
	// <li>origins：源站组。</li>不填保持原有值。
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 端口，支持格式：
	// 80：80端口
	// 81-90：81至90端口。不填保持原有值。
	Port []*string `json:"Port,omitempty" name:"Port"`

	// 协议，取值有：
	// <li>TCP：TCP协议；</li>
	// <li>UDP：UDP协议。</li>不填保持原有值。
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// 源站信息：
	// 当OriginType=custom时，表示一个或多个源站，如：
	// OriginValue=["8.8.8.8:80","9.9.9.9:80"]
	// OriginValue=["test.com:80"]；
	// 当OriginType=origins时，要求有且仅有一个元素，表示源站组ID，如：
	// OriginValue=["origin-537f5b41-162a-11ed-abaa-525400c5da15"]。
	// 不填保持原有值。
	OriginValue []*string `json:"OriginValue,omitempty" name:"OriginValue"`

	// 传递客户端IP，取值有：
	// <li>TOA：TOA（仅Proto=TCP时可选）；</li>
	// <li>PPV1：Proxy Protocol传递，协议版本V1（仅Proto=TCP时可选）；</li>
	// <li>PPV2：Proxy Protocol传递，协议版本V2；</li>
	// <li>OFF：不传递。</li>不填保持原有值。
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// 是否开启会话保持，取值有：
	// <li>true：开启；</li>
	// <li>false：关闭。</li>不填保持原有值。
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
	delete(f, "OriginType")
	delete(f, "Port")
	delete(f, "Proto")
	delete(f, "OriginValue")
	delete(f, "ForwardClientIp")
	delete(f, "SessionPersist")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApplicationProxyRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationProxyRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyApplicationProxyRuleStatusRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 状态，取值有：
	// <li>offline: 停用；</li>
	// <li>online: 启用。</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyApplicationProxyRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 状态，取值有：
	// <li>offline: 停用；</li>
	// <li>online: 启用。</li>
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

// Predefined struct for user
type ModifyApplicationProxyRuleStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApplicationProxyRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationProxyRuleStatusResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyApplicationProxyStatusRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 状态，取值有：
	// <li>offline: 停用；</li>
	// <li>online: 启用。</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyApplicationProxyStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 状态，取值有：
	// <li>offline: 停用；</li>
	// <li>online: 启用。</li>
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

// Predefined struct for user
type ModifyApplicationProxyStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApplicationProxyStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationProxyStatusResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDDoSPolicyHostRequestParams struct {
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 加速开关，取值有：
	// <li>on：开启加速；</li>
	// <li>off：关闭加速。</li>
	AccelerateType *string `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// 策略id。
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 安全开关，取值有：
	// <li>on：开启安全；</li>
	// <li>off：关闭安全。</li>
	SecurityType *string `json:"SecurityType,omitempty" name:"SecurityType"`
}

type ModifyDDoSPolicyHostRequest struct {
	*tchttp.BaseRequest
	
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 加速开关，取值有：
	// <li>on：开启加速；</li>
	// <li>off：关闭加速。</li>
	AccelerateType *string `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// 策略id。
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 安全开关，取值有：
	// <li>on：开启安全；</li>
	// <li>off：关闭安全。</li>
	SecurityType *string `json:"SecurityType,omitempty" name:"SecurityType"`
}

func (r *ModifyDDoSPolicyHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSPolicyHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Host")
	delete(f, "AccelerateType")
	delete(f, "PolicyId")
	delete(f, "SecurityType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSPolicyHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSPolicyHostResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDDoSPolicyHostResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSPolicyHostResponseParams `json:"Response"`
}

func (r *ModifyDDoSPolicyHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSPolicyHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSPolicyRequestParams struct {
	// 策略Id。
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// DDoS防护配置详情。
	DDoSRule *DDoSRule `json:"DDoSRule,omitempty" name:"DDoSRule"`
}

type ModifyDDoSPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略Id。
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// DDoS防护配置详情。
	DDoSRule *DDoSRule `json:"DDoSRule,omitempty" name:"DDoSRule"`
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
	delete(f, "PolicyId")
	delete(f, "ZoneId")
	delete(f, "DDoSRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSPolicyResponseParams struct {
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
type ModifyDefaultCertificateRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 默认证书ID。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 证书状态，取值有：
	// <li>deployed ：部署证书；</li>
	// <li>disabled ：禁用证书。</li>失败状态下重新deployed即可重试。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyDefaultCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 默认证书ID。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 证书状态，取值有：
	// <li>deployed ：部署证书；</li>
	// <li>disabled ：禁用证书。</li>失败状态下重新deployed即可重试。
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

// Predefined struct for user
type ModifyDefaultCertificateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDefaultCertificateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDefaultCertificateResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDnsRecordRequestParams struct {
	// 记录ID。
	DnsRecordId *string `json:"DnsRecordId,omitempty" name:"DnsRecordId"`

	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// DNS记录类型，取值有：
	// <li>A：将域名指向一个外网 IPv4 地址，如 8.8.8.8；</li>
	// <li>AAAA：将域名指向一个外网 IPv6 地址；</li>
	// <li>MX：用于邮箱服务器，相关记录值/优先级参数由邮件注册商提供。存在多条 MX 记录时，优先级越低越优先；</li>
	// <li>CNAME：将域名指向另一个域名，再由该域名解析出最终 IP 地址；</li>
	// <li>TXT：对域名进行标识和说明，常用于域名验证和 SPF 记录（反垃圾邮件）；</li>
	// <li>NS：如果需要将子域名交给其他 DNS 服务商解析，则需要添加 NS 记录。根域名无法添加 NS 记录；</li>
	// <li>CAA：指定可为本站点颁发证书的 CA；</li>
	// <li>SRV：标识某台服务器使用了某个服务，常见于微软系统的目录管理。</li>不填写保持原有配置。
	DnsRecordType *string `json:"DnsRecordType,omitempty" name:"DnsRecordType"`

	// 记录名称，由主机记录+站点名称组成，不填写保持原有配置。
	DnsRecordName *string `json:"DnsRecordName,omitempty" name:"DnsRecordName"`

	// 记录内容，不填写保持原有配置。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 缓存时间，数值越小，修改记录各地生效时间越快，默认为300，单位：秒，不填写保持原有配置。
	TTL *int64 `json:"TTL,omitempty" name:"TTL"`

	// 该参数在修改MX记录时生效，值越小优先级越高，用户可指定值范围为1~50，不指定默认为0，不填写保持原有配置。
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 代理模式，取值有：
	// <li>dns_only：仅DNS解析；</li>
	// <li>proxied：代理加速。</li>不填写保持原有配置。
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type ModifyDnsRecordRequest struct {
	*tchttp.BaseRequest
	
	// 记录ID。
	DnsRecordId *string `json:"DnsRecordId,omitempty" name:"DnsRecordId"`

	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// DNS记录类型，取值有：
	// <li>A：将域名指向一个外网 IPv4 地址，如 8.8.8.8；</li>
	// <li>AAAA：将域名指向一个外网 IPv6 地址；</li>
	// <li>MX：用于邮箱服务器，相关记录值/优先级参数由邮件注册商提供。存在多条 MX 记录时，优先级越低越优先；</li>
	// <li>CNAME：将域名指向另一个域名，再由该域名解析出最终 IP 地址；</li>
	// <li>TXT：对域名进行标识和说明，常用于域名验证和 SPF 记录（反垃圾邮件）；</li>
	// <li>NS：如果需要将子域名交给其他 DNS 服务商解析，则需要添加 NS 记录。根域名无法添加 NS 记录；</li>
	// <li>CAA：指定可为本站点颁发证书的 CA；</li>
	// <li>SRV：标识某台服务器使用了某个服务，常见于微软系统的目录管理。</li>不填写保持原有配置。
	DnsRecordType *string `json:"DnsRecordType,omitempty" name:"DnsRecordType"`

	// 记录名称，由主机记录+站点名称组成，不填写保持原有配置。
	DnsRecordName *string `json:"DnsRecordName,omitempty" name:"DnsRecordName"`

	// 记录内容，不填写保持原有配置。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 缓存时间，数值越小，修改记录各地生效时间越快，默认为300，单位：秒，不填写保持原有配置。
	TTL *int64 `json:"TTL,omitempty" name:"TTL"`

	// 该参数在修改MX记录时生效，值越小优先级越高，用户可指定值范围为1~50，不指定默认为0，不填写保持原有配置。
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 代理模式，取值有：
	// <li>dns_only：仅DNS解析；</li>
	// <li>proxied：代理加速。</li>不填写保持原有配置。
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
	delete(f, "DnsRecordId")
	delete(f, "ZoneId")
	delete(f, "DnsRecordType")
	delete(f, "DnsRecordName")
	delete(f, "Content")
	delete(f, "TTL")
	delete(f, "Priority")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDnsRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDnsRecordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDnsRecordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDnsRecordResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDnssecRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// DNSSEC状态，取值有
	// <li>enabled：开启；</li>
	// <li>disabled：关闭。</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyDnssecRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// DNSSEC状态，取值有
	// <li>enabled：开启；</li>
	// <li>disabled：关闭。</li>
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
	delete(f, "ZoneId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDnssecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDnssecResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDnssecResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDnssecResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyHostsCertificateRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 本次变更的域名列表。
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`

	// 证书信息, 只需要传入 CertId 即可, 如果为空, 则使用默认证书。
	ServerCertInfo []*ServerCertInfo `json:"ServerCertInfo,omitempty" name:"ServerCertInfo"`
}

type ModifyHostsCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 本次变更的域名列表。
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`

	// 证书信息, 只需要传入 CertId 即可, 如果为空, 则使用默认证书。
	ServerCertInfo []*ServerCertInfo `json:"ServerCertInfo,omitempty" name:"ServerCertInfo"`
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
	delete(f, "ServerCertInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHostsCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostsCertificateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyHostsCertificateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHostsCertificateResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyLoadBalancingRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 负载均衡ID。
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// 代理模式，取值有：
	// <li>dns_only：仅DNS；</li>
	// <li>proxied：开启代理。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 主源站源站组ID。
	OriginGroupId *string `json:"OriginGroupId,omitempty" name:"OriginGroupId"`

	// 备用源站源站组ID，当Type=proxied时可以填写，为空表示不使用备用源站。
	BackupOriginGroupId *string `json:"BackupOriginGroupId,omitempty" name:"BackupOriginGroupId"`

	// 当Type=dns_only时，指解析记录在DNS服务器缓存的生存时间。
	// 取值范围60-86400，单位：秒，不填写使用默认值：600。
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`
}

type ModifyLoadBalancingRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 负载均衡ID。
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// 代理模式，取值有：
	// <li>dns_only：仅DNS；</li>
	// <li>proxied：开启代理。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 主源站源站组ID。
	OriginGroupId *string `json:"OriginGroupId,omitempty" name:"OriginGroupId"`

	// 备用源站源站组ID，当Type=proxied时可以填写，为空表示不使用备用源站。
	BackupOriginGroupId *string `json:"BackupOriginGroupId,omitempty" name:"BackupOriginGroupId"`

	// 当Type=dns_only时，指解析记录在DNS服务器缓存的生存时间。
	// 取值范围60-86400，单位：秒，不填写使用默认值：600。
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
	delete(f, "OriginGroupId")
	delete(f, "BackupOriginGroupId")
	delete(f, "TTL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLoadBalancingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancingResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyLoadBalancingStatusRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 负载均衡ID。
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// 负载均衡状态，取值有：
	// <li>online：启用；</li>
	// <li>offline：停用。</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyLoadBalancingStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 负载均衡ID。
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// 负载均衡状态，取值有：
	// <li>online：启用；</li>
	// <li>offline：停用。</li>
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

// Predefined struct for user
type ModifyLoadBalancingStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLoadBalancingStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancingStatusResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyLogTopicTaskRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 日志集所属地区。
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`

	// 日志集ID。
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 日志主题ID。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 数据推送类型，可选的类型有：
	// <li>domain：七层代理日志；</li>
	// <li>application：四层代理日志；</li>
	// <li>web-rateLiming：速率限制日志；</li>
	// <li>web-attack：Web攻击防护日志；</li>
	// <li>web-rule：自定义规则日志；</li>
	// <li>web-bot：Bot管理日志。</li>
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// 推送任务名。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 待更新的主题名称，不填表示不更新主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 更新后的日志集名称。
	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`

	// 更新后的日志集保存时间。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 待添加的推送任务实体列表。
	DropEntityList []*string `json:"DropEntityList,omitempty" name:"DropEntityList"`

	// 待删除的推送任务实例列表。
	AddedEntityList []*string `json:"AddedEntityList,omitempty" name:"AddedEntityList"`
}

type ModifyLogTopicTaskRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 日志集所属地区。
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`

	// 日志集ID。
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 日志主题ID。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 数据推送类型，可选的类型有：
	// <li>domain：七层代理日志；</li>
	// <li>application：四层代理日志；</li>
	// <li>web-rateLiming：速率限制日志；</li>
	// <li>web-attack：Web攻击防护日志；</li>
	// <li>web-rule：自定义规则日志；</li>
	// <li>web-bot：Bot管理日志。</li>
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// 推送任务名。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 待更新的主题名称，不填表示不更新主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 更新后的日志集名称。
	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`

	// 更新后的日志集保存时间。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 待添加的推送任务实体列表。
	DropEntityList []*string `json:"DropEntityList,omitempty" name:"DropEntityList"`

	// 待删除的推送任务实例列表。
	AddedEntityList []*string `json:"AddedEntityList,omitempty" name:"AddedEntityList"`
}

func (r *ModifyLogTopicTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogTopicTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "LogSetRegion")
	delete(f, "LogSetId")
	delete(f, "TopicId")
	delete(f, "EntityType")
	delete(f, "TaskName")
	delete(f, "TopicName")
	delete(f, "LogSetName")
	delete(f, "Period")
	delete(f, "DropEntityList")
	delete(f, "AddedEntityList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLogTopicTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLogTopicTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLogTopicTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLogTopicTaskResponseParams `json:"Response"`
}

func (r *ModifyLogTopicTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogTopicTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOriginGroupRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 源站组ID。
	OriginGroupId *string `json:"OriginGroupId,omitempty" name:"OriginGroupId"`

	// 源站类型，取值有：
	// <li>self：自有源站；</li>
	// <li>third_party：第三方源站；</li>
	// <li>cos：腾讯云COS源站。</li>
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 源站组名称。
	OriginGroupName *string `json:"OriginGroupName,omitempty" name:"OriginGroupName"`

	// 源站配置类型，当OriginType=self时，取值有：
	// <li>area：按区域配置；</li>
	// <li>weight： 按权重配置；</li>
	// <li>proto： 按HTTP协议配置。</li>当OriginType=third_party/cos时放空。
	ConfigurationType *string `json:"ConfigurationType,omitempty" name:"ConfigurationType"`

	// 源站记录信息。
	OriginRecords []*OriginRecord `json:"OriginRecords,omitempty" name:"OriginRecords"`
}

type ModifyOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 源站组ID。
	OriginGroupId *string `json:"OriginGroupId,omitempty" name:"OriginGroupId"`

	// 源站类型，取值有：
	// <li>self：自有源站；</li>
	// <li>third_party：第三方源站；</li>
	// <li>cos：腾讯云COS源站。</li>
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 源站组名称。
	OriginGroupName *string `json:"OriginGroupName,omitempty" name:"OriginGroupName"`

	// 源站配置类型，当OriginType=self时，取值有：
	// <li>area：按区域配置；</li>
	// <li>weight： 按权重配置；</li>
	// <li>proto： 按HTTP协议配置。</li>当OriginType=third_party/cos时放空。
	ConfigurationType *string `json:"ConfigurationType,omitempty" name:"ConfigurationType"`

	// 源站记录信息。
	OriginRecords []*OriginRecord `json:"OriginRecords,omitempty" name:"OriginRecords"`
}

func (r *ModifyOriginGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOriginGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "OriginGroupId")
	delete(f, "OriginType")
	delete(f, "OriginGroupName")
	delete(f, "ConfigurationType")
	delete(f, "OriginRecords")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOriginGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyOriginGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyOriginGroupResponseParams `json:"Response"`
}

func (r *ModifyOriginGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOriginGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRulePriorityRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 规则 ID 的顺序，多条规则执行顺序依次往下。
	RuleIds []*string `json:"RuleIds,omitempty" name:"RuleIds"`
}

type ModifyRulePriorityRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 规则 ID 的顺序，多条规则执行顺序依次往下。
	RuleIds []*string `json:"RuleIds,omitempty" name:"RuleIds"`
}

func (r *ModifyRulePriorityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRulePriorityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRulePriorityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRulePriorityResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRulePriorityResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRulePriorityResponseParams `json:"Response"`
}

func (r *ModifyRulePriorityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRulePriorityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 规则名称，字符串名称长度 1~255。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则内容。
	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`

	// 规则 ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 规则状态，取值有：
	// <li> enable: 启用； </li>
	// <li> disable: 未启用。</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 规则名称，字符串名称长度 1~255。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则内容。
	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`

	// 规则 ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 规则状态，取值有：
	// <li> enable: 启用； </li>
	// <li> disable: 未启用。</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RuleName")
	delete(f, "Rules")
	delete(f, "RuleId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleResponseParams struct {
	// 规则 ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRuleResponseParams `json:"Response"`
}

func (r *ModifyRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityPolicyRequestParams struct {
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 安全配置。
	SecurityConfig *SecurityConfig `json:"SecurityConfig,omitempty" name:"SecurityConfig"`
}

type ModifySecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 安全配置。
	SecurityConfig *SecurityConfig `json:"SecurityConfig,omitempty" name:"SecurityConfig"`
}

func (r *ModifySecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entity")
	delete(f, "SecurityConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityPolicyResponseParams `json:"Response"`
}

func (r *ModifySecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityWafGroupPolicyRequestParams struct {
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 总开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>不填默认为上次的配置。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 规则等级，取值有：
	// <li> loose：宽松；</li>
	// <li> normal：正常；</li>
	// <li> strict：严格；</li>
	// <li> stricter：超严格；</li>
	// <li> custom：自定义。</li>不填默认为上次的配置。
	Level *string `json:"Level,omitempty" name:"Level"`

	// 处置方式，取值有：
	// <li> block：阻断；</li>
	// <li> observe：观察。</li>不填默认为上次的配置。
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 托管规则。不填默认为上次的配置。
	WafRules *WafRule `json:"WafRules,omitempty" name:"WafRules"`

	// AI引擎模式。不填默认为上次的配置。
	AiRule *AiRule `json:"AiRule,omitempty" name:"AiRule"`

	// 托管规则等级组。不填默认为上次的配置。
	WafGroups []*WafGroup `json:"WafGroups,omitempty" name:"WafGroups"`
}

type ModifySecurityWafGroupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 总开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>不填默认为上次的配置。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 规则等级，取值有：
	// <li> loose：宽松；</li>
	// <li> normal：正常；</li>
	// <li> strict：严格；</li>
	// <li> stricter：超严格；</li>
	// <li> custom：自定义。</li>不填默认为上次的配置。
	Level *string `json:"Level,omitempty" name:"Level"`

	// 处置方式，取值有：
	// <li> block：阻断；</li>
	// <li> observe：观察。</li>不填默认为上次的配置。
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 托管规则。不填默认为上次的配置。
	WafRules *WafRule `json:"WafRules,omitempty" name:"WafRules"`

	// AI引擎模式。不填默认为上次的配置。
	AiRule *AiRule `json:"AiRule,omitempty" name:"AiRule"`

	// 托管规则等级组。不填默认为上次的配置。
	WafGroups []*WafGroup `json:"WafGroups,omitempty" name:"WafGroups"`
}

func (r *ModifySecurityWafGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityWafGroupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entity")
	delete(f, "Switch")
	delete(f, "Level")
	delete(f, "Mode")
	delete(f, "WafRules")
	delete(f, "AiRule")
	delete(f, "WafGroups")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityWafGroupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityWafGroupPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySecurityWafGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityWafGroupPolicyResponseParams `json:"Response"`
}

func (r *ModifySecurityWafGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityWafGroupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneCnameSpeedUpRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// CNAME 加速状态，取值有：
	// <li> enabled：开启；</li>
	// <li> disabled：关闭。</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyZoneCnameSpeedUpRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// CNAME 加速状态，取值有：
	// <li> enabled：开启；</li>
	// <li> disabled：关闭。</li>
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
	delete(f, "ZoneId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneCnameSpeedUpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneCnameSpeedUpResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyZoneCnameSpeedUpResponse struct {
	*tchttp.BaseResponse
	Response *ModifyZoneCnameSpeedUpResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyZoneRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点接入方式，取值有：
	// <li> full：NS 接入；</li>
	// <li> partial：CNAME 接入。</li>不填写保持原有配置。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 自定义站点信息，以替代系统默认分配的名称服务器。不填写保持原有配置。
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitempty" name:"VanityNameServers"`
}

type ModifyZoneRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点接入方式，取值有：
	// <li> full：NS 接入；</li>
	// <li> partial：CNAME 接入。</li>不填写保持原有配置。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 自定义站点信息，以替代系统默认分配的名称服务器。不填写保持原有配置。
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
	delete(f, "ZoneId")
	delete(f, "Type")
	delete(f, "VanityNameServers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyZoneResponse struct {
	*tchttp.BaseResponse
	Response *ModifyZoneResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyZoneSettingRequestParams struct {
	// 待变更的站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 缓存过期时间配置。
	// 不填写表示保持原有配置。
	CacheConfig *CacheConfig `json:"CacheConfig,omitempty" name:"CacheConfig"`

	// 节点缓存键配置。
	// 不填写表示保持原有配置。
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// 浏览器缓存配置。
	// 不填写表示保持原有配置。
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// 离线缓存配置。
	// 不填写表示保持原有配置。
	OfflineCache *OfflineCache `json:"OfflineCache,omitempty" name:"OfflineCache"`

	// Quic访问配置。
	// 不填写表示保持原有配置。
	Quic *Quic `json:"Quic,omitempty" name:"Quic"`

	// Post请求传输配置。
	// 不填写表示保持原有配置。
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitempty" name:"PostMaxSize"`

	// 智能压缩配置。
	// 不填写表示保持原有配置。
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// Http2回源配置。
	// 不填写表示保持原有配置。
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitempty" name:"UpstreamHttp2"`

	// 访问协议强制Https跳转配置。
	// 不填写表示保持原有配置。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// Https加速配置。
	// 不填写表示保持原有配置。
	Https *Https `json:"Https,omitempty" name:"Https"`

	// 源站配置。
	// 不填写表示保持原有配置。
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// 智能加速配置。
	// 不填写表示保持原有配置。
	SmartRouting *SmartRouting `json:"SmartRouting,omitempty" name:"SmartRouting"`

	// WebSocket配置。
	// 不填写表示保持原有配置。
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`

	// 客户端IP回源请求头配置。
	// 不填写表示保持原有配置。
	ClientIpHeader *ClientIpHeader `json:"ClientIpHeader,omitempty" name:"ClientIpHeader"`

	// 缓存预刷新配置。
	// 不填写表示保持原有配置。
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitempty" name:"CachePrefresh"`

	// Ipv6访问配置。
	// 不填写表示保持原有配置。
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`
}

type ModifyZoneSettingRequest struct {
	*tchttp.BaseRequest
	
	// 待变更的站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 缓存过期时间配置。
	// 不填写表示保持原有配置。
	CacheConfig *CacheConfig `json:"CacheConfig,omitempty" name:"CacheConfig"`

	// 节点缓存键配置。
	// 不填写表示保持原有配置。
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// 浏览器缓存配置。
	// 不填写表示保持原有配置。
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// 离线缓存配置。
	// 不填写表示保持原有配置。
	OfflineCache *OfflineCache `json:"OfflineCache,omitempty" name:"OfflineCache"`

	// Quic访问配置。
	// 不填写表示保持原有配置。
	Quic *Quic `json:"Quic,omitempty" name:"Quic"`

	// Post请求传输配置。
	// 不填写表示保持原有配置。
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitempty" name:"PostMaxSize"`

	// 智能压缩配置。
	// 不填写表示保持原有配置。
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// Http2回源配置。
	// 不填写表示保持原有配置。
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitempty" name:"UpstreamHttp2"`

	// 访问协议强制Https跳转配置。
	// 不填写表示保持原有配置。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// Https加速配置。
	// 不填写表示保持原有配置。
	Https *Https `json:"Https,omitempty" name:"Https"`

	// 源站配置。
	// 不填写表示保持原有配置。
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// 智能加速配置。
	// 不填写表示保持原有配置。
	SmartRouting *SmartRouting `json:"SmartRouting,omitempty" name:"SmartRouting"`

	// WebSocket配置。
	// 不填写表示保持原有配置。
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`

	// 客户端IP回源请求头配置。
	// 不填写表示保持原有配置。
	ClientIpHeader *ClientIpHeader `json:"ClientIpHeader,omitempty" name:"ClientIpHeader"`

	// 缓存预刷新配置。
	// 不填写表示保持原有配置。
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitempty" name:"CachePrefresh"`

	// Ipv6访问配置。
	// 不填写表示保持原有配置。
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`
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
	delete(f, "CacheConfig")
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
	delete(f, "CachePrefresh")
	delete(f, "Ipv6")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneSettingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyZoneSettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyZoneSettingResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyZoneStatusRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点状态，取值有：
	// <li> false：开启站点；</li>
	// <li> true：关闭站点。</li>
	Paused *bool `json:"Paused,omitempty" name:"Paused"`
}

type ModifyZoneStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点状态，取值有：
	// <li> false：开启站点；</li>
	// <li> true：关闭站点。</li>
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
	delete(f, "ZoneId")
	delete(f, "Paused")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyZoneStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyZoneStatusResponseParams `json:"Response"`
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

type NoCache struct {
	// 不缓存配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type NormalAction struct {
	// 功能名称，功能名称填写规范可调用接口 [查询规则引擎的设置参数](https://tcloud4api.woa.com/document/product/1657/79433?!preview&!document=1) 查看。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 参数。
	Parameters []*RuleNormalActionParams `json:"Parameters,omitempty" name:"Parameters"`
}

type OfflineCache struct {
	// 离线缓存是否开启，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type OptimizeAction struct {
	// 站点性能优化配置项，取值有：
	// <li>Http2；</li>
	// <li>Http3；</li>
	// <li>Brotli。</li>
	Name *string `json:"Name,omitempty" name:"Name"`

	// 网络环境。
	Connectivity *string `json:"Connectivity,omitempty" name:"Connectivity"`

	// 开启配置项后，预估性能优化效果，单位ms。
	Value *int64 `json:"Value,omitempty" name:"Value"`

	// 开启配置项后，预估性能提升比例，单位%。
	Ratio *int64 `json:"Ratio,omitempty" name:"Ratio"`
}

type Origin struct {
	// 主源站列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Origins []*string `json:"Origins,omitempty" name:"Origins"`

	// 备源站列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupOrigins []*string `json:"BackupOrigins,omitempty" name:"BackupOrigins"`

	// 回源协议配置，取值有：
	// <li>http：强制 http 回源；</li>
	// <li>follow：协议跟随回源；</li>
	// <li>https：强制 https 回源，https 回源时仅支持源站 443 端口。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginPullProtocol *string `json:"OriginPullProtocol,omitempty" name:"OriginPullProtocol"`

	// OriginType 为对象存储（COS）时，可以指定是否允许访问私有 bucket。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPrivateAccess *string `json:"CosPrivateAccess,omitempty" name:"CosPrivateAccess"`
}

type OriginGroup struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 源站组ID。
	OriginGroupId *string `json:"OriginGroupId,omitempty" name:"OriginGroupId"`

	// 源站类型，取值有：
	// <li>self：自有源站；</li>
	// <li>third_party：第三方源站；</li>
	// <li>cos：腾讯云COS源站。</li>
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 源站组名称。
	OriginGroupName *string `json:"OriginGroupName,omitempty" name:"OriginGroupName"`

	// 源站配置类型，当OriginType=self时，取值有：
	// <li>area：按区域配置；</li>
	// <li>weight： 按权重配置。</li>
	// <li>proto： 按HTTP协议配置。</li>当OriginType=third_party/cos时放空。
	ConfigurationType *string `json:"ConfigurationType,omitempty" name:"ConfigurationType"`

	// 源站记录信息。
	OriginRecords []*OriginRecord `json:"OriginRecords,omitempty" name:"OriginRecords"`

	// 源站组更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type OriginRecord struct {
	// 源站记录值，不包含端口信息，可以为：IPv4，IPv6，域名格式。
	Record *string `json:"Record,omitempty" name:"Record"`

	// 源站记录ID。
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// 源站端口，取值范围：[1-65535]。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 当源站配置类型ConfigurationType=weight时，表示权重。
	// 不配置权重信息时，所有源站组记录统一填写为0或者不填写，表示多个源站轮询回源。
	// 配置权重信息时，取值为[1-100]，多个源站权重总和应为100，表示多个源站按照权重回源。
	// 当源站配置类型ConfigurationType=proto时，表示权重。
	// 不配置权重信息时，所有源站组记录统一填写为0或者不填写，表示多个源站轮询回源。
	// 配置权重信息时，取值为[1-100]，源站组内Proto相同的多个源站权重总和应为100，表示多个源站按照权重回源。
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// 当源站配置类型ConfigurationType=proto时，表示源站的协议类型，将按照客户端请求协议回到相应的源站，取值有：
	// <li>http：HTTP协议源站；</li>
	// <li>https：HTTPS协议源站。</li>
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// 当源站配置类型ConfigurationType=area时，表示区域，为空表示全部地区。取值为iso-3166中alpha-2编码或者大洲区域代码。大洲区域代码取值为：
	// <li>Asia：亚洲；</li>
	// <li>Europe：欧洲；</li>
	// <li>Africa：非洲；</li>
	// <li>Oceania：大洋洲；</li>
	// <li>Americas：美洲。</li>源站组记录中，至少需要有一项为全部地区。
	Area []*string `json:"Area,omitempty" name:"Area"`

	// 当源站类型OriginType=third_part时有效
	// 是否私有鉴权，取值有：
	// <li>true：使用私有鉴权；</li>
	// <li>false：不使用私有鉴权。</li>不填写，默认值为：false。
	Private *bool `json:"Private,omitempty" name:"Private"`

	// 当源站类型Private=true时有效，表示私有鉴权使用参数。
	PrivateParameters []*PrivateParameter `json:"PrivateParameters,omitempty" name:"PrivateParameters"`
}

type PlanInfo struct {
	// 结算货币类型，取值有：
	// <li> CNY ：人民币结算； </li>
	// <li> USD ：美元结算。</li>
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 套餐所含流量，该流量数值为安全加速流量，内容加速流量和智能加速流量的总和（单位：字节）。
	Flux *uint64 `json:"Flux,omitempty" name:"Flux"`

	// 结算周期，取值有：
	// <li> y ：按年结算； </li>
	// <li> m ：按月结算；</li>
	// <li> h ：按小时结算； </li>
	// <li> M ：按分钟结算；</li>
	// <li> s ：按秒结算。 </li>
	Frequency *string `json:"Frequency,omitempty" name:"Frequency"`

	// 套餐类型，取值有：
	// <li> sta ：全球内容分发网络（不包括中国大陆）标准版套餐； </li>
	// <li> sta_with_bot ：全球内容分发网络（不包括中国大陆）标准版套餐附带bot管理；</li>
	// <li> sta_cm ：中国大陆内容分发网络标准版套餐； </li>
	// <li> sta_cm_with_bot ：中国大陆内容分发网络标准版套餐附带bot管理；</li>
	// <li> ent ：全球内容分发网络（不包括中国大陆）企业版套餐； </li>
	// <li> ent_with_bot ： 全球内容分发网络（不包括中国大陆）企业版套餐附带bot管理；</li>
	// <li> ent_cm ：中国大陆内容分发网络企业版套餐； </li>
	// <li> ent_cm_with_bot ：中国大陆内容分发网络企业版套餐附带bot管理。</li>
	PlanType *string `json:"PlanType,omitempty" name:"PlanType"`

	// 套餐价格（单位：分）。
	Price *float64 `json:"Price,omitempty" name:"Price"`

	// 套餐所含请求次数，该请求次数为安全加速请求次数。（单位：次）。
	Request *uint64 `json:"Request,omitempty" name:"Request"`

	// 套餐所能绑定的站点个数。
	SiteNumber *uint64 `json:"SiteNumber,omitempty" name:"SiteNumber"`

	// 套餐加速区域类型，取值有：
	// <li> mainland ：中国大陆； </li>
	// <li> overseas ：全球（不包括中国大陆）。</li>
	Area *string `json:"Area,omitempty" name:"Area"`
}

type PortraitManagedRuleDetail struct {
	// 规则唯一id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则的描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 规则所属类型的名字, botdb(用户画像)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleTypeName *string `json:"RuleTypeName,omitempty" name:"RuleTypeName"`

	// 规则内的功能分类Id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassificationId *int64 `json:"ClassificationId,omitempty" name:"ClassificationId"`

	// 规则当前所属动作状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type PostMaxSize struct {
	// 是否开启POST请求上传文件限制，平台默认为限制为32MB，取值有：
	// <li>on：开启限制；</li>
	// <li>off：关闭限制。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 最大限制，取值在1MB和500MB之间。单位字节。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

type PrivateParameter struct {
	// 私有鉴权参数名称，取值有：
	// <li>AccessKeyId：鉴权参数Access Key ID；</li>
	// <li>SecretAccessKey：鉴权参数Secret Access Key。</li>
	Name *string `json:"Name,omitempty" name:"Name"`

	// 私有鉴权参数值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type QueryCondition struct {
	// 筛选条件的key。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 查询条件操作符，操作类型有：
	// <li>equals: 等于；</li>
	// <li>notEquals: 不等于；</li>
	// <li>include: 包含；</li>
	// <li>notInclude: 不包含; </li>
	// <li>startWith: 开始于；</li>
	// <li>notStartWith: 不开始于；</li>
	// <li>endWith: 结尾是；</li>
	// <li>notEndWith: 不结尾是。</li>
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 筛选条件的值。
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type QueryString struct {
	// CacheKey是否由QueryString组成，取值有：
	// <li>on：是；</li>
	// <li>off：否。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// CacheKey使用QueryString的方式，取值有：
	// <li>includeCustom：使用部分url参数；</li>
	// <li>excludeCustom：排除部分url参数。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 使用/排除的url参数数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type Quic struct {
	// 是否开启Quic配置，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Quota struct {
	// 单次批量提交配额上限。
	Batch *int64 `json:"Batch,omitempty" name:"Batch"`

	// 每日提交配额上限。
	Daily *int64 `json:"Daily,omitempty" name:"Daily"`

	// 每日剩余的可提交配额。
	DailyAvailable *int64 `json:"DailyAvailable,omitempty" name:"DailyAvailable"`

	// 配额类型，取值有：
	// <li> purge_prefix：前缀；</li>
	// <li> purge_url：URL；</li>
	// <li> purge_host：Hostname；</li>
	// <li> purge_all：全部缓存内容。</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

type RateLimitConfig struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 速率限制-用户规则列表。如果为null，默认使用历史配置。
	RateLimitUserRules []*RateLimitUserRule `json:"RateLimitUserRules,omitempty" name:"RateLimitUserRules"`

	// 速率限制模板功能。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLimitTemplate *RateLimitTemplate `json:"RateLimitTemplate,omitempty" name:"RateLimitTemplate"`

	// 智能客户端过滤。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLimitIntelligence *RateLimitIntelligence `json:"RateLimitIntelligence,omitempty" name:"RateLimitIntelligence"`
}

type RateLimitIntelligence struct {
	// 功能开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 执行动作，取值有：
	// <li>monitor：观察；</li>
	// <li>alg：挑战。</li>
	Action *string `json:"Action,omitempty" name:"Action"`
}

type RateLimitIntelligenceRuleDetail struct {
	// 智能识别到的客户端IP。
	MatchContent *string `json:"MatchContent,omitempty" name:"MatchContent"`

	// 应用的动作。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 更新时间。
	EffectiveTime *string `json:"EffectiveTime,omitempty" name:"EffectiveTime"`

	// 失效时间。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 规则id。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 处置状态，allowed即已经人为放行。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type RateLimitTemplate struct {
	// 模板等级名称，取值有：
	// <li>sup_loose：超级宽松；</li>
	// <li>loose：宽松；</li>
	// <li>emergency：紧急；</li>
	// <li>normal：适中；</li>
	// <li>strict：严格；</li>
	// <li>close：关闭 - 仅精准速率限制生效。</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 模板值详情。仅出参返回。
	RateLimitTemplateDetail *RateLimitTemplateDetail `json:"RateLimitTemplateDetail,omitempty" name:"RateLimitTemplateDetail"`
}

type RateLimitTemplateDetail struct {
	// 模板名称，取值有：
	// <li>sup_loose：超级宽松；</li>
	// <li>loose：宽松；</li>
	// <li>emergency：紧急；</li>
	// <li>normal：适中；</li>
	// <li>strict：严格；</li>
	// <li>close：关闭 - 仅精准速率限制生效。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 唯一id。
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 处置动作。模板阀值触发后的处罚行为。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 惩罚时间，0-2天，单位是秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PunishTime *int64 `json:"PunishTime,omitempty" name:"PunishTime"`

	// 统计阈值，单位是次，取值范围0-4294967294。
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`

	// 统计周期，取值范围0-120秒。
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

type RateLimitUserRule struct {
	// 速率限制统计阈值，单位是次，取值范围0-4294967294。
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`

	// 速率限制统计时间，取值范围 10/20/30/40/50/60 单位是秒。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 规则名，只能以英文字符，数字，下划线组合，且不能以下划线开头。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 处置动作，取值有：
	// <li>monitor：观察；</li>
	// <li>drop：拦截；</li>
	// <li>alg：JavaScript挑战。</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// 惩罚时长，0-2天。
	PunishTime *int64 `json:"PunishTime,omitempty" name:"PunishTime"`

	// 处罚时长单位，取值有：
	// <li>second：秒；</li>
	// <li>minutes：分钟；</li>
	// <li>hour：小时。</li>
	PunishTimeUnit *string `json:"PunishTimeUnit,omitempty" name:"PunishTimeUnit"`

	// 规则状态，取值有：
	// <li>on：生效；</li>
	// <li>off：不生效。</li>
	// <li>hour：小时。</li>默认on生效。
	RuleStatus *string `json:"RuleStatus,omitempty" name:"RuleStatus"`

	// 规则详情。
	AclConditions []*AclCondition `json:"AclConditions,omitempty" name:"AclConditions"`

	// 规则权重，取值范围0-100。
	RulePriority *int64 `json:"RulePriority,omitempty" name:"RulePriority"`

	// 规则id。仅出参使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// 过滤词，取值有：
	// <li>host：域名；</li>
	// <li>sip：客户端ip。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreqFields []*string `json:"FreqFields,omitempty" name:"FreqFields"`

	// 更新时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type ReclaimZoneRequestParams struct {
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type ReclaimZoneRequest struct {
	*tchttp.BaseRequest
	
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
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
	delete(f, "ZoneName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReclaimZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReclaimZoneResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReclaimZoneResponse struct {
	*tchttp.BaseResponse
	Response *ReclaimZoneResponseParams `json:"Response"`
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

type Resource struct {
	// 资源 ID。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 付费模式，取值有：
	// <li>0：后付费。</li>
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 生效时间。
	EnableTime *string `json:"EnableTime,omitempty" name:"EnableTime"`

	// 失效时间。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 套餐状态，取值有：
	// <li>normal：正常；</li>
	// <li>isolated：隔离；</li>
	// <li>destroyed：销毁。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 询价参数。
	Sv []*Sv `json:"Sv,omitempty" name:"Sv"`

	// 是否自动续费，取值有：
	// <li>0：默认状态；</li>
	// <li>1：自动续费；</li>
	// <li>2：不自动续费。</li>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 套餐关联资源 ID。
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 地域，取值有：
	// <li>mainland：国内；</li>
	// <li>overseas：海外。</li>
	Area *string `json:"Area,omitempty" name:"Area"`
}

type RewriteAction struct {
	// 功能名称，功能名称填写规范可调用接口 [查询规则引擎的设置参数](https://tcloud4api.woa.com/document/product/1657/79433?!preview&!document=1) 查看。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 参数。
	Parameters []*RuleRewriteActionParams `json:"Parameters,omitempty" name:"Parameters"`
}

type Rule struct {
	// 执行功能判断条件。
	// 注意：满足该数组内任意一项条件，功能即可执行。
	Conditions []*RuleAndConditions `json:"Conditions,omitempty" name:"Conditions"`

	// 执行的功能。
	Actions []*Action `json:"Actions,omitempty" name:"Actions"`
}

type RuleAndConditions struct {
	// 规则引擎条件，该数组内所有项全部满足即判断该条件满足。
	Conditions []*RuleCondition `json:"Conditions,omitempty" name:"Conditions"`
}

type RuleChoicePropertiesItem struct {
	// 参数名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数值类型。
	// <li> CHOICE：参数值只能在 ChoicesValue 中选择； </li>
	// <li> TOGGLE：参数值为开关类型，可在 ChoicesValue 中选择；</li>
	// <li> CUSTOM_NUM：参数值用户自定义，整型类型；</li>
	// <li> CUSTOM_STRING：参数值用户自定义，字符串类型。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 参数值的可选值。
	// 注意：若参数值为用户自定义则该数组为空数组。
	ChoicesValue []*string `json:"ChoicesValue,omitempty" name:"ChoicesValue"`

	// 数值参数的最小值，非数值参数或 Min 和 Max 值都为 0 则此项无意义。
	Min *int64 `json:"Min,omitempty" name:"Min"`

	// 数值参数的最大值，非数值参数或 Min 和 Max 值都为 0 则此项无意义。
	Max *int64 `json:"Max,omitempty" name:"Max"`

	// 参数值是否支持多选或者填写多个。
	IsMultiple *bool `json:"IsMultiple,omitempty" name:"IsMultiple"`

	// 是否允许为空。
	IsAllowEmpty *bool `json:"IsAllowEmpty,omitempty" name:"IsAllowEmpty"`

	// 特殊参数。
	// <li> 为 NULL：RuleAction 选择 NormalAction；</li>
	// <li> 成员参数 Id 为 Action：RuleAction 选择 RewirteAction；</li>
	// <li> 成员参数 Id 为 StatusCode：RuleAction 选择 CodeAction。</li>
	ExtraParameter *RuleExtraParameter `json:"ExtraParameter,omitempty" name:"ExtraParameter"`
}

type RuleCodeActionParams struct {
	// 状态 Code。
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 参数名称，参数填写规范可调用接口 [查询规则引擎的设置参数](https://tcloud4api.woa.com/document/product/1657/79433?!preview&!document=1) 查看。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type RuleCondition struct {
	// 运算符，取值有：
	// <li> equal: 等于； </li>
	// <li> notequal: 不等于。</li>
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 匹配类型，取值有：
	// <li> 全部（站点任意请求）: host。 </li>
	// <li> 文件名: filename； </li>
	// <li> 文件后缀: extension； </li>
	// <li> HOST: host； </li>
	// <li> URL Full: full_url，当前站点下完整 URL 路径，必须包含 HTTP 协议，Host 和 路径； </li>
	// <li> URL Path: url，当前站点下 URL 路径的请求。 </li>
	Target *string `json:"Target,omitempty" name:"Target"`

	// 对应匹配类型的参数值，对应匹配类型的取值有：
	// <li> 文件后缀：jpg、txt等文件后缀；</li>
	// <li> 文件名称：例如 foo.jpg 中的 foo；</li>
	// <li> 全部（站点任意请求）： all； </li>
	// <li> HOST：当前站点下的 host ，例如www.maxx55.com；</li>
	// <li> URL Path：当前站点下 URL 路径的请求，例如：/example；</li>
	// <li> URL Full：当前站点下完整 URL 请求，必须包含 HTTP 协议，Host 和 路径，例如：https://www.maxx55.cn/example。</li>
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type RuleExtraParameter struct {
	// 参数名，取值有：
	// <li> Action：修改 HTTP 头部所需参数，RuleAction 选择 RewirteAction；</li>
	// <li> StatusCode：状态码相关功能所需参数，RuleAction 选择 CodeAction。</li>
	Id *string `json:"Id,omitempty" name:"Id"`

	// 参数值类型。
	// <li> CHOICE：参数值只能在 Values 中选择； </li>
	// <li> CUSTOM_NUM：参数值用户自定义，整型类型；</li>
	// <li> CUSTOM_STRING：参数值用户自定义，字符串类型。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 可选参数值。
	// 注意：当 Id 的值为 StatusCode 时数组中的值为整型，填写参数值时请填写字符串的整型数值。
	Choices []*string `json:"Choices,omitempty" name:"Choices"`
}

type RuleItem struct {
	// 规则ID。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 规则名称，名称字符串长度 1~255。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则状态，取值有:
	// <li> enable: 启用； </li>
	// <li> disable: 未启用。 </li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 规则内容。
	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`

	// 规则优先级, 值越大优先级越高，最小为 1。
	RulePriority *int64 `json:"RulePriority,omitempty" name:"RulePriority"`
}

type RuleNormalActionParams struct {
	// 参数名称，参数填写规范可调用接口 [查询规则引擎的设置参数](https://tcloud4api.woa.com/document/product/1657/79433?!preview&!document=1) 查看。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type RuleRewriteActionParams struct {
	// 功能参数名称，参数填写规范可调用接口 [查询规则引擎的设置参数](https://tcloud4api.woa.com/document/product/1657/79433?!preview&!document=1) 查看。现在只有三种取值：
	// <li> add：添加 HTTP 头部；</li>
	// <li> set：重写 HTTP 头部；</li>
	// <li> del：删除 HTTP 头部。</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// 参数名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type RulesProperties struct {
	// 值为参数名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数值参数的最小值，非数值参数或 Min 和 Max 值都为 0 则此项无意义。
	Min *int64 `json:"Min,omitempty" name:"Min"`

	// 参数值的可选值。
	// 注意：若参数值为用户自定义则该数组为空数组。
	ChoicesValue []*string `json:"ChoicesValue,omitempty" name:"ChoicesValue"`

	// 参数值类型。
	// <li> CHOICE：参数值只能在 ChoicesValue 中选择； </li>
	// <li> TOGGLE：参数值为开关类型，可在 ChoicesValue 中选择；</li>
	// <li> OBJECT：参数值为对象类型，ChoiceProperties 为改对象类型关联的属性；</li>
	// <li> CUSTOM_NUM：参数值用户自定义，整型类型；</li>
	// <li> CUSTOM_STRING：参数值用户自定义，字符串类型。</li>注意：当参数类型为 OBJECT 类型时，请注意参考 [示例2 参数为 OBJECT 类型的创建](https://tcloud4api.woa.com/document/product/1657/79382?!preview&!document=1)
	Type *string `json:"Type,omitempty" name:"Type"`

	// 数值参数的最大值，非数值参数或 Min 和 Max 值都为 0 则此项无意义。
	Max *int64 `json:"Max,omitempty" name:"Max"`

	// 参数值是否支持多选或者填写多个。
	IsMultiple *bool `json:"IsMultiple,omitempty" name:"IsMultiple"`

	// 是否允许为空。
	IsAllowEmpty *bool `json:"IsAllowEmpty,omitempty" name:"IsAllowEmpty"`

	// 该参数对应的关联配置参数，属于调用接口的必填参数。
	// 注意：如果可选参数无特殊新增参数则该数组为空数组。
	ChoiceProperties []*RuleChoicePropertiesItem `json:"ChoiceProperties,omitempty" name:"ChoiceProperties"`

	// <li> 为 NULL：无特殊参数，RuleAction 选择 NormalAction；</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraParameter *RuleExtraParameter `json:"ExtraParameter,omitempty" name:"ExtraParameter"`
}

type RulesSettingAction struct {
	// 功能名称，取值有：
	// <li> 访问URL 重写（AccessUrlRedirect）；</li>
	// <li> 回源 URL 重写 （UpstreamUrlRedirect）；</li>
	// <li> 自定义错误页面
	// (ErrorPage)；</li>
	// <li> QUIC（QUIC）；</li>
	// <li> WebSocket （WebSocket）；</li>
	// <li> 视频拖拽（VideoSeek）；</li>
	// <li> Token 鉴权（Authentication）；</li>
	// <li> 自定义CacheKey（CacheKey）；</li>
	// <li> 节点缓存 TTL （Cache）；</li>
	// <li> 浏览器缓存 TTL（MaxAge）；</li>
	// <li> 离线缓存（OfflineCache）；</li>
	// <li> 智能加速（SmartRouting）；</li>
	// <li> 分片回源（RangeOriginPull）；</li>
	// <li> HTTP/2 回源（UpstreamHttp2）；</li>
	// <li> Host Header 重写（HostHeader）；</li>
	// <li> 强制 HTTPS（ForceRedirect）；</li>
	// <li> 回源 HTTPS（OriginPullProtocol）；</li>
	// <li> 缓存预刷新（CachePrefresh）；</li>
	// <li> 智能压缩（Compression）；</li>
	// <li> 修改 HTTP 请求头（RequestHeader）；</li>
	// <li> 修改HTTP响应头（ResponseHeader）;</li>
	// <li> 状态码缓存 TTL（StatusCodeCache）;</li>
	// <li> Hsts；</li>
	// <li> ClientIpHeader；</li>
	// <li> TlsVersion；</li>
	// <li> OcspStapling。</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// 参数信息。
	Properties []*RulesProperties `json:"Properties,omitempty" name:"Properties"`
}

type SecClientIp struct {
	// 客户端ip。
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 最大qps。
	RequestMaxQps *int64 `json:"RequestMaxQps,omitempty" name:"RequestMaxQps"`

	// 请求数。
	RequestNum *int64 `json:"RequestNum,omitempty" name:"RequestNum"`
}

type SecEntry struct {
	// 查询维度值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 查询维度下详细数据。
	Value []*SecEntryValue `json:"Value,omitempty" name:"Value"`
}

type SecEntryValue struct {
	// 指标名称。
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 时序数据详情。
	Detail []*TimingDataItem `json:"Detail,omitempty" name:"Detail"`

	// 最大值。
	Max *int64 `json:"Max,omitempty" name:"Max"`

	// 平均值。
	Avg *float64 `json:"Avg,omitempty" name:"Avg"`

	// 数据总和。
	Sum *float64 `json:"Sum,omitempty" name:"Sum"`
}

type SecHitRuleInfo struct {
	// 规则ID。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则类型名称。
	RuleTypeName *string `json:"RuleTypeName,omitempty" name:"RuleTypeName"`

	// 执行动作（处置方式），取值有：
	// <li>trans ：通过 ；</li>
	// <li>alg ：算法挑战 ；</li>
	// <li>drop ：丢弃 ；</li>
	// <li>ban ：封禁源ip ；</li>
	// <li>redirect ：重定向 ；</li>
	// <li>page ：返回指定页面 ；</li>
	// <li>monitor ：观察 。</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// 命中时间，采用unix秒级时间戳。
	HitTime *int64 `json:"HitTime,omitempty" name:"HitTime"`

	// 请求数。
	RequestNum *int64 `json:"RequestNum,omitempty" name:"RequestNum"`

	// 规则描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 子域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type SecRuleRelatedInfo struct {
	// 规则ID列表（99999为无效id）。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 执行动作（处置方式），取值有：
	// <li>trans ：通过 ；</li>
	// <li>alg ：算法挑战 ；</li>
	// <li>drop ：丢弃 ；</li>
	// <li>ban ：封禁源ip ；</li>
	// <li>redirect ：重定向 ；</li>
	// <li>page ：返回指定页面 ；</li>
	// <li>monitor ：观察 。</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// 风险等级（waf日志中独有），取值有：
	// <li>high risk ：高危 ；</li>
	// <li>middle risk ：中危 ；</li>
	// <li>low risk ：低危 ；</li>
	// <li>unkonw ：未知 。</li>
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 规则等级，取值有：
	// <li>normal  ：正常 。</li>
	RuleLevel *string `json:"RuleLevel,omitempty" name:"RuleLevel"`

	// 规则描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 规则类型名称。
	RuleTypeName *string `json:"RuleTypeName,omitempty" name:"RuleTypeName"`
}

type SecurityConfig struct {
	// 托管规则。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WafConfig *WafConfig `json:"WafConfig,omitempty" name:"WafConfig"`

	// 速率限制。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLimitConfig *RateLimitConfig `json:"RateLimitConfig,omitempty" name:"RateLimitConfig"`

	// 自定义规则。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclConfig *AclConfig `json:"AclConfig,omitempty" name:"AclConfig"`

	// Bot配置。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotConfig *BotConfig `json:"BotConfig,omitempty" name:"BotConfig"`

	// 七层防护总开关。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SwitchConfig *SwitchConfig `json:"SwitchConfig,omitempty" name:"SwitchConfig"`

	// 基础访问管控。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpTableConfig *IpTableConfig `json:"IpTableConfig,omitempty" name:"IpTableConfig"`

	// 例外规则配置。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptConfig *ExceptConfig `json:"ExceptConfig,omitempty" name:"ExceptConfig"`

	// 自定义拦截页面配置。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DropPageConfig *DropPageConfig `json:"DropPageConfig,omitempty" name:"DropPageConfig"`
}

type SecurityEntity struct {
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名/应用名。
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// 类型，取值有：
	// <li> domain：7层子域名； </li>
	// <li> application：4层应用名。 </li>
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`
}

type SecurityType struct {
	// 安全类型开关，取值为：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ServerCertInfo struct {
	// 服务器证书 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 证书备注名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 证书类型，取值有：
	// <li>default：默认证书；</lil>
	// <li>upload：用户上传；</li>
	// <li>managed:腾讯云托管。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 证书过期时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 证书部署时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployTime *string `json:"DeployTime,omitempty" name:"DeployTime"`

	// 签名算法。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignAlgo *string `json:"SignAlgo,omitempty" name:"SignAlgo"`
}

type ShieldArea struct {
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// DDoS策略Id。
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 防护类型，参考值：
	// <li>domain：7层子域；</li>
	// <li>application：4层应用。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 7层站点名。
	EntityName *string `json:"EntityName,omitempty" name:"EntityName"`

	// 该防护分区下的7层子域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DDoSHosts []*DDoSHost `json:"DDoSHosts,omitempty" name:"DDoSHosts"`

	// 四层tcp转发规则数。
	TcpNum *int64 `json:"TcpNum,omitempty" name:"TcpNum"`

	// 四层udp转发规则数。
	UdpNum *int64 `json:"UdpNum,omitempty" name:"UdpNum"`

	// 实例名称。
	Entity *string `json:"Entity,omitempty" name:"Entity"`
}

type SingleDataRecord struct {
	// 查询维度值。
	TypeKey *string `json:"TypeKey,omitempty" name:"TypeKey"`

	// 查询维度下具体指标值。
	TypeValue []*SingleTypeValue `json:"TypeValue,omitempty" name:"TypeValue"`
}

type SingleTypeValue struct {
	// 指标名。
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 指标值。
	DetailData *int64 `json:"DetailData,omitempty" name:"DetailData"`
}

type SmartRouting struct {
	// 智能加速配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Sort struct {
	// 排序字段，当前支持：
	// createTime，域名创建时间
	// certExpireTime，证书过期时间
	// certDeployTime,  证书部署时间
	Key *string `json:"Key,omitempty" name:"Key"`

	// asc/desc，默认desc。
	Sequence *string `json:"Sequence,omitempty" name:"Sequence"`
}

type SpeedTestingConfig struct {
	// 任务类型，取值有：
	// <li>1：页面性能;</li>
	// <li>2：文件上传;</li>
	// <li>3：文件下载;</li>
	// <li>4：端口性能;</li>
	// <li>5：网络质量;</li>
	// <li>6：音视频体验。</li>
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 拨测 url。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 拨测 UA。
	UA *string `json:"UA,omitempty" name:"UA"`

	// 网络类型。
	Connectivity *string `json:"Connectivity,omitempty" name:"Connectivity"`
}

type SpeedTestingDetailData struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 地域性能数据。
	DistrictStatistics []*DistrictStatistics `json:"DistrictStatistics,omitempty" name:"DistrictStatistics"`
}

type SpeedTestingInfo struct {
	// 任务状态，取值有：
	// <li> 200：任务完成;</li>
	// <li> 100：任务进行中；</li>
	// <li> 503: 任务失败。</li>
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 拨测任务 ID。
	TestId *string `json:"TestId,omitempty" name:"TestId"`

	// 拨测任务配置。
	SpeedTestingConfig *SpeedTestingConfig `json:"SpeedTestingConfig,omitempty" name:"SpeedTestingConfig"`

	// 拨测任务统计结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpeedTestingStatistics *SpeedTestingStatistics `json:"SpeedTestingStatistics,omitempty" name:"SpeedTestingStatistics"`
}

type SpeedTestingMetricData struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 源站拨测信息。
	OriginSpeedTestingInfo []*SpeedTestingInfo `json:"OriginSpeedTestingInfo,omitempty" name:"OriginSpeedTestingInfo"`

	// EO 拨测信息。
	ProxySpeedTestingInfo []*SpeedTestingInfo `json:"ProxySpeedTestingInfo,omitempty" name:"ProxySpeedTestingInfo"`

	// 站点状态。
	SpeedTestingStatus *SpeedTestingStatus `json:"SpeedTestingStatus,omitempty" name:"SpeedTestingStatus"`

	// 优化建议。
	OptimizeAction []*OptimizeAction `json:"OptimizeAction,omitempty" name:"OptimizeAction"`

	// EO 整体性能，单位ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyLoadTime *int64 `json:"ProxyLoadTime,omitempty" name:"ProxyLoadTime"`

	// 源站整体性能，单位ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginLoadTime *int64 `json:"OriginLoadTime,omitempty" name:"OriginLoadTime"`
}

type SpeedTestingQuota struct {
	// 站点总拨测次数。
	TotalTestRuns *int64 `json:"TotalTestRuns,omitempty" name:"TotalTestRuns"`

	// 站点剩余可用拨测次数。
	AvailableTestRuns *int64 `json:"AvailableTestRuns,omitempty" name:"AvailableTestRuns"`
}

type SpeedTestingStatistics struct {
	// 首屏时间，单位 ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstContentfulPaint *int64 `json:"FirstContentfulPaint,omitempty" name:"FirstContentfulPaint"`

	// 首屏完全渲染时间，单位 ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstMeaningfulPaint *int64 `json:"FirstMeaningfulPaint,omitempty" name:"FirstMeaningfulPaint"`

	// 整体下载速度，单位 KB/s。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OverallDownloadSpeed *float64 `json:"OverallDownloadSpeed,omitempty" name:"OverallDownloadSpeed"`

	// 渲染时间，单位 ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenderTime *int64 `json:"RenderTime,omitempty" name:"RenderTime"`

	// 文档完成时间, 单位 ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocumentFinishTime *int64 `json:"DocumentFinishTime,omitempty" name:"DocumentFinishTime"`

	// 基础文档TCP连接时间，单位 ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TcpConnectionTime *int64 `json:"TcpConnectionTime,omitempty" name:"TcpConnectionTime"`

	// 基础文档服务器响应时间，单位 ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseTime *int64 `json:"ResponseTime,omitempty" name:"ResponseTime"`

	// 基础文档下载时间，单位 ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileDownloadTime *int64 `json:"FileDownloadTime,omitempty" name:"FileDownloadTime"`

	// 整体性能，测试总时间，单位 ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadTime *int64 `json:"LoadTime,omitempty" name:"LoadTime"`
}

type SpeedTestingStatus struct {
	// 拨测 url。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 拨测 url 是否使用 https。
	Tls *bool `json:"Tls,omitempty" name:"Tls"`

	// 任务创建时间。
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 任务状态，取值有：
	// <li> 200：任务完成;</li>
	// <li> 100：任务进行中。</li>
	// <li> 503: 任务失败。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 拨测 UA。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UA *string `json:"UA,omitempty" name:"UA"`

	// 网络环境。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Connectivity *string `json:"Connectivity,omitempty" name:"Connectivity"`

	// 是否可达，取值：
	// <li> true：可达；</li>
	// <li> false：不可达。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reachable *bool `json:"Reachable,omitempty" name:"Reachable"`

	// 是否超时，取值：
	// <li> true：超时；</li>
	// <li> false：不超时。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimedOut *bool `json:"TimedOut,omitempty" name:"TimedOut"`
}

type Sv struct {
	// 询价参数键。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 询价参数值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type SwitchConfig struct {
	// Web类型的安全总开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>不影响DDoS与Bot的开关。
	WebSwitch *string `json:"WebSwitch,omitempty" name:"WebSwitch"`
}

// Predefined struct for user
type SwitchLogTopicTaskRequestParams struct {
	// 推送任务的主题ID。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 是否开启推送，可选的动作有：
	// <li>true：开启推送任务；</li>
	// <li>false：关闭推送任务。</li>
	IsOpen *bool `json:"IsOpen,omitempty" name:"IsOpen"`
}

type SwitchLogTopicTaskRequest struct {
	*tchttp.BaseRequest
	
	// 推送任务的主题ID。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 是否开启推送，可选的动作有：
	// <li>true：开启推送任务；</li>
	// <li>false：关闭推送任务。</li>
	IsOpen *bool `json:"IsOpen,omitempty" name:"IsOpen"`
}

func (r *SwitchLogTopicTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchLogTopicTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "IsOpen")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchLogTopicTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchLogTopicTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwitchLogTopicTaskResponse struct {
	*tchttp.BaseResponse
	Response *SwitchLogTopicTaskResponseParams `json:"Response"`
}

func (r *SwitchLogTopicTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchLogTopicTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签键。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type Task struct {
	// 任务 ID。
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 状态。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 资源。
	Target *string `json:"Target,omitempty" name:"Target"`

	// 任务类型。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务完成时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type TimingDataItem struct {
	// 返回数据对应时间点，采用unix秒级时间戳。
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 具体数值。
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type TimingDataRecord struct {
	// 查询维度值。
	TypeKey *string `json:"TypeKey,omitempty" name:"TypeKey"`

	// 详细时序数据。
	TypeValue []*TimingTypeValue `json:"TypeValue,omitempty" name:"TypeValue"`
}

type TimingTypeValue struct {
	// 数据和。
	Sum *int64 `json:"Sum,omitempty" name:"Sum"`

	// 最大值。
	Max *int64 `json:"Max,omitempty" name:"Max"`

	// 平均值。
	Avg *int64 `json:"Avg,omitempty" name:"Avg"`

	// 指标名。
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 详细数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*TimingDataItem `json:"Detail,omitempty" name:"Detail"`
}

type TopDataRecord struct {
	// 查询维度值。
	TypeKey *string `json:"TypeKey,omitempty" name:"TypeKey"`

	// top数据排行。
	DetailData []*TopDetailData `json:"DetailData,omitempty" name:"DetailData"`
}

type TopDetailData struct {
	// 字段名。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 字段值。
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type TopEntry struct {
	// top查询维度值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 查询具体数据。
	Value []*TopEntryValue `json:"Value,omitempty" name:"Value"`
}

type TopEntryValue struct {
	// 排序实体名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 排序实体数量。
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type UpstreamHttp2 struct {
	// http2回源配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type VanityNameServers struct {
	// 自定义 ns 开关，取值有：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 自定义 ns 列表。
	Servers []*string `json:"Servers,omitempty" name:"Servers"`
}

type VanityNameServersIps struct {
	// 自定义名字服务器名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 自定义名字服务器 IPv4 地址。
	IPv4 *string `json:"IPv4,omitempty" name:"IPv4"`
}

type Waf struct {
	// Waf开关，取值为：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

type WafConfig struct {
	// WafConfig开关，取值有：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>开关仅与配置是否生效有关，即使为off（关闭），也可以正常修改配置的内容。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 上一次设置的防护级别，取值有：
	// <li> loose：宽松；</li>
	// <li> normal：正常；</li>
	// <li> strict：严格；</li>
	// <li> stricter：超严格；</li>
	// <li> custom：自定义。</li>
	Level *string `json:"Level,omitempty" name:"Level"`

	// 全局WAF模式，取值有：
	// <li> block：阻断（全局阻断，但可对详细规则配置观察）；</li>
	// <li> observe：观察（无论详细规则配置什么，都为观察）。</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 托管规则详细配置。如果为null，默认使用历史配置。
	WafRule *WafRule `json:"WafRule,omitempty" name:"WafRule"`

	// AI规则引擎防护配置。如果为null，默认使用历史配置。
	AiRule *AiRule `json:"AiRule,omitempty" name:"AiRule"`
}

type WafGroup struct {
	// 执行动作，取值有：
	// <li> block：阻断；</li>
	// <li> observe：观察。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 防护级别，取值有：
	// <li> loose：宽松；</li>
	// <li> normal：正常；</li>
	// <li> strict：严格；</li>
	// <li> stricter：超严格；</li>
	// <li> custom：自定义。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *string `json:"Level,omitempty" name:"Level"`

	// 规则类型id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeId *int64 `json:"TypeId,omitempty" name:"TypeId"`
}

type WafGroupDetail struct {
	// 规则类型ID。
	RuleTypeId *int64 `json:"RuleTypeId,omitempty" name:"RuleTypeId"`

	// 规则类型名称。
	RuleTypeName *string `json:"RuleTypeName,omitempty" name:"RuleTypeName"`

	// 规则类型描述。
	RuleTypeDesc *string `json:"RuleTypeDesc,omitempty" name:"RuleTypeDesc"`

	// 规则列表。
	WafGroupRules []*WafGroupRule `json:"WafGroupRules,omitempty" name:"WafGroupRules"`

	// 规则等级。
	Level *string `json:"Level,omitempty" name:"Level"`

	// 动作。
	Action *string `json:"Action,omitempty" name:"Action"`
}

type WafGroupInfo struct {
	// 托管规则组列表。
	WafGroupDetails []*WafGroupDetail `json:"WafGroupDetails,omitempty" name:"WafGroupDetails"`

	// 规则组等级，取值有：
	// <li> loose：宽松；</li>
	// <li> normal：正常；</li>
	// <li> strict：严格；</li>
	// <li> stricter：超严格。</li>
	Level *string `json:"Level,omitempty" name:"Level"`

	// 保留字段。
	Act *string `json:"Act,omitempty" name:"Act"`

	// 模式，取值有：
	// <li> block：阻断；</li>
	// <li> observe：观察。</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 开关，取值有：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type WafGroupRule struct {
	// 规则id。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 等级描述。
	RuleLevelDesc *string `json:"RuleLevelDesc,omitempty" name:"RuleLevelDesc"`

	// 规则标签。部分类型的规则不存在该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleTags []*string `json:"RuleTags,omitempty" name:"RuleTags"`

	// 更新时间，格式为YYYY-MM-DD hh:mm:ss。部分类型的规则不存在该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 状态，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>为空时对应接口Status无意义，例如仅查询规则详情时。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 规则类型名。
	RuleTypeName *string `json:"RuleTypeName,omitempty" name:"RuleTypeName"`

	// 规则类型id。
	RuleTypeId *int64 `json:"RuleTypeId,omitempty" name:"RuleTypeId"`

	// 规则类型描述。
	RuleTypeDesc *string `json:"RuleTypeDesc,omitempty" name:"RuleTypeDesc"`
}

type WafRule struct {
	// 托管规则开关，取值有：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 黑名单，ID参考接口 [DescribeSecurityGroupManagedRules](https://tcloud4api.woa.com/document/product/1657/80807?!preview&!document=1)。
	BlockRuleIDs []*int64 `json:"BlockRuleIDs,omitempty" name:"BlockRuleIDs"`

	// 观察模式ID列表，将规则ID加入本参数列表中代表该ID使用观察模式生效，即该规则ID进入观察模式。ID参考接口 [DescribeSecurityGroupManagedRules](https://tcloud4api.woa.com/document/product/1657/80807?!preview&!document=1)。
	ObserveRuleIDs []*int64 `json:"ObserveRuleIDs,omitempty" name:"ObserveRuleIDs"`
}

type WebLogs struct {
	// 请求（事件）ID。
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 攻击源（客户端）Ip。
	AttackIp *string `json:"AttackIp,omitempty" name:"AttackIp"`

	// 受攻击子域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// http 日志内容。
	HttpLog *string `json:"HttpLog,omitempty" name:"HttpLog"`

	// IP所在国家iso-3166中alpha-2编码，编码信息请参考[ISO-3166](https://git.woa.com/edgeone/iso-3166/blob/master/all/all.json)
	SipCountryCode *string `json:"SipCountryCode,omitempty" name:"SipCountryCode"`

	// 攻击时间，采用unix秒级时间戳。
	AttackTime *uint64 `json:"AttackTime,omitempty" name:"AttackTime"`

	// 请求地址。
	RequestUri *string `json:"RequestUri,omitempty" name:"RequestUri"`

	// 攻击内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackContent *string `json:"AttackContent,omitempty" name:"AttackContent"`

	// 规则相关信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleDetailList []*SecRuleRelatedInfo `json:"RuleDetailList,omitempty" name:"RuleDetailList"`

	// 请求类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReqMethod *string `json:"ReqMethod,omitempty" name:"ReqMethod"`
}

type WebSocket struct {
	// WebSocket 超时时间配置开关，取值有：
	// <li>on：使用Timeout作为WebSocket超时时间；</li>
	// <li>off：平台仍支持WebSocket连接，此时使用系统默认的15秒为超时时间。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 超时时间，单位为秒，最大超时时间120秒。
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
}

type Zone struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 站点当前使用的 NS 列表。
	OriginalNameServers []*string `json:"OriginalNameServers,omitempty" name:"OriginalNameServers"`

	// 腾讯云分配的 NS 列表。
	NameServers []*string `json:"NameServers,omitempty" name:"NameServers"`

	// 站点状态，取值有：
	// <li> active：NS 已切换； </li>
	// <li> pending：NS 未切换；</li>
	// <li> moved：NS 已切走；</li>
	// <li> deactivated：被封禁。 </li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 站点接入方式，取值有
	// <li> full：NS 接入； </li>
	// <li> partial：CNAME 接入。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 站点是否关闭。
	Paused *bool `json:"Paused,omitempty" name:"Paused"`

	// 是否开启cname加速，取值有：
	// <li> enabled：开启；</li>
	// <li> disabled：关闭。</li>
	CnameSpeedUp *string `json:"CnameSpeedUp,omitempty" name:"CnameSpeedUp"`

	// cname 接入状态，取值有：
	// <li> finished：站点已验证；</li>
	// <li> pending：站点验证中。</li>
	CnameStatus *string `json:"CnameStatus,omitempty" name:"CnameStatus"`

	// 资源标签列表。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 计费资源列表。
	Resources []*Resource `json:"Resources,omitempty" name:"Resources"`

	// 站点创建时间。
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 站点修改时间。
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// 站点接入地域，取值有：
	// <li> global：全球；</li>
	// <li> mainland：中国大陆；</li>
	// <li> overseas：境外区域。</li>
	Area *string `json:"Area,omitempty" name:"Area"`

	// 用户自定义 NS 信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitempty" name:"VanityNameServers"`

	// 用户自定义 NS IP 信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VanityNameServersIps []*VanityNameServersIps `json:"VanityNameServersIps,omitempty" name:"VanityNameServersIps"`
}

type ZoneSetting struct {
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 站点加速区域信息，取值有：
	// <li> mainland：中国境内加速；</li>
	// <li> overseas：中国境外加速。</li>
	Area *string `json:"Area,omitempty" name:"Area"`

	// 节点缓存键配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// Quic访问配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quic *Quic `json:"Quic,omitempty" name:"Quic"`

	// POST请求传输配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitempty" name:"PostMaxSize"`

	// 智能压缩配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// Http2回源配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitempty" name:"UpstreamHttp2"`

	// 访问协议强制Https跳转配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// 缓存过期时间配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheConfig *CacheConfig `json:"CacheConfig,omitempty" name:"CacheConfig"`

	// 源站配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// 智能加速配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmartRouting *SmartRouting `json:"SmartRouting,omitempty" name:"SmartRouting"`

	// 浏览器缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// 离线缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineCache *OfflineCache `json:"OfflineCache,omitempty" name:"OfflineCache"`

	// WebSocket配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`

	// 客户端IP回源请求头配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIpHeader *ClientIpHeader `json:"ClientIpHeader,omitempty" name:"ClientIpHeader"`

	// 缓存预刷新配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitempty" name:"CachePrefresh"`

	// Ipv6访问配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// Https 加速配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Https *Https `json:"Https,omitempty" name:"Https"`
}